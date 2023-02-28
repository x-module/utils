package gamelift

import (
	"encoding/json"
	"fmt"
	"github.com/go-xmodule/module/global"
	"github.com/go-xmodule/utils/gamelift/pkg/proto/pbuffer"
	"github.com/go-xmodule/utils/gamelift/pkg/socketio"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

const (
	healthCheckTimeout = 60
)

type Handler interface {
	StartGameSession(event *pbuffer.ActivateGameSession)
	UpdateGameSession(event *pbuffer.UpdateGameSession)
	ProcessTerminate(event *pbuffer.TerminateProcess)
	HealthCheck() bool
}
type Config struct {
	Domain string
	Port   int
}
type Client interface {
	Handle(h Handler)

	Open() error
	ProcessReady(event *pbuffer.ProcessReady) error
	ProcessEnding(event *pbuffer.ProcessEnding) error
	ActivateGameSession(event *pbuffer.GameSessionActivate) error
	TerminateGameSession(event *pbuffer.GameSessionTerminate) error
	StartMatchBackfill(event *pbuffer.BackfillMatchmakingRequest) (*pbuffer.BackfillMatchmakingResponse, error)
	StopMatchBackfill(event *pbuffer.StopMatchmakingRequest) error
	UpdatePlayerSessionCreationPolicy(event *pbuffer.UpdatePlayerSessionCreationPolicy) error
	AcceptPlayerSession(event *pbuffer.AcceptPlayerSession) error
	RemovePlayerSession(event *pbuffer.RemovePlayerSession) error
	DescribePlayerSessions(event *pbuffer.DescribePlayerSessionsRequest) (*pbuffer.DescribePlayerSessionsResponse, error)

	GetGameSessionId() *string
	GetTerminationTime() *time.Time
}

type client struct {
	config               Config
	client               *socketio.Client
	handler              Handler
	isReady              bool
	logger               *logrus.Logger
	gameSessionID        *string
	processTerminateTime *time.Time
}

func NewClient(logger *logrus.Logger, config Config) Client {
	return &client{logger: logger, config: config}
}

func (c *client) Handle(handler Handler) {
	c.handler = handler
}

func (c *client) HandleReceivedMessage(str string, msg interface{}, p *socketio.Packet) error {
	err := json.Unmarshal([]byte(str), &msg)
	if err != nil {
		c.logger.WithField(global.ErrField, err).Error("failed to unmarshal Message")
		return err
	}
	ackPacket := socketio.NewAckPacket(p, []interface{}{true})
	if err := c.client.SendPacket(ackPacket); err != nil {
		c.logger.Error("failed to send ack packet", err)
	}
	return nil
}

func (c *client) Open() error {
	q := url.Values{}
	if ppid := os.Getenv("MAIN_PID"); ppid != "" {
		q.Set("pID", ppid)
	} else {
		q.Set("pID", fmt.Sprint(os.Getpid()))
	}
	q.Set("sdkVersion", "3.4.0")
	q.Set("sdkLanguage", "Go")
	socketUrl := fmt.Sprintf("ws://%s:%d/socket.io/?%s", c.config.Domain, c.config.Port, q.Encode())
	fmt.Println(socketUrl)
	c.client = socketio.NewClient(socketUrl, c.logger)
	c.client.HandleFunc(func(p *socketio.Packet) {
		name := string(p.Data[0].(json.RawMessage))
		var str string
		err := json.Unmarshal([]byte(p.Data[1].(json.RawMessage)), &str)
		if err != nil {
			c.logger.Error("failed to unmarshal packet name", err)
		}
		switch name {
		case `"StartGameSession"`:
			msg := &pbuffer.ActivateGameSession{}
			if err := c.HandleReceivedMessage(str, msg, p); err != nil {
				c.logger.Error("failed to parse received ActivateGameSession", err)
			}
			c.gameSessionID = stringAddr(msg.GetGameSession().GetGameSessionId())
			go c.handler.StartGameSession(msg)
		case `"UpdateGameSession"`:
			msg := &pbuffer.UpdateGameSession{}
			if err := c.HandleReceivedMessage(str, msg, p); err != nil {
				c.logger.Error("failed to parse received UpdateGameSession", err)
			}
			go c.handler.UpdateGameSession(msg)
		case `"TerminateProcess"`:
			msg := &pbuffer.TerminateProcess{}
			if err := c.HandleReceivedMessage(str, msg, p); err != nil {
				c.logger.Error("failed to parse received TerminateProcess", err)
			}
			c.processTerminateTime = timeAddr(time.Unix(msg.GetTerminationTime(), 0))
			go c.handler.ProcessTerminate(msg)
		default:
			c.logger.Info("unhandled packet", name)
		}
	})
	return c.client.Open()
}

func (c *client) ReportHealth() {
	// TODO: nonblocking
	health := c.handler.HealthCheck()
	event := &pbuffer.ReportHealth{HealthStatus: health}
	data, err := proto.Marshal(event)
	if err != nil {
		c.logger.Error("failed to marshal ReportHealth", err)
	}
	var rmsg []interface{}
	rmsg = append(rmsg, proto.MessageName(event), data)
	c.client.Send(rmsg)
}

type GenericError struct {
	pbuffer.GameLiftResponse
}

func (err *GenericError) Error() string {
	return fmt.Sprintf("%v:%v:%v", err.GetStatus().String(), err.GetErrorMessage(), err.GetResponseData())
}

func ParseGameLiftResponse(data []interface{}) error {
	var success bool
	if err := json.Unmarshal(data[0].(json.RawMessage), &success); err != nil {
		return err
	}
	if success {
		return nil
	}
	var str string
	err := json.Unmarshal([]byte(data[1].(json.RawMessage)), &str)
	if err != nil {
		return err
	}
	var msg pbuffer.GameLiftResponse
	if err := jsonpb.Unmarshal(strings.NewReader(str), &msg); err != nil {
		return err
	}
	return &GenericError{msg}
}

func (c *client) ProcessReady(event *pbuffer.ProcessReady) error {
	err := c.call(event)
	if err != nil {
		return err
	}

	c.isReady = true
	// wake healthcheck goroutine
	go func() {
		for c.isReady {
			c.ReportHealth()
			time.Sleep(time.Second * healthCheckTimeout)
		}
	}()
	return nil
}

func (c *client) call(event proto.Message) error {
	data, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	var rmsg []interface{}
	rmsg = append(rmsg, proto.MessageName(event), data)
	ack, err := c.client.SendAck(rmsg)
	if err != nil {
		return err
	}
	if err := ParseGameLiftResponse(ack); err != nil {
		return err
	}
	return nil
}

func (c *client) callReturn(event proto.Message, result proto.Message) error {
	data, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	var rmsg []interface{}
	rmsg = append(rmsg, proto.MessageName(event), data)
	ack, err := c.client.SendAck(rmsg)
	if err != nil {
		return err
	}
	if err := ParseGameLiftResponse(ack); err != nil {
		return err
	}
	var str string
	if err := json.Unmarshal(ack[1].(json.RawMessage), &str); err != nil {
		return err
	}
	if err := jsonpb.Unmarshal(strings.NewReader(str), result); err != nil {
		return err
	}
	return nil
}

func (c *client) ProcessEnding(event *pbuffer.ProcessEnding) error {
	return c.call(event)
}

func (c *client) ActivateGameSession(event *pbuffer.GameSessionActivate) error {
	return c.call(event)
}

func (c *client) TerminateGameSession(event *pbuffer.GameSessionTerminate) error {
	return c.call(event)
}

func (c *client) StartMatchBackfill(event *pbuffer.BackfillMatchmakingRequest) (*pbuffer.BackfillMatchmakingResponse, error) {
	result := &pbuffer.BackfillMatchmakingResponse{}
	return result, c.callReturn(event, result)
}

func (c *client) StopMatchBackfill(event *pbuffer.StopMatchmakingRequest) error {
	return c.call(event)
}

func (c *client) UpdatePlayerSessionCreationPolicy(event *pbuffer.UpdatePlayerSessionCreationPolicy) error {
	return c.call(event)
}

func (c *client) AcceptPlayerSession(event *pbuffer.AcceptPlayerSession) error {
	return c.call(event)
}

func (c *client) RemovePlayerSession(event *pbuffer.RemovePlayerSession) error {
	return c.call(event)
}

func (c *client) DescribePlayerSessions(event *pbuffer.DescribePlayerSessionsRequest) (*pbuffer.DescribePlayerSessionsResponse, error) {
	result := &pbuffer.DescribePlayerSessionsResponse{}
	return result, c.callReturn(event, result)
}

func (c *client) GetGameSessionId() *string {
	return c.gameSessionID
}

func (c *client) GetTerminationTime() *time.Time {
	return c.processTerminateTime
}

func stringAddr(s string) *string {
	return &s
}

func timeAddr(t time.Time) *time.Time {
	return &t
}

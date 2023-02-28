package socketio

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-xmodule/utils/gamelift/pkg/eventio"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

var (
	ErrorEmptyPacket = errors.New("packet length should be at least 1 byte")
	ErrorNullPacket  = errors.New("packet length should be at least 1")
)

type PacketType int

const (
	Connect PacketType = iota
	Disconnect
	Event
	Ack
	Error
	BinaryEvent
	BinaryAck
)

type HandlerFunc func(packet *Packet)

type Handler interface {
	HandleMessage(packet *Packet)
}

type nullHandler struct{}

func (pt PacketType) String() string {
	switch pt {
	case Connect:
		return "Connect"
	case Disconnect:
		return "Disconnect"
	case Event:
		return "Event"
	case Ack:
		return "Ack"
	case Error:
		return "Error"
	case BinaryEvent:
		return "BinaryEvent"
	case BinaryAck:
		return "BinaryAck"
	default:
		return "Unknown"
	}
}

type Packet struct {
	Type PacketType
	ID   *int
	Data []interface{}
}

func NewAckPacket(p *Packet, data []interface{}) Packet {
	return Packet{
		Type: Ack,
		ID:   p.ID,
		Data: data,
	}
}

func EncodePacket(p Packet) (string, error) {
	var (
		data []byte
		err  error
	)
	if len(p.Data) > 0 {
		data, err = json.Marshal(p.Data)
		if err != nil {
			return "", err
		}
	} else {
		data = []byte{}
	}
	var idStr string
	if p.ID != nil {
		idStr = fmt.Sprint(*p.ID)
	}
	return fmt.Sprintf("%d%v%v", p.Type, idStr, string(data)), nil
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func DecodePacket(data string) (Packet, error) {
	var p Packet
	if len(data) == 0 {
		return p, ErrorEmptyPacket
	}
	typ, err := strconv.Atoi(data[0:1])
	if err != nil {
		return p, err
	}
	p.Type = PacketType(typ)
	if p.Type != Event && p.Type != Ack && p.Type != Error {
		return p, nil
	}
	var i int
	for i = 1; i < len(data); i++ {
		if !isDigit(data[i]) {
			break
		}
	}
	// log.Println(data, i)
	if i > 1 {
		pid, err := strconv.Atoi(data[1:i])
		if err != nil {
			return p, err
		}
		p.ID = &pid
	}
	// log.Println(data, i, data[i:])
	var msgs []json.RawMessage
	if err := json.Unmarshal([]byte(data[i:]), &msgs); err != nil {
		return p, err
	}
	if len(msgs) == 0 {
		return p, ErrorNullPacket
	}
	// log.Println(msgs)
	for _, msg := range msgs {
		p.Data = append(p.Data, msg)
	}
	return p, nil
}

func (h nullHandler) HandleMessage(packet *Packet) {
}

func (f HandlerFunc) HandleMessage(packet *Packet) {
	f(packet)
}

type Client struct {
	client  *eventio.Client
	handler Handler
	reqId   int
	ackCh   map[int]chan []interface{}
	ackChMu sync.Mutex
	logger  *logrus.Logger
}

func NewClient(url string, logger *logrus.Logger) *Client {
	ec := eventio.NewClient(url, logger)
	client := &Client{
		client:  ec,
		handler: &nullHandler{},
		reqId:   10000,
		ackCh:   make(map[int]chan []interface{}),
		logger:  logger,
	}
	ec.Handle(client)
	return client
}

func (c *Client) NextReqID() int {
	c.reqId++
	reqID := c.reqId
	return reqID
}

func (c *Client) Handle(h Handler) {
	c.handler = h
}

func (c *Client) HandleFunc(fn func(p *Packet)) {
	c.handler = HandlerFunc(fn)
}

// HandleMessage handles Event.IO Message.
func (c *Client) HandleMessage(msg string) error {
	p, err := DecodePacket(msg)
	if err != nil {
		c.logger.Error("failed to DecodePacket", err)
	}
	c.logger.Info("recv", p.Type)
	switch p.Type {
	case Event:
		c.handler.HandleMessage(&p)
	case Ack:
		c.logger.Info("recv ack id", *p.ID)
		c.ackChMu.Lock()
		if ackCh, ok := c.ackCh[*p.ID]; ok {
			ackCh <- p.Data
			delete(c.ackCh, *p.ID)
		}
		c.ackChMu.Unlock()
	default:
		c.logger.Info("received ignoring type", p.Type)
	}
	return err
}

func (c *Client) Open() error {
	return c.client.Open()
}

func (c *Client) SendPacket(p Packet) error {
	s, err := EncodePacket(p)
	if err != nil {
		return err
	}
	c.logger.Info("sending", s)
	c.client.Send(s)
	return nil
}

func (c *Client) SendPacketAck(p Packet) ([]interface{}, error) {
	reqID := c.NextReqID()
	p.ID = &reqID
	s, err := EncodePacket(p)
	if err != nil {
		return nil, err
	}
	c.ackChMu.Lock()
	c.ackCh[reqID] = make(chan []interface{})
	c.ackChMu.Unlock()
	c.logger.Info("sending need ack", reqID, s)
	c.client.Send(s)

	ack := <-c.ackCh[reqID]
	c.logger.Info("received ack", ack)
	return ack, nil
}

func (c *Client) Send(data []interface{}) error {
	p := Packet{
		Data: data,
		ID:   nil,
		Type: Event,
	}
	return c.SendPacket(p)
}

func (c *Client) SendAck(data []interface{}) ([]interface{}, error) {
	p := Packet{
		Data: data,
		ID:   nil,
		Type: Event,
	}
	return c.SendPacketAck(p)
}

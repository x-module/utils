package server

import (
	"github.com/go-xmodule/utils/gamelift/pkg/gamelift"
	"github.com/go-xmodule/utils/gamelift/pkg/proto/pbuffer"
	"github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/utils"
	"github.com/go-xmodule/utils/utils/xlog"
	"net"
)

type Handler struct {
	client gamelift.Client
	port   int
}

func (h *Handler) StartGameSession(event *pbuffer.ActivateGameSession) {
	xlog.Logger.Println("StartGameSession", event)
	if err := h.client.ActivateGameSession(&pbuffer.GameSessionActivate{
		GameSessionId: event.GetGameSession().GetGameSessionId(),
		MaxPlayers:    event.GetGameSession().GetMaxPlayers(),
		Port:          event.GetGameSession().GetPort(),
	}); err != nil {
		xlog.Logger.WithField(global.ErrField, err).Error("activate game session error")
		return
	}
	xlog.Logger.Println("ActivateGameSession complete. sessionID:", *h.client.GetGameSessionId())
}

func (h *Handler) UpdateGameSession(event *pbuffer.UpdateGameSession) {
	xlog.Logger.Println("UpdateGameSession", event)
}

func (h *Handler) ProcessTerminate(event *pbuffer.TerminateProcess) {
	xlog.Logger.Println("ProcessTerminate", event)
}

func (h *Handler) HealthCheck() bool {
	xlog.Logger.Println("HealthCheck")
	return true
}

type GameLift struct {
	config gamelift.Config
}

func NewGameLift() *GameLift {
	return new(GameLift)
}
func (g *GameLift) Init(config gamelift.Config) *GameLift {
	g.config = config
	return g
}
func (g *GameLift) Server() error {
	conn, port, err := utils.OpenFreeUDPPort(9000, 100)
	if err != nil {
		xlog.Logger.WithField(global.ErrField, err).Error("opens free UDP port error")
		return err
	}
	defer func(conn net.PacketConn) {
		err = conn.Close()
		if err != nil {
			xlog.Logger.WithField(global.ErrField, err).Error("close udp collect error")
		}
	}(conn)
	client := gamelift.NewClient(xlog.Logger, g.config)
	h := &Handler{client: client}
	client.Handle(h)
	if err = client.Open(); err != nil {
		xlog.Logger.WithField(global.ErrField, err).Error("open socket error!")
		return err
	}
	if err = client.ProcessReady(&pbuffer.ProcessReady{
		LogPathsToUpload: []string{},
		Port:             int32(port),
		// MaxConcurrentGameSessions: 0, // not set in original ServerSDK
	}); err != nil {
		xlog.Logger.WithField(global.ErrField, err).Error("gamelift process ready error!")
		return err
	}
	return nil
}

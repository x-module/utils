/**
 * Created by goland.
 * @file   websocket.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/22 15:42
 * @desc   websocket.go
 */

package socket

import (
	"fmt"
	"github.com/go-xmodule/utils/utils/xlog"
	"golang.org/x/net/websocket"
	"net/http"
)

type WebSocket struct {
	Pattern   string
	Address   string
	wsHandler *websocket.Conn
}

func NewWebSocket() *WebSocket {
	return new(WebSocket)
}

func (w *WebSocket) Init(pattern string, address string) *WebSocket {
	w.Pattern = pattern
	w.Address = address
	xlog.Logger.Debug("Received back from client: ")

	return w
}
func (w *WebSocket) SendMessage(message any) error {
	if err := websocket.Message.Send(w.wsHandler, message); err != nil {
		xlog.Logger.WithField("err", err).Error("Can't send message")
		return err
	}
	return nil
}

func (w *WebSocket) Handler(ws *websocket.Conn) {
	w.wsHandler = ws
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			xlog.Logger.WithField("err", err).Error("Can't receive")
			break
		}
		xlog.Logger.Debugf("Received back from client: " + reply)
		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

	}
}

func (w *WebSocket) Server() error {
	http.Handle(w.Pattern, websocket.Handler(w.Handler))
	fmt.Println("-----------", w.Address)
	if err := http.ListenAndServe(w.Address, nil); err != nil {
		xlog.Logger.WithField("err", err).Error("websocket listenAndServe err")
		return err
	} else {

		xlog.Logger.Debug("start websocket server  success")
	}
	return nil
}

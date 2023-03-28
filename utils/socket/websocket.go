/**
 * Created by goland.
 * @file   websocket.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/22 15:42
 * @desc   websocket.go
 */

package socket

import (
	"encoding/json"
	"fmt"
	"github.com/go-xmodule/utils/utils"
	"github.com/go-xmodule/utils/utils/xlog"
	"golang.org/x/net/websocket"
	"net/http"
)

const (
	MessageTypeString = "message"
	MessageTypeEvent  = "event"
)

type MessageCallback func(message Message, socket *WebSocket) error
type EventCallback func(message Message, socket *WebSocket) error

type Message struct {
	Type    string `json:"Type"`
	Event   string `json:"Event"`
	Message string `json:"Message"`
}

type WebSocket struct {
	Pattern         string
	Address         string
	wsHandler       *websocket.Conn
	messageCallback MessageCallback
	eventFunMap     map[string]EventCallback
}

func NewWebSocket() *WebSocket {
	return &WebSocket{
		eventFunMap: map[string]EventCallback{},
	}
}

func (w *WebSocket) OnMessage(action MessageCallback) *WebSocket {
	w.messageCallback = action
	return w
}
func (w *WebSocket) _message(message Message) {
	if err := w.messageCallback(message, w); err != nil {
		xlog.Logger.WithField("err", err).Error("execute message function error!")
	}
}

func (w *WebSocket) Event(event string, message string) error {
	msg := Message{
		Type:    MessageTypeEvent,
		Event:   event,
		Message: message,
	}
	return w.Send(msg)
}

func (w *WebSocket) OnEvent(event string, action EventCallback) *WebSocket {
	w.eventFunMap[event] = action
	return w
}

func (w *WebSocket) _event(message Message) bool {
	if action, exist := w.eventFunMap[message.Event]; exist {
		err := action(message, w)
		if err != nil {
			xlog.Logger.WithField("err", err).Error("execute event action error! event:" + utils.JsonString(message))
			return false
		}
	}
	return true
}

func (w *WebSocket) Init(pattern string, address string) *WebSocket {
	w.Pattern = pattern
	w.Address = address
	return w
}
func (w *WebSocket) SendMessage(message string) error {
	msg := Message{
		Message: message,
		Type:    MessageTypeString,
	}
	return w.Send(msg)
}

func (w *WebSocket) Send(message Message) error {
	if err := websocket.Message.Send(w.wsHandler, utils.JsonString(message)); err != nil {
		xlog.Logger.WithField("err", err).Error("Can't send message")
		return err
	}
	fmt.Println("send back")
	return nil
}

func (w *WebSocket) Handler(ws *websocket.Conn) {
	w.wsHandler = ws
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			xlog.Logger.Info("Can't receive")
			break
		}
		xlog.Logger.Debugf("Received back from client: " + reply)
		var message Message
		err := json.Unmarshal([]byte(reply), &message)
		if err != nil {
			xlog.Logger.WithField("err", err).Error("parse json error")
			return
		}
		if message.Type == MessageTypeString {
			if message.Message == "ping" {
				_ = w.SendMessage("pong")
			} else {
				w._message(message)
			}
		} else if message.Type == MessageTypeEvent {
			w._event(message)
		}
	}
}

func (w *WebSocket) Server() error {
	http.Handle(w.Pattern, websocket.Handler(w.Handler))
	if err := http.ListenAndServe(w.Address, nil); err != nil {
		xlog.Logger.WithField("err", err).Error("websocket listenAndServe err")
		return err
	} else {
		xlog.Logger.Debug("start websocket server  success")
	}
	return nil
}

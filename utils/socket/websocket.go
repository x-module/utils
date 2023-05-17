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
	"github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/utils"
	"github.com/go-xmodule/utils/utils/xlog"
	"github.com/gorilla/websocket"
	"log"
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

var messageChan = make(chan string, 1000)

func (w *WebSocket) Send(message Message) error {
	msg, _ := json.Marshal(message)
	messageChan <- string(msg)
	return nil
}

func (w *WebSocket) _send() {
	defer func() {
		if err := recover(); err != nil {
			xlog.Logger.WithField(global.ErrField, err).Error("write connect message error")
			w._send()
		}
	}()
	for msg := range messageChan {
		err := w.wsHandler.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			xlog.Logger.WithField(global.ErrField, err).Error("send message error")
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (w *WebSocket) handler(ws http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(ws, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	w.wsHandler = c
	go w._send()
	defer c.Close()
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			xlog.Logger.WithField(global.ErrField, err).Error("Can't receive")
			break
		}
		xlog.Logger.Debugf("Received message from client, msgType: %d  msg:%s", mt, string(msg))
		if string(msg) == "p" {
			_ = w.SendMessage("h")
		} else {
			var message Message
			err := json.Unmarshal(msg, &message)
			if err != nil {
				xlog.Logger.WithField("err", err).Error("parse json error")
				return
			}
			if message.Type == MessageTypeString {
				w._message(message)
			} else if message.Type == MessageTypeEvent {
				w._event(message)
			}
		}

	}
}

func (w *WebSocket) Server() error {
	http.HandleFunc(w.Pattern, w.handler)
	if err := http.ListenAndServe(w.Address, nil); err != nil {
		xlog.Logger.WithField("err", err).Error("websocket listenAndServe err")
		return err
	} else {
		xlog.Logger.Debug("start websocket server  success")
	}
	return nil
}

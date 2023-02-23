/**
 * Created by goland.
 * @file   websocket.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/22 15:53
 * @desc   websocket.go
 */

package handler

import "github.com/go-xmodule/utils/utils/socket"

var WebsocketHandler *socket.WebSocket

func InitializeWebsocket(pattern string, address string) *socket.WebSocket {
	WebsocketHandler = socket.NewWebSocket()
	go WebsocketHandler.Init(pattern, address).Server()
	return WebsocketHandler
}

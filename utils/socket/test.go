/**
 * Created by PhpStorm.
 * @file   test.go
 * @author 李锦 <Ljin@cavemanstudio.net>
 * @date   2023/3/21 10:10
 * @desc   test.go
 */

package socket

//
// import (
// "time"
// "vue/server/socket"
// )
//
// var socketHandler *socket.WebSocket
//
// func main() {

// system.InitializeLogger(config.Log{
//  Path: "./data/log",
//  File: "run.log",
//  Mode: "debug",
// })
//
// socketHandler = socket.NewWebSocket()
// socketHandler.Init("/ws", ":2000")
// socketHandler.OnEvent("sync", func(message socket.Message, socket *socket.WebSocket) error {
//  xlog.Logger.Debug("event call", message.Event)
//  socket.Event("sync", "server-sync-message")
//  return nil
// })
// socketHandler.OnMessage(func(message socket.Message, socket *socket.WebSocket) error {
//  xlog.Logger.Debug("receive message", message.Message)
//  socket.SendMessage("我收到了啊")
//  socket.Event("sync", "server-sync")
//  return nil
// })
// go socketHandler.Server()
//
// // 开启websocket
// //handler.InitializeWebsocket("/ws", ":2000")
// //time.Sleep(5 * time.Second)
// //fmt.Println("============================================")
// //_ = handler.WebsocketHandler.SendMessage(" start create party matchmaker")
//
// for {
//  time.Sleep(time.Second)
// }
// }

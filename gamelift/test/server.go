/**
 * Created by goland.
 * @file   server.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/28 10:28
 * @desc   server.go
 */

package main

//
// // InitializeLogger 初始化日志配置
// func InitializeLogger(config config.Log) {
// 	if !gotool.FileUtils.Exists(config.Path) {
// 		err := os.MkdirAll(config.Path, os.ModePerm)
// 		if err != nil {
// 			panic("init system error. make log data err.path:" + config.Path)
// 		}
// 	}
// 	// 日志文件
// 	fileName := path.Join(config.Path, config.File)
// 	if !gotool.FileUtils.Exists(fileName) {
// 		openfile.Create(fileName)
// 		if !gotool.FileUtils.Exists(fileName) {
// 			panic("init system error. create log file err. log file:" + fileName)
// 		}
// 	}
// 	xlog.InitLogger(config.Path, config.File, config.Mode)
// }
//
// func main() {
// 	InitializeLogger(config.Log{
// 		Path: "gamelift/log",
// 		File: "log.log",
// 	})
// 	server.NewGameLift().Init(gamelift.Config{
// 		Domain: "127.0.0.1",
// 		Port:   5757,
// 	}).Server()
// 	for {
//
// 	}
// }

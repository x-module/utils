/**
 * Created by goland.
 * @file   connect.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/7 15:34
 * @desc   connect.go
 */

package server

import (
	"github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"log"
)

// GetNacosClient 获取系统配置
func GetNacosClient(nacosConfig nacos.ConnectConfig) naming_client.INamingClient {
	client, err := nacos.GetNamingClient(nacosConfig)
	if err != nil {
		log.Fatal(global.SystemInitFail.String())
	}
	return client
}

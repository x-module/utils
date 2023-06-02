/**
 * Created by goland.
 * @file   server.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/7 14:48
 * @desc   server.go
 */

package server

import (
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/x-module/utils/global"
	"github.com/x-module/utils/nacos"
	"github.com/x-module/utils/utils/convertor"
	"github.com/x-module/utils/utils/xlog"
	"log"
	"strings"
)

type Server struct {
	Name  string `yaml:"name"`
	Group string `yaml:"group"`
	Desc  string `yaml:"desc"`
}

type ServerInfo struct {
	Describe string `json:"describe,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
}

var ServerList = map[string]ServerInfo{}

type ServerGroup struct {
	ServiceName string   `param:"serviceName"` // required
	Clusters    []string `param:"clusters"`    // optional,default:DEFAULT
	GroupName   string   `param:"groupName"`   // optional,default:DEFAULT_GROUP
}

// SubscribeServer 订阅服务
func SubscribeServer(client naming_client.INamingClient, serverList []ServerGroup) {
	for _, server := range serverList {
		err := nacos.SubscribeServer(nacos.SubscribeServerParams{
			Client:      client,
			ServiceName: server.ServiceName,
			GroupName:   server.GroupName,
			SubscribeCallback: func(services []model.Instance, err error) {
				xlog.Logger.Debugf("server config change:%+v", services)
				serviceName := strings.Split(services[0].ServiceName, "@@")
				var serverInfo ServerInfo
				_ = convertor.TransInterfaceToStruct(services[0].Metadata, &serverInfo)
				ServerList[serviceName[0]] = serverInfo
			},
		})
		if err != nil {
			log.Fatal(global.SystemInitFail.String())
		}
	}
}

// InitServer 当前服务配置初始化
func InitServer(client naming_client.INamingClient, serverList []ServerGroup) {
	for _, server := range serverList {
		service, err := nacos.GetService(nacos.GetServerParams{
			Client:      client,
			ServiceName: server.ServiceName,
			GroupName:   server.GroupName,
		})
		if err != nil {
			log.Fatal(global.SystemInitFail.String())
		}
		if len(service.Hosts) > 0 {
			serviceName := strings.Split(service.Hosts[0].ServiceName, "@@")
			var serverInfo ServerInfo
			_ = convertor.TransInterfaceToStruct(service.Hosts[0].Metadata, &serverInfo)
			ServerList[serviceName[0]] = serverInfo
		}
	}
}

// GetServerAddress 获取服务信息
func GetServerAddress(serverName string) (string, error) {
	if server, ok := ServerList[serverName]; !ok {
		xlog.Logger.Errorf("%s,server:%s", global.UnknownServerErr.String(), serverName)
		return "", errors.New(global.UnknownServerErr.String())
	} else {
		return fmt.Sprintf("%s:%s", server.Host, server.Port), nil
	}
}

// InitServerConfig 初始化个服务配置信息
func InitServerConfig(config nacos.ConnectConfig, serverName string, ServerList []Server) {
	client := GetNacosClient(config)
	var serverList []ServerGroup
	for _, server := range ServerList {
		if server.Name != serverName {
			serverList = append(serverList, ServerGroup{
				ServiceName: server.Name,
				GroupName:   server.Group,
			})
		}
	}
	InitServer(client, serverList)
	SubscribeServer(client, serverList)
}

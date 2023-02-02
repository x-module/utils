/**
 * Created by goland.
 * @file   server.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/30 10:16
 * @desc   server.go
 */

package nacos

import (
	"github.com/go-utils-module/utils/global"
	"github.com/go-utils-module/utils/utils/netutil"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

// RegisterServerParams 服务注册参数
type RegisterServerParams struct {
	Client      naming_client.INamingClient `json:"client,omitempty"`
	Port        uint64                      `json:"port,omitempty"`         // required
	ServiceName string                      `json:"service_name,omitempty"` // required
	ClusterName string                      `json:"cluster_name,omitempty"` // optional,default:DEFAULT
	GroupName   string                      `json:"group_name,omitempty"`   // optional,default:DEFAULT_GROUP
	Metadata    map[string]string           `json:"metadata,omitempty"`     // optional
}

// RegisterServer 注册服务
func RegisterServer(params RegisterServerParams) (bool, error) {
	ip := netutil.GetInternalIp()
	result, err := params.Client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        params.Port,
		ServiceName: params.ServiceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    params.Metadata,
		ClusterName: params.ClusterName,
		GroupName:   params.GroupName, // 默认值DEFAULT_GROUP
	})
	if err != nil {
		log.Printf("%s,err:%s", global.RegisterServerErr.String(), err.Error())
		return false, err
	}
	return result, err
}

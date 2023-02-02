/**
 * Created by goland.
 * @file   config.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/30 10:16
 * @desc   config.go
 */

package nacos

import (
	"github.com/go-utils-module/utils/global"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v2"
	"log"
)

type OnChange func(namespace, group, dataId, data string)

type ListenConfigParams struct {
	Client config_client.IConfigClient `json:"client,omitempty"`
	DataId string                      `json:"data_id,omitempty"`
	Group  string                      `json:"group,omitempty"`
	// AppName  string                      `json:"app_name,omitempty"`
	OnChange OnChange `json:"on_change,omitempty"`
}
type GetConfigParams struct {
	Client config_client.IConfigClient `json:"client,omitempty"`
	DataId string                      `json:"data_id,omitempty"`
	Group  string                      `json:"group,omitempty"`
}

// GetConfig 获取配置
func GetConfig(params GetConfigParams, config any) error {
	content, err := params.Client.GetConfig(vo.ConfigParam{
		Group:  params.Group,
		DataId: params.DataId,
		// AppName: params.AppName,
	})
	if err != nil {
		log.Printf("%s,err:%s", global.GetConfigErr.String(), err.Error())
		return err
	}
	if content == "" {
		log.Fatalf("%s,config:%s", global.GetConfigErr.String(), params.DataId)
	}
	return yaml.Unmarshal([]byte(content), config)
}

// ListenConfig 监听配置文件变化
func ListenConfig(params ListenConfigParams) error {
	err := params.Client.ListenConfig(vo.ConfigParam{
		DataId: params.DataId,
		Group:  params.Group,
		// AppName:  params.AppName,
		OnChange: params.OnChange,
	})
	if err != nil {
		log.Printf("%s,err:%s", global.ListenConfigErr.String(), err.Error())
	}
	return err
}

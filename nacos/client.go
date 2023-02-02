/**
 * Created by goland.
 * @file   client.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/30 10:16
 * @desc   client.go
 */

package nacos

import (
	"github.com/go-utils-module/utils/global"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

type GetAllServerInfoParams struct {
	Client    naming_client.INamingClient `json:"client,omitempty"`
	NameSpace string                      `json:"name_space,omitempty"` // required
	PageNo    uint32                      `param:"pageNo"`              // optional,default:1
	PageSize  uint32                      `param:"pageSize"`            // optional,default:10
	GroupName string                      `json:"group_name,omitempty"` // optional,default:DEFAULT_GROUP
}
type GetServerParams struct {
	Client      naming_client.INamingClient `json:"client,omitempty"`
	ServiceName string                      `json:"service_name,omitempty"` // required
	Cluster     []string                    `json:"cluster_name,omitempty"` // optional,default:DEFAULT
	GroupName   string                      `json:"group_name,omitempty"`   // optional,default:DEFAULT_GROUP
}
type GetInstanceParams struct {
	Client      naming_client.INamingClient `json:"client,omitempty"`
	ServiceName string                      `json:"service_name,omitempty"` // required
	Cluster     []string                    `json:"cluster_name,omitempty"` // optional,default:DEFAULT
	GroupName   string                      `json:"group_name,omitempty"`   // optional,default:DEFAULT_GROUP
}

type SubscribeServerParams struct {
	Client            naming_client.INamingClient                `json:"client,omitempty"`
	ServiceName       string                                     `param:"serviceName"` // required
	Clusters          []string                                   `param:"clusters"`    // optional,default:DEFAULT
	GroupName         string                                     `param:"groupName"`   // optional,default:DEFAULT_GROUP
	SubscribeCallback func(services []model.Instance, err error) // required
}

// GetAllServicesInfo 获取注册中信的服务
func GetAllServicesInfo(params GetAllServerInfoParams) (model.ServiceList, error) {
	service, err := params.Client.GetAllServicesInfo(
		vo.GetAllServiceInfoParam{
			NameSpace: params.NameSpace,
			GroupName: params.GroupName,
			PageNo:    params.PageNo,
			PageSize:  params.PageSize,
		})
	if err != nil {
		log.Printf("%s,err:%s", global.GetServerErr.String(), err.Error())
		return model.ServiceList{}, err
	}
	return service, nil
}

// GetService 获取注册中信的服务
func GetService(params GetServerParams) (model.Service, error) {
	service, err := params.Client.GetService(
		vo.GetServiceParam{
			ServiceName: params.ServiceName,
			GroupName:   params.GroupName,
			Clusters:    params.Cluster,
		})
	if err != nil {
		log.Printf("%s,err:%s", global.GetServerErr.String(), err.Error())
		return model.Service{}, err
	}
	return service, nil
}

// GetInstance 获取注册中信的服务实例
func GetInstance(params GetInstanceParams) ([]model.Instance, error) {
	instances, err := params.Client.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: params.ServiceName,
		GroupName:   params.GroupName,
		Clusters:    params.Cluster,
	})
	if err != nil {
		log.Printf("%s,err:%s", global.GetInstanceErr.String(), err.Error())
		return nil, err
	}
	return instances, nil
}

// SubscribeServer 服务订阅
func SubscribeServer(params SubscribeServerParams) error {
	param := &vo.SubscribeParam{
		ServiceName:       params.ServiceName,
		GroupName:         params.GroupName,
		SubscribeCallback: params.SubscribeCallback,
	}
	err := params.Client.Subscribe(param)
	if err != nil {
		log.Printf("%s,err:%s", global.SubscribeServerErr.String(), err.Error())
		return err
	}
	return nil
}

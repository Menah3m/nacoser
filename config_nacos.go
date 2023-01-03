package nacoser

/*
   @Auth: menah3m
   @Desc: 创建nacos config client
*/

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

// 用于连接nacos的client 的config信息
type Config struct {
	ClientConfig constant.ClientConfig
	ServerConfig []constant.ServerConfig
}

// 需要获取的nacos上的config
type Params struct {
	DataID  string
	Group   string
	Content string
}

// CreateNewConfigClient 创建用于查询配置的客户端
func (c *Config) CreateNewConfigClient() config_client.IConfigClient {
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &c.ClientConfig,
			ServerConfigs: c.ServerConfig,
		})
	if err != nil {
		log.Println("Config init: create new config client failed,err:", err)
	}

	return configClient
}

// PublishNacosConfig 从Nacos中读取配置
func (p *Params) PublishNacosConfig(configClient config_client.IConfigClient) (bool, error) {
	success, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  p.DataID,
		Group:   p.Group,
		Content: p.Content,
	})
	if err != nil {
		log.Println("GetNacosConfig: get nacos config failed,err:", err)
		return false, err
	}
	return success, nil
}

// GetNacosConfig 从Nacos中读取配置
func (p *Params) GetNacosConfig(configClient config_client.IConfigClient) (string, error) {
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: p.DataID,
		Group:  p.Group,
	})
	if err != nil {

		log.Println("GetNacosConfig: get nacos config failed,err:", err)
		return "", err
	}
	return content, nil
}

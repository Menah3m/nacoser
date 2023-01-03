// 实现了读取nacos配置的封装
package nacoser

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

/*
   @Auth: menah3m
   @Desc:
*/

// type NacosConfig struct {
// 	ServerConfig NacosServerConfig
// 	ClientConfig NacosClientConfig
// }

type NacosServerConfig struct {
	IpAddr      string
	Port        string
	ContextPath string
	Scheme      string
}

type NacosClientConfig struct {
	NamespaceId string
	TimeoutMs   string
	Username    string
	Password    string
}

type FilePath struct {
	Type string
	Name string
	Path string
}

// 包含了 viper.Viper 的结构体
type Configer struct {
	Viper        *viper.Viper
	ServerConfig NacosServerConfig
	ClientConfig NacosClientConfig
	Params       Params
}

// 从配置文件中获取连接nacos的配置信息
func (c *Configer) ReadParamsFromFile(filePath string) error {
	parse := ParseFilePath(filePath)
	c.Viper.SetConfigType(parse.Type)
	c.Viper.SetConfigName(parse.Name)
	c.Viper.AddConfigPath(parse.Path)
	c.Viper.AddConfigPath(".")
	err := c.Viper.ReadInConfig()
	if err != nil {
		return err
	}
	// server config params
	c.ServerConfig.Scheme = c.Viper.GetString("nacosserver.scheme")
	c.ServerConfig.IpAddr = c.Viper.GetString("nacosserver.ipaddr")
	c.ServerConfig.Port = c.Viper.GetString("nacosserver.port")
	c.ServerConfig.ContextPath = c.Viper.GetString("nacosserver.contextpath")
	// client config params
	c.ClientConfig.NamespaceId = c.Viper.GetString("nacosclient.namespaceid")
	c.ClientConfig.TimeoutMs = c.Viper.GetString("nacosclient.timeoutms")
	c.ClientConfig.Password = c.Viper.GetString("nacosclient.password")
	c.ClientConfig.Username = c.Viper.GetString("nacosclient.username")
	// target params
	c.Params.DataID = c.Viper.GetString("target.dataID")
	c.Params.Group = c.Viper.GetString("tartget.group")
	c.Params.Content = c.Viper.GetString("target.content")
	//

	fmt.Println(c)
	return nil
}

func (c *Configer) BindNacosClientParams() Config {
	timeoutms, _ := strconv.Atoi(c.ClientConfig.TimeoutMs)
	clientConfig := constant.ClientConfig{
		TimeoutMs:   uint64(timeoutms),
		NamespaceId: c.ClientConfig.NamespaceId,
		Username:    c.ClientConfig.Username,
		Password:    c.ClientConfig.Password,
		LogDir:      ".",
	}

	port, _ := strconv.Atoi(c.ServerConfig.Port)
	serverConfig := []constant.ServerConfig{
		{
			Scheme:      c.ServerConfig.Scheme,
			ContextPath: c.ServerConfig.ContextPath,
			IpAddr:      c.ServerConfig.IpAddr,
			Port:        uint64(port),
		},
	}
	fmt.Println("BindNacosClientParams")
	return Config{
		ClientConfig: clientConfig,
		ServerConfig: serverConfig,
	}
}

func (c *Configer) BindTargetParams() Params {
	fmt.Println("BindTargetParams")
	return Params{
		DataID:  c.Params.DataID,
		Group:   c.Params.Group,
		Content: c.Params.Content,
	}

}

func ParseFilePath(filepath string) FilePath {
	fields := strings.Split(filepath, "/")
	fileName := fields[len(fields)-1]

	subfields := strings.Split(fileName, ".")
	name := subfields[0]
	typ := subfields[len(subfields)-1]
	path := strings.Join(fields[:len(fields)-1], "/")

	return FilePath{
		Type: typ,
		Name: name,
		Path: path,
	}
}

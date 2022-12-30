// 示例 演示了如何使用
package main

import (
	"fmt"
	nacoser "github.com/menah3m/nacoser"
	"github.com/spf13/viper"
	"log"
)

/*
   @Auth: menah3m
   @Desc:
*/

var config = viper.New()

func main() {
	c := nacoser.Configer{
		Viper: config,
	}
	fp := "./config.yaml"
	err := c.ReadParamsFromFile(fp)
	if err != nil {
		log.Panicln("fatal error config file:", err)
	}
	clientParams := c.BindNacosClientParams()
	targetParams := c.BindTargetParams()
	client := clientParams.CreateNewConfigClient()

	_, err = targetParams.PublishNacosConfig(client)
	if err != nil {
		fmt.Println("publish config err:", err)
	}
	fmt.Println("publish config success!")

	remoteConfigInfo, err := targetParams.GetNacosConfig(client)
	if err != nil {
		log.Println("getNacosConfig err:", err)
	}
	fmt.Println("[config name] ", targetParams.DataID)
	fmt.Println(remoteConfigInfo)
}

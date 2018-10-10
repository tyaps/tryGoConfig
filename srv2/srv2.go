package main

import (

	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"

	"poc_config/consul/srv2/service"

	lib_conf "poc_config/consul/lib/configuration"
	"poc_config/consul/srv2/configuration"
	//"net/http"
	//"os"
	//"poc_config/consul/srv1/pack1"
	//"poc_config/consul/srv1/service"
	//"strings"

	//"time"

	//"github.com/alexdzyoba/webkv/service"

	//To enable remote support in Viper, do a blank import of the viper/remote package:

	_ "github.com/spf13/viper/remote"
)




//var myConfiguration *myConfig
var confManager *lib_conf.Manager


func ParseConfigFromViper() (*configuration.AppConfig, error) {

	var c = new(configuration.AppConfig)
	c.Key1=viper.GetString("Key1")
	c.Title =viper.GetString("Title")
	c.Service1Url = viper.GetString("Service1Url")

	return c, nil
}

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	//прочитать локальный конфиг, чтобы знать, откуда забираю конфиг + список сервисов (пока не дискавери)
	//viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigFile("config.yml")
	viper.ReadInConfig()
	remoteConfigKey:=viper.GetString("remoteConfigKey")


	confManager = new(lib_conf.Manager) //configuration.New() //

	err:= confManager.LoadConfiguration(remoteConfigKey) // getConfig()
	if err!=nil{
		panic(err)
	}

	//ура, работаем, конфиг настроен
	//вывести конфигурацию
	appConfig,_:= ParseConfigFromViper()
	log.Printf("appConfig %v ", appConfig)

	configuration.AppConfiguration = appConfig


	s, err := service.New()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", s)


	l := fmt.Sprintf(":%d", *port)
	log.Print("Listening on ", l)
	log.Fatal(http.ListenAndServe(l, nil))
}

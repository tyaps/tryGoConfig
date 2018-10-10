package main

import (
	"flag"
	"fmt"
	"github.com/armon/consul-api"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"poc_config/consul/srv1/configuration"
	"poc_config/consul/srv1/service"

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
func checkConsul2(){
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}
func checkConsl(){



	config := consulapi.DefaultConfig()
	//config.Address = "127.0.0.1:8500"
	consul, err := consulapi.NewClient(config)
	if err!=nil{
		log.Print("%v", err)
	}

	kv := consul.KV()
	d := &consulapi.KVPair{Key: "sites", Value: []byte("example.com")}
	_, err = kv.Put(d, nil)

	//d := &consulapi.KVPair{Key: "sites", Value: []byte("example.com")}
	//var b bool
	//var mt *consulapi.WriteMeta
	//b, mt, err = kv.Acquire(d, nil)
	//
	//log.Print("%v", b)
	//log.Print("%v", mt)
	//log.Print("%v", err)
	//
	//kvp, qm, error := kv.Get("sites", nil)
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Println(qm)
	//	fmt.Println(error)
	//} else {
	//	fmt.Println(string(kvp.Value))
	//}


}




//var myConfiguration *myConfig
var confManager *configuration.Manager


func ParseConfigFromViper() (*configuration.AppConfig, error) {

	var c = new(configuration.AppConfig)
	c.Key1=viper.GetString("Key1")
	c.Title=viper.GetString("Title")

	return c, nil
}

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	confManager = new(configuration.Manager) //configuration.New() //

	err:= confManager.LoadConfiguration("configs/srv1") // getConfig()
	if err!=nil{
		panic(err)
	}

	//ура, работаем, конфиг настроен
	//вывести конфигурацию
	appConfig,_:= ParseConfigFromViper()
	log.Printf("appConfig %v ", appConfig)

	configuration.AppConfiguration = appConfig

	//потом подписаться на обновления






	s, err := service.New()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", s)


	l := fmt.Sprintf(":%d", *port)
	log.Print("Listening on ", l)
	log.Fatal(http.ListenAndServe(l, nil))
}

func comments(){
//	viper.SetDefault("c1", "content1")
//	log.Print(viper.Get("c1"))
//
//	//viper.SetConfigName("config1")
//	//viper.AddConfigPath(".")               // optionally look for config in the working directory
//	//err := viper.ReadInConfig() // Find and read the config file
//	//if err != nil { // Handle errors reading the config file
//	//	panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	//}
//	//log.Print(viper.Get("c2"))
//
//
//	var yamlExample = []byte(`
//Hacker: true
//#comm
//name: steve #comment
//hobbies:
//- skateboarding
//- snowboarding
//- go
//clothing:
//  jacket: leather
//  trousers: denim
//age: 35 /*comment*/
//eyes : brown --comment
//beard: true dfg
//`)
//
//	viper.SetConfigType("yaml")
//	var err1 = viper.ReadConfig(bytes.NewBuffer(yamlExample))
//	log.Printf("error %v", err1)
//
//	log.Print(viper.Get("name")) // this would be "steve"
//	log.Print(viper.Get("c1")) // this would be "steve"
//
//
//	var jsExample=`
//{"name": "mark"}
//`
//
//
//	////	jsExample=`
//	////{"name":"123",
//	////"b":"243"
//	////}
//	//`
//	log.Print(viper.Get("name"))
//	viper.SetConfigType("json")
//	err1 = viper.ReadConfig(bytes.NewBuffer([]byte(jsExample)))
//	log.Printf("error %v", err1)
//	log.Print(viper.Get("name"))
//	log.Print(viper.Get("a"))
//	log.Print(viper.Get("b"))
//
//	//viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
//	//viper.SetConfigType("json") // Need to explicitly set this to json
//	//err := viper.ReadRemoteConfig()
//	//log.Print(viper.Get("c3"))
//
//	//viper.WatchConfig()
//	//viper.OnConfigChange(func(e fsnotify.Event) {
//	//	fmt.Println("Config file changed:", e.Name)
//	//})
//
//	a:=pack1.Ttt()
//	log.Print(a)
//
//	return
}
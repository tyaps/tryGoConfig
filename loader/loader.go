package main

import (
	"fmt"
	"github.com/armon/consul-api"
	"github.com/hashicorp/consul/api"
	"log"
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

type config struct{
	Path string
	Value string
}

func readConfigs() ([]config, error){

	s := []config {
		{"srv1", `
key1: 1
key2: 2`,
	},
		{"srv2", `
key1: 10
key2: 20`,
		},

	}

	return s, nil
}

func writeConfigsToStore(configs []config) error{


	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err!=nil{
		log.Print("%v", err)
	}

	kv := consul.KV()

	for _, c:=range configs{
		d := &consulapi.KVPair{Key: c.Path, Value: []byte(c.Value)}
		_, err = kv.Put(d, nil)

	}

	return nil //TODO: errors

}

func main() {

	c, err:=readConfigs();

	if  err!=nil{
		panic(err)
	}

	log.Printf("res=%v %v", len(c), c)

	writeConfigsToStore(c)


}

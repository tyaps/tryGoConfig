package configuration

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

type Manager struct{}



func New() (*Manager, error) {
	s := new(Manager)



	return s, nil
}


func (*Manager) LoadConfiguration(key string) (error){

	//загружаем из consul
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	pair, _, err := kv.Get(key, nil) //TODO: serviceName

	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)

	yamlConfig:=string(pair.Value)


	//распарсить конфиг через viper
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer([]byte(yamlConfig)))
	if err != nil {
		panic(err)
	}


	return  nil

}




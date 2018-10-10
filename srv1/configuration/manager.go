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


func (*Manager) LoadConfiguration() (error){

	//загружаем из consul
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	pair, _, err := kv.Get("srv1", nil) //TODO: serviceName

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

	c,_:= ParseConfigFromViper()
	fmt.Printf("KV: %v", c)

	return  nil

}


//not used
func ParseConfigFromViper() (*AppConfig, error) {

	var c = new(AppConfig)
	c.Key1=viper.GetString("Key1")
	c.Key2=viper.GetString("Key2")

	return c, nil
}

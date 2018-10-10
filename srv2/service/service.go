package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"poc_config/consul/srv2/configuration"
	"time"

	"github.com/go-redis/redis"
	consul "github.com/hashicorp/consul/api"
)

type Service struct {
	Name        string
	TTL         time.Duration
	RedisClient redis.UniversalClient
	ConsulAgent *consul.Agent
}

//func New(addrs []string, ttl time.Duration) (*Service, error) {
//	s := new(Service)
//	s.Name = "webkv"
//	s.TTL = ttl
//	s.RedisClient = redis.NewUniversalClient(&redis.UniversalOptions{
//		Addrs: addrs,
//	})
//
//	ok, err := s.Check()
//	if !ok {
//		return nil, err
//	}
//
//	c, err := consul.NewClient(consul.DefaultConfig())
//	if err != nil {
//		return nil, err
//	}
//	s.ConsulAgent = c.Agent()
//
//	serviceDef := &consul.AgentServiceRegistration{
//		Name: s.Name,
//		Check: &consul.AgentServiceCheck{
//			TTL: s.TTL.String(),
//		},
//	}
//
//	if err := s.ConsulAgent.ServiceRegister(serviceDef); err != nil {
//		return nil, err
//	}
//	go s.UpdateTTL(s.Check)
//
//	return s, nil
//}

func New() (*Service, error) {
	s := new(Service)

	s.Name = "webkv"

	return s, nil
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var greeting string

	//ask greeting
	response, err := http.Get(configuration.AppConfiguration.Service1Url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		greeting=string(contents)
		//fmt.Printf("%s\n", string(contents))
	}



	fmt.Fprintf(w, "%v %v", greeting, configuration.AppConfiguration.Title)
	//
	//log.Printf("url=\"%s\" remote=\"%s\" key=\"%s\" status=%d\n",
	//	r.URL, r.RemoteAddr, key, status)
}


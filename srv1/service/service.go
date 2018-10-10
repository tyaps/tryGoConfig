package service

import (
	"fmt"
	"net/http"
	"poc_config/consul/srv1/configuration"
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
	//status := 200
	//
	//key := strings.Trim(r.URL.Path, "/")
	//val, err := s.RedisClient.Get(key).Result()
	//if err != nil {
	//	http.Error(w, "Key not found", http.StatusNotFound)
	//	status = 404
	//}

	fmt.Fprint(w,  configuration.AppConfiguration.Greeting)
	//
	//log.Printf("url=\"%s\" remote=\"%s\" key=\"%s\" status=%d\n",
	//	r.URL, r.RemoteAddr, key, status)
}


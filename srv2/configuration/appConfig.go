package configuration

//конфиг самого приложения
type AppConfig struct{
	Key1     	string
	Title 		string
	Service1Url string
}

var AppConfiguration *AppConfig

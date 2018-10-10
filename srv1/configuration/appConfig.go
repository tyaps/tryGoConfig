package configuration

//конфиг самого приложения
type AppConfig struct{
	Key1     string
	Greeting string
}

var AppConfiguration *AppConfig

package common

import "github.com/spf13/viper"

type IConfig struct {
	DNSPort int
	TTL     int
	Suffix  string
}

var Config IConfig

func InitConfig() {
	// read config from config.yml
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Config = IConfig{
		DNSPort: viper.GetInt("dns.port"),
		TTL:     viper.GetInt("dns.ttl"),
		Suffix:  viper.GetString("dns.suffix"),
	}
}

package infra

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	cfg *viper.Viper
)

type Infra interface {
	ApiServer() string
	Config() *viper.Viper
}

type infra struct{}

func NewInfra() Infra {
	return &infra{}
}

func (i *infra) ApiServer() string {
	host := i.Config().GetString("HTTPHOST")
	port := i.Config().GetString("HTTPPORT")
	return fmt.Sprintf("%s:%s", host, port)
}
func (i *infra) Config() *viper.Viper {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	cfg = viper.GetViper()
	return cfg
}

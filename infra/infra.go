package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	cfg *viper.Viper
)

type Infra interface {
	ApiServer() string
	Config() *viper.Viper
	SqlDb() *sql.DB
}

type infra struct{}

func NewInfra() Infra {
	return &infra{}
}

func (i *infra) SqlDb() *sql.DB {
	dbUser := i.Config().GetString("DBUSER")
	dbPassword := i.Config().GetString("DBPASSWORD")
	dbHost := i.Config().GetString("DBHOST")
	dbPort := i.Config().GetString("DBPORT")
	schema := i.Config().GetString("DBSCHEMA")
	dbEngine := i.Config().GetString("DBENGINE")

	db, err := sql.Open(dbEngine, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, schema))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
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

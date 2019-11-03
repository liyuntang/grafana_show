package database

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_"github.com/Go-SQL-Driver/MySQL"
	"log"
	"os"
)

type grafanadb struct {
	User string	`toml:"user"`
	Passwd string	`toml:"passwd"`
	Address string	`toml:"address"`
	Port int	`toml:"port"`
	Charset string	`toml:"charset"`
	Database string	`toml:"database"`
	Table string	`toml:"table"`
	Logger *log.Logger
}

type Grafana_show struct {
	Endpoint string	`xorm:"endpoint"`
	Metric string	`xorm:"metric"`
	Value int		`xorm:"value"`
	Monit_time int64	`xorm:"monit_time"`
}




func (g *grafanadb)GetEngine() *xorm.Engine {
	endPoint := fmt.Sprintf("%s:%d", g.Address, g.Port)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true", g.User, g.Passwd, endPoint, g.Database, g.Charset)
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		g.Logger.Println("init db connection", endPoint, "is bad")
		os.Exit(1)
	}
	return engine
}

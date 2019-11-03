package monit

import (
	"grafana_show/tomlConfig"
	"log"
)

func GetMonitHost(configration *tomlConfig.Config, logger *log.Logger) []tomlConfig.MonitHost {
	db := tomlConfig.Grafanadb{}
	db.User = configration.Grafanadb.User
	db.Passwd = configration.Grafanadb.Passwd
	db.Address = configration.Grafanadb.Address
	db.Port = configration.Grafanadb.Port
	db.Database = configration.Grafanadb.Database
	db.Charset = configration.Grafanadb.Charset
	db.Logger = logger
	return db.GetMonitHost("host_info")
	//return db.GetMonitHost("test")
}

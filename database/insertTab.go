package database

import (
	"grafana_show/tomlConfig"
	"log"
)

func InsertTab(configration *tomlConfig.Config, logger *log.Logger, monitTime int64, ch chan map[string][]map[string]int) {
	db := grafanadb{}
	db.User = configration.Grafanadb.User
	db.Passwd = configration.Grafanadb.Passwd
	db.Address = configration.Grafanadb.Address
	db.Port = configration.Grafanadb.Port
	db.Database = configration.Grafanadb.Database
	db.Charset = configration.Grafanadb.Charset
	db.Logger = logger
	engine := db.GetEngine()
	rows := []Grafana_show{}
	for dict := range ch {
		for endpoint, monitElementMap := range dict {
			for _, elementDict := range monitElementMap {
				for element, value := range elementDict {
					//fmt.Println("enpoint is", endpoint, "element is", element, "value is", value)
					row := Grafana_show{}
					row.Endpoint = endpoint
					row.Metric = element
					row.Value = value
					row.Monit_time = monitTime
					rows = append(rows, row)
					//fmt.Println("endpoint", reflect.TypeOf(endpoint), "metric", reflect.TypeOf(element), "value", reflect.TypeOf(value))
				}
			}
		}
	}
	//fmt.Println(rows)
	num, err := engine.Insert(&rows)
	if err != nil {
		logger.Println("insert to table is bad, err is", err)
	}
	logger.Println("写入了", num, "行数据")
}

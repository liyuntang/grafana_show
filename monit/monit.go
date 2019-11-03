package monit

import (
	"grafana_show/tomlConfig"
	"strings"
	"sync"
)

func Monit(host tomlConfig.MonitHost, monitElement []string, ch chan map[string][]map[string]int, wait *sync.WaitGroup)  {
	// 定义一个map[string]int类型的slice用于存放分析后的结果
	resultMap := map[string][]map[string]int{}
	firstMap, secondMap := host.GetStatus()
	replicationMap := host.GetReplication()
	//遍历monitElement
	for _, element := range monitElement {
		tmpMap := map[string]int{}
		switch strings.ToLower(element) {
		case "threads_connected":
			tmpMap[element] = Threads_connected(firstMap, host.HostStatus)
		case "threads_running":
			tmpMap[element] = Threads_running(firstMap, host.HostStatus)
		case "max_used_connections":
			tmpMap[element] = Max_used_connections(firstMap, host.HostStatus)
		case "aborted_connects":
			tmpMap[element] = Aborted_connects(firstMap, host.HostStatus)
		case "com_select":
			tmpMap[element] = Com_select(firstMap, secondMap, host.HostStatus)
		case "com_insert":
			tmpMap[element] = Com_insert(firstMap, secondMap, host.HostStatus)
		case "com_update":
			tmpMap[element] = Com_update(firstMap, secondMap, host.HostStatus)
		case "com_delete":
			tmpMap[element] = Com_delete(firstMap, secondMap, host.HostStatus)
		case "com_rollback":
			tmpMap[element] = Com_rollback(firstMap, secondMap, host.HostStatus)
		case "com_commit":
			tmpMap[element] = Com_commit(firstMap, secondMap, host.HostStatus)
		case "queries":
			tmpMap[element] = Queries(firstMap, secondMap, host.HostStatus)
		case "tps":
			tmpMap[element] = tps(firstMap, secondMap, host.HostStatus)
		case "qps":
			tmpMap[element] = qps(firstMap, secondMap, host.HostStatus)
		case "uptime":
			tmpMap[element] = Uptime(firstMap, host.HostStatus)
		case "slave_io_running":
			if replicationMap != nil {
				tmpMap[element] = Slave_IO_Running(replicationMap, host.HostStatus)
			}
		case "slave_sql_running":
			if replicationMap != nil {
				tmpMap[element] = Slave_SQL_Running(replicationMap, host.HostStatus)
			}
		case "seconds_behind_master":
			if replicationMap != nil {
				tmpMap[element] = Seconds_Behind_Master(replicationMap, host.HostStatus)
			}
		case "binlogdelay":
			if replicationMap != nil {
				tmpMap[element] = binlogDelay(replicationMap, host.HostStatus)
			}
		case "postiondelay":
			if replicationMap != nil {
				tmpMap[element] = postionDelay(replicationMap, host.HostStatus)
			}
		case "slowlog":{
			tmpMap[element] = slowLog(firstMap, secondMap, host.HostStatus)
		}
		default:
			//host.Logger.Println(host.Endpoint, "monit element is error, elemet is", element)
			// 存在两种情况：1、mysql宕机或者不可达导致返回的空map，2、监控项没匹配上，增加监控项即可
		}
		resultMap[host.Endpoint] = append(resultMap[host.Endpoint], tmpMap)

	}
	//fmt.Println(resultMap)
	wait.Done()
	ch <- resultMap
}

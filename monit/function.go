package monit

import (
	"strconv"
	"strings"
)

func Slave_IO_Running(replicationMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	// 还有一种情况就是，从库由于下线维护可能会吧同步进程给停止或者reset，这个地方需要标示出来或者报警
	if len(replicationMap) == 0 {
		return -1
	}
	if strings.ToLower(replicationMap["Slave_IO_Running"]) == "yes" {
		return 1
	} else if strings.ToLower(replicationMap["Slave_IO_Running"]) == "no" {
		return 0
	} else {
		return -1
	}
}

func Slave_SQL_Running(replicationMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	// 还有一种情况就是，从库由于下线维护可能会吧同步进程给停止或者reset，这个地方需要标示出来或者报警
	if len(replicationMap) == 0 {
		return -1
	}
	switch strings.ToLower(replicationMap["Slave_SQL_Running"]) {
	case "yes":
		return 1
	case "no":
		return 0
	default:
		return -1
	}
}

func Seconds_Behind_Master(replicationMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	// 还有一种情况就是，从库由于下线维护可能会吧同步进程给停止或者reset，这个地方需要标示出来或者报警
	if len(replicationMap) == 0 {
		return -1
	}
	if strings.ToLower(replicationMap["Seconds_Behind_Master"]) == "null" {
		return -1
	} else {
		value, err := strconv.Atoi(replicationMap["Seconds_Behind_Master"])
		if err != nil {
			return -1
		}
		return value
	}
}

func binlogDelay(replicationMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	// 还有一种情况就是，从库由于下线维护可能会吧同步进程给停止或者reset，这个地方需要标示出来或者报警
	if len(replicationMap) == 0 {
		return -1
	}
	Master_Log_File, err3 := strconv.Atoi(strings.Trim(strings.Split(replicationMap["Master_Log_File"],".")[1], "0"))
	if err3 != nil {
		return -1
	}
	Relay_Master_Log_File, err4 := strconv.Atoi(strings.Trim(strings.Split(replicationMap["Relay_Master_Log_File"],".")[1], "0"))
	if err4 != nil {
		return -1
	}
	return Master_Log_File - Relay_Master_Log_File
}

func postionDelay(replicationMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	// 还有一种情况就是，从库由于下线维护可能会吧同步进程给停止或者reset，这个地方需要标示出来或者报警
	if len(replicationMap) == 0 {
		return -1
	}
	Read_Master_Log_Pos, err1 := strconv.Atoi(replicationMap["Read_Master_Log_Pos"])
	if err1 != nil {
		return -1
	}
	Exec_Master_Log_Pos, err2 := strconv.Atoi(replicationMap["Exec_Master_Log_Pos"])
	if err2 != nil {
		return -1
	}
	return Read_Master_Log_Pos - Exec_Master_Log_Pos
}

func Uptime(firstMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value, err := strconv.Atoi(firstMap["Uptime"])
	if err != nil {
		return -1
	}
	return value
}

func Threads_connected(firstMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value, err := strconv.Atoi(firstMap["Threads_connected"])
	if err != nil {
		return -1
	}
	return value
}

func Threads_running(firstMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value, err := strconv.Atoi(firstMap["Threads_running"])
	if err != nil {
		return -1
	}
	return value
}

func Max_used_connections(firstMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value, err := strconv.Atoi(firstMap["Max_used_connections"])
	if err != nil {
		return -1
	}
	return value
}

func Aborted_connects(firstMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value, err := strconv.Atoi(firstMap["Aborted_connects"])
	if err != nil {
		return -1
	}
	return value
}

func Com_select(firstMap, secondMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value1, err := strconv.Atoi(firstMap["Com_select"])
	if err != nil {
		return -1
	}
	value2, err := strconv.Atoi(secondMap["Com_select"])
	if err != nil {
		return -1
	}
	return value2 - value1
}

func Com_insert(firstMap, secondMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value1, err := strconv.Atoi(firstMap["Com_insert"])
	if err != nil {
		return -1
	}
	value2, err := strconv.Atoi(secondMap["Com_insert"])
	if err != nil {
		return -1
	}
	return value2 - value1
}

func Com_update(firstMap, secondMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value1, err := strconv.Atoi(firstMap["Com_update"])
	if err != nil {
		return -1
	}
	value2, err := strconv.Atoi(secondMap["Com_update"])
	if err != nil {
		return -1
	}
	return value2 - value1
}

func Com_delete(firstMap, secondMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value1, err := strconv.Atoi(firstMap["Com_delete"])
	if err != nil {
		return -1
	}
	value2, err := strconv.Atoi(secondMap["Com_delete"])
	if err != nil {
		return -1
	}
	return value2 - value1
}

func Com_rollback(firstMap, secondMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value1, err := strconv.Atoi(firstMap["Com_rollback"])
	if err != nil {
		return -1
	}
	value2, err := strconv.Atoi(secondMap["Com_rollback"])
	if err != nil {
		return -1
	}
	return value2 - value1
}


func Com_commit(firstMap, secondMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	insert := Com_insert(firstMap, secondMap, hostStatus)
	update := Com_update(firstMap, secondMap, hostStatus)
	delete := Com_delete(firstMap,secondMap, hostStatus)
	return insert+update+delete
}

func Queries(firstMap, secondMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value1, err := strconv.Atoi(firstMap["Queries"])
	if err != nil {
		return -1
	}
	value2, err := strconv.Atoi(secondMap["Queries"])
	if err != nil {
		return -1
	}
	return value2 - value1
}

func tps(firstMap, secondMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	commit := Com_commit(firstMap, secondMap, hostStatus)
	rollbak := Com_rollback(firstMap, secondMap, hostStatus)
	return commit+rollbak
}

func qps(firstMap, secondMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	return Queries(firstMap, secondMap, hostStatus)
}

func slowLog(firstMap, secondMap map[string]string, hostStatus bool) int {
	if !hostStatus {
		return -1
	}
	value1, err := strconv.Atoi(firstMap["Slow_queries"])
	if err != nil {
		return -1
	}
	value2, err := strconv.Atoi(secondMap["Slow_queries"])
	if err != nil {
		return -1
	}
	return value2 - value1
}

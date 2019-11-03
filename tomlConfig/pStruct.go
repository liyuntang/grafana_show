package tomlConfig

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/go-xorm/xorm"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	System system
	Grafanadb Grafanadb
}

type system struct {
	LogFile string	`toml:"logFile"`
	MonitIntervalTime string	`toml:"monitIntervalTime"`
	MonitElement []string	`toml:"monitElement"`
}

type Grafanadb struct {
	User string	`toml:"user"`
	Passwd string	`toml:"passwd"`
	Address string	`toml:"address"`
	Port int	`toml:"port"`
	Charset string	`toml:"charset"`
	Database string	`toml:"database"`
	Table string	`toml:"table"`
	Logger *log.Logger
}

type MonitHost struct {
	HostStatus bool		// 纪录mysql的状态，true：正常  false：宕机或者不可达
	Endpoint string
	User string
	Passwd string
	Address string
	Port int
	Charset string
	Database string
	Logger *log.Logger

}

func (mh *MonitHost) GetStatus() (FirstMap, SecondMap map[string]string) {
	engine := mh.GetEngine()
	if engine == nil {
		// 说明mysql宕机或者不可达,将mysql的状态设置为false
		mh.HostStatus = false
		return nil, nil
	} else {
		// 说明mysql正常,将mysql的状态设置为true
		mh.HostStatus = true
		// 第一次获取status
		firstMap := getStatus(engine, mh.Logger)
		// 间隔1s
		time.Sleep(1*time.Second)
		// 第二次获取status
		secondMap := getStatus(engine, mh.Logger)
		// 返回两个status
		return firstMap, secondMap
	}

}

// 获取数据同步相关的监控值
func (mh *MonitHost) GetReplication() (replicationMap map[string]string) {
	if strings.Contains(mh.Endpoint, "master") && !strings.Contains(mh.Endpoint, "backup") {
		// 说明该host为master主库，如果是主库的话则没有replication状态
		//fmt.Println("sorry master is not have replication")
		return nil
	}
	if !mh.HostStatus {
		return nil
	}
	tmpMap := map[string]string{}
	engine := mh.GetEngine()
	resultSlice, err6 := engine.QueryString("show slave status;")
	if err6 != nil {
		mh.Logger.Println("get replication status of", mh.Endpoint, "is bad")
		return nil
	}
	//fmt.Println(resultSlice)
	for _, dict := range resultSlice {
		for key, value := range dict {
			tmpMap[key] = value
		}
	}
	return tmpMap
}

func getStatus(engine *xorm.Engine, logger *log.Logger) map[string]string {
	statMap := map[string]string{}
	sql := "show global status;"
	resultSlice, err5 := engine.QueryString(sql)
	if err5 != nil {
		fmt.Println("get status is bad, err5 is", err5)
	}
	for _, dict := range resultSlice {
		statMap[dict["Variable_name"]] = dict["Value"]
	}
	return statMap
}

func (mh *MonitHost)GetEngine() *xorm.Engine {
	endPoint := fmt.Sprintf("%s:%d", mh.Address, mh.Port)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true", mh.User, mh.Passwd, endPoint, mh.Database, mh.Charset)
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		mh.Logger.Println("init db connection", endPoint, "is bad")
		return nil
	}
	//fmt.Println("get engine is ok")
	// 虽然我们此处已经获取了engine，但并不能说明mysql是正常的，这个有xorm的个性导致的，xorm的engine并不会操作数据库，除非手动触发
	// 这里我们确认下mysql是否正常
	err1 := engine.Ping()
	if err1 != nil {
		//fmt.Println("mysql 宕机了")
		//fmt.Println("============================")
		return nil
	}
	return engine
}

func (g *Grafanadb)GetMonitHost(table string) []MonitHost {
	// 定义一个存放monithost的切片，类型是monithost
	hostSlice := []MonitHost{}
	engine := g.GetEngine()
	sql := fmt.Sprintf("select hostname, address, product, cluster_name, role, port from %s;", table)
	//sql := fmt.Sprintf("select hostname, address, product, cluster_name, role, port from %s where cluster_name = 'newdefault';", table)
	resultSlice, err1 := engine.QueryString(sql)
	if err1 != nil {
		g.Logger.Println("run sql", sql, "is bad, err1 is", err1)
	}
	for _, dict := range resultSlice {
		mh := MonitHost{}
		mh.Endpoint = dict["product"]+"_"+dict["cluster_name"]+"_"+dict["role"]+"_"+dict["port"]+"_"+dict["hostname"]
		mh.User = g.User
		mh.Passwd = g.Passwd
		mh.Address = dict["address"]
		port, err2 := strconv.Atoi(dict["port"])
		if err2 != nil {
			mh.Logger.Println("transfer port is bad, err2 is", err2)
		}
		mh.Port = port
		mh.Charset = g.Charset
		mh.Database = ""
		mh.Logger = g.Logger
		hostSlice = append(hostSlice, mh)
	}
	return hostSlice
}

// method
func (g *Grafanadb)GetEngine() *xorm.Engine {
	endPoint := fmt.Sprintf("%s:%d", g.Address, g.Port)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true", g.User, g.Passwd, endPoint, g.Database, g.Charset)
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		g.Logger.Println("init db connection", endPoint, "is bad")
		os.Exit(1)
	}
	return engine
}












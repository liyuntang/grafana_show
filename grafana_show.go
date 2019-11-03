package main

import (
	"flag"
	"grafana_show/common"
	"grafana_show/database"
	"grafana_show/monit"
	"grafana_show/tomlConfig"
	"sync"
	"time"
)

var (
	config string
	wg sync.WaitGroup
)

func init()  {
	flag.StringVar(&config, "config", "src/grafana_show/conf/monit.file", "configration file")
}

func main()  {
	flag.Parse()
	// 解析配置文件

	configration := tomlConfig.TomlCofig(config)
	// 初始化日志
	Logger := common.InitLog(configration.System.LogFile)


	// 获取监控时间间隔
	intervalTime := configration.System.MonitIntervalTime
	for {
		// 获取监控主机列表
		hostSlice := monit.GetMonitHost(configration, Logger)

		// 获取监控时间
		monitTime := time.Now().Unix()

		// 定义一个channel存放监控分析结果
		ch := make(chan map[string][]map[string]int, 100000)

		aaa := time.Now()
		//fmt.Println("monitTime is", monitTime)
		// 开始检测并发
		for _, host := range hostSlice {
			wg.Add(1)
			go monit.Monit(host, configration.System.MonitElement, ch, &wg)

		}
		wg.Wait()
		close(ch)
		//os.Exit(0)
		Logger.Println("get monit host time is", time.Since(aaa))
		bbb := time.Now()
		//将监控结果入库
		database.InsertTab(configration, Logger, monitTime, ch)
		Logger.Println("inser db time is", time.Since(bbb))
		// 休息片刻
		waitTime, err := time.ParseDuration(intervalTime)
		if err != nil {
			Logger.Println("parse interval time is bad, err is", err)
		}
		time.Sleep(waitTime)
	}

}



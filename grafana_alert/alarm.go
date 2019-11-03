package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	address string
	port int
)

func init()  {
	flag.StringVar(&address, "h", "0.0.0.0", "ip address")
	flag.IntVar(&port, "P", 8000, "port")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", alert)
	http.ListenAndServe(":8000", nil)


}

func alert(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(time.Now())
	var data AlterData
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err is", err)
	}
	//fmt.Println(string(buf))
	err1 := json.Unmarshal(buf, &data)
	if err1 != nil {
		fmt.Println("json data is bad, err1 is", err1)
	}
	//fmt.Println(data)
	if len(data.EvalMatches) == 0 {
		alertInfo := fmt.Sprintf("没有抓到数据")
		baojing(alertInfo)
	} else {
		for _, monitData := range data.EvalMatches {
			metric := monitData["metric"]
			value := monitData["value"]
			alertInfo := fmt.Sprintf("状态：异常,报警项：%s,主机：%s,当前值：%f,报警规则：%s", data.RuleName, metric, value, data.Message)
			baojing(alertInfo)
		}
	}

}

func baojing(content string)  {
	url:= fmt.Sprintf("http://smscenter.niceprivate.com/smscenter.php?g=11&p=jqYgH2OhDk&&desc=[Mysql/all]all:%s&deploy_path=/&type=1",content)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		by,_ :=ioutil.ReadAll(res.Body)
		fmt.Println(string(by))
	}
}

type AlterData struct {
	EvalMatches []map[string]interface{}
	Message string
	RuleId int
	RuleName string
	RuleUrl string
	Title string
}













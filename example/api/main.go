package main

import (
	"github.com/astaxie/beego"
	"github.com/slene/iploc"
	"path/filepath"
	"os"
)


type MainController struct {
	beego.Controller
}

func (m *MainController) Get() {
	m.Post()
}

func (m *MainController) Post() {
	result := map[string]interface{}{"success":false}
	ip := m.Ctx.Request.Form.Get("ip")
	ipinfo, err := iploc.GetIpInfo(ip)
	if err != nil {
		result["msg"] = err.Error()
	} else {
		result["success"] = true
		result["ip"] = ip
		result["flag"] = ipinfo.Flag
		result["code"] = ipinfo.Code
		result["country"] = ipinfo.Country
		result["region"] = ipinfo.Region
		result["city"] = ipinfo.City
		result["isp"] = ipinfo.Isp
		if ipinfo.Flag != iploc.FLAG_INUSE {
			result["note"] = ipinfo.Note
		}
	}
	m.Data["json"] = result
	m.ServeJson()
}

func init() {
	pwd, _ := os.Getwd()
	iplocFilePath, _ := filepath.Abs(filepath.Join(pwd, "src/github.com/slene/iploc/iploc.dat"))
	iploc.IpLocInit(iplocFilePath)
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"net/http"
	"strings"
	"github.com/devfeel/dotweb"
	"github.com/zhujq/phonedata"
)

type App struct {
	Web      *dotweb.DotWeb
}


type ResBody struct {
	Status      string `json:"status"`
	PhoneNum   string `json:"PhoneNum"`
	Province     string `json:"Province"`
	City   string `json:"City "`
	ZipCode    string `json:"ZipCode"`
	AreaZone string `json:"AreaZone"`
	CardType string `json:"CardType"`
}


func NewApp() *App {
	var a = &App{}
	a.Web = dotweb.New()
	return a
}

var app = NewApp()


func main() {
	var err error

	var (
		version = flag.Bool("version", false, "version v1.0")
		port    = flag.Int("port", 8080, "listen port.")
	)

	flag.Parse()

	if *version {
		fmt.Println("v1.0")
		os.Exit(0)
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	InitRoute(app.Web.HttpServer)
	log.Println("Start China Phone Query Server on ", *port)
	app.Web.StartServer(*port)
}

func indexHandler(ctx dotweb.Context) error {
	phonenum := ctx.QueryString("phonenum")

	if phonenum == "" {
		log.Println("ERROR: 没有提供电话号码")
		return ctx.WriteString("欢迎使用中国电话号码查询系统，请在网址后输入 /?phonenum=电话号码 查询，如https://phone.zhujq.ml/?phonenum=13988888888")
	}

	var message = ResBody{
		Status:      "failed",
		PhoneNum: "",
		Province: "",
		City: "",
		ZipCode: "",
		AreaZone: "",
		CardType: "",
	}

	log.Println(phonenum)	

	message.PhoneNum = phonenum

	if strings.HasPrefix(phonenum,"86"){

		phonenum = strings.TrimPrefix(phonenum,"86")

	}

	phoneresult, err := phonedata.Find(phonenum)

	if err != nil {	
		log.Println("error:", err)	
		return ctx.WriteJsonC(http.StatusNotFound, message)
	}
	
	message.Status = "success"
	message.Province = phoneresult.Province
	message.City = phoneresult.City
	message.ZipCode = phoneresult.ZipCode
	message.AreaZone = phoneresult.AreaZone
	message.CardType = phoneresult.CardType
	
	return ctx.WriteJson(message)	
}

func InitRoute(server *dotweb.HttpServer) {
	server.GET("/", indexHandler)
}


package main

import (
	"net/http"
	"youxi/service"

	"youxi/router"
	"youxi/pkg/setting"
	"youxi/models"
	"fmt"
)

func init() {
	setting.Setup()
	models.Setup()
	service.CronStart()
}


func main() {


	router := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerConfig.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerConfig.ReadTimeout,
		WriteTimeout:   setting.ServerConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}


package main

import (
	"fmt"
	"github.com/rongqinzheng/go-application-go-gin/go-gin/pkg/setting"
	"github.com/rongqinzheng/go-application-go-gin/go-gin/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

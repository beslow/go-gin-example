package main

import (
	"fmt"
	"net/http"

	"github.com/beslow/go-gin-example/pkg/setting"

	"github.com/beslow/go-gin-example/routers/router"
)

func main() {
	router := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

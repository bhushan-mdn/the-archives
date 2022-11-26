package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhushan-mdn/the-archives/models"
	"github.com/bhushan-mdn/the-archives/pkg/config"
	"github.com/bhushan-mdn/the-archives/pkg/storage"
	"github.com/bhushan-mdn/the-archives/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Setup()
	models.Setup()
	storage.Setup()
}

func main() {
	gin.SetMode(config.Server.Mode)

	routersInit := routers.InitRouter()
	readTimeout := config.Server.ReadTimeout
	writeTimeout := config.Server.WriteTimeout
	port := fmt.Sprintf(":%d", config.Server.Port)
	maxHeaderBytes := 1 << 20 // 1 MB

	server := &http.Server{
		Addr:           port,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[INFO] HTTP server listening on %s", port)

	server.ListenAndServe()

}

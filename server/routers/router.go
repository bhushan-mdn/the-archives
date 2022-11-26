package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/login", login)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", upload)
	r.GET("/health", CheckHealth)

	r.POST("/blob", CreateBlob)
	r.GET("/blob/:id", GetBlob)

	r.POST("/files", CreateFile)
	r.GET("/files", GetAllFiles)
	r.HEAD("/file/:id", GetFile)
	r.GET("/file/:id", GetFile)
	r.PUT("/file/:id", UpdateFile)
	r.DELETE("/file/:id", DeleteFile)

	return r
}

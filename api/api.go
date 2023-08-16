package api

import (
	_ "api_gateway/api/docs"

	"api_gateway/api/handlers"
	"api_gateway/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	r.Use(customCORSMiddleware())

	// book
	r.POST("/books", h.CreateBook)
	r.GET("/books", h.GetBookList)
	// r.GET("/book/:id", h.GetBookByID)
	r.GET("/books/:title", h.GetBookByTitle)
	r.PUT("/books/:id", h.UpdateBook)
	r.PATCH("/books/:id", h.UpdatePatchBook)
	r.DELETE("/books/:id", h.DeleteBook)

	//user service
	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUserById)
	r.GET("/myself", h.GetMyself)
	r.GET("/user", h.AuthMiddleware(), h.GetUserList)

	//login
	r.POST("/login", h.Login)
	r.POST("/signup", h.Register)
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Acces-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSF-TOKEN, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

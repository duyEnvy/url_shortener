package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"url-shortener/component/appctx"
	"url-shortener/module/url/transport/ginurl"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db = db.Debug()

	appContext := appctx.NewAppContext(db)

	r := gin.Default()

	r.GET("/:short_code", ginurl.FindUrl(appContext))
	v1 := r.Group("/api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		}) // listen and serve on
	})
	shortenUrl := v1.Group("/shorten-url")

	shortenUrl.POST("", ginurl.CreateUrl(appContext))

	err = r.Run()
	if err != nil {
		log.Fatalln(err)
	} // listen and serve on 0.0.0.0:8080
}

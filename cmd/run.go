package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"SnowFlakeGo/service"
)

func main(){
	sfWorker, err := service.NewWorker(int64(1))
	if err != nil {
		fmt.Println(err)
        return
	}

	r := gin.Default()
	r.GET("/getid", func(c *gin.Context){
		id := sfWorker.GetId()
		fmt.Println(id)
		c.JSON(200, gin.H{
			"message" : "ok",
			"data" : id,
		})
	})
	r.Run()
}
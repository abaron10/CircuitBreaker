
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restAPI/server/domain"
	"strconv"
)

func getTask(c * gin.Context){
	c.JSON(200, domain.Tasks)

}

func addTask(c * gin.Context){
	var task domain.Task
	list := &domain.Tasks
	decodingError := c.BindJSON(&task)
	if decodingError != nil{
		fmt.Println("error decoding json")
	}
	(*list) = append((*list),task)
	c.JSON(200, domain.Tasks)
}

func searchTask(c * gin.Context){
	queryParamId,ok := c.GetQuery("key")
	if !ok {
		fmt.Println("Non existing query param")
	}
	id , err := strconv.Atoi(queryParamId)
	if err != nil {
		fmt.Println("Error casting param")
	}

	var response domain.Task
	for _,task := range domain.Tasks {
		if(task.ID == id){
			response = task
			break
		}
	}
	c.JSON(200, response)
}


func main(){
	r := gin.Default()
	r.GET("/ping", test)
	r.GET("/getTasks", getTask)
	r.POST("/createTask", addTask)
	r.GET("/searchTask", searchTask)
	r.Run(":8089")
	t := domain.Task{}
	fmt.Println(t)
}

func test(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
package controllers

import (
	"gupta/first/db"
	"gupta/first/types"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Tasks(c *gin.Context) {
	c.JSON(200, db.Tasks)
}

func PostTask(c *gin.Context) {
	var task types.Task
	id := uuid.New().ID()
	task.IsDone = false
	c.BindJSON(&task)
	task.ID = int(id)
	db.Tasks = append(db.Tasks, task)
	c.JSON(200, db.Tasks)
}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func PutTask(c *gin.Context) {
	var task types.Task
	id,_ := strconv.Atoi(c.Param("id"))
	for	i, t := range db.Tasks {
		if t.ID == id {
			c.BindJSON(&task)
			db.Tasks[i] =	task
			c.JSON(200, db.Tasks)
			return
		}
	}
}
package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"toDoList/pkg/utils"
	"toDoList/service"
)

func ShowTasks(c *gin.Context) {
	var listTasks service.ListTaskService

	//先进行身份验证
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTasks); err == nil {
		res := listTasks.List(claim.Id) //后面的id为备忘录的id,前面的id为用户id可有cookie获得
		c.JSON(200, res)
	} else {
		//有错误的话，返回并打印日志
		logging.Error(err)
		c.JSON(400, err)
	}
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService

	//先进行身份验证
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(claim.Id, c.Param("id")) //后面的id为备忘录的id,前面的id为用户id可有cookie获得
		c.JSON(200, res)
	} else {
		//有错误的话，返回并打印日志
		logging.Error(err)
		c.JSON(400, err)
	}
}

func TasksCreate(c *gin.Context) {
	var createTask service.CreateTasksService

	//先进行身份验证
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		//有错误的话，返回并打印日志
		logging.Error(err)
		c.JSON(400, err)
	}
}

func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTasksService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(claim.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//先进行身份验证
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claim.Id)
		c.JSON(200, res)
	} else {
		//有错误的话，返回并打印日志
		logging.Error(err)
		c.JSON(400, err)
	}
}

func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(claim.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

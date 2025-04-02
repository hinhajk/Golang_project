package service

import (
	"time"
	"toDoList/models"
	"toDoList/serializer"
)

type CreateTasksService struct {
	TaskName  string `form:"tasks_name"  json:"tasks_name"`
	User      models.User
	Uid       uint
	Title     string `gorm:"index;not null"`
	Status    int    `json:"status" form:"status"`
	Content   string
	StartTime int
	EndTime   int
}

// 业务逻辑
func (createTask *CreateTasksService) Create(id uint) serializer.Response {
	var user models.User
	// 先找到用户
	models.DB.First(&user, id)
	task := models.Tasks{
		User:      user,
		Uid:       user.ID,
		TaskName:  createTask.TaskName, //这些数据是从前端传过来的
		Title:     createTask.Title,
		Status:    0,
		Content:   createTask.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := models.DB.Create(&task).Error //创建
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "创建失败",
		}
	}
	return serializer.Response{
		Status:  200,
		Message: "创建成功",
	}
}

type ShowTaskService struct {
}

func (showTask *ShowTaskService) Show(uid uint, tid string) serializer.Response {
	var task models.Tasks
	err := models.DB.First(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "查询失败",
		}
	}
	return serializer.Response{
		Status:  200,
		Data:    serializer.BuildTask(task),
		Message: "内容如下",
	}

}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (listTasks *ListTaskService) List(uid uint) serializer.Response {
	var task []models.Tasks
	var count int64 = 0
	if listTasks.PageSize == 0 {
		listTasks.PageSize = 10
	}
	//prelod将外键预加载出来，再进行分页操作
	models.DB.Model(&models.Tasks{}).Preload("User").Where("uid = ?", uid).Count(&count).Limit(listTasks.PageSize).Offset((listTasks.PageNum - 1) * listTasks.PageSize).Find(&task)

	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTasks(task),
		Count:  count,
	}
}

type UpdateTasksService struct {
	TaskName  string `form:"tasks_name"  json:"tasks_name"`
	User      models.User
	Uid       uint
	Title     string `gorm:"index;not null"`
	Status    int    `json:"status" form:"status"`
	Content   string
	StartTime int
	EndTime   int
}

func (updateTask *UpdateTasksService) Update(uid uint, tid string) serializer.Response {
	//首先查找到这个任务
	var task models.Tasks
	models.DB.First(&task, tid)
	task.Content = updateTask.Content
	task.Status = updateTask.Status
	task.Title = updateTask.Title
	models.DB.Save(&task)
	return serializer.Response{
		Status:  200,
		Data:    serializer.BuildTask(task),
		Message: "更新完成",
	}
}

type SearchTaskService struct {
	Info     string `form:"info" json:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

func (searchTasks *SearchTaskService) Search(uid uint) serializer.Response {
	var task []models.Tasks
	var count int64 = 0
	if searchTasks.PageSize == 0 {
		searchTasks.PageSize = 10
	}
	//prelod将外键预加载出来，再进行分页操作
	models.DB.Model(&models.Tasks{}).Preload("User").Where("uid = ?", uid).
		Where("title LIKE ? OR content LIKE ?", searchTasks.Info+"%", "%"+searchTasks.Info+"%").
		Count(&count).Limit(searchTasks.PageSize).Offset((searchTasks.PageNum - 1) * searchTasks.PageSize).Find(&task)

	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTasks(task),
	}
}

type DeleteTaskService struct {
}

func (deleteTask *DeleteTaskService) Delete(uid uint, tid string) serializer.Response {
	var task models.Tasks
	models.DB.First(&task, tid)
	models.DB.Delete(&task)
	return serializer.Response{
		Status:  200,
		Message: "删除成功",
	}
}

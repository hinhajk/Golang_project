package serializer

import "toDoList/models"

type Task struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title" example:"吃饭"`
	Content   string `json:"content" example:"睡觉"` //任务内容，longtext代表内容字段较长
	View      uint64 `json:"view" example:"1"`
	CreateAt  int64  `json:"create_at" example:"1"`
	StartTime int64  `json:"start_time" example:"1"`
	EndTime   int64  `json:"end_time" example:"1"`
	Status    int    `json:"status" example:"0"`
}

func BuildTask(task models.Tasks) Task {
	return Task{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
		CreateAt:  task.CreatedAt.Unix(),
	}
}

func BuildTasks(tasks []models.Tasks) []Task {
	result := make([]Task, len(tasks))
	for i, task := range tasks {
		result[i] = BuildTask(task)
	}
	return result
}

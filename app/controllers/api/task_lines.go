package api

import (
	"openRPA-log-server/app/models"

	"openRPA-log-server/app/manager"
)

type TaskLines struct{}

type TaskLine struct {
	TaskRequests models.TaskRequestSlice `json:"tasks"`
}

func (ctrl TaskLines) Show(c Api) *Api {
	var taskLine TaskLine
	var tasks models.TaskSlice

	id, err := c.GetStrIDFromParam()
	if err != nil {
		return c.SetMessage("IDが不正です").HandleBadRequestError()
	}

	if err := models.DB.Find(&tasks, "task_line_id = ?", id).Error; err != nil {
		return c.SetMessage(err.Error()).HandleNotFoundError()
	}

	taskRequests, err := manager.ParseTasksToTaskRequests(tasks)

	if err != nil {
		return c.SetMessage("JSONのパースに失敗しました err:" + err.Error()).HandleInternalServerError()
	}

	taskLine.TaskRequests = taskRequests

	return c.RenderJSON(struct {
		TaskLine TaskLine `json:"task_line"`
	}{
		TaskLine: taskLine,
	})
}

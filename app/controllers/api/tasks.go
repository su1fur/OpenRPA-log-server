package api

import (
	"openRPA-log-server/app/models"
	"openRPA-log-server/app/lib/tool"
	"time"
	"openRPA-log-server/app/manager"
	"openRPA-log-server/app/validators"
)

type Tasks struct{}

func (ctrl Tasks) Index(c Api) *Api {
	var tasks models.TaskSlice

	if err := models.DB.Find(&tasks).Error; err != nil {
		return c.SetMessage(err.Error()).HandleBadRequestError()
	}

	taskRequests, err := manager.ParseTasksToTaskRequests(tasks)
	if err != nil {
		return c.SetMessage("JSONのパースに失敗しました err:" + err.Error()).HandleInternalServerError()
	}

	return c.RenderJSON(struct {
		Tasks models.TaskRequestSlice `json:"tasks"`
	}{
		Tasks: taskRequests,
	})
}

func (ctrl Tasks) Show(c Api) *Api {
	task := models.Task{}

	id, err := c.GetUint64IDFromParam()
	if err != nil {
		return c.SetMessage(err.Error()).HandleBadRequestError()
	}

	if err := models.DB.First(&task, id).Error; err != nil {
		return c.SetMessage(err.Error()).HandleNotFoundError()
	}

	taskRequest, err := manager.ParseTaskToTaskRequest(task)
	if err != nil {
		return c.SetMessage("JSONのパースに失敗しました err:" + err.Error()).HandleInternalServerError()
	}

	return c.RenderJSON(struct {
		Task models.TaskRequest `json:"task"`
	}{
		Task: taskRequest,
	})
}

func (ctrl Tasks) Create(c Api) *Api {

	taskRequest := models.TaskRequest{}
	if err := c.BindJSON(&taskRequest); err != nil {
		return c.SetMessage(err.Error()).HandleBadRequestError()
	}

	if taskRequest.ID != 0 {
		return c.UnSpecifiedId().HandleBadRequestError()
	}

	if err := validators.Struct(&taskRequest); err != nil {
		return c.SetMessage("不正なリクエストです。 err:"+err.Error()).HandleBadRequestError()
	}

	if taskRequest.TaskLineID == "" {
		taskRequest.TaskLineID = tool.GenerateSha256(time.Date(2215, 10, 10, 12, 3, 0, 0, time.UTC).Sub(time.Now()).String())
	}

	task, err := manager.ParseTaskRequestToTask(taskRequest)
	if err != nil {
		return c.SetMessage("JSONのパースに失敗しました err:" + err.Error()).HandleInternalServerError()
	}

	if err := models.DB.Create(&task).Error; err != nil {
		return c.SetMessage(err.Error()).HandleInternalServerError()
	}

	taskRequest, err = manager.ParseTaskToTaskRequest(task)
	if err != nil {
		return c.SetMessage("JSONのパースに失敗しました err:" + err.Error()).HandleInternalServerError()
	}

	return c.SetMessage("Task Create Success!!").RenderJSON(struct {
		Task models.TaskRequest `json:"task"`
	}{
		Task: taskRequest,
	})
}

func (ctrl Tasks) Update(c Api) *Api {

	var task models.Task

	id, err := c.GetUint64IDFromParam()
	if err != nil {
		return c.SetMessage(err.Error()).HandleBadRequestError()
	}

	if err := models.DB.First(&task, id).Error; err != nil {
		return c.SetMessage(err.Error()).HandleNotFoundError()
	}

	taskRequest := models.TaskRequest{}
	if err := c.BindJSON(&taskRequest); err != nil {
		return c.SetMessage(err.Error()).HandleBadRequestError()
	}

	if err := validators.Struct(&taskRequest); err != nil {
		return c.SetMessage("不正なリクエストです。 err:"+err.Error()).HandleBadRequestError()
	}

	if taskRequest.TaskLineID == "" {
		taskRequest.TaskLineID = tool.GenerateSha256(time.Date(2215, 10, 10, 12, 3, 0, 0, time.UTC).Sub(time.Now()).String())
	}

	task, err = manager.ParseTaskRequestToTask(taskRequest)
	if err != nil {
		return c.SetMessage("JSONのパースに失敗しました err:" + err.Error()).HandleInternalServerError()
	}

	if err := models.DB.Save(&task).Error; err != nil {
		return c.SetMessage(err.Error()).HandleInternalServerError()
	}

	return ctrl.Show(c)
}

func (ctrl Tasks) Delete(c Api) *Api {
	task := models.Task{}

	id, err := c.GetUint64IDFromParam()
	if err != nil {
		return c.SetMessage(err.Error()).HandleBadRequestError()
	}

	if err := models.DB.First(&task, id).Error; err != nil {
		return c.SetMessage(err.Error()).HandleNotFoundError()
	}

	if err := models.DB.Delete(&task).Error; err != nil {
		return c.SetMessage(err.Error()).HandleInternalServerError()
	}

	return c.Success()
}

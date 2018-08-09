package manager

import (
	"openRPA-log-server/app/models"
	"openRPA-log-server/app/lib/tool"
)

func ParseTasksToTaskRequests(tasks models.TaskSlice) (taskRequests models.TaskRequestSlice, err error) {

	for _, v := range tasks {

		var taskRequest models.TaskRequest
		var err error
		taskRequest.Input, err = tool.ParseInterfaceFromStringJSON(v.Input)
		taskRequest.Arguments, err = tool.ParseInterfaceFromStringJSON(v.Arguments)
		taskRequest.Returns, err = tool.ParseInterfaceFromStringJSON(v.Returns)
		taskRequest.Model = v.Model
		taskRequest.TaskLineID = v.TaskLineID
		taskRequest.Status = v.Status
		taskRequest.EndPoint = v.EndPoint

		if err != nil {
			return models.TaskRequestSlice{}, err
		}

		taskRequests = append(taskRequests, taskRequest)
	}

	return taskRequests, nil
}

func ParseTaskToTaskRequest(task models.Task) (taskRequest models.TaskRequest, err error) {

	taskRequest.Input, err = tool.ParseInterfaceFromStringJSON(task.Input)
	taskRequest.Arguments, err = tool.ParseInterfaceFromStringJSON(task.Arguments)
	taskRequest.Returns, err = tool.ParseInterfaceFromStringJSON(task.Returns)
	taskRequest.Model = task.Model
	taskRequest.TaskLineID = task.TaskLineID
	taskRequest.Status = task.Status
	taskRequest.EndPoint = task.EndPoint

	if err != nil {
		return models.TaskRequest{}, err
	}

	return taskRequest, nil
}

func ParseTaskRequestToTask(taskRequest models.TaskRequest) (task models.Task, err error) {

	task.Input, err = tool.ParseStringJSONFromInterface(taskRequest.Input)
	task.Arguments, err = tool.ParseStringJSONFromInterface(taskRequest.Input)
	task.Returns, err = tool.ParseStringJSONFromInterface(taskRequest.Input)
	task.EndPoint = taskRequest.EndPoint
	task.TaskLineID = taskRequest.TaskLineID
	task.Status = taskRequest.Status
	task.Model = taskRequest.Model

	if err != nil {
		return models.Task{},err
	}

	return task, nil
}

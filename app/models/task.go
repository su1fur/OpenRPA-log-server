package models

import "sort"

type Task struct {
	Model
	TaskLineID string `json:"task_line_id"`
	EndPoint   string `json:"end_point"`
	Status     string `json:"status"`
	Input      string `json:"input" gorm:"type:varchar(2048)"`
	Arguments  string `json:"arguments" gorm:"type:varchar(2048)"`
	Returns    string `json:"returns" gorm:"type:varchar(2048)"`
}

type TaskRequest struct {
	Model
	TaskLineID string      `json:"task_line_id" validator:"existTaskLine"`
	EndPoint   string      `json:"end_point" validator:"url"`
	Status     string      `json:"status" validate:"eq=started|eq=processing|eq=succeeded|eq=failed"`
	Input      interface{} `json:"input"`
	Arguments  interface{} `json:"arguments"`
	Returns    interface{} `json:"returns"`
}

/**
TaskのSlice
*/
type TaskSlice []Task

func (p TaskSlice) Len() int           { return len(p) }
func (p TaskSlice) Less(i, j int) bool { return p[i].ID > p[j].ID }
func (p TaskSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p TaskSlice) Sort()              { sort.Sort(p) }

/**
TaskRequestのSlice
*/
type TaskRequestSlice []TaskRequest

func (p TaskRequestSlice) Len() int           { return len(p) }
func (p TaskRequestSlice) Less(i, j int) bool { return p[i].ID > p[j].ID }
func (p TaskRequestSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p TaskRequestSlice) Sort()              { sort.Sort(p) }

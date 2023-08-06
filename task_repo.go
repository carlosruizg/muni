package main

import "github.com/carlosruizg/muni/graph/model"

type TaskRepo struct {
	Tasks []*model.LabellingTask
}

func NewTaskRepo() *TaskRepo {
	example := "example"
	tasks := []*model.LabellingTask{
		{
			ID:      "1",
			Title:   "Summarization",
			Example: &example,
			User: &model.User{
				ID:   "1",
				Name: "peter",
			},
		},
		{
			ID:      "2",
			Title:   "Text Extraction",
			Example: &example,
			User: &model.User{
				ID:   "1",
				Name: "peter",
			},
		},
		{
			ID:      "3",
			Title:   "Chat",
			Example: &example,
			User: &model.User{
				ID:   "1",
				Name: "peter",
			},
		},
	}
	return &TaskRepo{
		Tasks: tasks,
	}
}

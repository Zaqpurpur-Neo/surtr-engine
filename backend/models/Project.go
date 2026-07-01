package models

import "surtr-engine/commons"

type Project struct {
	Id       string                `json:"id"`
	Name     string                `json:"name"`
	Category string                `json:"category"`
	Status   commons.StatusProject `json:"status"`
}

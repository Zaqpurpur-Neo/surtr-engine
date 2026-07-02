package models

import "surtr-engine/commons"

type Project struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Status      commons.StatusProject `json:"status"`
	StartAt     *string               `json:"startAt"`
	EndAt       *string               `json:"endAt"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	CategoryID string `json:"categoryId"`
}

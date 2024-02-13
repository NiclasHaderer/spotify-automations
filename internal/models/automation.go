package models

type Automation[T any] struct {
	Name   string `json:"name"`
	Config T      `json:"config"`
}

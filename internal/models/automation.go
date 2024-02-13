package models

type Automation[T any] struct {
	Config T `json:"config"`
}

type AutomationOption struct {
	Name           string
	CreateOrModify func()
	Run            func()
}

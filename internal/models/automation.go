package models

type Automation struct {
	Config string `json:"config"`
}

type AutomationOption struct {
	Name           string
	CreateOrModify func()
	Run            func()
}

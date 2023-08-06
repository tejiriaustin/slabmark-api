package models

type LabReading struct {
	Shared
	AccountInfo `json:"account_info" bson:"account_info"`
}

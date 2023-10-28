package models

import "time"

type Activity struct {
	Shared      `bson:",inline"`
	ActionTaken string      `bson:"action_taken"`
	Time        time.Time   `bson:"time"`
	AccountInfo AccountInfo `bson:"account_info"`
}

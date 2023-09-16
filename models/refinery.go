package models

import "time"

type (
	HourlyReport struct {
		TakeTime  *time.Time `json:"take_time" bson:"take_time"`
		FlowRate  string     `json:"flow_rate" bson:"flow_rate"`
		FfaOfRbdo string     `json:"ffa_of_rbdo" bson:"ffa_of_rbdo"`
		FfaOfDfa  string     `json:"ffa_of_dfa" bson:"ffa_of_dfa"`
		FfaOfSpo  string     `json:"k" bson:"ffa_of_spo"`
	}
	RefineryReport struct {
		Shared         `bson:",inline"`
		PlantSituation string         `json:"plant_situation" bson:"plant_situation"`
		AccountInfo    AccountInfo    `json:"account_info" bson:"account_info"`
		HourlyReports  []HourlyReport `json:"hourly_reports" bson:"hourly_reports"`
	}
)

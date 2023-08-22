package models

import "time"

type (
	HourlyQualityReadings struct {
		Shared        `bson:",inline"`
		TimeOfReading *time.Time `json:"time" bson:"time"`
		D4            string     `json:"d4" bson:"d4"`
		D1            string     `json:"d1" bson:"d1"`
		DFA           string     `json:"dfa" bson:"dfa"`
		Remark        string     `json:"remark" bson:"remark"`
	}

	DailyQualityReadings struct {
		Shared        `bson:",inline"`
		ProductCode   string      `json:"product_code" bson:"product_code"`
		OverallRemark string      `json:"overall_remark" bson:"overall_remark"`
		AccountInfo   AccountInfo `json:"account_info" bson:"account_info"`
		IdsOfReadings []string    `json:"ids_of_readings" bson:"ids_of_readings"`
	}
)

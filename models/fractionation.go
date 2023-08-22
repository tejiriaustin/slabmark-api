package models

type (
	ResumptionStock struct {
		OleinOne    string `json:"olein_one" bson:"olein_one"`
		OleinTwo    string `json:"olein_two" bson:"olein_two"`
		Product     string `json:"product" bson:"product"`
		Stearin     string `json:"stearin" bson:"stearin"`
		Filtration  string `json:"filtration" bson:"filtration"`
		FastCooling string `json:"fast_cooling" bson:"fast_cooling"`
	}
	ClosingStock struct {
		OleinOne    string `json:"olein_one" bson:"olein_one"`
		OleinTwo    string `json:"olein_two" bson:"olein_two"`
		Product     string `json:"product" bson:"product"`
		Stearin     string `json:"stearin" bson:"stearin"`
		Filtration  string `json:"filtration" bson:"filtration"`
		FastCooling string `json:"fast_cooling" bson:"fast_cooling"`
	}
	Batch struct {
		NumberOfCycles  int    `json:"number_of_cycles" bson:"number_of_cycles"`
		BatchNumber     string `json:"batch_number" bson:"batch_number"`
		OleinQuantity   string `json:"olein_quantity" bson:"olein_quantity"`
		OleinYield      string `json:"olein_yield" bson:"olein_yield"`
		StearinQuantity string `json:"stearin_quantity" bson:"stearin_quantity"`
		StearinYield    string `json:"stearin_yield" bson:"stearin_yield"`
	}
	Filtration struct {
		BatchOne Batch `json:"batch_one" bson:"batch_one"`
		BatchTwo Batch `json:"batch_two" bson:"batch_two"`
	}
	LoadingBatch struct {
		CRBatchNumber       string `json:"cr_batch_number" bson:"cr_batch_number"`
		MeterReading        string `json:"meter_reading" bson:"meter_reading"`
		CrystallizerReading string `json:"crystallizer_reading" bson:"crystallizer_reading"`
	}
	Loading struct {
		LoadingBatchOne LoadingBatch `json:"loading_batch_one" bson:"loading_batch_one"`
		LoadingBatchTwo LoadingBatch `json:"loading_batch_two" bson:"loading_batch_two"`
	}
	FractionationReport struct {
		Shared          `bson:",inline"`
		Status          string          `json:"status" bson:"status"`
		ResumptionStock ResumptionStock `Json:"resumption_stock" bson:"resumption_stock"`
		ClosingStock    ClosingStock    `Json:"closing_stock" bson:"closing_stock"`
		Filtration      Filtration      `json:"filtration" bson:"filtration"`
		Loading         Loading         `json:"loading" bson:"loading"`
		AccountInfo     *AccountInfo    `json:"account_info" bson:"account_info"`
	}
)

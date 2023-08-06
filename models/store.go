package models

type StoreItem struct {
	Shared
	AccountInfo `json:"account_info" bson:"account_info"`
	ItemName    string `json:"item_name" bson:"item_name"`
	Quantity    int    `json:"quantity" bson:"quantity"`
}

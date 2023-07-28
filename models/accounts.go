package models

type AccountKind string

const (
	AdminAccountKind    AccountKind = "SLABMARK.ACCOUNT.KIND.ADMIN"
	EmployeeAccountKind AccountKind = "SLABMARK.ACCOUNT.EMPLOYEE"
)

type Status string

const (
	ActiveStatus    Status = "ACTIVE"
	SuspendedStatus Status = "SUSPENDED"
)

type Account struct {
	General
	FirstName   string      `json:"first_name" bson:"first_name"`
	LastName    string      `json:"last_name" bson:"last_name"`
	Phone       string      `json:"phone" bson:"phone"`
	Email       string      `json:"email" bson:"email"`
	AccountKind AccountKind `json:"account_kind" bson:"account_kind"`
	Role        string      `json:"role" bson:"role"`
	Status      Status      `json:"status" bson:"status"`
}

type Role struct {
	General
	RoleName string `json:"role_name" bson:"role_name"`
}

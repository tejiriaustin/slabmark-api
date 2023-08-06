package models

type (
	Status string // Status of an account type

	Kind string // Kind of account
)

const (
	ActiveStatus    Status = "ACTIVE"
	SuspendedStatus Status = "SUSPENDED"

	AdminAccountKind    Kind = "SLABMARK.ACCOUNT.KIND.ADMIN"
	EmployeeAccountKind Kind = "SLABMARK.ACCOUNT.KIND.EMPLOYEE"
)

type (
	Role struct {
		Shared
		RoleName string `json:"role_name" bson:"role_name"`
	}

	Account struct {
		Shared
		FirstName string `json:"first_name" bson:"first_name"`
		LastName  string `json:"last_name" bson:"last_name"`
		FullName  string `json:"full_name" bson:"full_name"`
		Phone     string `json:"phone" bson:"phone"`
		Email     string `json:"email" bson:"email"`
		Kind      Kind   `json:"kind" bson:"kind"`
		Role      Role   `json:"role" bson:"role"`
		Status    Status `json:"status" bson:"status"`
	}
)

type AccountInfo struct {
	Id        string `json:"id" bson:"id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	FullName  string `json:"full_name"`
	Kind      Kind   `json:"kind" bson:"kind"`
	Role      Role   `json:"role" bson:"role"`
}
type AccountInterface interface {
	GetFullName() string
}

func NewAccount() AccountInterface {
	return &Account{}
}

func (a Account) GetFullName() string {
	return a.FirstName + " " + a.LastName
}

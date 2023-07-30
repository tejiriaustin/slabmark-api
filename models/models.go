package models

import (
	"time"

	"github.com/google/uuid"
)

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

type Shared struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (s Shared) InitObject() {
	s.ID = uuid.New()
	now := time.Now().UTC()
	s.CreatedAt = &now
}

type Account struct {
	Shared
	FirstName   string      `json:"first_name" bson:"first_name"`
	LastName    string      `json:"last_name" bson:"last_name"`
	FullName    string      `json:"full_name"`
	Phone       string      `json:"phone" bson:"phone"`
	Email       string      `json:"email" bson:"email"`
	AccountKind AccountKind `json:"account_kind" bson:"account_kind"`
	Role        Role        `json:"role" bson:"role"`
	Status      Status      `json:"status" bson:"status"`
}
type AccountInfo struct {
	Id        uuid.UUID `json:"id" bson:"id"`
	FirstName string    `json:"first_name" bson:"first_name"`
	LastName  string    `json:"last_name" bson:"last_name"`
	FullName  string    `json:"full_name"`
	Role      Role      `json:"role" bson:"role"`
}

type Role struct {
	Shared
	RoleName string `json:"role_name" bson:"role_name"`
}

type StoreItem struct {
	Shared
	AccountInfo `json:"account_info" bson:"account_info"`
	ItemName    string `json:"item_name" bson:"item_name"`
	Quantity    int    `json:"quantity" bson:"quantity"`
}

type LabReading struct {
	Shared
	AccountInfo `json:"account_info" bson:"account_info"`
}

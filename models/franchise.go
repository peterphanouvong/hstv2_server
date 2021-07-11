package models

import "time"

type Franchise struct {
	ID *int `json:"id"`
	FranchiseeId *int `json:"franchisee_id"`
	Email *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	CreatedAt *time.Time `json:"created_at"`
	AddressId *int `json:"address_id"`
}

type FranchiseList struct {
	Franchises []Franchise `json:"franchises"`
}
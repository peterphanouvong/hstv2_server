package models

import (
	"time"

	"github.com/jackc/pgtype"
)

type User struct {
	ID *int `json:"id"`
	FirstName *string `json:"first_name"`
	MiddleName *string `json:"middle_name"`
	LastName *string `json:"last_name"`
	DOB pgtype.Date `json:"d_o_b"`
	PhoneNumber *string `json:"phone_number"`
	AddressId *int `json:"address_id"`
	CreatedAt *time.Time `json:"created_at"`
	CognitoId *string `json:"cognito_id"`
	Email *string `json:"email"`
	UserTypeId *int `json:"user_type_id"`
}

type UserList struct {
	Users []User `json:"users"`
}
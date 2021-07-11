package models

import (
	"net/http"
	"time"

	"github.com/jackc/pgtype"
)
type Person struct {
	ID *int `json:"id"`
	FirstName *string `json:"first_name"`
	MiddleName *string `json:"middle_name"`
	LastName *string `json:"last_name"`
	DOB pgtype.Date `json:"d_o_b"`
	CreatedAt *time.Time `json:"created_at"`
	CognitoId *string `json:"cognito_id"`
}

type PersonList struct {
	People []Person `json:"people"`
}

func (i *Person) Bind(r *http.Request) error {
	return nil
}

func (*PersonList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Person) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

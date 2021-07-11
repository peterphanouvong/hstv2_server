package models

import (
	"net/http"
)

type Address struct {
	ID *int `json:"id"`
	Line1 *string `json:"line_1"`
	Line2 *string `json:"line_2"`
	City *string `json:"city"`
	State *string `json:"state"`
	Postcode *string `json:"postcode"`
	Country *string `json:"country"`
}

type AddressList struct {
	People []Person `json:"people"`
}

func (i *Address) Bind(r *http.Request) error {
	return nil
}

func (*AddressList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Address) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
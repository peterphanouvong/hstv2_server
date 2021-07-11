package db

import (
	"context"
	"fmt"

	"github.com/peterphanouvong/hst/models"
)

func (db Database) AddAddress(addr models.Address) (*models.Address, error) {
	query := `
		insert into address (line_1, line_2, city, state, postcode, country)
		values ($1, $2, $3, $4, $5, $6)
		returning *
	`
	err := db.Conn.QueryRow(context.Background(), query, addr.Line1, addr.Line2, addr.City, addr.State, addr.Postcode, addr.Country).Scan(
		&addr.ID,
		&addr.Line1,
		&addr.Line2,
		&addr.City,
		&addr.State,
		&addr.Postcode,
		&addr.Country,
	)
	fmt.Println(err);

	return &addr, err;
}

func (db Database) GetAddress(id string) (*models.Address, error) {
	var address = &models.Address{}
	
	query := `
		select id, line_1, line_2, city, state, postcode, country
		from address
		where id = $1
	`

	err := db.Conn.QueryRow(context.Background(), query, id).Scan(
		&address.ID,
		&address.Line1,
		&address.Line2,
		&address.City,
		&address.State,
		&address.Postcode,
		&address.Country,
	)

	return address, err
}

func (db Database) UpdateAddress(address *models.Address) (*models.Address, error) {
	query := `
		update address
		set
			line_1 = $1,
			line_2 = $2,
			city = $3,
			state = $4,
			postcode = $5
		where id = $6
		returning *
	`

	var updatedAddress = models.Address{}
	err := db.Conn.QueryRow(context.Background(), query, address.Line1, address.Line2, address.City, address.State, address.Postcode, address.ID).Scan(
		&updatedAddress.ID,
		&updatedAddress.Line1,
		&updatedAddress.Line2,
		&updatedAddress.City,
		&updatedAddress.State,
		&updatedAddress.Postcode,
		&updatedAddress.Country,
	)

	return &updatedAddress, err
}
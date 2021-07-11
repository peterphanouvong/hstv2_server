package db

import (
	"context"
	"fmt"

	"github.com/peterphanouvong/hst/models"
)

func (db Database) AddFranchise(franchise models.Franchise) (*models.Franchise, error) {
	query := `
		insert into franchise (address_id, phone_number, email, franchisee_id)
		values ($1, $2, $3, $4)
		returning *
	`
	var newFranchise models.Franchise
	err := db.Conn.QueryRow(context.Background(), query, franchise.AddressId, franchise.PhoneNumber, franchise.Email, franchise.FranchiseeId).Scan(
		&newFranchise.ID,
		&newFranchise.AddressId,
		&newFranchise.PhoneNumber,
		&newFranchise.Email,
		&newFranchise.FranchiseeId,
		&newFranchise.CreatedAt,
	)
	fmt.Println(err);

	return &newFranchise, err;
}

func (db Database) GetFranchiseList() (*models.FranchiseList, error) {
	list := &models.FranchiseList{}
	rows, err := db.Conn.Query(context.Background(), "select * from franchise")

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var franchise models.Franchise
		err := rows.Scan(
			&franchise.ID,
			&franchise.AddressId,
			&franchise.PhoneNumber,
			&franchise.Email,
			&franchise.FranchiseeId,
			&franchise.CreatedAt,
		)

		if err != nil {
			return list, err
		}

		list.Franchises = append(list.Franchises, franchise)
	}
	return list, nil
}

// func (db Database) GetAddress(id string) (*models.Address, error) {
// 	var address = &models.Address{}
	
// 	query := `
// 		select id, line_1, line_2, city, state, postcode, country
// 		from address
// 		where id = $1
// 	`

// 	err := db.Conn.QueryRow(context.Background(), query, id).Scan(
// 		&address.ID,
// 		&address.Line1,
// 		&address.Line2,
// 		&address.City,
// 		&address.State,
// 		&address.Postcode,
// 		&address.Country,
// 	)

// 	return address, err
// }

// func (db Database) UpdateAddress(address *models.Address) (*models.Address, error) {
// 	query := `
// 		update address
// 		set
// 			line_1 = $1,
// 			line_2 = $2,
// 			city = $3,
// 			state = $4,
// 			postcode = $5
// 		where id = $6
// 		returning *
// 	`

// 	var updatedAddress = models.Address{}
// 	err := db.Conn.QueryRow(context.Background(), query, address.Line1, address.Line2, address.City, address.State, address.Postcode, address.ID).Scan(
// 		&updatedAddress.ID,
// 		&updatedAddress.Line1,
// 		&updatedAddress.Line2,
// 		&updatedAddress.City,
// 		&updatedAddress.State,
// 		&updatedAddress.Postcode,
// 		&updatedAddress.Country,
// 	)

// 	return &updatedAddress, err
// }
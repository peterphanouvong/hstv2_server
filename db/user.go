package db

import (
	"context"

	"github.com/peterphanouvong/hst/models"
)

func (db Database) GetAllUsers() (*models.UserList, error) {
	list := &models.UserList{}
    query := `
        select 
            p.id,
            first_name,
            middle_name,
            last_name,
            d_o_b,
            phone_number,
            address_id,
            created_at,
            u.cognito_id,
						email,
						user_type_id
        from users u
		join people p on u.id = p.id
    `
	rows, err := db.Conn.Query(context.Background(), query)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName, 
			&user.MiddleName, 
			&user.LastName, 
			&user.DOB, 
			&user.PhoneNumber, 
			&user.AddressId, 
			&user.CreatedAt,
			&user.CognitoId,
			&user.Email,
			&user.UserTypeId,
        )
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}

func (db Database) GetUserByCognitoId(cognitoId string) (*models.User, error) {
	var user = &models.User{}
	query := `
		select 
			p.id,
			first_name,
			coalesce(middle_name, '') as middle_name,
			last_name,
			coalesce(d_o_b, $2) as d_o_b,
			coalesce(phone_number, '') as phone_number,
			coalesce(address_id, -2) as address_id,
			created_at,
			p.cognito_id,
			email,
			user_type_id
		from people p
		join users u on u.id = p.id
		where p.cognito_id = $1
	`
	row := db.Conn.QueryRow(context.Background(), query, cognitoId, DEFAULT_DATE)
	err := row.Scan(
		&user.ID,
		&user.FirstName, 
		&user.MiddleName, 
		&user.LastName, 
		&user.DOB, 
		&user.PhoneNumber, 
		&user.AddressId, 
		&user.CreatedAt,
		&user.CognitoId,
		&user.Email,
		&user.UserTypeId,
	)

	return user, err
}

func (db Database) AddUser(user models.User) (*models.User, error) {
	var newUser models.User

	query := `
		insert into users (id, email, cognito_id, user_type_id)
		values (
			$1, 
			$2, 
			$3, 
			coalesce($4, 1000)
		)
		returning id
	`
	db.Conn.QueryRow(context.Background(), query, user.ID, user.Email, user.CognitoId, user.UserTypeId)

	query = `
		select 
			p.id,
			first_name,
			middle_name,
			last_name,
			d_o_b,
			phone_number,
			address_id,
			created_at,
			u.cognito_id,
			email,
			user_type_id
				from users u
			join people p on u.id = p.id
			where p.id = $1
	`

	err := db.Conn.QueryRow(context.Background(), query, user.ID).Scan(
		&newUser.ID,
		&newUser.FirstName, 
		&newUser.MiddleName, 
		&newUser.LastName, 
		&newUser.DOB, 
		&newUser.PhoneNumber, 
		&newUser.AddressId, 
		&newUser.CreatedAt,
		&newUser.CognitoId,
		&newUser.Email,
		&newUser.UserTypeId,
	)
	
	return &newUser, err;
}

func (db Database) GetUsersByType(user_type_id string) (*models.UserList, error) {
	list := &models.UserList{}
	query := `
		select 
			p.id,
			first_name,
			middle_name,
			last_name,
			d_o_b,
			phone_number,
			address_id,
			created_at,
			u.cognito_id,
			email,
			user_type_id
				from users u
			join people p on u.id = p.id
			where user_type_id = $1
	`
	rows, err := db.Conn.Query(context.Background(), query, user_type_id)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName, 
			&user.MiddleName, 
			&user.LastName, 
			&user.DOB, 
			&user.PhoneNumber, 
			&user.AddressId, 
			&user.CreatedAt,
			&user.CognitoId,
			&user.Email,
			&user.UserTypeId,
    )
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}

func (db Database) UpdateUser(user models.User) (*models.User, error) {

	var updatedUser models.User

	query := `
		update users
		set
			email = $1,
			phone_number = $2
		where id = $3
	`
	db.Conn.QueryRow(context.Background(), query, user.Email, user.PhoneNumber, user.ID)

	query = `
		select 
			p.id,
			first_name,
			middle_name,
			last_name,
			d_o_b,
			phone_number,
			address_id,
			created_at,
			u.cognito_id,
			email,
			user_type_id
				from users u
			join people p on u.id = p.id
			where p.id = $1
	`

	err := db.Conn.QueryRow(context.Background(), query, user.ID).Scan(
		&updatedUser.ID,
		&updatedUser.FirstName, 
		&updatedUser.MiddleName, 
		&updatedUser.LastName, 
		&updatedUser.DOB, 
		&updatedUser.PhoneNumber, 
		&updatedUser.AddressId, 
		&updatedUser.CreatedAt,
		&updatedUser.CognitoId,
		&updatedUser.Email,
		&updatedUser.UserTypeId,
	)
	
	return &updatedUser, err;
}

func (db Database) GetUserById(id string) (*models.User, error) {
	var user models.User
	query := `
		select 
			p.id,
			first_name,
			middle_name,
			last_name,
			d_o_b,
			phone_number,
			address_id,
			created_at,
			u.cognito_id,
			email,
			user_type_id
			from users u
			join people p on u.id = p.id
			where p.id = $1
	`

	err := db.Conn.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.MiddleName,
		&user.LastName,
		&user.DOB,
		&user.PhoneNumber,
		&user.AddressId,
		&user.CreatedAt,
		&user.CognitoId,
		&user.Email,
		&user.UserTypeId,
	)

	return &user, err

}
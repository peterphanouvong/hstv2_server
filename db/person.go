package db

import (
	"context"
	"fmt"

	//"fmt"
	//"time"

	"github.com/peterphanouvong/hst/models"
)

func (db Database) GetAllPersons() (*models.PersonList, error) {
	list := &models.PersonList{}
    query := `
        select 
            id,
            first_name,
            coalesce(middle_name, '') as middle_name,
            last_name,
            coalesce(d_o_b, $1) as d_o_b,
            created_at,
            coalesce(cognito_id, '') as cognito_id
        from people
    `
	rows, err := db.Conn.Query(context.Background(), query, DEFAULT_DATE)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var person models.Person
		err := rows.Scan(
            &person.ID,
            &person.FirstName, 
            &person.MiddleName, 
            &person.LastName, 
            &person.DOB, 
            &person.CreatedAt,
            &person.CognitoId,
        )
		if err != nil {
			return list, err
		}
		list.People = append(list.People, person)
	}
	return list, nil
}

func (db Database) GetPersonByCognitoId(cognitoId string) (models.Person, error) {
    person := models.Person{}
    query := `
        select 
            id,
            first_name,
            coalesce(middle_name, '') as middle_name,
            last_name,
            coalesce(d_o_b, $2) as d_o_b,
            created_at,
            cognito_id
        from people where cognito_id = $1
    `
    row := db.Conn.QueryRow(context.Background(), query, cognitoId, DEFAULT_DATE)
    err := row.Scan(
        &person.ID,
        &person.FirstName, 
        &person.MiddleName, 
        &person.LastName, 
        &person.DOB, 
        &person.CreatedAt,
        &person.CognitoId,
    )

    return person, err
}

func (db Database) AddPerson(person *models.Person) (*models.Person, error) {
    newPerson := models.Person{}

    query := `
        insert into people (first_name, middle_name, last_name, d_o_b, cognito_id)
        values (
            $1,
            $2,
            $3,
            $4,
            $5
        )
        returning id, first_name, middle_name, last_name, d_o_b, cognito_id, created_at
    `
    
    // fmt.Println(person.DOB);
    row := db.Conn.QueryRow(context.Background(), query, person.FirstName, person.MiddleName, person.LastName, person.DOB, person.CognitoId)

    fmt.Println(person.FirstName, person.MiddleName, person.LastName, person.DOB.Time, person.CognitoId)

    err := row.Scan(
        &newPerson.ID,
        &newPerson.FirstName,
        &newPerson.MiddleName,
        &newPerson.LastName,
        &newPerson.DOB,
        &newPerson.CognitoId,
        &newPerson.CreatedAt,
    )

    fmt.Println("ROW")
    fmt.Println(row)

    return &newPerson, err
}

func (db Database) UpdatePerson(person *models.Person) (*models.Person, error) {
    
    fmt.Println(person)

    query := `
        update people
        set
            first_name = $1,
            middle_name = $2,
            last_name = $3,
            d_o_b = $4
        where id = $5
    `

    db.Conn.QueryRow(context.Background(), query, person.FirstName, person.MiddleName, person.LastName, person.DOB, person.ID)

    query = `
        select 
            id,
            first_name,
            middle_name,
            last_name,
            d_o_b,
            created_at,
            cognito_id
        from people where id = $1
    `
    updatedPerson := models.Person{}
    err := db.Conn.QueryRow(context.Background(), query, person.ID).Scan(
        &updatedPerson.ID,
        &updatedPerson.FirstName, 
        &updatedPerson.MiddleName, 
        &updatedPerson.LastName, 
        &updatedPerson.DOB, 
        &updatedPerson.CreatedAt,
        &updatedPerson.CognitoId,
    )


    return &updatedPerson, err
}


// func (db Database) AddItem(item *models.Item) error {
//     var id int
//     var createdAt time.Time
//     query := `INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id, created_at`
//     err := db.Conn.QueryRow(context.Background(), query, item.Name, item.Description).Scan(&id, &createdAt)
//     if err != nil {
// 			fmt.Println(err)
//       return err
//     }
//     item.ID = id
//     item.CreatedAt = createdAt
//     return nil
// }


// func (db Database) GetItemById(itemId int) (models.Item, error) {
//     item := models.Item{}
//     query := `SELECT * FROM items WHERE id = $1;`
//     row := db.Conn.QueryRow(context.Background(), query, itemId)
//     err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)
//     return item, err
// }
// func (db Database) DeleteItem(itemId int) error {
//     query := `DELETE FROM items WHERE id = $1;`
//     _, err := db.Conn.Exec(context.Background(), query, itemId)
//     return err
// }
// func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
//     item := models.Item{}
//     query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`
//     err := db.Conn.QueryRow(context.Background(), query, itemData.Name, itemData.Description, itemId).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)
//     if err != nil {
//         return item, err
//     }
//     return item, nil
// }
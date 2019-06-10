package services

import (
	"fmt"
	"github.com/boblancer/Post-API/models"
	_ "github.com/go-sql-driver/mysql"
	"post-api/database"


)

type UserService struct {
}

/*
	ID			*string `json:"id"`
	FirstName 	*string	`json:"first_name"`
	LastName 	*string	`json:"last"`
	Posts 		[]*post	`json:"posts"`

	ID		*string `json:"id"`
	Header 	*string `json:"header"`
	Body 	*string `json:"body"`

SELECT column_name(s)
FROM table1
LEFT JOIN table2
ON table1.column_name = table2.column_name;

 */
func (UserService) FindAll() ([]models.User, error) {
	DB, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer DB.Close()

	statement := fmt.Sprintf("SELECT * FROM users")
	rows, err := DB.Query(statement)
	if err != nil {
		fmt.Println(err.Error())
	}
	users := []models.User{}
	post := []models.Post{}
	/*
	for rows.Next() {
		var person Person
		var add Address
		if err := rows.Scan(&person.ID, &person.FirstName, &person.LastName, &add.City, &add.State); err != nil {
			fmt.Println(err.Error())
		}
		person.Address = add
		people = append(people, person)
	}
	*/
	//return dataUsers, nil
}

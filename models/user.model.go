package models

type User struct{
	ID			*string `json:"id"`
	FirstName 	*string	`json:"first_name"`
	LastName 	*string	`json:"last"`
	Posts 		[]*Post	`json:"posts"`

}

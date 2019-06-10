package models

type user struct{
	ID			*string
	FirstName 	*string
	Lastname 	*string
	Posts 		[]*post

}

package api

// Profile is the id of the user
type Profile string

// Class of objects that can be granted permission
type Class string

// ID of object that can be granted permission
type ID string

// Object contains both an Id and a type, called Class
type Object struct {
	ID    ID    `json:"id"`
	Class Class `json:"class"`
}

/*
	Resource is the Resource for which Permission is granted
	Examples:
	* Address
	* Product
	* Transaction
	* Role
*/
type Resource string

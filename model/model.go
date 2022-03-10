package model

type Book struct {
	Name        string `form:"name" json:"name"`
	Author      string `from:"author" json:"author"`
	Publication string `from:"publication" json:"publication"`
}

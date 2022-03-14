package model

type Book struct {
	Number      string `from:"number" json:"number"`
	Name        string `form:"name" json:"name"`
	Publication string `from:"publication" json:"publication"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Book
}

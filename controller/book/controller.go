package book

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yogesh-p-thakare3110/bookstore/controller"
	"github.com/yogesh-p-thakare3110/bookstore/database"
	"github.com/yogesh-p-thakare3110/bookstore/model"
)

type bookController struct {
	db database.IDB
}

func NewBook(db database.IDB) controller.IController {
	return &bookController{
		db: db,
	}
}

func (b bookController) AllBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.Response

		rows, err := b.db.GetBook()
		if err != nil {
			log.Println(err.Error())
			return
		}

		response.Status = 200
		response.Message = "Success"
		response.Data = rows

		w.WriteHeader(202)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}

func (b bookController) InsertBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.Response
		var books model.Book
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			return
		}

		err = json.Unmarshal(body, &books)
		if err != nil {
			log.Println(err.Error())
			return
		}
		number, err := b.db.InsertBook(books)
		if err != nil {
			log.Println(err.Error())
			return
		}

		books.Number = fmt.Sprintf("%d", number)
		response.Status = 200
		response.Message = "Insert data Succesfully"
		response.Data = []model.Book{
			books,
		}
		fmt.Print("Insert data to database")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}

func (b bookController) DeleteBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.Response
		var number int64
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			return
		}
		err = json.Unmarshal(body, &number)
		if err != nil {
			log.Println(err.Error())
			return
		}
		err = b.db.DeleteBook(number)
		if err != nil {
			log.Println(err.Error())
			return
		}

		response.Status = 200
		response.Message = "Delete book succesfulluy"
		fmt.Print("Delete data succesfully")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (b bookController) UpdateBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.Response
		var books model.Book
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			return
		}
		err = json.Unmarshal(body, &books)

		number, err := b.db.UpdateBook(books)
		if err != nil {
			log.Println(err.Error())
			return
		}
		books.Name = fmt.Sprintf("%d", number)
		response.Status = 200
		response.Message = "Update book succesfulluy"
		response.Data = []model.Book{
			books,
		}

		fmt.Print("Update book succesfully")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

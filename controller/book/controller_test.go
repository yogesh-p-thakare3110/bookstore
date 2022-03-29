package book_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/yogesh-p-thakare3110/bookstore/controller/book"
	"github.com/yogesh-p-thakare3110/bookstore/model"
)

type mockDB struct {
	f      func() (interface{}, error)
	number int64
}

func (m mockDB) GetBook() ([]model.Book, error) {
	fmt.Println("Mock DB")
	resp, err := m.f()
	return resp.([]model.Book), err
}

func (m mockDB) InsertBook(book model.Book) (string, error) {
	fmt.Println("Mock DB")
	resp, err := m.f()
	return resp.(string), err
}

func (m mockDB) UpdateBook(book model.Book) (string, error) {
	fmt.Println("Mock DB")
	resp, err := m.f()
	return resp.(string), err
}

func (m mockDB) DeleteBook(number int64) error {
	fmt.Println("Mock DB")
	_, err := m.f()
	return err
}

func TestAllbook(t *testing.T) {
	errorCase := true

	db := &mockDB{
		f: func() (interface{}, error) {
			if errorCase {
				return []model.Book{}, errors.New("error case")
			}
			return []model.Book{
				{
					Number:      "10",
					Name:        "A B of Yogi",
					Publication: "Paramhansa",
				},
			}, nil
		},
	}

	eController := book.NewBook(db)

	allBookController := eController.AllBook()

	tests := []struct {
		name       string
		isError    bool
		statusCode int
	}{
		{
			name:       "Failure Case",
			isError:    true,
			statusCode: 200,
		},
		{
			name:       "Success Case",
			isError:    false,
			statusCode: 202,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			r := httptest.NewRequest("GETBOOK", "getallbooks", nil)

			errorCase = tt.isError

			allBookController(w, r)

			t.Log(w.Header())
			t.Log(w.Code)
			t.Log(w.Body.String())

			if w.Code != tt.statusCode {
				t.Error("Want: %v, Got: %v", tt.statusCode, w.Code)
			}
		})
	}
}

func TestInsertBook(t *testing.T) {
	errorCase := true

	db := &mockDB{
		f: func() (interface{}, error) {
			if errorCase {
				return []model.Book{}, errors.New("error case")
			}
			return []model.Book{
				{
					Number:      "10",
					Name:        "A B Of Yogi",
					Publication: "Paramhansa",
				},
			}, nil
		},
	}
	eController := book.NewBook(db)

	insertBookController := eController.InsertBook()

	tests := []struct {
		name       string
		isError    bool
		statusCode int
	}{
		{
			name:       "Failure Case",
			isError:    true,
			statusCode: 200,
		},
		{
			name:       "Success Case",
			isError:    false,
			statusCode: 202,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var newstr = []byte(`{"Number": 1001, "Name" : RDPD, "Publication": Zensar}`)
			r := httptest.NewRequest("POST", "/insertbook", bytes.NewBuffer(newstr))

			errorCase = tt.isError

			insertBookController(w, r)

			t.Log(w.Header())
			t.Log(w.Code)
			t.Log(w.Body.String())

			if w.Code != tt.statusCode {
				t.Errorf("Want: %v, Got: %v", tt.statusCode, w.Code)
			}
		})
	}
}

func TestUpdateBook(t *testing.T) {
	errorCase := true

	db := &mockDB{
		f: func() (interface{}, error) {
			if errorCase {
				return []model.Book{}, errors.New("error case")
			}
			return []model.Book{
				{
					Number:      "10",
					Name:        "A B of Yogi",
					Publication: "Paramhansa",
				},
			}, nil
		},
	}

	eController := book.NewBook(db)

	updateBookController := eController.UpdateBook()

	tests := []struct {
		name       string
		isError    bool
		statusCode int
	}{
		{
			name:       "Failure Case",
			isError:    true,
			statusCode: 200,
		},
		{
			name:       "Success Case",
			isError:    false,
			statusCode: 202,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var newstr = []byte(`{"Number": 1001, "Name" : RDPD, "Publication": Zensar}`)
			r := httptest.NewRequest("PUT", "/updateBooks", bytes.NewBuffer(newstr))

			errorCase = tt.isError

			updateBookController(w, r)

			t.Log(w.Header())
			t.Log(w.Code)
			t.Log(w.Body.String())

			if w.Code != tt.statusCode {
				t.Errorf("Want: %v, Got: %v", tt.statusCode, w.Code)
			}
		})
	}
}

func TestDeleteBook(t *testing.T) {
	errorCase := true

	db := &mockDB{
		f: func() (interface{}, error) {
			if errorCase {
				return []model.Book{}, errors.New("error case")
			}
			return []model.Book{
				{
					Number:      "10",
					Name:        "A B of Yogi",
					Publication: "Paramhansa",
				},
			}, nil
		},
	}

	eController := book.NewBook(db)

	deleteBookController := eController.DeleteBook()

	tests := []struct {
		name       string
		isError    bool
		statusCode int
	}{
		{
			name:       "Failure case",
			isError:    true,
			statusCode: 200,
		},
		{
			name:       "Success Case",
			isError:    false,
			statusCode: 202,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var newstr = []byte(`{"Id":100}`)
			r := httptest.NewRequest("DELETE", "/deleteBooks", bytes.NewBuffer(newstr))

			errorCase = tt.isError

			deleteBookController(w, r)

			t.Log(w.Header())
			t.Log(w.Code)
			t.Log(w.Body.String())

			if w.Code != tt.statusCode {
				t.Errorf("Want: %v, Got: %v", tt.statusCode, w.Code)
			}
		})
	}
}

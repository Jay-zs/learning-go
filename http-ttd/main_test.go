package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetBook(t *testing.T) {
	testcases := []struct {
		desc      string
		req       string
		expRes    []Book
		expStatus int
	}{
		{"get all books", "/book", []Book{
			{1, "Jay", Author{Id: 1}, "Penguin", "11/03/2002"},
		}, http.StatusOK},
		{"get all books with query param", "/book?title=Jay", []Book{
			{1, "Jay", Author{Id: 1}, "Penguin", "11/03/2002"}}, http.StatusOK},
		{"get all books with query param", "/book?includeAuthor=true", []Book{
			{1, "Jay", Author{1, "RD", "Sharma", "2/11/1989", "Sharma"}, "Penguin", "11/03/2002"}}, http.StatusOK},
	}
	for j, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "localhost:8000"+tc.req, nil)

		getBook(w, req)

		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", j, tc.desc)
		}

		res, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		resBooks := []Book{}

		err = json.Unmarshal(res, &resBooks)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		if !reflect.DeepEqual(resBooks, tc.expRes) {
			t.Errorf("%v test failed %v", j, tc.desc)
		}
	}
}

func TestGetBookById(t *testing.T) {
	testcases := []struct {
		desc      string
		req       string
		expRes    Book
		expStatus int
	}{
		{"get book", "1", Book{1, "Jay", Author{1, "RD", "Sharma", "2/11/1989", "Sharma"}, "Penguin", "11/03/2002"}, http.StatusOK},
		{"Id doesn't exist", "1000", Book{}, http.StatusNotFound},
		{"Invalid Id", "abc", Book{}, http.StatusBadRequest},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/book/{id}", nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.req})
		getBookById(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}

		res, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		resBook := Book{}

		err = json.Unmarshal(res, &resBook)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		if resBook != tc.expRes {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc      string
		reqBody   Book
		expRes    Book
		expStatus int
	}{
		{"Valid Details", Book{Title: "Jay", Author: Author{Id: 1}, Publication: "Penguin", PublishedDate: "11/03/2002"}, Book{1, "Jay", Author{Id: 1}, "Penguin", "11/03/2002"}, 200},
		{"Publication should be Scholastic/Penguin/Arihanth", Book{Title: "Jay", Author: Author{Id: 1}, Publication: "Jay", PublishedDate: "11/03/2002"}, Book{}, http.StatusBadRequest},
		{"Published date should be between 1880 and 2022", Book{Title: "", Author: Author{Id: 1}, Publication: "", PublishedDate: "1/1/1870"}, Book{}, http.StatusBadRequest},
		{"Published date should be between 1880 and 2022", Book{Title: "", Author: Author{Id: 1}, Publication: "", PublishedDate: "1/1/2222"}, Book{}, http.StatusBadRequest},
		{"Author should exist", Book{Title: "Jay", Author: Author{Id: 2}, Publication: "Penguin", PublishedDate: "11/03/2002"}, Book{}, http.StatusBadRequest},
		{"Title can't be empty", Book{Title: "", Author: Author{Id: 1}, Publication: "", PublishedDate: ""}, Book{}, http.StatusBadRequest},
		{"Book already exists", Book{Title: "Jay", Author: Author{Id: 1}, Publication: "Penguin", PublishedDate: "11/03/2002"}, Book{}, http.StatusBadRequest},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tc.reqBody)
		req := httptest.NewRequest(http.MethodPost, "localhost:8000/book/", bytes.NewReader(body))
		postBook(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
		res, _ := io.ReadAll(w.Result().Body)
		resBook := Book{}
		json.Unmarshal(res, &resBook)
		if resBook != tc.expRes {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc      string
		reqBody   Author
		expRes    Author
		expStatus int
	}{
		{"Valid details", Author{FirstName: "RD", LastName: "Sharma", Dob: "2/11/1989", PenName: "Sharma"}, Author{1, "RD", "Sharma", "2/11/1989", "Sharma"}, http.StatusOK},
		{"InValid details", Author{FirstName: "", LastName: "Sharma", Dob: "2/11/1989", PenName: "Sharma"}, Author{}, http.StatusBadRequest},
		{"Author already exists", Author{FirstName: "RD", LastName: "Sharma", Dob: "2/11/1989", PenName: "Sharma"}, Author{}, http.StatusBadRequest},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tc.reqBody)
		req := httptest.NewRequest(http.MethodPost, "localhost:8000/author", bytes.NewReader(body))
		postAuthor(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
		res, _ := io.ReadAll(w.Result().Body)
		resAuthor := Author{}
		json.Unmarshal(res, &resAuthor)
		if resAuthor != tc.expRes {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc      string
		reqId     string
		reqBody   Book
		expStatus int
	}{
		{"valid case id exist", "1", Book{1, "title", Author{Id: 1},
			"Arihanth", "18/08/2018"}, http.StatusOK},
		{"invalid case id not exist", "1000", Book{1000, "title1", Author{Id: 9},
			"Arihanth", "18/08/2018"}, http.StatusBadRequest},
		{"Invalid book name.", "1", Book{1, "", Author{Id: 1},
			"Oxford", "21/04/1985"}, http.StatusBadRequest},
		{"Invalid author.", "1", Book{1, "title3", Author{Id: 2},
			"Oxford", "21/04/1985"}, http.StatusBadRequest},
	}

	for i, tc := range testcases {
		body, err := json.Marshal(tc.reqBody)
		if err != nil {
			t.Errorf("can not convert data into []byte")
		}
		req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/book/{id}", bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.reqId})
		putBook(res, req)
		if res.Result().StatusCode != tc.expStatus {
			t.Errorf("test cases fail at %d", i)
		}

	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc      string
		reqId     string
		reqData   Author
		expStatus int
	}{
		{"Valid case update firstname.", "1", Author{1, "Rohan", "gupta", "01/07/2001", "GCC"}, http.StatusOK},

		{"Valid case id not present.", "1000", Author{1000, "Mohan", "chandra", "01/07/2001", "GCC"}, http.StatusBadRequest},
	}

	for i, tc := range testcases {

		body, err := json.Marshal(tc.reqData)
		if err != nil {
			t.Errorf("can not convert data into []byte")
		}
		req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/author/{id}", bytes.NewReader(body))
		res := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.reqId})
		putAuthor(res, req)
		if res.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test cases fail at %v", i, tc.desc)
		}

	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc      string
		reqId     string
		expStatus int
	}{
		{"Valid Details", "5", http.StatusOK},
		{"Book does not exists", "100", http.StatusBadRequest},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/book/{id}", nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.reqId})
		deleteBook(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc      string
		reqId     string
		expStatus int
	}{
		{"Valid Details", "1", http.StatusOK},
		{"Author does not exists", "100", http.StatusBadRequest},
		{"invalid id", "abc", http.StatusBadRequest},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "localhost:8000/author/{id}", nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.reqId})
		deleteAuthor(w, req)
		defer w.Result().Body.Close()

		if w.Result().StatusCode != tc.expStatus {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

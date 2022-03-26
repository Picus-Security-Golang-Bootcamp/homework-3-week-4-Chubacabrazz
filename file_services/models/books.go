package models

type Book struct {
	Book_ID    string `json:"Book_ID"`
	Book_Name  string `json:"Book_Name"`
	Book_Page  string `json:"Book_Page"`
	Book_Stock string `json:"Book_Stock"`
	Book_Price string `json:"Book_Price"`
	Book_Scode string `json:"Book_Scode"`
	Book_ISBN  string `json:"Book_ISBN"`
	Author     string `json:"Author"`
	Author_ID  string `json:"Author_ID"`
}

type BookSlice []Book

package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Item []*BookList
}

type BookList struct {
	Name   string
	Genre  string
	Amount int32
}

func main() {
	book, bookList := BookStore("Harry Potter", 20)

	fmt.Println("test", book, string(bookList))
}

func BookStore(bookName string, amount int32) (result string, bookList []byte) {
	book := Book{
		Item: []*BookList{
			{
				Name:   "Harry Potter",
				Genre:  "Fantasy",
				Amount: 20,
			},
			{
				Name:   "Secrets Of Divine Love",
				Genre:  "Religion",
				Amount: 30,
			},
			{
				Name:   "The Martian",
				Genre:  "Sci-Fi",
				Amount: 14,
			},
			{
				Name:   "Java Programming for Kids",
				Genre:  "Technology",
				Amount: 0,
			},
		},
	}

	for _, v := range book.Item {
		if v.Name == bookName {
			if v.Amount < amount {
				result := fmt.Sprintf("Book named %s is out of stock", v.Name)
				list, _ := json.MarshalIndent(book.Item, "", "\t")

				return result, list
			}

			if v.Amount >= amount {
				v.Amount -= amount
				result := fmt.Sprintf("Book name %s successfully purchased", v.Name)
				list, _ := json.MarshalIndent(book.Item, "", "\t")

				return result, list
			}

		} else {
			list, _ := json.MarshalIndent(book.Item, "", "\t")
			return "Book not found", list
		}
	}

	return
}

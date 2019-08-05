package main

import (
	"fmt"
	"go-ddd/domain/repository"
	"go-ddd/infrastructure/persistence"
)

func main(){
	r := infrastructure.BookRepository{}
	repository.InitBookRepository(&r)
	book, _ := r.Get(1)

	fmt.Println(book)
}

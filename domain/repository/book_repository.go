package repository

import "go-ddd/domain/entity"

type BookRepository interface {
	Get(Id int64) (*entity.Book, error)
	GetAll() ([]*entity.Book, error)
	Save(duck *entity.Book) error
}

var bookRepository BookRepository

func GetBookRepository() BookRepository {
	return bookRepository
}

func InitBookRepository(r BookRepository) {
	bookRepository = r
}

package repository

import "go-ddd/domain/entity"

type BookRepository interface {
	Get(ID int64) (*entity.Book, error)
	GetAll() ([]*entity.Book, error)
	Save(duck *entity.Book) error
}

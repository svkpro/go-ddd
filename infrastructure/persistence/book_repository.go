package infrastructure

import (
	"go-ddd/domain/entity"
)

type BookRepository struct {

}

func (r *BookRepository) Get(Id int64) (*entity.Book, error) {

	book:= entity.Book{1, "DDD in Golang", "Me"}

	return &book, nil
}

func (r *BookRepository) GetAll() ([]*entity.Book, error) {
	panic("implement me")
}

func (r *BookRepository) Save(duck *entity.Book) error {
	panic("implement me")
}

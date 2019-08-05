package service

import (
	"errors"
	"go-ddd/domain/entity"
	"go-ddd/domain/repository"
)

const NilBook = "The book is nil"

type Matcher interface {
	FindMatch(book *entity.Book) (*entity.Book, error)
}

type MatchingService struct{}

func (s *MatchingService) FindMatch(book *entity.Book) (*entity.Book, error) {
	if book == nil {
		return nil, errors.New(NilBook)
	}

	book, err := repository.GetBookRepository().Get(book.Id)

	if err != nil {
		return nil, err
	}

	return book, nil
}

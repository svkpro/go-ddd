package service

import "go-ddd/domain/entity"

type MatchingService interface {
	FindMatch(book *entity.Book) (*entity.Book, error)
}

package repository

import (
	"github.com/go-pp/pp"
	"github.com/ryicoh/nest-option-example/entity"
	"github.com/ryicoh/nest-option-example/repository/option"
)

type (
	UserRepository interface {
		Fetch(ables ...option.UserOptionable) []*entity.User
	}
	userRepository struct{}
)

func NewRepository() UserRepository {
	return new(userRepository)
}

func (r *userRepository) Fetch(ables ...option.UserOptionable) []*entity.User {
	chained := option.DefaultUserOpts
	for _, able := range ables {
		chained = able.UserOption()(chained)
	}

	pp.ColoringEnabled = false
	pp.Println(chained)

	return nil
}

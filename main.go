package main

import (
	"github.com/ryicoh/nest-option-example/repository"
	"github.com/ryicoh/nest-option-example/repository/option"
)

func main() {
	repo := repository.NewRepository()

	repo.Fetch(
		option.ID(1),
		option.Limit(10),
		option.Offset(10),
	)

	repo.Fetch(
		option.ID(1),
	)
}

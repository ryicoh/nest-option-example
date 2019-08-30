package main

import "github.com/go-pp/pp"

func main() {
	repo := new(UserRepo)
	repo.Fetch(
		WithID(1), WithLimit(10), WithOffset(10),
	)
}

type UserRepo struct{}

func (r *UserRepo) Fetch(opts ...UserOptionable) *User {
	chained := defaultUserOpts
	for _, opt := range opts {
		opt.UserOption()(chained)
	}

	pp.ColoringEnabled = false
	pp.Println(chained)

	return new(User)
}

type UserOptionable interface {
	UserOption() UserOption
}

type UserOption func(*UserOptions)

func (o UserOption) UserOption() UserOption {
	return o
}

var defaultUserOpts = &UserOptions{
	Options: defaultOpts,
}

type UserOptions struct {
	*Options
	ID int64
}

type User struct {
	ID   int64
	Name string
}

type Option func(*Options)

func (o Option) UserOption() UserOption {
	return func(opts *UserOptions) {
		if opts.Options == nil {
			opts.Options = new(Options)
		}
		o(opts.Options)
	}
}

var defaultOpts = &Options{
	Limit: 20,
}

type Options struct {
	Limit, Offset int64
}

func WithID(id int64) UserOption {
	return func(opts *UserOptions) {
		opts.ID = id
	}
}

func WithLimit(limit int64) Option {
	return func(opts *Options) {
		opts.Limit = limit
	}
}

func WithOffset(offset int64) Option {
	return func(opts *Options) {
		opts.Offset = offset
	}
}

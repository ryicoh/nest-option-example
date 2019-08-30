package option

type (
	UserOptions struct {
		Options
		ID int64
	}

	UserOption func(UserOptions) UserOptions

	UserOptionable interface {
		UserOption() UserOption
	}
)

var DefaultUserOpts = UserOptions{
	Options: defaultOpts,
}

func (o UserOption) UserOption() UserOption {
	return o
}

func ID(id int64) UserOption {
	return func(opts UserOptions) UserOptions {
		opts.ID = id
		return opts
	}
}

package option

type Options struct {
	Limit, Offset int64
}

type Option func(Options) Options

func (o Option) UserOption() UserOption {
	return func(opts UserOptions) UserOptions {
		opts.Options = o(opts.Options)
		return opts
	}
}

var defaultOpts = Options{
	Limit: 20,
}

func Limit(limit int64) Option {
	return func(opts Options) Options {
		opts.Limit = limit
		return opts
	}
}

func Offset(offset int64) Option {
	return func(opts Options) Options {
		opts.Offset = offset
		return opts
	}
}

package bk_converter

import "context"

type Loader interface {
	Load(ctx context.Context, args Args, ch chan<- Record) error
}

type Store interface {
	Store(ctx context.Context, args Args, ch <-chan Record) error
}

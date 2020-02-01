package bk_converter

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/multierr"
)

var (
	loaders map[string]Loader
	stores  map[string]Store
)

func Register(name string, i interface{}) {
	switch it := i.(type) {
	case Loader:
		loaders[name] = it
	case Store:
		stores[name] = it
	default:
		panic(errors.New("Loader or Store required"))
	}
}

type Converter struct {
	config Config
}

func New(config Config) Converter {
	validate(config)
	return Converter{config: config}
}

func validate(config Config) {
	for _, conversion := range config.Conversions {
		from, to := conversion.From, conversion.To
		if _, ok := loaders[from.Name]; !ok {
			panic(fmt.Errorf("Loader %s does not existed", from))
		}
		if _, ok := stores[to.Name]; !ok {
			panic(fmt.Errorf("Store %s does not existed", to))
		}
	}
}

func (c *Converter) Run(ctx context.Context) (err error) {

	concurrent := c.config.Concurrent
	conversion := c.config.Conversions

	for _, conversion := range conversion {
		from, to := conversion.From, conversion.To
		loader := loaders[from.Name]
		store := stores[to.Name]
		if concurrent {
			go func() {
				if err2 := run(ctx, from.Args, loader, store); err2 != nil {
					err = multierr.Append(err, err2)
				}
			}()
		} else {
			if err := run(ctx, to.Args, loader, store); err != nil {
				return err
			}
		}
	}

	return
}

func run(parent context.Context, args Args, loader Loader, store Store) (err error) {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()
	ch := make(chan Record)
	go func() {
		defer close(ch)
		if err2 := loader.Load(ctx, args, ch); err2 != nil {
			err = multierr.Append(err, err2)
			return
		}
	}()
	if err2 := store.Store(ctx, args, ch); err2 != nil {
		err = multierr.Append(err, err2)
	}
	return err
}

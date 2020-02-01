package dummy

import (
	"context"
	"encoding/json"
	"github.com/zhenzou/bk_converter"
)

func init() {
	bk_converter.Register("dummy", &DummyStore{})
}

type DummyStore struct {
}

func (d *DummyStore) Store(ctx context.Context, args bk_converter.Args, ch <-chan bk_converter.Record) error {
	for record := range ch {
		bytes, err := json.Marshal(record)
		if err != nil {
			println("store error:", err.Error())
			continue
		}
		println("store succeed:", string(bytes))
	}
	return nil
}

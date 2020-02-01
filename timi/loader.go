package timi

import (
	"context"
	"encoding/csv"
	"io"
	"os"

	"github.com/jszwec/csvutil"

	"github.com/zhenzou/bk_converter"
)

func init() {
	bk_converter.Register("timi", &TimiLoader{})
}

type TimiLoader struct {
}

func (t *TimiLoader) Load(ctx context.Context, args bk_converter.Args, ch chan<- bk_converter.Record) error {

	file, err := os.Open(args.In)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder, err := csvutil.NewDecoder(csv.NewReader(file))
	if err != nil {
		return err
	}

	for {
		record := Record{}
		err := decoder.Decode(&record)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		ch <- ConvertToRecord(record)
	}

	return nil
}

func ConvertToRecord(record Record) bk_converter.Record {

	result := bk_converter.Record{}

	return result
}

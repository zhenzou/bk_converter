package timi

import (
	"context"
	"encoding/csv"
	"fmt"
	"gopkg.in/errgo.v2/errors"
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
		return errors.Wrap(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	decoder, err := csvutil.NewDecoder(reader)
	if err != nil {
		return err
	}

	for {
		record := Record{}
		err := decoder.Decode(&record)
		//fmt.Println("load:", record)
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

	class, ok := NameMapping[record.Name]
	if !ok {
		panic(fmt.Errorf("unknown name %s", record.Name))
	}
	result.FirstClass = class.First
	result.SecondClass = class.Second
	if record.Type == "支出" {
		result.Type = bk_converter.Out
	} else if record.Type == "收入" {
		result.Type = bk_converter.In
	}
	result.Remark = record.Remark
	result.Amount = record.Amount
	result.Time = record.Time.Time

	return result
}

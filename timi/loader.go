package timi

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/jszwec/csvutil"
	"gopkg.in/errgo.v2/errors"
	"gopkg.in/yaml.v2"

	"github.com/zhenzou/bk_converter"
)

func init() {
	bk_converter.Register("timi", &TimiLoader{})
}

type TimiLoader struct {
}

func (t *TimiLoader) Load(ctx context.Context, args bk_converter.Args, ch chan<- bk_converter.Record) error {
	mapping, err := loadMapping(args.Mapping)
	if err != nil {
		return errors.Wrap(err)
	}

	file, err := bk_converter.Open(args.In)
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
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		ch <- ConvertToRecord(mapping, record)
	}

	return nil
}

func loadMapping(name string) (mapping map[string]Class, err error) {
	data, err := bk_converter.ReadAll(name)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &mapping)
	return
}

func ConvertToRecord(mapping map[string]Class, record Record) bk_converter.Record {

	result := bk_converter.Record{}

	class, ok := mapping[record.Name]
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
	result.FromAccount = "现金"
	result.Remark = record.Remark
	result.Amount = record.Amount
	result.Time = record.Time.Time

	return result
}

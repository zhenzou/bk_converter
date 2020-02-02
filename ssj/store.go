package ssj

import (
	"context"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/zhenzou/bk_converter"
	"github.com/zuijinbuzai/excelclaim/excel"
)

var (
	header = []string{"交易种类", "类别", "子类", "账户1", "账户2", "花费", "人员", "商家", "事件", "日期", "详细"}
)

func init() {
	bk_converter.Register("ssj", &SSJStore{})
}

type SSJStore struct {
}

func (d *SSJStore) Store(ctx context.Context, args bk_converter.Args, ch <-chan bk_converter.Record) error {

	file := excelize.NewFile()

	file.NewSheet("支出")
	file.NewSheet("收入")

	sheetOut := excel.NewSheet(file, "支出", 11, 28)
	sheetIn := excel.NewSheet(file, "收入", 11, 28)

	sheetIn.WriteRow(header...)
	sheetOut.WriteRow(header...)

	for record := range ch {
		row := ConvertToRow(record)
		if record.Type == bk_converter.In {
			sheetIn.WriteRow(row...)
		} else if record.Type == bk_converter.Out {
			sheetOut.WriteRow(row...)
		}
	}
	return file.SaveAs(args.Out)
}

func ConvertToRow(r bk_converter.Record) []string {
	slice := []string{}
	if r.Type == bk_converter.In {
		slice = append(slice, "收入")
	} else {
		slice = append(slice, "支出")
	}

	slice = append(slice, r.FirstClass)
	slice = append(slice, r.SecondClass)
	slice = append(slice, r.FromAccount)
	slice = append(slice, " ")
	slice = append(slice, strconv.FormatFloat(r.Amount, 'f', 2, 64))
	slice = append(slice, " ")
	slice = append(slice, " ")
	slice = append(slice, " ")
	slice = append(slice, r.Time.Format("2006-01-02 15:04:05"))
	slice = append(slice, r.Remark)
	return slice
}

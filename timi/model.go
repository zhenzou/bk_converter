package timi

import "time"

var (
	NameMapping = map[string]Class{}
)

type Class struct {
	First  string
	Second string
}

type Record struct {
	ID     string  `json:"id" csv:"账目编号"`
	Type   string  `json:"type" csv:"类型"`
	Name   string  `json:"name" csv:"账目名称"`
	Amount float64 `json:"amount" csv:"金额"`
	Time   Time    `json:"time" csv:"时间"`
	Remark string  `json:"remark" csv:"备注"`
	Image  string  `json:"image" csv:"相关图片"`
}

type Time struct {
	time.Time
}

const format = "2006-01-02 15:04:05"

func (t Time) MarshalCSV() ([]byte, error) {
	var b [len(format)]byte
	return t.AppendFormat(b[:0], format), nil
}

func (t *Time) UnmarshalCSV(data []byte) error {
	tt, err := time.Parse(format, string(data))
	if err != nil {
		return err
	}
	*t = Time{Time: tt}
	return nil
}

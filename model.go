package bk_converter

import "time"

type Type int

const (
	In  Type = 1
	Out      = 2
)

// 以随手记的为中间格式
type Record struct {
	Type        Type
	FirstClass  string
	SecondClass string
	FromAccount string
	ToAccount   string
	Amount      float64
	Recorder    string
	Remark      string
	Time        time.Time
}

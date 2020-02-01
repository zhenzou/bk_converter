package timi

import "time"

var (
	NameMapping = map[string]Class{
		"住房": {
			First:  "居家物业",
			Second: "房租",
		},
		"数码": {
			First:  "购物剁手",
			Second: "数码电器",
		},
		"服饰": {
			First:  "购物剁手",
			Second: "衣服裤子",
		},
		"日用品": {
			First:  "居家物业",
			Second: "日常用品",
		},
		"家居": {
			First:  "居家物业",
			Second: "家具家居",
		},
		"医疗": {
			First:  "医疗保健",
			Second: "治疗费",
		},
		"人情": {
			First:  "人情往来",
			Second: "送礼请客",
		},
		"娱乐": {
			First:  "休闲娱乐",
			Second: "休闲玩乐",
		},
		"用餐": {
			First:  "食品酒水",
			Second: "早午晚餐",
		},
		"零食": {
			First:  "食品酒水",
			Second: "水果零食",
		},
		"交通": {
			First:  "行车交通",
			Second: "公共交通",
		},
		"学习": {
			First:  "学习进修",
			Second: "数码装备",
		},
		"通讯": {
			First:  "交流通讯",
			Second: "手机费",
		},
		"旅游": {
			First:  "休闲娱乐",
			Second: "旅游度假",
		},
		"丽人": {
			First:  "医疗保健",
			Second: "美容费",
		},
		"一般": {
			First:  "其他杂项",
			Second: "其他支出",
		},
		"投资": {
			First:  "金融保险",
			Second: "投资亏损",
		},
		"亲属卡": {
			First:  "人情往来",
			Second: "孝敬家长",
		},
		"还贷": {
			First:  "金融保险",
			Second: "按揭还款",
		},
		"借款": {
			First:  "人情往来",
			Second: "借款",
		},
		"工资": {
			First:  "职业收入",
			Second: "工资收入",
		},
		"奖金": {
			First:  "职业收入",
			Second: "奖金收入",
		},
		"兼职": {
			First:  "职业收入",
			Second: "兼职收入",
		},
		"公积金": {
			First:  "其他收入",
			Second: "公积金",
		},
		"还款": {
			First:  "其他收入",
			Second: "还款",
		},
		"退款": {
			First:  "其他收入",
			Second: "退款",
		},
		"红包": {
			First:  "其他收入",
			Second: "红包",
		},
		"理财": {
			First:  "职业收入",
			Second: "投资收入",
		},
		"贷款": {
			First:  "其他收入",
			Second: "贷款",
		},
	}
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

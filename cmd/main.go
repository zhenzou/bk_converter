package main

import (
	"context"
	"os"

	"gopkg.in/errgo.v2/errors"

	"github.com/zhenzou/bk_converter"
	_ "github.com/zhenzou/bk_converter/dummy"
	_ "github.com/zhenzou/bk_converter/ssj"
	_ "github.com/zhenzou/bk_converter/timi"
)

func loadConfig() bk_converter.Config {
	return bk_converter.Config{
		Concurrent: false,
		Conversions: []bk_converter.Conversion{{
			From: bk_converter.Args{
				Name:    "timi",
				In:      "/Users/zouzhen/Downloads/Timi571265945_20200131_utf8/日常账本_36604458_utf8.csv",
				Mapping: "",
				Others:  nil,
			},
			To: bk_converter.Args{
				Name:    "ssj",
				In:      "",
				Out:     "/Users/zouzhen/Downloads/ssj.xls",
				Mapping: "",
				Others:  nil,
			},
		}},
	}
}

func main() {

	config := loadConfig()

	converter := bk_converter.New(config)
	err := converter.Run(context.Background())
	if err != nil {

		println("err:", errors.Details(err))
		os.Exit(-1)
	}
}

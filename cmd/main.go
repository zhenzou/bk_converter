package main

import (
	"context"
	"flag"
	"os"

	"gopkg.in/errgo.v2/errors"
	"gopkg.in/yaml.v2"

	"github.com/zhenzou/bk_converter"
	_ "github.com/zhenzou/bk_converter/dummy"
	_ "github.com/zhenzou/bk_converter/ssj"
	_ "github.com/zhenzou/bk_converter/timi"
)

var (
	conf string
)

func init() {
	flag.StringVar(&conf, "conf", "config.yml", "config path")
	flag.Parse()
}

func loadConfig() bk_converter.Config {
	data, err := bk_converter.ReadAll(conf)
	if err != nil {
		println("load config error:", err.Error())
		os.Exit(-1)
	}
	config := bk_converter.Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		println("load config error:", err.Error())
		os.Exit(-1)
	}
	return config
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

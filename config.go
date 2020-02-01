package bk_converter

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Concurrent  bool         `yaml:"concurrent"`
	Conversions []Conversion `yaml:"conversions"`
}

type Conversion struct {
	From Args `yaml:"from"`
	To   Args `yaml:"to"`
}

type Args struct {
	Name    string            `yaml:"name"`
	In      string            `yaml:"in"`
	Out     string            `yaml:"out"`
	Mapping string            `yaml:"mapping"`
	Others  map[string]string `yaml:"-"`
}

func (a *Args) MarshalYAML() ([]byte, error) {
	m := map[string]string{}
	for k, v := range a.Others {
		m[k] = v
	}
	m["name"] = a.Name
	m["in"] = a.In
	m["out"] = a.Out
	m["mapping"] = a.Mapping
	return yaml.Marshal(m)
}

func (a *Args) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := map[string]string{}
	err := unmarshal(&m)
	if err != nil {
		return err
	}
	a.Name = m["name"]
	a.In = m["in"]
	a.Out = m["out"]
	a.Mapping = m["mapping"]
	a.Others = m
	delete(m, "name")
	delete(m, "in")
	delete(m, "out")
	delete(m, "mapping")
	return nil
}

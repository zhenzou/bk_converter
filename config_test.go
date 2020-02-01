package bk_converter

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/stretchr/testify/assert"
)

func TestArgs_MarshalYAML(t *testing.T) {
	type fields struct {
		In      string
		Out     string
		Mapping string
		Others  map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "not nil others",
			fields: fields{
				In:      "in",
				Out:     "out",
				Mapping: "path",
				Others: map[string]string{
					"other":  "other",
					"other1": "other1",
				},
			},
			want: []byte(
				`in: in
mapping: path
other: other
other1: other1
out: out
`),
			wantErr: false,
		},
		{
			name: "nil others",
			fields: fields{
				In:      "in",
				Out:     "out",
				Mapping: "path",
				Others:  nil,
			},
			want: []byte(
				`in: in
mapping: path
out: out
`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Args{
				In:      tt.fields.In,
				Out:     tt.fields.Out,
				Mapping: tt.fields.Mapping,
				Others:  tt.fields.Others,
			}
			got, err := a.MarshalYAML()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalYAML() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgs_UnmarshalYAML(t *testing.T) {
	y1 := `in: in
mapping: path
other: other
other1: other1
out: out
`

	a1 := Args{}
	err := yaml.Unmarshal([]byte(y1), &a1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "in", a1.In)
	assert.Equal(t, "other1", a1.Others["other1"])
}

package structs

import (
	"testing"
)

type s struct {
	F int64       `json:"f"`
	G float32     `json:"g"`
	H interface{} `json:"h"`
}

type NestStruct struct {
	A string  `json:"a"`
	B int     `json:"b"`
	C float64 `json:"c"`
	D bool    `json:"d"`
	E s       `json:"e"`
}

func TestStructs_Map(t *testing.T) {
	type fields struct{}
	f1 := fields{}
	type args struct {
		itf interface{}
	}
	c1 := args{
		itf: &NestStruct{
			A: "A",
			B: 1,
			C: 1.1,
			D: false,
			E: s{
				F: 1,
				G: 1.2,
				H: nil,
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "c1", fields: f1, args: c1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Structs{}
			got := s.Map(tt.args.itf)
			t.Logf("%+v", got)
		})
	}
}

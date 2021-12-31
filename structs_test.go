package structs

import (
	"fmt"
	"testing"
)

type s struct {
	F int64   `json:"f"`
	G float32 `json:"g"`
	H interface{}
}

type NestStruct struct {
	A  string  `json:"a"`
	B  int     `json:"b"`
	C  float64 `json:"c"`
	D  bool    `json:"d"`
	E  s       `json:"e"`
	EF int64   `json:"ef"`
}

type NestStruct2 struct {
	A int     `json:"a"`
	B int     `json:"b"`
	C float64 `json:"c"`
	D bool    `json:"d"`
	E s       `json:"e"`
}

func TestStructs_Map(t *testing.T) {
	//f1 := New([]string{"A", "B"})
	f1 := New(nil)
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
			},
			EF: 10,
		},
	}
	tests := []struct {
		name    string
		fields  *Structs
		args    args
		wantErr bool
	}{
		{name: "c1", fields: f1, args: c1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.fields.IgnoreFields)
			got := s.Map(tt.args.itf)
			t.Logf("%+v", got)
		})
	}
}

func TestStructs_Facade(t *testing.T) {
	f1 := New([]string{"A"})
	type args struct {
		src interface{}
		dst interface{}
	}
	var d NestStruct2
	c1 := args{
		src: &NestStruct{
			A: "A",
			B: 1,
			C: 1.1,
			D: false,
			E: s{
				F: 1,
				G: 1.2,
			},
		},
		dst: &d,
	}
	tests := []struct {
		name    string
		fields  *Structs
		args    args
		wantErr bool
	}{
		{name: "c1", fields: f1, args: c1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Structs{
				IgnoreFields: tt.fields.IgnoreFields,
				AliasFields:  tt.fields.AliasFields,
			}
			if err := s.Facade(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("Facade() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Printf("%+v", tt.args.dst)
		})
	}
}

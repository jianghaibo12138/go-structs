package structs

import "testing"

type s struct {
	F int64
	G float32
	H interface{}
}

type NestStruct struct {
	A string
	B int
	C float64
	D bool
	E s
}

func TestMap(t *testing.T) {
	type args struct {
		s interface{}
	}
	c1 := args{
		s: NestStruct{
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
		name string
		args args
	}{
		{name: "c1", args: c1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Map(tt.args.s)
		})
	}
}

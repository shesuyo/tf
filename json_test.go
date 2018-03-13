package wtf

import (
	"testing"

	"github.com/pquerna/ffjson/ffjson"
)

func TestJSONStringify(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JSONStringify(tt.args.obj); got != tt.want {
				t.Errorf("JSONStringify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ffJSONStringify(obj interface{}) string {
	bs, _ := ffjson.Marshal(obj)
	return BytesString(bs)
}

type benchJSONStruct struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Benchmark_JSONStringifyMap(b *testing.B) {
	a := map[string]interface{}{"code": 0, "data": "1", "msg": "ok"}
	for i := 0; i < b.N; i++ {
		JSONStringify(a)
	}
}

func Benchmark_ffJSONStringifyMap(b *testing.B) {
	a := map[string]interface{}{"code": 0, "data": "1", "msg": "ok"}
	for i := 0; i < b.N; i++ {
		ffJSONStringify(a)
	}
}

func Benchmark_JSONStringifyStruct(b *testing.B) {
	a := benchJSONStruct{Code: 0, Data: 1, Msg: "ok"}
	for i := 0; i < b.N; i++ {
		JSONStringify(a)
	}
}

func Benchmark_ffJSONStringifyStruct(b *testing.B) {
	a := benchJSONStruct{Code: 0, Data: 1, Msg: "ok"}
	for i := 0; i < b.N; i++ {
		ffJSONStringify(a)
	}
}

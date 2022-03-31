package tf

import (
	"strconv"
	"testing"
)

var m map[string]int

func init() {
	m = make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		m[strconv.Itoa(i)] = i
	}
}

func TestInt(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.v); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Int1(b *testing.B) {
	var a interface{} = "1"
	for i := 0; i < b.N; i++ {
		v, ok := a.(string)
		_, _ = v, ok
	}
}

func Benchmark_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(i)
	}
}

// Benchmark_MapGet-4   	100000000	        17.9 ns/op	       0 B/op	       0 allocs/op
func Benchmark_MapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := m["72"]
		_ = a
	}
}

//Benchmark_MapGetWithoutAssign-4   	100000000	        14.9 ns/op	       0 B/op	       0 allocs/op
func Benchmark_MapGetWithoutAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m["72"]
	}
}

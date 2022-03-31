package tf

import (
	"reflect"
	"testing"
)

func TestStringBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{"", args{"asd"}, []byte("asd")},
		{"", args{"aa@@#%$!%^*&&*!@"}, []byte("aa@@#%$!%^*&&*!@")},
		{"", args{"102548 5485648"}, []byte("102548 5485648")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

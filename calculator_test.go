package calculator

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TestCase1", args{10, 2}, 12},
		{"TestCase2", args{4, 3}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleAdd() {
	var sum int
	a := 3
	b := 2

	sum = Add(a, b)
	fmt.Println(a, " + ", b, " = ", sum)

	// Output:
	// 3 + 2 = 5
}

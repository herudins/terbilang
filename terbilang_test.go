package terbilang

import (
	"fmt"
	"testing"
)

func TestTerbilang(t *testing.T) {
	angka := 12142343546456
	terbilangStr := Terbilang(angka)
	fmt.Println(terbilangStr)
}

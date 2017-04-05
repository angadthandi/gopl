package ch1_test

import (
	"github.com/angadthandi/gopl/ch1"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	ch1.Echo1()
}

func BenchmarkEcho2(b *testing.B) {
	ch1.Echo2()
}

func BenchmarkEcho3(b *testing.B) {
	ch1.Echo3()
}

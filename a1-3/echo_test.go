package echo

import(
	"testing"
)

//!-test
//!+bench
func BenchmarkEcho1(b *testing.B){
	for i :=0;i<b.N;i++{
		Echo1([]string{"I","Love","You"})
	}
}
/*
func BenchmarkEcho2(b *testing.B){
	for i :=0;i<b.N;i++{
		Echo2([]string{0:"I",1:"Love",2:"You"})
	}
}*/

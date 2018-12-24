// goruntime_demo_test
package main

import "testing"

func TestDemo(t *testing.T) {
	demo()
}

func BenchmarkAdd(t *testing.B) {

	//重置时间点
	t.ReportAllocs()
	demo()
}

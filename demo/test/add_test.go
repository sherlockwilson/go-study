package add

import "testing"

func TestAdd(t *testing.T) {
	//定义测试数据
	tests := []struct{ a, b, c int }{
		{3, 4, 7},
		{5, 12, 17},
		{8, 15, 23},
		{12, 35, 47},
		{30000, 40000, 70000},
	}
	//测试逻辑
	for _, tt := range tests {
		if actual := add(tt.a, tt.b); actual != tt.c {
			t.Errorf("Add(%d,%d) got %d;expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkAdd(t *testing.B) {

	//重置时间点
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		add(1, 2)
	}
}

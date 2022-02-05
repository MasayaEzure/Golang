package index

// テスト用パッケージをインポート
import "testing"

var Debug bool = false

// テスト処理
func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("スキップ")
	}
	v := Average([]int{1,2,3,4,5})
	if v != 3 {
		t.Errorf("%v != %v", v, 3)
	}
}
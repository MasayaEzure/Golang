package main

// テスト用のパッケージをインポート
import "testing"

var Debug bool = false

// テスト処理
func TestIsTrue(t *testing.T) {
	i := 10
	if Debug {
		t.Skip("スキップ")
	}
	v := IsTrue(i)

	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}
package main

import (
	"fmt"
	"regexp"
)

// regex

func main() {
	// Goの正規表現の基本
	m, _ := regexp.MatchString("a", "xami")
	fmt.Println(m)

	// Compile
	r, _ := regexp.Compile("f")
	m = r.MatchString("bjlfyho")
	fmt.Println(m)

	// MustCompile
	aa := regexp.MustCompile("x")
	m = aa.MatchString("gafjdlx")
	fmt.Println(m)
	
	/*
	フラグ一覧
	i 大文字と小文字を区別しない
	m マルチラインモード（^と&が文頭、文末に加えて行頭、行末にマッチ）
	s .が\nにマッチ
	U 最小マッチへの変換 (x*はx*?、x+はx+?)
	*/
	bb := regexp.MustCompile(`(?is)abc`)
	a := bb.MatchString("ABC")
	fmt.Println(a)

	/*
	幅をもたない正規表現のパターン
	パターン一覧
	^ 文頭
	$ 文末
	\A 文頭
	\z 文末
	\b ASCIIによるワード協会
	\B 非ASCIIによるワード協会
	*/
	cc := regexp.MustCompile(`^ABC$`)
	n := cc.MatchString("ABC")
	fmt.Println(n)
	
	m = cc.MatchString("  ABC ")
	fmt.Println(m)

	/*
	繰り返しを表す正規表現
	パターン一覧
	x* 0回以上
	x+ 1回以上
	x? 0 <= x <= 1
	x{n,m} n <= x <= m (回)
	x{n, } n <= x (回)
	x{n} n = x (回)
	x*? 0回以上
	x+? 1回以上
	x?? 0 <= x <= 1 (0を優先)
	x{n,m}? n <= x <= m (回)
	x{n, }? N <= x (回)
	x{n}? n = x (回)
	*/
	
	dd := regexp.MustCompile("a+b*")
	fmt.Println(dd.MatchString("ab"))
	fmt.Println(dd.MatchString("a"))
	fmt.Println(dd.MatchString("aaaaaabbbbb"))
	fmt.Println(dd.MatchString("b"))

	// 正規表現の文字クラス
	ee := regexp.MustCompile(`[XYZ]`)
	fmt.Println(ee.MatchString("Y"))

	ff := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(ff.MatchString("abc"))
	fmt.Println(ff.MatchString("abcdefg"))

	gg := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(gg.MatchString("abc"))
	fmt.Println(gg.MatchString("あ"))

	// 正規表現のグループ
	hh := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	fmt.Println(hh.MatchString("abcxyz"))
	fmt.Println(hh.MatchString("ABCXYZ"))
	fmt.Println(hh.MatchString("ABCxyz"))
	fmt.Println(hh.MatchString("abcXYZ"))
	fmt.Println(hh.MatchString("ABCabc"))

	// 正規表現にマッチした文字列の取得(FindString)
	fa := hh.FindString("AAAAABCXYZZZZZ")
	fmt.Println(fa)
	fb := hh.FindAllString("ABCXYZABCXYZ", -1)
	fmt.Println(fb)

	// 正規表のグループによるサブマッチ
	ii := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
		0123-456-789
		111-222-333
		556-787-889
		`

	ms := ii.FindAllStringSubmatch(s, -1)
	fmt.Println(ms)

	for _, v := range ms {
		fmt.Println(v)
	}

	fmt.Println(ms[0][0])
	fmt.Println(ms[0][1])
	fmt.Println(ms[0][2])
	fmt.Println(ms[0][3])

	// 正規表現による文字列の置換
	jj := regexp.MustCompile(`\s+`)
	fmt.Println(jj.ReplaceAllString("A B C", ":"))

	kk := regexp.MustCompile(`、|。`)
	fmt.Println(kk.ReplaceAllString("あいう、えおかきくけこ、さしすせそ。", ""))

	// 正規表現による文字列の分割
	ll := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	fmt.Println(ll.Split("DEFGHABCXYZIJKLMABCXYZNOPQRSTUVWXYZ", -1))

	mm := regexp.MustCompile(`\s+`)
	fmt.Println(mm.Split("a                b       c", -1))

}
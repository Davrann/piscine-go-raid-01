package student

import "github.com/01-edu/z01"

func Raid1d(x, y int) {
	var tl rune = 'A'
	var tr rune = 'C'
	var bl rune = 'A'
	var br rune = 'C'
	var w rune = 'B'
	var h rune = 'B'
	var s rune = ' '
	var eol rune = '\n'

	z01.PrintRune(tl)
	for i := 0; i < x-2; i++ {
		z01.PrintRune(w)
	}
	if x > 1 {
		z01.PrintRune(tr)
	}
	z01.PrintRune(eol)
	for i := 0; i < y-2; i++ {
		z01.PrintRune(h)
		for j := 0; j < x-2; j++ {
			z01.PrintRune(s)
		}
		if x > 1 {
			z01.PrintRune(h)
		}
		z01.PrintRune(eol)
	}
	if y > 1 {
		z01.PrintRune(bl)
		for i := 0; i < x-2; i++ {
			z01.PrintRune(w)
		}
		if x > 1 {
			z01.PrintRune(br)
		}
		z01.PrintRune(eol)
	}
}

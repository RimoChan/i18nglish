package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var r4p = regexp.MustCompile(`[\w\x{3005}-\x{3007}\x{3021}-\x{3029}\x{3031}-\x{3035}\x{3038}-\x{303C}\x{3220}-\x{3229}\x{3248}-\x{324F}\x{3251}-\x{325F}\x{3280}-\x{3289}\x{32B1}-\x{32BF}\x{3400}-\x{4DB5}\x{4E00}-\x{9FEA}\x{F900}-\x{FA6D}\x{FA70}-\x{FAD9}\x{20000}-\x{2A6D6}\x{2A700}-\x{2B734}\x{2B740}-\x{2B81D}\x{2B820}-\x{2CEA1}\x{2CEB0}-\x{2EBE0}\x{2F800}-\x{2FA1D}]+`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(r4p.ReplaceAllStringFunc(scanner.Text(), i7h))
	}
}

func i7h(w2d string) string {
	r2n := []rune(strings.TrimSpace(w2d))
	if len(r2n) <= 3 {
		return w2d
	}
	return i7h(fmt.Sprintf("%c%d%c", r2n[0], len(r2n)-2, r2n[len(r2n)-1]))
}

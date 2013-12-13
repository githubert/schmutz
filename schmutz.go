package main

import (
	"bufio"
	"flag"
	"io"
	"math/rand"
	"os"
)

func main() {
	small := []rune{'\u0323', '\u0307', '\u0312'}
	large := []rune{'\u0314', '\u031C', '\u0358', '\u0353', '\u0335'}
	strikeOut := []rune{'\u0337', '\u0338', '\u0336', '\u0335',
		'\u20d2', '\u20d3', '\u20e5', '\u20e6', '\u20eb'}

	var useSmall = flag.Bool("feinstaub", false, "spray dust on your text")
	var useLarge = flag.Bool("grobe-mettwurst", false, "a cookie got mangled in your typewriter")
	var useStrikeout = flag.Bool("nein-nein-nein", false, "this is unacceptable")
	var prob = flag.Float64("p", 0.1, "probability between 0.0 and 1.0")

	flag.Parse()

	runes := []rune{}

	if *useSmall {
		runes = append(runes, small...)
	}

	if *useLarge {
		runes = append(runes, large...)
	}

	if *useStrikeout {
		runes = append(runes, strikeOut...)
	}

	if len(runes) == 0 {
		runes = append(runes, small...)
	}

	in := bufio.NewReader(os.Stdin)

	for {
		str, err := in.ReadString('\n')

		if err != nil {
			break
		}

		io.WriteString(os.Stdout, mangle(str, runes, *prob))
	}
}

func mangle(text string, runes []rune, prob float64) (result string) {
	for _, rune := range text {
		result += string(rune)

		if rand.Intn(100) <= int(100.0*prob) {
			result += string(runes[rand.Intn(len(runes))])
		}
	}

	return
}

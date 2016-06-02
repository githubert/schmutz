/* schmutz.go: You̇r screen is quiṭe dirty, please cleȧn it.
 *
 * Copyright (C) 2016 Clemens Fries <github-schmutz@xenoworld.de>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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

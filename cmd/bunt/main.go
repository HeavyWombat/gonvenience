// Copyright © 2018 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"fmt"

	"github.com/homeport/gonvenience/pkg/v1/bunt"
	"github.com/homeport/gonvenience/pkg/v1/neat"
	"github.com/homeport/gonvenience/pkg/v1/term"
)

func main() {
	list := []string{
		bunt.Sprintf("Pink{Pink}"),
		bunt.Sprintf("LightPink{LightPink}"),
		bunt.Sprintf("HotPink{HotPink}"),
		bunt.Sprintf("DeepPink{DeepPink}"),
		bunt.Sprintf("PaleVioletRed{PaleVioletRed}"),
		bunt.Sprintf("MediumVioletRed{MediumVioletRed}"),
		bunt.Sprintf("LightSalmon{LightSalmon}"),
		bunt.Sprintf("Salmon{Salmon}"),
		bunt.Sprintf("DarkSalmon{DarkSalmon}"),
		bunt.Sprintf("LightCoral{LightCoral}"),
		bunt.Sprintf("IndianRed{IndianRed}"),
		bunt.Sprintf("Crimson{Crimson}"),
		bunt.Sprintf("FireBrick{FireBrick}"),
		bunt.Sprintf("DarkRed{DarkRed}"),
		bunt.Sprintf("Red{Red}"),
		bunt.Sprintf("OrangeRed{OrangeRed}"),
		bunt.Sprintf("Tomato{Tomato}"),
		bunt.Sprintf("Coral{Coral}"),
		bunt.Sprintf("DarkOrange{DarkOrange}"),
		bunt.Sprintf("Orange{Orange}"),
		bunt.Sprintf("Yellow{Yellow}"),
		bunt.Sprintf("LightYellow{LightYellow}"),
		bunt.Sprintf("LemonChiffon{LemonChiffon}"),
		bunt.Sprintf("LightGoldenrodYellow{LightGoldenrodYellow}"),
		bunt.Sprintf("PapayaWhip{PapayaWhip}"),
		bunt.Sprintf("Moccasin{Moccasin}"),
		bunt.Sprintf("PeachPuff{PeachPuff}"),
		bunt.Sprintf("PaleGoldenrod{PaleGoldenrod}"),
		bunt.Sprintf("Khaki{Khaki}"),
		bunt.Sprintf("DarkKhaki{DarkKhaki}"),
		bunt.Sprintf("Gold{Gold}"),
		bunt.Sprintf("Cornsilk{Cornsilk}"),
		bunt.Sprintf("BlanchedAlmond{BlanchedAlmond}"),
		bunt.Sprintf("Bisque{Bisque}"),
		bunt.Sprintf("NavajoWhite{NavajoWhite}"),
		bunt.Sprintf("Wheat{Wheat}"),
		bunt.Sprintf("BurlyWood{BurlyWood}"),
		bunt.Sprintf("Tan{Tan}"),
		bunt.Sprintf("RosyBrown{RosyBrown}"),
		bunt.Sprintf("SandyBrown{SandyBrown}"),
		bunt.Sprintf("Goldenrod{Goldenrod}"),
		bunt.Sprintf("DarkGoldenrod{DarkGoldenrod}"),
		bunt.Sprintf("Peru{Peru}"),
		bunt.Sprintf("Chocolate{Chocolate}"),
		bunt.Sprintf("SaddleBrown{SaddleBrown}"),
		bunt.Sprintf("Sienna{Sienna}"),
		bunt.Sprintf("Brown{Brown}"),
		bunt.Sprintf("Maroon{Maroon}"),
		bunt.Sprintf("DarkOliveGreen{DarkOliveGreen}"),
		bunt.Sprintf("Olive{Olive}"),
		bunt.Sprintf("OliveDrab{OliveDrab}"),
		bunt.Sprintf("YellowGreen{YellowGreen}"),
		bunt.Sprintf("LimeGreen{LimeGreen}"),
		bunt.Sprintf("Lime{Lime}"),
		bunt.Sprintf("LawnGreen{LawnGreen}"),
		bunt.Sprintf("Chartreuse{Chartreuse}"),
		bunt.Sprintf("GreenYellow{GreenYellow}"),
		bunt.Sprintf("SpringGreen{SpringGreen}"),
		bunt.Sprintf("MediumSpringGreen{MediumSpringGreen}"),
		bunt.Sprintf("LightGreen{LightGreen}"),
		bunt.Sprintf("PaleGreen{PaleGreen}"),
		bunt.Sprintf("DarkSeaGreen{DarkSeaGreen}"),
		bunt.Sprintf("MediumAquamarine{MediumAquamarine}"),
		bunt.Sprintf("MediumSeaGreen{MediumSeaGreen}"),
		bunt.Sprintf("SeaGreen{SeaGreen}"),
		bunt.Sprintf("ForestGreen{ForestGreen}"),
		bunt.Sprintf("Green{Green}"),
		bunt.Sprintf("DarkGreen{DarkGreen}"),
		bunt.Sprintf("Aqua{Aqua}"),
		bunt.Sprintf("Cyan{Cyan}"),
		bunt.Sprintf("LightCyan{LightCyan}"),
		bunt.Sprintf("PaleTurquoise{PaleTurquoise}"),
		bunt.Sprintf("Aquamarine{Aquamarine}"),
		bunt.Sprintf("Turquoise{Turquoise}"),
		bunt.Sprintf("MediumTurquoise{MediumTurquoise}"),
		bunt.Sprintf("DarkTurquoise{DarkTurquoise}"),
		bunt.Sprintf("LightSeaGreen{LightSeaGreen}"),
		bunt.Sprintf("CadetBlue{CadetBlue}"),
		bunt.Sprintf("DarkCyan{DarkCyan}"),
		bunt.Sprintf("Teal{Teal}"),
		bunt.Sprintf("LightSteelBlue{LightSteelBlue}"),
		bunt.Sprintf("PowderBlue{PowderBlue}"),
		bunt.Sprintf("LightBlue{LightBlue}"),
		bunt.Sprintf("SkyBlue{SkyBlue}"),
		bunt.Sprintf("LightSkyBlue{LightSkyBlue}"),
		bunt.Sprintf("DeepSkyBlue{DeepSkyBlue}"),
		bunt.Sprintf("DodgerBlue{DodgerBlue}"),
		bunt.Sprintf("CornflowerBlue{CornflowerBlue}"),
		bunt.Sprintf("SteelBlue{SteelBlue}"),
		bunt.Sprintf("RoyalBlue{RoyalBlue}"),
		bunt.Sprintf("Blue{Blue}"),
		bunt.Sprintf("MediumBlue{MediumBlue}"),
		bunt.Sprintf("DarkBlue{DarkBlue}"),
		bunt.Sprintf("Navy{Navy}"),
		bunt.Sprintf("MidnightBlue{MidnightBlue}"),
		bunt.Sprintf("Lavender{Lavender}"),
		bunt.Sprintf("Thistle{Thistle}"),
		bunt.Sprintf("Plum{Plum}"),
		bunt.Sprintf("Violet{Violet}"),
		bunt.Sprintf("Orchid{Orchid}"),
		bunt.Sprintf("Fuchsia{Fuchsia}"),
		bunt.Sprintf("Magenta{Magenta}"),
		bunt.Sprintf("MediumOrchid{MediumOrchid}"),
		bunt.Sprintf("MediumPurple{MediumPurple}"),
		bunt.Sprintf("BlueViolet{BlueViolet}"),
		bunt.Sprintf("DarkViolet{DarkViolet}"),
		bunt.Sprintf("DarkOrchid{DarkOrchid}"),
		bunt.Sprintf("DarkMagenta{DarkMagenta}"),
		bunt.Sprintf("Purple{Purple}"),
		bunt.Sprintf("Indigo{Indigo}"),
		bunt.Sprintf("DarkSlateBlue{DarkSlateBlue}"),
		bunt.Sprintf("SlateBlue{SlateBlue}"),
		bunt.Sprintf("MediumSlateBlue{MediumSlateBlue}"),
		bunt.Sprintf("White{White}"),
		bunt.Sprintf("Snow{Snow}"),
		bunt.Sprintf("Honeydew{Honeydew}"),
		bunt.Sprintf("MintCream{MintCream}"),
		bunt.Sprintf("Azure{Azure}"),
		bunt.Sprintf("AliceBlue{AliceBlue}"),
		bunt.Sprintf("GhostWhite{GhostWhite}"),
		bunt.Sprintf("WhiteSmoke{WhiteSmoke}"),
		bunt.Sprintf("Seashell{Seashell}"),
		bunt.Sprintf("Beige{Beige}"),
		bunt.Sprintf("OldLace{OldLace}"),
		bunt.Sprintf("FloralWhite{FloralWhite}"),
		bunt.Sprintf("Ivory{Ivory}"),
		bunt.Sprintf("AntiqueWhite{AntiqueWhite}"),
		bunt.Sprintf("Linen{Linen}"),
		bunt.Sprintf("LavenderBlush{LavenderBlush}"),
		bunt.Sprintf("MistyRose{MistyRose}"),
		bunt.Sprintf("Gainsboro{Gainsboro}"),
		bunt.Sprintf("LightGray{LightGray}"),
		bunt.Sprintf("Silver{Silver}"),
		bunt.Sprintf("DarkGray{DarkGray}"),
		bunt.Sprintf("Gray{Gray}"),
		bunt.Sprintf("DimGray{DimGray}"),
		bunt.Sprintf("LightSlateGray{LightSlateGray}"),
		bunt.Sprintf("SlateGray{SlateGray}"),
		bunt.Sprintf("DarkSlateGray{DarkSlateGray}"),
		bunt.Sprintf("Black{Black}"),
		bunt.Sprintf("AdditionGreen{AdditionGreen}"),
		bunt.Sprintf("RemovalRed{RemovalRed}"),
		bunt.Sprintf("ModificationYellow{ModificationYellow}"),
	}

	// find the longest word in the list
	max := 0
	for _, entry := range list {
		if length := bunt.PlainTextLength(entry); length > max {
			max = length
		}
	}

	// find a table size with rows and columns that fits the terminal size
	columns := term.GetTerminalWidth() / max
	rows := (len(list) / columns) + 1

	// just a list with 0..n, for AlignCenter
	columnIdxs := make([]int, columns)
	for c := range columnIdxs {
		columnIdxs[c] = c
	}

	// prepare a table with the size of rows x columns
	table := make([][]string, rows)
	for r := range table {
		table[r] = make([]string, columns)
	}

	// fill the table
	for idx, entry := range list {
		c := idx / rows
		r := idx % rows
		table[r][c] = entry
	}

	// print the table
	if rendered, err := neat.Table(table, neat.DesiredWidth(term.GetTerminalWidth()), neat.AlignCenter(columnIdxs...)); err == nil {
		fmt.Println(rendered)
	}
}

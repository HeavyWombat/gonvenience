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
	"sort"

	"github.com/homeport/gonvenience/pkg/v1/bunt"
	"github.com/homeport/gonvenience/pkg/v1/neat"
	"github.com/homeport/gonvenience/pkg/v1/term"
	"github.com/lucasb-eyer/go-colorful"
)

var colorMap = map[colorful.Color]string{
	bunt.Pink:                 "Pink",
	bunt.LightPink:            "LightPink",
	bunt.HotPink:              "HotPink",
	bunt.DeepPink:             "DeepPink",
	bunt.PaleVioletRed:        "PaleVioletRed",
	bunt.MediumVioletRed:      "MediumVioletRed",
	bunt.LightSalmon:          "LightSalmon",
	bunt.Salmon:               "Salmon",
	bunt.DarkSalmon:           "DarkSalmon",
	bunt.LightCoral:           "LightCoral",
	bunt.IndianRed:            "IndianRed",
	bunt.Crimson:              "Crimson",
	bunt.FireBrick:            "FireBrick",
	bunt.DarkRed:              "DarkRed",
	bunt.Red:                  "Red",
	bunt.OrangeRed:            "OrangeRed",
	bunt.Tomato:               "Tomato",
	bunt.Coral:                "Coral",
	bunt.DarkOrange:           "DarkOrange",
	bunt.Orange:               "Orange",
	bunt.Yellow:               "Yellow",
	bunt.LightYellow:          "LightYellow",
	bunt.LemonChiffon:         "LemonChiffon",
	bunt.LightGoldenrodYellow: "LightGoldenrodYellow",
	bunt.PapayaWhip:           "PapayaWhip",
	bunt.Moccasin:             "Moccasin",
	bunt.PeachPuff:            "PeachPuff",
	bunt.PaleGoldenrod:        "PaleGoldenrod",
	bunt.Khaki:                "Khaki",
	bunt.DarkKhaki:            "DarkKhaki",
	bunt.Gold:                 "Gold",
	bunt.Cornsilk:             "Cornsilk",
	bunt.BlanchedAlmond:       "BlanchedAlmond",
	bunt.Bisque:               "Bisque",
	bunt.NavajoWhite:          "NavajoWhite",
	bunt.Wheat:                "Wheat",
	bunt.BurlyWood:            "BurlyWood",
	bunt.Tan:                  "Tan",
	bunt.RosyBrown:            "RosyBrown",
	bunt.SandyBrown:           "SandyBrown",
	bunt.Goldenrod:            "Goldenrod",
	bunt.DarkGoldenrod:        "DarkGoldenrod",
	bunt.Peru:                 "Peru",
	bunt.Chocolate:            "Chocolate",
	bunt.SaddleBrown:          "SaddleBrown",
	bunt.Sienna:               "Sienna",
	bunt.Brown:                "Brown",
	bunt.Maroon:               "Maroon",
	bunt.DarkOliveGreen:       "DarkOliveGreen",
	bunt.Olive:                "Olive",
	bunt.OliveDrab:            "OliveDrab",
	bunt.YellowGreen:          "YellowGreen",
	bunt.LimeGreen:            "LimeGreen",
	bunt.Lime:                 "Lime",
	bunt.LawnGreen:            "LawnGreen",
	bunt.Chartreuse:           "Chartreuse",
	bunt.GreenYellow:          "GreenYellow",
	bunt.SpringGreen:          "SpringGreen",
	bunt.MediumSpringGreen:    "MediumSpringGreen",
	bunt.LightGreen:           "LightGreen",
	bunt.PaleGreen:            "PaleGreen",
	bunt.DarkSeaGreen:         "DarkSeaGreen",
	bunt.MediumAquamarine:     "MediumAquamarine",
	bunt.MediumSeaGreen:       "MediumSeaGreen",
	bunt.SeaGreen:             "SeaGreen",
	bunt.ForestGreen:          "ForestGreen",
	bunt.Green:                "Green",
	bunt.DarkGreen:            "DarkGreen",
	bunt.Aqua:                 "Aqua",
	bunt.Cyan:                 "Cyan",
	bunt.LightCyan:            "LightCyan",
	bunt.PaleTurquoise:        "PaleTurquoise",
	bunt.Aquamarine:           "Aquamarine",
	bunt.Turquoise:            "Turquoise",
	bunt.MediumTurquoise:      "MediumTurquoise",
	bunt.DarkTurquoise:        "DarkTurquoise",
	bunt.LightSeaGreen:        "LightSeaGreen",
	bunt.CadetBlue:            "CadetBlue",
	bunt.DarkCyan:             "DarkCyan",
	bunt.Teal:                 "Teal",
	bunt.LightSteelBlue:       "LightSteelBlue",
	bunt.PowderBlue:           "PowderBlue",
	bunt.LightBlue:            "LightBlue",
	bunt.SkyBlue:              "SkyBlue",
	bunt.LightSkyBlue:         "LightSkyBlue",
	bunt.DeepSkyBlue:          "DeepSkyBlue",
	bunt.DodgerBlue:           "DodgerBlue",
	bunt.CornflowerBlue:       "CornflowerBlue",
	bunt.SteelBlue:            "SteelBlue",
	bunt.RoyalBlue:            "RoyalBlue",
	bunt.Blue:                 "Blue",
	bunt.MediumBlue:           "MediumBlue",
	bunt.DarkBlue:             "DarkBlue",
	bunt.Navy:                 "Navy",
	bunt.MidnightBlue:         "MidnightBlue",
	bunt.Lavender:             "Lavender",
	bunt.Thistle:              "Thistle",
	bunt.Plum:                 "Plum",
	bunt.Violet:               "Violet",
	bunt.Orchid:               "Orchid",
	bunt.Fuchsia:              "Fuchsia",
	bunt.Magenta:              "Magenta",
	bunt.MediumOrchid:         "MediumOrchid",
	bunt.MediumPurple:         "MediumPurple",
	bunt.BlueViolet:           "BlueViolet",
	bunt.DarkViolet:           "DarkViolet",
	bunt.DarkOrchid:           "DarkOrchid",
	bunt.DarkMagenta:          "DarkMagenta",
	bunt.Purple:               "Purple",
	bunt.Indigo:               "Indigo",
	bunt.DarkSlateBlue:        "DarkSlateBlue",
	bunt.SlateBlue:            "SlateBlue",
	bunt.MediumSlateBlue:      "MediumSlateBlue",
	bunt.White:                "White",
	bunt.Snow:                 "Snow",
	bunt.Honeydew:             "Honeydew",
	bunt.MintCream:            "MintCream",
	bunt.Azure:                "Azure",
	bunt.AliceBlue:            "AliceBlue",
	bunt.GhostWhite:           "GhostWhite",
	bunt.WhiteSmoke:           "WhiteSmoke",
	bunt.Seashell:             "Seashell",
	bunt.Beige:                "Beige",
	bunt.OldLace:              "OldLace",
	bunt.FloralWhite:          "FloralWhite",
	bunt.Ivory:                "Ivory",
	bunt.AntiqueWhite:         "AntiqueWhite",
	bunt.Linen:                "Linen",
	bunt.LavenderBlush:        "LavenderBlush",
	bunt.MistyRose:            "MistyRose",
	bunt.Gainsboro:            "Gainsboro",
	bunt.LightGray:            "LightGray",
	bunt.Silver:               "Silver",
	bunt.DarkGray:             "DarkGray",
	bunt.Gray:                 "Gray",
	bunt.DimGray:              "DimGray",
	bunt.LightSlateGray:       "LightSlateGray",
	bunt.SlateGray:            "SlateGray",
	bunt.DarkSlateGray:        "DarkSlateGray",
	bunt.Black:                "Black",
	bunt.AdditionGreen:        "AdditionGreen",
	bunt.RemovalRed:           "RemovalRed",
	bunt.ModificationYellow:   "ModificationYellow",
}

func main() {
	// compile a list of the colors
	colors := []colorful.Color{}
	for color := range colorMap {
		colors = append(colors, color)
	}

	// sort based on their 4-bit equivalent
	sort.Slice(colors, func(i, j int) bool {
		return bunt.Get4bitEquivalentColorAttribute(colors[i]) < bunt.Get4bitEquivalentColorAttribute(colors[j])
	})

	// create a list of examples
	list := []string{}
	for _, color := range colors {
		name := colorMap[color]
		list = append(list,
			fmt.Sprintf("%s (%s)",
				bunt.Colorize(name, color),
				bunt.Colorize4bit("4bit", color),
			),
		)
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

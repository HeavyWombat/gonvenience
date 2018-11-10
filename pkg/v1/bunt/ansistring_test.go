// Copyright © 2018 Matthias Diester
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

package bunt_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/HeavyWombat/gonvenience/pkg/v1/bunt"
)

var _ = Describe("ANSI string tests", func() {
	BeforeEach(func() {
		ColorSetting = ON
		TrueColorSetting = ON
	})

	It("should preserve already existing styles in an input string", func() {
		alreadyStyledText := "Preserve styles like \x1b[1mbold\x1b[0m, \x1b[3mitalic\x1b[0m, \x1b[1;38;2;255;128;16mcolors\x1b[0m, and \x1b[4;38;2;128;128;128mstyled colors\x1b[0m in input strings."
		Expect(Sprint(alreadyStyledText)).To(BeEquivalentTo(alreadyStyledText))
	})

	It("should translate *bold* syntax into styled text", func() {
		Expect(Sprint(">>> *bold* <<<")).To(BeEquivalentTo(
			fmt.Sprintf(">>> %s <<<",
				Style("bold", Bold),
			)))
	})

	It("should translate _italic_ syntax into styled text", func() {
		Expect(Sprint(">>> _italic_ <<<")).To(BeEquivalentTo(
			fmt.Sprintf(">>> %s <<<",
				Style("italic", Italic),
			)))
	})

	It("should translate ~underline~ syntax into styled text", func() {
		Expect(Sprint(">>> ~underline~ <<<")).To(BeEquivalentTo(
			fmt.Sprintf(">>> %s <<<",
				Style("underline", Underline),
			)))
	})

	It("should translate Blue{Blue} color syntax into styled text", func() {
		Expect(Sprint(">>> Blue{Blue} <<<")).To(BeEquivalentTo(
			fmt.Sprintf(">>> %s <<<",
				Colorize("Blue", Blue),
			)))
	})

	It("should translate multiple occurrences of *bold* into styled text", func() {
		Expect(Sprint("Text *with* multiple *bold* elements.")).To(BeEquivalentTo(
			fmt.Sprintf("Text %s multiple %s elements.",
				Style("with", Bold),
				Style("bold", Bold),
			)))
	})

	It("should translate *bold* and _italic_ syntax into styled text", func() {
		Expect(Sprint("Use *bold* and _italic_ in texts.")).To(BeEquivalentTo(
			fmt.Sprintf("Use %s and %s in texts.",
				Style("bold", Bold),
				Style("italic", Italic),
			)))
	})

	It("should translate nested text marker in color marker into the correct output", func() {
		Expect(Sprint(" --- >>>CornflowerBlue{~http://github.com~}<<< ---")).To(BeEquivalentTo(
			fmt.Sprintf(" --- >>>\x1b[38;2;100;149;237m\x1b[4m%s\x1b[0m\x1b[0m<<< ---",
				"http://github.com",
			)))
	})

	It("should translate nested color marker in text marker into the correct output", func() {
		Expect(Sprint(" --- >>>~CornflowerBlue{http://github.com}~<<< ---")).To(BeEquivalentTo(
			fmt.Sprintf(" --- >>>\x1b[4m\x1b[38;2;100;149;237m%s\x1b[0m\x1b[0m<<< ---",
				"http://github.com",
			)))
	})
})

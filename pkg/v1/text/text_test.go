// Copyright © 2019 The Homeport Team
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

package text_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/homeport/gonvenience/pkg/v1/text"
)

var _ = Describe("Generate random strings with fixed length", func() {
	Context("Random string with no prefix", func() {
		It("should generate a random string with fixed length", func() {
			Expect(len(RandomString(32))).To(BeEquivalentTo(32))
		})

		It("should fail when negative length is given", func() {
			defer func() {
				Expect(recover()).To(HaveOccurred())
			}()

			RandomString(-1)
		})
	})

	Context("Random string with given prefix", func() {
		It("should generate a random string with fixed length", func() {
			Expect(len(RandomStringWithPrefix("foobar", 32))).To(BeEquivalentTo(32))
		})

		It("should fail when the prefix is already longer than the fixed length", func() {
			defer func() {
				Expect(recover()).To(HaveOccurred())
			}()

			RandomStringWithPrefix("foobar", 4)
		})

		It("should fail when negative length is given", func() {
			defer func() {
				Expect(recover()).To(HaveOccurred())
			}()

			RandomStringWithPrefix("foobar", -1)
		})
	})

	Context("Text with given fixed length", func() {
		It("should create a string with the text and enough padding to fill it up to the required length", func() {
			Expect(FixedLength("Foobar", 10)).To(BeEquivalentTo("Foobar    "))
		})

		It("should trim the text if the text alone exceeds the provided desired length", func() {
			Expect(FixedLength("This text is too long", 10)).To(BeEquivalentTo("This [...]"))
		})

		It("should return the text as-is if it already has the perfect length", func() {
			Expect(FixedLength("Foobar", 6)).To(BeEquivalentTo("Foobar"))
		})
	})
})

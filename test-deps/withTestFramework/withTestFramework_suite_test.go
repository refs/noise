package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWithTestFramework(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WithTestFramework Suite")
}

var _ = Describe("Book", func() {
	r := sum(2, 2)

	Describe("Sum", func() {
		Context("With correct result", func() {
			It("should pass", func() {
				Expect(r).To(Equal(4))
			})
		})
	})
})

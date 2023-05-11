package main

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Test", func() {
	gomega.Expect(1).To(gomega.Equal(1))
})

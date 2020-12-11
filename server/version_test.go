package server_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// The actual test suite
var _ = t.Describe("Version", func() {
	// Prepare the sut
	BeforeEach(func() {
		beforeEach()
		setupSUT()
	})
	AfterEach(afterEach)

	t.Describe("Version", func() {
		It("should succeed", func() {
			// Given
			// When
			response, err := sut.Version(context.Background())

			// Then
			Expect(err).To(BeNil())
			Expect(response).NotTo(BeNil())
			Expect(response.Version).NotTo(BeEmpty())
			Expect(response.RuntimeName).NotTo(BeEmpty())
			Expect(response.RuntimeName).NotTo(BeEmpty())
			Expect(response.RuntimeAPIVersion).NotTo(BeEmpty())
		})
	})
})

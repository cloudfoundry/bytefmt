package bytefmt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-golang/bytefmt"
)

var _ = Describe("bytefmt", func() {
	Context("BPSSize", func() {
		It("Prints in the largest possible unit", func() {
			Expect(BPSSize(10 * TERABITSPS)).To(Equal("10Tbps"))
			Expect(BPSSize(10.5 * TERABITSPS)).To(Equal("10.5Tbps"))

			Expect(BPSSize(10 * GIGABITSPS)).To(Equal("10Gbps"))
			Expect(BPSSize(10.5 * GIGABITSPS)).To(Equal("10.5Gbps"))

			Expect(BPSSize(100 * MEGABITSPS)).To(Equal("100Mbps"))
			Expect(BPSSize(100.5 * MEGABITSPS)).To(Equal("100.5Mbps"))

			Expect(BPSSize(100 * KILOBITSPS)).To(Equal("100Kbps"))
			Expect(BPSSize(100.5 * KILOBITSPS)).To(Equal("100.5Kbps"))

			Expect(BPSSize(1)).To(Equal("1bps"))
		})

		It("prints '0' for zero bytes", func() {
			Expect(BPSSize(0)).To(Equal("0"))
		})
	})
})

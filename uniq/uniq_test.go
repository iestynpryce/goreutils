package main_test

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Uniq", func() {
	Context("when presented with a duplicate free file", func() {
		var fname string = "test_nodups"
		var benchmark string = "one\ntwo\nthree\n"

		BeforeEach(func() {
			f, err := os.OpenFile(fname,
				os.O_CREATE|os.O_WRONLY,
				os.FileMode(0666))
			Expect(err).NotTo(HaveOccurred())
			defer f.Close()
			f.WriteString(benchmark)
		})

		AfterEach(func() {
			err := os.Remove(fname)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should output all the contents of the file", func() {
			command := exec.Command(pathToUniq, fname)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say(benchmark))
		})
	})

	Context("when presented with a file with duplicates on adjacent lines", func() {
		var fname string = "test_adjdups"
		var input string = "one\none\ntwo\nthree\nthree"
		var benchmark string = "one\ntwo\nthree\n"

		BeforeEach(func() {
			f, err := os.OpenFile(fname,
				os.O_CREATE|os.O_WRONLY,
				os.FileMode(0666))
			Expect(err).NotTo(HaveOccurred())
			defer f.Close()
			f.WriteString(input)
		})

		AfterEach(func() {
			err := os.Remove(fname)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should output all but duplicate lines in the file", func() {
			command := exec.Command(pathToUniq, fname)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say(benchmark))
		})
	})

	Context("when presented with a file with duplicates on non-adjacent lines", func() {
		var fname string = "test_adjdups"
		var benchmark string = "one\ntwo\none\n"

		BeforeEach(func() {
			f, err := os.OpenFile(fname,
				os.O_CREATE|os.O_WRONLY,
				os.FileMode(0666))
			Expect(err).NotTo(HaveOccurred())
			defer f.Close()
			f.WriteString(benchmark)
		})

		AfterEach(func() {
			err := os.Remove(fname)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should output all the content of the file", func() {
			command := exec.Command(pathToUniq, fname)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say(benchmark))
		})
	})

	Context("when presented with data from stdin", func() {
		It("should remove adjacent duplicates", func() {
			var input string = "dog\ndog\ncat\ncat\ndog\nfish\n"
			var output string = "dog\ncat\ndog\nfish\n"
			command := exec.Command(pathToUniq)
			stdin, err := command.StdinPipe()
			Expect(err).NotTo(HaveOccurred())
			defer stdin.Close()

			session, _ := gexec.Start(command, GinkgoWriter, GinkgoWriter)

			fmt.Fprintf(stdin, input)
			Eventually(session.Out).Should(gbytes.Say(output))
		})
	})

	Context("when presented with an unreadable file", func() {
		var fname string = "test_badfile"

		BeforeEach(func() {
			f, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, os.FileMode(0000))
			Expect(err).NotTo(HaveOccurred())
			f.Close()
		})

		AfterEach(func() {
			err := os.Remove(fname)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return a useful error message", func() {

			command := exec.Command(pathToUniq, fname)

			session, _ := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Eventually(session.Err).Should(gbytes.Say(".*Error:.*permission denied\n"))
			session.Wait()
		})
	})
})

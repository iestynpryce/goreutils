package main_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
		var fout string = "test_nodups_out"
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

			if _, e := os.Stat(fout); e == nil {
				err = os.Remove(fout)
				Expect(err).NotTo(HaveOccurred())
			}
		})

		It("should output all the contents of the file (stdout)", func() {
			command := exec.Command(pathToUniq, fname)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say(benchmark))
		})

		It("should output all the contents of the file (outfile)", func() {
			command := exec.Command(pathToUniq, fname, fout)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			session.Wait()

			buf, _ := ioutil.ReadFile(fout)
			Expect(bytes.Equal(buf, []byte(benchmark))).To(Equal(true))
		})
	})

	Context("when presented with a file with duplicates on adjacent lines", func() {
		var fname string = "test_adjdups"
		var fout string = "test_adjdups_out"
		var input string = "one\none\ntwo\nthree\nthree\n"
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

			if _, e := os.Stat(fout); e == nil {
				err = os.Remove(fout)
				Expect(err).NotTo(HaveOccurred())
			}
		})

		It("should output all but duplicate lines in the file (stdout)", func() {
			command := exec.Command(pathToUniq, fname)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say(benchmark))
		})

		It("should output all but duplicate lines in the file (outfile)", func() {
			command := exec.Command(pathToUniq, fname, fout)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			session.Wait()

			buf, _ := ioutil.ReadFile(fout)
			Expect(bytes.Equal(buf, []byte(benchmark))).To(Equal(true))
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

	Context("when using the -d option", func() {
		It("should only print duplicate lines, one for each group",
			func() {
				var input string = "foo\nfoo\nfoo\bar\n"
				var output string = "foo\n"
				command := exec.Command(pathToUniq, "-d")
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

		It("should return a useful error message when reading", func() {

			command := exec.Command(pathToUniq, fname)

			session, _ := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Eventually(session.Err).Should(gbytes.Say(".*Error:.*permission denied\n"))
			session.Wait()
		})

		It("should return a useful error message when writing", func() {
			var fread string = "test_file"
			f, err := os.OpenFile(fread, os.O_CREATE|os.O_WRONLY, os.FileMode(0666))
			Expect(err).NotTo(HaveOccurred())
			f.Close()

			command := exec.Command(pathToUniq, fread, fname)

			session, _ := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Eventually(session.Err).Should(gbytes.Say(".*Error:.*permission denied\n"))
			session.Wait()

			err = os.Remove(fread)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})

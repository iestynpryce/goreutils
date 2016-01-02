package main_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Tee", func() {

	Context("when presented with a normal file", func() {
		var command *exec.Cmd
		var fname string = "test_tee.out"

		BeforeEach(func() {
			command = exec.Command(pathToTee, fname)
		})

		AfterEach(func() {
			err := os.Remove(fname)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return no error when run on a normal file", func() {
			stdin, e := command.StdinPipe()
			Expect(e).NotTo(HaveOccurred())
			defer stdin.Close()

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			fmt.Fprintf(stdin, "test string\n")
			Eventually(session.Out).Should(gbytes.Say("test string.*"))
		})
	})

	Context("when appending an existing file", func() {
		var command *exec.Cmd
		var fname = "test_append.out"

		BeforeEach(func() {
			command = exec.Command(pathToTee,
				"-a",
				fname)

			// Create an empty file
			f, _ := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY,
				os.FileMode(0666))
			f.Close()
		})

		AfterEach(func() {
			err := os.Remove(fname)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should append to an empty file", func() {
			stdin, e := command.StdinPipe()
			Expect(e).NotTo(HaveOccurred())
			defer stdin.Close()

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			fmt.Fprintf(stdin, "test\n")
			Eventually(session).Should(gbytes.Say(".*test.*"))

			buf, _ := ioutil.ReadFile(fname)
			benchmark := []byte("test\n")
			Expect(bytes.Equal(buf, benchmark)).To(Equal(true))
		})

		FIt("should append to a file with content", func() {
			stdin, e := command.StdinPipe()
			Expect(e).NotTo(HaveOccurred())
			defer stdin.Close()

			e = ioutil.WriteFile(fname, []byte("test\n"), os.FileMode(0666))
			Expect(e).NotTo(HaveOccurred())

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			fmt.Fprintf(stdin, "another test\n")
			Eventually(session).Should(gbytes.Say("another test\n"))

			buf, _ := ioutil.ReadFile(fname)
			benchmark := []byte("test\nanother test\n")
			Expect(bytes.Equal(buf, benchmark)).To(Equal(true))
		})
	})

	Context("when presented with an invalid file", func() {
		var command *exec.Cmd
		var session *gexec.Session
		var stdin io.WriteCloser

		rand.Seed(time.Now().UTC().UnixNano())
		const chars = "abcdefghijklmnopqrstuvwxyx0123456789"
		nameByte := make([]byte, 64)
		for i := 0; i < 64; i++ {
			nameByte[i] = chars[rand.Intn(len(chars))]
		}
		name := string(nameByte)

		JustBeforeEach(func() {
			var err error

			f, ferr := os.OpenFile(name, os.O_CREATE, os.FileMode(0000))
			if ferr != nil {
				log.Fatalf("Error: %v\n", ferr)
			}
			f.Close()

			command = exec.Command(pathToTee, name)
			stdin, err = command.StdinPipe()
			Expect(err).NotTo(HaveOccurred())
			defer stdin.Close()

			session, _ = gexec.Start(command, GinkgoWriter,
				GinkgoWriter)
		})

		AfterEach(func() {
			err := os.Remove(name)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return a useful error message", func() {
			Eventually(session.Err).Should(gbytes.Say(".*Error:.*"))
			fmt.Fprintf(stdin, "test")
			session.Wait()
		})

		It("should return a non zero return code", func() {
			Eventually(session).ShouldNot(gexec.Exit(0))
			session.Wait()
		})

	})

})

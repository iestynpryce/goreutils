package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	//	"time"
	//. "github.com/iestynpryce/goreutils/nl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Nl", func() {

	It("should return no error when run on a normal file", func() {
		command := exec.Command(pathToNl)
		_, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("when presented with an invalid file", func() {
		var command *exec.Cmd
		var session *gexec.Session

		BeforeEach(func() {
			command = exec.Command(pathToNl, "doesNotExist")
			session, _ = gexec.Start(command, GinkgoWriter,
				GinkgoWriter)
		})

		It("should return a non zero return code", func() {
			Eventually(session).Should(gexec.Exit(-1))
		})

		It("should return a useful error message", func() {
			Eventually(session.Err).Should(gbytes.Say("open doesNotExist: no such file or directory"))
		})
	})

	Context("when presented with a valid file", func() {
		var session *gexec.Session
		var fileName string
		var fileHandle *os.File

		BeforeEach(func() {
			fileHandle, _ = ioutil.TempFile("", "")
			defer fileHandle.Close()
			fileName = fileHandle.Name()

			fileHandle.WriteString("line 1\nline 2\n    \nline 3\n")
			fileHandle.Sync()

			command := exec.Command(pathToNl, fileName)
			session, _ = gexec.Start(command, GinkgoWriter,
				GinkgoWriter)
		})

		AfterEach(func() {
			os.Remove(fileName)
		})

		It("should return line numbers at the start of non-blank lines",
			func() {
				Eventually(session.Out).Should(gbytes.Say("\\s+1.*\n\\s+2"))
			})

		FIt("should not return line numbers at the start of blank lines",
			func() {
				Eventually(session.Out).Should(gbytes.Say("\\s+1.*\n\\s+2.*\n    \n\\s+3"))
			})
	})
})

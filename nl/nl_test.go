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

		It("should not return line numbers at the start of blank lines",
			func() {
				Eventually(session.Out).Should(gbytes.Say("\\s+1.*\n\\s+2.*\n    \n\\s+3"))
			})
	})

	Context("when using the -i option", func() {
		var session *gexec.Session
		var fileName string
		var fileHandle *os.File

		BeforeEach(func() {
			fileHandle, _ = ioutil.TempFile("", "")
			defer fileHandle.Close()
			fileName = fileHandle.Name()

			fileHandle.WriteString("line 1\nline 2\nline 3\n")
			fileHandle.Sync()

		})

		AfterEach(func() {
			os.Remove(fileName)
		})

		It("should return line numbers incremented by 1",
			func() {
				localCommand := exec.Command(pathToNl, "-i", "1", fileName)
				session, _ = gexec.Start(localCommand, GinkgoWriter,
					GinkgoWriter)
				Eventually(session.Out).Should(gbytes.Say("\\s+1.*\n\\s+2.*\n\\s+3"))
			})

		It("should return line numbers incremented by 2",
			func() {
				localCommand := exec.Command(pathToNl, "-i", "2", fileName)
				session, _ = gexec.Start(localCommand, GinkgoWriter,
					GinkgoWriter)
				Eventually(session.Out).Should(gbytes.Say("\\s+1.*\n\\s+3.*\n\\s+5"))
			})

		It("should return line numbers incremented by 10",
			func() {
				localCommand := exec.Command(pathToNl, "-i", "10", fileName)
				session, _ = gexec.Start(localCommand, GinkgoWriter,
					GinkgoWriter)
				Eventually(session.Out).Should(gbytes.Say("\\s+1.*\n\\s+11.*\n\\s+21"))
			})

		It("should return an error when line numbers incremented by -1",
			func() {
				localCommand := exec.Command(pathToNl, "-i", "-1", fileName)
				session, _ = gexec.Start(localCommand, GinkgoWriter,
					GinkgoWriter)
				Eventually(session.Err).Should(gbytes.Say("Error: invalid line-increment value \"-1\""))
			})

		It("should return an error when line numbers incremented by 0",
			func() {
				localCommand := exec.Command(pathToNl, "-i", "0", fileName)
				session, _ = gexec.Start(localCommand, GinkgoWriter,
					GinkgoWriter)
				Eventually(session.Err).Should(gbytes.Say("Error: invalid line-increment value \"0\""))
			})
	})

	Context("when using the -s option", func() {
		var session *gexec.Session
		var fileName string
		var fileHandle *os.File

		BeforeEach(func() {
			fileHandle, _ = ioutil.TempFile("", "")
			defer fileHandle.Close()
			fileName = fileHandle.Name()

			fileHandle.WriteString("line 1\nline 2\nline 3\n")
			fileHandle.Sync()

		})

		AfterEach(func() {
			os.Remove(fileName)
		})

		It("should return line numbers followed by \\t by default",
			func() {
				localCommand := exec.Command(pathToNl, fileName)
				session, _ = gexec.Start(localCommand, GinkgoWriter,
					GinkgoWriter)
				Eventually(session.Out).Should(gbytes.Say("\\s+1\t.*\n\\s+2\t.*\n\\s+3\t"))
			})

		It("should return line numbers followed by \\t",
			func() {
				localCommand := exec.Command(pathToNl, "-s", "\t", fileName)
				session, _ = gexec.Start(localCommand, GinkgoWriter,
					GinkgoWriter)
				Eventually(session.Out).Should(gbytes.Say("\\s+1\t.*\n\\s+2\t.*\n\\s+3\t"))
			})

		It("should return line numbers followed by no space",
			func() {
				localCommand := exec.Command(pathToNl, "-s", "", fileName)
				session, _ = gexec.Start(localCommand, GinkgoWriter,
					GinkgoWriter)
				Eventually(session.Out).Should(gbytes.Say("\\s+1line.*\n\\s+2line.*\n\\s+3line"))
			})

		It("should return line numbers followed by --",
			func() {
				localCommand := exec.Command(pathToNl, "-s", "--", fileName)
				session, _ = gexec.Start(localCommand, GinkgoWriter,
					GinkgoWriter)
				Eventually(session.Out).Should(gbytes.Say("\\s+1--.*\n\\s+2--.*\n\\s+3--"))
			})
	})

})

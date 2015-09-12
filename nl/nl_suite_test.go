package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var pathToNl string

func TestNl(t *testing.T) {

	BeforeSuite(func() {
		var err error
		pathToNl, err = gexec.Build("github.com/iestynpryce/goreutils/nl")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Nl Suite")
}

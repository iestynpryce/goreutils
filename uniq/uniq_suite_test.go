package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var pathToUniq string

func TestUniq(t *testing.T) {

	BeforeSuite(func() {
		var err error
		pathToUniq, err = gexec.Build("github.com/iestynpryce/goreutils/uniq")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Uniq Suite")
}

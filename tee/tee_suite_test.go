package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var pathToTee string

func TestTee(t *testing.T) {

	BeforeSuite(func() {
		var err error
		pathToTee, err = gexec.Build("github.com/iestynpryce/goreutils/tee")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Tee Suite")
}

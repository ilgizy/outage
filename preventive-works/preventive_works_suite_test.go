package preventive_works_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPreventiveWorks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PreventiveWorks Suite")
}

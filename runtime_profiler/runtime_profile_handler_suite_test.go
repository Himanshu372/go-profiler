package runtime_profiler_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRuntimeProfiler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RuntimeProfiler Suite")
}

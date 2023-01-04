package integration_test

import (
	"github.com/mvach/bosh-azure-storage-cli/integration"
	"github.com/onsi/gomega/gexec"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var cliPath string
var largeContent string

var _ = BeforeSuite(func() {
	cliPath = os.Getenv("AZURE_STORAGE_CLI_PATH")
	largeContent = integration.GenerateRandomString(1024 * 1024 * 6)

	if len(cliPath) == 0 {
		var err error
		cliPath, err = gexec.Build("github.com/mvach/bosh-azure-storage-cli")
		Expect(err).ShouldNot(HaveOccurred())
	}
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

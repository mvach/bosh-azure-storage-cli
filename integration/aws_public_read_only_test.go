package integration_test

import (
	"os"

	"github.com/mvach/bosh-azure-storage-cli/config"
	"github.com/mvach/bosh-azure-storage-cli/integration"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing gets against a public AWS S3 bucket", func() {
	Context("with PUBLIC READ ONLY (no creds) configuration", func() {
		It("can successfully get a publicly readable file", func() {
			accountName := os.Getenv("ACCOUNT_NAME")
			Expect(accountName).ToNot(BeEmpty(), "ACCOUNT_NAME must be set")

			accountKey := os.Getenv("ACCOUNT_KEY")
			Expect(accountKey).ToNot(BeEmpty(), "ACCOUNT_KEY must be set")

			containerName := os.Getenv("CONTAINER_NAME")
			Expect(containerName).ToNot(BeEmpty(), "CONTAINER_NAME must be set")

			s3Filename := integration.GenerateRandomString()
			s3FileContents := integration.GenerateRandomString()

			cfg := &config.AZStorageConfig{
				AccountName:   accountName,
				AccountKey:    accountKey,
				ContainerName: containerName,
			}

			configPath := integration.MakeConfigFile(cfg)
			defer func() { _ = os.Remove(configPath) }()

			contentFile := integration.MakeContentFile(s3FileContents)
			defer func() { _ = os.Remove(contentFile) }()

			s3CLISession, err := integration.RunCli(cliPath, configPath, "put", contentFile, s3Filename)
			Expect(err).ToNot(HaveOccurred())
			Expect(s3CLISession.ExitCode()).ToNot(BeZero())

			s3CLISession, err = integration.RunCli(cliPath, configPath, "get", s3Filename, "public-file")
			Expect(err).ToNot(HaveOccurred())

			defer func() { _ = os.Remove("public-file") }()
			Expect(s3CLISession.ExitCode()).To(BeZero())

			gottenBytes, err := os.ReadFile("public-file")
			Expect(err).ToNot(HaveOccurred())
			Expect(string(gottenBytes)).To(Equal(s3FileContents))

			s3CLISession, err = integration.RunCli(cliPath, configPath, "exists", s3Filename)
			Expect(err).ToNot(HaveOccurred())
			Expect(s3CLISession.ExitCode()).To(BeZero())
			Expect(s3CLISession.Err.Contents()).To(MatchRegexp("File '.*' exists in bucket '.*'"))
		})
	})
})
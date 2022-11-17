package client_test

import (
	"errors"
	"github.com/mvach/bosh-azure-storage-cli/blob"
	"github.com/mvach/bosh-azure-storage-cli/client"
	"github.com/mvach/bosh-azure-storage-cli/client/clientfakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Client", func() {

	It("put file uploads to a blob", func() {
		storageClient := clientfakes.FakeStorageClient{}

		azBlobstore, err := client.New(&storageClient)
		Expect(err).ToNot(HaveOccurred())

		file, _ := os.CreateTemp("", "tmpfile")

		azBlobstore.Put(file, "target/blob")

		Expect(storageClient.UploadCallCount()).To(Equal(1))
		source, dest := storageClient.UploadArgsForCall(0)

		Expect(source).To(Equal(file))
		Expect(dest).To(Equal("target/blob"))
	})

	It("get blob downloads to a file", func() {
		storageClient := clientfakes.FakeStorageClient{}

		azBlobstore, err := client.New(&storageClient)
		Expect(err).ToNot(HaveOccurred())

		file, _ := os.CreateTemp("", "tmpfile")

		azBlobstore.Get("source/blob", file)

		Expect(storageClient.DownloadCallCount()).To(Equal(1))
		source, dest := storageClient.DownloadArgsForCall(0)

		Expect(source).To(Equal("source/blob"))
		Expect(dest).To(Equal(file))
	})

	It("delete blob deletes the blob", func() {
		storageClient := clientfakes.FakeStorageClient{}

		azBlobstore, err := client.New(&storageClient)
		Expect(err).ToNot(HaveOccurred())

		azBlobstore.Delete("blob")

		Expect(storageClient.DeleteCallCount()).To(Equal(1))
		dest := storageClient.DeleteArgsForCall(0)

		Expect(dest).To(Equal("blob"))
	})

	Context("if the blob existence is checked", func() {
		It("returns blob.Existing on success", func() {
			storageClient := clientfakes.FakeStorageClient{}
			storageClient.ExistsReturns(blob.Existing, nil)

			azBlobstore, _ := client.New(&storageClient)
			existsState, err := azBlobstore.Exists("blob")
			Expect(existsState == blob.Existing).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())

			dest := storageClient.ExistsArgsForCall(0)
			Expect(dest).To(Equal("blob"))
		})

		It("returns blob.NotExisting for not existing blobs", func() {
			storageClient := clientfakes.FakeStorageClient{}
			storageClient.ExistsReturns(blob.NotExisting, nil)

			azBlobstore, _ := client.New(&storageClient)
			existsState, err := azBlobstore.Exists("blob")
			Expect(existsState == blob.NotExisting).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())

			dest := storageClient.ExistsArgsForCall(0)
			Expect(dest).To(Equal("blob"))
		})

		It("returns blob.ExistenceUnknown and an error in case an error occured", func() {
			storageClient := clientfakes.FakeStorageClient{}
			storageClient.ExistsReturns(blob.ExistenceUnknown, errors.New("booom"))

			azBlobstore, _ := client.New(&storageClient)
			existsState, err := azBlobstore.Exists("blob")
			Expect(existsState == blob.ExistenceUnknown).To(BeTrue())
			Expect(err).To(HaveOccurred())

			dest := storageClient.ExistsArgsForCall(0)
			Expect(dest).To(Equal("blob"))
		})
	})

})

package client_test

import (
	"github.com/mvach/bosh-azure-storage-cli/client"
	"github.com/mvach/bosh-azure-storage-cli/client/clientfakes"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Config", func() {

	It("put", func() {
		storageClient := clientfakes.FakeStorageClient{}

		client.New(&storageClient)
		//azBlobstore.Put(...)
	})

})

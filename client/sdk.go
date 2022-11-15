package client

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/mvach/bosh-azure-storage-cli/config"
)

func NewAZClient(storageConfig config.AZStorageConfig) (*blockblob.Client, error) {
	credential, err := azblob.NewSharedKeyCredential(storageConfig.AccountName, storageConfig.AccountKey)
	if err != nil {
		return nil, err
	}

	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", storageConfig.AccountName)

	client, err := blockblob.NewClientWithSharedKeyCredential(serviceURL, credential, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

package client

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/mvach/bosh-azure-storage-cli/config"
)

type AzBlobstore struct {
	azClient      *azblob.Client
	storageConfig *config.AZStorageConfig
}

func New(azClient *azblob.Client, storageConfig *config.AZStorageConfig) (AzBlobstore, error) {
	return AzBlobstore{azClient: azClient, storageConfig: storageConfig}, nil
}

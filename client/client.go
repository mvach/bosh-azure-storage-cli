package client

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/mvach/bosh-azure-storage-cli/config"
	"log"
	"os"
)

type AzBlobstore struct {
	azClient      *azblob.Client
	storageConfig *config.AZStorageConfig
}

func New(azClient *azblob.Client, storageConfig *config.AZStorageConfig) (AzBlobstore, error) {
	return AzBlobstore{azClient: azClient, storageConfig: storageConfig}, nil
}

func (client *AzBlobstore) Put(file *os.File, dest string) error {
	//TODO: Check if we need to set a "WithTimeout(parent Context, timeout time.Duration)" context here:
	response, err := client.azClient.UploadFile(context.Background(), "test2", "dev1", file, nil)
	if err == nil {
		log.Println(response)
	}

	return err
}

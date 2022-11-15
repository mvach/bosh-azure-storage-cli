package client

import (
	"context"
	"os"
)

type AzBlobstore struct {
	storageClient StorageClient
}

func New(storageClient StorageClient) (AzBlobstore, error) {
	return AzBlobstore{storageClient: storageClient}, nil
}

func (client *AzBlobstore) Put(sourceFile *os.File, destPath string) error {

	_, err := client.storageClient.Upload(context.Background(), sourceFile, destPath, nil)

	return err
}

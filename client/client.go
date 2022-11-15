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

func (client *AzBlobstore) Put(file *os.File, dest string) error {

	_, err := client.storageClient.Upload(context.Background(), file, nil)

	return err
}

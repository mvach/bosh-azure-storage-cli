package client

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"io"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . StorageClient
type StorageClient interface {
	Upload(
		context context.Context,
		body io.ReadSeekCloser,
		options *blockblob.UploadOptions) (StorageResponse, error)
}

type DefaultStorageClient struct {
	azClient blockblob.Client
}

func NewStorageClient(azClient blockblob.Client) (StorageClient, error) {
	return DefaultStorageClient{azClient: azClient}, nil
}

func (dsc DefaultStorageClient) Upload(
	context context.Context,
	body io.ReadSeekCloser,
	options *blockblob.UploadOptions) (StorageResponse, error) {

	resp, err := dsc.azClient.Upload(context, body, options)

	return StorageResponse{
		ClientRequestID:     resp.ClientRequestID,
		ContentMD5:          resp.ContentMD5,
		Date:                resp.Date,
		ETag:                resp.ETag,
		EncryptionKeySHA256: resp.EncryptionKeySHA256,
		EncryptionScope:     resp.EncryptionScope,
		IsServerEncrypted:   resp.IsServerEncrypted,
		LastModified:        resp.LastModified,
		RequestID:           resp.RequestID,
		Version:             resp.Version,
		VersionID:           resp.VersionID,
	}, err
}

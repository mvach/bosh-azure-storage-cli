# Azure Storage CLI

The Azure Storage CLI is for uploading, fetching and deleting content to and from an Azure blobstore.
It is highly inspired by the https://github.com/cloudfoundry/bosh-s3cli.

## Usage

Given a JSON config file (`config.json`)...

``` json
{
  "account-name":           "<string> (required)",
  "account-key":            "<string> (required)",
  "container-name":         "<string> (required)"
}
```

``` bash
# Command: "put"
# Upload a blob to the blobstore.
./bosh-azure-storage-cli -c config.json put <path/to/file> <remote-blob> 

# Command: "get"
# Fetch a blob from the blobstore.
# Destination file will be overwritten if exists.
./bosh-azure-storage-cli -c config.json get <remote-blob> <path/to/file>

# Command: "delete"
# Remove a blob from the blobstore.
./bosh-azure-storage-cli -c config.json delete <remote-blob>

# Command: "exists"
# Checks if blob exists in the blobstore.
./bosh-azure-storage-cli -c config.json exists <remote-blob>

# Command: "sign"
# Create a self-signed url for s blob in the blobstore.
./bosh-azure-storage-cli -c config.json sign <remote-blob> <get|put> <seconds-to-expiration>
```

## Running integration tests

To run the integration tests:
- Export the following variables into your environment:
  ```
  export ACCOUNT_NAME=<your Azure accounnt name>
  export ACCOUNT_KEY=<your Azure account key>
  export CONTAINER_NAME=<the target container name>
  ```
- Run `go run github.com/onsi/ginkgo/ginkgo -r  integration/`
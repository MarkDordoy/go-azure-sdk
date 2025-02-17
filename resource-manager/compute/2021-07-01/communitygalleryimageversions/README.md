
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/communitygalleryimageversions` Documentation

The `communitygalleryimageversions` SDK allows for interaction with the Azure Resource Manager Service `compute` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/communitygalleryimageversions"
```


### Client Initialization

```go
client := communitygalleryimageversions.NewCommunityGalleryImageVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CommunityGalleryImageVersionsClient.Get`

```go
ctx := context.TODO()
id := communitygalleryimageversions.NewVersionID("12345678-1234-9876-4563-123456789012", "locationValue", "sharedGalleryValue", "imageValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```

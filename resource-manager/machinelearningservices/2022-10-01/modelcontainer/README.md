
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-10-01/modelcontainer` Documentation

The `modelcontainer` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-10-01/modelcontainer"
```


### Client Initialization

```go
client := modelcontainer.NewModelContainerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ModelContainerClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := modelcontainer.NewModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "modelValue")

payload := modelcontainer.ModelContainerResource{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ModelContainerClient.Delete`

```go
ctx := context.TODO()
id := modelcontainer.NewModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "modelValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ModelContainerClient.Get`

```go
ctx := context.TODO()
id := modelcontainer.NewModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "modelValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ModelContainerClient.List`

```go
ctx := context.TODO()
id := modelcontainer.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.List(ctx, id, modelcontainer.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, modelcontainer.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```

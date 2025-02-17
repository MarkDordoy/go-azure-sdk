
## `github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/targets` Documentation

The `targets` SDK allows for interaction with the Azure Resource Manager Service `chaosstudio` (API Version `2022-07-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/targets"
```


### Client Initialization

```go
client := targets.NewTargetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TargetsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := targets.NewTargetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "parentResourceTypeValue", "parentResourceValue", "targetValue")

payload := targets.Target{
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


### Example Usage: `TargetsClient.Delete`

```go
ctx := context.TODO()
id := targets.NewTargetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "parentResourceTypeValue", "parentResourceValue", "targetValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TargetsClient.Get`

```go
ctx := context.TODO()
id := targets.NewTargetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "parentResourceTypeValue", "parentResourceValue", "targetValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TargetsClient.List`

```go
ctx := context.TODO()
id := targets.NewProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "parentResourceTypeValue", "parentResourceValue")

// alternatively `client.List(ctx, id, targets.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, targets.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


## `github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2019-08-01/agentpools` Documentation

The `agentpools` SDK allows for interaction with the Azure Resource Manager Service `containerservice` (API Version `2019-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2019-08-01/agentpools"
```


### Client Initialization

```go
client := agentpools.NewAgentPoolsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AgentPoolsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterValue", "agentPoolValue")

payload := agentpools.AgentPool{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AgentPoolsClient.Delete`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterValue", "agentPoolValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AgentPoolsClient.Get`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterValue", "agentPoolValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentPoolsClient.GetAvailableAgentPoolVersions`

```go
ctx := context.TODO()
id := agentpools.NewManagedClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterValue")

read, err := client.GetAvailableAgentPoolVersions(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentPoolsClient.GetUpgradeProfile`

```go
ctx := context.TODO()
id := agentpools.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterValue", "agentPoolValue")

read, err := client.GetUpgradeProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AgentPoolsClient.List`

```go
ctx := context.TODO()
id := agentpools.NewManagedClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```

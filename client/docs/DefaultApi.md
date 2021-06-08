# \DefaultApi

All URIs are relative to *http://127.0.0.1:8010*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionPost**](DefaultApi.md#TransactionPost) | **Post** /transaction | Issues a transaction to be executed



## TransactionPost

> TransactionResult TransactionPost(ctx).Transaction(transaction).Execute()

Issues a transaction to be executed

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transaction := *openapiclient.NewTransaction(false, "Dbname_example", "Mode_example", false, false, "Type_example") // Transaction | Optional description in *Markdown*

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TransactionPost(context.Background()).Transaction(transaction).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TransactionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionPost`: TransactionResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TransactionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction** | [**Transaction**](Transaction.md) | Optional description in *Markdown* | 

### Return type

[**TransactionResult**](TransactionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


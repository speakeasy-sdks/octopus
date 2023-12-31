# PostUnsubscribeResponse


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ContentType`                                                        | *string*                                                             | :heavy_check_mark:                                                   | HTTP response content type for this operation                        |
| `ErrorResponse`                                                      | [*shared.ErrorResponse](../../../pkg/models/shared/errorresponse.md) | :heavy_minus_sign:                                                   | Error                                                                |
| `StatusCode`                                                         | *int*                                                                | :heavy_check_mark:                                                   | HTTP response status code for this operation                         |
| `RawResponse`                                                        | [*http.Response](https://pkg.go.dev/net/http#Response)               | :heavy_check_mark:                                                   | Raw HTTP response; suitable for custom response parsing              |
| `Subscription`                                                       | [*shared.Subscription](../../../pkg/models/shared/subscription.md)   | :heavy_minus_sign:                                                   | Unsubscribed successfully                                            |
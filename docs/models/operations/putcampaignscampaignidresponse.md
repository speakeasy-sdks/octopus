# PutCampaignsCampaignIDResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Campaign`                                                    | [*shared.Campaign](../../models/shared/campaign.md)           | :heavy_minus_sign:                                            | Campaign updated successfully                                 |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `ErrorResponse`                                               | [*shared.ErrorResponse](../../models/shared/errorresponse.md) | :heavy_minus_sign:                                            | Error                                                         |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | N/A                                                           |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | N/A                                                           |
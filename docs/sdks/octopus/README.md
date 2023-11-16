# Octopus SDK


## Overview

Octopus API: API for Octopus - email marketing software.

### Available Operations

* [GetCampaignsCampaignID](#getcampaignscampaignid) - Get Campaign
* [PostCampaigns](#postcampaigns) - Create Campaign
* [PostCampaignsCampaignIDSend](#postcampaignscampaignidsend) - Send Campaign
* [PostSubscribe](#postsubscribe) - Subscribe
* [PostUnsubscribe](#postunsubscribe) - Unsubscribe
* [PutCampaignsCampaignID](#putcampaignscampaignid) - Update Campaign

## GetCampaignsCampaignID

Get Campaign

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"github.com/speakeasy-sdks/octopus"
	"context"
	"github.com/speakeasy-sdks/octopus/pkg/models/operations"
	"log"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.GetCampaignsCampaignID(ctx, operations.GetCampaignsCampaignIDRequest{
        CampaignID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Campaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCampaignsCampaignIDRequest](../../pkg/models/operations/getcampaignscampaignidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetCampaignsCampaignIDResponse](../../pkg/models/operations/getcampaignscampaignidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostCampaigns

Create Campaign

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"github.com/speakeasy-sdks/octopus"
	"context"
	"log"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PostCampaigns(ctx, shared.CampaignRequest{
        Body: octopus.String("Welcome to Octopus"),
        From: octopus.String("string"),
        Name: octopus.String("Welcome Campaign"),
        Schedule: octopus.String("2021-01-01T00:00:00Z"),
        Subject: octopus.String("Welcome to Octopus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Campaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.CampaignRequest](../../pkg/models/shared/campaignrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.PostCampaignsResponse](../../pkg/models/operations/postcampaignsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostCampaignsCampaignIDSend

Send Campaign

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"github.com/speakeasy-sdks/octopus"
	"context"
	"github.com/speakeasy-sdks/octopus/pkg/models/operations"
	"log"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PostCampaignsCampaignIDSend(ctx, operations.PostCampaignsCampaignIDSendRequest{
        CampaignSendingRequest: shared.CampaignSendingRequest{
            CampaignID: octopus.String("123456"),
        },
        CampaignID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CampaignSending != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PostCampaignsCampaignIDSendRequest](../../pkg/models/operations/postcampaignscampaignidsendrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PostCampaignsCampaignIDSendResponse](../../pkg/models/operations/postcampaignscampaignidsendresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostSubscribe

Subscribe

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"github.com/speakeasy-sdks/octopus"
	"context"
	"log"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PostSubscribe(ctx, shared.SubscriptionRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Subscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.SubscriptionRequest](../../pkg/models/shared/subscriptionrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.PostSubscribeResponse](../../pkg/models/operations/postsubscriberesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostUnsubscribe

Unsubscribe

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"github.com/speakeasy-sdks/octopus"
	"context"
	"log"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PostUnsubscribe(ctx, shared.SubscriptionRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Subscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.SubscriptionRequest](../../pkg/models/shared/subscriptionrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.PostUnsubscribeResponse](../../pkg/models/operations/postunsubscriberesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PutCampaignsCampaignID

Update Campaign

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"github.com/speakeasy-sdks/octopus"
	"context"
	"github.com/speakeasy-sdks/octopus/pkg/models/operations"
	"log"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PutCampaignsCampaignID(ctx, operations.PutCampaignsCampaignIDRequest{
        CampaignRequest: shared.CampaignRequest{
            Body: octopus.String("Welcome to Octopus"),
            From: octopus.String("string"),
            Name: octopus.String("Welcome Campaign"),
            Schedule: octopus.String("2021-01-01T00:00:00Z"),
            Subject: octopus.String("Welcome to Octopus"),
        },
        CampaignID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Campaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutCampaignsCampaignIDRequest](../../pkg/models/operations/putcampaignscampaignidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutCampaignsCampaignIDResponse](../../pkg/models/operations/putcampaignscampaignidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

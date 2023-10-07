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
	"context"
	"log"
	"github.com/speakeasy-sdks/octopus"
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"github.com/speakeasy-sdks/octopus/pkg/models/operations"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Octopus.GetCampaignsCampaignID(ctx, operations.GetCampaignsCampaignIDRequest{
        CampaignID: "Southeast firewall gray",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCampaignsCampaignIDRequest](../../models/operations/getcampaignscampaignidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetCampaignsCampaignIDResponse](../../models/operations/getcampaignscampaignidresponse.md), error**


## PostCampaigns

Create Campaign

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/octopus"
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Octopus.PostCampaigns(ctx, shared.CampaignRequest{
        Body: octopus.String("Welcome to Octopus"),
        From: octopus.String("deposit"),
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

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.CampaignRequest](../../models/shared/campaignrequest.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.PostCampaignsResponse](../../models/operations/postcampaignsresponse.md), error**


## PostCampaignsCampaignIDSend

Send Campaign

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/octopus"
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"github.com/speakeasy-sdks/octopus/pkg/models/operations"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Octopus.PostCampaignsCampaignIDSend(ctx, operations.PostCampaignsCampaignIDSendRequest{
        CampaignSendingRequest: shared.CampaignSendingRequest{
            CampaignID: octopus.String("123456"),
        },
        CampaignID: "within Oxnard",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PostCampaignsCampaignIDSendRequest](../../models/operations/postcampaignscampaignidsendrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PostCampaignsCampaignIDSendResponse](../../models/operations/postcampaignscampaignidsendresponse.md), error**


## PostSubscribe

Subscribe

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/octopus"
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Octopus.PostSubscribe(ctx, shared.SubscriptionRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Subscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.SubscriptionRequest](../../models/shared/subscriptionrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PostSubscribeResponse](../../models/operations/postsubscriberesponse.md), error**


## PostUnsubscribe

Unsubscribe

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/octopus"
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Octopus.PostUnsubscribe(ctx, shared.SubscriptionRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Subscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.SubscriptionRequest](../../models/shared/subscriptionrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PostUnsubscribeResponse](../../models/operations/postunsubscriberesponse.md), error**


## PutCampaignsCampaignID

Update Campaign

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/octopus"
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"github.com/speakeasy-sdks/octopus/pkg/models/operations"
)

func main() {
    s := octopus.New(
        octopus.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Octopus.PutCampaignsCampaignID(ctx, operations.PutCampaignsCampaignIDRequest{
        CampaignRequest: shared.CampaignRequest{
            Body: octopus.String("Welcome to Octopus"),
            From: octopus.String("Lutetium"),
            Name: octopus.String("Welcome Campaign"),
            Schedule: octopus.String("2021-01-01T00:00:00Z"),
            Subject: octopus.String("Welcome to Octopus"),
        },
        CampaignID: "warmly Hybrid",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PutCampaignsCampaignIDRequest](../../models/operations/putcampaignscampaignidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.PutCampaignsCampaignIDResponse](../../models/operations/putcampaignscampaignidresponse.md), error**


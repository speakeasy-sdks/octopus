# github.com/speakeasy-sdks/octopus

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/octopus
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/octopus"
	"github.com/speakeasy-sdks/octopus/pkg/models/operations"
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"log"
)

func main() {
	s := octopus.New(
		octopus.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Octopus.GetCampaignsCampaignID(ctx, operations.GetCampaignsCampaignIDRequest{
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [Octopus SDK](docs/sdks/octopus/README.md)

* [GetCampaignsCampaignID](docs/sdks/octopus/README.md#getcampaignscampaignid) - Get Campaign
* [PostCampaigns](docs/sdks/octopus/README.md#postcampaigns) - Create Campaign
* [PostCampaignsCampaignIDSend](docs/sdks/octopus/README.md#postcampaignscampaignidsend) - Send Campaign
* [PostSubscribe](docs/sdks/octopus/README.md#postsubscribe) - Subscribe
* [PostUnsubscribe](docs/sdks/octopus/README.md#postunsubscribe) - Unsubscribe
* [PutCampaignsCampaignID](docs/sdks/octopus/README.md#putcampaignscampaignid) - Update Campaign
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

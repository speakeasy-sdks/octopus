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
<!-- End SDK Example Usage -->
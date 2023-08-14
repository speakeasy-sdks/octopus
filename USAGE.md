<!-- Start SDK Example Usage -->


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
    res, err := s.GetCampaignsCampaignID(ctx, operations.GetCampaignsCampaignIDRequest{
        CampaignID: "corrupti",
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
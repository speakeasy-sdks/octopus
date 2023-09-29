// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/octopus/pkg/models/shared"
	"net/http"
)

type PostCampaignsCampaignIDSendRequest struct {
	CampaignSendingRequest shared.CampaignSendingRequest `request:"mediaType=application/json"`
	CampaignID             string                        `pathParam:"style=simple,explode=false,name=campaignId"`
}

func (o *PostCampaignsCampaignIDSendRequest) GetCampaignSendingRequest() shared.CampaignSendingRequest {
	if o == nil {
		return shared.CampaignSendingRequest{}
	}
	return o.CampaignSendingRequest
}

func (o *PostCampaignsCampaignIDSendRequest) GetCampaignID() string {
	if o == nil {
		return ""
	}
	return o.CampaignID
}

type PostCampaignsCampaignIDSendResponse struct {
	// Campaign sent successfully
	CampaignSending *shared.CampaignSending
	// HTTP response content type for this operation
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostCampaignsCampaignIDSendResponse) GetCampaignSending() *shared.CampaignSending {
	if o == nil {
		return nil
	}
	return o.CampaignSending
}

func (o *PostCampaignsCampaignIDSendResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostCampaignsCampaignIDSendResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *PostCampaignsCampaignIDSendResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostCampaignsCampaignIDSendResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

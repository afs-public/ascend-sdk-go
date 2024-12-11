// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type SnapshotsListSnapshotsRequest struct {
	// A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:
	//  `snapshot_id`
	//  `process_date`
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
	// The number of snapshots to be returned per page. Defaults to 500. Maximum is 1000.
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// The token used to retrieve a page of snapshots.
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *SnapshotsListSnapshotsRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *SnapshotsListSnapshotsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *SnapshotsListSnapshotsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type SnapshotsListSnapshotsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListSnapshotsResponse *components.ListSnapshotsResponse
	// PERMISSION_DENIED: The user does not have access to the requested resource.
	Status *components.Status
}

func (o *SnapshotsListSnapshotsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SnapshotsListSnapshotsResponse) GetListSnapshotsResponse() *components.ListSnapshotsResponse {
	if o == nil {
		return nil
	}
	return o.ListSnapshotsResponse
}

func (o *SnapshotsListSnapshotsResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type WatchlistServiceGetWatchlistItemRequest struct {
	// The watchlist id.
	WatchlistID string `pathParam:"style=simple,explode=false,name=watchlist_id"`
	// The item id.
	ItemID string `pathParam:"style=simple,explode=false,name=item_id"`
}

func (o *WatchlistServiceGetWatchlistItemRequest) GetWatchlistID() string {
	if o == nil {
		return ""
	}
	return o.WatchlistID
}

func (o *WatchlistServiceGetWatchlistItemRequest) GetItemID() string {
	if o == nil {
		return ""
	}
	return o.ItemID
}

type WatchlistServiceGetWatchlistItemResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	WatchlistItem *components.WatchlistItem
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *WatchlistServiceGetWatchlistItemResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *WatchlistServiceGetWatchlistItemResponse) GetWatchlistItem() *components.WatchlistItem {
	if o == nil {
		return nil
	}
	return o.WatchlistItem
}

func (o *WatchlistServiceGetWatchlistItemResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
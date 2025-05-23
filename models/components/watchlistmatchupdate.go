// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// WatchlistMatchUpdateMatchState - Match state - whether or not the match is confirmed
type WatchlistMatchUpdateMatchState string

const (
	WatchlistMatchUpdateMatchStateMatchUnspecified WatchlistMatchUpdateMatchState = "MATCH_UNSPECIFIED"
	WatchlistMatchUpdateMatchStateConfirmedMatch   WatchlistMatchUpdateMatchState = "CONFIRMED_MATCH"
	WatchlistMatchUpdateMatchStatePotentialMatch   WatchlistMatchUpdateMatchState = "POTENTIAL_MATCH"
	WatchlistMatchUpdateMatchStateNoMatch          WatchlistMatchUpdateMatchState = "NO_MATCH"
	WatchlistMatchUpdateMatchStateInconclusive     WatchlistMatchUpdateMatchState = "INCONCLUSIVE"
)

func (e WatchlistMatchUpdateMatchState) ToPointer() *WatchlistMatchUpdateMatchState {
	return &e
}

// WatchlistMatchUpdate - Matched profile details
type WatchlistMatchUpdate struct {
	// Identifies that a confirmed watchlist match can be excluded when calculating the related screen state
	ExcludeFromScreening *bool `json:"exclude_from_screening,omitempty"`
	// Match state - whether or not the match is confirmed
	MatchState *WatchlistMatchUpdateMatchState `json:"match_state,omitempty"`
	// Indicates the watchlist source for a given match
	WatchlistID *string `json:"watchlist_id,omitempty"`
	// Identification number for the watchlist item that was matched
	WatchlistItemID *int `json:"watchlist_item_id,omitempty"`
}

func (o *WatchlistMatchUpdate) GetExcludeFromScreening() *bool {
	if o == nil {
		return nil
	}
	return o.ExcludeFromScreening
}

func (o *WatchlistMatchUpdate) GetMatchState() *WatchlistMatchUpdateMatchState {
	if o == nil {
		return nil
	}
	return o.MatchState
}

func (o *WatchlistMatchUpdate) GetWatchlistID() *string {
	if o == nil {
		return nil
	}
	return o.WatchlistID
}

func (o *WatchlistMatchUpdate) GetWatchlistItemID() *int {
	if o == nil {
		return nil
	}
	return o.WatchlistItemID
}

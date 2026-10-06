package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPagination(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	// Make the initial request to list assets (first page)
	res, err := sdk.Assets.ListAssets(ctx, ascendsdkgo.String(""), ascendsdkgo.Int(25), ascendsdkgo.String(""), ascendsdkgo.String(""))
	require.NoError(t, err, "Initial assets request should not error")

	// Verify the response structure is valid
	assert.NotNil(t, res.HTTPMeta, "HTTP metadata should be present in response")
	assert.NotNil(t, res.HTTPMeta.Response, "HTTP response object should be present")
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode, "Should receive successful HTTP 200 response")
	assert.NotNil(t, res.ListAssetsResponse, "Assets response data should be present")

	// Test pagination by only checking the first few pages
	pageCount := 0
	maxPages := 3 // Limit to the first 3 pages for testing to avoid long-running tests

	// Iterate through pages until we reach the limit or no more pages exist
	for res != nil && pageCount < maxPages {
		fmt.Printf("Processing page %d\n", pageCount+1)

		// Verify current page has valid response structure
		assert.NotNil(t, res.ListAssetsResponse, "Page %d should have valid assets response", pageCount+1)

		// Attempt to get the next page using the SDK's pagination method
		nextRes, err := res.Next()
		require.NoError(t, err, "Next page request should not error")

		// If we successfully got the next page, verify it's also valid
		if nextRes != nil {
			assert.Equal(t, 200, nextRes.HTTPMeta.Response.StatusCode, "Next page should also return HTTP 200")
		}

		// Move to the next page for the next iteration
		res = nextRes
		pageCount++
	}

	fmt.Println("Pagination test completed successfully.")
}

func TestEmptyInitialResponseHandling(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	// Use a filter that should return no results to test empty response handling
	// This filter looks for assets with an impossible condition
	emptyFilter := `name == "NONEXISTENT_ASSET_12345_IMPOSSIBLE"`

	res, err := sdk.Assets.ListAssets(ctx, ascendsdkgo.String(""), ascendsdkgo.Int(25), ascendsdkgo.String(""), ascendsdkgo.String(emptyFilter))
	require.NoError(t, err, "Empty filter request should not error")

	// Verify the response structure is still valid
	assert.NotNil(t, res.HTTPMeta, "HTTP metadata should be present")
	assert.NotNil(t, res.HTTPMeta.Response, "HTTP response object should be present")
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode, "Should receive successful HTTP 200 response")
	assert.NotNil(t, res.ListAssetsResponse, "Assets response should be present")

	// Verify that Next() returns nil for empty responses
	for res != nil {
		// Check if the current response has assets
		if res.ListAssetsResponse.Assets != nil && len(res.ListAssetsResponse.Assets) > 0 {
			assert.Fail(t, "Expected no assets in response, but found some")
		}

		// Attempt to get the next page
		nextRes, err := res.Next()
		require.NoError(t, err, "Next page request should not error even for empty results")
		res = nextRes
	}

	fmt.Println("Empty initial response handling test completed successfully.")
}

func TestPaginationWithSymbolFilters(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	// Test filter by symbol
	equityFilter := `symbol == "WTG"`
	res, err := sdk.Assets.ListAssets(ctx, ascendsdkgo.String(""), ascendsdkgo.Int(25), ascendsdkgo.String(""), ascendsdkgo.String(equityFilter))
	require.NoError(t, err, "Symbol filter request should not error")

	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode, "Should receive successful HTTP 200 response")
	assert.NotNil(t, res.ListAssetsResponse, "Assets response should be present")

	// Count pages with an equity filter
	equityPages := 0
	equityAssets := 0
	maxPages := 3

	for res != nil && equityPages < maxPages {
		equityPages++

		if res.ListAssetsResponse.Assets != nil && len(res.ListAssetsResponse.Assets) > 0 {
			currentCount := len(res.ListAssetsResponse.Assets)
			equityAssets += currentCount
			fmt.Printf("Symbol filter - Page %d: %d assets\n", equityPages, currentCount)

			// Verify all returned assets match the filter (if any)
			for _, asset := range res.ListAssetsResponse.Assets {
				if asset.Symbol != nil {
					assert.Equal(t, "WTG", *asset.Symbol, "Symbol %s doesn't match filter", *asset.Symbol)
				}
			}
		}

		nextRes, err := res.Next()
		require.NoError(t, err, "Next page request should not error")
		res = nextRes
	}

	fmt.Printf("Filter test completed. Symbol assets found: %d across %d pages\n", equityAssets, equityPages)
}

func TestPaginationWithUsableFilter(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	// Test with a usable filter
	activeFilter := `usable == true`
	res, err := sdk.Assets.ListAssets(ctx, ascendsdkgo.String(""), ascendsdkgo.Int(25), ascendsdkgo.String(""), ascendsdkgo.String(activeFilter))
	require.NoError(t, err, "Usable filter request should not error")

	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode, "Should receive successful HTTP 200 response")
	assert.NotNil(t, res.ListAssetsResponse, "Assets response should be present")

	pageCount := 0
	totalActiveAssets := 0
	maxPages := 3

	for res != nil && pageCount < maxPages {
		pageCount++

		if res.ListAssetsResponse.Assets != nil && len(res.ListAssetsResponse.Assets) > 0 {
			pageAssets := len(res.ListAssetsResponse.Assets)
			totalActiveAssets += pageAssets
			fmt.Printf("Usable filter - Page %d: %d assets\n", pageCount, pageAssets)

			// Verify all returned assets match the filter
			for _, asset := range res.ListAssetsResponse.Assets {
				if asset.Usable != nil {
					assert.True(t, *asset.Usable, "Asset usable status %v doesn't match filter", *asset.Usable)
				}
			}
		}

		nextRes, err := res.Next()
		require.NoError(t, err, "Next page request should not error")
		res = nextRes
	}

	fmt.Printf("Status filter test completed. Active assets: %d across %d pages\n", totalActiveAssets, pageCount)
}

func TestComplexFilterPagination(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()
	// Test with a complex filter combining multiple conditions
	complexFilter := `type == "EQUITY" && usable == true`
	res, err := sdk.Assets.ListAssets(ctx, ascendsdkgo.String(""), ascendsdkgo.Int(25), ascendsdkgo.String(""), ascendsdkgo.String(complexFilter))
	require.NoError(t, err, "Complex filter request should not error")

	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode, "Should receive successful HTTP 200 response")
	assert.NotNil(t, res.ListAssetsResponse, "Assets response should be present")

	pageCount := 0
	matchingAssets := 0
	maxPages := 3

	for res != nil && pageCount < maxPages {
		pageCount++

		if res.ListAssetsResponse.Assets != nil && len(res.ListAssetsResponse.Assets) > 0 {
			pageAssets := len(res.ListAssetsResponse.Assets)
			matchingAssets += pageAssets
			fmt.Printf("Complex filter - Page %d: %d assets\n", pageCount, pageAssets)

			// Verify assets match both conditions
			for _, asset := range res.ListAssetsResponse.Assets {
				if asset.Type != nil {
					assert.Equal(t, components.AssetType1Equity, *asset.Type, "Asset type %v doesn't match filter", *asset.Type)
				}
				if asset.Usable != nil {
					assert.True(t, *asset.Usable, "Asset usable status %v doesn't match filter", *asset.Usable)
				}
			}
		}

		nextRes, err := res.Next()
		require.NoError(t, err, "Next page request should not error")
		res = nextRes
	}

	fmt.Printf("Complex filter test completed. Matching assets: %d across %d pages\n", matchingAssets, pageCount)
}

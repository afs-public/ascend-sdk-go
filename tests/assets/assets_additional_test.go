package assets

import (
	"context"
	"os"
	"testing"
	"time"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestAssetsAdditional(t *testing.T) {
	ctx := context.Background()
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	time.Sleep(5 * time.Second)

	t.Run("Test Assets Assets List Assets By_correspondent1 Assets List Assets By_correspondent1", func(t *testing.T) {
		result, err := sdk.Assets.ListAssetsCorrespondent(ctx, os.Getenv("CORRESPONDENT_ID"), ascendsdk.Pointer(10), ascendsdk.Pointer(""), ascendsdk.String(""))
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Assets Assets Get Asset By_correspondent Assets Get Asset By_correspondent1", func(t *testing.T) {
		result, err := sdk.Assets.GetAssetCorrespondent(ctx, os.Getenv("CORRESPONDENT_ID"), "8395")
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})
}

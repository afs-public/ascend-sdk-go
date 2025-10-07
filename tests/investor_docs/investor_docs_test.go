package investor_docs

import (
	"context"
	"fmt"
	"os"
	"testing"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvestorDocs(t *testing.T) {
	ctx := context.Background()
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	accountId, err := helpers.CreateAccountId(sdk, ctx)
	require.NoError(t, err)

	t.Run("Investor Docs Investor Docs Batch Create Upload Links Batch Create Upload Links1", func(t *testing.T) {
		request := components.BatchCreateUploadLinksRequestCreate{
			CreateDocumentUploadLinkRequest: []components.CreateUploadLinkRequestCreate{
				components.CreateUploadLinkRequestCreate{
					AccountDocumentUploadRequest: &components.AccountDocumentUploadRequestCreate{
						AccountID:       *accountId,
						CorrespondentID: os.Getenv("CORRESPONDENT_ID"),
						DocumentType:    components.DocumentTypeFdicSweepProgramAgreement,
					},
					ClientBatchSourceID: "cda89bd0-a6bc-4acc-89da-d35bde30cbf4",
					MimeType:            "image/jpeg",
				},
			},
		}

		result, err := sdk.InvestorDocs.BatchCreateUploadLinks(ctx, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Investor Docs Investor Docs List Documents List Documents1", func(t *testing.T) {
		filter := fmt.Sprintf(`correspondent_id=="%s"`, os.Getenv("CORRESPONDENT_ID"))
		result, err := sdk.InvestorDocs.ListDocuments(ctx, ascendsdk.Pointer(50), ascendsdk.Pointer(""), &filter)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})
}

package accounts

import (
	"context"
	"fmt"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
)

func EnrollAccountDeactivateId(sdk *ascendsdk.SDK, ctx context.Context, accountId string) (*string, error) {
	enrollAccountRequest := components.EnrollAccountRequestCreate{
		Enrollment: components.EnrollmentCreate{
			PrincipalApproverID: helpers.EnrollmentPrincipalApproverID,
			Type:                components.EnrollmentCreateTypeCashFdicCashSweep,
		},
	}

	res, err := sdk.EnrollmentsAndAgreements.EnrollAccount(ctx, accountId, enrollAccountRequest)
	if err != nil {
		return nil, err
	}

	if len(res.EnrollAccountResponse.Agreements) < 1 {
		return nil, fmt.Errorf("insufficient agreements returned")
	}

	return res.EnrollAccountResponse.Agreements[0].EnrollmentID, nil
}

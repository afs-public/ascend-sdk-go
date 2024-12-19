package accounts

import (
	"context"
	"fmt"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
)

func GetEnrollmentToDeactivate(sdk *ascendsdk.SDK, ctx context.Context, accountId string) (*string, error) {
	res, err := sdk.EnrollmentsAndAgreements.ListEnrollments(ctx, accountId, nil, nil)

	if err != nil {
		return nil, err
	}

	enrollments := res.ListEnrollmentsResponse.Enrollments

	if enrollments != nil {
		for _, enrollment := range enrollments {
			if *enrollment.Type == "DIVIDEND_REINVESTMENT_PLAN" {
				return enrollment.EnrollmentID, nil
			}
		}
	}

	return nil, fmt.Errorf("no enrollment found")
}

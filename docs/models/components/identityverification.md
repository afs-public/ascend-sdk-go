# IdentityVerification

The screen state of one screening within an investigation, one of:
- `SCREEN_STATE_UNSPECIFIED` - Default/Null value.
- `PENDING` - Screen result is pending.
- `PASSED` - Screen result has passed.
- `FAILED` - Screen result has failed.
- `NEEDS_REVIEW` - Screen result needs manual review.
- `DEFERRED_REVIEW` - Screen result is deferred for review at a later date.
- `OUT_OF_SCOPE` - Screen state is out of scope for this investigation type.


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `IdentityVerificationScreenStateUnspecified` | SCREEN_STATE_UNSPECIFIED                     |
| `IdentityVerificationPending`                | PENDING                                      |
| `IdentityVerificationPassed`                 | PASSED                                       |
| `IdentityVerificationFailed`                 | FAILED                                       |
| `IdentityVerificationNeedsReview`            | NEEDS_REVIEW                                 |
| `IdentityVerificationDeferredReview`         | DEFERRED_REVIEW                              |
| `IdentityVerificationOutOfScope`             | OUT_OF_SCOPE                                 |
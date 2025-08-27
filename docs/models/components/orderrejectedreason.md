# OrderRejectedReason

When an order has the REJECTED status, this will be populated with a system code describing the rejection.


## Values

| Name                                                                   | Value                                                                  |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `OrderRejectedReasonOrderRejectReasonUnspecified`                      | ORDER_REJECT_REASON_UNSPECIFIED                                        |
| `OrderRejectedReasonBrokerOption`                                      | BROKER_OPTION                                                          |
| `OrderRejectedReasonUnknownSecurity`                                   | UNKNOWN_SECURITY                                                       |
| `OrderRejectedReasonExchangeClosed`                                    | EXCHANGE_CLOSED                                                        |
| `OrderRejectedReasonOrderExceedsLimit`                                 | ORDER_EXCEEDS_LIMIT                                                    |
| `OrderRejectedReasonTooLateToEnter`                                    | TOO_LATE_TO_ENTER                                                      |
| `OrderRejectedReasonUnknownOrder`                                      | UNKNOWN_ORDER                                                          |
| `OrderRejectedReasonDuplicateOrder`                                    | DUPLICATE_ORDER                                                        |
| `OrderRejectedReasonStaleOrder`                                        | STALE_ORDER                                                            |
| `OrderRejectedReasonBelowNotionalMinimum`                              | BELOW_NOTIONAL_MINIMUM                                                 |
| `OrderRejectedReasonOrderDateUnavailable`                              | ORDER_DATE_UNAVAILABLE                                                 |
| `OrderRejectedReasonAggressiveLimitPrice`                              | AGGRESSIVE_LIMIT_PRICE                                                 |
| `OrderRejectedReasonAccountNotEntitled`                                | ACCOUNT_NOT_ENTITLED                                                   |
| `OrderRejectedReasonSystemError`                                       | SYSTEM_ERROR                                                           |
| `OrderRejectedReasonBlockingCorporateAction`                           | BLOCKING_CORPORATE_ACTION                                              |
| `OrderRejectedReasonUnavailablePriceQuote`                             | UNAVAILABLE_PRICE_QUOTE                                                |
| `OrderRejectedReasonExecutionMisconfiguredClient`                      | EXECUTION_MISCONFIGURED_CLIENT                                         |
| `OrderRejectedReasonNotionalQuantityNotAllowedForSecurity`             | NOTIONAL_QUANTITY_NOT_ALLOWED_FOR_SECURITY                             |
| `OrderRejectedReasonFractionalQuantityNotAllowedForSecurity`           | FRACTIONAL_QUANTITY_NOT_ALLOWED_FOR_SECURITY                           |
| `OrderRejectedReasonOnlyFractionalSellOrWholeOrdersAllowedForSecurity` | ONLY_FRACTIONAL_SELL_OR_WHOLE_ORDERS_ALLOWED_FOR_SECURITY              |
| `OrderRejectedReasonSymbolNotTradeable`                                | SYMBOL_NOT_TRADEABLE                                                   |
| `OrderRejectedReasonAboveNotionalMaximum`                              | ABOVE_NOTIONAL_MAXIMUM                                                 |
| `OrderRejectedReasonAboveShareMaximum`                                 | ABOVE_SHARE_MAXIMUM                                                    |
| `OrderRejectedReasonFailedBuyingPower`                                 | FAILED_BUYING_POWER                                                    |
| `OrderRejectedReasonInsufficientPosition`                              | INSUFFICIENT_POSITION                                                  |
| `OrderRejectedReasonMaxSellQuantityRequired`                           | MAX_SELL_QUANTITY_REQUIRED                                             |
| `OrderRejectedReasonMaxSellQuantityProhibited`                         | MAX_SELL_QUANTITY_PROHIBITED                                           |
| `OrderRejectedReasonStopPriceExceedsMarketPrice`                       | STOP_PRICE_EXCEEDS_MARKET_PRICE                                        |
| `OrderRejectedReasonTradesDisabledForAssetType`                        | TRADES_DISABLED_FOR_ASSET_TYPE                                         |
| `OrderRejectedReasonCommissionNotAllowedForNonBrokerDealer`            | COMMISSION_NOT_ALLOWED_FOR_NON_BROKER_DEALER                           |
| `OrderRejectedReasonAssetNotSetUpToTrade`                              | ASSET_NOT_SET_UP_TO_TRADE                                              |
| `OrderRejectedReasonInvalidOrderQuantity`                              | INVALID_ORDER_QUANTITY                                                 |
| `OrderRejectedReasonClientReceivedTimeRequired`                        | CLIENT_RECEIVED_TIME_REQUIRED                                          |
| `OrderRejectedReasonClientNotPermittedToUseTradingSession`             | CLIENT_NOT_PERMITTED_TO_USE_TRADING_SESSION                            |
| `OrderRejectedReasonStopPriceBelowMarketPrice`                         | STOP_PRICE_BELOW_MARKET_PRICE                                          |
# ExemptCustomerReason

**Field Dependencies:**

Exempt entities must set `exempt_verifying_beneficial_owners` to `true` and provide an `exempt_customer_reason` on the owner record.

Required if `exempt_verifying_beneficial_owners` is `true`.

Otherwise, must be empty.


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ExemptCustomerReasonExemptReasonUnspecified`                        | EXEMPT_REASON_UNSPECIFIED                                            |
| `ExemptCustomerReasonRegulatedFinancialInstitution`                  | REGULATED_FINANCIAL_INSTITUTION                                      |
| `ExemptCustomerReasonDepartmentOrAgencyOfFederalStateOrSubdivision`  | DEPARTMENT_OR_AGENCY_OF_FEDERAL_STATE_OR_SUBDIVISION                 |
| `ExemptCustomerReasonNonBankListedEntity`                            | NON_BANK_LISTED_ENTITY                                               |
| `ExemptCustomerReasonSection12SecuritiesExchangeAct1934Or15D`        | SECTION_12_SECURITIES_EXCHANGE_ACT_1934_OR_15D                       |
| `ExemptCustomerReasonSection3InvestmentCompanyAct1940`               | SECTION_3_INVESTMENT_COMPANY_ACT_1940                                |
| `ExemptCustomerReasonSection202AInvestmentAdvisorsAct1940`           | SECTION_202A_INVESTMENT_ADVISORS_ACT_1940                            |
| `ExemptCustomerReasonSection3SecuritiesExchangeAct1934Section6Or17A` | SECTION_3_SECURITIES_EXCHANGE_ACT_1934_SECTION_6_OR_17A              |
| `ExemptCustomerReasonAnyOtherSecuritiesExchangeAct1934`              | ANY_OTHER_SECURITIES_EXCHANGE_ACT_1934                               |
| `ExemptCustomerReasonCommodityFuturesTradingCommissionRegistered`    | COMMODITY_FUTURES_TRADING_COMMISSION_REGISTERED                      |
| `ExemptCustomerReasonPublicAccountingFirmSection102SarbanesOxley`    | PUBLIC_ACCOUNTING_FIRM_SECTION_102_SARBANES_OXLEY                    |
| `ExemptCustomerReasonStateRegulatedInsuranceCompany`                 | STATE_REGULATED_INSURANCE_COMPANY                                    |
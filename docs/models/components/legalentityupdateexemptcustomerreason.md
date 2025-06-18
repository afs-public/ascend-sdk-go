# LegalEntityUpdateExemptCustomerReason

**Field Dependencies:**

Exempt entities must set `exempt_verifying_beneficial_owners` to `true` and provide an `exempt_customer_reason` on the owner record.

Required if `exempt_verifying_beneficial_owners` is `true`.

Otherwise, must be empty.


## Values

| Name                                                                                  | Value                                                                                 |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `LegalEntityUpdateExemptCustomerReasonExemptReasonUnspecified`                        | EXEMPT_REASON_UNSPECIFIED                                                             |
| `LegalEntityUpdateExemptCustomerReasonRegulatedFinancialInstitution`                  | REGULATED_FINANCIAL_INSTITUTION                                                       |
| `LegalEntityUpdateExemptCustomerReasonDepartmentOrAgencyOfFederalStateOrSubdivision`  | DEPARTMENT_OR_AGENCY_OF_FEDERAL_STATE_OR_SUBDIVISION                                  |
| `LegalEntityUpdateExemptCustomerReasonNonBankListedEntity`                            | NON_BANK_LISTED_ENTITY                                                                |
| `LegalEntityUpdateExemptCustomerReasonSection12SecuritiesExchangeAct1934Or15D`        | SECTION_12_SECURITIES_EXCHANGE_ACT_1934_OR_15D                                        |
| `LegalEntityUpdateExemptCustomerReasonSection3InvestmentCompanyAct1940`               | SECTION_3_INVESTMENT_COMPANY_ACT_1940                                                 |
| `LegalEntityUpdateExemptCustomerReasonSection202AInvestmentAdvisorsAct1940`           | SECTION_202A_INVESTMENT_ADVISORS_ACT_1940                                             |
| `LegalEntityUpdateExemptCustomerReasonSection3SecuritiesExchangeAct1934Section6Or17A` | SECTION_3_SECURITIES_EXCHANGE_ACT_1934_SECTION_6_OR_17A                               |
| `LegalEntityUpdateExemptCustomerReasonAnyOtherSecuritiesExchangeAct1934`              | ANY_OTHER_SECURITIES_EXCHANGE_ACT_1934                                                |
| `LegalEntityUpdateExemptCustomerReasonCommodityFuturesTradingCommissionRegistered`    | COMMODITY_FUTURES_TRADING_COMMISSION_REGISTERED                                       |
| `LegalEntityUpdateExemptCustomerReasonPublicAccountingFirmSection102SarbanesOxley`    | PUBLIC_ACCOUNTING_FIRM_SECTION_102_SARBANES_OXLEY                                     |
| `LegalEntityUpdateExemptCustomerReasonStateRegulatedInsuranceCompany`                 | STATE_REGULATED_INSURANCE_COMPANY                                                     |
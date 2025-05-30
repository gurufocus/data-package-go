# FundamentalsNNOREITNODIRECTCashflowStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetImpairmentCharge** | Pointer to **float32** |  | [optional] 
**BeginningCashPosition** | Pointer to **float32** |  | [optional] 
**CashFlowCapitalExpenditure** | Pointer to **float32** | &lt;p&gt;{{Cash_Flow_CPEX}} refers to the funds spent for a company to acquire or upgrade physical assets such as property, industrial buildings or equipment.&lt;/p&gt; | [optional] 
**CashFlowDeferredTax** | Pointer to **float32** | &lt;p&gt;{{CF_DeferredTax}} represents future tax liability or asset, resulting from temporary differences between book (accounting) value of assets and liabilities, and their tax value. This arises due to differences between financial accounting for shareholders and tax accounting.&lt;/p&gt; | [optional] 
**CashFlowDepreciationDepletionAmortization** | Pointer to **float32** | &lt;p&gt;{{DDA}} is a present expense that accounts for the past cost of an asset that is now providing benefits. Depletion and amortization are synonyms for depreciation. Generally: The term depreciation is used when discussing man made tangible assets. The term depletion is used when discussing natural tangible assets. The term amortization is used when discussing intangible assets&lt;/p&gt; | [optional] 
**CashFlowForLeaseFinancing** | Pointer to **float32** | https://www.gurufocus.com/glossary/cash_flow_for_lease_financing | [optional] 
**CashFlowFromInvesting** | Pointer to **float32** | &lt;p&gt;{{Cash_Flow_from_Investing}} covers the cash a company gains or spends from investment activities in financial market and operating subsidiaries. It also includes the cash the company used for {{Net_PPE}}(PPE). If a company spends cash on {{Net_PPE}} (PPE), this will reduce their cash position. This is called {{Cash_Flow_CPEX}} (CPEX). Likewise, if a company buys another company for cash, this will reduce their cash position. &lt;br&gt;{{Cash_Flow_from_Investing}} is calculated as {{Cash_Flow_from_Investing}} &#x3D; {{PurchaseOfPPE}} + {{SaleOfPPE}} + {{PurchaseOfBusiness}} + {{SaleOfBusiness}} + {{PurchaseOfInvestment}} + {{SaleOfInvestment}} + {{NetIntangiblesPurchaseAndSale}} + {{CashFromDiscontinuedInvestingActivities}} + {{CashFromOtherInvestingActivities}}&lt;/p&gt; | [optional] 
**CashFlowFromOperations** | Pointer to **float32** | &lt;p&gt;{{Cash_Flow_from_Operations}} refers to the cash brought in through a company&#39;s sales. &lt;br&gt;Therefore, {{Cash_Flow_from_Operations}} &#x3D; {{NetIncomeFromContinuingOperations}} + {{CF_DDA}} + {ChangeInWorkingCapital}} + Deferred Tax + {{Cash_Flow_from_Disc_Op}} + {{AssetImpairmentCharge}} + {{StockBasedCompensation}} + {{Cash_Flow_from_Others}}&lt;/p&gt; | [optional] 
**CashFlowFromOthers** | Pointer to **float32** | &lt;p&gt;{{Cash_Flow_from_Others}} may include {{ChangeInWorkingCapital}}. These are cash differences caused by the {{ChangeInInventory}}, {{AccountsPayable}}, {{Accts_Rec}} etc. For instance, if a company pays its suppliers slower, its cash position will build up faster. If a company receives payments from its customers slower, its {{Accts_Rec}} will rise, and its cash position will grow more slowly (or even shrink).&lt;/p&gt; | [optional] 
**CashFromDiscontinuedInvestingActivities** | Pointer to **float32** | &lt;p&gt;{{CashFromDiscontinuedInvestingActivities}} means the cash received by a company that comes from the discontinued investing activities.&lt;/p&gt; | [optional] 
**CashFromDiscontinuedOperatingActivities** | Pointer to **float32** | The cash generated from discontinued operations | [optional] 
**CashFromFinancing** | Pointer to **float32** | &lt;p&gt;{{Cash_from_Financing}} is the cash generated/spent from financial activities such as share issuance (buy back), debt issuance (repayment), and dividends paid to preferred and common stockholders. In the calculation of {{total_freecashflow}}, {{Cash_from_Financing}} is not calculated because it is not related to operating activities. &lt;br&gt;{{Cash_from_Financing}} &#x3D; {{Issuance_of_Stock}} + {{Repurchase_of_Stock}} + {{Net_Issuance_of_Debt}} + {{Net_Issuance_of_preferred}} + {{Dividends}} + Other Financing&lt;/p&gt; | [optional] 
**CashFromOtherInvestingActivities** | Pointer to **float32** | &lt;p&gt;{{CashFromOtherInvestingActivities}} means the cash received by a company that comes from other investing activities.&lt;/p&gt; | [optional] 
**ChangeInInventory** | Pointer to **float32** | &lt;p&gt;{{ChangeInInventory}} is the difference between last period&#39;s ending inventory and the current period&#39;s ending inventory.&lt;/p&gt; | [optional] 
**ChangeInOtherWorkingCapital** | Pointer to **float32** |  | [optional] 
**ChangeInPayablesAndAccruedExpense** | Pointer to **float32** | &lt;p&gt;{{ChangeInPayablesAndAccruedExpense}} is the increase or decrease between periods of the {{Accts_Payable}}. Accrued expenses represent expenses incurred at the end of the reporting period but not yet paid; also called accrued liabilities. The accrued liability is shown under Liabilities section in the balance sheet.&lt;/p&gt; | [optional] 
**ChangeInPrepaidAssets** | Pointer to **float32** | &lt;p&gt;{{ChangeInPrepaidAssets}} is any increase or decrease between periods of the prepaid assets.&lt;/p&gt; | [optional] 
**ChangeInReceivables** | Pointer to **float32** | &lt;p&gt;Change In {{Accts_Rec}} relative to the previous period. It is any increase or decrease in the cash a company is owed by its customers.&lt;/p&gt; | [optional] 
**ChangeInWorkingCapital** | Pointer to **float32** | &lt;p&gt;Working Capital is a measure of a company&#39;s short term liquidity or its ability to cover short term liabilities. It is defined as the difference between a company&#39;s {{Total_Current_Assets}} and {{Total_Current_Liabilities}}. &lt;br&gt;Working Capital is calculated as: Working Capital &#x3D; {{Total_Current_Assets}} - {{Total_Current_Liabilities}} &lt;br&gt;{{ChangeInWorkingCapital}} is reported in the cash flow statement since it is one of the major ways in which {{Net_Income}} can differ from operating cash flow.&lt;/p&gt; | [optional] 
**DebtIssuance** | Pointer to **float32** |  | [optional] 
**DebtPayments** | Pointer to **float32** |  | [optional] 
**Dividends** | Pointer to **float32** | &lt;p&gt;{{Dividends}} refers to the payment of cash to shareholders as dividends when the company generates income.&lt;/p&gt; | [optional] 
**EffectOfExchangeRateChanges** | Pointer to **float32** |  | [optional] 
**EndingCashPosition** | Pointer to **float32** |  | [optional] 
**IssuanceOfStock** | Pointer to **float32** | &lt;p&gt;A company may raise cash from issuing new shares. It can also use cash to buy back shares. If this number is positive, it means that the company has received more cash from issuing shares than it has paid to buy back shares. If this number is negative, it means that company has paid more cash to buy back shares than it has received for issuing shares.&lt;/p&gt; | [optional] 
**NetChangeInCash** | Pointer to **float32** | &lt;p&gt;{{Net_Change_in_Cash}} is calculated as {{Net_Change_in_Cash}} &#x3D; {{Cash_Flow_from_Operations}} + {{Cash_Flow_from_Investing}} + {{Cash_from_Financing}} + {{effect_of_exchange_rate_changes}}&lt;/p&gt; | [optional] 
**NetIncomeFromContinuingOperations** | Pointer to **float32** | &lt;p&gt;{{NetIncomeFromContinuingOperations}} indicates the {{Net_Income}} that a firm brings in from ongoing business activities. These activities are expected to continue into the next reporting period. It excludes extraordinary items, income from the cumulative effects of accounting changes, non-recurring items, income from tax loss carry forward, and {{IS_preferred_dividends}}.&lt;/p&gt; | [optional] 
**NetIntangiblesPurchaseAndSale** | Pointer to **float32** | &lt;p&gt;{{NetIntangiblesPurchaseAndSale}} means the net cash inflow received by a company that comes from the purchase and sale of intangibles. It equals the cash received from sale of intangibles minus the cash spent on purchasing intangibles.&lt;/p&gt; | [optional] 
**NetIssuanceOfDebt** | Pointer to **float32** | &lt;p&gt;{{Net_Issuance_of_Debt}} is the cash a company received or spent through debt related activities such as debt issuance or debt repayment. If a company pays down its debt during the period, this number will be negative. If a company issued more debt, it receives cash and this number is positive.&lt;/p&gt; | [optional] 
**NetIssuanceOfPreferred** | Pointer to **float32** | &lt;p&gt;A company may raise cash from issuing new preferred shares. It can also use cash to buy back preferred shares. If this number is positive, it means that the company has received more cash from issuing preferred shares than it has paid to buy back preferred shares. If this number is negative, it means that company has paid more cash to buy back preferred shares than it has received for issuing preferred shares.&lt;/p&gt; | [optional] 
**OtherFinancing** | Pointer to **float32** | &lt;p&gt;{{Other_Financing}} represents other {{Cash_from_Financing}} activity that not otherwise classified, which includes: Proceeds From Stock Option Exercised, Other Financing Charges.&lt;/p&gt; | [optional] 
**PurchaseOfBusiness** | Pointer to **float32** | &lt;p&gt;{{PurchaseOfBusiness}} is the amount used to purchase business.&lt;/p&gt; | [optional] 
**PurchaseOfInvestment** | Pointer to **float32** | &lt;p&gt;{{PurchaseOfInvestment}} represents cash outflow on the purchase of investments in securities.&lt;/p&gt; | [optional] 
**PurchaseOfPpe** | Pointer to **float32** | &lt;p&gt;{{PurchaseOfPPE}} is the amount used to purchase Property, Plant and Equipment.&lt;/p&gt; | [optional] 
**RepurchaseOfStock** | Pointer to **float32** | &lt;p&gt;A company may raise cash from issuing new shares. It can also use cash to buy back shares. {{Repurchase_of_Stock}} represents the cash outflow to reacquire common stock during the period.&lt;/p&gt; | [optional] 
**SaleOfBusiness** | Pointer to **float32** | &lt;p&gt;{{SaleOfBusiness}} is the amount earned to sell business.&lt;/p&gt; | [optional] 
**SaleOfInvestment** | Pointer to **float32** | &lt;p&gt;{{SaleOfInvestment}} represents cash inflow on the sale of investments in securities.&lt;/p&gt; | [optional] 
**SaleOfPpe** | Pointer to **float32** | &lt;p&gt;{{SaleOfPPE}} is the amount earned to sell {{Net_PPE}}. &lt;/p&gt; | [optional] 
**StockBasedCompensation** | Pointer to **float32** | &lt;p&gt;{{StockBasedCompensation}} is a way corporations use stock options to reward employees. It provides executives and employees the opportunity to share in the growth of the company and, if structured properly, can align their interests with the interests of the company&#39;s shareholders and investors, without burning the company&#39;s cash on hand.&lt;/p&gt; | [optional] 
**TotalFreeCashFlow** | Pointer to **float32** | &lt;p&gt;{{total_freecashflow}} is considered one of the most important parameters to measure a company&#39;s earnings power by value investors because it is not subject to estimates of {{DDA}} (DDA). However, when we look at the {{total_freecashflow}}, we should look from a long term perspective, because any year&#39;s {{total_freecashflow}} can be drastically affected by the spending on {{Net_PPE}} (PPE) of the business in that year. Over the long term, {{total_freecashflow}} should give pretty good picture on the real earnings power of the company. &lt;br&gt;{{total_freecashflow}} is calculated as {{total_freecashflow}} &#x3D; {{cash_Flow_from_Operations}} + {{Cash_Flow_CPEX}}&lt;/p&gt; | [optional] 

## Methods

### NewFundamentalsNNOREITNODIRECTCashflowStatement

`func NewFundamentalsNNOREITNODIRECTCashflowStatement() *FundamentalsNNOREITNODIRECTCashflowStatement`

NewFundamentalsNNOREITNODIRECTCashflowStatement instantiates a new FundamentalsNNOREITNODIRECTCashflowStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsNNOREITNODIRECTCashflowStatementWithDefaults

`func NewFundamentalsNNOREITNODIRECTCashflowStatementWithDefaults() *FundamentalsNNOREITNODIRECTCashflowStatement`

NewFundamentalsNNOREITNODIRECTCashflowStatementWithDefaults instantiates a new FundamentalsNNOREITNODIRECTCashflowStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetImpairmentCharge

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetAssetImpairmentCharge() float32`

GetAssetImpairmentCharge returns the AssetImpairmentCharge field if non-nil, zero value otherwise.

### GetAssetImpairmentChargeOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetAssetImpairmentChargeOk() (*float32, bool)`

GetAssetImpairmentChargeOk returns a tuple with the AssetImpairmentCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetImpairmentCharge

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetAssetImpairmentCharge(v float32)`

SetAssetImpairmentCharge sets AssetImpairmentCharge field to given value.

### HasAssetImpairmentCharge

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasAssetImpairmentCharge() bool`

HasAssetImpairmentCharge returns a boolean if a field has been set.

### GetBeginningCashPosition

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetBeginningCashPosition() float32`

GetBeginningCashPosition returns the BeginningCashPosition field if non-nil, zero value otherwise.

### GetBeginningCashPositionOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetBeginningCashPositionOk() (*float32, bool)`

GetBeginningCashPositionOk returns a tuple with the BeginningCashPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginningCashPosition

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetBeginningCashPosition(v float32)`

SetBeginningCashPosition sets BeginningCashPosition field to given value.

### HasBeginningCashPosition

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasBeginningCashPosition() bool`

HasBeginningCashPosition returns a boolean if a field has been set.

### GetCashFlowCapitalExpenditure

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowCapitalExpenditure() float32`

GetCashFlowCapitalExpenditure returns the CashFlowCapitalExpenditure field if non-nil, zero value otherwise.

### GetCashFlowCapitalExpenditureOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowCapitalExpenditureOk() (*float32, bool)`

GetCashFlowCapitalExpenditureOk returns a tuple with the CashFlowCapitalExpenditure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowCapitalExpenditure

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFlowCapitalExpenditure(v float32)`

SetCashFlowCapitalExpenditure sets CashFlowCapitalExpenditure field to given value.

### HasCashFlowCapitalExpenditure

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFlowCapitalExpenditure() bool`

HasCashFlowCapitalExpenditure returns a boolean if a field has been set.

### GetCashFlowDeferredTax

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowDeferredTax() float32`

GetCashFlowDeferredTax returns the CashFlowDeferredTax field if non-nil, zero value otherwise.

### GetCashFlowDeferredTaxOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowDeferredTaxOk() (*float32, bool)`

GetCashFlowDeferredTaxOk returns a tuple with the CashFlowDeferredTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowDeferredTax

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFlowDeferredTax(v float32)`

SetCashFlowDeferredTax sets CashFlowDeferredTax field to given value.

### HasCashFlowDeferredTax

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFlowDeferredTax() bool`

HasCashFlowDeferredTax returns a boolean if a field has been set.

### GetCashFlowDepreciationDepletionAmortization

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowDepreciationDepletionAmortization() float32`

GetCashFlowDepreciationDepletionAmortization returns the CashFlowDepreciationDepletionAmortization field if non-nil, zero value otherwise.

### GetCashFlowDepreciationDepletionAmortizationOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowDepreciationDepletionAmortizationOk() (*float32, bool)`

GetCashFlowDepreciationDepletionAmortizationOk returns a tuple with the CashFlowDepreciationDepletionAmortization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowDepreciationDepletionAmortization

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFlowDepreciationDepletionAmortization(v float32)`

SetCashFlowDepreciationDepletionAmortization sets CashFlowDepreciationDepletionAmortization field to given value.

### HasCashFlowDepreciationDepletionAmortization

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFlowDepreciationDepletionAmortization() bool`

HasCashFlowDepreciationDepletionAmortization returns a boolean if a field has been set.

### GetCashFlowForLeaseFinancing

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowForLeaseFinancing() float32`

GetCashFlowForLeaseFinancing returns the CashFlowForLeaseFinancing field if non-nil, zero value otherwise.

### GetCashFlowForLeaseFinancingOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowForLeaseFinancingOk() (*float32, bool)`

GetCashFlowForLeaseFinancingOk returns a tuple with the CashFlowForLeaseFinancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowForLeaseFinancing

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFlowForLeaseFinancing(v float32)`

SetCashFlowForLeaseFinancing sets CashFlowForLeaseFinancing field to given value.

### HasCashFlowForLeaseFinancing

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFlowForLeaseFinancing() bool`

HasCashFlowForLeaseFinancing returns a boolean if a field has been set.

### GetCashFlowFromInvesting

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowFromInvesting() float32`

GetCashFlowFromInvesting returns the CashFlowFromInvesting field if non-nil, zero value otherwise.

### GetCashFlowFromInvestingOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowFromInvestingOk() (*float32, bool)`

GetCashFlowFromInvestingOk returns a tuple with the CashFlowFromInvesting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowFromInvesting

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFlowFromInvesting(v float32)`

SetCashFlowFromInvesting sets CashFlowFromInvesting field to given value.

### HasCashFlowFromInvesting

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFlowFromInvesting() bool`

HasCashFlowFromInvesting returns a boolean if a field has been set.

### GetCashFlowFromOperations

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowFromOperations() float32`

GetCashFlowFromOperations returns the CashFlowFromOperations field if non-nil, zero value otherwise.

### GetCashFlowFromOperationsOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowFromOperationsOk() (*float32, bool)`

GetCashFlowFromOperationsOk returns a tuple with the CashFlowFromOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowFromOperations

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFlowFromOperations(v float32)`

SetCashFlowFromOperations sets CashFlowFromOperations field to given value.

### HasCashFlowFromOperations

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFlowFromOperations() bool`

HasCashFlowFromOperations returns a boolean if a field has been set.

### GetCashFlowFromOthers

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowFromOthers() float32`

GetCashFlowFromOthers returns the CashFlowFromOthers field if non-nil, zero value otherwise.

### GetCashFlowFromOthersOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFlowFromOthersOk() (*float32, bool)`

GetCashFlowFromOthersOk returns a tuple with the CashFlowFromOthers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowFromOthers

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFlowFromOthers(v float32)`

SetCashFlowFromOthers sets CashFlowFromOthers field to given value.

### HasCashFlowFromOthers

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFlowFromOthers() bool`

HasCashFlowFromOthers returns a boolean if a field has been set.

### GetCashFromDiscontinuedInvestingActivities

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFromDiscontinuedInvestingActivities() float32`

GetCashFromDiscontinuedInvestingActivities returns the CashFromDiscontinuedInvestingActivities field if non-nil, zero value otherwise.

### GetCashFromDiscontinuedInvestingActivitiesOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFromDiscontinuedInvestingActivitiesOk() (*float32, bool)`

GetCashFromDiscontinuedInvestingActivitiesOk returns a tuple with the CashFromDiscontinuedInvestingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFromDiscontinuedInvestingActivities

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFromDiscontinuedInvestingActivities(v float32)`

SetCashFromDiscontinuedInvestingActivities sets CashFromDiscontinuedInvestingActivities field to given value.

### HasCashFromDiscontinuedInvestingActivities

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFromDiscontinuedInvestingActivities() bool`

HasCashFromDiscontinuedInvestingActivities returns a boolean if a field has been set.

### GetCashFromDiscontinuedOperatingActivities

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFromDiscontinuedOperatingActivities() float32`

GetCashFromDiscontinuedOperatingActivities returns the CashFromDiscontinuedOperatingActivities field if non-nil, zero value otherwise.

### GetCashFromDiscontinuedOperatingActivitiesOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFromDiscontinuedOperatingActivitiesOk() (*float32, bool)`

GetCashFromDiscontinuedOperatingActivitiesOk returns a tuple with the CashFromDiscontinuedOperatingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFromDiscontinuedOperatingActivities

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFromDiscontinuedOperatingActivities(v float32)`

SetCashFromDiscontinuedOperatingActivities sets CashFromDiscontinuedOperatingActivities field to given value.

### HasCashFromDiscontinuedOperatingActivities

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFromDiscontinuedOperatingActivities() bool`

HasCashFromDiscontinuedOperatingActivities returns a boolean if a field has been set.

### GetCashFromFinancing

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFromFinancing() float32`

GetCashFromFinancing returns the CashFromFinancing field if non-nil, zero value otherwise.

### GetCashFromFinancingOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFromFinancingOk() (*float32, bool)`

GetCashFromFinancingOk returns a tuple with the CashFromFinancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFromFinancing

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFromFinancing(v float32)`

SetCashFromFinancing sets CashFromFinancing field to given value.

### HasCashFromFinancing

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFromFinancing() bool`

HasCashFromFinancing returns a boolean if a field has been set.

### GetCashFromOtherInvestingActivities

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFromOtherInvestingActivities() float32`

GetCashFromOtherInvestingActivities returns the CashFromOtherInvestingActivities field if non-nil, zero value otherwise.

### GetCashFromOtherInvestingActivitiesOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetCashFromOtherInvestingActivitiesOk() (*float32, bool)`

GetCashFromOtherInvestingActivitiesOk returns a tuple with the CashFromOtherInvestingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFromOtherInvestingActivities

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetCashFromOtherInvestingActivities(v float32)`

SetCashFromOtherInvestingActivities sets CashFromOtherInvestingActivities field to given value.

### HasCashFromOtherInvestingActivities

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasCashFromOtherInvestingActivities() bool`

HasCashFromOtherInvestingActivities returns a boolean if a field has been set.

### GetChangeInInventory

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInInventory() float32`

GetChangeInInventory returns the ChangeInInventory field if non-nil, zero value otherwise.

### GetChangeInInventoryOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInInventoryOk() (*float32, bool)`

GetChangeInInventoryOk returns a tuple with the ChangeInInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeInInventory

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetChangeInInventory(v float32)`

SetChangeInInventory sets ChangeInInventory field to given value.

### HasChangeInInventory

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasChangeInInventory() bool`

HasChangeInInventory returns a boolean if a field has been set.

### GetChangeInOtherWorkingCapital

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInOtherWorkingCapital() float32`

GetChangeInOtherWorkingCapital returns the ChangeInOtherWorkingCapital field if non-nil, zero value otherwise.

### GetChangeInOtherWorkingCapitalOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInOtherWorkingCapitalOk() (*float32, bool)`

GetChangeInOtherWorkingCapitalOk returns a tuple with the ChangeInOtherWorkingCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeInOtherWorkingCapital

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetChangeInOtherWorkingCapital(v float32)`

SetChangeInOtherWorkingCapital sets ChangeInOtherWorkingCapital field to given value.

### HasChangeInOtherWorkingCapital

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasChangeInOtherWorkingCapital() bool`

HasChangeInOtherWorkingCapital returns a boolean if a field has been set.

### GetChangeInPayablesAndAccruedExpense

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInPayablesAndAccruedExpense() float32`

GetChangeInPayablesAndAccruedExpense returns the ChangeInPayablesAndAccruedExpense field if non-nil, zero value otherwise.

### GetChangeInPayablesAndAccruedExpenseOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInPayablesAndAccruedExpenseOk() (*float32, bool)`

GetChangeInPayablesAndAccruedExpenseOk returns a tuple with the ChangeInPayablesAndAccruedExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeInPayablesAndAccruedExpense

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetChangeInPayablesAndAccruedExpense(v float32)`

SetChangeInPayablesAndAccruedExpense sets ChangeInPayablesAndAccruedExpense field to given value.

### HasChangeInPayablesAndAccruedExpense

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasChangeInPayablesAndAccruedExpense() bool`

HasChangeInPayablesAndAccruedExpense returns a boolean if a field has been set.

### GetChangeInPrepaidAssets

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInPrepaidAssets() float32`

GetChangeInPrepaidAssets returns the ChangeInPrepaidAssets field if non-nil, zero value otherwise.

### GetChangeInPrepaidAssetsOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInPrepaidAssetsOk() (*float32, bool)`

GetChangeInPrepaidAssetsOk returns a tuple with the ChangeInPrepaidAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeInPrepaidAssets

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetChangeInPrepaidAssets(v float32)`

SetChangeInPrepaidAssets sets ChangeInPrepaidAssets field to given value.

### HasChangeInPrepaidAssets

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasChangeInPrepaidAssets() bool`

HasChangeInPrepaidAssets returns a boolean if a field has been set.

### GetChangeInReceivables

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInReceivables() float32`

GetChangeInReceivables returns the ChangeInReceivables field if non-nil, zero value otherwise.

### GetChangeInReceivablesOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInReceivablesOk() (*float32, bool)`

GetChangeInReceivablesOk returns a tuple with the ChangeInReceivables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeInReceivables

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetChangeInReceivables(v float32)`

SetChangeInReceivables sets ChangeInReceivables field to given value.

### HasChangeInReceivables

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasChangeInReceivables() bool`

HasChangeInReceivables returns a boolean if a field has been set.

### GetChangeInWorkingCapital

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInWorkingCapital() float32`

GetChangeInWorkingCapital returns the ChangeInWorkingCapital field if non-nil, zero value otherwise.

### GetChangeInWorkingCapitalOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetChangeInWorkingCapitalOk() (*float32, bool)`

GetChangeInWorkingCapitalOk returns a tuple with the ChangeInWorkingCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeInWorkingCapital

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetChangeInWorkingCapital(v float32)`

SetChangeInWorkingCapital sets ChangeInWorkingCapital field to given value.

### HasChangeInWorkingCapital

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasChangeInWorkingCapital() bool`

HasChangeInWorkingCapital returns a boolean if a field has been set.

### GetDebtIssuance

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetDebtIssuance() float32`

GetDebtIssuance returns the DebtIssuance field if non-nil, zero value otherwise.

### GetDebtIssuanceOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetDebtIssuanceOk() (*float32, bool)`

GetDebtIssuanceOk returns a tuple with the DebtIssuance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtIssuance

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetDebtIssuance(v float32)`

SetDebtIssuance sets DebtIssuance field to given value.

### HasDebtIssuance

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasDebtIssuance() bool`

HasDebtIssuance returns a boolean if a field has been set.

### GetDebtPayments

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetDebtPayments() float32`

GetDebtPayments returns the DebtPayments field if non-nil, zero value otherwise.

### GetDebtPaymentsOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetDebtPaymentsOk() (*float32, bool)`

GetDebtPaymentsOk returns a tuple with the DebtPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtPayments

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetDebtPayments(v float32)`

SetDebtPayments sets DebtPayments field to given value.

### HasDebtPayments

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasDebtPayments() bool`

HasDebtPayments returns a boolean if a field has been set.

### GetDividends

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetDividends() float32`

GetDividends returns the Dividends field if non-nil, zero value otherwise.

### GetDividendsOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetDividendsOk() (*float32, bool)`

GetDividendsOk returns a tuple with the Dividends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividends

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetDividends(v float32)`

SetDividends sets Dividends field to given value.

### HasDividends

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasDividends() bool`

HasDividends returns a boolean if a field has been set.

### GetEffectOfExchangeRateChanges

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetEffectOfExchangeRateChanges() float32`

GetEffectOfExchangeRateChanges returns the EffectOfExchangeRateChanges field if non-nil, zero value otherwise.

### GetEffectOfExchangeRateChangesOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetEffectOfExchangeRateChangesOk() (*float32, bool)`

GetEffectOfExchangeRateChangesOk returns a tuple with the EffectOfExchangeRateChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectOfExchangeRateChanges

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetEffectOfExchangeRateChanges(v float32)`

SetEffectOfExchangeRateChanges sets EffectOfExchangeRateChanges field to given value.

### HasEffectOfExchangeRateChanges

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasEffectOfExchangeRateChanges() bool`

HasEffectOfExchangeRateChanges returns a boolean if a field has been set.

### GetEndingCashPosition

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetEndingCashPosition() float32`

GetEndingCashPosition returns the EndingCashPosition field if non-nil, zero value otherwise.

### GetEndingCashPositionOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetEndingCashPositionOk() (*float32, bool)`

GetEndingCashPositionOk returns a tuple with the EndingCashPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingCashPosition

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetEndingCashPosition(v float32)`

SetEndingCashPosition sets EndingCashPosition field to given value.

### HasEndingCashPosition

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasEndingCashPosition() bool`

HasEndingCashPosition returns a boolean if a field has been set.

### GetIssuanceOfStock

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetIssuanceOfStock() float32`

GetIssuanceOfStock returns the IssuanceOfStock field if non-nil, zero value otherwise.

### GetIssuanceOfStockOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetIssuanceOfStockOk() (*float32, bool)`

GetIssuanceOfStockOk returns a tuple with the IssuanceOfStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceOfStock

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetIssuanceOfStock(v float32)`

SetIssuanceOfStock sets IssuanceOfStock field to given value.

### HasIssuanceOfStock

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasIssuanceOfStock() bool`

HasIssuanceOfStock returns a boolean if a field has been set.

### GetNetChangeInCash

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetNetChangeInCash() float32`

GetNetChangeInCash returns the NetChangeInCash field if non-nil, zero value otherwise.

### GetNetChangeInCashOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetNetChangeInCashOk() (*float32, bool)`

GetNetChangeInCashOk returns a tuple with the NetChangeInCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChangeInCash

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetNetChangeInCash(v float32)`

SetNetChangeInCash sets NetChangeInCash field to given value.

### HasNetChangeInCash

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasNetChangeInCash() bool`

HasNetChangeInCash returns a boolean if a field has been set.

### GetNetIncomeFromContinuingOperations

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetNetIncomeFromContinuingOperations() float32`

GetNetIncomeFromContinuingOperations returns the NetIncomeFromContinuingOperations field if non-nil, zero value otherwise.

### GetNetIncomeFromContinuingOperationsOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetNetIncomeFromContinuingOperationsOk() (*float32, bool)`

GetNetIncomeFromContinuingOperationsOk returns a tuple with the NetIncomeFromContinuingOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncomeFromContinuingOperations

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetNetIncomeFromContinuingOperations(v float32)`

SetNetIncomeFromContinuingOperations sets NetIncomeFromContinuingOperations field to given value.

### HasNetIncomeFromContinuingOperations

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasNetIncomeFromContinuingOperations() bool`

HasNetIncomeFromContinuingOperations returns a boolean if a field has been set.

### GetNetIntangiblesPurchaseAndSale

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetNetIntangiblesPurchaseAndSale() float32`

GetNetIntangiblesPurchaseAndSale returns the NetIntangiblesPurchaseAndSale field if non-nil, zero value otherwise.

### GetNetIntangiblesPurchaseAndSaleOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetNetIntangiblesPurchaseAndSaleOk() (*float32, bool)`

GetNetIntangiblesPurchaseAndSaleOk returns a tuple with the NetIntangiblesPurchaseAndSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIntangiblesPurchaseAndSale

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetNetIntangiblesPurchaseAndSale(v float32)`

SetNetIntangiblesPurchaseAndSale sets NetIntangiblesPurchaseAndSale field to given value.

### HasNetIntangiblesPurchaseAndSale

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasNetIntangiblesPurchaseAndSale() bool`

HasNetIntangiblesPurchaseAndSale returns a boolean if a field has been set.

### GetNetIssuanceOfDebt

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetNetIssuanceOfDebt() float32`

GetNetIssuanceOfDebt returns the NetIssuanceOfDebt field if non-nil, zero value otherwise.

### GetNetIssuanceOfDebtOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetNetIssuanceOfDebtOk() (*float32, bool)`

GetNetIssuanceOfDebtOk returns a tuple with the NetIssuanceOfDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIssuanceOfDebt

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetNetIssuanceOfDebt(v float32)`

SetNetIssuanceOfDebt sets NetIssuanceOfDebt field to given value.

### HasNetIssuanceOfDebt

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasNetIssuanceOfDebt() bool`

HasNetIssuanceOfDebt returns a boolean if a field has been set.

### GetNetIssuanceOfPreferred

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetNetIssuanceOfPreferred() float32`

GetNetIssuanceOfPreferred returns the NetIssuanceOfPreferred field if non-nil, zero value otherwise.

### GetNetIssuanceOfPreferredOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetNetIssuanceOfPreferredOk() (*float32, bool)`

GetNetIssuanceOfPreferredOk returns a tuple with the NetIssuanceOfPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIssuanceOfPreferred

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetNetIssuanceOfPreferred(v float32)`

SetNetIssuanceOfPreferred sets NetIssuanceOfPreferred field to given value.

### HasNetIssuanceOfPreferred

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasNetIssuanceOfPreferred() bool`

HasNetIssuanceOfPreferred returns a boolean if a field has been set.

### GetOtherFinancing

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetOtherFinancing() float32`

GetOtherFinancing returns the OtherFinancing field if non-nil, zero value otherwise.

### GetOtherFinancingOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetOtherFinancingOk() (*float32, bool)`

GetOtherFinancingOk returns a tuple with the OtherFinancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFinancing

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetOtherFinancing(v float32)`

SetOtherFinancing sets OtherFinancing field to given value.

### HasOtherFinancing

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasOtherFinancing() bool`

HasOtherFinancing returns a boolean if a field has been set.

### GetPurchaseOfBusiness

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetPurchaseOfBusiness() float32`

GetPurchaseOfBusiness returns the PurchaseOfBusiness field if non-nil, zero value otherwise.

### GetPurchaseOfBusinessOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetPurchaseOfBusinessOk() (*float32, bool)`

GetPurchaseOfBusinessOk returns a tuple with the PurchaseOfBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOfBusiness

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetPurchaseOfBusiness(v float32)`

SetPurchaseOfBusiness sets PurchaseOfBusiness field to given value.

### HasPurchaseOfBusiness

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasPurchaseOfBusiness() bool`

HasPurchaseOfBusiness returns a boolean if a field has been set.

### GetPurchaseOfInvestment

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetPurchaseOfInvestment() float32`

GetPurchaseOfInvestment returns the PurchaseOfInvestment field if non-nil, zero value otherwise.

### GetPurchaseOfInvestmentOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetPurchaseOfInvestmentOk() (*float32, bool)`

GetPurchaseOfInvestmentOk returns a tuple with the PurchaseOfInvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOfInvestment

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetPurchaseOfInvestment(v float32)`

SetPurchaseOfInvestment sets PurchaseOfInvestment field to given value.

### HasPurchaseOfInvestment

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasPurchaseOfInvestment() bool`

HasPurchaseOfInvestment returns a boolean if a field has been set.

### GetPurchaseOfPpe

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetPurchaseOfPpe() float32`

GetPurchaseOfPpe returns the PurchaseOfPpe field if non-nil, zero value otherwise.

### GetPurchaseOfPpeOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetPurchaseOfPpeOk() (*float32, bool)`

GetPurchaseOfPpeOk returns a tuple with the PurchaseOfPpe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOfPpe

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetPurchaseOfPpe(v float32)`

SetPurchaseOfPpe sets PurchaseOfPpe field to given value.

### HasPurchaseOfPpe

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasPurchaseOfPpe() bool`

HasPurchaseOfPpe returns a boolean if a field has been set.

### GetRepurchaseOfStock

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetRepurchaseOfStock() float32`

GetRepurchaseOfStock returns the RepurchaseOfStock field if non-nil, zero value otherwise.

### GetRepurchaseOfStockOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetRepurchaseOfStockOk() (*float32, bool)`

GetRepurchaseOfStockOk returns a tuple with the RepurchaseOfStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepurchaseOfStock

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetRepurchaseOfStock(v float32)`

SetRepurchaseOfStock sets RepurchaseOfStock field to given value.

### HasRepurchaseOfStock

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasRepurchaseOfStock() bool`

HasRepurchaseOfStock returns a boolean if a field has been set.

### GetSaleOfBusiness

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetSaleOfBusiness() float32`

GetSaleOfBusiness returns the SaleOfBusiness field if non-nil, zero value otherwise.

### GetSaleOfBusinessOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetSaleOfBusinessOk() (*float32, bool)`

GetSaleOfBusinessOk returns a tuple with the SaleOfBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleOfBusiness

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetSaleOfBusiness(v float32)`

SetSaleOfBusiness sets SaleOfBusiness field to given value.

### HasSaleOfBusiness

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasSaleOfBusiness() bool`

HasSaleOfBusiness returns a boolean if a field has been set.

### GetSaleOfInvestment

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetSaleOfInvestment() float32`

GetSaleOfInvestment returns the SaleOfInvestment field if non-nil, zero value otherwise.

### GetSaleOfInvestmentOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetSaleOfInvestmentOk() (*float32, bool)`

GetSaleOfInvestmentOk returns a tuple with the SaleOfInvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleOfInvestment

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetSaleOfInvestment(v float32)`

SetSaleOfInvestment sets SaleOfInvestment field to given value.

### HasSaleOfInvestment

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasSaleOfInvestment() bool`

HasSaleOfInvestment returns a boolean if a field has been set.

### GetSaleOfPpe

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetSaleOfPpe() float32`

GetSaleOfPpe returns the SaleOfPpe field if non-nil, zero value otherwise.

### GetSaleOfPpeOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetSaleOfPpeOk() (*float32, bool)`

GetSaleOfPpeOk returns a tuple with the SaleOfPpe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleOfPpe

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetSaleOfPpe(v float32)`

SetSaleOfPpe sets SaleOfPpe field to given value.

### HasSaleOfPpe

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasSaleOfPpe() bool`

HasSaleOfPpe returns a boolean if a field has been set.

### GetStockBasedCompensation

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetStockBasedCompensation() float32`

GetStockBasedCompensation returns the StockBasedCompensation field if non-nil, zero value otherwise.

### GetStockBasedCompensationOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetStockBasedCompensationOk() (*float32, bool)`

GetStockBasedCompensationOk returns a tuple with the StockBasedCompensation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockBasedCompensation

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetStockBasedCompensation(v float32)`

SetStockBasedCompensation sets StockBasedCompensation field to given value.

### HasStockBasedCompensation

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasStockBasedCompensation() bool`

HasStockBasedCompensation returns a boolean if a field has been set.

### GetTotalFreeCashFlow

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetTotalFreeCashFlow() float32`

GetTotalFreeCashFlow returns the TotalFreeCashFlow field if non-nil, zero value otherwise.

### GetTotalFreeCashFlowOk

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) GetTotalFreeCashFlowOk() (*float32, bool)`

GetTotalFreeCashFlowOk returns a tuple with the TotalFreeCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFreeCashFlow

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) SetTotalFreeCashFlow(v float32)`

SetTotalFreeCashFlow sets TotalFreeCashFlow field to given value.

### HasTotalFreeCashFlow

`func (o *FundamentalsNNOREITNODIRECTCashflowStatement) HasTotalFreeCashFlow() bool`

HasTotalFreeCashFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



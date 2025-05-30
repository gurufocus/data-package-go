# ValuationsNNOREITNODIRECTValuationRatios

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CyclicallyAdjustedPbRatio** | Pointer to **float32** | The Cyclically Adjusted PB Ratio, also known as the CAPB Ratio, is the stock price divided by the average of the inflation adjusted book value per share of a company over the past 10 years. | [optional] 
**CyclicallyAdjustedPriceToFcf** | Pointer to **float32** |  | [optional] 
**CyclicallyAdjustedPsRatio** | Pointer to **float32** | The Cyclically Adjusted PS Ratio, also referred to as the CAPS Ratio, is the stock price divided by the average of the inflation adjusted revenue per share of a company over the past 10 years. | [optional] 
**EarningYieldGreenblatt** | Pointer to **float32** | The standard definition of earnings yield is the earnings per share divided by the price of a share. It&#39;s the inverse of P/E and shows the amount of money earned compared to the price you pay for a share. | [optional] 
**EnterpriseValueToEbit** | Pointer to **float32** | EV-to-EBIT is calculated as Enterprise Value divided by its EBIT. | [optional] 
**EnterpriseValueToEbitda** | Pointer to **float32** | EV-to-EBITDA is calculated as enterprise value divided by its EBITDA. | [optional] 
**EnterpriseValueToFcf** | Pointer to **float32** |  | [optional] 
**EnterpriseValueToRevenue** | Pointer to **float32** | EV-to-Revenue is calculated as enterprise value divided by its revenue. | [optional] 
**FcfYield** | Pointer to **float32** | Free cash flow yield: the free cash flow divided by share price | [optional] 
**PbRatio** | Pointer to **float32** | Companies use the price-to-book ratio to compare a firm&#39;s market to book value by dividing price per share by book value per share. Some people know it as the price-equity ratio. | [optional] 
**PeRatio** | Pointer to **float32** | &lt;p&gt;The PE ratio, or Price-to-Earnings ratio, is the most widely used metric in the valuation of stocks. It is calculated as:  &lt;b&gt;PE Ratio &#x3D; Share Price / {{eps_diluated}}&lt;/b&gt;.   It can also be calculated from the numbers for the whole company:  &lt;b&gt;PE Ratio &#x3D; {{mktcap}} / {{net_income}}&lt;/b&gt;.&lt;/p&gt;  &lt;p&gt;There are at least three kinds of PE ratios used among different investors. They are Trailing Twelve Month PE Ratio({{pettm}}), {{forwardPE}}, and {{penri}}. A new PE ratio based on inflation-adjusted normalized PE ratio is called {{ShillerPE}}, after Yale professor Robert Shiller.&lt;/p&gt;  &lt;p&gt;In the calculation of {{pettm}}, the earnings per share used are the earnings per share over the past 12 months({{ttm_eps}}). For {{forwardPE}}, the earnings are the expected earnings for the next twelve months({{forward_eps}}). In the case of {{penri}}, the reported earnings less the non-recurring items are used({{eps_nri}}).For the {{ShillerPE}}, the earnings of the past 10 years are inflation-adjusted and averaged. Since {{ShillerPE}} looks at the average over the last 10 years, it is also called PE10.&lt;/p&gt; | [optional] 
**PegRatio** | Pointer to **float32** | PEG is defined as the PE Ratio without NRI divided by the growth ratio. The growth rate we use is the 5-year average EBITDA growth rate. | [optional] 
**Penri** | Pointer to **float32** |  | [optional] 
**PriceToFreeCashFlow** | Pointer to **float32** | Price to free cash flow is an equity valuation metric used to compare a company&#39;s per-share market price to its per-share amount of free cash flow (FCF). | [optional] 
**PriceToOperatingCashFlow** | Pointer to **float32** | The price/cash flow ratio (also called price-to-cash flow ratio or P/CF), is a ratio used to compare a company&#39;s market value to its cash flow. | [optional] 
**PriceToOwnerEarnings** | Pointer to **float32** | In 1986 Berkshire Hathaway Shareholder Letter, Warren Buffett defined owner earnings as follows:  \&quot;These represent (a) reported earnings plus (b) depreciation, depletion, amortization, and certain other non-cash charges...less (c) the average annual amount of capitalized expenditures for plant and equipment, etc. that the business requires to fully maintain its long-term competitive position and its unit volume. (If the business requires additional working capital to maintain its competitive position and unit volume, the increment also should be included in (c))...Our owner-earnings equation does not yield the deceptively precise figures provided by GAAP, since (c) must be a guess - and one sometimes very difficult to make. Despite this problem, we consider the owner earnings figure, not the GAAP figure, to be the relevant item for valuation purposes - both for investors in buying stocks and for managers in buying entire businesses...All of this points up the absurdity of the &#39;cash flow&#39; numbers that are often set forth in Wall Street reports. These numbers routinely include (a) plus (b) - but do not subtract (c).\&quot; | [optional] 
**PriceToTangibleBook** | Pointer to **float32** | The Price to Tangible Book Value ratio (PTBV) expresses share price as a proportion of the company&#39;s tangible book value reported on the company&#39;s balance sheet. | [optional] 
**PsRatio** | Pointer to **float32** | Price–sales ratio, P/S ratio, or PSR, is a valuation metric for stocks. | [optional] 
**RateOfReturnValue** | Pointer to **float32** | Yacktman defines forward rate of return as the normalized free cash flow yield plus real growth plus inflation. | [optional] 
**ShillerPeRatio** | Pointer to **float32** | Price earnings ratio is based on average inflation-adjusted earnings from the previous 10 years, known as the Cyclically Adjusted PE Ratio (CAPE Ratio), Shiller PE Ratio, or PE 10 | [optional] 
**Yield** | Pointer to **float32** | The dividend yield is the ratio of a company&#39;s annual dividend compared to its share price. | [optional] 

## Methods

### NewValuationsNNOREITNODIRECTValuationRatios

`func NewValuationsNNOREITNODIRECTValuationRatios() *ValuationsNNOREITNODIRECTValuationRatios`

NewValuationsNNOREITNODIRECTValuationRatios instantiates a new ValuationsNNOREITNODIRECTValuationRatios object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuationsNNOREITNODIRECTValuationRatiosWithDefaults

`func NewValuationsNNOREITNODIRECTValuationRatiosWithDefaults() *ValuationsNNOREITNODIRECTValuationRatios`

NewValuationsNNOREITNODIRECTValuationRatiosWithDefaults instantiates a new ValuationsNNOREITNODIRECTValuationRatios object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCyclicallyAdjustedPbRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetCyclicallyAdjustedPbRatio() float32`

GetCyclicallyAdjustedPbRatio returns the CyclicallyAdjustedPbRatio field if non-nil, zero value otherwise.

### GetCyclicallyAdjustedPbRatioOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetCyclicallyAdjustedPbRatioOk() (*float32, bool)`

GetCyclicallyAdjustedPbRatioOk returns a tuple with the CyclicallyAdjustedPbRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyclicallyAdjustedPbRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetCyclicallyAdjustedPbRatio(v float32)`

SetCyclicallyAdjustedPbRatio sets CyclicallyAdjustedPbRatio field to given value.

### HasCyclicallyAdjustedPbRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasCyclicallyAdjustedPbRatio() bool`

HasCyclicallyAdjustedPbRatio returns a boolean if a field has been set.

### GetCyclicallyAdjustedPriceToFcf

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetCyclicallyAdjustedPriceToFcf() float32`

GetCyclicallyAdjustedPriceToFcf returns the CyclicallyAdjustedPriceToFcf field if non-nil, zero value otherwise.

### GetCyclicallyAdjustedPriceToFcfOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetCyclicallyAdjustedPriceToFcfOk() (*float32, bool)`

GetCyclicallyAdjustedPriceToFcfOk returns a tuple with the CyclicallyAdjustedPriceToFcf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyclicallyAdjustedPriceToFcf

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetCyclicallyAdjustedPriceToFcf(v float32)`

SetCyclicallyAdjustedPriceToFcf sets CyclicallyAdjustedPriceToFcf field to given value.

### HasCyclicallyAdjustedPriceToFcf

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasCyclicallyAdjustedPriceToFcf() bool`

HasCyclicallyAdjustedPriceToFcf returns a boolean if a field has been set.

### GetCyclicallyAdjustedPsRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetCyclicallyAdjustedPsRatio() float32`

GetCyclicallyAdjustedPsRatio returns the CyclicallyAdjustedPsRatio field if non-nil, zero value otherwise.

### GetCyclicallyAdjustedPsRatioOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetCyclicallyAdjustedPsRatioOk() (*float32, bool)`

GetCyclicallyAdjustedPsRatioOk returns a tuple with the CyclicallyAdjustedPsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyclicallyAdjustedPsRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetCyclicallyAdjustedPsRatio(v float32)`

SetCyclicallyAdjustedPsRatio sets CyclicallyAdjustedPsRatio field to given value.

### HasCyclicallyAdjustedPsRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasCyclicallyAdjustedPsRatio() bool`

HasCyclicallyAdjustedPsRatio returns a boolean if a field has been set.

### GetEarningYieldGreenblatt

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetEarningYieldGreenblatt() float32`

GetEarningYieldGreenblatt returns the EarningYieldGreenblatt field if non-nil, zero value otherwise.

### GetEarningYieldGreenblattOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetEarningYieldGreenblattOk() (*float32, bool)`

GetEarningYieldGreenblattOk returns a tuple with the EarningYieldGreenblatt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningYieldGreenblatt

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetEarningYieldGreenblatt(v float32)`

SetEarningYieldGreenblatt sets EarningYieldGreenblatt field to given value.

### HasEarningYieldGreenblatt

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasEarningYieldGreenblatt() bool`

HasEarningYieldGreenblatt returns a boolean if a field has been set.

### GetEnterpriseValueToEbit

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetEnterpriseValueToEbit() float32`

GetEnterpriseValueToEbit returns the EnterpriseValueToEbit field if non-nil, zero value otherwise.

### GetEnterpriseValueToEbitOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetEnterpriseValueToEbitOk() (*float32, bool)`

GetEnterpriseValueToEbitOk returns a tuple with the EnterpriseValueToEbit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseValueToEbit

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetEnterpriseValueToEbit(v float32)`

SetEnterpriseValueToEbit sets EnterpriseValueToEbit field to given value.

### HasEnterpriseValueToEbit

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasEnterpriseValueToEbit() bool`

HasEnterpriseValueToEbit returns a boolean if a field has been set.

### GetEnterpriseValueToEbitda

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetEnterpriseValueToEbitda() float32`

GetEnterpriseValueToEbitda returns the EnterpriseValueToEbitda field if non-nil, zero value otherwise.

### GetEnterpriseValueToEbitdaOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetEnterpriseValueToEbitdaOk() (*float32, bool)`

GetEnterpriseValueToEbitdaOk returns a tuple with the EnterpriseValueToEbitda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseValueToEbitda

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetEnterpriseValueToEbitda(v float32)`

SetEnterpriseValueToEbitda sets EnterpriseValueToEbitda field to given value.

### HasEnterpriseValueToEbitda

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasEnterpriseValueToEbitda() bool`

HasEnterpriseValueToEbitda returns a boolean if a field has been set.

### GetEnterpriseValueToFcf

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetEnterpriseValueToFcf() float32`

GetEnterpriseValueToFcf returns the EnterpriseValueToFcf field if non-nil, zero value otherwise.

### GetEnterpriseValueToFcfOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetEnterpriseValueToFcfOk() (*float32, bool)`

GetEnterpriseValueToFcfOk returns a tuple with the EnterpriseValueToFcf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseValueToFcf

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetEnterpriseValueToFcf(v float32)`

SetEnterpriseValueToFcf sets EnterpriseValueToFcf field to given value.

### HasEnterpriseValueToFcf

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasEnterpriseValueToFcf() bool`

HasEnterpriseValueToFcf returns a boolean if a field has been set.

### GetEnterpriseValueToRevenue

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetEnterpriseValueToRevenue() float32`

GetEnterpriseValueToRevenue returns the EnterpriseValueToRevenue field if non-nil, zero value otherwise.

### GetEnterpriseValueToRevenueOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetEnterpriseValueToRevenueOk() (*float32, bool)`

GetEnterpriseValueToRevenueOk returns a tuple with the EnterpriseValueToRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseValueToRevenue

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetEnterpriseValueToRevenue(v float32)`

SetEnterpriseValueToRevenue sets EnterpriseValueToRevenue field to given value.

### HasEnterpriseValueToRevenue

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasEnterpriseValueToRevenue() bool`

HasEnterpriseValueToRevenue returns a boolean if a field has been set.

### GetFcfYield

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetFcfYield() float32`

GetFcfYield returns the FcfYield field if non-nil, zero value otherwise.

### GetFcfYieldOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetFcfYieldOk() (*float32, bool)`

GetFcfYieldOk returns a tuple with the FcfYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcfYield

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetFcfYield(v float32)`

SetFcfYield sets FcfYield field to given value.

### HasFcfYield

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasFcfYield() bool`

HasFcfYield returns a boolean if a field has been set.

### GetPbRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPbRatio() float32`

GetPbRatio returns the PbRatio field if non-nil, zero value otherwise.

### GetPbRatioOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPbRatioOk() (*float32, bool)`

GetPbRatioOk returns a tuple with the PbRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetPbRatio(v float32)`

SetPbRatio sets PbRatio field to given value.

### HasPbRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasPbRatio() bool`

HasPbRatio returns a boolean if a field has been set.

### GetPeRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPeRatio() float32`

GetPeRatio returns the PeRatio field if non-nil, zero value otherwise.

### GetPeRatioOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPeRatioOk() (*float32, bool)`

GetPeRatioOk returns a tuple with the PeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetPeRatio(v float32)`

SetPeRatio sets PeRatio field to given value.

### HasPeRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasPeRatio() bool`

HasPeRatio returns a boolean if a field has been set.

### GetPegRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPegRatio() float32`

GetPegRatio returns the PegRatio field if non-nil, zero value otherwise.

### GetPegRatioOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPegRatioOk() (*float32, bool)`

GetPegRatioOk returns a tuple with the PegRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetPegRatio(v float32)`

SetPegRatio sets PegRatio field to given value.

### HasPegRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasPegRatio() bool`

HasPegRatio returns a boolean if a field has been set.

### GetPenri

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPenri() float32`

GetPenri returns the Penri field if non-nil, zero value otherwise.

### GetPenriOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPenriOk() (*float32, bool)`

GetPenriOk returns a tuple with the Penri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenri

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetPenri(v float32)`

SetPenri sets Penri field to given value.

### HasPenri

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasPenri() bool`

HasPenri returns a boolean if a field has been set.

### GetPriceToFreeCashFlow

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPriceToFreeCashFlow() float32`

GetPriceToFreeCashFlow returns the PriceToFreeCashFlow field if non-nil, zero value otherwise.

### GetPriceToFreeCashFlowOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPriceToFreeCashFlowOk() (*float32, bool)`

GetPriceToFreeCashFlowOk returns a tuple with the PriceToFreeCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceToFreeCashFlow

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetPriceToFreeCashFlow(v float32)`

SetPriceToFreeCashFlow sets PriceToFreeCashFlow field to given value.

### HasPriceToFreeCashFlow

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasPriceToFreeCashFlow() bool`

HasPriceToFreeCashFlow returns a boolean if a field has been set.

### GetPriceToOperatingCashFlow

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPriceToOperatingCashFlow() float32`

GetPriceToOperatingCashFlow returns the PriceToOperatingCashFlow field if non-nil, zero value otherwise.

### GetPriceToOperatingCashFlowOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPriceToOperatingCashFlowOk() (*float32, bool)`

GetPriceToOperatingCashFlowOk returns a tuple with the PriceToOperatingCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceToOperatingCashFlow

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetPriceToOperatingCashFlow(v float32)`

SetPriceToOperatingCashFlow sets PriceToOperatingCashFlow field to given value.

### HasPriceToOperatingCashFlow

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasPriceToOperatingCashFlow() bool`

HasPriceToOperatingCashFlow returns a boolean if a field has been set.

### GetPriceToOwnerEarnings

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPriceToOwnerEarnings() float32`

GetPriceToOwnerEarnings returns the PriceToOwnerEarnings field if non-nil, zero value otherwise.

### GetPriceToOwnerEarningsOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPriceToOwnerEarningsOk() (*float32, bool)`

GetPriceToOwnerEarningsOk returns a tuple with the PriceToOwnerEarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceToOwnerEarnings

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetPriceToOwnerEarnings(v float32)`

SetPriceToOwnerEarnings sets PriceToOwnerEarnings field to given value.

### HasPriceToOwnerEarnings

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasPriceToOwnerEarnings() bool`

HasPriceToOwnerEarnings returns a boolean if a field has been set.

### GetPriceToTangibleBook

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPriceToTangibleBook() float32`

GetPriceToTangibleBook returns the PriceToTangibleBook field if non-nil, zero value otherwise.

### GetPriceToTangibleBookOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPriceToTangibleBookOk() (*float32, bool)`

GetPriceToTangibleBookOk returns a tuple with the PriceToTangibleBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceToTangibleBook

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetPriceToTangibleBook(v float32)`

SetPriceToTangibleBook sets PriceToTangibleBook field to given value.

### HasPriceToTangibleBook

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasPriceToTangibleBook() bool`

HasPriceToTangibleBook returns a boolean if a field has been set.

### GetPsRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPsRatio() float32`

GetPsRatio returns the PsRatio field if non-nil, zero value otherwise.

### GetPsRatioOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetPsRatioOk() (*float32, bool)`

GetPsRatioOk returns a tuple with the PsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetPsRatio(v float32)`

SetPsRatio sets PsRatio field to given value.

### HasPsRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasPsRatio() bool`

HasPsRatio returns a boolean if a field has been set.

### GetRateOfReturnValue

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetRateOfReturnValue() float32`

GetRateOfReturnValue returns the RateOfReturnValue field if non-nil, zero value otherwise.

### GetRateOfReturnValueOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetRateOfReturnValueOk() (*float32, bool)`

GetRateOfReturnValueOk returns a tuple with the RateOfReturnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateOfReturnValue

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetRateOfReturnValue(v float32)`

SetRateOfReturnValue sets RateOfReturnValue field to given value.

### HasRateOfReturnValue

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasRateOfReturnValue() bool`

HasRateOfReturnValue returns a boolean if a field has been set.

### GetShillerPeRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetShillerPeRatio() float32`

GetShillerPeRatio returns the ShillerPeRatio field if non-nil, zero value otherwise.

### GetShillerPeRatioOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetShillerPeRatioOk() (*float32, bool)`

GetShillerPeRatioOk returns a tuple with the ShillerPeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShillerPeRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetShillerPeRatio(v float32)`

SetShillerPeRatio sets ShillerPeRatio field to given value.

### HasShillerPeRatio

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasShillerPeRatio() bool`

HasShillerPeRatio returns a boolean if a field has been set.

### GetYield

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetYield() float32`

GetYield returns the Yield field if non-nil, zero value otherwise.

### GetYieldOk

`func (o *ValuationsNNOREITNODIRECTValuationRatios) GetYieldOk() (*float32, bool)`

GetYieldOk returns a tuple with the Yield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYield

`func (o *ValuationsNNOREITNODIRECTValuationRatios) SetYield(v float32)`

SetYield sets Yield field to given value.

### HasYield

`func (o *ValuationsNNOREITNODIRECTValuationRatios) HasYield() bool`

HasYield returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Go API client for openapi

API for accessing Gurufocus data packages, please go to [https://www.gurufocus.com/user/me?tab=account&subtab=api-token](https://www.gurufocus.com/user/me?tab=account&subtab=api-token) to view or generate authorization keys.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.12.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/gurufocus/data-package-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.gurufocus.com/data*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DataPackagesAPI* | [**DownloadIdGet**](docs/DataPackagesAPI.md#downloadidget) | **Get** /download/{id} | Get download url of the data package
*DataPackagesAPI* | [**DownloadListGet**](docs/DataPackagesAPI.md#downloadlistget) | **Get** /download/list | List available data packages
*ETFAPI* | [**EtfSymbolGet**](docs/ETFAPI.md#etfsymbolget) | **Get** /etf/{symbol} | ETF profile, key statistics and holdings.
*GURUAPI* | [**GurusIdGet**](docs/GURUAPI.md#gurusidget) | **Get** /gurus/{id} | Access the holdings and trades of over 4,500 institutional investors, enabling broad market trend analysis or fueling investment research tools with extensive institutional activity data
*GURUAPI* | [**GurusListGet**](docs/GURUAPI.md#guruslistget) | **Get** /gurus/list | Get the list of available gurus&#39; basic information
*INSIDERAPI* | [**InsidersDateGet**](docs/INSIDERAPI.md#insidersdateget) | **Get** /insiders/{date} | A comprehensive record of insider trading and institutional ownership for investment research or powering data-driven financial tools that track executive trading behavior over time
*STOCKAPI* | [**StocksSymbolFundamentalsGet**](docs/STOCKAPI.md#stockssymbolfundamentalsget) | **Get** /stocks/{symbol}/fundamentals | Complete historical financial statements providing the data needed for in-depth analysis, valuation modeling, and building investment research platforms.
*STOCKAPI* | [**StocksSymbolProfileGet**](docs/STOCKAPI.md#stockssymbolprofileget) | **Get** /stocks/{symbol}/profile | Access essential company details and current valuation metrics to power investment platforms, populate company directories, or enhance financial analysis tools with up-to-date market data.
*STOCKAPI* | [**StocksSymbolRankingsGet**](docs/STOCKAPI.md#stockssymbolrankingsget) | **Get** /stocks/{symbol}/rankings | Proprietary scoring and ranking systems that assess company quality, valuation, and performance for powering data-driven investment platforms.
*STOCKAPI* | [**StocksSymbolSegmentGet**](docs/STOCKAPI.md#stockssymbolsegmentget) | **Get** /stocks/{symbol}/segment | Gain insights into a company&#39;s revenue breakdown by product and geography to build detailed financial visualizations, enhance stock research platforms, or create data-driven market analysis tools
*STOCKAPI* | [**StocksSymbolValuationsGet**](docs/STOCKAPI.md#stockssymbolvaluationsget) | **Get** /stocks/{symbol}/valuations | A deep dataset of historical valuation metrics to support investors and entrepreneurs in the development of data-driven investment platforms.


## Documentation For Models

 - [DownloadIdGet200Response](docs/DownloadIdGet200Response.md)
 - [DownloadListGet200ResponseInner](docs/DownloadListGet200ResponseInner.md)
 - [DownloadListGet403Response](docs/DownloadListGet403Response.md)
 - [DownloadListGet404Response](docs/DownloadListGet404Response.md)
 - [DownloadListGet500Response](docs/DownloadListGet500Response.md)
 - [EtfEtfBasicInformation](docs/EtfEtfBasicInformation.md)
 - [EtfEtfDividends](docs/EtfEtfDividends.md)
 - [EtfEtfFundamental](docs/EtfEtfFundamental.md)
 - [EtfEtfKeyStatistics](docs/EtfEtfKeyStatistics.md)
 - [EtfEtfPortfolioHoldings](docs/EtfEtfPortfolioHoldings.md)
 - [EtfEtfSectorBreakdowns](docs/EtfEtfSectorBreakdowns.md)
 - [EtfSymbolGet200Response](docs/EtfSymbolGet200Response.md)
 - [FundamentalsINOREITNODIRECT](docs/FundamentalsINOREITNODIRECT.md)
 - [FundamentalsINOREITNODIRECTBalanceSheet](docs/FundamentalsINOREITNODIRECTBalanceSheet.md)
 - [FundamentalsINOREITNODIRECTBasicInformation](docs/FundamentalsINOREITNODIRECTBasicInformation.md)
 - [FundamentalsINOREITNODIRECTCashflowStatement](docs/FundamentalsINOREITNODIRECTCashflowStatement.md)
 - [FundamentalsINOREITNODIRECTIncomeStatement](docs/FundamentalsINOREITNODIRECTIncomeStatement.md)
 - [FundamentalsIREITNODIRECT](docs/FundamentalsIREITNODIRECT.md)
 - [FundamentalsIREITNODIRECTBalanceSheet](docs/FundamentalsIREITNODIRECTBalanceSheet.md)
 - [FundamentalsIREITNODIRECTBasicInformation](docs/FundamentalsIREITNODIRECTBasicInformation.md)
 - [FundamentalsIREITNODIRECTCashflowStatement](docs/FundamentalsIREITNODIRECTCashflowStatement.md)
 - [FundamentalsIREITNODIRECTIncomeStatement](docs/FundamentalsIREITNODIRECTIncomeStatement.md)
 - [FundamentalsNNOREITDIRECT](docs/FundamentalsNNOREITDIRECT.md)
 - [FundamentalsNNOREITDIRECTBalanceSheet](docs/FundamentalsNNOREITDIRECTBalanceSheet.md)
 - [FundamentalsNNOREITDIRECTBasicInformation](docs/FundamentalsNNOREITDIRECTBasicInformation.md)
 - [FundamentalsNNOREITDIRECTCashflowStatement](docs/FundamentalsNNOREITDIRECTCashflowStatement.md)
 - [FundamentalsNNOREITDIRECTIncomeStatement](docs/FundamentalsNNOREITDIRECTIncomeStatement.md)
 - [FundamentalsNNOREITNODIRECT](docs/FundamentalsNNOREITNODIRECT.md)
 - [FundamentalsNNOREITNODIRECTBalanceSheet](docs/FundamentalsNNOREITNODIRECTBalanceSheet.md)
 - [FundamentalsNNOREITNODIRECTBasicInformation](docs/FundamentalsNNOREITNODIRECTBasicInformation.md)
 - [FundamentalsNNOREITNODIRECTCashflowStatement](docs/FundamentalsNNOREITNODIRECTCashflowStatement.md)
 - [FundamentalsNNOREITNODIRECTIncomeStatement](docs/FundamentalsNNOREITNODIRECTIncomeStatement.md)
 - [FundamentalsNREITDIRECT](docs/FundamentalsNREITDIRECT.md)
 - [FundamentalsNREITDIRECTBalanceSheet](docs/FundamentalsNREITDIRECTBalanceSheet.md)
 - [FundamentalsNREITDIRECTBasicInformation](docs/FundamentalsNREITDIRECTBasicInformation.md)
 - [FundamentalsNREITDIRECTCashflowStatement](docs/FundamentalsNREITDIRECTCashflowStatement.md)
 - [FundamentalsNREITDIRECTIncomeStatement](docs/FundamentalsNREITDIRECTIncomeStatement.md)
 - [FundamentalsNREITNODIRECT](docs/FundamentalsNREITNODIRECT.md)
 - [FundamentalsNREITNODIRECTBalanceSheet](docs/FundamentalsNREITNODIRECTBalanceSheet.md)
 - [FundamentalsNREITNODIRECTBasicInformation](docs/FundamentalsNREITNODIRECTBasicInformation.md)
 - [FundamentalsNREITNODIRECTCashflowStatement](docs/FundamentalsNREITNODIRECTCashflowStatement.md)
 - [FundamentalsNREITNODIRECTIncomeStatement](docs/FundamentalsNREITNODIRECTIncomeStatement.md)
 - [GuruTransaction](docs/GuruTransaction.md)
 - [GurusIdGet200Response](docs/GurusIdGet200Response.md)
 - [GurusListGet200Response](docs/GurusListGet200Response.md)
 - [GurusListGet200ResponseDataInner](docs/GurusListGet200ResponseDataInner.md)
 - [InsiderTransaction](docs/InsiderTransaction.md)
 - [StockFundamentalsAnnuallyInner](docs/StockFundamentalsAnnuallyInner.md)
 - [StockFundamentalsBasicInformation](docs/StockFundamentalsBasicInformation.md)
 - [StockFundamentalsTtm](docs/StockFundamentalsTtm.md)
 - [StockProfileBasicInformation](docs/StockProfileBasicInformation.md)
 - [StockProfileDividends](docs/StockProfileDividends.md)
 - [StockProfileFundamental](docs/StockProfileFundamental.md)
 - [StockProfileGeneral](docs/StockProfileGeneral.md)
 - [StockProfileGrowth](docs/StockProfileGrowth.md)
 - [StockProfilePrice](docs/StockProfilePrice.md)
 - [StockProfileProfitability](docs/StockProfileProfitability.md)
 - [StockProfileValuationRatio](docs/StockProfileValuationRatio.md)
 - [StockRankingsBasicInformation](docs/StockRankingsBasicInformation.md)
 - [StockRankingsGuruFocusRankings](docs/StockRankingsGuruFocusRankings.md)
 - [StockSegmentBasicInformation](docs/StockSegmentBasicInformation.md)
 - [StockValuationsAnnuallyInner](docs/StockValuationsAnnuallyInner.md)
 - [StockValuationsBasicInformation](docs/StockValuationsBasicInformation.md)
 - [StockValuationsTtm](docs/StockValuationsTtm.md)
 - [StocksSymbolFundamentalsGet200Response](docs/StocksSymbolFundamentalsGet200Response.md)
 - [StocksSymbolProfileGet200Response](docs/StocksSymbolProfileGet200Response.md)
 - [StocksSymbolRankingsGet200Response](docs/StocksSymbolRankingsGet200Response.md)
 - [StocksSymbolSegmentGet200Response](docs/StocksSymbolSegmentGet200Response.md)
 - [StocksSymbolValuationsGet200Response](docs/StocksSymbolValuationsGet200Response.md)
 - [ValuationsINOREITNODIRECT](docs/ValuationsINOREITNODIRECT.md)
 - [ValuationsINOREITNODIRECTBasicInformation](docs/ValuationsINOREITNODIRECTBasicInformation.md)
 - [ValuationsINOREITNODIRECTPerShareData](docs/ValuationsINOREITNODIRECTPerShareData.md)
 - [ValuationsINOREITNODIRECTRatios](docs/ValuationsINOREITNODIRECTRatios.md)
 - [ValuationsINOREITNODIRECTValuationRatios](docs/ValuationsINOREITNODIRECTValuationRatios.md)
 - [ValuationsINOREITNODIRECTValuationandQuality](docs/ValuationsINOREITNODIRECTValuationandQuality.md)
 - [ValuationsIREITNODIRECT](docs/ValuationsIREITNODIRECT.md)
 - [ValuationsIREITNODIRECTBasicInformation](docs/ValuationsIREITNODIRECTBasicInformation.md)
 - [ValuationsIREITNODIRECTPerShareData](docs/ValuationsIREITNODIRECTPerShareData.md)
 - [ValuationsIREITNODIRECTRatios](docs/ValuationsIREITNODIRECTRatios.md)
 - [ValuationsIREITNODIRECTValuationRatios](docs/ValuationsIREITNODIRECTValuationRatios.md)
 - [ValuationsIREITNODIRECTValuationandQuality](docs/ValuationsIREITNODIRECTValuationandQuality.md)
 - [ValuationsNNOREITDIRECT](docs/ValuationsNNOREITDIRECT.md)
 - [ValuationsNNOREITDIRECTBasicInformation](docs/ValuationsNNOREITDIRECTBasicInformation.md)
 - [ValuationsNNOREITDIRECTPerShareData](docs/ValuationsNNOREITDIRECTPerShareData.md)
 - [ValuationsNNOREITDIRECTRatios](docs/ValuationsNNOREITDIRECTRatios.md)
 - [ValuationsNNOREITDIRECTValuationRatios](docs/ValuationsNNOREITDIRECTValuationRatios.md)
 - [ValuationsNNOREITDIRECTValuationandQuality](docs/ValuationsNNOREITDIRECTValuationandQuality.md)
 - [ValuationsNNOREITNODIRECT](docs/ValuationsNNOREITNODIRECT.md)
 - [ValuationsNNOREITNODIRECTBasicInformation](docs/ValuationsNNOREITNODIRECTBasicInformation.md)
 - [ValuationsNNOREITNODIRECTPerShareData](docs/ValuationsNNOREITNODIRECTPerShareData.md)
 - [ValuationsNNOREITNODIRECTRatios](docs/ValuationsNNOREITNODIRECTRatios.md)
 - [ValuationsNNOREITNODIRECTValuationRatios](docs/ValuationsNNOREITNODIRECTValuationRatios.md)
 - [ValuationsNNOREITNODIRECTValuationandQuality](docs/ValuationsNNOREITNODIRECTValuationandQuality.md)
 - [ValuationsNREITDIRECT](docs/ValuationsNREITDIRECT.md)
 - [ValuationsNREITDIRECTBasicInformation](docs/ValuationsNREITDIRECTBasicInformation.md)
 - [ValuationsNREITDIRECTPerShareData](docs/ValuationsNREITDIRECTPerShareData.md)
 - [ValuationsNREITDIRECTRatios](docs/ValuationsNREITDIRECTRatios.md)
 - [ValuationsNREITDIRECTValuationRatios](docs/ValuationsNREITDIRECTValuationRatios.md)
 - [ValuationsNREITDIRECTValuationandQuality](docs/ValuationsNREITDIRECTValuationandQuality.md)
 - [ValuationsNREITNODIRECT](docs/ValuationsNREITNODIRECT.md)
 - [ValuationsNREITNODIRECTBasicInformation](docs/ValuationsNREITNODIRECTBasicInformation.md)
 - [ValuationsNREITNODIRECTPerShareData](docs/ValuationsNREITNODIRECTPerShareData.md)
 - [ValuationsNREITNODIRECTRatios](docs/ValuationsNREITNODIRECTRatios.md)
 - [ValuationsNREITNODIRECTValuationRatios](docs/ValuationsNREITNODIRECTValuationRatios.md)
 - [ValuationsNREITNODIRECTValuationandQuality](docs/ValuationsNREITNODIRECTValuationandQuality.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ApiKeyAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"ApiKeyAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author




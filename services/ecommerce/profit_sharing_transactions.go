package ecommerce

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

type TransactionsApiService services.Service

// QueryOrderAmount 查询剩余待分金额API
//
// 可调用此接口查询订单剩余待分金额
//
// 错误码列表
// |名称|状态码|描述示例|原因|解决方案
// |-|-|-|-|-|
// |SYSTEM_ERROR|500|系统错误|系统超时|系统异常，请使用相同参数稍后重新调用|
// |PARAM_ERROR|400|商户号未设置|请求参数不符合参数格式|请使用正确的参数重新调用|
// |RATELIMIT_EXCEED|429|商户发起查询的频率过高|接口有频率限制|请降低频率后重试|
// |INVALID_REQUEST|400|非分账订单不支持此查询|请求参数符合参数格式，但不符合业务规则|请根据返回的错误信息确认违反的业务规则|
// |INVALID_REQUEST|400|非本商户的订单不支持查询|请求参数符合参数格式，但不符合业务规则|请根据返回的错误信息确认违反的业务规则|
func (a *TransactionsApiService) QueryOrderAmount(ctx context.Context, req QueryOrderAmountRequest) (resp *QueryOrderAmountResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.TransactionId == nil {
		return nil, nil, fmt.Errorf("field `TransactionId` is required and must be specified in QueryOrderAmountRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "v3/ecommerce/profitsharing/orders/{transaction_id}/amounts"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"transaction_id"+"}", neturl.PathEscape(core.ParameterToString(*req.TransactionId, "")), -1)

	// Make sure All Required Params are properly set

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract QueryOrderAmountResponse from Http Response
	resp = new(QueryOrderAmountResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

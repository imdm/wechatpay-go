package ecommerce

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
	nethttp "net/http"
	neturl "net/url"
)

type ReturnOrdersApiService services.Service

// CreateReturnOrder 请求分账回退API
//
// 如果订单已经分账，在退款时，可以先调此接口，将已分账的资金从分账接收方的账户回退给分账方，再发起退款
// 注意：
// - 分账回退以原分账单为依据，支持多次回退，申请回退总金额不能超过原分账单分给该接收方的金额
// - 此接口采用同步处理模式，即在接收到商户请求后，会实时返回处理结果
// - 对同一笔分账单最多能发起20次分账回退请求
// - 退款和分账回退没有耦合，分账回退可以先于退款请求，也可以后于退款请求
// - 此功能需要接收方在商户平台-交易中心-分账-分账接收设置下，开启同意分账回退后，才能使用
//
// 错误码列表
// 名称|状态码|描述示例|原因|解决方案
// |-|-|-|-|-|
// SYSTEM_ERROR|500|系统错误|系统超时|系统异常，请使用相同参数稍后重新调用
// PARAM_ERROR|400|订单号格式不正确|请求参数不符合参数格式|请使用正确的参数重新调用
// INVALID_REQUEST|400|回退方不存在|请求参数符合参数格式，但不符合业务规则|请根据返回的错误信息确认违反的业务规则
// RATELIMIT_EXCEED|429|商户发起分账回退的频率过高|接口有频率限制|请降低频率后重试
// NO_AUTH|403|回退方未开通分账回退功能|未开通分账权限|请先让回退方开通分账回退功能
func (a *ReturnOrdersApiService) CreateReturnOrder(ctx context.Context, req CreateReturnOrderRequest) (resp *ReturnOrdersEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/ecommerce/profitsharing/returnorders"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract ReturnOrdersEntity from Http Response
	resp = new(ReturnOrdersEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryReturnOrder 查询分账回退结果API
//
// 商户需要核实回退结果，可调用此接口查询回退结果
// 注意：
// - 如果分账回退接口返回状态为处理中，可调用此接口查询回退结果
//
// 错误码列表
// 名称|状态码|描述示例|原因|解决方案
// |-|-|-|-|-|
// SYSTEM_ERROR|500|系统错误|系统超时|系统异常，请使用相同参数稍后重新调用
// PARAM_ERROR|400|商户号未设置|请求参数不符合参数格式|请使用正确的参数重新调用
// RATELIMIT_EXCEED|429|商户发起分账回退查询的频率过高|接口有频率限制|请降低频率后重试
// RESOURCE_NOT_EXISTS|404|记录不存在|分账回退单不存在|请检查请求的单号是否正确
func (a *ReturnOrdersApiService) QueryReturnOrder(ctx context.Context, req QueryReturnOrderRequest) (resp *ReturnOrdersEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutReturnNo == nil {
		return nil, nil, fmt.Errorf("field `OutReturnNo` is required and must be specified in QueryReturnOrderRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/ecommerce/profitsharing/returnorders"
	// Build Path with Path Params
	//localVarPath = strings.Replace(localVarPath, "{"+"out_return_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutReturnNo, "")), -1)

	// Make sure All Required Params are properly set
	if req.OutReturnNo == nil {
		return nil, nil, fmt.Errorf("field `OutReturnNo` is required and must be specified in QueryReturnOrderRequest")
	}
	if req.OutOrderNo == nil {
		return nil, nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in QueryReturnOrderRequest")
	}
	if req.SubMchid == nil {
		return nil, nil, fmt.Errorf("field `SubMchid` is required and must be specified in QueryReturnOrderRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("out_return_no", core.ParameterToString(*req.OutReturnNo, ""))
	localVarQueryParams.Add("sub_mchid", core.ParameterToString(*req.SubMchid, ""))
	localVarQueryParams.Add("out_order_no", core.ParameterToString(*req.OutOrderNo, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract ReturnOrdersEntity from Http Response
	resp = new(ReturnOrdersEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

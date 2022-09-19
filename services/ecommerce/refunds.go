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

type RefundsApiService services.Service

// Create 退款申请
//
// # 应用场景
// 当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，微信支付将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。
//
// 注意：
// 1、交易时间超过一年的订单无法提交退款
// 2、微信支付退款支持单笔交易分多次退款，多次退款需要提交原支付订单的商户订单号和设置不同的退款单号。申请退款总金额不能超过订单金额。 一笔退款失败后重新提交，请不要更换退款单号，请使用原商户退款单号
// 3、请求频率限制：150qps，即每秒钟正常的申请退款请求次数不超过150次
//
//	错误或无效请求频率限制：6qps，即每秒钟异常或错误的退款申请请求不超过6次
//
// 4、每个支付订单的部分退款次数不能超过50次
// 5、如果同一个用户有多笔退款，建议分不同批次进行退款，避免并发退款导致退款失败
// 6、申请退款接口的返回仅代表业务的受理情况，具体退款是否成功，需要通过退款查询接口获取结果
//
// # 错误码
// |名称|描述|原因|解决方案|
// |-|-|-|-|
// |SYSTEM_ERROR|接口返回错误|系统超时等|请不要更换商户退款单号，请使用相同参数再次调用API。|
// |USER_ACCOUNT_ABNORMAL|退款请求失败|用户帐号注销|此状态代表退款申请失败，商户可自行处理退款。|
// |NOT_ENOUGH|余额不足|商户可用退款余额不足|此状态代表退款申请失败，商户可根据具体的错误提示做相应的处理。|
// |PARAM_ERROR|参数错误|请求参数未按指引进行填写|请求参数错误，请重新检查再调用退款申请|
// |MCH_NOT_EXISTS|MCHID不存在|参数中缺少MCHID|请检查MCHID是否正确|
// |RESOURCE_NOT_EXISTS|订单号不存在|缺少有效的订单号|请检查你的订单号是否正确且是否已支付，未支付的订单不能发起退款|
// |SIGN_ERROR|签名错误|参数签名结果不正确|请检查签名参数和方法是否都符合签名算法要求|
// |FREQUENCY_LIMITED|频率限制|2个月之前的订单申请退款有频率限制|该笔退款未受理，请降低频率后重试|
// |INVALID_REQUEST|请求参数符合参数格式，但不符合业务规则|不符合业务规则|此状态代表退款申请失败，商户可根据具体的错误提示做相应的处理。|
// |NO_AUTH|没有退款权限|没有此单的退款权限|此状态代表退款申请失败，请检查是否有退这笔订单的权限|
func (a *RefundsApiService) Create(ctx context.Context, req CreateRefundRequest) (resp *Refund, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/ecommerce/refunds/apply"
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

	// Extract Refund from Http Response
	resp = new(Refund)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryByOutRefundNo 查询单笔退款（通过商户退款单号）
//
// # 应用场景
// 提交退款申请后，通过调用该接口查询退款状态。退款有一定延时，建议查询退款状态在提交退款申请后1分钟发起，一般来说零钱支付的退款5分钟内到账，银行卡支付的退款1-3个工作日到账。
// # 错误码
// |名称|描述|原因|解决方案|
// |-|-|-|-|
// |SYSTEM_ERROR|接口返回错误|系统超时|请尝试再次调用API。|
// |RESOURCE_NOT_EXISTS|退款单查询失败|退款单号错误或退款单状态不正确|请检查退款单号是否有误以及订单状态是否正确，如：未支付、已支付未退款|
// |PARAM_ERROR|参数错误|请求参数未按指引进行填写|请求参数错误，请检查参数再调用退款查询|
// |MCH_NOT_EXISTS|MCHID不存在|参数中缺少MCHID|请检查MCHID是否正确|
// |SIGN_ERROR|签名错误|参数签名结果不正确|请检查签名参数和方法是否都符合签名算法要求|
func (a *RefundsApiService) QueryByOutRefundNo(ctx context.Context, req QueryByOutRefundNoRequest) (resp *QueryRefundResp, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutRefundNo == nil {
		return nil, nil, fmt.Errorf("field `OutRefundNo` is required and must be specified in QueryByOutRefundNoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/ecommerce/refunds/out-refund-no/{out_refund_no}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"out_refund_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutRefundNo, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	if req.SubMchid != nil {
		localVarQueryParams.Add("sub_mchid", core.ParameterToString(*req.SubMchid, ""))
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract Refund from Http Response
	resp = new(QueryRefundResp)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

package ecommerce

import (
	"encoding/json"
	"fmt"
	"time"
)

// AddReceiverRequest
type AddReceiverRequest struct {
	// 微信分配的公众账号ID
	Appid *string       `json:"appid"`
	Type  *ReceiverType `json:"type"`
	// 类型是MERCHANT_ID时，是商户号 类型是PERSONAL_OPENID时，是个人openid
	Account *string `json:"account"`
	// 分账接收方的名称，当type为MERCHANT_ID时，接收方名称是商户全称。
	Name *string `json:"name,omitempty"`
	// 分账接收方类型是MERCHANT_ID时，是商户全称（必传），当商户是小微商户或个体户时，是开户人姓名 分账接收方类型是PERSONAL_OPENID时，是个人姓名（选传，传则校验）  1、此字段需要加密，的加密方法详见：敏感信息加密说明 2、使用微信支付平台证书中的公钥 3、使用RSAES-OAEP算法进行加密 4、将请求中HTTP头部的Wechatpay-Serial设置为证书序列号
	EncryptedName *string `json:"encrypted_name,omitempty" encryption:"EM_APIV3"`
	// 子商户与接收方的关系。 本字段值为枚举：
	// SUPPLIER：供应商
	// DISTRIBUTOR：分销商
	// SERVICE_PROVIDER：服务商
	// PLATFORM：平台
	// OTHERS：其他
	RelationType *ReceiverRelationType `json:"relation_type"`
}

func (o AddReceiverRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Account == nil {
		return nil, fmt.Errorf("field `Account` is required and must be specified in AddReceiverRequest")
	}
	toSerialize["account"] = o.Account

	if o.Appid == nil {
		return nil, fmt.Errorf("field `Appid` is required and must be specified in AddReceiverRequest")
	}
	toSerialize["appid"] = o.Appid

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.EncryptedName != nil {
		toSerialize["encrypted_name"] = o.EncryptedName
	}

	if o.RelationType == nil {
		return nil, fmt.Errorf("field `RelationType` is required and must be specified in AddReceiverRequest")
	}
	toSerialize["relation_type"] = o.RelationType

	if o.Type == nil {
		return nil, fmt.Errorf("field `Type` is required and must be specified in AddReceiverRequest")
	}
	toSerialize["type"] = o.Type
	return json.Marshal(toSerialize)
}

func (o AddReceiverRequest) String() string {
	var ret string
	if o.Account == nil {
		ret += "Account:<nil>, "
	} else {
		ret += fmt.Sprintf("Account:%v, ", *o.Account)
	}

	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}

	if o.Name == nil {
		ret += "Name:<nil>, "
	} else {
		ret += fmt.Sprintf("Name:%v, ", *o.Name)
	}

	if o.EncryptedName == nil {
		ret += "EncryptedName:<nil>, "
	} else {
		ret += fmt.Sprintf("EncryptedName:%v, ", *o.EncryptedName)
	}

	if o.RelationType == nil {
		ret += "RelationType:<nil>, "
	} else {
		ret += fmt.Sprintf("RelationType:%v, ", *o.RelationType)
	}

	if o.Type == nil {
		ret += "Type:<nil>"
	} else {
		ret += fmt.Sprintf("Type:%v", *o.Type)
	}

	return fmt.Sprintf("AddReceiverRequest{%s}", ret)
}

func (o AddReceiverRequest) Clone() *AddReceiverRequest {
	ret := AddReceiverRequest{}

	if o.Account != nil {
		ret.Account = new(string)
		*ret.Account = *o.Account
	}

	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}

	if o.Name != nil {
		ret.Name = new(string)
		*ret.Name = *o.Name
	}

	if o.EncryptedName != nil {
		ret.EncryptedName = new(string)
		*ret.EncryptedName = *o.EncryptedName
	}

	if o.RelationType != nil {
		ret.RelationType = new(ReceiverRelationType)
		*ret.RelationType = *o.RelationType
	}

	if o.Type != nil {
		ret.Type = new(ReceiverType)
		*ret.Type = *o.Type
	}

	return &ret
}

// AddReceiverResponse
type AddReceiverResponse struct {
	// 参考请求参数
	Account *string `json:"account"`
	// 参考请求参数  * `MERCHANT_ID` - 商户号，  * `PERSONAL_OPENID` - 个人openid（由父商户APPID转换得到）
	Type *ReceiverType `json:"type"`
}

func (o AddReceiverResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Account == nil {
		return nil, fmt.Errorf("field `Account` is required and must be specified in AddReceiverResponse")
	}
	toSerialize["account"] = o.Account

	if o.Type == nil {
		return nil, fmt.Errorf("field `Type` is required and must be specified in AddReceiverResponse")
	}
	toSerialize["type"] = o.Type
	return json.Marshal(toSerialize)
}

func (o AddReceiverResponse) String() string {
	var ret string
	if o.Account == nil {
		ret += "Account:<nil>, "
	} else {
		ret += fmt.Sprintf("Account:%v, ", *o.Account)
	}

	if o.Type == nil {
		ret += "Type:<nil>"
	} else {
		ret += fmt.Sprintf("Type:%v", *o.Type)
	}

	return fmt.Sprintf("AddReceiverResponse{%s}", ret)
}

func (o AddReceiverResponse) Clone() *AddReceiverResponse {
	ret := AddReceiverResponse{}

	if o.Account != nil {
		ret.Account = new(string)
		*ret.Account = *o.Account
	}

	if o.Type != nil {
		ret.Type = new(ReceiverType)
		*ret.Type = *o.Type
	}

	return &ret
}

// CreateOrderReceiver
type CreateOrderReceiver struct {
	// 1、类型是MERCHANT_ID时，是商户号（mch_id或者sub_mch_id） 2、类型是PERSONAL_OPENID时，是个人openid
	Account *string `json:"receiver_account"`
	// 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额
	Amount *int64 `json:"amount"`
	// 分账的原因描述，分账账单中需要体现
	Description *string `json:"description"`
	// 可选项，在接收方类型为个人的时可选填，若有值，会检查与 receiver_name 是否实名匹配，不匹配会拒绝分账请求 1、分账接收方类型是PERSONAL_OPENID或PERSONAL_SUB_OPENID时，是个人姓名的密文（选传，传则校验） 此字段的加密的方式为：敏感信息加密说明 2、使用微信支付平台证书中的公钥 3、使用RSAES-OAEP算法进行加密 4、将请求中HTTP头部的Wechatpay-Serial设置为证书序列号
	Name *string `json:"receiver_name,omitempty" encryption:"EM_APIV3"`
	// 1、MERCHANT_ID：商户号 2、PERSONAL_OPENID：个人openid（由父商户APPID转换得到）
	Type *string `json:"type"`
}

func (o CreateOrderReceiver) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Account == nil {
		return nil, fmt.Errorf("field `Account` is required and must be specified in CreateOrderReceiver")
	}
	toSerialize["receiver_account"] = o.Account

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in CreateOrderReceiver")
	}
	toSerialize["amount"] = o.Amount

	if o.Description == nil {
		return nil, fmt.Errorf("field `Description` is required and must be specified in CreateOrderReceiver")
	}
	toSerialize["description"] = o.Description

	if o.Name != nil {
		toSerialize["receiver_name"] = o.Name
	}

	if o.Type == nil {
		return nil, fmt.Errorf("field `Type` is required and must be specified in CreateOrderReceiver")
	}
	toSerialize["type"] = o.Type
	return json.Marshal(toSerialize)
}

func (o CreateOrderReceiver) String() string {
	var ret string
	if o.Account == nil {
		ret += "Account:<nil>, "
	} else {
		ret += fmt.Sprintf("Account:%v, ", *o.Account)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>, "
	} else {
		ret += fmt.Sprintf("Amount:%v, ", *o.Amount)
	}

	if o.Description == nil {
		ret += "Description:<nil>, "
	} else {
		ret += fmt.Sprintf("Description:%v, ", *o.Description)
	}

	if o.Name == nil {
		ret += "EncryptedName:<nil>, "
	} else {
		ret += fmt.Sprintf("EncryptedName:%v, ", *o.Name)
	}

	if o.Type == nil {
		ret += "Type:<nil>"
	} else {
		ret += fmt.Sprintf("Type:%v", *o.Type)
	}

	return fmt.Sprintf("CreateOrderReceiver{%s}", ret)
}

func (o CreateOrderReceiver) Clone() *CreateOrderReceiver {
	ret := CreateOrderReceiver{}

	if o.Account != nil {
		ret.Account = new(string)
		*ret.Account = *o.Account
	}

	if o.Amount != nil {
		ret.Amount = new(int64)
		*ret.Amount = *o.Amount
	}

	if o.Description != nil {
		ret.Description = new(string)
		*ret.Description = *o.Description
	}

	if o.Name != nil {
		ret.Name = new(string)
		*ret.Name = *o.Name
	}

	if o.Type != nil {
		ret.Type = new(string)
		*ret.Type = *o.Type
	}

	return &ret
}

// CreateOrderRequest
type CreateOrderRequest struct {
	// 电商平台的appid（公众号APPID或者小程序APPID）
	Appid *string `json:"appid"`
	// 服务商系统内部的分账单号，在服务商系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutOrderNo *string `json:"out_order_no"`
	// 分账接收方列表，可以设置出资商户作为分账接受方，最多可有50个分账接收方
	Receivers []CreateOrderReceiver `json:"receivers"`
	// 分账出资的电商平台二级商户，填写微信支付分配的商户号
	SubMchid *string `json:"sub_mchid"`
	// 微信支付订单号
	TransactionId *string `json:"transaction_id"`
	// 是否完成分账
	//1、如果为true，该笔订单剩余未分账的金额会解冻回电商平台二级商户；
	//2、如果为false，该笔订单剩余未分账的金额不会解冻回电商平台二级商户，可以对该笔订单再次进行分账。
	Finish *bool `json:"finish"`
}

func (o CreateOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Appid == nil {
		return nil, fmt.Errorf("field `Appid` is required and must be specified in CreateOrderRequest")
	}
	toSerialize["appid"] = o.Appid

	if o.OutOrderNo == nil {
		return nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in CreateOrderRequest")
	}
	toSerialize["out_order_no"] = o.OutOrderNo

	if o.Receivers == nil {
		return nil, fmt.Errorf("field `Receivers` is required and must be specified in CreateOrderRequest")
	}
	toSerialize["receivers"] = o.Receivers

	if o.SubMchid == nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in CreateOrderRequest")
	}
	toSerialize["sub_mchid"] = o.SubMchid

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in CreateOrderRequest")
	}
	toSerialize["transaction_id"] = o.TransactionId

	if o.Finish == nil {
		return nil, fmt.Errorf("field `Finish` is required and must be specified in CreateOrderRequest")
	}
	toSerialize["finish"] = o.Finish
	return json.Marshal(toSerialize)
}

func (o CreateOrderRequest) String() string {
	var ret string
	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}

	if o.OutOrderNo == nil {
		ret += "OutOrderNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutOrderNo:%v, ", *o.OutOrderNo)
	}

	ret += fmt.Sprintf("Receivers:%v, ", o.Receivers)

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}

	if o.Finish == nil {
		ret += "Finish:<nil>"
	} else {
		ret += fmt.Sprintf("Finish:%v", *o.Finish)
	}

	return fmt.Sprintf("CreateOrderRequest{%s}", ret)
}

func (o CreateOrderRequest) Clone() *CreateOrderRequest {
	ret := CreateOrderRequest{}

	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}

	if o.OutOrderNo != nil {
		ret.OutOrderNo = new(string)
		*ret.OutOrderNo = *o.OutOrderNo
	}

	if o.Receivers != nil {
		ret.Receivers = make([]CreateOrderReceiver, len(o.Receivers))
		for i, item := range o.Receivers {
			ret.Receivers[i] = *item.Clone()
		}
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.Finish != nil {
		ret.Finish = new(bool)
		*ret.Finish = *o.Finish
	}

	return &ret
}

// CreateReturnOrderRequest
type CreateReturnOrderRequest struct {
	// 分账回退的接收商户，对应原分账出资的分账方商户，填写微信支付分配的商户号。
	SubMchid *string `json:"sub_mchid"`
	// 微信分账单号，微信系统返回的唯一标识。微信分账单号和商户分账单号二选一填写
	OrderId *string `json:"order_id,omitempty"`
	// 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。微信分账单号和商户分账单号二选一填写 取值范围：[0-9a-zA-Z_*@-]
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// 此回退单号是商户在自己后台生成的一个新的回退单号，在商户后台唯一
	OutReturnNo *string `json:"out_return_no"`
	// 分账回退的出资商户，只能对原分账请求中成功分给商户接收方进行回退。回退商户号只能填写原分账请求中接收分账的商户号
	ReturnMchid *string `json:"return_mchid"`
	// 需要从分账接收方回退的金额，单位为分，只能为整数，不能超过原始分账单分出给该接收方的金额
	Amount *int64 `json:"amount"`
	// 分账回退的原因描述
	Description *string `json:"description"`
}

func (o CreateReturnOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in CreateReturnOrderRequest")
	}
	toSerialize["amount"] = o.Amount

	if o.Description == nil {
		return nil, fmt.Errorf("field `Description` is required and must be specified in CreateReturnOrderRequest")
	}
	toSerialize["description"] = o.Description

	if o.OrderId != nil {
		toSerialize["order_id"] = o.OrderId
	}

	if o.OutOrderNo != nil {
		toSerialize["out_order_no"] = o.OutOrderNo
	}

	if o.OutReturnNo == nil {
		return nil, fmt.Errorf("field `OutReturnNo` is required and must be specified in CreateReturnOrderRequest")
	}
	toSerialize["out_return_no"] = o.OutReturnNo

	if o.ReturnMchid == nil {
		return nil, fmt.Errorf("field `ReturnMchid` is required and must be specified in CreateReturnOrderRequest")
	}
	toSerialize["return_mchid"] = o.ReturnMchid

	if o.SubMchid == nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in CreateReturnOrderRequest")
	}
	toSerialize["sub_mchid"] = o.SubMchid
	return json.Marshal(toSerialize)
}

func (o CreateReturnOrderRequest) String() string {
	var ret string
	if o.Amount == nil {
		ret += "Amount:<nil>, "
	} else {
		ret += fmt.Sprintf("Amount:%v, ", *o.Amount)
	}

	if o.Description == nil {
		ret += "Description:<nil>, "
	} else {
		ret += fmt.Sprintf("Description:%v, ", *o.Description)
	}

	if o.OrderId == nil {
		ret += "OrderId:<nil>, "
	} else {
		ret += fmt.Sprintf("OrderId:%v, ", *o.OrderId)
	}

	if o.OutOrderNo == nil {
		ret += "OutOrderNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutOrderNo:%v, ", *o.OutOrderNo)
	}

	if o.OutReturnNo == nil {
		ret += "OutReturnNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutReturnNo:%v, ", *o.OutReturnNo)
	}

	if o.ReturnMchid == nil {
		ret += "ReturnMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("ReturnMchid:%v, ", *o.ReturnMchid)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>"
	} else {
		ret += fmt.Sprintf("SubMchid:%v", *o.SubMchid)
	}

	return fmt.Sprintf("CreateReturnOrderRequest{%s}", ret)
}

func (o CreateReturnOrderRequest) Clone() *CreateReturnOrderRequest {
	ret := CreateReturnOrderRequest{}

	if o.Amount != nil {
		ret.Amount = new(int64)
		*ret.Amount = *o.Amount
	}

	if o.Description != nil {
		ret.Description = new(string)
		*ret.Description = *o.Description
	}

	if o.OrderId != nil {
		ret.OrderId = new(string)
		*ret.OrderId = *o.OrderId
	}

	if o.OutOrderNo != nil {
		ret.OutOrderNo = new(string)
		*ret.OutOrderNo = *o.OutOrderNo
	}

	if o.OutReturnNo != nil {
		ret.OutReturnNo = new(string)
		*ret.OutReturnNo = *o.OutReturnNo
	}

	if o.ReturnMchid != nil {
		ret.ReturnMchid = new(string)
		*ret.ReturnMchid = *o.ReturnMchid
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	return &ret
}

// DeleteReceiverRequest
type DeleteReceiverRequest struct {
	// 微信分配的公众账号ID
	Appid *string `json:"appid"`
	// 分账接收方的类型，枚举值： MERCHANT_ID：商户 PERSONAL_OPENID：个人
	Type *ReceiverType `json:"type"`
	// 类型是MERCHANT_ID时，是商户号 类型是PERSONAL_OPENID时，是个人openid 类型是PERSONAL_SUB_OPENID时，是个人sub_openid
	Account *string `json:"account"`
}

func (o DeleteReceiverRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Account == nil {
		return nil, fmt.Errorf("field `Account` is required and must be specified in DeleteReceiverRequest")
	}
	toSerialize["account"] = o.Account

	if o.Appid == nil {
		return nil, fmt.Errorf("field `Appid` is required and must be specified in DeleteReceiverRequest")
	}
	toSerialize["appid"] = o.Appid

	if o.Type == nil {
		return nil, fmt.Errorf("field `Type` is required and must be specified in DeleteReceiverRequest")
	}
	toSerialize["type"] = o.Type
	return json.Marshal(toSerialize)
}

func (o DeleteReceiverRequest) String() string {
	var ret string
	if o.Account == nil {
		ret += "Account:<nil>, "
	} else {
		ret += fmt.Sprintf("Account:%v, ", *o.Account)
	}

	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}

	if o.Type == nil {
		ret += "Type:<nil>"
	} else {
		ret += fmt.Sprintf("Type:%v", *o.Type)
	}

	return fmt.Sprintf("DeleteReceiverRequest{%s}", ret)
}

func (o DeleteReceiverRequest) Clone() *DeleteReceiverRequest {
	ret := DeleteReceiverRequest{}

	if o.Account != nil {
		ret.Account = new(string)
		*ret.Account = *o.Account
	}

	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}

	if o.Type != nil {
		ret.Type = new(ReceiverType)
		*ret.Type = *o.Type
	}

	return &ret
}

// DeleteReceiverResponse
type DeleteReceiverResponse struct {
	// 参考请求参数
	Account *string `json:"account"`
	// 参考请求参数  * `MERCHANT_ID` - 商户号，  * `PERSONAL_OPENID` - 个人openid（由父商户APPID转换得到），  * `PERSONAL_SUB_OPENID` - 个人sub_openid（由子商户APPID转换得到）（直连商户不需要，服务商需要），
	Type *ReceiverType `json:"type"`
}

func (o DeleteReceiverResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Account == nil {
		return nil, fmt.Errorf("field `Account` is required and must be specified in DeleteReceiverResponse")
	}
	toSerialize["account"] = o.Account

	if o.Type == nil {
		return nil, fmt.Errorf("field `Type` is required and must be specified in DeleteReceiverResponse")
	}
	toSerialize["type"] = o.Type
	return json.Marshal(toSerialize)
}

func (o DeleteReceiverResponse) String() string {
	var ret string
	if o.Account == nil {
		ret += "Account:<nil>, "
	} else {
		ret += fmt.Sprintf("Account:%v, ", *o.Account)
	}

	if o.Type == nil {
		ret += "Type:<nil>"
	} else {
		ret += fmt.Sprintf("Type:%v", *o.Type)
	}

	return fmt.Sprintf("DeleteReceiverResponse{%s}", ret)
}

func (o DeleteReceiverResponse) Clone() *DeleteReceiverResponse {
	ret := DeleteReceiverResponse{}

	if o.Account != nil {
		ret.Account = new(string)
		*ret.Account = *o.Account
	}

	if o.Type != nil {
		ret.Type = new(ReceiverType)
		*ret.Type = *o.Type
	}

	return &ret
}

// DetailFailReason   * `ACCOUNT_ABNORMAL` - 分账接收账户异常，  * `NO_RELATION` - 分账关系已解除，  * `RECEIVER_HIGH_RISK` - 高风险接收方，  * `RECEIVER_REAL_NAME_NOT_VERIFIED` - 接收方未实名，  * `NO_AUTH` - 分账权限已解除，
type DetailFailReason string

func (e DetailFailReason) Ptr() *DetailFailReason {
	return &e
}

// Enums of DetailFailReason
const (
	DETAILFAILREASON_ACCOUNT_ABNORMAL                DetailFailReason = "ACCOUNT_ABNORMAL"
	DETAILFAILREASON_NO_RELATION                     DetailFailReason = "NO_RELATION"
	DETAILFAILREASON_RECEIVER_HIGH_RISK              DetailFailReason = "RECEIVER_HIGH_RISK"
	DETAILFAILREASON_RECEIVER_REAL_NAME_NOT_VERIFIED DetailFailReason = "RECEIVER_REAL_NAME_NOT_VERIFIED"
	DETAILFAILREASON_NO_AUTH                         DetailFailReason = "NO_AUTH"
)

// DetailStatus   * `PENDING` - 待分账，  * `SUCCESS` - 分账成功，  * `CLOSED` - 已关闭，
type DetailStatus string

func (e DetailStatus) Ptr() *DetailStatus {
	return &e
}

// Enums of DetailStatus
const (
	DETAILSTATUS_PENDING DetailStatus = "PENDING"
	DETAILSTATUS_SUCCESS DetailStatus = "SUCCESS"
	DETAILSTATUS_CLOSED  DetailStatus = "CLOSED"
)

// OrderReceiverDetail
type OrderReceiverDetail struct {
	// 仅分账接收方类型为MERCHANT_ID时，填写微信支付分配的商户号
	ReceiverMchID *string `json:"receiver_mchid"`
	// 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额
	Amount *int64 `json:"amount"`
	// 分账的原因描述，分账账单中需要体现
	Description *string `json:"description"`
	// 枚举值： 1、PENDING：待分账 2、SUCCESS：分账成功 3、CLOSED：已关闭  * `PENDING` - 待分账，  * `SUCCESS` - 分账成功，  * `CLOSED` - 已关闭，
	Result *DetailStatus `json:"result"`
	// 分账完成时间，遵循RFC3339标准格式
	FinishTime *time.Time `json:"finish_time"`
	// 分账失败原因。包含以下枚举值： 1、ACCOUNT_ABNORMAL : 分账接收账户异常 2、NO_RELATION : 分账关系已解除 3、RECEIVER_HIGH_RISK : 高风险接收方 4、RECEIVER_REAL_NAME_NOT_VERIFIED : 接收方未实名 5、NO_AUTH : 分账权限已解除  * `ACCOUNT_ABNORMAL` - 分账接收账户异常，  * `NO_RELATION` - 分账关系已解除，  * `RECEIVER_HIGH_RISK` - 高风险接收方，  * `RECEIVER_REAL_NAME_NOT_VERIFIED` - 接收方未实名，  * `NO_AUTH` - 分账权限已解除，
	FailReason *DetailFailReason `json:"fail_reason,omitempty"`
	// 1、MERCHANT_ID：商户号 2、PERSONAL_OPENID：个人openid（由父商户APPID转换得到） 3、PERSONAL_SUB_OPENID: 个人sub_openid（由子商户APPID转换得到）  * `MERCHANT_ID` - 商户号，  * `PERSONAL_OPENID` - 个人openid（由父商户APPID转换得到），  * `PERSONAL_SUB_OPENID` - 个人sub_openid（由子商户APPID转换得到）（直连商户不需要，服务商需要），
	Type *ReceiverType `json:"type"`
	// 1、类型是MERCHANT_ID时，是商户号 2、类型是PERSONAL_OPENID时，是个人openid 3、类型是PERSONAL_SUB_OPENID时，是个人sub_openid
	ReceiverAccount *string `json:"receiver_account"`
	//// 分账创建时间，遵循RFC3339标准格式
	//CreateTime *time.Time `json:"create_time"`
	// 微信分账明细单号，每笔分账业务执行的明细单号，可与资金账单对账使用
	DetailId *string `json:"detail_id"`
}

func (o OrderReceiverDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.ReceiverAccount == nil {
		return nil, fmt.Errorf("field `ReceiverAccount` is required and must be specified in OrderReceiverDetail")
	}
	toSerialize["receiver_account"] = o.ReceiverAccount

	if o.ReceiverMchID == nil {
		return nil, fmt.Errorf("field `ReceiverMchID` is required and must be specified in OrderReceiverDetail")
	}
	toSerialize["receiver_mchid"] = o.ReceiverMchID

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in OrderReceiverDetail")
	}
	toSerialize["amount"] = o.Amount

	//if o.CreateTime == nil {
	//	return nil, fmt.Errorf("field `CreateTime` is required and must be specified in OrderReceiverDetail")
	//}
	//toSerialize["create_time"] = o.CreateTime.Format(time.RFC3339)

	if o.Description == nil {
		return nil, fmt.Errorf("field `Description` is required and must be specified in OrderReceiverDetail")
	}
	toSerialize["description"] = o.Description

	if o.DetailId == nil {
		return nil, fmt.Errorf("field `DetailId` is required and must be specified in OrderReceiverDetail")
	}
	toSerialize["detail_id"] = o.DetailId

	if o.FailReason != nil {
		toSerialize["fail_reason"] = o.FailReason
	}

	if o.FinishTime == nil {
		return nil, fmt.Errorf("field `FinishTime` is required and must be specified in OrderReceiverDetail")
	}
	toSerialize["finish_time"] = o.FinishTime.Format(time.RFC3339)

	if o.Result == nil {
		return nil, fmt.Errorf("field `Result` is required and must be specified in OrderReceiverDetail")
	}
	toSerialize["result"] = o.Result

	if o.Type == nil {
		return nil, fmt.Errorf("field `Type` is required and must be specified in OrderReceiverDetail")
	}
	toSerialize["type"] = o.Type
	return json.Marshal(toSerialize)
}

func (o OrderReceiverDetail) String() string {
	var ret string
	if o.ReceiverAccount == nil {
		ret += "ReceiverAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("ReceiverAccount:%v, ", *o.ReceiverAccount)
	}

	if o.ReceiverMchID == nil {
		ret += "ReceiverMchID:<nil>, "
	} else {
		ret += fmt.Sprintf("ReceiverMchID:%v, ", *o.ReceiverMchID)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>, "
	} else {
		ret += fmt.Sprintf("Amount:%v, ", *o.Amount)
	}

	//if o.CreateTime == nil {
	//	ret += "CreateTime:<nil>, "
	//} else {
	//	ret += fmt.Sprintf("CreateTime:%v, ", *o.CreateTime)
	//}

	if o.Description == nil {
		ret += "Description:<nil>, "
	} else {
		ret += fmt.Sprintf("Description:%v, ", *o.Description)
	}

	if o.DetailId == nil {
		ret += "DetailId:<nil>, "
	} else {
		ret += fmt.Sprintf("DetailId:%v, ", *o.DetailId)
	}

	if o.FailReason == nil {
		ret += "FailReason:<nil>, "
	} else {
		ret += fmt.Sprintf("FailReason:%v, ", *o.FailReason)
	}

	if o.FinishTime == nil {
		ret += "FinishTime:<nil>, "
	} else {
		ret += fmt.Sprintf("FinishTime:%v, ", *o.FinishTime)
	}

	if o.Result == nil {
		ret += "Result:<nil>, "
	} else {
		ret += fmt.Sprintf("Result:%v, ", *o.Result)
	}

	if o.Type == nil {
		ret += "Type:<nil>"
	} else {
		ret += fmt.Sprintf("Type:%v", *o.Type)
	}

	return fmt.Sprintf("OrderReceiverDetail{%s}", ret)
}

func (o OrderReceiverDetail) Clone() *OrderReceiverDetail {
	ret := OrderReceiverDetail{}

	if o.ReceiverAccount != nil {
		ret.ReceiverAccount = new(string)
		*ret.ReceiverAccount = *o.ReceiverAccount
	}

	if o.ReceiverMchID != nil {
		ret.ReceiverMchID = new(string)
		*ret.ReceiverMchID = *o.ReceiverMchID
	}

	if o.Amount != nil {
		ret.Amount = new(int64)
		*ret.Amount = *o.Amount
	}

	//if o.CreateTime != nil {
	//	ret.CreateTime = new(time.Time)
	//	*ret.CreateTime = *o.CreateTime
	//}

	if o.Description != nil {
		ret.Description = new(string)
		*ret.Description = *o.Description
	}

	if o.DetailId != nil {
		ret.DetailId = new(string)
		*ret.DetailId = *o.DetailId
	}

	if o.FailReason != nil {
		ret.FailReason = new(DetailFailReason)
		*ret.FailReason = *o.FailReason
	}

	if o.FinishTime != nil {
		ret.FinishTime = new(time.Time)
		*ret.FinishTime = *o.FinishTime
	}

	if o.Result != nil {
		ret.Result = new(DetailStatus)
		*ret.Result = *o.Result
	}

	if o.Type != nil {
		ret.Type = new(ReceiverType)
		*ret.Type = *o.Type
	}

	return &ret
}

// OrderStatus   * `PROCESSING` - 处理中，  * `FINISHED` - 分账完成，
type OrderStatus string

func (e OrderStatus) Ptr() *OrderStatus {
	return &e
}

// Enums of OrderStatus
const (
	ORDERSTATUS_PROCESSING OrderStatus = "PROCESSING"
	ORDERSTATUS_FINISHED   OrderStatus = "FINISHED"
)

// OrdersEntity
type OrdersEntity struct {
	// 微信分账单号，微信系统返回的唯一标识
	OrderId *string `json:"order_id"`
	// 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutOrderNo *string `json:"out_order_no"`
	// 分账接收方列表
	Receivers []OrderReceiverDetail `json:"receivers"`
	// 分账单状态（每个接收方的分账结果请查看receivers中的result字段），枚举值： 1、PROCESSING：处理中 2、FINISHED：分账完成  * `PROCESSING` - 处理中，  * `FINISHED` - 分账完成，
	Status *OrderStatus `json:"status"`
	// 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要）
	SubMchid *string `json:"sub_mchid"`
	// 微信支付订单号
	TransactionId *string `json:"transaction_id"`
}

func (o OrdersEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OrderId == nil {
		return nil, fmt.Errorf("field `OrderId` is required and must be specified in OrdersEntity")
	}
	toSerialize["order_id"] = o.OrderId

	if o.OutOrderNo == nil {
		return nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in OrdersEntity")
	}
	toSerialize["out_order_no"] = o.OutOrderNo

	if o.Receivers != nil {
		return nil, fmt.Errorf("field `Receivers` is required and must be specified in OrdersEntity")
	}
	toSerialize["receivers"] = o.Receivers

	if o.Status == nil {
		return nil, fmt.Errorf("field `Status` is required and must be specified in OrdersEntity")
	}
	toSerialize["status"] = o.Status

	if o.SubMchid != nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in OrdersEntity")
	}
	toSerialize["sub_mchid"] = o.SubMchid

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in OrdersEntity")
	}
	toSerialize["transaction_id"] = o.TransactionId
	return json.Marshal(toSerialize)
}

func (o OrdersEntity) String() string {
	var ret string
	if o.OrderId == nil {
		ret += "OrderId:<nil>, "
	} else {
		ret += fmt.Sprintf("OrderId:%v, ", *o.OrderId)
	}

	if o.OutOrderNo == nil {
		ret += "OutOrderNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutOrderNo:%v, ", *o.OutOrderNo)
	}

	ret += fmt.Sprintf("Receivers:%v, ", o.Receivers)

	if o.Status == nil {
		ret += "Status:<nil>, "
	} else {
		ret += fmt.Sprintf("Status:%v, ", *o.Status)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>"
	} else {
		ret += fmt.Sprintf("TransactionId:%v", *o.TransactionId)
	}

	return fmt.Sprintf("OrdersEntity{%s}", ret)
}

func (o OrdersEntity) Clone() *OrdersEntity {
	ret := OrdersEntity{}

	if o.OrderId != nil {
		ret.OrderId = new(string)
		*ret.OrderId = *o.OrderId
	}

	if o.OutOrderNo != nil {
		ret.OutOrderNo = new(string)
		*ret.OutOrderNo = *o.OutOrderNo
	}

	if o.Receivers != nil {
		ret.Receivers = make([]OrderReceiverDetail, len(o.Receivers))
		for i, item := range o.Receivers {
			ret.Receivers[i] = *item.Clone()
		}
	}

	if o.Status != nil {
		ret.Status = new(OrderStatus)
		*ret.Status = *o.Status
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	return &ret
}

// QueryMerchantRatioRequest
type QueryMerchantRatioRequest struct {
	// 微信支付分配的子商户号，即分账的出资商户号。
	SubMchid *string `json:"sub_mchid"`
}

func (o QueryMerchantRatioRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.SubMchid == nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in QueryMerchantRatioRequest")
	}
	toSerialize["sub_mchid"] = o.SubMchid
	return json.Marshal(toSerialize)
}

func (o QueryMerchantRatioRequest) String() string {
	var ret string
	if o.SubMchid == nil {
		ret += "SubMchid:<nil>"
	} else {
		ret += fmt.Sprintf("SubMchid:%v", *o.SubMchid)
	}

	return fmt.Sprintf("QueryMerchantRatioRequest{%s}", ret)
}

func (o QueryMerchantRatioRequest) Clone() *QueryMerchantRatioRequest {
	ret := QueryMerchantRatioRequest{}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	return &ret
}

// QueryMerchantRatioResponse
type QueryMerchantRatioResponse struct {
	// 子商户允许父商户分账的最大比例，单位万分比，比如2000表示20%
	MaxRatio *int64 `json:"max_ratio"`
	// 参考请求参数
	SubMchid *string `json:"sub_mchid"`
}

func (o QueryMerchantRatioResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.MaxRatio == nil {
		return nil, fmt.Errorf("field `MaxRatio` is required and must be specified in QueryMerchantRatioResponse")
	}
	toSerialize["max_ratio"] = o.MaxRatio

	if o.SubMchid == nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in QueryMerchantRatioResponse")
	}
	toSerialize["sub_mchid"] = o.SubMchid
	return json.Marshal(toSerialize)
}

func (o QueryMerchantRatioResponse) String() string {
	var ret string
	if o.MaxRatio == nil {
		ret += "MaxRatio:<nil>, "
	} else {
		ret += fmt.Sprintf("MaxRatio:%v, ", *o.MaxRatio)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>"
	} else {
		ret += fmt.Sprintf("SubMchid:%v", *o.SubMchid)
	}

	return fmt.Sprintf("QueryMerchantRatioResponse{%s}", ret)
}

func (o QueryMerchantRatioResponse) Clone() *QueryMerchantRatioResponse {
	ret := QueryMerchantRatioResponse{}

	if o.MaxRatio != nil {
		ret.MaxRatio = new(int64)
		*ret.MaxRatio = *o.MaxRatio
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	return &ret
}

// QueryOrderAmountRequest
type QueryOrderAmountRequest struct {
	// 微信支付订单号
	TransactionId *string `json:"transaction_id"`
}

func (o QueryOrderAmountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in QueryOrderAmountRequest")
	}
	toSerialize["transaction_id"] = o.TransactionId
	return json.Marshal(toSerialize)
}

func (o QueryOrderAmountRequest) String() string {
	var ret string
	if o.TransactionId == nil {
		ret += "TransactionId:<nil>"
	} else {
		ret += fmt.Sprintf("TransactionId:%v", *o.TransactionId)
	}

	return fmt.Sprintf("QueryOrderAmountRequest{%s}", ret)
}

func (o QueryOrderAmountRequest) Clone() *QueryOrderAmountRequest {
	ret := QueryOrderAmountRequest{}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	return &ret
}

// QueryOrderAmountResponse
type QueryOrderAmountResponse struct {
	// 微信支付订单号
	TransactionId *string `json:"transaction_id"`
	// 订单剩余待分金额，整数，单元为分
	UnsplitAmount *int64 `json:"unsplit_amount"`
}

func (o QueryOrderAmountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in QueryOrderAmountResponse")
	}
	toSerialize["transaction_id"] = o.TransactionId

	if o.UnsplitAmount == nil {
		return nil, fmt.Errorf("field `UnsplitAmount` is required and must be specified in QueryOrderAmountResponse")
	}
	toSerialize["unsplit_amount"] = o.UnsplitAmount
	return json.Marshal(toSerialize)
}

func (o QueryOrderAmountResponse) String() string {
	var ret string
	if o.TransactionId == nil {
		ret += "TransactionId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}

	if o.UnsplitAmount == nil {
		ret += "UnsplitAmount:<nil>"
	} else {
		ret += fmt.Sprintf("UnsplitAmount:%v", *o.UnsplitAmount)
	}

	return fmt.Sprintf("QueryOrderAmountResponse{%s}", ret)
}

func (o QueryOrderAmountResponse) Clone() *QueryOrderAmountResponse {
	ret := QueryOrderAmountResponse{}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.UnsplitAmount != nil {
		ret.UnsplitAmount = new(int64)
		*ret.UnsplitAmount = *o.UnsplitAmount
	}

	return &ret
}

// QueryOrderRequest
type QueryOrderRequest struct {
	// 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要）
	SubMchid *string `json:"sub_mchid,omitempty"`
	// 微信支付订单号
	TransactionId *string `json:"transaction_id"`
	// 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@ 。 微信分账单号与商户分账单号二选一填写
	OutOrderNo *string `json:"out_order_no"`
}

func (o QueryOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.SubMchid != nil {
		toSerialize["sub_mchid"] = o.SubMchid
	}

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in QueryOrderRequest")
	}
	toSerialize["transaction_id"] = o.TransactionId

	if o.OutOrderNo == nil {
		return nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in QueryOrderRequest")
	}
	toSerialize["out_order_no"] = o.OutOrderNo
	return json.Marshal(toSerialize)
}

func (o QueryOrderRequest) String() string {
	var ret string
	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}

	if o.OutOrderNo == nil {
		ret += "OutOrderNo:<nil>"
	} else {
		ret += fmt.Sprintf("OutOrderNo:%v", *o.OutOrderNo)
	}

	return fmt.Sprintf("QueryOrderRequest{%s}", ret)
}

func (o QueryOrderRequest) Clone() *QueryOrderRequest {
	ret := QueryOrderRequest{}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.OutOrderNo != nil {
		ret.OutOrderNo = new(string)
		*ret.OutOrderNo = *o.OutOrderNo
	}

	return &ret
}

// QueryReturnOrderRequest
type QueryReturnOrderRequest struct {
	// 分账回退的接收商户，对应原分账出资的分账方商户，填写微信支付分配的商户号。（直连商户不需要，服务商需要）
	SubMchid *string `json:"sub_mchid"`
	// 调用回退接口提供的商户系统内部的回退单号。微信支付回退单号与商户回退单号二选一填写
	OutReturnNo *string `json:"out_return_no"`
	// 原发起分账请求时使用的商户系统内部的分账单号。微信分账单号与商户分账单号二选一填写
	OutOrderNo *string `json:"out_order_no"`
}

func (o QueryReturnOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.SubMchid == nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in QueryReturnOrderRequest")
	}
	toSerialize["sub_mchid"] = o.SubMchid

	if o.OutReturnNo == nil {
		return nil, fmt.Errorf("field `OutReturnNo` is required and must be specified in QueryReturnOrderRequest")
	}
	toSerialize["out_return_no"] = o.OutReturnNo

	if o.OutOrderNo == nil {
		return nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in QueryReturnOrderRequest")
	}
	toSerialize["out_order_no"] = o.OutOrderNo
	return json.Marshal(toSerialize)
}

func (o QueryReturnOrderRequest) String() string {
	var ret string
	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.OutReturnNo == nil {
		ret += "OutReturnNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutReturnNo:%v, ", *o.OutReturnNo)
	}

	if o.OutOrderNo == nil {
		ret += "OutOrderNo:<nil>"
	} else {
		ret += fmt.Sprintf("OutOrderNo:%v", *o.OutOrderNo)
	}

	return fmt.Sprintf("QueryReturnOrderRequest{%s}", ret)
}

func (o QueryReturnOrderRequest) Clone() *QueryReturnOrderRequest {
	ret := QueryReturnOrderRequest{}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.OutReturnNo != nil {
		ret.OutReturnNo = new(string)
		*ret.OutReturnNo = *o.OutReturnNo
	}

	if o.OutOrderNo != nil {
		ret.OutOrderNo = new(string)
		*ret.OutOrderNo = *o.OutOrderNo
	}

	return &ret
}

type ReceiverRelationType string

func (e ReceiverRelationType) Ptr() *ReceiverRelationType {
	return &e
}

// Enums of ReceiverRelationType
const (
	RECEIVERRELATIONTYPE_SUPPLIER         ReceiverRelationType = "SUPPLIER"         // 供应商
	RECEIVERRELATIONTYPE_DISTRIBUTOR      ReceiverRelationType = "DISTRIBUTOR"      // 分销商
	RECEIVERRELATIONTYPE_SERVICE_PROVIDER ReceiverRelationType = "SERVICE_PROVIDER" // 服务提供商
	RECEIVERRELATIONTYPE_PLATFORM         ReceiverRelationType = "PLATFORM"         // 平台
	RECEIVERRELATIONTYPE_OTHERS           ReceiverRelationType = "OTHERS"           // 其他
)

// ReceiverType   * `MERCHANT_ID` - 商户号，  * `PERSONAL_OPENID` - 个人openid（由父商户APPID转换得到），  * `PERSONAL_SUB_OPENID` - 个人sub_openid（由子商户APPID转换得到）（直连商户不需要，服务商需要），
type ReceiverType string

func (e ReceiverType) Ptr() *ReceiverType {
	return &e
}

// Enums of ReceiverType
const (
	RECEIVERTYPE_MERCHANT_ID     ReceiverType = "MERCHANT_ID"
	RECEIVERTYPE_PERSONAL_OPENID ReceiverType = "PERSONAL_OPENID"
)

// ReturnOrderFailReason   * `ACCOUNT_ABNORMAL` - 分账接收方账户异常，  * `BALANCE_NOT_ENOUGH` - 余额不足，  * `TIME_OUT_CLOSED` - 超时关单，
type ReturnOrderFailReason string

func (e ReturnOrderFailReason) Ptr() *ReturnOrderFailReason {
	return &e
}

// Enums of ReturnOrderFailReason
const (
	RETURNORDERFAILREASON_ACCOUNT_ABNORMAL       ReturnOrderFailReason = "ACCOUNT_ABNORMAL"       // 原分账接收方账户异常
	RETURNORDERFAILREASON_PAYER_ACCOUNT_ABNORMAL ReturnOrderFailReason = "PAYER_ACCOUNT_ABNORMAL" // 原分账分出方账户异常
	RETURNORDERFAILREASON_TIME_OUT_CLOSED        ReturnOrderFailReason = "TIME_OUT_CLOSED"        // 超时关单
)

// ReturnOrderStatus   * `PROCESSING` - 处理中，  * `SUCCESS` - 已成功，  * `FAILED` - 已失败，
type ReturnOrderStatus string

func (e ReturnOrderStatus) Ptr() *ReturnOrderStatus {
	return &e
}

// Enums of ReturnOrderStatus
const (
	RETURNORDERSTATUS_PROCESSING ReturnOrderStatus = "PROCESSING"
	RETURNORDERSTATUS_SUCCESS    ReturnOrderStatus = "SUCCESS"
	RETURNORDERSTATUS_FAILED     ReturnOrderStatus = "FAILED"
)

// ReturnOrdersEntity
type ReturnOrdersEntity struct {
	// 参考请求参数
	SubMchid *string `json:"sub_mchid"`
	// 参考请求参数
	OrderId *string `json:"order_id"`
	// 参考请求参数
	OutOrderNo *string `json:"out_order_no"`
	// 参考请求参数
	OutReturnNo *string `json:"out_return_no"`
	// 只能对原分账请求中成功分给商户接收方进行回退
	ReturnMchid *string `json:"return_mchid"`
	// 微信分账回退单号，微信系统返回的唯一标识
	ReturnNo *string `json:"return_no"`
	// 需要从分账接收方回退的金额，单位为分，只能为整数
	Amount *int64 `json:"amount"`
	// 如果请求返回为处理中，则商户可以通过调用回退结果查询接口获取请求的最终处理结果。如果查询到回退结果在处理中，请勿变更商户回退单号，使用相同的参数再次发起分账回退，否则会出现资金风险。在处理中状态的回退单如果5天没有成功，会因为超时被设置为已失败。 枚举值： PROCESSING：处理中 SUCCESS：已成功 FAILED：已失败   * `PROCESSING` - 处理中，  * `SUCCESS` - 已成功，  * `FAILED` - 已失败，
	Result *ReturnOrderStatus `json:"result"`
	// 失败原因。包含以下枚举值： ACCOUNT_ABNORMAL : 分账接收方账户异常 TIME_OUT_CLOSED : 超时关单  * `ACCOUNT_ABNORMAL` - 分账接收方账户异常，  * `BALANCE_NOT_ENOUGH` - 余额不足，  * `TIME_OUT_CLOSED` - 超时关单，
	FailReason *ReturnOrderFailReason `json:"fail_reason,omitempty"`
	//// 分账回退创建时间，遵循RFC3339标准格式
	//CreateTime *time.Time `json:"create_time"`
	//// 分账回退的原因描述
	//Description *string `json:"description"`
	// 分账回退完成时间，遵循RFC3339标准格式
	FinishTime *time.Time `json:"finish_time"`
}

func (o ReturnOrdersEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in ReturnOrdersEntity")
	}
	toSerialize["amount"] = o.Amount

	//if o.CreateTime == nil {
	//	return nil, fmt.Errorf("field `CreateTime` is required and must be specified in ReturnOrdersEntity")
	//}
	//toSerialize["create_time"] = o.CreateTime.Format(time.RFC3339)
	//
	//if o.Description == nil {
	//	return nil, fmt.Errorf("field `Description` is required and must be specified in ReturnOrdersEntity")
	//}
	//toSerialize["description"] = o.Description

	if o.FailReason != nil {
		toSerialize["fail_reason"] = o.FailReason
	}

	if o.FinishTime == nil {
		return nil, fmt.Errorf("field `FinishTime` is required and must be specified in ReturnOrdersEntity")
	}
	toSerialize["finish_time"] = o.FinishTime.Format(time.RFC3339)

	if o.OrderId == nil {
		return nil, fmt.Errorf("field `OrderId` is required and must be specified in ReturnOrdersEntity")
	}
	toSerialize["order_id"] = o.OrderId

	if o.OutOrderNo == nil {
		return nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in ReturnOrdersEntity")
	}
	toSerialize["out_order_no"] = o.OutOrderNo

	if o.OutReturnNo == nil {
		return nil, fmt.Errorf("field `OutReturnNo` is required and must be specified in ReturnOrdersEntity")
	}
	toSerialize["out_return_no"] = o.OutReturnNo

	if o.Result == nil {
		return nil, fmt.Errorf("field `Result` is required and must be specified in ReturnOrdersEntity")
	}
	toSerialize["result"] = o.Result

	if o.ReturnNo == nil {
		return nil, fmt.Errorf("field `ReturnNo` is required and must be specified in ReturnOrdersEntity")
	}
	toSerialize["return_no"] = o.ReturnNo

	if o.ReturnMchid == nil {
		return nil, fmt.Errorf("field `ReturnMchid` is required and must be specified in ReturnOrdersEntity")
	}
	toSerialize["return_mchid"] = o.ReturnMchid

	if o.SubMchid == nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in ReturnOrdersEntity")
	}
	toSerialize["sub_mchid"] = o.SubMchid
	return json.Marshal(toSerialize)
}

func (o ReturnOrdersEntity) String() string {
	var ret string
	if o.Amount == nil {
		ret += "Amount:<nil>, "
	} else {
		ret += fmt.Sprintf("Amount:%v, ", *o.Amount)
	}

	//if o.CreateTime == nil {
	//	ret += "CreateTime:<nil>, "
	//} else {
	//	ret += fmt.Sprintf("CreateTime:%v, ", *o.CreateTime)
	//}
	//
	//if o.Description == nil {
	//	ret += "Description:<nil>, "
	//} else {
	//	ret += fmt.Sprintf("Description:%v, ", *o.Description)
	//}

	if o.FailReason == nil {
		ret += "FailReason:<nil>, "
	} else {
		ret += fmt.Sprintf("FailReason:%v, ", *o.FailReason)
	}

	if o.FinishTime == nil {
		ret += "FinishTime:<nil>, "
	} else {
		ret += fmt.Sprintf("FinishTime:%v, ", *o.FinishTime)
	}

	if o.OrderId == nil {
		ret += "OrderId:<nil>, "
	} else {
		ret += fmt.Sprintf("OrderId:%v, ", *o.OrderId)
	}

	if o.OutOrderNo == nil {
		ret += "OutOrderNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutOrderNo:%v, ", *o.OutOrderNo)
	}

	if o.OutReturnNo == nil {
		ret += "OutReturnNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutReturnNo:%v, ", *o.OutReturnNo)
	}

	if o.Result == nil {
		ret += "Result:<nil>, "
	} else {
		ret += fmt.Sprintf("Result:%v, ", *o.Result)
	}

	if o.ReturnNo == nil {
		ret += "ReturnNo:<nil>, "
	} else {
		ret += fmt.Sprintf("ReturnNo:%v, ", *o.ReturnNo)
	}

	if o.ReturnMchid == nil {
		ret += "ReturnMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("ReturnMchid:%v, ", *o.ReturnMchid)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>"
	} else {
		ret += fmt.Sprintf("SubMchid:%v", *o.SubMchid)
	}

	return fmt.Sprintf("ReturnOrdersEntity{%s}", ret)
}

func (o ReturnOrdersEntity) Clone() *ReturnOrdersEntity {
	ret := ReturnOrdersEntity{}

	if o.Amount != nil {
		ret.Amount = new(int64)
		*ret.Amount = *o.Amount
	}

	//if o.CreateTime != nil {
	//	ret.CreateTime = new(time.Time)
	//	*ret.CreateTime = *o.CreateTime
	//}
	//
	//if o.Description != nil {
	//	ret.Description = new(string)
	//	*ret.Description = *o.Description
	//}

	if o.FailReason != nil {
		ret.FailReason = new(ReturnOrderFailReason)
		*ret.FailReason = *o.FailReason
	}

	if o.FinishTime != nil {
		ret.FinishTime = new(time.Time)
		*ret.FinishTime = *o.FinishTime
	}

	if o.OrderId != nil {
		ret.OrderId = new(string)
		*ret.OrderId = *o.OrderId
	}

	if o.OutOrderNo != nil {
		ret.OutOrderNo = new(string)
		*ret.OutOrderNo = *o.OutOrderNo
	}

	if o.OutReturnNo != nil {
		ret.OutReturnNo = new(string)
		*ret.OutReturnNo = *o.OutReturnNo
	}

	if o.Result != nil {
		ret.Result = new(ReturnOrderStatus)
		*ret.Result = *o.Result
	}

	if o.ReturnNo != nil {
		ret.ReturnNo = new(string)
		*ret.ReturnNo = *o.ReturnNo
	}

	if o.ReturnMchid != nil {
		ret.ReturnMchid = new(string)
		*ret.ReturnMchid = *o.ReturnMchid
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	return &ret
}

// SplitBillHashType   * `SHA1` - SHA1，Secure Hash Algorithm 1
type SplitBillHashType string

func (e SplitBillHashType) Ptr() *SplitBillHashType {
	return &e
}

// Enums of SplitBillHashType
const (
	SPLITBILLHASHTYPE_SHA1 SplitBillHashType = "SHA1"
)

// SplitBillRequest
type SplitBillRequest struct {
	// 若商户是直连商户： - 无需填写该字段  若商户是服务商： - 不填则默认返回服务商下的所有分账账单。 - 如需下载某个子商户下的分账账单，则填指定的子商户号。
	SubMchid *string `json:"sub_mchid,omitempty"`
	// 格式YYYY-MM-DD。仅支持三个月内的账单下载申请。
	BillDate *string `json:"bill_date"`
	// 不填则以不压缩的方式返回数据流
	TarType *SplitBillTarType `json:"tar_type,omitempty"`
}

func (o SplitBillRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.SubMchid != nil {
		toSerialize["sub_mchid"] = o.SubMchid
	}

	if o.BillDate == nil {
		return nil, fmt.Errorf("field `BillDate` is required and must be specified in SplitBillRequest")
	}
	toSerialize["bill_date"] = o.BillDate

	if o.TarType != nil {
		toSerialize["tar_type"] = o.TarType
	}
	return json.Marshal(toSerialize)
}

func (o SplitBillRequest) String() string {
	var ret string
	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.BillDate == nil {
		ret += "BillDate:<nil>, "
	} else {
		ret += fmt.Sprintf("BillDate:%v, ", *o.BillDate)
	}

	if o.TarType == nil {
		ret += "TarType:<nil>"
	} else {
		ret += fmt.Sprintf("TarType:%v", *o.TarType)
	}

	return fmt.Sprintf("SplitBillRequest{%s}", ret)
}

func (o SplitBillRequest) Clone() *SplitBillRequest {
	ret := SplitBillRequest{}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.BillDate != nil {
		ret.BillDate = new(string)
		*ret.BillDate = *o.BillDate
	}

	if o.TarType != nil {
		ret.TarType = new(SplitBillTarType)
		*ret.TarType = *o.TarType
	}

	return &ret
}

// SplitBillResponse
type SplitBillResponse struct {
	// 供下一步请求账单文件的下载地址，该地址30s内有效
	DownloadUrl *string `json:"download_url"`
	// 原始账单（gzip需要解压缩）的摘要算法，用于校验文件的完整性  * `SHA1` - SHA1，Secure Hash Algorithm 1
	HashType *SplitBillHashType `json:"hash_type"`
	// 原始账单（gzip需要解压缩）的摘要值，用于校验文件的完整性
	HashValue *string `json:"hash_value"`
}

func (o SplitBillResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.DownloadUrl == nil {
		return nil, fmt.Errorf("field `DownloadUrl` is required and must be specified in SplitBillResponse")
	}
	toSerialize["download_url"] = o.DownloadUrl

	if o.HashType == nil {
		return nil, fmt.Errorf("field `HashType` is required and must be specified in SplitBillResponse")
	}
	toSerialize["hash_type"] = o.HashType

	if o.HashValue == nil {
		return nil, fmt.Errorf("field `HashValue` is required and must be specified in SplitBillResponse")
	}
	toSerialize["hash_value"] = o.HashValue
	return json.Marshal(toSerialize)
}

func (o SplitBillResponse) String() string {
	var ret string
	if o.DownloadUrl == nil {
		ret += "DownloadUrl:<nil>, "
	} else {
		ret += fmt.Sprintf("DownloadUrl:%v, ", *o.DownloadUrl)
	}

	if o.HashType == nil {
		ret += "HashType:<nil>, "
	} else {
		ret += fmt.Sprintf("HashType:%v, ", *o.HashType)
	}

	if o.HashValue == nil {
		ret += "HashValue:<nil>"
	} else {
		ret += fmt.Sprintf("HashValue:%v", *o.HashValue)
	}

	return fmt.Sprintf("SplitBillResponse{%s}", ret)
}

func (o SplitBillResponse) Clone() *SplitBillResponse {
	ret := SplitBillResponse{}

	if o.DownloadUrl != nil {
		ret.DownloadUrl = new(string)
		*ret.DownloadUrl = *o.DownloadUrl
	}

	if o.HashType != nil {
		ret.HashType = new(SplitBillHashType)
		*ret.HashType = *o.HashType
	}

	if o.HashValue != nil {
		ret.HashValue = new(string)
		*ret.HashValue = *o.HashValue
	}

	return &ret
}

// SplitBillTarType   * `GZIP` - GZIP格式压缩，返回格式为.gzip的压缩包账单
type SplitBillTarType string

func (e SplitBillTarType) Ptr() *SplitBillTarType {
	return &e
}

// Enums of SplitBillTarType
const (
	SPLITBILLTARTYPE_GZIP SplitBillTarType = "GZIP"
)

// UnfreezeOrderRequest
type UnfreezeOrderRequest struct {
	// 分账的原因描述，分账账单中需要体现
	Description *string `json:"description"`
	// 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutOrderNo *string `json:"out_order_no"`
	// 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要）
	SubMchid *string `json:"sub_mchid"`
	// 微信支付订单号
	TransactionId *string `json:"transaction_id"`
}

func (o UnfreezeOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Description == nil {
		return nil, fmt.Errorf("field `Description` is required and must be specified in UnfreezeOrderRequest")
	}
	toSerialize["description"] = o.Description

	if o.OutOrderNo == nil {
		return nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in UnfreezeOrderRequest")
	}
	toSerialize["out_order_no"] = o.OutOrderNo

	if o.SubMchid == nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in UnfreezeOrderRequest")
	}
	toSerialize["sub_mchid"] = o.SubMchid

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in UnfreezeOrderRequest")
	}
	toSerialize["transaction_id"] = o.TransactionId
	return json.Marshal(toSerialize)
}

func (o UnfreezeOrderRequest) String() string {
	var ret string
	if o.Description == nil {
		ret += "Description:<nil>, "
	} else {
		ret += fmt.Sprintf("Description:%v, ", *o.Description)
	}

	if o.OutOrderNo == nil {
		ret += "OutOrderNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutOrderNo:%v, ", *o.OutOrderNo)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>"
	} else {
		ret += fmt.Sprintf("TransactionId:%v", *o.TransactionId)
	}

	return fmt.Sprintf("UnfreezeOrderRequest{%s}", ret)
}

func (o UnfreezeOrderRequest) Clone() *UnfreezeOrderRequest {
	ret := UnfreezeOrderRequest{}

	if o.Description != nil {
		ret.Description = new(string)
		*ret.Description = *o.Description
	}

	if o.OutOrderNo != nil {
		ret.OutOrderNo = new(string)
		*ret.OutOrderNo = *o.OutOrderNo
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	return &ret
}

// FinishOrderResp .
type FinishOrderResp struct {
	// 微信分账单号，微信系统返回的唯一标识
	OrderId *string `json:"order_id"`
	// 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutOrderNo *string `json:"out_order_no"`
	// 分账出资的电商平台二级商户，填写微信支付分配的商户号。
	SubMchid *string `json:"sub_mchid"`
	// 微信支付订单号
	TransactionId *string `json:"transaction_id"`
}

func (o FinishOrderResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OrderId == nil {
		return nil, fmt.Errorf("field `OrderId` is required and must be specified in OrdersEntity")
	}
	toSerialize["order_id"] = o.OrderId

	if o.OutOrderNo == nil {
		return nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in OrdersEntity")
	}
	toSerialize["out_order_no"] = o.OutOrderNo

	if o.SubMchid != nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in OrdersEntity")
	}
	toSerialize["sub_mchid"] = o.SubMchid

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in OrdersEntity")
	}
	toSerialize["transaction_id"] = o.TransactionId
	return json.Marshal(toSerialize)
}

func (o FinishOrderResp) String() string {
	var ret string
	if o.OrderId == nil {
		ret += "OrderId:<nil>, "
	} else {
		ret += fmt.Sprintf("OrderId:%v, ", *o.OrderId)
	}

	if o.OutOrderNo == nil {
		ret += "OutOrderNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutOrderNo:%v, ", *o.OutOrderNo)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>"
	} else {
		ret += fmt.Sprintf("TransactionId:%v", *o.TransactionId)
	}

	return fmt.Sprintf("FinishOrderResp{%s}", ret)
}

func (o FinishOrderResp) Clone() *FinishOrderResp {
	ret := FinishOrderResp{}

	if o.OrderId != nil {
		ret.OrderId = new(string)
		*ret.OrderId = *o.OrderId
	}

	if o.OutOrderNo != nil {
		ret.OutOrderNo = new(string)
		*ret.OutOrderNo = *o.OutOrderNo
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	return &ret
}

// QueryOrderResp .
type QueryOrderResp struct {
	// 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要）
	SubMchid *string `json:"sub_mchid"`
	// 微信支付订单号
	TransactionId *string `json:"transaction_id"`
	// 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutOrderNo *string `json:"out_order_no"`
	// 微信分账单号，微信系统返回的唯一标识
	OrderId *string `json:"order_id"`
	// 分账单状态（每个接收方的分账结果请查看receivers中的result字段），枚举值： 1、PROCESSING：处理中 2、FINISHED：分账完成  * `PROCESSING` - 处理中，  * `FINISHED` - 分账完成，
	Status *OrderStatus `json:"status"`
	// 分账接收方列表
	Receivers []OrderReceiverDetail `json:"receivers"`
	// 分账完结的分账金额，单位为分， 仅当查询分账完结的执行结果时，存在本字段
	FinishAmount *int64 `json:"finish_amount,omitempty"`
	// 分账完结的原因描述，仅当查询分账完结的执行结果时，存在本字段。
	FinishDescription *string `json:"finish_description,omitempty"`
}

func (o QueryOrderResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OrderId == nil {
		return nil, fmt.Errorf("field `OrderId` is required and must be specified in OrdersEntity")
	}
	toSerialize["order_id"] = o.OrderId

	if o.OutOrderNo == nil {
		return nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in OrdersEntity")
	}
	toSerialize["out_order_no"] = o.OutOrderNo

	if o.Receivers != nil {
		return nil, fmt.Errorf("field `Receivers` is required and must be specified in OrdersEntity")
	}
	toSerialize["receivers"] = o.Receivers

	if o.Status == nil {
		return nil, fmt.Errorf("field `Status` is required and must be specified in OrdersEntity")
	}
	toSerialize["state"] = o.Status

	if o.SubMchid != nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in OrdersEntity")
	}
	toSerialize["sub_mchid"] = o.SubMchid

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in OrdersEntity")
	}
	toSerialize["transaction_id"] = o.TransactionId

	if o.FinishAmount != nil {
		toSerialize["finish_amount"] = o.FinishAmount
	}
	if o.FinishDescription != nil {
		toSerialize["finish_description"] = o.FinishDescription
	}
	return json.Marshal(toSerialize)
}

func (o QueryOrderResp) String() string {
	var ret string
	if o.OrderId == nil {
		ret += "OrderId:<nil>, "
	} else {
		ret += fmt.Sprintf("OrderId:%v, ", *o.OrderId)
	}

	if o.OutOrderNo == nil {
		ret += "OutOrderNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutOrderNo:%v, ", *o.OutOrderNo)
	}

	ret += fmt.Sprintf("Receivers:%v, ", o.Receivers)

	if o.Status == nil {
		ret += "Status:<nil>, "
	} else {
		ret += fmt.Sprintf("Status:%v, ", *o.Status)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>"
	} else {
		ret += fmt.Sprintf("TransactionId:%v", *o.TransactionId)
	}

	if o.FinishAmount == nil {
		ret += "FinishAmount:<nil>"
	} else {
		ret += fmt.Sprintf("FinishAmount:%v", *o.FinishAmount)
	}

	if o.FinishDescription == nil {
		ret += "FinishDescription:<nil>"
	} else {
		ret += fmt.Sprintf("FinishDescription:%v", *o.FinishDescription)
	}

	return fmt.Sprintf("QueryOrderResp{%s}", ret)
}

func (o QueryOrderResp) Clone() *QueryOrderResp {
	ret := QueryOrderResp{}

	if o.OrderId != nil {
		ret.OrderId = new(string)
		*ret.OrderId = *o.OrderId
	}

	if o.OutOrderNo != nil {
		ret.OutOrderNo = new(string)
		*ret.OutOrderNo = *o.OutOrderNo
	}

	if o.Receivers != nil {
		ret.Receivers = make([]OrderReceiverDetail, len(o.Receivers))
		for i, item := range o.Receivers {
			ret.Receivers[i] = *item.Clone()
		}
	}

	if o.Status != nil {
		ret.Status = new(OrderStatus)
		*ret.Status = *o.Status
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.FinishAmount != nil {
		ret.FinishAmount = new(int64)
		*ret.FinishAmount = *o.FinishAmount
	}

	if o.FinishDescription != nil {
		ret.FinishDescription = new(string)
		*ret.FinishDescription = *o.FinishDescription
	}

	return &ret
}

// Channel * `ORIGINAL` - 原路退款, 退款渠道 * `BALANCE` - 退回到余额, 退款渠道 * `OTHER_BALANCE` - 原账户异常退到其他余额账户, 退款渠道 * `OTHER_BANKCARD` - 原银行卡异常退到其他银行卡, 退款渠道
type Channel string

func (e Channel) Ptr() *Channel {
	return &e
}

// Enums of Channel
const (
	CHANNEL_ORIGINAL       Channel = "ORIGINAL"
	CHANNEL_BALANCE        Channel = "BALANCE"
	CHANNEL_OTHER_BALANCE  Channel = "OTHER_BALANCE"
	CHANNEL_OTHER_BANKCARD Channel = "OTHER_BANKCARD"
)

// Scope * `GLOBAL` - 全场代金券, 全场优惠类型 * `SINGLE` - 单品优惠, 单品优惠类型
type Scope string

func (e Scope) Ptr() *Scope {
	return &e
}

// Enums of Scope
const (
	SCOPE_GLOBAL Scope = "GLOBAL"
	SCOPE_SINGLE Scope = "SINGLE"
)

// Status * `SUCCESS` - 退款成功, 退款状态 * `CLOSED` - 退款关闭, 退款状态 * `PROCESSING` - 退款处理中, 退款状态 * `ABNORMAL` - 退款异常, 退款状态
type Status string

func (e Status) Ptr() *Status {
	return &e
}

// Enums of Status
const (
	STATUS_SUCCESS    Status = "SUCCESS"
	STATUS_CLOSE      Status = "CLOSE"
	STATUS_PROCESSING Status = "PROCESSING"
	STATUS_ABNORMAL   Status = "ABNORMAL"
)

// Type * `COUPON` - 代金券, 代金券类型，需要走结算资金的充值型代金券 * `DISCOUNT` - 优惠券, 优惠券类型，不走结算资金的免充值型优惠券
type Type string

func (e Type) Ptr() *Type {
	return &e
}

// Enums of Type
const (
	TYPE_COUPON   Type = "COUPON"
	TYPE_DISCOUNT Type = "DISCOUNT"
)

// FundsAccount * `UNSETTLED` - 未结算资金, 退款所使用资金对应的资金账户类型 * `AVAILABLE` - 可用余额, 退款所使用资金对应的资金账户类型 * `UNAVAILABLE` - 不可用余额, 退款所使用资金对应的资金账户类型 * `OPERATION` - 运营户, 退款所使用资金对应的资金账户类型 * `BASIC` - 基本账户（含可用余额和不可用余额）, 退款所使用资金对应的资金账户类型
type FundsAccount string

func (e FundsAccount) Ptr() *FundsAccount {
	return &e
}

// Enums of FundsAccount
const (
	FUNDSACCOUNT_UNSETTLED   FundsAccount = "UNSETTLED"
	FUNDSACCOUNT_AVAILABLE   FundsAccount = "AVAILABLE"
	FUNDSACCOUNT_UNAVAILABLE FundsAccount = "UNAVAILABLE"
	FUNDSACCOUNT_OPERATION   FundsAccount = "OPERATION"
	FUNDSACCOUNT_BASIC       FundsAccount = "BASIC"
)

// RefundAccount 电商平台垫资退款专用参数。需先确认已开通此功能后，才能使用。若需要开通，请联系微信支付客服。
type RefundAccount string

func (e RefundAccount) Ptr() *RefundAccount {
	return &e
}

// Enums of RefundAccount
const (
	REFUNDACCOUNT_REFUND_SOURCE_PARTNER_ADVANCE RefundAccount = "REFUND_SOURCE_PARTNER_ADVANCE" // 电商平台垫付，需要向微信支付申请开通
	REUFNDACCOUNT_REFUND_SOURCE_SUB_MERCHANT    RefundAccount = "REFUND_SOURCE_SUB_MERCHANT"    // 二级商户，默认值
)

// AmountReq .
type AmountReq struct {
	// 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额。
	Refund *int64 `json:"refund"`
	// 原支付交易的订单总金额，币种的最小单位，只能为整数。
	Total *int64 `json:"total"`
	// 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY。
	Currency *string `json:"currency"`
}

func (o AmountReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Refund == nil {
		return nil, fmt.Errorf("field `Refund` is required and must be specified in AmountReq")
	}
	toSerialize["refund"] = o.Refund

	if o.Total == nil {
		return nil, fmt.Errorf("field `Total` is required and must be specified in AmountReq")
	}
	toSerialize["total"] = o.Total

	if o.Currency == nil {
		return nil, fmt.Errorf("field `Currency` is required and must be specified in AmountReq")
	}
	toSerialize["currency"] = o.Currency
	return json.Marshal(toSerialize)
}

func (o AmountReq) String() string {
	var ret string
	if o.Refund == nil {
		ret += "Refund:<nil>, "
	} else {
		ret += fmt.Sprintf("Refund:%v, ", *o.Refund)
	}

	if o.Total == nil {
		ret += "Total:<nil>, "
	} else {
		ret += fmt.Sprintf("Total:%v, ", *o.Total)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>"
	} else {
		ret += fmt.Sprintf("Currency:%v", *o.Currency)
	}

	return fmt.Sprintf("AmountReq{%s}", ret)
}

func (o AmountReq) Clone() *AmountReq {
	ret := AmountReq{}

	if o.Refund != nil {
		ret.Refund = new(int64)
		*ret.Refund = *o.Refund
	}

	if o.Total != nil {
		ret.Total = new(int64)
		*ret.Total = *o.Total
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	return &ret
}

// CreateRefundRequest .
type CreateRefundRequest struct {
	// 微信支付分配二级商户的商户号。
	SubMchid *string `json:"sub_mchid"`
	// 电商平台在微信公众平台申请服务号对应的APPID，申请商户功能的时候微信支付会配置绑定关系。
	SpAppID *string `json:"sp_appid"`
	// 原支付交易对应的微信订单号
	SubAppID *string `json:"sub_appid,omitempty"`
	// 原支付交易对应的微信订单号
	TransactionId *string `json:"transaction_id,omitempty"`
	// 原支付交易对应的商户订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`
	// 若商户传入，会在下发给用户的退款消息中体现退款原因
	Reason *string `json:"reason,omitempty"`
	// 订单金额信息
	Amount *AmountReq `json:"amount"`
	// 异步接收微信支付退款结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效，优先回调当前传的这个地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
	// 若传递此参数则使用对应的资金账户退款，否则默认使用未结算资金退款（仅对老资金流商户适用）  枚举值： - AVAILABLE：可用余额账户    * `AVAILABLE` - 可用余额
	FundsAccount *string `json:"funds_account,omitempty"`
	// 电商平台垫资退款专用参数。需先确认已开通此功能后，才能使用。若需要开通，请联系微信支付客服。
	//枚举值：
	//REFUND_SOURCE_PARTNER_ADVANCE : 电商平台垫付，需要向微信支付申请开通
	//REFUND_SOURCE_SUB_MERCHANT : 二级商户，默认值
	//注意：
	//若传入REFUND_SOURCE_PARTNER_ADVANCE，仅代表可以使用垫付退款，实际出款账户需以退款申请受理结果或查单结果为准。
	RefundAccount *string `json:"refund_account,omitempty"`
}

func (o CreateRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.SubMchid == nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in CreateRefundRequest")
	}
	toSerialize["sub_mchid"] = o.SubMchid

	if o.SpAppID == nil {
		return nil, fmt.Errorf("field `SpAppID` is required and must be specified in CreateRefundRequest")
	}
	toSerialize["sp_appid"] = o.SpAppID

	if o.SubAppID != nil {
		toSerialize["sub_appid"] = o.SubAppID
	}

	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
	}

	if o.OutTradeNo != nil {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}

	if o.OutRefundNo == nil {
		return nil, fmt.Errorf("field `OutRefundNo` is required and must be specified in CreateRefundRequest")
	}
	toSerialize["out_refund_no"] = o.OutRefundNo

	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	if o.NotifyUrl != nil {
		toSerialize["notify_url"] = o.NotifyUrl
	}

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in CreateRefundRequest")
	}
	toSerialize["amount"] = o.Amount

	if o.FundsAccount != nil {
		toSerialize["funds_account"] = o.FundsAccount
	}

	if o.RefundAccount != nil {
		toSerialize["refund_account"] = o.RefundAccount
	}
	return json.Marshal(toSerialize)
}

func (o CreateRefundRequest) String() string {
	var ret string
	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.SpAppID == nil {
		ret += "SpAppID:<nil>, "
	} else {
		ret += fmt.Sprintf("SpAppID:%v, ", *o.SpAppID)
	}

	if o.SubAppID == nil {
		ret += "SubAppID:<nil>, "
	} else {
		ret += fmt.Sprintf("SubAppID:%v, ", *o.SubAppID)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.Reason == nil {
		ret += "Reason:<nil>, "
	} else {
		ret += fmt.Sprintf("Reason:%v, ", *o.Reason)
	}

	if o.NotifyUrl == nil {
		ret += "NotifyUrl:<nil>, "
	} else {
		ret += fmt.Sprintf("NotifyUrl:%v, ", *o.NotifyUrl)
	}

	if o.FundsAccount == nil {
		ret += "FundsAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("FundsAccount:%v, ", *o.FundsAccount)
	}

	if o.RefundAccount == nil {
		ret += "RefundAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundAccount:%v, ", *o.RefundAccount)
	}

	ret += fmt.Sprintf("Amount:%v, ", o.Amount)

	return fmt.Sprintf("CreateRefundRequest{%s}", ret)
}

func (o CreateRefundRequest) Clone() *CreateRefundRequest {
	ret := CreateRefundRequest{}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.SpAppID != nil {
		ret.SpAppID = new(string)
		*ret.SpAppID = *o.SpAppID
	}

	if o.SubAppID != nil {
		ret.SubAppID = new(string)
		*ret.SubAppID = *o.SubAppID
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.OutRefundNo != nil {
		ret.OutRefundNo = new(string)
		*ret.OutRefundNo = *o.OutRefundNo
	}

	if o.Reason != nil {
		ret.Reason = new(string)
		*ret.Reason = *o.Reason
	}

	if o.NotifyUrl != nil {
		ret.NotifyUrl = new(string)
		*ret.NotifyUrl = *o.NotifyUrl
	}

	if o.FundsAccount != nil {
		ret.FundsAccount = new(string)
		*ret.FundsAccount = *o.FundsAccount
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.RefundAccount != nil {
		ret.RefundAccount = new(string)
		*ret.RefundAccount = *o.RefundAccount
	}

	return &ret
}

// Amount
type Amount struct {
	// 退款标价金额，单位为分，可以做部分退款
	Refund *int64 `json:"refund"`
	// 退款给用户的金额，不包含所有优惠券金额
	PayerRefund *int64 `json:"payer_refund"`
	// 优惠退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金，说明详见代金券或立减优惠，单位为分
	DiscountRefund *int64 `json:"discount_refund"`
	// 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY。
	Currency *string `json:"currency"`
}

func (o Amount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Refund == nil {
		return nil, fmt.Errorf("field `Refund` is required and must be specified in Amount")
	}
	toSerialize["refund"] = o.Refund

	if o.PayerRefund == nil {
		return nil, fmt.Errorf("field `PayerRefund` is required and must be specified in Amount")
	}
	toSerialize["payer_refund"] = o.PayerRefund

	if o.DiscountRefund == nil {
		return nil, fmt.Errorf("field `DiscountRefund` is required and must be specified in Amount")
	}
	toSerialize["discount_refund"] = o.DiscountRefund

	if o.Currency == nil {
		return nil, fmt.Errorf("field `Currency` is required and must be specified in Amount")
	}
	toSerialize["currency"] = o.Currency
	return json.Marshal(toSerialize)
}

func (o Amount) String() string {
	var ret string

	if o.Refund == nil {
		ret += "Refund:<nil>, "
	} else {
		ret += fmt.Sprintf("Refund:%v, ", *o.Refund)
	}

	if o.PayerRefund == nil {
		ret += "PayerRefund:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerRefund:%v, ", *o.PayerRefund)
	}

	if o.DiscountRefund == nil {
		ret += "DiscountRefund:<nil>, "
	} else {
		ret += fmt.Sprintf("DiscountRefund:%v, ", *o.DiscountRefund)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>"
	} else {
		ret += fmt.Sprintf("Currency:%v", *o.Currency)
	}

	return fmt.Sprintf("Amount{%s}", ret)
}

func (o Amount) Clone() *Amount {
	ret := Amount{}

	if o.Refund != nil {
		ret.Refund = new(int64)
		*ret.Refund = *o.Refund
	}

	if o.PayerRefund != nil {
		ret.PayerRefund = new(int64)
		*ret.PayerRefund = *o.PayerRefund
	}

	if o.DiscountRefund != nil {
		ret.DiscountRefund = new(int64)
		*ret.DiscountRefund = *o.DiscountRefund
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	return &ret
}

// Promotion .
type Promotion struct {
	// 券或者立减优惠id
	PromotionId *string `json:"promotion_id"`
	// 枚举值： - GLOBAL- 全场代金券 - SINGLE- 单品优惠 * `GLOBAL` - 全场代金券 * `SINGLE` - 单品优惠
	Scope *Scope `json:"scope"`
	// 枚举值： - COUPON- 代金券，需要走结算资金的充值型代金券 - DISCOUNT- 优惠券，不走结算资金的免充值型优惠券 * `COUPON` - 代金券 * `DISCOUNT` - 优惠券
	Type *Type `json:"type"`
	// 用户享受优惠的金额（优惠券面额=微信出资金额+商家出资金额+其他出资方金额 ），单位为分
	Amount *int64 `json:"amount"`
	// 优惠退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为用户支付的现金，说明详见代金券或立减优惠，单位为分
	RefundAmount *int64 `json:"refund_amount"`
}

func (o Promotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.PromotionId == nil {
		return nil, fmt.Errorf("field `PromotionId` is required and must be specified in Promotion")
	}
	toSerialize["promotion_id"] = o.PromotionId

	if o.Scope == nil {
		return nil, fmt.Errorf("field `Scope` is required and must be specified in Promotion")
	}
	toSerialize["scope"] = o.Scope

	if o.Type == nil {
		return nil, fmt.Errorf("field `Type` is required and must be specified in Promotion")
	}
	toSerialize["type"] = o.Type

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in Promotion")
	}
	toSerialize["amount"] = o.Amount

	if o.RefundAmount == nil {
		return nil, fmt.Errorf("field `RefundAmount` is required and must be specified in Promotion")
	}
	toSerialize["refund_amount"] = o.RefundAmount

	return json.Marshal(toSerialize)
}

func (o Promotion) String() string {
	var ret string
	if o.PromotionId == nil {
		ret += "PromotionId:<nil>, "
	} else {
		ret += fmt.Sprintf("PromotionId:%v, ", *o.PromotionId)
	}

	if o.Scope == nil {
		ret += "Scope:<nil>, "
	} else {
		ret += fmt.Sprintf("Scope:%v, ", *o.Scope)
	}

	if o.Type == nil {
		ret += "Type:<nil>, "
	} else {
		ret += fmt.Sprintf("Type:%v, ", *o.Type)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>, "
	} else {
		ret += fmt.Sprintf("Amount:%v, ", *o.Amount)
	}

	if o.RefundAmount == nil {
		ret += "RefundAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundAmount:%v, ", *o.RefundAmount)
	}

	return fmt.Sprintf("Promotion{%s}", ret)
}

func (o Promotion) Clone() *Promotion {
	ret := Promotion{}

	if o.PromotionId != nil {
		ret.PromotionId = new(string)
		*ret.PromotionId = *o.PromotionId
	}

	if o.Scope != nil {
		ret.Scope = new(Scope)
		*ret.Scope = *o.Scope
	}

	if o.Type != nil {
		ret.Type = new(Type)
		*ret.Type = *o.Type
	}

	if o.Amount != nil {
		ret.Amount = new(int64)
		*ret.Amount = *o.Amount
	}

	if o.RefundAmount != nil {
		ret.RefundAmount = new(int64)
		*ret.RefundAmount = *o.RefundAmount
	}

	return &ret
}

// Refund
type Refund struct {
	// 微信支付退款号
	RefundId *string `json:"refund_id"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`
	// 退款受理时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。
	CreateTime *time.Time `json:"create_time"`
	// 退款所使用资金对应的资金账户类型 枚举值：
	//REFUND_SOURCE_PARTNER_ADVANCE : 电商平台垫付
	//REFUND_SOURCE_SUB_MERCHANT : 二级商户，默认值
	RefundAccount *RefundAccount `json:"refund_account,omitempty"`
	// 金额详细信息
	Amount *Amount `json:"amount"`
	// 优惠退款信息
	PromotionDetail []Promotion `json:"promotion_detail,omitempty"`
}

func (o Refund) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.RefundId == nil {
		return nil, fmt.Errorf("field `RefundId` is required and must be specified in Refund")
	}
	toSerialize["refund_id"] = o.RefundId

	if o.OutRefundNo == nil {
		return nil, fmt.Errorf("field `OutRefundNo` is required and must be specified in Refund")
	}
	toSerialize["out_refund_no"] = o.OutRefundNo

	if o.CreateTime == nil {
		return nil, fmt.Errorf("field `CreateTime` is required and must be specified in Refund")
	}
	toSerialize["create_time"] = o.CreateTime.Format(time.RFC3339)

	if o.RefundAccount != nil {
		toSerialize["refund_account"] = o.RefundAccount
	}

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in Refund")
	}
	toSerialize["amount"] = o.Amount

	if o.PromotionDetail != nil {
		toSerialize["promotion_detail"] = o.PromotionDetail
	}
	return json.Marshal(toSerialize)
}

func (o Refund) String() string {
	var ret string
	if o.RefundId == nil {
		ret += "RefundId:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundId:%v, ", *o.RefundId)
	}

	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.CreateTime == nil {
		ret += "CreateTime:<nil>, "
	} else {
		ret += fmt.Sprintf("CreateTime:%v, ", *o.CreateTime)
	}

	if o.RefundAccount == nil {
		ret += "RefundAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundAccount:%v, ", *o.RefundAccount)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>, "
	} else {
		ret += fmt.Sprintf("Amount:%v, ", *o.Amount)
	}

	if o.PromotionDetail == nil {
		ret += "PromotionDetail:<nil>, "
	} else {
		ret += fmt.Sprintf("PromotionDetail:%v, ", o.PromotionDetail)
	}

	return fmt.Sprintf("Refund{%s}", ret)
}

func (o Refund) Clone() *Refund {
	ret := Refund{}

	if o.RefundId != nil {
		ret.RefundId = new(string)
		*ret.RefundId = *o.RefundId
	}

	if o.OutRefundNo != nil {
		ret.OutRefundNo = new(string)
		*ret.OutRefundNo = *o.OutRefundNo
	}

	if o.CreateTime != nil {
		ret.CreateTime = new(time.Time)
		*ret.CreateTime = *o.CreateTime
	}

	if o.RefundAccount != nil {
		ret.RefundAccount = new(RefundAccount)
		*ret.RefundAccount = *o.RefundAccount
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.PromotionDetail != nil {
		ret.PromotionDetail = make([]Promotion, len(o.PromotionDetail))
		for i, item := range o.PromotionDetail {
			ret.PromotionDetail[i] = *item.Clone()
		}
	}

	return &ret
}

// QueryByOutRefundNoRequest
type QueryByOutRefundNoRequest struct {
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`
	// 子商户的商户号，由微信支付生成并下发。服务商模式下必须传递此参数
	SubMchid *string `json:"sub_mchid"`
}

func (o QueryByOutRefundNoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OutRefundNo == nil {
		return nil, fmt.Errorf("field `OutRefundNo` is required and must be specified in QueryByOutRefundNoRequest")
	}
	toSerialize["out_refund_no"] = o.OutRefundNo

	if o.SubMchid == nil {
		return nil, fmt.Errorf("field `SubMchid` is required and must be specified in QueryByOutRefundNoRequest")
	}
	toSerialize["sub_mchid"] = o.SubMchid

	return json.Marshal(toSerialize)
}

func (o QueryByOutRefundNoRequest) String() string {
	var ret string
	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>"
	} else {
		ret += fmt.Sprintf("SubMchid:%v", *o.SubMchid)
	}

	return fmt.Sprintf("QueryByOutRefundNoRequest{%s}", ret)
}

func (o QueryByOutRefundNoRequest) Clone() *QueryByOutRefundNoRequest {
	ret := QueryByOutRefundNoRequest{}

	if o.OutRefundNo != nil {
		ret.OutRefundNo = new(string)
		*ret.OutRefundNo = *o.OutRefundNo
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	return &ret
}

// QueryRefundResp .
type QueryRefundResp struct {
	// 微信支付退款号
	RefundId *string `json:"refund_id"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`
	// 微信支付交易订单号
	TransactionId *string `json:"transaction_id"`
	// 原支付交易对应的商户订单号
	OutTradeNo *string `json:"out_trade_no"`
	// 枚举值： - ORIGINAL—原路退款 - BALANCE—退回到余额 - OTHER_BALANCE—原账户异常退到其他余额账户 - OTHER_BANKCARD—原银行卡异常退到其他银行卡 * `ORIGINAL` - 原路退款 * `BALANCE` - 退回到余额 * `OTHER_BALANCE` - 原账户异常退到其他余额账户 * `OTHER_BANKCARD` - 原银行卡异常退到其他银行卡
	Channel *Channel `json:"channel"`
	// 取当前退款单的退款入账方，有以下几种情况： 1）退回银行卡：{银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱:支付用户零钱 3）退还商户:商户基本账户商户结算银行账户 4）退回支付用户零钱通:支付用户零钱通
	UserReceivedAccount *string `json:"user_received_account"`
	// 退款成功时间，退款状态status为SUCCESS（退款成功）时，返回该字段。遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。
	SuccessTime *time.Time `json:"success_time,omitempty"`
	// 退款受理时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。
	CreateTime *time.Time `json:"create_time"`
	// 退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往商户平台（pay.weixin.qq.com）-交易中心，手动处理此笔退款。 枚举值： - SUCCESS—退款成功 - CLOSED—退款关闭 - PROCESSING—退款处理中 - ABNORMAL—退款异常 * `SUCCESS` - 退款成功 * `CLOSED` - 退款关闭 * `PROCESSING` - 退款处理中 * `ABNORMAL` - 退款异常
	Status *Status `json:"status"`
	// 金额详细信息
	Amount *Amount `json:"amount"`
	// 优惠退款信息
	PromotionDetail []Promotion `json:"promotion_detail,omitempty"`
	// 电商平台垫资退款专用参数。需先确认已开通此功能后，才能使用。若需要开通，请联系微信支付客服。
	// 枚举值：
	// REFUND_SOURCE_PARTNER_ADVANCE : 电商平台垫付，需要向微信支付申请开通
	// REFUND_SOURCE_SUB_MERCHANT : 二级商户，默认值
	// 注意：若传入REFUND_SOURCE_PARTNER_ADVANCE，仅代表可以使用垫付退款，实际出款账户需以退款申请受理结果或查单结果为准。
	RefundAccount *RefundAccount `json:"refund_account,omitempty"`

	// 退款所使用资金对应的资金账户类型 枚举值： - UNSETTLED : 未结算资金 - AVAILABLE : 可用余额 - UNAVAILABLE : 不可用余额 - OPERATION : 运营户 - BASIC : 基本账户（含可用余额和不可用余额） * `UNSETTLED` - 未结算资金 * `AVAILABLE` - 可用余额 * `UNAVAILABLE` - 不可用余额 * `OPERATION` - 运营户 * `BASIC` - 基本账户（含可用余额和不可用余额）
	FundsAccount *FundsAccount `json:"funds_account,omitempty"`
}

func (o QueryRefundResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.RefundId == nil {
		return nil, fmt.Errorf("field `RefundId` is required and must be specified in Refund")
	}
	toSerialize["refund_id"] = o.RefundId

	if o.OutRefundNo == nil {
		return nil, fmt.Errorf("field `OutRefundNo` is required and must be specified in Refund")
	}
	toSerialize["out_refund_no"] = o.OutRefundNo

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in Refund")
	}
	toSerialize["transaction_id"] = o.TransactionId

	if o.OutTradeNo == nil {
		return nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in Refund")
	}
	toSerialize["out_trade_no"] = o.OutTradeNo

	if o.Channel == nil {
		return nil, fmt.Errorf("field `Channel` is required and must be specified in Refund")
	}
	toSerialize["channel"] = o.Channel

	if o.UserReceivedAccount == nil {
		return nil, fmt.Errorf("field `UserReceivedAccount` is required and must be specified in Refund")
	}
	toSerialize["user_received_account"] = o.UserReceivedAccount

	if o.SuccessTime != nil {
		toSerialize["success_time"] = o.SuccessTime.Format(time.RFC3339)
	}

	if o.CreateTime == nil {
		return nil, fmt.Errorf("field `CreateTime` is required and must be specified in Refund")
	}
	toSerialize["create_time"] = o.CreateTime.Format(time.RFC3339)

	if o.Status == nil {
		return nil, fmt.Errorf("field `Status` is required and must be specified in Refund")
	}
	toSerialize["status"] = o.Status

	if o.FundsAccount != nil {
		toSerialize["funds_account"] = o.FundsAccount
	}

	if o.RefundAccount != nil {
		toSerialize["refund_account"] = o.RefundAccount
	}

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in Refund")
	}
	toSerialize["amount"] = o.Amount

	if o.PromotionDetail != nil {
		toSerialize["promotion_detail"] = o.PromotionDetail
	}
	return json.Marshal(toSerialize)
}

func (o QueryRefundResp) String() string {
	var ret string
	if o.RefundId == nil {
		ret += "RefundId:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundId:%v, ", *o.RefundId)
	}

	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.Channel == nil {
		ret += "Channel:<nil>, "
	} else {
		ret += fmt.Sprintf("Channel:%v, ", *o.Channel)
	}

	if o.UserReceivedAccount == nil {
		ret += "UserReceivedAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("UserReceivedAccount:%v, ", *o.UserReceivedAccount)
	}

	if o.SuccessTime == nil {
		ret += "SuccessTime:<nil>, "
	} else {
		ret += fmt.Sprintf("SuccessTime:%v, ", *o.SuccessTime)
	}

	if o.CreateTime == nil {
		ret += "CreateTime:<nil>, "
	} else {
		ret += fmt.Sprintf("CreateTime:%v, ", *o.CreateTime)
	}

	if o.Status == nil {
		ret += "Status:<nil>, "
	} else {
		ret += fmt.Sprintf("Status:%v, ", *o.Status)
	}

	if o.FundsAccount == nil {
		ret += "FundsAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("FundsAccount:%v, ", *o.FundsAccount)
	}

	if o.RefundAccount == nil {
		ret += "RefundAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundAccount:%v, ", *o.RefundAccount)
	}

	ret += fmt.Sprintf("Amount:%v, ", o.Amount)

	ret += fmt.Sprintf("PromotionDetail:%v", o.PromotionDetail)

	return fmt.Sprintf("QueryRefundResp{%s}", ret)
}

func (o QueryRefundResp) Clone() *QueryRefundResp {
	ret := QueryRefundResp{}

	if o.RefundId != nil {
		ret.RefundId = new(string)
		*ret.RefundId = *o.RefundId
	}

	if o.OutRefundNo != nil {
		ret.OutRefundNo = new(string)
		*ret.OutRefundNo = *o.OutRefundNo
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.Channel != nil {
		ret.Channel = new(Channel)
		*ret.Channel = *o.Channel
	}

	if o.UserReceivedAccount != nil {
		ret.UserReceivedAccount = new(string)
		*ret.UserReceivedAccount = *o.UserReceivedAccount
	}

	if o.SuccessTime != nil {
		ret.SuccessTime = new(time.Time)
		*ret.SuccessTime = *o.SuccessTime
	}

	if o.CreateTime != nil {
		ret.CreateTime = new(time.Time)
		*ret.CreateTime = *o.CreateTime
	}

	if o.Status != nil {
		ret.Status = new(Status)
		*ret.Status = *o.Status
	}

	if o.FundsAccount != nil {
		ret.FundsAccount = new(FundsAccount)
		*ret.FundsAccount = *o.FundsAccount
	}

	if o.RefundAccount != nil {
		ret.RefundAccount = new(RefundAccount)
		*ret.RefundAccount = *o.RefundAccount
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.PromotionDetail != nil {
		ret.PromotionDetail = make([]Promotion, len(o.PromotionDetail))
		for i, item := range o.PromotionDetail {
			ret.PromotionDetail[i] = *item.Clone()
		}
	}

	return &ret
}

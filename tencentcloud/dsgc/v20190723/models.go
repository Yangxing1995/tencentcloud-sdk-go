// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190723

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AKSKLeak struct {
	// AK编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	AK *string `json:"AK,omitempty" name:"AK"`

	// SK编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SK *string `json:"SK,omitempty" name:"SK"`

	// URL编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	URL *string `json:"URL,omitempty" name:"URL"`
}

type AccountRisk struct {
	// id（可不参考）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 风险账户
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskAccount *string `json:"RiskAccount,omitempty" name:"RiskAccount"`
}

type AssessmentControlItem struct {
	// 评估项Id
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 评估项名称
	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`

	// 评估项描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 评估项来源，内置/用户自定，取值（system，user）
	Source *string `json:"Source,omitempty" name:"Source"`

	// 评估项类型，问卷/自动化，取值（questionnaire，auto）
	ItemType *string `json:"ItemType,omitempty" name:"ItemType"`

	// 评估项子类型，单选/多选/时间/文本/AKSK等，取值（singlechoice，multichoice，date，text，AKSK……等）
	ItemSubType *string `json:"ItemSubType,omitempty" name:"ItemSubType"`

	// 评估项创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 评估项启用状态，启用/未启用，取值draft / launched
	Status *string `json:"Status,omitempty" name:"Status"`

	// 评估项关联的模版数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateCount *int64 `json:"TemplateCount,omitempty" name:"TemplateCount"`
}

type AssessmentRisk struct {
	// 风险项Id
	RiskId *string `json:"RiskId,omitempty" name:"RiskId"`

	// 风险项描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskDescription *string `json:"RiskDescription,omitempty" name:"RiskDescription"`

	// 评估模版Id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 评估模版名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 评估项Id
	ControlItemId *string `json:"ControlItemId,omitempty" name:"ControlItemId"`

	// 评估项名称
	ControlItemName *string `json:"ControlItemName,omitempty" name:"ControlItemName"`

	// 评估描述
	ControlItemDesc *string `json:"ControlItemDesc,omitempty" name:"ControlItemDesc"`

	// 风险等级，取值（high，medium，low）
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 风险缓解措施
	RiskMitigation *string `json:"RiskMitigation,omitempty" name:"RiskMitigation"`

	// 风险处理状态。(waiting待处理, processing处理中, finished已处理)
	Status *string `json:"Status,omitempty" name:"Status"`

	// 风险生成时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 风险负责人
	RiskOwner *string `json:"RiskOwner,omitempty" name:"RiskOwner"`

	// 风险涉及资产
	RelatedAsset *string `json:"RelatedAsset,omitempty" name:"RelatedAsset"`

	// 风险涉及资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 风险涉及资产名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceName *string `json:"DataSourceName,omitempty" name:"DataSourceName"`

	// 资产名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 建议使用安全产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityProduct []*SecurityProduct `json:"SecurityProduct,omitempty" name:"SecurityProduct"`

	// 风险类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`
}

type AssessmentRiskItem struct {
	// 脆弱项id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskName *string `json:"RiskName,omitempty" name:"RiskName"`

	// 脆弱性级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 风险类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`

	// 关联模版个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReferTemplateCount *int64 `json:"ReferTemplateCount,omitempty" name:"ReferTemplateCount"`

	// 支持的数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportDataSource []*string `json:"SupportDataSource,omitempty" name:"SupportDataSource"`
}

type AssessmentTask struct {
	// 评估任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 评估任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 业务名称
	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`

	// 业务所属部门
	BusinessDept *string `json:"BusinessDept,omitempty" name:"BusinessDept"`

	// 业务负责人
	BusinessOwner *string `json:"BusinessOwner,omitempty" name:"BusinessOwner"`

	// 评估模版Id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 评估模版名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 分类分级模版Id
	ComplianceGroupId *int64 `json:"ComplianceGroupId,omitempty" name:"ComplianceGroupId"`

	// 分类分级模版名称
	ComplianceGroupName *string `json:"ComplianceGroupName,omitempty" name:"ComplianceGroupName"`

	// 评估项数量
	ControlItemCount *int64 `json:"ControlItemCount,omitempty" name:"ControlItemCount"`

	// 风险项数量（仅状态为finished的风险项不计入总数，其余状态均算入该数量）
	RiskCount *int64 `json:"RiskCount,omitempty" name:"RiskCount"`

	// 评估任务完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedTime *string `json:"FinishedTime,omitempty" name:"FinishedTime"`

	// 评估任务发起时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 评估状态。(waiting待评估，processing评估中, , finished已评估, failed评估失败)
	Status *string `json:"Status,omitempty" name:"Status"`

	// 待处理各等级风险项信息
	RiskCountInfoList []*RiskCountInfo `json:"RiskCountInfoList,omitempty" name:"RiskCountInfoList"`

	// 数据源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscoveryCondition *DiscoveryCondition `json:"DiscoveryCondition,omitempty" name:"DiscoveryCondition"`

	// 评估任务失败信息
	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`
}

type AssessmentTemplate struct {
	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 评估模版Id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 评估模版名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 模版来源，内置/用户自定，取值（system，user）
	Source *string `json:"Source,omitempty" name:"Source"`

	// 模版类型，自动化/半自动化/问卷，取值（auto，semi-auto，law）等
	UseType *string `json:"UseType,omitempty" name:"UseType"`

	// 评估模版创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 模版关联的评估项数量
	ControlItemCount *int64 `json:"ControlItemCount,omitempty" name:"ControlItemCount"`

	// 模版已启用的评估项数量
	AppliedItemCount *int64 `json:"AppliedItemCount,omitempty" name:"AppliedItemCount"`

	// 模版启用状态，草稿/已启用，取值draft / launched
	Status *string `json:"Status,omitempty" name:"Status"`

	// 支持的数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportDataSource []*string `json:"SupportDataSource,omitempty" name:"SupportDataSource"`
}

type AssetCosDetail struct {
	// 桶的名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *string `json:"DataType,omitempty" name:"DataType"`

	// 文件的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNums *int64 `json:"FileNums,omitempty" name:"FileNums"`

	// 敏感的文件个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveFileNums *int64 `json:"SensitiveFileNums,omitempty" name:"SensitiveFileNums"`

	// 敏感分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributionData []*Note `json:"DistributionData,omitempty" name:"DistributionData"`

	// cos文件的敏感数据个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchedNum *int64 `json:"MatchedNum,omitempty" name:"MatchedNum"`
}

type AssetDBDetail struct {
	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdName *string `json:"DdName,omitempty" name:"DdName"`

	// 数据库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *string `json:"DataType,omitempty" name:"DataType"`

	// 表的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNums *int64 `json:"TableNums,omitempty" name:"TableNums"`

	// 敏感表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveTableNums *int64 `json:"SensitiveTableNums,omitempty" name:"SensitiveTableNums"`

	// 字段的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldNums *int64 `json:"FieldNums,omitempty" name:"FieldNums"`

	// 敏感字段的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveFieldNums *int64 `json:"SensitiveFieldNums,omitempty" name:"SensitiveFieldNums"`

	// 敏感数据分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributionData []*Note `json:"DistributionData,omitempty" name:"DistributionData"`
}

type AssetList struct {
	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 数据源信息
	DataSourceInfo []*DataSourceInfo `json:"DataSourceInfo,omitempty" name:"DataSourceInfo"`
}

// Predefined struct for user
type AuthorizeDSPAMetaResourcesRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 授权方式，可选：automatic(一键自动授权) 、 account(指定用户名授权)。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 资源类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 用户授权的账户信息，如果是一键自动授权模式，则不需要填写账户名与密码。
	ResourcesAccount []*DspaResourceAccount `json:"ResourcesAccount,omitempty" name:"ResourcesAccount"`
}

type AuthorizeDSPAMetaResourcesRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 授权方式，可选：automatic(一键自动授权) 、 account(指定用户名授权)。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 资源类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 用户授权的账户信息，如果是一键自动授权模式，则不需要填写账户名与密码。
	ResourcesAccount []*DspaResourceAccount `json:"ResourcesAccount,omitempty" name:"ResourcesAccount"`
}

func (r *AuthorizeDSPAMetaResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuthorizeDSPAMetaResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "AuthType")
	delete(f, "MetaType")
	delete(f, "ResourceRegion")
	delete(f, "ResourcesAccount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AuthorizeDSPAMetaResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuthorizeDSPAMetaResourcesResponseParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 授权结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*DspaTaskResult `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AuthorizeDSPAMetaResourcesResponse struct {
	*tchttp.BaseResponse
	Response *AuthorizeDSPAMetaResourcesResponseParams `json:"Response"`
}

func (r *AuthorizeDSPAMetaResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuthorizeDSPAMetaResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDSPAResourceCosBucketsRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 绑定的COS桶信息。
	BindCosResourceItems []*CosResourceItem `json:"BindCosResourceItems,omitempty" name:"BindCosResourceItems"`

	// 解绑的COS桶信息。
	UnbindCosResourceItems []*CosResourceItem `json:"UnbindCosResourceItems,omitempty" name:"UnbindCosResourceItems"`
}

type BindDSPAResourceCosBucketsRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 绑定的COS桶信息。
	BindCosResourceItems []*CosResourceItem `json:"BindCosResourceItems,omitempty" name:"BindCosResourceItems"`

	// 解绑的COS桶信息。
	UnbindCosResourceItems []*CosResourceItem `json:"UnbindCosResourceItems,omitempty" name:"UnbindCosResourceItems"`
}

func (r *BindDSPAResourceCosBucketsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDSPAResourceCosBucketsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "BindCosResourceItems")
	delete(f, "UnbindCosResourceItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindDSPAResourceCosBucketsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDSPAResourceCosBucketsResponseParams struct {
	// 绑定结果数组
	CosTaskResults []*CosTaskResult `json:"CosTaskResults,omitempty" name:"CosTaskResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindDSPAResourceCosBucketsResponse struct {
	*tchttp.BaseResponse
	Response *BindDSPAResourceCosBucketsResponseParams `json:"Response"`
}

func (r *BindDSPAResourceCosBucketsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDSPAResourceCosBucketsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDSPAResourceDatabasesRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据库实例ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 数据库实例类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 绑定DB列表。
	BindDbItems []*DbResourceItem `json:"BindDbItems,omitempty" name:"BindDbItems"`

	// 解绑DB列表。
	UnbindDbItems []*DbResourceItem `json:"UnbindDbItems,omitempty" name:"UnbindDbItems"`
}

type BindDSPAResourceDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据库实例ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 数据库实例类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 绑定DB列表。
	BindDbItems []*DbResourceItem `json:"BindDbItems,omitempty" name:"BindDbItems"`

	// 解绑DB列表。
	UnbindDbItems []*DbResourceItem `json:"UnbindDbItems,omitempty" name:"UnbindDbItems"`
}

func (r *BindDSPAResourceDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDSPAResourceDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ResourceId")
	delete(f, "MetaType")
	delete(f, "BindDbItems")
	delete(f, "UnbindDbItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindDSPAResourceDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDSPAResourceDatabasesResponseParams struct {
	// 绑定结果数组
	DbTaskResults []*DbTaskResult `json:"DbTaskResults,omitempty" name:"DbTaskResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindDSPAResourceDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *BindDSPAResourceDatabasesResponseParams `json:"Response"`
}

func (r *BindDSPAResourceDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDSPAResourceDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type COSDataRule struct {
	// 只能取and 、or两个值其中之一，and：字段和内容同时满足，or：字段和内容满足其一.
	// 默认值为or
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 规则内容
	Contents []*DataContent `json:"Contents,omitempty" name:"Contents"`
}

type COSInstance struct {
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 桶名
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
}

type CategoryRule struct {
	// 分类id
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 规则id
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 级别id
	LevelId *uint64 `json:"LevelId,omitempty" name:"LevelId"`

	// 级别名称
	LevelName *string `json:"LevelName,omitempty" name:"LevelName"`

	// 分类规则绑定关系id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 别名ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasRuleId *int64 `json:"AliasRuleId,omitempty" name:"AliasRuleId"`

	// 别名规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasRuleName *string `json:"AliasRuleName,omitempty" name:"AliasRuleName"`
}

type CategoryRuleStatistic struct {
	// 分类id
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 规则数量
	RuleCount *uint64 `json:"RuleCount,omitempty" name:"RuleCount"`

	// 分类名称
	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`
}

type ComplianceGroupDetail struct {
	// 模板id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 模板类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceGroupType *int64 `json:"ComplianceGroupType,omitempty" name:"ComplianceGroupType"`

	// 模板分级方案id
	LevelGroupId *uint64 `json:"LevelGroupId,omitempty" name:"LevelGroupId"`

	// 模板分级方案名称
	LevelGroupName *string `json:"LevelGroupName,omitempty" name:"LevelGroupName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 是否开启别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAlias *bool `json:"IsAlias,omitempty" name:"IsAlias"`
}

type ComplianceGroupRuleIdInfo struct {
	// 敏感数据识别规则ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 敏感数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分级标识ID, 系统支持高、中、低三级，也支持自定义分级
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`
}

// Predefined struct for user
type CopyDSPATemplateRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type CopyDSPATemplateRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *CopyDSPATemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyDSPATemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyDSPATemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyDSPATemplateResponseParams struct {
	// 模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CopyDSPATemplateResponse struct {
	*tchttp.BaseResponse
	Response *CopyDSPATemplateResponseParams `json:"Response"`
}

func (r *CopyDSPATemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyDSPATemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosAsset struct {
	// 桶的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketNums *int64 `json:"BucketNums,omitempty" name:"BucketNums"`

	// 敏感桶的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveBucketNums *int64 `json:"SensitiveBucketNums,omitempty" name:"SensitiveBucketNums"`

	// 文件个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNums *int64 `json:"FileNums,omitempty" name:"FileNums"`

	// 敏感文件的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveFileNums *int64 `json:"SensitiveFileNums,omitempty" name:"SensitiveFileNums"`
}

type CosResourceItem struct {
	// cos数据源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 桶所在地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 桶名称。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type CosTaskResult struct {
	// 结果类型。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 结果描述。
	ResultDescription *string `json:"ResultDescription,omitempty" name:"ResultDescription"`

	// 错误信息描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrDescription *ErrDescription `json:"ErrDescription,omitempty" name:"ErrDescription"`

	// 资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

// Predefined struct for user
type CreateAssetSortingReportRetryTaskRequestParams struct {
	// 任务id
	ReportTaskId *uint64 `json:"ReportTaskId,omitempty" name:"ReportTaskId"`

	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

type CreateAssetSortingReportRetryTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	ReportTaskId *uint64 `json:"ReportTaskId,omitempty" name:"ReportTaskId"`

	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

func (r *CreateAssetSortingReportRetryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetSortingReportRetryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReportTaskId")
	delete(f, "DspaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetSortingReportRetryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetSortingReportRetryTaskResponseParams struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTaskId *uint64 `json:"ReportTaskId,omitempty" name:"ReportTaskId"`

	// 提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAssetSortingReportRetryTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetSortingReportRetryTaskResponseParams `json:"Response"`
}

func (r *CreateAssetSortingReportRetryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetSortingReportRetryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetSortingReportTaskRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 识别模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 选中资产列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

type CreateAssetSortingReportTaskRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 识别模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 选中资产列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

func (r *CreateAssetSortingReportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetSortingReportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "AssetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetSortingReportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetSortingReportTaskResponseParams struct {
	// 报表任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTaskId *uint64 `json:"ReportTaskId,omitempty" name:"ReportTaskId"`

	// 提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAssetSortingReportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetSortingReportTaskResponseParams `json:"Response"`
}

func (r *CreateAssetSortingReportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetSortingReportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClassificationRuleRequestParams struct {

}

type CreateClassificationRuleRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateClassificationRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClassificationRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClassificationRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClassificationRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClassificationRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateClassificationRuleResponseParams `json:"Response"`
}

func (r *CreateClassificationRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClassificationRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateComplianceRules struct {
	// 规则id
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 级别id
	LevelId *uint64 `json:"LevelId,omitempty" name:"LevelId"`
}

// Predefined struct for user
type CreateDSPAAssessmentRiskLevelRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 风险等级名称
	RiskLevelName *string `json:"RiskLevelName,omitempty" name:"RiskLevelName"`

	// 识别模版
	IdentifyComplianceId *int64 `json:"IdentifyComplianceId,omitempty" name:"IdentifyComplianceId"`

	// 风险等级矩阵
	RiskLevelRule []*RiskLevelMatrix `json:"RiskLevelRule,omitempty" name:"RiskLevelRule"`

	// 风险等级的描述
	RiskLevelDescription *string `json:"RiskLevelDescription,omitempty" name:"RiskLevelDescription"`
}

type CreateDSPAAssessmentRiskLevelRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 风险等级名称
	RiskLevelName *string `json:"RiskLevelName,omitempty" name:"RiskLevelName"`

	// 识别模版
	IdentifyComplianceId *int64 `json:"IdentifyComplianceId,omitempty" name:"IdentifyComplianceId"`

	// 风险等级矩阵
	RiskLevelRule []*RiskLevelMatrix `json:"RiskLevelRule,omitempty" name:"RiskLevelRule"`

	// 风险等级的描述
	RiskLevelDescription *string `json:"RiskLevelDescription,omitempty" name:"RiskLevelDescription"`
}

func (r *CreateDSPAAssessmentRiskLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAAssessmentRiskLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "RiskLevelName")
	delete(f, "IdentifyComplianceId")
	delete(f, "RiskLevelRule")
	delete(f, "RiskLevelDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPAAssessmentRiskLevelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAAssessmentRiskLevelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPAAssessmentRiskLevelResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPAAssessmentRiskLevelResponseParams `json:"Response"`
}

func (r *CreateDSPAAssessmentRiskLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAAssessmentRiskLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAAssessmentRiskTemplateRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 模版名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 风险等级id
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" name:"RiskLevelId"`

	// 风险id列表
	RiskIdList []*int64 `json:"RiskIdList,omitempty" name:"RiskIdList"`

	// 模版描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`
}

type CreateDSPAAssessmentRiskTemplateRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 模版名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 风险等级id
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" name:"RiskLevelId"`

	// 风险id列表
	RiskIdList []*int64 `json:"RiskIdList,omitempty" name:"RiskIdList"`

	// 模版描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`
}

func (r *CreateDSPAAssessmentRiskTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAAssessmentRiskTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateName")
	delete(f, "RiskLevelId")
	delete(f, "RiskIdList")
	delete(f, "TemplateDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPAAssessmentRiskTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAAssessmentRiskTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPAAssessmentRiskTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPAAssessmentRiskTemplateResponseParams `json:"Response"`
}

func (r *CreateDSPAAssessmentRiskTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAAssessmentRiskTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAAssessmentTaskRequestParams struct {
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估任务名称。1-20个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 评估模版Id，格式“template-xxxxxxxx”
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 敏感数据扫描数据源条件。
	DiscoveryCondition *DiscoveryCondition `json:"DiscoveryCondition,omitempty" name:"DiscoveryCondition"`

	// 评估业务名称。1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字
	//
	// Deprecated: BusinessName is deprecated.
	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`

	// 业务所属部门。1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字
	//
	// Deprecated: BusinessDept is deprecated.
	BusinessDept *string `json:"BusinessDept,omitempty" name:"BusinessDept"`

	// 业务负责人。1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字
	//
	// Deprecated: BusinessOwner is deprecated.
	BusinessOwner *string `json:"BusinessOwner,omitempty" name:"BusinessOwner"`

	// 分类分级模版Id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 说明
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateDSPAAssessmentTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估任务名称。1-20个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 评估模版Id，格式“template-xxxxxxxx”
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 敏感数据扫描数据源条件。
	DiscoveryCondition *DiscoveryCondition `json:"DiscoveryCondition,omitempty" name:"DiscoveryCondition"`

	// 评估业务名称。1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字
	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`

	// 业务所属部门。1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字
	BusinessDept *string `json:"BusinessDept,omitempty" name:"BusinessDept"`

	// 业务负责人。1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字
	BusinessOwner *string `json:"BusinessOwner,omitempty" name:"BusinessOwner"`

	// 分类分级模版Id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 说明
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateDSPAAssessmentTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAAssessmentTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Name")
	delete(f, "TemplateId")
	delete(f, "DiscoveryCondition")
	delete(f, "BusinessName")
	delete(f, "BusinessDept")
	delete(f, "BusinessOwner")
	delete(f, "ComplianceId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPAAssessmentTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAAssessmentTaskResponseParams struct {
	// 评估任务Id，格式“task-xxxxxxxx”
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPAAssessmentTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPAAssessmentTaskResponseParams `json:"Response"`
}

func (r *CreateDSPAAssessmentTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAAssessmentTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPACOSDiscoveryTaskRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 任务开关，0 关闭，1 启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 桶名
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 通用规则集开关，0 关闭，1 启用
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 执行计划， 0立即 1定时，选择“立即”时，扫描周期只能选择单次。
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 扫描周期，0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 待扫描文件类型，用逗号隔开，格式如：[".txt", ".csv", ".log", ".xml",".html", ".json"]。
	FileTypes []*string `json:"FileTypes,omitempty" name:"FileTypes"`

	// 文件大小上限，单位为KB，如1000, 目前单个文件最大只支持100MB（102400KB）
	FileSizeLimit *int64 `json:"FileSizeLimit,omitempty" name:"FileSizeLimit"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 任务描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 合规组ID列表，最多支持添加5个
	ComplianceGroupIds []*int64 `json:"ComplianceGroupIds,omitempty" name:"ComplianceGroupIds"`

	// 任务定时启动时间，格式如：2006-01-02 15:04:05
	// 当执行计划（Plan字段）为”立即“时，定时启动时间不会生效，此场景下给该字段传值不会被保存。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`
}

type CreateDSPACOSDiscoveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 任务开关，0 关闭，1 启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 桶名
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 通用规则集开关，0 关闭，1 启用
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 执行计划， 0立即 1定时，选择“立即”时，扫描周期只能选择单次。
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 扫描周期，0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 待扫描文件类型，用逗号隔开，格式如：[".txt", ".csv", ".log", ".xml",".html", ".json"]。
	FileTypes []*string `json:"FileTypes,omitempty" name:"FileTypes"`

	// 文件大小上限，单位为KB，如1000, 目前单个文件最大只支持100MB（102400KB）
	FileSizeLimit *int64 `json:"FileSizeLimit,omitempty" name:"FileSizeLimit"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 任务描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 合规组ID列表，最多支持添加5个
	ComplianceGroupIds []*int64 `json:"ComplianceGroupIds,omitempty" name:"ComplianceGroupIds"`

	// 任务定时启动时间，格式如：2006-01-02 15:04:05
	// 当执行计划（Plan字段）为”立即“时，定时启动时间不会生效，此场景下给该字段传值不会被保存。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`
}

func (r *CreateDSPACOSDiscoveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPACOSDiscoveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Name")
	delete(f, "DataSourceId")
	delete(f, "Enable")
	delete(f, "Bucket")
	delete(f, "GeneralRuleSetEnable")
	delete(f, "Plan")
	delete(f, "Period")
	delete(f, "FileTypes")
	delete(f, "FileSizeLimit")
	delete(f, "ResourceRegion")
	delete(f, "Description")
	delete(f, "ComplianceGroupIds")
	delete(f, "TimingStartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPACOSDiscoveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPACOSDiscoveryTaskResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 扫描结果ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultId *int64 `json:"ResultId,omitempty" name:"ResultId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPACOSDiscoveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPACOSDiscoveryTaskResponseParams `json:"Response"`
}

func (r *CreateDSPACOSDiscoveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPACOSDiscoveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPACategoryRelationRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分类id
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 父级分类id（无父级分类传-1）
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" name:"ParentCategoryId"`

	// 分类模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type CreateDSPACategoryRelationRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分类id
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 父级分类id（无父级分类传-1）
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" name:"ParentCategoryId"`

	// 分类模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *CreateDSPACategoryRelationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPACategoryRelationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "CategoryId")
	delete(f, "ParentCategoryId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPACategoryRelationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPACategoryRelationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPACategoryRelationResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPACategoryRelationResponseParams `json:"Response"`
}

func (r *CreateDSPACategoryRelationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPACategoryRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPACategoryRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 敏感数据分类名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateDSPACategoryRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 敏感数据分类名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateDSPACategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPACategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPACategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPACategoryResponseParams struct {
	// 敏感数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPACategoryResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPACategoryResponseParams `json:"Response"`
}

func (r *CreateDSPACategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPACategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAComplianceGroupRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 合规组描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 合规组规则配置（参数已废弃，请传空数组）
	ComplianceGroupRules []*ComplianceGroupRuleIdInfo `json:"ComplianceGroupRules,omitempty" name:"ComplianceGroupRules"`

	// 分级组ID，默认值为1，新增参数，可选
	LevelGroupId *uint64 `json:"LevelGroupId,omitempty" name:"LevelGroupId"`
}

type CreateDSPAComplianceGroupRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 合规组描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 合规组规则配置（参数已废弃，请传空数组）
	ComplianceGroupRules []*ComplianceGroupRuleIdInfo `json:"ComplianceGroupRules,omitempty" name:"ComplianceGroupRules"`

	// 分级组ID，默认值为1，新增参数，可选
	LevelGroupId *uint64 `json:"LevelGroupId,omitempty" name:"LevelGroupId"`
}

func (r *CreateDSPAComplianceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAComplianceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ComplianceGroupRules")
	delete(f, "LevelGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPAComplianceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAComplianceGroupResponseParams struct {
	// 合规组ID
	ComplianceGroupId *int64 `json:"ComplianceGroupId,omitempty" name:"ComplianceGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPAComplianceGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPAComplianceGroupResponseParams `json:"Response"`
}

func (r *CreateDSPAComplianceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAComplianceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAComplianceRulesRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分类id
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 合规组模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 规则列表
	Rules []*CreateComplianceRules `json:"Rules,omitempty" name:"Rules"`
}

type CreateDSPAComplianceRulesRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分类id
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 合规组模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 规则列表
	Rules []*CreateComplianceRules `json:"Rules,omitempty" name:"Rules"`
}

func (r *CreateDSPAComplianceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAComplianceRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "CategoryId")
	delete(f, "ComplianceId")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPAComplianceRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAComplianceRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPAComplianceRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPAComplianceRulesResponseParams `json:"Response"`
}

func (r *CreateDSPAComplianceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAComplianceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPACosMetaResourcesRequestParams struct {
	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// COS桶列表
	Buckets []*string `json:"Buckets,omitempty" name:"Buckets"`
}

type CreateDSPACosMetaResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// COS桶列表
	Buckets []*string `json:"Buckets,omitempty" name:"Buckets"`
}

func (r *CreateDSPACosMetaResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPACosMetaResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceRegion")
	delete(f, "DspaId")
	delete(f, "Buckets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPACosMetaResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPACosMetaResourcesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPACosMetaResourcesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPACosMetaResourcesResponseParams `json:"Response"`
}

func (r *CreateDSPACosMetaResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPACosMetaResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPADbMetaResourcesRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源类型，支持：cdb（云数据库 MySQL）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbpg（TDSQL-C PostgreSQL版）、cynosdbmysql（TDSQL-C MySQL版）
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 用来标记本次更新是否已经是最后一次，可选值：continue（后续还需要更新）、finished（本次是最后一次更新）。
	UpdateStatus *string `json:"UpdateStatus,omitempty" name:"UpdateStatus"`

	// 本次更新的ID号，用来标记一次完整的更新过程。
	UpdateId *string `json:"UpdateId,omitempty" name:"UpdateId"`

	// 云上资源列表。
	Items []*DspaCloudResourceMeta `json:"Items,omitempty" name:"Items"`
}

type CreateDSPADbMetaResourcesRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源类型，支持：cdb（云数据库 MySQL）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbpg（TDSQL-C PostgreSQL版）、cynosdbmysql（TDSQL-C MySQL版）
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 用来标记本次更新是否已经是最后一次，可选值：continue（后续还需要更新）、finished（本次是最后一次更新）。
	UpdateStatus *string `json:"UpdateStatus,omitempty" name:"UpdateStatus"`

	// 本次更新的ID号，用来标记一次完整的更新过程。
	UpdateId *string `json:"UpdateId,omitempty" name:"UpdateId"`

	// 云上资源列表。
	Items []*DspaCloudResourceMeta `json:"Items,omitempty" name:"Items"`
}

func (r *CreateDSPADbMetaResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPADbMetaResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "MetaType")
	delete(f, "ResourceRegion")
	delete(f, "UpdateStatus")
	delete(f, "UpdateId")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPADbMetaResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPADbMetaResourcesResponseParams struct {
	// 本次更新的ID号，用来标记一次完整的更新过程。
	UpdateId *string `json:"UpdateId,omitempty" name:"UpdateId"`

	// 资源类型，支持：cdb（云数据库 MySQL）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbpg（TDSQL-C PostgreSQL版）、cynosdbmysql（TDSQL-C MySQL版）
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPADbMetaResourcesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPADbMetaResourcesResponseParams `json:"Response"`
}

func (r *CreateDSPADbMetaResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPADbMetaResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPADiscoveryRuleRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 规则名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// RDB类敏感数据识别规则
	RDBRules *DspaDiscoveryRDBRules `json:"RDBRules,omitempty" name:"RDBRules"`

	// COS类敏感数据识别规则
	COSRules *DspaDiscoveryCOSRules `json:"COSRules,omitempty" name:"COSRules"`
}

type CreateDSPADiscoveryRuleRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 规则名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// RDB类敏感数据识别规则
	RDBRules *DspaDiscoveryRDBRules `json:"RDBRules,omitempty" name:"RDBRules"`

	// COS类敏感数据识别规则
	COSRules *DspaDiscoveryCOSRules `json:"COSRules,omitempty" name:"COSRules"`
}

func (r *CreateDSPADiscoveryRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPADiscoveryRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "RDBRules")
	delete(f, "COSRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPADiscoveryRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPADiscoveryRuleResponseParams struct {
	// 规则ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPADiscoveryRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPADiscoveryRuleResponseParams `json:"Response"`
}

func (r *CreateDSPADiscoveryRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPADiscoveryRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPADiscoveryTaskRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 任务开关，0 关闭，1 启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 通用规则集开关，0 关闭，1 启用
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 执行计划， 0立即 1定时，选择“立即”时，扫描周期只能选择单次
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 扫描周期，0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 数据源类型，可取值如下：
	// cdb 表示云数据库 MySQL,
	// dcdb 表示TDSQL MySQL版,
	// mariadb 表示云数据库 MariaDB,
	// postgres 表示云数据库 PostgreSQL,
	// cynosdbpg 表示TDSQL-C PostgreSQL版,
	// cynosdbmysql 表示TDSQL-C MySQL版,
	// selfbuilt-db 表示自建数据库
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 任务描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用于传入的数据源的条件，目前只支持数据库，所以目前表示数据库的名称，选择多个数据库，之间通过逗号分隔，若不选，则默认选择全部数据库
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 合规组ID列表，最多支持添加5个
	ComplianceGroupIds []*int64 `json:"ComplianceGroupIds,omitempty" name:"ComplianceGroupIds"`

	// 任务定时启动时间，格式如：2006-01-02 15:04:05
	// 当执行计划（Plan字段）为”立即“时，定时启动时间不会生效，此场景下给该字段传值不会被保存。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`
}

type CreateDSPADiscoveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 任务开关，0 关闭，1 启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 通用规则集开关，0 关闭，1 启用
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 执行计划， 0立即 1定时，选择“立即”时，扫描周期只能选择单次
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 扫描周期，0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 数据源类型，可取值如下：
	// cdb 表示云数据库 MySQL,
	// dcdb 表示TDSQL MySQL版,
	// mariadb 表示云数据库 MariaDB,
	// postgres 表示云数据库 PostgreSQL,
	// cynosdbpg 表示TDSQL-C PostgreSQL版,
	// cynosdbmysql 表示TDSQL-C MySQL版,
	// selfbuilt-db 表示自建数据库
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 任务描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用于传入的数据源的条件，目前只支持数据库，所以目前表示数据库的名称，选择多个数据库，之间通过逗号分隔，若不选，则默认选择全部数据库
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 合规组ID列表，最多支持添加5个
	ComplianceGroupIds []*int64 `json:"ComplianceGroupIds,omitempty" name:"ComplianceGroupIds"`

	// 任务定时启动时间，格式如：2006-01-02 15:04:05
	// 当执行计划（Plan字段）为”立即“时，定时启动时间不会生效，此场景下给该字段传值不会被保存。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`
}

func (r *CreateDSPADiscoveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPADiscoveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Name")
	delete(f, "DataSourceId")
	delete(f, "Enable")
	delete(f, "GeneralRuleSetEnable")
	delete(f, "Plan")
	delete(f, "Period")
	delete(f, "ResourceRegion")
	delete(f, "DataSourceType")
	delete(f, "Description")
	delete(f, "Condition")
	delete(f, "ComplianceGroupIds")
	delete(f, "TimingStartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPADiscoveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPADiscoveryTaskResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 扫描结果ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultId *int64 `json:"ResultId,omitempty" name:"ResultId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPADiscoveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPADiscoveryTaskResponseParams `json:"Response"`
}

func (r *CreateDSPADiscoveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPADiscoveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPALevelGroupRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分级组名称，唯一性约束，最多60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分级标识配置
	ItemLevels []*ItemLevel `json:"ItemLevels,omitempty" name:"ItemLevels"`

	// 分级组描述，最多1024字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateDSPALevelGroupRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分级组名称，唯一性约束，最多60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分级标识配置
	ItemLevels []*ItemLevel `json:"ItemLevels,omitempty" name:"ItemLevels"`

	// 分级组描述，最多1024字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateDSPALevelGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPALevelGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Name")
	delete(f, "ItemLevels")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPALevelGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPALevelGroupResponseParams struct {
	// 分级组ID
	LevelGroupId *uint64 `json:"LevelGroupId,omitempty" name:"LevelGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPALevelGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPALevelGroupResponseParams `json:"Response"`
}

func (r *CreateDSPALevelGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPALevelGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAMetaResourcesRequestParams struct {
	// 资源类型，支持：cdb（云数据库 MySQL）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbpg（TDSQL-C PostgreSQL版）、cynosdbmysql（TDSQL-C MySQL版）
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 用来标记本次更新是否已经是最后一次，可选值：continue（后续还需要更新）、finished（本次是最后一次更新）。
	UpdateStatus *string `json:"UpdateStatus,omitempty" name:"UpdateStatus"`

	// 本次更新的ID号，用来标记一次完整的更新过程。
	UpdateId *string `json:"UpdateId,omitempty" name:"UpdateId"`

	// 资源列表。
	Items []*DspaUserResourceMeta `json:"Items,omitempty" name:"Items"`
}

type CreateDSPAMetaResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 资源类型，支持：cdb（云数据库 MySQL）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbpg（TDSQL-C PostgreSQL版）、cynosdbmysql（TDSQL-C MySQL版）
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 用来标记本次更新是否已经是最后一次，可选值：continue（后续还需要更新）、finished（本次是最后一次更新）。
	UpdateStatus *string `json:"UpdateStatus,omitempty" name:"UpdateStatus"`

	// 本次更新的ID号，用来标记一次完整的更新过程。
	UpdateId *string `json:"UpdateId,omitempty" name:"UpdateId"`

	// 资源列表。
	Items []*DspaUserResourceMeta `json:"Items,omitempty" name:"Items"`
}

func (r *CreateDSPAMetaResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAMetaResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MetaType")
	delete(f, "ResourceRegion")
	delete(f, "DspaId")
	delete(f, "UpdateStatus")
	delete(f, "UpdateId")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPAMetaResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPAMetaResourcesResponseParams struct {
	// 本次更新的ID号，用来标记一次完整的更新过程。
	UpdateId *string `json:"UpdateId,omitempty" name:"UpdateId"`

	// 资源类型，支持：cdb（云数据库 MySQL）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbpg（TDSQL-C PostgreSQL版）、cynosdbmysql（TDSQL-C MySQL版）
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPAMetaResourcesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPAMetaResourcesResponseParams `json:"Response"`
}

func (r *CreateDSPAMetaResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPAMetaResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPASelfBuildMetaResourceRequestParams struct {
	// Dspa实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 自建数据库类型。目前支持的自建数据库类型按照协议进行区分，支持两种开源数据库协议：
	// mysql_like_proto -- Mysql协议类关系型数据库，
	// postgre_like_proto -- Postgre协议类关系型数据库。
	// 其他闭源协议的数据库如SqlServer、Oracle等暂不支持。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 自建云资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 可用于访问自建云资源的IP。
	ResourceVip *string `json:"ResourceVip,omitempty" name:"ResourceVip"`

	// 可用于访问自建云资源的端口。
	ResourceVPort *uint64 `json:"ResourceVPort,omitempty" name:"ResourceVPort"`

	// 自建云资源的VPC ID。
	ResourceUniqueVpcId *string `json:"ResourceUniqueVpcId,omitempty" name:"ResourceUniqueVpcId"`

	// 自建云资源的Subnet ID。
	ResourceUniqueSubnetId *string `json:"ResourceUniqueSubnetId,omitempty" name:"ResourceUniqueSubnetId"`

	// 自建云资源所处的服务类型，可选：
	// cvm - 通过云服务器直接访问。
	// clb - 通过LB的方式进行访问。
	ResourceAccessType *string `json:"ResourceAccessType,omitempty" name:"ResourceAccessType"`

	// 账户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 账户密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 资源名称，1-60个字符。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type CreateDSPASelfBuildMetaResourceRequest struct {
	*tchttp.BaseRequest
	
	// Dspa实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 自建数据库类型。目前支持的自建数据库类型按照协议进行区分，支持两种开源数据库协议：
	// mysql_like_proto -- Mysql协议类关系型数据库，
	// postgre_like_proto -- Postgre协议类关系型数据库。
	// 其他闭源协议的数据库如SqlServer、Oracle等暂不支持。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 自建云资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 可用于访问自建云资源的IP。
	ResourceVip *string `json:"ResourceVip,omitempty" name:"ResourceVip"`

	// 可用于访问自建云资源的端口。
	ResourceVPort *uint64 `json:"ResourceVPort,omitempty" name:"ResourceVPort"`

	// 自建云资源的VPC ID。
	ResourceUniqueVpcId *string `json:"ResourceUniqueVpcId,omitempty" name:"ResourceUniqueVpcId"`

	// 自建云资源的Subnet ID。
	ResourceUniqueSubnetId *string `json:"ResourceUniqueSubnetId,omitempty" name:"ResourceUniqueSubnetId"`

	// 自建云资源所处的服务类型，可选：
	// cvm - 通过云服务器直接访问。
	// clb - 通过LB的方式进行访问。
	ResourceAccessType *string `json:"ResourceAccessType,omitempty" name:"ResourceAccessType"`

	// 账户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 账户密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 资源名称，1-60个字符。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

func (r *CreateDSPASelfBuildMetaResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPASelfBuildMetaResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "MetaType")
	delete(f, "ResourceRegion")
	delete(f, "ResourceId")
	delete(f, "ResourceVip")
	delete(f, "ResourceVPort")
	delete(f, "ResourceUniqueVpcId")
	delete(f, "ResourceUniqueSubnetId")
	delete(f, "ResourceAccessType")
	delete(f, "UserName")
	delete(f, "Password")
	delete(f, "ResourceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPASelfBuildMetaResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPASelfBuildMetaResourceResponseParams struct {
	// 连通性测试情况，success表示可正常访问，failed表示无法访问。
	ConnectivityStatus *string `json:"ConnectivityStatus,omitempty" name:"ConnectivityStatus"`

	// 连通性描述字段，如果连通性测试失败，这里会返回无法访问的相关信息说明。
	ConnectivityDescription *string `json:"ConnectivityDescription,omitempty" name:"ConnectivityDescription"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPASelfBuildMetaResourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPASelfBuildMetaResourceResponseParams `json:"Response"`
}

func (r *CreateDSPASelfBuildMetaResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPASelfBuildMetaResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIdentifyRuleAnotherNameRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 规则绑定的分类id
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则别名
	AnotherName *string `json:"AnotherName,omitempty" name:"AnotherName"`

	// 别名规则id
	AliasRuleId *int64 `json:"AliasRuleId,omitempty" name:"AliasRuleId"`

	// 别名规则名称
	AliasRuleName *string `json:"AliasRuleName,omitempty" name:"AliasRuleName"`
}

type CreateIdentifyRuleAnotherNameRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 规则绑定的分类id
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则别名
	AnotherName *string `json:"AnotherName,omitempty" name:"AnotherName"`

	// 别名规则id
	AliasRuleId *int64 `json:"AliasRuleId,omitempty" name:"AliasRuleId"`

	// 别名规则名称
	AliasRuleName *string `json:"AliasRuleName,omitempty" name:"AliasRuleName"`
}

func (r *CreateIdentifyRuleAnotherNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIdentifyRuleAnotherNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "CategoryId")
	delete(f, "RuleId")
	delete(f, "RuleName")
	delete(f, "AnotherName")
	delete(f, "AliasRuleId")
	delete(f, "AliasRuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIdentifyRuleAnotherNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIdentifyRuleAnotherNameResponseParams struct {
	// 创建的别名规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasRuleId *int64 `json:"AliasRuleId,omitempty" name:"AliasRuleId"`

	// 别名规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasRuleName *string `json:"AliasRuleName,omitempty" name:"AliasRuleName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIdentifyRuleAnotherNameResponse struct {
	*tchttp.BaseResponse
	Response *CreateIdentifyRuleAnotherNameResponseParams `json:"Response"`
}

func (r *CreateIdentifyRuleAnotherNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIdentifyRuleAnotherNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewClassificationRequestParams struct {

}

type CreateNewClassificationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateNewClassificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewClassificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNewClassificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewClassificationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNewClassificationResponse struct {
	*tchttp.BaseResponse
	Response *CreateNewClassificationResponseParams `json:"Response"`
}

func (r *CreateNewClassificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewClassificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrCopyStandardRequestParams struct {

}

type CreateOrCopyStandardRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateOrCopyStandardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrCopyStandardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrCopyStandardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrCopyStandardResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOrCopyStandardResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrCopyStandardResponseParams `json:"Response"`
}

func (r *CreateOrCopyStandardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrCopyStandardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBInstanceInfo struct {
	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 数据源绑定的db信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbInfos []*DbInfo `json:"DbInfos,omitempty" name:"DbInfos"`
}

type DBStatements struct {
	// 数据库名称
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// 数据库Schema
	DBSchema *string `json:"DBSchema,omitempty" name:"DBSchema"`
}

type DSPACosMetaDataInfo struct {
	// COS桶名
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// COS桶创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 1 -- 有效，0 -- 无效，资源可能已被删除。
	Valid *int64 `json:"Valid,omitempty" name:"Valid"`

	// DSPA为COS资源生成的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// COS资源所处的地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// COS桶绑定状态
	BindStatus *string `json:"BindStatus,omitempty" name:"BindStatus"`

	// COS桶存储量
	Storage *float64 `json:"Storage,omitempty" name:"Storage"`
}

type DSPADataSourceDbInfo struct {
	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

type DSPAMetaType struct {
	// 元数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 支持的此元数据类型的地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regions []*string `json:"Regions,omitempty" name:"Regions"`

	// 此元数据类型支持的授权类型：
	// account    -- 账户名密码授权，账户的最高只读权限需要由用户自行赋予；
	// automatic -- 一键授权，由DSPA自动生成账户名密码并自动在实例中给账户名赋予最高只读权限；
	// 如果此列表为空，表明此类资源不支持以上的授权机制，无法通过后台进行授权。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportedAuthTypes []*string `json:"SupportedAuthTypes,omitempty" name:"SupportedAuthTypes"`
}

type DSPATableInfo struct {
	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`
}

type DataCategory struct {
	// 分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分类名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 敏感数据分类来源，取值：0 内置, 1 自定义
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 关联模板数量
	RelateComplianceCount *uint64 `json:"RelateComplianceCount,omitempty" name:"RelateComplianceCount"`
}

type DataContent struct {
	// 规则内容，可以是正则规则，关键词，
	// 忽略词扥
	RuleContent *string `json:"RuleContent,omitempty" name:"RuleContent"`

	// 是否区分大小写
	// false: 不区分大小写
	// true:区分大小写
	IsIgnoreCase *bool `json:"IsIgnoreCase,omitempty" name:"IsIgnoreCase"`
}

type DataRule struct {
	// 规则类型；取值：
	// keyword 关键字, 
	// regex 正则
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 内容
	RuleContent *string `json:"RuleContent,omitempty" name:"RuleContent"`

	// 该字段是针对规则类型RuleType为keyword类型时的一个扩展属性
	ExtendParameters []*DatagovRuleExtendParameter `json:"ExtendParameters,omitempty" name:"ExtendParameters"`
}

type DataRules struct {
	// 操作符；只能取and, or的其中一种
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contents []*DataRule `json:"Contents,omitempty" name:"Contents"`
}

type DataSourceInfo struct {
	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 针对rbd-就是绑定的db_name
	BindList []*string `json:"BindList,omitempty" name:"BindList"`
}

type DatagovRuleExtendParameter struct {
	// 扩展参数名称，目前支持如下两个扩展属性名称：
	// IsFullWordMatch，表示是否全文匹配，该Name对应的Value可取值为"true"或"false":，默认值为"false"，
	// IsIgnoreCase，表示是否忽略大小写，该Name对应的Value可取值为"true"或"false"，默认值为"true"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 扩展参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DbInfo struct {
	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 绑定的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidStatus *string `json:"ValidStatus,omitempty" name:"ValidStatus"`
}

type DbRelationStatusItem struct {
	// DB名称。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// DB绑定状态。
	BindStatus *string `json:"BindStatus,omitempty" name:"BindStatus"`

	// DB有效性状态。
	ValidStatus *string `json:"ValidStatus,omitempty" name:"ValidStatus"`
}

type DbResourceItem struct {
	// DB名称。
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

type DbTaskResult struct {
	// 结果类型。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 结果描述。
	ResultDescription *string `json:"ResultDescription,omitempty" name:"ResultDescription"`

	// 错误信息描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrDescription *ErrDescription `json:"ErrDescription,omitempty" name:"ErrDescription"`

	// 资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// database名称。
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

// Predefined struct for user
type DecribeSuggestRiskLevelMatrixRequestParams struct {
	// dspaId
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分类分级等级列表
	SensitiveLevelList []*RiskMatrixLevel `json:"SensitiveLevelList,omitempty" name:"SensitiveLevelList"`

	// 脆弱项等级列表
	VulnerabilityLevelList []*RiskMatrixLevel `json:"VulnerabilityLevelList,omitempty" name:"VulnerabilityLevelList"`
}

type DecribeSuggestRiskLevelMatrixRequest struct {
	*tchttp.BaseRequest
	
	// dspaId
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分类分级等级列表
	SensitiveLevelList []*RiskMatrixLevel `json:"SensitiveLevelList,omitempty" name:"SensitiveLevelList"`

	// 脆弱项等级列表
	VulnerabilityLevelList []*RiskMatrixLevel `json:"VulnerabilityLevelList,omitempty" name:"VulnerabilityLevelList"`
}

func (r *DecribeSuggestRiskLevelMatrixRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DecribeSuggestRiskLevelMatrixRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "SensitiveLevelList")
	delete(f, "VulnerabilityLevelList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DecribeSuggestRiskLevelMatrixRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DecribeSuggestRiskLevelMatrixResponseParams struct {
	// 矩阵
	SuggestRiskLevelMatrix []*SuggestRiskLevelMatrix `json:"SuggestRiskLevelMatrix,omitempty" name:"SuggestRiskLevelMatrix"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DecribeSuggestRiskLevelMatrixResponse struct {
	*tchttp.BaseResponse
	Response *DecribeSuggestRiskLevelMatrixResponseParams `json:"Response"`
}

func (r *DecribeSuggestRiskLevelMatrixResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DecribeSuggestRiskLevelMatrixResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCosMetaResourceRequestParams struct {
	// 实例Id。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源实例ID。
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

type DeleteCosMetaResourceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源实例ID。
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

func (r *DeleteCosMetaResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCosMetaResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ResourceRegion")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCosMetaResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCosMetaResourceResponseParams struct {
	// 结果集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*DspaTaskResult `json:"Results,omitempty" name:"Results"`

	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCosMetaResourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCosMetaResourceResponseParams `json:"Response"`
}

func (r *DeleteCosMetaResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCosMetaResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPAAssessmentTaskRequestParams struct {
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估任务Id，格式“task-xxxxxxxx”
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DeleteDSPAAssessmentTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估任务Id，格式“task-xxxxxxxx”
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteDSPAAssessmentTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPAAssessmentTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDSPAAssessmentTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPAAssessmentTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDSPAAssessmentTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDSPAAssessmentTaskResponseParams `json:"Response"`
}

func (r *DeleteDSPAAssessmentTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPAAssessmentTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPACOSDiscoveryTaskRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DeleteDSPACOSDiscoveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteDSPACOSDiscoveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPACOSDiscoveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDSPACOSDiscoveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPACOSDiscoveryTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDSPACOSDiscoveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDSPACOSDiscoveryTaskResponseParams `json:"Response"`
}

func (r *DeleteDSPACOSDiscoveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPACOSDiscoveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPACOSDiscoveryTaskResultRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 扫描bucket结果ID
	BucketResultId *int64 `json:"BucketResultId,omitempty" name:"BucketResultId"`
}

type DeleteDSPACOSDiscoveryTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 扫描bucket结果ID
	BucketResultId *int64 `json:"BucketResultId,omitempty" name:"BucketResultId"`
}

func (r *DeleteDSPACOSDiscoveryTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPACOSDiscoveryTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "BucketResultId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDSPACOSDiscoveryTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPACOSDiscoveryTaskResultResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDSPACOSDiscoveryTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDSPACOSDiscoveryTaskResultResponseParams `json:"Response"`
}

func (r *DeleteDSPACOSDiscoveryTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPACOSDiscoveryTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPADiscoveryTaskRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`
}

type DeleteDSPADiscoveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`
}

func (r *DeleteDSPADiscoveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPADiscoveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "DataSourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDSPADiscoveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPADiscoveryTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDSPADiscoveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDSPADiscoveryTaskResponseParams `json:"Response"`
}

func (r *DeleteDSPADiscoveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPADiscoveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPADiscoveryTaskResultRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 扫描数据库结果ID
	DbResultId *int64 `json:"DbResultId,omitempty" name:"DbResultId"`
}

type DeleteDSPADiscoveryTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 扫描数据库结果ID
	DbResultId *int64 `json:"DbResultId,omitempty" name:"DbResultId"`
}

func (r *DeleteDSPADiscoveryTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPADiscoveryTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "DbResultId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDSPADiscoveryTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPADiscoveryTaskResultResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDSPADiscoveryTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDSPADiscoveryTaskResultResponseParams `json:"Response"`
}

func (r *DeleteDSPADiscoveryTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPADiscoveryTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPAMetaResourceRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 用户云资源ID。
	ResourceIDs []*string `json:"ResourceIDs,omitempty" name:"ResourceIDs"`
}

type DeleteDSPAMetaResourceRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 用户云资源ID。
	ResourceIDs []*string `json:"ResourceIDs,omitempty" name:"ResourceIDs"`
}

func (r *DeleteDSPAMetaResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPAMetaResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ResourceIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDSPAMetaResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDSPAMetaResourceResponseParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 删除结果。
	Results []*DspaTaskResult `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDSPAMetaResourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDSPAMetaResourceResponseParams `json:"Response"`
}

func (r *DeleteDSPAMetaResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDSPAMetaResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDetailDataExportResultRequestParams struct {
	// 导出任务id
	ExportTaskId *uint64 `json:"ExportTaskId,omitempty" name:"ExportTaskId"`

	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

type DescribeAssetDetailDataExportResultRequest struct {
	*tchttp.BaseRequest
	
	// 导出任务id
	ExportTaskId *uint64 `json:"ExportTaskId,omitempty" name:"ExportTaskId"`

	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

func (r *DescribeAssetDetailDataExportResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDetailDataExportResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportTaskId")
	delete(f, "DspaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetDetailDataExportResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDetailDataExportResultResponseParams struct {
	// 导出结果
	ExportResult *string `json:"ExportResult,omitempty" name:"ExportResult"`

	// 导出文件地址
	ExportFileUrl *string `json:"ExportFileUrl,omitempty" name:"ExportFileUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetDetailDataExportResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetDetailDataExportResultResponseParams `json:"Response"`
}

func (r *DescribeAssetDetailDataExportResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDetailDataExportResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetOverviewRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

type DescribeAssetOverviewRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

func (r *DescribeAssetOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "AssetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetOverviewResponseParams struct {
	// 数据库实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceNums *int64 `json:"DBInstanceNums,omitempty" name:"DBInstanceNums"`

	// 数据库个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBNums *int64 `json:"DBNums,omitempty" name:"DBNums"`

	// 表的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNums *int64 `json:"TableNums,omitempty" name:"TableNums"`

	// 字段个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldNums *int64 `json:"FieldNums,omitempty" name:"FieldNums"`

	// 数据库实例的分布情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBInstanceDistribution []*Note `json:"DBInstanceDistribution,omitempty" name:"DBInstanceDistribution"`

	// db分布情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBDistribution []*Note `json:"DBDistribution,omitempty" name:"DBDistribution"`

	// cos桶的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketNums *int64 `json:"BucketNums,omitempty" name:"BucketNums"`

	// 文件个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNums *int64 `json:"FileNums,omitempty" name:"FileNums"`

	// 用于对用户进行提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// es实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsInstanceNums *uint64 `json:"EsInstanceNums,omitempty" name:"EsInstanceNums"`

	// es索引数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsIndexNums *uint64 `json:"EsIndexNums,omitempty" name:"EsIndexNums"`

	// es字段数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsFieldNums *uint64 `json:"EsFieldNums,omitempty" name:"EsFieldNums"`

	// mongo实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoInstanceNums *uint64 `json:"MongoInstanceNums,omitempty" name:"MongoInstanceNums"`

	// mongo数据库数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoDbNums *uint64 `json:"MongoDbNums,omitempty" name:"MongoDbNums"`

	// mongo集合数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoColNums *uint64 `json:"MongoColNums,omitempty" name:"MongoColNums"`

	// mongo字段数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoFieldNums *uint64 `json:"MongoFieldNums,omitempty" name:"MongoFieldNums"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetOverviewResponseParams `json:"Response"`
}

func (r *DescribeAssetOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindDBListRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`
}

type DescribeBindDBListRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`
}

func (r *DescribeBindDBListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindDBListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "DataSourceType")
	delete(f, "DataSourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindDBListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindDBListResponseParams struct {
	// 绑定的DB列表（已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindDBList []*string `json:"BindDBList,omitempty" name:"BindDBList"`

	// 绑定信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindList []*DBInstanceInfo `json:"BindList,omitempty" name:"BindList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBindDBListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindDBListResponseParams `json:"Response"`
}

func (r *DescribeBindDBListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindDBListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCOSAssetSensitiveDistributionRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

type DescribeCOSAssetSensitiveDistributionRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

func (r *DescribeCOSAssetSensitiveDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCOSAssetSensitiveDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "AssetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCOSAssetSensitiveDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCOSAssetSensitiveDistributionResponseParams struct {
	// cos的涉敏资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosAsset *CosAsset `json:"CosAsset,omitempty" name:"CosAsset"`

	// 涉敏top
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopAsset []*TopAsset `json:"TopAsset,omitempty" name:"TopAsset"`

	// cos资产详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosDetail []*AssetCosDetail `json:"CosDetail,omitempty" name:"CosDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCOSAssetSensitiveDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCOSAssetSensitiveDistributionResponseParams `json:"Response"`
}

func (r *DescribeCOSAssetSensitiveDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCOSAssetSensitiveDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassificationInfoRequestParams struct {

}

type DescribeClassificationInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeClassificationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassificationInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassificationInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassificationInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClassificationInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassificationInfoResponseParams `json:"Response"`
}

func (r *DescribeClassificationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassificationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassificationRuleCountRequestParams struct {

}

type DescribeClassificationRuleCountRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeClassificationRuleCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassificationRuleCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassificationRuleCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassificationRuleCountResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClassificationRuleCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassificationRuleCountResponseParams `json:"Response"`
}

func (r *DescribeClassificationRuleCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassificationRuleCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentHighRiskTop10OverviewRequestParams struct {
	// dspa实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeDSPAAssessmentHighRiskTop10OverviewRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeDSPAAssessmentHighRiskTop10OverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentHighRiskTop10OverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentHighRiskTop10OverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentHighRiskTop10OverviewResponseParams struct {
	// 含高风险资产TOP10的列表数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetsList []*HighRiskAssetsDetail `json:"AssetsList,omitempty" name:"AssetsList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentHighRiskTop10OverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentHighRiskTop10OverviewResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentHighRiskTop10OverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentHighRiskTop10OverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentLatestRiskDetailInfoRequestParams struct {
	// dspa实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 风险id
	RiskId *int64 `json:"RiskId,omitempty" name:"RiskId"`
}

type DescribeDSPAAssessmentLatestRiskDetailInfoRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 风险id
	RiskId *int64 `json:"RiskId,omitempty" name:"RiskId"`
}

func (r *DescribeDSPAAssessmentLatestRiskDetailInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentLatestRiskDetailInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	delete(f, "RiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentLatestRiskDetailInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentLatestRiskDetailInfoResponseParams struct {
	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 数据源name
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceName *string `json:"DataSourceName,omitempty" name:"DataSourceName"`

	// 资产对象名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 风险评估模版id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssessmentTemplateId *int64 `json:"AssessmentTemplateId,omitempty" name:"AssessmentTemplateId"`

	// 分类分级的模版id
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentifyTemplateId *int64 `json:"IdentifyTemplateId,omitempty" name:"IdentifyTemplateId"`

	// 风险类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`

	// 风险项
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskName *string `json:"RiskName,omitempty" name:"RiskName"`

	// 风险的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskDescription *string `json:"RiskDescription,omitempty" name:"RiskDescription"`

	// 风险的级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 处置的建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuggestAction *string `json:"SuggestAction,omitempty" name:"SuggestAction"`

	// 处置状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 安全产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityProduct []*SecurityProduct `json:"SecurityProduct,omitempty" name:"SecurityProduct"`

	// 风险归属
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskDimension *string `json:"RiskDimension,omitempty" name:"RiskDimension"`

	// 关联数据库（如果风险归属是instance）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelationAsset []*string `json:"RelationAsset,omitempty" name:"RelationAsset"`

	// 风险账号详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountRiskDetail []*AccountRisk `json:"AccountRiskDetail,omitempty" name:"AccountRiskDetail"`

	// 权限风险详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivilegeRiskDetail []*PrivilegeRisk `json:"PrivilegeRiskDetail,omitempty" name:"PrivilegeRiskDetail"`

	// 策略风险的cos风险文件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyRiskCosFileList []*string `json:"PolicyRiskCosFileList,omitempty" name:"PolicyRiskCosFileList"`

	// AKSK泄漏列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AKSKLeak []*AKSKLeak `json:"AKSKLeak,omitempty" name:"AKSKLeak"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentLatestRiskDetailInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentLatestRiskDetailInfoResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentLatestRiskDetailInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentLatestRiskDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentLatestRiskListRequestParams struct {
	// dspa实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 资产名称，数据源id
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 风险类型
	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`

	// 风险项
	ControlItemId *string `json:"ControlItemId,omitempty" name:"ControlItemId"`

	// 风险状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 扫描开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 扫描结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 风险等级筛选
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
}

type DescribeDSPAAssessmentLatestRiskListRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 资产名称，数据源id
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 风险类型
	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`

	// 风险项
	ControlItemId *string `json:"ControlItemId,omitempty" name:"ControlItemId"`

	// 风险状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 扫描开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 扫描结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 风险等级筛选
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
}

func (r *DescribeDSPAAssessmentLatestRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentLatestRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DataSourceId")
	delete(f, "RiskType")
	delete(f, "ControlItemId")
	delete(f, "Status")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "RiskLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentLatestRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentLatestRiskListResponseParams struct {
	// 最新风险详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestRiskList []*RiskItemInfo `json:"LatestRiskList,omitempty" name:"LatestRiskList"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentLatestRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentLatestRiskListResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentLatestRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentLatestRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentNewDiscoveredRiskOverviewRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeDSPAAssessmentNewDiscoveredRiskOverviewRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeDSPAAssessmentNewDiscoveredRiskOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentNewDiscoveredRiskOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentNewDiscoveredRiskOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentNewDiscoveredRiskOverviewResponseParams struct {
	// 待处理的风险数
	NewDiscoveredRiskCount *int64 `json:"NewDiscoveredRiskCount,omitempty" name:"NewDiscoveredRiskCount"`

	// 受影响的资产数
	AffectedAssetCount *int64 `json:"AffectedAssetCount,omitempty" name:"AffectedAssetCount"`

	// 周同比
	WeekRatio *float64 `json:"WeekRatio,omitempty" name:"WeekRatio"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentNewDiscoveredRiskOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentNewDiscoveredRiskOverviewResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentNewDiscoveredRiskOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentNewDiscoveredRiskOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentPendingRiskOverviewRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeDSPAAssessmentPendingRiskOverviewRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeDSPAAssessmentPendingRiskOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentPendingRiskOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentPendingRiskOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentPendingRiskOverviewResponseParams struct {
	// 待处理的风险数
	PendingRiskCount *int64 `json:"PendingRiskCount,omitempty" name:"PendingRiskCount"`

	// 受影响的资产数
	AffectedAssetCount *int64 `json:"AffectedAssetCount,omitempty" name:"AffectedAssetCount"`

	// 周同比
	WeekRatio *float64 `json:"WeekRatio,omitempty" name:"WeekRatio"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentPendingRiskOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentPendingRiskOverviewResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentPendingRiskOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentPendingRiskOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentProcessingRiskOverviewRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeDSPAAssessmentProcessingRiskOverviewRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeDSPAAssessmentProcessingRiskOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentProcessingRiskOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentProcessingRiskOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentProcessingRiskOverviewResponseParams struct {
	// 待处理的风险数
	ProcessingRiskCount *int64 `json:"ProcessingRiskCount,omitempty" name:"ProcessingRiskCount"`

	// 受影响的资产数
	AffectedAssetCount *int64 `json:"AffectedAssetCount,omitempty" name:"AffectedAssetCount"`

	// 周同比
	WeekRatio *float64 `json:"WeekRatio,omitempty" name:"WeekRatio"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentProcessingRiskOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentProcessingRiskOverviewResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentProcessingRiskOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentProcessingRiskOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskAmountOverviewRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeDSPAAssessmentRiskAmountOverviewRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeDSPAAssessmentRiskAmountOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskAmountOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskAmountOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskAmountOverviewResponseParams struct {
	// 风险总数
	TotalRiskCount *int64 `json:"TotalRiskCount,omitempty" name:"TotalRiskCount"`

	// 受影响的资产数
	TotalAffectedAssetCount *int64 `json:"TotalAffectedAssetCount,omitempty" name:"TotalAffectedAssetCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskAmountOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskAmountOverviewResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskAmountOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskAmountOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskDatasourceTop5RequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

type DescribeDSPAAssessmentRiskDatasourceTop5Request struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

func (r *DescribeDSPAAssessmentRiskDatasourceTop5Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskDatasourceTop5Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskDatasourceTop5Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskDatasourceTop5ResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*RiskItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskDatasourceTop5Response struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskDatasourceTop5ResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskDatasourceTop5Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskDatasourceTop5Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskDealedOverviewRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

type DescribeDSPAAssessmentRiskDealedOverviewRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

func (r *DescribeDSPAAssessmentRiskDealedOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskDealedOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskDealedOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskDealedOverviewResponseParams struct {
	// 遗留待处理风险总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 昨日完成风险处置数
	YesterdayDealedCount *uint64 `json:"YesterdayDealedCount,omitempty" name:"YesterdayDealedCount"`

	// 遗留待处理风险数周同比
	UnDealedRiskWeekRatio *float64 `json:"UnDealedRiskWeekRatio,omitempty" name:"UnDealedRiskWeekRatio"`

	// 遗留待处理风险数日环比
	UnDealedRiskDayRatio *float64 `json:"UnDealedRiskDayRatio,omitempty" name:"UnDealedRiskDayRatio"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskDealedOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskDealedOverviewResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskDealedOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskDealedOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskDealedTrendRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 开始日期
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束日期
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 评估模版id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeDSPAAssessmentRiskDealedTrendRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 开始日期
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束日期
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 评估模版id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeDSPAAssessmentRiskDealedTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskDealedTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskDealedTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskDealedTrendResponseParams struct {
	// 趋势统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*RiskDealedTrendItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskDealedTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskDealedTrendResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskDealedTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskDealedTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskDistributionOverviewRequestParams struct {
	// dspa实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeDSPAAssessmentRiskDistributionOverviewRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeDSPAAssessmentRiskDistributionOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskDistributionOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskDistributionOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskDistributionOverviewResponseParams struct {
	// 风险类型分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskTypeDistribution []*Note `json:"RiskTypeDistribution,omitempty" name:"RiskTypeDistribution"`

	// 风险详情分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskDetailDistribution []*Note `json:"RiskDetailDistribution,omitempty" name:"RiskDetailDistribution"`

	// 风险资产详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskAssetsDistribution []*Note `json:"RiskAssetsDistribution,omitempty" name:"RiskAssetsDistribution"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskDistributionOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskDistributionOverviewResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskDistributionOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskDistributionOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskItemTop5RequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

type DescribeDSPAAssessmentRiskItemTop5Request struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

func (r *DescribeDSPAAssessmentRiskItemTop5Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskItemTop5Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskItemTop5Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskItemTop5ResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*RiskItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskItemTop5Response struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskItemTop5ResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskItemTop5Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskItemTop5Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskLevelDetailRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 风险级别id
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" name:"RiskLevelId"`
}

type DescribeDSPAAssessmentRiskLevelDetailRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 风险级别id
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" name:"RiskLevelId"`
}

func (r *DescribeDSPAAssessmentRiskLevelDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskLevelDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "RiskLevelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskLevelDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskLevelDetailResponseParams struct {
	// 风险级别名称
	RiskLevelName *string `json:"RiskLevelName,omitempty" name:"RiskLevelName"`

	// 风险级别描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevelDescription *string `json:"RiskLevelDescription,omitempty" name:"RiskLevelDescription"`

	// 分类分级id
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentifyComplianceId *int64 `json:"IdentifyComplianceId,omitempty" name:"IdentifyComplianceId"`

	// 分类分级模版名称
	IdentifyComplianceName *string `json:"IdentifyComplianceName,omitempty" name:"IdentifyComplianceName"`

	// 风险数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevelMatrix []*RiskLevelMatrix `json:"RiskLevelMatrix,omitempty" name:"RiskLevelMatrix"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskLevelDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskLevelDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskLevelDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskLevelDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskLevelListRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDSPAAssessmentRiskLevelListRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDSPAAssessmentRiskLevelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskLevelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskLevelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskLevelListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 风险等级列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevelList []*RiskLevelRisk `json:"RiskLevelList,omitempty" name:"RiskLevelList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskLevelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskLevelListResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskLevelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskLevelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskLevelTrendRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 开始日期
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时日期
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 评估模版id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeDSPAAssessmentRiskLevelTrendRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 开始日期
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时日期
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 评估模版id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeDSPAAssessmentRiskLevelTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskLevelTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskLevelTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskLevelTrendResponseParams struct {
	// 结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*RiskLevelTrendItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskLevelTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskLevelTrendResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskLevelTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskLevelTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskOverviewRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

type DescribeDSPAAssessmentRiskOverviewRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

func (r *DescribeDSPAAssessmentRiskOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskOverviewResponseParams struct {
	// 风险总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 高危风险数
	HighRiskCount *uint64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`

	// 周同比
	HighRiskWeekRatio *float64 `json:"HighRiskWeekRatio,omitempty" name:"HighRiskWeekRatio"`

	// 高危风险数日环比
	HighRiskDayRatio *float64 `json:"HighRiskDayRatio,omitempty" name:"HighRiskDayRatio"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskOverviewResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskProcessHistoryRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// xxx
	RiskId *int64 `json:"RiskId,omitempty" name:"RiskId"`
}

type DescribeDSPAAssessmentRiskProcessHistoryRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// xxx
	RiskId *int64 `json:"RiskId,omitempty" name:"RiskId"`
}

func (r *DescribeDSPAAssessmentRiskProcessHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskProcessHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "RiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskProcessHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskProcessHistoryResponseParams struct {
	// 处理的历史
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessHistory []*ProcessHistory `json:"ProcessHistory,omitempty" name:"ProcessHistory"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskProcessHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskProcessHistoryResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskProcessHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskProcessHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskTemplateDetailRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDSPAAssessmentRiskTemplateDetailRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDSPAAssessmentRiskTemplateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskTemplateDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskTemplateDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskTemplateDetailResponseParams struct {
	// 模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模版名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模版的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 风险等级
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" name:"RiskLevelId"`

	// 风险等级名称
	RiskLevelName *string `json:"RiskLevelName,omitempty" name:"RiskLevelName"`

	// 脆弱项配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskItemList []*AssessmentRiskItem `json:"RiskItemList,omitempty" name:"RiskItemList"`

	// 脆弱项配置条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 被任务引用次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCitations *int64 `json:"TaskCitations,omitempty" name:"TaskCitations"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskTemplateDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskTemplateDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskTemplateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskTemplateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskTemplateVulnerableListRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 风险类型
	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`

	// 风险名称
	RiskName *string `json:"RiskName,omitempty" name:"RiskName"`
}

type DescribeDSPAAssessmentRiskTemplateVulnerableListRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 风险类型
	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`

	// 风险名称
	RiskName *string `json:"RiskName,omitempty" name:"RiskName"`
}

func (r *DescribeDSPAAssessmentRiskTemplateVulnerableListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskTemplateVulnerableListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "RiskType")
	delete(f, "RiskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRiskTemplateVulnerableListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRiskTemplateVulnerableListResponseParams struct {
	// 脆弱项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskItemList []*AssessmentRiskItem `json:"RiskItemList,omitempty" name:"RiskItemList"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRiskTemplateVulnerableListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRiskTemplateVulnerableListResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRiskTemplateVulnerableListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRiskTemplateVulnerableListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRisksRequestParams struct {
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估任务Id，格式“task-xxxxxxxx”
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 偏移量。默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果集个数限制。默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤项。
	// 支持模糊搜索：ControlItemName。
	// 支持过滤：
	// RiskLevel：风险等级（high，medium，low）
	// Status：风险处理状态(waiting待处理, processing处理中, stopped处理暂停, finished已处理, failed处理失败)
	Filters []*DspaAssessmentFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDSPAAssessmentRisksRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估任务Id，格式“task-xxxxxxxx”
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 偏移量。默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果集个数限制。默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤项。
	// 支持模糊搜索：ControlItemName。
	// 支持过滤：
	// RiskLevel：风险等级（high，medium，low）
	// Status：风险处理状态(waiting待处理, processing处理中, stopped处理暂停, finished已处理, failed处理失败)
	Filters []*DspaAssessmentFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDSPAAssessmentRisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentRisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentRisksResponseParams struct {
	// 符合条件的评估风险项数目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 评估风险项列表
	Items []*AssessmentRisk `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentRisksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentRisksResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentRisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentRisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentTasksRequestParams struct {
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 偏移量。默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果集个数限制。默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤项。
	// 支持模糊搜索：TaskId，TaskName
	// 支持过滤：
	// BusinessName：业务名称
	// BusinessDept：业务部门名称
	// TemplateName：评估模版名称
	// Status：评估状态 (waiting待评估，processing评估中, , finished已评估, failed评估失败)
	Filters []*DspaAssessmentFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDSPAAssessmentTasksRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 偏移量。默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果集个数限制。默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤项。
	// 支持模糊搜索：TaskId，TaskName
	// 支持过滤：
	// BusinessName：业务名称
	// BusinessDept：业务部门名称
	// TemplateName：评估模版名称
	// Status：评估状态 (waiting待评估，processing评估中, , finished已评估, failed评估失败)
	Filters []*DspaAssessmentFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDSPAAssessmentTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentTasksResponseParams struct {
	// 符合条件的评估任务数目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 评估任务列表
	Items []*AssessmentTask `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentTasksResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentTemplateControlItemsRequestParams struct {
	// DSPA实例Id。格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版Id。格式“template-xxxxxxxx”
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 偏移量。默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果集个数限制。默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤项。
	// 支持模糊搜索：ItemId，ItemName
	// 支持过滤：
	// Source：评估项来源，system / user
	// ItemType：评估项类型，questionnaire / auto
	// ItemSubType：评估项子类型
	// Status：评估项启用状态，draft / launched
	Filters []*DspaAssessmentFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDSPAAssessmentTemplateControlItemsRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id。格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估模版Id。格式“template-xxxxxxxx”
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 偏移量。默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果集个数限制。默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤项。
	// 支持模糊搜索：ItemId，ItemName
	// 支持过滤：
	// Source：评估项来源，system / user
	// ItemType：评估项类型，questionnaire / auto
	// ItemSubType：评估项子类型
	// Status：评估项启用状态，draft / launched
	Filters []*DspaAssessmentFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDSPAAssessmentTemplateControlItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentTemplateControlItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentTemplateControlItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentTemplateControlItemsResponseParams struct {
	// 符合条件的评估项数目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 模版关联的评估项列表
	Items []*AssessmentControlItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentTemplateControlItemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentTemplateControlItemsResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentTemplateControlItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentTemplateControlItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentTemplatesRequestParams struct {
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 偏移量。默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果集个数限制。默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤项。
	// 支持模糊搜索：（TemplateId，TemplateName）
	// 支持过滤：
	// Source：模版来源，system / user
	// UseType：模版类型，auto，semi-auto，law等
	// Status：模版启用状态，draft / launched
	Filters []*DspaAssessmentFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDSPAAssessmentTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 偏移量。默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果集个数限制。默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤项。
	// 支持模糊搜索：（TemplateId，TemplateName）
	// 支持过滤：
	// Source：模版来源，system / user
	// UseType：模版类型，auto，semi-auto，law等
	// Status：模版启用状态，draft / launched
	Filters []*DspaAssessmentFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDSPAAssessmentTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAAssessmentTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAAssessmentTemplatesResponseParams struct {
	// 符合条件的模板数目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 模板列表。
	Items []*AssessmentTemplate `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAAssessmentTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAAssessmentTemplatesResponseParams `json:"Response"`
}

func (r *DescribeDSPAAssessmentTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAAssessmentTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDataAssetBucketsRequestParams struct {
	// DSPA实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组Id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type DescribeDSPACOSDataAssetBucketsRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组Id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *DescribeDSPACOSDataAssetBucketsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDataAssetBucketsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACOSDataAssetBucketsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDataAssetBucketsResponseParams struct {
	// COS对象存储敏感数据资产已扫描的桶集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Buckets []*string `json:"Buckets,omitempty" name:"Buckets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACOSDataAssetBucketsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACOSDataAssetBucketsResponseParams `json:"Response"`
}

func (r *DescribeDSPACOSDataAssetBucketsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDataAssetBucketsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDataAssetByComplianceIdRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type DescribeDSPACOSDataAssetByComplianceIdRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *DescribeDSPACOSDataAssetByComplianceIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDataAssetByComplianceIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACOSDataAssetByComplianceIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDataAssetByComplianceIdResponseParams struct {
	// 符合条件的COS存储对象的敏感数据资产统计记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stats *DspaCOSDataAssetCount `json:"Stats,omitempty" name:"Stats"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACOSDataAssetByComplianceIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACOSDataAssetByComplianceIdResponseParams `json:"Response"`
}

func (r *DescribeDSPACOSDataAssetByComplianceIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDataAssetByComplianceIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDataAssetDetailRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type DescribeDSPACOSDataAssetDetailRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *DescribeDSPACOSDataAssetDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDataAssetDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACOSDataAssetDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDataAssetDetailResponseParams struct {
	// COS对象存储敏感数据资产详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*DspaCOSDataAssetDetail `json:"Details,omitempty" name:"Details"`

	// 符合条件的COS对象存储敏感数据资产数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACOSDataAssetDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACOSDataAssetDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPACOSDataAssetDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDataAssetDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDiscoveryTaskDetailRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeDSPACOSDiscoveryTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDSPACOSDiscoveryTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDiscoveryTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACOSDiscoveryTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDiscoveryTaskDetailResponseParams struct {
	// 任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Task *DspaCOSDiscoveryTaskDetail `json:"Task,omitempty" name:"Task"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACOSDiscoveryTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACOSDiscoveryTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPACOSDiscoveryTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDiscoveryTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDiscoveryTaskFilesRequestParams struct {
	// DSPA实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 扫描任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 扫描Bucket任务结果ID
	BucketResultId *int64 `json:"BucketResultId,omitempty" name:"BucketResultId"`
}

type DescribeDSPACOSDiscoveryTaskFilesRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 扫描任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 扫描Bucket任务结果ID
	BucketResultId *int64 `json:"BucketResultId,omitempty" name:"BucketResultId"`
}

func (r *DescribeDSPACOSDiscoveryTaskFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDiscoveryTaskFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "BucketResultId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACOSDiscoveryTaskFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDiscoveryTaskFilesResponseParams struct {
	// 文件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Files []*string `json:"Files,omitempty" name:"Files"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACOSDiscoveryTaskFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACOSDiscoveryTaskFilesResponseParams `json:"Response"`
}

func (r *DescribeDSPACOSDiscoveryTaskFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDiscoveryTaskFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDiscoveryTaskResultRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Array of Filter	此参数对外不可见。过滤数组。支持的Name：
	// BucketName 对象桶名
	// TaskID 任务ID，
	// TaskName 任务名，
	// DataSourceId：数据源ID，
	// ResourceRegion：资源所在地域
	// 每项过滤条件最多支持5个。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDSPACOSDiscoveryTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Array of Filter	此参数对外不可见。过滤数组。支持的Name：
	// BucketName 对象桶名
	// TaskID 任务ID，
	// TaskName 任务名，
	// DataSourceId：数据源ID，
	// ResourceRegion：资源所在地域
	// 每项过滤条件最多支持5个。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDSPACOSDiscoveryTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDiscoveryTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACOSDiscoveryTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDiscoveryTaskResultResponseParams struct {
	// 扫描任务结果项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DspaCOSDiscoveryTaskResult `json:"Items,omitempty" name:"Items"`

	// 符合条件的数据结果数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACOSDiscoveryTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACOSDiscoveryTaskResultResponseParams `json:"Response"`
}

func (r *DescribeDSPACOSDiscoveryTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDiscoveryTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDiscoveryTasksRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务扫描结果状态，可供选择的状态值有：-1待触发 0待扫描 1扫描中 2扫描终止 3扫描成功 4扫描失败
	StatusList []*int64 `json:"StatusList,omitempty" name:"StatusList"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果集数量，默认值是10000，最大值为10000，根据该资源的个数限制条件，该资源的个数不会超过10000，所以如果不输入该字段，默认获取全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDSPACOSDiscoveryTasksRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务扫描结果状态，可供选择的状态值有：-1待触发 0待扫描 1扫描中 2扫描终止 3扫描成功 4扫描失败
	StatusList []*int64 `json:"StatusList,omitempty" name:"StatusList"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果集数量，默认值是10000，最大值为10000，根据该资源的个数限制条件，该资源的个数不会超过10000，所以如果不输入该字段，默认获取全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDSPACOSDiscoveryTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDiscoveryTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "StatusList")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACOSDiscoveryTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSDiscoveryTasksResponseParams struct {
	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DspaCOSDiscoveryTask `json:"Items,omitempty" name:"Items"`

	// 符合条件的任务列表数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACOSDiscoveryTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACOSDiscoveryTasksResponseParams `json:"Response"`
}

func (r *DescribeDSPACOSDiscoveryTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSDiscoveryTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSTaskResultDetailRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 扫描Bucket结果ID
	BucketResultId *int64 `json:"BucketResultId,omitempty" name:"BucketResultId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 敏感数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分级ID
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 扫描桶名称。
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 多级分类的分类ID集合
	CategoryIdList []*int64 `json:"CategoryIdList,omitempty" name:"CategoryIdList"`
}

type DescribeDSPACOSTaskResultDetailRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 扫描Bucket结果ID
	BucketResultId *int64 `json:"BucketResultId,omitempty" name:"BucketResultId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 敏感数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分级ID
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 扫描桶名称。
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 多级分类的分类ID集合
	CategoryIdList []*int64 `json:"CategoryIdList,omitempty" name:"CategoryIdList"`
}

func (r *DescribeDSPACOSTaskResultDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSTaskResultDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "BucketResultId")
	delete(f, "ComplianceId")
	delete(f, "FileName")
	delete(f, "CategoryId")
	delete(f, "LevelId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "BucketName")
	delete(f, "CategoryIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACOSTaskResultDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACOSTaskResultDetailResponseParams struct {
	// 扫描结果详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DspaDiscoveryCOSTaskResultDetail `json:"Items,omitempty" name:"Items"`

	// 符合条件的详情数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACOSTaskResultDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACOSTaskResultDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPACOSTaskResultDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACOSTaskResultDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACategoriesRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分类名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果集数量，默认值是10000，最大值为10000，根据该资源的个数限制条件，该资源的个数不会超过10000，所以如果不输入该字段，默认获取全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDSPACategoriesRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分类名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果集数量，默认值是10000，最大值为10000，根据该资源的个数限制条件，该资源的个数不会超过10000，所以如果不输入该字段，默认获取全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDSPACategoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACategoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "CategoryId")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACategoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACategoriesResponseParams struct {
	// 敏感数据分类列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DataCategory `json:"Items,omitempty" name:"Items"`

	// 符合条件的敏感数据分类数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACategoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACategoriesResponseParams `json:"Response"`
}

func (r *DescribeDSPACategoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACategoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACategoryRuleStatisticRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type DescribeDSPACategoryRuleStatisticRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *DescribeDSPACategoryRuleStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACategoryRuleStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACategoryRuleStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACategoryRuleStatisticResponseParams struct {
	// 分类规则统计信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticSet []*CategoryRuleStatistic `json:"StatisticSet,omitempty" name:"StatisticSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACategoryRuleStatisticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACategoryRuleStatisticResponseParams `json:"Response"`
}

func (r *DescribeDSPACategoryRuleStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACategoryRuleStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACategoryRulesRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分类id
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 合规组模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type DescribeDSPACategoryRulesRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分类id
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 合规组模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *DescribeDSPACategoryRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACategoryRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "CategoryId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACategoryRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACategoryRulesResponseParams struct {
	// 分类规则信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryRules []*CategoryRule `json:"CategoryRules,omitempty" name:"CategoryRules"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACategoryRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACategoryRulesResponseParams `json:"Response"`
}

func (r *DescribeDSPACategoryRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACategoryRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACategoryTreeRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type DescribeDSPACategoryTreeRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *DescribeDSPACategoryTreeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACategoryTreeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACategoryTreeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACategoryTreeResponseParams struct {
	// 分类树json
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultJson *string `json:"ResultJson,omitempty" name:"ResultJson"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACategoryTreeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACategoryTreeResponseParams `json:"Response"`
}

func (r *DescribeDSPACategoryTreeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACategoryTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACategoryTreeWithRulesRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 分类id
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
}

type DescribeDSPACategoryTreeWithRulesRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 分类id
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`
}

func (r *DescribeDSPACategoryTreeWithRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACategoryTreeWithRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "CategoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPACategoryTreeWithRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPACategoryTreeWithRulesResponseParams struct {
	// 分类树json
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultJson *string `json:"ResultJson,omitempty" name:"ResultJson"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPACategoryTreeWithRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPACategoryTreeWithRulesResponseParams `json:"Response"`
}

func (r *DescribeDSPACategoryTreeWithRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPACategoryTreeWithRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAComplianceGroupDetailRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type DescribeDSPAComplianceGroupDetailRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *DescribeDSPAComplianceGroupDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAComplianceGroupDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAComplianceGroupDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAComplianceGroupDetailResponseParams struct {
	// 模板详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *ComplianceGroupDetail `json:"Detail,omitempty" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAComplianceGroupDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAComplianceGroupDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPAComplianceGroupDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAComplianceGroupDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAComplianceGroupsRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceGroupId *int64 `json:"ComplianceGroupId,omitempty" name:"ComplianceGroupId"`

	// 合规组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果集数量，默认值是10000，最大值为10000，根据该资源的个数限制条件，该资源的个数不会超过10000，所以如果不输入该字段，默认获取全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 合规组类型可选值：0 默认合规组, 1 系统合规组, 2 自定义合规组
	ComplianceGroupTypeList []*int64 `json:"ComplianceGroupTypeList,omitempty" name:"ComplianceGroupTypeList"`
}

type DescribeDSPAComplianceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceGroupId *int64 `json:"ComplianceGroupId,omitempty" name:"ComplianceGroupId"`

	// 合规组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果集数量，默认值是10000，最大值为10000，根据该资源的个数限制条件，该资源的个数不会超过10000，所以如果不输入该字段，默认获取全量数据
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 合规组类型可选值：0 默认合规组, 1 系统合规组, 2 自定义合规组
	ComplianceGroupTypeList []*int64 `json:"ComplianceGroupTypeList,omitempty" name:"ComplianceGroupTypeList"`
}

func (r *DescribeDSPAComplianceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAComplianceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceGroupId")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ComplianceGroupTypeList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAComplianceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAComplianceGroupsResponseParams struct {
	// 合规组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DspaDiscoveryComplianceGroupInfo `json:"Items,omitempty" name:"Items"`

	// 符合条件的合规组列表数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAComplianceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAComplianceGroupsResponseParams `json:"Response"`
}

func (r *DescribeDSPAComplianceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAComplianceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAComplianceUpdateNotificationRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组分类模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type DescribeDSPAComplianceUpdateNotificationRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组分类模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *DescribeDSPAComplianceUpdateNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAComplianceUpdateNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAComplianceUpdateNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAComplianceUpdateNotificationResponseParams struct {
	// 模板是否更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdated *bool `json:"IsUpdated,omitempty" name:"IsUpdated"`

	// 任务名称集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNameSet []*string `json:"TaskNameSet,omitempty" name:"TaskNameSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAComplianceUpdateNotificationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAComplianceUpdateNotificationResponseParams `json:"Response"`
}

func (r *DescribeDSPAComplianceUpdateNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAComplianceUpdateNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADataSourceDbInfoRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`
}

type DescribeDSPADataSourceDbInfoRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`
}

func (r *DescribeDSPADataSourceDbInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADataSourceDbInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "DataSourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPADataSourceDbInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADataSourceDbInfoResponseParams struct {
	// 数据库信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DSPADataSourceDbInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPADataSourceDbInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPADataSourceDbInfoResponseParams `json:"Response"`
}

func (r *DescribeDSPADataSourceDbInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADataSourceDbInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryRulesRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回上限，默认值10， 最大值10000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要过滤别名
	FilterRuleSource *bool `json:"FilterRuleSource,omitempty" name:"FilterRuleSource"`
}

type DescribeDSPADiscoveryRulesRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回上限，默认值10， 最大值10000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否需要过滤别名
	FilterRuleSource *bool `json:"FilterRuleSource,omitempty" name:"FilterRuleSource"`
}

func (r *DescribeDSPADiscoveryRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RuleId")
	delete(f, "Name")
	delete(f, "FilterRuleSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPADiscoveryRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryRulesResponseParams struct {
	// 规则ID
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 规则集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DspaDiscoveryRuleDetail `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPADiscoveryRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPADiscoveryRulesResponseParams `json:"Response"`
}

func (r *DescribeDSPADiscoveryRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryServiceStatusRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

type DescribeDSPADiscoveryServiceStatusRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

func (r *DescribeDSPADiscoveryServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryServiceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPADiscoveryServiceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryServiceStatusResponseParams struct {
	// 分类分级服务是否开通，true 表示已开通，false表示未开通
	ServiceEnabled *bool `json:"ServiceEnabled,omitempty" name:"ServiceEnabled"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPADiscoveryServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPADiscoveryServiceStatusResponseParams `json:"Response"`
}

func (r *DescribeDSPADiscoveryServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryTaskDetailRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeDSPADiscoveryTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDSPADiscoveryTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPADiscoveryTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryTaskDetailResponseParams struct {
	// 任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Task *DspaDiscoveryTaskDetail `json:"Task,omitempty" name:"Task"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPADiscoveryTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPADiscoveryTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPADiscoveryTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryTaskResultDetailRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 扫描数据库结果ID
	DbResultId *int64 `json:"DbResultId,omitempty" name:"DbResultId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 数据库名
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 所属数据表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 敏感数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分级ID
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 多级分类的分类ID集合
	CategoryIdList []*int64 `json:"CategoryIdList,omitempty" name:"CategoryIdList"`
}

type DescribeDSPADiscoveryTaskResultDetailRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 扫描数据库结果ID
	DbResultId *int64 `json:"DbResultId,omitempty" name:"DbResultId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 数据库名
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 所属数据表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 敏感数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分级ID
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 多级分类的分类ID集合
	CategoryIdList []*int64 `json:"CategoryIdList,omitempty" name:"CategoryIdList"`
}

func (r *DescribeDSPADiscoveryTaskResultDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryTaskResultDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "DbResultId")
	delete(f, "ComplianceId")
	delete(f, "DbName")
	delete(f, "TableName")
	delete(f, "CategoryId")
	delete(f, "LevelId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CategoryIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPADiscoveryTaskResultDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryTaskResultDetailResponseParams struct {
	// 扫描结果详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DspaDiscoveryTaskResultDetail `json:"Items,omitempty" name:"Items"`

	// 符合条件的扫描结果详情记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPADiscoveryTaskResultDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPADiscoveryTaskResultDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPADiscoveryTaskResultDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryTaskResultDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryTaskResultRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据源类型，可取值如下：
	// cdb 表示云数据库 MySQL,
	// dcdb 表示TDSQL MySQL版,
	// mariadb 表示云数据库 MariaDB,
	// postgres 表示云数据库 PostgreSQL,
	// cynosdbpg 表示TDSQL-C PostgreSQL版,
	// cynosdbmysql 表示TDSQL-C MySQL版,
	// selfbuilt-db 表示自建数据库
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
}

type DescribeDSPADiscoveryTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据源类型，可取值如下：
	// cdb 表示云数据库 MySQL,
	// dcdb 表示TDSQL MySQL版,
	// mariadb 表示云数据库 MariaDB,
	// postgres 表示云数据库 PostgreSQL,
	// cynosdbpg 表示TDSQL-C PostgreSQL版,
	// cynosdbmysql 表示TDSQL-C MySQL版,
	// selfbuilt-db 表示自建数据库
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
}

func (r *DescribeDSPADiscoveryTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "DataSourceType")
	delete(f, "TaskId")
	delete(f, "TaskName")
	delete(f, "DataSourceId")
	delete(f, "DbName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPADiscoveryTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryTaskResultResponseParams struct {
	// 扫描任务结果项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DspaDiscoveryTaskDbResult `json:"Items,omitempty" name:"Items"`

	// 符合条件的扫描任务结果记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPADiscoveryTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPADiscoveryTaskResultResponseParams `json:"Response"`
}

func (r *DescribeDSPADiscoveryTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryTaskTablesRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 数据库扫描结果ID
	DbResultId *int64 `json:"DbResultId,omitempty" name:"DbResultId"`

	// db名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

type DescribeDSPADiscoveryTaskTablesRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 数据库扫描结果ID
	DbResultId *int64 `json:"DbResultId,omitempty" name:"DbResultId"`

	// db名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *DescribeDSPADiscoveryTaskTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryTaskTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "DbResultId")
	delete(f, "DbName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPADiscoveryTaskTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPADiscoveryTaskTablesResponseParams struct {
	// 分类分级扫描表集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DSPATableInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPADiscoveryTaskTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPADiscoveryTaskTablesResponseParams `json:"Response"`
}

func (r *DescribeDSPADiscoveryTaskTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPADiscoveryTaskTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAESDataAssetByComplianceIdRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type DescribeDSPAESDataAssetByComplianceIdRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *DescribeDSPAESDataAssetByComplianceIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAESDataAssetByComplianceIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAESDataAssetByComplianceIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAESDataAssetByComplianceIdResponseParams struct {
	// 概览统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stats *ESDataAssetCountDto `json:"Stats,omitempty" name:"Stats"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAESDataAssetByComplianceIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAESDataAssetByComplianceIdResponseParams `json:"Response"`
}

func (r *DescribeDSPAESDataAssetByComplianceIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAESDataAssetByComplianceIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAESDataAssetDetailRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制条目数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDSPAESDataAssetDetailRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制条目数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDSPAESDataAssetDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAESDataAssetDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAESDataAssetDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAESDataAssetDetailResponseParams struct {
	// 总的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 概览详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*ESDataAssetDetail `json:"Details,omitempty" name:"Details"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAESDataAssetDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAESDataAssetDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPAESDataAssetDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAESDataAssetDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAESDataSampleRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 字段扫描结果ID
	FieldResultId *int64 `json:"FieldResultId,omitempty" name:"FieldResultId"`
}

type DescribeDSPAESDataSampleRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 字段扫描结果ID
	FieldResultId *int64 `json:"FieldResultId,omitempty" name:"FieldResultId"`
}

func (r *DescribeDSPAESDataSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAESDataSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "FieldResultId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAESDataSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAESDataSampleResponseParams struct {
	// 数据样本列表，最多10条数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*string `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAESDataSampleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAESDataSampleResponseParams `json:"Response"`
}

func (r *DescribeDSPAESDataSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAESDataSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAESDiscoveryTaskResultDetailRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 多级分类的分类ID集合
	CategoryIdList []*int64 `json:"CategoryIdList,omitempty" name:"CategoryIdList"`

	// 敏感数据分级ID
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`
}

type DescribeDSPAESDiscoveryTaskResultDetailRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 多级分类的分类ID集合
	CategoryIdList []*int64 `json:"CategoryIdList,omitempty" name:"CategoryIdList"`

	// 敏感数据分级ID
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`
}

func (r *DescribeDSPAESDiscoveryTaskResultDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAESDiscoveryTaskResultDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "ComplianceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CategoryIdList")
	delete(f, "LevelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPAESDiscoveryTaskResultDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPAESDiscoveryTaskResultDetailResponseParams struct {
	// ES扫描结果详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ESTaskResultDetail `json:"Items,omitempty" name:"Items"`

	// 符合条件的扫描结果详情记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPAESDiscoveryTaskResultDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPAESDiscoveryTaskResultDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPAESDiscoveryTaskResultDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPAESDiscoveryTaskResultDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPALevelDetailRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 过滤数组。支持的Name：
	// ComplianceId 合规组ID
	// LevelGroupId 敏感分级组ID
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDSPALevelDetailRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 过滤数组。支持的Name：
	// ComplianceId 合规组ID
	// LevelGroupId 敏感分级组ID
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDSPALevelDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPALevelDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPALevelDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPALevelDetailResponseParams struct {
	// 符合条件的敏感数据分级标识记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*LevelItem `json:"Items,omitempty" name:"Items"`

	// 符合条件的敏感数据分级标识记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPALevelDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPALevelDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPALevelDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPALevelDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPALevelGroupsRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分级组名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeDSPALevelGroupsRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分级组名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeDSPALevelGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPALevelGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPALevelGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPALevelGroupsResponseParams struct {
	// 符合条件的敏感数据分级标识记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DspaDiscoveryLevelDetail `json:"Items,omitempty" name:"Items"`

	// 符合条件的敏感数据分级标识记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPALevelGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPALevelGroupsResponseParams `json:"Response"`
}

func (r *DescribeDSPALevelGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPALevelGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPARDBDataAssetByComplianceIdRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 数据源类型，不填默认过滤非自建的所有关系型数据源类型，填selfbuilt-db只过滤自建类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`
}

type DescribeDSPARDBDataAssetByComplianceIdRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 数据源类型，不填默认过滤非自建的所有关系型数据源类型，填selfbuilt-db只过滤自建类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`
}

func (r *DescribeDSPARDBDataAssetByComplianceIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPARDBDataAssetByComplianceIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "DataSourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPARDBDataAssetByComplianceIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPARDBDataAssetByComplianceIdResponseParams struct {
	// 符合条件的RDB关系数据库敏感数据资产统计记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stats *DspaRDBDataAssetCount `json:"Stats,omitempty" name:"Stats"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPARDBDataAssetByComplianceIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPARDBDataAssetByComplianceIdResponseParams `json:"Response"`
}

func (r *DescribeDSPARDBDataAssetByComplianceIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPARDBDataAssetByComplianceIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPARDBDataAssetDetailRequestParams struct {
	// DSPA实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组Id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type DescribeDSPARDBDataAssetDetailRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组Id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *DescribeDSPARDBDataAssetDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPARDBDataAssetDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPARDBDataAssetDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPARDBDataAssetDetailResponseParams struct {
	// RDB关系数据库敏感数据资产详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*DspaRDBDataAssetDetail `json:"Details,omitempty" name:"Details"`

	// 符合条件的RDB关系数据库敏感数据资产数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPARDBDataAssetDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPARDBDataAssetDetailResponseParams `json:"Response"`
}

func (r *DescribeDSPARDBDataAssetDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPARDBDataAssetDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPASupportedMetasRequestParams struct {
	// 元数据类型
	MetaTypes []*string `json:"MetaTypes,omitempty" name:"MetaTypes"`
}

type DescribeDSPASupportedMetasRequest struct {
	*tchttp.BaseRequest
	
	// 元数据类型
	MetaTypes []*string `json:"MetaTypes,omitempty" name:"MetaTypes"`
}

func (r *DescribeDSPASupportedMetasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPASupportedMetasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MetaTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPASupportedMetasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPASupportedMetasResponseParams struct {
	// 支持的元数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metas []*DSPAMetaType `json:"Metas,omitempty" name:"Metas"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPASupportedMetasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPASupportedMetasResponseParams `json:"Response"`
}

func (r *DescribeDSPASupportedMetasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPASupportedMetasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPATaskResultDataSampleRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 字段扫描结果ID
	FieldResultId *int64 `json:"FieldResultId,omitempty" name:"FieldResultId"`
}

type DescribeDSPATaskResultDataSampleRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 字段扫描结果ID
	FieldResultId *int64 `json:"FieldResultId,omitempty" name:"FieldResultId"`
}

func (r *DescribeDSPATaskResultDataSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPATaskResultDataSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "FieldResultId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDSPATaskResultDataSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDSPATaskResultDataSampleResponseParams struct {
	// 数据样本列表，最多10条数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DspaFieldResultDataSample `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDSPATaskResultDataSampleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDSPATaskResultDataSampleResponseParams `json:"Response"`
}

func (r *DescribeDSPATaskResultDataSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDSPATaskResultDataSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeESAssetSensitiveDistributionRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

type DescribeESAssetSensitiveDistributionRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

func (r *DescribeESAssetSensitiveDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeESAssetSensitiveDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "AssetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeESAssetSensitiveDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeESAssetSensitiveDistributionResponseParams struct {
	// ES的资产统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ESAsset *ESAsset `json:"ESAsset,omitempty" name:"ESAsset"`

	// 涉敏top数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopAsset []*TopAsset `json:"TopAsset,omitempty" name:"TopAsset"`

	// ES的详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ESDetail []*ESAssetDBDetail `json:"ESDetail,omitempty" name:"ESDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeESAssetSensitiveDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeESAssetSensitiveDistributionResponseParams `json:"Response"`
}

func (r *DescribeESAssetSensitiveDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeESAssetSensitiveDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportTaskResultRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 导出任务id
	ExportTaskId *uint64 `json:"ExportTaskId,omitempty" name:"ExportTaskId"`
}

type DescribeExportTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 导出任务id
	ExportTaskId *uint64 `json:"ExportTaskId,omitempty" name:"ExportTaskId"`
}

func (r *DescribeExportTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ExportTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportTaskResultResponseParams struct {
	// 导出任务结果
	ExportResult *string `json:"ExportResult,omitempty" name:"ExportResult"`

	// 导出文件地址
	ExportFileUrl *string `json:"ExportFileUrl,omitempty" name:"ExportFileUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExportTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExportTaskResultResponseParams `json:"Response"`
}

func (r *DescribeExportTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLeafClassificationRequestParams struct {

}

type DescribeLeafClassificationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLeafClassificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLeafClassificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLeafClassificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLeafClassificationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLeafClassificationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLeafClassificationResponseParams `json:"Response"`
}

func (r *DescribeLeafClassificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLeafClassificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMongoAssetSensitiveDistributionRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

type DescribeMongoAssetSensitiveDistributionRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

func (r *DescribeMongoAssetSensitiveDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMongoAssetSensitiveDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "AssetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMongoAssetSensitiveDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMongoAssetSensitiveDistributionResponseParams struct {
	// mongo的资产统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoAsset *MongoAsset `json:"MongoAsset,omitempty" name:"MongoAsset"`

	// 涉敏top数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopAsset []*TopAsset `json:"TopAsset,omitempty" name:"TopAsset"`

	// mongo的详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongoDetail []*MongoAssetDBDetail `json:"MongoDetail,omitempty" name:"MongoDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMongoAssetSensitiveDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMongoAssetSensitiveDistributionResponseParams `json:"Response"`
}

func (r *DescribeMongoAssetSensitiveDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMongoAssetSensitiveDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRDBAssetSensitiveDistributionRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

type DescribeRDBAssetSensitiveDistributionRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

func (r *DescribeRDBAssetSensitiveDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRDBAssetSensitiveDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "AssetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRDBAssetSensitiveDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRDBAssetSensitiveDistributionResponseParams struct {
	// rdb的资产统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	RDBAsset *RDBAsset `json:"RDBAsset,omitempty" name:"RDBAsset"`

	// 涉敏top数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopAsset []*TopAsset `json:"TopAsset,omitempty" name:"TopAsset"`

	// rdb的详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RDBDetail []*AssetDBDetail `json:"RDBDetail,omitempty" name:"RDBDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRDBAssetSensitiveDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRDBAssetSensitiveDistributionResponseParams `json:"Response"`
}

func (r *DescribeRDBAssetSensitiveDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRDBAssetSensitiveDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportTaskDownloadUrlRequestParams struct {
	// 任务id
	ReportTaskId *uint64 `json:"ReportTaskId,omitempty" name:"ReportTaskId"`

	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 是否同时下载敏感资产详情报告
	IsWithSensitiveDetailReport *bool `json:"IsWithSensitiveDetailReport,omitempty" name:"IsWithSensitiveDetailReport"`
}

type DescribeReportTaskDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	ReportTaskId *uint64 `json:"ReportTaskId,omitempty" name:"ReportTaskId"`

	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 是否同时下载敏感资产详情报告
	IsWithSensitiveDetailReport *bool `json:"IsWithSensitiveDetailReport,omitempty" name:"IsWithSensitiveDetailReport"`
}

func (r *DescribeReportTaskDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportTaskDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReportTaskId")
	delete(f, "DspaId")
	delete(f, "IsWithSensitiveDetailReport")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReportTaskDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportTaskDownloadUrlResponseParams struct {
	// 下载链接集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrlSet []*string `json:"DownloadUrlSet,omitempty" name:"DownloadUrlSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReportTaskDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReportTaskDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeReportTaskDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportTaskDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportTasksRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 报表名称
	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
}

type DescribeReportTasksRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 报表名称
	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`
}

func (r *DescribeReportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ReportName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportTasksResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 报表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*ReportInfo `json:"ItemSet,omitempty" name:"ItemSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReportTasksResponseParams `json:"Response"`
}

func (r *DescribeReportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleDetailRequestParams struct {

}

type DescribeRuleDetailRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleDetailResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleListRequestParams struct {

}

type DescribeRuleListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleListResponseParams `json:"Response"`
}

func (r *DescribeRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSensitiveCOSDataDistributionRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 2331
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

type DescribeSensitiveCOSDataDistributionRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 2331
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

func (r *DescribeSensitiveCOSDataDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSensitiveCOSDataDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "AssetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSensitiveCOSDataDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSensitiveCOSDataDistributionResponseParams struct {
	// 分级分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelDistribution []*Note `json:"LevelDistribution,omitempty" name:"LevelDistribution"`

	// 分类分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryDistribution []*Note `json:"CategoryDistribution,omitempty" name:"CategoryDistribution"`

	// 规则分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleDistribution []*RuleDistribution `json:"RuleDistribution,omitempty" name:"RuleDistribution"`

	// 计算占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveDataNum *int64 `json:"SensitiveDataNum,omitempty" name:"SensitiveDataNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSensitiveCOSDataDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSensitiveCOSDataDistributionResponseParams `json:"Response"`
}

func (r *DescribeSensitiveCOSDataDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSensitiveCOSDataDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSensitiveRDBDataDistributionRequestParams struct {
	// dspa-实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

type DescribeSensitiveRDBDataDistributionRequest struct {
	*tchttp.BaseRequest
	
	// dspa-实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 查询的资产信息列表
	AssetList []*AssetList `json:"AssetList,omitempty" name:"AssetList"`
}

func (r *DescribeSensitiveRDBDataDistributionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSensitiveRDBDataDistributionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "AssetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSensitiveRDBDataDistributionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSensitiveRDBDataDistributionResponseParams struct {
	// 分级分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelDistribution []*Note `json:"LevelDistribution,omitempty" name:"LevelDistribution"`

	// 分类分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryDistribution []*Note `json:"CategoryDistribution,omitempty" name:"CategoryDistribution"`

	// 敏感规则分布详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleDistribution []*RuleDistribution `json:"RuleDistribution,omitempty" name:"RuleDistribution"`

	// 计算占比字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveDataNum *int64 `json:"SensitiveDataNum,omitempty" name:"SensitiveDataNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSensitiveRDBDataDistributionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSensitiveRDBDataDistributionResponseParams `json:"Response"`
}

func (r *DescribeSensitiveRDBDataDistributionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSensitiveRDBDataDistributionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableDSPAMetaResourceAuthRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 用户云资源ID列表。
	ResourceIDs []*string `json:"ResourceIDs,omitempty" name:"ResourceIDs"`
}

type DisableDSPAMetaResourceAuthRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 用户云资源ID列表。
	ResourceIDs []*string `json:"ResourceIDs,omitempty" name:"ResourceIDs"`
}

func (r *DisableDSPAMetaResourceAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableDSPAMetaResourceAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "MetaType")
	delete(f, "ResourceRegion")
	delete(f, "ResourceIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableDSPAMetaResourceAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableDSPAMetaResourceAuthResponseParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 授权结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*DspaTaskResult `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableDSPAMetaResourceAuthResponse struct {
	*tchttp.BaseResponse
	Response *DisableDSPAMetaResourceAuthResponseParams `json:"Response"`
}

func (r *DisableDSPAMetaResourceAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableDSPAMetaResourceAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiscoveryCondition struct {
	// RDB实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RDBInstances []*RDBInstance `json:"RDBInstances,omitempty" name:"RDBInstances"`

	// COS实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	COSInstances []*COSInstance `json:"COSInstances,omitempty" name:"COSInstances"`

	// Mongo实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NOSQLInstances []*NOSQLInstance `json:"NOSQLInstances,omitempty" name:"NOSQLInstances"`

	// ES实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ESInstances []*ESInstance `json:"ESInstances,omitempty" name:"ESInstances"`
}

type DspaAssessmentFilter struct {
	// 过滤类型。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤类型的值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DspaCOSDataAssetCount struct {
	// 数组资产类型，0代表关系型数据库资产，1代表对象存储COS资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataAssetType *int64 `json:"DataAssetType,omitempty" name:"DataAssetType"`

	// 已扫描的存储桶的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalBucketCnt *int64 `json:"TotalBucketCnt,omitempty" name:"TotalBucketCnt"`

	// 对象总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalObjectCnt *int64 `json:"TotalObjectCnt,omitempty" name:"TotalObjectCnt"`

	// 敏感数据类型个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveCategoryCnt *int64 `json:"SensitiveCategoryCnt,omitempty" name:"SensitiveCategoryCnt"`

	// 敏感数据条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveDataCnt *int64 `json:"SensitiveDataCnt,omitempty" name:"SensitiveDataCnt"`

	// 敏感等级分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveLevel []*SensitiveLevel `json:"SensitiveLevel,omitempty" name:"SensitiveLevel"`

	// 敏感存储桶个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveBucketCnt *int64 `json:"SensitiveBucketCnt,omitempty" name:"SensitiveBucketCnt"`

	// 敏感对象个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveObjectCnt *int64 `json:"SensitiveObjectCnt,omitempty" name:"SensitiveObjectCnt"`

	// 数据分类分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryDistributed []*DspaDataCategoryDistributed `json:"CategoryDistributed,omitempty" name:"CategoryDistributed"`
}

type DspaCOSDataAssetDetail struct {
	// 对象桶
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 对象名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 出现次数
	SensitiveDataCount *int64 `json:"SensitiveDataCount,omitempty" name:"SensitiveDataCount"`

	// 敏感数据分类
	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`

	// 敏感等级
	LevelRiskName *string `json:"LevelRiskName,omitempty" name:"LevelRiskName"`

	// KMS加密
	KMSEncrypted *bool `json:"KMSEncrypted,omitempty" name:"KMSEncrypted"`

	// 文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件大小
	FileSize *string `json:"FileSize,omitempty" name:"FileSize"`

	// 敏感数据分级分数
	LevelRiskScore *int64 `json:"LevelRiskScore,omitempty" name:"LevelRiskScore"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 资源所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 分类ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 分级ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 文件扫描结果ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileResultId *int64 `json:"FileResultId,omitempty" name:"FileResultId"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceName *string `json:"DataSourceName,omitempty" name:"DataSourceName"`

	// 分类路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryFullPath *string `json:"CategoryFullPath,omitempty" name:"CategoryFullPath"`
}

type DspaCOSDiscoveryTask struct {
	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 执行周期，0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 执行计划，0立即 1定时
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 任务开关；1 打开，0 关闭
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 数据源对象信息
	DataSourceInfo *DspaCOSDiscoveryTaskDataSourceInfo `json:"DataSourceInfo,omitempty" name:"DataSourceInfo"`

	// 通用规则集开关，0 关闭，1 启用
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 任务最新的一次执行结果信息，该字段用于查询任务列表接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ScanTaskResult `json:"Result,omitempty" name:"Result"`

	// 定时开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`

	// 关联分类模板是否更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceUpdate *bool `json:"ComplianceUpdate,omitempty" name:"ComplianceUpdate"`
}

type DspaCOSDiscoveryTaskDataSourceInfo struct {
	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 代理地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyAddress []*string `json:"ProxyAddress,omitempty" name:"ProxyAddress"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceName *string `json:"DataSourceName,omitempty" name:"DataSourceName"`

	// 扫描任务条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Condition *DspaDiscoveryTaskCOSCondition `json:"Condition,omitempty" name:"Condition"`

	// 资源所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
}

type DspaCOSDiscoveryTaskDetail struct {
	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 执行周期，0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 执行计划，0立即 1定时
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 任务开关；1 打开，0 关闭
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 数据源对象信息
	DataSourceInfo *DspaCOSDiscoveryTaskDataSourceInfo `json:"DataSourceInfo,omitempty" name:"DataSourceInfo"`

	// 通用规则集开关，0 关闭，1 启用
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 当创建任务时，用户打开了通用规则集开关，则该字段就会保存默认合规组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultComplianceInfo []*ScanTaskComplianceInfo `json:"DefaultComplianceInfo,omitempty" name:"DefaultComplianceInfo"`

	// 该任务中用户选择的合规组信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomComplianceInfo []*ScanTaskComplianceInfo `json:"CustomComplianceInfo,omitempty" name:"CustomComplianceInfo"`

	// 定时开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`
}

type DspaCOSDiscoveryTaskResult struct {
	// 扫描bucket结果ID
	BucketResultId *int64 `json:"BucketResultId,omitempty" name:"BucketResultId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 扫描任务最新一次扫描结果ID
	ResultId *int64 `json:"ResultId,omitempty" name:"ResultId"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 桶名称
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 总文件数
	TotalFiles *int64 `json:"TotalFiles,omitempty" name:"TotalFiles"`

	// 被识别出的敏感数据数
	SensitiveDataNums *int64 `json:"SensitiveDataNums,omitempty" name:"SensitiveDataNums"`

	// Bucket扫描的结束时间，格式如：2006-01-02 15:04:05
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceName *string `json:"DataSourceName,omitempty" name:"DataSourceName"`

	// Bucket扫描状态，0待扫描 1扫描中 2扫描终止 3扫描成功 4扫描失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Bucket扫描结果错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`

	// 资源所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 是否超额
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverSize *string `json:"OverSize,omitempty" name:"OverSize"`
}

type DspaCloudResourceMeta struct {
	// 用户资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源VIP。
	ResourceVip *string `json:"ResourceVip,omitempty" name:"ResourceVip"`

	// 资源端口。
	ResourceVPort *uint64 `json:"ResourceVPort,omitempty" name:"ResourceVPort"`

	// 资源被创建时间。
	ResourceCreateTime *string `json:"ResourceCreateTime,omitempty" name:"ResourceCreateTime"`

	// 用户资源VPC ID 字符串。
	ResourceUniqueVpcId *string `json:"ResourceUniqueVpcId,omitempty" name:"ResourceUniqueVpcId"`

	// 用户资源Subnet ID 字符串。
	ResourceUniqueSubnetId *string `json:"ResourceUniqueSubnetId,omitempty" name:"ResourceUniqueSubnetId"`
}

type DspaDataCategoryDistributed struct {
	// 数据分类ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 数据分类名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`

	// 数据分类统计个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 分类路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryFullPath *string `json:"CategoryFullPath,omitempty" name:"CategoryFullPath"`
}

type DspaDataSourceMngFilter struct {
	// 过滤类型。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤类型的值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DspaDiscoveryCOSDataRule struct {
	// 只能取and 、or两个值其中之一，and：字段和内容同时满足，or：字段和内容满足其一.
	// 默认值为or
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 规则内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contents []*DspaDiscoveryDataContent `json:"Contents,omitempty" name:"Contents"`
}

type DspaDiscoveryCOSRules struct {
	// 规则状态；0 不启用, 1 启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// regex规则内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegexRule *DspaDiscoveryCOSDataRule `json:"RegexRule,omitempty" name:"RegexRule"`

	// 关键词规则内容组，最大支持5个关键词。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeywordRule *DspaDiscoveryCOSDataRule `json:"KeywordRule,omitempty" name:"KeywordRule"`

	// 忽略词规则内容组，最大支持5个忽略词。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreStringRule *DspaDiscoveryCOSDataRule `json:"IgnoreStringRule,omitempty" name:"IgnoreStringRule"`

	// 最大匹配距离，默认值为100。上限为500.
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMatch *int64 `json:"MaxMatch,omitempty" name:"MaxMatch"`
}

type DspaDiscoveryCOSTaskResultDetail struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 扫描File结果详情ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileResultId *int64 `json:"FileResultId,omitempty" name:"FileResultId"`

	// 所属桶名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 所属文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 敏感数据分类ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分类名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`

	// 敏感数据分级ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 敏感数据分级名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelName *string `json:"LevelName,omitempty" name:"LevelName"`

	// KMS加密，true or false
	// 注意：此字段可能返回 null，表示取不到有效值。
	KMSEncrypted *bool `json:"KMSEncrypted,omitempty" name:"KMSEncrypted"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 敏感数据分级分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelRiskScore *int64 `json:"LevelRiskScore,omitempty" name:"LevelRiskScore"`

	// 文件大小，单位为KB
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件类型，如csv，txt
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 敏感数据出现次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveDataCount *int64 `json:"SensitiveDataCount,omitempty" name:"SensitiveDataCount"`

	// 分类树路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryFullPath []*string `json:"CategoryFullPath,omitempty" name:"CategoryFullPath"`

	// 合规组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 结果id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultId *int64 `json:"ResultId,omitempty" name:"ResultId"`
}

type DspaDiscoveryComplianceGroup struct {
	// 合规组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceGroupId *int64 `json:"ComplianceGroupId,omitempty" name:"ComplianceGroupId"`

	// 合规组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 合规组描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 合规组类型；0 默认合规组，1 系统合规组（除默认合规组外）, 2 自定义合规组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceGroupType *int64 `json:"ComplianceGroupType,omitempty" name:"ComplianceGroupType"`

	// 合规组对应的规则项
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceGroupRules []*DspaDiscoveryComplianceGroupRuleInfo `json:"ComplianceGroupRules,omitempty" name:"ComplianceGroupRules"`

	// 合规组对应的分级组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelGroupId *uint64 `json:"LevelGroupId,omitempty" name:"LevelGroupId"`
}

type DspaDiscoveryComplianceGroupInfo struct {
	// 合规组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceGroupId *int64 `json:"ComplianceGroupId,omitempty" name:"ComplianceGroupId"`

	// 合规组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 合规组描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 合规组类型；0 默认合规组，1 系统合规组（除默认合规组外）, 2 自定义合规组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceGroupType *int64 `json:"ComplianceGroupType,omitempty" name:"ComplianceGroupType"`

	// 合规组对应的规则项
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceGroupRules []*DspaDiscoveryComplianceGroupRule `json:"ComplianceGroupRules,omitempty" name:"ComplianceGroupRules"`

	// 合规组对应的分级组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelGroupId *uint64 `json:"LevelGroupId,omitempty" name:"LevelGroupId"`

	// 是否禁止使用（true，禁止使用，false，可以使用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`

	// 是否别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAlias *bool `json:"IsAlias,omitempty" name:"IsAlias"`
}

type DspaDiscoveryComplianceGroupRule struct {
	// 敏感数据识别规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 敏感数据识别规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 敏感数据分类ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分级ID, 目前只支持高、中、低三级
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 合规组对应的分类信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`

	// 分级名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelRiskName *string `json:"LevelRiskName,omitempty" name:"LevelRiskName"`
}

type DspaDiscoveryComplianceGroupRuleInfo struct {
	// 敏感数据识别规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 敏感数据识别规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 敏感数据分类ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分级ID, 目前只支持高、中、低三级
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`
}

type DspaDiscoveryDataContent struct {
	// 规则内容，可以是正则规则，关键词，
	// 忽略词扥
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleContent *string `json:"RuleContent,omitempty" name:"RuleContent"`

	// 是否区分大小写
	// false: 不区分大小写
	// true:区分大小写
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsIgnoreCase *bool `json:"IsIgnoreCase,omitempty" name:"IsIgnoreCase"`
}

type DspaDiscoveryDataRule struct {
	// 规则类型；取值：
	// keyword 关键字, 
	// regex 正则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleContent *string `json:"RuleContent,omitempty" name:"RuleContent"`

	// 该字段是针对规则类型RuleType为keyword类型时的一个扩展属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendParameters []*DatagovRuleExtendParameter `json:"ExtendParameters,omitempty" name:"ExtendParameters"`
}

type DspaDiscoveryDataRules struct {
	// 操作符；只能取and, or的其中一种
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contents []*DspaDiscoveryDataRule `json:"Contents,omitempty" name:"Contents"`
}

type DspaDiscoveryLevelDetail struct {
	// 分级组名称，唯一性约束，最多60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelGroupName *string `json:"LevelGroupName,omitempty" name:"LevelGroupName"`

	// 分级组来源，0为内置，1为自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 分级组描述，最多1024字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelGroupDesc *string `json:"LevelGroupDesc,omitempty" name:"LevelGroupDesc"`

	// 具体分级标识详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelDetail []*LevelItem `json:"LevelDetail,omitempty" name:"LevelDetail"`

	// 引用合规组次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefComplianceCnt *uint64 `json:"RefComplianceCnt,omitempty" name:"RefComplianceCnt"`

	// 引用合规组
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefCompliance []*DspaDiscoveryComplianceGroup `json:"RefCompliance,omitempty" name:"RefCompliance"`

	// 分级组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelGroupId *uint64 `json:"LevelGroupId,omitempty" name:"LevelGroupId"`
}

type DspaDiscoveryRDBRules struct {
	// 规则状态；0 不启用, 1 启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 只能取and 、or两个值其中之一，and：字段和内容同时满足，or：字段和内容满足其一
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchOperator *string `json:"MatchOperator,omitempty" name:"MatchOperator"`

	// 字段名包含规则，最大支持选择9项
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaRule *DspaDiscoveryDataRules `json:"MetaRule,omitempty" name:"MetaRule"`

	// 内容包含规则，最大支持选择9项
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentRule *DspaDiscoveryDataRules `json:"ContentRule,omitempty" name:"ContentRule"`
}

type DspaDiscoveryRuleDetail struct {
	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则来源，取值：0 内置, 1 自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// RDB规则详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	RDBRules *DspaDiscoveryRDBRules `json:"RDBRules,omitempty" name:"RDBRules"`

	// COS规则详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	COSRules *DspaDiscoveryCOSRules `json:"COSRules,omitempty" name:"COSRules"`
}

type DspaDiscoveryTaskCOSCondition struct {
	// 数据桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileTypes []*string `json:"FileTypes,omitempty" name:"FileTypes"`

	// 文件大小上限，单位为KB，如1000, 目前单个文件最大只支持1GB（1048576KB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSizeLimit *int64 `json:"FileSizeLimit,omitempty" name:"FileSizeLimit"`
}

type DspaDiscoveryTaskDataSource struct {
	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 用于传入的数据源的条件，可以选择多个数据库，数据库之间通过逗号分隔，如果为空，默认是全部数据库
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 代理地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyAddress []*string `json:"ProxyAddress,omitempty" name:"ProxyAddress"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceName *string `json:"DataSourceName,omitempty" name:"DataSourceName"`

	// 资源所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
}

type DspaDiscoveryTaskDbResult struct {
	// 扫描数据库结果ID
	DbResultId *int64 `json:"DbResultId,omitempty" name:"DbResultId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 扫描任务最新一次扫描结果ID
	ResultId *int64 `json:"ResultId,omitempty" name:"ResultId"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 总表数
	TotalTables *int64 `json:"TotalTables,omitempty" name:"TotalTables"`

	// 敏感表数
	SensitiveTables *int64 `json:"SensitiveTables,omitempty" name:"SensitiveTables"`

	// DB扫描的结束时间，格式如：2006-01-02 15:04:05
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceName *string `json:"DataSourceName,omitempty" name:"DataSourceName"`

	// DB扫描状态，0待扫描 1扫描中 2扫描终止 3扫描成功 4扫描失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// DB扫描结果错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *string `json:"ErrorInfo,omitempty" name:"ErrorInfo"`

	// 资源所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 敏感字段数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveField *int64 `json:"SensitiveField,omitempty" name:"SensitiveField"`

	// 总的字段数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalField *int64 `json:"TotalField,omitempty" name:"TotalField"`
}

type DspaDiscoveryTaskDetail struct {
	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 执行周期，0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 执行计划，0立即 1定时
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 任务开关；1 打开，0 关闭
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 元数据对象信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceInfo *DspaDiscoveryTaskDataSource `json:"DataSourceInfo,omitempty" name:"DataSourceInfo"`

	// 通用规则集开关，0 关闭，1 启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 当创建任务时，用户打开了通用规则集开关，则该字段就会保存默认合规组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultComplianceInfo []*ScanTaskComplianceInfo `json:"DefaultComplianceInfo,omitempty" name:"DefaultComplianceInfo"`

	// 该任务中用户选择的合规组信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomComplianceInfo []*ScanTaskComplianceInfo `json:"CustomComplianceInfo,omitempty" name:"CustomComplianceInfo"`

	// 定时开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`
}

type DspaDiscoveryTaskResultDetail struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 扫描结果详情ID
	FieldResultId *int64 `json:"FieldResultId,omitempty" name:"FieldResultId"`

	// 所属数据表名
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 字段名
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 敏感数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分类名称
	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`

	// 敏感数据分级ID
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 敏感数据分级名称
	LevelName *string `json:"LevelName,omitempty" name:"LevelName"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 敏感数据分级分数
	LevelRiskScore *int64 `json:"LevelRiskScore,omitempty" name:"LevelRiskScore"`

	// 保护措施
	// 注意：此字段可能返回 null，表示取不到有效值。
	SafeGuard *DspaSafeGuard `json:"SafeGuard,omitempty" name:"SafeGuard"`

	// 分类路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryFullPath *string `json:"CategoryFullPath,omitempty" name:"CategoryFullPath"`
}

type DspaFieldResultDataSample struct {
	// 数据样本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSample *string `json:"DataSample,omitempty" name:"DataSample"`
}

type DspaInstance struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// DSPA实例名称。
	DspaName *string `json:"DspaName,omitempty" name:"DspaName"`

	// DSPA实例描述信息。
	DspaDescription *string `json:"DspaDescription,omitempty" name:"DspaDescription"`

	// DSPA实例已授权的数据库实例数量。
	DBAuthCount *uint64 `json:"DBAuthCount,omitempty" name:"DBAuthCount"`

	// DSPA实例已绑定的cos桶数量。
	CosBindCount *uint64 `json:"CosBindCount,omitempty" name:"CosBindCount"`

	// DSPA实例版本。
	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`

	// DSPA实例状态。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例过期时间戳。
	ExpiredAt *uint64 `json:"ExpiredAt,omitempty" name:"ExpiredAt"`

	// 账户APPID。
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 体验版本信息。
	TrialVersion *string `json:"TrialVersion,omitempty" name:"TrialVersion"`

	// 体验版本过期时间戳。
	TrialEndAt *uint64 `json:"TrialEndAt,omitempty" name:"TrialEndAt"`

	// DB已购配额。
	DbTotalQuota *int64 `json:"DbTotalQuota,omitempty" name:"DbTotalQuota"`

	// COS已购配额。
	CosTotalQuota *int64 `json:"CosTotalQuota,omitempty" name:"CosTotalQuota"`

	// COS配额单位，比如TB。
	CosQuotaUnit *string `json:"CosQuotaUnit,omitempty" name:"CosQuotaUnit"`

	// 0: 默认状态(用户未设置)
	// 1: 开启自动续费
	// 2: 明确不自动续费
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type DspaRDBDataAssetCount struct {
	// 数组资产类型，0代表关系型数据库资产，1代表对象存储COS资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataAssetType *int64 `json:"DataAssetType,omitempty" name:"DataAssetType"`

	// 已扫描的数据库的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalDbCnt *int64 `json:"TotalDbCnt,omitempty" name:"TotalDbCnt"`

	// 数据库表的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTableCnt *int64 `json:"TotalTableCnt,omitempty" name:"TotalTableCnt"`

	// 敏感数据类型个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveCategoryCnt *int64 `json:"SensitiveCategoryCnt,omitempty" name:"SensitiveCategoryCnt"`

	// 敏感字段的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveFieldCnt *int64 `json:"SensitiveFieldCnt,omitempty" name:"SensitiveFieldCnt"`

	// 敏感等级分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveLevel []*SensitiveLevel `json:"SensitiveLevel,omitempty" name:"SensitiveLevel"`

	// 敏感数据库的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveDbCnt *int64 `json:"SensitiveDbCnt,omitempty" name:"SensitiveDbCnt"`

	// 敏感数据库表的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveTableCnt *int64 `json:"SensitiveTableCnt,omitempty" name:"SensitiveTableCnt"`

	// 扫描字段的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalFieldCnt *int64 `json:"TotalFieldCnt,omitempty" name:"TotalFieldCnt"`

	// 数据分类分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryDistributed []*DspaDataCategoryDistributed `json:"CategoryDistributed,omitempty" name:"CategoryDistributed"`
}

type DspaRDBDataAssetDetail struct {
	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 数据库类型
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 数据库表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 数据库表字段名称
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 数据分类
	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`

	// 敏感等级
	LevelRiskName *string `json:"LevelRiskName,omitempty" name:"LevelRiskName"`

	// 分级风险分数，1-10，最小值为1，最大值为10
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelRiskScore *int64 `json:"LevelRiskScore,omitempty" name:"LevelRiskScore"`

	// 可信分
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrustedScore *string `json:"TrustedScore,omitempty" name:"TrustedScore"`

	// 资源所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 字段扫描结果ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldResultId *string `json:"FieldResultId,omitempty" name:"FieldResultId"`

	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 分级ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 分类ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceName *string `json:"DataSourceName,omitempty" name:"DataSourceName"`

	// 保护措施
	// 注意：此字段可能返回 null，表示取不到有效值。
	SafeGuard *DspaSafeGuard `json:"SafeGuard,omitempty" name:"SafeGuard"`

	// 分类路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryFullPath *string `json:"CategoryFullPath,omitempty" name:"CategoryFullPath"`
}

type DspaResourceAccount struct {
	// 资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 用户名。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码。
	Password *string `json:"Password,omitempty" name:"Password"`
}

type DspaSafeGuard struct {
	// 加密状态，可取值如下：
	// UNSET 未设置
	// DISABLE 规则设置未启用
	// ENABLE 规则设置并启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`

	// 脱敏状态，可取值如下：
	// UNSET 未设置
	// DISABLE 规则设置未启用
	// ENABLE 规则设置并启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desensitization *string `json:"Desensitization,omitempty" name:"Desensitization"`
}

type DspaTaskResult struct {
	// 任务结果。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 结果描述。
	ResultDescription *string `json:"ResultDescription,omitempty" name:"ResultDescription"`

	// 资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`
}

type DspaUserResourceMeta struct {
	// 用户资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源VIP。
	ResourceVip *string `json:"ResourceVip,omitempty" name:"ResourceVip"`

	// 资源端口。
	ResourceVPort *uint64 `json:"ResourceVPort,omitempty" name:"ResourceVPort"`

	// 资源被创建时间。
	ResourceCreateTime *string `json:"ResourceCreateTime,omitempty" name:"ResourceCreateTime"`

	// 用户资源VPC ID 字符串。
	ResourceUniqueVpcId *string `json:"ResourceUniqueVpcId,omitempty" name:"ResourceUniqueVpcId"`

	// 用户资源Subnet ID 字符串。
	ResourceUniqueSubnetId *string `json:"ResourceUniqueSubnetId,omitempty" name:"ResourceUniqueSubnetId"`

	// 用户资源类型信息。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源被同步时间。
	ResourceSyncTime *string `json:"ResourceSyncTime,omitempty" name:"ResourceSyncTime"`

	// 资源被授权状态。
	AuthStatus *string `json:"AuthStatus,omitempty" name:"AuthStatus"`

	// 资源创建类型，cloud-云原生资源，build-用户自建资源。
	BuildType *string `json:"BuildType,omitempty" name:"BuildType"`

	// 主实例ID。
	MasterInsId *string `json:"MasterInsId,omitempty" name:"MasterInsId"`

	// 用户资源VPC ID 整数。
	ResourceVpcId *uint64 `json:"ResourceVpcId,omitempty" name:"ResourceVpcId"`

	// 用户资源Subnet ID 整数。
	ResourceSubnetId *uint64 `json:"ResourceSubnetId,omitempty" name:"ResourceSubnetId"`

	// 协议类型。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 资源版本号。
	ResourceVersion *string `json:"ResourceVersion,omitempty" name:"ResourceVersion"`

	// 授权方式
	ResourceAuthType *string `json:"ResourceAuthType,omitempty" name:"ResourceAuthType"`

	// 授权账号名
	ResourceAuthAccount *string `json:"ResourceAuthAccount,omitempty" name:"ResourceAuthAccount"`
}

type ESAsset struct {
	// 索引总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexNums *int64 `json:"IndexNums,omitempty" name:"IndexNums"`

	// 敏感索引的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveIndexNums *int64 `json:"SensitiveIndexNums,omitempty" name:"SensitiveIndexNums"`

	// 字段数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldNums *int64 `json:"FieldNums,omitempty" name:"FieldNums"`

	// 敏感的字段数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveFieldNums *int64 `json:"SensitiveFieldNums,omitempty" name:"SensitiveFieldNums"`
}

type ESAssetDBDetail struct {
	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 索引名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexName *string `json:"IndexName,omitempty" name:"IndexName"`

	// 数据库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *string `json:"DataType,omitempty" name:"DataType"`

	// 字段的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldNums *int64 `json:"FieldNums,omitempty" name:"FieldNums"`

	// 敏感字段的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveFieldNums *int64 `json:"SensitiveFieldNums,omitempty" name:"SensitiveFieldNums"`

	// 敏感数据分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributionData []*Note `json:"DistributionData,omitempty" name:"DistributionData"`
}

type ESDataAssetCountDto struct {
	// es
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataAssetType *int64 `json:"DataAssetType,omitempty" name:"DataAssetType"`

	// 敏感索引个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveIndexCnt *int64 `json:"SensitiveIndexCnt,omitempty" name:"SensitiveIndexCnt"`

	// 总的索引个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalIndexCnt *int64 `json:"TotalIndexCnt,omitempty" name:"TotalIndexCnt"`

	// 敏感字段个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveFieldCnt *int64 `json:"SensitiveFieldCnt,omitempty" name:"SensitiveFieldCnt"`

	// 总的字段个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalFieldCnt *int64 `json:"TotalFieldCnt,omitempty" name:"TotalFieldCnt"`

	// 敏感分类的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveCategoryCnt *int64 `json:"SensitiveCategoryCnt,omitempty" name:"SensitiveCategoryCnt"`

	// 敏感分级的分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveLevel []*SensitiveLevel `json:"SensitiveLevel,omitempty" name:"SensitiveLevel"`

	// 敏感分类的分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryDistributed []*DspaDataCategoryDistributed `json:"CategoryDistributed,omitempty" name:"CategoryDistributed"`
}

type ESDataAssetDetail struct {
	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldResultId *int64 `json:"FieldResultId,omitempty" name:"FieldResultId"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceName *string `json:"DataSourceName,omitempty" name:"DataSourceName"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 索引名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexName *string `json:"IndexName,omitempty" name:"IndexName"`

	// 字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 分类id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 分类名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`

	// 分类路径数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryArr []*string `json:"CategoryArr,omitempty" name:"CategoryArr"`

	// 等级id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 分级名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelRiskName *string `json:"LevelRiskName,omitempty" name:"LevelRiskName"`

	// 分级分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelRiskScore *int64 `json:"LevelRiskScore,omitempty" name:"LevelRiskScore"`

	// 可信分
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrustedScore *float64 `json:"TrustedScore,omitempty" name:"TrustedScore"`

	// 规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

type ESInstance struct {
	// 数据源id
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 扫描任务ID
	DiscoveryTaskId *int64 `json:"DiscoveryTaskId,omitempty" name:"DiscoveryTaskId"`

	// 扫描任务实例ID
	DiscoveryTaskInstanceID *int64 `json:"DiscoveryTaskInstanceID,omitempty" name:"DiscoveryTaskInstanceID"`
}

type ESTaskResultDetail struct {
	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 分类id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 分类名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryName *string `json:"CategoryName,omitempty" name:"CategoryName"`

	// 多级分类的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryArr []*string `json:"CategoryArr,omitempty" name:"CategoryArr"`

	// 分级id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 分级名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelName *string `json:"LevelName,omitempty" name:"LevelName"`
}

// Predefined struct for user
type EnableDSPADiscoveryRuleRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 开关ScanRule
	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type EnableDSPADiscoveryRuleRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 开关ScanRule
	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

func (r *EnableDSPADiscoveryRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableDSPADiscoveryRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "RuleId")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableDSPADiscoveryRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableDSPADiscoveryRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableDSPADiscoveryRuleResponse struct {
	*tchttp.BaseResponse
	Response *EnableDSPADiscoveryRuleResponseParams `json:"Response"`
}

func (r *EnableDSPADiscoveryRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableDSPADiscoveryRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTrialVersionRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 体验版本名称。
	TrialVersion *string `json:"TrialVersion,omitempty" name:"TrialVersion"`
}

type EnableTrialVersionRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 体验版本名称。
	TrialVersion *string `json:"TrialVersion,omitempty" name:"TrialVersion"`
}

func (r *EnableTrialVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTrialVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TrialVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableTrialVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTrialVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableTrialVersionResponse struct {
	*tchttp.BaseResponse
	Response *EnableTrialVersionResponseParams `json:"Response"`
}

func (r *EnableTrialVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTrialVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrDescription struct {
	// 错误码。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 具体错误信息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`
}

// Predefined struct for user
type ExportAssetDetailDataRequestParams struct {
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 资产类型（rdb,cvm_db,cos）
	MetaDataType *string `json:"MetaDataType,omitempty" name:"MetaDataType"`

	// 过滤数组。支持的Name：
	// DataSourceID 数据源ID
	// DbName 数据库名称
	// CategoryID 敏感数据分类ID
	// RuleID 规则ID
	// LevelID 敏感分级ID
	// ResourceRegion 资源所在地域
	// DataSourceType 数据源类型，不填默认过滤非自建的所有关系型数据源类型，填selfbuilt-db只过滤自建类型
	// 注意：每个name默认支持最多5个values。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type ExportAssetDetailDataRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 资产类型（rdb,cvm_db,cos）
	MetaDataType *string `json:"MetaDataType,omitempty" name:"MetaDataType"`

	// 过滤数组。支持的Name：
	// DataSourceID 数据源ID
	// DbName 数据库名称
	// CategoryID 敏感数据分类ID
	// RuleID 规则ID
	// LevelID 敏感分级ID
	// ResourceRegion 资源所在地域
	// DataSourceType 数据源类型，不填默认过滤非自建的所有关系型数据源类型，填selfbuilt-db只过滤自建类型
	// 注意：每个name默认支持最多5个values。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportAssetDetailDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAssetDetailDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "MetaDataType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportAssetDetailDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportAssetDetailDataResponseParams struct {
	// 导出任务id
	ExportTaskId *uint64 `json:"ExportTaskId,omitempty" name:"ExportTaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportAssetDetailDataResponse struct {
	*tchttp.BaseResponse
	Response *ExportAssetDetailDataResponseParams `json:"Response"`
}

func (r *ExportAssetDetailDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAssetDetailDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type GetResourceConnectionStatusRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源列表中展示的资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type GetResourceConnectionStatusRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源列表中展示的资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *GetResourceConnectionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourceConnectionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "MetaType")
	delete(f, "ResourceRegion")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetResourceConnectionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetResourceConnectionStatusResponseParams struct {
	// 连接状态，success -- 连接成功，failed -- 连接失败
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" name:"ConnectionStatus"`

	// 连接状态的描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionDesc *string `json:"ConnectionDesc,omitempty" name:"ConnectionDesc"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetResourceConnectionStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetResourceConnectionStatusResponseParams `json:"Response"`
}

func (r *GetResourceConnectionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourceConnectionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTrialVersionRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

type GetTrialVersionRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

func (r *GetTrialVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTrialVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTrialVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTrialVersionResponseParams struct {
	// 体验版本名称。
	TrialVersion *string `json:"TrialVersion,omitempty" name:"TrialVersion"`

	// 版本体验结束时间戳。
	TrialEndAt *uint64 `json:"TrialEndAt,omitempty" name:"TrialEndAt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTrialVersionResponse struct {
	*tchttp.BaseResponse
	Response *GetTrialVersionResponseParams `json:"Response"`
}

func (r *GetTrialVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTrialVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserQuotaInfoRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

type GetUserQuotaInfoRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`
}

func (r *GetUserQuotaInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserQuotaInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUserQuotaInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserQuotaInfoResponseParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 用户购买的DB配额。
	DbTotalQuota *int64 `json:"DbTotalQuota,omitempty" name:"DbTotalQuota"`

	// 用户购买的COS存储量配额。
	CosTotalQuota *int64 `json:"CosTotalQuota,omitempty" name:"CosTotalQuota"`

	// 用户可用的DB配额。
	DbRemainQuota *int64 `json:"DbRemainQuota,omitempty" name:"DbRemainQuota"`

	// 用户可用的COS存储量配额。
	CosRemainQuota *float64 `json:"CosRemainQuota,omitempty" name:"CosRemainQuota"`

	// COS存储量单位，比如TB。
	CosQuotaUnit *string `json:"CosQuotaUnit,omitempty" name:"CosQuotaUnit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetUserQuotaInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetUserQuotaInfoResponseParams `json:"Response"`
}

func (r *GetUserQuotaInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserQuotaInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HighRiskAssetsDetail struct {
	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资产对象名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetsName *string `json:"AssetsName,omitempty" name:"AssetsName"`

	// 高风险个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighRiskCount *int64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`

	// 风险类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`

	// 总的风险个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalRiskCount *int64 `json:"TotalRiskCount,omitempty" name:"TotalRiskCount"`
}

type ItemLevel struct {
	// 分级标识名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	LevelRiskName *string `json:"LevelRiskName,omitempty" name:"LevelRiskName"`

	// 分级标识对应的风险分数值，1-10，最小为1，最大为10
	LevelRiskScore *int64 `json:"LevelRiskScore,omitempty" name:"LevelRiskScore"`
}

type LevelItem struct {
	// 分级ID
	LevelId *uint64 `json:"LevelId,omitempty" name:"LevelId"`

	// 分级组ID
	LevelGroupId *uint64 `json:"LevelGroupId,omitempty" name:"LevelGroupId"`

	// 分级标识名称，支持内置分级，内置分级取值：高，中，低，也可以自定义
	LevelRiskName *string `json:"LevelRiskName,omitempty" name:"LevelRiskName"`

	// 分级风险分数，1-10，最小值为1，最大值为10
	LevelRiskScore *uint64 `json:"LevelRiskScore,omitempty" name:"LevelRiskScore"`
}

// Predefined struct for user
type ListDSPAClustersRequestParams struct {
	// 分页步长，默认为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤项。
	// 支持的过滤项包括：DspaId、Status、Version、DspaName。
	// DspaId和DspaName支持模糊搜索。
	// Status支持的可选值：enabled、disabled。
	// Version支持的可选值：trial、official。
	Filters []*DspaDataSourceMngFilter `json:"Filters,omitempty" name:"Filters"`

	// 展示模式。
	ListMode *string `json:"ListMode,omitempty" name:"ListMode"`
}

type ListDSPAClustersRequest struct {
	*tchttp.BaseRequest
	
	// 分页步长，默认为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤项。
	// 支持的过滤项包括：DspaId、Status、Version、DspaName。
	// DspaId和DspaName支持模糊搜索。
	// Status支持的可选值：enabled、disabled。
	// Version支持的可选值：trial、official。
	Filters []*DspaDataSourceMngFilter `json:"Filters,omitempty" name:"Filters"`

	// 展示模式。
	ListMode *string `json:"ListMode,omitempty" name:"ListMode"`
}

func (r *ListDSPAClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDSPAClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ListMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDSPAClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDSPAClustersResponseParams struct {
	// 资源总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资源列表。
	InstanceList []*DspaInstance `json:"InstanceList,omitempty" name:"InstanceList"`

	// 是否被拒绝访问所有dspa实例资源。
	DenyAll *bool `json:"DenyAll,omitempty" name:"DenyAll"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListDSPAClustersResponse struct {
	*tchttp.BaseResponse
	Response *ListDSPAClustersResponseParams `json:"Response"`
}

func (r *ListDSPAClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDSPAClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDSPACosMetaResourcesRequestParams struct {
	// 实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 过滤数组。支持的Name：
	// Bucket - 桶名，支持模糊匹配
	// 
	// ResoureRegion - 资源所处地域，需要填写完整地域名称，不支持模糊匹配。
	// 
	// Valid -- 资源是否有效，"1" 表示有效，"0"表示无效。
	Filters []*DspaDataSourceMngFilter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源绑定状态过滤，默认为全部
	BindType *string `json:"BindType,omitempty" name:"BindType"`
}

type ListDSPACosMetaResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 过滤数组。支持的Name：
	// Bucket - 桶名，支持模糊匹配
	// 
	// ResoureRegion - 资源所处地域，需要填写完整地域名称，不支持模糊匹配。
	// 
	// Valid -- 资源是否有效，"1" 表示有效，"0"表示无效。
	Filters []*DspaDataSourceMngFilter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源绑定状态过滤，默认为全部
	BindType *string `json:"BindType,omitempty" name:"BindType"`
}

func (r *ListDSPACosMetaResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDSPACosMetaResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDSPACosMetaResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDSPACosMetaResourcesResponseParams struct {
	// 符合条件的COS元数据数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// COS元数据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DSPACosMetaDataInfo `json:"Items,omitempty" name:"Items"`

	// DSPA实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListDSPACosMetaResourcesResponse struct {
	*tchttp.BaseResponse
	Response *ListDSPACosMetaResourcesResponseParams `json:"Response"`
}

func (r *ListDSPACosMetaResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDSPACosMetaResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDSPAMetaResourcesRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 过滤项。
	// 可过滤值包括：
	// ResoureRegion - 资源所处地域，需要填写完整地域名称，不支持模糊匹配。
	// 
	// AuthStatus - authorized（已授权）、unauthorized（未授权）、deleted（资源已被删除），不支持模糊匹配，需要填写完整。
	// 
	// BuildType - cloud（云原生资源）、build（用户自建资源），不支持模糊匹配，需要填写完整。
	// 
	// MetaType - cdb（云数据Mysql）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbmysql（TDSQL-C MySQL版）、cos（对象存储）、mysql_like_proto（自建型Mysql协议类关系型数据库）、postgre_like_proto（自建型Postgre协议类关系型数据库）。
	// 
	// ResourceId - 资源ID，支持模糊搜索。
	Filters []*DspaDataSourceMngFilter `json:"Filters,omitempty" name:"Filters"`

	// 分页步长，默认为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 资源绑定状态过滤，默认为全部
	BindType *string `json:"BindType,omitempty" name:"BindType"`
}

type ListDSPAMetaResourcesRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 过滤项。
	// 可过滤值包括：
	// ResoureRegion - 资源所处地域，需要填写完整地域名称，不支持模糊匹配。
	// 
	// AuthStatus - authorized（已授权）、unauthorized（未授权）、deleted（资源已被删除），不支持模糊匹配，需要填写完整。
	// 
	// BuildType - cloud（云原生资源）、build（用户自建资源），不支持模糊匹配，需要填写完整。
	// 
	// MetaType - cdb（云数据Mysql）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbmysql（TDSQL-C MySQL版）、cos（对象存储）、mysql_like_proto（自建型Mysql协议类关系型数据库）、postgre_like_proto（自建型Postgre协议类关系型数据库）。
	// 
	// ResourceId - 资源ID，支持模糊搜索。
	Filters []*DspaDataSourceMngFilter `json:"Filters,omitempty" name:"Filters"`

	// 分页步长，默认为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 资源绑定状态过滤，默认为全部
	BindType *string `json:"BindType,omitempty" name:"BindType"`
}

func (r *ListDSPAMetaResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDSPAMetaResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDSPAMetaResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDSPAMetaResourcesResponseParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 用户资源列表。
	Resources []*DspaUserResourceMeta `json:"Resources,omitempty" name:"Resources"`

	// 资源总量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListDSPAMetaResourcesResponse struct {
	*tchttp.BaseResponse
	Response *ListDSPAMetaResourcesResponseParams `json:"Response"`
}

func (r *ListDSPAMetaResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDSPAMetaResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClassificationRuleRequestParams struct {

}

type ModifyClassificationRuleRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyClassificationRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClassificationRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClassificationRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClassificationRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClassificationRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClassificationRuleResponseParams `json:"Response"`
}

func (r *ModifyClassificationRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClassificationRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClassificationRuleStateRequestParams struct {

}

type ModifyClassificationRuleStateRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyClassificationRuleStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClassificationRuleStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClassificationRuleStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClassificationRuleStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClassificationRuleStateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClassificationRuleStateResponseParams `json:"Response"`
}

func (r *ModifyClassificationRuleStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClassificationRuleStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAAssessmentRiskLatestRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 最新风险项Id
	RiskLatestTableId *uint64 `json:"RiskLatestTableId,omitempty" name:"RiskLatestTableId"`

	// 风险状态（waiting:待处理，processing:处理中，finished:已处理，ignored:已忽略）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注
	Note *string `json:"Note,omitempty" name:"Note"`

	// 处置人
	ProcessPeople *string `json:"ProcessPeople,omitempty" name:"ProcessPeople"`
}

type ModifyDSPAAssessmentRiskLatestRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 最新风险项Id
	RiskLatestTableId *uint64 `json:"RiskLatestTableId,omitempty" name:"RiskLatestTableId"`

	// 风险状态（waiting:待处理，processing:处理中，finished:已处理，ignored:已忽略）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注
	Note *string `json:"Note,omitempty" name:"Note"`

	// 处置人
	ProcessPeople *string `json:"ProcessPeople,omitempty" name:"ProcessPeople"`
}

func (r *ModifyDSPAAssessmentRiskLatestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAAssessmentRiskLatestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "RiskLatestTableId")
	delete(f, "Status")
	delete(f, "Note")
	delete(f, "ProcessPeople")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPAAssessmentRiskLatestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAAssessmentRiskLatestResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPAAssessmentRiskLatestResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPAAssessmentRiskLatestResponseParams `json:"Response"`
}

func (r *ModifyDSPAAssessmentRiskLatestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAAssessmentRiskLatestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAAssessmentRiskLevelRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 风险等级名称
	RiskLevelName *string `json:"RiskLevelName,omitempty" name:"RiskLevelName"`

	// 风险的描述
	RiskLevelDescription *string `json:"RiskLevelDescription,omitempty" name:"RiskLevelDescription"`

	// 风险id
	RiskId *int64 `json:"RiskId,omitempty" name:"RiskId"`

	// 需要修改的风险列表
	ModifyRiskItem []*RiskLevelMatrix `json:"ModifyRiskItem,omitempty" name:"ModifyRiskItem"`
}

type ModifyDSPAAssessmentRiskLevelRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 风险等级名称
	RiskLevelName *string `json:"RiskLevelName,omitempty" name:"RiskLevelName"`

	// 风险的描述
	RiskLevelDescription *string `json:"RiskLevelDescription,omitempty" name:"RiskLevelDescription"`

	// 风险id
	RiskId *int64 `json:"RiskId,omitempty" name:"RiskId"`

	// 需要修改的风险列表
	ModifyRiskItem []*RiskLevelMatrix `json:"ModifyRiskItem,omitempty" name:"ModifyRiskItem"`
}

func (r *ModifyDSPAAssessmentRiskLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAAssessmentRiskLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "RiskLevelName")
	delete(f, "RiskLevelDescription")
	delete(f, "RiskId")
	delete(f, "ModifyRiskItem")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPAAssessmentRiskLevelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAAssessmentRiskLevelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPAAssessmentRiskLevelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPAAssessmentRiskLevelResponseParams `json:"Response"`
}

func (r *ModifyDSPAAssessmentRiskLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAAssessmentRiskLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAAssessmentRiskRequestParams struct {
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估任务Id，格式“task-xxxxxxxx”
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 风险项Id，格式“risk-xxxxxxxx”
	RiskId *string `json:"RiskId,omitempty" name:"RiskId"`

	// 风险项状态。（waiting:待处理，processing:处理中，finished:已处理）
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyDSPAAssessmentRiskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估任务Id，格式“task-xxxxxxxx”
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 风险项Id，格式“risk-xxxxxxxx”
	RiskId *string `json:"RiskId,omitempty" name:"RiskId"`

	// 风险项状态。（waiting:待处理，processing:处理中，finished:已处理）
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDSPAAssessmentRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAAssessmentRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "RiskId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPAAssessmentRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAAssessmentRiskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPAAssessmentRiskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPAAssessmentRiskResponseParams `json:"Response"`
}

func (r *ModifyDSPAAssessmentRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAAssessmentRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAAssessmentRiskTemplateRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 模版名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模版的描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 修改的风险等级id
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" name:"RiskLevelId"`

	// 脆弱项列表
	RiskIdList []*int64 `json:"RiskIdList,omitempty" name:"RiskIdList"`
}

type ModifyDSPAAssessmentRiskTemplateRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 模版名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模版的描述
	TemplateDescription *string `json:"TemplateDescription,omitempty" name:"TemplateDescription"`

	// 模版id
	TemplateId *int64 `json:"TemplateId,omitempty" name:"TemplateId"`

	// 修改的风险等级id
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" name:"RiskLevelId"`

	// 脆弱项列表
	RiskIdList []*int64 `json:"RiskIdList,omitempty" name:"RiskIdList"`
}

func (r *ModifyDSPAAssessmentRiskTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAAssessmentRiskTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TemplateName")
	delete(f, "TemplateDescription")
	delete(f, "TemplateId")
	delete(f, "RiskLevelId")
	delete(f, "RiskIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPAAssessmentRiskTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAAssessmentRiskTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPAAssessmentRiskTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPAAssessmentRiskTemplateResponseParams `json:"Response"`
}

func (r *ModifyDSPAAssessmentRiskTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAAssessmentRiskTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPACOSDiscoveryTaskRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 任务开关，0 关闭，1 启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 通用规则集开关；0 关闭，1 启用
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 合规组ID列表，最多支持添加5个
	ComplianceGroupIds []*int64 `json:"ComplianceGroupIds,omitempty" name:"ComplianceGroupIds"`

	// 执行计划； 0立即 1定时，选择“立即”时，扫描周期只能选择单次
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 扫描周期；0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 任务定时启动时间，格式：2006-01-02 15:04:05
	// 当执行计划（Plan字段）为”立即“时，定时启动时间不会生效，此场景下给该字段传值不会被保存。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`

	// 待扫描文件类型，用逗号隔开，格式如：[".txt", ".csv", ".log", ".xml",".html", ".json"]。
	FileTypes []*string `json:"FileTypes,omitempty" name:"FileTypes"`

	// 文件大小上限，单位为KB，如1000, 目前单个文件最大只支持100MB（102400KB）
	FileSizeLimit *int64 `json:"FileSizeLimit,omitempty" name:"FileSizeLimit"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
}

type ModifyDSPACOSDiscoveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 任务开关，0 关闭，1 启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 通用规则集开关；0 关闭，1 启用
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 合规组ID列表，最多支持添加5个
	ComplianceGroupIds []*int64 `json:"ComplianceGroupIds,omitempty" name:"ComplianceGroupIds"`

	// 执行计划； 0立即 1定时，选择“立即”时，扫描周期只能选择单次
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 扫描周期；0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 任务定时启动时间，格式：2006-01-02 15:04:05
	// 当执行计划（Plan字段）为”立即“时，定时启动时间不会生效，此场景下给该字段传值不会被保存。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`

	// 待扫描文件类型，用逗号隔开，格式如：[".txt", ".csv", ".log", ".xml",".html", ".json"]。
	FileTypes []*string `json:"FileTypes,omitempty" name:"FileTypes"`

	// 文件大小上限，单位为KB，如1000, 目前单个文件最大只支持100MB（102400KB）
	FileSizeLimit *int64 `json:"FileSizeLimit,omitempty" name:"FileSizeLimit"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`
}

func (r *ModifyDSPACOSDiscoveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPACOSDiscoveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Enable")
	delete(f, "GeneralRuleSetEnable")
	delete(f, "ComplianceGroupIds")
	delete(f, "Plan")
	delete(f, "Period")
	delete(f, "TimingStartTime")
	delete(f, "FileTypes")
	delete(f, "FileSizeLimit")
	delete(f, "ResourceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPACOSDiscoveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPACOSDiscoveryTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPACOSDiscoveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPACOSDiscoveryTaskResponseParams `json:"Response"`
}

func (r *ModifyDSPACOSDiscoveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPACOSDiscoveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPACOSTaskResultRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 文件扫描结果ID
	FileResultId *int64 `json:"FileResultId,omitempty" name:"FileResultId"`

	// 是否设置为非敏感文件
	IsSetNonSensitiveFile *bool `json:"IsSetNonSensitiveFile,omitempty" name:"IsSetNonSensitiveFile"`
}

type ModifyDSPACOSTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 文件扫描结果ID
	FileResultId *int64 `json:"FileResultId,omitempty" name:"FileResultId"`

	// 是否设置为非敏感文件
	IsSetNonSensitiveFile *bool `json:"IsSetNonSensitiveFile,omitempty" name:"IsSetNonSensitiveFile"`
}

func (r *ModifyDSPACOSTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPACOSTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceId")
	delete(f, "FileResultId")
	delete(f, "IsSetNonSensitiveFile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPACOSTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPACOSTaskResultResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPACOSTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPACOSTaskResultResponseParams `json:"Response"`
}

func (r *ModifyDSPACOSTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPACOSTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPACategoryRelationRequestParams struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 当前分类id
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 合并到的分类id
	MergedCategoryId *int64 `json:"MergedCategoryId,omitempty" name:"MergedCategoryId"`

	// 合规组模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

type ModifyDSPACategoryRelationRequest struct {
	*tchttp.BaseRequest
	
	// dspa实例id
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 当前分类id
	CategoryId *uint64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 合并到的分类id
	MergedCategoryId *int64 `json:"MergedCategoryId,omitempty" name:"MergedCategoryId"`

	// 合规组模板id
	ComplianceId *uint64 `json:"ComplianceId,omitempty" name:"ComplianceId"`
}

func (r *ModifyDSPACategoryRelationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPACategoryRelationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "CategoryId")
	delete(f, "MergedCategoryId")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPACategoryRelationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPACategoryRelationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPACategoryRelationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPACategoryRelationResponseParams `json:"Response"`
}

func (r *ModifyDSPACategoryRelationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPACategoryRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPACategoryRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分类名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ModifyDSPACategoryRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据分类ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 敏感数据分类名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyDSPACategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPACategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "CategoryId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPACategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPACategoryResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPACategoryResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPACategoryResponseParams `json:"Response"`
}

func (r *ModifyDSPACategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPACategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAClusterInfoRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// DSPA实例名。1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字。
	DspaName *string `json:"DspaName,omitempty" name:"DspaName"`

	// DSPA实例描述信息。最长1024个字符。
	DspaDescription *string `json:"DspaDescription,omitempty" name:"DspaDescription"`
}

type ModifyDSPAClusterInfoRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// DSPA实例名。1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字。
	DspaName *string `json:"DspaName,omitempty" name:"DspaName"`

	// DSPA实例描述信息。最长1024个字符。
	DspaDescription *string `json:"DspaDescription,omitempty" name:"DspaDescription"`
}

func (r *ModifyDSPAClusterInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAClusterInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "DspaName")
	delete(f, "DspaDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPAClusterInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAClusterInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPAClusterInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPAClusterInfoResponseParams `json:"Response"`
}

func (r *ModifyDSPAClusterInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAClusterInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAComplianceGroupRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceGroupId *int64 `json:"ComplianceGroupId,omitempty" name:"ComplianceGroupId"`

	// 合规组名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 合规组描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 合规组规则配置（参数已废弃，请传空数组）
	ComplianceGroupRules []*ComplianceGroupRuleIdInfo `json:"ComplianceGroupRules,omitempty" name:"ComplianceGroupRules"`

	// 分级组ID，新增参数，可选参数，默认值为1
	LevelGroupId *uint64 `json:"LevelGroupId,omitempty" name:"LevelGroupId"`

	// 是否开启别名
	RuleAlias *bool `json:"RuleAlias,omitempty" name:"RuleAlias"`
}

type ModifyDSPAComplianceGroupRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 合规组ID
	ComplianceGroupId *int64 `json:"ComplianceGroupId,omitempty" name:"ComplianceGroupId"`

	// 合规组名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 合规组描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 合规组规则配置（参数已废弃，请传空数组）
	ComplianceGroupRules []*ComplianceGroupRuleIdInfo `json:"ComplianceGroupRules,omitempty" name:"ComplianceGroupRules"`

	// 分级组ID，新增参数，可选参数，默认值为1
	LevelGroupId *uint64 `json:"LevelGroupId,omitempty" name:"LevelGroupId"`

	// 是否开启别名
	RuleAlias *bool `json:"RuleAlias,omitempty" name:"RuleAlias"`
}

func (r *ModifyDSPAComplianceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAComplianceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ComplianceGroupId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ComplianceGroupRules")
	delete(f, "LevelGroupId")
	delete(f, "RuleAlias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPAComplianceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAComplianceGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPAComplianceGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPAComplianceGroupResponseParams `json:"Response"`
}

func (r *ModifyDSPAComplianceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAComplianceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPADiscoveryRuleRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 规则名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// RDB类敏感数据识别规则
	RDBRules *ScanTaskRDBRules `json:"RDBRules,omitempty" name:"RDBRules"`

	// COS类敏感数据识别规则
	COSRules *ScanTaskCOSRules `json:"COSRules,omitempty" name:"COSRules"`
}

type ModifyDSPADiscoveryRuleRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 规则名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// RDB类敏感数据识别规则
	RDBRules *ScanTaskRDBRules `json:"RDBRules,omitempty" name:"RDBRules"`

	// COS类敏感数据识别规则
	COSRules *ScanTaskCOSRules `json:"COSRules,omitempty" name:"COSRules"`
}

func (r *ModifyDSPADiscoveryRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPADiscoveryRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "Name")
	delete(f, "RuleId")
	delete(f, "Description")
	delete(f, "RDBRules")
	delete(f, "COSRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPADiscoveryRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPADiscoveryRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPADiscoveryRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPADiscoveryRuleResponseParams `json:"Response"`
}

func (r *ModifyDSPADiscoveryRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPADiscoveryRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPADiscoveryTaskRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 任务开关，0 关闭，1 启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 用于传入的数据源的条件，目前只支持数据库，所以目前表示数据库的名称，最多添加5个数据库，之间通过逗号分隔
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 通用规则集开关；0 关闭，1 启用
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 合规组ID列表，最多支持添加5个
	ComplianceGroupIds []*int64 `json:"ComplianceGroupIds,omitempty" name:"ComplianceGroupIds"`

	// 执行计划； 0立即 1定时，选择“立即”时，扫描周期只能选择单次
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 扫描周期；0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 任务定时启动时间，格式：2006-01-02 15:04:05
	// 当执行计划（Plan字段）为”立即“时，定时启动时间不会生效，此场景下给该字段传值不会被保存。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 数据源类型，可取值如下：
	// cdb 表示云数据库 MySQL,
	// dcdb 表示TDSQL MySQL版,
	// mariadb 表示云数据库 MariaDB,
	// postgres 表示云数据库 PostgreSQL,
	// cynosdbpg 表示TDSQL-C PostgreSQL版,
	// cynosdbmysql 表示TDSQL-C MySQL版,
	// selfbuilt-db 表示自建数据库
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`
}

type ModifyDSPADiscoveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称，1-60个字符，仅允许输入中文、英文字母、数字、'_'、'-'，并且开头和结尾需为中文、英文字母或者数字，Name不可重复
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务描述，最大长度为1024个字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 任务开关，0 关闭，1 启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 数据源ID
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 用于传入的数据源的条件，目前只支持数据库，所以目前表示数据库的名称，最多添加5个数据库，之间通过逗号分隔
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 通用规则集开关；0 关闭，1 启用
	GeneralRuleSetEnable *int64 `json:"GeneralRuleSetEnable,omitempty" name:"GeneralRuleSetEnable"`

	// 合规组ID列表，最多支持添加5个
	ComplianceGroupIds []*int64 `json:"ComplianceGroupIds,omitempty" name:"ComplianceGroupIds"`

	// 执行计划； 0立即 1定时，选择“立即”时，扫描周期只能选择单次
	Plan *int64 `json:"Plan,omitempty" name:"Plan"`

	// 扫描周期；0单次 1每天 2每周 3每月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 任务定时启动时间，格式：2006-01-02 15:04:05
	// 当执行计划（Plan字段）为”立即“时，定时启动时间不会生效，此场景下给该字段传值不会被保存。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 数据源类型，可取值如下：
	// cdb 表示云数据库 MySQL,
	// dcdb 表示TDSQL MySQL版,
	// mariadb 表示云数据库 MariaDB,
	// postgres 表示云数据库 PostgreSQL,
	// cynosdbpg 表示TDSQL-C PostgreSQL版,
	// cynosdbmysql 表示TDSQL-C MySQL版,
	// selfbuilt-db 表示自建数据库
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`
}

func (r *ModifyDSPADiscoveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPADiscoveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Enable")
	delete(f, "DataSourceId")
	delete(f, "Condition")
	delete(f, "GeneralRuleSetEnable")
	delete(f, "ComplianceGroupIds")
	delete(f, "Plan")
	delete(f, "Period")
	delete(f, "TimingStartTime")
	delete(f, "ResourceRegion")
	delete(f, "DataSourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPADiscoveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPADiscoveryTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPADiscoveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPADiscoveryTaskResponseParams `json:"Response"`
}

func (r *ModifyDSPADiscoveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPADiscoveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAESTaskResultRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 字段扫描结果ID
	FieldResultId *int64 `json:"FieldResultId,omitempty" name:"FieldResultId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 是否设置为非敏感字段
	IsSetNonSensitiveField *bool `json:"IsSetNonSensitiveField,omitempty" name:"IsSetNonSensitiveField"`

	// 调整后新的规则ID
	DestRuleId *int64 `json:"DestRuleId,omitempty" name:"DestRuleId"`

	// 调整后新的分类ID
	DestCategoryId *int64 `json:"DestCategoryId,omitempty" name:"DestCategoryId"`

	// 调整后新的分级ID
	DestLevelId *int64 `json:"DestLevelId,omitempty" name:"DestLevelId"`
}

type ModifyDSPAESTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 字段扫描结果ID
	FieldResultId *int64 `json:"FieldResultId,omitempty" name:"FieldResultId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 是否设置为非敏感字段
	IsSetNonSensitiveField *bool `json:"IsSetNonSensitiveField,omitempty" name:"IsSetNonSensitiveField"`

	// 调整后新的规则ID
	DestRuleId *int64 `json:"DestRuleId,omitempty" name:"DestRuleId"`

	// 调整后新的分类ID
	DestCategoryId *int64 `json:"DestCategoryId,omitempty" name:"DestCategoryId"`

	// 调整后新的分级ID
	DestLevelId *int64 `json:"DestLevelId,omitempty" name:"DestLevelId"`
}

func (r *ModifyDSPAESTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAESTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "FieldResultId")
	delete(f, "ComplianceId")
	delete(f, "IsSetNonSensitiveField")
	delete(f, "DestRuleId")
	delete(f, "DestCategoryId")
	delete(f, "DestLevelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPAESTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPAESTaskResultResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPAESTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPAESTaskResultResponseParams `json:"Response"`
}

func (r *ModifyDSPAESTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPAESTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPATaskResultRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 字段扫描结果ID
	FieldResultId *int64 `json:"FieldResultId,omitempty" name:"FieldResultId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 是否设置为非敏感字段
	IsSetNonSensitiveField *bool `json:"IsSetNonSensitiveField,omitempty" name:"IsSetNonSensitiveField"`

	// 调整后新的规则ID
	DestRuleId *int64 `json:"DestRuleId,omitempty" name:"DestRuleId"`

	// 调整后新的分类ID
	DestCategoryId *int64 `json:"DestCategoryId,omitempty" name:"DestCategoryId"`

	// 调整后新的分级ID
	DestLevelId *int64 `json:"DestLevelId,omitempty" name:"DestLevelId"`
}

type ModifyDSPATaskResultRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 字段扫描结果ID
	FieldResultId *int64 `json:"FieldResultId,omitempty" name:"FieldResultId"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitempty" name:"ComplianceId"`

	// 是否设置为非敏感字段
	IsSetNonSensitiveField *bool `json:"IsSetNonSensitiveField,omitempty" name:"IsSetNonSensitiveField"`

	// 调整后新的规则ID
	DestRuleId *int64 `json:"DestRuleId,omitempty" name:"DestRuleId"`

	// 调整后新的分类ID
	DestCategoryId *int64 `json:"DestCategoryId,omitempty" name:"DestCategoryId"`

	// 调整后新的分级ID
	DestLevelId *int64 `json:"DestLevelId,omitempty" name:"DestLevelId"`
}

func (r *ModifyDSPATaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPATaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "FieldResultId")
	delete(f, "ComplianceId")
	delete(f, "IsSetNonSensitiveField")
	delete(f, "DestRuleId")
	delete(f, "DestCategoryId")
	delete(f, "DestLevelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDSPATaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDSPATaskResultResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDSPATaskResultResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDSPATaskResultResponseParams `json:"Response"`
}

func (r *ModifyDSPATaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDSPATaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLevelInfoRequestParams struct {

}

type ModifyLevelInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyLevelInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLevelInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLevelInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLevelInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLevelInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLevelInfoResponseParams `json:"Response"`
}

func (r *ModifyLevelInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLevelInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLevelNameRequestParams struct {

}

type ModifyLevelNameRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyLevelNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLevelNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLevelNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLevelNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLevelNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLevelNameResponseParams `json:"Response"`
}

func (r *ModifyLevelNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLevelNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLevelStateRequestParams struct {

}

type ModifyLevelStateRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyLevelStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLevelStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLevelStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLevelStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLevelStateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLevelStateResponseParams `json:"Response"`
}

func (r *ModifyLevelStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLevelStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMergeClassificationRequestParams struct {

}

type ModifyMergeClassificationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyMergeClassificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMergeClassificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMergeClassificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMergeClassificationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMergeClassificationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMergeClassificationResponseParams `json:"Response"`
}

func (r *ModifyMergeClassificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMergeClassificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNewClassificationRequestParams struct {

}

type ModifyNewClassificationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyNewClassificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewClassificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNewClassificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNewClassificationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNewClassificationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNewClassificationResponseParams `json:"Response"`
}

func (r *ModifyNewClassificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewClassificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStandardInfoRequestParams struct {

}

type ModifyStandardInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyStandardInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStandardInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStandardInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStandardInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyStandardInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStandardInfoResponseParams `json:"Response"`
}

func (r *ModifyStandardInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStandardInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MongoAsset struct {
	// DB总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbNums *int64 `json:"DbNums,omitempty" name:"DbNums"`

	// 敏感DB数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveDbNums *int64 `json:"SensitiveDbNums,omitempty" name:"SensitiveDbNums"`

	// 集合数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColNums *int64 `json:"ColNums,omitempty" name:"ColNums"`

	// 敏感集合的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveColNums *int64 `json:"SensitiveColNums,omitempty" name:"SensitiveColNums"`

	// 字段数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldNums *int64 `json:"FieldNums,omitempty" name:"FieldNums"`

	// 敏感的字段数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveFieldNums *int64 `json:"SensitiveFieldNums,omitempty" name:"SensitiveFieldNums"`
}

type MongoAssetDBDetail struct {
	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdName *string `json:"DdName,omitempty" name:"DdName"`

	// 数据库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *string `json:"DataType,omitempty" name:"DataType"`

	// 集合的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColNums *int64 `json:"ColNums,omitempty" name:"ColNums"`

	// 敏感集合数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveColNums *int64 `json:"SensitiveColNums,omitempty" name:"SensitiveColNums"`

	// 字段的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldNums *int64 `json:"FieldNums,omitempty" name:"FieldNums"`

	// 敏感字段的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveFieldNums *int64 `json:"SensitiveFieldNums,omitempty" name:"SensitiveFieldNums"`

	// 敏感数据分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributionData []*Note `json:"DistributionData,omitempty" name:"DistributionData"`
}

type NOSQLInstance struct {
	// 数据源id
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// cdb, dcdb, mariadb, postgres, cynosdbpg, cynosdbmysql, cos, mysql_like_proto, postgre_like_proto,mongodb
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 根据实例创建的敏感数据识别扫描任务Id
	DiscoveryTaskId *int64 `json:"DiscoveryTaskId,omitempty" name:"DiscoveryTaskId"`

	// 敏感数据识别任务实例id
	DiscoveryTaskInstanceID *int64 `json:"DiscoveryTaskInstanceID,omitempty" name:"DiscoveryTaskInstanceID"`
}

type Note struct {
	// 通用key，比如分类名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 通用value，比如分类个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type PrivilegeRisk struct {
	// 账户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountName []*string `json:"AccountName,omitempty" name:"AccountName"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ProcessHistory struct {
	// 处理时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 处理人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Handler *string `json:"Handler,omitempty" name:"Handler"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitempty" name:"Note"`
}

// Predefined struct for user
type QueryDSPAMetaResourceDbListRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据库实例ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 数据库实例所在地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 数据库实例类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`
}

type QueryDSPAMetaResourceDbListRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 数据库实例ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 数据库实例所在地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 数据库实例类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`
}

func (r *QueryDSPAMetaResourceDbListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDSPAMetaResourceDbListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ResourceId")
	delete(f, "ResourceRegion")
	delete(f, "MetaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryDSPAMetaResourceDbListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDSPAMetaResourceDbListResponseParams struct {
	// 数据库实例DB列表的查询结果。
	DbRelationStatusItems []*DbRelationStatusItem `json:"DbRelationStatusItems,omitempty" name:"DbRelationStatusItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryDSPAMetaResourceDbListResponse struct {
	*tchttp.BaseResponse
	Response *QueryDSPAMetaResourceDbListResponseParams `json:"Response"`
}

func (r *QueryDSPAMetaResourceDbListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDSPAMetaResourceDbListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryResourceDbBindStatusRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源所在地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`
}

type QueryResourceDbBindStatusRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源所在地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 资源类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`
}

func (r *QueryResourceDbBindStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResourceDbBindStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ResourceId")
	delete(f, "ResourceRegion")
	delete(f, "MetaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryResourceDbBindStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryResourceDbBindStatusResponseParams struct {
	// 绑定DB数量。
	BindDbNums *uint64 `json:"BindDbNums,omitempty" name:"BindDbNums"`

	// 未绑定DB数量。
	UnbindDbNums *uint64 `json:"UnbindDbNums,omitempty" name:"UnbindDbNums"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryResourceDbBindStatusResponse struct {
	*tchttp.BaseResponse
	Response *QueryResourceDbBindStatusResponseParams `json:"Response"`
}

func (r *QueryResourceDbBindStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResourceDbBindStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RDBAsset struct {
	// DB总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbNums *int64 `json:"DbNums,omitempty" name:"DbNums"`

	// 敏感DB数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveDbNums *int64 `json:"SensitiveDbNums,omitempty" name:"SensitiveDbNums"`

	// 表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNums *int64 `json:"TableNums,omitempty" name:"TableNums"`

	// 敏感表的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveTableNums *int64 `json:"SensitiveTableNums,omitempty" name:"SensitiveTableNums"`

	// 字段数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldNums *int64 `json:"FieldNums,omitempty" name:"FieldNums"`

	// 敏感的字段数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveFieldNums *int64 `json:"SensitiveFieldNums,omitempty" name:"SensitiveFieldNums"`
}

type RDBInstance struct {
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// cdb, dcdb, mariadb, postgres, cynosdbpg, cynosdbmysql, cos, mysql_like_proto, postgre_like_proto
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 资源所在地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 若未来扩展到DBName粒度，可采用
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBs []*DBStatements `json:"DBs,omitempty" name:"DBs"`
}

type ReportInfo struct {
	// 任务id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 报告名称
	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`

	// 报告类型（AssetSorting:资产梳理）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportType *string `json:"ReportType,omitempty" name:"ReportType"`

	// 报告周期（0单次 1每天 2每周 3每月）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportPeriod *uint64 `json:"ReportPeriod,omitempty" name:"ReportPeriod"`

	// 执行计划 （0:单次报告 1:定时报告）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportPlan *uint64 `json:"ReportPlan,omitempty" name:"ReportPlan"`

	// 报告导出状态（Success 成功, Failed 失败, InProgress 进行中）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportStatus *string `json:"ReportStatus,omitempty" name:"ReportStatus"`

	// 任务下次启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimingStartTime *string `json:"TimingStartTime,omitempty" name:"TimingStartTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedTime *string `json:"FinishedTime,omitempty" name:"FinishedTime"`

	// 子账号uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`

	// 失败信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedMessage *string `json:"FailedMessage,omitempty" name:"FailedMessage"`

	// 是否启用（0：否 1：是）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// 识别模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceName *string `json:"ComplianceName,omitempty" name:"ComplianceName"`

	// 进度百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgressPercent *uint64 `json:"ProgressPercent,omitempty" name:"ProgressPercent"`
}

// Predefined struct for user
type RestartDSPAAssessmentTaskRequestParams struct {
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估任务Id，格式“task-xxxxxxxx”
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type RestartDSPAAssessmentTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例Id，格式“dspa-xxxxxxxx”
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 评估任务Id，格式“task-xxxxxxxx”
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RestartDSPAAssessmentTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDSPAAssessmentTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDSPAAssessmentTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDSPAAssessmentTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestartDSPAAssessmentTaskResponse struct {
	*tchttp.BaseResponse
	Response *RestartDSPAAssessmentTaskResponseParams `json:"Response"`
}

func (r *RestartDSPAAssessmentTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDSPAAssessmentTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskCountInfo struct {
	// 风险等级
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 该等级风险项数量
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type RiskDealedTrendItem struct {
	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 未解决数量
	Unhandled *uint64 `json:"Unhandled,omitempty" name:"Unhandled"`

	// 已解决数量
	Handled *uint64 `json:"Handled,omitempty" name:"Handled"`

	// 新发现
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewDiscoveryHandled *uint64 `json:"NewDiscoveryHandled,omitempty" name:"NewDiscoveryHandled"`
}

type RiskItem struct {
	// 名称
	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`

	// 风险数量
	RiskNum *uint64 `json:"RiskNum,omitempty" name:"RiskNum"`
}

type RiskItemInfo struct {
	// 最新风险项id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceName *string `json:"DataSourceName,omitempty" name:"DataSourceName"`

	// 数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 资产名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 风险类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`

	// 风险项
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskName *string `json:"RiskName,omitempty" name:"RiskName"`

	// 风险级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 风险描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskDescription *string `json:"RiskDescription,omitempty" name:"RiskDescription"`

	// 建议措施
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuggestAction *string `json:"SuggestAction,omitempty" name:"SuggestAction"`

	// 安全产品（可能有多个）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityProduct []*SecurityProduct `json:"SecurityProduct,omitempty" name:"SecurityProduct"`

	// 状态(waiting:待处理，processing:处理中，finished:已处理，ignored:已忽略)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 最后处置时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastProcessTime *string `json:"LastProcessTime,omitempty" name:"LastProcessTime"`

	// 分类分级合规组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentifyComplianceId *int64 `json:"IdentifyComplianceId,omitempty" name:"IdentifyComplianceId"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSubType *string `json:"ItemSubType,omitempty" name:"ItemSubType"`
}

type RiskLevelMatrix struct {
	// 存储id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 分类分级levelID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveLevelId *int64 `json:"SensitiveLevelId,omitempty" name:"SensitiveLevelId"`

	// 分类分级名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveLevelName *string `json:"SensitiveLevelName,omitempty" name:"SensitiveLevelName"`

	// 漏洞级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityLevel *string `json:"VulnerabilityLevel,omitempty" name:"VulnerabilityLevel"`

	// 风险级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
}

type RiskLevelRisk struct {
	// 风险id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 风险等级列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevelName *string `json:"RiskLevelName,omitempty" name:"RiskLevelName"`

	// 风险级别描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevelDescription *string `json:"RiskLevelDescription,omitempty" name:"RiskLevelDescription"`

	// 引用的分类分级模版
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentifyComplianceName *string `json:"IdentifyComplianceName,omitempty" name:"IdentifyComplianceName"`

	// 类型，区分自定义还是系统内置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type RiskLevelTrendItem struct {
	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 高风险数量
	High *uint64 `json:"High,omitempty" name:"High"`

	// 中风险数量
	Medium *uint64 `json:"Medium,omitempty" name:"Medium"`

	// 低风险数量
	Low *uint64 `json:"Low,omitempty" name:"Low"`

	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type RiskMatrixLevel struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 就是id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitempty" name:"Score"`
}

type RuleDistribution struct {
	// 规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 分级id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 分级名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelName *string `json:"LevelName,omitempty" name:"LevelName"`

	// 规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleCnt *int64 `json:"RuleCnt,omitempty" name:"RuleCnt"`
}

type ScanTaskCOSRules struct {
	// regex规则内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegexRule *COSDataRule `json:"RegexRule,omitempty" name:"RegexRule"`

	// 规则状态；0 不启用, 1 启用
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 关键词规则内容组，最大支持5个关键词。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeywordRule *COSDataRule `json:"KeywordRule,omitempty" name:"KeywordRule"`

	// 忽略词规则内容组，最大支持5个忽略词。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreStringRule *COSDataRule `json:"IgnoreStringRule,omitempty" name:"IgnoreStringRule"`

	// 最大匹配距离，默认值为100。上限为500.
	MaxMatch *int64 `json:"MaxMatch,omitempty" name:"MaxMatch"`
}

type ScanTaskComplianceInfo struct {
	// 合规组ID
	ComplianceGroupId *int64 `json:"ComplianceGroupId,omitempty" name:"ComplianceGroupId"`

	// 合规组名称
	ComplianceGroupName *string `json:"ComplianceGroupName,omitempty" name:"ComplianceGroupName"`
}

type ScanTaskRDBRules struct {
	// 规则状态；0 不启用, 1 启用
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 只能取and 、or两个值其中之一，and：字段和内容同时满足，or：字段和内容满足其一
	MatchOperator *string `json:"MatchOperator,omitempty" name:"MatchOperator"`

	// 字段名包含规则，最大支持选择9项
	MetaRule *DataRules `json:"MetaRule,omitempty" name:"MetaRule"`

	// 内容包含规则，最大支持选择9项
	ContentRule *DataRules `json:"ContentRule,omitempty" name:"ContentRule"`
}

type ScanTaskResult struct {
	// 任务最新一次运行结果ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 任务扫描结束的时间，格式如：2021-12-12 12:12:12
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务状态，-1待触发 0待扫描 1扫描中 2扫描终止 3扫描成功 4扫描失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 扫描任务结果展示，如果扫描失败，则显示失败原因
	Result *string `json:"Result,omitempty" name:"Result"`
}

type SecurityProduct struct {
	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReferUrl *string `json:"ReferUrl,omitempty" name:"ReferUrl"`
}

type SensitiveLevel struct {
	// 分级标识ID
	LevelId *int64 `json:"LevelId,omitempty" name:"LevelId"`

	// 分级标识统计
	LevelCnt *int64 `json:"LevelCnt,omitempty" name:"LevelCnt"`

	// 分级标识名称
	LevelRiskName *string `json:"LevelRiskName,omitempty" name:"LevelRiskName"`

	// 分级标识分数
	LevelRiskScore *int64 `json:"LevelRiskScore,omitempty" name:"LevelRiskScore"`
}

// Predefined struct for user
type StartDSPADiscoveryTaskRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type StartDSPADiscoveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StartDSPADiscoveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartDSPADiscoveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartDSPADiscoveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartDSPADiscoveryTaskResponseParams struct {
	// 任务扫描结果ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultId *int64 `json:"ResultId,omitempty" name:"ResultId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartDSPADiscoveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *StartDSPADiscoveryTaskResponseParams `json:"Response"`
}

func (r *StartDSPADiscoveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartDSPADiscoveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDSPADiscoveryTaskRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

type StopDSPADiscoveryTaskRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopDSPADiscoveryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDSPADiscoveryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopDSPADiscoveryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDSPADiscoveryTaskResponseParams struct {
	// 任务扫描结果ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultId *int64 `json:"ResultId,omitempty" name:"ResultId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopDSPADiscoveryTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopDSPADiscoveryTaskResponseParams `json:"Response"`
}

func (r *StopDSPADiscoveryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDSPADiscoveryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SuggestRiskLevelMatrix struct {
	// 矩阵
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevelMatrix []*SuggestRiskLevelMatrixItem `json:"RiskLevelMatrix,omitempty" name:"RiskLevelMatrix"`
}

type SuggestRiskLevelMatrixItem struct {
	// 分类分级等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveLevel *RiskMatrixLevel `json:"SensitiveLevel,omitempty" name:"SensitiveLevel"`

	// 脆弱项等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityLevel *RiskMatrixLevel `json:"VulnerabilityLevel,omitempty" name:"VulnerabilityLevel"`

	// 风险名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskName *string `json:"RiskName,omitempty" name:"RiskName"`

	// 分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskScore *float64 `json:"RiskScore,omitempty" name:"RiskScore"`
}

type TopAsset struct {
	// 分级名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelName *string `json:"LevelName,omitempty" name:"LevelName"`

	// top数据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopStat []*TopAssetStat `json:"TopStat,omitempty" name:"TopStat"`
}

type TopAssetStat struct {
	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceId *string `json:"DataSourceId,omitempty" name:"DataSourceId"`

	// db_name
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubData *string `json:"SubData,omitempty" name:"SubData"`

	// 敏感个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveCnt *int64 `json:"SensitiveCnt,omitempty" name:"SensitiveCnt"`
}

// Predefined struct for user
type UpdateDSPASelfBuildResourceRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 云资源名称，如果是通过CVM访问则填写CVM的资源ID，如果是通过LB访问则填写LB的资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源绑定的端口，为0则表示不更新。
	ResourceVPort *uint64 `json:"ResourceVPort,omitempty" name:"ResourceVPort"`

	// 账户名，为空则表示不更新。
	// UserName和Password必须同时填写或同时为空。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 账户密码，为空则表示不更新。
	// UserName和Password必须同时填写或同时为空。
	Password *string `json:"Password,omitempty" name:"Password"`
}

type UpdateDSPASelfBuildResourceRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 云资源名称，如果是通过CVM访问则填写CVM的资源ID，如果是通过LB访问则填写LB的资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源绑定的端口，为0则表示不更新。
	ResourceVPort *uint64 `json:"ResourceVPort,omitempty" name:"ResourceVPort"`

	// 账户名，为空则表示不更新。
	// UserName和Password必须同时填写或同时为空。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 账户密码，为空则表示不更新。
	// UserName和Password必须同时填写或同时为空。
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *UpdateDSPASelfBuildResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDSPASelfBuildResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "ResourceId")
	delete(f, "ResourceVPort")
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDSPASelfBuildResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDSPASelfBuildResourceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDSPASelfBuildResourceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDSPASelfBuildResourceResponseParams `json:"Response"`
}

func (r *UpdateDSPASelfBuildResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDSPASelfBuildResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyDSPACOSRuleRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 待验证COS规则
	COSRules *ScanTaskCOSRules `json:"COSRules,omitempty" name:"COSRules"`

	// 待验证数据
	Data *string `json:"Data,omitempty" name:"Data"`
}

type VerifyDSPACOSRuleRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 待验证COS规则
	COSRules *ScanTaskCOSRules `json:"COSRules,omitempty" name:"COSRules"`

	// 待验证数据
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *VerifyDSPACOSRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyDSPACOSRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "COSRules")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyDSPACOSRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyDSPACOSRuleResponseParams struct {
	// 验证结果
	// Success 验证成功
	// Failed 验证失败
	VerifyResult *string `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 验证结果详情
	DetailInfo *string `json:"DetailInfo,omitempty" name:"DetailInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyDSPACOSRuleResponse struct {
	*tchttp.BaseResponse
	Response *VerifyDSPACOSRuleResponseParams `json:"Response"`
}

func (r *VerifyDSPACOSRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyDSPACOSRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyDSPADiscoveryRuleRequestParams struct {
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 只能取and 、or两个值其中之一，and：字段和内容同时满足，or：字段和内容满足其一
	MatchOperator *string `json:"MatchOperator,omitempty" name:"MatchOperator"`

	// 字段名包含规则，最大支持选择9项
	MetaRule *DataRules `json:"MetaRule,omitempty" name:"MetaRule"`

	// 内容包含规则，最大支持选择9项
	ContentRule *DataRules `json:"ContentRule,omitempty" name:"ContentRule"`

	// 验证规则字段名，最大长度为1024个字符
	VerifyMeta *string `json:"VerifyMeta,omitempty" name:"VerifyMeta"`

	// 验证规则数据内容，最大长度为1024个字符
	VerifyContent *string `json:"VerifyContent,omitempty" name:"VerifyContent"`
}

type VerifyDSPADiscoveryRuleRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 只能取and 、or两个值其中之一，and：字段和内容同时满足，or：字段和内容满足其一
	MatchOperator *string `json:"MatchOperator,omitempty" name:"MatchOperator"`

	// 字段名包含规则，最大支持选择9项
	MetaRule *DataRules `json:"MetaRule,omitempty" name:"MetaRule"`

	// 内容包含规则，最大支持选择9项
	ContentRule *DataRules `json:"ContentRule,omitempty" name:"ContentRule"`

	// 验证规则字段名，最大长度为1024个字符
	VerifyMeta *string `json:"VerifyMeta,omitempty" name:"VerifyMeta"`

	// 验证规则数据内容，最大长度为1024个字符
	VerifyContent *string `json:"VerifyContent,omitempty" name:"VerifyContent"`
}

func (r *VerifyDSPADiscoveryRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyDSPADiscoveryRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "MatchOperator")
	delete(f, "MetaRule")
	delete(f, "ContentRule")
	delete(f, "VerifyMeta")
	delete(f, "VerifyContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyDSPADiscoveryRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyDSPADiscoveryRuleResponseParams struct {
	// 验证结果
	// Success 验证成功
	// Failed 验证失败
	VerifyResult *string `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 验证结果详情
	DetailInfo *string `json:"DetailInfo,omitempty" name:"DetailInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyDSPADiscoveryRuleResponse struct {
	*tchttp.BaseResponse
	Response *VerifyDSPADiscoveryRuleResponseParams `json:"Response"`
}

func (r *VerifyDSPADiscoveryRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyDSPADiscoveryRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
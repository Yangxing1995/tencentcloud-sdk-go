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

package v20230321

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddClusterStorageOptionRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 集群存储选项。
	StorageOption *StorageOption `json:"StorageOption,omitnil" name:"StorageOption"`
}

type AddClusterStorageOptionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 集群存储选项。
	StorageOption *StorageOption `json:"StorageOption,omitnil" name:"StorageOption"`
}

func (r *AddClusterStorageOptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterStorageOptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "StorageOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddClusterStorageOptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterStorageOptionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddClusterStorageOptionResponse struct {
	*tchttp.BaseResponse
	Response *AddClusterStorageOptionResponseParams `json:"Response"`
}

func (r *AddClusterStorageOptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterStorageOptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodesRequestParams struct {
	// 集群中实例所在的位置。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 私有网络相关信息配置。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil" name:"VirtualPrivateCloud"`

	// 添加节点数量。
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。目前仅支持公有镜像和特定自定义镜像。
	ImageId *string `json:"ImageId,omitnil" name:"ImageId"`

	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>SPOTPAID：竞价付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。<br><li>具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil" name:"InternetAccessible"`

	// 节点显示名称。
	// 不指定节点显示名称则默认显示‘未命名’。
	// 最多支持60个字符。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 集群登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// 集群中实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 队列名称。不指定则为默认队列。<li>SLURM默认队列为：compute。<li>SGE默认队列为：all.q。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// 添加节点角色。默认值：Compute<br><li>Compute：计算节点。<br><li>Login：登录节点。
	NodeRole *string `json:"NodeRole,omitnil" name:"NodeRole"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`

	// 添加节点类型。默认取值：STATIC。<li>STATIC：静态节点，不会参与弹性伸缩流程。<li>DYNAMIC：弹性节点，会被弹性缩容的节点。管控节点和登录节点不支持此参数。
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type AddNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群中实例所在的位置。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 私有网络相关信息配置。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil" name:"VirtualPrivateCloud"`

	// 添加节点数量。
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。目前仅支持公有镜像和特定自定义镜像。
	ImageId *string `json:"ImageId,omitnil" name:"ImageId"`

	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>SPOTPAID：竞价付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。<br><li>具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil" name:"InternetAccessible"`

	// 节点显示名称。
	// 不指定节点显示名称则默认显示‘未命名’。
	// 最多支持60个字符。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 集群登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// 集群中实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 队列名称。不指定则为默认队列。<li>SLURM默认队列为：compute。<li>SGE默认队列为：all.q。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// 添加节点角色。默认值：Compute<br><li>Compute：计算节点。<br><li>Login：登录节点。
	NodeRole *string `json:"NodeRole,omitnil" name:"NodeRole"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`

	// 添加节点类型。默认取值：STATIC。<li>STATIC：静态节点，不会参与弹性伸缩流程。<li>DYNAMIC：弹性节点，会被弹性缩容的节点。管控节点和登录节点不支持此参数。
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *AddNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "ClusterId")
	delete(f, "VirtualPrivateCloud")
	delete(f, "Count")
	delete(f, "ImageId")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "InstanceType")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "InternetAccessible")
	delete(f, "InstanceName")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "ClientToken")
	delete(f, "QueueName")
	delete(f, "NodeRole")
	delete(f, "DryRun")
	delete(f, "NodeType")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddNodesResponse struct {
	*tchttp.BaseResponse
	Response *AddNodesResponseParams `json:"Response"`
}

func (r *AddNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddQueueRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 队列名称。<br><li>最多支持32个字符。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
}

type AddQueueRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 队列名称。<br><li>最多支持32个字符。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
}

func (r *AddQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddQueueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddQueueResponse struct {
	*tchttp.BaseResponse
	Response *AddQueueResponseParams `json:"Response"`
}

func (r *AddQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CFSOption struct {
	// 文件系统本地挂载路径。
	LocalPath *string `json:"LocalPath,omitnil" name:"LocalPath"`

	// 文件系统远程挂载ip及路径。
	RemotePath *string `json:"RemotePath,omitnil" name:"RemotePath"`

	// 文件系统协议类型，默认值NFS 3.0。
	// <li>NFS 3.0。
	// <li>NFS 4.0。
	// <li>TURBO。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 文件系统存储类型，默认值SD；其中 SD 为通用标准型标准型存储， HP为通用性能型存储， TB为turbo标准型， TP 为turbo性能型。
	StorageType *string `json:"StorageType,omitnil" name:"StorageType"`

	// 文件系统挂载挂载命令参数选项。
	// 
	// - NFS 3.0默认值：vers=3,nolock,proto=tcp,noresvport
	// - NFS 4.0默认值：vers=4.0,noresvport
	// - TURBO默认值：user_xattr
	MountOption *string `json:"MountOption,omitnil" name:"MountOption"`
}

type CFSOptionOverview struct {
	// 文件系统本地挂载路径。
	LocalPath *string `json:"LocalPath,omitnil" name:"LocalPath"`

	// 文件系统远程挂载ip及路径。
	RemotePath *string `json:"RemotePath,omitnil" name:"RemotePath"`

	// 文件系统协议类型。
	// <li>NFS 3.0。
	// <li>NFS 4.0。
	// <li>TURBO。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 文件系统存储类型，默认值SD；其中 SD 为通用标准型标准型存储， HP为通用性能型存储， TB为turbo标准型， TP 为turbo性能型。
	StorageType *string `json:"StorageType,omitnil" name:"StorageType"`

	// 文件系统挂载命令参数选项。
	MountOption *string `json:"MountOption,omitnil" name:"MountOption"`
}

type ClusterActivity struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 集群活动ID。
	ActivityId *string `json:"ActivityId,omitnil" name:"ActivityId"`

	// 集群活动类型。取值范围：<br><li>CreateAndAddNodes：创建实例并添加进集群<br><li>RemoveNodesFromCluster：从集群移除实例<br><li>TerminateNodes：销毁实例<br><li>MountStorageOption：增加挂载选项并进行挂载<br><li>UmountStorageOption：删除集群挂载存储选项并解挂载
	ActivityType *string `json:"ActivityType,omitnil" name:"ActivityType"`

	// 集群活动状态。取值范围：<br><li>PENDING：等待运行<br><li>RUNNING：运行中<br><li>SUCCESSFUL：活动成功<br><li>PARTIALLY_SUCCESSFUL：活动部分成功<br><li>FAILED：活动失败
	ActivityStatus *string `json:"ActivityStatus,omitnil" name:"ActivityStatus"`

	// 集群活动状态码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityStatusCode *string `json:"ActivityStatusCode,omitnil" name:"ActivityStatusCode"`

	// 集群活动结果详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultDetail *string `json:"ResultDetail,omitnil" name:"ResultDetail"`

	// 集群活动起因。
	Cause *string `json:"Cause,omitnil" name:"Cause"`

	// 集群活动描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 集群活动相关节点活动集合。
	RelatedNodeActivitySet []*NodeActivity `json:"RelatedNodeActivitySet,omitnil" name:"RelatedNodeActivitySet"`

	// 集群活动开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 集群活动结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type ClusterOverview struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 集群状态。取值范围：<br><li>PENDING：创建中<br><li>INITING：初始化中<br><li>INIT_FAILED：初始化失败<br><li>RUNNING：运行中<br><li>TERMINATING：销毁中
	ClusterStatus *string `json:"ClusterStatus,omitnil" name:"ClusterStatus"`

	// 集群名称。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 集群位置信息。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 集群创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 集群调度器。
	SchedulerType *string `json:"SchedulerType,omitnil" name:"SchedulerType"`

	// 计算节点数量。
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitnil" name:"ComputeNodeCount"`

	// 计算节点概览。
	ComputeNodeSet []*ComputeNodeOverview `json:"ComputeNodeSet,omitnil" name:"ComputeNodeSet"`

	// 管控节点数量。
	ManagerNodeCount *int64 `json:"ManagerNodeCount,omitnil" name:"ManagerNodeCount"`

	// 管控节点概览。
	ManagerNodeSet []*ManagerNodeOverview `json:"ManagerNodeSet,omitnil" name:"ManagerNodeSet"`

	// 登录节点概览。
	LoginNodeSet []*LoginNodeOverview `json:"LoginNodeSet,omitnil" name:"LoginNodeSet"`

	// 登录节点数量。
	LoginNodeCount *int64 `json:"LoginNodeCount,omitnil" name:"LoginNodeCount"`

	// 弹性伸缩类型。<br><li>THPC_AS：集群自动扩缩容由THPC产品内部实现。<br><li>AS：集群自动扩缩容由[弹性伸缩](https://cloud.tencent.com/document/product/377/3154)产品实现。
	AutoScalingType *string `json:"AutoScalingType,omitnil" name:"AutoScalingType"`

	// 集群所属私有网络ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

type ComputeNode struct {
	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>SPOTPAID：竞价付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。
	// <br><li>具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil" name:"InternetAccessible"`

	// 节点显示名称。<br><li>
	// 不指定节点显示名称则默认显示‘未命名’。
	// 最多支持60个字符。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type ComputeNodeOverview struct {
	// 计算节点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// 集群中实例所在的位置。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 指定管理节点。
	ManagerNode *ManagerNode `json:"ManagerNode,omitnil" name:"ManagerNode"`

	// 指定管理节点的数量。默认取值：1。取值范围：1～2。
	ManagerNodeCount *int64 `json:"ManagerNodeCount,omitnil" name:"ManagerNodeCount"`

	// 指定计算节点。
	ComputeNode *ComputeNode `json:"ComputeNode,omitnil" name:"ComputeNode"`

	// 指定计算节点的数量。默认取值：0。
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitnil" name:"ComputeNodeCount"`

	// 调度器类型。默认取值：SLURM。<br><li>SGE：SGE调度器。<br><li>SLURM：SLURM调度器。
	SchedulerType *string `json:"SchedulerType,omitnil" name:"SchedulerType"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。目前支持部分公有镜像和自定义镜像。
	ImageId *string `json:"ImageId,omitnil" name:"ImageId"`

	// 私有网络相关信息配置。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil" name:"VirtualPrivateCloud"`

	// 集群登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// 集群中实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`

	// 域名字服务类型。默认取值：NIS。
	// <li>NIS：NIS域名字服务。
	AccountType *string `json:"AccountType,omitnil" name:"AccountType"`

	// 集群显示名称。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 集群存储选项
	StorageOption *StorageOption `json:"StorageOption,omitnil" name:"StorageOption"`

	// 指定登录节点。
	LoginNode *LoginNode `json:"LoginNode,omitnil" name:"LoginNode"`

	// 指定登录节点的数量。默认取值：0。取值范围：0～10。
	LoginNodeCount *int64 `json:"LoginNodeCount,omitnil" name:"LoginNodeCount"`

	// 创建集群时同时绑定的标签对说明。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 弹性伸缩类型。默认值：THPC_AS<br><li>THPC_AS：集群自动扩缩容由THPC产品内部实现。<br><li>AS：集群自动扩缩容由[弹性伸缩](https://cloud.tencent.com/document/product/377/3154)产品实现。
	AutoScalingType *string `json:"AutoScalingType,omitnil" name:"AutoScalingType"`

	// 节点初始化脚本信息列表。
	InitNodeScripts []*NodeScript `json:"InitNodeScripts,omitnil" name:"InitNodeScripts"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群中实例所在的位置。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 指定管理节点。
	ManagerNode *ManagerNode `json:"ManagerNode,omitnil" name:"ManagerNode"`

	// 指定管理节点的数量。默认取值：1。取值范围：1～2。
	ManagerNodeCount *int64 `json:"ManagerNodeCount,omitnil" name:"ManagerNodeCount"`

	// 指定计算节点。
	ComputeNode *ComputeNode `json:"ComputeNode,omitnil" name:"ComputeNode"`

	// 指定计算节点的数量。默认取值：0。
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitnil" name:"ComputeNodeCount"`

	// 调度器类型。默认取值：SLURM。<br><li>SGE：SGE调度器。<br><li>SLURM：SLURM调度器。
	SchedulerType *string `json:"SchedulerType,omitnil" name:"SchedulerType"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。目前支持部分公有镜像和自定义镜像。
	ImageId *string `json:"ImageId,omitnil" name:"ImageId"`

	// 私有网络相关信息配置。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil" name:"VirtualPrivateCloud"`

	// 集群登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// 集群中实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`

	// 域名字服务类型。默认取值：NIS。
	// <li>NIS：NIS域名字服务。
	AccountType *string `json:"AccountType,omitnil" name:"AccountType"`

	// 集群显示名称。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 集群存储选项
	StorageOption *StorageOption `json:"StorageOption,omitnil" name:"StorageOption"`

	// 指定登录节点。
	LoginNode *LoginNode `json:"LoginNode,omitnil" name:"LoginNode"`

	// 指定登录节点的数量。默认取值：0。取值范围：0～10。
	LoginNodeCount *int64 `json:"LoginNodeCount,omitnil" name:"LoginNodeCount"`

	// 创建集群时同时绑定的标签对说明。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 弹性伸缩类型。默认值：THPC_AS<br><li>THPC_AS：集群自动扩缩容由THPC产品内部实现。<br><li>AS：集群自动扩缩容由[弹性伸缩](https://cloud.tencent.com/document/product/377/3154)产品实现。
	AutoScalingType *string `json:"AutoScalingType,omitnil" name:"AutoScalingType"`

	// 节点初始化脚本信息列表。
	InitNodeScripts []*NodeScript `json:"InitNodeScripts,omitnil" name:"InitNodeScripts"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "ManagerNode")
	delete(f, "ManagerNodeCount")
	delete(f, "ComputeNode")
	delete(f, "ComputeNodeCount")
	delete(f, "SchedulerType")
	delete(f, "ImageId")
	delete(f, "VirtualPrivateCloud")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	delete(f, "AccountType")
	delete(f, "ClusterName")
	delete(f, "StorageOption")
	delete(f, "LoginNode")
	delete(f, "LoginNodeCount")
	delete(f, "Tags")
	delete(f, "AutoScalingType")
	delete(f, "InitNodeScripts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// 集群ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同，具体限制详见：[存储概述](https://cloud.tencent.com/document/product/213/4952)。默认值为0，表示不购买数据盘。更多限制详见产品文档。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 数据盘类型。数据盘类型限制详见[存储概述](https://cloud.tencent.com/document/product/213/4952)。取值范围：<br><li>LOCAL_NVME：本地NVME硬盘，与InstanceType强相关，不支持指定<br><li>LOCAL_PRO：本地HDD硬盘，与InstanceType强相关，不支持指定<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><li>CLOUD_HSSD：增强型SSD云硬盘<br><li>CLOUD_TSSD：极速型SSD云硬盘<br><li>CLOUD_BSSD：通用型SSD云硬盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`
}

// Predefined struct for user
type DeleteClusterRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterResponseParams `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterStorageOptionRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 本地挂载路径。
	LocalPath *string `json:"LocalPath,omitnil" name:"LocalPath"`
}

type DeleteClusterStorageOptionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 本地挂载路径。
	LocalPath *string `json:"LocalPath,omitnil" name:"LocalPath"`
}

func (r *DeleteClusterStorageOptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterStorageOptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "LocalPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterStorageOptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterStorageOptionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteClusterStorageOptionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterStorageOptionResponseParams `json:"Response"`
}

func (r *DeleteClusterStorageOptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterStorageOptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNodesRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 节点ID。
	NodeIds []*string `json:"NodeIds,omitnil" name:"NodeIds"`
}

type DeleteNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 节点ID。
	NodeIds []*string `json:"NodeIds,omitnil" name:"NodeIds"`
}

func (r *DeleteNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNodesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNodesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNodesResponseParams `json:"Response"`
}

func (r *DeleteNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQueueRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 队列名称。<br><li>最多支持32个字符。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
}

type DeleteQueueRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 队列名称。<br><li>最多支持32个字符。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
}

func (r *DeleteQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQueueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteQueueResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQueueResponseParams `json:"Response"`
}

func (r *DeleteQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingConfigurationRequestParams struct {
	// 集群ID。	
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribeAutoScalingConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。	
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *DescribeAutoScalingConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingConfigurationResponseParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 任务连续等待时间，队列的任务处于连续等待的时间。单位秒。
	ExpansionBusyTime *int64 `json:"ExpansionBusyTime,omitnil" name:"ExpansionBusyTime"`

	// 节点连续空闲（未运行作业）时间，一个节点连续处于空闲状态时间。
	ShrinkIdleTime *int64 `json:"ShrinkIdleTime,omitnil" name:"ShrinkIdleTime"`

	// 扩容队列配置概览列表。
	QueueConfigs []*QueueConfigOverview `json:"QueueConfigs,omitnil" name:"QueueConfigs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAutoScalingConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingConfigurationResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterActivitiesRequestParams struct {
	// 集群ID。通过该参数指定需要查询活动历史记录的集群。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeClusterActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。通过该参数指定需要查询活动历史记录的集群。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeClusterActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterActivitiesResponseParams struct {
	// 集群活动历史记录列表。
	ClusterActivitySet []*ClusterActivity `json:"ClusterActivitySet,omitnil" name:"ClusterActivitySet"`

	// 集群活动历史记录数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClusterActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterActivitiesResponseParams `json:"Response"`
}

func (r *DescribeClusterActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterStorageOptionRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribeClusterStorageOptionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *DescribeClusterStorageOptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterStorageOptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterStorageOptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterStorageOptionResponseParams struct {
	// 集群存储选项信息概览。
	StorageOption *StorageOptionOverview `json:"StorageOption,omitnil" name:"StorageOption"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClusterStorageOptionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterStorageOptionResponseParams `json:"Response"`
}

func (r *DescribeClusterStorageOptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterStorageOptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// 集群ID列表。通过该参数可以指定需要查询信息的集群列表。<br>如果您不指定该参数，则返回Limit数量以内的集群信息。
	ClusterIds []*string `json:"ClusterIds,omitnil" name:"ClusterIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表。通过该参数可以指定需要查询信息的集群列表。<br>如果您不指定该参数，则返回Limit数量以内的集群信息。
	ClusterIds []*string `json:"ClusterIds,omitnil" name:"ClusterIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// 集群概览信息列表。
	ClusterSet []*ClusterOverview `json:"ClusterSet,omitnil" name:"ClusterSet"`

	// 集群数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInitNodeScriptsRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribeInitNodeScriptsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *DescribeInitNodeScriptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInitNodeScriptsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInitNodeScriptsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInitNodeScriptsResponseParams struct {
	// 节点初始化脚本列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitNodeScriptSet []*NodeScript `json:"InitNodeScriptSet,omitnil" name:"InitNodeScriptSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInitNodeScriptsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInitNodeScriptsResponseParams `json:"Response"`
}

func (r *DescribeInitNodeScriptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInitNodeScriptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodesRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// <li><strong>queue-name</strong></li> <p style="padding-left: 30px;">按照【<strong>队列名称</strong>】进行过滤。队列名称形如：compute。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"><li><strong>node-role</strong></li> <p style="padding-left: 30px;">按照【<strong>节点角色</strong>】进行过滤。节点角色形如：Manager。（Manager：管控节点。Compute：计算节点。Login：登录节点。ManagerBackup：备用管控节点。）</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"><li><strong>node-type</strong></li> <p style="padding-left: 30px;">按照【<strong>节点类型</strong>】进行过滤。节点类型形如：STATIC。(STATIC：静态节点。DYNAMIC：弹性节点。)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// <li><strong>queue-name</strong></li> <p style="padding-left: 30px;">按照【<strong>队列名称</strong>】进行过滤。队列名称形如：compute。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"><li><strong>node-role</strong></li> <p style="padding-left: 30px;">按照【<strong>节点角色</strong>】进行过滤。节点角色形如：Manager。（Manager：管控节点。Compute：计算节点。Login：登录节点。ManagerBackup：备用管控节点。）</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"><li><strong>node-type</strong></li> <p style="padding-left: 30px;">按照【<strong>节点类型</strong>】进行过滤。节点类型形如：STATIC。(STATIC：静态节点。DYNAMIC：弹性节点。)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodesResponseParams struct {
	// 节点概览信息列表。
	NodeSet []*NodeOverview `json:"NodeSet,omitnil" name:"NodeSet"`

	// 符合条件的节点数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodesResponseParams `json:"Response"`
}

func (r *DescribeNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQueuesRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeQueuesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQueuesResponseParams struct {
	// 队列概览信息列表。
	QueueSet []*QueueOverview `json:"QueueSet,omitnil" name:"QueueSet"`

	// 符合条件的节点数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeQueuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQueuesResponseParams `json:"Response"`
}

func (r *DescribeQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpansionNodeConfig struct {
	// 扩容实例所在的位置。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>SPOTPAID：竞价付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。
	// <br><li>具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 私有网络相关信息配置。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil" name:"VirtualPrivateCloud"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type ExpansionNodeConfigOverview struct {
	// 节点机型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 扩容实例所在的位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 私有网络相关信息配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil" name:"VirtualPrivateCloud"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitnil" name:"ImageId"`

	// 公网带宽相关信息设置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil" name:"InternetAccessible"`

	// 节点系统盘配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 节点数据盘配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisks []*DataDisk `json:"DataDisks,omitnil" name:"DataDisks"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type GooseFSOption struct {
	// 文件系统本地挂载路径。
	LocalPath *string `json:"LocalPath,omitnil" name:"LocalPath"`

	// 文件系统远程挂载路径。
	RemotePath *string `json:"RemotePath,omitnil" name:"RemotePath"`

	// 文件系统master的ip和端口。
	Masters []*string `json:"Masters,omitnil" name:"Masters"`
}

type GooseFSOptionOverview struct {
	// 文件系统本地挂载路径。
	LocalPath *string `json:"LocalPath,omitnil" name:"LocalPath"`

	// 文件系统远程挂载路径。
	RemotePath *string `json:"RemotePath,omitnil" name:"RemotePath"`

	// 文件系统master的ip和端口。
	Masters []*string `json:"Masters,omitnil" name:"Masters"`
}

type GooseFSxOption struct {
	// 文件系统master的ip和端口列表。
	Masters []*string `json:"Masters,omitnil" name:"Masters"`

	// 文件系统的本地挂载路径。GooseFSx目前只支持挂载在/goosefsx/{文件系统ID}_proxy/目录下。
	LocalPath *string `json:"LocalPath,omitnil" name:"LocalPath"`
}

type GooseFSxOptionOverview struct {
	// 文件系统master的ip和端口列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Masters []*string `json:"Masters,omitnil" name:"Masters"`

	// 文件系统的本地挂载路径。GooseFSx目前只支持挂载在/goosefsx/{文件系统ID}_proxy/目录下。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalPath *string `json:"LocalPath,omitnil" name:"LocalPath"`
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 自动续费标识。取值范围：
	// NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费
	// DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type InternetAccessible struct {
	// 网络计费类型。取值范围：
	// BANDWIDTH_PREPAID：预付费按带宽结算
	// TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费
	// BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费
	// BANDWIDTH_PACKAGE：带宽包用户
	// 默认取值：非带宽包用户默认与子机付费类型保持一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetChargeType *string `json:"InternetChargeType,omitnil" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见购买网络带宽。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`
}

type LoginNode struct {
	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>SPOTPAID：竞价付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。
	// <br><li>具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil" name:"InternetAccessible"`

	// 节点显示名称。<br><li>
	// 不指定节点显示名称则默认显示‘未命名’。
	// 最多支持60个字符。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type LoginNodeOverview struct {
	// 登录节点ID。
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`
}

type LoginSettings struct {
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到30位，至少包括两项[a-z]，[A-Z]、[0-9] 和 [( ) \` ~ ! @ # $ % ^ & *  - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到30位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) \` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitnil" name:"Password"`
}

type ManagerNode struct {
	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID\_BY\_HOUR的实例创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil" name:"InternetAccessible"`

	// 节点显示名称。<br><li>
	// 不指定节点显示名称则默认显示‘未命名’。
	// </li><li>购买多个节点，如果指定模式串`{R:x}`，表示生成数字[`[x, x+n-1]`，其中`n`表示购买节点的数量，例如`server_{R:3}`，购买1个时，节点显示名称为`server_3`；购买2个时，节点显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。
	// 购买多个节点，如果不指定模式串，则在节点显示名称添加后缀`1、2...n`，其中`n`表示购买节点的数量，例如`server_`，购买2个时，节点显示名称分别为`server_1`，`server_2`。</li><li>
	// 最多支持60个字符（包含模式串）。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type ManagerNodeOverview struct {
	// 管控节点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`
}

// Predefined struct for user
type ModifyInitNodeScriptsRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 节点初始化脚本信息列表。
	InitNodeScripts []*NodeScript `json:"InitNodeScripts,omitnil" name:"InitNodeScripts"`
}

type ModifyInitNodeScriptsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 节点初始化脚本信息列表。
	InitNodeScripts []*NodeScript `json:"InitNodeScripts,omitnil" name:"InitNodeScripts"`
}

func (r *ModifyInitNodeScriptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInitNodeScriptsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InitNodeScripts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInitNodeScriptsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInitNodeScriptsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInitNodeScriptsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInitNodeScriptsResponseParams `json:"Response"`
}

func (r *ModifyInitNodeScriptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInitNodeScriptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeActivity struct {
	// 节点活动所在的实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInstanceId *string `json:"NodeInstanceId,omitnil" name:"NodeInstanceId"`

	// 节点活动状态。取值范围：<br><li>RUNNING：运行中<br><li>SUCCESSFUL：活动成功<br><li>FAILED：活动失败
	NodeActivityStatus *string `json:"NodeActivityStatus,omitnil" name:"NodeActivityStatus"`

	// 节点活动状态码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeActivityStatusCode *string `json:"NodeActivityStatusCode,omitnil" name:"NodeActivityStatusCode"`

	// 节点活动状态原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeActivityStatusReason *string `json:"NodeActivityStatusReason,omitnil" name:"NodeActivityStatusReason"`
}

type NodeOverview struct {
	// 节点实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点所在可用区信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 节点状态。<br><li>SUBMITTED：已完成提交。<br><li>CREATING：创建中。<br><li>CREATED：完成创建。<br><li>INITING：初始化中。<br><li>INIT_FAILED：初始化失败。<br><li>RUNNING：运行中。<br><li>DELETING：销毁中。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeState *string `json:"NodeState,omitnil" name:"NodeState"`

	// 镜像ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitnil" name:"ImageId"`

	// 节点所属队列名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// 节点角色。<br><li>Manager：管控节点。<br><li>Compute：计算节点。<br><li>Login：登录节点。<br><li>ManagerBackup：备用管控节点。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeRole *string `json:"NodeRole,omitnil" name:"NodeRole"`

	// 节点类型。<br><li>STATIC：静态节点。<br><li>DYNAMIC：弹性节点。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`
}

type NodeScript struct {
	// 节点执行脚本获取地址。
	// 目前仅支持cos地址。地址最大长度：255。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptPath *string `json:"ScriptPath,omitnil" name:"ScriptPath"`

	// 脚本执行超时时间（包含拉取脚本的时间）。单位秒，默认值：30。取值范围：10～1200。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`
}

type Placement struct {
	// 实例所属的可用区名称。该参数可以通过调用  [DescribeZones](https://cloud.tencent.com/document/product/213/15707) 的返回值中的Zone字段来获取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

type QueueConfig struct {
	// 队列名称。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// 队列中弹性节点数量最小值。默认值：0。取值范围：0～200。
	MinSize *uint64 `json:"MinSize,omitnil" name:"MinSize"`

	// 队列中弹性节点数量最大值。默认值：10。取值范围：0～200。
	MaxSize *uint64 `json:"MaxSize,omitnil" name:"MaxSize"`

	// 是否开启自动扩容。
	EnableAutoExpansion *bool `json:"EnableAutoExpansion,omitnil" name:"EnableAutoExpansion"`

	// 是否开启自动缩容。
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitnil" name:"EnableAutoShrink"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。目前仅支持公有镜和特定自定义镜像。
	ImageId *string `json:"ImageId,omitnil" name:"ImageId"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil" name:"InternetAccessible"`

	// 扩容节点配置信息。
	ExpansionNodeConfigs []*ExpansionNodeConfig `json:"ExpansionNodeConfigs,omitnil" name:"ExpansionNodeConfigs"`

	// 队列中期望的空闲节点数量（包含弹性节点和静态节点）。默认值：0。队列中，处于空闲状态的节点小于此值，集群会扩容弹性节点；处于空闲状态的节点大于此值，集群会缩容弹性节点。
	DesiredIdleNodeCapacity *int64 `json:"DesiredIdleNodeCapacity,omitnil" name:"DesiredIdleNodeCapacity"`

	// 扩容比例。默认值：100。取值范围：1～100。
	// 如果扩容比例为50，那么每轮只会扩容当前作业负载所需的50%数量的节点。
	ScaleOutRatio *int64 `json:"ScaleOutRatio,omitnil" name:"ScaleOutRatio"`

	// 比例扩容阈值。默认值：0。取值范围：0～200。
	// 当作业负载需要扩容节点数量大于此值，当前扩容轮次按照ScaleOutRatio配置的比例进行扩容。当作业负载需要扩容节点数量小于此值，当前扩容轮次扩容当前作业负载所需数量的节点。
	// 此参数配合ScaleOutRatio参数进行使用，用于比例扩容场景下，在作业负载所需节点数量较小时，加快收敛速度。
	ScaleOutNodeThreshold *int64 `json:"ScaleOutNodeThreshold,omitnil" name:"ScaleOutNodeThreshold"`

	// 每轮扩容最大节点个数。默认值：100。取值范围：1～100。
	MaxNodesPerCycle *int64 `json:"MaxNodesPerCycle,omitnil" name:"MaxNodesPerCycle"`

	// 扩容过程中，作业的内存在匹配实例机型时增大比例（不会影响作业提交的内存大小，只影响匹配计算过程）。<br/>
	// 针对场景：由于实例机型的总内存会大于实例内部的可用内存，16GB内存规格的实例，实例操作系统内的可用内存只有约14.9GB内存。假设此时提交一个需要15GB内存的作业，
	// 
	// - 当ScaleUpMemRatio=0时，会匹配到16GB内存规格的实例,但是由于操作系统内的可用内存为14.9GB小于作业所需的15GB，扩容出来的实例作业无法运行起来。
	// - 当ScaleUpMemRatio=10时，匹配实例规格会按照15*(1+10%)=16.5GB来进行实例规格匹配，则不会匹配到16GB的实例，而是更大内存规格的实例来保证作业能够被运行起来。
	ScaleUpMemRatio *int64 `json:"ScaleUpMemRatio,omitnil" name:"ScaleUpMemRatio"`
}

type QueueConfigOverview struct {
	// 队列名称。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// 队列中弹性节点数量最小值。取值范围0～200。
	MinSize *int64 `json:"MinSize,omitnil" name:"MinSize"`

	// 队列中弹性节点数量最大值。取值范围0～200。
	MaxSize *int64 `json:"MaxSize,omitnil" name:"MaxSize"`

	// 是否开启自动扩容。
	EnableAutoExpansion *bool `json:"EnableAutoExpansion,omitnil" name:"EnableAutoExpansion"`

	// 是否开启自动缩容。
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitnil" name:"EnableAutoShrink"`

	// 扩容节点配置信息。
	ExpansionNodeConfigs []*ExpansionNodeConfigOverview `json:"ExpansionNodeConfigs,omitnil" name:"ExpansionNodeConfigs"`

	// 队列中期望的空闲节点数量（包含弹性节点和静态节点）。默认值：0。队列中，处于空闲状态的节点小于此值，集群会扩容弹性节点；处于空闲状态的节点大于此值，集群会缩容弹性节点。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredIdleNodeCapacity *int64 `json:"DesiredIdleNodeCapacity,omitnil" name:"DesiredIdleNodeCapacity"`

	// 扩容比例。默认值：100。取值范围：1～100。
	// 如果扩容比例为50，那么每轮只会扩容当前作业负载所需的50%数量的节点。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleOutRatio *int64 `json:"ScaleOutRatio,omitnil" name:"ScaleOutRatio"`

	// 比例扩容阈值。默认值：0。取值范围：0～200。
	// 当作业负载需要扩容节点数量大于此值，当前扩容轮次按照ScaleOutRatio配置的的比例进行扩容。当作业负载需要扩容节点数量小于此值，当前扩容轮次扩容当前作业负载所需数量的节点。
	// 此参数配合ScaleOutRatio参数进行使用，用于比例扩容场景下，在作业负载所需节点数量较小时，加快收敛速度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleOutNodeThreshold *int64 `json:"ScaleOutNodeThreshold,omitnil" name:"ScaleOutNodeThreshold"`

	// 每轮扩容最大节点个数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxNodesPerCycle *int64 `json:"MaxNodesPerCycle,omitnil" name:"MaxNodesPerCycle"`

	// 扩容过程中，作业的内存在匹配实例机型时增大比例（不会影响作业提交的内存大小，只影响匹配计算过程）。<br/>
	// 针对场景：由于实例机型的总内存会大于实例内部的可用内存，16GB内存规格的实例，实例操作系统内的可用内存只有约14.9GB内存。假设此时提交一个需要15GB内存的作业，
	// 
	// - 当ScaleUpMemRatio=0时，会匹配到16GB内存规格的实例,但是由于操作系统内的可用内存为14.9GB小于作业所需的15GB，扩容出来的实例作业无法运行起来。
	// - 当ScaleUpMemRatio=10时，匹配实例规格会按照15*(1+10%)=16.5GB来进行实例规格匹配，则不会匹配到16GB的实例，而是更大内存规格的实例来保证作业能够被运行起来。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleUpMemRatio *int64 `json:"ScaleUpMemRatio,omitnil" name:"ScaleUpMemRatio"`
}

type QueueOverview struct {
	// 队列名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
}

// Predefined struct for user
type SetAutoScalingConfigurationRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 任务连续等待时间，队列的任务处于连续等待的时间。单位秒。默认值120。
	ExpansionBusyTime *int64 `json:"ExpansionBusyTime,omitnil" name:"ExpansionBusyTime"`

	// 节点连续空闲（未运行作业）时间，一个节点连续处于空闲状态时间。单位秒。默认值300。
	ShrinkIdleTime *int64 `json:"ShrinkIdleTime,omitnil" name:"ShrinkIdleTime"`

	// 扩容队列配置列表。
	QueueConfigs []*QueueConfig `json:"QueueConfigs,omitnil" name:"QueueConfigs"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会绑定弹性伸缩组。检查项包括是否填写了必需参数，请求格式，业务限制。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId。
	// false（默认）：发送正常请求，通过检查后直接绑定弹性伸缩组。
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`
}

type SetAutoScalingConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 任务连续等待时间，队列的任务处于连续等待的时间。单位秒。默认值120。
	ExpansionBusyTime *int64 `json:"ExpansionBusyTime,omitnil" name:"ExpansionBusyTime"`

	// 节点连续空闲（未运行作业）时间，一个节点连续处于空闲状态时间。单位秒。默认值300。
	ShrinkIdleTime *int64 `json:"ShrinkIdleTime,omitnil" name:"ShrinkIdleTime"`

	// 扩容队列配置列表。
	QueueConfigs []*QueueConfig `json:"QueueConfigs,omitnil" name:"QueueConfigs"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会绑定弹性伸缩组。检查项包括是否填写了必需参数，请求格式，业务限制。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId。
	// false（默认）：发送正常请求，通过检查后直接绑定弹性伸缩组。
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`
}

func (r *SetAutoScalingConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoScalingConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ExpansionBusyTime")
	delete(f, "ShrinkIdleTime")
	delete(f, "QueueConfigs")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAutoScalingConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAutoScalingConfigurationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetAutoScalingConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *SetAutoScalingConfigurationResponseParams `json:"Response"`
}

func (r *SetAutoScalingConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoScalingConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageOption struct {
	// 集群挂载CFS文件系统选项。
	CFSOptions []*CFSOption `json:"CFSOptions,omitnil" name:"CFSOptions"`

	// 集群挂载GooseFS文件系统选项。
	GooseFSOptions []*GooseFSOption `json:"GooseFSOptions,omitnil" name:"GooseFSOptions"`

	// 集群挂载GooseFSx文件系统选项。
	GooseFSxOptions []*GooseFSxOption `json:"GooseFSxOptions,omitnil" name:"GooseFSxOptions"`
}

type StorageOptionOverview struct {
	// CFS存储选项概览信息列表。
	CFSOptions []*CFSOptionOverview `json:"CFSOptions,omitnil" name:"CFSOptions"`

	// GooseFS存储选项概览信息列表。
	GooseFSOptions []*GooseFSOptionOverview `json:"GooseFSOptions,omitnil" name:"GooseFSOptions"`

	// GooseFSx存储选项概览信息列表。
	GooseFSxOptions []*GooseFSxOptionOverview `json:"GooseFSxOptions,omitnil" name:"GooseFSxOptions"`
}

type SystemDisk struct {
	// 系统盘类型。系统盘类型限制详见存储概述。取值范围：
	// CLOUD_BASIC：普通云硬盘
	// CLOUD_SSD：SSD云硬盘
	// CLOUD_PREMIUM：高性能云硬盘
	// 
	// 默认取值：当前有库存的硬盘类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 系统盘大小，单位：GB。默认值为 50
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type VirtualPrivateCloud struct {
	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台](https://console.cloud.tencent.com/vpc/vpc?rid=1)查询；也可以调用接口 [DescribeVpcEx](/document/api/215/1372) ，从接口返回中的`unVpcId`字段获取。若在创建子机时VpcId与SubnetId同时传入`DEFAULT`，则强制使用默认vpc网络。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台](https://console.cloud.tencent.com/vpc/subnet?rid=1)查询；也可以调用接口  [DescribeSubnets](/document/api/215/15784) ，从接口返回中的`unSubnetId`字段获取。若在创建子机时SubnetId与VpcId同时传入`DEFAULT`，则强制使用默认vpc网络。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}
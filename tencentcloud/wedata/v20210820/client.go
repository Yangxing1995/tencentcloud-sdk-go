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

package v20210820

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-08-20"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewBatchCreateIntegrationTaskAlarmsRequest() (request *BatchCreateIntegrationTaskAlarmsRequest) {
    request = &BatchCreateIntegrationTaskAlarmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchCreateIntegrationTaskAlarms")
    
    
    return
}

func NewBatchCreateIntegrationTaskAlarmsResponse() (response *BatchCreateIntegrationTaskAlarmsResponse) {
    response = &BatchCreateIntegrationTaskAlarmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchCreateIntegrationTaskAlarms
// 批量创建任务告警规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchCreateIntegrationTaskAlarms(request *BatchCreateIntegrationTaskAlarmsRequest) (response *BatchCreateIntegrationTaskAlarmsResponse, err error) {
    return c.BatchCreateIntegrationTaskAlarmsWithContext(context.Background(), request)
}

// BatchCreateIntegrationTaskAlarms
// 批量创建任务告警规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchCreateIntegrationTaskAlarmsWithContext(ctx context.Context, request *BatchCreateIntegrationTaskAlarmsRequest) (response *BatchCreateIntegrationTaskAlarmsResponse, err error) {
    if request == nil {
        request = NewBatchCreateIntegrationTaskAlarmsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchCreateIntegrationTaskAlarms require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchCreateIntegrationTaskAlarmsResponse()
    err = c.Send(request, response)
    return
}

func NewBatchCreateTaskVersionDsRequest() (request *BatchCreateTaskVersionDsRequest) {
    request = &BatchCreateTaskVersionDsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchCreateTaskVersionDs")
    
    
    return
}

func NewBatchCreateTaskVersionDsResponse() (response *BatchCreateTaskVersionDsResponse) {
    response = &BatchCreateTaskVersionDsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchCreateTaskVersionDs
// 批量创建任务版本Ds
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BatchCreateTaskVersionDs(request *BatchCreateTaskVersionDsRequest) (response *BatchCreateTaskVersionDsResponse, err error) {
    return c.BatchCreateTaskVersionDsWithContext(context.Background(), request)
}

// BatchCreateTaskVersionDs
// 批量创建任务版本Ds
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BatchCreateTaskVersionDsWithContext(ctx context.Context, request *BatchCreateTaskVersionDsRequest) (response *BatchCreateTaskVersionDsResponse, err error) {
    if request == nil {
        request = NewBatchCreateTaskVersionDsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchCreateTaskVersionDs require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchCreateTaskVersionDsResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteIntegrationTasksRequest() (request *BatchDeleteIntegrationTasksRequest) {
    request = &BatchDeleteIntegrationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchDeleteIntegrationTasks")
    
    
    return
}

func NewBatchDeleteIntegrationTasksResponse() (response *BatchDeleteIntegrationTasksResponse) {
    response = &BatchDeleteIntegrationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteIntegrationTasks
// 批量删除集成任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) BatchDeleteIntegrationTasks(request *BatchDeleteIntegrationTasksRequest) (response *BatchDeleteIntegrationTasksResponse, err error) {
    return c.BatchDeleteIntegrationTasksWithContext(context.Background(), request)
}

// BatchDeleteIntegrationTasks
// 批量删除集成任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) BatchDeleteIntegrationTasksWithContext(ctx context.Context, request *BatchDeleteIntegrationTasksRequest) (response *BatchDeleteIntegrationTasksResponse, err error) {
    if request == nil {
        request = NewBatchDeleteIntegrationTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteIntegrationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteIntegrationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteOpsTasksRequest() (request *BatchDeleteOpsTasksRequest) {
    request = &BatchDeleteOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchDeleteOpsTasks")
    
    
    return
}

func NewBatchDeleteOpsTasksResponse() (response *BatchDeleteOpsTasksResponse) {
    response = &BatchDeleteOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteOpsTasks
// 智能运维-批量删除任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchDeleteOpsTasks(request *BatchDeleteOpsTasksRequest) (response *BatchDeleteOpsTasksResponse, err error) {
    return c.BatchDeleteOpsTasksWithContext(context.Background(), request)
}

// BatchDeleteOpsTasks
// 智能运维-批量删除任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchDeleteOpsTasksWithContext(ctx context.Context, request *BatchDeleteOpsTasksRequest) (response *BatchDeleteOpsTasksResponse, err error) {
    if request == nil {
        request = NewBatchDeleteOpsTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteTasksDsRequest() (request *BatchDeleteTasksDsRequest) {
    request = &BatchDeleteTasksDsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchDeleteTasksDs")
    
    
    return
}

func NewBatchDeleteTasksDsResponse() (response *BatchDeleteTasksDsResponse) {
    response = &BatchDeleteTasksDsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteTasksDs
// Ds批量删除任务，仅对任务状态为”已停止“有效；
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchDeleteTasksDs(request *BatchDeleteTasksDsRequest) (response *BatchDeleteTasksDsResponse, err error) {
    return c.BatchDeleteTasksDsWithContext(context.Background(), request)
}

// BatchDeleteTasksDs
// Ds批量删除任务，仅对任务状态为”已停止“有效；
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchDeleteTasksDsWithContext(ctx context.Context, request *BatchDeleteTasksDsRequest) (response *BatchDeleteTasksDsResponse, err error) {
    if request == nil {
        request = NewBatchDeleteTasksDsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteTasksDs require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteTasksDsResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteTasksNewRequest() (request *BatchDeleteTasksNewRequest) {
    request = &BatchDeleteTasksNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchDeleteTasksNew")
    
    
    return
}

func NewBatchDeleteTasksNewResponse() (response *BatchDeleteTasksNewResponse) {
    response = &BatchDeleteTasksNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量删除任务，仅对任务状态为”已停止“有效；
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchDeleteTasksNew(request *BatchDeleteTasksNewRequest) (response *BatchDeleteTasksNewResponse, err error) {
    return c.BatchDeleteTasksNewWithContext(context.Background(), request)
}

// BatchDeleteTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量删除任务，仅对任务状态为”已停止“有效；
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchDeleteTasksNewWithContext(ctx context.Context, request *BatchDeleteTasksNewRequest) (response *BatchDeleteTasksNewResponse, err error) {
    if request == nil {
        request = NewBatchDeleteTasksNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteTasksNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteTasksNewResponse()
    err = c.Send(request, response)
    return
}

func NewBatchForceSuccessIntegrationTaskInstancesRequest() (request *BatchForceSuccessIntegrationTaskInstancesRequest) {
    request = &BatchForceSuccessIntegrationTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchForceSuccessIntegrationTaskInstances")
    
    
    return
}

func NewBatchForceSuccessIntegrationTaskInstancesResponse() (response *BatchForceSuccessIntegrationTaskInstancesResponse) {
    response = &BatchForceSuccessIntegrationTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchForceSuccessIntegrationTaskInstances
// 批量置成功集成任务实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchForceSuccessIntegrationTaskInstances(request *BatchForceSuccessIntegrationTaskInstancesRequest) (response *BatchForceSuccessIntegrationTaskInstancesResponse, err error) {
    return c.BatchForceSuccessIntegrationTaskInstancesWithContext(context.Background(), request)
}

// BatchForceSuccessIntegrationTaskInstances
// 批量置成功集成任务实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchForceSuccessIntegrationTaskInstancesWithContext(ctx context.Context, request *BatchForceSuccessIntegrationTaskInstancesRequest) (response *BatchForceSuccessIntegrationTaskInstancesResponse, err error) {
    if request == nil {
        request = NewBatchForceSuccessIntegrationTaskInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchForceSuccessIntegrationTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchForceSuccessIntegrationTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewBatchKillIntegrationTaskInstancesRequest() (request *BatchKillIntegrationTaskInstancesRequest) {
    request = &BatchKillIntegrationTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchKillIntegrationTaskInstances")
    
    
    return
}

func NewBatchKillIntegrationTaskInstancesResponse() (response *BatchKillIntegrationTaskInstancesResponse) {
    response = &BatchKillIntegrationTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchKillIntegrationTaskInstances
// 批量终止集成任务实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchKillIntegrationTaskInstances(request *BatchKillIntegrationTaskInstancesRequest) (response *BatchKillIntegrationTaskInstancesResponse, err error) {
    return c.BatchKillIntegrationTaskInstancesWithContext(context.Background(), request)
}

// BatchKillIntegrationTaskInstances
// 批量终止集成任务实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchKillIntegrationTaskInstancesWithContext(ctx context.Context, request *BatchKillIntegrationTaskInstancesRequest) (response *BatchKillIntegrationTaskInstancesResponse, err error) {
    if request == nil {
        request = NewBatchKillIntegrationTaskInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchKillIntegrationTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchKillIntegrationTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewBatchMakeUpIntegrationTasksRequest() (request *BatchMakeUpIntegrationTasksRequest) {
    request = &BatchMakeUpIntegrationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchMakeUpIntegrationTasks")
    
    
    return
}

func NewBatchMakeUpIntegrationTasksResponse() (response *BatchMakeUpIntegrationTasksResponse) {
    response = &BatchMakeUpIntegrationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchMakeUpIntegrationTasks
// 对集成离线任务执行批量补数据操作
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchMakeUpIntegrationTasks(request *BatchMakeUpIntegrationTasksRequest) (response *BatchMakeUpIntegrationTasksResponse, err error) {
    return c.BatchMakeUpIntegrationTasksWithContext(context.Background(), request)
}

// BatchMakeUpIntegrationTasks
// 对集成离线任务执行批量补数据操作
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchMakeUpIntegrationTasksWithContext(ctx context.Context, request *BatchMakeUpIntegrationTasksRequest) (response *BatchMakeUpIntegrationTasksResponse, err error) {
    if request == nil {
        request = NewBatchMakeUpIntegrationTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchMakeUpIntegrationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchMakeUpIntegrationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyOpsOwnersRequest() (request *BatchModifyOpsOwnersRequest) {
    request = &BatchModifyOpsOwnersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchModifyOpsOwners")
    
    
    return
}

func NewBatchModifyOpsOwnersResponse() (response *BatchModifyOpsOwnersResponse) {
    response = &BatchModifyOpsOwnersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchModifyOpsOwners
// 批量修改任务责任人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchModifyOpsOwners(request *BatchModifyOpsOwnersRequest) (response *BatchModifyOpsOwnersResponse, err error) {
    return c.BatchModifyOpsOwnersWithContext(context.Background(), request)
}

// BatchModifyOpsOwners
// 批量修改任务责任人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchModifyOpsOwnersWithContext(ctx context.Context, request *BatchModifyOpsOwnersRequest) (response *BatchModifyOpsOwnersResponse, err error) {
    if request == nil {
        request = NewBatchModifyOpsOwnersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyOpsOwners require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyOpsOwnersResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyOwnersNewRequest() (request *BatchModifyOwnersNewRequest) {
    request = &BatchModifyOwnersNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchModifyOwnersNew")
    
    
    return
}

func NewBatchModifyOwnersNewResponse() (response *BatchModifyOwnersNewResponse) {
    response = &BatchModifyOwnersNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchModifyOwnersNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量修改任务责任人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchModifyOwnersNew(request *BatchModifyOwnersNewRequest) (response *BatchModifyOwnersNewResponse, err error) {
    return c.BatchModifyOwnersNewWithContext(context.Background(), request)
}

// BatchModifyOwnersNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量修改任务责任人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchModifyOwnersNewWithContext(ctx context.Context, request *BatchModifyOwnersNewRequest) (response *BatchModifyOwnersNewResponse, err error) {
    if request == nil {
        request = NewBatchModifyOwnersNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyOwnersNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyOwnersNewResponse()
    err = c.Send(request, response)
    return
}

func NewBatchRerunIntegrationTaskInstancesRequest() (request *BatchRerunIntegrationTaskInstancesRequest) {
    request = &BatchRerunIntegrationTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchRerunIntegrationTaskInstances")
    
    
    return
}

func NewBatchRerunIntegrationTaskInstancesResponse() (response *BatchRerunIntegrationTaskInstancesResponse) {
    response = &BatchRerunIntegrationTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchRerunIntegrationTaskInstances
// 批量重跑集成任务实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) BatchRerunIntegrationTaskInstances(request *BatchRerunIntegrationTaskInstancesRequest) (response *BatchRerunIntegrationTaskInstancesResponse, err error) {
    return c.BatchRerunIntegrationTaskInstancesWithContext(context.Background(), request)
}

// BatchRerunIntegrationTaskInstances
// 批量重跑集成任务实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) BatchRerunIntegrationTaskInstancesWithContext(ctx context.Context, request *BatchRerunIntegrationTaskInstancesRequest) (response *BatchRerunIntegrationTaskInstancesResponse, err error) {
    if request == nil {
        request = NewBatchRerunIntegrationTaskInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchRerunIntegrationTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchRerunIntegrationTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewBatchResumeIntegrationTasksRequest() (request *BatchResumeIntegrationTasksRequest) {
    request = &BatchResumeIntegrationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchResumeIntegrationTasks")
    
    
    return
}

func NewBatchResumeIntegrationTasksResponse() (response *BatchResumeIntegrationTasksResponse) {
    response = &BatchResumeIntegrationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchResumeIntegrationTasks
// 批量继续执行集成实时任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) BatchResumeIntegrationTasks(request *BatchResumeIntegrationTasksRequest) (response *BatchResumeIntegrationTasksResponse, err error) {
    return c.BatchResumeIntegrationTasksWithContext(context.Background(), request)
}

// BatchResumeIntegrationTasks
// 批量继续执行集成实时任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) BatchResumeIntegrationTasksWithContext(ctx context.Context, request *BatchResumeIntegrationTasksRequest) (response *BatchResumeIntegrationTasksResponse, err error) {
    if request == nil {
        request = NewBatchResumeIntegrationTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchResumeIntegrationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchResumeIntegrationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewBatchRunOpsTaskRequest() (request *BatchRunOpsTaskRequest) {
    request = &BatchRunOpsTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchRunOpsTask")
    
    
    return
}

func NewBatchRunOpsTaskResponse() (response *BatchRunOpsTaskResponse) {
    response = &BatchRunOpsTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchRunOpsTask
// 任务运维-任务列表 批量运行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchRunOpsTask(request *BatchRunOpsTaskRequest) (response *BatchRunOpsTaskResponse, err error) {
    return c.BatchRunOpsTaskWithContext(context.Background(), request)
}

// BatchRunOpsTask
// 任务运维-任务列表 批量运行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchRunOpsTaskWithContext(ctx context.Context, request *BatchRunOpsTaskRequest) (response *BatchRunOpsTaskResponse, err error) {
    if request == nil {
        request = NewBatchRunOpsTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchRunOpsTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchRunOpsTaskResponse()
    err = c.Send(request, response)
    return
}

func NewBatchStartIntegrationTasksRequest() (request *BatchStartIntegrationTasksRequest) {
    request = &BatchStartIntegrationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchStartIntegrationTasks")
    
    
    return
}

func NewBatchStartIntegrationTasksResponse() (response *BatchStartIntegrationTasksResponse) {
    response = &BatchStartIntegrationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchStartIntegrationTasks
// 批量运行集成任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchStartIntegrationTasks(request *BatchStartIntegrationTasksRequest) (response *BatchStartIntegrationTasksResponse, err error) {
    return c.BatchStartIntegrationTasksWithContext(context.Background(), request)
}

// BatchStartIntegrationTasks
// 批量运行集成任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchStartIntegrationTasksWithContext(ctx context.Context, request *BatchStartIntegrationTasksRequest) (response *BatchStartIntegrationTasksResponse, err error) {
    if request == nil {
        request = NewBatchStartIntegrationTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchStartIntegrationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchStartIntegrationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewBatchStopIntegrationTasksRequest() (request *BatchStopIntegrationTasksRequest) {
    request = &BatchStopIntegrationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchStopIntegrationTasks")
    
    
    return
}

func NewBatchStopIntegrationTasksResponse() (response *BatchStopIntegrationTasksResponse) {
    response = &BatchStopIntegrationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchStopIntegrationTasks
// 批量停止集成任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) BatchStopIntegrationTasks(request *BatchStopIntegrationTasksRequest) (response *BatchStopIntegrationTasksResponse, err error) {
    return c.BatchStopIntegrationTasksWithContext(context.Background(), request)
}

// BatchStopIntegrationTasks
// 批量停止集成任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) BatchStopIntegrationTasksWithContext(ctx context.Context, request *BatchStopIntegrationTasksRequest) (response *BatchStopIntegrationTasksResponse, err error) {
    if request == nil {
        request = NewBatchStopIntegrationTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchStopIntegrationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchStopIntegrationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewBatchStopOpsTasksRequest() (request *BatchStopOpsTasksRequest) {
    request = &BatchStopOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchStopOpsTasks")
    
    
    return
}

func NewBatchStopOpsTasksResponse() (response *BatchStopOpsTasksResponse) {
    response = &BatchStopOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchStopOpsTasks
// 仅对任务状态为”调度中“和”已暂停“有效，对所选任务的任务实例进行终止，并停止调度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchStopOpsTasks(request *BatchStopOpsTasksRequest) (response *BatchStopOpsTasksResponse, err error) {
    return c.BatchStopOpsTasksWithContext(context.Background(), request)
}

// BatchStopOpsTasks
// 仅对任务状态为”调度中“和”已暂停“有效，对所选任务的任务实例进行终止，并停止调度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchStopOpsTasksWithContext(ctx context.Context, request *BatchStopOpsTasksRequest) (response *BatchStopOpsTasksResponse, err error) {
    if request == nil {
        request = NewBatchStopOpsTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchStopOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchStopOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewBatchStopTasksNewRequest() (request *BatchStopTasksNewRequest) {
    request = &BatchStopTasksNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchStopTasksNew")
    
    
    return
}

func NewBatchStopTasksNewResponse() (response *BatchStopTasksNewResponse) {
    response = &BatchStopTasksNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchStopTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 仅对任务状态为”调度中“和”已暂停“有效，对所选任务的任务实例进行终止，并停止调度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchStopTasksNew(request *BatchStopTasksNewRequest) (response *BatchStopTasksNewResponse, err error) {
    return c.BatchStopTasksNewWithContext(context.Background(), request)
}

// BatchStopTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 仅对任务状态为”调度中“和”已暂停“有效，对所选任务的任务实例进行终止，并停止调度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchStopTasksNewWithContext(ctx context.Context, request *BatchStopTasksNewRequest) (response *BatchStopTasksNewResponse, err error) {
    if request == nil {
        request = NewBatchStopTasksNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchStopTasksNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchStopTasksNewResponse()
    err = c.Send(request, response)
    return
}

func NewBatchStopWorkflowsByIdsRequest() (request *BatchStopWorkflowsByIdsRequest) {
    request = &BatchStopWorkflowsByIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchStopWorkflowsByIds")
    
    
    return
}

func NewBatchStopWorkflowsByIdsResponse() (response *BatchStopWorkflowsByIdsResponse) {
    response = &BatchStopWorkflowsByIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchStopWorkflowsByIds
// 批量停止工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchStopWorkflowsByIds(request *BatchStopWorkflowsByIdsRequest) (response *BatchStopWorkflowsByIdsResponse, err error) {
    return c.BatchStopWorkflowsByIdsWithContext(context.Background(), request)
}

// BatchStopWorkflowsByIds
// 批量停止工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) BatchStopWorkflowsByIdsWithContext(ctx context.Context, request *BatchStopWorkflowsByIdsRequest) (response *BatchStopWorkflowsByIdsResponse, err error) {
    if request == nil {
        request = NewBatchStopWorkflowsByIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchStopWorkflowsByIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchStopWorkflowsByIdsResponse()
    err = c.Send(request, response)
    return
}

func NewBatchSuspendIntegrationTasksRequest() (request *BatchSuspendIntegrationTasksRequest) {
    request = &BatchSuspendIntegrationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchSuspendIntegrationTasks")
    
    
    return
}

func NewBatchSuspendIntegrationTasksResponse() (response *BatchSuspendIntegrationTasksResponse) {
    response = &BatchSuspendIntegrationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchSuspendIntegrationTasks
// 批量暂停集成任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchSuspendIntegrationTasks(request *BatchSuspendIntegrationTasksRequest) (response *BatchSuspendIntegrationTasksResponse, err error) {
    return c.BatchSuspendIntegrationTasksWithContext(context.Background(), request)
}

// BatchSuspendIntegrationTasks
// 批量暂停集成任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchSuspendIntegrationTasksWithContext(ctx context.Context, request *BatchSuspendIntegrationTasksRequest) (response *BatchSuspendIntegrationTasksResponse, err error) {
    if request == nil {
        request = NewBatchSuspendIntegrationTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchSuspendIntegrationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchSuspendIntegrationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewBatchUpdateIntegrationTasksRequest() (request *BatchUpdateIntegrationTasksRequest) {
    request = &BatchUpdateIntegrationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "BatchUpdateIntegrationTasks")
    
    
    return
}

func NewBatchUpdateIntegrationTasksResponse() (response *BatchUpdateIntegrationTasksResponse) {
    response = &BatchUpdateIntegrationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchUpdateIntegrationTasks
// 批量更新集成任务（暂时仅支持批量更新责任人）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchUpdateIntegrationTasks(request *BatchUpdateIntegrationTasksRequest) (response *BatchUpdateIntegrationTasksResponse, err error) {
    return c.BatchUpdateIntegrationTasksWithContext(context.Background(), request)
}

// BatchUpdateIntegrationTasks
// 批量更新集成任务（暂时仅支持批量更新责任人）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) BatchUpdateIntegrationTasksWithContext(ctx context.Context, request *BatchUpdateIntegrationTasksRequest) (response *BatchUpdateIntegrationTasksResponse, err error) {
    if request == nil {
        request = NewBatchUpdateIntegrationTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchUpdateIntegrationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchUpdateIntegrationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAlarmRegularNameExistRequest() (request *CheckAlarmRegularNameExistRequest) {
    request = &CheckAlarmRegularNameExistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CheckAlarmRegularNameExist")
    
    
    return
}

func NewCheckAlarmRegularNameExistResponse() (response *CheckAlarmRegularNameExistResponse) {
    response = &CheckAlarmRegularNameExistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckAlarmRegularNameExist
// 判断告警规则重名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckAlarmRegularNameExist(request *CheckAlarmRegularNameExistRequest) (response *CheckAlarmRegularNameExistResponse, err error) {
    return c.CheckAlarmRegularNameExistWithContext(context.Background(), request)
}

// CheckAlarmRegularNameExist
// 判断告警规则重名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckAlarmRegularNameExistWithContext(ctx context.Context, request *CheckAlarmRegularNameExistRequest) (response *CheckAlarmRegularNameExistResponse, err error) {
    if request == nil {
        request = NewCheckAlarmRegularNameExistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAlarmRegularNameExist require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAlarmRegularNameExistResponse()
    err = c.Send(request, response)
    return
}

func NewCheckDuplicateRuleNameRequest() (request *CheckDuplicateRuleNameRequest) {
    request = &CheckDuplicateRuleNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CheckDuplicateRuleName")
    
    
    return
}

func NewCheckDuplicateRuleNameResponse() (response *CheckDuplicateRuleNameResponse) {
    response = &CheckDuplicateRuleNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckDuplicateRuleName
// 检查规则名称是否重复
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckDuplicateRuleName(request *CheckDuplicateRuleNameRequest) (response *CheckDuplicateRuleNameResponse, err error) {
    return c.CheckDuplicateRuleNameWithContext(context.Background(), request)
}

// CheckDuplicateRuleName
// 检查规则名称是否重复
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckDuplicateRuleNameWithContext(ctx context.Context, request *CheckDuplicateRuleNameRequest) (response *CheckDuplicateRuleNameResponse, err error) {
    if request == nil {
        request = NewCheckDuplicateRuleNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckDuplicateRuleName require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckDuplicateRuleNameResponse()
    err = c.Send(request, response)
    return
}

func NewCheckDuplicateTemplateNameRequest() (request *CheckDuplicateTemplateNameRequest) {
    request = &CheckDuplicateTemplateNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CheckDuplicateTemplateName")
    
    
    return
}

func NewCheckDuplicateTemplateNameResponse() (response *CheckDuplicateTemplateNameResponse) {
    response = &CheckDuplicateTemplateNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckDuplicateTemplateName
// 检查规则模板名称是否重复
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckDuplicateTemplateName(request *CheckDuplicateTemplateNameRequest) (response *CheckDuplicateTemplateNameResponse, err error) {
    return c.CheckDuplicateTemplateNameWithContext(context.Background(), request)
}

// CheckDuplicateTemplateName
// 检查规则模板名称是否重复
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckDuplicateTemplateNameWithContext(ctx context.Context, request *CheckDuplicateTemplateNameRequest) (response *CheckDuplicateTemplateNameResponse, err error) {
    if request == nil {
        request = NewCheckDuplicateTemplateNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckDuplicateTemplateName require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckDuplicateTemplateNameResponse()
    err = c.Send(request, response)
    return
}

func NewCheckIntegrationNodeNameExistsRequest() (request *CheckIntegrationNodeNameExistsRequest) {
    request = &CheckIntegrationNodeNameExistsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CheckIntegrationNodeNameExists")
    
    
    return
}

func NewCheckIntegrationNodeNameExistsResponse() (response *CheckIntegrationNodeNameExistsResponse) {
    response = &CheckIntegrationNodeNameExistsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckIntegrationNodeNameExists
// 判断集成节点名称是否存在
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckIntegrationNodeNameExists(request *CheckIntegrationNodeNameExistsRequest) (response *CheckIntegrationNodeNameExistsResponse, err error) {
    return c.CheckIntegrationNodeNameExistsWithContext(context.Background(), request)
}

// CheckIntegrationNodeNameExists
// 判断集成节点名称是否存在
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckIntegrationNodeNameExistsWithContext(ctx context.Context, request *CheckIntegrationNodeNameExistsRequest) (response *CheckIntegrationNodeNameExistsResponse, err error) {
    if request == nil {
        request = NewCheckIntegrationNodeNameExistsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckIntegrationNodeNameExists require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckIntegrationNodeNameExistsResponse()
    err = c.Send(request, response)
    return
}

func NewCheckIntegrationTaskNameExistsRequest() (request *CheckIntegrationTaskNameExistsRequest) {
    request = &CheckIntegrationTaskNameExistsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CheckIntegrationTaskNameExists")
    
    
    return
}

func NewCheckIntegrationTaskNameExistsResponse() (response *CheckIntegrationTaskNameExistsResponse) {
    response = &CheckIntegrationTaskNameExistsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckIntegrationTaskNameExists
// 判断集成任务名称是否存在
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckIntegrationTaskNameExists(request *CheckIntegrationTaskNameExistsRequest) (response *CheckIntegrationTaskNameExistsResponse, err error) {
    return c.CheckIntegrationTaskNameExistsWithContext(context.Background(), request)
}

// CheckIntegrationTaskNameExists
// 判断集成任务名称是否存在
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckIntegrationTaskNameExistsWithContext(ctx context.Context, request *CheckIntegrationTaskNameExistsRequest) (response *CheckIntegrationTaskNameExistsResponse, err error) {
    if request == nil {
        request = NewCheckIntegrationTaskNameExistsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckIntegrationTaskNameExists require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckIntegrationTaskNameExistsResponse()
    err = c.Send(request, response)
    return
}

func NewCheckTaskNameExistRequest() (request *CheckTaskNameExistRequest) {
    request = &CheckTaskNameExistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CheckTaskNameExist")
    
    
    return
}

func NewCheckTaskNameExistResponse() (response *CheckTaskNameExistResponse) {
    response = &CheckTaskNameExistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckTaskNameExist
// 离线任务重名校验
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckTaskNameExist(request *CheckTaskNameExistRequest) (response *CheckTaskNameExistResponse, err error) {
    return c.CheckTaskNameExistWithContext(context.Background(), request)
}

// CheckTaskNameExist
// 离线任务重名校验
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CheckTaskNameExistWithContext(ctx context.Context, request *CheckTaskNameExistRequest) (response *CheckTaskNameExistResponse, err error) {
    if request == nil {
        request = NewCheckTaskNameExistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckTaskNameExist require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckTaskNameExistResponse()
    err = c.Send(request, response)
    return
}

func NewCommitExportTaskRequest() (request *CommitExportTaskRequest) {
    request = &CommitExportTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CommitExportTask")
    
    
    return
}

func NewCommitExportTaskResponse() (response *CommitExportTaskResponse) {
    response = &CommitExportTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CommitExportTask
// 提交数据导出任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CommitExportTask(request *CommitExportTaskRequest) (response *CommitExportTaskResponse, err error) {
    return c.CommitExportTaskWithContext(context.Background(), request)
}

// CommitExportTask
// 提交数据导出任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CommitExportTaskWithContext(ctx context.Context, request *CommitExportTaskRequest) (response *CommitExportTaskResponse, err error) {
    if request == nil {
        request = NewCommitExportTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CommitExportTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCommitExportTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCommitIntegrationTaskRequest() (request *CommitIntegrationTaskRequest) {
    request = &CommitIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CommitIntegrationTask")
    
    
    return
}

func NewCommitIntegrationTaskResponse() (response *CommitIntegrationTaskResponse) {
    response = &CommitIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CommitIntegrationTask
// 提交集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CommitIntegrationTask(request *CommitIntegrationTaskRequest) (response *CommitIntegrationTaskResponse, err error) {
    return c.CommitIntegrationTaskWithContext(context.Background(), request)
}

// CommitIntegrationTask
// 提交集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CommitIntegrationTaskWithContext(ctx context.Context, request *CommitIntegrationTaskRequest) (response *CommitIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewCommitIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CommitIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCommitIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCommitRuleGroupExecResultRequest() (request *CommitRuleGroupExecResultRequest) {
    request = &CommitRuleGroupExecResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CommitRuleGroupExecResult")
    
    
    return
}

func NewCommitRuleGroupExecResultResponse() (response *CommitRuleGroupExecResultResponse) {
    response = &CommitRuleGroupExecResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CommitRuleGroupExecResult
// Runner 规则检测结果上报
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CommitRuleGroupExecResult(request *CommitRuleGroupExecResultRequest) (response *CommitRuleGroupExecResultResponse, err error) {
    return c.CommitRuleGroupExecResultWithContext(context.Background(), request)
}

// CommitRuleGroupExecResult
// Runner 规则检测结果上报
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CommitRuleGroupExecResultWithContext(ctx context.Context, request *CommitRuleGroupExecResultRequest) (response *CommitRuleGroupExecResultResponse, err error) {
    if request == nil {
        request = NewCommitRuleGroupExecResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CommitRuleGroupExecResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewCommitRuleGroupExecResultResponse()
    err = c.Send(request, response)
    return
}

func NewCommitRuleGroupTaskRequest() (request *CommitRuleGroupTaskRequest) {
    request = &CommitRuleGroupTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CommitRuleGroupTask")
    
    
    return
}

func NewCommitRuleGroupTaskResponse() (response *CommitRuleGroupTaskResponse) {
    response = &CommitRuleGroupTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CommitRuleGroupTask
// 提交规则组运行任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CommitRuleGroupTask(request *CommitRuleGroupTaskRequest) (response *CommitRuleGroupTaskResponse, err error) {
    return c.CommitRuleGroupTaskWithContext(context.Background(), request)
}

// CommitRuleGroupTask
// 提交规则组运行任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CommitRuleGroupTaskWithContext(ctx context.Context, request *CommitRuleGroupTaskRequest) (response *CommitRuleGroupTaskResponse, err error) {
    if request == nil {
        request = NewCommitRuleGroupTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CommitRuleGroupTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCommitRuleGroupTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCountOpsInstanceStateRequest() (request *CountOpsInstanceStateRequest) {
    request = &CountOpsInstanceStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CountOpsInstanceState")
    
    
    return
}

func NewCountOpsInstanceStateResponse() (response *CountOpsInstanceStateResponse) {
    response = &CountOpsInstanceStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CountOpsInstanceState
// 统计任务实例状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CountOpsInstanceState(request *CountOpsInstanceStateRequest) (response *CountOpsInstanceStateResponse, err error) {
    return c.CountOpsInstanceStateWithContext(context.Background(), request)
}

// CountOpsInstanceState
// 统计任务实例状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CountOpsInstanceStateWithContext(ctx context.Context, request *CountOpsInstanceStateRequest) (response *CountOpsInstanceStateResponse, err error) {
    if request == nil {
        request = NewCountOpsInstanceStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CountOpsInstanceState require credential")
    }

    request.SetContext(ctx)
    
    response = NewCountOpsInstanceStateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBaselineRequest() (request *CreateBaselineRequest) {
    request = &CreateBaselineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateBaseline")
    
    
    return
}

func NewCreateBaselineResponse() (response *CreateBaselineResponse) {
    response = &CreateBaselineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBaseline
// 创建基线
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateBaseline(request *CreateBaselineRequest) (response *CreateBaselineResponse, err error) {
    return c.CreateBaselineWithContext(context.Background(), request)
}

// CreateBaseline
// 创建基线
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateBaselineWithContext(ctx context.Context, request *CreateBaselineRequest) (response *CreateBaselineResponse, err error) {
    if request == nil {
        request = NewCreateBaselineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBaseline require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBaselineResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomFunctionRequest() (request *CreateCustomFunctionRequest) {
    request = &CreateCustomFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateCustomFunction")
    
    
    return
}

func NewCreateCustomFunctionResponse() (response *CreateCustomFunctionResponse) {
    response = &CreateCustomFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCustomFunction
// 创建用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCustomFunction(request *CreateCustomFunctionRequest) (response *CreateCustomFunctionResponse, err error) {
    return c.CreateCustomFunctionWithContext(context.Background(), request)
}

// CreateCustomFunction
// 创建用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCustomFunctionWithContext(ctx context.Context, request *CreateCustomFunctionRequest) (response *CreateCustomFunctionResponse, err error) {
    if request == nil {
        request = NewCreateCustomFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataSourceRequest() (request *CreateDataSourceRequest) {
    request = &CreateDataSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateDataSource")
    
    
    return
}

func NewCreateDataSourceResponse() (response *CreateDataSourceResponse) {
    response = &CreateDataSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataSource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDataSource(request *CreateDataSourceRequest) (response *CreateDataSourceResponse, err error) {
    return c.CreateDataSourceWithContext(context.Background(), request)
}

// CreateDataSource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDataSourceWithContext(ctx context.Context, request *CreateDataSourceRequest) (response *CreateDataSourceResponse, err error) {
    if request == nil {
        request = NewCreateDataSourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataSourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDsFolderRequest() (request *CreateDsFolderRequest) {
    request = &CreateDsFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateDsFolder")
    
    
    return
}

func NewCreateDsFolderResponse() (response *CreateDsFolderResponse) {
    response = &CreateDsFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDsFolder
// 编排空间-创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDsFolder(request *CreateDsFolderRequest) (response *CreateDsFolderResponse, err error) {
    return c.CreateDsFolderWithContext(context.Background(), request)
}

// CreateDsFolder
// 编排空间-创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDsFolderWithContext(ctx context.Context, request *CreateDsFolderRequest) (response *CreateDsFolderResponse, err error) {
    if request == nil {
        request = NewCreateDsFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDsFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDsFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFolderRequest() (request *CreateFolderRequest) {
    request = &CreateFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateFolder")
    
    
    return
}

func NewCreateFolderResponse() (response *CreateFolderResponse) {
    response = &CreateFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateFolder(request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    return c.CreateFolderWithContext(context.Background(), request)
}

// CreateFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateFolderWithContext(ctx context.Context, request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    if request == nil {
        request = NewCreateFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHiveTableRequest() (request *CreateHiveTableRequest) {
    request = &CreateHiveTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateHiveTable")
    
    
    return
}

func NewCreateHiveTableResponse() (response *CreateHiveTableResponse) {
    response = &CreateHiveTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHiveTable
// 建hive表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateHiveTable(request *CreateHiveTableRequest) (response *CreateHiveTableResponse, err error) {
    return c.CreateHiveTableWithContext(context.Background(), request)
}

// CreateHiveTable
// 建hive表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateHiveTableWithContext(ctx context.Context, request *CreateHiveTableRequest) (response *CreateHiveTableResponse, err error) {
    if request == nil {
        request = NewCreateHiveTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHiveTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHiveTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHiveTableByDDLRequest() (request *CreateHiveTableByDDLRequest) {
    request = &CreateHiveTableByDDLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateHiveTableByDDL")
    
    
    return
}

func NewCreateHiveTableByDDLResponse() (response *CreateHiveTableByDDLResponse) {
    response = &CreateHiveTableByDDLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateHiveTableByDDL
// 创建hive表，返回表名称
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateHiveTableByDDL(request *CreateHiveTableByDDLRequest) (response *CreateHiveTableByDDLResponse, err error) {
    return c.CreateHiveTableByDDLWithContext(context.Background(), request)
}

// CreateHiveTableByDDL
// 创建hive表，返回表名称
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateHiveTableByDDLWithContext(ctx context.Context, request *CreateHiveTableByDDLRequest) (response *CreateHiveTableByDDLResponse, err error) {
    if request == nil {
        request = NewCreateHiveTableByDDLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHiveTableByDDL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHiveTableByDDLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInLongAgentRequest() (request *CreateInLongAgentRequest) {
    request = &CreateInLongAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateInLongAgent")
    
    
    return
}

func NewCreateInLongAgentResponse() (response *CreateInLongAgentResponse) {
    response = &CreateInLongAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInLongAgent
// 注册采集器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateInLongAgent(request *CreateInLongAgentRequest) (response *CreateInLongAgentResponse, err error) {
    return c.CreateInLongAgentWithContext(context.Background(), request)
}

// CreateInLongAgent
// 注册采集器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateInLongAgentWithContext(ctx context.Context, request *CreateInLongAgentRequest) (response *CreateInLongAgentResponse, err error) {
    if request == nil {
        request = NewCreateInLongAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInLongAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInLongAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntegrationNodeRequest() (request *CreateIntegrationNodeRequest) {
    request = &CreateIntegrationNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateIntegrationNode")
    
    
    return
}

func NewCreateIntegrationNodeResponse() (response *CreateIntegrationNodeResponse) {
    response = &CreateIntegrationNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIntegrationNode
// 创建集成节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIntegrationNode(request *CreateIntegrationNodeRequest) (response *CreateIntegrationNodeResponse, err error) {
    return c.CreateIntegrationNodeWithContext(context.Background(), request)
}

// CreateIntegrationNode
// 创建集成节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIntegrationNodeWithContext(ctx context.Context, request *CreateIntegrationNodeRequest) (response *CreateIntegrationNodeResponse, err error) {
    if request == nil {
        request = NewCreateIntegrationNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntegrationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntegrationNodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntegrationTaskRequest() (request *CreateIntegrationTaskRequest) {
    request = &CreateIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateIntegrationTask")
    
    
    return
}

func NewCreateIntegrationTaskResponse() (response *CreateIntegrationTaskResponse) {
    response = &CreateIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIntegrationTask
// 创建集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIntegrationTask(request *CreateIntegrationTaskRequest) (response *CreateIntegrationTaskResponse, err error) {
    return c.CreateIntegrationTaskWithContext(context.Background(), request)
}

// CreateIntegrationTask
// 创建集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIntegrationTaskWithContext(ctx context.Context, request *CreateIntegrationTaskRequest) (response *CreateIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewCreateIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOfflineTaskRequest() (request *CreateOfflineTaskRequest) {
    request = &CreateOfflineTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateOfflineTask")
    
    
    return
}

func NewCreateOfflineTaskResponse() (response *CreateOfflineTaskResponse) {
    response = &CreateOfflineTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOfflineTask
// 创建离线任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateOfflineTask(request *CreateOfflineTaskRequest) (response *CreateOfflineTaskResponse, err error) {
    return c.CreateOfflineTaskWithContext(context.Background(), request)
}

// CreateOfflineTask
// 创建离线任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateOfflineTaskWithContext(ctx context.Context, request *CreateOfflineTaskRequest) (response *CreateOfflineTaskResponse, err error) {
    if request == nil {
        request = NewCreateOfflineTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOfflineTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOfflineTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpsMakePlanRequest() (request *CreateOpsMakePlanRequest) {
    request = &CreateOpsMakePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateOpsMakePlan")
    
    
    return
}

func NewCreateOpsMakePlanResponse() (response *CreateOpsMakePlanResponse) {
    response = &CreateOpsMakePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOpsMakePlan
// 补录任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateOpsMakePlan(request *CreateOpsMakePlanRequest) (response *CreateOpsMakePlanResponse, err error) {
    return c.CreateOpsMakePlanWithContext(context.Background(), request)
}

// CreateOpsMakePlan
// 补录任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateOpsMakePlanWithContext(ctx context.Context, request *CreateOpsMakePlanRequest) (response *CreateOpsMakePlanResponse, err error) {
    if request == nil {
        request = NewCreateOpsMakePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpsMakePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpsMakePlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrUpdateResourceRequest() (request *CreateOrUpdateResourceRequest) {
    request = &CreateOrUpdateResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateOrUpdateResource")
    
    
    return
}

func NewCreateOrUpdateResourceResponse() (response *CreateOrUpdateResourceResponse) {
    response = &CreateOrUpdateResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrUpdateResource
// 资源管理需要先将资源上传到cos中，然后调用该接口，将cos资源绑定到wedata
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateOrUpdateResource(request *CreateOrUpdateResourceRequest) (response *CreateOrUpdateResourceResponse, err error) {
    return c.CreateOrUpdateResourceWithContext(context.Background(), request)
}

// CreateOrUpdateResource
// 资源管理需要先将资源上传到cos中，然后调用该接口，将cos资源绑定到wedata
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateOrUpdateResourceWithContext(ctx context.Context, request *CreateOrUpdateResourceRequest) (response *CreateOrUpdateResourceResponse, err error) {
    if request == nil {
        request = NewCreateOrUpdateResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrUpdateResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrUpdateResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourcePathRequest() (request *CreateResourcePathRequest) {
    request = &CreateResourcePathRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateResourcePath")
    
    
    return
}

func NewCreateResourcePathResponse() (response *CreateResourcePathResponse) {
    response = &CreateResourcePathResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateResourcePath
// 文件路径的根目录为 /datastudio/resource，如果要在根目录下创建 aaa 文件夹，FilePath的值应该为 /datastudio/resource，如果根目录下已经创建了 aaa 文件夹，要在 aaa 下创建  bbb 文件夹，FilePath的值应该为 /datastudio/resource/aaa
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateResourcePath(request *CreateResourcePathRequest) (response *CreateResourcePathResponse, err error) {
    return c.CreateResourcePathWithContext(context.Background(), request)
}

// CreateResourcePath
// 文件路径的根目录为 /datastudio/resource，如果要在根目录下创建 aaa 文件夹，FilePath的值应该为 /datastudio/resource，如果根目录下已经创建了 aaa 文件夹，要在 aaa 下创建  bbb 文件夹，FilePath的值应该为 /datastudio/resource/aaa
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateResourcePathWithContext(ctx context.Context, request *CreateResourcePathRequest) (response *CreateResourcePathResponse, err error) {
    if request == nil {
        request = NewCreateResourcePathRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResourcePath require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourcePathResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateRule")
    
    
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRule
// 创建质量规则接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPLICATENAME = "InvalidParameter.DuplicateName"
//  INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    return c.CreateRuleWithContext(context.Background(), request)
}

// CreateRule
// 创建质量规则接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPLICATENAME = "InvalidParameter.DuplicateName"
//  INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"
func (c *Client) CreateRuleWithContext(ctx context.Context, request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleTemplateRequest() (request *CreateRuleTemplateRequest) {
    request = &CreateRuleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateRuleTemplate")
    
    
    return
}

func NewCreateRuleTemplateResponse() (response *CreateRuleTemplateResponse) {
    response = &CreateRuleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRuleTemplate
// 创建规则模版
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateRuleTemplate(request *CreateRuleTemplateRequest) (response *CreateRuleTemplateResponse, err error) {
    return c.CreateRuleTemplateWithContext(context.Background(), request)
}

// CreateRuleTemplate
// 创建规则模版
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateRuleTemplateWithContext(ctx context.Context, request *CreateRuleTemplateRequest) (response *CreateRuleTemplateResponse, err error) {
    if request == nil {
        request = NewCreateRuleTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRuleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRuleTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    return c.CreateTaskWithContext(context.Background(), request)
}

// CreateTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskAlarmRegularRequest() (request *CreateTaskAlarmRegularRequest) {
    request = &CreateTaskAlarmRegularRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateTaskAlarmRegular")
    
    
    return
}

func NewCreateTaskAlarmRegularResponse() (response *CreateTaskAlarmRegularResponse) {
    response = &CreateTaskAlarmRegularResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTaskAlarmRegular
// 创建任务告警规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTaskAlarmRegular(request *CreateTaskAlarmRegularRequest) (response *CreateTaskAlarmRegularResponse, err error) {
    return c.CreateTaskAlarmRegularWithContext(context.Background(), request)
}

// CreateTaskAlarmRegular
// 创建任务告警规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTaskAlarmRegularWithContext(ctx context.Context, request *CreateTaskAlarmRegularRequest) (response *CreateTaskAlarmRegularResponse, err error) {
    if request == nil {
        request = NewCreateTaskAlarmRegularRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskAlarmRegular require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskAlarmRegularResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskVersionDsRequest() (request *CreateTaskVersionDsRequest) {
    request = &CreateTaskVersionDsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateTaskVersionDs")
    
    
    return
}

func NewCreateTaskVersionDsResponse() (response *CreateTaskVersionDsResponse) {
    response = &CreateTaskVersionDsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTaskVersionDs
// 创建任务版本
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTaskVersionDs(request *CreateTaskVersionDsRequest) (response *CreateTaskVersionDsResponse, err error) {
    return c.CreateTaskVersionDsWithContext(context.Background(), request)
}

// CreateTaskVersionDs
// 创建任务版本
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTaskVersionDsWithContext(ctx context.Context, request *CreateTaskVersionDsRequest) (response *CreateTaskVersionDsResponse, err error) {
    if request == nil {
        request = NewCreateTaskVersionDsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskVersionDs require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskVersionDsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowRequest() (request *CreateWorkflowRequest) {
    request = &CreateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateWorkflow")
    
    
    return
}

func NewCreateWorkflowResponse() (response *CreateWorkflowResponse) {
    response = &CreateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateWorkflow(request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    return c.CreateWorkflowWithContext(context.Background(), request)
}

// CreateWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateWorkflowWithContext(ctx context.Context, request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDagInstancesRequest() (request *DagInstancesRequest) {
    request = &DagInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DagInstances")
    
    
    return
}

func NewDagInstancesResponse() (response *DagInstancesResponse) {
    response = &DagInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DagInstances
// 拉取dag实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DagInstances(request *DagInstancesRequest) (response *DagInstancesResponse, err error) {
    return c.DagInstancesWithContext(context.Background(), request)
}

// DagInstances
// 拉取dag实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DagInstancesWithContext(ctx context.Context, request *DagInstancesRequest) (response *DagInstancesResponse, err error) {
    if request == nil {
        request = NewDagInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DagInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDagInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBaselineRequest() (request *DeleteBaselineRequest) {
    request = &DeleteBaselineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteBaseline")
    
    
    return
}

func NewDeleteBaselineResponse() (response *DeleteBaselineResponse) {
    response = &DeleteBaselineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBaseline
// DeleteBaseline
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteBaseline(request *DeleteBaselineRequest) (response *DeleteBaselineResponse, err error) {
    return c.DeleteBaselineWithContext(context.Background(), request)
}

// DeleteBaseline
// DeleteBaseline
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteBaselineWithContext(ctx context.Context, request *DeleteBaselineRequest) (response *DeleteBaselineResponse, err error) {
    if request == nil {
        request = NewDeleteBaselineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBaseline require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBaselineResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomFunctionRequest() (request *DeleteCustomFunctionRequest) {
    request = &DeleteCustomFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteCustomFunction")
    
    
    return
}

func NewDeleteCustomFunctionResponse() (response *DeleteCustomFunctionResponse) {
    response = &DeleteCustomFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomFunction
// 删除用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCustomFunction(request *DeleteCustomFunctionRequest) (response *DeleteCustomFunctionResponse, err error) {
    return c.DeleteCustomFunctionWithContext(context.Background(), request)
}

// DeleteCustomFunction
// 删除用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCustomFunctionWithContext(ctx context.Context, request *DeleteCustomFunctionRequest) (response *DeleteCustomFunctionResponse, err error) {
    if request == nil {
        request = NewDeleteCustomFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDataSourcesRequest() (request *DeleteDataSourcesRequest) {
    request = &DeleteDataSourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteDataSources")
    
    
    return
}

func NewDeleteDataSourcesResponse() (response *DeleteDataSourcesResponse) {
    response = &DeleteDataSourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDataSources
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDataSources(request *DeleteDataSourcesRequest) (response *DeleteDataSourcesResponse, err error) {
    return c.DeleteDataSourcesWithContext(context.Background(), request)
}

// DeleteDataSources
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDataSourcesWithContext(ctx context.Context, request *DeleteDataSourcesRequest) (response *DeleteDataSourcesResponse, err error) {
    if request == nil {
        request = NewDeleteDataSourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataSources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDataSourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDsFolderRequest() (request *DeleteDsFolderRequest) {
    request = &DeleteDsFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteDsFolder")
    
    
    return
}

func NewDeleteDsFolderResponse() (response *DeleteDsFolderResponse) {
    response = &DeleteDsFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDsFolder
// 编排空间-删除文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDsFolder(request *DeleteDsFolderRequest) (response *DeleteDsFolderResponse, err error) {
    return c.DeleteDsFolderWithContext(context.Background(), request)
}

// DeleteDsFolder
// 编排空间-删除文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDsFolderWithContext(ctx context.Context, request *DeleteDsFolderRequest) (response *DeleteDsFolderResponse, err error) {
    if request == nil {
        request = NewDeleteDsFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDsFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDsFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileRequest() (request *DeleteFileRequest) {
    request = &DeleteFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteFile")
    
    
    return
}

func NewDeleteFileResponse() (response *DeleteFileResponse) {
    response = &DeleteFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFile
// 删除文件
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEINSTANCENOTEXISTS = "InvalidParameter.DataEngineInstanceNotExists"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteFile(request *DeleteFileRequest) (response *DeleteFileResponse, err error) {
    return c.DeleteFileWithContext(context.Background(), request)
}

// DeleteFile
// 删除文件
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEINSTANCENOTEXISTS = "InvalidParameter.DataEngineInstanceNotExists"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteFileWithContext(ctx context.Context, request *DeleteFileRequest) (response *DeleteFileResponse, err error) {
    if request == nil {
        request = NewDeleteFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFilePathRequest() (request *DeleteFilePathRequest) {
    request = &DeleteFilePathRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteFilePath")
    
    
    return
}

func NewDeleteFilePathResponse() (response *DeleteFilePathResponse) {
    response = &DeleteFilePathResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFilePath
// 开发空间-批量删除目录和文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteFilePath(request *DeleteFilePathRequest) (response *DeleteFilePathResponse, err error) {
    return c.DeleteFilePathWithContext(context.Background(), request)
}

// DeleteFilePath
// 开发空间-批量删除目录和文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteFilePathWithContext(ctx context.Context, request *DeleteFilePathRequest) (response *DeleteFilePathResponse, err error) {
    if request == nil {
        request = NewDeleteFilePathRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFilePath require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFilePathResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFolderRequest() (request *DeleteFolderRequest) {
    request = &DeleteFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteFolder")
    
    
    return
}

func NewDeleteFolderResponse() (response *DeleteFolderResponse) {
    response = &DeleteFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteFolder(request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
    return c.DeleteFolderWithContext(context.Background(), request)
}

// DeleteFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteFolderWithContext(ctx context.Context, request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
    if request == nil {
        request = NewDeleteFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInLongAgentRequest() (request *DeleteInLongAgentRequest) {
    request = &DeleteInLongAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteInLongAgent")
    
    
    return
}

func NewDeleteInLongAgentResponse() (response *DeleteInLongAgentResponse) {
    response = &DeleteInLongAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInLongAgent
// 删除采集器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteInLongAgent(request *DeleteInLongAgentRequest) (response *DeleteInLongAgentResponse, err error) {
    return c.DeleteInLongAgentWithContext(context.Background(), request)
}

// DeleteInLongAgent
// 删除采集器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteInLongAgentWithContext(ctx context.Context, request *DeleteInLongAgentRequest) (response *DeleteInLongAgentResponse, err error) {
    if request == nil {
        request = NewDeleteInLongAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInLongAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInLongAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIntegrationNodeRequest() (request *DeleteIntegrationNodeRequest) {
    request = &DeleteIntegrationNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteIntegrationNode")
    
    
    return
}

func NewDeleteIntegrationNodeResponse() (response *DeleteIntegrationNodeResponse) {
    response = &DeleteIntegrationNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIntegrationNode
// 删除集成节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIntegrationNode(request *DeleteIntegrationNodeRequest) (response *DeleteIntegrationNodeResponse, err error) {
    return c.DeleteIntegrationNodeWithContext(context.Background(), request)
}

// DeleteIntegrationNode
// 删除集成节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIntegrationNodeWithContext(ctx context.Context, request *DeleteIntegrationNodeRequest) (response *DeleteIntegrationNodeResponse, err error) {
    if request == nil {
        request = NewDeleteIntegrationNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIntegrationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIntegrationNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIntegrationTaskRequest() (request *DeleteIntegrationTaskRequest) {
    request = &DeleteIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteIntegrationTask")
    
    
    return
}

func NewDeleteIntegrationTaskResponse() (response *DeleteIntegrationTaskResponse) {
    response = &DeleteIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIntegrationTask
// 删除集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIntegrationTask(request *DeleteIntegrationTaskRequest) (response *DeleteIntegrationTaskResponse, err error) {
    return c.DeleteIntegrationTaskWithContext(context.Background(), request)
}

// DeleteIntegrationTask
// 删除集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIntegrationTaskWithContext(ctx context.Context, request *DeleteIntegrationTaskRequest) (response *DeleteIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewDeleteIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOfflineTaskRequest() (request *DeleteOfflineTaskRequest) {
    request = &DeleteOfflineTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteOfflineTask")
    
    
    return
}

func NewDeleteOfflineTaskResponse() (response *DeleteOfflineTaskResponse) {
    response = &DeleteOfflineTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOfflineTask
// 删除任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteOfflineTask(request *DeleteOfflineTaskRequest) (response *DeleteOfflineTaskResponse, err error) {
    return c.DeleteOfflineTaskWithContext(context.Background(), request)
}

// DeleteOfflineTask
// 删除任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteOfflineTaskWithContext(ctx context.Context, request *DeleteOfflineTaskRequest) (response *DeleteOfflineTaskResponse, err error) {
    if request == nil {
        request = NewDeleteOfflineTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOfflineTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOfflineTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectParamDsRequest() (request *DeleteProjectParamDsRequest) {
    request = &DeleteProjectParamDsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteProjectParamDs")
    
    
    return
}

func NewDeleteProjectParamDsResponse() (response *DeleteProjectParamDsResponse) {
    response = &DeleteProjectParamDsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProjectParamDs
// 删除项目参数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteProjectParamDs(request *DeleteProjectParamDsRequest) (response *DeleteProjectParamDsResponse, err error) {
    return c.DeleteProjectParamDsWithContext(context.Background(), request)
}

// DeleteProjectParamDs
// 删除项目参数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteProjectParamDsWithContext(ctx context.Context, request *DeleteProjectParamDsRequest) (response *DeleteProjectParamDsResponse, err error) {
    if request == nil {
        request = NewDeleteProjectParamDsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProjectParamDs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectParamDsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceRequest() (request *DeleteResourceRequest) {
    request = &DeleteResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteResource")
    
    
    return
}

func NewDeleteResourceResponse() (response *DeleteResourceResponse) {
    response = &DeleteResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResource
// 资源管理删除资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteResource(request *DeleteResourceRequest) (response *DeleteResourceResponse, err error) {
    return c.DeleteResourceWithContext(context.Background(), request)
}

// DeleteResource
// 资源管理删除资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteResourceWithContext(ctx context.Context, request *DeleteResourceRequest) (response *DeleteResourceResponse, err error) {
    if request == nil {
        request = NewDeleteResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceFileRequest() (request *DeleteResourceFileRequest) {
    request = &DeleteResourceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteResourceFile")
    
    
    return
}

func NewDeleteResourceFileResponse() (response *DeleteResourceFileResponse) {
    response = &DeleteResourceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResourceFile
// 资源管理-删除资源文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteResourceFile(request *DeleteResourceFileRequest) (response *DeleteResourceFileResponse, err error) {
    return c.DeleteResourceFileWithContext(context.Background(), request)
}

// DeleteResourceFile
// 资源管理-删除资源文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteResourceFileWithContext(ctx context.Context, request *DeleteResourceFileRequest) (response *DeleteResourceFileResponse, err error) {
    if request == nil {
        request = NewDeleteResourceFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceFilesRequest() (request *DeleteResourceFilesRequest) {
    request = &DeleteResourceFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteResourceFiles")
    
    
    return
}

func NewDeleteResourceFilesResponse() (response *DeleteResourceFilesResponse) {
    response = &DeleteResourceFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResourceFiles
// 资源管理-批量删除资源文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteResourceFiles(request *DeleteResourceFilesRequest) (response *DeleteResourceFilesResponse, err error) {
    return c.DeleteResourceFilesWithContext(context.Background(), request)
}

// DeleteResourceFiles
// 资源管理-批量删除资源文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteResourceFilesWithContext(ctx context.Context, request *DeleteResourceFilesRequest) (response *DeleteResourceFilesResponse, err error) {
    if request == nil {
        request = NewDeleteResourceFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteRule")
    
    
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRule
// 删除质量规则接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    return c.DeleteRuleWithContext(context.Background(), request)
}

// DeleteRule
// 删除质量规则接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleTemplateRequest() (request *DeleteRuleTemplateRequest) {
    request = &DeleteRuleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteRuleTemplate")
    
    
    return
}

func NewDeleteRuleTemplateResponse() (response *DeleteRuleTemplateResponse) {
    response = &DeleteRuleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRuleTemplate
// 删除规则模版
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteRuleTemplate(request *DeleteRuleTemplateRequest) (response *DeleteRuleTemplateResponse, err error) {
    return c.DeleteRuleTemplateWithContext(context.Background(), request)
}

// DeleteRuleTemplate
// 删除规则模版
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteRuleTemplateWithContext(ctx context.Context, request *DeleteRuleTemplateRequest) (response *DeleteRuleTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteRuleTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRuleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRuleTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskAlarmRegularRequest() (request *DeleteTaskAlarmRegularRequest) {
    request = &DeleteTaskAlarmRegularRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteTaskAlarmRegular")
    
    
    return
}

func NewDeleteTaskAlarmRegularResponse() (response *DeleteTaskAlarmRegularResponse) {
    response = &DeleteTaskAlarmRegularResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTaskAlarmRegular
// 删除任务告警规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTaskAlarmRegular(request *DeleteTaskAlarmRegularRequest) (response *DeleteTaskAlarmRegularResponse, err error) {
    return c.DeleteTaskAlarmRegularWithContext(context.Background(), request)
}

// DeleteTaskAlarmRegular
// 删除任务告警规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTaskAlarmRegularWithContext(ctx context.Context, request *DeleteTaskAlarmRegularRequest) (response *DeleteTaskAlarmRegularResponse, err error) {
    if request == nil {
        request = NewDeleteTaskAlarmRegularRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTaskAlarmRegular require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTaskAlarmRegularResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskDsRequest() (request *DeleteTaskDsRequest) {
    request = &DeleteTaskDsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteTaskDs")
    
    
    return
}

func NewDeleteTaskDsResponse() (response *DeleteTaskDsResponse) {
    response = &DeleteTaskDsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTaskDs
// 删除任务Ds
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteTaskDs(request *DeleteTaskDsRequest) (response *DeleteTaskDsResponse, err error) {
    return c.DeleteTaskDsWithContext(context.Background(), request)
}

// DeleteTaskDs
// 删除任务Ds
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteTaskDsWithContext(ctx context.Context, request *DeleteTaskDsRequest) (response *DeleteTaskDsResponse, err error) {
    if request == nil {
        request = NewDeleteTaskDsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTaskDs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTaskDsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowByIdRequest() (request *DeleteWorkflowByIdRequest) {
    request = &DeleteWorkflowByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteWorkflowById")
    
    
    return
}

func NewDeleteWorkflowByIdResponse() (response *DeleteWorkflowByIdResponse) {
    response = &DeleteWorkflowByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkflowById
// 通过工作流Id删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteWorkflowById(request *DeleteWorkflowByIdRequest) (response *DeleteWorkflowByIdResponse, err error) {
    return c.DeleteWorkflowByIdWithContext(context.Background(), request)
}

// DeleteWorkflowById
// 通过工作流Id删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteWorkflowByIdWithContext(ctx context.Context, request *DeleteWorkflowByIdRequest) (response *DeleteWorkflowByIdResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflowById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowNewRequest() (request *DeleteWorkflowNewRequest) {
    request = &DeleteWorkflowNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteWorkflowNew")
    
    
    return
}

func NewDeleteWorkflowNewResponse() (response *DeleteWorkflowNewResponse) {
    response = &DeleteWorkflowNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkflowNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteWorkflowNew(request *DeleteWorkflowNewRequest) (response *DeleteWorkflowNewResponse, err error) {
    return c.DeleteWorkflowNewWithContext(context.Background(), request)
}

// DeleteWorkflowNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteWorkflowNewWithContext(ctx context.Context, request *DeleteWorkflowNewRequest) (response *DeleteWorkflowNewResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflowNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmEventsRequest() (request *DescribeAlarmEventsRequest) {
    request = &DescribeAlarmEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeAlarmEvents")
    
    
    return
}

func NewDescribeAlarmEventsResponse() (response *DescribeAlarmEventsResponse) {
    response = &DescribeAlarmEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmEvents
// 告警事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmEvents(request *DescribeAlarmEventsRequest) (response *DescribeAlarmEventsResponse, err error) {
    return c.DescribeAlarmEventsWithContext(context.Background(), request)
}

// DescribeAlarmEvents
// 告警事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmEventsWithContext(ctx context.Context, request *DescribeAlarmEventsRequest) (response *DescribeAlarmEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmReceiverRequest() (request *DescribeAlarmReceiverRequest) {
    request = &DescribeAlarmReceiverRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeAlarmReceiver")
    
    
    return
}

func NewDescribeAlarmReceiverResponse() (response *DescribeAlarmReceiverResponse) {
    response = &DescribeAlarmReceiverResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmReceiver
// 告警接收人详情
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmReceiver(request *DescribeAlarmReceiverRequest) (response *DescribeAlarmReceiverResponse, err error) {
    return c.DescribeAlarmReceiverWithContext(context.Background(), request)
}

// DescribeAlarmReceiver
// 告警接收人详情
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmReceiverWithContext(ctx context.Context, request *DescribeAlarmReceiverRequest) (response *DescribeAlarmReceiverResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmReceiverRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmReceiver require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmReceiverResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllByFolderNewRequest() (request *DescribeAllByFolderNewRequest) {
    request = &DescribeAllByFolderNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeAllByFolderNew")
    
    
    return
}

func NewDescribeAllByFolderNewResponse() (response *DescribeAllByFolderNewResponse) {
    response = &DescribeAllByFolderNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllByFolderNew
// 查询父目录下所有子文件夹+工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAllByFolderNew(request *DescribeAllByFolderNewRequest) (response *DescribeAllByFolderNewResponse, err error) {
    return c.DescribeAllByFolderNewWithContext(context.Background(), request)
}

// DescribeAllByFolderNew
// 查询父目录下所有子文件夹+工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAllByFolderNewWithContext(ctx context.Context, request *DescribeAllByFolderNewRequest) (response *DescribeAllByFolderNewResponse, err error) {
    if request == nil {
        request = NewDescribeAllByFolderNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllByFolderNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllByFolderNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllTaskTypeRequest() (request *DescribeAllTaskTypeRequest) {
    request = &DescribeAllTaskTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeAllTaskType")
    
    
    return
}

func NewDescribeAllTaskTypeResponse() (response *DescribeAllTaskTypeResponse) {
    response = &DescribeAllTaskTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllTaskType
// 获取所有任务类型
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAllTaskType(request *DescribeAllTaskTypeRequest) (response *DescribeAllTaskTypeResponse, err error) {
    return c.DescribeAllTaskTypeWithContext(context.Background(), request)
}

// DescribeAllTaskType
// 获取所有任务类型
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAllTaskTypeWithContext(ctx context.Context, request *DescribeAllTaskTypeRequest) (response *DescribeAllTaskTypeResponse, err error) {
    if request == nil {
        request = NewDescribeAllTaskTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllTaskType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllTaskTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllUsedVersionSonRequest() (request *DescribeAllUsedVersionSonRequest) {
    request = &DescribeAllUsedVersionSonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeAllUsedVersionSon")
    
    
    return
}

func NewDescribeAllUsedVersionSonResponse() (response *DescribeAllUsedVersionSonResponse) {
    response = &DescribeAllUsedVersionSonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllUsedVersionSon
// 根据任务Id查找生产态子任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAllUsedVersionSon(request *DescribeAllUsedVersionSonRequest) (response *DescribeAllUsedVersionSonResponse, err error) {
    return c.DescribeAllUsedVersionSonWithContext(context.Background(), request)
}

// DescribeAllUsedVersionSon
// 根据任务Id查找生产态子任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAllUsedVersionSonWithContext(ctx context.Context, request *DescribeAllUsedVersionSonRequest) (response *DescribeAllUsedVersionSonResponse, err error) {
    if request == nil {
        request = NewDescribeAllUsedVersionSonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllUsedVersionSon require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllUsedVersionSonResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApproveListRequest() (request *DescribeApproveListRequest) {
    request = &DescribeApproveListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeApproveList")
    
    
    return
}

func NewDescribeApproveListResponse() (response *DescribeApproveListResponse) {
    response = &DescribeApproveListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApproveList
// 获取待审批列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeApproveList(request *DescribeApproveListRequest) (response *DescribeApproveListResponse, err error) {
    return c.DescribeApproveListWithContext(context.Background(), request)
}

// DescribeApproveList
// 获取待审批列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeApproveListWithContext(ctx context.Context, request *DescribeApproveListRequest) (response *DescribeApproveListResponse, err error) {
    if request == nil {
        request = NewDescribeApproveListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApproveList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApproveListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApproveTypeListRequest() (request *DescribeApproveTypeListRequest) {
    request = &DescribeApproveTypeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeApproveTypeList")
    
    
    return
}

func NewDescribeApproveTypeListResponse() (response *DescribeApproveTypeListResponse) {
    response = &DescribeApproveTypeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApproveTypeList
// 获取审批分类列表
//
// 可能返回的错误码:
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeApproveTypeList(request *DescribeApproveTypeListRequest) (response *DescribeApproveTypeListResponse, err error) {
    return c.DescribeApproveTypeListWithContext(context.Background(), request)
}

// DescribeApproveTypeList
// 获取审批分类列表
//
// 可能返回的错误码:
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeApproveTypeListWithContext(ctx context.Context, request *DescribeApproveTypeListRequest) (response *DescribeApproveTypeListResponse, err error) {
    if request == nil {
        request = NewDescribeApproveTypeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApproveTypeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApproveTypeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineAllTaskDagRequest() (request *DescribeBaselineAllTaskDagRequest) {
    request = &DescribeBaselineAllTaskDagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeBaselineAllTaskDag")
    
    
    return
}

func NewDescribeBaselineAllTaskDagResponse() (response *DescribeBaselineAllTaskDagResponse) {
    response = &DescribeBaselineAllTaskDagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaselineAllTaskDag
// 查询基线DAG
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineAllTaskDag(request *DescribeBaselineAllTaskDagRequest) (response *DescribeBaselineAllTaskDagResponse, err error) {
    return c.DescribeBaselineAllTaskDagWithContext(context.Background(), request)
}

// DescribeBaselineAllTaskDag
// 查询基线DAG
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineAllTaskDagWithContext(ctx context.Context, request *DescribeBaselineAllTaskDagRequest) (response *DescribeBaselineAllTaskDagResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineAllTaskDagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineAllTaskDag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineAllTaskDagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineByIdRequest() (request *DescribeBaselineByIdRequest) {
    request = &DescribeBaselineByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeBaselineById")
    
    
    return
}

func NewDescribeBaselineByIdResponse() (response *DescribeBaselineByIdResponse) {
    response = &DescribeBaselineByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaselineById
// 基线列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineById(request *DescribeBaselineByIdRequest) (response *DescribeBaselineByIdResponse, err error) {
    return c.DescribeBaselineByIdWithContext(context.Background(), request)
}

// DescribeBaselineById
// 基线列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineByIdWithContext(ctx context.Context, request *DescribeBaselineByIdRequest) (response *DescribeBaselineByIdResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineInstanceDagRequest() (request *DescribeBaselineInstanceDagRequest) {
    request = &DescribeBaselineInstanceDagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeBaselineInstanceDag")
    
    
    return
}

func NewDescribeBaselineInstanceDagResponse() (response *DescribeBaselineInstanceDagResponse) {
    response = &DescribeBaselineInstanceDagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaselineInstanceDag
// 查询基线实例DAG
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineInstanceDag(request *DescribeBaselineInstanceDagRequest) (response *DescribeBaselineInstanceDagResponse, err error) {
    return c.DescribeBaselineInstanceDagWithContext(context.Background(), request)
}

// DescribeBaselineInstanceDag
// 查询基线实例DAG
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineInstanceDagWithContext(ctx context.Context, request *DescribeBaselineInstanceDagRequest) (response *DescribeBaselineInstanceDagResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineInstanceDagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineInstanceDag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineInstanceDagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineInstanceGanttRequest() (request *DescribeBaselineInstanceGanttRequest) {
    request = &DescribeBaselineInstanceGanttRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeBaselineInstanceGantt")
    
    
    return
}

func NewDescribeBaselineInstanceGanttResponse() (response *DescribeBaselineInstanceGanttResponse) {
    response = &DescribeBaselineInstanceGanttResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaselineInstanceGantt
// 查询基线实例关键任务实例甘特图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineInstanceGantt(request *DescribeBaselineInstanceGanttRequest) (response *DescribeBaselineInstanceGanttResponse, err error) {
    return c.DescribeBaselineInstanceGanttWithContext(context.Background(), request)
}

// DescribeBaselineInstanceGantt
// 查询基线实例关键任务实例甘特图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselineInstanceGanttWithContext(ctx context.Context, request *DescribeBaselineInstanceGanttRequest) (response *DescribeBaselineInstanceGanttResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineInstanceGanttRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineInstanceGantt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineInstanceGanttResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselineInstancesRequest() (request *DescribeBaselineInstancesRequest) {
    request = &DescribeBaselineInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeBaselineInstances")
    
    
    return
}

func NewDescribeBaselineInstancesResponse() (response *DescribeBaselineInstancesResponse) {
    response = &DescribeBaselineInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaselineInstances
// 查询基线实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBaselineInstances(request *DescribeBaselineInstancesRequest) (response *DescribeBaselineInstancesResponse, err error) {
    return c.DescribeBaselineInstancesWithContext(context.Background(), request)
}

// DescribeBaselineInstances
// 查询基线实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBaselineInstancesWithContext(ctx context.Context, request *DescribeBaselineInstancesRequest) (response *DescribeBaselineInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeBaselineInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselineInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselineInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaselinesRequest() (request *DescribeBaselinesRequest) {
    request = &DescribeBaselinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeBaselines")
    
    
    return
}

func NewDescribeBaselinesResponse() (response *DescribeBaselinesResponse) {
    response = &DescribeBaselinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaselines
// 基线列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselines(request *DescribeBaselinesRequest) (response *DescribeBaselinesResponse, err error) {
    return c.DescribeBaselinesWithContext(context.Background(), request)
}

// DescribeBaselines
// 基线列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBaselinesWithContext(ctx context.Context, request *DescribeBaselinesRequest) (response *DescribeBaselinesResponse, err error) {
    if request == nil {
        request = NewDescribeBaselinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaselines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaselinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchOperateTaskRequest() (request *DescribeBatchOperateTaskRequest) {
    request = &DescribeBatchOperateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeBatchOperateTask")
    
    
    return
}

func NewDescribeBatchOperateTaskResponse() (response *DescribeBatchOperateTaskResponse) {
    response = &DescribeBatchOperateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBatchOperateTask
// 批量操作任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBatchOperateTask(request *DescribeBatchOperateTaskRequest) (response *DescribeBatchOperateTaskResponse, err error) {
    return c.DescribeBatchOperateTaskWithContext(context.Background(), request)
}

// DescribeBatchOperateTask
// 批量操作任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBatchOperateTaskWithContext(ctx context.Context, request *DescribeBatchOperateTaskRequest) (response *DescribeBatchOperateTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBatchOperateTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchOperateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchOperateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBelongToRequest() (request *DescribeBelongToRequest) {
    request = &DescribeBelongToRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeBelongTo")
    
    
    return
}

func NewDescribeBelongToResponse() (response *DescribeBelongToResponse) {
    response = &DescribeBelongToResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBelongTo
// 智能运维-事件列表-所属任务/基线过滤列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBelongTo(request *DescribeBelongToRequest) (response *DescribeBelongToResponse, err error) {
    return c.DescribeBelongToWithContext(context.Background(), request)
}

// DescribeBelongTo
// 智能运维-事件列表-所属任务/基线过滤列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBelongToWithContext(ctx context.Context, request *DescribeBelongToRequest) (response *DescribeBelongToResponse, err error) {
    if request == nil {
        request = NewDescribeBelongToRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBelongTo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBelongToResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterNamespaceListRequest() (request *DescribeClusterNamespaceListRequest) {
    request = &DescribeClusterNamespaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeClusterNamespaceList")
    
    
    return
}

func NewDescribeClusterNamespaceListResponse() (response *DescribeClusterNamespaceListResponse) {
    response = &DescribeClusterNamespaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterNamespaceList
// 获取集群命名空间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeClusterNamespaceList(request *DescribeClusterNamespaceListRequest) (response *DescribeClusterNamespaceListResponse, err error) {
    return c.DescribeClusterNamespaceListWithContext(context.Background(), request)
}

// DescribeClusterNamespaceList
// 获取集群命名空间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeClusterNamespaceListWithContext(ctx context.Context, request *DescribeClusterNamespaceListRequest) (response *DescribeClusterNamespaceListResponse, err error) {
    if request == nil {
        request = NewDescribeClusterNamespaceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterNamespaceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterNamespaceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeColumnLineageRequest() (request *DescribeColumnLineageRequest) {
    request = &DescribeColumnLineageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeColumnLineage")
    
    
    return
}

func NewDescribeColumnLineageResponse() (response *DescribeColumnLineageResponse) {
    response = &DescribeColumnLineageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeColumnLineage
// 列出字段血缘信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeColumnLineage(request *DescribeColumnLineageRequest) (response *DescribeColumnLineageResponse, err error) {
    return c.DescribeColumnLineageWithContext(context.Background(), request)
}

// DescribeColumnLineage
// 列出字段血缘信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeColumnLineageWithContext(ctx context.Context, request *DescribeColumnLineageRequest) (response *DescribeColumnLineageResponse, err error) {
    if request == nil {
        request = NewDescribeColumnLineageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeColumnLineage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeColumnLineageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeColumnsMetaRequest() (request *DescribeColumnsMetaRequest) {
    request = &DescribeColumnsMetaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeColumnsMeta")
    
    
    return
}

func NewDescribeColumnsMetaResponse() (response *DescribeColumnsMetaResponse) {
    response = &DescribeColumnsMetaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeColumnsMeta
// 查询表的所有列元数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeColumnsMeta(request *DescribeColumnsMetaRequest) (response *DescribeColumnsMetaResponse, err error) {
    return c.DescribeColumnsMetaWithContext(context.Background(), request)
}

// DescribeColumnsMeta
// 查询表的所有列元数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeColumnsMetaWithContext(ctx context.Context, request *DescribeColumnsMetaRequest) (response *DescribeColumnsMetaResponse, err error) {
    if request == nil {
        request = NewDescribeColumnsMetaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeColumnsMeta require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeColumnsMetaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataBasesRequest() (request *DescribeDataBasesRequest) {
    request = &DescribeDataBasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDataBases")
    
    
    return
}

func NewDescribeDataBasesResponse() (response *DescribeDataBasesResponse) {
    response = &DescribeDataBasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataBases
// 查询数据来源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeDataBases(request *DescribeDataBasesRequest) (response *DescribeDataBasesResponse, err error) {
    return c.DescribeDataBasesWithContext(context.Background(), request)
}

// DescribeDataBases
// 查询数据来源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeDataBasesWithContext(ctx context.Context, request *DescribeDataBasesRequest) (response *DescribeDataBasesResponse, err error) {
    if request == nil {
        request = NewDescribeDataBasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataBases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataBasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataCheckStatRequest() (request *DescribeDataCheckStatRequest) {
    request = &DescribeDataCheckStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDataCheckStat")
    
    
    return
}

func NewDescribeDataCheckStatResponse() (response *DescribeDataCheckStatResponse) {
    response = &DescribeDataCheckStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataCheckStat
// 数据质量的概览页面数据监测情况接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeDataCheckStat(request *DescribeDataCheckStatRequest) (response *DescribeDataCheckStatResponse, err error) {
    return c.DescribeDataCheckStatWithContext(context.Background(), request)
}

// DescribeDataCheckStat
// 数据质量的概览页面数据监测情况接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeDataCheckStatWithContext(ctx context.Context, request *DescribeDataCheckStatRequest) (response *DescribeDataCheckStatResponse, err error) {
    if request == nil {
        request = NewDescribeDataCheckStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataCheckStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataCheckStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataObjectsRequest() (request *DescribeDataObjectsRequest) {
    request = &DescribeDataObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDataObjects")
    
    
    return
}

func NewDescribeDataObjectsResponse() (response *DescribeDataObjectsResponse) {
    response = &DescribeDataObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataObjects
// 查询规则组数据对象列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeDataObjects(request *DescribeDataObjectsRequest) (response *DescribeDataObjectsResponse, err error) {
    return c.DescribeDataObjectsWithContext(context.Background(), request)
}

// DescribeDataObjects
// 查询规则组数据对象列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeDataObjectsWithContext(ctx context.Context, request *DescribeDataObjectsRequest) (response *DescribeDataObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeDataObjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataObjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataSourceInfoListRequest() (request *DescribeDataSourceInfoListRequest) {
    request = &DescribeDataSourceInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDataSourceInfoList")
    
    
    return
}

func NewDescribeDataSourceInfoListResponse() (response *DescribeDataSourceInfoListResponse) {
    response = &DescribeDataSourceInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataSourceInfoList
// 获取数据源信息-数据源分页列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDataSourceInfoList(request *DescribeDataSourceInfoListRequest) (response *DescribeDataSourceInfoListResponse, err error) {
    return c.DescribeDataSourceInfoListWithContext(context.Background(), request)
}

// DescribeDataSourceInfoList
// 获取数据源信息-数据源分页列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDataSourceInfoListWithContext(ctx context.Context, request *DescribeDataSourceInfoListRequest) (response *DescribeDataSourceInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeDataSourceInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataSourceInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataSourceInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataSourceListRequest() (request *DescribeDataSourceListRequest) {
    request = &DescribeDataSourceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDataSourceList")
    
    
    return
}

func NewDescribeDataSourceListResponse() (response *DescribeDataSourceListResponse) {
    response = &DescribeDataSourceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataSourceList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeDataSourceList(request *DescribeDataSourceListRequest) (response *DescribeDataSourceListResponse, err error) {
    return c.DescribeDataSourceListWithContext(context.Background(), request)
}

// DescribeDataSourceList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeDataSourceListWithContext(ctx context.Context, request *DescribeDataSourceListRequest) (response *DescribeDataSourceListResponse, err error) {
    if request == nil {
        request = NewDescribeDataSourceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataSourceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataSourceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataSourceWithoutInfoRequest() (request *DescribeDataSourceWithoutInfoRequest) {
    request = &DescribeDataSourceWithoutInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDataSourceWithoutInfo")
    
    
    return
}

func NewDescribeDataSourceWithoutInfoResponse() (response *DescribeDataSourceWithoutInfoResponse) {
    response = &DescribeDataSourceWithoutInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataSourceWithoutInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
func (c *Client) DescribeDataSourceWithoutInfo(request *DescribeDataSourceWithoutInfoRequest) (response *DescribeDataSourceWithoutInfoResponse, err error) {
    return c.DescribeDataSourceWithoutInfoWithContext(context.Background(), request)
}

// DescribeDataSourceWithoutInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
func (c *Client) DescribeDataSourceWithoutInfoWithContext(ctx context.Context, request *DescribeDataSourceWithoutInfoRequest) (response *DescribeDataSourceWithoutInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataSourceWithoutInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataSourceWithoutInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataSourceWithoutInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataTypesRequest() (request *DescribeDataTypesRequest) {
    request = &DescribeDataTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDataTypes")
    
    
    return
}

func NewDescribeDataTypesResponse() (response *DescribeDataTypesResponse) {
    response = &DescribeDataTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataTypes
// 获取字段类型列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataTypes(request *DescribeDataTypesRequest) (response *DescribeDataTypesResponse, err error) {
    return c.DescribeDataTypesWithContext(context.Background(), request)
}

// DescribeDataTypes
// 获取字段类型列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataTypesWithContext(ctx context.Context, request *DescribeDataTypesRequest) (response *DescribeDataTypesResponse, err error) {
    if request == nil {
        request = NewDescribeDataTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseInfoListRequest() (request *DescribeDatabaseInfoListRequest) {
    request = &DescribeDatabaseInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDatabaseInfoList")
    
    
    return
}

func NewDescribeDatabaseInfoListResponse() (response *DescribeDatabaseInfoListResponse) {
    response = &DescribeDatabaseInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseInfoList
// 获取数据库信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDatabaseInfoList(request *DescribeDatabaseInfoListRequest) (response *DescribeDatabaseInfoListResponse, err error) {
    return c.DescribeDatabaseInfoListWithContext(context.Background(), request)
}

// DescribeDatabaseInfoList
// 获取数据库信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDatabaseInfoListWithContext(ctx context.Context, request *DescribeDatabaseInfoListRequest) (response *DescribeDatabaseInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseMetasRequest() (request *DescribeDatabaseMetasRequest) {
    request = &DescribeDatabaseMetasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDatabaseMetas")
    
    
    return
}

func NewDescribeDatabaseMetasResponse() (response *DescribeDatabaseMetasResponse) {
    response = &DescribeDatabaseMetasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseMetas
// 查询数据库列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDatabaseMetas(request *DescribeDatabaseMetasRequest) (response *DescribeDatabaseMetasResponse, err error) {
    return c.DescribeDatabaseMetasWithContext(context.Background(), request)
}

// DescribeDatabaseMetas
// 查询数据库列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDatabaseMetasWithContext(ctx context.Context, request *DescribeDatabaseMetasRequest) (response *DescribeDatabaseMetasResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseMetasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseMetas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseMetasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatasourceRequest() (request *DescribeDatasourceRequest) {
    request = &DescribeDatasourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDatasource")
    
    
    return
}

func NewDescribeDatasourceResponse() (response *DescribeDatasourceResponse) {
    response = &DescribeDatasourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatasource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDatasource(request *DescribeDatasourceRequest) (response *DescribeDatasourceResponse, err error) {
    return c.DescribeDatasourceWithContext(context.Background(), request)
}

// DescribeDatasource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 数据源详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDatasourceWithContext(ctx context.Context, request *DescribeDatasourceRequest) (response *DescribeDatasourceResponse, err error) {
    if request == nil {
        request = NewDescribeDatasourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDependOpsTaskListRequest() (request *DescribeDependOpsTaskListRequest) {
    request = &DescribeDependOpsTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDependOpsTaskList")
    
    
    return
}

func NewDescribeDependOpsTaskListResponse() (response *DescribeDependOpsTaskListResponse) {
    response = &DescribeDependOpsTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDependOpsTaskList
// 根据任务id获取下游依赖任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDependOpsTaskList(request *DescribeDependOpsTaskListRequest) (response *DescribeDependOpsTaskListResponse, err error) {
    return c.DescribeDependOpsTaskListWithContext(context.Background(), request)
}

// DescribeDependOpsTaskList
// 根据任务id获取下游依赖任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDependOpsTaskListWithContext(ctx context.Context, request *DescribeDependOpsTaskListRequest) (response *DescribeDependOpsTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeDependOpsTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDependOpsTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDependOpsTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDependOpsTasksRequest() (request *DescribeDependOpsTasksRequest) {
    request = &DescribeDependOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDependOpsTasks")
    
    
    return
}

func NewDescribeDependOpsTasksResponse() (response *DescribeDependOpsTasksResponse) {
    response = &DescribeDependOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDependOpsTasks
// 根据层级查找上/下游任务节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDependOpsTasks(request *DescribeDependOpsTasksRequest) (response *DescribeDependOpsTasksResponse, err error) {
    return c.DescribeDependOpsTasksWithContext(context.Background(), request)
}

// DescribeDependOpsTasks
// 根据层级查找上/下游任务节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDependOpsTasksWithContext(ctx context.Context, request *DescribeDependOpsTasksRequest) (response *DescribeDependOpsTasksResponse, err error) {
    if request == nil {
        request = NewDescribeDependOpsTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDependOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDependOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDependTaskListsRequest() (request *DescribeDependTaskListsRequest) {
    request = &DescribeDependTaskListsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDependTaskLists")
    
    
    return
}

func NewDescribeDependTaskListsResponse() (response *DescribeDependTaskListsResponse) {
    response = &DescribeDependTaskListsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDependTaskLists
// 通过taskIds查询task详情列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDependTaskLists(request *DescribeDependTaskListsRequest) (response *DescribeDependTaskListsResponse, err error) {
    return c.DescribeDependTaskListsWithContext(context.Background(), request)
}

// DescribeDependTaskLists
// 通过taskIds查询task详情列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDependTaskListsWithContext(ctx context.Context, request *DescribeDependTaskListsRequest) (response *DescribeDependTaskListsResponse, err error) {
    if request == nil {
        request = NewDescribeDependTaskListsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDependTaskLists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDependTaskListsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDependTasksNewRequest() (request *DescribeDependTasksNewRequest) {
    request = &DescribeDependTasksNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDependTasksNew")
    
    
    return
}

func NewDescribeDependTasksNewResponse() (response *DescribeDependTasksNewResponse) {
    response = &DescribeDependTasksNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDependTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 根据层级查找上/下游任务节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDependTasksNew(request *DescribeDependTasksNewRequest) (response *DescribeDependTasksNewResponse, err error) {
    return c.DescribeDependTasksNewWithContext(context.Background(), request)
}

// DescribeDependTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 根据层级查找上/下游任务节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDependTasksNewWithContext(ctx context.Context, request *DescribeDependTasksNewRequest) (response *DescribeDependTasksNewResponse, err error) {
    if request == nil {
        request = NewDescribeDependTasksNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDependTasksNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDependTasksNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiagnosticInfoByBaselineIdRequest() (request *DescribeDiagnosticInfoByBaselineIdRequest) {
    request = &DescribeDiagnosticInfoByBaselineIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDiagnosticInfoByBaselineId")
    
    
    return
}

func NewDescribeDiagnosticInfoByBaselineIdResponse() (response *DescribeDiagnosticInfoByBaselineIdResponse) {
    response = &DescribeDiagnosticInfoByBaselineIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDiagnosticInfoByBaselineId
// 查询基线诊断信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDiagnosticInfoByBaselineId(request *DescribeDiagnosticInfoByBaselineIdRequest) (response *DescribeDiagnosticInfoByBaselineIdResponse, err error) {
    return c.DescribeDiagnosticInfoByBaselineIdWithContext(context.Background(), request)
}

// DescribeDiagnosticInfoByBaselineId
// 查询基线诊断信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDiagnosticInfoByBaselineIdWithContext(ctx context.Context, request *DescribeDiagnosticInfoByBaselineIdRequest) (response *DescribeDiagnosticInfoByBaselineIdResponse, err error) {
    if request == nil {
        request = NewDescribeDiagnosticInfoByBaselineIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiagnosticInfoByBaselineId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiagnosticInfoByBaselineIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDimensionScoreRequest() (request *DescribeDimensionScoreRequest) {
    request = &DescribeDimensionScoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDimensionScore")
    
    
    return
}

func NewDescribeDimensionScoreResponse() (response *DescribeDimensionScoreResponse) {
    response = &DescribeDimensionScoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDimensionScore
// 质量报告-查询质量评分
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDimensionScore(request *DescribeDimensionScoreRequest) (response *DescribeDimensionScoreResponse, err error) {
    return c.DescribeDimensionScoreWithContext(context.Background(), request)
}

// DescribeDimensionScore
// 质量报告-查询质量评分
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDimensionScoreWithContext(ctx context.Context, request *DescribeDimensionScoreRequest) (response *DescribeDimensionScoreResponse, err error) {
    if request == nil {
        request = NewDescribeDimensionScoreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDimensionScore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDimensionScoreResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDrInstancePageRequest() (request *DescribeDrInstancePageRequest) {
    request = &DescribeDrInstancePageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDrInstancePage")
    
    
    return
}

func NewDescribeDrInstancePageResponse() (response *DescribeDrInstancePageResponse) {
    response = &DescribeDrInstancePageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDrInstancePage
// 分页查询试运行实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDrInstancePage(request *DescribeDrInstancePageRequest) (response *DescribeDrInstancePageResponse, err error) {
    return c.DescribeDrInstancePageWithContext(context.Background(), request)
}

// DescribeDrInstancePage
// 分页查询试运行实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDrInstancePageWithContext(ctx context.Context, request *DescribeDrInstancePageRequest) (response *DescribeDrInstancePageResponse, err error) {
    if request == nil {
        request = NewDescribeDrInstancePageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDrInstancePage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDrInstancePageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDrInstanceScriptContentRequest() (request *DescribeDrInstanceScriptContentRequest) {
    request = &DescribeDrInstanceScriptContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDrInstanceScriptContent")
    
    
    return
}

func NewDescribeDrInstanceScriptContentResponse() (response *DescribeDrInstanceScriptContentResponse) {
    response = &DescribeDrInstanceScriptContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDrInstanceScriptContent
// 查询试运行实例执行内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDrInstanceScriptContent(request *DescribeDrInstanceScriptContentRequest) (response *DescribeDrInstanceScriptContentResponse, err error) {
    return c.DescribeDrInstanceScriptContentWithContext(context.Background(), request)
}

// DescribeDrInstanceScriptContent
// 查询试运行实例执行内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDrInstanceScriptContentWithContext(ctx context.Context, request *DescribeDrInstanceScriptContentRequest) (response *DescribeDrInstanceScriptContentResponse, err error) {
    if request == nil {
        request = NewDescribeDrInstanceScriptContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDrInstanceScriptContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDrInstanceScriptContentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDrSonInstanceRequest() (request *DescribeDrSonInstanceRequest) {
    request = &DescribeDrSonInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDrSonInstance")
    
    
    return
}

func NewDescribeDrSonInstanceResponse() (response *DescribeDrSonInstanceResponse) {
    response = &DescribeDrSonInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDrSonInstance
// 查询试运行实例子实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDrSonInstance(request *DescribeDrSonInstanceRequest) (response *DescribeDrSonInstanceResponse, err error) {
    return c.DescribeDrSonInstanceWithContext(context.Background(), request)
}

// DescribeDrSonInstance
// 查询试运行实例子实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDrSonInstanceWithContext(ctx context.Context, request *DescribeDrSonInstanceRequest) (response *DescribeDrSonInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeDrSonInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDrSonInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDrSonInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDsFolderTreeRequest() (request *DescribeDsFolderTreeRequest) {
    request = &DescribeDsFolderTreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDsFolderTree")
    
    
    return
}

func NewDescribeDsFolderTreeResponse() (response *DescribeDsFolderTreeResponse) {
    response = &DescribeDsFolderTreeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDsFolderTree
// 查询目录树
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDsFolderTree(request *DescribeDsFolderTreeRequest) (response *DescribeDsFolderTreeResponse, err error) {
    return c.DescribeDsFolderTreeWithContext(context.Background(), request)
}

// DescribeDsFolderTree
// 查询目录树
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDsFolderTreeWithContext(ctx context.Context, request *DescribeDsFolderTreeRequest) (response *DescribeDsFolderTreeResponse, err error) {
    if request == nil {
        request = NewDescribeDsFolderTreeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDsFolderTree require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDsFolderTreeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDsParentFolderTreeRequest() (request *DescribeDsParentFolderTreeRequest) {
    request = &DescribeDsParentFolderTreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDsParentFolderTree")
    
    
    return
}

func NewDescribeDsParentFolderTreeResponse() (response *DescribeDsParentFolderTreeResponse) {
    response = &DescribeDsParentFolderTreeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDsParentFolderTree
// 查询父目录树，用于工作流、任务定位
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDsParentFolderTree(request *DescribeDsParentFolderTreeRequest) (response *DescribeDsParentFolderTreeResponse, err error) {
    return c.DescribeDsParentFolderTreeWithContext(context.Background(), request)
}

// DescribeDsParentFolderTree
// 查询父目录树，用于工作流、任务定位
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDsParentFolderTreeWithContext(ctx context.Context, request *DescribeDsParentFolderTreeRequest) (response *DescribeDsParentFolderTreeResponse, err error) {
    if request == nil {
        request = NewDescribeDsParentFolderTreeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDsParentFolderTree require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDsParentFolderTreeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventRequest() (request *DescribeEventRequest) {
    request = &DescribeEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeEvent")
    
    
    return
}

func NewDescribeEventResponse() (response *DescribeEventResponse) {
    response = &DescribeEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEvent
// 根据项目ID和事件名称查看事件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEvent(request *DescribeEventRequest) (response *DescribeEventResponse, err error) {
    return c.DescribeEventWithContext(context.Background(), request)
}

// DescribeEvent
// 根据项目ID和事件名称查看事件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventWithContext(ctx context.Context, request *DescribeEventRequest) (response *DescribeEventResponse, err error) {
    if request == nil {
        request = NewDescribeEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventCasesRequest() (request *DescribeEventCasesRequest) {
    request = &DescribeEventCasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeEventCases")
    
    
    return
}

func NewDescribeEventCasesResponse() (response *DescribeEventCasesResponse) {
    response = &DescribeEventCasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEventCases
// 根据条件查找事件实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventCases(request *DescribeEventCasesRequest) (response *DescribeEventCasesResponse, err error) {
    return c.DescribeEventCasesWithContext(context.Background(), request)
}

// DescribeEventCases
// 根据条件查找事件实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventCasesWithContext(ctx context.Context, request *DescribeEventCasesRequest) (response *DescribeEventCasesResponse, err error) {
    if request == nil {
        request = NewDescribeEventCasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventCases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventCasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventConsumeTasksRequest() (request *DescribeEventConsumeTasksRequest) {
    request = &DescribeEventConsumeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeEventConsumeTasks")
    
    
    return
}

func NewDescribeEventConsumeTasksResponse() (response *DescribeEventConsumeTasksResponse) {
    response = &DescribeEventConsumeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEventConsumeTasks
// 查看事件实例的消费任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventConsumeTasks(request *DescribeEventConsumeTasksRequest) (response *DescribeEventConsumeTasksResponse, err error) {
    return c.DescribeEventConsumeTasksWithContext(context.Background(), request)
}

// DescribeEventConsumeTasks
// 查看事件实例的消费任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventConsumeTasksWithContext(ctx context.Context, request *DescribeEventConsumeTasksRequest) (response *DescribeEventConsumeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeEventConsumeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventConsumeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventConsumeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventDetailRequest() (request *DescribeEventDetailRequest) {
    request = &DescribeEventDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeEventDetail")
    
    
    return
}

func NewDescribeEventDetailResponse() (response *DescribeEventDetailResponse) {
    response = &DescribeEventDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEventDetail
// 智能运维事件详情1
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventDetail(request *DescribeEventDetailRequest) (response *DescribeEventDetailResponse, err error) {
    return c.DescribeEventDetailWithContext(context.Background(), request)
}

// DescribeEventDetail
// 智能运维事件详情1
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventDetailWithContext(ctx context.Context, request *DescribeEventDetailRequest) (response *DescribeEventDetailResponse, err error) {
    if request == nil {
        request = NewDescribeEventDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventIsAlarmTypesRequest() (request *DescribeEventIsAlarmTypesRequest) {
    request = &DescribeEventIsAlarmTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeEventIsAlarmTypes")
    
    
    return
}

func NewDescribeEventIsAlarmTypesResponse() (response *DescribeEventIsAlarmTypesResponse) {
    response = &DescribeEventIsAlarmTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEventIsAlarmTypes
// 事件是否告警过滤条件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventIsAlarmTypes(request *DescribeEventIsAlarmTypesRequest) (response *DescribeEventIsAlarmTypesResponse, err error) {
    return c.DescribeEventIsAlarmTypesWithContext(context.Background(), request)
}

// DescribeEventIsAlarmTypes
// 事件是否告警过滤条件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventIsAlarmTypesWithContext(ctx context.Context, request *DescribeEventIsAlarmTypesRequest) (response *DescribeEventIsAlarmTypesResponse, err error) {
    if request == nil {
        request = NewDescribeEventIsAlarmTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventIsAlarmTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventIsAlarmTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventTypesRequest() (request *DescribeEventTypesRequest) {
    request = &DescribeEventTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeEventTypes")
    
    
    return
}

func NewDescribeEventTypesResponse() (response *DescribeEventTypesResponse) {
    response = &DescribeEventTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEventTypes
// 查询事件类型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventTypes(request *DescribeEventTypesRequest) (response *DescribeEventTypesResponse, err error) {
    return c.DescribeEventTypesWithContext(context.Background(), request)
}

// DescribeEventTypes
// 查询事件类型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventTypesWithContext(ctx context.Context, request *DescribeEventTypesRequest) (response *DescribeEventTypesResponse, err error) {
    if request == nil {
        request = NewDescribeEventTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventsRequest() (request *DescribeEventsRequest) {
    request = &DescribeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeEvents")
    
    
    return
}

func NewDescribeEventsResponse() (response *DescribeEventsResponse) {
    response = &DescribeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEvents
// 智能运维事件查询列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    return c.DescribeEventsWithContext(context.Background(), request)
}

// DescribeEvents
// 智能运维事件查询列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    if request == nil {
        request = NewDescribeEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExecStrategyRequest() (request *DescribeExecStrategyRequest) {
    request = &DescribeExecStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeExecStrategy")
    
    
    return
}

func NewDescribeExecStrategyResponse() (response *DescribeExecStrategyResponse) {
    response = &DescribeExecStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExecStrategy
// 查询规则组执行策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeExecStrategy(request *DescribeExecStrategyRequest) (response *DescribeExecStrategyResponse, err error) {
    return c.DescribeExecStrategyWithContext(context.Background(), request)
}

// DescribeExecStrategy
// 查询规则组执行策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeExecStrategyWithContext(ctx context.Context, request *DescribeExecStrategyRequest) (response *DescribeExecStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeExecStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExecStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExecStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFathersRequest() (request *DescribeFathersRequest) {
    request = &DescribeFathersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeFathers")
    
    
    return
}

func NewDescribeFathersResponse() (response *DescribeFathersResponse) {
    response = &DescribeFathersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFathers
// 获取关联父实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFathers(request *DescribeFathersRequest) (response *DescribeFathersResponse, err error) {
    return c.DescribeFathersWithContext(context.Background(), request)
}

// DescribeFathers
// 获取关联父实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFathersWithContext(ctx context.Context, request *DescribeFathersRequest) (response *DescribeFathersResponse, err error) {
    if request == nil {
        request = NewDescribeFathersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFathers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFathersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFolderListRequest() (request *DescribeFolderListRequest) {
    request = &DescribeFolderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeFolderList")
    
    
    return
}

func NewDescribeFolderListResponse() (response *DescribeFolderListResponse) {
    response = &DescribeFolderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFolderList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 拉取文件夹目录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFolderList(request *DescribeFolderListRequest) (response *DescribeFolderListResponse, err error) {
    return c.DescribeFolderListWithContext(context.Background(), request)
}

// DescribeFolderList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 拉取文件夹目录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFolderListWithContext(ctx context.Context, request *DescribeFolderListRequest) (response *DescribeFolderListResponse, err error) {
    if request == nil {
        request = NewDescribeFolderListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFolderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFolderListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFolderWorkflowListRequest() (request *DescribeFolderWorkflowListRequest) {
    request = &DescribeFolderWorkflowListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeFolderWorkflowList")
    
    
    return
}

func NewDescribeFolderWorkflowListResponse() (response *DescribeFolderWorkflowListResponse) {
    response = &DescribeFolderWorkflowListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFolderWorkflowList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 拉取文件夹下的工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFolderWorkflowList(request *DescribeFolderWorkflowListRequest) (response *DescribeFolderWorkflowListResponse, err error) {
    return c.DescribeFolderWorkflowListWithContext(context.Background(), request)
}

// DescribeFolderWorkflowList
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 拉取文件夹下的工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFolderWorkflowListWithContext(ctx context.Context, request *DescribeFolderWorkflowListRequest) (response *DescribeFolderWorkflowListResponse, err error) {
    if request == nil {
        request = NewDescribeFolderWorkflowListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFolderWorkflowList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFolderWorkflowListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFunctionKindsRequest() (request *DescribeFunctionKindsRequest) {
    request = &DescribeFunctionKindsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeFunctionKinds")
    
    
    return
}

func NewDescribeFunctionKindsResponse() (response *DescribeFunctionKindsResponse) {
    response = &DescribeFunctionKindsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFunctionKinds
// 查询函数分类
//
// 可能返回的错误码:
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeFunctionKinds(request *DescribeFunctionKindsRequest) (response *DescribeFunctionKindsResponse, err error) {
    return c.DescribeFunctionKindsWithContext(context.Background(), request)
}

// DescribeFunctionKinds
// 查询函数分类
//
// 可能返回的错误码:
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeFunctionKindsWithContext(ctx context.Context, request *DescribeFunctionKindsRequest) (response *DescribeFunctionKindsResponse, err error) {
    if request == nil {
        request = NewDescribeFunctionKindsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFunctionKinds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFunctionKindsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFunctionTypesRequest() (request *DescribeFunctionTypesRequest) {
    request = &DescribeFunctionTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeFunctionTypes")
    
    
    return
}

func NewDescribeFunctionTypesResponse() (response *DescribeFunctionTypesResponse) {
    response = &DescribeFunctionTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFunctionTypes
// 查询函数类型
//
// 可能返回的错误码:
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeFunctionTypes(request *DescribeFunctionTypesRequest) (response *DescribeFunctionTypesResponse, err error) {
    return c.DescribeFunctionTypesWithContext(context.Background(), request)
}

// DescribeFunctionTypes
// 查询函数类型
//
// 可能返回的错误码:
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) DescribeFunctionTypesWithContext(ctx context.Context, request *DescribeFunctionTypesRequest) (response *DescribeFunctionTypesResponse, err error) {
    if request == nil {
        request = NewDescribeFunctionTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFunctionTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFunctionTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInLongAgentListRequest() (request *DescribeInLongAgentListRequest) {
    request = &DescribeInLongAgentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInLongAgentList")
    
    
    return
}

func NewDescribeInLongAgentListResponse() (response *DescribeInLongAgentListResponse) {
    response = &DescribeInLongAgentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInLongAgentList
// 获取采集器列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInLongAgentList(request *DescribeInLongAgentListRequest) (response *DescribeInLongAgentListResponse, err error) {
    return c.DescribeInLongAgentListWithContext(context.Background(), request)
}

// DescribeInLongAgentList
// 获取采集器列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInLongAgentListWithContext(ctx context.Context, request *DescribeInLongAgentListRequest) (response *DescribeInLongAgentListResponse, err error) {
    if request == nil {
        request = NewDescribeInLongAgentListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInLongAgentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInLongAgentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInLongAgentTaskListRequest() (request *DescribeInLongAgentTaskListRequest) {
    request = &DescribeInLongAgentTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInLongAgentTaskList")
    
    
    return
}

func NewDescribeInLongAgentTaskListResponse() (response *DescribeInLongAgentTaskListResponse) {
    response = &DescribeInLongAgentTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInLongAgentTaskList
// 查询采集器关联的任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInLongAgentTaskList(request *DescribeInLongAgentTaskListRequest) (response *DescribeInLongAgentTaskListResponse, err error) {
    return c.DescribeInLongAgentTaskListWithContext(context.Background(), request)
}

// DescribeInLongAgentTaskList
// 查询采集器关联的任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInLongAgentTaskListWithContext(ctx context.Context, request *DescribeInLongAgentTaskListRequest) (response *DescribeInLongAgentTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeInLongAgentTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInLongAgentTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInLongAgentTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInLongAgentVpcListRequest() (request *DescribeInLongAgentVpcListRequest) {
    request = &DescribeInLongAgentVpcListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInLongAgentVpcList")
    
    
    return
}

func NewDescribeInLongAgentVpcListResponse() (response *DescribeInLongAgentVpcListResponse) {
    response = &DescribeInLongAgentVpcListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInLongAgentVpcList
// 获取采集器所在集群的VPC列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInLongAgentVpcList(request *DescribeInLongAgentVpcListRequest) (response *DescribeInLongAgentVpcListResponse, err error) {
    return c.DescribeInLongAgentVpcListWithContext(context.Background(), request)
}

// DescribeInLongAgentVpcList
// 获取采集器所在集群的VPC列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInLongAgentVpcListWithContext(ctx context.Context, request *DescribeInLongAgentVpcListRequest) (response *DescribeInLongAgentVpcListResponse, err error) {
    if request == nil {
        request = NewDescribeInLongAgentVpcListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInLongAgentVpcList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInLongAgentVpcListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInLongTkeClusterListRequest() (request *DescribeInLongTkeClusterListRequest) {
    request = &DescribeInLongTkeClusterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInLongTkeClusterList")
    
    
    return
}

func NewDescribeInLongTkeClusterListResponse() (response *DescribeInLongTkeClusterListResponse) {
    response = &DescribeInLongTkeClusterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInLongTkeClusterList
// 获取TKE集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInLongTkeClusterList(request *DescribeInLongTkeClusterListRequest) (response *DescribeInLongTkeClusterListResponse, err error) {
    return c.DescribeInLongTkeClusterListWithContext(context.Background(), request)
}

// DescribeInLongTkeClusterList
// 获取TKE集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInLongTkeClusterListWithContext(ctx context.Context, request *DescribeInLongTkeClusterListRequest) (response *DescribeInLongTkeClusterListResponse, err error) {
    if request == nil {
        request = NewDescribeInLongTkeClusterListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInLongTkeClusterList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInLongTkeClusterListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceByCycleRequest() (request *DescribeInstanceByCycleRequest) {
    request = &DescribeInstanceByCycleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstanceByCycle")
    
    
    return
}

func NewDescribeInstanceByCycleResponse() (response *DescribeInstanceByCycleResponse) {
    response = &DescribeInstanceByCycleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceByCycle
// 根据周期类型查询所有实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceByCycle(request *DescribeInstanceByCycleRequest) (response *DescribeInstanceByCycleResponse, err error) {
    return c.DescribeInstanceByCycleWithContext(context.Background(), request)
}

// DescribeInstanceByCycle
// 根据周期类型查询所有实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceByCycleWithContext(ctx context.Context, request *DescribeInstanceByCycleRequest) (response *DescribeInstanceByCycleResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceByCycleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceByCycle require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceByCycleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceByCycleReportRequest() (request *DescribeInstanceByCycleReportRequest) {
    request = &DescribeInstanceByCycleReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstanceByCycleReport")
    
    
    return
}

func NewDescribeInstanceByCycleReportResponse() (response *DescribeInstanceByCycleReportResponse) {
    response = &DescribeInstanceByCycleReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceByCycleReport
// 实例状态周期增长趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceByCycleReport(request *DescribeInstanceByCycleReportRequest) (response *DescribeInstanceByCycleReportResponse, err error) {
    return c.DescribeInstanceByCycleReportWithContext(context.Background(), request)
}

// DescribeInstanceByCycleReport
// 实例状态周期增长趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceByCycleReportWithContext(ctx context.Context, request *DescribeInstanceByCycleReportRequest) (response *DescribeInstanceByCycleReportResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceByCycleReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceByCycleReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceByCycleReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLastLogRequest() (request *DescribeInstanceLastLogRequest) {
    request = &DescribeInstanceLastLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstanceLastLog")
    
    
    return
}

func NewDescribeInstanceLastLogResponse() (response *DescribeInstanceLastLogResponse) {
    response = &DescribeInstanceLastLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceLastLog
// 日志获取详情页面
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLastLog(request *DescribeInstanceLastLogRequest) (response *DescribeInstanceLastLogResponse, err error) {
    return c.DescribeInstanceLastLogWithContext(context.Background(), request)
}

// DescribeInstanceLastLog
// 日志获取详情页面
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLastLogWithContext(ctx context.Context, request *DescribeInstanceLastLogRequest) (response *DescribeInstanceLastLogResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLastLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLastLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLastLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceListRequest() (request *DescribeInstanceListRequest) {
    request = &DescribeInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstanceList")
    
    
    return
}

func NewDescribeInstanceListResponse() (response *DescribeInstanceListResponse) {
    response = &DescribeInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceList
// 获取实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    return c.DescribeInstanceListWithContext(context.Background(), request)
}

// DescribeInstanceList
// 获取实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceListWithContext(ctx context.Context, request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogRequest() (request *DescribeInstanceLogRequest) {
    request = &DescribeInstanceLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstanceLog")
    
    
    return
}

func NewDescribeInstanceLogResponse() (response *DescribeInstanceLogResponse) {
    response = &DescribeInstanceLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceLog
// 获取实例运行日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLog(request *DescribeInstanceLogRequest) (response *DescribeInstanceLogResponse, err error) {
    return c.DescribeInstanceLogWithContext(context.Background(), request)
}

// DescribeInstanceLog
// 获取实例运行日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogWithContext(ctx context.Context, request *DescribeInstanceLogRequest) (response *DescribeInstanceLogResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogDetailRequest() (request *DescribeInstanceLogDetailRequest) {
    request = &DescribeInstanceLogDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstanceLogDetail")
    
    
    return
}

func NewDescribeInstanceLogDetailResponse() (response *DescribeInstanceLogDetailResponse) {
    response = &DescribeInstanceLogDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceLogDetail
// 获取具体实例相关日志信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogDetail(request *DescribeInstanceLogDetailRequest) (response *DescribeInstanceLogDetailResponse, err error) {
    return c.DescribeInstanceLogDetailWithContext(context.Background(), request)
}

// DescribeInstanceLogDetail
// 获取具体实例相关日志信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogDetailWithContext(ctx context.Context, request *DescribeInstanceLogDetailRequest) (response *DescribeInstanceLogDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLogDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLogDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogFileRequest() (request *DescribeInstanceLogFileRequest) {
    request = &DescribeInstanceLogFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstanceLogFile")
    
    
    return
}

func NewDescribeInstanceLogFileResponse() (response *DescribeInstanceLogFileResponse) {
    response = &DescribeInstanceLogFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceLogFile
// 下载日志文件，返回日志URL
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogFile(request *DescribeInstanceLogFileRequest) (response *DescribeInstanceLogFileResponse, err error) {
    return c.DescribeInstanceLogFileWithContext(context.Background(), request)
}

// DescribeInstanceLogFile
// 下载日志文件，返回日志URL
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogFileWithContext(ctx context.Context, request *DescribeInstanceLogFileRequest) (response *DescribeInstanceLogFileResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLogFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLogFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogListRequest() (request *DescribeInstanceLogListRequest) {
    request = &DescribeInstanceLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstanceLogList")
    
    
    return
}

func NewDescribeInstanceLogListResponse() (response *DescribeInstanceLogListResponse) {
    response = &DescribeInstanceLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceLogList
// 离线任务实例运行日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogList(request *DescribeInstanceLogListRequest) (response *DescribeInstanceLogListResponse, err error) {
    return c.DescribeInstanceLogListWithContext(context.Background(), request)
}

// DescribeInstanceLogList
// 离线任务实例运行日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogListWithContext(ctx context.Context, request *DescribeInstanceLogListRequest) (response *DescribeInstanceLogListResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogsRequest() (request *DescribeInstanceLogsRequest) {
    request = &DescribeInstanceLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstanceLogs")
    
    
    return
}

func NewDescribeInstanceLogsResponse() (response *DescribeInstanceLogsResponse) {
    response = &DescribeInstanceLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceLogs
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 获取实例日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogs(request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    return c.DescribeInstanceLogsWithContext(context.Background(), request)
}

// DescribeInstanceLogs
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 获取实例日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceLogsWithContext(ctx context.Context, request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 数据质量，查询调度任务的实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 数据质量，查询调度任务的实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesInfoWithTaskInfoRequest() (request *DescribeInstancesInfoWithTaskInfoRequest) {
    request = &DescribeInstancesInfoWithTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeInstancesInfoWithTaskInfo")
    
    
    return
}

func NewDescribeInstancesInfoWithTaskInfoResponse() (response *DescribeInstancesInfoWithTaskInfoResponse) {
    response = &DescribeInstancesInfoWithTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesInfoWithTaskInfo
// 拉取实例列表，join task表一些信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstancesInfoWithTaskInfo(request *DescribeInstancesInfoWithTaskInfoRequest) (response *DescribeInstancesInfoWithTaskInfoResponse, err error) {
    return c.DescribeInstancesInfoWithTaskInfoWithContext(context.Background(), request)
}

// DescribeInstancesInfoWithTaskInfo
// 拉取实例列表，join task表一些信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstancesInfoWithTaskInfoWithContext(ctx context.Context, request *DescribeInstancesInfoWithTaskInfoRequest) (response *DescribeInstancesInfoWithTaskInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesInfoWithTaskInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesInfoWithTaskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesInfoWithTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationNodeRequest() (request *DescribeIntegrationNodeRequest) {
    request = &DescribeIntegrationNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeIntegrationNode")
    
    
    return
}

func NewDescribeIntegrationNodeResponse() (response *DescribeIntegrationNodeResponse) {
    response = &DescribeIntegrationNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationNode
// 查询集成节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIntegrationNode(request *DescribeIntegrationNodeRequest) (response *DescribeIntegrationNodeResponse, err error) {
    return c.DescribeIntegrationNodeWithContext(context.Background(), request)
}

// DescribeIntegrationNode
// 查询集成节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIntegrationNodeWithContext(ctx context.Context, request *DescribeIntegrationNodeRequest) (response *DescribeIntegrationNodeResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationStatisticsRequest() (request *DescribeIntegrationStatisticsRequest) {
    request = &DescribeIntegrationStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeIntegrationStatistics")
    
    
    return
}

func NewDescribeIntegrationStatisticsResponse() (response *DescribeIntegrationStatisticsResponse) {
    response = &DescribeIntegrationStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationStatistics
// 数据集成大屏概览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatistics(request *DescribeIntegrationStatisticsRequest) (response *DescribeIntegrationStatisticsResponse, err error) {
    return c.DescribeIntegrationStatisticsWithContext(context.Background(), request)
}

// DescribeIntegrationStatistics
// 数据集成大屏概览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsWithContext(ctx context.Context, request *DescribeIntegrationStatisticsRequest) (response *DescribeIntegrationStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationStatisticsAgentStatusRequest() (request *DescribeIntegrationStatisticsAgentStatusRequest) {
    request = &DescribeIntegrationStatisticsAgentStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeIntegrationStatisticsAgentStatus")
    
    
    return
}

func NewDescribeIntegrationStatisticsAgentStatusResponse() (response *DescribeIntegrationStatisticsAgentStatusResponse) {
    response = &DescribeIntegrationStatisticsAgentStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationStatisticsAgentStatus
// 数据集成大屏采集器状态分布统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsAgentStatus(request *DescribeIntegrationStatisticsAgentStatusRequest) (response *DescribeIntegrationStatisticsAgentStatusResponse, err error) {
    return c.DescribeIntegrationStatisticsAgentStatusWithContext(context.Background(), request)
}

// DescribeIntegrationStatisticsAgentStatus
// 数据集成大屏采集器状态分布统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsAgentStatusWithContext(ctx context.Context, request *DescribeIntegrationStatisticsAgentStatusRequest) (response *DescribeIntegrationStatisticsAgentStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationStatisticsAgentStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationStatisticsAgentStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationStatisticsAgentStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationStatisticsInstanceTrendRequest() (request *DescribeIntegrationStatisticsInstanceTrendRequest) {
    request = &DescribeIntegrationStatisticsInstanceTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeIntegrationStatisticsInstanceTrend")
    
    
    return
}

func NewDescribeIntegrationStatisticsInstanceTrendResponse() (response *DescribeIntegrationStatisticsInstanceTrendResponse) {
    response = &DescribeIntegrationStatisticsInstanceTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationStatisticsInstanceTrend
// 数据集成大屏实例状态统计趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsInstanceTrend(request *DescribeIntegrationStatisticsInstanceTrendRequest) (response *DescribeIntegrationStatisticsInstanceTrendResponse, err error) {
    return c.DescribeIntegrationStatisticsInstanceTrendWithContext(context.Background(), request)
}

// DescribeIntegrationStatisticsInstanceTrend
// 数据集成大屏实例状态统计趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsInstanceTrendWithContext(ctx context.Context, request *DescribeIntegrationStatisticsInstanceTrendRequest) (response *DescribeIntegrationStatisticsInstanceTrendResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationStatisticsInstanceTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationStatisticsInstanceTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationStatisticsInstanceTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationStatisticsRecordsTrendRequest() (request *DescribeIntegrationStatisticsRecordsTrendRequest) {
    request = &DescribeIntegrationStatisticsRecordsTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeIntegrationStatisticsRecordsTrend")
    
    
    return
}

func NewDescribeIntegrationStatisticsRecordsTrendResponse() (response *DescribeIntegrationStatisticsRecordsTrendResponse) {
    response = &DescribeIntegrationStatisticsRecordsTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationStatisticsRecordsTrend
// 数据集成大屏同步条数统计趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsRecordsTrend(request *DescribeIntegrationStatisticsRecordsTrendRequest) (response *DescribeIntegrationStatisticsRecordsTrendResponse, err error) {
    return c.DescribeIntegrationStatisticsRecordsTrendWithContext(context.Background(), request)
}

// DescribeIntegrationStatisticsRecordsTrend
// 数据集成大屏同步条数统计趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsRecordsTrendWithContext(ctx context.Context, request *DescribeIntegrationStatisticsRecordsTrendRequest) (response *DescribeIntegrationStatisticsRecordsTrendResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationStatisticsRecordsTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationStatisticsRecordsTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationStatisticsRecordsTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationStatisticsTaskStatusRequest() (request *DescribeIntegrationStatisticsTaskStatusRequest) {
    request = &DescribeIntegrationStatisticsTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeIntegrationStatisticsTaskStatus")
    
    
    return
}

func NewDescribeIntegrationStatisticsTaskStatusResponse() (response *DescribeIntegrationStatisticsTaskStatusResponse) {
    response = &DescribeIntegrationStatisticsTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationStatisticsTaskStatus
// 数据集成大屏任务状态分布统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsTaskStatus(request *DescribeIntegrationStatisticsTaskStatusRequest) (response *DescribeIntegrationStatisticsTaskStatusResponse, err error) {
    return c.DescribeIntegrationStatisticsTaskStatusWithContext(context.Background(), request)
}

// DescribeIntegrationStatisticsTaskStatus
// 数据集成大屏任务状态分布统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsTaskStatusWithContext(ctx context.Context, request *DescribeIntegrationStatisticsTaskStatusRequest) (response *DescribeIntegrationStatisticsTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationStatisticsTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationStatisticsTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationStatisticsTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationStatisticsTaskStatusTrendRequest() (request *DescribeIntegrationStatisticsTaskStatusTrendRequest) {
    request = &DescribeIntegrationStatisticsTaskStatusTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeIntegrationStatisticsTaskStatusTrend")
    
    
    return
}

func NewDescribeIntegrationStatisticsTaskStatusTrendResponse() (response *DescribeIntegrationStatisticsTaskStatusTrendResponse) {
    response = &DescribeIntegrationStatisticsTaskStatusTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationStatisticsTaskStatusTrend
// 数据集成大屏任务状态统计趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsTaskStatusTrend(request *DescribeIntegrationStatisticsTaskStatusTrendRequest) (response *DescribeIntegrationStatisticsTaskStatusTrendResponse, err error) {
    return c.DescribeIntegrationStatisticsTaskStatusTrendWithContext(context.Background(), request)
}

// DescribeIntegrationStatisticsTaskStatusTrend
// 数据集成大屏任务状态统计趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeIntegrationStatisticsTaskStatusTrendWithContext(ctx context.Context, request *DescribeIntegrationStatisticsTaskStatusTrendRequest) (response *DescribeIntegrationStatisticsTaskStatusTrendResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationStatisticsTaskStatusTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationStatisticsTaskStatusTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationStatisticsTaskStatusTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationTaskRequest() (request *DescribeIntegrationTaskRequest) {
    request = &DescribeIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeIntegrationTask")
    
    
    return
}

func NewDescribeIntegrationTaskResponse() (response *DescribeIntegrationTaskResponse) {
    response = &DescribeIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationTask
// 查询集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIntegrationTask(request *DescribeIntegrationTaskRequest) (response *DescribeIntegrationTaskResponse, err error) {
    return c.DescribeIntegrationTaskWithContext(context.Background(), request)
}

// DescribeIntegrationTask
// 查询集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIntegrationTaskWithContext(ctx context.Context, request *DescribeIntegrationTaskRequest) (response *DescribeIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationTasksRequest() (request *DescribeIntegrationTasksRequest) {
    request = &DescribeIntegrationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeIntegrationTasks")
    
    
    return
}

func NewDescribeIntegrationTasksResponse() (response *DescribeIntegrationTasksResponse) {
    response = &DescribeIntegrationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationTasks
// 查询集成任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIntegrationTasks(request *DescribeIntegrationTasksRequest) (response *DescribeIntegrationTasksResponse, err error) {
    return c.DescribeIntegrationTasksWithContext(context.Background(), request)
}

// DescribeIntegrationTasks
// 查询集成任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIntegrationTasksWithContext(ctx context.Context, request *DescribeIntegrationTasksRequest) (response *DescribeIntegrationTasksResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationVersionNodesInfoRequest() (request *DescribeIntegrationVersionNodesInfoRequest) {
    request = &DescribeIntegrationVersionNodesInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeIntegrationVersionNodesInfo")
    
    
    return
}

func NewDescribeIntegrationVersionNodesInfoResponse() (response *DescribeIntegrationVersionNodesInfoResponse) {
    response = &DescribeIntegrationVersionNodesInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationVersionNodesInfo
// 查询集成任务版本节点信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIntegrationVersionNodesInfo(request *DescribeIntegrationVersionNodesInfoRequest) (response *DescribeIntegrationVersionNodesInfoResponse, err error) {
    return c.DescribeIntegrationVersionNodesInfoWithContext(context.Background(), request)
}

// DescribeIntegrationVersionNodesInfo
// 查询集成任务版本节点信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIntegrationVersionNodesInfoWithContext(ctx context.Context, request *DescribeIntegrationVersionNodesInfoRequest) (response *DescribeIntegrationVersionNodesInfoResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationVersionNodesInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationVersionNodesInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationVersionNodesInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKafkaTopicInfoRequest() (request *DescribeKafkaTopicInfoRequest) {
    request = &DescribeKafkaTopicInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeKafkaTopicInfo")
    
    
    return
}

func NewDescribeKafkaTopicInfoResponse() (response *DescribeKafkaTopicInfoResponse) {
    response = &DescribeKafkaTopicInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKafkaTopicInfo
// 没用到
//
// 
//
// 获取kafka的topic信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKafkaTopicInfo(request *DescribeKafkaTopicInfoRequest) (response *DescribeKafkaTopicInfoResponse, err error) {
    return c.DescribeKafkaTopicInfoWithContext(context.Background(), request)
}

// DescribeKafkaTopicInfo
// 没用到
//
// 
//
// 获取kafka的topic信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKafkaTopicInfoWithContext(ctx context.Context, request *DescribeKafkaTopicInfoRequest) (response *DescribeKafkaTopicInfoResponse, err error) {
    if request == nil {
        request = NewDescribeKafkaTopicInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKafkaTopicInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKafkaTopicInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorsByPageRequest() (request *DescribeMonitorsByPageRequest) {
    request = &DescribeMonitorsByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeMonitorsByPage")
    
    
    return
}

func NewDescribeMonitorsByPageResponse() (response *DescribeMonitorsByPageResponse) {
    response = &DescribeMonitorsByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMonitorsByPage
// 分页查询质量监控组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMonitorsByPage(request *DescribeMonitorsByPageRequest) (response *DescribeMonitorsByPageResponse, err error) {
    return c.DescribeMonitorsByPageWithContext(context.Background(), request)
}

// DescribeMonitorsByPage
// 分页查询质量监控组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMonitorsByPageWithContext(ctx context.Context, request *DescribeMonitorsByPageRequest) (response *DescribeMonitorsByPageResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorsByPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMonitorsByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMonitorsByPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineTaskTokenRequest() (request *DescribeOfflineTaskTokenRequest) {
    request = &DescribeOfflineTaskTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOfflineTaskToken")
    
    
    return
}

func NewDescribeOfflineTaskTokenResponse() (response *DescribeOfflineTaskTokenResponse) {
    response = &DescribeOfflineTaskTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOfflineTaskToken
// 获取离线任务长连接Token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOfflineTaskToken(request *DescribeOfflineTaskTokenRequest) (response *DescribeOfflineTaskTokenResponse, err error) {
    return c.DescribeOfflineTaskTokenWithContext(context.Background(), request)
}

// DescribeOfflineTaskToken
// 获取离线任务长连接Token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOfflineTaskTokenWithContext(ctx context.Context, request *DescribeOfflineTaskTokenRequest) (response *DescribeOfflineTaskTokenResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineTaskTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfflineTaskToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfflineTaskTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOperateOpsTaskDatasourceRequest() (request *DescribeOperateOpsTaskDatasourceRequest) {
    request = &DescribeOperateOpsTaskDatasourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOperateOpsTaskDatasource")
    
    
    return
}

func NewDescribeOperateOpsTaskDatasourceResponse() (response *DescribeOperateOpsTaskDatasourceResponse) {
    response = &DescribeOperateOpsTaskDatasourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOperateOpsTaskDatasource
// 任务运维搜索 查询生产态任务数据源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOperateOpsTaskDatasource(request *DescribeOperateOpsTaskDatasourceRequest) (response *DescribeOperateOpsTaskDatasourceResponse, err error) {
    return c.DescribeOperateOpsTaskDatasourceWithContext(context.Background(), request)
}

// DescribeOperateOpsTaskDatasource
// 任务运维搜索 查询生产态任务数据源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOperateOpsTaskDatasourceWithContext(ctx context.Context, request *DescribeOperateOpsTaskDatasourceRequest) (response *DescribeOperateOpsTaskDatasourceResponse, err error) {
    if request == nil {
        request = NewDescribeOperateOpsTaskDatasourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOperateOpsTaskDatasource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOperateOpsTaskDatasourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOperateOpsTaskDatasourceTypeRequest() (request *DescribeOperateOpsTaskDatasourceTypeRequest) {
    request = &DescribeOperateOpsTaskDatasourceTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOperateOpsTaskDatasourceType")
    
    
    return
}

func NewDescribeOperateOpsTaskDatasourceTypeResponse() (response *DescribeOperateOpsTaskDatasourceTypeResponse) {
    response = &DescribeOperateOpsTaskDatasourceTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOperateOpsTaskDatasourceType
// 任务运维-查询生产态任务数据源类型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOperateOpsTaskDatasourceType(request *DescribeOperateOpsTaskDatasourceTypeRequest) (response *DescribeOperateOpsTaskDatasourceTypeResponse, err error) {
    return c.DescribeOperateOpsTaskDatasourceTypeWithContext(context.Background(), request)
}

// DescribeOperateOpsTaskDatasourceType
// 任务运维-查询生产态任务数据源类型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOperateOpsTaskDatasourceTypeWithContext(ctx context.Context, request *DescribeOperateOpsTaskDatasourceTypeRequest) (response *DescribeOperateOpsTaskDatasourceTypeResponse, err error) {
    if request == nil {
        request = NewDescribeOperateOpsTaskDatasourceTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOperateOpsTaskDatasourceType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOperateOpsTaskDatasourceTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOperateOpsTasksRequest() (request *DescribeOperateOpsTasksRequest) {
    request = &DescribeOperateOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOperateOpsTasks")
    
    
    return
}

func NewDescribeOperateOpsTasksResponse() (response *DescribeOperateOpsTasksResponse) {
    response = &DescribeOperateOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOperateOpsTasks
// 任务运维列表组合条件查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOperateOpsTasks(request *DescribeOperateOpsTasksRequest) (response *DescribeOperateOpsTasksResponse, err error) {
    return c.DescribeOperateOpsTasksWithContext(context.Background(), request)
}

// DescribeOperateOpsTasks
// 任务运维列表组合条件查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOperateOpsTasksWithContext(ctx context.Context, request *DescribeOperateOpsTasksRequest) (response *DescribeOperateOpsTasksResponse, err error) {
    if request == nil {
        request = NewDescribeOperateOpsTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOperateOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOperateOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOperateTasksRequest() (request *DescribeOperateTasksRequest) {
    request = &DescribeOperateTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOperateTasks")
    
    
    return
}

func NewDescribeOperateTasksResponse() (response *DescribeOperateTasksResponse) {
    response = &DescribeOperateTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOperateTasks
// 任务运维列表组合条件查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOperateTasks(request *DescribeOperateTasksRequest) (response *DescribeOperateTasksResponse, err error) {
    return c.DescribeOperateTasksWithContext(context.Background(), request)
}

// DescribeOperateTasks
// 任务运维列表组合条件查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOperateTasksWithContext(ctx context.Context, request *DescribeOperateTasksRequest) (response *DescribeOperateTasksResponse, err error) {
    if request == nil {
        request = NewDescribeOperateTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOperateTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOperateTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpsInstanceLogListRequest() (request *DescribeOpsInstanceLogListRequest) {
    request = &DescribeOpsInstanceLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOpsInstanceLogList")
    
    
    return
}

func NewDescribeOpsInstanceLogListResponse() (response *DescribeOpsInstanceLogListResponse) {
    response = &DescribeOpsInstanceLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOpsInstanceLogList
// 实例运维-获取实例日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOpsInstanceLogList(request *DescribeOpsInstanceLogListRequest) (response *DescribeOpsInstanceLogListResponse, err error) {
    return c.DescribeOpsInstanceLogListWithContext(context.Background(), request)
}

// DescribeOpsInstanceLogList
// 实例运维-获取实例日志列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOpsInstanceLogListWithContext(ctx context.Context, request *DescribeOpsInstanceLogListRequest) (response *DescribeOpsInstanceLogListResponse, err error) {
    if request == nil {
        request = NewDescribeOpsInstanceLogListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpsInstanceLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpsInstanceLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpsMakePlanInstancesRequest() (request *DescribeOpsMakePlanInstancesRequest) {
    request = &DescribeOpsMakePlanInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOpsMakePlanInstances")
    
    
    return
}

func NewDescribeOpsMakePlanInstancesResponse() (response *DescribeOpsMakePlanInstancesResponse) {
    response = &DescribeOpsMakePlanInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOpsMakePlanInstances
// 根据补录计划和补录任务获取补录实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOpsMakePlanInstances(request *DescribeOpsMakePlanInstancesRequest) (response *DescribeOpsMakePlanInstancesResponse, err error) {
    return c.DescribeOpsMakePlanInstancesWithContext(context.Background(), request)
}

// DescribeOpsMakePlanInstances
// 根据补录计划和补录任务获取补录实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOpsMakePlanInstancesWithContext(ctx context.Context, request *DescribeOpsMakePlanInstancesRequest) (response *DescribeOpsMakePlanInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeOpsMakePlanInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpsMakePlanInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpsMakePlanInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpsMakePlanTasksRequest() (request *DescribeOpsMakePlanTasksRequest) {
    request = &DescribeOpsMakePlanTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOpsMakePlanTasks")
    
    
    return
}

func NewDescribeOpsMakePlanTasksResponse() (response *DescribeOpsMakePlanTasksResponse) {
    response = &DescribeOpsMakePlanTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOpsMakePlanTasks
// 查看补录计划任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOpsMakePlanTasks(request *DescribeOpsMakePlanTasksRequest) (response *DescribeOpsMakePlanTasksResponse, err error) {
    return c.DescribeOpsMakePlanTasksWithContext(context.Background(), request)
}

// DescribeOpsMakePlanTasks
// 查看补录计划任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOpsMakePlanTasksWithContext(ctx context.Context, request *DescribeOpsMakePlanTasksRequest) (response *DescribeOpsMakePlanTasksResponse, err error) {
    if request == nil {
        request = NewDescribeOpsMakePlanTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpsMakePlanTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpsMakePlanTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpsMakePlansRequest() (request *DescribeOpsMakePlansRequest) {
    request = &DescribeOpsMakePlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOpsMakePlans")
    
    
    return
}

func NewDescribeOpsMakePlansResponse() (response *DescribeOpsMakePlansResponse) {
    response = &DescribeOpsMakePlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOpsMakePlans
// 根据条件分页查询补录计划
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOpsMakePlans(request *DescribeOpsMakePlansRequest) (response *DescribeOpsMakePlansResponse, err error) {
    return c.DescribeOpsMakePlansWithContext(context.Background(), request)
}

// DescribeOpsMakePlans
// 根据条件分页查询补录计划
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOpsMakePlansWithContext(ctx context.Context, request *DescribeOpsMakePlansRequest) (response *DescribeOpsMakePlansResponse, err error) {
    if request == nil {
        request = NewDescribeOpsMakePlansRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpsMakePlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpsMakePlansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOpsWorkflowsRequest() (request *DescribeOpsWorkflowsRequest) {
    request = &DescribeOpsWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOpsWorkflows")
    
    
    return
}

func NewDescribeOpsWorkflowsResponse() (response *DescribeOpsWorkflowsResponse) {
    response = &DescribeOpsWorkflowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOpsWorkflows
// 查询用户生产工作流列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOpsWorkflows(request *DescribeOpsWorkflowsRequest) (response *DescribeOpsWorkflowsResponse, err error) {
    return c.DescribeOpsWorkflowsWithContext(context.Background(), request)
}

// DescribeOpsWorkflows
// 查询用户生产工作流列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOpsWorkflowsWithContext(ctx context.Context, request *DescribeOpsWorkflowsRequest) (response *DescribeOpsWorkflowsResponse, err error) {
    if request == nil {
        request = NewDescribeOpsWorkflowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOpsWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOpsWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationalFunctionsRequest() (request *DescribeOrganizationalFunctionsRequest) {
    request = &DescribeOrganizationalFunctionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeOrganizationalFunctions")
    
    
    return
}

func NewDescribeOrganizationalFunctionsResponse() (response *DescribeOrganizationalFunctionsResponse) {
    response = &DescribeOrganizationalFunctionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationalFunctions
// 查询全量函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOrganizationalFunctions(request *DescribeOrganizationalFunctionsRequest) (response *DescribeOrganizationalFunctionsResponse, err error) {
    return c.DescribeOrganizationalFunctionsWithContext(context.Background(), request)
}

// DescribeOrganizationalFunctions
// 查询全量函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeOrganizationalFunctionsWithContext(ctx context.Context, request *DescribeOrganizationalFunctionsRequest) (response *DescribeOrganizationalFunctionsResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationalFunctionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationalFunctions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationalFunctionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProdTasksRequest() (request *DescribeProdTasksRequest) {
    request = &DescribeProdTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeProdTasks")
    
    
    return
}

func NewDescribeProdTasksResponse() (response *DescribeProdTasksResponse) {
    response = &DescribeProdTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProdTasks
// 数据质量获取生产调度任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeProdTasks(request *DescribeProdTasksRequest) (response *DescribeProdTasksResponse, err error) {
    return c.DescribeProdTasksWithContext(context.Background(), request)
}

// DescribeProdTasks
// 数据质量获取生产调度任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeProdTasksWithContext(ctx context.Context, request *DescribeProdTasksRequest) (response *DescribeProdTasksResponse, err error) {
    if request == nil {
        request = NewDescribeProdTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProdTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProdTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectRequest() (request *DescribeProjectRequest) {
    request = &DescribeProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeProject")
    
    
    return
}

func NewDescribeProjectResponse() (response *DescribeProjectResponse) {
    response = &DescribeProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProject
// 获取项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
func (c *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    return c.DescribeProjectWithContext(context.Background(), request)
}

// DescribeProject
// 获取项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
func (c *Client) DescribeProjectWithContext(ctx context.Context, request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
    if request == nil {
        request = NewDescribeProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQualityScoreRequest() (request *DescribeQualityScoreRequest) {
    request = &DescribeQualityScoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeQualityScore")
    
    
    return
}

func NewDescribeQualityScoreResponse() (response *DescribeQualityScoreResponse) {
    response = &DescribeQualityScoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQualityScore
// 质量报告-质量评分
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQualityScore(request *DescribeQualityScoreRequest) (response *DescribeQualityScoreResponse, err error) {
    return c.DescribeQualityScoreWithContext(context.Background(), request)
}

// DescribeQualityScore
// 质量报告-质量评分
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQualityScoreWithContext(ctx context.Context, request *DescribeQualityScoreRequest) (response *DescribeQualityScoreResponse, err error) {
    if request == nil {
        request = NewDescribeQualityScoreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQualityScore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQualityScoreResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQualityScoreTrendRequest() (request *DescribeQualityScoreTrendRequest) {
    request = &DescribeQualityScoreTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeQualityScoreTrend")
    
    
    return
}

func NewDescribeQualityScoreTrendResponse() (response *DescribeQualityScoreTrendResponse) {
    response = &DescribeQualityScoreTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQualityScoreTrend
// 质量报告-质量分周期趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQualityScoreTrend(request *DescribeQualityScoreTrendRequest) (response *DescribeQualityScoreTrendResponse, err error) {
    return c.DescribeQualityScoreTrendWithContext(context.Background(), request)
}

// DescribeQualityScoreTrend
// 质量报告-质量分周期趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQualityScoreTrendWithContext(ctx context.Context, request *DescribeQualityScoreTrendRequest) (response *DescribeQualityScoreTrendResponse, err error) {
    if request == nil {
        request = NewDescribeQualityScoreTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQualityScoreTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQualityScoreTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealTimeTaskInstanceNodeInfoRequest() (request *DescribeRealTimeTaskInstanceNodeInfoRequest) {
    request = &DescribeRealTimeTaskInstanceNodeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRealTimeTaskInstanceNodeInfo")
    
    
    return
}

func NewDescribeRealTimeTaskInstanceNodeInfoResponse() (response *DescribeRealTimeTaskInstanceNodeInfoResponse) {
    response = &DescribeRealTimeTaskInstanceNodeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRealTimeTaskInstanceNodeInfo
// 查询实时任务实例节点信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRealTimeTaskInstanceNodeInfo(request *DescribeRealTimeTaskInstanceNodeInfoRequest) (response *DescribeRealTimeTaskInstanceNodeInfoResponse, err error) {
    return c.DescribeRealTimeTaskInstanceNodeInfoWithContext(context.Background(), request)
}

// DescribeRealTimeTaskInstanceNodeInfo
// 查询实时任务实例节点信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRealTimeTaskInstanceNodeInfoWithContext(ctx context.Context, request *DescribeRealTimeTaskInstanceNodeInfoRequest) (response *DescribeRealTimeTaskInstanceNodeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRealTimeTaskInstanceNodeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRealTimeTaskInstanceNodeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRealTimeTaskInstanceNodeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealTimeTaskMetricOverviewRequest() (request *DescribeRealTimeTaskMetricOverviewRequest) {
    request = &DescribeRealTimeTaskMetricOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRealTimeTaskMetricOverview")
    
    
    return
}

func NewDescribeRealTimeTaskMetricOverviewResponse() (response *DescribeRealTimeTaskMetricOverviewResponse) {
    response = &DescribeRealTimeTaskMetricOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRealTimeTaskMetricOverview
// 实时任务运行指标概览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRealTimeTaskMetricOverview(request *DescribeRealTimeTaskMetricOverviewRequest) (response *DescribeRealTimeTaskMetricOverviewResponse, err error) {
    return c.DescribeRealTimeTaskMetricOverviewWithContext(context.Background(), request)
}

// DescribeRealTimeTaskMetricOverview
// 实时任务运行指标概览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRealTimeTaskMetricOverviewWithContext(ctx context.Context, request *DescribeRealTimeTaskMetricOverviewRequest) (response *DescribeRealTimeTaskMetricOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeRealTimeTaskMetricOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRealTimeTaskMetricOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRealTimeTaskMetricOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealTimeTaskSpeedRequest() (request *DescribeRealTimeTaskSpeedRequest) {
    request = &DescribeRealTimeTaskSpeedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRealTimeTaskSpeed")
    
    
    return
}

func NewDescribeRealTimeTaskSpeedResponse() (response *DescribeRealTimeTaskSpeedResponse) {
    response = &DescribeRealTimeTaskSpeedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRealTimeTaskSpeed
// 实时任务同步速度趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRealTimeTaskSpeed(request *DescribeRealTimeTaskSpeedRequest) (response *DescribeRealTimeTaskSpeedResponse, err error) {
    return c.DescribeRealTimeTaskSpeedWithContext(context.Background(), request)
}

// DescribeRealTimeTaskSpeed
// 实时任务同步速度趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRealTimeTaskSpeedWithContext(ctx context.Context, request *DescribeRealTimeTaskSpeedRequest) (response *DescribeRealTimeTaskSpeedResponse, err error) {
    if request == nil {
        request = NewDescribeRealTimeTaskSpeedRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRealTimeTaskSpeed require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRealTimeTaskSpeedResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRelatedInstancesRequest() (request *DescribeRelatedInstancesRequest) {
    request = &DescribeRelatedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRelatedInstances")
    
    
    return
}

func NewDescribeRelatedInstancesResponse() (response *DescribeRelatedInstancesResponse) {
    response = &DescribeRelatedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRelatedInstances
// 查询任务实例的关联实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRelatedInstances(request *DescribeRelatedInstancesRequest) (response *DescribeRelatedInstancesResponse, err error) {
    return c.DescribeRelatedInstancesWithContext(context.Background(), request)
}

// DescribeRelatedInstances
// 查询任务实例的关联实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRelatedInstancesWithContext(ctx context.Context, request *DescribeRelatedInstancesRequest) (response *DescribeRelatedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeRelatedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRelatedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRelatedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceManagePathTreesRequest() (request *DescribeResourceManagePathTreesRequest) {
    request = &DescribeResourceManagePathTreesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeResourceManagePathTrees")
    
    
    return
}

func NewDescribeResourceManagePathTreesResponse() (response *DescribeResourceManagePathTreesResponse) {
    response = &DescribeResourceManagePathTreesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceManagePathTrees
// 获取资源管理目录树
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeResourceManagePathTrees(request *DescribeResourceManagePathTreesRequest) (response *DescribeResourceManagePathTreesResponse, err error) {
    return c.DescribeResourceManagePathTreesWithContext(context.Background(), request)
}

// DescribeResourceManagePathTrees
// 获取资源管理目录树
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeResourceManagePathTreesWithContext(ctx context.Context, request *DescribeResourceManagePathTreesRequest) (response *DescribeResourceManagePathTreesResponse, err error) {
    if request == nil {
        request = NewDescribeResourceManagePathTreesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceManagePathTrees require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceManagePathTreesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleRequest() (request *DescribeRuleRequest) {
    request = &DescribeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRule")
    
    
    return
}

func NewDescribeRuleResponse() (response *DescribeRuleResponse) {
    response = &DescribeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRule
// 查询规则详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
func (c *Client) DescribeRule(request *DescribeRuleRequest) (response *DescribeRuleResponse, err error) {
    return c.DescribeRuleWithContext(context.Background(), request)
}

// DescribeRule
// 查询规则详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
func (c *Client) DescribeRuleWithContext(ctx context.Context, request *DescribeRuleRequest) (response *DescribeRuleResponse, err error) {
    if request == nil {
        request = NewDescribeRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleDataSourcesRequest() (request *DescribeRuleDataSourcesRequest) {
    request = &DescribeRuleDataSourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleDataSources")
    
    
    return
}

func NewDescribeRuleDataSourcesResponse() (response *DescribeRuleDataSourcesResponse) {
    response = &DescribeRuleDataSourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleDataSources
// 查询质量规则数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleDataSources(request *DescribeRuleDataSourcesRequest) (response *DescribeRuleDataSourcesResponse, err error) {
    return c.DescribeRuleDataSourcesWithContext(context.Background(), request)
}

// DescribeRuleDataSources
// 查询质量规则数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleDataSourcesWithContext(ctx context.Context, request *DescribeRuleDataSourcesRequest) (response *DescribeRuleDataSourcesResponse, err error) {
    if request == nil {
        request = NewDescribeRuleDataSourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleDataSources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleDataSourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleDimStatRequest() (request *DescribeRuleDimStatRequest) {
    request = &DescribeRuleDimStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleDimStat")
    
    
    return
}

func NewDescribeRuleDimStatResponse() (response *DescribeRuleDimStatResponse) {
    response = &DescribeRuleDimStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleDimStat
// 数据质量概览页面触发维度分布统计接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleDimStat(request *DescribeRuleDimStatRequest) (response *DescribeRuleDimStatResponse, err error) {
    return c.DescribeRuleDimStatWithContext(context.Background(), request)
}

// DescribeRuleDimStat
// 数据质量概览页面触发维度分布统计接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleDimStatWithContext(ctx context.Context, request *DescribeRuleDimStatRequest) (response *DescribeRuleDimStatResponse, err error) {
    if request == nil {
        request = NewDescribeRuleDimStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleDimStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleDimStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleExecDetailRequest() (request *DescribeRuleExecDetailRequest) {
    request = &DescribeRuleExecDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleExecDetail")
    
    
    return
}

func NewDescribeRuleExecDetailResponse() (response *DescribeRuleExecDetailResponse) {
    response = &DescribeRuleExecDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleExecDetail
// 查询规则执行结果详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRuleExecDetail(request *DescribeRuleExecDetailRequest) (response *DescribeRuleExecDetailResponse, err error) {
    return c.DescribeRuleExecDetailWithContext(context.Background(), request)
}

// DescribeRuleExecDetail
// 查询规则执行结果详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRuleExecDetailWithContext(ctx context.Context, request *DescribeRuleExecDetailRequest) (response *DescribeRuleExecDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRuleExecDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleExecDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleExecDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleExecExportResultRequest() (request *DescribeRuleExecExportResultRequest) {
    request = &DescribeRuleExecExportResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleExecExportResult")
    
    
    return
}

func NewDescribeRuleExecExportResultResponse() (response *DescribeRuleExecExportResultResponse) {
    response = &DescribeRuleExecExportResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleExecExportResult
// 查询规则执行导出结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRuleExecExportResult(request *DescribeRuleExecExportResultRequest) (response *DescribeRuleExecExportResultResponse, err error) {
    return c.DescribeRuleExecExportResultWithContext(context.Background(), request)
}

// DescribeRuleExecExportResult
// 查询规则执行导出结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRuleExecExportResultWithContext(ctx context.Context, request *DescribeRuleExecExportResultRequest) (response *DescribeRuleExecExportResultResponse, err error) {
    if request == nil {
        request = NewDescribeRuleExecExportResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleExecExportResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleExecExportResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleExecHistoryRequest() (request *DescribeRuleExecHistoryRequest) {
    request = &DescribeRuleExecHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleExecHistory")
    
    
    return
}

func NewDescribeRuleExecHistoryResponse() (response *DescribeRuleExecHistoryResponse) {
    response = &DescribeRuleExecHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleExecHistory
// 查询规则执行历史， 最近30条
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRuleExecHistory(request *DescribeRuleExecHistoryRequest) (response *DescribeRuleExecHistoryResponse, err error) {
    return c.DescribeRuleExecHistoryWithContext(context.Background(), request)
}

// DescribeRuleExecHistory
// 查询规则执行历史， 最近30条
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRuleExecHistoryWithContext(ctx context.Context, request *DescribeRuleExecHistoryRequest) (response *DescribeRuleExecHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeRuleExecHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleExecHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleExecHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleExecLogRequest() (request *DescribeRuleExecLogRequest) {
    request = &DescribeRuleExecLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleExecLog")
    
    
    return
}

func NewDescribeRuleExecLogResponse() (response *DescribeRuleExecLogResponse) {
    response = &DescribeRuleExecLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleExecLog
// 规则执行日志查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
func (c *Client) DescribeRuleExecLog(request *DescribeRuleExecLogRequest) (response *DescribeRuleExecLogResponse, err error) {
    return c.DescribeRuleExecLogWithContext(context.Background(), request)
}

// DescribeRuleExecLog
// 规则执行日志查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
func (c *Client) DescribeRuleExecLogWithContext(ctx context.Context, request *DescribeRuleExecLogRequest) (response *DescribeRuleExecLogResponse, err error) {
    if request == nil {
        request = NewDescribeRuleExecLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleExecLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleExecLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleExecResultsRequest() (request *DescribeRuleExecResultsRequest) {
    request = &DescribeRuleExecResultsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleExecResults")
    
    
    return
}

func NewDescribeRuleExecResultsResponse() (response *DescribeRuleExecResultsResponse) {
    response = &DescribeRuleExecResultsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleExecResults
// 规则执行结果列表查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleExecResults(request *DescribeRuleExecResultsRequest) (response *DescribeRuleExecResultsResponse, err error) {
    return c.DescribeRuleExecResultsWithContext(context.Background(), request)
}

// DescribeRuleExecResults
// 规则执行结果列表查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleExecResultsWithContext(ctx context.Context, request *DescribeRuleExecResultsRequest) (response *DescribeRuleExecResultsResponse, err error) {
    if request == nil {
        request = NewDescribeRuleExecResultsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleExecResults require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleExecResultsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleExecResultsByPageRequest() (request *DescribeRuleExecResultsByPageRequest) {
    request = &DescribeRuleExecResultsByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleExecResultsByPage")
    
    
    return
}

func NewDescribeRuleExecResultsByPageResponse() (response *DescribeRuleExecResultsByPageResponse) {
    response = &DescribeRuleExecResultsByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleExecResultsByPage
// 分页查询规则执行结果列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleExecResultsByPage(request *DescribeRuleExecResultsByPageRequest) (response *DescribeRuleExecResultsByPageResponse, err error) {
    return c.DescribeRuleExecResultsByPageWithContext(context.Background(), request)
}

// DescribeRuleExecResultsByPage
// 分页查询规则执行结果列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleExecResultsByPageWithContext(ctx context.Context, request *DescribeRuleExecResultsByPageRequest) (response *DescribeRuleExecResultsByPageResponse, err error) {
    if request == nil {
        request = NewDescribeRuleExecResultsByPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleExecResultsByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleExecResultsByPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleExecStatRequest() (request *DescribeRuleExecStatRequest) {
    request = &DescribeRuleExecStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleExecStat")
    
    
    return
}

func NewDescribeRuleExecStatResponse() (response *DescribeRuleExecStatResponse) {
    response = &DescribeRuleExecStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleExecStat
// 数据质量概览页面规则运行情况接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleExecStat(request *DescribeRuleExecStatRequest) (response *DescribeRuleExecStatResponse, err error) {
    return c.DescribeRuleExecStatWithContext(context.Background(), request)
}

// DescribeRuleExecStat
// 数据质量概览页面规则运行情况接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleExecStatWithContext(ctx context.Context, request *DescribeRuleExecStatRequest) (response *DescribeRuleExecStatResponse, err error) {
    if request == nil {
        request = NewDescribeRuleExecStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleExecStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleExecStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleGroupRequest() (request *DescribeRuleGroupRequest) {
    request = &DescribeRuleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleGroup")
    
    
    return
}

func NewDescribeRuleGroupResponse() (response *DescribeRuleGroupResponse) {
    response = &DescribeRuleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleGroup
// 查询规则组详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleGroup(request *DescribeRuleGroupRequest) (response *DescribeRuleGroupResponse, err error) {
    return c.DescribeRuleGroupWithContext(context.Background(), request)
}

// DescribeRuleGroup
// 查询规则组详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleGroupWithContext(ctx context.Context, request *DescribeRuleGroupRequest) (response *DescribeRuleGroupResponse, err error) {
    if request == nil {
        request = NewDescribeRuleGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleGroupExecResultsByPageRequest() (request *DescribeRuleGroupExecResultsByPageRequest) {
    request = &DescribeRuleGroupExecResultsByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleGroupExecResultsByPage")
    
    
    return
}

func NewDescribeRuleGroupExecResultsByPageResponse() (response *DescribeRuleGroupExecResultsByPageResponse) {
    response = &DescribeRuleGroupExecResultsByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleGroupExecResultsByPage
// 规则组执行结果分页查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
func (c *Client) DescribeRuleGroupExecResultsByPage(request *DescribeRuleGroupExecResultsByPageRequest) (response *DescribeRuleGroupExecResultsByPageResponse, err error) {
    return c.DescribeRuleGroupExecResultsByPageWithContext(context.Background(), request)
}

// DescribeRuleGroupExecResultsByPage
// 规则组执行结果分页查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
func (c *Client) DescribeRuleGroupExecResultsByPageWithContext(ctx context.Context, request *DescribeRuleGroupExecResultsByPageRequest) (response *DescribeRuleGroupExecResultsByPageResponse, err error) {
    if request == nil {
        request = NewDescribeRuleGroupExecResultsByPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleGroupExecResultsByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleGroupExecResultsByPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleGroupExecResultsByPageWithoutAuthRequest() (request *DescribeRuleGroupExecResultsByPageWithoutAuthRequest) {
    request = &DescribeRuleGroupExecResultsByPageWithoutAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleGroupExecResultsByPageWithoutAuth")
    
    
    return
}

func NewDescribeRuleGroupExecResultsByPageWithoutAuthResponse() (response *DescribeRuleGroupExecResultsByPageWithoutAuthResponse) {
    response = &DescribeRuleGroupExecResultsByPageWithoutAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleGroupExecResultsByPageWithoutAuth
// 规则组执行结果分页查询接口不带鉴权
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
func (c *Client) DescribeRuleGroupExecResultsByPageWithoutAuth(request *DescribeRuleGroupExecResultsByPageWithoutAuthRequest) (response *DescribeRuleGroupExecResultsByPageWithoutAuthResponse, err error) {
    return c.DescribeRuleGroupExecResultsByPageWithoutAuthWithContext(context.Background(), request)
}

// DescribeRuleGroupExecResultsByPageWithoutAuth
// 规则组执行结果分页查询接口不带鉴权
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
func (c *Client) DescribeRuleGroupExecResultsByPageWithoutAuthWithContext(ctx context.Context, request *DescribeRuleGroupExecResultsByPageWithoutAuthRequest) (response *DescribeRuleGroupExecResultsByPageWithoutAuthResponse, err error) {
    if request == nil {
        request = NewDescribeRuleGroupExecResultsByPageWithoutAuthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleGroupExecResultsByPageWithoutAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleGroupExecResultsByPageWithoutAuthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleGroupSubscriptionRequest() (request *DescribeRuleGroupSubscriptionRequest) {
    request = &DescribeRuleGroupSubscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleGroupSubscription")
    
    
    return
}

func NewDescribeRuleGroupSubscriptionResponse() (response *DescribeRuleGroupSubscriptionResponse) {
    response = &DescribeRuleGroupSubscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleGroupSubscription
// 查询规则组订阅信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRuleGroupSubscription(request *DescribeRuleGroupSubscriptionRequest) (response *DescribeRuleGroupSubscriptionResponse, err error) {
    return c.DescribeRuleGroupSubscriptionWithContext(context.Background(), request)
}

// DescribeRuleGroupSubscription
// 查询规则组订阅信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRuleGroupSubscriptionWithContext(ctx context.Context, request *DescribeRuleGroupSubscriptionRequest) (response *DescribeRuleGroupSubscriptionResponse, err error) {
    if request == nil {
        request = NewDescribeRuleGroupSubscriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleGroupSubscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleGroupSubscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleGroupTableRequest() (request *DescribeRuleGroupTableRequest) {
    request = &DescribeRuleGroupTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleGroupTable")
    
    
    return
}

func NewDescribeRuleGroupTableResponse() (response *DescribeRuleGroupTableResponse) {
    response = &DescribeRuleGroupTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleGroupTable
// 查询表绑定执行规则组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRuleGroupTable(request *DescribeRuleGroupTableRequest) (response *DescribeRuleGroupTableResponse, err error) {
    return c.DescribeRuleGroupTableWithContext(context.Background(), request)
}

// DescribeRuleGroupTable
// 查询表绑定执行规则组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRuleGroupTableWithContext(ctx context.Context, request *DescribeRuleGroupTableRequest) (response *DescribeRuleGroupTableResponse, err error) {
    if request == nil {
        request = NewDescribeRuleGroupTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleGroupTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleGroupTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleGroupsByPageRequest() (request *DescribeRuleGroupsByPageRequest) {
    request = &DescribeRuleGroupsByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleGroupsByPage")
    
    
    return
}

func NewDescribeRuleGroupsByPageResponse() (response *DescribeRuleGroupsByPageResponse) {
    response = &DescribeRuleGroupsByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleGroupsByPage
// 【过滤条件】
//
// {表名称TableName,支持模糊匹配}       {表负责人TableOwnerName,支持模糊匹配}      {监控方式MonitorTypes，1.未配置 2.关联生产调度 3.离线周期检测,支持多选}  {订阅人ReceiverUin}
//
// 【必要字段】
//
// {数据来源DatasourceId}
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleGroupsByPage(request *DescribeRuleGroupsByPageRequest) (response *DescribeRuleGroupsByPageResponse, err error) {
    return c.DescribeRuleGroupsByPageWithContext(context.Background(), request)
}

// DescribeRuleGroupsByPage
// 【过滤条件】
//
// {表名称TableName,支持模糊匹配}       {表负责人TableOwnerName,支持模糊匹配}      {监控方式MonitorTypes，1.未配置 2.关联生产调度 3.离线周期检测,支持多选}  {订阅人ReceiverUin}
//
// 【必要字段】
//
// {数据来源DatasourceId}
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleGroupsByPageWithContext(ctx context.Context, request *DescribeRuleGroupsByPageRequest) (response *DescribeRuleGroupsByPageResponse, err error) {
    if request == nil {
        request = NewDescribeRuleGroupsByPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleGroupsByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleGroupsByPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleHistoryByPageRequest() (request *DescribeRuleHistoryByPageRequest) {
    request = &DescribeRuleHistoryByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleHistoryByPage")
    
    
    return
}

func NewDescribeRuleHistoryByPageResponse() (response *DescribeRuleHistoryByPageResponse) {
    response = &DescribeRuleHistoryByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleHistoryByPage
// 过滤条件【必要字段】{ruleId}
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
func (c *Client) DescribeRuleHistoryByPage(request *DescribeRuleHistoryByPageRequest) (response *DescribeRuleHistoryByPageResponse, err error) {
    return c.DescribeRuleHistoryByPageWithContext(context.Background(), request)
}

// DescribeRuleHistoryByPage
// 过滤条件【必要字段】{ruleId}
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
func (c *Client) DescribeRuleHistoryByPageWithContext(ctx context.Context, request *DescribeRuleHistoryByPageRequest) (response *DescribeRuleHistoryByPageResponse, err error) {
    if request == nil {
        request = NewDescribeRuleHistoryByPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleHistoryByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleHistoryByPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleTablesByPageRequest() (request *DescribeRuleTablesByPageRequest) {
    request = &DescribeRuleTablesByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleTablesByPage")
    
    
    return
}

func NewDescribeRuleTablesByPageResponse() (response *DescribeRuleTablesByPageResponse) {
    response = &DescribeRuleTablesByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleTablesByPage
// 获取表列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleTablesByPage(request *DescribeRuleTablesByPageRequest) (response *DescribeRuleTablesByPageResponse, err error) {
    return c.DescribeRuleTablesByPageWithContext(context.Background(), request)
}

// DescribeRuleTablesByPage
// 获取表列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleTablesByPageWithContext(ctx context.Context, request *DescribeRuleTablesByPageRequest) (response *DescribeRuleTablesByPageResponse, err error) {
    if request == nil {
        request = NewDescribeRuleTablesByPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleTablesByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleTablesByPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleTemplateRequest() (request *DescribeRuleTemplateRequest) {
    request = &DescribeRuleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleTemplate")
    
    
    return
}

func NewDescribeRuleTemplateResponse() (response *DescribeRuleTemplateResponse) {
    response = &DescribeRuleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleTemplate
// 查询模板详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RULETEMPLATENOTEXIST = "InvalidParameter.RuleTemplateNotExist"
func (c *Client) DescribeRuleTemplate(request *DescribeRuleTemplateRequest) (response *DescribeRuleTemplateResponse, err error) {
    return c.DescribeRuleTemplateWithContext(context.Background(), request)
}

// DescribeRuleTemplate
// 查询模板详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RULETEMPLATENOTEXIST = "InvalidParameter.RuleTemplateNotExist"
func (c *Client) DescribeRuleTemplateWithContext(ctx context.Context, request *DescribeRuleTemplateRequest) (response *DescribeRuleTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeRuleTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleTemplatesRequest() (request *DescribeRuleTemplatesRequest) {
    request = &DescribeRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleTemplates")
    
    
    return
}

func NewDescribeRuleTemplatesResponse() (response *DescribeRuleTemplatesResponse) {
    response = &DescribeRuleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleTemplates
// 查询规则模版列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRuleTemplates(request *DescribeRuleTemplatesRequest) (response *DescribeRuleTemplatesResponse, err error) {
    return c.DescribeRuleTemplatesWithContext(context.Background(), request)
}

// DescribeRuleTemplates
// 查询规则模版列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRuleTemplatesWithContext(ctx context.Context, request *DescribeRuleTemplatesRequest) (response *DescribeRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeRuleTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleTemplatesByPageRequest() (request *DescribeRuleTemplatesByPageRequest) {
    request = &DescribeRuleTemplatesByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRuleTemplatesByPage")
    
    
    return
}

func NewDescribeRuleTemplatesByPageResponse() (response *DescribeRuleTemplatesByPageResponse) {
    response = &DescribeRuleTemplatesByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleTemplatesByPage
// 过滤条件】 {模版名称Name,支持模糊匹配} {模版类型type，1.系统模版 2.自定义模版} {质量检测维度QualityDims, 1.准确性 2.唯一性 3.完整性 4.一致性 5.及时性 6.有效性} 【排序字段】 { 引用数排序类型CitationOrderType，根据引用数量排序 ASC DESC}
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleTemplatesByPage(request *DescribeRuleTemplatesByPageRequest) (response *DescribeRuleTemplatesByPageResponse, err error) {
    return c.DescribeRuleTemplatesByPageWithContext(context.Background(), request)
}

// DescribeRuleTemplatesByPage
// 过滤条件】 {模版名称Name,支持模糊匹配} {模版类型type，1.系统模版 2.自定义模版} {质量检测维度QualityDims, 1.准确性 2.唯一性 3.完整性 4.一致性 5.及时性 6.有效性} 【排序字段】 { 引用数排序类型CitationOrderType，根据引用数量排序 ASC DESC}
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRuleTemplatesByPageWithContext(ctx context.Context, request *DescribeRuleTemplatesByPageRequest) (response *DescribeRuleTemplatesByPageResponse, err error) {
    if request == nil {
        request = NewDescribeRuleTemplatesByPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleTemplatesByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleTemplatesByPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRules")
    
    
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRules
// 查询质量规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

// DescribeRules
// 查询质量规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesByPageRequest() (request *DescribeRulesByPageRequest) {
    request = &DescribeRulesByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeRulesByPage")
    
    
    return
}

func NewDescribeRulesByPageResponse() (response *DescribeRulesByPageResponse) {
    response = &DescribeRulesByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRulesByPage
// 分页查询质量规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRulesByPage(request *DescribeRulesByPageRequest) (response *DescribeRulesByPageResponse, err error) {
    return c.DescribeRulesByPageWithContext(context.Background(), request)
}

// DescribeRulesByPage
// 分页查询质量规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeRulesByPageWithContext(ctx context.Context, request *DescribeRulesByPageRequest) (response *DescribeRulesByPageResponse, err error) {
    if request == nil {
        request = NewDescribeRulesByPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRulesByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesByPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScheduleInstanceRequest() (request *DescribeScheduleInstanceRequest) {
    request = &DescribeScheduleInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeScheduleInstance")
    
    
    return
}

func NewDescribeScheduleInstanceResponse() (response *DescribeScheduleInstanceResponse) {
    response = &DescribeScheduleInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScheduleInstance
// 在基线系统内查询单个调度任务实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScheduleInstance(request *DescribeScheduleInstanceRequest) (response *DescribeScheduleInstanceResponse, err error) {
    return c.DescribeScheduleInstanceWithContext(context.Background(), request)
}

// DescribeScheduleInstance
// 在基线系统内查询单个调度任务实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeScheduleInstanceWithContext(ctx context.Context, request *DescribeScheduleInstanceRequest) (response *DescribeScheduleInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeScheduleInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScheduleInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScheduleInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScheduleInstancesRequest() (request *DescribeScheduleInstancesRequest) {
    request = &DescribeScheduleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeScheduleInstances")
    
    
    return
}

func NewDescribeScheduleInstancesResponse() (response *DescribeScheduleInstancesResponse) {
    response = &DescribeScheduleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScheduleInstances
// 获取实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeScheduleInstances(request *DescribeScheduleInstancesRequest) (response *DescribeScheduleInstancesResponse, err error) {
    return c.DescribeScheduleInstancesWithContext(context.Background(), request)
}

// DescribeScheduleInstances
// 获取实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeScheduleInstancesWithContext(ctx context.Context, request *DescribeScheduleInstancesRequest) (response *DescribeScheduleInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeScheduleInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScheduleInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScheduleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSchedulerInstanceStatusRequest() (request *DescribeSchedulerInstanceStatusRequest) {
    request = &DescribeSchedulerInstanceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeSchedulerInstanceStatus")
    
    
    return
}

func NewDescribeSchedulerInstanceStatusResponse() (response *DescribeSchedulerInstanceStatusResponse) {
    response = &DescribeSchedulerInstanceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSchedulerInstanceStatus
// 运维大屏-实例状态分布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSchedulerInstanceStatus(request *DescribeSchedulerInstanceStatusRequest) (response *DescribeSchedulerInstanceStatusResponse, err error) {
    return c.DescribeSchedulerInstanceStatusWithContext(context.Background(), request)
}

// DescribeSchedulerInstanceStatus
// 运维大屏-实例状态分布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSchedulerInstanceStatusWithContext(ctx context.Context, request *DescribeSchedulerInstanceStatusRequest) (response *DescribeSchedulerInstanceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSchedulerInstanceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSchedulerInstanceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSchedulerInstanceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSchedulerRunTimeInstanceCntByStatusRequest() (request *DescribeSchedulerRunTimeInstanceCntByStatusRequest) {
    request = &DescribeSchedulerRunTimeInstanceCntByStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeSchedulerRunTimeInstanceCntByStatus")
    
    
    return
}

func NewDescribeSchedulerRunTimeInstanceCntByStatusResponse() (response *DescribeSchedulerRunTimeInstanceCntByStatusResponse) {
    response = &DescribeSchedulerRunTimeInstanceCntByStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSchedulerRunTimeInstanceCntByStatus
// 运维大屏-实例运行时长排行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSchedulerRunTimeInstanceCntByStatus(request *DescribeSchedulerRunTimeInstanceCntByStatusRequest) (response *DescribeSchedulerRunTimeInstanceCntByStatusResponse, err error) {
    return c.DescribeSchedulerRunTimeInstanceCntByStatusWithContext(context.Background(), request)
}

// DescribeSchedulerRunTimeInstanceCntByStatus
// 运维大屏-实例运行时长排行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSchedulerRunTimeInstanceCntByStatusWithContext(ctx context.Context, request *DescribeSchedulerRunTimeInstanceCntByStatusRequest) (response *DescribeSchedulerRunTimeInstanceCntByStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSchedulerRunTimeInstanceCntByStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSchedulerRunTimeInstanceCntByStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSchedulerRunTimeInstanceCntByStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSchedulerTaskCntByStatusRequest() (request *DescribeSchedulerTaskCntByStatusRequest) {
    request = &DescribeSchedulerTaskCntByStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeSchedulerTaskCntByStatus")
    
    
    return
}

func NewDescribeSchedulerTaskCntByStatusResponse() (response *DescribeSchedulerTaskCntByStatusResponse) {
    response = &DescribeSchedulerTaskCntByStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSchedulerTaskCntByStatus
// 任务状态统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSchedulerTaskCntByStatus(request *DescribeSchedulerTaskCntByStatusRequest) (response *DescribeSchedulerTaskCntByStatusResponse, err error) {
    return c.DescribeSchedulerTaskCntByStatusWithContext(context.Background(), request)
}

// DescribeSchedulerTaskCntByStatus
// 任务状态统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSchedulerTaskCntByStatusWithContext(ctx context.Context, request *DescribeSchedulerTaskCntByStatusRequest) (response *DescribeSchedulerTaskCntByStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSchedulerTaskCntByStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSchedulerTaskCntByStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSchedulerTaskCntByStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSchedulerTaskTypeCntRequest() (request *DescribeSchedulerTaskTypeCntRequest) {
    request = &DescribeSchedulerTaskTypeCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeSchedulerTaskTypeCnt")
    
    
    return
}

func NewDescribeSchedulerTaskTypeCntResponse() (response *DescribeSchedulerTaskTypeCntResponse) {
    response = &DescribeSchedulerTaskTypeCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSchedulerTaskTypeCnt
// 运维大屏-任务状态分布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSchedulerTaskTypeCnt(request *DescribeSchedulerTaskTypeCntRequest) (response *DescribeSchedulerTaskTypeCntResponse, err error) {
    return c.DescribeSchedulerTaskTypeCntWithContext(context.Background(), request)
}

// DescribeSchedulerTaskTypeCnt
// 运维大屏-任务状态分布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSchedulerTaskTypeCntWithContext(ctx context.Context, request *DescribeSchedulerTaskTypeCntRequest) (response *DescribeSchedulerTaskTypeCntResponse, err error) {
    if request == nil {
        request = NewDescribeSchedulerTaskTypeCntRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSchedulerTaskTypeCnt require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSchedulerTaskTypeCntResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSonInstancesRequest() (request *DescribeSonInstancesRequest) {
    request = &DescribeSonInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeSonInstances")
    
    
    return
}

func NewDescribeSonInstancesResponse() (response *DescribeSonInstancesResponse) {
    response = &DescribeSonInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSonInstances
// 获取关联子实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSonInstances(request *DescribeSonInstancesRequest) (response *DescribeSonInstancesResponse, err error) {
    return c.DescribeSonInstancesWithContext(context.Background(), request)
}

// DescribeSonInstances
// 获取关联子实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSonInstancesWithContext(ctx context.Context, request *DescribeSonInstancesRequest) (response *DescribeSonInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeSonInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSonInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSonInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStandardRuleDetailInfoListRequest() (request *DescribeStandardRuleDetailInfoListRequest) {
    request = &DescribeStandardRuleDetailInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeStandardRuleDetailInfoList")
    
    
    return
}

func NewDescribeStandardRuleDetailInfoListResponse() (response *DescribeStandardRuleDetailInfoListResponse) {
    response = &DescribeStandardRuleDetailInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStandardRuleDetailInfoList
// 获取数据标准规则详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStandardRuleDetailInfoList(request *DescribeStandardRuleDetailInfoListRequest) (response *DescribeStandardRuleDetailInfoListResponse, err error) {
    return c.DescribeStandardRuleDetailInfoListWithContext(context.Background(), request)
}

// DescribeStandardRuleDetailInfoList
// 获取数据标准规则详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStandardRuleDetailInfoListWithContext(ctx context.Context, request *DescribeStandardRuleDetailInfoListRequest) (response *DescribeStandardRuleDetailInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStandardRuleDetailInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStandardRuleDetailInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStandardRuleDetailInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStatisticInstanceStatusTrendOpsRequest() (request *DescribeStatisticInstanceStatusTrendOpsRequest) {
    request = &DescribeStatisticInstanceStatusTrendOpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeStatisticInstanceStatusTrendOps")
    
    
    return
}

func NewDescribeStatisticInstanceStatusTrendOpsResponse() (response *DescribeStatisticInstanceStatusTrendOpsResponse) {
    response = &DescribeStatisticInstanceStatusTrendOpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStatisticInstanceStatusTrendOps
// 任务状态趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStatisticInstanceStatusTrendOps(request *DescribeStatisticInstanceStatusTrendOpsRequest) (response *DescribeStatisticInstanceStatusTrendOpsResponse, err error) {
    return c.DescribeStatisticInstanceStatusTrendOpsWithContext(context.Background(), request)
}

// DescribeStatisticInstanceStatusTrendOps
// 任务状态趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStatisticInstanceStatusTrendOpsWithContext(ctx context.Context, request *DescribeStatisticInstanceStatusTrendOpsRequest) (response *DescribeStatisticInstanceStatusTrendOpsResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticInstanceStatusTrendOpsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStatisticInstanceStatusTrendOps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStatisticInstanceStatusTrendOpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamTaskLogListRequest() (request *DescribeStreamTaskLogListRequest) {
    request = &DescribeStreamTaskLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeStreamTaskLogList")
    
    
    return
}

func NewDescribeStreamTaskLogListResponse() (response *DescribeStreamTaskLogListResponse) {
    response = &DescribeStreamTaskLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamTaskLogList
// 查询实时任务日志列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALCALLCLOUDAPIERROR = "InternalError.InternalCallCloudApiError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStreamTaskLogList(request *DescribeStreamTaskLogListRequest) (response *DescribeStreamTaskLogListResponse, err error) {
    return c.DescribeStreamTaskLogListWithContext(context.Background(), request)
}

// DescribeStreamTaskLogList
// 查询实时任务日志列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALCALLCLOUDAPIERROR = "InternalError.InternalCallCloudApiError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStreamTaskLogListWithContext(ctx context.Context, request *DescribeStreamTaskLogListRequest) (response *DescribeStreamTaskLogListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamTaskLogListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamTaskLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamTaskLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSuccessorOpsTaskInfosRequest() (request *DescribeSuccessorOpsTaskInfosRequest) {
    request = &DescribeSuccessorOpsTaskInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeSuccessorOpsTaskInfos")
    
    
    return
}

func NewDescribeSuccessorOpsTaskInfosResponse() (response *DescribeSuccessorOpsTaskInfosResponse) {
    response = &DescribeSuccessorOpsTaskInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSuccessorOpsTaskInfos
// 获取下游任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALCALLCLOUDAPIERROR = "InternalError.InternalCallCloudApiError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSuccessorOpsTaskInfos(request *DescribeSuccessorOpsTaskInfosRequest) (response *DescribeSuccessorOpsTaskInfosResponse, err error) {
    return c.DescribeSuccessorOpsTaskInfosWithContext(context.Background(), request)
}

// DescribeSuccessorOpsTaskInfos
// 获取下游任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALCALLCLOUDAPIERROR = "InternalError.InternalCallCloudApiError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSuccessorOpsTaskInfosWithContext(ctx context.Context, request *DescribeSuccessorOpsTaskInfosRequest) (response *DescribeSuccessorOpsTaskInfosResponse, err error) {
    if request == nil {
        request = NewDescribeSuccessorOpsTaskInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSuccessorOpsTaskInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSuccessorOpsTaskInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableInfoListRequest() (request *DescribeTableInfoListRequest) {
    request = &DescribeTableInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTableInfoList")
    
    
    return
}

func NewDescribeTableInfoListResponse() (response *DescribeTableInfoListResponse) {
    response = &DescribeTableInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableInfoList
// 获取数据表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTableInfoList(request *DescribeTableInfoListRequest) (response *DescribeTableInfoListResponse, err error) {
    return c.DescribeTableInfoListWithContext(context.Background(), request)
}

// DescribeTableInfoList
// 获取数据表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTableInfoListWithContext(ctx context.Context, request *DescribeTableInfoListRequest) (response *DescribeTableInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeTableInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableLineageRequest() (request *DescribeTableLineageRequest) {
    request = &DescribeTableLineageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTableLineage")
    
    
    return
}

func NewDescribeTableLineageResponse() (response *DescribeTableLineageResponse) {
    response = &DescribeTableLineageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableLineage
// 列出表血缘信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTableLineage(request *DescribeTableLineageRequest) (response *DescribeTableLineageResponse, err error) {
    return c.DescribeTableLineageWithContext(context.Background(), request)
}

// DescribeTableLineage
// 列出表血缘信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTableLineageWithContext(ctx context.Context, request *DescribeTableLineageRequest) (response *DescribeTableLineageResponse, err error) {
    if request == nil {
        request = NewDescribeTableLineageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableLineage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableLineageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableMetasRequest() (request *DescribeTableMetasRequest) {
    request = &DescribeTableMetasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTableMetas")
    
    
    return
}

func NewDescribeTableMetasResponse() (response *DescribeTableMetasResponse) {
    response = &DescribeTableMetasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableMetas
// 获取表元数据list
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTableMetas(request *DescribeTableMetasRequest) (response *DescribeTableMetasResponse, err error) {
    return c.DescribeTableMetasWithContext(context.Background(), request)
}

// DescribeTableMetas
// 获取表元数据list
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTableMetasWithContext(ctx context.Context, request *DescribeTableMetasRequest) (response *DescribeTableMetasResponse, err error) {
    if request == nil {
        request = NewDescribeTableMetasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableMetas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableMetasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableQualityDetailsRequest() (request *DescribeTableQualityDetailsRequest) {
    request = &DescribeTableQualityDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTableQualityDetails")
    
    
    return
}

func NewDescribeTableQualityDetailsResponse() (response *DescribeTableQualityDetailsResponse) {
    response = &DescribeTableQualityDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableQualityDetails
// 质量报告-查询表质量详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTableQualityDetails(request *DescribeTableQualityDetailsRequest) (response *DescribeTableQualityDetailsResponse, err error) {
    return c.DescribeTableQualityDetailsWithContext(context.Background(), request)
}

// DescribeTableQualityDetails
// 质量报告-查询表质量详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTableQualityDetailsWithContext(ctx context.Context, request *DescribeTableQualityDetailsRequest) (response *DescribeTableQualityDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeTableQualityDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableQualityDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableQualityDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableSchemaInfoRequest() (request *DescribeTableSchemaInfoRequest) {
    request = &DescribeTableSchemaInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTableSchemaInfo")
    
    
    return
}

func NewDescribeTableSchemaInfoResponse() (response *DescribeTableSchemaInfoResponse) {
    response = &DescribeTableSchemaInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableSchemaInfo
// 获取表schema信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTableSchemaInfo(request *DescribeTableSchemaInfoRequest) (response *DescribeTableSchemaInfoResponse, err error) {
    return c.DescribeTableSchemaInfoWithContext(context.Background(), request)
}

// DescribeTableSchemaInfo
// 获取表schema信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTableSchemaInfoWithContext(ctx context.Context, request *DescribeTableSchemaInfoRequest) (response *DescribeTableSchemaInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTableSchemaInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableSchemaInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableSchemaInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableScoreTrendRequest() (request *DescribeTableScoreTrendRequest) {
    request = &DescribeTableScoreTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTableScoreTrend")
    
    
    return
}

func NewDescribeTableScoreTrendResponse() (response *DescribeTableScoreTrendResponse) {
    response = &DescribeTableScoreTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableScoreTrend
// 查询表得分趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTableScoreTrend(request *DescribeTableScoreTrendRequest) (response *DescribeTableScoreTrendResponse, err error) {
    return c.DescribeTableScoreTrendWithContext(context.Background(), request)
}

// DescribeTableScoreTrend
// 查询表得分趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTableScoreTrendWithContext(ctx context.Context, request *DescribeTableScoreTrendRequest) (response *DescribeTableScoreTrendResponse, err error) {
    if request == nil {
        request = NewDescribeTableScoreTrendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableScoreTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableScoreTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskAlarmRegulationsRequest() (request *DescribeTaskAlarmRegulationsRequest) {
    request = &DescribeTaskAlarmRegulationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskAlarmRegulations")
    
    
    return
}

func NewDescribeTaskAlarmRegulationsResponse() (response *DescribeTaskAlarmRegulationsResponse) {
    response = &DescribeTaskAlarmRegulationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskAlarmRegulations
// 查询任务告警规则列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTaskAlarmRegulations(request *DescribeTaskAlarmRegulationsRequest) (response *DescribeTaskAlarmRegulationsResponse, err error) {
    return c.DescribeTaskAlarmRegulationsWithContext(context.Background(), request)
}

// DescribeTaskAlarmRegulations
// 查询任务告警规则列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTaskAlarmRegulationsWithContext(ctx context.Context, request *DescribeTaskAlarmRegulationsRequest) (response *DescribeTaskAlarmRegulationsResponse, err error) {
    if request == nil {
        request = NewDescribeTaskAlarmRegulationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskAlarmRegulations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskAlarmRegulationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskByCycleRequest() (request *DescribeTaskByCycleRequest) {
    request = &DescribeTaskByCycleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskByCycle")
    
    
    return
}

func NewDescribeTaskByCycleResponse() (response *DescribeTaskByCycleResponse) {
    response = &DescribeTaskByCycleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskByCycle
// 根据周期类型 查询所有任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskByCycle(request *DescribeTaskByCycleRequest) (response *DescribeTaskByCycleResponse, err error) {
    return c.DescribeTaskByCycleWithContext(context.Background(), request)
}

// DescribeTaskByCycle
// 根据周期类型 查询所有任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskByCycleWithContext(ctx context.Context, request *DescribeTaskByCycleRequest) (response *DescribeTaskByCycleResponse, err error) {
    if request == nil {
        request = NewDescribeTaskByCycleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskByCycle require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskByCycleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskByCycleReportRequest() (request *DescribeTaskByCycleReportRequest) {
    request = &DescribeTaskByCycleReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskByCycleReport")
    
    
    return
}

func NewDescribeTaskByCycleReportResponse() (response *DescribeTaskByCycleReportResponse) {
    response = &DescribeTaskByCycleReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskByCycleReport
// 任务状态周期增长趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskByCycleReport(request *DescribeTaskByCycleReportRequest) (response *DescribeTaskByCycleReportResponse, err error) {
    return c.DescribeTaskByCycleReportWithContext(context.Background(), request)
}

// DescribeTaskByCycleReport
// 任务状态周期增长趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskByCycleReportWithContext(ctx context.Context, request *DescribeTaskByCycleReportRequest) (response *DescribeTaskByCycleReportResponse, err error) {
    if request == nil {
        request = NewDescribeTaskByCycleReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskByCycleReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskByCycleReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskByStatusReportRequest() (request *DescribeTaskByStatusReportRequest) {
    request = &DescribeTaskByStatusReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskByStatusReport")
    
    
    return
}

func NewDescribeTaskByStatusReportResponse() (response *DescribeTaskByStatusReportResponse) {
    response = &DescribeTaskByStatusReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskByStatusReport
// 任务状态趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskByStatusReport(request *DescribeTaskByStatusReportRequest) (response *DescribeTaskByStatusReportResponse, err error) {
    return c.DescribeTaskByStatusReportWithContext(context.Background(), request)
}

// DescribeTaskByStatusReport
// 任务状态趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskByStatusReportWithContext(ctx context.Context, request *DescribeTaskByStatusReportRequest) (response *DescribeTaskByStatusReportResponse, err error) {
    if request == nil {
        request = NewDescribeTaskByStatusReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskByStatusReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskByStatusReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskDetail")
    
    
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskDetail
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 查询任务具体详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    return c.DescribeTaskDetailWithContext(context.Background(), request)
}

// DescribeTaskDetail
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 查询任务具体详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInstanceRequest() (request *DescribeTaskInstanceRequest) {
    request = &DescribeTaskInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskInstance")
    
    
    return
}

func NewDescribeTaskInstanceResponse() (response *DescribeTaskInstanceResponse) {
    response = &DescribeTaskInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskInstance
// 离线任务实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
func (c *Client) DescribeTaskInstance(request *DescribeTaskInstanceRequest) (response *DescribeTaskInstanceResponse, err error) {
    return c.DescribeTaskInstanceWithContext(context.Background(), request)
}

// DescribeTaskInstance
// 离线任务实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
func (c *Client) DescribeTaskInstanceWithContext(ctx context.Context, request *DescribeTaskInstanceRequest) (response *DescribeTaskInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInstanceReportDetailRequest() (request *DescribeTaskInstanceReportDetailRequest) {
    request = &DescribeTaskInstanceReportDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskInstanceReportDetail")
    
    
    return
}

func NewDescribeTaskInstanceReportDetailResponse() (response *DescribeTaskInstanceReportDetailResponse) {
    response = &DescribeTaskInstanceReportDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskInstanceReportDetail
// 离线任务实例统计明细
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskInstanceReportDetail(request *DescribeTaskInstanceReportDetailRequest) (response *DescribeTaskInstanceReportDetailResponse, err error) {
    return c.DescribeTaskInstanceReportDetailWithContext(context.Background(), request)
}

// DescribeTaskInstanceReportDetail
// 离线任务实例统计明细
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskInstanceReportDetailWithContext(ctx context.Context, request *DescribeTaskInstanceReportDetailRequest) (response *DescribeTaskInstanceReportDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInstanceReportDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskInstanceReportDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskInstanceReportDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInstancesRequest() (request *DescribeTaskInstancesRequest) {
    request = &DescribeTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskInstances")
    
    
    return
}

func NewDescribeTaskInstancesResponse() (response *DescribeTaskInstancesResponse) {
    response = &DescribeTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskInstances
// 查询任务实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskInstances(request *DescribeTaskInstancesRequest) (response *DescribeTaskInstancesResponse, err error) {
    return c.DescribeTaskInstancesWithContext(context.Background(), request)
}

// DescribeTaskInstances
// 查询任务实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskInstancesWithContext(ctx context.Context, request *DescribeTaskInstancesRequest) (response *DescribeTaskInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLockStatusRequest() (request *DescribeTaskLockStatusRequest) {
    request = &DescribeTaskLockStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskLockStatus")
    
    
    return
}

func NewDescribeTaskLockStatusResponse() (response *DescribeTaskLockStatusResponse) {
    response = &DescribeTaskLockStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskLockStatus
// 查看任务锁状态信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTaskLockStatus(request *DescribeTaskLockStatusRequest) (response *DescribeTaskLockStatusResponse, err error) {
    return c.DescribeTaskLockStatusWithContext(context.Background(), request)
}

// DescribeTaskLockStatus
// 查看任务锁状态信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTaskLockStatusWithContext(ctx context.Context, request *DescribeTaskLockStatusRequest) (response *DescribeTaskLockStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLockStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskLockStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskLockStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskReportRequest() (request *DescribeTaskReportRequest) {
    request = &DescribeTaskReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskReport")
    
    
    return
}

func NewDescribeTaskReportResponse() (response *DescribeTaskReportResponse) {
    response = &DescribeTaskReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskReport
// 按起止日期统计离线任务的所有实例的运行指标总和
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskReport(request *DescribeTaskReportRequest) (response *DescribeTaskReportResponse, err error) {
    return c.DescribeTaskReportWithContext(context.Background(), request)
}

// DescribeTaskReport
// 按起止日期统计离线任务的所有实例的运行指标总和
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskReportWithContext(ctx context.Context, request *DescribeTaskReportRequest) (response *DescribeTaskReportResponse, err error) {
    if request == nil {
        request = NewDescribeTaskReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskReportDetailListRequest() (request *DescribeTaskReportDetailListRequest) {
    request = &DescribeTaskReportDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskReportDetailList")
    
    
    return
}

func NewDescribeTaskReportDetailListResponse() (response *DescribeTaskReportDetailListResponse) {
    response = &DescribeTaskReportDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskReportDetailList
// 离线任务周期统计明细
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskReportDetailList(request *DescribeTaskReportDetailListRequest) (response *DescribeTaskReportDetailListResponse, err error) {
    return c.DescribeTaskReportDetailListWithContext(context.Background(), request)
}

// DescribeTaskReportDetailList
// 离线任务周期统计明细
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskReportDetailListWithContext(ctx context.Context, request *DescribeTaskReportDetailListRequest) (response *DescribeTaskReportDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeTaskReportDetailListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskReportDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskReportDetailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRunHistoryRequest() (request *DescribeTaskRunHistoryRequest) {
    request = &DescribeTaskRunHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskRunHistory")
    
    
    return
}

func NewDescribeTaskRunHistoryResponse() (response *DescribeTaskRunHistoryResponse) {
    response = &DescribeTaskRunHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskRunHistory
// 分页查询任务运行历史
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskRunHistory(request *DescribeTaskRunHistoryRequest) (response *DescribeTaskRunHistoryResponse, err error) {
    return c.DescribeTaskRunHistoryWithContext(context.Background(), request)
}

// DescribeTaskRunHistory
// 分页查询任务运行历史
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskRunHistoryWithContext(ctx context.Context, request *DescribeTaskRunHistoryRequest) (response *DescribeTaskRunHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRunHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskRunHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskRunHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskScriptRequest() (request *DescribeTaskScriptRequest) {
    request = &DescribeTaskScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTaskScript")
    
    
    return
}

func NewDescribeTaskScriptResponse() (response *DescribeTaskScriptResponse) {
    response = &DescribeTaskScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskScript
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 查询任务脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskScript(request *DescribeTaskScriptRequest) (response *DescribeTaskScriptResponse, err error) {
    return c.DescribeTaskScriptWithContext(context.Background(), request)
}

// DescribeTaskScript
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 查询任务脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTaskScriptWithContext(ctx context.Context, request *DescribeTaskScriptRequest) (response *DescribeTaskScriptResponse, err error) {
    if request == nil {
        request = NewDescribeTaskScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksByPageRequest() (request *DescribeTasksByPageRequest) {
    request = &DescribeTasksByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTasksByPage")
    
    
    return
}

func NewDescribeTasksByPageResponse() (response *DescribeTasksByPageResponse) {
    response = &DescribeTasksByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTasksByPage
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 根据工作流分页查询任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTasksByPage(request *DescribeTasksByPageRequest) (response *DescribeTasksByPageResponse, err error) {
    return c.DescribeTasksByPageWithContext(context.Background(), request)
}

// DescribeTasksByPage
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 根据工作流分页查询任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTasksByPageWithContext(ctx context.Context, request *DescribeTasksByPageRequest) (response *DescribeTasksByPageResponse, err error) {
    if request == nil {
        request = NewDescribeTasksByPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasksByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksByPageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateDimCountRequest() (request *DescribeTemplateDimCountRequest) {
    request = &DescribeTemplateDimCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTemplateDimCount")
    
    
    return
}

func NewDescribeTemplateDimCountResponse() (response *DescribeTemplateDimCountResponse) {
    response = &DescribeTemplateDimCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTemplateDimCount
// 查询规则模板维度分布情况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeTemplateDimCount(request *DescribeTemplateDimCountRequest) (response *DescribeTemplateDimCountResponse, err error) {
    return c.DescribeTemplateDimCountWithContext(context.Background(), request)
}

// DescribeTemplateDimCount
// 查询规则模板维度分布情况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeTemplateDimCountWithContext(ctx context.Context, request *DescribeTemplateDimCountRequest) (response *DescribeTemplateDimCountResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateDimCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTemplateDimCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTemplateDimCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateHistoryRequest() (request *DescribeTemplateHistoryRequest) {
    request = &DescribeTemplateHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTemplateHistory")
    
    
    return
}

func NewDescribeTemplateHistoryResponse() (response *DescribeTemplateHistoryResponse) {
    response = &DescribeTemplateHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTemplateHistory
// 查询规则模板操作记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
func (c *Client) DescribeTemplateHistory(request *DescribeTemplateHistoryRequest) (response *DescribeTemplateHistoryResponse, err error) {
    return c.DescribeTemplateHistoryWithContext(context.Background(), request)
}

// DescribeTemplateHistory
// 查询规则模板操作记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
func (c *Client) DescribeTemplateHistoryWithContext(ctx context.Context, request *DescribeTemplateHistoryRequest) (response *DescribeTemplateHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTemplateHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTemplateHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeThirdTaskRunLogRequest() (request *DescribeThirdTaskRunLogRequest) {
    request = &DescribeThirdTaskRunLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeThirdTaskRunLog")
    
    
    return
}

func NewDescribeThirdTaskRunLogResponse() (response *DescribeThirdTaskRunLogResponse) {
    response = &DescribeThirdTaskRunLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeThirdTaskRunLog
// 获取第三方运行日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeThirdTaskRunLog(request *DescribeThirdTaskRunLogRequest) (response *DescribeThirdTaskRunLogResponse, err error) {
    return c.DescribeThirdTaskRunLogWithContext(context.Background(), request)
}

// DescribeThirdTaskRunLog
// 获取第三方运行日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeThirdTaskRunLogWithContext(ctx context.Context, request *DescribeThirdTaskRunLogRequest) (response *DescribeThirdTaskRunLogResponse, err error) {
    if request == nil {
        request = NewDescribeThirdTaskRunLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeThirdTaskRunLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeThirdTaskRunLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopTableStatRequest() (request *DescribeTopTableStatRequest) {
    request = &DescribeTopTableStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTopTableStat")
    
    
    return
}

func NewDescribeTopTableStatResponse() (response *DescribeTopTableStatResponse) {
    response = &DescribeTopTableStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopTableStat
// 数据质量概览页面表排行接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeTopTableStat(request *DescribeTopTableStatRequest) (response *DescribeTopTableStatResponse, err error) {
    return c.DescribeTopTableStatWithContext(context.Background(), request)
}

// DescribeTopTableStat
// 数据质量概览页面表排行接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeTopTableStatWithContext(ctx context.Context, request *DescribeTopTableStatRequest) (response *DescribeTopTableStatResponse, err error) {
    if request == nil {
        request = NewDescribeTopTableStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopTableStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopTableStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrendStatRequest() (request *DescribeTrendStatRequest) {
    request = &DescribeTrendStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeTrendStat")
    
    
    return
}

func NewDescribeTrendStatResponse() (response *DescribeTrendStatResponse) {
    response = &DescribeTrendStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrendStat
// 数据质量概览页面趋势变化接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeTrendStat(request *DescribeTrendStatRequest) (response *DescribeTrendStatResponse, err error) {
    return c.DescribeTrendStatWithContext(context.Background(), request)
}

// DescribeTrendStat
// 数据质量概览页面趋势变化接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) DescribeTrendStatWithContext(ctx context.Context, request *DescribeTrendStatRequest) (response *DescribeTrendStatResponse, err error) {
    if request == nil {
        request = NewDescribeTrendStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrendStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrendStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkflowCanvasInfoRequest() (request *DescribeWorkflowCanvasInfoRequest) {
    request = &DescribeWorkflowCanvasInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeWorkflowCanvasInfo")
    
    
    return
}

func NewDescribeWorkflowCanvasInfoResponse() (response *DescribeWorkflowCanvasInfoResponse) {
    response = &DescribeWorkflowCanvasInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkflowCanvasInfo
// 查询工作流画布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowCanvasInfo(request *DescribeWorkflowCanvasInfoRequest) (response *DescribeWorkflowCanvasInfoResponse, err error) {
    return c.DescribeWorkflowCanvasInfoWithContext(context.Background(), request)
}

// DescribeWorkflowCanvasInfo
// 查询工作流画布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowCanvasInfoWithContext(ctx context.Context, request *DescribeWorkflowCanvasInfoRequest) (response *DescribeWorkflowCanvasInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowCanvasInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkflowCanvasInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkflowCanvasInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkflowExecuteByIdRequest() (request *DescribeWorkflowExecuteByIdRequest) {
    request = &DescribeWorkflowExecuteByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeWorkflowExecuteById")
    
    
    return
}

func NewDescribeWorkflowExecuteByIdResponse() (response *DescribeWorkflowExecuteByIdResponse) {
    response = &DescribeWorkflowExecuteByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkflowExecuteById
// 查询工作流画布运行起止时间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowExecuteById(request *DescribeWorkflowExecuteByIdRequest) (response *DescribeWorkflowExecuteByIdResponse, err error) {
    return c.DescribeWorkflowExecuteByIdWithContext(context.Background(), request)
}

// DescribeWorkflowExecuteById
// 查询工作流画布运行起止时间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowExecuteByIdWithContext(ctx context.Context, request *DescribeWorkflowExecuteByIdRequest) (response *DescribeWorkflowExecuteByIdResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowExecuteByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkflowExecuteById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkflowExecuteByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkflowInfoByIdRequest() (request *DescribeWorkflowInfoByIdRequest) {
    request = &DescribeWorkflowInfoByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeWorkflowInfoById")
    
    
    return
}

func NewDescribeWorkflowInfoByIdResponse() (response *DescribeWorkflowInfoByIdResponse) {
    response = &DescribeWorkflowInfoByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkflowInfoById
// 通过工作流id，查询工作流详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowInfoById(request *DescribeWorkflowInfoByIdRequest) (response *DescribeWorkflowInfoByIdResponse, err error) {
    return c.DescribeWorkflowInfoByIdWithContext(context.Background(), request)
}

// DescribeWorkflowInfoById
// 通过工作流id，查询工作流详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowInfoByIdWithContext(ctx context.Context, request *DescribeWorkflowInfoByIdRequest) (response *DescribeWorkflowInfoByIdResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowInfoByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkflowInfoById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkflowInfoByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkflowListByProjectIdRequest() (request *DescribeWorkflowListByProjectIdRequest) {
    request = &DescribeWorkflowListByProjectIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeWorkflowListByProjectId")
    
    
    return
}

func NewDescribeWorkflowListByProjectIdResponse() (response *DescribeWorkflowListByProjectIdResponse) {
    response = &DescribeWorkflowListByProjectIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkflowListByProjectId
// 根据项目id 获取项目下所有工作流列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowListByProjectId(request *DescribeWorkflowListByProjectIdRequest) (response *DescribeWorkflowListByProjectIdResponse, err error) {
    return c.DescribeWorkflowListByProjectIdWithContext(context.Background(), request)
}

// DescribeWorkflowListByProjectId
// 根据项目id 获取项目下所有工作流列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowListByProjectIdWithContext(ctx context.Context, request *DescribeWorkflowListByProjectIdRequest) (response *DescribeWorkflowListByProjectIdResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowListByProjectIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkflowListByProjectId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkflowListByProjectIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkflowOpsCanvasInfoRequest() (request *DescribeWorkflowOpsCanvasInfoRequest) {
    request = &DescribeWorkflowOpsCanvasInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeWorkflowOpsCanvasInfo")
    
    
    return
}

func NewDescribeWorkflowOpsCanvasInfoResponse() (response *DescribeWorkflowOpsCanvasInfoResponse) {
    response = &DescribeWorkflowOpsCanvasInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkflowOpsCanvasInfo
// 查询运维画布信息，只需要获取边和节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowOpsCanvasInfo(request *DescribeWorkflowOpsCanvasInfoRequest) (response *DescribeWorkflowOpsCanvasInfoResponse, err error) {
    return c.DescribeWorkflowOpsCanvasInfoWithContext(context.Background(), request)
}

// DescribeWorkflowOpsCanvasInfo
// 查询运维画布信息，只需要获取边和节点
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowOpsCanvasInfoWithContext(ctx context.Context, request *DescribeWorkflowOpsCanvasInfoRequest) (response *DescribeWorkflowOpsCanvasInfoResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowOpsCanvasInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkflowOpsCanvasInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkflowOpsCanvasInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkflowTaskCountRequest() (request *DescribeWorkflowTaskCountRequest) {
    request = &DescribeWorkflowTaskCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeWorkflowTaskCount")
    
    
    return
}

func NewDescribeWorkflowTaskCountResponse() (response *DescribeWorkflowTaskCountResponse) {
    response = &DescribeWorkflowTaskCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkflowTaskCount
// 查询工作流任务数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowTaskCount(request *DescribeWorkflowTaskCountRequest) (response *DescribeWorkflowTaskCountResponse, err error) {
    return c.DescribeWorkflowTaskCountWithContext(context.Background(), request)
}

// DescribeWorkflowTaskCount
// 查询工作流任务数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeWorkflowTaskCountWithContext(ctx context.Context, request *DescribeWorkflowTaskCountRequest) (response *DescribeWorkflowTaskCountResponse, err error) {
    if request == nil {
        request = NewDescribeWorkflowTaskCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkflowTaskCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkflowTaskCountResponse()
    err = c.Send(request, response)
    return
}

func NewDiagnosePlusRequest() (request *DiagnosePlusRequest) {
    request = &DiagnosePlusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DiagnosePlus")
    
    
    return
}

func NewDiagnosePlusResponse() (response *DiagnosePlusResponse) {
    response = &DiagnosePlusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DiagnosePlus
// 实例诊断信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DiagnosePlus(request *DiagnosePlusRequest) (response *DiagnosePlusResponse, err error) {
    return c.DiagnosePlusWithContext(context.Background(), request)
}

// DiagnosePlus
// 实例诊断信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DiagnosePlusWithContext(ctx context.Context, request *DiagnosePlusRequest) (response *DiagnosePlusResponse, err error) {
    if request == nil {
        request = NewDiagnosePlusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DiagnosePlus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDiagnosePlusResponse()
    err = c.Send(request, response)
    return
}

func NewDryRunDIOfflineTaskRequest() (request *DryRunDIOfflineTaskRequest) {
    request = &DryRunDIOfflineTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DryRunDIOfflineTask")
    
    
    return
}

func NewDryRunDIOfflineTaskResponse() (response *DryRunDIOfflineTaskResponse) {
    response = &DryRunDIOfflineTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DryRunDIOfflineTask
// 调试运行集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DryRunDIOfflineTask(request *DryRunDIOfflineTaskRequest) (response *DryRunDIOfflineTaskResponse, err error) {
    return c.DryRunDIOfflineTaskWithContext(context.Background(), request)
}

// DryRunDIOfflineTask
// 调试运行集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DryRunDIOfflineTaskWithContext(ctx context.Context, request *DryRunDIOfflineTaskRequest) (response *DryRunDIOfflineTaskResponse, err error) {
    if request == nil {
        request = NewDryRunDIOfflineTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DryRunDIOfflineTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDryRunDIOfflineTaskResponse()
    err = c.Send(request, response)
    return
}

func NewEditBaselineRequest() (request *EditBaselineRequest) {
    request = &EditBaselineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "EditBaseline")
    
    
    return
}

func NewEditBaselineResponse() (response *EditBaselineResponse) {
    response = &EditBaselineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EditBaseline
// 编辑基线
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditBaseline(request *EditBaselineRequest) (response *EditBaselineResponse, err error) {
    return c.EditBaselineWithContext(context.Background(), request)
}

// EditBaseline
// 编辑基线
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EditBaselineWithContext(ctx context.Context, request *EditBaselineRequest) (response *EditBaselineResponse, err error) {
    if request == nil {
        request = NewEditBaselineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EditBaseline require credential")
    }

    request.SetContext(ctx)
    
    response = NewEditBaselineResponse()
    err = c.Send(request, response)
    return
}

func NewFindAllFolderRequest() (request *FindAllFolderRequest) {
    request = &FindAllFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "FindAllFolder")
    
    
    return
}

func NewFindAllFolderResponse() (response *FindAllFolderResponse) {
    response = &FindAllFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FindAllFolder
// 查找全部的文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FindAllFolder(request *FindAllFolderRequest) (response *FindAllFolderResponse, err error) {
    return c.FindAllFolderWithContext(context.Background(), request)
}

// FindAllFolder
// 查找全部的文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FindAllFolderWithContext(ctx context.Context, request *FindAllFolderRequest) (response *FindAllFolderResponse, err error) {
    if request == nil {
        request = NewFindAllFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FindAllFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewFindAllFolderResponse()
    err = c.Send(request, response)
    return
}

func NewForceSucInstancesRequest() (request *ForceSucInstancesRequest) {
    request = &ForceSucInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ForceSucInstances")
    
    
    return
}

func NewForceSucInstancesResponse() (response *ForceSucInstancesResponse) {
    response = &ForceSucInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ForceSucInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量置成功
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ForceSucInstances(request *ForceSucInstancesRequest) (response *ForceSucInstancesResponse, err error) {
    return c.ForceSucInstancesWithContext(context.Background(), request)
}

// ForceSucInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量置成功
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ForceSucInstancesWithContext(ctx context.Context, request *ForceSucInstancesRequest) (response *ForceSucInstancesResponse, err error) {
    if request == nil {
        request = NewForceSucInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForceSucInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewForceSucInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewForceSucScheduleInstancesRequest() (request *ForceSucScheduleInstancesRequest) {
    request = &ForceSucScheduleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ForceSucScheduleInstances")
    
    
    return
}

func NewForceSucScheduleInstancesResponse() (response *ForceSucScheduleInstancesResponse) {
    response = &ForceSucScheduleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ForceSucScheduleInstances
// 实例强制成功
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ForceSucScheduleInstances(request *ForceSucScheduleInstancesRequest) (response *ForceSucScheduleInstancesResponse, err error) {
    return c.ForceSucScheduleInstancesWithContext(context.Background(), request)
}

// ForceSucScheduleInstances
// 实例强制成功
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ForceSucScheduleInstancesWithContext(ctx context.Context, request *ForceSucScheduleInstancesRequest) (response *ForceSucScheduleInstancesResponse, err error) {
    if request == nil {
        request = NewForceSucScheduleInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ForceSucScheduleInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewForceSucScheduleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewFreezeOpsTasksRequest() (request *FreezeOpsTasksRequest) {
    request = &FreezeOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "FreezeOpsTasks")
    
    
    return
}

func NewFreezeOpsTasksResponse() (response *FreezeOpsTasksResponse) {
    response = &FreezeOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FreezeOpsTasks
// 任务运维-批量冻结任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeOpsTasks(request *FreezeOpsTasksRequest) (response *FreezeOpsTasksResponse, err error) {
    return c.FreezeOpsTasksWithContext(context.Background(), request)
}

// FreezeOpsTasks
// 任务运维-批量冻结任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeOpsTasksWithContext(ctx context.Context, request *FreezeOpsTasksRequest) (response *FreezeOpsTasksResponse, err error) {
    if request == nil {
        request = NewFreezeOpsTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FreezeOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewFreezeOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewFreezeTasksRequest() (request *FreezeTasksRequest) {
    request = &FreezeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "FreezeTasks")
    
    
    return
}

func NewFreezeTasksResponse() (response *FreezeTasksResponse) {
    response = &FreezeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FreezeTasks
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量冻结任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeTasks(request *FreezeTasksRequest) (response *FreezeTasksResponse, err error) {
    return c.FreezeTasksWithContext(context.Background(), request)
}

// FreezeTasks
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 批量冻结任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeTasksWithContext(ctx context.Context, request *FreezeTasksRequest) (response *FreezeTasksResponse, err error) {
    if request == nil {
        request = NewFreezeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FreezeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewFreezeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewFreezeTasksByMultiWorkflowRequest() (request *FreezeTasksByMultiWorkflowRequest) {
    request = &FreezeTasksByMultiWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "FreezeTasksByMultiWorkflow")
    
    
    return
}

func NewFreezeTasksByMultiWorkflowResponse() (response *FreezeTasksByMultiWorkflowResponse) {
    response = &FreezeTasksByMultiWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FreezeTasksByMultiWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 基于多个工作流进行批量冻结任务操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeTasksByMultiWorkflow(request *FreezeTasksByMultiWorkflowRequest) (response *FreezeTasksByMultiWorkflowResponse, err error) {
    return c.FreezeTasksByMultiWorkflowWithContext(context.Background(), request)
}

// FreezeTasksByMultiWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 基于多个工作流进行批量冻结任务操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeTasksByMultiWorkflowWithContext(ctx context.Context, request *FreezeTasksByMultiWorkflowRequest) (response *FreezeTasksByMultiWorkflowResponse, err error) {
    if request == nil {
        request = NewFreezeTasksByMultiWorkflowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FreezeTasksByMultiWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewFreezeTasksByMultiWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewFreezeTasksByWorkflowIdsRequest() (request *FreezeTasksByWorkflowIdsRequest) {
    request = &FreezeTasksByWorkflowIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "FreezeTasksByWorkflowIds")
    
    
    return
}

func NewFreezeTasksByWorkflowIdsResponse() (response *FreezeTasksByWorkflowIdsResponse) {
    response = &FreezeTasksByWorkflowIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FreezeTasksByWorkflowIds
// 暂停工作流下的所有任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeTasksByWorkflowIds(request *FreezeTasksByWorkflowIdsRequest) (response *FreezeTasksByWorkflowIdsResponse, err error) {
    return c.FreezeTasksByWorkflowIdsWithContext(context.Background(), request)
}

// FreezeTasksByWorkflowIds
// 暂停工作流下的所有任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FreezeTasksByWorkflowIdsWithContext(ctx context.Context, request *FreezeTasksByWorkflowIdsRequest) (response *FreezeTasksByWorkflowIdsResponse, err error) {
    if request == nil {
        request = NewFreezeTasksByWorkflowIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FreezeTasksByWorkflowIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewFreezeTasksByWorkflowIdsResponse()
    err = c.Send(request, response)
    return
}

func NewGenHiveTableDDLSqlRequest() (request *GenHiveTableDDLSqlRequest) {
    request = &GenHiveTableDDLSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GenHiveTableDDLSql")
    
    
    return
}

func NewGenHiveTableDDLSqlResponse() (response *GenHiveTableDDLSqlResponse) {
    response = &GenHiveTableDDLSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenHiveTableDDLSql
// 生成建hive表的sql
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GenHiveTableDDLSql(request *GenHiveTableDDLSqlRequest) (response *GenHiveTableDDLSqlResponse, err error) {
    return c.GenHiveTableDDLSqlWithContext(context.Background(), request)
}

// GenHiveTableDDLSql
// 生成建hive表的sql
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GenHiveTableDDLSqlWithContext(ctx context.Context, request *GenHiveTableDDLSqlRequest) (response *GenHiveTableDDLSqlResponse, err error) {
    if request == nil {
        request = NewGenHiveTableDDLSqlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenHiveTableDDLSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenHiveTableDDLSqlResponse()
    err = c.Send(request, response)
    return
}

func NewGetFileInfoRequest() (request *GetFileInfoRequest) {
    request = &GetFileInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetFileInfo")
    
    
    return
}

func NewGetFileInfoResponse() (response *GetFileInfoResponse) {
    response = &GetFileInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFileInfo
// 开发空间-获取数据开发脚本信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetFileInfo(request *GetFileInfoRequest) (response *GetFileInfoResponse, err error) {
    return c.GetFileInfoWithContext(context.Background(), request)
}

// GetFileInfo
// 开发空间-获取数据开发脚本信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetFileInfoWithContext(ctx context.Context, request *GetFileInfoRequest) (response *GetFileInfoResponse, err error) {
    if request == nil {
        request = NewGetFileInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFileInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFileInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetIntegrationNodeColumnSchemaRequest() (request *GetIntegrationNodeColumnSchemaRequest) {
    request = &GetIntegrationNodeColumnSchemaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetIntegrationNodeColumnSchema")
    
    
    return
}

func NewGetIntegrationNodeColumnSchemaResponse() (response *GetIntegrationNodeColumnSchemaResponse) {
    response = &GetIntegrationNodeColumnSchemaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetIntegrationNodeColumnSchema
// 提取数据集成节点字段Schema
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) GetIntegrationNodeColumnSchema(request *GetIntegrationNodeColumnSchemaRequest) (response *GetIntegrationNodeColumnSchemaResponse, err error) {
    return c.GetIntegrationNodeColumnSchemaWithContext(context.Background(), request)
}

// GetIntegrationNodeColumnSchema
// 提取数据集成节点字段Schema
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) GetIntegrationNodeColumnSchemaWithContext(ctx context.Context, request *GetIntegrationNodeColumnSchemaRequest) (response *GetIntegrationNodeColumnSchemaResponse, err error) {
    if request == nil {
        request = NewGetIntegrationNodeColumnSchemaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetIntegrationNodeColumnSchema require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetIntegrationNodeColumnSchemaResponse()
    err = c.Send(request, response)
    return
}

func NewGetOfflineDIInstanceListRequest() (request *GetOfflineDIInstanceListRequest) {
    request = &GetOfflineDIInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOfflineDIInstanceList")
    
    
    return
}

func NewGetOfflineDIInstanceListResponse() (response *GetOfflineDIInstanceListResponse) {
    response = &GetOfflineDIInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOfflineDIInstanceList
// 获取离线任务实例列表(新)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetOfflineDIInstanceList(request *GetOfflineDIInstanceListRequest) (response *GetOfflineDIInstanceListResponse, err error) {
    return c.GetOfflineDIInstanceListWithContext(context.Background(), request)
}

// GetOfflineDIInstanceList
// 获取离线任务实例列表(新)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetOfflineDIInstanceListWithContext(ctx context.Context, request *GetOfflineDIInstanceListRequest) (response *GetOfflineDIInstanceListResponse, err error) {
    if request == nil {
        request = NewGetOfflineDIInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOfflineDIInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOfflineDIInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewGetOfflineInstanceListRequest() (request *GetOfflineInstanceListRequest) {
    request = &GetOfflineInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOfflineInstanceList")
    
    
    return
}

func NewGetOfflineInstanceListResponse() (response *GetOfflineInstanceListResponse) {
    response = &GetOfflineInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOfflineInstanceList
// 获取离线任务实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetOfflineInstanceList(request *GetOfflineInstanceListRequest) (response *GetOfflineInstanceListResponse, err error) {
    return c.GetOfflineInstanceListWithContext(context.Background(), request)
}

// GetOfflineInstanceList
// 获取离线任务实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetOfflineInstanceListWithContext(ctx context.Context, request *GetOfflineInstanceListRequest) (response *GetOfflineInstanceListResponse, err error) {
    if request == nil {
        request = NewGetOfflineInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOfflineInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOfflineInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewKillInstancesRequest() (request *KillInstancesRequest) {
    request = &KillInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "KillInstances")
    
    
    return
}

func NewKillInstancesResponse() (response *KillInstancesResponse) {
    response = &KillInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量终止操作
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) KillInstances(request *KillInstancesRequest) (response *KillInstancesResponse, err error) {
    return c.KillInstancesWithContext(context.Background(), request)
}

// KillInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量终止操作
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) KillInstancesWithContext(ctx context.Context, request *KillInstancesRequest) (response *KillInstancesResponse, err error) {
    if request == nil {
        request = NewKillInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewKillOpsMakePlanInstancesRequest() (request *KillOpsMakePlanInstancesRequest) {
    request = &KillOpsMakePlanInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "KillOpsMakePlanInstances")
    
    
    return
}

func NewKillOpsMakePlanInstancesResponse() (response *KillOpsMakePlanInstancesResponse) {
    response = &KillOpsMakePlanInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillOpsMakePlanInstances
// 按补录计划批量终止实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) KillOpsMakePlanInstances(request *KillOpsMakePlanInstancesRequest) (response *KillOpsMakePlanInstancesResponse, err error) {
    return c.KillOpsMakePlanInstancesWithContext(context.Background(), request)
}

// KillOpsMakePlanInstances
// 按补录计划批量终止实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) KillOpsMakePlanInstancesWithContext(ctx context.Context, request *KillOpsMakePlanInstancesRequest) (response *KillOpsMakePlanInstancesResponse, err error) {
    if request == nil {
        request = NewKillOpsMakePlanInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillOpsMakePlanInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillOpsMakePlanInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewKillScheduleInstancesRequest() (request *KillScheduleInstancesRequest) {
    request = &KillScheduleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "KillScheduleInstances")
    
    
    return
}

func NewKillScheduleInstancesResponse() (response *KillScheduleInstancesResponse) {
    response = &KillScheduleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillScheduleInstances
// 批量kill实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) KillScheduleInstances(request *KillScheduleInstancesRequest) (response *KillScheduleInstancesResponse, err error) {
    return c.KillScheduleInstancesWithContext(context.Background(), request)
}

// KillScheduleInstances
// 批量kill实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) KillScheduleInstancesWithContext(ctx context.Context, request *KillScheduleInstancesRequest) (response *KillScheduleInstancesResponse, err error) {
    if request == nil {
        request = NewKillScheduleInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillScheduleInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillScheduleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewLockIntegrationTaskRequest() (request *LockIntegrationTaskRequest) {
    request = &LockIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "LockIntegrationTask")
    
    
    return
}

func NewLockIntegrationTaskResponse() (response *LockIntegrationTaskResponse) {
    response = &LockIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// LockIntegrationTask
// 锁定集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LockIntegrationTask(request *LockIntegrationTaskRequest) (response *LockIntegrationTaskResponse, err error) {
    return c.LockIntegrationTaskWithContext(context.Background(), request)
}

// LockIntegrationTask
// 锁定集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LockIntegrationTaskWithContext(ctx context.Context, request *LockIntegrationTaskRequest) (response *LockIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewLockIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LockIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewLockIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewMakeUpOpsTasksRequest() (request *MakeUpOpsTasksRequest) {
    request = &MakeUpOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "MakeUpOpsTasks")
    
    
    return
}

func NewMakeUpOpsTasksResponse() (response *MakeUpOpsTasksResponse) {
    response = &MakeUpOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MakeUpOpsTasks
// 任务批量补录，调度状态任务才可以补录；
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpOpsTasks(request *MakeUpOpsTasksRequest) (response *MakeUpOpsTasksResponse, err error) {
    return c.MakeUpOpsTasksWithContext(context.Background(), request)
}

// MakeUpOpsTasks
// 任务批量补录，调度状态任务才可以补录；
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpOpsTasksWithContext(ctx context.Context, request *MakeUpOpsTasksRequest) (response *MakeUpOpsTasksResponse, err error) {
    if request == nil {
        request = NewMakeUpOpsTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MakeUpOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewMakeUpOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewMakeUpTasksByWorkflowRequest() (request *MakeUpTasksByWorkflowRequest) {
    request = &MakeUpTasksByWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "MakeUpTasksByWorkflow")
    
    
    return
}

func NewMakeUpTasksByWorkflowResponse() (response *MakeUpTasksByWorkflowResponse) {
    response = &MakeUpTasksByWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MakeUpTasksByWorkflow
// 工作流补数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpTasksByWorkflow(request *MakeUpTasksByWorkflowRequest) (response *MakeUpTasksByWorkflowResponse, err error) {
    return c.MakeUpTasksByWorkflowWithContext(context.Background(), request)
}

// MakeUpTasksByWorkflow
// 工作流补数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpTasksByWorkflowWithContext(ctx context.Context, request *MakeUpTasksByWorkflowRequest) (response *MakeUpTasksByWorkflowResponse, err error) {
    if request == nil {
        request = NewMakeUpTasksByWorkflowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MakeUpTasksByWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewMakeUpTasksByWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewMakeUpTasksNewRequest() (request *MakeUpTasksNewRequest) {
    request = &MakeUpTasksNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "MakeUpTasksNew")
    
    
    return
}

func NewMakeUpTasksNewResponse() (response *MakeUpTasksNewResponse) {
    response = &MakeUpTasksNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MakeUpTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 任务批量补录，调度状态任务才可以补录；
//
// 
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpTasksNew(request *MakeUpTasksNewRequest) (response *MakeUpTasksNewResponse, err error) {
    return c.MakeUpTasksNewWithContext(context.Background(), request)
}

// MakeUpTasksNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 任务批量补录，调度状态任务才可以补录；
//
// 
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpTasksNewWithContext(ctx context.Context, request *MakeUpTasksNewRequest) (response *MakeUpTasksNewResponse, err error) {
    if request == nil {
        request = NewMakeUpTasksNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MakeUpTasksNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewMakeUpTasksNewResponse()
    err = c.Send(request, response)
    return
}

func NewMakeUpWorkflowNewRequest() (request *MakeUpWorkflowNewRequest) {
    request = &MakeUpWorkflowNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "MakeUpWorkflowNew")
    
    
    return
}

func NewMakeUpWorkflowNewResponse() (response *MakeUpWorkflowNewResponse) {
    response = &MakeUpWorkflowNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MakeUpWorkflowNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 工作流下所有任务的补录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpWorkflowNew(request *MakeUpWorkflowNewRequest) (response *MakeUpWorkflowNewResponse, err error) {
    return c.MakeUpWorkflowNewWithContext(context.Background(), request)
}

// MakeUpWorkflowNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 工作流下所有任务的补录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) MakeUpWorkflowNewWithContext(ctx context.Context, request *MakeUpWorkflowNewRequest) (response *MakeUpWorkflowNewResponse, err error) {
    if request == nil {
        request = NewMakeUpWorkflowNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MakeUpWorkflowNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewMakeUpWorkflowNewResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApproveStatusRequest() (request *ModifyApproveStatusRequest) {
    request = &ModifyApproveStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyApproveStatus")
    
    
    return
}

func NewModifyApproveStatusResponse() (response *ModifyApproveStatusResponse) {
    response = &ModifyApproveStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApproveStatus
// 修改审批单状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyApproveStatus(request *ModifyApproveStatusRequest) (response *ModifyApproveStatusResponse, err error) {
    return c.ModifyApproveStatusWithContext(context.Background(), request)
}

// ModifyApproveStatus
// 修改审批单状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyApproveStatusWithContext(ctx context.Context, request *ModifyApproveStatusRequest) (response *ModifyApproveStatusResponse, err error) {
    if request == nil {
        request = NewModifyApproveStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApproveStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApproveStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBaselineAlarmStatusRequest() (request *ModifyBaselineAlarmStatusRequest) {
    request = &ModifyBaselineAlarmStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyBaselineAlarmStatus")
    
    
    return
}

func NewModifyBaselineAlarmStatusResponse() (response *ModifyBaselineAlarmStatusResponse) {
    response = &ModifyBaselineAlarmStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBaselineAlarmStatus
// 编辑基线告警状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyBaselineAlarmStatus(request *ModifyBaselineAlarmStatusRequest) (response *ModifyBaselineAlarmStatusResponse, err error) {
    return c.ModifyBaselineAlarmStatusWithContext(context.Background(), request)
}

// ModifyBaselineAlarmStatus
// 编辑基线告警状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyBaselineAlarmStatusWithContext(ctx context.Context, request *ModifyBaselineAlarmStatusRequest) (response *ModifyBaselineAlarmStatusResponse, err error) {
    if request == nil {
        request = NewModifyBaselineAlarmStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBaselineAlarmStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBaselineAlarmStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBaselineTaskAlarmStatusRequest() (request *ModifyBaselineTaskAlarmStatusRequest) {
    request = &ModifyBaselineTaskAlarmStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyBaselineTaskAlarmStatus")
    
    
    return
}

func NewModifyBaselineTaskAlarmStatusResponse() (response *ModifyBaselineTaskAlarmStatusResponse) {
    response = &ModifyBaselineTaskAlarmStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBaselineTaskAlarmStatus
// 编辑基线实例中任务告警状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyBaselineTaskAlarmStatus(request *ModifyBaselineTaskAlarmStatusRequest) (response *ModifyBaselineTaskAlarmStatusResponse, err error) {
    return c.ModifyBaselineTaskAlarmStatusWithContext(context.Background(), request)
}

// ModifyBaselineTaskAlarmStatus
// 编辑基线实例中任务告警状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyBaselineTaskAlarmStatusWithContext(ctx context.Context, request *ModifyBaselineTaskAlarmStatusRequest) (response *ModifyBaselineTaskAlarmStatusResponse, err error) {
    if request == nil {
        request = NewModifyBaselineTaskAlarmStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBaselineTaskAlarmStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBaselineTaskAlarmStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDataSourceRequest() (request *ModifyDataSourceRequest) {
    request = &ModifyDataSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyDataSource")
    
    
    return
}

func NewModifyDataSourceResponse() (response *ModifyDataSourceResponse) {
    response = &ModifyDataSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDataSource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 修改数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDataSource(request *ModifyDataSourceRequest) (response *ModifyDataSourceResponse, err error) {
    return c.ModifyDataSourceWithContext(context.Background(), request)
}

// ModifyDataSource
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 修改数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDataSourceWithContext(ctx context.Context, request *ModifyDataSourceRequest) (response *ModifyDataSourceResponse, err error) {
    if request == nil {
        request = NewModifyDataSourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataSourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDimensionWeightRequest() (request *ModifyDimensionWeightRequest) {
    request = &ModifyDimensionWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyDimensionWeight")
    
    
    return
}

func NewModifyDimensionWeightResponse() (response *ModifyDimensionWeightResponse) {
    response = &ModifyDimensionWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDimensionWeight
// 质量报告-修改维度权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDimensionWeight(request *ModifyDimensionWeightRequest) (response *ModifyDimensionWeightResponse, err error) {
    return c.ModifyDimensionWeightWithContext(context.Background(), request)
}

// ModifyDimensionWeight
// 质量报告-修改维度权限
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDimensionWeightWithContext(ctx context.Context, request *ModifyDimensionWeightRequest) (response *ModifyDimensionWeightResponse, err error) {
    if request == nil {
        request = NewModifyDimensionWeightRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDimensionWeight require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDimensionWeightResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDsFolderRequest() (request *ModifyDsFolderRequest) {
    request = &ModifyDsFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyDsFolder")
    
    
    return
}

func NewModifyDsFolderResponse() (response *ModifyDsFolderResponse) {
    response = &ModifyDsFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDsFolder
// 数据开发模块-文件夹更新
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDsFolder(request *ModifyDsFolderRequest) (response *ModifyDsFolderResponse, err error) {
    return c.ModifyDsFolderWithContext(context.Background(), request)
}

// ModifyDsFolder
// 数据开发模块-文件夹更新
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDsFolderWithContext(ctx context.Context, request *ModifyDsFolderRequest) (response *ModifyDsFolderResponse, err error) {
    if request == nil {
        request = NewModifyDsFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDsFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDsFolderResponse()
    err = c.Send(request, response)
    return
}

func NewModifyExecStrategyRequest() (request *ModifyExecStrategyRequest) {
    request = &ModifyExecStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyExecStrategy")
    
    
    return
}

func NewModifyExecStrategyResponse() (response *ModifyExecStrategyResponse) {
    response = &ModifyExecStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyExecStrategy
// 更新规则组执行策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  UNSUPPORTEDOPERATION_NORULEINRULEGROUP = "UnsupportedOperation.NoRuleInRuleGroup"
func (c *Client) ModifyExecStrategy(request *ModifyExecStrategyRequest) (response *ModifyExecStrategyResponse, err error) {
    return c.ModifyExecStrategyWithContext(context.Background(), request)
}

// ModifyExecStrategy
// 更新规则组执行策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  UNSUPPORTEDOPERATION_NORULEINRULEGROUP = "UnsupportedOperation.NoRuleInRuleGroup"
func (c *Client) ModifyExecStrategyWithContext(ctx context.Context, request *ModifyExecStrategyRequest) (response *ModifyExecStrategyResponse, err error) {
    if request == nil {
        request = NewModifyExecStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyExecStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyExecStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFolderRequest() (request *ModifyFolderRequest) {
    request = &ModifyFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyFolder")
    
    
    return
}

func NewModifyFolderResponse() (response *ModifyFolderResponse) {
    response = &ModifyFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 文件夹更新
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  UNSUPPORTEDOPERATION_NORULEINRULEGROUP = "UnsupportedOperation.NoRuleInRuleGroup"
func (c *Client) ModifyFolder(request *ModifyFolderRequest) (response *ModifyFolderResponse, err error) {
    return c.ModifyFolderWithContext(context.Background(), request)
}

// ModifyFolder
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 文件夹更新
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  UNSUPPORTEDOPERATION_NORULEINRULEGROUP = "UnsupportedOperation.NoRuleInRuleGroup"
func (c *Client) ModifyFolderWithContext(ctx context.Context, request *ModifyFolderRequest) (response *ModifyFolderResponse, err error) {
    if request == nil {
        request = NewModifyFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFolderResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIntegrationNodeRequest() (request *ModifyIntegrationNodeRequest) {
    request = &ModifyIntegrationNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyIntegrationNode")
    
    
    return
}

func NewModifyIntegrationNodeResponse() (response *ModifyIntegrationNodeResponse) {
    response = &ModifyIntegrationNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIntegrationNode
// 更新集成节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIntegrationNode(request *ModifyIntegrationNodeRequest) (response *ModifyIntegrationNodeResponse, err error) {
    return c.ModifyIntegrationNodeWithContext(context.Background(), request)
}

// ModifyIntegrationNode
// 更新集成节点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIntegrationNodeWithContext(ctx context.Context, request *ModifyIntegrationNodeRequest) (response *ModifyIntegrationNodeResponse, err error) {
    if request == nil {
        request = NewModifyIntegrationNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIntegrationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIntegrationNodeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIntegrationTaskRequest() (request *ModifyIntegrationTaskRequest) {
    request = &ModifyIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyIntegrationTask")
    
    
    return
}

func NewModifyIntegrationTaskResponse() (response *ModifyIntegrationTaskResponse) {
    response = &ModifyIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIntegrationTask
// 更新集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIntegrationTask(request *ModifyIntegrationTaskRequest) (response *ModifyIntegrationTaskResponse, err error) {
    return c.ModifyIntegrationTaskWithContext(context.Background(), request)
}

// ModifyIntegrationTask
// 更新集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIntegrationTaskWithContext(ctx context.Context, request *ModifyIntegrationTaskRequest) (response *ModifyIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewModifyIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMonitorStatusRequest() (request *ModifyMonitorStatusRequest) {
    request = &ModifyMonitorStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyMonitorStatus")
    
    
    return
}

func NewModifyMonitorStatusResponse() (response *ModifyMonitorStatusResponse) {
    response = &ModifyMonitorStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMonitorStatus
// 更新监控状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMonitorStatus(request *ModifyMonitorStatusRequest) (response *ModifyMonitorStatusResponse, err error) {
    return c.ModifyMonitorStatusWithContext(context.Background(), request)
}

// ModifyMonitorStatus
// 更新监控状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMonitorStatusWithContext(ctx context.Context, request *ModifyMonitorStatusRequest) (response *ModifyMonitorStatusResponse, err error) {
    if request == nil {
        request = NewModifyMonitorStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMonitorStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMonitorStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleRequest() (request *ModifyRuleRequest) {
    request = &ModifyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyRule")
    
    
    return
}

func NewModifyRuleResponse() (response *ModifyRuleResponse) {
    response = &ModifyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRule
// 更新质量规则接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    return c.ModifyRuleWithContext(context.Background(), request)
}

// ModifyRule
// 更新质量规则接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
func (c *Client) ModifyRuleWithContext(ctx context.Context, request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    if request == nil {
        request = NewModifyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleGroupSubscriptionRequest() (request *ModifyRuleGroupSubscriptionRequest) {
    request = &ModifyRuleGroupSubscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyRuleGroupSubscription")
    
    
    return
}

func NewModifyRuleGroupSubscriptionResponse() (response *ModifyRuleGroupSubscriptionResponse) {
    response = &ModifyRuleGroupSubscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRuleGroupSubscription
// 更新规则组订阅信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"
func (c *Client) ModifyRuleGroupSubscription(request *ModifyRuleGroupSubscriptionRequest) (response *ModifyRuleGroupSubscriptionResponse, err error) {
    return c.ModifyRuleGroupSubscriptionWithContext(context.Background(), request)
}

// ModifyRuleGroupSubscription
// 更新规则组订阅信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"
func (c *Client) ModifyRuleGroupSubscriptionWithContext(ctx context.Context, request *ModifyRuleGroupSubscriptionRequest) (response *ModifyRuleGroupSubscriptionResponse, err error) {
    if request == nil {
        request = NewModifyRuleGroupSubscriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRuleGroupSubscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRuleGroupSubscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleTemplateRequest() (request *ModifyRuleTemplateRequest) {
    request = &ModifyRuleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyRuleTemplate")
    
    
    return
}

func NewModifyRuleTemplateResponse() (response *ModifyRuleTemplateResponse) {
    response = &ModifyRuleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRuleTemplate
// 编辑规则模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RULETEMPLATENOTEXIST = "InvalidParameter.RuleTemplateNotExist"
func (c *Client) ModifyRuleTemplate(request *ModifyRuleTemplateRequest) (response *ModifyRuleTemplateResponse, err error) {
    return c.ModifyRuleTemplateWithContext(context.Background(), request)
}

// ModifyRuleTemplate
// 编辑规则模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RULETEMPLATENOTEXIST = "InvalidParameter.RuleTemplateNotExist"
func (c *Client) ModifyRuleTemplateWithContext(ctx context.Context, request *ModifyRuleTemplateRequest) (response *ModifyRuleTemplateResponse, err error) {
    if request == nil {
        request = NewModifyRuleTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRuleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRuleTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskAlarmRegularRequest() (request *ModifyTaskAlarmRegularRequest) {
    request = &ModifyTaskAlarmRegularRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyTaskAlarmRegular")
    
    
    return
}

func NewModifyTaskAlarmRegularResponse() (response *ModifyTaskAlarmRegularResponse) {
    response = &ModifyTaskAlarmRegularResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTaskAlarmRegular
// 修改任务告警规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTaskAlarmRegular(request *ModifyTaskAlarmRegularRequest) (response *ModifyTaskAlarmRegularResponse, err error) {
    return c.ModifyTaskAlarmRegularWithContext(context.Background(), request)
}

// ModifyTaskAlarmRegular
// 修改任务告警规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyTaskAlarmRegularWithContext(ctx context.Context, request *ModifyTaskAlarmRegularRequest) (response *ModifyTaskAlarmRegularResponse, err error) {
    if request == nil {
        request = NewModifyTaskAlarmRegularRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskAlarmRegular require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskAlarmRegularResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskInfoRequest() (request *ModifyTaskInfoRequest) {
    request = &ModifyTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyTaskInfo")
    
    
    return
}

func NewModifyTaskInfoResponse() (response *ModifyTaskInfoResponse) {
    response = &ModifyTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTaskInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskInfo(request *ModifyTaskInfoRequest) (response *ModifyTaskInfoResponse, err error) {
    return c.ModifyTaskInfoWithContext(context.Background(), request)
}

// ModifyTaskInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskInfoWithContext(ctx context.Context, request *ModifyTaskInfoRequest) (response *ModifyTaskInfoResponse, err error) {
    if request == nil {
        request = NewModifyTaskInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskLinksRequest() (request *ModifyTaskLinksRequest) {
    request = &ModifyTaskLinksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyTaskLinks")
    
    
    return
}

func NewModifyTaskLinksResponse() (response *ModifyTaskLinksResponse) {
    response = &ModifyTaskLinksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTaskLinks
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 添加父任务依赖
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskLinks(request *ModifyTaskLinksRequest) (response *ModifyTaskLinksResponse, err error) {
    return c.ModifyTaskLinksWithContext(context.Background(), request)
}

// ModifyTaskLinks
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 添加父任务依赖
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskLinksWithContext(ctx context.Context, request *ModifyTaskLinksRequest) (response *ModifyTaskLinksResponse, err error) {
    if request == nil {
        request = NewModifyTaskLinksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskLinks require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskLinksResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskNameRequest() (request *ModifyTaskNameRequest) {
    request = &ModifyTaskNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyTaskName")
    
    
    return
}

func NewModifyTaskNameResponse() (response *ModifyTaskNameResponse) {
    response = &ModifyTaskNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTaskName
// 重命名任务（任务编辑）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskName(request *ModifyTaskNameRequest) (response *ModifyTaskNameResponse, err error) {
    return c.ModifyTaskNameWithContext(context.Background(), request)
}

// ModifyTaskName
// 重命名任务（任务编辑）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskNameWithContext(ctx context.Context, request *ModifyTaskNameRequest) (response *ModifyTaskNameResponse, err error) {
    if request == nil {
        request = NewModifyTaskNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskScriptRequest() (request *ModifyTaskScriptRequest) {
    request = &ModifyTaskScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyTaskScript")
    
    
    return
}

func NewModifyTaskScriptResponse() (response *ModifyTaskScriptResponse) {
    response = &ModifyTaskScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTaskScript
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 修改任务脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskScript(request *ModifyTaskScriptRequest) (response *ModifyTaskScriptResponse, err error) {
    return c.ModifyTaskScriptWithContext(context.Background(), request)
}

// ModifyTaskScript
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 修改任务脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTaskScriptWithContext(ctx context.Context, request *ModifyTaskScriptRequest) (response *ModifyTaskScriptResponse, err error) {
    if request == nil {
        request = NewModifyTaskScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskScriptResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkflowInfoRequest() (request *ModifyWorkflowInfoRequest) {
    request = &ModifyWorkflowInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyWorkflowInfo")
    
    
    return
}

func NewModifyWorkflowInfoResponse() (response *ModifyWorkflowInfoResponse) {
    response = &ModifyWorkflowInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWorkflowInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyWorkflowInfo(request *ModifyWorkflowInfoRequest) (response *ModifyWorkflowInfoResponse, err error) {
    return c.ModifyWorkflowInfoWithContext(context.Background(), request)
}

// ModifyWorkflowInfo
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyWorkflowInfoWithContext(ctx context.Context, request *ModifyWorkflowInfoRequest) (response *ModifyWorkflowInfoResponse, err error) {
    if request == nil {
        request = NewModifyWorkflowInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkflowInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkflowInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkflowScheduleRequest() (request *ModifyWorkflowScheduleRequest) {
    request = &ModifyWorkflowScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyWorkflowSchedule")
    
    
    return
}

func NewModifyWorkflowScheduleResponse() (response *ModifyWorkflowScheduleResponse) {
    response = &ModifyWorkflowScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWorkflowSchedule
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新工作流调度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyWorkflowSchedule(request *ModifyWorkflowScheduleRequest) (response *ModifyWorkflowScheduleResponse, err error) {
    return c.ModifyWorkflowScheduleWithContext(context.Background(), request)
}

// ModifyWorkflowSchedule
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 更新工作流调度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyWorkflowScheduleWithContext(ctx context.Context, request *ModifyWorkflowScheduleRequest) (response *ModifyWorkflowScheduleResponse, err error) {
    if request == nil {
        request = NewModifyWorkflowScheduleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkflowSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkflowScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterEventRequest() (request *RegisterEventRequest) {
    request = &RegisterEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RegisterEvent")
    
    
    return
}

func NewRegisterEventResponse() (response *RegisterEventResponse) {
    response = &RegisterEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterEvent
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 注册事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RegisterEvent(request *RegisterEventRequest) (response *RegisterEventResponse, err error) {
    return c.RegisterEventWithContext(context.Background(), request)
}

// RegisterEvent
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 注册事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RegisterEventWithContext(ctx context.Context, request *RegisterEventRequest) (response *RegisterEventResponse, err error) {
    if request == nil {
        request = NewRegisterEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterEventResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterEventListenerRequest() (request *RegisterEventListenerRequest) {
    request = &RegisterEventListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RegisterEventListener")
    
    
    return
}

func NewRegisterEventListenerResponse() (response *RegisterEventListenerResponse) {
    response = &RegisterEventListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterEventListener
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 注册事件监听器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RegisterEventListener(request *RegisterEventListenerRequest) (response *RegisterEventListenerResponse, err error) {
    return c.RegisterEventListenerWithContext(context.Background(), request)
}

// RegisterEventListener
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 注册事件监听器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RegisterEventListenerWithContext(ctx context.Context, request *RegisterEventListenerRequest) (response *RegisterEventListenerResponse, err error) {
    if request == nil {
        request = NewRegisterEventListenerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterEventListener require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterEventListenerResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveWorkflowDsRequest() (request *RemoveWorkflowDsRequest) {
    request = &RemoveWorkflowDsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RemoveWorkflowDs")
    
    
    return
}

func NewRemoveWorkflowDsResponse() (response *RemoveWorkflowDsResponse) {
    response = &RemoveWorkflowDsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveWorkflowDs
// 删除编排空间工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RemoveWorkflowDs(request *RemoveWorkflowDsRequest) (response *RemoveWorkflowDsResponse, err error) {
    return c.RemoveWorkflowDsWithContext(context.Background(), request)
}

// RemoveWorkflowDs
// 删除编排空间工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RemoveWorkflowDsWithContext(ctx context.Context, request *RemoveWorkflowDsRequest) (response *RemoveWorkflowDsResponse, err error) {
    if request == nil {
        request = NewRemoveWorkflowDsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveWorkflowDs require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveWorkflowDsResponse()
    err = c.Send(request, response)
    return
}

func NewRerunInstancesRequest() (request *RerunInstancesRequest) {
    request = &RerunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RerunInstances")
    
    
    return
}

func NewRerunInstancesResponse() (response *RerunInstancesResponse) {
    response = &RerunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RerunInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量重跑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunInstances(request *RerunInstancesRequest) (response *RerunInstancesResponse, err error) {
    return c.RerunInstancesWithContext(context.Background(), request)
}

// RerunInstances
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 实例批量重跑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunInstancesWithContext(ctx context.Context, request *RerunInstancesRequest) (response *RerunInstancesResponse, err error) {
    if request == nil {
        request = NewRerunInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RerunInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRerunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRerunOpsMakePlanInstancesRequest() (request *RerunOpsMakePlanInstancesRequest) {
    request = &RerunOpsMakePlanInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RerunOpsMakePlanInstances")
    
    
    return
}

func NewRerunOpsMakePlanInstancesResponse() (response *RerunOpsMakePlanInstancesResponse) {
    response = &RerunOpsMakePlanInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RerunOpsMakePlanInstances
// 按补录计划批量重跑/选择补录计划→补录任务→补录实例，点击重跑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunOpsMakePlanInstances(request *RerunOpsMakePlanInstancesRequest) (response *RerunOpsMakePlanInstancesResponse, err error) {
    return c.RerunOpsMakePlanInstancesWithContext(context.Background(), request)
}

// RerunOpsMakePlanInstances
// 按补录计划批量重跑/选择补录计划→补录任务→补录实例，点击重跑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunOpsMakePlanInstancesWithContext(ctx context.Context, request *RerunOpsMakePlanInstancesRequest) (response *RerunOpsMakePlanInstancesResponse, err error) {
    if request == nil {
        request = NewRerunOpsMakePlanInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RerunOpsMakePlanInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRerunOpsMakePlanInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRerunScheduleInstancesRequest() (request *RerunScheduleInstancesRequest) {
    request = &RerunScheduleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RerunScheduleInstances")
    
    
    return
}

func NewRerunScheduleInstancesResponse() (response *RerunScheduleInstancesResponse) {
    response = &RerunScheduleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RerunScheduleInstances
// 实例批量重跑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunScheduleInstances(request *RerunScheduleInstancesRequest) (response *RerunScheduleInstancesResponse, err error) {
    return c.RerunScheduleInstancesWithContext(context.Background(), request)
}

// RerunScheduleInstances
// 实例批量重跑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunScheduleInstancesWithContext(ctx context.Context, request *RerunScheduleInstancesRequest) (response *RerunScheduleInstancesResponse, err error) {
    if request == nil {
        request = NewRerunScheduleInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RerunScheduleInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRerunScheduleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRestartInLongAgentRequest() (request *RestartInLongAgentRequest) {
    request = &RestartInLongAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RestartInLongAgent")
    
    
    return
}

func NewRestartInLongAgentResponse() (response *RestartInLongAgentResponse) {
    response = &RestartInLongAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartInLongAgent
// 重启采集器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RestartInLongAgent(request *RestartInLongAgentRequest) (response *RestartInLongAgentResponse, err error) {
    return c.RestartInLongAgentWithContext(context.Background(), request)
}

// RestartInLongAgent
// 重启采集器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RestartInLongAgentWithContext(ctx context.Context, request *RestartInLongAgentRequest) (response *RestartInLongAgentResponse, err error) {
    if request == nil {
        request = NewRestartInLongAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartInLongAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartInLongAgentResponse()
    err = c.Send(request, response)
    return
}

func NewResumeIntegrationTaskRequest() (request *ResumeIntegrationTaskRequest) {
    request = &ResumeIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ResumeIntegrationTask")
    
    
    return
}

func NewResumeIntegrationTaskResponse() (response *ResumeIntegrationTaskResponse) {
    response = &ResumeIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeIntegrationTask
// 继续集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeIntegrationTask(request *ResumeIntegrationTaskRequest) (response *ResumeIntegrationTaskResponse, err error) {
    return c.ResumeIntegrationTaskWithContext(context.Background(), request)
}

// ResumeIntegrationTask
// 继续集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeIntegrationTaskWithContext(ctx context.Context, request *ResumeIntegrationTaskRequest) (response *ResumeIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewResumeIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewRobAndLockIntegrationTaskRequest() (request *RobAndLockIntegrationTaskRequest) {
    request = &RobAndLockIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RobAndLockIntegrationTask")
    
    
    return
}

func NewRobAndLockIntegrationTaskResponse() (response *RobAndLockIntegrationTaskResponse) {
    response = &RobAndLockIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RobAndLockIntegrationTask
// 抢占锁定集成任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RobAndLockIntegrationTask(request *RobAndLockIntegrationTaskRequest) (response *RobAndLockIntegrationTaskResponse, err error) {
    return c.RobAndLockIntegrationTaskWithContext(context.Background(), request)
}

// RobAndLockIntegrationTask
// 抢占锁定集成任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RobAndLockIntegrationTaskWithContext(ctx context.Context, request *RobAndLockIntegrationTaskRequest) (response *RobAndLockIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewRobAndLockIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RobAndLockIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRobAndLockIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewRunForceSucScheduleInstancesRequest() (request *RunForceSucScheduleInstancesRequest) {
    request = &RunForceSucScheduleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RunForceSucScheduleInstances")
    
    
    return
}

func NewRunForceSucScheduleInstancesResponse() (response *RunForceSucScheduleInstancesResponse) {
    response = &RunForceSucScheduleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunForceSucScheduleInstances
// 实例强制成功
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunForceSucScheduleInstances(request *RunForceSucScheduleInstancesRequest) (response *RunForceSucScheduleInstancesResponse, err error) {
    return c.RunForceSucScheduleInstancesWithContext(context.Background(), request)
}

// RunForceSucScheduleInstances
// 实例强制成功
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunForceSucScheduleInstancesWithContext(ctx context.Context, request *RunForceSucScheduleInstancesRequest) (response *RunForceSucScheduleInstancesResponse, err error) {
    if request == nil {
        request = NewRunForceSucScheduleInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunForceSucScheduleInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunForceSucScheduleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRunRerunScheduleInstancesRequest() (request *RunRerunScheduleInstancesRequest) {
    request = &RunRerunScheduleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RunRerunScheduleInstances")
    
    
    return
}

func NewRunRerunScheduleInstancesResponse() (response *RunRerunScheduleInstancesResponse) {
    response = &RunRerunScheduleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunRerunScheduleInstances
// 实例批量重跑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunRerunScheduleInstances(request *RunRerunScheduleInstancesRequest) (response *RunRerunScheduleInstancesResponse, err error) {
    return c.RunRerunScheduleInstancesWithContext(context.Background(), request)
}

// RunRerunScheduleInstances
// 实例批量重跑
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunRerunScheduleInstancesWithContext(ctx context.Context, request *RunRerunScheduleInstancesRequest) (response *RunRerunScheduleInstancesResponse, err error) {
    if request == nil {
        request = NewRunRerunScheduleInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunRerunScheduleInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunRerunScheduleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRunTaskRequest() (request *RunTaskRequest) {
    request = &RunTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RunTask")
    
    
    return
}

func NewRunTaskResponse() (response *RunTaskResponse) {
    response = &RunTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 运行任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunTask(request *RunTaskRequest) (response *RunTaskResponse, err error) {
    return c.RunTaskWithContext(context.Background(), request)
}

// RunTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 运行任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunTaskWithContext(ctx context.Context, request *RunTaskRequest) (response *RunTaskResponse, err error) {
    if request == nil {
        request = NewRunTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunTaskResponse()
    err = c.Send(request, response)
    return
}

func NewRunTasksByMultiWorkflowRequest() (request *RunTasksByMultiWorkflowRequest) {
    request = &RunTasksByMultiWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RunTasksByMultiWorkflow")
    
    
    return
}

func NewRunTasksByMultiWorkflowResponse() (response *RunTasksByMultiWorkflowResponse) {
    response = &RunTasksByMultiWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunTasksByMultiWorkflow
// 批量启动工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunTasksByMultiWorkflow(request *RunTasksByMultiWorkflowRequest) (response *RunTasksByMultiWorkflowResponse, err error) {
    return c.RunTasksByMultiWorkflowWithContext(context.Background(), request)
}

// RunTasksByMultiWorkflow
// 批量启动工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunTasksByMultiWorkflowWithContext(ctx context.Context, request *RunTasksByMultiWorkflowRequest) (response *RunTasksByMultiWorkflowResponse, err error) {
    if request == nil {
        request = NewRunTasksByMultiWorkflowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunTasksByMultiWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunTasksByMultiWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewSaveCustomFunctionRequest() (request *SaveCustomFunctionRequest) {
    request = &SaveCustomFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SaveCustomFunction")
    
    
    return
}

func NewSaveCustomFunctionResponse() (response *SaveCustomFunctionResponse) {
    response = &SaveCustomFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SaveCustomFunction
// 保存用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SaveCustomFunction(request *SaveCustomFunctionRequest) (response *SaveCustomFunctionResponse, err error) {
    return c.SaveCustomFunctionWithContext(context.Background(), request)
}

// SaveCustomFunction
// 保存用户自定义函数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SaveCustomFunctionWithContext(ctx context.Context, request *SaveCustomFunctionRequest) (response *SaveCustomFunctionResponse, err error) {
    if request == nil {
        request = NewSaveCustomFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaveCustomFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaveCustomFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewSetTaskAlarmNewRequest() (request *SetTaskAlarmNewRequest) {
    request = &SetTaskAlarmNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SetTaskAlarmNew")
    
    
    return
}

func NewSetTaskAlarmNewResponse() (response *SetTaskAlarmNewResponse) {
    response = &SetTaskAlarmNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetTaskAlarmNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 设置任务告警，新建/更新告警信息（最新）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SetTaskAlarmNew(request *SetTaskAlarmNewRequest) (response *SetTaskAlarmNewResponse, err error) {
    return c.SetTaskAlarmNewWithContext(context.Background(), request)
}

// SetTaskAlarmNew
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 设置任务告警，新建/更新告警信息（最新）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SetTaskAlarmNewWithContext(ctx context.Context, request *SetTaskAlarmNewRequest) (response *SetTaskAlarmNewResponse, err error) {
    if request == nil {
        request = NewSetTaskAlarmNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTaskAlarmNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTaskAlarmNewResponse()
    err = c.Send(request, response)
    return
}

func NewStartIntegrationTaskRequest() (request *StartIntegrationTaskRequest) {
    request = &StartIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "StartIntegrationTask")
    
    
    return
}

func NewStartIntegrationTaskResponse() (response *StartIntegrationTaskResponse) {
    response = &StartIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartIntegrationTask
// 启动集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartIntegrationTask(request *StartIntegrationTaskRequest) (response *StartIntegrationTaskResponse, err error) {
    return c.StartIntegrationTaskWithContext(context.Background(), request)
}

// StartIntegrationTask
// 启动集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartIntegrationTaskWithContext(ctx context.Context, request *StartIntegrationTaskRequest) (response *StartIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewStartIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopBaselineRequest() (request *StopBaselineRequest) {
    request = &StopBaselineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "StopBaseline")
    
    
    return
}

func NewStopBaselineResponse() (response *StopBaselineResponse) {
    response = &StopBaselineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopBaseline
// 提交基线
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopBaseline(request *StopBaselineRequest) (response *StopBaselineResponse, err error) {
    return c.StopBaselineWithContext(context.Background(), request)
}

// StopBaseline
// 提交基线
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopBaselineWithContext(ctx context.Context, request *StopBaselineRequest) (response *StopBaselineResponse, err error) {
    if request == nil {
        request = NewStopBaselineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopBaseline require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopBaselineResponse()
    err = c.Send(request, response)
    return
}

func NewStopIntegrationTaskRequest() (request *StopIntegrationTaskRequest) {
    request = &StopIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "StopIntegrationTask")
    
    
    return
}

func NewStopIntegrationTaskResponse() (response *StopIntegrationTaskResponse) {
    response = &StopIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopIntegrationTask
// 停止集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopIntegrationTask(request *StopIntegrationTaskRequest) (response *StopIntegrationTaskResponse, err error) {
    return c.StopIntegrationTaskWithContext(context.Background(), request)
}

// StopIntegrationTask
// 停止集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopIntegrationTaskWithContext(ctx context.Context, request *StopIntegrationTaskRequest) (response *StopIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewStopIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitBaselineRequest() (request *SubmitBaselineRequest) {
    request = &SubmitBaselineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitBaseline")
    
    
    return
}

func NewSubmitBaselineResponse() (response *SubmitBaselineResponse) {
    response = &SubmitBaselineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitBaseline
// 提交基线
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitBaseline(request *SubmitBaselineRequest) (response *SubmitBaselineResponse, err error) {
    return c.SubmitBaselineWithContext(context.Background(), request)
}

// SubmitBaseline
// 提交基线
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitBaselineWithContext(ctx context.Context, request *SubmitBaselineRequest) (response *SubmitBaselineResponse, err error) {
    if request == nil {
        request = NewSubmitBaselineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitBaseline require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitBaselineResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitCustomFunctionRequest() (request *SubmitCustomFunctionRequest) {
    request = &SubmitCustomFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitCustomFunction")
    
    
    return
}

func NewSubmitCustomFunctionResponse() (response *SubmitCustomFunctionResponse) {
    response = &SubmitCustomFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitCustomFunction
// 提交自定义函数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitCustomFunction(request *SubmitCustomFunctionRequest) (response *SubmitCustomFunctionResponse, err error) {
    return c.SubmitCustomFunctionWithContext(context.Background(), request)
}

// SubmitCustomFunction
// 提交自定义函数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitCustomFunctionWithContext(ctx context.Context, request *SubmitCustomFunctionRequest) (response *SubmitCustomFunctionResponse, err error) {
    if request == nil {
        request = NewSubmitCustomFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitCustomFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitCustomFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitSqlTaskRequest() (request *SubmitSqlTaskRequest) {
    request = &SubmitSqlTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitSqlTask")
    
    
    return
}

func NewSubmitSqlTaskResponse() (response *SubmitSqlTaskResponse) {
    response = &SubmitSqlTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitSqlTask
// 即席分析提交SQL任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALCALLCLOUDAPIERROR = "InternalError.InternalCallCloudApiError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEINSTANCENOTEXISTS = "InvalidParameter.DataEngineInstanceNotExists"
//  INVALIDPARAMETER_WEDATAROLENOTEXISTS = "InvalidParameter.WeDataRoleNotExists"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitSqlTask(request *SubmitSqlTaskRequest) (response *SubmitSqlTaskResponse, err error) {
    return c.SubmitSqlTaskWithContext(context.Background(), request)
}

// SubmitSqlTask
// 即席分析提交SQL任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALCALLCLOUDAPIERROR = "InternalError.InternalCallCloudApiError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEINSTANCENOTEXISTS = "InvalidParameter.DataEngineInstanceNotExists"
//  INVALIDPARAMETER_WEDATAROLENOTEXISTS = "InvalidParameter.WeDataRoleNotExists"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitSqlTaskWithContext(ctx context.Context, request *SubmitSqlTaskRequest) (response *SubmitSqlTaskResponse, err error) {
    if request == nil {
        request = NewSubmitSqlTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitSqlTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitSqlTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTaskRequest() (request *SubmitTaskRequest) {
    request = &SubmitTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitTask")
    
    
    return
}

func NewSubmitTaskResponse() (response *SubmitTaskResponse) {
    response = &SubmitTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 提交任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALCALLCLOUDAPIERROR = "InternalError.InternalCallCloudApiError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEINSTANCENOTEXISTS = "InvalidParameter.DataEngineInstanceNotExists"
//  INVALIDPARAMETER_WEDATAROLENOTEXISTS = "InvalidParameter.WeDataRoleNotExists"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitTask(request *SubmitTaskRequest) (response *SubmitTaskResponse, err error) {
    return c.SubmitTaskWithContext(context.Background(), request)
}

// SubmitTask
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 提交任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALCALLCLOUDAPIERROR = "InternalError.InternalCallCloudApiError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEINSTANCENOTEXISTS = "InvalidParameter.DataEngineInstanceNotExists"
//  INVALIDPARAMETER_WEDATAROLENOTEXISTS = "InvalidParameter.WeDataRoleNotExists"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitTaskWithContext(ctx context.Context, request *SubmitTaskRequest) (response *SubmitTaskResponse, err error) {
    if request == nil {
        request = NewSubmitTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTaskTestRunRequest() (request *SubmitTaskTestRunRequest) {
    request = &SubmitTaskTestRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitTaskTestRun")
    
    
    return
}

func NewSubmitTaskTestRunResponse() (response *SubmitTaskTestRunResponse) {
    response = &SubmitTaskTestRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTaskTestRun
// 无
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitTaskTestRun(request *SubmitTaskTestRunRequest) (response *SubmitTaskTestRunResponse, err error) {
    return c.SubmitTaskTestRunWithContext(context.Background(), request)
}

// SubmitTaskTestRun
// 无
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitTaskTestRunWithContext(ctx context.Context, request *SubmitTaskTestRunRequest) (response *SubmitTaskTestRunResponse, err error) {
    if request == nil {
        request = NewSubmitTaskTestRunRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTaskTestRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTaskTestRunResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitWorkflowRequest() (request *SubmitWorkflowRequest) {
    request = &SubmitWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitWorkflow")
    
    
    return
}

func NewSubmitWorkflowResponse() (response *SubmitWorkflowResponse) {
    response = &SubmitWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 提交工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitWorkflow(request *SubmitWorkflowRequest) (response *SubmitWorkflowResponse, err error) {
    return c.SubmitWorkflowWithContext(context.Background(), request)
}

// SubmitWorkflow
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 提交工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitWorkflowWithContext(ctx context.Context, request *SubmitWorkflowRequest) (response *SubmitWorkflowResponse, err error) {
    if request == nil {
        request = NewSubmitWorkflowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewSuspendIntegrationTaskRequest() (request *SuspendIntegrationTaskRequest) {
    request = &SuspendIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SuspendIntegrationTask")
    
    
    return
}

func NewSuspendIntegrationTaskResponse() (response *SuspendIntegrationTaskResponse) {
    response = &SuspendIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SuspendIntegrationTask
// 暂停集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SuspendIntegrationTask(request *SuspendIntegrationTaskRequest) (response *SuspendIntegrationTaskResponse, err error) {
    return c.SuspendIntegrationTaskWithContext(context.Background(), request)
}

// SuspendIntegrationTask
// 暂停集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SuspendIntegrationTaskWithContext(ctx context.Context, request *SuspendIntegrationTaskRequest) (response *SuspendIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewSuspendIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SuspendIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSuspendIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewTaskLogRequest() (request *TaskLogRequest) {
    request = &TaskLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "TaskLog")
    
    
    return
}

func NewTaskLogResponse() (response *TaskLogResponse) {
    response = &TaskLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TaskLog
// 查询Inlong manager日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TaskLog(request *TaskLogRequest) (response *TaskLogResponse, err error) {
    return c.TaskLogWithContext(context.Background(), request)
}

// TaskLog
// 查询Inlong manager日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TaskLogWithContext(ctx context.Context, request *TaskLogRequest) (response *TaskLogResponse, err error) {
    if request == nil {
        request = NewTaskLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TaskLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewTaskLogResponse()
    err = c.Send(request, response)
    return
}

func NewTriggerEventRequest() (request *TriggerEventRequest) {
    request = &TriggerEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "TriggerEvent")
    
    
    return
}

func NewTriggerEventResponse() (response *TriggerEventResponse) {
    response = &TriggerEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TriggerEvent
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 触发事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) TriggerEvent(request *TriggerEventRequest) (response *TriggerEventResponse, err error) {
    return c.TriggerEventWithContext(context.Background(), request)
}

// TriggerEvent
// <p style="color:red;">[注意：该Beta版本只满足广州区部分白名单客户使用]</p>
//
// 触发事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) TriggerEventWithContext(ctx context.Context, request *TriggerEventRequest) (response *TriggerEventResponse, err error) {
    if request == nil {
        request = NewTriggerEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TriggerEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewTriggerEventResponse()
    err = c.Send(request, response)
    return
}

func NewUnlockIntegrationTaskRequest() (request *UnlockIntegrationTaskRequest) {
    request = &UnlockIntegrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UnlockIntegrationTask")
    
    
    return
}

func NewUnlockIntegrationTaskResponse() (response *UnlockIntegrationTaskResponse) {
    response = &UnlockIntegrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnlockIntegrationTask
// 解锁集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnlockIntegrationTask(request *UnlockIntegrationTaskRequest) (response *UnlockIntegrationTaskResponse, err error) {
    return c.UnlockIntegrationTaskWithContext(context.Background(), request)
}

// UnlockIntegrationTask
// 解锁集成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnlockIntegrationTaskWithContext(ctx context.Context, request *UnlockIntegrationTaskRequest) (response *UnlockIntegrationTaskResponse, err error) {
    if request == nil {
        request = NewUnlockIntegrationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnlockIntegrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnlockIntegrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateInLongAgentRequest() (request *UpdateInLongAgentRequest) {
    request = &UpdateInLongAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateInLongAgent")
    
    
    return
}

func NewUpdateInLongAgentResponse() (response *UpdateInLongAgentResponse) {
    response = &UpdateInLongAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateInLongAgent
// 更新采集器
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateInLongAgent(request *UpdateInLongAgentRequest) (response *UpdateInLongAgentResponse, err error) {
    return c.UpdateInLongAgentWithContext(context.Background(), request)
}

// UpdateInLongAgent
// 更新采集器
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateInLongAgentWithContext(ctx context.Context, request *UpdateInLongAgentRequest) (response *UpdateInLongAgentResponse, err error) {
    if request == nil {
        request = NewUpdateInLongAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateInLongAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateInLongAgentResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkflowOwnerRequest() (request *UpdateWorkflowOwnerRequest) {
    request = &UpdateWorkflowOwnerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateWorkflowOwner")
    
    
    return
}

func NewUpdateWorkflowOwnerResponse() (response *UpdateWorkflowOwnerResponse) {
    response = &UpdateWorkflowOwnerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkflowOwner
// 修改工作流责任人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) UpdateWorkflowOwner(request *UpdateWorkflowOwnerRequest) (response *UpdateWorkflowOwnerResponse, err error) {
    return c.UpdateWorkflowOwnerWithContext(context.Background(), request)
}

// UpdateWorkflowOwner
// 修改工作流责任人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) UpdateWorkflowOwnerWithContext(ctx context.Context, request *UpdateWorkflowOwnerRequest) (response *UpdateWorkflowOwnerResponse, err error) {
    if request == nil {
        request = NewUpdateWorkflowOwnerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkflowOwner require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkflowOwnerResponse()
    err = c.Send(request, response)
    return
}

func NewUploadContentRequest() (request *UploadContentRequest) {
    request = &UploadContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UploadContent")
    
    
    return
}

func NewUploadContentResponse() (response *UploadContentResponse) {
    response = &UploadContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadContent
// 保存任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEINSTANCENOTEXISTS = "InvalidParameter.DataEngineInstanceNotExists"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UploadContent(request *UploadContentRequest) (response *UploadContentResponse, err error) {
    return c.UploadContentWithContext(context.Background(), request)
}

// UploadContent
// 保存任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAENGINEINSTANCENOTEXISTS = "InvalidParameter.DataEngineInstanceNotExists"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UploadContentWithContext(ctx context.Context, request *UploadContentRequest) (response *UploadContentResponse, err error) {
    if request == nil {
        request = NewUploadContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadContentResponse()
    err = c.Send(request, response)
    return
}

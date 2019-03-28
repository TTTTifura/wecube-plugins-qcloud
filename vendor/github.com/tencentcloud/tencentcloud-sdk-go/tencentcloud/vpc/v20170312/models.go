// Copyright 1999-2018 Tencent Ltd.
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

package v20170312

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Address struct {
	// `EIP`的`ID`，是`EIP`的唯一标识。
	AddressId *string `json:"AddressId" name:"AddressId"`
	// `EIP`名称。
	AddressName *string `json:"AddressName" name:"AddressName"`
	// `EIP`状态。
	AddressStatus *string `json:"AddressStatus" name:"AddressStatus"`
	// 弹性外网IP
	AddressIp *string `json:"AddressIp" name:"AddressIp"`
	// 绑定的资源实例`ID`。可能是一个`CVM`，`NAT`，或是弹性网卡。
	BindedResourceId *string `json:"BindedResourceId" name:"BindedResourceId"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type AddressTemplate struct {
	// IP地址模板名称。
	AddressTemplateName *string `json:"AddressTemplateName" name:"AddressTemplateName"`
	// IP地址模板实例唯一ID。
	AddressTemplateId *string `json:"AddressTemplateId" name:"AddressTemplateId"`
	// IP地址信息。
	AddressSet []*string `json:"AddressSet" name:"AddressSet" list`
	// 创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type AddressTemplateGroup struct {
	// IP地址模板集合名称。
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName" name:"AddressTemplateGroupName"`
	// IP地址模板集合实例ID，例如：ipmg-dih8xdbq。
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId" name:"AddressTemplateGroupId"`
	// IP地址模板ID。
	AddressTemplateIdSet []*string `json:"AddressTemplateIdSet" name:"AddressTemplateIdSet" list`
	// 创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type AllocateAddressesRequest struct {
	*tchttp.BaseRequest
	// 申请 EIP 数量，默认值为1。
	AddressCount *int64 `json:"AddressCount" name:"AddressCount"`
}

func (r *AllocateAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 申请到的 EIP 的唯一 ID 列表。
		AddressSet []*string `json:"AddressSet" name:"AddressSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	// 指定的内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses" name:"PrivateIpAddresses" list`
	// 新申请的内网IP地址个数。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount" name:"SecondaryPrivateIpAddressCount"`
}

func (r *AssignPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignPrivateIpAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 内网IP详细信息。
		PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet" name:"PrivateIpAddressSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignPrivateIpAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressRequest struct {
	*tchttp.BaseRequest
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId" name:"AddressId"`
	// 要绑定的实例 ID。实例 ID 形如：`ins-11112222`。可通过登录[控制台](https://console.cloud.tencent.com/cvm)查询，也可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/9389) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 要绑定的弹性网卡 ID。 弹性网卡 ID 形如：`eni-11112222`。`NetworkInterfaceId` 与 `InstanceId` 不可同时指定。弹性网卡 ID 可通过登录[控制台](https://console.cloud.tencent.com/vpc/eni)查询，也可通过[DescribeNetworkInterfaces](https://cloud.tencent.com/document/api/215/4814)接口返回值中的`networkInterfaceId`获取。
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	// 要绑定的内网 IP。如果指定了 `NetworkInterfaceId` 则也必须指定 `PrivateIpAddress` ，表示将 EIP 绑定到指定弹性网卡的指定内网 IP 上。同时要确保指定的 `PrivateIpAddress` 是指定的 `NetworkInterfaceId` 上的一个内网 IP。指定弹性网卡的内网 IP 可通过登录[控制台](https://console.cloud.tencent.com/vpc/eni)查询，也可通过[DescribeNetworkInterfaces](https://cloud.tencent.com/document/api/215/4814)接口返回值中的`privateIpAddress`获取。
	PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
}

func (r *AssociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachClassicLinkVpcRequest struct {
	*tchttp.BaseRequest
	// VPC实例ID
	VpcId *string `json:"VpcId" name:"VpcId"`
	// CVM实例ID
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
}

func (r *AttachClassicLinkVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachClassicLinkVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachClassicLinkVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachClassicLinkVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachClassicLinkVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	// CVM实例ID。形如：ins-r8hr2upy。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *AttachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClassicLinkInstance struct {
	// VPC实例ID
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 云服务器实例唯一ID
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

type CreateAddressTemplateGroupRequest struct {
	*tchttp.BaseRequest
	// IP地址模版集合名称。
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName" name:"AddressTemplateGroupName"`
	// IP地址模版实例ID，例如：ipm-mdunqeb6。
	AddressTemplateIds []*string `json:"AddressTemplateIds" name:"AddressTemplateIds" list`
}

func (r *CreateAddressTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAddressTemplateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAddressTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// IP地址模板集合对象。
		AddressTemplateGroup *AddressTemplateGroup `json:"AddressTemplateGroup" name:"AddressTemplateGroup"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAddressTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAddressTemplateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAddressTemplateRequest struct {
	*tchttp.BaseRequest
	// IP地址模版名称
	AddressTemplateName *string `json:"AddressTemplateName" name:"AddressTemplateName"`
	// 地址信息，支持 IP、CIDR、IP 范围。
	Addresses []*string `json:"Addresses" name:"Addresses" list`
}

func (r *CreateAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAddressTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// IP地址模板对象。
		AddressTemplate *AddressTemplate `json:"AddressTemplate" name:"AddressTemplate"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAddressTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerGatewayRequest struct {
	*tchttp.BaseRequest
	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName *string `json:"CustomerGatewayName" name:"CustomerGatewayName"`
	// 对端网关公网IP。
	IpAddress *string `json:"IpAddress" name:"IpAddress"`
}

func (r *CreateCustomerGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCustomerGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 对端网关对象
		CustomerGateway *CustomerGateway `json:"CustomerGateway" name:"CustomerGateway"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCustomerGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName" name:"NetworkInterfaceName"`
	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription" name:"NetworkInterfaceDescription"`
	// 弹性网卡所在的子网实例ID，例如：subnet-0ap8nwca。
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 新申请的内网IP地址个数。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount" name:"SecondaryPrivateIpAddressCount"`
	// 指定绑定的安全组，例如：['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds" name:"SecurityGroupIds" list`
	// 指定内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses" name:"PrivateIpAddresses" list`
}

func (r *CreateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 弹性网卡实例。
		NetworkInterface *NetworkInterface `json:"NetworkInterface" name:"NetworkInterface"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRouteTableRequest struct {
	*tchttp.BaseRequest
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName *string `json:"RouteTableName" name:"RouteTableName"`
}

func (r *CreateRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRouteTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 路由表对象。
		RouteTable *RouteTable `json:"RouteTable" name:"RouteTable"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRouteTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoutesRequest struct {
	*tchttp.BaseRequest
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId" name:"RouteTableId"`
	// 路由策略对象。
	Routes []*string `json:"Routes" name:"Routes" list`
}

func (r *CreateRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
	// 安全组规则集合。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet" name:"SecurityGroupPolicySet"`
}

func (r *CreateSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRequest struct {
	*tchttp.BaseRequest
	// 项目id，默认0。可在qcloud控制台项目管理页面查询到。
	ProjectId *string `json:"ProjectId" name:"ProjectId"`
	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName" name:"GroupName"`
	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription" name:"GroupDescription"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 安全组对象。
		SecurityGroup *SecurityGroup `json:"SecurityGroup" name:"SecurityGroup"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceTemplateGroupRequest struct {
	*tchttp.BaseRequest
	// 协议端口模板集合名称
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName" name:"ServiceTemplateGroupName"`
	// 协议端口模板实例ID，例如：ppm-4dw6agho。
	ServiceTemplateIds []*string `json:"ServiceTemplateIds" name:"ServiceTemplateIds" list`
}

func (r *CreateServiceTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceTemplateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 协议端口模板集合对象。
		ServiceTemplateGroup *ServiceTemplateGroup `json:"ServiceTemplateGroup" name:"ServiceTemplateGroup"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceTemplateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceTemplateRequest struct {
	*tchttp.BaseRequest
	// 协议端口模板名称
	ServiceTemplateName *string `json:"ServiceTemplateName" name:"ServiceTemplateName"`
	// 支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。
	Services []*string `json:"Services" name:"Services" list`
}

func (r *CreateServiceTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 协议端口模板对象。
		ServiceTemplate *ServiceTemplate `json:"ServiceTemplate" name:"ServiceTemplate"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetRequest struct {
	*tchttp.BaseRequest
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `json:"SubnetName" name:"SubnetName"`
	// 子网网段，子网网段必须在VPC网段内，相同VPC内子网网段不能重叠。
	CidrBlock *string `json:"CidrBlock" name:"CidrBlock"`
	// 子网所在的可用区ID，不同子网选择不同可用区可以做跨可用区灾备。
	Zone *string `json:"Zone" name:"Zone"`
}

func (r *CreateSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubnetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 子网对象。
		Subnet *Subnet `json:"Subnet" name:"Subnet"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubnetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcRequest struct {
	*tchttp.BaseRequest
	// vpc名称，最大长度不能超过60个字节。
	VpcName *string `json:"VpcName" name:"VpcName"`
	// vpc的cidr，只能为10.0.0.0/16，172.16.0.0/12，192.168.0.0/16这三个内网网段内。
	CidrBlock *string `json:"CidrBlock" name:"CidrBlock"`
	// 是否开启组播。true: 开启, false: 不开启。
	EnableMulticast *string `json:"EnableMulticast" name:"EnableMulticast"`
	// DNS地址，最多支持4个
	DnsServers []*string `json:"DnsServers" name:"DnsServers" list`
	// 域名
	DomainName *string `json:"DomainName" name:"DomainName"`
}

func (r *CreateVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// Vpc对象。
		Vpc *Vpc `json:"Vpc" name:"Vpc"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpnConnectionRequest struct {
	*tchttp.BaseRequest
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// 对端网关ID，例如：cgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId" name:"CustomerGatewayId"`
	// 通道名称，可任意命名，但不得超过60个字符。
	VpnConnectionName *string `json:"VpnConnectionName" name:"VpnConnectionName"`
	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey" name:"PreShareKey"`
	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases" name:"SecurityPolicyDatabases" list`
	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自保护机制，用户配置网络安全协议
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification" name:"IKEOptionsSpecification"`
	// IPSec配置，腾讯云提供IPSec安全会话设置
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification" name:"IPSECOptionsSpecification"`
}

func (r *CreateVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpnConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 通道实例对象。
		VpnConnection *VpnConnection `json:"VpnConnection" name:"VpnConnection"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpnConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpnGatewayRequest struct {
	*tchttp.BaseRequest
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// VPN网关名称，最大长度不能超过60个字节。
	VpnGatewayName *string `json:"VpnGatewayName" name:"VpnGatewayName"`
	// VPN网关计费模式，PREPAID：表示预付费，即包年包月，POSTPAID_BY_HOUR：表示后付费，即按量计费。默认：POSTPAID_BY_HOUR，如果指定预付费模式，参数InstanceChargePrepaid必填。
	InstanceChargeType *string `json:"InstanceChargeType" name:"InstanceChargeType"`
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut" name:"InternetMaxBandwidthOut"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid" name:"InstanceChargePrepaid"`
}

func (r *CreateVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// VPN网关对象
		VpnGateway *VpnGateway `json:"VpnGateway" name:"VpnGateway"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomerGateway struct {
	// 用户网关唯一ID
	CustomerGatewayId *string `json:"CustomerGatewayId" name:"CustomerGatewayId"`
	// 网关名称
	CustomerGatewayName *string `json:"CustomerGatewayName" name:"CustomerGatewayName"`
	// 公网地址
	IpAddress *string `json:"IpAddress" name:"IpAddress"`
	// 创建时间
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type CustomerGatewayVendor struct {
	// 平台。
	Platform *string `json:"Platform" name:"Platform"`
	// 软件版本。
	SoftwareVersion *string `json:"SoftwareVersion" name:"SoftwareVersion"`
	// 供应商名称。
	VendorName *string `json:"VendorName" name:"VendorName"`
}

type DeleteAddressTemplateGroupRequest struct {
	*tchttp.BaseRequest
	// IP地址模板集合实例ID，例如：ipmg-90cex8mq。
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId" name:"AddressTemplateGroupId"`
}

func (r *DeleteAddressTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAddressTemplateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAddressTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAddressTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAddressTemplateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAddressTemplateRequest struct {
	*tchttp.BaseRequest
	// IP地址模板实例ID，例如：ipm-09o5m8kc。
	AddressTemplateId *string `json:"AddressTemplateId" name:"AddressTemplateId"`
}

func (r *DeleteAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAddressTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAddressTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomerGatewayRequest struct {
	*tchttp.BaseRequest
	// 对端网关ID，例如：cgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId" name:"CustomerGatewayId"`
}

func (r *DeleteCustomerGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCustomerGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCustomerGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
}

func (r *DeleteNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteTableRequest struct {
	*tchttp.BaseRequest
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId" name:"RouteTableId"`
}

func (r *DeleteRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRouteTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRouteTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoutesRequest struct {
	*tchttp.BaseRequest
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId" name:"RouteTableId"`
	// 路由策略对象。
	Routes []*string `json:"Routes" name:"Routes" list`
}

func (r *DeleteRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
	// 安全组规则集合。一个请求中只能删除单个方向的一条或多条规则。支持指定索引（PolicyIndex） 匹配删除和安全组规则匹配删除两种方式，一个请求中只能使用一种匹配方式。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet" name:"SecurityGroupPolicySet"`
}

func (r *DeleteSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRequest struct {
	*tchttp.BaseRequest
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateGroupRequest struct {
	*tchttp.BaseRequest
	// 协议端口模板集合实例ID，例如：ppmg-n17uxvve。
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId" name:"ServiceTemplateGroupId"`
}

func (r *DeleteServiceTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceTemplateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceTemplateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateRequest struct {
	*tchttp.BaseRequest
	// 协议端口模板实例ID，例如：ppm-e6dy460g。
	ServiceTemplateId *string `json:"ServiceTemplateId" name:"ServiceTemplateId"`
}

func (r *DeleteServiceTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetRequest struct {
	*tchttp.BaseRequest
	// 子网实例ID。可通过DescribeSubnets接口返回值中的SubnetId获取。
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
}

func (r *DeleteSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubnetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubnetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcRequest struct {
	*tchttp.BaseRequest
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId" name:"VpcId"`
}

func (r *DeleteVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnConnectionRequest struct {
	*tchttp.BaseRequest
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId" name:"VpnConnectionId"`
}

func (r *DeleteVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayRequest struct {
	*tchttp.BaseRequest
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
}

func (r *DeleteVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAddressQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 账户 EIP 配额信息。
		QuotaSet []*Quota `json:"QuotaSet" name:"QuotaSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplateGroupsRequest struct {
	*tchttp.BaseRequest
	// 过滤条件。
	// <li>address-template-group-name - String - （过滤条件）IP地址模板集合名称。</li>
	// <li>address-template-group-id - String - （过滤条件）IP地址模板实集合例ID，例如：ipmg-mdunqeb6。</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。
	Offset *string `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit" name:"Limit"`
}

func (r *DescribeAddressTemplateGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressTemplateGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplateGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// IP地址模板。
		AddressTemplateGroupSet []*AddressTemplateGroup `json:"AddressTemplateGroupSet" name:"AddressTemplateGroupSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressTemplateGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressTemplateGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplatesRequest struct {
	*tchttp.BaseRequest
	// 过滤条件。
	// <li>address-template-name - String - （过滤条件）IP地址模板名称。</li>
	// <li>address-template-id - String - （过滤条件）IP地址模板实例ID，例如：ipm-mdunqeb6。</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。
	Offset *string `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit" name:"Limit"`
}

func (r *DescribeAddressTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// IP地址模版。
		AddressTemplateSet []*AddressTemplate `json:"AddressTemplateSet" name:"AddressTemplateSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesRequest struct {
	*tchttp.BaseRequest
	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：`eip-11112222`。参数不支持同时指定`AddressIds`和`Filters`。
	AddressIds []*string `json:"AddressIds" name:"AddressIds" list`
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AddressIds`和`Filters`。详细的过滤条件如下：
	// <li> address-id - String - 是否必填：否 - （过滤条件）按照 EIP 的唯一 ID 过滤。EIP 唯一 ID 形如：eip-11112222。</li>
	// <li> address-name - String - 是否必填：否 - （过滤条件）按照 EIP 名称过滤。不支持模糊过滤。</li>
	// <li> address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。</li>
	// <li> address-status - String - 是否必填：否 - （过滤条件）按照 EIP 的状态过滤。取值范围：[详见EIP状态列表](https://cloud.tencent.com/document/api/213/9452#eip_state)。</li>
	// <li> instance-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的实例 ID 过滤。实例 ID 形如：ins-11112222。</li>
	// <li> private-ip-address - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的内网 IP 过滤。</li>
	// <li> network-interface-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的弹性网卡 ID 过滤。弹性网卡 ID 形如：eni-11112222。</li>
	// <li> is-arrears - String - 是否必填：否 - （过滤条件）按照 EIP 是否欠费进行过滤。（TRUE：EIP 处于欠费状态|FALSE：EIP 费用状态正常）</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *int64 `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Limit *int64 `json:"Limit" name:"Limit"`
}

func (r *DescribeAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的 EIP 数量。
		TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
		// EIP 详细信息列表。
		AddressSet []*Address `json:"AddressSet" name:"AddressSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicLinkInstancesRequest struct {
	*tchttp.BaseRequest
	// 过滤条件。
	// <li>vpc-id - String - （过滤条件）VPC实例ID。</li>
	// <li>vm-ip - String - （过滤条件）基础网络云主机IP。</li>
	Filters []*FilterObject `json:"Filters" name:"Filters" list`
	// 偏移量
	Offset *string `json:"Offset" name:"Offset"`
	// 返回数量
	Limit *string `json:"Limit" name:"Limit"`
}

func (r *DescribeClassicLinkInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicLinkInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClassicLinkInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 私有网络和基础网络互通设备。
		ClassicLinkInstanceSet []*ClassicLinkInstance `json:"ClassicLinkInstanceSet" name:"ClassicLinkInstanceSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClassicLinkInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClassicLinkInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewayVendorsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCustomerGatewayVendorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomerGatewayVendorsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewayVendorsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 对端网关厂商信息对象。
		CustomerGatewayVendorSet []*CustomerGatewayVendor `json:"CustomerGatewayVendorSet" name:"CustomerGatewayVendorSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerGatewayVendorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomerGatewayVendorsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewaysRequest struct {
	*tchttp.BaseRequest
	// 对端网关ID，例如：cgw-2wqq41m9。每次请求的实例的上限为100。参数不支持同时指定CustomerGatewayIds和Filters。
	CustomerGatewayIds []*string `json:"CustomerGatewayIds" name:"CustomerGatewayIds" list`
	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定CustomerGatewayIds和Filters。
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeCustomerGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomerGatewaysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 对端网关对象列表
		CustomerGatewaySet []*CustomerGateway `json:"CustomerGatewaySet" name:"CustomerGatewaySet" list`
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomerGatewaysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesRequest struct {
	*tchttp.BaseRequest
	// 弹性网卡实例ID查询。形如：eni-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定NetworkInterfaceIds和Filters。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds" name:"NetworkInterfaceIds" list`
	// 过滤条件，参数不支持同时指定NetworkInterfaceIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>subnet-id - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。</li>
	// <li>network-interface-id - String - （过滤条件）弹性网卡实例ID，形如：eni-5k56k7k7。</li>
	// <li>attachment.instance-id - String - （过滤条件）绑定的云服务器实例ID，形如：ins-3nqpdn3i。</li>
	// <li>groups.security-group-id - String - （过滤条件）绑定的安全组实例ID，例如：sg-f9ekbxeq。</li>
	// <li>network-interface-name - String - （过滤条件）网卡实例名称。</li>
	// <li>network-interface-description - String - （过滤条件）网卡实例描述。</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeNetworkInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 实例详细信息列表。
		NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet" name:"NetworkInterfaceSet" list`
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTablesRequest struct {
	*tchttp.BaseRequest
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableIds []*string `json:"RouteTableIds" name:"RouteTableIds" list`
	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// <li>route-table-id - String - （过滤条件）路由表实例ID。</li>
	// <li>route-table-name - String - （过滤条件）路由表名称。</li>
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>association.main - String - （过滤条件）是否主路由表。</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量。
	Offset *string `json:"Offset" name:"Offset"`
	// 请求对象个数。
	Limit *string `json:"Limit" name:"Limit"`
}

func (r *DescribeRouteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 路由表对象。
		RouteTableSet []*RouteTable `json:"RouteTableSet" name:"RouteTableSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupAssociationStatisticsRequest struct {
	*tchttp.BaseRequest
	// 安全实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds" name:"SecurityGroupIds" list`
}

func (r *DescribeSecurityGroupAssociationStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupAssociationStatisticsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupAssociationStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 安全组关联实例统计。
		SecurityGroupAssociationStatisticsSet []*SecurityGroupAssociationStatistics `json:"SecurityGroupAssociationStatisticsSet" name:"SecurityGroupAssociationStatisticsSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupAssociationStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupAssociationStatisticsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
}

func (r *DescribeSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 安全组规则集合。
		SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet" name:"SecurityGroupPolicySet"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。每次请求的实例的上限为100。参数不支持同时指定SecurityGroupIds和Filters。
	SecurityGroupIds []*string `json:"SecurityGroupIds" name:"SecurityGroupIds" list`
	// 过滤条件，参数不支持同时指定SecurityGroupIds和Filters。
	// <li>project-id - Integer - （过滤条件）项目id。</li>
	// <li>security-group-name - String - （过滤条件）安全组名称。</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量。
	Offset *string `json:"Offset" name:"Offset"`
	// 返回数量。
	Limit *string `json:"Limit" name:"Limit"`
}

func (r *DescribeSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 安全组对象。
		SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet" name:"SecurityGroupSet" list`
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplateGroupsRequest struct {
	*tchttp.BaseRequest
	// 过滤条件。
	// <li>service-template-group-name - String - （过滤条件）协议端口模板集合名称。</li>
	// <li>service-template-group-id - String - （过滤条件）协议端口模板集合实例ID，例如：ppmg-e6dy460g。</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。
	Offset *string `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit" name:"Limit"`
}

func (r *DescribeServiceTemplateGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTemplateGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplateGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 协议端口模板集合。
		ServiceTemplateGroupSet []*ServiceTemplateGroup `json:"ServiceTemplateGroupSet" name:"ServiceTemplateGroupSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTemplateGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTemplateGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplatesRequest struct {
	*tchttp.BaseRequest
	// 过滤条件。
	// <li>service-template-name - String - （过滤条件）协议端口模板名称。</li>
	// <li>service-template-id - String - （过滤条件）协议端口模板实例ID，例如：ppm-e6dy460g。</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。
	Offset *string `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit" name:"Limit"`
}

func (r *DescribeServiceTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 协议端口模板对象。
		ServiceTemplateSet []*ServiceTemplate `json:"ServiceTemplateSet" name:"ServiceTemplateSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest
	// 子网实例ID查询。形如：subnet-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定SubnetIds和Filters。
	SubnetIds []*string `json:"SubnetIds" name:"SubnetIds" list`
	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// <li>subnet-id - String - （过滤条件）Subnet实例名称。</li>
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>cidr-block - String - （过滤条件）vpc的cidr。</li>
	// <li>is-default - Boolean - （过滤条件）是否是默认子网。</li>
	// <li>subnet-name - String - （过滤条件）子网名称。</li>
	// <li>zone - String - （过滤条件）可用区。</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量
	Offset *string `json:"Offset" name:"Offset"`
	// 返回数量
	Limit *string `json:"Limit" name:"Limit"`
}

func (r *DescribeSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 子网对象。
		SubnetSet []*Subnet `json:"SubnetSet" name:"SubnetSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsRequest struct {
	*tchttp.BaseRequest
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcIds []*string `json:"VpcIds" name:"VpcIds" list`
	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// <li>vpc-name - String - （过滤条件）VPC实例名称。</li>
	// <li>is-default - Boolean - （过滤条件）是否默认VPC。</li>
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>cidr-block - String - （过滤条件）vpc的cidr。</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量
	Offset *string `json:"Offset" name:"Offset"`
	// 返回数量
	Limit *string `json:"Limit" name:"Limit"`
}

func (r *DescribeVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的对象数。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// VPC对象。
		VpcSet []*Vpc `json:"VpcSet" name:"VpcSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnConnectionsRequest struct {
	*tchttp.BaseRequest
	// VPN通道实例ID。形如：vpnx-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnConnectionIds和Filters。
	VpnConnectionIds []*string `json:"VpnConnectionIds" name:"VpnConnectionIds" list`
	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定VpnConnectionIds和Filters。
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeVpnConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnConnectionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// VPN通道实例。
		VpnConnectionSet []*VpnConnection `json:"VpnConnectionSet" name:"VpnConnectionSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpnConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnConnectionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewaysRequest struct {
	*tchttp.BaseRequest
	// VPN网关实例ID。形如：vpngw-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnGatewayIds和Filters。
	VpnGatewayIds []*string `json:"VpnGatewayIds" name:"VpnGatewayIds" list`
	// 过滤器对象属性
	Filters []*FilterObject `json:"Filters" name:"Filters" list`
	// 偏移量
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 请求对象个数
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeVpnGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnGatewaysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// VPN网关实例详细信息列表。
		VpnGatewaySet []*VpnGateway `json:"VpnGatewaySet" name:"VpnGatewaySet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpnGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpnGatewaysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachClassicLinkVpcRequest struct {
	*tchttp.BaseRequest
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// CVM实例ID查询。形如：ins-r8hr2upy。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
}

func (r *DetachClassicLinkVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachClassicLinkVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachClassicLinkVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachClassicLinkVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachClassicLinkVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	// CVM实例ID。形如：ins-r8hr2upy。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *DetachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressRequest struct {
	*tchttp.BaseRequest
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId" name:"AddressId"`
	// 表示解绑 EIP 之后是否分配普通公网 IP。取值范围：<br><li>TRUE：表示解绑 EIP 之后分配普通公网 IP。<br><li>FALSE：表示解绑 EIP 之后不分配普通公网 IP。<br>默认取值：FALSE。<br><br>只有满足以下条件时才能指定该参数：<br><li> 只有在解绑主网卡的主内网 IP 上的 EIP 时才能指定该参数。<br><li>解绑 EIP 后重新分配普通公网 IP 操作一个账号每天最多操作 10 次；详情可通过 [DescribeAddressQuota](https://cloud.tencent.com/document/api/213/1378) 接口获取。
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp" name:"ReallocateNormalPublicIp"`
}

func (r *DisassociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCustomerGatewayConfigurationRequest struct {
	*tchttp.BaseRequest
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId" name:"VpnConnectionId"`
	// 对端网关厂商信息对象，可通过DescribeCustomerGatewayVendors获取。
	CustomerGatewayVendor *CustomerGatewayVendor `json:"CustomerGatewayVendor" name:"CustomerGatewayVendor"`
	// 通道接入设备物理接口名称。
	InterfaceName *string `json:"InterfaceName" name:"InterfaceName"`
}

func (r *DownloadCustomerGatewayConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadCustomerGatewayConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCustomerGatewayConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// XML格式配置信息。
		CustomerGatewayConfiguration *string `json:"CustomerGatewayConfiguration" name:"CustomerGatewayConfiguration"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadCustomerGatewayConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadCustomerGatewayConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name" name:"Name"`
	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values" name:"Values" list`
}

type FilterObject struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name" name:"Name"`
	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values" name:"Values" list`
}

type IKEOptionsSpecification struct {
	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBS-192', 'AES-CBC-256', 'DES-CBC'，默认为3DES-CBC
	PropoEncryAlgorithm *string `json:"PropoEncryAlgorithm" name:"PropoEncryAlgorithm"`
	// 认证算法：可选值：'MD5', 'SHA1'，默认为MD5
	PropoAuthenAlgorithm *string `json:"PropoAuthenAlgorithm" name:"PropoAuthenAlgorithm"`
	// 协商模式：可选值：'AGGRESSIVE', 'MAIN'，默认为MAIN
	ExchangeMode *string `json:"ExchangeMode" name:"ExchangeMode"`
	// 本端标识类型：可选值：'ADDRESS', 'FQDN'，默认为ADDRESS
	LocalIdentity *string `json:"LocalIdentity" name:"LocalIdentity"`
	// 对端标识类型：可选值：'ADDRESS', 'FQDN'，默认为ADDRESS
	RemoteIdentity *string `json:"RemoteIdentity" name:"RemoteIdentity"`
	// 本端标识，当LocalIdentity选为ADDRESS时，LocalAddress必填。localAddress默认为vpn网关公网IP
	LocalAddress *string `json:"LocalAddress" name:"LocalAddress"`
	// 对端标识，当RemoteIdentity选为ADDRESS时，RemoteAddress必填
	RemoteAddress *string `json:"RemoteAddress" name:"RemoteAddress"`
	// 本端标识，当LocalIdentity选为FQDN时，LocalFqdnName必填
	LocalFqdnName *string `json:"LocalFqdnName" name:"LocalFqdnName"`
	// 对端标识，当remoteIdentity选为FQDN时，RemoteFqdnName必填
	RemoteFqdnName *string `json:"RemoteFqdnName" name:"RemoteFqdnName"`
	// DH group，指定IKE交换密钥时使用的DH组，可选值：'GROUP1', 'GROUP2', 'GROUP5', 'GROUP14', 'GROUP24'，
	DhGroupName *string `json:"DhGroupName" name:"DhGroupName"`
	// IKE SA Lifetime，单位：秒，设置IKE SA的生存周期，取值范围：60-604800
	IKESaLifetimeSeconds *uint64 `json:"IKESaLifetimeSeconds" name:"IKESaLifetimeSeconds"`
	// IKE版本
	IKEVersion *string `json:"IKEVersion" name:"IKEVersion"`
}

type IPSECOptionsSpecification struct {
	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBC-192', 'AES-CBC-256', 'DES-CBC', 'NULL'， 默认为AES-CBC-128
	EncryptAlgorithm *string `json:"EncryptAlgorithm" name:"EncryptAlgorithm"`
	// 认证算法：可选值：'MD5', 'SHA1'，默认为
	IntegrityAlgorith *string `json:"IntegrityAlgorith" name:"IntegrityAlgorith"`
	// IPsec SA lifetime(s)：单位秒，取值范围：180-604800
	IPSECSaLifetimeSeconds *uint64 `json:"IPSECSaLifetimeSeconds" name:"IPSECSaLifetimeSeconds"`
	// PFS：可选值：'NULL', 'DH-GROUP1', 'DH-GROUP2', 'DH-GROUP5', 'DH-GROUP14', 'DH-GROUP24'，默认为NULL
	PfsDhGroup *string `json:"PfsDhGroup" name:"PfsDhGroup"`
	// IPsec SA lifetime(KB)：单位KB，取值范围：2560-604800
	IPSECSaLifetimeTraffic *uint64 `json:"IPSECSaLifetimeTraffic" name:"IPSECSaLifetimeTraffic"`
}

type InquiryPriceCreateVpnGatewayRequest struct {
	*tchttp.BaseRequest
	// VPN网关计费模式，PREPAID：表示预付费，即包年包月，POSTPAID_BY_HOUR：表示后付费，即按量计费。默认：POSTPAID_BY_HOUR，如果指定预付费模式，参数InstanceChargePrepaid必填。
	InstanceChargeType *string `json:"InstanceChargeType" name:"InstanceChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid" name:"InstanceChargePrepaid"`
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut" name:"InternetMaxBandwidthOut"`
}

func (r *InquiryPriceCreateVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 商品价格。
		Price *Price `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewVpnGatewayRequest struct {
	*tchttp.BaseRequest
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid" name:"InstanceChargePrepaid"`
}

func (r *InquiryPriceRenewVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 商品价格。
		Price *Price `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut" name:"InternetMaxBandwidthOut"`
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 商品价格。
		Price *Price `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *uint64 `json:"Period" name:"Period"`
	// 自动续费标识。取值范围： NOTIFY_AND_AUTO_RENEW：通知过期且自动续费， NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费。默认：NOTIFY_AND_MANUAL_RENEW
	RenewFlag *string `json:"RenewFlag" name:"RenewFlag"`
}

type ItemPrice struct {
	// 按量计费后付费单价，单位：元。
	UnitPrice *float64 `json:"UnitPrice" name:"UnitPrice"`
	// 按量计费后付费计价单元，可取值范围： HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）： GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。
	ChargeUnit *string `json:"ChargeUnit" name:"ChargeUnit"`
	// 预付费商品的原价，单位：元。
	OriginalPrice *float64 `json:"OriginalPrice" name:"OriginalPrice"`
	// 预付费商品的折扣价，单位：元。
	DiscountPrice *float64 `json:"DiscountPrice" name:"DiscountPrice"`
}

type MigrateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	// 弹性网卡当前绑定的CVM实例ID。形如：ins-r8hr2upy。
	SourceInstanceId *string `json:"SourceInstanceId" name:"SourceInstanceId"`
	// 待迁移的目的CVM实例ID。
	DestinationInstanceId *string `json:"DestinationInstanceId" name:"DestinationInstanceId"`
}

func (r *MigrateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MigrateNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MigrateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MigrateNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MigratePrivateIpAddressRequest struct {
	*tchttp.BaseRequest
	// 当内网IP绑定的弹性网卡实例ID，例如：eni-m6dyj72l。
	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId" name:"SourceNetworkInterfaceId"`
	// 待迁移的目的弹性网卡实例ID。
	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId" name:"DestinationNetworkInterfaceId"`
	// 迁移的内网IP地址，例如：10.0.0.6。
	PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
}

func (r *MigratePrivateIpAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MigratePrivateIpAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MigratePrivateIpAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigratePrivateIpAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MigratePrivateIpAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressAttributeRequest struct {
	*tchttp.BaseRequest
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId" name:"AddressId"`
	// 修改后的 EIP 名称。长度上限为20个字符。
	AddressName *string `json:"AddressName" name:"AddressName"`
}

func (r *ModifyAddressAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateAttributeRequest struct {
	*tchttp.BaseRequest
	// IP地址模板实例ID，例如：ipm-mdunqeb6。
	AddressTemplateId *string `json:"AddressTemplateId" name:"AddressTemplateId"`
	// IP地址模板名称。
	AddressTemplateName *string `json:"AddressTemplateName" name:"AddressTemplateName"`
	// 地址信息，支持 IP、CIDR、IP 范围。
	Addresses []*string `json:"Addresses" name:"Addresses" list`
}

func (r *ModifyAddressTemplateAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressTemplateAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressTemplateAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressTemplateAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateGroupAttributeRequest struct {
	*tchttp.BaseRequest
	// IP地址模板集合实例ID，例如：ipmg-2uw6ujo6。
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId" name:"AddressTemplateGroupId"`
	// IP地址模板集合名称。
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName" name:"AddressTemplateGroupName"`
	// IP地址模板实例ID， 例如：ipm-mdunqeb6。
	AddressTemplateIds []*string `json:"AddressTemplateIds" name:"AddressTemplateIds" list`
}

func (r *ModifyAddressTemplateGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressTemplateGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressTemplateGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressTemplateGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressTemplateGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGatewayAttributeRequest struct {
	*tchttp.BaseRequest
	// 对端网关ID，例如：cgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId" name:"CustomerGatewayId"`
	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName *string `json:"CustomerGatewayName" name:"CustomerGatewayName"`
}

func (r *ModifyCustomerGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomerGatewayAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomerGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomerGatewayAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkInterfaceAttributeRequest struct {
	*tchttp.BaseRequest
	// 弹性网卡实例ID，例如：eni-pxir56ns。
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName" name:"NetworkInterfaceName"`
	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription" name:"NetworkInterfaceDescription"`
	// 指定绑定的安全组，例如:['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds" name:"SecurityGroupIds" list`
}

func (r *ModifyNetworkInterfaceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkInterfaceAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkInterfaceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkInterfaceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkInterfaceAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateIpAddressesAttributeRequest struct {
	*tchttp.BaseRequest
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	// 指定的内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses" name:"PrivateIpAddresses" list`
}

func (r *ModifyPrivateIpAddressesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPrivateIpAddressesAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateIpAddressesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateIpAddressesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPrivateIpAddressesAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableAttributeRequest struct {
	*tchttp.BaseRequest
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId" name:"RouteTableId"`
	// 路由表名称。
	RouteTableName *string `json:"RouteTableName" name:"RouteTableName"`
}

func (r *ModifyRouteTableAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRouteTableAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRouteTableAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRouteTableAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAttributeRequest struct {
	*tchttp.BaseRequest
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName" name:"GroupName"`
	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription" name:"GroupDescription"`
}

func (r *ModifySecurityGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
	// 安全组规则集合。 SecurityGroupPolicySet对象必须同时指定新的出（Egress）入（Ingress）站规则。 SecurityGroupPolicy对象不支持自定义索引（PolicyIndex）。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet" name:"SecurityGroupPolicySet"`
}

func (r *ModifySecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySecurityGroupPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateAttributeRequest struct {
	*tchttp.BaseRequest
	// 协议端口模板实例ID，例如：ppm-529nwwj8。
	ServiceTemplateId *string `json:"ServiceTemplateId" name:"ServiceTemplateId"`
	// 协议端口模板名称。
	ServiceTemplateName *string `json:"ServiceTemplateName" name:"ServiceTemplateName"`
	// 支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。
	Services []*string `json:"Services" name:"Services" list`
}

func (r *ModifyServiceTemplateAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceTemplateAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceTemplateAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceTemplateAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateGroupAttributeRequest struct {
	*tchttp.BaseRequest
	// 协议端口模板集合实例ID，例如：ppmg-ei8hfd9a。
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId" name:"ServiceTemplateGroupId"`
	// 协议端口模板集合名称。
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName" name:"ServiceTemplateGroupName"`
	// 协议端口模板实例ID，例如：ppm-4dw6agho。
	ServiceTemplateIds []*string `json:"ServiceTemplateIds" name:"ServiceTemplateIds" list`
}

func (r *ModifyServiceTemplateGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceTemplateGroupAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceTemplateGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceTemplateGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceTemplateGroupAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetAttributeRequest struct {
	*tchttp.BaseRequest
	// 子网实例ID。形如：subnet-pxir56ns。
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `json:"SubnetName" name:"SubnetName"`
	// 子网是否开启广播。
	EnableBroadcast *string `json:"EnableBroadcast" name:"EnableBroadcast"`
}

func (r *ModifySubnetAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubnetAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubnetAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubnetAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAttributeRequest struct {
	*tchttp.BaseRequest
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 私有网络名称，可任意命名，但不得超过60个字符。
	VpcName *string `json:"VpcName" name:"VpcName"`
	// 是否开启组播。true: 开启, false: 关闭。
	EnableMulticast *string `json:"EnableMulticast" name:"EnableMulticast"`
	// DNS地址，最多支持4个，第1个默认为主，其余为备
	DnsServers []*string `json:"DnsServers" name:"DnsServers" list`
	// 域名
	DomainName *string `json:"DomainName" name:"DomainName"`
}

func (r *ModifyVpcAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpcAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpcAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnConnectionAttributeRequest struct {
	*tchttp.BaseRequest
	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId" name:"VpnConnectionId"`
	// VPN通道名称，可任意命名，但不得超过60个字符。
	VpnConnectionName *string `json:"VpnConnectionName" name:"VpnConnectionName"`
	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey" name:"PreShareKey"`
	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases" name:"SecurityPolicyDatabases" list`
	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自保护机制，用户配置网络安全协议。
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification" name:"IKEOptionsSpecification"`
	// IPSec配置，腾讯云提供IPSec安全会话设置。
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification" name:"IPSECOptionsSpecification"`
}

func (r *ModifyVpnConnectionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnConnectionAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnConnectionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpnConnectionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnConnectionAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayAttributeRequest struct {
	*tchttp.BaseRequest
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// VPN网关名称，最大长度不能超过60个字节。
	VpnGatewayName *string `json:"VpnGatewayName" name:"VpnGatewayName"`
	// VPN网关计费模式，目前只支持预付费（即包年包月）到后付费（即按量计费）的转换。即参数只支持：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType" name:"InstanceChargeType"`
}

func (r *ModifyVpnGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnGatewayAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpnGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnGatewayAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NetworkInterface struct {
	// 弹性网卡实例ID，例如：eni-f1xjkw1b。
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	// 弹性网卡名称。
	NetworkInterfaceName *string `json:"NetworkInterfaceName" name:"NetworkInterfaceName"`
	// 弹性网卡描述。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription" name:"NetworkInterfaceDescription"`
	// 子网实例ID。
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// VPC实例ID。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 绑定的安全组。
	GroupSet []*string `json:"GroupSet" name:"GroupSet" list`
	// 是否是主网卡。
	Primary *bool `json:"Primary" name:"Primary"`
	// MAC地址。
	MacAddress *string `json:"MacAddress" name:"MacAddress"`
	// 取值范围：PENDING|AVAILABLE|ATTACHING|DETACHING|DELETING。
	State *string `json:"State" name:"State"`
	// 内网IP信息。
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet" name:"PrivateIpAddressSet" list`
	// 绑定的云服务器对象。
	Attachment []*InstanceChargePrepaid `json:"Attachment" name:"Attachment" list`
	// 可用区。
	Zone *string `json:"Zone" name:"Zone"`
	// 创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type Price struct {
	// 实例价格。
	InstancePrice *ItemPrice `json:"InstancePrice" name:"InstancePrice"`
	// 网络价格。
	BandwidthPrice *ItemPrice `json:"BandwidthPrice" name:"BandwidthPrice"`
}

type PrivateIpAddressSpecification struct {
	// 内网IP地址。
	PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
	// 是否是主IP。
	Primary *bool `json:"Primary" name:"Primary"`
	// 公网IP地址。
	PublicIpAddress *string `json:"PublicIpAddress" name:"PublicIpAddress"`
	// EIP实例ID，例如：eip-11112222。
	AddressId *string `json:"AddressId" name:"AddressId"`
	// 内网IP描述信息。
	Description *string `json:"Description" name:"Description"`
	// 公网IP是否被封堵。
	IsWanIpBlocked *bool `json:"IsWanIpBlocked" name:"IsWanIpBlocked"`
}

type Quota struct {
	// 配额名称，取值范围：<br><li>`TOTAL_EIP_QUOTA`：用户当前地域下EIP的配额数；<br><li>`DAILY_EIP_APPLY`：用户当前地域下今日申购次数；<br><li>`DAILY_PUBLIC_IP_ASSIGN`：用户当前地域下，重新分配公网 IP次数。
	QuotaId *string `json:"QuotaId" name:"QuotaId"`
	// 当前数量
	QuotaCurrent *int64 `json:"QuotaCurrent" name:"QuotaCurrent"`
	// 配额数量
	QuotaLimit *int64 `json:"QuotaLimit" name:"QuotaLimit"`
}

type ReleaseAddressesRequest struct {
	*tchttp.BaseRequest
	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：`eip-11112222`。
	AddressIds []*string `json:"AddressIds" name:"AddressIds" list`
}

func (r *ReleaseAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewVpnGatewayRequest struct {
	*tchttp.BaseRequest
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// 预付费计费模式。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid" name:"InstanceChargePrepaid"`
}

func (r *RenewVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceRouteTableAssociationRequest struct {
	*tchttp.BaseRequest
	// 子网实例ID，例如：subnet-3x5lf5q0。可通过DescribeSubnets接口查询。
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId" name:"RouteTableId"`
}

func (r *ReplaceRouteTableAssociationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceRouteTableAssociationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceRouteTableAssociationResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceRouteTableAssociationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceRouteTableAssociationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceRoutesRequest struct {
	*tchttp.BaseRequest
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId" name:"RouteTableId"`
	// 路由策略对象。只需要指定路由策略ID（RouteId）。
	Routes []*Route `json:"Routes" name:"Routes" list`
}

func (r *ReplaceRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceSecurityGroupPolicyRequest struct {
	*tchttp.BaseRequest
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
	// 安全组规则集合对象。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet" name:"SecurityGroupPolicySet"`
}

func (r *ReplaceSecurityGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceSecurityGroupPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceSecurityGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceSecurityGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceSecurityGroupPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetRoutesRequest struct {
	*tchttp.BaseRequest
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId" name:"RouteTableId"`
	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName *string `json:"RouteTableName" name:"RouteTableName"`
	// 路由策略。
	Routes []*Route `json:"Routes" name:"Routes" list`
}

func (r *ResetRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetRoutesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetRoutesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetRoutesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetVpnConnectionRequest struct {
	*tchttp.BaseRequest
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId" name:"VpnConnectionId"`
}

func (r *ResetVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetVpnConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetVpnConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetVpnGatewayInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut" name:"InternetMaxBandwidthOut"`
}

func (r *ResetVpnGatewayInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetVpnGatewayInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetVpnGatewayInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetVpnGatewayInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetVpnGatewayInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Route struct {
	// 路由策略ID。
	RouteId *uint64 `json:"RouteId" name:"RouteId"`
	// 目的网段，取值不能在私有网络网段内，例如：112.20.51.0/24。
	DestinationCidrBlock *string `json:"DestinationCidrBlock" name:"DestinationCidrBlock"`
	// 下一跳类型，目前我们支持的类型有：CVM：公网网关类型的云主机；VPN：vpn网关； DIRECTCONNECT：专线网关；PEERCONNECTION：对等连接；SSLVPN：sslvpn网关；NAT：nat网关; NORMAL_CVM：普通云主机。
	GatewayType *string `json:"GatewayType" name:"GatewayType"`
	// 下一跳地址，这里只需要指定不同下一跳类型的网关ID，系统会自动匹配到下一跳地址。
	GatewayId *string `json:"GatewayId" name:"GatewayId"`
	// 路由策略描述。
	RouteDescription *string `json:"RouteDescription" name:"RouteDescription"`
}

type RouteTable struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId" name:"RouteTableId"`
	// 路由表名称。
	RouteTableName *string `json:"RouteTableName" name:"RouteTableName"`
	// 路由表关联关系。
	AssociationSet []*RouteTableAssociation `json:"AssociationSet" name:"AssociationSet" list`
	// 路由表策略集合。
	RouteSet []*Route `json:"RouteSet" name:"RouteSet" list`
	// 是否默认路由表。
	Main *bool `json:"Main" name:"Main"`
	// 创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type RouteTableAssociation struct {
	// 子网实例ID。
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId" name:"RouteTableId"`
}

type SecurityGroup struct {
	// 项目id，默认0。可在qcloud控制台项目管理页面查询到。
	ProjectId *string `json:"ProjectId" name:"ProjectId"`
	// 安全组实例ID，例如：sg-ohuuioma。
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
	// 安全组名称，可任意命名，但不得超过60个字符。
	SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
	// 安全组备注，最多100个字符。
	SecurityGroupDesc *string `json:"SecurityGroupDesc" name:"SecurityGroupDesc"`
	// 是否是默认安全组，默认安全组不支持删除。
	IsDefault *bool `json:"IsDefault" name:"IsDefault"`
	// 安全组创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type SecurityGroupAssociationStatistics struct {
	// 安全组实例ID。
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
	// 云主机实例数。
	CVM *uint64 `json:"CVM" name:"CVM"`
	// 数据库实例数。
	CDB *uint64 `json:"CDB" name:"CDB"`
	// 弹性网卡实例数。
	ENI *uint64 `json:"ENI" name:"ENI"`
	// 被安全组引用数。
	SG *uint64 `json:"SG" name:"SG"`
}

type SecurityGroupPolicy struct {
	// 安全组规则索引号。
	PolicyIndex *int64 `json:"PolicyIndex" name:"PolicyIndex"`
	// 协议, 取值: TCP,UDP, ICMP。
	Protocol *string `json:"Protocol" name:"Protocol"`
	// 端口(all, 离散port,  range)。
	Port *string `json:"Port" name:"Port"`
	// 协议端口ID或者协议端口组ID。ServiceTemplate和Protocol+Port互斥。
	ServiceTemplate []*string `json:"ServiceTemplate" name:"ServiceTemplate" list`
	// 网段或IP(互斥)。
	CidrBlock *string `json:"CidrBlock" name:"CidrBlock"`
	// 已绑定安全组的网段或IP。
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
	// IP地址ID或者ID地址组ID。
	AddressTemplate *string `json:"AddressTemplate" name:"AddressTemplate"`
	// ACCEPT 或 DROP。
	Action *string `json:"Action" name:"Action"`
	// 安全组规则描述。
	PolicyDescription *string `json:"PolicyDescription" name:"PolicyDescription"`
}

type SecurityGroupPolicySet struct {
	// 安全组规则当前版本。用户每次更新安全规则版本会自动加1，防止更新的路由规则已过期，不填不考虑冲突。
	Version *string `json:"Version" name:"Version"`
	// 出站规则。
	Egress []*SecurityGroupPolicy `json:"Egress" name:"Egress" list`
	// 入站规则。
	Ingress []*SecurityGroupPolicy `json:"Ingress" name:"Ingress" list`
}

type SecurityPolicyDatabase struct {
	// 本端网段
	LocalCidrBlock *string `json:"LocalCidrBlock" name:"LocalCidrBlock"`
	// 对端网段
	RemoteCidrBlock []*string `json:"RemoteCidrBlock" name:"RemoteCidrBlock" list`
}

type ServiceTemplate struct {
	// 协议端口实例ID，例如：ppm-f5n1f8da。
	ServiceTemplateId *string `json:"ServiceTemplateId" name:"ServiceTemplateId"`
	// 模板名称。
	ServiceTemplateName *string `json:"ServiceTemplateName" name:"ServiceTemplateName"`
	// 协议端口信息。
	ServiceSet []*string `json:"ServiceSet" name:"ServiceSet" list`
	// 创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type ServiceTemplateGroup struct {
	// 协议端口模板集合实例ID，例如：ppmg-2klmrefu。
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId" name:"ServiceTemplateGroupId"`
	// 协议端口模板集合名称。
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName" name:"ServiceTemplateGroupName"`
	// 协议端口模板实例ID。
	ServiceTemplateIdSet []*string `json:"ServiceTemplateIdSet" name:"ServiceTemplateIdSet" list`
	// 创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type Subnet struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 子网实例ID，例如：subnet-bthucmmy。
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 子网名称。
	SubnetName *string `json:"SubnetName" name:"SubnetName"`
	// 子网的CIDR。
	CidrBlock *string `json:"CidrBlock" name:"CidrBlock"`
	// 是否默认子网。
	IsDefault *bool `json:"IsDefault" name:"IsDefault"`
	// 是否开启广播。
	EnableBroadcast *bool `json:"EnableBroadcast" name:"EnableBroadcast"`
	// 可用区。
	Zone *string `json:"Zone" name:"Zone"`
	// 路由表实例ID，例如：rtb-l2h8d7c2。
	RouteTableId *string `json:"RouteTableId" name:"RouteTableId"`
	// 创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
	// 可用IP数。
	AvailableIpAddressCount *uint64 `json:"AvailableIpAddressCount" name:"AvailableIpAddressCount"`
}

type TransformAddressRequest struct {
	*tchttp.BaseRequest
	// 待操作有普通公网 IP 的实例 ID。实例 ID 形如：`ins-11112222`。可通过登录[控制台](https://console.cloud.tencent.com/cvm)查询，也可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/9389) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *TransformAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransformAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransformAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransformAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransformAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnassignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	// 指定的内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses" name:"PrivateIpAddresses" list`
}

func (r *UnassignPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnassignPrivateIpAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnassignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnassignPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnassignPrivateIpAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Vpc struct {
	// Vpc名称。
	VpcName *string `json:"VpcName" name:"VpcName"`
	// VPC实例ID，例如：vpc-azd4dt1c。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// VPC的cidr，只能为10.0.0.0/16，172.16.0.0/12，192.168.0.0/16这三个内网网段内。
	CidrBlock *string `json:"CidrBlock" name:"CidrBlock"`
	// 是否默认VPC。
	IsDefault *bool `json:"IsDefault" name:"IsDefault"`
	// 是否开启组播。
	EnableMulticast *bool `json:"EnableMulticast" name:"EnableMulticast"`
	// 创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type VpnConnection struct {
	// 通道实例ID。
	VpnConnectionId *string `json:"VpnConnectionId" name:"VpnConnectionId"`
	// 通道名称。
	VpnConnectionName *string `json:"VpnConnectionName" name:"VpnConnectionName"`
	// VPC实例ID。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// 对端网关实例ID。
	CustomerGatewayId *string `json:"CustomerGatewayId" name:"CustomerGatewayId"`
	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey" name:"PreShareKey"`
	// 通道传输协议。
	VpnProto *string `json:"VpnProto" name:"VpnProto"`
	// 通道加密协议。
	EncryptProto *string `json:"EncryptProto" name:"EncryptProto"`
	// 路由类型。
	RouteType *string `json:"RouteType" name:"RouteType"`
	// 创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
	// 通道的生产状态，PENDING：生产中，AVAILABLE：运行中，DELETING：删除中。
	State *string `json:"State" name:"State"`
	// 通道连接状态，AVAILABLE：已连接。
	NetStatus *string `json:"NetStatus" name:"NetStatus"`
	// SPD。
	SecurityPolicyDatabaseSet []*SecurityPolicyDatabase `json:"SecurityPolicyDatabaseSet" name:"SecurityPolicyDatabaseSet" list`
	// IKE选项。
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification" name:"IKEOptionsSpecification"`
	// IPSEC选择。
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification" name:"IPSECOptionsSpecification"`
}

type VpnGateway struct {
	// 网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	// VPC实例ID。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 网关实例名称。
	VpnGatewayName *string `json:"VpnGatewayName" name:"VpnGatewayName"`
	// 网关实例类型：'IPSEC', 'SSL'。
	Type *string `json:"Type" name:"Type"`
	// 网关实例状态， 'PENDING'：生产中，'DELETING'：删除中，'AVAILABLE'：运行中。
	State *string `json:"State" name:"State"`
	// 网关公网IP。
	PublicIpAddress *string `json:"PublicIpAddress" name:"PublicIpAddress"`
	// 网关续费类型：'NOTIFY_AND_MANUAL_RENEW'：手动续费，'NOTIFY_AND_AUTO_RENEW'：自动续费
	RenewFlag *string `json:"RenewFlag" name:"RenewFlag"`
	// 网关付费类型：POSTPAID_BY_HOUR：按小时后付费，PREPAID：包年包月预付费，
	InstanceChargeType *string `json:"InstanceChargeType" name:"InstanceChargeType"`
	// 网关出带宽。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut" name:"InternetMaxBandwidthOut"`
	// 创建时间。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
	// 预付费网关过期时间。
	ExpiredTime *string `json:"ExpiredTime" name:"ExpiredTime"`
	// 公网IP是否被封堵。
	IsAddressBlocked *bool `json:"IsAddressBlocked" name:"IsAddressBlocked"`
	// 计费模式变更，PREPAID_TO_POSTPAID：包年包月预付费到期转按小时后付费。
	NewPurchasePlan *string `json:"NewPurchasePlan" name:"NewPurchasePlan"`
	// 网关计费装，PROTECTIVELY_ISOLATED：被安全隔离的实例，NORMAL：正常。
	RestrictState *string `json:"RestrictState" name:"RestrictState"`
}

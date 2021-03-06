# QCloud Plugin
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![](https://img.shields.io/badge/language-golang-orang.svg)

English / [中文](README_CN.md)

## Introduction

QCloud plugin is an open-source project used by WeCube to manage the life cycle of IaaS and PaaS resources on Tencent Cloud.

The QCloud plugin makes resource management feasible by providing two kinds of APIs such as:

1. Basic API
	
	It simplifies the native QCloud API parameters to make these APIs more friendly and easyier to use. We now support the following resource APIs: VPC, SUBNET, ROUTETABLE, CVM, CLB, MySQL, Redis, MariaDB, etc.

1. Advanced API
	
	It provides combinational business capabilities based on TencentCloud's native APIs to complete more complex operations.

QCloud plugin 1.0.0 is now released, its architecture & APIs are as follows: 
<img src="./docs/compile/images/plugin_function_en.png" />


## Build and Run Docker Image

Before executing the following commands, please make sure docker command is installed on the CentOS host.

[How to Install Docker](https://docs.docker.com/install/linux/docker-ce/centos/)

1. Git clone source code 

```
git clone https://github.com/WeBankPartners/wecube-plugins-qcloud.git
```

2. Build plugin binary

```shell script
make build 
```

![qcloud_build](docs/compile/images/qcloud_build.png)

3. Build plugin docker image, the docker image tag is Github's commit number.

```shell script
make image
```

![qcloud_image](docs/compile/images/qcloud_image.png)

4. Run plugin container. Please replace variable `{$IMAGE_TAG}` with your image tag and execute the following command.

```shell script
docker run -d -p 8081:8081 --restart=unless-stopped -v /etc/localtime:/etc/localtime  wecube-plugins-qcloud:{$IMAGE_TAG}
```

5. On the same CentOS server, use curl command to check if QCloud plugin works fine. Please replace variable `{$your_SecretID}` and `{$your_SecretKey}` with your Tencent Cloud account's secretID and secretKey. If you see a new vpc with CIDR 10.5.0.0/16 has been created on Tencent Cloud, that means the plugin works fine.

```shell script
curl -X POST http://127.0.0.1:8081/v1/qcloud/vpc/create -H "cache-control: no-cache" -H "content-type: application/json" -d "{\"inputs\":[{\"provider_params\": \"Region=ap-shanghai;AvailableZone=ap-shanghai-1;SecretID={$your_SecretID};SecretKey={$your_SecretKey}\",\"name\": \"api_test_vpc\",\"cidr_block\": \"10.5.0.0/16\"}]}"
```

## Build Plugin Package for WeCube

If you want to build a plugin package to work with WeCube, please execute the following command. You can replace variable `{$package_version}` with the version number you want.

```shell script
make package PLUGIN_VERSION=v{$package_version}
```

![qcloud_package](docs/compile/images/qcloud_plugin_package.png)

## How to Work with Wecube
Wecube work with plugin by use Wecube's plugin management function. Plugin package have a config file called register.xml,which describe the plugin api's input and output parameters.

A register.xml's content have following parts:

1. run plugin instance as docker container parameter. 
```
    <package name="qcloud-resource-management" version="v1.10">
    <docker-image-file>wecube-plugins-qcloud.tar</docker-image-file>
    <docker-image-repository>wecube-plugins-qcloud</docker-image-repository>
    <docker-image-tag>v1.10</docker-image-tag>
    <container-port>8081</container-port>
    <container-start-param>-v /etc/localtime:/etc/localtime -v /home/app/wecube-plugins-qcloud/logs:/home/app/wecube-plugins-qcloud/logs</container-start-param>
```

2. plugin api description.
```
<plugin id="vpc" name="Vpc Management" >
        <interface name="create" path="/v1/qcloud/vpc/create">
            <input-parameters>
                <parameter datatype="string">guid</parameter>
                <parameter datatype="string">provider_params</parameter>
                <parameter datatype="string">name</parameter>
                <parameter datatype="string">cidr_block</parameter>
                <parameter datatype="string">id</parameter>
            </input-parameters>
            <output-parameters>
                <parameter datatype="string">guid</parameter>
                <parameter datatype="string">id</parameter>
            </output-parameters>
        </interface>
</plugin>

```

The above api descirbe the input and output as following curl command:

```
curl -X POST \
  http://127.0.0.1:8080/v1/qcloud/vpc/create \ // example
  -H 'content-type: application/json' \
  -d '{
	"inputs":[
		{
		"guid": "1234",
		"name": "VPC-C",
		"cidr_block": "10.5.0.0/16",
		"provider_params": "Region=ap-chengdu;AvailableZone=ap-chengdu-1;SecretID=xxx;SecretKey=xxxx"
		}
	]
}'

curl result as following:
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "guid": "1234",
                "id": "vpc-12345"
            }
        ]
    }
} 

```

## License
QCloud Plugin is licensed under the Apache License Version 2.0.

## Community
- For quick response, please [raise an issue](https://github.com/WeBankPartners/wecube-plugins-qcloud/issues/new/choose) to us, or you can also scan the following QR code to join our community, we will provide feedback as quickly as we can.

	<div align="left">
	<img src="https://github.com/WeBankPartners/we-cmdb/blob/master/cmdb-wiki/images/wecube_qr_code.png"  height="200" width="200">
	</div>

- Contact us: fintech@webank.com






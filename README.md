# terrafmtter
[![OSCS Status](https://www.oscs1024.com/platform/badge/Pangjiping/terrafmtter.svg?size=small)](https://www.oscs1024.com/project/Pangjiping/terrafmtter?ref=badge_small)

Terraform code generation tool by Alibaba Cloud Container Service for Kubernetes [(ACK)](https://www.alibabacloud.com/product/kubernetes).

- This tool will be compatible with all the latest terraform services of Alibaba Cloud. 
- The latest provider in [terraform-provider-alicloud](https://registry.terraform.io/providers/aliyun/alicloud/latest).
- The github repo in [terraform-provider-alicloud](https://github.com/aliyun/terraform-provider-alicloud).


**NOTE:**
- It's only for MacOS now.

## Requirements
- GoLang version 1.17+
- Git
- Network with access to GitHub


## Install
You could install this tool by `git`.
```shell
mkdir -p $GOPATH/src/github.com/Pangjiping
cd $GOPATH/src/github.com/Pangjiping
git clone git@github.com:Pangjiping/terrafmtter.git
cd terrafmtter
make install
```

## Usage
We will add the executable file to your environment variables during installation.
So you could run `terrafmtter` in any directory.
`terrafmtter` supports four command line arguments:
- `resource`: It's required. You could select multiple resources which you want to create, like `cs_kubernetes,cs_kubernetes_node_pool`. Please use `,` to separate multiple resources.
- `data`: It's required. You could select multiple datasources which you want to query, like `cs_kubernetes_version,cs_kubernetes_clusters`. Please use `,` to separate multiple datasources.
- `version`: It's optional. You could specify the alicloud-provider version like `1.173.0`. It's will be `latest` version if not assigned.
- `file`: It's optional. You could specify the location of the output terraform file. 

Below is a simple use case for terrafmtter：
```shell
terrafmtter --data cs_kubernetes_version \
--version 1.173.0 \
--file main.tf
```

The auto-generated `main.tf` is follows:
```terraform
# Code generated by Alibaba Cloud Container Service for Kubernetes.
# More in https://www.alibabacloud.com/product/kubernetes
# We would not try to get your accessKey and secretKey for safety.
# PLEASE EDIT TODOs as needed. ^_^

# [TODO] PLEASE EDIT your secret information.
# Details in https://registry.terraform.io/providers/aliyun/alicloud/latest/docs
provider "alicloud" {
    access_key = "********"
    secret_key = "********"
    region = "cn-hangzhou"
}

# [TODO] PLEASE EDIT.
variable "default" {
    default = "tf-create"
}


# [TODO] PLEASE EDIT.
# Details in https://registry.terraform.io/providers/aliyun/alicloud/1.173.0/docs/data-sources/cs_kubernetes_version
data "alicloud_cs_kubernetes_version" "default" {

  # Details in https://registry.terraform.io/providers/aliyun/alicloud/1.173.0/docs/data-sources/cs_kubernetes_version#cluster_type
  cluster_type = "TODO"
  # Details in https://registry.terraform.io/providers/aliyun/alicloud/1.173.0/docs/data-sources/cs_kubernetes_version#version
  version = "TODO"
  # Details in https://registry.terraform.io/providers/aliyun/alicloud/1.173.0/docs/data-sources/cs_kubernetes_version#profile
  profile = "TODO"

}
```

## Enhancements for ack (future)
- auto attach resource and datasource.
- set recommended fields.
- full field enumeration.

## TODOs
- template enhancement.
- setting default value.
- optimize optional and required.
- optimize for ack.



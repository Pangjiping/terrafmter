terraform-provider-alicloud/cs\_managed\_kubernetes\_clusters.html.markdown at v1.173.0 · aliyun/terraform-provider-alicloud · GitHub

[Skip to content](#start-of-content)

[Homepage](https://github.com/)

[Sign up](/signup?ref_cta=Sign+up&ref_loc=header+logged+out&ref_page=%2F%3Cuser-name%3E%2F%3Crepo-name%3E%2Fblob%2Fshow&source=header-repo)

- Product




- [Features](/features)
- [Mobile](/mobile)
- [Actions](/features/actions)
- [Codespaces](/features/codespaces)
- [Copilot](/features/copilot)
- [Packages](/features/packages)
- [Security](/features/security)
- [Code review](/features/code-review)
- [Issues](/features/issues)
- [Integrations](/features/integrations)
- [GitHub Sponsors](/sponsors)
- [Customer stories](/customer-stories)

- [Team](/team)
- [Enterprise](/enterprise)
- Explore




- [Explore GitHub](/explore)
- Learn and contribute
- [Topics](/topics)
- [Collections](/collections)
- [Trending](/trending)
- [Skills](https://skills.github.com/)
- [GitHub Sponsors](/sponsors/explore)
- [Open source guides](https://opensource.guide)
- Connect with others
- [The ReadME Project](/readme)
- [Events](/events)
- [Community forum](https://github.community)
- [GitHub Education](https://education.github.com)
- [GitHub Stars program](https://stars.github.com)

- [Marketplace](/marketplace)
- Pricing




- [Plans](/pricing)
- [Compare plans](/pricing#compare-features)
- [Contact Sales](https://github.com/enterprise/contact)
- [Education](https://education.github.com)

- In this repository

All GitHub
↵


Jump to
↵


- No suggested jump to results

- In this repository

All GitHub
↵


Jump to
↵

- In this organization

All GitHub
↵


Jump to
↵

- In this repository

All GitHub
↵


Jump to
↵


[Sign in](/login?return_to=https%3A%2F%2Fgithub.com%2Faliyun%2Fterraform-provider-alicloud%2Fblob%2Fv1.173.0%2Fwebsite%2Fdocs%2Fd%2Fcs_managed_kubernetes_clusters.html.markdown)

[Sign up](/signup?ref_cta=Sign+up&ref_loc=header+logged+out&ref_page=%2F%3Cuser-name%3E%2F%3Crepo-name%3E%2Fblob%2Fshow&source=header-repo&source_repo=aliyun%2Fterraform-provider-alicloud)

{{ message }}

[aliyun](/aliyun)/ **[terraform-provider-alicloud](/aliyun/terraform-provider-alicloud)** Public

- [Notifications](/login?return_to=%2Faliyun%2Fterraform-provider-alicloud)
- [Fork\
425](/login?return_to=%2Faliyun%2Fterraform-provider-alicloud)
- [Star\
458](/login?return_to=%2Faliyun%2Fterraform-provider-alicloud)


- [Code](/aliyun/terraform-provider-alicloud/tree/v1.173.0)
- [Issues312](/aliyun/terraform-provider-alicloud/issues)
- [Pull requests104](/aliyun/terraform-provider-alicloud/pulls)
- [Actions](/aliyun/terraform-provider-alicloud/actions)
- [Security](/aliyun/terraform-provider-alicloud/security)
- [Insights](/aliyun/terraform-provider-alicloud/pulse)

More

- [Code](/aliyun/terraform-provider-alicloud/tree/v1.173.0)
- [Issues](/aliyun/terraform-provider-alicloud/issues)
- [Pull requests](/aliyun/terraform-provider-alicloud/pulls)
- [Actions](/aliyun/terraform-provider-alicloud/actions)
- [Security](/aliyun/terraform-provider-alicloud/security)
- [Insights](/aliyun/terraform-provider-alicloud/pulse)

[Permalink](/aliyun/terraform-provider-alicloud/blob/5651c2d752e81a14a712f8261bdc2b132d3a089b/website/docs/d/cs_managed_kubernetes_clusters.html.markdown)

v1.173.0

Switch branches/tags

BranchesTags

Could not load branches

Nothing to show

[{{ refName }}default](https://github.com/aliyun/terraform-provider-alicloud/blob/{{ urlEncodedRefName }}/website/docs/d/cs_managed_kubernetes_clusters.html.markdown)[View all branches](/aliyun/terraform-provider-alicloud/branches)

Could not load tags

Nothing to show

[{{ refName }}default](https://github.com/aliyun/terraform-provider-alicloud/blob/{{ urlEncodedRefName }}/website/docs/d/cs_managed_kubernetes_clusters.html.markdown)

[View all tags](/aliyun/terraform-provider-alicloud/tags)

## [terraform-provider-alicloud](/aliyun/terraform-provider-alicloud/tree/v1.173.0)/[website](/aliyun/terraform-provider-alicloud/tree/v1.173.0/website)/[docs](/aliyun/terraform-provider-alicloud/tree/v1.173.0/website/docs)/[d](/aliyun/terraform-provider-alicloud/tree/v1.173.0/website/docs/d)/ **cs\_managed\_kubernetes\_clusters.html.markdown**

[Go to file](/aliyun/terraform-provider-alicloud/find/v1.173.0)

- [Go to fileT](/aliyun/terraform-provider-alicloud/find/v1.173.0)
- Go to lineL
-
Copy path

- Copy permalink

This commit does not belong to any branch on this repository, and may belong to a fork outside of the repository.

Cannot retrieve contributors at this time

[alicloud\_cs\_managed\_kubernetes\_clusters](#alicloud_cs_managed_kubernetes_clusters) [Example Usage](#example-usage) [Argument Reference](#argument-reference) [Attributes Reference](#attributes-reference) [Block Nodes](#block-nodes) [Block Connections](#block-connections)

77 lines (61 sloc)

3.52 KB


[Display the source blob](/aliyun/terraform-provider-alicloud/blob/v1.173.0/website/docs/d/cs_managed_kubernetes_clusters.html.markdown?plain=1)[Display the rendered blob](/aliyun/terraform-provider-alicloud/blob/v1.173.0/website/docs/d/cs_managed_kubernetes_clusters.html.markdown)

[Raw](/aliyun/terraform-provider-alicloud/raw/v1.173.0/website/docs/d/cs_managed_kubernetes_clusters.html.markdown) [Blame](/aliyun/terraform-provider-alicloud/blame/v1.173.0/website/docs/d/cs_managed_kubernetes_clusters.html.markdown)

Edit this file

E
Open in GitHub Desktop


- [View raw](/aliyun/terraform-provider-alicloud/raw/v1.173.0/website/docs/d/cs_managed_kubernetes_clusters.html.markdown)
- [View blame](/aliyun/terraform-provider-alicloud/blame/v1.173.0/website/docs/d/cs_managed_kubernetes_clusters.html.markdown)

subcategorylayoutpage\_titlesidebar\_currentdescription

Container Service for Kubernetes (ACK)

alicloud

Alicloud: alicloud\_cs\_managed\_kubernetes\_clusters

docs-alicloud-datasource-cs-managed-kubernetes-clusters

Provides a list of Container Service Managed Kubernetes Clusters to be used by the alicloud\_cs\_managed\_kubernetes\_clusters resource.

# alicloud\_cs\_managed\_kubernetes\_clusters

This data source provides a list Container Service Managed Kubernetes Clusters on Alibaba Cloud.

-\> **NOTE:** Available in v1.35.0+

## Example Usage

```
# Declare the data source
data "alicloud_cs_managed_kubernetes_clusters" "k8s_clusters" {
name_regex  = "my-first-k8s"
output_file = "my-first-k8s-json"
}

output "output" {
value = "${data.alicloud_cs_managed_kubernetes_clusters.k8s_clusters.clusters}"
}
```

## Argument Reference

The following arguments are supported:

- `ids` \- (Optional) Cluster IDs to filter.
- `name_regex` \- (Optional) A regex string to filter results by cluster name.
- `output_file` \- (Optional) File name where to save data source results (after running `terraform plan`).
- `enabled_details` \- (Optional) Boolean, false by default, only `id` and `name` are exported. Set to true if more details are needed, e.g., `master_disk_category`, `slb_internet_enabled`, `connections`. See full list in attributes.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

- `ids` \- A list of matched Kubernetes clusters' ids.
- `names` \- A list of matched Kubernetes clusters' names.
- `clusters` \- A list of matched Kubernetes clusters. Each element contains the following attributes:

  - `id` \- The ID of the container cluster.
  - `name` \- The name of the container cluster.
  - `availability_zone` \- The ID of availability zone.
  - `key_name` \- The keypair of ssh login cluster node, you have to create it first.
  - `worker_numbers` \- The ECS instance node number in the current container cluster.
  - `vswitch_ids` \- The ID of VSwitches where the current cluster is located.
  - `vpc_id` \- The ID of VPC where the current cluster is located.
  - `security_group_id` \- The ID of security group where the current cluster worker node is located.
  - `nat_gateway_id` \- The ID of nat gateway used to launch kubernetes cluster.
  - `worker_nodes` \- List of cluster worker nodes. It contains several attributes to `Block Nodes`.
  - `connections` \- Map of kubernetes cluster connection information. It contains several attributes to `Block Connections`.
  - `log_config` \- A list of one element containing information about the associated log store. It contains the following attributes:

    - `type` \- Type of collecting logs.
    - `project` \- Log Service project name.
  - `resource_group_id` \- (Optional, ForceNew, Available in 1.101.0+) The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
  - `cluster_spec` \- (Optional, ForceNew, Available in 1.101.0+) The cluster specifications of kubernetes cluster,which can be empty.Valid values:

    - ack.standard: Standard managed clusters.
    - ack.pro.small: Professional managed clusters.

### Block Nodes

- `id` \- ID of the node.
- `name` \- Node name.
- `private_ip` \- The private IP address of node.
- `role` \- (Deprecated from version 1.9.4)

### Block Connections

- `api_server_internet` \- API Server Internet endpoint.
- `api_server_intranet` \- API Server Intranet endpoint.
- `master_public_ip` \- Master node SSH IP address.
- `service_domain` \- Service Access Domain.

Go


## Footer

[GitHub](https://github.com "GitHub")
© 2022 GitHub, Inc.


### Footer navigation

- [Terms](https://docs.github.com/en/github/site-policy/github-terms-of-service)
- [Privacy](https://docs.github.com/en/github/site-policy/github-privacy-statement)
- [Security](https://github.com/security)
- [Status](https://www.githubstatus.com/)
- [Docs](https://docs.github.com)
- [Contact GitHub](https://support.github.com?tags=dotcom-footer)
- [Pricing](https://github.com/pricing)
- [API](https://docs.github.com)
- [Training](https://services.github.com)
- [Blog](https://github.blog)
- [About](https://github.com/about)

You can’t perform that action at this time.


You signed in with another tab or window. Reload to refresh your session.You signed out in another tab or window. Reload to refresh your session.
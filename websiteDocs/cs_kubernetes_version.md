terraform-provider-alicloud/cs\_kubernetes\_version.html.markdown at v1.173.0 · aliyun/terraform-provider-alicloud · GitHub

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


[Sign in](/login?return_to=https%3A%2F%2Fgithub.com%2Faliyun%2Fterraform-provider-alicloud%2Fblob%2Fv1.173.0%2Fwebsite%2Fdocs%2Fd%2Fcs_kubernetes_version.html.markdown)

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
- [Pull requests102](/aliyun/terraform-provider-alicloud/pulls)
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

[Permalink](/aliyun/terraform-provider-alicloud/blob/5651c2d752e81a14a712f8261bdc2b132d3a089b/website/docs/d/cs_kubernetes_version.html.markdown)

v1.173.0

Switch branches/tags

BranchesTags

Could not load branches

Nothing to show

[{{ refName }}default](https://github.com/aliyun/terraform-provider-alicloud/blob/{{ urlEncodedRefName }}/website/docs/d/cs_kubernetes_version.html.markdown)[View all branches](/aliyun/terraform-provider-alicloud/branches)

Could not load tags

Nothing to show

[{{ refName }}default](https://github.com/aliyun/terraform-provider-alicloud/blob/{{ urlEncodedRefName }}/website/docs/d/cs_kubernetes_version.html.markdown)

[View all tags](/aliyun/terraform-provider-alicloud/tags)

## [terraform-provider-alicloud](/aliyun/terraform-provider-alicloud/tree/v1.173.0)/[website](/aliyun/terraform-provider-alicloud/tree/v1.173.0/website)/[docs](/aliyun/terraform-provider-alicloud/tree/v1.173.0/website/docs)/[d](/aliyun/terraform-provider-alicloud/tree/v1.173.0/website/docs/d)/ **cs\_kubernetes\_version.html.markdown**

[Go to file](/aliyun/terraform-provider-alicloud/find/v1.173.0)

- [Go to fileT](/aliyun/terraform-provider-alicloud/find/v1.173.0)
- Go to lineL
-
Copy path

- Copy permalink

This commit does not belong to any branch on this repository, and may belong to a fork outside of the repository.

Cannot retrieve contributors at this time

[alicloud\_cs\_kubernetes\_version](#alicloud_cs_kubernetes_version) [Example Usage](#example-usage) [Argument Reference](#argument-reference) [Attributes Reference](#attributes-reference)

81 lines (62 sloc)

2.48 KB


[Display the source blob](/aliyun/terraform-provider-alicloud/blob/v1.173.0/website/docs/d/cs_kubernetes_version.html.markdown?plain=1)[Display the rendered blob](/aliyun/terraform-provider-alicloud/blob/v1.173.0/website/docs/d/cs_kubernetes_version.html.markdown)

[Raw](/aliyun/terraform-provider-alicloud/raw/v1.173.0/website/docs/d/cs_kubernetes_version.html.markdown) [Blame](/aliyun/terraform-provider-alicloud/blame/v1.173.0/website/docs/d/cs_kubernetes_version.html.markdown)

Edit this file

E
Open in GitHub Desktop


- [View raw](/aliyun/terraform-provider-alicloud/raw/v1.173.0/website/docs/d/cs_kubernetes_version.html.markdown)
- [View blame](/aliyun/terraform-provider-alicloud/blame/v1.173.0/website/docs/d/cs_kubernetes_version.html.markdown)

subcategorylayoutpage\_titlesidebar\_currentdescription

Container Service for Kubernetes (ACK)

alicloud

Alicloud: alicloud\_cs\_kubernetes\_version

docs-alicloud-datasource-cs-kubernetes-version

Provides a data source to query the details of the Kubernetes version supported by ACK.

# alicloud\_cs\_kubernetes\_version

This data source provides the details of the Kubernetes version supported by ACK.

-\> **NOTE:** Available in 1.170.0+.

## Example Usage

```
# Query the managed kubernetes cluster metadata of version 1.22.3-aliyun.1 in the region specified by the client.
data "alicloud_cs_kubernetes_version" "default" {
cluster_type = "ManagedKubernetes"
version      = "1.22.3-aliyun.1"
profile      = "Default"
}
output "metadata" {
value = data.alicloud_cs_kubernetes_version.default.metadata
}

# Query the kubernetes cluster metadata of version 1.22.3-aliyun.1 in the region specified by the client.
data "alicloud_cs_kubernetes_version" "default" {
cluster_type = "Kubernetes"
version      = "1.22.3-aliyun.1"
profile      = "Default"
}
output "metadata" {
value = data.alicloud_cs_kubernetes_version.default.metadata
}

# Query the serverless kubernetes cluster metadata of version 1.22.3-aliyun.1 in the region specified by the client.
data "alicloud_cs_kubernetes_version" "default" {
cluster_type = "ManagedKubernetes"
version      = "1.22.3-aliyun.1"
profile      = "Serverless"
}
output "metadata" {
value = data.alicloud_cs_kubernetes_version.default.metadata
}

# Query the edge kubernetes cluster metadata of version 1.20.11-aliyunedge.1 in the region specified by the client.
data "alicloud_cs_kubernetes_version" "default" {
cluster_type = "ManagedKubernetes"
version      = "1.20.11-aliyunedge.1"
profile      = "Edge"
}
output "metadata" {
value = data.alicloud_cs_kubernetes_version.default.metadata
}

```

## Argument Reference

The following arguments are supported.

- `cluster_type` \- (Required) The type of cluster. Its valid value are `Kubernetes` and `ManagedKubernetes`.
- `version` \- (Optional) The ACK released kubernetes version.
- `profile` \- (Optional) The profile of cluster. Its valid value are `Default`, `Serverless` and `Edge`.

## Attributes Reference

The following attributes are exported.

- `metadata` \- A list of metadata of kubernetes version.

  - `version` \- The ACK released kubernetes version.
  - `runtime` \- The list of supported runtime.

    - `name` \- The runtime name.
    - `version` \- The runtime version.

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
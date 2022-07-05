package ack

// attach到一个ack资源可能需要的所有resource和datasource

// attachMap key: resource or data
//           value: name
var (
	attachResourceMap   map[string]string
	attachDataSourceMap map[string]string
)

func init() {
	attachResourceMap = map[string]string{
		"cs_kubernetes_node_pool":   "",
		"cs_autoscaling_config":     "",
		"cs_edge_kubernetes":        "",
		"cs_kubernetes":             "",
		"cs_kubernetes_autoscaler":  "",
		"cs_kubernetes_permissions": "",
		"cs_managed_kubernetes":     "vpc,vswitch,key_pair",
		"cs_serverless_kubernetes":  "",
	}

	attachDataSourceMap = map[string]string{
		"cs_kubernetes_node_pool":   "cs_kubernetes_clusters,cs_managed_kubernetes_clusters",
		"cs_autoscaling_config":     "",
		"cs_edge_kubernetes":        "",
		"cs_kubernetes":             "",
		"cs_kubernetes_autoscaler":  "",
		"cs_kubernetes_permissions": "",
		"cs_managed_kubernetes":     "zones,instance_types,cs_kubernetes_version",
		"cs_serverless_kubernetes":  "",
	}
}

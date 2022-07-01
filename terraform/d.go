package terraform

var DataSourceMap map[string]string

func init() {
	DataSourceMap = map[string]string{
		// Container Service for Kubernetes (ACK)
		"ack_service":                       "ack_service",
		"cs_edge_kubernetes_clusters":       "cs_edge_kubernetes_clusters",
		"cs_kubernetes_addon_metadata":      "cs_kubernetes_addon_metadata",
		"cs_kubernetes_addons":              "cs_kubernetes_addons",
		"cs_kubernetes_clusters":            "cs_kubernetes_clusters",
		"cs_kubernetes_permissions":         "cs_kubernetes_permissions",
		"cs_kubernetes_version":             "cs_kubernetes_version",
		"cs_managed_kubernetes_clusters":    "cs_managed_kubernetes_clusters",
		"cs_serverless_kubernetes_clusters": "cs_serverless_kubernetes_clusters",

		// other data sources ...
	}
}

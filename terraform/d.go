package terraform

var DataSourceMap map[string]string

func init() {
	DataSourceMap = map[string]string{
		// Container Service for Kubernetes (ACK)
		"service":                        "ack_service",
		"edge_kubernetes_clusters":       "cs_edge_kubernetes_clusters",
		"kubernetes_addon_metadata":      "cs_kubernetes_addon_metadata",
		"kubernetes_addons":              "cs_kubernetes_addons",
		"kubernetes_clusters":            "cs_kubernetes_clusters",
		"kubernetes_permissions":         "cs_kubernetes_permissions",
		"kubernetes_version":             "cs_kubernetes_version",
		"managed_kubernetes_clusters":    "cs_managed_kubernetes_clusters",
		"serverless_kubernetes_clusters": "cs_serverless_kubernetes_clusters",
	}
}

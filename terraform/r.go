package terraform

var ResourceMap map[string]string

func init() {
	ResourceMap = map[string]string{
		// Container Service for Kubernetes (ACK)
		"kubernetes":             "cs_kubernetes",
		"managed_kubernetes":     "cs_managed_kubernetes",
		"serverless_kubernetes":  "cs_serverless_kubernetes",
		"edge_kubernetes":        "cs_edge_kubernetes",
		"autoscaling_config":     "cs_autoscaling_config",
		"kubernetes_addon":       "cs_kubernetes_addon",
		"kubernetes_autoscaler":  "cs_kubernetes_autoscaler",
		"kubernetes_node_pool":   "cs_kubernetes_node_pool",
		"kubernetes_permissions": "cs_kubernetes_permissions",

		// other resources ...
	}
}

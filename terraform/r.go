package terraform

var ResourceMap map[string]string

func init() {
	ResourceMap = map[string]string{
		// Container Service for Kubernetes (ACK)
		"cs_kubernetes":             "",
		"cs_managed_kubernetes":     "",
		"cs_serverless_kubernetes":  "",
		"cs_edge_kubernetes":        "",
		"cs_autoscaling_config":     "",
		"cs_kubernetes_addon":       "",
		"cs_kubernetes_autoscaler":  "",
		"cs_kubernetes_node_pool":   "",
		"cs_kubernetes_permissions": "",
	}
}

package terraform

var DataSourceMap map[string]string

func init() {
	DataSourceMap = map[string]string{
		"account":                         "account",
		"caller_identity":                 "caller_identity",
		"images":"images",
		"regions":                "regions",
		"zones":               "zones",
		"db_zones":               "db_zones",
		"instance_type_families": "instance_type_families",
		"instance_types":         "instance_types",
		"instances":              "instances",
		"disks":                  "disks",
		"network_interfaces":     "network_interfaces",
		"snapshots":              "snapshots",
		"vpcs":                   "vpcs",
		"vswitches":              "vswitches",
		"eips":                   "eips",
		"key_pairs":              "key_pairs",
		"kms_keys":               "kms_keys",
		"kms_ciphertext":         "kms_ciphertext",
		"kms_plaintext":          "kms_plaintext",
		"dns_resolution_lines":   "dns_resolution_lines",
		"dns_domains":            "dns_domains",
		"dns_groups":             "dns_groups",
		"dns_records":            "dns_records",
		// alicloud_dns_domain_groups, alicloud_dns_domain_records have been deprecated.
		"dns_domain_groups":  "dns_domain_groups",
		"dns_domain_records": "dns_domain_records",
		// alicloud_ram_account_alias has been deprecated
		"ram_account_alias":                           "ram_account_alias",
		"ram_account_aliases":  "ram_account_aliases",
		"ram_groups":  "ram_groups",
		"ram_users":"ram_users",
		"ram_roles":"ram_roles",
		"ram_policies":"ram_policies",
		"security_groups":"security_groups",
		"security_group_rules":"security_group_rules",
		"slbs":"slbs",
		"slb_attachments":"slb_attachments",
		"slb_backend_servers":"slb_backend_servers",
		"slb_listeners":"slb_listeners",
		"slb_rules":  "slb_rules",
		"slb_server_groups":  "slb_server_groups",
		"slb_master_slave_server_groups":"slb_master_slave_server_groups",
		"slb_acls":  "slb_acls",
		"slb_server_certificates":  "slb_server_certificates",
		"slb_ca_certificates":  "slb_ca_certificates",
		"slb_domain_extensions":  "slb_domain_extensions",
		"slb_zones":  "slb_zones",
		"oss_service": "oss_service",
		"oss_bucket_objects":"oss_bucket_objects",
		"oss_buckets":"oss_buckets",
		"ons_instances": "ons_instances",
		"ons_topics":  "ons_topics",
		"ons_groups": "ons_groups",
		"alikafka_consumer_groups":  "alikafka_consumer_groups",
		"alikafka_instances": "alikafka_instances",
		"alikafka_topics": "alikafka_topics",
		"alikafka_sasl_users":"alikafka_sasl_users",
		"alikafka_sasl_acls":  "alikafka_sasl_acls",
		"fc_functions": "fc_functions",
		"file_crc64_checksum": "file_crc64_checksum",
		"fc_services":  "fc_services",
		"fc_triggers": "fc_triggers",
		"fc_custom_domains": "fc_custom_domains",
		"fc_zones": "fc_zones",
		"db_instances": "db_instances",
		"db_instance_engines":  "db_instance_engines",
		"db_instance_classes": "db_instance_classes",
		"rds_backups":  "rds_backups",
		"rds_modify_parameter_logs":"rds_modify_parameter_logs",
		"pvtz_zones": "pvtz_zones",
		"pvtz_zone_records": "pvtz_zone_records",
		"router_interfaces":  "router_interfaces",
		"vpn_gateways":  "vpn_gateways",
		"vpn_customer_gateways": "vpn_customer_gateways",
		"vpn_connections":  "vpn_connections",
		"ssl_vpn_servers": "ssl_vpn_servers",
		"ssl_vpn_client_certs":  "ssl_vpn_client_certs",
		"mongo_instances": "mongo_instances",
		"mongodb_instances": "mongodb_instances",
		"mongodb_zones":   "mongodb_zones",
		"gpdb_instances":   "gpdb_instances",
		"gpdb_zones": "gpdb_zones",
		"kvstore_instances": "kvstore_instances",
		"kvstore_zones": "kvstore_zones",
		"kvstore_permission": "kvstore_permission",
		"kvstore_instance_classes": "kvstore_instance_classes",
		"kvstore_instance_engines": "kvstore_instance_engines",
		"cen_instances":  "cen_instances",
		"cen_bandwidth_packages":"cen_bandwidth_packages",
		"cen_bandwidth_limits": "cen_bandwidth_limits",
		"cen_route_entries":  "cen_route_entries",
		"cen_region_route_entries": "cen_region_route_entries",
		"cen_transit_router_route_entries": "cen_transit_router_route_entries",
		"cen_transit_router_route_table_associations": "cen_transit_router_route_table_associationsd",
		"cen_transit_router_route_table_propagations": "cen_transit_router_route_table_propagations",
		"cen_transit_router_route_tables":"cen_transit_router_route_tables",
		"cen_transit_router_vbr_attachments":  "cen_transit_router_vbr_attachments" ,
		"cen_transit_router_vpc_attachments": "cen_transit_router_vpc_attachments",
		"cen_transit_routers":  "cen_transit_routers",
		"cs_kubernetes_clusters": "cs_kubernetes_clusters",
		"cs_managed_kubernetes_clusters": "cs_managed_kubernetes_clusters",
		"cs_edge_kubernetes_clusters": "cs_edge_kubernetes_clusters",
		"cs_serverless_kubernetes_clusters": "cs_serverless_kubernetes_clusters",
		"cs_kubernetes_permissions": "cs_kubernetes_permissions",
		"cs_kubernetes_addons": "cs_kubernetes_addons",
		"cs_kubernetes_version": "cs_kubernetes_version",
		"cs_kubernetes_addon_metadata": "cs_kubernetes_addon_metadata",
		"cr_namespaces":  "cr_namespaces",
		"cr_repos": "cr_repos",
		"cr_ee_instances": "cr_ee_instances",
		"cr_ee_namespaces": "cr_ee_namespaces",
		"cr_ee_repos": "cr_ee_repos",
		"cr_ee_sync_rules": "cr_ee_sync_rules",
		"mns_queues": "mns_queues",
		"mns_topics": "mns_topics",
		"mns_topic_subscriptions": "mns_topic_subscriptions",
		"api_gateway_service":  "api_gateway_service",
		"api_gateway_apis": "api_gateway_apis",
		"api_gateway_groups":  "api_gateway_groups",
		"api_gateway_apps":  "api_gateway_apps",
		"elasticsearch_instances": "elasticsearch_instances",
		"elasticsearch_zones":  "elasticsearch_zones",
		"drds_instances": "drds_instances",
		"nas_service":  "nas_service",
		"nas_access_groups":"nas_access_groups",
		"nas_access_rules": "nas_access_rules",
		"nas_mount_targets": "nas_mount_targets",
		"nas_file_systems": "nas_file_systems",
		"nas_protocols":  "nas_protocols",
		"cas_certificates": "cas_certificates",
		"common_bandwidth_packages": "common_bandwidth_packages",
		"route_tables":  "route_tables",
		"route_entries": "route_entries",
		"nat_gateways": "nat_gateways",
		"snat_entries": "snat_entries",
		"forward_entries":"forward_entries",
		"ddoscoo_instances":  "ddoscoo_instances",
		"ddosbgp_instances": "ddosbgp_instances",
		"ess_alarms":  "ess_alarms",
		"ess_notifications":  "ess_notifications",
		"ess_scaling_groups": "ess_scaling_groups",
		"ess_scaling_rules":"ess_scaling_rules",
		"ess_scaling_configurations": "ess_scaling_configurations",
		"ess_lifecycle_hooks":  "ess_lifecycle_hooks",
		"ess_scheduled_tasks":"ess_scheduled_tasks",
		"ots_service":  "ots_service",
		"ots_instances": "ots_instances",
		"ots_instance_attachments": "ots_instance_attachments",
		"ots_tables": "ots_tables",
		"ots_tunnels": "ots_tunnels",
		"cloud_connect_networks":  "cloud_connect_networks",
		"emr_instance_types": "emr_instance_types",
		"emr_disk_types":  "emr_disk_types",
		"emr_main_versions": "emr_main_versions",
		"sag_acls": "sag_acls",
		"yundun_dbaudit_instance":"yundun_dbaudit_instance",
		"yundun_bastionhost_instances": "yundun_bastionhost_instances",
		"bastionhost_instances":  "bastionhost_instances",
		"market_product": "market_product",
		"market_products": "market_products" ,
		"polardb_clusters":"polardb_clusters",
		"polardb_node_classes":  "polardb_node_classes",
		"polardb_endpoints": "polardb_endpoints",
		"polardb_accounts":  "polardb_accounts",
		"polardb_databases":  "polardb_databases",
		"polardb_zones":   "polardb_zones",
		"hbase_instances": "hbase_instances",
		"hbase_zones":  "hbase_zones",
		"hbase_instance_types":"hbase_instance_types",
		"adb_clusters": "adb_clusters",
		"adb_zones":  "adb_zones",
		"cen_flowlogs": "cen_flowlogs",
		"kms_aliases":  "kms_aliases",
		"dns_domain_txt_guid": "dns_domain_txt_guid",
		"edas_service": "edas_service",
		"fnf_service": "fnf_service",
		"kms_service":"kms_service",
		"sae_service": "sae_service",
		"dataworks_service":"dataworks_service",
		"data_works_service": "data_works_service",
		"mns_service": "mns_service",
		"cloud_storage_gateway_service":"cloud_storage_gateway_service",
		"vs_service":  "vs_service",
		"pvtz_service": "pvtz_service",
		"cms_service": "cms_service",
		"maxcompute_service": "maxcompute_service",
		"brain_industrial_service":"brain_industrial_service",
		"iot_service": "iot_service",
		"ack_service":  "ack_service",
		"cr_service": "cr_service",
		"dcdn_service":  "dcdn_service",
		"datahub_service": "datahub_service",
		"ons_service": "ons_service",
		"fc_service":  "fc_service",
		"privatelink_service":"privatelink_service",
		"edas_applications": "edas_applications",
		"edas_deploy_groups": "edas_deploy_groups",
		"edas_clusters":  "edas_clusters",
		"resource_manager_folders": "resource_manager_folders",
		"dns_instances": "dns_instances",
		"resource_manager_policies":  "resource_manager_policies",
		"resource_manager_resource_groups": "resource_manager_resource_groups",
		"resource_manager_roles": "resource_manager_roles",
		"resource_manager_policy_versions": "resource_manager_policy_versions",
		"alidns_domain_groups":  "alidns_domain_groups",
		"kms_key_versions": "kms_key_versions",
		"alidns_records": "alidns_records",
		"resource_manager_accounts": "resource_manager_accounts",
		"resource_manager_resource_directories": "resource_manager_resource_directories",
		"resource_manager_handshakes": "resource_manager_handshakes",
		"waf_domains": "waf_domains",
		"kms_secrets": "kms_secrets",
		"cen_route_maps": "cen_route_maps",
		"cen_private_zones":  "cen_private_zones",
		"dms_enterprise_instances": "dms_enterprise_instances",
		"cassandra_clusters": "cassandra_clusters",
		"cassandra_data_centers": "cassandra_data_centers",
		"cassandra_zones":"cassandra_zones",
		"kms_secret_versions": "kms_secret_versions",
		"waf_instances": "waf_instances",
		"eci_image_caches": "eci_image_caches",
		"dms_enterprise_users":"dms_enterprise_users",
		"dms_user_tenants":"dms_user_tenants",
		"ecs_dedicated_hosts":  "ecs_dedicated_hosts",
		"oos_templates": "oos_templates",
		"oos_executions": "oos_executions",
		"resource_manager_policy_attachments":"resource_manager_policy_attachments",
		"dcdn_domains":    "dcdn_domains",
		"mse_clusters": "mse_clusters",
		"actiontrail_trails": "actiontrail_trails",
		"actiontrails": "actiontrails",
		"alidns_instances": "alidns_instances",
		"alidns_domains":  "alidns_domains",
		"log_alert_resource":  "log_alert_resource",
		"log_service":  "log_service",
		"cen_instance_attachments": "cen_instance_attachments",
		"cdn_service":  "cdn_service",
		"cen_vbr_health_checks": "cen_vbr_health_checks",
		"config_rules":"config_rules",
		"config_configuration_recorders": "config_configuration_recorders",
		"config_delivery_channels":"config_delivery_channels",
		"cms_alarm_contacts":  "cms_alarm_contacts",
		"kvstore_connections":  "kvstore_connections",
		"cms_alarm_contact_groups": "cms_alarm_contact_groups",
		"enhanced_nat_available_zones": "enhanced_nat_available_zones",
		"cen_route_services":  "cen_route_services",
		"kvstore_accounts":  "kvstore_accounts",
		"cms_group_metric_rules": "cms_group_metric_rules",
		"fnf_flows":  "fnf_flows",
		"fnf_schedules":  "fnf_schedules",
		"ros_change_sets": "ros_change_sets",
		"ros_stacks":  "ros_stacks",
		"ros_stack_groups": "ros_stack_groups",
		"ros_templates":  "ros_templates",
		"privatelink_vpc_endpoint_services":"privatelink_vpc_endpoint_services",
		"privatelink_vpc_endpoints": "privatelink_vpc_endpoints",
		"privatelink_vpc_endpoint_connections":  "privatelink_vpc_endpoint_connections",
		"privatelink_vpc_endpoint_service_resources":  "privatelink_vpc_endpoint_service_resources",
		"alicloud_privatelink_vpc_endpoint_service_users":      dataSourceAlicloudPrivatelinkVpcEndpointServiceUsers(),
		"alicloud_resource_manager_resource_shares":            dataSourceAlicloudResourceManagerResourceShares(),
		"alicloud_privatelink_vpc_endpoint_zones":              dataSourceAlicloudPrivatelinkVpcEndpointZones(),
		"alicloud_ga_accelerators":                             dataSourceAlicloudGaAccelerators(),
		"alicloud_eci_container_groups":                        dataSourceAlicloudEciContainerGroups(),
		"alicloud_resource_manager_shared_resources":           dataSourceAlicloudResourceManagerSharedResources(),
		"alicloud_resource_manager_shared_targets":             dataSourceAlicloudResourceManagerSharedTargets(),
		"alicloud_ga_listeners":                                dataSourceAlicloudGaListeners(),
		"alicloud_tsdb_instances":                              dataSourceAlicloudTsdbInstances(),
		"alicloud_tsdb_zones":                                  dataSourceAlicloudTsdbZones(),
		"alicloud_ga_bandwidth_packages":                       dataSourceAlicloudGaBandwidthPackages(),
		"alicloud_ga_endpoint_groups":                          dataSourceAlicloudGaEndpointGroups(),
		"alicloud_brain_industrial_pid_organizations":          dataSourceAlicloudBrainIndustrialPidOrganizations(),
		"alicloud_ga_ip_sets":                                  dataSourceAlicloudGaIpSets(),
		"alicloud_ga_forwarding_rules":                         dataSourceAlicloudGaForwardingRules(),
		"alicloud_eipanycast_anycast_eip_addresses":            dataSourceAlicloudEipanycastAnycastEipAddresses(),
		"alicloud_brain_industrial_pid_projects":               dataSourceAlicloudBrainIndustrialPidProjects(),
		"alicloud_cms_monitor_groups":                          dataSourceAlicloudCmsMonitorGroups(),
		"alicloud_ram_saml_providers":                          dataSourceAlicloudRamSamlProviders(),
		"alicloud_quotas_quotas":
		"alicloud_quotas_application_infos":
		"alicloud_cms_monitor_group_instanceses":
		"alicloud_cms_monitor_group_instances":
		"alicloud_quotas_quota_alarms":
		"alicloud_ecs_commands":
		"alicloud_cloud_storage_gateway_storage_bundles":
		"alicloud_ecs_hpc_clusters":
		"alicloud_brain_industrial_pid_loops":
		"alicloud_quotas_quota_applications":
		"alicloud_ecs_auto_snapshot_policies":
		"alicloud_rds_parameter_groups":
		"alicloud_ecs_launch_templates":
		"alicloud_resource_manager_control_policies":
		"alicloud_resource_manager_control_policy_attachments":
		"alicloud_rds_accounts":
		"alicloud_havips":
		"alicloud_ecs_snapshots":
		"alicloud_ecs_key_pairs":
		"alicloud_adb_db_clusters":
		"alicloud_vpc_flow_logs":
		"alicloud_network_acls":
		"alicloud_ecs_disks":
		"alicloud_ddoscoo_domain_resources":
		"alicloud_ddoscoo_ports":
		"alicloud_slb_load_balancers":
		"alicloud_ecs_network_interfaces":
		"alicloud_config_aggregators":
		"alicloud_config_aggregate_config_rules":
		"alicloud_config_aggregate_compliance_packs":
		"alicloud_config_compliance_packs":
		"alicloud_eip_addresses":
		"alicloud_direct_mail_receiverses":
		"alicloud_log_projects":
		"alicloud_log_stores":
		"alicloud_event_bridge_service":
		"alicloud_event_bridge_event_buses":
		"alicloud_amqp_virtual_hosts":
		"alicloud_amqp_queues":
		"alicloud_amqp_exchanges":
		"alicloud_cassandra_backup_plans":
		"alicloud_cen_transit_router_peer_attachments":
		"alicloud_amqp_instances":
		"alicloud_hbr_vaults":
		"alicloud_ssl_certificates_service_certificates":
		"alicloud_arms_alert_contacts":
		"alicloud_event_bridge_rules":
		"alicloud_cloud_firewall_control_policies":
		"alicloud_sae_namespaces":
		"alicloud_sae_config_maps":
		"alicloud_alb_security_policies":
		"alicloud_event_bridge_event_sources":
		"alicloud_ecd_policy_groups":
		"alicloud_ecp_key_pairs":
		"alicloud_hbr_ecs_backup_plans":
		"alicloud_hbr_nas_backup_plans":
		"alicloud_hbr_oss_backup_plans":
		"alicloud_scdn_domains":
		"alicloud_alb_server_groups":
		"alicloud_data_works_folders":
		"alicloud_arms_alert_contact_groups":
		"alicloud_express_connect_access_points":
		"alicloud_cloud_storage_gateway_gateways":
		"alicloud_lindorm_instances":
		"alicloud_express_connect_physical_connection_service":
		"alicloud_cddc_dedicated_host_groups":
		"alicloud_hbr_ecs_backup_clients":
		"alicloud_msc_sub_contacts":
		"alicloud_express_connect_physical_connections":
		"alicloud_alb_load_balancers":
		"alicloud_alb_zones":
		"alicloud_sddp_rules":
		"alicloud_bastionhost_user_groups":
		"alicloud_security_center_groups":
		"alicloud_alb_acls":
		"alicloud_hbr_snapshots":
		"alicloud_bastionhost_users":
		"alicloud_dfs_access_groups":
		"alicloud_ehpc_job_templates":
		"alicloud_sddp_configs":
		"alicloud_hbr_restore_jobs":
		"alicloud_alb_listeners":
		"alicloud_ens_key_pairs":
		"alicloud_sae_applications":
		"alicloud_alb_rules":
		"alicloud_cms_metric_rule_templates":
		"alicloud_iot_device_groups":
		"alicloud_express_connect_virtual_border_routers":
		"alicloud_imm_projects":
		"alicloud_click_house_db_clusters":
		"alicloud_direct_mail_domains":
		"alicloud_bastionhost_host_groups":
		"alicloud_vpc_dhcp_options_sets":
		"alicloud_alb_health_check_templates":
		"alicloud_cdn_real_time_log_deliveries":
		"alicloud_click_house_accounts":
		"alicloud_direct_mail_mail_addresses":
		"alicloud_database_gateway_gateways":
		"alicloud_bastionhost_hosts":
		"amqp_bindings":"amqp_bindings",
		"slb_tls_cipher_policies":"slb_tls_cipher_policies",
		"cloud_sso_directories":"cloud_sso_directories",
		"bastionhost_host_accounts":"bastionhost_host_accounts",
		"waf_certificates":"waf_certificates",
		"simple_application_server_instances":"simple_application_server_instances",
		"simple_application_server_plans":"simple_application_server_plans",
		"simple_application_server_images":"simple_application_server_images",
		"video_surveillance_system_groups":"video_surveillance_system_groups",
		"msc_sub_subscriptions":"msc_sub_subscriptions",
		"sddp_instances":"sddp_instances",
		"vpc_nat_ip_cidrs":"vpc_nat_ip_cidrs",
		"vpc_nat_ips":"vpc_nat_ips",
		"quick_bi_users":"quick_bi_users",
		"vod_domains":"vod_domains",
		"arms_dispatch_rules":"arms_dispatch_rules",
		"open_search_app_groups":"open_search_app_groups",
		"graph_database_db_instances":"graph_database_db_instances",
		"arms_prometheus_alert_rules":"arms_prometheus_alert_rules",
		"dbfs_instances":"dbfs_instances",
		"rdc_organizations":"rdc_organizations",
		"eais_instances":"eais_instances",
		"sae_ingresses":"sae_ingresses",
		"cloudauth_face_configs":"cloudauth_face_configs",
		"imp_app_templates":"imp_app_templates",
		"mhub_products":"mhub_products",
		"cloud_sso_scim_server_credentials":  "cloud_sso_scim_server_credentials",
		"dts_subscription_jobs": "dts_subscription_jobs",
		"service_mesh_service_meshes":"service_mesh_service_meshes",
		"service_mesh_versions":"service_mesh_versions",
		"mhub_apps":"mhub_apps",
		"cloud_sso_groups":"cloud_sso_groups",
		"hbr_backup_jobs":"hbr_backup_jobs",
		"click_house_regions":"click_house_regions",
		"dts_synchronization_jobs":"dts_synchronization_jobs",
		"cloud_firewall_instances":"cloud_firewall_instances",
		"cr_endpoint_acl_policies":"cr_endpoint_acl_policies",
		"cr_endpoint_acl_service":"cr_endpoint_acl_service",
		"actiontrail_history_delivery_jobs":"actiontrail_history_delivery_jobs",
		"sae_instance_specifications":"sae_instance_specifications",
		"cen_transit_router_service":"cen_transit_router_service",
		"ecs_deployment_sets":"ecs_deployment_sets",
		"cloud_sso_users":"cloud_sso_users",
		"cloud_sso_access_configurations":  "cloud_sso_access_configurations",
		"dfs_file_systems": "dfs_file_systems",
		"dfs_zones":  "dfs_zones",
		"vpc_traffic_mirror_filters":  "vpc_traffic_mirror_filters",
		"dfs_access_rules":"dfs_access_rules",
		"nas_zones": "nas_zones",
		"dfs_mount_points": "dfs_mount_points",
		"vpc_traffic_mirror_filter_egress_rules":  "vpc_traffic_mirror_filter_egress_rules",
		"ecd_simple_office_sites": "ecd_simple_office_sites",
		"vpc_traffic_mirror_filter_ingress_rules": "vpc_traffic_mirror_filter_ingress_rules",
		"ecd_nas_file_systems":  "ecd_nas_file_systems",
		"vpc_traffic_mirror_service": "vpc_traffic_mirror_service",
		"msc_sub_webhooks":"msc_sub_webhooks",
		"ecd_users": "ecd_users",
		"vpc_traffic_mirror_sessions":"vpc_traffic_mirror_sessions",
		"gpdb_accounts":  "gpdb_accounts",
		"vpc_ipv6_gateways":  "vpc_ipv6_gateways",
		"vpc_ipv6_egress_rules": "vpc_ipv6_egress_rules",
		"vpc_ipv6_addresses":   "vpc_ipv6_addresses",
		"hbr_server_backup_plans": "hbr_server_backup_plans",
		"cms_dynamic_tag_groups": "cms_dynamic_tag_groups",
		"ecd_network_packages":  "ecd_network_packages",
		"cloud_storage_gateway_gateway_smb_users":"cloud_storage_gateway_gateway_smb_users",
		"vpc_ipv6_internet_bandwidths": "vpc_ipv6_internet_bandwidths",
		"simple_application_server_firewall_rules": "simple_application_server_firewall_rules",
		"pvtz_endpoints": "pvtz_endpoints",
		"pvtz_resolver_zones":   "pvtz_resolver_zones",
		"pvtz_rules": "pvtz_rules",
		"ecd_bundles":   "ecd_bundles",
		"simple_application_server_disks":  "simple_application_server_disks",
		"simple_application_server_snapshots": "simple_application_server_snapshots",
		"simple_application_server_custom_images": "simple_application_server_custom_images",
		"cloud_storage_gateway_stocks":  "cloud_storage_gateway_stocks",
		"cloud_storage_gateway_gateway_cache_disks": "cloud_storage_gateway_gateway_cache_disks",
		"cloud_storage_gateway_gateway_block_volumes":"cloud_storage_gateway_gateway_block_volumes",
		"direct_mail_tags":  "direct_mail_tags",
		"cloud_storage_gateway_gateway_file_shares":"cloud_storage_gateway_gateway_file_shares",
		"ecd_desktops":  "ecd_desktops",
		"cloud_storage_gateway_express_syncs":"cloud_storage_gateway_express_syncs",
		"oos_applications": "oos_applications",
		"eci_virtual_nodes":  "eci_virtual_nodes",
		"eci_zones": "eci_zones",
		"ros_stack_instances":  "ros_stack_instances",
		"ros_regions":  "ros_regions",
		"ecs_dedicated_host_clusters": "ecs_dedicated_host_clusters",
		"oos_application_groups": "oos_application_groups",
		"dts_consumer_channels": "dts_consumer_channels",
		"emr_clusters": "emr_clusters",
		"ecd_images":"ecd_images",
		"oos_patch_baselines": "oos_patch_baselines",
		"ecd_commands": "ecd_commands",
		"cddc_zones":  "cddc_zones",
		"cddc_host_ecs_level_infos":  "cddc_host_ecs_level_infos",
		"cddc_dedicated_hosts": "cddc_dedicated_hosts",
		"oos_parameters": "oos_parameters",
		"oos_state_configurations": "oos_state_configurations",
		"oos_secret_parameters":  "oos_secret_parameters",
		"click_house_backup_policies":"click_house_backup_policies",
		"cloud_sso_service":  "cloud_sso_service",
		"mongodb_audit_policies":   "mongodb_audit_policies",
		"mongodb_accounts":   "mongodb_accounts",
		"mongodb_serverless_instances":"mongodb_serverless_instances",
		"cddc_dedicated_host_accounts":"cddc_dedicated_host_accounts",
		"cr_chart_namespaces": "cr_chart_namespaces",
		"fnf_executions":   "fnf_executions",
		"cr_chart_repositories": "cr_chart_repositories",
		"mongodb_sharding_network_public_addresses":"mongodb_sharding_network_public_addresses",
		"ga_acls": "ga_acls",
		"ga_additional_certificates":  "ga_additional_certificates",
		"alidns_custom_lines":  "alidns_custom_lines",
		"ros_template_scratches":"ros_template_scratches",
		"alidns_gtm_instances":  "alidns_gtm_instances",
		"vpc_bgp_groups": "vpc_bgp_groups",
		"nas_snapshots":  "nas_snapshots",
		"hbr_replication_vault_regions": "hbr_replication_vault_regions",
		"alidns_address_pools": "alidns_address_pools",
		"ecs_prefix_lists":  "ecs_prefix_lists",
		"alidns_access_strategies": "alidns_access_strategies",
		"vpc_bgp_peers": "vpc_bgp_peers",
		"nas_filesets": "nas_filesets",
		"cdn_ip_info": "cdn_ip_info",
		"nas_auto_snapshot_policies": "nas_auto_snapshot_policies",
		"nas_lifecycle_policies":   "nas_lifecycle_policies",
		"vpc_bgp_networks":"vpc_bgp_networks",
		"nas_data_flows": "nas_data_flows",
		"ecs_storage_capacity_units":                           "ecs_storage_capacity_units",
		"dbfs_snapshots":                                       "dbfs_snapshots",
		"msc_sub_contact_verification_message":                 "msc_sub_contact_verification_message",
		"dts_migration_jobs":                                   "dts_migration_jobs",
		"mse_gateways":                                         "mse_gateways",
		"mongodb_sharding_network_private_addresses":           "mongodb_sharding_network_private_addresses",
		"ecp_instances":  "ecp_instances",                                   
		"ecp_zones":  "ecp_zones",                                         
		"ecp_instance_types": "ecp_instance_types",
		"dcdn_ipa_domains": "dcdn_ipa_domains",                                  
		"sddp_data_limits":  "sddp_data_limits",                                 
		"ecs_image_components":  "ecs_image_components",
		"sae_application_scaling_rules":  "sae_application_scaling_rules",
		"sae_grey_tag_routes":   "sae_grey_tag_routes",
		"ecs_snapshot_groups": "ecs_snapshot_groups",
		"vpn_ipsec_servers": "vpn_ipsec_servers",
		"cr_chains":  "cr_chains",
		"vpn_pbr_route_entries": "vpn_pbr_route_entries",
		"mse_znodes": "mse_znodes",
		"cen_transit_router_available_resources":  "cen_transit_router_available_resources",
		"ecs_image_pipelines": "ecs_image_pipelines",
		"hbr_ots_backup_plans": "hbr_ots_backup_plans",
		"hbr_ots_snapshots": "hbr_ots_snapshots",
		"bastionhost_host_share_keys": "bastionhost_host_share_keys",
		"ecs_network_interface_permissions":   "ecs_network_interface_permissions",
		"mse_engine_namespaces":   "mse_engine_namespaces",
		"ga_accelerator_spare_ip_attachments": "ga_accelerator_spare_ip_attachments",
		"smartag_flow_logs": "smartag_flow_logs",
		"ecs_invocations": "ecs_invocations",
		"ecd_snapshots":   "ecd_snapshots",
		"tag_meta_tags":"tag_meta_tags",
		"ecd_desktop_types": "ecd_desktop_types",
		"config_deliveries": "config_deliveries",
		"cms_namespaces":  "cms_namespaces",
		"cms_sls_groups":  "cms_sls_groups",
		"config_aggregate_deliveries":  "config_aggregate_deliveries",
		"edas_namespaces": "edas_namespaces",
		"cdn_blocked_regions": "cdn_blocked_regions",
		"schedulerx_namespaces":                                "schedulerx_namespaces",
		"ehpc_clusters": "ehpc_clusters", 
		"cen_traffic_marking_policies": "cen_traffic_marking_policies",
	}

}

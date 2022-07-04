package terraform

var DataSourceMap map[string]string

func init() {
	DataSourceMap = map[string]string{
		"account":                         "account",
		"caller_identity":                 "caller_identity",
		"alicloud_images":                 dataSourceAlicloudImages(),
		"alicloud_regions":                dataSourceAlicloudRegions(),
		"alicloud_zones":                  dataSourceAlicloudZones(),
		"alicloud_db_zones":               dataSourceAlicloudDBZones(),
		"alicloud_instance_type_families": dataSourceAlicloudInstanceTypeFamilies(),
		"alicloud_instance_types":         dataSourceAlicloudInstanceTypes(),
		"alicloud_instances":              dataSourceAlicloudInstances(),
		"alicloud_disks":                  dataSourceAlicloudEcsDisks(),
		"alicloud_network_interfaces":     dataSourceAlicloudEcsNetworkInterfaces(),
		"alicloud_snapshots":              dataSourceAlicloudEcsSnapshots(),
		"alicloud_vpcs":                   dataSourceAlicloudVpcs(),
		"alicloud_vswitches":              dataSourceAlicloudVswitches(),
		"alicloud_eips":                   dataSourceAlicloudEipAddresses(),
		"alicloud_key_pairs":              dataSourceAlicloudEcsKeyPairs(),
		"alicloud_kms_keys":               dataSourceAlicloudKmsKeys(),
		"alicloud_kms_ciphertext":         dataSourceAlicloudKmsCiphertext(),
		"alicloud_kms_plaintext":          dataSourceAlicloudKmsPlaintext(),
		"alicloud_dns_resolution_lines":   dataSourceAlicloudDnsResolutionLines(),
		"alicloud_dns_domains":            dataSourceAlicloudAlidnsDomains(),
		"alicloud_dns_groups":             dataSourceAlicloudDnsGroups(),
		"alicloud_dns_records":            dataSourceAlicloudDnsRecords(),
		// alicloud_dns_domain_groups, alicloud_dns_domain_records have been deprecated.
		"alicloud_dns_domain_groups":  dataSourceAlicloudDnsGroups(),
		"alicloud_dns_domain_records": dataSourceAlicloudDnsRecords(),
		// alicloud_ram_account_alias has been deprecated
		"alicloud_ram_account_alias":                           dataSourceAlicloudRamAccountAlias(),
		"alicloud_ram_account_aliases":                         dataSourceAlicloudRamAccountAlias(),
		"alicloud_ram_groups":                                  dataSourceAlicloudRamGroups(),
		"alicloud_ram_users":                                   dataSourceAlicloudRamUsers(),
		"alicloud_ram_roles":                                   dataSourceAlicloudRamRoles(),
		"alicloud_ram_policies":                                dataSourceAlicloudRamPolicies(),
		"alicloud_security_groups":                             dataSourceAlicloudSecurityGroups(),
		"alicloud_security_group_rules":                        dataSourceAlicloudSecurityGroupRules(),
		"alicloud_slbs":                                        dataSourceAlicloudSlbLoadBalancers(),
		"alicloud_slb_attachments":                             dataSourceAlicloudSlbAttachments(),
		"alicloud_slb_backend_servers":                         dataSourceAlicloudSlbBackendServers(),
		"alicloud_slb_listeners":                               dataSourceAlicloudSlbListeners(),
		"alicloud_slb_rules":                                   dataSourceAlicloudSlbRules(),
		"alicloud_slb_server_groups":                           dataSourceAlicloudSlbServerGroups(),
		"alicloud_slb_master_slave_server_groups":              dataSourceAlicloudSlbMasterSlaveServerGroups(),
		"alicloud_slb_acls":                                    dataSourceAlicloudSlbAcls(),
		"alicloud_slb_server_certificates":                     dataSourceAlicloudSlbServerCertificates(),
		"alicloud_slb_ca_certificates":                         dataSourceAlicloudSlbCaCertificates(),
		"alicloud_slb_domain_extensions":                       dataSourceAlicloudSlbDomainExtensions(),
		"alicloud_slb_zones":                                   dataSourceAlicloudSlbZones(),
		"alicloud_oss_service":                                 dataSourceAlicloudOssService(),
		"alicloud_oss_bucket_objects":                          dataSourceAlicloudOssBucketObjects(),
		"alicloud_oss_buckets":                                 dataSourceAlicloudOssBuckets(),
		"alicloud_ons_instances":                               dataSourceAlicloudOnsInstances(),
		"alicloud_ons_topics":                                  dataSourceAlicloudOnsTopics(),
		"alicloud_ons_groups":                                  dataSourceAlicloudOnsGroups(),
		"alicloud_alikafka_consumer_groups":                    dataSourceAlicloudAlikafkaConsumerGroups(),
		"alicloud_alikafka_instances":                          dataSourceAlicloudAlikafkaInstances(),
		"alicloud_alikafka_topics":                             dataSourceAlicloudAlikafkaTopics(),
		"alicloud_alikafka_sasl_users":                         dataSourceAlicloudAlikafkaSaslUsers(),
		"alicloud_alikafka_sasl_acls":                          dataSourceAlicloudAlikafkaSaslAcls(),
		"alicloud_fc_functions":                                dataSourceAlicloudFcFunctions(),
		"alicloud_file_crc64_checksum":                         dataSourceAlicloudFileCRC64Checksum(),
		"alicloud_fc_services":                                 dataSourceAlicloudFcServices(),
		"alicloud_fc_triggers":                                 dataSourceAlicloudFcTriggers(),
		"alicloud_fc_custom_domains":                           dataSourceAlicloudFcCustomDomains(),
		"alicloud_fc_zones":                                    dataSourceAlicloudFcZones(),
		"alicloud_db_instances":                                dataSourceAlicloudDBInstances(),
		"alicloud_db_instance_engines":                         dataSourceAlicloudDBInstanceEngines(),
		"alicloud_db_instance_classes":                         dataSourceAlicloudDBInstanceClasses(),
		"alicloud_rds_backups":                                 dataSourceAlicloudRdsBackups(),
		"alicloud_rds_modify_parameter_logs":                   dataSourceAlicloudRdsModifyParameterLogs(),
		"alicloud_pvtz_zones":                                  dataSourceAlicloudPvtzZones(),
		"alicloud_pvtz_zone_records":                           dataSourceAlicloudPvtzZoneRecords(),
		"alicloud_router_interfaces":                           dataSourceAlicloudRouterInterfaces(),
		"alicloud_vpn_gateways":                                dataSourceAlicloudVpnGateways(),
		"alicloud_vpn_customer_gateways":                       dataSourceAlicloudVpnCustomerGateways(),
		"alicloud_vpn_connections":                             dataSourceAlicloudVpnConnections(),
		"alicloud_ssl_vpn_servers":                             dataSourceAlicloudSslVpnServers(),
		"alicloud_ssl_vpn_client_certs":                        dataSourceAlicloudSslVpnClientCerts(),
		"alicloud_mongo_instances":                             dataSourceAlicloudMongoDBInstances(),
		"alicloud_mongodb_instances":                           dataSourceAlicloudMongoDBInstances(),
		"alicloud_mongodb_zones":                               dataSourceAlicloudMongoDBZones(),
		"alicloud_gpdb_instances":                              dataSourceAlicloudGpdbInstances(),
		"alicloud_gpdb_zones":                                  dataSourceAlicloudGpdbZones(),
		"alicloud_kvstore_instances":                           dataSourceAlicloudKvstoreInstances(),
		"alicloud_kvstore_zones":                               dataSourceAlicloudKVStoreZones(),
		"alicloud_kvstore_permission":                          dataSourceAlicloudKVStorePermission(),
		"alicloud_kvstore_instance_classes":                    dataSourceAlicloudKVStoreInstanceClasses(),
		"alicloud_kvstore_instance_engines":                    dataSourceAlicloudKVStoreInstanceEngines(),
		"alicloud_cen_instances":                               dataSourceAlicloudCenInstances(),
		"alicloud_cen_bandwidth_packages":                      dataSourceAlicloudCenBandwidthPackages(),
		"alicloud_cen_bandwidth_limits":                        dataSourceAlicloudCenBandwidthLimits(),
		"alicloud_cen_route_entries":                           dataSourceAlicloudCenRouteEntries(),
		"alicloud_cen_region_route_entries":                    dataSourceAlicloudCenRegionRouteEntries(),
		"alicloud_cen_transit_router_route_entries":            dataSourceAlicloudCenTransitRouterRouteEntries(),
		"alicloud_cen_transit_router_route_table_associations": dataSourceAlicloudCenTransitRouterRouteTableAssociations(),
		"alicloud_cen_transit_router_route_table_propagations": dataSourceAlicloudCenTransitRouterRouteTablePropagations(),
		"alicloud_cen_transit_router_route_tables":             dataSourceAlicloudCenTransitRouterRouteTables(),
		"alicloud_cen_transit_router_vbr_attachments":          dataSourceAlicloudCenTransitRouterVbrAttachments(),
		"alicloud_cen_transit_router_vpc_attachments":          dataSourceAlicloudCenTransitRouterVpcAttachments(),
		"alicloud_cen_transit_routers":                         dataSourceAlicloudCenTransitRouters(),
		"alicloud_cs_kubernetes_clusters":                      dataSourceAlicloudCSKubernetesClusters(),
		"alicloud_cs_managed_kubernetes_clusters":              dataSourceAlicloudCSManagerKubernetesClusters(),
		"alicloud_cs_edge_kubernetes_clusters":                 dataSourceAlicloudCSEdgeKubernetesClusters(),
		"alicloud_cs_serverless_kubernetes_clusters":           dataSourceAlicloudCSServerlessKubernetesClusters(),
		"alicloud_cs_kubernetes_permissions":                   dataSourceAlicloudCSKubernetesPermissions(),
		"alicloud_cs_kubernetes_addons":                        dataSourceAlicloudCSKubernetesAddons(),
		"alicloud_cs_kubernetes_version":                       dataSourceAlicloudCSKubernetesVersion(),
		"alicloud_cs_kubernetes_addon_metadata":                dataSourceAlicloudCSKubernetesAddonMetadata(),
		"alicloud_cr_namespaces":                               dataSourceAlicloudCRNamespaces(),
		"alicloud_cr_repos":                                    dataSourceAlicloudCRRepos(),
		"alicloud_cr_ee_instances":                             dataSourceAlicloudCrEEInstances(),
		"alicloud_cr_ee_namespaces":                            dataSourceAlicloudCrEENamespaces(),
		"alicloud_cr_ee_repos":                                 dataSourceAlicloudCrEERepos(),
		"alicloud_cr_ee_sync_rules":                            dataSourceAlicloudCrEESyncRules(),
		"alicloud_mns_queues":                                  dataSourceAlicloudMNSQueues(),
		"alicloud_mns_topics":                                  dataSourceAlicloudMNSTopics(),
		"alicloud_mns_topic_subscriptions":                     dataSourceAlicloudMNSTopicSubscriptions(),
		"alicloud_api_gateway_service":                         dataSourceAlicloudApiGatewayService(),
		"alicloud_api_gateway_apis":                            dataSourceAlicloudApiGatewayApis(),
		"alicloud_api_gateway_groups":                          dataSourceAlicloudApiGatewayGroups(),
		"alicloud_api_gateway_apps":                            dataSourceAlicloudApiGatewayApps(),
		"alicloud_elasticsearch_instances":                     dataSourceAlicloudElasticsearch(),
		"alicloud_elasticsearch_zones":                         dataSourceAlicloudElaticsearchZones(),
		"alicloud_drds_instances":                              dataSourceAlicloudDRDSInstances(),
		"alicloud_nas_service":                                 dataSourceAlicloudNasService(),
		"alicloud_nas_access_groups":                           dataSourceAlicloudNasAccessGroups(),
		"alicloud_nas_access_rules":                            dataSourceAlicloudAccessRules(),
		"alicloud_nas_mount_targets":                           dataSourceAlicloudNasMountTargets(),
		"alicloud_nas_file_systems":                            dataSourceAlicloudFileSystems(),
		"alicloud_nas_protocols":                               dataSourceAlicloudNasProtocols(),
		"alicloud_cas_certificates":                            dataSourceAlicloudSslCertificatesServiceCertificates(),
		"alicloud_common_bandwidth_packages":                   dataSourceAlicloudCommonBandwidthPackages(),
		"alicloud_route_tables":                                dataSourceAlicloudRouteTables(),
		"alicloud_route_entries":                               dataSourceAlicloudRouteEntries(),
		"alicloud_nat_gateways":                                dataSourceAlicloudNatGateways(),
		"alicloud_snat_entries":                                dataSourceAlicloudSnatEntries(),
		"alicloud_forward_entries":                             dataSourceAlicloudForwardEntries(),
		"alicloud_ddoscoo_instances":                           dataSourceAlicloudDdoscooInstances(),
		"alicloud_ddosbgp_instances":                           dataSourceAlicloudDdosbgpInstances(),
		"alicloud_ess_alarms":                                  dataSourceAlicloudEssAlarms(),
		"alicloud_ess_notifications":                           dataSourceAlicloudEssNotifications(),
		"alicloud_ess_scaling_groups":                          dataSourceAlicloudEssScalingGroups(),
		"alicloud_ess_scaling_rules":                           dataSourceAlicloudEssScalingRules(),
		"alicloud_ess_scaling_configurations":                  dataSourceAlicloudEssScalingConfigurations(),
		"alicloud_ess_lifecycle_hooks":                         dataSourceAlicloudEssLifecycleHooks(),
		"alicloud_ess_scheduled_tasks":                         dataSourceAlicloudEssScheduledTasks(),
		"alicloud_ots_service":                                 dataSourceAlicloudOtsService(),
		"alicloud_ots_instances":                               dataSourceAlicloudOtsInstances(),
		"alicloud_ots_instance_attachments":                    dataSourceAlicloudOtsInstanceAttachments(),
		"alicloud_ots_tables":                                  dataSourceAlicloudOtsTables(),
		"alicloud_ots_tunnels":                                 dataSourceAlicloudOtsTunnels(),
		"alicloud_cloud_connect_networks":                      dataSourceAlicloudCloudConnectNetworks(),
		"alicloud_emr_instance_types":                          dataSourceAlicloudEmrInstanceTypes(),
		"alicloud_emr_disk_types":                              dataSourceAlicloudEmrDiskTypes(),
		"alicloud_emr_main_versions":                           dataSourceAlicloudEmrMainVersions(),
		"alicloud_sag_acls":                                    dataSourceAlicloudSagAcls(),
		"alicloud_yundun_dbaudit_instance":                     dataSourceAlicloudDbauditInstances(),
		"alicloud_yundun_bastionhost_instances":                dataSourceAlicloudBastionhostInstances(),
		"alicloud_bastionhost_instances":                       dataSourceAlicloudBastionhostInstances(),
		"alicloud_market_product":                              dataSourceAlicloudProduct(),
		"alicloud_market_products":                             dataSourceAlicloudProducts(),
		"alicloud_polardb_clusters":                            dataSourceAlicloudPolarDBClusters(),
		"alicloud_polardb_node_classes":                        dataSourceAlicloudPolarDBNodeClasses(),
		"alicloud_polardb_endpoints":                           dataSourceAlicloudPolarDBEndpoints(),
		"alicloud_polardb_accounts":                            dataSourceAlicloudPolarDBAccounts(),
		"alicloud_polardb_databases":                           dataSourceAlicloudPolarDBDatabases(),
		"alicloud_polardb_zones":                               dataSourceAlicloudPolarDBZones(),
		"alicloud_hbase_instances":                             dataSourceAlicloudHBaseInstances(),
		"alicloud_hbase_zones":                                 dataSourceAlicloudHBaseZones(),
		"alicloud_hbase_instance_types":                        dataSourceAlicloudHBaseInstanceTypes(),
		"alicloud_adb_clusters":                                dataSourceAlicloudAdbDbClusters(),
		"alicloud_adb_zones":                                   dataSourceAlicloudAdbZones(),
		"alicloud_cen_flowlogs":                                dataSourceAlicloudCenFlowlogs(),
		"alicloud_kms_aliases":                                 dataSourceAlicloudKmsAliases(),
		"alicloud_dns_domain_txt_guid":                         dataSourceAlicloudDnsDomainTxtGuid(),
		"alicloud_edas_service":                                dataSourceAlicloudEdasService(),
		"alicloud_fnf_service":                                 dataSourceAlicloudFnfService(),
		"alicloud_kms_service":                                 dataSourceAlicloudKmsService(),
		"alicloud_sae_service":                                 dataSourceAlicloudSaeService(),
		"alicloud_dataworks_service":                           dataSourceAlicloudDataWorksService(),
		"alicloud_data_works_service":                          dataSourceAlicloudDataWorksService(),
		"alicloud_mns_service":                                 dataSourceAlicloudMnsService(),
		"alicloud_cloud_storage_gateway_service":               dataSourceAlicloudCloudStorageGatewayService(),
		"alicloud_vs_service":                                  dataSourceAlicloudVsService(),
		"alicloud_pvtz_service":                                dataSourceAlicloudPvtzService(),
		"alicloud_cms_service":                                 dataSourceAlicloudCmsService(),
		"alicloud_maxcompute_service":                          dataSourceAlicloudMaxcomputeService(),
		"alicloud_brain_industrial_service":                    dataSourceAlicloudBrainIndustrialService(),
		"alicloud_iot_service":                                 dataSourceAlicloudIotService(),
		"alicloud_ack_service":                                 dataSourceAlicloudAckService(),
		"alicloud_cr_service":                                  dataSourceAlicloudCrService(),
		"alicloud_dcdn_service":                                dataSourceAlicloudDcdnService(),
		"alicloud_datahub_service":                             dataSourceAlicloudDatahubService(),
		"alicloud_ons_service":                                 dataSourceAlicloudOnsService(),
		"alicloud_fc_service":                                  dataSourceAlicloudFcService(),
		"alicloud_privatelink_service":                         dataSourceAlicloudPrivateLinkService(),
		"alicloud_edas_applications":                           dataSourceAlicloudEdasApplications(),
		"alicloud_edas_deploy_groups":                          dataSourceAlicloudEdasDeployGroups(),
		"alicloud_edas_clusters":                               dataSourceAlicloudEdasClusters(),
		"alicloud_resource_manager_folders":                    dataSourceAlicloudResourceManagerFolders(),
		"alicloud_dns_instances":                               dataSourceAlicloudAlidnsInstances(),
		"alicloud_resource_manager_policies":                   dataSourceAlicloudResourceManagerPolicies(),
		"alicloud_resource_manager_resource_groups":            dataSourceAlicloudResourceManagerResourceGroups(),
		"alicloud_resource_manager_roles":                      dataSourceAlicloudResourceManagerRoles(),
		"alicloud_resource_manager_policy_versions":            dataSourceAlicloudResourceManagerPolicyVersions(),
		"alicloud_alidns_domain_groups":                        dataSourceAlicloudAlidnsDomainGroups(),
		"alicloud_kms_key_versions":                            dataSourceAlicloudKmsKeyVersions(),
		"alicloud_alidns_records":                              dataSourceAlicloudAlidnsRecords(),
		"alicloud_resource_manager_accounts":                   dataSourceAlicloudResourceManagerAccounts(),
		"alicloud_resource_manager_resource_directories":       dataSourceAlicloudResourceManagerResourceDirectories(),
		"alicloud_resource_manager_handshakes":                 dataSourceAlicloudResourceManagerHandshakes(),
		"alicloud_waf_domains":                                 dataSourceAlicloudWafDomains(),
		"alicloud_kms_secrets":                                 dataSourceAlicloudKmsSecrets(),
		"alicloud_cen_route_maps":                              dataSourceAlicloudCenRouteMaps(),
		"alicloud_cen_private_zones":                           dataSourceAlicloudCenPrivateZones(),
		"alicloud_dms_enterprise_instances":                    dataSourceAlicloudDmsEnterpriseInstances(),
		"alicloud_cassandra_clusters":                          dataSourceAlicloudCassandraClusters(),
		"alicloud_cassandra_data_centers":                      dataSourceAlicloudCassandraDataCenters(),
		"alicloud_cassandra_zones":                             dataSourceAlicloudCassandraZones(),
		"alicloud_kms_secret_versions":                         dataSourceAlicloudKmsSecretVersions(),
		"alicloud_waf_instances":                               dataSourceAlicloudWafInstances(),
		"alicloud_eci_image_caches":                            dataSourceAlicloudEciImageCaches(),
		"alicloud_dms_enterprise_users":                        dataSourceAlicloudDmsEnterpriseUsers(),
		"alicloud_dms_user_tenants":                            dataSourceAlicloudDmsUserTenants(),
		"alicloud_ecs_dedicated_hosts":                         dataSourceAlicloudEcsDedicatedHosts(),
		"alicloud_oos_templates":                               dataSourceAlicloudOosTemplates(),
		"alicloud_oos_executions":                              dataSourceAlicloudOosExecutions(),
		"alicloud_resource_manager_policy_attachments":         dataSourceAlicloudResourceManagerPolicyAttachments(),
		"alicloud_dcdn_domains":                                dataSourceAlicloudDcdnDomains(),
		"alicloud_mse_clusters":                                dataSourceAlicloudMseClusters(),
		"alicloud_actiontrail_trails":                          dataSourceAlicloudActiontrailTrails(),
		"alicloud_actiontrails":                                dataSourceAlicloudActiontrailTrails(),
		"alicloud_alidns_instances":                            dataSourceAlicloudAlidnsInstances(),
		"alicloud_alidns_domains":                              dataSourceAlicloudAlidnsDomains(),
		"alicloud_log_alert_resource":                          dataSourceAlicloudLogAlertResource(),
		"alicloud_log_service":                                 dataSourceAlicloudLogService(),
		"alicloud_cen_instance_attachments":                    dataSourceAlicloudCenInstanceAttachments(),
		"alicloud_cdn_service":                                 dataSourceAlicloudCdnService(),
		"alicloud_cen_vbr_health_checks":                       dataSourceAlicloudCenVbrHealthChecks(),
		"alicloud_config_rules":                                dataSourceAlicloudConfigRules(),
		"alicloud_config_configuration_recorders":              dataSourceAlicloudConfigConfigurationRecorders(),
		"alicloud_config_delivery_channels":                    dataSourceAlicloudConfigDeliveryChannels(),
		"alicloud_cms_alarm_contacts":                          dataSourceAlicloudCmsAlarmContacts(),
		"alicloud_kvstore_connections":                         dataSourceAlicloudKvstoreConnections(),
		"alicloud_cms_alarm_contact_groups":                    dataSourceAlicloudCmsAlarmContactGroups(),
		"alicloud_enhanced_nat_available_zones":                dataSourceAlicloudEnhancedNatAvailableZones(),
		"alicloud_cen_route_services":                          dataSourceAlicloudCenRouteServices(),
		"alicloud_kvstore_accounts":                            dataSourceAlicloudKvstoreAccounts(),
		"alicloud_cms_group_metric_rules":                      dataSourceAlicloudCmsGroupMetricRules(),
		"alicloud_fnf_flows":                                   dataSourceAlicloudFnfFlows(),
		"alicloud_fnf_schedules":                               dataSourceAlicloudFnfSchedules(),
		"alicloud_ros_change_sets":                             dataSourceAlicloudRosChangeSets(),
		"alicloud_ros_stacks":                                  dataSourceAlicloudRosStacks(),
		"alicloud_ros_stack_groups":                            dataSourceAlicloudRosStackGroups(),
		"alicloud_ros_templates":                               dataSourceAlicloudRosTemplates(),
		"alicloud_privatelink_vpc_endpoint_services":           dataSourceAlicloudPrivatelinkVpcEndpointServices(),
		"alicloud_privatelink_vpc_endpoints":                   dataSourceAlicloudPrivatelinkVpcEndpoints(),
		"alicloud_privatelink_vpc_endpoint_connections":        dataSourceAlicloudPrivatelinkVpcEndpointConnections(),
		"alicloud_privatelink_vpc_endpoint_service_resources":  dataSourceAlicloudPrivatelinkVpcEndpointServiceResources(),
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
		"alicloud_quotas_quotas":                               dataSourceAlicloudQuotasQuotas(),
		"alicloud_quotas_application_infos":                    dataSourceAlicloudQuotasQuotaApplications(),
		"alicloud_cms_monitor_group_instanceses":               dataSourceAlicloudCmsMonitorGroupInstances(),
		"alicloud_cms_monitor_group_instances":                 dataSourceAlicloudCmsMonitorGroupInstances(),
		"alicloud_quotas_quota_alarms":                         dataSourceAlicloudQuotasQuotaAlarms(),
		"alicloud_ecs_commands":                                dataSourceAlicloudEcsCommands(),
		"alicloud_cloud_storage_gateway_storage_bundles":       dataSourceAlicloudCloudStorageGatewayStorageBundles(),
		"alicloud_ecs_hpc_clusters":                            dataSourceAlicloudEcsHpcClusters(),
		"alicloud_brain_industrial_pid_loops":                  dataSourceAlicloudBrainIndustrialPidLoops(),
		"alicloud_quotas_quota_applications":                   dataSourceAlicloudQuotasQuotaApplications(),
		"alicloud_ecs_auto_snapshot_policies":                  dataSourceAlicloudEcsAutoSnapshotPolicies(),
		"alicloud_rds_parameter_groups":                        dataSourceAlicloudRdsParameterGroups(),
		"alicloud_ecs_launch_templates":                        dataSourceAlicloudEcsLaunchTemplates(),
		"alicloud_resource_manager_control_policies":           dataSourceAlicloudResourceManagerControlPolicies(),
		"alicloud_resource_manager_control_policy_attachments": dataSourceAlicloudResourceManagerControlPolicyAttachments(),
		"alicloud_rds_accounts":                                dataSourceAlicloudRdsAccounts(),
		"alicloud_havips":                                      dataSourceAlicloudHavips(),
		"alicloud_ecs_snapshots":                               dataSourceAlicloudEcsSnapshots(),
		"alicloud_ecs_key_pairs":                               dataSourceAlicloudEcsKeyPairs(),
		"alicloud_adb_db_clusters":                             dataSourceAlicloudAdbDbClusters(),
		"alicloud_vpc_flow_logs":                               dataSourceAlicloudVpcFlowLogs(),
		"alicloud_network_acls":                                dataSourceAlicloudNetworkAcls(),
		"alicloud_ecs_disks":                                   dataSourceAlicloudEcsDisks(),
		"alicloud_ddoscoo_domain_resources":                    dataSourceAlicloudDdoscooDomainResources(),
		"alicloud_ddoscoo_ports":                               dataSourceAlicloudDdoscooPorts(),
		"alicloud_slb_load_balancers":                          dataSourceAlicloudSlbLoadBalancers(),
		"alicloud_ecs_network_interfaces":                      dataSourceAlicloudEcsNetworkInterfaces(),
		"alicloud_config_aggregators":                          dataSourceAlicloudConfigAggregators(),
		"alicloud_config_aggregate_config_rules":               dataSourceAlicloudConfigAggregateConfigRules(),
		"alicloud_config_aggregate_compliance_packs":           dataSourceAlicloudConfigAggregateCompliancePacks(),
		"alicloud_config_compliance_packs":                     dataSourceAlicloudConfigCompliancePacks(),
		"alicloud_eip_addresses":                               dataSourceAlicloudEipAddresses(),
		"alicloud_direct_mail_receiverses":                     dataSourceAlicloudDirectMailReceiverses(),
		"alicloud_log_projects":                                dataSourceAlicloudLogProjects(),
		"alicloud_log_stores":                                  dataSourceAlicloudLogStores(),
		"alicloud_event_bridge_service":                        dataSourceAlicloudEventBridgeService(),
		"alicloud_event_bridge_event_buses":                    dataSourceAlicloudEventBridgeEventBuses(),
		"alicloud_amqp_virtual_hosts":                          dataSourceAlicloudAmqpVirtualHosts(),
		"alicloud_amqp_queues":                                 dataSourceAlicloudAmqpQueues(),
		"alicloud_amqp_exchanges":                              dataSourceAlicloudAmqpExchanges(),
		"alicloud_cassandra_backup_plans":                      dataSourceAlicloudCassandraBackupPlans(),
		"alicloud_cen_transit_router_peer_attachments":         dataSourceAlicloudCenTransitRouterPeerAttachments(),
		"alicloud_amqp_instances":                              dataSourceAlicloudAmqpInstances(),
		"alicloud_hbr_vaults":                                  dataSourceAlicloudHbrVaults(),
		"alicloud_ssl_certificates_service_certificates":       dataSourceAlicloudSslCertificatesServiceCertificates(),
		"alicloud_arms_alert_contacts":                         dataSourceAlicloudArmsAlertContacts(),
		"alicloud_event_bridge_rules":                          dataSourceAlicloudEventBridgeRules(),
		"alicloud_cloud_firewall_control_policies":             dataSourceAlicloudCloudFirewallControlPolicies(),
		"alicloud_sae_namespaces":                              dataSourceAlicloudSaeNamespaces(),
		"alicloud_sae_config_maps":                             dataSourceAlicloudSaeConfigMaps(),
		"alicloud_alb_security_policies":                       dataSourceAlicloudAlbSecurityPolicies(),
		"alicloud_event_bridge_event_sources":                  dataSourceAlicloudEventBridgeEventSources(),
		"alicloud_ecd_policy_groups":                           dataSourceAlicloudEcdPolicyGroups(),
		"alicloud_ecp_key_pairs":                               dataSourceAlicloudEcpKeyPairs(),
		"alicloud_hbr_ecs_backup_plans":                        dataSourceAlicloudHbrEcsBackupPlans(),
		"alicloud_hbr_nas_backup_plans":                        dataSourceAlicloudHbrNasBackupPlans(),
		"alicloud_hbr_oss_backup_plans":                        dataSourceAlicloudHbrOssBackupPlans(),
		"alicloud_scdn_domains":                                dataSourceAlicloudScdnDomains(),
		"alicloud_alb_server_groups":                           dataSourceAlicloudAlbServerGroups(),
		"alicloud_data_works_folders":                          dataSourceAlicloudDataWorksFolders(),
		"alicloud_arms_alert_contact_groups":                   dataSourceAlicloudArmsAlertContactGroups(),
		"alicloud_express_connect_access_points":               dataSourceAlicloudExpressConnectAccessPoints(),
		"alicloud_cloud_storage_gateway_gateways":              dataSourceAlicloudCloudStorageGatewayGateways(),
		"alicloud_lindorm_instances":                           dataSourceAlicloudLindormInstances(),
		"alicloud_express_connect_physical_connection_service": dataSourceAlicloudExpressConnectPhysicalConnectionService(),
		"alicloud_cddc_dedicated_host_groups":                  dataSourceAlicloudCddcDedicatedHostGroups(),
		"alicloud_hbr_ecs_backup_clients":                      dataSourceAlicloudHbrEcsBackupClients(),
		"alicloud_msc_sub_contacts":                            dataSourceAlicloudMscSubContacts(),
		"alicloud_express_connect_physical_connections":        dataSourceAlicloudExpressConnectPhysicalConnections(),
		"alicloud_alb_load_balancers":                          dataSourceAlicloudAlbLoadBalancers(),
		"alicloud_alb_zones":                                   dataSourceAlicloudAlbZones(),
		"alicloud_sddp_rules":                                  dataSourceAlicloudSddpRules(),
		"alicloud_bastionhost_user_groups":                     dataSourceAlicloudBastionhostUserGroups(),
		"alicloud_security_center_groups":                      dataSourceAlicloudSecurityCenterGroups(),
		"alicloud_alb_acls":                                    dataSourceAlicloudAlbAcls(),
		"alicloud_hbr_snapshots":                               dataSourceAlicloudHbrSnapshots(),
		"alicloud_bastionhost_users":                           dataSourceAlicloudBastionhostUsers(),
		"alicloud_dfs_access_groups":                           dataSourceAlicloudDfsAccessGroups(),
		"alicloud_ehpc_job_templates":                          dataSourceAlicloudEhpcJobTemplates(),
		"alicloud_sddp_configs":                                dataSourceAlicloudSddpConfigs(),
		"alicloud_hbr_restore_jobs":                            dataSourceAlicloudHbrRestoreJobs(),
		"alicloud_alb_listeners":                               dataSourceAlicloudAlbListeners(),
		"alicloud_ens_key_pairs":                               dataSourceAlicloudEnsKeyPairs(),
		"alicloud_sae_applications":                            dataSourceAlicloudSaeApplications(),
		"alicloud_alb_rules":                                   dataSourceAlicloudAlbRules(),
		"alicloud_cms_metric_rule_templates":                   dataSourceAlicloudCmsMetricRuleTemplates(),
		"alicloud_iot_device_groups":                           dataSourceAlicloudIotDeviceGroups(),
		"alicloud_express_connect_virtual_border_routers":      dataSourceAlicloudExpressConnectVirtualBorderRouters(),
		"alicloud_imm_projects":                                dataSourceAlicloudImmProjects(),
		"alicloud_click_house_db_clusters":                     dataSourceAlicloudClickHouseDbClusters(),
		"alicloud_direct_mail_domains":                         dataSourceAlicloudDirectMailDomains(),
		"alicloud_bastionhost_host_groups":                     dataSourceAlicloudBastionhostHostGroups(),
		"alicloud_vpc_dhcp_options_sets":                       dataSourceAlicloudVpcDhcpOptionsSets(),
		"alicloud_alb_health_check_templates":                  dataSourceAlicloudAlbHealthCheckTemplates(),
		"alicloud_cdn_real_time_log_deliveries":                dataSourceAlicloudCdnRealTimeLogDeliveries(),
		"alicloud_click_house_accounts":                        dataSourceAlicloudClickHouseAccounts(),
		"alicloud_direct_mail_mail_addresses":                  dataSourceAlicloudDirectMailMailAddresses(),
		"alicloud_database_gateway_gateways":                   dataSourceAlicloudDatabaseGatewayGateways(),
		"alicloud_bastionhost_hosts":                           dataSourceAlicloudBastionhostHosts(),
		"alicloud_amqp_bindings":                               dataSourceAlicloudAmqpBindings(),
		"alicloud_slb_tls_cipher_policies":                     dataSourceAlicloudSlbTlsCipherPolicies(),
		"alicloud_cloud_sso_directories":                       dataSourceAlicloudCloudSsoDirectories(),
		"alicloud_bastionhost_host_accounts":                   dataSourceAlicloudBastionhostHostAccounts(),
		"alicloud_waf_certificates":                            dataSourceAlicloudWafCertificates(),
		"alicloud_simple_application_server_instances":         dataSourceAlicloudSimpleApplicationServerInstances(),
		"alicloud_simple_application_server_plans":             dataSourceAlicloudSimpleApplicationServerPlans(),
		"alicloud_simple_application_server_images":            dataSourceAlicloudSimpleApplicationServerImages(),
		"alicloud_video_surveillance_system_groups":            dataSourceAlicloudVideoSurveillanceSystemGroups(),
		"alicloud_msc_sub_subscriptions":                       dataSourceAlicloudMscSubSubscriptions(),
		"alicloud_sddp_instances":                              dataSourceAlicloudSddpInstances(),
		"alicloud_vpc_nat_ip_cidrs":                            dataSourceAlicloudVpcNatIpCidrs(),
		"alicloud_vpc_nat_ips":                                 dataSourceAlicloudVpcNatIps(),
		"alicloud_quick_bi_users":                              dataSourceAlicloudQuickBiUsers(),
		"alicloud_vod_domains":                                 dataSourceAlicloudVodDomains(),
		"alicloud_arms_dispatch_rules":                         dataSourceAlicloudArmsDispatchRules(),
		"alicloud_open_search_app_groups":                      dataSourceAlicloudOpenSearchAppGroups(),
		"alicloud_graph_database_db_instances":                 dataSourceAlicloudGraphDatabaseDbInstances(),
		"alicloud_arms_prometheus_alert_rules":                 dataSourceAlicloudArmsPrometheusAlertRules(),
		"alicloud_dbfs_instances":                              dataSourceAlicloudDbfsInstances(),
		"alicloud_rdc_organizations":                           dataSourceAlicloudRdcOrganizations(),
		"alicloud_eais_instances":                              dataSourceAlicloudEaisInstances(),
		"alicloud_sae_ingresses":                               dataSourceAlicloudSaeIngresses(),
		"alicloud_cloudauth_face_configs":                      dataSourceAlicloudCloudauthFaceConfigs(),
		"alicloud_imp_app_templates":                           dataSourceAlicloudImpAppTemplates(),
		"alicloud_mhub_products":                               dataSourceAlicloudMhubProducts(),
		"alicloud_cloud_sso_scim_server_credentials":           dataSourceAlicloudCloudSsoScimServerCredentials(),
		"alicloud_dts_subscription_jobs":                       dataSourceAlicloudDtsSubscriptionJobs(),
		"alicloud_service_mesh_service_meshes":                 dataSourceAlicloudServiceMeshServiceMeshes(),
		"alicloud_service_mesh_versions":                       dataSourceAlicloudServiceMeshVersions(),
		"alicloud_mhub_apps":                                   dataSourceAlicloudMhubApps(),
		"alicloud_cloud_sso_groups":                            dataSourceAlicloudCloudSsoGroups(),
		"alicloud_hbr_backup_jobs":                             dataSourceAlicloudHbrBackupJobs(),
		"alicloud_click_house_regions":                         dataSourceAlicloudClickHouseRegions(),
		"alicloud_dts_synchronization_jobs":                    dataSourceAlicloudDtsSynchronizationJobs(),
		"alicloud_cloud_firewall_instances":                    dataSourceAlicloudCloudFirewallInstances(),
		"alicloud_cr_endpoint_acl_policies":                    dataSourceAlicloudCrEndpointAclPolicies(),
		"alicloud_cr_endpoint_acl_service":                     dataSourceAlicloudCrEndpointAclService(),
		"alicloud_actiontrail_history_delivery_jobs":           dataSourceAlicloudActiontrailHistoryDeliveryJobs(),
		"alicloud_sae_instance_specifications":                 dataSourceAlicloudSaeInstanceSpecifications(),
		"alicloud_cen_transit_router_service":                  dataSourceAlicloudCenTransitRouterService(),
		"alicloud_ecs_deployment_sets":                         dataSourceAlicloudEcsDeploymentSets(),
		"alicloud_cloud_sso_users":                             dataSourceAlicloudCloudSsoUsers(),
		"alicloud_cloud_sso_access_configurations":             
		"alicloud_dfs_file_systems":                          
		"alicloud_dfs_zones":                                 
		"alicloud_vpc_traffic_mirror_filters":                 
		"alicloud_dfs_access_rules":                         
		"alicloud_nas_zones":                                  
		"alicloud_dfs_mount_points":                           
		"alicloud_vpc_traffic_mirror_filter_egress_rules":     
		"alicloud_ecd_simple_office_sites":                    
		"alicloud_vpc_traffic_mirror_filter_ingress_rules":    
		"alicloud_ecd_nas_file_systems":                       
		"alicloud_vpc_traffic_mirror_service":                
		"alicloud_msc_sub_webhooks":                         
		"alicloud_ecd_users":                                 
		"alicloud_vpc_traffic_mirror_sessions":              
		"alicloud_gpdb_accounts":                             
		"alicloud_vpc_ipv6_gateways":                      
		"alicloud_vpc_ipv6_egress_rules":                     
		"alicloud_vpc_ipv6_addresses":                        
		"alicloud_hbr_server_backup_plans":                 
		"alicloud_cms_dynamic_tag_groups":                  
		"alicloud_ecd_network_packages":                     
		"alicloud_cloud_storage_gateway_gateway_smb_users":    
		"alicloud_vpc_ipv6_internet_bandwidths":                
		"alicloud_simple_application_server_firewall_rules":   
		"alicloud_pvtz_endpoints":                             
		"alicloud_pvtz_resolver_zones":                      
		"alicloud_pvtz_rules":                                
		"alicloud_ecd_bundles":                               
		"alicloud_simple_application_server_disks":            
		"alicloud_simple_application_server_snapshots":       
		"alicloud_simple_application_server_custom_images":     
		"alicloud_cloud_storage_gateway_stocks":                
		"alicloud_cloud_storage_gateway_gateway_cache_disks":  
		"alicloud_cloud_storage_gateway_gateway_block_volumes":
		"alicloud_direct_mail_tags":                           
		"alicloud_cloud_storage_gateway_gateway_file_shares": 
		"alicloud_ecd_desktops":                               
		"alicloud_cloud_storage_gateway_express_syncs":        
		"alicloud_oos_applications":                           
		"alicloud_eci_virtual_nodes":                         
		"alicloud_eci_zones":                                
		"alicloud_ros_stack_instances":                     
		"alicloud_ros_regions":                             
		"alicloud_ecs_dedicated_host_clusters":            
		"alicloud_oos_application_groups":                  
		"alicloud_dts_consumer_channels":                
		"alicloud_emr_clusters":                              
		"alicloud_ecd_images":                                
		"alicloud_oos_patch_baselines":                       
		"alicloud_ecd_commands":                              
		"alicloud_cddc_zones":                               
		"alicloud_cddc_host_ecs_level_infos":             
		"alicloud_cddc_dedicated_hosts":                       
		"alicloud_oos_parameters":                             
		"alicloud_oos_state_configurations":                 
		"alicloud_oos_secret_parameters":                     
		"alicloud_click_house_backup_policies":              
		"alicloud_cloud_sso_service":                          
		"alicloud_mongodb_audit_policies":                  
		"alicloud_mongodb_accounts":                           
		"alicloud_mongodb_serverless_instances":               
		"alicloud_cddc_dedicated_host_accounts":              
		"alicloud_cr_chart_namespaces":                        
		"alicloud_fnf_executions":                        
		"alicloud_cr_chart_repositories":                   
		"alicloud_mongodb_sharding_network_public_addresses":   
		"alicloud_ga_acls":                                     
		"alicloud_ga_additional_certificates":                
		"alicloud_alidns_custom_lines":                  
		"alicloud_ros_template_scratches":                  
		"alicloud_alidns_gtm_instances":                   
		"alicloud_vpc_bgp_groups":                         
		"alicloud_nas_snapshots":                           
		"alicloud_hbr_replication_vault_regions":          
		"alicloud_alidns_address_pools":                     
		"alicloud_ecs_prefix_lists":                          
		"alicloud_alidns_access_strategies":                   
		"alicloud_vpc_bgp_peers":                             
		"alicloud_nas_filesets":                            
		"cdn_ip_info":                                     
		"nas_auto_snapshot_policies":                       
		"nas_lifecycle_policies":                        
		"vpc_bgp_networks":                              
		"nas_data_flows":                                       
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
		"ecs_image_components":                               
		"sae_application_scaling_rules":                      
		"sae_grey_tag_routes":                               
		"ecs_snapshot_groups":                                 
		"vpn_ipsec_servers":                                   
		"cr_chains":                                         
		"vpn_pbr_route_entries":                               
		"mse_znodes":                                         
		"cen_transit_router_available_resources":      
		"ecs_image_pipelines":                             
		"hbr_ots_backup_plans":                                
		"hbr_ots_snapshots":                                   
		"bastionhost_host_share_keys":                   
		"ecs_network_interface_permissions":                 
		"mse_engine_namespaces":                              
		"ga_accelerator_spare_ip_attachments":              
		"smartag_flow_logs":                                    
		"ecs_invocations":                                     
		"ecd_snapshots":                                       
		"tag_meta_tags":                                      
		"ecd_desktop_types":                                   
		"config_deliveries":                                   
		"cms_namespaces":                                     
		"cms_sls_groups":                                     
		"config_aggregate_deliveries":                        
		"edas_namespaces":                                  
		"cdn_blocked_regions":                                 
		"schedulerx_namespaces":                                "schedulerx_namespaces",
		"ehpc_clusters": "ehpc_clusters", 
		"cen_traffic_marking_policies": "cen_traffic_marking_policies",
	}

}

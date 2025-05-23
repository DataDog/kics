package generic.terraform

import data.generic.common as common_lib

check_cidr(rule) {
	rule.cidr_blocks[_] == "0.0.0.0/0"
} else {
	rule.cidr_block == "0.0.0.0/0"
}

# Checks if a TCP port is open in a rule
portOpenToInternet(rule, port) {
	check_cidr(rule)
	rule.protocol == "tcp"
	containsPort(rule, port)
}

portOpenToInternet(rules, port) {
	rule := rules[_]
	check_cidr(rule)
	rule.protocol == "tcp"
	containsPort(rule, port)
}

# Checks if a port is included in a rule
containsPort(rule, port) {
	rule.from_port <= port
	rule.to_port >= port
} else {
	rule.from_port == 0
	rule.to_port == 0
} else {
	regex.match(sprintf("(^|\\s|,)%d(-|,|$|\\s)", [port]), rule.destination_port_range)
} else {
	ports := split(rule.destination_port_range, ",")
	sublist := split(ports[var], "-")
	to_number(trim(sublist[0], " ")) <= port
	to_number(trim(sublist[1], " ")) >= port
} else {
	ports := split(rule.port_range, ",")
	sublist := split(ports[var], "/")
	to_number(trim(sublist[0], " ")) <= port
	to_number(trim(sublist[1], " ")) >= port
}

# Gets the list of protocols
getProtocolList("-1") = protocols {
	protocols := ["TCP", "UDP", "ICMP"]
}

getProtocolList("*") = protocols {
	protocols := ["TCP", "UDP", "ICMP"]
}

getProtocolList(protocol) = protocols {
	upper(protocol) == "TCP"
	protocols := ["TCP"]
}

getProtocolList(protocol) = protocols {
	upper(protocol) == "UDP"
	protocols := ["UDP"]
}

# Checks if any principal are allowed in a policy
anyPrincipal(statement) {
  # Case: Principal = "*"
  statement.Principal == "*"
}

anyPrincipal(statement) {
  # Case: Principal = {"AWS": "*"}
  is_string(statement.Principal.AWS)
  statement.Principal.AWS == "*"
}

anyPrincipal(statement) {
  # Case: Principal = {"AWS": ["*"]}
  is_array(statement.Principal.AWS)
  "*" == statement.Principal.AWS[_]
}

anyPrincipal(statement) {
  # Case: Principal = {"Service": "*"}  (uncommon, but valid)
  is_string(statement.Principal.Service)
  statement.Principal.Service == "*"
}

anyPrincipal(statement) {
  # Case: Principal = {"Service": ["*"]}
  is_array(statement.Principal.Service)
  "*" == statement.Principal.Service[_]
}

getSpecInfo(resource) = specInfo { # this one can be also used for the result
	spec := resource.spec.job_template.spec.template.spec
	specInfo := {"spec": spec, "path": "spec.job_template.spec.template.spec"}
} else = specInfo {
	spec := resource.spec.template.spec
	specInfo := {"spec": spec, "path": "spec.template.spec"}
} else = specInfo {
	spec := resource.spec
	specInfo := {"spec": spec, "path": "spec"}
}

check_aws_resource_supports_tags(p) {
	resource = {
		"aws_acm_certificate": "",
		"aws_acmpca_certificate_authority": "",
		"aws_api_gateway_api_key": "",
		"aws_api_gateway_client_certificate": "",
		"aws_api_gateway_domain_name": "",
		"aws_api_gateway_rest_api": "",
		"aws_api_gateway_stage": "",
		"aws_api_gateway_usage_plan": "",
		"aws_api_gateway_vpc_link": "",
		"aws_apigatewayv2_api": "",
		"aws_apigatewayv2_domain_name": "",
		"aws_apigatewayv2_stage": "",
		"aws_apigatewayv2_vpc_link": "",
		"aws_accessanalyzer_analyzer": "",
		"aws_appmesh_gateway_route": "",
		"aws_appmesh_mesh": "",
		"aws_appmesh_route": "",
		"aws_appmesh_virtual_gateway": "",
		"aws_appmesh_virtual_node": "",
		"aws_appmesh_virtual_router": "",
		"aws_appmesh_virtual_service": "",
		"aws_appsync_graphql_api": "",
		"aws_athena_workgroup": "",
		"aws_autoscaling_group": "",
		"aws_backup_plan": "",
		"aws_backup_vault": "",
		"aws_batch_compute_environment": "",
		"aws_batch_job_definition": "",
		"aws_batch_job_queue": "",
		"aws_cloud9_environment_ec2": "",
		"aws_cloudformation_stack": "",
		"aws_cloudformation_stack_set": "",
		"aws_cloudfront_distribution": "",
		"aws_cloudhsm_v2_cluster": "",
		"aws_cloudtrail": "",
		"aws_cloudwatch_composite_alarm": "",
		"aws_cloudwatch_log_group": "",
		"aws_cloudwatch_metric_alarm": "",
		"aws_codeartifact_domain": "",
		"aws_codeartifact_repository": "",
		"aws_codebuild_project": "",
		"aws_codebuild_report_group": "",
		"aws_codecommit_repository": "",
		"aws_codedeploy_app": "",
		"aws_codedeploy_deployment_group": "",
		"aws_codepipeline_webhook": "",
		"aws_codestarconnections_connection": "",
		"aws_codestarnotifications_notification_rule": "",
		"aws_cognito_identity_pool": "",
		"aws_cognito_user_pool": "",
		"aws_config_aggregate_authorization": "",
		"aws_config_config_rule": "",
		"aws_config_configuration_aggregator": "",
		"aws_dlm_lifecycle_policy": "",
		"aws_datapipeline_pipeline": "",
		"aws_datasync_agent": "",
		"aws_datasync_location_efs": "",
		"aws_datasync_location_fsx_windows_file_system": "",
		"aws_datasync_location_nfs": "",
		"aws_datasync_location_s3": "",
		"aws_datasync_location_smb": "",
		"aws_datasync_task": "",
		"aws_dms_certificate": "",
		"aws_dms_endpoint": "",
		"aws_dms_event_subscription": "",
		"aws_dms_replication_instance": "",
		"aws_dms_replication_subnet_group": "",
		"aws_dms_replication_task": "",
		"aws_dx_connection": "",
		"aws_dx_hosted_public_virtual_interface": "",
		"aws_dx_hosted_private_virtual_interface_accepter": "",
		"aws_dx_hosted_transit_virtual_interface_accepter": "",
		"aws_dx_lag": "",
		"aws_dx_private_virtual_interface": "",
		"aws_dx_public_virtual_interface": "",
		"aws_dx_transit_virtual_interface": "",
		"aws_directory_service_directory": "",
		"aws_docdb_cluster": "",
		"aws_docdb_cluster_instance": "",
		"aws_docdb_cluster_parameter_group": "",
		"aws_docdb_subnet_group": "",
		"aws_dynamodb_table": "",
		"aws_dax_cluster": "",
		"aws_ami": "",
		"aws_ami_copy": "",
		"aws_ami_from_instance": "",
		"aws_ebs_snapshot": "",
		"aws_ebs_snapshot_copy": "",
		"aws_ebs_volume": "",
		"aws_ec2_capacity_reservation": "",
		"aws_ec2_carrier_gateway": "",
		"aws_ec2_client_vpn_endpoint": "",
		"aws_ec2_fleet": "",
		"aws_ec2_local_gateway_route_table_vpc_association": "",
		"aws_ec2_traffic_mirror_filter": "",
		"aws_ec2_traffic_mirror_session": "",
		"aws_ec2_traffic_mirror_target": "",
		"aws_ec2_transit_gateway": "",
		"aws_ec2_transit_gateway_peering_attachment": "",
		"aws_ec2_transit_gateway_peering_attachment_accepter": "",
		"aws_ec2_transit_gateway_route_table": "",
		"aws_ec2_transit_gateway_vpc_attachment": "",
		"aws_ec2_transit_gateway_vpc_attachment_accepter": "",
		"aws_eip": "",
		"aws_instance": "",
		"aws_key_pair": "",
		"aws_launch_template": "",
		"aws_placement_group": "",
		"aws_spot_fleet_request": "",
		"aws_spot_instance_request": "",
		"aws_ecr_repository": "",
		"aws_ecs_capacity_provider": "",
		"aws_ecs_cluster": "",
		"aws_ecs_service": "",
		"aws_ecs_task_definition": "",
		"aws_efs_access_point": "",
		"aws_efs_file_system": "",
		"aws_eks_addon": "",
		"aws_eks_cluster": "",
		"aws_eks_fargate_profile": "",
		"aws_eks_node_group": "",
		"aws_elasticache_cluster": "",
		"aws_elasticache_replication_group": "",
		"aws_elasticache_subnet_group": "",
		"aws_elastic_beanstalk_application": "",
		"aws_elastic_beanstalk_application_version": "",
		"aws_elastic_beanstalk_environment": "",
		"aws_elb": "",
		"aws_lb": "",
		"aws_lb_target_group": "",
		"aws_emr_cluster": "",
		"aws_elasticsearch_domain": "",
		"aws_cloudwatch_event_bus": "",
		"aws_cloudwatch_event_rule": "",
		"aws_fsx_lustre_file_system": "",
		"aws_fsx_windows_file_system": "",
		"aws_gamelift_alias": "",
		"aws_gamelift_build": "",
		"aws_gamelift_fleet": "",
		"aws_gamelift_game_session_queue": "",
		"aws_glacier_vault": "",
		"aws_globalaccelerator_accelerator": "",
		"aws_glue_crawler": "",
		"aws_glue_dev_endpoint": "",
		"aws_glue_job": "",
		"aws_glue_ml_transform": "",
		"aws_glue_registry": "",
		"aws_glue_schema": "",
		"aws_glue_trigger": "",
		"aws_glue_workflow": "",
		"aws_guardduty_detector": "",
		"aws_guardduty_filter": "",
		"aws_guardduty_ipset": "",
		"aws_guardduty_threatintelset": "",
		"aws_iam_instance_profile": "",
		"aws_iam_openid_connect_provider": "",
		"aws_iam_policy": "",
		"aws_iam_role": "",
		"aws_iam_saml_provider": "",
		"aws_iam_server_certificate": "",
		"aws_iam_user": "",
		"aws_imagebuilder_component": "",
		"aws_imagebuilder_distribution_configuration": "",
		"aws_imagebuilder_image": "",
		"aws_imagebuilder_image_pipeline": "",
		"aws_imagebuilder_image_recipe": "",
		"aws_imagebuilder_infrastructure_configuration": "",
		"aws_inspector_assessment_template": "",
		"aws_iot_topic_rule": "",
		"aws_kms_external_key": "",
		"aws_kms_key": "",
		"aws_kinesis_stream": "",
		"aws_kinesis_analytics_application": "",
		"aws_kinesisanalyticsv2_application": "",
		"aws_kinesis_firehose_delivery_stream": "",
		"aws_kinesis_video_stream": "",
		"aws_lambda_function": "",
		"aws_licensemanager_license_configuration": "",
		"aws_lightsail_instance": "",
		"aws_mq_broker": "",
		"aws_mq_configuration": "",
		"aws_msk_cluster": "",
		"aws_mwaa_environment": "",
		"aws_media_convert_queue": "",
		"aws_media_package_channel": "",
		"aws_media_store_container": "",
		"aws_neptune_cluster": "",
		"aws_neptune_cluster_instance": "",
		"aws_neptune_cluster_parameter_group": "",
		"aws_neptune_event_subscription": "",
		"aws_neptune_parameter_group": "",
		"aws_neptune_subnet_group": "",
		"aws_networkfirewall_firewall": "",
		"aws_networkfirewall_firewall_policy": "",
		"aws_networkfirewall_rule_group": "",
		"aws_opsworks_custom_layer": "",
		"aws_opsworks_ganglia_layer": "",
		"aws_opsworks_haproxy_layer": "",
		"aws_opsworks_java_app_layer": "",
		"aws_opsworks_memcached_layer": "",
		"aws_opsworks_mysql_layer": "",
		"aws_opsworks_nodejs_app_layer": "",
		"aws_opsworks_php_app_layer": "",
		"aws_opsworks_rails_app_layer": "",
		"aws_opsworks_stack": "",
		"aws_opsworks_static_web_layer": "",
		"aws_organizations_account": "",
		"aws_organizations_organizational_unit": "",
		"aws_organizations_policy": "",
		"aws_pinpoint_app": "",
		"aws_qldb_ledger": "",
		"aws_ram_resource_share": "",
		"aws_db_cluster_snapshot": "",
		"aws_db_event_subscription": "",
		"aws_db_instance": "",
		"aws_db_option_group": "",
		"aws_db_parameter_group": "",
		"aws_db_proxy": "",
		"aws_db_proxy_endpoint": "",
		"aws_db_security_group": "",
		"aws_db_subnet_group": "",
		"aws_rds_cluster": "",
		"aws_rds_cluster_endpoint": "",
		"aws_rds_cluster_instance": "",
		"aws_rds_cluster_parameter_group": "",
		"aws_redshift_cluster": "",
		"aws_redshift_event_subscription": "",
		"aws_redshift_parameter_group": "",
		"aws_redshift_snapshot_copy_grant": "",
		"aws_redshift_snapshot_schedule": "",
		"aws_redshift_subnet_group": "",
		"aws_resourcegroups_group": "",
		"aws_route53_health_check": "",
		"aws_route53_zone": "",
		"aws_route53_resolver_endpoint": "",
		"aws_route53_resolver_firewall_domain_list": "",
		"aws_route53_resolver_firewall_rule_group": "",
		"aws_route53_resolver_firewall_rule_group_association": "",
		"aws_route53_resolver_query_log_config": "",
		"aws_route53_resolver_rule": "",
		"aws_s3_bucket": "",
		"aws_s3_bucket_object": "",
		"aws_s3_object_copy": "",
		"aws_s3control_bucket": "",
		"aws_sns_topic": "",
		"aws_sqs_queue": "",
		"aws_ssm_activation": "",
		"aws_ssm_document": "",
		"aws_ssm_maintenance_window": "",
		"aws_ssm_parameter": "",
		"aws_ssm_patch_baseline": "",
		"aws_ssoadmin_permission_set": "",
		"aws_swf_domain": "",
		"aws_sagemaker_app": "",
		"aws_sagemaker_domain": "",
		"aws_sagemaker_endpoint": "",
		"aws_sagemaker_endpoint_configuration": "",
		"aws_sagemaker_feature_group": "",
		"aws_sagemaker_image": "",
		"aws_sagemaker_model": "",
		"aws_sagemaker_model_package_group": "",
		"aws_sagemaker_notebook_instance": "",
		"aws_sagemaker_user_profile": "",
		"aws_secretsmanager_secret": "",
		"aws_serverlessapplicationrepository_cloudformation_stack": "",
		"aws_servicecatalog_portfolio": "",
		"aws_service_discovery_http_namespace": "",
		"aws_service_discovery_private_dns_namespace": "",
		"aws_service_discovery_public_dns_namespace": "",
		"aws_service_discovery_service": "",
		"aws_shield_protection": "",
		"aws_signer_signing_profile": "",
		"aws_sfn_activity": "",
		"aws_sfn_state_machine": "",
		"aws_storagegateway_cached_iscsi_volume": "",
		"aws_storagegateway_gateway": "",
		"aws_storagegateway_nfs_file_share": "",
		"aws_storagegateway_smb_file_share": "",
		"aws_storagegateway_stored_iscsi_volume": "",
		"aws_storagegateway_tape_pool": "",
		"aws_synthetics_canary": "",
		"aws_transfer_server": "",
		"aws_transfer_user": "",
		"aws_customer_gateway": "",
		"aws_default_network_acl": "",
		"aws_default_route_table": "",
		"aws_default_security_group": "",
		"aws_default_subnet": "",
		"aws_ec2_managed_prefix_list": "",
		"aws_egress_only_internet_gateway": "",
		"aws_flow_log": "",
		"aws_internet_gateway": "",
		"aws_nat_gateway": "",
		"aws_network_acl": "",
		"aws_network_interface": "",
		"aws_route_table": "",
		"aws_security_group": "",
		"aws_subnet": "",
		"aws_vpc": "",
		"aws_vpc_dhcp_options": "",
		"aws_vpc_endpoint": "",
		"aws_vpc_endpoint_service": "",
		"aws_vpc_peering_connection": "",
		"aws_vpc_peering_connection_accepter": "",
		"aws_vpn_connection": "",
		"aws_vpn_gateway": "",
		"aws_waf_rate_based_rule": "",
		"aws_waf_rule": "",
		"aws_waf_rule_group": "",
		"aws_waf_web_acl": "",
		"aws_wafregional_rate_based_rule": "",
		"aws_wafregional_rule": "",
		"aws_wafregional_rule_group": "",
		"aws_wafregional_web_acl": "",
		"aws_wafv2_ip_set": "",
		"aws_wafv2_regex_pattern_set": "",
		"aws_wafv2_rule_group": "",
		"aws_wafv2_web_acl": "",
		"aws_workspaces_directory": "",
		"aws_workspaces_ip_group": "",
		"aws_workspaces_workspace": "",
		"aws_xray_group": "",
		"aws_xray_sampling_rule": "",
	}

	resource[p]
}

empty_array(arr) {
	arr == []
} else {
	arr == null
} else {
	false
}

uses_aws_managed_key(key, awsManagedKey) {
	key == awsManagedKey
} else {
	keyName := split(key, ".")[2]
	kms := input.document[z].data.aws_kms_key[keyName]
	kms.key_id == awsManagedKey
}

getStatement(policy) = st {
	is_array(policy.Statement)
	st = policy.Statement
} else = st {
	st := [policy.Statement]
}

is_publicly_accessible(policy) {
	statements := getStatement(policy)
	statement:= statements[_]
	statement.Effect == "Allow"
	anyPrincipal(statement)
}

get_accessibility(resource, name, resourcePolicyName, resourceTarget) = info {
	policy := common_lib.json_unmarshal(resource.policy)
	is_publicly_accessible(policy)
	info = {"accessibility": "public", "policy": policy}
} else = info {
	policy := common_lib.json_unmarshal(resource.policy)
	not is_publicly_accessible(policy)
	info = {"accessibility": "hasPolicy", "policy": policy}
} else = info {
	not common_lib.valid_key(resource, "policy")

	resourcePolicy := input.document[_].resource[resourcePolicyName][_]
	split(resourcePolicy[resourceTarget], ".")[1] == name

	policy := common_lib.json_unmarshal(resourcePolicy.policy)
	is_publicly_accessible(policy)
	info = {"accessibility": "public", "policy": policy}
} else = info {
	not common_lib.valid_key(resource, "policy")

	resourcePolicy := input.document[_].resource[resourcePolicyName][_]
	split(resourcePolicy[resourceTarget], ".")[1] == name

	policy := common_lib.json_unmarshal(resourcePolicy.policy)
	not is_publicly_accessible(policy)
	info = {"accessibility": "hasPolicy", "policy": policy}
} else = info {
	info = {"accessibility": "unknown", "policy": ""}
}

is_default_password(password) = output {
   contains(password, data.common_lib.default_passwords[_])
   output = true
} else = output {
	# repetition of the same number more than three times
	regex.match(`(0{3,}|1{3,}|2{3,}|3{3,}|4{3,}|5{3,}|6{3,}|7{3,}|8{3,}|9{3,})`, password) == true
	output = true
} else = output {
	output = false
}

matches(target, name) {
	split(target, ".")[1] == name
} else {
	target == name
}

check_member(attribute, search) {
	startswith(attribute.members[_], search)
} else {
	startswith(attribute.member, search)
}

check_member(attribute, search) {
	startswith(attribute.members[_], search)
} else {
	startswith(attribute.member, search)
}

matches(target, name) {
	split(target, ".")[1] == name
} else {
	target == name
}

matches(target, name) {
	split(target, ".")[1] == name
} else {
	target == name
}


has_target_resource(bucketName, resourceName) {
	resource := input.document[i].resource[resourceName][_]

	split(resource.bucket, ".")[1] == bucketName
}

#Checks if an action is allowed for all principals
allows_action_from_all_principals(json_policy, action) {
 	policy := common_lib.json_unmarshal(json_policy)
	st := common_lib.get_statement(policy)
	statement := st[_]
	statement.Effect == "Allow"
    anyPrincipal(statement)
    common_lib.containsOrInArrayContains(statement.Action, action)
}

resourceFieldName = {
	"google_bigquery_dataset": "friendly_name",
	"alicloud_actiontrail_trail": "trail_name",
	"alicloud_ros_stack": "stack_name",
	"alicloud_oss_bucket": "bucket",
	"aws_s3_bucket": "bucket",
	"aws_msk_cluster": "cluster_name",
	"aws_mq_broker": "broker_name",
	"aws_elasticache_cluster": "cluster_id",
}

compute_raw_resource_name(resource, resourceDefinitionName) = name {
	name := resource["name"]
} else = name {
	name := resource["display_name"]
}  else = name {
	name := resource.metadata.name
} else = name {
	prefix := resource.name_prefix
	name := sprintf("%s<unknown-sufix>", [prefix])
} else = name {
	name := common_lib.get_tag_name_if_exists(resource)
} else = name {
	name := resourceDefinitionName
}

get_resource_name(resource, resourceDefinitionName) = final_name {
    final_name := resourceDefinitionName
}

get_specific_resource_name(resource, resourceType, resourceDefinitionName) = name {
	field := resourceFieldName[resourceType]
	name := resource[field]
} else = name {
	name := get_resource_name(resource, resourceDefinitionName)
}

check_key_empty(disk_encryption_key) = key {
	common_lib.valid_key(disk_encryption_key, "raw_key")
	common_lib.emptyOrNull(disk_encryption_key.raw_key)
	key := "raw_key"
} else = key {
	common_lib.valid_key(disk_encryption_key, "kms_key_self_link")
	common_lib.emptyOrNull(disk_encryption_key.kms_key_self_link)
	key := "kms_key_self_link"
}

check_gcp_resource_supports_labels(p) {
	resource = {
        "google_active_directory_domain": "",
        "google_alloydb_backup": "",
        "google_alloydb_cluster": "",
        "google_alloydb_instance": "",
        "google_artifact_registry_repository": "",
        "google_assured_workloads_workload": "",
        "google_beyondcorp_app_connection": "",
        "google_beyondcorp_app_connector": "",
        "google_beyondcorp_app_gateway": "",
        "google_bigquery_dataset": "",
        "google_bigquery_job": "",
        "google_bigquery_table": "",
        "google_bigtable_instance": "",
        "google_certificate_manager_certificate": "",
        "google_certificate_manager_certificate_issuance_config": "",
        "google_certificate_manager_certificate_map": "",
        "google_certificate_manager_certificate_map_entry": "",
        "google_certificate_manager_dns_authorization": "",
        "google_certificate_manager_trust_config": "",
        "google_cloud_identity_group": "",
        "google_cloud_run_v2_job": "",
        "google_cloud_run_v2_service": "",
        "google_clouddeploy_delivery_pipeline": "",
        "google_clouddeploy_target": "",
        "google_cloudfunctions2_function": "",
        "google_cloudfunctions_function": "",
        "google_composer_environment": "",
        "google_compute_address": "",
        "google_compute_disk": "",
        "google_compute_external_vpn_gateway": "",
        "google_compute_forwarding_rule": "",
        "google_compute_global_forwarding_rule": "",
        "google_compute_image": "",
        "google_compute_instance": "",
        "google_compute_instance_from_template": "",
        "google_compute_instance_template": "",
        "google_compute_region_disk": "",
        "google_compute_region_instance_template": "",
        "google_compute_snapshot": "",
        "google_compute_vpn_tunnel": "",
        "google_data_fusion_instance": "",
        "google_database_migration_service_connection_profile": "",
        "google_database_migration_service_private_connection": "",
        "google_dataflow_job": "",
        "google_dataplex_asset": "",
        "google_dataplex_datascan": "",
        "google_dataplex_lake": "",
        "google_dataplex_task": "",
        "google_dataplex_zone": "",
        "google_dataproc_cluster": "",
        "google_dataproc_job": "",
        "google_dataproc_metastore_service": "",
        "google_dataproc_workflow_template": "",
        "google_datastream_connection_profile": "",
        "google_datastream_private_connection": "",
        "google_datastream_stream": "",
        "google_dialogflow_cx_intent": "",
        "google_dns_managed_zone": "",
        "google_edgecontainer_cluster": "",
        "google_edgecontainer_node_pool": "",
        "google_edgecontainer_vpn_connection": "",
        "google_edgenetwork_network": "",
        "google_edgenetwork_subnet": "",
        "google_eventarc_trigger": "",
        "google_filestore_backup": "",
        "google_filestore_instance": "",
        "google_filestore_snapshot": "",
        "google_gke_backup_backup_plan": "",
        "google_gke_backup_restore_plan": "",
        "google_gke_hub_feature": "",
        "google_gke_hub_membership": "",
        "google_gke_hub_membership_binding": "",
        "google_gke_hub_namespace": "",
        "google_gke_hub_scope": "",
        "google_gke_hub_scope_rbac_role_binding": "",
        "google_healthcare_consent_store": "",
        "google_healthcare_dicom_store": "",
        "google_healthcare_fhir_store": "",
        "google_healthcare_hl7_v2_store": "",
        "google_integration_connectors_connection": "",
        "google_kms_crypto_key": "",
        "google_memcache_instance": "",
        "google_ml_engine_model": "",
        "google_monitoring_notification_channel": "",
        "google_network_connectivity_hub": "",
        "google_network_connectivity_policy_based_route": "",
        "google_network_connectivity_service_connection_policy": "",
        "google_network_connectivity_spoke": "",
        "google_network_management_connectivity_test": "",
        "google_network_security_address_group": "",
        "google_network_services_edge_cache_keyset": "",
        "google_network_services_edge_cache_origin": "",
        "google_network_services_edge_cache_service": "",
        "google_network_services_gateway": "",
        "google_notebooks_instance": "",
        "google_privateca_ca_pool": "",
        "google_privateca_certificate": "",
        "google_privateca_certificate_authority": "",
        "google_privateca_certificate_template": "",
        "google_project": "",
        "google_pubsub_subscription": "",
        "google_pubsub_topic": "",
        "google_recaptcha_enterprise_key": "",
        "google_redis_instance": "",
        "google_secret_manager_secret": "",
        "google_spanner_instance": "",
        "google_storage_bucket": "",
        "google_tpu_node": "",
        "google_vertex_ai_dataset": "",
        "google_vertex_ai_endpoint": "",
        "google_vertex_ai_featurestore": "",
        "google_vertex_ai_featurestore_entitytype": "",
        "google_vertex_ai_featurestore_entitytype_feature": "",
        "google_vertex_ai_index": "",
        "google_vertex_ai_index_endpoint": "",
        "google_vertex_ai_tensorboard": "",
        "google_workflows_workflow": ""
    }
	resource[p]
}

check_azure_resource_supports_tags(p) {
	resource = {
        "azurerm_aadb2c_directory": "",
        "azurerm_active_directory_domain_service": "",
        "azurerm_analysis_services_server": "",
        "azurerm_api_connection": "",
        "azurerm_api_management": "",
        "azurerm_api_management_named_value": "",
        "azurerm_app_configuration": "",
        "azurerm_app_configuration_feature": "",
        "azurerm_app_configuration_key": "",
        "azurerm_app_service": "",
        "azurerm_app_service_certificate": "",
        "azurerm_app_service_certificate_order": "",
        "azurerm_app_service_environment": "",
        "azurerm_app_service_environment_v3": "",
        "azurerm_app_service_managed_certificate": "",
        "azurerm_app_service_plan": "",
        "azurerm_app_service_slot": "",
        "azurerm_application_gateway": "",
        "azurerm_application_insights": "",
        "azurerm_application_insights_standard_web_test": "",
        "azurerm_application_insights_web_test": "",
        "azurerm_application_insights_workbook": "",
        "azurerm_application_insights_workbook_template": "",
        "azurerm_application_load_balancer": "",
        "azurerm_application_load_balancer_frontend": "",
        "azurerm_application_load_balancer_subnet_association": "",
        "azurerm_application_security_group": "",
        "azurerm_arc_kubernetes_cluster": "",
        "azurerm_arc_machine_extension": "",
        "azurerm_arc_private_link_scope": "",
        "azurerm_arc_resource_bridge_appliance": "",
        "azurerm_attestation_provider": "",
        "azurerm_automanage_configuration": "",
        "azurerm_automation_account": "",
        "azurerm_automation_dsc_configuration": "",
        "azurerm_automation_python3_package": "",
        "azurerm_automation_runbook": "",
        "azurerm_automation_watcher": "",
        "azurerm_availability_set": "",
        "azurerm_bastion_host": "",
        "azurerm_batch_account": "",
        "azurerm_bot_channels_registration": "",
        "azurerm_bot_connection": "",
        "azurerm_bot_service_azure_bot": "",
        "azurerm_bot_web_app": "",
        "azurerm_capacity_reservation": "",
        "azurerm_capacity_reservation_group": "",
        "azurerm_cdn_endpoint": "",
        "azurerm_cdn_frontdoor_endpoint": "",
        "azurerm_cdn_frontdoor_firewall_policy": "",
        "azurerm_cdn_frontdoor_profile": "",
        "azurerm_cdn_profile": "",
        "azurerm_cognitive_account": "",
        "azurerm_communication_service": "",
        "azurerm_confidential_ledger": "",
        "azurerm_container_app": "",
        "azurerm_container_app_environment": "",
        "azurerm_container_app_environment_certificate": "",
        "azurerm_container_group": "",
        "azurerm_container_registry": "",
        "azurerm_container_registry_agent_pool": "",
        "azurerm_container_registry_task": "",
        "azurerm_container_registry_webhook": "",
        "azurerm_cosmosdb_account": "",
        "azurerm_cosmosdb_cassandra_cluster": "",
        "azurerm_cosmosdb_postgresql_cluster": "",
        "azurerm_custom_ip_prefix": "",
        "azurerm_custom_provider": "",
        "azurerm_dashboard": "",
        "azurerm_dashboard_grafana": "",
        "azurerm_data_factory": "",
        "azurerm_data_protection_backup_vault": "",
        "azurerm_data_protection_resource_guard": "",
        "azurerm_data_share_account": "",
        "azurerm_database_migration_project": "",
        "azurerm_database_migration_service": "",
        "azurerm_databox_edge_device": "",
        "azurerm_databricks_access_connector": "",
        "azurerm_databricks_workspace": "",
        "azurerm_datadog_monitor": "",
        "azurerm_dedicated_hardware_security_module": "",
        "azurerm_dedicated_host": "",
        "azurerm_dedicated_host_group": "",
        "azurerm_dev_center": "",
        "azurerm_dev_center_project": "",
        "azurerm_dev_test_global_vm_shutdown_schedule": "",
        "azurerm_dev_test_lab": "",
        "azurerm_dev_test_linux_virtual_machine": "",
        "azurerm_dev_test_policy": "",
        "azurerm_dev_test_schedule": "",
        "azurerm_dev_test_virtual_network": "",
        "azurerm_dev_test_windows_virtual_machine": "",
        "azurerm_digital_twins_instance": "",
        "azurerm_disk_access": "",
        "azurerm_disk_encryption_set": "",
        "azurerm_disk_pool": "",
        "azurerm_dns_a_record": "",
        "azurerm_dns_aaaa_record": "",
        "azurerm_dns_caa_record": "",
        "azurerm_dns_cname_record": "",
        "azurerm_dns_mx_record": "",
        "azurerm_dns_ns_record": "",
        "azurerm_dns_ptr_record": "",
        "azurerm_dns_srv_record": "",
        "azurerm_dns_txt_record": "",
        "azurerm_dns_zone": "",
        "azurerm_elastic_cloud_elasticsearch": "",
        "azurerm_email_communication_service": "",
        "azurerm_eventgrid_domain": "",
        "azurerm_eventgrid_system_topic": "",
        "azurerm_eventgrid_topic": "",
        "azurerm_eventhub_cluster": "",
        "azurerm_eventhub_namespace": "",
        "azurerm_express_route_circuit": "",
        "azurerm_express_route_gateway": "",
        "azurerm_express_route_port": "",
        "azurerm_firewall": "",
        "azurerm_firewall_policy": "",
        "azurerm_fluid_relay_server": "",
        "azurerm_frontdoor": "",
        "azurerm_frontdoor_firewall_policy": "",
        "azurerm_function_app": "",
        "azurerm_function_app_slot": "",
        "azurerm_gallery_application": "",
        "azurerm_gallery_application_version": "",
        "azurerm_graph_account": "",
        "azurerm_graph_services_account": "",
        "azurerm_hdinsight_hadoop_cluster": "",
        "azurerm_hdinsight_hbase_cluster": "",
        "azurerm_hdinsight_interactive_query_cluster": "",
        "azurerm_hdinsight_kafka_cluster": "",
        "azurerm_hdinsight_spark_cluster": "",
        "azurerm_healthbot": "",
        "azurerm_healthcare_dicom_service": "",
        "azurerm_healthcare_fhir_service": "",
        "azurerm_healthcare_medtech_service": "",
        "azurerm_healthcare_service": "",
        "azurerm_healthcare_workspace": "",
        "azurerm_hpc_cache": "",
        "azurerm_image": "",
        "azurerm_integration_service_environment": "",
        "azurerm_iot_security_solution": "",
        "azurerm_iot_time_series_insights_event_source_eventhub": "",
        "azurerm_iot_time_series_insights_event_source_iothub": "",
        "azurerm_iot_time_series_insights_gen2_environment": "",
        "azurerm_iot_time_series_insights_reference_data_set": "",
        "azurerm_iot_time_series_insights_standard_environment": "",
        "azurerm_iotcentral_application": "",
        "azurerm_iothub": "",
        "azurerm_iothub_device_update_account": "",
        "azurerm_iothub_device_update_instance": "",
        "azurerm_iothub_dps": "",
        "azurerm_ip_group": "",
        "azurerm_key_vault": "",
        "azurerm_key_vault_certificate": "",
        "azurerm_key_vault_key": "",
        "azurerm_key_vault_managed_hardware_security_module": "",
        "azurerm_key_vault_managed_storage_account": "",
        "azurerm_key_vault_managed_storage_account_sas_token_definition": "",
        "azurerm_key_vault_secret": "",
        "azurerm_kubernetes_cluster": "",
        "azurerm_kubernetes_cluster_node_pool": "",
        "azurerm_kubernetes_fleet_manager": "",
        "azurerm_kusto_cluster": "",
        "azurerm_lab_service_lab": "",
        "azurerm_lab_service_plan": "",
        "azurerm_lb": "",
        "azurerm_linux_function_app": "",
        "azurerm_linux_function_app_slot": "",
        "azurerm_linux_virtual_machine": "",
        "azurerm_linux_virtual_machine_scale_set": "",
        "azurerm_linux_web_app": "",
        "azurerm_linux_web_app_slot": "",
        "azurerm_load_test": "",
        "azurerm_local_network_gateway": "",
        "azurerm_log_analytics_cluster": "",
        "azurerm_log_analytics_query_pack": "",
        "azurerm_log_analytics_query_pack_query": "",
        "azurerm_log_analytics_saved_search": "",
        "azurerm_log_analytics_solution": "",
        "azurerm_log_analytics_workspace": "",
        "azurerm_logic_app_integration_account": "",
        "azurerm_logic_app_standard": "",
        "azurerm_logic_app_workflow": "",
        "azurerm_logz_monitor": "",
        "azurerm_logz_sub_account": "",
        "azurerm_machine_learning_compute_cluster": "",
        "azurerm_machine_learning_compute_instance": "",
        "azurerm_machine_learning_datastore_blobstorage": "",
        "azurerm_machine_learning_datastore_datalake_gen2": "",
        "azurerm_machine_learning_datastore_fileshare": "",
        "azurerm_machine_learning_inference_cluster": "",
        "azurerm_machine_learning_synapse_spark": "",
        "azurerm_machine_learning_workspace": "",
        "azurerm_maintenance_configuration": "",
        "azurerm_managed_application": "",
        "azurerm_managed_application_definition": "",
        "azurerm_managed_disk": "",
        "azurerm_managed_lustre_file_system": "",
        "azurerm_management_group_template_deployment": "",
        "azurerm_maps_account": "",
        "azurerm_maps_creator": "",
        "azurerm_mariadb_server": "",
        "azurerm_media_live_event": "",
        "azurerm_media_services_account": "",
        "azurerm_media_streaming_endpoint": "",
        "azurerm_mobile_network": "",
        "azurerm_mobile_network_attached_data_network": "",
        "azurerm_mobile_network_data_network": "",
        "azurerm_mobile_network_packet_core_control_plane": "",
        "azurerm_mobile_network_packet_core_data_plane": "",
        "azurerm_mobile_network_service": "",
        "azurerm_mobile_network_sim_group": "",
        "azurerm_mobile_network_sim_policy": "",
        "azurerm_mobile_network_site": "",
        "azurerm_mobile_network_slice": "",
        "azurerm_monitor_action_group": "",
        "azurerm_monitor_action_rule_action_group": "",
        "azurerm_monitor_action_rule_suppression": "",
        "azurerm_monitor_activity_log_alert": "",
        "azurerm_monitor_alert_processing_rule_action_group": "",
        "azurerm_monitor_alert_processing_rule_suppression": "",
        "azurerm_monitor_alert_prometheus_rule_group": "",
        "azurerm_monitor_autoscale_setting": "",
        "azurerm_monitor_data_collection_endpoint": "",
        "azurerm_monitor_data_collection_rule": "",
        "azurerm_monitor_metric_alert": "",
        "azurerm_monitor_private_link_scope": "",
        "azurerm_monitor_scheduled_query_rules_alert": "",
        "azurerm_monitor_scheduled_query_rules_alert_v2": "",
        "azurerm_monitor_scheduled_query_rules_log": "",
        "azurerm_monitor_smart_detector_alert_rule": "",
        "azurerm_monitor_workspace": "",
        "azurerm_mssql_database": "",
        "azurerm_mssql_elasticpool": "",
        "azurerm_mssql_failover_group": "",
        "azurerm_mssql_job_agent": "",
        "azurerm_mssql_managed_instance": "",
        "azurerm_mssql_server": "",
        "azurerm_mssql_virtual_machine": "",
        "azurerm_mssql_virtual_machine_group": "",
        "azurerm_mysql_flexible_server": "",
        "azurerm_mysql_server": "",
        "azurerm_nat_gateway": "",
        "azurerm_netapp_account": "",
        "azurerm_netapp_pool": "",
        "azurerm_netapp_snapshot_policy": "",
        "azurerm_netapp_volume": "",
        "azurerm_network_connection_monitor": "",
        "azurerm_network_ddos_protection_plan": "",
        "azurerm_network_function_azure_traffic_collector": "",
        "azurerm_network_function_collector_policy": "",
        "azurerm_network_interface": "",
        "azurerm_network_manager": "",
        "azurerm_network_profile": "",
        "azurerm_network_security_group": "",
        "azurerm_network_watcher": "",
        "azurerm_network_watcher_flow_log": "",
        "azurerm_nginx_deployment": "",
        "azurerm_notification_hub": "",
        "azurerm_notification_hub_namespace": "",
        "azurerm_orbital_contact_profile": "",
        "azurerm_orbital_spacecraft": "",
        "azurerm_orchestrated_virtual_machine_scale_set": "",
        "azurerm_palo_alto_local_rulestack_rule": "",
        "azurerm_palo_alto_next_generation_firewall_virtual_hub_local_rulestack": "",
        "azurerm_palo_alto_next_generation_firewall_virtual_hub_panorama": "",
        "azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack": "",
        "azurerm_palo_alto_next_generation_firewall_virtual_network_panorama": "",
        "azurerm_point_to_site_vpn_gateway": "",
        "azurerm_portal_dashboard": "",
        "azurerm_postgresql_flexible_server": "",
        "azurerm_postgresql_server": "",
        "azurerm_powerbi_embedded": "",
        "azurerm_private_dns_a_record": "",
        "azurerm_private_dns_aaaa_record": "",
        "azurerm_private_dns_cname_record": "",
        "azurerm_private_dns_mx_record": "",
        "azurerm_private_dns_ptr_record": "",
        "azurerm_private_dns_resolver": "",
        "azurerm_private_dns_resolver_dns_forwarding_ruleset": "",
        "azurerm_private_dns_resolver_inbound_endpoint": "",
        "azurerm_private_dns_resolver_outbound_endpoint": "",
        "azurerm_private_dns_srv_record": "",
        "azurerm_private_dns_txt_record": "",
        "azurerm_private_dns_zone": "",
        "azurerm_private_dns_zone_virtual_network_link": "",
        "azurerm_private_endpoint": "",
        "azurerm_private_link_service": "",
        "azurerm_proximity_placement_group": "",
        "azurerm_public_ip": "",
        "azurerm_public_ip_prefix": "",
        "azurerm_purview_account": "",
        "azurerm_recovery_services_vault": "",
        "azurerm_redis_cache": "",
        "azurerm_redis_enterprise_cluster": "",
        "azurerm_relay_namespace": "",
        "azurerm_resource_deployment_script_azure_cli": "",
        "azurerm_resource_deployment_script_azure_power_shell": "",
        "azurerm_resource_group": "",
        "azurerm_resource_group_template_deployment": "",
        "azurerm_route_filter": "",
        "azurerm_route_server": "",
        "azurerm_route_table": "",
        "azurerm_search_service": "",
        "azurerm_security_center_automation": "",
        "azurerm_sentinel_threat_intelligence_indicator": "",
        "azurerm_service_fabric_cluster": "",
        "azurerm_service_fabric_managed_cluster": "",
        "azurerm_service_plan": "",
        "azurerm_servicebus_namespace": "",
        "azurerm_shared_image": "",
        "azurerm_shared_image_gallery": "",
        "azurerm_shared_image_version": "",
        "azurerm_signalr_service": "",
        "azurerm_snapshot": "",
        "azurerm_spatial_anchors_account": "",
        "azurerm_spring_cloud_service": "",
        "azurerm_sql_database": "",
        "azurerm_sql_elasticpool": "",
        "azurerm_sql_failover_group": "",
        "azurerm_sql_managed_instance": "",
        "azurerm_sql_server": "",
        "azurerm_ssh_public_key": "",
        "azurerm_stack_hci_cluster": "",
        "azurerm_static_site": "",
        "azurerm_storage_account": "",
        "azurerm_storage_mover": "",
        "azurerm_storage_sync": "",
        "azurerm_stream_analytics_cluster": "",
        "azurerm_stream_analytics_job": "",
        "azurerm_subnet_service_endpoint_storage_policy": "",
        "azurerm_subscription": "",
        "azurerm_subscription_template_deployment": "",
        "azurerm_synapse_private_link_hub": "",
        "azurerm_synapse_spark_pool": "",
        "azurerm_synapse_sql_pool": "",
        "azurerm_synapse_workspace": "",
        "azurerm_tenant_template_deployment": "",
        "azurerm_traffic_manager_profile": "",
        "azurerm_user_assigned_identity": "",
        "azurerm_video_analyzer": "",
        "azurerm_virtual_desktop_application_group": "",
        "azurerm_virtual_desktop_host_pool": "",
        "azurerm_virtual_desktop_scaling_plan": "",
        "azurerm_virtual_desktop_workspace": "",
        "azurerm_virtual_hub": "",
        "azurerm_virtual_hub_security_partner_provider": "",
        "azurerm_virtual_machine": "",
        "azurerm_virtual_machine_extension": "",
        "azurerm_virtual_machine_scale_set": "",
        "azurerm_virtual_network": "",
        "azurerm_virtual_network_gateway": "",
        "azurerm_virtual_network_gateway_connection": "",
        "azurerm_virtual_wan": "",
        "azurerm_vmware_private_cloud": "",
        "azurerm_voice_services_communications_gateway": "",
        "azurerm_voice_services_communications_gateway_test_line": "",
        "azurerm_vpn_gateway": "",
        "azurerm_vpn_server_configuration": "",
        "azurerm_vpn_site": "",
        "azurerm_web_application_firewall_policy": "",
        "azurerm_web_pubsub": "",
        "azurerm_windows_function_app": "",
        "azurerm_windows_function_app_slot": "",
        "azurerm_windows_virtual_machine": "",
        "azurerm_windows_virtual_machine_scale_set": "",
        "azurerm_windows_web_app": "",
        "azurerm_windows_web_app_slot": ""
    }
	resource[p]
}
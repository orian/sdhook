package sdhook

// ResType is a monitored resource descriptor type.
//
// See https://cloud.google.com/logging/docs/api/v2/resource-list
type ResType string

const (
	// ResTypeApi, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_api
	ResTypeApi ResType = "api"
	// ResTypeAppScriptFunction, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_app_script_function
	ResTypeAppScriptFunction ResType = "app_script_function"
	// ResTypeAssistantAction, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_assistant_action
	ResTypeAssistantAction ResType = "assistant_action"
	// ResTypeAuditedResource, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_audited_resource
	ResTypeAuditedResource ResType = "audited_resource"
	// ResTypeAwsAlbLoadBalancer, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_alb_load_balancer
	ResTypeAwsAlbLoadBalancer ResType = "aws_alb_load_balancer"
	// ResTypeAwsCloudfrontDistribution, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_cloudfront_distribution
	ResTypeAwsCloudfrontDistribution ResType = "aws_cloudfront_distribution"
	// ResTypeAwsDynamodbTable, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_dynamodb_table
	ResTypeAwsDynamodbTable ResType = "aws_dynamodb_table"
	// ResTypeAwsEbsVolume, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_ebs_volume
	ResTypeAwsEbsVolume ResType = "aws_ebs_volume"
	// ResTypeAwsEc2_instance, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_ec2_instance
	ResTypeAwsEc2_instance ResType = "aws_ec2_instance"
	// ResTypeAwsElasticacheCluster, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_elasticache_cluster
	ResTypeAwsElasticacheCluster ResType = "aws_elasticache_cluster"
	// ResTypeAwsElbLoadBalancer, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_elb_load_balancer
	ResTypeAwsElbLoadBalancer ResType = "aws_elb_load_balancer"
	// ResTypeAwsEmrCluster, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_emr_cluster
	ResTypeAwsEmrCluster ResType = "aws_emr_cluster"
	// ResTypeAwsKinesisStream, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_kinesis_stream
	ResTypeAwsKinesisStream ResType = "aws_kinesis_stream"
	// ResTypeAwsLambdaFunction, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_lambda_function
	ResTypeAwsLambdaFunction ResType = "aws_lambda_function"
	// ResTypeAwsRdsDatabase, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_rds_database
	ResTypeAwsRdsDatabase ResType = "aws_rds_database"
	// ResTypeAwsRedshiftCluster, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_redshift_cluster
	ResTypeAwsRedshiftCluster ResType = "aws_redshift_cluster"
	// ResTypeAwsS3_bucket, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_s3_bucket
	ResTypeAwsS3_bucket ResType = "aws_s3_bucket"
	// ResTypeAwsSes, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_ses
	ResTypeAwsSes ResType = "aws_ses"
	// ResTypeAwsSnsTopic, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_sns_topic
	ResTypeAwsSnsTopic ResType = "aws_sns_topic"
	// ResTypeAwsSqsQueue, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_aws_sqs_queue
	ResTypeAwsSqsQueue ResType = "aws_sqs_queue"
	// ResTypeBigqueryDataset, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_bigquery_dataset
	ResTypeBigqueryDataset ResType = "bigquery_dataset"
	// ResTypeBigqueryProject, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_bigquery_project
	ResTypeBigqueryProject ResType = "bigquery_project"
	// ResTypeBigtableCluster, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_bigtable_cluster
	ResTypeBigtableCluster ResType = "bigtable_cluster"
	// ResTypeBigtableTable, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_bigtable_table
	ResTypeBigtableTable ResType = "bigtable_table"
	// ResTypeBuild, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_build
	ResTypeBuild ResType = "build"
	// ResTypeCloudComposerEnvironment, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloud_composer_environment
	ResTypeCloudComposerEnvironment ResType = "cloud_composer_environment"
	// ResTypeCloudDataprocCluster, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloud_dataproc_cluster
	ResTypeCloudDataprocCluster ResType = "cloud_dataproc_cluster"
	// ResTypeCloudDataprocJob, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloud_dataproc_job
	ResTypeCloudDataprocJob ResType = "cloud_dataproc_job"
	// ResTypeCloudFunction, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloud_function
	ResTypeCloudFunction ResType = "cloud_function"
	// ResTypeCloudRunRevision, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloud_run_revision
	ResTypeCloudRunRevision ResType = "cloud_run_revision"
	// ResTypeCloudSchedulerJob, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloud_scheduler_job
	ResTypeCloudSchedulerJob ResType = "cloud_scheduler_job"
	// ResTypeCloudTasksQueue, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloud_tasks_queue
	ResTypeCloudTasksQueue ResType = "cloud_tasks_queue"
	// ResTypeCloudiotDevice, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloudiot_device
	ResTypeCloudiotDevice ResType = "cloudiot_device"
	// ResTypeCloudiotDeviceRegistry, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloudiot_device_registry
	ResTypeCloudiotDeviceRegistry ResType = "cloudiot_device_registry"
	// ResTypeCloudmlJob, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloudml_job
	ResTypeCloudmlJob ResType = "cloudml_job"
	// ResTypeCloudmlModelVersion, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloudml_model_version
	ResTypeCloudmlModelVersion ResType = "cloudml_model_version"
	// ResTypeCloudsqlDatabase, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_cloudsql_database
	ResTypeCloudsqlDatabase ResType = "cloudsql_database"
	// ResTypeConsumedApi, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_consumed_api
	ResTypeConsumedApi ResType = "consumed_api"
	// ResTypeCsrRepository, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_csr_repository
	ResTypeCsrRepository ResType = "csr_repository"
	// ResTypeDataflowJob, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_dataflow_job
	ResTypeDataflowJob ResType = "dataflow_job"
	// ResTypeDatastoreRequest, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_datastore_request
	ResTypeDatastoreRequest ResType = "datastore_request"
	// ResTypeDnsManagedZone, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_dns_managed_zone
	ResTypeDnsManagedZone ResType = "dns_managed_zone"
	// ResTypeDnsPolicy, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_dns_policy
	ResTypeDnsPolicy ResType = "dns_policy"
	// ResTypeDnsQuery, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_dns_query
	ResTypeDnsQuery ResType = "dns_query"
	// ResTypeFilestoreInstance, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_filestore_instance
	ResTypeFilestoreInstance ResType = "filestore_instance"
	// ResTypeFirebaseDomain, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_firebase_domain
	ResTypeFirebaseDomain ResType = "firebase_domain"
	// ResTypeFirebaseNamespace, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_firebase_namespace
	ResTypeFirebaseNamespace ResType = "firebase_namespace"
	// ResTypeFirestoreInstance, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_firestore_instance
	ResTypeFirestoreInstance ResType = "firestore_instance"
	// ResTypeGaeApp, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_gae_app
	ResTypeGaeApp ResType = "gae_app"
	// ResTypeGaeInstance, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_gae_instance
	ResTypeGaeInstance ResType = "gae_instance"
	ResTypeGceProject  ResType = "gce_project"
	// ResTypeGceDisk, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_gce_disk
	ResTypeGceDisk ResType = "gce_disk"
	// ResTypeGceInstance, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_gce_instance
	ResTypeGceInstance ResType = "gce_instance"
	// ResTypeGceNodeGroup, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_gce_node_group
	ResTypeGceNodeGroup ResType = "gce_node_group"
	// ResTypeGceNodeTemplate, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_gce_node_template
	ResTypeGceNodeTemplate ResType = "gce_node_template"
	// ResTypeGceResourcePolicy, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_gce_resource_policy
	ResTypeGceResourcePolicy ResType = "gce_resource_policy"
	// ResTypeGceRouter, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_gce_router
	ResTypeGceRouter ResType = "gce_router"
	// ResTypeGcsBucket, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_gcs_bucket
	ResTypeGcsBucket ResType = "gcs_bucket"
	// ResTypeGenericNode, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_generic_node
	ResTypeGenericNode ResType = "generic_node"
	// ResTypeGenericTask, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_generic_task
	ResTypeGenericTask ResType = "generic_task"
	// ResTypeGkeContainer, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_gke_container
	ResTypeGkeContainer ResType = "gke_container"
	// ResTypeGlobal, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_global
	ResTypeGlobal ResType = "global"
	// ResTypeHttpsLbRule, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_https_lb_rule
	ResTypeHttpsLbRule ResType = "https_lb_rule"
	// ResTypeIdentitytoolkitProject, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_identitytoolkit_project
	ResTypeIdentitytoolkitProject ResType = "identitytoolkit_project"
	// ResTypeIdentitytoolkitTenant, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_identitytoolkit_tenant
	ResTypeIdentitytoolkitTenant ResType = "identitytoolkit_tenant"
	// ResTypeInterconnect, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_interconnect
	ResTypeInterconnect ResType = "interconnect"
	// ResTypeInterconnectAttachment, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_interconnect_attachment
	ResTypeInterconnectAttachment ResType = "interconnect_attachment"
	// ResTypeInternalHttpLbRule, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_internal_http_lb_rule
	ResTypeInternalHttpLbRule ResType = "internal_http_lb_rule"
	// ResTypeInternalTcpLbRule, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_internal_tcp_lb_rule
	ResTypeInternalTcpLbRule ResType = "internal_tcp_lb_rule"
	// ResTypeInternalUdpLbRule, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_internal_udp_lb_rule
	ResTypeInternalUdpLbRule ResType = "internal_udp_lb_rule"
	// ResTypeK8s_cluster, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_k8s_cluster
	ResTypeK8s_cluster ResType = "k8s_cluster"
	// ResTypeK8s_container, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_k8s_container
	ResTypeK8s_container ResType = "k8s_container"
	// ResTypeK8s_node, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_k8s_node
	ResTypeK8s_node ResType = "k8s_node"
	// ResTypeK8s_pod, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_k8s_pod
	ResTypeK8s_pod ResType = "k8s_pod"
	// ResTypeKnativeRevision, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_knative_revision
	ResTypeKnativeRevision ResType = "knative_revision"
	// ResTypeL7_lb_rule, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_l7_lb_rule
	ResTypeL7_lb_rule ResType = "l7_lb_rule"
	// ResTypeLoggingExclusion, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_logging_exclusion
	ResTypeLoggingExclusion ResType = "logging_exclusion"
	// ResTypeLoggingSink, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_logging_sink
	ResTypeLoggingSink ResType = "logging_sink"
	// ResTypeMetric, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_metric
	ResTypeMetric ResType = "metric"
	// ResTypeMicrosoftAdDomain, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_microsoft_ad_domain
	ResTypeMicrosoftAdDomain ResType = "microsoft_ad_domain"
	// ResTypeNatGateway, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_nat_gateway
	ResTypeNatGateway ResType = "nat_gateway"
	// ResTypeNetappCloudVolume, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_netapp_cloud_volume
	ResTypeNetappCloudVolume ResType = "netapp_cloud_volume"
	// ResTypeNetworkSecurityPolicy, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_network_security_policy
	ResTypeNetworkSecurityPolicy ResType = "network_security_policy"
	// ResTypePubsubSnapshot, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_pubsub_snapshot
	ResTypePubsubSnapshot ResType = "pubsub_snapshot"
	// ResTypePubsubSubscription, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_pubsub_subscription
	ResTypePubsubSubscription ResType = "pubsub_subscription"
	// ResTypePubsubTopic, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_pubsub_topic
	ResTypePubsubTopic ResType = "pubsub_topic"
	// ResTypeRecommender, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_recommender
	ResTypeRecommender ResType = "recommender"
	// ResTypeRedisInstance, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_redis_instance
	ResTypeRedisInstance ResType = "redis_instance"
	// ResTypeSpannerInstance, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_spanner_instance
	ResTypeSpannerInstance ResType = "spanner_instance"
	// ResTypeTcpLbRule, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_tcp_lb_rule
	ResTypeTcpLbRule ResType = "tcp_lb_rule"
	// ResTypeTcpSslProxyRule, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_tcp_ssl_proxy_rule
	ResTypeTcpSslProxyRule ResType = "tcp_ssl_proxy_rule"
	// ResTypeTpuWorker, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_tpu_worker
	ResTypeTpuWorker ResType = "tpu_worker"
	// ResTypeUdpLbRule, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_udp_lb_rule
	ResTypeUdpLbRule ResType = "udp_lb_rule"
	// ResTypeUptimeUrl, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_uptime_url
	ResTypeUptimeUrl ResType = "uptime_url"
	// ResTypeVpnGateway, for description and list of labels https://cloud.google.com/monitoring/api/resources#tag_vpn_gateway
	ResTypeVpnGateway ResType = "vpn_gateway"
)

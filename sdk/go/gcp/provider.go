// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the google-beta package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	AccessApprovalCustomEndpoint       pulumi.StringPtrOutput `pulumi:"accessApprovalCustomEndpoint"`
	AccessContextManagerCustomEndpoint pulumi.StringPtrOutput `pulumi:"accessContextManagerCustomEndpoint"`
	AccessToken                        pulumi.StringPtrOutput `pulumi:"accessToken"`
	ActiveDirectoryCustomEndpoint      pulumi.StringPtrOutput `pulumi:"activeDirectoryCustomEndpoint"`
	ApiGatewayCustomEndpoint           pulumi.StringPtrOutput `pulumi:"apiGatewayCustomEndpoint"`
	ApigeeCustomEndpoint               pulumi.StringPtrOutput `pulumi:"apigeeCustomEndpoint"`
	AppEngineCustomEndpoint            pulumi.StringPtrOutput `pulumi:"appEngineCustomEndpoint"`
	ArtifactRegistryCustomEndpoint     pulumi.StringPtrOutput `pulumi:"artifactRegistryCustomEndpoint"`
	AssuredWorkloadsCustomEndpoint     pulumi.StringPtrOutput `pulumi:"assuredWorkloadsCustomEndpoint"`
	BigQueryCustomEndpoint             pulumi.StringPtrOutput `pulumi:"bigQueryCustomEndpoint"`
	BigqueryConnectionCustomEndpoint   pulumi.StringPtrOutput `pulumi:"bigqueryConnectionCustomEndpoint"`
	BigqueryDataTransferCustomEndpoint pulumi.StringPtrOutput `pulumi:"bigqueryDataTransferCustomEndpoint"`
	BigqueryReservationCustomEndpoint  pulumi.StringPtrOutput `pulumi:"bigqueryReservationCustomEndpoint"`
	BigtableCustomEndpoint             pulumi.StringPtrOutput `pulumi:"bigtableCustomEndpoint"`
	BillingCustomEndpoint              pulumi.StringPtrOutput `pulumi:"billingCustomEndpoint"`
	BillingProject                     pulumi.StringPtrOutput `pulumi:"billingProject"`
	BinaryAuthorizationCustomEndpoint  pulumi.StringPtrOutput `pulumi:"binaryAuthorizationCustomEndpoint"`
	CloudAssetCustomEndpoint           pulumi.StringPtrOutput `pulumi:"cloudAssetCustomEndpoint"`
	CloudBillingCustomEndpoint         pulumi.StringPtrOutput `pulumi:"cloudBillingCustomEndpoint"`
	CloudBuildCustomEndpoint           pulumi.StringPtrOutput `pulumi:"cloudBuildCustomEndpoint"`
	CloudBuildWorkerPoolCustomEndpoint pulumi.StringPtrOutput `pulumi:"cloudBuildWorkerPoolCustomEndpoint"`
	CloudFunctionsCustomEndpoint       pulumi.StringPtrOutput `pulumi:"cloudFunctionsCustomEndpoint"`
	CloudIdentityCustomEndpoint        pulumi.StringPtrOutput `pulumi:"cloudIdentityCustomEndpoint"`
	CloudIotCustomEndpoint             pulumi.StringPtrOutput `pulumi:"cloudIotCustomEndpoint"`
	CloudRunCustomEndpoint             pulumi.StringPtrOutput `pulumi:"cloudRunCustomEndpoint"`
	CloudSchedulerCustomEndpoint       pulumi.StringPtrOutput `pulumi:"cloudSchedulerCustomEndpoint"`
	CloudTasksCustomEndpoint           pulumi.StringPtrOutput `pulumi:"cloudTasksCustomEndpoint"`
	ComposerCustomEndpoint             pulumi.StringPtrOutput `pulumi:"composerCustomEndpoint"`
	ComputeBetaCustomEndpoint          pulumi.StringPtrOutput `pulumi:"computeBetaCustomEndpoint"`
	ComputeCustomEndpoint              pulumi.StringPtrOutput `pulumi:"computeCustomEndpoint"`
	ContainerAnalysisCustomEndpoint    pulumi.StringPtrOutput `pulumi:"containerAnalysisCustomEndpoint"`
	ContainerBetaCustomEndpoint        pulumi.StringPtrOutput `pulumi:"containerBetaCustomEndpoint"`
	ContainerCustomEndpoint            pulumi.StringPtrOutput `pulumi:"containerCustomEndpoint"`
	Credentials                        pulumi.StringPtrOutput `pulumi:"credentials"`
	DataCatalogCustomEndpoint          pulumi.StringPtrOutput `pulumi:"dataCatalogCustomEndpoint"`
	DataFusionCustomEndpoint           pulumi.StringPtrOutput `pulumi:"dataFusionCustomEndpoint"`
	DataLossPreventionCustomEndpoint   pulumi.StringPtrOutput `pulumi:"dataLossPreventionCustomEndpoint"`
	DataflowCustomEndpoint             pulumi.StringPtrOutput `pulumi:"dataflowCustomEndpoint"`
	DataprocBetaCustomEndpoint         pulumi.StringPtrOutput `pulumi:"dataprocBetaCustomEndpoint"`
	DataprocCustomEndpoint             pulumi.StringPtrOutput `pulumi:"dataprocCustomEndpoint"`
	DataprocMetastoreCustomEndpoint    pulumi.StringPtrOutput `pulumi:"dataprocMetastoreCustomEndpoint"`
	DatastoreCustomEndpoint            pulumi.StringPtrOutput `pulumi:"datastoreCustomEndpoint"`
	DeploymentManagerCustomEndpoint    pulumi.StringPtrOutput `pulumi:"deploymentManagerCustomEndpoint"`
	DialogflowCustomEndpoint           pulumi.StringPtrOutput `pulumi:"dialogflowCustomEndpoint"`
	DialogflowCxCustomEndpoint         pulumi.StringPtrOutput `pulumi:"dialogflowCxCustomEndpoint"`
	DnsCustomEndpoint                  pulumi.StringPtrOutput `pulumi:"dnsCustomEndpoint"`
	EssentialContactsCustomEndpoint    pulumi.StringPtrOutput `pulumi:"essentialContactsCustomEndpoint"`
	EventarcCustomEndpoint             pulumi.StringPtrOutput `pulumi:"eventarcCustomEndpoint"`
	FilestoreCustomEndpoint            pulumi.StringPtrOutput `pulumi:"filestoreCustomEndpoint"`
	FirebaseCustomEndpoint             pulumi.StringPtrOutput `pulumi:"firebaseCustomEndpoint"`
	FirestoreCustomEndpoint            pulumi.StringPtrOutput `pulumi:"firestoreCustomEndpoint"`
	GameServicesCustomEndpoint         pulumi.StringPtrOutput `pulumi:"gameServicesCustomEndpoint"`
	GkeHubCustomEndpoint               pulumi.StringPtrOutput `pulumi:"gkeHubCustomEndpoint"`
	GkehubFeatureCustomEndpoint        pulumi.StringPtrOutput `pulumi:"gkehubFeatureCustomEndpoint"`
	GooglePartnerName                  pulumi.StringPtrOutput `pulumi:"googlePartnerName"`
	HealthcareCustomEndpoint           pulumi.StringPtrOutput `pulumi:"healthcareCustomEndpoint"`
	IamBetaCustomEndpoint              pulumi.StringPtrOutput `pulumi:"iamBetaCustomEndpoint"`
	IamCredentialsCustomEndpoint       pulumi.StringPtrOutput `pulumi:"iamCredentialsCustomEndpoint"`
	IamCustomEndpoint                  pulumi.StringPtrOutput `pulumi:"iamCustomEndpoint"`
	IapCustomEndpoint                  pulumi.StringPtrOutput `pulumi:"iapCustomEndpoint"`
	IdentityPlatformCustomEndpoint     pulumi.StringPtrOutput `pulumi:"identityPlatformCustomEndpoint"`
	ImpersonateServiceAccount          pulumi.StringPtrOutput `pulumi:"impersonateServiceAccount"`
	KmsCustomEndpoint                  pulumi.StringPtrOutput `pulumi:"kmsCustomEndpoint"`
	LoggingCustomEndpoint              pulumi.StringPtrOutput `pulumi:"loggingCustomEndpoint"`
	MemcacheCustomEndpoint             pulumi.StringPtrOutput `pulumi:"memcacheCustomEndpoint"`
	MlEngineCustomEndpoint             pulumi.StringPtrOutput `pulumi:"mlEngineCustomEndpoint"`
	MonitoringCustomEndpoint           pulumi.StringPtrOutput `pulumi:"monitoringCustomEndpoint"`
	NetworkManagementCustomEndpoint    pulumi.StringPtrOutput `pulumi:"networkManagementCustomEndpoint"`
	NetworkServicesCustomEndpoint      pulumi.StringPtrOutput `pulumi:"networkServicesCustomEndpoint"`
	NotebooksCustomEndpoint            pulumi.StringPtrOutput `pulumi:"notebooksCustomEndpoint"`
	OsConfigCustomEndpoint             pulumi.StringPtrOutput `pulumi:"osConfigCustomEndpoint"`
	OsLoginCustomEndpoint              pulumi.StringPtrOutput `pulumi:"osLoginCustomEndpoint"`
	PrivatecaCustomEndpoint            pulumi.StringPtrOutput `pulumi:"privatecaCustomEndpoint"`
	Project                            pulumi.StringPtrOutput `pulumi:"project"`
	PubsubCustomEndpoint               pulumi.StringPtrOutput `pulumi:"pubsubCustomEndpoint"`
	PubsubLiteCustomEndpoint           pulumi.StringPtrOutput `pulumi:"pubsubLiteCustomEndpoint"`
	RedisCustomEndpoint                pulumi.StringPtrOutput `pulumi:"redisCustomEndpoint"`
	Region                             pulumi.StringPtrOutput `pulumi:"region"`
	RequestReason                      pulumi.StringPtrOutput `pulumi:"requestReason"`
	RequestTimeout                     pulumi.StringPtrOutput `pulumi:"requestTimeout"`
	ResourceManagerCustomEndpoint      pulumi.StringPtrOutput `pulumi:"resourceManagerCustomEndpoint"`
	ResourceManagerV2CustomEndpoint    pulumi.StringPtrOutput `pulumi:"resourceManagerV2CustomEndpoint"`
	RuntimeConfigCustomEndpoint        pulumi.StringPtrOutput `pulumi:"runtimeConfigCustomEndpoint"`
	RuntimeconfigCustomEndpoint        pulumi.StringPtrOutput `pulumi:"runtimeconfigCustomEndpoint"`
	SecretManagerCustomEndpoint        pulumi.StringPtrOutput `pulumi:"secretManagerCustomEndpoint"`
	SecurityCenterCustomEndpoint       pulumi.StringPtrOutput `pulumi:"securityCenterCustomEndpoint"`
	SecurityScannerCustomEndpoint      pulumi.StringPtrOutput `pulumi:"securityScannerCustomEndpoint"`
	ServiceDirectoryCustomEndpoint     pulumi.StringPtrOutput `pulumi:"serviceDirectoryCustomEndpoint"`
	ServiceManagementCustomEndpoint    pulumi.StringPtrOutput `pulumi:"serviceManagementCustomEndpoint"`
	ServiceNetworkingCustomEndpoint    pulumi.StringPtrOutput `pulumi:"serviceNetworkingCustomEndpoint"`
	ServiceUsageCustomEndpoint         pulumi.StringPtrOutput `pulumi:"serviceUsageCustomEndpoint"`
	SourceRepoCustomEndpoint           pulumi.StringPtrOutput `pulumi:"sourceRepoCustomEndpoint"`
	SpannerCustomEndpoint              pulumi.StringPtrOutput `pulumi:"spannerCustomEndpoint"`
	SqlCustomEndpoint                  pulumi.StringPtrOutput `pulumi:"sqlCustomEndpoint"`
	StorageCustomEndpoint              pulumi.StringPtrOutput `pulumi:"storageCustomEndpoint"`
	StorageTransferCustomEndpoint      pulumi.StringPtrOutput `pulumi:"storageTransferCustomEndpoint"`
	TagsCustomEndpoint                 pulumi.StringPtrOutput `pulumi:"tagsCustomEndpoint"`
	TpuCustomEndpoint                  pulumi.StringPtrOutput `pulumi:"tpuCustomEndpoint"`
	VertexAiCustomEndpoint             pulumi.StringPtrOutput `pulumi:"vertexAiCustomEndpoint"`
	VpcAccessCustomEndpoint            pulumi.StringPtrOutput `pulumi:"vpcAccessCustomEndpoint"`
	WorkflowsCustomEndpoint            pulumi.StringPtrOutput `pulumi:"workflowsCustomEndpoint"`
	Zone                               pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.Project == nil {
		args.Project = pulumi.StringPtr(getEnvOrDefault("", nil, "GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT").(string))
	}
	if args.Region == nil {
		args.Region = pulumi.StringPtr(getEnvOrDefault("", nil, "GOOGLE_REGION", "GCLOUD_REGION", "CLOUDSDK_COMPUTE_REGION").(string))
	}
	if args.Zone == nil {
		args.Zone = pulumi.StringPtr(getEnvOrDefault("", nil, "GOOGLE_ZONE", "GCLOUD_ZONE", "CLOUDSDK_COMPUTE_ZONE").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:gcp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	AccessApprovalCustomEndpoint       *string           `pulumi:"accessApprovalCustomEndpoint"`
	AccessContextManagerCustomEndpoint *string           `pulumi:"accessContextManagerCustomEndpoint"`
	AccessToken                        *string           `pulumi:"accessToken"`
	ActiveDirectoryCustomEndpoint      *string           `pulumi:"activeDirectoryCustomEndpoint"`
	ApiGatewayCustomEndpoint           *string           `pulumi:"apiGatewayCustomEndpoint"`
	ApigeeCustomEndpoint               *string           `pulumi:"apigeeCustomEndpoint"`
	AppEngineCustomEndpoint            *string           `pulumi:"appEngineCustomEndpoint"`
	ArtifactRegistryCustomEndpoint     *string           `pulumi:"artifactRegistryCustomEndpoint"`
	AssuredWorkloadsCustomEndpoint     *string           `pulumi:"assuredWorkloadsCustomEndpoint"`
	Batching                           *ProviderBatching `pulumi:"batching"`
	BigQueryCustomEndpoint             *string           `pulumi:"bigQueryCustomEndpoint"`
	BigqueryConnectionCustomEndpoint   *string           `pulumi:"bigqueryConnectionCustomEndpoint"`
	BigqueryDataTransferCustomEndpoint *string           `pulumi:"bigqueryDataTransferCustomEndpoint"`
	BigqueryReservationCustomEndpoint  *string           `pulumi:"bigqueryReservationCustomEndpoint"`
	BigtableCustomEndpoint             *string           `pulumi:"bigtableCustomEndpoint"`
	BillingCustomEndpoint              *string           `pulumi:"billingCustomEndpoint"`
	BillingProject                     *string           `pulumi:"billingProject"`
	BinaryAuthorizationCustomEndpoint  *string           `pulumi:"binaryAuthorizationCustomEndpoint"`
	CloudAssetCustomEndpoint           *string           `pulumi:"cloudAssetCustomEndpoint"`
	CloudBillingCustomEndpoint         *string           `pulumi:"cloudBillingCustomEndpoint"`
	CloudBuildCustomEndpoint           *string           `pulumi:"cloudBuildCustomEndpoint"`
	CloudBuildWorkerPoolCustomEndpoint *string           `pulumi:"cloudBuildWorkerPoolCustomEndpoint"`
	CloudFunctionsCustomEndpoint       *string           `pulumi:"cloudFunctionsCustomEndpoint"`
	CloudIdentityCustomEndpoint        *string           `pulumi:"cloudIdentityCustomEndpoint"`
	CloudIotCustomEndpoint             *string           `pulumi:"cloudIotCustomEndpoint"`
	CloudRunCustomEndpoint             *string           `pulumi:"cloudRunCustomEndpoint"`
	CloudSchedulerCustomEndpoint       *string           `pulumi:"cloudSchedulerCustomEndpoint"`
	CloudTasksCustomEndpoint           *string           `pulumi:"cloudTasksCustomEndpoint"`
	ComposerCustomEndpoint             *string           `pulumi:"composerCustomEndpoint"`
	ComputeBetaCustomEndpoint          *string           `pulumi:"computeBetaCustomEndpoint"`
	ComputeCustomEndpoint              *string           `pulumi:"computeCustomEndpoint"`
	ContainerAnalysisCustomEndpoint    *string           `pulumi:"containerAnalysisCustomEndpoint"`
	ContainerBetaCustomEndpoint        *string           `pulumi:"containerBetaCustomEndpoint"`
	ContainerCustomEndpoint            *string           `pulumi:"containerCustomEndpoint"`
	Credentials                        *string           `pulumi:"credentials"`
	DataCatalogCustomEndpoint          *string           `pulumi:"dataCatalogCustomEndpoint"`
	DataFusionCustomEndpoint           *string           `pulumi:"dataFusionCustomEndpoint"`
	DataLossPreventionCustomEndpoint   *string           `pulumi:"dataLossPreventionCustomEndpoint"`
	DataflowCustomEndpoint             *string           `pulumi:"dataflowCustomEndpoint"`
	DataprocBetaCustomEndpoint         *string           `pulumi:"dataprocBetaCustomEndpoint"`
	DataprocCustomEndpoint             *string           `pulumi:"dataprocCustomEndpoint"`
	DataprocMetastoreCustomEndpoint    *string           `pulumi:"dataprocMetastoreCustomEndpoint"`
	DatastoreCustomEndpoint            *string           `pulumi:"datastoreCustomEndpoint"`
	DeploymentManagerCustomEndpoint    *string           `pulumi:"deploymentManagerCustomEndpoint"`
	DialogflowCustomEndpoint           *string           `pulumi:"dialogflowCustomEndpoint"`
	DialogflowCxCustomEndpoint         *string           `pulumi:"dialogflowCxCustomEndpoint"`
	DisableGooglePartnerName           *bool             `pulumi:"disableGooglePartnerName"`
	DnsCustomEndpoint                  *string           `pulumi:"dnsCustomEndpoint"`
	EssentialContactsCustomEndpoint    *string           `pulumi:"essentialContactsCustomEndpoint"`
	EventarcCustomEndpoint             *string           `pulumi:"eventarcCustomEndpoint"`
	FilestoreCustomEndpoint            *string           `pulumi:"filestoreCustomEndpoint"`
	FirebaseCustomEndpoint             *string           `pulumi:"firebaseCustomEndpoint"`
	FirestoreCustomEndpoint            *string           `pulumi:"firestoreCustomEndpoint"`
	GameServicesCustomEndpoint         *string           `pulumi:"gameServicesCustomEndpoint"`
	GkeHubCustomEndpoint               *string           `pulumi:"gkeHubCustomEndpoint"`
	GkehubFeatureCustomEndpoint        *string           `pulumi:"gkehubFeatureCustomEndpoint"`
	GooglePartnerName                  *string           `pulumi:"googlePartnerName"`
	HealthcareCustomEndpoint           *string           `pulumi:"healthcareCustomEndpoint"`
	IamBetaCustomEndpoint              *string           `pulumi:"iamBetaCustomEndpoint"`
	IamCredentialsCustomEndpoint       *string           `pulumi:"iamCredentialsCustomEndpoint"`
	IamCustomEndpoint                  *string           `pulumi:"iamCustomEndpoint"`
	IapCustomEndpoint                  *string           `pulumi:"iapCustomEndpoint"`
	IdentityPlatformCustomEndpoint     *string           `pulumi:"identityPlatformCustomEndpoint"`
	ImpersonateServiceAccount          *string           `pulumi:"impersonateServiceAccount"`
	ImpersonateServiceAccountDelegates []string          `pulumi:"impersonateServiceAccountDelegates"`
	KmsCustomEndpoint                  *string           `pulumi:"kmsCustomEndpoint"`
	LoggingCustomEndpoint              *string           `pulumi:"loggingCustomEndpoint"`
	MemcacheCustomEndpoint             *string           `pulumi:"memcacheCustomEndpoint"`
	MlEngineCustomEndpoint             *string           `pulumi:"mlEngineCustomEndpoint"`
	MonitoringCustomEndpoint           *string           `pulumi:"monitoringCustomEndpoint"`
	NetworkManagementCustomEndpoint    *string           `pulumi:"networkManagementCustomEndpoint"`
	NetworkServicesCustomEndpoint      *string           `pulumi:"networkServicesCustomEndpoint"`
	NotebooksCustomEndpoint            *string           `pulumi:"notebooksCustomEndpoint"`
	OsConfigCustomEndpoint             *string           `pulumi:"osConfigCustomEndpoint"`
	OsLoginCustomEndpoint              *string           `pulumi:"osLoginCustomEndpoint"`
	PrivatecaCustomEndpoint            *string           `pulumi:"privatecaCustomEndpoint"`
	Project                            *string           `pulumi:"project"`
	PubsubCustomEndpoint               *string           `pulumi:"pubsubCustomEndpoint"`
	PubsubLiteCustomEndpoint           *string           `pulumi:"pubsubLiteCustomEndpoint"`
	RedisCustomEndpoint                *string           `pulumi:"redisCustomEndpoint"`
	Region                             *string           `pulumi:"region"`
	RequestReason                      *string           `pulumi:"requestReason"`
	RequestTimeout                     *string           `pulumi:"requestTimeout"`
	ResourceManagerCustomEndpoint      *string           `pulumi:"resourceManagerCustomEndpoint"`
	ResourceManagerV2CustomEndpoint    *string           `pulumi:"resourceManagerV2CustomEndpoint"`
	RuntimeConfigCustomEndpoint        *string           `pulumi:"runtimeConfigCustomEndpoint"`
	RuntimeconfigCustomEndpoint        *string           `pulumi:"runtimeconfigCustomEndpoint"`
	Scopes                             []string          `pulumi:"scopes"`
	SecretManagerCustomEndpoint        *string           `pulumi:"secretManagerCustomEndpoint"`
	SecurityCenterCustomEndpoint       *string           `pulumi:"securityCenterCustomEndpoint"`
	SecurityScannerCustomEndpoint      *string           `pulumi:"securityScannerCustomEndpoint"`
	ServiceDirectoryCustomEndpoint     *string           `pulumi:"serviceDirectoryCustomEndpoint"`
	ServiceManagementCustomEndpoint    *string           `pulumi:"serviceManagementCustomEndpoint"`
	ServiceNetworkingCustomEndpoint    *string           `pulumi:"serviceNetworkingCustomEndpoint"`
	ServiceUsageCustomEndpoint         *string           `pulumi:"serviceUsageCustomEndpoint"`
	SourceRepoCustomEndpoint           *string           `pulumi:"sourceRepoCustomEndpoint"`
	SpannerCustomEndpoint              *string           `pulumi:"spannerCustomEndpoint"`
	SqlCustomEndpoint                  *string           `pulumi:"sqlCustomEndpoint"`
	StorageCustomEndpoint              *string           `pulumi:"storageCustomEndpoint"`
	StorageTransferCustomEndpoint      *string           `pulumi:"storageTransferCustomEndpoint"`
	TagsCustomEndpoint                 *string           `pulumi:"tagsCustomEndpoint"`
	TpuCustomEndpoint                  *string           `pulumi:"tpuCustomEndpoint"`
	UserProjectOverride                *bool             `pulumi:"userProjectOverride"`
	VertexAiCustomEndpoint             *string           `pulumi:"vertexAiCustomEndpoint"`
	VpcAccessCustomEndpoint            *string           `pulumi:"vpcAccessCustomEndpoint"`
	WorkflowsCustomEndpoint            *string           `pulumi:"workflowsCustomEndpoint"`
	Zone                               *string           `pulumi:"zone"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	AccessApprovalCustomEndpoint       pulumi.StringPtrInput
	AccessContextManagerCustomEndpoint pulumi.StringPtrInput
	AccessToken                        pulumi.StringPtrInput
	ActiveDirectoryCustomEndpoint      pulumi.StringPtrInput
	ApiGatewayCustomEndpoint           pulumi.StringPtrInput
	ApigeeCustomEndpoint               pulumi.StringPtrInput
	AppEngineCustomEndpoint            pulumi.StringPtrInput
	ArtifactRegistryCustomEndpoint     pulumi.StringPtrInput
	AssuredWorkloadsCustomEndpoint     pulumi.StringPtrInput
	Batching                           ProviderBatchingPtrInput
	BigQueryCustomEndpoint             pulumi.StringPtrInput
	BigqueryConnectionCustomEndpoint   pulumi.StringPtrInput
	BigqueryDataTransferCustomEndpoint pulumi.StringPtrInput
	BigqueryReservationCustomEndpoint  pulumi.StringPtrInput
	BigtableCustomEndpoint             pulumi.StringPtrInput
	BillingCustomEndpoint              pulumi.StringPtrInput
	BillingProject                     pulumi.StringPtrInput
	BinaryAuthorizationCustomEndpoint  pulumi.StringPtrInput
	CloudAssetCustomEndpoint           pulumi.StringPtrInput
	CloudBillingCustomEndpoint         pulumi.StringPtrInput
	CloudBuildCustomEndpoint           pulumi.StringPtrInput
	CloudBuildWorkerPoolCustomEndpoint pulumi.StringPtrInput
	CloudFunctionsCustomEndpoint       pulumi.StringPtrInput
	CloudIdentityCustomEndpoint        pulumi.StringPtrInput
	CloudIotCustomEndpoint             pulumi.StringPtrInput
	CloudRunCustomEndpoint             pulumi.StringPtrInput
	CloudSchedulerCustomEndpoint       pulumi.StringPtrInput
	CloudTasksCustomEndpoint           pulumi.StringPtrInput
	ComposerCustomEndpoint             pulumi.StringPtrInput
	ComputeBetaCustomEndpoint          pulumi.StringPtrInput
	ComputeCustomEndpoint              pulumi.StringPtrInput
	ContainerAnalysisCustomEndpoint    pulumi.StringPtrInput
	ContainerBetaCustomEndpoint        pulumi.StringPtrInput
	ContainerCustomEndpoint            pulumi.StringPtrInput
	Credentials                        pulumi.StringPtrInput
	DataCatalogCustomEndpoint          pulumi.StringPtrInput
	DataFusionCustomEndpoint           pulumi.StringPtrInput
	DataLossPreventionCustomEndpoint   pulumi.StringPtrInput
	DataflowCustomEndpoint             pulumi.StringPtrInput
	DataprocBetaCustomEndpoint         pulumi.StringPtrInput
	DataprocCustomEndpoint             pulumi.StringPtrInput
	DataprocMetastoreCustomEndpoint    pulumi.StringPtrInput
	DatastoreCustomEndpoint            pulumi.StringPtrInput
	DeploymentManagerCustomEndpoint    pulumi.StringPtrInput
	DialogflowCustomEndpoint           pulumi.StringPtrInput
	DialogflowCxCustomEndpoint         pulumi.StringPtrInput
	DisableGooglePartnerName           pulumi.BoolPtrInput
	DnsCustomEndpoint                  pulumi.StringPtrInput
	EssentialContactsCustomEndpoint    pulumi.StringPtrInput
	EventarcCustomEndpoint             pulumi.StringPtrInput
	FilestoreCustomEndpoint            pulumi.StringPtrInput
	FirebaseCustomEndpoint             pulumi.StringPtrInput
	FirestoreCustomEndpoint            pulumi.StringPtrInput
	GameServicesCustomEndpoint         pulumi.StringPtrInput
	GkeHubCustomEndpoint               pulumi.StringPtrInput
	GkehubFeatureCustomEndpoint        pulumi.StringPtrInput
	GooglePartnerName                  pulumi.StringPtrInput
	HealthcareCustomEndpoint           pulumi.StringPtrInput
	IamBetaCustomEndpoint              pulumi.StringPtrInput
	IamCredentialsCustomEndpoint       pulumi.StringPtrInput
	IamCustomEndpoint                  pulumi.StringPtrInput
	IapCustomEndpoint                  pulumi.StringPtrInput
	IdentityPlatformCustomEndpoint     pulumi.StringPtrInput
	ImpersonateServiceAccount          pulumi.StringPtrInput
	ImpersonateServiceAccountDelegates pulumi.StringArrayInput
	KmsCustomEndpoint                  pulumi.StringPtrInput
	LoggingCustomEndpoint              pulumi.StringPtrInput
	MemcacheCustomEndpoint             pulumi.StringPtrInput
	MlEngineCustomEndpoint             pulumi.StringPtrInput
	MonitoringCustomEndpoint           pulumi.StringPtrInput
	NetworkManagementCustomEndpoint    pulumi.StringPtrInput
	NetworkServicesCustomEndpoint      pulumi.StringPtrInput
	NotebooksCustomEndpoint            pulumi.StringPtrInput
	OsConfigCustomEndpoint             pulumi.StringPtrInput
	OsLoginCustomEndpoint              pulumi.StringPtrInput
	PrivatecaCustomEndpoint            pulumi.StringPtrInput
	Project                            pulumi.StringPtrInput
	PubsubCustomEndpoint               pulumi.StringPtrInput
	PubsubLiteCustomEndpoint           pulumi.StringPtrInput
	RedisCustomEndpoint                pulumi.StringPtrInput
	Region                             pulumi.StringPtrInput
	RequestReason                      pulumi.StringPtrInput
	RequestTimeout                     pulumi.StringPtrInput
	ResourceManagerCustomEndpoint      pulumi.StringPtrInput
	ResourceManagerV2CustomEndpoint    pulumi.StringPtrInput
	RuntimeConfigCustomEndpoint        pulumi.StringPtrInput
	RuntimeconfigCustomEndpoint        pulumi.StringPtrInput
	Scopes                             pulumi.StringArrayInput
	SecretManagerCustomEndpoint        pulumi.StringPtrInput
	SecurityCenterCustomEndpoint       pulumi.StringPtrInput
	SecurityScannerCustomEndpoint      pulumi.StringPtrInput
	ServiceDirectoryCustomEndpoint     pulumi.StringPtrInput
	ServiceManagementCustomEndpoint    pulumi.StringPtrInput
	ServiceNetworkingCustomEndpoint    pulumi.StringPtrInput
	ServiceUsageCustomEndpoint         pulumi.StringPtrInput
	SourceRepoCustomEndpoint           pulumi.StringPtrInput
	SpannerCustomEndpoint              pulumi.StringPtrInput
	SqlCustomEndpoint                  pulumi.StringPtrInput
	StorageCustomEndpoint              pulumi.StringPtrInput
	StorageTransferCustomEndpoint      pulumi.StringPtrInput
	TagsCustomEndpoint                 pulumi.StringPtrInput
	TpuCustomEndpoint                  pulumi.StringPtrInput
	UserProjectOverride                pulumi.BoolPtrInput
	VertexAiCustomEndpoint             pulumi.StringPtrInput
	VpcAccessCustomEndpoint            pulumi.StringPtrInput
	WorkflowsCustomEndpoint            pulumi.StringPtrInput
	Zone                               pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *Provider) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderPtrInput interface {
	pulumi.Input

	ToProviderPtrOutput() ProviderPtrOutput
	ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput
}

type providerPtrType ProviderArgs

func (*providerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (i *providerPtrType) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *providerPtrType) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o.ToProviderPtrOutputWithContext(context.Background())
}

func (o ProviderOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Provider) *Provider {
		return &v
	}).(ProviderPtrOutput)
}

type ProviderPtrOutput struct{ *pulumi.OutputState }

func (ProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (o ProviderPtrOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) Elem() ProviderOutput {
	return o.ApplyT(func(v *Provider) Provider {
		if v != nil {
			return *v
		}
		var ret Provider
		return ret
	}).(ProviderOutput)
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderPtrOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure connection to SDN Connector.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewSystemSdnConnector(ctx, "trname", &fortios.SystemSdnConnectorArgs{
// 			AzureRegion:    pulumi.String("global"),
// 			HaStatus:       pulumi.String("disable"),
// 			Password:       pulumi.String("deWdf321ds"),
// 			Server:         pulumi.String("1.1.1.1"),
// 			ServerPort:     pulumi.Int(3),
// 			Status:         pulumi.String("disable"),
// 			Type:           pulumi.String("aci"),
// 			UpdateInterval: pulumi.Int(60),
// 			UseMetadataIam: pulumi.String("disable"),
// 			Username:       pulumi.String("sg"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// System SdnConnector can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemSdnConnector:SystemSdnConnector labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemSdnConnector:SystemSdnConnector labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemSdnConnector struct {
	pulumi.CustomResourceState

	// AWS access key ID.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// IBM cloud API key or service ID API key.
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// Azure server region. Valid values: `global`, `china`, `germany`, `usgov`, `local`.
	AzureRegion pulumi.StringOutput `pulumi:"azureRegion"`
	// Azure client ID (application ID).
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Azure client secret (application key).
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// Compartment ID.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// Compute generation for IBM cloud infrastructure.
	ComputeGeneration pulumi.IntOutput `pulumi:"computeGeneration"`
	// Domain name.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Configure AWS external account list. The structure of `externalAccountList` block is documented below.
	ExternalAccountLists SystemSdnConnectorExternalAccountListArrayOutput `pulumi:"externalAccountLists"`
	// Configure GCP external IP. The structure of `externalIp` block is documented below.
	ExternalIps SystemSdnConnectorExternalIpArrayOutput `pulumi:"externalIps"`
	// Configure GCP forwarding rule. The structure of `forwardingRule` block is documented below.
	ForwardingRules SystemSdnConnectorForwardingRuleArrayOutput `pulumi:"forwardingRules"`
	// GCP project name.
	GcpProject pulumi.StringOutput `pulumi:"gcpProject"`
	// Configure GCP project list. The structure of `gcpProjectList` block is documented below.
	GcpProjectLists SystemSdnConnectorGcpProjectListArrayOutput `pulumi:"gcpProjectLists"`
	// Group name of computers.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// Enable/disable use for FortiGate HA service. Valid values: `disable`, `enable`.
	HaStatus pulumi.StringOutput `pulumi:"haStatus"`
	// IBM cloud region name.
	IbmRegion pulumi.StringOutput `pulumi:"ibmRegion"`
	// IBM cloud compute generation 1 region name. Valid values: `us-south`, `us-east`, `germany`, `great-britain`, `japan`, `australia`.
	IbmRegionGen1 pulumi.StringOutput `pulumi:"ibmRegionGen1"`
	// IBM cloud compute generation 2 region name. Valid values: `us-south`, `us-east`, `great-britain`.
	IbmRegionGen2 pulumi.StringOutput `pulumi:"ibmRegionGen2"`
	// Private key password.
	KeyPasswd pulumi.StringPtrOutput `pulumi:"keyPasswd"`
	// Azure Stack login endpoint.
	LoginEndpoint pulumi.StringOutput `pulumi:"loginEndpoint"`
	// GCP zone name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configure Azure network interface. The structure of `nic` block is documented below.
	Nics SystemSdnConnectorNicArrayOutput `pulumi:"nics"`
	// OCI certificate.
	OciCert pulumi.StringOutput `pulumi:"ociCert"`
	// OCI pubkey fingerprint.
	OciFingerprint pulumi.StringOutput `pulumi:"ociFingerprint"`
	// OCI server region.
	OciRegion pulumi.StringOutput `pulumi:"ociRegion"`
	// OCI region type. Valid values: `commercial`, `government`.
	OciRegionType pulumi.StringOutput `pulumi:"ociRegionType"`
	// Password of the remote SDN connector as login credentials.
	Password pulumi.StringOutput `pulumi:"password"`
	// Private key of GCP service account.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// AWS region name.
	Region pulumi.StringOutput `pulumi:"region"`
	// Resource group of Azure route table.
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// Azure Stack resource URL.
	ResourceUrl pulumi.StringOutput `pulumi:"resourceUrl"`
	// Configure Azure route table. The structure of `routeTable` block is documented below.
	RouteTables SystemSdnConnectorRouteTableArrayOutput `pulumi:"routeTables"`
	// Configure Azure route. The structure of `route` block is documented below.
	Routes SystemSdnConnectorRouteArrayOutput `pulumi:"routes"`
	// AWS secret access key.
	SecretKey pulumi.StringPtrOutput `pulumi:"secretKey"`
	// Secret token of Kubernetes service account.
	SecretToken pulumi.StringOutput `pulumi:"secretToken"`
	// Server address of the remote SDN connector.
	Server pulumi.StringOutput `pulumi:"server"`
	// Server address list of the remote SDN connector. The structure of `serverList` block is documented below.
	ServerLists SystemSdnConnectorServerListArrayOutput `pulumi:"serverLists"`
	// Port number of the remote SDN connector.
	ServerPort pulumi.IntOutput `pulumi:"serverPort"`
	// GCP service account email.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// Enable/disable connection to the remote SDN connector. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Subscription ID of Azure route table.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// Tenant ID (directory ID).
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Type of SDN connector.
	Type pulumi.StringOutput `pulumi:"type"`
	// Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
	UpdateInterval pulumi.IntOutput `pulumi:"updateInterval"`
	// Enable/disable using IAM role from metadata to call API. Valid values: `disable`, `enable`.
	UseMetadataIam pulumi.StringOutput `pulumi:"useMetadataIam"`
	// User ID.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// Username of the remote SDN connector as login credentials.
	Username pulumi.StringOutput `pulumi:"username"`
	// vCenter server password for NSX quarantine.
	VcenterPassword pulumi.StringPtrOutput `pulumi:"vcenterPassword"`
	// vCenter server address for NSX quarantine.
	VcenterServer pulumi.StringOutput `pulumi:"vcenterServer"`
	// vCenter server username for NSX quarantine.
	VcenterUsername pulumi.StringOutput `pulumi:"vcenterUsername"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable server certificate verification. Valid values: `disable`, `enable`.
	VerifyCertificate pulumi.StringOutput `pulumi:"verifyCertificate"`
	// AWS VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSystemSdnConnector registers a new resource with the given unique name, arguments, and options.
func NewSystemSdnConnector(ctx *pulumi.Context,
	name string, args *SystemSdnConnectorArgs, opts ...pulumi.ResourceOption) (*SystemSdnConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSdnConnector
	err := ctx.RegisterResource("fortios:index/systemSdnConnector:SystemSdnConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSdnConnector gets an existing SystemSdnConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSdnConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSdnConnectorState, opts ...pulumi.ResourceOption) (*SystemSdnConnector, error) {
	var resource SystemSdnConnector
	err := ctx.ReadResource("fortios:index/systemSdnConnector:SystemSdnConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSdnConnector resources.
type systemSdnConnectorState struct {
	// AWS access key ID.
	AccessKey *string `pulumi:"accessKey"`
	// IBM cloud API key or service ID API key.
	ApiKey *string `pulumi:"apiKey"`
	// Azure server region. Valid values: `global`, `china`, `germany`, `usgov`, `local`.
	AzureRegion *string `pulumi:"azureRegion"`
	// Azure client ID (application ID).
	ClientId *string `pulumi:"clientId"`
	// Azure client secret (application key).
	ClientSecret *string `pulumi:"clientSecret"`
	// Compartment ID.
	CompartmentId *string `pulumi:"compartmentId"`
	// Compute generation for IBM cloud infrastructure.
	ComputeGeneration *int `pulumi:"computeGeneration"`
	// Domain name.
	Domain *string `pulumi:"domain"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure AWS external account list. The structure of `externalAccountList` block is documented below.
	ExternalAccountLists []SystemSdnConnectorExternalAccountList `pulumi:"externalAccountLists"`
	// Configure GCP external IP. The structure of `externalIp` block is documented below.
	ExternalIps []SystemSdnConnectorExternalIp `pulumi:"externalIps"`
	// Configure GCP forwarding rule. The structure of `forwardingRule` block is documented below.
	ForwardingRules []SystemSdnConnectorForwardingRule `pulumi:"forwardingRules"`
	// GCP project name.
	GcpProject *string `pulumi:"gcpProject"`
	// Configure GCP project list. The structure of `gcpProjectList` block is documented below.
	GcpProjectLists []SystemSdnConnectorGcpProjectList `pulumi:"gcpProjectLists"`
	// Group name of computers.
	GroupName *string `pulumi:"groupName"`
	// Enable/disable use for FortiGate HA service. Valid values: `disable`, `enable`.
	HaStatus *string `pulumi:"haStatus"`
	// IBM cloud region name.
	IbmRegion *string `pulumi:"ibmRegion"`
	// IBM cloud compute generation 1 region name. Valid values: `us-south`, `us-east`, `germany`, `great-britain`, `japan`, `australia`.
	IbmRegionGen1 *string `pulumi:"ibmRegionGen1"`
	// IBM cloud compute generation 2 region name. Valid values: `us-south`, `us-east`, `great-britain`.
	IbmRegionGen2 *string `pulumi:"ibmRegionGen2"`
	// Private key password.
	KeyPasswd *string `pulumi:"keyPasswd"`
	// Azure Stack login endpoint.
	LoginEndpoint *string `pulumi:"loginEndpoint"`
	// GCP zone name.
	Name *string `pulumi:"name"`
	// Configure Azure network interface. The structure of `nic` block is documented below.
	Nics []SystemSdnConnectorNic `pulumi:"nics"`
	// OCI certificate.
	OciCert *string `pulumi:"ociCert"`
	// OCI pubkey fingerprint.
	OciFingerprint *string `pulumi:"ociFingerprint"`
	// OCI server region.
	OciRegion *string `pulumi:"ociRegion"`
	// OCI region type. Valid values: `commercial`, `government`.
	OciRegionType *string `pulumi:"ociRegionType"`
	// Password of the remote SDN connector as login credentials.
	Password *string `pulumi:"password"`
	// Private key of GCP service account.
	PrivateKey *string `pulumi:"privateKey"`
	// AWS region name.
	Region *string `pulumi:"region"`
	// Resource group of Azure route table.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Azure Stack resource URL.
	ResourceUrl *string `pulumi:"resourceUrl"`
	// Configure Azure route table. The structure of `routeTable` block is documented below.
	RouteTables []SystemSdnConnectorRouteTable `pulumi:"routeTables"`
	// Configure Azure route. The structure of `route` block is documented below.
	Routes []SystemSdnConnectorRoute `pulumi:"routes"`
	// AWS secret access key.
	SecretKey *string `pulumi:"secretKey"`
	// Secret token of Kubernetes service account.
	SecretToken *string `pulumi:"secretToken"`
	// Server address of the remote SDN connector.
	Server *string `pulumi:"server"`
	// Server address list of the remote SDN connector. The structure of `serverList` block is documented below.
	ServerLists []SystemSdnConnectorServerList `pulumi:"serverLists"`
	// Port number of the remote SDN connector.
	ServerPort *int `pulumi:"serverPort"`
	// GCP service account email.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Enable/disable connection to the remote SDN connector. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Subscription ID of Azure route table.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Tenant ID (directory ID).
	TenantId *string `pulumi:"tenantId"`
	// Type of SDN connector.
	Type *string `pulumi:"type"`
	// Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
	UpdateInterval *int `pulumi:"updateInterval"`
	// Enable/disable using IAM role from metadata to call API. Valid values: `disable`, `enable`.
	UseMetadataIam *string `pulumi:"useMetadataIam"`
	// User ID.
	UserId *string `pulumi:"userId"`
	// Username of the remote SDN connector as login credentials.
	Username *string `pulumi:"username"`
	// vCenter server password for NSX quarantine.
	VcenterPassword *string `pulumi:"vcenterPassword"`
	// vCenter server address for NSX quarantine.
	VcenterServer *string `pulumi:"vcenterServer"`
	// vCenter server username for NSX quarantine.
	VcenterUsername *string `pulumi:"vcenterUsername"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable server certificate verification. Valid values: `disable`, `enable`.
	VerifyCertificate *string `pulumi:"verifyCertificate"`
	// AWS VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type SystemSdnConnectorState struct {
	// AWS access key ID.
	AccessKey pulumi.StringPtrInput
	// IBM cloud API key or service ID API key.
	ApiKey pulumi.StringPtrInput
	// Azure server region. Valid values: `global`, `china`, `germany`, `usgov`, `local`.
	AzureRegion pulumi.StringPtrInput
	// Azure client ID (application ID).
	ClientId pulumi.StringPtrInput
	// Azure client secret (application key).
	ClientSecret pulumi.StringPtrInput
	// Compartment ID.
	CompartmentId pulumi.StringPtrInput
	// Compute generation for IBM cloud infrastructure.
	ComputeGeneration pulumi.IntPtrInput
	// Domain name.
	Domain pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure AWS external account list. The structure of `externalAccountList` block is documented below.
	ExternalAccountLists SystemSdnConnectorExternalAccountListArrayInput
	// Configure GCP external IP. The structure of `externalIp` block is documented below.
	ExternalIps SystemSdnConnectorExternalIpArrayInput
	// Configure GCP forwarding rule. The structure of `forwardingRule` block is documented below.
	ForwardingRules SystemSdnConnectorForwardingRuleArrayInput
	// GCP project name.
	GcpProject pulumi.StringPtrInput
	// Configure GCP project list. The structure of `gcpProjectList` block is documented below.
	GcpProjectLists SystemSdnConnectorGcpProjectListArrayInput
	// Group name of computers.
	GroupName pulumi.StringPtrInput
	// Enable/disable use for FortiGate HA service. Valid values: `disable`, `enable`.
	HaStatus pulumi.StringPtrInput
	// IBM cloud region name.
	IbmRegion pulumi.StringPtrInput
	// IBM cloud compute generation 1 region name. Valid values: `us-south`, `us-east`, `germany`, `great-britain`, `japan`, `australia`.
	IbmRegionGen1 pulumi.StringPtrInput
	// IBM cloud compute generation 2 region name. Valid values: `us-south`, `us-east`, `great-britain`.
	IbmRegionGen2 pulumi.StringPtrInput
	// Private key password.
	KeyPasswd pulumi.StringPtrInput
	// Azure Stack login endpoint.
	LoginEndpoint pulumi.StringPtrInput
	// GCP zone name.
	Name pulumi.StringPtrInput
	// Configure Azure network interface. The structure of `nic` block is documented below.
	Nics SystemSdnConnectorNicArrayInput
	// OCI certificate.
	OciCert pulumi.StringPtrInput
	// OCI pubkey fingerprint.
	OciFingerprint pulumi.StringPtrInput
	// OCI server region.
	OciRegion pulumi.StringPtrInput
	// OCI region type. Valid values: `commercial`, `government`.
	OciRegionType pulumi.StringPtrInput
	// Password of the remote SDN connector as login credentials.
	Password pulumi.StringPtrInput
	// Private key of GCP service account.
	PrivateKey pulumi.StringPtrInput
	// AWS region name.
	Region pulumi.StringPtrInput
	// Resource group of Azure route table.
	ResourceGroup pulumi.StringPtrInput
	// Azure Stack resource URL.
	ResourceUrl pulumi.StringPtrInput
	// Configure Azure route table. The structure of `routeTable` block is documented below.
	RouteTables SystemSdnConnectorRouteTableArrayInput
	// Configure Azure route. The structure of `route` block is documented below.
	Routes SystemSdnConnectorRouteArrayInput
	// AWS secret access key.
	SecretKey pulumi.StringPtrInput
	// Secret token of Kubernetes service account.
	SecretToken pulumi.StringPtrInput
	// Server address of the remote SDN connector.
	Server pulumi.StringPtrInput
	// Server address list of the remote SDN connector. The structure of `serverList` block is documented below.
	ServerLists SystemSdnConnectorServerListArrayInput
	// Port number of the remote SDN connector.
	ServerPort pulumi.IntPtrInput
	// GCP service account email.
	ServiceAccount pulumi.StringPtrInput
	// Enable/disable connection to the remote SDN connector. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Subscription ID of Azure route table.
	SubscriptionId pulumi.StringPtrInput
	// Tenant ID (directory ID).
	TenantId pulumi.StringPtrInput
	// Type of SDN connector.
	Type pulumi.StringPtrInput
	// Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
	UpdateInterval pulumi.IntPtrInput
	// Enable/disable using IAM role from metadata to call API. Valid values: `disable`, `enable`.
	UseMetadataIam pulumi.StringPtrInput
	// User ID.
	UserId pulumi.StringPtrInput
	// Username of the remote SDN connector as login credentials.
	Username pulumi.StringPtrInput
	// vCenter server password for NSX quarantine.
	VcenterPassword pulumi.StringPtrInput
	// vCenter server address for NSX quarantine.
	VcenterServer pulumi.StringPtrInput
	// vCenter server username for NSX quarantine.
	VcenterUsername pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable server certificate verification. Valid values: `disable`, `enable`.
	VerifyCertificate pulumi.StringPtrInput
	// AWS VPC ID.
	VpcId pulumi.StringPtrInput
}

func (SystemSdnConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSdnConnectorState)(nil)).Elem()
}

type systemSdnConnectorArgs struct {
	// AWS access key ID.
	AccessKey *string `pulumi:"accessKey"`
	// IBM cloud API key or service ID API key.
	ApiKey *string `pulumi:"apiKey"`
	// Azure server region. Valid values: `global`, `china`, `germany`, `usgov`, `local`.
	AzureRegion *string `pulumi:"azureRegion"`
	// Azure client ID (application ID).
	ClientId *string `pulumi:"clientId"`
	// Azure client secret (application key).
	ClientSecret *string `pulumi:"clientSecret"`
	// Compartment ID.
	CompartmentId *string `pulumi:"compartmentId"`
	// Compute generation for IBM cloud infrastructure.
	ComputeGeneration *int `pulumi:"computeGeneration"`
	// Domain name.
	Domain *string `pulumi:"domain"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure AWS external account list. The structure of `externalAccountList` block is documented below.
	ExternalAccountLists []SystemSdnConnectorExternalAccountList `pulumi:"externalAccountLists"`
	// Configure GCP external IP. The structure of `externalIp` block is documented below.
	ExternalIps []SystemSdnConnectorExternalIp `pulumi:"externalIps"`
	// Configure GCP forwarding rule. The structure of `forwardingRule` block is documented below.
	ForwardingRules []SystemSdnConnectorForwardingRule `pulumi:"forwardingRules"`
	// GCP project name.
	GcpProject *string `pulumi:"gcpProject"`
	// Configure GCP project list. The structure of `gcpProjectList` block is documented below.
	GcpProjectLists []SystemSdnConnectorGcpProjectList `pulumi:"gcpProjectLists"`
	// Group name of computers.
	GroupName *string `pulumi:"groupName"`
	// Enable/disable use for FortiGate HA service. Valid values: `disable`, `enable`.
	HaStatus *string `pulumi:"haStatus"`
	// IBM cloud region name.
	IbmRegion *string `pulumi:"ibmRegion"`
	// IBM cloud compute generation 1 region name. Valid values: `us-south`, `us-east`, `germany`, `great-britain`, `japan`, `australia`.
	IbmRegionGen1 *string `pulumi:"ibmRegionGen1"`
	// IBM cloud compute generation 2 region name. Valid values: `us-south`, `us-east`, `great-britain`.
	IbmRegionGen2 *string `pulumi:"ibmRegionGen2"`
	// Private key password.
	KeyPasswd *string `pulumi:"keyPasswd"`
	// Azure Stack login endpoint.
	LoginEndpoint *string `pulumi:"loginEndpoint"`
	// GCP zone name.
	Name *string `pulumi:"name"`
	// Configure Azure network interface. The structure of `nic` block is documented below.
	Nics []SystemSdnConnectorNic `pulumi:"nics"`
	// OCI certificate.
	OciCert *string `pulumi:"ociCert"`
	// OCI pubkey fingerprint.
	OciFingerprint *string `pulumi:"ociFingerprint"`
	// OCI server region.
	OciRegion *string `pulumi:"ociRegion"`
	// OCI region type. Valid values: `commercial`, `government`.
	OciRegionType *string `pulumi:"ociRegionType"`
	// Password of the remote SDN connector as login credentials.
	Password *string `pulumi:"password"`
	// Private key of GCP service account.
	PrivateKey *string `pulumi:"privateKey"`
	// AWS region name.
	Region *string `pulumi:"region"`
	// Resource group of Azure route table.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Azure Stack resource URL.
	ResourceUrl *string `pulumi:"resourceUrl"`
	// Configure Azure route table. The structure of `routeTable` block is documented below.
	RouteTables []SystemSdnConnectorRouteTable `pulumi:"routeTables"`
	// Configure Azure route. The structure of `route` block is documented below.
	Routes []SystemSdnConnectorRoute `pulumi:"routes"`
	// AWS secret access key.
	SecretKey *string `pulumi:"secretKey"`
	// Secret token of Kubernetes service account.
	SecretToken *string `pulumi:"secretToken"`
	// Server address of the remote SDN connector.
	Server *string `pulumi:"server"`
	// Server address list of the remote SDN connector. The structure of `serverList` block is documented below.
	ServerLists []SystemSdnConnectorServerList `pulumi:"serverLists"`
	// Port number of the remote SDN connector.
	ServerPort *int `pulumi:"serverPort"`
	// GCP service account email.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Enable/disable connection to the remote SDN connector. Valid values: `disable`, `enable`.
	Status string `pulumi:"status"`
	// Subscription ID of Azure route table.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Tenant ID (directory ID).
	TenantId *string `pulumi:"tenantId"`
	// Type of SDN connector.
	Type string `pulumi:"type"`
	// Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
	UpdateInterval *int `pulumi:"updateInterval"`
	// Enable/disable using IAM role from metadata to call API. Valid values: `disable`, `enable`.
	UseMetadataIam *string `pulumi:"useMetadataIam"`
	// User ID.
	UserId *string `pulumi:"userId"`
	// Username of the remote SDN connector as login credentials.
	Username *string `pulumi:"username"`
	// vCenter server password for NSX quarantine.
	VcenterPassword *string `pulumi:"vcenterPassword"`
	// vCenter server address for NSX quarantine.
	VcenterServer *string `pulumi:"vcenterServer"`
	// vCenter server username for NSX quarantine.
	VcenterUsername *string `pulumi:"vcenterUsername"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable server certificate verification. Valid values: `disable`, `enable`.
	VerifyCertificate *string `pulumi:"verifyCertificate"`
	// AWS VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a SystemSdnConnector resource.
type SystemSdnConnectorArgs struct {
	// AWS access key ID.
	AccessKey pulumi.StringPtrInput
	// IBM cloud API key or service ID API key.
	ApiKey pulumi.StringPtrInput
	// Azure server region. Valid values: `global`, `china`, `germany`, `usgov`, `local`.
	AzureRegion pulumi.StringPtrInput
	// Azure client ID (application ID).
	ClientId pulumi.StringPtrInput
	// Azure client secret (application key).
	ClientSecret pulumi.StringPtrInput
	// Compartment ID.
	CompartmentId pulumi.StringPtrInput
	// Compute generation for IBM cloud infrastructure.
	ComputeGeneration pulumi.IntPtrInput
	// Domain name.
	Domain pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure AWS external account list. The structure of `externalAccountList` block is documented below.
	ExternalAccountLists SystemSdnConnectorExternalAccountListArrayInput
	// Configure GCP external IP. The structure of `externalIp` block is documented below.
	ExternalIps SystemSdnConnectorExternalIpArrayInput
	// Configure GCP forwarding rule. The structure of `forwardingRule` block is documented below.
	ForwardingRules SystemSdnConnectorForwardingRuleArrayInput
	// GCP project name.
	GcpProject pulumi.StringPtrInput
	// Configure GCP project list. The structure of `gcpProjectList` block is documented below.
	GcpProjectLists SystemSdnConnectorGcpProjectListArrayInput
	// Group name of computers.
	GroupName pulumi.StringPtrInput
	// Enable/disable use for FortiGate HA service. Valid values: `disable`, `enable`.
	HaStatus pulumi.StringPtrInput
	// IBM cloud region name.
	IbmRegion pulumi.StringPtrInput
	// IBM cloud compute generation 1 region name. Valid values: `us-south`, `us-east`, `germany`, `great-britain`, `japan`, `australia`.
	IbmRegionGen1 pulumi.StringPtrInput
	// IBM cloud compute generation 2 region name. Valid values: `us-south`, `us-east`, `great-britain`.
	IbmRegionGen2 pulumi.StringPtrInput
	// Private key password.
	KeyPasswd pulumi.StringPtrInput
	// Azure Stack login endpoint.
	LoginEndpoint pulumi.StringPtrInput
	// GCP zone name.
	Name pulumi.StringPtrInput
	// Configure Azure network interface. The structure of `nic` block is documented below.
	Nics SystemSdnConnectorNicArrayInput
	// OCI certificate.
	OciCert pulumi.StringPtrInput
	// OCI pubkey fingerprint.
	OciFingerprint pulumi.StringPtrInput
	// OCI server region.
	OciRegion pulumi.StringPtrInput
	// OCI region type. Valid values: `commercial`, `government`.
	OciRegionType pulumi.StringPtrInput
	// Password of the remote SDN connector as login credentials.
	Password pulumi.StringPtrInput
	// Private key of GCP service account.
	PrivateKey pulumi.StringPtrInput
	// AWS region name.
	Region pulumi.StringPtrInput
	// Resource group of Azure route table.
	ResourceGroup pulumi.StringPtrInput
	// Azure Stack resource URL.
	ResourceUrl pulumi.StringPtrInput
	// Configure Azure route table. The structure of `routeTable` block is documented below.
	RouteTables SystemSdnConnectorRouteTableArrayInput
	// Configure Azure route. The structure of `route` block is documented below.
	Routes SystemSdnConnectorRouteArrayInput
	// AWS secret access key.
	SecretKey pulumi.StringPtrInput
	// Secret token of Kubernetes service account.
	SecretToken pulumi.StringPtrInput
	// Server address of the remote SDN connector.
	Server pulumi.StringPtrInput
	// Server address list of the remote SDN connector. The structure of `serverList` block is documented below.
	ServerLists SystemSdnConnectorServerListArrayInput
	// Port number of the remote SDN connector.
	ServerPort pulumi.IntPtrInput
	// GCP service account email.
	ServiceAccount pulumi.StringPtrInput
	// Enable/disable connection to the remote SDN connector. Valid values: `disable`, `enable`.
	Status pulumi.StringInput
	// Subscription ID of Azure route table.
	SubscriptionId pulumi.StringPtrInput
	// Tenant ID (directory ID).
	TenantId pulumi.StringPtrInput
	// Type of SDN connector.
	Type pulumi.StringInput
	// Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
	UpdateInterval pulumi.IntPtrInput
	// Enable/disable using IAM role from metadata to call API. Valid values: `disable`, `enable`.
	UseMetadataIam pulumi.StringPtrInput
	// User ID.
	UserId pulumi.StringPtrInput
	// Username of the remote SDN connector as login credentials.
	Username pulumi.StringPtrInput
	// vCenter server password for NSX quarantine.
	VcenterPassword pulumi.StringPtrInput
	// vCenter server address for NSX quarantine.
	VcenterServer pulumi.StringPtrInput
	// vCenter server username for NSX quarantine.
	VcenterUsername pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable server certificate verification. Valid values: `disable`, `enable`.
	VerifyCertificate pulumi.StringPtrInput
	// AWS VPC ID.
	VpcId pulumi.StringPtrInput
}

func (SystemSdnConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSdnConnectorArgs)(nil)).Elem()
}

type SystemSdnConnectorInput interface {
	pulumi.Input

	ToSystemSdnConnectorOutput() SystemSdnConnectorOutput
	ToSystemSdnConnectorOutputWithContext(ctx context.Context) SystemSdnConnectorOutput
}

func (*SystemSdnConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSdnConnector)(nil)).Elem()
}

func (i *SystemSdnConnector) ToSystemSdnConnectorOutput() SystemSdnConnectorOutput {
	return i.ToSystemSdnConnectorOutputWithContext(context.Background())
}

func (i *SystemSdnConnector) ToSystemSdnConnectorOutputWithContext(ctx context.Context) SystemSdnConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSdnConnectorOutput)
}

// SystemSdnConnectorArrayInput is an input type that accepts SystemSdnConnectorArray and SystemSdnConnectorArrayOutput values.
// You can construct a concrete instance of `SystemSdnConnectorArrayInput` via:
//
//          SystemSdnConnectorArray{ SystemSdnConnectorArgs{...} }
type SystemSdnConnectorArrayInput interface {
	pulumi.Input

	ToSystemSdnConnectorArrayOutput() SystemSdnConnectorArrayOutput
	ToSystemSdnConnectorArrayOutputWithContext(context.Context) SystemSdnConnectorArrayOutput
}

type SystemSdnConnectorArray []SystemSdnConnectorInput

func (SystemSdnConnectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSdnConnector)(nil)).Elem()
}

func (i SystemSdnConnectorArray) ToSystemSdnConnectorArrayOutput() SystemSdnConnectorArrayOutput {
	return i.ToSystemSdnConnectorArrayOutputWithContext(context.Background())
}

func (i SystemSdnConnectorArray) ToSystemSdnConnectorArrayOutputWithContext(ctx context.Context) SystemSdnConnectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSdnConnectorArrayOutput)
}

// SystemSdnConnectorMapInput is an input type that accepts SystemSdnConnectorMap and SystemSdnConnectorMapOutput values.
// You can construct a concrete instance of `SystemSdnConnectorMapInput` via:
//
//          SystemSdnConnectorMap{ "key": SystemSdnConnectorArgs{...} }
type SystemSdnConnectorMapInput interface {
	pulumi.Input

	ToSystemSdnConnectorMapOutput() SystemSdnConnectorMapOutput
	ToSystemSdnConnectorMapOutputWithContext(context.Context) SystemSdnConnectorMapOutput
}

type SystemSdnConnectorMap map[string]SystemSdnConnectorInput

func (SystemSdnConnectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSdnConnector)(nil)).Elem()
}

func (i SystemSdnConnectorMap) ToSystemSdnConnectorMapOutput() SystemSdnConnectorMapOutput {
	return i.ToSystemSdnConnectorMapOutputWithContext(context.Background())
}

func (i SystemSdnConnectorMap) ToSystemSdnConnectorMapOutputWithContext(ctx context.Context) SystemSdnConnectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSdnConnectorMapOutput)
}

type SystemSdnConnectorOutput struct{ *pulumi.OutputState }

func (SystemSdnConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSdnConnector)(nil)).Elem()
}

func (o SystemSdnConnectorOutput) ToSystemSdnConnectorOutput() SystemSdnConnectorOutput {
	return o
}

func (o SystemSdnConnectorOutput) ToSystemSdnConnectorOutputWithContext(ctx context.Context) SystemSdnConnectorOutput {
	return o
}

type SystemSdnConnectorArrayOutput struct{ *pulumi.OutputState }

func (SystemSdnConnectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSdnConnector)(nil)).Elem()
}

func (o SystemSdnConnectorArrayOutput) ToSystemSdnConnectorArrayOutput() SystemSdnConnectorArrayOutput {
	return o
}

func (o SystemSdnConnectorArrayOutput) ToSystemSdnConnectorArrayOutputWithContext(ctx context.Context) SystemSdnConnectorArrayOutput {
	return o
}

func (o SystemSdnConnectorArrayOutput) Index(i pulumi.IntInput) SystemSdnConnectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSdnConnector {
		return vs[0].([]*SystemSdnConnector)[vs[1].(int)]
	}).(SystemSdnConnectorOutput)
}

type SystemSdnConnectorMapOutput struct{ *pulumi.OutputState }

func (SystemSdnConnectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSdnConnector)(nil)).Elem()
}

func (o SystemSdnConnectorMapOutput) ToSystemSdnConnectorMapOutput() SystemSdnConnectorMapOutput {
	return o
}

func (o SystemSdnConnectorMapOutput) ToSystemSdnConnectorMapOutputWithContext(ctx context.Context) SystemSdnConnectorMapOutput {
	return o
}

func (o SystemSdnConnectorMapOutput) MapIndex(k pulumi.StringInput) SystemSdnConnectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSdnConnector {
		return vs[0].(map[string]*SystemSdnConnector)[vs[1].(string)]
	}).(SystemSdnConnectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSdnConnectorInput)(nil)).Elem(), &SystemSdnConnector{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSdnConnectorArrayInput)(nil)).Elem(), SystemSdnConnectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSdnConnectorMapInput)(nil)).Elem(), SystemSdnConnectorMap{})
	pulumi.RegisterOutputType(SystemSdnConnectorOutput{})
	pulumi.RegisterOutputType(SystemSdnConnectorArrayOutput{})
	pulumi.RegisterOutputType(SystemSdnConnectorMapOutput{})
}

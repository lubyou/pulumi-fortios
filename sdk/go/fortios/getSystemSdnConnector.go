// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system sdnconnector
func LookupSystemSdnConnector(ctx *pulumi.Context, args *LookupSystemSdnConnectorArgs, opts ...pulumi.InvokeOption) (*LookupSystemSdnConnectorResult, error) {
	var rv LookupSystemSdnConnectorResult
	err := ctx.Invoke("fortios:index/getSystemSdnConnector:GetSystemSdnConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemSdnConnector.
type LookupSystemSdnConnectorArgs struct {
	// Specify the name of the desired system sdnconnector.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemSdnConnector.
type LookupSystemSdnConnectorResult struct {
	// AWS access key ID.
	AccessKey string `pulumi:"accessKey"`
	// IBM cloud API key or service ID API key.
	ApiKey string `pulumi:"apiKey"`
	// Azure server region.
	AzureRegion string `pulumi:"azureRegion"`
	// Azure client ID (application ID).
	ClientId string `pulumi:"clientId"`
	// Azure client secret (application key).
	ClientSecret string `pulumi:"clientSecret"`
	// Compartment ID.
	CompartmentId string `pulumi:"compartmentId"`
	// Compute generation for IBM cloud infrastructure.
	ComputeGeneration int `pulumi:"computeGeneration"`
	// Domain name.
	Domain string `pulumi:"domain"`
	// Configure GCP external IP. The structure of `externalIp` block is documented below.
	ExternalIps []GetSystemSdnConnectorExternalIp `pulumi:"externalIps"`
	// GCP project name.
	GcpProject string `pulumi:"gcpProject"`
	// Group name of computers.
	GroupName string `pulumi:"groupName"`
	// Enable/disable use for FortiGate HA service.
	HaStatus string `pulumi:"haStatus"`
	// IBM cloud region name.
	IbmRegion string `pulumi:"ibmRegion"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Private key password.
	KeyPasswd string `pulumi:"keyPasswd"`
	// Azure Stack login endpoint.
	LoginEndpoint string `pulumi:"loginEndpoint"`
	// Route name.
	Name string `pulumi:"name"`
	// Configure Azure network interface. The structure of `nic` block is documented below.
	Nics []GetSystemSdnConnectorNic `pulumi:"nics"`
	// OCI certificate.
	OciCert string `pulumi:"ociCert"`
	// OCI pubkey fingerprint.
	OciFingerprint string `pulumi:"ociFingerprint"`
	// OCI server region.
	OciRegion string `pulumi:"ociRegion"`
	// OCI region type.
	OciRegionType string `pulumi:"ociRegionType"`
	// Password of the remote SDN connector as login credentials.
	Password string `pulumi:"password"`
	// Private key of GCP service account.
	PrivateKey string `pulumi:"privateKey"`
	// AWS region name.
	Region string `pulumi:"region"`
	// Resource group of Azure route table.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Azure Stack resource URL.
	ResourceUrl string `pulumi:"resourceUrl"`
	// Configure Azure route table. The structure of `routeTable` block is documented below.
	RouteTables []GetSystemSdnConnectorRouteTable `pulumi:"routeTables"`
	// Configure Azure route. The structure of `route` block is documented below.
	Routes []GetSystemSdnConnectorRoute `pulumi:"routes"`
	// AWS secret access key.
	SecretKey string `pulumi:"secretKey"`
	// Secret token of Kubernetes service account.
	SecretToken string `pulumi:"secretToken"`
	// Server address of the remote SDN connector.
	Server string `pulumi:"server"`
	// Port number of the remote SDN connector.
	ServerPort int `pulumi:"serverPort"`
	// GCP service account email.
	ServiceAccount string `pulumi:"serviceAccount"`
	// Enable/disable connection to the remote SDN connector.
	Status string `pulumi:"status"`
	// Subscription ID of Azure route table.
	SubscriptionId string `pulumi:"subscriptionId"`
	// Tenant ID (directory ID).
	TenantId string `pulumi:"tenantId"`
	// Type of SDN connector.
	Type string `pulumi:"type"`
	// Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
	UpdateInterval int `pulumi:"updateInterval"`
	// Enable/disable using IAM role from metadata to call API.
	UseMetadataIam string `pulumi:"useMetadataIam"`
	// User ID.
	UserId string `pulumi:"userId"`
	// Username of the remote SDN connector as login credentials.
	Username string `pulumi:"username"`
	// vCenter server password for NSX quarantine.
	VcenterPassword string `pulumi:"vcenterPassword"`
	// vCenter server address for NSX quarantine.
	VcenterServer string `pulumi:"vcenterServer"`
	// vCenter server username for NSX quarantine.
	VcenterUsername string  `pulumi:"vcenterUsername"`
	Vdomparam       *string `pulumi:"vdomparam"`
	// AWS VPC ID.
	VpcId string `pulumi:"vpcId"`
}
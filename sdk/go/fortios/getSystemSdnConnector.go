// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemSdnConnector(ctx *pulumi.Context, args *LookupSystemSdnConnectorArgs, opts ...pulumi.InvokeOption) (*LookupSystemSdnConnectorResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemSdnConnectorResult
	err := ctx.Invoke("fortios:index/getSystemSdnConnector:GetSystemSdnConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemSdnConnector.
type LookupSystemSdnConnectorArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemSdnConnector.
type LookupSystemSdnConnectorResult struct {
	AccessKey            string                                     `pulumi:"accessKey"`
	ApiKey               string                                     `pulumi:"apiKey"`
	AzureRegion          string                                     `pulumi:"azureRegion"`
	ClientId             string                                     `pulumi:"clientId"`
	ClientSecret         string                                     `pulumi:"clientSecret"`
	CompartmentId        string                                     `pulumi:"compartmentId"`
	ComputeGeneration    int                                        `pulumi:"computeGeneration"`
	Domain               string                                     `pulumi:"domain"`
	ExternalAccountLists []GetSystemSdnConnectorExternalAccountList `pulumi:"externalAccountLists"`
	ExternalIps          []GetSystemSdnConnectorExternalIp          `pulumi:"externalIps"`
	ForwardingRules      []GetSystemSdnConnectorForwardingRule      `pulumi:"forwardingRules"`
	GcpProject           string                                     `pulumi:"gcpProject"`
	GcpProjectLists      []GetSystemSdnConnectorGcpProjectList      `pulumi:"gcpProjectLists"`
	GroupName            string                                     `pulumi:"groupName"`
	HaStatus             string                                     `pulumi:"haStatus"`
	IbmRegion            string                                     `pulumi:"ibmRegion"`
	IbmRegionGen1        string                                     `pulumi:"ibmRegionGen1"`
	IbmRegionGen2        string                                     `pulumi:"ibmRegionGen2"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                            `pulumi:"id"`
	KeyPasswd         string                            `pulumi:"keyPasswd"`
	LoginEndpoint     string                            `pulumi:"loginEndpoint"`
	Name              string                            `pulumi:"name"`
	Nics              []GetSystemSdnConnectorNic        `pulumi:"nics"`
	OciCert           string                            `pulumi:"ociCert"`
	OciFingerprint    string                            `pulumi:"ociFingerprint"`
	OciRegion         string                            `pulumi:"ociRegion"`
	OciRegionType     string                            `pulumi:"ociRegionType"`
	Password          string                            `pulumi:"password"`
	PrivateKey        string                            `pulumi:"privateKey"`
	Region            string                            `pulumi:"region"`
	ResourceGroup     string                            `pulumi:"resourceGroup"`
	ResourceUrl       string                            `pulumi:"resourceUrl"`
	RouteTables       []GetSystemSdnConnectorRouteTable `pulumi:"routeTables"`
	Routes            []GetSystemSdnConnectorRoute      `pulumi:"routes"`
	SecretKey         string                            `pulumi:"secretKey"`
	SecretToken       string                            `pulumi:"secretToken"`
	Server            string                            `pulumi:"server"`
	ServerLists       []GetSystemSdnConnectorServerList `pulumi:"serverLists"`
	ServerPort        int                               `pulumi:"serverPort"`
	ServiceAccount    string                            `pulumi:"serviceAccount"`
	Status            string                            `pulumi:"status"`
	SubscriptionId    string                            `pulumi:"subscriptionId"`
	TenantId          string                            `pulumi:"tenantId"`
	Type              string                            `pulumi:"type"`
	UpdateInterval    int                               `pulumi:"updateInterval"`
	UseMetadataIam    string                            `pulumi:"useMetadataIam"`
	UserId            string                            `pulumi:"userId"`
	Username          string                            `pulumi:"username"`
	VcenterPassword   string                            `pulumi:"vcenterPassword"`
	VcenterServer     string                            `pulumi:"vcenterServer"`
	VcenterUsername   string                            `pulumi:"vcenterUsername"`
	Vdomparam         *string                           `pulumi:"vdomparam"`
	VerifyCertificate string                            `pulumi:"verifyCertificate"`
	VpcId             string                            `pulumi:"vpcId"`
}

func LookupSystemSdnConnectorOutput(ctx *pulumi.Context, args LookupSystemSdnConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupSystemSdnConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemSdnConnectorResult, error) {
			args := v.(LookupSystemSdnConnectorArgs)
			r, err := LookupSystemSdnConnector(ctx, &args, opts...)
			var s LookupSystemSdnConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemSdnConnectorResultOutput)
}

// A collection of arguments for invoking GetSystemSdnConnector.
type LookupSystemSdnConnectorOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemSdnConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemSdnConnectorArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemSdnConnector.
type LookupSystemSdnConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupSystemSdnConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemSdnConnectorResult)(nil)).Elem()
}

func (o LookupSystemSdnConnectorResultOutput) ToLookupSystemSdnConnectorResultOutput() LookupSystemSdnConnectorResultOutput {
	return o
}

func (o LookupSystemSdnConnectorResultOutput) ToLookupSystemSdnConnectorResultOutputWithContext(ctx context.Context) LookupSystemSdnConnectorResultOutput {
	return o
}

func (o LookupSystemSdnConnectorResultOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.AccessKey }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.ApiKey }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) AzureRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.AzureRegion }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) CompartmentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.CompartmentId }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ComputeGeneration() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) int { return v.ComputeGeneration }).(pulumi.IntOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ExternalAccountLists() GetSystemSdnConnectorExternalAccountListArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) []GetSystemSdnConnectorExternalAccountList {
		return v.ExternalAccountLists
	}).(GetSystemSdnConnectorExternalAccountListArrayOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ExternalIps() GetSystemSdnConnectorExternalIpArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) []GetSystemSdnConnectorExternalIp { return v.ExternalIps }).(GetSystemSdnConnectorExternalIpArrayOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ForwardingRules() GetSystemSdnConnectorForwardingRuleArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) []GetSystemSdnConnectorForwardingRule { return v.ForwardingRules }).(GetSystemSdnConnectorForwardingRuleArrayOutput)
}

func (o LookupSystemSdnConnectorResultOutput) GcpProject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.GcpProject }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) GcpProjectLists() GetSystemSdnConnectorGcpProjectListArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) []GetSystemSdnConnectorGcpProjectList { return v.GcpProjectLists }).(GetSystemSdnConnectorGcpProjectListArrayOutput)
}

func (o LookupSystemSdnConnectorResultOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.GroupName }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) HaStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.HaStatus }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) IbmRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.IbmRegion }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) IbmRegionGen1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.IbmRegionGen1 }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) IbmRegionGen2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.IbmRegionGen2 }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemSdnConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) KeyPasswd() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.KeyPasswd }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) LoginEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.LoginEndpoint }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Nics() GetSystemSdnConnectorNicArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) []GetSystemSdnConnectorNic { return v.Nics }).(GetSystemSdnConnectorNicArrayOutput)
}

func (o LookupSystemSdnConnectorResultOutput) OciCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.OciCert }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) OciFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.OciFingerprint }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) OciRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.OciRegion }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) OciRegionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.OciRegionType }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ResourceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.ResourceUrl }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) RouteTables() GetSystemSdnConnectorRouteTableArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) []GetSystemSdnConnectorRouteTable { return v.RouteTables }).(GetSystemSdnConnectorRouteTableArrayOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Routes() GetSystemSdnConnectorRouteArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) []GetSystemSdnConnectorRoute { return v.Routes }).(GetSystemSdnConnectorRouteArrayOutput)
}

func (o LookupSystemSdnConnectorResultOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.SecretKey }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) SecretToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.SecretToken }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.Server }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ServerLists() GetSystemSdnConnectorServerListArrayOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) []GetSystemSdnConnectorServerList { return v.ServerLists }).(GetSystemSdnConnectorServerListArrayOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) int { return v.ServerPort }).(pulumi.IntOutput)
}

func (o LookupSystemSdnConnectorResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) UpdateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) int { return v.UpdateInterval }).(pulumi.IntOutput)
}

func (o LookupSystemSdnConnectorResultOutput) UseMetadataIam() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.UseMetadataIam }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.UserId }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.Username }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) VcenterPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.VcenterPassword }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) VcenterServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.VcenterServer }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) VcenterUsername() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.VcenterUsername }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupSystemSdnConnectorResultOutput) VerifyCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.VerifyCertificate }).(pulumi.StringOutput)
}

func (o LookupSystemSdnConnectorResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemSdnConnectorResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemSdnConnectorResultOutput{})
}

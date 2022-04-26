// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system automationaction
func LookupSystemAutomationAction(ctx *pulumi.Context, args *LookupSystemAutomationActionArgs, opts ...pulumi.InvokeOption) (*LookupSystemAutomationActionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemAutomationActionResult
	err := ctx.Invoke("fortios:index/getSystemAutomationAction:GetSystemAutomationAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAutomationAction.
type LookupSystemAutomationActionArgs struct {
	// Specify the name of the desired system automationaction.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAutomationAction.
type LookupSystemAutomationActionResult struct {
	// Access profile for CLI script action to access FortiGate features.
	Accprofile string `pulumi:"accprofile"`
	// Action type.
	ActionType string `pulumi:"actionType"`
	// AliCloud AccessKey ID.
	AlicloudAccessKeyId string `pulumi:"alicloudAccessKeyId"`
	// AliCloud AccessKey secret.
	AlicloudAccessKeySecret string `pulumi:"alicloudAccessKeySecret"`
	// AliCloud account ID.
	AlicloudAccountId string `pulumi:"alicloudAccountId"`
	// AliCloud function name.
	AlicloudFunction string `pulumi:"alicloudFunction"`
	// AliCloud function authorization type.
	AlicloudFunctionAuthorization string `pulumi:"alicloudFunctionAuthorization"`
	// AliCloud function domain.
	AlicloudFunctionDomain string `pulumi:"alicloudFunctionDomain"`
	// AliCloud region.
	AlicloudRegion string `pulumi:"alicloudRegion"`
	// AliCloud service name.
	AlicloudService string `pulumi:"alicloudService"`
	// AliCloud version.
	AlicloudVersion string `pulumi:"alicloudVersion"`
	// AWS API Gateway ID.
	AwsApiId string `pulumi:"awsApiId"`
	// AWS API Gateway API key.
	AwsApiKey string `pulumi:"awsApiKey"`
	// AWS API Gateway path.
	AwsApiPath string `pulumi:"awsApiPath"`
	// AWS API Gateway deployment stage name.
	AwsApiStage string `pulumi:"awsApiStage"`
	// AWS domain.
	AwsDomain string `pulumi:"awsDomain"`
	// AWS region.
	AwsRegion string `pulumi:"awsRegion"`
	// Azure function API key.
	AzureApiKey string `pulumi:"azureApiKey"`
	// Azure function application name.
	AzureApp string `pulumi:"azureApp"`
	// Azure function domain.
	AzureDomain string `pulumi:"azureDomain"`
	// Azure function name.
	AzureFunction string `pulumi:"azureFunction"`
	// Azure function authorization level.
	AzureFunctionAuthorization string `pulumi:"azureFunctionAuthorization"`
	// Delay before execution (in seconds).
	Delay int `pulumi:"delay"`
	// Description.
	Description string `pulumi:"description"`
	// Email body.
	EmailBody string `pulumi:"emailBody"`
	// Email sender name.
	EmailFrom string `pulumi:"emailFrom"`
	// Email subject.
	EmailSubject string `pulumi:"emailSubject"`
	// Email addresses. The structure of `emailTo` block is documented below.
	EmailTos []GetSystemAutomationActionEmailTo `pulumi:"emailTos"`
	// Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric.
	ExecuteSecurityFabric string `pulumi:"executeSecurityFabric"`
	// Google Cloud function name.
	GcpFunction string `pulumi:"gcpFunction"`
	// Google Cloud function domain.
	GcpFunctionDomain string `pulumi:"gcpFunctionDomain"`
	// Google Cloud function region.
	GcpFunctionRegion string `pulumi:"gcpFunctionRegion"`
	// Google Cloud Platform project name.
	GcpProject string `pulumi:"gcpProject"`
	// Request headers. The structure of `headers` block is documented below.
	Headers []GetSystemAutomationActionHeader `pulumi:"headers"`
	// Request body (if necessary). Should be serialized json string.
	HttpBody string `pulumi:"httpBody"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Message content.
	Message string `pulumi:"message"`
	// Message type.
	MessageType string `pulumi:"messageType"`
	// Request method (POST, PUT, GET, PATCH or DELETE).
	Method string `pulumi:"method"`
	// Limit execution to no more than once in this interval (in seconds).
	MinimumInterval int `pulumi:"minimumInterval"`
	// SDN connector name.
	Name string `pulumi:"name"`
	// Protocol port.
	Port int `pulumi:"port"`
	// Request protocol.
	Protocol string `pulumi:"protocol"`
	// Enable/disable replacement message.
	ReplacementMessage string `pulumi:"replacementMessage"`
	// Replacement message group.
	ReplacemsgGroup string `pulumi:"replacemsgGroup"`
	// Required in action chain.
	Required string `pulumi:"required"`
	// CLI script.
	Script string `pulumi:"script"`
	// NSX SDN connector names. The structure of `sdnConnector` block is documented below.
	SdnConnectors []GetSystemAutomationActionSdnConnector `pulumi:"sdnConnectors"`
	// NSX security tag.
	SecurityTag string `pulumi:"securityTag"`
	// Custom TLS certificate for API request.
	TlsCertificate string `pulumi:"tlsCertificate"`
	// Request API URI.
	Uri       string  `pulumi:"uri"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable verification of the remote host certificate.
	VerifyHostCert string `pulumi:"verifyHostCert"`
}

func LookupSystemAutomationActionOutput(ctx *pulumi.Context, args LookupSystemAutomationActionOutputArgs, opts ...pulumi.InvokeOption) LookupSystemAutomationActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemAutomationActionResult, error) {
			args := v.(LookupSystemAutomationActionArgs)
			r, err := LookupSystemAutomationAction(ctx, &args, opts...)
			var s LookupSystemAutomationActionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemAutomationActionResultOutput)
}

// A collection of arguments for invoking GetSystemAutomationAction.
type LookupSystemAutomationActionOutputArgs struct {
	// Specify the name of the desired system automationaction.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemAutomationActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAutomationActionArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemAutomationAction.
type LookupSystemAutomationActionResultOutput struct{ *pulumi.OutputState }

func (LookupSystemAutomationActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAutomationActionResult)(nil)).Elem()
}

func (o LookupSystemAutomationActionResultOutput) ToLookupSystemAutomationActionResultOutput() LookupSystemAutomationActionResultOutput {
	return o
}

func (o LookupSystemAutomationActionResultOutput) ToLookupSystemAutomationActionResultOutputWithContext(ctx context.Context) LookupSystemAutomationActionResultOutput {
	return o
}

// Access profile for CLI script action to access FortiGate features.
func (o LookupSystemAutomationActionResultOutput) Accprofile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.Accprofile }).(pulumi.StringOutput)
}

// Action type.
func (o LookupSystemAutomationActionResultOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.ActionType }).(pulumi.StringOutput)
}

// AliCloud AccessKey ID.
func (o LookupSystemAutomationActionResultOutput) AlicloudAccessKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AlicloudAccessKeyId }).(pulumi.StringOutput)
}

// AliCloud AccessKey secret.
func (o LookupSystemAutomationActionResultOutput) AlicloudAccessKeySecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AlicloudAccessKeySecret }).(pulumi.StringOutput)
}

// AliCloud account ID.
func (o LookupSystemAutomationActionResultOutput) AlicloudAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AlicloudAccountId }).(pulumi.StringOutput)
}

// AliCloud function name.
func (o LookupSystemAutomationActionResultOutput) AlicloudFunction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AlicloudFunction }).(pulumi.StringOutput)
}

// AliCloud function authorization type.
func (o LookupSystemAutomationActionResultOutput) AlicloudFunctionAuthorization() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AlicloudFunctionAuthorization }).(pulumi.StringOutput)
}

// AliCloud function domain.
func (o LookupSystemAutomationActionResultOutput) AlicloudFunctionDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AlicloudFunctionDomain }).(pulumi.StringOutput)
}

// AliCloud region.
func (o LookupSystemAutomationActionResultOutput) AlicloudRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AlicloudRegion }).(pulumi.StringOutput)
}

// AliCloud service name.
func (o LookupSystemAutomationActionResultOutput) AlicloudService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AlicloudService }).(pulumi.StringOutput)
}

// AliCloud version.
func (o LookupSystemAutomationActionResultOutput) AlicloudVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AlicloudVersion }).(pulumi.StringOutput)
}

// AWS API Gateway ID.
func (o LookupSystemAutomationActionResultOutput) AwsApiId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AwsApiId }).(pulumi.StringOutput)
}

// AWS API Gateway API key.
func (o LookupSystemAutomationActionResultOutput) AwsApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AwsApiKey }).(pulumi.StringOutput)
}

// AWS API Gateway path.
func (o LookupSystemAutomationActionResultOutput) AwsApiPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AwsApiPath }).(pulumi.StringOutput)
}

// AWS API Gateway deployment stage name.
func (o LookupSystemAutomationActionResultOutput) AwsApiStage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AwsApiStage }).(pulumi.StringOutput)
}

// AWS domain.
func (o LookupSystemAutomationActionResultOutput) AwsDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AwsDomain }).(pulumi.StringOutput)
}

// AWS region.
func (o LookupSystemAutomationActionResultOutput) AwsRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AwsRegion }).(pulumi.StringOutput)
}

// Azure function API key.
func (o LookupSystemAutomationActionResultOutput) AzureApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AzureApiKey }).(pulumi.StringOutput)
}

// Azure function application name.
func (o LookupSystemAutomationActionResultOutput) AzureApp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AzureApp }).(pulumi.StringOutput)
}

// Azure function domain.
func (o LookupSystemAutomationActionResultOutput) AzureDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AzureDomain }).(pulumi.StringOutput)
}

// Azure function name.
func (o LookupSystemAutomationActionResultOutput) AzureFunction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AzureFunction }).(pulumi.StringOutput)
}

// Azure function authorization level.
func (o LookupSystemAutomationActionResultOutput) AzureFunctionAuthorization() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.AzureFunctionAuthorization }).(pulumi.StringOutput)
}

// Delay before execution (in seconds).
func (o LookupSystemAutomationActionResultOutput) Delay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) int { return v.Delay }).(pulumi.IntOutput)
}

// Description.
func (o LookupSystemAutomationActionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.Description }).(pulumi.StringOutput)
}

// Email body.
func (o LookupSystemAutomationActionResultOutput) EmailBody() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.EmailBody }).(pulumi.StringOutput)
}

// Email sender name.
func (o LookupSystemAutomationActionResultOutput) EmailFrom() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.EmailFrom }).(pulumi.StringOutput)
}

// Email subject.
func (o LookupSystemAutomationActionResultOutput) EmailSubject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.EmailSubject }).(pulumi.StringOutput)
}

// Email addresses. The structure of `emailTo` block is documented below.
func (o LookupSystemAutomationActionResultOutput) EmailTos() GetSystemAutomationActionEmailToArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) []GetSystemAutomationActionEmailTo { return v.EmailTos }).(GetSystemAutomationActionEmailToArrayOutput)
}

// Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric.
func (o LookupSystemAutomationActionResultOutput) ExecuteSecurityFabric() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.ExecuteSecurityFabric }).(pulumi.StringOutput)
}

// Google Cloud function name.
func (o LookupSystemAutomationActionResultOutput) GcpFunction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.GcpFunction }).(pulumi.StringOutput)
}

// Google Cloud function domain.
func (o LookupSystemAutomationActionResultOutput) GcpFunctionDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.GcpFunctionDomain }).(pulumi.StringOutput)
}

// Google Cloud function region.
func (o LookupSystemAutomationActionResultOutput) GcpFunctionRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.GcpFunctionRegion }).(pulumi.StringOutput)
}

// Google Cloud Platform project name.
func (o LookupSystemAutomationActionResultOutput) GcpProject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.GcpProject }).(pulumi.StringOutput)
}

// Request headers. The structure of `headers` block is documented below.
func (o LookupSystemAutomationActionResultOutput) Headers() GetSystemAutomationActionHeaderArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) []GetSystemAutomationActionHeader { return v.Headers }).(GetSystemAutomationActionHeaderArrayOutput)
}

// Request body (if necessary). Should be serialized json string.
func (o LookupSystemAutomationActionResultOutput) HttpBody() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.HttpBody }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemAutomationActionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Message content.
func (o LookupSystemAutomationActionResultOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.Message }).(pulumi.StringOutput)
}

// Message type.
func (o LookupSystemAutomationActionResultOutput) MessageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.MessageType }).(pulumi.StringOutput)
}

// Request method (POST, PUT, GET, PATCH or DELETE).
func (o LookupSystemAutomationActionResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.Method }).(pulumi.StringOutput)
}

// Limit execution to no more than once in this interval (in seconds).
func (o LookupSystemAutomationActionResultOutput) MinimumInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) int { return v.MinimumInterval }).(pulumi.IntOutput)
}

// SDN connector name.
func (o LookupSystemAutomationActionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Protocol port.
func (o LookupSystemAutomationActionResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) int { return v.Port }).(pulumi.IntOutput)
}

// Request protocol.
func (o LookupSystemAutomationActionResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// Enable/disable replacement message.
func (o LookupSystemAutomationActionResultOutput) ReplacementMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.ReplacementMessage }).(pulumi.StringOutput)
}

// Replacement message group.
func (o LookupSystemAutomationActionResultOutput) ReplacemsgGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.ReplacemsgGroup }).(pulumi.StringOutput)
}

// Required in action chain.
func (o LookupSystemAutomationActionResultOutput) Required() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.Required }).(pulumi.StringOutput)
}

// CLI script.
func (o LookupSystemAutomationActionResultOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.Script }).(pulumi.StringOutput)
}

// NSX SDN connector names. The structure of `sdnConnector` block is documented below.
func (o LookupSystemAutomationActionResultOutput) SdnConnectors() GetSystemAutomationActionSdnConnectorArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) []GetSystemAutomationActionSdnConnector {
		return v.SdnConnectors
	}).(GetSystemAutomationActionSdnConnectorArrayOutput)
}

// NSX security tag.
func (o LookupSystemAutomationActionResultOutput) SecurityTag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.SecurityTag }).(pulumi.StringOutput)
}

// Custom TLS certificate for API request.
func (o LookupSystemAutomationActionResultOutput) TlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.TlsCertificate }).(pulumi.StringOutput)
}

// Request API URI.
func (o LookupSystemAutomationActionResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.Uri }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationActionResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable verification of the remote host certificate.
func (o LookupSystemAutomationActionResultOutput) VerifyHostCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationActionResult) string { return v.VerifyHostCert }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemAutomationActionResultOutput{})
}

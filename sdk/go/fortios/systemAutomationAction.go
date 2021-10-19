// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Action for automation stitches.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewSystemAutomationAction(ctx, "trname", &fortios.SystemAutomationActionArgs{
// 			ActionType:      pulumi.String("email"),
// 			AwsDomain:       pulumi.String("amazonaws.com"),
// 			Delay:           pulumi.Int(0),
// 			EmailSubject:    pulumi.String("SUBJECT1"),
// 			Method:          pulumi.String("post"),
// 			MinimumInterval: pulumi.Int(1),
// 			Protocol:        pulumi.String("http"),
// 			Required:        pulumi.String("disable"),
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
// System AutomationAction can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemAutomationAction:SystemAutomationAction labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemAutomationAction struct {
	pulumi.CustomResourceState

	// Access profile for CLI script action to access FortiGate features.
	Accprofile pulumi.StringOutput `pulumi:"accprofile"`
	// Action type.
	ActionType pulumi.StringOutput `pulumi:"actionType"`
	// AliCloud AccessKey ID.
	AlicloudAccessKeyId pulumi.StringOutput `pulumi:"alicloudAccessKeyId"`
	// AliCloud AccessKey secret.
	AlicloudAccessKeySecret pulumi.StringPtrOutput `pulumi:"alicloudAccessKeySecret"`
	// AliCloud account ID.
	AlicloudAccountId pulumi.StringOutput `pulumi:"alicloudAccountId"`
	// AliCloud function name.
	AlicloudFunction pulumi.StringOutput `pulumi:"alicloudFunction"`
	// AliCloud function authorization type. Valid values: `anonymous`, `function`.
	AlicloudFunctionAuthorization pulumi.StringOutput `pulumi:"alicloudFunctionAuthorization"`
	// AliCloud function domain.
	AlicloudFunctionDomain pulumi.StringOutput `pulumi:"alicloudFunctionDomain"`
	// AliCloud region.
	AlicloudRegion pulumi.StringOutput `pulumi:"alicloudRegion"`
	// AliCloud service name.
	AlicloudService pulumi.StringOutput `pulumi:"alicloudService"`
	// AliCloud version.
	AlicloudVersion pulumi.StringOutput `pulumi:"alicloudVersion"`
	// AWS API Gateway ID.
	AwsApiId pulumi.StringOutput `pulumi:"awsApiId"`
	// AWS API Gateway API key.
	AwsApiKey pulumi.StringPtrOutput `pulumi:"awsApiKey"`
	// AWS API Gateway path.
	AwsApiPath pulumi.StringOutput `pulumi:"awsApiPath"`
	// AWS API Gateway deployment stage name.
	AwsApiStage pulumi.StringOutput `pulumi:"awsApiStage"`
	// AWS domain.
	AwsDomain pulumi.StringOutput `pulumi:"awsDomain"`
	// AWS region.
	AwsRegion pulumi.StringOutput `pulumi:"awsRegion"`
	// Azure function API key.
	AzureApiKey pulumi.StringPtrOutput `pulumi:"azureApiKey"`
	// Azure function application name.
	AzureApp pulumi.StringOutput `pulumi:"azureApp"`
	// Azure function domain.
	AzureDomain pulumi.StringOutput `pulumi:"azureDomain"`
	// Azure function name.
	AzureFunction pulumi.StringOutput `pulumi:"azureFunction"`
	// Azure function authorization level. Valid values: `anonymous`, `function`, `admin`.
	AzureFunctionAuthorization pulumi.StringOutput `pulumi:"azureFunctionAuthorization"`
	// Delay before execution (in seconds).
	Delay pulumi.IntOutput `pulumi:"delay"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Email body.
	EmailBody pulumi.StringOutput `pulumi:"emailBody"`
	// Email sender name.
	EmailFrom pulumi.StringPtrOutput `pulumi:"emailFrom"`
	// Email subject.
	EmailSubject pulumi.StringPtrOutput `pulumi:"emailSubject"`
	// Email addresses. The structure of `emailTo` block is documented below.
	EmailTos SystemAutomationActionEmailToArrayOutput `pulumi:"emailTos"`
	// Google Cloud function name.
	GcpFunction pulumi.StringOutput `pulumi:"gcpFunction"`
	// Google Cloud function domain.
	GcpFunctionDomain pulumi.StringOutput `pulumi:"gcpFunctionDomain"`
	// Google Cloud function region.
	GcpFunctionRegion pulumi.StringOutput `pulumi:"gcpFunctionRegion"`
	// Google Cloud Platform project name.
	GcpProject pulumi.StringOutput `pulumi:"gcpProject"`
	// Request headers. The structure of `headers` block is documented below.
	Headers SystemAutomationActionHeaderArrayOutput `pulumi:"headers"`
	// Request body (if necessary). Should be serialized json string.
	HttpBody pulumi.StringPtrOutput `pulumi:"httpBody"`
	// Message content.
	Message pulumi.StringOutput `pulumi:"message"`
	// Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `post`, `put`, `get`, `patch`, `delete`.
	Method pulumi.StringOutput `pulumi:"method"`
	// Limit execution to no more than once in this interval (in seconds).
	MinimumInterval pulumi.IntOutput `pulumi:"minimumInterval"`
	// SDN connector name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Protocol port.
	Port pulumi.IntOutput `pulumi:"port"`
	// Request protocol. Valid values: `http`, `https`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Enable/disable replacement message. Valid values: `enable`, `disable`.
	ReplacementMessage pulumi.StringOutput `pulumi:"replacementMessage"`
	// Required in action chain. Valid values: `enable`, `disable`.
	Required pulumi.StringOutput `pulumi:"required"`
	// CLI script.
	Script pulumi.StringPtrOutput `pulumi:"script"`
	// NSX SDN connector names. The structure of `sdnConnector` block is documented below.
	SdnConnectors SystemAutomationActionSdnConnectorArrayOutput `pulumi:"sdnConnectors"`
	// NSX security tag.
	SecurityTag pulumi.StringOutput `pulumi:"securityTag"`
	// Custom TLS certificate for API request.
	TlsCertificate pulumi.StringOutput `pulumi:"tlsCertificate"`
	// Request API URI.
	Uri pulumi.StringPtrOutput `pulumi:"uri"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemAutomationAction registers a new resource with the given unique name, arguments, and options.
func NewSystemAutomationAction(ctx *pulumi.Context,
	name string, args *SystemAutomationActionArgs, opts ...pulumi.ResourceOption) (*SystemAutomationAction, error) {
	if args == nil {
		args = &SystemAutomationActionArgs{}
	}

	var resource SystemAutomationAction
	err := ctx.RegisterResource("fortios:index/systemAutomationAction:SystemAutomationAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAutomationAction gets an existing SystemAutomationAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAutomationAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAutomationActionState, opts ...pulumi.ResourceOption) (*SystemAutomationAction, error) {
	var resource SystemAutomationAction
	err := ctx.ReadResource("fortios:index/systemAutomationAction:SystemAutomationAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAutomationAction resources.
type systemAutomationActionState struct {
	// Access profile for CLI script action to access FortiGate features.
	Accprofile *string `pulumi:"accprofile"`
	// Action type.
	ActionType *string `pulumi:"actionType"`
	// AliCloud AccessKey ID.
	AlicloudAccessKeyId *string `pulumi:"alicloudAccessKeyId"`
	// AliCloud AccessKey secret.
	AlicloudAccessKeySecret *string `pulumi:"alicloudAccessKeySecret"`
	// AliCloud account ID.
	AlicloudAccountId *string `pulumi:"alicloudAccountId"`
	// AliCloud function name.
	AlicloudFunction *string `pulumi:"alicloudFunction"`
	// AliCloud function authorization type. Valid values: `anonymous`, `function`.
	AlicloudFunctionAuthorization *string `pulumi:"alicloudFunctionAuthorization"`
	// AliCloud function domain.
	AlicloudFunctionDomain *string `pulumi:"alicloudFunctionDomain"`
	// AliCloud region.
	AlicloudRegion *string `pulumi:"alicloudRegion"`
	// AliCloud service name.
	AlicloudService *string `pulumi:"alicloudService"`
	// AliCloud version.
	AlicloudVersion *string `pulumi:"alicloudVersion"`
	// AWS API Gateway ID.
	AwsApiId *string `pulumi:"awsApiId"`
	// AWS API Gateway API key.
	AwsApiKey *string `pulumi:"awsApiKey"`
	// AWS API Gateway path.
	AwsApiPath *string `pulumi:"awsApiPath"`
	// AWS API Gateway deployment stage name.
	AwsApiStage *string `pulumi:"awsApiStage"`
	// AWS domain.
	AwsDomain *string `pulumi:"awsDomain"`
	// AWS region.
	AwsRegion *string `pulumi:"awsRegion"`
	// Azure function API key.
	AzureApiKey *string `pulumi:"azureApiKey"`
	// Azure function application name.
	AzureApp *string `pulumi:"azureApp"`
	// Azure function domain.
	AzureDomain *string `pulumi:"azureDomain"`
	// Azure function name.
	AzureFunction *string `pulumi:"azureFunction"`
	// Azure function authorization level. Valid values: `anonymous`, `function`, `admin`.
	AzureFunctionAuthorization *string `pulumi:"azureFunctionAuthorization"`
	// Delay before execution (in seconds).
	Delay *int `pulumi:"delay"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Email body.
	EmailBody *string `pulumi:"emailBody"`
	// Email sender name.
	EmailFrom *string `pulumi:"emailFrom"`
	// Email subject.
	EmailSubject *string `pulumi:"emailSubject"`
	// Email addresses. The structure of `emailTo` block is documented below.
	EmailTos []SystemAutomationActionEmailTo `pulumi:"emailTos"`
	// Google Cloud function name.
	GcpFunction *string `pulumi:"gcpFunction"`
	// Google Cloud function domain.
	GcpFunctionDomain *string `pulumi:"gcpFunctionDomain"`
	// Google Cloud function region.
	GcpFunctionRegion *string `pulumi:"gcpFunctionRegion"`
	// Google Cloud Platform project name.
	GcpProject *string `pulumi:"gcpProject"`
	// Request headers. The structure of `headers` block is documented below.
	Headers []SystemAutomationActionHeader `pulumi:"headers"`
	// Request body (if necessary). Should be serialized json string.
	HttpBody *string `pulumi:"httpBody"`
	// Message content.
	Message *string `pulumi:"message"`
	// Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `post`, `put`, `get`, `patch`, `delete`.
	Method *string `pulumi:"method"`
	// Limit execution to no more than once in this interval (in seconds).
	MinimumInterval *int `pulumi:"minimumInterval"`
	// SDN connector name.
	Name *string `pulumi:"name"`
	// Protocol port.
	Port *int `pulumi:"port"`
	// Request protocol. Valid values: `http`, `https`.
	Protocol *string `pulumi:"protocol"`
	// Enable/disable replacement message. Valid values: `enable`, `disable`.
	ReplacementMessage *string `pulumi:"replacementMessage"`
	// Required in action chain. Valid values: `enable`, `disable`.
	Required *string `pulumi:"required"`
	// CLI script.
	Script *string `pulumi:"script"`
	// NSX SDN connector names. The structure of `sdnConnector` block is documented below.
	SdnConnectors []SystemAutomationActionSdnConnector `pulumi:"sdnConnectors"`
	// NSX security tag.
	SecurityTag *string `pulumi:"securityTag"`
	// Custom TLS certificate for API request.
	TlsCertificate *string `pulumi:"tlsCertificate"`
	// Request API URI.
	Uri *string `pulumi:"uri"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemAutomationActionState struct {
	// Access profile for CLI script action to access FortiGate features.
	Accprofile pulumi.StringPtrInput
	// Action type.
	ActionType pulumi.StringPtrInput
	// AliCloud AccessKey ID.
	AlicloudAccessKeyId pulumi.StringPtrInput
	// AliCloud AccessKey secret.
	AlicloudAccessKeySecret pulumi.StringPtrInput
	// AliCloud account ID.
	AlicloudAccountId pulumi.StringPtrInput
	// AliCloud function name.
	AlicloudFunction pulumi.StringPtrInput
	// AliCloud function authorization type. Valid values: `anonymous`, `function`.
	AlicloudFunctionAuthorization pulumi.StringPtrInput
	// AliCloud function domain.
	AlicloudFunctionDomain pulumi.StringPtrInput
	// AliCloud region.
	AlicloudRegion pulumi.StringPtrInput
	// AliCloud service name.
	AlicloudService pulumi.StringPtrInput
	// AliCloud version.
	AlicloudVersion pulumi.StringPtrInput
	// AWS API Gateway ID.
	AwsApiId pulumi.StringPtrInput
	// AWS API Gateway API key.
	AwsApiKey pulumi.StringPtrInput
	// AWS API Gateway path.
	AwsApiPath pulumi.StringPtrInput
	// AWS API Gateway deployment stage name.
	AwsApiStage pulumi.StringPtrInput
	// AWS domain.
	AwsDomain pulumi.StringPtrInput
	// AWS region.
	AwsRegion pulumi.StringPtrInput
	// Azure function API key.
	AzureApiKey pulumi.StringPtrInput
	// Azure function application name.
	AzureApp pulumi.StringPtrInput
	// Azure function domain.
	AzureDomain pulumi.StringPtrInput
	// Azure function name.
	AzureFunction pulumi.StringPtrInput
	// Azure function authorization level. Valid values: `anonymous`, `function`, `admin`.
	AzureFunctionAuthorization pulumi.StringPtrInput
	// Delay before execution (in seconds).
	Delay pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Email body.
	EmailBody pulumi.StringPtrInput
	// Email sender name.
	EmailFrom pulumi.StringPtrInput
	// Email subject.
	EmailSubject pulumi.StringPtrInput
	// Email addresses. The structure of `emailTo` block is documented below.
	EmailTos SystemAutomationActionEmailToArrayInput
	// Google Cloud function name.
	GcpFunction pulumi.StringPtrInput
	// Google Cloud function domain.
	GcpFunctionDomain pulumi.StringPtrInput
	// Google Cloud function region.
	GcpFunctionRegion pulumi.StringPtrInput
	// Google Cloud Platform project name.
	GcpProject pulumi.StringPtrInput
	// Request headers. The structure of `headers` block is documented below.
	Headers SystemAutomationActionHeaderArrayInput
	// Request body (if necessary). Should be serialized json string.
	HttpBody pulumi.StringPtrInput
	// Message content.
	Message pulumi.StringPtrInput
	// Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `post`, `put`, `get`, `patch`, `delete`.
	Method pulumi.StringPtrInput
	// Limit execution to no more than once in this interval (in seconds).
	MinimumInterval pulumi.IntPtrInput
	// SDN connector name.
	Name pulumi.StringPtrInput
	// Protocol port.
	Port pulumi.IntPtrInput
	// Request protocol. Valid values: `http`, `https`.
	Protocol pulumi.StringPtrInput
	// Enable/disable replacement message. Valid values: `enable`, `disable`.
	ReplacementMessage pulumi.StringPtrInput
	// Required in action chain. Valid values: `enable`, `disable`.
	Required pulumi.StringPtrInput
	// CLI script.
	Script pulumi.StringPtrInput
	// NSX SDN connector names. The structure of `sdnConnector` block is documented below.
	SdnConnectors SystemAutomationActionSdnConnectorArrayInput
	// NSX security tag.
	SecurityTag pulumi.StringPtrInput
	// Custom TLS certificate for API request.
	TlsCertificate pulumi.StringPtrInput
	// Request API URI.
	Uri pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAutomationActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutomationActionState)(nil)).Elem()
}

type systemAutomationActionArgs struct {
	// Access profile for CLI script action to access FortiGate features.
	Accprofile *string `pulumi:"accprofile"`
	// Action type.
	ActionType *string `pulumi:"actionType"`
	// AliCloud AccessKey ID.
	AlicloudAccessKeyId *string `pulumi:"alicloudAccessKeyId"`
	// AliCloud AccessKey secret.
	AlicloudAccessKeySecret *string `pulumi:"alicloudAccessKeySecret"`
	// AliCloud account ID.
	AlicloudAccountId *string `pulumi:"alicloudAccountId"`
	// AliCloud function name.
	AlicloudFunction *string `pulumi:"alicloudFunction"`
	// AliCloud function authorization type. Valid values: `anonymous`, `function`.
	AlicloudFunctionAuthorization *string `pulumi:"alicloudFunctionAuthorization"`
	// AliCloud function domain.
	AlicloudFunctionDomain *string `pulumi:"alicloudFunctionDomain"`
	// AliCloud region.
	AlicloudRegion *string `pulumi:"alicloudRegion"`
	// AliCloud service name.
	AlicloudService *string `pulumi:"alicloudService"`
	// AliCloud version.
	AlicloudVersion *string `pulumi:"alicloudVersion"`
	// AWS API Gateway ID.
	AwsApiId *string `pulumi:"awsApiId"`
	// AWS API Gateway API key.
	AwsApiKey *string `pulumi:"awsApiKey"`
	// AWS API Gateway path.
	AwsApiPath *string `pulumi:"awsApiPath"`
	// AWS API Gateway deployment stage name.
	AwsApiStage *string `pulumi:"awsApiStage"`
	// AWS domain.
	AwsDomain *string `pulumi:"awsDomain"`
	// AWS region.
	AwsRegion *string `pulumi:"awsRegion"`
	// Azure function API key.
	AzureApiKey *string `pulumi:"azureApiKey"`
	// Azure function application name.
	AzureApp *string `pulumi:"azureApp"`
	// Azure function domain.
	AzureDomain *string `pulumi:"azureDomain"`
	// Azure function name.
	AzureFunction *string `pulumi:"azureFunction"`
	// Azure function authorization level. Valid values: `anonymous`, `function`, `admin`.
	AzureFunctionAuthorization *string `pulumi:"azureFunctionAuthorization"`
	// Delay before execution (in seconds).
	Delay *int `pulumi:"delay"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Email body.
	EmailBody *string `pulumi:"emailBody"`
	// Email sender name.
	EmailFrom *string `pulumi:"emailFrom"`
	// Email subject.
	EmailSubject *string `pulumi:"emailSubject"`
	// Email addresses. The structure of `emailTo` block is documented below.
	EmailTos []SystemAutomationActionEmailTo `pulumi:"emailTos"`
	// Google Cloud function name.
	GcpFunction *string `pulumi:"gcpFunction"`
	// Google Cloud function domain.
	GcpFunctionDomain *string `pulumi:"gcpFunctionDomain"`
	// Google Cloud function region.
	GcpFunctionRegion *string `pulumi:"gcpFunctionRegion"`
	// Google Cloud Platform project name.
	GcpProject *string `pulumi:"gcpProject"`
	// Request headers. The structure of `headers` block is documented below.
	Headers []SystemAutomationActionHeader `pulumi:"headers"`
	// Request body (if necessary). Should be serialized json string.
	HttpBody *string `pulumi:"httpBody"`
	// Message content.
	Message *string `pulumi:"message"`
	// Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `post`, `put`, `get`, `patch`, `delete`.
	Method *string `pulumi:"method"`
	// Limit execution to no more than once in this interval (in seconds).
	MinimumInterval *int `pulumi:"minimumInterval"`
	// SDN connector name.
	Name *string `pulumi:"name"`
	// Protocol port.
	Port *int `pulumi:"port"`
	// Request protocol. Valid values: `http`, `https`.
	Protocol *string `pulumi:"protocol"`
	// Enable/disable replacement message. Valid values: `enable`, `disable`.
	ReplacementMessage *string `pulumi:"replacementMessage"`
	// Required in action chain. Valid values: `enable`, `disable`.
	Required *string `pulumi:"required"`
	// CLI script.
	Script *string `pulumi:"script"`
	// NSX SDN connector names. The structure of `sdnConnector` block is documented below.
	SdnConnectors []SystemAutomationActionSdnConnector `pulumi:"sdnConnectors"`
	// NSX security tag.
	SecurityTag *string `pulumi:"securityTag"`
	// Custom TLS certificate for API request.
	TlsCertificate *string `pulumi:"tlsCertificate"`
	// Request API URI.
	Uri *string `pulumi:"uri"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAutomationAction resource.
type SystemAutomationActionArgs struct {
	// Access profile for CLI script action to access FortiGate features.
	Accprofile pulumi.StringPtrInput
	// Action type.
	ActionType pulumi.StringPtrInput
	// AliCloud AccessKey ID.
	AlicloudAccessKeyId pulumi.StringPtrInput
	// AliCloud AccessKey secret.
	AlicloudAccessKeySecret pulumi.StringPtrInput
	// AliCloud account ID.
	AlicloudAccountId pulumi.StringPtrInput
	// AliCloud function name.
	AlicloudFunction pulumi.StringPtrInput
	// AliCloud function authorization type. Valid values: `anonymous`, `function`.
	AlicloudFunctionAuthorization pulumi.StringPtrInput
	// AliCloud function domain.
	AlicloudFunctionDomain pulumi.StringPtrInput
	// AliCloud region.
	AlicloudRegion pulumi.StringPtrInput
	// AliCloud service name.
	AlicloudService pulumi.StringPtrInput
	// AliCloud version.
	AlicloudVersion pulumi.StringPtrInput
	// AWS API Gateway ID.
	AwsApiId pulumi.StringPtrInput
	// AWS API Gateway API key.
	AwsApiKey pulumi.StringPtrInput
	// AWS API Gateway path.
	AwsApiPath pulumi.StringPtrInput
	// AWS API Gateway deployment stage name.
	AwsApiStage pulumi.StringPtrInput
	// AWS domain.
	AwsDomain pulumi.StringPtrInput
	// AWS region.
	AwsRegion pulumi.StringPtrInput
	// Azure function API key.
	AzureApiKey pulumi.StringPtrInput
	// Azure function application name.
	AzureApp pulumi.StringPtrInput
	// Azure function domain.
	AzureDomain pulumi.StringPtrInput
	// Azure function name.
	AzureFunction pulumi.StringPtrInput
	// Azure function authorization level. Valid values: `anonymous`, `function`, `admin`.
	AzureFunctionAuthorization pulumi.StringPtrInput
	// Delay before execution (in seconds).
	Delay pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Email body.
	EmailBody pulumi.StringPtrInput
	// Email sender name.
	EmailFrom pulumi.StringPtrInput
	// Email subject.
	EmailSubject pulumi.StringPtrInput
	// Email addresses. The structure of `emailTo` block is documented below.
	EmailTos SystemAutomationActionEmailToArrayInput
	// Google Cloud function name.
	GcpFunction pulumi.StringPtrInput
	// Google Cloud function domain.
	GcpFunctionDomain pulumi.StringPtrInput
	// Google Cloud function region.
	GcpFunctionRegion pulumi.StringPtrInput
	// Google Cloud Platform project name.
	GcpProject pulumi.StringPtrInput
	// Request headers. The structure of `headers` block is documented below.
	Headers SystemAutomationActionHeaderArrayInput
	// Request body (if necessary). Should be serialized json string.
	HttpBody pulumi.StringPtrInput
	// Message content.
	Message pulumi.StringPtrInput
	// Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `post`, `put`, `get`, `patch`, `delete`.
	Method pulumi.StringPtrInput
	// Limit execution to no more than once in this interval (in seconds).
	MinimumInterval pulumi.IntPtrInput
	// SDN connector name.
	Name pulumi.StringPtrInput
	// Protocol port.
	Port pulumi.IntPtrInput
	// Request protocol. Valid values: `http`, `https`.
	Protocol pulumi.StringPtrInput
	// Enable/disable replacement message. Valid values: `enable`, `disable`.
	ReplacementMessage pulumi.StringPtrInput
	// Required in action chain. Valid values: `enable`, `disable`.
	Required pulumi.StringPtrInput
	// CLI script.
	Script pulumi.StringPtrInput
	// NSX SDN connector names. The structure of `sdnConnector` block is documented below.
	SdnConnectors SystemAutomationActionSdnConnectorArrayInput
	// NSX security tag.
	SecurityTag pulumi.StringPtrInput
	// Custom TLS certificate for API request.
	TlsCertificate pulumi.StringPtrInput
	// Request API URI.
	Uri pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAutomationActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutomationActionArgs)(nil)).Elem()
}

type SystemAutomationActionInput interface {
	pulumi.Input

	ToSystemAutomationActionOutput() SystemAutomationActionOutput
	ToSystemAutomationActionOutputWithContext(ctx context.Context) SystemAutomationActionOutput
}

func (*SystemAutomationAction) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAutomationAction)(nil))
}

func (i *SystemAutomationAction) ToSystemAutomationActionOutput() SystemAutomationActionOutput {
	return i.ToSystemAutomationActionOutputWithContext(context.Background())
}

func (i *SystemAutomationAction) ToSystemAutomationActionOutputWithContext(ctx context.Context) SystemAutomationActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationActionOutput)
}

func (i *SystemAutomationAction) ToSystemAutomationActionPtrOutput() SystemAutomationActionPtrOutput {
	return i.ToSystemAutomationActionPtrOutputWithContext(context.Background())
}

func (i *SystemAutomationAction) ToSystemAutomationActionPtrOutputWithContext(ctx context.Context) SystemAutomationActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationActionPtrOutput)
}

type SystemAutomationActionPtrInput interface {
	pulumi.Input

	ToSystemAutomationActionPtrOutput() SystemAutomationActionPtrOutput
	ToSystemAutomationActionPtrOutputWithContext(ctx context.Context) SystemAutomationActionPtrOutput
}

type systemAutomationActionPtrType SystemAutomationActionArgs

func (*systemAutomationActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutomationAction)(nil))
}

func (i *systemAutomationActionPtrType) ToSystemAutomationActionPtrOutput() SystemAutomationActionPtrOutput {
	return i.ToSystemAutomationActionPtrOutputWithContext(context.Background())
}

func (i *systemAutomationActionPtrType) ToSystemAutomationActionPtrOutputWithContext(ctx context.Context) SystemAutomationActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationActionPtrOutput)
}

// SystemAutomationActionArrayInput is an input type that accepts SystemAutomationActionArray and SystemAutomationActionArrayOutput values.
// You can construct a concrete instance of `SystemAutomationActionArrayInput` via:
//
//          SystemAutomationActionArray{ SystemAutomationActionArgs{...} }
type SystemAutomationActionArrayInput interface {
	pulumi.Input

	ToSystemAutomationActionArrayOutput() SystemAutomationActionArrayOutput
	ToSystemAutomationActionArrayOutputWithContext(context.Context) SystemAutomationActionArrayOutput
}

type SystemAutomationActionArray []SystemAutomationActionInput

func (SystemAutomationActionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemAutomationAction)(nil))
}

func (i SystemAutomationActionArray) ToSystemAutomationActionArrayOutput() SystemAutomationActionArrayOutput {
	return i.ToSystemAutomationActionArrayOutputWithContext(context.Background())
}

func (i SystemAutomationActionArray) ToSystemAutomationActionArrayOutputWithContext(ctx context.Context) SystemAutomationActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationActionArrayOutput)
}

// SystemAutomationActionMapInput is an input type that accepts SystemAutomationActionMap and SystemAutomationActionMapOutput values.
// You can construct a concrete instance of `SystemAutomationActionMapInput` via:
//
//          SystemAutomationActionMap{ "key": SystemAutomationActionArgs{...} }
type SystemAutomationActionMapInput interface {
	pulumi.Input

	ToSystemAutomationActionMapOutput() SystemAutomationActionMapOutput
	ToSystemAutomationActionMapOutputWithContext(context.Context) SystemAutomationActionMapOutput
}

type SystemAutomationActionMap map[string]SystemAutomationActionInput

func (SystemAutomationActionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemAutomationAction)(nil))
}

func (i SystemAutomationActionMap) ToSystemAutomationActionMapOutput() SystemAutomationActionMapOutput {
	return i.ToSystemAutomationActionMapOutputWithContext(context.Background())
}

func (i SystemAutomationActionMap) ToSystemAutomationActionMapOutputWithContext(ctx context.Context) SystemAutomationActionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationActionMapOutput)
}

type SystemAutomationActionOutput struct {
	*pulumi.OutputState
}

func (SystemAutomationActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAutomationAction)(nil))
}

func (o SystemAutomationActionOutput) ToSystemAutomationActionOutput() SystemAutomationActionOutput {
	return o
}

func (o SystemAutomationActionOutput) ToSystemAutomationActionOutputWithContext(ctx context.Context) SystemAutomationActionOutput {
	return o
}

func (o SystemAutomationActionOutput) ToSystemAutomationActionPtrOutput() SystemAutomationActionPtrOutput {
	return o.ToSystemAutomationActionPtrOutputWithContext(context.Background())
}

func (o SystemAutomationActionOutput) ToSystemAutomationActionPtrOutputWithContext(ctx context.Context) SystemAutomationActionPtrOutput {
	return o.ApplyT(func(v SystemAutomationAction) *SystemAutomationAction {
		return &v
	}).(SystemAutomationActionPtrOutput)
}

type SystemAutomationActionPtrOutput struct {
	*pulumi.OutputState
}

func (SystemAutomationActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutomationAction)(nil))
}

func (o SystemAutomationActionPtrOutput) ToSystemAutomationActionPtrOutput() SystemAutomationActionPtrOutput {
	return o
}

func (o SystemAutomationActionPtrOutput) ToSystemAutomationActionPtrOutputWithContext(ctx context.Context) SystemAutomationActionPtrOutput {
	return o
}

type SystemAutomationActionArrayOutput struct{ *pulumi.OutputState }

func (SystemAutomationActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemAutomationAction)(nil))
}

func (o SystemAutomationActionArrayOutput) ToSystemAutomationActionArrayOutput() SystemAutomationActionArrayOutput {
	return o
}

func (o SystemAutomationActionArrayOutput) ToSystemAutomationActionArrayOutputWithContext(ctx context.Context) SystemAutomationActionArrayOutput {
	return o
}

func (o SystemAutomationActionArrayOutput) Index(i pulumi.IntInput) SystemAutomationActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemAutomationAction {
		return vs[0].([]SystemAutomationAction)[vs[1].(int)]
	}).(SystemAutomationActionOutput)
}

type SystemAutomationActionMapOutput struct{ *pulumi.OutputState }

func (SystemAutomationActionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemAutomationAction)(nil))
}

func (o SystemAutomationActionMapOutput) ToSystemAutomationActionMapOutput() SystemAutomationActionMapOutput {
	return o
}

func (o SystemAutomationActionMapOutput) ToSystemAutomationActionMapOutputWithContext(ctx context.Context) SystemAutomationActionMapOutput {
	return o
}

func (o SystemAutomationActionMapOutput) MapIndex(k pulumi.StringInput) SystemAutomationActionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemAutomationAction {
		return vs[0].(map[string]SystemAutomationAction)[vs[1].(string)]
	}).(SystemAutomationActionOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemAutomationActionOutput{})
	pulumi.RegisterOutputType(SystemAutomationActionPtrOutput{})
	pulumi.RegisterOutputType(SystemAutomationActionArrayOutput{})
	pulumi.RegisterOutputType(SystemAutomationActionMapOutput{})
}

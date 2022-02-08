// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure general log settings.
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
// 		_, err := fortios.NewLogSetting(ctx, "trname", &fortios.LogSettingArgs{
// 			BriefTrafficFormat:   pulumi.String("disable"),
// 			DaemonLog:            pulumi.String("disable"),
// 			ExpolicyImplicitLog:  pulumi.String("disable"),
// 			FazOverride:          pulumi.String("disable"),
// 			Fwpolicy6ImplicitLog: pulumi.String("disable"),
// 			FwpolicyImplicitLog:  pulumi.String("disable"),
// 			LocalInAllow:         pulumi.String("disable"),
// 			LocalInDenyBroadcast: pulumi.String("disable"),
// 			LocalInDenyUnicast:   pulumi.String("disable"),
// 			LocalOut:             pulumi.String("disable"),
// 			LogInvalidPacket:     pulumi.String("disable"),
// 			LogPolicyComment:     pulumi.String("disable"),
// 			LogPolicyName:        pulumi.String("disable"),
// 			LogUserInUpper:       pulumi.String("disable"),
// 			NeighborEvent:        pulumi.String("disable"),
// 			ResolveIp:            pulumi.String("disable"),
// 			ResolvePort:          pulumi.String("enable"),
// 			SyslogOverride:       pulumi.String("disable"),
// 			UserAnonymize:        pulumi.String("disable"),
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
// Log Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logSetting:LogSetting labelname LogSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSetting struct {
	pulumi.CustomResourceState

	// User name anonymization hash salt.
	AnonymizationHash pulumi.StringOutput `pulumi:"anonymizationHash"`
	// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
	BriefTrafficFormat pulumi.StringOutput `pulumi:"briefTrafficFormat"`
	// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
	CustomLogFields LogSettingCustomLogFieldArrayOutput `pulumi:"customLogFields"`
	// Enable/disable daemon logging. Valid values: `enable`, `disable`.
	DaemonLog pulumi.StringOutput `pulumi:"daemonLog"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
	ExpolicyImplicitLog pulumi.StringOutput `pulumi:"expolicyImplicitLog"`
	// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
	FazOverride pulumi.StringOutput `pulumi:"fazOverride"`
	// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
	Fwpolicy6ImplicitLog pulumi.StringOutput `pulumi:"fwpolicy6ImplicitLog"`
	// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
	FwpolicyImplicitLog pulumi.StringOutput `pulumi:"fwpolicyImplicitLog"`
	// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
	LocalInAllow pulumi.StringOutput `pulumi:"localInAllow"`
	// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
	LocalInDenyBroadcast pulumi.StringOutput `pulumi:"localInDenyBroadcast"`
	// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
	LocalInDenyUnicast pulumi.StringOutput `pulumi:"localInDenyUnicast"`
	// Enable/disable local-out logging. Valid values: `enable`, `disable`.
	LocalOut pulumi.StringOutput `pulumi:"localOut"`
	// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
	LogInvalidPacket pulumi.StringOutput `pulumi:"logInvalidPacket"`
	// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyComment pulumi.StringOutput `pulumi:"logPolicyComment"`
	// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyName pulumi.StringOutput `pulumi:"logPolicyName"`
	// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
	LogUserInUpper pulumi.StringOutput `pulumi:"logUserInUpper"`
	// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
	NeighborEvent pulumi.StringOutput `pulumi:"neighborEvent"`
	// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
	ResolveIp pulumi.StringOutput `pulumi:"resolveIp"`
	// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
	ResolvePort pulumi.StringOutput `pulumi:"resolvePort"`
	// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
	RestApiGet pulumi.StringOutput `pulumi:"restApiGet"`
	// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
	RestApiSet pulumi.StringOutput `pulumi:"restApiSet"`
	// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
	SyslogOverride pulumi.StringOutput `pulumi:"syslogOverride"`
	// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
	UserAnonymize pulumi.StringOutput `pulumi:"userAnonymize"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogSetting registers a new resource with the given unique name, arguments, and options.
func NewLogSetting(ctx *pulumi.Context,
	name string, args *LogSettingArgs, opts ...pulumi.ResourceOption) (*LogSetting, error) {
	if args == nil {
		args = &LogSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSetting
	err := ctx.RegisterResource("fortios:index/logSetting:LogSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSetting gets an existing LogSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSettingState, opts ...pulumi.ResourceOption) (*LogSetting, error) {
	var resource LogSetting
	err := ctx.ReadResource("fortios:index/logSetting:LogSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSetting resources.
type logSettingState struct {
	// User name anonymization hash salt.
	AnonymizationHash *string `pulumi:"anonymizationHash"`
	// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
	BriefTrafficFormat *string `pulumi:"briefTrafficFormat"`
	// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
	CustomLogFields []LogSettingCustomLogField `pulumi:"customLogFields"`
	// Enable/disable daemon logging. Valid values: `enable`, `disable`.
	DaemonLog *string `pulumi:"daemonLog"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
	ExpolicyImplicitLog *string `pulumi:"expolicyImplicitLog"`
	// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
	FazOverride *string `pulumi:"fazOverride"`
	// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
	Fwpolicy6ImplicitLog *string `pulumi:"fwpolicy6ImplicitLog"`
	// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
	FwpolicyImplicitLog *string `pulumi:"fwpolicyImplicitLog"`
	// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
	LocalInAllow *string `pulumi:"localInAllow"`
	// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
	LocalInDenyBroadcast *string `pulumi:"localInDenyBroadcast"`
	// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
	LocalInDenyUnicast *string `pulumi:"localInDenyUnicast"`
	// Enable/disable local-out logging. Valid values: `enable`, `disable`.
	LocalOut *string `pulumi:"localOut"`
	// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
	LogInvalidPacket *string `pulumi:"logInvalidPacket"`
	// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyComment *string `pulumi:"logPolicyComment"`
	// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyName *string `pulumi:"logPolicyName"`
	// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
	LogUserInUpper *string `pulumi:"logUserInUpper"`
	// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
	NeighborEvent *string `pulumi:"neighborEvent"`
	// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
	ResolveIp *string `pulumi:"resolveIp"`
	// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
	ResolvePort *string `pulumi:"resolvePort"`
	// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
	RestApiGet *string `pulumi:"restApiGet"`
	// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
	RestApiSet *string `pulumi:"restApiSet"`
	// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
	SyslogOverride *string `pulumi:"syslogOverride"`
	// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
	UserAnonymize *string `pulumi:"userAnonymize"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogSettingState struct {
	// User name anonymization hash salt.
	AnonymizationHash pulumi.StringPtrInput
	// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
	BriefTrafficFormat pulumi.StringPtrInput
	// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
	CustomLogFields LogSettingCustomLogFieldArrayInput
	// Enable/disable daemon logging. Valid values: `enable`, `disable`.
	DaemonLog pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
	ExpolicyImplicitLog pulumi.StringPtrInput
	// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
	FazOverride pulumi.StringPtrInput
	// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
	Fwpolicy6ImplicitLog pulumi.StringPtrInput
	// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
	FwpolicyImplicitLog pulumi.StringPtrInput
	// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
	LocalInAllow pulumi.StringPtrInput
	// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
	LocalInDenyBroadcast pulumi.StringPtrInput
	// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
	LocalInDenyUnicast pulumi.StringPtrInput
	// Enable/disable local-out logging. Valid values: `enable`, `disable`.
	LocalOut pulumi.StringPtrInput
	// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
	LogInvalidPacket pulumi.StringPtrInput
	// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyComment pulumi.StringPtrInput
	// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyName pulumi.StringPtrInput
	// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
	LogUserInUpper pulumi.StringPtrInput
	// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
	NeighborEvent pulumi.StringPtrInput
	// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
	ResolveIp pulumi.StringPtrInput
	// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
	ResolvePort pulumi.StringPtrInput
	// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
	RestApiGet pulumi.StringPtrInput
	// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
	RestApiSet pulumi.StringPtrInput
	// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
	SyslogOverride pulumi.StringPtrInput
	// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
	UserAnonymize pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSettingState)(nil)).Elem()
}

type logSettingArgs struct {
	// User name anonymization hash salt.
	AnonymizationHash *string `pulumi:"anonymizationHash"`
	// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
	BriefTrafficFormat *string `pulumi:"briefTrafficFormat"`
	// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
	CustomLogFields []LogSettingCustomLogField `pulumi:"customLogFields"`
	// Enable/disable daemon logging. Valid values: `enable`, `disable`.
	DaemonLog *string `pulumi:"daemonLog"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
	ExpolicyImplicitLog *string `pulumi:"expolicyImplicitLog"`
	// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
	FazOverride *string `pulumi:"fazOverride"`
	// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
	Fwpolicy6ImplicitLog *string `pulumi:"fwpolicy6ImplicitLog"`
	// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
	FwpolicyImplicitLog *string `pulumi:"fwpolicyImplicitLog"`
	// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
	LocalInAllow *string `pulumi:"localInAllow"`
	// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
	LocalInDenyBroadcast *string `pulumi:"localInDenyBroadcast"`
	// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
	LocalInDenyUnicast *string `pulumi:"localInDenyUnicast"`
	// Enable/disable local-out logging. Valid values: `enable`, `disable`.
	LocalOut *string `pulumi:"localOut"`
	// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
	LogInvalidPacket *string `pulumi:"logInvalidPacket"`
	// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyComment *string `pulumi:"logPolicyComment"`
	// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyName *string `pulumi:"logPolicyName"`
	// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
	LogUserInUpper *string `pulumi:"logUserInUpper"`
	// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
	NeighborEvent *string `pulumi:"neighborEvent"`
	// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
	ResolveIp *string `pulumi:"resolveIp"`
	// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
	ResolvePort *string `pulumi:"resolvePort"`
	// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
	RestApiGet *string `pulumi:"restApiGet"`
	// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
	RestApiSet *string `pulumi:"restApiSet"`
	// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
	SyslogOverride *string `pulumi:"syslogOverride"`
	// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
	UserAnonymize *string `pulumi:"userAnonymize"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogSetting resource.
type LogSettingArgs struct {
	// User name anonymization hash salt.
	AnonymizationHash pulumi.StringPtrInput
	// Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
	BriefTrafficFormat pulumi.StringPtrInput
	// Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
	CustomLogFields LogSettingCustomLogFieldArrayInput
	// Enable/disable daemon logging. Valid values: `enable`, `disable`.
	DaemonLog pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
	ExpolicyImplicitLog pulumi.StringPtrInput
	// Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
	FazOverride pulumi.StringPtrInput
	// Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
	Fwpolicy6ImplicitLog pulumi.StringPtrInput
	// Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
	FwpolicyImplicitLog pulumi.StringPtrInput
	// Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
	LocalInAllow pulumi.StringPtrInput
	// Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
	LocalInDenyBroadcast pulumi.StringPtrInput
	// Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
	LocalInDenyUnicast pulumi.StringPtrInput
	// Enable/disable local-out logging. Valid values: `enable`, `disable`.
	LocalOut pulumi.StringPtrInput
	// Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
	LogInvalidPacket pulumi.StringPtrInput
	// Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyComment pulumi.StringPtrInput
	// Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
	LogPolicyName pulumi.StringPtrInput
	// Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
	LogUserInUpper pulumi.StringPtrInput
	// Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
	NeighborEvent pulumi.StringPtrInput
	// Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
	ResolveIp pulumi.StringPtrInput
	// Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
	ResolvePort pulumi.StringPtrInput
	// Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
	RestApiGet pulumi.StringPtrInput
	// Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
	RestApiSet pulumi.StringPtrInput
	// Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
	SyslogOverride pulumi.StringPtrInput
	// Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
	UserAnonymize pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSettingArgs)(nil)).Elem()
}

type LogSettingInput interface {
	pulumi.Input

	ToLogSettingOutput() LogSettingOutput
	ToLogSettingOutputWithContext(ctx context.Context) LogSettingOutput
}

func (*LogSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSetting)(nil)).Elem()
}

func (i *LogSetting) ToLogSettingOutput() LogSettingOutput {
	return i.ToLogSettingOutputWithContext(context.Background())
}

func (i *LogSetting) ToLogSettingOutputWithContext(ctx context.Context) LogSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingOutput)
}

// LogSettingArrayInput is an input type that accepts LogSettingArray and LogSettingArrayOutput values.
// You can construct a concrete instance of `LogSettingArrayInput` via:
//
//          LogSettingArray{ LogSettingArgs{...} }
type LogSettingArrayInput interface {
	pulumi.Input

	ToLogSettingArrayOutput() LogSettingArrayOutput
	ToLogSettingArrayOutputWithContext(context.Context) LogSettingArrayOutput
}

type LogSettingArray []LogSettingInput

func (LogSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSetting)(nil)).Elem()
}

func (i LogSettingArray) ToLogSettingArrayOutput() LogSettingArrayOutput {
	return i.ToLogSettingArrayOutputWithContext(context.Background())
}

func (i LogSettingArray) ToLogSettingArrayOutputWithContext(ctx context.Context) LogSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingArrayOutput)
}

// LogSettingMapInput is an input type that accepts LogSettingMap and LogSettingMapOutput values.
// You can construct a concrete instance of `LogSettingMapInput` via:
//
//          LogSettingMap{ "key": LogSettingArgs{...} }
type LogSettingMapInput interface {
	pulumi.Input

	ToLogSettingMapOutput() LogSettingMapOutput
	ToLogSettingMapOutputWithContext(context.Context) LogSettingMapOutput
}

type LogSettingMap map[string]LogSettingInput

func (LogSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSetting)(nil)).Elem()
}

func (i LogSettingMap) ToLogSettingMapOutput() LogSettingMapOutput {
	return i.ToLogSettingMapOutputWithContext(context.Background())
}

func (i LogSettingMap) ToLogSettingMapOutputWithContext(ctx context.Context) LogSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingMapOutput)
}

type LogSettingOutput struct{ *pulumi.OutputState }

func (LogSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSetting)(nil)).Elem()
}

func (o LogSettingOutput) ToLogSettingOutput() LogSettingOutput {
	return o
}

func (o LogSettingOutput) ToLogSettingOutputWithContext(ctx context.Context) LogSettingOutput {
	return o
}

type LogSettingArrayOutput struct{ *pulumi.OutputState }

func (LogSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSetting)(nil)).Elem()
}

func (o LogSettingArrayOutput) ToLogSettingArrayOutput() LogSettingArrayOutput {
	return o
}

func (o LogSettingArrayOutput) ToLogSettingArrayOutputWithContext(ctx context.Context) LogSettingArrayOutput {
	return o
}

func (o LogSettingArrayOutput) Index(i pulumi.IntInput) LogSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSetting {
		return vs[0].([]*LogSetting)[vs[1].(int)]
	}).(LogSettingOutput)
}

type LogSettingMapOutput struct{ *pulumi.OutputState }

func (LogSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSetting)(nil)).Elem()
}

func (o LogSettingMapOutput) ToLogSettingMapOutput() LogSettingMapOutput {
	return o
}

func (o LogSettingMapOutput) ToLogSettingMapOutputWithContext(ctx context.Context) LogSettingMapOutput {
	return o
}

func (o LogSettingMapOutput) MapIndex(k pulumi.StringInput) LogSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSetting {
		return vs[0].(map[string]*LogSetting)[vs[1].(string)]
	}).(LogSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSettingInput)(nil)).Elem(), &LogSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSettingArrayInput)(nil)).Elem(), LogSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSettingMapInput)(nil)).Elem(), LogSettingMap{})
	pulumi.RegisterOutputType(LogSettingOutput{})
	pulumi.RegisterOutputType(LogSettingArrayOutput{})
	pulumi.RegisterOutputType(LogSettingMapOutput{})
}

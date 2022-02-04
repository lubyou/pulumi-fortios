// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings for TACACS+ accounting events filter. Applies to FortiOS Version `>= 7.0.2`.
//
// ## Import
//
// LogTacacsAccounting3 Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logTacacsAccounting3Filter:LogTacacsAccounting3Filter labelname LogTacacsAccounting3Filter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogTacacsAccounting3Filter struct {
	pulumi.CustomResourceState

	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit pulumi.StringOutput `pulumi:"cliCmdAudit"`
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit pulumi.StringOutput `pulumi:"configChangeAudit"`
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit pulumi.StringOutput `pulumi:"loginAudit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogTacacsAccounting3Filter registers a new resource with the given unique name, arguments, and options.
func NewLogTacacsAccounting3Filter(ctx *pulumi.Context,
	name string, args *LogTacacsAccounting3FilterArgs, opts ...pulumi.ResourceOption) (*LogTacacsAccounting3Filter, error) {
	if args == nil {
		args = &LogTacacsAccounting3FilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogTacacsAccounting3Filter
	err := ctx.RegisterResource("fortios:index/logTacacsAccounting3Filter:LogTacacsAccounting3Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogTacacsAccounting3Filter gets an existing LogTacacsAccounting3Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogTacacsAccounting3Filter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogTacacsAccounting3FilterState, opts ...pulumi.ResourceOption) (*LogTacacsAccounting3Filter, error) {
	var resource LogTacacsAccounting3Filter
	err := ctx.ReadResource("fortios:index/logTacacsAccounting3Filter:LogTacacsAccounting3Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogTacacsAccounting3Filter resources.
type logTacacsAccounting3FilterState struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit *string `pulumi:"cliCmdAudit"`
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit *string `pulumi:"configChangeAudit"`
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit *string `pulumi:"loginAudit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogTacacsAccounting3FilterState struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogTacacsAccounting3FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logTacacsAccounting3FilterState)(nil)).Elem()
}

type logTacacsAccounting3FilterArgs struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit *string `pulumi:"cliCmdAudit"`
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit *string `pulumi:"configChangeAudit"`
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit *string `pulumi:"loginAudit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogTacacsAccounting3Filter resource.
type LogTacacsAccounting3FilterArgs struct {
	// Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
	CliCmdAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
	ConfigChangeAudit pulumi.StringPtrInput
	// Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
	LoginAudit pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogTacacsAccounting3FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logTacacsAccounting3FilterArgs)(nil)).Elem()
}

type LogTacacsAccounting3FilterInput interface {
	pulumi.Input

	ToLogTacacsAccounting3FilterOutput() LogTacacsAccounting3FilterOutput
	ToLogTacacsAccounting3FilterOutputWithContext(ctx context.Context) LogTacacsAccounting3FilterOutput
}

func (*LogTacacsAccounting3Filter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogTacacsAccounting3Filter)(nil)).Elem()
}

func (i *LogTacacsAccounting3Filter) ToLogTacacsAccounting3FilterOutput() LogTacacsAccounting3FilterOutput {
	return i.ToLogTacacsAccounting3FilterOutputWithContext(context.Background())
}

func (i *LogTacacsAccounting3Filter) ToLogTacacsAccounting3FilterOutputWithContext(ctx context.Context) LogTacacsAccounting3FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTacacsAccounting3FilterOutput)
}

// LogTacacsAccounting3FilterArrayInput is an input type that accepts LogTacacsAccounting3FilterArray and LogTacacsAccounting3FilterArrayOutput values.
// You can construct a concrete instance of `LogTacacsAccounting3FilterArrayInput` via:
//
//          LogTacacsAccounting3FilterArray{ LogTacacsAccounting3FilterArgs{...} }
type LogTacacsAccounting3FilterArrayInput interface {
	pulumi.Input

	ToLogTacacsAccounting3FilterArrayOutput() LogTacacsAccounting3FilterArrayOutput
	ToLogTacacsAccounting3FilterArrayOutputWithContext(context.Context) LogTacacsAccounting3FilterArrayOutput
}

type LogTacacsAccounting3FilterArray []LogTacacsAccounting3FilterInput

func (LogTacacsAccounting3FilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogTacacsAccounting3Filter)(nil)).Elem()
}

func (i LogTacacsAccounting3FilterArray) ToLogTacacsAccounting3FilterArrayOutput() LogTacacsAccounting3FilterArrayOutput {
	return i.ToLogTacacsAccounting3FilterArrayOutputWithContext(context.Background())
}

func (i LogTacacsAccounting3FilterArray) ToLogTacacsAccounting3FilterArrayOutputWithContext(ctx context.Context) LogTacacsAccounting3FilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTacacsAccounting3FilterArrayOutput)
}

// LogTacacsAccounting3FilterMapInput is an input type that accepts LogTacacsAccounting3FilterMap and LogTacacsAccounting3FilterMapOutput values.
// You can construct a concrete instance of `LogTacacsAccounting3FilterMapInput` via:
//
//          LogTacacsAccounting3FilterMap{ "key": LogTacacsAccounting3FilterArgs{...} }
type LogTacacsAccounting3FilterMapInput interface {
	pulumi.Input

	ToLogTacacsAccounting3FilterMapOutput() LogTacacsAccounting3FilterMapOutput
	ToLogTacacsAccounting3FilterMapOutputWithContext(context.Context) LogTacacsAccounting3FilterMapOutput
}

type LogTacacsAccounting3FilterMap map[string]LogTacacsAccounting3FilterInput

func (LogTacacsAccounting3FilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogTacacsAccounting3Filter)(nil)).Elem()
}

func (i LogTacacsAccounting3FilterMap) ToLogTacacsAccounting3FilterMapOutput() LogTacacsAccounting3FilterMapOutput {
	return i.ToLogTacacsAccounting3FilterMapOutputWithContext(context.Background())
}

func (i LogTacacsAccounting3FilterMap) ToLogTacacsAccounting3FilterMapOutputWithContext(ctx context.Context) LogTacacsAccounting3FilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTacacsAccounting3FilterMapOutput)
}

type LogTacacsAccounting3FilterOutput struct{ *pulumi.OutputState }

func (LogTacacsAccounting3FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogTacacsAccounting3Filter)(nil)).Elem()
}

func (o LogTacacsAccounting3FilterOutput) ToLogTacacsAccounting3FilterOutput() LogTacacsAccounting3FilterOutput {
	return o
}

func (o LogTacacsAccounting3FilterOutput) ToLogTacacsAccounting3FilterOutputWithContext(ctx context.Context) LogTacacsAccounting3FilterOutput {
	return o
}

type LogTacacsAccounting3FilterArrayOutput struct{ *pulumi.OutputState }

func (LogTacacsAccounting3FilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogTacacsAccounting3Filter)(nil)).Elem()
}

func (o LogTacacsAccounting3FilterArrayOutput) ToLogTacacsAccounting3FilterArrayOutput() LogTacacsAccounting3FilterArrayOutput {
	return o
}

func (o LogTacacsAccounting3FilterArrayOutput) ToLogTacacsAccounting3FilterArrayOutputWithContext(ctx context.Context) LogTacacsAccounting3FilterArrayOutput {
	return o
}

func (o LogTacacsAccounting3FilterArrayOutput) Index(i pulumi.IntInput) LogTacacsAccounting3FilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogTacacsAccounting3Filter {
		return vs[0].([]*LogTacacsAccounting3Filter)[vs[1].(int)]
	}).(LogTacacsAccounting3FilterOutput)
}

type LogTacacsAccounting3FilterMapOutput struct{ *pulumi.OutputState }

func (LogTacacsAccounting3FilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogTacacsAccounting3Filter)(nil)).Elem()
}

func (o LogTacacsAccounting3FilterMapOutput) ToLogTacacsAccounting3FilterMapOutput() LogTacacsAccounting3FilterMapOutput {
	return o
}

func (o LogTacacsAccounting3FilterMapOutput) ToLogTacacsAccounting3FilterMapOutputWithContext(ctx context.Context) LogTacacsAccounting3FilterMapOutput {
	return o
}

func (o LogTacacsAccounting3FilterMapOutput) MapIndex(k pulumi.StringInput) LogTacacsAccounting3FilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogTacacsAccounting3Filter {
		return vs[0].(map[string]*LogTacacsAccounting3Filter)[vs[1].(string)]
	}).(LogTacacsAccounting3FilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogTacacsAccounting3FilterInput)(nil)).Elem(), &LogTacacsAccounting3Filter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogTacacsAccounting3FilterArrayInput)(nil)).Elem(), LogTacacsAccounting3FilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogTacacsAccounting3FilterMapInput)(nil)).Elem(), LogTacacsAccounting3FilterMap{})
	pulumi.RegisterOutputType(LogTacacsAccounting3FilterOutput{})
	pulumi.RegisterOutputType(LogTacacsAccounting3FilterArrayOutput{})
	pulumi.RegisterOutputType(LogTacacsAccounting3FilterMapOutput{})
}
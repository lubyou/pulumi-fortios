// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings for WebTrends.
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
// 		_, err := fortios.NewLogWebtrendsSetting(ctx, "trname", &fortios.LogWebtrendsSettingArgs{
// 			Status: pulumi.String("disable"),
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
// LogWebtrends Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logWebtrendsSetting:LogWebtrendsSetting labelname LogWebtrendsSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogWebtrendsSetting struct {
	pulumi.CustomResourceState

	// Address of the remote WebTrends server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogWebtrendsSetting registers a new resource with the given unique name, arguments, and options.
func NewLogWebtrendsSetting(ctx *pulumi.Context,
	name string, args *LogWebtrendsSettingArgs, opts ...pulumi.ResourceOption) (*LogWebtrendsSetting, error) {
	if args == nil {
		args = &LogWebtrendsSettingArgs{}
	}

	var resource LogWebtrendsSetting
	err := ctx.RegisterResource("fortios:index/logWebtrendsSetting:LogWebtrendsSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogWebtrendsSetting gets an existing LogWebtrendsSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogWebtrendsSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogWebtrendsSettingState, opts ...pulumi.ResourceOption) (*LogWebtrendsSetting, error) {
	var resource LogWebtrendsSetting
	err := ctx.ReadResource("fortios:index/logWebtrendsSetting:LogWebtrendsSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogWebtrendsSetting resources.
type logWebtrendsSettingState struct {
	// Address of the remote WebTrends server.
	Server *string `pulumi:"server"`
	// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogWebtrendsSettingState struct {
	// Address of the remote WebTrends server.
	Server pulumi.StringPtrInput
	// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogWebtrendsSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logWebtrendsSettingState)(nil)).Elem()
}

type logWebtrendsSettingArgs struct {
	// Address of the remote WebTrends server.
	Server *string `pulumi:"server"`
	// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogWebtrendsSetting resource.
type LogWebtrendsSettingArgs struct {
	// Address of the remote WebTrends server.
	Server pulumi.StringPtrInput
	// Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogWebtrendsSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logWebtrendsSettingArgs)(nil)).Elem()
}

type LogWebtrendsSettingInput interface {
	pulumi.Input

	ToLogWebtrendsSettingOutput() LogWebtrendsSettingOutput
	ToLogWebtrendsSettingOutputWithContext(ctx context.Context) LogWebtrendsSettingOutput
}

func (*LogWebtrendsSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*LogWebtrendsSetting)(nil))
}

func (i *LogWebtrendsSetting) ToLogWebtrendsSettingOutput() LogWebtrendsSettingOutput {
	return i.ToLogWebtrendsSettingOutputWithContext(context.Background())
}

func (i *LogWebtrendsSetting) ToLogWebtrendsSettingOutputWithContext(ctx context.Context) LogWebtrendsSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogWebtrendsSettingOutput)
}

func (i *LogWebtrendsSetting) ToLogWebtrendsSettingPtrOutput() LogWebtrendsSettingPtrOutput {
	return i.ToLogWebtrendsSettingPtrOutputWithContext(context.Background())
}

func (i *LogWebtrendsSetting) ToLogWebtrendsSettingPtrOutputWithContext(ctx context.Context) LogWebtrendsSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogWebtrendsSettingPtrOutput)
}

type LogWebtrendsSettingPtrInput interface {
	pulumi.Input

	ToLogWebtrendsSettingPtrOutput() LogWebtrendsSettingPtrOutput
	ToLogWebtrendsSettingPtrOutputWithContext(ctx context.Context) LogWebtrendsSettingPtrOutput
}

type logWebtrendsSettingPtrType LogWebtrendsSettingArgs

func (*logWebtrendsSettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogWebtrendsSetting)(nil))
}

func (i *logWebtrendsSettingPtrType) ToLogWebtrendsSettingPtrOutput() LogWebtrendsSettingPtrOutput {
	return i.ToLogWebtrendsSettingPtrOutputWithContext(context.Background())
}

func (i *logWebtrendsSettingPtrType) ToLogWebtrendsSettingPtrOutputWithContext(ctx context.Context) LogWebtrendsSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogWebtrendsSettingPtrOutput)
}

// LogWebtrendsSettingArrayInput is an input type that accepts LogWebtrendsSettingArray and LogWebtrendsSettingArrayOutput values.
// You can construct a concrete instance of `LogWebtrendsSettingArrayInput` via:
//
//          LogWebtrendsSettingArray{ LogWebtrendsSettingArgs{...} }
type LogWebtrendsSettingArrayInput interface {
	pulumi.Input

	ToLogWebtrendsSettingArrayOutput() LogWebtrendsSettingArrayOutput
	ToLogWebtrendsSettingArrayOutputWithContext(context.Context) LogWebtrendsSettingArrayOutput
}

type LogWebtrendsSettingArray []LogWebtrendsSettingInput

func (LogWebtrendsSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogWebtrendsSetting)(nil))
}

func (i LogWebtrendsSettingArray) ToLogWebtrendsSettingArrayOutput() LogWebtrendsSettingArrayOutput {
	return i.ToLogWebtrendsSettingArrayOutputWithContext(context.Background())
}

func (i LogWebtrendsSettingArray) ToLogWebtrendsSettingArrayOutputWithContext(ctx context.Context) LogWebtrendsSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogWebtrendsSettingArrayOutput)
}

// LogWebtrendsSettingMapInput is an input type that accepts LogWebtrendsSettingMap and LogWebtrendsSettingMapOutput values.
// You can construct a concrete instance of `LogWebtrendsSettingMapInput` via:
//
//          LogWebtrendsSettingMap{ "key": LogWebtrendsSettingArgs{...} }
type LogWebtrendsSettingMapInput interface {
	pulumi.Input

	ToLogWebtrendsSettingMapOutput() LogWebtrendsSettingMapOutput
	ToLogWebtrendsSettingMapOutputWithContext(context.Context) LogWebtrendsSettingMapOutput
}

type LogWebtrendsSettingMap map[string]LogWebtrendsSettingInput

func (LogWebtrendsSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogWebtrendsSetting)(nil))
}

func (i LogWebtrendsSettingMap) ToLogWebtrendsSettingMapOutput() LogWebtrendsSettingMapOutput {
	return i.ToLogWebtrendsSettingMapOutputWithContext(context.Background())
}

func (i LogWebtrendsSettingMap) ToLogWebtrendsSettingMapOutputWithContext(ctx context.Context) LogWebtrendsSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogWebtrendsSettingMapOutput)
}

type LogWebtrendsSettingOutput struct {
	*pulumi.OutputState
}

func (LogWebtrendsSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogWebtrendsSetting)(nil))
}

func (o LogWebtrendsSettingOutput) ToLogWebtrendsSettingOutput() LogWebtrendsSettingOutput {
	return o
}

func (o LogWebtrendsSettingOutput) ToLogWebtrendsSettingOutputWithContext(ctx context.Context) LogWebtrendsSettingOutput {
	return o
}

func (o LogWebtrendsSettingOutput) ToLogWebtrendsSettingPtrOutput() LogWebtrendsSettingPtrOutput {
	return o.ToLogWebtrendsSettingPtrOutputWithContext(context.Background())
}

func (o LogWebtrendsSettingOutput) ToLogWebtrendsSettingPtrOutputWithContext(ctx context.Context) LogWebtrendsSettingPtrOutput {
	return o.ApplyT(func(v LogWebtrendsSetting) *LogWebtrendsSetting {
		return &v
	}).(LogWebtrendsSettingPtrOutput)
}

type LogWebtrendsSettingPtrOutput struct {
	*pulumi.OutputState
}

func (LogWebtrendsSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogWebtrendsSetting)(nil))
}

func (o LogWebtrendsSettingPtrOutput) ToLogWebtrendsSettingPtrOutput() LogWebtrendsSettingPtrOutput {
	return o
}

func (o LogWebtrendsSettingPtrOutput) ToLogWebtrendsSettingPtrOutputWithContext(ctx context.Context) LogWebtrendsSettingPtrOutput {
	return o
}

type LogWebtrendsSettingArrayOutput struct{ *pulumi.OutputState }

func (LogWebtrendsSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogWebtrendsSetting)(nil))
}

func (o LogWebtrendsSettingArrayOutput) ToLogWebtrendsSettingArrayOutput() LogWebtrendsSettingArrayOutput {
	return o
}

func (o LogWebtrendsSettingArrayOutput) ToLogWebtrendsSettingArrayOutputWithContext(ctx context.Context) LogWebtrendsSettingArrayOutput {
	return o
}

func (o LogWebtrendsSettingArrayOutput) Index(i pulumi.IntInput) LogWebtrendsSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogWebtrendsSetting {
		return vs[0].([]LogWebtrendsSetting)[vs[1].(int)]
	}).(LogWebtrendsSettingOutput)
}

type LogWebtrendsSettingMapOutput struct{ *pulumi.OutputState }

func (LogWebtrendsSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogWebtrendsSetting)(nil))
}

func (o LogWebtrendsSettingMapOutput) ToLogWebtrendsSettingMapOutput() LogWebtrendsSettingMapOutput {
	return o
}

func (o LogWebtrendsSettingMapOutput) ToLogWebtrendsSettingMapOutputWithContext(ctx context.Context) LogWebtrendsSettingMapOutput {
	return o
}

func (o LogWebtrendsSettingMapOutput) MapIndex(k pulumi.StringInput) LogWebtrendsSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogWebtrendsSetting {
		return vs[0].(map[string]LogWebtrendsSetting)[vs[1].(string)]
	}).(LogWebtrendsSettingOutput)
}

func init() {
	pulumi.RegisterOutputType(LogWebtrendsSettingOutput{})
	pulumi.RegisterOutputType(LogWebtrendsSettingPtrOutput{})
	pulumi.RegisterOutputType(LogWebtrendsSettingArrayOutput{})
	pulumi.RegisterOutputType(LogWebtrendsSettingMapOutput{})
}
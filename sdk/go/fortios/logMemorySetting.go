// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings for memory buffer.
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
// 		_, err := fortios.NewLogMemorySetting(ctx, "trname", &fortios.LogMemorySettingArgs{
// 			Diskfull: pulumi.String("overwrite"),
// 			Status:   pulumi.String("disable"),
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
// LogMemory Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logMemorySetting:LogMemorySetting labelname LogMemorySetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogMemorySetting struct {
	pulumi.CustomResourceState

	// Action to take when memory is full. Valid values: `overwrite`.
	Diskfull pulumi.StringOutput `pulumi:"diskfull"`
	// Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogMemorySetting registers a new resource with the given unique name, arguments, and options.
func NewLogMemorySetting(ctx *pulumi.Context,
	name string, args *LogMemorySettingArgs, opts ...pulumi.ResourceOption) (*LogMemorySetting, error) {
	if args == nil {
		args = &LogMemorySettingArgs{}
	}

	var resource LogMemorySetting
	err := ctx.RegisterResource("fortios:index/logMemorySetting:LogMemorySetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogMemorySetting gets an existing LogMemorySetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogMemorySetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogMemorySettingState, opts ...pulumi.ResourceOption) (*LogMemorySetting, error) {
	var resource LogMemorySetting
	err := ctx.ReadResource("fortios:index/logMemorySetting:LogMemorySetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogMemorySetting resources.
type logMemorySettingState struct {
	// Action to take when memory is full. Valid values: `overwrite`.
	Diskfull *string `pulumi:"diskfull"`
	// Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogMemorySettingState struct {
	// Action to take when memory is full. Valid values: `overwrite`.
	Diskfull pulumi.StringPtrInput
	// Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogMemorySettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logMemorySettingState)(nil)).Elem()
}

type logMemorySettingArgs struct {
	// Action to take when memory is full. Valid values: `overwrite`.
	Diskfull *string `pulumi:"diskfull"`
	// Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogMemorySetting resource.
type LogMemorySettingArgs struct {
	// Action to take when memory is full. Valid values: `overwrite`.
	Diskfull pulumi.StringPtrInput
	// Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogMemorySettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logMemorySettingArgs)(nil)).Elem()
}

type LogMemorySettingInput interface {
	pulumi.Input

	ToLogMemorySettingOutput() LogMemorySettingOutput
	ToLogMemorySettingOutputWithContext(ctx context.Context) LogMemorySettingOutput
}

func (*LogMemorySetting) ElementType() reflect.Type {
	return reflect.TypeOf((*LogMemorySetting)(nil))
}

func (i *LogMemorySetting) ToLogMemorySettingOutput() LogMemorySettingOutput {
	return i.ToLogMemorySettingOutputWithContext(context.Background())
}

func (i *LogMemorySetting) ToLogMemorySettingOutputWithContext(ctx context.Context) LogMemorySettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMemorySettingOutput)
}

func (i *LogMemorySetting) ToLogMemorySettingPtrOutput() LogMemorySettingPtrOutput {
	return i.ToLogMemorySettingPtrOutputWithContext(context.Background())
}

func (i *LogMemorySetting) ToLogMemorySettingPtrOutputWithContext(ctx context.Context) LogMemorySettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMemorySettingPtrOutput)
}

type LogMemorySettingPtrInput interface {
	pulumi.Input

	ToLogMemorySettingPtrOutput() LogMemorySettingPtrOutput
	ToLogMemorySettingPtrOutputWithContext(ctx context.Context) LogMemorySettingPtrOutput
}

type logMemorySettingPtrType LogMemorySettingArgs

func (*logMemorySettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogMemorySetting)(nil))
}

func (i *logMemorySettingPtrType) ToLogMemorySettingPtrOutput() LogMemorySettingPtrOutput {
	return i.ToLogMemorySettingPtrOutputWithContext(context.Background())
}

func (i *logMemorySettingPtrType) ToLogMemorySettingPtrOutputWithContext(ctx context.Context) LogMemorySettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMemorySettingPtrOutput)
}

// LogMemorySettingArrayInput is an input type that accepts LogMemorySettingArray and LogMemorySettingArrayOutput values.
// You can construct a concrete instance of `LogMemorySettingArrayInput` via:
//
//          LogMemorySettingArray{ LogMemorySettingArgs{...} }
type LogMemorySettingArrayInput interface {
	pulumi.Input

	ToLogMemorySettingArrayOutput() LogMemorySettingArrayOutput
	ToLogMemorySettingArrayOutputWithContext(context.Context) LogMemorySettingArrayOutput
}

type LogMemorySettingArray []LogMemorySettingInput

func (LogMemorySettingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogMemorySetting)(nil))
}

func (i LogMemorySettingArray) ToLogMemorySettingArrayOutput() LogMemorySettingArrayOutput {
	return i.ToLogMemorySettingArrayOutputWithContext(context.Background())
}

func (i LogMemorySettingArray) ToLogMemorySettingArrayOutputWithContext(ctx context.Context) LogMemorySettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMemorySettingArrayOutput)
}

// LogMemorySettingMapInput is an input type that accepts LogMemorySettingMap and LogMemorySettingMapOutput values.
// You can construct a concrete instance of `LogMemorySettingMapInput` via:
//
//          LogMemorySettingMap{ "key": LogMemorySettingArgs{...} }
type LogMemorySettingMapInput interface {
	pulumi.Input

	ToLogMemorySettingMapOutput() LogMemorySettingMapOutput
	ToLogMemorySettingMapOutputWithContext(context.Context) LogMemorySettingMapOutput
}

type LogMemorySettingMap map[string]LogMemorySettingInput

func (LogMemorySettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogMemorySetting)(nil))
}

func (i LogMemorySettingMap) ToLogMemorySettingMapOutput() LogMemorySettingMapOutput {
	return i.ToLogMemorySettingMapOutputWithContext(context.Background())
}

func (i LogMemorySettingMap) ToLogMemorySettingMapOutputWithContext(ctx context.Context) LogMemorySettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMemorySettingMapOutput)
}

type LogMemorySettingOutput struct {
	*pulumi.OutputState
}

func (LogMemorySettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogMemorySetting)(nil))
}

func (o LogMemorySettingOutput) ToLogMemorySettingOutput() LogMemorySettingOutput {
	return o
}

func (o LogMemorySettingOutput) ToLogMemorySettingOutputWithContext(ctx context.Context) LogMemorySettingOutput {
	return o
}

func (o LogMemorySettingOutput) ToLogMemorySettingPtrOutput() LogMemorySettingPtrOutput {
	return o.ToLogMemorySettingPtrOutputWithContext(context.Background())
}

func (o LogMemorySettingOutput) ToLogMemorySettingPtrOutputWithContext(ctx context.Context) LogMemorySettingPtrOutput {
	return o.ApplyT(func(v LogMemorySetting) *LogMemorySetting {
		return &v
	}).(LogMemorySettingPtrOutput)
}

type LogMemorySettingPtrOutput struct {
	*pulumi.OutputState
}

func (LogMemorySettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogMemorySetting)(nil))
}

func (o LogMemorySettingPtrOutput) ToLogMemorySettingPtrOutput() LogMemorySettingPtrOutput {
	return o
}

func (o LogMemorySettingPtrOutput) ToLogMemorySettingPtrOutputWithContext(ctx context.Context) LogMemorySettingPtrOutput {
	return o
}

type LogMemorySettingArrayOutput struct{ *pulumi.OutputState }

func (LogMemorySettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogMemorySetting)(nil))
}

func (o LogMemorySettingArrayOutput) ToLogMemorySettingArrayOutput() LogMemorySettingArrayOutput {
	return o
}

func (o LogMemorySettingArrayOutput) ToLogMemorySettingArrayOutputWithContext(ctx context.Context) LogMemorySettingArrayOutput {
	return o
}

func (o LogMemorySettingArrayOutput) Index(i pulumi.IntInput) LogMemorySettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogMemorySetting {
		return vs[0].([]LogMemorySetting)[vs[1].(int)]
	}).(LogMemorySettingOutput)
}

type LogMemorySettingMapOutput struct{ *pulumi.OutputState }

func (LogMemorySettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogMemorySetting)(nil))
}

func (o LogMemorySettingMapOutput) ToLogMemorySettingMapOutput() LogMemorySettingMapOutput {
	return o
}

func (o LogMemorySettingMapOutput) ToLogMemorySettingMapOutputWithContext(ctx context.Context) LogMemorySettingMapOutput {
	return o
}

func (o LogMemorySettingMapOutput) MapIndex(k pulumi.StringInput) LogMemorySettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogMemorySetting {
		return vs[0].(map[string]LogMemorySetting)[vs[1].(string)]
	}).(LogMemorySettingOutput)
}

func init() {
	pulumi.RegisterOutputType(LogMemorySettingOutput{})
	pulumi.RegisterOutputType(LogMemorySettingPtrOutput{})
	pulumi.RegisterOutputType(LogMemorySettingArrayOutput{})
	pulumi.RegisterOutputType(LogMemorySettingMapOutput{})
}

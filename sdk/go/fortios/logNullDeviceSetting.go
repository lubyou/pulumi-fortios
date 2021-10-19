// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings for null device logging.
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
// 		_, err := fortios.NewLogNullDeviceSetting(ctx, "trname", &fortios.LogNullDeviceSettingArgs{
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
// LogNullDevice Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logNullDeviceSetting:LogNullDeviceSetting labelname LogNullDeviceSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogNullDeviceSetting struct {
	pulumi.CustomResourceState

	// Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved). Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogNullDeviceSetting registers a new resource with the given unique name, arguments, and options.
func NewLogNullDeviceSetting(ctx *pulumi.Context,
	name string, args *LogNullDeviceSettingArgs, opts ...pulumi.ResourceOption) (*LogNullDeviceSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	var resource LogNullDeviceSetting
	err := ctx.RegisterResource("fortios:index/logNullDeviceSetting:LogNullDeviceSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogNullDeviceSetting gets an existing LogNullDeviceSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogNullDeviceSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogNullDeviceSettingState, opts ...pulumi.ResourceOption) (*LogNullDeviceSetting, error) {
	var resource LogNullDeviceSetting
	err := ctx.ReadResource("fortios:index/logNullDeviceSetting:LogNullDeviceSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogNullDeviceSetting resources.
type logNullDeviceSettingState struct {
	// Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved). Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogNullDeviceSettingState struct {
	// Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved). Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogNullDeviceSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logNullDeviceSettingState)(nil)).Elem()
}

type logNullDeviceSettingArgs struct {
	// Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved). Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogNullDeviceSetting resource.
type LogNullDeviceSettingArgs struct {
	// Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved). Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogNullDeviceSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logNullDeviceSettingArgs)(nil)).Elem()
}

type LogNullDeviceSettingInput interface {
	pulumi.Input

	ToLogNullDeviceSettingOutput() LogNullDeviceSettingOutput
	ToLogNullDeviceSettingOutputWithContext(ctx context.Context) LogNullDeviceSettingOutput
}

func (*LogNullDeviceSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*LogNullDeviceSetting)(nil))
}

func (i *LogNullDeviceSetting) ToLogNullDeviceSettingOutput() LogNullDeviceSettingOutput {
	return i.ToLogNullDeviceSettingOutputWithContext(context.Background())
}

func (i *LogNullDeviceSetting) ToLogNullDeviceSettingOutputWithContext(ctx context.Context) LogNullDeviceSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogNullDeviceSettingOutput)
}

func (i *LogNullDeviceSetting) ToLogNullDeviceSettingPtrOutput() LogNullDeviceSettingPtrOutput {
	return i.ToLogNullDeviceSettingPtrOutputWithContext(context.Background())
}

func (i *LogNullDeviceSetting) ToLogNullDeviceSettingPtrOutputWithContext(ctx context.Context) LogNullDeviceSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogNullDeviceSettingPtrOutput)
}

type LogNullDeviceSettingPtrInput interface {
	pulumi.Input

	ToLogNullDeviceSettingPtrOutput() LogNullDeviceSettingPtrOutput
	ToLogNullDeviceSettingPtrOutputWithContext(ctx context.Context) LogNullDeviceSettingPtrOutput
}

type logNullDeviceSettingPtrType LogNullDeviceSettingArgs

func (*logNullDeviceSettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogNullDeviceSetting)(nil))
}

func (i *logNullDeviceSettingPtrType) ToLogNullDeviceSettingPtrOutput() LogNullDeviceSettingPtrOutput {
	return i.ToLogNullDeviceSettingPtrOutputWithContext(context.Background())
}

func (i *logNullDeviceSettingPtrType) ToLogNullDeviceSettingPtrOutputWithContext(ctx context.Context) LogNullDeviceSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogNullDeviceSettingPtrOutput)
}

// LogNullDeviceSettingArrayInput is an input type that accepts LogNullDeviceSettingArray and LogNullDeviceSettingArrayOutput values.
// You can construct a concrete instance of `LogNullDeviceSettingArrayInput` via:
//
//          LogNullDeviceSettingArray{ LogNullDeviceSettingArgs{...} }
type LogNullDeviceSettingArrayInput interface {
	pulumi.Input

	ToLogNullDeviceSettingArrayOutput() LogNullDeviceSettingArrayOutput
	ToLogNullDeviceSettingArrayOutputWithContext(context.Context) LogNullDeviceSettingArrayOutput
}

type LogNullDeviceSettingArray []LogNullDeviceSettingInput

func (LogNullDeviceSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogNullDeviceSetting)(nil))
}

func (i LogNullDeviceSettingArray) ToLogNullDeviceSettingArrayOutput() LogNullDeviceSettingArrayOutput {
	return i.ToLogNullDeviceSettingArrayOutputWithContext(context.Background())
}

func (i LogNullDeviceSettingArray) ToLogNullDeviceSettingArrayOutputWithContext(ctx context.Context) LogNullDeviceSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogNullDeviceSettingArrayOutput)
}

// LogNullDeviceSettingMapInput is an input type that accepts LogNullDeviceSettingMap and LogNullDeviceSettingMapOutput values.
// You can construct a concrete instance of `LogNullDeviceSettingMapInput` via:
//
//          LogNullDeviceSettingMap{ "key": LogNullDeviceSettingArgs{...} }
type LogNullDeviceSettingMapInput interface {
	pulumi.Input

	ToLogNullDeviceSettingMapOutput() LogNullDeviceSettingMapOutput
	ToLogNullDeviceSettingMapOutputWithContext(context.Context) LogNullDeviceSettingMapOutput
}

type LogNullDeviceSettingMap map[string]LogNullDeviceSettingInput

func (LogNullDeviceSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogNullDeviceSetting)(nil))
}

func (i LogNullDeviceSettingMap) ToLogNullDeviceSettingMapOutput() LogNullDeviceSettingMapOutput {
	return i.ToLogNullDeviceSettingMapOutputWithContext(context.Background())
}

func (i LogNullDeviceSettingMap) ToLogNullDeviceSettingMapOutputWithContext(ctx context.Context) LogNullDeviceSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogNullDeviceSettingMapOutput)
}

type LogNullDeviceSettingOutput struct {
	*pulumi.OutputState
}

func (LogNullDeviceSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogNullDeviceSetting)(nil))
}

func (o LogNullDeviceSettingOutput) ToLogNullDeviceSettingOutput() LogNullDeviceSettingOutput {
	return o
}

func (o LogNullDeviceSettingOutput) ToLogNullDeviceSettingOutputWithContext(ctx context.Context) LogNullDeviceSettingOutput {
	return o
}

func (o LogNullDeviceSettingOutput) ToLogNullDeviceSettingPtrOutput() LogNullDeviceSettingPtrOutput {
	return o.ToLogNullDeviceSettingPtrOutputWithContext(context.Background())
}

func (o LogNullDeviceSettingOutput) ToLogNullDeviceSettingPtrOutputWithContext(ctx context.Context) LogNullDeviceSettingPtrOutput {
	return o.ApplyT(func(v LogNullDeviceSetting) *LogNullDeviceSetting {
		return &v
	}).(LogNullDeviceSettingPtrOutput)
}

type LogNullDeviceSettingPtrOutput struct {
	*pulumi.OutputState
}

func (LogNullDeviceSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogNullDeviceSetting)(nil))
}

func (o LogNullDeviceSettingPtrOutput) ToLogNullDeviceSettingPtrOutput() LogNullDeviceSettingPtrOutput {
	return o
}

func (o LogNullDeviceSettingPtrOutput) ToLogNullDeviceSettingPtrOutputWithContext(ctx context.Context) LogNullDeviceSettingPtrOutput {
	return o
}

type LogNullDeviceSettingArrayOutput struct{ *pulumi.OutputState }

func (LogNullDeviceSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogNullDeviceSetting)(nil))
}

func (o LogNullDeviceSettingArrayOutput) ToLogNullDeviceSettingArrayOutput() LogNullDeviceSettingArrayOutput {
	return o
}

func (o LogNullDeviceSettingArrayOutput) ToLogNullDeviceSettingArrayOutputWithContext(ctx context.Context) LogNullDeviceSettingArrayOutput {
	return o
}

func (o LogNullDeviceSettingArrayOutput) Index(i pulumi.IntInput) LogNullDeviceSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogNullDeviceSetting {
		return vs[0].([]LogNullDeviceSetting)[vs[1].(int)]
	}).(LogNullDeviceSettingOutput)
}

type LogNullDeviceSettingMapOutput struct{ *pulumi.OutputState }

func (LogNullDeviceSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogNullDeviceSetting)(nil))
}

func (o LogNullDeviceSettingMapOutput) ToLogNullDeviceSettingMapOutput() LogNullDeviceSettingMapOutput {
	return o
}

func (o LogNullDeviceSettingMapOutput) ToLogNullDeviceSettingMapOutputWithContext(ctx context.Context) LogNullDeviceSettingMapOutput {
	return o
}

func (o LogNullDeviceSettingMapOutput) MapIndex(k pulumi.StringInput) LogNullDeviceSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogNullDeviceSetting {
		return vs[0].(map[string]LogNullDeviceSetting)[vs[1].(string)]
	}).(LogNullDeviceSettingOutput)
}

func init() {
	pulumi.RegisterOutputType(LogNullDeviceSettingOutput{})
	pulumi.RegisterOutputType(LogNullDeviceSettingPtrOutput{})
	pulumi.RegisterOutputType(LogNullDeviceSettingArrayOutput{})
	pulumi.RegisterOutputType(LogNullDeviceSettingMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch spanning tree protocol (STP).
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
// 		_, err := fortios.NewSwitchControllerStpSettings(ctx, "trname", &fortios.SwitchControllerStpSettingsArgs{
// 			ForwardTime:  pulumi.Int(15),
// 			HelloTime:    pulumi.Int(2),
// 			MaxAge:       pulumi.Int(20),
// 			MaxHops:      pulumi.Int(20),
// 			PendingTimer: pulumi.Int(4),
// 			Revision:     pulumi.Int(0),
// 			Status:       pulumi.String("enable"),
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
// SwitchController StpSettings can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/switchControllerStpSettings:SwitchControllerStpSettings labelname SwitchControllerStpSettings
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerStpSettings:SwitchControllerStpSettings labelname SwitchControllerStpSettings
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerStpSettings struct {
	pulumi.CustomResourceState

	// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
	ForwardTime pulumi.IntOutput `pulumi:"forwardTime"`
	// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
	HelloTime pulumi.IntOutput `pulumi:"helloTime"`
	// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
	MaxAge pulumi.IntOutput `pulumi:"maxAge"`
	// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
	MaxHops pulumi.IntOutput `pulumi:"maxHops"`
	// Name of global STP settings configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Pending time (1 - 15 sec, default = 4).
	PendingTimer pulumi.IntOutput `pulumi:"pendingTimer"`
	// STP revision number (0 - 65535).
	Revision pulumi.IntOutput `pulumi:"revision"`
	// Enable/disable STP. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerStpSettings registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerStpSettings(ctx *pulumi.Context,
	name string, args *SwitchControllerStpSettingsArgs, opts ...pulumi.ResourceOption) (*SwitchControllerStpSettings, error) {
	if args == nil {
		args = &SwitchControllerStpSettingsArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerStpSettings
	err := ctx.RegisterResource("fortios:index/switchControllerStpSettings:SwitchControllerStpSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerStpSettings gets an existing SwitchControllerStpSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerStpSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerStpSettingsState, opts ...pulumi.ResourceOption) (*SwitchControllerStpSettings, error) {
	var resource SwitchControllerStpSettings
	err := ctx.ReadResource("fortios:index/switchControllerStpSettings:SwitchControllerStpSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerStpSettings resources.
type switchControllerStpSettingsState struct {
	// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
	ForwardTime *int `pulumi:"forwardTime"`
	// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
	HelloTime *int `pulumi:"helloTime"`
	// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
	MaxAge *int `pulumi:"maxAge"`
	// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
	MaxHops *int `pulumi:"maxHops"`
	// Name of global STP settings configuration.
	Name *string `pulumi:"name"`
	// Pending time (1 - 15 sec, default = 4).
	PendingTimer *int `pulumi:"pendingTimer"`
	// STP revision number (0 - 65535).
	Revision *int `pulumi:"revision"`
	// Enable/disable STP. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerStpSettingsState struct {
	// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
	ForwardTime pulumi.IntPtrInput
	// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
	HelloTime pulumi.IntPtrInput
	// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
	MaxAge pulumi.IntPtrInput
	// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
	MaxHops pulumi.IntPtrInput
	// Name of global STP settings configuration.
	Name pulumi.StringPtrInput
	// Pending time (1 - 15 sec, default = 4).
	PendingTimer pulumi.IntPtrInput
	// STP revision number (0 - 65535).
	Revision pulumi.IntPtrInput
	// Enable/disable STP. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerStpSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerStpSettingsState)(nil)).Elem()
}

type switchControllerStpSettingsArgs struct {
	// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
	ForwardTime *int `pulumi:"forwardTime"`
	// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
	HelloTime *int `pulumi:"helloTime"`
	// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
	MaxAge *int `pulumi:"maxAge"`
	// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
	MaxHops *int `pulumi:"maxHops"`
	// Name of global STP settings configuration.
	Name *string `pulumi:"name"`
	// Pending time (1 - 15 sec, default = 4).
	PendingTimer *int `pulumi:"pendingTimer"`
	// STP revision number (0 - 65535).
	Revision *int `pulumi:"revision"`
	// Enable/disable STP. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerStpSettings resource.
type SwitchControllerStpSettingsArgs struct {
	// Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
	ForwardTime pulumi.IntPtrInput
	// Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
	HelloTime pulumi.IntPtrInput
	// Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
	MaxAge pulumi.IntPtrInput
	// Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
	MaxHops pulumi.IntPtrInput
	// Name of global STP settings configuration.
	Name pulumi.StringPtrInput
	// Pending time (1 - 15 sec, default = 4).
	PendingTimer pulumi.IntPtrInput
	// STP revision number (0 - 65535).
	Revision pulumi.IntPtrInput
	// Enable/disable STP. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerStpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerStpSettingsArgs)(nil)).Elem()
}

type SwitchControllerStpSettingsInput interface {
	pulumi.Input

	ToSwitchControllerStpSettingsOutput() SwitchControllerStpSettingsOutput
	ToSwitchControllerStpSettingsOutputWithContext(ctx context.Context) SwitchControllerStpSettingsOutput
}

func (*SwitchControllerStpSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerStpSettings)(nil)).Elem()
}

func (i *SwitchControllerStpSettings) ToSwitchControllerStpSettingsOutput() SwitchControllerStpSettingsOutput {
	return i.ToSwitchControllerStpSettingsOutputWithContext(context.Background())
}

func (i *SwitchControllerStpSettings) ToSwitchControllerStpSettingsOutputWithContext(ctx context.Context) SwitchControllerStpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerStpSettingsOutput)
}

// SwitchControllerStpSettingsArrayInput is an input type that accepts SwitchControllerStpSettingsArray and SwitchControllerStpSettingsArrayOutput values.
// You can construct a concrete instance of `SwitchControllerStpSettingsArrayInput` via:
//
//          SwitchControllerStpSettingsArray{ SwitchControllerStpSettingsArgs{...} }
type SwitchControllerStpSettingsArrayInput interface {
	pulumi.Input

	ToSwitchControllerStpSettingsArrayOutput() SwitchControllerStpSettingsArrayOutput
	ToSwitchControllerStpSettingsArrayOutputWithContext(context.Context) SwitchControllerStpSettingsArrayOutput
}

type SwitchControllerStpSettingsArray []SwitchControllerStpSettingsInput

func (SwitchControllerStpSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerStpSettings)(nil)).Elem()
}

func (i SwitchControllerStpSettingsArray) ToSwitchControllerStpSettingsArrayOutput() SwitchControllerStpSettingsArrayOutput {
	return i.ToSwitchControllerStpSettingsArrayOutputWithContext(context.Background())
}

func (i SwitchControllerStpSettingsArray) ToSwitchControllerStpSettingsArrayOutputWithContext(ctx context.Context) SwitchControllerStpSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerStpSettingsArrayOutput)
}

// SwitchControllerStpSettingsMapInput is an input type that accepts SwitchControllerStpSettingsMap and SwitchControllerStpSettingsMapOutput values.
// You can construct a concrete instance of `SwitchControllerStpSettingsMapInput` via:
//
//          SwitchControllerStpSettingsMap{ "key": SwitchControllerStpSettingsArgs{...} }
type SwitchControllerStpSettingsMapInput interface {
	pulumi.Input

	ToSwitchControllerStpSettingsMapOutput() SwitchControllerStpSettingsMapOutput
	ToSwitchControllerStpSettingsMapOutputWithContext(context.Context) SwitchControllerStpSettingsMapOutput
}

type SwitchControllerStpSettingsMap map[string]SwitchControllerStpSettingsInput

func (SwitchControllerStpSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerStpSettings)(nil)).Elem()
}

func (i SwitchControllerStpSettingsMap) ToSwitchControllerStpSettingsMapOutput() SwitchControllerStpSettingsMapOutput {
	return i.ToSwitchControllerStpSettingsMapOutputWithContext(context.Background())
}

func (i SwitchControllerStpSettingsMap) ToSwitchControllerStpSettingsMapOutputWithContext(ctx context.Context) SwitchControllerStpSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerStpSettingsMapOutput)
}

type SwitchControllerStpSettingsOutput struct{ *pulumi.OutputState }

func (SwitchControllerStpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerStpSettings)(nil)).Elem()
}

func (o SwitchControllerStpSettingsOutput) ToSwitchControllerStpSettingsOutput() SwitchControllerStpSettingsOutput {
	return o
}

func (o SwitchControllerStpSettingsOutput) ToSwitchControllerStpSettingsOutputWithContext(ctx context.Context) SwitchControllerStpSettingsOutput {
	return o
}

type SwitchControllerStpSettingsArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerStpSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerStpSettings)(nil)).Elem()
}

func (o SwitchControllerStpSettingsArrayOutput) ToSwitchControllerStpSettingsArrayOutput() SwitchControllerStpSettingsArrayOutput {
	return o
}

func (o SwitchControllerStpSettingsArrayOutput) ToSwitchControllerStpSettingsArrayOutputWithContext(ctx context.Context) SwitchControllerStpSettingsArrayOutput {
	return o
}

func (o SwitchControllerStpSettingsArrayOutput) Index(i pulumi.IntInput) SwitchControllerStpSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerStpSettings {
		return vs[0].([]*SwitchControllerStpSettings)[vs[1].(int)]
	}).(SwitchControllerStpSettingsOutput)
}

type SwitchControllerStpSettingsMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerStpSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerStpSettings)(nil)).Elem()
}

func (o SwitchControllerStpSettingsMapOutput) ToSwitchControllerStpSettingsMapOutput() SwitchControllerStpSettingsMapOutput {
	return o
}

func (o SwitchControllerStpSettingsMapOutput) ToSwitchControllerStpSettingsMapOutputWithContext(ctx context.Context) SwitchControllerStpSettingsMapOutput {
	return o
}

func (o SwitchControllerStpSettingsMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerStpSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerStpSettings {
		return vs[0].(map[string]*SwitchControllerStpSettings)[vs[1].(string)]
	}).(SwitchControllerStpSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerStpSettingsInput)(nil)).Elem(), &SwitchControllerStpSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerStpSettingsArrayInput)(nil)).Elem(), SwitchControllerStpSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerStpSettingsMapInput)(nil)).Elem(), SwitchControllerStpSettingsMap{})
	pulumi.RegisterOutputType(SwitchControllerStpSettingsOutput{})
	pulumi.RegisterOutputType(SwitchControllerStpSettingsArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerStpSettingsMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure HA monitor.
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
// 		_, err := fortios.NewSystemHaMonitor(ctx, "trname", &fortios.SystemHaMonitorArgs{
// 			MonitorVlan:         pulumi.String("disable"),
// 			VlanHbInterval:      pulumi.Int(5),
// 			VlanHbLostThreshold: pulumi.Int(3),
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
// System HaMonitor can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemHaMonitor:SystemHaMonitor labelname SystemHaMonitor
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemHaMonitor:SystemHaMonitor labelname SystemHaMonitor
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemHaMonitor struct {
	pulumi.CustomResourceState

	// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
	MonitorVlan pulumi.StringOutput `pulumi:"monitorVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Configure heartbeat interval (seconds).
	VlanHbInterval pulumi.IntOutput `pulumi:"vlanHbInterval"`
	// VLAN lost heartbeat threshold (1 - 60).
	VlanHbLostThreshold pulumi.IntOutput `pulumi:"vlanHbLostThreshold"`
}

// NewSystemHaMonitor registers a new resource with the given unique name, arguments, and options.
func NewSystemHaMonitor(ctx *pulumi.Context,
	name string, args *SystemHaMonitorArgs, opts ...pulumi.ResourceOption) (*SystemHaMonitor, error) {
	if args == nil {
		args = &SystemHaMonitorArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemHaMonitor
	err := ctx.RegisterResource("fortios:index/systemHaMonitor:SystemHaMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemHaMonitor gets an existing SystemHaMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemHaMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemHaMonitorState, opts ...pulumi.ResourceOption) (*SystemHaMonitor, error) {
	var resource SystemHaMonitor
	err := ctx.ReadResource("fortios:index/systemHaMonitor:SystemHaMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemHaMonitor resources.
type systemHaMonitorState struct {
	// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
	MonitorVlan *string `pulumi:"monitorVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure heartbeat interval (seconds).
	VlanHbInterval *int `pulumi:"vlanHbInterval"`
	// VLAN lost heartbeat threshold (1 - 60).
	VlanHbLostThreshold *int `pulumi:"vlanHbLostThreshold"`
}

type SystemHaMonitorState struct {
	// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
	MonitorVlan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure heartbeat interval (seconds).
	VlanHbInterval pulumi.IntPtrInput
	// VLAN lost heartbeat threshold (1 - 60).
	VlanHbLostThreshold pulumi.IntPtrInput
}

func (SystemHaMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemHaMonitorState)(nil)).Elem()
}

type systemHaMonitorArgs struct {
	// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
	MonitorVlan *string `pulumi:"monitorVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure heartbeat interval (seconds).
	VlanHbInterval *int `pulumi:"vlanHbInterval"`
	// VLAN lost heartbeat threshold (1 - 60).
	VlanHbLostThreshold *int `pulumi:"vlanHbLostThreshold"`
}

// The set of arguments for constructing a SystemHaMonitor resource.
type SystemHaMonitorArgs struct {
	// Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
	MonitorVlan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure heartbeat interval (seconds).
	VlanHbInterval pulumi.IntPtrInput
	// VLAN lost heartbeat threshold (1 - 60).
	VlanHbLostThreshold pulumi.IntPtrInput
}

func (SystemHaMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemHaMonitorArgs)(nil)).Elem()
}

type SystemHaMonitorInput interface {
	pulumi.Input

	ToSystemHaMonitorOutput() SystemHaMonitorOutput
	ToSystemHaMonitorOutputWithContext(ctx context.Context) SystemHaMonitorOutput
}

func (*SystemHaMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemHaMonitor)(nil)).Elem()
}

func (i *SystemHaMonitor) ToSystemHaMonitorOutput() SystemHaMonitorOutput {
	return i.ToSystemHaMonitorOutputWithContext(context.Background())
}

func (i *SystemHaMonitor) ToSystemHaMonitorOutputWithContext(ctx context.Context) SystemHaMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemHaMonitorOutput)
}

// SystemHaMonitorArrayInput is an input type that accepts SystemHaMonitorArray and SystemHaMonitorArrayOutput values.
// You can construct a concrete instance of `SystemHaMonitorArrayInput` via:
//
//          SystemHaMonitorArray{ SystemHaMonitorArgs{...} }
type SystemHaMonitorArrayInput interface {
	pulumi.Input

	ToSystemHaMonitorArrayOutput() SystemHaMonitorArrayOutput
	ToSystemHaMonitorArrayOutputWithContext(context.Context) SystemHaMonitorArrayOutput
}

type SystemHaMonitorArray []SystemHaMonitorInput

func (SystemHaMonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemHaMonitor)(nil)).Elem()
}

func (i SystemHaMonitorArray) ToSystemHaMonitorArrayOutput() SystemHaMonitorArrayOutput {
	return i.ToSystemHaMonitorArrayOutputWithContext(context.Background())
}

func (i SystemHaMonitorArray) ToSystemHaMonitorArrayOutputWithContext(ctx context.Context) SystemHaMonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemHaMonitorArrayOutput)
}

// SystemHaMonitorMapInput is an input type that accepts SystemHaMonitorMap and SystemHaMonitorMapOutput values.
// You can construct a concrete instance of `SystemHaMonitorMapInput` via:
//
//          SystemHaMonitorMap{ "key": SystemHaMonitorArgs{...} }
type SystemHaMonitorMapInput interface {
	pulumi.Input

	ToSystemHaMonitorMapOutput() SystemHaMonitorMapOutput
	ToSystemHaMonitorMapOutputWithContext(context.Context) SystemHaMonitorMapOutput
}

type SystemHaMonitorMap map[string]SystemHaMonitorInput

func (SystemHaMonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemHaMonitor)(nil)).Elem()
}

func (i SystemHaMonitorMap) ToSystemHaMonitorMapOutput() SystemHaMonitorMapOutput {
	return i.ToSystemHaMonitorMapOutputWithContext(context.Background())
}

func (i SystemHaMonitorMap) ToSystemHaMonitorMapOutputWithContext(ctx context.Context) SystemHaMonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemHaMonitorMapOutput)
}

type SystemHaMonitorOutput struct{ *pulumi.OutputState }

func (SystemHaMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemHaMonitor)(nil)).Elem()
}

func (o SystemHaMonitorOutput) ToSystemHaMonitorOutput() SystemHaMonitorOutput {
	return o
}

func (o SystemHaMonitorOutput) ToSystemHaMonitorOutputWithContext(ctx context.Context) SystemHaMonitorOutput {
	return o
}

type SystemHaMonitorArrayOutput struct{ *pulumi.OutputState }

func (SystemHaMonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemHaMonitor)(nil)).Elem()
}

func (o SystemHaMonitorArrayOutput) ToSystemHaMonitorArrayOutput() SystemHaMonitorArrayOutput {
	return o
}

func (o SystemHaMonitorArrayOutput) ToSystemHaMonitorArrayOutputWithContext(ctx context.Context) SystemHaMonitorArrayOutput {
	return o
}

func (o SystemHaMonitorArrayOutput) Index(i pulumi.IntInput) SystemHaMonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemHaMonitor {
		return vs[0].([]*SystemHaMonitor)[vs[1].(int)]
	}).(SystemHaMonitorOutput)
}

type SystemHaMonitorMapOutput struct{ *pulumi.OutputState }

func (SystemHaMonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemHaMonitor)(nil)).Elem()
}

func (o SystemHaMonitorMapOutput) ToSystemHaMonitorMapOutput() SystemHaMonitorMapOutput {
	return o
}

func (o SystemHaMonitorMapOutput) ToSystemHaMonitorMapOutputWithContext(ctx context.Context) SystemHaMonitorMapOutput {
	return o
}

func (o SystemHaMonitorMapOutput) MapIndex(k pulumi.StringInput) SystemHaMonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemHaMonitor {
		return vs[0].(map[string]*SystemHaMonitor)[vs[1].(string)]
	}).(SystemHaMonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemHaMonitorInput)(nil)).Elem(), &SystemHaMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemHaMonitorArrayInput)(nil)).Elem(), SystemHaMonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemHaMonitorMapInput)(nil)).Elem(), SystemHaMonitorMap{})
	pulumi.RegisterOutputType(SystemHaMonitorOutput{})
	pulumi.RegisterOutputType(SystemHaMonitorArrayOutput{})
	pulumi.RegisterOutputType(SystemHaMonitorMapOutput{})
}

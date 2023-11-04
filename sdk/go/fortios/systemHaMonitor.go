// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type SystemHaMonitor struct {
	pulumi.CustomResourceState

	MonitorVlan         pulumi.StringOutput    `pulumi:"monitorVlan"`
	Vdomparam           pulumi.StringPtrOutput `pulumi:"vdomparam"`
	VlanHbInterval      pulumi.IntOutput       `pulumi:"vlanHbInterval"`
	VlanHbLostThreshold pulumi.IntOutput       `pulumi:"vlanHbLostThreshold"`
}

// NewSystemHaMonitor registers a new resource with the given unique name, arguments, and options.
func NewSystemHaMonitor(ctx *pulumi.Context,
	name string, args *SystemHaMonitorArgs, opts ...pulumi.ResourceOption) (*SystemHaMonitor, error) {
	if args == nil {
		args = &SystemHaMonitorArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
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
	MonitorVlan         *string `pulumi:"monitorVlan"`
	Vdomparam           *string `pulumi:"vdomparam"`
	VlanHbInterval      *int    `pulumi:"vlanHbInterval"`
	VlanHbLostThreshold *int    `pulumi:"vlanHbLostThreshold"`
}

type SystemHaMonitorState struct {
	MonitorVlan         pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	VlanHbInterval      pulumi.IntPtrInput
	VlanHbLostThreshold pulumi.IntPtrInput
}

func (SystemHaMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemHaMonitorState)(nil)).Elem()
}

type systemHaMonitorArgs struct {
	MonitorVlan         *string `pulumi:"monitorVlan"`
	Vdomparam           *string `pulumi:"vdomparam"`
	VlanHbInterval      *int    `pulumi:"vlanHbInterval"`
	VlanHbLostThreshold *int    `pulumi:"vlanHbLostThreshold"`
}

// The set of arguments for constructing a SystemHaMonitor resource.
type SystemHaMonitorArgs struct {
	MonitorVlan         pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	VlanHbInterval      pulumi.IntPtrInput
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

func (i *SystemHaMonitor) ToOutput(ctx context.Context) pulumix.Output[*SystemHaMonitor] {
	return pulumix.Output[*SystemHaMonitor]{
		OutputState: i.ToSystemHaMonitorOutputWithContext(ctx).OutputState,
	}
}

// SystemHaMonitorArrayInput is an input type that accepts SystemHaMonitorArray and SystemHaMonitorArrayOutput values.
// You can construct a concrete instance of `SystemHaMonitorArrayInput` via:
//
//	SystemHaMonitorArray{ SystemHaMonitorArgs{...} }
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

func (i SystemHaMonitorArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemHaMonitor] {
	return pulumix.Output[[]*SystemHaMonitor]{
		OutputState: i.ToSystemHaMonitorArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemHaMonitorMapInput is an input type that accepts SystemHaMonitorMap and SystemHaMonitorMapOutput values.
// You can construct a concrete instance of `SystemHaMonitorMapInput` via:
//
//	SystemHaMonitorMap{ "key": SystemHaMonitorArgs{...} }
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

func (i SystemHaMonitorMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemHaMonitor] {
	return pulumix.Output[map[string]*SystemHaMonitor]{
		OutputState: i.ToSystemHaMonitorMapOutputWithContext(ctx).OutputState,
	}
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

func (o SystemHaMonitorOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemHaMonitor] {
	return pulumix.Output[*SystemHaMonitor]{
		OutputState: o.OutputState,
	}
}

func (o SystemHaMonitorOutput) MonitorVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemHaMonitor) pulumi.StringOutput { return v.MonitorVlan }).(pulumi.StringOutput)
}

func (o SystemHaMonitorOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemHaMonitor) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o SystemHaMonitorOutput) VlanHbInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemHaMonitor) pulumi.IntOutput { return v.VlanHbInterval }).(pulumi.IntOutput)
}

func (o SystemHaMonitorOutput) VlanHbLostThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemHaMonitor) pulumi.IntOutput { return v.VlanHbLostThreshold }).(pulumi.IntOutput)
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

func (o SystemHaMonitorArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemHaMonitor] {
	return pulumix.Output[[]*SystemHaMonitor]{
		OutputState: o.OutputState,
	}
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

func (o SystemHaMonitorMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemHaMonitor] {
	return pulumix.Output[map[string]*SystemHaMonitor]{
		OutputState: o.OutputState,
	}
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

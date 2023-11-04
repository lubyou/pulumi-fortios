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

type SystemManagementTunnel struct {
	pulumi.CustomResourceState

	AllowCollectStatistics pulumi.StringOutput    `pulumi:"allowCollectStatistics"`
	AllowConfigRestore     pulumi.StringOutput    `pulumi:"allowConfigRestore"`
	AllowPushConfiguration pulumi.StringOutput    `pulumi:"allowPushConfiguration"`
	AllowPushFirmware      pulumi.StringOutput    `pulumi:"allowPushFirmware"`
	AuthorizedManagerOnly  pulumi.StringOutput    `pulumi:"authorizedManagerOnly"`
	SerialNumber           pulumi.StringOutput    `pulumi:"serialNumber"`
	Status                 pulumi.StringOutput    `pulumi:"status"`
	Vdomparam              pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemManagementTunnel registers a new resource with the given unique name, arguments, and options.
func NewSystemManagementTunnel(ctx *pulumi.Context,
	name string, args *SystemManagementTunnelArgs, opts ...pulumi.ResourceOption) (*SystemManagementTunnel, error) {
	if args == nil {
		args = &SystemManagementTunnelArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemManagementTunnel
	err := ctx.RegisterResource("fortios:index/systemManagementTunnel:SystemManagementTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemManagementTunnel gets an existing SystemManagementTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemManagementTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemManagementTunnelState, opts ...pulumi.ResourceOption) (*SystemManagementTunnel, error) {
	var resource SystemManagementTunnel
	err := ctx.ReadResource("fortios:index/systemManagementTunnel:SystemManagementTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemManagementTunnel resources.
type systemManagementTunnelState struct {
	AllowCollectStatistics *string `pulumi:"allowCollectStatistics"`
	AllowConfigRestore     *string `pulumi:"allowConfigRestore"`
	AllowPushConfiguration *string `pulumi:"allowPushConfiguration"`
	AllowPushFirmware      *string `pulumi:"allowPushFirmware"`
	AuthorizedManagerOnly  *string `pulumi:"authorizedManagerOnly"`
	SerialNumber           *string `pulumi:"serialNumber"`
	Status                 *string `pulumi:"status"`
	Vdomparam              *string `pulumi:"vdomparam"`
}

type SystemManagementTunnelState struct {
	AllowCollectStatistics pulumi.StringPtrInput
	AllowConfigRestore     pulumi.StringPtrInput
	AllowPushConfiguration pulumi.StringPtrInput
	AllowPushFirmware      pulumi.StringPtrInput
	AuthorizedManagerOnly  pulumi.StringPtrInput
	SerialNumber           pulumi.StringPtrInput
	Status                 pulumi.StringPtrInput
	Vdomparam              pulumi.StringPtrInput
}

func (SystemManagementTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemManagementTunnelState)(nil)).Elem()
}

type systemManagementTunnelArgs struct {
	AllowCollectStatistics *string `pulumi:"allowCollectStatistics"`
	AllowConfigRestore     *string `pulumi:"allowConfigRestore"`
	AllowPushConfiguration *string `pulumi:"allowPushConfiguration"`
	AllowPushFirmware      *string `pulumi:"allowPushFirmware"`
	AuthorizedManagerOnly  *string `pulumi:"authorizedManagerOnly"`
	SerialNumber           *string `pulumi:"serialNumber"`
	Status                 *string `pulumi:"status"`
	Vdomparam              *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemManagementTunnel resource.
type SystemManagementTunnelArgs struct {
	AllowCollectStatistics pulumi.StringPtrInput
	AllowConfigRestore     pulumi.StringPtrInput
	AllowPushConfiguration pulumi.StringPtrInput
	AllowPushFirmware      pulumi.StringPtrInput
	AuthorizedManagerOnly  pulumi.StringPtrInput
	SerialNumber           pulumi.StringPtrInput
	Status                 pulumi.StringPtrInput
	Vdomparam              pulumi.StringPtrInput
}

func (SystemManagementTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemManagementTunnelArgs)(nil)).Elem()
}

type SystemManagementTunnelInput interface {
	pulumi.Input

	ToSystemManagementTunnelOutput() SystemManagementTunnelOutput
	ToSystemManagementTunnelOutputWithContext(ctx context.Context) SystemManagementTunnelOutput
}

func (*SystemManagementTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemManagementTunnel)(nil)).Elem()
}

func (i *SystemManagementTunnel) ToSystemManagementTunnelOutput() SystemManagementTunnelOutput {
	return i.ToSystemManagementTunnelOutputWithContext(context.Background())
}

func (i *SystemManagementTunnel) ToSystemManagementTunnelOutputWithContext(ctx context.Context) SystemManagementTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemManagementTunnelOutput)
}

func (i *SystemManagementTunnel) ToOutput(ctx context.Context) pulumix.Output[*SystemManagementTunnel] {
	return pulumix.Output[*SystemManagementTunnel]{
		OutputState: i.ToSystemManagementTunnelOutputWithContext(ctx).OutputState,
	}
}

// SystemManagementTunnelArrayInput is an input type that accepts SystemManagementTunnelArray and SystemManagementTunnelArrayOutput values.
// You can construct a concrete instance of `SystemManagementTunnelArrayInput` via:
//
//	SystemManagementTunnelArray{ SystemManagementTunnelArgs{...} }
type SystemManagementTunnelArrayInput interface {
	pulumi.Input

	ToSystemManagementTunnelArrayOutput() SystemManagementTunnelArrayOutput
	ToSystemManagementTunnelArrayOutputWithContext(context.Context) SystemManagementTunnelArrayOutput
}

type SystemManagementTunnelArray []SystemManagementTunnelInput

func (SystemManagementTunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemManagementTunnel)(nil)).Elem()
}

func (i SystemManagementTunnelArray) ToSystemManagementTunnelArrayOutput() SystemManagementTunnelArrayOutput {
	return i.ToSystemManagementTunnelArrayOutputWithContext(context.Background())
}

func (i SystemManagementTunnelArray) ToSystemManagementTunnelArrayOutputWithContext(ctx context.Context) SystemManagementTunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemManagementTunnelArrayOutput)
}

func (i SystemManagementTunnelArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemManagementTunnel] {
	return pulumix.Output[[]*SystemManagementTunnel]{
		OutputState: i.ToSystemManagementTunnelArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemManagementTunnelMapInput is an input type that accepts SystemManagementTunnelMap and SystemManagementTunnelMapOutput values.
// You can construct a concrete instance of `SystemManagementTunnelMapInput` via:
//
//	SystemManagementTunnelMap{ "key": SystemManagementTunnelArgs{...} }
type SystemManagementTunnelMapInput interface {
	pulumi.Input

	ToSystemManagementTunnelMapOutput() SystemManagementTunnelMapOutput
	ToSystemManagementTunnelMapOutputWithContext(context.Context) SystemManagementTunnelMapOutput
}

type SystemManagementTunnelMap map[string]SystemManagementTunnelInput

func (SystemManagementTunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemManagementTunnel)(nil)).Elem()
}

func (i SystemManagementTunnelMap) ToSystemManagementTunnelMapOutput() SystemManagementTunnelMapOutput {
	return i.ToSystemManagementTunnelMapOutputWithContext(context.Background())
}

func (i SystemManagementTunnelMap) ToSystemManagementTunnelMapOutputWithContext(ctx context.Context) SystemManagementTunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemManagementTunnelMapOutput)
}

func (i SystemManagementTunnelMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemManagementTunnel] {
	return pulumix.Output[map[string]*SystemManagementTunnel]{
		OutputState: i.ToSystemManagementTunnelMapOutputWithContext(ctx).OutputState,
	}
}

type SystemManagementTunnelOutput struct{ *pulumi.OutputState }

func (SystemManagementTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemManagementTunnel)(nil)).Elem()
}

func (o SystemManagementTunnelOutput) ToSystemManagementTunnelOutput() SystemManagementTunnelOutput {
	return o
}

func (o SystemManagementTunnelOutput) ToSystemManagementTunnelOutputWithContext(ctx context.Context) SystemManagementTunnelOutput {
	return o
}

func (o SystemManagementTunnelOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemManagementTunnel] {
	return pulumix.Output[*SystemManagementTunnel]{
		OutputState: o.OutputState,
	}
}

func (o SystemManagementTunnelOutput) AllowCollectStatistics() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemManagementTunnel) pulumi.StringOutput { return v.AllowCollectStatistics }).(pulumi.StringOutput)
}

func (o SystemManagementTunnelOutput) AllowConfigRestore() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemManagementTunnel) pulumi.StringOutput { return v.AllowConfigRestore }).(pulumi.StringOutput)
}

func (o SystemManagementTunnelOutput) AllowPushConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemManagementTunnel) pulumi.StringOutput { return v.AllowPushConfiguration }).(pulumi.StringOutput)
}

func (o SystemManagementTunnelOutput) AllowPushFirmware() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemManagementTunnel) pulumi.StringOutput { return v.AllowPushFirmware }).(pulumi.StringOutput)
}

func (o SystemManagementTunnelOutput) AuthorizedManagerOnly() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemManagementTunnel) pulumi.StringOutput { return v.AuthorizedManagerOnly }).(pulumi.StringOutput)
}

func (o SystemManagementTunnelOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemManagementTunnel) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o SystemManagementTunnelOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemManagementTunnel) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SystemManagementTunnelOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemManagementTunnel) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemManagementTunnelArrayOutput struct{ *pulumi.OutputState }

func (SystemManagementTunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemManagementTunnel)(nil)).Elem()
}

func (o SystemManagementTunnelArrayOutput) ToSystemManagementTunnelArrayOutput() SystemManagementTunnelArrayOutput {
	return o
}

func (o SystemManagementTunnelArrayOutput) ToSystemManagementTunnelArrayOutputWithContext(ctx context.Context) SystemManagementTunnelArrayOutput {
	return o
}

func (o SystemManagementTunnelArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemManagementTunnel] {
	return pulumix.Output[[]*SystemManagementTunnel]{
		OutputState: o.OutputState,
	}
}

func (o SystemManagementTunnelArrayOutput) Index(i pulumi.IntInput) SystemManagementTunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemManagementTunnel {
		return vs[0].([]*SystemManagementTunnel)[vs[1].(int)]
	}).(SystemManagementTunnelOutput)
}

type SystemManagementTunnelMapOutput struct{ *pulumi.OutputState }

func (SystemManagementTunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemManagementTunnel)(nil)).Elem()
}

func (o SystemManagementTunnelMapOutput) ToSystemManagementTunnelMapOutput() SystemManagementTunnelMapOutput {
	return o
}

func (o SystemManagementTunnelMapOutput) ToSystemManagementTunnelMapOutputWithContext(ctx context.Context) SystemManagementTunnelMapOutput {
	return o
}

func (o SystemManagementTunnelMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemManagementTunnel] {
	return pulumix.Output[map[string]*SystemManagementTunnel]{
		OutputState: o.OutputState,
	}
}

func (o SystemManagementTunnelMapOutput) MapIndex(k pulumi.StringInput) SystemManagementTunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemManagementTunnel {
		return vs[0].(map[string]*SystemManagementTunnel)[vs[1].(string)]
	}).(SystemManagementTunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemManagementTunnelInput)(nil)).Elem(), &SystemManagementTunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemManagementTunnelArrayInput)(nil)).Elem(), SystemManagementTunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemManagementTunnelMapInput)(nil)).Elem(), SystemManagementTunnelMap{})
	pulumi.RegisterOutputType(SystemManagementTunnelOutput{})
	pulumi.RegisterOutputType(SystemManagementTunnelArrayOutput{})
	pulumi.RegisterOutputType(SystemManagementTunnelMapOutput{})
}

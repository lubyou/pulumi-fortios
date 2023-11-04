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

type SystemFtmPush struct {
	pulumi.CustomResourceState

	Proxy      pulumi.StringOutput    `pulumi:"proxy"`
	Server     pulumi.StringOutput    `pulumi:"server"`
	ServerCert pulumi.StringOutput    `pulumi:"serverCert"`
	ServerIp   pulumi.StringOutput    `pulumi:"serverIp"`
	ServerPort pulumi.IntOutput       `pulumi:"serverPort"`
	Status     pulumi.StringOutput    `pulumi:"status"`
	Vdomparam  pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemFtmPush registers a new resource with the given unique name, arguments, and options.
func NewSystemFtmPush(ctx *pulumi.Context,
	name string, args *SystemFtmPushArgs, opts ...pulumi.ResourceOption) (*SystemFtmPush, error) {
	if args == nil {
		args = &SystemFtmPushArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemFtmPush
	err := ctx.RegisterResource("fortios:index/systemFtmPush:SystemFtmPush", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFtmPush gets an existing SystemFtmPush resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFtmPush(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFtmPushState, opts ...pulumi.ResourceOption) (*SystemFtmPush, error) {
	var resource SystemFtmPush
	err := ctx.ReadResource("fortios:index/systemFtmPush:SystemFtmPush", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFtmPush resources.
type systemFtmPushState struct {
	Proxy      *string `pulumi:"proxy"`
	Server     *string `pulumi:"server"`
	ServerCert *string `pulumi:"serverCert"`
	ServerIp   *string `pulumi:"serverIp"`
	ServerPort *int    `pulumi:"serverPort"`
	Status     *string `pulumi:"status"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

type SystemFtmPushState struct {
	Proxy      pulumi.StringPtrInput
	Server     pulumi.StringPtrInput
	ServerCert pulumi.StringPtrInput
	ServerIp   pulumi.StringPtrInput
	ServerPort pulumi.IntPtrInput
	Status     pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (SystemFtmPushState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFtmPushState)(nil)).Elem()
}

type systemFtmPushArgs struct {
	Proxy      *string `pulumi:"proxy"`
	Server     *string `pulumi:"server"`
	ServerCert *string `pulumi:"serverCert"`
	ServerIp   *string `pulumi:"serverIp"`
	ServerPort *int    `pulumi:"serverPort"`
	Status     *string `pulumi:"status"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFtmPush resource.
type SystemFtmPushArgs struct {
	Proxy      pulumi.StringPtrInput
	Server     pulumi.StringPtrInput
	ServerCert pulumi.StringPtrInput
	ServerIp   pulumi.StringPtrInput
	ServerPort pulumi.IntPtrInput
	Status     pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (SystemFtmPushArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFtmPushArgs)(nil)).Elem()
}

type SystemFtmPushInput interface {
	pulumi.Input

	ToSystemFtmPushOutput() SystemFtmPushOutput
	ToSystemFtmPushOutputWithContext(ctx context.Context) SystemFtmPushOutput
}

func (*SystemFtmPush) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFtmPush)(nil)).Elem()
}

func (i *SystemFtmPush) ToSystemFtmPushOutput() SystemFtmPushOutput {
	return i.ToSystemFtmPushOutputWithContext(context.Background())
}

func (i *SystemFtmPush) ToSystemFtmPushOutputWithContext(ctx context.Context) SystemFtmPushOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFtmPushOutput)
}

func (i *SystemFtmPush) ToOutput(ctx context.Context) pulumix.Output[*SystemFtmPush] {
	return pulumix.Output[*SystemFtmPush]{
		OutputState: i.ToSystemFtmPushOutputWithContext(ctx).OutputState,
	}
}

// SystemFtmPushArrayInput is an input type that accepts SystemFtmPushArray and SystemFtmPushArrayOutput values.
// You can construct a concrete instance of `SystemFtmPushArrayInput` via:
//
//	SystemFtmPushArray{ SystemFtmPushArgs{...} }
type SystemFtmPushArrayInput interface {
	pulumi.Input

	ToSystemFtmPushArrayOutput() SystemFtmPushArrayOutput
	ToSystemFtmPushArrayOutputWithContext(context.Context) SystemFtmPushArrayOutput
}

type SystemFtmPushArray []SystemFtmPushInput

func (SystemFtmPushArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFtmPush)(nil)).Elem()
}

func (i SystemFtmPushArray) ToSystemFtmPushArrayOutput() SystemFtmPushArrayOutput {
	return i.ToSystemFtmPushArrayOutputWithContext(context.Background())
}

func (i SystemFtmPushArray) ToSystemFtmPushArrayOutputWithContext(ctx context.Context) SystemFtmPushArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFtmPushArrayOutput)
}

func (i SystemFtmPushArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemFtmPush] {
	return pulumix.Output[[]*SystemFtmPush]{
		OutputState: i.ToSystemFtmPushArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemFtmPushMapInput is an input type that accepts SystemFtmPushMap and SystemFtmPushMapOutput values.
// You can construct a concrete instance of `SystemFtmPushMapInput` via:
//
//	SystemFtmPushMap{ "key": SystemFtmPushArgs{...} }
type SystemFtmPushMapInput interface {
	pulumi.Input

	ToSystemFtmPushMapOutput() SystemFtmPushMapOutput
	ToSystemFtmPushMapOutputWithContext(context.Context) SystemFtmPushMapOutput
}

type SystemFtmPushMap map[string]SystemFtmPushInput

func (SystemFtmPushMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFtmPush)(nil)).Elem()
}

func (i SystemFtmPushMap) ToSystemFtmPushMapOutput() SystemFtmPushMapOutput {
	return i.ToSystemFtmPushMapOutputWithContext(context.Background())
}

func (i SystemFtmPushMap) ToSystemFtmPushMapOutputWithContext(ctx context.Context) SystemFtmPushMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFtmPushMapOutput)
}

func (i SystemFtmPushMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemFtmPush] {
	return pulumix.Output[map[string]*SystemFtmPush]{
		OutputState: i.ToSystemFtmPushMapOutputWithContext(ctx).OutputState,
	}
}

type SystemFtmPushOutput struct{ *pulumi.OutputState }

func (SystemFtmPushOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFtmPush)(nil)).Elem()
}

func (o SystemFtmPushOutput) ToSystemFtmPushOutput() SystemFtmPushOutput {
	return o
}

func (o SystemFtmPushOutput) ToSystemFtmPushOutputWithContext(ctx context.Context) SystemFtmPushOutput {
	return o
}

func (o SystemFtmPushOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemFtmPush] {
	return pulumix.Output[*SystemFtmPush]{
		OutputState: o.OutputState,
	}
}

func (o SystemFtmPushOutput) Proxy() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFtmPush) pulumi.StringOutput { return v.Proxy }).(pulumi.StringOutput)
}

func (o SystemFtmPushOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFtmPush) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o SystemFtmPushOutput) ServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFtmPush) pulumi.StringOutput { return v.ServerCert }).(pulumi.StringOutput)
}

func (o SystemFtmPushOutput) ServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFtmPush) pulumi.StringOutput { return v.ServerIp }).(pulumi.StringOutput)
}

func (o SystemFtmPushOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemFtmPush) pulumi.IntOutput { return v.ServerPort }).(pulumi.IntOutput)
}

func (o SystemFtmPushOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFtmPush) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SystemFtmPushOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFtmPush) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemFtmPushArrayOutput struct{ *pulumi.OutputState }

func (SystemFtmPushArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFtmPush)(nil)).Elem()
}

func (o SystemFtmPushArrayOutput) ToSystemFtmPushArrayOutput() SystemFtmPushArrayOutput {
	return o
}

func (o SystemFtmPushArrayOutput) ToSystemFtmPushArrayOutputWithContext(ctx context.Context) SystemFtmPushArrayOutput {
	return o
}

func (o SystemFtmPushArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemFtmPush] {
	return pulumix.Output[[]*SystemFtmPush]{
		OutputState: o.OutputState,
	}
}

func (o SystemFtmPushArrayOutput) Index(i pulumi.IntInput) SystemFtmPushOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemFtmPush {
		return vs[0].([]*SystemFtmPush)[vs[1].(int)]
	}).(SystemFtmPushOutput)
}

type SystemFtmPushMapOutput struct{ *pulumi.OutputState }

func (SystemFtmPushMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFtmPush)(nil)).Elem()
}

func (o SystemFtmPushMapOutput) ToSystemFtmPushMapOutput() SystemFtmPushMapOutput {
	return o
}

func (o SystemFtmPushMapOutput) ToSystemFtmPushMapOutputWithContext(ctx context.Context) SystemFtmPushMapOutput {
	return o
}

func (o SystemFtmPushMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemFtmPush] {
	return pulumix.Output[map[string]*SystemFtmPush]{
		OutputState: o.OutputState,
	}
}

func (o SystemFtmPushMapOutput) MapIndex(k pulumi.StringInput) SystemFtmPushOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemFtmPush {
		return vs[0].(map[string]*SystemFtmPush)[vs[1].(string)]
	}).(SystemFtmPushOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFtmPushInput)(nil)).Elem(), &SystemFtmPush{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFtmPushArrayInput)(nil)).Elem(), SystemFtmPushArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFtmPushMapInput)(nil)).Elem(), SystemFtmPushMap{})
	pulumi.RegisterOutputType(SystemFtmPushOutput{})
	pulumi.RegisterOutputType(SystemFtmPushArrayOutput{})
	pulumi.RegisterOutputType(SystemFtmPushMapOutput{})
}

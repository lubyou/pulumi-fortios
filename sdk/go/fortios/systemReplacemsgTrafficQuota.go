// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type SystemReplacemsgTrafficQuota struct {
	pulumi.CustomResourceState

	Buffer    pulumi.StringPtrOutput `pulumi:"buffer"`
	Format    pulumi.StringOutput    `pulumi:"format"`
	Header    pulumi.StringOutput    `pulumi:"header"`
	MsgType   pulumi.StringOutput    `pulumi:"msgType"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgTrafficQuota registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgTrafficQuota(ctx *pulumi.Context,
	name string, args *SystemReplacemsgTrafficQuotaArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgTrafficQuota, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgTrafficQuota
	err := ctx.RegisterResource("fortios:index/systemReplacemsgTrafficQuota:SystemReplacemsgTrafficQuota", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgTrafficQuota gets an existing SystemReplacemsgTrafficQuota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgTrafficQuota(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgTrafficQuotaState, opts ...pulumi.ResourceOption) (*SystemReplacemsgTrafficQuota, error) {
	var resource SystemReplacemsgTrafficQuota
	err := ctx.ReadResource("fortios:index/systemReplacemsgTrafficQuota:SystemReplacemsgTrafficQuota", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgTrafficQuota resources.
type systemReplacemsgTrafficQuotaState struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgTrafficQuotaState struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgTrafficQuotaState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgTrafficQuotaState)(nil)).Elem()
}

type systemReplacemsgTrafficQuotaArgs struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   string  `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgTrafficQuota resource.
type SystemReplacemsgTrafficQuotaArgs struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgTrafficQuotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgTrafficQuotaArgs)(nil)).Elem()
}

type SystemReplacemsgTrafficQuotaInput interface {
	pulumi.Input

	ToSystemReplacemsgTrafficQuotaOutput() SystemReplacemsgTrafficQuotaOutput
	ToSystemReplacemsgTrafficQuotaOutputWithContext(ctx context.Context) SystemReplacemsgTrafficQuotaOutput
}

func (*SystemReplacemsgTrafficQuota) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgTrafficQuota)(nil)).Elem()
}

func (i *SystemReplacemsgTrafficQuota) ToSystemReplacemsgTrafficQuotaOutput() SystemReplacemsgTrafficQuotaOutput {
	return i.ToSystemReplacemsgTrafficQuotaOutputWithContext(context.Background())
}

func (i *SystemReplacemsgTrafficQuota) ToSystemReplacemsgTrafficQuotaOutputWithContext(ctx context.Context) SystemReplacemsgTrafficQuotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgTrafficQuotaOutput)
}

func (i *SystemReplacemsgTrafficQuota) ToOutput(ctx context.Context) pulumix.Output[*SystemReplacemsgTrafficQuota] {
	return pulumix.Output[*SystemReplacemsgTrafficQuota]{
		OutputState: i.ToSystemReplacemsgTrafficQuotaOutputWithContext(ctx).OutputState,
	}
}

// SystemReplacemsgTrafficQuotaArrayInput is an input type that accepts SystemReplacemsgTrafficQuotaArray and SystemReplacemsgTrafficQuotaArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgTrafficQuotaArrayInput` via:
//
//	SystemReplacemsgTrafficQuotaArray{ SystemReplacemsgTrafficQuotaArgs{...} }
type SystemReplacemsgTrafficQuotaArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgTrafficQuotaArrayOutput() SystemReplacemsgTrafficQuotaArrayOutput
	ToSystemReplacemsgTrafficQuotaArrayOutputWithContext(context.Context) SystemReplacemsgTrafficQuotaArrayOutput
}

type SystemReplacemsgTrafficQuotaArray []SystemReplacemsgTrafficQuotaInput

func (SystemReplacemsgTrafficQuotaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgTrafficQuota)(nil)).Elem()
}

func (i SystemReplacemsgTrafficQuotaArray) ToSystemReplacemsgTrafficQuotaArrayOutput() SystemReplacemsgTrafficQuotaArrayOutput {
	return i.ToSystemReplacemsgTrafficQuotaArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgTrafficQuotaArray) ToSystemReplacemsgTrafficQuotaArrayOutputWithContext(ctx context.Context) SystemReplacemsgTrafficQuotaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgTrafficQuotaArrayOutput)
}

func (i SystemReplacemsgTrafficQuotaArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemReplacemsgTrafficQuota] {
	return pulumix.Output[[]*SystemReplacemsgTrafficQuota]{
		OutputState: i.ToSystemReplacemsgTrafficQuotaArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemReplacemsgTrafficQuotaMapInput is an input type that accepts SystemReplacemsgTrafficQuotaMap and SystemReplacemsgTrafficQuotaMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgTrafficQuotaMapInput` via:
//
//	SystemReplacemsgTrafficQuotaMap{ "key": SystemReplacemsgTrafficQuotaArgs{...} }
type SystemReplacemsgTrafficQuotaMapInput interface {
	pulumi.Input

	ToSystemReplacemsgTrafficQuotaMapOutput() SystemReplacemsgTrafficQuotaMapOutput
	ToSystemReplacemsgTrafficQuotaMapOutputWithContext(context.Context) SystemReplacemsgTrafficQuotaMapOutput
}

type SystemReplacemsgTrafficQuotaMap map[string]SystemReplacemsgTrafficQuotaInput

func (SystemReplacemsgTrafficQuotaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgTrafficQuota)(nil)).Elem()
}

func (i SystemReplacemsgTrafficQuotaMap) ToSystemReplacemsgTrafficQuotaMapOutput() SystemReplacemsgTrafficQuotaMapOutput {
	return i.ToSystemReplacemsgTrafficQuotaMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgTrafficQuotaMap) ToSystemReplacemsgTrafficQuotaMapOutputWithContext(ctx context.Context) SystemReplacemsgTrafficQuotaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgTrafficQuotaMapOutput)
}

func (i SystemReplacemsgTrafficQuotaMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemReplacemsgTrafficQuota] {
	return pulumix.Output[map[string]*SystemReplacemsgTrafficQuota]{
		OutputState: i.ToSystemReplacemsgTrafficQuotaMapOutputWithContext(ctx).OutputState,
	}
}

type SystemReplacemsgTrafficQuotaOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgTrafficQuotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgTrafficQuota)(nil)).Elem()
}

func (o SystemReplacemsgTrafficQuotaOutput) ToSystemReplacemsgTrafficQuotaOutput() SystemReplacemsgTrafficQuotaOutput {
	return o
}

func (o SystemReplacemsgTrafficQuotaOutput) ToSystemReplacemsgTrafficQuotaOutputWithContext(ctx context.Context) SystemReplacemsgTrafficQuotaOutput {
	return o
}

func (o SystemReplacemsgTrafficQuotaOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemReplacemsgTrafficQuota] {
	return pulumix.Output[*SystemReplacemsgTrafficQuota]{
		OutputState: o.OutputState,
	}
}

func (o SystemReplacemsgTrafficQuotaOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgTrafficQuota) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

func (o SystemReplacemsgTrafficQuotaOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgTrafficQuota) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o SystemReplacemsgTrafficQuotaOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgTrafficQuota) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

func (o SystemReplacemsgTrafficQuotaOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgTrafficQuota) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

func (o SystemReplacemsgTrafficQuotaOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgTrafficQuota) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemReplacemsgTrafficQuotaArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgTrafficQuotaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgTrafficQuota)(nil)).Elem()
}

func (o SystemReplacemsgTrafficQuotaArrayOutput) ToSystemReplacemsgTrafficQuotaArrayOutput() SystemReplacemsgTrafficQuotaArrayOutput {
	return o
}

func (o SystemReplacemsgTrafficQuotaArrayOutput) ToSystemReplacemsgTrafficQuotaArrayOutputWithContext(ctx context.Context) SystemReplacemsgTrafficQuotaArrayOutput {
	return o
}

func (o SystemReplacemsgTrafficQuotaArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemReplacemsgTrafficQuota] {
	return pulumix.Output[[]*SystemReplacemsgTrafficQuota]{
		OutputState: o.OutputState,
	}
}

func (o SystemReplacemsgTrafficQuotaArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgTrafficQuotaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgTrafficQuota {
		return vs[0].([]*SystemReplacemsgTrafficQuota)[vs[1].(int)]
	}).(SystemReplacemsgTrafficQuotaOutput)
}

type SystemReplacemsgTrafficQuotaMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgTrafficQuotaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgTrafficQuota)(nil)).Elem()
}

func (o SystemReplacemsgTrafficQuotaMapOutput) ToSystemReplacemsgTrafficQuotaMapOutput() SystemReplacemsgTrafficQuotaMapOutput {
	return o
}

func (o SystemReplacemsgTrafficQuotaMapOutput) ToSystemReplacemsgTrafficQuotaMapOutputWithContext(ctx context.Context) SystemReplacemsgTrafficQuotaMapOutput {
	return o
}

func (o SystemReplacemsgTrafficQuotaMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemReplacemsgTrafficQuota] {
	return pulumix.Output[map[string]*SystemReplacemsgTrafficQuota]{
		OutputState: o.OutputState,
	}
}

func (o SystemReplacemsgTrafficQuotaMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgTrafficQuotaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgTrafficQuota {
		return vs[0].(map[string]*SystemReplacemsgTrafficQuota)[vs[1].(string)]
	}).(SystemReplacemsgTrafficQuotaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgTrafficQuotaInput)(nil)).Elem(), &SystemReplacemsgTrafficQuota{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgTrafficQuotaArrayInput)(nil)).Elem(), SystemReplacemsgTrafficQuotaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgTrafficQuotaMapInput)(nil)).Elem(), SystemReplacemsgTrafficQuotaMap{})
	pulumi.RegisterOutputType(SystemReplacemsgTrafficQuotaOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgTrafficQuotaArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgTrafficQuotaMapOutput{})
}

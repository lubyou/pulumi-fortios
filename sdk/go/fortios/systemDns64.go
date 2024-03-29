// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemDns64 struct {
	pulumi.CustomResourceState

	AlwaysSynthesizeAaaaRecord pulumi.StringOutput    `pulumi:"alwaysSynthesizeAaaaRecord"`
	Dns64Prefix                pulumi.StringOutput    `pulumi:"dns64Prefix"`
	Status                     pulumi.StringOutput    `pulumi:"status"`
	Vdomparam                  pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemDns64 registers a new resource with the given unique name, arguments, and options.
func NewSystemDns64(ctx *pulumi.Context,
	name string, args *SystemDns64Args, opts ...pulumi.ResourceOption) (*SystemDns64, error) {
	if args == nil {
		args = &SystemDns64Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemDns64
	err := ctx.RegisterResource("fortios:index/systemDns64:SystemDns64", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemDns64 gets an existing SystemDns64 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemDns64(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemDns64State, opts ...pulumi.ResourceOption) (*SystemDns64, error) {
	var resource SystemDns64
	err := ctx.ReadResource("fortios:index/systemDns64:SystemDns64", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemDns64 resources.
type systemDns64State struct {
	AlwaysSynthesizeAaaaRecord *string `pulumi:"alwaysSynthesizeAaaaRecord"`
	Dns64Prefix                *string `pulumi:"dns64Prefix"`
	Status                     *string `pulumi:"status"`
	Vdomparam                  *string `pulumi:"vdomparam"`
}

type SystemDns64State struct {
	AlwaysSynthesizeAaaaRecord pulumi.StringPtrInput
	Dns64Prefix                pulumi.StringPtrInput
	Status                     pulumi.StringPtrInput
	Vdomparam                  pulumi.StringPtrInput
}

func (SystemDns64State) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDns64State)(nil)).Elem()
}

type systemDns64Args struct {
	AlwaysSynthesizeAaaaRecord *string `pulumi:"alwaysSynthesizeAaaaRecord"`
	Dns64Prefix                *string `pulumi:"dns64Prefix"`
	Status                     *string `pulumi:"status"`
	Vdomparam                  *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemDns64 resource.
type SystemDns64Args struct {
	AlwaysSynthesizeAaaaRecord pulumi.StringPtrInput
	Dns64Prefix                pulumi.StringPtrInput
	Status                     pulumi.StringPtrInput
	Vdomparam                  pulumi.StringPtrInput
}

func (SystemDns64Args) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDns64Args)(nil)).Elem()
}

type SystemDns64Input interface {
	pulumi.Input

	ToSystemDns64Output() SystemDns64Output
	ToSystemDns64OutputWithContext(ctx context.Context) SystemDns64Output
}

func (*SystemDns64) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDns64)(nil)).Elem()
}

func (i *SystemDns64) ToSystemDns64Output() SystemDns64Output {
	return i.ToSystemDns64OutputWithContext(context.Background())
}

func (i *SystemDns64) ToSystemDns64OutputWithContext(ctx context.Context) SystemDns64Output {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDns64Output)
}

// SystemDns64ArrayInput is an input type that accepts SystemDns64Array and SystemDns64ArrayOutput values.
// You can construct a concrete instance of `SystemDns64ArrayInput` via:
//
//	SystemDns64Array{ SystemDns64Args{...} }
type SystemDns64ArrayInput interface {
	pulumi.Input

	ToSystemDns64ArrayOutput() SystemDns64ArrayOutput
	ToSystemDns64ArrayOutputWithContext(context.Context) SystemDns64ArrayOutput
}

type SystemDns64Array []SystemDns64Input

func (SystemDns64Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDns64)(nil)).Elem()
}

func (i SystemDns64Array) ToSystemDns64ArrayOutput() SystemDns64ArrayOutput {
	return i.ToSystemDns64ArrayOutputWithContext(context.Background())
}

func (i SystemDns64Array) ToSystemDns64ArrayOutputWithContext(ctx context.Context) SystemDns64ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDns64ArrayOutput)
}

// SystemDns64MapInput is an input type that accepts SystemDns64Map and SystemDns64MapOutput values.
// You can construct a concrete instance of `SystemDns64MapInput` via:
//
//	SystemDns64Map{ "key": SystemDns64Args{...} }
type SystemDns64MapInput interface {
	pulumi.Input

	ToSystemDns64MapOutput() SystemDns64MapOutput
	ToSystemDns64MapOutputWithContext(context.Context) SystemDns64MapOutput
}

type SystemDns64Map map[string]SystemDns64Input

func (SystemDns64Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDns64)(nil)).Elem()
}

func (i SystemDns64Map) ToSystemDns64MapOutput() SystemDns64MapOutput {
	return i.ToSystemDns64MapOutputWithContext(context.Background())
}

func (i SystemDns64Map) ToSystemDns64MapOutputWithContext(ctx context.Context) SystemDns64MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDns64MapOutput)
}

type SystemDns64Output struct{ *pulumi.OutputState }

func (SystemDns64Output) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDns64)(nil)).Elem()
}

func (o SystemDns64Output) ToSystemDns64Output() SystemDns64Output {
	return o
}

func (o SystemDns64Output) ToSystemDns64OutputWithContext(ctx context.Context) SystemDns64Output {
	return o
}

func (o SystemDns64Output) AlwaysSynthesizeAaaaRecord() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDns64) pulumi.StringOutput { return v.AlwaysSynthesizeAaaaRecord }).(pulumi.StringOutput)
}

func (o SystemDns64Output) Dns64Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDns64) pulumi.StringOutput { return v.Dns64Prefix }).(pulumi.StringOutput)
}

func (o SystemDns64Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDns64) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SystemDns64Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDns64) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemDns64ArrayOutput struct{ *pulumi.OutputState }

func (SystemDns64ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDns64)(nil)).Elem()
}

func (o SystemDns64ArrayOutput) ToSystemDns64ArrayOutput() SystemDns64ArrayOutput {
	return o
}

func (o SystemDns64ArrayOutput) ToSystemDns64ArrayOutputWithContext(ctx context.Context) SystemDns64ArrayOutput {
	return o
}

func (o SystemDns64ArrayOutput) Index(i pulumi.IntInput) SystemDns64Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemDns64 {
		return vs[0].([]*SystemDns64)[vs[1].(int)]
	}).(SystemDns64Output)
}

type SystemDns64MapOutput struct{ *pulumi.OutputState }

func (SystemDns64MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDns64)(nil)).Elem()
}

func (o SystemDns64MapOutput) ToSystemDns64MapOutput() SystemDns64MapOutput {
	return o
}

func (o SystemDns64MapOutput) ToSystemDns64MapOutputWithContext(ctx context.Context) SystemDns64MapOutput {
	return o
}

func (o SystemDns64MapOutput) MapIndex(k pulumi.StringInput) SystemDns64Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemDns64 {
		return vs[0].(map[string]*SystemDns64)[vs[1].(string)]
	}).(SystemDns64Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDns64Input)(nil)).Elem(), &SystemDns64{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDns64ArrayInput)(nil)).Elem(), SystemDns64Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDns64MapInput)(nil)).Elem(), SystemDns64Map{})
	pulumi.RegisterOutputType(SystemDns64Output{})
	pulumi.RegisterOutputType(SystemDns64ArrayOutput{})
	pulumi.RegisterOutputType(SystemDns64MapOutput{})
}

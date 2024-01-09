// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemVdomSflow struct {
	pulumi.CustomResourceState

	CollectorIp           pulumi.StringOutput    `pulumi:"collectorIp"`
	CollectorPort         pulumi.IntOutput       `pulumi:"collectorPort"`
	Interface             pulumi.StringOutput    `pulumi:"interface"`
	InterfaceSelectMethod pulumi.StringOutput    `pulumi:"interfaceSelectMethod"`
	SourceIp              pulumi.StringOutput    `pulumi:"sourceIp"`
	VdomSflow             pulumi.StringOutput    `pulumi:"vdomSflow"`
	Vdomparam             pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemVdomSflow registers a new resource with the given unique name, arguments, and options.
func NewSystemVdomSflow(ctx *pulumi.Context,
	name string, args *SystemVdomSflowArgs, opts ...pulumi.ResourceOption) (*SystemVdomSflow, error) {
	if args == nil {
		args = &SystemVdomSflowArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemVdomSflow
	err := ctx.RegisterResource("fortios:index/systemVdomSflow:SystemVdomSflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVdomSflow gets an existing SystemVdomSflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVdomSflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVdomSflowState, opts ...pulumi.ResourceOption) (*SystemVdomSflow, error) {
	var resource SystemVdomSflow
	err := ctx.ReadResource("fortios:index/systemVdomSflow:SystemVdomSflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVdomSflow resources.
type systemVdomSflowState struct {
	CollectorIp           *string `pulumi:"collectorIp"`
	CollectorPort         *int    `pulumi:"collectorPort"`
	Interface             *string `pulumi:"interface"`
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	SourceIp              *string `pulumi:"sourceIp"`
	VdomSflow             *string `pulumi:"vdomSflow"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

type SystemVdomSflowState struct {
	CollectorIp           pulumi.StringPtrInput
	CollectorPort         pulumi.IntPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	VdomSflow             pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (SystemVdomSflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomSflowState)(nil)).Elem()
}

type systemVdomSflowArgs struct {
	CollectorIp           *string `pulumi:"collectorIp"`
	CollectorPort         *int    `pulumi:"collectorPort"`
	Interface             *string `pulumi:"interface"`
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	SourceIp              *string `pulumi:"sourceIp"`
	VdomSflow             *string `pulumi:"vdomSflow"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemVdomSflow resource.
type SystemVdomSflowArgs struct {
	CollectorIp           pulumi.StringPtrInput
	CollectorPort         pulumi.IntPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	VdomSflow             pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (SystemVdomSflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomSflowArgs)(nil)).Elem()
}

type SystemVdomSflowInput interface {
	pulumi.Input

	ToSystemVdomSflowOutput() SystemVdomSflowOutput
	ToSystemVdomSflowOutputWithContext(ctx context.Context) SystemVdomSflowOutput
}

func (*SystemVdomSflow) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomSflow)(nil)).Elem()
}

func (i *SystemVdomSflow) ToSystemVdomSflowOutput() SystemVdomSflowOutput {
	return i.ToSystemVdomSflowOutputWithContext(context.Background())
}

func (i *SystemVdomSflow) ToSystemVdomSflowOutputWithContext(ctx context.Context) SystemVdomSflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomSflowOutput)
}

// SystemVdomSflowArrayInput is an input type that accepts SystemVdomSflowArray and SystemVdomSflowArrayOutput values.
// You can construct a concrete instance of `SystemVdomSflowArrayInput` via:
//
//	SystemVdomSflowArray{ SystemVdomSflowArgs{...} }
type SystemVdomSflowArrayInput interface {
	pulumi.Input

	ToSystemVdomSflowArrayOutput() SystemVdomSflowArrayOutput
	ToSystemVdomSflowArrayOutputWithContext(context.Context) SystemVdomSflowArrayOutput
}

type SystemVdomSflowArray []SystemVdomSflowInput

func (SystemVdomSflowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomSflow)(nil)).Elem()
}

func (i SystemVdomSflowArray) ToSystemVdomSflowArrayOutput() SystemVdomSflowArrayOutput {
	return i.ToSystemVdomSflowArrayOutputWithContext(context.Background())
}

func (i SystemVdomSflowArray) ToSystemVdomSflowArrayOutputWithContext(ctx context.Context) SystemVdomSflowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomSflowArrayOutput)
}

// SystemVdomSflowMapInput is an input type that accepts SystemVdomSflowMap and SystemVdomSflowMapOutput values.
// You can construct a concrete instance of `SystemVdomSflowMapInput` via:
//
//	SystemVdomSflowMap{ "key": SystemVdomSflowArgs{...} }
type SystemVdomSflowMapInput interface {
	pulumi.Input

	ToSystemVdomSflowMapOutput() SystemVdomSflowMapOutput
	ToSystemVdomSflowMapOutputWithContext(context.Context) SystemVdomSflowMapOutput
}

type SystemVdomSflowMap map[string]SystemVdomSflowInput

func (SystemVdomSflowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomSflow)(nil)).Elem()
}

func (i SystemVdomSflowMap) ToSystemVdomSflowMapOutput() SystemVdomSflowMapOutput {
	return i.ToSystemVdomSflowMapOutputWithContext(context.Background())
}

func (i SystemVdomSflowMap) ToSystemVdomSflowMapOutputWithContext(ctx context.Context) SystemVdomSflowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomSflowMapOutput)
}

type SystemVdomSflowOutput struct{ *pulumi.OutputState }

func (SystemVdomSflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomSflow)(nil)).Elem()
}

func (o SystemVdomSflowOutput) ToSystemVdomSflowOutput() SystemVdomSflowOutput {
	return o
}

func (o SystemVdomSflowOutput) ToSystemVdomSflowOutputWithContext(ctx context.Context) SystemVdomSflowOutput {
	return o
}

func (o SystemVdomSflowOutput) CollectorIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomSflow) pulumi.StringOutput { return v.CollectorIp }).(pulumi.StringOutput)
}

func (o SystemVdomSflowOutput) CollectorPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemVdomSflow) pulumi.IntOutput { return v.CollectorPort }).(pulumi.IntOutput)
}

func (o SystemVdomSflowOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomSflow) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o SystemVdomSflowOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomSflow) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o SystemVdomSflowOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomSflow) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o SystemVdomSflowOutput) VdomSflow() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomSflow) pulumi.StringOutput { return v.VdomSflow }).(pulumi.StringOutput)
}

func (o SystemVdomSflowOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemVdomSflow) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemVdomSflowArrayOutput struct{ *pulumi.OutputState }

func (SystemVdomSflowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomSflow)(nil)).Elem()
}

func (o SystemVdomSflowArrayOutput) ToSystemVdomSflowArrayOutput() SystemVdomSflowArrayOutput {
	return o
}

func (o SystemVdomSflowArrayOutput) ToSystemVdomSflowArrayOutputWithContext(ctx context.Context) SystemVdomSflowArrayOutput {
	return o
}

func (o SystemVdomSflowArrayOutput) Index(i pulumi.IntInput) SystemVdomSflowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemVdomSflow {
		return vs[0].([]*SystemVdomSflow)[vs[1].(int)]
	}).(SystemVdomSflowOutput)
}

type SystemVdomSflowMapOutput struct{ *pulumi.OutputState }

func (SystemVdomSflowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomSflow)(nil)).Elem()
}

func (o SystemVdomSflowMapOutput) ToSystemVdomSflowMapOutput() SystemVdomSflowMapOutput {
	return o
}

func (o SystemVdomSflowMapOutput) ToSystemVdomSflowMapOutputWithContext(ctx context.Context) SystemVdomSflowMapOutput {
	return o
}

func (o SystemVdomSflowMapOutput) MapIndex(k pulumi.StringInput) SystemVdomSflowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemVdomSflow {
		return vs[0].(map[string]*SystemVdomSflow)[vs[1].(string)]
	}).(SystemVdomSflowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomSflowInput)(nil)).Elem(), &SystemVdomSflow{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomSflowArrayInput)(nil)).Elem(), SystemVdomSflowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomSflowMapInput)(nil)).Elem(), SystemVdomSflowMap{})
	pulumi.RegisterOutputType(SystemVdomSflowOutput{})
	pulumi.RegisterOutputType(SystemVdomSflowArrayOutput{})
	pulumi.RegisterOutputType(SystemVdomSflowMapOutput{})
}

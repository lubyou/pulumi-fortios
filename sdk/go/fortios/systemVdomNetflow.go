// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemVdomNetflow struct {
	pulumi.CustomResourceState

	CollectorIp           pulumi.StringOutput    `pulumi:"collectorIp"`
	CollectorPort         pulumi.IntOutput       `pulumi:"collectorPort"`
	Interface             pulumi.StringOutput    `pulumi:"interface"`
	InterfaceSelectMethod pulumi.StringOutput    `pulumi:"interfaceSelectMethod"`
	SourceIp              pulumi.StringOutput    `pulumi:"sourceIp"`
	VdomNetflow           pulumi.StringOutput    `pulumi:"vdomNetflow"`
	Vdomparam             pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemVdomNetflow registers a new resource with the given unique name, arguments, and options.
func NewSystemVdomNetflow(ctx *pulumi.Context,
	name string, args *SystemVdomNetflowArgs, opts ...pulumi.ResourceOption) (*SystemVdomNetflow, error) {
	if args == nil {
		args = &SystemVdomNetflowArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemVdomNetflow
	err := ctx.RegisterResource("fortios:index/systemVdomNetflow:SystemVdomNetflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVdomNetflow gets an existing SystemVdomNetflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVdomNetflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVdomNetflowState, opts ...pulumi.ResourceOption) (*SystemVdomNetflow, error) {
	var resource SystemVdomNetflow
	err := ctx.ReadResource("fortios:index/systemVdomNetflow:SystemVdomNetflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVdomNetflow resources.
type systemVdomNetflowState struct {
	CollectorIp           *string `pulumi:"collectorIp"`
	CollectorPort         *int    `pulumi:"collectorPort"`
	Interface             *string `pulumi:"interface"`
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	SourceIp              *string `pulumi:"sourceIp"`
	VdomNetflow           *string `pulumi:"vdomNetflow"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

type SystemVdomNetflowState struct {
	CollectorIp           pulumi.StringPtrInput
	CollectorPort         pulumi.IntPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	VdomNetflow           pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (SystemVdomNetflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomNetflowState)(nil)).Elem()
}

type systemVdomNetflowArgs struct {
	CollectorIp           *string `pulumi:"collectorIp"`
	CollectorPort         *int    `pulumi:"collectorPort"`
	Interface             *string `pulumi:"interface"`
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	SourceIp              *string `pulumi:"sourceIp"`
	VdomNetflow           *string `pulumi:"vdomNetflow"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemVdomNetflow resource.
type SystemVdomNetflowArgs struct {
	CollectorIp           pulumi.StringPtrInput
	CollectorPort         pulumi.IntPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	VdomNetflow           pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (SystemVdomNetflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomNetflowArgs)(nil)).Elem()
}

type SystemVdomNetflowInput interface {
	pulumi.Input

	ToSystemVdomNetflowOutput() SystemVdomNetflowOutput
	ToSystemVdomNetflowOutputWithContext(ctx context.Context) SystemVdomNetflowOutput
}

func (*SystemVdomNetflow) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomNetflow)(nil)).Elem()
}

func (i *SystemVdomNetflow) ToSystemVdomNetflowOutput() SystemVdomNetflowOutput {
	return i.ToSystemVdomNetflowOutputWithContext(context.Background())
}

func (i *SystemVdomNetflow) ToSystemVdomNetflowOutputWithContext(ctx context.Context) SystemVdomNetflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomNetflowOutput)
}

// SystemVdomNetflowArrayInput is an input type that accepts SystemVdomNetflowArray and SystemVdomNetflowArrayOutput values.
// You can construct a concrete instance of `SystemVdomNetflowArrayInput` via:
//
//	SystemVdomNetflowArray{ SystemVdomNetflowArgs{...} }
type SystemVdomNetflowArrayInput interface {
	pulumi.Input

	ToSystemVdomNetflowArrayOutput() SystemVdomNetflowArrayOutput
	ToSystemVdomNetflowArrayOutputWithContext(context.Context) SystemVdomNetflowArrayOutput
}

type SystemVdomNetflowArray []SystemVdomNetflowInput

func (SystemVdomNetflowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomNetflow)(nil)).Elem()
}

func (i SystemVdomNetflowArray) ToSystemVdomNetflowArrayOutput() SystemVdomNetflowArrayOutput {
	return i.ToSystemVdomNetflowArrayOutputWithContext(context.Background())
}

func (i SystemVdomNetflowArray) ToSystemVdomNetflowArrayOutputWithContext(ctx context.Context) SystemVdomNetflowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomNetflowArrayOutput)
}

// SystemVdomNetflowMapInput is an input type that accepts SystemVdomNetflowMap and SystemVdomNetflowMapOutput values.
// You can construct a concrete instance of `SystemVdomNetflowMapInput` via:
//
//	SystemVdomNetflowMap{ "key": SystemVdomNetflowArgs{...} }
type SystemVdomNetflowMapInput interface {
	pulumi.Input

	ToSystemVdomNetflowMapOutput() SystemVdomNetflowMapOutput
	ToSystemVdomNetflowMapOutputWithContext(context.Context) SystemVdomNetflowMapOutput
}

type SystemVdomNetflowMap map[string]SystemVdomNetflowInput

func (SystemVdomNetflowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomNetflow)(nil)).Elem()
}

func (i SystemVdomNetflowMap) ToSystemVdomNetflowMapOutput() SystemVdomNetflowMapOutput {
	return i.ToSystemVdomNetflowMapOutputWithContext(context.Background())
}

func (i SystemVdomNetflowMap) ToSystemVdomNetflowMapOutputWithContext(ctx context.Context) SystemVdomNetflowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomNetflowMapOutput)
}

type SystemVdomNetflowOutput struct{ *pulumi.OutputState }

func (SystemVdomNetflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomNetflow)(nil)).Elem()
}

func (o SystemVdomNetflowOutput) ToSystemVdomNetflowOutput() SystemVdomNetflowOutput {
	return o
}

func (o SystemVdomNetflowOutput) ToSystemVdomNetflowOutputWithContext(ctx context.Context) SystemVdomNetflowOutput {
	return o
}

func (o SystemVdomNetflowOutput) CollectorIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomNetflow) pulumi.StringOutput { return v.CollectorIp }).(pulumi.StringOutput)
}

func (o SystemVdomNetflowOutput) CollectorPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemVdomNetflow) pulumi.IntOutput { return v.CollectorPort }).(pulumi.IntOutput)
}

func (o SystemVdomNetflowOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomNetflow) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o SystemVdomNetflowOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomNetflow) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o SystemVdomNetflowOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomNetflow) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o SystemVdomNetflowOutput) VdomNetflow() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomNetflow) pulumi.StringOutput { return v.VdomNetflow }).(pulumi.StringOutput)
}

func (o SystemVdomNetflowOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemVdomNetflow) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemVdomNetflowArrayOutput struct{ *pulumi.OutputState }

func (SystemVdomNetflowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomNetflow)(nil)).Elem()
}

func (o SystemVdomNetflowArrayOutput) ToSystemVdomNetflowArrayOutput() SystemVdomNetflowArrayOutput {
	return o
}

func (o SystemVdomNetflowArrayOutput) ToSystemVdomNetflowArrayOutputWithContext(ctx context.Context) SystemVdomNetflowArrayOutput {
	return o
}

func (o SystemVdomNetflowArrayOutput) Index(i pulumi.IntInput) SystemVdomNetflowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemVdomNetflow {
		return vs[0].([]*SystemVdomNetflow)[vs[1].(int)]
	}).(SystemVdomNetflowOutput)
}

type SystemVdomNetflowMapOutput struct{ *pulumi.OutputState }

func (SystemVdomNetflowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomNetflow)(nil)).Elem()
}

func (o SystemVdomNetflowMapOutput) ToSystemVdomNetflowMapOutput() SystemVdomNetflowMapOutput {
	return o
}

func (o SystemVdomNetflowMapOutput) ToSystemVdomNetflowMapOutputWithContext(ctx context.Context) SystemVdomNetflowMapOutput {
	return o
}

func (o SystemVdomNetflowMapOutput) MapIndex(k pulumi.StringInput) SystemVdomNetflowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemVdomNetflow {
		return vs[0].(map[string]*SystemVdomNetflow)[vs[1].(string)]
	}).(SystemVdomNetflowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomNetflowInput)(nil)).Elem(), &SystemVdomNetflow{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomNetflowArrayInput)(nil)).Elem(), SystemVdomNetflowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomNetflowMapInput)(nil)).Elem(), SystemVdomNetflowMap{})
	pulumi.RegisterOutputType(SystemVdomNetflowOutput{})
	pulumi.RegisterOutputType(SystemVdomNetflowArrayOutput{})
	pulumi.RegisterOutputType(SystemVdomNetflowMapOutput{})
}

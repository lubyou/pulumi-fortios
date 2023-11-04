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

type SwitchControllerAutoConfigPolicy struct {
	pulumi.CustomResourceState

	IgmpFloodReport    pulumi.StringOutput    `pulumi:"igmpFloodReport"`
	IgmpFloodTraffic   pulumi.StringOutput    `pulumi:"igmpFloodTraffic"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	PoeStatus          pulumi.StringOutput    `pulumi:"poeStatus"`
	QosPolicy          pulumi.StringOutput    `pulumi:"qosPolicy"`
	StormControlPolicy pulumi.StringOutput    `pulumi:"stormControlPolicy"`
	Vdomparam          pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerAutoConfigPolicy registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerAutoConfigPolicy(ctx *pulumi.Context,
	name string, args *SwitchControllerAutoConfigPolicyArgs, opts ...pulumi.ResourceOption) (*SwitchControllerAutoConfigPolicy, error) {
	if args == nil {
		args = &SwitchControllerAutoConfigPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwitchControllerAutoConfigPolicy
	err := ctx.RegisterResource("fortios:index/switchControllerAutoConfigPolicy:SwitchControllerAutoConfigPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerAutoConfigPolicy gets an existing SwitchControllerAutoConfigPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerAutoConfigPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerAutoConfigPolicyState, opts ...pulumi.ResourceOption) (*SwitchControllerAutoConfigPolicy, error) {
	var resource SwitchControllerAutoConfigPolicy
	err := ctx.ReadResource("fortios:index/switchControllerAutoConfigPolicy:SwitchControllerAutoConfigPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerAutoConfigPolicy resources.
type switchControllerAutoConfigPolicyState struct {
	IgmpFloodReport    *string `pulumi:"igmpFloodReport"`
	IgmpFloodTraffic   *string `pulumi:"igmpFloodTraffic"`
	Name               *string `pulumi:"name"`
	PoeStatus          *string `pulumi:"poeStatus"`
	QosPolicy          *string `pulumi:"qosPolicy"`
	StormControlPolicy *string `pulumi:"stormControlPolicy"`
	Vdomparam          *string `pulumi:"vdomparam"`
}

type SwitchControllerAutoConfigPolicyState struct {
	IgmpFloodReport    pulumi.StringPtrInput
	IgmpFloodTraffic   pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	PoeStatus          pulumi.StringPtrInput
	QosPolicy          pulumi.StringPtrInput
	StormControlPolicy pulumi.StringPtrInput
	Vdomparam          pulumi.StringPtrInput
}

func (SwitchControllerAutoConfigPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerAutoConfigPolicyState)(nil)).Elem()
}

type switchControllerAutoConfigPolicyArgs struct {
	IgmpFloodReport    *string `pulumi:"igmpFloodReport"`
	IgmpFloodTraffic   *string `pulumi:"igmpFloodTraffic"`
	Name               *string `pulumi:"name"`
	PoeStatus          *string `pulumi:"poeStatus"`
	QosPolicy          *string `pulumi:"qosPolicy"`
	StormControlPolicy *string `pulumi:"stormControlPolicy"`
	Vdomparam          *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerAutoConfigPolicy resource.
type SwitchControllerAutoConfigPolicyArgs struct {
	IgmpFloodReport    pulumi.StringPtrInput
	IgmpFloodTraffic   pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	PoeStatus          pulumi.StringPtrInput
	QosPolicy          pulumi.StringPtrInput
	StormControlPolicy pulumi.StringPtrInput
	Vdomparam          pulumi.StringPtrInput
}

func (SwitchControllerAutoConfigPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerAutoConfigPolicyArgs)(nil)).Elem()
}

type SwitchControllerAutoConfigPolicyInput interface {
	pulumi.Input

	ToSwitchControllerAutoConfigPolicyOutput() SwitchControllerAutoConfigPolicyOutput
	ToSwitchControllerAutoConfigPolicyOutputWithContext(ctx context.Context) SwitchControllerAutoConfigPolicyOutput
}

func (*SwitchControllerAutoConfigPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerAutoConfigPolicy)(nil)).Elem()
}

func (i *SwitchControllerAutoConfigPolicy) ToSwitchControllerAutoConfigPolicyOutput() SwitchControllerAutoConfigPolicyOutput {
	return i.ToSwitchControllerAutoConfigPolicyOutputWithContext(context.Background())
}

func (i *SwitchControllerAutoConfigPolicy) ToSwitchControllerAutoConfigPolicyOutputWithContext(ctx context.Context) SwitchControllerAutoConfigPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerAutoConfigPolicyOutput)
}

func (i *SwitchControllerAutoConfigPolicy) ToOutput(ctx context.Context) pulumix.Output[*SwitchControllerAutoConfigPolicy] {
	return pulumix.Output[*SwitchControllerAutoConfigPolicy]{
		OutputState: i.ToSwitchControllerAutoConfigPolicyOutputWithContext(ctx).OutputState,
	}
}

// SwitchControllerAutoConfigPolicyArrayInput is an input type that accepts SwitchControllerAutoConfigPolicyArray and SwitchControllerAutoConfigPolicyArrayOutput values.
// You can construct a concrete instance of `SwitchControllerAutoConfigPolicyArrayInput` via:
//
//	SwitchControllerAutoConfigPolicyArray{ SwitchControllerAutoConfigPolicyArgs{...} }
type SwitchControllerAutoConfigPolicyArrayInput interface {
	pulumi.Input

	ToSwitchControllerAutoConfigPolicyArrayOutput() SwitchControllerAutoConfigPolicyArrayOutput
	ToSwitchControllerAutoConfigPolicyArrayOutputWithContext(context.Context) SwitchControllerAutoConfigPolicyArrayOutput
}

type SwitchControllerAutoConfigPolicyArray []SwitchControllerAutoConfigPolicyInput

func (SwitchControllerAutoConfigPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerAutoConfigPolicy)(nil)).Elem()
}

func (i SwitchControllerAutoConfigPolicyArray) ToSwitchControllerAutoConfigPolicyArrayOutput() SwitchControllerAutoConfigPolicyArrayOutput {
	return i.ToSwitchControllerAutoConfigPolicyArrayOutputWithContext(context.Background())
}

func (i SwitchControllerAutoConfigPolicyArray) ToSwitchControllerAutoConfigPolicyArrayOutputWithContext(ctx context.Context) SwitchControllerAutoConfigPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerAutoConfigPolicyArrayOutput)
}

func (i SwitchControllerAutoConfigPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*SwitchControllerAutoConfigPolicy] {
	return pulumix.Output[[]*SwitchControllerAutoConfigPolicy]{
		OutputState: i.ToSwitchControllerAutoConfigPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// SwitchControllerAutoConfigPolicyMapInput is an input type that accepts SwitchControllerAutoConfigPolicyMap and SwitchControllerAutoConfigPolicyMapOutput values.
// You can construct a concrete instance of `SwitchControllerAutoConfigPolicyMapInput` via:
//
//	SwitchControllerAutoConfigPolicyMap{ "key": SwitchControllerAutoConfigPolicyArgs{...} }
type SwitchControllerAutoConfigPolicyMapInput interface {
	pulumi.Input

	ToSwitchControllerAutoConfigPolicyMapOutput() SwitchControllerAutoConfigPolicyMapOutput
	ToSwitchControllerAutoConfigPolicyMapOutputWithContext(context.Context) SwitchControllerAutoConfigPolicyMapOutput
}

type SwitchControllerAutoConfigPolicyMap map[string]SwitchControllerAutoConfigPolicyInput

func (SwitchControllerAutoConfigPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerAutoConfigPolicy)(nil)).Elem()
}

func (i SwitchControllerAutoConfigPolicyMap) ToSwitchControllerAutoConfigPolicyMapOutput() SwitchControllerAutoConfigPolicyMapOutput {
	return i.ToSwitchControllerAutoConfigPolicyMapOutputWithContext(context.Background())
}

func (i SwitchControllerAutoConfigPolicyMap) ToSwitchControllerAutoConfigPolicyMapOutputWithContext(ctx context.Context) SwitchControllerAutoConfigPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerAutoConfigPolicyMapOutput)
}

func (i SwitchControllerAutoConfigPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SwitchControllerAutoConfigPolicy] {
	return pulumix.Output[map[string]*SwitchControllerAutoConfigPolicy]{
		OutputState: i.ToSwitchControllerAutoConfigPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type SwitchControllerAutoConfigPolicyOutput struct{ *pulumi.OutputState }

func (SwitchControllerAutoConfigPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerAutoConfigPolicy)(nil)).Elem()
}

func (o SwitchControllerAutoConfigPolicyOutput) ToSwitchControllerAutoConfigPolicyOutput() SwitchControllerAutoConfigPolicyOutput {
	return o
}

func (o SwitchControllerAutoConfigPolicyOutput) ToSwitchControllerAutoConfigPolicyOutputWithContext(ctx context.Context) SwitchControllerAutoConfigPolicyOutput {
	return o
}

func (o SwitchControllerAutoConfigPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*SwitchControllerAutoConfigPolicy] {
	return pulumix.Output[*SwitchControllerAutoConfigPolicy]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerAutoConfigPolicyOutput) IgmpFloodReport() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigPolicy) pulumi.StringOutput { return v.IgmpFloodReport }).(pulumi.StringOutput)
}

func (o SwitchControllerAutoConfigPolicyOutput) IgmpFloodTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigPolicy) pulumi.StringOutput { return v.IgmpFloodTraffic }).(pulumi.StringOutput)
}

func (o SwitchControllerAutoConfigPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwitchControllerAutoConfigPolicyOutput) PoeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigPolicy) pulumi.StringOutput { return v.PoeStatus }).(pulumi.StringOutput)
}

func (o SwitchControllerAutoConfigPolicyOutput) QosPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigPolicy) pulumi.StringOutput { return v.QosPolicy }).(pulumi.StringOutput)
}

func (o SwitchControllerAutoConfigPolicyOutput) StormControlPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigPolicy) pulumi.StringOutput { return v.StormControlPolicy }).(pulumi.StringOutput)
}

func (o SwitchControllerAutoConfigPolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigPolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchControllerAutoConfigPolicyArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerAutoConfigPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerAutoConfigPolicy)(nil)).Elem()
}

func (o SwitchControllerAutoConfigPolicyArrayOutput) ToSwitchControllerAutoConfigPolicyArrayOutput() SwitchControllerAutoConfigPolicyArrayOutput {
	return o
}

func (o SwitchControllerAutoConfigPolicyArrayOutput) ToSwitchControllerAutoConfigPolicyArrayOutputWithContext(ctx context.Context) SwitchControllerAutoConfigPolicyArrayOutput {
	return o
}

func (o SwitchControllerAutoConfigPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SwitchControllerAutoConfigPolicy] {
	return pulumix.Output[[]*SwitchControllerAutoConfigPolicy]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerAutoConfigPolicyArrayOutput) Index(i pulumi.IntInput) SwitchControllerAutoConfigPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerAutoConfigPolicy {
		return vs[0].([]*SwitchControllerAutoConfigPolicy)[vs[1].(int)]
	}).(SwitchControllerAutoConfigPolicyOutput)
}

type SwitchControllerAutoConfigPolicyMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerAutoConfigPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerAutoConfigPolicy)(nil)).Elem()
}

func (o SwitchControllerAutoConfigPolicyMapOutput) ToSwitchControllerAutoConfigPolicyMapOutput() SwitchControllerAutoConfigPolicyMapOutput {
	return o
}

func (o SwitchControllerAutoConfigPolicyMapOutput) ToSwitchControllerAutoConfigPolicyMapOutputWithContext(ctx context.Context) SwitchControllerAutoConfigPolicyMapOutput {
	return o
}

func (o SwitchControllerAutoConfigPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SwitchControllerAutoConfigPolicy] {
	return pulumix.Output[map[string]*SwitchControllerAutoConfigPolicy]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerAutoConfigPolicyMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerAutoConfigPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerAutoConfigPolicy {
		return vs[0].(map[string]*SwitchControllerAutoConfigPolicy)[vs[1].(string)]
	}).(SwitchControllerAutoConfigPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerAutoConfigPolicyInput)(nil)).Elem(), &SwitchControllerAutoConfigPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerAutoConfigPolicyArrayInput)(nil)).Elem(), SwitchControllerAutoConfigPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerAutoConfigPolicyMapInput)(nil)).Elem(), SwitchControllerAutoConfigPolicyMap{})
	pulumi.RegisterOutputType(SwitchControllerAutoConfigPolicyOutput{})
	pulumi.RegisterOutputType(SwitchControllerAutoConfigPolicyArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerAutoConfigPolicyMapOutput{})
}

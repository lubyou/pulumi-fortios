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

type SwitchControllerStormControlPolicy struct {
	pulumi.CustomResourceState

	Broadcast        pulumi.StringOutput    `pulumi:"broadcast"`
	Description      pulumi.StringOutput    `pulumi:"description"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Rate             pulumi.IntOutput       `pulumi:"rate"`
	StormControlMode pulumi.StringOutput    `pulumi:"stormControlMode"`
	UnknownMulticast pulumi.StringOutput    `pulumi:"unknownMulticast"`
	UnknownUnicast   pulumi.StringOutput    `pulumi:"unknownUnicast"`
	Vdomparam        pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerStormControlPolicy registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerStormControlPolicy(ctx *pulumi.Context,
	name string, args *SwitchControllerStormControlPolicyArgs, opts ...pulumi.ResourceOption) (*SwitchControllerStormControlPolicy, error) {
	if args == nil {
		args = &SwitchControllerStormControlPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwitchControllerStormControlPolicy
	err := ctx.RegisterResource("fortios:index/switchControllerStormControlPolicy:SwitchControllerStormControlPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerStormControlPolicy gets an existing SwitchControllerStormControlPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerStormControlPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerStormControlPolicyState, opts ...pulumi.ResourceOption) (*SwitchControllerStormControlPolicy, error) {
	var resource SwitchControllerStormControlPolicy
	err := ctx.ReadResource("fortios:index/switchControllerStormControlPolicy:SwitchControllerStormControlPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerStormControlPolicy resources.
type switchControllerStormControlPolicyState struct {
	Broadcast        *string `pulumi:"broadcast"`
	Description      *string `pulumi:"description"`
	Name             *string `pulumi:"name"`
	Rate             *int    `pulumi:"rate"`
	StormControlMode *string `pulumi:"stormControlMode"`
	UnknownMulticast *string `pulumi:"unknownMulticast"`
	UnknownUnicast   *string `pulumi:"unknownUnicast"`
	Vdomparam        *string `pulumi:"vdomparam"`
}

type SwitchControllerStormControlPolicyState struct {
	Broadcast        pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Rate             pulumi.IntPtrInput
	StormControlMode pulumi.StringPtrInput
	UnknownMulticast pulumi.StringPtrInput
	UnknownUnicast   pulumi.StringPtrInput
	Vdomparam        pulumi.StringPtrInput
}

func (SwitchControllerStormControlPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerStormControlPolicyState)(nil)).Elem()
}

type switchControllerStormControlPolicyArgs struct {
	Broadcast        *string `pulumi:"broadcast"`
	Description      *string `pulumi:"description"`
	Name             *string `pulumi:"name"`
	Rate             *int    `pulumi:"rate"`
	StormControlMode *string `pulumi:"stormControlMode"`
	UnknownMulticast *string `pulumi:"unknownMulticast"`
	UnknownUnicast   *string `pulumi:"unknownUnicast"`
	Vdomparam        *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerStormControlPolicy resource.
type SwitchControllerStormControlPolicyArgs struct {
	Broadcast        pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Rate             pulumi.IntPtrInput
	StormControlMode pulumi.StringPtrInput
	UnknownMulticast pulumi.StringPtrInput
	UnknownUnicast   pulumi.StringPtrInput
	Vdomparam        pulumi.StringPtrInput
}

func (SwitchControllerStormControlPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerStormControlPolicyArgs)(nil)).Elem()
}

type SwitchControllerStormControlPolicyInput interface {
	pulumi.Input

	ToSwitchControllerStormControlPolicyOutput() SwitchControllerStormControlPolicyOutput
	ToSwitchControllerStormControlPolicyOutputWithContext(ctx context.Context) SwitchControllerStormControlPolicyOutput
}

func (*SwitchControllerStormControlPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerStormControlPolicy)(nil)).Elem()
}

func (i *SwitchControllerStormControlPolicy) ToSwitchControllerStormControlPolicyOutput() SwitchControllerStormControlPolicyOutput {
	return i.ToSwitchControllerStormControlPolicyOutputWithContext(context.Background())
}

func (i *SwitchControllerStormControlPolicy) ToSwitchControllerStormControlPolicyOutputWithContext(ctx context.Context) SwitchControllerStormControlPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerStormControlPolicyOutput)
}

func (i *SwitchControllerStormControlPolicy) ToOutput(ctx context.Context) pulumix.Output[*SwitchControllerStormControlPolicy] {
	return pulumix.Output[*SwitchControllerStormControlPolicy]{
		OutputState: i.ToSwitchControllerStormControlPolicyOutputWithContext(ctx).OutputState,
	}
}

// SwitchControllerStormControlPolicyArrayInput is an input type that accepts SwitchControllerStormControlPolicyArray and SwitchControllerStormControlPolicyArrayOutput values.
// You can construct a concrete instance of `SwitchControllerStormControlPolicyArrayInput` via:
//
//	SwitchControllerStormControlPolicyArray{ SwitchControllerStormControlPolicyArgs{...} }
type SwitchControllerStormControlPolicyArrayInput interface {
	pulumi.Input

	ToSwitchControllerStormControlPolicyArrayOutput() SwitchControllerStormControlPolicyArrayOutput
	ToSwitchControllerStormControlPolicyArrayOutputWithContext(context.Context) SwitchControllerStormControlPolicyArrayOutput
}

type SwitchControllerStormControlPolicyArray []SwitchControllerStormControlPolicyInput

func (SwitchControllerStormControlPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerStormControlPolicy)(nil)).Elem()
}

func (i SwitchControllerStormControlPolicyArray) ToSwitchControllerStormControlPolicyArrayOutput() SwitchControllerStormControlPolicyArrayOutput {
	return i.ToSwitchControllerStormControlPolicyArrayOutputWithContext(context.Background())
}

func (i SwitchControllerStormControlPolicyArray) ToSwitchControllerStormControlPolicyArrayOutputWithContext(ctx context.Context) SwitchControllerStormControlPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerStormControlPolicyArrayOutput)
}

func (i SwitchControllerStormControlPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*SwitchControllerStormControlPolicy] {
	return pulumix.Output[[]*SwitchControllerStormControlPolicy]{
		OutputState: i.ToSwitchControllerStormControlPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// SwitchControllerStormControlPolicyMapInput is an input type that accepts SwitchControllerStormControlPolicyMap and SwitchControllerStormControlPolicyMapOutput values.
// You can construct a concrete instance of `SwitchControllerStormControlPolicyMapInput` via:
//
//	SwitchControllerStormControlPolicyMap{ "key": SwitchControllerStormControlPolicyArgs{...} }
type SwitchControllerStormControlPolicyMapInput interface {
	pulumi.Input

	ToSwitchControllerStormControlPolicyMapOutput() SwitchControllerStormControlPolicyMapOutput
	ToSwitchControllerStormControlPolicyMapOutputWithContext(context.Context) SwitchControllerStormControlPolicyMapOutput
}

type SwitchControllerStormControlPolicyMap map[string]SwitchControllerStormControlPolicyInput

func (SwitchControllerStormControlPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerStormControlPolicy)(nil)).Elem()
}

func (i SwitchControllerStormControlPolicyMap) ToSwitchControllerStormControlPolicyMapOutput() SwitchControllerStormControlPolicyMapOutput {
	return i.ToSwitchControllerStormControlPolicyMapOutputWithContext(context.Background())
}

func (i SwitchControllerStormControlPolicyMap) ToSwitchControllerStormControlPolicyMapOutputWithContext(ctx context.Context) SwitchControllerStormControlPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerStormControlPolicyMapOutput)
}

func (i SwitchControllerStormControlPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SwitchControllerStormControlPolicy] {
	return pulumix.Output[map[string]*SwitchControllerStormControlPolicy]{
		OutputState: i.ToSwitchControllerStormControlPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type SwitchControllerStormControlPolicyOutput struct{ *pulumi.OutputState }

func (SwitchControllerStormControlPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerStormControlPolicy)(nil)).Elem()
}

func (o SwitchControllerStormControlPolicyOutput) ToSwitchControllerStormControlPolicyOutput() SwitchControllerStormControlPolicyOutput {
	return o
}

func (o SwitchControllerStormControlPolicyOutput) ToSwitchControllerStormControlPolicyOutputWithContext(ctx context.Context) SwitchControllerStormControlPolicyOutput {
	return o
}

func (o SwitchControllerStormControlPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*SwitchControllerStormControlPolicy] {
	return pulumix.Output[*SwitchControllerStormControlPolicy]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerStormControlPolicyOutput) Broadcast() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerStormControlPolicy) pulumi.StringOutput { return v.Broadcast }).(pulumi.StringOutput)
}

func (o SwitchControllerStormControlPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerStormControlPolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o SwitchControllerStormControlPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerStormControlPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwitchControllerStormControlPolicyOutput) Rate() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerStormControlPolicy) pulumi.IntOutput { return v.Rate }).(pulumi.IntOutput)
}

func (o SwitchControllerStormControlPolicyOutput) StormControlMode() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerStormControlPolicy) pulumi.StringOutput { return v.StormControlMode }).(pulumi.StringOutput)
}

func (o SwitchControllerStormControlPolicyOutput) UnknownMulticast() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerStormControlPolicy) pulumi.StringOutput { return v.UnknownMulticast }).(pulumi.StringOutput)
}

func (o SwitchControllerStormControlPolicyOutput) UnknownUnicast() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerStormControlPolicy) pulumi.StringOutput { return v.UnknownUnicast }).(pulumi.StringOutput)
}

func (o SwitchControllerStormControlPolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerStormControlPolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchControllerStormControlPolicyArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerStormControlPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerStormControlPolicy)(nil)).Elem()
}

func (o SwitchControllerStormControlPolicyArrayOutput) ToSwitchControllerStormControlPolicyArrayOutput() SwitchControllerStormControlPolicyArrayOutput {
	return o
}

func (o SwitchControllerStormControlPolicyArrayOutput) ToSwitchControllerStormControlPolicyArrayOutputWithContext(ctx context.Context) SwitchControllerStormControlPolicyArrayOutput {
	return o
}

func (o SwitchControllerStormControlPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SwitchControllerStormControlPolicy] {
	return pulumix.Output[[]*SwitchControllerStormControlPolicy]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerStormControlPolicyArrayOutput) Index(i pulumi.IntInput) SwitchControllerStormControlPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerStormControlPolicy {
		return vs[0].([]*SwitchControllerStormControlPolicy)[vs[1].(int)]
	}).(SwitchControllerStormControlPolicyOutput)
}

type SwitchControllerStormControlPolicyMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerStormControlPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerStormControlPolicy)(nil)).Elem()
}

func (o SwitchControllerStormControlPolicyMapOutput) ToSwitchControllerStormControlPolicyMapOutput() SwitchControllerStormControlPolicyMapOutput {
	return o
}

func (o SwitchControllerStormControlPolicyMapOutput) ToSwitchControllerStormControlPolicyMapOutputWithContext(ctx context.Context) SwitchControllerStormControlPolicyMapOutput {
	return o
}

func (o SwitchControllerStormControlPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SwitchControllerStormControlPolicy] {
	return pulumix.Output[map[string]*SwitchControllerStormControlPolicy]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerStormControlPolicyMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerStormControlPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerStormControlPolicy {
		return vs[0].(map[string]*SwitchControllerStormControlPolicy)[vs[1].(string)]
	}).(SwitchControllerStormControlPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerStormControlPolicyInput)(nil)).Elem(), &SwitchControllerStormControlPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerStormControlPolicyArrayInput)(nil)).Elem(), SwitchControllerStormControlPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerStormControlPolicyMapInput)(nil)).Elem(), SwitchControllerStormControlPolicyMap{})
	pulumi.RegisterOutputType(SwitchControllerStormControlPolicyOutput{})
	pulumi.RegisterOutputType(SwitchControllerStormControlPolicyArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerStormControlPolicyMapOutput{})
}

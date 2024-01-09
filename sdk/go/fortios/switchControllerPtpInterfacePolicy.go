// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SwitchControllerPtpInterfacePolicy struct {
	pulumi.CustomResourceState

	Description pulumi.StringOutput    `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Vdomparam   pulumi.StringPtrOutput `pulumi:"vdomparam"`
	Vlan        pulumi.StringOutput    `pulumi:"vlan"`
	VlanPri     pulumi.IntOutput       `pulumi:"vlanPri"`
}

// NewSwitchControllerPtpInterfacePolicy registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerPtpInterfacePolicy(ctx *pulumi.Context,
	name string, args *SwitchControllerPtpInterfacePolicyArgs, opts ...pulumi.ResourceOption) (*SwitchControllerPtpInterfacePolicy, error) {
	if args == nil {
		args = &SwitchControllerPtpInterfacePolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwitchControllerPtpInterfacePolicy
	err := ctx.RegisterResource("fortios:index/switchControllerPtpInterfacePolicy:SwitchControllerPtpInterfacePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerPtpInterfacePolicy gets an existing SwitchControllerPtpInterfacePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerPtpInterfacePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerPtpInterfacePolicyState, opts ...pulumi.ResourceOption) (*SwitchControllerPtpInterfacePolicy, error) {
	var resource SwitchControllerPtpInterfacePolicy
	err := ctx.ReadResource("fortios:index/switchControllerPtpInterfacePolicy:SwitchControllerPtpInterfacePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerPtpInterfacePolicy resources.
type switchControllerPtpInterfacePolicyState struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Vdomparam   *string `pulumi:"vdomparam"`
	Vlan        *string `pulumi:"vlan"`
	VlanPri     *int    `pulumi:"vlanPri"`
}

type SwitchControllerPtpInterfacePolicyState struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Vdomparam   pulumi.StringPtrInput
	Vlan        pulumi.StringPtrInput
	VlanPri     pulumi.IntPtrInput
}

func (SwitchControllerPtpInterfacePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerPtpInterfacePolicyState)(nil)).Elem()
}

type switchControllerPtpInterfacePolicyArgs struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Vdomparam   *string `pulumi:"vdomparam"`
	Vlan        *string `pulumi:"vlan"`
	VlanPri     *int    `pulumi:"vlanPri"`
}

// The set of arguments for constructing a SwitchControllerPtpInterfacePolicy resource.
type SwitchControllerPtpInterfacePolicyArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Vdomparam   pulumi.StringPtrInput
	Vlan        pulumi.StringPtrInput
	VlanPri     pulumi.IntPtrInput
}

func (SwitchControllerPtpInterfacePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerPtpInterfacePolicyArgs)(nil)).Elem()
}

type SwitchControllerPtpInterfacePolicyInput interface {
	pulumi.Input

	ToSwitchControllerPtpInterfacePolicyOutput() SwitchControllerPtpInterfacePolicyOutput
	ToSwitchControllerPtpInterfacePolicyOutputWithContext(ctx context.Context) SwitchControllerPtpInterfacePolicyOutput
}

func (*SwitchControllerPtpInterfacePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerPtpInterfacePolicy)(nil)).Elem()
}

func (i *SwitchControllerPtpInterfacePolicy) ToSwitchControllerPtpInterfacePolicyOutput() SwitchControllerPtpInterfacePolicyOutput {
	return i.ToSwitchControllerPtpInterfacePolicyOutputWithContext(context.Background())
}

func (i *SwitchControllerPtpInterfacePolicy) ToSwitchControllerPtpInterfacePolicyOutputWithContext(ctx context.Context) SwitchControllerPtpInterfacePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerPtpInterfacePolicyOutput)
}

// SwitchControllerPtpInterfacePolicyArrayInput is an input type that accepts SwitchControllerPtpInterfacePolicyArray and SwitchControllerPtpInterfacePolicyArrayOutput values.
// You can construct a concrete instance of `SwitchControllerPtpInterfacePolicyArrayInput` via:
//
//	SwitchControllerPtpInterfacePolicyArray{ SwitchControllerPtpInterfacePolicyArgs{...} }
type SwitchControllerPtpInterfacePolicyArrayInput interface {
	pulumi.Input

	ToSwitchControllerPtpInterfacePolicyArrayOutput() SwitchControllerPtpInterfacePolicyArrayOutput
	ToSwitchControllerPtpInterfacePolicyArrayOutputWithContext(context.Context) SwitchControllerPtpInterfacePolicyArrayOutput
}

type SwitchControllerPtpInterfacePolicyArray []SwitchControllerPtpInterfacePolicyInput

func (SwitchControllerPtpInterfacePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerPtpInterfacePolicy)(nil)).Elem()
}

func (i SwitchControllerPtpInterfacePolicyArray) ToSwitchControllerPtpInterfacePolicyArrayOutput() SwitchControllerPtpInterfacePolicyArrayOutput {
	return i.ToSwitchControllerPtpInterfacePolicyArrayOutputWithContext(context.Background())
}

func (i SwitchControllerPtpInterfacePolicyArray) ToSwitchControllerPtpInterfacePolicyArrayOutputWithContext(ctx context.Context) SwitchControllerPtpInterfacePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerPtpInterfacePolicyArrayOutput)
}

// SwitchControllerPtpInterfacePolicyMapInput is an input type that accepts SwitchControllerPtpInterfacePolicyMap and SwitchControllerPtpInterfacePolicyMapOutput values.
// You can construct a concrete instance of `SwitchControllerPtpInterfacePolicyMapInput` via:
//
//	SwitchControllerPtpInterfacePolicyMap{ "key": SwitchControllerPtpInterfacePolicyArgs{...} }
type SwitchControllerPtpInterfacePolicyMapInput interface {
	pulumi.Input

	ToSwitchControllerPtpInterfacePolicyMapOutput() SwitchControllerPtpInterfacePolicyMapOutput
	ToSwitchControllerPtpInterfacePolicyMapOutputWithContext(context.Context) SwitchControllerPtpInterfacePolicyMapOutput
}

type SwitchControllerPtpInterfacePolicyMap map[string]SwitchControllerPtpInterfacePolicyInput

func (SwitchControllerPtpInterfacePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerPtpInterfacePolicy)(nil)).Elem()
}

func (i SwitchControllerPtpInterfacePolicyMap) ToSwitchControllerPtpInterfacePolicyMapOutput() SwitchControllerPtpInterfacePolicyMapOutput {
	return i.ToSwitchControllerPtpInterfacePolicyMapOutputWithContext(context.Background())
}

func (i SwitchControllerPtpInterfacePolicyMap) ToSwitchControllerPtpInterfacePolicyMapOutputWithContext(ctx context.Context) SwitchControllerPtpInterfacePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerPtpInterfacePolicyMapOutput)
}

type SwitchControllerPtpInterfacePolicyOutput struct{ *pulumi.OutputState }

func (SwitchControllerPtpInterfacePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerPtpInterfacePolicy)(nil)).Elem()
}

func (o SwitchControllerPtpInterfacePolicyOutput) ToSwitchControllerPtpInterfacePolicyOutput() SwitchControllerPtpInterfacePolicyOutput {
	return o
}

func (o SwitchControllerPtpInterfacePolicyOutput) ToSwitchControllerPtpInterfacePolicyOutputWithContext(ctx context.Context) SwitchControllerPtpInterfacePolicyOutput {
	return o
}

func (o SwitchControllerPtpInterfacePolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerPtpInterfacePolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o SwitchControllerPtpInterfacePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerPtpInterfacePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwitchControllerPtpInterfacePolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerPtpInterfacePolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o SwitchControllerPtpInterfacePolicyOutput) Vlan() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerPtpInterfacePolicy) pulumi.StringOutput { return v.Vlan }).(pulumi.StringOutput)
}

func (o SwitchControllerPtpInterfacePolicyOutput) VlanPri() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerPtpInterfacePolicy) pulumi.IntOutput { return v.VlanPri }).(pulumi.IntOutput)
}

type SwitchControllerPtpInterfacePolicyArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerPtpInterfacePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerPtpInterfacePolicy)(nil)).Elem()
}

func (o SwitchControllerPtpInterfacePolicyArrayOutput) ToSwitchControllerPtpInterfacePolicyArrayOutput() SwitchControllerPtpInterfacePolicyArrayOutput {
	return o
}

func (o SwitchControllerPtpInterfacePolicyArrayOutput) ToSwitchControllerPtpInterfacePolicyArrayOutputWithContext(ctx context.Context) SwitchControllerPtpInterfacePolicyArrayOutput {
	return o
}

func (o SwitchControllerPtpInterfacePolicyArrayOutput) Index(i pulumi.IntInput) SwitchControllerPtpInterfacePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerPtpInterfacePolicy {
		return vs[0].([]*SwitchControllerPtpInterfacePolicy)[vs[1].(int)]
	}).(SwitchControllerPtpInterfacePolicyOutput)
}

type SwitchControllerPtpInterfacePolicyMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerPtpInterfacePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerPtpInterfacePolicy)(nil)).Elem()
}

func (o SwitchControllerPtpInterfacePolicyMapOutput) ToSwitchControllerPtpInterfacePolicyMapOutput() SwitchControllerPtpInterfacePolicyMapOutput {
	return o
}

func (o SwitchControllerPtpInterfacePolicyMapOutput) ToSwitchControllerPtpInterfacePolicyMapOutputWithContext(ctx context.Context) SwitchControllerPtpInterfacePolicyMapOutput {
	return o
}

func (o SwitchControllerPtpInterfacePolicyMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerPtpInterfacePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerPtpInterfacePolicy {
		return vs[0].(map[string]*SwitchControllerPtpInterfacePolicy)[vs[1].(string)]
	}).(SwitchControllerPtpInterfacePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerPtpInterfacePolicyInput)(nil)).Elem(), &SwitchControllerPtpInterfacePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerPtpInterfacePolicyArrayInput)(nil)).Elem(), SwitchControllerPtpInterfacePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerPtpInterfacePolicyMapInput)(nil)).Elem(), SwitchControllerPtpInterfacePolicyMap{})
	pulumi.RegisterOutputType(SwitchControllerPtpInterfacePolicyOutput{})
	pulumi.RegisterOutputType(SwitchControllerPtpInterfacePolicyArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerPtpInterfacePolicyMapOutput{})
}

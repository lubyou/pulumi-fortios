// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SwitchControllerSecurityPolicyCaptivePortal struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput    `pulumi:"name"`
	PolicyType pulumi.StringOutput    `pulumi:"policyType"`
	Vdomparam  pulumi.StringPtrOutput `pulumi:"vdomparam"`
	Vlan       pulumi.StringOutput    `pulumi:"vlan"`
}

// NewSwitchControllerSecurityPolicyCaptivePortal registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerSecurityPolicyCaptivePortal(ctx *pulumi.Context,
	name string, args *SwitchControllerSecurityPolicyCaptivePortalArgs, opts ...pulumi.ResourceOption) (*SwitchControllerSecurityPolicyCaptivePortal, error) {
	if args == nil {
		args = &SwitchControllerSecurityPolicyCaptivePortalArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerSecurityPolicyCaptivePortal
	err := ctx.RegisterResource("fortios:index/switchControllerSecurityPolicyCaptivePortal:SwitchControllerSecurityPolicyCaptivePortal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerSecurityPolicyCaptivePortal gets an existing SwitchControllerSecurityPolicyCaptivePortal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerSecurityPolicyCaptivePortal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerSecurityPolicyCaptivePortalState, opts ...pulumi.ResourceOption) (*SwitchControllerSecurityPolicyCaptivePortal, error) {
	var resource SwitchControllerSecurityPolicyCaptivePortal
	err := ctx.ReadResource("fortios:index/switchControllerSecurityPolicyCaptivePortal:SwitchControllerSecurityPolicyCaptivePortal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerSecurityPolicyCaptivePortal resources.
type switchControllerSecurityPolicyCaptivePortalState struct {
	Name       *string `pulumi:"name"`
	PolicyType *string `pulumi:"policyType"`
	Vdomparam  *string `pulumi:"vdomparam"`
	Vlan       *string `pulumi:"vlan"`
}

type SwitchControllerSecurityPolicyCaptivePortalState struct {
	Name       pulumi.StringPtrInput
	PolicyType pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
	Vlan       pulumi.StringPtrInput
}

func (SwitchControllerSecurityPolicyCaptivePortalState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSecurityPolicyCaptivePortalState)(nil)).Elem()
}

type switchControllerSecurityPolicyCaptivePortalArgs struct {
	Name       *string `pulumi:"name"`
	PolicyType *string `pulumi:"policyType"`
	Vdomparam  *string `pulumi:"vdomparam"`
	Vlan       *string `pulumi:"vlan"`
}

// The set of arguments for constructing a SwitchControllerSecurityPolicyCaptivePortal resource.
type SwitchControllerSecurityPolicyCaptivePortalArgs struct {
	Name       pulumi.StringPtrInput
	PolicyType pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
	Vlan       pulumi.StringPtrInput
}

func (SwitchControllerSecurityPolicyCaptivePortalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSecurityPolicyCaptivePortalArgs)(nil)).Elem()
}

type SwitchControllerSecurityPolicyCaptivePortalInput interface {
	pulumi.Input

	ToSwitchControllerSecurityPolicyCaptivePortalOutput() SwitchControllerSecurityPolicyCaptivePortalOutput
	ToSwitchControllerSecurityPolicyCaptivePortalOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyCaptivePortalOutput
}

func (*SwitchControllerSecurityPolicyCaptivePortal) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSecurityPolicyCaptivePortal)(nil)).Elem()
}

func (i *SwitchControllerSecurityPolicyCaptivePortal) ToSwitchControllerSecurityPolicyCaptivePortalOutput() SwitchControllerSecurityPolicyCaptivePortalOutput {
	return i.ToSwitchControllerSecurityPolicyCaptivePortalOutputWithContext(context.Background())
}

func (i *SwitchControllerSecurityPolicyCaptivePortal) ToSwitchControllerSecurityPolicyCaptivePortalOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyCaptivePortalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSecurityPolicyCaptivePortalOutput)
}

// SwitchControllerSecurityPolicyCaptivePortalArrayInput is an input type that accepts SwitchControllerSecurityPolicyCaptivePortalArray and SwitchControllerSecurityPolicyCaptivePortalArrayOutput values.
// You can construct a concrete instance of `SwitchControllerSecurityPolicyCaptivePortalArrayInput` via:
//
//	SwitchControllerSecurityPolicyCaptivePortalArray{ SwitchControllerSecurityPolicyCaptivePortalArgs{...} }
type SwitchControllerSecurityPolicyCaptivePortalArrayInput interface {
	pulumi.Input

	ToSwitchControllerSecurityPolicyCaptivePortalArrayOutput() SwitchControllerSecurityPolicyCaptivePortalArrayOutput
	ToSwitchControllerSecurityPolicyCaptivePortalArrayOutputWithContext(context.Context) SwitchControllerSecurityPolicyCaptivePortalArrayOutput
}

type SwitchControllerSecurityPolicyCaptivePortalArray []SwitchControllerSecurityPolicyCaptivePortalInput

func (SwitchControllerSecurityPolicyCaptivePortalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSecurityPolicyCaptivePortal)(nil)).Elem()
}

func (i SwitchControllerSecurityPolicyCaptivePortalArray) ToSwitchControllerSecurityPolicyCaptivePortalArrayOutput() SwitchControllerSecurityPolicyCaptivePortalArrayOutput {
	return i.ToSwitchControllerSecurityPolicyCaptivePortalArrayOutputWithContext(context.Background())
}

func (i SwitchControllerSecurityPolicyCaptivePortalArray) ToSwitchControllerSecurityPolicyCaptivePortalArrayOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyCaptivePortalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSecurityPolicyCaptivePortalArrayOutput)
}

// SwitchControllerSecurityPolicyCaptivePortalMapInput is an input type that accepts SwitchControllerSecurityPolicyCaptivePortalMap and SwitchControllerSecurityPolicyCaptivePortalMapOutput values.
// You can construct a concrete instance of `SwitchControllerSecurityPolicyCaptivePortalMapInput` via:
//
//	SwitchControllerSecurityPolicyCaptivePortalMap{ "key": SwitchControllerSecurityPolicyCaptivePortalArgs{...} }
type SwitchControllerSecurityPolicyCaptivePortalMapInput interface {
	pulumi.Input

	ToSwitchControllerSecurityPolicyCaptivePortalMapOutput() SwitchControllerSecurityPolicyCaptivePortalMapOutput
	ToSwitchControllerSecurityPolicyCaptivePortalMapOutputWithContext(context.Context) SwitchControllerSecurityPolicyCaptivePortalMapOutput
}

type SwitchControllerSecurityPolicyCaptivePortalMap map[string]SwitchControllerSecurityPolicyCaptivePortalInput

func (SwitchControllerSecurityPolicyCaptivePortalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSecurityPolicyCaptivePortal)(nil)).Elem()
}

func (i SwitchControllerSecurityPolicyCaptivePortalMap) ToSwitchControllerSecurityPolicyCaptivePortalMapOutput() SwitchControllerSecurityPolicyCaptivePortalMapOutput {
	return i.ToSwitchControllerSecurityPolicyCaptivePortalMapOutputWithContext(context.Background())
}

func (i SwitchControllerSecurityPolicyCaptivePortalMap) ToSwitchControllerSecurityPolicyCaptivePortalMapOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyCaptivePortalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSecurityPolicyCaptivePortalMapOutput)
}

type SwitchControllerSecurityPolicyCaptivePortalOutput struct{ *pulumi.OutputState }

func (SwitchControllerSecurityPolicyCaptivePortalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSecurityPolicyCaptivePortal)(nil)).Elem()
}

func (o SwitchControllerSecurityPolicyCaptivePortalOutput) ToSwitchControllerSecurityPolicyCaptivePortalOutput() SwitchControllerSecurityPolicyCaptivePortalOutput {
	return o
}

func (o SwitchControllerSecurityPolicyCaptivePortalOutput) ToSwitchControllerSecurityPolicyCaptivePortalOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyCaptivePortalOutput {
	return o
}

func (o SwitchControllerSecurityPolicyCaptivePortalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerSecurityPolicyCaptivePortal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwitchControllerSecurityPolicyCaptivePortalOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerSecurityPolicyCaptivePortal) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

func (o SwitchControllerSecurityPolicyCaptivePortalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerSecurityPolicyCaptivePortal) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o SwitchControllerSecurityPolicyCaptivePortalOutput) Vlan() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerSecurityPolicyCaptivePortal) pulumi.StringOutput { return v.Vlan }).(pulumi.StringOutput)
}

type SwitchControllerSecurityPolicyCaptivePortalArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerSecurityPolicyCaptivePortalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSecurityPolicyCaptivePortal)(nil)).Elem()
}

func (o SwitchControllerSecurityPolicyCaptivePortalArrayOutput) ToSwitchControllerSecurityPolicyCaptivePortalArrayOutput() SwitchControllerSecurityPolicyCaptivePortalArrayOutput {
	return o
}

func (o SwitchControllerSecurityPolicyCaptivePortalArrayOutput) ToSwitchControllerSecurityPolicyCaptivePortalArrayOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyCaptivePortalArrayOutput {
	return o
}

func (o SwitchControllerSecurityPolicyCaptivePortalArrayOutput) Index(i pulumi.IntInput) SwitchControllerSecurityPolicyCaptivePortalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerSecurityPolicyCaptivePortal {
		return vs[0].([]*SwitchControllerSecurityPolicyCaptivePortal)[vs[1].(int)]
	}).(SwitchControllerSecurityPolicyCaptivePortalOutput)
}

type SwitchControllerSecurityPolicyCaptivePortalMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerSecurityPolicyCaptivePortalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSecurityPolicyCaptivePortal)(nil)).Elem()
}

func (o SwitchControllerSecurityPolicyCaptivePortalMapOutput) ToSwitchControllerSecurityPolicyCaptivePortalMapOutput() SwitchControllerSecurityPolicyCaptivePortalMapOutput {
	return o
}

func (o SwitchControllerSecurityPolicyCaptivePortalMapOutput) ToSwitchControllerSecurityPolicyCaptivePortalMapOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyCaptivePortalMapOutput {
	return o
}

func (o SwitchControllerSecurityPolicyCaptivePortalMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerSecurityPolicyCaptivePortalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerSecurityPolicyCaptivePortal {
		return vs[0].(map[string]*SwitchControllerSecurityPolicyCaptivePortal)[vs[1].(string)]
	}).(SwitchControllerSecurityPolicyCaptivePortalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSecurityPolicyCaptivePortalInput)(nil)).Elem(), &SwitchControllerSecurityPolicyCaptivePortal{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSecurityPolicyCaptivePortalArrayInput)(nil)).Elem(), SwitchControllerSecurityPolicyCaptivePortalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSecurityPolicyCaptivePortalMapInput)(nil)).Elem(), SwitchControllerSecurityPolicyCaptivePortalMap{})
	pulumi.RegisterOutputType(SwitchControllerSecurityPolicyCaptivePortalOutput{})
	pulumi.RegisterOutputType(SwitchControllerSecurityPolicyCaptivePortalArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerSecurityPolicyCaptivePortalMapOutput{})
}

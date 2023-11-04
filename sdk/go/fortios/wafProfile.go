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

type WafProfile struct {
	pulumi.CustomResourceState

	AddressList         WafProfileAddressListOutput    `pulumi:"addressList"`
	Comment             pulumi.StringPtrOutput         `pulumi:"comment"`
	Constraint          WafProfileConstraintOutput     `pulumi:"constraint"`
	DynamicSortSubtable pulumi.StringPtrOutput         `pulumi:"dynamicSortSubtable"`
	ExtendedLog         pulumi.StringOutput            `pulumi:"extendedLog"`
	External            pulumi.StringOutput            `pulumi:"external"`
	GetAllTables        pulumi.StringPtrOutput         `pulumi:"getAllTables"`
	Method              WafProfileMethodOutput         `pulumi:"method"`
	Name                pulumi.StringOutput            `pulumi:"name"`
	Signature           WafProfileSignatureOutput      `pulumi:"signature"`
	UrlAccesses         WafProfileUrlAccessArrayOutput `pulumi:"urlAccesses"`
	Vdomparam           pulumi.StringPtrOutput         `pulumi:"vdomparam"`
}

// NewWafProfile registers a new resource with the given unique name, arguments, and options.
func NewWafProfile(ctx *pulumi.Context,
	name string, args *WafProfileArgs, opts ...pulumi.ResourceOption) (*WafProfile, error) {
	if args == nil {
		args = &WafProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WafProfile
	err := ctx.RegisterResource("fortios:index/wafProfile:WafProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWafProfile gets an existing WafProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWafProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WafProfileState, opts ...pulumi.ResourceOption) (*WafProfile, error) {
	var resource WafProfile
	err := ctx.ReadResource("fortios:index/wafProfile:WafProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WafProfile resources.
type wafProfileState struct {
	AddressList         *WafProfileAddressList `pulumi:"addressList"`
	Comment             *string                `pulumi:"comment"`
	Constraint          *WafProfileConstraint  `pulumi:"constraint"`
	DynamicSortSubtable *string                `pulumi:"dynamicSortSubtable"`
	ExtendedLog         *string                `pulumi:"extendedLog"`
	External            *string                `pulumi:"external"`
	GetAllTables        *string                `pulumi:"getAllTables"`
	Method              *WafProfileMethod      `pulumi:"method"`
	Name                *string                `pulumi:"name"`
	Signature           *WafProfileSignature   `pulumi:"signature"`
	UrlAccesses         []WafProfileUrlAccess  `pulumi:"urlAccesses"`
	Vdomparam           *string                `pulumi:"vdomparam"`
}

type WafProfileState struct {
	AddressList         WafProfileAddressListPtrInput
	Comment             pulumi.StringPtrInput
	Constraint          WafProfileConstraintPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	ExtendedLog         pulumi.StringPtrInput
	External            pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Method              WafProfileMethodPtrInput
	Name                pulumi.StringPtrInput
	Signature           WafProfileSignaturePtrInput
	UrlAccesses         WafProfileUrlAccessArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WafProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*wafProfileState)(nil)).Elem()
}

type wafProfileArgs struct {
	AddressList         *WafProfileAddressList `pulumi:"addressList"`
	Comment             *string                `pulumi:"comment"`
	Constraint          *WafProfileConstraint  `pulumi:"constraint"`
	DynamicSortSubtable *string                `pulumi:"dynamicSortSubtable"`
	ExtendedLog         *string                `pulumi:"extendedLog"`
	External            *string                `pulumi:"external"`
	GetAllTables        *string                `pulumi:"getAllTables"`
	Method              *WafProfileMethod      `pulumi:"method"`
	Name                *string                `pulumi:"name"`
	Signature           *WafProfileSignature   `pulumi:"signature"`
	UrlAccesses         []WafProfileUrlAccess  `pulumi:"urlAccesses"`
	Vdomparam           *string                `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WafProfile resource.
type WafProfileArgs struct {
	AddressList         WafProfileAddressListPtrInput
	Comment             pulumi.StringPtrInput
	Constraint          WafProfileConstraintPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	ExtendedLog         pulumi.StringPtrInput
	External            pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Method              WafProfileMethodPtrInput
	Name                pulumi.StringPtrInput
	Signature           WafProfileSignaturePtrInput
	UrlAccesses         WafProfileUrlAccessArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WafProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wafProfileArgs)(nil)).Elem()
}

type WafProfileInput interface {
	pulumi.Input

	ToWafProfileOutput() WafProfileOutput
	ToWafProfileOutputWithContext(ctx context.Context) WafProfileOutput
}

func (*WafProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**WafProfile)(nil)).Elem()
}

func (i *WafProfile) ToWafProfileOutput() WafProfileOutput {
	return i.ToWafProfileOutputWithContext(context.Background())
}

func (i *WafProfile) ToWafProfileOutputWithContext(ctx context.Context) WafProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafProfileOutput)
}

func (i *WafProfile) ToOutput(ctx context.Context) pulumix.Output[*WafProfile] {
	return pulumix.Output[*WafProfile]{
		OutputState: i.ToWafProfileOutputWithContext(ctx).OutputState,
	}
}

// WafProfileArrayInput is an input type that accepts WafProfileArray and WafProfileArrayOutput values.
// You can construct a concrete instance of `WafProfileArrayInput` via:
//
//	WafProfileArray{ WafProfileArgs{...} }
type WafProfileArrayInput interface {
	pulumi.Input

	ToWafProfileArrayOutput() WafProfileArrayOutput
	ToWafProfileArrayOutputWithContext(context.Context) WafProfileArrayOutput
}

type WafProfileArray []WafProfileInput

func (WafProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafProfile)(nil)).Elem()
}

func (i WafProfileArray) ToWafProfileArrayOutput() WafProfileArrayOutput {
	return i.ToWafProfileArrayOutputWithContext(context.Background())
}

func (i WafProfileArray) ToWafProfileArrayOutputWithContext(ctx context.Context) WafProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafProfileArrayOutput)
}

func (i WafProfileArray) ToOutput(ctx context.Context) pulumix.Output[[]*WafProfile] {
	return pulumix.Output[[]*WafProfile]{
		OutputState: i.ToWafProfileArrayOutputWithContext(ctx).OutputState,
	}
}

// WafProfileMapInput is an input type that accepts WafProfileMap and WafProfileMapOutput values.
// You can construct a concrete instance of `WafProfileMapInput` via:
//
//	WafProfileMap{ "key": WafProfileArgs{...} }
type WafProfileMapInput interface {
	pulumi.Input

	ToWafProfileMapOutput() WafProfileMapOutput
	ToWafProfileMapOutputWithContext(context.Context) WafProfileMapOutput
}

type WafProfileMap map[string]WafProfileInput

func (WafProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafProfile)(nil)).Elem()
}

func (i WafProfileMap) ToWafProfileMapOutput() WafProfileMapOutput {
	return i.ToWafProfileMapOutputWithContext(context.Background())
}

func (i WafProfileMap) ToWafProfileMapOutputWithContext(ctx context.Context) WafProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafProfileMapOutput)
}

func (i WafProfileMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WafProfile] {
	return pulumix.Output[map[string]*WafProfile]{
		OutputState: i.ToWafProfileMapOutputWithContext(ctx).OutputState,
	}
}

type WafProfileOutput struct{ *pulumi.OutputState }

func (WafProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WafProfile)(nil)).Elem()
}

func (o WafProfileOutput) ToWafProfileOutput() WafProfileOutput {
	return o
}

func (o WafProfileOutput) ToWafProfileOutputWithContext(ctx context.Context) WafProfileOutput {
	return o
}

func (o WafProfileOutput) ToOutput(ctx context.Context) pulumix.Output[*WafProfile] {
	return pulumix.Output[*WafProfile]{
		OutputState: o.OutputState,
	}
}

func (o WafProfileOutput) AddressList() WafProfileAddressListOutput {
	return o.ApplyT(func(v *WafProfile) WafProfileAddressListOutput { return v.AddressList }).(WafProfileAddressListOutput)
}

func (o WafProfileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafProfile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o WafProfileOutput) Constraint() WafProfileConstraintOutput {
	return o.ApplyT(func(v *WafProfile) WafProfileConstraintOutput { return v.Constraint }).(WafProfileConstraintOutput)
}

func (o WafProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafProfile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WafProfileOutput) ExtendedLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WafProfile) pulumi.StringOutput { return v.ExtendedLog }).(pulumi.StringOutput)
}

func (o WafProfileOutput) External() pulumi.StringOutput {
	return o.ApplyT(func(v *WafProfile) pulumi.StringOutput { return v.External }).(pulumi.StringOutput)
}

func (o WafProfileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafProfile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WafProfileOutput) Method() WafProfileMethodOutput {
	return o.ApplyT(func(v *WafProfile) WafProfileMethodOutput { return v.Method }).(WafProfileMethodOutput)
}

func (o WafProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WafProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WafProfileOutput) Signature() WafProfileSignatureOutput {
	return o.ApplyT(func(v *WafProfile) WafProfileSignatureOutput { return v.Signature }).(WafProfileSignatureOutput)
}

func (o WafProfileOutput) UrlAccesses() WafProfileUrlAccessArrayOutput {
	return o.ApplyT(func(v *WafProfile) WafProfileUrlAccessArrayOutput { return v.UrlAccesses }).(WafProfileUrlAccessArrayOutput)
}

func (o WafProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WafProfileArrayOutput struct{ *pulumi.OutputState }

func (WafProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafProfile)(nil)).Elem()
}

func (o WafProfileArrayOutput) ToWafProfileArrayOutput() WafProfileArrayOutput {
	return o
}

func (o WafProfileArrayOutput) ToWafProfileArrayOutputWithContext(ctx context.Context) WafProfileArrayOutput {
	return o
}

func (o WafProfileArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WafProfile] {
	return pulumix.Output[[]*WafProfile]{
		OutputState: o.OutputState,
	}
}

func (o WafProfileArrayOutput) Index(i pulumi.IntInput) WafProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WafProfile {
		return vs[0].([]*WafProfile)[vs[1].(int)]
	}).(WafProfileOutput)
}

type WafProfileMapOutput struct{ *pulumi.OutputState }

func (WafProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafProfile)(nil)).Elem()
}

func (o WafProfileMapOutput) ToWafProfileMapOutput() WafProfileMapOutput {
	return o
}

func (o WafProfileMapOutput) ToWafProfileMapOutputWithContext(ctx context.Context) WafProfileMapOutput {
	return o
}

func (o WafProfileMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WafProfile] {
	return pulumix.Output[map[string]*WafProfile]{
		OutputState: o.OutputState,
	}
}

func (o WafProfileMapOutput) MapIndex(k pulumi.StringInput) WafProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WafProfile {
		return vs[0].(map[string]*WafProfile)[vs[1].(string)]
	}).(WafProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WafProfileInput)(nil)).Elem(), &WafProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafProfileArrayInput)(nil)).Elem(), WafProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafProfileMapInput)(nil)).Elem(), WafProfileMap{})
	pulumi.RegisterOutputType(WafProfileOutput{})
	pulumi.RegisterOutputType(WafProfileArrayOutput{})
	pulumi.RegisterOutputType(WafProfileMapOutput{})
}

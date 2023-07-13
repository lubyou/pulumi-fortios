// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WirelessControllerBonjourProfile struct {
	pulumi.CustomResourceState

	Comment             pulumi.StringOutput                                   `pulumi:"comment"`
	DynamicSortSubtable pulumi.StringPtrOutput                                `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                                `pulumi:"getAllTables"`
	Name                pulumi.StringOutput                                   `pulumi:"name"`
	PolicyLists         WirelessControllerBonjourProfilePolicyListArrayOutput `pulumi:"policyLists"`
	Vdomparam           pulumi.StringPtrOutput                                `pulumi:"vdomparam"`
}

// NewWirelessControllerBonjourProfile registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerBonjourProfile(ctx *pulumi.Context,
	name string, args *WirelessControllerBonjourProfileArgs, opts ...pulumi.ResourceOption) (*WirelessControllerBonjourProfile, error) {
	if args == nil {
		args = &WirelessControllerBonjourProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerBonjourProfile
	err := ctx.RegisterResource("fortios:index/wirelessControllerBonjourProfile:WirelessControllerBonjourProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerBonjourProfile gets an existing WirelessControllerBonjourProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerBonjourProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerBonjourProfileState, opts ...pulumi.ResourceOption) (*WirelessControllerBonjourProfile, error) {
	var resource WirelessControllerBonjourProfile
	err := ctx.ReadResource("fortios:index/wirelessControllerBonjourProfile:WirelessControllerBonjourProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerBonjourProfile resources.
type wirelessControllerBonjourProfileState struct {
	Comment             *string                                      `pulumi:"comment"`
	DynamicSortSubtable *string                                      `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                      `pulumi:"getAllTables"`
	Name                *string                                      `pulumi:"name"`
	PolicyLists         []WirelessControllerBonjourProfilePolicyList `pulumi:"policyLists"`
	Vdomparam           *string                                      `pulumi:"vdomparam"`
}

type WirelessControllerBonjourProfileState struct {
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	PolicyLists         WirelessControllerBonjourProfilePolicyListArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerBonjourProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerBonjourProfileState)(nil)).Elem()
}

type wirelessControllerBonjourProfileArgs struct {
	Comment             *string                                      `pulumi:"comment"`
	DynamicSortSubtable *string                                      `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                      `pulumi:"getAllTables"`
	Name                *string                                      `pulumi:"name"`
	PolicyLists         []WirelessControllerBonjourProfilePolicyList `pulumi:"policyLists"`
	Vdomparam           *string                                      `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerBonjourProfile resource.
type WirelessControllerBonjourProfileArgs struct {
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	PolicyLists         WirelessControllerBonjourProfilePolicyListArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerBonjourProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerBonjourProfileArgs)(nil)).Elem()
}

type WirelessControllerBonjourProfileInput interface {
	pulumi.Input

	ToWirelessControllerBonjourProfileOutput() WirelessControllerBonjourProfileOutput
	ToWirelessControllerBonjourProfileOutputWithContext(ctx context.Context) WirelessControllerBonjourProfileOutput
}

func (*WirelessControllerBonjourProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerBonjourProfile)(nil)).Elem()
}

func (i *WirelessControllerBonjourProfile) ToWirelessControllerBonjourProfileOutput() WirelessControllerBonjourProfileOutput {
	return i.ToWirelessControllerBonjourProfileOutputWithContext(context.Background())
}

func (i *WirelessControllerBonjourProfile) ToWirelessControllerBonjourProfileOutputWithContext(ctx context.Context) WirelessControllerBonjourProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerBonjourProfileOutput)
}

// WirelessControllerBonjourProfileArrayInput is an input type that accepts WirelessControllerBonjourProfileArray and WirelessControllerBonjourProfileArrayOutput values.
// You can construct a concrete instance of `WirelessControllerBonjourProfileArrayInput` via:
//
//	WirelessControllerBonjourProfileArray{ WirelessControllerBonjourProfileArgs{...} }
type WirelessControllerBonjourProfileArrayInput interface {
	pulumi.Input

	ToWirelessControllerBonjourProfileArrayOutput() WirelessControllerBonjourProfileArrayOutput
	ToWirelessControllerBonjourProfileArrayOutputWithContext(context.Context) WirelessControllerBonjourProfileArrayOutput
}

type WirelessControllerBonjourProfileArray []WirelessControllerBonjourProfileInput

func (WirelessControllerBonjourProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerBonjourProfile)(nil)).Elem()
}

func (i WirelessControllerBonjourProfileArray) ToWirelessControllerBonjourProfileArrayOutput() WirelessControllerBonjourProfileArrayOutput {
	return i.ToWirelessControllerBonjourProfileArrayOutputWithContext(context.Background())
}

func (i WirelessControllerBonjourProfileArray) ToWirelessControllerBonjourProfileArrayOutputWithContext(ctx context.Context) WirelessControllerBonjourProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerBonjourProfileArrayOutput)
}

// WirelessControllerBonjourProfileMapInput is an input type that accepts WirelessControllerBonjourProfileMap and WirelessControllerBonjourProfileMapOutput values.
// You can construct a concrete instance of `WirelessControllerBonjourProfileMapInput` via:
//
//	WirelessControllerBonjourProfileMap{ "key": WirelessControllerBonjourProfileArgs{...} }
type WirelessControllerBonjourProfileMapInput interface {
	pulumi.Input

	ToWirelessControllerBonjourProfileMapOutput() WirelessControllerBonjourProfileMapOutput
	ToWirelessControllerBonjourProfileMapOutputWithContext(context.Context) WirelessControllerBonjourProfileMapOutput
}

type WirelessControllerBonjourProfileMap map[string]WirelessControllerBonjourProfileInput

func (WirelessControllerBonjourProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerBonjourProfile)(nil)).Elem()
}

func (i WirelessControllerBonjourProfileMap) ToWirelessControllerBonjourProfileMapOutput() WirelessControllerBonjourProfileMapOutput {
	return i.ToWirelessControllerBonjourProfileMapOutputWithContext(context.Background())
}

func (i WirelessControllerBonjourProfileMap) ToWirelessControllerBonjourProfileMapOutputWithContext(ctx context.Context) WirelessControllerBonjourProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerBonjourProfileMapOutput)
}

type WirelessControllerBonjourProfileOutput struct{ *pulumi.OutputState }

func (WirelessControllerBonjourProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerBonjourProfile)(nil)).Elem()
}

func (o WirelessControllerBonjourProfileOutput) ToWirelessControllerBonjourProfileOutput() WirelessControllerBonjourProfileOutput {
	return o
}

func (o WirelessControllerBonjourProfileOutput) ToWirelessControllerBonjourProfileOutputWithContext(ctx context.Context) WirelessControllerBonjourProfileOutput {
	return o
}

func (o WirelessControllerBonjourProfileOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerBonjourProfile) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

func (o WirelessControllerBonjourProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerBonjourProfile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerBonjourProfileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerBonjourProfile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerBonjourProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerBonjourProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WirelessControllerBonjourProfileOutput) PolicyLists() WirelessControllerBonjourProfilePolicyListArrayOutput {
	return o.ApplyT(func(v *WirelessControllerBonjourProfile) WirelessControllerBonjourProfilePolicyListArrayOutput {
		return v.PolicyLists
	}).(WirelessControllerBonjourProfilePolicyListArrayOutput)
}

func (o WirelessControllerBonjourProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerBonjourProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelessControllerBonjourProfileArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerBonjourProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerBonjourProfile)(nil)).Elem()
}

func (o WirelessControllerBonjourProfileArrayOutput) ToWirelessControllerBonjourProfileArrayOutput() WirelessControllerBonjourProfileArrayOutput {
	return o
}

func (o WirelessControllerBonjourProfileArrayOutput) ToWirelessControllerBonjourProfileArrayOutputWithContext(ctx context.Context) WirelessControllerBonjourProfileArrayOutput {
	return o
}

func (o WirelessControllerBonjourProfileArrayOutput) Index(i pulumi.IntInput) WirelessControllerBonjourProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerBonjourProfile {
		return vs[0].([]*WirelessControllerBonjourProfile)[vs[1].(int)]
	}).(WirelessControllerBonjourProfileOutput)
}

type WirelessControllerBonjourProfileMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerBonjourProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerBonjourProfile)(nil)).Elem()
}

func (o WirelessControllerBonjourProfileMapOutput) ToWirelessControllerBonjourProfileMapOutput() WirelessControllerBonjourProfileMapOutput {
	return o
}

func (o WirelessControllerBonjourProfileMapOutput) ToWirelessControllerBonjourProfileMapOutputWithContext(ctx context.Context) WirelessControllerBonjourProfileMapOutput {
	return o
}

func (o WirelessControllerBonjourProfileMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerBonjourProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerBonjourProfile {
		return vs[0].(map[string]*WirelessControllerBonjourProfile)[vs[1].(string)]
	}).(WirelessControllerBonjourProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerBonjourProfileInput)(nil)).Elem(), &WirelessControllerBonjourProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerBonjourProfileArrayInput)(nil)).Elem(), WirelessControllerBonjourProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerBonjourProfileMapInput)(nil)).Elem(), WirelessControllerBonjourProfileMap{})
	pulumi.RegisterOutputType(WirelessControllerBonjourProfileOutput{})
	pulumi.RegisterOutputType(WirelessControllerBonjourProfileArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerBonjourProfileMapOutput{})
}

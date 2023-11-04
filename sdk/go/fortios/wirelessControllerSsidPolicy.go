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

type WirelessControllerSsidPolicy struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Vdomparam   pulumi.StringPtrOutput `pulumi:"vdomparam"`
	Vlan        pulumi.StringOutput    `pulumi:"vlan"`
}

// NewWirelessControllerSsidPolicy registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerSsidPolicy(ctx *pulumi.Context,
	name string, args *WirelessControllerSsidPolicyArgs, opts ...pulumi.ResourceOption) (*WirelessControllerSsidPolicy, error) {
	if args == nil {
		args = &WirelessControllerSsidPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerSsidPolicy
	err := ctx.RegisterResource("fortios:index/wirelessControllerSsidPolicy:WirelessControllerSsidPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerSsidPolicy gets an existing WirelessControllerSsidPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerSsidPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerSsidPolicyState, opts ...pulumi.ResourceOption) (*WirelessControllerSsidPolicy, error) {
	var resource WirelessControllerSsidPolicy
	err := ctx.ReadResource("fortios:index/wirelessControllerSsidPolicy:WirelessControllerSsidPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerSsidPolicy resources.
type wirelessControllerSsidPolicyState struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Vdomparam   *string `pulumi:"vdomparam"`
	Vlan        *string `pulumi:"vlan"`
}

type WirelessControllerSsidPolicyState struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Vdomparam   pulumi.StringPtrInput
	Vlan        pulumi.StringPtrInput
}

func (WirelessControllerSsidPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerSsidPolicyState)(nil)).Elem()
}

type wirelessControllerSsidPolicyArgs struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Vdomparam   *string `pulumi:"vdomparam"`
	Vlan        *string `pulumi:"vlan"`
}

// The set of arguments for constructing a WirelessControllerSsidPolicy resource.
type WirelessControllerSsidPolicyArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Vdomparam   pulumi.StringPtrInput
	Vlan        pulumi.StringPtrInput
}

func (WirelessControllerSsidPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerSsidPolicyArgs)(nil)).Elem()
}

type WirelessControllerSsidPolicyInput interface {
	pulumi.Input

	ToWirelessControllerSsidPolicyOutput() WirelessControllerSsidPolicyOutput
	ToWirelessControllerSsidPolicyOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyOutput
}

func (*WirelessControllerSsidPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerSsidPolicy)(nil)).Elem()
}

func (i *WirelessControllerSsidPolicy) ToWirelessControllerSsidPolicyOutput() WirelessControllerSsidPolicyOutput {
	return i.ToWirelessControllerSsidPolicyOutputWithContext(context.Background())
}

func (i *WirelessControllerSsidPolicy) ToWirelessControllerSsidPolicyOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSsidPolicyOutput)
}

func (i *WirelessControllerSsidPolicy) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerSsidPolicy] {
	return pulumix.Output[*WirelessControllerSsidPolicy]{
		OutputState: i.ToWirelessControllerSsidPolicyOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerSsidPolicyArrayInput is an input type that accepts WirelessControllerSsidPolicyArray and WirelessControllerSsidPolicyArrayOutput values.
// You can construct a concrete instance of `WirelessControllerSsidPolicyArrayInput` via:
//
//	WirelessControllerSsidPolicyArray{ WirelessControllerSsidPolicyArgs{...} }
type WirelessControllerSsidPolicyArrayInput interface {
	pulumi.Input

	ToWirelessControllerSsidPolicyArrayOutput() WirelessControllerSsidPolicyArrayOutput
	ToWirelessControllerSsidPolicyArrayOutputWithContext(context.Context) WirelessControllerSsidPolicyArrayOutput
}

type WirelessControllerSsidPolicyArray []WirelessControllerSsidPolicyInput

func (WirelessControllerSsidPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerSsidPolicy)(nil)).Elem()
}

func (i WirelessControllerSsidPolicyArray) ToWirelessControllerSsidPolicyArrayOutput() WirelessControllerSsidPolicyArrayOutput {
	return i.ToWirelessControllerSsidPolicyArrayOutputWithContext(context.Background())
}

func (i WirelessControllerSsidPolicyArray) ToWirelessControllerSsidPolicyArrayOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSsidPolicyArrayOutput)
}

func (i WirelessControllerSsidPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerSsidPolicy] {
	return pulumix.Output[[]*WirelessControllerSsidPolicy]{
		OutputState: i.ToWirelessControllerSsidPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerSsidPolicyMapInput is an input type that accepts WirelessControllerSsidPolicyMap and WirelessControllerSsidPolicyMapOutput values.
// You can construct a concrete instance of `WirelessControllerSsidPolicyMapInput` via:
//
//	WirelessControllerSsidPolicyMap{ "key": WirelessControllerSsidPolicyArgs{...} }
type WirelessControllerSsidPolicyMapInput interface {
	pulumi.Input

	ToWirelessControllerSsidPolicyMapOutput() WirelessControllerSsidPolicyMapOutput
	ToWirelessControllerSsidPolicyMapOutputWithContext(context.Context) WirelessControllerSsidPolicyMapOutput
}

type WirelessControllerSsidPolicyMap map[string]WirelessControllerSsidPolicyInput

func (WirelessControllerSsidPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerSsidPolicy)(nil)).Elem()
}

func (i WirelessControllerSsidPolicyMap) ToWirelessControllerSsidPolicyMapOutput() WirelessControllerSsidPolicyMapOutput {
	return i.ToWirelessControllerSsidPolicyMapOutputWithContext(context.Background())
}

func (i WirelessControllerSsidPolicyMap) ToWirelessControllerSsidPolicyMapOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSsidPolicyMapOutput)
}

func (i WirelessControllerSsidPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerSsidPolicy] {
	return pulumix.Output[map[string]*WirelessControllerSsidPolicy]{
		OutputState: i.ToWirelessControllerSsidPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type WirelessControllerSsidPolicyOutput struct{ *pulumi.OutputState }

func (WirelessControllerSsidPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerSsidPolicy)(nil)).Elem()
}

func (o WirelessControllerSsidPolicyOutput) ToWirelessControllerSsidPolicyOutput() WirelessControllerSsidPolicyOutput {
	return o
}

func (o WirelessControllerSsidPolicyOutput) ToWirelessControllerSsidPolicyOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyOutput {
	return o
}

func (o WirelessControllerSsidPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerSsidPolicy] {
	return pulumix.Output[*WirelessControllerSsidPolicy]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerSsidPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerSsidPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerSsidPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerSsidPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WirelessControllerSsidPolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerSsidPolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerSsidPolicyOutput) Vlan() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerSsidPolicy) pulumi.StringOutput { return v.Vlan }).(pulumi.StringOutput)
}

type WirelessControllerSsidPolicyArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerSsidPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerSsidPolicy)(nil)).Elem()
}

func (o WirelessControllerSsidPolicyArrayOutput) ToWirelessControllerSsidPolicyArrayOutput() WirelessControllerSsidPolicyArrayOutput {
	return o
}

func (o WirelessControllerSsidPolicyArrayOutput) ToWirelessControllerSsidPolicyArrayOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyArrayOutput {
	return o
}

func (o WirelessControllerSsidPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerSsidPolicy] {
	return pulumix.Output[[]*WirelessControllerSsidPolicy]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerSsidPolicyArrayOutput) Index(i pulumi.IntInput) WirelessControllerSsidPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerSsidPolicy {
		return vs[0].([]*WirelessControllerSsidPolicy)[vs[1].(int)]
	}).(WirelessControllerSsidPolicyOutput)
}

type WirelessControllerSsidPolicyMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerSsidPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerSsidPolicy)(nil)).Elem()
}

func (o WirelessControllerSsidPolicyMapOutput) ToWirelessControllerSsidPolicyMapOutput() WirelessControllerSsidPolicyMapOutput {
	return o
}

func (o WirelessControllerSsidPolicyMapOutput) ToWirelessControllerSsidPolicyMapOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyMapOutput {
	return o
}

func (o WirelessControllerSsidPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerSsidPolicy] {
	return pulumix.Output[map[string]*WirelessControllerSsidPolicy]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerSsidPolicyMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerSsidPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerSsidPolicy {
		return vs[0].(map[string]*WirelessControllerSsidPolicy)[vs[1].(string)]
	}).(WirelessControllerSsidPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerSsidPolicyInput)(nil)).Elem(), &WirelessControllerSsidPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerSsidPolicyArrayInput)(nil)).Elem(), WirelessControllerSsidPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerSsidPolicyMapInput)(nil)).Elem(), WirelessControllerSsidPolicyMap{})
	pulumi.RegisterOutputType(WirelessControllerSsidPolicyOutput{})
	pulumi.RegisterOutputType(WirelessControllerSsidPolicyArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerSsidPolicyMapOutput{})
}

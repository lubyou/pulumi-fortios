// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnIpsecConcentrator struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput                `pulumi:"dynamicSortSubtable"`
	Fosid               pulumi.IntOutput                      `pulumi:"fosid"`
	Members             VpnIpsecConcentratorMemberArrayOutput `pulumi:"members"`
	Name                pulumi.StringOutput                   `pulumi:"name"`
	SrcCheck            pulumi.StringOutput                   `pulumi:"srcCheck"`
	Vdomparam           pulumi.StringPtrOutput                `pulumi:"vdomparam"`
}

// NewVpnIpsecConcentrator registers a new resource with the given unique name, arguments, and options.
func NewVpnIpsecConcentrator(ctx *pulumi.Context,
	name string, args *VpnIpsecConcentratorArgs, opts ...pulumi.ResourceOption) (*VpnIpsecConcentrator, error) {
	if args == nil {
		args = &VpnIpsecConcentratorArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource VpnIpsecConcentrator
	err := ctx.RegisterResource("fortios:index/vpnIpsecConcentrator:VpnIpsecConcentrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnIpsecConcentrator gets an existing VpnIpsecConcentrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnIpsecConcentrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnIpsecConcentratorState, opts ...pulumi.ResourceOption) (*VpnIpsecConcentrator, error) {
	var resource VpnIpsecConcentrator
	err := ctx.ReadResource("fortios:index/vpnIpsecConcentrator:VpnIpsecConcentrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnIpsecConcentrator resources.
type vpnIpsecConcentratorState struct {
	DynamicSortSubtable *string                      `pulumi:"dynamicSortSubtable"`
	Fosid               *int                         `pulumi:"fosid"`
	Members             []VpnIpsecConcentratorMember `pulumi:"members"`
	Name                *string                      `pulumi:"name"`
	SrcCheck            *string                      `pulumi:"srcCheck"`
	Vdomparam           *string                      `pulumi:"vdomparam"`
}

type VpnIpsecConcentratorState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	Fosid               pulumi.IntPtrInput
	Members             VpnIpsecConcentratorMemberArrayInput
	Name                pulumi.StringPtrInput
	SrcCheck            pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (VpnIpsecConcentratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnIpsecConcentratorState)(nil)).Elem()
}

type vpnIpsecConcentratorArgs struct {
	DynamicSortSubtable *string                      `pulumi:"dynamicSortSubtable"`
	Fosid               *int                         `pulumi:"fosid"`
	Members             []VpnIpsecConcentratorMember `pulumi:"members"`
	Name                *string                      `pulumi:"name"`
	SrcCheck            *string                      `pulumi:"srcCheck"`
	Vdomparam           *string                      `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnIpsecConcentrator resource.
type VpnIpsecConcentratorArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	Fosid               pulumi.IntPtrInput
	Members             VpnIpsecConcentratorMemberArrayInput
	Name                pulumi.StringPtrInput
	SrcCheck            pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (VpnIpsecConcentratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnIpsecConcentratorArgs)(nil)).Elem()
}

type VpnIpsecConcentratorInput interface {
	pulumi.Input

	ToVpnIpsecConcentratorOutput() VpnIpsecConcentratorOutput
	ToVpnIpsecConcentratorOutputWithContext(ctx context.Context) VpnIpsecConcentratorOutput
}

func (*VpnIpsecConcentrator) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnIpsecConcentrator)(nil)).Elem()
}

func (i *VpnIpsecConcentrator) ToVpnIpsecConcentratorOutput() VpnIpsecConcentratorOutput {
	return i.ToVpnIpsecConcentratorOutputWithContext(context.Background())
}

func (i *VpnIpsecConcentrator) ToVpnIpsecConcentratorOutputWithContext(ctx context.Context) VpnIpsecConcentratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecConcentratorOutput)
}

// VpnIpsecConcentratorArrayInput is an input type that accepts VpnIpsecConcentratorArray and VpnIpsecConcentratorArrayOutput values.
// You can construct a concrete instance of `VpnIpsecConcentratorArrayInput` via:
//
//	VpnIpsecConcentratorArray{ VpnIpsecConcentratorArgs{...} }
type VpnIpsecConcentratorArrayInput interface {
	pulumi.Input

	ToVpnIpsecConcentratorArrayOutput() VpnIpsecConcentratorArrayOutput
	ToVpnIpsecConcentratorArrayOutputWithContext(context.Context) VpnIpsecConcentratorArrayOutput
}

type VpnIpsecConcentratorArray []VpnIpsecConcentratorInput

func (VpnIpsecConcentratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnIpsecConcentrator)(nil)).Elem()
}

func (i VpnIpsecConcentratorArray) ToVpnIpsecConcentratorArrayOutput() VpnIpsecConcentratorArrayOutput {
	return i.ToVpnIpsecConcentratorArrayOutputWithContext(context.Background())
}

func (i VpnIpsecConcentratorArray) ToVpnIpsecConcentratorArrayOutputWithContext(ctx context.Context) VpnIpsecConcentratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecConcentratorArrayOutput)
}

// VpnIpsecConcentratorMapInput is an input type that accepts VpnIpsecConcentratorMap and VpnIpsecConcentratorMapOutput values.
// You can construct a concrete instance of `VpnIpsecConcentratorMapInput` via:
//
//	VpnIpsecConcentratorMap{ "key": VpnIpsecConcentratorArgs{...} }
type VpnIpsecConcentratorMapInput interface {
	pulumi.Input

	ToVpnIpsecConcentratorMapOutput() VpnIpsecConcentratorMapOutput
	ToVpnIpsecConcentratorMapOutputWithContext(context.Context) VpnIpsecConcentratorMapOutput
}

type VpnIpsecConcentratorMap map[string]VpnIpsecConcentratorInput

func (VpnIpsecConcentratorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnIpsecConcentrator)(nil)).Elem()
}

func (i VpnIpsecConcentratorMap) ToVpnIpsecConcentratorMapOutput() VpnIpsecConcentratorMapOutput {
	return i.ToVpnIpsecConcentratorMapOutputWithContext(context.Background())
}

func (i VpnIpsecConcentratorMap) ToVpnIpsecConcentratorMapOutputWithContext(ctx context.Context) VpnIpsecConcentratorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecConcentratorMapOutput)
}

type VpnIpsecConcentratorOutput struct{ *pulumi.OutputState }

func (VpnIpsecConcentratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnIpsecConcentrator)(nil)).Elem()
}

func (o VpnIpsecConcentratorOutput) ToVpnIpsecConcentratorOutput() VpnIpsecConcentratorOutput {
	return o
}

func (o VpnIpsecConcentratorOutput) ToVpnIpsecConcentratorOutputWithContext(ctx context.Context) VpnIpsecConcentratorOutput {
	return o
}

func (o VpnIpsecConcentratorOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnIpsecConcentrator) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o VpnIpsecConcentratorOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnIpsecConcentrator) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o VpnIpsecConcentratorOutput) Members() VpnIpsecConcentratorMemberArrayOutput {
	return o.ApplyT(func(v *VpnIpsecConcentrator) VpnIpsecConcentratorMemberArrayOutput { return v.Members }).(VpnIpsecConcentratorMemberArrayOutput)
}

func (o VpnIpsecConcentratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecConcentrator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnIpsecConcentratorOutput) SrcCheck() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecConcentrator) pulumi.StringOutput { return v.SrcCheck }).(pulumi.StringOutput)
}

func (o VpnIpsecConcentratorOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnIpsecConcentrator) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpnIpsecConcentratorArrayOutput struct{ *pulumi.OutputState }

func (VpnIpsecConcentratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnIpsecConcentrator)(nil)).Elem()
}

func (o VpnIpsecConcentratorArrayOutput) ToVpnIpsecConcentratorArrayOutput() VpnIpsecConcentratorArrayOutput {
	return o
}

func (o VpnIpsecConcentratorArrayOutput) ToVpnIpsecConcentratorArrayOutputWithContext(ctx context.Context) VpnIpsecConcentratorArrayOutput {
	return o
}

func (o VpnIpsecConcentratorArrayOutput) Index(i pulumi.IntInput) VpnIpsecConcentratorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnIpsecConcentrator {
		return vs[0].([]*VpnIpsecConcentrator)[vs[1].(int)]
	}).(VpnIpsecConcentratorOutput)
}

type VpnIpsecConcentratorMapOutput struct{ *pulumi.OutputState }

func (VpnIpsecConcentratorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnIpsecConcentrator)(nil)).Elem()
}

func (o VpnIpsecConcentratorMapOutput) ToVpnIpsecConcentratorMapOutput() VpnIpsecConcentratorMapOutput {
	return o
}

func (o VpnIpsecConcentratorMapOutput) ToVpnIpsecConcentratorMapOutputWithContext(ctx context.Context) VpnIpsecConcentratorMapOutput {
	return o
}

func (o VpnIpsecConcentratorMapOutput) MapIndex(k pulumi.StringInput) VpnIpsecConcentratorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnIpsecConcentrator {
		return vs[0].(map[string]*VpnIpsecConcentrator)[vs[1].(string)]
	}).(VpnIpsecConcentratorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecConcentratorInput)(nil)).Elem(), &VpnIpsecConcentrator{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecConcentratorArrayInput)(nil)).Elem(), VpnIpsecConcentratorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecConcentratorMapInput)(nil)).Elem(), VpnIpsecConcentratorMap{})
	pulumi.RegisterOutputType(VpnIpsecConcentratorOutput{})
	pulumi.RegisterOutputType(VpnIpsecConcentratorArrayOutput{})
	pulumi.RegisterOutputType(VpnIpsecConcentratorMapOutput{})
}

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

type VpnIpsecFec struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput        `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput        `pulumi:"getAllTables"`
	Mappings            VpnIpsecFecMappingArrayOutput `pulumi:"mappings"`
	Name                pulumi.StringOutput           `pulumi:"name"`
	Vdomparam           pulumi.StringPtrOutput        `pulumi:"vdomparam"`
}

// NewVpnIpsecFec registers a new resource with the given unique name, arguments, and options.
func NewVpnIpsecFec(ctx *pulumi.Context,
	name string, args *VpnIpsecFecArgs, opts ...pulumi.ResourceOption) (*VpnIpsecFec, error) {
	if args == nil {
		args = &VpnIpsecFecArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpnIpsecFec
	err := ctx.RegisterResource("fortios:index/vpnIpsecFec:VpnIpsecFec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnIpsecFec gets an existing VpnIpsecFec resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnIpsecFec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnIpsecFecState, opts ...pulumi.ResourceOption) (*VpnIpsecFec, error) {
	var resource VpnIpsecFec
	err := ctx.ReadResource("fortios:index/vpnIpsecFec:VpnIpsecFec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnIpsecFec resources.
type vpnIpsecFecState struct {
	DynamicSortSubtable *string              `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string              `pulumi:"getAllTables"`
	Mappings            []VpnIpsecFecMapping `pulumi:"mappings"`
	Name                *string              `pulumi:"name"`
	Vdomparam           *string              `pulumi:"vdomparam"`
}

type VpnIpsecFecState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Mappings            VpnIpsecFecMappingArrayInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (VpnIpsecFecState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnIpsecFecState)(nil)).Elem()
}

type vpnIpsecFecArgs struct {
	DynamicSortSubtable *string              `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string              `pulumi:"getAllTables"`
	Mappings            []VpnIpsecFecMapping `pulumi:"mappings"`
	Name                *string              `pulumi:"name"`
	Vdomparam           *string              `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnIpsecFec resource.
type VpnIpsecFecArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Mappings            VpnIpsecFecMappingArrayInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (VpnIpsecFecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnIpsecFecArgs)(nil)).Elem()
}

type VpnIpsecFecInput interface {
	pulumi.Input

	ToVpnIpsecFecOutput() VpnIpsecFecOutput
	ToVpnIpsecFecOutputWithContext(ctx context.Context) VpnIpsecFecOutput
}

func (*VpnIpsecFec) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnIpsecFec)(nil)).Elem()
}

func (i *VpnIpsecFec) ToVpnIpsecFecOutput() VpnIpsecFecOutput {
	return i.ToVpnIpsecFecOutputWithContext(context.Background())
}

func (i *VpnIpsecFec) ToVpnIpsecFecOutputWithContext(ctx context.Context) VpnIpsecFecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecFecOutput)
}

func (i *VpnIpsecFec) ToOutput(ctx context.Context) pulumix.Output[*VpnIpsecFec] {
	return pulumix.Output[*VpnIpsecFec]{
		OutputState: i.ToVpnIpsecFecOutputWithContext(ctx).OutputState,
	}
}

// VpnIpsecFecArrayInput is an input type that accepts VpnIpsecFecArray and VpnIpsecFecArrayOutput values.
// You can construct a concrete instance of `VpnIpsecFecArrayInput` via:
//
//	VpnIpsecFecArray{ VpnIpsecFecArgs{...} }
type VpnIpsecFecArrayInput interface {
	pulumi.Input

	ToVpnIpsecFecArrayOutput() VpnIpsecFecArrayOutput
	ToVpnIpsecFecArrayOutputWithContext(context.Context) VpnIpsecFecArrayOutput
}

type VpnIpsecFecArray []VpnIpsecFecInput

func (VpnIpsecFecArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnIpsecFec)(nil)).Elem()
}

func (i VpnIpsecFecArray) ToVpnIpsecFecArrayOutput() VpnIpsecFecArrayOutput {
	return i.ToVpnIpsecFecArrayOutputWithContext(context.Background())
}

func (i VpnIpsecFecArray) ToVpnIpsecFecArrayOutputWithContext(ctx context.Context) VpnIpsecFecArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecFecArrayOutput)
}

func (i VpnIpsecFecArray) ToOutput(ctx context.Context) pulumix.Output[[]*VpnIpsecFec] {
	return pulumix.Output[[]*VpnIpsecFec]{
		OutputState: i.ToVpnIpsecFecArrayOutputWithContext(ctx).OutputState,
	}
}

// VpnIpsecFecMapInput is an input type that accepts VpnIpsecFecMap and VpnIpsecFecMapOutput values.
// You can construct a concrete instance of `VpnIpsecFecMapInput` via:
//
//	VpnIpsecFecMap{ "key": VpnIpsecFecArgs{...} }
type VpnIpsecFecMapInput interface {
	pulumi.Input

	ToVpnIpsecFecMapOutput() VpnIpsecFecMapOutput
	ToVpnIpsecFecMapOutputWithContext(context.Context) VpnIpsecFecMapOutput
}

type VpnIpsecFecMap map[string]VpnIpsecFecInput

func (VpnIpsecFecMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnIpsecFec)(nil)).Elem()
}

func (i VpnIpsecFecMap) ToVpnIpsecFecMapOutput() VpnIpsecFecMapOutput {
	return i.ToVpnIpsecFecMapOutputWithContext(context.Background())
}

func (i VpnIpsecFecMap) ToVpnIpsecFecMapOutputWithContext(ctx context.Context) VpnIpsecFecMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecFecMapOutput)
}

func (i VpnIpsecFecMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpnIpsecFec] {
	return pulumix.Output[map[string]*VpnIpsecFec]{
		OutputState: i.ToVpnIpsecFecMapOutputWithContext(ctx).OutputState,
	}
}

type VpnIpsecFecOutput struct{ *pulumi.OutputState }

func (VpnIpsecFecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnIpsecFec)(nil)).Elem()
}

func (o VpnIpsecFecOutput) ToVpnIpsecFecOutput() VpnIpsecFecOutput {
	return o
}

func (o VpnIpsecFecOutput) ToVpnIpsecFecOutputWithContext(ctx context.Context) VpnIpsecFecOutput {
	return o
}

func (o VpnIpsecFecOutput) ToOutput(ctx context.Context) pulumix.Output[*VpnIpsecFec] {
	return pulumix.Output[*VpnIpsecFec]{
		OutputState: o.OutputState,
	}
}

func (o VpnIpsecFecOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnIpsecFec) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o VpnIpsecFecOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnIpsecFec) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o VpnIpsecFecOutput) Mappings() VpnIpsecFecMappingArrayOutput {
	return o.ApplyT(func(v *VpnIpsecFec) VpnIpsecFecMappingArrayOutput { return v.Mappings }).(VpnIpsecFecMappingArrayOutput)
}

func (o VpnIpsecFecOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecFec) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnIpsecFecOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnIpsecFec) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpnIpsecFecArrayOutput struct{ *pulumi.OutputState }

func (VpnIpsecFecArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnIpsecFec)(nil)).Elem()
}

func (o VpnIpsecFecArrayOutput) ToVpnIpsecFecArrayOutput() VpnIpsecFecArrayOutput {
	return o
}

func (o VpnIpsecFecArrayOutput) ToVpnIpsecFecArrayOutputWithContext(ctx context.Context) VpnIpsecFecArrayOutput {
	return o
}

func (o VpnIpsecFecArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VpnIpsecFec] {
	return pulumix.Output[[]*VpnIpsecFec]{
		OutputState: o.OutputState,
	}
}

func (o VpnIpsecFecArrayOutput) Index(i pulumi.IntInput) VpnIpsecFecOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnIpsecFec {
		return vs[0].([]*VpnIpsecFec)[vs[1].(int)]
	}).(VpnIpsecFecOutput)
}

type VpnIpsecFecMapOutput struct{ *pulumi.OutputState }

func (VpnIpsecFecMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnIpsecFec)(nil)).Elem()
}

func (o VpnIpsecFecMapOutput) ToVpnIpsecFecMapOutput() VpnIpsecFecMapOutput {
	return o
}

func (o VpnIpsecFecMapOutput) ToVpnIpsecFecMapOutputWithContext(ctx context.Context) VpnIpsecFecMapOutput {
	return o
}

func (o VpnIpsecFecMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpnIpsecFec] {
	return pulumix.Output[map[string]*VpnIpsecFec]{
		OutputState: o.OutputState,
	}
}

func (o VpnIpsecFecMapOutput) MapIndex(k pulumi.StringInput) VpnIpsecFecOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnIpsecFec {
		return vs[0].(map[string]*VpnIpsecFec)[vs[1].(string)]
	}).(VpnIpsecFecOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecFecInput)(nil)).Elem(), &VpnIpsecFec{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecFecArrayInput)(nil)).Elem(), VpnIpsecFecArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecFecMapInput)(nil)).Elem(), VpnIpsecFecMap{})
	pulumi.RegisterOutputType(VpnIpsecFecOutput{})
	pulumi.RegisterOutputType(VpnIpsecFecArrayOutput{})
	pulumi.RegisterOutputType(VpnIpsecFecMapOutput{})
}

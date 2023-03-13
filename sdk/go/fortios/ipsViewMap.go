// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IpsViewMap struct {
	pulumi.CustomResourceState

	Fosid      pulumi.IntOutput       `pulumi:"fosid"`
	IdPolicyId pulumi.IntOutput       `pulumi:"idPolicyId"`
	PolicyId   pulumi.IntOutput       `pulumi:"policyId"`
	VdomId     pulumi.IntOutput       `pulumi:"vdomId"`
	Vdomparam  pulumi.StringPtrOutput `pulumi:"vdomparam"`
	Which      pulumi.StringOutput    `pulumi:"which"`
}

// NewIpsViewMap registers a new resource with the given unique name, arguments, and options.
func NewIpsViewMap(ctx *pulumi.Context,
	name string, args *IpsViewMapArgs, opts ...pulumi.ResourceOption) (*IpsViewMap, error) {
	if args == nil {
		args = &IpsViewMapArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource IpsViewMap
	err := ctx.RegisterResource("fortios:index/ipsViewMap:IpsViewMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsViewMap gets an existing IpsViewMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsViewMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsViewMapState, opts ...pulumi.ResourceOption) (*IpsViewMap, error) {
	var resource IpsViewMap
	err := ctx.ReadResource("fortios:index/ipsViewMap:IpsViewMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsViewMap resources.
type ipsViewMapState struct {
	Fosid      *int    `pulumi:"fosid"`
	IdPolicyId *int    `pulumi:"idPolicyId"`
	PolicyId   *int    `pulumi:"policyId"`
	VdomId     *int    `pulumi:"vdomId"`
	Vdomparam  *string `pulumi:"vdomparam"`
	Which      *string `pulumi:"which"`
}

type IpsViewMapState struct {
	Fosid      pulumi.IntPtrInput
	IdPolicyId pulumi.IntPtrInput
	PolicyId   pulumi.IntPtrInput
	VdomId     pulumi.IntPtrInput
	Vdomparam  pulumi.StringPtrInput
	Which      pulumi.StringPtrInput
}

func (IpsViewMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsViewMapState)(nil)).Elem()
}

type ipsViewMapArgs struct {
	Fosid      *int    `pulumi:"fosid"`
	IdPolicyId *int    `pulumi:"idPolicyId"`
	PolicyId   *int    `pulumi:"policyId"`
	VdomId     *int    `pulumi:"vdomId"`
	Vdomparam  *string `pulumi:"vdomparam"`
	Which      *string `pulumi:"which"`
}

// The set of arguments for constructing a IpsViewMap resource.
type IpsViewMapArgs struct {
	Fosid      pulumi.IntPtrInput
	IdPolicyId pulumi.IntPtrInput
	PolicyId   pulumi.IntPtrInput
	VdomId     pulumi.IntPtrInput
	Vdomparam  pulumi.StringPtrInput
	Which      pulumi.StringPtrInput
}

func (IpsViewMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsViewMapArgs)(nil)).Elem()
}

type IpsViewMapInput interface {
	pulumi.Input

	ToIpsViewMapOutput() IpsViewMapOutput
	ToIpsViewMapOutputWithContext(ctx context.Context) IpsViewMapOutput
}

func (*IpsViewMap) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsViewMap)(nil)).Elem()
}

func (i *IpsViewMap) ToIpsViewMapOutput() IpsViewMapOutput {
	return i.ToIpsViewMapOutputWithContext(context.Background())
}

func (i *IpsViewMap) ToIpsViewMapOutputWithContext(ctx context.Context) IpsViewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsViewMapOutput)
}

// IpsViewMapArrayInput is an input type that accepts IpsViewMapArray and IpsViewMapArrayOutput values.
// You can construct a concrete instance of `IpsViewMapArrayInput` via:
//
//	IpsViewMapArray{ IpsViewMapArgs{...} }
type IpsViewMapArrayInput interface {
	pulumi.Input

	ToIpsViewMapArrayOutput() IpsViewMapArrayOutput
	ToIpsViewMapArrayOutputWithContext(context.Context) IpsViewMapArrayOutput
}

type IpsViewMapArray []IpsViewMapInput

func (IpsViewMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsViewMap)(nil)).Elem()
}

func (i IpsViewMapArray) ToIpsViewMapArrayOutput() IpsViewMapArrayOutput {
	return i.ToIpsViewMapArrayOutputWithContext(context.Background())
}

func (i IpsViewMapArray) ToIpsViewMapArrayOutputWithContext(ctx context.Context) IpsViewMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsViewMapArrayOutput)
}

// IpsViewMapMapInput is an input type that accepts IpsViewMapMap and IpsViewMapMapOutput values.
// You can construct a concrete instance of `IpsViewMapMapInput` via:
//
//	IpsViewMapMap{ "key": IpsViewMapArgs{...} }
type IpsViewMapMapInput interface {
	pulumi.Input

	ToIpsViewMapMapOutput() IpsViewMapMapOutput
	ToIpsViewMapMapOutputWithContext(context.Context) IpsViewMapMapOutput
}

type IpsViewMapMap map[string]IpsViewMapInput

func (IpsViewMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsViewMap)(nil)).Elem()
}

func (i IpsViewMapMap) ToIpsViewMapMapOutput() IpsViewMapMapOutput {
	return i.ToIpsViewMapMapOutputWithContext(context.Background())
}

func (i IpsViewMapMap) ToIpsViewMapMapOutputWithContext(ctx context.Context) IpsViewMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsViewMapMapOutput)
}

type IpsViewMapOutput struct{ *pulumi.OutputState }

func (IpsViewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsViewMap)(nil)).Elem()
}

func (o IpsViewMapOutput) ToIpsViewMapOutput() IpsViewMapOutput {
	return o
}

func (o IpsViewMapOutput) ToIpsViewMapOutputWithContext(ctx context.Context) IpsViewMapOutput {
	return o
}

func (o IpsViewMapOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsViewMap) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o IpsViewMapOutput) IdPolicyId() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsViewMap) pulumi.IntOutput { return v.IdPolicyId }).(pulumi.IntOutput)
}

func (o IpsViewMapOutput) PolicyId() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsViewMap) pulumi.IntOutput { return v.PolicyId }).(pulumi.IntOutput)
}

func (o IpsViewMapOutput) VdomId() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsViewMap) pulumi.IntOutput { return v.VdomId }).(pulumi.IntOutput)
}

func (o IpsViewMapOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpsViewMap) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o IpsViewMapOutput) Which() pulumi.StringOutput {
	return o.ApplyT(func(v *IpsViewMap) pulumi.StringOutput { return v.Which }).(pulumi.StringOutput)
}

type IpsViewMapArrayOutput struct{ *pulumi.OutputState }

func (IpsViewMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsViewMap)(nil)).Elem()
}

func (o IpsViewMapArrayOutput) ToIpsViewMapArrayOutput() IpsViewMapArrayOutput {
	return o
}

func (o IpsViewMapArrayOutput) ToIpsViewMapArrayOutputWithContext(ctx context.Context) IpsViewMapArrayOutput {
	return o
}

func (o IpsViewMapArrayOutput) Index(i pulumi.IntInput) IpsViewMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsViewMap {
		return vs[0].([]*IpsViewMap)[vs[1].(int)]
	}).(IpsViewMapOutput)
}

type IpsViewMapMapOutput struct{ *pulumi.OutputState }

func (IpsViewMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsViewMap)(nil)).Elem()
}

func (o IpsViewMapMapOutput) ToIpsViewMapMapOutput() IpsViewMapMapOutput {
	return o
}

func (o IpsViewMapMapOutput) ToIpsViewMapMapOutputWithContext(ctx context.Context) IpsViewMapMapOutput {
	return o
}

func (o IpsViewMapMapOutput) MapIndex(k pulumi.StringInput) IpsViewMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsViewMap {
		return vs[0].(map[string]*IpsViewMap)[vs[1].(string)]
	}).(IpsViewMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsViewMapInput)(nil)).Elem(), &IpsViewMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsViewMapArrayInput)(nil)).Elem(), IpsViewMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsViewMapMapInput)(nil)).Elem(), IpsViewMapMap{})
	pulumi.RegisterOutputType(IpsViewMapOutput{})
	pulumi.RegisterOutputType(IpsViewMapArrayOutput{})
	pulumi.RegisterOutputType(IpsViewMapMapOutput{})
}

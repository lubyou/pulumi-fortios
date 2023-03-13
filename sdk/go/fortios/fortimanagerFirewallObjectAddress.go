// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FortimanagerFirewallObjectAddress struct {
	pulumi.CustomResourceState

	Adom           pulumi.StringPtrOutput `pulumi:"adom"`
	AllowRouting   pulumi.StringPtrOutput `pulumi:"allowRouting"`
	AssociatedIntf pulumi.StringPtrOutput `pulumi:"associatedIntf"`
	Comment        pulumi.StringPtrOutput `pulumi:"comment"`
	EndIp          pulumi.StringPtrOutput `pulumi:"endIp"`
	Fqdn           pulumi.StringPtrOutput `pulumi:"fqdn"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	StartIp        pulumi.StringPtrOutput `pulumi:"startIp"`
	Subnet         pulumi.StringPtrOutput `pulumi:"subnet"`
	Type           pulumi.StringPtrOutput `pulumi:"type"`
}

// NewFortimanagerFirewallObjectAddress registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerFirewallObjectAddress(ctx *pulumi.Context,
	name string, args *FortimanagerFirewallObjectAddressArgs, opts ...pulumi.ResourceOption) (*FortimanagerFirewallObjectAddress, error) {
	if args == nil {
		args = &FortimanagerFirewallObjectAddressArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FortimanagerFirewallObjectAddress
	err := ctx.RegisterResource("fortios:index/fortimanagerFirewallObjectAddress:FortimanagerFirewallObjectAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerFirewallObjectAddress gets an existing FortimanagerFirewallObjectAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerFirewallObjectAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerFirewallObjectAddressState, opts ...pulumi.ResourceOption) (*FortimanagerFirewallObjectAddress, error) {
	var resource FortimanagerFirewallObjectAddress
	err := ctx.ReadResource("fortios:index/fortimanagerFirewallObjectAddress:FortimanagerFirewallObjectAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerFirewallObjectAddress resources.
type fortimanagerFirewallObjectAddressState struct {
	Adom           *string `pulumi:"adom"`
	AllowRouting   *string `pulumi:"allowRouting"`
	AssociatedIntf *string `pulumi:"associatedIntf"`
	Comment        *string `pulumi:"comment"`
	EndIp          *string `pulumi:"endIp"`
	Fqdn           *string `pulumi:"fqdn"`
	Name           *string `pulumi:"name"`
	StartIp        *string `pulumi:"startIp"`
	Subnet         *string `pulumi:"subnet"`
	Type           *string `pulumi:"type"`
}

type FortimanagerFirewallObjectAddressState struct {
	Adom           pulumi.StringPtrInput
	AllowRouting   pulumi.StringPtrInput
	AssociatedIntf pulumi.StringPtrInput
	Comment        pulumi.StringPtrInput
	EndIp          pulumi.StringPtrInput
	Fqdn           pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	StartIp        pulumi.StringPtrInput
	Subnet         pulumi.StringPtrInput
	Type           pulumi.StringPtrInput
}

func (FortimanagerFirewallObjectAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerFirewallObjectAddressState)(nil)).Elem()
}

type fortimanagerFirewallObjectAddressArgs struct {
	Adom           *string `pulumi:"adom"`
	AllowRouting   *string `pulumi:"allowRouting"`
	AssociatedIntf *string `pulumi:"associatedIntf"`
	Comment        *string `pulumi:"comment"`
	EndIp          *string `pulumi:"endIp"`
	Fqdn           *string `pulumi:"fqdn"`
	Name           *string `pulumi:"name"`
	StartIp        *string `pulumi:"startIp"`
	Subnet         *string `pulumi:"subnet"`
	Type           *string `pulumi:"type"`
}

// The set of arguments for constructing a FortimanagerFirewallObjectAddress resource.
type FortimanagerFirewallObjectAddressArgs struct {
	Adom           pulumi.StringPtrInput
	AllowRouting   pulumi.StringPtrInput
	AssociatedIntf pulumi.StringPtrInput
	Comment        pulumi.StringPtrInput
	EndIp          pulumi.StringPtrInput
	Fqdn           pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	StartIp        pulumi.StringPtrInput
	Subnet         pulumi.StringPtrInput
	Type           pulumi.StringPtrInput
}

func (FortimanagerFirewallObjectAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerFirewallObjectAddressArgs)(nil)).Elem()
}

type FortimanagerFirewallObjectAddressInput interface {
	pulumi.Input

	ToFortimanagerFirewallObjectAddressOutput() FortimanagerFirewallObjectAddressOutput
	ToFortimanagerFirewallObjectAddressOutputWithContext(ctx context.Context) FortimanagerFirewallObjectAddressOutput
}

func (*FortimanagerFirewallObjectAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerFirewallObjectAddress)(nil)).Elem()
}

func (i *FortimanagerFirewallObjectAddress) ToFortimanagerFirewallObjectAddressOutput() FortimanagerFirewallObjectAddressOutput {
	return i.ToFortimanagerFirewallObjectAddressOutputWithContext(context.Background())
}

func (i *FortimanagerFirewallObjectAddress) ToFortimanagerFirewallObjectAddressOutputWithContext(ctx context.Context) FortimanagerFirewallObjectAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectAddressOutput)
}

// FortimanagerFirewallObjectAddressArrayInput is an input type that accepts FortimanagerFirewallObjectAddressArray and FortimanagerFirewallObjectAddressArrayOutput values.
// You can construct a concrete instance of `FortimanagerFirewallObjectAddressArrayInput` via:
//
//	FortimanagerFirewallObjectAddressArray{ FortimanagerFirewallObjectAddressArgs{...} }
type FortimanagerFirewallObjectAddressArrayInput interface {
	pulumi.Input

	ToFortimanagerFirewallObjectAddressArrayOutput() FortimanagerFirewallObjectAddressArrayOutput
	ToFortimanagerFirewallObjectAddressArrayOutputWithContext(context.Context) FortimanagerFirewallObjectAddressArrayOutput
}

type FortimanagerFirewallObjectAddressArray []FortimanagerFirewallObjectAddressInput

func (FortimanagerFirewallObjectAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerFirewallObjectAddress)(nil)).Elem()
}

func (i FortimanagerFirewallObjectAddressArray) ToFortimanagerFirewallObjectAddressArrayOutput() FortimanagerFirewallObjectAddressArrayOutput {
	return i.ToFortimanagerFirewallObjectAddressArrayOutputWithContext(context.Background())
}

func (i FortimanagerFirewallObjectAddressArray) ToFortimanagerFirewallObjectAddressArrayOutputWithContext(ctx context.Context) FortimanagerFirewallObjectAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectAddressArrayOutput)
}

// FortimanagerFirewallObjectAddressMapInput is an input type that accepts FortimanagerFirewallObjectAddressMap and FortimanagerFirewallObjectAddressMapOutput values.
// You can construct a concrete instance of `FortimanagerFirewallObjectAddressMapInput` via:
//
//	FortimanagerFirewallObjectAddressMap{ "key": FortimanagerFirewallObjectAddressArgs{...} }
type FortimanagerFirewallObjectAddressMapInput interface {
	pulumi.Input

	ToFortimanagerFirewallObjectAddressMapOutput() FortimanagerFirewallObjectAddressMapOutput
	ToFortimanagerFirewallObjectAddressMapOutputWithContext(context.Context) FortimanagerFirewallObjectAddressMapOutput
}

type FortimanagerFirewallObjectAddressMap map[string]FortimanagerFirewallObjectAddressInput

func (FortimanagerFirewallObjectAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerFirewallObjectAddress)(nil)).Elem()
}

func (i FortimanagerFirewallObjectAddressMap) ToFortimanagerFirewallObjectAddressMapOutput() FortimanagerFirewallObjectAddressMapOutput {
	return i.ToFortimanagerFirewallObjectAddressMapOutputWithContext(context.Background())
}

func (i FortimanagerFirewallObjectAddressMap) ToFortimanagerFirewallObjectAddressMapOutputWithContext(ctx context.Context) FortimanagerFirewallObjectAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectAddressMapOutput)
}

type FortimanagerFirewallObjectAddressOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallObjectAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerFirewallObjectAddress)(nil)).Elem()
}

func (o FortimanagerFirewallObjectAddressOutput) ToFortimanagerFirewallObjectAddressOutput() FortimanagerFirewallObjectAddressOutput {
	return o
}

func (o FortimanagerFirewallObjectAddressOutput) ToFortimanagerFirewallObjectAddressOutputWithContext(ctx context.Context) FortimanagerFirewallObjectAddressOutput {
	return o
}

func (o FortimanagerFirewallObjectAddressOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerFirewallObjectAddress) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

func (o FortimanagerFirewallObjectAddressOutput) AllowRouting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerFirewallObjectAddress) pulumi.StringPtrOutput { return v.AllowRouting }).(pulumi.StringPtrOutput)
}

func (o FortimanagerFirewallObjectAddressOutput) AssociatedIntf() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerFirewallObjectAddress) pulumi.StringPtrOutput { return v.AssociatedIntf }).(pulumi.StringPtrOutput)
}

func (o FortimanagerFirewallObjectAddressOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerFirewallObjectAddress) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FortimanagerFirewallObjectAddressOutput) EndIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerFirewallObjectAddress) pulumi.StringPtrOutput { return v.EndIp }).(pulumi.StringPtrOutput)
}

func (o FortimanagerFirewallObjectAddressOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerFirewallObjectAddress) pulumi.StringPtrOutput { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o FortimanagerFirewallObjectAddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerFirewallObjectAddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FortimanagerFirewallObjectAddressOutput) StartIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerFirewallObjectAddress) pulumi.StringPtrOutput { return v.StartIp }).(pulumi.StringPtrOutput)
}

func (o FortimanagerFirewallObjectAddressOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerFirewallObjectAddress) pulumi.StringPtrOutput { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o FortimanagerFirewallObjectAddressOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerFirewallObjectAddress) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type FortimanagerFirewallObjectAddressArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallObjectAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerFirewallObjectAddress)(nil)).Elem()
}

func (o FortimanagerFirewallObjectAddressArrayOutput) ToFortimanagerFirewallObjectAddressArrayOutput() FortimanagerFirewallObjectAddressArrayOutput {
	return o
}

func (o FortimanagerFirewallObjectAddressArrayOutput) ToFortimanagerFirewallObjectAddressArrayOutputWithContext(ctx context.Context) FortimanagerFirewallObjectAddressArrayOutput {
	return o
}

func (o FortimanagerFirewallObjectAddressArrayOutput) Index(i pulumi.IntInput) FortimanagerFirewallObjectAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerFirewallObjectAddress {
		return vs[0].([]*FortimanagerFirewallObjectAddress)[vs[1].(int)]
	}).(FortimanagerFirewallObjectAddressOutput)
}

type FortimanagerFirewallObjectAddressMapOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallObjectAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerFirewallObjectAddress)(nil)).Elem()
}

func (o FortimanagerFirewallObjectAddressMapOutput) ToFortimanagerFirewallObjectAddressMapOutput() FortimanagerFirewallObjectAddressMapOutput {
	return o
}

func (o FortimanagerFirewallObjectAddressMapOutput) ToFortimanagerFirewallObjectAddressMapOutputWithContext(ctx context.Context) FortimanagerFirewallObjectAddressMapOutput {
	return o
}

func (o FortimanagerFirewallObjectAddressMapOutput) MapIndex(k pulumi.StringInput) FortimanagerFirewallObjectAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerFirewallObjectAddress {
		return vs[0].(map[string]*FortimanagerFirewallObjectAddress)[vs[1].(string)]
	}).(FortimanagerFirewallObjectAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerFirewallObjectAddressInput)(nil)).Elem(), &FortimanagerFirewallObjectAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerFirewallObjectAddressArrayInput)(nil)).Elem(), FortimanagerFirewallObjectAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerFirewallObjectAddressMapInput)(nil)).Elem(), FortimanagerFirewallObjectAddressMap{})
	pulumi.RegisterOutputType(FortimanagerFirewallObjectAddressOutput{})
	pulumi.RegisterOutputType(FortimanagerFirewallObjectAddressArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerFirewallObjectAddressMapOutput{})
}

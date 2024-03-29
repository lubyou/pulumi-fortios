// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallVendorMac struct {
	pulumi.CustomResourceState

	Fosid     pulumi.IntOutput       `pulumi:"fosid"`
	MacNumber pulumi.IntOutput       `pulumi:"macNumber"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	Obsolete  pulumi.IntOutput       `pulumi:"obsolete"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallVendorMac registers a new resource with the given unique name, arguments, and options.
func NewFirewallVendorMac(ctx *pulumi.Context,
	name string, args *FirewallVendorMacArgs, opts ...pulumi.ResourceOption) (*FirewallVendorMac, error) {
	if args == nil {
		args = &FirewallVendorMacArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallVendorMac
	err := ctx.RegisterResource("fortios:index/firewallVendorMac:FirewallVendorMac", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallVendorMac gets an existing FirewallVendorMac resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallVendorMac(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallVendorMacState, opts ...pulumi.ResourceOption) (*FirewallVendorMac, error) {
	var resource FirewallVendorMac
	err := ctx.ReadResource("fortios:index/firewallVendorMac:FirewallVendorMac", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallVendorMac resources.
type firewallVendorMacState struct {
	Fosid     *int    `pulumi:"fosid"`
	MacNumber *int    `pulumi:"macNumber"`
	Name      *string `pulumi:"name"`
	Obsolete  *int    `pulumi:"obsolete"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallVendorMacState struct {
	Fosid     pulumi.IntPtrInput
	MacNumber pulumi.IntPtrInput
	Name      pulumi.StringPtrInput
	Obsolete  pulumi.IntPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (FirewallVendorMacState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVendorMacState)(nil)).Elem()
}

type firewallVendorMacArgs struct {
	Fosid     *int    `pulumi:"fosid"`
	MacNumber *int    `pulumi:"macNumber"`
	Name      *string `pulumi:"name"`
	Obsolete  *int    `pulumi:"obsolete"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallVendorMac resource.
type FirewallVendorMacArgs struct {
	Fosid     pulumi.IntPtrInput
	MacNumber pulumi.IntPtrInput
	Name      pulumi.StringPtrInput
	Obsolete  pulumi.IntPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (FirewallVendorMacArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVendorMacArgs)(nil)).Elem()
}

type FirewallVendorMacInput interface {
	pulumi.Input

	ToFirewallVendorMacOutput() FirewallVendorMacOutput
	ToFirewallVendorMacOutputWithContext(ctx context.Context) FirewallVendorMacOutput
}

func (*FirewallVendorMac) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVendorMac)(nil)).Elem()
}

func (i *FirewallVendorMac) ToFirewallVendorMacOutput() FirewallVendorMacOutput {
	return i.ToFirewallVendorMacOutputWithContext(context.Background())
}

func (i *FirewallVendorMac) ToFirewallVendorMacOutputWithContext(ctx context.Context) FirewallVendorMacOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVendorMacOutput)
}

// FirewallVendorMacArrayInput is an input type that accepts FirewallVendorMacArray and FirewallVendorMacArrayOutput values.
// You can construct a concrete instance of `FirewallVendorMacArrayInput` via:
//
//	FirewallVendorMacArray{ FirewallVendorMacArgs{...} }
type FirewallVendorMacArrayInput interface {
	pulumi.Input

	ToFirewallVendorMacArrayOutput() FirewallVendorMacArrayOutput
	ToFirewallVendorMacArrayOutputWithContext(context.Context) FirewallVendorMacArrayOutput
}

type FirewallVendorMacArray []FirewallVendorMacInput

func (FirewallVendorMacArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVendorMac)(nil)).Elem()
}

func (i FirewallVendorMacArray) ToFirewallVendorMacArrayOutput() FirewallVendorMacArrayOutput {
	return i.ToFirewallVendorMacArrayOutputWithContext(context.Background())
}

func (i FirewallVendorMacArray) ToFirewallVendorMacArrayOutputWithContext(ctx context.Context) FirewallVendorMacArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVendorMacArrayOutput)
}

// FirewallVendorMacMapInput is an input type that accepts FirewallVendorMacMap and FirewallVendorMacMapOutput values.
// You can construct a concrete instance of `FirewallVendorMacMapInput` via:
//
//	FirewallVendorMacMap{ "key": FirewallVendorMacArgs{...} }
type FirewallVendorMacMapInput interface {
	pulumi.Input

	ToFirewallVendorMacMapOutput() FirewallVendorMacMapOutput
	ToFirewallVendorMacMapOutputWithContext(context.Context) FirewallVendorMacMapOutput
}

type FirewallVendorMacMap map[string]FirewallVendorMacInput

func (FirewallVendorMacMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVendorMac)(nil)).Elem()
}

func (i FirewallVendorMacMap) ToFirewallVendorMacMapOutput() FirewallVendorMacMapOutput {
	return i.ToFirewallVendorMacMapOutputWithContext(context.Background())
}

func (i FirewallVendorMacMap) ToFirewallVendorMacMapOutputWithContext(ctx context.Context) FirewallVendorMacMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVendorMacMapOutput)
}

type FirewallVendorMacOutput struct{ *pulumi.OutputState }

func (FirewallVendorMacOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVendorMac)(nil)).Elem()
}

func (o FirewallVendorMacOutput) ToFirewallVendorMacOutput() FirewallVendorMacOutput {
	return o
}

func (o FirewallVendorMacOutput) ToFirewallVendorMacOutputWithContext(ctx context.Context) FirewallVendorMacOutput {
	return o
}

func (o FirewallVendorMacOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallVendorMac) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o FirewallVendorMacOutput) MacNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallVendorMac) pulumi.IntOutput { return v.MacNumber }).(pulumi.IntOutput)
}

func (o FirewallVendorMacOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVendorMac) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallVendorMacOutput) Obsolete() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallVendorMac) pulumi.IntOutput { return v.Obsolete }).(pulumi.IntOutput)
}

func (o FirewallVendorMacOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVendorMac) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallVendorMacArrayOutput struct{ *pulumi.OutputState }

func (FirewallVendorMacArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVendorMac)(nil)).Elem()
}

func (o FirewallVendorMacArrayOutput) ToFirewallVendorMacArrayOutput() FirewallVendorMacArrayOutput {
	return o
}

func (o FirewallVendorMacArrayOutput) ToFirewallVendorMacArrayOutputWithContext(ctx context.Context) FirewallVendorMacArrayOutput {
	return o
}

func (o FirewallVendorMacArrayOutput) Index(i pulumi.IntInput) FirewallVendorMacOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallVendorMac {
		return vs[0].([]*FirewallVendorMac)[vs[1].(int)]
	}).(FirewallVendorMacOutput)
}

type FirewallVendorMacMapOutput struct{ *pulumi.OutputState }

func (FirewallVendorMacMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVendorMac)(nil)).Elem()
}

func (o FirewallVendorMacMapOutput) ToFirewallVendorMacMapOutput() FirewallVendorMacMapOutput {
	return o
}

func (o FirewallVendorMacMapOutput) ToFirewallVendorMacMapOutputWithContext(ctx context.Context) FirewallVendorMacMapOutput {
	return o
}

func (o FirewallVendorMacMapOutput) MapIndex(k pulumi.StringInput) FirewallVendorMacOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallVendorMac {
		return vs[0].(map[string]*FirewallVendorMac)[vs[1].(string)]
	}).(FirewallVendorMacOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVendorMacInput)(nil)).Elem(), &FirewallVendorMac{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVendorMacArrayInput)(nil)).Elem(), FirewallVendorMacArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVendorMacMapInput)(nil)).Elem(), FirewallVendorMacMap{})
	pulumi.RegisterOutputType(FirewallVendorMacOutput{})
	pulumi.RegisterOutputType(FirewallVendorMacArrayOutput{})
	pulumi.RegisterOutputType(FirewallVendorMacMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallIpTranslation struct {
	pulumi.CustomResourceState

	Endip      pulumi.StringOutput    `pulumi:"endip"`
	MapStartip pulumi.StringOutput    `pulumi:"mapStartip"`
	Startip    pulumi.StringOutput    `pulumi:"startip"`
	Transid    pulumi.IntOutput       `pulumi:"transid"`
	Type       pulumi.StringOutput    `pulumi:"type"`
	Vdomparam  pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallIpTranslation registers a new resource with the given unique name, arguments, and options.
func NewFirewallIpTranslation(ctx *pulumi.Context,
	name string, args *FirewallIpTranslationArgs, opts ...pulumi.ResourceOption) (*FirewallIpTranslation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endip == nil {
		return nil, errors.New("invalid value for required argument 'Endip'")
	}
	if args.MapStartip == nil {
		return nil, errors.New("invalid value for required argument 'MapStartip'")
	}
	if args.Startip == nil {
		return nil, errors.New("invalid value for required argument 'Startip'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallIpTranslation
	err := ctx.RegisterResource("fortios:index/firewallIpTranslation:FirewallIpTranslation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallIpTranslation gets an existing FirewallIpTranslation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallIpTranslation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallIpTranslationState, opts ...pulumi.ResourceOption) (*FirewallIpTranslation, error) {
	var resource FirewallIpTranslation
	err := ctx.ReadResource("fortios:index/firewallIpTranslation:FirewallIpTranslation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallIpTranslation resources.
type firewallIpTranslationState struct {
	Endip      *string `pulumi:"endip"`
	MapStartip *string `pulumi:"mapStartip"`
	Startip    *string `pulumi:"startip"`
	Transid    *int    `pulumi:"transid"`
	Type       *string `pulumi:"type"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

type FirewallIpTranslationState struct {
	Endip      pulumi.StringPtrInput
	MapStartip pulumi.StringPtrInput
	Startip    pulumi.StringPtrInput
	Transid    pulumi.IntPtrInput
	Type       pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (FirewallIpTranslationState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIpTranslationState)(nil)).Elem()
}

type firewallIpTranslationArgs struct {
	Endip      string  `pulumi:"endip"`
	MapStartip string  `pulumi:"mapStartip"`
	Startip    string  `pulumi:"startip"`
	Transid    *int    `pulumi:"transid"`
	Type       *string `pulumi:"type"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallIpTranslation resource.
type FirewallIpTranslationArgs struct {
	Endip      pulumi.StringInput
	MapStartip pulumi.StringInput
	Startip    pulumi.StringInput
	Transid    pulumi.IntPtrInput
	Type       pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (FirewallIpTranslationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIpTranslationArgs)(nil)).Elem()
}

type FirewallIpTranslationInput interface {
	pulumi.Input

	ToFirewallIpTranslationOutput() FirewallIpTranslationOutput
	ToFirewallIpTranslationOutputWithContext(ctx context.Context) FirewallIpTranslationOutput
}

func (*FirewallIpTranslation) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIpTranslation)(nil)).Elem()
}

func (i *FirewallIpTranslation) ToFirewallIpTranslationOutput() FirewallIpTranslationOutput {
	return i.ToFirewallIpTranslationOutputWithContext(context.Background())
}

func (i *FirewallIpTranslation) ToFirewallIpTranslationOutputWithContext(ctx context.Context) FirewallIpTranslationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIpTranslationOutput)
}

// FirewallIpTranslationArrayInput is an input type that accepts FirewallIpTranslationArray and FirewallIpTranslationArrayOutput values.
// You can construct a concrete instance of `FirewallIpTranslationArrayInput` via:
//
//	FirewallIpTranslationArray{ FirewallIpTranslationArgs{...} }
type FirewallIpTranslationArrayInput interface {
	pulumi.Input

	ToFirewallIpTranslationArrayOutput() FirewallIpTranslationArrayOutput
	ToFirewallIpTranslationArrayOutputWithContext(context.Context) FirewallIpTranslationArrayOutput
}

type FirewallIpTranslationArray []FirewallIpTranslationInput

func (FirewallIpTranslationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallIpTranslation)(nil)).Elem()
}

func (i FirewallIpTranslationArray) ToFirewallIpTranslationArrayOutput() FirewallIpTranslationArrayOutput {
	return i.ToFirewallIpTranslationArrayOutputWithContext(context.Background())
}

func (i FirewallIpTranslationArray) ToFirewallIpTranslationArrayOutputWithContext(ctx context.Context) FirewallIpTranslationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIpTranslationArrayOutput)
}

// FirewallIpTranslationMapInput is an input type that accepts FirewallIpTranslationMap and FirewallIpTranslationMapOutput values.
// You can construct a concrete instance of `FirewallIpTranslationMapInput` via:
//
//	FirewallIpTranslationMap{ "key": FirewallIpTranslationArgs{...} }
type FirewallIpTranslationMapInput interface {
	pulumi.Input

	ToFirewallIpTranslationMapOutput() FirewallIpTranslationMapOutput
	ToFirewallIpTranslationMapOutputWithContext(context.Context) FirewallIpTranslationMapOutput
}

type FirewallIpTranslationMap map[string]FirewallIpTranslationInput

func (FirewallIpTranslationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallIpTranslation)(nil)).Elem()
}

func (i FirewallIpTranslationMap) ToFirewallIpTranslationMapOutput() FirewallIpTranslationMapOutput {
	return i.ToFirewallIpTranslationMapOutputWithContext(context.Background())
}

func (i FirewallIpTranslationMap) ToFirewallIpTranslationMapOutputWithContext(ctx context.Context) FirewallIpTranslationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIpTranslationMapOutput)
}

type FirewallIpTranslationOutput struct{ *pulumi.OutputState }

func (FirewallIpTranslationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIpTranslation)(nil)).Elem()
}

func (o FirewallIpTranslationOutput) ToFirewallIpTranslationOutput() FirewallIpTranslationOutput {
	return o
}

func (o FirewallIpTranslationOutput) ToFirewallIpTranslationOutputWithContext(ctx context.Context) FirewallIpTranslationOutput {
	return o
}

func (o FirewallIpTranslationOutput) Endip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallIpTranslation) pulumi.StringOutput { return v.Endip }).(pulumi.StringOutput)
}

func (o FirewallIpTranslationOutput) MapStartip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallIpTranslation) pulumi.StringOutput { return v.MapStartip }).(pulumi.StringOutput)
}

func (o FirewallIpTranslationOutput) Startip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallIpTranslation) pulumi.StringOutput { return v.Startip }).(pulumi.StringOutput)
}

func (o FirewallIpTranslationOutput) Transid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallIpTranslation) pulumi.IntOutput { return v.Transid }).(pulumi.IntOutput)
}

func (o FirewallIpTranslationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallIpTranslation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o FirewallIpTranslationOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallIpTranslation) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallIpTranslationArrayOutput struct{ *pulumi.OutputState }

func (FirewallIpTranslationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallIpTranslation)(nil)).Elem()
}

func (o FirewallIpTranslationArrayOutput) ToFirewallIpTranslationArrayOutput() FirewallIpTranslationArrayOutput {
	return o
}

func (o FirewallIpTranslationArrayOutput) ToFirewallIpTranslationArrayOutputWithContext(ctx context.Context) FirewallIpTranslationArrayOutput {
	return o
}

func (o FirewallIpTranslationArrayOutput) Index(i pulumi.IntInput) FirewallIpTranslationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallIpTranslation {
		return vs[0].([]*FirewallIpTranslation)[vs[1].(int)]
	}).(FirewallIpTranslationOutput)
}

type FirewallIpTranslationMapOutput struct{ *pulumi.OutputState }

func (FirewallIpTranslationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallIpTranslation)(nil)).Elem()
}

func (o FirewallIpTranslationMapOutput) ToFirewallIpTranslationMapOutput() FirewallIpTranslationMapOutput {
	return o
}

func (o FirewallIpTranslationMapOutput) ToFirewallIpTranslationMapOutputWithContext(ctx context.Context) FirewallIpTranslationMapOutput {
	return o
}

func (o FirewallIpTranslationMapOutput) MapIndex(k pulumi.StringInput) FirewallIpTranslationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallIpTranslation {
		return vs[0].(map[string]*FirewallIpTranslation)[vs[1].(string)]
	}).(FirewallIpTranslationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIpTranslationInput)(nil)).Elem(), &FirewallIpTranslation{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIpTranslationArrayInput)(nil)).Elem(), FirewallIpTranslationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIpTranslationMapInput)(nil)).Elem(), FirewallIpTranslationMap{})
	pulumi.RegisterOutputType(FirewallIpTranslationOutput{})
	pulumi.RegisterOutputType(FirewallIpTranslationArrayOutput{})
	pulumi.RegisterOutputType(FirewallIpTranslationMapOutput{})
}

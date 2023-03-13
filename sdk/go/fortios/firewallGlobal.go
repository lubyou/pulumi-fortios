// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallGlobal struct {
	pulumi.CustomResourceState

	BannedIpPersistency pulumi.StringOutput    `pulumi:"bannedIpPersistency"`
	Vdomparam           pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallGlobal registers a new resource with the given unique name, arguments, and options.
func NewFirewallGlobal(ctx *pulumi.Context,
	name string, args *FirewallGlobalArgs, opts ...pulumi.ResourceOption) (*FirewallGlobal, error) {
	if args == nil {
		args = &FirewallGlobalArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallGlobal
	err := ctx.RegisterResource("fortios:index/firewallGlobal:FirewallGlobal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallGlobal gets an existing FirewallGlobal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallGlobal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallGlobalState, opts ...pulumi.ResourceOption) (*FirewallGlobal, error) {
	var resource FirewallGlobal
	err := ctx.ReadResource("fortios:index/firewallGlobal:FirewallGlobal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallGlobal resources.
type firewallGlobalState struct {
	BannedIpPersistency *string `pulumi:"bannedIpPersistency"`
	Vdomparam           *string `pulumi:"vdomparam"`
}

type FirewallGlobalState struct {
	BannedIpPersistency pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallGlobalState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallGlobalState)(nil)).Elem()
}

type firewallGlobalArgs struct {
	BannedIpPersistency *string `pulumi:"bannedIpPersistency"`
	Vdomparam           *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallGlobal resource.
type FirewallGlobalArgs struct {
	BannedIpPersistency pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallGlobalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallGlobalArgs)(nil)).Elem()
}

type FirewallGlobalInput interface {
	pulumi.Input

	ToFirewallGlobalOutput() FirewallGlobalOutput
	ToFirewallGlobalOutputWithContext(ctx context.Context) FirewallGlobalOutput
}

func (*FirewallGlobal) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallGlobal)(nil)).Elem()
}

func (i *FirewallGlobal) ToFirewallGlobalOutput() FirewallGlobalOutput {
	return i.ToFirewallGlobalOutputWithContext(context.Background())
}

func (i *FirewallGlobal) ToFirewallGlobalOutputWithContext(ctx context.Context) FirewallGlobalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallGlobalOutput)
}

// FirewallGlobalArrayInput is an input type that accepts FirewallGlobalArray and FirewallGlobalArrayOutput values.
// You can construct a concrete instance of `FirewallGlobalArrayInput` via:
//
//	FirewallGlobalArray{ FirewallGlobalArgs{...} }
type FirewallGlobalArrayInput interface {
	pulumi.Input

	ToFirewallGlobalArrayOutput() FirewallGlobalArrayOutput
	ToFirewallGlobalArrayOutputWithContext(context.Context) FirewallGlobalArrayOutput
}

type FirewallGlobalArray []FirewallGlobalInput

func (FirewallGlobalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallGlobal)(nil)).Elem()
}

func (i FirewallGlobalArray) ToFirewallGlobalArrayOutput() FirewallGlobalArrayOutput {
	return i.ToFirewallGlobalArrayOutputWithContext(context.Background())
}

func (i FirewallGlobalArray) ToFirewallGlobalArrayOutputWithContext(ctx context.Context) FirewallGlobalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallGlobalArrayOutput)
}

// FirewallGlobalMapInput is an input type that accepts FirewallGlobalMap and FirewallGlobalMapOutput values.
// You can construct a concrete instance of `FirewallGlobalMapInput` via:
//
//	FirewallGlobalMap{ "key": FirewallGlobalArgs{...} }
type FirewallGlobalMapInput interface {
	pulumi.Input

	ToFirewallGlobalMapOutput() FirewallGlobalMapOutput
	ToFirewallGlobalMapOutputWithContext(context.Context) FirewallGlobalMapOutput
}

type FirewallGlobalMap map[string]FirewallGlobalInput

func (FirewallGlobalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallGlobal)(nil)).Elem()
}

func (i FirewallGlobalMap) ToFirewallGlobalMapOutput() FirewallGlobalMapOutput {
	return i.ToFirewallGlobalMapOutputWithContext(context.Background())
}

func (i FirewallGlobalMap) ToFirewallGlobalMapOutputWithContext(ctx context.Context) FirewallGlobalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallGlobalMapOutput)
}

type FirewallGlobalOutput struct{ *pulumi.OutputState }

func (FirewallGlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallGlobal)(nil)).Elem()
}

func (o FirewallGlobalOutput) ToFirewallGlobalOutput() FirewallGlobalOutput {
	return o
}

func (o FirewallGlobalOutput) ToFirewallGlobalOutputWithContext(ctx context.Context) FirewallGlobalOutput {
	return o
}

func (o FirewallGlobalOutput) BannedIpPersistency() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallGlobal) pulumi.StringOutput { return v.BannedIpPersistency }).(pulumi.StringOutput)
}

func (o FirewallGlobalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallGlobal) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallGlobalArrayOutput struct{ *pulumi.OutputState }

func (FirewallGlobalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallGlobal)(nil)).Elem()
}

func (o FirewallGlobalArrayOutput) ToFirewallGlobalArrayOutput() FirewallGlobalArrayOutput {
	return o
}

func (o FirewallGlobalArrayOutput) ToFirewallGlobalArrayOutputWithContext(ctx context.Context) FirewallGlobalArrayOutput {
	return o
}

func (o FirewallGlobalArrayOutput) Index(i pulumi.IntInput) FirewallGlobalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallGlobal {
		return vs[0].([]*FirewallGlobal)[vs[1].(int)]
	}).(FirewallGlobalOutput)
}

type FirewallGlobalMapOutput struct{ *pulumi.OutputState }

func (FirewallGlobalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallGlobal)(nil)).Elem()
}

func (o FirewallGlobalMapOutput) ToFirewallGlobalMapOutput() FirewallGlobalMapOutput {
	return o
}

func (o FirewallGlobalMapOutput) ToFirewallGlobalMapOutputWithContext(ctx context.Context) FirewallGlobalMapOutput {
	return o
}

func (o FirewallGlobalMapOutput) MapIndex(k pulumi.StringInput) FirewallGlobalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallGlobal {
		return vs[0].(map[string]*FirewallGlobal)[vs[1].(string)]
	}).(FirewallGlobalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallGlobalInput)(nil)).Elem(), &FirewallGlobal{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallGlobalArrayInput)(nil)).Elem(), FirewallGlobalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallGlobalMapInput)(nil)).Elem(), FirewallGlobalMap{})
	pulumi.RegisterOutputType(FirewallGlobalOutput{})
	pulumi.RegisterOutputType(FirewallGlobalArrayOutput{})
	pulumi.RegisterOutputType(FirewallGlobalMapOutput{})
}

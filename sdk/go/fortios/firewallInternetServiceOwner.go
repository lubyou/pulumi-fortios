// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallInternetServiceOwner struct {
	pulumi.CustomResourceState

	Fosid     pulumi.IntOutput       `pulumi:"fosid"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetServiceOwner registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetServiceOwner(ctx *pulumi.Context,
	name string, args *FirewallInternetServiceOwnerArgs, opts ...pulumi.ResourceOption) (*FirewallInternetServiceOwner, error) {
	if args == nil {
		args = &FirewallInternetServiceOwnerArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetServiceOwner
	err := ctx.RegisterResource("fortios:index/firewallInternetServiceOwner:FirewallInternetServiceOwner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetServiceOwner gets an existing FirewallInternetServiceOwner resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetServiceOwner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetServiceOwnerState, opts ...pulumi.ResourceOption) (*FirewallInternetServiceOwner, error) {
	var resource FirewallInternetServiceOwner
	err := ctx.ReadResource("fortios:index/firewallInternetServiceOwner:FirewallInternetServiceOwner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetServiceOwner resources.
type firewallInternetServiceOwnerState struct {
	Fosid     *int    `pulumi:"fosid"`
	Name      *string `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetServiceOwnerState struct {
	Fosid     pulumi.IntPtrInput
	Name      pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceOwnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceOwnerState)(nil)).Elem()
}

type firewallInternetServiceOwnerArgs struct {
	Fosid     *int    `pulumi:"fosid"`
	Name      *string `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetServiceOwner resource.
type FirewallInternetServiceOwnerArgs struct {
	Fosid     pulumi.IntPtrInput
	Name      pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceOwnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceOwnerArgs)(nil)).Elem()
}

type FirewallInternetServiceOwnerInput interface {
	pulumi.Input

	ToFirewallInternetServiceOwnerOutput() FirewallInternetServiceOwnerOutput
	ToFirewallInternetServiceOwnerOutputWithContext(ctx context.Context) FirewallInternetServiceOwnerOutput
}

func (*FirewallInternetServiceOwner) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceOwner)(nil)).Elem()
}

func (i *FirewallInternetServiceOwner) ToFirewallInternetServiceOwnerOutput() FirewallInternetServiceOwnerOutput {
	return i.ToFirewallInternetServiceOwnerOutputWithContext(context.Background())
}

func (i *FirewallInternetServiceOwner) ToFirewallInternetServiceOwnerOutputWithContext(ctx context.Context) FirewallInternetServiceOwnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceOwnerOutput)
}

// FirewallInternetServiceOwnerArrayInput is an input type that accepts FirewallInternetServiceOwnerArray and FirewallInternetServiceOwnerArrayOutput values.
// You can construct a concrete instance of `FirewallInternetServiceOwnerArrayInput` via:
//
//	FirewallInternetServiceOwnerArray{ FirewallInternetServiceOwnerArgs{...} }
type FirewallInternetServiceOwnerArrayInput interface {
	pulumi.Input

	ToFirewallInternetServiceOwnerArrayOutput() FirewallInternetServiceOwnerArrayOutput
	ToFirewallInternetServiceOwnerArrayOutputWithContext(context.Context) FirewallInternetServiceOwnerArrayOutput
}

type FirewallInternetServiceOwnerArray []FirewallInternetServiceOwnerInput

func (FirewallInternetServiceOwnerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceOwner)(nil)).Elem()
}

func (i FirewallInternetServiceOwnerArray) ToFirewallInternetServiceOwnerArrayOutput() FirewallInternetServiceOwnerArrayOutput {
	return i.ToFirewallInternetServiceOwnerArrayOutputWithContext(context.Background())
}

func (i FirewallInternetServiceOwnerArray) ToFirewallInternetServiceOwnerArrayOutputWithContext(ctx context.Context) FirewallInternetServiceOwnerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceOwnerArrayOutput)
}

// FirewallInternetServiceOwnerMapInput is an input type that accepts FirewallInternetServiceOwnerMap and FirewallInternetServiceOwnerMapOutput values.
// You can construct a concrete instance of `FirewallInternetServiceOwnerMapInput` via:
//
//	FirewallInternetServiceOwnerMap{ "key": FirewallInternetServiceOwnerArgs{...} }
type FirewallInternetServiceOwnerMapInput interface {
	pulumi.Input

	ToFirewallInternetServiceOwnerMapOutput() FirewallInternetServiceOwnerMapOutput
	ToFirewallInternetServiceOwnerMapOutputWithContext(context.Context) FirewallInternetServiceOwnerMapOutput
}

type FirewallInternetServiceOwnerMap map[string]FirewallInternetServiceOwnerInput

func (FirewallInternetServiceOwnerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceOwner)(nil)).Elem()
}

func (i FirewallInternetServiceOwnerMap) ToFirewallInternetServiceOwnerMapOutput() FirewallInternetServiceOwnerMapOutput {
	return i.ToFirewallInternetServiceOwnerMapOutputWithContext(context.Background())
}

func (i FirewallInternetServiceOwnerMap) ToFirewallInternetServiceOwnerMapOutputWithContext(ctx context.Context) FirewallInternetServiceOwnerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceOwnerMapOutput)
}

type FirewallInternetServiceOwnerOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceOwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceOwner)(nil)).Elem()
}

func (o FirewallInternetServiceOwnerOutput) ToFirewallInternetServiceOwnerOutput() FirewallInternetServiceOwnerOutput {
	return o
}

func (o FirewallInternetServiceOwnerOutput) ToFirewallInternetServiceOwnerOutputWithContext(ctx context.Context) FirewallInternetServiceOwnerOutput {
	return o
}

func (o FirewallInternetServiceOwnerOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallInternetServiceOwner) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o FirewallInternetServiceOwnerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInternetServiceOwner) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallInternetServiceOwnerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetServiceOwner) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallInternetServiceOwnerArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceOwnerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceOwner)(nil)).Elem()
}

func (o FirewallInternetServiceOwnerArrayOutput) ToFirewallInternetServiceOwnerArrayOutput() FirewallInternetServiceOwnerArrayOutput {
	return o
}

func (o FirewallInternetServiceOwnerArrayOutput) ToFirewallInternetServiceOwnerArrayOutputWithContext(ctx context.Context) FirewallInternetServiceOwnerArrayOutput {
	return o
}

func (o FirewallInternetServiceOwnerArrayOutput) Index(i pulumi.IntInput) FirewallInternetServiceOwnerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetServiceOwner {
		return vs[0].([]*FirewallInternetServiceOwner)[vs[1].(int)]
	}).(FirewallInternetServiceOwnerOutput)
}

type FirewallInternetServiceOwnerMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceOwnerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceOwner)(nil)).Elem()
}

func (o FirewallInternetServiceOwnerMapOutput) ToFirewallInternetServiceOwnerMapOutput() FirewallInternetServiceOwnerMapOutput {
	return o
}

func (o FirewallInternetServiceOwnerMapOutput) ToFirewallInternetServiceOwnerMapOutputWithContext(ctx context.Context) FirewallInternetServiceOwnerMapOutput {
	return o
}

func (o FirewallInternetServiceOwnerMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetServiceOwnerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetServiceOwner {
		return vs[0].(map[string]*FirewallInternetServiceOwner)[vs[1].(string)]
	}).(FirewallInternetServiceOwnerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceOwnerInput)(nil)).Elem(), &FirewallInternetServiceOwner{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceOwnerArrayInput)(nil)).Elem(), FirewallInternetServiceOwnerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceOwnerMapInput)(nil)).Elem(), FirewallInternetServiceOwnerMap{})
	pulumi.RegisterOutputType(FirewallInternetServiceOwnerOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceOwnerArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceOwnerMapOutput{})
}

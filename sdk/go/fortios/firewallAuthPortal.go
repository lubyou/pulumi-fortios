// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallAuthPortal struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput             `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput             `pulumi:"getAllTables"`
	Groups              FirewallAuthPortalGroupArrayOutput `pulumi:"groups"`
	IdentityBasedRoute  pulumi.StringOutput                `pulumi:"identityBasedRoute"`
	PortalAddr          pulumi.StringOutput                `pulumi:"portalAddr"`
	PortalAddr6         pulumi.StringOutput                `pulumi:"portalAddr6"`
	ProxyAuth           pulumi.StringOutput                `pulumi:"proxyAuth"`
	Vdomparam           pulumi.StringPtrOutput             `pulumi:"vdomparam"`
}

// NewFirewallAuthPortal registers a new resource with the given unique name, arguments, and options.
func NewFirewallAuthPortal(ctx *pulumi.Context,
	name string, args *FirewallAuthPortalArgs, opts ...pulumi.ResourceOption) (*FirewallAuthPortal, error) {
	if args == nil {
		args = &FirewallAuthPortalArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallAuthPortal
	err := ctx.RegisterResource("fortios:index/firewallAuthPortal:FirewallAuthPortal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallAuthPortal gets an existing FirewallAuthPortal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallAuthPortal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallAuthPortalState, opts ...pulumi.ResourceOption) (*FirewallAuthPortal, error) {
	var resource FirewallAuthPortal
	err := ctx.ReadResource("fortios:index/firewallAuthPortal:FirewallAuthPortal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallAuthPortal resources.
type firewallAuthPortalState struct {
	DynamicSortSubtable *string                   `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                   `pulumi:"getAllTables"`
	Groups              []FirewallAuthPortalGroup `pulumi:"groups"`
	IdentityBasedRoute  *string                   `pulumi:"identityBasedRoute"`
	PortalAddr          *string                   `pulumi:"portalAddr"`
	PortalAddr6         *string                   `pulumi:"portalAddr6"`
	ProxyAuth           *string                   `pulumi:"proxyAuth"`
	Vdomparam           *string                   `pulumi:"vdomparam"`
}

type FirewallAuthPortalState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Groups              FirewallAuthPortalGroupArrayInput
	IdentityBasedRoute  pulumi.StringPtrInput
	PortalAddr          pulumi.StringPtrInput
	PortalAddr6         pulumi.StringPtrInput
	ProxyAuth           pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallAuthPortalState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAuthPortalState)(nil)).Elem()
}

type firewallAuthPortalArgs struct {
	DynamicSortSubtable *string                   `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                   `pulumi:"getAllTables"`
	Groups              []FirewallAuthPortalGroup `pulumi:"groups"`
	IdentityBasedRoute  *string                   `pulumi:"identityBasedRoute"`
	PortalAddr          *string                   `pulumi:"portalAddr"`
	PortalAddr6         *string                   `pulumi:"portalAddr6"`
	ProxyAuth           *string                   `pulumi:"proxyAuth"`
	Vdomparam           *string                   `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallAuthPortal resource.
type FirewallAuthPortalArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Groups              FirewallAuthPortalGroupArrayInput
	IdentityBasedRoute  pulumi.StringPtrInput
	PortalAddr          pulumi.StringPtrInput
	PortalAddr6         pulumi.StringPtrInput
	ProxyAuth           pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallAuthPortalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAuthPortalArgs)(nil)).Elem()
}

type FirewallAuthPortalInput interface {
	pulumi.Input

	ToFirewallAuthPortalOutput() FirewallAuthPortalOutput
	ToFirewallAuthPortalOutputWithContext(ctx context.Context) FirewallAuthPortalOutput
}

func (*FirewallAuthPortal) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAuthPortal)(nil)).Elem()
}

func (i *FirewallAuthPortal) ToFirewallAuthPortalOutput() FirewallAuthPortalOutput {
	return i.ToFirewallAuthPortalOutputWithContext(context.Background())
}

func (i *FirewallAuthPortal) ToFirewallAuthPortalOutputWithContext(ctx context.Context) FirewallAuthPortalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAuthPortalOutput)
}

// FirewallAuthPortalArrayInput is an input type that accepts FirewallAuthPortalArray and FirewallAuthPortalArrayOutput values.
// You can construct a concrete instance of `FirewallAuthPortalArrayInput` via:
//
//	FirewallAuthPortalArray{ FirewallAuthPortalArgs{...} }
type FirewallAuthPortalArrayInput interface {
	pulumi.Input

	ToFirewallAuthPortalArrayOutput() FirewallAuthPortalArrayOutput
	ToFirewallAuthPortalArrayOutputWithContext(context.Context) FirewallAuthPortalArrayOutput
}

type FirewallAuthPortalArray []FirewallAuthPortalInput

func (FirewallAuthPortalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallAuthPortal)(nil)).Elem()
}

func (i FirewallAuthPortalArray) ToFirewallAuthPortalArrayOutput() FirewallAuthPortalArrayOutput {
	return i.ToFirewallAuthPortalArrayOutputWithContext(context.Background())
}

func (i FirewallAuthPortalArray) ToFirewallAuthPortalArrayOutputWithContext(ctx context.Context) FirewallAuthPortalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAuthPortalArrayOutput)
}

// FirewallAuthPortalMapInput is an input type that accepts FirewallAuthPortalMap and FirewallAuthPortalMapOutput values.
// You can construct a concrete instance of `FirewallAuthPortalMapInput` via:
//
//	FirewallAuthPortalMap{ "key": FirewallAuthPortalArgs{...} }
type FirewallAuthPortalMapInput interface {
	pulumi.Input

	ToFirewallAuthPortalMapOutput() FirewallAuthPortalMapOutput
	ToFirewallAuthPortalMapOutputWithContext(context.Context) FirewallAuthPortalMapOutput
}

type FirewallAuthPortalMap map[string]FirewallAuthPortalInput

func (FirewallAuthPortalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallAuthPortal)(nil)).Elem()
}

func (i FirewallAuthPortalMap) ToFirewallAuthPortalMapOutput() FirewallAuthPortalMapOutput {
	return i.ToFirewallAuthPortalMapOutputWithContext(context.Background())
}

func (i FirewallAuthPortalMap) ToFirewallAuthPortalMapOutputWithContext(ctx context.Context) FirewallAuthPortalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAuthPortalMapOutput)
}

type FirewallAuthPortalOutput struct{ *pulumi.OutputState }

func (FirewallAuthPortalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAuthPortal)(nil)).Elem()
}

func (o FirewallAuthPortalOutput) ToFirewallAuthPortalOutput() FirewallAuthPortalOutput {
	return o
}

func (o FirewallAuthPortalOutput) ToFirewallAuthPortalOutputWithContext(ctx context.Context) FirewallAuthPortalOutput {
	return o
}

func (o FirewallAuthPortalOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAuthPortal) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallAuthPortalOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAuthPortal) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o FirewallAuthPortalOutput) Groups() FirewallAuthPortalGroupArrayOutput {
	return o.ApplyT(func(v *FirewallAuthPortal) FirewallAuthPortalGroupArrayOutput { return v.Groups }).(FirewallAuthPortalGroupArrayOutput)
}

func (o FirewallAuthPortalOutput) IdentityBasedRoute() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAuthPortal) pulumi.StringOutput { return v.IdentityBasedRoute }).(pulumi.StringOutput)
}

func (o FirewallAuthPortalOutput) PortalAddr() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAuthPortal) pulumi.StringOutput { return v.PortalAddr }).(pulumi.StringOutput)
}

func (o FirewallAuthPortalOutput) PortalAddr6() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAuthPortal) pulumi.StringOutput { return v.PortalAddr6 }).(pulumi.StringOutput)
}

func (o FirewallAuthPortalOutput) ProxyAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAuthPortal) pulumi.StringOutput { return v.ProxyAuth }).(pulumi.StringOutput)
}

func (o FirewallAuthPortalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAuthPortal) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallAuthPortalArrayOutput struct{ *pulumi.OutputState }

func (FirewallAuthPortalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallAuthPortal)(nil)).Elem()
}

func (o FirewallAuthPortalArrayOutput) ToFirewallAuthPortalArrayOutput() FirewallAuthPortalArrayOutput {
	return o
}

func (o FirewallAuthPortalArrayOutput) ToFirewallAuthPortalArrayOutputWithContext(ctx context.Context) FirewallAuthPortalArrayOutput {
	return o
}

func (o FirewallAuthPortalArrayOutput) Index(i pulumi.IntInput) FirewallAuthPortalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallAuthPortal {
		return vs[0].([]*FirewallAuthPortal)[vs[1].(int)]
	}).(FirewallAuthPortalOutput)
}

type FirewallAuthPortalMapOutput struct{ *pulumi.OutputState }

func (FirewallAuthPortalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallAuthPortal)(nil)).Elem()
}

func (o FirewallAuthPortalMapOutput) ToFirewallAuthPortalMapOutput() FirewallAuthPortalMapOutput {
	return o
}

func (o FirewallAuthPortalMapOutput) ToFirewallAuthPortalMapOutputWithContext(ctx context.Context) FirewallAuthPortalMapOutput {
	return o
}

func (o FirewallAuthPortalMapOutput) MapIndex(k pulumi.StringInput) FirewallAuthPortalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallAuthPortal {
		return vs[0].(map[string]*FirewallAuthPortal)[vs[1].(string)]
	}).(FirewallAuthPortalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAuthPortalInput)(nil)).Elem(), &FirewallAuthPortal{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAuthPortalArrayInput)(nil)).Elem(), FirewallAuthPortalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAuthPortalMapInput)(nil)).Elem(), FirewallAuthPortalMap{})
	pulumi.RegisterOutputType(FirewallAuthPortalOutput{})
	pulumi.RegisterOutputType(FirewallAuthPortalArrayOutput{})
	pulumi.RegisterOutputType(FirewallAuthPortalMapOutput{})
}

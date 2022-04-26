// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure firewall authentication portals.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallAuthPortal(ctx, "trname", &fortios.FirewallAuthPortalArgs{
// 			Groups: FirewallAuthPortalGroupArray{
// 				&FirewallAuthPortalGroupArgs{
// 					Name: pulumi.String("Guest-group"),
// 				},
// 			},
// 			PortalAddr: pulumi.String("1.1.1.1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Firewall AuthPortal can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallAuthPortal:FirewallAuthPortal labelname FirewallAuthPortal
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallAuthPortal:FirewallAuthPortal labelname FirewallAuthPortal
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallAuthPortal struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
	Groups FirewallAuthPortalGroupArrayOutput `pulumi:"groups"`
	// Name of the identity-based route that applies to this portal.
	IdentityBasedRoute pulumi.StringOutput `pulumi:"identityBasedRoute"`
	// Address (or FQDN) of the authentication portal.
	PortalAddr pulumi.StringOutput `pulumi:"portalAddr"`
	// IPv6 address (or FQDN) of authentication portal.
	PortalAddr6 pulumi.StringOutput `pulumi:"portalAddr6"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallAuthPortal registers a new resource with the given unique name, arguments, and options.
func NewFirewallAuthPortal(ctx *pulumi.Context,
	name string, args *FirewallAuthPortalArgs, opts ...pulumi.ResourceOption) (*FirewallAuthPortal, error) {
	if args == nil {
		args = &FirewallAuthPortalArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
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
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
	Groups []FirewallAuthPortalGroup `pulumi:"groups"`
	// Name of the identity-based route that applies to this portal.
	IdentityBasedRoute *string `pulumi:"identityBasedRoute"`
	// Address (or FQDN) of the authentication portal.
	PortalAddr *string `pulumi:"portalAddr"`
	// IPv6 address (or FQDN) of authentication portal.
	PortalAddr6 *string `pulumi:"portalAddr6"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallAuthPortalState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
	Groups FirewallAuthPortalGroupArrayInput
	// Name of the identity-based route that applies to this portal.
	IdentityBasedRoute pulumi.StringPtrInput
	// Address (or FQDN) of the authentication portal.
	PortalAddr pulumi.StringPtrInput
	// IPv6 address (or FQDN) of authentication portal.
	PortalAddr6 pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallAuthPortalState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAuthPortalState)(nil)).Elem()
}

type firewallAuthPortalArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
	Groups []FirewallAuthPortalGroup `pulumi:"groups"`
	// Name of the identity-based route that applies to this portal.
	IdentityBasedRoute *string `pulumi:"identityBasedRoute"`
	// Address (or FQDN) of the authentication portal.
	PortalAddr *string `pulumi:"portalAddr"`
	// IPv6 address (or FQDN) of authentication portal.
	PortalAddr6 *string `pulumi:"portalAddr6"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallAuthPortal resource.
type FirewallAuthPortalArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
	Groups FirewallAuthPortalGroupArrayInput
	// Name of the identity-based route that applies to this portal.
	IdentityBasedRoute pulumi.StringPtrInput
	// Address (or FQDN) of the authentication portal.
	PortalAddr pulumi.StringPtrInput
	// IPv6 address (or FQDN) of authentication portal.
	PortalAddr6 pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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
//          FirewallAuthPortalArray{ FirewallAuthPortalArgs{...} }
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
//          FirewallAuthPortalMap{ "key": FirewallAuthPortalArgs{...} }
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

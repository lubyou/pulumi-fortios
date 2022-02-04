// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 IP pools.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallIppool6(ctx, "trname", &fortios.FirewallIppool6Args{
// 			Endip:   pulumi.String("2001:3ca1:10f:1a:121b::19"),
// 			Startip: pulumi.String("2001:3ca1:10f:1a:121b::10"),
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
// Firewall Ippool6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallIppool6:FirewallIppool6 labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallIppool6 struct {
	pulumi.CustomResourceState

	// Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
	AddNat46Route pulumi.StringOutput `pulumi:"addNat46Route"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
	Endip pulumi.StringOutput `pulumi:"endip"`
	// IPv6 IP pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable NAT46. Valid values: `disable`, `enable`.
	Nat46 pulumi.StringOutput `pulumi:"nat46"`
	// First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
	Startip pulumi.StringOutput `pulumi:"startip"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallIppool6 registers a new resource with the given unique name, arguments, and options.
func NewFirewallIppool6(ctx *pulumi.Context,
	name string, args *FirewallIppool6Args, opts ...pulumi.ResourceOption) (*FirewallIppool6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endip == nil {
		return nil, errors.New("invalid value for required argument 'Endip'")
	}
	if args.Startip == nil {
		return nil, errors.New("invalid value for required argument 'Startip'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallIppool6
	err := ctx.RegisterResource("fortios:index/firewallIppool6:FirewallIppool6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallIppool6 gets an existing FirewallIppool6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallIppool6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallIppool6State, opts ...pulumi.ResourceOption) (*FirewallIppool6, error) {
	var resource FirewallIppool6
	err := ctx.ReadResource("fortios:index/firewallIppool6:FirewallIppool6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallIppool6 resources.
type firewallIppool6State struct {
	// Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
	AddNat46Route *string `pulumi:"addNat46Route"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
	Endip *string `pulumi:"endip"`
	// IPv6 IP pool name.
	Name *string `pulumi:"name"`
	// Enable/disable NAT46. Valid values: `disable`, `enable`.
	Nat46 *string `pulumi:"nat46"`
	// First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
	Startip *string `pulumi:"startip"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallIppool6State struct {
	// Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
	AddNat46Route pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
	Endip pulumi.StringPtrInput
	// IPv6 IP pool name.
	Name pulumi.StringPtrInput
	// Enable/disable NAT46. Valid values: `disable`, `enable`.
	Nat46 pulumi.StringPtrInput
	// First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
	Startip pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallIppool6State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIppool6State)(nil)).Elem()
}

type firewallIppool6Args struct {
	// Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
	AddNat46Route *string `pulumi:"addNat46Route"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
	Endip string `pulumi:"endip"`
	// IPv6 IP pool name.
	Name *string `pulumi:"name"`
	// Enable/disable NAT46. Valid values: `disable`, `enable`.
	Nat46 *string `pulumi:"nat46"`
	// First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
	Startip string `pulumi:"startip"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallIppool6 resource.
type FirewallIppool6Args struct {
	// Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
	AddNat46Route pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
	Endip pulumi.StringInput
	// IPv6 IP pool name.
	Name pulumi.StringPtrInput
	// Enable/disable NAT46. Valid values: `disable`, `enable`.
	Nat46 pulumi.StringPtrInput
	// First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
	Startip pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallIppool6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIppool6Args)(nil)).Elem()
}

type FirewallIppool6Input interface {
	pulumi.Input

	ToFirewallIppool6Output() FirewallIppool6Output
	ToFirewallIppool6OutputWithContext(ctx context.Context) FirewallIppool6Output
}

func (*FirewallIppool6) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIppool6)(nil)).Elem()
}

func (i *FirewallIppool6) ToFirewallIppool6Output() FirewallIppool6Output {
	return i.ToFirewallIppool6OutputWithContext(context.Background())
}

func (i *FirewallIppool6) ToFirewallIppool6OutputWithContext(ctx context.Context) FirewallIppool6Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIppool6Output)
}

// FirewallIppool6ArrayInput is an input type that accepts FirewallIppool6Array and FirewallIppool6ArrayOutput values.
// You can construct a concrete instance of `FirewallIppool6ArrayInput` via:
//
//          FirewallIppool6Array{ FirewallIppool6Args{...} }
type FirewallIppool6ArrayInput interface {
	pulumi.Input

	ToFirewallIppool6ArrayOutput() FirewallIppool6ArrayOutput
	ToFirewallIppool6ArrayOutputWithContext(context.Context) FirewallIppool6ArrayOutput
}

type FirewallIppool6Array []FirewallIppool6Input

func (FirewallIppool6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallIppool6)(nil)).Elem()
}

func (i FirewallIppool6Array) ToFirewallIppool6ArrayOutput() FirewallIppool6ArrayOutput {
	return i.ToFirewallIppool6ArrayOutputWithContext(context.Background())
}

func (i FirewallIppool6Array) ToFirewallIppool6ArrayOutputWithContext(ctx context.Context) FirewallIppool6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIppool6ArrayOutput)
}

// FirewallIppool6MapInput is an input type that accepts FirewallIppool6Map and FirewallIppool6MapOutput values.
// You can construct a concrete instance of `FirewallIppool6MapInput` via:
//
//          FirewallIppool6Map{ "key": FirewallIppool6Args{...} }
type FirewallIppool6MapInput interface {
	pulumi.Input

	ToFirewallIppool6MapOutput() FirewallIppool6MapOutput
	ToFirewallIppool6MapOutputWithContext(context.Context) FirewallIppool6MapOutput
}

type FirewallIppool6Map map[string]FirewallIppool6Input

func (FirewallIppool6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallIppool6)(nil)).Elem()
}

func (i FirewallIppool6Map) ToFirewallIppool6MapOutput() FirewallIppool6MapOutput {
	return i.ToFirewallIppool6MapOutputWithContext(context.Background())
}

func (i FirewallIppool6Map) ToFirewallIppool6MapOutputWithContext(ctx context.Context) FirewallIppool6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIppool6MapOutput)
}

type FirewallIppool6Output struct{ *pulumi.OutputState }

func (FirewallIppool6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIppool6)(nil)).Elem()
}

func (o FirewallIppool6Output) ToFirewallIppool6Output() FirewallIppool6Output {
	return o
}

func (o FirewallIppool6Output) ToFirewallIppool6OutputWithContext(ctx context.Context) FirewallIppool6Output {
	return o
}

type FirewallIppool6ArrayOutput struct{ *pulumi.OutputState }

func (FirewallIppool6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallIppool6)(nil)).Elem()
}

func (o FirewallIppool6ArrayOutput) ToFirewallIppool6ArrayOutput() FirewallIppool6ArrayOutput {
	return o
}

func (o FirewallIppool6ArrayOutput) ToFirewallIppool6ArrayOutputWithContext(ctx context.Context) FirewallIppool6ArrayOutput {
	return o
}

func (o FirewallIppool6ArrayOutput) Index(i pulumi.IntInput) FirewallIppool6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallIppool6 {
		return vs[0].([]*FirewallIppool6)[vs[1].(int)]
	}).(FirewallIppool6Output)
}

type FirewallIppool6MapOutput struct{ *pulumi.OutputState }

func (FirewallIppool6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallIppool6)(nil)).Elem()
}

func (o FirewallIppool6MapOutput) ToFirewallIppool6MapOutput() FirewallIppool6MapOutput {
	return o
}

func (o FirewallIppool6MapOutput) ToFirewallIppool6MapOutputWithContext(ctx context.Context) FirewallIppool6MapOutput {
	return o
}

func (o FirewallIppool6MapOutput) MapIndex(k pulumi.StringInput) FirewallIppool6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallIppool6 {
		return vs[0].(map[string]*FirewallIppool6)[vs[1].(string)]
	}).(FirewallIppool6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIppool6Input)(nil)).Elem(), &FirewallIppool6{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIppool6ArrayInput)(nil)).Elem(), FirewallIppool6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIppool6MapInput)(nil)).Elem(), FirewallIppool6Map{})
	pulumi.RegisterOutputType(FirewallIppool6Output{})
	pulumi.RegisterOutputType(FirewallIppool6ArrayOutput{})
	pulumi.RegisterOutputType(FirewallIppool6MapOutput{})
}

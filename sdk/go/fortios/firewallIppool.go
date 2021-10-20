// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 IP pools.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallIppool(ctx, "trname", &fortios.FirewallIppoolArgs{
// 			ArpReply:         pulumi.String("enable"),
// 			BlockSize:        pulumi.Int(128),
// 			Endip:            pulumi.String("1.0.0.20"),
// 			NumBlocksPerUser: pulumi.Int(8),
// 			PbaTimeout:       pulumi.Int(30),
// 			PermitAnyHost:    pulumi.String("disable"),
// 			SourceEndip:      pulumi.String("0.0.0.0"),
// 			SourceStartip:    pulumi.String("0.0.0.0"),
// 			Startip:          pulumi.String("1.0.0.0"),
// 			Type:             pulumi.String("overload"),
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
// Firewall Ippool can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallIppool:FirewallIppool labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallIppool struct {
	pulumi.CustomResourceState

	// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
	ArpIntf pulumi.StringOutput `pulumi:"arpIntf"`
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
	ArpReply pulumi.StringOutput `pulumi:"arpReply"`
	// Associated interface name.
	AssociatedInterface pulumi.StringOutput `pulumi:"associatedInterface"`
	// Number of addresses in a block (64 to 4096, default = 128).
	BlockSize pulumi.IntOutput `pulumi:"blockSize"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Endip pulumi.StringOutput `pulumi:"endip"`
	// Final port number (inclusive) in the range for the address pool (Default: 65533).
	Endport pulumi.IntOutput `pulumi:"endport"`
	// IP pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
	NumBlocksPerUser pulumi.IntOutput `pulumi:"numBlocksPerUser"`
	// Port block allocation timeout (seconds).
	PbaTimeout pulumi.IntOutput `pulumi:"pbaTimeout"`
	// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
	PermitAnyHost pulumi.StringOutput `pulumi:"permitAnyHost"`
	// Number of port for each user (32 to 60416, default = 0, auto).
	PortPerUser pulumi.IntOutput `pulumi:"portPerUser"`
	// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceEndip pulumi.StringOutput `pulumi:"sourceEndip"`
	// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceStartip pulumi.StringOutput `pulumi:"sourceStartip"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Startip pulumi.StringOutput `pulumi:"startip"`
	// First port number (inclusive) in the range for the address pool (Default: 5117).
	Startport pulumi.IntOutput `pulumi:"startport"`
	// IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallIppool registers a new resource with the given unique name, arguments, and options.
func NewFirewallIppool(ctx *pulumi.Context,
	name string, args *FirewallIppoolArgs, opts ...pulumi.ResourceOption) (*FirewallIppool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endip == nil {
		return nil, errors.New("invalid value for required argument 'Endip'")
	}
	if args.Startip == nil {
		return nil, errors.New("invalid value for required argument 'Startip'")
	}
	var resource FirewallIppool
	err := ctx.RegisterResource("fortios:index/firewallIppool:FirewallIppool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallIppool gets an existing FirewallIppool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallIppool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallIppoolState, opts ...pulumi.ResourceOption) (*FirewallIppool, error) {
	var resource FirewallIppool
	err := ctx.ReadResource("fortios:index/firewallIppool:FirewallIppool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallIppool resources.
type firewallIppoolState struct {
	// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
	ArpIntf *string `pulumi:"arpIntf"`
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
	ArpReply *string `pulumi:"arpReply"`
	// Associated interface name.
	AssociatedInterface *string `pulumi:"associatedInterface"`
	// Number of addresses in a block (64 to 4096, default = 128).
	BlockSize *int `pulumi:"blockSize"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Endip *string `pulumi:"endip"`
	// Final port number (inclusive) in the range for the address pool (Default: 65533).
	Endport *int `pulumi:"endport"`
	// IP pool name.
	Name *string `pulumi:"name"`
	// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
	NumBlocksPerUser *int `pulumi:"numBlocksPerUser"`
	// Port block allocation timeout (seconds).
	PbaTimeout *int `pulumi:"pbaTimeout"`
	// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
	PermitAnyHost *string `pulumi:"permitAnyHost"`
	// Number of port for each user (32 to 60416, default = 0, auto).
	PortPerUser *int `pulumi:"portPerUser"`
	// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceEndip *string `pulumi:"sourceEndip"`
	// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceStartip *string `pulumi:"sourceStartip"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Startip *string `pulumi:"startip"`
	// First port number (inclusive) in the range for the address pool (Default: 5117).
	Startport *int `pulumi:"startport"`
	// IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallIppoolState struct {
	// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
	ArpIntf pulumi.StringPtrInput
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
	ArpReply pulumi.StringPtrInput
	// Associated interface name.
	AssociatedInterface pulumi.StringPtrInput
	// Number of addresses in a block (64 to 4096, default = 128).
	BlockSize pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Endip pulumi.StringPtrInput
	// Final port number (inclusive) in the range for the address pool (Default: 65533).
	Endport pulumi.IntPtrInput
	// IP pool name.
	Name pulumi.StringPtrInput
	// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
	NumBlocksPerUser pulumi.IntPtrInput
	// Port block allocation timeout (seconds).
	PbaTimeout pulumi.IntPtrInput
	// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
	PermitAnyHost pulumi.StringPtrInput
	// Number of port for each user (32 to 60416, default = 0, auto).
	PortPerUser pulumi.IntPtrInput
	// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceEndip pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceStartip pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Startip pulumi.StringPtrInput
	// First port number (inclusive) in the range for the address pool (Default: 5117).
	Startport pulumi.IntPtrInput
	// IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallIppoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIppoolState)(nil)).Elem()
}

type firewallIppoolArgs struct {
	// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
	ArpIntf *string `pulumi:"arpIntf"`
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
	ArpReply *string `pulumi:"arpReply"`
	// Associated interface name.
	AssociatedInterface *string `pulumi:"associatedInterface"`
	// Number of addresses in a block (64 to 4096, default = 128).
	BlockSize *int `pulumi:"blockSize"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Endip string `pulumi:"endip"`
	// Final port number (inclusive) in the range for the address pool (Default: 65533).
	Endport *int `pulumi:"endport"`
	// IP pool name.
	Name *string `pulumi:"name"`
	// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
	NumBlocksPerUser *int `pulumi:"numBlocksPerUser"`
	// Port block allocation timeout (seconds).
	PbaTimeout *int `pulumi:"pbaTimeout"`
	// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
	PermitAnyHost *string `pulumi:"permitAnyHost"`
	// Number of port for each user (32 to 60416, default = 0, auto).
	PortPerUser *int `pulumi:"portPerUser"`
	// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceEndip *string `pulumi:"sourceEndip"`
	// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceStartip *string `pulumi:"sourceStartip"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Startip string `pulumi:"startip"`
	// First port number (inclusive) in the range for the address pool (Default: 5117).
	Startport *int `pulumi:"startport"`
	// IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallIppool resource.
type FirewallIppoolArgs struct {
	// Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
	ArpIntf pulumi.StringPtrInput
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
	ArpReply pulumi.StringPtrInput
	// Associated interface name.
	AssociatedInterface pulumi.StringPtrInput
	// Number of addresses in a block (64 to 4096, default = 128).
	BlockSize pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Endip pulumi.StringInput
	// Final port number (inclusive) in the range for the address pool (Default: 65533).
	Endport pulumi.IntPtrInput
	// IP pool name.
	Name pulumi.StringPtrInput
	// Number of addresses blocks that can be used by a user (1 to 128, default = 8).
	NumBlocksPerUser pulumi.IntPtrInput
	// Port block allocation timeout (seconds).
	PbaTimeout pulumi.IntPtrInput
	// Enable/disable full cone NAT. Valid values: `disable`, `enable`.
	PermitAnyHost pulumi.StringPtrInput
	// Number of port for each user (32 to 60416, default = 0, auto).
	PortPerUser pulumi.IntPtrInput
	// Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceEndip pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	SourceStartip pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
	Startip pulumi.StringInput
	// First port number (inclusive) in the range for the address pool (Default: 5117).
	Startport pulumi.IntPtrInput
	// IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallIppoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIppoolArgs)(nil)).Elem()
}

type FirewallIppoolInput interface {
	pulumi.Input

	ToFirewallIppoolOutput() FirewallIppoolOutput
	ToFirewallIppoolOutputWithContext(ctx context.Context) FirewallIppoolOutput
}

func (*FirewallIppool) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallIppool)(nil))
}

func (i *FirewallIppool) ToFirewallIppoolOutput() FirewallIppoolOutput {
	return i.ToFirewallIppoolOutputWithContext(context.Background())
}

func (i *FirewallIppool) ToFirewallIppoolOutputWithContext(ctx context.Context) FirewallIppoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIppoolOutput)
}

func (i *FirewallIppool) ToFirewallIppoolPtrOutput() FirewallIppoolPtrOutput {
	return i.ToFirewallIppoolPtrOutputWithContext(context.Background())
}

func (i *FirewallIppool) ToFirewallIppoolPtrOutputWithContext(ctx context.Context) FirewallIppoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIppoolPtrOutput)
}

type FirewallIppoolPtrInput interface {
	pulumi.Input

	ToFirewallIppoolPtrOutput() FirewallIppoolPtrOutput
	ToFirewallIppoolPtrOutputWithContext(ctx context.Context) FirewallIppoolPtrOutput
}

type firewallIppoolPtrType FirewallIppoolArgs

func (*firewallIppoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIppool)(nil))
}

func (i *firewallIppoolPtrType) ToFirewallIppoolPtrOutput() FirewallIppoolPtrOutput {
	return i.ToFirewallIppoolPtrOutputWithContext(context.Background())
}

func (i *firewallIppoolPtrType) ToFirewallIppoolPtrOutputWithContext(ctx context.Context) FirewallIppoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIppoolPtrOutput)
}

// FirewallIppoolArrayInput is an input type that accepts FirewallIppoolArray and FirewallIppoolArrayOutput values.
// You can construct a concrete instance of `FirewallIppoolArrayInput` via:
//
//          FirewallIppoolArray{ FirewallIppoolArgs{...} }
type FirewallIppoolArrayInput interface {
	pulumi.Input

	ToFirewallIppoolArrayOutput() FirewallIppoolArrayOutput
	ToFirewallIppoolArrayOutputWithContext(context.Context) FirewallIppoolArrayOutput
}

type FirewallIppoolArray []FirewallIppoolInput

func (FirewallIppoolArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallIppool)(nil))
}

func (i FirewallIppoolArray) ToFirewallIppoolArrayOutput() FirewallIppoolArrayOutput {
	return i.ToFirewallIppoolArrayOutputWithContext(context.Background())
}

func (i FirewallIppoolArray) ToFirewallIppoolArrayOutputWithContext(ctx context.Context) FirewallIppoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIppoolArrayOutput)
}

// FirewallIppoolMapInput is an input type that accepts FirewallIppoolMap and FirewallIppoolMapOutput values.
// You can construct a concrete instance of `FirewallIppoolMapInput` via:
//
//          FirewallIppoolMap{ "key": FirewallIppoolArgs{...} }
type FirewallIppoolMapInput interface {
	pulumi.Input

	ToFirewallIppoolMapOutput() FirewallIppoolMapOutput
	ToFirewallIppoolMapOutputWithContext(context.Context) FirewallIppoolMapOutput
}

type FirewallIppoolMap map[string]FirewallIppoolInput

func (FirewallIppoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallIppool)(nil))
}

func (i FirewallIppoolMap) ToFirewallIppoolMapOutput() FirewallIppoolMapOutput {
	return i.ToFirewallIppoolMapOutputWithContext(context.Background())
}

func (i FirewallIppoolMap) ToFirewallIppoolMapOutputWithContext(ctx context.Context) FirewallIppoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIppoolMapOutput)
}

type FirewallIppoolOutput struct {
	*pulumi.OutputState
}

func (FirewallIppoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallIppool)(nil))
}

func (o FirewallIppoolOutput) ToFirewallIppoolOutput() FirewallIppoolOutput {
	return o
}

func (o FirewallIppoolOutput) ToFirewallIppoolOutputWithContext(ctx context.Context) FirewallIppoolOutput {
	return o
}

func (o FirewallIppoolOutput) ToFirewallIppoolPtrOutput() FirewallIppoolPtrOutput {
	return o.ToFirewallIppoolPtrOutputWithContext(context.Background())
}

func (o FirewallIppoolOutput) ToFirewallIppoolPtrOutputWithContext(ctx context.Context) FirewallIppoolPtrOutput {
	return o.ApplyT(func(v FirewallIppool) *FirewallIppool {
		return &v
	}).(FirewallIppoolPtrOutput)
}

type FirewallIppoolPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallIppoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIppool)(nil))
}

func (o FirewallIppoolPtrOutput) ToFirewallIppoolPtrOutput() FirewallIppoolPtrOutput {
	return o
}

func (o FirewallIppoolPtrOutput) ToFirewallIppoolPtrOutputWithContext(ctx context.Context) FirewallIppoolPtrOutput {
	return o
}

type FirewallIppoolArrayOutput struct{ *pulumi.OutputState }

func (FirewallIppoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallIppool)(nil))
}

func (o FirewallIppoolArrayOutput) ToFirewallIppoolArrayOutput() FirewallIppoolArrayOutput {
	return o
}

func (o FirewallIppoolArrayOutput) ToFirewallIppoolArrayOutputWithContext(ctx context.Context) FirewallIppoolArrayOutput {
	return o
}

func (o FirewallIppoolArrayOutput) Index(i pulumi.IntInput) FirewallIppoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallIppool {
		return vs[0].([]FirewallIppool)[vs[1].(int)]
	}).(FirewallIppoolOutput)
}

type FirewallIppoolMapOutput struct{ *pulumi.OutputState }

func (FirewallIppoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallIppool)(nil))
}

func (o FirewallIppoolMapOutput) ToFirewallIppoolMapOutput() FirewallIppoolMapOutput {
	return o
}

func (o FirewallIppoolMapOutput) ToFirewallIppoolMapOutputWithContext(ctx context.Context) FirewallIppoolMapOutput {
	return o
}

func (o FirewallIppoolMapOutput) MapIndex(k pulumi.StringInput) FirewallIppoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallIppool {
		return vs[0].(map[string]FirewallIppool)[vs[1].(string)]
	}).(FirewallIppoolOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallIppoolOutput{})
	pulumi.RegisterOutputType(FirewallIppoolPtrOutput{})
	pulumi.RegisterOutputType(FirewallIppoolArrayOutput{})
	pulumi.RegisterOutputType(FirewallIppoolMapOutput{})
}
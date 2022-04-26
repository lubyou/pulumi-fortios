// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IP to MAC address pairs in the IP/MAC binding table.
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
// 		_, err := fortios.NewFirewallIpmacbindingTable(ctx, "trname", &fortios.FirewallIpmacbindingTableArgs{
// 			Ip:     pulumi.String("1.1.1.1"),
// 			Mac:    pulumi.String("00:01:6c:06:a6:29"),
// 			SeqNum: pulumi.Int(1),
// 			Status: pulumi.String("disable"),
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
// FirewallIpmacbinding Table can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallIpmacbindingTable:FirewallIpmacbindingTable labelname {{seq_num}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallIpmacbindingTable:FirewallIpmacbindingTable labelname {{seq_num}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallIpmacbindingTable struct {
	pulumi.CustomResourceState

	// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
	Ip pulumi.StringOutput `pulumi:"ip"`
	// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Name of the pair (optional, default = no name).
	Name pulumi.StringOutput `pulumi:"name"`
	// Entry number.
	SeqNum pulumi.IntOutput `pulumi:"seqNum"`
	// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallIpmacbindingTable registers a new resource with the given unique name, arguments, and options.
func NewFirewallIpmacbindingTable(ctx *pulumi.Context,
	name string, args *FirewallIpmacbindingTableArgs, opts ...pulumi.ResourceOption) (*FirewallIpmacbindingTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallIpmacbindingTable
	err := ctx.RegisterResource("fortios:index/firewallIpmacbindingTable:FirewallIpmacbindingTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallIpmacbindingTable gets an existing FirewallIpmacbindingTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallIpmacbindingTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallIpmacbindingTableState, opts ...pulumi.ResourceOption) (*FirewallIpmacbindingTable, error) {
	var resource FirewallIpmacbindingTable
	err := ctx.ReadResource("fortios:index/firewallIpmacbindingTable:FirewallIpmacbindingTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallIpmacbindingTable resources.
type firewallIpmacbindingTableState struct {
	// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
	Ip *string `pulumi:"ip"`
	// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
	Mac *string `pulumi:"mac"`
	// Name of the pair (optional, default = no name).
	Name *string `pulumi:"name"`
	// Entry number.
	SeqNum *int `pulumi:"seqNum"`
	// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallIpmacbindingTableState struct {
	// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
	Ip pulumi.StringPtrInput
	// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
	Mac pulumi.StringPtrInput
	// Name of the pair (optional, default = no name).
	Name pulumi.StringPtrInput
	// Entry number.
	SeqNum pulumi.IntPtrInput
	// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallIpmacbindingTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIpmacbindingTableState)(nil)).Elem()
}

type firewallIpmacbindingTableArgs struct {
	// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
	Ip string `pulumi:"ip"`
	// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
	Mac *string `pulumi:"mac"`
	// Name of the pair (optional, default = no name).
	Name *string `pulumi:"name"`
	// Entry number.
	SeqNum *int `pulumi:"seqNum"`
	// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallIpmacbindingTable resource.
type FirewallIpmacbindingTableArgs struct {
	// IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
	Ip pulumi.StringInput
	// MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
	Mac pulumi.StringPtrInput
	// Name of the pair (optional, default = no name).
	Name pulumi.StringPtrInput
	// Entry number.
	SeqNum pulumi.IntPtrInput
	// Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallIpmacbindingTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIpmacbindingTableArgs)(nil)).Elem()
}

type FirewallIpmacbindingTableInput interface {
	pulumi.Input

	ToFirewallIpmacbindingTableOutput() FirewallIpmacbindingTableOutput
	ToFirewallIpmacbindingTableOutputWithContext(ctx context.Context) FirewallIpmacbindingTableOutput
}

func (*FirewallIpmacbindingTable) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIpmacbindingTable)(nil)).Elem()
}

func (i *FirewallIpmacbindingTable) ToFirewallIpmacbindingTableOutput() FirewallIpmacbindingTableOutput {
	return i.ToFirewallIpmacbindingTableOutputWithContext(context.Background())
}

func (i *FirewallIpmacbindingTable) ToFirewallIpmacbindingTableOutputWithContext(ctx context.Context) FirewallIpmacbindingTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIpmacbindingTableOutput)
}

// FirewallIpmacbindingTableArrayInput is an input type that accepts FirewallIpmacbindingTableArray and FirewallIpmacbindingTableArrayOutput values.
// You can construct a concrete instance of `FirewallIpmacbindingTableArrayInput` via:
//
//          FirewallIpmacbindingTableArray{ FirewallIpmacbindingTableArgs{...} }
type FirewallIpmacbindingTableArrayInput interface {
	pulumi.Input

	ToFirewallIpmacbindingTableArrayOutput() FirewallIpmacbindingTableArrayOutput
	ToFirewallIpmacbindingTableArrayOutputWithContext(context.Context) FirewallIpmacbindingTableArrayOutput
}

type FirewallIpmacbindingTableArray []FirewallIpmacbindingTableInput

func (FirewallIpmacbindingTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallIpmacbindingTable)(nil)).Elem()
}

func (i FirewallIpmacbindingTableArray) ToFirewallIpmacbindingTableArrayOutput() FirewallIpmacbindingTableArrayOutput {
	return i.ToFirewallIpmacbindingTableArrayOutputWithContext(context.Background())
}

func (i FirewallIpmacbindingTableArray) ToFirewallIpmacbindingTableArrayOutputWithContext(ctx context.Context) FirewallIpmacbindingTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIpmacbindingTableArrayOutput)
}

// FirewallIpmacbindingTableMapInput is an input type that accepts FirewallIpmacbindingTableMap and FirewallIpmacbindingTableMapOutput values.
// You can construct a concrete instance of `FirewallIpmacbindingTableMapInput` via:
//
//          FirewallIpmacbindingTableMap{ "key": FirewallIpmacbindingTableArgs{...} }
type FirewallIpmacbindingTableMapInput interface {
	pulumi.Input

	ToFirewallIpmacbindingTableMapOutput() FirewallIpmacbindingTableMapOutput
	ToFirewallIpmacbindingTableMapOutputWithContext(context.Context) FirewallIpmacbindingTableMapOutput
}

type FirewallIpmacbindingTableMap map[string]FirewallIpmacbindingTableInput

func (FirewallIpmacbindingTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallIpmacbindingTable)(nil)).Elem()
}

func (i FirewallIpmacbindingTableMap) ToFirewallIpmacbindingTableMapOutput() FirewallIpmacbindingTableMapOutput {
	return i.ToFirewallIpmacbindingTableMapOutputWithContext(context.Background())
}

func (i FirewallIpmacbindingTableMap) ToFirewallIpmacbindingTableMapOutputWithContext(ctx context.Context) FirewallIpmacbindingTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIpmacbindingTableMapOutput)
}

type FirewallIpmacbindingTableOutput struct{ *pulumi.OutputState }

func (FirewallIpmacbindingTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIpmacbindingTable)(nil)).Elem()
}

func (o FirewallIpmacbindingTableOutput) ToFirewallIpmacbindingTableOutput() FirewallIpmacbindingTableOutput {
	return o
}

func (o FirewallIpmacbindingTableOutput) ToFirewallIpmacbindingTableOutputWithContext(ctx context.Context) FirewallIpmacbindingTableOutput {
	return o
}

type FirewallIpmacbindingTableArrayOutput struct{ *pulumi.OutputState }

func (FirewallIpmacbindingTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallIpmacbindingTable)(nil)).Elem()
}

func (o FirewallIpmacbindingTableArrayOutput) ToFirewallIpmacbindingTableArrayOutput() FirewallIpmacbindingTableArrayOutput {
	return o
}

func (o FirewallIpmacbindingTableArrayOutput) ToFirewallIpmacbindingTableArrayOutputWithContext(ctx context.Context) FirewallIpmacbindingTableArrayOutput {
	return o
}

func (o FirewallIpmacbindingTableArrayOutput) Index(i pulumi.IntInput) FirewallIpmacbindingTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallIpmacbindingTable {
		return vs[0].([]*FirewallIpmacbindingTable)[vs[1].(int)]
	}).(FirewallIpmacbindingTableOutput)
}

type FirewallIpmacbindingTableMapOutput struct{ *pulumi.OutputState }

func (FirewallIpmacbindingTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallIpmacbindingTable)(nil)).Elem()
}

func (o FirewallIpmacbindingTableMapOutput) ToFirewallIpmacbindingTableMapOutput() FirewallIpmacbindingTableMapOutput {
	return o
}

func (o FirewallIpmacbindingTableMapOutput) ToFirewallIpmacbindingTableMapOutputWithContext(ctx context.Context) FirewallIpmacbindingTableMapOutput {
	return o
}

func (o FirewallIpmacbindingTableMapOutput) MapIndex(k pulumi.StringInput) FirewallIpmacbindingTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallIpmacbindingTable {
		return vs[0].(map[string]*FirewallIpmacbindingTable)[vs[1].(string)]
	}).(FirewallIpmacbindingTableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIpmacbindingTableInput)(nil)).Elem(), &FirewallIpmacbindingTable{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIpmacbindingTableArrayInput)(nil)).Elem(), FirewallIpmacbindingTableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIpmacbindingTableMapInput)(nil)).Elem(), FirewallIpmacbindingTableMap{})
	pulumi.RegisterOutputType(FirewallIpmacbindingTableOutput{})
	pulumi.RegisterOutputType(FirewallIpmacbindingTableArrayOutput{})
	pulumi.RegisterOutputType(FirewallIpmacbindingTableMapOutput{})
}

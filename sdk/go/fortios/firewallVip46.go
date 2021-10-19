// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 to IPv6 virtual IPs.
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
// 		_, err := fortios.NewFirewallVip46(ctx, "trname", &fortios.FirewallVip46Args{
// 			ArpReply:    pulumi.String("enable"),
// 			Color:       pulumi.Int(0),
// 			Extip:       pulumi.String("10.202.1.200"),
// 			Extport:     pulumi.String("0-65535"),
// 			Fosid:       pulumi.Int(0),
// 			LdbMethod:   pulumi.String("static"),
// 			Mappedip:    pulumi.String("2001:1:1:2::200"),
// 			Mappedport:  pulumi.String("0-65535"),
// 			Portforward: pulumi.String("disable"),
// 			Protocol:    pulumi.String("tcp"),
// 			Type:        pulumi.String("static-nat"),
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
// Firewall Vip46 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallVip46:FirewallVip46 labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallVip46 struct {
	pulumi.CustomResourceState

	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply pulumi.StringOutput `pulumi:"arpReply"`
	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Start-external-IP [-end-external-IP].
	Extip pulumi.StringOutput `pulumi:"extip"`
	// External service port.
	Extport pulumi.StringOutput `pulumi:"extport"`
	// Custom defined id.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod pulumi.StringOutput `pulumi:"ldbMethod"`
	// Start-mapped-IP [-end mapped-IP].
	Mappedip pulumi.StringOutput `pulumi:"mappedip"`
	// Mapped service port.
	Mappedport pulumi.StringOutput `pulumi:"mappedport"`
	// Health monitors.
	Monitors FirewallVip46MonitorArrayOutput `pulumi:"monitors"`
	// Health monitor name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward pulumi.StringOutput `pulumi:"portforward"`
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Real servers. The structure of `realservers` block is documented below.
	Realservers FirewallVip46RealserverArrayOutput `pulumi:"realservers"`
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType pulumi.StringOutput `pulumi:"serverType"`
	// Source IP filter (x.x.x.x/x). The structure of `srcFilter` block is documented below.
	SrcFilters FirewallVip46SrcFilterArrayOutput `pulumi:"srcFilters"`
	// Interfaces to which the VIP46 applies. Separate the names with spaces. The structure of `srcintfFilter` block is documented below.
	SrcintfFilters FirewallVip46SrcintfFilterArrayOutput `pulumi:"srcintfFilters"`
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallVip46 registers a new resource with the given unique name, arguments, and options.
func NewFirewallVip46(ctx *pulumi.Context,
	name string, args *FirewallVip46Args, opts ...pulumi.ResourceOption) (*FirewallVip46, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Extip == nil {
		return nil, errors.New("invalid value for required argument 'Extip'")
	}
	if args.Mappedip == nil {
		return nil, errors.New("invalid value for required argument 'Mappedip'")
	}
	var resource FirewallVip46
	err := ctx.RegisterResource("fortios:index/firewallVip46:FirewallVip46", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallVip46 gets an existing FirewallVip46 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallVip46(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallVip46State, opts ...pulumi.ResourceOption) (*FirewallVip46, error) {
	var resource FirewallVip46
	err := ctx.ReadResource("fortios:index/firewallVip46:FirewallVip46", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallVip46 resources.
type firewallVip46State struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply *string `pulumi:"arpReply"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Start-external-IP [-end-external-IP].
	Extip *string `pulumi:"extip"`
	// External service port.
	Extport *string `pulumi:"extport"`
	// Custom defined id.
	Fosid *int `pulumi:"fosid"`
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Start-mapped-IP [-end mapped-IP].
	Mappedip *string `pulumi:"mappedip"`
	// Mapped service port.
	Mappedport *string `pulumi:"mappedport"`
	// Health monitors.
	Monitors []FirewallVip46Monitor `pulumi:"monitors"`
	// Health monitor name.
	Name *string `pulumi:"name"`
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward *string `pulumi:"portforward"`
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol *string `pulumi:"protocol"`
	// Real servers. The structure of `realservers` block is documented below.
	Realservers []FirewallVip46Realserver `pulumi:"realservers"`
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType *string `pulumi:"serverType"`
	// Source IP filter (x.x.x.x/x). The structure of `srcFilter` block is documented below.
	SrcFilters []FirewallVip46SrcFilter `pulumi:"srcFilters"`
	// Interfaces to which the VIP46 applies. Separate the names with spaces. The structure of `srcintfFilter` block is documented below.
	SrcintfFilters []FirewallVip46SrcintfFilter `pulumi:"srcintfFilters"`
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallVip46State struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Start-external-IP [-end-external-IP].
	Extip pulumi.StringPtrInput
	// External service port.
	Extport pulumi.StringPtrInput
	// Custom defined id.
	Fosid pulumi.IntPtrInput
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod pulumi.StringPtrInput
	// Start-mapped-IP [-end mapped-IP].
	Mappedip pulumi.StringPtrInput
	// Mapped service port.
	Mappedport pulumi.StringPtrInput
	// Health monitors.
	Monitors FirewallVip46MonitorArrayInput
	// Health monitor name.
	Name pulumi.StringPtrInput
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward pulumi.StringPtrInput
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol pulumi.StringPtrInput
	// Real servers. The structure of `realservers` block is documented below.
	Realservers FirewallVip46RealserverArrayInput
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType pulumi.StringPtrInput
	// Source IP filter (x.x.x.x/x). The structure of `srcFilter` block is documented below.
	SrcFilters FirewallVip46SrcFilterArrayInput
	// Interfaces to which the VIP46 applies. Separate the names with spaces. The structure of `srcintfFilter` block is documented below.
	SrcintfFilters FirewallVip46SrcintfFilterArrayInput
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallVip46State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVip46State)(nil)).Elem()
}

type firewallVip46Args struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply *string `pulumi:"arpReply"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Start-external-IP [-end-external-IP].
	Extip string `pulumi:"extip"`
	// External service port.
	Extport *string `pulumi:"extport"`
	// Custom defined id.
	Fosid *int `pulumi:"fosid"`
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod *string `pulumi:"ldbMethod"`
	// Start-mapped-IP [-end mapped-IP].
	Mappedip string `pulumi:"mappedip"`
	// Mapped service port.
	Mappedport *string `pulumi:"mappedport"`
	// Health monitors.
	Monitors []FirewallVip46Monitor `pulumi:"monitors"`
	// Health monitor name.
	Name *string `pulumi:"name"`
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward *string `pulumi:"portforward"`
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol *string `pulumi:"protocol"`
	// Real servers. The structure of `realservers` block is documented below.
	Realservers []FirewallVip46Realserver `pulumi:"realservers"`
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType *string `pulumi:"serverType"`
	// Source IP filter (x.x.x.x/x). The structure of `srcFilter` block is documented below.
	SrcFilters []FirewallVip46SrcFilter `pulumi:"srcFilters"`
	// Interfaces to which the VIP46 applies. Separate the names with spaces. The structure of `srcintfFilter` block is documented below.
	SrcintfFilters []FirewallVip46SrcintfFilter `pulumi:"srcintfFilters"`
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallVip46 resource.
type FirewallVip46Args struct {
	// Enable ARP reply. Valid values: `disable`, `enable`.
	ArpReply pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Start-external-IP [-end-external-IP].
	Extip pulumi.StringInput
	// External service port.
	Extport pulumi.StringPtrInput
	// Custom defined id.
	Fosid pulumi.IntPtrInput
	// Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
	LdbMethod pulumi.StringPtrInput
	// Start-mapped-IP [-end mapped-IP].
	Mappedip pulumi.StringInput
	// Mapped service port.
	Mappedport pulumi.StringPtrInput
	// Health monitors.
	Monitors FirewallVip46MonitorArrayInput
	// Health monitor name.
	Name pulumi.StringPtrInput
	// Enable port forwarding. Valid values: `disable`, `enable`.
	Portforward pulumi.StringPtrInput
	// Mapped port protocol. Valid values: `tcp`, `udp`.
	Protocol pulumi.StringPtrInput
	// Real servers. The structure of `realservers` block is documented below.
	Realservers FirewallVip46RealserverArrayInput
	// Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
	ServerType pulumi.StringPtrInput
	// Source IP filter (x.x.x.x/x). The structure of `srcFilter` block is documented below.
	SrcFilters FirewallVip46SrcFilterArrayInput
	// Interfaces to which the VIP46 applies. Separate the names with spaces. The structure of `srcintfFilter` block is documented below.
	SrcintfFilters FirewallVip46SrcintfFilterArrayInput
	// VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallVip46Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVip46Args)(nil)).Elem()
}

type FirewallVip46Input interface {
	pulumi.Input

	ToFirewallVip46Output() FirewallVip46Output
	ToFirewallVip46OutputWithContext(ctx context.Context) FirewallVip46Output
}

func (*FirewallVip46) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallVip46)(nil))
}

func (i *FirewallVip46) ToFirewallVip46Output() FirewallVip46Output {
	return i.ToFirewallVip46OutputWithContext(context.Background())
}

func (i *FirewallVip46) ToFirewallVip46OutputWithContext(ctx context.Context) FirewallVip46Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip46Output)
}

func (i *FirewallVip46) ToFirewallVip46PtrOutput() FirewallVip46PtrOutput {
	return i.ToFirewallVip46PtrOutputWithContext(context.Background())
}

func (i *FirewallVip46) ToFirewallVip46PtrOutputWithContext(ctx context.Context) FirewallVip46PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip46PtrOutput)
}

type FirewallVip46PtrInput interface {
	pulumi.Input

	ToFirewallVip46PtrOutput() FirewallVip46PtrOutput
	ToFirewallVip46PtrOutputWithContext(ctx context.Context) FirewallVip46PtrOutput
}

type firewallVip46PtrType FirewallVip46Args

func (*firewallVip46PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVip46)(nil))
}

func (i *firewallVip46PtrType) ToFirewallVip46PtrOutput() FirewallVip46PtrOutput {
	return i.ToFirewallVip46PtrOutputWithContext(context.Background())
}

func (i *firewallVip46PtrType) ToFirewallVip46PtrOutputWithContext(ctx context.Context) FirewallVip46PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip46PtrOutput)
}

// FirewallVip46ArrayInput is an input type that accepts FirewallVip46Array and FirewallVip46ArrayOutput values.
// You can construct a concrete instance of `FirewallVip46ArrayInput` via:
//
//          FirewallVip46Array{ FirewallVip46Args{...} }
type FirewallVip46ArrayInput interface {
	pulumi.Input

	ToFirewallVip46ArrayOutput() FirewallVip46ArrayOutput
	ToFirewallVip46ArrayOutputWithContext(context.Context) FirewallVip46ArrayOutput
}

type FirewallVip46Array []FirewallVip46Input

func (FirewallVip46Array) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallVip46)(nil))
}

func (i FirewallVip46Array) ToFirewallVip46ArrayOutput() FirewallVip46ArrayOutput {
	return i.ToFirewallVip46ArrayOutputWithContext(context.Background())
}

func (i FirewallVip46Array) ToFirewallVip46ArrayOutputWithContext(ctx context.Context) FirewallVip46ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip46ArrayOutput)
}

// FirewallVip46MapInput is an input type that accepts FirewallVip46Map and FirewallVip46MapOutput values.
// You can construct a concrete instance of `FirewallVip46MapInput` via:
//
//          FirewallVip46Map{ "key": FirewallVip46Args{...} }
type FirewallVip46MapInput interface {
	pulumi.Input

	ToFirewallVip46MapOutput() FirewallVip46MapOutput
	ToFirewallVip46MapOutputWithContext(context.Context) FirewallVip46MapOutput
}

type FirewallVip46Map map[string]FirewallVip46Input

func (FirewallVip46Map) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallVip46)(nil))
}

func (i FirewallVip46Map) ToFirewallVip46MapOutput() FirewallVip46MapOutput {
	return i.ToFirewallVip46MapOutputWithContext(context.Background())
}

func (i FirewallVip46Map) ToFirewallVip46MapOutputWithContext(ctx context.Context) FirewallVip46MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip46MapOutput)
}

type FirewallVip46Output struct {
	*pulumi.OutputState
}

func (FirewallVip46Output) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallVip46)(nil))
}

func (o FirewallVip46Output) ToFirewallVip46Output() FirewallVip46Output {
	return o
}

func (o FirewallVip46Output) ToFirewallVip46OutputWithContext(ctx context.Context) FirewallVip46Output {
	return o
}

func (o FirewallVip46Output) ToFirewallVip46PtrOutput() FirewallVip46PtrOutput {
	return o.ToFirewallVip46PtrOutputWithContext(context.Background())
}

func (o FirewallVip46Output) ToFirewallVip46PtrOutputWithContext(ctx context.Context) FirewallVip46PtrOutput {
	return o.ApplyT(func(v FirewallVip46) *FirewallVip46 {
		return &v
	}).(FirewallVip46PtrOutput)
}

type FirewallVip46PtrOutput struct {
	*pulumi.OutputState
}

func (FirewallVip46PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVip46)(nil))
}

func (o FirewallVip46PtrOutput) ToFirewallVip46PtrOutput() FirewallVip46PtrOutput {
	return o
}

func (o FirewallVip46PtrOutput) ToFirewallVip46PtrOutputWithContext(ctx context.Context) FirewallVip46PtrOutput {
	return o
}

type FirewallVip46ArrayOutput struct{ *pulumi.OutputState }

func (FirewallVip46ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallVip46)(nil))
}

func (o FirewallVip46ArrayOutput) ToFirewallVip46ArrayOutput() FirewallVip46ArrayOutput {
	return o
}

func (o FirewallVip46ArrayOutput) ToFirewallVip46ArrayOutputWithContext(ctx context.Context) FirewallVip46ArrayOutput {
	return o
}

func (o FirewallVip46ArrayOutput) Index(i pulumi.IntInput) FirewallVip46Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallVip46 {
		return vs[0].([]FirewallVip46)[vs[1].(int)]
	}).(FirewallVip46Output)
}

type FirewallVip46MapOutput struct{ *pulumi.OutputState }

func (FirewallVip46MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallVip46)(nil))
}

func (o FirewallVip46MapOutput) ToFirewallVip46MapOutput() FirewallVip46MapOutput {
	return o
}

func (o FirewallVip46MapOutput) ToFirewallVip46MapOutputWithContext(ctx context.Context) FirewallVip46MapOutput {
	return o
}

func (o FirewallVip46MapOutput) MapIndex(k pulumi.StringInput) FirewallVip46Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallVip46 {
		return vs[0].(map[string]FirewallVip46)[vs[1].(string)]
	}).(FirewallVip46Output)
}

func init() {
	pulumi.RegisterOutputType(FirewallVip46Output{})
	pulumi.RegisterOutputType(FirewallVip46PtrOutput{})
	pulumi.RegisterOutputType(FirewallVip46ArrayOutput{})
	pulumi.RegisterOutputType(FirewallVip46MapOutput{})
}

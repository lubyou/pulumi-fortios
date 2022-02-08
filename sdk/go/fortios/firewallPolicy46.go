// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 to IPv6 policies. Applies to FortiOS Version `<= 7.0.0`.
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
// 		trnameFirewallVip46, err := fortios.NewFirewallVip46(ctx, "trnameFirewallVip46", &fortios.FirewallVip46Args{
// 			ArpReply:    pulumi.String("enable"),
// 			Color:       pulumi.Int(0),
// 			Extip:       pulumi.String("10.1.100.55"),
// 			Extport:     pulumi.String("0-65535"),
// 			Fosid:       pulumi.Int(0),
// 			LdbMethod:   pulumi.String("static"),
// 			Mappedip:    pulumi.String("2000:172:16:200::55"),
// 			Mappedport:  pulumi.String("0-65535"),
// 			Portforward: pulumi.String("disable"),
// 			Protocol:    pulumi.String("tcp"),
// 			Type:        pulumi.String("static-nat"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fortios.NewFirewallPolicy46(ctx, "trnameFirewallPolicy46", &fortios.FirewallPolicy46Args{
// 			Action:         pulumi.String("deny"),
// 			Dstintf:        pulumi.String("port3"),
// 			Fixedport:      pulumi.String("disable"),
// 			Ippool:         pulumi.String("disable"),
// 			Logtraffic:     pulumi.String("disable"),
// 			PermitAnyHost:  pulumi.String("disable"),
// 			Policyid:       pulumi.Int(2),
// 			Schedule:       pulumi.String("always"),
// 			Srcintf:        pulumi.String("port2"),
// 			Status:         pulumi.String("enable"),
// 			TcpMssReceiver: pulumi.Int(0),
// 			TcpMssSender:   pulumi.Int(0),
// 			Dstaddrs: FirewallPolicy46DstaddrArray{
// 				&FirewallPolicy46DstaddrArgs{
// 					Name: trnameFirewallVip46.Name,
// 				},
// 			},
// 			Services: FirewallPolicy46ServiceArray{
// 				&FirewallPolicy46ServiceArgs{
// 					Name: pulumi.String("ALL"),
// 				},
// 			},
// 			Srcaddrs: FirewallPolicy46SrcaddrArray{
// 				&FirewallPolicy46SrcaddrArgs{
// 					Name: pulumi.String("FIREWALL_AUTH_PORTAL_ADDRESS"),
// 				},
// 			},
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
// Firewall Policy46 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallPolicy46:FirewallPolicy46 labelname {{policyid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallPolicy46 struct {
	pulumi.CustomResourceState

	// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Destination address objects. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallPolicy46DstaddrArrayOutput `pulumi:"dstaddrs"`
	// Destination interface name.
	Dstintf pulumi.StringOutput `pulumi:"dstintf"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable fixed port for this policy. Valid values: `enable`, `disable`.
	Fixedport pulumi.StringOutput `pulumi:"fixedport"`
	// Enable/disable use of IP Pools for source NAT. Valid values: `enable`, `disable`.
	Ippool pulumi.StringOutput `pulumi:"ippool"`
	// Enable/disable traffic logging for this policy. Valid values: `enable`, `disable`.
	Logtraffic pulumi.StringOutput `pulumi:"logtraffic"`
	// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
	LogtrafficStart pulumi.StringOutput `pulumi:"logtrafficStart"`
	// IP pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Per IP traffic shaper.
	PerIpShaper pulumi.StringOutput `pulumi:"perIpShaper"`
	// Enable/disable allowing any host. Valid values: `enable`, `disable`.
	PermitAnyHost pulumi.StringOutput `pulumi:"permitAnyHost"`
	// Policy ID.
	Policyid pulumi.IntOutput `pulumi:"policyid"`
	// IP Pool names. The structure of `poolname` block is documented below.
	Poolnames FirewallPolicy46PoolnameArrayOutput `pulumi:"poolnames"`
	// Schedule name.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// Service name. The structure of `service` block is documented below.
	Services FirewallPolicy46ServiceArrayOutput `pulumi:"services"`
	// Source address objects. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallPolicy46SrcaddrArrayOutput `pulumi:"srcaddrs"`
	// Source interface name.
	Srcintf pulumi.StringOutput `pulumi:"srcintf"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
	TcpMssReceiver pulumi.IntOutput `pulumi:"tcpMssReceiver"`
	// TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
	TcpMssSender pulumi.IntOutput `pulumi:"tcpMssSender"`
	// Traffic shaper.
	TrafficShaper pulumi.StringOutput `pulumi:"trafficShaper"`
	// Reverse traffic shaper.
	TrafficShaperReverse pulumi.StringOutput `pulumi:"trafficShaperReverse"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallPolicy46 registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy46(ctx *pulumi.Context,
	name string, args *FirewallPolicy46Args, opts ...pulumi.ResourceOption) (*FirewallPolicy46, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dstaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Dstaddrs'")
	}
	if args.Dstintf == nil {
		return nil, errors.New("invalid value for required argument 'Dstintf'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.Srcaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Srcaddrs'")
	}
	if args.Srcintf == nil {
		return nil, errors.New("invalid value for required argument 'Srcintf'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallPolicy46
	err := ctx.RegisterResource("fortios:index/firewallPolicy46:FirewallPolicy46", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy46 gets an existing FirewallPolicy46 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy46(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicy46State, opts ...pulumi.ResourceOption) (*FirewallPolicy46, error) {
	var resource FirewallPolicy46
	err := ctx.ReadResource("fortios:index/firewallPolicy46:FirewallPolicy46", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy46 resources.
type firewallPolicy46State struct {
	// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
	Action *string `pulumi:"action"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address objects. The structure of `dstaddr` block is documented below.
	Dstaddrs []FirewallPolicy46Dstaddr `pulumi:"dstaddrs"`
	// Destination interface name.
	Dstintf *string `pulumi:"dstintf"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable fixed port for this policy. Valid values: `enable`, `disable`.
	Fixedport *string `pulumi:"fixedport"`
	// Enable/disable use of IP Pools for source NAT. Valid values: `enable`, `disable`.
	Ippool *string `pulumi:"ippool"`
	// Enable/disable traffic logging for this policy. Valid values: `enable`, `disable`.
	Logtraffic *string `pulumi:"logtraffic"`
	// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
	LogtrafficStart *string `pulumi:"logtrafficStart"`
	// IP pool name.
	Name *string `pulumi:"name"`
	// Per IP traffic shaper.
	PerIpShaper *string `pulumi:"perIpShaper"`
	// Enable/disable allowing any host. Valid values: `enable`, `disable`.
	PermitAnyHost *string `pulumi:"permitAnyHost"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// IP Pool names. The structure of `poolname` block is documented below.
	Poolnames []FirewallPolicy46Poolname `pulumi:"poolnames"`
	// Schedule name.
	Schedule *string `pulumi:"schedule"`
	// Service name. The structure of `service` block is documented below.
	Services []FirewallPolicy46Service `pulumi:"services"`
	// Source address objects. The structure of `srcaddr` block is documented below.
	Srcaddrs []FirewallPolicy46Srcaddr `pulumi:"srcaddrs"`
	// Source interface name.
	Srcintf *string `pulumi:"srcintf"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
	TcpMssReceiver *int `pulumi:"tcpMssReceiver"`
	// TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
	TcpMssSender *int `pulumi:"tcpMssSender"`
	// Traffic shaper.
	TrafficShaper *string `pulumi:"trafficShaper"`
	// Reverse traffic shaper.
	TrafficShaperReverse *string `pulumi:"trafficShaperReverse"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallPolicy46State struct {
	// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
	Action pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address objects. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallPolicy46DstaddrArrayInput
	// Destination interface name.
	Dstintf pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable fixed port for this policy. Valid values: `enable`, `disable`.
	Fixedport pulumi.StringPtrInput
	// Enable/disable use of IP Pools for source NAT. Valid values: `enable`, `disable`.
	Ippool pulumi.StringPtrInput
	// Enable/disable traffic logging for this policy. Valid values: `enable`, `disable`.
	Logtraffic pulumi.StringPtrInput
	// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
	LogtrafficStart pulumi.StringPtrInput
	// IP pool name.
	Name pulumi.StringPtrInput
	// Per IP traffic shaper.
	PerIpShaper pulumi.StringPtrInput
	// Enable/disable allowing any host. Valid values: `enable`, `disable`.
	PermitAnyHost pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// IP Pool names. The structure of `poolname` block is documented below.
	Poolnames FirewallPolicy46PoolnameArrayInput
	// Schedule name.
	Schedule pulumi.StringPtrInput
	// Service name. The structure of `service` block is documented below.
	Services FirewallPolicy46ServiceArrayInput
	// Source address objects. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallPolicy46SrcaddrArrayInput
	// Source interface name.
	Srcintf pulumi.StringPtrInput
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
	TcpMssReceiver pulumi.IntPtrInput
	// TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
	TcpMssSender pulumi.IntPtrInput
	// Traffic shaper.
	TrafficShaper pulumi.StringPtrInput
	// Reverse traffic shaper.
	TrafficShaperReverse pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallPolicy46State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicy46State)(nil)).Elem()
}

type firewallPolicy46Args struct {
	// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
	Action *string `pulumi:"action"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address objects. The structure of `dstaddr` block is documented below.
	Dstaddrs []FirewallPolicy46Dstaddr `pulumi:"dstaddrs"`
	// Destination interface name.
	Dstintf string `pulumi:"dstintf"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable fixed port for this policy. Valid values: `enable`, `disable`.
	Fixedport *string `pulumi:"fixedport"`
	// Enable/disable use of IP Pools for source NAT. Valid values: `enable`, `disable`.
	Ippool *string `pulumi:"ippool"`
	// Enable/disable traffic logging for this policy. Valid values: `enable`, `disable`.
	Logtraffic *string `pulumi:"logtraffic"`
	// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
	LogtrafficStart *string `pulumi:"logtrafficStart"`
	// IP pool name.
	Name *string `pulumi:"name"`
	// Per IP traffic shaper.
	PerIpShaper *string `pulumi:"perIpShaper"`
	// Enable/disable allowing any host. Valid values: `enable`, `disable`.
	PermitAnyHost *string `pulumi:"permitAnyHost"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// IP Pool names. The structure of `poolname` block is documented below.
	Poolnames []FirewallPolicy46Poolname `pulumi:"poolnames"`
	// Schedule name.
	Schedule string `pulumi:"schedule"`
	// Service name. The structure of `service` block is documented below.
	Services []FirewallPolicy46Service `pulumi:"services"`
	// Source address objects. The structure of `srcaddr` block is documented below.
	Srcaddrs []FirewallPolicy46Srcaddr `pulumi:"srcaddrs"`
	// Source interface name.
	Srcintf string `pulumi:"srcintf"`
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
	TcpMssReceiver *int `pulumi:"tcpMssReceiver"`
	// TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
	TcpMssSender *int `pulumi:"tcpMssSender"`
	// Traffic shaper.
	TrafficShaper *string `pulumi:"trafficShaper"`
	// Reverse traffic shaper.
	TrafficShaperReverse *string `pulumi:"trafficShaperReverse"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallPolicy46 resource.
type FirewallPolicy46Args struct {
	// Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
	Action pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address objects. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallPolicy46DstaddrArrayInput
	// Destination interface name.
	Dstintf pulumi.StringInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable fixed port for this policy. Valid values: `enable`, `disable`.
	Fixedport pulumi.StringPtrInput
	// Enable/disable use of IP Pools for source NAT. Valid values: `enable`, `disable`.
	Ippool pulumi.StringPtrInput
	// Enable/disable traffic logging for this policy. Valid values: `enable`, `disable`.
	Logtraffic pulumi.StringPtrInput
	// Record logs when a session starts and ends. Valid values: `enable`, `disable`.
	LogtrafficStart pulumi.StringPtrInput
	// IP pool name.
	Name pulumi.StringPtrInput
	// Per IP traffic shaper.
	PerIpShaper pulumi.StringPtrInput
	// Enable/disable allowing any host. Valid values: `enable`, `disable`.
	PermitAnyHost pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// IP Pool names. The structure of `poolname` block is documented below.
	Poolnames FirewallPolicy46PoolnameArrayInput
	// Schedule name.
	Schedule pulumi.StringInput
	// Service name. The structure of `service` block is documented below.
	Services FirewallPolicy46ServiceArrayInput
	// Source address objects. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallPolicy46SrcaddrArrayInput
	// Source interface name.
	Srcintf pulumi.StringInput
	// Enable/disable this policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
	TcpMssReceiver pulumi.IntPtrInput
	// TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
	TcpMssSender pulumi.IntPtrInput
	// Traffic shaper.
	TrafficShaper pulumi.StringPtrInput
	// Reverse traffic shaper.
	TrafficShaperReverse pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallPolicy46Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicy46Args)(nil)).Elem()
}

type FirewallPolicy46Input interface {
	pulumi.Input

	ToFirewallPolicy46Output() FirewallPolicy46Output
	ToFirewallPolicy46OutputWithContext(ctx context.Context) FirewallPolicy46Output
}

func (*FirewallPolicy46) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy46)(nil)).Elem()
}

func (i *FirewallPolicy46) ToFirewallPolicy46Output() FirewallPolicy46Output {
	return i.ToFirewallPolicy46OutputWithContext(context.Background())
}

func (i *FirewallPolicy46) ToFirewallPolicy46OutputWithContext(ctx context.Context) FirewallPolicy46Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicy46Output)
}

// FirewallPolicy46ArrayInput is an input type that accepts FirewallPolicy46Array and FirewallPolicy46ArrayOutput values.
// You can construct a concrete instance of `FirewallPolicy46ArrayInput` via:
//
//          FirewallPolicy46Array{ FirewallPolicy46Args{...} }
type FirewallPolicy46ArrayInput interface {
	pulumi.Input

	ToFirewallPolicy46ArrayOutput() FirewallPolicy46ArrayOutput
	ToFirewallPolicy46ArrayOutputWithContext(context.Context) FirewallPolicy46ArrayOutput
}

type FirewallPolicy46Array []FirewallPolicy46Input

func (FirewallPolicy46Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallPolicy46)(nil)).Elem()
}

func (i FirewallPolicy46Array) ToFirewallPolicy46ArrayOutput() FirewallPolicy46ArrayOutput {
	return i.ToFirewallPolicy46ArrayOutputWithContext(context.Background())
}

func (i FirewallPolicy46Array) ToFirewallPolicy46ArrayOutputWithContext(ctx context.Context) FirewallPolicy46ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicy46ArrayOutput)
}

// FirewallPolicy46MapInput is an input type that accepts FirewallPolicy46Map and FirewallPolicy46MapOutput values.
// You can construct a concrete instance of `FirewallPolicy46MapInput` via:
//
//          FirewallPolicy46Map{ "key": FirewallPolicy46Args{...} }
type FirewallPolicy46MapInput interface {
	pulumi.Input

	ToFirewallPolicy46MapOutput() FirewallPolicy46MapOutput
	ToFirewallPolicy46MapOutputWithContext(context.Context) FirewallPolicy46MapOutput
}

type FirewallPolicy46Map map[string]FirewallPolicy46Input

func (FirewallPolicy46Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallPolicy46)(nil)).Elem()
}

func (i FirewallPolicy46Map) ToFirewallPolicy46MapOutput() FirewallPolicy46MapOutput {
	return i.ToFirewallPolicy46MapOutputWithContext(context.Background())
}

func (i FirewallPolicy46Map) ToFirewallPolicy46MapOutputWithContext(ctx context.Context) FirewallPolicy46MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicy46MapOutput)
}

type FirewallPolicy46Output struct{ *pulumi.OutputState }

func (FirewallPolicy46Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy46)(nil)).Elem()
}

func (o FirewallPolicy46Output) ToFirewallPolicy46Output() FirewallPolicy46Output {
	return o
}

func (o FirewallPolicy46Output) ToFirewallPolicy46OutputWithContext(ctx context.Context) FirewallPolicy46Output {
	return o
}

type FirewallPolicy46ArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicy46ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallPolicy46)(nil)).Elem()
}

func (o FirewallPolicy46ArrayOutput) ToFirewallPolicy46ArrayOutput() FirewallPolicy46ArrayOutput {
	return o
}

func (o FirewallPolicy46ArrayOutput) ToFirewallPolicy46ArrayOutputWithContext(ctx context.Context) FirewallPolicy46ArrayOutput {
	return o
}

func (o FirewallPolicy46ArrayOutput) Index(i pulumi.IntInput) FirewallPolicy46Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallPolicy46 {
		return vs[0].([]*FirewallPolicy46)[vs[1].(int)]
	}).(FirewallPolicy46Output)
}

type FirewallPolicy46MapOutput struct{ *pulumi.OutputState }

func (FirewallPolicy46MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallPolicy46)(nil)).Elem()
}

func (o FirewallPolicy46MapOutput) ToFirewallPolicy46MapOutput() FirewallPolicy46MapOutput {
	return o
}

func (o FirewallPolicy46MapOutput) ToFirewallPolicy46MapOutputWithContext(ctx context.Context) FirewallPolicy46MapOutput {
	return o
}

func (o FirewallPolicy46MapOutput) MapIndex(k pulumi.StringInput) FirewallPolicy46Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallPolicy46 {
		return vs[0].(map[string]*FirewallPolicy46)[vs[1].(string)]
	}).(FirewallPolicy46Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicy46Input)(nil)).Elem(), &FirewallPolicy46{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicy46ArrayInput)(nil)).Elem(), FirewallPolicy46Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicy46MapInput)(nil)).Elem(), FirewallPolicy46Map{})
	pulumi.RegisterOutputType(FirewallPolicy46Output{})
	pulumi.RegisterOutputType(FirewallPolicy46ArrayOutput{})
	pulumi.RegisterOutputType(FirewallPolicy46MapOutput{})
}

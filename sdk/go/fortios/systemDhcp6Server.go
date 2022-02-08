// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure DHCPv6 servers.
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
// 		_, err := fortios.NewSystemDhcp6Server(ctx, "trname", &fortios.SystemDhcp6ServerArgs{
// 			Fosid:       pulumi.Int(1),
// 			Interface:   pulumi.String("port3"),
// 			LeaseTime:   pulumi.Int(604800),
// 			RapidCommit: pulumi.String("disable"),
// 			Status:      pulumi.String("enable"),
// 			Subnet:      pulumi.String("2001:db8:1234:113::/64"),
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
// SystemDhcp6 Server can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemDhcp6Server:SystemDhcp6Server labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemDhcp6Server struct {
	pulumi.CustomResourceState

	// IAID of obtained delegated-prefix from the upstream interface.
	DelegatedPrefixIaid pulumi.IntOutput `pulumi:"delegatedPrefixIaid"`
	// DNS search list options. Valid values: `delegated`, `specify`.
	DnsSearchList pulumi.StringOutput `pulumi:"dnsSearchList"`
	// DNS server 1.
	DnsServer1 pulumi.StringOutput `pulumi:"dnsServer1"`
	// DNS server 2.
	DnsServer2 pulumi.StringOutput `pulumi:"dnsServer2"`
	// DNS server 3.
	DnsServer3 pulumi.StringOutput `pulumi:"dnsServer3"`
	// DNS server 4.
	DnsServer4 pulumi.StringOutput `pulumi:"dnsServer4"`
	// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
	DnsService pulumi.StringOutput `pulumi:"dnsService"`
	// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode pulumi.StringOutput `pulumi:"ipMode"`
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges SystemDhcp6ServerIpRangeArrayOutput `pulumi:"ipRanges"`
	// Lease time in seconds, 0 means unlimited.
	LeaseTime pulumi.IntOutput `pulumi:"leaseTime"`
	// Option 1.
	Option1 pulumi.StringOutput `pulumi:"option1"`
	// Option 2.
	Option2 pulumi.StringOutput `pulumi:"option2"`
	// Option 3.
	Option3 pulumi.StringOutput `pulumi:"option3"`
	// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
	PrefixMode pulumi.StringOutput `pulumi:"prefixMode"`
	// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
	PrefixRanges SystemDhcp6ServerPrefixRangeArrayOutput `pulumi:"prefixRanges"`
	// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
	RapidCommit pulumi.StringOutput `pulumi:"rapidCommit"`
	// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Subnet or subnet-id if the IP mode is delegated.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Interface name from where delegated information is provided.
	UpstreamInterface pulumi.StringOutput `pulumi:"upstreamInterface"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemDhcp6Server registers a new resource with the given unique name, arguments, and options.
func NewSystemDhcp6Server(ctx *pulumi.Context,
	name string, args *SystemDhcp6ServerArgs, opts ...pulumi.ResourceOption) (*SystemDhcp6Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemDhcp6Server
	err := ctx.RegisterResource("fortios:index/systemDhcp6Server:SystemDhcp6Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemDhcp6Server gets an existing SystemDhcp6Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemDhcp6Server(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemDhcp6ServerState, opts ...pulumi.ResourceOption) (*SystemDhcp6Server, error) {
	var resource SystemDhcp6Server
	err := ctx.ReadResource("fortios:index/systemDhcp6Server:SystemDhcp6Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemDhcp6Server resources.
type systemDhcp6ServerState struct {
	// IAID of obtained delegated-prefix from the upstream interface.
	DelegatedPrefixIaid *int `pulumi:"delegatedPrefixIaid"`
	// DNS search list options. Valid values: `delegated`, `specify`.
	DnsSearchList *string `pulumi:"dnsSearchList"`
	// DNS server 1.
	DnsServer1 *string `pulumi:"dnsServer1"`
	// DNS server 2.
	DnsServer2 *string `pulumi:"dnsServer2"`
	// DNS server 3.
	DnsServer3 *string `pulumi:"dnsServer3"`
	// DNS server 4.
	DnsServer4 *string `pulumi:"dnsServer4"`
	// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
	DnsService *string `pulumi:"dnsService"`
	// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
	Domain *string `pulumi:"domain"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface *string `pulumi:"interface"`
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode *string `pulumi:"ipMode"`
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges []SystemDhcp6ServerIpRange `pulumi:"ipRanges"`
	// Lease time in seconds, 0 means unlimited.
	LeaseTime *int `pulumi:"leaseTime"`
	// Option 1.
	Option1 *string `pulumi:"option1"`
	// Option 2.
	Option2 *string `pulumi:"option2"`
	// Option 3.
	Option3 *string `pulumi:"option3"`
	// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
	PrefixMode *string `pulumi:"prefixMode"`
	// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
	PrefixRanges []SystemDhcp6ServerPrefixRange `pulumi:"prefixRanges"`
	// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
	RapidCommit *string `pulumi:"rapidCommit"`
	// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Subnet or subnet-id if the IP mode is delegated.
	Subnet *string `pulumi:"subnet"`
	// Interface name from where delegated information is provided.
	UpstreamInterface *string `pulumi:"upstreamInterface"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemDhcp6ServerState struct {
	// IAID of obtained delegated-prefix from the upstream interface.
	DelegatedPrefixIaid pulumi.IntPtrInput
	// DNS search list options. Valid values: `delegated`, `specify`.
	DnsSearchList pulumi.StringPtrInput
	// DNS server 1.
	DnsServer1 pulumi.StringPtrInput
	// DNS server 2.
	DnsServer2 pulumi.StringPtrInput
	// DNS server 3.
	DnsServer3 pulumi.StringPtrInput
	// DNS server 4.
	DnsServer4 pulumi.StringPtrInput
	// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
	DnsService pulumi.StringPtrInput
	// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
	Domain pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntPtrInput
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface pulumi.StringPtrInput
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode pulumi.StringPtrInput
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges SystemDhcp6ServerIpRangeArrayInput
	// Lease time in seconds, 0 means unlimited.
	LeaseTime pulumi.IntPtrInput
	// Option 1.
	Option1 pulumi.StringPtrInput
	// Option 2.
	Option2 pulumi.StringPtrInput
	// Option 3.
	Option3 pulumi.StringPtrInput
	// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
	PrefixMode pulumi.StringPtrInput
	// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
	PrefixRanges SystemDhcp6ServerPrefixRangeArrayInput
	// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
	RapidCommit pulumi.StringPtrInput
	// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Subnet or subnet-id if the IP mode is delegated.
	Subnet pulumi.StringPtrInput
	// Interface name from where delegated information is provided.
	UpstreamInterface pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemDhcp6ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDhcp6ServerState)(nil)).Elem()
}

type systemDhcp6ServerArgs struct {
	// IAID of obtained delegated-prefix from the upstream interface.
	DelegatedPrefixIaid *int `pulumi:"delegatedPrefixIaid"`
	// DNS search list options. Valid values: `delegated`, `specify`.
	DnsSearchList *string `pulumi:"dnsSearchList"`
	// DNS server 1.
	DnsServer1 *string `pulumi:"dnsServer1"`
	// DNS server 2.
	DnsServer2 *string `pulumi:"dnsServer2"`
	// DNS server 3.
	DnsServer3 *string `pulumi:"dnsServer3"`
	// DNS server 4.
	DnsServer4 *string `pulumi:"dnsServer4"`
	// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
	DnsService *string `pulumi:"dnsService"`
	// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
	Domain *string `pulumi:"domain"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface string `pulumi:"interface"`
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode *string `pulumi:"ipMode"`
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges []SystemDhcp6ServerIpRange `pulumi:"ipRanges"`
	// Lease time in seconds, 0 means unlimited.
	LeaseTime *int `pulumi:"leaseTime"`
	// Option 1.
	Option1 *string `pulumi:"option1"`
	// Option 2.
	Option2 *string `pulumi:"option2"`
	// Option 3.
	Option3 *string `pulumi:"option3"`
	// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
	PrefixMode *string `pulumi:"prefixMode"`
	// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
	PrefixRanges []SystemDhcp6ServerPrefixRange `pulumi:"prefixRanges"`
	// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
	RapidCommit *string `pulumi:"rapidCommit"`
	// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Subnet or subnet-id if the IP mode is delegated.
	Subnet string `pulumi:"subnet"`
	// Interface name from where delegated information is provided.
	UpstreamInterface *string `pulumi:"upstreamInterface"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemDhcp6Server resource.
type SystemDhcp6ServerArgs struct {
	// IAID of obtained delegated-prefix from the upstream interface.
	DelegatedPrefixIaid pulumi.IntPtrInput
	// DNS search list options. Valid values: `delegated`, `specify`.
	DnsSearchList pulumi.StringPtrInput
	// DNS server 1.
	DnsServer1 pulumi.StringPtrInput
	// DNS server 2.
	DnsServer2 pulumi.StringPtrInput
	// DNS server 3.
	DnsServer3 pulumi.StringPtrInput
	// DNS server 4.
	DnsServer4 pulumi.StringPtrInput
	// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
	DnsService pulumi.StringPtrInput
	// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
	Domain pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// ID.
	Fosid pulumi.IntInput
	// DHCP server can assign IP configurations to clients connected to this interface.
	Interface pulumi.StringInput
	// Method used to assign client IP. Valid values: `range`, `delegated`.
	IpMode pulumi.StringPtrInput
	// DHCP IP range configuration. The structure of `ipRange` block is documented below.
	IpRanges SystemDhcp6ServerIpRangeArrayInput
	// Lease time in seconds, 0 means unlimited.
	LeaseTime pulumi.IntPtrInput
	// Option 1.
	Option1 pulumi.StringPtrInput
	// Option 2.
	Option2 pulumi.StringPtrInput
	// Option 3.
	Option3 pulumi.StringPtrInput
	// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
	PrefixMode pulumi.StringPtrInput
	// DHCP prefix configuration. The structure of `prefixRange` block is documented below.
	PrefixRanges SystemDhcp6ServerPrefixRangeArrayInput
	// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
	RapidCommit pulumi.StringPtrInput
	// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Subnet or subnet-id if the IP mode is delegated.
	Subnet pulumi.StringInput
	// Interface name from where delegated information is provided.
	UpstreamInterface pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemDhcp6ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDhcp6ServerArgs)(nil)).Elem()
}

type SystemDhcp6ServerInput interface {
	pulumi.Input

	ToSystemDhcp6ServerOutput() SystemDhcp6ServerOutput
	ToSystemDhcp6ServerOutputWithContext(ctx context.Context) SystemDhcp6ServerOutput
}

func (*SystemDhcp6Server) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDhcp6Server)(nil)).Elem()
}

func (i *SystemDhcp6Server) ToSystemDhcp6ServerOutput() SystemDhcp6ServerOutput {
	return i.ToSystemDhcp6ServerOutputWithContext(context.Background())
}

func (i *SystemDhcp6Server) ToSystemDhcp6ServerOutputWithContext(ctx context.Context) SystemDhcp6ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDhcp6ServerOutput)
}

// SystemDhcp6ServerArrayInput is an input type that accepts SystemDhcp6ServerArray and SystemDhcp6ServerArrayOutput values.
// You can construct a concrete instance of `SystemDhcp6ServerArrayInput` via:
//
//          SystemDhcp6ServerArray{ SystemDhcp6ServerArgs{...} }
type SystemDhcp6ServerArrayInput interface {
	pulumi.Input

	ToSystemDhcp6ServerArrayOutput() SystemDhcp6ServerArrayOutput
	ToSystemDhcp6ServerArrayOutputWithContext(context.Context) SystemDhcp6ServerArrayOutput
}

type SystemDhcp6ServerArray []SystemDhcp6ServerInput

func (SystemDhcp6ServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDhcp6Server)(nil)).Elem()
}

func (i SystemDhcp6ServerArray) ToSystemDhcp6ServerArrayOutput() SystemDhcp6ServerArrayOutput {
	return i.ToSystemDhcp6ServerArrayOutputWithContext(context.Background())
}

func (i SystemDhcp6ServerArray) ToSystemDhcp6ServerArrayOutputWithContext(ctx context.Context) SystemDhcp6ServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDhcp6ServerArrayOutput)
}

// SystemDhcp6ServerMapInput is an input type that accepts SystemDhcp6ServerMap and SystemDhcp6ServerMapOutput values.
// You can construct a concrete instance of `SystemDhcp6ServerMapInput` via:
//
//          SystemDhcp6ServerMap{ "key": SystemDhcp6ServerArgs{...} }
type SystemDhcp6ServerMapInput interface {
	pulumi.Input

	ToSystemDhcp6ServerMapOutput() SystemDhcp6ServerMapOutput
	ToSystemDhcp6ServerMapOutputWithContext(context.Context) SystemDhcp6ServerMapOutput
}

type SystemDhcp6ServerMap map[string]SystemDhcp6ServerInput

func (SystemDhcp6ServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDhcp6Server)(nil)).Elem()
}

func (i SystemDhcp6ServerMap) ToSystemDhcp6ServerMapOutput() SystemDhcp6ServerMapOutput {
	return i.ToSystemDhcp6ServerMapOutputWithContext(context.Background())
}

func (i SystemDhcp6ServerMap) ToSystemDhcp6ServerMapOutputWithContext(ctx context.Context) SystemDhcp6ServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDhcp6ServerMapOutput)
}

type SystemDhcp6ServerOutput struct{ *pulumi.OutputState }

func (SystemDhcp6ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDhcp6Server)(nil)).Elem()
}

func (o SystemDhcp6ServerOutput) ToSystemDhcp6ServerOutput() SystemDhcp6ServerOutput {
	return o
}

func (o SystemDhcp6ServerOutput) ToSystemDhcp6ServerOutputWithContext(ctx context.Context) SystemDhcp6ServerOutput {
	return o
}

type SystemDhcp6ServerArrayOutput struct{ *pulumi.OutputState }

func (SystemDhcp6ServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDhcp6Server)(nil)).Elem()
}

func (o SystemDhcp6ServerArrayOutput) ToSystemDhcp6ServerArrayOutput() SystemDhcp6ServerArrayOutput {
	return o
}

func (o SystemDhcp6ServerArrayOutput) ToSystemDhcp6ServerArrayOutputWithContext(ctx context.Context) SystemDhcp6ServerArrayOutput {
	return o
}

func (o SystemDhcp6ServerArrayOutput) Index(i pulumi.IntInput) SystemDhcp6ServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemDhcp6Server {
		return vs[0].([]*SystemDhcp6Server)[vs[1].(int)]
	}).(SystemDhcp6ServerOutput)
}

type SystemDhcp6ServerMapOutput struct{ *pulumi.OutputState }

func (SystemDhcp6ServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDhcp6Server)(nil)).Elem()
}

func (o SystemDhcp6ServerMapOutput) ToSystemDhcp6ServerMapOutput() SystemDhcp6ServerMapOutput {
	return o
}

func (o SystemDhcp6ServerMapOutput) ToSystemDhcp6ServerMapOutputWithContext(ctx context.Context) SystemDhcp6ServerMapOutput {
	return o
}

func (o SystemDhcp6ServerMapOutput) MapIndex(k pulumi.StringInput) SystemDhcp6ServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemDhcp6Server {
		return vs[0].(map[string]*SystemDhcp6Server)[vs[1].(string)]
	}).(SystemDhcp6ServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDhcp6ServerInput)(nil)).Elem(), &SystemDhcp6Server{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDhcp6ServerArrayInput)(nil)).Elem(), SystemDhcp6ServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDhcp6ServerMapInput)(nil)).Elem(), SystemDhcp6ServerMap{})
	pulumi.RegisterOutputType(SystemDhcp6ServerOutput{})
	pulumi.RegisterOutputType(SystemDhcp6ServerArrayOutput{})
	pulumi.RegisterOutputType(SystemDhcp6ServerMapOutput{})
}

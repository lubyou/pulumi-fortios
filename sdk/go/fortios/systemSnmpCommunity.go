// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SNMP community configuration.
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
// 		_, err := fortios.NewSystemSnmpCommunity(ctx, "trname", &fortios.SystemSnmpCommunityArgs{
// 			Events:         pulumi.String("cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high"),
// 			Fosid:          pulumi.Int(1),
// 			QueryV1Port:    pulumi.Int(161),
// 			QueryV1Status:  pulumi.String("enable"),
// 			QueryV2cPort:   pulumi.Int(161),
// 			QueryV2cStatus: pulumi.String("enable"),
// 			Status:         pulumi.String("enable"),
// 			TrapV1Lport:    pulumi.Int(162),
// 			TrapV1Rport:    pulumi.Int(162),
// 			TrapV1Status:   pulumi.String("enable"),
// 			TrapV2cLport:   pulumi.Int(162),
// 			TrapV2cRport:   pulumi.Int(162),
// 			TrapV2cStatus:  pulumi.String("enable"),
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
// SystemSnmp Community can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemSnmpCommunity:SystemSnmpCommunity labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemSnmpCommunity struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// SNMP trap events.
	Events pulumi.StringOutput `pulumi:"events"`
	// Community ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
	Hosts SystemSnmpCommunityHostArrayOutput `pulumi:"hosts"`
	// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
	Hosts6s SystemSnmpCommunityHosts6ArrayOutput `pulumi:"hosts6s"`
	// Community name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SNMP v1 query port (default = 161).
	QueryV1Port pulumi.IntOutput `pulumi:"queryV1Port"`
	// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
	QueryV1Status pulumi.StringOutput `pulumi:"queryV1Status"`
	// SNMP v2c query port (default = 161).
	QueryV2cPort pulumi.IntOutput `pulumi:"queryV2cPort"`
	// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
	QueryV2cStatus pulumi.StringOutput `pulumi:"queryV2cStatus"`
	// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// SNMP v1 trap local port (default = 162).
	TrapV1Lport pulumi.IntOutput `pulumi:"trapV1Lport"`
	// SNMP v1 trap remote port (default = 162).
	TrapV1Rport pulumi.IntOutput `pulumi:"trapV1Rport"`
	// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
	TrapV1Status pulumi.StringOutput `pulumi:"trapV1Status"`
	// SNMP v2c trap local port (default = 162).
	TrapV2cLport pulumi.IntOutput `pulumi:"trapV2cLport"`
	// SNMP v2c trap remote port (default = 162).
	TrapV2cRport pulumi.IntOutput `pulumi:"trapV2cRport"`
	// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
	TrapV2cStatus pulumi.StringOutput `pulumi:"trapV2cStatus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSnmpCommunity registers a new resource with the given unique name, arguments, and options.
func NewSystemSnmpCommunity(ctx *pulumi.Context,
	name string, args *SystemSnmpCommunityArgs, opts ...pulumi.ResourceOption) (*SystemSnmpCommunity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSnmpCommunity
	err := ctx.RegisterResource("fortios:index/systemSnmpCommunity:SystemSnmpCommunity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSnmpCommunity gets an existing SystemSnmpCommunity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSnmpCommunity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSnmpCommunityState, opts ...pulumi.ResourceOption) (*SystemSnmpCommunity, error) {
	var resource SystemSnmpCommunity
	err := ctx.ReadResource("fortios:index/systemSnmpCommunity:SystemSnmpCommunity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSnmpCommunity resources.
type systemSnmpCommunityState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// SNMP trap events.
	Events *string `pulumi:"events"`
	// Community ID.
	Fosid *int `pulumi:"fosid"`
	// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
	Hosts []SystemSnmpCommunityHost `pulumi:"hosts"`
	// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
	Hosts6s []SystemSnmpCommunityHosts6 `pulumi:"hosts6s"`
	// Community name.
	Name *string `pulumi:"name"`
	// SNMP v1 query port (default = 161).
	QueryV1Port *int `pulumi:"queryV1Port"`
	// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
	QueryV1Status *string `pulumi:"queryV1Status"`
	// SNMP v2c query port (default = 161).
	QueryV2cPort *int `pulumi:"queryV2cPort"`
	// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
	QueryV2cStatus *string `pulumi:"queryV2cStatus"`
	// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// SNMP v1 trap local port (default = 162).
	TrapV1Lport *int `pulumi:"trapV1Lport"`
	// SNMP v1 trap remote port (default = 162).
	TrapV1Rport *int `pulumi:"trapV1Rport"`
	// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
	TrapV1Status *string `pulumi:"trapV1Status"`
	// SNMP v2c trap local port (default = 162).
	TrapV2cLport *int `pulumi:"trapV2cLport"`
	// SNMP v2c trap remote port (default = 162).
	TrapV2cRport *int `pulumi:"trapV2cRport"`
	// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
	TrapV2cStatus *string `pulumi:"trapV2cStatus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemSnmpCommunityState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// SNMP trap events.
	Events pulumi.StringPtrInput
	// Community ID.
	Fosid pulumi.IntPtrInput
	// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
	Hosts SystemSnmpCommunityHostArrayInput
	// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
	Hosts6s SystemSnmpCommunityHosts6ArrayInput
	// Community name.
	Name pulumi.StringPtrInput
	// SNMP v1 query port (default = 161).
	QueryV1Port pulumi.IntPtrInput
	// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
	QueryV1Status pulumi.StringPtrInput
	// SNMP v2c query port (default = 161).
	QueryV2cPort pulumi.IntPtrInput
	// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
	QueryV2cStatus pulumi.StringPtrInput
	// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// SNMP v1 trap local port (default = 162).
	TrapV1Lport pulumi.IntPtrInput
	// SNMP v1 trap remote port (default = 162).
	TrapV1Rport pulumi.IntPtrInput
	// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
	TrapV1Status pulumi.StringPtrInput
	// SNMP v2c trap local port (default = 162).
	TrapV2cLport pulumi.IntPtrInput
	// SNMP v2c trap remote port (default = 162).
	TrapV2cRport pulumi.IntPtrInput
	// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
	TrapV2cStatus pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSnmpCommunityState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSnmpCommunityState)(nil)).Elem()
}

type systemSnmpCommunityArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// SNMP trap events.
	Events *string `pulumi:"events"`
	// Community ID.
	Fosid int `pulumi:"fosid"`
	// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
	Hosts []SystemSnmpCommunityHost `pulumi:"hosts"`
	// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
	Hosts6s []SystemSnmpCommunityHosts6 `pulumi:"hosts6s"`
	// Community name.
	Name *string `pulumi:"name"`
	// SNMP v1 query port (default = 161).
	QueryV1Port *int `pulumi:"queryV1Port"`
	// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
	QueryV1Status *string `pulumi:"queryV1Status"`
	// SNMP v2c query port (default = 161).
	QueryV2cPort *int `pulumi:"queryV2cPort"`
	// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
	QueryV2cStatus *string `pulumi:"queryV2cStatus"`
	// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// SNMP v1 trap local port (default = 162).
	TrapV1Lport *int `pulumi:"trapV1Lport"`
	// SNMP v1 trap remote port (default = 162).
	TrapV1Rport *int `pulumi:"trapV1Rport"`
	// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
	TrapV1Status *string `pulumi:"trapV1Status"`
	// SNMP v2c trap local port (default = 162).
	TrapV2cLport *int `pulumi:"trapV2cLport"`
	// SNMP v2c trap remote port (default = 162).
	TrapV2cRport *int `pulumi:"trapV2cRport"`
	// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
	TrapV2cStatus *string `pulumi:"trapV2cStatus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSnmpCommunity resource.
type SystemSnmpCommunityArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// SNMP trap events.
	Events pulumi.StringPtrInput
	// Community ID.
	Fosid pulumi.IntInput
	// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
	Hosts SystemSnmpCommunityHostArrayInput
	// Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
	Hosts6s SystemSnmpCommunityHosts6ArrayInput
	// Community name.
	Name pulumi.StringPtrInput
	// SNMP v1 query port (default = 161).
	QueryV1Port pulumi.IntPtrInput
	// Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
	QueryV1Status pulumi.StringPtrInput
	// SNMP v2c query port (default = 161).
	QueryV2cPort pulumi.IntPtrInput
	// Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
	QueryV2cStatus pulumi.StringPtrInput
	// Enable/disable this SNMP community. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// SNMP v1 trap local port (default = 162).
	TrapV1Lport pulumi.IntPtrInput
	// SNMP v1 trap remote port (default = 162).
	TrapV1Rport pulumi.IntPtrInput
	// Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
	TrapV1Status pulumi.StringPtrInput
	// SNMP v2c trap local port (default = 162).
	TrapV2cLport pulumi.IntPtrInput
	// SNMP v2c trap remote port (default = 162).
	TrapV2cRport pulumi.IntPtrInput
	// Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
	TrapV2cStatus pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSnmpCommunityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSnmpCommunityArgs)(nil)).Elem()
}

type SystemSnmpCommunityInput interface {
	pulumi.Input

	ToSystemSnmpCommunityOutput() SystemSnmpCommunityOutput
	ToSystemSnmpCommunityOutputWithContext(ctx context.Context) SystemSnmpCommunityOutput
}

func (*SystemSnmpCommunity) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSnmpCommunity)(nil)).Elem()
}

func (i *SystemSnmpCommunity) ToSystemSnmpCommunityOutput() SystemSnmpCommunityOutput {
	return i.ToSystemSnmpCommunityOutputWithContext(context.Background())
}

func (i *SystemSnmpCommunity) ToSystemSnmpCommunityOutputWithContext(ctx context.Context) SystemSnmpCommunityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSnmpCommunityOutput)
}

// SystemSnmpCommunityArrayInput is an input type that accepts SystemSnmpCommunityArray and SystemSnmpCommunityArrayOutput values.
// You can construct a concrete instance of `SystemSnmpCommunityArrayInput` via:
//
//          SystemSnmpCommunityArray{ SystemSnmpCommunityArgs{...} }
type SystemSnmpCommunityArrayInput interface {
	pulumi.Input

	ToSystemSnmpCommunityArrayOutput() SystemSnmpCommunityArrayOutput
	ToSystemSnmpCommunityArrayOutputWithContext(context.Context) SystemSnmpCommunityArrayOutput
}

type SystemSnmpCommunityArray []SystemSnmpCommunityInput

func (SystemSnmpCommunityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSnmpCommunity)(nil)).Elem()
}

func (i SystemSnmpCommunityArray) ToSystemSnmpCommunityArrayOutput() SystemSnmpCommunityArrayOutput {
	return i.ToSystemSnmpCommunityArrayOutputWithContext(context.Background())
}

func (i SystemSnmpCommunityArray) ToSystemSnmpCommunityArrayOutputWithContext(ctx context.Context) SystemSnmpCommunityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSnmpCommunityArrayOutput)
}

// SystemSnmpCommunityMapInput is an input type that accepts SystemSnmpCommunityMap and SystemSnmpCommunityMapOutput values.
// You can construct a concrete instance of `SystemSnmpCommunityMapInput` via:
//
//          SystemSnmpCommunityMap{ "key": SystemSnmpCommunityArgs{...} }
type SystemSnmpCommunityMapInput interface {
	pulumi.Input

	ToSystemSnmpCommunityMapOutput() SystemSnmpCommunityMapOutput
	ToSystemSnmpCommunityMapOutputWithContext(context.Context) SystemSnmpCommunityMapOutput
}

type SystemSnmpCommunityMap map[string]SystemSnmpCommunityInput

func (SystemSnmpCommunityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSnmpCommunity)(nil)).Elem()
}

func (i SystemSnmpCommunityMap) ToSystemSnmpCommunityMapOutput() SystemSnmpCommunityMapOutput {
	return i.ToSystemSnmpCommunityMapOutputWithContext(context.Background())
}

func (i SystemSnmpCommunityMap) ToSystemSnmpCommunityMapOutputWithContext(ctx context.Context) SystemSnmpCommunityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSnmpCommunityMapOutput)
}

type SystemSnmpCommunityOutput struct{ *pulumi.OutputState }

func (SystemSnmpCommunityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSnmpCommunity)(nil)).Elem()
}

func (o SystemSnmpCommunityOutput) ToSystemSnmpCommunityOutput() SystemSnmpCommunityOutput {
	return o
}

func (o SystemSnmpCommunityOutput) ToSystemSnmpCommunityOutputWithContext(ctx context.Context) SystemSnmpCommunityOutput {
	return o
}

type SystemSnmpCommunityArrayOutput struct{ *pulumi.OutputState }

func (SystemSnmpCommunityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSnmpCommunity)(nil)).Elem()
}

func (o SystemSnmpCommunityArrayOutput) ToSystemSnmpCommunityArrayOutput() SystemSnmpCommunityArrayOutput {
	return o
}

func (o SystemSnmpCommunityArrayOutput) ToSystemSnmpCommunityArrayOutputWithContext(ctx context.Context) SystemSnmpCommunityArrayOutput {
	return o
}

func (o SystemSnmpCommunityArrayOutput) Index(i pulumi.IntInput) SystemSnmpCommunityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSnmpCommunity {
		return vs[0].([]*SystemSnmpCommunity)[vs[1].(int)]
	}).(SystemSnmpCommunityOutput)
}

type SystemSnmpCommunityMapOutput struct{ *pulumi.OutputState }

func (SystemSnmpCommunityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSnmpCommunity)(nil)).Elem()
}

func (o SystemSnmpCommunityMapOutput) ToSystemSnmpCommunityMapOutput() SystemSnmpCommunityMapOutput {
	return o
}

func (o SystemSnmpCommunityMapOutput) ToSystemSnmpCommunityMapOutputWithContext(ctx context.Context) SystemSnmpCommunityMapOutput {
	return o
}

func (o SystemSnmpCommunityMapOutput) MapIndex(k pulumi.StringInput) SystemSnmpCommunityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSnmpCommunity {
		return vs[0].(map[string]*SystemSnmpCommunity)[vs[1].(string)]
	}).(SystemSnmpCommunityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSnmpCommunityInput)(nil)).Elem(), &SystemSnmpCommunity{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSnmpCommunityArrayInput)(nil)).Elem(), SystemSnmpCommunityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSnmpCommunityMapInput)(nil)).Elem(), SystemSnmpCommunityMap{})
	pulumi.RegisterOutputType(SystemSnmpCommunityOutput{})
	pulumi.RegisterOutputType(SystemSnmpCommunityArrayOutput{})
	pulumi.RegisterOutputType(SystemSnmpCommunityMapOutput{})
}

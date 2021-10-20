// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure wireless controller global settings.
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
// 		_, err := fortios.NewWirelessControllerGlobal(ctx, "trname", &fortios.WirelessControllerGlobalArgs{
// 			ApLogServer:           pulumi.String("disable"),
// 			ApLogServerIp:         pulumi.String("0.0.0.0"),
// 			ApLogServerPort:       pulumi.Int(0),
// 			ControlMessageOffload: pulumi.String("ebp-frame aeroscout-tag ap-list sta-list sta-cap-list stats aeroscout-mu"),
// 			DataEthernetIi:        pulumi.String("disable"),
// 			DiscoveryMcAddr:       pulumi.String("224.0.1.140"),
// 			FiappEthType:          pulumi.Int(5252),
// 			ImageDownload:         pulumi.String("enable"),
// 			IpsecBaseIp:           pulumi.String("169.254.0.1"),
// 			LinkAggregation:       pulumi.String("disable"),
// 			MaxClients:            pulumi.Int(0),
// 			MaxRetransmit:         pulumi.Int(3),
// 			MeshEthType:           pulumi.Int(8755),
// 			RogueScanMacAdjacency: pulumi.Int(7),
// 			WtpShare:              pulumi.String("disable"),
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
// WirelessController Global can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerGlobal:WirelessControllerGlobal labelname WirelessControllerGlobal
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerGlobal struct {
	pulumi.CustomResourceState

	// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
	ApLogServer pulumi.StringOutput `pulumi:"apLogServer"`
	// IP address that APs or FortiAPs send log messages to.
	ApLogServerIp pulumi.StringOutput `pulumi:"apLogServerIp"`
	// Port that APs or FortiAPs send log messages to.
	ApLogServerPort pulumi.IntOutput `pulumi:"apLogServerPort"`
	// Configure CAPWAP control message data channel offload.
	ControlMessageOffload pulumi.StringOutput `pulumi:"controlMessageOffload"`
	DataEthernetIi        pulumi.StringOutput `pulumi:"dataEthernetIi"`
	// Multicast IP address for AP discovery (default = 244.0.1.140).
	DiscoveryMcAddr pulumi.StringOutput `pulumi:"discoveryMcAddr"`
	// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
	FiappEthType pulumi.IntOutput `pulumi:"fiappEthType"`
	// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
	ImageDownload pulumi.StringOutput `pulumi:"imageDownload"`
	// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
	IpsecBaseIp pulumi.StringOutput `pulumi:"ipsecBaseIp"`
	// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
	LinkAggregation pulumi.StringOutput `pulumi:"linkAggregation"`
	// Description of the location of the wireless controller.
	Location pulumi.StringOutput `pulumi:"location"`
	// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
	MaxClients pulumi.IntOutput `pulumi:"maxClients"`
	// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
	MaxRetransmit pulumi.IntOutput `pulumi:"maxRetransmit"`
	// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
	MeshEthType pulumi.IntOutput `pulumi:"meshEthType"`
	// Name of the wireless controller.
	Name pulumi.StringOutput `pulumi:"name"`
	// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
	RogueScanMacAdjacency pulumi.IntOutput `pulumi:"rogueScanMacAdjacency"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
	WtpShare pulumi.StringOutput `pulumi:"wtpShare"`
}

// NewWirelessControllerGlobal registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerGlobal(ctx *pulumi.Context,
	name string, args *WirelessControllerGlobalArgs, opts ...pulumi.ResourceOption) (*WirelessControllerGlobal, error) {
	if args == nil {
		args = &WirelessControllerGlobalArgs{}
	}

	var resource WirelessControllerGlobal
	err := ctx.RegisterResource("fortios:index/wirelessControllerGlobal:WirelessControllerGlobal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerGlobal gets an existing WirelessControllerGlobal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerGlobal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerGlobalState, opts ...pulumi.ResourceOption) (*WirelessControllerGlobal, error) {
	var resource WirelessControllerGlobal
	err := ctx.ReadResource("fortios:index/wirelessControllerGlobal:WirelessControllerGlobal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerGlobal resources.
type wirelessControllerGlobalState struct {
	// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
	ApLogServer *string `pulumi:"apLogServer"`
	// IP address that APs or FortiAPs send log messages to.
	ApLogServerIp *string `pulumi:"apLogServerIp"`
	// Port that APs or FortiAPs send log messages to.
	ApLogServerPort *int `pulumi:"apLogServerPort"`
	// Configure CAPWAP control message data channel offload.
	ControlMessageOffload *string `pulumi:"controlMessageOffload"`
	DataEthernetIi        *string `pulumi:"dataEthernetIi"`
	// Multicast IP address for AP discovery (default = 244.0.1.140).
	DiscoveryMcAddr *string `pulumi:"discoveryMcAddr"`
	// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
	FiappEthType *int `pulumi:"fiappEthType"`
	// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
	ImageDownload *string `pulumi:"imageDownload"`
	// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
	IpsecBaseIp *string `pulumi:"ipsecBaseIp"`
	// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
	LinkAggregation *string `pulumi:"linkAggregation"`
	// Description of the location of the wireless controller.
	Location *string `pulumi:"location"`
	// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
	MaxClients *int `pulumi:"maxClients"`
	// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
	MaxRetransmit *int `pulumi:"maxRetransmit"`
	// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
	MeshEthType *int `pulumi:"meshEthType"`
	// Name of the wireless controller.
	Name *string `pulumi:"name"`
	// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
	RogueScanMacAdjacency *int `pulumi:"rogueScanMacAdjacency"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
	WtpShare *string `pulumi:"wtpShare"`
}

type WirelessControllerGlobalState struct {
	// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
	ApLogServer pulumi.StringPtrInput
	// IP address that APs or FortiAPs send log messages to.
	ApLogServerIp pulumi.StringPtrInput
	// Port that APs or FortiAPs send log messages to.
	ApLogServerPort pulumi.IntPtrInput
	// Configure CAPWAP control message data channel offload.
	ControlMessageOffload pulumi.StringPtrInput
	DataEthernetIi        pulumi.StringPtrInput
	// Multicast IP address for AP discovery (default = 244.0.1.140).
	DiscoveryMcAddr pulumi.StringPtrInput
	// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
	FiappEthType pulumi.IntPtrInput
	// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
	ImageDownload pulumi.StringPtrInput
	// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
	IpsecBaseIp pulumi.StringPtrInput
	// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
	LinkAggregation pulumi.StringPtrInput
	// Description of the location of the wireless controller.
	Location pulumi.StringPtrInput
	// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
	MaxClients pulumi.IntPtrInput
	// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
	MaxRetransmit pulumi.IntPtrInput
	// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
	MeshEthType pulumi.IntPtrInput
	// Name of the wireless controller.
	Name pulumi.StringPtrInput
	// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
	RogueScanMacAdjacency pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
	WtpShare pulumi.StringPtrInput
}

func (WirelessControllerGlobalState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerGlobalState)(nil)).Elem()
}

type wirelessControllerGlobalArgs struct {
	// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
	ApLogServer *string `pulumi:"apLogServer"`
	// IP address that APs or FortiAPs send log messages to.
	ApLogServerIp *string `pulumi:"apLogServerIp"`
	// Port that APs or FortiAPs send log messages to.
	ApLogServerPort *int `pulumi:"apLogServerPort"`
	// Configure CAPWAP control message data channel offload.
	ControlMessageOffload *string `pulumi:"controlMessageOffload"`
	DataEthernetIi        *string `pulumi:"dataEthernetIi"`
	// Multicast IP address for AP discovery (default = 244.0.1.140).
	DiscoveryMcAddr *string `pulumi:"discoveryMcAddr"`
	// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
	FiappEthType *int `pulumi:"fiappEthType"`
	// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
	ImageDownload *string `pulumi:"imageDownload"`
	// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
	IpsecBaseIp *string `pulumi:"ipsecBaseIp"`
	// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
	LinkAggregation *string `pulumi:"linkAggregation"`
	// Description of the location of the wireless controller.
	Location *string `pulumi:"location"`
	// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
	MaxClients *int `pulumi:"maxClients"`
	// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
	MaxRetransmit *int `pulumi:"maxRetransmit"`
	// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
	MeshEthType *int `pulumi:"meshEthType"`
	// Name of the wireless controller.
	Name *string `pulumi:"name"`
	// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
	RogueScanMacAdjacency *int `pulumi:"rogueScanMacAdjacency"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
	WtpShare *string `pulumi:"wtpShare"`
}

// The set of arguments for constructing a WirelessControllerGlobal resource.
type WirelessControllerGlobalArgs struct {
	// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
	ApLogServer pulumi.StringPtrInput
	// IP address that APs or FortiAPs send log messages to.
	ApLogServerIp pulumi.StringPtrInput
	// Port that APs or FortiAPs send log messages to.
	ApLogServerPort pulumi.IntPtrInput
	// Configure CAPWAP control message data channel offload.
	ControlMessageOffload pulumi.StringPtrInput
	DataEthernetIi        pulumi.StringPtrInput
	// Multicast IP address for AP discovery (default = 244.0.1.140).
	DiscoveryMcAddr pulumi.StringPtrInput
	// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
	FiappEthType pulumi.IntPtrInput
	// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
	ImageDownload pulumi.StringPtrInput
	// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
	IpsecBaseIp pulumi.StringPtrInput
	// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
	LinkAggregation pulumi.StringPtrInput
	// Description of the location of the wireless controller.
	Location pulumi.StringPtrInput
	// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
	MaxClients pulumi.IntPtrInput
	// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
	MaxRetransmit pulumi.IntPtrInput
	// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
	MeshEthType pulumi.IntPtrInput
	// Name of the wireless controller.
	Name pulumi.StringPtrInput
	// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
	RogueScanMacAdjacency pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
	WtpShare pulumi.StringPtrInput
}

func (WirelessControllerGlobalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerGlobalArgs)(nil)).Elem()
}

type WirelessControllerGlobalInput interface {
	pulumi.Input

	ToWirelessControllerGlobalOutput() WirelessControllerGlobalOutput
	ToWirelessControllerGlobalOutputWithContext(ctx context.Context) WirelessControllerGlobalOutput
}

func (*WirelessControllerGlobal) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerGlobal)(nil))
}

func (i *WirelessControllerGlobal) ToWirelessControllerGlobalOutput() WirelessControllerGlobalOutput {
	return i.ToWirelessControllerGlobalOutputWithContext(context.Background())
}

func (i *WirelessControllerGlobal) ToWirelessControllerGlobalOutputWithContext(ctx context.Context) WirelessControllerGlobalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerGlobalOutput)
}

func (i *WirelessControllerGlobal) ToWirelessControllerGlobalPtrOutput() WirelessControllerGlobalPtrOutput {
	return i.ToWirelessControllerGlobalPtrOutputWithContext(context.Background())
}

func (i *WirelessControllerGlobal) ToWirelessControllerGlobalPtrOutputWithContext(ctx context.Context) WirelessControllerGlobalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerGlobalPtrOutput)
}

type WirelessControllerGlobalPtrInput interface {
	pulumi.Input

	ToWirelessControllerGlobalPtrOutput() WirelessControllerGlobalPtrOutput
	ToWirelessControllerGlobalPtrOutputWithContext(ctx context.Context) WirelessControllerGlobalPtrOutput
}

type wirelessControllerGlobalPtrType WirelessControllerGlobalArgs

func (*wirelessControllerGlobalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerGlobal)(nil))
}

func (i *wirelessControllerGlobalPtrType) ToWirelessControllerGlobalPtrOutput() WirelessControllerGlobalPtrOutput {
	return i.ToWirelessControllerGlobalPtrOutputWithContext(context.Background())
}

func (i *wirelessControllerGlobalPtrType) ToWirelessControllerGlobalPtrOutputWithContext(ctx context.Context) WirelessControllerGlobalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerGlobalPtrOutput)
}

// WirelessControllerGlobalArrayInput is an input type that accepts WirelessControllerGlobalArray and WirelessControllerGlobalArrayOutput values.
// You can construct a concrete instance of `WirelessControllerGlobalArrayInput` via:
//
//          WirelessControllerGlobalArray{ WirelessControllerGlobalArgs{...} }
type WirelessControllerGlobalArrayInput interface {
	pulumi.Input

	ToWirelessControllerGlobalArrayOutput() WirelessControllerGlobalArrayOutput
	ToWirelessControllerGlobalArrayOutputWithContext(context.Context) WirelessControllerGlobalArrayOutput
}

type WirelessControllerGlobalArray []WirelessControllerGlobalInput

func (WirelessControllerGlobalArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WirelessControllerGlobal)(nil))
}

func (i WirelessControllerGlobalArray) ToWirelessControllerGlobalArrayOutput() WirelessControllerGlobalArrayOutput {
	return i.ToWirelessControllerGlobalArrayOutputWithContext(context.Background())
}

func (i WirelessControllerGlobalArray) ToWirelessControllerGlobalArrayOutputWithContext(ctx context.Context) WirelessControllerGlobalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerGlobalArrayOutput)
}

// WirelessControllerGlobalMapInput is an input type that accepts WirelessControllerGlobalMap and WirelessControllerGlobalMapOutput values.
// You can construct a concrete instance of `WirelessControllerGlobalMapInput` via:
//
//          WirelessControllerGlobalMap{ "key": WirelessControllerGlobalArgs{...} }
type WirelessControllerGlobalMapInput interface {
	pulumi.Input

	ToWirelessControllerGlobalMapOutput() WirelessControllerGlobalMapOutput
	ToWirelessControllerGlobalMapOutputWithContext(context.Context) WirelessControllerGlobalMapOutput
}

type WirelessControllerGlobalMap map[string]WirelessControllerGlobalInput

func (WirelessControllerGlobalMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WirelessControllerGlobal)(nil))
}

func (i WirelessControllerGlobalMap) ToWirelessControllerGlobalMapOutput() WirelessControllerGlobalMapOutput {
	return i.ToWirelessControllerGlobalMapOutputWithContext(context.Background())
}

func (i WirelessControllerGlobalMap) ToWirelessControllerGlobalMapOutputWithContext(ctx context.Context) WirelessControllerGlobalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerGlobalMapOutput)
}

type WirelessControllerGlobalOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerGlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerGlobal)(nil))
}

func (o WirelessControllerGlobalOutput) ToWirelessControllerGlobalOutput() WirelessControllerGlobalOutput {
	return o
}

func (o WirelessControllerGlobalOutput) ToWirelessControllerGlobalOutputWithContext(ctx context.Context) WirelessControllerGlobalOutput {
	return o
}

func (o WirelessControllerGlobalOutput) ToWirelessControllerGlobalPtrOutput() WirelessControllerGlobalPtrOutput {
	return o.ToWirelessControllerGlobalPtrOutputWithContext(context.Background())
}

func (o WirelessControllerGlobalOutput) ToWirelessControllerGlobalPtrOutputWithContext(ctx context.Context) WirelessControllerGlobalPtrOutput {
	return o.ApplyT(func(v WirelessControllerGlobal) *WirelessControllerGlobal {
		return &v
	}).(WirelessControllerGlobalPtrOutput)
}

type WirelessControllerGlobalPtrOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerGlobalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerGlobal)(nil))
}

func (o WirelessControllerGlobalPtrOutput) ToWirelessControllerGlobalPtrOutput() WirelessControllerGlobalPtrOutput {
	return o
}

func (o WirelessControllerGlobalPtrOutput) ToWirelessControllerGlobalPtrOutputWithContext(ctx context.Context) WirelessControllerGlobalPtrOutput {
	return o
}

type WirelessControllerGlobalArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerGlobalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WirelessControllerGlobal)(nil))
}

func (o WirelessControllerGlobalArrayOutput) ToWirelessControllerGlobalArrayOutput() WirelessControllerGlobalArrayOutput {
	return o
}

func (o WirelessControllerGlobalArrayOutput) ToWirelessControllerGlobalArrayOutputWithContext(ctx context.Context) WirelessControllerGlobalArrayOutput {
	return o
}

func (o WirelessControllerGlobalArrayOutput) Index(i pulumi.IntInput) WirelessControllerGlobalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WirelessControllerGlobal {
		return vs[0].([]WirelessControllerGlobal)[vs[1].(int)]
	}).(WirelessControllerGlobalOutput)
}

type WirelessControllerGlobalMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerGlobalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WirelessControllerGlobal)(nil))
}

func (o WirelessControllerGlobalMapOutput) ToWirelessControllerGlobalMapOutput() WirelessControllerGlobalMapOutput {
	return o
}

func (o WirelessControllerGlobalMapOutput) ToWirelessControllerGlobalMapOutputWithContext(ctx context.Context) WirelessControllerGlobalMapOutput {
	return o
}

func (o WirelessControllerGlobalMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerGlobalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WirelessControllerGlobal {
		return vs[0].(map[string]WirelessControllerGlobal)[vs[1].(string)]
	}).(WirelessControllerGlobalOutput)
}

func init() {
	pulumi.RegisterOutputType(WirelessControllerGlobalOutput{})
	pulumi.RegisterOutputType(WirelessControllerGlobalPtrOutput{})
	pulumi.RegisterOutputType(WirelessControllerGlobalArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerGlobalMapOutput{})
}
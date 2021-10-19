// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure CAPWAP timers.
//
// ## Import
//
// WirelessController Timers can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerTimers:WirelessControllerTimers labelname WirelessControllerTimers
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerTimers struct {
	pulumi.CustomResourceState

	// Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
	BleScanReportIntv pulumi.IntOutput `pulumi:"bleScanReportIntv"`
	// Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
	ClientIdleTimeout pulumi.IntOutput `pulumi:"clientIdleTimeout"`
	// Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	DarrpDay pulumi.StringOutput `pulumi:"darrpDay"`
	// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
	DarrpOptimize pulumi.IntOutput `pulumi:"darrpOptimize"`
	// Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrpTime` block is documented below.
	DarrpTimes WirelessControllerTimersDarrpTimeArrayOutput `pulumi:"darrpTimes"`
	// Time between discovery requests (2 - 180 sec, default = 5).
	DiscoveryInterval pulumi.IntOutput `pulumi:"discoveryInterval"`
	// Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
	DrmaInterval pulumi.IntOutput `pulumi:"drmaInterval"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
	EchoInterval pulumi.IntOutput `pulumi:"echoInterval"`
	// Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
	FakeApLog pulumi.IntOutput `pulumi:"fakeApLog"`
	// Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
	IpsecIntfCleanup pulumi.IntOutput `pulumi:"ipsecIntfCleanup"`
	// Time between running radio reports (1 - 255 sec, default = 15).
	RadioStatsInterval pulumi.IntOutput `pulumi:"radioStatsInterval"`
	// Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
	RogueApLog pulumi.IntOutput `pulumi:"rogueApLog"`
	// Time between running station capability reports (1 - 255 sec, default = 30).
	StaCapabilityInterval pulumi.IntOutput `pulumi:"staCapabilityInterval"`
	// Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
	StaLocateTimer pulumi.IntOutput `pulumi:"staLocateTimer"`
	// Time between running client (station) reports (1 - 255 sec, default = 1).
	StaStatsInterval pulumi.IntOutput `pulumi:"staStatsInterval"`
	// Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
	VapStatsInterval pulumi.IntOutput `pulumi:"vapStatsInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerTimers registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerTimers(ctx *pulumi.Context,
	name string, args *WirelessControllerTimersArgs, opts ...pulumi.ResourceOption) (*WirelessControllerTimers, error) {
	if args == nil {
		args = &WirelessControllerTimersArgs{}
	}

	var resource WirelessControllerTimers
	err := ctx.RegisterResource("fortios:index/wirelessControllerTimers:WirelessControllerTimers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerTimers gets an existing WirelessControllerTimers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerTimers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerTimersState, opts ...pulumi.ResourceOption) (*WirelessControllerTimers, error) {
	var resource WirelessControllerTimers
	err := ctx.ReadResource("fortios:index/wirelessControllerTimers:WirelessControllerTimers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerTimers resources.
type wirelessControllerTimersState struct {
	// Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
	BleScanReportIntv *int `pulumi:"bleScanReportIntv"`
	// Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
	ClientIdleTimeout *int `pulumi:"clientIdleTimeout"`
	// Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	DarrpDay *string `pulumi:"darrpDay"`
	// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
	DarrpOptimize *int `pulumi:"darrpOptimize"`
	// Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrpTime` block is documented below.
	DarrpTimes []WirelessControllerTimersDarrpTime `pulumi:"darrpTimes"`
	// Time between discovery requests (2 - 180 sec, default = 5).
	DiscoveryInterval *int `pulumi:"discoveryInterval"`
	// Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
	DrmaInterval *int `pulumi:"drmaInterval"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
	EchoInterval *int `pulumi:"echoInterval"`
	// Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
	FakeApLog *int `pulumi:"fakeApLog"`
	// Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
	IpsecIntfCleanup *int `pulumi:"ipsecIntfCleanup"`
	// Time between running radio reports (1 - 255 sec, default = 15).
	RadioStatsInterval *int `pulumi:"radioStatsInterval"`
	// Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
	RogueApLog *int `pulumi:"rogueApLog"`
	// Time between running station capability reports (1 - 255 sec, default = 30).
	StaCapabilityInterval *int `pulumi:"staCapabilityInterval"`
	// Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
	StaLocateTimer *int `pulumi:"staLocateTimer"`
	// Time between running client (station) reports (1 - 255 sec, default = 1).
	StaStatsInterval *int `pulumi:"staStatsInterval"`
	// Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
	VapStatsInterval *int `pulumi:"vapStatsInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerTimersState struct {
	// Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
	BleScanReportIntv pulumi.IntPtrInput
	// Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
	ClientIdleTimeout pulumi.IntPtrInput
	// Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	DarrpDay pulumi.StringPtrInput
	// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
	DarrpOptimize pulumi.IntPtrInput
	// Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrpTime` block is documented below.
	DarrpTimes WirelessControllerTimersDarrpTimeArrayInput
	// Time between discovery requests (2 - 180 sec, default = 5).
	DiscoveryInterval pulumi.IntPtrInput
	// Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
	DrmaInterval pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
	EchoInterval pulumi.IntPtrInput
	// Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
	FakeApLog pulumi.IntPtrInput
	// Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
	IpsecIntfCleanup pulumi.IntPtrInput
	// Time between running radio reports (1 - 255 sec, default = 15).
	RadioStatsInterval pulumi.IntPtrInput
	// Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
	RogueApLog pulumi.IntPtrInput
	// Time between running station capability reports (1 - 255 sec, default = 30).
	StaCapabilityInterval pulumi.IntPtrInput
	// Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
	StaLocateTimer pulumi.IntPtrInput
	// Time between running client (station) reports (1 - 255 sec, default = 1).
	StaStatsInterval pulumi.IntPtrInput
	// Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
	VapStatsInterval pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerTimersState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerTimersState)(nil)).Elem()
}

type wirelessControllerTimersArgs struct {
	// Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
	BleScanReportIntv *int `pulumi:"bleScanReportIntv"`
	// Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
	ClientIdleTimeout *int `pulumi:"clientIdleTimeout"`
	// Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	DarrpDay *string `pulumi:"darrpDay"`
	// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
	DarrpOptimize *int `pulumi:"darrpOptimize"`
	// Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrpTime` block is documented below.
	DarrpTimes []WirelessControllerTimersDarrpTime `pulumi:"darrpTimes"`
	// Time between discovery requests (2 - 180 sec, default = 5).
	DiscoveryInterval *int `pulumi:"discoveryInterval"`
	// Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
	DrmaInterval *int `pulumi:"drmaInterval"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
	EchoInterval *int `pulumi:"echoInterval"`
	// Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
	FakeApLog *int `pulumi:"fakeApLog"`
	// Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
	IpsecIntfCleanup *int `pulumi:"ipsecIntfCleanup"`
	// Time between running radio reports (1 - 255 sec, default = 15).
	RadioStatsInterval *int `pulumi:"radioStatsInterval"`
	// Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
	RogueApLog *int `pulumi:"rogueApLog"`
	// Time between running station capability reports (1 - 255 sec, default = 30).
	StaCapabilityInterval *int `pulumi:"staCapabilityInterval"`
	// Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
	StaLocateTimer *int `pulumi:"staLocateTimer"`
	// Time between running client (station) reports (1 - 255 sec, default = 1).
	StaStatsInterval *int `pulumi:"staStatsInterval"`
	// Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
	VapStatsInterval *int `pulumi:"vapStatsInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerTimers resource.
type WirelessControllerTimersArgs struct {
	// Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
	BleScanReportIntv pulumi.IntPtrInput
	// Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
	ClientIdleTimeout pulumi.IntPtrInput
	// Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
	DarrpDay pulumi.StringPtrInput
	// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
	DarrpOptimize pulumi.IntPtrInput
	// Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrpTime` block is documented below.
	DarrpTimes WirelessControllerTimersDarrpTimeArrayInput
	// Time between discovery requests (2 - 180 sec, default = 5).
	DiscoveryInterval pulumi.IntPtrInput
	// Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
	DrmaInterval pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
	EchoInterval pulumi.IntPtrInput
	// Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
	FakeApLog pulumi.IntPtrInput
	// Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
	IpsecIntfCleanup pulumi.IntPtrInput
	// Time between running radio reports (1 - 255 sec, default = 15).
	RadioStatsInterval pulumi.IntPtrInput
	// Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
	RogueApLog pulumi.IntPtrInput
	// Time between running station capability reports (1 - 255 sec, default = 30).
	StaCapabilityInterval pulumi.IntPtrInput
	// Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
	StaLocateTimer pulumi.IntPtrInput
	// Time between running client (station) reports (1 - 255 sec, default = 1).
	StaStatsInterval pulumi.IntPtrInput
	// Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
	VapStatsInterval pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerTimersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerTimersArgs)(nil)).Elem()
}

type WirelessControllerTimersInput interface {
	pulumi.Input

	ToWirelessControllerTimersOutput() WirelessControllerTimersOutput
	ToWirelessControllerTimersOutputWithContext(ctx context.Context) WirelessControllerTimersOutput
}

func (*WirelessControllerTimers) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerTimers)(nil))
}

func (i *WirelessControllerTimers) ToWirelessControllerTimersOutput() WirelessControllerTimersOutput {
	return i.ToWirelessControllerTimersOutputWithContext(context.Background())
}

func (i *WirelessControllerTimers) ToWirelessControllerTimersOutputWithContext(ctx context.Context) WirelessControllerTimersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerTimersOutput)
}

func (i *WirelessControllerTimers) ToWirelessControllerTimersPtrOutput() WirelessControllerTimersPtrOutput {
	return i.ToWirelessControllerTimersPtrOutputWithContext(context.Background())
}

func (i *WirelessControllerTimers) ToWirelessControllerTimersPtrOutputWithContext(ctx context.Context) WirelessControllerTimersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerTimersPtrOutput)
}

type WirelessControllerTimersPtrInput interface {
	pulumi.Input

	ToWirelessControllerTimersPtrOutput() WirelessControllerTimersPtrOutput
	ToWirelessControllerTimersPtrOutputWithContext(ctx context.Context) WirelessControllerTimersPtrOutput
}

type wirelessControllerTimersPtrType WirelessControllerTimersArgs

func (*wirelessControllerTimersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerTimers)(nil))
}

func (i *wirelessControllerTimersPtrType) ToWirelessControllerTimersPtrOutput() WirelessControllerTimersPtrOutput {
	return i.ToWirelessControllerTimersPtrOutputWithContext(context.Background())
}

func (i *wirelessControllerTimersPtrType) ToWirelessControllerTimersPtrOutputWithContext(ctx context.Context) WirelessControllerTimersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerTimersPtrOutput)
}

// WirelessControllerTimersArrayInput is an input type that accepts WirelessControllerTimersArray and WirelessControllerTimersArrayOutput values.
// You can construct a concrete instance of `WirelessControllerTimersArrayInput` via:
//
//          WirelessControllerTimersArray{ WirelessControllerTimersArgs{...} }
type WirelessControllerTimersArrayInput interface {
	pulumi.Input

	ToWirelessControllerTimersArrayOutput() WirelessControllerTimersArrayOutput
	ToWirelessControllerTimersArrayOutputWithContext(context.Context) WirelessControllerTimersArrayOutput
}

type WirelessControllerTimersArray []WirelessControllerTimersInput

func (WirelessControllerTimersArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WirelessControllerTimers)(nil))
}

func (i WirelessControllerTimersArray) ToWirelessControllerTimersArrayOutput() WirelessControllerTimersArrayOutput {
	return i.ToWirelessControllerTimersArrayOutputWithContext(context.Background())
}

func (i WirelessControllerTimersArray) ToWirelessControllerTimersArrayOutputWithContext(ctx context.Context) WirelessControllerTimersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerTimersArrayOutput)
}

// WirelessControllerTimersMapInput is an input type that accepts WirelessControllerTimersMap and WirelessControllerTimersMapOutput values.
// You can construct a concrete instance of `WirelessControllerTimersMapInput` via:
//
//          WirelessControllerTimersMap{ "key": WirelessControllerTimersArgs{...} }
type WirelessControllerTimersMapInput interface {
	pulumi.Input

	ToWirelessControllerTimersMapOutput() WirelessControllerTimersMapOutput
	ToWirelessControllerTimersMapOutputWithContext(context.Context) WirelessControllerTimersMapOutput
}

type WirelessControllerTimersMap map[string]WirelessControllerTimersInput

func (WirelessControllerTimersMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WirelessControllerTimers)(nil))
}

func (i WirelessControllerTimersMap) ToWirelessControllerTimersMapOutput() WirelessControllerTimersMapOutput {
	return i.ToWirelessControllerTimersMapOutputWithContext(context.Background())
}

func (i WirelessControllerTimersMap) ToWirelessControllerTimersMapOutputWithContext(ctx context.Context) WirelessControllerTimersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerTimersMapOutput)
}

type WirelessControllerTimersOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerTimersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerTimers)(nil))
}

func (o WirelessControllerTimersOutput) ToWirelessControllerTimersOutput() WirelessControllerTimersOutput {
	return o
}

func (o WirelessControllerTimersOutput) ToWirelessControllerTimersOutputWithContext(ctx context.Context) WirelessControllerTimersOutput {
	return o
}

func (o WirelessControllerTimersOutput) ToWirelessControllerTimersPtrOutput() WirelessControllerTimersPtrOutput {
	return o.ToWirelessControllerTimersPtrOutputWithContext(context.Background())
}

func (o WirelessControllerTimersOutput) ToWirelessControllerTimersPtrOutputWithContext(ctx context.Context) WirelessControllerTimersPtrOutput {
	return o.ApplyT(func(v WirelessControllerTimers) *WirelessControllerTimers {
		return &v
	}).(WirelessControllerTimersPtrOutput)
}

type WirelessControllerTimersPtrOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerTimersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerTimers)(nil))
}

func (o WirelessControllerTimersPtrOutput) ToWirelessControllerTimersPtrOutput() WirelessControllerTimersPtrOutput {
	return o
}

func (o WirelessControllerTimersPtrOutput) ToWirelessControllerTimersPtrOutputWithContext(ctx context.Context) WirelessControllerTimersPtrOutput {
	return o
}

type WirelessControllerTimersArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerTimersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WirelessControllerTimers)(nil))
}

func (o WirelessControllerTimersArrayOutput) ToWirelessControllerTimersArrayOutput() WirelessControllerTimersArrayOutput {
	return o
}

func (o WirelessControllerTimersArrayOutput) ToWirelessControllerTimersArrayOutputWithContext(ctx context.Context) WirelessControllerTimersArrayOutput {
	return o
}

func (o WirelessControllerTimersArrayOutput) Index(i pulumi.IntInput) WirelessControllerTimersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WirelessControllerTimers {
		return vs[0].([]WirelessControllerTimers)[vs[1].(int)]
	}).(WirelessControllerTimersOutput)
}

type WirelessControllerTimersMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerTimersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WirelessControllerTimers)(nil))
}

func (o WirelessControllerTimersMapOutput) ToWirelessControllerTimersMapOutput() WirelessControllerTimersMapOutput {
	return o
}

func (o WirelessControllerTimersMapOutput) ToWirelessControllerTimersMapOutputWithContext(ctx context.Context) WirelessControllerTimersMapOutput {
	return o
}

func (o WirelessControllerTimersMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerTimersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WirelessControllerTimers {
		return vs[0].(map[string]WirelessControllerTimers)[vs[1].(string)]
	}).(WirelessControllerTimersOutput)
}

func init() {
	pulumi.RegisterOutputType(WirelessControllerTimersOutput{})
	pulumi.RegisterOutputType(WirelessControllerTimersPtrOutput{})
	pulumi.RegisterOutputType(WirelessControllerTimersArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerTimersMapOutput{})
}

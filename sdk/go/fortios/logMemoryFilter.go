// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Filters for memory buffer.
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
// 		_, err := fortios.NewLogMemoryFilter(ctx, "trname", &fortios.LogMemoryFilterArgs{
// 			Anomaly:          pulumi.String("enable"),
// 			Dns:              pulumi.String("enable"),
// 			FilterType:       pulumi.String("include"),
// 			ForwardTraffic:   pulumi.String("enable"),
// 			Gtp:              pulumi.String("enable"),
// 			LocalTraffic:     pulumi.String("enable"),
// 			MulticastTraffic: pulumi.String("enable"),
// 			Severity:         pulumi.String("information"),
// 			SnifferTraffic:   pulumi.String("enable"),
// 			Ssh:              pulumi.String("enable"),
// 			Voip:             pulumi.String("enable"),
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
// LogMemory Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logMemoryFilter:LogMemoryFilter labelname LogMemoryFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogMemoryFilter struct {
	pulumi.CustomResourceState

	// Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
	Admin pulumi.StringOutput `pulumi:"admin"`
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringOutput `pulumi:"anomaly"`
	// Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
	Auth pulumi.StringOutput `pulumi:"auth"`
	// Enable/disable CPU & memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
	CpuMemoryUsage pulumi.StringOutput `pulumi:"cpuMemoryUsage"`
	// Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
	Dhcp pulumi.StringOutput `pulumi:"dhcp"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringOutput `pulumi:"dns"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable event logging. Valid values: `enable`, `disable`.
	Event pulumi.StringOutput `pulumi:"event"`
	// Free style filter string.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringOutput `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringOutput `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles LogMemoryFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringOutput `pulumi:"gtp"`
	// Enable/disable HA logging. Valid values: `enable`, `disable`.
	Ha pulumi.StringOutput `pulumi:"ha"`
	// Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
	Ipsec pulumi.StringOutput `pulumi:"ipsec"`
	// Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
	LdbMonitor pulumi.StringOutput `pulumi:"ldbMonitor"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringOutput `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringOutput `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringOutput `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringOutput `pulumi:"netscanVulnerability"`
	// Enable/disable pattern update logging. Valid values: `enable`, `disable`.
	Pattern pulumi.StringOutput `pulumi:"pattern"`
	// Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
	Ppp pulumi.StringOutput `pulumi:"ppp"`
	// Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
	Radius pulumi.StringOutput `pulumi:"radius"`
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringOutput `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringOutput `pulumi:"ssh"`
	// Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
	SslvpnLogAdm pulumi.StringOutput `pulumi:"sslvpnLogAdm"`
	// Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
	SslvpnLogAuth pulumi.StringOutput `pulumi:"sslvpnLogAuth"`
	// Enable/disable SSL session logging. Valid values: `enable`, `disable`.
	SslvpnLogSession pulumi.StringOutput `pulumi:"sslvpnLogSession"`
	// Enable/disable system activity logging. Valid values: `enable`, `disable`.
	System pulumi.StringOutput `pulumi:"system"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
	VipSsl pulumi.StringOutput `pulumi:"vipSsl"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringOutput `pulumi:"voip"`
	// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
	WanOpt pulumi.StringOutput `pulumi:"wanOpt"`
	// Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
	WirelessActivity pulumi.StringOutput `pulumi:"wirelessActivity"`
}

// NewLogMemoryFilter registers a new resource with the given unique name, arguments, and options.
func NewLogMemoryFilter(ctx *pulumi.Context,
	name string, args *LogMemoryFilterArgs, opts ...pulumi.ResourceOption) (*LogMemoryFilter, error) {
	if args == nil {
		args = &LogMemoryFilterArgs{}
	}

	var resource LogMemoryFilter
	err := ctx.RegisterResource("fortios:index/logMemoryFilter:LogMemoryFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogMemoryFilter gets an existing LogMemoryFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogMemoryFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogMemoryFilterState, opts ...pulumi.ResourceOption) (*LogMemoryFilter, error) {
	var resource LogMemoryFilter
	err := ctx.ReadResource("fortios:index/logMemoryFilter:LogMemoryFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogMemoryFilter resources.
type logMemoryFilterState struct {
	// Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
	Admin *string `pulumi:"admin"`
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
	Auth *string `pulumi:"auth"`
	// Enable/disable CPU & memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
	CpuMemoryUsage *string `pulumi:"cpuMemoryUsage"`
	// Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
	Dhcp *string `pulumi:"dhcp"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns *string `pulumi:"dns"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable event logging. Valid values: `enable`, `disable`.
	Event *string `pulumi:"event"`
	// Free style filter string.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []LogMemoryFilterFreeStyle `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp *string `pulumi:"gtp"`
	// Enable/disable HA logging. Valid values: `enable`, `disable`.
	Ha *string `pulumi:"ha"`
	// Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
	Ipsec *string `pulumi:"ipsec"`
	// Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
	LdbMonitor *string `pulumi:"ldbMonitor"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic *string `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic *string `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery *string `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability *string `pulumi:"netscanVulnerability"`
	// Enable/disable pattern update logging. Valid values: `enable`, `disable`.
	Pattern *string `pulumi:"pattern"`
	// Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
	Ppp *string `pulumi:"ppp"`
	// Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
	Radius *string `pulumi:"radius"`
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic *string `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh *string `pulumi:"ssh"`
	// Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
	SslvpnLogAdm *string `pulumi:"sslvpnLogAdm"`
	// Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
	SslvpnLogAuth *string `pulumi:"sslvpnLogAuth"`
	// Enable/disable SSL session logging. Valid values: `enable`, `disable`.
	SslvpnLogSession *string `pulumi:"sslvpnLogSession"`
	// Enable/disable system activity logging. Valid values: `enable`, `disable`.
	System *string `pulumi:"system"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
	VipSsl *string `pulumi:"vipSsl"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip *string `pulumi:"voip"`
	// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
	WanOpt *string `pulumi:"wanOpt"`
	// Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
	WirelessActivity *string `pulumi:"wirelessActivity"`
}

type LogMemoryFilterState struct {
	// Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
	Admin pulumi.StringPtrInput
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
	Auth pulumi.StringPtrInput
	// Enable/disable CPU & memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
	CpuMemoryUsage pulumi.StringPtrInput
	// Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
	Dhcp pulumi.StringPtrInput
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable event logging. Valid values: `enable`, `disable`.
	Event pulumi.StringPtrInput
	// Free style filter string.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles LogMemoryFilterFreeStyleArrayInput
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringPtrInput
	// Enable/disable HA logging. Valid values: `enable`, `disable`.
	Ha pulumi.StringPtrInput
	// Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
	Ipsec pulumi.StringPtrInput
	// Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
	LdbMonitor pulumi.StringPtrInput
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringPtrInput
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringPtrInput
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringPtrInput
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringPtrInput
	// Enable/disable pattern update logging. Valid values: `enable`, `disable`.
	Pattern pulumi.StringPtrInput
	// Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
	Ppp pulumi.StringPtrInput
	// Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
	Radius pulumi.StringPtrInput
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringPtrInput
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringPtrInput
	// Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
	SslvpnLogAdm pulumi.StringPtrInput
	// Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
	SslvpnLogAuth pulumi.StringPtrInput
	// Enable/disable SSL session logging. Valid values: `enable`, `disable`.
	SslvpnLogSession pulumi.StringPtrInput
	// Enable/disable system activity logging. Valid values: `enable`, `disable`.
	System pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
	VipSsl pulumi.StringPtrInput
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringPtrInput
	// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
	WanOpt pulumi.StringPtrInput
	// Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
	WirelessActivity pulumi.StringPtrInput
}

func (LogMemoryFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logMemoryFilterState)(nil)).Elem()
}

type logMemoryFilterArgs struct {
	// Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
	Admin *string `pulumi:"admin"`
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
	Auth *string `pulumi:"auth"`
	// Enable/disable CPU & memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
	CpuMemoryUsage *string `pulumi:"cpuMemoryUsage"`
	// Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
	Dhcp *string `pulumi:"dhcp"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns *string `pulumi:"dns"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable event logging. Valid values: `enable`, `disable`.
	Event *string `pulumi:"event"`
	// Free style filter string.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []LogMemoryFilterFreeStyle `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp *string `pulumi:"gtp"`
	// Enable/disable HA logging. Valid values: `enable`, `disable`.
	Ha *string `pulumi:"ha"`
	// Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
	Ipsec *string `pulumi:"ipsec"`
	// Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
	LdbMonitor *string `pulumi:"ldbMonitor"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic *string `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic *string `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery *string `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability *string `pulumi:"netscanVulnerability"`
	// Enable/disable pattern update logging. Valid values: `enable`, `disable`.
	Pattern *string `pulumi:"pattern"`
	// Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
	Ppp *string `pulumi:"ppp"`
	// Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
	Radius *string `pulumi:"radius"`
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic *string `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh *string `pulumi:"ssh"`
	// Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
	SslvpnLogAdm *string `pulumi:"sslvpnLogAdm"`
	// Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
	SslvpnLogAuth *string `pulumi:"sslvpnLogAuth"`
	// Enable/disable SSL session logging. Valid values: `enable`, `disable`.
	SslvpnLogSession *string `pulumi:"sslvpnLogSession"`
	// Enable/disable system activity logging. Valid values: `enable`, `disable`.
	System *string `pulumi:"system"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
	VipSsl *string `pulumi:"vipSsl"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip *string `pulumi:"voip"`
	// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
	WanOpt *string `pulumi:"wanOpt"`
	// Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
	WirelessActivity *string `pulumi:"wirelessActivity"`
}

// The set of arguments for constructing a LogMemoryFilter resource.
type LogMemoryFilterArgs struct {
	// Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
	Admin pulumi.StringPtrInput
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
	Auth pulumi.StringPtrInput
	// Enable/disable CPU & memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
	CpuMemoryUsage pulumi.StringPtrInput
	// Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
	Dhcp pulumi.StringPtrInput
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable event logging. Valid values: `enable`, `disable`.
	Event pulumi.StringPtrInput
	// Free style filter string.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles LogMemoryFilterFreeStyleArrayInput
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringPtrInput
	// Enable/disable HA logging. Valid values: `enable`, `disable`.
	Ha pulumi.StringPtrInput
	// Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
	Ipsec pulumi.StringPtrInput
	// Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
	LdbMonitor pulumi.StringPtrInput
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringPtrInput
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringPtrInput
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringPtrInput
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringPtrInput
	// Enable/disable pattern update logging. Valid values: `enable`, `disable`.
	Pattern pulumi.StringPtrInput
	// Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
	Ppp pulumi.StringPtrInput
	// Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
	Radius pulumi.StringPtrInput
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringPtrInput
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringPtrInput
	// Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
	SslvpnLogAdm pulumi.StringPtrInput
	// Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
	SslvpnLogAuth pulumi.StringPtrInput
	// Enable/disable SSL session logging. Valid values: `enable`, `disable`.
	SslvpnLogSession pulumi.StringPtrInput
	// Enable/disable system activity logging. Valid values: `enable`, `disable`.
	System pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
	VipSsl pulumi.StringPtrInput
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringPtrInput
	// Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
	WanOpt pulumi.StringPtrInput
	// Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
	WirelessActivity pulumi.StringPtrInput
}

func (LogMemoryFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logMemoryFilterArgs)(nil)).Elem()
}

type LogMemoryFilterInput interface {
	pulumi.Input

	ToLogMemoryFilterOutput() LogMemoryFilterOutput
	ToLogMemoryFilterOutputWithContext(ctx context.Context) LogMemoryFilterOutput
}

func (*LogMemoryFilter) ElementType() reflect.Type {
	return reflect.TypeOf((*LogMemoryFilter)(nil))
}

func (i *LogMemoryFilter) ToLogMemoryFilterOutput() LogMemoryFilterOutput {
	return i.ToLogMemoryFilterOutputWithContext(context.Background())
}

func (i *LogMemoryFilter) ToLogMemoryFilterOutputWithContext(ctx context.Context) LogMemoryFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMemoryFilterOutput)
}

func (i *LogMemoryFilter) ToLogMemoryFilterPtrOutput() LogMemoryFilterPtrOutput {
	return i.ToLogMemoryFilterPtrOutputWithContext(context.Background())
}

func (i *LogMemoryFilter) ToLogMemoryFilterPtrOutputWithContext(ctx context.Context) LogMemoryFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMemoryFilterPtrOutput)
}

type LogMemoryFilterPtrInput interface {
	pulumi.Input

	ToLogMemoryFilterPtrOutput() LogMemoryFilterPtrOutput
	ToLogMemoryFilterPtrOutputWithContext(ctx context.Context) LogMemoryFilterPtrOutput
}

type logMemoryFilterPtrType LogMemoryFilterArgs

func (*logMemoryFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogMemoryFilter)(nil))
}

func (i *logMemoryFilterPtrType) ToLogMemoryFilterPtrOutput() LogMemoryFilterPtrOutput {
	return i.ToLogMemoryFilterPtrOutputWithContext(context.Background())
}

func (i *logMemoryFilterPtrType) ToLogMemoryFilterPtrOutputWithContext(ctx context.Context) LogMemoryFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMemoryFilterPtrOutput)
}

// LogMemoryFilterArrayInput is an input type that accepts LogMemoryFilterArray and LogMemoryFilterArrayOutput values.
// You can construct a concrete instance of `LogMemoryFilterArrayInput` via:
//
//          LogMemoryFilterArray{ LogMemoryFilterArgs{...} }
type LogMemoryFilterArrayInput interface {
	pulumi.Input

	ToLogMemoryFilterArrayOutput() LogMemoryFilterArrayOutput
	ToLogMemoryFilterArrayOutputWithContext(context.Context) LogMemoryFilterArrayOutput
}

type LogMemoryFilterArray []LogMemoryFilterInput

func (LogMemoryFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogMemoryFilter)(nil))
}

func (i LogMemoryFilterArray) ToLogMemoryFilterArrayOutput() LogMemoryFilterArrayOutput {
	return i.ToLogMemoryFilterArrayOutputWithContext(context.Background())
}

func (i LogMemoryFilterArray) ToLogMemoryFilterArrayOutputWithContext(ctx context.Context) LogMemoryFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMemoryFilterArrayOutput)
}

// LogMemoryFilterMapInput is an input type that accepts LogMemoryFilterMap and LogMemoryFilterMapOutput values.
// You can construct a concrete instance of `LogMemoryFilterMapInput` via:
//
//          LogMemoryFilterMap{ "key": LogMemoryFilterArgs{...} }
type LogMemoryFilterMapInput interface {
	pulumi.Input

	ToLogMemoryFilterMapOutput() LogMemoryFilterMapOutput
	ToLogMemoryFilterMapOutputWithContext(context.Context) LogMemoryFilterMapOutput
}

type LogMemoryFilterMap map[string]LogMemoryFilterInput

func (LogMemoryFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogMemoryFilter)(nil))
}

func (i LogMemoryFilterMap) ToLogMemoryFilterMapOutput() LogMemoryFilterMapOutput {
	return i.ToLogMemoryFilterMapOutputWithContext(context.Background())
}

func (i LogMemoryFilterMap) ToLogMemoryFilterMapOutputWithContext(ctx context.Context) LogMemoryFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMemoryFilterMapOutput)
}

type LogMemoryFilterOutput struct {
	*pulumi.OutputState
}

func (LogMemoryFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogMemoryFilter)(nil))
}

func (o LogMemoryFilterOutput) ToLogMemoryFilterOutput() LogMemoryFilterOutput {
	return o
}

func (o LogMemoryFilterOutput) ToLogMemoryFilterOutputWithContext(ctx context.Context) LogMemoryFilterOutput {
	return o
}

func (o LogMemoryFilterOutput) ToLogMemoryFilterPtrOutput() LogMemoryFilterPtrOutput {
	return o.ToLogMemoryFilterPtrOutputWithContext(context.Background())
}

func (o LogMemoryFilterOutput) ToLogMemoryFilterPtrOutputWithContext(ctx context.Context) LogMemoryFilterPtrOutput {
	return o.ApplyT(func(v LogMemoryFilter) *LogMemoryFilter {
		return &v
	}).(LogMemoryFilterPtrOutput)
}

type LogMemoryFilterPtrOutput struct {
	*pulumi.OutputState
}

func (LogMemoryFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogMemoryFilter)(nil))
}

func (o LogMemoryFilterPtrOutput) ToLogMemoryFilterPtrOutput() LogMemoryFilterPtrOutput {
	return o
}

func (o LogMemoryFilterPtrOutput) ToLogMemoryFilterPtrOutputWithContext(ctx context.Context) LogMemoryFilterPtrOutput {
	return o
}

type LogMemoryFilterArrayOutput struct{ *pulumi.OutputState }

func (LogMemoryFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogMemoryFilter)(nil))
}

func (o LogMemoryFilterArrayOutput) ToLogMemoryFilterArrayOutput() LogMemoryFilterArrayOutput {
	return o
}

func (o LogMemoryFilterArrayOutput) ToLogMemoryFilterArrayOutputWithContext(ctx context.Context) LogMemoryFilterArrayOutput {
	return o
}

func (o LogMemoryFilterArrayOutput) Index(i pulumi.IntInput) LogMemoryFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogMemoryFilter {
		return vs[0].([]LogMemoryFilter)[vs[1].(int)]
	}).(LogMemoryFilterOutput)
}

type LogMemoryFilterMapOutput struct{ *pulumi.OutputState }

func (LogMemoryFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogMemoryFilter)(nil))
}

func (o LogMemoryFilterMapOutput) ToLogMemoryFilterMapOutput() LogMemoryFilterMapOutput {
	return o
}

func (o LogMemoryFilterMapOutput) ToLogMemoryFilterMapOutputWithContext(ctx context.Context) LogMemoryFilterMapOutput {
	return o
}

func (o LogMemoryFilterMapOutput) MapIndex(k pulumi.StringInput) LogMemoryFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogMemoryFilter {
		return vs[0].(map[string]LogMemoryFilter)[vs[1].(string)]
	}).(LogMemoryFilterOutput)
}

func init() {
	pulumi.RegisterOutputType(LogMemoryFilterOutput{})
	pulumi.RegisterOutputType(LogMemoryFilterPtrOutput{})
	pulumi.RegisterOutputType(LogMemoryFilterArrayOutput{})
	pulumi.RegisterOutputType(LogMemoryFilterMapOutput{})
}
// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Override filters for remote system server.
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
// 		_, err := fortios.NewLogSyslogd4OverrideFilter(ctx, "trname", &fortios.LogSyslogd4OverrideFilterArgs{
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
// LogSyslogd4 OverrideFilter can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd4OverrideFilter:LogSyslogd4OverrideFilter labelname LogSyslogd4OverrideFilter
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd4OverrideFilter:LogSyslogd4OverrideFilter labelname LogSyslogd4OverrideFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSyslogd4OverrideFilter struct {
	pulumi.CustomResourceState

	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringOutput `pulumi:"anomaly"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringOutput `pulumi:"dns"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Free style filter string.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringOutput `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringOutput `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles LogSyslogd4OverrideFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringOutput `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringOutput `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringOutput `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringOutput `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringOutput `pulumi:"netscanVulnerability"`
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringOutput `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringOutput `pulumi:"ssh"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringOutput `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringOutput `pulumi:"ztnaTraffic"`
}

// NewLogSyslogd4OverrideFilter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd4OverrideFilter(ctx *pulumi.Context,
	name string, args *LogSyslogd4OverrideFilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogd4OverrideFilter, error) {
	if args == nil {
		args = &LogSyslogd4OverrideFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogd4OverrideFilter
	err := ctx.RegisterResource("fortios:index/logSyslogd4OverrideFilter:LogSyslogd4OverrideFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogd4OverrideFilter gets an existing LogSyslogd4OverrideFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogd4OverrideFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogd4OverrideFilterState, opts ...pulumi.ResourceOption) (*LogSyslogd4OverrideFilter, error) {
	var resource LogSyslogd4OverrideFilter
	err := ctx.ReadResource("fortios:index/logSyslogd4OverrideFilter:LogSyslogd4OverrideFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogd4OverrideFilter resources.
type logSyslogd4OverrideFilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns *string `pulumi:"dns"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Free style filter string.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []LogSyslogd4OverrideFilterFreeStyle `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp *string `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic *string `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic *string `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery *string `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability *string `pulumi:"netscanVulnerability"`
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic *string `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh *string `pulumi:"ssh"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip *string `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic *string `pulumi:"ztnaTraffic"`
}

type LogSyslogd4OverrideFilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Free style filter string.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles LogSyslogd4OverrideFilterFreeStyleArrayInput
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringPtrInput
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringPtrInput
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringPtrInput
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringPtrInput
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringPtrInput
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringPtrInput
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringPtrInput
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringPtrInput
}

func (LogSyslogd4OverrideFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd4OverrideFilterState)(nil)).Elem()
}

type logSyslogd4OverrideFilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns *string `pulumi:"dns"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Free style filter string.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []LogSyslogd4OverrideFilterFreeStyle `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp *string `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic *string `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic *string `pulumi:"multicastTraffic"`
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery *string `pulumi:"netscanDiscovery"`
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability *string `pulumi:"netscanVulnerability"`
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic *string `pulumi:"snifferTraffic"`
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh *string `pulumi:"ssh"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip *string `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic *string `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a LogSyslogd4OverrideFilter resource.
type LogSyslogd4OverrideFilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
	Dns pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Free style filter string.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles LogSyslogd4OverrideFilterFreeStyleArrayInput
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringPtrInput
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringPtrInput
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringPtrInput
	// Enable/disable netscan discovery event logging.
	NetscanDiscovery pulumi.StringPtrInput
	// Enable/disable netscan vulnerability event logging.
	NetscanVulnerability pulumi.StringPtrInput
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringPtrInput
	// Enable/disable SSH logging. Valid values: `enable`, `disable`.
	Ssh pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringPtrInput
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringPtrInput
}

func (LogSyslogd4OverrideFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd4OverrideFilterArgs)(nil)).Elem()
}

type LogSyslogd4OverrideFilterInput interface {
	pulumi.Input

	ToLogSyslogd4OverrideFilterOutput() LogSyslogd4OverrideFilterOutput
	ToLogSyslogd4OverrideFilterOutputWithContext(ctx context.Context) LogSyslogd4OverrideFilterOutput
}

func (*LogSyslogd4OverrideFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd4OverrideFilter)(nil)).Elem()
}

func (i *LogSyslogd4OverrideFilter) ToLogSyslogd4OverrideFilterOutput() LogSyslogd4OverrideFilterOutput {
	return i.ToLogSyslogd4OverrideFilterOutputWithContext(context.Background())
}

func (i *LogSyslogd4OverrideFilter) ToLogSyslogd4OverrideFilterOutputWithContext(ctx context.Context) LogSyslogd4OverrideFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4OverrideFilterOutput)
}

// LogSyslogd4OverrideFilterArrayInput is an input type that accepts LogSyslogd4OverrideFilterArray and LogSyslogd4OverrideFilterArrayOutput values.
// You can construct a concrete instance of `LogSyslogd4OverrideFilterArrayInput` via:
//
//          LogSyslogd4OverrideFilterArray{ LogSyslogd4OverrideFilterArgs{...} }
type LogSyslogd4OverrideFilterArrayInput interface {
	pulumi.Input

	ToLogSyslogd4OverrideFilterArrayOutput() LogSyslogd4OverrideFilterArrayOutput
	ToLogSyslogd4OverrideFilterArrayOutputWithContext(context.Context) LogSyslogd4OverrideFilterArrayOutput
}

type LogSyslogd4OverrideFilterArray []LogSyslogd4OverrideFilterInput

func (LogSyslogd4OverrideFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd4OverrideFilter)(nil)).Elem()
}

func (i LogSyslogd4OverrideFilterArray) ToLogSyslogd4OverrideFilterArrayOutput() LogSyslogd4OverrideFilterArrayOutput {
	return i.ToLogSyslogd4OverrideFilterArrayOutputWithContext(context.Background())
}

func (i LogSyslogd4OverrideFilterArray) ToLogSyslogd4OverrideFilterArrayOutputWithContext(ctx context.Context) LogSyslogd4OverrideFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4OverrideFilterArrayOutput)
}

// LogSyslogd4OverrideFilterMapInput is an input type that accepts LogSyslogd4OverrideFilterMap and LogSyslogd4OverrideFilterMapOutput values.
// You can construct a concrete instance of `LogSyslogd4OverrideFilterMapInput` via:
//
//          LogSyslogd4OverrideFilterMap{ "key": LogSyslogd4OverrideFilterArgs{...} }
type LogSyslogd4OverrideFilterMapInput interface {
	pulumi.Input

	ToLogSyslogd4OverrideFilterMapOutput() LogSyslogd4OverrideFilterMapOutput
	ToLogSyslogd4OverrideFilterMapOutputWithContext(context.Context) LogSyslogd4OverrideFilterMapOutput
}

type LogSyslogd4OverrideFilterMap map[string]LogSyslogd4OverrideFilterInput

func (LogSyslogd4OverrideFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd4OverrideFilter)(nil)).Elem()
}

func (i LogSyslogd4OverrideFilterMap) ToLogSyslogd4OverrideFilterMapOutput() LogSyslogd4OverrideFilterMapOutput {
	return i.ToLogSyslogd4OverrideFilterMapOutputWithContext(context.Background())
}

func (i LogSyslogd4OverrideFilterMap) ToLogSyslogd4OverrideFilterMapOutputWithContext(ctx context.Context) LogSyslogd4OverrideFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4OverrideFilterMapOutput)
}

type LogSyslogd4OverrideFilterOutput struct{ *pulumi.OutputState }

func (LogSyslogd4OverrideFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd4OverrideFilter)(nil)).Elem()
}

func (o LogSyslogd4OverrideFilterOutput) ToLogSyslogd4OverrideFilterOutput() LogSyslogd4OverrideFilterOutput {
	return o
}

func (o LogSyslogd4OverrideFilterOutput) ToLogSyslogd4OverrideFilterOutputWithContext(ctx context.Context) LogSyslogd4OverrideFilterOutput {
	return o
}

type LogSyslogd4OverrideFilterArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogd4OverrideFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd4OverrideFilter)(nil)).Elem()
}

func (o LogSyslogd4OverrideFilterArrayOutput) ToLogSyslogd4OverrideFilterArrayOutput() LogSyslogd4OverrideFilterArrayOutput {
	return o
}

func (o LogSyslogd4OverrideFilterArrayOutput) ToLogSyslogd4OverrideFilterArrayOutputWithContext(ctx context.Context) LogSyslogd4OverrideFilterArrayOutput {
	return o
}

func (o LogSyslogd4OverrideFilterArrayOutput) Index(i pulumi.IntInput) LogSyslogd4OverrideFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogd4OverrideFilter {
		return vs[0].([]*LogSyslogd4OverrideFilter)[vs[1].(int)]
	}).(LogSyslogd4OverrideFilterOutput)
}

type LogSyslogd4OverrideFilterMapOutput struct{ *pulumi.OutputState }

func (LogSyslogd4OverrideFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd4OverrideFilter)(nil)).Elem()
}

func (o LogSyslogd4OverrideFilterMapOutput) ToLogSyslogd4OverrideFilterMapOutput() LogSyslogd4OverrideFilterMapOutput {
	return o
}

func (o LogSyslogd4OverrideFilterMapOutput) ToLogSyslogd4OverrideFilterMapOutputWithContext(ctx context.Context) LogSyslogd4OverrideFilterMapOutput {
	return o
}

func (o LogSyslogd4OverrideFilterMapOutput) MapIndex(k pulumi.StringInput) LogSyslogd4OverrideFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogd4OverrideFilter {
		return vs[0].(map[string]*LogSyslogd4OverrideFilter)[vs[1].(string)]
	}).(LogSyslogd4OverrideFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4OverrideFilterInput)(nil)).Elem(), &LogSyslogd4OverrideFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4OverrideFilterArrayInput)(nil)).Elem(), LogSyslogd4OverrideFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4OverrideFilterMapInput)(nil)).Elem(), LogSyslogd4OverrideFilterMap{})
	pulumi.RegisterOutputType(LogSyslogd4OverrideFilterOutput{})
	pulumi.RegisterOutputType(LogSyslogd4OverrideFilterArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogd4OverrideFilterMapOutput{})
}

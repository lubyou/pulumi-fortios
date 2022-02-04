// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
// 		_, err := fortios.NewLogSyslogd2OverrideFilter(ctx, "trname", &fortios.LogSyslogd2OverrideFilterArgs{
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
// LogSyslogd2 OverrideFilter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd2OverrideFilter:LogSyslogd2OverrideFilter labelname LogSyslogd2OverrideFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSyslogd2OverrideFilter struct {
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
	FreeStyles LogSyslogd2OverrideFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogSyslogd2OverrideFilter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd2OverrideFilter(ctx *pulumi.Context,
	name string, args *LogSyslogd2OverrideFilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogd2OverrideFilter, error) {
	if args == nil {
		args = &LogSyslogd2OverrideFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogd2OverrideFilter
	err := ctx.RegisterResource("fortios:index/logSyslogd2OverrideFilter:LogSyslogd2OverrideFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogd2OverrideFilter gets an existing LogSyslogd2OverrideFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogd2OverrideFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogd2OverrideFilterState, opts ...pulumi.ResourceOption) (*LogSyslogd2OverrideFilter, error) {
	var resource LogSyslogd2OverrideFilter
	err := ctx.ReadResource("fortios:index/logSyslogd2OverrideFilter:LogSyslogd2OverrideFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogd2OverrideFilter resources.
type logSyslogd2OverrideFilterState struct {
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
	FreeStyles []LogSyslogd2OverrideFilterFreeStyle `pulumi:"freeStyles"`
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

type LogSyslogd2OverrideFilterState struct {
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
	FreeStyles LogSyslogd2OverrideFilterFreeStyleArrayInput
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

func (LogSyslogd2OverrideFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd2OverrideFilterState)(nil)).Elem()
}

type logSyslogd2OverrideFilterArgs struct {
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
	FreeStyles []LogSyslogd2OverrideFilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogSyslogd2OverrideFilter resource.
type LogSyslogd2OverrideFilterArgs struct {
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
	FreeStyles LogSyslogd2OverrideFilterFreeStyleArrayInput
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

func (LogSyslogd2OverrideFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd2OverrideFilterArgs)(nil)).Elem()
}

type LogSyslogd2OverrideFilterInput interface {
	pulumi.Input

	ToLogSyslogd2OverrideFilterOutput() LogSyslogd2OverrideFilterOutput
	ToLogSyslogd2OverrideFilterOutputWithContext(ctx context.Context) LogSyslogd2OverrideFilterOutput
}

func (*LogSyslogd2OverrideFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd2OverrideFilter)(nil)).Elem()
}

func (i *LogSyslogd2OverrideFilter) ToLogSyslogd2OverrideFilterOutput() LogSyslogd2OverrideFilterOutput {
	return i.ToLogSyslogd2OverrideFilterOutputWithContext(context.Background())
}

func (i *LogSyslogd2OverrideFilter) ToLogSyslogd2OverrideFilterOutputWithContext(ctx context.Context) LogSyslogd2OverrideFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2OverrideFilterOutput)
}

// LogSyslogd2OverrideFilterArrayInput is an input type that accepts LogSyslogd2OverrideFilterArray and LogSyslogd2OverrideFilterArrayOutput values.
// You can construct a concrete instance of `LogSyslogd2OverrideFilterArrayInput` via:
//
//          LogSyslogd2OverrideFilterArray{ LogSyslogd2OverrideFilterArgs{...} }
type LogSyslogd2OverrideFilterArrayInput interface {
	pulumi.Input

	ToLogSyslogd2OverrideFilterArrayOutput() LogSyslogd2OverrideFilterArrayOutput
	ToLogSyslogd2OverrideFilterArrayOutputWithContext(context.Context) LogSyslogd2OverrideFilterArrayOutput
}

type LogSyslogd2OverrideFilterArray []LogSyslogd2OverrideFilterInput

func (LogSyslogd2OverrideFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd2OverrideFilter)(nil)).Elem()
}

func (i LogSyslogd2OverrideFilterArray) ToLogSyslogd2OverrideFilterArrayOutput() LogSyslogd2OverrideFilterArrayOutput {
	return i.ToLogSyslogd2OverrideFilterArrayOutputWithContext(context.Background())
}

func (i LogSyslogd2OverrideFilterArray) ToLogSyslogd2OverrideFilterArrayOutputWithContext(ctx context.Context) LogSyslogd2OverrideFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2OverrideFilterArrayOutput)
}

// LogSyslogd2OverrideFilterMapInput is an input type that accepts LogSyslogd2OverrideFilterMap and LogSyslogd2OverrideFilterMapOutput values.
// You can construct a concrete instance of `LogSyslogd2OverrideFilterMapInput` via:
//
//          LogSyslogd2OverrideFilterMap{ "key": LogSyslogd2OverrideFilterArgs{...} }
type LogSyslogd2OverrideFilterMapInput interface {
	pulumi.Input

	ToLogSyslogd2OverrideFilterMapOutput() LogSyslogd2OverrideFilterMapOutput
	ToLogSyslogd2OverrideFilterMapOutputWithContext(context.Context) LogSyslogd2OverrideFilterMapOutput
}

type LogSyslogd2OverrideFilterMap map[string]LogSyslogd2OverrideFilterInput

func (LogSyslogd2OverrideFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd2OverrideFilter)(nil)).Elem()
}

func (i LogSyslogd2OverrideFilterMap) ToLogSyslogd2OverrideFilterMapOutput() LogSyslogd2OverrideFilterMapOutput {
	return i.ToLogSyslogd2OverrideFilterMapOutputWithContext(context.Background())
}

func (i LogSyslogd2OverrideFilterMap) ToLogSyslogd2OverrideFilterMapOutputWithContext(ctx context.Context) LogSyslogd2OverrideFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2OverrideFilterMapOutput)
}

type LogSyslogd2OverrideFilterOutput struct{ *pulumi.OutputState }

func (LogSyslogd2OverrideFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd2OverrideFilter)(nil)).Elem()
}

func (o LogSyslogd2OverrideFilterOutput) ToLogSyslogd2OverrideFilterOutput() LogSyslogd2OverrideFilterOutput {
	return o
}

func (o LogSyslogd2OverrideFilterOutput) ToLogSyslogd2OverrideFilterOutputWithContext(ctx context.Context) LogSyslogd2OverrideFilterOutput {
	return o
}

type LogSyslogd2OverrideFilterArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogd2OverrideFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd2OverrideFilter)(nil)).Elem()
}

func (o LogSyslogd2OverrideFilterArrayOutput) ToLogSyslogd2OverrideFilterArrayOutput() LogSyslogd2OverrideFilterArrayOutput {
	return o
}

func (o LogSyslogd2OverrideFilterArrayOutput) ToLogSyslogd2OverrideFilterArrayOutputWithContext(ctx context.Context) LogSyslogd2OverrideFilterArrayOutput {
	return o
}

func (o LogSyslogd2OverrideFilterArrayOutput) Index(i pulumi.IntInput) LogSyslogd2OverrideFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogd2OverrideFilter {
		return vs[0].([]*LogSyslogd2OverrideFilter)[vs[1].(int)]
	}).(LogSyslogd2OverrideFilterOutput)
}

type LogSyslogd2OverrideFilterMapOutput struct{ *pulumi.OutputState }

func (LogSyslogd2OverrideFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd2OverrideFilter)(nil)).Elem()
}

func (o LogSyslogd2OverrideFilterMapOutput) ToLogSyslogd2OverrideFilterMapOutput() LogSyslogd2OverrideFilterMapOutput {
	return o
}

func (o LogSyslogd2OverrideFilterMapOutput) ToLogSyslogd2OverrideFilterMapOutputWithContext(ctx context.Context) LogSyslogd2OverrideFilterMapOutput {
	return o
}

func (o LogSyslogd2OverrideFilterMapOutput) MapIndex(k pulumi.StringInput) LogSyslogd2OverrideFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogd2OverrideFilter {
		return vs[0].(map[string]*LogSyslogd2OverrideFilter)[vs[1].(string)]
	}).(LogSyslogd2OverrideFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd2OverrideFilterInput)(nil)).Elem(), &LogSyslogd2OverrideFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd2OverrideFilterArrayInput)(nil)).Elem(), LogSyslogd2OverrideFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd2OverrideFilterMapInput)(nil)).Elem(), LogSyslogd2OverrideFilterMap{})
	pulumi.RegisterOutputType(LogSyslogd2OverrideFilterOutput{})
	pulumi.RegisterOutputType(LogSyslogd2OverrideFilterArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogd2OverrideFilterMapOutput{})
}

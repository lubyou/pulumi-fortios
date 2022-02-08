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
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewLogSyslogdOverrideFilter(ctx, "trname", &fortios.LogSyslogdOverrideFilterArgs{
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
// LogSyslogd OverrideFilter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logSyslogdOverrideFilter:LogSyslogdOverrideFilter labelname LogSyslogdOverrideFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSyslogdOverrideFilter struct {
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
	FreeStyles LogSyslogdOverrideFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogSyslogdOverrideFilter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogdOverrideFilter(ctx *pulumi.Context,
	name string, args *LogSyslogdOverrideFilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogdOverrideFilter, error) {
	if args == nil {
		args = &LogSyslogdOverrideFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogdOverrideFilter
	err := ctx.RegisterResource("fortios:index/logSyslogdOverrideFilter:LogSyslogdOverrideFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogdOverrideFilter gets an existing LogSyslogdOverrideFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogdOverrideFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogdOverrideFilterState, opts ...pulumi.ResourceOption) (*LogSyslogdOverrideFilter, error) {
	var resource LogSyslogdOverrideFilter
	err := ctx.ReadResource("fortios:index/logSyslogdOverrideFilter:LogSyslogdOverrideFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogdOverrideFilter resources.
type logSyslogdOverrideFilterState struct {
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
	FreeStyles []LogSyslogdOverrideFilterFreeStyle `pulumi:"freeStyles"`
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

type LogSyslogdOverrideFilterState struct {
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
	FreeStyles LogSyslogdOverrideFilterFreeStyleArrayInput
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

func (LogSyslogdOverrideFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogdOverrideFilterState)(nil)).Elem()
}

type logSyslogdOverrideFilterArgs struct {
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
	FreeStyles []LogSyslogdOverrideFilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogSyslogdOverrideFilter resource.
type LogSyslogdOverrideFilterArgs struct {
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
	FreeStyles LogSyslogdOverrideFilterFreeStyleArrayInput
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

func (LogSyslogdOverrideFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogdOverrideFilterArgs)(nil)).Elem()
}

type LogSyslogdOverrideFilterInput interface {
	pulumi.Input

	ToLogSyslogdOverrideFilterOutput() LogSyslogdOverrideFilterOutput
	ToLogSyslogdOverrideFilterOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterOutput
}

func (*LogSyslogdOverrideFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogdOverrideFilter)(nil)).Elem()
}

func (i *LogSyslogdOverrideFilter) ToLogSyslogdOverrideFilterOutput() LogSyslogdOverrideFilterOutput {
	return i.ToLogSyslogdOverrideFilterOutputWithContext(context.Background())
}

func (i *LogSyslogdOverrideFilter) ToLogSyslogdOverrideFilterOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdOverrideFilterOutput)
}

// LogSyslogdOverrideFilterArrayInput is an input type that accepts LogSyslogdOverrideFilterArray and LogSyslogdOverrideFilterArrayOutput values.
// You can construct a concrete instance of `LogSyslogdOverrideFilterArrayInput` via:
//
//          LogSyslogdOverrideFilterArray{ LogSyslogdOverrideFilterArgs{...} }
type LogSyslogdOverrideFilterArrayInput interface {
	pulumi.Input

	ToLogSyslogdOverrideFilterArrayOutput() LogSyslogdOverrideFilterArrayOutput
	ToLogSyslogdOverrideFilterArrayOutputWithContext(context.Context) LogSyslogdOverrideFilterArrayOutput
}

type LogSyslogdOverrideFilterArray []LogSyslogdOverrideFilterInput

func (LogSyslogdOverrideFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogdOverrideFilter)(nil)).Elem()
}

func (i LogSyslogdOverrideFilterArray) ToLogSyslogdOverrideFilterArrayOutput() LogSyslogdOverrideFilterArrayOutput {
	return i.ToLogSyslogdOverrideFilterArrayOutputWithContext(context.Background())
}

func (i LogSyslogdOverrideFilterArray) ToLogSyslogdOverrideFilterArrayOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdOverrideFilterArrayOutput)
}

// LogSyslogdOverrideFilterMapInput is an input type that accepts LogSyslogdOverrideFilterMap and LogSyslogdOverrideFilterMapOutput values.
// You can construct a concrete instance of `LogSyslogdOverrideFilterMapInput` via:
//
//          LogSyslogdOverrideFilterMap{ "key": LogSyslogdOverrideFilterArgs{...} }
type LogSyslogdOverrideFilterMapInput interface {
	pulumi.Input

	ToLogSyslogdOverrideFilterMapOutput() LogSyslogdOverrideFilterMapOutput
	ToLogSyslogdOverrideFilterMapOutputWithContext(context.Context) LogSyslogdOverrideFilterMapOutput
}

type LogSyslogdOverrideFilterMap map[string]LogSyslogdOverrideFilterInput

func (LogSyslogdOverrideFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogdOverrideFilter)(nil)).Elem()
}

func (i LogSyslogdOverrideFilterMap) ToLogSyslogdOverrideFilterMapOutput() LogSyslogdOverrideFilterMapOutput {
	return i.ToLogSyslogdOverrideFilterMapOutputWithContext(context.Background())
}

func (i LogSyslogdOverrideFilterMap) ToLogSyslogdOverrideFilterMapOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdOverrideFilterMapOutput)
}

type LogSyslogdOverrideFilterOutput struct{ *pulumi.OutputState }

func (LogSyslogdOverrideFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogdOverrideFilter)(nil)).Elem()
}

func (o LogSyslogdOverrideFilterOutput) ToLogSyslogdOverrideFilterOutput() LogSyslogdOverrideFilterOutput {
	return o
}

func (o LogSyslogdOverrideFilterOutput) ToLogSyslogdOverrideFilterOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterOutput {
	return o
}

type LogSyslogdOverrideFilterArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogdOverrideFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogdOverrideFilter)(nil)).Elem()
}

func (o LogSyslogdOverrideFilterArrayOutput) ToLogSyslogdOverrideFilterArrayOutput() LogSyslogdOverrideFilterArrayOutput {
	return o
}

func (o LogSyslogdOverrideFilterArrayOutput) ToLogSyslogdOverrideFilterArrayOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterArrayOutput {
	return o
}

func (o LogSyslogdOverrideFilterArrayOutput) Index(i pulumi.IntInput) LogSyslogdOverrideFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogdOverrideFilter {
		return vs[0].([]*LogSyslogdOverrideFilter)[vs[1].(int)]
	}).(LogSyslogdOverrideFilterOutput)
}

type LogSyslogdOverrideFilterMapOutput struct{ *pulumi.OutputState }

func (LogSyslogdOverrideFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogdOverrideFilter)(nil)).Elem()
}

func (o LogSyslogdOverrideFilterMapOutput) ToLogSyslogdOverrideFilterMapOutput() LogSyslogdOverrideFilterMapOutput {
	return o
}

func (o LogSyslogdOverrideFilterMapOutput) ToLogSyslogdOverrideFilterMapOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterMapOutput {
	return o
}

func (o LogSyslogdOverrideFilterMapOutput) MapIndex(k pulumi.StringInput) LogSyslogdOverrideFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogdOverrideFilter {
		return vs[0].(map[string]*LogSyslogdOverrideFilter)[vs[1].(string)]
	}).(LogSyslogdOverrideFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdOverrideFilterInput)(nil)).Elem(), &LogSyslogdOverrideFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdOverrideFilterArrayInput)(nil)).Elem(), LogSyslogdOverrideFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdOverrideFilterMapInput)(nil)).Elem(), LogSyslogdOverrideFilterMap{})
	pulumi.RegisterOutputType(LogSyslogdOverrideFilterOutput{})
	pulumi.RegisterOutputType(LogSyslogdOverrideFilterArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogdOverrideFilterMapOutput{})
}

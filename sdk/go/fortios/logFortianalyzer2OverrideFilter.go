// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Override filters for FortiAnalyzer.
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
// 		_, err := fortios.NewLogFortianalyzer2OverrideFilter(ctx, "trname", &fortios.LogFortianalyzer2OverrideFilterArgs{
// 			Anomaly:          pulumi.String("enable"),
// 			DlpArchive:       pulumi.String("enable"),
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
// LogFortianalyzer2 OverrideFilter can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzer2OverrideFilter:LogFortianalyzer2OverrideFilter labelname LogFortianalyzer2OverrideFilter
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzer2OverrideFilter:LogFortianalyzer2OverrideFilter labelname LogFortianalyzer2OverrideFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortianalyzer2OverrideFilter struct {
	pulumi.CustomResourceState

	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringOutput `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringOutput `pulumi:"dlpArchive"`
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
	FreeStyles LogFortianalyzer2OverrideFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
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

// NewLogFortianalyzer2OverrideFilter registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzer2OverrideFilter(ctx *pulumi.Context,
	name string, args *LogFortianalyzer2OverrideFilterArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzer2OverrideFilter, error) {
	if args == nil {
		args = &LogFortianalyzer2OverrideFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogFortianalyzer2OverrideFilter
	err := ctx.RegisterResource("fortios:index/logFortianalyzer2OverrideFilter:LogFortianalyzer2OverrideFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzer2OverrideFilter gets an existing LogFortianalyzer2OverrideFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzer2OverrideFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzer2OverrideFilterState, opts ...pulumi.ResourceOption) (*LogFortianalyzer2OverrideFilter, error) {
	var resource LogFortianalyzer2OverrideFilter
	err := ctx.ReadResource("fortios:index/logFortianalyzer2OverrideFilter:LogFortianalyzer2OverrideFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzer2OverrideFilter resources.
type logFortianalyzer2OverrideFilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive *string `pulumi:"dlpArchive"`
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
	FreeStyles []LogFortianalyzer2OverrideFilterFreeStyle `pulumi:"freeStyles"`
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
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
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

type LogFortianalyzer2OverrideFilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringPtrInput
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
	FreeStyles LogFortianalyzer2OverrideFilterFreeStyleArrayInput
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
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
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

func (LogFortianalyzer2OverrideFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer2OverrideFilterState)(nil)).Elem()
}

type logFortianalyzer2OverrideFilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive *string `pulumi:"dlpArchive"`
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
	FreeStyles []LogFortianalyzer2OverrideFilterFreeStyle `pulumi:"freeStyles"`
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
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
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

// The set of arguments for constructing a LogFortianalyzer2OverrideFilter resource.
type LogFortianalyzer2OverrideFilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringPtrInput
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
	FreeStyles LogFortianalyzer2OverrideFilterFreeStyleArrayInput
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
	// Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
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

func (LogFortianalyzer2OverrideFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer2OverrideFilterArgs)(nil)).Elem()
}

type LogFortianalyzer2OverrideFilterInput interface {
	pulumi.Input

	ToLogFortianalyzer2OverrideFilterOutput() LogFortianalyzer2OverrideFilterOutput
	ToLogFortianalyzer2OverrideFilterOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideFilterOutput
}

func (*LogFortianalyzer2OverrideFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer2OverrideFilter)(nil)).Elem()
}

func (i *LogFortianalyzer2OverrideFilter) ToLogFortianalyzer2OverrideFilterOutput() LogFortianalyzer2OverrideFilterOutput {
	return i.ToLogFortianalyzer2OverrideFilterOutputWithContext(context.Background())
}

func (i *LogFortianalyzer2OverrideFilter) ToLogFortianalyzer2OverrideFilterOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer2OverrideFilterOutput)
}

// LogFortianalyzer2OverrideFilterArrayInput is an input type that accepts LogFortianalyzer2OverrideFilterArray and LogFortianalyzer2OverrideFilterArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzer2OverrideFilterArrayInput` via:
//
//          LogFortianalyzer2OverrideFilterArray{ LogFortianalyzer2OverrideFilterArgs{...} }
type LogFortianalyzer2OverrideFilterArrayInput interface {
	pulumi.Input

	ToLogFortianalyzer2OverrideFilterArrayOutput() LogFortianalyzer2OverrideFilterArrayOutput
	ToLogFortianalyzer2OverrideFilterArrayOutputWithContext(context.Context) LogFortianalyzer2OverrideFilterArrayOutput
}

type LogFortianalyzer2OverrideFilterArray []LogFortianalyzer2OverrideFilterInput

func (LogFortianalyzer2OverrideFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzer2OverrideFilter)(nil)).Elem()
}

func (i LogFortianalyzer2OverrideFilterArray) ToLogFortianalyzer2OverrideFilterArrayOutput() LogFortianalyzer2OverrideFilterArrayOutput {
	return i.ToLogFortianalyzer2OverrideFilterArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzer2OverrideFilterArray) ToLogFortianalyzer2OverrideFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer2OverrideFilterArrayOutput)
}

// LogFortianalyzer2OverrideFilterMapInput is an input type that accepts LogFortianalyzer2OverrideFilterMap and LogFortianalyzer2OverrideFilterMapOutput values.
// You can construct a concrete instance of `LogFortianalyzer2OverrideFilterMapInput` via:
//
//          LogFortianalyzer2OverrideFilterMap{ "key": LogFortianalyzer2OverrideFilterArgs{...} }
type LogFortianalyzer2OverrideFilterMapInput interface {
	pulumi.Input

	ToLogFortianalyzer2OverrideFilterMapOutput() LogFortianalyzer2OverrideFilterMapOutput
	ToLogFortianalyzer2OverrideFilterMapOutputWithContext(context.Context) LogFortianalyzer2OverrideFilterMapOutput
}

type LogFortianalyzer2OverrideFilterMap map[string]LogFortianalyzer2OverrideFilterInput

func (LogFortianalyzer2OverrideFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzer2OverrideFilter)(nil)).Elem()
}

func (i LogFortianalyzer2OverrideFilterMap) ToLogFortianalyzer2OverrideFilterMapOutput() LogFortianalyzer2OverrideFilterMapOutput {
	return i.ToLogFortianalyzer2OverrideFilterMapOutputWithContext(context.Background())
}

func (i LogFortianalyzer2OverrideFilterMap) ToLogFortianalyzer2OverrideFilterMapOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer2OverrideFilterMapOutput)
}

type LogFortianalyzer2OverrideFilterOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer2OverrideFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer2OverrideFilter)(nil)).Elem()
}

func (o LogFortianalyzer2OverrideFilterOutput) ToLogFortianalyzer2OverrideFilterOutput() LogFortianalyzer2OverrideFilterOutput {
	return o
}

func (o LogFortianalyzer2OverrideFilterOutput) ToLogFortianalyzer2OverrideFilterOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideFilterOutput {
	return o
}

type LogFortianalyzer2OverrideFilterArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer2OverrideFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzer2OverrideFilter)(nil)).Elem()
}

func (o LogFortianalyzer2OverrideFilterArrayOutput) ToLogFortianalyzer2OverrideFilterArrayOutput() LogFortianalyzer2OverrideFilterArrayOutput {
	return o
}

func (o LogFortianalyzer2OverrideFilterArrayOutput) ToLogFortianalyzer2OverrideFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideFilterArrayOutput {
	return o
}

func (o LogFortianalyzer2OverrideFilterArrayOutput) Index(i pulumi.IntInput) LogFortianalyzer2OverrideFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortianalyzer2OverrideFilter {
		return vs[0].([]*LogFortianalyzer2OverrideFilter)[vs[1].(int)]
	}).(LogFortianalyzer2OverrideFilterOutput)
}

type LogFortianalyzer2OverrideFilterMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer2OverrideFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzer2OverrideFilter)(nil)).Elem()
}

func (o LogFortianalyzer2OverrideFilterMapOutput) ToLogFortianalyzer2OverrideFilterMapOutput() LogFortianalyzer2OverrideFilterMapOutput {
	return o
}

func (o LogFortianalyzer2OverrideFilterMapOutput) ToLogFortianalyzer2OverrideFilterMapOutputWithContext(ctx context.Context) LogFortianalyzer2OverrideFilterMapOutput {
	return o
}

func (o LogFortianalyzer2OverrideFilterMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzer2OverrideFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortianalyzer2OverrideFilter {
		return vs[0].(map[string]*LogFortianalyzer2OverrideFilter)[vs[1].(string)]
	}).(LogFortianalyzer2OverrideFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer2OverrideFilterInput)(nil)).Elem(), &LogFortianalyzer2OverrideFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer2OverrideFilterArrayInput)(nil)).Elem(), LogFortianalyzer2OverrideFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer2OverrideFilterMapInput)(nil)).Elem(), LogFortianalyzer2OverrideFilterMap{})
	pulumi.RegisterOutputType(LogFortianalyzer2OverrideFilterOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer2OverrideFilterArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer2OverrideFilterMapOutput{})
}

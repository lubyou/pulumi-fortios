// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
// 		_, err := fortios.NewLogFortianalyzerOverrideFilter(ctx, "trname", &fortios.LogFortianalyzerOverrideFilterArgs{
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
// LogFortianalyzer OverrideFilter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzerOverrideFilter:LogFortianalyzerOverrideFilter labelname LogFortianalyzerOverrideFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortianalyzerOverrideFilter struct {
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
	FreeStyles LogFortianalyzerOverrideFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogFortianalyzerOverrideFilter registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzerOverrideFilter(ctx *pulumi.Context,
	name string, args *LogFortianalyzerOverrideFilterArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzerOverrideFilter, error) {
	if args == nil {
		args = &LogFortianalyzerOverrideFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogFortianalyzerOverrideFilter
	err := ctx.RegisterResource("fortios:index/logFortianalyzerOverrideFilter:LogFortianalyzerOverrideFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzerOverrideFilter gets an existing LogFortianalyzerOverrideFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzerOverrideFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzerOverrideFilterState, opts ...pulumi.ResourceOption) (*LogFortianalyzerOverrideFilter, error) {
	var resource LogFortianalyzerOverrideFilter
	err := ctx.ReadResource("fortios:index/logFortianalyzerOverrideFilter:LogFortianalyzerOverrideFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzerOverrideFilter resources.
type logFortianalyzerOverrideFilterState struct {
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
	FreeStyles []LogFortianalyzerOverrideFilterFreeStyle `pulumi:"freeStyles"`
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

type LogFortianalyzerOverrideFilterState struct {
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
	FreeStyles LogFortianalyzerOverrideFilterFreeStyleArrayInput
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

func (LogFortianalyzerOverrideFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzerOverrideFilterState)(nil)).Elem()
}

type logFortianalyzerOverrideFilterArgs struct {
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
	FreeStyles []LogFortianalyzerOverrideFilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogFortianalyzerOverrideFilter resource.
type LogFortianalyzerOverrideFilterArgs struct {
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
	FreeStyles LogFortianalyzerOverrideFilterFreeStyleArrayInput
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

func (LogFortianalyzerOverrideFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzerOverrideFilterArgs)(nil)).Elem()
}

type LogFortianalyzerOverrideFilterInput interface {
	pulumi.Input

	ToLogFortianalyzerOverrideFilterOutput() LogFortianalyzerOverrideFilterOutput
	ToLogFortianalyzerOverrideFilterOutputWithContext(ctx context.Context) LogFortianalyzerOverrideFilterOutput
}

func (*LogFortianalyzerOverrideFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzerOverrideFilter)(nil)).Elem()
}

func (i *LogFortianalyzerOverrideFilter) ToLogFortianalyzerOverrideFilterOutput() LogFortianalyzerOverrideFilterOutput {
	return i.ToLogFortianalyzerOverrideFilterOutputWithContext(context.Background())
}

func (i *LogFortianalyzerOverrideFilter) ToLogFortianalyzerOverrideFilterOutputWithContext(ctx context.Context) LogFortianalyzerOverrideFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerOverrideFilterOutput)
}

// LogFortianalyzerOverrideFilterArrayInput is an input type that accepts LogFortianalyzerOverrideFilterArray and LogFortianalyzerOverrideFilterArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzerOverrideFilterArrayInput` via:
//
//          LogFortianalyzerOverrideFilterArray{ LogFortianalyzerOverrideFilterArgs{...} }
type LogFortianalyzerOverrideFilterArrayInput interface {
	pulumi.Input

	ToLogFortianalyzerOverrideFilterArrayOutput() LogFortianalyzerOverrideFilterArrayOutput
	ToLogFortianalyzerOverrideFilterArrayOutputWithContext(context.Context) LogFortianalyzerOverrideFilterArrayOutput
}

type LogFortianalyzerOverrideFilterArray []LogFortianalyzerOverrideFilterInput

func (LogFortianalyzerOverrideFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzerOverrideFilter)(nil)).Elem()
}

func (i LogFortianalyzerOverrideFilterArray) ToLogFortianalyzerOverrideFilterArrayOutput() LogFortianalyzerOverrideFilterArrayOutput {
	return i.ToLogFortianalyzerOverrideFilterArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzerOverrideFilterArray) ToLogFortianalyzerOverrideFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzerOverrideFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerOverrideFilterArrayOutput)
}

// LogFortianalyzerOverrideFilterMapInput is an input type that accepts LogFortianalyzerOverrideFilterMap and LogFortianalyzerOverrideFilterMapOutput values.
// You can construct a concrete instance of `LogFortianalyzerOverrideFilterMapInput` via:
//
//          LogFortianalyzerOverrideFilterMap{ "key": LogFortianalyzerOverrideFilterArgs{...} }
type LogFortianalyzerOverrideFilterMapInput interface {
	pulumi.Input

	ToLogFortianalyzerOverrideFilterMapOutput() LogFortianalyzerOverrideFilterMapOutput
	ToLogFortianalyzerOverrideFilterMapOutputWithContext(context.Context) LogFortianalyzerOverrideFilterMapOutput
}

type LogFortianalyzerOverrideFilterMap map[string]LogFortianalyzerOverrideFilterInput

func (LogFortianalyzerOverrideFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzerOverrideFilter)(nil)).Elem()
}

func (i LogFortianalyzerOverrideFilterMap) ToLogFortianalyzerOverrideFilterMapOutput() LogFortianalyzerOverrideFilterMapOutput {
	return i.ToLogFortianalyzerOverrideFilterMapOutputWithContext(context.Background())
}

func (i LogFortianalyzerOverrideFilterMap) ToLogFortianalyzerOverrideFilterMapOutputWithContext(ctx context.Context) LogFortianalyzerOverrideFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerOverrideFilterMapOutput)
}

type LogFortianalyzerOverrideFilterOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerOverrideFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzerOverrideFilter)(nil)).Elem()
}

func (o LogFortianalyzerOverrideFilterOutput) ToLogFortianalyzerOverrideFilterOutput() LogFortianalyzerOverrideFilterOutput {
	return o
}

func (o LogFortianalyzerOverrideFilterOutput) ToLogFortianalyzerOverrideFilterOutputWithContext(ctx context.Context) LogFortianalyzerOverrideFilterOutput {
	return o
}

type LogFortianalyzerOverrideFilterArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerOverrideFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzerOverrideFilter)(nil)).Elem()
}

func (o LogFortianalyzerOverrideFilterArrayOutput) ToLogFortianalyzerOverrideFilterArrayOutput() LogFortianalyzerOverrideFilterArrayOutput {
	return o
}

func (o LogFortianalyzerOverrideFilterArrayOutput) ToLogFortianalyzerOverrideFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzerOverrideFilterArrayOutput {
	return o
}

func (o LogFortianalyzerOverrideFilterArrayOutput) Index(i pulumi.IntInput) LogFortianalyzerOverrideFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortianalyzerOverrideFilter {
		return vs[0].([]*LogFortianalyzerOverrideFilter)[vs[1].(int)]
	}).(LogFortianalyzerOverrideFilterOutput)
}

type LogFortianalyzerOverrideFilterMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerOverrideFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzerOverrideFilter)(nil)).Elem()
}

func (o LogFortianalyzerOverrideFilterMapOutput) ToLogFortianalyzerOverrideFilterMapOutput() LogFortianalyzerOverrideFilterMapOutput {
	return o
}

func (o LogFortianalyzerOverrideFilterMapOutput) ToLogFortianalyzerOverrideFilterMapOutputWithContext(ctx context.Context) LogFortianalyzerOverrideFilterMapOutput {
	return o
}

func (o LogFortianalyzerOverrideFilterMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzerOverrideFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortianalyzerOverrideFilter {
		return vs[0].(map[string]*LogFortianalyzerOverrideFilter)[vs[1].(string)]
	}).(LogFortianalyzerOverrideFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerOverrideFilterInput)(nil)).Elem(), &LogFortianalyzerOverrideFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerOverrideFilterArrayInput)(nil)).Elem(), LogFortianalyzerOverrideFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerOverrideFilterMapInput)(nil)).Elem(), LogFortianalyzerOverrideFilterMap{})
	pulumi.RegisterOutputType(LogFortianalyzerOverrideFilterOutput{})
	pulumi.RegisterOutputType(LogFortianalyzerOverrideFilterArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzerOverrideFilterMapOutput{})
}

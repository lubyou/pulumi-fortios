// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Filters for FortiAnalyzer.
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
// 		_, err := fortios.NewLogFortianalyzerFilter(ctx, "trname", &fortios.LogFortianalyzerFilterArgs{
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
// LogFortianalyzer Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzerFilter:LogFortianalyzerFilter labelname LogFortianalyzerFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortianalyzerFilter struct {
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
	FreeStyles LogFortianalyzerFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogFortianalyzerFilter registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzerFilter(ctx *pulumi.Context,
	name string, args *LogFortianalyzerFilterArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzerFilter, error) {
	if args == nil {
		args = &LogFortianalyzerFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogFortianalyzerFilter
	err := ctx.RegisterResource("fortios:index/logFortianalyzerFilter:LogFortianalyzerFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzerFilter gets an existing LogFortianalyzerFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzerFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzerFilterState, opts ...pulumi.ResourceOption) (*LogFortianalyzerFilter, error) {
	var resource LogFortianalyzerFilter
	err := ctx.ReadResource("fortios:index/logFortianalyzerFilter:LogFortianalyzerFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzerFilter resources.
type logFortianalyzerFilterState struct {
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
	FreeStyles []LogFortianalyzerFilterFreeStyle `pulumi:"freeStyles"`
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

type LogFortianalyzerFilterState struct {
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
	FreeStyles LogFortianalyzerFilterFreeStyleArrayInput
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

func (LogFortianalyzerFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzerFilterState)(nil)).Elem()
}

type logFortianalyzerFilterArgs struct {
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
	FreeStyles []LogFortianalyzerFilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogFortianalyzerFilter resource.
type LogFortianalyzerFilterArgs struct {
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
	FreeStyles LogFortianalyzerFilterFreeStyleArrayInput
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

func (LogFortianalyzerFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzerFilterArgs)(nil)).Elem()
}

type LogFortianalyzerFilterInput interface {
	pulumi.Input

	ToLogFortianalyzerFilterOutput() LogFortianalyzerFilterOutput
	ToLogFortianalyzerFilterOutputWithContext(ctx context.Context) LogFortianalyzerFilterOutput
}

func (*LogFortianalyzerFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzerFilter)(nil)).Elem()
}

func (i *LogFortianalyzerFilter) ToLogFortianalyzerFilterOutput() LogFortianalyzerFilterOutput {
	return i.ToLogFortianalyzerFilterOutputWithContext(context.Background())
}

func (i *LogFortianalyzerFilter) ToLogFortianalyzerFilterOutputWithContext(ctx context.Context) LogFortianalyzerFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerFilterOutput)
}

// LogFortianalyzerFilterArrayInput is an input type that accepts LogFortianalyzerFilterArray and LogFortianalyzerFilterArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzerFilterArrayInput` via:
//
//          LogFortianalyzerFilterArray{ LogFortianalyzerFilterArgs{...} }
type LogFortianalyzerFilterArrayInput interface {
	pulumi.Input

	ToLogFortianalyzerFilterArrayOutput() LogFortianalyzerFilterArrayOutput
	ToLogFortianalyzerFilterArrayOutputWithContext(context.Context) LogFortianalyzerFilterArrayOutput
}

type LogFortianalyzerFilterArray []LogFortianalyzerFilterInput

func (LogFortianalyzerFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzerFilter)(nil)).Elem()
}

func (i LogFortianalyzerFilterArray) ToLogFortianalyzerFilterArrayOutput() LogFortianalyzerFilterArrayOutput {
	return i.ToLogFortianalyzerFilterArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzerFilterArray) ToLogFortianalyzerFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzerFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerFilterArrayOutput)
}

// LogFortianalyzerFilterMapInput is an input type that accepts LogFortianalyzerFilterMap and LogFortianalyzerFilterMapOutput values.
// You can construct a concrete instance of `LogFortianalyzerFilterMapInput` via:
//
//          LogFortianalyzerFilterMap{ "key": LogFortianalyzerFilterArgs{...} }
type LogFortianalyzerFilterMapInput interface {
	pulumi.Input

	ToLogFortianalyzerFilterMapOutput() LogFortianalyzerFilterMapOutput
	ToLogFortianalyzerFilterMapOutputWithContext(context.Context) LogFortianalyzerFilterMapOutput
}

type LogFortianalyzerFilterMap map[string]LogFortianalyzerFilterInput

func (LogFortianalyzerFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzerFilter)(nil)).Elem()
}

func (i LogFortianalyzerFilterMap) ToLogFortianalyzerFilterMapOutput() LogFortianalyzerFilterMapOutput {
	return i.ToLogFortianalyzerFilterMapOutputWithContext(context.Background())
}

func (i LogFortianalyzerFilterMap) ToLogFortianalyzerFilterMapOutputWithContext(ctx context.Context) LogFortianalyzerFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerFilterMapOutput)
}

type LogFortianalyzerFilterOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzerFilter)(nil)).Elem()
}

func (o LogFortianalyzerFilterOutput) ToLogFortianalyzerFilterOutput() LogFortianalyzerFilterOutput {
	return o
}

func (o LogFortianalyzerFilterOutput) ToLogFortianalyzerFilterOutputWithContext(ctx context.Context) LogFortianalyzerFilterOutput {
	return o
}

type LogFortianalyzerFilterArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzerFilter)(nil)).Elem()
}

func (o LogFortianalyzerFilterArrayOutput) ToLogFortianalyzerFilterArrayOutput() LogFortianalyzerFilterArrayOutput {
	return o
}

func (o LogFortianalyzerFilterArrayOutput) ToLogFortianalyzerFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzerFilterArrayOutput {
	return o
}

func (o LogFortianalyzerFilterArrayOutput) Index(i pulumi.IntInput) LogFortianalyzerFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortianalyzerFilter {
		return vs[0].([]*LogFortianalyzerFilter)[vs[1].(int)]
	}).(LogFortianalyzerFilterOutput)
}

type LogFortianalyzerFilterMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzerFilter)(nil)).Elem()
}

func (o LogFortianalyzerFilterMapOutput) ToLogFortianalyzerFilterMapOutput() LogFortianalyzerFilterMapOutput {
	return o
}

func (o LogFortianalyzerFilterMapOutput) ToLogFortianalyzerFilterMapOutputWithContext(ctx context.Context) LogFortianalyzerFilterMapOutput {
	return o
}

func (o LogFortianalyzerFilterMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzerFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortianalyzerFilter {
		return vs[0].(map[string]*LogFortianalyzerFilter)[vs[1].(string)]
	}).(LogFortianalyzerFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerFilterInput)(nil)).Elem(), &LogFortianalyzerFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerFilterArrayInput)(nil)).Elem(), LogFortianalyzerFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerFilterMapInput)(nil)).Elem(), LogFortianalyzerFilterMap{})
	pulumi.RegisterOutputType(LogFortianalyzerFilterOutput{})
	pulumi.RegisterOutputType(LogFortianalyzerFilterArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzerFilterMapOutput{})
}

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
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewLogFortianalyzer3OverrideFilter(ctx, "trname", &fortios.LogFortianalyzer3OverrideFilterArgs{
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
// LogFortianalyzer3 OverrideFilter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzer3OverrideFilter:LogFortianalyzer3OverrideFilter labelname LogFortianalyzer3OverrideFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortianalyzer3OverrideFilter struct {
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
	FreeStyles LogFortianalyzer3OverrideFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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
}

// NewLogFortianalyzer3OverrideFilter registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzer3OverrideFilter(ctx *pulumi.Context,
	name string, args *LogFortianalyzer3OverrideFilterArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzer3OverrideFilter, error) {
	if args == nil {
		args = &LogFortianalyzer3OverrideFilterArgs{}
	}

	var resource LogFortianalyzer3OverrideFilter
	err := ctx.RegisterResource("fortios:index/logFortianalyzer3OverrideFilter:LogFortianalyzer3OverrideFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzer3OverrideFilter gets an existing LogFortianalyzer3OverrideFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzer3OverrideFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzer3OverrideFilterState, opts ...pulumi.ResourceOption) (*LogFortianalyzer3OverrideFilter, error) {
	var resource LogFortianalyzer3OverrideFilter
	err := ctx.ReadResource("fortios:index/logFortianalyzer3OverrideFilter:LogFortianalyzer3OverrideFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzer3OverrideFilter resources.
type logFortianalyzer3OverrideFilterState struct {
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
	FreeStyles []LogFortianalyzer3OverrideFilterFreeStyle `pulumi:"freeStyles"`
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
}

type LogFortianalyzer3OverrideFilterState struct {
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
	FreeStyles LogFortianalyzer3OverrideFilterFreeStyleArrayInput
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
}

func (LogFortianalyzer3OverrideFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3OverrideFilterState)(nil)).Elem()
}

type logFortianalyzer3OverrideFilterArgs struct {
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
	FreeStyles []LogFortianalyzer3OverrideFilterFreeStyle `pulumi:"freeStyles"`
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
}

// The set of arguments for constructing a LogFortianalyzer3OverrideFilter resource.
type LogFortianalyzer3OverrideFilterArgs struct {
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
	FreeStyles LogFortianalyzer3OverrideFilterFreeStyleArrayInput
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
}

func (LogFortianalyzer3OverrideFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3OverrideFilterArgs)(nil)).Elem()
}

type LogFortianalyzer3OverrideFilterInput interface {
	pulumi.Input

	ToLogFortianalyzer3OverrideFilterOutput() LogFortianalyzer3OverrideFilterOutput
	ToLogFortianalyzer3OverrideFilterOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterOutput
}

func (*LogFortianalyzer3OverrideFilter) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFortianalyzer3OverrideFilter)(nil))
}

func (i *LogFortianalyzer3OverrideFilter) ToLogFortianalyzer3OverrideFilterOutput() LogFortianalyzer3OverrideFilterOutput {
	return i.ToLogFortianalyzer3OverrideFilterOutputWithContext(context.Background())
}

func (i *LogFortianalyzer3OverrideFilter) ToLogFortianalyzer3OverrideFilterOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3OverrideFilterOutput)
}

func (i *LogFortianalyzer3OverrideFilter) ToLogFortianalyzer3OverrideFilterPtrOutput() LogFortianalyzer3OverrideFilterPtrOutput {
	return i.ToLogFortianalyzer3OverrideFilterPtrOutputWithContext(context.Background())
}

func (i *LogFortianalyzer3OverrideFilter) ToLogFortianalyzer3OverrideFilterPtrOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3OverrideFilterPtrOutput)
}

type LogFortianalyzer3OverrideFilterPtrInput interface {
	pulumi.Input

	ToLogFortianalyzer3OverrideFilterPtrOutput() LogFortianalyzer3OverrideFilterPtrOutput
	ToLogFortianalyzer3OverrideFilterPtrOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterPtrOutput
}

type logFortianalyzer3OverrideFilterPtrType LogFortianalyzer3OverrideFilterArgs

func (*logFortianalyzer3OverrideFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer3OverrideFilter)(nil))
}

func (i *logFortianalyzer3OverrideFilterPtrType) ToLogFortianalyzer3OverrideFilterPtrOutput() LogFortianalyzer3OverrideFilterPtrOutput {
	return i.ToLogFortianalyzer3OverrideFilterPtrOutputWithContext(context.Background())
}

func (i *logFortianalyzer3OverrideFilterPtrType) ToLogFortianalyzer3OverrideFilterPtrOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3OverrideFilterPtrOutput)
}

// LogFortianalyzer3OverrideFilterArrayInput is an input type that accepts LogFortianalyzer3OverrideFilterArray and LogFortianalyzer3OverrideFilterArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzer3OverrideFilterArrayInput` via:
//
//          LogFortianalyzer3OverrideFilterArray{ LogFortianalyzer3OverrideFilterArgs{...} }
type LogFortianalyzer3OverrideFilterArrayInput interface {
	pulumi.Input

	ToLogFortianalyzer3OverrideFilterArrayOutput() LogFortianalyzer3OverrideFilterArrayOutput
	ToLogFortianalyzer3OverrideFilterArrayOutputWithContext(context.Context) LogFortianalyzer3OverrideFilterArrayOutput
}

type LogFortianalyzer3OverrideFilterArray []LogFortianalyzer3OverrideFilterInput

func (LogFortianalyzer3OverrideFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogFortianalyzer3OverrideFilter)(nil))
}

func (i LogFortianalyzer3OverrideFilterArray) ToLogFortianalyzer3OverrideFilterArrayOutput() LogFortianalyzer3OverrideFilterArrayOutput {
	return i.ToLogFortianalyzer3OverrideFilterArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzer3OverrideFilterArray) ToLogFortianalyzer3OverrideFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3OverrideFilterArrayOutput)
}

// LogFortianalyzer3OverrideFilterMapInput is an input type that accepts LogFortianalyzer3OverrideFilterMap and LogFortianalyzer3OverrideFilterMapOutput values.
// You can construct a concrete instance of `LogFortianalyzer3OverrideFilterMapInput` via:
//
//          LogFortianalyzer3OverrideFilterMap{ "key": LogFortianalyzer3OverrideFilterArgs{...} }
type LogFortianalyzer3OverrideFilterMapInput interface {
	pulumi.Input

	ToLogFortianalyzer3OverrideFilterMapOutput() LogFortianalyzer3OverrideFilterMapOutput
	ToLogFortianalyzer3OverrideFilterMapOutputWithContext(context.Context) LogFortianalyzer3OverrideFilterMapOutput
}

type LogFortianalyzer3OverrideFilterMap map[string]LogFortianalyzer3OverrideFilterInput

func (LogFortianalyzer3OverrideFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogFortianalyzer3OverrideFilter)(nil))
}

func (i LogFortianalyzer3OverrideFilterMap) ToLogFortianalyzer3OverrideFilterMapOutput() LogFortianalyzer3OverrideFilterMapOutput {
	return i.ToLogFortianalyzer3OverrideFilterMapOutputWithContext(context.Background())
}

func (i LogFortianalyzer3OverrideFilterMap) ToLogFortianalyzer3OverrideFilterMapOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3OverrideFilterMapOutput)
}

type LogFortianalyzer3OverrideFilterOutput struct {
	*pulumi.OutputState
}

func (LogFortianalyzer3OverrideFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFortianalyzer3OverrideFilter)(nil))
}

func (o LogFortianalyzer3OverrideFilterOutput) ToLogFortianalyzer3OverrideFilterOutput() LogFortianalyzer3OverrideFilterOutput {
	return o
}

func (o LogFortianalyzer3OverrideFilterOutput) ToLogFortianalyzer3OverrideFilterOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterOutput {
	return o
}

func (o LogFortianalyzer3OverrideFilterOutput) ToLogFortianalyzer3OverrideFilterPtrOutput() LogFortianalyzer3OverrideFilterPtrOutput {
	return o.ToLogFortianalyzer3OverrideFilterPtrOutputWithContext(context.Background())
}

func (o LogFortianalyzer3OverrideFilterOutput) ToLogFortianalyzer3OverrideFilterPtrOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterPtrOutput {
	return o.ApplyT(func(v LogFortianalyzer3OverrideFilter) *LogFortianalyzer3OverrideFilter {
		return &v
	}).(LogFortianalyzer3OverrideFilterPtrOutput)
}

type LogFortianalyzer3OverrideFilterPtrOutput struct {
	*pulumi.OutputState
}

func (LogFortianalyzer3OverrideFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer3OverrideFilter)(nil))
}

func (o LogFortianalyzer3OverrideFilterPtrOutput) ToLogFortianalyzer3OverrideFilterPtrOutput() LogFortianalyzer3OverrideFilterPtrOutput {
	return o
}

func (o LogFortianalyzer3OverrideFilterPtrOutput) ToLogFortianalyzer3OverrideFilterPtrOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterPtrOutput {
	return o
}

type LogFortianalyzer3OverrideFilterArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3OverrideFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogFortianalyzer3OverrideFilter)(nil))
}

func (o LogFortianalyzer3OverrideFilterArrayOutput) ToLogFortianalyzer3OverrideFilterArrayOutput() LogFortianalyzer3OverrideFilterArrayOutput {
	return o
}

func (o LogFortianalyzer3OverrideFilterArrayOutput) ToLogFortianalyzer3OverrideFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterArrayOutput {
	return o
}

func (o LogFortianalyzer3OverrideFilterArrayOutput) Index(i pulumi.IntInput) LogFortianalyzer3OverrideFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogFortianalyzer3OverrideFilter {
		return vs[0].([]LogFortianalyzer3OverrideFilter)[vs[1].(int)]
	}).(LogFortianalyzer3OverrideFilterOutput)
}

type LogFortianalyzer3OverrideFilterMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3OverrideFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogFortianalyzer3OverrideFilter)(nil))
}

func (o LogFortianalyzer3OverrideFilterMapOutput) ToLogFortianalyzer3OverrideFilterMapOutput() LogFortianalyzer3OverrideFilterMapOutput {
	return o
}

func (o LogFortianalyzer3OverrideFilterMapOutput) ToLogFortianalyzer3OverrideFilterMapOutputWithContext(ctx context.Context) LogFortianalyzer3OverrideFilterMapOutput {
	return o
}

func (o LogFortianalyzer3OverrideFilterMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzer3OverrideFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogFortianalyzer3OverrideFilter {
		return vs[0].(map[string]LogFortianalyzer3OverrideFilter)[vs[1].(string)]
	}).(LogFortianalyzer3OverrideFilterOutput)
}

func init() {
	pulumi.RegisterOutputType(LogFortianalyzer3OverrideFilterOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3OverrideFilterPtrOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3OverrideFilterArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3OverrideFilterMapOutput{})
}
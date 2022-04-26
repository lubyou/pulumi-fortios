// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Filters for FortiAnalyzer Cloud. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// LogFortianalyzerCloud Filter can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzerCloudFilter:LogFortianalyzerCloudFilter labelname LogFortianalyzerCloudFilter
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzerCloudFilter:LogFortianalyzerCloudFilter labelname LogFortianalyzerCloudFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortianalyzerCloudFilter struct {
	pulumi.CustomResourceState

	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringOutput `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringOutput `pulumi:"dlpArchive"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Free style filter string.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringOutput `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringOutput `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles LogFortianalyzerCloudFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringOutput `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringOutput `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringOutput `pulumi:"multicastTraffic"`
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringOutput `pulumi:"snifferTraffic"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringOutput `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringOutput `pulumi:"ztnaTraffic"`
}

// NewLogFortianalyzerCloudFilter registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzerCloudFilter(ctx *pulumi.Context,
	name string, args *LogFortianalyzerCloudFilterArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzerCloudFilter, error) {
	if args == nil {
		args = &LogFortianalyzerCloudFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogFortianalyzerCloudFilter
	err := ctx.RegisterResource("fortios:index/logFortianalyzerCloudFilter:LogFortianalyzerCloudFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzerCloudFilter gets an existing LogFortianalyzerCloudFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzerCloudFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzerCloudFilterState, opts ...pulumi.ResourceOption) (*LogFortianalyzerCloudFilter, error) {
	var resource LogFortianalyzerCloudFilter
	err := ctx.ReadResource("fortios:index/logFortianalyzerCloudFilter:LogFortianalyzerCloudFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzerCloudFilter resources.
type logFortianalyzerCloudFilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive *string `pulumi:"dlpArchive"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Free style filter string.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []LogFortianalyzerCloudFilterFreeStyle `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp *string `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic *string `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic *string `pulumi:"multicastTraffic"`
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic *string `pulumi:"snifferTraffic"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip *string `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic *string `pulumi:"ztnaTraffic"`
}

type LogFortianalyzerCloudFilterState struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Free style filter string.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles LogFortianalyzerCloudFilterFreeStyleArrayInput
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringPtrInput
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringPtrInput
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringPtrInput
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringPtrInput
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringPtrInput
}

func (LogFortianalyzerCloudFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzerCloudFilterState)(nil)).Elem()
}

type logFortianalyzerCloudFilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly *string `pulumi:"anomaly"`
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive *string `pulumi:"dlpArchive"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Free style filter string.
	Filter *string `pulumi:"filter"`
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType *string `pulumi:"filterType"`
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic *string `pulumi:"forwardTraffic"`
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles []LogFortianalyzerCloudFilterFreeStyle `pulumi:"freeStyles"`
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp *string `pulumi:"gtp"`
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic *string `pulumi:"localTraffic"`
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic *string `pulumi:"multicastTraffic"`
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic *string `pulumi:"snifferTraffic"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip *string `pulumi:"voip"`
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic *string `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a LogFortianalyzerCloudFilter resource.
type LogFortianalyzerCloudFilterArgs struct {
	// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
	Anomaly pulumi.StringPtrInput
	// Enable/disable DLP archive logging. Valid values: `enable`, `disable`.
	DlpArchive pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Free style filter string.
	Filter pulumi.StringPtrInput
	// Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
	FilterType pulumi.StringPtrInput
	// Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
	ForwardTraffic pulumi.StringPtrInput
	// Free Style Filters The structure of `freeStyle` block is documented below.
	FreeStyles LogFortianalyzerCloudFilterFreeStyleArrayInput
	// Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
	Gtp pulumi.StringPtrInput
	// Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
	LocalTraffic pulumi.StringPtrInput
	// Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
	MulticastTraffic pulumi.StringPtrInput
	// Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
	SnifferTraffic pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable VoIP logging. Valid values: `enable`, `disable`.
	Voip pulumi.StringPtrInput
	// Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
	ZtnaTraffic pulumi.StringPtrInput
}

func (LogFortianalyzerCloudFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzerCloudFilterArgs)(nil)).Elem()
}

type LogFortianalyzerCloudFilterInput interface {
	pulumi.Input

	ToLogFortianalyzerCloudFilterOutput() LogFortianalyzerCloudFilterOutput
	ToLogFortianalyzerCloudFilterOutputWithContext(ctx context.Context) LogFortianalyzerCloudFilterOutput
}

func (*LogFortianalyzerCloudFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzerCloudFilter)(nil)).Elem()
}

func (i *LogFortianalyzerCloudFilter) ToLogFortianalyzerCloudFilterOutput() LogFortianalyzerCloudFilterOutput {
	return i.ToLogFortianalyzerCloudFilterOutputWithContext(context.Background())
}

func (i *LogFortianalyzerCloudFilter) ToLogFortianalyzerCloudFilterOutputWithContext(ctx context.Context) LogFortianalyzerCloudFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerCloudFilterOutput)
}

// LogFortianalyzerCloudFilterArrayInput is an input type that accepts LogFortianalyzerCloudFilterArray and LogFortianalyzerCloudFilterArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzerCloudFilterArrayInput` via:
//
//          LogFortianalyzerCloudFilterArray{ LogFortianalyzerCloudFilterArgs{...} }
type LogFortianalyzerCloudFilterArrayInput interface {
	pulumi.Input

	ToLogFortianalyzerCloudFilterArrayOutput() LogFortianalyzerCloudFilterArrayOutput
	ToLogFortianalyzerCloudFilterArrayOutputWithContext(context.Context) LogFortianalyzerCloudFilterArrayOutput
}

type LogFortianalyzerCloudFilterArray []LogFortianalyzerCloudFilterInput

func (LogFortianalyzerCloudFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzerCloudFilter)(nil)).Elem()
}

func (i LogFortianalyzerCloudFilterArray) ToLogFortianalyzerCloudFilterArrayOutput() LogFortianalyzerCloudFilterArrayOutput {
	return i.ToLogFortianalyzerCloudFilterArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzerCloudFilterArray) ToLogFortianalyzerCloudFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzerCloudFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerCloudFilterArrayOutput)
}

// LogFortianalyzerCloudFilterMapInput is an input type that accepts LogFortianalyzerCloudFilterMap and LogFortianalyzerCloudFilterMapOutput values.
// You can construct a concrete instance of `LogFortianalyzerCloudFilterMapInput` via:
//
//          LogFortianalyzerCloudFilterMap{ "key": LogFortianalyzerCloudFilterArgs{...} }
type LogFortianalyzerCloudFilterMapInput interface {
	pulumi.Input

	ToLogFortianalyzerCloudFilterMapOutput() LogFortianalyzerCloudFilterMapOutput
	ToLogFortianalyzerCloudFilterMapOutputWithContext(context.Context) LogFortianalyzerCloudFilterMapOutput
}

type LogFortianalyzerCloudFilterMap map[string]LogFortianalyzerCloudFilterInput

func (LogFortianalyzerCloudFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzerCloudFilter)(nil)).Elem()
}

func (i LogFortianalyzerCloudFilterMap) ToLogFortianalyzerCloudFilterMapOutput() LogFortianalyzerCloudFilterMapOutput {
	return i.ToLogFortianalyzerCloudFilterMapOutputWithContext(context.Background())
}

func (i LogFortianalyzerCloudFilterMap) ToLogFortianalyzerCloudFilterMapOutputWithContext(ctx context.Context) LogFortianalyzerCloudFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerCloudFilterMapOutput)
}

type LogFortianalyzerCloudFilterOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerCloudFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzerCloudFilter)(nil)).Elem()
}

func (o LogFortianalyzerCloudFilterOutput) ToLogFortianalyzerCloudFilterOutput() LogFortianalyzerCloudFilterOutput {
	return o
}

func (o LogFortianalyzerCloudFilterOutput) ToLogFortianalyzerCloudFilterOutputWithContext(ctx context.Context) LogFortianalyzerCloudFilterOutput {
	return o
}

type LogFortianalyzerCloudFilterArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerCloudFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzerCloudFilter)(nil)).Elem()
}

func (o LogFortianalyzerCloudFilterArrayOutput) ToLogFortianalyzerCloudFilterArrayOutput() LogFortianalyzerCloudFilterArrayOutput {
	return o
}

func (o LogFortianalyzerCloudFilterArrayOutput) ToLogFortianalyzerCloudFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzerCloudFilterArrayOutput {
	return o
}

func (o LogFortianalyzerCloudFilterArrayOutput) Index(i pulumi.IntInput) LogFortianalyzerCloudFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortianalyzerCloudFilter {
		return vs[0].([]*LogFortianalyzerCloudFilter)[vs[1].(int)]
	}).(LogFortianalyzerCloudFilterOutput)
}

type LogFortianalyzerCloudFilterMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerCloudFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzerCloudFilter)(nil)).Elem()
}

func (o LogFortianalyzerCloudFilterMapOutput) ToLogFortianalyzerCloudFilterMapOutput() LogFortianalyzerCloudFilterMapOutput {
	return o
}

func (o LogFortianalyzerCloudFilterMapOutput) ToLogFortianalyzerCloudFilterMapOutputWithContext(ctx context.Context) LogFortianalyzerCloudFilterMapOutput {
	return o
}

func (o LogFortianalyzerCloudFilterMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzerCloudFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortianalyzerCloudFilter {
		return vs[0].(map[string]*LogFortianalyzerCloudFilter)[vs[1].(string)]
	}).(LogFortianalyzerCloudFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerCloudFilterInput)(nil)).Elem(), &LogFortianalyzerCloudFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerCloudFilterArrayInput)(nil)).Elem(), LogFortianalyzerCloudFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerCloudFilterMapInput)(nil)).Elem(), LogFortianalyzerCloudFilterMap{})
	pulumi.RegisterOutputType(LogFortianalyzerCloudFilterOutput{})
	pulumi.RegisterOutputType(LogFortianalyzerCloudFilterArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzerCloudFilterMapOutput{})
}

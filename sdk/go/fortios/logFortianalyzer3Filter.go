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
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewLogFortianalyzer3Filter(ctx, "trname", &fortios.LogFortianalyzer3FilterArgs{
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
// LogFortianalyzer3 Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logFortianalyzer3Filter:LogFortianalyzer3Filter labelname LogFortianalyzer3Filter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortianalyzer3Filter struct {
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
	FreeStyles LogFortianalyzer3FilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogFortianalyzer3Filter registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzer3Filter(ctx *pulumi.Context,
	name string, args *LogFortianalyzer3FilterArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzer3Filter, error) {
	if args == nil {
		args = &LogFortianalyzer3FilterArgs{}
	}

	var resource LogFortianalyzer3Filter
	err := ctx.RegisterResource("fortios:index/logFortianalyzer3Filter:LogFortianalyzer3Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzer3Filter gets an existing LogFortianalyzer3Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzer3Filter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzer3FilterState, opts ...pulumi.ResourceOption) (*LogFortianalyzer3Filter, error) {
	var resource LogFortianalyzer3Filter
	err := ctx.ReadResource("fortios:index/logFortianalyzer3Filter:LogFortianalyzer3Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzer3Filter resources.
type logFortianalyzer3FilterState struct {
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
	FreeStyles []LogFortianalyzer3FilterFreeStyle `pulumi:"freeStyles"`
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

type LogFortianalyzer3FilterState struct {
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
	FreeStyles LogFortianalyzer3FilterFreeStyleArrayInput
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

func (LogFortianalyzer3FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3FilterState)(nil)).Elem()
}

type logFortianalyzer3FilterArgs struct {
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
	FreeStyles []LogFortianalyzer3FilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogFortianalyzer3Filter resource.
type LogFortianalyzer3FilterArgs struct {
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
	FreeStyles LogFortianalyzer3FilterFreeStyleArrayInput
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

func (LogFortianalyzer3FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3FilterArgs)(nil)).Elem()
}

type LogFortianalyzer3FilterInput interface {
	pulumi.Input

	ToLogFortianalyzer3FilterOutput() LogFortianalyzer3FilterOutput
	ToLogFortianalyzer3FilterOutputWithContext(ctx context.Context) LogFortianalyzer3FilterOutput
}

func (*LogFortianalyzer3Filter) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFortianalyzer3Filter)(nil))
}

func (i *LogFortianalyzer3Filter) ToLogFortianalyzer3FilterOutput() LogFortianalyzer3FilterOutput {
	return i.ToLogFortianalyzer3FilterOutputWithContext(context.Background())
}

func (i *LogFortianalyzer3Filter) ToLogFortianalyzer3FilterOutputWithContext(ctx context.Context) LogFortianalyzer3FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3FilterOutput)
}

func (i *LogFortianalyzer3Filter) ToLogFortianalyzer3FilterPtrOutput() LogFortianalyzer3FilterPtrOutput {
	return i.ToLogFortianalyzer3FilterPtrOutputWithContext(context.Background())
}

func (i *LogFortianalyzer3Filter) ToLogFortianalyzer3FilterPtrOutputWithContext(ctx context.Context) LogFortianalyzer3FilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3FilterPtrOutput)
}

type LogFortianalyzer3FilterPtrInput interface {
	pulumi.Input

	ToLogFortianalyzer3FilterPtrOutput() LogFortianalyzer3FilterPtrOutput
	ToLogFortianalyzer3FilterPtrOutputWithContext(ctx context.Context) LogFortianalyzer3FilterPtrOutput
}

type logFortianalyzer3FilterPtrType LogFortianalyzer3FilterArgs

func (*logFortianalyzer3FilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer3Filter)(nil))
}

func (i *logFortianalyzer3FilterPtrType) ToLogFortianalyzer3FilterPtrOutput() LogFortianalyzer3FilterPtrOutput {
	return i.ToLogFortianalyzer3FilterPtrOutputWithContext(context.Background())
}

func (i *logFortianalyzer3FilterPtrType) ToLogFortianalyzer3FilterPtrOutputWithContext(ctx context.Context) LogFortianalyzer3FilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3FilterPtrOutput)
}

// LogFortianalyzer3FilterArrayInput is an input type that accepts LogFortianalyzer3FilterArray and LogFortianalyzer3FilterArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzer3FilterArrayInput` via:
//
//          LogFortianalyzer3FilterArray{ LogFortianalyzer3FilterArgs{...} }
type LogFortianalyzer3FilterArrayInput interface {
	pulumi.Input

	ToLogFortianalyzer3FilterArrayOutput() LogFortianalyzer3FilterArrayOutput
	ToLogFortianalyzer3FilterArrayOutputWithContext(context.Context) LogFortianalyzer3FilterArrayOutput
}

type LogFortianalyzer3FilterArray []LogFortianalyzer3FilterInput

func (LogFortianalyzer3FilterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogFortianalyzer3Filter)(nil))
}

func (i LogFortianalyzer3FilterArray) ToLogFortianalyzer3FilterArrayOutput() LogFortianalyzer3FilterArrayOutput {
	return i.ToLogFortianalyzer3FilterArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzer3FilterArray) ToLogFortianalyzer3FilterArrayOutputWithContext(ctx context.Context) LogFortianalyzer3FilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3FilterArrayOutput)
}

// LogFortianalyzer3FilterMapInput is an input type that accepts LogFortianalyzer3FilterMap and LogFortianalyzer3FilterMapOutput values.
// You can construct a concrete instance of `LogFortianalyzer3FilterMapInput` via:
//
//          LogFortianalyzer3FilterMap{ "key": LogFortianalyzer3FilterArgs{...} }
type LogFortianalyzer3FilterMapInput interface {
	pulumi.Input

	ToLogFortianalyzer3FilterMapOutput() LogFortianalyzer3FilterMapOutput
	ToLogFortianalyzer3FilterMapOutputWithContext(context.Context) LogFortianalyzer3FilterMapOutput
}

type LogFortianalyzer3FilterMap map[string]LogFortianalyzer3FilterInput

func (LogFortianalyzer3FilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogFortianalyzer3Filter)(nil))
}

func (i LogFortianalyzer3FilterMap) ToLogFortianalyzer3FilterMapOutput() LogFortianalyzer3FilterMapOutput {
	return i.ToLogFortianalyzer3FilterMapOutputWithContext(context.Background())
}

func (i LogFortianalyzer3FilterMap) ToLogFortianalyzer3FilterMapOutputWithContext(ctx context.Context) LogFortianalyzer3FilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3FilterMapOutput)
}

type LogFortianalyzer3FilterOutput struct {
	*pulumi.OutputState
}

func (LogFortianalyzer3FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFortianalyzer3Filter)(nil))
}

func (o LogFortianalyzer3FilterOutput) ToLogFortianalyzer3FilterOutput() LogFortianalyzer3FilterOutput {
	return o
}

func (o LogFortianalyzer3FilterOutput) ToLogFortianalyzer3FilterOutputWithContext(ctx context.Context) LogFortianalyzer3FilterOutput {
	return o
}

func (o LogFortianalyzer3FilterOutput) ToLogFortianalyzer3FilterPtrOutput() LogFortianalyzer3FilterPtrOutput {
	return o.ToLogFortianalyzer3FilterPtrOutputWithContext(context.Background())
}

func (o LogFortianalyzer3FilterOutput) ToLogFortianalyzer3FilterPtrOutputWithContext(ctx context.Context) LogFortianalyzer3FilterPtrOutput {
	return o.ApplyT(func(v LogFortianalyzer3Filter) *LogFortianalyzer3Filter {
		return &v
	}).(LogFortianalyzer3FilterPtrOutput)
}

type LogFortianalyzer3FilterPtrOutput struct {
	*pulumi.OutputState
}

func (LogFortianalyzer3FilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer3Filter)(nil))
}

func (o LogFortianalyzer3FilterPtrOutput) ToLogFortianalyzer3FilterPtrOutput() LogFortianalyzer3FilterPtrOutput {
	return o
}

func (o LogFortianalyzer3FilterPtrOutput) ToLogFortianalyzer3FilterPtrOutputWithContext(ctx context.Context) LogFortianalyzer3FilterPtrOutput {
	return o
}

type LogFortianalyzer3FilterArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3FilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogFortianalyzer3Filter)(nil))
}

func (o LogFortianalyzer3FilterArrayOutput) ToLogFortianalyzer3FilterArrayOutput() LogFortianalyzer3FilterArrayOutput {
	return o
}

func (o LogFortianalyzer3FilterArrayOutput) ToLogFortianalyzer3FilterArrayOutputWithContext(ctx context.Context) LogFortianalyzer3FilterArrayOutput {
	return o
}

func (o LogFortianalyzer3FilterArrayOutput) Index(i pulumi.IntInput) LogFortianalyzer3FilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogFortianalyzer3Filter {
		return vs[0].([]LogFortianalyzer3Filter)[vs[1].(int)]
	}).(LogFortianalyzer3FilterOutput)
}

type LogFortianalyzer3FilterMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3FilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogFortianalyzer3Filter)(nil))
}

func (o LogFortianalyzer3FilterMapOutput) ToLogFortianalyzer3FilterMapOutput() LogFortianalyzer3FilterMapOutput {
	return o
}

func (o LogFortianalyzer3FilterMapOutput) ToLogFortianalyzer3FilterMapOutputWithContext(ctx context.Context) LogFortianalyzer3FilterMapOutput {
	return o
}

func (o LogFortianalyzer3FilterMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzer3FilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogFortianalyzer3Filter {
		return vs[0].(map[string]LogFortianalyzer3Filter)[vs[1].(string)]
	}).(LogFortianalyzer3FilterOutput)
}

func init() {
	pulumi.RegisterOutputType(LogFortianalyzer3FilterOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3FilterPtrOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3FilterArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3FilterMapOutput{})
}
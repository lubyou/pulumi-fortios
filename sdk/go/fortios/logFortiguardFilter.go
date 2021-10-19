// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Filters for FortiCloud.
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
// 		_, err := fortios.NewLogFortiguardFilter(ctx, "trname", &fortios.LogFortiguardFilterArgs{
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
// LogFortiguard Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logFortiguardFilter:LogFortiguardFilter labelname LogFortiguardFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogFortiguardFilter struct {
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
	FreeStyles LogFortiguardFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogFortiguardFilter registers a new resource with the given unique name, arguments, and options.
func NewLogFortiguardFilter(ctx *pulumi.Context,
	name string, args *LogFortiguardFilterArgs, opts ...pulumi.ResourceOption) (*LogFortiguardFilter, error) {
	if args == nil {
		args = &LogFortiguardFilterArgs{}
	}

	var resource LogFortiguardFilter
	err := ctx.RegisterResource("fortios:index/logFortiguardFilter:LogFortiguardFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortiguardFilter gets an existing LogFortiguardFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortiguardFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortiguardFilterState, opts ...pulumi.ResourceOption) (*LogFortiguardFilter, error) {
	var resource LogFortiguardFilter
	err := ctx.ReadResource("fortios:index/logFortiguardFilter:LogFortiguardFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortiguardFilter resources.
type logFortiguardFilterState struct {
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
	FreeStyles []LogFortiguardFilterFreeStyle `pulumi:"freeStyles"`
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

type LogFortiguardFilterState struct {
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
	FreeStyles LogFortiguardFilterFreeStyleArrayInput
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

func (LogFortiguardFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortiguardFilterState)(nil)).Elem()
}

type logFortiguardFilterArgs struct {
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
	FreeStyles []LogFortiguardFilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogFortiguardFilter resource.
type LogFortiguardFilterArgs struct {
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
	FreeStyles LogFortiguardFilterFreeStyleArrayInput
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

func (LogFortiguardFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortiguardFilterArgs)(nil)).Elem()
}

type LogFortiguardFilterInput interface {
	pulumi.Input

	ToLogFortiguardFilterOutput() LogFortiguardFilterOutput
	ToLogFortiguardFilterOutputWithContext(ctx context.Context) LogFortiguardFilterOutput
}

func (*LogFortiguardFilter) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFortiguardFilter)(nil))
}

func (i *LogFortiguardFilter) ToLogFortiguardFilterOutput() LogFortiguardFilterOutput {
	return i.ToLogFortiguardFilterOutputWithContext(context.Background())
}

func (i *LogFortiguardFilter) ToLogFortiguardFilterOutputWithContext(ctx context.Context) LogFortiguardFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiguardFilterOutput)
}

func (i *LogFortiguardFilter) ToLogFortiguardFilterPtrOutput() LogFortiguardFilterPtrOutput {
	return i.ToLogFortiguardFilterPtrOutputWithContext(context.Background())
}

func (i *LogFortiguardFilter) ToLogFortiguardFilterPtrOutputWithContext(ctx context.Context) LogFortiguardFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiguardFilterPtrOutput)
}

type LogFortiguardFilterPtrInput interface {
	pulumi.Input

	ToLogFortiguardFilterPtrOutput() LogFortiguardFilterPtrOutput
	ToLogFortiguardFilterPtrOutputWithContext(ctx context.Context) LogFortiguardFilterPtrOutput
}

type logFortiguardFilterPtrType LogFortiguardFilterArgs

func (*logFortiguardFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortiguardFilter)(nil))
}

func (i *logFortiguardFilterPtrType) ToLogFortiguardFilterPtrOutput() LogFortiguardFilterPtrOutput {
	return i.ToLogFortiguardFilterPtrOutputWithContext(context.Background())
}

func (i *logFortiguardFilterPtrType) ToLogFortiguardFilterPtrOutputWithContext(ctx context.Context) LogFortiguardFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiguardFilterPtrOutput)
}

// LogFortiguardFilterArrayInput is an input type that accepts LogFortiguardFilterArray and LogFortiguardFilterArrayOutput values.
// You can construct a concrete instance of `LogFortiguardFilterArrayInput` via:
//
//          LogFortiguardFilterArray{ LogFortiguardFilterArgs{...} }
type LogFortiguardFilterArrayInput interface {
	pulumi.Input

	ToLogFortiguardFilterArrayOutput() LogFortiguardFilterArrayOutput
	ToLogFortiguardFilterArrayOutputWithContext(context.Context) LogFortiguardFilterArrayOutput
}

type LogFortiguardFilterArray []LogFortiguardFilterInput

func (LogFortiguardFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogFortiguardFilter)(nil))
}

func (i LogFortiguardFilterArray) ToLogFortiguardFilterArrayOutput() LogFortiguardFilterArrayOutput {
	return i.ToLogFortiguardFilterArrayOutputWithContext(context.Background())
}

func (i LogFortiguardFilterArray) ToLogFortiguardFilterArrayOutputWithContext(ctx context.Context) LogFortiguardFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiguardFilterArrayOutput)
}

// LogFortiguardFilterMapInput is an input type that accepts LogFortiguardFilterMap and LogFortiguardFilterMapOutput values.
// You can construct a concrete instance of `LogFortiguardFilterMapInput` via:
//
//          LogFortiguardFilterMap{ "key": LogFortiguardFilterArgs{...} }
type LogFortiguardFilterMapInput interface {
	pulumi.Input

	ToLogFortiguardFilterMapOutput() LogFortiguardFilterMapOutput
	ToLogFortiguardFilterMapOutputWithContext(context.Context) LogFortiguardFilterMapOutput
}

type LogFortiguardFilterMap map[string]LogFortiguardFilterInput

func (LogFortiguardFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogFortiguardFilter)(nil))
}

func (i LogFortiguardFilterMap) ToLogFortiguardFilterMapOutput() LogFortiguardFilterMapOutput {
	return i.ToLogFortiguardFilterMapOutputWithContext(context.Background())
}

func (i LogFortiguardFilterMap) ToLogFortiguardFilterMapOutputWithContext(ctx context.Context) LogFortiguardFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiguardFilterMapOutput)
}

type LogFortiguardFilterOutput struct {
	*pulumi.OutputState
}

func (LogFortiguardFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFortiguardFilter)(nil))
}

func (o LogFortiguardFilterOutput) ToLogFortiguardFilterOutput() LogFortiguardFilterOutput {
	return o
}

func (o LogFortiguardFilterOutput) ToLogFortiguardFilterOutputWithContext(ctx context.Context) LogFortiguardFilterOutput {
	return o
}

func (o LogFortiguardFilterOutput) ToLogFortiguardFilterPtrOutput() LogFortiguardFilterPtrOutput {
	return o.ToLogFortiguardFilterPtrOutputWithContext(context.Background())
}

func (o LogFortiguardFilterOutput) ToLogFortiguardFilterPtrOutputWithContext(ctx context.Context) LogFortiguardFilterPtrOutput {
	return o.ApplyT(func(v LogFortiguardFilter) *LogFortiguardFilter {
		return &v
	}).(LogFortiguardFilterPtrOutput)
}

type LogFortiguardFilterPtrOutput struct {
	*pulumi.OutputState
}

func (LogFortiguardFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortiguardFilter)(nil))
}

func (o LogFortiguardFilterPtrOutput) ToLogFortiguardFilterPtrOutput() LogFortiguardFilterPtrOutput {
	return o
}

func (o LogFortiguardFilterPtrOutput) ToLogFortiguardFilterPtrOutputWithContext(ctx context.Context) LogFortiguardFilterPtrOutput {
	return o
}

type LogFortiguardFilterArrayOutput struct{ *pulumi.OutputState }

func (LogFortiguardFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogFortiguardFilter)(nil))
}

func (o LogFortiguardFilterArrayOutput) ToLogFortiguardFilterArrayOutput() LogFortiguardFilterArrayOutput {
	return o
}

func (o LogFortiguardFilterArrayOutput) ToLogFortiguardFilterArrayOutputWithContext(ctx context.Context) LogFortiguardFilterArrayOutput {
	return o
}

func (o LogFortiguardFilterArrayOutput) Index(i pulumi.IntInput) LogFortiguardFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogFortiguardFilter {
		return vs[0].([]LogFortiguardFilter)[vs[1].(int)]
	}).(LogFortiguardFilterOutput)
}

type LogFortiguardFilterMapOutput struct{ *pulumi.OutputState }

func (LogFortiguardFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogFortiguardFilter)(nil))
}

func (o LogFortiguardFilterMapOutput) ToLogFortiguardFilterMapOutput() LogFortiguardFilterMapOutput {
	return o
}

func (o LogFortiguardFilterMapOutput) ToLogFortiguardFilterMapOutputWithContext(ctx context.Context) LogFortiguardFilterMapOutput {
	return o
}

func (o LogFortiguardFilterMapOutput) MapIndex(k pulumi.StringInput) LogFortiguardFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogFortiguardFilter {
		return vs[0].(map[string]LogFortiguardFilter)[vs[1].(string)]
	}).(LogFortiguardFilterOutput)
}

func init() {
	pulumi.RegisterOutputType(LogFortiguardFilterOutput{})
	pulumi.RegisterOutputType(LogFortiguardFilterPtrOutput{})
	pulumi.RegisterOutputType(LogFortiguardFilterArrayOutput{})
	pulumi.RegisterOutputType(LogFortiguardFilterMapOutput{})
}

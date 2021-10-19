// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Filters for remote system server.
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
// 		_, err := fortios.NewLogSyslogd3Filter(ctx, "trname", &fortios.LogSyslogd3FilterArgs{
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
// LogSyslogd3 Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd3Filter:LogSyslogd3Filter labelname LogSyslogd3Filter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSyslogd3Filter struct {
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
	FreeStyles LogSyslogd3FilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogSyslogd3Filter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd3Filter(ctx *pulumi.Context,
	name string, args *LogSyslogd3FilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogd3Filter, error) {
	if args == nil {
		args = &LogSyslogd3FilterArgs{}
	}

	var resource LogSyslogd3Filter
	err := ctx.RegisterResource("fortios:index/logSyslogd3Filter:LogSyslogd3Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogd3Filter gets an existing LogSyslogd3Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogd3Filter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogd3FilterState, opts ...pulumi.ResourceOption) (*LogSyslogd3Filter, error) {
	var resource LogSyslogd3Filter
	err := ctx.ReadResource("fortios:index/logSyslogd3Filter:LogSyslogd3Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogd3Filter resources.
type logSyslogd3FilterState struct {
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
	FreeStyles []LogSyslogd3FilterFreeStyle `pulumi:"freeStyles"`
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

type LogSyslogd3FilterState struct {
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
	FreeStyles LogSyslogd3FilterFreeStyleArrayInput
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

func (LogSyslogd3FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd3FilterState)(nil)).Elem()
}

type logSyslogd3FilterArgs struct {
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
	FreeStyles []LogSyslogd3FilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogSyslogd3Filter resource.
type LogSyslogd3FilterArgs struct {
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
	FreeStyles LogSyslogd3FilterFreeStyleArrayInput
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

func (LogSyslogd3FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd3FilterArgs)(nil)).Elem()
}

type LogSyslogd3FilterInput interface {
	pulumi.Input

	ToLogSyslogd3FilterOutput() LogSyslogd3FilterOutput
	ToLogSyslogd3FilterOutputWithContext(ctx context.Context) LogSyslogd3FilterOutput
}

func (*LogSyslogd3Filter) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSyslogd3Filter)(nil))
}

func (i *LogSyslogd3Filter) ToLogSyslogd3FilterOutput() LogSyslogd3FilterOutput {
	return i.ToLogSyslogd3FilterOutputWithContext(context.Background())
}

func (i *LogSyslogd3Filter) ToLogSyslogd3FilterOutputWithContext(ctx context.Context) LogSyslogd3FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd3FilterOutput)
}

func (i *LogSyslogd3Filter) ToLogSyslogd3FilterPtrOutput() LogSyslogd3FilterPtrOutput {
	return i.ToLogSyslogd3FilterPtrOutputWithContext(context.Background())
}

func (i *LogSyslogd3Filter) ToLogSyslogd3FilterPtrOutputWithContext(ctx context.Context) LogSyslogd3FilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd3FilterPtrOutput)
}

type LogSyslogd3FilterPtrInput interface {
	pulumi.Input

	ToLogSyslogd3FilterPtrOutput() LogSyslogd3FilterPtrOutput
	ToLogSyslogd3FilterPtrOutputWithContext(ctx context.Context) LogSyslogd3FilterPtrOutput
}

type logSyslogd3FilterPtrType LogSyslogd3FilterArgs

func (*logSyslogd3FilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd3Filter)(nil))
}

func (i *logSyslogd3FilterPtrType) ToLogSyslogd3FilterPtrOutput() LogSyslogd3FilterPtrOutput {
	return i.ToLogSyslogd3FilterPtrOutputWithContext(context.Background())
}

func (i *logSyslogd3FilterPtrType) ToLogSyslogd3FilterPtrOutputWithContext(ctx context.Context) LogSyslogd3FilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd3FilterPtrOutput)
}

// LogSyslogd3FilterArrayInput is an input type that accepts LogSyslogd3FilterArray and LogSyslogd3FilterArrayOutput values.
// You can construct a concrete instance of `LogSyslogd3FilterArrayInput` via:
//
//          LogSyslogd3FilterArray{ LogSyslogd3FilterArgs{...} }
type LogSyslogd3FilterArrayInput interface {
	pulumi.Input

	ToLogSyslogd3FilterArrayOutput() LogSyslogd3FilterArrayOutput
	ToLogSyslogd3FilterArrayOutputWithContext(context.Context) LogSyslogd3FilterArrayOutput
}

type LogSyslogd3FilterArray []LogSyslogd3FilterInput

func (LogSyslogd3FilterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogSyslogd3Filter)(nil))
}

func (i LogSyslogd3FilterArray) ToLogSyslogd3FilterArrayOutput() LogSyslogd3FilterArrayOutput {
	return i.ToLogSyslogd3FilterArrayOutputWithContext(context.Background())
}

func (i LogSyslogd3FilterArray) ToLogSyslogd3FilterArrayOutputWithContext(ctx context.Context) LogSyslogd3FilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd3FilterArrayOutput)
}

// LogSyslogd3FilterMapInput is an input type that accepts LogSyslogd3FilterMap and LogSyslogd3FilterMapOutput values.
// You can construct a concrete instance of `LogSyslogd3FilterMapInput` via:
//
//          LogSyslogd3FilterMap{ "key": LogSyslogd3FilterArgs{...} }
type LogSyslogd3FilterMapInput interface {
	pulumi.Input

	ToLogSyslogd3FilterMapOutput() LogSyslogd3FilterMapOutput
	ToLogSyslogd3FilterMapOutputWithContext(context.Context) LogSyslogd3FilterMapOutput
}

type LogSyslogd3FilterMap map[string]LogSyslogd3FilterInput

func (LogSyslogd3FilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogSyslogd3Filter)(nil))
}

func (i LogSyslogd3FilterMap) ToLogSyslogd3FilterMapOutput() LogSyslogd3FilterMapOutput {
	return i.ToLogSyslogd3FilterMapOutputWithContext(context.Background())
}

func (i LogSyslogd3FilterMap) ToLogSyslogd3FilterMapOutputWithContext(ctx context.Context) LogSyslogd3FilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd3FilterMapOutput)
}

type LogSyslogd3FilterOutput struct {
	*pulumi.OutputState
}

func (LogSyslogd3FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSyslogd3Filter)(nil))
}

func (o LogSyslogd3FilterOutput) ToLogSyslogd3FilterOutput() LogSyslogd3FilterOutput {
	return o
}

func (o LogSyslogd3FilterOutput) ToLogSyslogd3FilterOutputWithContext(ctx context.Context) LogSyslogd3FilterOutput {
	return o
}

func (o LogSyslogd3FilterOutput) ToLogSyslogd3FilterPtrOutput() LogSyslogd3FilterPtrOutput {
	return o.ToLogSyslogd3FilterPtrOutputWithContext(context.Background())
}

func (o LogSyslogd3FilterOutput) ToLogSyslogd3FilterPtrOutputWithContext(ctx context.Context) LogSyslogd3FilterPtrOutput {
	return o.ApplyT(func(v LogSyslogd3Filter) *LogSyslogd3Filter {
		return &v
	}).(LogSyslogd3FilterPtrOutput)
}

type LogSyslogd3FilterPtrOutput struct {
	*pulumi.OutputState
}

func (LogSyslogd3FilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd3Filter)(nil))
}

func (o LogSyslogd3FilterPtrOutput) ToLogSyslogd3FilterPtrOutput() LogSyslogd3FilterPtrOutput {
	return o
}

func (o LogSyslogd3FilterPtrOutput) ToLogSyslogd3FilterPtrOutputWithContext(ctx context.Context) LogSyslogd3FilterPtrOutput {
	return o
}

type LogSyslogd3FilterArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogd3FilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSyslogd3Filter)(nil))
}

func (o LogSyslogd3FilterArrayOutput) ToLogSyslogd3FilterArrayOutput() LogSyslogd3FilterArrayOutput {
	return o
}

func (o LogSyslogd3FilterArrayOutput) ToLogSyslogd3FilterArrayOutputWithContext(ctx context.Context) LogSyslogd3FilterArrayOutput {
	return o
}

func (o LogSyslogd3FilterArrayOutput) Index(i pulumi.IntInput) LogSyslogd3FilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogSyslogd3Filter {
		return vs[0].([]LogSyslogd3Filter)[vs[1].(int)]
	}).(LogSyslogd3FilterOutput)
}

type LogSyslogd3FilterMapOutput struct{ *pulumi.OutputState }

func (LogSyslogd3FilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogSyslogd3Filter)(nil))
}

func (o LogSyslogd3FilterMapOutput) ToLogSyslogd3FilterMapOutput() LogSyslogd3FilterMapOutput {
	return o
}

func (o LogSyslogd3FilterMapOutput) ToLogSyslogd3FilterMapOutputWithContext(ctx context.Context) LogSyslogd3FilterMapOutput {
	return o
}

func (o LogSyslogd3FilterMapOutput) MapIndex(k pulumi.StringInput) LogSyslogd3FilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogSyslogd3Filter {
		return vs[0].(map[string]LogSyslogd3Filter)[vs[1].(string)]
	}).(LogSyslogd3FilterOutput)
}

func init() {
	pulumi.RegisterOutputType(LogSyslogd3FilterOutput{})
	pulumi.RegisterOutputType(LogSyslogd3FilterPtrOutput{})
	pulumi.RegisterOutputType(LogSyslogd3FilterArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogd3FilterMapOutput{})
}

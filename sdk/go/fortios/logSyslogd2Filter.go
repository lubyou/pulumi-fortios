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
// 		_, err := fortios.NewLogSyslogd2Filter(ctx, "trname", &fortios.LogSyslogd2FilterArgs{
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
// LogSyslogd2 Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd2Filter:LogSyslogd2Filter labelname LogSyslogd2Filter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSyslogd2Filter struct {
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
	FreeStyles LogSyslogd2FilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogSyslogd2Filter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd2Filter(ctx *pulumi.Context,
	name string, args *LogSyslogd2FilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogd2Filter, error) {
	if args == nil {
		args = &LogSyslogd2FilterArgs{}
	}

	var resource LogSyslogd2Filter
	err := ctx.RegisterResource("fortios:index/logSyslogd2Filter:LogSyslogd2Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogd2Filter gets an existing LogSyslogd2Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogd2Filter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogd2FilterState, opts ...pulumi.ResourceOption) (*LogSyslogd2Filter, error) {
	var resource LogSyslogd2Filter
	err := ctx.ReadResource("fortios:index/logSyslogd2Filter:LogSyslogd2Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogd2Filter resources.
type logSyslogd2FilterState struct {
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
	FreeStyles []LogSyslogd2FilterFreeStyle `pulumi:"freeStyles"`
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

type LogSyslogd2FilterState struct {
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
	FreeStyles LogSyslogd2FilterFreeStyleArrayInput
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

func (LogSyslogd2FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd2FilterState)(nil)).Elem()
}

type logSyslogd2FilterArgs struct {
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
	FreeStyles []LogSyslogd2FilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogSyslogd2Filter resource.
type LogSyslogd2FilterArgs struct {
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
	FreeStyles LogSyslogd2FilterFreeStyleArrayInput
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

func (LogSyslogd2FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd2FilterArgs)(nil)).Elem()
}

type LogSyslogd2FilterInput interface {
	pulumi.Input

	ToLogSyslogd2FilterOutput() LogSyslogd2FilterOutput
	ToLogSyslogd2FilterOutputWithContext(ctx context.Context) LogSyslogd2FilterOutput
}

func (*LogSyslogd2Filter) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSyslogd2Filter)(nil))
}

func (i *LogSyslogd2Filter) ToLogSyslogd2FilterOutput() LogSyslogd2FilterOutput {
	return i.ToLogSyslogd2FilterOutputWithContext(context.Background())
}

func (i *LogSyslogd2Filter) ToLogSyslogd2FilterOutputWithContext(ctx context.Context) LogSyslogd2FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2FilterOutput)
}

func (i *LogSyslogd2Filter) ToLogSyslogd2FilterPtrOutput() LogSyslogd2FilterPtrOutput {
	return i.ToLogSyslogd2FilterPtrOutputWithContext(context.Background())
}

func (i *LogSyslogd2Filter) ToLogSyslogd2FilterPtrOutputWithContext(ctx context.Context) LogSyslogd2FilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2FilterPtrOutput)
}

type LogSyslogd2FilterPtrInput interface {
	pulumi.Input

	ToLogSyslogd2FilterPtrOutput() LogSyslogd2FilterPtrOutput
	ToLogSyslogd2FilterPtrOutputWithContext(ctx context.Context) LogSyslogd2FilterPtrOutput
}

type logSyslogd2FilterPtrType LogSyslogd2FilterArgs

func (*logSyslogd2FilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd2Filter)(nil))
}

func (i *logSyslogd2FilterPtrType) ToLogSyslogd2FilterPtrOutput() LogSyslogd2FilterPtrOutput {
	return i.ToLogSyslogd2FilterPtrOutputWithContext(context.Background())
}

func (i *logSyslogd2FilterPtrType) ToLogSyslogd2FilterPtrOutputWithContext(ctx context.Context) LogSyslogd2FilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2FilterPtrOutput)
}

// LogSyslogd2FilterArrayInput is an input type that accepts LogSyslogd2FilterArray and LogSyslogd2FilterArrayOutput values.
// You can construct a concrete instance of `LogSyslogd2FilterArrayInput` via:
//
//          LogSyslogd2FilterArray{ LogSyslogd2FilterArgs{...} }
type LogSyslogd2FilterArrayInput interface {
	pulumi.Input

	ToLogSyslogd2FilterArrayOutput() LogSyslogd2FilterArrayOutput
	ToLogSyslogd2FilterArrayOutputWithContext(context.Context) LogSyslogd2FilterArrayOutput
}

type LogSyslogd2FilterArray []LogSyslogd2FilterInput

func (LogSyslogd2FilterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogSyslogd2Filter)(nil))
}

func (i LogSyslogd2FilterArray) ToLogSyslogd2FilterArrayOutput() LogSyslogd2FilterArrayOutput {
	return i.ToLogSyslogd2FilterArrayOutputWithContext(context.Background())
}

func (i LogSyslogd2FilterArray) ToLogSyslogd2FilterArrayOutputWithContext(ctx context.Context) LogSyslogd2FilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2FilterArrayOutput)
}

// LogSyslogd2FilterMapInput is an input type that accepts LogSyslogd2FilterMap and LogSyslogd2FilterMapOutput values.
// You can construct a concrete instance of `LogSyslogd2FilterMapInput` via:
//
//          LogSyslogd2FilterMap{ "key": LogSyslogd2FilterArgs{...} }
type LogSyslogd2FilterMapInput interface {
	pulumi.Input

	ToLogSyslogd2FilterMapOutput() LogSyslogd2FilterMapOutput
	ToLogSyslogd2FilterMapOutputWithContext(context.Context) LogSyslogd2FilterMapOutput
}

type LogSyslogd2FilterMap map[string]LogSyslogd2FilterInput

func (LogSyslogd2FilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogSyslogd2Filter)(nil))
}

func (i LogSyslogd2FilterMap) ToLogSyslogd2FilterMapOutput() LogSyslogd2FilterMapOutput {
	return i.ToLogSyslogd2FilterMapOutputWithContext(context.Background())
}

func (i LogSyslogd2FilterMap) ToLogSyslogd2FilterMapOutputWithContext(ctx context.Context) LogSyslogd2FilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2FilterMapOutput)
}

type LogSyslogd2FilterOutput struct {
	*pulumi.OutputState
}

func (LogSyslogd2FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSyslogd2Filter)(nil))
}

func (o LogSyslogd2FilterOutput) ToLogSyslogd2FilterOutput() LogSyslogd2FilterOutput {
	return o
}

func (o LogSyslogd2FilterOutput) ToLogSyslogd2FilterOutputWithContext(ctx context.Context) LogSyslogd2FilterOutput {
	return o
}

func (o LogSyslogd2FilterOutput) ToLogSyslogd2FilterPtrOutput() LogSyslogd2FilterPtrOutput {
	return o.ToLogSyslogd2FilterPtrOutputWithContext(context.Background())
}

func (o LogSyslogd2FilterOutput) ToLogSyslogd2FilterPtrOutputWithContext(ctx context.Context) LogSyslogd2FilterPtrOutput {
	return o.ApplyT(func(v LogSyslogd2Filter) *LogSyslogd2Filter {
		return &v
	}).(LogSyslogd2FilterPtrOutput)
}

type LogSyslogd2FilterPtrOutput struct {
	*pulumi.OutputState
}

func (LogSyslogd2FilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd2Filter)(nil))
}

func (o LogSyslogd2FilterPtrOutput) ToLogSyslogd2FilterPtrOutput() LogSyslogd2FilterPtrOutput {
	return o
}

func (o LogSyslogd2FilterPtrOutput) ToLogSyslogd2FilterPtrOutputWithContext(ctx context.Context) LogSyslogd2FilterPtrOutput {
	return o
}

type LogSyslogd2FilterArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogd2FilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSyslogd2Filter)(nil))
}

func (o LogSyslogd2FilterArrayOutput) ToLogSyslogd2FilterArrayOutput() LogSyslogd2FilterArrayOutput {
	return o
}

func (o LogSyslogd2FilterArrayOutput) ToLogSyslogd2FilterArrayOutputWithContext(ctx context.Context) LogSyslogd2FilterArrayOutput {
	return o
}

func (o LogSyslogd2FilterArrayOutput) Index(i pulumi.IntInput) LogSyslogd2FilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogSyslogd2Filter {
		return vs[0].([]LogSyslogd2Filter)[vs[1].(int)]
	}).(LogSyslogd2FilterOutput)
}

type LogSyslogd2FilterMapOutput struct{ *pulumi.OutputState }

func (LogSyslogd2FilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogSyslogd2Filter)(nil))
}

func (o LogSyslogd2FilterMapOutput) ToLogSyslogd2FilterMapOutput() LogSyslogd2FilterMapOutput {
	return o
}

func (o LogSyslogd2FilterMapOutput) ToLogSyslogd2FilterMapOutputWithContext(ctx context.Context) LogSyslogd2FilterMapOutput {
	return o
}

func (o LogSyslogd2FilterMapOutput) MapIndex(k pulumi.StringInput) LogSyslogd2FilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogSyslogd2Filter {
		return vs[0].(map[string]LogSyslogd2Filter)[vs[1].(string)]
	}).(LogSyslogd2FilterOutput)
}

func init() {
	pulumi.RegisterOutputType(LogSyslogd2FilterOutput{})
	pulumi.RegisterOutputType(LogSyslogd2FilterPtrOutput{})
	pulumi.RegisterOutputType(LogSyslogd2FilterArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogd2FilterMapOutput{})
}

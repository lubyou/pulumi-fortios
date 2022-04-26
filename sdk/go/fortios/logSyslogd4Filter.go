// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

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
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewLogSyslogd4Filter(ctx, "trname", &fortios.LogSyslogd4FilterArgs{
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
// LogSyslogd4 Filter can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd4Filter:LogSyslogd4Filter labelname LogSyslogd4Filter
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/logSyslogd4Filter:LogSyslogd4Filter labelname LogSyslogd4Filter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogSyslogd4Filter struct {
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
	FreeStyles LogSyslogd4FilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogSyslogd4Filter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd4Filter(ctx *pulumi.Context,
	name string, args *LogSyslogd4FilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogd4Filter, error) {
	if args == nil {
		args = &LogSyslogd4FilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogd4Filter
	err := ctx.RegisterResource("fortios:index/logSyslogd4Filter:LogSyslogd4Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogd4Filter gets an existing LogSyslogd4Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogd4Filter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogd4FilterState, opts ...pulumi.ResourceOption) (*LogSyslogd4Filter, error) {
	var resource LogSyslogd4Filter
	err := ctx.ReadResource("fortios:index/logSyslogd4Filter:LogSyslogd4Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogd4Filter resources.
type logSyslogd4FilterState struct {
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
	FreeStyles []LogSyslogd4FilterFreeStyle `pulumi:"freeStyles"`
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

type LogSyslogd4FilterState struct {
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
	FreeStyles LogSyslogd4FilterFreeStyleArrayInput
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

func (LogSyslogd4FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd4FilterState)(nil)).Elem()
}

type logSyslogd4FilterArgs struct {
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
	FreeStyles []LogSyslogd4FilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogSyslogd4Filter resource.
type LogSyslogd4FilterArgs struct {
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
	FreeStyles LogSyslogd4FilterFreeStyleArrayInput
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

func (LogSyslogd4FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd4FilterArgs)(nil)).Elem()
}

type LogSyslogd4FilterInput interface {
	pulumi.Input

	ToLogSyslogd4FilterOutput() LogSyslogd4FilterOutput
	ToLogSyslogd4FilterOutputWithContext(ctx context.Context) LogSyslogd4FilterOutput
}

func (*LogSyslogd4Filter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd4Filter)(nil)).Elem()
}

func (i *LogSyslogd4Filter) ToLogSyslogd4FilterOutput() LogSyslogd4FilterOutput {
	return i.ToLogSyslogd4FilterOutputWithContext(context.Background())
}

func (i *LogSyslogd4Filter) ToLogSyslogd4FilterOutputWithContext(ctx context.Context) LogSyslogd4FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4FilterOutput)
}

// LogSyslogd4FilterArrayInput is an input type that accepts LogSyslogd4FilterArray and LogSyslogd4FilterArrayOutput values.
// You can construct a concrete instance of `LogSyslogd4FilterArrayInput` via:
//
//          LogSyslogd4FilterArray{ LogSyslogd4FilterArgs{...} }
type LogSyslogd4FilterArrayInput interface {
	pulumi.Input

	ToLogSyslogd4FilterArrayOutput() LogSyslogd4FilterArrayOutput
	ToLogSyslogd4FilterArrayOutputWithContext(context.Context) LogSyslogd4FilterArrayOutput
}

type LogSyslogd4FilterArray []LogSyslogd4FilterInput

func (LogSyslogd4FilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd4Filter)(nil)).Elem()
}

func (i LogSyslogd4FilterArray) ToLogSyslogd4FilterArrayOutput() LogSyslogd4FilterArrayOutput {
	return i.ToLogSyslogd4FilterArrayOutputWithContext(context.Background())
}

func (i LogSyslogd4FilterArray) ToLogSyslogd4FilterArrayOutputWithContext(ctx context.Context) LogSyslogd4FilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4FilterArrayOutput)
}

// LogSyslogd4FilterMapInput is an input type that accepts LogSyslogd4FilterMap and LogSyslogd4FilterMapOutput values.
// You can construct a concrete instance of `LogSyslogd4FilterMapInput` via:
//
//          LogSyslogd4FilterMap{ "key": LogSyslogd4FilterArgs{...} }
type LogSyslogd4FilterMapInput interface {
	pulumi.Input

	ToLogSyslogd4FilterMapOutput() LogSyslogd4FilterMapOutput
	ToLogSyslogd4FilterMapOutputWithContext(context.Context) LogSyslogd4FilterMapOutput
}

type LogSyslogd4FilterMap map[string]LogSyslogd4FilterInput

func (LogSyslogd4FilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd4Filter)(nil)).Elem()
}

func (i LogSyslogd4FilterMap) ToLogSyslogd4FilterMapOutput() LogSyslogd4FilterMapOutput {
	return i.ToLogSyslogd4FilterMapOutputWithContext(context.Background())
}

func (i LogSyslogd4FilterMap) ToLogSyslogd4FilterMapOutputWithContext(ctx context.Context) LogSyslogd4FilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd4FilterMapOutput)
}

type LogSyslogd4FilterOutput struct{ *pulumi.OutputState }

func (LogSyslogd4FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd4Filter)(nil)).Elem()
}

func (o LogSyslogd4FilterOutput) ToLogSyslogd4FilterOutput() LogSyslogd4FilterOutput {
	return o
}

func (o LogSyslogd4FilterOutput) ToLogSyslogd4FilterOutputWithContext(ctx context.Context) LogSyslogd4FilterOutput {
	return o
}

type LogSyslogd4FilterArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogd4FilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd4Filter)(nil)).Elem()
}

func (o LogSyslogd4FilterArrayOutput) ToLogSyslogd4FilterArrayOutput() LogSyslogd4FilterArrayOutput {
	return o
}

func (o LogSyslogd4FilterArrayOutput) ToLogSyslogd4FilterArrayOutputWithContext(ctx context.Context) LogSyslogd4FilterArrayOutput {
	return o
}

func (o LogSyslogd4FilterArrayOutput) Index(i pulumi.IntInput) LogSyslogd4FilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogd4Filter {
		return vs[0].([]*LogSyslogd4Filter)[vs[1].(int)]
	}).(LogSyslogd4FilterOutput)
}

type LogSyslogd4FilterMapOutput struct{ *pulumi.OutputState }

func (LogSyslogd4FilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd4Filter)(nil)).Elem()
}

func (o LogSyslogd4FilterMapOutput) ToLogSyslogd4FilterMapOutput() LogSyslogd4FilterMapOutput {
	return o
}

func (o LogSyslogd4FilterMapOutput) ToLogSyslogd4FilterMapOutputWithContext(ctx context.Context) LogSyslogd4FilterMapOutput {
	return o
}

func (o LogSyslogd4FilterMapOutput) MapIndex(k pulumi.StringInput) LogSyslogd4FilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogd4Filter {
		return vs[0].(map[string]*LogSyslogd4Filter)[vs[1].(string)]
	}).(LogSyslogd4FilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4FilterInput)(nil)).Elem(), &LogSyslogd4Filter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4FilterArrayInput)(nil)).Elem(), LogSyslogd4FilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd4FilterMapInput)(nil)).Elem(), LogSyslogd4FilterMap{})
	pulumi.RegisterOutputType(LogSyslogd4FilterOutput{})
	pulumi.RegisterOutputType(LogSyslogd4FilterArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogd4FilterMapOutput{})
}

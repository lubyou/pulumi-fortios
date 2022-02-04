// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Filters for null device logging.
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
// 		_, err := fortios.NewLogNullDeviceFilter(ctx, "trname", &fortios.LogNullDeviceFilterArgs{
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
// LogNullDevice Filter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logNullDeviceFilter:LogNullDeviceFilter labelname LogNullDeviceFilter
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogNullDeviceFilter struct {
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
	FreeStyles LogNullDeviceFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogNullDeviceFilter registers a new resource with the given unique name, arguments, and options.
func NewLogNullDeviceFilter(ctx *pulumi.Context,
	name string, args *LogNullDeviceFilterArgs, opts ...pulumi.ResourceOption) (*LogNullDeviceFilter, error) {
	if args == nil {
		args = &LogNullDeviceFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogNullDeviceFilter
	err := ctx.RegisterResource("fortios:index/logNullDeviceFilter:LogNullDeviceFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogNullDeviceFilter gets an existing LogNullDeviceFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogNullDeviceFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogNullDeviceFilterState, opts ...pulumi.ResourceOption) (*LogNullDeviceFilter, error) {
	var resource LogNullDeviceFilter
	err := ctx.ReadResource("fortios:index/logNullDeviceFilter:LogNullDeviceFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogNullDeviceFilter resources.
type logNullDeviceFilterState struct {
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
	FreeStyles []LogNullDeviceFilterFreeStyle `pulumi:"freeStyles"`
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

type LogNullDeviceFilterState struct {
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
	FreeStyles LogNullDeviceFilterFreeStyleArrayInput
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

func (LogNullDeviceFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logNullDeviceFilterState)(nil)).Elem()
}

type logNullDeviceFilterArgs struct {
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
	FreeStyles []LogNullDeviceFilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogNullDeviceFilter resource.
type LogNullDeviceFilterArgs struct {
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
	FreeStyles LogNullDeviceFilterFreeStyleArrayInput
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

func (LogNullDeviceFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logNullDeviceFilterArgs)(nil)).Elem()
}

type LogNullDeviceFilterInput interface {
	pulumi.Input

	ToLogNullDeviceFilterOutput() LogNullDeviceFilterOutput
	ToLogNullDeviceFilterOutputWithContext(ctx context.Context) LogNullDeviceFilterOutput
}

func (*LogNullDeviceFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogNullDeviceFilter)(nil)).Elem()
}

func (i *LogNullDeviceFilter) ToLogNullDeviceFilterOutput() LogNullDeviceFilterOutput {
	return i.ToLogNullDeviceFilterOutputWithContext(context.Background())
}

func (i *LogNullDeviceFilter) ToLogNullDeviceFilterOutputWithContext(ctx context.Context) LogNullDeviceFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogNullDeviceFilterOutput)
}

// LogNullDeviceFilterArrayInput is an input type that accepts LogNullDeviceFilterArray and LogNullDeviceFilterArrayOutput values.
// You can construct a concrete instance of `LogNullDeviceFilterArrayInput` via:
//
//          LogNullDeviceFilterArray{ LogNullDeviceFilterArgs{...} }
type LogNullDeviceFilterArrayInput interface {
	pulumi.Input

	ToLogNullDeviceFilterArrayOutput() LogNullDeviceFilterArrayOutput
	ToLogNullDeviceFilterArrayOutputWithContext(context.Context) LogNullDeviceFilterArrayOutput
}

type LogNullDeviceFilterArray []LogNullDeviceFilterInput

func (LogNullDeviceFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogNullDeviceFilter)(nil)).Elem()
}

func (i LogNullDeviceFilterArray) ToLogNullDeviceFilterArrayOutput() LogNullDeviceFilterArrayOutput {
	return i.ToLogNullDeviceFilterArrayOutputWithContext(context.Background())
}

func (i LogNullDeviceFilterArray) ToLogNullDeviceFilterArrayOutputWithContext(ctx context.Context) LogNullDeviceFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogNullDeviceFilterArrayOutput)
}

// LogNullDeviceFilterMapInput is an input type that accepts LogNullDeviceFilterMap and LogNullDeviceFilterMapOutput values.
// You can construct a concrete instance of `LogNullDeviceFilterMapInput` via:
//
//          LogNullDeviceFilterMap{ "key": LogNullDeviceFilterArgs{...} }
type LogNullDeviceFilterMapInput interface {
	pulumi.Input

	ToLogNullDeviceFilterMapOutput() LogNullDeviceFilterMapOutput
	ToLogNullDeviceFilterMapOutputWithContext(context.Context) LogNullDeviceFilterMapOutput
}

type LogNullDeviceFilterMap map[string]LogNullDeviceFilterInput

func (LogNullDeviceFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogNullDeviceFilter)(nil)).Elem()
}

func (i LogNullDeviceFilterMap) ToLogNullDeviceFilterMapOutput() LogNullDeviceFilterMapOutput {
	return i.ToLogNullDeviceFilterMapOutputWithContext(context.Background())
}

func (i LogNullDeviceFilterMap) ToLogNullDeviceFilterMapOutputWithContext(ctx context.Context) LogNullDeviceFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogNullDeviceFilterMapOutput)
}

type LogNullDeviceFilterOutput struct{ *pulumi.OutputState }

func (LogNullDeviceFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogNullDeviceFilter)(nil)).Elem()
}

func (o LogNullDeviceFilterOutput) ToLogNullDeviceFilterOutput() LogNullDeviceFilterOutput {
	return o
}

func (o LogNullDeviceFilterOutput) ToLogNullDeviceFilterOutputWithContext(ctx context.Context) LogNullDeviceFilterOutput {
	return o
}

type LogNullDeviceFilterArrayOutput struct{ *pulumi.OutputState }

func (LogNullDeviceFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogNullDeviceFilter)(nil)).Elem()
}

func (o LogNullDeviceFilterArrayOutput) ToLogNullDeviceFilterArrayOutput() LogNullDeviceFilterArrayOutput {
	return o
}

func (o LogNullDeviceFilterArrayOutput) ToLogNullDeviceFilterArrayOutputWithContext(ctx context.Context) LogNullDeviceFilterArrayOutput {
	return o
}

func (o LogNullDeviceFilterArrayOutput) Index(i pulumi.IntInput) LogNullDeviceFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogNullDeviceFilter {
		return vs[0].([]*LogNullDeviceFilter)[vs[1].(int)]
	}).(LogNullDeviceFilterOutput)
}

type LogNullDeviceFilterMapOutput struct{ *pulumi.OutputState }

func (LogNullDeviceFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogNullDeviceFilter)(nil)).Elem()
}

func (o LogNullDeviceFilterMapOutput) ToLogNullDeviceFilterMapOutput() LogNullDeviceFilterMapOutput {
	return o
}

func (o LogNullDeviceFilterMapOutput) ToLogNullDeviceFilterMapOutputWithContext(ctx context.Context) LogNullDeviceFilterMapOutput {
	return o
}

func (o LogNullDeviceFilterMapOutput) MapIndex(k pulumi.StringInput) LogNullDeviceFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogNullDeviceFilter {
		return vs[0].(map[string]*LogNullDeviceFilter)[vs[1].(string)]
	}).(LogNullDeviceFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogNullDeviceFilterInput)(nil)).Elem(), &LogNullDeviceFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogNullDeviceFilterArrayInput)(nil)).Elem(), LogNullDeviceFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogNullDeviceFilterMapInput)(nil)).Elem(), LogNullDeviceFilterMap{})
	pulumi.RegisterOutputType(LogNullDeviceFilterOutput{})
	pulumi.RegisterOutputType(LogNullDeviceFilterArrayOutput{})
	pulumi.RegisterOutputType(LogNullDeviceFilterMapOutput{})
}

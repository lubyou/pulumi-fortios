// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogSyslogdFilter struct {
	pulumi.CustomResourceState

	Anomaly              pulumi.StringOutput                  `pulumi:"anomaly"`
	Dns                  pulumi.StringOutput                  `pulumi:"dns"`
	DynamicSortSubtable  pulumi.StringPtrOutput               `pulumi:"dynamicSortSubtable"`
	Filter               pulumi.StringOutput                  `pulumi:"filter"`
	FilterType           pulumi.StringOutput                  `pulumi:"filterType"`
	ForwardTraffic       pulumi.StringOutput                  `pulumi:"forwardTraffic"`
	FreeStyles           LogSyslogdFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	Gtp                  pulumi.StringOutput                  `pulumi:"gtp"`
	LocalTraffic         pulumi.StringOutput                  `pulumi:"localTraffic"`
	MulticastTraffic     pulumi.StringOutput                  `pulumi:"multicastTraffic"`
	NetscanDiscovery     pulumi.StringOutput                  `pulumi:"netscanDiscovery"`
	NetscanVulnerability pulumi.StringOutput                  `pulumi:"netscanVulnerability"`
	Severity             pulumi.StringOutput                  `pulumi:"severity"`
	SnifferTraffic       pulumi.StringOutput                  `pulumi:"snifferTraffic"`
	Ssh                  pulumi.StringOutput                  `pulumi:"ssh"`
	Vdomparam            pulumi.StringPtrOutput               `pulumi:"vdomparam"`
	Voip                 pulumi.StringOutput                  `pulumi:"voip"`
	ZtnaTraffic          pulumi.StringOutput                  `pulumi:"ztnaTraffic"`
}

// NewLogSyslogdFilter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogdFilter(ctx *pulumi.Context,
	name string, args *LogSyslogdFilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogdFilter, error) {
	if args == nil {
		args = &LogSyslogdFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogdFilter
	err := ctx.RegisterResource("fortios:index/logSyslogdFilter:LogSyslogdFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogdFilter gets an existing LogSyslogdFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogdFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogdFilterState, opts ...pulumi.ResourceOption) (*LogSyslogdFilter, error) {
	var resource LogSyslogdFilter
	err := ctx.ReadResource("fortios:index/logSyslogdFilter:LogSyslogdFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogdFilter resources.
type logSyslogdFilterState struct {
	Anomaly              *string                     `pulumi:"anomaly"`
	Dns                  *string                     `pulumi:"dns"`
	DynamicSortSubtable  *string                     `pulumi:"dynamicSortSubtable"`
	Filter               *string                     `pulumi:"filter"`
	FilterType           *string                     `pulumi:"filterType"`
	ForwardTraffic       *string                     `pulumi:"forwardTraffic"`
	FreeStyles           []LogSyslogdFilterFreeStyle `pulumi:"freeStyles"`
	Gtp                  *string                     `pulumi:"gtp"`
	LocalTraffic         *string                     `pulumi:"localTraffic"`
	MulticastTraffic     *string                     `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                     `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                     `pulumi:"netscanVulnerability"`
	Severity             *string                     `pulumi:"severity"`
	SnifferTraffic       *string                     `pulumi:"snifferTraffic"`
	Ssh                  *string                     `pulumi:"ssh"`
	Vdomparam            *string                     `pulumi:"vdomparam"`
	Voip                 *string                     `pulumi:"voip"`
	ZtnaTraffic          *string                     `pulumi:"ztnaTraffic"`
}

type LogSyslogdFilterState struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogSyslogdFilterFreeStyleArrayInput
	Gtp                  pulumi.StringPtrInput
	LocalTraffic         pulumi.StringPtrInput
	MulticastTraffic     pulumi.StringPtrInput
	NetscanDiscovery     pulumi.StringPtrInput
	NetscanVulnerability pulumi.StringPtrInput
	Severity             pulumi.StringPtrInput
	SnifferTraffic       pulumi.StringPtrInput
	Ssh                  pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
	Voip                 pulumi.StringPtrInput
	ZtnaTraffic          pulumi.StringPtrInput
}

func (LogSyslogdFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogdFilterState)(nil)).Elem()
}

type logSyslogdFilterArgs struct {
	Anomaly              *string                     `pulumi:"anomaly"`
	Dns                  *string                     `pulumi:"dns"`
	DynamicSortSubtable  *string                     `pulumi:"dynamicSortSubtable"`
	Filter               *string                     `pulumi:"filter"`
	FilterType           *string                     `pulumi:"filterType"`
	ForwardTraffic       *string                     `pulumi:"forwardTraffic"`
	FreeStyles           []LogSyslogdFilterFreeStyle `pulumi:"freeStyles"`
	Gtp                  *string                     `pulumi:"gtp"`
	LocalTraffic         *string                     `pulumi:"localTraffic"`
	MulticastTraffic     *string                     `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                     `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                     `pulumi:"netscanVulnerability"`
	Severity             *string                     `pulumi:"severity"`
	SnifferTraffic       *string                     `pulumi:"snifferTraffic"`
	Ssh                  *string                     `pulumi:"ssh"`
	Vdomparam            *string                     `pulumi:"vdomparam"`
	Voip                 *string                     `pulumi:"voip"`
	ZtnaTraffic          *string                     `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a LogSyslogdFilter resource.
type LogSyslogdFilterArgs struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogSyslogdFilterFreeStyleArrayInput
	Gtp                  pulumi.StringPtrInput
	LocalTraffic         pulumi.StringPtrInput
	MulticastTraffic     pulumi.StringPtrInput
	NetscanDiscovery     pulumi.StringPtrInput
	NetscanVulnerability pulumi.StringPtrInput
	Severity             pulumi.StringPtrInput
	SnifferTraffic       pulumi.StringPtrInput
	Ssh                  pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
	Voip                 pulumi.StringPtrInput
	ZtnaTraffic          pulumi.StringPtrInput
}

func (LogSyslogdFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogdFilterArgs)(nil)).Elem()
}

type LogSyslogdFilterInput interface {
	pulumi.Input

	ToLogSyslogdFilterOutput() LogSyslogdFilterOutput
	ToLogSyslogdFilterOutputWithContext(ctx context.Context) LogSyslogdFilterOutput
}

func (*LogSyslogdFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogdFilter)(nil)).Elem()
}

func (i *LogSyslogdFilter) ToLogSyslogdFilterOutput() LogSyslogdFilterOutput {
	return i.ToLogSyslogdFilterOutputWithContext(context.Background())
}

func (i *LogSyslogdFilter) ToLogSyslogdFilterOutputWithContext(ctx context.Context) LogSyslogdFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdFilterOutput)
}

// LogSyslogdFilterArrayInput is an input type that accepts LogSyslogdFilterArray and LogSyslogdFilterArrayOutput values.
// You can construct a concrete instance of `LogSyslogdFilterArrayInput` via:
//
//	LogSyslogdFilterArray{ LogSyslogdFilterArgs{...} }
type LogSyslogdFilterArrayInput interface {
	pulumi.Input

	ToLogSyslogdFilterArrayOutput() LogSyslogdFilterArrayOutput
	ToLogSyslogdFilterArrayOutputWithContext(context.Context) LogSyslogdFilterArrayOutput
}

type LogSyslogdFilterArray []LogSyslogdFilterInput

func (LogSyslogdFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogdFilter)(nil)).Elem()
}

func (i LogSyslogdFilterArray) ToLogSyslogdFilterArrayOutput() LogSyslogdFilterArrayOutput {
	return i.ToLogSyslogdFilterArrayOutputWithContext(context.Background())
}

func (i LogSyslogdFilterArray) ToLogSyslogdFilterArrayOutputWithContext(ctx context.Context) LogSyslogdFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdFilterArrayOutput)
}

// LogSyslogdFilterMapInput is an input type that accepts LogSyslogdFilterMap and LogSyslogdFilterMapOutput values.
// You can construct a concrete instance of `LogSyslogdFilterMapInput` via:
//
//	LogSyslogdFilterMap{ "key": LogSyslogdFilterArgs{...} }
type LogSyslogdFilterMapInput interface {
	pulumi.Input

	ToLogSyslogdFilterMapOutput() LogSyslogdFilterMapOutput
	ToLogSyslogdFilterMapOutputWithContext(context.Context) LogSyslogdFilterMapOutput
}

type LogSyslogdFilterMap map[string]LogSyslogdFilterInput

func (LogSyslogdFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogdFilter)(nil)).Elem()
}

func (i LogSyslogdFilterMap) ToLogSyslogdFilterMapOutput() LogSyslogdFilterMapOutput {
	return i.ToLogSyslogdFilterMapOutputWithContext(context.Background())
}

func (i LogSyslogdFilterMap) ToLogSyslogdFilterMapOutputWithContext(ctx context.Context) LogSyslogdFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdFilterMapOutput)
}

type LogSyslogdFilterOutput struct{ *pulumi.OutputState }

func (LogSyslogdFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogdFilter)(nil)).Elem()
}

func (o LogSyslogdFilterOutput) ToLogSyslogdFilterOutput() LogSyslogdFilterOutput {
	return o
}

func (o LogSyslogdFilterOutput) ToLogSyslogdFilterOutputWithContext(ctx context.Context) LogSyslogdFilterOutput {
	return o
}

func (o LogSyslogdFilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogSyslogdFilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) FreeStyles() LogSyslogdFilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) LogSyslogdFilterFreeStyleArrayOutput { return v.FreeStyles }).(LogSyslogdFilterFreeStyleArrayOutput)
}

func (o LogSyslogdFilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LogSyslogdFilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

func (o LogSyslogdFilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdFilter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
}

type LogSyslogdFilterArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogdFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogdFilter)(nil)).Elem()
}

func (o LogSyslogdFilterArrayOutput) ToLogSyslogdFilterArrayOutput() LogSyslogdFilterArrayOutput {
	return o
}

func (o LogSyslogdFilterArrayOutput) ToLogSyslogdFilterArrayOutputWithContext(ctx context.Context) LogSyslogdFilterArrayOutput {
	return o
}

func (o LogSyslogdFilterArrayOutput) Index(i pulumi.IntInput) LogSyslogdFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogdFilter {
		return vs[0].([]*LogSyslogdFilter)[vs[1].(int)]
	}).(LogSyslogdFilterOutput)
}

type LogSyslogdFilterMapOutput struct{ *pulumi.OutputState }

func (LogSyslogdFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogdFilter)(nil)).Elem()
}

func (o LogSyslogdFilterMapOutput) ToLogSyslogdFilterMapOutput() LogSyslogdFilterMapOutput {
	return o
}

func (o LogSyslogdFilterMapOutput) ToLogSyslogdFilterMapOutputWithContext(ctx context.Context) LogSyslogdFilterMapOutput {
	return o
}

func (o LogSyslogdFilterMapOutput) MapIndex(k pulumi.StringInput) LogSyslogdFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogdFilter {
		return vs[0].(map[string]*LogSyslogdFilter)[vs[1].(string)]
	}).(LogSyslogdFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdFilterInput)(nil)).Elem(), &LogSyslogdFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdFilterArrayInput)(nil)).Elem(), LogSyslogdFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdFilterMapInput)(nil)).Elem(), LogSyslogdFilterMap{})
	pulumi.RegisterOutputType(LogSyslogdFilterOutput{})
	pulumi.RegisterOutputType(LogSyslogdFilterArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogdFilterMapOutput{})
}

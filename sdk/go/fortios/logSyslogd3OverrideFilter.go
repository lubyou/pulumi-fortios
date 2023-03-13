// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogSyslogd3OverrideFilter struct {
	pulumi.CustomResourceState

	Anomaly              pulumi.StringOutput                           `pulumi:"anomaly"`
	Dns                  pulumi.StringOutput                           `pulumi:"dns"`
	DynamicSortSubtable  pulumi.StringPtrOutput                        `pulumi:"dynamicSortSubtable"`
	Filter               pulumi.StringOutput                           `pulumi:"filter"`
	FilterType           pulumi.StringOutput                           `pulumi:"filterType"`
	ForwardTraffic       pulumi.StringOutput                           `pulumi:"forwardTraffic"`
	FreeStyles           LogSyslogd3OverrideFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	Gtp                  pulumi.StringOutput                           `pulumi:"gtp"`
	LocalTraffic         pulumi.StringOutput                           `pulumi:"localTraffic"`
	MulticastTraffic     pulumi.StringOutput                           `pulumi:"multicastTraffic"`
	NetscanDiscovery     pulumi.StringOutput                           `pulumi:"netscanDiscovery"`
	NetscanVulnerability pulumi.StringOutput                           `pulumi:"netscanVulnerability"`
	Severity             pulumi.StringOutput                           `pulumi:"severity"`
	SnifferTraffic       pulumi.StringOutput                           `pulumi:"snifferTraffic"`
	Ssh                  pulumi.StringOutput                           `pulumi:"ssh"`
	Vdomparam            pulumi.StringPtrOutput                        `pulumi:"vdomparam"`
	Voip                 pulumi.StringOutput                           `pulumi:"voip"`
	ZtnaTraffic          pulumi.StringOutput                           `pulumi:"ztnaTraffic"`
}

// NewLogSyslogd3OverrideFilter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd3OverrideFilter(ctx *pulumi.Context,
	name string, args *LogSyslogd3OverrideFilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogd3OverrideFilter, error) {
	if args == nil {
		args = &LogSyslogd3OverrideFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSyslogd3OverrideFilter
	err := ctx.RegisterResource("fortios:index/logSyslogd3OverrideFilter:LogSyslogd3OverrideFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogd3OverrideFilter gets an existing LogSyslogd3OverrideFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogd3OverrideFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogd3OverrideFilterState, opts ...pulumi.ResourceOption) (*LogSyslogd3OverrideFilter, error) {
	var resource LogSyslogd3OverrideFilter
	err := ctx.ReadResource("fortios:index/logSyslogd3OverrideFilter:LogSyslogd3OverrideFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogd3OverrideFilter resources.
type logSyslogd3OverrideFilterState struct {
	Anomaly              *string                              `pulumi:"anomaly"`
	Dns                  *string                              `pulumi:"dns"`
	DynamicSortSubtable  *string                              `pulumi:"dynamicSortSubtable"`
	Filter               *string                              `pulumi:"filter"`
	FilterType           *string                              `pulumi:"filterType"`
	ForwardTraffic       *string                              `pulumi:"forwardTraffic"`
	FreeStyles           []LogSyslogd3OverrideFilterFreeStyle `pulumi:"freeStyles"`
	Gtp                  *string                              `pulumi:"gtp"`
	LocalTraffic         *string                              `pulumi:"localTraffic"`
	MulticastTraffic     *string                              `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                              `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                              `pulumi:"netscanVulnerability"`
	Severity             *string                              `pulumi:"severity"`
	SnifferTraffic       *string                              `pulumi:"snifferTraffic"`
	Ssh                  *string                              `pulumi:"ssh"`
	Vdomparam            *string                              `pulumi:"vdomparam"`
	Voip                 *string                              `pulumi:"voip"`
	ZtnaTraffic          *string                              `pulumi:"ztnaTraffic"`
}

type LogSyslogd3OverrideFilterState struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogSyslogd3OverrideFilterFreeStyleArrayInput
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

func (LogSyslogd3OverrideFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd3OverrideFilterState)(nil)).Elem()
}

type logSyslogd3OverrideFilterArgs struct {
	Anomaly              *string                              `pulumi:"anomaly"`
	Dns                  *string                              `pulumi:"dns"`
	DynamicSortSubtable  *string                              `pulumi:"dynamicSortSubtable"`
	Filter               *string                              `pulumi:"filter"`
	FilterType           *string                              `pulumi:"filterType"`
	ForwardTraffic       *string                              `pulumi:"forwardTraffic"`
	FreeStyles           []LogSyslogd3OverrideFilterFreeStyle `pulumi:"freeStyles"`
	Gtp                  *string                              `pulumi:"gtp"`
	LocalTraffic         *string                              `pulumi:"localTraffic"`
	MulticastTraffic     *string                              `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                              `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                              `pulumi:"netscanVulnerability"`
	Severity             *string                              `pulumi:"severity"`
	SnifferTraffic       *string                              `pulumi:"snifferTraffic"`
	Ssh                  *string                              `pulumi:"ssh"`
	Vdomparam            *string                              `pulumi:"vdomparam"`
	Voip                 *string                              `pulumi:"voip"`
	ZtnaTraffic          *string                              `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a LogSyslogd3OverrideFilter resource.
type LogSyslogd3OverrideFilterArgs struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogSyslogd3OverrideFilterFreeStyleArrayInput
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

func (LogSyslogd3OverrideFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd3OverrideFilterArgs)(nil)).Elem()
}

type LogSyslogd3OverrideFilterInput interface {
	pulumi.Input

	ToLogSyslogd3OverrideFilterOutput() LogSyslogd3OverrideFilterOutput
	ToLogSyslogd3OverrideFilterOutputWithContext(ctx context.Context) LogSyslogd3OverrideFilterOutput
}

func (*LogSyslogd3OverrideFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd3OverrideFilter)(nil)).Elem()
}

func (i *LogSyslogd3OverrideFilter) ToLogSyslogd3OverrideFilterOutput() LogSyslogd3OverrideFilterOutput {
	return i.ToLogSyslogd3OverrideFilterOutputWithContext(context.Background())
}

func (i *LogSyslogd3OverrideFilter) ToLogSyslogd3OverrideFilterOutputWithContext(ctx context.Context) LogSyslogd3OverrideFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd3OverrideFilterOutput)
}

// LogSyslogd3OverrideFilterArrayInput is an input type that accepts LogSyslogd3OverrideFilterArray and LogSyslogd3OverrideFilterArrayOutput values.
// You can construct a concrete instance of `LogSyslogd3OverrideFilterArrayInput` via:
//
//	LogSyslogd3OverrideFilterArray{ LogSyslogd3OverrideFilterArgs{...} }
type LogSyslogd3OverrideFilterArrayInput interface {
	pulumi.Input

	ToLogSyslogd3OverrideFilterArrayOutput() LogSyslogd3OverrideFilterArrayOutput
	ToLogSyslogd3OverrideFilterArrayOutputWithContext(context.Context) LogSyslogd3OverrideFilterArrayOutput
}

type LogSyslogd3OverrideFilterArray []LogSyslogd3OverrideFilterInput

func (LogSyslogd3OverrideFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd3OverrideFilter)(nil)).Elem()
}

func (i LogSyslogd3OverrideFilterArray) ToLogSyslogd3OverrideFilterArrayOutput() LogSyslogd3OverrideFilterArrayOutput {
	return i.ToLogSyslogd3OverrideFilterArrayOutputWithContext(context.Background())
}

func (i LogSyslogd3OverrideFilterArray) ToLogSyslogd3OverrideFilterArrayOutputWithContext(ctx context.Context) LogSyslogd3OverrideFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd3OverrideFilterArrayOutput)
}

// LogSyslogd3OverrideFilterMapInput is an input type that accepts LogSyslogd3OverrideFilterMap and LogSyslogd3OverrideFilterMapOutput values.
// You can construct a concrete instance of `LogSyslogd3OverrideFilterMapInput` via:
//
//	LogSyslogd3OverrideFilterMap{ "key": LogSyslogd3OverrideFilterArgs{...} }
type LogSyslogd3OverrideFilterMapInput interface {
	pulumi.Input

	ToLogSyslogd3OverrideFilterMapOutput() LogSyslogd3OverrideFilterMapOutput
	ToLogSyslogd3OverrideFilterMapOutputWithContext(context.Context) LogSyslogd3OverrideFilterMapOutput
}

type LogSyslogd3OverrideFilterMap map[string]LogSyslogd3OverrideFilterInput

func (LogSyslogd3OverrideFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd3OverrideFilter)(nil)).Elem()
}

func (i LogSyslogd3OverrideFilterMap) ToLogSyslogd3OverrideFilterMapOutput() LogSyslogd3OverrideFilterMapOutput {
	return i.ToLogSyslogd3OverrideFilterMapOutputWithContext(context.Background())
}

func (i LogSyslogd3OverrideFilterMap) ToLogSyslogd3OverrideFilterMapOutputWithContext(ctx context.Context) LogSyslogd3OverrideFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd3OverrideFilterMapOutput)
}

type LogSyslogd3OverrideFilterOutput struct{ *pulumi.OutputState }

func (LogSyslogd3OverrideFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd3OverrideFilter)(nil)).Elem()
}

func (o LogSyslogd3OverrideFilterOutput) ToLogSyslogd3OverrideFilterOutput() LogSyslogd3OverrideFilterOutput {
	return o
}

func (o LogSyslogd3OverrideFilterOutput) ToLogSyslogd3OverrideFilterOutputWithContext(ctx context.Context) LogSyslogd3OverrideFilterOutput {
	return o
}

func (o LogSyslogd3OverrideFilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogSyslogd3OverrideFilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) FreeStyles() LogSyslogd3OverrideFilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) LogSyslogd3OverrideFilterFreeStyleArrayOutput { return v.FreeStyles }).(LogSyslogd3OverrideFilterFreeStyleArrayOutput)
}

func (o LogSyslogd3OverrideFilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LogSyslogd3OverrideFilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

func (o LogSyslogd3OverrideFilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd3OverrideFilter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
}

type LogSyslogd3OverrideFilterArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogd3OverrideFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd3OverrideFilter)(nil)).Elem()
}

func (o LogSyslogd3OverrideFilterArrayOutput) ToLogSyslogd3OverrideFilterArrayOutput() LogSyslogd3OverrideFilterArrayOutput {
	return o
}

func (o LogSyslogd3OverrideFilterArrayOutput) ToLogSyslogd3OverrideFilterArrayOutputWithContext(ctx context.Context) LogSyslogd3OverrideFilterArrayOutput {
	return o
}

func (o LogSyslogd3OverrideFilterArrayOutput) Index(i pulumi.IntInput) LogSyslogd3OverrideFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogd3OverrideFilter {
		return vs[0].([]*LogSyslogd3OverrideFilter)[vs[1].(int)]
	}).(LogSyslogd3OverrideFilterOutput)
}

type LogSyslogd3OverrideFilterMapOutput struct{ *pulumi.OutputState }

func (LogSyslogd3OverrideFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd3OverrideFilter)(nil)).Elem()
}

func (o LogSyslogd3OverrideFilterMapOutput) ToLogSyslogd3OverrideFilterMapOutput() LogSyslogd3OverrideFilterMapOutput {
	return o
}

func (o LogSyslogd3OverrideFilterMapOutput) ToLogSyslogd3OverrideFilterMapOutputWithContext(ctx context.Context) LogSyslogd3OverrideFilterMapOutput {
	return o
}

func (o LogSyslogd3OverrideFilterMapOutput) MapIndex(k pulumi.StringInput) LogSyslogd3OverrideFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogd3OverrideFilter {
		return vs[0].(map[string]*LogSyslogd3OverrideFilter)[vs[1].(string)]
	}).(LogSyslogd3OverrideFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd3OverrideFilterInput)(nil)).Elem(), &LogSyslogd3OverrideFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd3OverrideFilterArrayInput)(nil)).Elem(), LogSyslogd3OverrideFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd3OverrideFilterMapInput)(nil)).Elem(), LogSyslogd3OverrideFilterMap{})
	pulumi.RegisterOutputType(LogSyslogd3OverrideFilterOutput{})
	pulumi.RegisterOutputType(LogSyslogd3OverrideFilterArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogd3OverrideFilterMapOutput{})
}

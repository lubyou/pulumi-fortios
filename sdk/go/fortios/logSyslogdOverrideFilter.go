// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type LogSyslogdOverrideFilter struct {
	pulumi.CustomResourceState

	Anomaly              pulumi.StringOutput                          `pulumi:"anomaly"`
	Dns                  pulumi.StringOutput                          `pulumi:"dns"`
	DynamicSortSubtable  pulumi.StringPtrOutput                       `pulumi:"dynamicSortSubtable"`
	Filter               pulumi.StringOutput                          `pulumi:"filter"`
	FilterType           pulumi.StringOutput                          `pulumi:"filterType"`
	ForwardTraffic       pulumi.StringOutput                          `pulumi:"forwardTraffic"`
	FreeStyles           LogSyslogdOverrideFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	GetAllTables         pulumi.StringPtrOutput                       `pulumi:"getAllTables"`
	Gtp                  pulumi.StringOutput                          `pulumi:"gtp"`
	LocalTraffic         pulumi.StringOutput                          `pulumi:"localTraffic"`
	MulticastTraffic     pulumi.StringOutput                          `pulumi:"multicastTraffic"`
	NetscanDiscovery     pulumi.StringOutput                          `pulumi:"netscanDiscovery"`
	NetscanVulnerability pulumi.StringOutput                          `pulumi:"netscanVulnerability"`
	Severity             pulumi.StringOutput                          `pulumi:"severity"`
	SnifferTraffic       pulumi.StringOutput                          `pulumi:"snifferTraffic"`
	Ssh                  pulumi.StringOutput                          `pulumi:"ssh"`
	Vdomparam            pulumi.StringPtrOutput                       `pulumi:"vdomparam"`
	Voip                 pulumi.StringOutput                          `pulumi:"voip"`
	ZtnaTraffic          pulumi.StringOutput                          `pulumi:"ztnaTraffic"`
}

// NewLogSyslogdOverrideFilter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogdOverrideFilter(ctx *pulumi.Context,
	name string, args *LogSyslogdOverrideFilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogdOverrideFilter, error) {
	if args == nil {
		args = &LogSyslogdOverrideFilterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogSyslogdOverrideFilter
	err := ctx.RegisterResource("fortios:index/logSyslogdOverrideFilter:LogSyslogdOverrideFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogdOverrideFilter gets an existing LogSyslogdOverrideFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogdOverrideFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogdOverrideFilterState, opts ...pulumi.ResourceOption) (*LogSyslogdOverrideFilter, error) {
	var resource LogSyslogdOverrideFilter
	err := ctx.ReadResource("fortios:index/logSyslogdOverrideFilter:LogSyslogdOverrideFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogdOverrideFilter resources.
type logSyslogdOverrideFilterState struct {
	Anomaly              *string                             `pulumi:"anomaly"`
	Dns                  *string                             `pulumi:"dns"`
	DynamicSortSubtable  *string                             `pulumi:"dynamicSortSubtable"`
	Filter               *string                             `pulumi:"filter"`
	FilterType           *string                             `pulumi:"filterType"`
	ForwardTraffic       *string                             `pulumi:"forwardTraffic"`
	FreeStyles           []LogSyslogdOverrideFilterFreeStyle `pulumi:"freeStyles"`
	GetAllTables         *string                             `pulumi:"getAllTables"`
	Gtp                  *string                             `pulumi:"gtp"`
	LocalTraffic         *string                             `pulumi:"localTraffic"`
	MulticastTraffic     *string                             `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                             `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                             `pulumi:"netscanVulnerability"`
	Severity             *string                             `pulumi:"severity"`
	SnifferTraffic       *string                             `pulumi:"snifferTraffic"`
	Ssh                  *string                             `pulumi:"ssh"`
	Vdomparam            *string                             `pulumi:"vdomparam"`
	Voip                 *string                             `pulumi:"voip"`
	ZtnaTraffic          *string                             `pulumi:"ztnaTraffic"`
}

type LogSyslogdOverrideFilterState struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogSyslogdOverrideFilterFreeStyleArrayInput
	GetAllTables         pulumi.StringPtrInput
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

func (LogSyslogdOverrideFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogdOverrideFilterState)(nil)).Elem()
}

type logSyslogdOverrideFilterArgs struct {
	Anomaly              *string                             `pulumi:"anomaly"`
	Dns                  *string                             `pulumi:"dns"`
	DynamicSortSubtable  *string                             `pulumi:"dynamicSortSubtable"`
	Filter               *string                             `pulumi:"filter"`
	FilterType           *string                             `pulumi:"filterType"`
	ForwardTraffic       *string                             `pulumi:"forwardTraffic"`
	FreeStyles           []LogSyslogdOverrideFilterFreeStyle `pulumi:"freeStyles"`
	GetAllTables         *string                             `pulumi:"getAllTables"`
	Gtp                  *string                             `pulumi:"gtp"`
	LocalTraffic         *string                             `pulumi:"localTraffic"`
	MulticastTraffic     *string                             `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                             `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                             `pulumi:"netscanVulnerability"`
	Severity             *string                             `pulumi:"severity"`
	SnifferTraffic       *string                             `pulumi:"snifferTraffic"`
	Ssh                  *string                             `pulumi:"ssh"`
	Vdomparam            *string                             `pulumi:"vdomparam"`
	Voip                 *string                             `pulumi:"voip"`
	ZtnaTraffic          *string                             `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a LogSyslogdOverrideFilter resource.
type LogSyslogdOverrideFilterArgs struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogSyslogdOverrideFilterFreeStyleArrayInput
	GetAllTables         pulumi.StringPtrInput
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

func (LogSyslogdOverrideFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogdOverrideFilterArgs)(nil)).Elem()
}

type LogSyslogdOverrideFilterInput interface {
	pulumi.Input

	ToLogSyslogdOverrideFilterOutput() LogSyslogdOverrideFilterOutput
	ToLogSyslogdOverrideFilterOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterOutput
}

func (*LogSyslogdOverrideFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogdOverrideFilter)(nil)).Elem()
}

func (i *LogSyslogdOverrideFilter) ToLogSyslogdOverrideFilterOutput() LogSyslogdOverrideFilterOutput {
	return i.ToLogSyslogdOverrideFilterOutputWithContext(context.Background())
}

func (i *LogSyslogdOverrideFilter) ToLogSyslogdOverrideFilterOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdOverrideFilterOutput)
}

func (i *LogSyslogdOverrideFilter) ToOutput(ctx context.Context) pulumix.Output[*LogSyslogdOverrideFilter] {
	return pulumix.Output[*LogSyslogdOverrideFilter]{
		OutputState: i.ToLogSyslogdOverrideFilterOutputWithContext(ctx).OutputState,
	}
}

// LogSyslogdOverrideFilterArrayInput is an input type that accepts LogSyslogdOverrideFilterArray and LogSyslogdOverrideFilterArrayOutput values.
// You can construct a concrete instance of `LogSyslogdOverrideFilterArrayInput` via:
//
//	LogSyslogdOverrideFilterArray{ LogSyslogdOverrideFilterArgs{...} }
type LogSyslogdOverrideFilterArrayInput interface {
	pulumi.Input

	ToLogSyslogdOverrideFilterArrayOutput() LogSyslogdOverrideFilterArrayOutput
	ToLogSyslogdOverrideFilterArrayOutputWithContext(context.Context) LogSyslogdOverrideFilterArrayOutput
}

type LogSyslogdOverrideFilterArray []LogSyslogdOverrideFilterInput

func (LogSyslogdOverrideFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogdOverrideFilter)(nil)).Elem()
}

func (i LogSyslogdOverrideFilterArray) ToLogSyslogdOverrideFilterArrayOutput() LogSyslogdOverrideFilterArrayOutput {
	return i.ToLogSyslogdOverrideFilterArrayOutputWithContext(context.Background())
}

func (i LogSyslogdOverrideFilterArray) ToLogSyslogdOverrideFilterArrayOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdOverrideFilterArrayOutput)
}

func (i LogSyslogdOverrideFilterArray) ToOutput(ctx context.Context) pulumix.Output[[]*LogSyslogdOverrideFilter] {
	return pulumix.Output[[]*LogSyslogdOverrideFilter]{
		OutputState: i.ToLogSyslogdOverrideFilterArrayOutputWithContext(ctx).OutputState,
	}
}

// LogSyslogdOverrideFilterMapInput is an input type that accepts LogSyslogdOverrideFilterMap and LogSyslogdOverrideFilterMapOutput values.
// You can construct a concrete instance of `LogSyslogdOverrideFilterMapInput` via:
//
//	LogSyslogdOverrideFilterMap{ "key": LogSyslogdOverrideFilterArgs{...} }
type LogSyslogdOverrideFilterMapInput interface {
	pulumi.Input

	ToLogSyslogdOverrideFilterMapOutput() LogSyslogdOverrideFilterMapOutput
	ToLogSyslogdOverrideFilterMapOutputWithContext(context.Context) LogSyslogdOverrideFilterMapOutput
}

type LogSyslogdOverrideFilterMap map[string]LogSyslogdOverrideFilterInput

func (LogSyslogdOverrideFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogdOverrideFilter)(nil)).Elem()
}

func (i LogSyslogdOverrideFilterMap) ToLogSyslogdOverrideFilterMapOutput() LogSyslogdOverrideFilterMapOutput {
	return i.ToLogSyslogdOverrideFilterMapOutputWithContext(context.Background())
}

func (i LogSyslogdOverrideFilterMap) ToLogSyslogdOverrideFilterMapOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdOverrideFilterMapOutput)
}

func (i LogSyslogdOverrideFilterMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LogSyslogdOverrideFilter] {
	return pulumix.Output[map[string]*LogSyslogdOverrideFilter]{
		OutputState: i.ToLogSyslogdOverrideFilterMapOutputWithContext(ctx).OutputState,
	}
}

type LogSyslogdOverrideFilterOutput struct{ *pulumi.OutputState }

func (LogSyslogdOverrideFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogdOverrideFilter)(nil)).Elem()
}

func (o LogSyslogdOverrideFilterOutput) ToLogSyslogdOverrideFilterOutput() LogSyslogdOverrideFilterOutput {
	return o
}

func (o LogSyslogdOverrideFilterOutput) ToLogSyslogdOverrideFilterOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterOutput {
	return o
}

func (o LogSyslogdOverrideFilterOutput) ToOutput(ctx context.Context) pulumix.Output[*LogSyslogdOverrideFilter] {
	return pulumix.Output[*LogSyslogdOverrideFilter]{
		OutputState: o.OutputState,
	}
}

func (o LogSyslogdOverrideFilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogSyslogdOverrideFilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) FreeStyles() LogSyslogdOverrideFilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) LogSyslogdOverrideFilterFreeStyleArrayOutput { return v.FreeStyles }).(LogSyslogdOverrideFilterFreeStyleArrayOutput)
}

func (o LogSyslogdOverrideFilterOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o LogSyslogdOverrideFilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LogSyslogdOverrideFilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

func (o LogSyslogdOverrideFilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdOverrideFilter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
}

type LogSyslogdOverrideFilterArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogdOverrideFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogdOverrideFilter)(nil)).Elem()
}

func (o LogSyslogdOverrideFilterArrayOutput) ToLogSyslogdOverrideFilterArrayOutput() LogSyslogdOverrideFilterArrayOutput {
	return o
}

func (o LogSyslogdOverrideFilterArrayOutput) ToLogSyslogdOverrideFilterArrayOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterArrayOutput {
	return o
}

func (o LogSyslogdOverrideFilterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LogSyslogdOverrideFilter] {
	return pulumix.Output[[]*LogSyslogdOverrideFilter]{
		OutputState: o.OutputState,
	}
}

func (o LogSyslogdOverrideFilterArrayOutput) Index(i pulumi.IntInput) LogSyslogdOverrideFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogdOverrideFilter {
		return vs[0].([]*LogSyslogdOverrideFilter)[vs[1].(int)]
	}).(LogSyslogdOverrideFilterOutput)
}

type LogSyslogdOverrideFilterMapOutput struct{ *pulumi.OutputState }

func (LogSyslogdOverrideFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogdOverrideFilter)(nil)).Elem()
}

func (o LogSyslogdOverrideFilterMapOutput) ToLogSyslogdOverrideFilterMapOutput() LogSyslogdOverrideFilterMapOutput {
	return o
}

func (o LogSyslogdOverrideFilterMapOutput) ToLogSyslogdOverrideFilterMapOutputWithContext(ctx context.Context) LogSyslogdOverrideFilterMapOutput {
	return o
}

func (o LogSyslogdOverrideFilterMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LogSyslogdOverrideFilter] {
	return pulumix.Output[map[string]*LogSyslogdOverrideFilter]{
		OutputState: o.OutputState,
	}
}

func (o LogSyslogdOverrideFilterMapOutput) MapIndex(k pulumi.StringInput) LogSyslogdOverrideFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogdOverrideFilter {
		return vs[0].(map[string]*LogSyslogdOverrideFilter)[vs[1].(string)]
	}).(LogSyslogdOverrideFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdOverrideFilterInput)(nil)).Elem(), &LogSyslogdOverrideFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdOverrideFilterArrayInput)(nil)).Elem(), LogSyslogdOverrideFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdOverrideFilterMapInput)(nil)).Elem(), LogSyslogdOverrideFilterMap{})
	pulumi.RegisterOutputType(LogSyslogdOverrideFilterOutput{})
	pulumi.RegisterOutputType(LogSyslogdOverrideFilterArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogdOverrideFilterMapOutput{})
}

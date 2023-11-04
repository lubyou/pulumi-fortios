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

type LogSyslogd2Filter struct {
	pulumi.CustomResourceState

	Anomaly              pulumi.StringOutput                   `pulumi:"anomaly"`
	Dns                  pulumi.StringOutput                   `pulumi:"dns"`
	DynamicSortSubtable  pulumi.StringPtrOutput                `pulumi:"dynamicSortSubtable"`
	Filter               pulumi.StringOutput                   `pulumi:"filter"`
	FilterType           pulumi.StringOutput                   `pulumi:"filterType"`
	ForwardTraffic       pulumi.StringOutput                   `pulumi:"forwardTraffic"`
	FreeStyles           LogSyslogd2FilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	GetAllTables         pulumi.StringPtrOutput                `pulumi:"getAllTables"`
	Gtp                  pulumi.StringOutput                   `pulumi:"gtp"`
	LocalTraffic         pulumi.StringOutput                   `pulumi:"localTraffic"`
	MulticastTraffic     pulumi.StringOutput                   `pulumi:"multicastTraffic"`
	NetscanDiscovery     pulumi.StringOutput                   `pulumi:"netscanDiscovery"`
	NetscanVulnerability pulumi.StringOutput                   `pulumi:"netscanVulnerability"`
	Severity             pulumi.StringOutput                   `pulumi:"severity"`
	SnifferTraffic       pulumi.StringOutput                   `pulumi:"snifferTraffic"`
	Ssh                  pulumi.StringOutput                   `pulumi:"ssh"`
	Vdomparam            pulumi.StringPtrOutput                `pulumi:"vdomparam"`
	Voip                 pulumi.StringOutput                   `pulumi:"voip"`
	ZtnaTraffic          pulumi.StringOutput                   `pulumi:"ztnaTraffic"`
}

// NewLogSyslogd2Filter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd2Filter(ctx *pulumi.Context,
	name string, args *LogSyslogd2FilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogd2Filter, error) {
	if args == nil {
		args = &LogSyslogd2FilterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
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
	Anomaly              *string                      `pulumi:"anomaly"`
	Dns                  *string                      `pulumi:"dns"`
	DynamicSortSubtable  *string                      `pulumi:"dynamicSortSubtable"`
	Filter               *string                      `pulumi:"filter"`
	FilterType           *string                      `pulumi:"filterType"`
	ForwardTraffic       *string                      `pulumi:"forwardTraffic"`
	FreeStyles           []LogSyslogd2FilterFreeStyle `pulumi:"freeStyles"`
	GetAllTables         *string                      `pulumi:"getAllTables"`
	Gtp                  *string                      `pulumi:"gtp"`
	LocalTraffic         *string                      `pulumi:"localTraffic"`
	MulticastTraffic     *string                      `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                      `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                      `pulumi:"netscanVulnerability"`
	Severity             *string                      `pulumi:"severity"`
	SnifferTraffic       *string                      `pulumi:"snifferTraffic"`
	Ssh                  *string                      `pulumi:"ssh"`
	Vdomparam            *string                      `pulumi:"vdomparam"`
	Voip                 *string                      `pulumi:"voip"`
	ZtnaTraffic          *string                      `pulumi:"ztnaTraffic"`
}

type LogSyslogd2FilterState struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogSyslogd2FilterFreeStyleArrayInput
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

func (LogSyslogd2FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd2FilterState)(nil)).Elem()
}

type logSyslogd2FilterArgs struct {
	Anomaly              *string                      `pulumi:"anomaly"`
	Dns                  *string                      `pulumi:"dns"`
	DynamicSortSubtable  *string                      `pulumi:"dynamicSortSubtable"`
	Filter               *string                      `pulumi:"filter"`
	FilterType           *string                      `pulumi:"filterType"`
	ForwardTraffic       *string                      `pulumi:"forwardTraffic"`
	FreeStyles           []LogSyslogd2FilterFreeStyle `pulumi:"freeStyles"`
	GetAllTables         *string                      `pulumi:"getAllTables"`
	Gtp                  *string                      `pulumi:"gtp"`
	LocalTraffic         *string                      `pulumi:"localTraffic"`
	MulticastTraffic     *string                      `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                      `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                      `pulumi:"netscanVulnerability"`
	Severity             *string                      `pulumi:"severity"`
	SnifferTraffic       *string                      `pulumi:"snifferTraffic"`
	Ssh                  *string                      `pulumi:"ssh"`
	Vdomparam            *string                      `pulumi:"vdomparam"`
	Voip                 *string                      `pulumi:"voip"`
	ZtnaTraffic          *string                      `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a LogSyslogd2Filter resource.
type LogSyslogd2FilterArgs struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogSyslogd2FilterFreeStyleArrayInput
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

func (LogSyslogd2FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd2FilterArgs)(nil)).Elem()
}

type LogSyslogd2FilterInput interface {
	pulumi.Input

	ToLogSyslogd2FilterOutput() LogSyslogd2FilterOutput
	ToLogSyslogd2FilterOutputWithContext(ctx context.Context) LogSyslogd2FilterOutput
}

func (*LogSyslogd2Filter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd2Filter)(nil)).Elem()
}

func (i *LogSyslogd2Filter) ToLogSyslogd2FilterOutput() LogSyslogd2FilterOutput {
	return i.ToLogSyslogd2FilterOutputWithContext(context.Background())
}

func (i *LogSyslogd2Filter) ToLogSyslogd2FilterOutputWithContext(ctx context.Context) LogSyslogd2FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2FilterOutput)
}

func (i *LogSyslogd2Filter) ToOutput(ctx context.Context) pulumix.Output[*LogSyslogd2Filter] {
	return pulumix.Output[*LogSyslogd2Filter]{
		OutputState: i.ToLogSyslogd2FilterOutputWithContext(ctx).OutputState,
	}
}

// LogSyslogd2FilterArrayInput is an input type that accepts LogSyslogd2FilterArray and LogSyslogd2FilterArrayOutput values.
// You can construct a concrete instance of `LogSyslogd2FilterArrayInput` via:
//
//	LogSyslogd2FilterArray{ LogSyslogd2FilterArgs{...} }
type LogSyslogd2FilterArrayInput interface {
	pulumi.Input

	ToLogSyslogd2FilterArrayOutput() LogSyslogd2FilterArrayOutput
	ToLogSyslogd2FilterArrayOutputWithContext(context.Context) LogSyslogd2FilterArrayOutput
}

type LogSyslogd2FilterArray []LogSyslogd2FilterInput

func (LogSyslogd2FilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd2Filter)(nil)).Elem()
}

func (i LogSyslogd2FilterArray) ToLogSyslogd2FilterArrayOutput() LogSyslogd2FilterArrayOutput {
	return i.ToLogSyslogd2FilterArrayOutputWithContext(context.Background())
}

func (i LogSyslogd2FilterArray) ToLogSyslogd2FilterArrayOutputWithContext(ctx context.Context) LogSyslogd2FilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2FilterArrayOutput)
}

func (i LogSyslogd2FilterArray) ToOutput(ctx context.Context) pulumix.Output[[]*LogSyslogd2Filter] {
	return pulumix.Output[[]*LogSyslogd2Filter]{
		OutputState: i.ToLogSyslogd2FilterArrayOutputWithContext(ctx).OutputState,
	}
}

// LogSyslogd2FilterMapInput is an input type that accepts LogSyslogd2FilterMap and LogSyslogd2FilterMapOutput values.
// You can construct a concrete instance of `LogSyslogd2FilterMapInput` via:
//
//	LogSyslogd2FilterMap{ "key": LogSyslogd2FilterArgs{...} }
type LogSyslogd2FilterMapInput interface {
	pulumi.Input

	ToLogSyslogd2FilterMapOutput() LogSyslogd2FilterMapOutput
	ToLogSyslogd2FilterMapOutputWithContext(context.Context) LogSyslogd2FilterMapOutput
}

type LogSyslogd2FilterMap map[string]LogSyslogd2FilterInput

func (LogSyslogd2FilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd2Filter)(nil)).Elem()
}

func (i LogSyslogd2FilterMap) ToLogSyslogd2FilterMapOutput() LogSyslogd2FilterMapOutput {
	return i.ToLogSyslogd2FilterMapOutputWithContext(context.Background())
}

func (i LogSyslogd2FilterMap) ToLogSyslogd2FilterMapOutputWithContext(ctx context.Context) LogSyslogd2FilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogd2FilterMapOutput)
}

func (i LogSyslogd2FilterMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LogSyslogd2Filter] {
	return pulumix.Output[map[string]*LogSyslogd2Filter]{
		OutputState: i.ToLogSyslogd2FilterMapOutputWithContext(ctx).OutputState,
	}
}

type LogSyslogd2FilterOutput struct{ *pulumi.OutputState }

func (LogSyslogd2FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogd2Filter)(nil)).Elem()
}

func (o LogSyslogd2FilterOutput) ToLogSyslogd2FilterOutput() LogSyslogd2FilterOutput {
	return o
}

func (o LogSyslogd2FilterOutput) ToLogSyslogd2FilterOutputWithContext(ctx context.Context) LogSyslogd2FilterOutput {
	return o
}

func (o LogSyslogd2FilterOutput) ToOutput(ctx context.Context) pulumix.Output[*LogSyslogd2Filter] {
	return pulumix.Output[*LogSyslogd2Filter]{
		OutputState: o.OutputState,
	}
}

func (o LogSyslogd2FilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogSyslogd2FilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) FreeStyles() LogSyslogd2FilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) LogSyslogd2FilterFreeStyleArrayOutput { return v.FreeStyles }).(LogSyslogd2FilterFreeStyleArrayOutput)
}

func (o LogSyslogd2FilterOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o LogSyslogd2FilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LogSyslogd2FilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

func (o LogSyslogd2FilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd2Filter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
}

type LogSyslogd2FilterArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogd2FilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogd2Filter)(nil)).Elem()
}

func (o LogSyslogd2FilterArrayOutput) ToLogSyslogd2FilterArrayOutput() LogSyslogd2FilterArrayOutput {
	return o
}

func (o LogSyslogd2FilterArrayOutput) ToLogSyslogd2FilterArrayOutputWithContext(ctx context.Context) LogSyslogd2FilterArrayOutput {
	return o
}

func (o LogSyslogd2FilterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LogSyslogd2Filter] {
	return pulumix.Output[[]*LogSyslogd2Filter]{
		OutputState: o.OutputState,
	}
}

func (o LogSyslogd2FilterArrayOutput) Index(i pulumi.IntInput) LogSyslogd2FilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogd2Filter {
		return vs[0].([]*LogSyslogd2Filter)[vs[1].(int)]
	}).(LogSyslogd2FilterOutput)
}

type LogSyslogd2FilterMapOutput struct{ *pulumi.OutputState }

func (LogSyslogd2FilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogd2Filter)(nil)).Elem()
}

func (o LogSyslogd2FilterMapOutput) ToLogSyslogd2FilterMapOutput() LogSyslogd2FilterMapOutput {
	return o
}

func (o LogSyslogd2FilterMapOutput) ToLogSyslogd2FilterMapOutputWithContext(ctx context.Context) LogSyslogd2FilterMapOutput {
	return o
}

func (o LogSyslogd2FilterMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LogSyslogd2Filter] {
	return pulumix.Output[map[string]*LogSyslogd2Filter]{
		OutputState: o.OutputState,
	}
}

func (o LogSyslogd2FilterMapOutput) MapIndex(k pulumi.StringInput) LogSyslogd2FilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogd2Filter {
		return vs[0].(map[string]*LogSyslogd2Filter)[vs[1].(string)]
	}).(LogSyslogd2FilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd2FilterInput)(nil)).Elem(), &LogSyslogd2Filter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd2FilterArrayInput)(nil)).Elem(), LogSyslogd2FilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogd2FilterMapInput)(nil)).Elem(), LogSyslogd2FilterMap{})
	pulumi.RegisterOutputType(LogSyslogd2FilterOutput{})
	pulumi.RegisterOutputType(LogSyslogd2FilterArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogd2FilterMapOutput{})
}

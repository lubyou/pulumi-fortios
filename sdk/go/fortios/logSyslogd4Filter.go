// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogSyslogd4Filter struct {
	pulumi.CustomResourceState

	Anomaly              pulumi.StringOutput                   `pulumi:"anomaly"`
	Dns                  pulumi.StringOutput                   `pulumi:"dns"`
	DynamicSortSubtable  pulumi.StringPtrOutput                `pulumi:"dynamicSortSubtable"`
	Filter               pulumi.StringOutput                   `pulumi:"filter"`
	FilterType           pulumi.StringOutput                   `pulumi:"filterType"`
	ForwardTraffic       pulumi.StringOutput                   `pulumi:"forwardTraffic"`
	FreeStyles           LogSyslogd4FilterFreeStyleArrayOutput `pulumi:"freeStyles"`
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

// NewLogSyslogd4Filter registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogd4Filter(ctx *pulumi.Context,
	name string, args *LogSyslogd4FilterArgs, opts ...pulumi.ResourceOption) (*LogSyslogd4Filter, error) {
	if args == nil {
		args = &LogSyslogd4FilterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
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
	Anomaly              *string                      `pulumi:"anomaly"`
	Dns                  *string                      `pulumi:"dns"`
	DynamicSortSubtable  *string                      `pulumi:"dynamicSortSubtable"`
	Filter               *string                      `pulumi:"filter"`
	FilterType           *string                      `pulumi:"filterType"`
	ForwardTraffic       *string                      `pulumi:"forwardTraffic"`
	FreeStyles           []LogSyslogd4FilterFreeStyle `pulumi:"freeStyles"`
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

type LogSyslogd4FilterState struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogSyslogd4FilterFreeStyleArrayInput
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

func (LogSyslogd4FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogd4FilterState)(nil)).Elem()
}

type logSyslogd4FilterArgs struct {
	Anomaly              *string                      `pulumi:"anomaly"`
	Dns                  *string                      `pulumi:"dns"`
	DynamicSortSubtable  *string                      `pulumi:"dynamicSortSubtable"`
	Filter               *string                      `pulumi:"filter"`
	FilterType           *string                      `pulumi:"filterType"`
	ForwardTraffic       *string                      `pulumi:"forwardTraffic"`
	FreeStyles           []LogSyslogd4FilterFreeStyle `pulumi:"freeStyles"`
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

// The set of arguments for constructing a LogSyslogd4Filter resource.
type LogSyslogd4FilterArgs struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogSyslogd4FilterFreeStyleArrayInput
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
//	LogSyslogd4FilterArray{ LogSyslogd4FilterArgs{...} }
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
//	LogSyslogd4FilterMap{ "key": LogSyslogd4FilterArgs{...} }
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

func (o LogSyslogd4FilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogSyslogd4FilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) FreeStyles() LogSyslogd4FilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) LogSyslogd4FilterFreeStyleArrayOutput { return v.FreeStyles }).(LogSyslogd4FilterFreeStyleArrayOutput)
}

func (o LogSyslogd4FilterOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o LogSyslogd4FilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LogSyslogd4FilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

func (o LogSyslogd4FilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogd4Filter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
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

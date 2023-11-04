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

type LogFortianalyzerFilter struct {
	pulumi.CustomResourceState

	Anomaly              pulumi.StringOutput                        `pulumi:"anomaly"`
	DlpArchive           pulumi.StringOutput                        `pulumi:"dlpArchive"`
	Dns                  pulumi.StringOutput                        `pulumi:"dns"`
	DynamicSortSubtable  pulumi.StringPtrOutput                     `pulumi:"dynamicSortSubtable"`
	Filter               pulumi.StringOutput                        `pulumi:"filter"`
	FilterType           pulumi.StringOutput                        `pulumi:"filterType"`
	ForwardTraffic       pulumi.StringOutput                        `pulumi:"forwardTraffic"`
	FreeStyles           LogFortianalyzerFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	GetAllTables         pulumi.StringPtrOutput                     `pulumi:"getAllTables"`
	Gtp                  pulumi.StringOutput                        `pulumi:"gtp"`
	LocalTraffic         pulumi.StringOutput                        `pulumi:"localTraffic"`
	MulticastTraffic     pulumi.StringOutput                        `pulumi:"multicastTraffic"`
	NetscanDiscovery     pulumi.StringOutput                        `pulumi:"netscanDiscovery"`
	NetscanVulnerability pulumi.StringOutput                        `pulumi:"netscanVulnerability"`
	Severity             pulumi.StringOutput                        `pulumi:"severity"`
	SnifferTraffic       pulumi.StringOutput                        `pulumi:"snifferTraffic"`
	Ssh                  pulumi.StringOutput                        `pulumi:"ssh"`
	Vdomparam            pulumi.StringPtrOutput                     `pulumi:"vdomparam"`
	Voip                 pulumi.StringOutput                        `pulumi:"voip"`
	ZtnaTraffic          pulumi.StringOutput                        `pulumi:"ztnaTraffic"`
}

// NewLogFortianalyzerFilter registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzerFilter(ctx *pulumi.Context,
	name string, args *LogFortianalyzerFilterArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzerFilter, error) {
	if args == nil {
		args = &LogFortianalyzerFilterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogFortianalyzerFilter
	err := ctx.RegisterResource("fortios:index/logFortianalyzerFilter:LogFortianalyzerFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzerFilter gets an existing LogFortianalyzerFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzerFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzerFilterState, opts ...pulumi.ResourceOption) (*LogFortianalyzerFilter, error) {
	var resource LogFortianalyzerFilter
	err := ctx.ReadResource("fortios:index/logFortianalyzerFilter:LogFortianalyzerFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzerFilter resources.
type logFortianalyzerFilterState struct {
	Anomaly              *string                           `pulumi:"anomaly"`
	DlpArchive           *string                           `pulumi:"dlpArchive"`
	Dns                  *string                           `pulumi:"dns"`
	DynamicSortSubtable  *string                           `pulumi:"dynamicSortSubtable"`
	Filter               *string                           `pulumi:"filter"`
	FilterType           *string                           `pulumi:"filterType"`
	ForwardTraffic       *string                           `pulumi:"forwardTraffic"`
	FreeStyles           []LogFortianalyzerFilterFreeStyle `pulumi:"freeStyles"`
	GetAllTables         *string                           `pulumi:"getAllTables"`
	Gtp                  *string                           `pulumi:"gtp"`
	LocalTraffic         *string                           `pulumi:"localTraffic"`
	MulticastTraffic     *string                           `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                           `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                           `pulumi:"netscanVulnerability"`
	Severity             *string                           `pulumi:"severity"`
	SnifferTraffic       *string                           `pulumi:"snifferTraffic"`
	Ssh                  *string                           `pulumi:"ssh"`
	Vdomparam            *string                           `pulumi:"vdomparam"`
	Voip                 *string                           `pulumi:"voip"`
	ZtnaTraffic          *string                           `pulumi:"ztnaTraffic"`
}

type LogFortianalyzerFilterState struct {
	Anomaly              pulumi.StringPtrInput
	DlpArchive           pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogFortianalyzerFilterFreeStyleArrayInput
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

func (LogFortianalyzerFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzerFilterState)(nil)).Elem()
}

type logFortianalyzerFilterArgs struct {
	Anomaly              *string                           `pulumi:"anomaly"`
	DlpArchive           *string                           `pulumi:"dlpArchive"`
	Dns                  *string                           `pulumi:"dns"`
	DynamicSortSubtable  *string                           `pulumi:"dynamicSortSubtable"`
	Filter               *string                           `pulumi:"filter"`
	FilterType           *string                           `pulumi:"filterType"`
	ForwardTraffic       *string                           `pulumi:"forwardTraffic"`
	FreeStyles           []LogFortianalyzerFilterFreeStyle `pulumi:"freeStyles"`
	GetAllTables         *string                           `pulumi:"getAllTables"`
	Gtp                  *string                           `pulumi:"gtp"`
	LocalTraffic         *string                           `pulumi:"localTraffic"`
	MulticastTraffic     *string                           `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                           `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                           `pulumi:"netscanVulnerability"`
	Severity             *string                           `pulumi:"severity"`
	SnifferTraffic       *string                           `pulumi:"snifferTraffic"`
	Ssh                  *string                           `pulumi:"ssh"`
	Vdomparam            *string                           `pulumi:"vdomparam"`
	Voip                 *string                           `pulumi:"voip"`
	ZtnaTraffic          *string                           `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a LogFortianalyzerFilter resource.
type LogFortianalyzerFilterArgs struct {
	Anomaly              pulumi.StringPtrInput
	DlpArchive           pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogFortianalyzerFilterFreeStyleArrayInput
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

func (LogFortianalyzerFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzerFilterArgs)(nil)).Elem()
}

type LogFortianalyzerFilterInput interface {
	pulumi.Input

	ToLogFortianalyzerFilterOutput() LogFortianalyzerFilterOutput
	ToLogFortianalyzerFilterOutputWithContext(ctx context.Context) LogFortianalyzerFilterOutput
}

func (*LogFortianalyzerFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzerFilter)(nil)).Elem()
}

func (i *LogFortianalyzerFilter) ToLogFortianalyzerFilterOutput() LogFortianalyzerFilterOutput {
	return i.ToLogFortianalyzerFilterOutputWithContext(context.Background())
}

func (i *LogFortianalyzerFilter) ToLogFortianalyzerFilterOutputWithContext(ctx context.Context) LogFortianalyzerFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerFilterOutput)
}

func (i *LogFortianalyzerFilter) ToOutput(ctx context.Context) pulumix.Output[*LogFortianalyzerFilter] {
	return pulumix.Output[*LogFortianalyzerFilter]{
		OutputState: i.ToLogFortianalyzerFilterOutputWithContext(ctx).OutputState,
	}
}

// LogFortianalyzerFilterArrayInput is an input type that accepts LogFortianalyzerFilterArray and LogFortianalyzerFilterArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzerFilterArrayInput` via:
//
//	LogFortianalyzerFilterArray{ LogFortianalyzerFilterArgs{...} }
type LogFortianalyzerFilterArrayInput interface {
	pulumi.Input

	ToLogFortianalyzerFilterArrayOutput() LogFortianalyzerFilterArrayOutput
	ToLogFortianalyzerFilterArrayOutputWithContext(context.Context) LogFortianalyzerFilterArrayOutput
}

type LogFortianalyzerFilterArray []LogFortianalyzerFilterInput

func (LogFortianalyzerFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzerFilter)(nil)).Elem()
}

func (i LogFortianalyzerFilterArray) ToLogFortianalyzerFilterArrayOutput() LogFortianalyzerFilterArrayOutput {
	return i.ToLogFortianalyzerFilterArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzerFilterArray) ToLogFortianalyzerFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzerFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerFilterArrayOutput)
}

func (i LogFortianalyzerFilterArray) ToOutput(ctx context.Context) pulumix.Output[[]*LogFortianalyzerFilter] {
	return pulumix.Output[[]*LogFortianalyzerFilter]{
		OutputState: i.ToLogFortianalyzerFilterArrayOutputWithContext(ctx).OutputState,
	}
}

// LogFortianalyzerFilterMapInput is an input type that accepts LogFortianalyzerFilterMap and LogFortianalyzerFilterMapOutput values.
// You can construct a concrete instance of `LogFortianalyzerFilterMapInput` via:
//
//	LogFortianalyzerFilterMap{ "key": LogFortianalyzerFilterArgs{...} }
type LogFortianalyzerFilterMapInput interface {
	pulumi.Input

	ToLogFortianalyzerFilterMapOutput() LogFortianalyzerFilterMapOutput
	ToLogFortianalyzerFilterMapOutputWithContext(context.Context) LogFortianalyzerFilterMapOutput
}

type LogFortianalyzerFilterMap map[string]LogFortianalyzerFilterInput

func (LogFortianalyzerFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzerFilter)(nil)).Elem()
}

func (i LogFortianalyzerFilterMap) ToLogFortianalyzerFilterMapOutput() LogFortianalyzerFilterMapOutput {
	return i.ToLogFortianalyzerFilterMapOutputWithContext(context.Background())
}

func (i LogFortianalyzerFilterMap) ToLogFortianalyzerFilterMapOutputWithContext(ctx context.Context) LogFortianalyzerFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzerFilterMapOutput)
}

func (i LogFortianalyzerFilterMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LogFortianalyzerFilter] {
	return pulumix.Output[map[string]*LogFortianalyzerFilter]{
		OutputState: i.ToLogFortianalyzerFilterMapOutputWithContext(ctx).OutputState,
	}
}

type LogFortianalyzerFilterOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzerFilter)(nil)).Elem()
}

func (o LogFortianalyzerFilterOutput) ToLogFortianalyzerFilterOutput() LogFortianalyzerFilterOutput {
	return o
}

func (o LogFortianalyzerFilterOutput) ToLogFortianalyzerFilterOutputWithContext(ctx context.Context) LogFortianalyzerFilterOutput {
	return o
}

func (o LogFortianalyzerFilterOutput) ToOutput(ctx context.Context) pulumix.Output[*LogFortianalyzerFilter] {
	return pulumix.Output[*LogFortianalyzerFilter]{
		OutputState: o.OutputState,
	}
}

func (o LogFortianalyzerFilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) DlpArchive() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.DlpArchive }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzerFilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) FreeStyles() LogFortianalyzerFilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) LogFortianalyzerFilterFreeStyleArrayOutput { return v.FreeStyles }).(LogFortianalyzerFilterFreeStyleArrayOutput)
}

func (o LogFortianalyzerFilterOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzerFilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzerFilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

func (o LogFortianalyzerFilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzerFilter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
}

type LogFortianalyzerFilterArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzerFilter)(nil)).Elem()
}

func (o LogFortianalyzerFilterArrayOutput) ToLogFortianalyzerFilterArrayOutput() LogFortianalyzerFilterArrayOutput {
	return o
}

func (o LogFortianalyzerFilterArrayOutput) ToLogFortianalyzerFilterArrayOutputWithContext(ctx context.Context) LogFortianalyzerFilterArrayOutput {
	return o
}

func (o LogFortianalyzerFilterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LogFortianalyzerFilter] {
	return pulumix.Output[[]*LogFortianalyzerFilter]{
		OutputState: o.OutputState,
	}
}

func (o LogFortianalyzerFilterArrayOutput) Index(i pulumi.IntInput) LogFortianalyzerFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortianalyzerFilter {
		return vs[0].([]*LogFortianalyzerFilter)[vs[1].(int)]
	}).(LogFortianalyzerFilterOutput)
}

type LogFortianalyzerFilterMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzerFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzerFilter)(nil)).Elem()
}

func (o LogFortianalyzerFilterMapOutput) ToLogFortianalyzerFilterMapOutput() LogFortianalyzerFilterMapOutput {
	return o
}

func (o LogFortianalyzerFilterMapOutput) ToLogFortianalyzerFilterMapOutputWithContext(ctx context.Context) LogFortianalyzerFilterMapOutput {
	return o
}

func (o LogFortianalyzerFilterMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LogFortianalyzerFilter] {
	return pulumix.Output[map[string]*LogFortianalyzerFilter]{
		OutputState: o.OutputState,
	}
}

func (o LogFortianalyzerFilterMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzerFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortianalyzerFilter {
		return vs[0].(map[string]*LogFortianalyzerFilter)[vs[1].(string)]
	}).(LogFortianalyzerFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerFilterInput)(nil)).Elem(), &LogFortianalyzerFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerFilterArrayInput)(nil)).Elem(), LogFortianalyzerFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzerFilterMapInput)(nil)).Elem(), LogFortianalyzerFilterMap{})
	pulumi.RegisterOutputType(LogFortianalyzerFilterOutput{})
	pulumi.RegisterOutputType(LogFortianalyzerFilterArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzerFilterMapOutput{})
}

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

type LogFortianalyzer3Filter struct {
	pulumi.CustomResourceState

	Anomaly              pulumi.StringOutput                         `pulumi:"anomaly"`
	DlpArchive           pulumi.StringOutput                         `pulumi:"dlpArchive"`
	Dns                  pulumi.StringOutput                         `pulumi:"dns"`
	DynamicSortSubtable  pulumi.StringPtrOutput                      `pulumi:"dynamicSortSubtable"`
	Filter               pulumi.StringOutput                         `pulumi:"filter"`
	FilterType           pulumi.StringOutput                         `pulumi:"filterType"`
	ForwardTraffic       pulumi.StringOutput                         `pulumi:"forwardTraffic"`
	FreeStyles           LogFortianalyzer3FilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	GetAllTables         pulumi.StringPtrOutput                      `pulumi:"getAllTables"`
	Gtp                  pulumi.StringOutput                         `pulumi:"gtp"`
	LocalTraffic         pulumi.StringOutput                         `pulumi:"localTraffic"`
	MulticastTraffic     pulumi.StringOutput                         `pulumi:"multicastTraffic"`
	NetscanDiscovery     pulumi.StringOutput                         `pulumi:"netscanDiscovery"`
	NetscanVulnerability pulumi.StringOutput                         `pulumi:"netscanVulnerability"`
	Severity             pulumi.StringOutput                         `pulumi:"severity"`
	SnifferTraffic       pulumi.StringOutput                         `pulumi:"snifferTraffic"`
	Ssh                  pulumi.StringOutput                         `pulumi:"ssh"`
	Vdomparam            pulumi.StringPtrOutput                      `pulumi:"vdomparam"`
	Voip                 pulumi.StringOutput                         `pulumi:"voip"`
	ZtnaTraffic          pulumi.StringOutput                         `pulumi:"ztnaTraffic"`
}

// NewLogFortianalyzer3Filter registers a new resource with the given unique name, arguments, and options.
func NewLogFortianalyzer3Filter(ctx *pulumi.Context,
	name string, args *LogFortianalyzer3FilterArgs, opts ...pulumi.ResourceOption) (*LogFortianalyzer3Filter, error) {
	if args == nil {
		args = &LogFortianalyzer3FilterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogFortianalyzer3Filter
	err := ctx.RegisterResource("fortios:index/logFortianalyzer3Filter:LogFortianalyzer3Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortianalyzer3Filter gets an existing LogFortianalyzer3Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortianalyzer3Filter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortianalyzer3FilterState, opts ...pulumi.ResourceOption) (*LogFortianalyzer3Filter, error) {
	var resource LogFortianalyzer3Filter
	err := ctx.ReadResource("fortios:index/logFortianalyzer3Filter:LogFortianalyzer3Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortianalyzer3Filter resources.
type logFortianalyzer3FilterState struct {
	Anomaly              *string                            `pulumi:"anomaly"`
	DlpArchive           *string                            `pulumi:"dlpArchive"`
	Dns                  *string                            `pulumi:"dns"`
	DynamicSortSubtable  *string                            `pulumi:"dynamicSortSubtable"`
	Filter               *string                            `pulumi:"filter"`
	FilterType           *string                            `pulumi:"filterType"`
	ForwardTraffic       *string                            `pulumi:"forwardTraffic"`
	FreeStyles           []LogFortianalyzer3FilterFreeStyle `pulumi:"freeStyles"`
	GetAllTables         *string                            `pulumi:"getAllTables"`
	Gtp                  *string                            `pulumi:"gtp"`
	LocalTraffic         *string                            `pulumi:"localTraffic"`
	MulticastTraffic     *string                            `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                            `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                            `pulumi:"netscanVulnerability"`
	Severity             *string                            `pulumi:"severity"`
	SnifferTraffic       *string                            `pulumi:"snifferTraffic"`
	Ssh                  *string                            `pulumi:"ssh"`
	Vdomparam            *string                            `pulumi:"vdomparam"`
	Voip                 *string                            `pulumi:"voip"`
	ZtnaTraffic          *string                            `pulumi:"ztnaTraffic"`
}

type LogFortianalyzer3FilterState struct {
	Anomaly              pulumi.StringPtrInput
	DlpArchive           pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogFortianalyzer3FilterFreeStyleArrayInput
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

func (LogFortianalyzer3FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3FilterState)(nil)).Elem()
}

type logFortianalyzer3FilterArgs struct {
	Anomaly              *string                            `pulumi:"anomaly"`
	DlpArchive           *string                            `pulumi:"dlpArchive"`
	Dns                  *string                            `pulumi:"dns"`
	DynamicSortSubtable  *string                            `pulumi:"dynamicSortSubtable"`
	Filter               *string                            `pulumi:"filter"`
	FilterType           *string                            `pulumi:"filterType"`
	ForwardTraffic       *string                            `pulumi:"forwardTraffic"`
	FreeStyles           []LogFortianalyzer3FilterFreeStyle `pulumi:"freeStyles"`
	GetAllTables         *string                            `pulumi:"getAllTables"`
	Gtp                  *string                            `pulumi:"gtp"`
	LocalTraffic         *string                            `pulumi:"localTraffic"`
	MulticastTraffic     *string                            `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                            `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                            `pulumi:"netscanVulnerability"`
	Severity             *string                            `pulumi:"severity"`
	SnifferTraffic       *string                            `pulumi:"snifferTraffic"`
	Ssh                  *string                            `pulumi:"ssh"`
	Vdomparam            *string                            `pulumi:"vdomparam"`
	Voip                 *string                            `pulumi:"voip"`
	ZtnaTraffic          *string                            `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a LogFortianalyzer3Filter resource.
type LogFortianalyzer3FilterArgs struct {
	Anomaly              pulumi.StringPtrInput
	DlpArchive           pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogFortianalyzer3FilterFreeStyleArrayInput
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

func (LogFortianalyzer3FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortianalyzer3FilterArgs)(nil)).Elem()
}

type LogFortianalyzer3FilterInput interface {
	pulumi.Input

	ToLogFortianalyzer3FilterOutput() LogFortianalyzer3FilterOutput
	ToLogFortianalyzer3FilterOutputWithContext(ctx context.Context) LogFortianalyzer3FilterOutput
}

func (*LogFortianalyzer3Filter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer3Filter)(nil)).Elem()
}

func (i *LogFortianalyzer3Filter) ToLogFortianalyzer3FilterOutput() LogFortianalyzer3FilterOutput {
	return i.ToLogFortianalyzer3FilterOutputWithContext(context.Background())
}

func (i *LogFortianalyzer3Filter) ToLogFortianalyzer3FilterOutputWithContext(ctx context.Context) LogFortianalyzer3FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3FilterOutput)
}

func (i *LogFortianalyzer3Filter) ToOutput(ctx context.Context) pulumix.Output[*LogFortianalyzer3Filter] {
	return pulumix.Output[*LogFortianalyzer3Filter]{
		OutputState: i.ToLogFortianalyzer3FilterOutputWithContext(ctx).OutputState,
	}
}

// LogFortianalyzer3FilterArrayInput is an input type that accepts LogFortianalyzer3FilterArray and LogFortianalyzer3FilterArrayOutput values.
// You can construct a concrete instance of `LogFortianalyzer3FilterArrayInput` via:
//
//	LogFortianalyzer3FilterArray{ LogFortianalyzer3FilterArgs{...} }
type LogFortianalyzer3FilterArrayInput interface {
	pulumi.Input

	ToLogFortianalyzer3FilterArrayOutput() LogFortianalyzer3FilterArrayOutput
	ToLogFortianalyzer3FilterArrayOutputWithContext(context.Context) LogFortianalyzer3FilterArrayOutput
}

type LogFortianalyzer3FilterArray []LogFortianalyzer3FilterInput

func (LogFortianalyzer3FilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzer3Filter)(nil)).Elem()
}

func (i LogFortianalyzer3FilterArray) ToLogFortianalyzer3FilterArrayOutput() LogFortianalyzer3FilterArrayOutput {
	return i.ToLogFortianalyzer3FilterArrayOutputWithContext(context.Background())
}

func (i LogFortianalyzer3FilterArray) ToLogFortianalyzer3FilterArrayOutputWithContext(ctx context.Context) LogFortianalyzer3FilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3FilterArrayOutput)
}

func (i LogFortianalyzer3FilterArray) ToOutput(ctx context.Context) pulumix.Output[[]*LogFortianalyzer3Filter] {
	return pulumix.Output[[]*LogFortianalyzer3Filter]{
		OutputState: i.ToLogFortianalyzer3FilterArrayOutputWithContext(ctx).OutputState,
	}
}

// LogFortianalyzer3FilterMapInput is an input type that accepts LogFortianalyzer3FilterMap and LogFortianalyzer3FilterMapOutput values.
// You can construct a concrete instance of `LogFortianalyzer3FilterMapInput` via:
//
//	LogFortianalyzer3FilterMap{ "key": LogFortianalyzer3FilterArgs{...} }
type LogFortianalyzer3FilterMapInput interface {
	pulumi.Input

	ToLogFortianalyzer3FilterMapOutput() LogFortianalyzer3FilterMapOutput
	ToLogFortianalyzer3FilterMapOutputWithContext(context.Context) LogFortianalyzer3FilterMapOutput
}

type LogFortianalyzer3FilterMap map[string]LogFortianalyzer3FilterInput

func (LogFortianalyzer3FilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzer3Filter)(nil)).Elem()
}

func (i LogFortianalyzer3FilterMap) ToLogFortianalyzer3FilterMapOutput() LogFortianalyzer3FilterMapOutput {
	return i.ToLogFortianalyzer3FilterMapOutputWithContext(context.Background())
}

func (i LogFortianalyzer3FilterMap) ToLogFortianalyzer3FilterMapOutputWithContext(ctx context.Context) LogFortianalyzer3FilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortianalyzer3FilterMapOutput)
}

func (i LogFortianalyzer3FilterMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LogFortianalyzer3Filter] {
	return pulumix.Output[map[string]*LogFortianalyzer3Filter]{
		OutputState: i.ToLogFortianalyzer3FilterMapOutputWithContext(ctx).OutputState,
	}
}

type LogFortianalyzer3FilterOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortianalyzer3Filter)(nil)).Elem()
}

func (o LogFortianalyzer3FilterOutput) ToLogFortianalyzer3FilterOutput() LogFortianalyzer3FilterOutput {
	return o
}

func (o LogFortianalyzer3FilterOutput) ToLogFortianalyzer3FilterOutputWithContext(ctx context.Context) LogFortianalyzer3FilterOutput {
	return o
}

func (o LogFortianalyzer3FilterOutput) ToOutput(ctx context.Context) pulumix.Output[*LogFortianalyzer3Filter] {
	return pulumix.Output[*LogFortianalyzer3Filter]{
		OutputState: o.OutputState,
	}
}

func (o LogFortianalyzer3FilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) DlpArchive() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.DlpArchive }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzer3FilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) FreeStyles() LogFortianalyzer3FilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) LogFortianalyzer3FilterFreeStyleArrayOutput { return v.FreeStyles }).(LogFortianalyzer3FilterFreeStyleArrayOutput)
}

func (o LogFortianalyzer3FilterOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzer3FilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LogFortianalyzer3FilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

func (o LogFortianalyzer3FilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortianalyzer3Filter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
}

type LogFortianalyzer3FilterArrayOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3FilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortianalyzer3Filter)(nil)).Elem()
}

func (o LogFortianalyzer3FilterArrayOutput) ToLogFortianalyzer3FilterArrayOutput() LogFortianalyzer3FilterArrayOutput {
	return o
}

func (o LogFortianalyzer3FilterArrayOutput) ToLogFortianalyzer3FilterArrayOutputWithContext(ctx context.Context) LogFortianalyzer3FilterArrayOutput {
	return o
}

func (o LogFortianalyzer3FilterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LogFortianalyzer3Filter] {
	return pulumix.Output[[]*LogFortianalyzer3Filter]{
		OutputState: o.OutputState,
	}
}

func (o LogFortianalyzer3FilterArrayOutput) Index(i pulumi.IntInput) LogFortianalyzer3FilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortianalyzer3Filter {
		return vs[0].([]*LogFortianalyzer3Filter)[vs[1].(int)]
	}).(LogFortianalyzer3FilterOutput)
}

type LogFortianalyzer3FilterMapOutput struct{ *pulumi.OutputState }

func (LogFortianalyzer3FilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortianalyzer3Filter)(nil)).Elem()
}

func (o LogFortianalyzer3FilterMapOutput) ToLogFortianalyzer3FilterMapOutput() LogFortianalyzer3FilterMapOutput {
	return o
}

func (o LogFortianalyzer3FilterMapOutput) ToLogFortianalyzer3FilterMapOutputWithContext(ctx context.Context) LogFortianalyzer3FilterMapOutput {
	return o
}

func (o LogFortianalyzer3FilterMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LogFortianalyzer3Filter] {
	return pulumix.Output[map[string]*LogFortianalyzer3Filter]{
		OutputState: o.OutputState,
	}
}

func (o LogFortianalyzer3FilterMapOutput) MapIndex(k pulumi.StringInput) LogFortianalyzer3FilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortianalyzer3Filter {
		return vs[0].(map[string]*LogFortianalyzer3Filter)[vs[1].(string)]
	}).(LogFortianalyzer3FilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer3FilterInput)(nil)).Elem(), &LogFortianalyzer3Filter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer3FilterArrayInput)(nil)).Elem(), LogFortianalyzer3FilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortianalyzer3FilterMapInput)(nil)).Elem(), LogFortianalyzer3FilterMap{})
	pulumi.RegisterOutputType(LogFortianalyzer3FilterOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3FilterArrayOutput{})
	pulumi.RegisterOutputType(LogFortianalyzer3FilterMapOutput{})
}

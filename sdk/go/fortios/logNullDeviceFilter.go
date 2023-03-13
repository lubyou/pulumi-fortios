// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogNullDeviceFilter struct {
	pulumi.CustomResourceState

	Anomaly              pulumi.StringOutput                     `pulumi:"anomaly"`
	Dns                  pulumi.StringOutput                     `pulumi:"dns"`
	DynamicSortSubtable  pulumi.StringPtrOutput                  `pulumi:"dynamicSortSubtable"`
	Filter               pulumi.StringOutput                     `pulumi:"filter"`
	FilterType           pulumi.StringOutput                     `pulumi:"filterType"`
	ForwardTraffic       pulumi.StringOutput                     `pulumi:"forwardTraffic"`
	FreeStyles           LogNullDeviceFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	Gtp                  pulumi.StringOutput                     `pulumi:"gtp"`
	LocalTraffic         pulumi.StringOutput                     `pulumi:"localTraffic"`
	MulticastTraffic     pulumi.StringOutput                     `pulumi:"multicastTraffic"`
	NetscanDiscovery     pulumi.StringOutput                     `pulumi:"netscanDiscovery"`
	NetscanVulnerability pulumi.StringOutput                     `pulumi:"netscanVulnerability"`
	Severity             pulumi.StringOutput                     `pulumi:"severity"`
	SnifferTraffic       pulumi.StringOutput                     `pulumi:"snifferTraffic"`
	Ssh                  pulumi.StringOutput                     `pulumi:"ssh"`
	Vdomparam            pulumi.StringPtrOutput                  `pulumi:"vdomparam"`
	Voip                 pulumi.StringOutput                     `pulumi:"voip"`
	ZtnaTraffic          pulumi.StringOutput                     `pulumi:"ztnaTraffic"`
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
	Anomaly              *string                        `pulumi:"anomaly"`
	Dns                  *string                        `pulumi:"dns"`
	DynamicSortSubtable  *string                        `pulumi:"dynamicSortSubtable"`
	Filter               *string                        `pulumi:"filter"`
	FilterType           *string                        `pulumi:"filterType"`
	ForwardTraffic       *string                        `pulumi:"forwardTraffic"`
	FreeStyles           []LogNullDeviceFilterFreeStyle `pulumi:"freeStyles"`
	Gtp                  *string                        `pulumi:"gtp"`
	LocalTraffic         *string                        `pulumi:"localTraffic"`
	MulticastTraffic     *string                        `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                        `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                        `pulumi:"netscanVulnerability"`
	Severity             *string                        `pulumi:"severity"`
	SnifferTraffic       *string                        `pulumi:"snifferTraffic"`
	Ssh                  *string                        `pulumi:"ssh"`
	Vdomparam            *string                        `pulumi:"vdomparam"`
	Voip                 *string                        `pulumi:"voip"`
	ZtnaTraffic          *string                        `pulumi:"ztnaTraffic"`
}

type LogNullDeviceFilterState struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogNullDeviceFilterFreeStyleArrayInput
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

func (LogNullDeviceFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logNullDeviceFilterState)(nil)).Elem()
}

type logNullDeviceFilterArgs struct {
	Anomaly              *string                        `pulumi:"anomaly"`
	Dns                  *string                        `pulumi:"dns"`
	DynamicSortSubtable  *string                        `pulumi:"dynamicSortSubtable"`
	Filter               *string                        `pulumi:"filter"`
	FilterType           *string                        `pulumi:"filterType"`
	ForwardTraffic       *string                        `pulumi:"forwardTraffic"`
	FreeStyles           []LogNullDeviceFilterFreeStyle `pulumi:"freeStyles"`
	Gtp                  *string                        `pulumi:"gtp"`
	LocalTraffic         *string                        `pulumi:"localTraffic"`
	MulticastTraffic     *string                        `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                        `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                        `pulumi:"netscanVulnerability"`
	Severity             *string                        `pulumi:"severity"`
	SnifferTraffic       *string                        `pulumi:"snifferTraffic"`
	Ssh                  *string                        `pulumi:"ssh"`
	Vdomparam            *string                        `pulumi:"vdomparam"`
	Voip                 *string                        `pulumi:"voip"`
	ZtnaTraffic          *string                        `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a LogNullDeviceFilter resource.
type LogNullDeviceFilterArgs struct {
	Anomaly              pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogNullDeviceFilterFreeStyleArrayInput
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
//	LogNullDeviceFilterArray{ LogNullDeviceFilterArgs{...} }
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
//	LogNullDeviceFilterMap{ "key": LogNullDeviceFilterArgs{...} }
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

func (o LogNullDeviceFilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogNullDeviceFilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) FreeStyles() LogNullDeviceFilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) LogNullDeviceFilterFreeStyleArrayOutput { return v.FreeStyles }).(LogNullDeviceFilterFreeStyleArrayOutput)
}

func (o LogNullDeviceFilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LogNullDeviceFilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

func (o LogNullDeviceFilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogNullDeviceFilter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
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

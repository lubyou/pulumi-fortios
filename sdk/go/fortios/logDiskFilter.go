// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogDiskFilter struct {
	pulumi.CustomResourceState

	Admin                pulumi.StringOutput               `pulumi:"admin"`
	Anomaly              pulumi.StringOutput               `pulumi:"anomaly"`
	Auth                 pulumi.StringOutput               `pulumi:"auth"`
	CpuMemoryUsage       pulumi.StringOutput               `pulumi:"cpuMemoryUsage"`
	Dhcp                 pulumi.StringOutput               `pulumi:"dhcp"`
	DlpArchive           pulumi.StringOutput               `pulumi:"dlpArchive"`
	Dns                  pulumi.StringOutput               `pulumi:"dns"`
	DynamicSortSubtable  pulumi.StringPtrOutput            `pulumi:"dynamicSortSubtable"`
	Event                pulumi.StringOutput               `pulumi:"event"`
	Filter               pulumi.StringOutput               `pulumi:"filter"`
	FilterType           pulumi.StringOutput               `pulumi:"filterType"`
	ForwardTraffic       pulumi.StringOutput               `pulumi:"forwardTraffic"`
	FreeStyles           LogDiskFilterFreeStyleArrayOutput `pulumi:"freeStyles"`
	Gtp                  pulumi.StringOutput               `pulumi:"gtp"`
	Ha                   pulumi.StringOutput               `pulumi:"ha"`
	Ipsec                pulumi.StringOutput               `pulumi:"ipsec"`
	LdbMonitor           pulumi.StringOutput               `pulumi:"ldbMonitor"`
	LocalTraffic         pulumi.StringOutput               `pulumi:"localTraffic"`
	MulticastTraffic     pulumi.StringOutput               `pulumi:"multicastTraffic"`
	NetscanDiscovery     pulumi.StringOutput               `pulumi:"netscanDiscovery"`
	NetscanVulnerability pulumi.StringOutput               `pulumi:"netscanVulnerability"`
	Pattern              pulumi.StringOutput               `pulumi:"pattern"`
	Ppp                  pulumi.StringOutput               `pulumi:"ppp"`
	Radius               pulumi.StringOutput               `pulumi:"radius"`
	Severity             pulumi.StringOutput               `pulumi:"severity"`
	SnifferTraffic       pulumi.StringOutput               `pulumi:"snifferTraffic"`
	Ssh                  pulumi.StringOutput               `pulumi:"ssh"`
	SslvpnLogAdm         pulumi.StringOutput               `pulumi:"sslvpnLogAdm"`
	SslvpnLogAuth        pulumi.StringOutput               `pulumi:"sslvpnLogAuth"`
	SslvpnLogSession     pulumi.StringOutput               `pulumi:"sslvpnLogSession"`
	System               pulumi.StringOutput               `pulumi:"system"`
	Vdomparam            pulumi.StringPtrOutput            `pulumi:"vdomparam"`
	VipSsl               pulumi.StringOutput               `pulumi:"vipSsl"`
	Voip                 pulumi.StringOutput               `pulumi:"voip"`
	WanOpt               pulumi.StringOutput               `pulumi:"wanOpt"`
	WirelessActivity     pulumi.StringOutput               `pulumi:"wirelessActivity"`
	ZtnaTraffic          pulumi.StringOutput               `pulumi:"ztnaTraffic"`
}

// NewLogDiskFilter registers a new resource with the given unique name, arguments, and options.
func NewLogDiskFilter(ctx *pulumi.Context,
	name string, args *LogDiskFilterArgs, opts ...pulumi.ResourceOption) (*LogDiskFilter, error) {
	if args == nil {
		args = &LogDiskFilterArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogDiskFilter
	err := ctx.RegisterResource("fortios:index/logDiskFilter:LogDiskFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogDiskFilter gets an existing LogDiskFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogDiskFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogDiskFilterState, opts ...pulumi.ResourceOption) (*LogDiskFilter, error) {
	var resource LogDiskFilter
	err := ctx.ReadResource("fortios:index/logDiskFilter:LogDiskFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogDiskFilter resources.
type logDiskFilterState struct {
	Admin                *string                  `pulumi:"admin"`
	Anomaly              *string                  `pulumi:"anomaly"`
	Auth                 *string                  `pulumi:"auth"`
	CpuMemoryUsage       *string                  `pulumi:"cpuMemoryUsage"`
	Dhcp                 *string                  `pulumi:"dhcp"`
	DlpArchive           *string                  `pulumi:"dlpArchive"`
	Dns                  *string                  `pulumi:"dns"`
	DynamicSortSubtable  *string                  `pulumi:"dynamicSortSubtable"`
	Event                *string                  `pulumi:"event"`
	Filter               *string                  `pulumi:"filter"`
	FilterType           *string                  `pulumi:"filterType"`
	ForwardTraffic       *string                  `pulumi:"forwardTraffic"`
	FreeStyles           []LogDiskFilterFreeStyle `pulumi:"freeStyles"`
	Gtp                  *string                  `pulumi:"gtp"`
	Ha                   *string                  `pulumi:"ha"`
	Ipsec                *string                  `pulumi:"ipsec"`
	LdbMonitor           *string                  `pulumi:"ldbMonitor"`
	LocalTraffic         *string                  `pulumi:"localTraffic"`
	MulticastTraffic     *string                  `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                  `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                  `pulumi:"netscanVulnerability"`
	Pattern              *string                  `pulumi:"pattern"`
	Ppp                  *string                  `pulumi:"ppp"`
	Radius               *string                  `pulumi:"radius"`
	Severity             *string                  `pulumi:"severity"`
	SnifferTraffic       *string                  `pulumi:"snifferTraffic"`
	Ssh                  *string                  `pulumi:"ssh"`
	SslvpnLogAdm         *string                  `pulumi:"sslvpnLogAdm"`
	SslvpnLogAuth        *string                  `pulumi:"sslvpnLogAuth"`
	SslvpnLogSession     *string                  `pulumi:"sslvpnLogSession"`
	System               *string                  `pulumi:"system"`
	Vdomparam            *string                  `pulumi:"vdomparam"`
	VipSsl               *string                  `pulumi:"vipSsl"`
	Voip                 *string                  `pulumi:"voip"`
	WanOpt               *string                  `pulumi:"wanOpt"`
	WirelessActivity     *string                  `pulumi:"wirelessActivity"`
	ZtnaTraffic          *string                  `pulumi:"ztnaTraffic"`
}

type LogDiskFilterState struct {
	Admin                pulumi.StringPtrInput
	Anomaly              pulumi.StringPtrInput
	Auth                 pulumi.StringPtrInput
	CpuMemoryUsage       pulumi.StringPtrInput
	Dhcp                 pulumi.StringPtrInput
	DlpArchive           pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Event                pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogDiskFilterFreeStyleArrayInput
	Gtp                  pulumi.StringPtrInput
	Ha                   pulumi.StringPtrInput
	Ipsec                pulumi.StringPtrInput
	LdbMonitor           pulumi.StringPtrInput
	LocalTraffic         pulumi.StringPtrInput
	MulticastTraffic     pulumi.StringPtrInput
	NetscanDiscovery     pulumi.StringPtrInput
	NetscanVulnerability pulumi.StringPtrInput
	Pattern              pulumi.StringPtrInput
	Ppp                  pulumi.StringPtrInput
	Radius               pulumi.StringPtrInput
	Severity             pulumi.StringPtrInput
	SnifferTraffic       pulumi.StringPtrInput
	Ssh                  pulumi.StringPtrInput
	SslvpnLogAdm         pulumi.StringPtrInput
	SslvpnLogAuth        pulumi.StringPtrInput
	SslvpnLogSession     pulumi.StringPtrInput
	System               pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
	VipSsl               pulumi.StringPtrInput
	Voip                 pulumi.StringPtrInput
	WanOpt               pulumi.StringPtrInput
	WirelessActivity     pulumi.StringPtrInput
	ZtnaTraffic          pulumi.StringPtrInput
}

func (LogDiskFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logDiskFilterState)(nil)).Elem()
}

type logDiskFilterArgs struct {
	Admin                *string                  `pulumi:"admin"`
	Anomaly              *string                  `pulumi:"anomaly"`
	Auth                 *string                  `pulumi:"auth"`
	CpuMemoryUsage       *string                  `pulumi:"cpuMemoryUsage"`
	Dhcp                 *string                  `pulumi:"dhcp"`
	DlpArchive           *string                  `pulumi:"dlpArchive"`
	Dns                  *string                  `pulumi:"dns"`
	DynamicSortSubtable  *string                  `pulumi:"dynamicSortSubtable"`
	Event                *string                  `pulumi:"event"`
	Filter               *string                  `pulumi:"filter"`
	FilterType           *string                  `pulumi:"filterType"`
	ForwardTraffic       *string                  `pulumi:"forwardTraffic"`
	FreeStyles           []LogDiskFilterFreeStyle `pulumi:"freeStyles"`
	Gtp                  *string                  `pulumi:"gtp"`
	Ha                   *string                  `pulumi:"ha"`
	Ipsec                *string                  `pulumi:"ipsec"`
	LdbMonitor           *string                  `pulumi:"ldbMonitor"`
	LocalTraffic         *string                  `pulumi:"localTraffic"`
	MulticastTraffic     *string                  `pulumi:"multicastTraffic"`
	NetscanDiscovery     *string                  `pulumi:"netscanDiscovery"`
	NetscanVulnerability *string                  `pulumi:"netscanVulnerability"`
	Pattern              *string                  `pulumi:"pattern"`
	Ppp                  *string                  `pulumi:"ppp"`
	Radius               *string                  `pulumi:"radius"`
	Severity             *string                  `pulumi:"severity"`
	SnifferTraffic       *string                  `pulumi:"snifferTraffic"`
	Ssh                  *string                  `pulumi:"ssh"`
	SslvpnLogAdm         *string                  `pulumi:"sslvpnLogAdm"`
	SslvpnLogAuth        *string                  `pulumi:"sslvpnLogAuth"`
	SslvpnLogSession     *string                  `pulumi:"sslvpnLogSession"`
	System               *string                  `pulumi:"system"`
	Vdomparam            *string                  `pulumi:"vdomparam"`
	VipSsl               *string                  `pulumi:"vipSsl"`
	Voip                 *string                  `pulumi:"voip"`
	WanOpt               *string                  `pulumi:"wanOpt"`
	WirelessActivity     *string                  `pulumi:"wirelessActivity"`
	ZtnaTraffic          *string                  `pulumi:"ztnaTraffic"`
}

// The set of arguments for constructing a LogDiskFilter resource.
type LogDiskFilterArgs struct {
	Admin                pulumi.StringPtrInput
	Anomaly              pulumi.StringPtrInput
	Auth                 pulumi.StringPtrInput
	CpuMemoryUsage       pulumi.StringPtrInput
	Dhcp                 pulumi.StringPtrInput
	DlpArchive           pulumi.StringPtrInput
	Dns                  pulumi.StringPtrInput
	DynamicSortSubtable  pulumi.StringPtrInput
	Event                pulumi.StringPtrInput
	Filter               pulumi.StringPtrInput
	FilterType           pulumi.StringPtrInput
	ForwardTraffic       pulumi.StringPtrInput
	FreeStyles           LogDiskFilterFreeStyleArrayInput
	Gtp                  pulumi.StringPtrInput
	Ha                   pulumi.StringPtrInput
	Ipsec                pulumi.StringPtrInput
	LdbMonitor           pulumi.StringPtrInput
	LocalTraffic         pulumi.StringPtrInput
	MulticastTraffic     pulumi.StringPtrInput
	NetscanDiscovery     pulumi.StringPtrInput
	NetscanVulnerability pulumi.StringPtrInput
	Pattern              pulumi.StringPtrInput
	Ppp                  pulumi.StringPtrInput
	Radius               pulumi.StringPtrInput
	Severity             pulumi.StringPtrInput
	SnifferTraffic       pulumi.StringPtrInput
	Ssh                  pulumi.StringPtrInput
	SslvpnLogAdm         pulumi.StringPtrInput
	SslvpnLogAuth        pulumi.StringPtrInput
	SslvpnLogSession     pulumi.StringPtrInput
	System               pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
	VipSsl               pulumi.StringPtrInput
	Voip                 pulumi.StringPtrInput
	WanOpt               pulumi.StringPtrInput
	WirelessActivity     pulumi.StringPtrInput
	ZtnaTraffic          pulumi.StringPtrInput
}

func (LogDiskFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logDiskFilterArgs)(nil)).Elem()
}

type LogDiskFilterInput interface {
	pulumi.Input

	ToLogDiskFilterOutput() LogDiskFilterOutput
	ToLogDiskFilterOutputWithContext(ctx context.Context) LogDiskFilterOutput
}

func (*LogDiskFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**LogDiskFilter)(nil)).Elem()
}

func (i *LogDiskFilter) ToLogDiskFilterOutput() LogDiskFilterOutput {
	return i.ToLogDiskFilterOutputWithContext(context.Background())
}

func (i *LogDiskFilter) ToLogDiskFilterOutputWithContext(ctx context.Context) LogDiskFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDiskFilterOutput)
}

// LogDiskFilterArrayInput is an input type that accepts LogDiskFilterArray and LogDiskFilterArrayOutput values.
// You can construct a concrete instance of `LogDiskFilterArrayInput` via:
//
//	LogDiskFilterArray{ LogDiskFilterArgs{...} }
type LogDiskFilterArrayInput interface {
	pulumi.Input

	ToLogDiskFilterArrayOutput() LogDiskFilterArrayOutput
	ToLogDiskFilterArrayOutputWithContext(context.Context) LogDiskFilterArrayOutput
}

type LogDiskFilterArray []LogDiskFilterInput

func (LogDiskFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogDiskFilter)(nil)).Elem()
}

func (i LogDiskFilterArray) ToLogDiskFilterArrayOutput() LogDiskFilterArrayOutput {
	return i.ToLogDiskFilterArrayOutputWithContext(context.Background())
}

func (i LogDiskFilterArray) ToLogDiskFilterArrayOutputWithContext(ctx context.Context) LogDiskFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDiskFilterArrayOutput)
}

// LogDiskFilterMapInput is an input type that accepts LogDiskFilterMap and LogDiskFilterMapOutput values.
// You can construct a concrete instance of `LogDiskFilterMapInput` via:
//
//	LogDiskFilterMap{ "key": LogDiskFilterArgs{...} }
type LogDiskFilterMapInput interface {
	pulumi.Input

	ToLogDiskFilterMapOutput() LogDiskFilterMapOutput
	ToLogDiskFilterMapOutputWithContext(context.Context) LogDiskFilterMapOutput
}

type LogDiskFilterMap map[string]LogDiskFilterInput

func (LogDiskFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogDiskFilter)(nil)).Elem()
}

func (i LogDiskFilterMap) ToLogDiskFilterMapOutput() LogDiskFilterMapOutput {
	return i.ToLogDiskFilterMapOutputWithContext(context.Background())
}

func (i LogDiskFilterMap) ToLogDiskFilterMapOutputWithContext(ctx context.Context) LogDiskFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDiskFilterMapOutput)
}

type LogDiskFilterOutput struct{ *pulumi.OutputState }

func (LogDiskFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogDiskFilter)(nil)).Elem()
}

func (o LogDiskFilterOutput) ToLogDiskFilterOutput() LogDiskFilterOutput {
	return o
}

func (o LogDiskFilterOutput) ToLogDiskFilterOutputWithContext(ctx context.Context) LogDiskFilterOutput {
	return o
}

func (o LogDiskFilterOutput) Admin() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Admin }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Anomaly() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Anomaly }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Auth() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Auth }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) CpuMemoryUsage() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.CpuMemoryUsage }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Dhcp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Dhcp }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) DlpArchive() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.DlpArchive }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Dns }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogDiskFilterOutput) Event() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Event }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) FilterType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.FilterType }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) ForwardTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.ForwardTraffic }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) FreeStyles() LogDiskFilterFreeStyleArrayOutput {
	return o.ApplyT(func(v *LogDiskFilter) LogDiskFilterFreeStyleArrayOutput { return v.FreeStyles }).(LogDiskFilterFreeStyleArrayOutput)
}

func (o LogDiskFilterOutput) Gtp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Gtp }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Ha() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Ha }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Ipsec() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Ipsec }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) LdbMonitor() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.LdbMonitor }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) LocalTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.LocalTraffic }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) MulticastTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.MulticastTraffic }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) NetscanDiscovery() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.NetscanDiscovery }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) NetscanVulnerability() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.NetscanVulnerability }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Pattern }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Ppp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Ppp }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Radius() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Radius }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) SnifferTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.SnifferTraffic }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Ssh() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Ssh }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) SslvpnLogAdm() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.SslvpnLogAdm }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) SslvpnLogAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.SslvpnLogAuth }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) SslvpnLogSession() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.SslvpnLogSession }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) System() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.System }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LogDiskFilterOutput) VipSsl() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.VipSsl }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) Voip() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.Voip }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) WanOpt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.WanOpt }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) WirelessActivity() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.WirelessActivity }).(pulumi.StringOutput)
}

func (o LogDiskFilterOutput) ZtnaTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *LogDiskFilter) pulumi.StringOutput { return v.ZtnaTraffic }).(pulumi.StringOutput)
}

type LogDiskFilterArrayOutput struct{ *pulumi.OutputState }

func (LogDiskFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogDiskFilter)(nil)).Elem()
}

func (o LogDiskFilterArrayOutput) ToLogDiskFilterArrayOutput() LogDiskFilterArrayOutput {
	return o
}

func (o LogDiskFilterArrayOutput) ToLogDiskFilterArrayOutputWithContext(ctx context.Context) LogDiskFilterArrayOutput {
	return o
}

func (o LogDiskFilterArrayOutput) Index(i pulumi.IntInput) LogDiskFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogDiskFilter {
		return vs[0].([]*LogDiskFilter)[vs[1].(int)]
	}).(LogDiskFilterOutput)
}

type LogDiskFilterMapOutput struct{ *pulumi.OutputState }

func (LogDiskFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogDiskFilter)(nil)).Elem()
}

func (o LogDiskFilterMapOutput) ToLogDiskFilterMapOutput() LogDiskFilterMapOutput {
	return o
}

func (o LogDiskFilterMapOutput) ToLogDiskFilterMapOutputWithContext(ctx context.Context) LogDiskFilterMapOutput {
	return o
}

func (o LogDiskFilterMapOutput) MapIndex(k pulumi.StringInput) LogDiskFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogDiskFilter {
		return vs[0].(map[string]*LogDiskFilter)[vs[1].(string)]
	}).(LogDiskFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogDiskFilterInput)(nil)).Elem(), &LogDiskFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogDiskFilterArrayInput)(nil)).Elem(), LogDiskFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogDiskFilterMapInput)(nil)).Elem(), LogDiskFilterMap{})
	pulumi.RegisterOutputType(LogDiskFilterOutput{})
	pulumi.RegisterOutputType(LogDiskFilterArrayOutput{})
	pulumi.RegisterOutputType(LogDiskFilterMapOutput{})
}

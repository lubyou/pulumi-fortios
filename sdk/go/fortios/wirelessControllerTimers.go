// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WirelessControllerTimers struct {
	pulumi.CustomResourceState

	AuthTimeout             pulumi.IntOutput                             `pulumi:"authTimeout"`
	BleScanReportIntv       pulumi.IntOutput                             `pulumi:"bleScanReportIntv"`
	ClientIdleRehomeTimeout pulumi.IntOutput                             `pulumi:"clientIdleRehomeTimeout"`
	ClientIdleTimeout       pulumi.IntOutput                             `pulumi:"clientIdleTimeout"`
	DarrpDay                pulumi.StringOutput                          `pulumi:"darrpDay"`
	DarrpOptimize           pulumi.IntOutput                             `pulumi:"darrpOptimize"`
	DarrpTimes              WirelessControllerTimersDarrpTimeArrayOutput `pulumi:"darrpTimes"`
	DiscoveryInterval       pulumi.IntOutput                             `pulumi:"discoveryInterval"`
	DrmaInterval            pulumi.IntOutput                             `pulumi:"drmaInterval"`
	DynamicSortSubtable     pulumi.StringPtrOutput                       `pulumi:"dynamicSortSubtable"`
	EchoInterval            pulumi.IntOutput                             `pulumi:"echoInterval"`
	FakeApLog               pulumi.IntOutput                             `pulumi:"fakeApLog"`
	GetAllTables            pulumi.StringPtrOutput                       `pulumi:"getAllTables"`
	IpsecIntfCleanup        pulumi.IntOutput                             `pulumi:"ipsecIntfCleanup"`
	RadioStatsInterval      pulumi.IntOutput                             `pulumi:"radioStatsInterval"`
	RogueApCleanup          pulumi.IntOutput                             `pulumi:"rogueApCleanup"`
	RogueApLog              pulumi.IntOutput                             `pulumi:"rogueApLog"`
	StaCapabilityInterval   pulumi.IntOutput                             `pulumi:"staCapabilityInterval"`
	StaLocateTimer          pulumi.IntOutput                             `pulumi:"staLocateTimer"`
	StaStatsInterval        pulumi.IntOutput                             `pulumi:"staStatsInterval"`
	VapStatsInterval        pulumi.IntOutput                             `pulumi:"vapStatsInterval"`
	Vdomparam               pulumi.StringPtrOutput                       `pulumi:"vdomparam"`
}

// NewWirelessControllerTimers registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerTimers(ctx *pulumi.Context,
	name string, args *WirelessControllerTimersArgs, opts ...pulumi.ResourceOption) (*WirelessControllerTimers, error) {
	if args == nil {
		args = &WirelessControllerTimersArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerTimers
	err := ctx.RegisterResource("fortios:index/wirelessControllerTimers:WirelessControllerTimers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerTimers gets an existing WirelessControllerTimers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerTimers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerTimersState, opts ...pulumi.ResourceOption) (*WirelessControllerTimers, error) {
	var resource WirelessControllerTimers
	err := ctx.ReadResource("fortios:index/wirelessControllerTimers:WirelessControllerTimers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerTimers resources.
type wirelessControllerTimersState struct {
	AuthTimeout             *int                                `pulumi:"authTimeout"`
	BleScanReportIntv       *int                                `pulumi:"bleScanReportIntv"`
	ClientIdleRehomeTimeout *int                                `pulumi:"clientIdleRehomeTimeout"`
	ClientIdleTimeout       *int                                `pulumi:"clientIdleTimeout"`
	DarrpDay                *string                             `pulumi:"darrpDay"`
	DarrpOptimize           *int                                `pulumi:"darrpOptimize"`
	DarrpTimes              []WirelessControllerTimersDarrpTime `pulumi:"darrpTimes"`
	DiscoveryInterval       *int                                `pulumi:"discoveryInterval"`
	DrmaInterval            *int                                `pulumi:"drmaInterval"`
	DynamicSortSubtable     *string                             `pulumi:"dynamicSortSubtable"`
	EchoInterval            *int                                `pulumi:"echoInterval"`
	FakeApLog               *int                                `pulumi:"fakeApLog"`
	GetAllTables            *string                             `pulumi:"getAllTables"`
	IpsecIntfCleanup        *int                                `pulumi:"ipsecIntfCleanup"`
	RadioStatsInterval      *int                                `pulumi:"radioStatsInterval"`
	RogueApCleanup          *int                                `pulumi:"rogueApCleanup"`
	RogueApLog              *int                                `pulumi:"rogueApLog"`
	StaCapabilityInterval   *int                                `pulumi:"staCapabilityInterval"`
	StaLocateTimer          *int                                `pulumi:"staLocateTimer"`
	StaStatsInterval        *int                                `pulumi:"staStatsInterval"`
	VapStatsInterval        *int                                `pulumi:"vapStatsInterval"`
	Vdomparam               *string                             `pulumi:"vdomparam"`
}

type WirelessControllerTimersState struct {
	AuthTimeout             pulumi.IntPtrInput
	BleScanReportIntv       pulumi.IntPtrInput
	ClientIdleRehomeTimeout pulumi.IntPtrInput
	ClientIdleTimeout       pulumi.IntPtrInput
	DarrpDay                pulumi.StringPtrInput
	DarrpOptimize           pulumi.IntPtrInput
	DarrpTimes              WirelessControllerTimersDarrpTimeArrayInput
	DiscoveryInterval       pulumi.IntPtrInput
	DrmaInterval            pulumi.IntPtrInput
	DynamicSortSubtable     pulumi.StringPtrInput
	EchoInterval            pulumi.IntPtrInput
	FakeApLog               pulumi.IntPtrInput
	GetAllTables            pulumi.StringPtrInput
	IpsecIntfCleanup        pulumi.IntPtrInput
	RadioStatsInterval      pulumi.IntPtrInput
	RogueApCleanup          pulumi.IntPtrInput
	RogueApLog              pulumi.IntPtrInput
	StaCapabilityInterval   pulumi.IntPtrInput
	StaLocateTimer          pulumi.IntPtrInput
	StaStatsInterval        pulumi.IntPtrInput
	VapStatsInterval        pulumi.IntPtrInput
	Vdomparam               pulumi.StringPtrInput
}

func (WirelessControllerTimersState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerTimersState)(nil)).Elem()
}

type wirelessControllerTimersArgs struct {
	AuthTimeout             *int                                `pulumi:"authTimeout"`
	BleScanReportIntv       *int                                `pulumi:"bleScanReportIntv"`
	ClientIdleRehomeTimeout *int                                `pulumi:"clientIdleRehomeTimeout"`
	ClientIdleTimeout       *int                                `pulumi:"clientIdleTimeout"`
	DarrpDay                *string                             `pulumi:"darrpDay"`
	DarrpOptimize           *int                                `pulumi:"darrpOptimize"`
	DarrpTimes              []WirelessControllerTimersDarrpTime `pulumi:"darrpTimes"`
	DiscoveryInterval       *int                                `pulumi:"discoveryInterval"`
	DrmaInterval            *int                                `pulumi:"drmaInterval"`
	DynamicSortSubtable     *string                             `pulumi:"dynamicSortSubtable"`
	EchoInterval            *int                                `pulumi:"echoInterval"`
	FakeApLog               *int                                `pulumi:"fakeApLog"`
	GetAllTables            *string                             `pulumi:"getAllTables"`
	IpsecIntfCleanup        *int                                `pulumi:"ipsecIntfCleanup"`
	RadioStatsInterval      *int                                `pulumi:"radioStatsInterval"`
	RogueApCleanup          *int                                `pulumi:"rogueApCleanup"`
	RogueApLog              *int                                `pulumi:"rogueApLog"`
	StaCapabilityInterval   *int                                `pulumi:"staCapabilityInterval"`
	StaLocateTimer          *int                                `pulumi:"staLocateTimer"`
	StaStatsInterval        *int                                `pulumi:"staStatsInterval"`
	VapStatsInterval        *int                                `pulumi:"vapStatsInterval"`
	Vdomparam               *string                             `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerTimers resource.
type WirelessControllerTimersArgs struct {
	AuthTimeout             pulumi.IntPtrInput
	BleScanReportIntv       pulumi.IntPtrInput
	ClientIdleRehomeTimeout pulumi.IntPtrInput
	ClientIdleTimeout       pulumi.IntPtrInput
	DarrpDay                pulumi.StringPtrInput
	DarrpOptimize           pulumi.IntPtrInput
	DarrpTimes              WirelessControllerTimersDarrpTimeArrayInput
	DiscoveryInterval       pulumi.IntPtrInput
	DrmaInterval            pulumi.IntPtrInput
	DynamicSortSubtable     pulumi.StringPtrInput
	EchoInterval            pulumi.IntPtrInput
	FakeApLog               pulumi.IntPtrInput
	GetAllTables            pulumi.StringPtrInput
	IpsecIntfCleanup        pulumi.IntPtrInput
	RadioStatsInterval      pulumi.IntPtrInput
	RogueApCleanup          pulumi.IntPtrInput
	RogueApLog              pulumi.IntPtrInput
	StaCapabilityInterval   pulumi.IntPtrInput
	StaLocateTimer          pulumi.IntPtrInput
	StaStatsInterval        pulumi.IntPtrInput
	VapStatsInterval        pulumi.IntPtrInput
	Vdomparam               pulumi.StringPtrInput
}

func (WirelessControllerTimersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerTimersArgs)(nil)).Elem()
}

type WirelessControllerTimersInput interface {
	pulumi.Input

	ToWirelessControllerTimersOutput() WirelessControllerTimersOutput
	ToWirelessControllerTimersOutputWithContext(ctx context.Context) WirelessControllerTimersOutput
}

func (*WirelessControllerTimers) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerTimers)(nil)).Elem()
}

func (i *WirelessControllerTimers) ToWirelessControllerTimersOutput() WirelessControllerTimersOutput {
	return i.ToWirelessControllerTimersOutputWithContext(context.Background())
}

func (i *WirelessControllerTimers) ToWirelessControllerTimersOutputWithContext(ctx context.Context) WirelessControllerTimersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerTimersOutput)
}

// WirelessControllerTimersArrayInput is an input type that accepts WirelessControllerTimersArray and WirelessControllerTimersArrayOutput values.
// You can construct a concrete instance of `WirelessControllerTimersArrayInput` via:
//
//	WirelessControllerTimersArray{ WirelessControllerTimersArgs{...} }
type WirelessControllerTimersArrayInput interface {
	pulumi.Input

	ToWirelessControllerTimersArrayOutput() WirelessControllerTimersArrayOutput
	ToWirelessControllerTimersArrayOutputWithContext(context.Context) WirelessControllerTimersArrayOutput
}

type WirelessControllerTimersArray []WirelessControllerTimersInput

func (WirelessControllerTimersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerTimers)(nil)).Elem()
}

func (i WirelessControllerTimersArray) ToWirelessControllerTimersArrayOutput() WirelessControllerTimersArrayOutput {
	return i.ToWirelessControllerTimersArrayOutputWithContext(context.Background())
}

func (i WirelessControllerTimersArray) ToWirelessControllerTimersArrayOutputWithContext(ctx context.Context) WirelessControllerTimersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerTimersArrayOutput)
}

// WirelessControllerTimersMapInput is an input type that accepts WirelessControllerTimersMap and WirelessControllerTimersMapOutput values.
// You can construct a concrete instance of `WirelessControllerTimersMapInput` via:
//
//	WirelessControllerTimersMap{ "key": WirelessControllerTimersArgs{...} }
type WirelessControllerTimersMapInput interface {
	pulumi.Input

	ToWirelessControllerTimersMapOutput() WirelessControllerTimersMapOutput
	ToWirelessControllerTimersMapOutputWithContext(context.Context) WirelessControllerTimersMapOutput
}

type WirelessControllerTimersMap map[string]WirelessControllerTimersInput

func (WirelessControllerTimersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerTimers)(nil)).Elem()
}

func (i WirelessControllerTimersMap) ToWirelessControllerTimersMapOutput() WirelessControllerTimersMapOutput {
	return i.ToWirelessControllerTimersMapOutputWithContext(context.Background())
}

func (i WirelessControllerTimersMap) ToWirelessControllerTimersMapOutputWithContext(ctx context.Context) WirelessControllerTimersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerTimersMapOutput)
}

type WirelessControllerTimersOutput struct{ *pulumi.OutputState }

func (WirelessControllerTimersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerTimers)(nil)).Elem()
}

func (o WirelessControllerTimersOutput) ToWirelessControllerTimersOutput() WirelessControllerTimersOutput {
	return o
}

func (o WirelessControllerTimersOutput) ToWirelessControllerTimersOutputWithContext(ctx context.Context) WirelessControllerTimersOutput {
	return o
}

func (o WirelessControllerTimersOutput) AuthTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.AuthTimeout }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) BleScanReportIntv() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.BleScanReportIntv }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) ClientIdleRehomeTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.ClientIdleRehomeTimeout }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) ClientIdleTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.ClientIdleTimeout }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) DarrpDay() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.StringOutput { return v.DarrpDay }).(pulumi.StringOutput)
}

func (o WirelessControllerTimersOutput) DarrpOptimize() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.DarrpOptimize }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) DarrpTimes() WirelessControllerTimersDarrpTimeArrayOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) WirelessControllerTimersDarrpTimeArrayOutput { return v.DarrpTimes }).(WirelessControllerTimersDarrpTimeArrayOutput)
}

func (o WirelessControllerTimersOutput) DiscoveryInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.DiscoveryInterval }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) DrmaInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.DrmaInterval }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerTimersOutput) EchoInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.EchoInterval }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) FakeApLog() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.FakeApLog }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerTimersOutput) IpsecIntfCleanup() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.IpsecIntfCleanup }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) RadioStatsInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.RadioStatsInterval }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) RogueApCleanup() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.RogueApCleanup }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) RogueApLog() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.RogueApLog }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) StaCapabilityInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.StaCapabilityInterval }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) StaLocateTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.StaLocateTimer }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) StaStatsInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.StaStatsInterval }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) VapStatsInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.IntOutput { return v.VapStatsInterval }).(pulumi.IntOutput)
}

func (o WirelessControllerTimersOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerTimers) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelessControllerTimersArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerTimersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerTimers)(nil)).Elem()
}

func (o WirelessControllerTimersArrayOutput) ToWirelessControllerTimersArrayOutput() WirelessControllerTimersArrayOutput {
	return o
}

func (o WirelessControllerTimersArrayOutput) ToWirelessControllerTimersArrayOutputWithContext(ctx context.Context) WirelessControllerTimersArrayOutput {
	return o
}

func (o WirelessControllerTimersArrayOutput) Index(i pulumi.IntInput) WirelessControllerTimersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerTimers {
		return vs[0].([]*WirelessControllerTimers)[vs[1].(int)]
	}).(WirelessControllerTimersOutput)
}

type WirelessControllerTimersMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerTimersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerTimers)(nil)).Elem()
}

func (o WirelessControllerTimersMapOutput) ToWirelessControllerTimersMapOutput() WirelessControllerTimersMapOutput {
	return o
}

func (o WirelessControllerTimersMapOutput) ToWirelessControllerTimersMapOutputWithContext(ctx context.Context) WirelessControllerTimersMapOutput {
	return o
}

func (o WirelessControllerTimersMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerTimersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerTimers {
		return vs[0].(map[string]*WirelessControllerTimers)[vs[1].(string)]
	}).(WirelessControllerTimersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerTimersInput)(nil)).Elem(), &WirelessControllerTimers{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerTimersArrayInput)(nil)).Elem(), WirelessControllerTimersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerTimersMapInput)(nil)).Elem(), WirelessControllerTimersMap{})
	pulumi.RegisterOutputType(WirelessControllerTimersOutput{})
	pulumi.RegisterOutputType(WirelessControllerTimersArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerTimersMapOutput{})
}

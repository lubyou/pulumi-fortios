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

type WirelessControllerQosProfile struct {
	pulumi.CustomResourceState

	BandwidthAdmissionControl pulumi.StringOutput                              `pulumi:"bandwidthAdmissionControl"`
	BandwidthCapacity         pulumi.IntOutput                                 `pulumi:"bandwidthCapacity"`
	Burst                     pulumi.StringOutput                              `pulumi:"burst"`
	CallAdmissionControl      pulumi.StringOutput                              `pulumi:"callAdmissionControl"`
	CallCapacity              pulumi.IntOutput                                 `pulumi:"callCapacity"`
	Comment                   pulumi.StringOutput                              `pulumi:"comment"`
	Downlink                  pulumi.IntOutput                                 `pulumi:"downlink"`
	DownlinkSta               pulumi.IntOutput                                 `pulumi:"downlinkSta"`
	DscpWmmBes                WirelessControllerQosProfileDscpWmmBeArrayOutput `pulumi:"dscpWmmBes"`
	DscpWmmBks                WirelessControllerQosProfileDscpWmmBkArrayOutput `pulumi:"dscpWmmBks"`
	DscpWmmMapping            pulumi.StringOutput                              `pulumi:"dscpWmmMapping"`
	DscpWmmVis                WirelessControllerQosProfileDscpWmmViArrayOutput `pulumi:"dscpWmmVis"`
	DscpWmmVos                WirelessControllerQosProfileDscpWmmVoArrayOutput `pulumi:"dscpWmmVos"`
	DynamicSortSubtable       pulumi.StringPtrOutput                           `pulumi:"dynamicSortSubtable"`
	GetAllTables              pulumi.StringPtrOutput                           `pulumi:"getAllTables"`
	Name                      pulumi.StringOutput                              `pulumi:"name"`
	Uplink                    pulumi.IntOutput                                 `pulumi:"uplink"`
	UplinkSta                 pulumi.IntOutput                                 `pulumi:"uplinkSta"`
	Vdomparam                 pulumi.StringPtrOutput                           `pulumi:"vdomparam"`
	Wmm                       pulumi.StringOutput                              `pulumi:"wmm"`
	WmmBeDscp                 pulumi.IntOutput                                 `pulumi:"wmmBeDscp"`
	WmmBkDscp                 pulumi.IntOutput                                 `pulumi:"wmmBkDscp"`
	WmmDscpMarking            pulumi.StringOutput                              `pulumi:"wmmDscpMarking"`
	WmmUapsd                  pulumi.StringOutput                              `pulumi:"wmmUapsd"`
	WmmViDscp                 pulumi.IntOutput                                 `pulumi:"wmmViDscp"`
	WmmVoDscp                 pulumi.IntOutput                                 `pulumi:"wmmVoDscp"`
}

// NewWirelessControllerQosProfile registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerQosProfile(ctx *pulumi.Context,
	name string, args *WirelessControllerQosProfileArgs, opts ...pulumi.ResourceOption) (*WirelessControllerQosProfile, error) {
	if args == nil {
		args = &WirelessControllerQosProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerQosProfile
	err := ctx.RegisterResource("fortios:index/wirelessControllerQosProfile:WirelessControllerQosProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerQosProfile gets an existing WirelessControllerQosProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerQosProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerQosProfileState, opts ...pulumi.ResourceOption) (*WirelessControllerQosProfile, error) {
	var resource WirelessControllerQosProfile
	err := ctx.ReadResource("fortios:index/wirelessControllerQosProfile:WirelessControllerQosProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerQosProfile resources.
type wirelessControllerQosProfileState struct {
	BandwidthAdmissionControl *string                                 `pulumi:"bandwidthAdmissionControl"`
	BandwidthCapacity         *int                                    `pulumi:"bandwidthCapacity"`
	Burst                     *string                                 `pulumi:"burst"`
	CallAdmissionControl      *string                                 `pulumi:"callAdmissionControl"`
	CallCapacity              *int                                    `pulumi:"callCapacity"`
	Comment                   *string                                 `pulumi:"comment"`
	Downlink                  *int                                    `pulumi:"downlink"`
	DownlinkSta               *int                                    `pulumi:"downlinkSta"`
	DscpWmmBes                []WirelessControllerQosProfileDscpWmmBe `pulumi:"dscpWmmBes"`
	DscpWmmBks                []WirelessControllerQosProfileDscpWmmBk `pulumi:"dscpWmmBks"`
	DscpWmmMapping            *string                                 `pulumi:"dscpWmmMapping"`
	DscpWmmVis                []WirelessControllerQosProfileDscpWmmVi `pulumi:"dscpWmmVis"`
	DscpWmmVos                []WirelessControllerQosProfileDscpWmmVo `pulumi:"dscpWmmVos"`
	DynamicSortSubtable       *string                                 `pulumi:"dynamicSortSubtable"`
	GetAllTables              *string                                 `pulumi:"getAllTables"`
	Name                      *string                                 `pulumi:"name"`
	Uplink                    *int                                    `pulumi:"uplink"`
	UplinkSta                 *int                                    `pulumi:"uplinkSta"`
	Vdomparam                 *string                                 `pulumi:"vdomparam"`
	Wmm                       *string                                 `pulumi:"wmm"`
	WmmBeDscp                 *int                                    `pulumi:"wmmBeDscp"`
	WmmBkDscp                 *int                                    `pulumi:"wmmBkDscp"`
	WmmDscpMarking            *string                                 `pulumi:"wmmDscpMarking"`
	WmmUapsd                  *string                                 `pulumi:"wmmUapsd"`
	WmmViDscp                 *int                                    `pulumi:"wmmViDscp"`
	WmmVoDscp                 *int                                    `pulumi:"wmmVoDscp"`
}

type WirelessControllerQosProfileState struct {
	BandwidthAdmissionControl pulumi.StringPtrInput
	BandwidthCapacity         pulumi.IntPtrInput
	Burst                     pulumi.StringPtrInput
	CallAdmissionControl      pulumi.StringPtrInput
	CallCapacity              pulumi.IntPtrInput
	Comment                   pulumi.StringPtrInput
	Downlink                  pulumi.IntPtrInput
	DownlinkSta               pulumi.IntPtrInput
	DscpWmmBes                WirelessControllerQosProfileDscpWmmBeArrayInput
	DscpWmmBks                WirelessControllerQosProfileDscpWmmBkArrayInput
	DscpWmmMapping            pulumi.StringPtrInput
	DscpWmmVis                WirelessControllerQosProfileDscpWmmViArrayInput
	DscpWmmVos                WirelessControllerQosProfileDscpWmmVoArrayInput
	DynamicSortSubtable       pulumi.StringPtrInput
	GetAllTables              pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	Uplink                    pulumi.IntPtrInput
	UplinkSta                 pulumi.IntPtrInput
	Vdomparam                 pulumi.StringPtrInput
	Wmm                       pulumi.StringPtrInput
	WmmBeDscp                 pulumi.IntPtrInput
	WmmBkDscp                 pulumi.IntPtrInput
	WmmDscpMarking            pulumi.StringPtrInput
	WmmUapsd                  pulumi.StringPtrInput
	WmmViDscp                 pulumi.IntPtrInput
	WmmVoDscp                 pulumi.IntPtrInput
}

func (WirelessControllerQosProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerQosProfileState)(nil)).Elem()
}

type wirelessControllerQosProfileArgs struct {
	BandwidthAdmissionControl *string                                 `pulumi:"bandwidthAdmissionControl"`
	BandwidthCapacity         *int                                    `pulumi:"bandwidthCapacity"`
	Burst                     *string                                 `pulumi:"burst"`
	CallAdmissionControl      *string                                 `pulumi:"callAdmissionControl"`
	CallCapacity              *int                                    `pulumi:"callCapacity"`
	Comment                   *string                                 `pulumi:"comment"`
	Downlink                  *int                                    `pulumi:"downlink"`
	DownlinkSta               *int                                    `pulumi:"downlinkSta"`
	DscpWmmBes                []WirelessControllerQosProfileDscpWmmBe `pulumi:"dscpWmmBes"`
	DscpWmmBks                []WirelessControllerQosProfileDscpWmmBk `pulumi:"dscpWmmBks"`
	DscpWmmMapping            *string                                 `pulumi:"dscpWmmMapping"`
	DscpWmmVis                []WirelessControllerQosProfileDscpWmmVi `pulumi:"dscpWmmVis"`
	DscpWmmVos                []WirelessControllerQosProfileDscpWmmVo `pulumi:"dscpWmmVos"`
	DynamicSortSubtable       *string                                 `pulumi:"dynamicSortSubtable"`
	GetAllTables              *string                                 `pulumi:"getAllTables"`
	Name                      *string                                 `pulumi:"name"`
	Uplink                    *int                                    `pulumi:"uplink"`
	UplinkSta                 *int                                    `pulumi:"uplinkSta"`
	Vdomparam                 *string                                 `pulumi:"vdomparam"`
	Wmm                       *string                                 `pulumi:"wmm"`
	WmmBeDscp                 *int                                    `pulumi:"wmmBeDscp"`
	WmmBkDscp                 *int                                    `pulumi:"wmmBkDscp"`
	WmmDscpMarking            *string                                 `pulumi:"wmmDscpMarking"`
	WmmUapsd                  *string                                 `pulumi:"wmmUapsd"`
	WmmViDscp                 *int                                    `pulumi:"wmmViDscp"`
	WmmVoDscp                 *int                                    `pulumi:"wmmVoDscp"`
}

// The set of arguments for constructing a WirelessControllerQosProfile resource.
type WirelessControllerQosProfileArgs struct {
	BandwidthAdmissionControl pulumi.StringPtrInput
	BandwidthCapacity         pulumi.IntPtrInput
	Burst                     pulumi.StringPtrInput
	CallAdmissionControl      pulumi.StringPtrInput
	CallCapacity              pulumi.IntPtrInput
	Comment                   pulumi.StringPtrInput
	Downlink                  pulumi.IntPtrInput
	DownlinkSta               pulumi.IntPtrInput
	DscpWmmBes                WirelessControllerQosProfileDscpWmmBeArrayInput
	DscpWmmBks                WirelessControllerQosProfileDscpWmmBkArrayInput
	DscpWmmMapping            pulumi.StringPtrInput
	DscpWmmVis                WirelessControllerQosProfileDscpWmmViArrayInput
	DscpWmmVos                WirelessControllerQosProfileDscpWmmVoArrayInput
	DynamicSortSubtable       pulumi.StringPtrInput
	GetAllTables              pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	Uplink                    pulumi.IntPtrInput
	UplinkSta                 pulumi.IntPtrInput
	Vdomparam                 pulumi.StringPtrInput
	Wmm                       pulumi.StringPtrInput
	WmmBeDscp                 pulumi.IntPtrInput
	WmmBkDscp                 pulumi.IntPtrInput
	WmmDscpMarking            pulumi.StringPtrInput
	WmmUapsd                  pulumi.StringPtrInput
	WmmViDscp                 pulumi.IntPtrInput
	WmmVoDscp                 pulumi.IntPtrInput
}

func (WirelessControllerQosProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerQosProfileArgs)(nil)).Elem()
}

type WirelessControllerQosProfileInput interface {
	pulumi.Input

	ToWirelessControllerQosProfileOutput() WirelessControllerQosProfileOutput
	ToWirelessControllerQosProfileOutputWithContext(ctx context.Context) WirelessControllerQosProfileOutput
}

func (*WirelessControllerQosProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerQosProfile)(nil)).Elem()
}

func (i *WirelessControllerQosProfile) ToWirelessControllerQosProfileOutput() WirelessControllerQosProfileOutput {
	return i.ToWirelessControllerQosProfileOutputWithContext(context.Background())
}

func (i *WirelessControllerQosProfile) ToWirelessControllerQosProfileOutputWithContext(ctx context.Context) WirelessControllerQosProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerQosProfileOutput)
}

func (i *WirelessControllerQosProfile) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerQosProfile] {
	return pulumix.Output[*WirelessControllerQosProfile]{
		OutputState: i.ToWirelessControllerQosProfileOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerQosProfileArrayInput is an input type that accepts WirelessControllerQosProfileArray and WirelessControllerQosProfileArrayOutput values.
// You can construct a concrete instance of `WirelessControllerQosProfileArrayInput` via:
//
//	WirelessControllerQosProfileArray{ WirelessControllerQosProfileArgs{...} }
type WirelessControllerQosProfileArrayInput interface {
	pulumi.Input

	ToWirelessControllerQosProfileArrayOutput() WirelessControllerQosProfileArrayOutput
	ToWirelessControllerQosProfileArrayOutputWithContext(context.Context) WirelessControllerQosProfileArrayOutput
}

type WirelessControllerQosProfileArray []WirelessControllerQosProfileInput

func (WirelessControllerQosProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerQosProfile)(nil)).Elem()
}

func (i WirelessControllerQosProfileArray) ToWirelessControllerQosProfileArrayOutput() WirelessControllerQosProfileArrayOutput {
	return i.ToWirelessControllerQosProfileArrayOutputWithContext(context.Background())
}

func (i WirelessControllerQosProfileArray) ToWirelessControllerQosProfileArrayOutputWithContext(ctx context.Context) WirelessControllerQosProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerQosProfileArrayOutput)
}

func (i WirelessControllerQosProfileArray) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerQosProfile] {
	return pulumix.Output[[]*WirelessControllerQosProfile]{
		OutputState: i.ToWirelessControllerQosProfileArrayOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerQosProfileMapInput is an input type that accepts WirelessControllerQosProfileMap and WirelessControllerQosProfileMapOutput values.
// You can construct a concrete instance of `WirelessControllerQosProfileMapInput` via:
//
//	WirelessControllerQosProfileMap{ "key": WirelessControllerQosProfileArgs{...} }
type WirelessControllerQosProfileMapInput interface {
	pulumi.Input

	ToWirelessControllerQosProfileMapOutput() WirelessControllerQosProfileMapOutput
	ToWirelessControllerQosProfileMapOutputWithContext(context.Context) WirelessControllerQosProfileMapOutput
}

type WirelessControllerQosProfileMap map[string]WirelessControllerQosProfileInput

func (WirelessControllerQosProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerQosProfile)(nil)).Elem()
}

func (i WirelessControllerQosProfileMap) ToWirelessControllerQosProfileMapOutput() WirelessControllerQosProfileMapOutput {
	return i.ToWirelessControllerQosProfileMapOutputWithContext(context.Background())
}

func (i WirelessControllerQosProfileMap) ToWirelessControllerQosProfileMapOutputWithContext(ctx context.Context) WirelessControllerQosProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerQosProfileMapOutput)
}

func (i WirelessControllerQosProfileMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerQosProfile] {
	return pulumix.Output[map[string]*WirelessControllerQosProfile]{
		OutputState: i.ToWirelessControllerQosProfileMapOutputWithContext(ctx).OutputState,
	}
}

type WirelessControllerQosProfileOutput struct{ *pulumi.OutputState }

func (WirelessControllerQosProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerQosProfile)(nil)).Elem()
}

func (o WirelessControllerQosProfileOutput) ToWirelessControllerQosProfileOutput() WirelessControllerQosProfileOutput {
	return o
}

func (o WirelessControllerQosProfileOutput) ToWirelessControllerQosProfileOutputWithContext(ctx context.Context) WirelessControllerQosProfileOutput {
	return o
}

func (o WirelessControllerQosProfileOutput) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerQosProfile] {
	return pulumix.Output[*WirelessControllerQosProfile]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerQosProfileOutput) BandwidthAdmissionControl() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringOutput { return v.BandwidthAdmissionControl }).(pulumi.StringOutput)
}

func (o WirelessControllerQosProfileOutput) BandwidthCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.IntOutput { return v.BandwidthCapacity }).(pulumi.IntOutput)
}

func (o WirelessControllerQosProfileOutput) Burst() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringOutput { return v.Burst }).(pulumi.StringOutput)
}

func (o WirelessControllerQosProfileOutput) CallAdmissionControl() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringOutput { return v.CallAdmissionControl }).(pulumi.StringOutput)
}

func (o WirelessControllerQosProfileOutput) CallCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.IntOutput { return v.CallCapacity }).(pulumi.IntOutput)
}

func (o WirelessControllerQosProfileOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

func (o WirelessControllerQosProfileOutput) Downlink() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.IntOutput { return v.Downlink }).(pulumi.IntOutput)
}

func (o WirelessControllerQosProfileOutput) DownlinkSta() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.IntOutput { return v.DownlinkSta }).(pulumi.IntOutput)
}

func (o WirelessControllerQosProfileOutput) DscpWmmBes() WirelessControllerQosProfileDscpWmmBeArrayOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) WirelessControllerQosProfileDscpWmmBeArrayOutput {
		return v.DscpWmmBes
	}).(WirelessControllerQosProfileDscpWmmBeArrayOutput)
}

func (o WirelessControllerQosProfileOutput) DscpWmmBks() WirelessControllerQosProfileDscpWmmBkArrayOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) WirelessControllerQosProfileDscpWmmBkArrayOutput {
		return v.DscpWmmBks
	}).(WirelessControllerQosProfileDscpWmmBkArrayOutput)
}

func (o WirelessControllerQosProfileOutput) DscpWmmMapping() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringOutput { return v.DscpWmmMapping }).(pulumi.StringOutput)
}

func (o WirelessControllerQosProfileOutput) DscpWmmVis() WirelessControllerQosProfileDscpWmmViArrayOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) WirelessControllerQosProfileDscpWmmViArrayOutput {
		return v.DscpWmmVis
	}).(WirelessControllerQosProfileDscpWmmViArrayOutput)
}

func (o WirelessControllerQosProfileOutput) DscpWmmVos() WirelessControllerQosProfileDscpWmmVoArrayOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) WirelessControllerQosProfileDscpWmmVoArrayOutput {
		return v.DscpWmmVos
	}).(WirelessControllerQosProfileDscpWmmVoArrayOutput)
}

func (o WirelessControllerQosProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerQosProfileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerQosProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WirelessControllerQosProfileOutput) Uplink() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.IntOutput { return v.Uplink }).(pulumi.IntOutput)
}

func (o WirelessControllerQosProfileOutput) UplinkSta() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.IntOutput { return v.UplinkSta }).(pulumi.IntOutput)
}

func (o WirelessControllerQosProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerQosProfileOutput) Wmm() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringOutput { return v.Wmm }).(pulumi.StringOutput)
}

func (o WirelessControllerQosProfileOutput) WmmBeDscp() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.IntOutput { return v.WmmBeDscp }).(pulumi.IntOutput)
}

func (o WirelessControllerQosProfileOutput) WmmBkDscp() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.IntOutput { return v.WmmBkDscp }).(pulumi.IntOutput)
}

func (o WirelessControllerQosProfileOutput) WmmDscpMarking() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringOutput { return v.WmmDscpMarking }).(pulumi.StringOutput)
}

func (o WirelessControllerQosProfileOutput) WmmUapsd() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.StringOutput { return v.WmmUapsd }).(pulumi.StringOutput)
}

func (o WirelessControllerQosProfileOutput) WmmViDscp() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.IntOutput { return v.WmmViDscp }).(pulumi.IntOutput)
}

func (o WirelessControllerQosProfileOutput) WmmVoDscp() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerQosProfile) pulumi.IntOutput { return v.WmmVoDscp }).(pulumi.IntOutput)
}

type WirelessControllerQosProfileArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerQosProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerQosProfile)(nil)).Elem()
}

func (o WirelessControllerQosProfileArrayOutput) ToWirelessControllerQosProfileArrayOutput() WirelessControllerQosProfileArrayOutput {
	return o
}

func (o WirelessControllerQosProfileArrayOutput) ToWirelessControllerQosProfileArrayOutputWithContext(ctx context.Context) WirelessControllerQosProfileArrayOutput {
	return o
}

func (o WirelessControllerQosProfileArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerQosProfile] {
	return pulumix.Output[[]*WirelessControllerQosProfile]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerQosProfileArrayOutput) Index(i pulumi.IntInput) WirelessControllerQosProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerQosProfile {
		return vs[0].([]*WirelessControllerQosProfile)[vs[1].(int)]
	}).(WirelessControllerQosProfileOutput)
}

type WirelessControllerQosProfileMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerQosProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerQosProfile)(nil)).Elem()
}

func (o WirelessControllerQosProfileMapOutput) ToWirelessControllerQosProfileMapOutput() WirelessControllerQosProfileMapOutput {
	return o
}

func (o WirelessControllerQosProfileMapOutput) ToWirelessControllerQosProfileMapOutputWithContext(ctx context.Context) WirelessControllerQosProfileMapOutput {
	return o
}

func (o WirelessControllerQosProfileMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerQosProfile] {
	return pulumix.Output[map[string]*WirelessControllerQosProfile]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerQosProfileMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerQosProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerQosProfile {
		return vs[0].(map[string]*WirelessControllerQosProfile)[vs[1].(string)]
	}).(WirelessControllerQosProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerQosProfileInput)(nil)).Elem(), &WirelessControllerQosProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerQosProfileArrayInput)(nil)).Elem(), WirelessControllerQosProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerQosProfileMapInput)(nil)).Elem(), WirelessControllerQosProfileMap{})
	pulumi.RegisterOutputType(WirelessControllerQosProfileOutput{})
	pulumi.RegisterOutputType(WirelessControllerQosProfileArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerQosProfileMapOutput{})
}

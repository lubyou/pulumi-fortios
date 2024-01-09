// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SwitchControllerLldpSettings struct {
	pulumi.CustomResourceState

	DeviceDetection     pulumi.StringOutput    `pulumi:"deviceDetection"`
	FastStartInterval   pulumi.IntOutput       `pulumi:"fastStartInterval"`
	ManagementInterface pulumi.StringOutput    `pulumi:"managementInterface"`
	Status              pulumi.StringOutput    `pulumi:"status"`
	TxHold              pulumi.IntOutput       `pulumi:"txHold"`
	TxInterval          pulumi.IntOutput       `pulumi:"txInterval"`
	Vdomparam           pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerLldpSettings registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerLldpSettings(ctx *pulumi.Context,
	name string, args *SwitchControllerLldpSettingsArgs, opts ...pulumi.ResourceOption) (*SwitchControllerLldpSettings, error) {
	if args == nil {
		args = &SwitchControllerLldpSettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwitchControllerLldpSettings
	err := ctx.RegisterResource("fortios:index/switchControllerLldpSettings:SwitchControllerLldpSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerLldpSettings gets an existing SwitchControllerLldpSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerLldpSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerLldpSettingsState, opts ...pulumi.ResourceOption) (*SwitchControllerLldpSettings, error) {
	var resource SwitchControllerLldpSettings
	err := ctx.ReadResource("fortios:index/switchControllerLldpSettings:SwitchControllerLldpSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerLldpSettings resources.
type switchControllerLldpSettingsState struct {
	DeviceDetection     *string `pulumi:"deviceDetection"`
	FastStartInterval   *int    `pulumi:"fastStartInterval"`
	ManagementInterface *string `pulumi:"managementInterface"`
	Status              *string `pulumi:"status"`
	TxHold              *int    `pulumi:"txHold"`
	TxInterval          *int    `pulumi:"txInterval"`
	Vdomparam           *string `pulumi:"vdomparam"`
}

type SwitchControllerLldpSettingsState struct {
	DeviceDetection     pulumi.StringPtrInput
	FastStartInterval   pulumi.IntPtrInput
	ManagementInterface pulumi.StringPtrInput
	Status              pulumi.StringPtrInput
	TxHold              pulumi.IntPtrInput
	TxInterval          pulumi.IntPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SwitchControllerLldpSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerLldpSettingsState)(nil)).Elem()
}

type switchControllerLldpSettingsArgs struct {
	DeviceDetection     *string `pulumi:"deviceDetection"`
	FastStartInterval   *int    `pulumi:"fastStartInterval"`
	ManagementInterface *string `pulumi:"managementInterface"`
	Status              *string `pulumi:"status"`
	TxHold              *int    `pulumi:"txHold"`
	TxInterval          *int    `pulumi:"txInterval"`
	Vdomparam           *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerLldpSettings resource.
type SwitchControllerLldpSettingsArgs struct {
	DeviceDetection     pulumi.StringPtrInput
	FastStartInterval   pulumi.IntPtrInput
	ManagementInterface pulumi.StringPtrInput
	Status              pulumi.StringPtrInput
	TxHold              pulumi.IntPtrInput
	TxInterval          pulumi.IntPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SwitchControllerLldpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerLldpSettingsArgs)(nil)).Elem()
}

type SwitchControllerLldpSettingsInput interface {
	pulumi.Input

	ToSwitchControllerLldpSettingsOutput() SwitchControllerLldpSettingsOutput
	ToSwitchControllerLldpSettingsOutputWithContext(ctx context.Context) SwitchControllerLldpSettingsOutput
}

func (*SwitchControllerLldpSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerLldpSettings)(nil)).Elem()
}

func (i *SwitchControllerLldpSettings) ToSwitchControllerLldpSettingsOutput() SwitchControllerLldpSettingsOutput {
	return i.ToSwitchControllerLldpSettingsOutputWithContext(context.Background())
}

func (i *SwitchControllerLldpSettings) ToSwitchControllerLldpSettingsOutputWithContext(ctx context.Context) SwitchControllerLldpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerLldpSettingsOutput)
}

// SwitchControllerLldpSettingsArrayInput is an input type that accepts SwitchControllerLldpSettingsArray and SwitchControllerLldpSettingsArrayOutput values.
// You can construct a concrete instance of `SwitchControllerLldpSettingsArrayInput` via:
//
//	SwitchControllerLldpSettingsArray{ SwitchControllerLldpSettingsArgs{...} }
type SwitchControllerLldpSettingsArrayInput interface {
	pulumi.Input

	ToSwitchControllerLldpSettingsArrayOutput() SwitchControllerLldpSettingsArrayOutput
	ToSwitchControllerLldpSettingsArrayOutputWithContext(context.Context) SwitchControllerLldpSettingsArrayOutput
}

type SwitchControllerLldpSettingsArray []SwitchControllerLldpSettingsInput

func (SwitchControllerLldpSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerLldpSettings)(nil)).Elem()
}

func (i SwitchControllerLldpSettingsArray) ToSwitchControllerLldpSettingsArrayOutput() SwitchControllerLldpSettingsArrayOutput {
	return i.ToSwitchControllerLldpSettingsArrayOutputWithContext(context.Background())
}

func (i SwitchControllerLldpSettingsArray) ToSwitchControllerLldpSettingsArrayOutputWithContext(ctx context.Context) SwitchControllerLldpSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerLldpSettingsArrayOutput)
}

// SwitchControllerLldpSettingsMapInput is an input type that accepts SwitchControllerLldpSettingsMap and SwitchControllerLldpSettingsMapOutput values.
// You can construct a concrete instance of `SwitchControllerLldpSettingsMapInput` via:
//
//	SwitchControllerLldpSettingsMap{ "key": SwitchControllerLldpSettingsArgs{...} }
type SwitchControllerLldpSettingsMapInput interface {
	pulumi.Input

	ToSwitchControllerLldpSettingsMapOutput() SwitchControllerLldpSettingsMapOutput
	ToSwitchControllerLldpSettingsMapOutputWithContext(context.Context) SwitchControllerLldpSettingsMapOutput
}

type SwitchControllerLldpSettingsMap map[string]SwitchControllerLldpSettingsInput

func (SwitchControllerLldpSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerLldpSettings)(nil)).Elem()
}

func (i SwitchControllerLldpSettingsMap) ToSwitchControllerLldpSettingsMapOutput() SwitchControllerLldpSettingsMapOutput {
	return i.ToSwitchControllerLldpSettingsMapOutputWithContext(context.Background())
}

func (i SwitchControllerLldpSettingsMap) ToSwitchControllerLldpSettingsMapOutputWithContext(ctx context.Context) SwitchControllerLldpSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerLldpSettingsMapOutput)
}

type SwitchControllerLldpSettingsOutput struct{ *pulumi.OutputState }

func (SwitchControllerLldpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerLldpSettings)(nil)).Elem()
}

func (o SwitchControllerLldpSettingsOutput) ToSwitchControllerLldpSettingsOutput() SwitchControllerLldpSettingsOutput {
	return o
}

func (o SwitchControllerLldpSettingsOutput) ToSwitchControllerLldpSettingsOutputWithContext(ctx context.Context) SwitchControllerLldpSettingsOutput {
	return o
}

func (o SwitchControllerLldpSettingsOutput) DeviceDetection() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpSettings) pulumi.StringOutput { return v.DeviceDetection }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpSettingsOutput) FastStartInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerLldpSettings) pulumi.IntOutput { return v.FastStartInterval }).(pulumi.IntOutput)
}

func (o SwitchControllerLldpSettingsOutput) ManagementInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpSettings) pulumi.StringOutput { return v.ManagementInterface }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpSettingsOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpSettings) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpSettingsOutput) TxHold() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerLldpSettings) pulumi.IntOutput { return v.TxHold }).(pulumi.IntOutput)
}

func (o SwitchControllerLldpSettingsOutput) TxInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerLldpSettings) pulumi.IntOutput { return v.TxInterval }).(pulumi.IntOutput)
}

func (o SwitchControllerLldpSettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerLldpSettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchControllerLldpSettingsArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerLldpSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerLldpSettings)(nil)).Elem()
}

func (o SwitchControllerLldpSettingsArrayOutput) ToSwitchControllerLldpSettingsArrayOutput() SwitchControllerLldpSettingsArrayOutput {
	return o
}

func (o SwitchControllerLldpSettingsArrayOutput) ToSwitchControllerLldpSettingsArrayOutputWithContext(ctx context.Context) SwitchControllerLldpSettingsArrayOutput {
	return o
}

func (o SwitchControllerLldpSettingsArrayOutput) Index(i pulumi.IntInput) SwitchControllerLldpSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerLldpSettings {
		return vs[0].([]*SwitchControllerLldpSettings)[vs[1].(int)]
	}).(SwitchControllerLldpSettingsOutput)
}

type SwitchControllerLldpSettingsMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerLldpSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerLldpSettings)(nil)).Elem()
}

func (o SwitchControllerLldpSettingsMapOutput) ToSwitchControllerLldpSettingsMapOutput() SwitchControllerLldpSettingsMapOutput {
	return o
}

func (o SwitchControllerLldpSettingsMapOutput) ToSwitchControllerLldpSettingsMapOutputWithContext(ctx context.Context) SwitchControllerLldpSettingsMapOutput {
	return o
}

func (o SwitchControllerLldpSettingsMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerLldpSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerLldpSettings {
		return vs[0].(map[string]*SwitchControllerLldpSettings)[vs[1].(string)]
	}).(SwitchControllerLldpSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerLldpSettingsInput)(nil)).Elem(), &SwitchControllerLldpSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerLldpSettingsArrayInput)(nil)).Elem(), SwitchControllerLldpSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerLldpSettingsMapInput)(nil)).Elem(), SwitchControllerLldpSettingsMap{})
	pulumi.RegisterOutputType(SwitchControllerLldpSettingsOutput{})
	pulumi.RegisterOutputType(SwitchControllerLldpSettingsArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerLldpSettingsMapOutput{})
}

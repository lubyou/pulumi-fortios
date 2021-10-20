// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure integrated NAC settings for FortiSwitch. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// SwitchController NacSettings can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerNacSettings:SwitchControllerNacSettings labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerNacSettings struct {
	pulumi.CustomResourceState

	// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
	AutoAuth pulumi.StringOutput `pulumi:"autoAuth"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceNacPort pulumi.StringOutput `pulumi:"bounceNacPort"`
	// Time interval after which inactive NAC devices will be expired (in minutes, 0 means no expiry).
	InactiveTimer pulumi.IntOutput `pulumi:"inactiveTimer"`
	// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush pulumi.StringOutput `pulumi:"linkDownFlush"`
	// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// NAC settings name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Default NAC Onboarding VLAN when NAC devices are discovered.
	OnboardingVlan pulumi.StringOutput `pulumi:"onboardingVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerNacSettings registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerNacSettings(ctx *pulumi.Context,
	name string, args *SwitchControllerNacSettingsArgs, opts ...pulumi.ResourceOption) (*SwitchControllerNacSettings, error) {
	if args == nil {
		args = &SwitchControllerNacSettingsArgs{}
	}

	var resource SwitchControllerNacSettings
	err := ctx.RegisterResource("fortios:index/switchControllerNacSettings:SwitchControllerNacSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerNacSettings gets an existing SwitchControllerNacSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerNacSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerNacSettingsState, opts ...pulumi.ResourceOption) (*SwitchControllerNacSettings, error) {
	var resource SwitchControllerNacSettings
	err := ctx.ReadResource("fortios:index/switchControllerNacSettings:SwitchControllerNacSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerNacSettings resources.
type switchControllerNacSettingsState struct {
	// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
	AutoAuth *string `pulumi:"autoAuth"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceNacPort *string `pulumi:"bounceNacPort"`
	// Time interval after which inactive NAC devices will be expired (in minutes, 0 means no expiry).
	InactiveTimer *int `pulumi:"inactiveTimer"`
	// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush *string `pulumi:"linkDownFlush"`
	// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
	Mode *string `pulumi:"mode"`
	// NAC settings name.
	Name *string `pulumi:"name"`
	// Default NAC Onboarding VLAN when NAC devices are discovered.
	OnboardingVlan *string `pulumi:"onboardingVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerNacSettingsState struct {
	// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
	AutoAuth pulumi.StringPtrInput
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceNacPort pulumi.StringPtrInput
	// Time interval after which inactive NAC devices will be expired (in minutes, 0 means no expiry).
	InactiveTimer pulumi.IntPtrInput
	// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush pulumi.StringPtrInput
	// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
	Mode pulumi.StringPtrInput
	// NAC settings name.
	Name pulumi.StringPtrInput
	// Default NAC Onboarding VLAN when NAC devices are discovered.
	OnboardingVlan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerNacSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerNacSettingsState)(nil)).Elem()
}

type switchControllerNacSettingsArgs struct {
	// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
	AutoAuth *string `pulumi:"autoAuth"`
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceNacPort *string `pulumi:"bounceNacPort"`
	// Time interval after which inactive NAC devices will be expired (in minutes, 0 means no expiry).
	InactiveTimer *int `pulumi:"inactiveTimer"`
	// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush *string `pulumi:"linkDownFlush"`
	// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
	Mode *string `pulumi:"mode"`
	// NAC settings name.
	Name *string `pulumi:"name"`
	// Default NAC Onboarding VLAN when NAC devices are discovered.
	OnboardingVlan *string `pulumi:"onboardingVlan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerNacSettings resource.
type SwitchControllerNacSettingsArgs struct {
	// Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
	AutoAuth pulumi.StringPtrInput
	// Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
	BounceNacPort pulumi.StringPtrInput
	// Time interval after which inactive NAC devices will be expired (in minutes, 0 means no expiry).
	InactiveTimer pulumi.IntPtrInput
	// Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
	LinkDownFlush pulumi.StringPtrInput
	// Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
	Mode pulumi.StringPtrInput
	// NAC settings name.
	Name pulumi.StringPtrInput
	// Default NAC Onboarding VLAN when NAC devices are discovered.
	OnboardingVlan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerNacSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerNacSettingsArgs)(nil)).Elem()
}

type SwitchControllerNacSettingsInput interface {
	pulumi.Input

	ToSwitchControllerNacSettingsOutput() SwitchControllerNacSettingsOutput
	ToSwitchControllerNacSettingsOutputWithContext(ctx context.Context) SwitchControllerNacSettingsOutput
}

func (*SwitchControllerNacSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerNacSettings)(nil))
}

func (i *SwitchControllerNacSettings) ToSwitchControllerNacSettingsOutput() SwitchControllerNacSettingsOutput {
	return i.ToSwitchControllerNacSettingsOutputWithContext(context.Background())
}

func (i *SwitchControllerNacSettings) ToSwitchControllerNacSettingsOutputWithContext(ctx context.Context) SwitchControllerNacSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerNacSettingsOutput)
}

func (i *SwitchControllerNacSettings) ToSwitchControllerNacSettingsPtrOutput() SwitchControllerNacSettingsPtrOutput {
	return i.ToSwitchControllerNacSettingsPtrOutputWithContext(context.Background())
}

func (i *SwitchControllerNacSettings) ToSwitchControllerNacSettingsPtrOutputWithContext(ctx context.Context) SwitchControllerNacSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerNacSettingsPtrOutput)
}

type SwitchControllerNacSettingsPtrInput interface {
	pulumi.Input

	ToSwitchControllerNacSettingsPtrOutput() SwitchControllerNacSettingsPtrOutput
	ToSwitchControllerNacSettingsPtrOutputWithContext(ctx context.Context) SwitchControllerNacSettingsPtrOutput
}

type switchControllerNacSettingsPtrType SwitchControllerNacSettingsArgs

func (*switchControllerNacSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerNacSettings)(nil))
}

func (i *switchControllerNacSettingsPtrType) ToSwitchControllerNacSettingsPtrOutput() SwitchControllerNacSettingsPtrOutput {
	return i.ToSwitchControllerNacSettingsPtrOutputWithContext(context.Background())
}

func (i *switchControllerNacSettingsPtrType) ToSwitchControllerNacSettingsPtrOutputWithContext(ctx context.Context) SwitchControllerNacSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerNacSettingsPtrOutput)
}

// SwitchControllerNacSettingsArrayInput is an input type that accepts SwitchControllerNacSettingsArray and SwitchControllerNacSettingsArrayOutput values.
// You can construct a concrete instance of `SwitchControllerNacSettingsArrayInput` via:
//
//          SwitchControllerNacSettingsArray{ SwitchControllerNacSettingsArgs{...} }
type SwitchControllerNacSettingsArrayInput interface {
	pulumi.Input

	ToSwitchControllerNacSettingsArrayOutput() SwitchControllerNacSettingsArrayOutput
	ToSwitchControllerNacSettingsArrayOutputWithContext(context.Context) SwitchControllerNacSettingsArrayOutput
}

type SwitchControllerNacSettingsArray []SwitchControllerNacSettingsInput

func (SwitchControllerNacSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SwitchControllerNacSettings)(nil))
}

func (i SwitchControllerNacSettingsArray) ToSwitchControllerNacSettingsArrayOutput() SwitchControllerNacSettingsArrayOutput {
	return i.ToSwitchControllerNacSettingsArrayOutputWithContext(context.Background())
}

func (i SwitchControllerNacSettingsArray) ToSwitchControllerNacSettingsArrayOutputWithContext(ctx context.Context) SwitchControllerNacSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerNacSettingsArrayOutput)
}

// SwitchControllerNacSettingsMapInput is an input type that accepts SwitchControllerNacSettingsMap and SwitchControllerNacSettingsMapOutput values.
// You can construct a concrete instance of `SwitchControllerNacSettingsMapInput` via:
//
//          SwitchControllerNacSettingsMap{ "key": SwitchControllerNacSettingsArgs{...} }
type SwitchControllerNacSettingsMapInput interface {
	pulumi.Input

	ToSwitchControllerNacSettingsMapOutput() SwitchControllerNacSettingsMapOutput
	ToSwitchControllerNacSettingsMapOutputWithContext(context.Context) SwitchControllerNacSettingsMapOutput
}

type SwitchControllerNacSettingsMap map[string]SwitchControllerNacSettingsInput

func (SwitchControllerNacSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SwitchControllerNacSettings)(nil))
}

func (i SwitchControllerNacSettingsMap) ToSwitchControllerNacSettingsMapOutput() SwitchControllerNacSettingsMapOutput {
	return i.ToSwitchControllerNacSettingsMapOutputWithContext(context.Background())
}

func (i SwitchControllerNacSettingsMap) ToSwitchControllerNacSettingsMapOutputWithContext(ctx context.Context) SwitchControllerNacSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerNacSettingsMapOutput)
}

type SwitchControllerNacSettingsOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerNacSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerNacSettings)(nil))
}

func (o SwitchControllerNacSettingsOutput) ToSwitchControllerNacSettingsOutput() SwitchControllerNacSettingsOutput {
	return o
}

func (o SwitchControllerNacSettingsOutput) ToSwitchControllerNacSettingsOutputWithContext(ctx context.Context) SwitchControllerNacSettingsOutput {
	return o
}

func (o SwitchControllerNacSettingsOutput) ToSwitchControllerNacSettingsPtrOutput() SwitchControllerNacSettingsPtrOutput {
	return o.ToSwitchControllerNacSettingsPtrOutputWithContext(context.Background())
}

func (o SwitchControllerNacSettingsOutput) ToSwitchControllerNacSettingsPtrOutputWithContext(ctx context.Context) SwitchControllerNacSettingsPtrOutput {
	return o.ApplyT(func(v SwitchControllerNacSettings) *SwitchControllerNacSettings {
		return &v
	}).(SwitchControllerNacSettingsPtrOutput)
}

type SwitchControllerNacSettingsPtrOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerNacSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerNacSettings)(nil))
}

func (o SwitchControllerNacSettingsPtrOutput) ToSwitchControllerNacSettingsPtrOutput() SwitchControllerNacSettingsPtrOutput {
	return o
}

func (o SwitchControllerNacSettingsPtrOutput) ToSwitchControllerNacSettingsPtrOutputWithContext(ctx context.Context) SwitchControllerNacSettingsPtrOutput {
	return o
}

type SwitchControllerNacSettingsArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerNacSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SwitchControllerNacSettings)(nil))
}

func (o SwitchControllerNacSettingsArrayOutput) ToSwitchControllerNacSettingsArrayOutput() SwitchControllerNacSettingsArrayOutput {
	return o
}

func (o SwitchControllerNacSettingsArrayOutput) ToSwitchControllerNacSettingsArrayOutputWithContext(ctx context.Context) SwitchControllerNacSettingsArrayOutput {
	return o
}

func (o SwitchControllerNacSettingsArrayOutput) Index(i pulumi.IntInput) SwitchControllerNacSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SwitchControllerNacSettings {
		return vs[0].([]SwitchControllerNacSettings)[vs[1].(int)]
	}).(SwitchControllerNacSettingsOutput)
}

type SwitchControllerNacSettingsMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerNacSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SwitchControllerNacSettings)(nil))
}

func (o SwitchControllerNacSettingsMapOutput) ToSwitchControllerNacSettingsMapOutput() SwitchControllerNacSettingsMapOutput {
	return o
}

func (o SwitchControllerNacSettingsMapOutput) ToSwitchControllerNacSettingsMapOutputWithContext(ctx context.Context) SwitchControllerNacSettingsMapOutput {
	return o
}

func (o SwitchControllerNacSettingsMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerNacSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SwitchControllerNacSettings {
		return vs[0].(map[string]SwitchControllerNacSettings)[vs[1].(string)]
	}).(SwitchControllerNacSettingsOutput)
}

func init() {
	pulumi.RegisterOutputType(SwitchControllerNacSettingsOutput{})
	pulumi.RegisterOutputType(SwitchControllerNacSettingsPtrOutput{})
	pulumi.RegisterOutputType(SwitchControllerNacSettingsArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerNacSettingsMapOutput{})
}
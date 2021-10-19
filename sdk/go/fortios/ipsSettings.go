// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS VDOM parameter.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewIpsSettings(ctx, "trname", &fortios.IpsSettingsArgs{
// 			IpsPacketQuota:      pulumi.Int(0),
// 			PacketLogHistory:    pulumi.Int(1),
// 			PacketLogMemory:     pulumi.Int(256),
// 			PacketLogPostAttack: pulumi.Int(0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Ips Settings can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/ipsSettings:IpsSettings labelname IpsSettings
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type IpsSettings struct {
	pulumi.CustomResourceState

	// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
	IpsPacketQuota pulumi.IntOutput `pulumi:"ipsPacketQuota"`
	// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
	PacketLogHistory pulumi.IntOutput `pulumi:"packetLogHistory"`
	// Maximum memory can be used by packet log (64 - 8192 kB).
	PacketLogMemory pulumi.IntOutput `pulumi:"packetLogMemory"`
	// Number of packets to log after the IPS signature is detected (0 - 255).
	PacketLogPostAttack pulumi.IntOutput `pulumi:"packetLogPostAttack"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIpsSettings registers a new resource with the given unique name, arguments, and options.
func NewIpsSettings(ctx *pulumi.Context,
	name string, args *IpsSettingsArgs, opts ...pulumi.ResourceOption) (*IpsSettings, error) {
	if args == nil {
		args = &IpsSettingsArgs{}
	}

	var resource IpsSettings
	err := ctx.RegisterResource("fortios:index/ipsSettings:IpsSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsSettings gets an existing IpsSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsSettingsState, opts ...pulumi.ResourceOption) (*IpsSettings, error) {
	var resource IpsSettings
	err := ctx.ReadResource("fortios:index/ipsSettings:IpsSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsSettings resources.
type ipsSettingsState struct {
	// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
	IpsPacketQuota *int `pulumi:"ipsPacketQuota"`
	// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
	PacketLogHistory *int `pulumi:"packetLogHistory"`
	// Maximum memory can be used by packet log (64 - 8192 kB).
	PacketLogMemory *int `pulumi:"packetLogMemory"`
	// Number of packets to log after the IPS signature is detected (0 - 255).
	PacketLogPostAttack *int `pulumi:"packetLogPostAttack"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IpsSettingsState struct {
	// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
	IpsPacketQuota pulumi.IntPtrInput
	// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
	PacketLogHistory pulumi.IntPtrInput
	// Maximum memory can be used by packet log (64 - 8192 kB).
	PacketLogMemory pulumi.IntPtrInput
	// Number of packets to log after the IPS signature is detected (0 - 255).
	PacketLogPostAttack pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsSettingsState)(nil)).Elem()
}

type ipsSettingsArgs struct {
	// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
	IpsPacketQuota *int `pulumi:"ipsPacketQuota"`
	// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
	PacketLogHistory *int `pulumi:"packetLogHistory"`
	// Maximum memory can be used by packet log (64 - 8192 kB).
	PacketLogMemory *int `pulumi:"packetLogMemory"`
	// Number of packets to log after the IPS signature is detected (0 - 255).
	PacketLogPostAttack *int `pulumi:"packetLogPostAttack"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a IpsSettings resource.
type IpsSettingsArgs struct {
	// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
	IpsPacketQuota pulumi.IntPtrInput
	// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
	PacketLogHistory pulumi.IntPtrInput
	// Maximum memory can be used by packet log (64 - 8192 kB).
	PacketLogMemory pulumi.IntPtrInput
	// Number of packets to log after the IPS signature is detected (0 - 255).
	PacketLogPostAttack pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsSettingsArgs)(nil)).Elem()
}

type IpsSettingsInput interface {
	pulumi.Input

	ToIpsSettingsOutput() IpsSettingsOutput
	ToIpsSettingsOutputWithContext(ctx context.Context) IpsSettingsOutput
}

func (*IpsSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsSettings)(nil))
}

func (i *IpsSettings) ToIpsSettingsOutput() IpsSettingsOutput {
	return i.ToIpsSettingsOutputWithContext(context.Background())
}

func (i *IpsSettings) ToIpsSettingsOutputWithContext(ctx context.Context) IpsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSettingsOutput)
}

func (i *IpsSettings) ToIpsSettingsPtrOutput() IpsSettingsPtrOutput {
	return i.ToIpsSettingsPtrOutputWithContext(context.Background())
}

func (i *IpsSettings) ToIpsSettingsPtrOutputWithContext(ctx context.Context) IpsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSettingsPtrOutput)
}

type IpsSettingsPtrInput interface {
	pulumi.Input

	ToIpsSettingsPtrOutput() IpsSettingsPtrOutput
	ToIpsSettingsPtrOutputWithContext(ctx context.Context) IpsSettingsPtrOutput
}

type ipsSettingsPtrType IpsSettingsArgs

func (*ipsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsSettings)(nil))
}

func (i *ipsSettingsPtrType) ToIpsSettingsPtrOutput() IpsSettingsPtrOutput {
	return i.ToIpsSettingsPtrOutputWithContext(context.Background())
}

func (i *ipsSettingsPtrType) ToIpsSettingsPtrOutputWithContext(ctx context.Context) IpsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSettingsPtrOutput)
}

// IpsSettingsArrayInput is an input type that accepts IpsSettingsArray and IpsSettingsArrayOutput values.
// You can construct a concrete instance of `IpsSettingsArrayInput` via:
//
//          IpsSettingsArray{ IpsSettingsArgs{...} }
type IpsSettingsArrayInput interface {
	pulumi.Input

	ToIpsSettingsArrayOutput() IpsSettingsArrayOutput
	ToIpsSettingsArrayOutputWithContext(context.Context) IpsSettingsArrayOutput
}

type IpsSettingsArray []IpsSettingsInput

func (IpsSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IpsSettings)(nil))
}

func (i IpsSettingsArray) ToIpsSettingsArrayOutput() IpsSettingsArrayOutput {
	return i.ToIpsSettingsArrayOutputWithContext(context.Background())
}

func (i IpsSettingsArray) ToIpsSettingsArrayOutputWithContext(ctx context.Context) IpsSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSettingsArrayOutput)
}

// IpsSettingsMapInput is an input type that accepts IpsSettingsMap and IpsSettingsMapOutput values.
// You can construct a concrete instance of `IpsSettingsMapInput` via:
//
//          IpsSettingsMap{ "key": IpsSettingsArgs{...} }
type IpsSettingsMapInput interface {
	pulumi.Input

	ToIpsSettingsMapOutput() IpsSettingsMapOutput
	ToIpsSettingsMapOutputWithContext(context.Context) IpsSettingsMapOutput
}

type IpsSettingsMap map[string]IpsSettingsInput

func (IpsSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IpsSettings)(nil))
}

func (i IpsSettingsMap) ToIpsSettingsMapOutput() IpsSettingsMapOutput {
	return i.ToIpsSettingsMapOutputWithContext(context.Background())
}

func (i IpsSettingsMap) ToIpsSettingsMapOutputWithContext(ctx context.Context) IpsSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSettingsMapOutput)
}

type IpsSettingsOutput struct {
	*pulumi.OutputState
}

func (IpsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsSettings)(nil))
}

func (o IpsSettingsOutput) ToIpsSettingsOutput() IpsSettingsOutput {
	return o
}

func (o IpsSettingsOutput) ToIpsSettingsOutputWithContext(ctx context.Context) IpsSettingsOutput {
	return o
}

func (o IpsSettingsOutput) ToIpsSettingsPtrOutput() IpsSettingsPtrOutput {
	return o.ToIpsSettingsPtrOutputWithContext(context.Background())
}

func (o IpsSettingsOutput) ToIpsSettingsPtrOutputWithContext(ctx context.Context) IpsSettingsPtrOutput {
	return o.ApplyT(func(v IpsSettings) *IpsSettings {
		return &v
	}).(IpsSettingsPtrOutput)
}

type IpsSettingsPtrOutput struct {
	*pulumi.OutputState
}

func (IpsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsSettings)(nil))
}

func (o IpsSettingsPtrOutput) ToIpsSettingsPtrOutput() IpsSettingsPtrOutput {
	return o
}

func (o IpsSettingsPtrOutput) ToIpsSettingsPtrOutputWithContext(ctx context.Context) IpsSettingsPtrOutput {
	return o
}

type IpsSettingsArrayOutput struct{ *pulumi.OutputState }

func (IpsSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpsSettings)(nil))
}

func (o IpsSettingsArrayOutput) ToIpsSettingsArrayOutput() IpsSettingsArrayOutput {
	return o
}

func (o IpsSettingsArrayOutput) ToIpsSettingsArrayOutputWithContext(ctx context.Context) IpsSettingsArrayOutput {
	return o
}

func (o IpsSettingsArrayOutput) Index(i pulumi.IntInput) IpsSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpsSettings {
		return vs[0].([]IpsSettings)[vs[1].(int)]
	}).(IpsSettingsOutput)
}

type IpsSettingsMapOutput struct{ *pulumi.OutputState }

func (IpsSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IpsSettings)(nil))
}

func (o IpsSettingsMapOutput) ToIpsSettingsMapOutput() IpsSettingsMapOutput {
	return o
}

func (o IpsSettingsMapOutput) ToIpsSettingsMapOutputWithContext(ctx context.Context) IpsSettingsMapOutput {
	return o
}

func (o IpsSettingsMapOutput) MapIndex(k pulumi.StringInput) IpsSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IpsSettings {
		return vs[0].(map[string]IpsSettings)[vs[1].(string)]
	}).(IpsSettingsOutput)
}

func init() {
	pulumi.RegisterOutputType(IpsSettingsOutput{})
	pulumi.RegisterOutputType(IpsSettingsPtrOutput{})
	pulumi.RegisterOutputType(IpsSettingsArrayOutput{})
	pulumi.RegisterOutputType(IpsSettingsMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure endpoint control settings.
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
// 		_, err := fortios.NewEndpointControlSettings(ctx, "trname", &fortios.EndpointControlSettingsArgs{
// 			DownloadLocation:                  pulumi.String("fortiguard"),
// 			ForticlientAvdbUpdateInterval:     pulumi.Int(8),
// 			ForticlientDeregUnsupportedClient: pulumi.String("enable"),
// 			ForticlientEmsRestApiCallTimeout:  pulumi.Int(5000),
// 			ForticlientKeepaliveInterval:      pulumi.Int(60),
// 			ForticlientOfflineGrace:           pulumi.String("disable"),
// 			ForticlientOfflineGraceInterval:   pulumi.Int(120),
// 			ForticlientRegKeyEnforce:          pulumi.String("disable"),
// 			ForticlientRegTimeout:             pulumi.Int(7),
// 			ForticlientSysUpdateInterval:      pulumi.Int(720),
// 			ForticlientUserAvatar:             pulumi.String("enable"),
// 			ForticlientWarningInterval:        pulumi.Int(1),
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
// EndpointControl Settings can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/endpointControlSettings:EndpointControlSettings labelname EndpointControlSettings
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type EndpointControlSettings struct {
	pulumi.CustomResourceState

	// Customized URL for downloading FortiClient.
	DownloadCustomLink pulumi.StringOutput `pulumi:"downloadCustomLink"`
	// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
	DownloadLocation pulumi.StringOutput `pulumi:"downloadLocation"`
	// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
	ForticlientAvdbUpdateInterval pulumi.IntOutput `pulumi:"forticlientAvdbUpdateInterval"`
	// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDeregUnsupportedClient pulumi.StringOutput `pulumi:"forticlientDeregUnsupportedClient"`
	// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDisconnectUnsupportedClient pulumi.StringOutput `pulumi:"forticlientDisconnectUnsupportedClient"`
	// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
	ForticlientEmsRestApiCallTimeout pulumi.IntOutput `pulumi:"forticlientEmsRestApiCallTimeout"`
	// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
	ForticlientKeepaliveInterval pulumi.IntOutput `pulumi:"forticlientKeepaliveInterval"`
	// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
	ForticlientOfflineGrace pulumi.StringOutput `pulumi:"forticlientOfflineGrace"`
	// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
	ForticlientOfflineGraceInterval pulumi.IntOutput `pulumi:"forticlientOfflineGraceInterval"`
	// FortiClient registration key.
	ForticlientRegKey pulumi.StringPtrOutput `pulumi:"forticlientRegKey"`
	// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
	ForticlientRegKeyEnforce pulumi.StringOutput `pulumi:"forticlientRegKeyEnforce"`
	// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
	ForticlientRegTimeout pulumi.IntOutput `pulumi:"forticlientRegTimeout"`
	// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
	ForticlientSysUpdateInterval pulumi.IntOutput `pulumi:"forticlientSysUpdateInterval"`
	// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
	ForticlientUserAvatar pulumi.StringOutput `pulumi:"forticlientUserAvatar"`
	// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
	ForticlientWarningInterval pulumi.IntOutput `pulumi:"forticlientWarningInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewEndpointControlSettings registers a new resource with the given unique name, arguments, and options.
func NewEndpointControlSettings(ctx *pulumi.Context,
	name string, args *EndpointControlSettingsArgs, opts ...pulumi.ResourceOption) (*EndpointControlSettings, error) {
	if args == nil {
		args = &EndpointControlSettingsArgs{}
	}

	var resource EndpointControlSettings
	err := ctx.RegisterResource("fortios:index/endpointControlSettings:EndpointControlSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointControlSettings gets an existing EndpointControlSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointControlSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointControlSettingsState, opts ...pulumi.ResourceOption) (*EndpointControlSettings, error) {
	var resource EndpointControlSettings
	err := ctx.ReadResource("fortios:index/endpointControlSettings:EndpointControlSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointControlSettings resources.
type endpointControlSettingsState struct {
	// Customized URL for downloading FortiClient.
	DownloadCustomLink *string `pulumi:"downloadCustomLink"`
	// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
	DownloadLocation *string `pulumi:"downloadLocation"`
	// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
	ForticlientAvdbUpdateInterval *int `pulumi:"forticlientAvdbUpdateInterval"`
	// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDeregUnsupportedClient *string `pulumi:"forticlientDeregUnsupportedClient"`
	// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDisconnectUnsupportedClient *string `pulumi:"forticlientDisconnectUnsupportedClient"`
	// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
	ForticlientEmsRestApiCallTimeout *int `pulumi:"forticlientEmsRestApiCallTimeout"`
	// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
	ForticlientKeepaliveInterval *int `pulumi:"forticlientKeepaliveInterval"`
	// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
	ForticlientOfflineGrace *string `pulumi:"forticlientOfflineGrace"`
	// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
	ForticlientOfflineGraceInterval *int `pulumi:"forticlientOfflineGraceInterval"`
	// FortiClient registration key.
	ForticlientRegKey *string `pulumi:"forticlientRegKey"`
	// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
	ForticlientRegKeyEnforce *string `pulumi:"forticlientRegKeyEnforce"`
	// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
	ForticlientRegTimeout *int `pulumi:"forticlientRegTimeout"`
	// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
	ForticlientSysUpdateInterval *int `pulumi:"forticlientSysUpdateInterval"`
	// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
	ForticlientUserAvatar *string `pulumi:"forticlientUserAvatar"`
	// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
	ForticlientWarningInterval *int `pulumi:"forticlientWarningInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type EndpointControlSettingsState struct {
	// Customized URL for downloading FortiClient.
	DownloadCustomLink pulumi.StringPtrInput
	// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
	DownloadLocation pulumi.StringPtrInput
	// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
	ForticlientAvdbUpdateInterval pulumi.IntPtrInput
	// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDeregUnsupportedClient pulumi.StringPtrInput
	// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDisconnectUnsupportedClient pulumi.StringPtrInput
	// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
	ForticlientEmsRestApiCallTimeout pulumi.IntPtrInput
	// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
	ForticlientKeepaliveInterval pulumi.IntPtrInput
	// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
	ForticlientOfflineGrace pulumi.StringPtrInput
	// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
	ForticlientOfflineGraceInterval pulumi.IntPtrInput
	// FortiClient registration key.
	ForticlientRegKey pulumi.StringPtrInput
	// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
	ForticlientRegKeyEnforce pulumi.StringPtrInput
	// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
	ForticlientRegTimeout pulumi.IntPtrInput
	// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
	ForticlientSysUpdateInterval pulumi.IntPtrInput
	// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
	ForticlientUserAvatar pulumi.StringPtrInput
	// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
	ForticlientWarningInterval pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EndpointControlSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlSettingsState)(nil)).Elem()
}

type endpointControlSettingsArgs struct {
	// Customized URL for downloading FortiClient.
	DownloadCustomLink *string `pulumi:"downloadCustomLink"`
	// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
	DownloadLocation *string `pulumi:"downloadLocation"`
	// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
	ForticlientAvdbUpdateInterval *int `pulumi:"forticlientAvdbUpdateInterval"`
	// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDeregUnsupportedClient *string `pulumi:"forticlientDeregUnsupportedClient"`
	// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDisconnectUnsupportedClient *string `pulumi:"forticlientDisconnectUnsupportedClient"`
	// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
	ForticlientEmsRestApiCallTimeout *int `pulumi:"forticlientEmsRestApiCallTimeout"`
	// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
	ForticlientKeepaliveInterval *int `pulumi:"forticlientKeepaliveInterval"`
	// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
	ForticlientOfflineGrace *string `pulumi:"forticlientOfflineGrace"`
	// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
	ForticlientOfflineGraceInterval *int `pulumi:"forticlientOfflineGraceInterval"`
	// FortiClient registration key.
	ForticlientRegKey *string `pulumi:"forticlientRegKey"`
	// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
	ForticlientRegKeyEnforce *string `pulumi:"forticlientRegKeyEnforce"`
	// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
	ForticlientRegTimeout *int `pulumi:"forticlientRegTimeout"`
	// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
	ForticlientSysUpdateInterval *int `pulumi:"forticlientSysUpdateInterval"`
	// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
	ForticlientUserAvatar *string `pulumi:"forticlientUserAvatar"`
	// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
	ForticlientWarningInterval *int `pulumi:"forticlientWarningInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EndpointControlSettings resource.
type EndpointControlSettingsArgs struct {
	// Customized URL for downloading FortiClient.
	DownloadCustomLink pulumi.StringPtrInput
	// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
	DownloadLocation pulumi.StringPtrInput
	// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
	ForticlientAvdbUpdateInterval pulumi.IntPtrInput
	// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDeregUnsupportedClient pulumi.StringPtrInput
	// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
	ForticlientDisconnectUnsupportedClient pulumi.StringPtrInput
	// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
	ForticlientEmsRestApiCallTimeout pulumi.IntPtrInput
	// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
	ForticlientKeepaliveInterval pulumi.IntPtrInput
	// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
	ForticlientOfflineGrace pulumi.StringPtrInput
	// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
	ForticlientOfflineGraceInterval pulumi.IntPtrInput
	// FortiClient registration key.
	ForticlientRegKey pulumi.StringPtrInput
	// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
	ForticlientRegKeyEnforce pulumi.StringPtrInput
	// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
	ForticlientRegTimeout pulumi.IntPtrInput
	// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
	ForticlientSysUpdateInterval pulumi.IntPtrInput
	// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
	ForticlientUserAvatar pulumi.StringPtrInput
	// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
	ForticlientWarningInterval pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EndpointControlSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlSettingsArgs)(nil)).Elem()
}

type EndpointControlSettingsInput interface {
	pulumi.Input

	ToEndpointControlSettingsOutput() EndpointControlSettingsOutput
	ToEndpointControlSettingsOutputWithContext(ctx context.Context) EndpointControlSettingsOutput
}

func (*EndpointControlSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointControlSettings)(nil))
}

func (i *EndpointControlSettings) ToEndpointControlSettingsOutput() EndpointControlSettingsOutput {
	return i.ToEndpointControlSettingsOutputWithContext(context.Background())
}

func (i *EndpointControlSettings) ToEndpointControlSettingsOutputWithContext(ctx context.Context) EndpointControlSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlSettingsOutput)
}

func (i *EndpointControlSettings) ToEndpointControlSettingsPtrOutput() EndpointControlSettingsPtrOutput {
	return i.ToEndpointControlSettingsPtrOutputWithContext(context.Background())
}

func (i *EndpointControlSettings) ToEndpointControlSettingsPtrOutputWithContext(ctx context.Context) EndpointControlSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlSettingsPtrOutput)
}

type EndpointControlSettingsPtrInput interface {
	pulumi.Input

	ToEndpointControlSettingsPtrOutput() EndpointControlSettingsPtrOutput
	ToEndpointControlSettingsPtrOutputWithContext(ctx context.Context) EndpointControlSettingsPtrOutput
}

type endpointControlSettingsPtrType EndpointControlSettingsArgs

func (*endpointControlSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointControlSettings)(nil))
}

func (i *endpointControlSettingsPtrType) ToEndpointControlSettingsPtrOutput() EndpointControlSettingsPtrOutput {
	return i.ToEndpointControlSettingsPtrOutputWithContext(context.Background())
}

func (i *endpointControlSettingsPtrType) ToEndpointControlSettingsPtrOutputWithContext(ctx context.Context) EndpointControlSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlSettingsPtrOutput)
}

// EndpointControlSettingsArrayInput is an input type that accepts EndpointControlSettingsArray and EndpointControlSettingsArrayOutput values.
// You can construct a concrete instance of `EndpointControlSettingsArrayInput` via:
//
//          EndpointControlSettingsArray{ EndpointControlSettingsArgs{...} }
type EndpointControlSettingsArrayInput interface {
	pulumi.Input

	ToEndpointControlSettingsArrayOutput() EndpointControlSettingsArrayOutput
	ToEndpointControlSettingsArrayOutputWithContext(context.Context) EndpointControlSettingsArrayOutput
}

type EndpointControlSettingsArray []EndpointControlSettingsInput

func (EndpointControlSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EndpointControlSettings)(nil))
}

func (i EndpointControlSettingsArray) ToEndpointControlSettingsArrayOutput() EndpointControlSettingsArrayOutput {
	return i.ToEndpointControlSettingsArrayOutputWithContext(context.Background())
}

func (i EndpointControlSettingsArray) ToEndpointControlSettingsArrayOutputWithContext(ctx context.Context) EndpointControlSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlSettingsArrayOutput)
}

// EndpointControlSettingsMapInput is an input type that accepts EndpointControlSettingsMap and EndpointControlSettingsMapOutput values.
// You can construct a concrete instance of `EndpointControlSettingsMapInput` via:
//
//          EndpointControlSettingsMap{ "key": EndpointControlSettingsArgs{...} }
type EndpointControlSettingsMapInput interface {
	pulumi.Input

	ToEndpointControlSettingsMapOutput() EndpointControlSettingsMapOutput
	ToEndpointControlSettingsMapOutputWithContext(context.Context) EndpointControlSettingsMapOutput
}

type EndpointControlSettingsMap map[string]EndpointControlSettingsInput

func (EndpointControlSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EndpointControlSettings)(nil))
}

func (i EndpointControlSettingsMap) ToEndpointControlSettingsMapOutput() EndpointControlSettingsMapOutput {
	return i.ToEndpointControlSettingsMapOutputWithContext(context.Background())
}

func (i EndpointControlSettingsMap) ToEndpointControlSettingsMapOutputWithContext(ctx context.Context) EndpointControlSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlSettingsMapOutput)
}

type EndpointControlSettingsOutput struct {
	*pulumi.OutputState
}

func (EndpointControlSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointControlSettings)(nil))
}

func (o EndpointControlSettingsOutput) ToEndpointControlSettingsOutput() EndpointControlSettingsOutput {
	return o
}

func (o EndpointControlSettingsOutput) ToEndpointControlSettingsOutputWithContext(ctx context.Context) EndpointControlSettingsOutput {
	return o
}

func (o EndpointControlSettingsOutput) ToEndpointControlSettingsPtrOutput() EndpointControlSettingsPtrOutput {
	return o.ToEndpointControlSettingsPtrOutputWithContext(context.Background())
}

func (o EndpointControlSettingsOutput) ToEndpointControlSettingsPtrOutputWithContext(ctx context.Context) EndpointControlSettingsPtrOutput {
	return o.ApplyT(func(v EndpointControlSettings) *EndpointControlSettings {
		return &v
	}).(EndpointControlSettingsPtrOutput)
}

type EndpointControlSettingsPtrOutput struct {
	*pulumi.OutputState
}

func (EndpointControlSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointControlSettings)(nil))
}

func (o EndpointControlSettingsPtrOutput) ToEndpointControlSettingsPtrOutput() EndpointControlSettingsPtrOutput {
	return o
}

func (o EndpointControlSettingsPtrOutput) ToEndpointControlSettingsPtrOutputWithContext(ctx context.Context) EndpointControlSettingsPtrOutput {
	return o
}

type EndpointControlSettingsArrayOutput struct{ *pulumi.OutputState }

func (EndpointControlSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointControlSettings)(nil))
}

func (o EndpointControlSettingsArrayOutput) ToEndpointControlSettingsArrayOutput() EndpointControlSettingsArrayOutput {
	return o
}

func (o EndpointControlSettingsArrayOutput) ToEndpointControlSettingsArrayOutputWithContext(ctx context.Context) EndpointControlSettingsArrayOutput {
	return o
}

func (o EndpointControlSettingsArrayOutput) Index(i pulumi.IntInput) EndpointControlSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointControlSettings {
		return vs[0].([]EndpointControlSettings)[vs[1].(int)]
	}).(EndpointControlSettingsOutput)
}

type EndpointControlSettingsMapOutput struct{ *pulumi.OutputState }

func (EndpointControlSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EndpointControlSettings)(nil))
}

func (o EndpointControlSettingsMapOutput) ToEndpointControlSettingsMapOutput() EndpointControlSettingsMapOutput {
	return o
}

func (o EndpointControlSettingsMapOutput) ToEndpointControlSettingsMapOutputWithContext(ctx context.Context) EndpointControlSettingsMapOutput {
	return o
}

func (o EndpointControlSettingsMapOutput) MapIndex(k pulumi.StringInput) EndpointControlSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EndpointControlSettings {
		return vs[0].(map[string]EndpointControlSettings)[vs[1].(string)]
	}).(EndpointControlSettingsOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointControlSettingsOutput{})
	pulumi.RegisterOutputType(EndpointControlSettingsPtrOutput{})
	pulumi.RegisterOutputType(EndpointControlSettingsArrayOutput{})
	pulumi.RegisterOutputType(EndpointControlSettingsMapOutput{})
}

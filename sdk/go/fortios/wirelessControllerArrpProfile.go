// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles. Applies to FortiOS Version `>= 6.4.2`.
//
// ## Import
//
// WirelessController ArrpProfile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerArrpProfile:WirelessControllerArrpProfile labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerArrpProfile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
	IncludeDfsChannel pulumi.StringOutput `pulumi:"includeDfsChannel"`
	// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
	IncludeWeatherChannel pulumi.StringOutput `pulumi:"includeWeatherChannel"`
	// Period in seconds to measure average transmit retries and receive errors (default = 300).
	MonitorPeriod pulumi.IntOutput `pulumi:"monitorPeriod"`
	// WiFi ARRP profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
	SelectionPeriod pulumi.IntOutput `pulumi:"selectionPeriod"`
	// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
	ThresholdAp pulumi.IntOutput `pulumi:"thresholdAp"`
	// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
	ThresholdChannelLoad pulumi.IntOutput `pulumi:"thresholdChannelLoad"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
	ThresholdNoiseFloor pulumi.StringOutput `pulumi:"thresholdNoiseFloor"`
	// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
	ThresholdRxErrors pulumi.IntOutput `pulumi:"thresholdRxErrors"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
	ThresholdSpectralRssi pulumi.StringOutput `pulumi:"thresholdSpectralRssi"`
	// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
	ThresholdTxRetries pulumi.IntOutput `pulumi:"thresholdTxRetries"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
	WeightChannelLoad pulumi.IntOutput `pulumi:"weightChannelLoad"`
	// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
	WeightDfsChannel pulumi.IntOutput `pulumi:"weightDfsChannel"`
	// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
	WeightManagedAp pulumi.IntOutput `pulumi:"weightManagedAp"`
	// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
	WeightNoiseFloor pulumi.IntOutput `pulumi:"weightNoiseFloor"`
	// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
	WeightRogueAp pulumi.IntOutput `pulumi:"weightRogueAp"`
	// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
	WeightSpectralRssi pulumi.IntOutput `pulumi:"weightSpectralRssi"`
	// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
	WeightWeatherChannel pulumi.IntOutput `pulumi:"weightWeatherChannel"`
}

// NewWirelessControllerArrpProfile registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerArrpProfile(ctx *pulumi.Context,
	name string, args *WirelessControllerArrpProfileArgs, opts ...pulumi.ResourceOption) (*WirelessControllerArrpProfile, error) {
	if args == nil {
		args = &WirelessControllerArrpProfileArgs{}
	}

	var resource WirelessControllerArrpProfile
	err := ctx.RegisterResource("fortios:index/wirelessControllerArrpProfile:WirelessControllerArrpProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerArrpProfile gets an existing WirelessControllerArrpProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerArrpProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerArrpProfileState, opts ...pulumi.ResourceOption) (*WirelessControllerArrpProfile, error) {
	var resource WirelessControllerArrpProfile
	err := ctx.ReadResource("fortios:index/wirelessControllerArrpProfile:WirelessControllerArrpProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerArrpProfile resources.
type wirelessControllerArrpProfileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
	IncludeDfsChannel *string `pulumi:"includeDfsChannel"`
	// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
	IncludeWeatherChannel *string `pulumi:"includeWeatherChannel"`
	// Period in seconds to measure average transmit retries and receive errors (default = 300).
	MonitorPeriod *int `pulumi:"monitorPeriod"`
	// WiFi ARRP profile name.
	Name *string `pulumi:"name"`
	// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
	SelectionPeriod *int `pulumi:"selectionPeriod"`
	// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
	ThresholdAp *int `pulumi:"thresholdAp"`
	// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
	ThresholdChannelLoad *int `pulumi:"thresholdChannelLoad"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
	ThresholdNoiseFloor *string `pulumi:"thresholdNoiseFloor"`
	// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
	ThresholdRxErrors *int `pulumi:"thresholdRxErrors"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
	ThresholdSpectralRssi *string `pulumi:"thresholdSpectralRssi"`
	// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
	ThresholdTxRetries *int `pulumi:"thresholdTxRetries"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
	WeightChannelLoad *int `pulumi:"weightChannelLoad"`
	// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
	WeightDfsChannel *int `pulumi:"weightDfsChannel"`
	// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
	WeightManagedAp *int `pulumi:"weightManagedAp"`
	// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
	WeightNoiseFloor *int `pulumi:"weightNoiseFloor"`
	// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
	WeightRogueAp *int `pulumi:"weightRogueAp"`
	// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
	WeightSpectralRssi *int `pulumi:"weightSpectralRssi"`
	// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
	WeightWeatherChannel *int `pulumi:"weightWeatherChannel"`
}

type WirelessControllerArrpProfileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
	IncludeDfsChannel pulumi.StringPtrInput
	// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
	IncludeWeatherChannel pulumi.StringPtrInput
	// Period in seconds to measure average transmit retries and receive errors (default = 300).
	MonitorPeriod pulumi.IntPtrInput
	// WiFi ARRP profile name.
	Name pulumi.StringPtrInput
	// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
	SelectionPeriod pulumi.IntPtrInput
	// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
	ThresholdAp pulumi.IntPtrInput
	// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
	ThresholdChannelLoad pulumi.IntPtrInput
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
	ThresholdNoiseFloor pulumi.StringPtrInput
	// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
	ThresholdRxErrors pulumi.IntPtrInput
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
	ThresholdSpectralRssi pulumi.StringPtrInput
	// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
	ThresholdTxRetries pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
	WeightChannelLoad pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
	WeightDfsChannel pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
	WeightManagedAp pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
	WeightNoiseFloor pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
	WeightRogueAp pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
	WeightSpectralRssi pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
	WeightWeatherChannel pulumi.IntPtrInput
}

func (WirelessControllerArrpProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerArrpProfileState)(nil)).Elem()
}

type wirelessControllerArrpProfileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
	IncludeDfsChannel *string `pulumi:"includeDfsChannel"`
	// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
	IncludeWeatherChannel *string `pulumi:"includeWeatherChannel"`
	// Period in seconds to measure average transmit retries and receive errors (default = 300).
	MonitorPeriod *int `pulumi:"monitorPeriod"`
	// WiFi ARRP profile name.
	Name *string `pulumi:"name"`
	// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
	SelectionPeriod *int `pulumi:"selectionPeriod"`
	// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
	ThresholdAp *int `pulumi:"thresholdAp"`
	// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
	ThresholdChannelLoad *int `pulumi:"thresholdChannelLoad"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
	ThresholdNoiseFloor *string `pulumi:"thresholdNoiseFloor"`
	// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
	ThresholdRxErrors *int `pulumi:"thresholdRxErrors"`
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
	ThresholdSpectralRssi *string `pulumi:"thresholdSpectralRssi"`
	// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
	ThresholdTxRetries *int `pulumi:"thresholdTxRetries"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
	WeightChannelLoad *int `pulumi:"weightChannelLoad"`
	// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
	WeightDfsChannel *int `pulumi:"weightDfsChannel"`
	// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
	WeightManagedAp *int `pulumi:"weightManagedAp"`
	// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
	WeightNoiseFloor *int `pulumi:"weightNoiseFloor"`
	// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
	WeightRogueAp *int `pulumi:"weightRogueAp"`
	// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
	WeightSpectralRssi *int `pulumi:"weightSpectralRssi"`
	// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
	WeightWeatherChannel *int `pulumi:"weightWeatherChannel"`
}

// The set of arguments for constructing a WirelessControllerArrpProfile resource.
type WirelessControllerArrpProfileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
	IncludeDfsChannel pulumi.StringPtrInput
	// Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
	IncludeWeatherChannel pulumi.StringPtrInput
	// Period in seconds to measure average transmit retries and receive errors (default = 300).
	MonitorPeriod pulumi.IntPtrInput
	// WiFi ARRP profile name.
	Name pulumi.StringPtrInput
	// Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
	SelectionPeriod pulumi.IntPtrInput
	// Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
	ThresholdAp pulumi.IntPtrInput
	// Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
	ThresholdChannelLoad pulumi.IntPtrInput
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
	ThresholdNoiseFloor pulumi.StringPtrInput
	// Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
	ThresholdRxErrors pulumi.IntPtrInput
	// Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
	ThresholdSpectralRssi pulumi.StringPtrInput
	// Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
	ThresholdTxRetries pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
	WeightChannelLoad pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
	WeightDfsChannel pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
	WeightManagedAp pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
	WeightNoiseFloor pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
	WeightRogueAp pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
	WeightSpectralRssi pulumi.IntPtrInput
	// Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
	WeightWeatherChannel pulumi.IntPtrInput
}

func (WirelessControllerArrpProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerArrpProfileArgs)(nil)).Elem()
}

type WirelessControllerArrpProfileInput interface {
	pulumi.Input

	ToWirelessControllerArrpProfileOutput() WirelessControllerArrpProfileOutput
	ToWirelessControllerArrpProfileOutputWithContext(ctx context.Context) WirelessControllerArrpProfileOutput
}

func (*WirelessControllerArrpProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerArrpProfile)(nil))
}

func (i *WirelessControllerArrpProfile) ToWirelessControllerArrpProfileOutput() WirelessControllerArrpProfileOutput {
	return i.ToWirelessControllerArrpProfileOutputWithContext(context.Background())
}

func (i *WirelessControllerArrpProfile) ToWirelessControllerArrpProfileOutputWithContext(ctx context.Context) WirelessControllerArrpProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerArrpProfileOutput)
}

func (i *WirelessControllerArrpProfile) ToWirelessControllerArrpProfilePtrOutput() WirelessControllerArrpProfilePtrOutput {
	return i.ToWirelessControllerArrpProfilePtrOutputWithContext(context.Background())
}

func (i *WirelessControllerArrpProfile) ToWirelessControllerArrpProfilePtrOutputWithContext(ctx context.Context) WirelessControllerArrpProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerArrpProfilePtrOutput)
}

type WirelessControllerArrpProfilePtrInput interface {
	pulumi.Input

	ToWirelessControllerArrpProfilePtrOutput() WirelessControllerArrpProfilePtrOutput
	ToWirelessControllerArrpProfilePtrOutputWithContext(ctx context.Context) WirelessControllerArrpProfilePtrOutput
}

type wirelessControllerArrpProfilePtrType WirelessControllerArrpProfileArgs

func (*wirelessControllerArrpProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerArrpProfile)(nil))
}

func (i *wirelessControllerArrpProfilePtrType) ToWirelessControllerArrpProfilePtrOutput() WirelessControllerArrpProfilePtrOutput {
	return i.ToWirelessControllerArrpProfilePtrOutputWithContext(context.Background())
}

func (i *wirelessControllerArrpProfilePtrType) ToWirelessControllerArrpProfilePtrOutputWithContext(ctx context.Context) WirelessControllerArrpProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerArrpProfilePtrOutput)
}

// WirelessControllerArrpProfileArrayInput is an input type that accepts WirelessControllerArrpProfileArray and WirelessControllerArrpProfileArrayOutput values.
// You can construct a concrete instance of `WirelessControllerArrpProfileArrayInput` via:
//
//          WirelessControllerArrpProfileArray{ WirelessControllerArrpProfileArgs{...} }
type WirelessControllerArrpProfileArrayInput interface {
	pulumi.Input

	ToWirelessControllerArrpProfileArrayOutput() WirelessControllerArrpProfileArrayOutput
	ToWirelessControllerArrpProfileArrayOutputWithContext(context.Context) WirelessControllerArrpProfileArrayOutput
}

type WirelessControllerArrpProfileArray []WirelessControllerArrpProfileInput

func (WirelessControllerArrpProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WirelessControllerArrpProfile)(nil))
}

func (i WirelessControllerArrpProfileArray) ToWirelessControllerArrpProfileArrayOutput() WirelessControllerArrpProfileArrayOutput {
	return i.ToWirelessControllerArrpProfileArrayOutputWithContext(context.Background())
}

func (i WirelessControllerArrpProfileArray) ToWirelessControllerArrpProfileArrayOutputWithContext(ctx context.Context) WirelessControllerArrpProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerArrpProfileArrayOutput)
}

// WirelessControllerArrpProfileMapInput is an input type that accepts WirelessControllerArrpProfileMap and WirelessControllerArrpProfileMapOutput values.
// You can construct a concrete instance of `WirelessControllerArrpProfileMapInput` via:
//
//          WirelessControllerArrpProfileMap{ "key": WirelessControllerArrpProfileArgs{...} }
type WirelessControllerArrpProfileMapInput interface {
	pulumi.Input

	ToWirelessControllerArrpProfileMapOutput() WirelessControllerArrpProfileMapOutput
	ToWirelessControllerArrpProfileMapOutputWithContext(context.Context) WirelessControllerArrpProfileMapOutput
}

type WirelessControllerArrpProfileMap map[string]WirelessControllerArrpProfileInput

func (WirelessControllerArrpProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WirelessControllerArrpProfile)(nil))
}

func (i WirelessControllerArrpProfileMap) ToWirelessControllerArrpProfileMapOutput() WirelessControllerArrpProfileMapOutput {
	return i.ToWirelessControllerArrpProfileMapOutputWithContext(context.Background())
}

func (i WirelessControllerArrpProfileMap) ToWirelessControllerArrpProfileMapOutputWithContext(ctx context.Context) WirelessControllerArrpProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerArrpProfileMapOutput)
}

type WirelessControllerArrpProfileOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerArrpProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerArrpProfile)(nil))
}

func (o WirelessControllerArrpProfileOutput) ToWirelessControllerArrpProfileOutput() WirelessControllerArrpProfileOutput {
	return o
}

func (o WirelessControllerArrpProfileOutput) ToWirelessControllerArrpProfileOutputWithContext(ctx context.Context) WirelessControllerArrpProfileOutput {
	return o
}

func (o WirelessControllerArrpProfileOutput) ToWirelessControllerArrpProfilePtrOutput() WirelessControllerArrpProfilePtrOutput {
	return o.ToWirelessControllerArrpProfilePtrOutputWithContext(context.Background())
}

func (o WirelessControllerArrpProfileOutput) ToWirelessControllerArrpProfilePtrOutputWithContext(ctx context.Context) WirelessControllerArrpProfilePtrOutput {
	return o.ApplyT(func(v WirelessControllerArrpProfile) *WirelessControllerArrpProfile {
		return &v
	}).(WirelessControllerArrpProfilePtrOutput)
}

type WirelessControllerArrpProfilePtrOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerArrpProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerArrpProfile)(nil))
}

func (o WirelessControllerArrpProfilePtrOutput) ToWirelessControllerArrpProfilePtrOutput() WirelessControllerArrpProfilePtrOutput {
	return o
}

func (o WirelessControllerArrpProfilePtrOutput) ToWirelessControllerArrpProfilePtrOutputWithContext(ctx context.Context) WirelessControllerArrpProfilePtrOutput {
	return o
}

type WirelessControllerArrpProfileArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerArrpProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WirelessControllerArrpProfile)(nil))
}

func (o WirelessControllerArrpProfileArrayOutput) ToWirelessControllerArrpProfileArrayOutput() WirelessControllerArrpProfileArrayOutput {
	return o
}

func (o WirelessControllerArrpProfileArrayOutput) ToWirelessControllerArrpProfileArrayOutputWithContext(ctx context.Context) WirelessControllerArrpProfileArrayOutput {
	return o
}

func (o WirelessControllerArrpProfileArrayOutput) Index(i pulumi.IntInput) WirelessControllerArrpProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WirelessControllerArrpProfile {
		return vs[0].([]WirelessControllerArrpProfile)[vs[1].(int)]
	}).(WirelessControllerArrpProfileOutput)
}

type WirelessControllerArrpProfileMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerArrpProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WirelessControllerArrpProfile)(nil))
}

func (o WirelessControllerArrpProfileMapOutput) ToWirelessControllerArrpProfileMapOutput() WirelessControllerArrpProfileMapOutput {
	return o
}

func (o WirelessControllerArrpProfileMapOutput) ToWirelessControllerArrpProfileMapOutputWithContext(ctx context.Context) WirelessControllerArrpProfileMapOutput {
	return o
}

func (o WirelessControllerArrpProfileMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerArrpProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WirelessControllerArrpProfile {
		return vs[0].(map[string]WirelessControllerArrpProfile)[vs[1].(string)]
	}).(WirelessControllerArrpProfileOutput)
}

func init() {
	pulumi.RegisterOutputType(WirelessControllerArrpProfileOutput{})
	pulumi.RegisterOutputType(WirelessControllerArrpProfilePtrOutput{})
	pulumi.RegisterOutputType(WirelessControllerArrpProfileArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerArrpProfileMapOutput{})
}

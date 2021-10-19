// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure access point status (rogue | accepted | suppressed).
//
// ## Import
//
// WirelessController ApStatus can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerApStatus:WirelessControllerApStatus labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerApStatus struct {
	pulumi.CustomResourceState

	// Access Point's (AP's) BSSID.
	Bssid pulumi.StringOutput `pulumi:"bssid"`
	// AP ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Access Point's (AP's) SSID.
	Ssid pulumi.StringOutput `pulumi:"ssid"`
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerApStatus registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerApStatus(ctx *pulumi.Context,
	name string, args *WirelessControllerApStatusArgs, opts ...pulumi.ResourceOption) (*WirelessControllerApStatus, error) {
	if args == nil {
		args = &WirelessControllerApStatusArgs{}
	}

	var resource WirelessControllerApStatus
	err := ctx.RegisterResource("fortios:index/wirelessControllerApStatus:WirelessControllerApStatus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerApStatus gets an existing WirelessControllerApStatus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerApStatus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerApStatusState, opts ...pulumi.ResourceOption) (*WirelessControllerApStatus, error) {
	var resource WirelessControllerApStatus
	err := ctx.ReadResource("fortios:index/wirelessControllerApStatus:WirelessControllerApStatus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerApStatus resources.
type wirelessControllerApStatusState struct {
	// Access Point's (AP's) BSSID.
	Bssid *string `pulumi:"bssid"`
	// AP ID.
	Fosid *int `pulumi:"fosid"`
	// Access Point's (AP's) SSID.
	Ssid *string `pulumi:"ssid"`
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerApStatusState struct {
	// Access Point's (AP's) BSSID.
	Bssid pulumi.StringPtrInput
	// AP ID.
	Fosid pulumi.IntPtrInput
	// Access Point's (AP's) SSID.
	Ssid pulumi.StringPtrInput
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerApStatusState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerApStatusState)(nil)).Elem()
}

type wirelessControllerApStatusArgs struct {
	// Access Point's (AP's) BSSID.
	Bssid *string `pulumi:"bssid"`
	// AP ID.
	Fosid *int `pulumi:"fosid"`
	// Access Point's (AP's) SSID.
	Ssid *string `pulumi:"ssid"`
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerApStatus resource.
type WirelessControllerApStatusArgs struct {
	// Access Point's (AP's) BSSID.
	Bssid pulumi.StringPtrInput
	// AP ID.
	Fosid pulumi.IntPtrInput
	// Access Point's (AP's) SSID.
	Ssid pulumi.StringPtrInput
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerApStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerApStatusArgs)(nil)).Elem()
}

type WirelessControllerApStatusInput interface {
	pulumi.Input

	ToWirelessControllerApStatusOutput() WirelessControllerApStatusOutput
	ToWirelessControllerApStatusOutputWithContext(ctx context.Context) WirelessControllerApStatusOutput
}

func (*WirelessControllerApStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerApStatus)(nil))
}

func (i *WirelessControllerApStatus) ToWirelessControllerApStatusOutput() WirelessControllerApStatusOutput {
	return i.ToWirelessControllerApStatusOutputWithContext(context.Background())
}

func (i *WirelessControllerApStatus) ToWirelessControllerApStatusOutputWithContext(ctx context.Context) WirelessControllerApStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerApStatusOutput)
}

func (i *WirelessControllerApStatus) ToWirelessControllerApStatusPtrOutput() WirelessControllerApStatusPtrOutput {
	return i.ToWirelessControllerApStatusPtrOutputWithContext(context.Background())
}

func (i *WirelessControllerApStatus) ToWirelessControllerApStatusPtrOutputWithContext(ctx context.Context) WirelessControllerApStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerApStatusPtrOutput)
}

type WirelessControllerApStatusPtrInput interface {
	pulumi.Input

	ToWirelessControllerApStatusPtrOutput() WirelessControllerApStatusPtrOutput
	ToWirelessControllerApStatusPtrOutputWithContext(ctx context.Context) WirelessControllerApStatusPtrOutput
}

type wirelessControllerApStatusPtrType WirelessControllerApStatusArgs

func (*wirelessControllerApStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerApStatus)(nil))
}

func (i *wirelessControllerApStatusPtrType) ToWirelessControllerApStatusPtrOutput() WirelessControllerApStatusPtrOutput {
	return i.ToWirelessControllerApStatusPtrOutputWithContext(context.Background())
}

func (i *wirelessControllerApStatusPtrType) ToWirelessControllerApStatusPtrOutputWithContext(ctx context.Context) WirelessControllerApStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerApStatusPtrOutput)
}

// WirelessControllerApStatusArrayInput is an input type that accepts WirelessControllerApStatusArray and WirelessControllerApStatusArrayOutput values.
// You can construct a concrete instance of `WirelessControllerApStatusArrayInput` via:
//
//          WirelessControllerApStatusArray{ WirelessControllerApStatusArgs{...} }
type WirelessControllerApStatusArrayInput interface {
	pulumi.Input

	ToWirelessControllerApStatusArrayOutput() WirelessControllerApStatusArrayOutput
	ToWirelessControllerApStatusArrayOutputWithContext(context.Context) WirelessControllerApStatusArrayOutput
}

type WirelessControllerApStatusArray []WirelessControllerApStatusInput

func (WirelessControllerApStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WirelessControllerApStatus)(nil))
}

func (i WirelessControllerApStatusArray) ToWirelessControllerApStatusArrayOutput() WirelessControllerApStatusArrayOutput {
	return i.ToWirelessControllerApStatusArrayOutputWithContext(context.Background())
}

func (i WirelessControllerApStatusArray) ToWirelessControllerApStatusArrayOutputWithContext(ctx context.Context) WirelessControllerApStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerApStatusArrayOutput)
}

// WirelessControllerApStatusMapInput is an input type that accepts WirelessControllerApStatusMap and WirelessControllerApStatusMapOutput values.
// You can construct a concrete instance of `WirelessControllerApStatusMapInput` via:
//
//          WirelessControllerApStatusMap{ "key": WirelessControllerApStatusArgs{...} }
type WirelessControllerApStatusMapInput interface {
	pulumi.Input

	ToWirelessControllerApStatusMapOutput() WirelessControllerApStatusMapOutput
	ToWirelessControllerApStatusMapOutputWithContext(context.Context) WirelessControllerApStatusMapOutput
}

type WirelessControllerApStatusMap map[string]WirelessControllerApStatusInput

func (WirelessControllerApStatusMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WirelessControllerApStatus)(nil))
}

func (i WirelessControllerApStatusMap) ToWirelessControllerApStatusMapOutput() WirelessControllerApStatusMapOutput {
	return i.ToWirelessControllerApStatusMapOutputWithContext(context.Background())
}

func (i WirelessControllerApStatusMap) ToWirelessControllerApStatusMapOutputWithContext(ctx context.Context) WirelessControllerApStatusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerApStatusMapOutput)
}

type WirelessControllerApStatusOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerApStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerApStatus)(nil))
}

func (o WirelessControllerApStatusOutput) ToWirelessControllerApStatusOutput() WirelessControllerApStatusOutput {
	return o
}

func (o WirelessControllerApStatusOutput) ToWirelessControllerApStatusOutputWithContext(ctx context.Context) WirelessControllerApStatusOutput {
	return o
}

func (o WirelessControllerApStatusOutput) ToWirelessControllerApStatusPtrOutput() WirelessControllerApStatusPtrOutput {
	return o.ToWirelessControllerApStatusPtrOutputWithContext(context.Background())
}

func (o WirelessControllerApStatusOutput) ToWirelessControllerApStatusPtrOutputWithContext(ctx context.Context) WirelessControllerApStatusPtrOutput {
	return o.ApplyT(func(v WirelessControllerApStatus) *WirelessControllerApStatus {
		return &v
	}).(WirelessControllerApStatusPtrOutput)
}

type WirelessControllerApStatusPtrOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerApStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerApStatus)(nil))
}

func (o WirelessControllerApStatusPtrOutput) ToWirelessControllerApStatusPtrOutput() WirelessControllerApStatusPtrOutput {
	return o
}

func (o WirelessControllerApStatusPtrOutput) ToWirelessControllerApStatusPtrOutputWithContext(ctx context.Context) WirelessControllerApStatusPtrOutput {
	return o
}

type WirelessControllerApStatusArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerApStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WirelessControllerApStatus)(nil))
}

func (o WirelessControllerApStatusArrayOutput) ToWirelessControllerApStatusArrayOutput() WirelessControllerApStatusArrayOutput {
	return o
}

func (o WirelessControllerApStatusArrayOutput) ToWirelessControllerApStatusArrayOutputWithContext(ctx context.Context) WirelessControllerApStatusArrayOutput {
	return o
}

func (o WirelessControllerApStatusArrayOutput) Index(i pulumi.IntInput) WirelessControllerApStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WirelessControllerApStatus {
		return vs[0].([]WirelessControllerApStatus)[vs[1].(int)]
	}).(WirelessControllerApStatusOutput)
}

type WirelessControllerApStatusMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerApStatusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WirelessControllerApStatus)(nil))
}

func (o WirelessControllerApStatusMapOutput) ToWirelessControllerApStatusMapOutput() WirelessControllerApStatusMapOutput {
	return o
}

func (o WirelessControllerApStatusMapOutput) ToWirelessControllerApStatusMapOutputWithContext(ctx context.Context) WirelessControllerApStatusMapOutput {
	return o
}

func (o WirelessControllerApStatusMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerApStatusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WirelessControllerApStatus {
		return vs[0].(map[string]WirelessControllerApStatus)[vs[1].(string)]
	}).(WirelessControllerApStatusOutput)
}

func init() {
	pulumi.RegisterOutputType(WirelessControllerApStatusOutput{})
	pulumi.RegisterOutputType(WirelessControllerApStatusPtrOutput{})
	pulumi.RegisterOutputType(WirelessControllerApStatusArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerApStatusMapOutput{})
}

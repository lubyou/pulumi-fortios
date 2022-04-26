// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure AP local configuration profiles. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// WirelessController ApcfgProfile can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerApcfgProfile:WirelessControllerApcfgProfile labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerApcfgProfile:WirelessControllerApcfgProfile labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerApcfgProfile struct {
	pulumi.CustomResourceState

	// IP address of the validation controller that AP must be able to join after applying AP local configuration.
	AcIp pulumi.StringOutput `pulumi:"acIp"`
	// Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
	AcPort pulumi.IntOutput `pulumi:"acPort"`
	// Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
	AcTimer pulumi.IntOutput `pulumi:"acTimer"`
	// Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
	AcType pulumi.StringOutput `pulumi:"acType"`
	// FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.
	ApFamily pulumi.StringOutput `pulumi:"apFamily"`
	// AP local configuration command list. The structure of `commandList` block is documented below.
	CommandLists WirelessControllerApcfgProfileCommandListArrayOutput `pulumi:"commandLists"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// AP local configuration command name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerApcfgProfile registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerApcfgProfile(ctx *pulumi.Context,
	name string, args *WirelessControllerApcfgProfileArgs, opts ...pulumi.ResourceOption) (*WirelessControllerApcfgProfile, error) {
	if args == nil {
		args = &WirelessControllerApcfgProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelessControllerApcfgProfile
	err := ctx.RegisterResource("fortios:index/wirelessControllerApcfgProfile:WirelessControllerApcfgProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerApcfgProfile gets an existing WirelessControllerApcfgProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerApcfgProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerApcfgProfileState, opts ...pulumi.ResourceOption) (*WirelessControllerApcfgProfile, error) {
	var resource WirelessControllerApcfgProfile
	err := ctx.ReadResource("fortios:index/wirelessControllerApcfgProfile:WirelessControllerApcfgProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerApcfgProfile resources.
type wirelessControllerApcfgProfileState struct {
	// IP address of the validation controller that AP must be able to join after applying AP local configuration.
	AcIp *string `pulumi:"acIp"`
	// Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
	AcPort *int `pulumi:"acPort"`
	// Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
	AcTimer *int `pulumi:"acTimer"`
	// Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
	AcType *string `pulumi:"acType"`
	// FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.
	ApFamily *string `pulumi:"apFamily"`
	// AP local configuration command list. The structure of `commandList` block is documented below.
	CommandLists []WirelessControllerApcfgProfileCommandList `pulumi:"commandLists"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AP local configuration command name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerApcfgProfileState struct {
	// IP address of the validation controller that AP must be able to join after applying AP local configuration.
	AcIp pulumi.StringPtrInput
	// Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
	AcPort pulumi.IntPtrInput
	// Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
	AcTimer pulumi.IntPtrInput
	// Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
	AcType pulumi.StringPtrInput
	// FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.
	ApFamily pulumi.StringPtrInput
	// AP local configuration command list. The structure of `commandList` block is documented below.
	CommandLists WirelessControllerApcfgProfileCommandListArrayInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// AP local configuration command name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerApcfgProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerApcfgProfileState)(nil)).Elem()
}

type wirelessControllerApcfgProfileArgs struct {
	// IP address of the validation controller that AP must be able to join after applying AP local configuration.
	AcIp *string `pulumi:"acIp"`
	// Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
	AcPort *int `pulumi:"acPort"`
	// Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
	AcTimer *int `pulumi:"acTimer"`
	// Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
	AcType *string `pulumi:"acType"`
	// FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.
	ApFamily *string `pulumi:"apFamily"`
	// AP local configuration command list. The structure of `commandList` block is documented below.
	CommandLists []WirelessControllerApcfgProfileCommandList `pulumi:"commandLists"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AP local configuration command name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerApcfgProfile resource.
type WirelessControllerApcfgProfileArgs struct {
	// IP address of the validation controller that AP must be able to join after applying AP local configuration.
	AcIp pulumi.StringPtrInput
	// Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
	AcPort pulumi.IntPtrInput
	// Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
	AcTimer pulumi.IntPtrInput
	// Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
	AcType pulumi.StringPtrInput
	// FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.
	ApFamily pulumi.StringPtrInput
	// AP local configuration command list. The structure of `commandList` block is documented below.
	CommandLists WirelessControllerApcfgProfileCommandListArrayInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// AP local configuration command name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerApcfgProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerApcfgProfileArgs)(nil)).Elem()
}

type WirelessControllerApcfgProfileInput interface {
	pulumi.Input

	ToWirelessControllerApcfgProfileOutput() WirelessControllerApcfgProfileOutput
	ToWirelessControllerApcfgProfileOutputWithContext(ctx context.Context) WirelessControllerApcfgProfileOutput
}

func (*WirelessControllerApcfgProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerApcfgProfile)(nil)).Elem()
}

func (i *WirelessControllerApcfgProfile) ToWirelessControllerApcfgProfileOutput() WirelessControllerApcfgProfileOutput {
	return i.ToWirelessControllerApcfgProfileOutputWithContext(context.Background())
}

func (i *WirelessControllerApcfgProfile) ToWirelessControllerApcfgProfileOutputWithContext(ctx context.Context) WirelessControllerApcfgProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerApcfgProfileOutput)
}

// WirelessControllerApcfgProfileArrayInput is an input type that accepts WirelessControllerApcfgProfileArray and WirelessControllerApcfgProfileArrayOutput values.
// You can construct a concrete instance of `WirelessControllerApcfgProfileArrayInput` via:
//
//          WirelessControllerApcfgProfileArray{ WirelessControllerApcfgProfileArgs{...} }
type WirelessControllerApcfgProfileArrayInput interface {
	pulumi.Input

	ToWirelessControllerApcfgProfileArrayOutput() WirelessControllerApcfgProfileArrayOutput
	ToWirelessControllerApcfgProfileArrayOutputWithContext(context.Context) WirelessControllerApcfgProfileArrayOutput
}

type WirelessControllerApcfgProfileArray []WirelessControllerApcfgProfileInput

func (WirelessControllerApcfgProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerApcfgProfile)(nil)).Elem()
}

func (i WirelessControllerApcfgProfileArray) ToWirelessControllerApcfgProfileArrayOutput() WirelessControllerApcfgProfileArrayOutput {
	return i.ToWirelessControllerApcfgProfileArrayOutputWithContext(context.Background())
}

func (i WirelessControllerApcfgProfileArray) ToWirelessControllerApcfgProfileArrayOutputWithContext(ctx context.Context) WirelessControllerApcfgProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerApcfgProfileArrayOutput)
}

// WirelessControllerApcfgProfileMapInput is an input type that accepts WirelessControllerApcfgProfileMap and WirelessControllerApcfgProfileMapOutput values.
// You can construct a concrete instance of `WirelessControllerApcfgProfileMapInput` via:
//
//          WirelessControllerApcfgProfileMap{ "key": WirelessControllerApcfgProfileArgs{...} }
type WirelessControllerApcfgProfileMapInput interface {
	pulumi.Input

	ToWirelessControllerApcfgProfileMapOutput() WirelessControllerApcfgProfileMapOutput
	ToWirelessControllerApcfgProfileMapOutputWithContext(context.Context) WirelessControllerApcfgProfileMapOutput
}

type WirelessControllerApcfgProfileMap map[string]WirelessControllerApcfgProfileInput

func (WirelessControllerApcfgProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerApcfgProfile)(nil)).Elem()
}

func (i WirelessControllerApcfgProfileMap) ToWirelessControllerApcfgProfileMapOutput() WirelessControllerApcfgProfileMapOutput {
	return i.ToWirelessControllerApcfgProfileMapOutputWithContext(context.Background())
}

func (i WirelessControllerApcfgProfileMap) ToWirelessControllerApcfgProfileMapOutputWithContext(ctx context.Context) WirelessControllerApcfgProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerApcfgProfileMapOutput)
}

type WirelessControllerApcfgProfileOutput struct{ *pulumi.OutputState }

func (WirelessControllerApcfgProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerApcfgProfile)(nil)).Elem()
}

func (o WirelessControllerApcfgProfileOutput) ToWirelessControllerApcfgProfileOutput() WirelessControllerApcfgProfileOutput {
	return o
}

func (o WirelessControllerApcfgProfileOutput) ToWirelessControllerApcfgProfileOutputWithContext(ctx context.Context) WirelessControllerApcfgProfileOutput {
	return o
}

type WirelessControllerApcfgProfileArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerApcfgProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerApcfgProfile)(nil)).Elem()
}

func (o WirelessControllerApcfgProfileArrayOutput) ToWirelessControllerApcfgProfileArrayOutput() WirelessControllerApcfgProfileArrayOutput {
	return o
}

func (o WirelessControllerApcfgProfileArrayOutput) ToWirelessControllerApcfgProfileArrayOutputWithContext(ctx context.Context) WirelessControllerApcfgProfileArrayOutput {
	return o
}

func (o WirelessControllerApcfgProfileArrayOutput) Index(i pulumi.IntInput) WirelessControllerApcfgProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerApcfgProfile {
		return vs[0].([]*WirelessControllerApcfgProfile)[vs[1].(int)]
	}).(WirelessControllerApcfgProfileOutput)
}

type WirelessControllerApcfgProfileMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerApcfgProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerApcfgProfile)(nil)).Elem()
}

func (o WirelessControllerApcfgProfileMapOutput) ToWirelessControllerApcfgProfileMapOutput() WirelessControllerApcfgProfileMapOutput {
	return o
}

func (o WirelessControllerApcfgProfileMapOutput) ToWirelessControllerApcfgProfileMapOutputWithContext(ctx context.Context) WirelessControllerApcfgProfileMapOutput {
	return o
}

func (o WirelessControllerApcfgProfileMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerApcfgProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerApcfgProfile {
		return vs[0].(map[string]*WirelessControllerApcfgProfile)[vs[1].(string)]
	}).(WirelessControllerApcfgProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerApcfgProfileInput)(nil)).Elem(), &WirelessControllerApcfgProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerApcfgProfileArrayInput)(nil)).Elem(), WirelessControllerApcfgProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerApcfgProfileMapInput)(nil)).Elem(), WirelessControllerApcfgProfileMap{})
	pulumi.RegisterOutputType(WirelessControllerApcfgProfileOutput{})
	pulumi.RegisterOutputType(WirelessControllerApcfgProfileArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerApcfgProfileMapOutput{})
}

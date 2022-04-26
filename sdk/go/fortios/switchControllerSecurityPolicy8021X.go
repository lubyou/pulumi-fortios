// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure 802.1x MAC Authentication Bypass (MAB) policies.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewSwitchControllerSecurityPolicy8021X(ctx, "trname", &fortios.SwitchControllerSecurityPolicy8021XArgs{
// 			AuthFailVlan:           pulumi.String("disable"),
// 			AuthFailVlanid:         pulumi.Int(0),
// 			EapPassthru:            pulumi.String("disable"),
// 			FramevidApply:          pulumi.String("enable"),
// 			GuestAuthDelay:         pulumi.Int(30),
// 			GuestVlan:              pulumi.String("disable"),
// 			GuestVlanid:            pulumi.Int(100),
// 			MacAuthBypass:          pulumi.String("disable"),
// 			OpenAuth:               pulumi.String("disable"),
// 			PolicyType:             pulumi.String("802.1X"),
// 			RadiusTimeoutOverwrite: pulumi.String("disable"),
// 			SecurityMode:           pulumi.String("802.1X"),
// 			UserGroups: SwitchControllerSecurityPolicy8021XUserGroupArray{
// 				&SwitchControllerSecurityPolicy8021XUserGroupArgs{
// 					Name: pulumi.String("Guest-group"),
// 				},
// 			},
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
// SwitchControllerSecurityPolicy 8021X can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/switchControllerSecurityPolicy8021X:SwitchControllerSecurityPolicy8021X labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerSecurityPolicy8021X:SwitchControllerSecurityPolicy8021X labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerSecurityPolicy8021X struct {
	pulumi.CustomResourceState

	// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
	AuthFailVlan pulumi.StringOutput `pulumi:"authFailVlan"`
	// VLAN ID on which authentication failed.
	AuthFailVlanId pulumi.StringOutput `pulumi:"authFailVlanId"`
	// VLAN ID on which authentication failed.
	AuthFailVlanid pulumi.IntOutput `pulumi:"authFailVlanid"`
	// Authentication server timeout period (3 - 15 sec, default = 3).
	AuthserverTimeoutPeriod pulumi.IntOutput `pulumi:"authserverTimeoutPeriod"`
	// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
	AuthserverTimeoutVlan pulumi.StringOutput `pulumi:"authserverTimeoutVlan"`
	// Authentication server timeout VLAN name.
	AuthserverTimeoutVlanid pulumi.StringOutput `pulumi:"authserverTimeoutVlanid"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
	EapAutoUntaggedVlans pulumi.StringOutput `pulumi:"eapAutoUntaggedVlans"`
	// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
	EapPassthru pulumi.StringOutput `pulumi:"eapPassthru"`
	// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
	FramevidApply pulumi.StringOutput `pulumi:"framevidApply"`
	// Guest authentication delay (1 - 900  sec, default = 30).
	GuestAuthDelay pulumi.IntOutput `pulumi:"guestAuthDelay"`
	// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
	GuestVlan pulumi.StringOutput `pulumi:"guestVlan"`
	// Guest VLAN name.
	GuestVlanId pulumi.StringOutput `pulumi:"guestVlanId"`
	// Guest VLAN ID.
	GuestVlanid pulumi.IntOutput `pulumi:"guestVlanid"`
	// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
	MacAuthBypass pulumi.StringOutput `pulumi:"macAuthBypass"`
	// Group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
	OpenAuth pulumi.StringOutput `pulumi:"openAuth"`
	// Policy type. Valid values: `802.1X`.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
	// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
	RadiusTimeoutOverwrite pulumi.StringOutput `pulumi:"radiusTimeoutOverwrite"`
	// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
	SecurityMode pulumi.StringOutput `pulumi:"securityMode"`
	// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
	UserGroups SwitchControllerSecurityPolicy8021XUserGroupArrayOutput `pulumi:"userGroups"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerSecurityPolicy8021X registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerSecurityPolicy8021X(ctx *pulumi.Context,
	name string, args *SwitchControllerSecurityPolicy8021XArgs, opts ...pulumi.ResourceOption) (*SwitchControllerSecurityPolicy8021X, error) {
	if args == nil {
		args = &SwitchControllerSecurityPolicy8021XArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerSecurityPolicy8021X
	err := ctx.RegisterResource("fortios:index/switchControllerSecurityPolicy8021X:SwitchControllerSecurityPolicy8021X", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerSecurityPolicy8021X gets an existing SwitchControllerSecurityPolicy8021X resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerSecurityPolicy8021X(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerSecurityPolicy8021XState, opts ...pulumi.ResourceOption) (*SwitchControllerSecurityPolicy8021X, error) {
	var resource SwitchControllerSecurityPolicy8021X
	err := ctx.ReadResource("fortios:index/switchControllerSecurityPolicy8021X:SwitchControllerSecurityPolicy8021X", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerSecurityPolicy8021X resources.
type switchControllerSecurityPolicy8021XState struct {
	// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
	AuthFailVlan *string `pulumi:"authFailVlan"`
	// VLAN ID on which authentication failed.
	AuthFailVlanId *string `pulumi:"authFailVlanId"`
	// VLAN ID on which authentication failed.
	AuthFailVlanid *int `pulumi:"authFailVlanid"`
	// Authentication server timeout period (3 - 15 sec, default = 3).
	AuthserverTimeoutPeriod *int `pulumi:"authserverTimeoutPeriod"`
	// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
	AuthserverTimeoutVlan *string `pulumi:"authserverTimeoutVlan"`
	// Authentication server timeout VLAN name.
	AuthserverTimeoutVlanid *string `pulumi:"authserverTimeoutVlanid"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
	EapAutoUntaggedVlans *string `pulumi:"eapAutoUntaggedVlans"`
	// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
	EapPassthru *string `pulumi:"eapPassthru"`
	// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
	FramevidApply *string `pulumi:"framevidApply"`
	// Guest authentication delay (1 - 900  sec, default = 30).
	GuestAuthDelay *int `pulumi:"guestAuthDelay"`
	// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
	GuestVlan *string `pulumi:"guestVlan"`
	// Guest VLAN name.
	GuestVlanId *string `pulumi:"guestVlanId"`
	// Guest VLAN ID.
	GuestVlanid *int `pulumi:"guestVlanid"`
	// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
	MacAuthBypass *string `pulumi:"macAuthBypass"`
	// Group name.
	Name *string `pulumi:"name"`
	// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
	OpenAuth *string `pulumi:"openAuth"`
	// Policy type. Valid values: `802.1X`.
	PolicyType *string `pulumi:"policyType"`
	// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
	RadiusTimeoutOverwrite *string `pulumi:"radiusTimeoutOverwrite"`
	// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
	SecurityMode *string `pulumi:"securityMode"`
	// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
	UserGroups []SwitchControllerSecurityPolicy8021XUserGroup `pulumi:"userGroups"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerSecurityPolicy8021XState struct {
	// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
	AuthFailVlan pulumi.StringPtrInput
	// VLAN ID on which authentication failed.
	AuthFailVlanId pulumi.StringPtrInput
	// VLAN ID on which authentication failed.
	AuthFailVlanid pulumi.IntPtrInput
	// Authentication server timeout period (3 - 15 sec, default = 3).
	AuthserverTimeoutPeriod pulumi.IntPtrInput
	// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
	AuthserverTimeoutVlan pulumi.StringPtrInput
	// Authentication server timeout VLAN name.
	AuthserverTimeoutVlanid pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
	EapAutoUntaggedVlans pulumi.StringPtrInput
	// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
	EapPassthru pulumi.StringPtrInput
	// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
	FramevidApply pulumi.StringPtrInput
	// Guest authentication delay (1 - 900  sec, default = 30).
	GuestAuthDelay pulumi.IntPtrInput
	// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
	GuestVlan pulumi.StringPtrInput
	// Guest VLAN name.
	GuestVlanId pulumi.StringPtrInput
	// Guest VLAN ID.
	GuestVlanid pulumi.IntPtrInput
	// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
	MacAuthBypass pulumi.StringPtrInput
	// Group name.
	Name pulumi.StringPtrInput
	// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
	OpenAuth pulumi.StringPtrInput
	// Policy type. Valid values: `802.1X`.
	PolicyType pulumi.StringPtrInput
	// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
	RadiusTimeoutOverwrite pulumi.StringPtrInput
	// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
	SecurityMode pulumi.StringPtrInput
	// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
	UserGroups SwitchControllerSecurityPolicy8021XUserGroupArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSecurityPolicy8021XState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSecurityPolicy8021XState)(nil)).Elem()
}

type switchControllerSecurityPolicy8021XArgs struct {
	// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
	AuthFailVlan *string `pulumi:"authFailVlan"`
	// VLAN ID on which authentication failed.
	AuthFailVlanId *string `pulumi:"authFailVlanId"`
	// VLAN ID on which authentication failed.
	AuthFailVlanid *int `pulumi:"authFailVlanid"`
	// Authentication server timeout period (3 - 15 sec, default = 3).
	AuthserverTimeoutPeriod *int `pulumi:"authserverTimeoutPeriod"`
	// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
	AuthserverTimeoutVlan *string `pulumi:"authserverTimeoutVlan"`
	// Authentication server timeout VLAN name.
	AuthserverTimeoutVlanid *string `pulumi:"authserverTimeoutVlanid"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
	EapAutoUntaggedVlans *string `pulumi:"eapAutoUntaggedVlans"`
	// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
	EapPassthru *string `pulumi:"eapPassthru"`
	// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
	FramevidApply *string `pulumi:"framevidApply"`
	// Guest authentication delay (1 - 900  sec, default = 30).
	GuestAuthDelay *int `pulumi:"guestAuthDelay"`
	// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
	GuestVlan *string `pulumi:"guestVlan"`
	// Guest VLAN name.
	GuestVlanId *string `pulumi:"guestVlanId"`
	// Guest VLAN ID.
	GuestVlanid *int `pulumi:"guestVlanid"`
	// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
	MacAuthBypass *string `pulumi:"macAuthBypass"`
	// Group name.
	Name *string `pulumi:"name"`
	// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
	OpenAuth *string `pulumi:"openAuth"`
	// Policy type. Valid values: `802.1X`.
	PolicyType *string `pulumi:"policyType"`
	// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
	RadiusTimeoutOverwrite *string `pulumi:"radiusTimeoutOverwrite"`
	// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
	SecurityMode *string `pulumi:"securityMode"`
	// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
	UserGroups []SwitchControllerSecurityPolicy8021XUserGroup `pulumi:"userGroups"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerSecurityPolicy8021X resource.
type SwitchControllerSecurityPolicy8021XArgs struct {
	// Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
	AuthFailVlan pulumi.StringPtrInput
	// VLAN ID on which authentication failed.
	AuthFailVlanId pulumi.StringPtrInput
	// VLAN ID on which authentication failed.
	AuthFailVlanid pulumi.IntPtrInput
	// Authentication server timeout period (3 - 15 sec, default = 3).
	AuthserverTimeoutPeriod pulumi.IntPtrInput
	// Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
	AuthserverTimeoutVlan pulumi.StringPtrInput
	// Authentication server timeout VLAN name.
	AuthserverTimeoutVlanid pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
	EapAutoUntaggedVlans pulumi.StringPtrInput
	// Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
	EapPassthru pulumi.StringPtrInput
	// Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
	FramevidApply pulumi.StringPtrInput
	// Guest authentication delay (1 - 900  sec, default = 30).
	GuestAuthDelay pulumi.IntPtrInput
	// Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
	GuestVlan pulumi.StringPtrInput
	// Guest VLAN name.
	GuestVlanId pulumi.StringPtrInput
	// Guest VLAN ID.
	GuestVlanid pulumi.IntPtrInput
	// Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
	MacAuthBypass pulumi.StringPtrInput
	// Group name.
	Name pulumi.StringPtrInput
	// Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
	OpenAuth pulumi.StringPtrInput
	// Policy type. Valid values: `802.1X`.
	PolicyType pulumi.StringPtrInput
	// Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
	RadiusTimeoutOverwrite pulumi.StringPtrInput
	// Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
	SecurityMode pulumi.StringPtrInput
	// Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `userGroup` block is documented below.
	UserGroups SwitchControllerSecurityPolicy8021XUserGroupArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSecurityPolicy8021XArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSecurityPolicy8021XArgs)(nil)).Elem()
}

type SwitchControllerSecurityPolicy8021XInput interface {
	pulumi.Input

	ToSwitchControllerSecurityPolicy8021XOutput() SwitchControllerSecurityPolicy8021XOutput
	ToSwitchControllerSecurityPolicy8021XOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicy8021XOutput
}

func (*SwitchControllerSecurityPolicy8021X) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSecurityPolicy8021X)(nil)).Elem()
}

func (i *SwitchControllerSecurityPolicy8021X) ToSwitchControllerSecurityPolicy8021XOutput() SwitchControllerSecurityPolicy8021XOutput {
	return i.ToSwitchControllerSecurityPolicy8021XOutputWithContext(context.Background())
}

func (i *SwitchControllerSecurityPolicy8021X) ToSwitchControllerSecurityPolicy8021XOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicy8021XOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSecurityPolicy8021XOutput)
}

// SwitchControllerSecurityPolicy8021XArrayInput is an input type that accepts SwitchControllerSecurityPolicy8021XArray and SwitchControllerSecurityPolicy8021XArrayOutput values.
// You can construct a concrete instance of `SwitchControllerSecurityPolicy8021XArrayInput` via:
//
//          SwitchControllerSecurityPolicy8021XArray{ SwitchControllerSecurityPolicy8021XArgs{...} }
type SwitchControllerSecurityPolicy8021XArrayInput interface {
	pulumi.Input

	ToSwitchControllerSecurityPolicy8021XArrayOutput() SwitchControllerSecurityPolicy8021XArrayOutput
	ToSwitchControllerSecurityPolicy8021XArrayOutputWithContext(context.Context) SwitchControllerSecurityPolicy8021XArrayOutput
}

type SwitchControllerSecurityPolicy8021XArray []SwitchControllerSecurityPolicy8021XInput

func (SwitchControllerSecurityPolicy8021XArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSecurityPolicy8021X)(nil)).Elem()
}

func (i SwitchControllerSecurityPolicy8021XArray) ToSwitchControllerSecurityPolicy8021XArrayOutput() SwitchControllerSecurityPolicy8021XArrayOutput {
	return i.ToSwitchControllerSecurityPolicy8021XArrayOutputWithContext(context.Background())
}

func (i SwitchControllerSecurityPolicy8021XArray) ToSwitchControllerSecurityPolicy8021XArrayOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicy8021XArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSecurityPolicy8021XArrayOutput)
}

// SwitchControllerSecurityPolicy8021XMapInput is an input type that accepts SwitchControllerSecurityPolicy8021XMap and SwitchControllerSecurityPolicy8021XMapOutput values.
// You can construct a concrete instance of `SwitchControllerSecurityPolicy8021XMapInput` via:
//
//          SwitchControllerSecurityPolicy8021XMap{ "key": SwitchControllerSecurityPolicy8021XArgs{...} }
type SwitchControllerSecurityPolicy8021XMapInput interface {
	pulumi.Input

	ToSwitchControllerSecurityPolicy8021XMapOutput() SwitchControllerSecurityPolicy8021XMapOutput
	ToSwitchControllerSecurityPolicy8021XMapOutputWithContext(context.Context) SwitchControllerSecurityPolicy8021XMapOutput
}

type SwitchControllerSecurityPolicy8021XMap map[string]SwitchControllerSecurityPolicy8021XInput

func (SwitchControllerSecurityPolicy8021XMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSecurityPolicy8021X)(nil)).Elem()
}

func (i SwitchControllerSecurityPolicy8021XMap) ToSwitchControllerSecurityPolicy8021XMapOutput() SwitchControllerSecurityPolicy8021XMapOutput {
	return i.ToSwitchControllerSecurityPolicy8021XMapOutputWithContext(context.Background())
}

func (i SwitchControllerSecurityPolicy8021XMap) ToSwitchControllerSecurityPolicy8021XMapOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicy8021XMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSecurityPolicy8021XMapOutput)
}

type SwitchControllerSecurityPolicy8021XOutput struct{ *pulumi.OutputState }

func (SwitchControllerSecurityPolicy8021XOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSecurityPolicy8021X)(nil)).Elem()
}

func (o SwitchControllerSecurityPolicy8021XOutput) ToSwitchControllerSecurityPolicy8021XOutput() SwitchControllerSecurityPolicy8021XOutput {
	return o
}

func (o SwitchControllerSecurityPolicy8021XOutput) ToSwitchControllerSecurityPolicy8021XOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicy8021XOutput {
	return o
}

type SwitchControllerSecurityPolicy8021XArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerSecurityPolicy8021XArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSecurityPolicy8021X)(nil)).Elem()
}

func (o SwitchControllerSecurityPolicy8021XArrayOutput) ToSwitchControllerSecurityPolicy8021XArrayOutput() SwitchControllerSecurityPolicy8021XArrayOutput {
	return o
}

func (o SwitchControllerSecurityPolicy8021XArrayOutput) ToSwitchControllerSecurityPolicy8021XArrayOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicy8021XArrayOutput {
	return o
}

func (o SwitchControllerSecurityPolicy8021XArrayOutput) Index(i pulumi.IntInput) SwitchControllerSecurityPolicy8021XOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerSecurityPolicy8021X {
		return vs[0].([]*SwitchControllerSecurityPolicy8021X)[vs[1].(int)]
	}).(SwitchControllerSecurityPolicy8021XOutput)
}

type SwitchControllerSecurityPolicy8021XMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerSecurityPolicy8021XMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSecurityPolicy8021X)(nil)).Elem()
}

func (o SwitchControllerSecurityPolicy8021XMapOutput) ToSwitchControllerSecurityPolicy8021XMapOutput() SwitchControllerSecurityPolicy8021XMapOutput {
	return o
}

func (o SwitchControllerSecurityPolicy8021XMapOutput) ToSwitchControllerSecurityPolicy8021XMapOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicy8021XMapOutput {
	return o
}

func (o SwitchControllerSecurityPolicy8021XMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerSecurityPolicy8021XOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerSecurityPolicy8021X {
		return vs[0].(map[string]*SwitchControllerSecurityPolicy8021X)[vs[1].(string)]
	}).(SwitchControllerSecurityPolicy8021XOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSecurityPolicy8021XInput)(nil)).Elem(), &SwitchControllerSecurityPolicy8021X{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSecurityPolicy8021XArrayInput)(nil)).Elem(), SwitchControllerSecurityPolicy8021XArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSecurityPolicy8021XMapInput)(nil)).Elem(), SwitchControllerSecurityPolicy8021XMap{})
	pulumi.RegisterOutputType(SwitchControllerSecurityPolicy8021XOutput{})
	pulumi.RegisterOutputType(SwitchControllerSecurityPolicy8021XArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerSecurityPolicy8021XMapOutput{})
}

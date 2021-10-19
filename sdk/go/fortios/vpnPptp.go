// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure PPTP.
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
// 		_, err := fortios.NewVpnPptp(ctx, "trname", &fortios.VpnPptpArgs{
// 			Eip:     pulumi.String("1.1.1.22"),
// 			IpMode:  pulumi.String("range"),
// 			LocalIp: pulumi.String("0.0.0.0"),
// 			Sip:     pulumi.String("1.1.1.1"),
// 			Status:  pulumi.String("enable"),
// 			Usrgrp:  pulumi.String("Guest-group"),
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
// Vpn Pptp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/vpnPptp:VpnPptp labelname VpnPptp
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type VpnPptp struct {
	pulumi.CustomResourceState

	// End IP.
	Eip pulumi.StringOutput `pulumi:"eip"`
	// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
	IpMode pulumi.StringOutput `pulumi:"ipMode"`
	// Local IP to be used for peer's remote IP.
	LocalIp pulumi.StringOutput `pulumi:"localIp"`
	// Start IP.
	Sip pulumi.StringOutput `pulumi:"sip"`
	// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// User group.
	Usrgrp pulumi.StringOutput `pulumi:"usrgrp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpnPptp registers a new resource with the given unique name, arguments, and options.
func NewVpnPptp(ctx *pulumi.Context,
	name string, args *VpnPptpArgs, opts ...pulumi.ResourceOption) (*VpnPptp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	var resource VpnPptp
	err := ctx.RegisterResource("fortios:index/vpnPptp:VpnPptp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnPptp gets an existing VpnPptp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnPptp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnPptpState, opts ...pulumi.ResourceOption) (*VpnPptp, error) {
	var resource VpnPptp
	err := ctx.ReadResource("fortios:index/vpnPptp:VpnPptp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnPptp resources.
type vpnPptpState struct {
	// End IP.
	Eip *string `pulumi:"eip"`
	// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
	IpMode *string `pulumi:"ipMode"`
	// Local IP to be used for peer's remote IP.
	LocalIp *string `pulumi:"localIp"`
	// Start IP.
	Sip *string `pulumi:"sip"`
	// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// User group.
	Usrgrp *string `pulumi:"usrgrp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VpnPptpState struct {
	// End IP.
	Eip pulumi.StringPtrInput
	// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
	IpMode pulumi.StringPtrInput
	// Local IP to be used for peer's remote IP.
	LocalIp pulumi.StringPtrInput
	// Start IP.
	Sip pulumi.StringPtrInput
	// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// User group.
	Usrgrp pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpnPptpState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnPptpState)(nil)).Elem()
}

type vpnPptpArgs struct {
	// End IP.
	Eip *string `pulumi:"eip"`
	// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
	IpMode *string `pulumi:"ipMode"`
	// Local IP to be used for peer's remote IP.
	LocalIp *string `pulumi:"localIp"`
	// Start IP.
	Sip *string `pulumi:"sip"`
	// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// User group.
	Usrgrp *string `pulumi:"usrgrp"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnPptp resource.
type VpnPptpArgs struct {
	// End IP.
	Eip pulumi.StringPtrInput
	// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
	IpMode pulumi.StringPtrInput
	// Local IP to be used for peer's remote IP.
	LocalIp pulumi.StringPtrInput
	// Start IP.
	Sip pulumi.StringPtrInput
	// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// User group.
	Usrgrp pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpnPptpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnPptpArgs)(nil)).Elem()
}

type VpnPptpInput interface {
	pulumi.Input

	ToVpnPptpOutput() VpnPptpOutput
	ToVpnPptpOutputWithContext(ctx context.Context) VpnPptpOutput
}

func (*VpnPptp) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnPptp)(nil))
}

func (i *VpnPptp) ToVpnPptpOutput() VpnPptpOutput {
	return i.ToVpnPptpOutputWithContext(context.Background())
}

func (i *VpnPptp) ToVpnPptpOutputWithContext(ctx context.Context) VpnPptpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnPptpOutput)
}

func (i *VpnPptp) ToVpnPptpPtrOutput() VpnPptpPtrOutput {
	return i.ToVpnPptpPtrOutputWithContext(context.Background())
}

func (i *VpnPptp) ToVpnPptpPtrOutputWithContext(ctx context.Context) VpnPptpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnPptpPtrOutput)
}

type VpnPptpPtrInput interface {
	pulumi.Input

	ToVpnPptpPtrOutput() VpnPptpPtrOutput
	ToVpnPptpPtrOutputWithContext(ctx context.Context) VpnPptpPtrOutput
}

type vpnPptpPtrType VpnPptpArgs

func (*vpnPptpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnPptp)(nil))
}

func (i *vpnPptpPtrType) ToVpnPptpPtrOutput() VpnPptpPtrOutput {
	return i.ToVpnPptpPtrOutputWithContext(context.Background())
}

func (i *vpnPptpPtrType) ToVpnPptpPtrOutputWithContext(ctx context.Context) VpnPptpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnPptpPtrOutput)
}

// VpnPptpArrayInput is an input type that accepts VpnPptpArray and VpnPptpArrayOutput values.
// You can construct a concrete instance of `VpnPptpArrayInput` via:
//
//          VpnPptpArray{ VpnPptpArgs{...} }
type VpnPptpArrayInput interface {
	pulumi.Input

	ToVpnPptpArrayOutput() VpnPptpArrayOutput
	ToVpnPptpArrayOutputWithContext(context.Context) VpnPptpArrayOutput
}

type VpnPptpArray []VpnPptpInput

func (VpnPptpArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VpnPptp)(nil))
}

func (i VpnPptpArray) ToVpnPptpArrayOutput() VpnPptpArrayOutput {
	return i.ToVpnPptpArrayOutputWithContext(context.Background())
}

func (i VpnPptpArray) ToVpnPptpArrayOutputWithContext(ctx context.Context) VpnPptpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnPptpArrayOutput)
}

// VpnPptpMapInput is an input type that accepts VpnPptpMap and VpnPptpMapOutput values.
// You can construct a concrete instance of `VpnPptpMapInput` via:
//
//          VpnPptpMap{ "key": VpnPptpArgs{...} }
type VpnPptpMapInput interface {
	pulumi.Input

	ToVpnPptpMapOutput() VpnPptpMapOutput
	ToVpnPptpMapOutputWithContext(context.Context) VpnPptpMapOutput
}

type VpnPptpMap map[string]VpnPptpInput

func (VpnPptpMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VpnPptp)(nil))
}

func (i VpnPptpMap) ToVpnPptpMapOutput() VpnPptpMapOutput {
	return i.ToVpnPptpMapOutputWithContext(context.Background())
}

func (i VpnPptpMap) ToVpnPptpMapOutputWithContext(ctx context.Context) VpnPptpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnPptpMapOutput)
}

type VpnPptpOutput struct {
	*pulumi.OutputState
}

func (VpnPptpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnPptp)(nil))
}

func (o VpnPptpOutput) ToVpnPptpOutput() VpnPptpOutput {
	return o
}

func (o VpnPptpOutput) ToVpnPptpOutputWithContext(ctx context.Context) VpnPptpOutput {
	return o
}

func (o VpnPptpOutput) ToVpnPptpPtrOutput() VpnPptpPtrOutput {
	return o.ToVpnPptpPtrOutputWithContext(context.Background())
}

func (o VpnPptpOutput) ToVpnPptpPtrOutputWithContext(ctx context.Context) VpnPptpPtrOutput {
	return o.ApplyT(func(v VpnPptp) *VpnPptp {
		return &v
	}).(VpnPptpPtrOutput)
}

type VpnPptpPtrOutput struct {
	*pulumi.OutputState
}

func (VpnPptpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnPptp)(nil))
}

func (o VpnPptpPtrOutput) ToVpnPptpPtrOutput() VpnPptpPtrOutput {
	return o
}

func (o VpnPptpPtrOutput) ToVpnPptpPtrOutputWithContext(ctx context.Context) VpnPptpPtrOutput {
	return o
}

type VpnPptpArrayOutput struct{ *pulumi.OutputState }

func (VpnPptpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnPptp)(nil))
}

func (o VpnPptpArrayOutput) ToVpnPptpArrayOutput() VpnPptpArrayOutput {
	return o
}

func (o VpnPptpArrayOutput) ToVpnPptpArrayOutputWithContext(ctx context.Context) VpnPptpArrayOutput {
	return o
}

func (o VpnPptpArrayOutput) Index(i pulumi.IntInput) VpnPptpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnPptp {
		return vs[0].([]VpnPptp)[vs[1].(int)]
	}).(VpnPptpOutput)
}

type VpnPptpMapOutput struct{ *pulumi.OutputState }

func (VpnPptpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpnPptp)(nil))
}

func (o VpnPptpMapOutput) ToVpnPptpMapOutput() VpnPptpMapOutput {
	return o
}

func (o VpnPptpMapOutput) ToVpnPptpMapOutputWithContext(ctx context.Context) VpnPptpMapOutput {
	return o
}

func (o VpnPptpMapOutput) MapIndex(k pulumi.StringInput) VpnPptpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpnPptp {
		return vs[0].(map[string]VpnPptp)[vs[1].(string)]
	}).(VpnPptpOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnPptpOutput{})
	pulumi.RegisterOutputType(VpnPptpPtrOutput{})
	pulumi.RegisterOutputType(VpnPptpArrayOutput{})
	pulumi.RegisterOutputType(VpnPptpMapOutput{})
}

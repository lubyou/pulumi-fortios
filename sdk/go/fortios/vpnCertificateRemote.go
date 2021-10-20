// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Remote certificate as a PEM file.
//
// ## Import
//
// VpnCertificate Remote can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/vpnCertificateRemote:VpnCertificateRemote labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type VpnCertificateRemote struct {
	pulumi.CustomResourceState

	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringOutput `pulumi:"range"`
	// Remote certificate.
	Remote pulumi.StringOutput `pulumi:"remote"`
	// Remote certificate source type.
	Source pulumi.StringOutput `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpnCertificateRemote registers a new resource with the given unique name, arguments, and options.
func NewVpnCertificateRemote(ctx *pulumi.Context,
	name string, args *VpnCertificateRemoteArgs, opts ...pulumi.ResourceOption) (*VpnCertificateRemote, error) {
	if args == nil {
		args = &VpnCertificateRemoteArgs{}
	}

	var resource VpnCertificateRemote
	err := ctx.RegisterResource("fortios:index/vpnCertificateRemote:VpnCertificateRemote", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnCertificateRemote gets an existing VpnCertificateRemote resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnCertificateRemote(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnCertificateRemoteState, opts ...pulumi.ResourceOption) (*VpnCertificateRemote, error) {
	var resource VpnCertificateRemote
	err := ctx.ReadResource("fortios:index/vpnCertificateRemote:VpnCertificateRemote", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnCertificateRemote resources.
type vpnCertificateRemoteState struct {
	// Name.
	Name *string `pulumi:"name"`
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// Remote certificate.
	Remote *string `pulumi:"remote"`
	// Remote certificate source type.
	Source *string `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VpnCertificateRemoteState struct {
	// Name.
	Name pulumi.StringPtrInput
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// Remote certificate.
	Remote pulumi.StringPtrInput
	// Remote certificate source type.
	Source pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpnCertificateRemoteState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCertificateRemoteState)(nil)).Elem()
}

type vpnCertificateRemoteArgs struct {
	// Name.
	Name *string `pulumi:"name"`
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// Remote certificate.
	Remote *string `pulumi:"remote"`
	// Remote certificate source type.
	Source *string `pulumi:"source"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnCertificateRemote resource.
type VpnCertificateRemoteArgs struct {
	// Name.
	Name pulumi.StringPtrInput
	// Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// Remote certificate.
	Remote pulumi.StringPtrInput
	// Remote certificate source type.
	Source pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpnCertificateRemoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCertificateRemoteArgs)(nil)).Elem()
}

type VpnCertificateRemoteInput interface {
	pulumi.Input

	ToVpnCertificateRemoteOutput() VpnCertificateRemoteOutput
	ToVpnCertificateRemoteOutputWithContext(ctx context.Context) VpnCertificateRemoteOutput
}

func (*VpnCertificateRemote) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnCertificateRemote)(nil))
}

func (i *VpnCertificateRemote) ToVpnCertificateRemoteOutput() VpnCertificateRemoteOutput {
	return i.ToVpnCertificateRemoteOutputWithContext(context.Background())
}

func (i *VpnCertificateRemote) ToVpnCertificateRemoteOutputWithContext(ctx context.Context) VpnCertificateRemoteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateRemoteOutput)
}

func (i *VpnCertificateRemote) ToVpnCertificateRemotePtrOutput() VpnCertificateRemotePtrOutput {
	return i.ToVpnCertificateRemotePtrOutputWithContext(context.Background())
}

func (i *VpnCertificateRemote) ToVpnCertificateRemotePtrOutputWithContext(ctx context.Context) VpnCertificateRemotePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateRemotePtrOutput)
}

type VpnCertificateRemotePtrInput interface {
	pulumi.Input

	ToVpnCertificateRemotePtrOutput() VpnCertificateRemotePtrOutput
	ToVpnCertificateRemotePtrOutputWithContext(ctx context.Context) VpnCertificateRemotePtrOutput
}

type vpnCertificateRemotePtrType VpnCertificateRemoteArgs

func (*vpnCertificateRemotePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnCertificateRemote)(nil))
}

func (i *vpnCertificateRemotePtrType) ToVpnCertificateRemotePtrOutput() VpnCertificateRemotePtrOutput {
	return i.ToVpnCertificateRemotePtrOutputWithContext(context.Background())
}

func (i *vpnCertificateRemotePtrType) ToVpnCertificateRemotePtrOutputWithContext(ctx context.Context) VpnCertificateRemotePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateRemotePtrOutput)
}

// VpnCertificateRemoteArrayInput is an input type that accepts VpnCertificateRemoteArray and VpnCertificateRemoteArrayOutput values.
// You can construct a concrete instance of `VpnCertificateRemoteArrayInput` via:
//
//          VpnCertificateRemoteArray{ VpnCertificateRemoteArgs{...} }
type VpnCertificateRemoteArrayInput interface {
	pulumi.Input

	ToVpnCertificateRemoteArrayOutput() VpnCertificateRemoteArrayOutput
	ToVpnCertificateRemoteArrayOutputWithContext(context.Context) VpnCertificateRemoteArrayOutput
}

type VpnCertificateRemoteArray []VpnCertificateRemoteInput

func (VpnCertificateRemoteArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VpnCertificateRemote)(nil))
}

func (i VpnCertificateRemoteArray) ToVpnCertificateRemoteArrayOutput() VpnCertificateRemoteArrayOutput {
	return i.ToVpnCertificateRemoteArrayOutputWithContext(context.Background())
}

func (i VpnCertificateRemoteArray) ToVpnCertificateRemoteArrayOutputWithContext(ctx context.Context) VpnCertificateRemoteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateRemoteArrayOutput)
}

// VpnCertificateRemoteMapInput is an input type that accepts VpnCertificateRemoteMap and VpnCertificateRemoteMapOutput values.
// You can construct a concrete instance of `VpnCertificateRemoteMapInput` via:
//
//          VpnCertificateRemoteMap{ "key": VpnCertificateRemoteArgs{...} }
type VpnCertificateRemoteMapInput interface {
	pulumi.Input

	ToVpnCertificateRemoteMapOutput() VpnCertificateRemoteMapOutput
	ToVpnCertificateRemoteMapOutputWithContext(context.Context) VpnCertificateRemoteMapOutput
}

type VpnCertificateRemoteMap map[string]VpnCertificateRemoteInput

func (VpnCertificateRemoteMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VpnCertificateRemote)(nil))
}

func (i VpnCertificateRemoteMap) ToVpnCertificateRemoteMapOutput() VpnCertificateRemoteMapOutput {
	return i.ToVpnCertificateRemoteMapOutputWithContext(context.Background())
}

func (i VpnCertificateRemoteMap) ToVpnCertificateRemoteMapOutputWithContext(ctx context.Context) VpnCertificateRemoteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateRemoteMapOutput)
}

type VpnCertificateRemoteOutput struct {
	*pulumi.OutputState
}

func (VpnCertificateRemoteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnCertificateRemote)(nil))
}

func (o VpnCertificateRemoteOutput) ToVpnCertificateRemoteOutput() VpnCertificateRemoteOutput {
	return o
}

func (o VpnCertificateRemoteOutput) ToVpnCertificateRemoteOutputWithContext(ctx context.Context) VpnCertificateRemoteOutput {
	return o
}

func (o VpnCertificateRemoteOutput) ToVpnCertificateRemotePtrOutput() VpnCertificateRemotePtrOutput {
	return o.ToVpnCertificateRemotePtrOutputWithContext(context.Background())
}

func (o VpnCertificateRemoteOutput) ToVpnCertificateRemotePtrOutputWithContext(ctx context.Context) VpnCertificateRemotePtrOutput {
	return o.ApplyT(func(v VpnCertificateRemote) *VpnCertificateRemote {
		return &v
	}).(VpnCertificateRemotePtrOutput)
}

type VpnCertificateRemotePtrOutput struct {
	*pulumi.OutputState
}

func (VpnCertificateRemotePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnCertificateRemote)(nil))
}

func (o VpnCertificateRemotePtrOutput) ToVpnCertificateRemotePtrOutput() VpnCertificateRemotePtrOutput {
	return o
}

func (o VpnCertificateRemotePtrOutput) ToVpnCertificateRemotePtrOutputWithContext(ctx context.Context) VpnCertificateRemotePtrOutput {
	return o
}

type VpnCertificateRemoteArrayOutput struct{ *pulumi.OutputState }

func (VpnCertificateRemoteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnCertificateRemote)(nil))
}

func (o VpnCertificateRemoteArrayOutput) ToVpnCertificateRemoteArrayOutput() VpnCertificateRemoteArrayOutput {
	return o
}

func (o VpnCertificateRemoteArrayOutput) ToVpnCertificateRemoteArrayOutputWithContext(ctx context.Context) VpnCertificateRemoteArrayOutput {
	return o
}

func (o VpnCertificateRemoteArrayOutput) Index(i pulumi.IntInput) VpnCertificateRemoteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnCertificateRemote {
		return vs[0].([]VpnCertificateRemote)[vs[1].(int)]
	}).(VpnCertificateRemoteOutput)
}

type VpnCertificateRemoteMapOutput struct{ *pulumi.OutputState }

func (VpnCertificateRemoteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpnCertificateRemote)(nil))
}

func (o VpnCertificateRemoteMapOutput) ToVpnCertificateRemoteMapOutput() VpnCertificateRemoteMapOutput {
	return o
}

func (o VpnCertificateRemoteMapOutput) ToVpnCertificateRemoteMapOutputWithContext(ctx context.Context) VpnCertificateRemoteMapOutput {
	return o
}

func (o VpnCertificateRemoteMapOutput) MapIndex(k pulumi.StringInput) VpnCertificateRemoteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpnCertificateRemote {
		return vs[0].(map[string]VpnCertificateRemote)[vs[1].(string)]
	}).(VpnCertificateRemoteOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnCertificateRemoteOutput{})
	pulumi.RegisterOutputType(VpnCertificateRemotePtrOutput{})
	pulumi.RegisterOutputType(VpnCertificateRemoteArrayOutput{})
	pulumi.RegisterOutputType(VpnCertificateRemoteMapOutput{})
}
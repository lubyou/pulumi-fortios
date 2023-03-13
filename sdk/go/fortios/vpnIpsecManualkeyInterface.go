// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnIpsecManualkeyInterface struct {
	pulumi.CustomResourceState

	AddrType   pulumi.StringOutput    `pulumi:"addrType"`
	AuthAlg    pulumi.StringOutput    `pulumi:"authAlg"`
	AuthKey    pulumi.StringOutput    `pulumi:"authKey"`
	EncAlg     pulumi.StringOutput    `pulumi:"encAlg"`
	EncKey     pulumi.StringOutput    `pulumi:"encKey"`
	Interface  pulumi.StringOutput    `pulumi:"interface"`
	IpVersion  pulumi.StringOutput    `pulumi:"ipVersion"`
	LocalGw    pulumi.StringOutput    `pulumi:"localGw"`
	LocalGw6   pulumi.StringOutput    `pulumi:"localGw6"`
	LocalSpi   pulumi.StringOutput    `pulumi:"localSpi"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	NpuOffload pulumi.StringOutput    `pulumi:"npuOffload"`
	RemoteGw   pulumi.StringOutput    `pulumi:"remoteGw"`
	RemoteGw6  pulumi.StringOutput    `pulumi:"remoteGw6"`
	RemoteSpi  pulumi.StringOutput    `pulumi:"remoteSpi"`
	Vdomparam  pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpnIpsecManualkeyInterface registers a new resource with the given unique name, arguments, and options.
func NewVpnIpsecManualkeyInterface(ctx *pulumi.Context,
	name string, args *VpnIpsecManualkeyInterfaceArgs, opts ...pulumi.ResourceOption) (*VpnIpsecManualkeyInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthAlg == nil {
		return nil, errors.New("invalid value for required argument 'AuthAlg'")
	}
	if args.EncAlg == nil {
		return nil, errors.New("invalid value for required argument 'EncAlg'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.RemoteGw == nil {
		return nil, errors.New("invalid value for required argument 'RemoteGw'")
	}
	if args.RemoteGw6 == nil {
		return nil, errors.New("invalid value for required argument 'RemoteGw6'")
	}
	if args.AuthKey != nil {
		args.AuthKey = pulumi.ToSecret(args.AuthKey).(pulumi.StringPtrInput)
	}
	if args.EncKey != nil {
		args.EncKey = pulumi.ToSecret(args.EncKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"authKey",
		"encKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource VpnIpsecManualkeyInterface
	err := ctx.RegisterResource("fortios:index/vpnIpsecManualkeyInterface:VpnIpsecManualkeyInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnIpsecManualkeyInterface gets an existing VpnIpsecManualkeyInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnIpsecManualkeyInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnIpsecManualkeyInterfaceState, opts ...pulumi.ResourceOption) (*VpnIpsecManualkeyInterface, error) {
	var resource VpnIpsecManualkeyInterface
	err := ctx.ReadResource("fortios:index/vpnIpsecManualkeyInterface:VpnIpsecManualkeyInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnIpsecManualkeyInterface resources.
type vpnIpsecManualkeyInterfaceState struct {
	AddrType   *string `pulumi:"addrType"`
	AuthAlg    *string `pulumi:"authAlg"`
	AuthKey    *string `pulumi:"authKey"`
	EncAlg     *string `pulumi:"encAlg"`
	EncKey     *string `pulumi:"encKey"`
	Interface  *string `pulumi:"interface"`
	IpVersion  *string `pulumi:"ipVersion"`
	LocalGw    *string `pulumi:"localGw"`
	LocalGw6   *string `pulumi:"localGw6"`
	LocalSpi   *string `pulumi:"localSpi"`
	Name       *string `pulumi:"name"`
	NpuOffload *string `pulumi:"npuOffload"`
	RemoteGw   *string `pulumi:"remoteGw"`
	RemoteGw6  *string `pulumi:"remoteGw6"`
	RemoteSpi  *string `pulumi:"remoteSpi"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

type VpnIpsecManualkeyInterfaceState struct {
	AddrType   pulumi.StringPtrInput
	AuthAlg    pulumi.StringPtrInput
	AuthKey    pulumi.StringPtrInput
	EncAlg     pulumi.StringPtrInput
	EncKey     pulumi.StringPtrInput
	Interface  pulumi.StringPtrInput
	IpVersion  pulumi.StringPtrInput
	LocalGw    pulumi.StringPtrInput
	LocalGw6   pulumi.StringPtrInput
	LocalSpi   pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
	NpuOffload pulumi.StringPtrInput
	RemoteGw   pulumi.StringPtrInput
	RemoteGw6  pulumi.StringPtrInput
	RemoteSpi  pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (VpnIpsecManualkeyInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnIpsecManualkeyInterfaceState)(nil)).Elem()
}

type vpnIpsecManualkeyInterfaceArgs struct {
	AddrType   *string `pulumi:"addrType"`
	AuthAlg    string  `pulumi:"authAlg"`
	AuthKey    *string `pulumi:"authKey"`
	EncAlg     string  `pulumi:"encAlg"`
	EncKey     *string `pulumi:"encKey"`
	Interface  string  `pulumi:"interface"`
	IpVersion  *string `pulumi:"ipVersion"`
	LocalGw    *string `pulumi:"localGw"`
	LocalGw6   *string `pulumi:"localGw6"`
	LocalSpi   *string `pulumi:"localSpi"`
	Name       *string `pulumi:"name"`
	NpuOffload *string `pulumi:"npuOffload"`
	RemoteGw   string  `pulumi:"remoteGw"`
	RemoteGw6  string  `pulumi:"remoteGw6"`
	RemoteSpi  *string `pulumi:"remoteSpi"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnIpsecManualkeyInterface resource.
type VpnIpsecManualkeyInterfaceArgs struct {
	AddrType   pulumi.StringPtrInput
	AuthAlg    pulumi.StringInput
	AuthKey    pulumi.StringPtrInput
	EncAlg     pulumi.StringInput
	EncKey     pulumi.StringPtrInput
	Interface  pulumi.StringInput
	IpVersion  pulumi.StringPtrInput
	LocalGw    pulumi.StringPtrInput
	LocalGw6   pulumi.StringPtrInput
	LocalSpi   pulumi.StringPtrInput
	Name       pulumi.StringPtrInput
	NpuOffload pulumi.StringPtrInput
	RemoteGw   pulumi.StringInput
	RemoteGw6  pulumi.StringInput
	RemoteSpi  pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (VpnIpsecManualkeyInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnIpsecManualkeyInterfaceArgs)(nil)).Elem()
}

type VpnIpsecManualkeyInterfaceInput interface {
	pulumi.Input

	ToVpnIpsecManualkeyInterfaceOutput() VpnIpsecManualkeyInterfaceOutput
	ToVpnIpsecManualkeyInterfaceOutputWithContext(ctx context.Context) VpnIpsecManualkeyInterfaceOutput
}

func (*VpnIpsecManualkeyInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnIpsecManualkeyInterface)(nil)).Elem()
}

func (i *VpnIpsecManualkeyInterface) ToVpnIpsecManualkeyInterfaceOutput() VpnIpsecManualkeyInterfaceOutput {
	return i.ToVpnIpsecManualkeyInterfaceOutputWithContext(context.Background())
}

func (i *VpnIpsecManualkeyInterface) ToVpnIpsecManualkeyInterfaceOutputWithContext(ctx context.Context) VpnIpsecManualkeyInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecManualkeyInterfaceOutput)
}

// VpnIpsecManualkeyInterfaceArrayInput is an input type that accepts VpnIpsecManualkeyInterfaceArray and VpnIpsecManualkeyInterfaceArrayOutput values.
// You can construct a concrete instance of `VpnIpsecManualkeyInterfaceArrayInput` via:
//
//	VpnIpsecManualkeyInterfaceArray{ VpnIpsecManualkeyInterfaceArgs{...} }
type VpnIpsecManualkeyInterfaceArrayInput interface {
	pulumi.Input

	ToVpnIpsecManualkeyInterfaceArrayOutput() VpnIpsecManualkeyInterfaceArrayOutput
	ToVpnIpsecManualkeyInterfaceArrayOutputWithContext(context.Context) VpnIpsecManualkeyInterfaceArrayOutput
}

type VpnIpsecManualkeyInterfaceArray []VpnIpsecManualkeyInterfaceInput

func (VpnIpsecManualkeyInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnIpsecManualkeyInterface)(nil)).Elem()
}

func (i VpnIpsecManualkeyInterfaceArray) ToVpnIpsecManualkeyInterfaceArrayOutput() VpnIpsecManualkeyInterfaceArrayOutput {
	return i.ToVpnIpsecManualkeyInterfaceArrayOutputWithContext(context.Background())
}

func (i VpnIpsecManualkeyInterfaceArray) ToVpnIpsecManualkeyInterfaceArrayOutputWithContext(ctx context.Context) VpnIpsecManualkeyInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecManualkeyInterfaceArrayOutput)
}

// VpnIpsecManualkeyInterfaceMapInput is an input type that accepts VpnIpsecManualkeyInterfaceMap and VpnIpsecManualkeyInterfaceMapOutput values.
// You can construct a concrete instance of `VpnIpsecManualkeyInterfaceMapInput` via:
//
//	VpnIpsecManualkeyInterfaceMap{ "key": VpnIpsecManualkeyInterfaceArgs{...} }
type VpnIpsecManualkeyInterfaceMapInput interface {
	pulumi.Input

	ToVpnIpsecManualkeyInterfaceMapOutput() VpnIpsecManualkeyInterfaceMapOutput
	ToVpnIpsecManualkeyInterfaceMapOutputWithContext(context.Context) VpnIpsecManualkeyInterfaceMapOutput
}

type VpnIpsecManualkeyInterfaceMap map[string]VpnIpsecManualkeyInterfaceInput

func (VpnIpsecManualkeyInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnIpsecManualkeyInterface)(nil)).Elem()
}

func (i VpnIpsecManualkeyInterfaceMap) ToVpnIpsecManualkeyInterfaceMapOutput() VpnIpsecManualkeyInterfaceMapOutput {
	return i.ToVpnIpsecManualkeyInterfaceMapOutputWithContext(context.Background())
}

func (i VpnIpsecManualkeyInterfaceMap) ToVpnIpsecManualkeyInterfaceMapOutputWithContext(ctx context.Context) VpnIpsecManualkeyInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecManualkeyInterfaceMapOutput)
}

type VpnIpsecManualkeyInterfaceOutput struct{ *pulumi.OutputState }

func (VpnIpsecManualkeyInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnIpsecManualkeyInterface)(nil)).Elem()
}

func (o VpnIpsecManualkeyInterfaceOutput) ToVpnIpsecManualkeyInterfaceOutput() VpnIpsecManualkeyInterfaceOutput {
	return o
}

func (o VpnIpsecManualkeyInterfaceOutput) ToVpnIpsecManualkeyInterfaceOutputWithContext(ctx context.Context) VpnIpsecManualkeyInterfaceOutput {
	return o
}

func (o VpnIpsecManualkeyInterfaceOutput) AddrType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.AddrType }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) AuthAlg() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.AuthAlg }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) AuthKey() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.AuthKey }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) EncAlg() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.EncAlg }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) EncKey() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.EncKey }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.IpVersion }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) LocalGw() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.LocalGw }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) LocalGw6() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.LocalGw6 }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) LocalSpi() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.LocalSpi }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) NpuOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.NpuOffload }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) RemoteGw() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.RemoteGw }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) RemoteGw6() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.RemoteGw6 }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) RemoteSpi() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringOutput { return v.RemoteSpi }).(pulumi.StringOutput)
}

func (o VpnIpsecManualkeyInterfaceOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnIpsecManualkeyInterface) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpnIpsecManualkeyInterfaceArrayOutput struct{ *pulumi.OutputState }

func (VpnIpsecManualkeyInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnIpsecManualkeyInterface)(nil)).Elem()
}

func (o VpnIpsecManualkeyInterfaceArrayOutput) ToVpnIpsecManualkeyInterfaceArrayOutput() VpnIpsecManualkeyInterfaceArrayOutput {
	return o
}

func (o VpnIpsecManualkeyInterfaceArrayOutput) ToVpnIpsecManualkeyInterfaceArrayOutputWithContext(ctx context.Context) VpnIpsecManualkeyInterfaceArrayOutput {
	return o
}

func (o VpnIpsecManualkeyInterfaceArrayOutput) Index(i pulumi.IntInput) VpnIpsecManualkeyInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnIpsecManualkeyInterface {
		return vs[0].([]*VpnIpsecManualkeyInterface)[vs[1].(int)]
	}).(VpnIpsecManualkeyInterfaceOutput)
}

type VpnIpsecManualkeyInterfaceMapOutput struct{ *pulumi.OutputState }

func (VpnIpsecManualkeyInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnIpsecManualkeyInterface)(nil)).Elem()
}

func (o VpnIpsecManualkeyInterfaceMapOutput) ToVpnIpsecManualkeyInterfaceMapOutput() VpnIpsecManualkeyInterfaceMapOutput {
	return o
}

func (o VpnIpsecManualkeyInterfaceMapOutput) ToVpnIpsecManualkeyInterfaceMapOutputWithContext(ctx context.Context) VpnIpsecManualkeyInterfaceMapOutput {
	return o
}

func (o VpnIpsecManualkeyInterfaceMapOutput) MapIndex(k pulumi.StringInput) VpnIpsecManualkeyInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnIpsecManualkeyInterface {
		return vs[0].(map[string]*VpnIpsecManualkeyInterface)[vs[1].(string)]
	}).(VpnIpsecManualkeyInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecManualkeyInterfaceInput)(nil)).Elem(), &VpnIpsecManualkeyInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecManualkeyInterfaceArrayInput)(nil)).Elem(), VpnIpsecManualkeyInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecManualkeyInterfaceMapInput)(nil)).Elem(), VpnIpsecManualkeyInterfaceMap{})
	pulumi.RegisterOutputType(VpnIpsecManualkeyInterfaceOutput{})
	pulumi.RegisterOutputType(VpnIpsecManualkeyInterfaceArrayOutput{})
	pulumi.RegisterOutputType(VpnIpsecManualkeyInterfaceMapOutput{})
}

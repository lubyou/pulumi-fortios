// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnSslWebRealm struct {
	pulumi.CustomResourceState

	LoginPage             pulumi.StringPtrOutput `pulumi:"loginPage"`
	MaxConcurrentUser     pulumi.IntOutput       `pulumi:"maxConcurrentUser"`
	NasIp                 pulumi.StringOutput    `pulumi:"nasIp"`
	RadiusPort            pulumi.IntOutput       `pulumi:"radiusPort"`
	RadiusServer          pulumi.StringOutput    `pulumi:"radiusServer"`
	UrlPath               pulumi.StringOutput    `pulumi:"urlPath"`
	Vdomparam             pulumi.StringPtrOutput `pulumi:"vdomparam"`
	VirtualHost           pulumi.StringPtrOutput `pulumi:"virtualHost"`
	VirtualHostOnly       pulumi.StringOutput    `pulumi:"virtualHostOnly"`
	VirtualHostServerCert pulumi.StringOutput    `pulumi:"virtualHostServerCert"`
}

// NewVpnSslWebRealm registers a new resource with the given unique name, arguments, and options.
func NewVpnSslWebRealm(ctx *pulumi.Context,
	name string, args *VpnSslWebRealmArgs, opts ...pulumi.ResourceOption) (*VpnSslWebRealm, error) {
	if args == nil {
		args = &VpnSslWebRealmArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource VpnSslWebRealm
	err := ctx.RegisterResource("fortios:index/vpnSslWebRealm:VpnSslWebRealm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnSslWebRealm gets an existing VpnSslWebRealm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnSslWebRealm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnSslWebRealmState, opts ...pulumi.ResourceOption) (*VpnSslWebRealm, error) {
	var resource VpnSslWebRealm
	err := ctx.ReadResource("fortios:index/vpnSslWebRealm:VpnSslWebRealm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnSslWebRealm resources.
type vpnSslWebRealmState struct {
	LoginPage             *string `pulumi:"loginPage"`
	MaxConcurrentUser     *int    `pulumi:"maxConcurrentUser"`
	NasIp                 *string `pulumi:"nasIp"`
	RadiusPort            *int    `pulumi:"radiusPort"`
	RadiusServer          *string `pulumi:"radiusServer"`
	UrlPath               *string `pulumi:"urlPath"`
	Vdomparam             *string `pulumi:"vdomparam"`
	VirtualHost           *string `pulumi:"virtualHost"`
	VirtualHostOnly       *string `pulumi:"virtualHostOnly"`
	VirtualHostServerCert *string `pulumi:"virtualHostServerCert"`
}

type VpnSslWebRealmState struct {
	LoginPage             pulumi.StringPtrInput
	MaxConcurrentUser     pulumi.IntPtrInput
	NasIp                 pulumi.StringPtrInput
	RadiusPort            pulumi.IntPtrInput
	RadiusServer          pulumi.StringPtrInput
	UrlPath               pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
	VirtualHost           pulumi.StringPtrInput
	VirtualHostOnly       pulumi.StringPtrInput
	VirtualHostServerCert pulumi.StringPtrInput
}

func (VpnSslWebRealmState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSslWebRealmState)(nil)).Elem()
}

type vpnSslWebRealmArgs struct {
	LoginPage             *string `pulumi:"loginPage"`
	MaxConcurrentUser     *int    `pulumi:"maxConcurrentUser"`
	NasIp                 *string `pulumi:"nasIp"`
	RadiusPort            *int    `pulumi:"radiusPort"`
	RadiusServer          *string `pulumi:"radiusServer"`
	UrlPath               *string `pulumi:"urlPath"`
	Vdomparam             *string `pulumi:"vdomparam"`
	VirtualHost           *string `pulumi:"virtualHost"`
	VirtualHostOnly       *string `pulumi:"virtualHostOnly"`
	VirtualHostServerCert *string `pulumi:"virtualHostServerCert"`
}

// The set of arguments for constructing a VpnSslWebRealm resource.
type VpnSslWebRealmArgs struct {
	LoginPage             pulumi.StringPtrInput
	MaxConcurrentUser     pulumi.IntPtrInput
	NasIp                 pulumi.StringPtrInput
	RadiusPort            pulumi.IntPtrInput
	RadiusServer          pulumi.StringPtrInput
	UrlPath               pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
	VirtualHost           pulumi.StringPtrInput
	VirtualHostOnly       pulumi.StringPtrInput
	VirtualHostServerCert pulumi.StringPtrInput
}

func (VpnSslWebRealmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSslWebRealmArgs)(nil)).Elem()
}

type VpnSslWebRealmInput interface {
	pulumi.Input

	ToVpnSslWebRealmOutput() VpnSslWebRealmOutput
	ToVpnSslWebRealmOutputWithContext(ctx context.Context) VpnSslWebRealmOutput
}

func (*VpnSslWebRealm) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnSslWebRealm)(nil)).Elem()
}

func (i *VpnSslWebRealm) ToVpnSslWebRealmOutput() VpnSslWebRealmOutput {
	return i.ToVpnSslWebRealmOutputWithContext(context.Background())
}

func (i *VpnSslWebRealm) ToVpnSslWebRealmOutputWithContext(ctx context.Context) VpnSslWebRealmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSslWebRealmOutput)
}

// VpnSslWebRealmArrayInput is an input type that accepts VpnSslWebRealmArray and VpnSslWebRealmArrayOutput values.
// You can construct a concrete instance of `VpnSslWebRealmArrayInput` via:
//
//	VpnSslWebRealmArray{ VpnSslWebRealmArgs{...} }
type VpnSslWebRealmArrayInput interface {
	pulumi.Input

	ToVpnSslWebRealmArrayOutput() VpnSslWebRealmArrayOutput
	ToVpnSslWebRealmArrayOutputWithContext(context.Context) VpnSslWebRealmArrayOutput
}

type VpnSslWebRealmArray []VpnSslWebRealmInput

func (VpnSslWebRealmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnSslWebRealm)(nil)).Elem()
}

func (i VpnSslWebRealmArray) ToVpnSslWebRealmArrayOutput() VpnSslWebRealmArrayOutput {
	return i.ToVpnSslWebRealmArrayOutputWithContext(context.Background())
}

func (i VpnSslWebRealmArray) ToVpnSslWebRealmArrayOutputWithContext(ctx context.Context) VpnSslWebRealmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSslWebRealmArrayOutput)
}

// VpnSslWebRealmMapInput is an input type that accepts VpnSslWebRealmMap and VpnSslWebRealmMapOutput values.
// You can construct a concrete instance of `VpnSslWebRealmMapInput` via:
//
//	VpnSslWebRealmMap{ "key": VpnSslWebRealmArgs{...} }
type VpnSslWebRealmMapInput interface {
	pulumi.Input

	ToVpnSslWebRealmMapOutput() VpnSslWebRealmMapOutput
	ToVpnSslWebRealmMapOutputWithContext(context.Context) VpnSslWebRealmMapOutput
}

type VpnSslWebRealmMap map[string]VpnSslWebRealmInput

func (VpnSslWebRealmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnSslWebRealm)(nil)).Elem()
}

func (i VpnSslWebRealmMap) ToVpnSslWebRealmMapOutput() VpnSslWebRealmMapOutput {
	return i.ToVpnSslWebRealmMapOutputWithContext(context.Background())
}

func (i VpnSslWebRealmMap) ToVpnSslWebRealmMapOutputWithContext(ctx context.Context) VpnSslWebRealmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSslWebRealmMapOutput)
}

type VpnSslWebRealmOutput struct{ *pulumi.OutputState }

func (VpnSslWebRealmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnSslWebRealm)(nil)).Elem()
}

func (o VpnSslWebRealmOutput) ToVpnSslWebRealmOutput() VpnSslWebRealmOutput {
	return o
}

func (o VpnSslWebRealmOutput) ToVpnSslWebRealmOutputWithContext(ctx context.Context) VpnSslWebRealmOutput {
	return o
}

func (o VpnSslWebRealmOutput) LoginPage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnSslWebRealm) pulumi.StringPtrOutput { return v.LoginPage }).(pulumi.StringPtrOutput)
}

func (o VpnSslWebRealmOutput) MaxConcurrentUser() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnSslWebRealm) pulumi.IntOutput { return v.MaxConcurrentUser }).(pulumi.IntOutput)
}

func (o VpnSslWebRealmOutput) NasIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSslWebRealm) pulumi.StringOutput { return v.NasIp }).(pulumi.StringOutput)
}

func (o VpnSslWebRealmOutput) RadiusPort() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnSslWebRealm) pulumi.IntOutput { return v.RadiusPort }).(pulumi.IntOutput)
}

func (o VpnSslWebRealmOutput) RadiusServer() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSslWebRealm) pulumi.StringOutput { return v.RadiusServer }).(pulumi.StringOutput)
}

func (o VpnSslWebRealmOutput) UrlPath() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSslWebRealm) pulumi.StringOutput { return v.UrlPath }).(pulumi.StringOutput)
}

func (o VpnSslWebRealmOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnSslWebRealm) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o VpnSslWebRealmOutput) VirtualHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnSslWebRealm) pulumi.StringPtrOutput { return v.VirtualHost }).(pulumi.StringPtrOutput)
}

func (o VpnSslWebRealmOutput) VirtualHostOnly() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSslWebRealm) pulumi.StringOutput { return v.VirtualHostOnly }).(pulumi.StringOutput)
}

func (o VpnSslWebRealmOutput) VirtualHostServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSslWebRealm) pulumi.StringOutput { return v.VirtualHostServerCert }).(pulumi.StringOutput)
}

type VpnSslWebRealmArrayOutput struct{ *pulumi.OutputState }

func (VpnSslWebRealmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnSslWebRealm)(nil)).Elem()
}

func (o VpnSslWebRealmArrayOutput) ToVpnSslWebRealmArrayOutput() VpnSslWebRealmArrayOutput {
	return o
}

func (o VpnSslWebRealmArrayOutput) ToVpnSslWebRealmArrayOutputWithContext(ctx context.Context) VpnSslWebRealmArrayOutput {
	return o
}

func (o VpnSslWebRealmArrayOutput) Index(i pulumi.IntInput) VpnSslWebRealmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnSslWebRealm {
		return vs[0].([]*VpnSslWebRealm)[vs[1].(int)]
	}).(VpnSslWebRealmOutput)
}

type VpnSslWebRealmMapOutput struct{ *pulumi.OutputState }

func (VpnSslWebRealmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnSslWebRealm)(nil)).Elem()
}

func (o VpnSslWebRealmMapOutput) ToVpnSslWebRealmMapOutput() VpnSslWebRealmMapOutput {
	return o
}

func (o VpnSslWebRealmMapOutput) ToVpnSslWebRealmMapOutputWithContext(ctx context.Context) VpnSslWebRealmMapOutput {
	return o
}

func (o VpnSslWebRealmMapOutput) MapIndex(k pulumi.StringInput) VpnSslWebRealmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnSslWebRealm {
		return vs[0].(map[string]*VpnSslWebRealm)[vs[1].(string)]
	}).(VpnSslWebRealmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnSslWebRealmInput)(nil)).Elem(), &VpnSslWebRealm{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnSslWebRealmArrayInput)(nil)).Elem(), VpnSslWebRealmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnSslWebRealmMapInput)(nil)).Elem(), VpnSslWebRealmMap{})
	pulumi.RegisterOutputType(VpnSslWebRealmOutput{})
	pulumi.RegisterOutputType(VpnSslWebRealmArrayOutput{})
	pulumi.RegisterOutputType(VpnSslWebRealmMapOutput{})
}

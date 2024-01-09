// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemGreTunnel struct {
	pulumi.CustomResourceState

	ChecksumReception          pulumi.StringOutput    `pulumi:"checksumReception"`
	ChecksumTransmission       pulumi.StringOutput    `pulumi:"checksumTransmission"`
	Diffservcode               pulumi.StringOutput    `pulumi:"diffservcode"`
	DscpCopying                pulumi.StringOutput    `pulumi:"dscpCopying"`
	Interface                  pulumi.StringOutput    `pulumi:"interface"`
	IpVersion                  pulumi.StringOutput    `pulumi:"ipVersion"`
	KeepaliveFailtimes         pulumi.IntOutput       `pulumi:"keepaliveFailtimes"`
	KeepaliveInterval          pulumi.IntOutput       `pulumi:"keepaliveInterval"`
	KeyInbound                 pulumi.IntOutput       `pulumi:"keyInbound"`
	KeyOutbound                pulumi.IntOutput       `pulumi:"keyOutbound"`
	LocalGw                    pulumi.StringOutput    `pulumi:"localGw"`
	LocalGw6                   pulumi.StringOutput    `pulumi:"localGw6"`
	Name                       pulumi.StringOutput    `pulumi:"name"`
	RemoteGw                   pulumi.StringOutput    `pulumi:"remoteGw"`
	RemoteGw6                  pulumi.StringOutput    `pulumi:"remoteGw6"`
	SequenceNumberReception    pulumi.StringOutput    `pulumi:"sequenceNumberReception"`
	SequenceNumberTransmission pulumi.StringOutput    `pulumi:"sequenceNumberTransmission"`
	UseSdwan                   pulumi.StringOutput    `pulumi:"useSdwan"`
	Vdomparam                  pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemGreTunnel registers a new resource with the given unique name, arguments, and options.
func NewSystemGreTunnel(ctx *pulumi.Context,
	name string, args *SystemGreTunnelArgs, opts ...pulumi.ResourceOption) (*SystemGreTunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocalGw == nil {
		return nil, errors.New("invalid value for required argument 'LocalGw'")
	}
	if args.RemoteGw == nil {
		return nil, errors.New("invalid value for required argument 'RemoteGw'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemGreTunnel
	err := ctx.RegisterResource("fortios:index/systemGreTunnel:SystemGreTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemGreTunnel gets an existing SystemGreTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemGreTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemGreTunnelState, opts ...pulumi.ResourceOption) (*SystemGreTunnel, error) {
	var resource SystemGreTunnel
	err := ctx.ReadResource("fortios:index/systemGreTunnel:SystemGreTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemGreTunnel resources.
type systemGreTunnelState struct {
	ChecksumReception          *string `pulumi:"checksumReception"`
	ChecksumTransmission       *string `pulumi:"checksumTransmission"`
	Diffservcode               *string `pulumi:"diffservcode"`
	DscpCopying                *string `pulumi:"dscpCopying"`
	Interface                  *string `pulumi:"interface"`
	IpVersion                  *string `pulumi:"ipVersion"`
	KeepaliveFailtimes         *int    `pulumi:"keepaliveFailtimes"`
	KeepaliveInterval          *int    `pulumi:"keepaliveInterval"`
	KeyInbound                 *int    `pulumi:"keyInbound"`
	KeyOutbound                *int    `pulumi:"keyOutbound"`
	LocalGw                    *string `pulumi:"localGw"`
	LocalGw6                   *string `pulumi:"localGw6"`
	Name                       *string `pulumi:"name"`
	RemoteGw                   *string `pulumi:"remoteGw"`
	RemoteGw6                  *string `pulumi:"remoteGw6"`
	SequenceNumberReception    *string `pulumi:"sequenceNumberReception"`
	SequenceNumberTransmission *string `pulumi:"sequenceNumberTransmission"`
	UseSdwan                   *string `pulumi:"useSdwan"`
	Vdomparam                  *string `pulumi:"vdomparam"`
}

type SystemGreTunnelState struct {
	ChecksumReception          pulumi.StringPtrInput
	ChecksumTransmission       pulumi.StringPtrInput
	Diffservcode               pulumi.StringPtrInput
	DscpCopying                pulumi.StringPtrInput
	Interface                  pulumi.StringPtrInput
	IpVersion                  pulumi.StringPtrInput
	KeepaliveFailtimes         pulumi.IntPtrInput
	KeepaliveInterval          pulumi.IntPtrInput
	KeyInbound                 pulumi.IntPtrInput
	KeyOutbound                pulumi.IntPtrInput
	LocalGw                    pulumi.StringPtrInput
	LocalGw6                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	RemoteGw                   pulumi.StringPtrInput
	RemoteGw6                  pulumi.StringPtrInput
	SequenceNumberReception    pulumi.StringPtrInput
	SequenceNumberTransmission pulumi.StringPtrInput
	UseSdwan                   pulumi.StringPtrInput
	Vdomparam                  pulumi.StringPtrInput
}

func (SystemGreTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGreTunnelState)(nil)).Elem()
}

type systemGreTunnelArgs struct {
	ChecksumReception          *string `pulumi:"checksumReception"`
	ChecksumTransmission       *string `pulumi:"checksumTransmission"`
	Diffservcode               *string `pulumi:"diffservcode"`
	DscpCopying                *string `pulumi:"dscpCopying"`
	Interface                  *string `pulumi:"interface"`
	IpVersion                  *string `pulumi:"ipVersion"`
	KeepaliveFailtimes         *int    `pulumi:"keepaliveFailtimes"`
	KeepaliveInterval          *int    `pulumi:"keepaliveInterval"`
	KeyInbound                 *int    `pulumi:"keyInbound"`
	KeyOutbound                *int    `pulumi:"keyOutbound"`
	LocalGw                    string  `pulumi:"localGw"`
	LocalGw6                   *string `pulumi:"localGw6"`
	Name                       *string `pulumi:"name"`
	RemoteGw                   string  `pulumi:"remoteGw"`
	RemoteGw6                  *string `pulumi:"remoteGw6"`
	SequenceNumberReception    *string `pulumi:"sequenceNumberReception"`
	SequenceNumberTransmission *string `pulumi:"sequenceNumberTransmission"`
	UseSdwan                   *string `pulumi:"useSdwan"`
	Vdomparam                  *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemGreTunnel resource.
type SystemGreTunnelArgs struct {
	ChecksumReception          pulumi.StringPtrInput
	ChecksumTransmission       pulumi.StringPtrInput
	Diffservcode               pulumi.StringPtrInput
	DscpCopying                pulumi.StringPtrInput
	Interface                  pulumi.StringPtrInput
	IpVersion                  pulumi.StringPtrInput
	KeepaliveFailtimes         pulumi.IntPtrInput
	KeepaliveInterval          pulumi.IntPtrInput
	KeyInbound                 pulumi.IntPtrInput
	KeyOutbound                pulumi.IntPtrInput
	LocalGw                    pulumi.StringInput
	LocalGw6                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	RemoteGw                   pulumi.StringInput
	RemoteGw6                  pulumi.StringPtrInput
	SequenceNumberReception    pulumi.StringPtrInput
	SequenceNumberTransmission pulumi.StringPtrInput
	UseSdwan                   pulumi.StringPtrInput
	Vdomparam                  pulumi.StringPtrInput
}

func (SystemGreTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGreTunnelArgs)(nil)).Elem()
}

type SystemGreTunnelInput interface {
	pulumi.Input

	ToSystemGreTunnelOutput() SystemGreTunnelOutput
	ToSystemGreTunnelOutputWithContext(ctx context.Context) SystemGreTunnelOutput
}

func (*SystemGreTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGreTunnel)(nil)).Elem()
}

func (i *SystemGreTunnel) ToSystemGreTunnelOutput() SystemGreTunnelOutput {
	return i.ToSystemGreTunnelOutputWithContext(context.Background())
}

func (i *SystemGreTunnel) ToSystemGreTunnelOutputWithContext(ctx context.Context) SystemGreTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGreTunnelOutput)
}

// SystemGreTunnelArrayInput is an input type that accepts SystemGreTunnelArray and SystemGreTunnelArrayOutput values.
// You can construct a concrete instance of `SystemGreTunnelArrayInput` via:
//
//	SystemGreTunnelArray{ SystemGreTunnelArgs{...} }
type SystemGreTunnelArrayInput interface {
	pulumi.Input

	ToSystemGreTunnelArrayOutput() SystemGreTunnelArrayOutput
	ToSystemGreTunnelArrayOutputWithContext(context.Context) SystemGreTunnelArrayOutput
}

type SystemGreTunnelArray []SystemGreTunnelInput

func (SystemGreTunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemGreTunnel)(nil)).Elem()
}

func (i SystemGreTunnelArray) ToSystemGreTunnelArrayOutput() SystemGreTunnelArrayOutput {
	return i.ToSystemGreTunnelArrayOutputWithContext(context.Background())
}

func (i SystemGreTunnelArray) ToSystemGreTunnelArrayOutputWithContext(ctx context.Context) SystemGreTunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGreTunnelArrayOutput)
}

// SystemGreTunnelMapInput is an input type that accepts SystemGreTunnelMap and SystemGreTunnelMapOutput values.
// You can construct a concrete instance of `SystemGreTunnelMapInput` via:
//
//	SystemGreTunnelMap{ "key": SystemGreTunnelArgs{...} }
type SystemGreTunnelMapInput interface {
	pulumi.Input

	ToSystemGreTunnelMapOutput() SystemGreTunnelMapOutput
	ToSystemGreTunnelMapOutputWithContext(context.Context) SystemGreTunnelMapOutput
}

type SystemGreTunnelMap map[string]SystemGreTunnelInput

func (SystemGreTunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemGreTunnel)(nil)).Elem()
}

func (i SystemGreTunnelMap) ToSystemGreTunnelMapOutput() SystemGreTunnelMapOutput {
	return i.ToSystemGreTunnelMapOutputWithContext(context.Background())
}

func (i SystemGreTunnelMap) ToSystemGreTunnelMapOutputWithContext(ctx context.Context) SystemGreTunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGreTunnelMapOutput)
}

type SystemGreTunnelOutput struct{ *pulumi.OutputState }

func (SystemGreTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGreTunnel)(nil)).Elem()
}

func (o SystemGreTunnelOutput) ToSystemGreTunnelOutput() SystemGreTunnelOutput {
	return o
}

func (o SystemGreTunnelOutput) ToSystemGreTunnelOutputWithContext(ctx context.Context) SystemGreTunnelOutput {
	return o
}

func (o SystemGreTunnelOutput) ChecksumReception() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.ChecksumReception }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) ChecksumTransmission() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.ChecksumTransmission }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) Diffservcode() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.Diffservcode }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) DscpCopying() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.DscpCopying }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.IpVersion }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) KeepaliveFailtimes() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.IntOutput { return v.KeepaliveFailtimes }).(pulumi.IntOutput)
}

func (o SystemGreTunnelOutput) KeepaliveInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.IntOutput { return v.KeepaliveInterval }).(pulumi.IntOutput)
}

func (o SystemGreTunnelOutput) KeyInbound() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.IntOutput { return v.KeyInbound }).(pulumi.IntOutput)
}

func (o SystemGreTunnelOutput) KeyOutbound() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.IntOutput { return v.KeyOutbound }).(pulumi.IntOutput)
}

func (o SystemGreTunnelOutput) LocalGw() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.LocalGw }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) LocalGw6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.LocalGw6 }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) RemoteGw() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.RemoteGw }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) RemoteGw6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.RemoteGw6 }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) SequenceNumberReception() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.SequenceNumberReception }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) SequenceNumberTransmission() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.SequenceNumberTransmission }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) UseSdwan() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringOutput { return v.UseSdwan }).(pulumi.StringOutput)
}

func (o SystemGreTunnelOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemGreTunnel) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemGreTunnelArrayOutput struct{ *pulumi.OutputState }

func (SystemGreTunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemGreTunnel)(nil)).Elem()
}

func (o SystemGreTunnelArrayOutput) ToSystemGreTunnelArrayOutput() SystemGreTunnelArrayOutput {
	return o
}

func (o SystemGreTunnelArrayOutput) ToSystemGreTunnelArrayOutputWithContext(ctx context.Context) SystemGreTunnelArrayOutput {
	return o
}

func (o SystemGreTunnelArrayOutput) Index(i pulumi.IntInput) SystemGreTunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemGreTunnel {
		return vs[0].([]*SystemGreTunnel)[vs[1].(int)]
	}).(SystemGreTunnelOutput)
}

type SystemGreTunnelMapOutput struct{ *pulumi.OutputState }

func (SystemGreTunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemGreTunnel)(nil)).Elem()
}

func (o SystemGreTunnelMapOutput) ToSystemGreTunnelMapOutput() SystemGreTunnelMapOutput {
	return o
}

func (o SystemGreTunnelMapOutput) ToSystemGreTunnelMapOutputWithContext(ctx context.Context) SystemGreTunnelMapOutput {
	return o
}

func (o SystemGreTunnelMapOutput) MapIndex(k pulumi.StringInput) SystemGreTunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemGreTunnel {
		return vs[0].(map[string]*SystemGreTunnel)[vs[1].(string)]
	}).(SystemGreTunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGreTunnelInput)(nil)).Elem(), &SystemGreTunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGreTunnelArrayInput)(nil)).Elem(), SystemGreTunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGreTunnelMapInput)(nil)).Elem(), SystemGreTunnelMap{})
	pulumi.RegisterOutputType(SystemGreTunnelOutput{})
	pulumi.RegisterOutputType(SystemGreTunnelArrayOutput{})
	pulumi.RegisterOutputType(SystemGreTunnelMapOutput{})
}

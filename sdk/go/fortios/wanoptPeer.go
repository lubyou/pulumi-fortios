// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WanoptPeer struct {
	pulumi.CustomResourceState

	Ip         pulumi.StringOutput    `pulumi:"ip"`
	PeerHostId pulumi.StringOutput    `pulumi:"peerHostId"`
	Vdomparam  pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWanoptPeer registers a new resource with the given unique name, arguments, and options.
func NewWanoptPeer(ctx *pulumi.Context,
	name string, args *WanoptPeerArgs, opts ...pulumi.ResourceOption) (*WanoptPeer, error) {
	if args == nil {
		args = &WanoptPeerArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WanoptPeer
	err := ctx.RegisterResource("fortios:index/wanoptPeer:WanoptPeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWanoptPeer gets an existing WanoptPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWanoptPeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WanoptPeerState, opts ...pulumi.ResourceOption) (*WanoptPeer, error) {
	var resource WanoptPeer
	err := ctx.ReadResource("fortios:index/wanoptPeer:WanoptPeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WanoptPeer resources.
type wanoptPeerState struct {
	Ip         *string `pulumi:"ip"`
	PeerHostId *string `pulumi:"peerHostId"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

type WanoptPeerState struct {
	Ip         pulumi.StringPtrInput
	PeerHostId pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (WanoptPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptPeerState)(nil)).Elem()
}

type wanoptPeerArgs struct {
	Ip         *string `pulumi:"ip"`
	PeerHostId *string `pulumi:"peerHostId"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WanoptPeer resource.
type WanoptPeerArgs struct {
	Ip         pulumi.StringPtrInput
	PeerHostId pulumi.StringPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (WanoptPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptPeerArgs)(nil)).Elem()
}

type WanoptPeerInput interface {
	pulumi.Input

	ToWanoptPeerOutput() WanoptPeerOutput
	ToWanoptPeerOutputWithContext(ctx context.Context) WanoptPeerOutput
}

func (*WanoptPeer) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptPeer)(nil)).Elem()
}

func (i *WanoptPeer) ToWanoptPeerOutput() WanoptPeerOutput {
	return i.ToWanoptPeerOutputWithContext(context.Background())
}

func (i *WanoptPeer) ToWanoptPeerOutputWithContext(ctx context.Context) WanoptPeerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptPeerOutput)
}

// WanoptPeerArrayInput is an input type that accepts WanoptPeerArray and WanoptPeerArrayOutput values.
// You can construct a concrete instance of `WanoptPeerArrayInput` via:
//
//	WanoptPeerArray{ WanoptPeerArgs{...} }
type WanoptPeerArrayInput interface {
	pulumi.Input

	ToWanoptPeerArrayOutput() WanoptPeerArrayOutput
	ToWanoptPeerArrayOutputWithContext(context.Context) WanoptPeerArrayOutput
}

type WanoptPeerArray []WanoptPeerInput

func (WanoptPeerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WanoptPeer)(nil)).Elem()
}

func (i WanoptPeerArray) ToWanoptPeerArrayOutput() WanoptPeerArrayOutput {
	return i.ToWanoptPeerArrayOutputWithContext(context.Background())
}

func (i WanoptPeerArray) ToWanoptPeerArrayOutputWithContext(ctx context.Context) WanoptPeerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptPeerArrayOutput)
}

// WanoptPeerMapInput is an input type that accepts WanoptPeerMap and WanoptPeerMapOutput values.
// You can construct a concrete instance of `WanoptPeerMapInput` via:
//
//	WanoptPeerMap{ "key": WanoptPeerArgs{...} }
type WanoptPeerMapInput interface {
	pulumi.Input

	ToWanoptPeerMapOutput() WanoptPeerMapOutput
	ToWanoptPeerMapOutputWithContext(context.Context) WanoptPeerMapOutput
}

type WanoptPeerMap map[string]WanoptPeerInput

func (WanoptPeerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WanoptPeer)(nil)).Elem()
}

func (i WanoptPeerMap) ToWanoptPeerMapOutput() WanoptPeerMapOutput {
	return i.ToWanoptPeerMapOutputWithContext(context.Background())
}

func (i WanoptPeerMap) ToWanoptPeerMapOutputWithContext(ctx context.Context) WanoptPeerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptPeerMapOutput)
}

type WanoptPeerOutput struct{ *pulumi.OutputState }

func (WanoptPeerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptPeer)(nil)).Elem()
}

func (o WanoptPeerOutput) ToWanoptPeerOutput() WanoptPeerOutput {
	return o
}

func (o WanoptPeerOutput) ToWanoptPeerOutputWithContext(ctx context.Context) WanoptPeerOutput {
	return o
}

func (o WanoptPeerOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptPeer) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

func (o WanoptPeerOutput) PeerHostId() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptPeer) pulumi.StringOutput { return v.PeerHostId }).(pulumi.StringOutput)
}

func (o WanoptPeerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WanoptPeer) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WanoptPeerArrayOutput struct{ *pulumi.OutputState }

func (WanoptPeerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WanoptPeer)(nil)).Elem()
}

func (o WanoptPeerArrayOutput) ToWanoptPeerArrayOutput() WanoptPeerArrayOutput {
	return o
}

func (o WanoptPeerArrayOutput) ToWanoptPeerArrayOutputWithContext(ctx context.Context) WanoptPeerArrayOutput {
	return o
}

func (o WanoptPeerArrayOutput) Index(i pulumi.IntInput) WanoptPeerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WanoptPeer {
		return vs[0].([]*WanoptPeer)[vs[1].(int)]
	}).(WanoptPeerOutput)
}

type WanoptPeerMapOutput struct{ *pulumi.OutputState }

func (WanoptPeerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WanoptPeer)(nil)).Elem()
}

func (o WanoptPeerMapOutput) ToWanoptPeerMapOutput() WanoptPeerMapOutput {
	return o
}

func (o WanoptPeerMapOutput) ToWanoptPeerMapOutputWithContext(ctx context.Context) WanoptPeerMapOutput {
	return o
}

func (o WanoptPeerMapOutput) MapIndex(k pulumi.StringInput) WanoptPeerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WanoptPeer {
		return vs[0].(map[string]*WanoptPeer)[vs[1].(string)]
	}).(WanoptPeerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptPeerInput)(nil)).Elem(), &WanoptPeer{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptPeerArrayInput)(nil)).Elem(), WanoptPeerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptPeerMapInput)(nil)).Elem(), WanoptPeerMap{})
	pulumi.RegisterOutputType(WanoptPeerOutput{})
	pulumi.RegisterOutputType(WanoptPeerArrayOutput{})
	pulumi.RegisterOutputType(WanoptPeerMapOutput{})
}

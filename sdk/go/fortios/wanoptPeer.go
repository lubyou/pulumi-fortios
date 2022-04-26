// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure WAN optimization peers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewWanoptPeer(ctx, "trname", &fortios.WanoptPeerArgs{
// 			Ip:         pulumi.String("1.1.1.1"),
// 			PeerHostId: pulumi.String("1"),
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
// Wanopt Peer can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/wanoptPeer:WanoptPeer labelname {{peer_host_id}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/wanoptPeer:WanoptPeer labelname {{peer_host_id}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WanoptPeer struct {
	pulumi.CustomResourceState

	// Peer IP address.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Peer host ID.
	PeerHostId pulumi.StringOutput `pulumi:"peerHostId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
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
	// Peer IP address.
	Ip *string `pulumi:"ip"`
	// Peer host ID.
	PeerHostId *string `pulumi:"peerHostId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WanoptPeerState struct {
	// Peer IP address.
	Ip pulumi.StringPtrInput
	// Peer host ID.
	PeerHostId pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptPeerState)(nil)).Elem()
}

type wanoptPeerArgs struct {
	// Peer IP address.
	Ip *string `pulumi:"ip"`
	// Peer host ID.
	PeerHostId *string `pulumi:"peerHostId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WanoptPeer resource.
type WanoptPeerArgs struct {
	// Peer IP address.
	Ip pulumi.StringPtrInput
	// Peer host ID.
	PeerHostId pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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
//          WanoptPeerArray{ WanoptPeerArgs{...} }
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
//          WanoptPeerMap{ "key": WanoptPeerArgs{...} }
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

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Designate cache-service for wan-optimization and webcache.
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
// 		_, err := fortios.NewWanoptCacheService(ctx, "trname", &fortios.WanoptCacheServiceArgs{
// 			AcceptableConnections: pulumi.String("any"),
// 			Collaboration:         pulumi.String("disable"),
// 			DeviceId:              pulumi.String("default_dev_id"),
// 			PreferScenario:        pulumi.String("balance"),
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
// Wanopt CacheService can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wanoptCacheService:WanoptCacheService labelname WanoptCacheService
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WanoptCacheService struct {
	pulumi.CustomResourceState

	// Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
	AcceptableConnections pulumi.StringOutput `pulumi:"acceptableConnections"`
	// Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
	Collaboration pulumi.StringOutput `pulumi:"collaboration"`
	// Device ID of this peer.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// Modify cache-service destination peer list. The structure of `dstPeer` block is documented below.
	DstPeers WanoptCacheServiceDstPeerArrayOutput `pulumi:"dstPeers"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
	PreferScenario pulumi.StringOutput `pulumi:"preferScenario"`
	// Modify cache-service source peer list. The structure of `srcPeer` block is documented below.
	SrcPeers WanoptCacheServiceSrcPeerArrayOutput `pulumi:"srcPeers"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWanoptCacheService registers a new resource with the given unique name, arguments, and options.
func NewWanoptCacheService(ctx *pulumi.Context,
	name string, args *WanoptCacheServiceArgs, opts ...pulumi.ResourceOption) (*WanoptCacheService, error) {
	if args == nil {
		args = &WanoptCacheServiceArgs{}
	}

	var resource WanoptCacheService
	err := ctx.RegisterResource("fortios:index/wanoptCacheService:WanoptCacheService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWanoptCacheService gets an existing WanoptCacheService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWanoptCacheService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WanoptCacheServiceState, opts ...pulumi.ResourceOption) (*WanoptCacheService, error) {
	var resource WanoptCacheService
	err := ctx.ReadResource("fortios:index/wanoptCacheService:WanoptCacheService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WanoptCacheService resources.
type wanoptCacheServiceState struct {
	// Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
	AcceptableConnections *string `pulumi:"acceptableConnections"`
	// Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
	Collaboration *string `pulumi:"collaboration"`
	// Device ID of this peer.
	DeviceId *string `pulumi:"deviceId"`
	// Modify cache-service destination peer list. The structure of `dstPeer` block is documented below.
	DstPeers []WanoptCacheServiceDstPeer `pulumi:"dstPeers"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
	PreferScenario *string `pulumi:"preferScenario"`
	// Modify cache-service source peer list. The structure of `srcPeer` block is documented below.
	SrcPeers []WanoptCacheServiceSrcPeer `pulumi:"srcPeers"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WanoptCacheServiceState struct {
	// Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
	AcceptableConnections pulumi.StringPtrInput
	// Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
	Collaboration pulumi.StringPtrInput
	// Device ID of this peer.
	DeviceId pulumi.StringPtrInput
	// Modify cache-service destination peer list. The structure of `dstPeer` block is documented below.
	DstPeers WanoptCacheServiceDstPeerArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
	PreferScenario pulumi.StringPtrInput
	// Modify cache-service source peer list. The structure of `srcPeer` block is documented below.
	SrcPeers WanoptCacheServiceSrcPeerArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptCacheServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptCacheServiceState)(nil)).Elem()
}

type wanoptCacheServiceArgs struct {
	// Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
	AcceptableConnections *string `pulumi:"acceptableConnections"`
	// Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
	Collaboration *string `pulumi:"collaboration"`
	// Device ID of this peer.
	DeviceId *string `pulumi:"deviceId"`
	// Modify cache-service destination peer list. The structure of `dstPeer` block is documented below.
	DstPeers []WanoptCacheServiceDstPeer `pulumi:"dstPeers"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
	PreferScenario *string `pulumi:"preferScenario"`
	// Modify cache-service source peer list. The structure of `srcPeer` block is documented below.
	SrcPeers []WanoptCacheServiceSrcPeer `pulumi:"srcPeers"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WanoptCacheService resource.
type WanoptCacheServiceArgs struct {
	// Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
	AcceptableConnections pulumi.StringPtrInput
	// Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
	Collaboration pulumi.StringPtrInput
	// Device ID of this peer.
	DeviceId pulumi.StringPtrInput
	// Modify cache-service destination peer list. The structure of `dstPeer` block is documented below.
	DstPeers WanoptCacheServiceDstPeerArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
	PreferScenario pulumi.StringPtrInput
	// Modify cache-service source peer list. The structure of `srcPeer` block is documented below.
	SrcPeers WanoptCacheServiceSrcPeerArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptCacheServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptCacheServiceArgs)(nil)).Elem()
}

type WanoptCacheServiceInput interface {
	pulumi.Input

	ToWanoptCacheServiceOutput() WanoptCacheServiceOutput
	ToWanoptCacheServiceOutputWithContext(ctx context.Context) WanoptCacheServiceOutput
}

func (*WanoptCacheService) ElementType() reflect.Type {
	return reflect.TypeOf((*WanoptCacheService)(nil))
}

func (i *WanoptCacheService) ToWanoptCacheServiceOutput() WanoptCacheServiceOutput {
	return i.ToWanoptCacheServiceOutputWithContext(context.Background())
}

func (i *WanoptCacheService) ToWanoptCacheServiceOutputWithContext(ctx context.Context) WanoptCacheServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptCacheServiceOutput)
}

func (i *WanoptCacheService) ToWanoptCacheServicePtrOutput() WanoptCacheServicePtrOutput {
	return i.ToWanoptCacheServicePtrOutputWithContext(context.Background())
}

func (i *WanoptCacheService) ToWanoptCacheServicePtrOutputWithContext(ctx context.Context) WanoptCacheServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptCacheServicePtrOutput)
}

type WanoptCacheServicePtrInput interface {
	pulumi.Input

	ToWanoptCacheServicePtrOutput() WanoptCacheServicePtrOutput
	ToWanoptCacheServicePtrOutputWithContext(ctx context.Context) WanoptCacheServicePtrOutput
}

type wanoptCacheServicePtrType WanoptCacheServiceArgs

func (*wanoptCacheServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptCacheService)(nil))
}

func (i *wanoptCacheServicePtrType) ToWanoptCacheServicePtrOutput() WanoptCacheServicePtrOutput {
	return i.ToWanoptCacheServicePtrOutputWithContext(context.Background())
}

func (i *wanoptCacheServicePtrType) ToWanoptCacheServicePtrOutputWithContext(ctx context.Context) WanoptCacheServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptCacheServicePtrOutput)
}

// WanoptCacheServiceArrayInput is an input type that accepts WanoptCacheServiceArray and WanoptCacheServiceArrayOutput values.
// You can construct a concrete instance of `WanoptCacheServiceArrayInput` via:
//
//          WanoptCacheServiceArray{ WanoptCacheServiceArgs{...} }
type WanoptCacheServiceArrayInput interface {
	pulumi.Input

	ToWanoptCacheServiceArrayOutput() WanoptCacheServiceArrayOutput
	ToWanoptCacheServiceArrayOutputWithContext(context.Context) WanoptCacheServiceArrayOutput
}

type WanoptCacheServiceArray []WanoptCacheServiceInput

func (WanoptCacheServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WanoptCacheService)(nil))
}

func (i WanoptCacheServiceArray) ToWanoptCacheServiceArrayOutput() WanoptCacheServiceArrayOutput {
	return i.ToWanoptCacheServiceArrayOutputWithContext(context.Background())
}

func (i WanoptCacheServiceArray) ToWanoptCacheServiceArrayOutputWithContext(ctx context.Context) WanoptCacheServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptCacheServiceArrayOutput)
}

// WanoptCacheServiceMapInput is an input type that accepts WanoptCacheServiceMap and WanoptCacheServiceMapOutput values.
// You can construct a concrete instance of `WanoptCacheServiceMapInput` via:
//
//          WanoptCacheServiceMap{ "key": WanoptCacheServiceArgs{...} }
type WanoptCacheServiceMapInput interface {
	pulumi.Input

	ToWanoptCacheServiceMapOutput() WanoptCacheServiceMapOutput
	ToWanoptCacheServiceMapOutputWithContext(context.Context) WanoptCacheServiceMapOutput
}

type WanoptCacheServiceMap map[string]WanoptCacheServiceInput

func (WanoptCacheServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WanoptCacheService)(nil))
}

func (i WanoptCacheServiceMap) ToWanoptCacheServiceMapOutput() WanoptCacheServiceMapOutput {
	return i.ToWanoptCacheServiceMapOutputWithContext(context.Background())
}

func (i WanoptCacheServiceMap) ToWanoptCacheServiceMapOutputWithContext(ctx context.Context) WanoptCacheServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptCacheServiceMapOutput)
}

type WanoptCacheServiceOutput struct {
	*pulumi.OutputState
}

func (WanoptCacheServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WanoptCacheService)(nil))
}

func (o WanoptCacheServiceOutput) ToWanoptCacheServiceOutput() WanoptCacheServiceOutput {
	return o
}

func (o WanoptCacheServiceOutput) ToWanoptCacheServiceOutputWithContext(ctx context.Context) WanoptCacheServiceOutput {
	return o
}

func (o WanoptCacheServiceOutput) ToWanoptCacheServicePtrOutput() WanoptCacheServicePtrOutput {
	return o.ToWanoptCacheServicePtrOutputWithContext(context.Background())
}

func (o WanoptCacheServiceOutput) ToWanoptCacheServicePtrOutputWithContext(ctx context.Context) WanoptCacheServicePtrOutput {
	return o.ApplyT(func(v WanoptCacheService) *WanoptCacheService {
		return &v
	}).(WanoptCacheServicePtrOutput)
}

type WanoptCacheServicePtrOutput struct {
	*pulumi.OutputState
}

func (WanoptCacheServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptCacheService)(nil))
}

func (o WanoptCacheServicePtrOutput) ToWanoptCacheServicePtrOutput() WanoptCacheServicePtrOutput {
	return o
}

func (o WanoptCacheServicePtrOutput) ToWanoptCacheServicePtrOutputWithContext(ctx context.Context) WanoptCacheServicePtrOutput {
	return o
}

type WanoptCacheServiceArrayOutput struct{ *pulumi.OutputState }

func (WanoptCacheServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WanoptCacheService)(nil))
}

func (o WanoptCacheServiceArrayOutput) ToWanoptCacheServiceArrayOutput() WanoptCacheServiceArrayOutput {
	return o
}

func (o WanoptCacheServiceArrayOutput) ToWanoptCacheServiceArrayOutputWithContext(ctx context.Context) WanoptCacheServiceArrayOutput {
	return o
}

func (o WanoptCacheServiceArrayOutput) Index(i pulumi.IntInput) WanoptCacheServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WanoptCacheService {
		return vs[0].([]WanoptCacheService)[vs[1].(int)]
	}).(WanoptCacheServiceOutput)
}

type WanoptCacheServiceMapOutput struct{ *pulumi.OutputState }

func (WanoptCacheServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WanoptCacheService)(nil))
}

func (o WanoptCacheServiceMapOutput) ToWanoptCacheServiceMapOutput() WanoptCacheServiceMapOutput {
	return o
}

func (o WanoptCacheServiceMapOutput) ToWanoptCacheServiceMapOutputWithContext(ctx context.Context) WanoptCacheServiceMapOutput {
	return o
}

func (o WanoptCacheServiceMapOutput) MapIndex(k pulumi.StringInput) WanoptCacheServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WanoptCacheService {
		return vs[0].(map[string]WanoptCacheService)[vs[1].(string)]
	}).(WanoptCacheServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(WanoptCacheServiceOutput{})
	pulumi.RegisterOutputType(WanoptCacheServicePtrOutput{})
	pulumi.RegisterOutputType(WanoptCacheServiceArrayOutput{})
	pulumi.RegisterOutputType(WanoptCacheServiceMapOutput{})
}

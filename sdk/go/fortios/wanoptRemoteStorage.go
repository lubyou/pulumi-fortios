// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure a remote cache device as Web cache storage.
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
// 		_, err := fortios.NewWanoptRemoteStorage(ctx, "trname", &fortios.WanoptRemoteStorageArgs{
// 			RemoteCacheIp: pulumi.String("0.0.0.0"),
// 			Status:        pulumi.String("disable"),
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
// Wanopt RemoteStorage can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wanoptRemoteStorage:WanoptRemoteStorage labelname WanoptRemoteStorage
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WanoptRemoteStorage struct {
	pulumi.CustomResourceState

	// ID that this device uses to connect to the remote device.
	LocalCacheId pulumi.StringOutput `pulumi:"localCacheId"`
	// ID of the remote device to which the device connects.
	RemoteCacheId pulumi.StringOutput `pulumi:"remoteCacheId"`
	// IP address of the remote device to which the device connects.
	RemoteCacheIp pulumi.StringOutput `pulumi:"remoteCacheIp"`
	// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWanoptRemoteStorage registers a new resource with the given unique name, arguments, and options.
func NewWanoptRemoteStorage(ctx *pulumi.Context,
	name string, args *WanoptRemoteStorageArgs, opts ...pulumi.ResourceOption) (*WanoptRemoteStorage, error) {
	if args == nil {
		args = &WanoptRemoteStorageArgs{}
	}

	var resource WanoptRemoteStorage
	err := ctx.RegisterResource("fortios:index/wanoptRemoteStorage:WanoptRemoteStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWanoptRemoteStorage gets an existing WanoptRemoteStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWanoptRemoteStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WanoptRemoteStorageState, opts ...pulumi.ResourceOption) (*WanoptRemoteStorage, error) {
	var resource WanoptRemoteStorage
	err := ctx.ReadResource("fortios:index/wanoptRemoteStorage:WanoptRemoteStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WanoptRemoteStorage resources.
type wanoptRemoteStorageState struct {
	// ID that this device uses to connect to the remote device.
	LocalCacheId *string `pulumi:"localCacheId"`
	// ID of the remote device to which the device connects.
	RemoteCacheId *string `pulumi:"remoteCacheId"`
	// IP address of the remote device to which the device connects.
	RemoteCacheIp *string `pulumi:"remoteCacheIp"`
	// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WanoptRemoteStorageState struct {
	// ID that this device uses to connect to the remote device.
	LocalCacheId pulumi.StringPtrInput
	// ID of the remote device to which the device connects.
	RemoteCacheId pulumi.StringPtrInput
	// IP address of the remote device to which the device connects.
	RemoteCacheIp pulumi.StringPtrInput
	// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptRemoteStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptRemoteStorageState)(nil)).Elem()
}

type wanoptRemoteStorageArgs struct {
	// ID that this device uses to connect to the remote device.
	LocalCacheId *string `pulumi:"localCacheId"`
	// ID of the remote device to which the device connects.
	RemoteCacheId *string `pulumi:"remoteCacheId"`
	// IP address of the remote device to which the device connects.
	RemoteCacheIp *string `pulumi:"remoteCacheIp"`
	// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WanoptRemoteStorage resource.
type WanoptRemoteStorageArgs struct {
	// ID that this device uses to connect to the remote device.
	LocalCacheId pulumi.StringPtrInput
	// ID of the remote device to which the device connects.
	RemoteCacheId pulumi.StringPtrInput
	// IP address of the remote device to which the device connects.
	RemoteCacheIp pulumi.StringPtrInput
	// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptRemoteStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptRemoteStorageArgs)(nil)).Elem()
}

type WanoptRemoteStorageInput interface {
	pulumi.Input

	ToWanoptRemoteStorageOutput() WanoptRemoteStorageOutput
	ToWanoptRemoteStorageOutputWithContext(ctx context.Context) WanoptRemoteStorageOutput
}

func (*WanoptRemoteStorage) ElementType() reflect.Type {
	return reflect.TypeOf((*WanoptRemoteStorage)(nil))
}

func (i *WanoptRemoteStorage) ToWanoptRemoteStorageOutput() WanoptRemoteStorageOutput {
	return i.ToWanoptRemoteStorageOutputWithContext(context.Background())
}

func (i *WanoptRemoteStorage) ToWanoptRemoteStorageOutputWithContext(ctx context.Context) WanoptRemoteStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptRemoteStorageOutput)
}

func (i *WanoptRemoteStorage) ToWanoptRemoteStoragePtrOutput() WanoptRemoteStoragePtrOutput {
	return i.ToWanoptRemoteStoragePtrOutputWithContext(context.Background())
}

func (i *WanoptRemoteStorage) ToWanoptRemoteStoragePtrOutputWithContext(ctx context.Context) WanoptRemoteStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptRemoteStoragePtrOutput)
}

type WanoptRemoteStoragePtrInput interface {
	pulumi.Input

	ToWanoptRemoteStoragePtrOutput() WanoptRemoteStoragePtrOutput
	ToWanoptRemoteStoragePtrOutputWithContext(ctx context.Context) WanoptRemoteStoragePtrOutput
}

type wanoptRemoteStoragePtrType WanoptRemoteStorageArgs

func (*wanoptRemoteStoragePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptRemoteStorage)(nil))
}

func (i *wanoptRemoteStoragePtrType) ToWanoptRemoteStoragePtrOutput() WanoptRemoteStoragePtrOutput {
	return i.ToWanoptRemoteStoragePtrOutputWithContext(context.Background())
}

func (i *wanoptRemoteStoragePtrType) ToWanoptRemoteStoragePtrOutputWithContext(ctx context.Context) WanoptRemoteStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptRemoteStoragePtrOutput)
}

// WanoptRemoteStorageArrayInput is an input type that accepts WanoptRemoteStorageArray and WanoptRemoteStorageArrayOutput values.
// You can construct a concrete instance of `WanoptRemoteStorageArrayInput` via:
//
//          WanoptRemoteStorageArray{ WanoptRemoteStorageArgs{...} }
type WanoptRemoteStorageArrayInput interface {
	pulumi.Input

	ToWanoptRemoteStorageArrayOutput() WanoptRemoteStorageArrayOutput
	ToWanoptRemoteStorageArrayOutputWithContext(context.Context) WanoptRemoteStorageArrayOutput
}

type WanoptRemoteStorageArray []WanoptRemoteStorageInput

func (WanoptRemoteStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WanoptRemoteStorage)(nil))
}

func (i WanoptRemoteStorageArray) ToWanoptRemoteStorageArrayOutput() WanoptRemoteStorageArrayOutput {
	return i.ToWanoptRemoteStorageArrayOutputWithContext(context.Background())
}

func (i WanoptRemoteStorageArray) ToWanoptRemoteStorageArrayOutputWithContext(ctx context.Context) WanoptRemoteStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptRemoteStorageArrayOutput)
}

// WanoptRemoteStorageMapInput is an input type that accepts WanoptRemoteStorageMap and WanoptRemoteStorageMapOutput values.
// You can construct a concrete instance of `WanoptRemoteStorageMapInput` via:
//
//          WanoptRemoteStorageMap{ "key": WanoptRemoteStorageArgs{...} }
type WanoptRemoteStorageMapInput interface {
	pulumi.Input

	ToWanoptRemoteStorageMapOutput() WanoptRemoteStorageMapOutput
	ToWanoptRemoteStorageMapOutputWithContext(context.Context) WanoptRemoteStorageMapOutput
}

type WanoptRemoteStorageMap map[string]WanoptRemoteStorageInput

func (WanoptRemoteStorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WanoptRemoteStorage)(nil))
}

func (i WanoptRemoteStorageMap) ToWanoptRemoteStorageMapOutput() WanoptRemoteStorageMapOutput {
	return i.ToWanoptRemoteStorageMapOutputWithContext(context.Background())
}

func (i WanoptRemoteStorageMap) ToWanoptRemoteStorageMapOutputWithContext(ctx context.Context) WanoptRemoteStorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptRemoteStorageMapOutput)
}

type WanoptRemoteStorageOutput struct {
	*pulumi.OutputState
}

func (WanoptRemoteStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WanoptRemoteStorage)(nil))
}

func (o WanoptRemoteStorageOutput) ToWanoptRemoteStorageOutput() WanoptRemoteStorageOutput {
	return o
}

func (o WanoptRemoteStorageOutput) ToWanoptRemoteStorageOutputWithContext(ctx context.Context) WanoptRemoteStorageOutput {
	return o
}

func (o WanoptRemoteStorageOutput) ToWanoptRemoteStoragePtrOutput() WanoptRemoteStoragePtrOutput {
	return o.ToWanoptRemoteStoragePtrOutputWithContext(context.Background())
}

func (o WanoptRemoteStorageOutput) ToWanoptRemoteStoragePtrOutputWithContext(ctx context.Context) WanoptRemoteStoragePtrOutput {
	return o.ApplyT(func(v WanoptRemoteStorage) *WanoptRemoteStorage {
		return &v
	}).(WanoptRemoteStoragePtrOutput)
}

type WanoptRemoteStoragePtrOutput struct {
	*pulumi.OutputState
}

func (WanoptRemoteStoragePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptRemoteStorage)(nil))
}

func (o WanoptRemoteStoragePtrOutput) ToWanoptRemoteStoragePtrOutput() WanoptRemoteStoragePtrOutput {
	return o
}

func (o WanoptRemoteStoragePtrOutput) ToWanoptRemoteStoragePtrOutputWithContext(ctx context.Context) WanoptRemoteStoragePtrOutput {
	return o
}

type WanoptRemoteStorageArrayOutput struct{ *pulumi.OutputState }

func (WanoptRemoteStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WanoptRemoteStorage)(nil))
}

func (o WanoptRemoteStorageArrayOutput) ToWanoptRemoteStorageArrayOutput() WanoptRemoteStorageArrayOutput {
	return o
}

func (o WanoptRemoteStorageArrayOutput) ToWanoptRemoteStorageArrayOutputWithContext(ctx context.Context) WanoptRemoteStorageArrayOutput {
	return o
}

func (o WanoptRemoteStorageArrayOutput) Index(i pulumi.IntInput) WanoptRemoteStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WanoptRemoteStorage {
		return vs[0].([]WanoptRemoteStorage)[vs[1].(int)]
	}).(WanoptRemoteStorageOutput)
}

type WanoptRemoteStorageMapOutput struct{ *pulumi.OutputState }

func (WanoptRemoteStorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WanoptRemoteStorage)(nil))
}

func (o WanoptRemoteStorageMapOutput) ToWanoptRemoteStorageMapOutput() WanoptRemoteStorageMapOutput {
	return o
}

func (o WanoptRemoteStorageMapOutput) ToWanoptRemoteStorageMapOutputWithContext(ctx context.Context) WanoptRemoteStorageMapOutput {
	return o
}

func (o WanoptRemoteStorageMapOutput) MapIndex(k pulumi.StringInput) WanoptRemoteStorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WanoptRemoteStorage {
		return vs[0].(map[string]WanoptRemoteStorage)[vs[1].(string)]
	}).(WanoptRemoteStorageOutput)
}

func init() {
	pulumi.RegisterOutputType(WanoptRemoteStorageOutput{})
	pulumi.RegisterOutputType(WanoptRemoteStoragePtrOutput{})
	pulumi.RegisterOutputType(WanoptRemoteStorageArrayOutput{})
	pulumi.RegisterOutputType(WanoptRemoteStorageMapOutput{})
}
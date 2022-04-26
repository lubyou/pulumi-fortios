// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure logical storage.
//
// ## Import
//
// System Storage can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemStorage:SystemStorage labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemStorage:SystemStorage labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemStorage struct {
	pulumi.CustomResourceState

	// Partition device.
	Device pulumi.StringOutput `pulumi:"device"`
	// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
	MediaStatus pulumi.StringOutput `pulumi:"mediaStatus"`
	// Storage name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set storage order.
	Order pulumi.IntOutput `pulumi:"order"`
	// Label of underlying partition.
	Partition pulumi.StringOutput `pulumi:"partition"`
	// Partition size.
	Size pulumi.IntOutput `pulumi:"size"`
	// Enable/disable storage. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Use hard disk for logging or WAN Optimization (default = log).
	Usage pulumi.StringOutput `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
	WanoptMode pulumi.StringOutput `pulumi:"wanoptMode"`
}

// NewSystemStorage registers a new resource with the given unique name, arguments, and options.
func NewSystemStorage(ctx *pulumi.Context,
	name string, args *SystemStorageArgs, opts ...pulumi.ResourceOption) (*SystemStorage, error) {
	if args == nil {
		args = &SystemStorageArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemStorage
	err := ctx.RegisterResource("fortios:index/systemStorage:SystemStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemStorage gets an existing SystemStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemStorageState, opts ...pulumi.ResourceOption) (*SystemStorage, error) {
	var resource SystemStorage
	err := ctx.ReadResource("fortios:index/systemStorage:SystemStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemStorage resources.
type systemStorageState struct {
	// Partition device.
	Device *string `pulumi:"device"`
	// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
	MediaStatus *string `pulumi:"mediaStatus"`
	// Storage name.
	Name *string `pulumi:"name"`
	// Set storage order.
	Order *int `pulumi:"order"`
	// Label of underlying partition.
	Partition *string `pulumi:"partition"`
	// Partition size.
	Size *int `pulumi:"size"`
	// Enable/disable storage. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Use hard disk for logging or WAN Optimization (default = log).
	Usage *string `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
	WanoptMode *string `pulumi:"wanoptMode"`
}

type SystemStorageState struct {
	// Partition device.
	Device pulumi.StringPtrInput
	// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
	MediaStatus pulumi.StringPtrInput
	// Storage name.
	Name pulumi.StringPtrInput
	// Set storage order.
	Order pulumi.IntPtrInput
	// Label of underlying partition.
	Partition pulumi.StringPtrInput
	// Partition size.
	Size pulumi.IntPtrInput
	// Enable/disable storage. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Use hard disk for logging or WAN Optimization (default = log).
	Usage pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
	WanoptMode pulumi.StringPtrInput
}

func (SystemStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemStorageState)(nil)).Elem()
}

type systemStorageArgs struct {
	// Partition device.
	Device *string `pulumi:"device"`
	// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
	MediaStatus *string `pulumi:"mediaStatus"`
	// Storage name.
	Name *string `pulumi:"name"`
	// Set storage order.
	Order *int `pulumi:"order"`
	// Label of underlying partition.
	Partition *string `pulumi:"partition"`
	// Partition size.
	Size *int `pulumi:"size"`
	// Enable/disable storage. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Use hard disk for logging or WAN Optimization (default = log).
	Usage *string `pulumi:"usage"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
	WanoptMode *string `pulumi:"wanoptMode"`
}

// The set of arguments for constructing a SystemStorage resource.
type SystemStorageArgs struct {
	// Partition device.
	Device pulumi.StringPtrInput
	// The physical status of current media. Valid values: `enable`, `disable`, `fail`.
	MediaStatus pulumi.StringPtrInput
	// Storage name.
	Name pulumi.StringPtrInput
	// Set storage order.
	Order pulumi.IntPtrInput
	// Label of underlying partition.
	Partition pulumi.StringPtrInput
	// Partition size.
	Size pulumi.IntPtrInput
	// Enable/disable storage. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Use hard disk for logging or WAN Optimization (default = log).
	Usage pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
	WanoptMode pulumi.StringPtrInput
}

func (SystemStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemStorageArgs)(nil)).Elem()
}

type SystemStorageInput interface {
	pulumi.Input

	ToSystemStorageOutput() SystemStorageOutput
	ToSystemStorageOutputWithContext(ctx context.Context) SystemStorageOutput
}

func (*SystemStorage) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemStorage)(nil)).Elem()
}

func (i *SystemStorage) ToSystemStorageOutput() SystemStorageOutput {
	return i.ToSystemStorageOutputWithContext(context.Background())
}

func (i *SystemStorage) ToSystemStorageOutputWithContext(ctx context.Context) SystemStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStorageOutput)
}

// SystemStorageArrayInput is an input type that accepts SystemStorageArray and SystemStorageArrayOutput values.
// You can construct a concrete instance of `SystemStorageArrayInput` via:
//
//          SystemStorageArray{ SystemStorageArgs{...} }
type SystemStorageArrayInput interface {
	pulumi.Input

	ToSystemStorageArrayOutput() SystemStorageArrayOutput
	ToSystemStorageArrayOutputWithContext(context.Context) SystemStorageArrayOutput
}

type SystemStorageArray []SystemStorageInput

func (SystemStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemStorage)(nil)).Elem()
}

func (i SystemStorageArray) ToSystemStorageArrayOutput() SystemStorageArrayOutput {
	return i.ToSystemStorageArrayOutputWithContext(context.Background())
}

func (i SystemStorageArray) ToSystemStorageArrayOutputWithContext(ctx context.Context) SystemStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStorageArrayOutput)
}

// SystemStorageMapInput is an input type that accepts SystemStorageMap and SystemStorageMapOutput values.
// You can construct a concrete instance of `SystemStorageMapInput` via:
//
//          SystemStorageMap{ "key": SystemStorageArgs{...} }
type SystemStorageMapInput interface {
	pulumi.Input

	ToSystemStorageMapOutput() SystemStorageMapOutput
	ToSystemStorageMapOutputWithContext(context.Context) SystemStorageMapOutput
}

type SystemStorageMap map[string]SystemStorageInput

func (SystemStorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemStorage)(nil)).Elem()
}

func (i SystemStorageMap) ToSystemStorageMapOutput() SystemStorageMapOutput {
	return i.ToSystemStorageMapOutputWithContext(context.Background())
}

func (i SystemStorageMap) ToSystemStorageMapOutputWithContext(ctx context.Context) SystemStorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStorageMapOutput)
}

type SystemStorageOutput struct{ *pulumi.OutputState }

func (SystemStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemStorage)(nil)).Elem()
}

func (o SystemStorageOutput) ToSystemStorageOutput() SystemStorageOutput {
	return o
}

func (o SystemStorageOutput) ToSystemStorageOutputWithContext(ctx context.Context) SystemStorageOutput {
	return o
}

type SystemStorageArrayOutput struct{ *pulumi.OutputState }

func (SystemStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemStorage)(nil)).Elem()
}

func (o SystemStorageArrayOutput) ToSystemStorageArrayOutput() SystemStorageArrayOutput {
	return o
}

func (o SystemStorageArrayOutput) ToSystemStorageArrayOutputWithContext(ctx context.Context) SystemStorageArrayOutput {
	return o
}

func (o SystemStorageArrayOutput) Index(i pulumi.IntInput) SystemStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemStorage {
		return vs[0].([]*SystemStorage)[vs[1].(int)]
	}).(SystemStorageOutput)
}

type SystemStorageMapOutput struct{ *pulumi.OutputState }

func (SystemStorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemStorage)(nil)).Elem()
}

func (o SystemStorageMapOutput) ToSystemStorageMapOutput() SystemStorageMapOutput {
	return o
}

func (o SystemStorageMapOutput) ToSystemStorageMapOutputWithContext(ctx context.Context) SystemStorageMapOutput {
	return o
}

func (o SystemStorageMapOutput) MapIndex(k pulumi.StringInput) SystemStorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemStorage {
		return vs[0].(map[string]*SystemStorage)[vs[1].(string)]
	}).(SystemStorageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStorageInput)(nil)).Elem(), &SystemStorage{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStorageArrayInput)(nil)).Elem(), SystemStorageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStorageMapInput)(nil)).Elem(), SystemStorageMap{})
	pulumi.RegisterOutputType(SystemStorageOutput{})
	pulumi.RegisterOutputType(SystemStorageArrayOutput{})
	pulumi.RegisterOutputType(SystemStorageMapOutput{})
}

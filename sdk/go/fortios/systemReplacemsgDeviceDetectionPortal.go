// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replacement messages. Applies to FortiOS Version `<= 6.4.1`.
//
// ## Import
//
// SystemReplacemsg DeviceDetectionPortal can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgDeviceDetectionPortal:SystemReplacemsgDeviceDetectionPortal labelname {{msg_type}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgDeviceDetectionPortal:SystemReplacemsgDeviceDetectionPortal labelname {{msg_type}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemReplacemsgDeviceDetectionPortal struct {
	pulumi.CustomResourceState

	// Message string.
	Buffer pulumi.StringPtrOutput `pulumi:"buffer"`
	// Format flag.
	Format pulumi.StringOutput `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringOutput `pulumi:"header"`
	// Message type.
	MsgType pulumi.StringOutput `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgDeviceDetectionPortal registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgDeviceDetectionPortal(ctx *pulumi.Context,
	name string, args *SystemReplacemsgDeviceDetectionPortalArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgDeviceDetectionPortal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgDeviceDetectionPortal
	err := ctx.RegisterResource("fortios:index/systemReplacemsgDeviceDetectionPortal:SystemReplacemsgDeviceDetectionPortal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgDeviceDetectionPortal gets an existing SystemReplacemsgDeviceDetectionPortal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgDeviceDetectionPortal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgDeviceDetectionPortalState, opts ...pulumi.ResourceOption) (*SystemReplacemsgDeviceDetectionPortal, error) {
	var resource SystemReplacemsgDeviceDetectionPortal
	err := ctx.ReadResource("fortios:index/systemReplacemsgDeviceDetectionPortal:SystemReplacemsgDeviceDetectionPortal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgDeviceDetectionPortal resources.
type systemReplacemsgDeviceDetectionPortalState struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgDeviceDetectionPortalState struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgDeviceDetectionPortalState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgDeviceDetectionPortalState)(nil)).Elem()
}

type systemReplacemsgDeviceDetectionPortalArgs struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgDeviceDetectionPortal resource.
type SystemReplacemsgDeviceDetectionPortalArgs struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgDeviceDetectionPortalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgDeviceDetectionPortalArgs)(nil)).Elem()
}

type SystemReplacemsgDeviceDetectionPortalInput interface {
	pulumi.Input

	ToSystemReplacemsgDeviceDetectionPortalOutput() SystemReplacemsgDeviceDetectionPortalOutput
	ToSystemReplacemsgDeviceDetectionPortalOutputWithContext(ctx context.Context) SystemReplacemsgDeviceDetectionPortalOutput
}

func (*SystemReplacemsgDeviceDetectionPortal) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgDeviceDetectionPortal)(nil)).Elem()
}

func (i *SystemReplacemsgDeviceDetectionPortal) ToSystemReplacemsgDeviceDetectionPortalOutput() SystemReplacemsgDeviceDetectionPortalOutput {
	return i.ToSystemReplacemsgDeviceDetectionPortalOutputWithContext(context.Background())
}

func (i *SystemReplacemsgDeviceDetectionPortal) ToSystemReplacemsgDeviceDetectionPortalOutputWithContext(ctx context.Context) SystemReplacemsgDeviceDetectionPortalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgDeviceDetectionPortalOutput)
}

// SystemReplacemsgDeviceDetectionPortalArrayInput is an input type that accepts SystemReplacemsgDeviceDetectionPortalArray and SystemReplacemsgDeviceDetectionPortalArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgDeviceDetectionPortalArrayInput` via:
//
//          SystemReplacemsgDeviceDetectionPortalArray{ SystemReplacemsgDeviceDetectionPortalArgs{...} }
type SystemReplacemsgDeviceDetectionPortalArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgDeviceDetectionPortalArrayOutput() SystemReplacemsgDeviceDetectionPortalArrayOutput
	ToSystemReplacemsgDeviceDetectionPortalArrayOutputWithContext(context.Context) SystemReplacemsgDeviceDetectionPortalArrayOutput
}

type SystemReplacemsgDeviceDetectionPortalArray []SystemReplacemsgDeviceDetectionPortalInput

func (SystemReplacemsgDeviceDetectionPortalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgDeviceDetectionPortal)(nil)).Elem()
}

func (i SystemReplacemsgDeviceDetectionPortalArray) ToSystemReplacemsgDeviceDetectionPortalArrayOutput() SystemReplacemsgDeviceDetectionPortalArrayOutput {
	return i.ToSystemReplacemsgDeviceDetectionPortalArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgDeviceDetectionPortalArray) ToSystemReplacemsgDeviceDetectionPortalArrayOutputWithContext(ctx context.Context) SystemReplacemsgDeviceDetectionPortalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgDeviceDetectionPortalArrayOutput)
}

// SystemReplacemsgDeviceDetectionPortalMapInput is an input type that accepts SystemReplacemsgDeviceDetectionPortalMap and SystemReplacemsgDeviceDetectionPortalMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgDeviceDetectionPortalMapInput` via:
//
//          SystemReplacemsgDeviceDetectionPortalMap{ "key": SystemReplacemsgDeviceDetectionPortalArgs{...} }
type SystemReplacemsgDeviceDetectionPortalMapInput interface {
	pulumi.Input

	ToSystemReplacemsgDeviceDetectionPortalMapOutput() SystemReplacemsgDeviceDetectionPortalMapOutput
	ToSystemReplacemsgDeviceDetectionPortalMapOutputWithContext(context.Context) SystemReplacemsgDeviceDetectionPortalMapOutput
}

type SystemReplacemsgDeviceDetectionPortalMap map[string]SystemReplacemsgDeviceDetectionPortalInput

func (SystemReplacemsgDeviceDetectionPortalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgDeviceDetectionPortal)(nil)).Elem()
}

func (i SystemReplacemsgDeviceDetectionPortalMap) ToSystemReplacemsgDeviceDetectionPortalMapOutput() SystemReplacemsgDeviceDetectionPortalMapOutput {
	return i.ToSystemReplacemsgDeviceDetectionPortalMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgDeviceDetectionPortalMap) ToSystemReplacemsgDeviceDetectionPortalMapOutputWithContext(ctx context.Context) SystemReplacemsgDeviceDetectionPortalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgDeviceDetectionPortalMapOutput)
}

type SystemReplacemsgDeviceDetectionPortalOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgDeviceDetectionPortalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgDeviceDetectionPortal)(nil)).Elem()
}

func (o SystemReplacemsgDeviceDetectionPortalOutput) ToSystemReplacemsgDeviceDetectionPortalOutput() SystemReplacemsgDeviceDetectionPortalOutput {
	return o
}

func (o SystemReplacemsgDeviceDetectionPortalOutput) ToSystemReplacemsgDeviceDetectionPortalOutputWithContext(ctx context.Context) SystemReplacemsgDeviceDetectionPortalOutput {
	return o
}

type SystemReplacemsgDeviceDetectionPortalArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgDeviceDetectionPortalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgDeviceDetectionPortal)(nil)).Elem()
}

func (o SystemReplacemsgDeviceDetectionPortalArrayOutput) ToSystemReplacemsgDeviceDetectionPortalArrayOutput() SystemReplacemsgDeviceDetectionPortalArrayOutput {
	return o
}

func (o SystemReplacemsgDeviceDetectionPortalArrayOutput) ToSystemReplacemsgDeviceDetectionPortalArrayOutputWithContext(ctx context.Context) SystemReplacemsgDeviceDetectionPortalArrayOutput {
	return o
}

func (o SystemReplacemsgDeviceDetectionPortalArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgDeviceDetectionPortalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgDeviceDetectionPortal {
		return vs[0].([]*SystemReplacemsgDeviceDetectionPortal)[vs[1].(int)]
	}).(SystemReplacemsgDeviceDetectionPortalOutput)
}

type SystemReplacemsgDeviceDetectionPortalMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgDeviceDetectionPortalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgDeviceDetectionPortal)(nil)).Elem()
}

func (o SystemReplacemsgDeviceDetectionPortalMapOutput) ToSystemReplacemsgDeviceDetectionPortalMapOutput() SystemReplacemsgDeviceDetectionPortalMapOutput {
	return o
}

func (o SystemReplacemsgDeviceDetectionPortalMapOutput) ToSystemReplacemsgDeviceDetectionPortalMapOutputWithContext(ctx context.Context) SystemReplacemsgDeviceDetectionPortalMapOutput {
	return o
}

func (o SystemReplacemsgDeviceDetectionPortalMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgDeviceDetectionPortalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgDeviceDetectionPortal {
		return vs[0].(map[string]*SystemReplacemsgDeviceDetectionPortal)[vs[1].(string)]
	}).(SystemReplacemsgDeviceDetectionPortalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgDeviceDetectionPortalInput)(nil)).Elem(), &SystemReplacemsgDeviceDetectionPortal{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgDeviceDetectionPortalArrayInput)(nil)).Elem(), SystemReplacemsgDeviceDetectionPortalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgDeviceDetectionPortalMapInput)(nil)).Elem(), SystemReplacemsgDeviceDetectionPortalMap{})
	pulumi.RegisterOutputType(SystemReplacemsgDeviceDetectionPortalOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgDeviceDetectionPortalArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgDeviceDetectionPortalMapOutput{})
}

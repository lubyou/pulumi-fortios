// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replacement messages.
//
// ## Import
//
// SystemReplacemsg Utm can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgUtm:SystemReplacemsgUtm labelname {{msg_type}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgUtm:SystemReplacemsgUtm labelname {{msg_type}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemReplacemsgUtm struct {
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

// NewSystemReplacemsgUtm registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgUtm(ctx *pulumi.Context,
	name string, args *SystemReplacemsgUtmArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgUtm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgUtm
	err := ctx.RegisterResource("fortios:index/systemReplacemsgUtm:SystemReplacemsgUtm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgUtm gets an existing SystemReplacemsgUtm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgUtm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgUtmState, opts ...pulumi.ResourceOption) (*SystemReplacemsgUtm, error) {
	var resource SystemReplacemsgUtm
	err := ctx.ReadResource("fortios:index/systemReplacemsgUtm:SystemReplacemsgUtm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgUtm resources.
type systemReplacemsgUtmState struct {
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

type SystemReplacemsgUtmState struct {
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

func (SystemReplacemsgUtmState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgUtmState)(nil)).Elem()
}

type systemReplacemsgUtmArgs struct {
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

// The set of arguments for constructing a SystemReplacemsgUtm resource.
type SystemReplacemsgUtmArgs struct {
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

func (SystemReplacemsgUtmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgUtmArgs)(nil)).Elem()
}

type SystemReplacemsgUtmInput interface {
	pulumi.Input

	ToSystemReplacemsgUtmOutput() SystemReplacemsgUtmOutput
	ToSystemReplacemsgUtmOutputWithContext(ctx context.Context) SystemReplacemsgUtmOutput
}

func (*SystemReplacemsgUtm) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgUtm)(nil)).Elem()
}

func (i *SystemReplacemsgUtm) ToSystemReplacemsgUtmOutput() SystemReplacemsgUtmOutput {
	return i.ToSystemReplacemsgUtmOutputWithContext(context.Background())
}

func (i *SystemReplacemsgUtm) ToSystemReplacemsgUtmOutputWithContext(ctx context.Context) SystemReplacemsgUtmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgUtmOutput)
}

// SystemReplacemsgUtmArrayInput is an input type that accepts SystemReplacemsgUtmArray and SystemReplacemsgUtmArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgUtmArrayInput` via:
//
//          SystemReplacemsgUtmArray{ SystemReplacemsgUtmArgs{...} }
type SystemReplacemsgUtmArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgUtmArrayOutput() SystemReplacemsgUtmArrayOutput
	ToSystemReplacemsgUtmArrayOutputWithContext(context.Context) SystemReplacemsgUtmArrayOutput
}

type SystemReplacemsgUtmArray []SystemReplacemsgUtmInput

func (SystemReplacemsgUtmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgUtm)(nil)).Elem()
}

func (i SystemReplacemsgUtmArray) ToSystemReplacemsgUtmArrayOutput() SystemReplacemsgUtmArrayOutput {
	return i.ToSystemReplacemsgUtmArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgUtmArray) ToSystemReplacemsgUtmArrayOutputWithContext(ctx context.Context) SystemReplacemsgUtmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgUtmArrayOutput)
}

// SystemReplacemsgUtmMapInput is an input type that accepts SystemReplacemsgUtmMap and SystemReplacemsgUtmMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgUtmMapInput` via:
//
//          SystemReplacemsgUtmMap{ "key": SystemReplacemsgUtmArgs{...} }
type SystemReplacemsgUtmMapInput interface {
	pulumi.Input

	ToSystemReplacemsgUtmMapOutput() SystemReplacemsgUtmMapOutput
	ToSystemReplacemsgUtmMapOutputWithContext(context.Context) SystemReplacemsgUtmMapOutput
}

type SystemReplacemsgUtmMap map[string]SystemReplacemsgUtmInput

func (SystemReplacemsgUtmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgUtm)(nil)).Elem()
}

func (i SystemReplacemsgUtmMap) ToSystemReplacemsgUtmMapOutput() SystemReplacemsgUtmMapOutput {
	return i.ToSystemReplacemsgUtmMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgUtmMap) ToSystemReplacemsgUtmMapOutputWithContext(ctx context.Context) SystemReplacemsgUtmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgUtmMapOutput)
}

type SystemReplacemsgUtmOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgUtmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgUtm)(nil)).Elem()
}

func (o SystemReplacemsgUtmOutput) ToSystemReplacemsgUtmOutput() SystemReplacemsgUtmOutput {
	return o
}

func (o SystemReplacemsgUtmOutput) ToSystemReplacemsgUtmOutputWithContext(ctx context.Context) SystemReplacemsgUtmOutput {
	return o
}

type SystemReplacemsgUtmArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgUtmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgUtm)(nil)).Elem()
}

func (o SystemReplacemsgUtmArrayOutput) ToSystemReplacemsgUtmArrayOutput() SystemReplacemsgUtmArrayOutput {
	return o
}

func (o SystemReplacemsgUtmArrayOutput) ToSystemReplacemsgUtmArrayOutputWithContext(ctx context.Context) SystemReplacemsgUtmArrayOutput {
	return o
}

func (o SystemReplacemsgUtmArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgUtmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgUtm {
		return vs[0].([]*SystemReplacemsgUtm)[vs[1].(int)]
	}).(SystemReplacemsgUtmOutput)
}

type SystemReplacemsgUtmMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgUtmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgUtm)(nil)).Elem()
}

func (o SystemReplacemsgUtmMapOutput) ToSystemReplacemsgUtmMapOutput() SystemReplacemsgUtmMapOutput {
	return o
}

func (o SystemReplacemsgUtmMapOutput) ToSystemReplacemsgUtmMapOutputWithContext(ctx context.Context) SystemReplacemsgUtmMapOutput {
	return o
}

func (o SystemReplacemsgUtmMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgUtmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgUtm {
		return vs[0].(map[string]*SystemReplacemsgUtm)[vs[1].(string)]
	}).(SystemReplacemsgUtmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgUtmInput)(nil)).Elem(), &SystemReplacemsgUtm{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgUtmArrayInput)(nil)).Elem(), SystemReplacemsgUtmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgUtmMapInput)(nil)).Elem(), SystemReplacemsgUtmMap{})
	pulumi.RegisterOutputType(SystemReplacemsgUtmOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgUtmArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgUtmMapOutput{})
}

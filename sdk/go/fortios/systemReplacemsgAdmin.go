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
// SystemReplacemsg Admin can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgAdmin:SystemReplacemsgAdmin labelname {{msg_type}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgAdmin:SystemReplacemsgAdmin labelname {{msg_type}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemReplacemsgAdmin struct {
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

// NewSystemReplacemsgAdmin registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgAdmin(ctx *pulumi.Context,
	name string, args *SystemReplacemsgAdminArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgAdmin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgAdmin
	err := ctx.RegisterResource("fortios:index/systemReplacemsgAdmin:SystemReplacemsgAdmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgAdmin gets an existing SystemReplacemsgAdmin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgAdmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgAdminState, opts ...pulumi.ResourceOption) (*SystemReplacemsgAdmin, error) {
	var resource SystemReplacemsgAdmin
	err := ctx.ReadResource("fortios:index/systemReplacemsgAdmin:SystemReplacemsgAdmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgAdmin resources.
type systemReplacemsgAdminState struct {
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

type SystemReplacemsgAdminState struct {
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

func (SystemReplacemsgAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgAdminState)(nil)).Elem()
}

type systemReplacemsgAdminArgs struct {
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

// The set of arguments for constructing a SystemReplacemsgAdmin resource.
type SystemReplacemsgAdminArgs struct {
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

func (SystemReplacemsgAdminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgAdminArgs)(nil)).Elem()
}

type SystemReplacemsgAdminInput interface {
	pulumi.Input

	ToSystemReplacemsgAdminOutput() SystemReplacemsgAdminOutput
	ToSystemReplacemsgAdminOutputWithContext(ctx context.Context) SystemReplacemsgAdminOutput
}

func (*SystemReplacemsgAdmin) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgAdmin)(nil)).Elem()
}

func (i *SystemReplacemsgAdmin) ToSystemReplacemsgAdminOutput() SystemReplacemsgAdminOutput {
	return i.ToSystemReplacemsgAdminOutputWithContext(context.Background())
}

func (i *SystemReplacemsgAdmin) ToSystemReplacemsgAdminOutputWithContext(ctx context.Context) SystemReplacemsgAdminOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAdminOutput)
}

// SystemReplacemsgAdminArrayInput is an input type that accepts SystemReplacemsgAdminArray and SystemReplacemsgAdminArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgAdminArrayInput` via:
//
//          SystemReplacemsgAdminArray{ SystemReplacemsgAdminArgs{...} }
type SystemReplacemsgAdminArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgAdminArrayOutput() SystemReplacemsgAdminArrayOutput
	ToSystemReplacemsgAdminArrayOutputWithContext(context.Context) SystemReplacemsgAdminArrayOutput
}

type SystemReplacemsgAdminArray []SystemReplacemsgAdminInput

func (SystemReplacemsgAdminArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgAdmin)(nil)).Elem()
}

func (i SystemReplacemsgAdminArray) ToSystemReplacemsgAdminArrayOutput() SystemReplacemsgAdminArrayOutput {
	return i.ToSystemReplacemsgAdminArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgAdminArray) ToSystemReplacemsgAdminArrayOutputWithContext(ctx context.Context) SystemReplacemsgAdminArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAdminArrayOutput)
}

// SystemReplacemsgAdminMapInput is an input type that accepts SystemReplacemsgAdminMap and SystemReplacemsgAdminMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgAdminMapInput` via:
//
//          SystemReplacemsgAdminMap{ "key": SystemReplacemsgAdminArgs{...} }
type SystemReplacemsgAdminMapInput interface {
	pulumi.Input

	ToSystemReplacemsgAdminMapOutput() SystemReplacemsgAdminMapOutput
	ToSystemReplacemsgAdminMapOutputWithContext(context.Context) SystemReplacemsgAdminMapOutput
}

type SystemReplacemsgAdminMap map[string]SystemReplacemsgAdminInput

func (SystemReplacemsgAdminMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgAdmin)(nil)).Elem()
}

func (i SystemReplacemsgAdminMap) ToSystemReplacemsgAdminMapOutput() SystemReplacemsgAdminMapOutput {
	return i.ToSystemReplacemsgAdminMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgAdminMap) ToSystemReplacemsgAdminMapOutputWithContext(ctx context.Context) SystemReplacemsgAdminMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAdminMapOutput)
}

type SystemReplacemsgAdminOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAdminOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgAdmin)(nil)).Elem()
}

func (o SystemReplacemsgAdminOutput) ToSystemReplacemsgAdminOutput() SystemReplacemsgAdminOutput {
	return o
}

func (o SystemReplacemsgAdminOutput) ToSystemReplacemsgAdminOutputWithContext(ctx context.Context) SystemReplacemsgAdminOutput {
	return o
}

type SystemReplacemsgAdminArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAdminArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgAdmin)(nil)).Elem()
}

func (o SystemReplacemsgAdminArrayOutput) ToSystemReplacemsgAdminArrayOutput() SystemReplacemsgAdminArrayOutput {
	return o
}

func (o SystemReplacemsgAdminArrayOutput) ToSystemReplacemsgAdminArrayOutputWithContext(ctx context.Context) SystemReplacemsgAdminArrayOutput {
	return o
}

func (o SystemReplacemsgAdminArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgAdminOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgAdmin {
		return vs[0].([]*SystemReplacemsgAdmin)[vs[1].(int)]
	}).(SystemReplacemsgAdminOutput)
}

type SystemReplacemsgAdminMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAdminMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgAdmin)(nil)).Elem()
}

func (o SystemReplacemsgAdminMapOutput) ToSystemReplacemsgAdminMapOutput() SystemReplacemsgAdminMapOutput {
	return o
}

func (o SystemReplacemsgAdminMapOutput) ToSystemReplacemsgAdminMapOutputWithContext(ctx context.Context) SystemReplacemsgAdminMapOutput {
	return o
}

func (o SystemReplacemsgAdminMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgAdminOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgAdmin {
		return vs[0].(map[string]*SystemReplacemsgAdmin)[vs[1].(string)]
	}).(SystemReplacemsgAdminOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAdminInput)(nil)).Elem(), &SystemReplacemsgAdmin{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAdminArrayInput)(nil)).Elem(), SystemReplacemsgAdminArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAdminMapInput)(nil)).Elem(), SystemReplacemsgAdminMap{})
	pulumi.RegisterOutputType(SystemReplacemsgAdminOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgAdminArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgAdminMapOutput{})
}

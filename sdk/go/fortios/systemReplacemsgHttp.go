// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
// SystemReplacemsg Http can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgHttp:SystemReplacemsgHttp labelname {{msg_type}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemReplacemsgHttp struct {
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

// NewSystemReplacemsgHttp registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgHttp(ctx *pulumi.Context,
	name string, args *SystemReplacemsgHttpArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgHttp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	var resource SystemReplacemsgHttp
	err := ctx.RegisterResource("fortios:index/systemReplacemsgHttp:SystemReplacemsgHttp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgHttp gets an existing SystemReplacemsgHttp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgHttp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgHttpState, opts ...pulumi.ResourceOption) (*SystemReplacemsgHttp, error) {
	var resource SystemReplacemsgHttp
	err := ctx.ReadResource("fortios:index/systemReplacemsgHttp:SystemReplacemsgHttp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgHttp resources.
type systemReplacemsgHttpState struct {
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

type SystemReplacemsgHttpState struct {
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

func (SystemReplacemsgHttpState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgHttpState)(nil)).Elem()
}

type systemReplacemsgHttpArgs struct {
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

// The set of arguments for constructing a SystemReplacemsgHttp resource.
type SystemReplacemsgHttpArgs struct {
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

func (SystemReplacemsgHttpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgHttpArgs)(nil)).Elem()
}

type SystemReplacemsgHttpInput interface {
	pulumi.Input

	ToSystemReplacemsgHttpOutput() SystemReplacemsgHttpOutput
	ToSystemReplacemsgHttpOutputWithContext(ctx context.Context) SystemReplacemsgHttpOutput
}

func (*SystemReplacemsgHttp) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemReplacemsgHttp)(nil))
}

func (i *SystemReplacemsgHttp) ToSystemReplacemsgHttpOutput() SystemReplacemsgHttpOutput {
	return i.ToSystemReplacemsgHttpOutputWithContext(context.Background())
}

func (i *SystemReplacemsgHttp) ToSystemReplacemsgHttpOutputWithContext(ctx context.Context) SystemReplacemsgHttpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgHttpOutput)
}

func (i *SystemReplacemsgHttp) ToSystemReplacemsgHttpPtrOutput() SystemReplacemsgHttpPtrOutput {
	return i.ToSystemReplacemsgHttpPtrOutputWithContext(context.Background())
}

func (i *SystemReplacemsgHttp) ToSystemReplacemsgHttpPtrOutputWithContext(ctx context.Context) SystemReplacemsgHttpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgHttpPtrOutput)
}

type SystemReplacemsgHttpPtrInput interface {
	pulumi.Input

	ToSystemReplacemsgHttpPtrOutput() SystemReplacemsgHttpPtrOutput
	ToSystemReplacemsgHttpPtrOutputWithContext(ctx context.Context) SystemReplacemsgHttpPtrOutput
}

type systemReplacemsgHttpPtrType SystemReplacemsgHttpArgs

func (*systemReplacemsgHttpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgHttp)(nil))
}

func (i *systemReplacemsgHttpPtrType) ToSystemReplacemsgHttpPtrOutput() SystemReplacemsgHttpPtrOutput {
	return i.ToSystemReplacemsgHttpPtrOutputWithContext(context.Background())
}

func (i *systemReplacemsgHttpPtrType) ToSystemReplacemsgHttpPtrOutputWithContext(ctx context.Context) SystemReplacemsgHttpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgHttpPtrOutput)
}

// SystemReplacemsgHttpArrayInput is an input type that accepts SystemReplacemsgHttpArray and SystemReplacemsgHttpArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgHttpArrayInput` via:
//
//          SystemReplacemsgHttpArray{ SystemReplacemsgHttpArgs{...} }
type SystemReplacemsgHttpArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgHttpArrayOutput() SystemReplacemsgHttpArrayOutput
	ToSystemReplacemsgHttpArrayOutputWithContext(context.Context) SystemReplacemsgHttpArrayOutput
}

type SystemReplacemsgHttpArray []SystemReplacemsgHttpInput

func (SystemReplacemsgHttpArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemReplacemsgHttp)(nil))
}

func (i SystemReplacemsgHttpArray) ToSystemReplacemsgHttpArrayOutput() SystemReplacemsgHttpArrayOutput {
	return i.ToSystemReplacemsgHttpArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgHttpArray) ToSystemReplacemsgHttpArrayOutputWithContext(ctx context.Context) SystemReplacemsgHttpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgHttpArrayOutput)
}

// SystemReplacemsgHttpMapInput is an input type that accepts SystemReplacemsgHttpMap and SystemReplacemsgHttpMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgHttpMapInput` via:
//
//          SystemReplacemsgHttpMap{ "key": SystemReplacemsgHttpArgs{...} }
type SystemReplacemsgHttpMapInput interface {
	pulumi.Input

	ToSystemReplacemsgHttpMapOutput() SystemReplacemsgHttpMapOutput
	ToSystemReplacemsgHttpMapOutputWithContext(context.Context) SystemReplacemsgHttpMapOutput
}

type SystemReplacemsgHttpMap map[string]SystemReplacemsgHttpInput

func (SystemReplacemsgHttpMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemReplacemsgHttp)(nil))
}

func (i SystemReplacemsgHttpMap) ToSystemReplacemsgHttpMapOutput() SystemReplacemsgHttpMapOutput {
	return i.ToSystemReplacemsgHttpMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgHttpMap) ToSystemReplacemsgHttpMapOutputWithContext(ctx context.Context) SystemReplacemsgHttpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgHttpMapOutput)
}

type SystemReplacemsgHttpOutput struct {
	*pulumi.OutputState
}

func (SystemReplacemsgHttpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemReplacemsgHttp)(nil))
}

func (o SystemReplacemsgHttpOutput) ToSystemReplacemsgHttpOutput() SystemReplacemsgHttpOutput {
	return o
}

func (o SystemReplacemsgHttpOutput) ToSystemReplacemsgHttpOutputWithContext(ctx context.Context) SystemReplacemsgHttpOutput {
	return o
}

func (o SystemReplacemsgHttpOutput) ToSystemReplacemsgHttpPtrOutput() SystemReplacemsgHttpPtrOutput {
	return o.ToSystemReplacemsgHttpPtrOutputWithContext(context.Background())
}

func (o SystemReplacemsgHttpOutput) ToSystemReplacemsgHttpPtrOutputWithContext(ctx context.Context) SystemReplacemsgHttpPtrOutput {
	return o.ApplyT(func(v SystemReplacemsgHttp) *SystemReplacemsgHttp {
		return &v
	}).(SystemReplacemsgHttpPtrOutput)
}

type SystemReplacemsgHttpPtrOutput struct {
	*pulumi.OutputState
}

func (SystemReplacemsgHttpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgHttp)(nil))
}

func (o SystemReplacemsgHttpPtrOutput) ToSystemReplacemsgHttpPtrOutput() SystemReplacemsgHttpPtrOutput {
	return o
}

func (o SystemReplacemsgHttpPtrOutput) ToSystemReplacemsgHttpPtrOutputWithContext(ctx context.Context) SystemReplacemsgHttpPtrOutput {
	return o
}

type SystemReplacemsgHttpArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgHttpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemReplacemsgHttp)(nil))
}

func (o SystemReplacemsgHttpArrayOutput) ToSystemReplacemsgHttpArrayOutput() SystemReplacemsgHttpArrayOutput {
	return o
}

func (o SystemReplacemsgHttpArrayOutput) ToSystemReplacemsgHttpArrayOutputWithContext(ctx context.Context) SystemReplacemsgHttpArrayOutput {
	return o
}

func (o SystemReplacemsgHttpArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgHttpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemReplacemsgHttp {
		return vs[0].([]SystemReplacemsgHttp)[vs[1].(int)]
	}).(SystemReplacemsgHttpOutput)
}

type SystemReplacemsgHttpMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgHttpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemReplacemsgHttp)(nil))
}

func (o SystemReplacemsgHttpMapOutput) ToSystemReplacemsgHttpMapOutput() SystemReplacemsgHttpMapOutput {
	return o
}

func (o SystemReplacemsgHttpMapOutput) ToSystemReplacemsgHttpMapOutputWithContext(ctx context.Context) SystemReplacemsgHttpMapOutput {
	return o
}

func (o SystemReplacemsgHttpMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgHttpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemReplacemsgHttp {
		return vs[0].(map[string]SystemReplacemsgHttp)[vs[1].(string)]
	}).(SystemReplacemsgHttpOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemReplacemsgHttpOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgHttpPtrOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgHttpArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgHttpMapOutput{})
}

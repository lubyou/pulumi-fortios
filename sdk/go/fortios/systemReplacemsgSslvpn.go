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
// SystemReplacemsg Sslvpn can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgSslvpn:SystemReplacemsgSslvpn labelname {{msg_type}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgSslvpn:SystemReplacemsgSslvpn labelname {{msg_type}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemReplacemsgSslvpn struct {
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

// NewSystemReplacemsgSslvpn registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgSslvpn(ctx *pulumi.Context,
	name string, args *SystemReplacemsgSslvpnArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgSslvpn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgSslvpn
	err := ctx.RegisterResource("fortios:index/systemReplacemsgSslvpn:SystemReplacemsgSslvpn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgSslvpn gets an existing SystemReplacemsgSslvpn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgSslvpn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgSslvpnState, opts ...pulumi.ResourceOption) (*SystemReplacemsgSslvpn, error) {
	var resource SystemReplacemsgSslvpn
	err := ctx.ReadResource("fortios:index/systemReplacemsgSslvpn:SystemReplacemsgSslvpn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgSslvpn resources.
type systemReplacemsgSslvpnState struct {
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

type SystemReplacemsgSslvpnState struct {
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

func (SystemReplacemsgSslvpnState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgSslvpnState)(nil)).Elem()
}

type systemReplacemsgSslvpnArgs struct {
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

// The set of arguments for constructing a SystemReplacemsgSslvpn resource.
type SystemReplacemsgSslvpnArgs struct {
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

func (SystemReplacemsgSslvpnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgSslvpnArgs)(nil)).Elem()
}

type SystemReplacemsgSslvpnInput interface {
	pulumi.Input

	ToSystemReplacemsgSslvpnOutput() SystemReplacemsgSslvpnOutput
	ToSystemReplacemsgSslvpnOutputWithContext(ctx context.Context) SystemReplacemsgSslvpnOutput
}

func (*SystemReplacemsgSslvpn) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgSslvpn)(nil)).Elem()
}

func (i *SystemReplacemsgSslvpn) ToSystemReplacemsgSslvpnOutput() SystemReplacemsgSslvpnOutput {
	return i.ToSystemReplacemsgSslvpnOutputWithContext(context.Background())
}

func (i *SystemReplacemsgSslvpn) ToSystemReplacemsgSslvpnOutputWithContext(ctx context.Context) SystemReplacemsgSslvpnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgSslvpnOutput)
}

// SystemReplacemsgSslvpnArrayInput is an input type that accepts SystemReplacemsgSslvpnArray and SystemReplacemsgSslvpnArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgSslvpnArrayInput` via:
//
//          SystemReplacemsgSslvpnArray{ SystemReplacemsgSslvpnArgs{...} }
type SystemReplacemsgSslvpnArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgSslvpnArrayOutput() SystemReplacemsgSslvpnArrayOutput
	ToSystemReplacemsgSslvpnArrayOutputWithContext(context.Context) SystemReplacemsgSslvpnArrayOutput
}

type SystemReplacemsgSslvpnArray []SystemReplacemsgSslvpnInput

func (SystemReplacemsgSslvpnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgSslvpn)(nil)).Elem()
}

func (i SystemReplacemsgSslvpnArray) ToSystemReplacemsgSslvpnArrayOutput() SystemReplacemsgSslvpnArrayOutput {
	return i.ToSystemReplacemsgSslvpnArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgSslvpnArray) ToSystemReplacemsgSslvpnArrayOutputWithContext(ctx context.Context) SystemReplacemsgSslvpnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgSslvpnArrayOutput)
}

// SystemReplacemsgSslvpnMapInput is an input type that accepts SystemReplacemsgSslvpnMap and SystemReplacemsgSslvpnMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgSslvpnMapInput` via:
//
//          SystemReplacemsgSslvpnMap{ "key": SystemReplacemsgSslvpnArgs{...} }
type SystemReplacemsgSslvpnMapInput interface {
	pulumi.Input

	ToSystemReplacemsgSslvpnMapOutput() SystemReplacemsgSslvpnMapOutput
	ToSystemReplacemsgSslvpnMapOutputWithContext(context.Context) SystemReplacemsgSslvpnMapOutput
}

type SystemReplacemsgSslvpnMap map[string]SystemReplacemsgSslvpnInput

func (SystemReplacemsgSslvpnMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgSslvpn)(nil)).Elem()
}

func (i SystemReplacemsgSslvpnMap) ToSystemReplacemsgSslvpnMapOutput() SystemReplacemsgSslvpnMapOutput {
	return i.ToSystemReplacemsgSslvpnMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgSslvpnMap) ToSystemReplacemsgSslvpnMapOutputWithContext(ctx context.Context) SystemReplacemsgSslvpnMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgSslvpnMapOutput)
}

type SystemReplacemsgSslvpnOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgSslvpnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgSslvpn)(nil)).Elem()
}

func (o SystemReplacemsgSslvpnOutput) ToSystemReplacemsgSslvpnOutput() SystemReplacemsgSslvpnOutput {
	return o
}

func (o SystemReplacemsgSslvpnOutput) ToSystemReplacemsgSslvpnOutputWithContext(ctx context.Context) SystemReplacemsgSslvpnOutput {
	return o
}

type SystemReplacemsgSslvpnArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgSslvpnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgSslvpn)(nil)).Elem()
}

func (o SystemReplacemsgSslvpnArrayOutput) ToSystemReplacemsgSslvpnArrayOutput() SystemReplacemsgSslvpnArrayOutput {
	return o
}

func (o SystemReplacemsgSslvpnArrayOutput) ToSystemReplacemsgSslvpnArrayOutputWithContext(ctx context.Context) SystemReplacemsgSslvpnArrayOutput {
	return o
}

func (o SystemReplacemsgSslvpnArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgSslvpnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgSslvpn {
		return vs[0].([]*SystemReplacemsgSslvpn)[vs[1].(int)]
	}).(SystemReplacemsgSslvpnOutput)
}

type SystemReplacemsgSslvpnMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgSslvpnMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgSslvpn)(nil)).Elem()
}

func (o SystemReplacemsgSslvpnMapOutput) ToSystemReplacemsgSslvpnMapOutput() SystemReplacemsgSslvpnMapOutput {
	return o
}

func (o SystemReplacemsgSslvpnMapOutput) ToSystemReplacemsgSslvpnMapOutputWithContext(ctx context.Context) SystemReplacemsgSslvpnMapOutput {
	return o
}

func (o SystemReplacemsgSslvpnMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgSslvpnOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgSslvpn {
		return vs[0].(map[string]*SystemReplacemsgSslvpn)[vs[1].(string)]
	}).(SystemReplacemsgSslvpnOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgSslvpnInput)(nil)).Elem(), &SystemReplacemsgSslvpn{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgSslvpnArrayInput)(nil)).Elem(), SystemReplacemsgSslvpnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgSslvpnMapInput)(nil)).Elem(), SystemReplacemsgSslvpnMap{})
	pulumi.RegisterOutputType(SystemReplacemsgSslvpnOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgSslvpnArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgSslvpnMapOutput{})
}

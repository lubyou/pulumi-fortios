// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemReplacemsgAuth struct {
	pulumi.CustomResourceState

	Buffer    pulumi.StringPtrOutput `pulumi:"buffer"`
	Format    pulumi.StringOutput    `pulumi:"format"`
	Header    pulumi.StringOutput    `pulumi:"header"`
	MsgType   pulumi.StringOutput    `pulumi:"msgType"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgAuth registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgAuth(ctx *pulumi.Context,
	name string, args *SystemReplacemsgAuthArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgAuth, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgAuth
	err := ctx.RegisterResource("fortios:index/systemReplacemsgAuth:SystemReplacemsgAuth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgAuth gets an existing SystemReplacemsgAuth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgAuth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgAuthState, opts ...pulumi.ResourceOption) (*SystemReplacemsgAuth, error) {
	var resource SystemReplacemsgAuth
	err := ctx.ReadResource("fortios:index/systemReplacemsgAuth:SystemReplacemsgAuth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgAuth resources.
type systemReplacemsgAuthState struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgAuthState struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgAuthState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgAuthState)(nil)).Elem()
}

type systemReplacemsgAuthArgs struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   string  `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgAuth resource.
type SystemReplacemsgAuthArgs struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgAuthArgs)(nil)).Elem()
}

type SystemReplacemsgAuthInput interface {
	pulumi.Input

	ToSystemReplacemsgAuthOutput() SystemReplacemsgAuthOutput
	ToSystemReplacemsgAuthOutputWithContext(ctx context.Context) SystemReplacemsgAuthOutput
}

func (*SystemReplacemsgAuth) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgAuth)(nil)).Elem()
}

func (i *SystemReplacemsgAuth) ToSystemReplacemsgAuthOutput() SystemReplacemsgAuthOutput {
	return i.ToSystemReplacemsgAuthOutputWithContext(context.Background())
}

func (i *SystemReplacemsgAuth) ToSystemReplacemsgAuthOutputWithContext(ctx context.Context) SystemReplacemsgAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAuthOutput)
}

// SystemReplacemsgAuthArrayInput is an input type that accepts SystemReplacemsgAuthArray and SystemReplacemsgAuthArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgAuthArrayInput` via:
//
//	SystemReplacemsgAuthArray{ SystemReplacemsgAuthArgs{...} }
type SystemReplacemsgAuthArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgAuthArrayOutput() SystemReplacemsgAuthArrayOutput
	ToSystemReplacemsgAuthArrayOutputWithContext(context.Context) SystemReplacemsgAuthArrayOutput
}

type SystemReplacemsgAuthArray []SystemReplacemsgAuthInput

func (SystemReplacemsgAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgAuth)(nil)).Elem()
}

func (i SystemReplacemsgAuthArray) ToSystemReplacemsgAuthArrayOutput() SystemReplacemsgAuthArrayOutput {
	return i.ToSystemReplacemsgAuthArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgAuthArray) ToSystemReplacemsgAuthArrayOutputWithContext(ctx context.Context) SystemReplacemsgAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAuthArrayOutput)
}

// SystemReplacemsgAuthMapInput is an input type that accepts SystemReplacemsgAuthMap and SystemReplacemsgAuthMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgAuthMapInput` via:
//
//	SystemReplacemsgAuthMap{ "key": SystemReplacemsgAuthArgs{...} }
type SystemReplacemsgAuthMapInput interface {
	pulumi.Input

	ToSystemReplacemsgAuthMapOutput() SystemReplacemsgAuthMapOutput
	ToSystemReplacemsgAuthMapOutputWithContext(context.Context) SystemReplacemsgAuthMapOutput
}

type SystemReplacemsgAuthMap map[string]SystemReplacemsgAuthInput

func (SystemReplacemsgAuthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgAuth)(nil)).Elem()
}

func (i SystemReplacemsgAuthMap) ToSystemReplacemsgAuthMapOutput() SystemReplacemsgAuthMapOutput {
	return i.ToSystemReplacemsgAuthMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgAuthMap) ToSystemReplacemsgAuthMapOutputWithContext(ctx context.Context) SystemReplacemsgAuthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAuthMapOutput)
}

type SystemReplacemsgAuthOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgAuth)(nil)).Elem()
}

func (o SystemReplacemsgAuthOutput) ToSystemReplacemsgAuthOutput() SystemReplacemsgAuthOutput {
	return o
}

func (o SystemReplacemsgAuthOutput) ToSystemReplacemsgAuthOutputWithContext(ctx context.Context) SystemReplacemsgAuthOutput {
	return o
}

func (o SystemReplacemsgAuthOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgAuth) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

func (o SystemReplacemsgAuthOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgAuth) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o SystemReplacemsgAuthOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgAuth) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

func (o SystemReplacemsgAuthOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgAuth) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

func (o SystemReplacemsgAuthOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgAuth) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemReplacemsgAuthArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgAuth)(nil)).Elem()
}

func (o SystemReplacemsgAuthArrayOutput) ToSystemReplacemsgAuthArrayOutput() SystemReplacemsgAuthArrayOutput {
	return o
}

func (o SystemReplacemsgAuthArrayOutput) ToSystemReplacemsgAuthArrayOutputWithContext(ctx context.Context) SystemReplacemsgAuthArrayOutput {
	return o
}

func (o SystemReplacemsgAuthArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgAuthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgAuth {
		return vs[0].([]*SystemReplacemsgAuth)[vs[1].(int)]
	}).(SystemReplacemsgAuthOutput)
}

type SystemReplacemsgAuthMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAuthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgAuth)(nil)).Elem()
}

func (o SystemReplacemsgAuthMapOutput) ToSystemReplacemsgAuthMapOutput() SystemReplacemsgAuthMapOutput {
	return o
}

func (o SystemReplacemsgAuthMapOutput) ToSystemReplacemsgAuthMapOutputWithContext(ctx context.Context) SystemReplacemsgAuthMapOutput {
	return o
}

func (o SystemReplacemsgAuthMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgAuthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgAuth {
		return vs[0].(map[string]*SystemReplacemsgAuth)[vs[1].(string)]
	}).(SystemReplacemsgAuthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAuthInput)(nil)).Elem(), &SystemReplacemsgAuth{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAuthArrayInput)(nil)).Elem(), SystemReplacemsgAuthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAuthMapInput)(nil)).Elem(), SystemReplacemsgAuthMap{})
	pulumi.RegisterOutputType(SystemReplacemsgAuthOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgAuthArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgAuthMapOutput{})
}

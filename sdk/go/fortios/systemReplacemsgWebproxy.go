// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemReplacemsgWebproxy struct {
	pulumi.CustomResourceState

	Buffer    pulumi.StringPtrOutput `pulumi:"buffer"`
	Format    pulumi.StringOutput    `pulumi:"format"`
	Header    pulumi.StringOutput    `pulumi:"header"`
	MsgType   pulumi.StringOutput    `pulumi:"msgType"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgWebproxy registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgWebproxy(ctx *pulumi.Context,
	name string, args *SystemReplacemsgWebproxyArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgWebproxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgWebproxy
	err := ctx.RegisterResource("fortios:index/systemReplacemsgWebproxy:SystemReplacemsgWebproxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgWebproxy gets an existing SystemReplacemsgWebproxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgWebproxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgWebproxyState, opts ...pulumi.ResourceOption) (*SystemReplacemsgWebproxy, error) {
	var resource SystemReplacemsgWebproxy
	err := ctx.ReadResource("fortios:index/systemReplacemsgWebproxy:SystemReplacemsgWebproxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgWebproxy resources.
type systemReplacemsgWebproxyState struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgWebproxyState struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgWebproxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgWebproxyState)(nil)).Elem()
}

type systemReplacemsgWebproxyArgs struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   string  `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgWebproxy resource.
type SystemReplacemsgWebproxyArgs struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgWebproxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgWebproxyArgs)(nil)).Elem()
}

type SystemReplacemsgWebproxyInput interface {
	pulumi.Input

	ToSystemReplacemsgWebproxyOutput() SystemReplacemsgWebproxyOutput
	ToSystemReplacemsgWebproxyOutputWithContext(ctx context.Context) SystemReplacemsgWebproxyOutput
}

func (*SystemReplacemsgWebproxy) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgWebproxy)(nil)).Elem()
}

func (i *SystemReplacemsgWebproxy) ToSystemReplacemsgWebproxyOutput() SystemReplacemsgWebproxyOutput {
	return i.ToSystemReplacemsgWebproxyOutputWithContext(context.Background())
}

func (i *SystemReplacemsgWebproxy) ToSystemReplacemsgWebproxyOutputWithContext(ctx context.Context) SystemReplacemsgWebproxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgWebproxyOutput)
}

// SystemReplacemsgWebproxyArrayInput is an input type that accepts SystemReplacemsgWebproxyArray and SystemReplacemsgWebproxyArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgWebproxyArrayInput` via:
//
//	SystemReplacemsgWebproxyArray{ SystemReplacemsgWebproxyArgs{...} }
type SystemReplacemsgWebproxyArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgWebproxyArrayOutput() SystemReplacemsgWebproxyArrayOutput
	ToSystemReplacemsgWebproxyArrayOutputWithContext(context.Context) SystemReplacemsgWebproxyArrayOutput
}

type SystemReplacemsgWebproxyArray []SystemReplacemsgWebproxyInput

func (SystemReplacemsgWebproxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgWebproxy)(nil)).Elem()
}

func (i SystemReplacemsgWebproxyArray) ToSystemReplacemsgWebproxyArrayOutput() SystemReplacemsgWebproxyArrayOutput {
	return i.ToSystemReplacemsgWebproxyArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgWebproxyArray) ToSystemReplacemsgWebproxyArrayOutputWithContext(ctx context.Context) SystemReplacemsgWebproxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgWebproxyArrayOutput)
}

// SystemReplacemsgWebproxyMapInput is an input type that accepts SystemReplacemsgWebproxyMap and SystemReplacemsgWebproxyMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgWebproxyMapInput` via:
//
//	SystemReplacemsgWebproxyMap{ "key": SystemReplacemsgWebproxyArgs{...} }
type SystemReplacemsgWebproxyMapInput interface {
	pulumi.Input

	ToSystemReplacemsgWebproxyMapOutput() SystemReplacemsgWebproxyMapOutput
	ToSystemReplacemsgWebproxyMapOutputWithContext(context.Context) SystemReplacemsgWebproxyMapOutput
}

type SystemReplacemsgWebproxyMap map[string]SystemReplacemsgWebproxyInput

func (SystemReplacemsgWebproxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgWebproxy)(nil)).Elem()
}

func (i SystemReplacemsgWebproxyMap) ToSystemReplacemsgWebproxyMapOutput() SystemReplacemsgWebproxyMapOutput {
	return i.ToSystemReplacemsgWebproxyMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgWebproxyMap) ToSystemReplacemsgWebproxyMapOutputWithContext(ctx context.Context) SystemReplacemsgWebproxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgWebproxyMapOutput)
}

type SystemReplacemsgWebproxyOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgWebproxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgWebproxy)(nil)).Elem()
}

func (o SystemReplacemsgWebproxyOutput) ToSystemReplacemsgWebproxyOutput() SystemReplacemsgWebproxyOutput {
	return o
}

func (o SystemReplacemsgWebproxyOutput) ToSystemReplacemsgWebproxyOutputWithContext(ctx context.Context) SystemReplacemsgWebproxyOutput {
	return o
}

func (o SystemReplacemsgWebproxyOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgWebproxy) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

func (o SystemReplacemsgWebproxyOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgWebproxy) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o SystemReplacemsgWebproxyOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgWebproxy) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

func (o SystemReplacemsgWebproxyOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgWebproxy) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

func (o SystemReplacemsgWebproxyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgWebproxy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemReplacemsgWebproxyArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgWebproxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgWebproxy)(nil)).Elem()
}

func (o SystemReplacemsgWebproxyArrayOutput) ToSystemReplacemsgWebproxyArrayOutput() SystemReplacemsgWebproxyArrayOutput {
	return o
}

func (o SystemReplacemsgWebproxyArrayOutput) ToSystemReplacemsgWebproxyArrayOutputWithContext(ctx context.Context) SystemReplacemsgWebproxyArrayOutput {
	return o
}

func (o SystemReplacemsgWebproxyArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgWebproxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgWebproxy {
		return vs[0].([]*SystemReplacemsgWebproxy)[vs[1].(int)]
	}).(SystemReplacemsgWebproxyOutput)
}

type SystemReplacemsgWebproxyMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgWebproxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgWebproxy)(nil)).Elem()
}

func (o SystemReplacemsgWebproxyMapOutput) ToSystemReplacemsgWebproxyMapOutput() SystemReplacemsgWebproxyMapOutput {
	return o
}

func (o SystemReplacemsgWebproxyMapOutput) ToSystemReplacemsgWebproxyMapOutputWithContext(ctx context.Context) SystemReplacemsgWebproxyMapOutput {
	return o
}

func (o SystemReplacemsgWebproxyMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgWebproxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgWebproxy {
		return vs[0].(map[string]*SystemReplacemsgWebproxy)[vs[1].(string)]
	}).(SystemReplacemsgWebproxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgWebproxyInput)(nil)).Elem(), &SystemReplacemsgWebproxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgWebproxyArrayInput)(nil)).Elem(), SystemReplacemsgWebproxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgWebproxyMapInput)(nil)).Elem(), SystemReplacemsgWebproxyMap{})
	pulumi.RegisterOutputType(SystemReplacemsgWebproxyOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgWebproxyArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgWebproxyMapOutput{})
}

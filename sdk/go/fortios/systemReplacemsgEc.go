// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemReplacemsgEc struct {
	pulumi.CustomResourceState

	Buffer    pulumi.StringPtrOutput `pulumi:"buffer"`
	Format    pulumi.StringOutput    `pulumi:"format"`
	Header    pulumi.StringOutput    `pulumi:"header"`
	MsgType   pulumi.StringOutput    `pulumi:"msgType"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgEc registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgEc(ctx *pulumi.Context,
	name string, args *SystemReplacemsgEcArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgEc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgEc
	err := ctx.RegisterResource("fortios:index/systemReplacemsgEc:SystemReplacemsgEc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgEc gets an existing SystemReplacemsgEc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgEc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgEcState, opts ...pulumi.ResourceOption) (*SystemReplacemsgEc, error) {
	var resource SystemReplacemsgEc
	err := ctx.ReadResource("fortios:index/systemReplacemsgEc:SystemReplacemsgEc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgEc resources.
type systemReplacemsgEcState struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgEcState struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgEcState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgEcState)(nil)).Elem()
}

type systemReplacemsgEcArgs struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   string  `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgEc resource.
type SystemReplacemsgEcArgs struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgEcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgEcArgs)(nil)).Elem()
}

type SystemReplacemsgEcInput interface {
	pulumi.Input

	ToSystemReplacemsgEcOutput() SystemReplacemsgEcOutput
	ToSystemReplacemsgEcOutputWithContext(ctx context.Context) SystemReplacemsgEcOutput
}

func (*SystemReplacemsgEc) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgEc)(nil)).Elem()
}

func (i *SystemReplacemsgEc) ToSystemReplacemsgEcOutput() SystemReplacemsgEcOutput {
	return i.ToSystemReplacemsgEcOutputWithContext(context.Background())
}

func (i *SystemReplacemsgEc) ToSystemReplacemsgEcOutputWithContext(ctx context.Context) SystemReplacemsgEcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgEcOutput)
}

// SystemReplacemsgEcArrayInput is an input type that accepts SystemReplacemsgEcArray and SystemReplacemsgEcArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgEcArrayInput` via:
//
//	SystemReplacemsgEcArray{ SystemReplacemsgEcArgs{...} }
type SystemReplacemsgEcArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgEcArrayOutput() SystemReplacemsgEcArrayOutput
	ToSystemReplacemsgEcArrayOutputWithContext(context.Context) SystemReplacemsgEcArrayOutput
}

type SystemReplacemsgEcArray []SystemReplacemsgEcInput

func (SystemReplacemsgEcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgEc)(nil)).Elem()
}

func (i SystemReplacemsgEcArray) ToSystemReplacemsgEcArrayOutput() SystemReplacemsgEcArrayOutput {
	return i.ToSystemReplacemsgEcArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgEcArray) ToSystemReplacemsgEcArrayOutputWithContext(ctx context.Context) SystemReplacemsgEcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgEcArrayOutput)
}

// SystemReplacemsgEcMapInput is an input type that accepts SystemReplacemsgEcMap and SystemReplacemsgEcMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgEcMapInput` via:
//
//	SystemReplacemsgEcMap{ "key": SystemReplacemsgEcArgs{...} }
type SystemReplacemsgEcMapInput interface {
	pulumi.Input

	ToSystemReplacemsgEcMapOutput() SystemReplacemsgEcMapOutput
	ToSystemReplacemsgEcMapOutputWithContext(context.Context) SystemReplacemsgEcMapOutput
}

type SystemReplacemsgEcMap map[string]SystemReplacemsgEcInput

func (SystemReplacemsgEcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgEc)(nil)).Elem()
}

func (i SystemReplacemsgEcMap) ToSystemReplacemsgEcMapOutput() SystemReplacemsgEcMapOutput {
	return i.ToSystemReplacemsgEcMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgEcMap) ToSystemReplacemsgEcMapOutputWithContext(ctx context.Context) SystemReplacemsgEcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgEcMapOutput)
}

type SystemReplacemsgEcOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgEcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgEc)(nil)).Elem()
}

func (o SystemReplacemsgEcOutput) ToSystemReplacemsgEcOutput() SystemReplacemsgEcOutput {
	return o
}

func (o SystemReplacemsgEcOutput) ToSystemReplacemsgEcOutputWithContext(ctx context.Context) SystemReplacemsgEcOutput {
	return o
}

func (o SystemReplacemsgEcOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgEc) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

func (o SystemReplacemsgEcOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgEc) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o SystemReplacemsgEcOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgEc) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

func (o SystemReplacemsgEcOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgEc) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

func (o SystemReplacemsgEcOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgEc) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemReplacemsgEcArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgEcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgEc)(nil)).Elem()
}

func (o SystemReplacemsgEcArrayOutput) ToSystemReplacemsgEcArrayOutput() SystemReplacemsgEcArrayOutput {
	return o
}

func (o SystemReplacemsgEcArrayOutput) ToSystemReplacemsgEcArrayOutputWithContext(ctx context.Context) SystemReplacemsgEcArrayOutput {
	return o
}

func (o SystemReplacemsgEcArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgEcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgEc {
		return vs[0].([]*SystemReplacemsgEc)[vs[1].(int)]
	}).(SystemReplacemsgEcOutput)
}

type SystemReplacemsgEcMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgEcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgEc)(nil)).Elem()
}

func (o SystemReplacemsgEcMapOutput) ToSystemReplacemsgEcMapOutput() SystemReplacemsgEcMapOutput {
	return o
}

func (o SystemReplacemsgEcMapOutput) ToSystemReplacemsgEcMapOutputWithContext(ctx context.Context) SystemReplacemsgEcMapOutput {
	return o
}

func (o SystemReplacemsgEcMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgEcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgEc {
		return vs[0].(map[string]*SystemReplacemsgEc)[vs[1].(string)]
	}).(SystemReplacemsgEcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgEcInput)(nil)).Elem(), &SystemReplacemsgEc{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgEcArrayInput)(nil)).Elem(), SystemReplacemsgEcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgEcMapInput)(nil)).Elem(), SystemReplacemsgEcMap{})
	pulumi.RegisterOutputType(SystemReplacemsgEcOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgEcArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgEcMapOutput{})
}

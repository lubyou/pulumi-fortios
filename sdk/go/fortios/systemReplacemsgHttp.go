// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemReplacemsgHttp struct {
	pulumi.CustomResourceState

	Buffer    pulumi.StringPtrOutput `pulumi:"buffer"`
	Format    pulumi.StringOutput    `pulumi:"format"`
	Header    pulumi.StringOutput    `pulumi:"header"`
	MsgType   pulumi.StringOutput    `pulumi:"msgType"`
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
	opts = pkgResourceDefaultOpts(opts)
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
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgHttpState struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgHttpState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgHttpState)(nil)).Elem()
}

type systemReplacemsgHttpArgs struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   string  `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgHttp resource.
type SystemReplacemsgHttpArgs struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringInput
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
	return reflect.TypeOf((**SystemReplacemsgHttp)(nil)).Elem()
}

func (i *SystemReplacemsgHttp) ToSystemReplacemsgHttpOutput() SystemReplacemsgHttpOutput {
	return i.ToSystemReplacemsgHttpOutputWithContext(context.Background())
}

func (i *SystemReplacemsgHttp) ToSystemReplacemsgHttpOutputWithContext(ctx context.Context) SystemReplacemsgHttpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgHttpOutput)
}

// SystemReplacemsgHttpArrayInput is an input type that accepts SystemReplacemsgHttpArray and SystemReplacemsgHttpArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgHttpArrayInput` via:
//
//	SystemReplacemsgHttpArray{ SystemReplacemsgHttpArgs{...} }
type SystemReplacemsgHttpArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgHttpArrayOutput() SystemReplacemsgHttpArrayOutput
	ToSystemReplacemsgHttpArrayOutputWithContext(context.Context) SystemReplacemsgHttpArrayOutput
}

type SystemReplacemsgHttpArray []SystemReplacemsgHttpInput

func (SystemReplacemsgHttpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgHttp)(nil)).Elem()
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
//	SystemReplacemsgHttpMap{ "key": SystemReplacemsgHttpArgs{...} }
type SystemReplacemsgHttpMapInput interface {
	pulumi.Input

	ToSystemReplacemsgHttpMapOutput() SystemReplacemsgHttpMapOutput
	ToSystemReplacemsgHttpMapOutputWithContext(context.Context) SystemReplacemsgHttpMapOutput
}

type SystemReplacemsgHttpMap map[string]SystemReplacemsgHttpInput

func (SystemReplacemsgHttpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgHttp)(nil)).Elem()
}

func (i SystemReplacemsgHttpMap) ToSystemReplacemsgHttpMapOutput() SystemReplacemsgHttpMapOutput {
	return i.ToSystemReplacemsgHttpMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgHttpMap) ToSystemReplacemsgHttpMapOutputWithContext(ctx context.Context) SystemReplacemsgHttpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgHttpMapOutput)
}

type SystemReplacemsgHttpOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgHttpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgHttp)(nil)).Elem()
}

func (o SystemReplacemsgHttpOutput) ToSystemReplacemsgHttpOutput() SystemReplacemsgHttpOutput {
	return o
}

func (o SystemReplacemsgHttpOutput) ToSystemReplacemsgHttpOutputWithContext(ctx context.Context) SystemReplacemsgHttpOutput {
	return o
}

func (o SystemReplacemsgHttpOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgHttp) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

func (o SystemReplacemsgHttpOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgHttp) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o SystemReplacemsgHttpOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgHttp) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

func (o SystemReplacemsgHttpOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgHttp) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

func (o SystemReplacemsgHttpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgHttp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemReplacemsgHttpArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgHttpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgHttp)(nil)).Elem()
}

func (o SystemReplacemsgHttpArrayOutput) ToSystemReplacemsgHttpArrayOutput() SystemReplacemsgHttpArrayOutput {
	return o
}

func (o SystemReplacemsgHttpArrayOutput) ToSystemReplacemsgHttpArrayOutputWithContext(ctx context.Context) SystemReplacemsgHttpArrayOutput {
	return o
}

func (o SystemReplacemsgHttpArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgHttpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgHttp {
		return vs[0].([]*SystemReplacemsgHttp)[vs[1].(int)]
	}).(SystemReplacemsgHttpOutput)
}

type SystemReplacemsgHttpMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgHttpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgHttp)(nil)).Elem()
}

func (o SystemReplacemsgHttpMapOutput) ToSystemReplacemsgHttpMapOutput() SystemReplacemsgHttpMapOutput {
	return o
}

func (o SystemReplacemsgHttpMapOutput) ToSystemReplacemsgHttpMapOutputWithContext(ctx context.Context) SystemReplacemsgHttpMapOutput {
	return o
}

func (o SystemReplacemsgHttpMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgHttpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgHttp {
		return vs[0].(map[string]*SystemReplacemsgHttp)[vs[1].(string)]
	}).(SystemReplacemsgHttpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgHttpInput)(nil)).Elem(), &SystemReplacemsgHttp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgHttpArrayInput)(nil)).Elem(), SystemReplacemsgHttpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgHttpMapInput)(nil)).Elem(), SystemReplacemsgHttpMap{})
	pulumi.RegisterOutputType(SystemReplacemsgHttpOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgHttpArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgHttpMapOutput{})
}

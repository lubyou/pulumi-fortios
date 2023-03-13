// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemDscpBasedPriority struct {
	pulumi.CustomResourceState

	Ds        pulumi.IntOutput       `pulumi:"ds"`
	Fosid     pulumi.IntOutput       `pulumi:"fosid"`
	Priority  pulumi.StringOutput    `pulumi:"priority"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemDscpBasedPriority registers a new resource with the given unique name, arguments, and options.
func NewSystemDscpBasedPriority(ctx *pulumi.Context,
	name string, args *SystemDscpBasedPriorityArgs, opts ...pulumi.ResourceOption) (*SystemDscpBasedPriority, error) {
	if args == nil {
		args = &SystemDscpBasedPriorityArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemDscpBasedPriority
	err := ctx.RegisterResource("fortios:index/systemDscpBasedPriority:SystemDscpBasedPriority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemDscpBasedPriority gets an existing SystemDscpBasedPriority resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemDscpBasedPriority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemDscpBasedPriorityState, opts ...pulumi.ResourceOption) (*SystemDscpBasedPriority, error) {
	var resource SystemDscpBasedPriority
	err := ctx.ReadResource("fortios:index/systemDscpBasedPriority:SystemDscpBasedPriority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemDscpBasedPriority resources.
type systemDscpBasedPriorityState struct {
	Ds        *int    `pulumi:"ds"`
	Fosid     *int    `pulumi:"fosid"`
	Priority  *string `pulumi:"priority"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemDscpBasedPriorityState struct {
	Ds        pulumi.IntPtrInput
	Fosid     pulumi.IntPtrInput
	Priority  pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemDscpBasedPriorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDscpBasedPriorityState)(nil)).Elem()
}

type systemDscpBasedPriorityArgs struct {
	Ds        *int    `pulumi:"ds"`
	Fosid     *int    `pulumi:"fosid"`
	Priority  *string `pulumi:"priority"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemDscpBasedPriority resource.
type SystemDscpBasedPriorityArgs struct {
	Ds        pulumi.IntPtrInput
	Fosid     pulumi.IntPtrInput
	Priority  pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemDscpBasedPriorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDscpBasedPriorityArgs)(nil)).Elem()
}

type SystemDscpBasedPriorityInput interface {
	pulumi.Input

	ToSystemDscpBasedPriorityOutput() SystemDscpBasedPriorityOutput
	ToSystemDscpBasedPriorityOutputWithContext(ctx context.Context) SystemDscpBasedPriorityOutput
}

func (*SystemDscpBasedPriority) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDscpBasedPriority)(nil)).Elem()
}

func (i *SystemDscpBasedPriority) ToSystemDscpBasedPriorityOutput() SystemDscpBasedPriorityOutput {
	return i.ToSystemDscpBasedPriorityOutputWithContext(context.Background())
}

func (i *SystemDscpBasedPriority) ToSystemDscpBasedPriorityOutputWithContext(ctx context.Context) SystemDscpBasedPriorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDscpBasedPriorityOutput)
}

// SystemDscpBasedPriorityArrayInput is an input type that accepts SystemDscpBasedPriorityArray and SystemDscpBasedPriorityArrayOutput values.
// You can construct a concrete instance of `SystemDscpBasedPriorityArrayInput` via:
//
//	SystemDscpBasedPriorityArray{ SystemDscpBasedPriorityArgs{...} }
type SystemDscpBasedPriorityArrayInput interface {
	pulumi.Input

	ToSystemDscpBasedPriorityArrayOutput() SystemDscpBasedPriorityArrayOutput
	ToSystemDscpBasedPriorityArrayOutputWithContext(context.Context) SystemDscpBasedPriorityArrayOutput
}

type SystemDscpBasedPriorityArray []SystemDscpBasedPriorityInput

func (SystemDscpBasedPriorityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDscpBasedPriority)(nil)).Elem()
}

func (i SystemDscpBasedPriorityArray) ToSystemDscpBasedPriorityArrayOutput() SystemDscpBasedPriorityArrayOutput {
	return i.ToSystemDscpBasedPriorityArrayOutputWithContext(context.Background())
}

func (i SystemDscpBasedPriorityArray) ToSystemDscpBasedPriorityArrayOutputWithContext(ctx context.Context) SystemDscpBasedPriorityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDscpBasedPriorityArrayOutput)
}

// SystemDscpBasedPriorityMapInput is an input type that accepts SystemDscpBasedPriorityMap and SystemDscpBasedPriorityMapOutput values.
// You can construct a concrete instance of `SystemDscpBasedPriorityMapInput` via:
//
//	SystemDscpBasedPriorityMap{ "key": SystemDscpBasedPriorityArgs{...} }
type SystemDscpBasedPriorityMapInput interface {
	pulumi.Input

	ToSystemDscpBasedPriorityMapOutput() SystemDscpBasedPriorityMapOutput
	ToSystemDscpBasedPriorityMapOutputWithContext(context.Context) SystemDscpBasedPriorityMapOutput
}

type SystemDscpBasedPriorityMap map[string]SystemDscpBasedPriorityInput

func (SystemDscpBasedPriorityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDscpBasedPriority)(nil)).Elem()
}

func (i SystemDscpBasedPriorityMap) ToSystemDscpBasedPriorityMapOutput() SystemDscpBasedPriorityMapOutput {
	return i.ToSystemDscpBasedPriorityMapOutputWithContext(context.Background())
}

func (i SystemDscpBasedPriorityMap) ToSystemDscpBasedPriorityMapOutputWithContext(ctx context.Context) SystemDscpBasedPriorityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDscpBasedPriorityMapOutput)
}

type SystemDscpBasedPriorityOutput struct{ *pulumi.OutputState }

func (SystemDscpBasedPriorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDscpBasedPriority)(nil)).Elem()
}

func (o SystemDscpBasedPriorityOutput) ToSystemDscpBasedPriorityOutput() SystemDscpBasedPriorityOutput {
	return o
}

func (o SystemDscpBasedPriorityOutput) ToSystemDscpBasedPriorityOutputWithContext(ctx context.Context) SystemDscpBasedPriorityOutput {
	return o
}

func (o SystemDscpBasedPriorityOutput) Ds() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDscpBasedPriority) pulumi.IntOutput { return v.Ds }).(pulumi.IntOutput)
}

func (o SystemDscpBasedPriorityOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDscpBasedPriority) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o SystemDscpBasedPriorityOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDscpBasedPriority) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

func (o SystemDscpBasedPriorityOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDscpBasedPriority) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemDscpBasedPriorityArrayOutput struct{ *pulumi.OutputState }

func (SystemDscpBasedPriorityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDscpBasedPriority)(nil)).Elem()
}

func (o SystemDscpBasedPriorityArrayOutput) ToSystemDscpBasedPriorityArrayOutput() SystemDscpBasedPriorityArrayOutput {
	return o
}

func (o SystemDscpBasedPriorityArrayOutput) ToSystemDscpBasedPriorityArrayOutputWithContext(ctx context.Context) SystemDscpBasedPriorityArrayOutput {
	return o
}

func (o SystemDscpBasedPriorityArrayOutput) Index(i pulumi.IntInput) SystemDscpBasedPriorityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemDscpBasedPriority {
		return vs[0].([]*SystemDscpBasedPriority)[vs[1].(int)]
	}).(SystemDscpBasedPriorityOutput)
}

type SystemDscpBasedPriorityMapOutput struct{ *pulumi.OutputState }

func (SystemDscpBasedPriorityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDscpBasedPriority)(nil)).Elem()
}

func (o SystemDscpBasedPriorityMapOutput) ToSystemDscpBasedPriorityMapOutput() SystemDscpBasedPriorityMapOutput {
	return o
}

func (o SystemDscpBasedPriorityMapOutput) ToSystemDscpBasedPriorityMapOutputWithContext(ctx context.Context) SystemDscpBasedPriorityMapOutput {
	return o
}

func (o SystemDscpBasedPriorityMapOutput) MapIndex(k pulumi.StringInput) SystemDscpBasedPriorityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemDscpBasedPriority {
		return vs[0].(map[string]*SystemDscpBasedPriority)[vs[1].(string)]
	}).(SystemDscpBasedPriorityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDscpBasedPriorityInput)(nil)).Elem(), &SystemDscpBasedPriority{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDscpBasedPriorityArrayInput)(nil)).Elem(), SystemDscpBasedPriorityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDscpBasedPriorityMapInput)(nil)).Elem(), SystemDscpBasedPriorityMap{})
	pulumi.RegisterOutputType(SystemDscpBasedPriorityOutput{})
	pulumi.RegisterOutputType(SystemDscpBasedPriorityArrayOutput{})
	pulumi.RegisterOutputType(SystemDscpBasedPriorityMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type FortimanagerSystemNTP struct {
	pulumi.CustomResourceState

	Server       pulumi.StringOutput    `pulumi:"server"`
	Status       pulumi.StringPtrOutput `pulumi:"status"`
	SyncInterval pulumi.IntPtrOutput    `pulumi:"syncInterval"`
}

// NewFortimanagerSystemNTP registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemNTP(ctx *pulumi.Context,
	name string, args *FortimanagerSystemNTPArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemNTP, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FortimanagerSystemNTP
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemNTP:FortimanagerSystemNTP", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemNTP gets an existing FortimanagerSystemNTP resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemNTP(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemNTPState, opts ...pulumi.ResourceOption) (*FortimanagerSystemNTP, error) {
	var resource FortimanagerSystemNTP
	err := ctx.ReadResource("fortios:index/fortimanagerSystemNTP:FortimanagerSystemNTP", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemNTP resources.
type fortimanagerSystemNTPState struct {
	Server       *string `pulumi:"server"`
	Status       *string `pulumi:"status"`
	SyncInterval *int    `pulumi:"syncInterval"`
}

type FortimanagerSystemNTPState struct {
	Server       pulumi.StringPtrInput
	Status       pulumi.StringPtrInput
	SyncInterval pulumi.IntPtrInput
}

func (FortimanagerSystemNTPState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemNTPState)(nil)).Elem()
}

type fortimanagerSystemNTPArgs struct {
	Server       string  `pulumi:"server"`
	Status       *string `pulumi:"status"`
	SyncInterval *int    `pulumi:"syncInterval"`
}

// The set of arguments for constructing a FortimanagerSystemNTP resource.
type FortimanagerSystemNTPArgs struct {
	Server       pulumi.StringInput
	Status       pulumi.StringPtrInput
	SyncInterval pulumi.IntPtrInput
}

func (FortimanagerSystemNTPArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemNTPArgs)(nil)).Elem()
}

type FortimanagerSystemNTPInput interface {
	pulumi.Input

	ToFortimanagerSystemNTPOutput() FortimanagerSystemNTPOutput
	ToFortimanagerSystemNTPOutputWithContext(ctx context.Context) FortimanagerSystemNTPOutput
}

func (*FortimanagerSystemNTP) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemNTP)(nil)).Elem()
}

func (i *FortimanagerSystemNTP) ToFortimanagerSystemNTPOutput() FortimanagerSystemNTPOutput {
	return i.ToFortimanagerSystemNTPOutputWithContext(context.Background())
}

func (i *FortimanagerSystemNTP) ToFortimanagerSystemNTPOutputWithContext(ctx context.Context) FortimanagerSystemNTPOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemNTPOutput)
}

func (i *FortimanagerSystemNTP) ToOutput(ctx context.Context) pulumix.Output[*FortimanagerSystemNTP] {
	return pulumix.Output[*FortimanagerSystemNTP]{
		OutputState: i.ToFortimanagerSystemNTPOutputWithContext(ctx).OutputState,
	}
}

// FortimanagerSystemNTPArrayInput is an input type that accepts FortimanagerSystemNTPArray and FortimanagerSystemNTPArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemNTPArrayInput` via:
//
//	FortimanagerSystemNTPArray{ FortimanagerSystemNTPArgs{...} }
type FortimanagerSystemNTPArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemNTPArrayOutput() FortimanagerSystemNTPArrayOutput
	ToFortimanagerSystemNTPArrayOutputWithContext(context.Context) FortimanagerSystemNTPArrayOutput
}

type FortimanagerSystemNTPArray []FortimanagerSystemNTPInput

func (FortimanagerSystemNTPArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemNTP)(nil)).Elem()
}

func (i FortimanagerSystemNTPArray) ToFortimanagerSystemNTPArrayOutput() FortimanagerSystemNTPArrayOutput {
	return i.ToFortimanagerSystemNTPArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemNTPArray) ToFortimanagerSystemNTPArrayOutputWithContext(ctx context.Context) FortimanagerSystemNTPArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemNTPArrayOutput)
}

func (i FortimanagerSystemNTPArray) ToOutput(ctx context.Context) pulumix.Output[[]*FortimanagerSystemNTP] {
	return pulumix.Output[[]*FortimanagerSystemNTP]{
		OutputState: i.ToFortimanagerSystemNTPArrayOutputWithContext(ctx).OutputState,
	}
}

// FortimanagerSystemNTPMapInput is an input type that accepts FortimanagerSystemNTPMap and FortimanagerSystemNTPMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemNTPMapInput` via:
//
//	FortimanagerSystemNTPMap{ "key": FortimanagerSystemNTPArgs{...} }
type FortimanagerSystemNTPMapInput interface {
	pulumi.Input

	ToFortimanagerSystemNTPMapOutput() FortimanagerSystemNTPMapOutput
	ToFortimanagerSystemNTPMapOutputWithContext(context.Context) FortimanagerSystemNTPMapOutput
}

type FortimanagerSystemNTPMap map[string]FortimanagerSystemNTPInput

func (FortimanagerSystemNTPMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemNTP)(nil)).Elem()
}

func (i FortimanagerSystemNTPMap) ToFortimanagerSystemNTPMapOutput() FortimanagerSystemNTPMapOutput {
	return i.ToFortimanagerSystemNTPMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemNTPMap) ToFortimanagerSystemNTPMapOutputWithContext(ctx context.Context) FortimanagerSystemNTPMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemNTPMapOutput)
}

func (i FortimanagerSystemNTPMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FortimanagerSystemNTP] {
	return pulumix.Output[map[string]*FortimanagerSystemNTP]{
		OutputState: i.ToFortimanagerSystemNTPMapOutputWithContext(ctx).OutputState,
	}
}

type FortimanagerSystemNTPOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemNTPOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemNTP)(nil)).Elem()
}

func (o FortimanagerSystemNTPOutput) ToFortimanagerSystemNTPOutput() FortimanagerSystemNTPOutput {
	return o
}

func (o FortimanagerSystemNTPOutput) ToFortimanagerSystemNTPOutputWithContext(ctx context.Context) FortimanagerSystemNTPOutput {
	return o
}

func (o FortimanagerSystemNTPOutput) ToOutput(ctx context.Context) pulumix.Output[*FortimanagerSystemNTP] {
	return pulumix.Output[*FortimanagerSystemNTP]{
		OutputState: o.OutputState,
	}
}

func (o FortimanagerSystemNTPOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerSystemNTP) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o FortimanagerSystemNTPOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemNTP) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemNTPOutput) SyncInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemNTP) pulumi.IntPtrOutput { return v.SyncInterval }).(pulumi.IntPtrOutput)
}

type FortimanagerSystemNTPArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemNTPArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemNTP)(nil)).Elem()
}

func (o FortimanagerSystemNTPArrayOutput) ToFortimanagerSystemNTPArrayOutput() FortimanagerSystemNTPArrayOutput {
	return o
}

func (o FortimanagerSystemNTPArrayOutput) ToFortimanagerSystemNTPArrayOutputWithContext(ctx context.Context) FortimanagerSystemNTPArrayOutput {
	return o
}

func (o FortimanagerSystemNTPArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FortimanagerSystemNTP] {
	return pulumix.Output[[]*FortimanagerSystemNTP]{
		OutputState: o.OutputState,
	}
}

func (o FortimanagerSystemNTPArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemNTPOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerSystemNTP {
		return vs[0].([]*FortimanagerSystemNTP)[vs[1].(int)]
	}).(FortimanagerSystemNTPOutput)
}

type FortimanagerSystemNTPMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemNTPMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemNTP)(nil)).Elem()
}

func (o FortimanagerSystemNTPMapOutput) ToFortimanagerSystemNTPMapOutput() FortimanagerSystemNTPMapOutput {
	return o
}

func (o FortimanagerSystemNTPMapOutput) ToFortimanagerSystemNTPMapOutputWithContext(ctx context.Context) FortimanagerSystemNTPMapOutput {
	return o
}

func (o FortimanagerSystemNTPMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FortimanagerSystemNTP] {
	return pulumix.Output[map[string]*FortimanagerSystemNTP]{
		OutputState: o.OutputState,
	}
}

func (o FortimanagerSystemNTPMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemNTPOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerSystemNTP {
		return vs[0].(map[string]*FortimanagerSystemNTP)[vs[1].(string)]
	}).(FortimanagerSystemNTPOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemNTPInput)(nil)).Elem(), &FortimanagerSystemNTP{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemNTPArrayInput)(nil)).Elem(), FortimanagerSystemNTPArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemNTPMapInput)(nil)).Elem(), FortimanagerSystemNTPMap{})
	pulumi.RegisterOutputType(FortimanagerSystemNTPOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemNTPArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemNTPMapOutput{})
}

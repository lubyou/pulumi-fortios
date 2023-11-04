// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type ExtensionControllerFortigate struct {
	pulumi.CustomResourceState

	Authorized  pulumi.StringOutput    `pulumi:"authorized"`
	Description pulumi.StringOutput    `pulumi:"description"`
	DeviceId    pulumi.IntOutput       `pulumi:"deviceId"`
	Fosid       pulumi.StringOutput    `pulumi:"fosid"`
	Hostname    pulumi.StringOutput    `pulumi:"hostname"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Profile     pulumi.StringOutput    `pulumi:"profile"`
	Vdom        pulumi.IntOutput       `pulumi:"vdom"`
	Vdomparam   pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewExtensionControllerFortigate registers a new resource with the given unique name, arguments, and options.
func NewExtensionControllerFortigate(ctx *pulumi.Context,
	name string, args *ExtensionControllerFortigateArgs, opts ...pulumi.ResourceOption) (*ExtensionControllerFortigate, error) {
	if args == nil {
		args = &ExtensionControllerFortigateArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExtensionControllerFortigate
	err := ctx.RegisterResource("fortios:index/extensionControllerFortigate:ExtensionControllerFortigate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtensionControllerFortigate gets an existing ExtensionControllerFortigate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtensionControllerFortigate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionControllerFortigateState, opts ...pulumi.ResourceOption) (*ExtensionControllerFortigate, error) {
	var resource ExtensionControllerFortigate
	err := ctx.ReadResource("fortios:index/extensionControllerFortigate:ExtensionControllerFortigate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExtensionControllerFortigate resources.
type extensionControllerFortigateState struct {
	Authorized  *string `pulumi:"authorized"`
	Description *string `pulumi:"description"`
	DeviceId    *int    `pulumi:"deviceId"`
	Fosid       *string `pulumi:"fosid"`
	Hostname    *string `pulumi:"hostname"`
	Name        *string `pulumi:"name"`
	Profile     *string `pulumi:"profile"`
	Vdom        *int    `pulumi:"vdom"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

type ExtensionControllerFortigateState struct {
	Authorized  pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	DeviceId    pulumi.IntPtrInput
	Fosid       pulumi.StringPtrInput
	Hostname    pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Profile     pulumi.StringPtrInput
	Vdom        pulumi.IntPtrInput
	Vdomparam   pulumi.StringPtrInput
}

func (ExtensionControllerFortigateState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionControllerFortigateState)(nil)).Elem()
}

type extensionControllerFortigateArgs struct {
	Authorized  *string `pulumi:"authorized"`
	Description *string `pulumi:"description"`
	DeviceId    *int    `pulumi:"deviceId"`
	Fosid       *string `pulumi:"fosid"`
	Hostname    *string `pulumi:"hostname"`
	Name        *string `pulumi:"name"`
	Profile     *string `pulumi:"profile"`
	Vdom        *int    `pulumi:"vdom"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a ExtensionControllerFortigate resource.
type ExtensionControllerFortigateArgs struct {
	Authorized  pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	DeviceId    pulumi.IntPtrInput
	Fosid       pulumi.StringPtrInput
	Hostname    pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Profile     pulumi.StringPtrInput
	Vdom        pulumi.IntPtrInput
	Vdomparam   pulumi.StringPtrInput
}

func (ExtensionControllerFortigateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionControllerFortigateArgs)(nil)).Elem()
}

type ExtensionControllerFortigateInput interface {
	pulumi.Input

	ToExtensionControllerFortigateOutput() ExtensionControllerFortigateOutput
	ToExtensionControllerFortigateOutputWithContext(ctx context.Context) ExtensionControllerFortigateOutput
}

func (*ExtensionControllerFortigate) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionControllerFortigate)(nil)).Elem()
}

func (i *ExtensionControllerFortigate) ToExtensionControllerFortigateOutput() ExtensionControllerFortigateOutput {
	return i.ToExtensionControllerFortigateOutputWithContext(context.Background())
}

func (i *ExtensionControllerFortigate) ToExtensionControllerFortigateOutputWithContext(ctx context.Context) ExtensionControllerFortigateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionControllerFortigateOutput)
}

func (i *ExtensionControllerFortigate) ToOutput(ctx context.Context) pulumix.Output[*ExtensionControllerFortigate] {
	return pulumix.Output[*ExtensionControllerFortigate]{
		OutputState: i.ToExtensionControllerFortigateOutputWithContext(ctx).OutputState,
	}
}

// ExtensionControllerFortigateArrayInput is an input type that accepts ExtensionControllerFortigateArray and ExtensionControllerFortigateArrayOutput values.
// You can construct a concrete instance of `ExtensionControllerFortigateArrayInput` via:
//
//	ExtensionControllerFortigateArray{ ExtensionControllerFortigateArgs{...} }
type ExtensionControllerFortigateArrayInput interface {
	pulumi.Input

	ToExtensionControllerFortigateArrayOutput() ExtensionControllerFortigateArrayOutput
	ToExtensionControllerFortigateArrayOutputWithContext(context.Context) ExtensionControllerFortigateArrayOutput
}

type ExtensionControllerFortigateArray []ExtensionControllerFortigateInput

func (ExtensionControllerFortigateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensionControllerFortigate)(nil)).Elem()
}

func (i ExtensionControllerFortigateArray) ToExtensionControllerFortigateArrayOutput() ExtensionControllerFortigateArrayOutput {
	return i.ToExtensionControllerFortigateArrayOutputWithContext(context.Background())
}

func (i ExtensionControllerFortigateArray) ToExtensionControllerFortigateArrayOutputWithContext(ctx context.Context) ExtensionControllerFortigateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionControllerFortigateArrayOutput)
}

func (i ExtensionControllerFortigateArray) ToOutput(ctx context.Context) pulumix.Output[[]*ExtensionControllerFortigate] {
	return pulumix.Output[[]*ExtensionControllerFortigate]{
		OutputState: i.ToExtensionControllerFortigateArrayOutputWithContext(ctx).OutputState,
	}
}

// ExtensionControllerFortigateMapInput is an input type that accepts ExtensionControllerFortigateMap and ExtensionControllerFortigateMapOutput values.
// You can construct a concrete instance of `ExtensionControllerFortigateMapInput` via:
//
//	ExtensionControllerFortigateMap{ "key": ExtensionControllerFortigateArgs{...} }
type ExtensionControllerFortigateMapInput interface {
	pulumi.Input

	ToExtensionControllerFortigateMapOutput() ExtensionControllerFortigateMapOutput
	ToExtensionControllerFortigateMapOutputWithContext(context.Context) ExtensionControllerFortigateMapOutput
}

type ExtensionControllerFortigateMap map[string]ExtensionControllerFortigateInput

func (ExtensionControllerFortigateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensionControllerFortigate)(nil)).Elem()
}

func (i ExtensionControllerFortigateMap) ToExtensionControllerFortigateMapOutput() ExtensionControllerFortigateMapOutput {
	return i.ToExtensionControllerFortigateMapOutputWithContext(context.Background())
}

func (i ExtensionControllerFortigateMap) ToExtensionControllerFortigateMapOutputWithContext(ctx context.Context) ExtensionControllerFortigateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionControllerFortigateMapOutput)
}

func (i ExtensionControllerFortigateMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ExtensionControllerFortigate] {
	return pulumix.Output[map[string]*ExtensionControllerFortigate]{
		OutputState: i.ToExtensionControllerFortigateMapOutputWithContext(ctx).OutputState,
	}
}

type ExtensionControllerFortigateOutput struct{ *pulumi.OutputState }

func (ExtensionControllerFortigateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionControllerFortigate)(nil)).Elem()
}

func (o ExtensionControllerFortigateOutput) ToExtensionControllerFortigateOutput() ExtensionControllerFortigateOutput {
	return o
}

func (o ExtensionControllerFortigateOutput) ToExtensionControllerFortigateOutputWithContext(ctx context.Context) ExtensionControllerFortigateOutput {
	return o
}

func (o ExtensionControllerFortigateOutput) ToOutput(ctx context.Context) pulumix.Output[*ExtensionControllerFortigate] {
	return pulumix.Output[*ExtensionControllerFortigate]{
		OutputState: o.OutputState,
	}
}

func (o ExtensionControllerFortigateOutput) Authorized() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigate) pulumi.StringOutput { return v.Authorized }).(pulumi.StringOutput)
}

func (o ExtensionControllerFortigateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ExtensionControllerFortigateOutput) DeviceId() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigate) pulumi.IntOutput { return v.DeviceId }).(pulumi.IntOutput)
}

func (o ExtensionControllerFortigateOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigate) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

func (o ExtensionControllerFortigateOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigate) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

func (o ExtensionControllerFortigateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExtensionControllerFortigateOutput) Profile() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigate) pulumi.StringOutput { return v.Profile }).(pulumi.StringOutput)
}

func (o ExtensionControllerFortigateOutput) Vdom() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigate) pulumi.IntOutput { return v.Vdom }).(pulumi.IntOutput)
}

func (o ExtensionControllerFortigateOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigate) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ExtensionControllerFortigateArrayOutput struct{ *pulumi.OutputState }

func (ExtensionControllerFortigateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensionControllerFortigate)(nil)).Elem()
}

func (o ExtensionControllerFortigateArrayOutput) ToExtensionControllerFortigateArrayOutput() ExtensionControllerFortigateArrayOutput {
	return o
}

func (o ExtensionControllerFortigateArrayOutput) ToExtensionControllerFortigateArrayOutputWithContext(ctx context.Context) ExtensionControllerFortigateArrayOutput {
	return o
}

func (o ExtensionControllerFortigateArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ExtensionControllerFortigate] {
	return pulumix.Output[[]*ExtensionControllerFortigate]{
		OutputState: o.OutputState,
	}
}

func (o ExtensionControllerFortigateArrayOutput) Index(i pulumi.IntInput) ExtensionControllerFortigateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExtensionControllerFortigate {
		return vs[0].([]*ExtensionControllerFortigate)[vs[1].(int)]
	}).(ExtensionControllerFortigateOutput)
}

type ExtensionControllerFortigateMapOutput struct{ *pulumi.OutputState }

func (ExtensionControllerFortigateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensionControllerFortigate)(nil)).Elem()
}

func (o ExtensionControllerFortigateMapOutput) ToExtensionControllerFortigateMapOutput() ExtensionControllerFortigateMapOutput {
	return o
}

func (o ExtensionControllerFortigateMapOutput) ToExtensionControllerFortigateMapOutputWithContext(ctx context.Context) ExtensionControllerFortigateMapOutput {
	return o
}

func (o ExtensionControllerFortigateMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ExtensionControllerFortigate] {
	return pulumix.Output[map[string]*ExtensionControllerFortigate]{
		OutputState: o.OutputState,
	}
}

func (o ExtensionControllerFortigateMapOutput) MapIndex(k pulumi.StringInput) ExtensionControllerFortigateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExtensionControllerFortigate {
		return vs[0].(map[string]*ExtensionControllerFortigate)[vs[1].(string)]
	}).(ExtensionControllerFortigateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionControllerFortigateInput)(nil)).Elem(), &ExtensionControllerFortigate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionControllerFortigateArrayInput)(nil)).Elem(), ExtensionControllerFortigateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionControllerFortigateMapInput)(nil)).Elem(), ExtensionControllerFortigateMap{})
	pulumi.RegisterOutputType(ExtensionControllerFortigateOutput{})
	pulumi.RegisterOutputType(ExtensionControllerFortigateArrayOutput{})
	pulumi.RegisterOutputType(ExtensionControllerFortigateMapOutput{})
}

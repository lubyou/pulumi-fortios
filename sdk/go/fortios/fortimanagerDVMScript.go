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

type FortimanagerDVMScript struct {
	pulumi.CustomResourceState

	Adom        pulumi.StringPtrOutput `pulumi:"adom"`
	Content     pulumi.StringOutput    `pulumi:"content"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Target      pulumi.StringPtrOutput `pulumi:"target"`
}

// NewFortimanagerDVMScript registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerDVMScript(ctx *pulumi.Context,
	name string, args *FortimanagerDVMScriptArgs, opts ...pulumi.ResourceOption) (*FortimanagerDVMScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FortimanagerDVMScript
	err := ctx.RegisterResource("fortios:index/fortimanagerDVMScript:FortimanagerDVMScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerDVMScript gets an existing FortimanagerDVMScript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerDVMScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerDVMScriptState, opts ...pulumi.ResourceOption) (*FortimanagerDVMScript, error) {
	var resource FortimanagerDVMScript
	err := ctx.ReadResource("fortios:index/fortimanagerDVMScript:FortimanagerDVMScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerDVMScript resources.
type fortimanagerDVMScriptState struct {
	Adom        *string `pulumi:"adom"`
	Content     *string `pulumi:"content"`
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Target      *string `pulumi:"target"`
}

type FortimanagerDVMScriptState struct {
	Adom        pulumi.StringPtrInput
	Content     pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Target      pulumi.StringPtrInput
}

func (FortimanagerDVMScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerDVMScriptState)(nil)).Elem()
}

type fortimanagerDVMScriptArgs struct {
	Adom        *string `pulumi:"adom"`
	Content     string  `pulumi:"content"`
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Target      *string `pulumi:"target"`
}

// The set of arguments for constructing a FortimanagerDVMScript resource.
type FortimanagerDVMScriptArgs struct {
	Adom        pulumi.StringPtrInput
	Content     pulumi.StringInput
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Target      pulumi.StringPtrInput
}

func (FortimanagerDVMScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerDVMScriptArgs)(nil)).Elem()
}

type FortimanagerDVMScriptInput interface {
	pulumi.Input

	ToFortimanagerDVMScriptOutput() FortimanagerDVMScriptOutput
	ToFortimanagerDVMScriptOutputWithContext(ctx context.Context) FortimanagerDVMScriptOutput
}

func (*FortimanagerDVMScript) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerDVMScript)(nil)).Elem()
}

func (i *FortimanagerDVMScript) ToFortimanagerDVMScriptOutput() FortimanagerDVMScriptOutput {
	return i.ToFortimanagerDVMScriptOutputWithContext(context.Background())
}

func (i *FortimanagerDVMScript) ToFortimanagerDVMScriptOutputWithContext(ctx context.Context) FortimanagerDVMScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerDVMScriptOutput)
}

// FortimanagerDVMScriptArrayInput is an input type that accepts FortimanagerDVMScriptArray and FortimanagerDVMScriptArrayOutput values.
// You can construct a concrete instance of `FortimanagerDVMScriptArrayInput` via:
//
//	FortimanagerDVMScriptArray{ FortimanagerDVMScriptArgs{...} }
type FortimanagerDVMScriptArrayInput interface {
	pulumi.Input

	ToFortimanagerDVMScriptArrayOutput() FortimanagerDVMScriptArrayOutput
	ToFortimanagerDVMScriptArrayOutputWithContext(context.Context) FortimanagerDVMScriptArrayOutput
}

type FortimanagerDVMScriptArray []FortimanagerDVMScriptInput

func (FortimanagerDVMScriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerDVMScript)(nil)).Elem()
}

func (i FortimanagerDVMScriptArray) ToFortimanagerDVMScriptArrayOutput() FortimanagerDVMScriptArrayOutput {
	return i.ToFortimanagerDVMScriptArrayOutputWithContext(context.Background())
}

func (i FortimanagerDVMScriptArray) ToFortimanagerDVMScriptArrayOutputWithContext(ctx context.Context) FortimanagerDVMScriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerDVMScriptArrayOutput)
}

// FortimanagerDVMScriptMapInput is an input type that accepts FortimanagerDVMScriptMap and FortimanagerDVMScriptMapOutput values.
// You can construct a concrete instance of `FortimanagerDVMScriptMapInput` via:
//
//	FortimanagerDVMScriptMap{ "key": FortimanagerDVMScriptArgs{...} }
type FortimanagerDVMScriptMapInput interface {
	pulumi.Input

	ToFortimanagerDVMScriptMapOutput() FortimanagerDVMScriptMapOutput
	ToFortimanagerDVMScriptMapOutputWithContext(context.Context) FortimanagerDVMScriptMapOutput
}

type FortimanagerDVMScriptMap map[string]FortimanagerDVMScriptInput

func (FortimanagerDVMScriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerDVMScript)(nil)).Elem()
}

func (i FortimanagerDVMScriptMap) ToFortimanagerDVMScriptMapOutput() FortimanagerDVMScriptMapOutput {
	return i.ToFortimanagerDVMScriptMapOutputWithContext(context.Background())
}

func (i FortimanagerDVMScriptMap) ToFortimanagerDVMScriptMapOutputWithContext(ctx context.Context) FortimanagerDVMScriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerDVMScriptMapOutput)
}

type FortimanagerDVMScriptOutput struct{ *pulumi.OutputState }

func (FortimanagerDVMScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerDVMScript)(nil)).Elem()
}

func (o FortimanagerDVMScriptOutput) ToFortimanagerDVMScriptOutput() FortimanagerDVMScriptOutput {
	return o
}

func (o FortimanagerDVMScriptOutput) ToFortimanagerDVMScriptOutputWithContext(ctx context.Context) FortimanagerDVMScriptOutput {
	return o
}

func (o FortimanagerDVMScriptOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerDVMScript) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

func (o FortimanagerDVMScriptOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerDVMScript) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

func (o FortimanagerDVMScriptOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerDVMScript) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o FortimanagerDVMScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerDVMScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FortimanagerDVMScriptOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerDVMScript) pulumi.StringPtrOutput { return v.Target }).(pulumi.StringPtrOutput)
}

type FortimanagerDVMScriptArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerDVMScriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerDVMScript)(nil)).Elem()
}

func (o FortimanagerDVMScriptArrayOutput) ToFortimanagerDVMScriptArrayOutput() FortimanagerDVMScriptArrayOutput {
	return o
}

func (o FortimanagerDVMScriptArrayOutput) ToFortimanagerDVMScriptArrayOutputWithContext(ctx context.Context) FortimanagerDVMScriptArrayOutput {
	return o
}

func (o FortimanagerDVMScriptArrayOutput) Index(i pulumi.IntInput) FortimanagerDVMScriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerDVMScript {
		return vs[0].([]*FortimanagerDVMScript)[vs[1].(int)]
	}).(FortimanagerDVMScriptOutput)
}

type FortimanagerDVMScriptMapOutput struct{ *pulumi.OutputState }

func (FortimanagerDVMScriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerDVMScript)(nil)).Elem()
}

func (o FortimanagerDVMScriptMapOutput) ToFortimanagerDVMScriptMapOutput() FortimanagerDVMScriptMapOutput {
	return o
}

func (o FortimanagerDVMScriptMapOutput) ToFortimanagerDVMScriptMapOutputWithContext(ctx context.Context) FortimanagerDVMScriptMapOutput {
	return o
}

func (o FortimanagerDVMScriptMapOutput) MapIndex(k pulumi.StringInput) FortimanagerDVMScriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerDVMScript {
		return vs[0].(map[string]*FortimanagerDVMScript)[vs[1].(string)]
	}).(FortimanagerDVMScriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerDVMScriptInput)(nil)).Elem(), &FortimanagerDVMScript{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerDVMScriptArrayInput)(nil)).Elem(), FortimanagerDVMScriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerDVMScriptMapInput)(nil)).Elem(), FortimanagerDVMScriptMap{})
	pulumi.RegisterOutputType(FortimanagerDVMScriptOutput{})
	pulumi.RegisterOutputType(FortimanagerDVMScriptArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerDVMScriptMapOutput{})
}

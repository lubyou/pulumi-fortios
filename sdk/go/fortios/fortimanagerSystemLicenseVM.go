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

type FortimanagerSystemLicenseVM struct {
	pulumi.CustomResourceState

	Adom        pulumi.StringPtrOutput `pulumi:"adom"`
	FileContent pulumi.StringOutput    `pulumi:"fileContent"`
	Target      pulumi.StringOutput    `pulumi:"target"`
}

// NewFortimanagerSystemLicenseVM registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemLicenseVM(ctx *pulumi.Context,
	name string, args *FortimanagerSystemLicenseVMArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemLicenseVM, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileContent == nil {
		return nil, errors.New("invalid value for required argument 'FileContent'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FortimanagerSystemLicenseVM
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemLicenseVM:FortimanagerSystemLicenseVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemLicenseVM gets an existing FortimanagerSystemLicenseVM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemLicenseVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemLicenseVMState, opts ...pulumi.ResourceOption) (*FortimanagerSystemLicenseVM, error) {
	var resource FortimanagerSystemLicenseVM
	err := ctx.ReadResource("fortios:index/fortimanagerSystemLicenseVM:FortimanagerSystemLicenseVM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemLicenseVM resources.
type fortimanagerSystemLicenseVMState struct {
	Adom        *string `pulumi:"adom"`
	FileContent *string `pulumi:"fileContent"`
	Target      *string `pulumi:"target"`
}

type FortimanagerSystemLicenseVMState struct {
	Adom        pulumi.StringPtrInput
	FileContent pulumi.StringPtrInput
	Target      pulumi.StringPtrInput
}

func (FortimanagerSystemLicenseVMState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemLicenseVMState)(nil)).Elem()
}

type fortimanagerSystemLicenseVMArgs struct {
	Adom        *string `pulumi:"adom"`
	FileContent string  `pulumi:"fileContent"`
	Target      string  `pulumi:"target"`
}

// The set of arguments for constructing a FortimanagerSystemLicenseVM resource.
type FortimanagerSystemLicenseVMArgs struct {
	Adom        pulumi.StringPtrInput
	FileContent pulumi.StringInput
	Target      pulumi.StringInput
}

func (FortimanagerSystemLicenseVMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemLicenseVMArgs)(nil)).Elem()
}

type FortimanagerSystemLicenseVMInput interface {
	pulumi.Input

	ToFortimanagerSystemLicenseVMOutput() FortimanagerSystemLicenseVMOutput
	ToFortimanagerSystemLicenseVMOutputWithContext(ctx context.Context) FortimanagerSystemLicenseVMOutput
}

func (*FortimanagerSystemLicenseVM) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemLicenseVM)(nil)).Elem()
}

func (i *FortimanagerSystemLicenseVM) ToFortimanagerSystemLicenseVMOutput() FortimanagerSystemLicenseVMOutput {
	return i.ToFortimanagerSystemLicenseVMOutputWithContext(context.Background())
}

func (i *FortimanagerSystemLicenseVM) ToFortimanagerSystemLicenseVMOutputWithContext(ctx context.Context) FortimanagerSystemLicenseVMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemLicenseVMOutput)
}

// FortimanagerSystemLicenseVMArrayInput is an input type that accepts FortimanagerSystemLicenseVMArray and FortimanagerSystemLicenseVMArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemLicenseVMArrayInput` via:
//
//	FortimanagerSystemLicenseVMArray{ FortimanagerSystemLicenseVMArgs{...} }
type FortimanagerSystemLicenseVMArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemLicenseVMArrayOutput() FortimanagerSystemLicenseVMArrayOutput
	ToFortimanagerSystemLicenseVMArrayOutputWithContext(context.Context) FortimanagerSystemLicenseVMArrayOutput
}

type FortimanagerSystemLicenseVMArray []FortimanagerSystemLicenseVMInput

func (FortimanagerSystemLicenseVMArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemLicenseVM)(nil)).Elem()
}

func (i FortimanagerSystemLicenseVMArray) ToFortimanagerSystemLicenseVMArrayOutput() FortimanagerSystemLicenseVMArrayOutput {
	return i.ToFortimanagerSystemLicenseVMArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemLicenseVMArray) ToFortimanagerSystemLicenseVMArrayOutputWithContext(ctx context.Context) FortimanagerSystemLicenseVMArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemLicenseVMArrayOutput)
}

// FortimanagerSystemLicenseVMMapInput is an input type that accepts FortimanagerSystemLicenseVMMap and FortimanagerSystemLicenseVMMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemLicenseVMMapInput` via:
//
//	FortimanagerSystemLicenseVMMap{ "key": FortimanagerSystemLicenseVMArgs{...} }
type FortimanagerSystemLicenseVMMapInput interface {
	pulumi.Input

	ToFortimanagerSystemLicenseVMMapOutput() FortimanagerSystemLicenseVMMapOutput
	ToFortimanagerSystemLicenseVMMapOutputWithContext(context.Context) FortimanagerSystemLicenseVMMapOutput
}

type FortimanagerSystemLicenseVMMap map[string]FortimanagerSystemLicenseVMInput

func (FortimanagerSystemLicenseVMMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemLicenseVM)(nil)).Elem()
}

func (i FortimanagerSystemLicenseVMMap) ToFortimanagerSystemLicenseVMMapOutput() FortimanagerSystemLicenseVMMapOutput {
	return i.ToFortimanagerSystemLicenseVMMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemLicenseVMMap) ToFortimanagerSystemLicenseVMMapOutputWithContext(ctx context.Context) FortimanagerSystemLicenseVMMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemLicenseVMMapOutput)
}

type FortimanagerSystemLicenseVMOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemLicenseVMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemLicenseVM)(nil)).Elem()
}

func (o FortimanagerSystemLicenseVMOutput) ToFortimanagerSystemLicenseVMOutput() FortimanagerSystemLicenseVMOutput {
	return o
}

func (o FortimanagerSystemLicenseVMOutput) ToFortimanagerSystemLicenseVMOutputWithContext(ctx context.Context) FortimanagerSystemLicenseVMOutput {
	return o
}

func (o FortimanagerSystemLicenseVMOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemLicenseVM) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemLicenseVMOutput) FileContent() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerSystemLicenseVM) pulumi.StringOutput { return v.FileContent }).(pulumi.StringOutput)
}

func (o FortimanagerSystemLicenseVMOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerSystemLicenseVM) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

type FortimanagerSystemLicenseVMArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemLicenseVMArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemLicenseVM)(nil)).Elem()
}

func (o FortimanagerSystemLicenseVMArrayOutput) ToFortimanagerSystemLicenseVMArrayOutput() FortimanagerSystemLicenseVMArrayOutput {
	return o
}

func (o FortimanagerSystemLicenseVMArrayOutput) ToFortimanagerSystemLicenseVMArrayOutputWithContext(ctx context.Context) FortimanagerSystemLicenseVMArrayOutput {
	return o
}

func (o FortimanagerSystemLicenseVMArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemLicenseVMOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerSystemLicenseVM {
		return vs[0].([]*FortimanagerSystemLicenseVM)[vs[1].(int)]
	}).(FortimanagerSystemLicenseVMOutput)
}

type FortimanagerSystemLicenseVMMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemLicenseVMMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemLicenseVM)(nil)).Elem()
}

func (o FortimanagerSystemLicenseVMMapOutput) ToFortimanagerSystemLicenseVMMapOutput() FortimanagerSystemLicenseVMMapOutput {
	return o
}

func (o FortimanagerSystemLicenseVMMapOutput) ToFortimanagerSystemLicenseVMMapOutputWithContext(ctx context.Context) FortimanagerSystemLicenseVMMapOutput {
	return o
}

func (o FortimanagerSystemLicenseVMMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemLicenseVMOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerSystemLicenseVM {
		return vs[0].(map[string]*FortimanagerSystemLicenseVM)[vs[1].(string)]
	}).(FortimanagerSystemLicenseVMOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemLicenseVMInput)(nil)).Elem(), &FortimanagerSystemLicenseVM{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemLicenseVMArrayInput)(nil)).Elem(), FortimanagerSystemLicenseVMArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemLicenseVMMapInput)(nil)).Elem(), FortimanagerSystemLicenseVMMap{})
	pulumi.RegisterOutputType(FortimanagerSystemLicenseVMOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemLicenseVMArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemLicenseVMMapOutput{})
}

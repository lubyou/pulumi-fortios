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

type FortimanagerSystemLicenseFortiCare struct {
	pulumi.CustomResourceState

	Adom             pulumi.StringPtrOutput `pulumi:"adom"`
	RegistrationCode pulumi.StringOutput    `pulumi:"registrationCode"`
	Target           pulumi.StringOutput    `pulumi:"target"`
}

// NewFortimanagerSystemLicenseFortiCare registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemLicenseFortiCare(ctx *pulumi.Context,
	name string, args *FortimanagerSystemLicenseFortiCareArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemLicenseFortiCare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistrationCode == nil {
		return nil, errors.New("invalid value for required argument 'RegistrationCode'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FortimanagerSystemLicenseFortiCare
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemLicenseFortiCare:FortimanagerSystemLicenseFortiCare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemLicenseFortiCare gets an existing FortimanagerSystemLicenseFortiCare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemLicenseFortiCare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemLicenseFortiCareState, opts ...pulumi.ResourceOption) (*FortimanagerSystemLicenseFortiCare, error) {
	var resource FortimanagerSystemLicenseFortiCare
	err := ctx.ReadResource("fortios:index/fortimanagerSystemLicenseFortiCare:FortimanagerSystemLicenseFortiCare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemLicenseFortiCare resources.
type fortimanagerSystemLicenseFortiCareState struct {
	Adom             *string `pulumi:"adom"`
	RegistrationCode *string `pulumi:"registrationCode"`
	Target           *string `pulumi:"target"`
}

type FortimanagerSystemLicenseFortiCareState struct {
	Adom             pulumi.StringPtrInput
	RegistrationCode pulumi.StringPtrInput
	Target           pulumi.StringPtrInput
}

func (FortimanagerSystemLicenseFortiCareState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemLicenseFortiCareState)(nil)).Elem()
}

type fortimanagerSystemLicenseFortiCareArgs struct {
	Adom             *string `pulumi:"adom"`
	RegistrationCode string  `pulumi:"registrationCode"`
	Target           string  `pulumi:"target"`
}

// The set of arguments for constructing a FortimanagerSystemLicenseFortiCare resource.
type FortimanagerSystemLicenseFortiCareArgs struct {
	Adom             pulumi.StringPtrInput
	RegistrationCode pulumi.StringInput
	Target           pulumi.StringInput
}

func (FortimanagerSystemLicenseFortiCareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemLicenseFortiCareArgs)(nil)).Elem()
}

type FortimanagerSystemLicenseFortiCareInput interface {
	pulumi.Input

	ToFortimanagerSystemLicenseFortiCareOutput() FortimanagerSystemLicenseFortiCareOutput
	ToFortimanagerSystemLicenseFortiCareOutputWithContext(ctx context.Context) FortimanagerSystemLicenseFortiCareOutput
}

func (*FortimanagerSystemLicenseFortiCare) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemLicenseFortiCare)(nil)).Elem()
}

func (i *FortimanagerSystemLicenseFortiCare) ToFortimanagerSystemLicenseFortiCareOutput() FortimanagerSystemLicenseFortiCareOutput {
	return i.ToFortimanagerSystemLicenseFortiCareOutputWithContext(context.Background())
}

func (i *FortimanagerSystemLicenseFortiCare) ToFortimanagerSystemLicenseFortiCareOutputWithContext(ctx context.Context) FortimanagerSystemLicenseFortiCareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemLicenseFortiCareOutput)
}

// FortimanagerSystemLicenseFortiCareArrayInput is an input type that accepts FortimanagerSystemLicenseFortiCareArray and FortimanagerSystemLicenseFortiCareArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemLicenseFortiCareArrayInput` via:
//
//	FortimanagerSystemLicenseFortiCareArray{ FortimanagerSystemLicenseFortiCareArgs{...} }
type FortimanagerSystemLicenseFortiCareArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemLicenseFortiCareArrayOutput() FortimanagerSystemLicenseFortiCareArrayOutput
	ToFortimanagerSystemLicenseFortiCareArrayOutputWithContext(context.Context) FortimanagerSystemLicenseFortiCareArrayOutput
}

type FortimanagerSystemLicenseFortiCareArray []FortimanagerSystemLicenseFortiCareInput

func (FortimanagerSystemLicenseFortiCareArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemLicenseFortiCare)(nil)).Elem()
}

func (i FortimanagerSystemLicenseFortiCareArray) ToFortimanagerSystemLicenseFortiCareArrayOutput() FortimanagerSystemLicenseFortiCareArrayOutput {
	return i.ToFortimanagerSystemLicenseFortiCareArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemLicenseFortiCareArray) ToFortimanagerSystemLicenseFortiCareArrayOutputWithContext(ctx context.Context) FortimanagerSystemLicenseFortiCareArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemLicenseFortiCareArrayOutput)
}

// FortimanagerSystemLicenseFortiCareMapInput is an input type that accepts FortimanagerSystemLicenseFortiCareMap and FortimanagerSystemLicenseFortiCareMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemLicenseFortiCareMapInput` via:
//
//	FortimanagerSystemLicenseFortiCareMap{ "key": FortimanagerSystemLicenseFortiCareArgs{...} }
type FortimanagerSystemLicenseFortiCareMapInput interface {
	pulumi.Input

	ToFortimanagerSystemLicenseFortiCareMapOutput() FortimanagerSystemLicenseFortiCareMapOutput
	ToFortimanagerSystemLicenseFortiCareMapOutputWithContext(context.Context) FortimanagerSystemLicenseFortiCareMapOutput
}

type FortimanagerSystemLicenseFortiCareMap map[string]FortimanagerSystemLicenseFortiCareInput

func (FortimanagerSystemLicenseFortiCareMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemLicenseFortiCare)(nil)).Elem()
}

func (i FortimanagerSystemLicenseFortiCareMap) ToFortimanagerSystemLicenseFortiCareMapOutput() FortimanagerSystemLicenseFortiCareMapOutput {
	return i.ToFortimanagerSystemLicenseFortiCareMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemLicenseFortiCareMap) ToFortimanagerSystemLicenseFortiCareMapOutputWithContext(ctx context.Context) FortimanagerSystemLicenseFortiCareMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemLicenseFortiCareMapOutput)
}

type FortimanagerSystemLicenseFortiCareOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemLicenseFortiCareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemLicenseFortiCare)(nil)).Elem()
}

func (o FortimanagerSystemLicenseFortiCareOutput) ToFortimanagerSystemLicenseFortiCareOutput() FortimanagerSystemLicenseFortiCareOutput {
	return o
}

func (o FortimanagerSystemLicenseFortiCareOutput) ToFortimanagerSystemLicenseFortiCareOutputWithContext(ctx context.Context) FortimanagerSystemLicenseFortiCareOutput {
	return o
}

func (o FortimanagerSystemLicenseFortiCareOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemLicenseFortiCare) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemLicenseFortiCareOutput) RegistrationCode() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerSystemLicenseFortiCare) pulumi.StringOutput { return v.RegistrationCode }).(pulumi.StringOutput)
}

func (o FortimanagerSystemLicenseFortiCareOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerSystemLicenseFortiCare) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

type FortimanagerSystemLicenseFortiCareArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemLicenseFortiCareArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemLicenseFortiCare)(nil)).Elem()
}

func (o FortimanagerSystemLicenseFortiCareArrayOutput) ToFortimanagerSystemLicenseFortiCareArrayOutput() FortimanagerSystemLicenseFortiCareArrayOutput {
	return o
}

func (o FortimanagerSystemLicenseFortiCareArrayOutput) ToFortimanagerSystemLicenseFortiCareArrayOutputWithContext(ctx context.Context) FortimanagerSystemLicenseFortiCareArrayOutput {
	return o
}

func (o FortimanagerSystemLicenseFortiCareArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemLicenseFortiCareOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerSystemLicenseFortiCare {
		return vs[0].([]*FortimanagerSystemLicenseFortiCare)[vs[1].(int)]
	}).(FortimanagerSystemLicenseFortiCareOutput)
}

type FortimanagerSystemLicenseFortiCareMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemLicenseFortiCareMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemLicenseFortiCare)(nil)).Elem()
}

func (o FortimanagerSystemLicenseFortiCareMapOutput) ToFortimanagerSystemLicenseFortiCareMapOutput() FortimanagerSystemLicenseFortiCareMapOutput {
	return o
}

func (o FortimanagerSystemLicenseFortiCareMapOutput) ToFortimanagerSystemLicenseFortiCareMapOutputWithContext(ctx context.Context) FortimanagerSystemLicenseFortiCareMapOutput {
	return o
}

func (o FortimanagerSystemLicenseFortiCareMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemLicenseFortiCareOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerSystemLicenseFortiCare {
		return vs[0].(map[string]*FortimanagerSystemLicenseFortiCare)[vs[1].(string)]
	}).(FortimanagerSystemLicenseFortiCareOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemLicenseFortiCareInput)(nil)).Elem(), &FortimanagerSystemLicenseFortiCare{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemLicenseFortiCareArrayInput)(nil)).Elem(), FortimanagerSystemLicenseFortiCareArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemLicenseFortiCareMapInput)(nil)).Elem(), FortimanagerSystemLicenseFortiCareMap{})
	pulumi.RegisterOutputType(FortimanagerSystemLicenseFortiCareOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemLicenseFortiCareArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemLicenseFortiCareMapOutput{})
}

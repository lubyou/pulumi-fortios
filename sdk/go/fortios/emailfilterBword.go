// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EmailfilterBword struct {
	pulumi.CustomResourceState

	Comment             pulumi.StringPtrOutput           `pulumi:"comment"`
	DynamicSortSubtable pulumi.StringPtrOutput           `pulumi:"dynamicSortSubtable"`
	Entries             EmailfilterBwordEntryArrayOutput `pulumi:"entries"`
	Fosid               pulumi.IntOutput                 `pulumi:"fosid"`
	GetAllTables        pulumi.StringPtrOutput           `pulumi:"getAllTables"`
	Name                pulumi.StringOutput              `pulumi:"name"`
	Vdomparam           pulumi.StringPtrOutput           `pulumi:"vdomparam"`
}

// NewEmailfilterBword registers a new resource with the given unique name, arguments, and options.
func NewEmailfilterBword(ctx *pulumi.Context,
	name string, args *EmailfilterBwordArgs, opts ...pulumi.ResourceOption) (*EmailfilterBword, error) {
	if args == nil {
		args = &EmailfilterBwordArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EmailfilterBword
	err := ctx.RegisterResource("fortios:index/emailfilterBword:EmailfilterBword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailfilterBword gets an existing EmailfilterBword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailfilterBword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailfilterBwordState, opts ...pulumi.ResourceOption) (*EmailfilterBword, error) {
	var resource EmailfilterBword
	err := ctx.ReadResource("fortios:index/emailfilterBword:EmailfilterBword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailfilterBword resources.
type emailfilterBwordState struct {
	Comment             *string                 `pulumi:"comment"`
	DynamicSortSubtable *string                 `pulumi:"dynamicSortSubtable"`
	Entries             []EmailfilterBwordEntry `pulumi:"entries"`
	Fosid               *int                    `pulumi:"fosid"`
	GetAllTables        *string                 `pulumi:"getAllTables"`
	Name                *string                 `pulumi:"name"`
	Vdomparam           *string                 `pulumi:"vdomparam"`
}

type EmailfilterBwordState struct {
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Entries             EmailfilterBwordEntryArrayInput
	Fosid               pulumi.IntPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (EmailfilterBwordState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterBwordState)(nil)).Elem()
}

type emailfilterBwordArgs struct {
	Comment             *string                 `pulumi:"comment"`
	DynamicSortSubtable *string                 `pulumi:"dynamicSortSubtable"`
	Entries             []EmailfilterBwordEntry `pulumi:"entries"`
	Fosid               *int                    `pulumi:"fosid"`
	GetAllTables        *string                 `pulumi:"getAllTables"`
	Name                *string                 `pulumi:"name"`
	Vdomparam           *string                 `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EmailfilterBword resource.
type EmailfilterBwordArgs struct {
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Entries             EmailfilterBwordEntryArrayInput
	Fosid               pulumi.IntPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (EmailfilterBwordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterBwordArgs)(nil)).Elem()
}

type EmailfilterBwordInput interface {
	pulumi.Input

	ToEmailfilterBwordOutput() EmailfilterBwordOutput
	ToEmailfilterBwordOutputWithContext(ctx context.Context) EmailfilterBwordOutput
}

func (*EmailfilterBword) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterBword)(nil)).Elem()
}

func (i *EmailfilterBword) ToEmailfilterBwordOutput() EmailfilterBwordOutput {
	return i.ToEmailfilterBwordOutputWithContext(context.Background())
}

func (i *EmailfilterBword) ToEmailfilterBwordOutputWithContext(ctx context.Context) EmailfilterBwordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterBwordOutput)
}

// EmailfilterBwordArrayInput is an input type that accepts EmailfilterBwordArray and EmailfilterBwordArrayOutput values.
// You can construct a concrete instance of `EmailfilterBwordArrayInput` via:
//
//	EmailfilterBwordArray{ EmailfilterBwordArgs{...} }
type EmailfilterBwordArrayInput interface {
	pulumi.Input

	ToEmailfilterBwordArrayOutput() EmailfilterBwordArrayOutput
	ToEmailfilterBwordArrayOutputWithContext(context.Context) EmailfilterBwordArrayOutput
}

type EmailfilterBwordArray []EmailfilterBwordInput

func (EmailfilterBwordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterBword)(nil)).Elem()
}

func (i EmailfilterBwordArray) ToEmailfilterBwordArrayOutput() EmailfilterBwordArrayOutput {
	return i.ToEmailfilterBwordArrayOutputWithContext(context.Background())
}

func (i EmailfilterBwordArray) ToEmailfilterBwordArrayOutputWithContext(ctx context.Context) EmailfilterBwordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterBwordArrayOutput)
}

// EmailfilterBwordMapInput is an input type that accepts EmailfilterBwordMap and EmailfilterBwordMapOutput values.
// You can construct a concrete instance of `EmailfilterBwordMapInput` via:
//
//	EmailfilterBwordMap{ "key": EmailfilterBwordArgs{...} }
type EmailfilterBwordMapInput interface {
	pulumi.Input

	ToEmailfilterBwordMapOutput() EmailfilterBwordMapOutput
	ToEmailfilterBwordMapOutputWithContext(context.Context) EmailfilterBwordMapOutput
}

type EmailfilterBwordMap map[string]EmailfilterBwordInput

func (EmailfilterBwordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterBword)(nil)).Elem()
}

func (i EmailfilterBwordMap) ToEmailfilterBwordMapOutput() EmailfilterBwordMapOutput {
	return i.ToEmailfilterBwordMapOutputWithContext(context.Background())
}

func (i EmailfilterBwordMap) ToEmailfilterBwordMapOutputWithContext(ctx context.Context) EmailfilterBwordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterBwordMapOutput)
}

type EmailfilterBwordOutput struct{ *pulumi.OutputState }

func (EmailfilterBwordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterBword)(nil)).Elem()
}

func (o EmailfilterBwordOutput) ToEmailfilterBwordOutput() EmailfilterBwordOutput {
	return o
}

func (o EmailfilterBwordOutput) ToEmailfilterBwordOutputWithContext(ctx context.Context) EmailfilterBwordOutput {
	return o
}

func (o EmailfilterBwordOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailfilterBword) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o EmailfilterBwordOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailfilterBword) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o EmailfilterBwordOutput) Entries() EmailfilterBwordEntryArrayOutput {
	return o.ApplyT(func(v *EmailfilterBword) EmailfilterBwordEntryArrayOutput { return v.Entries }).(EmailfilterBwordEntryArrayOutput)
}

func (o EmailfilterBwordOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *EmailfilterBword) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o EmailfilterBwordOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailfilterBword) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o EmailfilterBwordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailfilterBword) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EmailfilterBwordOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailfilterBword) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type EmailfilterBwordArrayOutput struct{ *pulumi.OutputState }

func (EmailfilterBwordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterBword)(nil)).Elem()
}

func (o EmailfilterBwordArrayOutput) ToEmailfilterBwordArrayOutput() EmailfilterBwordArrayOutput {
	return o
}

func (o EmailfilterBwordArrayOutput) ToEmailfilterBwordArrayOutputWithContext(ctx context.Context) EmailfilterBwordArrayOutput {
	return o
}

func (o EmailfilterBwordArrayOutput) Index(i pulumi.IntInput) EmailfilterBwordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmailfilterBword {
		return vs[0].([]*EmailfilterBword)[vs[1].(int)]
	}).(EmailfilterBwordOutput)
}

type EmailfilterBwordMapOutput struct{ *pulumi.OutputState }

func (EmailfilterBwordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterBword)(nil)).Elem()
}

func (o EmailfilterBwordMapOutput) ToEmailfilterBwordMapOutput() EmailfilterBwordMapOutput {
	return o
}

func (o EmailfilterBwordMapOutput) ToEmailfilterBwordMapOutputWithContext(ctx context.Context) EmailfilterBwordMapOutput {
	return o
}

func (o EmailfilterBwordMapOutput) MapIndex(k pulumi.StringInput) EmailfilterBwordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmailfilterBword {
		return vs[0].(map[string]*EmailfilterBword)[vs[1].(string)]
	}).(EmailfilterBwordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterBwordInput)(nil)).Elem(), &EmailfilterBword{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterBwordArrayInput)(nil)).Elem(), EmailfilterBwordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterBwordMapInput)(nil)).Elem(), EmailfilterBwordMap{})
	pulumi.RegisterOutputType(EmailfilterBwordOutput{})
	pulumi.RegisterOutputType(EmailfilterBwordArrayOutput{})
	pulumi.RegisterOutputType(EmailfilterBwordMapOutput{})
}

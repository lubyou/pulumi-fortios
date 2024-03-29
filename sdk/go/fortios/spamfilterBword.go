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

type SpamfilterBword struct {
	pulumi.CustomResourceState

	Comment             pulumi.StringPtrOutput          `pulumi:"comment"`
	DynamicSortSubtable pulumi.StringPtrOutput          `pulumi:"dynamicSortSubtable"`
	Entries             SpamfilterBwordEntryArrayOutput `pulumi:"entries"`
	Fosid               pulumi.IntOutput                `pulumi:"fosid"`
	GetAllTables        pulumi.StringPtrOutput          `pulumi:"getAllTables"`
	Name                pulumi.StringOutput             `pulumi:"name"`
	Vdomparam           pulumi.StringPtrOutput          `pulumi:"vdomparam"`
}

// NewSpamfilterBword registers a new resource with the given unique name, arguments, and options.
func NewSpamfilterBword(ctx *pulumi.Context,
	name string, args *SpamfilterBwordArgs, opts ...pulumi.ResourceOption) (*SpamfilterBword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SpamfilterBword
	err := ctx.RegisterResource("fortios:index/spamfilterBword:SpamfilterBword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpamfilterBword gets an existing SpamfilterBword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpamfilterBword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpamfilterBwordState, opts ...pulumi.ResourceOption) (*SpamfilterBword, error) {
	var resource SpamfilterBword
	err := ctx.ReadResource("fortios:index/spamfilterBword:SpamfilterBword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpamfilterBword resources.
type spamfilterBwordState struct {
	Comment             *string                `pulumi:"comment"`
	DynamicSortSubtable *string                `pulumi:"dynamicSortSubtable"`
	Entries             []SpamfilterBwordEntry `pulumi:"entries"`
	Fosid               *int                   `pulumi:"fosid"`
	GetAllTables        *string                `pulumi:"getAllTables"`
	Name                *string                `pulumi:"name"`
	Vdomparam           *string                `pulumi:"vdomparam"`
}

type SpamfilterBwordState struct {
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Entries             SpamfilterBwordEntryArrayInput
	Fosid               pulumi.IntPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SpamfilterBwordState) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterBwordState)(nil)).Elem()
}

type spamfilterBwordArgs struct {
	Comment             *string                `pulumi:"comment"`
	DynamicSortSubtable *string                `pulumi:"dynamicSortSubtable"`
	Entries             []SpamfilterBwordEntry `pulumi:"entries"`
	Fosid               int                    `pulumi:"fosid"`
	GetAllTables        *string                `pulumi:"getAllTables"`
	Name                *string                `pulumi:"name"`
	Vdomparam           *string                `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SpamfilterBword resource.
type SpamfilterBwordArgs struct {
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Entries             SpamfilterBwordEntryArrayInput
	Fosid               pulumi.IntInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SpamfilterBwordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterBwordArgs)(nil)).Elem()
}

type SpamfilterBwordInput interface {
	pulumi.Input

	ToSpamfilterBwordOutput() SpamfilterBwordOutput
	ToSpamfilterBwordOutputWithContext(ctx context.Context) SpamfilterBwordOutput
}

func (*SpamfilterBword) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterBword)(nil)).Elem()
}

func (i *SpamfilterBword) ToSpamfilterBwordOutput() SpamfilterBwordOutput {
	return i.ToSpamfilterBwordOutputWithContext(context.Background())
}

func (i *SpamfilterBword) ToSpamfilterBwordOutputWithContext(ctx context.Context) SpamfilterBwordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterBwordOutput)
}

// SpamfilterBwordArrayInput is an input type that accepts SpamfilterBwordArray and SpamfilterBwordArrayOutput values.
// You can construct a concrete instance of `SpamfilterBwordArrayInput` via:
//
//	SpamfilterBwordArray{ SpamfilterBwordArgs{...} }
type SpamfilterBwordArrayInput interface {
	pulumi.Input

	ToSpamfilterBwordArrayOutput() SpamfilterBwordArrayOutput
	ToSpamfilterBwordArrayOutputWithContext(context.Context) SpamfilterBwordArrayOutput
}

type SpamfilterBwordArray []SpamfilterBwordInput

func (SpamfilterBwordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterBword)(nil)).Elem()
}

func (i SpamfilterBwordArray) ToSpamfilterBwordArrayOutput() SpamfilterBwordArrayOutput {
	return i.ToSpamfilterBwordArrayOutputWithContext(context.Background())
}

func (i SpamfilterBwordArray) ToSpamfilterBwordArrayOutputWithContext(ctx context.Context) SpamfilterBwordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterBwordArrayOutput)
}

// SpamfilterBwordMapInput is an input type that accepts SpamfilterBwordMap and SpamfilterBwordMapOutput values.
// You can construct a concrete instance of `SpamfilterBwordMapInput` via:
//
//	SpamfilterBwordMap{ "key": SpamfilterBwordArgs{...} }
type SpamfilterBwordMapInput interface {
	pulumi.Input

	ToSpamfilterBwordMapOutput() SpamfilterBwordMapOutput
	ToSpamfilterBwordMapOutputWithContext(context.Context) SpamfilterBwordMapOutput
}

type SpamfilterBwordMap map[string]SpamfilterBwordInput

func (SpamfilterBwordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterBword)(nil)).Elem()
}

func (i SpamfilterBwordMap) ToSpamfilterBwordMapOutput() SpamfilterBwordMapOutput {
	return i.ToSpamfilterBwordMapOutputWithContext(context.Background())
}

func (i SpamfilterBwordMap) ToSpamfilterBwordMapOutputWithContext(ctx context.Context) SpamfilterBwordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterBwordMapOutput)
}

type SpamfilterBwordOutput struct{ *pulumi.OutputState }

func (SpamfilterBwordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterBword)(nil)).Elem()
}

func (o SpamfilterBwordOutput) ToSpamfilterBwordOutput() SpamfilterBwordOutput {
	return o
}

func (o SpamfilterBwordOutput) ToSpamfilterBwordOutputWithContext(ctx context.Context) SpamfilterBwordOutput {
	return o
}

func (o SpamfilterBwordOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpamfilterBword) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o SpamfilterBwordOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpamfilterBword) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SpamfilterBwordOutput) Entries() SpamfilterBwordEntryArrayOutput {
	return o.ApplyT(func(v *SpamfilterBword) SpamfilterBwordEntryArrayOutput { return v.Entries }).(SpamfilterBwordEntryArrayOutput)
}

func (o SpamfilterBwordOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *SpamfilterBword) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o SpamfilterBwordOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpamfilterBword) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SpamfilterBwordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SpamfilterBword) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SpamfilterBwordOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpamfilterBword) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SpamfilterBwordArrayOutput struct{ *pulumi.OutputState }

func (SpamfilterBwordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterBword)(nil)).Elem()
}

func (o SpamfilterBwordArrayOutput) ToSpamfilterBwordArrayOutput() SpamfilterBwordArrayOutput {
	return o
}

func (o SpamfilterBwordArrayOutput) ToSpamfilterBwordArrayOutputWithContext(ctx context.Context) SpamfilterBwordArrayOutput {
	return o
}

func (o SpamfilterBwordArrayOutput) Index(i pulumi.IntInput) SpamfilterBwordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SpamfilterBword {
		return vs[0].([]*SpamfilterBword)[vs[1].(int)]
	}).(SpamfilterBwordOutput)
}

type SpamfilterBwordMapOutput struct{ *pulumi.OutputState }

func (SpamfilterBwordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterBword)(nil)).Elem()
}

func (o SpamfilterBwordMapOutput) ToSpamfilterBwordMapOutput() SpamfilterBwordMapOutput {
	return o
}

func (o SpamfilterBwordMapOutput) ToSpamfilterBwordMapOutputWithContext(ctx context.Context) SpamfilterBwordMapOutput {
	return o
}

func (o SpamfilterBwordMapOutput) MapIndex(k pulumi.StringInput) SpamfilterBwordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SpamfilterBword {
		return vs[0].(map[string]*SpamfilterBword)[vs[1].(string)]
	}).(SpamfilterBwordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterBwordInput)(nil)).Elem(), &SpamfilterBword{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterBwordArrayInput)(nil)).Elem(), SpamfilterBwordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterBwordMapInput)(nil)).Elem(), SpamfilterBwordMap{})
	pulumi.RegisterOutputType(SpamfilterBwordOutput{})
	pulumi.RegisterOutputType(SpamfilterBwordArrayOutput{})
	pulumi.RegisterOutputType(SpamfilterBwordMapOutput{})
}

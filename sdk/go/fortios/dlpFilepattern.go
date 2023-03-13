// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DlpFilepattern struct {
	pulumi.CustomResourceState

	Comment             pulumi.StringPtrOutput         `pulumi:"comment"`
	DynamicSortSubtable pulumi.StringPtrOutput         `pulumi:"dynamicSortSubtable"`
	Entries             DlpFilepatternEntryArrayOutput `pulumi:"entries"`
	Fosid               pulumi.IntOutput               `pulumi:"fosid"`
	Name                pulumi.StringOutput            `pulumi:"name"`
	Vdomparam           pulumi.StringPtrOutput         `pulumi:"vdomparam"`
}

// NewDlpFilepattern registers a new resource with the given unique name, arguments, and options.
func NewDlpFilepattern(ctx *pulumi.Context,
	name string, args *DlpFilepatternArgs, opts ...pulumi.ResourceOption) (*DlpFilepattern, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DlpFilepattern
	err := ctx.RegisterResource("fortios:index/dlpFilepattern:DlpFilepattern", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDlpFilepattern gets an existing DlpFilepattern resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDlpFilepattern(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DlpFilepatternState, opts ...pulumi.ResourceOption) (*DlpFilepattern, error) {
	var resource DlpFilepattern
	err := ctx.ReadResource("fortios:index/dlpFilepattern:DlpFilepattern", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DlpFilepattern resources.
type dlpFilepatternState struct {
	Comment             *string               `pulumi:"comment"`
	DynamicSortSubtable *string               `pulumi:"dynamicSortSubtable"`
	Entries             []DlpFilepatternEntry `pulumi:"entries"`
	Fosid               *int                  `pulumi:"fosid"`
	Name                *string               `pulumi:"name"`
	Vdomparam           *string               `pulumi:"vdomparam"`
}

type DlpFilepatternState struct {
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Entries             DlpFilepatternEntryArrayInput
	Fosid               pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (DlpFilepatternState) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpFilepatternState)(nil)).Elem()
}

type dlpFilepatternArgs struct {
	Comment             *string               `pulumi:"comment"`
	DynamicSortSubtable *string               `pulumi:"dynamicSortSubtable"`
	Entries             []DlpFilepatternEntry `pulumi:"entries"`
	Fosid               int                   `pulumi:"fosid"`
	Name                *string               `pulumi:"name"`
	Vdomparam           *string               `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DlpFilepattern resource.
type DlpFilepatternArgs struct {
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Entries             DlpFilepatternEntryArrayInput
	Fosid               pulumi.IntInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (DlpFilepatternArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpFilepatternArgs)(nil)).Elem()
}

type DlpFilepatternInput interface {
	pulumi.Input

	ToDlpFilepatternOutput() DlpFilepatternOutput
	ToDlpFilepatternOutputWithContext(ctx context.Context) DlpFilepatternOutput
}

func (*DlpFilepattern) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpFilepattern)(nil)).Elem()
}

func (i *DlpFilepattern) ToDlpFilepatternOutput() DlpFilepatternOutput {
	return i.ToDlpFilepatternOutputWithContext(context.Background())
}

func (i *DlpFilepattern) ToDlpFilepatternOutputWithContext(ctx context.Context) DlpFilepatternOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFilepatternOutput)
}

// DlpFilepatternArrayInput is an input type that accepts DlpFilepatternArray and DlpFilepatternArrayOutput values.
// You can construct a concrete instance of `DlpFilepatternArrayInput` via:
//
//	DlpFilepatternArray{ DlpFilepatternArgs{...} }
type DlpFilepatternArrayInput interface {
	pulumi.Input

	ToDlpFilepatternArrayOutput() DlpFilepatternArrayOutput
	ToDlpFilepatternArrayOutputWithContext(context.Context) DlpFilepatternArrayOutput
}

type DlpFilepatternArray []DlpFilepatternInput

func (DlpFilepatternArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpFilepattern)(nil)).Elem()
}

func (i DlpFilepatternArray) ToDlpFilepatternArrayOutput() DlpFilepatternArrayOutput {
	return i.ToDlpFilepatternArrayOutputWithContext(context.Background())
}

func (i DlpFilepatternArray) ToDlpFilepatternArrayOutputWithContext(ctx context.Context) DlpFilepatternArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFilepatternArrayOutput)
}

// DlpFilepatternMapInput is an input type that accepts DlpFilepatternMap and DlpFilepatternMapOutput values.
// You can construct a concrete instance of `DlpFilepatternMapInput` via:
//
//	DlpFilepatternMap{ "key": DlpFilepatternArgs{...} }
type DlpFilepatternMapInput interface {
	pulumi.Input

	ToDlpFilepatternMapOutput() DlpFilepatternMapOutput
	ToDlpFilepatternMapOutputWithContext(context.Context) DlpFilepatternMapOutput
}

type DlpFilepatternMap map[string]DlpFilepatternInput

func (DlpFilepatternMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpFilepattern)(nil)).Elem()
}

func (i DlpFilepatternMap) ToDlpFilepatternMapOutput() DlpFilepatternMapOutput {
	return i.ToDlpFilepatternMapOutputWithContext(context.Background())
}

func (i DlpFilepatternMap) ToDlpFilepatternMapOutputWithContext(ctx context.Context) DlpFilepatternMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFilepatternMapOutput)
}

type DlpFilepatternOutput struct{ *pulumi.OutputState }

func (DlpFilepatternOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpFilepattern)(nil)).Elem()
}

func (o DlpFilepatternOutput) ToDlpFilepatternOutput() DlpFilepatternOutput {
	return o
}

func (o DlpFilepatternOutput) ToDlpFilepatternOutputWithContext(ctx context.Context) DlpFilepatternOutput {
	return o
}

func (o DlpFilepatternOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpFilepattern) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o DlpFilepatternOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpFilepattern) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o DlpFilepatternOutput) Entries() DlpFilepatternEntryArrayOutput {
	return o.ApplyT(func(v *DlpFilepattern) DlpFilepatternEntryArrayOutput { return v.Entries }).(DlpFilepatternEntryArrayOutput)
}

func (o DlpFilepatternOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *DlpFilepattern) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o DlpFilepatternOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DlpFilepattern) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DlpFilepatternOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DlpFilepattern) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type DlpFilepatternArrayOutput struct{ *pulumi.OutputState }

func (DlpFilepatternArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpFilepattern)(nil)).Elem()
}

func (o DlpFilepatternArrayOutput) ToDlpFilepatternArrayOutput() DlpFilepatternArrayOutput {
	return o
}

func (o DlpFilepatternArrayOutput) ToDlpFilepatternArrayOutputWithContext(ctx context.Context) DlpFilepatternArrayOutput {
	return o
}

func (o DlpFilepatternArrayOutput) Index(i pulumi.IntInput) DlpFilepatternOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DlpFilepattern {
		return vs[0].([]*DlpFilepattern)[vs[1].(int)]
	}).(DlpFilepatternOutput)
}

type DlpFilepatternMapOutput struct{ *pulumi.OutputState }

func (DlpFilepatternMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpFilepattern)(nil)).Elem()
}

func (o DlpFilepatternMapOutput) ToDlpFilepatternMapOutput() DlpFilepatternMapOutput {
	return o
}

func (o DlpFilepatternMapOutput) ToDlpFilepatternMapOutputWithContext(ctx context.Context) DlpFilepatternMapOutput {
	return o
}

func (o DlpFilepatternMapOutput) MapIndex(k pulumi.StringInput) DlpFilepatternOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DlpFilepattern {
		return vs[0].(map[string]*DlpFilepattern)[vs[1].(string)]
	}).(DlpFilepatternOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DlpFilepatternInput)(nil)).Elem(), &DlpFilepattern{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpFilepatternArrayInput)(nil)).Elem(), DlpFilepatternArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpFilepatternMapInput)(nil)).Elem(), DlpFilepatternMap{})
	pulumi.RegisterOutputType(DlpFilepatternOutput{})
	pulumi.RegisterOutputType(DlpFilepatternArrayOutput{})
	pulumi.RegisterOutputType(DlpFilepatternMapOutput{})
}

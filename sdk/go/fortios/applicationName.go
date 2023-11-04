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

type ApplicationName struct {
	pulumi.CustomResourceState

	Behavior            pulumi.StringOutput                 `pulumi:"behavior"`
	Category            pulumi.IntOutput                    `pulumi:"category"`
	DynamicSortSubtable pulumi.StringPtrOutput              `pulumi:"dynamicSortSubtable"`
	Fosid               pulumi.IntOutput                    `pulumi:"fosid"`
	GetAllTables        pulumi.StringPtrOutput              `pulumi:"getAllTables"`
	Metadatas           ApplicationNameMetadataArrayOutput  `pulumi:"metadatas"`
	Name                pulumi.StringOutput                 `pulumi:"name"`
	Parameter           pulumi.StringOutput                 `pulumi:"parameter"`
	Parameters          ApplicationNameParameterArrayOutput `pulumi:"parameters"`
	Popularity          pulumi.IntOutput                    `pulumi:"popularity"`
	Protocol            pulumi.StringOutput                 `pulumi:"protocol"`
	Risk                pulumi.IntOutput                    `pulumi:"risk"`
	SubCategory         pulumi.IntOutput                    `pulumi:"subCategory"`
	Technology          pulumi.StringOutput                 `pulumi:"technology"`
	Vdomparam           pulumi.StringPtrOutput              `pulumi:"vdomparam"`
	Vendor              pulumi.StringOutput                 `pulumi:"vendor"`
	Weight              pulumi.IntOutput                    `pulumi:"weight"`
}

// NewApplicationName registers a new resource with the given unique name, arguments, and options.
func NewApplicationName(ctx *pulumi.Context,
	name string, args *ApplicationNameArgs, opts ...pulumi.ResourceOption) (*ApplicationName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationName
	err := ctx.RegisterResource("fortios:index/applicationName:ApplicationName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationName gets an existing ApplicationName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationNameState, opts ...pulumi.ResourceOption) (*ApplicationName, error) {
	var resource ApplicationName
	err := ctx.ReadResource("fortios:index/applicationName:ApplicationName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationName resources.
type applicationNameState struct {
	Behavior            *string                    `pulumi:"behavior"`
	Category            *int                       `pulumi:"category"`
	DynamicSortSubtable *string                    `pulumi:"dynamicSortSubtable"`
	Fosid               *int                       `pulumi:"fosid"`
	GetAllTables        *string                    `pulumi:"getAllTables"`
	Metadatas           []ApplicationNameMetadata  `pulumi:"metadatas"`
	Name                *string                    `pulumi:"name"`
	Parameter           *string                    `pulumi:"parameter"`
	Parameters          []ApplicationNameParameter `pulumi:"parameters"`
	Popularity          *int                       `pulumi:"popularity"`
	Protocol            *string                    `pulumi:"protocol"`
	Risk                *int                       `pulumi:"risk"`
	SubCategory         *int                       `pulumi:"subCategory"`
	Technology          *string                    `pulumi:"technology"`
	Vdomparam           *string                    `pulumi:"vdomparam"`
	Vendor              *string                    `pulumi:"vendor"`
	Weight              *int                       `pulumi:"weight"`
}

type ApplicationNameState struct {
	Behavior            pulumi.StringPtrInput
	Category            pulumi.IntPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Fosid               pulumi.IntPtrInput
	GetAllTables        pulumi.StringPtrInput
	Metadatas           ApplicationNameMetadataArrayInput
	Name                pulumi.StringPtrInput
	Parameter           pulumi.StringPtrInput
	Parameters          ApplicationNameParameterArrayInput
	Popularity          pulumi.IntPtrInput
	Protocol            pulumi.StringPtrInput
	Risk                pulumi.IntPtrInput
	SubCategory         pulumi.IntPtrInput
	Technology          pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	Vendor              pulumi.StringPtrInput
	Weight              pulumi.IntPtrInput
}

func (ApplicationNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationNameState)(nil)).Elem()
}

type applicationNameArgs struct {
	Behavior            *string                    `pulumi:"behavior"`
	Category            int                        `pulumi:"category"`
	DynamicSortSubtable *string                    `pulumi:"dynamicSortSubtable"`
	Fosid               *int                       `pulumi:"fosid"`
	GetAllTables        *string                    `pulumi:"getAllTables"`
	Metadatas           []ApplicationNameMetadata  `pulumi:"metadatas"`
	Name                *string                    `pulumi:"name"`
	Parameter           *string                    `pulumi:"parameter"`
	Parameters          []ApplicationNameParameter `pulumi:"parameters"`
	Popularity          *int                       `pulumi:"popularity"`
	Protocol            *string                    `pulumi:"protocol"`
	Risk                *int                       `pulumi:"risk"`
	SubCategory         *int                       `pulumi:"subCategory"`
	Technology          *string                    `pulumi:"technology"`
	Vdomparam           *string                    `pulumi:"vdomparam"`
	Vendor              *string                    `pulumi:"vendor"`
	Weight              *int                       `pulumi:"weight"`
}

// The set of arguments for constructing a ApplicationName resource.
type ApplicationNameArgs struct {
	Behavior            pulumi.StringPtrInput
	Category            pulumi.IntInput
	DynamicSortSubtable pulumi.StringPtrInput
	Fosid               pulumi.IntPtrInput
	GetAllTables        pulumi.StringPtrInput
	Metadatas           ApplicationNameMetadataArrayInput
	Name                pulumi.StringPtrInput
	Parameter           pulumi.StringPtrInput
	Parameters          ApplicationNameParameterArrayInput
	Popularity          pulumi.IntPtrInput
	Protocol            pulumi.StringPtrInput
	Risk                pulumi.IntPtrInput
	SubCategory         pulumi.IntPtrInput
	Technology          pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	Vendor              pulumi.StringPtrInput
	Weight              pulumi.IntPtrInput
}

func (ApplicationNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationNameArgs)(nil)).Elem()
}

type ApplicationNameInput interface {
	pulumi.Input

	ToApplicationNameOutput() ApplicationNameOutput
	ToApplicationNameOutputWithContext(ctx context.Context) ApplicationNameOutput
}

func (*ApplicationName) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationName)(nil)).Elem()
}

func (i *ApplicationName) ToApplicationNameOutput() ApplicationNameOutput {
	return i.ToApplicationNameOutputWithContext(context.Background())
}

func (i *ApplicationName) ToApplicationNameOutputWithContext(ctx context.Context) ApplicationNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNameOutput)
}

func (i *ApplicationName) ToOutput(ctx context.Context) pulumix.Output[*ApplicationName] {
	return pulumix.Output[*ApplicationName]{
		OutputState: i.ToApplicationNameOutputWithContext(ctx).OutputState,
	}
}

// ApplicationNameArrayInput is an input type that accepts ApplicationNameArray and ApplicationNameArrayOutput values.
// You can construct a concrete instance of `ApplicationNameArrayInput` via:
//
//	ApplicationNameArray{ ApplicationNameArgs{...} }
type ApplicationNameArrayInput interface {
	pulumi.Input

	ToApplicationNameArrayOutput() ApplicationNameArrayOutput
	ToApplicationNameArrayOutputWithContext(context.Context) ApplicationNameArrayOutput
}

type ApplicationNameArray []ApplicationNameInput

func (ApplicationNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationName)(nil)).Elem()
}

func (i ApplicationNameArray) ToApplicationNameArrayOutput() ApplicationNameArrayOutput {
	return i.ToApplicationNameArrayOutputWithContext(context.Background())
}

func (i ApplicationNameArray) ToApplicationNameArrayOutputWithContext(ctx context.Context) ApplicationNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNameArrayOutput)
}

func (i ApplicationNameArray) ToOutput(ctx context.Context) pulumix.Output[[]*ApplicationName] {
	return pulumix.Output[[]*ApplicationName]{
		OutputState: i.ToApplicationNameArrayOutputWithContext(ctx).OutputState,
	}
}

// ApplicationNameMapInput is an input type that accepts ApplicationNameMap and ApplicationNameMapOutput values.
// You can construct a concrete instance of `ApplicationNameMapInput` via:
//
//	ApplicationNameMap{ "key": ApplicationNameArgs{...} }
type ApplicationNameMapInput interface {
	pulumi.Input

	ToApplicationNameMapOutput() ApplicationNameMapOutput
	ToApplicationNameMapOutputWithContext(context.Context) ApplicationNameMapOutput
}

type ApplicationNameMap map[string]ApplicationNameInput

func (ApplicationNameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationName)(nil)).Elem()
}

func (i ApplicationNameMap) ToApplicationNameMapOutput() ApplicationNameMapOutput {
	return i.ToApplicationNameMapOutputWithContext(context.Background())
}

func (i ApplicationNameMap) ToApplicationNameMapOutputWithContext(ctx context.Context) ApplicationNameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationNameMapOutput)
}

func (i ApplicationNameMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApplicationName] {
	return pulumix.Output[map[string]*ApplicationName]{
		OutputState: i.ToApplicationNameMapOutputWithContext(ctx).OutputState,
	}
}

type ApplicationNameOutput struct{ *pulumi.OutputState }

func (ApplicationNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationName)(nil)).Elem()
}

func (o ApplicationNameOutput) ToApplicationNameOutput() ApplicationNameOutput {
	return o
}

func (o ApplicationNameOutput) ToApplicationNameOutputWithContext(ctx context.Context) ApplicationNameOutput {
	return o
}

func (o ApplicationNameOutput) ToOutput(ctx context.Context) pulumix.Output[*ApplicationName] {
	return pulumix.Output[*ApplicationName]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationNameOutput) Behavior() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.StringOutput { return v.Behavior }).(pulumi.StringOutput)
}

func (o ApplicationNameOutput) Category() pulumi.IntOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.IntOutput { return v.Category }).(pulumi.IntOutput)
}

func (o ApplicationNameOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o ApplicationNameOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o ApplicationNameOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o ApplicationNameOutput) Metadatas() ApplicationNameMetadataArrayOutput {
	return o.ApplyT(func(v *ApplicationName) ApplicationNameMetadataArrayOutput { return v.Metadatas }).(ApplicationNameMetadataArrayOutput)
}

func (o ApplicationNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationNameOutput) Parameter() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.StringOutput { return v.Parameter }).(pulumi.StringOutput)
}

func (o ApplicationNameOutput) Parameters() ApplicationNameParameterArrayOutput {
	return o.ApplyT(func(v *ApplicationName) ApplicationNameParameterArrayOutput { return v.Parameters }).(ApplicationNameParameterArrayOutput)
}

func (o ApplicationNameOutput) Popularity() pulumi.IntOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.IntOutput { return v.Popularity }).(pulumi.IntOutput)
}

func (o ApplicationNameOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o ApplicationNameOutput) Risk() pulumi.IntOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.IntOutput { return v.Risk }).(pulumi.IntOutput)
}

func (o ApplicationNameOutput) SubCategory() pulumi.IntOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.IntOutput { return v.SubCategory }).(pulumi.IntOutput)
}

func (o ApplicationNameOutput) Technology() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.StringOutput { return v.Technology }).(pulumi.StringOutput)
}

func (o ApplicationNameOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o ApplicationNameOutput) Vendor() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.StringOutput { return v.Vendor }).(pulumi.StringOutput)
}

func (o ApplicationNameOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *ApplicationName) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type ApplicationNameArrayOutput struct{ *pulumi.OutputState }

func (ApplicationNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationName)(nil)).Elem()
}

func (o ApplicationNameArrayOutput) ToApplicationNameArrayOutput() ApplicationNameArrayOutput {
	return o
}

func (o ApplicationNameArrayOutput) ToApplicationNameArrayOutputWithContext(ctx context.Context) ApplicationNameArrayOutput {
	return o
}

func (o ApplicationNameArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ApplicationName] {
	return pulumix.Output[[]*ApplicationName]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationNameArrayOutput) Index(i pulumi.IntInput) ApplicationNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationName {
		return vs[0].([]*ApplicationName)[vs[1].(int)]
	}).(ApplicationNameOutput)
}

type ApplicationNameMapOutput struct{ *pulumi.OutputState }

func (ApplicationNameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationName)(nil)).Elem()
}

func (o ApplicationNameMapOutput) ToApplicationNameMapOutput() ApplicationNameMapOutput {
	return o
}

func (o ApplicationNameMapOutput) ToApplicationNameMapOutputWithContext(ctx context.Context) ApplicationNameMapOutput {
	return o
}

func (o ApplicationNameMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApplicationName] {
	return pulumix.Output[map[string]*ApplicationName]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationNameMapOutput) MapIndex(k pulumi.StringInput) ApplicationNameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationName {
		return vs[0].(map[string]*ApplicationName)[vs[1].(string)]
	}).(ApplicationNameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationNameInput)(nil)).Elem(), &ApplicationName{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationNameArrayInput)(nil)).Elem(), ApplicationNameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationNameMapInput)(nil)).Elem(), ApplicationNameMap{})
	pulumi.RegisterOutputType(ApplicationNameOutput{})
	pulumi.RegisterOutputType(ApplicationNameArrayOutput{})
	pulumi.RegisterOutputType(ApplicationNameMapOutput{})
}

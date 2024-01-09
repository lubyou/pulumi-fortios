// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemAcme struct {
	pulumi.CustomResourceState

	Accounts            SystemAcmeAccountArrayOutput   `pulumi:"accounts"`
	DynamicSortSubtable pulumi.StringPtrOutput         `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput         `pulumi:"getAllTables"`
	Interfaces          SystemAcmeInterfaceArrayOutput `pulumi:"interfaces"`
	SourceIp            pulumi.StringOutput            `pulumi:"sourceIp"`
	SourceIp6           pulumi.StringOutput            `pulumi:"sourceIp6"`
	UseHaDirect         pulumi.StringOutput            `pulumi:"useHaDirect"`
	Vdomparam           pulumi.StringPtrOutput         `pulumi:"vdomparam"`
}

// NewSystemAcme registers a new resource with the given unique name, arguments, and options.
func NewSystemAcme(ctx *pulumi.Context,
	name string, args *SystemAcmeArgs, opts ...pulumi.ResourceOption) (*SystemAcme, error) {
	if args == nil {
		args = &SystemAcmeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemAcme
	err := ctx.RegisterResource("fortios:index/systemAcme:SystemAcme", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAcme gets an existing SystemAcme resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAcme(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAcmeState, opts ...pulumi.ResourceOption) (*SystemAcme, error) {
	var resource SystemAcme
	err := ctx.ReadResource("fortios:index/systemAcme:SystemAcme", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAcme resources.
type systemAcmeState struct {
	Accounts            []SystemAcmeAccount   `pulumi:"accounts"`
	DynamicSortSubtable *string               `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string               `pulumi:"getAllTables"`
	Interfaces          []SystemAcmeInterface `pulumi:"interfaces"`
	SourceIp            *string               `pulumi:"sourceIp"`
	SourceIp6           *string               `pulumi:"sourceIp6"`
	UseHaDirect         *string               `pulumi:"useHaDirect"`
	Vdomparam           *string               `pulumi:"vdomparam"`
}

type SystemAcmeState struct {
	Accounts            SystemAcmeAccountArrayInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Interfaces          SystemAcmeInterfaceArrayInput
	SourceIp            pulumi.StringPtrInput
	SourceIp6           pulumi.StringPtrInput
	UseHaDirect         pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SystemAcmeState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAcmeState)(nil)).Elem()
}

type systemAcmeArgs struct {
	Accounts            []SystemAcmeAccount   `pulumi:"accounts"`
	DynamicSortSubtable *string               `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string               `pulumi:"getAllTables"`
	Interfaces          []SystemAcmeInterface `pulumi:"interfaces"`
	SourceIp            *string               `pulumi:"sourceIp"`
	SourceIp6           *string               `pulumi:"sourceIp6"`
	UseHaDirect         *string               `pulumi:"useHaDirect"`
	Vdomparam           *string               `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAcme resource.
type SystemAcmeArgs struct {
	Accounts            SystemAcmeAccountArrayInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Interfaces          SystemAcmeInterfaceArrayInput
	SourceIp            pulumi.StringPtrInput
	SourceIp6           pulumi.StringPtrInput
	UseHaDirect         pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SystemAcmeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAcmeArgs)(nil)).Elem()
}

type SystemAcmeInput interface {
	pulumi.Input

	ToSystemAcmeOutput() SystemAcmeOutput
	ToSystemAcmeOutputWithContext(ctx context.Context) SystemAcmeOutput
}

func (*SystemAcme) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAcme)(nil)).Elem()
}

func (i *SystemAcme) ToSystemAcmeOutput() SystemAcmeOutput {
	return i.ToSystemAcmeOutputWithContext(context.Background())
}

func (i *SystemAcme) ToSystemAcmeOutputWithContext(ctx context.Context) SystemAcmeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAcmeOutput)
}

// SystemAcmeArrayInput is an input type that accepts SystemAcmeArray and SystemAcmeArrayOutput values.
// You can construct a concrete instance of `SystemAcmeArrayInput` via:
//
//	SystemAcmeArray{ SystemAcmeArgs{...} }
type SystemAcmeArrayInput interface {
	pulumi.Input

	ToSystemAcmeArrayOutput() SystemAcmeArrayOutput
	ToSystemAcmeArrayOutputWithContext(context.Context) SystemAcmeArrayOutput
}

type SystemAcmeArray []SystemAcmeInput

func (SystemAcmeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAcme)(nil)).Elem()
}

func (i SystemAcmeArray) ToSystemAcmeArrayOutput() SystemAcmeArrayOutput {
	return i.ToSystemAcmeArrayOutputWithContext(context.Background())
}

func (i SystemAcmeArray) ToSystemAcmeArrayOutputWithContext(ctx context.Context) SystemAcmeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAcmeArrayOutput)
}

// SystemAcmeMapInput is an input type that accepts SystemAcmeMap and SystemAcmeMapOutput values.
// You can construct a concrete instance of `SystemAcmeMapInput` via:
//
//	SystemAcmeMap{ "key": SystemAcmeArgs{...} }
type SystemAcmeMapInput interface {
	pulumi.Input

	ToSystemAcmeMapOutput() SystemAcmeMapOutput
	ToSystemAcmeMapOutputWithContext(context.Context) SystemAcmeMapOutput
}

type SystemAcmeMap map[string]SystemAcmeInput

func (SystemAcmeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAcme)(nil)).Elem()
}

func (i SystemAcmeMap) ToSystemAcmeMapOutput() SystemAcmeMapOutput {
	return i.ToSystemAcmeMapOutputWithContext(context.Background())
}

func (i SystemAcmeMap) ToSystemAcmeMapOutputWithContext(ctx context.Context) SystemAcmeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAcmeMapOutput)
}

type SystemAcmeOutput struct{ *pulumi.OutputState }

func (SystemAcmeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAcme)(nil)).Elem()
}

func (o SystemAcmeOutput) ToSystemAcmeOutput() SystemAcmeOutput {
	return o
}

func (o SystemAcmeOutput) ToSystemAcmeOutputWithContext(ctx context.Context) SystemAcmeOutput {
	return o
}

func (o SystemAcmeOutput) Accounts() SystemAcmeAccountArrayOutput {
	return o.ApplyT(func(v *SystemAcme) SystemAcmeAccountArrayOutput { return v.Accounts }).(SystemAcmeAccountArrayOutput)
}

func (o SystemAcmeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAcme) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SystemAcmeOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAcme) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SystemAcmeOutput) Interfaces() SystemAcmeInterfaceArrayOutput {
	return o.ApplyT(func(v *SystemAcme) SystemAcmeInterfaceArrayOutput { return v.Interfaces }).(SystemAcmeInterfaceArrayOutput)
}

func (o SystemAcmeOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAcme) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o SystemAcmeOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAcme) pulumi.StringOutput { return v.SourceIp6 }).(pulumi.StringOutput)
}

func (o SystemAcmeOutput) UseHaDirect() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAcme) pulumi.StringOutput { return v.UseHaDirect }).(pulumi.StringOutput)
}

func (o SystemAcmeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAcme) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemAcmeArrayOutput struct{ *pulumi.OutputState }

func (SystemAcmeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAcme)(nil)).Elem()
}

func (o SystemAcmeArrayOutput) ToSystemAcmeArrayOutput() SystemAcmeArrayOutput {
	return o
}

func (o SystemAcmeArrayOutput) ToSystemAcmeArrayOutputWithContext(ctx context.Context) SystemAcmeArrayOutput {
	return o
}

func (o SystemAcmeArrayOutput) Index(i pulumi.IntInput) SystemAcmeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAcme {
		return vs[0].([]*SystemAcme)[vs[1].(int)]
	}).(SystemAcmeOutput)
}

type SystemAcmeMapOutput struct{ *pulumi.OutputState }

func (SystemAcmeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAcme)(nil)).Elem()
}

func (o SystemAcmeMapOutput) ToSystemAcmeMapOutput() SystemAcmeMapOutput {
	return o
}

func (o SystemAcmeMapOutput) ToSystemAcmeMapOutputWithContext(ctx context.Context) SystemAcmeMapOutput {
	return o
}

func (o SystemAcmeMapOutput) MapIndex(k pulumi.StringInput) SystemAcmeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAcme {
		return vs[0].(map[string]*SystemAcme)[vs[1].(string)]
	}).(SystemAcmeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAcmeInput)(nil)).Elem(), &SystemAcme{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAcmeArrayInput)(nil)).Elem(), SystemAcmeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAcmeMapInput)(nil)).Elem(), SystemAcmeMap{})
	pulumi.RegisterOutputType(SystemAcmeOutput{})
	pulumi.RegisterOutputType(SystemAcmeArrayOutput{})
	pulumi.RegisterOutputType(SystemAcmeMapOutput{})
}

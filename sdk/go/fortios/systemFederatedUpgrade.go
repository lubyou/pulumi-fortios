// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemFederatedUpgrade struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput                    `pulumi:"dynamicSortSubtable"`
	FailureDevice       pulumi.StringOutput                       `pulumi:"failureDevice"`
	FailureReason       pulumi.StringOutput                       `pulumi:"failureReason"`
	GetAllTables        pulumi.StringPtrOutput                    `pulumi:"getAllTables"`
	HaRebootController  pulumi.StringOutput                       `pulumi:"haRebootController"`
	NextPathIndex       pulumi.IntOutput                          `pulumi:"nextPathIndex"`
	NodeLists           SystemFederatedUpgradeNodeListArrayOutput `pulumi:"nodeLists"`
	Status              pulumi.StringOutput                       `pulumi:"status"`
	UpgradeId           pulumi.IntOutput                          `pulumi:"upgradeId"`
	Vdomparam           pulumi.StringPtrOutput                    `pulumi:"vdomparam"`
}

// NewSystemFederatedUpgrade registers a new resource with the given unique name, arguments, and options.
func NewSystemFederatedUpgrade(ctx *pulumi.Context,
	name string, args *SystemFederatedUpgradeArgs, opts ...pulumi.ResourceOption) (*SystemFederatedUpgrade, error) {
	if args == nil {
		args = &SystemFederatedUpgradeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemFederatedUpgrade
	err := ctx.RegisterResource("fortios:index/systemFederatedUpgrade:SystemFederatedUpgrade", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFederatedUpgrade gets an existing SystemFederatedUpgrade resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFederatedUpgrade(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFederatedUpgradeState, opts ...pulumi.ResourceOption) (*SystemFederatedUpgrade, error) {
	var resource SystemFederatedUpgrade
	err := ctx.ReadResource("fortios:index/systemFederatedUpgrade:SystemFederatedUpgrade", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFederatedUpgrade resources.
type systemFederatedUpgradeState struct {
	DynamicSortSubtable *string                          `pulumi:"dynamicSortSubtable"`
	FailureDevice       *string                          `pulumi:"failureDevice"`
	FailureReason       *string                          `pulumi:"failureReason"`
	GetAllTables        *string                          `pulumi:"getAllTables"`
	HaRebootController  *string                          `pulumi:"haRebootController"`
	NextPathIndex       *int                             `pulumi:"nextPathIndex"`
	NodeLists           []SystemFederatedUpgradeNodeList `pulumi:"nodeLists"`
	Status              *string                          `pulumi:"status"`
	UpgradeId           *int                             `pulumi:"upgradeId"`
	Vdomparam           *string                          `pulumi:"vdomparam"`
}

type SystemFederatedUpgradeState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	FailureDevice       pulumi.StringPtrInput
	FailureReason       pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	HaRebootController  pulumi.StringPtrInput
	NextPathIndex       pulumi.IntPtrInput
	NodeLists           SystemFederatedUpgradeNodeListArrayInput
	Status              pulumi.StringPtrInput
	UpgradeId           pulumi.IntPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SystemFederatedUpgradeState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFederatedUpgradeState)(nil)).Elem()
}

type systemFederatedUpgradeArgs struct {
	DynamicSortSubtable *string                          `pulumi:"dynamicSortSubtable"`
	FailureDevice       *string                          `pulumi:"failureDevice"`
	FailureReason       *string                          `pulumi:"failureReason"`
	GetAllTables        *string                          `pulumi:"getAllTables"`
	HaRebootController  *string                          `pulumi:"haRebootController"`
	NextPathIndex       *int                             `pulumi:"nextPathIndex"`
	NodeLists           []SystemFederatedUpgradeNodeList `pulumi:"nodeLists"`
	Status              *string                          `pulumi:"status"`
	UpgradeId           *int                             `pulumi:"upgradeId"`
	Vdomparam           *string                          `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFederatedUpgrade resource.
type SystemFederatedUpgradeArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	FailureDevice       pulumi.StringPtrInput
	FailureReason       pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	HaRebootController  pulumi.StringPtrInput
	NextPathIndex       pulumi.IntPtrInput
	NodeLists           SystemFederatedUpgradeNodeListArrayInput
	Status              pulumi.StringPtrInput
	UpgradeId           pulumi.IntPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SystemFederatedUpgradeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFederatedUpgradeArgs)(nil)).Elem()
}

type SystemFederatedUpgradeInput interface {
	pulumi.Input

	ToSystemFederatedUpgradeOutput() SystemFederatedUpgradeOutput
	ToSystemFederatedUpgradeOutputWithContext(ctx context.Context) SystemFederatedUpgradeOutput
}

func (*SystemFederatedUpgrade) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFederatedUpgrade)(nil)).Elem()
}

func (i *SystemFederatedUpgrade) ToSystemFederatedUpgradeOutput() SystemFederatedUpgradeOutput {
	return i.ToSystemFederatedUpgradeOutputWithContext(context.Background())
}

func (i *SystemFederatedUpgrade) ToSystemFederatedUpgradeOutputWithContext(ctx context.Context) SystemFederatedUpgradeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFederatedUpgradeOutput)
}

// SystemFederatedUpgradeArrayInput is an input type that accepts SystemFederatedUpgradeArray and SystemFederatedUpgradeArrayOutput values.
// You can construct a concrete instance of `SystemFederatedUpgradeArrayInput` via:
//
//	SystemFederatedUpgradeArray{ SystemFederatedUpgradeArgs{...} }
type SystemFederatedUpgradeArrayInput interface {
	pulumi.Input

	ToSystemFederatedUpgradeArrayOutput() SystemFederatedUpgradeArrayOutput
	ToSystemFederatedUpgradeArrayOutputWithContext(context.Context) SystemFederatedUpgradeArrayOutput
}

type SystemFederatedUpgradeArray []SystemFederatedUpgradeInput

func (SystemFederatedUpgradeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFederatedUpgrade)(nil)).Elem()
}

func (i SystemFederatedUpgradeArray) ToSystemFederatedUpgradeArrayOutput() SystemFederatedUpgradeArrayOutput {
	return i.ToSystemFederatedUpgradeArrayOutputWithContext(context.Background())
}

func (i SystemFederatedUpgradeArray) ToSystemFederatedUpgradeArrayOutputWithContext(ctx context.Context) SystemFederatedUpgradeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFederatedUpgradeArrayOutput)
}

// SystemFederatedUpgradeMapInput is an input type that accepts SystemFederatedUpgradeMap and SystemFederatedUpgradeMapOutput values.
// You can construct a concrete instance of `SystemFederatedUpgradeMapInput` via:
//
//	SystemFederatedUpgradeMap{ "key": SystemFederatedUpgradeArgs{...} }
type SystemFederatedUpgradeMapInput interface {
	pulumi.Input

	ToSystemFederatedUpgradeMapOutput() SystemFederatedUpgradeMapOutput
	ToSystemFederatedUpgradeMapOutputWithContext(context.Context) SystemFederatedUpgradeMapOutput
}

type SystemFederatedUpgradeMap map[string]SystemFederatedUpgradeInput

func (SystemFederatedUpgradeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFederatedUpgrade)(nil)).Elem()
}

func (i SystemFederatedUpgradeMap) ToSystemFederatedUpgradeMapOutput() SystemFederatedUpgradeMapOutput {
	return i.ToSystemFederatedUpgradeMapOutputWithContext(context.Background())
}

func (i SystemFederatedUpgradeMap) ToSystemFederatedUpgradeMapOutputWithContext(ctx context.Context) SystemFederatedUpgradeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFederatedUpgradeMapOutput)
}

type SystemFederatedUpgradeOutput struct{ *pulumi.OutputState }

func (SystemFederatedUpgradeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFederatedUpgrade)(nil)).Elem()
}

func (o SystemFederatedUpgradeOutput) ToSystemFederatedUpgradeOutput() SystemFederatedUpgradeOutput {
	return o
}

func (o SystemFederatedUpgradeOutput) ToSystemFederatedUpgradeOutputWithContext(ctx context.Context) SystemFederatedUpgradeOutput {
	return o
}

func (o SystemFederatedUpgradeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFederatedUpgrade) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SystemFederatedUpgradeOutput) FailureDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFederatedUpgrade) pulumi.StringOutput { return v.FailureDevice }).(pulumi.StringOutput)
}

func (o SystemFederatedUpgradeOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFederatedUpgrade) pulumi.StringOutput { return v.FailureReason }).(pulumi.StringOutput)
}

func (o SystemFederatedUpgradeOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFederatedUpgrade) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SystemFederatedUpgradeOutput) HaRebootController() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFederatedUpgrade) pulumi.StringOutput { return v.HaRebootController }).(pulumi.StringOutput)
}

func (o SystemFederatedUpgradeOutput) NextPathIndex() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemFederatedUpgrade) pulumi.IntOutput { return v.NextPathIndex }).(pulumi.IntOutput)
}

func (o SystemFederatedUpgradeOutput) NodeLists() SystemFederatedUpgradeNodeListArrayOutput {
	return o.ApplyT(func(v *SystemFederatedUpgrade) SystemFederatedUpgradeNodeListArrayOutput { return v.NodeLists }).(SystemFederatedUpgradeNodeListArrayOutput)
}

func (o SystemFederatedUpgradeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFederatedUpgrade) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SystemFederatedUpgradeOutput) UpgradeId() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemFederatedUpgrade) pulumi.IntOutput { return v.UpgradeId }).(pulumi.IntOutput)
}

func (o SystemFederatedUpgradeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFederatedUpgrade) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemFederatedUpgradeArrayOutput struct{ *pulumi.OutputState }

func (SystemFederatedUpgradeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFederatedUpgrade)(nil)).Elem()
}

func (o SystemFederatedUpgradeArrayOutput) ToSystemFederatedUpgradeArrayOutput() SystemFederatedUpgradeArrayOutput {
	return o
}

func (o SystemFederatedUpgradeArrayOutput) ToSystemFederatedUpgradeArrayOutputWithContext(ctx context.Context) SystemFederatedUpgradeArrayOutput {
	return o
}

func (o SystemFederatedUpgradeArrayOutput) Index(i pulumi.IntInput) SystemFederatedUpgradeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemFederatedUpgrade {
		return vs[0].([]*SystemFederatedUpgrade)[vs[1].(int)]
	}).(SystemFederatedUpgradeOutput)
}

type SystemFederatedUpgradeMapOutput struct{ *pulumi.OutputState }

func (SystemFederatedUpgradeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFederatedUpgrade)(nil)).Elem()
}

func (o SystemFederatedUpgradeMapOutput) ToSystemFederatedUpgradeMapOutput() SystemFederatedUpgradeMapOutput {
	return o
}

func (o SystemFederatedUpgradeMapOutput) ToSystemFederatedUpgradeMapOutputWithContext(ctx context.Context) SystemFederatedUpgradeMapOutput {
	return o
}

func (o SystemFederatedUpgradeMapOutput) MapIndex(k pulumi.StringInput) SystemFederatedUpgradeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemFederatedUpgrade {
		return vs[0].(map[string]*SystemFederatedUpgrade)[vs[1].(string)]
	}).(SystemFederatedUpgradeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFederatedUpgradeInput)(nil)).Elem(), &SystemFederatedUpgrade{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFederatedUpgradeArrayInput)(nil)).Elem(), SystemFederatedUpgradeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFederatedUpgradeMapInput)(nil)).Elem(), SystemFederatedUpgradeMap{})
	pulumi.RegisterOutputType(SystemFederatedUpgradeOutput{})
	pulumi.RegisterOutputType(SystemFederatedUpgradeArrayOutput{})
	pulumi.RegisterOutputType(SystemFederatedUpgradeMapOutput{})
}

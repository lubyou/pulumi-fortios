// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FortimanagerObjectAdomRevision struct {
	pulumi.CustomResourceState

	Adom        pulumi.StringPtrOutput `pulumi:"adom"`
	CreatedBy   pulumi.StringPtrOutput `pulumi:"createdBy"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	Locked      pulumi.IntPtrOutput    `pulumi:"locked"`
	Name        pulumi.StringOutput    `pulumi:"name"`
}

// NewFortimanagerObjectAdomRevision registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerObjectAdomRevision(ctx *pulumi.Context,
	name string, args *FortimanagerObjectAdomRevisionArgs, opts ...pulumi.ResourceOption) (*FortimanagerObjectAdomRevision, error) {
	if args == nil {
		args = &FortimanagerObjectAdomRevisionArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FortimanagerObjectAdomRevision
	err := ctx.RegisterResource("fortios:index/fortimanagerObjectAdomRevision:FortimanagerObjectAdomRevision", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerObjectAdomRevision gets an existing FortimanagerObjectAdomRevision resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerObjectAdomRevision(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerObjectAdomRevisionState, opts ...pulumi.ResourceOption) (*FortimanagerObjectAdomRevision, error) {
	var resource FortimanagerObjectAdomRevision
	err := ctx.ReadResource("fortios:index/fortimanagerObjectAdomRevision:FortimanagerObjectAdomRevision", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerObjectAdomRevision resources.
type fortimanagerObjectAdomRevisionState struct {
	Adom        *string `pulumi:"adom"`
	CreatedBy   *string `pulumi:"createdBy"`
	Description *string `pulumi:"description"`
	Locked      *int    `pulumi:"locked"`
	Name        *string `pulumi:"name"`
}

type FortimanagerObjectAdomRevisionState struct {
	Adom        pulumi.StringPtrInput
	CreatedBy   pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	Locked      pulumi.IntPtrInput
	Name        pulumi.StringPtrInput
}

func (FortimanagerObjectAdomRevisionState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerObjectAdomRevisionState)(nil)).Elem()
}

type fortimanagerObjectAdomRevisionArgs struct {
	Adom        *string `pulumi:"adom"`
	CreatedBy   *string `pulumi:"createdBy"`
	Description *string `pulumi:"description"`
	Locked      *int    `pulumi:"locked"`
	Name        *string `pulumi:"name"`
}

// The set of arguments for constructing a FortimanagerObjectAdomRevision resource.
type FortimanagerObjectAdomRevisionArgs struct {
	Adom        pulumi.StringPtrInput
	CreatedBy   pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	Locked      pulumi.IntPtrInput
	Name        pulumi.StringPtrInput
}

func (FortimanagerObjectAdomRevisionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerObjectAdomRevisionArgs)(nil)).Elem()
}

type FortimanagerObjectAdomRevisionInput interface {
	pulumi.Input

	ToFortimanagerObjectAdomRevisionOutput() FortimanagerObjectAdomRevisionOutput
	ToFortimanagerObjectAdomRevisionOutputWithContext(ctx context.Context) FortimanagerObjectAdomRevisionOutput
}

func (*FortimanagerObjectAdomRevision) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerObjectAdomRevision)(nil)).Elem()
}

func (i *FortimanagerObjectAdomRevision) ToFortimanagerObjectAdomRevisionOutput() FortimanagerObjectAdomRevisionOutput {
	return i.ToFortimanagerObjectAdomRevisionOutputWithContext(context.Background())
}

func (i *FortimanagerObjectAdomRevision) ToFortimanagerObjectAdomRevisionOutputWithContext(ctx context.Context) FortimanagerObjectAdomRevisionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerObjectAdomRevisionOutput)
}

// FortimanagerObjectAdomRevisionArrayInput is an input type that accepts FortimanagerObjectAdomRevisionArray and FortimanagerObjectAdomRevisionArrayOutput values.
// You can construct a concrete instance of `FortimanagerObjectAdomRevisionArrayInput` via:
//
//	FortimanagerObjectAdomRevisionArray{ FortimanagerObjectAdomRevisionArgs{...} }
type FortimanagerObjectAdomRevisionArrayInput interface {
	pulumi.Input

	ToFortimanagerObjectAdomRevisionArrayOutput() FortimanagerObjectAdomRevisionArrayOutput
	ToFortimanagerObjectAdomRevisionArrayOutputWithContext(context.Context) FortimanagerObjectAdomRevisionArrayOutput
}

type FortimanagerObjectAdomRevisionArray []FortimanagerObjectAdomRevisionInput

func (FortimanagerObjectAdomRevisionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerObjectAdomRevision)(nil)).Elem()
}

func (i FortimanagerObjectAdomRevisionArray) ToFortimanagerObjectAdomRevisionArrayOutput() FortimanagerObjectAdomRevisionArrayOutput {
	return i.ToFortimanagerObjectAdomRevisionArrayOutputWithContext(context.Background())
}

func (i FortimanagerObjectAdomRevisionArray) ToFortimanagerObjectAdomRevisionArrayOutputWithContext(ctx context.Context) FortimanagerObjectAdomRevisionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerObjectAdomRevisionArrayOutput)
}

// FortimanagerObjectAdomRevisionMapInput is an input type that accepts FortimanagerObjectAdomRevisionMap and FortimanagerObjectAdomRevisionMapOutput values.
// You can construct a concrete instance of `FortimanagerObjectAdomRevisionMapInput` via:
//
//	FortimanagerObjectAdomRevisionMap{ "key": FortimanagerObjectAdomRevisionArgs{...} }
type FortimanagerObjectAdomRevisionMapInput interface {
	pulumi.Input

	ToFortimanagerObjectAdomRevisionMapOutput() FortimanagerObjectAdomRevisionMapOutput
	ToFortimanagerObjectAdomRevisionMapOutputWithContext(context.Context) FortimanagerObjectAdomRevisionMapOutput
}

type FortimanagerObjectAdomRevisionMap map[string]FortimanagerObjectAdomRevisionInput

func (FortimanagerObjectAdomRevisionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerObjectAdomRevision)(nil)).Elem()
}

func (i FortimanagerObjectAdomRevisionMap) ToFortimanagerObjectAdomRevisionMapOutput() FortimanagerObjectAdomRevisionMapOutput {
	return i.ToFortimanagerObjectAdomRevisionMapOutputWithContext(context.Background())
}

func (i FortimanagerObjectAdomRevisionMap) ToFortimanagerObjectAdomRevisionMapOutputWithContext(ctx context.Context) FortimanagerObjectAdomRevisionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerObjectAdomRevisionMapOutput)
}

type FortimanagerObjectAdomRevisionOutput struct{ *pulumi.OutputState }

func (FortimanagerObjectAdomRevisionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerObjectAdomRevision)(nil)).Elem()
}

func (o FortimanagerObjectAdomRevisionOutput) ToFortimanagerObjectAdomRevisionOutput() FortimanagerObjectAdomRevisionOutput {
	return o
}

func (o FortimanagerObjectAdomRevisionOutput) ToFortimanagerObjectAdomRevisionOutputWithContext(ctx context.Context) FortimanagerObjectAdomRevisionOutput {
	return o
}

func (o FortimanagerObjectAdomRevisionOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerObjectAdomRevision) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

func (o FortimanagerObjectAdomRevisionOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerObjectAdomRevision) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o FortimanagerObjectAdomRevisionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerObjectAdomRevision) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o FortimanagerObjectAdomRevisionOutput) Locked() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FortimanagerObjectAdomRevision) pulumi.IntPtrOutput { return v.Locked }).(pulumi.IntPtrOutput)
}

func (o FortimanagerObjectAdomRevisionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerObjectAdomRevision) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type FortimanagerObjectAdomRevisionArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerObjectAdomRevisionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerObjectAdomRevision)(nil)).Elem()
}

func (o FortimanagerObjectAdomRevisionArrayOutput) ToFortimanagerObjectAdomRevisionArrayOutput() FortimanagerObjectAdomRevisionArrayOutput {
	return o
}

func (o FortimanagerObjectAdomRevisionArrayOutput) ToFortimanagerObjectAdomRevisionArrayOutputWithContext(ctx context.Context) FortimanagerObjectAdomRevisionArrayOutput {
	return o
}

func (o FortimanagerObjectAdomRevisionArrayOutput) Index(i pulumi.IntInput) FortimanagerObjectAdomRevisionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerObjectAdomRevision {
		return vs[0].([]*FortimanagerObjectAdomRevision)[vs[1].(int)]
	}).(FortimanagerObjectAdomRevisionOutput)
}

type FortimanagerObjectAdomRevisionMapOutput struct{ *pulumi.OutputState }

func (FortimanagerObjectAdomRevisionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerObjectAdomRevision)(nil)).Elem()
}

func (o FortimanagerObjectAdomRevisionMapOutput) ToFortimanagerObjectAdomRevisionMapOutput() FortimanagerObjectAdomRevisionMapOutput {
	return o
}

func (o FortimanagerObjectAdomRevisionMapOutput) ToFortimanagerObjectAdomRevisionMapOutputWithContext(ctx context.Context) FortimanagerObjectAdomRevisionMapOutput {
	return o
}

func (o FortimanagerObjectAdomRevisionMapOutput) MapIndex(k pulumi.StringInput) FortimanagerObjectAdomRevisionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerObjectAdomRevision {
		return vs[0].(map[string]*FortimanagerObjectAdomRevision)[vs[1].(string)]
	}).(FortimanagerObjectAdomRevisionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerObjectAdomRevisionInput)(nil)).Elem(), &FortimanagerObjectAdomRevision{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerObjectAdomRevisionArrayInput)(nil)).Elem(), FortimanagerObjectAdomRevisionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerObjectAdomRevisionMapInput)(nil)).Elem(), FortimanagerObjectAdomRevisionMap{})
	pulumi.RegisterOutputType(FortimanagerObjectAdomRevisionOutput{})
	pulumi.RegisterOutputType(FortimanagerObjectAdomRevisionArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerObjectAdomRevisionMapOutput{})
}

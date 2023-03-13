// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FortimanagerSystemDNS struct {
	pulumi.CustomResourceState

	Primary   pulumi.StringPtrOutput `pulumi:"primary"`
	Secondary pulumi.StringPtrOutput `pulumi:"secondary"`
}

// NewFortimanagerSystemDNS registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemDNS(ctx *pulumi.Context,
	name string, args *FortimanagerSystemDNSArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemDNS, error) {
	if args == nil {
		args = &FortimanagerSystemDNSArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FortimanagerSystemDNS
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemDNS:FortimanagerSystemDNS", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemDNS gets an existing FortimanagerSystemDNS resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemDNS(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemDNSState, opts ...pulumi.ResourceOption) (*FortimanagerSystemDNS, error) {
	var resource FortimanagerSystemDNS
	err := ctx.ReadResource("fortios:index/fortimanagerSystemDNS:FortimanagerSystemDNS", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemDNS resources.
type fortimanagerSystemDNSState struct {
	Primary   *string `pulumi:"primary"`
	Secondary *string `pulumi:"secondary"`
}

type FortimanagerSystemDNSState struct {
	Primary   pulumi.StringPtrInput
	Secondary pulumi.StringPtrInput
}

func (FortimanagerSystemDNSState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemDNSState)(nil)).Elem()
}

type fortimanagerSystemDNSArgs struct {
	Primary   *string `pulumi:"primary"`
	Secondary *string `pulumi:"secondary"`
}

// The set of arguments for constructing a FortimanagerSystemDNS resource.
type FortimanagerSystemDNSArgs struct {
	Primary   pulumi.StringPtrInput
	Secondary pulumi.StringPtrInput
}

func (FortimanagerSystemDNSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemDNSArgs)(nil)).Elem()
}

type FortimanagerSystemDNSInput interface {
	pulumi.Input

	ToFortimanagerSystemDNSOutput() FortimanagerSystemDNSOutput
	ToFortimanagerSystemDNSOutputWithContext(ctx context.Context) FortimanagerSystemDNSOutput
}

func (*FortimanagerSystemDNS) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemDNS)(nil)).Elem()
}

func (i *FortimanagerSystemDNS) ToFortimanagerSystemDNSOutput() FortimanagerSystemDNSOutput {
	return i.ToFortimanagerSystemDNSOutputWithContext(context.Background())
}

func (i *FortimanagerSystemDNS) ToFortimanagerSystemDNSOutputWithContext(ctx context.Context) FortimanagerSystemDNSOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemDNSOutput)
}

// FortimanagerSystemDNSArrayInput is an input type that accepts FortimanagerSystemDNSArray and FortimanagerSystemDNSArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemDNSArrayInput` via:
//
//	FortimanagerSystemDNSArray{ FortimanagerSystemDNSArgs{...} }
type FortimanagerSystemDNSArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemDNSArrayOutput() FortimanagerSystemDNSArrayOutput
	ToFortimanagerSystemDNSArrayOutputWithContext(context.Context) FortimanagerSystemDNSArrayOutput
}

type FortimanagerSystemDNSArray []FortimanagerSystemDNSInput

func (FortimanagerSystemDNSArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemDNS)(nil)).Elem()
}

func (i FortimanagerSystemDNSArray) ToFortimanagerSystemDNSArrayOutput() FortimanagerSystemDNSArrayOutput {
	return i.ToFortimanagerSystemDNSArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemDNSArray) ToFortimanagerSystemDNSArrayOutputWithContext(ctx context.Context) FortimanagerSystemDNSArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemDNSArrayOutput)
}

// FortimanagerSystemDNSMapInput is an input type that accepts FortimanagerSystemDNSMap and FortimanagerSystemDNSMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemDNSMapInput` via:
//
//	FortimanagerSystemDNSMap{ "key": FortimanagerSystemDNSArgs{...} }
type FortimanagerSystemDNSMapInput interface {
	pulumi.Input

	ToFortimanagerSystemDNSMapOutput() FortimanagerSystemDNSMapOutput
	ToFortimanagerSystemDNSMapOutputWithContext(context.Context) FortimanagerSystemDNSMapOutput
}

type FortimanagerSystemDNSMap map[string]FortimanagerSystemDNSInput

func (FortimanagerSystemDNSMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemDNS)(nil)).Elem()
}

func (i FortimanagerSystemDNSMap) ToFortimanagerSystemDNSMapOutput() FortimanagerSystemDNSMapOutput {
	return i.ToFortimanagerSystemDNSMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemDNSMap) ToFortimanagerSystemDNSMapOutputWithContext(ctx context.Context) FortimanagerSystemDNSMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemDNSMapOutput)
}

type FortimanagerSystemDNSOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemDNSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemDNS)(nil)).Elem()
}

func (o FortimanagerSystemDNSOutput) ToFortimanagerSystemDNSOutput() FortimanagerSystemDNSOutput {
	return o
}

func (o FortimanagerSystemDNSOutput) ToFortimanagerSystemDNSOutputWithContext(ctx context.Context) FortimanagerSystemDNSOutput {
	return o
}

func (o FortimanagerSystemDNSOutput) Primary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemDNS) pulumi.StringPtrOutput { return v.Primary }).(pulumi.StringPtrOutput)
}

func (o FortimanagerSystemDNSOutput) Secondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FortimanagerSystemDNS) pulumi.StringPtrOutput { return v.Secondary }).(pulumi.StringPtrOutput)
}

type FortimanagerSystemDNSArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemDNSArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemDNS)(nil)).Elem()
}

func (o FortimanagerSystemDNSArrayOutput) ToFortimanagerSystemDNSArrayOutput() FortimanagerSystemDNSArrayOutput {
	return o
}

func (o FortimanagerSystemDNSArrayOutput) ToFortimanagerSystemDNSArrayOutputWithContext(ctx context.Context) FortimanagerSystemDNSArrayOutput {
	return o
}

func (o FortimanagerSystemDNSArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemDNSOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerSystemDNS {
		return vs[0].([]*FortimanagerSystemDNS)[vs[1].(int)]
	}).(FortimanagerSystemDNSOutput)
}

type FortimanagerSystemDNSMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemDNSMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemDNS)(nil)).Elem()
}

func (o FortimanagerSystemDNSMapOutput) ToFortimanagerSystemDNSMapOutput() FortimanagerSystemDNSMapOutput {
	return o
}

func (o FortimanagerSystemDNSMapOutput) ToFortimanagerSystemDNSMapOutputWithContext(ctx context.Context) FortimanagerSystemDNSMapOutput {
	return o
}

func (o FortimanagerSystemDNSMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemDNSOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerSystemDNS {
		return vs[0].(map[string]*FortimanagerSystemDNS)[vs[1].(string)]
	}).(FortimanagerSystemDNSOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemDNSInput)(nil)).Elem(), &FortimanagerSystemDNS{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemDNSArrayInput)(nil)).Elem(), FortimanagerSystemDNSArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemDNSMapInput)(nil)).Elem(), FortimanagerSystemDNSMap{})
	pulumi.RegisterOutputType(FortimanagerSystemDNSOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemDNSArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemDNSMapOutput{})
}

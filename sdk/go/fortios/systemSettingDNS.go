// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemSettingDNS struct {
	pulumi.CustomResourceState

	DnsOverTls pulumi.StringOutput `pulumi:"dnsOverTls"`
	Primary    pulumi.StringOutput `pulumi:"primary"`
	Secondary  pulumi.StringOutput `pulumi:"secondary"`
}

// NewSystemSettingDNS registers a new resource with the given unique name, arguments, and options.
func NewSystemSettingDNS(ctx *pulumi.Context,
	name string, args *SystemSettingDNSArgs, opts ...pulumi.ResourceOption) (*SystemSettingDNS, error) {
	if args == nil {
		args = &SystemSettingDNSArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSettingDNS
	err := ctx.RegisterResource("fortios:index/systemSettingDNS:SystemSettingDNS", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSettingDNS gets an existing SystemSettingDNS resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSettingDNS(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSettingDNSState, opts ...pulumi.ResourceOption) (*SystemSettingDNS, error) {
	var resource SystemSettingDNS
	err := ctx.ReadResource("fortios:index/systemSettingDNS:SystemSettingDNS", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSettingDNS resources.
type systemSettingDNSState struct {
	DnsOverTls *string `pulumi:"dnsOverTls"`
	Primary    *string `pulumi:"primary"`
	Secondary  *string `pulumi:"secondary"`
}

type SystemSettingDNSState struct {
	DnsOverTls pulumi.StringPtrInput
	Primary    pulumi.StringPtrInput
	Secondary  pulumi.StringPtrInput
}

func (SystemSettingDNSState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSettingDNSState)(nil)).Elem()
}

type systemSettingDNSArgs struct {
	DnsOverTls *string `pulumi:"dnsOverTls"`
	Primary    *string `pulumi:"primary"`
	Secondary  *string `pulumi:"secondary"`
}

// The set of arguments for constructing a SystemSettingDNS resource.
type SystemSettingDNSArgs struct {
	DnsOverTls pulumi.StringPtrInput
	Primary    pulumi.StringPtrInput
	Secondary  pulumi.StringPtrInput
}

func (SystemSettingDNSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSettingDNSArgs)(nil)).Elem()
}

type SystemSettingDNSInput interface {
	pulumi.Input

	ToSystemSettingDNSOutput() SystemSettingDNSOutput
	ToSystemSettingDNSOutputWithContext(ctx context.Context) SystemSettingDNSOutput
}

func (*SystemSettingDNS) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSettingDNS)(nil)).Elem()
}

func (i *SystemSettingDNS) ToSystemSettingDNSOutput() SystemSettingDNSOutput {
	return i.ToSystemSettingDNSOutputWithContext(context.Background())
}

func (i *SystemSettingDNS) ToSystemSettingDNSOutputWithContext(ctx context.Context) SystemSettingDNSOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSettingDNSOutput)
}

// SystemSettingDNSArrayInput is an input type that accepts SystemSettingDNSArray and SystemSettingDNSArrayOutput values.
// You can construct a concrete instance of `SystemSettingDNSArrayInput` via:
//
//	SystemSettingDNSArray{ SystemSettingDNSArgs{...} }
type SystemSettingDNSArrayInput interface {
	pulumi.Input

	ToSystemSettingDNSArrayOutput() SystemSettingDNSArrayOutput
	ToSystemSettingDNSArrayOutputWithContext(context.Context) SystemSettingDNSArrayOutput
}

type SystemSettingDNSArray []SystemSettingDNSInput

func (SystemSettingDNSArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSettingDNS)(nil)).Elem()
}

func (i SystemSettingDNSArray) ToSystemSettingDNSArrayOutput() SystemSettingDNSArrayOutput {
	return i.ToSystemSettingDNSArrayOutputWithContext(context.Background())
}

func (i SystemSettingDNSArray) ToSystemSettingDNSArrayOutputWithContext(ctx context.Context) SystemSettingDNSArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSettingDNSArrayOutput)
}

// SystemSettingDNSMapInput is an input type that accepts SystemSettingDNSMap and SystemSettingDNSMapOutput values.
// You can construct a concrete instance of `SystemSettingDNSMapInput` via:
//
//	SystemSettingDNSMap{ "key": SystemSettingDNSArgs{...} }
type SystemSettingDNSMapInput interface {
	pulumi.Input

	ToSystemSettingDNSMapOutput() SystemSettingDNSMapOutput
	ToSystemSettingDNSMapOutputWithContext(context.Context) SystemSettingDNSMapOutput
}

type SystemSettingDNSMap map[string]SystemSettingDNSInput

func (SystemSettingDNSMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSettingDNS)(nil)).Elem()
}

func (i SystemSettingDNSMap) ToSystemSettingDNSMapOutput() SystemSettingDNSMapOutput {
	return i.ToSystemSettingDNSMapOutputWithContext(context.Background())
}

func (i SystemSettingDNSMap) ToSystemSettingDNSMapOutputWithContext(ctx context.Context) SystemSettingDNSMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSettingDNSMapOutput)
}

type SystemSettingDNSOutput struct{ *pulumi.OutputState }

func (SystemSettingDNSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSettingDNS)(nil)).Elem()
}

func (o SystemSettingDNSOutput) ToSystemSettingDNSOutput() SystemSettingDNSOutput {
	return o
}

func (o SystemSettingDNSOutput) ToSystemSettingDNSOutputWithContext(ctx context.Context) SystemSettingDNSOutput {
	return o
}

func (o SystemSettingDNSOutput) DnsOverTls() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSettingDNS) pulumi.StringOutput { return v.DnsOverTls }).(pulumi.StringOutput)
}

func (o SystemSettingDNSOutput) Primary() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSettingDNS) pulumi.StringOutput { return v.Primary }).(pulumi.StringOutput)
}

func (o SystemSettingDNSOutput) Secondary() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSettingDNS) pulumi.StringOutput { return v.Secondary }).(pulumi.StringOutput)
}

type SystemSettingDNSArrayOutput struct{ *pulumi.OutputState }

func (SystemSettingDNSArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSettingDNS)(nil)).Elem()
}

func (o SystemSettingDNSArrayOutput) ToSystemSettingDNSArrayOutput() SystemSettingDNSArrayOutput {
	return o
}

func (o SystemSettingDNSArrayOutput) ToSystemSettingDNSArrayOutputWithContext(ctx context.Context) SystemSettingDNSArrayOutput {
	return o
}

func (o SystemSettingDNSArrayOutput) Index(i pulumi.IntInput) SystemSettingDNSOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSettingDNS {
		return vs[0].([]*SystemSettingDNS)[vs[1].(int)]
	}).(SystemSettingDNSOutput)
}

type SystemSettingDNSMapOutput struct{ *pulumi.OutputState }

func (SystemSettingDNSMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSettingDNS)(nil)).Elem()
}

func (o SystemSettingDNSMapOutput) ToSystemSettingDNSMapOutput() SystemSettingDNSMapOutput {
	return o
}

func (o SystemSettingDNSMapOutput) ToSystemSettingDNSMapOutputWithContext(ctx context.Context) SystemSettingDNSMapOutput {
	return o
}

func (o SystemSettingDNSMapOutput) MapIndex(k pulumi.StringInput) SystemSettingDNSOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSettingDNS {
		return vs[0].(map[string]*SystemSettingDNS)[vs[1].(string)]
	}).(SystemSettingDNSOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSettingDNSInput)(nil)).Elem(), &SystemSettingDNS{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSettingDNSArrayInput)(nil)).Elem(), SystemSettingDNSArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSettingDNSMapInput)(nil)).Elem(), SystemSettingDNSMap{})
	pulumi.RegisterOutputType(SystemSettingDNSOutput{})
	pulumi.RegisterOutputType(SystemSettingDNSArrayOutput{})
	pulumi.RegisterOutputType(SystemSettingDNSMapOutput{})
}

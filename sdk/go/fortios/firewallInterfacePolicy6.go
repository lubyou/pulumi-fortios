// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallInterfacePolicy6 struct {
	pulumi.CustomResourceState

	AddressType              pulumi.StringOutput                         `pulumi:"addressType"`
	ApplicationList          pulumi.StringOutput                         `pulumi:"applicationList"`
	ApplicationListStatus    pulumi.StringOutput                         `pulumi:"applicationListStatus"`
	AvProfile                pulumi.StringOutput                         `pulumi:"avProfile"`
	AvProfileStatus          pulumi.StringOutput                         `pulumi:"avProfileStatus"`
	Comments                 pulumi.StringPtrOutput                      `pulumi:"comments"`
	DlpProfile               pulumi.StringOutput                         `pulumi:"dlpProfile"`
	DlpProfileStatus         pulumi.StringOutput                         `pulumi:"dlpProfileStatus"`
	DlpSensor                pulumi.StringOutput                         `pulumi:"dlpSensor"`
	DlpSensorStatus          pulumi.StringOutput                         `pulumi:"dlpSensorStatus"`
	Dsri                     pulumi.StringOutput                         `pulumi:"dsri"`
	Dstaddr6s                FirewallInterfacePolicy6Dstaddr6ArrayOutput `pulumi:"dstaddr6s"`
	DynamicSortSubtable      pulumi.StringPtrOutput                      `pulumi:"dynamicSortSubtable"`
	EmailfilterProfile       pulumi.StringOutput                         `pulumi:"emailfilterProfile"`
	EmailfilterProfileStatus pulumi.StringOutput                         `pulumi:"emailfilterProfileStatus"`
	Interface                pulumi.StringOutput                         `pulumi:"interface"`
	IpsSensor                pulumi.StringOutput                         `pulumi:"ipsSensor"`
	IpsSensorStatus          pulumi.StringOutput                         `pulumi:"ipsSensorStatus"`
	Label                    pulumi.StringOutput                         `pulumi:"label"`
	Logtraffic               pulumi.StringOutput                         `pulumi:"logtraffic"`
	Policyid                 pulumi.IntOutput                            `pulumi:"policyid"`
	ScanBotnetConnections    pulumi.StringOutput                         `pulumi:"scanBotnetConnections"`
	Service6s                FirewallInterfacePolicy6Service6ArrayOutput `pulumi:"service6s"`
	SpamfilterProfile        pulumi.StringOutput                         `pulumi:"spamfilterProfile"`
	SpamfilterProfileStatus  pulumi.StringOutput                         `pulumi:"spamfilterProfileStatus"`
	Srcaddr6s                FirewallInterfacePolicy6Srcaddr6ArrayOutput `pulumi:"srcaddr6s"`
	Status                   pulumi.StringOutput                         `pulumi:"status"`
	Vdomparam                pulumi.StringPtrOutput                      `pulumi:"vdomparam"`
	WebfilterProfile         pulumi.StringOutput                         `pulumi:"webfilterProfile"`
	WebfilterProfileStatus   pulumi.StringOutput                         `pulumi:"webfilterProfileStatus"`
}

// NewFirewallInterfacePolicy6 registers a new resource with the given unique name, arguments, and options.
func NewFirewallInterfacePolicy6(ctx *pulumi.Context,
	name string, args *FirewallInterfacePolicy6Args, opts ...pulumi.ResourceOption) (*FirewallInterfacePolicy6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dstaddr6s == nil {
		return nil, errors.New("invalid value for required argument 'Dstaddr6s'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Srcaddr6s == nil {
		return nil, errors.New("invalid value for required argument 'Srcaddr6s'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInterfacePolicy6
	err := ctx.RegisterResource("fortios:index/firewallInterfacePolicy6:FirewallInterfacePolicy6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInterfacePolicy6 gets an existing FirewallInterfacePolicy6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInterfacePolicy6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInterfacePolicy6State, opts ...pulumi.ResourceOption) (*FirewallInterfacePolicy6, error) {
	var resource FirewallInterfacePolicy6
	err := ctx.ReadResource("fortios:index/firewallInterfacePolicy6:FirewallInterfacePolicy6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInterfacePolicy6 resources.
type firewallInterfacePolicy6State struct {
	AddressType              *string                            `pulumi:"addressType"`
	ApplicationList          *string                            `pulumi:"applicationList"`
	ApplicationListStatus    *string                            `pulumi:"applicationListStatus"`
	AvProfile                *string                            `pulumi:"avProfile"`
	AvProfileStatus          *string                            `pulumi:"avProfileStatus"`
	Comments                 *string                            `pulumi:"comments"`
	DlpProfile               *string                            `pulumi:"dlpProfile"`
	DlpProfileStatus         *string                            `pulumi:"dlpProfileStatus"`
	DlpSensor                *string                            `pulumi:"dlpSensor"`
	DlpSensorStatus          *string                            `pulumi:"dlpSensorStatus"`
	Dsri                     *string                            `pulumi:"dsri"`
	Dstaddr6s                []FirewallInterfacePolicy6Dstaddr6 `pulumi:"dstaddr6s"`
	DynamicSortSubtable      *string                            `pulumi:"dynamicSortSubtable"`
	EmailfilterProfile       *string                            `pulumi:"emailfilterProfile"`
	EmailfilterProfileStatus *string                            `pulumi:"emailfilterProfileStatus"`
	Interface                *string                            `pulumi:"interface"`
	IpsSensor                *string                            `pulumi:"ipsSensor"`
	IpsSensorStatus          *string                            `pulumi:"ipsSensorStatus"`
	Label                    *string                            `pulumi:"label"`
	Logtraffic               *string                            `pulumi:"logtraffic"`
	Policyid                 *int                               `pulumi:"policyid"`
	ScanBotnetConnections    *string                            `pulumi:"scanBotnetConnections"`
	Service6s                []FirewallInterfacePolicy6Service6 `pulumi:"service6s"`
	SpamfilterProfile        *string                            `pulumi:"spamfilterProfile"`
	SpamfilterProfileStatus  *string                            `pulumi:"spamfilterProfileStatus"`
	Srcaddr6s                []FirewallInterfacePolicy6Srcaddr6 `pulumi:"srcaddr6s"`
	Status                   *string                            `pulumi:"status"`
	Vdomparam                *string                            `pulumi:"vdomparam"`
	WebfilterProfile         *string                            `pulumi:"webfilterProfile"`
	WebfilterProfileStatus   *string                            `pulumi:"webfilterProfileStatus"`
}

type FirewallInterfacePolicy6State struct {
	AddressType              pulumi.StringPtrInput
	ApplicationList          pulumi.StringPtrInput
	ApplicationListStatus    pulumi.StringPtrInput
	AvProfile                pulumi.StringPtrInput
	AvProfileStatus          pulumi.StringPtrInput
	Comments                 pulumi.StringPtrInput
	DlpProfile               pulumi.StringPtrInput
	DlpProfileStatus         pulumi.StringPtrInput
	DlpSensor                pulumi.StringPtrInput
	DlpSensorStatus          pulumi.StringPtrInput
	Dsri                     pulumi.StringPtrInput
	Dstaddr6s                FirewallInterfacePolicy6Dstaddr6ArrayInput
	DynamicSortSubtable      pulumi.StringPtrInput
	EmailfilterProfile       pulumi.StringPtrInput
	EmailfilterProfileStatus pulumi.StringPtrInput
	Interface                pulumi.StringPtrInput
	IpsSensor                pulumi.StringPtrInput
	IpsSensorStatus          pulumi.StringPtrInput
	Label                    pulumi.StringPtrInput
	Logtraffic               pulumi.StringPtrInput
	Policyid                 pulumi.IntPtrInput
	ScanBotnetConnections    pulumi.StringPtrInput
	Service6s                FirewallInterfacePolicy6Service6ArrayInput
	SpamfilterProfile        pulumi.StringPtrInput
	SpamfilterProfileStatus  pulumi.StringPtrInput
	Srcaddr6s                FirewallInterfacePolicy6Srcaddr6ArrayInput
	Status                   pulumi.StringPtrInput
	Vdomparam                pulumi.StringPtrInput
	WebfilterProfile         pulumi.StringPtrInput
	WebfilterProfileStatus   pulumi.StringPtrInput
}

func (FirewallInterfacePolicy6State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInterfacePolicy6State)(nil)).Elem()
}

type firewallInterfacePolicy6Args struct {
	AddressType              *string                            `pulumi:"addressType"`
	ApplicationList          *string                            `pulumi:"applicationList"`
	ApplicationListStatus    *string                            `pulumi:"applicationListStatus"`
	AvProfile                *string                            `pulumi:"avProfile"`
	AvProfileStatus          *string                            `pulumi:"avProfileStatus"`
	Comments                 *string                            `pulumi:"comments"`
	DlpProfile               *string                            `pulumi:"dlpProfile"`
	DlpProfileStatus         *string                            `pulumi:"dlpProfileStatus"`
	DlpSensor                *string                            `pulumi:"dlpSensor"`
	DlpSensorStatus          *string                            `pulumi:"dlpSensorStatus"`
	Dsri                     *string                            `pulumi:"dsri"`
	Dstaddr6s                []FirewallInterfacePolicy6Dstaddr6 `pulumi:"dstaddr6s"`
	DynamicSortSubtable      *string                            `pulumi:"dynamicSortSubtable"`
	EmailfilterProfile       *string                            `pulumi:"emailfilterProfile"`
	EmailfilterProfileStatus *string                            `pulumi:"emailfilterProfileStatus"`
	Interface                string                             `pulumi:"interface"`
	IpsSensor                *string                            `pulumi:"ipsSensor"`
	IpsSensorStatus          *string                            `pulumi:"ipsSensorStatus"`
	Label                    *string                            `pulumi:"label"`
	Logtraffic               *string                            `pulumi:"logtraffic"`
	Policyid                 *int                               `pulumi:"policyid"`
	ScanBotnetConnections    *string                            `pulumi:"scanBotnetConnections"`
	Service6s                []FirewallInterfacePolicy6Service6 `pulumi:"service6s"`
	SpamfilterProfile        *string                            `pulumi:"spamfilterProfile"`
	SpamfilterProfileStatus  *string                            `pulumi:"spamfilterProfileStatus"`
	Srcaddr6s                []FirewallInterfacePolicy6Srcaddr6 `pulumi:"srcaddr6s"`
	Status                   *string                            `pulumi:"status"`
	Vdomparam                *string                            `pulumi:"vdomparam"`
	WebfilterProfile         *string                            `pulumi:"webfilterProfile"`
	WebfilterProfileStatus   *string                            `pulumi:"webfilterProfileStatus"`
}

// The set of arguments for constructing a FirewallInterfacePolicy6 resource.
type FirewallInterfacePolicy6Args struct {
	AddressType              pulumi.StringPtrInput
	ApplicationList          pulumi.StringPtrInput
	ApplicationListStatus    pulumi.StringPtrInput
	AvProfile                pulumi.StringPtrInput
	AvProfileStatus          pulumi.StringPtrInput
	Comments                 pulumi.StringPtrInput
	DlpProfile               pulumi.StringPtrInput
	DlpProfileStatus         pulumi.StringPtrInput
	DlpSensor                pulumi.StringPtrInput
	DlpSensorStatus          pulumi.StringPtrInput
	Dsri                     pulumi.StringPtrInput
	Dstaddr6s                FirewallInterfacePolicy6Dstaddr6ArrayInput
	DynamicSortSubtable      pulumi.StringPtrInput
	EmailfilterProfile       pulumi.StringPtrInput
	EmailfilterProfileStatus pulumi.StringPtrInput
	Interface                pulumi.StringInput
	IpsSensor                pulumi.StringPtrInput
	IpsSensorStatus          pulumi.StringPtrInput
	Label                    pulumi.StringPtrInput
	Logtraffic               pulumi.StringPtrInput
	Policyid                 pulumi.IntPtrInput
	ScanBotnetConnections    pulumi.StringPtrInput
	Service6s                FirewallInterfacePolicy6Service6ArrayInput
	SpamfilterProfile        pulumi.StringPtrInput
	SpamfilterProfileStatus  pulumi.StringPtrInput
	Srcaddr6s                FirewallInterfacePolicy6Srcaddr6ArrayInput
	Status                   pulumi.StringPtrInput
	Vdomparam                pulumi.StringPtrInput
	WebfilterProfile         pulumi.StringPtrInput
	WebfilterProfileStatus   pulumi.StringPtrInput
}

func (FirewallInterfacePolicy6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInterfacePolicy6Args)(nil)).Elem()
}

type FirewallInterfacePolicy6Input interface {
	pulumi.Input

	ToFirewallInterfacePolicy6Output() FirewallInterfacePolicy6Output
	ToFirewallInterfacePolicy6OutputWithContext(ctx context.Context) FirewallInterfacePolicy6Output
}

func (*FirewallInterfacePolicy6) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInterfacePolicy6)(nil)).Elem()
}

func (i *FirewallInterfacePolicy6) ToFirewallInterfacePolicy6Output() FirewallInterfacePolicy6Output {
	return i.ToFirewallInterfacePolicy6OutputWithContext(context.Background())
}

func (i *FirewallInterfacePolicy6) ToFirewallInterfacePolicy6OutputWithContext(ctx context.Context) FirewallInterfacePolicy6Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInterfacePolicy6Output)
}

// FirewallInterfacePolicy6ArrayInput is an input type that accepts FirewallInterfacePolicy6Array and FirewallInterfacePolicy6ArrayOutput values.
// You can construct a concrete instance of `FirewallInterfacePolicy6ArrayInput` via:
//
//	FirewallInterfacePolicy6Array{ FirewallInterfacePolicy6Args{...} }
type FirewallInterfacePolicy6ArrayInput interface {
	pulumi.Input

	ToFirewallInterfacePolicy6ArrayOutput() FirewallInterfacePolicy6ArrayOutput
	ToFirewallInterfacePolicy6ArrayOutputWithContext(context.Context) FirewallInterfacePolicy6ArrayOutput
}

type FirewallInterfacePolicy6Array []FirewallInterfacePolicy6Input

func (FirewallInterfacePolicy6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInterfacePolicy6)(nil)).Elem()
}

func (i FirewallInterfacePolicy6Array) ToFirewallInterfacePolicy6ArrayOutput() FirewallInterfacePolicy6ArrayOutput {
	return i.ToFirewallInterfacePolicy6ArrayOutputWithContext(context.Background())
}

func (i FirewallInterfacePolicy6Array) ToFirewallInterfacePolicy6ArrayOutputWithContext(ctx context.Context) FirewallInterfacePolicy6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInterfacePolicy6ArrayOutput)
}

// FirewallInterfacePolicy6MapInput is an input type that accepts FirewallInterfacePolicy6Map and FirewallInterfacePolicy6MapOutput values.
// You can construct a concrete instance of `FirewallInterfacePolicy6MapInput` via:
//
//	FirewallInterfacePolicy6Map{ "key": FirewallInterfacePolicy6Args{...} }
type FirewallInterfacePolicy6MapInput interface {
	pulumi.Input

	ToFirewallInterfacePolicy6MapOutput() FirewallInterfacePolicy6MapOutput
	ToFirewallInterfacePolicy6MapOutputWithContext(context.Context) FirewallInterfacePolicy6MapOutput
}

type FirewallInterfacePolicy6Map map[string]FirewallInterfacePolicy6Input

func (FirewallInterfacePolicy6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInterfacePolicy6)(nil)).Elem()
}

func (i FirewallInterfacePolicy6Map) ToFirewallInterfacePolicy6MapOutput() FirewallInterfacePolicy6MapOutput {
	return i.ToFirewallInterfacePolicy6MapOutputWithContext(context.Background())
}

func (i FirewallInterfacePolicy6Map) ToFirewallInterfacePolicy6MapOutputWithContext(ctx context.Context) FirewallInterfacePolicy6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInterfacePolicy6MapOutput)
}

type FirewallInterfacePolicy6Output struct{ *pulumi.OutputState }

func (FirewallInterfacePolicy6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInterfacePolicy6)(nil)).Elem()
}

func (o FirewallInterfacePolicy6Output) ToFirewallInterfacePolicy6Output() FirewallInterfacePolicy6Output {
	return o
}

func (o FirewallInterfacePolicy6Output) ToFirewallInterfacePolicy6OutputWithContext(ctx context.Context) FirewallInterfacePolicy6Output {
	return o
}

func (o FirewallInterfacePolicy6Output) AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.AddressType }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) ApplicationList() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.ApplicationList }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) ApplicationListStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.ApplicationListStatus }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) AvProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.AvProfile }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) AvProfileStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.AvProfileStatus }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o FirewallInterfacePolicy6Output) DlpProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.DlpProfile }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) DlpProfileStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.DlpProfileStatus }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) DlpSensor() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.DlpSensor }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) DlpSensorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.DlpSensorStatus }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) Dsri() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.Dsri }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) Dstaddr6s() FirewallInterfacePolicy6Dstaddr6ArrayOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) FirewallInterfacePolicy6Dstaddr6ArrayOutput { return v.Dstaddr6s }).(FirewallInterfacePolicy6Dstaddr6ArrayOutput)
}

func (o FirewallInterfacePolicy6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallInterfacePolicy6Output) EmailfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.EmailfilterProfile }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) EmailfilterProfileStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.EmailfilterProfileStatus }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) IpsSensor() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.IpsSensor }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) IpsSensorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.IpsSensorStatus }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) Label() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.Label }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) Logtraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.Logtraffic }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.IntOutput { return v.Policyid }).(pulumi.IntOutput)
}

func (o FirewallInterfacePolicy6Output) ScanBotnetConnections() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.ScanBotnetConnections }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) Service6s() FirewallInterfacePolicy6Service6ArrayOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) FirewallInterfacePolicy6Service6ArrayOutput { return v.Service6s }).(FirewallInterfacePolicy6Service6ArrayOutput)
}

func (o FirewallInterfacePolicy6Output) SpamfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.SpamfilterProfile }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) SpamfilterProfileStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.SpamfilterProfileStatus }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) Srcaddr6s() FirewallInterfacePolicy6Srcaddr6ArrayOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) FirewallInterfacePolicy6Srcaddr6ArrayOutput { return v.Srcaddr6s }).(FirewallInterfacePolicy6Srcaddr6ArrayOutput)
}

func (o FirewallInterfacePolicy6Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o FirewallInterfacePolicy6Output) WebfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.WebfilterProfile }).(pulumi.StringOutput)
}

func (o FirewallInterfacePolicy6Output) WebfilterProfileStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInterfacePolicy6) pulumi.StringOutput { return v.WebfilterProfileStatus }).(pulumi.StringOutput)
}

type FirewallInterfacePolicy6ArrayOutput struct{ *pulumi.OutputState }

func (FirewallInterfacePolicy6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInterfacePolicy6)(nil)).Elem()
}

func (o FirewallInterfacePolicy6ArrayOutput) ToFirewallInterfacePolicy6ArrayOutput() FirewallInterfacePolicy6ArrayOutput {
	return o
}

func (o FirewallInterfacePolicy6ArrayOutput) ToFirewallInterfacePolicy6ArrayOutputWithContext(ctx context.Context) FirewallInterfacePolicy6ArrayOutput {
	return o
}

func (o FirewallInterfacePolicy6ArrayOutput) Index(i pulumi.IntInput) FirewallInterfacePolicy6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInterfacePolicy6 {
		return vs[0].([]*FirewallInterfacePolicy6)[vs[1].(int)]
	}).(FirewallInterfacePolicy6Output)
}

type FirewallInterfacePolicy6MapOutput struct{ *pulumi.OutputState }

func (FirewallInterfacePolicy6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInterfacePolicy6)(nil)).Elem()
}

func (o FirewallInterfacePolicy6MapOutput) ToFirewallInterfacePolicy6MapOutput() FirewallInterfacePolicy6MapOutput {
	return o
}

func (o FirewallInterfacePolicy6MapOutput) ToFirewallInterfacePolicy6MapOutputWithContext(ctx context.Context) FirewallInterfacePolicy6MapOutput {
	return o
}

func (o FirewallInterfacePolicy6MapOutput) MapIndex(k pulumi.StringInput) FirewallInterfacePolicy6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInterfacePolicy6 {
		return vs[0].(map[string]*FirewallInterfacePolicy6)[vs[1].(string)]
	}).(FirewallInterfacePolicy6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInterfacePolicy6Input)(nil)).Elem(), &FirewallInterfacePolicy6{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInterfacePolicy6ArrayInput)(nil)).Elem(), FirewallInterfacePolicy6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInterfacePolicy6MapInput)(nil)).Elem(), FirewallInterfacePolicy6Map{})
	pulumi.RegisterOutputType(FirewallInterfacePolicy6Output{})
	pulumi.RegisterOutputType(FirewallInterfacePolicy6ArrayOutput{})
	pulumi.RegisterOutputType(FirewallInterfacePolicy6MapOutput{})
}

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

type FirewallSecurityPolicy1 struct {
	pulumi.CustomResourceState

	Action                 pulumi.StringOutput      `pulumi:"action"`
	ApplicationList        pulumi.StringOutput      `pulumi:"applicationList"`
	AvProfile              pulumi.StringOutput      `pulumi:"avProfile"`
	CapturePacket          pulumi.StringOutput      `pulumi:"capturePacket"`
	Comments               pulumi.StringPtrOutput   `pulumi:"comments"`
	Devices                pulumi.StringArrayOutput `pulumi:"devices"`
	DnsfilterProfile       pulumi.StringOutput      `pulumi:"dnsfilterProfile"`
	Dstaddrs               pulumi.StringArrayOutput `pulumi:"dstaddrs"`
	Dstintfs               pulumi.StringArrayOutput `pulumi:"dstintfs"`
	Groups                 pulumi.StringArrayOutput `pulumi:"groups"`
	InternetService        pulumi.StringOutput      `pulumi:"internetService"`
	InternetServiceIds     pulumi.IntArrayOutput    `pulumi:"internetServiceIds"`
	InternetServiceSrc     pulumi.StringOutput      `pulumi:"internetServiceSrc"`
	InternetServiceSrcIds  pulumi.IntArrayOutput    `pulumi:"internetServiceSrcIds"`
	Ippool                 pulumi.StringOutput      `pulumi:"ippool"`
	IpsSensor              pulumi.StringOutput      `pulumi:"ipsSensor"`
	Logtraffic             pulumi.StringOutput      `pulumi:"logtraffic"`
	LogtrafficStart        pulumi.StringOutput      `pulumi:"logtrafficStart"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	Nat                    pulumi.StringOutput      `pulumi:"nat"`
	Poolnames              pulumi.StringArrayOutput `pulumi:"poolnames"`
	ProfileProtocolOptions pulumi.StringOutput      `pulumi:"profileProtocolOptions"`
	Schedule               pulumi.StringOutput      `pulumi:"schedule"`
	Services               pulumi.StringArrayOutput `pulumi:"services"`
	Srcaddrs               pulumi.StringArrayOutput `pulumi:"srcaddrs"`
	Srcintfs               pulumi.StringArrayOutput `pulumi:"srcintfs"`
	SslSshProfile          pulumi.StringOutput      `pulumi:"sslSshProfile"`
	Status                 pulumi.StringOutput      `pulumi:"status"`
	Users                  pulumi.StringArrayOutput `pulumi:"users"`
	UtmStatus              pulumi.StringOutput      `pulumi:"utmStatus"`
	WebfilterProfile       pulumi.StringOutput      `pulumi:"webfilterProfile"`
}

// NewFirewallSecurityPolicy1 registers a new resource with the given unique name, arguments, and options.
func NewFirewallSecurityPolicy1(ctx *pulumi.Context,
	name string, args *FirewallSecurityPolicy1Args, opts ...pulumi.ResourceOption) (*FirewallSecurityPolicy1, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Dstaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Dstaddrs'")
	}
	if args.Dstintfs == nil {
		return nil, errors.New("invalid value for required argument 'Dstintfs'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.Services == nil {
		return nil, errors.New("invalid value for required argument 'Services'")
	}
	if args.Srcaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Srcaddrs'")
	}
	if args.Srcintfs == nil {
		return nil, errors.New("invalid value for required argument 'Srcintfs'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallSecurityPolicy1
	err := ctx.RegisterResource("fortios:index/firewallSecurityPolicy1:FirewallSecurityPolicy1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSecurityPolicy1 gets an existing FirewallSecurityPolicy1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSecurityPolicy1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSecurityPolicy1State, opts ...pulumi.ResourceOption) (*FirewallSecurityPolicy1, error) {
	var resource FirewallSecurityPolicy1
	err := ctx.ReadResource("fortios:index/firewallSecurityPolicy1:FirewallSecurityPolicy1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSecurityPolicy1 resources.
type firewallSecurityPolicy1State struct {
	Action                 *string  `pulumi:"action"`
	ApplicationList        *string  `pulumi:"applicationList"`
	AvProfile              *string  `pulumi:"avProfile"`
	CapturePacket          *string  `pulumi:"capturePacket"`
	Comments               *string  `pulumi:"comments"`
	Devices                []string `pulumi:"devices"`
	DnsfilterProfile       *string  `pulumi:"dnsfilterProfile"`
	Dstaddrs               []string `pulumi:"dstaddrs"`
	Dstintfs               []string `pulumi:"dstintfs"`
	Groups                 []string `pulumi:"groups"`
	InternetService        *string  `pulumi:"internetService"`
	InternetServiceIds     []int    `pulumi:"internetServiceIds"`
	InternetServiceSrc     *string  `pulumi:"internetServiceSrc"`
	InternetServiceSrcIds  []int    `pulumi:"internetServiceSrcIds"`
	Ippool                 *string  `pulumi:"ippool"`
	IpsSensor              *string  `pulumi:"ipsSensor"`
	Logtraffic             *string  `pulumi:"logtraffic"`
	LogtrafficStart        *string  `pulumi:"logtrafficStart"`
	Name                   *string  `pulumi:"name"`
	Nat                    *string  `pulumi:"nat"`
	Poolnames              []string `pulumi:"poolnames"`
	ProfileProtocolOptions *string  `pulumi:"profileProtocolOptions"`
	Schedule               *string  `pulumi:"schedule"`
	Services               []string `pulumi:"services"`
	Srcaddrs               []string `pulumi:"srcaddrs"`
	Srcintfs               []string `pulumi:"srcintfs"`
	SslSshProfile          *string  `pulumi:"sslSshProfile"`
	Status                 *string  `pulumi:"status"`
	Users                  []string `pulumi:"users"`
	UtmStatus              *string  `pulumi:"utmStatus"`
	WebfilterProfile       *string  `pulumi:"webfilterProfile"`
}

type FirewallSecurityPolicy1State struct {
	Action                 pulumi.StringPtrInput
	ApplicationList        pulumi.StringPtrInput
	AvProfile              pulumi.StringPtrInput
	CapturePacket          pulumi.StringPtrInput
	Comments               pulumi.StringPtrInput
	Devices                pulumi.StringArrayInput
	DnsfilterProfile       pulumi.StringPtrInput
	Dstaddrs               pulumi.StringArrayInput
	Dstintfs               pulumi.StringArrayInput
	Groups                 pulumi.StringArrayInput
	InternetService        pulumi.StringPtrInput
	InternetServiceIds     pulumi.IntArrayInput
	InternetServiceSrc     pulumi.StringPtrInput
	InternetServiceSrcIds  pulumi.IntArrayInput
	Ippool                 pulumi.StringPtrInput
	IpsSensor              pulumi.StringPtrInput
	Logtraffic             pulumi.StringPtrInput
	LogtrafficStart        pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	Nat                    pulumi.StringPtrInput
	Poolnames              pulumi.StringArrayInput
	ProfileProtocolOptions pulumi.StringPtrInput
	Schedule               pulumi.StringPtrInput
	Services               pulumi.StringArrayInput
	Srcaddrs               pulumi.StringArrayInput
	Srcintfs               pulumi.StringArrayInput
	SslSshProfile          pulumi.StringPtrInput
	Status                 pulumi.StringPtrInput
	Users                  pulumi.StringArrayInput
	UtmStatus              pulumi.StringPtrInput
	WebfilterProfile       pulumi.StringPtrInput
}

func (FirewallSecurityPolicy1State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSecurityPolicy1State)(nil)).Elem()
}

type firewallSecurityPolicy1Args struct {
	Action                 string   `pulumi:"action"`
	ApplicationList        *string  `pulumi:"applicationList"`
	AvProfile              *string  `pulumi:"avProfile"`
	CapturePacket          *string  `pulumi:"capturePacket"`
	Comments               *string  `pulumi:"comments"`
	Devices                []string `pulumi:"devices"`
	DnsfilterProfile       *string  `pulumi:"dnsfilterProfile"`
	Dstaddrs               []string `pulumi:"dstaddrs"`
	Dstintfs               []string `pulumi:"dstintfs"`
	Groups                 []string `pulumi:"groups"`
	InternetService        *string  `pulumi:"internetService"`
	InternetServiceIds     []int    `pulumi:"internetServiceIds"`
	InternetServiceSrc     *string  `pulumi:"internetServiceSrc"`
	InternetServiceSrcIds  []int    `pulumi:"internetServiceSrcIds"`
	Ippool                 *string  `pulumi:"ippool"`
	IpsSensor              *string  `pulumi:"ipsSensor"`
	Logtraffic             *string  `pulumi:"logtraffic"`
	LogtrafficStart        *string  `pulumi:"logtrafficStart"`
	Name                   *string  `pulumi:"name"`
	Nat                    *string  `pulumi:"nat"`
	Poolnames              []string `pulumi:"poolnames"`
	ProfileProtocolOptions *string  `pulumi:"profileProtocolOptions"`
	Schedule               string   `pulumi:"schedule"`
	Services               []string `pulumi:"services"`
	Srcaddrs               []string `pulumi:"srcaddrs"`
	Srcintfs               []string `pulumi:"srcintfs"`
	SslSshProfile          *string  `pulumi:"sslSshProfile"`
	Status                 *string  `pulumi:"status"`
	Users                  []string `pulumi:"users"`
	UtmStatus              *string  `pulumi:"utmStatus"`
	WebfilterProfile       *string  `pulumi:"webfilterProfile"`
}

// The set of arguments for constructing a FirewallSecurityPolicy1 resource.
type FirewallSecurityPolicy1Args struct {
	Action                 pulumi.StringInput
	ApplicationList        pulumi.StringPtrInput
	AvProfile              pulumi.StringPtrInput
	CapturePacket          pulumi.StringPtrInput
	Comments               pulumi.StringPtrInput
	Devices                pulumi.StringArrayInput
	DnsfilterProfile       pulumi.StringPtrInput
	Dstaddrs               pulumi.StringArrayInput
	Dstintfs               pulumi.StringArrayInput
	Groups                 pulumi.StringArrayInput
	InternetService        pulumi.StringPtrInput
	InternetServiceIds     pulumi.IntArrayInput
	InternetServiceSrc     pulumi.StringPtrInput
	InternetServiceSrcIds  pulumi.IntArrayInput
	Ippool                 pulumi.StringPtrInput
	IpsSensor              pulumi.StringPtrInput
	Logtraffic             pulumi.StringPtrInput
	LogtrafficStart        pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	Nat                    pulumi.StringPtrInput
	Poolnames              pulumi.StringArrayInput
	ProfileProtocolOptions pulumi.StringPtrInput
	Schedule               pulumi.StringInput
	Services               pulumi.StringArrayInput
	Srcaddrs               pulumi.StringArrayInput
	Srcintfs               pulumi.StringArrayInput
	SslSshProfile          pulumi.StringPtrInput
	Status                 pulumi.StringPtrInput
	Users                  pulumi.StringArrayInput
	UtmStatus              pulumi.StringPtrInput
	WebfilterProfile       pulumi.StringPtrInput
}

func (FirewallSecurityPolicy1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSecurityPolicy1Args)(nil)).Elem()
}

type FirewallSecurityPolicy1Input interface {
	pulumi.Input

	ToFirewallSecurityPolicy1Output() FirewallSecurityPolicy1Output
	ToFirewallSecurityPolicy1OutputWithContext(ctx context.Context) FirewallSecurityPolicy1Output
}

func (*FirewallSecurityPolicy1) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSecurityPolicy1)(nil)).Elem()
}

func (i *FirewallSecurityPolicy1) ToFirewallSecurityPolicy1Output() FirewallSecurityPolicy1Output {
	return i.ToFirewallSecurityPolicy1OutputWithContext(context.Background())
}

func (i *FirewallSecurityPolicy1) ToFirewallSecurityPolicy1OutputWithContext(ctx context.Context) FirewallSecurityPolicy1Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicy1Output)
}

// FirewallSecurityPolicy1ArrayInput is an input type that accepts FirewallSecurityPolicy1Array and FirewallSecurityPolicy1ArrayOutput values.
// You can construct a concrete instance of `FirewallSecurityPolicy1ArrayInput` via:
//
//	FirewallSecurityPolicy1Array{ FirewallSecurityPolicy1Args{...} }
type FirewallSecurityPolicy1ArrayInput interface {
	pulumi.Input

	ToFirewallSecurityPolicy1ArrayOutput() FirewallSecurityPolicy1ArrayOutput
	ToFirewallSecurityPolicy1ArrayOutputWithContext(context.Context) FirewallSecurityPolicy1ArrayOutput
}

type FirewallSecurityPolicy1Array []FirewallSecurityPolicy1Input

func (FirewallSecurityPolicy1Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSecurityPolicy1)(nil)).Elem()
}

func (i FirewallSecurityPolicy1Array) ToFirewallSecurityPolicy1ArrayOutput() FirewallSecurityPolicy1ArrayOutput {
	return i.ToFirewallSecurityPolicy1ArrayOutputWithContext(context.Background())
}

func (i FirewallSecurityPolicy1Array) ToFirewallSecurityPolicy1ArrayOutputWithContext(ctx context.Context) FirewallSecurityPolicy1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicy1ArrayOutput)
}

// FirewallSecurityPolicy1MapInput is an input type that accepts FirewallSecurityPolicy1Map and FirewallSecurityPolicy1MapOutput values.
// You can construct a concrete instance of `FirewallSecurityPolicy1MapInput` via:
//
//	FirewallSecurityPolicy1Map{ "key": FirewallSecurityPolicy1Args{...} }
type FirewallSecurityPolicy1MapInput interface {
	pulumi.Input

	ToFirewallSecurityPolicy1MapOutput() FirewallSecurityPolicy1MapOutput
	ToFirewallSecurityPolicy1MapOutputWithContext(context.Context) FirewallSecurityPolicy1MapOutput
}

type FirewallSecurityPolicy1Map map[string]FirewallSecurityPolicy1Input

func (FirewallSecurityPolicy1Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSecurityPolicy1)(nil)).Elem()
}

func (i FirewallSecurityPolicy1Map) ToFirewallSecurityPolicy1MapOutput() FirewallSecurityPolicy1MapOutput {
	return i.ToFirewallSecurityPolicy1MapOutputWithContext(context.Background())
}

func (i FirewallSecurityPolicy1Map) ToFirewallSecurityPolicy1MapOutputWithContext(ctx context.Context) FirewallSecurityPolicy1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicy1MapOutput)
}

type FirewallSecurityPolicy1Output struct{ *pulumi.OutputState }

func (FirewallSecurityPolicy1Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSecurityPolicy1)(nil)).Elem()
}

func (o FirewallSecurityPolicy1Output) ToFirewallSecurityPolicy1Output() FirewallSecurityPolicy1Output {
	return o
}

func (o FirewallSecurityPolicy1Output) ToFirewallSecurityPolicy1OutputWithContext(ctx context.Context) FirewallSecurityPolicy1Output {
	return o
}

func (o FirewallSecurityPolicy1Output) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) ApplicationList() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.ApplicationList }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) AvProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.AvProfile }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) CapturePacket() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.CapturePacket }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o FirewallSecurityPolicy1Output) Devices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringArrayOutput { return v.Devices }).(pulumi.StringArrayOutput)
}

func (o FirewallSecurityPolicy1Output) DnsfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.DnsfilterProfile }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) Dstaddrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringArrayOutput { return v.Dstaddrs }).(pulumi.StringArrayOutput)
}

func (o FirewallSecurityPolicy1Output) Dstintfs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringArrayOutput { return v.Dstintfs }).(pulumi.StringArrayOutput)
}

func (o FirewallSecurityPolicy1Output) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o FirewallSecurityPolicy1Output) InternetService() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.InternetService }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) InternetServiceIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.IntArrayOutput { return v.InternetServiceIds }).(pulumi.IntArrayOutput)
}

func (o FirewallSecurityPolicy1Output) InternetServiceSrc() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.InternetServiceSrc }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) InternetServiceSrcIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.IntArrayOutput { return v.InternetServiceSrcIds }).(pulumi.IntArrayOutput)
}

func (o FirewallSecurityPolicy1Output) Ippool() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.Ippool }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) IpsSensor() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.IpsSensor }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) Logtraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.Logtraffic }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) LogtrafficStart() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.LogtrafficStart }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) Nat() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.Nat }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) Poolnames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringArrayOutput { return v.Poolnames }).(pulumi.StringArrayOutput)
}

func (o FirewallSecurityPolicy1Output) ProfileProtocolOptions() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.ProfileProtocolOptions }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) Services() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringArrayOutput { return v.Services }).(pulumi.StringArrayOutput)
}

func (o FirewallSecurityPolicy1Output) Srcaddrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringArrayOutput { return v.Srcaddrs }).(pulumi.StringArrayOutput)
}

func (o FirewallSecurityPolicy1Output) Srcintfs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringArrayOutput { return v.Srcintfs }).(pulumi.StringArrayOutput)
}

func (o FirewallSecurityPolicy1Output) SslSshProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.SslSshProfile }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

func (o FirewallSecurityPolicy1Output) UtmStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.UtmStatus }).(pulumi.StringOutput)
}

func (o FirewallSecurityPolicy1Output) WebfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSecurityPolicy1) pulumi.StringOutput { return v.WebfilterProfile }).(pulumi.StringOutput)
}

type FirewallSecurityPolicy1ArrayOutput struct{ *pulumi.OutputState }

func (FirewallSecurityPolicy1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSecurityPolicy1)(nil)).Elem()
}

func (o FirewallSecurityPolicy1ArrayOutput) ToFirewallSecurityPolicy1ArrayOutput() FirewallSecurityPolicy1ArrayOutput {
	return o
}

func (o FirewallSecurityPolicy1ArrayOutput) ToFirewallSecurityPolicy1ArrayOutputWithContext(ctx context.Context) FirewallSecurityPolicy1ArrayOutput {
	return o
}

func (o FirewallSecurityPolicy1ArrayOutput) Index(i pulumi.IntInput) FirewallSecurityPolicy1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallSecurityPolicy1 {
		return vs[0].([]*FirewallSecurityPolicy1)[vs[1].(int)]
	}).(FirewallSecurityPolicy1Output)
}

type FirewallSecurityPolicy1MapOutput struct{ *pulumi.OutputState }

func (FirewallSecurityPolicy1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSecurityPolicy1)(nil)).Elem()
}

func (o FirewallSecurityPolicy1MapOutput) ToFirewallSecurityPolicy1MapOutput() FirewallSecurityPolicy1MapOutput {
	return o
}

func (o FirewallSecurityPolicy1MapOutput) ToFirewallSecurityPolicy1MapOutputWithContext(ctx context.Context) FirewallSecurityPolicy1MapOutput {
	return o
}

func (o FirewallSecurityPolicy1MapOutput) MapIndex(k pulumi.StringInput) FirewallSecurityPolicy1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallSecurityPolicy1 {
		return vs[0].(map[string]*FirewallSecurityPolicy1)[vs[1].(string)]
	}).(FirewallSecurityPolicy1Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSecurityPolicy1Input)(nil)).Elem(), &FirewallSecurityPolicy1{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSecurityPolicy1ArrayInput)(nil)).Elem(), FirewallSecurityPolicy1Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSecurityPolicy1MapInput)(nil)).Elem(), FirewallSecurityPolicy1Map{})
	pulumi.RegisterOutputType(FirewallSecurityPolicy1Output{})
	pulumi.RegisterOutputType(FirewallSecurityPolicy1ArrayOutput{})
	pulumi.RegisterOutputType(FirewallSecurityPolicy1MapOutput{})
}

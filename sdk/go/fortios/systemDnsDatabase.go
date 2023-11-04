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

type SystemDnsDatabase struct {
	pulumi.CustomResourceState

	AllowTransfer       pulumi.StringOutput                  `pulumi:"allowTransfer"`
	Authoritative       pulumi.StringOutput                  `pulumi:"authoritative"`
	Contact             pulumi.StringOutput                  `pulumi:"contact"`
	DnsEntries          SystemDnsDatabaseDnsEntryArrayOutput `pulumi:"dnsEntries"`
	Domain              pulumi.StringOutput                  `pulumi:"domain"`
	DynamicSortSubtable pulumi.StringPtrOutput               `pulumi:"dynamicSortSubtable"`
	Forwarder           pulumi.StringOutput                  `pulumi:"forwarder"`
	Forwarder6          pulumi.StringOutput                  `pulumi:"forwarder6"`
	GetAllTables        pulumi.StringPtrOutput               `pulumi:"getAllTables"`
	IpMaster            pulumi.StringOutput                  `pulumi:"ipMaster"`
	IpPrimary           pulumi.StringOutput                  `pulumi:"ipPrimary"`
	Name                pulumi.StringOutput                  `pulumi:"name"`
	PrimaryName         pulumi.StringOutput                  `pulumi:"primaryName"`
	RrMax               pulumi.IntOutput                     `pulumi:"rrMax"`
	SourceIp            pulumi.StringOutput                  `pulumi:"sourceIp"`
	SourceIp6           pulumi.StringOutput                  `pulumi:"sourceIp6"`
	Status              pulumi.StringOutput                  `pulumi:"status"`
	Ttl                 pulumi.IntOutput                     `pulumi:"ttl"`
	Type                pulumi.StringOutput                  `pulumi:"type"`
	Vdomparam           pulumi.StringPtrOutput               `pulumi:"vdomparam"`
	View                pulumi.StringOutput                  `pulumi:"view"`
}

// NewSystemDnsDatabase registers a new resource with the given unique name, arguments, and options.
func NewSystemDnsDatabase(ctx *pulumi.Context,
	name string, args *SystemDnsDatabaseArgs, opts ...pulumi.ResourceOption) (*SystemDnsDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authoritative == nil {
		return nil, errors.New("invalid value for required argument 'Authoritative'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Ttl == nil {
		return nil, errors.New("invalid value for required argument 'Ttl'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.View == nil {
		return nil, errors.New("invalid value for required argument 'View'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemDnsDatabase
	err := ctx.RegisterResource("fortios:index/systemDnsDatabase:SystemDnsDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemDnsDatabase gets an existing SystemDnsDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemDnsDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemDnsDatabaseState, opts ...pulumi.ResourceOption) (*SystemDnsDatabase, error) {
	var resource SystemDnsDatabase
	err := ctx.ReadResource("fortios:index/systemDnsDatabase:SystemDnsDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemDnsDatabase resources.
type systemDnsDatabaseState struct {
	AllowTransfer       *string                     `pulumi:"allowTransfer"`
	Authoritative       *string                     `pulumi:"authoritative"`
	Contact             *string                     `pulumi:"contact"`
	DnsEntries          []SystemDnsDatabaseDnsEntry `pulumi:"dnsEntries"`
	Domain              *string                     `pulumi:"domain"`
	DynamicSortSubtable *string                     `pulumi:"dynamicSortSubtable"`
	Forwarder           *string                     `pulumi:"forwarder"`
	Forwarder6          *string                     `pulumi:"forwarder6"`
	GetAllTables        *string                     `pulumi:"getAllTables"`
	IpMaster            *string                     `pulumi:"ipMaster"`
	IpPrimary           *string                     `pulumi:"ipPrimary"`
	Name                *string                     `pulumi:"name"`
	PrimaryName         *string                     `pulumi:"primaryName"`
	RrMax               *int                        `pulumi:"rrMax"`
	SourceIp            *string                     `pulumi:"sourceIp"`
	SourceIp6           *string                     `pulumi:"sourceIp6"`
	Status              *string                     `pulumi:"status"`
	Ttl                 *int                        `pulumi:"ttl"`
	Type                *string                     `pulumi:"type"`
	Vdomparam           *string                     `pulumi:"vdomparam"`
	View                *string                     `pulumi:"view"`
}

type SystemDnsDatabaseState struct {
	AllowTransfer       pulumi.StringPtrInput
	Authoritative       pulumi.StringPtrInput
	Contact             pulumi.StringPtrInput
	DnsEntries          SystemDnsDatabaseDnsEntryArrayInput
	Domain              pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Forwarder           pulumi.StringPtrInput
	Forwarder6          pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	IpMaster            pulumi.StringPtrInput
	IpPrimary           pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	PrimaryName         pulumi.StringPtrInput
	RrMax               pulumi.IntPtrInput
	SourceIp            pulumi.StringPtrInput
	SourceIp6           pulumi.StringPtrInput
	Status              pulumi.StringPtrInput
	Ttl                 pulumi.IntPtrInput
	Type                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	View                pulumi.StringPtrInput
}

func (SystemDnsDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDnsDatabaseState)(nil)).Elem()
}

type systemDnsDatabaseArgs struct {
	AllowTransfer       *string                     `pulumi:"allowTransfer"`
	Authoritative       string                      `pulumi:"authoritative"`
	Contact             *string                     `pulumi:"contact"`
	DnsEntries          []SystemDnsDatabaseDnsEntry `pulumi:"dnsEntries"`
	Domain              string                      `pulumi:"domain"`
	DynamicSortSubtable *string                     `pulumi:"dynamicSortSubtable"`
	Forwarder           *string                     `pulumi:"forwarder"`
	Forwarder6          *string                     `pulumi:"forwarder6"`
	GetAllTables        *string                     `pulumi:"getAllTables"`
	IpMaster            *string                     `pulumi:"ipMaster"`
	IpPrimary           *string                     `pulumi:"ipPrimary"`
	Name                *string                     `pulumi:"name"`
	PrimaryName         *string                     `pulumi:"primaryName"`
	RrMax               *int                        `pulumi:"rrMax"`
	SourceIp            *string                     `pulumi:"sourceIp"`
	SourceIp6           *string                     `pulumi:"sourceIp6"`
	Status              *string                     `pulumi:"status"`
	Ttl                 int                         `pulumi:"ttl"`
	Type                string                      `pulumi:"type"`
	Vdomparam           *string                     `pulumi:"vdomparam"`
	View                string                      `pulumi:"view"`
}

// The set of arguments for constructing a SystemDnsDatabase resource.
type SystemDnsDatabaseArgs struct {
	AllowTransfer       pulumi.StringPtrInput
	Authoritative       pulumi.StringInput
	Contact             pulumi.StringPtrInput
	DnsEntries          SystemDnsDatabaseDnsEntryArrayInput
	Domain              pulumi.StringInput
	DynamicSortSubtable pulumi.StringPtrInput
	Forwarder           pulumi.StringPtrInput
	Forwarder6          pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	IpMaster            pulumi.StringPtrInput
	IpPrimary           pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	PrimaryName         pulumi.StringPtrInput
	RrMax               pulumi.IntPtrInput
	SourceIp            pulumi.StringPtrInput
	SourceIp6           pulumi.StringPtrInput
	Status              pulumi.StringPtrInput
	Ttl                 pulumi.IntInput
	Type                pulumi.StringInput
	Vdomparam           pulumi.StringPtrInput
	View                pulumi.StringInput
}

func (SystemDnsDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDnsDatabaseArgs)(nil)).Elem()
}

type SystemDnsDatabaseInput interface {
	pulumi.Input

	ToSystemDnsDatabaseOutput() SystemDnsDatabaseOutput
	ToSystemDnsDatabaseOutputWithContext(ctx context.Context) SystemDnsDatabaseOutput
}

func (*SystemDnsDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDnsDatabase)(nil)).Elem()
}

func (i *SystemDnsDatabase) ToSystemDnsDatabaseOutput() SystemDnsDatabaseOutput {
	return i.ToSystemDnsDatabaseOutputWithContext(context.Background())
}

func (i *SystemDnsDatabase) ToSystemDnsDatabaseOutputWithContext(ctx context.Context) SystemDnsDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDnsDatabaseOutput)
}

func (i *SystemDnsDatabase) ToOutput(ctx context.Context) pulumix.Output[*SystemDnsDatabase] {
	return pulumix.Output[*SystemDnsDatabase]{
		OutputState: i.ToSystemDnsDatabaseOutputWithContext(ctx).OutputState,
	}
}

// SystemDnsDatabaseArrayInput is an input type that accepts SystemDnsDatabaseArray and SystemDnsDatabaseArrayOutput values.
// You can construct a concrete instance of `SystemDnsDatabaseArrayInput` via:
//
//	SystemDnsDatabaseArray{ SystemDnsDatabaseArgs{...} }
type SystemDnsDatabaseArrayInput interface {
	pulumi.Input

	ToSystemDnsDatabaseArrayOutput() SystemDnsDatabaseArrayOutput
	ToSystemDnsDatabaseArrayOutputWithContext(context.Context) SystemDnsDatabaseArrayOutput
}

type SystemDnsDatabaseArray []SystemDnsDatabaseInput

func (SystemDnsDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDnsDatabase)(nil)).Elem()
}

func (i SystemDnsDatabaseArray) ToSystemDnsDatabaseArrayOutput() SystemDnsDatabaseArrayOutput {
	return i.ToSystemDnsDatabaseArrayOutputWithContext(context.Background())
}

func (i SystemDnsDatabaseArray) ToSystemDnsDatabaseArrayOutputWithContext(ctx context.Context) SystemDnsDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDnsDatabaseArrayOutput)
}

func (i SystemDnsDatabaseArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemDnsDatabase] {
	return pulumix.Output[[]*SystemDnsDatabase]{
		OutputState: i.ToSystemDnsDatabaseArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemDnsDatabaseMapInput is an input type that accepts SystemDnsDatabaseMap and SystemDnsDatabaseMapOutput values.
// You can construct a concrete instance of `SystemDnsDatabaseMapInput` via:
//
//	SystemDnsDatabaseMap{ "key": SystemDnsDatabaseArgs{...} }
type SystemDnsDatabaseMapInput interface {
	pulumi.Input

	ToSystemDnsDatabaseMapOutput() SystemDnsDatabaseMapOutput
	ToSystemDnsDatabaseMapOutputWithContext(context.Context) SystemDnsDatabaseMapOutput
}

type SystemDnsDatabaseMap map[string]SystemDnsDatabaseInput

func (SystemDnsDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDnsDatabase)(nil)).Elem()
}

func (i SystemDnsDatabaseMap) ToSystemDnsDatabaseMapOutput() SystemDnsDatabaseMapOutput {
	return i.ToSystemDnsDatabaseMapOutputWithContext(context.Background())
}

func (i SystemDnsDatabaseMap) ToSystemDnsDatabaseMapOutputWithContext(ctx context.Context) SystemDnsDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDnsDatabaseMapOutput)
}

func (i SystemDnsDatabaseMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemDnsDatabase] {
	return pulumix.Output[map[string]*SystemDnsDatabase]{
		OutputState: i.ToSystemDnsDatabaseMapOutputWithContext(ctx).OutputState,
	}
}

type SystemDnsDatabaseOutput struct{ *pulumi.OutputState }

func (SystemDnsDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDnsDatabase)(nil)).Elem()
}

func (o SystemDnsDatabaseOutput) ToSystemDnsDatabaseOutput() SystemDnsDatabaseOutput {
	return o
}

func (o SystemDnsDatabaseOutput) ToSystemDnsDatabaseOutputWithContext(ctx context.Context) SystemDnsDatabaseOutput {
	return o
}

func (o SystemDnsDatabaseOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemDnsDatabase] {
	return pulumix.Output[*SystemDnsDatabase]{
		OutputState: o.OutputState,
	}
}

func (o SystemDnsDatabaseOutput) AllowTransfer() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.AllowTransfer }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) Authoritative() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.Authoritative }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) Contact() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.Contact }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) DnsEntries() SystemDnsDatabaseDnsEntryArrayOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) SystemDnsDatabaseDnsEntryArrayOutput { return v.DnsEntries }).(SystemDnsDatabaseDnsEntryArrayOutput)
}

func (o SystemDnsDatabaseOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SystemDnsDatabaseOutput) Forwarder() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.Forwarder }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) Forwarder6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.Forwarder6 }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SystemDnsDatabaseOutput) IpMaster() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.IpMaster }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) IpPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.IpPrimary }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) PrimaryName() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.PrimaryName }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) RrMax() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.IntOutput { return v.RrMax }).(pulumi.IntOutput)
}

func (o SystemDnsDatabaseOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.SourceIp6 }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.IntOutput { return v.Ttl }).(pulumi.IntOutput)
}

func (o SystemDnsDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SystemDnsDatabaseOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o SystemDnsDatabaseOutput) View() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDnsDatabase) pulumi.StringOutput { return v.View }).(pulumi.StringOutput)
}

type SystemDnsDatabaseArrayOutput struct{ *pulumi.OutputState }

func (SystemDnsDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDnsDatabase)(nil)).Elem()
}

func (o SystemDnsDatabaseArrayOutput) ToSystemDnsDatabaseArrayOutput() SystemDnsDatabaseArrayOutput {
	return o
}

func (o SystemDnsDatabaseArrayOutput) ToSystemDnsDatabaseArrayOutputWithContext(ctx context.Context) SystemDnsDatabaseArrayOutput {
	return o
}

func (o SystemDnsDatabaseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemDnsDatabase] {
	return pulumix.Output[[]*SystemDnsDatabase]{
		OutputState: o.OutputState,
	}
}

func (o SystemDnsDatabaseArrayOutput) Index(i pulumi.IntInput) SystemDnsDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemDnsDatabase {
		return vs[0].([]*SystemDnsDatabase)[vs[1].(int)]
	}).(SystemDnsDatabaseOutput)
}

type SystemDnsDatabaseMapOutput struct{ *pulumi.OutputState }

func (SystemDnsDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDnsDatabase)(nil)).Elem()
}

func (o SystemDnsDatabaseMapOutput) ToSystemDnsDatabaseMapOutput() SystemDnsDatabaseMapOutput {
	return o
}

func (o SystemDnsDatabaseMapOutput) ToSystemDnsDatabaseMapOutputWithContext(ctx context.Context) SystemDnsDatabaseMapOutput {
	return o
}

func (o SystemDnsDatabaseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemDnsDatabase] {
	return pulumix.Output[map[string]*SystemDnsDatabase]{
		OutputState: o.OutputState,
	}
}

func (o SystemDnsDatabaseMapOutput) MapIndex(k pulumi.StringInput) SystemDnsDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemDnsDatabase {
		return vs[0].(map[string]*SystemDnsDatabase)[vs[1].(string)]
	}).(SystemDnsDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDnsDatabaseInput)(nil)).Elem(), &SystemDnsDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDnsDatabaseArrayInput)(nil)).Elem(), SystemDnsDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDnsDatabaseMapInput)(nil)).Elem(), SystemDnsDatabaseMap{})
	pulumi.RegisterOutputType(SystemDnsDatabaseOutput{})
	pulumi.RegisterOutputType(SystemDnsDatabaseArrayOutput{})
	pulumi.RegisterOutputType(SystemDnsDatabaseMapOutput{})
}

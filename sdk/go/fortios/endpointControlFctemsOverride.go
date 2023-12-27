// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type EndpointControlFctemsOverride struct {
	pulumi.CustomResourceState

	CallTimeout                    pulumi.IntOutput       `pulumi:"callTimeout"`
	Capabilities                   pulumi.StringOutput    `pulumi:"capabilities"`
	CloudServerType                pulumi.StringOutput    `pulumi:"cloudServerType"`
	DirtyReason                    pulumi.StringOutput    `pulumi:"dirtyReason"`
	EmsId                          pulumi.IntOutput       `pulumi:"emsId"`
	FortinetoneCloudAuthentication pulumi.StringOutput    `pulumi:"fortinetoneCloudAuthentication"`
	HttpsPort                      pulumi.IntOutput       `pulumi:"httpsPort"`
	Interface                      pulumi.StringOutput    `pulumi:"interface"`
	InterfaceSelectMethod          pulumi.StringOutput    `pulumi:"interfaceSelectMethod"`
	Name                           pulumi.StringOutput    `pulumi:"name"`
	OutOfSyncThreshold             pulumi.IntOutput       `pulumi:"outOfSyncThreshold"`
	PreserveSslSession             pulumi.StringOutput    `pulumi:"preserveSslSession"`
	PullAvatars                    pulumi.StringOutput    `pulumi:"pullAvatars"`
	PullMalwareHash                pulumi.StringOutput    `pulumi:"pullMalwareHash"`
	PullSysinfo                    pulumi.StringOutput    `pulumi:"pullSysinfo"`
	PullTags                       pulumi.StringOutput    `pulumi:"pullTags"`
	PullVulnerabilities            pulumi.StringOutput    `pulumi:"pullVulnerabilities"`
	SerialNumber                   pulumi.StringOutput    `pulumi:"serialNumber"`
	Server                         pulumi.StringOutput    `pulumi:"server"`
	SourceIp                       pulumi.StringOutput    `pulumi:"sourceIp"`
	Status                         pulumi.StringOutput    `pulumi:"status"`
	TenantId                       pulumi.StringOutput    `pulumi:"tenantId"`
	TrustCaCn                      pulumi.StringOutput    `pulumi:"trustCaCn"`
	Vdomparam                      pulumi.StringPtrOutput `pulumi:"vdomparam"`
	WebsocketOverride              pulumi.StringOutput    `pulumi:"websocketOverride"`
}

// NewEndpointControlFctemsOverride registers a new resource with the given unique name, arguments, and options.
func NewEndpointControlFctemsOverride(ctx *pulumi.Context,
	name string, args *EndpointControlFctemsOverrideArgs, opts ...pulumi.ResourceOption) (*EndpointControlFctemsOverride, error) {
	if args == nil {
		args = &EndpointControlFctemsOverrideArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EndpointControlFctemsOverride
	err := ctx.RegisterResource("fortios:index/endpointControlFctemsOverride:EndpointControlFctemsOverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointControlFctemsOverride gets an existing EndpointControlFctemsOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointControlFctemsOverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointControlFctemsOverrideState, opts ...pulumi.ResourceOption) (*EndpointControlFctemsOverride, error) {
	var resource EndpointControlFctemsOverride
	err := ctx.ReadResource("fortios:index/endpointControlFctemsOverride:EndpointControlFctemsOverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointControlFctemsOverride resources.
type endpointControlFctemsOverrideState struct {
	CallTimeout                    *int    `pulumi:"callTimeout"`
	Capabilities                   *string `pulumi:"capabilities"`
	CloudServerType                *string `pulumi:"cloudServerType"`
	DirtyReason                    *string `pulumi:"dirtyReason"`
	EmsId                          *int    `pulumi:"emsId"`
	FortinetoneCloudAuthentication *string `pulumi:"fortinetoneCloudAuthentication"`
	HttpsPort                      *int    `pulumi:"httpsPort"`
	Interface                      *string `pulumi:"interface"`
	InterfaceSelectMethod          *string `pulumi:"interfaceSelectMethod"`
	Name                           *string `pulumi:"name"`
	OutOfSyncThreshold             *int    `pulumi:"outOfSyncThreshold"`
	PreserveSslSession             *string `pulumi:"preserveSslSession"`
	PullAvatars                    *string `pulumi:"pullAvatars"`
	PullMalwareHash                *string `pulumi:"pullMalwareHash"`
	PullSysinfo                    *string `pulumi:"pullSysinfo"`
	PullTags                       *string `pulumi:"pullTags"`
	PullVulnerabilities            *string `pulumi:"pullVulnerabilities"`
	SerialNumber                   *string `pulumi:"serialNumber"`
	Server                         *string `pulumi:"server"`
	SourceIp                       *string `pulumi:"sourceIp"`
	Status                         *string `pulumi:"status"`
	TenantId                       *string `pulumi:"tenantId"`
	TrustCaCn                      *string `pulumi:"trustCaCn"`
	Vdomparam                      *string `pulumi:"vdomparam"`
	WebsocketOverride              *string `pulumi:"websocketOverride"`
}

type EndpointControlFctemsOverrideState struct {
	CallTimeout                    pulumi.IntPtrInput
	Capabilities                   pulumi.StringPtrInput
	CloudServerType                pulumi.StringPtrInput
	DirtyReason                    pulumi.StringPtrInput
	EmsId                          pulumi.IntPtrInput
	FortinetoneCloudAuthentication pulumi.StringPtrInput
	HttpsPort                      pulumi.IntPtrInput
	Interface                      pulumi.StringPtrInput
	InterfaceSelectMethod          pulumi.StringPtrInput
	Name                           pulumi.StringPtrInput
	OutOfSyncThreshold             pulumi.IntPtrInput
	PreserveSslSession             pulumi.StringPtrInput
	PullAvatars                    pulumi.StringPtrInput
	PullMalwareHash                pulumi.StringPtrInput
	PullSysinfo                    pulumi.StringPtrInput
	PullTags                       pulumi.StringPtrInput
	PullVulnerabilities            pulumi.StringPtrInput
	SerialNumber                   pulumi.StringPtrInput
	Server                         pulumi.StringPtrInput
	SourceIp                       pulumi.StringPtrInput
	Status                         pulumi.StringPtrInput
	TenantId                       pulumi.StringPtrInput
	TrustCaCn                      pulumi.StringPtrInput
	Vdomparam                      pulumi.StringPtrInput
	WebsocketOverride              pulumi.StringPtrInput
}

func (EndpointControlFctemsOverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlFctemsOverrideState)(nil)).Elem()
}

type endpointControlFctemsOverrideArgs struct {
	CallTimeout                    *int    `pulumi:"callTimeout"`
	Capabilities                   *string `pulumi:"capabilities"`
	CloudServerType                *string `pulumi:"cloudServerType"`
	DirtyReason                    *string `pulumi:"dirtyReason"`
	EmsId                          *int    `pulumi:"emsId"`
	FortinetoneCloudAuthentication *string `pulumi:"fortinetoneCloudAuthentication"`
	HttpsPort                      *int    `pulumi:"httpsPort"`
	Interface                      *string `pulumi:"interface"`
	InterfaceSelectMethod          *string `pulumi:"interfaceSelectMethod"`
	Name                           *string `pulumi:"name"`
	OutOfSyncThreshold             *int    `pulumi:"outOfSyncThreshold"`
	PreserveSslSession             *string `pulumi:"preserveSslSession"`
	PullAvatars                    *string `pulumi:"pullAvatars"`
	PullMalwareHash                *string `pulumi:"pullMalwareHash"`
	PullSysinfo                    *string `pulumi:"pullSysinfo"`
	PullTags                       *string `pulumi:"pullTags"`
	PullVulnerabilities            *string `pulumi:"pullVulnerabilities"`
	SerialNumber                   *string `pulumi:"serialNumber"`
	Server                         *string `pulumi:"server"`
	SourceIp                       *string `pulumi:"sourceIp"`
	Status                         *string `pulumi:"status"`
	TenantId                       *string `pulumi:"tenantId"`
	TrustCaCn                      *string `pulumi:"trustCaCn"`
	Vdomparam                      *string `pulumi:"vdomparam"`
	WebsocketOverride              *string `pulumi:"websocketOverride"`
}

// The set of arguments for constructing a EndpointControlFctemsOverride resource.
type EndpointControlFctemsOverrideArgs struct {
	CallTimeout                    pulumi.IntPtrInput
	Capabilities                   pulumi.StringPtrInput
	CloudServerType                pulumi.StringPtrInput
	DirtyReason                    pulumi.StringPtrInput
	EmsId                          pulumi.IntPtrInput
	FortinetoneCloudAuthentication pulumi.StringPtrInput
	HttpsPort                      pulumi.IntPtrInput
	Interface                      pulumi.StringPtrInput
	InterfaceSelectMethod          pulumi.StringPtrInput
	Name                           pulumi.StringPtrInput
	OutOfSyncThreshold             pulumi.IntPtrInput
	PreserveSslSession             pulumi.StringPtrInput
	PullAvatars                    pulumi.StringPtrInput
	PullMalwareHash                pulumi.StringPtrInput
	PullSysinfo                    pulumi.StringPtrInput
	PullTags                       pulumi.StringPtrInput
	PullVulnerabilities            pulumi.StringPtrInput
	SerialNumber                   pulumi.StringPtrInput
	Server                         pulumi.StringPtrInput
	SourceIp                       pulumi.StringPtrInput
	Status                         pulumi.StringPtrInput
	TenantId                       pulumi.StringPtrInput
	TrustCaCn                      pulumi.StringPtrInput
	Vdomparam                      pulumi.StringPtrInput
	WebsocketOverride              pulumi.StringPtrInput
}

func (EndpointControlFctemsOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlFctemsOverrideArgs)(nil)).Elem()
}

type EndpointControlFctemsOverrideInput interface {
	pulumi.Input

	ToEndpointControlFctemsOverrideOutput() EndpointControlFctemsOverrideOutput
	ToEndpointControlFctemsOverrideOutputWithContext(ctx context.Context) EndpointControlFctemsOverrideOutput
}

func (*EndpointControlFctemsOverride) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointControlFctemsOverride)(nil)).Elem()
}

func (i *EndpointControlFctemsOverride) ToEndpointControlFctemsOverrideOutput() EndpointControlFctemsOverrideOutput {
	return i.ToEndpointControlFctemsOverrideOutputWithContext(context.Background())
}

func (i *EndpointControlFctemsOverride) ToEndpointControlFctemsOverrideOutputWithContext(ctx context.Context) EndpointControlFctemsOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlFctemsOverrideOutput)
}

func (i *EndpointControlFctemsOverride) ToOutput(ctx context.Context) pulumix.Output[*EndpointControlFctemsOverride] {
	return pulumix.Output[*EndpointControlFctemsOverride]{
		OutputState: i.ToEndpointControlFctemsOverrideOutputWithContext(ctx).OutputState,
	}
}

// EndpointControlFctemsOverrideArrayInput is an input type that accepts EndpointControlFctemsOverrideArray and EndpointControlFctemsOverrideArrayOutput values.
// You can construct a concrete instance of `EndpointControlFctemsOverrideArrayInput` via:
//
//	EndpointControlFctemsOverrideArray{ EndpointControlFctemsOverrideArgs{...} }
type EndpointControlFctemsOverrideArrayInput interface {
	pulumi.Input

	ToEndpointControlFctemsOverrideArrayOutput() EndpointControlFctemsOverrideArrayOutput
	ToEndpointControlFctemsOverrideArrayOutputWithContext(context.Context) EndpointControlFctemsOverrideArrayOutput
}

type EndpointControlFctemsOverrideArray []EndpointControlFctemsOverrideInput

func (EndpointControlFctemsOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointControlFctemsOverride)(nil)).Elem()
}

func (i EndpointControlFctemsOverrideArray) ToEndpointControlFctemsOverrideArrayOutput() EndpointControlFctemsOverrideArrayOutput {
	return i.ToEndpointControlFctemsOverrideArrayOutputWithContext(context.Background())
}

func (i EndpointControlFctemsOverrideArray) ToEndpointControlFctemsOverrideArrayOutputWithContext(ctx context.Context) EndpointControlFctemsOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlFctemsOverrideArrayOutput)
}

func (i EndpointControlFctemsOverrideArray) ToOutput(ctx context.Context) pulumix.Output[[]*EndpointControlFctemsOverride] {
	return pulumix.Output[[]*EndpointControlFctemsOverride]{
		OutputState: i.ToEndpointControlFctemsOverrideArrayOutputWithContext(ctx).OutputState,
	}
}

// EndpointControlFctemsOverrideMapInput is an input type that accepts EndpointControlFctemsOverrideMap and EndpointControlFctemsOverrideMapOutput values.
// You can construct a concrete instance of `EndpointControlFctemsOverrideMapInput` via:
//
//	EndpointControlFctemsOverrideMap{ "key": EndpointControlFctemsOverrideArgs{...} }
type EndpointControlFctemsOverrideMapInput interface {
	pulumi.Input

	ToEndpointControlFctemsOverrideMapOutput() EndpointControlFctemsOverrideMapOutput
	ToEndpointControlFctemsOverrideMapOutputWithContext(context.Context) EndpointControlFctemsOverrideMapOutput
}

type EndpointControlFctemsOverrideMap map[string]EndpointControlFctemsOverrideInput

func (EndpointControlFctemsOverrideMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointControlFctemsOverride)(nil)).Elem()
}

func (i EndpointControlFctemsOverrideMap) ToEndpointControlFctemsOverrideMapOutput() EndpointControlFctemsOverrideMapOutput {
	return i.ToEndpointControlFctemsOverrideMapOutputWithContext(context.Background())
}

func (i EndpointControlFctemsOverrideMap) ToEndpointControlFctemsOverrideMapOutputWithContext(ctx context.Context) EndpointControlFctemsOverrideMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlFctemsOverrideMapOutput)
}

func (i EndpointControlFctemsOverrideMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*EndpointControlFctemsOverride] {
	return pulumix.Output[map[string]*EndpointControlFctemsOverride]{
		OutputState: i.ToEndpointControlFctemsOverrideMapOutputWithContext(ctx).OutputState,
	}
}

type EndpointControlFctemsOverrideOutput struct{ *pulumi.OutputState }

func (EndpointControlFctemsOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointControlFctemsOverride)(nil)).Elem()
}

func (o EndpointControlFctemsOverrideOutput) ToEndpointControlFctemsOverrideOutput() EndpointControlFctemsOverrideOutput {
	return o
}

func (o EndpointControlFctemsOverrideOutput) ToEndpointControlFctemsOverrideOutputWithContext(ctx context.Context) EndpointControlFctemsOverrideOutput {
	return o
}

func (o EndpointControlFctemsOverrideOutput) ToOutput(ctx context.Context) pulumix.Output[*EndpointControlFctemsOverride] {
	return pulumix.Output[*EndpointControlFctemsOverride]{
		OutputState: o.OutputState,
	}
}

func (o EndpointControlFctemsOverrideOutput) CallTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.IntOutput { return v.CallTimeout }).(pulumi.IntOutput)
}

func (o EndpointControlFctemsOverrideOutput) Capabilities() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.Capabilities }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) CloudServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.CloudServerType }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) DirtyReason() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.DirtyReason }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) EmsId() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.IntOutput { return v.EmsId }).(pulumi.IntOutput)
}

func (o EndpointControlFctemsOverrideOutput) FortinetoneCloudAuthentication() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.FortinetoneCloudAuthentication }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) HttpsPort() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.IntOutput { return v.HttpsPort }).(pulumi.IntOutput)
}

func (o EndpointControlFctemsOverrideOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) OutOfSyncThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.IntOutput { return v.OutOfSyncThreshold }).(pulumi.IntOutput)
}

func (o EndpointControlFctemsOverrideOutput) PreserveSslSession() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.PreserveSslSession }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) PullAvatars() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.PullAvatars }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) PullMalwareHash() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.PullMalwareHash }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) PullSysinfo() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.PullSysinfo }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) PullTags() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.PullTags }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) PullVulnerabilities() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.PullVulnerabilities }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) TrustCaCn() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.TrustCaCn }).(pulumi.StringOutput)
}

func (o EndpointControlFctemsOverrideOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o EndpointControlFctemsOverrideOutput) WebsocketOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlFctemsOverride) pulumi.StringOutput { return v.WebsocketOverride }).(pulumi.StringOutput)
}

type EndpointControlFctemsOverrideArrayOutput struct{ *pulumi.OutputState }

func (EndpointControlFctemsOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointControlFctemsOverride)(nil)).Elem()
}

func (o EndpointControlFctemsOverrideArrayOutput) ToEndpointControlFctemsOverrideArrayOutput() EndpointControlFctemsOverrideArrayOutput {
	return o
}

func (o EndpointControlFctemsOverrideArrayOutput) ToEndpointControlFctemsOverrideArrayOutputWithContext(ctx context.Context) EndpointControlFctemsOverrideArrayOutput {
	return o
}

func (o EndpointControlFctemsOverrideArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*EndpointControlFctemsOverride] {
	return pulumix.Output[[]*EndpointControlFctemsOverride]{
		OutputState: o.OutputState,
	}
}

func (o EndpointControlFctemsOverrideArrayOutput) Index(i pulumi.IntInput) EndpointControlFctemsOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointControlFctemsOverride {
		return vs[0].([]*EndpointControlFctemsOverride)[vs[1].(int)]
	}).(EndpointControlFctemsOverrideOutput)
}

type EndpointControlFctemsOverrideMapOutput struct{ *pulumi.OutputState }

func (EndpointControlFctemsOverrideMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointControlFctemsOverride)(nil)).Elem()
}

func (o EndpointControlFctemsOverrideMapOutput) ToEndpointControlFctemsOverrideMapOutput() EndpointControlFctemsOverrideMapOutput {
	return o
}

func (o EndpointControlFctemsOverrideMapOutput) ToEndpointControlFctemsOverrideMapOutputWithContext(ctx context.Context) EndpointControlFctemsOverrideMapOutput {
	return o
}

func (o EndpointControlFctemsOverrideMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*EndpointControlFctemsOverride] {
	return pulumix.Output[map[string]*EndpointControlFctemsOverride]{
		OutputState: o.OutputState,
	}
}

func (o EndpointControlFctemsOverrideMapOutput) MapIndex(k pulumi.StringInput) EndpointControlFctemsOverrideOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointControlFctemsOverride {
		return vs[0].(map[string]*EndpointControlFctemsOverride)[vs[1].(string)]
	}).(EndpointControlFctemsOverrideOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointControlFctemsOverrideInput)(nil)).Elem(), &EndpointControlFctemsOverride{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointControlFctemsOverrideArrayInput)(nil)).Elem(), EndpointControlFctemsOverrideArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointControlFctemsOverrideMapInput)(nil)).Elem(), EndpointControlFctemsOverrideMap{})
	pulumi.RegisterOutputType(EndpointControlFctemsOverrideOutput{})
	pulumi.RegisterOutputType(EndpointControlFctemsOverrideArrayOutput{})
	pulumi.RegisterOutputType(EndpointControlFctemsOverrideMapOutput{})
}
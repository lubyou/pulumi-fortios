// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EndpointControlSettings struct {
	pulumi.CustomResourceState

	DownloadCustomLink                     pulumi.StringOutput    `pulumi:"downloadCustomLink"`
	DownloadLocation                       pulumi.StringOutput    `pulumi:"downloadLocation"`
	ForticlientAvdbUpdateInterval          pulumi.IntOutput       `pulumi:"forticlientAvdbUpdateInterval"`
	ForticlientDeregUnsupportedClient      pulumi.StringOutput    `pulumi:"forticlientDeregUnsupportedClient"`
	ForticlientDisconnectUnsupportedClient pulumi.StringOutput    `pulumi:"forticlientDisconnectUnsupportedClient"`
	ForticlientEmsRestApiCallTimeout       pulumi.IntOutput       `pulumi:"forticlientEmsRestApiCallTimeout"`
	ForticlientKeepaliveInterval           pulumi.IntOutput       `pulumi:"forticlientKeepaliveInterval"`
	ForticlientOfflineGrace                pulumi.StringOutput    `pulumi:"forticlientOfflineGrace"`
	ForticlientOfflineGraceInterval        pulumi.IntOutput       `pulumi:"forticlientOfflineGraceInterval"`
	ForticlientRegKey                      pulumi.StringPtrOutput `pulumi:"forticlientRegKey"`
	ForticlientRegKeyEnforce               pulumi.StringOutput    `pulumi:"forticlientRegKeyEnforce"`
	ForticlientRegTimeout                  pulumi.IntOutput       `pulumi:"forticlientRegTimeout"`
	ForticlientSysUpdateInterval           pulumi.IntOutput       `pulumi:"forticlientSysUpdateInterval"`
	ForticlientUserAvatar                  pulumi.StringOutput    `pulumi:"forticlientUserAvatar"`
	ForticlientWarningInterval             pulumi.IntOutput       `pulumi:"forticlientWarningInterval"`
	Vdomparam                              pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewEndpointControlSettings registers a new resource with the given unique name, arguments, and options.
func NewEndpointControlSettings(ctx *pulumi.Context,
	name string, args *EndpointControlSettingsArgs, opts ...pulumi.ResourceOption) (*EndpointControlSettings, error) {
	if args == nil {
		args = &EndpointControlSettingsArgs{}
	}

	if args.ForticlientRegKey != nil {
		args.ForticlientRegKey = pulumi.ToSecret(args.ForticlientRegKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"forticlientRegKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource EndpointControlSettings
	err := ctx.RegisterResource("fortios:index/endpointControlSettings:EndpointControlSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointControlSettings gets an existing EndpointControlSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointControlSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointControlSettingsState, opts ...pulumi.ResourceOption) (*EndpointControlSettings, error) {
	var resource EndpointControlSettings
	err := ctx.ReadResource("fortios:index/endpointControlSettings:EndpointControlSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointControlSettings resources.
type endpointControlSettingsState struct {
	DownloadCustomLink                     *string `pulumi:"downloadCustomLink"`
	DownloadLocation                       *string `pulumi:"downloadLocation"`
	ForticlientAvdbUpdateInterval          *int    `pulumi:"forticlientAvdbUpdateInterval"`
	ForticlientDeregUnsupportedClient      *string `pulumi:"forticlientDeregUnsupportedClient"`
	ForticlientDisconnectUnsupportedClient *string `pulumi:"forticlientDisconnectUnsupportedClient"`
	ForticlientEmsRestApiCallTimeout       *int    `pulumi:"forticlientEmsRestApiCallTimeout"`
	ForticlientKeepaliveInterval           *int    `pulumi:"forticlientKeepaliveInterval"`
	ForticlientOfflineGrace                *string `pulumi:"forticlientOfflineGrace"`
	ForticlientOfflineGraceInterval        *int    `pulumi:"forticlientOfflineGraceInterval"`
	ForticlientRegKey                      *string `pulumi:"forticlientRegKey"`
	ForticlientRegKeyEnforce               *string `pulumi:"forticlientRegKeyEnforce"`
	ForticlientRegTimeout                  *int    `pulumi:"forticlientRegTimeout"`
	ForticlientSysUpdateInterval           *int    `pulumi:"forticlientSysUpdateInterval"`
	ForticlientUserAvatar                  *string `pulumi:"forticlientUserAvatar"`
	ForticlientWarningInterval             *int    `pulumi:"forticlientWarningInterval"`
	Vdomparam                              *string `pulumi:"vdomparam"`
}

type EndpointControlSettingsState struct {
	DownloadCustomLink                     pulumi.StringPtrInput
	DownloadLocation                       pulumi.StringPtrInput
	ForticlientAvdbUpdateInterval          pulumi.IntPtrInput
	ForticlientDeregUnsupportedClient      pulumi.StringPtrInput
	ForticlientDisconnectUnsupportedClient pulumi.StringPtrInput
	ForticlientEmsRestApiCallTimeout       pulumi.IntPtrInput
	ForticlientKeepaliveInterval           pulumi.IntPtrInput
	ForticlientOfflineGrace                pulumi.StringPtrInput
	ForticlientOfflineGraceInterval        pulumi.IntPtrInput
	ForticlientRegKey                      pulumi.StringPtrInput
	ForticlientRegKeyEnforce               pulumi.StringPtrInput
	ForticlientRegTimeout                  pulumi.IntPtrInput
	ForticlientSysUpdateInterval           pulumi.IntPtrInput
	ForticlientUserAvatar                  pulumi.StringPtrInput
	ForticlientWarningInterval             pulumi.IntPtrInput
	Vdomparam                              pulumi.StringPtrInput
}

func (EndpointControlSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlSettingsState)(nil)).Elem()
}

type endpointControlSettingsArgs struct {
	DownloadCustomLink                     *string `pulumi:"downloadCustomLink"`
	DownloadLocation                       *string `pulumi:"downloadLocation"`
	ForticlientAvdbUpdateInterval          *int    `pulumi:"forticlientAvdbUpdateInterval"`
	ForticlientDeregUnsupportedClient      *string `pulumi:"forticlientDeregUnsupportedClient"`
	ForticlientDisconnectUnsupportedClient *string `pulumi:"forticlientDisconnectUnsupportedClient"`
	ForticlientEmsRestApiCallTimeout       *int    `pulumi:"forticlientEmsRestApiCallTimeout"`
	ForticlientKeepaliveInterval           *int    `pulumi:"forticlientKeepaliveInterval"`
	ForticlientOfflineGrace                *string `pulumi:"forticlientOfflineGrace"`
	ForticlientOfflineGraceInterval        *int    `pulumi:"forticlientOfflineGraceInterval"`
	ForticlientRegKey                      *string `pulumi:"forticlientRegKey"`
	ForticlientRegKeyEnforce               *string `pulumi:"forticlientRegKeyEnforce"`
	ForticlientRegTimeout                  *int    `pulumi:"forticlientRegTimeout"`
	ForticlientSysUpdateInterval           *int    `pulumi:"forticlientSysUpdateInterval"`
	ForticlientUserAvatar                  *string `pulumi:"forticlientUserAvatar"`
	ForticlientWarningInterval             *int    `pulumi:"forticlientWarningInterval"`
	Vdomparam                              *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EndpointControlSettings resource.
type EndpointControlSettingsArgs struct {
	DownloadCustomLink                     pulumi.StringPtrInput
	DownloadLocation                       pulumi.StringPtrInput
	ForticlientAvdbUpdateInterval          pulumi.IntPtrInput
	ForticlientDeregUnsupportedClient      pulumi.StringPtrInput
	ForticlientDisconnectUnsupportedClient pulumi.StringPtrInput
	ForticlientEmsRestApiCallTimeout       pulumi.IntPtrInput
	ForticlientKeepaliveInterval           pulumi.IntPtrInput
	ForticlientOfflineGrace                pulumi.StringPtrInput
	ForticlientOfflineGraceInterval        pulumi.IntPtrInput
	ForticlientRegKey                      pulumi.StringPtrInput
	ForticlientRegKeyEnforce               pulumi.StringPtrInput
	ForticlientRegTimeout                  pulumi.IntPtrInput
	ForticlientSysUpdateInterval           pulumi.IntPtrInput
	ForticlientUserAvatar                  pulumi.StringPtrInput
	ForticlientWarningInterval             pulumi.IntPtrInput
	Vdomparam                              pulumi.StringPtrInput
}

func (EndpointControlSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlSettingsArgs)(nil)).Elem()
}

type EndpointControlSettingsInput interface {
	pulumi.Input

	ToEndpointControlSettingsOutput() EndpointControlSettingsOutput
	ToEndpointControlSettingsOutputWithContext(ctx context.Context) EndpointControlSettingsOutput
}

func (*EndpointControlSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointControlSettings)(nil)).Elem()
}

func (i *EndpointControlSettings) ToEndpointControlSettingsOutput() EndpointControlSettingsOutput {
	return i.ToEndpointControlSettingsOutputWithContext(context.Background())
}

func (i *EndpointControlSettings) ToEndpointControlSettingsOutputWithContext(ctx context.Context) EndpointControlSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlSettingsOutput)
}

// EndpointControlSettingsArrayInput is an input type that accepts EndpointControlSettingsArray and EndpointControlSettingsArrayOutput values.
// You can construct a concrete instance of `EndpointControlSettingsArrayInput` via:
//
//	EndpointControlSettingsArray{ EndpointControlSettingsArgs{...} }
type EndpointControlSettingsArrayInput interface {
	pulumi.Input

	ToEndpointControlSettingsArrayOutput() EndpointControlSettingsArrayOutput
	ToEndpointControlSettingsArrayOutputWithContext(context.Context) EndpointControlSettingsArrayOutput
}

type EndpointControlSettingsArray []EndpointControlSettingsInput

func (EndpointControlSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointControlSettings)(nil)).Elem()
}

func (i EndpointControlSettingsArray) ToEndpointControlSettingsArrayOutput() EndpointControlSettingsArrayOutput {
	return i.ToEndpointControlSettingsArrayOutputWithContext(context.Background())
}

func (i EndpointControlSettingsArray) ToEndpointControlSettingsArrayOutputWithContext(ctx context.Context) EndpointControlSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlSettingsArrayOutput)
}

// EndpointControlSettingsMapInput is an input type that accepts EndpointControlSettingsMap and EndpointControlSettingsMapOutput values.
// You can construct a concrete instance of `EndpointControlSettingsMapInput` via:
//
//	EndpointControlSettingsMap{ "key": EndpointControlSettingsArgs{...} }
type EndpointControlSettingsMapInput interface {
	pulumi.Input

	ToEndpointControlSettingsMapOutput() EndpointControlSettingsMapOutput
	ToEndpointControlSettingsMapOutputWithContext(context.Context) EndpointControlSettingsMapOutput
}

type EndpointControlSettingsMap map[string]EndpointControlSettingsInput

func (EndpointControlSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointControlSettings)(nil)).Elem()
}

func (i EndpointControlSettingsMap) ToEndpointControlSettingsMapOutput() EndpointControlSettingsMapOutput {
	return i.ToEndpointControlSettingsMapOutputWithContext(context.Background())
}

func (i EndpointControlSettingsMap) ToEndpointControlSettingsMapOutputWithContext(ctx context.Context) EndpointControlSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlSettingsMapOutput)
}

type EndpointControlSettingsOutput struct{ *pulumi.OutputState }

func (EndpointControlSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointControlSettings)(nil)).Elem()
}

func (o EndpointControlSettingsOutput) ToEndpointControlSettingsOutput() EndpointControlSettingsOutput {
	return o
}

func (o EndpointControlSettingsOutput) ToEndpointControlSettingsOutputWithContext(ctx context.Context) EndpointControlSettingsOutput {
	return o
}

func (o EndpointControlSettingsOutput) DownloadCustomLink() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.StringOutput { return v.DownloadCustomLink }).(pulumi.StringOutput)
}

func (o EndpointControlSettingsOutput) DownloadLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.StringOutput { return v.DownloadLocation }).(pulumi.StringOutput)
}

func (o EndpointControlSettingsOutput) ForticlientAvdbUpdateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.IntOutput { return v.ForticlientAvdbUpdateInterval }).(pulumi.IntOutput)
}

func (o EndpointControlSettingsOutput) ForticlientDeregUnsupportedClient() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.StringOutput { return v.ForticlientDeregUnsupportedClient }).(pulumi.StringOutput)
}

func (o EndpointControlSettingsOutput) ForticlientDisconnectUnsupportedClient() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.StringOutput { return v.ForticlientDisconnectUnsupportedClient }).(pulumi.StringOutput)
}

func (o EndpointControlSettingsOutput) ForticlientEmsRestApiCallTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.IntOutput { return v.ForticlientEmsRestApiCallTimeout }).(pulumi.IntOutput)
}

func (o EndpointControlSettingsOutput) ForticlientKeepaliveInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.IntOutput { return v.ForticlientKeepaliveInterval }).(pulumi.IntOutput)
}

func (o EndpointControlSettingsOutput) ForticlientOfflineGrace() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.StringOutput { return v.ForticlientOfflineGrace }).(pulumi.StringOutput)
}

func (o EndpointControlSettingsOutput) ForticlientOfflineGraceInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.IntOutput { return v.ForticlientOfflineGraceInterval }).(pulumi.IntOutput)
}

func (o EndpointControlSettingsOutput) ForticlientRegKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.StringPtrOutput { return v.ForticlientRegKey }).(pulumi.StringPtrOutput)
}

func (o EndpointControlSettingsOutput) ForticlientRegKeyEnforce() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.StringOutput { return v.ForticlientRegKeyEnforce }).(pulumi.StringOutput)
}

func (o EndpointControlSettingsOutput) ForticlientRegTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.IntOutput { return v.ForticlientRegTimeout }).(pulumi.IntOutput)
}

func (o EndpointControlSettingsOutput) ForticlientSysUpdateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.IntOutput { return v.ForticlientSysUpdateInterval }).(pulumi.IntOutput)
}

func (o EndpointControlSettingsOutput) ForticlientUserAvatar() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.StringOutput { return v.ForticlientUserAvatar }).(pulumi.StringOutput)
}

func (o EndpointControlSettingsOutput) ForticlientWarningInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.IntOutput { return v.ForticlientWarningInterval }).(pulumi.IntOutput)
}

func (o EndpointControlSettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointControlSettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type EndpointControlSettingsArrayOutput struct{ *pulumi.OutputState }

func (EndpointControlSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointControlSettings)(nil)).Elem()
}

func (o EndpointControlSettingsArrayOutput) ToEndpointControlSettingsArrayOutput() EndpointControlSettingsArrayOutput {
	return o
}

func (o EndpointControlSettingsArrayOutput) ToEndpointControlSettingsArrayOutputWithContext(ctx context.Context) EndpointControlSettingsArrayOutput {
	return o
}

func (o EndpointControlSettingsArrayOutput) Index(i pulumi.IntInput) EndpointControlSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointControlSettings {
		return vs[0].([]*EndpointControlSettings)[vs[1].(int)]
	}).(EndpointControlSettingsOutput)
}

type EndpointControlSettingsMapOutput struct{ *pulumi.OutputState }

func (EndpointControlSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointControlSettings)(nil)).Elem()
}

func (o EndpointControlSettingsMapOutput) ToEndpointControlSettingsMapOutput() EndpointControlSettingsMapOutput {
	return o
}

func (o EndpointControlSettingsMapOutput) ToEndpointControlSettingsMapOutputWithContext(ctx context.Context) EndpointControlSettingsMapOutput {
	return o
}

func (o EndpointControlSettingsMapOutput) MapIndex(k pulumi.StringInput) EndpointControlSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointControlSettings {
		return vs[0].(map[string]*EndpointControlSettings)[vs[1].(string)]
	}).(EndpointControlSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointControlSettingsInput)(nil)).Elem(), &EndpointControlSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointControlSettingsArrayInput)(nil)).Elem(), EndpointControlSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointControlSettingsMapInput)(nil)).Elem(), EndpointControlSettingsMap{})
	pulumi.RegisterOutputType(EndpointControlSettingsOutput{})
	pulumi.RegisterOutputType(EndpointControlSettingsArrayOutput{})
	pulumi.RegisterOutputType(EndpointControlSettingsMapOutput{})
}

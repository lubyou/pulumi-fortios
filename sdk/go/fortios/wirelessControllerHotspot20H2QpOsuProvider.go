// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WirelessControllerHotspot20H2QpOsuProvider struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput                                                  `pulumi:"dynamicSortSubtable"`
	FriendlyNames       WirelessControllerHotspot20H2QpOsuProviderFriendlyNameArrayOutput       `pulumi:"friendlyNames"`
	Icon                pulumi.StringOutput                                                     `pulumi:"icon"`
	Name                pulumi.StringOutput                                                     `pulumi:"name"`
	OsuMethod           pulumi.StringOutput                                                     `pulumi:"osuMethod"`
	OsuNai              pulumi.StringOutput                                                     `pulumi:"osuNai"`
	ServerUri           pulumi.StringOutput                                                     `pulumi:"serverUri"`
	ServiceDescriptions WirelessControllerHotspot20H2QpOsuProviderServiceDescriptionArrayOutput `pulumi:"serviceDescriptions"`
	Vdomparam           pulumi.StringPtrOutput                                                  `pulumi:"vdomparam"`
}

// NewWirelessControllerHotspot20H2QpOsuProvider registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerHotspot20H2QpOsuProvider(ctx *pulumi.Context,
	name string, args *WirelessControllerHotspot20H2QpOsuProviderArgs, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20H2QpOsuProvider, error) {
	if args == nil {
		args = &WirelessControllerHotspot20H2QpOsuProviderArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelessControllerHotspot20H2QpOsuProvider
	err := ctx.RegisterResource("fortios:index/wirelessControllerHotspot20H2QpOsuProvider:WirelessControllerHotspot20H2QpOsuProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerHotspot20H2QpOsuProvider gets an existing WirelessControllerHotspot20H2QpOsuProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerHotspot20H2QpOsuProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerHotspot20H2QpOsuProviderState, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20H2QpOsuProvider, error) {
	var resource WirelessControllerHotspot20H2QpOsuProvider
	err := ctx.ReadResource("fortios:index/wirelessControllerHotspot20H2QpOsuProvider:WirelessControllerHotspot20H2QpOsuProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerHotspot20H2QpOsuProvider resources.
type wirelessControllerHotspot20H2QpOsuProviderState struct {
	DynamicSortSubtable *string                                                        `pulumi:"dynamicSortSubtable"`
	FriendlyNames       []WirelessControllerHotspot20H2QpOsuProviderFriendlyName       `pulumi:"friendlyNames"`
	Icon                *string                                                        `pulumi:"icon"`
	Name                *string                                                        `pulumi:"name"`
	OsuMethod           *string                                                        `pulumi:"osuMethod"`
	OsuNai              *string                                                        `pulumi:"osuNai"`
	ServerUri           *string                                                        `pulumi:"serverUri"`
	ServiceDescriptions []WirelessControllerHotspot20H2QpOsuProviderServiceDescription `pulumi:"serviceDescriptions"`
	Vdomparam           *string                                                        `pulumi:"vdomparam"`
}

type WirelessControllerHotspot20H2QpOsuProviderState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	FriendlyNames       WirelessControllerHotspot20H2QpOsuProviderFriendlyNameArrayInput
	Icon                pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	OsuMethod           pulumi.StringPtrInput
	OsuNai              pulumi.StringPtrInput
	ServerUri           pulumi.StringPtrInput
	ServiceDescriptions WirelessControllerHotspot20H2QpOsuProviderServiceDescriptionArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerHotspot20H2QpOsuProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20H2QpOsuProviderState)(nil)).Elem()
}

type wirelessControllerHotspot20H2QpOsuProviderArgs struct {
	DynamicSortSubtable *string                                                        `pulumi:"dynamicSortSubtable"`
	FriendlyNames       []WirelessControllerHotspot20H2QpOsuProviderFriendlyName       `pulumi:"friendlyNames"`
	Icon                *string                                                        `pulumi:"icon"`
	Name                *string                                                        `pulumi:"name"`
	OsuMethod           *string                                                        `pulumi:"osuMethod"`
	OsuNai              *string                                                        `pulumi:"osuNai"`
	ServerUri           *string                                                        `pulumi:"serverUri"`
	ServiceDescriptions []WirelessControllerHotspot20H2QpOsuProviderServiceDescription `pulumi:"serviceDescriptions"`
	Vdomparam           *string                                                        `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerHotspot20H2QpOsuProvider resource.
type WirelessControllerHotspot20H2QpOsuProviderArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	FriendlyNames       WirelessControllerHotspot20H2QpOsuProviderFriendlyNameArrayInput
	Icon                pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	OsuMethod           pulumi.StringPtrInput
	OsuNai              pulumi.StringPtrInput
	ServerUri           pulumi.StringPtrInput
	ServiceDescriptions WirelessControllerHotspot20H2QpOsuProviderServiceDescriptionArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerHotspot20H2QpOsuProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20H2QpOsuProviderArgs)(nil)).Elem()
}

type WirelessControllerHotspot20H2QpOsuProviderInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20H2QpOsuProviderOutput() WirelessControllerHotspot20H2QpOsuProviderOutput
	ToWirelessControllerHotspot20H2QpOsuProviderOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOsuProviderOutput
}

func (*WirelessControllerHotspot20H2QpOsuProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20H2QpOsuProvider)(nil)).Elem()
}

func (i *WirelessControllerHotspot20H2QpOsuProvider) ToWirelessControllerHotspot20H2QpOsuProviderOutput() WirelessControllerHotspot20H2QpOsuProviderOutput {
	return i.ToWirelessControllerHotspot20H2QpOsuProviderOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20H2QpOsuProvider) ToWirelessControllerHotspot20H2QpOsuProviderOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOsuProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20H2QpOsuProviderOutput)
}

// WirelessControllerHotspot20H2QpOsuProviderArrayInput is an input type that accepts WirelessControllerHotspot20H2QpOsuProviderArray and WirelessControllerHotspot20H2QpOsuProviderArrayOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20H2QpOsuProviderArrayInput` via:
//
//	WirelessControllerHotspot20H2QpOsuProviderArray{ WirelessControllerHotspot20H2QpOsuProviderArgs{...} }
type WirelessControllerHotspot20H2QpOsuProviderArrayInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20H2QpOsuProviderArrayOutput() WirelessControllerHotspot20H2QpOsuProviderArrayOutput
	ToWirelessControllerHotspot20H2QpOsuProviderArrayOutputWithContext(context.Context) WirelessControllerHotspot20H2QpOsuProviderArrayOutput
}

type WirelessControllerHotspot20H2QpOsuProviderArray []WirelessControllerHotspot20H2QpOsuProviderInput

func (WirelessControllerHotspot20H2QpOsuProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20H2QpOsuProvider)(nil)).Elem()
}

func (i WirelessControllerHotspot20H2QpOsuProviderArray) ToWirelessControllerHotspot20H2QpOsuProviderArrayOutput() WirelessControllerHotspot20H2QpOsuProviderArrayOutput {
	return i.ToWirelessControllerHotspot20H2QpOsuProviderArrayOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20H2QpOsuProviderArray) ToWirelessControllerHotspot20H2QpOsuProviderArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOsuProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20H2QpOsuProviderArrayOutput)
}

// WirelessControllerHotspot20H2QpOsuProviderMapInput is an input type that accepts WirelessControllerHotspot20H2QpOsuProviderMap and WirelessControllerHotspot20H2QpOsuProviderMapOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20H2QpOsuProviderMapInput` via:
//
//	WirelessControllerHotspot20H2QpOsuProviderMap{ "key": WirelessControllerHotspot20H2QpOsuProviderArgs{...} }
type WirelessControllerHotspot20H2QpOsuProviderMapInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20H2QpOsuProviderMapOutput() WirelessControllerHotspot20H2QpOsuProviderMapOutput
	ToWirelessControllerHotspot20H2QpOsuProviderMapOutputWithContext(context.Context) WirelessControllerHotspot20H2QpOsuProviderMapOutput
}

type WirelessControllerHotspot20H2QpOsuProviderMap map[string]WirelessControllerHotspot20H2QpOsuProviderInput

func (WirelessControllerHotspot20H2QpOsuProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20H2QpOsuProvider)(nil)).Elem()
}

func (i WirelessControllerHotspot20H2QpOsuProviderMap) ToWirelessControllerHotspot20H2QpOsuProviderMapOutput() WirelessControllerHotspot20H2QpOsuProviderMapOutput {
	return i.ToWirelessControllerHotspot20H2QpOsuProviderMapOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20H2QpOsuProviderMap) ToWirelessControllerHotspot20H2QpOsuProviderMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOsuProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20H2QpOsuProviderMapOutput)
}

type WirelessControllerHotspot20H2QpOsuProviderOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20H2QpOsuProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20H2QpOsuProvider)(nil)).Elem()
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) ToWirelessControllerHotspot20H2QpOsuProviderOutput() WirelessControllerHotspot20H2QpOsuProviderOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) ToWirelessControllerHotspot20H2QpOsuProviderOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOsuProviderOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20H2QpOsuProvider) pulumi.StringPtrOutput {
		return v.DynamicSortSubtable
	}).(pulumi.StringPtrOutput)
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) FriendlyNames() WirelessControllerHotspot20H2QpOsuProviderFriendlyNameArrayOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20H2QpOsuProvider) WirelessControllerHotspot20H2QpOsuProviderFriendlyNameArrayOutput {
		return v.FriendlyNames
	}).(WirelessControllerHotspot20H2QpOsuProviderFriendlyNameArrayOutput)
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) Icon() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20H2QpOsuProvider) pulumi.StringOutput { return v.Icon }).(pulumi.StringOutput)
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20H2QpOsuProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) OsuMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20H2QpOsuProvider) pulumi.StringOutput { return v.OsuMethod }).(pulumi.StringOutput)
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) OsuNai() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20H2QpOsuProvider) pulumi.StringOutput { return v.OsuNai }).(pulumi.StringOutput)
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) ServerUri() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20H2QpOsuProvider) pulumi.StringOutput { return v.ServerUri }).(pulumi.StringOutput)
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) ServiceDescriptions() WirelessControllerHotspot20H2QpOsuProviderServiceDescriptionArrayOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20H2QpOsuProvider) WirelessControllerHotspot20H2QpOsuProviderServiceDescriptionArrayOutput {
		return v.ServiceDescriptions
	}).(WirelessControllerHotspot20H2QpOsuProviderServiceDescriptionArrayOutput)
}

func (o WirelessControllerHotspot20H2QpOsuProviderOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20H2QpOsuProvider) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelessControllerHotspot20H2QpOsuProviderArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20H2QpOsuProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20H2QpOsuProvider)(nil)).Elem()
}

func (o WirelessControllerHotspot20H2QpOsuProviderArrayOutput) ToWirelessControllerHotspot20H2QpOsuProviderArrayOutput() WirelessControllerHotspot20H2QpOsuProviderArrayOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOsuProviderArrayOutput) ToWirelessControllerHotspot20H2QpOsuProviderArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOsuProviderArrayOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOsuProviderArrayOutput) Index(i pulumi.IntInput) WirelessControllerHotspot20H2QpOsuProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20H2QpOsuProvider {
		return vs[0].([]*WirelessControllerHotspot20H2QpOsuProvider)[vs[1].(int)]
	}).(WirelessControllerHotspot20H2QpOsuProviderOutput)
}

type WirelessControllerHotspot20H2QpOsuProviderMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20H2QpOsuProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20H2QpOsuProvider)(nil)).Elem()
}

func (o WirelessControllerHotspot20H2QpOsuProviderMapOutput) ToWirelessControllerHotspot20H2QpOsuProviderMapOutput() WirelessControllerHotspot20H2QpOsuProviderMapOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOsuProviderMapOutput) ToWirelessControllerHotspot20H2QpOsuProviderMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOsuProviderMapOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOsuProviderMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerHotspot20H2QpOsuProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20H2QpOsuProvider {
		return vs[0].(map[string]*WirelessControllerHotspot20H2QpOsuProvider)[vs[1].(string)]
	}).(WirelessControllerHotspot20H2QpOsuProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20H2QpOsuProviderInput)(nil)).Elem(), &WirelessControllerHotspot20H2QpOsuProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20H2QpOsuProviderArrayInput)(nil)).Elem(), WirelessControllerHotspot20H2QpOsuProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20H2QpOsuProviderMapInput)(nil)).Elem(), WirelessControllerHotspot20H2QpOsuProviderMap{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20H2QpOsuProviderOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20H2QpOsuProviderArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20H2QpOsuProviderMapOutput{})
}

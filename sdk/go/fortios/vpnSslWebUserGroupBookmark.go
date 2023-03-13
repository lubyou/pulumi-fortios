// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnSslWebUserGroupBookmark struct {
	pulumi.CustomResourceState

	Bookmarks           VpnSslWebUserGroupBookmarkBookmarkArrayOutput `pulumi:"bookmarks"`
	DynamicSortSubtable pulumi.StringPtrOutput                        `pulumi:"dynamicSortSubtable"`
	Name                pulumi.StringOutput                           `pulumi:"name"`
	Vdomparam           pulumi.StringPtrOutput                        `pulumi:"vdomparam"`
}

// NewVpnSslWebUserGroupBookmark registers a new resource with the given unique name, arguments, and options.
func NewVpnSslWebUserGroupBookmark(ctx *pulumi.Context,
	name string, args *VpnSslWebUserGroupBookmarkArgs, opts ...pulumi.ResourceOption) (*VpnSslWebUserGroupBookmark, error) {
	if args == nil {
		args = &VpnSslWebUserGroupBookmarkArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource VpnSslWebUserGroupBookmark
	err := ctx.RegisterResource("fortios:index/vpnSslWebUserGroupBookmark:VpnSslWebUserGroupBookmark", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnSslWebUserGroupBookmark gets an existing VpnSslWebUserGroupBookmark resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnSslWebUserGroupBookmark(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnSslWebUserGroupBookmarkState, opts ...pulumi.ResourceOption) (*VpnSslWebUserGroupBookmark, error) {
	var resource VpnSslWebUserGroupBookmark
	err := ctx.ReadResource("fortios:index/vpnSslWebUserGroupBookmark:VpnSslWebUserGroupBookmark", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnSslWebUserGroupBookmark resources.
type vpnSslWebUserGroupBookmarkState struct {
	Bookmarks           []VpnSslWebUserGroupBookmarkBookmark `pulumi:"bookmarks"`
	DynamicSortSubtable *string                              `pulumi:"dynamicSortSubtable"`
	Name                *string                              `pulumi:"name"`
	Vdomparam           *string                              `pulumi:"vdomparam"`
}

type VpnSslWebUserGroupBookmarkState struct {
	Bookmarks           VpnSslWebUserGroupBookmarkBookmarkArrayInput
	DynamicSortSubtable pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (VpnSslWebUserGroupBookmarkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSslWebUserGroupBookmarkState)(nil)).Elem()
}

type vpnSslWebUserGroupBookmarkArgs struct {
	Bookmarks           []VpnSslWebUserGroupBookmarkBookmark `pulumi:"bookmarks"`
	DynamicSortSubtable *string                              `pulumi:"dynamicSortSubtable"`
	Name                *string                              `pulumi:"name"`
	Vdomparam           *string                              `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnSslWebUserGroupBookmark resource.
type VpnSslWebUserGroupBookmarkArgs struct {
	Bookmarks           VpnSslWebUserGroupBookmarkBookmarkArrayInput
	DynamicSortSubtable pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (VpnSslWebUserGroupBookmarkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSslWebUserGroupBookmarkArgs)(nil)).Elem()
}

type VpnSslWebUserGroupBookmarkInput interface {
	pulumi.Input

	ToVpnSslWebUserGroupBookmarkOutput() VpnSslWebUserGroupBookmarkOutput
	ToVpnSslWebUserGroupBookmarkOutputWithContext(ctx context.Context) VpnSslWebUserGroupBookmarkOutput
}

func (*VpnSslWebUserGroupBookmark) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnSslWebUserGroupBookmark)(nil)).Elem()
}

func (i *VpnSslWebUserGroupBookmark) ToVpnSslWebUserGroupBookmarkOutput() VpnSslWebUserGroupBookmarkOutput {
	return i.ToVpnSslWebUserGroupBookmarkOutputWithContext(context.Background())
}

func (i *VpnSslWebUserGroupBookmark) ToVpnSslWebUserGroupBookmarkOutputWithContext(ctx context.Context) VpnSslWebUserGroupBookmarkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSslWebUserGroupBookmarkOutput)
}

// VpnSslWebUserGroupBookmarkArrayInput is an input type that accepts VpnSslWebUserGroupBookmarkArray and VpnSslWebUserGroupBookmarkArrayOutput values.
// You can construct a concrete instance of `VpnSslWebUserGroupBookmarkArrayInput` via:
//
//	VpnSslWebUserGroupBookmarkArray{ VpnSslWebUserGroupBookmarkArgs{...} }
type VpnSslWebUserGroupBookmarkArrayInput interface {
	pulumi.Input

	ToVpnSslWebUserGroupBookmarkArrayOutput() VpnSslWebUserGroupBookmarkArrayOutput
	ToVpnSslWebUserGroupBookmarkArrayOutputWithContext(context.Context) VpnSslWebUserGroupBookmarkArrayOutput
}

type VpnSslWebUserGroupBookmarkArray []VpnSslWebUserGroupBookmarkInput

func (VpnSslWebUserGroupBookmarkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnSslWebUserGroupBookmark)(nil)).Elem()
}

func (i VpnSslWebUserGroupBookmarkArray) ToVpnSslWebUserGroupBookmarkArrayOutput() VpnSslWebUserGroupBookmarkArrayOutput {
	return i.ToVpnSslWebUserGroupBookmarkArrayOutputWithContext(context.Background())
}

func (i VpnSslWebUserGroupBookmarkArray) ToVpnSslWebUserGroupBookmarkArrayOutputWithContext(ctx context.Context) VpnSslWebUserGroupBookmarkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSslWebUserGroupBookmarkArrayOutput)
}

// VpnSslWebUserGroupBookmarkMapInput is an input type that accepts VpnSslWebUserGroupBookmarkMap and VpnSslWebUserGroupBookmarkMapOutput values.
// You can construct a concrete instance of `VpnSslWebUserGroupBookmarkMapInput` via:
//
//	VpnSslWebUserGroupBookmarkMap{ "key": VpnSslWebUserGroupBookmarkArgs{...} }
type VpnSslWebUserGroupBookmarkMapInput interface {
	pulumi.Input

	ToVpnSslWebUserGroupBookmarkMapOutput() VpnSslWebUserGroupBookmarkMapOutput
	ToVpnSslWebUserGroupBookmarkMapOutputWithContext(context.Context) VpnSslWebUserGroupBookmarkMapOutput
}

type VpnSslWebUserGroupBookmarkMap map[string]VpnSslWebUserGroupBookmarkInput

func (VpnSslWebUserGroupBookmarkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnSslWebUserGroupBookmark)(nil)).Elem()
}

func (i VpnSslWebUserGroupBookmarkMap) ToVpnSslWebUserGroupBookmarkMapOutput() VpnSslWebUserGroupBookmarkMapOutput {
	return i.ToVpnSslWebUserGroupBookmarkMapOutputWithContext(context.Background())
}

func (i VpnSslWebUserGroupBookmarkMap) ToVpnSslWebUserGroupBookmarkMapOutputWithContext(ctx context.Context) VpnSslWebUserGroupBookmarkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSslWebUserGroupBookmarkMapOutput)
}

type VpnSslWebUserGroupBookmarkOutput struct{ *pulumi.OutputState }

func (VpnSslWebUserGroupBookmarkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnSslWebUserGroupBookmark)(nil)).Elem()
}

func (o VpnSslWebUserGroupBookmarkOutput) ToVpnSslWebUserGroupBookmarkOutput() VpnSslWebUserGroupBookmarkOutput {
	return o
}

func (o VpnSslWebUserGroupBookmarkOutput) ToVpnSslWebUserGroupBookmarkOutputWithContext(ctx context.Context) VpnSslWebUserGroupBookmarkOutput {
	return o
}

func (o VpnSslWebUserGroupBookmarkOutput) Bookmarks() VpnSslWebUserGroupBookmarkBookmarkArrayOutput {
	return o.ApplyT(func(v *VpnSslWebUserGroupBookmark) VpnSslWebUserGroupBookmarkBookmarkArrayOutput { return v.Bookmarks }).(VpnSslWebUserGroupBookmarkBookmarkArrayOutput)
}

func (o VpnSslWebUserGroupBookmarkOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnSslWebUserGroupBookmark) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o VpnSslWebUserGroupBookmarkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnSslWebUserGroupBookmark) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnSslWebUserGroupBookmarkOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnSslWebUserGroupBookmark) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpnSslWebUserGroupBookmarkArrayOutput struct{ *pulumi.OutputState }

func (VpnSslWebUserGroupBookmarkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnSslWebUserGroupBookmark)(nil)).Elem()
}

func (o VpnSslWebUserGroupBookmarkArrayOutput) ToVpnSslWebUserGroupBookmarkArrayOutput() VpnSslWebUserGroupBookmarkArrayOutput {
	return o
}

func (o VpnSslWebUserGroupBookmarkArrayOutput) ToVpnSslWebUserGroupBookmarkArrayOutputWithContext(ctx context.Context) VpnSslWebUserGroupBookmarkArrayOutput {
	return o
}

func (o VpnSslWebUserGroupBookmarkArrayOutput) Index(i pulumi.IntInput) VpnSslWebUserGroupBookmarkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnSslWebUserGroupBookmark {
		return vs[0].([]*VpnSslWebUserGroupBookmark)[vs[1].(int)]
	}).(VpnSslWebUserGroupBookmarkOutput)
}

type VpnSslWebUserGroupBookmarkMapOutput struct{ *pulumi.OutputState }

func (VpnSslWebUserGroupBookmarkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnSslWebUserGroupBookmark)(nil)).Elem()
}

func (o VpnSslWebUserGroupBookmarkMapOutput) ToVpnSslWebUserGroupBookmarkMapOutput() VpnSslWebUserGroupBookmarkMapOutput {
	return o
}

func (o VpnSslWebUserGroupBookmarkMapOutput) ToVpnSslWebUserGroupBookmarkMapOutputWithContext(ctx context.Context) VpnSslWebUserGroupBookmarkMapOutput {
	return o
}

func (o VpnSslWebUserGroupBookmarkMapOutput) MapIndex(k pulumi.StringInput) VpnSslWebUserGroupBookmarkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnSslWebUserGroupBookmark {
		return vs[0].(map[string]*VpnSslWebUserGroupBookmark)[vs[1].(string)]
	}).(VpnSslWebUserGroupBookmarkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnSslWebUserGroupBookmarkInput)(nil)).Elem(), &VpnSslWebUserGroupBookmark{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnSslWebUserGroupBookmarkArrayInput)(nil)).Elem(), VpnSslWebUserGroupBookmarkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnSslWebUserGroupBookmarkMapInput)(nil)).Elem(), VpnSslWebUserGroupBookmarkMap{})
	pulumi.RegisterOutputType(VpnSslWebUserGroupBookmarkOutput{})
	pulumi.RegisterOutputType(VpnSslWebUserGroupBookmarkArrayOutput{})
	pulumi.RegisterOutputType(VpnSslWebUserGroupBookmarkMapOutput{})
}

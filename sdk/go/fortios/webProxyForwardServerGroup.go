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

type WebProxyForwardServerGroup struct {
	pulumi.CustomResourceState

	Affinity            pulumi.StringOutput                             `pulumi:"affinity"`
	DynamicSortSubtable pulumi.StringPtrOutput                          `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                          `pulumi:"getAllTables"`
	GroupDownOption     pulumi.StringOutput                             `pulumi:"groupDownOption"`
	LdbMethod           pulumi.StringOutput                             `pulumi:"ldbMethod"`
	Name                pulumi.StringOutput                             `pulumi:"name"`
	ServerLists         WebProxyForwardServerGroupServerListArrayOutput `pulumi:"serverLists"`
	Vdomparam           pulumi.StringPtrOutput                          `pulumi:"vdomparam"`
}

// NewWebProxyForwardServerGroup registers a new resource with the given unique name, arguments, and options.
func NewWebProxyForwardServerGroup(ctx *pulumi.Context,
	name string, args *WebProxyForwardServerGroupArgs, opts ...pulumi.ResourceOption) (*WebProxyForwardServerGroup, error) {
	if args == nil {
		args = &WebProxyForwardServerGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebProxyForwardServerGroup
	err := ctx.RegisterResource("fortios:index/webProxyForwardServerGroup:WebProxyForwardServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebProxyForwardServerGroup gets an existing WebProxyForwardServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebProxyForwardServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebProxyForwardServerGroupState, opts ...pulumi.ResourceOption) (*WebProxyForwardServerGroup, error) {
	var resource WebProxyForwardServerGroup
	err := ctx.ReadResource("fortios:index/webProxyForwardServerGroup:WebProxyForwardServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebProxyForwardServerGroup resources.
type webProxyForwardServerGroupState struct {
	Affinity            *string                                `pulumi:"affinity"`
	DynamicSortSubtable *string                                `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                `pulumi:"getAllTables"`
	GroupDownOption     *string                                `pulumi:"groupDownOption"`
	LdbMethod           *string                                `pulumi:"ldbMethod"`
	Name                *string                                `pulumi:"name"`
	ServerLists         []WebProxyForwardServerGroupServerList `pulumi:"serverLists"`
	Vdomparam           *string                                `pulumi:"vdomparam"`
}

type WebProxyForwardServerGroupState struct {
	Affinity            pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	GroupDownOption     pulumi.StringPtrInput
	LdbMethod           pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ServerLists         WebProxyForwardServerGroupServerListArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WebProxyForwardServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyForwardServerGroupState)(nil)).Elem()
}

type webProxyForwardServerGroupArgs struct {
	Affinity            *string                                `pulumi:"affinity"`
	DynamicSortSubtable *string                                `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                `pulumi:"getAllTables"`
	GroupDownOption     *string                                `pulumi:"groupDownOption"`
	LdbMethod           *string                                `pulumi:"ldbMethod"`
	Name                *string                                `pulumi:"name"`
	ServerLists         []WebProxyForwardServerGroupServerList `pulumi:"serverLists"`
	Vdomparam           *string                                `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebProxyForwardServerGroup resource.
type WebProxyForwardServerGroupArgs struct {
	Affinity            pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	GroupDownOption     pulumi.StringPtrInput
	LdbMethod           pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ServerLists         WebProxyForwardServerGroupServerListArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WebProxyForwardServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyForwardServerGroupArgs)(nil)).Elem()
}

type WebProxyForwardServerGroupInput interface {
	pulumi.Input

	ToWebProxyForwardServerGroupOutput() WebProxyForwardServerGroupOutput
	ToWebProxyForwardServerGroupOutputWithContext(ctx context.Context) WebProxyForwardServerGroupOutput
}

func (*WebProxyForwardServerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WebProxyForwardServerGroup)(nil)).Elem()
}

func (i *WebProxyForwardServerGroup) ToWebProxyForwardServerGroupOutput() WebProxyForwardServerGroupOutput {
	return i.ToWebProxyForwardServerGroupOutputWithContext(context.Background())
}

func (i *WebProxyForwardServerGroup) ToWebProxyForwardServerGroupOutputWithContext(ctx context.Context) WebProxyForwardServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyForwardServerGroupOutput)
}

func (i *WebProxyForwardServerGroup) ToOutput(ctx context.Context) pulumix.Output[*WebProxyForwardServerGroup] {
	return pulumix.Output[*WebProxyForwardServerGroup]{
		OutputState: i.ToWebProxyForwardServerGroupOutputWithContext(ctx).OutputState,
	}
}

// WebProxyForwardServerGroupArrayInput is an input type that accepts WebProxyForwardServerGroupArray and WebProxyForwardServerGroupArrayOutput values.
// You can construct a concrete instance of `WebProxyForwardServerGroupArrayInput` via:
//
//	WebProxyForwardServerGroupArray{ WebProxyForwardServerGroupArgs{...} }
type WebProxyForwardServerGroupArrayInput interface {
	pulumi.Input

	ToWebProxyForwardServerGroupArrayOutput() WebProxyForwardServerGroupArrayOutput
	ToWebProxyForwardServerGroupArrayOutputWithContext(context.Context) WebProxyForwardServerGroupArrayOutput
}

type WebProxyForwardServerGroupArray []WebProxyForwardServerGroupInput

func (WebProxyForwardServerGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebProxyForwardServerGroup)(nil)).Elem()
}

func (i WebProxyForwardServerGroupArray) ToWebProxyForwardServerGroupArrayOutput() WebProxyForwardServerGroupArrayOutput {
	return i.ToWebProxyForwardServerGroupArrayOutputWithContext(context.Background())
}

func (i WebProxyForwardServerGroupArray) ToWebProxyForwardServerGroupArrayOutputWithContext(ctx context.Context) WebProxyForwardServerGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyForwardServerGroupArrayOutput)
}

func (i WebProxyForwardServerGroupArray) ToOutput(ctx context.Context) pulumix.Output[[]*WebProxyForwardServerGroup] {
	return pulumix.Output[[]*WebProxyForwardServerGroup]{
		OutputState: i.ToWebProxyForwardServerGroupArrayOutputWithContext(ctx).OutputState,
	}
}

// WebProxyForwardServerGroupMapInput is an input type that accepts WebProxyForwardServerGroupMap and WebProxyForwardServerGroupMapOutput values.
// You can construct a concrete instance of `WebProxyForwardServerGroupMapInput` via:
//
//	WebProxyForwardServerGroupMap{ "key": WebProxyForwardServerGroupArgs{...} }
type WebProxyForwardServerGroupMapInput interface {
	pulumi.Input

	ToWebProxyForwardServerGroupMapOutput() WebProxyForwardServerGroupMapOutput
	ToWebProxyForwardServerGroupMapOutputWithContext(context.Context) WebProxyForwardServerGroupMapOutput
}

type WebProxyForwardServerGroupMap map[string]WebProxyForwardServerGroupInput

func (WebProxyForwardServerGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebProxyForwardServerGroup)(nil)).Elem()
}

func (i WebProxyForwardServerGroupMap) ToWebProxyForwardServerGroupMapOutput() WebProxyForwardServerGroupMapOutput {
	return i.ToWebProxyForwardServerGroupMapOutputWithContext(context.Background())
}

func (i WebProxyForwardServerGroupMap) ToWebProxyForwardServerGroupMapOutputWithContext(ctx context.Context) WebProxyForwardServerGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyForwardServerGroupMapOutput)
}

func (i WebProxyForwardServerGroupMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WebProxyForwardServerGroup] {
	return pulumix.Output[map[string]*WebProxyForwardServerGroup]{
		OutputState: i.ToWebProxyForwardServerGroupMapOutputWithContext(ctx).OutputState,
	}
}

type WebProxyForwardServerGroupOutput struct{ *pulumi.OutputState }

func (WebProxyForwardServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebProxyForwardServerGroup)(nil)).Elem()
}

func (o WebProxyForwardServerGroupOutput) ToWebProxyForwardServerGroupOutput() WebProxyForwardServerGroupOutput {
	return o
}

func (o WebProxyForwardServerGroupOutput) ToWebProxyForwardServerGroupOutputWithContext(ctx context.Context) WebProxyForwardServerGroupOutput {
	return o
}

func (o WebProxyForwardServerGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*WebProxyForwardServerGroup] {
	return pulumix.Output[*WebProxyForwardServerGroup]{
		OutputState: o.OutputState,
	}
}

func (o WebProxyForwardServerGroupOutput) Affinity() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyForwardServerGroup) pulumi.StringOutput { return v.Affinity }).(pulumi.StringOutput)
}

func (o WebProxyForwardServerGroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebProxyForwardServerGroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WebProxyForwardServerGroupOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebProxyForwardServerGroup) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WebProxyForwardServerGroupOutput) GroupDownOption() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyForwardServerGroup) pulumi.StringOutput { return v.GroupDownOption }).(pulumi.StringOutput)
}

func (o WebProxyForwardServerGroupOutput) LdbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyForwardServerGroup) pulumi.StringOutput { return v.LdbMethod }).(pulumi.StringOutput)
}

func (o WebProxyForwardServerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyForwardServerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebProxyForwardServerGroupOutput) ServerLists() WebProxyForwardServerGroupServerListArrayOutput {
	return o.ApplyT(func(v *WebProxyForwardServerGroup) WebProxyForwardServerGroupServerListArrayOutput {
		return v.ServerLists
	}).(WebProxyForwardServerGroupServerListArrayOutput)
}

func (o WebProxyForwardServerGroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebProxyForwardServerGroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WebProxyForwardServerGroupArrayOutput struct{ *pulumi.OutputState }

func (WebProxyForwardServerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebProxyForwardServerGroup)(nil)).Elem()
}

func (o WebProxyForwardServerGroupArrayOutput) ToWebProxyForwardServerGroupArrayOutput() WebProxyForwardServerGroupArrayOutput {
	return o
}

func (o WebProxyForwardServerGroupArrayOutput) ToWebProxyForwardServerGroupArrayOutputWithContext(ctx context.Context) WebProxyForwardServerGroupArrayOutput {
	return o
}

func (o WebProxyForwardServerGroupArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WebProxyForwardServerGroup] {
	return pulumix.Output[[]*WebProxyForwardServerGroup]{
		OutputState: o.OutputState,
	}
}

func (o WebProxyForwardServerGroupArrayOutput) Index(i pulumi.IntInput) WebProxyForwardServerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebProxyForwardServerGroup {
		return vs[0].([]*WebProxyForwardServerGroup)[vs[1].(int)]
	}).(WebProxyForwardServerGroupOutput)
}

type WebProxyForwardServerGroupMapOutput struct{ *pulumi.OutputState }

func (WebProxyForwardServerGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebProxyForwardServerGroup)(nil)).Elem()
}

func (o WebProxyForwardServerGroupMapOutput) ToWebProxyForwardServerGroupMapOutput() WebProxyForwardServerGroupMapOutput {
	return o
}

func (o WebProxyForwardServerGroupMapOutput) ToWebProxyForwardServerGroupMapOutputWithContext(ctx context.Context) WebProxyForwardServerGroupMapOutput {
	return o
}

func (o WebProxyForwardServerGroupMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WebProxyForwardServerGroup] {
	return pulumix.Output[map[string]*WebProxyForwardServerGroup]{
		OutputState: o.OutputState,
	}
}

func (o WebProxyForwardServerGroupMapOutput) MapIndex(k pulumi.StringInput) WebProxyForwardServerGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebProxyForwardServerGroup {
		return vs[0].(map[string]*WebProxyForwardServerGroup)[vs[1].(string)]
	}).(WebProxyForwardServerGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyForwardServerGroupInput)(nil)).Elem(), &WebProxyForwardServerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyForwardServerGroupArrayInput)(nil)).Elem(), WebProxyForwardServerGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyForwardServerGroupMapInput)(nil)).Elem(), WebProxyForwardServerGroupMap{})
	pulumi.RegisterOutputType(WebProxyForwardServerGroupOutput{})
	pulumi.RegisterOutputType(WebProxyForwardServerGroupArrayOutput{})
	pulumi.RegisterOutputType(WebProxyForwardServerGroupMapOutput{})
}

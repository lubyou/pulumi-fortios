// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallProxypolicySort struct {
	pulumi.CustomResourceState

	Comment       pulumi.StringPtrOutput `pulumi:"comment"`
	ForceRecreate pulumi.StringPtrOutput `pulumi:"forceRecreate"`
	Sortby        pulumi.StringOutput    `pulumi:"sortby"`
	Sortdirection pulumi.StringOutput    `pulumi:"sortdirection"`
	Status        pulumi.StringPtrOutput `pulumi:"status"`
	Vdomparam     pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallProxypolicySort registers a new resource with the given unique name, arguments, and options.
func NewFirewallProxypolicySort(ctx *pulumi.Context,
	name string, args *FirewallProxypolicySortArgs, opts ...pulumi.ResourceOption) (*FirewallProxypolicySort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Sortby == nil {
		return nil, errors.New("invalid value for required argument 'Sortby'")
	}
	if args.Sortdirection == nil {
		return nil, errors.New("invalid value for required argument 'Sortdirection'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallProxypolicySort
	err := ctx.RegisterResource("fortios:index/firewallProxypolicySort:FirewallProxypolicySort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallProxypolicySort gets an existing FirewallProxypolicySort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallProxypolicySort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallProxypolicySortState, opts ...pulumi.ResourceOption) (*FirewallProxypolicySort, error) {
	var resource FirewallProxypolicySort
	err := ctx.ReadResource("fortios:index/firewallProxypolicySort:FirewallProxypolicySort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallProxypolicySort resources.
type firewallProxypolicySortState struct {
	Comment       *string `pulumi:"comment"`
	ForceRecreate *string `pulumi:"forceRecreate"`
	Sortby        *string `pulumi:"sortby"`
	Sortdirection *string `pulumi:"sortdirection"`
	Status        *string `pulumi:"status"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

type FirewallProxypolicySortState struct {
	Comment       pulumi.StringPtrInput
	ForceRecreate pulumi.StringPtrInput
	Sortby        pulumi.StringPtrInput
	Sortdirection pulumi.StringPtrInput
	Status        pulumi.StringPtrInput
	Vdomparam     pulumi.StringPtrInput
}

func (FirewallProxypolicySortState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxypolicySortState)(nil)).Elem()
}

type firewallProxypolicySortArgs struct {
	Comment       *string `pulumi:"comment"`
	ForceRecreate *string `pulumi:"forceRecreate"`
	Sortby        string  `pulumi:"sortby"`
	Sortdirection string  `pulumi:"sortdirection"`
	Status        *string `pulumi:"status"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallProxypolicySort resource.
type FirewallProxypolicySortArgs struct {
	Comment       pulumi.StringPtrInput
	ForceRecreate pulumi.StringPtrInput
	Sortby        pulumi.StringInput
	Sortdirection pulumi.StringInput
	Status        pulumi.StringPtrInput
	Vdomparam     pulumi.StringPtrInput
}

func (FirewallProxypolicySortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxypolicySortArgs)(nil)).Elem()
}

type FirewallProxypolicySortInput interface {
	pulumi.Input

	ToFirewallProxypolicySortOutput() FirewallProxypolicySortOutput
	ToFirewallProxypolicySortOutputWithContext(ctx context.Context) FirewallProxypolicySortOutput
}

func (*FirewallProxypolicySort) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxypolicySort)(nil)).Elem()
}

func (i *FirewallProxypolicySort) ToFirewallProxypolicySortOutput() FirewallProxypolicySortOutput {
	return i.ToFirewallProxypolicySortOutputWithContext(context.Background())
}

func (i *FirewallProxypolicySort) ToFirewallProxypolicySortOutputWithContext(ctx context.Context) FirewallProxypolicySortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxypolicySortOutput)
}

// FirewallProxypolicySortArrayInput is an input type that accepts FirewallProxypolicySortArray and FirewallProxypolicySortArrayOutput values.
// You can construct a concrete instance of `FirewallProxypolicySortArrayInput` via:
//
//          FirewallProxypolicySortArray{ FirewallProxypolicySortArgs{...} }
type FirewallProxypolicySortArrayInput interface {
	pulumi.Input

	ToFirewallProxypolicySortArrayOutput() FirewallProxypolicySortArrayOutput
	ToFirewallProxypolicySortArrayOutputWithContext(context.Context) FirewallProxypolicySortArrayOutput
}

type FirewallProxypolicySortArray []FirewallProxypolicySortInput

func (FirewallProxypolicySortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProxypolicySort)(nil)).Elem()
}

func (i FirewallProxypolicySortArray) ToFirewallProxypolicySortArrayOutput() FirewallProxypolicySortArrayOutput {
	return i.ToFirewallProxypolicySortArrayOutputWithContext(context.Background())
}

func (i FirewallProxypolicySortArray) ToFirewallProxypolicySortArrayOutputWithContext(ctx context.Context) FirewallProxypolicySortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxypolicySortArrayOutput)
}

// FirewallProxypolicySortMapInput is an input type that accepts FirewallProxypolicySortMap and FirewallProxypolicySortMapOutput values.
// You can construct a concrete instance of `FirewallProxypolicySortMapInput` via:
//
//          FirewallProxypolicySortMap{ "key": FirewallProxypolicySortArgs{...} }
type FirewallProxypolicySortMapInput interface {
	pulumi.Input

	ToFirewallProxypolicySortMapOutput() FirewallProxypolicySortMapOutput
	ToFirewallProxypolicySortMapOutputWithContext(context.Context) FirewallProxypolicySortMapOutput
}

type FirewallProxypolicySortMap map[string]FirewallProxypolicySortInput

func (FirewallProxypolicySortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProxypolicySort)(nil)).Elem()
}

func (i FirewallProxypolicySortMap) ToFirewallProxypolicySortMapOutput() FirewallProxypolicySortMapOutput {
	return i.ToFirewallProxypolicySortMapOutputWithContext(context.Background())
}

func (i FirewallProxypolicySortMap) ToFirewallProxypolicySortMapOutputWithContext(ctx context.Context) FirewallProxypolicySortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxypolicySortMapOutput)
}

type FirewallProxypolicySortOutput struct{ *pulumi.OutputState }

func (FirewallProxypolicySortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxypolicySort)(nil)).Elem()
}

func (o FirewallProxypolicySortOutput) ToFirewallProxypolicySortOutput() FirewallProxypolicySortOutput {
	return o
}

func (o FirewallProxypolicySortOutput) ToFirewallProxypolicySortOutputWithContext(ctx context.Context) FirewallProxypolicySortOutput {
	return o
}

type FirewallProxypolicySortArrayOutput struct{ *pulumi.OutputState }

func (FirewallProxypolicySortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProxypolicySort)(nil)).Elem()
}

func (o FirewallProxypolicySortArrayOutput) ToFirewallProxypolicySortArrayOutput() FirewallProxypolicySortArrayOutput {
	return o
}

func (o FirewallProxypolicySortArrayOutput) ToFirewallProxypolicySortArrayOutputWithContext(ctx context.Context) FirewallProxypolicySortArrayOutput {
	return o
}

func (o FirewallProxypolicySortArrayOutput) Index(i pulumi.IntInput) FirewallProxypolicySortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallProxypolicySort {
		return vs[0].([]*FirewallProxypolicySort)[vs[1].(int)]
	}).(FirewallProxypolicySortOutput)
}

type FirewallProxypolicySortMapOutput struct{ *pulumi.OutputState }

func (FirewallProxypolicySortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProxypolicySort)(nil)).Elem()
}

func (o FirewallProxypolicySortMapOutput) ToFirewallProxypolicySortMapOutput() FirewallProxypolicySortMapOutput {
	return o
}

func (o FirewallProxypolicySortMapOutput) ToFirewallProxypolicySortMapOutputWithContext(ctx context.Context) FirewallProxypolicySortMapOutput {
	return o
}

func (o FirewallProxypolicySortMapOutput) MapIndex(k pulumi.StringInput) FirewallProxypolicySortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallProxypolicySort {
		return vs[0].(map[string]*FirewallProxypolicySort)[vs[1].(string)]
	}).(FirewallProxypolicySortOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxypolicySortInput)(nil)).Elem(), &FirewallProxypolicySort{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxypolicySortArrayInput)(nil)).Elem(), FirewallProxypolicySortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxypolicySortMapInput)(nil)).Elem(), FirewallProxypolicySortMap{})
	pulumi.RegisterOutputType(FirewallProxypolicySortOutput{})
	pulumi.RegisterOutputType(FirewallProxypolicySortArrayOutput{})
	pulumi.RegisterOutputType(FirewallProxypolicySortMapOutput{})
}

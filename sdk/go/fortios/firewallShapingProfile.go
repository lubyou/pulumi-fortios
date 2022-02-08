// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure shaping profiles.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallShapingProfile(ctx, "trname", &fortios.FirewallShapingProfileArgs{
// 			DefaultClassId: pulumi.Int(2),
// 			ProfileName:    pulumi.String("shapingprofiles1"),
// 			ShapingEntries: FirewallShapingProfileShapingEntryArray{
// 				&FirewallShapingProfileShapingEntryArgs{
// 					ClassId:                       pulumi.Int(2),
// 					GuaranteedBandwidthPercentage: pulumi.Int(33),
// 					Id:                            pulumi.Int(1),
// 					MaximumBandwidthPercentage:    pulumi.Int(88),
// 					Priority:                      pulumi.String("high"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Firewall ShapingProfile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallShapingProfile:FirewallShapingProfile labelname {{profile_name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallShapingProfile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Default class ID to handle unclassified packets (including all local traffic).
	DefaultClassId pulumi.IntOutput `pulumi:"defaultClassId"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Shaping profile name.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
	ShapingEntries FirewallShapingProfileShapingEntryArrayOutput `pulumi:"shapingEntries"`
	// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallShapingProfile registers a new resource with the given unique name, arguments, and options.
func NewFirewallShapingProfile(ctx *pulumi.Context,
	name string, args *FirewallShapingProfileArgs, opts ...pulumi.ResourceOption) (*FirewallShapingProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultClassId == nil {
		return nil, errors.New("invalid value for required argument 'DefaultClassId'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallShapingProfile
	err := ctx.RegisterResource("fortios:index/firewallShapingProfile:FirewallShapingProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallShapingProfile gets an existing FirewallShapingProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallShapingProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallShapingProfileState, opts ...pulumi.ResourceOption) (*FirewallShapingProfile, error) {
	var resource FirewallShapingProfile
	err := ctx.ReadResource("fortios:index/firewallShapingProfile:FirewallShapingProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallShapingProfile resources.
type firewallShapingProfileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Default class ID to handle unclassified packets (including all local traffic).
	DefaultClassId *int `pulumi:"defaultClassId"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Shaping profile name.
	ProfileName *string `pulumi:"profileName"`
	// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
	ShapingEntries []FirewallShapingProfileShapingEntry `pulumi:"shapingEntries"`
	// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallShapingProfileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Default class ID to handle unclassified packets (including all local traffic).
	DefaultClassId pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Shaping profile name.
	ProfileName pulumi.StringPtrInput
	// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
	ShapingEntries FirewallShapingProfileShapingEntryArrayInput
	// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallShapingProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallShapingProfileState)(nil)).Elem()
}

type firewallShapingProfileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Default class ID to handle unclassified packets (including all local traffic).
	DefaultClassId int `pulumi:"defaultClassId"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Shaping profile name.
	ProfileName string `pulumi:"profileName"`
	// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
	ShapingEntries []FirewallShapingProfileShapingEntry `pulumi:"shapingEntries"`
	// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallShapingProfile resource.
type FirewallShapingProfileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Default class ID to handle unclassified packets (including all local traffic).
	DefaultClassId pulumi.IntInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Shaping profile name.
	ProfileName pulumi.StringInput
	// Define shaping entries of this shaping profile. The structure of `shapingEntries` block is documented below.
	ShapingEntries FirewallShapingProfileShapingEntryArrayInput
	// Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallShapingProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallShapingProfileArgs)(nil)).Elem()
}

type FirewallShapingProfileInput interface {
	pulumi.Input

	ToFirewallShapingProfileOutput() FirewallShapingProfileOutput
	ToFirewallShapingProfileOutputWithContext(ctx context.Context) FirewallShapingProfileOutput
}

func (*FirewallShapingProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallShapingProfile)(nil)).Elem()
}

func (i *FirewallShapingProfile) ToFirewallShapingProfileOutput() FirewallShapingProfileOutput {
	return i.ToFirewallShapingProfileOutputWithContext(context.Background())
}

func (i *FirewallShapingProfile) ToFirewallShapingProfileOutputWithContext(ctx context.Context) FirewallShapingProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShapingProfileOutput)
}

// FirewallShapingProfileArrayInput is an input type that accepts FirewallShapingProfileArray and FirewallShapingProfileArrayOutput values.
// You can construct a concrete instance of `FirewallShapingProfileArrayInput` via:
//
//          FirewallShapingProfileArray{ FirewallShapingProfileArgs{...} }
type FirewallShapingProfileArrayInput interface {
	pulumi.Input

	ToFirewallShapingProfileArrayOutput() FirewallShapingProfileArrayOutput
	ToFirewallShapingProfileArrayOutputWithContext(context.Context) FirewallShapingProfileArrayOutput
}

type FirewallShapingProfileArray []FirewallShapingProfileInput

func (FirewallShapingProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallShapingProfile)(nil)).Elem()
}

func (i FirewallShapingProfileArray) ToFirewallShapingProfileArrayOutput() FirewallShapingProfileArrayOutput {
	return i.ToFirewallShapingProfileArrayOutputWithContext(context.Background())
}

func (i FirewallShapingProfileArray) ToFirewallShapingProfileArrayOutputWithContext(ctx context.Context) FirewallShapingProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShapingProfileArrayOutput)
}

// FirewallShapingProfileMapInput is an input type that accepts FirewallShapingProfileMap and FirewallShapingProfileMapOutput values.
// You can construct a concrete instance of `FirewallShapingProfileMapInput` via:
//
//          FirewallShapingProfileMap{ "key": FirewallShapingProfileArgs{...} }
type FirewallShapingProfileMapInput interface {
	pulumi.Input

	ToFirewallShapingProfileMapOutput() FirewallShapingProfileMapOutput
	ToFirewallShapingProfileMapOutputWithContext(context.Context) FirewallShapingProfileMapOutput
}

type FirewallShapingProfileMap map[string]FirewallShapingProfileInput

func (FirewallShapingProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallShapingProfile)(nil)).Elem()
}

func (i FirewallShapingProfileMap) ToFirewallShapingProfileMapOutput() FirewallShapingProfileMapOutput {
	return i.ToFirewallShapingProfileMapOutputWithContext(context.Background())
}

func (i FirewallShapingProfileMap) ToFirewallShapingProfileMapOutputWithContext(ctx context.Context) FirewallShapingProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShapingProfileMapOutput)
}

type FirewallShapingProfileOutput struct{ *pulumi.OutputState }

func (FirewallShapingProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallShapingProfile)(nil)).Elem()
}

func (o FirewallShapingProfileOutput) ToFirewallShapingProfileOutput() FirewallShapingProfileOutput {
	return o
}

func (o FirewallShapingProfileOutput) ToFirewallShapingProfileOutputWithContext(ctx context.Context) FirewallShapingProfileOutput {
	return o
}

type FirewallShapingProfileArrayOutput struct{ *pulumi.OutputState }

func (FirewallShapingProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallShapingProfile)(nil)).Elem()
}

func (o FirewallShapingProfileArrayOutput) ToFirewallShapingProfileArrayOutput() FirewallShapingProfileArrayOutput {
	return o
}

func (o FirewallShapingProfileArrayOutput) ToFirewallShapingProfileArrayOutputWithContext(ctx context.Context) FirewallShapingProfileArrayOutput {
	return o
}

func (o FirewallShapingProfileArrayOutput) Index(i pulumi.IntInput) FirewallShapingProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallShapingProfile {
		return vs[0].([]*FirewallShapingProfile)[vs[1].(int)]
	}).(FirewallShapingProfileOutput)
}

type FirewallShapingProfileMapOutput struct{ *pulumi.OutputState }

func (FirewallShapingProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallShapingProfile)(nil)).Elem()
}

func (o FirewallShapingProfileMapOutput) ToFirewallShapingProfileMapOutput() FirewallShapingProfileMapOutput {
	return o
}

func (o FirewallShapingProfileMapOutput) ToFirewallShapingProfileMapOutputWithContext(ctx context.Context) FirewallShapingProfileMapOutput {
	return o
}

func (o FirewallShapingProfileMapOutput) MapIndex(k pulumi.StringInput) FirewallShapingProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallShapingProfile {
		return vs[0].(map[string]*FirewallShapingProfile)[vs[1].(string)]
	}).(FirewallShapingProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShapingProfileInput)(nil)).Elem(), &FirewallShapingProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShapingProfileArrayInput)(nil)).Elem(), FirewallShapingProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShapingProfileMapInput)(nil)).Elem(), FirewallShapingProfileMap{})
	pulumi.RegisterOutputType(FirewallShapingProfileOutput{})
	pulumi.RegisterOutputType(FirewallShapingProfileArrayOutput{})
	pulumi.RegisterOutputType(FirewallShapingProfileMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure SCTP filter profiles. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// SctpFilter Profile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/sctpFilterProfile:SctpFilterProfile labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SctpFilterProfile struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// PPID filters list. The structure of `ppidFilters` block is documented below.
	PpidFilters SctpFilterProfilePpidFilterArrayOutput `pulumi:"ppidFilters"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSctpFilterProfile registers a new resource with the given unique name, arguments, and options.
func NewSctpFilterProfile(ctx *pulumi.Context,
	name string, args *SctpFilterProfileArgs, opts ...pulumi.ResourceOption) (*SctpFilterProfile, error) {
	if args == nil {
		args = &SctpFilterProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SctpFilterProfile
	err := ctx.RegisterResource("fortios:index/sctpFilterProfile:SctpFilterProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSctpFilterProfile gets an existing SctpFilterProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSctpFilterProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SctpFilterProfileState, opts ...pulumi.ResourceOption) (*SctpFilterProfile, error) {
	var resource SctpFilterProfile
	err := ctx.ReadResource("fortios:index/sctpFilterProfile:SctpFilterProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SctpFilterProfile resources.
type sctpFilterProfileState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Profile name.
	Name *string `pulumi:"name"`
	// PPID filters list. The structure of `ppidFilters` block is documented below.
	PpidFilters []SctpFilterProfilePpidFilter `pulumi:"ppidFilters"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SctpFilterProfileState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// PPID filters list. The structure of `ppidFilters` block is documented below.
	PpidFilters SctpFilterProfilePpidFilterArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SctpFilterProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*sctpFilterProfileState)(nil)).Elem()
}

type sctpFilterProfileArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Profile name.
	Name *string `pulumi:"name"`
	// PPID filters list. The structure of `ppidFilters` block is documented below.
	PpidFilters []SctpFilterProfilePpidFilter `pulumi:"ppidFilters"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SctpFilterProfile resource.
type SctpFilterProfileArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// PPID filters list. The structure of `ppidFilters` block is documented below.
	PpidFilters SctpFilterProfilePpidFilterArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SctpFilterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sctpFilterProfileArgs)(nil)).Elem()
}

type SctpFilterProfileInput interface {
	pulumi.Input

	ToSctpFilterProfileOutput() SctpFilterProfileOutput
	ToSctpFilterProfileOutputWithContext(ctx context.Context) SctpFilterProfileOutput
}

func (*SctpFilterProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**SctpFilterProfile)(nil)).Elem()
}

func (i *SctpFilterProfile) ToSctpFilterProfileOutput() SctpFilterProfileOutput {
	return i.ToSctpFilterProfileOutputWithContext(context.Background())
}

func (i *SctpFilterProfile) ToSctpFilterProfileOutputWithContext(ctx context.Context) SctpFilterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SctpFilterProfileOutput)
}

// SctpFilterProfileArrayInput is an input type that accepts SctpFilterProfileArray and SctpFilterProfileArrayOutput values.
// You can construct a concrete instance of `SctpFilterProfileArrayInput` via:
//
//          SctpFilterProfileArray{ SctpFilterProfileArgs{...} }
type SctpFilterProfileArrayInput interface {
	pulumi.Input

	ToSctpFilterProfileArrayOutput() SctpFilterProfileArrayOutput
	ToSctpFilterProfileArrayOutputWithContext(context.Context) SctpFilterProfileArrayOutput
}

type SctpFilterProfileArray []SctpFilterProfileInput

func (SctpFilterProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SctpFilterProfile)(nil)).Elem()
}

func (i SctpFilterProfileArray) ToSctpFilterProfileArrayOutput() SctpFilterProfileArrayOutput {
	return i.ToSctpFilterProfileArrayOutputWithContext(context.Background())
}

func (i SctpFilterProfileArray) ToSctpFilterProfileArrayOutputWithContext(ctx context.Context) SctpFilterProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SctpFilterProfileArrayOutput)
}

// SctpFilterProfileMapInput is an input type that accepts SctpFilterProfileMap and SctpFilterProfileMapOutput values.
// You can construct a concrete instance of `SctpFilterProfileMapInput` via:
//
//          SctpFilterProfileMap{ "key": SctpFilterProfileArgs{...} }
type SctpFilterProfileMapInput interface {
	pulumi.Input

	ToSctpFilterProfileMapOutput() SctpFilterProfileMapOutput
	ToSctpFilterProfileMapOutputWithContext(context.Context) SctpFilterProfileMapOutput
}

type SctpFilterProfileMap map[string]SctpFilterProfileInput

func (SctpFilterProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SctpFilterProfile)(nil)).Elem()
}

func (i SctpFilterProfileMap) ToSctpFilterProfileMapOutput() SctpFilterProfileMapOutput {
	return i.ToSctpFilterProfileMapOutputWithContext(context.Background())
}

func (i SctpFilterProfileMap) ToSctpFilterProfileMapOutputWithContext(ctx context.Context) SctpFilterProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SctpFilterProfileMapOutput)
}

type SctpFilterProfileOutput struct{ *pulumi.OutputState }

func (SctpFilterProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SctpFilterProfile)(nil)).Elem()
}

func (o SctpFilterProfileOutput) ToSctpFilterProfileOutput() SctpFilterProfileOutput {
	return o
}

func (o SctpFilterProfileOutput) ToSctpFilterProfileOutputWithContext(ctx context.Context) SctpFilterProfileOutput {
	return o
}

type SctpFilterProfileArrayOutput struct{ *pulumi.OutputState }

func (SctpFilterProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SctpFilterProfile)(nil)).Elem()
}

func (o SctpFilterProfileArrayOutput) ToSctpFilterProfileArrayOutput() SctpFilterProfileArrayOutput {
	return o
}

func (o SctpFilterProfileArrayOutput) ToSctpFilterProfileArrayOutputWithContext(ctx context.Context) SctpFilterProfileArrayOutput {
	return o
}

func (o SctpFilterProfileArrayOutput) Index(i pulumi.IntInput) SctpFilterProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SctpFilterProfile {
		return vs[0].([]*SctpFilterProfile)[vs[1].(int)]
	}).(SctpFilterProfileOutput)
}

type SctpFilterProfileMapOutput struct{ *pulumi.OutputState }

func (SctpFilterProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SctpFilterProfile)(nil)).Elem()
}

func (o SctpFilterProfileMapOutput) ToSctpFilterProfileMapOutput() SctpFilterProfileMapOutput {
	return o
}

func (o SctpFilterProfileMapOutput) ToSctpFilterProfileMapOutputWithContext(ctx context.Context) SctpFilterProfileMapOutput {
	return o
}

func (o SctpFilterProfileMapOutput) MapIndex(k pulumi.StringInput) SctpFilterProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SctpFilterProfile {
		return vs[0].(map[string]*SctpFilterProfile)[vs[1].(string)]
	}).(SctpFilterProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SctpFilterProfileInput)(nil)).Elem(), &SctpFilterProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*SctpFilterProfileArrayInput)(nil)).Elem(), SctpFilterProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SctpFilterProfileMapInput)(nil)).Elem(), SctpFilterProfileMap{})
	pulumi.RegisterOutputType(SctpFilterProfileOutput{})
	pulumi.RegisterOutputType(SctpFilterProfileArrayOutput{})
	pulumi.RegisterOutputType(SctpFilterProfileMapOutput{})
}

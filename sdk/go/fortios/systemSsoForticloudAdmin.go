// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiCloud SSO admin users. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// System SsoForticloudAdmin can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemSsoForticloudAdmin:SystemSsoForticloudAdmin labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemSsoForticloudAdmin:SystemSsoForticloudAdmin labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemSsoForticloudAdmin struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Virtual domain name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms SystemSsoForticloudAdminVdomArrayOutput `pulumi:"vdoms"`
}

// NewSystemSsoForticloudAdmin registers a new resource with the given unique name, arguments, and options.
func NewSystemSsoForticloudAdmin(ctx *pulumi.Context,
	name string, args *SystemSsoForticloudAdminArgs, opts ...pulumi.ResourceOption) (*SystemSsoForticloudAdmin, error) {
	if args == nil {
		args = &SystemSsoForticloudAdminArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSsoForticloudAdmin
	err := ctx.RegisterResource("fortios:index/systemSsoForticloudAdmin:SystemSsoForticloudAdmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSsoForticloudAdmin gets an existing SystemSsoForticloudAdmin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSsoForticloudAdmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSsoForticloudAdminState, opts ...pulumi.ResourceOption) (*SystemSsoForticloudAdmin, error) {
	var resource SystemSsoForticloudAdmin
	err := ctx.ReadResource("fortios:index/systemSsoForticloudAdmin:SystemSsoForticloudAdmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSsoForticloudAdmin resources.
type systemSsoForticloudAdminState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Virtual domain name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms []SystemSsoForticloudAdminVdom `pulumi:"vdoms"`
}

type SystemSsoForticloudAdminState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Virtual domain name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms SystemSsoForticloudAdminVdomArrayInput
}

func (SystemSsoForticloudAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSsoForticloudAdminState)(nil)).Elem()
}

type systemSsoForticloudAdminArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Virtual domain name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms []SystemSsoForticloudAdminVdom `pulumi:"vdoms"`
}

// The set of arguments for constructing a SystemSsoForticloudAdmin resource.
type SystemSsoForticloudAdminArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Virtual domain name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms SystemSsoForticloudAdminVdomArrayInput
}

func (SystemSsoForticloudAdminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSsoForticloudAdminArgs)(nil)).Elem()
}

type SystemSsoForticloudAdminInput interface {
	pulumi.Input

	ToSystemSsoForticloudAdminOutput() SystemSsoForticloudAdminOutput
	ToSystemSsoForticloudAdminOutputWithContext(ctx context.Context) SystemSsoForticloudAdminOutput
}

func (*SystemSsoForticloudAdmin) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSsoForticloudAdmin)(nil)).Elem()
}

func (i *SystemSsoForticloudAdmin) ToSystemSsoForticloudAdminOutput() SystemSsoForticloudAdminOutput {
	return i.ToSystemSsoForticloudAdminOutputWithContext(context.Background())
}

func (i *SystemSsoForticloudAdmin) ToSystemSsoForticloudAdminOutputWithContext(ctx context.Context) SystemSsoForticloudAdminOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSsoForticloudAdminOutput)
}

// SystemSsoForticloudAdminArrayInput is an input type that accepts SystemSsoForticloudAdminArray and SystemSsoForticloudAdminArrayOutput values.
// You can construct a concrete instance of `SystemSsoForticloudAdminArrayInput` via:
//
//          SystemSsoForticloudAdminArray{ SystemSsoForticloudAdminArgs{...} }
type SystemSsoForticloudAdminArrayInput interface {
	pulumi.Input

	ToSystemSsoForticloudAdminArrayOutput() SystemSsoForticloudAdminArrayOutput
	ToSystemSsoForticloudAdminArrayOutputWithContext(context.Context) SystemSsoForticloudAdminArrayOutput
}

type SystemSsoForticloudAdminArray []SystemSsoForticloudAdminInput

func (SystemSsoForticloudAdminArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSsoForticloudAdmin)(nil)).Elem()
}

func (i SystemSsoForticloudAdminArray) ToSystemSsoForticloudAdminArrayOutput() SystemSsoForticloudAdminArrayOutput {
	return i.ToSystemSsoForticloudAdminArrayOutputWithContext(context.Background())
}

func (i SystemSsoForticloudAdminArray) ToSystemSsoForticloudAdminArrayOutputWithContext(ctx context.Context) SystemSsoForticloudAdminArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSsoForticloudAdminArrayOutput)
}

// SystemSsoForticloudAdminMapInput is an input type that accepts SystemSsoForticloudAdminMap and SystemSsoForticloudAdminMapOutput values.
// You can construct a concrete instance of `SystemSsoForticloudAdminMapInput` via:
//
//          SystemSsoForticloudAdminMap{ "key": SystemSsoForticloudAdminArgs{...} }
type SystemSsoForticloudAdminMapInput interface {
	pulumi.Input

	ToSystemSsoForticloudAdminMapOutput() SystemSsoForticloudAdminMapOutput
	ToSystemSsoForticloudAdminMapOutputWithContext(context.Context) SystemSsoForticloudAdminMapOutput
}

type SystemSsoForticloudAdminMap map[string]SystemSsoForticloudAdminInput

func (SystemSsoForticloudAdminMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSsoForticloudAdmin)(nil)).Elem()
}

func (i SystemSsoForticloudAdminMap) ToSystemSsoForticloudAdminMapOutput() SystemSsoForticloudAdminMapOutput {
	return i.ToSystemSsoForticloudAdminMapOutputWithContext(context.Background())
}

func (i SystemSsoForticloudAdminMap) ToSystemSsoForticloudAdminMapOutputWithContext(ctx context.Context) SystemSsoForticloudAdminMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSsoForticloudAdminMapOutput)
}

type SystemSsoForticloudAdminOutput struct{ *pulumi.OutputState }

func (SystemSsoForticloudAdminOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSsoForticloudAdmin)(nil)).Elem()
}

func (o SystemSsoForticloudAdminOutput) ToSystemSsoForticloudAdminOutput() SystemSsoForticloudAdminOutput {
	return o
}

func (o SystemSsoForticloudAdminOutput) ToSystemSsoForticloudAdminOutputWithContext(ctx context.Context) SystemSsoForticloudAdminOutput {
	return o
}

type SystemSsoForticloudAdminArrayOutput struct{ *pulumi.OutputState }

func (SystemSsoForticloudAdminArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSsoForticloudAdmin)(nil)).Elem()
}

func (o SystemSsoForticloudAdminArrayOutput) ToSystemSsoForticloudAdminArrayOutput() SystemSsoForticloudAdminArrayOutput {
	return o
}

func (o SystemSsoForticloudAdminArrayOutput) ToSystemSsoForticloudAdminArrayOutputWithContext(ctx context.Context) SystemSsoForticloudAdminArrayOutput {
	return o
}

func (o SystemSsoForticloudAdminArrayOutput) Index(i pulumi.IntInput) SystemSsoForticloudAdminOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSsoForticloudAdmin {
		return vs[0].([]*SystemSsoForticloudAdmin)[vs[1].(int)]
	}).(SystemSsoForticloudAdminOutput)
}

type SystemSsoForticloudAdminMapOutput struct{ *pulumi.OutputState }

func (SystemSsoForticloudAdminMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSsoForticloudAdmin)(nil)).Elem()
}

func (o SystemSsoForticloudAdminMapOutput) ToSystemSsoForticloudAdminMapOutput() SystemSsoForticloudAdminMapOutput {
	return o
}

func (o SystemSsoForticloudAdminMapOutput) ToSystemSsoForticloudAdminMapOutputWithContext(ctx context.Context) SystemSsoForticloudAdminMapOutput {
	return o
}

func (o SystemSsoForticloudAdminMapOutput) MapIndex(k pulumi.StringInput) SystemSsoForticloudAdminOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSsoForticloudAdmin {
		return vs[0].(map[string]*SystemSsoForticloudAdmin)[vs[1].(string)]
	}).(SystemSsoForticloudAdminOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSsoForticloudAdminInput)(nil)).Elem(), &SystemSsoForticloudAdmin{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSsoForticloudAdminArrayInput)(nil)).Elem(), SystemSsoForticloudAdminArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSsoForticloudAdminMapInput)(nil)).Elem(), SystemSsoForticloudAdminMap{})
	pulumi.RegisterOutputType(SystemSsoForticloudAdminOutput{})
	pulumi.RegisterOutputType(SystemSsoForticloudAdminArrayOutput{})
	pulumi.RegisterOutputType(SystemSsoForticloudAdminMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Automation stitches.
//
// ## Import
//
// System AutomationStitch can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemAutomationStitch:SystemAutomationStitch labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemAutomationStitch:SystemAutomationStitch labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemAutomationStitch struct {
	pulumi.CustomResourceState

	// Configure stitch actions. The structure of `actions` block is documented below.
	Actions SystemAutomationStitchActionArrayOutput `pulumi:"actions"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
	Destinations SystemAutomationStitchDestinationArrayOutput `pulumi:"destinations"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Destination name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable this stitch. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Trigger name.
	Trigger pulumi.StringOutput `pulumi:"trigger"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemAutomationStitch registers a new resource with the given unique name, arguments, and options.
func NewSystemAutomationStitch(ctx *pulumi.Context,
	name string, args *SystemAutomationStitchArgs, opts ...pulumi.ResourceOption) (*SystemAutomationStitch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.Trigger == nil {
		return nil, errors.New("invalid value for required argument 'Trigger'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAutomationStitch
	err := ctx.RegisterResource("fortios:index/systemAutomationStitch:SystemAutomationStitch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAutomationStitch gets an existing SystemAutomationStitch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAutomationStitch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAutomationStitchState, opts ...pulumi.ResourceOption) (*SystemAutomationStitch, error) {
	var resource SystemAutomationStitch
	err := ctx.ReadResource("fortios:index/systemAutomationStitch:SystemAutomationStitch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAutomationStitch resources.
type systemAutomationStitchState struct {
	// Configure stitch actions. The structure of `actions` block is documented below.
	Actions []SystemAutomationStitchAction `pulumi:"actions"`
	// Description.
	Description *string `pulumi:"description"`
	// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
	Destinations []SystemAutomationStitchDestination `pulumi:"destinations"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Destination name.
	Name *string `pulumi:"name"`
	// Enable/disable this stitch. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Trigger name.
	Trigger *string `pulumi:"trigger"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemAutomationStitchState struct {
	// Configure stitch actions. The structure of `actions` block is documented below.
	Actions SystemAutomationStitchActionArrayInput
	// Description.
	Description pulumi.StringPtrInput
	// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
	Destinations SystemAutomationStitchDestinationArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Destination name.
	Name pulumi.StringPtrInput
	// Enable/disable this stitch. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Trigger name.
	Trigger pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAutomationStitchState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutomationStitchState)(nil)).Elem()
}

type systemAutomationStitchArgs struct {
	// Configure stitch actions. The structure of `actions` block is documented below.
	Actions []SystemAutomationStitchAction `pulumi:"actions"`
	// Description.
	Description *string `pulumi:"description"`
	// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
	Destinations []SystemAutomationStitchDestination `pulumi:"destinations"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Destination name.
	Name *string `pulumi:"name"`
	// Enable/disable this stitch. Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// Trigger name.
	Trigger string `pulumi:"trigger"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAutomationStitch resource.
type SystemAutomationStitchArgs struct {
	// Configure stitch actions. The structure of `actions` block is documented below.
	Actions SystemAutomationStitchActionArrayInput
	// Description.
	Description pulumi.StringPtrInput
	// Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
	Destinations SystemAutomationStitchDestinationArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Destination name.
	Name pulumi.StringPtrInput
	// Enable/disable this stitch. Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// Trigger name.
	Trigger pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAutomationStitchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutomationStitchArgs)(nil)).Elem()
}

type SystemAutomationStitchInput interface {
	pulumi.Input

	ToSystemAutomationStitchOutput() SystemAutomationStitchOutput
	ToSystemAutomationStitchOutputWithContext(ctx context.Context) SystemAutomationStitchOutput
}

func (*SystemAutomationStitch) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutomationStitch)(nil)).Elem()
}

func (i *SystemAutomationStitch) ToSystemAutomationStitchOutput() SystemAutomationStitchOutput {
	return i.ToSystemAutomationStitchOutputWithContext(context.Background())
}

func (i *SystemAutomationStitch) ToSystemAutomationStitchOutputWithContext(ctx context.Context) SystemAutomationStitchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationStitchOutput)
}

// SystemAutomationStitchArrayInput is an input type that accepts SystemAutomationStitchArray and SystemAutomationStitchArrayOutput values.
// You can construct a concrete instance of `SystemAutomationStitchArrayInput` via:
//
//          SystemAutomationStitchArray{ SystemAutomationStitchArgs{...} }
type SystemAutomationStitchArrayInput interface {
	pulumi.Input

	ToSystemAutomationStitchArrayOutput() SystemAutomationStitchArrayOutput
	ToSystemAutomationStitchArrayOutputWithContext(context.Context) SystemAutomationStitchArrayOutput
}

type SystemAutomationStitchArray []SystemAutomationStitchInput

func (SystemAutomationStitchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutomationStitch)(nil)).Elem()
}

func (i SystemAutomationStitchArray) ToSystemAutomationStitchArrayOutput() SystemAutomationStitchArrayOutput {
	return i.ToSystemAutomationStitchArrayOutputWithContext(context.Background())
}

func (i SystemAutomationStitchArray) ToSystemAutomationStitchArrayOutputWithContext(ctx context.Context) SystemAutomationStitchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationStitchArrayOutput)
}

// SystemAutomationStitchMapInput is an input type that accepts SystemAutomationStitchMap and SystemAutomationStitchMapOutput values.
// You can construct a concrete instance of `SystemAutomationStitchMapInput` via:
//
//          SystemAutomationStitchMap{ "key": SystemAutomationStitchArgs{...} }
type SystemAutomationStitchMapInput interface {
	pulumi.Input

	ToSystemAutomationStitchMapOutput() SystemAutomationStitchMapOutput
	ToSystemAutomationStitchMapOutputWithContext(context.Context) SystemAutomationStitchMapOutput
}

type SystemAutomationStitchMap map[string]SystemAutomationStitchInput

func (SystemAutomationStitchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutomationStitch)(nil)).Elem()
}

func (i SystemAutomationStitchMap) ToSystemAutomationStitchMapOutput() SystemAutomationStitchMapOutput {
	return i.ToSystemAutomationStitchMapOutputWithContext(context.Background())
}

func (i SystemAutomationStitchMap) ToSystemAutomationStitchMapOutputWithContext(ctx context.Context) SystemAutomationStitchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutomationStitchMapOutput)
}

type SystemAutomationStitchOutput struct{ *pulumi.OutputState }

func (SystemAutomationStitchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutomationStitch)(nil)).Elem()
}

func (o SystemAutomationStitchOutput) ToSystemAutomationStitchOutput() SystemAutomationStitchOutput {
	return o
}

func (o SystemAutomationStitchOutput) ToSystemAutomationStitchOutputWithContext(ctx context.Context) SystemAutomationStitchOutput {
	return o
}

type SystemAutomationStitchArrayOutput struct{ *pulumi.OutputState }

func (SystemAutomationStitchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutomationStitch)(nil)).Elem()
}

func (o SystemAutomationStitchArrayOutput) ToSystemAutomationStitchArrayOutput() SystemAutomationStitchArrayOutput {
	return o
}

func (o SystemAutomationStitchArrayOutput) ToSystemAutomationStitchArrayOutputWithContext(ctx context.Context) SystemAutomationStitchArrayOutput {
	return o
}

func (o SystemAutomationStitchArrayOutput) Index(i pulumi.IntInput) SystemAutomationStitchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAutomationStitch {
		return vs[0].([]*SystemAutomationStitch)[vs[1].(int)]
	}).(SystemAutomationStitchOutput)
}

type SystemAutomationStitchMapOutput struct{ *pulumi.OutputState }

func (SystemAutomationStitchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutomationStitch)(nil)).Elem()
}

func (o SystemAutomationStitchMapOutput) ToSystemAutomationStitchMapOutput() SystemAutomationStitchMapOutput {
	return o
}

func (o SystemAutomationStitchMapOutput) ToSystemAutomationStitchMapOutputWithContext(ctx context.Context) SystemAutomationStitchMapOutput {
	return o
}

func (o SystemAutomationStitchMapOutput) MapIndex(k pulumi.StringInput) SystemAutomationStitchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAutomationStitch {
		return vs[0].(map[string]*SystemAutomationStitch)[vs[1].(string)]
	}).(SystemAutomationStitchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutomationStitchInput)(nil)).Elem(), &SystemAutomationStitch{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutomationStitchArrayInput)(nil)).Elem(), SystemAutomationStitchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutomationStitchMapInput)(nil)).Elem(), SystemAutomationStitchMap{})
	pulumi.RegisterOutputType(SystemAutomationStitchOutput{})
	pulumi.RegisterOutputType(SystemAutomationStitchArrayOutput{})
	pulumi.RegisterOutputType(SystemAutomationStitchMapOutput{})
}

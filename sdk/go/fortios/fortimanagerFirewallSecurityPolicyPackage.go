// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports Create/Read/Update/Delete firewall security policypackage on FortiManager which could be installed to the FortiGate later
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFortimanagerFirewallSecurityPolicyPackage(ctx, "test1", &fortios.FortimanagerFirewallSecurityPolicyPackageArgs{
// 			Target: pulumi.String("FGVM64-test"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FortimanagerFirewallSecurityPolicyPackage struct {
	pulumi.CustomResourceState

	// Source ADOM name. default is 'root'
	Adom pulumi.StringPtrOutput `pulumi:"adom"`
	// Inspection Mode. Enum:[flow, proxy]. default is 'flow'
	InspectionMode pulumi.StringPtrOutput `pulumi:"inspectionMode"`
	// Security policy package name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The installation target.
	Target pulumi.StringPtrOutput `pulumi:"target"`
	// Vdom of managed device. default is 'root'
	Vdom pulumi.StringPtrOutput `pulumi:"vdom"`
}

// NewFortimanagerFirewallSecurityPolicyPackage registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerFirewallSecurityPolicyPackage(ctx *pulumi.Context,
	name string, args *FortimanagerFirewallSecurityPolicyPackageArgs, opts ...pulumi.ResourceOption) (*FortimanagerFirewallSecurityPolicyPackage, error) {
	if args == nil {
		args = &FortimanagerFirewallSecurityPolicyPackageArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FortimanagerFirewallSecurityPolicyPackage
	err := ctx.RegisterResource("fortios:index/fortimanagerFirewallSecurityPolicyPackage:FortimanagerFirewallSecurityPolicyPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerFirewallSecurityPolicyPackage gets an existing FortimanagerFirewallSecurityPolicyPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerFirewallSecurityPolicyPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerFirewallSecurityPolicyPackageState, opts ...pulumi.ResourceOption) (*FortimanagerFirewallSecurityPolicyPackage, error) {
	var resource FortimanagerFirewallSecurityPolicyPackage
	err := ctx.ReadResource("fortios:index/fortimanagerFirewallSecurityPolicyPackage:FortimanagerFirewallSecurityPolicyPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerFirewallSecurityPolicyPackage resources.
type fortimanagerFirewallSecurityPolicyPackageState struct {
	// Source ADOM name. default is 'root'
	Adom *string `pulumi:"adom"`
	// Inspection Mode. Enum:[flow, proxy]. default is 'flow'
	InspectionMode *string `pulumi:"inspectionMode"`
	// Security policy package name.
	Name *string `pulumi:"name"`
	// The installation target.
	Target *string `pulumi:"target"`
	// Vdom of managed device. default is 'root'
	Vdom *string `pulumi:"vdom"`
}

type FortimanagerFirewallSecurityPolicyPackageState struct {
	// Source ADOM name. default is 'root'
	Adom pulumi.StringPtrInput
	// Inspection Mode. Enum:[flow, proxy]. default is 'flow'
	InspectionMode pulumi.StringPtrInput
	// Security policy package name.
	Name pulumi.StringPtrInput
	// The installation target.
	Target pulumi.StringPtrInput
	// Vdom of managed device. default is 'root'
	Vdom pulumi.StringPtrInput
}

func (FortimanagerFirewallSecurityPolicyPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerFirewallSecurityPolicyPackageState)(nil)).Elem()
}

type fortimanagerFirewallSecurityPolicyPackageArgs struct {
	// Source ADOM name. default is 'root'
	Adom *string `pulumi:"adom"`
	// Inspection Mode. Enum:[flow, proxy]. default is 'flow'
	InspectionMode *string `pulumi:"inspectionMode"`
	// Security policy package name.
	Name *string `pulumi:"name"`
	// The installation target.
	Target *string `pulumi:"target"`
	// Vdom of managed device. default is 'root'
	Vdom *string `pulumi:"vdom"`
}

// The set of arguments for constructing a FortimanagerFirewallSecurityPolicyPackage resource.
type FortimanagerFirewallSecurityPolicyPackageArgs struct {
	// Source ADOM name. default is 'root'
	Adom pulumi.StringPtrInput
	// Inspection Mode. Enum:[flow, proxy]. default is 'flow'
	InspectionMode pulumi.StringPtrInput
	// Security policy package name.
	Name pulumi.StringPtrInput
	// The installation target.
	Target pulumi.StringPtrInput
	// Vdom of managed device. default is 'root'
	Vdom pulumi.StringPtrInput
}

func (FortimanagerFirewallSecurityPolicyPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerFirewallSecurityPolicyPackageArgs)(nil)).Elem()
}

type FortimanagerFirewallSecurityPolicyPackageInput interface {
	pulumi.Input

	ToFortimanagerFirewallSecurityPolicyPackageOutput() FortimanagerFirewallSecurityPolicyPackageOutput
	ToFortimanagerFirewallSecurityPolicyPackageOutputWithContext(ctx context.Context) FortimanagerFirewallSecurityPolicyPackageOutput
}

func (*FortimanagerFirewallSecurityPolicyPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerFirewallSecurityPolicyPackage)(nil)).Elem()
}

func (i *FortimanagerFirewallSecurityPolicyPackage) ToFortimanagerFirewallSecurityPolicyPackageOutput() FortimanagerFirewallSecurityPolicyPackageOutput {
	return i.ToFortimanagerFirewallSecurityPolicyPackageOutputWithContext(context.Background())
}

func (i *FortimanagerFirewallSecurityPolicyPackage) ToFortimanagerFirewallSecurityPolicyPackageOutputWithContext(ctx context.Context) FortimanagerFirewallSecurityPolicyPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallSecurityPolicyPackageOutput)
}

// FortimanagerFirewallSecurityPolicyPackageArrayInput is an input type that accepts FortimanagerFirewallSecurityPolicyPackageArray and FortimanagerFirewallSecurityPolicyPackageArrayOutput values.
// You can construct a concrete instance of `FortimanagerFirewallSecurityPolicyPackageArrayInput` via:
//
//          FortimanagerFirewallSecurityPolicyPackageArray{ FortimanagerFirewallSecurityPolicyPackageArgs{...} }
type FortimanagerFirewallSecurityPolicyPackageArrayInput interface {
	pulumi.Input

	ToFortimanagerFirewallSecurityPolicyPackageArrayOutput() FortimanagerFirewallSecurityPolicyPackageArrayOutput
	ToFortimanagerFirewallSecurityPolicyPackageArrayOutputWithContext(context.Context) FortimanagerFirewallSecurityPolicyPackageArrayOutput
}

type FortimanagerFirewallSecurityPolicyPackageArray []FortimanagerFirewallSecurityPolicyPackageInput

func (FortimanagerFirewallSecurityPolicyPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerFirewallSecurityPolicyPackage)(nil)).Elem()
}

func (i FortimanagerFirewallSecurityPolicyPackageArray) ToFortimanagerFirewallSecurityPolicyPackageArrayOutput() FortimanagerFirewallSecurityPolicyPackageArrayOutput {
	return i.ToFortimanagerFirewallSecurityPolicyPackageArrayOutputWithContext(context.Background())
}

func (i FortimanagerFirewallSecurityPolicyPackageArray) ToFortimanagerFirewallSecurityPolicyPackageArrayOutputWithContext(ctx context.Context) FortimanagerFirewallSecurityPolicyPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallSecurityPolicyPackageArrayOutput)
}

// FortimanagerFirewallSecurityPolicyPackageMapInput is an input type that accepts FortimanagerFirewallSecurityPolicyPackageMap and FortimanagerFirewallSecurityPolicyPackageMapOutput values.
// You can construct a concrete instance of `FortimanagerFirewallSecurityPolicyPackageMapInput` via:
//
//          FortimanagerFirewallSecurityPolicyPackageMap{ "key": FortimanagerFirewallSecurityPolicyPackageArgs{...} }
type FortimanagerFirewallSecurityPolicyPackageMapInput interface {
	pulumi.Input

	ToFortimanagerFirewallSecurityPolicyPackageMapOutput() FortimanagerFirewallSecurityPolicyPackageMapOutput
	ToFortimanagerFirewallSecurityPolicyPackageMapOutputWithContext(context.Context) FortimanagerFirewallSecurityPolicyPackageMapOutput
}

type FortimanagerFirewallSecurityPolicyPackageMap map[string]FortimanagerFirewallSecurityPolicyPackageInput

func (FortimanagerFirewallSecurityPolicyPackageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerFirewallSecurityPolicyPackage)(nil)).Elem()
}

func (i FortimanagerFirewallSecurityPolicyPackageMap) ToFortimanagerFirewallSecurityPolicyPackageMapOutput() FortimanagerFirewallSecurityPolicyPackageMapOutput {
	return i.ToFortimanagerFirewallSecurityPolicyPackageMapOutputWithContext(context.Background())
}

func (i FortimanagerFirewallSecurityPolicyPackageMap) ToFortimanagerFirewallSecurityPolicyPackageMapOutputWithContext(ctx context.Context) FortimanagerFirewallSecurityPolicyPackageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallSecurityPolicyPackageMapOutput)
}

type FortimanagerFirewallSecurityPolicyPackageOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallSecurityPolicyPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerFirewallSecurityPolicyPackage)(nil)).Elem()
}

func (o FortimanagerFirewallSecurityPolicyPackageOutput) ToFortimanagerFirewallSecurityPolicyPackageOutput() FortimanagerFirewallSecurityPolicyPackageOutput {
	return o
}

func (o FortimanagerFirewallSecurityPolicyPackageOutput) ToFortimanagerFirewallSecurityPolicyPackageOutputWithContext(ctx context.Context) FortimanagerFirewallSecurityPolicyPackageOutput {
	return o
}

type FortimanagerFirewallSecurityPolicyPackageArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallSecurityPolicyPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerFirewallSecurityPolicyPackage)(nil)).Elem()
}

func (o FortimanagerFirewallSecurityPolicyPackageArrayOutput) ToFortimanagerFirewallSecurityPolicyPackageArrayOutput() FortimanagerFirewallSecurityPolicyPackageArrayOutput {
	return o
}

func (o FortimanagerFirewallSecurityPolicyPackageArrayOutput) ToFortimanagerFirewallSecurityPolicyPackageArrayOutputWithContext(ctx context.Context) FortimanagerFirewallSecurityPolicyPackageArrayOutput {
	return o
}

func (o FortimanagerFirewallSecurityPolicyPackageArrayOutput) Index(i pulumi.IntInput) FortimanagerFirewallSecurityPolicyPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerFirewallSecurityPolicyPackage {
		return vs[0].([]*FortimanagerFirewallSecurityPolicyPackage)[vs[1].(int)]
	}).(FortimanagerFirewallSecurityPolicyPackageOutput)
}

type FortimanagerFirewallSecurityPolicyPackageMapOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallSecurityPolicyPackageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerFirewallSecurityPolicyPackage)(nil)).Elem()
}

func (o FortimanagerFirewallSecurityPolicyPackageMapOutput) ToFortimanagerFirewallSecurityPolicyPackageMapOutput() FortimanagerFirewallSecurityPolicyPackageMapOutput {
	return o
}

func (o FortimanagerFirewallSecurityPolicyPackageMapOutput) ToFortimanagerFirewallSecurityPolicyPackageMapOutputWithContext(ctx context.Context) FortimanagerFirewallSecurityPolicyPackageMapOutput {
	return o
}

func (o FortimanagerFirewallSecurityPolicyPackageMapOutput) MapIndex(k pulumi.StringInput) FortimanagerFirewallSecurityPolicyPackageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerFirewallSecurityPolicyPackage {
		return vs[0].(map[string]*FortimanagerFirewallSecurityPolicyPackage)[vs[1].(string)]
	}).(FortimanagerFirewallSecurityPolicyPackageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerFirewallSecurityPolicyPackageInput)(nil)).Elem(), &FortimanagerFirewallSecurityPolicyPackage{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerFirewallSecurityPolicyPackageArrayInput)(nil)).Elem(), FortimanagerFirewallSecurityPolicyPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerFirewallSecurityPolicyPackageMapInput)(nil)).Elem(), FortimanagerFirewallSecurityPolicyPackageMap{})
	pulumi.RegisterOutputType(FortimanagerFirewallSecurityPolicyPackageOutput{})
	pulumi.RegisterOutputType(FortimanagerFirewallSecurityPolicyPackageArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerFirewallSecurityPolicyPackageMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports modifying system admin setting for FortiManager.
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
// 		_, err := fortios.NewFortimanagerSystemAdmin(ctx, "test1", &fortios.FortimanagerSystemAdminArgs{
// 			HttpPort:    pulumi.Int(80),
// 			HttpsPort:   pulumi.Int(443),
// 			IdleTimeout: pulumi.Int(20),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FortimanagerSystemAdmin struct {
	pulumi.CustomResourceState

	// Http port.
	HttpPort pulumi.IntPtrOutput `pulumi:"httpPort"`
	// Https port.
	HttpsPort pulumi.IntPtrOutput `pulumi:"httpsPort"`
	// Idle Timeout(1-480 minute).
	IdleTimeout pulumi.IntPtrOutput `pulumi:"idleTimeout"`
}

// NewFortimanagerSystemAdmin registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemAdmin(ctx *pulumi.Context,
	name string, args *FortimanagerSystemAdminArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemAdmin, error) {
	if args == nil {
		args = &FortimanagerSystemAdminArgs{}
	}

	var resource FortimanagerSystemAdmin
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemAdmin:FortimanagerSystemAdmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemAdmin gets an existing FortimanagerSystemAdmin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemAdmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemAdminState, opts ...pulumi.ResourceOption) (*FortimanagerSystemAdmin, error) {
	var resource FortimanagerSystemAdmin
	err := ctx.ReadResource("fortios:index/fortimanagerSystemAdmin:FortimanagerSystemAdmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemAdmin resources.
type fortimanagerSystemAdminState struct {
	// Http port.
	HttpPort *int `pulumi:"httpPort"`
	// Https port.
	HttpsPort *int `pulumi:"httpsPort"`
	// Idle Timeout(1-480 minute).
	IdleTimeout *int `pulumi:"idleTimeout"`
}

type FortimanagerSystemAdminState struct {
	// Http port.
	HttpPort pulumi.IntPtrInput
	// Https port.
	HttpsPort pulumi.IntPtrInput
	// Idle Timeout(1-480 minute).
	IdleTimeout pulumi.IntPtrInput
}

func (FortimanagerSystemAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemAdminState)(nil)).Elem()
}

type fortimanagerSystemAdminArgs struct {
	// Http port.
	HttpPort *int `pulumi:"httpPort"`
	// Https port.
	HttpsPort *int `pulumi:"httpsPort"`
	// Idle Timeout(1-480 minute).
	IdleTimeout *int `pulumi:"idleTimeout"`
}

// The set of arguments for constructing a FortimanagerSystemAdmin resource.
type FortimanagerSystemAdminArgs struct {
	// Http port.
	HttpPort pulumi.IntPtrInput
	// Https port.
	HttpsPort pulumi.IntPtrInput
	// Idle Timeout(1-480 minute).
	IdleTimeout pulumi.IntPtrInput
}

func (FortimanagerSystemAdminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemAdminArgs)(nil)).Elem()
}

type FortimanagerSystemAdminInput interface {
	pulumi.Input

	ToFortimanagerSystemAdminOutput() FortimanagerSystemAdminOutput
	ToFortimanagerSystemAdminOutputWithContext(ctx context.Context) FortimanagerSystemAdminOutput
}

func (*FortimanagerSystemAdmin) ElementType() reflect.Type {
	return reflect.TypeOf((*FortimanagerSystemAdmin)(nil))
}

func (i *FortimanagerSystemAdmin) ToFortimanagerSystemAdminOutput() FortimanagerSystemAdminOutput {
	return i.ToFortimanagerSystemAdminOutputWithContext(context.Background())
}

func (i *FortimanagerSystemAdmin) ToFortimanagerSystemAdminOutputWithContext(ctx context.Context) FortimanagerSystemAdminOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminOutput)
}

func (i *FortimanagerSystemAdmin) ToFortimanagerSystemAdminPtrOutput() FortimanagerSystemAdminPtrOutput {
	return i.ToFortimanagerSystemAdminPtrOutputWithContext(context.Background())
}

func (i *FortimanagerSystemAdmin) ToFortimanagerSystemAdminPtrOutputWithContext(ctx context.Context) FortimanagerSystemAdminPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminPtrOutput)
}

type FortimanagerSystemAdminPtrInput interface {
	pulumi.Input

	ToFortimanagerSystemAdminPtrOutput() FortimanagerSystemAdminPtrOutput
	ToFortimanagerSystemAdminPtrOutputWithContext(ctx context.Context) FortimanagerSystemAdminPtrOutput
}

type fortimanagerSystemAdminPtrType FortimanagerSystemAdminArgs

func (*fortimanagerSystemAdminPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemAdmin)(nil))
}

func (i *fortimanagerSystemAdminPtrType) ToFortimanagerSystemAdminPtrOutput() FortimanagerSystemAdminPtrOutput {
	return i.ToFortimanagerSystemAdminPtrOutputWithContext(context.Background())
}

func (i *fortimanagerSystemAdminPtrType) ToFortimanagerSystemAdminPtrOutputWithContext(ctx context.Context) FortimanagerSystemAdminPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminPtrOutput)
}

// FortimanagerSystemAdminArrayInput is an input type that accepts FortimanagerSystemAdminArray and FortimanagerSystemAdminArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemAdminArrayInput` via:
//
//          FortimanagerSystemAdminArray{ FortimanagerSystemAdminArgs{...} }
type FortimanagerSystemAdminArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemAdminArrayOutput() FortimanagerSystemAdminArrayOutput
	ToFortimanagerSystemAdminArrayOutputWithContext(context.Context) FortimanagerSystemAdminArrayOutput
}

type FortimanagerSystemAdminArray []FortimanagerSystemAdminInput

func (FortimanagerSystemAdminArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FortimanagerSystemAdmin)(nil))
}

func (i FortimanagerSystemAdminArray) ToFortimanagerSystemAdminArrayOutput() FortimanagerSystemAdminArrayOutput {
	return i.ToFortimanagerSystemAdminArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemAdminArray) ToFortimanagerSystemAdminArrayOutputWithContext(ctx context.Context) FortimanagerSystemAdminArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminArrayOutput)
}

// FortimanagerSystemAdminMapInput is an input type that accepts FortimanagerSystemAdminMap and FortimanagerSystemAdminMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemAdminMapInput` via:
//
//          FortimanagerSystemAdminMap{ "key": FortimanagerSystemAdminArgs{...} }
type FortimanagerSystemAdminMapInput interface {
	pulumi.Input

	ToFortimanagerSystemAdminMapOutput() FortimanagerSystemAdminMapOutput
	ToFortimanagerSystemAdminMapOutputWithContext(context.Context) FortimanagerSystemAdminMapOutput
}

type FortimanagerSystemAdminMap map[string]FortimanagerSystemAdminInput

func (FortimanagerSystemAdminMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FortimanagerSystemAdmin)(nil))
}

func (i FortimanagerSystemAdminMap) ToFortimanagerSystemAdminMapOutput() FortimanagerSystemAdminMapOutput {
	return i.ToFortimanagerSystemAdminMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemAdminMap) ToFortimanagerSystemAdminMapOutputWithContext(ctx context.Context) FortimanagerSystemAdminMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminMapOutput)
}

type FortimanagerSystemAdminOutput struct {
	*pulumi.OutputState
}

func (FortimanagerSystemAdminOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FortimanagerSystemAdmin)(nil))
}

func (o FortimanagerSystemAdminOutput) ToFortimanagerSystemAdminOutput() FortimanagerSystemAdminOutput {
	return o
}

func (o FortimanagerSystemAdminOutput) ToFortimanagerSystemAdminOutputWithContext(ctx context.Context) FortimanagerSystemAdminOutput {
	return o
}

func (o FortimanagerSystemAdminOutput) ToFortimanagerSystemAdminPtrOutput() FortimanagerSystemAdminPtrOutput {
	return o.ToFortimanagerSystemAdminPtrOutputWithContext(context.Background())
}

func (o FortimanagerSystemAdminOutput) ToFortimanagerSystemAdminPtrOutputWithContext(ctx context.Context) FortimanagerSystemAdminPtrOutput {
	return o.ApplyT(func(v FortimanagerSystemAdmin) *FortimanagerSystemAdmin {
		return &v
	}).(FortimanagerSystemAdminPtrOutput)
}

type FortimanagerSystemAdminPtrOutput struct {
	*pulumi.OutputState
}

func (FortimanagerSystemAdminPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemAdmin)(nil))
}

func (o FortimanagerSystemAdminPtrOutput) ToFortimanagerSystemAdminPtrOutput() FortimanagerSystemAdminPtrOutput {
	return o
}

func (o FortimanagerSystemAdminPtrOutput) ToFortimanagerSystemAdminPtrOutputWithContext(ctx context.Context) FortimanagerSystemAdminPtrOutput {
	return o
}

type FortimanagerSystemAdminArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdminArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FortimanagerSystemAdmin)(nil))
}

func (o FortimanagerSystemAdminArrayOutput) ToFortimanagerSystemAdminArrayOutput() FortimanagerSystemAdminArrayOutput {
	return o
}

func (o FortimanagerSystemAdminArrayOutput) ToFortimanagerSystemAdminArrayOutputWithContext(ctx context.Context) FortimanagerSystemAdminArrayOutput {
	return o
}

func (o FortimanagerSystemAdminArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemAdminOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FortimanagerSystemAdmin {
		return vs[0].([]FortimanagerSystemAdmin)[vs[1].(int)]
	}).(FortimanagerSystemAdminOutput)
}

type FortimanagerSystemAdminMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdminMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FortimanagerSystemAdmin)(nil))
}

func (o FortimanagerSystemAdminMapOutput) ToFortimanagerSystemAdminMapOutput() FortimanagerSystemAdminMapOutput {
	return o
}

func (o FortimanagerSystemAdminMapOutput) ToFortimanagerSystemAdminMapOutputWithContext(ctx context.Context) FortimanagerSystemAdminMapOutput {
	return o
}

func (o FortimanagerSystemAdminMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemAdminOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FortimanagerSystemAdmin {
		return vs[0].(map[string]FortimanagerSystemAdmin)[vs[1].(string)]
	}).(FortimanagerSystemAdminOutput)
}

func init() {
	pulumi.RegisterOutputType(FortimanagerSystemAdminOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemAdminPtrOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemAdminArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemAdminMapOutput{})
}

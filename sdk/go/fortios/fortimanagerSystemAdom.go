// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports Create/Read/Update/Delete system adom for FortiManager.
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
// 		_, err := fortios.NewFortimanagerSystemAdom(ctx, "test1", &fortios.FortimanagerSystemAdomArgs{
// 			ActionWhenConflictsOccurDuringPolicyCheck:  pulumi.String("Continue"),
// 			AutoPushPolicyPackagesWhenDeviceBackOnline: pulumi.String("Enable"),
// 			CentralManagementFortiap:                   pulumi.Bool(true),
// 			CentralManagementSdwan:                     pulumi.Bool(false),
// 			CentralManagementVpn:                       pulumi.Bool(false),
// 			Mode:                                       pulumi.String("Normal"),
// 			PerformPolicyCheckBeforeEveryInstall:       pulumi.Bool(true),
// 			Status:                                     pulumi.Int(1),
// 			Type:                                       pulumi.String("FortiCarrier"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FortimanagerSystemAdom struct {
	pulumi.CustomResourceState

	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck pulumi.StringPtrOutput `pulumi:"actionWhenConflictsOccurDuringPolicyCheck"`
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline pulumi.StringPtrOutput `pulumi:"autoPushPolicyPackagesWhenDeviceBackOnline"`
	// True or False.
	CentralManagementFortiap pulumi.BoolPtrOutput `pulumi:"centralManagementFortiap"`
	// True or False.
	CentralManagementSdwan pulumi.BoolPtrOutput `pulumi:"centralManagementSdwan"`
	// True or False.
	CentralManagementVpn pulumi.BoolPtrOutput `pulumi:"centralManagementVpn"`
	// Adom mode: Normal or Backup.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// Administrative Domain name.
	Name pulumi.StringOutput `pulumi:"name"`
	// True or False.
	PerformPolicyCheckBeforeEveryInstall pulumi.BoolPtrOutput `pulumi:"performPolicyCheckBeforeEveryInstall"`
	// Adom status. 0 means off and 1 means on.
	Status pulumi.IntPtrOutput `pulumi:"status"`
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewFortimanagerSystemAdom registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemAdom(ctx *pulumi.Context,
	name string, args *FortimanagerSystemAdomArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemAdom, error) {
	if args == nil {
		args = &FortimanagerSystemAdomArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FortimanagerSystemAdom
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemAdom:FortimanagerSystemAdom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemAdom gets an existing FortimanagerSystemAdom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemAdom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemAdomState, opts ...pulumi.ResourceOption) (*FortimanagerSystemAdom, error) {
	var resource FortimanagerSystemAdom
	err := ctx.ReadResource("fortios:index/fortimanagerSystemAdom:FortimanagerSystemAdom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemAdom resources.
type fortimanagerSystemAdomState struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck *string `pulumi:"actionWhenConflictsOccurDuringPolicyCheck"`
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline *string `pulumi:"autoPushPolicyPackagesWhenDeviceBackOnline"`
	// True or False.
	CentralManagementFortiap *bool `pulumi:"centralManagementFortiap"`
	// True or False.
	CentralManagementSdwan *bool `pulumi:"centralManagementSdwan"`
	// True or False.
	CentralManagementVpn *bool `pulumi:"centralManagementVpn"`
	// Adom mode: Normal or Backup.
	Mode *string `pulumi:"mode"`
	// Administrative Domain name.
	Name *string `pulumi:"name"`
	// True or False.
	PerformPolicyCheckBeforeEveryInstall *bool `pulumi:"performPolicyCheckBeforeEveryInstall"`
	// Adom status. 0 means off and 1 means on.
	Status *int `pulumi:"status"`
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type *string `pulumi:"type"`
}

type FortimanagerSystemAdomState struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck pulumi.StringPtrInput
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline pulumi.StringPtrInput
	// True or False.
	CentralManagementFortiap pulumi.BoolPtrInput
	// True or False.
	CentralManagementSdwan pulumi.BoolPtrInput
	// True or False.
	CentralManagementVpn pulumi.BoolPtrInput
	// Adom mode: Normal or Backup.
	Mode pulumi.StringPtrInput
	// Administrative Domain name.
	Name pulumi.StringPtrInput
	// True or False.
	PerformPolicyCheckBeforeEveryInstall pulumi.BoolPtrInput
	// Adom status. 0 means off and 1 means on.
	Status pulumi.IntPtrInput
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type pulumi.StringPtrInput
}

func (FortimanagerSystemAdomState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemAdomState)(nil)).Elem()
}

type fortimanagerSystemAdomArgs struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck *string `pulumi:"actionWhenConflictsOccurDuringPolicyCheck"`
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline *string `pulumi:"autoPushPolicyPackagesWhenDeviceBackOnline"`
	// True or False.
	CentralManagementFortiap *bool `pulumi:"centralManagementFortiap"`
	// True or False.
	CentralManagementSdwan *bool `pulumi:"centralManagementSdwan"`
	// True or False.
	CentralManagementVpn *bool `pulumi:"centralManagementVpn"`
	// Adom mode: Normal or Backup.
	Mode *string `pulumi:"mode"`
	// Administrative Domain name.
	Name *string `pulumi:"name"`
	// True or False.
	PerformPolicyCheckBeforeEveryInstall *bool `pulumi:"performPolicyCheckBeforeEveryInstall"`
	// Adom status. 0 means off and 1 means on.
	Status *int `pulumi:"status"`
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a FortimanagerSystemAdom resource.
type FortimanagerSystemAdomArgs struct {
	// True or False.
	ActionWhenConflictsOccurDuringPolicyCheck pulumi.StringPtrInput
	// True or False.
	AutoPushPolicyPackagesWhenDeviceBackOnline pulumi.StringPtrInput
	// True or False.
	CentralManagementFortiap pulumi.BoolPtrInput
	// True or False.
	CentralManagementSdwan pulumi.BoolPtrInput
	// True or False.
	CentralManagementVpn pulumi.BoolPtrInput
	// Adom mode: Normal or Backup.
	Mode pulumi.StringPtrInput
	// Administrative Domain name.
	Name pulumi.StringPtrInput
	// True or False.
	PerformPolicyCheckBeforeEveryInstall pulumi.BoolPtrInput
	// Adom status. 0 means off and 1 means on.
	Status pulumi.IntPtrInput
	// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
	Type pulumi.StringPtrInput
}

func (FortimanagerSystemAdomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemAdomArgs)(nil)).Elem()
}

type FortimanagerSystemAdomInput interface {
	pulumi.Input

	ToFortimanagerSystemAdomOutput() FortimanagerSystemAdomOutput
	ToFortimanagerSystemAdomOutputWithContext(ctx context.Context) FortimanagerSystemAdomOutput
}

func (*FortimanagerSystemAdom) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemAdom)(nil)).Elem()
}

func (i *FortimanagerSystemAdom) ToFortimanagerSystemAdomOutput() FortimanagerSystemAdomOutput {
	return i.ToFortimanagerSystemAdomOutputWithContext(context.Background())
}

func (i *FortimanagerSystemAdom) ToFortimanagerSystemAdomOutputWithContext(ctx context.Context) FortimanagerSystemAdomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdomOutput)
}

// FortimanagerSystemAdomArrayInput is an input type that accepts FortimanagerSystemAdomArray and FortimanagerSystemAdomArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemAdomArrayInput` via:
//
//          FortimanagerSystemAdomArray{ FortimanagerSystemAdomArgs{...} }
type FortimanagerSystemAdomArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemAdomArrayOutput() FortimanagerSystemAdomArrayOutput
	ToFortimanagerSystemAdomArrayOutputWithContext(context.Context) FortimanagerSystemAdomArrayOutput
}

type FortimanagerSystemAdomArray []FortimanagerSystemAdomInput

func (FortimanagerSystemAdomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemAdom)(nil)).Elem()
}

func (i FortimanagerSystemAdomArray) ToFortimanagerSystemAdomArrayOutput() FortimanagerSystemAdomArrayOutput {
	return i.ToFortimanagerSystemAdomArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemAdomArray) ToFortimanagerSystemAdomArrayOutputWithContext(ctx context.Context) FortimanagerSystemAdomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdomArrayOutput)
}

// FortimanagerSystemAdomMapInput is an input type that accepts FortimanagerSystemAdomMap and FortimanagerSystemAdomMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemAdomMapInput` via:
//
//          FortimanagerSystemAdomMap{ "key": FortimanagerSystemAdomArgs{...} }
type FortimanagerSystemAdomMapInput interface {
	pulumi.Input

	ToFortimanagerSystemAdomMapOutput() FortimanagerSystemAdomMapOutput
	ToFortimanagerSystemAdomMapOutputWithContext(context.Context) FortimanagerSystemAdomMapOutput
}

type FortimanagerSystemAdomMap map[string]FortimanagerSystemAdomInput

func (FortimanagerSystemAdomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemAdom)(nil)).Elem()
}

func (i FortimanagerSystemAdomMap) ToFortimanagerSystemAdomMapOutput() FortimanagerSystemAdomMapOutput {
	return i.ToFortimanagerSystemAdomMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemAdomMap) ToFortimanagerSystemAdomMapOutputWithContext(ctx context.Context) FortimanagerSystemAdomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdomMapOutput)
}

type FortimanagerSystemAdomOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemAdom)(nil)).Elem()
}

func (o FortimanagerSystemAdomOutput) ToFortimanagerSystemAdomOutput() FortimanagerSystemAdomOutput {
	return o
}

func (o FortimanagerSystemAdomOutput) ToFortimanagerSystemAdomOutputWithContext(ctx context.Context) FortimanagerSystemAdomOutput {
	return o
}

type FortimanagerSystemAdomArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemAdom)(nil)).Elem()
}

func (o FortimanagerSystemAdomArrayOutput) ToFortimanagerSystemAdomArrayOutput() FortimanagerSystemAdomArrayOutput {
	return o
}

func (o FortimanagerSystemAdomArrayOutput) ToFortimanagerSystemAdomArrayOutputWithContext(ctx context.Context) FortimanagerSystemAdomArrayOutput {
	return o
}

func (o FortimanagerSystemAdomArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemAdomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerSystemAdom {
		return vs[0].([]*FortimanagerSystemAdom)[vs[1].(int)]
	}).(FortimanagerSystemAdomOutput)
}

type FortimanagerSystemAdomMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemAdom)(nil)).Elem()
}

func (o FortimanagerSystemAdomMapOutput) ToFortimanagerSystemAdomMapOutput() FortimanagerSystemAdomMapOutput {
	return o
}

func (o FortimanagerSystemAdomMapOutput) ToFortimanagerSystemAdomMapOutputWithContext(ctx context.Context) FortimanagerSystemAdomMapOutput {
	return o
}

func (o FortimanagerSystemAdomMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemAdomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerSystemAdom {
		return vs[0].(map[string]*FortimanagerSystemAdom)[vs[1].(string)]
	}).(FortimanagerSystemAdomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemAdomInput)(nil)).Elem(), &FortimanagerSystemAdom{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemAdomArrayInput)(nil)).Elem(), FortimanagerSystemAdomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemAdomMapInput)(nil)).Elem(), FortimanagerSystemAdomMap{})
	pulumi.RegisterOutputType(FortimanagerSystemAdomOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemAdomArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemAdomMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemWccp struct {
	pulumi.CustomResourceState

	AssignmentBucketFormat pulumi.StringOutput    `pulumi:"assignmentBucketFormat"`
	AssignmentDstaddrMask  pulumi.StringOutput    `pulumi:"assignmentDstaddrMask"`
	AssignmentMethod       pulumi.StringOutput    `pulumi:"assignmentMethod"`
	AssignmentSrcaddrMask  pulumi.StringOutput    `pulumi:"assignmentSrcaddrMask"`
	AssignmentWeight       pulumi.IntOutput       `pulumi:"assignmentWeight"`
	Authentication         pulumi.StringOutput    `pulumi:"authentication"`
	CacheEngineMethod      pulumi.StringOutput    `pulumi:"cacheEngineMethod"`
	CacheId                pulumi.StringOutput    `pulumi:"cacheId"`
	ForwardMethod          pulumi.StringOutput    `pulumi:"forwardMethod"`
	GroupAddress           pulumi.StringOutput    `pulumi:"groupAddress"`
	Password               pulumi.StringPtrOutput `pulumi:"password"`
	Ports                  pulumi.StringOutput    `pulumi:"ports"`
	PortsDefined           pulumi.StringOutput    `pulumi:"portsDefined"`
	PrimaryHash            pulumi.StringOutput    `pulumi:"primaryHash"`
	Priority               pulumi.IntOutput       `pulumi:"priority"`
	Protocol               pulumi.IntOutput       `pulumi:"protocol"`
	ReturnMethod           pulumi.StringOutput    `pulumi:"returnMethod"`
	RouterId               pulumi.StringOutput    `pulumi:"routerId"`
	RouterList             pulumi.StringOutput    `pulumi:"routerList"`
	ServerList             pulumi.StringOutput    `pulumi:"serverList"`
	ServerType             pulumi.StringOutput    `pulumi:"serverType"`
	ServiceId              pulumi.StringOutput    `pulumi:"serviceId"`
	ServiceType            pulumi.StringOutput    `pulumi:"serviceType"`
	Vdomparam              pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemWccp registers a new resource with the given unique name, arguments, and options.
func NewSystemWccp(ctx *pulumi.Context,
	name string, args *SystemWccpArgs, opts ...pulumi.ResourceOption) (*SystemWccp, error) {
	if args == nil {
		args = &SystemWccpArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemWccp
	err := ctx.RegisterResource("fortios:index/systemWccp:SystemWccp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemWccp gets an existing SystemWccp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemWccp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemWccpState, opts ...pulumi.ResourceOption) (*SystemWccp, error) {
	var resource SystemWccp
	err := ctx.ReadResource("fortios:index/systemWccp:SystemWccp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemWccp resources.
type systemWccpState struct {
	AssignmentBucketFormat *string `pulumi:"assignmentBucketFormat"`
	AssignmentDstaddrMask  *string `pulumi:"assignmentDstaddrMask"`
	AssignmentMethod       *string `pulumi:"assignmentMethod"`
	AssignmentSrcaddrMask  *string `pulumi:"assignmentSrcaddrMask"`
	AssignmentWeight       *int    `pulumi:"assignmentWeight"`
	Authentication         *string `pulumi:"authentication"`
	CacheEngineMethod      *string `pulumi:"cacheEngineMethod"`
	CacheId                *string `pulumi:"cacheId"`
	ForwardMethod          *string `pulumi:"forwardMethod"`
	GroupAddress           *string `pulumi:"groupAddress"`
	Password               *string `pulumi:"password"`
	Ports                  *string `pulumi:"ports"`
	PortsDefined           *string `pulumi:"portsDefined"`
	PrimaryHash            *string `pulumi:"primaryHash"`
	Priority               *int    `pulumi:"priority"`
	Protocol               *int    `pulumi:"protocol"`
	ReturnMethod           *string `pulumi:"returnMethod"`
	RouterId               *string `pulumi:"routerId"`
	RouterList             *string `pulumi:"routerList"`
	ServerList             *string `pulumi:"serverList"`
	ServerType             *string `pulumi:"serverType"`
	ServiceId              *string `pulumi:"serviceId"`
	ServiceType            *string `pulumi:"serviceType"`
	Vdomparam              *string `pulumi:"vdomparam"`
}

type SystemWccpState struct {
	AssignmentBucketFormat pulumi.StringPtrInput
	AssignmentDstaddrMask  pulumi.StringPtrInput
	AssignmentMethod       pulumi.StringPtrInput
	AssignmentSrcaddrMask  pulumi.StringPtrInput
	AssignmentWeight       pulumi.IntPtrInput
	Authentication         pulumi.StringPtrInput
	CacheEngineMethod      pulumi.StringPtrInput
	CacheId                pulumi.StringPtrInput
	ForwardMethod          pulumi.StringPtrInput
	GroupAddress           pulumi.StringPtrInput
	Password               pulumi.StringPtrInput
	Ports                  pulumi.StringPtrInput
	PortsDefined           pulumi.StringPtrInput
	PrimaryHash            pulumi.StringPtrInput
	Priority               pulumi.IntPtrInput
	Protocol               pulumi.IntPtrInput
	ReturnMethod           pulumi.StringPtrInput
	RouterId               pulumi.StringPtrInput
	RouterList             pulumi.StringPtrInput
	ServerList             pulumi.StringPtrInput
	ServerType             pulumi.StringPtrInput
	ServiceId              pulumi.StringPtrInput
	ServiceType            pulumi.StringPtrInput
	Vdomparam              pulumi.StringPtrInput
}

func (SystemWccpState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemWccpState)(nil)).Elem()
}

type systemWccpArgs struct {
	AssignmentBucketFormat *string `pulumi:"assignmentBucketFormat"`
	AssignmentDstaddrMask  *string `pulumi:"assignmentDstaddrMask"`
	AssignmentMethod       *string `pulumi:"assignmentMethod"`
	AssignmentSrcaddrMask  *string `pulumi:"assignmentSrcaddrMask"`
	AssignmentWeight       *int    `pulumi:"assignmentWeight"`
	Authentication         *string `pulumi:"authentication"`
	CacheEngineMethod      *string `pulumi:"cacheEngineMethod"`
	CacheId                *string `pulumi:"cacheId"`
	ForwardMethod          *string `pulumi:"forwardMethod"`
	GroupAddress           *string `pulumi:"groupAddress"`
	Password               *string `pulumi:"password"`
	Ports                  *string `pulumi:"ports"`
	PortsDefined           *string `pulumi:"portsDefined"`
	PrimaryHash            *string `pulumi:"primaryHash"`
	Priority               *int    `pulumi:"priority"`
	Protocol               *int    `pulumi:"protocol"`
	ReturnMethod           *string `pulumi:"returnMethod"`
	RouterId               *string `pulumi:"routerId"`
	RouterList             *string `pulumi:"routerList"`
	ServerList             *string `pulumi:"serverList"`
	ServerType             *string `pulumi:"serverType"`
	ServiceId              *string `pulumi:"serviceId"`
	ServiceType            *string `pulumi:"serviceType"`
	Vdomparam              *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemWccp resource.
type SystemWccpArgs struct {
	AssignmentBucketFormat pulumi.StringPtrInput
	AssignmentDstaddrMask  pulumi.StringPtrInput
	AssignmentMethod       pulumi.StringPtrInput
	AssignmentSrcaddrMask  pulumi.StringPtrInput
	AssignmentWeight       pulumi.IntPtrInput
	Authentication         pulumi.StringPtrInput
	CacheEngineMethod      pulumi.StringPtrInput
	CacheId                pulumi.StringPtrInput
	ForwardMethod          pulumi.StringPtrInput
	GroupAddress           pulumi.StringPtrInput
	Password               pulumi.StringPtrInput
	Ports                  pulumi.StringPtrInput
	PortsDefined           pulumi.StringPtrInput
	PrimaryHash            pulumi.StringPtrInput
	Priority               pulumi.IntPtrInput
	Protocol               pulumi.IntPtrInput
	ReturnMethod           pulumi.StringPtrInput
	RouterId               pulumi.StringPtrInput
	RouterList             pulumi.StringPtrInput
	ServerList             pulumi.StringPtrInput
	ServerType             pulumi.StringPtrInput
	ServiceId              pulumi.StringPtrInput
	ServiceType            pulumi.StringPtrInput
	Vdomparam              pulumi.StringPtrInput
}

func (SystemWccpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemWccpArgs)(nil)).Elem()
}

type SystemWccpInput interface {
	pulumi.Input

	ToSystemWccpOutput() SystemWccpOutput
	ToSystemWccpOutputWithContext(ctx context.Context) SystemWccpOutput
}

func (*SystemWccp) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemWccp)(nil)).Elem()
}

func (i *SystemWccp) ToSystemWccpOutput() SystemWccpOutput {
	return i.ToSystemWccpOutputWithContext(context.Background())
}

func (i *SystemWccp) ToSystemWccpOutputWithContext(ctx context.Context) SystemWccpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemWccpOutput)
}

// SystemWccpArrayInput is an input type that accepts SystemWccpArray and SystemWccpArrayOutput values.
// You can construct a concrete instance of `SystemWccpArrayInput` via:
//
//	SystemWccpArray{ SystemWccpArgs{...} }
type SystemWccpArrayInput interface {
	pulumi.Input

	ToSystemWccpArrayOutput() SystemWccpArrayOutput
	ToSystemWccpArrayOutputWithContext(context.Context) SystemWccpArrayOutput
}

type SystemWccpArray []SystemWccpInput

func (SystemWccpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemWccp)(nil)).Elem()
}

func (i SystemWccpArray) ToSystemWccpArrayOutput() SystemWccpArrayOutput {
	return i.ToSystemWccpArrayOutputWithContext(context.Background())
}

func (i SystemWccpArray) ToSystemWccpArrayOutputWithContext(ctx context.Context) SystemWccpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemWccpArrayOutput)
}

// SystemWccpMapInput is an input type that accepts SystemWccpMap and SystemWccpMapOutput values.
// You can construct a concrete instance of `SystemWccpMapInput` via:
//
//	SystemWccpMap{ "key": SystemWccpArgs{...} }
type SystemWccpMapInput interface {
	pulumi.Input

	ToSystemWccpMapOutput() SystemWccpMapOutput
	ToSystemWccpMapOutputWithContext(context.Context) SystemWccpMapOutput
}

type SystemWccpMap map[string]SystemWccpInput

func (SystemWccpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemWccp)(nil)).Elem()
}

func (i SystemWccpMap) ToSystemWccpMapOutput() SystemWccpMapOutput {
	return i.ToSystemWccpMapOutputWithContext(context.Background())
}

func (i SystemWccpMap) ToSystemWccpMapOutputWithContext(ctx context.Context) SystemWccpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemWccpMapOutput)
}

type SystemWccpOutput struct{ *pulumi.OutputState }

func (SystemWccpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemWccp)(nil)).Elem()
}

func (o SystemWccpOutput) ToSystemWccpOutput() SystemWccpOutput {
	return o
}

func (o SystemWccpOutput) ToSystemWccpOutputWithContext(ctx context.Context) SystemWccpOutput {
	return o
}

func (o SystemWccpOutput) AssignmentBucketFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.AssignmentBucketFormat }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) AssignmentDstaddrMask() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.AssignmentDstaddrMask }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) AssignmentMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.AssignmentMethod }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) AssignmentSrcaddrMask() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.AssignmentSrcaddrMask }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) AssignmentWeight() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.IntOutput { return v.AssignmentWeight }).(pulumi.IntOutput)
}

func (o SystemWccpOutput) Authentication() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.Authentication }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) CacheEngineMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.CacheEngineMethod }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) CacheId() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.CacheId }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) ForwardMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.ForwardMethod }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) GroupAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.GroupAddress }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o SystemWccpOutput) Ports() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.Ports }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) PortsDefined() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.PortsDefined }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) PrimaryHash() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.PrimaryHash }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

func (o SystemWccpOutput) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.IntOutput { return v.Protocol }).(pulumi.IntOutput)
}

func (o SystemWccpOutput) ReturnMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.ReturnMethod }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.RouterId }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) RouterList() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.RouterList }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) ServerList() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.ServerList }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.ServerType }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) ServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringOutput { return v.ServiceType }).(pulumi.StringOutput)
}

func (o SystemWccpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemWccp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemWccpArrayOutput struct{ *pulumi.OutputState }

func (SystemWccpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemWccp)(nil)).Elem()
}

func (o SystemWccpArrayOutput) ToSystemWccpArrayOutput() SystemWccpArrayOutput {
	return o
}

func (o SystemWccpArrayOutput) ToSystemWccpArrayOutputWithContext(ctx context.Context) SystemWccpArrayOutput {
	return o
}

func (o SystemWccpArrayOutput) Index(i pulumi.IntInput) SystemWccpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemWccp {
		return vs[0].([]*SystemWccp)[vs[1].(int)]
	}).(SystemWccpOutput)
}

type SystemWccpMapOutput struct{ *pulumi.OutputState }

func (SystemWccpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemWccp)(nil)).Elem()
}

func (o SystemWccpMapOutput) ToSystemWccpMapOutput() SystemWccpMapOutput {
	return o
}

func (o SystemWccpMapOutput) ToSystemWccpMapOutputWithContext(ctx context.Context) SystemWccpMapOutput {
	return o
}

func (o SystemWccpMapOutput) MapIndex(k pulumi.StringInput) SystemWccpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemWccp {
		return vs[0].(map[string]*SystemWccp)[vs[1].(string)]
	}).(SystemWccpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemWccpInput)(nil)).Elem(), &SystemWccp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemWccpArrayInput)(nil)).Elem(), SystemWccpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemWccpMapInput)(nil)).Elem(), SystemWccpMap{})
	pulumi.RegisterOutputType(SystemWccpOutput{})
	pulumi.RegisterOutputType(SystemWccpArrayOutput{})
	pulumi.RegisterOutputType(SystemWccpMapOutput{})
}

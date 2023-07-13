// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemStandaloneCluster struct {
	pulumi.CustomResourceState

	ClusterPeers        SystemStandaloneClusterClusterPeerArrayOutput `pulumi:"clusterPeers"`
	DynamicSortSubtable pulumi.StringPtrOutput                        `pulumi:"dynamicSortSubtable"`
	Encryption          pulumi.StringOutput                           `pulumi:"encryption"`
	GetAllTables        pulumi.StringPtrOutput                        `pulumi:"getAllTables"`
	GroupMemberId       pulumi.IntOutput                              `pulumi:"groupMemberId"`
	Layer2Connection    pulumi.StringOutput                           `pulumi:"layer2Connection"`
	Psksecret           pulumi.StringPtrOutput                        `pulumi:"psksecret"`
	SessionSyncDev      pulumi.StringOutput                           `pulumi:"sessionSyncDev"`
	StandaloneGroupId   pulumi.IntOutput                              `pulumi:"standaloneGroupId"`
	Vdomparam           pulumi.StringPtrOutput                        `pulumi:"vdomparam"`
}

// NewSystemStandaloneCluster registers a new resource with the given unique name, arguments, and options.
func NewSystemStandaloneCluster(ctx *pulumi.Context,
	name string, args *SystemStandaloneClusterArgs, opts ...pulumi.ResourceOption) (*SystemStandaloneCluster, error) {
	if args == nil {
		args = &SystemStandaloneClusterArgs{}
	}

	if args.Psksecret != nil {
		args.Psksecret = pulumi.ToSecret(args.Psksecret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"psksecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemStandaloneCluster
	err := ctx.RegisterResource("fortios:index/systemStandaloneCluster:SystemStandaloneCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemStandaloneCluster gets an existing SystemStandaloneCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemStandaloneCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemStandaloneClusterState, opts ...pulumi.ResourceOption) (*SystemStandaloneCluster, error) {
	var resource SystemStandaloneCluster
	err := ctx.ReadResource("fortios:index/systemStandaloneCluster:SystemStandaloneCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemStandaloneCluster resources.
type systemStandaloneClusterState struct {
	ClusterPeers        []SystemStandaloneClusterClusterPeer `pulumi:"clusterPeers"`
	DynamicSortSubtable *string                              `pulumi:"dynamicSortSubtable"`
	Encryption          *string                              `pulumi:"encryption"`
	GetAllTables        *string                              `pulumi:"getAllTables"`
	GroupMemberId       *int                                 `pulumi:"groupMemberId"`
	Layer2Connection    *string                              `pulumi:"layer2Connection"`
	Psksecret           *string                              `pulumi:"psksecret"`
	SessionSyncDev      *string                              `pulumi:"sessionSyncDev"`
	StandaloneGroupId   *int                                 `pulumi:"standaloneGroupId"`
	Vdomparam           *string                              `pulumi:"vdomparam"`
}

type SystemStandaloneClusterState struct {
	ClusterPeers        SystemStandaloneClusterClusterPeerArrayInput
	DynamicSortSubtable pulumi.StringPtrInput
	Encryption          pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	GroupMemberId       pulumi.IntPtrInput
	Layer2Connection    pulumi.StringPtrInput
	Psksecret           pulumi.StringPtrInput
	SessionSyncDev      pulumi.StringPtrInput
	StandaloneGroupId   pulumi.IntPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SystemStandaloneClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemStandaloneClusterState)(nil)).Elem()
}

type systemStandaloneClusterArgs struct {
	ClusterPeers        []SystemStandaloneClusterClusterPeer `pulumi:"clusterPeers"`
	DynamicSortSubtable *string                              `pulumi:"dynamicSortSubtable"`
	Encryption          *string                              `pulumi:"encryption"`
	GetAllTables        *string                              `pulumi:"getAllTables"`
	GroupMemberId       *int                                 `pulumi:"groupMemberId"`
	Layer2Connection    *string                              `pulumi:"layer2Connection"`
	Psksecret           *string                              `pulumi:"psksecret"`
	SessionSyncDev      *string                              `pulumi:"sessionSyncDev"`
	StandaloneGroupId   *int                                 `pulumi:"standaloneGroupId"`
	Vdomparam           *string                              `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemStandaloneCluster resource.
type SystemStandaloneClusterArgs struct {
	ClusterPeers        SystemStandaloneClusterClusterPeerArrayInput
	DynamicSortSubtable pulumi.StringPtrInput
	Encryption          pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	GroupMemberId       pulumi.IntPtrInput
	Layer2Connection    pulumi.StringPtrInput
	Psksecret           pulumi.StringPtrInput
	SessionSyncDev      pulumi.StringPtrInput
	StandaloneGroupId   pulumi.IntPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SystemStandaloneClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemStandaloneClusterArgs)(nil)).Elem()
}

type SystemStandaloneClusterInput interface {
	pulumi.Input

	ToSystemStandaloneClusterOutput() SystemStandaloneClusterOutput
	ToSystemStandaloneClusterOutputWithContext(ctx context.Context) SystemStandaloneClusterOutput
}

func (*SystemStandaloneCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemStandaloneCluster)(nil)).Elem()
}

func (i *SystemStandaloneCluster) ToSystemStandaloneClusterOutput() SystemStandaloneClusterOutput {
	return i.ToSystemStandaloneClusterOutputWithContext(context.Background())
}

func (i *SystemStandaloneCluster) ToSystemStandaloneClusterOutputWithContext(ctx context.Context) SystemStandaloneClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStandaloneClusterOutput)
}

// SystemStandaloneClusterArrayInput is an input type that accepts SystemStandaloneClusterArray and SystemStandaloneClusterArrayOutput values.
// You can construct a concrete instance of `SystemStandaloneClusterArrayInput` via:
//
//	SystemStandaloneClusterArray{ SystemStandaloneClusterArgs{...} }
type SystemStandaloneClusterArrayInput interface {
	pulumi.Input

	ToSystemStandaloneClusterArrayOutput() SystemStandaloneClusterArrayOutput
	ToSystemStandaloneClusterArrayOutputWithContext(context.Context) SystemStandaloneClusterArrayOutput
}

type SystemStandaloneClusterArray []SystemStandaloneClusterInput

func (SystemStandaloneClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemStandaloneCluster)(nil)).Elem()
}

func (i SystemStandaloneClusterArray) ToSystemStandaloneClusterArrayOutput() SystemStandaloneClusterArrayOutput {
	return i.ToSystemStandaloneClusterArrayOutputWithContext(context.Background())
}

func (i SystemStandaloneClusterArray) ToSystemStandaloneClusterArrayOutputWithContext(ctx context.Context) SystemStandaloneClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStandaloneClusterArrayOutput)
}

// SystemStandaloneClusterMapInput is an input type that accepts SystemStandaloneClusterMap and SystemStandaloneClusterMapOutput values.
// You can construct a concrete instance of `SystemStandaloneClusterMapInput` via:
//
//	SystemStandaloneClusterMap{ "key": SystemStandaloneClusterArgs{...} }
type SystemStandaloneClusterMapInput interface {
	pulumi.Input

	ToSystemStandaloneClusterMapOutput() SystemStandaloneClusterMapOutput
	ToSystemStandaloneClusterMapOutputWithContext(context.Context) SystemStandaloneClusterMapOutput
}

type SystemStandaloneClusterMap map[string]SystemStandaloneClusterInput

func (SystemStandaloneClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemStandaloneCluster)(nil)).Elem()
}

func (i SystemStandaloneClusterMap) ToSystemStandaloneClusterMapOutput() SystemStandaloneClusterMapOutput {
	return i.ToSystemStandaloneClusterMapOutputWithContext(context.Background())
}

func (i SystemStandaloneClusterMap) ToSystemStandaloneClusterMapOutputWithContext(ctx context.Context) SystemStandaloneClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStandaloneClusterMapOutput)
}

type SystemStandaloneClusterOutput struct{ *pulumi.OutputState }

func (SystemStandaloneClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemStandaloneCluster)(nil)).Elem()
}

func (o SystemStandaloneClusterOutput) ToSystemStandaloneClusterOutput() SystemStandaloneClusterOutput {
	return o
}

func (o SystemStandaloneClusterOutput) ToSystemStandaloneClusterOutputWithContext(ctx context.Context) SystemStandaloneClusterOutput {
	return o
}

func (o SystemStandaloneClusterOutput) ClusterPeers() SystemStandaloneClusterClusterPeerArrayOutput {
	return o.ApplyT(func(v *SystemStandaloneCluster) SystemStandaloneClusterClusterPeerArrayOutput { return v.ClusterPeers }).(SystemStandaloneClusterClusterPeerArrayOutput)
}

func (o SystemStandaloneClusterOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemStandaloneCluster) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SystemStandaloneClusterOutput) Encryption() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemStandaloneCluster) pulumi.StringOutput { return v.Encryption }).(pulumi.StringOutput)
}

func (o SystemStandaloneClusterOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemStandaloneCluster) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SystemStandaloneClusterOutput) GroupMemberId() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemStandaloneCluster) pulumi.IntOutput { return v.GroupMemberId }).(pulumi.IntOutput)
}

func (o SystemStandaloneClusterOutput) Layer2Connection() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemStandaloneCluster) pulumi.StringOutput { return v.Layer2Connection }).(pulumi.StringOutput)
}

func (o SystemStandaloneClusterOutput) Psksecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemStandaloneCluster) pulumi.StringPtrOutput { return v.Psksecret }).(pulumi.StringPtrOutput)
}

func (o SystemStandaloneClusterOutput) SessionSyncDev() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemStandaloneCluster) pulumi.StringOutput { return v.SessionSyncDev }).(pulumi.StringOutput)
}

func (o SystemStandaloneClusterOutput) StandaloneGroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemStandaloneCluster) pulumi.IntOutput { return v.StandaloneGroupId }).(pulumi.IntOutput)
}

func (o SystemStandaloneClusterOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemStandaloneCluster) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemStandaloneClusterArrayOutput struct{ *pulumi.OutputState }

func (SystemStandaloneClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemStandaloneCluster)(nil)).Elem()
}

func (o SystemStandaloneClusterArrayOutput) ToSystemStandaloneClusterArrayOutput() SystemStandaloneClusterArrayOutput {
	return o
}

func (o SystemStandaloneClusterArrayOutput) ToSystemStandaloneClusterArrayOutputWithContext(ctx context.Context) SystemStandaloneClusterArrayOutput {
	return o
}

func (o SystemStandaloneClusterArrayOutput) Index(i pulumi.IntInput) SystemStandaloneClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemStandaloneCluster {
		return vs[0].([]*SystemStandaloneCluster)[vs[1].(int)]
	}).(SystemStandaloneClusterOutput)
}

type SystemStandaloneClusterMapOutput struct{ *pulumi.OutputState }

func (SystemStandaloneClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemStandaloneCluster)(nil)).Elem()
}

func (o SystemStandaloneClusterMapOutput) ToSystemStandaloneClusterMapOutput() SystemStandaloneClusterMapOutput {
	return o
}

func (o SystemStandaloneClusterMapOutput) ToSystemStandaloneClusterMapOutputWithContext(ctx context.Context) SystemStandaloneClusterMapOutput {
	return o
}

func (o SystemStandaloneClusterMapOutput) MapIndex(k pulumi.StringInput) SystemStandaloneClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemStandaloneCluster {
		return vs[0].(map[string]*SystemStandaloneCluster)[vs[1].(string)]
	}).(SystemStandaloneClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStandaloneClusterInput)(nil)).Elem(), &SystemStandaloneCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStandaloneClusterArrayInput)(nil)).Elem(), SystemStandaloneClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStandaloneClusterMapInput)(nil)).Elem(), SystemStandaloneClusterMap{})
	pulumi.RegisterOutputType(SystemStandaloneClusterOutput{})
	pulumi.RegisterOutputType(SystemStandaloneClusterArrayOutput{})
	pulumi.RegisterOutputType(SystemStandaloneClusterMapOutput{})
}

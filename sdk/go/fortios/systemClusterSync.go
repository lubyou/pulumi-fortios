// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemClusterSync struct {
	pulumi.CustomResourceState

	DownIntfsBeforeSessSyncs SystemClusterSyncDownIntfsBeforeSessSyncArrayOutput `pulumi:"downIntfsBeforeSessSyncs"`
	DynamicSortSubtable      pulumi.StringPtrOutput                              `pulumi:"dynamicSortSubtable"`
	HbInterval               pulumi.IntOutput                                    `pulumi:"hbInterval"`
	HbLostThreshold          pulumi.IntOutput                                    `pulumi:"hbLostThreshold"`
	IkeHeartbeatInterval     pulumi.IntOutput                                    `pulumi:"ikeHeartbeatInterval"`
	IkeMonitor               pulumi.StringOutput                                 `pulumi:"ikeMonitor"`
	IkeMonitorInterval       pulumi.IntOutput                                    `pulumi:"ikeMonitorInterval"`
	IpsecTunnelSync          pulumi.StringOutput                                 `pulumi:"ipsecTunnelSync"`
	Peerip                   pulumi.StringOutput                                 `pulumi:"peerip"`
	Peervd                   pulumi.StringOutput                                 `pulumi:"peervd"`
	SecondaryAddIpsecRoutes  pulumi.StringOutput                                 `pulumi:"secondaryAddIpsecRoutes"`
	SessionSyncFilter        SystemClusterSyncSessionSyncFilterOutput            `pulumi:"sessionSyncFilter"`
	SlaveAddIkeRoutes        pulumi.StringOutput                                 `pulumi:"slaveAddIkeRoutes"`
	SyncId                   pulumi.IntOutput                                    `pulumi:"syncId"`
	Syncvds                  SystemClusterSyncSyncvdArrayOutput                  `pulumi:"syncvds"`
	Vdomparam                pulumi.StringPtrOutput                              `pulumi:"vdomparam"`
}

// NewSystemClusterSync registers a new resource with the given unique name, arguments, and options.
func NewSystemClusterSync(ctx *pulumi.Context,
	name string, args *SystemClusterSyncArgs, opts ...pulumi.ResourceOption) (*SystemClusterSync, error) {
	if args == nil {
		args = &SystemClusterSyncArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemClusterSync
	err := ctx.RegisterResource("fortios:index/systemClusterSync:SystemClusterSync", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemClusterSync gets an existing SystemClusterSync resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemClusterSync(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemClusterSyncState, opts ...pulumi.ResourceOption) (*SystemClusterSync, error) {
	var resource SystemClusterSync
	err := ctx.ReadResource("fortios:index/systemClusterSync:SystemClusterSync", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemClusterSync resources.
type systemClusterSyncState struct {
	DownIntfsBeforeSessSyncs []SystemClusterSyncDownIntfsBeforeSessSync `pulumi:"downIntfsBeforeSessSyncs"`
	DynamicSortSubtable      *string                                    `pulumi:"dynamicSortSubtable"`
	HbInterval               *int                                       `pulumi:"hbInterval"`
	HbLostThreshold          *int                                       `pulumi:"hbLostThreshold"`
	IkeHeartbeatInterval     *int                                       `pulumi:"ikeHeartbeatInterval"`
	IkeMonitor               *string                                    `pulumi:"ikeMonitor"`
	IkeMonitorInterval       *int                                       `pulumi:"ikeMonitorInterval"`
	IpsecTunnelSync          *string                                    `pulumi:"ipsecTunnelSync"`
	Peerip                   *string                                    `pulumi:"peerip"`
	Peervd                   *string                                    `pulumi:"peervd"`
	SecondaryAddIpsecRoutes  *string                                    `pulumi:"secondaryAddIpsecRoutes"`
	SessionSyncFilter        *SystemClusterSyncSessionSyncFilter        `pulumi:"sessionSyncFilter"`
	SlaveAddIkeRoutes        *string                                    `pulumi:"slaveAddIkeRoutes"`
	SyncId                   *int                                       `pulumi:"syncId"`
	Syncvds                  []SystemClusterSyncSyncvd                  `pulumi:"syncvds"`
	Vdomparam                *string                                    `pulumi:"vdomparam"`
}

type SystemClusterSyncState struct {
	DownIntfsBeforeSessSyncs SystemClusterSyncDownIntfsBeforeSessSyncArrayInput
	DynamicSortSubtable      pulumi.StringPtrInput
	HbInterval               pulumi.IntPtrInput
	HbLostThreshold          pulumi.IntPtrInput
	IkeHeartbeatInterval     pulumi.IntPtrInput
	IkeMonitor               pulumi.StringPtrInput
	IkeMonitorInterval       pulumi.IntPtrInput
	IpsecTunnelSync          pulumi.StringPtrInput
	Peerip                   pulumi.StringPtrInput
	Peervd                   pulumi.StringPtrInput
	SecondaryAddIpsecRoutes  pulumi.StringPtrInput
	SessionSyncFilter        SystemClusterSyncSessionSyncFilterPtrInput
	SlaveAddIkeRoutes        pulumi.StringPtrInput
	SyncId                   pulumi.IntPtrInput
	Syncvds                  SystemClusterSyncSyncvdArrayInput
	Vdomparam                pulumi.StringPtrInput
}

func (SystemClusterSyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemClusterSyncState)(nil)).Elem()
}

type systemClusterSyncArgs struct {
	DownIntfsBeforeSessSyncs []SystemClusterSyncDownIntfsBeforeSessSync `pulumi:"downIntfsBeforeSessSyncs"`
	DynamicSortSubtable      *string                                    `pulumi:"dynamicSortSubtable"`
	HbInterval               *int                                       `pulumi:"hbInterval"`
	HbLostThreshold          *int                                       `pulumi:"hbLostThreshold"`
	IkeHeartbeatInterval     *int                                       `pulumi:"ikeHeartbeatInterval"`
	IkeMonitor               *string                                    `pulumi:"ikeMonitor"`
	IkeMonitorInterval       *int                                       `pulumi:"ikeMonitorInterval"`
	IpsecTunnelSync          *string                                    `pulumi:"ipsecTunnelSync"`
	Peerip                   *string                                    `pulumi:"peerip"`
	Peervd                   *string                                    `pulumi:"peervd"`
	SecondaryAddIpsecRoutes  *string                                    `pulumi:"secondaryAddIpsecRoutes"`
	SessionSyncFilter        *SystemClusterSyncSessionSyncFilter        `pulumi:"sessionSyncFilter"`
	SlaveAddIkeRoutes        *string                                    `pulumi:"slaveAddIkeRoutes"`
	SyncId                   *int                                       `pulumi:"syncId"`
	Syncvds                  []SystemClusterSyncSyncvd                  `pulumi:"syncvds"`
	Vdomparam                *string                                    `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemClusterSync resource.
type SystemClusterSyncArgs struct {
	DownIntfsBeforeSessSyncs SystemClusterSyncDownIntfsBeforeSessSyncArrayInput
	DynamicSortSubtable      pulumi.StringPtrInput
	HbInterval               pulumi.IntPtrInput
	HbLostThreshold          pulumi.IntPtrInput
	IkeHeartbeatInterval     pulumi.IntPtrInput
	IkeMonitor               pulumi.StringPtrInput
	IkeMonitorInterval       pulumi.IntPtrInput
	IpsecTunnelSync          pulumi.StringPtrInput
	Peerip                   pulumi.StringPtrInput
	Peervd                   pulumi.StringPtrInput
	SecondaryAddIpsecRoutes  pulumi.StringPtrInput
	SessionSyncFilter        SystemClusterSyncSessionSyncFilterPtrInput
	SlaveAddIkeRoutes        pulumi.StringPtrInput
	SyncId                   pulumi.IntPtrInput
	Syncvds                  SystemClusterSyncSyncvdArrayInput
	Vdomparam                pulumi.StringPtrInput
}

func (SystemClusterSyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemClusterSyncArgs)(nil)).Elem()
}

type SystemClusterSyncInput interface {
	pulumi.Input

	ToSystemClusterSyncOutput() SystemClusterSyncOutput
	ToSystemClusterSyncOutputWithContext(ctx context.Context) SystemClusterSyncOutput
}

func (*SystemClusterSync) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemClusterSync)(nil)).Elem()
}

func (i *SystemClusterSync) ToSystemClusterSyncOutput() SystemClusterSyncOutput {
	return i.ToSystemClusterSyncOutputWithContext(context.Background())
}

func (i *SystemClusterSync) ToSystemClusterSyncOutputWithContext(ctx context.Context) SystemClusterSyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemClusterSyncOutput)
}

// SystemClusterSyncArrayInput is an input type that accepts SystemClusterSyncArray and SystemClusterSyncArrayOutput values.
// You can construct a concrete instance of `SystemClusterSyncArrayInput` via:
//
//	SystemClusterSyncArray{ SystemClusterSyncArgs{...} }
type SystemClusterSyncArrayInput interface {
	pulumi.Input

	ToSystemClusterSyncArrayOutput() SystemClusterSyncArrayOutput
	ToSystemClusterSyncArrayOutputWithContext(context.Context) SystemClusterSyncArrayOutput
}

type SystemClusterSyncArray []SystemClusterSyncInput

func (SystemClusterSyncArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemClusterSync)(nil)).Elem()
}

func (i SystemClusterSyncArray) ToSystemClusterSyncArrayOutput() SystemClusterSyncArrayOutput {
	return i.ToSystemClusterSyncArrayOutputWithContext(context.Background())
}

func (i SystemClusterSyncArray) ToSystemClusterSyncArrayOutputWithContext(ctx context.Context) SystemClusterSyncArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemClusterSyncArrayOutput)
}

// SystemClusterSyncMapInput is an input type that accepts SystemClusterSyncMap and SystemClusterSyncMapOutput values.
// You can construct a concrete instance of `SystemClusterSyncMapInput` via:
//
//	SystemClusterSyncMap{ "key": SystemClusterSyncArgs{...} }
type SystemClusterSyncMapInput interface {
	pulumi.Input

	ToSystemClusterSyncMapOutput() SystemClusterSyncMapOutput
	ToSystemClusterSyncMapOutputWithContext(context.Context) SystemClusterSyncMapOutput
}

type SystemClusterSyncMap map[string]SystemClusterSyncInput

func (SystemClusterSyncMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemClusterSync)(nil)).Elem()
}

func (i SystemClusterSyncMap) ToSystemClusterSyncMapOutput() SystemClusterSyncMapOutput {
	return i.ToSystemClusterSyncMapOutputWithContext(context.Background())
}

func (i SystemClusterSyncMap) ToSystemClusterSyncMapOutputWithContext(ctx context.Context) SystemClusterSyncMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemClusterSyncMapOutput)
}

type SystemClusterSyncOutput struct{ *pulumi.OutputState }

func (SystemClusterSyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemClusterSync)(nil)).Elem()
}

func (o SystemClusterSyncOutput) ToSystemClusterSyncOutput() SystemClusterSyncOutput {
	return o
}

func (o SystemClusterSyncOutput) ToSystemClusterSyncOutputWithContext(ctx context.Context) SystemClusterSyncOutput {
	return o
}

func (o SystemClusterSyncOutput) DownIntfsBeforeSessSyncs() SystemClusterSyncDownIntfsBeforeSessSyncArrayOutput {
	return o.ApplyT(func(v *SystemClusterSync) SystemClusterSyncDownIntfsBeforeSessSyncArrayOutput {
		return v.DownIntfsBeforeSessSyncs
	}).(SystemClusterSyncDownIntfsBeforeSessSyncArrayOutput)
}

func (o SystemClusterSyncOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SystemClusterSyncOutput) HbInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.IntOutput { return v.HbInterval }).(pulumi.IntOutput)
}

func (o SystemClusterSyncOutput) HbLostThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.IntOutput { return v.HbLostThreshold }).(pulumi.IntOutput)
}

func (o SystemClusterSyncOutput) IkeHeartbeatInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.IntOutput { return v.IkeHeartbeatInterval }).(pulumi.IntOutput)
}

func (o SystemClusterSyncOutput) IkeMonitor() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.StringOutput { return v.IkeMonitor }).(pulumi.StringOutput)
}

func (o SystemClusterSyncOutput) IkeMonitorInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.IntOutput { return v.IkeMonitorInterval }).(pulumi.IntOutput)
}

func (o SystemClusterSyncOutput) IpsecTunnelSync() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.StringOutput { return v.IpsecTunnelSync }).(pulumi.StringOutput)
}

func (o SystemClusterSyncOutput) Peerip() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.StringOutput { return v.Peerip }).(pulumi.StringOutput)
}

func (o SystemClusterSyncOutput) Peervd() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.StringOutput { return v.Peervd }).(pulumi.StringOutput)
}

func (o SystemClusterSyncOutput) SecondaryAddIpsecRoutes() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.StringOutput { return v.SecondaryAddIpsecRoutes }).(pulumi.StringOutput)
}

func (o SystemClusterSyncOutput) SessionSyncFilter() SystemClusterSyncSessionSyncFilterOutput {
	return o.ApplyT(func(v *SystemClusterSync) SystemClusterSyncSessionSyncFilterOutput { return v.SessionSyncFilter }).(SystemClusterSyncSessionSyncFilterOutput)
}

func (o SystemClusterSyncOutput) SlaveAddIkeRoutes() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.StringOutput { return v.SlaveAddIkeRoutes }).(pulumi.StringOutput)
}

func (o SystemClusterSyncOutput) SyncId() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.IntOutput { return v.SyncId }).(pulumi.IntOutput)
}

func (o SystemClusterSyncOutput) Syncvds() SystemClusterSyncSyncvdArrayOutput {
	return o.ApplyT(func(v *SystemClusterSync) SystemClusterSyncSyncvdArrayOutput { return v.Syncvds }).(SystemClusterSyncSyncvdArrayOutput)
}

func (o SystemClusterSyncOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemClusterSync) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemClusterSyncArrayOutput struct{ *pulumi.OutputState }

func (SystemClusterSyncArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemClusterSync)(nil)).Elem()
}

func (o SystemClusterSyncArrayOutput) ToSystemClusterSyncArrayOutput() SystemClusterSyncArrayOutput {
	return o
}

func (o SystemClusterSyncArrayOutput) ToSystemClusterSyncArrayOutputWithContext(ctx context.Context) SystemClusterSyncArrayOutput {
	return o
}

func (o SystemClusterSyncArrayOutput) Index(i pulumi.IntInput) SystemClusterSyncOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemClusterSync {
		return vs[0].([]*SystemClusterSync)[vs[1].(int)]
	}).(SystemClusterSyncOutput)
}

type SystemClusterSyncMapOutput struct{ *pulumi.OutputState }

func (SystemClusterSyncMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemClusterSync)(nil)).Elem()
}

func (o SystemClusterSyncMapOutput) ToSystemClusterSyncMapOutput() SystemClusterSyncMapOutput {
	return o
}

func (o SystemClusterSyncMapOutput) ToSystemClusterSyncMapOutputWithContext(ctx context.Context) SystemClusterSyncMapOutput {
	return o
}

func (o SystemClusterSyncMapOutput) MapIndex(k pulumi.StringInput) SystemClusterSyncOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemClusterSync {
		return vs[0].(map[string]*SystemClusterSync)[vs[1].(string)]
	}).(SystemClusterSyncOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemClusterSyncInput)(nil)).Elem(), &SystemClusterSync{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemClusterSyncArrayInput)(nil)).Elem(), SystemClusterSyncArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemClusterSyncMapInput)(nil)).Elem(), SystemClusterSyncMap{})
	pulumi.RegisterOutputType(SystemClusterSyncOutput{})
	pulumi.RegisterOutputType(SystemClusterSyncArrayOutput{})
	pulumi.RegisterOutputType(SystemClusterSyncMapOutput{})
}

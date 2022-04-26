// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.
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
// 		_, err := fortios.NewSystemClusterSync(ctx, "trname", &fortios.SystemClusterSyncArgs{
// 			HbInterval:        pulumi.Int(3),
// 			HbLostThreshold:   pulumi.Int(3),
// 			Peerip:            pulumi.String("1.1.1.1"),
// 			Peervd:            pulumi.String("root"),
// 			SlaveAddIkeRoutes: pulumi.String("enable"),
// 			SyncId:            pulumi.Int(1),
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
// System ClusterSync can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemClusterSync:SystemClusterSync labelname {{sync_id}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemClusterSync:SystemClusterSync labelname {{sync_id}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemClusterSync struct {
	pulumi.CustomResourceState

	// List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
	DownIntfsBeforeSessSyncs SystemClusterSyncDownIntfsBeforeSessSyncArrayOutput `pulumi:"downIntfsBeforeSessSyncs"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Heartbeat interval (1 - 10 sec).
	HbInterval pulumi.IntOutput `pulumi:"hbInterval"`
	// Lost heartbeat threshold (1 - 10).
	HbLostThreshold pulumi.IntOutput `pulumi:"hbLostThreshold"`
	// IKE heartbeat interval (1 - 60 secs).
	IkeHeartbeatInterval pulumi.IntOutput `pulumi:"ikeHeartbeatInterval"`
	// Enable/disable IKE HA monitor. Valid values: `enable`, `disable`.
	IkeMonitor pulumi.StringOutput `pulumi:"ikeMonitor"`
	// IKE HA monitor interval (10 - 300 secs).
	IkeMonitorInterval pulumi.IntOutput `pulumi:"ikeMonitorInterval"`
	// Enable/disable IPsec tunnel synchronization. Valid values: `enable`, `disable`.
	IpsecTunnelSync pulumi.StringOutput `pulumi:"ipsecTunnelSync"`
	// IP address of the interface on the peer unit that is used for the session synchronization link.
	Peerip pulumi.StringOutput `pulumi:"peerip"`
	// VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
	Peervd pulumi.StringOutput `pulumi:"peervd"`
	// Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
	SecondaryAddIpsecRoutes pulumi.StringOutput `pulumi:"secondaryAddIpsecRoutes"`
	// Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
	SessionSyncFilter SystemClusterSyncSessionSyncFilterPtrOutput `pulumi:"sessionSyncFilter"`
	// Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
	SlaveAddIkeRoutes pulumi.StringOutput `pulumi:"slaveAddIkeRoutes"`
	// Sync ID.
	SyncId pulumi.IntOutput `pulumi:"syncId"`
	// Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
	Syncvds SystemClusterSyncSyncvdArrayOutput `pulumi:"syncvds"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
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
	// List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
	DownIntfsBeforeSessSyncs []SystemClusterSyncDownIntfsBeforeSessSync `pulumi:"downIntfsBeforeSessSyncs"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Heartbeat interval (1 - 10 sec).
	HbInterval *int `pulumi:"hbInterval"`
	// Lost heartbeat threshold (1 - 10).
	HbLostThreshold *int `pulumi:"hbLostThreshold"`
	// IKE heartbeat interval (1 - 60 secs).
	IkeHeartbeatInterval *int `pulumi:"ikeHeartbeatInterval"`
	// Enable/disable IKE HA monitor. Valid values: `enable`, `disable`.
	IkeMonitor *string `pulumi:"ikeMonitor"`
	// IKE HA monitor interval (10 - 300 secs).
	IkeMonitorInterval *int `pulumi:"ikeMonitorInterval"`
	// Enable/disable IPsec tunnel synchronization. Valid values: `enable`, `disable`.
	IpsecTunnelSync *string `pulumi:"ipsecTunnelSync"`
	// IP address of the interface on the peer unit that is used for the session synchronization link.
	Peerip *string `pulumi:"peerip"`
	// VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
	Peervd *string `pulumi:"peervd"`
	// Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
	SecondaryAddIpsecRoutes *string `pulumi:"secondaryAddIpsecRoutes"`
	// Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
	SessionSyncFilter *SystemClusterSyncSessionSyncFilter `pulumi:"sessionSyncFilter"`
	// Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
	SlaveAddIkeRoutes *string `pulumi:"slaveAddIkeRoutes"`
	// Sync ID.
	SyncId *int `pulumi:"syncId"`
	// Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
	Syncvds []SystemClusterSyncSyncvd `pulumi:"syncvds"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemClusterSyncState struct {
	// List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
	DownIntfsBeforeSessSyncs SystemClusterSyncDownIntfsBeforeSessSyncArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Heartbeat interval (1 - 10 sec).
	HbInterval pulumi.IntPtrInput
	// Lost heartbeat threshold (1 - 10).
	HbLostThreshold pulumi.IntPtrInput
	// IKE heartbeat interval (1 - 60 secs).
	IkeHeartbeatInterval pulumi.IntPtrInput
	// Enable/disable IKE HA monitor. Valid values: `enable`, `disable`.
	IkeMonitor pulumi.StringPtrInput
	// IKE HA monitor interval (10 - 300 secs).
	IkeMonitorInterval pulumi.IntPtrInput
	// Enable/disable IPsec tunnel synchronization. Valid values: `enable`, `disable`.
	IpsecTunnelSync pulumi.StringPtrInput
	// IP address of the interface on the peer unit that is used for the session synchronization link.
	Peerip pulumi.StringPtrInput
	// VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
	Peervd pulumi.StringPtrInput
	// Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
	SecondaryAddIpsecRoutes pulumi.StringPtrInput
	// Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
	SessionSyncFilter SystemClusterSyncSessionSyncFilterPtrInput
	// Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
	SlaveAddIkeRoutes pulumi.StringPtrInput
	// Sync ID.
	SyncId pulumi.IntPtrInput
	// Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
	Syncvds SystemClusterSyncSyncvdArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemClusterSyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemClusterSyncState)(nil)).Elem()
}

type systemClusterSyncArgs struct {
	// List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
	DownIntfsBeforeSessSyncs []SystemClusterSyncDownIntfsBeforeSessSync `pulumi:"downIntfsBeforeSessSyncs"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Heartbeat interval (1 - 10 sec).
	HbInterval *int `pulumi:"hbInterval"`
	// Lost heartbeat threshold (1 - 10).
	HbLostThreshold *int `pulumi:"hbLostThreshold"`
	// IKE heartbeat interval (1 - 60 secs).
	IkeHeartbeatInterval *int `pulumi:"ikeHeartbeatInterval"`
	// Enable/disable IKE HA monitor. Valid values: `enable`, `disable`.
	IkeMonitor *string `pulumi:"ikeMonitor"`
	// IKE HA monitor interval (10 - 300 secs).
	IkeMonitorInterval *int `pulumi:"ikeMonitorInterval"`
	// Enable/disable IPsec tunnel synchronization. Valid values: `enable`, `disable`.
	IpsecTunnelSync *string `pulumi:"ipsecTunnelSync"`
	// IP address of the interface on the peer unit that is used for the session synchronization link.
	Peerip *string `pulumi:"peerip"`
	// VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
	Peervd *string `pulumi:"peervd"`
	// Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
	SecondaryAddIpsecRoutes *string `pulumi:"secondaryAddIpsecRoutes"`
	// Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
	SessionSyncFilter *SystemClusterSyncSessionSyncFilter `pulumi:"sessionSyncFilter"`
	// Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
	SlaveAddIkeRoutes *string `pulumi:"slaveAddIkeRoutes"`
	// Sync ID.
	SyncId *int `pulumi:"syncId"`
	// Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
	Syncvds []SystemClusterSyncSyncvd `pulumi:"syncvds"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemClusterSync resource.
type SystemClusterSyncArgs struct {
	// List of interfaces to be turned down before session synchronization is complete. The structure of `downIntfsBeforeSessSync` block is documented below.
	DownIntfsBeforeSessSyncs SystemClusterSyncDownIntfsBeforeSessSyncArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Heartbeat interval (1 - 10 sec).
	HbInterval pulumi.IntPtrInput
	// Lost heartbeat threshold (1 - 10).
	HbLostThreshold pulumi.IntPtrInput
	// IKE heartbeat interval (1 - 60 secs).
	IkeHeartbeatInterval pulumi.IntPtrInput
	// Enable/disable IKE HA monitor. Valid values: `enable`, `disable`.
	IkeMonitor pulumi.StringPtrInput
	// IKE HA monitor interval (10 - 300 secs).
	IkeMonitorInterval pulumi.IntPtrInput
	// Enable/disable IPsec tunnel synchronization. Valid values: `enable`, `disable`.
	IpsecTunnelSync pulumi.StringPtrInput
	// IP address of the interface on the peer unit that is used for the session synchronization link.
	Peerip pulumi.StringPtrInput
	// VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
	Peervd pulumi.StringPtrInput
	// Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
	SecondaryAddIpsecRoutes pulumi.StringPtrInput
	// Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `sessionSyncFilter` block is documented below.
	SessionSyncFilter SystemClusterSyncSessionSyncFilterPtrInput
	// Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
	SlaveAddIkeRoutes pulumi.StringPtrInput
	// Sync ID.
	SyncId pulumi.IntPtrInput
	// Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
	Syncvds SystemClusterSyncSyncvdArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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
//          SystemClusterSyncArray{ SystemClusterSyncArgs{...} }
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
//          SystemClusterSyncMap{ "key": SystemClusterSyncArgs{...} }
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

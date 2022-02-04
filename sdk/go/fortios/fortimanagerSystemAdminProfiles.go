// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports Create/Read/Update/Delete admin profiles for FortiManager
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
// 		_, err := fortios.NewFortimanagerSystemAdminProfiles(ctx, "test1", &fortios.FortimanagerSystemAdminProfilesArgs{
// 			AdomPolicyPackages:                  pulumi.String("read"),
// 			AdomSwitch:                          pulumi.String("read"),
// 			Assignment:                          pulumi.String("read"),
// 			ConfigRetrieve:                      pulumi.String("read"),
// 			ConfigRevert:                        pulumi.String("read"),
// 			ConsistencyCheck:                    pulumi.String("read-write"),
// 			DeployManagement:                    pulumi.String("read"),
// 			Description:                         pulumi.String("11"),
// 			DeviceAp:                            pulumi.String("none"),
// 			DeviceConfig:                        pulumi.String("read"),
// 			DeviceForticlient:                   pulumi.String("read"),
// 			DeviceFortiswitch:                   pulumi.String("read"),
// 			DeviceManager:                       pulumi.String("read-write"),
// 			DeviceOperation:                     pulumi.String("read"),
// 			DeviceProfile:                       pulumi.String("read"),
// 			DeviceRevisionDeletion:              pulumi.String("read"),
// 			DeviceWanLinkLoadBalance:            pulumi.String("read"),
// 			FortiguardCenter:                    pulumi.String("read"),
// 			FortiguardCenterAdvanced:            pulumi.String("read"),
// 			FortiguardCenterFirmwareManagerment: pulumi.String("read"),
// 			FortiguardCenterLicensing:           pulumi.String("read"),
// 			GlobalPolicyPackages:                pulumi.String("read-write"),
// 			ImportPolicyPackages:                pulumi.String("read"),
// 			IntfMapping:                         pulumi.String("read-write"),
// 			LogViewer:                           pulumi.String("read"),
// 			PolicyObjects:                       pulumi.String("read-write"),
// 			Profileid:                           pulumi.String("terraform-test1"),
// 			SetInstallTargets:                   pulumi.String("read-write"),
// 			SystemSetting:                       pulumi.String("read"),
// 			TerminalAccess:                      pulumi.String("read"),
// 			VpnManager:                          pulumi.String("read"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FortimanagerSystemAdminProfiles struct {
	pulumi.CustomResourceState

	// Adom policy packages.
	AdomPolicyPackages pulumi.StringPtrOutput `pulumi:"adomPolicyPackages"`
	// Administrator Domain.
	AdomSwitch pulumi.StringPtrOutput `pulumi:"adomSwitch"`
	// Assignment Permission.
	Assignment pulumi.StringPtrOutput `pulumi:"assignment"`
	// Configuration Retrieve.
	ConfigRetrieve pulumi.StringPtrOutput `pulumi:"configRetrieve"`
	// Revert Configuration from Revision History.
	ConfigRevert pulumi.StringPtrOutput `pulumi:"configRevert"`
	// Consistency check.
	ConsistencyCheck pulumi.StringPtrOutput `pulumi:"consistencyCheck"`
	// Install to devices.
	DeployManagement pulumi.StringPtrOutput `pulumi:"deployManagement"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Manage AP.
	DeviceAp pulumi.StringPtrOutput `pulumi:"deviceAp"`
	// Manage device configurations.
	DeviceConfig pulumi.StringPtrOutput `pulumi:"deviceConfig"`
	// Manage FortiClient.
	DeviceForticlient pulumi.StringPtrOutput `pulumi:"deviceForticlient"`
	// Manage FortiSwitch.
	DeviceFortiswitch pulumi.StringPtrOutput `pulumi:"deviceFortiswitch"`
	// Device Manager.
	DeviceManager pulumi.StringPtrOutput `pulumi:"deviceManager"`
	// Device add/delete/edit.
	DeviceOperation pulumi.StringPtrOutput `pulumi:"deviceOperation"`
	// Device profile permission.
	DeviceProfile pulumi.StringPtrOutput `pulumi:"deviceProfile"`
	// Delete device revision.
	DeviceRevisionDeletion pulumi.StringPtrOutput `pulumi:"deviceRevisionDeletion"`
	// Manage WAN link load balance.
	DeviceWanLinkLoadBalance pulumi.StringPtrOutput `pulumi:"deviceWanLinkLoadBalance"`
	// FortiGuard Center.
	FortiguardCenter pulumi.StringPtrOutput `pulumi:"fortiguardCenter"`
	// FortiGuard Center Advanced.
	FortiguardCenterAdvanced pulumi.StringPtrOutput `pulumi:"fortiguardCenterAdvanced"`
	// FortiGuard Center Firmware Managerment.
	FortiguardCenterFirmwareManagerment pulumi.StringPtrOutput `pulumi:"fortiguardCenterFirmwareManagerment"`
	// FortiGuard Center Licensing.
	FortiguardCenterLicensing pulumi.StringPtrOutput `pulumi:"fortiguardCenterLicensing"`
	// Global policy packages.
	GlobalPolicyPackages pulumi.StringPtrOutput `pulumi:"globalPolicyPackages"`
	// Import Policy Package.
	ImportPolicyPackages pulumi.StringPtrOutput `pulumi:"importPolicyPackages"`
	// Interface Mapping.
	IntfMapping pulumi.StringPtrOutput `pulumi:"intfMapping"`
	// Log Viewer.
	LogViewer pulumi.StringPtrOutput `pulumi:"logViewer"`
	// Policy objects permission.
	PolicyObjects pulumi.StringPtrOutput `pulumi:"policyObjects"`
	// Profile name.
	Profileid pulumi.StringOutput `pulumi:"profileid"`
	// Edit installation targets.
	SetInstallTargets pulumi.StringPtrOutput `pulumi:"setInstallTargets"`
	// System Setting.
	SystemSetting pulumi.StringPtrOutput `pulumi:"systemSetting"`
	// Terminal access.
	TerminalAccess pulumi.StringPtrOutput `pulumi:"terminalAccess"`
	// VPN Manager.
	VpnManager pulumi.StringPtrOutput `pulumi:"vpnManager"`
}

// NewFortimanagerSystemAdminProfiles registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemAdminProfiles(ctx *pulumi.Context,
	name string, args *FortimanagerSystemAdminProfilesArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemAdminProfiles, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Profileid == nil {
		return nil, errors.New("invalid value for required argument 'Profileid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FortimanagerSystemAdminProfiles
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemAdminProfiles:FortimanagerSystemAdminProfiles", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemAdminProfiles gets an existing FortimanagerSystemAdminProfiles resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemAdminProfiles(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemAdminProfilesState, opts ...pulumi.ResourceOption) (*FortimanagerSystemAdminProfiles, error) {
	var resource FortimanagerSystemAdminProfiles
	err := ctx.ReadResource("fortios:index/fortimanagerSystemAdminProfiles:FortimanagerSystemAdminProfiles", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemAdminProfiles resources.
type fortimanagerSystemAdminProfilesState struct {
	// Adom policy packages.
	AdomPolicyPackages *string `pulumi:"adomPolicyPackages"`
	// Administrator Domain.
	AdomSwitch *string `pulumi:"adomSwitch"`
	// Assignment Permission.
	Assignment *string `pulumi:"assignment"`
	// Configuration Retrieve.
	ConfigRetrieve *string `pulumi:"configRetrieve"`
	// Revert Configuration from Revision History.
	ConfigRevert *string `pulumi:"configRevert"`
	// Consistency check.
	ConsistencyCheck *string `pulumi:"consistencyCheck"`
	// Install to devices.
	DeployManagement *string `pulumi:"deployManagement"`
	// Description.
	Description *string `pulumi:"description"`
	// Manage AP.
	DeviceAp *string `pulumi:"deviceAp"`
	// Manage device configurations.
	DeviceConfig *string `pulumi:"deviceConfig"`
	// Manage FortiClient.
	DeviceForticlient *string `pulumi:"deviceForticlient"`
	// Manage FortiSwitch.
	DeviceFortiswitch *string `pulumi:"deviceFortiswitch"`
	// Device Manager.
	DeviceManager *string `pulumi:"deviceManager"`
	// Device add/delete/edit.
	DeviceOperation *string `pulumi:"deviceOperation"`
	// Device profile permission.
	DeviceProfile *string `pulumi:"deviceProfile"`
	// Delete device revision.
	DeviceRevisionDeletion *string `pulumi:"deviceRevisionDeletion"`
	// Manage WAN link load balance.
	DeviceWanLinkLoadBalance *string `pulumi:"deviceWanLinkLoadBalance"`
	// FortiGuard Center.
	FortiguardCenter *string `pulumi:"fortiguardCenter"`
	// FortiGuard Center Advanced.
	FortiguardCenterAdvanced *string `pulumi:"fortiguardCenterAdvanced"`
	// FortiGuard Center Firmware Managerment.
	FortiguardCenterFirmwareManagerment *string `pulumi:"fortiguardCenterFirmwareManagerment"`
	// FortiGuard Center Licensing.
	FortiguardCenterLicensing *string `pulumi:"fortiguardCenterLicensing"`
	// Global policy packages.
	GlobalPolicyPackages *string `pulumi:"globalPolicyPackages"`
	// Import Policy Package.
	ImportPolicyPackages *string `pulumi:"importPolicyPackages"`
	// Interface Mapping.
	IntfMapping *string `pulumi:"intfMapping"`
	// Log Viewer.
	LogViewer *string `pulumi:"logViewer"`
	// Policy objects permission.
	PolicyObjects *string `pulumi:"policyObjects"`
	// Profile name.
	Profileid *string `pulumi:"profileid"`
	// Edit installation targets.
	SetInstallTargets *string `pulumi:"setInstallTargets"`
	// System Setting.
	SystemSetting *string `pulumi:"systemSetting"`
	// Terminal access.
	TerminalAccess *string `pulumi:"terminalAccess"`
	// VPN Manager.
	VpnManager *string `pulumi:"vpnManager"`
}

type FortimanagerSystemAdminProfilesState struct {
	// Adom policy packages.
	AdomPolicyPackages pulumi.StringPtrInput
	// Administrator Domain.
	AdomSwitch pulumi.StringPtrInput
	// Assignment Permission.
	Assignment pulumi.StringPtrInput
	// Configuration Retrieve.
	ConfigRetrieve pulumi.StringPtrInput
	// Revert Configuration from Revision History.
	ConfigRevert pulumi.StringPtrInput
	// Consistency check.
	ConsistencyCheck pulumi.StringPtrInput
	// Install to devices.
	DeployManagement pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Manage AP.
	DeviceAp pulumi.StringPtrInput
	// Manage device configurations.
	DeviceConfig pulumi.StringPtrInput
	// Manage FortiClient.
	DeviceForticlient pulumi.StringPtrInput
	// Manage FortiSwitch.
	DeviceFortiswitch pulumi.StringPtrInput
	// Device Manager.
	DeviceManager pulumi.StringPtrInput
	// Device add/delete/edit.
	DeviceOperation pulumi.StringPtrInput
	// Device profile permission.
	DeviceProfile pulumi.StringPtrInput
	// Delete device revision.
	DeviceRevisionDeletion pulumi.StringPtrInput
	// Manage WAN link load balance.
	DeviceWanLinkLoadBalance pulumi.StringPtrInput
	// FortiGuard Center.
	FortiguardCenter pulumi.StringPtrInput
	// FortiGuard Center Advanced.
	FortiguardCenterAdvanced pulumi.StringPtrInput
	// FortiGuard Center Firmware Managerment.
	FortiguardCenterFirmwareManagerment pulumi.StringPtrInput
	// FortiGuard Center Licensing.
	FortiguardCenterLicensing pulumi.StringPtrInput
	// Global policy packages.
	GlobalPolicyPackages pulumi.StringPtrInput
	// Import Policy Package.
	ImportPolicyPackages pulumi.StringPtrInput
	// Interface Mapping.
	IntfMapping pulumi.StringPtrInput
	// Log Viewer.
	LogViewer pulumi.StringPtrInput
	// Policy objects permission.
	PolicyObjects pulumi.StringPtrInput
	// Profile name.
	Profileid pulumi.StringPtrInput
	// Edit installation targets.
	SetInstallTargets pulumi.StringPtrInput
	// System Setting.
	SystemSetting pulumi.StringPtrInput
	// Terminal access.
	TerminalAccess pulumi.StringPtrInput
	// VPN Manager.
	VpnManager pulumi.StringPtrInput
}

func (FortimanagerSystemAdminProfilesState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemAdminProfilesState)(nil)).Elem()
}

type fortimanagerSystemAdminProfilesArgs struct {
	// Adom policy packages.
	AdomPolicyPackages *string `pulumi:"adomPolicyPackages"`
	// Administrator Domain.
	AdomSwitch *string `pulumi:"adomSwitch"`
	// Assignment Permission.
	Assignment *string `pulumi:"assignment"`
	// Configuration Retrieve.
	ConfigRetrieve *string `pulumi:"configRetrieve"`
	// Revert Configuration from Revision History.
	ConfigRevert *string `pulumi:"configRevert"`
	// Consistency check.
	ConsistencyCheck *string `pulumi:"consistencyCheck"`
	// Install to devices.
	DeployManagement *string `pulumi:"deployManagement"`
	// Description.
	Description *string `pulumi:"description"`
	// Manage AP.
	DeviceAp *string `pulumi:"deviceAp"`
	// Manage device configurations.
	DeviceConfig *string `pulumi:"deviceConfig"`
	// Manage FortiClient.
	DeviceForticlient *string `pulumi:"deviceForticlient"`
	// Manage FortiSwitch.
	DeviceFortiswitch *string `pulumi:"deviceFortiswitch"`
	// Device Manager.
	DeviceManager *string `pulumi:"deviceManager"`
	// Device add/delete/edit.
	DeviceOperation *string `pulumi:"deviceOperation"`
	// Device profile permission.
	DeviceProfile *string `pulumi:"deviceProfile"`
	// Delete device revision.
	DeviceRevisionDeletion *string `pulumi:"deviceRevisionDeletion"`
	// Manage WAN link load balance.
	DeviceWanLinkLoadBalance *string `pulumi:"deviceWanLinkLoadBalance"`
	// FortiGuard Center.
	FortiguardCenter *string `pulumi:"fortiguardCenter"`
	// FortiGuard Center Advanced.
	FortiguardCenterAdvanced *string `pulumi:"fortiguardCenterAdvanced"`
	// FortiGuard Center Firmware Managerment.
	FortiguardCenterFirmwareManagerment *string `pulumi:"fortiguardCenterFirmwareManagerment"`
	// FortiGuard Center Licensing.
	FortiguardCenterLicensing *string `pulumi:"fortiguardCenterLicensing"`
	// Global policy packages.
	GlobalPolicyPackages *string `pulumi:"globalPolicyPackages"`
	// Import Policy Package.
	ImportPolicyPackages *string `pulumi:"importPolicyPackages"`
	// Interface Mapping.
	IntfMapping *string `pulumi:"intfMapping"`
	// Log Viewer.
	LogViewer *string `pulumi:"logViewer"`
	// Policy objects permission.
	PolicyObjects *string `pulumi:"policyObjects"`
	// Profile name.
	Profileid string `pulumi:"profileid"`
	// Edit installation targets.
	SetInstallTargets *string `pulumi:"setInstallTargets"`
	// System Setting.
	SystemSetting *string `pulumi:"systemSetting"`
	// Terminal access.
	TerminalAccess *string `pulumi:"terminalAccess"`
	// VPN Manager.
	VpnManager *string `pulumi:"vpnManager"`
}

// The set of arguments for constructing a FortimanagerSystemAdminProfiles resource.
type FortimanagerSystemAdminProfilesArgs struct {
	// Adom policy packages.
	AdomPolicyPackages pulumi.StringPtrInput
	// Administrator Domain.
	AdomSwitch pulumi.StringPtrInput
	// Assignment Permission.
	Assignment pulumi.StringPtrInput
	// Configuration Retrieve.
	ConfigRetrieve pulumi.StringPtrInput
	// Revert Configuration from Revision History.
	ConfigRevert pulumi.StringPtrInput
	// Consistency check.
	ConsistencyCheck pulumi.StringPtrInput
	// Install to devices.
	DeployManagement pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Manage AP.
	DeviceAp pulumi.StringPtrInput
	// Manage device configurations.
	DeviceConfig pulumi.StringPtrInput
	// Manage FortiClient.
	DeviceForticlient pulumi.StringPtrInput
	// Manage FortiSwitch.
	DeviceFortiswitch pulumi.StringPtrInput
	// Device Manager.
	DeviceManager pulumi.StringPtrInput
	// Device add/delete/edit.
	DeviceOperation pulumi.StringPtrInput
	// Device profile permission.
	DeviceProfile pulumi.StringPtrInput
	// Delete device revision.
	DeviceRevisionDeletion pulumi.StringPtrInput
	// Manage WAN link load balance.
	DeviceWanLinkLoadBalance pulumi.StringPtrInput
	// FortiGuard Center.
	FortiguardCenter pulumi.StringPtrInput
	// FortiGuard Center Advanced.
	FortiguardCenterAdvanced pulumi.StringPtrInput
	// FortiGuard Center Firmware Managerment.
	FortiguardCenterFirmwareManagerment pulumi.StringPtrInput
	// FortiGuard Center Licensing.
	FortiguardCenterLicensing pulumi.StringPtrInput
	// Global policy packages.
	GlobalPolicyPackages pulumi.StringPtrInput
	// Import Policy Package.
	ImportPolicyPackages pulumi.StringPtrInput
	// Interface Mapping.
	IntfMapping pulumi.StringPtrInput
	// Log Viewer.
	LogViewer pulumi.StringPtrInput
	// Policy objects permission.
	PolicyObjects pulumi.StringPtrInput
	// Profile name.
	Profileid pulumi.StringInput
	// Edit installation targets.
	SetInstallTargets pulumi.StringPtrInput
	// System Setting.
	SystemSetting pulumi.StringPtrInput
	// Terminal access.
	TerminalAccess pulumi.StringPtrInput
	// VPN Manager.
	VpnManager pulumi.StringPtrInput
}

func (FortimanagerSystemAdminProfilesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemAdminProfilesArgs)(nil)).Elem()
}

type FortimanagerSystemAdminProfilesInput interface {
	pulumi.Input

	ToFortimanagerSystemAdminProfilesOutput() FortimanagerSystemAdminProfilesOutput
	ToFortimanagerSystemAdminProfilesOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesOutput
}

func (*FortimanagerSystemAdminProfiles) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (i *FortimanagerSystemAdminProfiles) ToFortimanagerSystemAdminProfilesOutput() FortimanagerSystemAdminProfilesOutput {
	return i.ToFortimanagerSystemAdminProfilesOutputWithContext(context.Background())
}

func (i *FortimanagerSystemAdminProfiles) ToFortimanagerSystemAdminProfilesOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminProfilesOutput)
}

// FortimanagerSystemAdminProfilesArrayInput is an input type that accepts FortimanagerSystemAdminProfilesArray and FortimanagerSystemAdminProfilesArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemAdminProfilesArrayInput` via:
//
//          FortimanagerSystemAdminProfilesArray{ FortimanagerSystemAdminProfilesArgs{...} }
type FortimanagerSystemAdminProfilesArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemAdminProfilesArrayOutput() FortimanagerSystemAdminProfilesArrayOutput
	ToFortimanagerSystemAdminProfilesArrayOutputWithContext(context.Context) FortimanagerSystemAdminProfilesArrayOutput
}

type FortimanagerSystemAdminProfilesArray []FortimanagerSystemAdminProfilesInput

func (FortimanagerSystemAdminProfilesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (i FortimanagerSystemAdminProfilesArray) ToFortimanagerSystemAdminProfilesArrayOutput() FortimanagerSystemAdminProfilesArrayOutput {
	return i.ToFortimanagerSystemAdminProfilesArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemAdminProfilesArray) ToFortimanagerSystemAdminProfilesArrayOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminProfilesArrayOutput)
}

// FortimanagerSystemAdminProfilesMapInput is an input type that accepts FortimanagerSystemAdminProfilesMap and FortimanagerSystemAdminProfilesMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemAdminProfilesMapInput` via:
//
//          FortimanagerSystemAdminProfilesMap{ "key": FortimanagerSystemAdminProfilesArgs{...} }
type FortimanagerSystemAdminProfilesMapInput interface {
	pulumi.Input

	ToFortimanagerSystemAdminProfilesMapOutput() FortimanagerSystemAdminProfilesMapOutput
	ToFortimanagerSystemAdminProfilesMapOutputWithContext(context.Context) FortimanagerSystemAdminProfilesMapOutput
}

type FortimanagerSystemAdminProfilesMap map[string]FortimanagerSystemAdminProfilesInput

func (FortimanagerSystemAdminProfilesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (i FortimanagerSystemAdminProfilesMap) ToFortimanagerSystemAdminProfilesMapOutput() FortimanagerSystemAdminProfilesMapOutput {
	return i.ToFortimanagerSystemAdminProfilesMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemAdminProfilesMap) ToFortimanagerSystemAdminProfilesMapOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemAdminProfilesMapOutput)
}

type FortimanagerSystemAdminProfilesOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdminProfilesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (o FortimanagerSystemAdminProfilesOutput) ToFortimanagerSystemAdminProfilesOutput() FortimanagerSystemAdminProfilesOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesOutput) ToFortimanagerSystemAdminProfilesOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesOutput {
	return o
}

type FortimanagerSystemAdminProfilesArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdminProfilesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (o FortimanagerSystemAdminProfilesArrayOutput) ToFortimanagerSystemAdminProfilesArrayOutput() FortimanagerSystemAdminProfilesArrayOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesArrayOutput) ToFortimanagerSystemAdminProfilesArrayOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesArrayOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemAdminProfilesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerSystemAdminProfiles {
		return vs[0].([]*FortimanagerSystemAdminProfiles)[vs[1].(int)]
	}).(FortimanagerSystemAdminProfilesOutput)
}

type FortimanagerSystemAdminProfilesMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemAdminProfilesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemAdminProfiles)(nil)).Elem()
}

func (o FortimanagerSystemAdminProfilesMapOutput) ToFortimanagerSystemAdminProfilesMapOutput() FortimanagerSystemAdminProfilesMapOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesMapOutput) ToFortimanagerSystemAdminProfilesMapOutputWithContext(ctx context.Context) FortimanagerSystemAdminProfilesMapOutput {
	return o
}

func (o FortimanagerSystemAdminProfilesMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemAdminProfilesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerSystemAdminProfiles {
		return vs[0].(map[string]*FortimanagerSystemAdminProfiles)[vs[1].(string)]
	}).(FortimanagerSystemAdminProfilesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemAdminProfilesInput)(nil)).Elem(), &FortimanagerSystemAdminProfiles{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemAdminProfilesArrayInput)(nil)).Elem(), FortimanagerSystemAdminProfilesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemAdminProfilesMapInput)(nil)).Elem(), FortimanagerSystemAdminProfilesMap{})
	pulumi.RegisterOutputType(FortimanagerSystemAdminProfilesOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemAdminProfilesArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemAdminProfilesMapOutput{})
}

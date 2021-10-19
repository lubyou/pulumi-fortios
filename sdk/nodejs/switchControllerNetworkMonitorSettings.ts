// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure network monitor settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SwitchControllerNetworkMonitorSettings("trname", {
 *     networkMonitoring: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * SwitchController NetworkMonitorSettings can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerNetworkMonitorSettings:SwitchControllerNetworkMonitorSettings labelname SwitchControllerNetworkMonitorSettings
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerNetworkMonitorSettings extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerNetworkMonitorSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerNetworkMonitorSettingsState, opts?: pulumi.CustomResourceOptions): SwitchControllerNetworkMonitorSettings {
        return new SwitchControllerNetworkMonitorSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerNetworkMonitorSettings:SwitchControllerNetworkMonitorSettings';

    /**
     * Returns true if the given object is an instance of SwitchControllerNetworkMonitorSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerNetworkMonitorSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerNetworkMonitorSettings.__pulumiType;
    }

    /**
     * Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
     */
    public readonly networkMonitoring!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerNetworkMonitorSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerNetworkMonitorSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerNetworkMonitorSettingsArgs | SwitchControllerNetworkMonitorSettingsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerNetworkMonitorSettingsState | undefined;
            inputs["networkMonitoring"] = state ? state.networkMonitoring : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerNetworkMonitorSettingsArgs | undefined;
            inputs["networkMonitoring"] = args ? args.networkMonitoring : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerNetworkMonitorSettings.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerNetworkMonitorSettings resources.
 */
export interface SwitchControllerNetworkMonitorSettingsState {
    /**
     * Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
     */
    networkMonitoring?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerNetworkMonitorSettings resource.
 */
export interface SwitchControllerNetworkMonitorSettingsArgs {
    /**
     * Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
     */
    networkMonitoring?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

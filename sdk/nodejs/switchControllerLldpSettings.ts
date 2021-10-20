// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch LLDP settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SwitchControllerLldpSettings("trname", {
 *     fastStartInterval: 2,
 *     managementInterface: "internal",
 *     status: "enable",
 *     txHold: 4,
 *     txInterval: 30,
 * });
 * ```
 *
 * ## Import
 *
 * SwitchController LldpSettings can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerLldpSettings:SwitchControllerLldpSettings labelname SwitchControllerLldpSettings
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerLldpSettings extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerLldpSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerLldpSettingsState, opts?: pulumi.CustomResourceOptions): SwitchControllerLldpSettings {
        return new SwitchControllerLldpSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerLldpSettings:SwitchControllerLldpSettings';

    /**
     * Returns true if the given object is an instance of SwitchControllerLldpSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerLldpSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerLldpSettings.__pulumiType;
    }

    /**
     * Enable/disable dynamic detection of LLDP neighbor devices for VLAN assignment. Valid values: `disable`, `enable`.
     */
    public readonly deviceDetection!: pulumi.Output<string>;
    /**
     * Frequency of LLDP PDU transmission from FortiSwitch for the first 4 packets when the link is up (2 - 5 sec, default = 2, 0 = disable fast start).
     */
    public readonly fastStartInterval!: pulumi.Output<number>;
    /**
     * Primary management interface to be advertised in LLDP and CDP PDUs. Valid values: `internal`, `mgmt`.
     */
    public readonly managementInterface!: pulumi.Output<string>;
    /**
     * Enable/disable LLDP global settings. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Number of tx-intervals before local LLDP data expires (1 - 16, default = 4). Packet TTL is tx-hold * tx-interval.
     */
    public readonly txHold!: pulumi.Output<number>;
    /**
     * Frequency of LLDP PDU transmission from FortiSwitch (5 - 4095 sec, default = 30). Packet TTL is tx-hold * tx-interval.
     */
    public readonly txInterval!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerLldpSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerLldpSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerLldpSettingsArgs | SwitchControllerLldpSettingsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerLldpSettingsState | undefined;
            inputs["deviceDetection"] = state ? state.deviceDetection : undefined;
            inputs["fastStartInterval"] = state ? state.fastStartInterval : undefined;
            inputs["managementInterface"] = state ? state.managementInterface : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["txHold"] = state ? state.txHold : undefined;
            inputs["txInterval"] = state ? state.txInterval : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerLldpSettingsArgs | undefined;
            inputs["deviceDetection"] = args ? args.deviceDetection : undefined;
            inputs["fastStartInterval"] = args ? args.fastStartInterval : undefined;
            inputs["managementInterface"] = args ? args.managementInterface : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["txHold"] = args ? args.txHold : undefined;
            inputs["txInterval"] = args ? args.txInterval : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerLldpSettings.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerLldpSettings resources.
 */
export interface SwitchControllerLldpSettingsState {
    /**
     * Enable/disable dynamic detection of LLDP neighbor devices for VLAN assignment. Valid values: `disable`, `enable`.
     */
    deviceDetection?: pulumi.Input<string>;
    /**
     * Frequency of LLDP PDU transmission from FortiSwitch for the first 4 packets when the link is up (2 - 5 sec, default = 2, 0 = disable fast start).
     */
    fastStartInterval?: pulumi.Input<number>;
    /**
     * Primary management interface to be advertised in LLDP and CDP PDUs. Valid values: `internal`, `mgmt`.
     */
    managementInterface?: pulumi.Input<string>;
    /**
     * Enable/disable LLDP global settings. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Number of tx-intervals before local LLDP data expires (1 - 16, default = 4). Packet TTL is tx-hold * tx-interval.
     */
    txHold?: pulumi.Input<number>;
    /**
     * Frequency of LLDP PDU transmission from FortiSwitch (5 - 4095 sec, default = 30). Packet TTL is tx-hold * tx-interval.
     */
    txInterval?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerLldpSettings resource.
 */
export interface SwitchControllerLldpSettingsArgs {
    /**
     * Enable/disable dynamic detection of LLDP neighbor devices for VLAN assignment. Valid values: `disable`, `enable`.
     */
    deviceDetection?: pulumi.Input<string>;
    /**
     * Frequency of LLDP PDU transmission from FortiSwitch for the first 4 packets when the link is up (2 - 5 sec, default = 2, 0 = disable fast start).
     */
    fastStartInterval?: pulumi.Input<number>;
    /**
     * Primary management interface to be advertised in LLDP and CDP PDUs. Valid values: `internal`, `mgmt`.
     */
    managementInterface?: pulumi.Input<string>;
    /**
     * Enable/disable LLDP global settings. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Number of tx-intervals before local LLDP data expires (1 - 16, default = 4). Packet TTL is tx-hold * tx-interval.
     */
    txHold?: pulumi.Input<number>;
    /**
     * Frequency of LLDP PDU transmission from FortiSwitch (5 - 4095 sec, default = 30). Packet TTL is tx-hold * tx-interval.
     */
    txInterval?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
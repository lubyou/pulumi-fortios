// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch spanning tree protocol (STP).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SwitchControllerStpSettings("trname", {
 *     forwardTime: 15,
 *     helloTime: 2,
 *     maxAge: 20,
 *     maxHops: 20,
 *     pendingTimer: 4,
 *     revision: 0,
 *     status: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * SwitchController StpSettings can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerStpSettings:SwitchControllerStpSettings labelname SwitchControllerStpSettings
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerStpSettings extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerStpSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerStpSettingsState, opts?: pulumi.CustomResourceOptions): SwitchControllerStpSettings {
        return new SwitchControllerStpSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerStpSettings:SwitchControllerStpSettings';

    /**
     * Returns true if the given object is an instance of SwitchControllerStpSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerStpSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerStpSettings.__pulumiType;
    }

    /**
     * Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
     */
    public readonly forwardTime!: pulumi.Output<number>;
    /**
     * Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
     */
    public readonly helloTime!: pulumi.Output<number>;
    /**
     * Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
     */
    public readonly maxAge!: pulumi.Output<number>;
    /**
     * Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
     */
    public readonly maxHops!: pulumi.Output<number>;
    /**
     * Name of global STP settings configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Pending time (1 - 15 sec, default = 4).
     */
    public readonly pendingTimer!: pulumi.Output<number>;
    /**
     * STP revision number (0 - 65535).
     */
    public readonly revision!: pulumi.Output<number>;
    /**
     * Enable/disable STP. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerStpSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerStpSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerStpSettingsArgs | SwitchControllerStpSettingsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerStpSettingsState | undefined;
            inputs["forwardTime"] = state ? state.forwardTime : undefined;
            inputs["helloTime"] = state ? state.helloTime : undefined;
            inputs["maxAge"] = state ? state.maxAge : undefined;
            inputs["maxHops"] = state ? state.maxHops : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pendingTimer"] = state ? state.pendingTimer : undefined;
            inputs["revision"] = state ? state.revision : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerStpSettingsArgs | undefined;
            inputs["forwardTime"] = args ? args.forwardTime : undefined;
            inputs["helloTime"] = args ? args.helloTime : undefined;
            inputs["maxAge"] = args ? args.maxAge : undefined;
            inputs["maxHops"] = args ? args.maxHops : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pendingTimer"] = args ? args.pendingTimer : undefined;
            inputs["revision"] = args ? args.revision : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerStpSettings.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerStpSettings resources.
 */
export interface SwitchControllerStpSettingsState {
    /**
     * Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
     */
    forwardTime?: pulumi.Input<number>;
    /**
     * Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
     */
    helloTime?: pulumi.Input<number>;
    /**
     * Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
     */
    maxAge?: pulumi.Input<number>;
    /**
     * Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
     */
    maxHops?: pulumi.Input<number>;
    /**
     * Name of global STP settings configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * Pending time (1 - 15 sec, default = 4).
     */
    pendingTimer?: pulumi.Input<number>;
    /**
     * STP revision number (0 - 65535).
     */
    revision?: pulumi.Input<number>;
    /**
     * Enable/disable STP. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerStpSettings resource.
 */
export interface SwitchControllerStpSettingsArgs {
    /**
     * Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
     */
    forwardTime?: pulumi.Input<number>;
    /**
     * Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
     */
    helloTime?: pulumi.Input<number>;
    /**
     * Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
     */
    maxAge?: pulumi.Input<number>;
    /**
     * Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
     */
    maxHops?: pulumi.Input<number>;
    /**
     * Name of global STP settings configuration.
     */
    name?: pulumi.Input<string>;
    /**
     * Pending time (1 - 15 sec, default = 4).
     */
    pendingTimer?: pulumi.Input<number>;
    /**
     * STP revision number (0 - 65535).
     */
    revision?: pulumi.Input<number>;
    /**
     * Enable/disable STP. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure HA monitor.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemHaMonitor("trname", {
 *     monitorVlan: "disable",
 *     vlanHbInterval: 5,
 *     vlanHbLostThreshold: 3,
 * });
 * ```
 *
 * ## Import
 *
 * System HaMonitor can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemHaMonitor:SystemHaMonitor labelname SystemHaMonitor
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemHaMonitor extends pulumi.CustomResource {
    /**
     * Get an existing SystemHaMonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemHaMonitorState, opts?: pulumi.CustomResourceOptions): SystemHaMonitor {
        return new SystemHaMonitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemHaMonitor:SystemHaMonitor';

    /**
     * Returns true if the given object is an instance of SystemHaMonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemHaMonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemHaMonitor.__pulumiType;
    }

    /**
     * Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
     */
    public readonly monitorVlan!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Configure heartbeat interval (seconds).
     */
    public readonly vlanHbInterval!: pulumi.Output<number>;
    /**
     * VLAN lost heartbeat threshold (1 - 60).
     */
    public readonly vlanHbLostThreshold!: pulumi.Output<number>;

    /**
     * Create a SystemHaMonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemHaMonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemHaMonitorArgs | SystemHaMonitorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemHaMonitorState | undefined;
            inputs["monitorVlan"] = state ? state.monitorVlan : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["vlanHbInterval"] = state ? state.vlanHbInterval : undefined;
            inputs["vlanHbLostThreshold"] = state ? state.vlanHbLostThreshold : undefined;
        } else {
            const args = argsOrState as SystemHaMonitorArgs | undefined;
            inputs["monitorVlan"] = args ? args.monitorVlan : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["vlanHbInterval"] = args ? args.vlanHbInterval : undefined;
            inputs["vlanHbLostThreshold"] = args ? args.vlanHbLostThreshold : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemHaMonitor.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemHaMonitor resources.
 */
export interface SystemHaMonitorState {
    /**
     * Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
     */
    monitorVlan?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Configure heartbeat interval (seconds).
     */
    vlanHbInterval?: pulumi.Input<number>;
    /**
     * VLAN lost heartbeat threshold (1 - 60).
     */
    vlanHbLostThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SystemHaMonitor resource.
 */
export interface SystemHaMonitorArgs {
    /**
     * Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
     */
    monitorVlan?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Configure heartbeat interval (seconds).
     */
    vlanHbInterval?: pulumi.Input<number>;
    /**
     * VLAN lost heartbeat threshold (1 - 60).
     */
    vlanHbLostThreshold?: pulumi.Input<number>;
}

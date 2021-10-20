// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure Type of Service (ToS) based priority table to set network traffic priorities.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemTosBasedPriority("trname", {
 *     fosid: 1,
 *     priority: "low",
 *     tos: 11,
 * });
 * ```
 *
 * ## Import
 *
 * System TosBasedPriority can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemTosBasedPriority:SystemTosBasedPriority labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemTosBasedPriority extends pulumi.CustomResource {
    /**
     * Get an existing SystemTosBasedPriority resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemTosBasedPriorityState, opts?: pulumi.CustomResourceOptions): SystemTosBasedPriority {
        return new SystemTosBasedPriority(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemTosBasedPriority:SystemTosBasedPriority';

    /**
     * Returns true if the given object is an instance of SystemTosBasedPriority.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemTosBasedPriority {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemTosBasedPriority.__pulumiType;
    }

    /**
     * Item ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium). Valid values: `low`, `medium`, `high`.
     */
    public readonly priority!: pulumi.Output<string>;
    /**
     * Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
     */
    public readonly tos!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemTosBasedPriority resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemTosBasedPriorityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemTosBasedPriorityArgs | SystemTosBasedPriorityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemTosBasedPriorityState | undefined;
            inputs["fosid"] = state ? state.fosid : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["tos"] = state ? state.tos : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemTosBasedPriorityArgs | undefined;
            inputs["fosid"] = args ? args.fosid : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["tos"] = args ? args.tos : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemTosBasedPriority.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemTosBasedPriority resources.
 */
export interface SystemTosBasedPriorityState {
    /**
     * Item ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium). Valid values: `low`, `medium`, `high`.
     */
    priority?: pulumi.Input<string>;
    /**
     * Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
     */
    tos?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemTosBasedPriority resource.
 */
export interface SystemTosBasedPriorityArgs {
    /**
     * Item ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium). Valid values: `low`, `medium`, `high`.
     */
    priority?: pulumi.Input<string>;
    /**
     * Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
     */
    tos?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
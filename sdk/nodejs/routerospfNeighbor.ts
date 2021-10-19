// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * OSPF neighbor configuration are used when OSPF runs on non-broadcast media
 *
 * > The provider supports the definition of Neighbor in Router Ospf `fortios.RouterOspf`, and also allows the definition of separate Neighbor resources `fortios.RouterospfNeighbor`, but do not use a `fortios.RouterOspf` with in-line Neighbor in conjunction with any `fortios.RouterospfNeighbor` resources, otherwise conflicts and overwrite will occur.
 *
 * ## Import
 *
 * Routerospf Neighbor can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/routerospfNeighbor:RouterospfNeighbor labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class RouterospfNeighbor extends pulumi.CustomResource {
    /**
     * Get an existing RouterospfNeighbor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterospfNeighborState, opts?: pulumi.CustomResourceOptions): RouterospfNeighbor {
        return new RouterospfNeighbor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerospfNeighbor:RouterospfNeighbor';

    /**
     * Returns true if the given object is an instance of RouterospfNeighbor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterospfNeighbor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterospfNeighbor.__pulumiType;
    }

    /**
     * Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
     */
    public readonly cost!: pulumi.Output<number>;
    /**
     * Neighbor entry ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Interface IP address of the neighbor.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Poll interval time in seconds.
     */
    public readonly pollInterval!: pulumi.Output<number>;
    /**
     * Priority.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterospfNeighbor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouterospfNeighborArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterospfNeighborArgs | RouterospfNeighborState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterospfNeighborState | undefined;
            inputs["cost"] = state ? state.cost : undefined;
            inputs["fosid"] = state ? state.fosid : undefined;
            inputs["ip"] = state ? state.ip : undefined;
            inputs["pollInterval"] = state ? state.pollInterval : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RouterospfNeighborArgs | undefined;
            inputs["cost"] = args ? args.cost : undefined;
            inputs["fosid"] = args ? args.fosid : undefined;
            inputs["ip"] = args ? args.ip : undefined;
            inputs["pollInterval"] = args ? args.pollInterval : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RouterospfNeighbor.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterospfNeighbor resources.
 */
export interface RouterospfNeighborState {
    /**
     * Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
     */
    cost?: pulumi.Input<number>;
    /**
     * Neighbor entry ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Interface IP address of the neighbor.
     */
    ip?: pulumi.Input<string>;
    /**
     * Poll interval time in seconds.
     */
    pollInterval?: pulumi.Input<number>;
    /**
     * Priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterospfNeighbor resource.
 */
export interface RouterospfNeighborArgs {
    /**
     * Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
     */
    cost?: pulumi.Input<number>;
    /**
     * Neighbor entry ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Interface IP address of the neighbor.
     */
    ip?: pulumi.Input<string>;
    /**
     * Poll interval time in seconds.
     */
    pollInterval?: pulumi.Input<number>;
    /**
     * Priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

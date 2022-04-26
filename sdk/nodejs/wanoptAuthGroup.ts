// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure WAN optimization authentication groups.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WanoptAuthGroup("trname", {
 *     authMethod: "cert",
 *     cert: "Fortinet_CA_SSL",
 *     peerAccept: "any",
 * });
 * ```
 *
 * ## Import
 *
 * Wanopt AuthGroup can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wanoptAuthGroup:WanoptAuthGroup labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wanoptAuthGroup:WanoptAuthGroup labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WanoptAuthGroup extends pulumi.CustomResource {
    /**
     * Get an existing WanoptAuthGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WanoptAuthGroupState, opts?: pulumi.CustomResourceOptions): WanoptAuthGroup {
        return new WanoptAuthGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wanoptAuthGroup:WanoptAuthGroup';

    /**
     * Returns true if the given object is an instance of WanoptAuthGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WanoptAuthGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WanoptAuthGroup.__pulumiType;
    }

    /**
     * Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
     */
    public readonly authMethod!: pulumi.Output<string>;
    /**
     * Name of certificate to identify this peer.
     */
    public readonly cert!: pulumi.Output<string>;
    /**
     * Auth-group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
     */
    public readonly peer!: pulumi.Output<string>;
    /**
     * Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
     */
    public readonly peerAccept!: pulumi.Output<string>;
    /**
     * Pre-shared key used by the peers in this authentication group.
     */
    public readonly psk!: pulumi.Output<string | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WanoptAuthGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WanoptAuthGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WanoptAuthGroupArgs | WanoptAuthGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WanoptAuthGroupState | undefined;
            resourceInputs["authMethod"] = state ? state.authMethod : undefined;
            resourceInputs["cert"] = state ? state.cert : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["peer"] = state ? state.peer : undefined;
            resourceInputs["peerAccept"] = state ? state.peerAccept : undefined;
            resourceInputs["psk"] = state ? state.psk : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WanoptAuthGroupArgs | undefined;
            if ((!args || args.cert === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cert'");
            }
            resourceInputs["authMethod"] = args ? args.authMethod : undefined;
            resourceInputs["cert"] = args ? args.cert : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["peer"] = args ? args.peer : undefined;
            resourceInputs["peerAccept"] = args ? args.peerAccept : undefined;
            resourceInputs["psk"] = args ? args.psk : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WanoptAuthGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WanoptAuthGroup resources.
 */
export interface WanoptAuthGroupState {
    /**
     * Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
     */
    authMethod?: pulumi.Input<string>;
    /**
     * Name of certificate to identify this peer.
     */
    cert?: pulumi.Input<string>;
    /**
     * Auth-group name.
     */
    name?: pulumi.Input<string>;
    /**
     * If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
     */
    peer?: pulumi.Input<string>;
    /**
     * Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
     */
    peerAccept?: pulumi.Input<string>;
    /**
     * Pre-shared key used by the peers in this authentication group.
     */
    psk?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WanoptAuthGroup resource.
 */
export interface WanoptAuthGroupArgs {
    /**
     * Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
     */
    authMethod?: pulumi.Input<string>;
    /**
     * Name of certificate to identify this peer.
     */
    cert: pulumi.Input<string>;
    /**
     * Auth-group name.
     */
    name?: pulumi.Input<string>;
    /**
     * If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
     */
    peer?: pulumi.Input<string>;
    /**
     * Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
     */
    peerAccept?: pulumi.Input<string>;
    /**
     * Pre-shared key used by the peers in this authentication group.
     */
    psk?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

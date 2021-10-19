// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure IPv6 routing policies.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.RouterPolicy6("trname", {
 *     dst: "::/0",
 *     endPort: 65535,
 *     gateway: "::",
 *     inputDevice: "port1",
 *     outputDevice: "port3",
 *     protocol: 33,
 *     seqNum: 1,
 *     src: "2001:db8:85a3::8a2e:370:7334/128",
 *     startPort: 1,
 *     status: "enable",
 *     tos: "0x00",
 *     tosMask: "0x00",
 * });
 * ```
 *
 * ## Import
 *
 * Router Policy6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/routerPolicy6:RouterPolicy6 labelname {{seq_num}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class RouterPolicy6 extends pulumi.CustomResource {
    /**
     * Get an existing RouterPolicy6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterPolicy6State, opts?: pulumi.CustomResourceOptions): RouterPolicy6 {
        return new RouterPolicy6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerPolicy6:RouterPolicy6';

    /**
     * Returns true if the given object is an instance of RouterPolicy6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterPolicy6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterPolicy6.__pulumiType;
    }

    /**
     * Optional comments.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Destination IPv6 prefix.
     */
    public readonly dst!: pulumi.Output<string>;
    /**
     * End destination port number (1 - 65535).
     */
    public readonly endPort!: pulumi.Output<number>;
    /**
     * IPv6 address of the gateway.
     */
    public readonly gateway!: pulumi.Output<string>;
    /**
     * Incoming interface name. Configuration examples: for FortiOS Version <= "6.2.4": `inputDevice  = "port2"`, for FortiOS Version >= "6.2.4": `inputDevice  = "\"fortilink\" \"port1\""`.
     */
    public readonly inputDevice!: pulumi.Output<string>;
    /**
     * Outgoing interface name.
     */
    public readonly outputDevice!: pulumi.Output<string>;
    /**
     * Protocol number (0 - 255).
     */
    public readonly protocol!: pulumi.Output<number>;
    /**
     * Sequence number.
     */
    public readonly seqNum!: pulumi.Output<number>;
    /**
     * Source IPv6 prefix.
     */
    public readonly src!: pulumi.Output<string>;
    /**
     * Start destination port number (1 - 65535).
     */
    public readonly startPort!: pulumi.Output<number>;
    /**
     * Enable/disable this policy route. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Type of service bit pattern.
     */
    public readonly tos!: pulumi.Output<string>;
    /**
     * Type of service evaluated bits.
     */
    public readonly tosMask!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterPolicy6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterPolicy6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterPolicy6Args | RouterPolicy6State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterPolicy6State | undefined;
            inputs["comments"] = state ? state.comments : undefined;
            inputs["dst"] = state ? state.dst : undefined;
            inputs["endPort"] = state ? state.endPort : undefined;
            inputs["gateway"] = state ? state.gateway : undefined;
            inputs["inputDevice"] = state ? state.inputDevice : undefined;
            inputs["outputDevice"] = state ? state.outputDevice : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["seqNum"] = state ? state.seqNum : undefined;
            inputs["src"] = state ? state.src : undefined;
            inputs["startPort"] = state ? state.startPort : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tos"] = state ? state.tos : undefined;
            inputs["tosMask"] = state ? state.tosMask : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RouterPolicy6Args | undefined;
            if ((!args || args.inputDevice === undefined) && !opts.urn) {
                throw new Error("Missing required property 'inputDevice'");
            }
            inputs["comments"] = args ? args.comments : undefined;
            inputs["dst"] = args ? args.dst : undefined;
            inputs["endPort"] = args ? args.endPort : undefined;
            inputs["gateway"] = args ? args.gateway : undefined;
            inputs["inputDevice"] = args ? args.inputDevice : undefined;
            inputs["outputDevice"] = args ? args.outputDevice : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["seqNum"] = args ? args.seqNum : undefined;
            inputs["src"] = args ? args.src : undefined;
            inputs["startPort"] = args ? args.startPort : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tos"] = args ? args.tos : undefined;
            inputs["tosMask"] = args ? args.tosMask : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RouterPolicy6.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterPolicy6 resources.
 */
export interface RouterPolicy6State {
    /**
     * Optional comments.
     */
    comments?: pulumi.Input<string>;
    /**
     * Destination IPv6 prefix.
     */
    dst?: pulumi.Input<string>;
    /**
     * End destination port number (1 - 65535).
     */
    endPort?: pulumi.Input<number>;
    /**
     * IPv6 address of the gateway.
     */
    gateway?: pulumi.Input<string>;
    /**
     * Incoming interface name. Configuration examples: for FortiOS Version <= "6.2.4": `inputDevice  = "port2"`, for FortiOS Version >= "6.2.4": `inputDevice  = "\"fortilink\" \"port1\""`.
     */
    inputDevice?: pulumi.Input<string>;
    /**
     * Outgoing interface name.
     */
    outputDevice?: pulumi.Input<string>;
    /**
     * Protocol number (0 - 255).
     */
    protocol?: pulumi.Input<number>;
    /**
     * Sequence number.
     */
    seqNum?: pulumi.Input<number>;
    /**
     * Source IPv6 prefix.
     */
    src?: pulumi.Input<string>;
    /**
     * Start destination port number (1 - 65535).
     */
    startPort?: pulumi.Input<number>;
    /**
     * Enable/disable this policy route. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Type of service bit pattern.
     */
    tos?: pulumi.Input<string>;
    /**
     * Type of service evaluated bits.
     */
    tosMask?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterPolicy6 resource.
 */
export interface RouterPolicy6Args {
    /**
     * Optional comments.
     */
    comments?: pulumi.Input<string>;
    /**
     * Destination IPv6 prefix.
     */
    dst?: pulumi.Input<string>;
    /**
     * End destination port number (1 - 65535).
     */
    endPort?: pulumi.Input<number>;
    /**
     * IPv6 address of the gateway.
     */
    gateway?: pulumi.Input<string>;
    /**
     * Incoming interface name. Configuration examples: for FortiOS Version <= "6.2.4": `inputDevice  = "port2"`, for FortiOS Version >= "6.2.4": `inputDevice  = "\"fortilink\" \"port1\""`.
     */
    inputDevice: pulumi.Input<string>;
    /**
     * Outgoing interface name.
     */
    outputDevice?: pulumi.Input<string>;
    /**
     * Protocol number (0 - 255).
     */
    protocol?: pulumi.Input<number>;
    /**
     * Sequence number.
     */
    seqNum?: pulumi.Input<number>;
    /**
     * Source IPv6 prefix.
     */
    src?: pulumi.Input<string>;
    /**
     * Start destination port number (1 - 65535).
     */
    startPort?: pulumi.Input<number>;
    /**
     * Enable/disable this policy route. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Type of service bit pattern.
     */
    tos?: pulumi.Input<string>;
    /**
     * Type of service evaluated bits.
     */
    tosMask?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

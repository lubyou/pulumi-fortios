// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure IPv6 tunnel over IPv4.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemSitTunnel("trname", {
 *     destination: "1.1.1.1",
 *     interface: "port2",
 *     ip6: "::/0",
 *     source: "2.2.2.2",
 * });
 * ```
 *
 * ## Import
 *
 * System SitTunnel can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemSitTunnel:SystemSitTunnel labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemSitTunnel:SystemSitTunnel labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemSitTunnel extends pulumi.CustomResource {
    /**
     * Get an existing SystemSitTunnel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemSitTunnelState, opts?: pulumi.CustomResourceOptions): SystemSitTunnel {
        return new SystemSitTunnel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemSitTunnel:SystemSitTunnel';

    /**
     * Returns true if the given object is an instance of SystemSitTunnel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemSitTunnel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemSitTunnel.__pulumiType;
    }

    /**
     * Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
     */
    public readonly autoAsicOffload!: pulumi.Output<string>;
    /**
     * Destination IP address of the tunnel.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * Interface name.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * IPv6 address of the tunnel.
     */
    public readonly ip6!: pulumi.Output<string>;
    /**
     * Tunnel name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Source IP address of the tunnel.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
     */
    public readonly useSdwan!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemSitTunnel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemSitTunnelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemSitTunnelArgs | SystemSitTunnelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemSitTunnelState | undefined;
            resourceInputs["autoAsicOffload"] = state ? state.autoAsicOffload : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ip6"] = state ? state.ip6 : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["useSdwan"] = state ? state.useSdwan : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemSitTunnelArgs | undefined;
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            resourceInputs["autoAsicOffload"] = args ? args.autoAsicOffload : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ip6"] = args ? args.ip6 : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["useSdwan"] = args ? args.useSdwan : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemSitTunnel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemSitTunnel resources.
 */
export interface SystemSitTunnelState {
    /**
     * Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
     */
    autoAsicOffload?: pulumi.Input<string>;
    /**
     * Destination IP address of the tunnel.
     */
    destination?: pulumi.Input<string>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * IPv6 address of the tunnel.
     */
    ip6?: pulumi.Input<string>;
    /**
     * Tunnel name.
     */
    name?: pulumi.Input<string>;
    /**
     * Source IP address of the tunnel.
     */
    source?: pulumi.Input<string>;
    /**
     * Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
     */
    useSdwan?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemSitTunnel resource.
 */
export interface SystemSitTunnelArgs {
    /**
     * Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
     */
    autoAsicOffload?: pulumi.Input<string>;
    /**
     * Destination IP address of the tunnel.
     */
    destination: pulumi.Input<string>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * IPv6 address of the tunnel.
     */
    ip6?: pulumi.Input<string>;
    /**
     * Tunnel name.
     */
    name?: pulumi.Input<string>;
    /**
     * Source IP address of the tunnel.
     */
    source?: pulumi.Input<string>;
    /**
     * Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
     */
    useSdwan?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure IP in IP Tunneling.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemIpipTunnel("trname", {
 *     interface: "port3",
 *     localGw: "1.1.1.1",
 *     remoteGw: "2.2.2.2",
 * });
 * ```
 *
 * ## Import
 *
 * System IpipTunnel can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemIpipTunnel:SystemIpipTunnel labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemIpipTunnel:SystemIpipTunnel labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemIpipTunnel extends pulumi.CustomResource {
    /**
     * Get an existing SystemIpipTunnel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemIpipTunnelState, opts?: pulumi.CustomResourceOptions): SystemIpipTunnel {
        return new SystemIpipTunnel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemIpipTunnel:SystemIpipTunnel';

    /**
     * Returns true if the given object is an instance of SystemIpipTunnel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemIpipTunnel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemIpipTunnel.__pulumiType;
    }

    /**
     * Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
     */
    public readonly autoAsicOffload!: pulumi.Output<string>;
    /**
     * Interface name that is associated with the incoming traffic from available options.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * IPv4 address for the local gateway.
     */
    public readonly localGw!: pulumi.Output<string>;
    /**
     * IPIP Tunnel name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * IPv4 address for the remote gateway.
     */
    public readonly remoteGw!: pulumi.Output<string>;
    /**
     * Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
     */
    public readonly useSdwan!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemIpipTunnel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemIpipTunnelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemIpipTunnelArgs | SystemIpipTunnelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemIpipTunnelState | undefined;
            resourceInputs["autoAsicOffload"] = state ? state.autoAsicOffload : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["localGw"] = state ? state.localGw : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["remoteGw"] = state ? state.remoteGw : undefined;
            resourceInputs["useSdwan"] = state ? state.useSdwan : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemIpipTunnelArgs | undefined;
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.localGw === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localGw'");
            }
            if ((!args || args.remoteGw === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteGw'");
            }
            resourceInputs["autoAsicOffload"] = args ? args.autoAsicOffload : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["localGw"] = args ? args.localGw : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["remoteGw"] = args ? args.remoteGw : undefined;
            resourceInputs["useSdwan"] = args ? args.useSdwan : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemIpipTunnel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemIpipTunnel resources.
 */
export interface SystemIpipTunnelState {
    /**
     * Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
     */
    autoAsicOffload?: pulumi.Input<string>;
    /**
     * Interface name that is associated with the incoming traffic from available options.
     */
    interface?: pulumi.Input<string>;
    /**
     * IPv4 address for the local gateway.
     */
    localGw?: pulumi.Input<string>;
    /**
     * IPIP Tunnel name.
     */
    name?: pulumi.Input<string>;
    /**
     * IPv4 address for the remote gateway.
     */
    remoteGw?: pulumi.Input<string>;
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
 * The set of arguments for constructing a SystemIpipTunnel resource.
 */
export interface SystemIpipTunnelArgs {
    /**
     * Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
     */
    autoAsicOffload?: pulumi.Input<string>;
    /**
     * Interface name that is associated with the incoming traffic from available options.
     */
    interface: pulumi.Input<string>;
    /**
     * IPv4 address for the local gateway.
     */
    localGw: pulumi.Input<string>;
    /**
     * IPIP Tunnel name.
     */
    name?: pulumi.Input<string>;
    /**
     * IPv4 address for the remote gateway.
     */
    remoteGw: pulumi.Input<string>;
    /**
     * Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
     */
    useSdwan?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

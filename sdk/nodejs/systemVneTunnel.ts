// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure virtual network enabler tunnel. Applies to FortiOS Version `>= 6.4.2`.
 *
 * ## Import
 *
 * System VneTunnel can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemVneTunnel:SystemVneTunnel labelname SystemVneTunnel
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemVneTunnel extends pulumi.CustomResource {
    /**
     * Get an existing SystemVneTunnel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemVneTunnelState, opts?: pulumi.CustomResourceOptions): SystemVneTunnel {
        return new SystemVneTunnel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemVneTunnel:SystemVneTunnel';

    /**
     * Returns true if the given object is an instance of SystemVneTunnel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemVneTunnel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemVneTunnel.__pulumiType;
    }

    /**
     * Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
     */
    public readonly autoAsicOffload!: pulumi.Output<string>;
    /**
     * BMR hostname.
     */
    public readonly bmrHostname!: pulumi.Output<string | undefined>;
    /**
     * Border relay IPv6 address.
     */
    public readonly br!: pulumi.Output<string>;
    /**
     * Interface name.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Tunnel IPv4 address and netmask.
     */
    public readonly ipv4Address!: pulumi.Output<string>;
    /**
     * VNE tunnel mode. Valid values: `map-e`, `fixed-ip`.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Name of local certificate for SSL connections.
     */
    public readonly sslCertificate!: pulumi.Output<string>;
    /**
     * Enable/disable VNE tunnel. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * URL of provisioning server.
     */
    public readonly updateUrl!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemVneTunnel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemVneTunnelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemVneTunnelArgs | SystemVneTunnelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemVneTunnelState | undefined;
            resourceInputs["autoAsicOffload"] = state ? state.autoAsicOffload : undefined;
            resourceInputs["bmrHostname"] = state ? state.bmrHostname : undefined;
            resourceInputs["br"] = state ? state.br : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ipv4Address"] = state ? state.ipv4Address : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["sslCertificate"] = state ? state.sslCertificate : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updateUrl"] = state ? state.updateUrl : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemVneTunnelArgs | undefined;
            resourceInputs["autoAsicOffload"] = args ? args.autoAsicOffload : undefined;
            resourceInputs["bmrHostname"] = args ? args.bmrHostname : undefined;
            resourceInputs["br"] = args ? args.br : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ipv4Address"] = args ? args.ipv4Address : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["sslCertificate"] = args ? args.sslCertificate : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["updateUrl"] = args ? args.updateUrl : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemVneTunnel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemVneTunnel resources.
 */
export interface SystemVneTunnelState {
    /**
     * Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
     */
    autoAsicOffload?: pulumi.Input<string>;
    /**
     * BMR hostname.
     */
    bmrHostname?: pulumi.Input<string>;
    /**
     * Border relay IPv6 address.
     */
    br?: pulumi.Input<string>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * Tunnel IPv4 address and netmask.
     */
    ipv4Address?: pulumi.Input<string>;
    /**
     * VNE tunnel mode. Valid values: `map-e`, `fixed-ip`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Name of local certificate for SSL connections.
     */
    sslCertificate?: pulumi.Input<string>;
    /**
     * Enable/disable VNE tunnel. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * URL of provisioning server.
     */
    updateUrl?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemVneTunnel resource.
 */
export interface SystemVneTunnelArgs {
    /**
     * Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
     */
    autoAsicOffload?: pulumi.Input<string>;
    /**
     * BMR hostname.
     */
    bmrHostname?: pulumi.Input<string>;
    /**
     * Border relay IPv6 address.
     */
    br?: pulumi.Input<string>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * Tunnel IPv4 address and netmask.
     */
    ipv4Address?: pulumi.Input<string>;
    /**
     * VNE tunnel mode. Valid values: `map-e`, `fixed-ip`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Name of local certificate for SSL connections.
     */
    sslCertificate?: pulumi.Input<string>;
    /**
     * Enable/disable VNE tunnel. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * URL of provisioning server.
     */
    updateUrl?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

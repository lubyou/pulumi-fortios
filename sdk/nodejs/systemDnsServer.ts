// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure DNS servers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemDnsServer("trname", {
 *     dnsfilterProfile: "default",
 *     mode: "forward-only",
 * });
 * ```
 *
 * ## Import
 *
 * System DnsServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemDnsServer:SystemDnsServer labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemDnsServer extends pulumi.CustomResource {
    /**
     * Get an existing SystemDnsServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemDnsServerState, opts?: pulumi.CustomResourceOptions): SystemDnsServer {
        return new SystemDnsServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemDnsServer:SystemDnsServer';

    /**
     * Returns true if the given object is an instance of SystemDnsServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemDnsServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemDnsServer.__pulumiType;
    }

    /**
     * DNS filter profile.
     */
    public readonly dnsfilterProfile!: pulumi.Output<string>;
    /**
     * DNS server mode. Valid values: `recursive`, `non-recursive`, `forward-only`.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * DNS server name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemDnsServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemDnsServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemDnsServerArgs | SystemDnsServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemDnsServerState | undefined;
            inputs["dnsfilterProfile"] = state ? state.dnsfilterProfile : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemDnsServerArgs | undefined;
            inputs["dnsfilterProfile"] = args ? args.dnsfilterProfile : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemDnsServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemDnsServer resources.
 */
export interface SystemDnsServerState {
    /**
     * DNS filter profile.
     */
    dnsfilterProfile?: pulumi.Input<string>;
    /**
     * DNS server mode. Valid values: `recursive`, `non-recursive`, `forward-only`.
     */
    mode?: pulumi.Input<string>;
    /**
     * DNS server name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemDnsServer resource.
 */
export interface SystemDnsServerArgs {
    /**
     * DNS filter profile.
     */
    dnsfilterProfile?: pulumi.Input<string>;
    /**
     * DNS server mode. Valid values: `recursive`, `non-recursive`, `forward-only`.
     */
    mode?: pulumi.Input<string>;
    /**
     * DNS server name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
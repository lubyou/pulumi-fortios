// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure DNS translation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallDnstranslation("trname", {
 *     dst: "2.2.2.2",
 *     fosid: 1,
 *     netmask: "255.0.0.0",
 *     src: "1.1.1.1",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Dnstranslation can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallDnstranslation:FirewallDnstranslation labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallDnstranslation extends pulumi.CustomResource {
    /**
     * Get an existing FirewallDnstranslation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallDnstranslationState, opts?: pulumi.CustomResourceOptions): FirewallDnstranslation {
        return new FirewallDnstranslation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallDnstranslation:FirewallDnstranslation';

    /**
     * Returns true if the given object is an instance of FirewallDnstranslation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallDnstranslation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallDnstranslation.__pulumiType;
    }

    /**
     * IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
     */
    public readonly dst!: pulumi.Output<string>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
     */
    public readonly netmask!: pulumi.Output<string>;
    /**
     * IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
     */
    public readonly src!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallDnstranslation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallDnstranslationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallDnstranslationArgs | FirewallDnstranslationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallDnstranslationState | undefined;
            resourceInputs["dst"] = state ? state.dst : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["netmask"] = state ? state.netmask : undefined;
            resourceInputs["src"] = state ? state.src : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallDnstranslationArgs | undefined;
            resourceInputs["dst"] = args ? args.dst : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["netmask"] = args ? args.netmask : undefined;
            resourceInputs["src"] = args ? args.src : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallDnstranslation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallDnstranslation resources.
 */
export interface FirewallDnstranslationState {
    /**
     * IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
     */
    dst?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
     */
    netmask?: pulumi.Input<string>;
    /**
     * IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
     */
    src?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallDnstranslation resource.
 */
export interface FirewallDnstranslationArgs {
    /**
     * IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
     */
    dst?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
     */
    netmask?: pulumi.Input<string>;
    /**
     * IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
     */
    src?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

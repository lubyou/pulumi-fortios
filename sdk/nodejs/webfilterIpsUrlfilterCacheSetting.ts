// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure IPS URL filter cache settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WebfilterIpsUrlfilterCacheSetting("trname", {
 *     dnsRetryInterval: 0,
 *     extendedTtl: 0,
 * });
 * ```
 *
 * ## Import
 *
 * Webfilter IpsUrlfilterCacheSetting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/webfilterIpsUrlfilterCacheSetting:WebfilterIpsUrlfilterCacheSetting labelname WebfilterIpsUrlfilterCacheSetting
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WebfilterIpsUrlfilterCacheSetting extends pulumi.CustomResource {
    /**
     * Get an existing WebfilterIpsUrlfilterCacheSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebfilterIpsUrlfilterCacheSettingState, opts?: pulumi.CustomResourceOptions): WebfilterIpsUrlfilterCacheSetting {
        return new WebfilterIpsUrlfilterCacheSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webfilterIpsUrlfilterCacheSetting:WebfilterIpsUrlfilterCacheSetting';

    /**
     * Returns true if the given object is an instance of WebfilterIpsUrlfilterCacheSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebfilterIpsUrlfilterCacheSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebfilterIpsUrlfilterCacheSetting.__pulumiType;
    }

    /**
     * Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.
     */
    public readonly dnsRetryInterval!: pulumi.Output<number>;
    /**
     * Extend time to live beyond reported by DNS. 0 means use DNS server's TTL
     */
    public readonly extendedTtl!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WebfilterIpsUrlfilterCacheSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WebfilterIpsUrlfilterCacheSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebfilterIpsUrlfilterCacheSettingArgs | WebfilterIpsUrlfilterCacheSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebfilterIpsUrlfilterCacheSettingState | undefined;
            resourceInputs["dnsRetryInterval"] = state ? state.dnsRetryInterval : undefined;
            resourceInputs["extendedTtl"] = state ? state.extendedTtl : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebfilterIpsUrlfilterCacheSettingArgs | undefined;
            resourceInputs["dnsRetryInterval"] = args ? args.dnsRetryInterval : undefined;
            resourceInputs["extendedTtl"] = args ? args.extendedTtl : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebfilterIpsUrlfilterCacheSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebfilterIpsUrlfilterCacheSetting resources.
 */
export interface WebfilterIpsUrlfilterCacheSettingState {
    /**
     * Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.
     */
    dnsRetryInterval?: pulumi.Input<number>;
    /**
     * Extend time to live beyond reported by DNS. 0 means use DNS server's TTL
     */
    extendedTtl?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebfilterIpsUrlfilterCacheSetting resource.
 */
export interface WebfilterIpsUrlfilterCacheSettingArgs {
    /**
     * Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.
     */
    dnsRetryInterval?: pulumi.Input<number>;
    /**
     * Extend time to live beyond reported by DNS. 0 means use DNS server's TTL
     */
    extendedTtl?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

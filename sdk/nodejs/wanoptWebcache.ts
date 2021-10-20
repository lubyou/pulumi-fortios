// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure global Web cache settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WanoptWebcache("trname", {
 *     alwaysRevalidate: "disable",
 *     cacheByDefault: "disable",
 *     cacheCookie: "disable",
 *     cacheExpired: "disable",
 *     defaultTtl: 1440,
 *     external: "disable",
 *     freshFactor: 100,
 *     hostValidate: "disable",
 *     ignoreConditional: "disable",
 *     ignoreIeReload: "enable",
 *     ignoreIms: "disable",
 *     ignorePnc: "disable",
 *     maxObjectSize: 512000,
 *     maxTtl: 7200,
 *     minTtl: 5,
 *     negRespTime: 0,
 *     revalPnc: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * Wanopt Webcache can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/wanoptWebcache:WanoptWebcache labelname WanoptWebcache
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WanoptWebcache extends pulumi.CustomResource {
    /**
     * Get an existing WanoptWebcache resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WanoptWebcacheState, opts?: pulumi.CustomResourceOptions): WanoptWebcache {
        return new WanoptWebcache(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wanoptWebcache:WanoptWebcache';

    /**
     * Returns true if the given object is an instance of WanoptWebcache.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WanoptWebcache {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WanoptWebcache.__pulumiType;
    }

    /**
     * Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
     */
    public readonly alwaysRevalidate!: pulumi.Output<string>;
    /**
     * Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
     */
    public readonly cacheByDefault!: pulumi.Output<string>;
    /**
     * Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
     */
    public readonly cacheCookie!: pulumi.Output<string>;
    /**
     * Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
     */
    public readonly cacheExpired!: pulumi.Output<string>;
    /**
     * Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
     */
    public readonly defaultTtl!: pulumi.Output<number>;
    /**
     * Enable/disable external Web caching. Valid values: `enable`, `disable`.
     */
    public readonly external!: pulumi.Output<string>;
    /**
     * Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
     */
    public readonly freshFactor!: pulumi.Output<number>;
    /**
     * Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
     */
    public readonly hostValidate!: pulumi.Output<string>;
    /**
     * Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
     */
    public readonly ignoreConditional!: pulumi.Output<string>;
    /**
     * Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
     */
    public readonly ignoreIeReload!: pulumi.Output<string>;
    /**
     * Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
     */
    public readonly ignoreIms!: pulumi.Output<string>;
    /**
     * Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
     */
    public readonly ignorePnc!: pulumi.Output<string>;
    /**
     * Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
     */
    public readonly maxObjectSize!: pulumi.Output<number>;
    /**
     * Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
     */
    public readonly maxTtl!: pulumi.Output<number>;
    /**
     * Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
     */
    public readonly minTtl!: pulumi.Output<number>;
    /**
     * Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
     */
    public readonly negRespTime!: pulumi.Output<number>;
    /**
     * Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
     */
    public readonly revalPnc!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WanoptWebcache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WanoptWebcacheArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WanoptWebcacheArgs | WanoptWebcacheState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WanoptWebcacheState | undefined;
            inputs["alwaysRevalidate"] = state ? state.alwaysRevalidate : undefined;
            inputs["cacheByDefault"] = state ? state.cacheByDefault : undefined;
            inputs["cacheCookie"] = state ? state.cacheCookie : undefined;
            inputs["cacheExpired"] = state ? state.cacheExpired : undefined;
            inputs["defaultTtl"] = state ? state.defaultTtl : undefined;
            inputs["external"] = state ? state.external : undefined;
            inputs["freshFactor"] = state ? state.freshFactor : undefined;
            inputs["hostValidate"] = state ? state.hostValidate : undefined;
            inputs["ignoreConditional"] = state ? state.ignoreConditional : undefined;
            inputs["ignoreIeReload"] = state ? state.ignoreIeReload : undefined;
            inputs["ignoreIms"] = state ? state.ignoreIms : undefined;
            inputs["ignorePnc"] = state ? state.ignorePnc : undefined;
            inputs["maxObjectSize"] = state ? state.maxObjectSize : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["minTtl"] = state ? state.minTtl : undefined;
            inputs["negRespTime"] = state ? state.negRespTime : undefined;
            inputs["revalPnc"] = state ? state.revalPnc : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WanoptWebcacheArgs | undefined;
            inputs["alwaysRevalidate"] = args ? args.alwaysRevalidate : undefined;
            inputs["cacheByDefault"] = args ? args.cacheByDefault : undefined;
            inputs["cacheCookie"] = args ? args.cacheCookie : undefined;
            inputs["cacheExpired"] = args ? args.cacheExpired : undefined;
            inputs["defaultTtl"] = args ? args.defaultTtl : undefined;
            inputs["external"] = args ? args.external : undefined;
            inputs["freshFactor"] = args ? args.freshFactor : undefined;
            inputs["hostValidate"] = args ? args.hostValidate : undefined;
            inputs["ignoreConditional"] = args ? args.ignoreConditional : undefined;
            inputs["ignoreIeReload"] = args ? args.ignoreIeReload : undefined;
            inputs["ignoreIms"] = args ? args.ignoreIms : undefined;
            inputs["ignorePnc"] = args ? args.ignorePnc : undefined;
            inputs["maxObjectSize"] = args ? args.maxObjectSize : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["minTtl"] = args ? args.minTtl : undefined;
            inputs["negRespTime"] = args ? args.negRespTime : undefined;
            inputs["revalPnc"] = args ? args.revalPnc : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WanoptWebcache.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WanoptWebcache resources.
 */
export interface WanoptWebcacheState {
    /**
     * Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
     */
    alwaysRevalidate?: pulumi.Input<string>;
    /**
     * Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
     */
    cacheByDefault?: pulumi.Input<string>;
    /**
     * Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
     */
    cacheCookie?: pulumi.Input<string>;
    /**
     * Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
     */
    cacheExpired?: pulumi.Input<string>;
    /**
     * Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
     */
    defaultTtl?: pulumi.Input<number>;
    /**
     * Enable/disable external Web caching. Valid values: `enable`, `disable`.
     */
    external?: pulumi.Input<string>;
    /**
     * Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
     */
    freshFactor?: pulumi.Input<number>;
    /**
     * Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
     */
    hostValidate?: pulumi.Input<string>;
    /**
     * Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
     */
    ignoreConditional?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
     */
    ignoreIeReload?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
     */
    ignoreIms?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
     */
    ignorePnc?: pulumi.Input<string>;
    /**
     * Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
     */
    maxObjectSize?: pulumi.Input<number>;
    /**
     * Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
     */
    minTtl?: pulumi.Input<number>;
    /**
     * Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
     */
    negRespTime?: pulumi.Input<number>;
    /**
     * Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
     */
    revalPnc?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WanoptWebcache resource.
 */
export interface WanoptWebcacheArgs {
    /**
     * Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
     */
    alwaysRevalidate?: pulumi.Input<string>;
    /**
     * Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
     */
    cacheByDefault?: pulumi.Input<string>;
    /**
     * Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
     */
    cacheCookie?: pulumi.Input<string>;
    /**
     * Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
     */
    cacheExpired?: pulumi.Input<string>;
    /**
     * Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
     */
    defaultTtl?: pulumi.Input<number>;
    /**
     * Enable/disable external Web caching. Valid values: `enable`, `disable`.
     */
    external?: pulumi.Input<string>;
    /**
     * Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
     */
    freshFactor?: pulumi.Input<number>;
    /**
     * Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
     */
    hostValidate?: pulumi.Input<string>;
    /**
     * Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
     */
    ignoreConditional?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
     */
    ignoreIeReload?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
     */
    ignoreIms?: pulumi.Input<string>;
    /**
     * Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
     */
    ignorePnc?: pulumi.Input<string>;
    /**
     * Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
     */
    maxObjectSize?: pulumi.Input<number>;
    /**
     * Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
     */
    minTtl?: pulumi.Input<number>;
    /**
     * Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
     */
    negRespTime?: pulumi.Input<number>;
    /**
     * Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
     */
    revalPnc?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
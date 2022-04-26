// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Exempt URLs from web proxy forwarding and caching.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi_fortios from "@lubyou/pulumi-fortios";
 *
 * const trname2 = new fortios.WebProxyForwardServer("trname2", {
 *     addrType: "fqdn",
 *     healthcheck: "disable",
 *     ip: "0.0.0.0",
 *     monitor: "http://www.google.com",
 *     port: 3128,
 *     serverDownOption: "block",
 * });
 * const trname = new fortios.WebProxyUrlMatch("trname", {
 *     cacheExemption: "disable",
 *     forwardServer: trname2.name,
 *     status: "enable",
 *     urlPattern: "/examples/servlet/*Servlet",
 * });
 * ```
 *
 * ## Import
 *
 * WebProxy UrlMatch can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/webProxyUrlMatch:WebProxyUrlMatch labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/webProxyUrlMatch:WebProxyUrlMatch labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WebProxyUrlMatch extends pulumi.CustomResource {
    /**
     * Get an existing WebProxyUrlMatch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebProxyUrlMatchState, opts?: pulumi.CustomResourceOptions): WebProxyUrlMatch {
        return new WebProxyUrlMatch(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webProxyUrlMatch:WebProxyUrlMatch';

    /**
     * Returns true if the given object is an instance of WebProxyUrlMatch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebProxyUrlMatch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebProxyUrlMatch.__pulumiType;
    }

    /**
     * Enable/disable exempting this URL pattern from caching. Valid values: `enable`, `disable`.
     */
    public readonly cacheExemption!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Forward server name.
     */
    public readonly forwardServer!: pulumi.Output<string>;
    /**
     * Configure a name for the URL to be exempted.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable exempting the URLs matching the URL pattern from web proxy forwarding and caching. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * URL pattern to be exempted from web proxy forwarding and caching.
     */
    public readonly urlPattern!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WebProxyUrlMatch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebProxyUrlMatchArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebProxyUrlMatchArgs | WebProxyUrlMatchState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebProxyUrlMatchState | undefined;
            resourceInputs["cacheExemption"] = state ? state.cacheExemption : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["forwardServer"] = state ? state.forwardServer : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["urlPattern"] = state ? state.urlPattern : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebProxyUrlMatchArgs | undefined;
            if ((!args || args.urlPattern === undefined) && !opts.urn) {
                throw new Error("Missing required property 'urlPattern'");
            }
            resourceInputs["cacheExemption"] = args ? args.cacheExemption : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["forwardServer"] = args ? args.forwardServer : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["urlPattern"] = args ? args.urlPattern : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebProxyUrlMatch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebProxyUrlMatch resources.
 */
export interface WebProxyUrlMatchState {
    /**
     * Enable/disable exempting this URL pattern from caching. Valid values: `enable`, `disable`.
     */
    cacheExemption?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Forward server name.
     */
    forwardServer?: pulumi.Input<string>;
    /**
     * Configure a name for the URL to be exempted.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable exempting the URLs matching the URL pattern from web proxy forwarding and caching. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * URL pattern to be exempted from web proxy forwarding and caching.
     */
    urlPattern?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebProxyUrlMatch resource.
 */
export interface WebProxyUrlMatchArgs {
    /**
     * Enable/disable exempting this URL pattern from caching. Valid values: `enable`, `disable`.
     */
    cacheExemption?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Forward server name.
     */
    forwardServer?: pulumi.Input<string>;
    /**
     * Configure a name for the URL to be exempted.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable exempting the URLs matching the URL pattern from web proxy forwarding and caching. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * URL pattern to be exempted from web proxy forwarding and caching.
     */
    urlPattern: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

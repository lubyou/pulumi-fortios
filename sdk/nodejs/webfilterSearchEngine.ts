// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure web filter search engines.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WebfilterSearchEngine("trname", {
 *     charset: "utf-8",
 *     hostname: "sg.eiwuc.com",
 *     query: "sc=",
 *     safesearch: "disable",
 *     url: "^\\/f",
 * });
 * ```
 *
 * ## Import
 *
 * Webfilter SearchEngine can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/webfilterSearchEngine:WebfilterSearchEngine labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WebfilterSearchEngine extends pulumi.CustomResource {
    /**
     * Get an existing WebfilterSearchEngine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebfilterSearchEngineState, opts?: pulumi.CustomResourceOptions): WebfilterSearchEngine {
        return new WebfilterSearchEngine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webfilterSearchEngine:WebfilterSearchEngine';

    /**
     * Returns true if the given object is an instance of WebfilterSearchEngine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebfilterSearchEngine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebfilterSearchEngine.__pulumiType;
    }

    /**
     * Search engine charset. Valid values: `utf-8`, `gb2312`.
     */
    public readonly charset!: pulumi.Output<string>;
    /**
     * Hostname (regular expression).
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * Search engine name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Code used to prefix a query (must end with an equals character).
     */
    public readonly query!: pulumi.Output<string>;
    /**
     * Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header. Valid values: `disable`, `url`, `header`.
     */
    public readonly safesearch!: pulumi.Output<string>;
    /**
     * Safe search parameter used in the URL.
     */
    public readonly safesearchStr!: pulumi.Output<string>;
    /**
     * URL (regular expression).
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WebfilterSearchEngine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WebfilterSearchEngineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebfilterSearchEngineArgs | WebfilterSearchEngineState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebfilterSearchEngineState | undefined;
            inputs["charset"] = state ? state.charset : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["query"] = state ? state.query : undefined;
            inputs["safesearch"] = state ? state.safesearch : undefined;
            inputs["safesearchStr"] = state ? state.safesearchStr : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebfilterSearchEngineArgs | undefined;
            inputs["charset"] = args ? args.charset : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["query"] = args ? args.query : undefined;
            inputs["safesearch"] = args ? args.safesearch : undefined;
            inputs["safesearchStr"] = args ? args.safesearchStr : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebfilterSearchEngine.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebfilterSearchEngine resources.
 */
export interface WebfilterSearchEngineState {
    /**
     * Search engine charset. Valid values: `utf-8`, `gb2312`.
     */
    charset?: pulumi.Input<string>;
    /**
     * Hostname (regular expression).
     */
    hostname?: pulumi.Input<string>;
    /**
     * Search engine name.
     */
    name?: pulumi.Input<string>;
    /**
     * Code used to prefix a query (must end with an equals character).
     */
    query?: pulumi.Input<string>;
    /**
     * Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header. Valid values: `disable`, `url`, `header`.
     */
    safesearch?: pulumi.Input<string>;
    /**
     * Safe search parameter used in the URL.
     */
    safesearchStr?: pulumi.Input<string>;
    /**
     * URL (regular expression).
     */
    url?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebfilterSearchEngine resource.
 */
export interface WebfilterSearchEngineArgs {
    /**
     * Search engine charset. Valid values: `utf-8`, `gb2312`.
     */
    charset?: pulumi.Input<string>;
    /**
     * Hostname (regular expression).
     */
    hostname?: pulumi.Input<string>;
    /**
     * Search engine name.
     */
    name?: pulumi.Input<string>;
    /**
     * Code used to prefix a query (must end with an equals character).
     */
    query?: pulumi.Input<string>;
    /**
     * Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header. Valid values: `disable`, `url`, `header`.
     */
    safesearch?: pulumi.Input<string>;
    /**
     * Safe search parameter used in the URL.
     */
    safesearchStr?: pulumi.Input<string>;
    /**
     * URL (regular expression).
     */
    url?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

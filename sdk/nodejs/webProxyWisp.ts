// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure Wireless Internet service provider (WISP) servers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WebProxyWisp("trname", {
 *     maxConnections: 64,
 *     outgoingIp: "0.0.0.0",
 *     serverIp: "1.1.1.1",
 *     serverPort: 15868,
 *     timeout: 5,
 * });
 * ```
 *
 * ## Import
 *
 * WebProxy Wisp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/webProxyWisp:WebProxyWisp labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WebProxyWisp extends pulumi.CustomResource {
    /**
     * Get an existing WebProxyWisp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebProxyWispState, opts?: pulumi.CustomResourceOptions): WebProxyWisp {
        return new WebProxyWisp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webProxyWisp:WebProxyWisp';

    /**
     * Returns true if the given object is an instance of WebProxyWisp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebProxyWisp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebProxyWisp.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Maximum number of web proxy WISP connections (4 - 4096, default = 64).
     */
    public readonly maxConnections!: pulumi.Output<number>;
    /**
     * Server name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * WISP outgoing IP address.
     */
    public readonly outgoingIp!: pulumi.Output<string>;
    /**
     * WISP server IP address.
     */
    public readonly serverIp!: pulumi.Output<string>;
    /**
     * WISP server port (1 - 65535, default = 15868).
     */
    public readonly serverPort!: pulumi.Output<number>;
    /**
     * Period of time before WISP requests time out (1 - 15 sec, default = 5).
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WebProxyWisp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebProxyWispArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebProxyWispArgs | WebProxyWispState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebProxyWispState | undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["maxConnections"] = state ? state.maxConnections : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["outgoingIp"] = state ? state.outgoingIp : undefined;
            inputs["serverIp"] = state ? state.serverIp : undefined;
            inputs["serverPort"] = state ? state.serverPort : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebProxyWispArgs | undefined;
            if ((!args || args.serverIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverIp'");
            }
            if ((!args || args.serverPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverPort'");
            }
            inputs["comment"] = args ? args.comment : undefined;
            inputs["maxConnections"] = args ? args.maxConnections : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["outgoingIp"] = args ? args.outgoingIp : undefined;
            inputs["serverIp"] = args ? args.serverIp : undefined;
            inputs["serverPort"] = args ? args.serverPort : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebProxyWisp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebProxyWisp resources.
 */
export interface WebProxyWispState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Maximum number of web proxy WISP connections (4 - 4096, default = 64).
     */
    maxConnections?: pulumi.Input<number>;
    /**
     * Server name.
     */
    name?: pulumi.Input<string>;
    /**
     * WISP outgoing IP address.
     */
    outgoingIp?: pulumi.Input<string>;
    /**
     * WISP server IP address.
     */
    serverIp?: pulumi.Input<string>;
    /**
     * WISP server port (1 - 65535, default = 15868).
     */
    serverPort?: pulumi.Input<number>;
    /**
     * Period of time before WISP requests time out (1 - 15 sec, default = 5).
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebProxyWisp resource.
 */
export interface WebProxyWispArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Maximum number of web proxy WISP connections (4 - 4096, default = 64).
     */
    maxConnections?: pulumi.Input<number>;
    /**
     * Server name.
     */
    name?: pulumi.Input<string>;
    /**
     * WISP outgoing IP address.
     */
    outgoingIp?: pulumi.Input<string>;
    /**
     * WISP server IP address.
     */
    serverIp: pulumi.Input<string>;
    /**
     * WISP server port (1 - 65535, default = 15868).
     */
    serverPort: pulumi.Input<number>;
    /**
     * Period of time before WISP requests time out (1 - 15 sec, default = 5).
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

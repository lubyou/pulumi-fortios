// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure forward-server addresses.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WebProxyForwardServer("trname", {
 *     addrType: "fqdn",
 *     healthcheck: "disable",
 *     ip: "0.0.0.0",
 *     monitor: "http://www.google.com",
 *     port: 3128,
 *     serverDownOption: "block",
 * });
 * ```
 *
 * ## Import
 *
 * WebProxy ForwardServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/webProxyForwardServer:WebProxyForwardServer labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WebProxyForwardServer extends pulumi.CustomResource {
    /**
     * Get an existing WebProxyForwardServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebProxyForwardServerState, opts?: pulumi.CustomResourceOptions): WebProxyForwardServer {
        return new WebProxyForwardServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webProxyForwardServer:WebProxyForwardServer';

    /**
     * Returns true if the given object is an instance of WebProxyForwardServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebProxyForwardServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebProxyForwardServer.__pulumiType;
    }

    /**
     * Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
     */
    public readonly addrType!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string>;
    /**
     * Forward server Fully Qualified Domain Name (FQDN).
     */
    public readonly fqdn!: pulumi.Output<string>;
    /**
     * Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
     */
    public readonly healthcheck!: pulumi.Output<string>;
    /**
     * Forward proxy server IP address.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * URL for forward server health check monitoring (default = http://www.google.com).
     */
    public readonly monitor!: pulumi.Output<string>;
    /**
     * Server name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * HTTP authentication password.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
     */
    public readonly serverDownOption!: pulumi.Output<string>;
    /**
     * HTTP authentication user name.
     */
    public readonly username!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WebProxyForwardServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WebProxyForwardServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebProxyForwardServerArgs | WebProxyForwardServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebProxyForwardServerState | undefined;
            inputs["addrType"] = state ? state.addrType : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["healthcheck"] = state ? state.healthcheck : undefined;
            inputs["ip"] = state ? state.ip : undefined;
            inputs["monitor"] = state ? state.monitor : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["serverDownOption"] = state ? state.serverDownOption : undefined;
            inputs["username"] = state ? state.username : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebProxyForwardServerArgs | undefined;
            inputs["addrType"] = args ? args.addrType : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["fqdn"] = args ? args.fqdn : undefined;
            inputs["healthcheck"] = args ? args.healthcheck : undefined;
            inputs["ip"] = args ? args.ip : undefined;
            inputs["monitor"] = args ? args.monitor : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["serverDownOption"] = args ? args.serverDownOption : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WebProxyForwardServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebProxyForwardServer resources.
 */
export interface WebProxyForwardServerState {
    /**
     * Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
     */
    addrType?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Forward server Fully Qualified Domain Name (FQDN).
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
     */
    healthcheck?: pulumi.Input<string>;
    /**
     * Forward proxy server IP address.
     */
    ip?: pulumi.Input<string>;
    /**
     * URL for forward server health check monitoring (default = http://www.google.com).
     */
    monitor?: pulumi.Input<string>;
    /**
     * Server name.
     */
    name?: pulumi.Input<string>;
    /**
     * HTTP authentication password.
     */
    password?: pulumi.Input<string>;
    /**
     * Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
     */
    port?: pulumi.Input<number>;
    /**
     * Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
     */
    serverDownOption?: pulumi.Input<string>;
    /**
     * HTTP authentication user name.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebProxyForwardServer resource.
 */
export interface WebProxyForwardServerArgs {
    /**
     * Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
     */
    addrType?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Forward server Fully Qualified Domain Name (FQDN).
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
     */
    healthcheck?: pulumi.Input<string>;
    /**
     * Forward proxy server IP address.
     */
    ip?: pulumi.Input<string>;
    /**
     * URL for forward server health check monitoring (default = http://www.google.com).
     */
    monitor?: pulumi.Input<string>;
    /**
     * Server name.
     */
    name?: pulumi.Input<string>;
    /**
     * HTTP authentication password.
     */
    password?: pulumi.Input<string>;
    /**
     * Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
     */
    port?: pulumi.Input<number>;
    /**
     * Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
     */
    serverDownOption?: pulumi.Input<string>;
    /**
     * HTTP authentication user name.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

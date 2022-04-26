// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SSL proxy settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallSslSetting("trname", {
 *     abbreviateHandshake: "enable",
 *     certCacheCapacity: 200,
 *     certCacheTimeout: 10,
 *     kxpQueueThreshold: 16,
 *     noMatchingCipherAction: "bypass",
 *     proxyConnectTimeout: 30,
 *     sessionCacheCapacity: 500,
 *     sessionCacheTimeout: 20,
 *     sslDhBits: "2048",
 *     sslQueueThreshold: 32,
 *     sslSendEmptyFrags: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * FirewallSsl Setting can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallSslSetting:FirewallSslSetting labelname FirewallSslSetting
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallSslSetting:FirewallSslSetting labelname FirewallSslSetting
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallSslSetting extends pulumi.CustomResource {
    /**
     * Get an existing FirewallSslSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallSslSettingState, opts?: pulumi.CustomResourceOptions): FirewallSslSetting {
        return new FirewallSslSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallSslSetting:FirewallSslSetting';

    /**
     * Returns true if the given object is an instance of FirewallSslSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallSslSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallSslSetting.__pulumiType;
    }

    /**
     * Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
     */
    public readonly abbreviateHandshake!: pulumi.Output<string>;
    /**
     * Maximum capacity of the host certificate cache (0 - 500, default = 200).
     */
    public readonly certCacheCapacity!: pulumi.Output<number>;
    /**
     * Time limit to keep certificate cache (1 - 120 min, default = 10).
     */
    public readonly certCacheTimeout!: pulumi.Output<number>;
    /**
     * Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
     */
    public readonly kxpQueueThreshold!: pulumi.Output<number>;
    /**
     * Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
     */
    public readonly noMatchingCipherAction!: pulumi.Output<string>;
    /**
     * Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
     */
    public readonly proxyConnectTimeout!: pulumi.Output<number>;
    /**
     * Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
     */
    public readonly sessionCacheCapacity!: pulumi.Output<number>;
    /**
     * Time limit to keep SSL session state (1 - 60 min, default = 20).
     */
    public readonly sessionCacheTimeout!: pulumi.Output<number>;
    /**
     * Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
     */
    public readonly sslDhBits!: pulumi.Output<string>;
    /**
     * Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
     */
    public readonly sslQueueThreshold!: pulumi.Output<number>;
    /**
     * Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
     */
    public readonly sslSendEmptyFrags!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallSslSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallSslSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallSslSettingArgs | FirewallSslSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallSslSettingState | undefined;
            resourceInputs["abbreviateHandshake"] = state ? state.abbreviateHandshake : undefined;
            resourceInputs["certCacheCapacity"] = state ? state.certCacheCapacity : undefined;
            resourceInputs["certCacheTimeout"] = state ? state.certCacheTimeout : undefined;
            resourceInputs["kxpQueueThreshold"] = state ? state.kxpQueueThreshold : undefined;
            resourceInputs["noMatchingCipherAction"] = state ? state.noMatchingCipherAction : undefined;
            resourceInputs["proxyConnectTimeout"] = state ? state.proxyConnectTimeout : undefined;
            resourceInputs["sessionCacheCapacity"] = state ? state.sessionCacheCapacity : undefined;
            resourceInputs["sessionCacheTimeout"] = state ? state.sessionCacheTimeout : undefined;
            resourceInputs["sslDhBits"] = state ? state.sslDhBits : undefined;
            resourceInputs["sslQueueThreshold"] = state ? state.sslQueueThreshold : undefined;
            resourceInputs["sslSendEmptyFrags"] = state ? state.sslSendEmptyFrags : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallSslSettingArgs | undefined;
            if ((!args || args.certCacheCapacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certCacheCapacity'");
            }
            if ((!args || args.certCacheTimeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certCacheTimeout'");
            }
            if ((!args || args.noMatchingCipherAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'noMatchingCipherAction'");
            }
            if ((!args || args.proxyConnectTimeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'proxyConnectTimeout'");
            }
            if ((!args || args.sessionCacheCapacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sessionCacheCapacity'");
            }
            if ((!args || args.sessionCacheTimeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sessionCacheTimeout'");
            }
            if ((!args || args.sslDhBits === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sslDhBits'");
            }
            if ((!args || args.sslSendEmptyFrags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sslSendEmptyFrags'");
            }
            resourceInputs["abbreviateHandshake"] = args ? args.abbreviateHandshake : undefined;
            resourceInputs["certCacheCapacity"] = args ? args.certCacheCapacity : undefined;
            resourceInputs["certCacheTimeout"] = args ? args.certCacheTimeout : undefined;
            resourceInputs["kxpQueueThreshold"] = args ? args.kxpQueueThreshold : undefined;
            resourceInputs["noMatchingCipherAction"] = args ? args.noMatchingCipherAction : undefined;
            resourceInputs["proxyConnectTimeout"] = args ? args.proxyConnectTimeout : undefined;
            resourceInputs["sessionCacheCapacity"] = args ? args.sessionCacheCapacity : undefined;
            resourceInputs["sessionCacheTimeout"] = args ? args.sessionCacheTimeout : undefined;
            resourceInputs["sslDhBits"] = args ? args.sslDhBits : undefined;
            resourceInputs["sslQueueThreshold"] = args ? args.sslQueueThreshold : undefined;
            resourceInputs["sslSendEmptyFrags"] = args ? args.sslSendEmptyFrags : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallSslSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallSslSetting resources.
 */
export interface FirewallSslSettingState {
    /**
     * Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
     */
    abbreviateHandshake?: pulumi.Input<string>;
    /**
     * Maximum capacity of the host certificate cache (0 - 500, default = 200).
     */
    certCacheCapacity?: pulumi.Input<number>;
    /**
     * Time limit to keep certificate cache (1 - 120 min, default = 10).
     */
    certCacheTimeout?: pulumi.Input<number>;
    /**
     * Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
     */
    kxpQueueThreshold?: pulumi.Input<number>;
    /**
     * Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
     */
    noMatchingCipherAction?: pulumi.Input<string>;
    /**
     * Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
     */
    proxyConnectTimeout?: pulumi.Input<number>;
    /**
     * Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
     */
    sessionCacheCapacity?: pulumi.Input<number>;
    /**
     * Time limit to keep SSL session state (1 - 60 min, default = 20).
     */
    sessionCacheTimeout?: pulumi.Input<number>;
    /**
     * Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
     */
    sslDhBits?: pulumi.Input<string>;
    /**
     * Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
     */
    sslQueueThreshold?: pulumi.Input<number>;
    /**
     * Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
     */
    sslSendEmptyFrags?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallSslSetting resource.
 */
export interface FirewallSslSettingArgs {
    /**
     * Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
     */
    abbreviateHandshake?: pulumi.Input<string>;
    /**
     * Maximum capacity of the host certificate cache (0 - 500, default = 200).
     */
    certCacheCapacity: pulumi.Input<number>;
    /**
     * Time limit to keep certificate cache (1 - 120 min, default = 10).
     */
    certCacheTimeout: pulumi.Input<number>;
    /**
     * Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
     */
    kxpQueueThreshold?: pulumi.Input<number>;
    /**
     * Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
     */
    noMatchingCipherAction: pulumi.Input<string>;
    /**
     * Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
     */
    proxyConnectTimeout: pulumi.Input<number>;
    /**
     * Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
     */
    sessionCacheCapacity: pulumi.Input<number>;
    /**
     * Time limit to keep SSL session state (1 - 60 min, default = 20).
     */
    sessionCacheTimeout: pulumi.Input<number>;
    /**
     * Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
     */
    sslDhBits: pulumi.Input<string>;
    /**
     * Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
     */
    sslQueueThreshold?: pulumi.Input<number>;
    /**
     * Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
     */
    sslSendEmptyFrags: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

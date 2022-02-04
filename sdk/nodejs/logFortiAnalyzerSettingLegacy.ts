// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to configure configure logging to FortiAnalyzer log management devices.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.LogFortianalyzerSetting`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const test1 = new fortios.LogFortiAnalyzerSettingLegacy("test1", {
 *     encAlgorithm: "high-medium",
 *     hmacAlgorithm: "sha256",
 *     reliable: "enable",
 *     server: "10.2.2.99",
 *     sourceIp: "10.2.2.99",
 *     status: "enable",
 *     uploadOption: "realtime",
 * });
 * ```
 */
export class LogFortiAnalyzerSettingLegacy extends pulumi.CustomResource {
    /**
     * Get an existing LogFortiAnalyzerSettingLegacy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogFortiAnalyzerSettingLegacyState, opts?: pulumi.CustomResourceOptions): LogFortiAnalyzerSettingLegacy {
        return new LogFortiAnalyzerSettingLegacy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logFortiAnalyzerSettingLegacy:LogFortiAnalyzerSettingLegacy';

    /**
     * Returns true if the given object is an instance of LogFortiAnalyzerSettingLegacy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogFortiAnalyzerSettingLegacy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogFortiAnalyzerSettingLegacy.__pulumiType;
    }

    /**
     * Enable/disable sending FortiAnalyzer log data with SSL encryption.
     */
    public readonly encAlgorithm!: pulumi.Output<string>;
    /**
     * FortiAnalyzer IPsec tunnel HMAC algorithm.
     */
    public readonly hmacAlgorithm!: pulumi.Output<string>;
    /**
     * Enable/disable reliable logging to FortiAnalyzer.
     */
    public readonly reliable!: pulumi.Output<string>;
    /**
     * The remote FortiAnalyzer.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Enable/disable logging to FortiAnalyzer.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Enable/disable logging to hard disk and then uploading to FortiAnalyzer.
     */
    public readonly uploadOption!: pulumi.Output<string>;

    /**
     * Create a LogFortiAnalyzerSettingLegacy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogFortiAnalyzerSettingLegacyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogFortiAnalyzerSettingLegacyArgs | LogFortiAnalyzerSettingLegacyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogFortiAnalyzerSettingLegacyState | undefined;
            resourceInputs["encAlgorithm"] = state ? state.encAlgorithm : undefined;
            resourceInputs["hmacAlgorithm"] = state ? state.hmacAlgorithm : undefined;
            resourceInputs["reliable"] = state ? state.reliable : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["uploadOption"] = state ? state.uploadOption : undefined;
        } else {
            const args = argsOrState as LogFortiAnalyzerSettingLegacyArgs | undefined;
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["encAlgorithm"] = args ? args.encAlgorithm : undefined;
            resourceInputs["hmacAlgorithm"] = args ? args.hmacAlgorithm : undefined;
            resourceInputs["reliable"] = args ? args.reliable : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["uploadOption"] = args ? args.uploadOption : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogFortiAnalyzerSettingLegacy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogFortiAnalyzerSettingLegacy resources.
 */
export interface LogFortiAnalyzerSettingLegacyState {
    /**
     * Enable/disable sending FortiAnalyzer log data with SSL encryption.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * FortiAnalyzer IPsec tunnel HMAC algorithm.
     */
    hmacAlgorithm?: pulumi.Input<string>;
    /**
     * Enable/disable reliable logging to FortiAnalyzer.
     */
    reliable?: pulumi.Input<string>;
    /**
     * The remote FortiAnalyzer.
     */
    server?: pulumi.Input<string>;
    /**
     * Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable logging to FortiAnalyzer.
     */
    status?: pulumi.Input<string>;
    /**
     * Enable/disable logging to hard disk and then uploading to FortiAnalyzer.
     */
    uploadOption?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogFortiAnalyzerSettingLegacy resource.
 */
export interface LogFortiAnalyzerSettingLegacyArgs {
    /**
     * Enable/disable sending FortiAnalyzer log data with SSL encryption.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * FortiAnalyzer IPsec tunnel HMAC algorithm.
     */
    hmacAlgorithm?: pulumi.Input<string>;
    /**
     * Enable/disable reliable logging to FortiAnalyzer.
     */
    reliable?: pulumi.Input<string>;
    /**
     * The remote FortiAnalyzer.
     */
    server?: pulumi.Input<string>;
    /**
     * Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable logging to FortiAnalyzer.
     */
    status: pulumi.Input<string>;
    /**
     * Enable/disable logging to hard disk and then uploading to FortiAnalyzer.
     */
    uploadOption?: pulumi.Input<string>;
}

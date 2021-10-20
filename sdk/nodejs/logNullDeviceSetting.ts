// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Settings for null device logging.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.LogNullDeviceSetting("trname", {
 *     status: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * LogNullDevice Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/logNullDeviceSetting:LogNullDeviceSetting labelname LogNullDeviceSetting
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class LogNullDeviceSetting extends pulumi.CustomResource {
    /**
     * Get an existing LogNullDeviceSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogNullDeviceSettingState, opts?: pulumi.CustomResourceOptions): LogNullDeviceSetting {
        return new LogNullDeviceSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logNullDeviceSetting:LogNullDeviceSetting';

    /**
     * Returns true if the given object is an instance of LogNullDeviceSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogNullDeviceSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogNullDeviceSetting.__pulumiType;
    }

    /**
     * Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved). Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a LogNullDeviceSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogNullDeviceSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogNullDeviceSettingArgs | LogNullDeviceSettingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogNullDeviceSettingState | undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LogNullDeviceSettingArgs | undefined;
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            inputs["status"] = args ? args.status : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LogNullDeviceSetting.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogNullDeviceSetting resources.
 */
export interface LogNullDeviceSettingState {
    /**
     * Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved). Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogNullDeviceSetting resource.
 */
export interface LogNullDeviceSettingArgs {
    /**
     * Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved). Valid values: `enable`, `disable`.
     */
    status: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
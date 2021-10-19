// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source. Applies to FortiOS Version `<= 6.2.0`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.DlpFpSensitivity("trname", {});
 * ```
 *
 * ## Import
 *
 * Dlp FpSensitivity can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/dlpFpSensitivity:DlpFpSensitivity labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class DlpFpSensitivity extends pulumi.CustomResource {
    /**
     * Get an existing DlpFpSensitivity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DlpFpSensitivityState, opts?: pulumi.CustomResourceOptions): DlpFpSensitivity {
        return new DlpFpSensitivity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/dlpFpSensitivity:DlpFpSensitivity';

    /**
     * Returns true if the given object is an instance of DlpFpSensitivity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DlpFpSensitivity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DlpFpSensitivity.__pulumiType;
    }

    /**
     * DLP Sensitivity Levels.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a DlpFpSensitivity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DlpFpSensitivityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DlpFpSensitivityArgs | DlpFpSensitivityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DlpFpSensitivityState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as DlpFpSensitivityArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DlpFpSensitivity.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DlpFpSensitivity resources.
 */
export interface DlpFpSensitivityState {
    /**
     * DLP Sensitivity Levels.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DlpFpSensitivity resource.
 */
export interface DlpFpSensitivityArgs {
    /**
     * DLP Sensitivity Levels.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source.
 *
 * ## Import
 *
 * Dlp Sensitivity can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/dlpSensitivity:DlpSensitivity labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class DlpSensitivity extends pulumi.CustomResource {
    /**
     * Get an existing DlpSensitivity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DlpSensitivityState, opts?: pulumi.CustomResourceOptions): DlpSensitivity {
        return new DlpSensitivity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/dlpSensitivity:DlpSensitivity';

    /**
     * Returns true if the given object is an instance of DlpSensitivity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DlpSensitivity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DlpSensitivity.__pulumiType;
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
     * Create a DlpSensitivity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DlpSensitivityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DlpSensitivityArgs | DlpSensitivityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DlpSensitivityState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as DlpSensitivityArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DlpSensitivity.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DlpSensitivity resources.
 */
export interface DlpSensitivityState {
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
 * The set of arguments for constructing a DlpSensitivity resource.
 */
export interface DlpSensitivityArgs {
    /**
     * DLP Sensitivity Levels.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

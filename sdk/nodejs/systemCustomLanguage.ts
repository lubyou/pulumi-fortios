// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure custom languages.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemCustomLanguage("trname", {
 *     filename: "en",
 * });
 * ```
 *
 * ## Import
 *
 * System CustomLanguage can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemCustomLanguage:SystemCustomLanguage labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemCustomLanguage extends pulumi.CustomResource {
    /**
     * Get an existing SystemCustomLanguage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemCustomLanguageState, opts?: pulumi.CustomResourceOptions): SystemCustomLanguage {
        return new SystemCustomLanguage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemCustomLanguage:SystemCustomLanguage';

    /**
     * Returns true if the given object is an instance of SystemCustomLanguage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemCustomLanguage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemCustomLanguage.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Custom language file path.
     */
    public readonly filename!: pulumi.Output<string>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemCustomLanguage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemCustomLanguageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemCustomLanguageArgs | SystemCustomLanguageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemCustomLanguageState | undefined;
            inputs["comments"] = state ? state.comments : undefined;
            inputs["filename"] = state ? state.filename : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemCustomLanguageArgs | undefined;
            if ((!args || args.filename === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filename'");
            }
            inputs["comments"] = args ? args.comments : undefined;
            inputs["filename"] = args ? args.filename : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemCustomLanguage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemCustomLanguage resources.
 */
export interface SystemCustomLanguageState {
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Custom language file path.
     */
    filename?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemCustomLanguage resource.
 */
export interface SystemCustomLanguageArgs {
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Custom language file path.
     */
    filename: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
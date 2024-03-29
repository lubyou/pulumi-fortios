// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ExtenderControllerExtenderProfile extends pulumi.CustomResource {
    /**
     * Get an existing ExtenderControllerExtenderProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtenderControllerExtenderProfileState, opts?: pulumi.CustomResourceOptions): ExtenderControllerExtenderProfile {
        return new ExtenderControllerExtenderProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/extenderControllerExtenderProfile:ExtenderControllerExtenderProfile';

    /**
     * Returns true if the given object is an instance of ExtenderControllerExtenderProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExtenderControllerExtenderProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExtenderControllerExtenderProfile.__pulumiType;
    }

    public readonly allowaccess!: pulumi.Output<string>;
    public readonly bandwidthLimit!: pulumi.Output<number>;
    public readonly cellular!: pulumi.Output<outputs.ExtenderControllerExtenderProfileCellular>;
    public readonly enforceBandwidth!: pulumi.Output<string>;
    public readonly extension!: pulumi.Output<string>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly lanExtension!: pulumi.Output<outputs.ExtenderControllerExtenderProfileLanExtension>;
    public readonly loginPassword!: pulumi.Output<string | undefined>;
    public readonly loginPasswordChange!: pulumi.Output<string>;
    public readonly model!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a ExtenderControllerExtenderProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExtenderControllerExtenderProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtenderControllerExtenderProfileArgs | ExtenderControllerExtenderProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExtenderControllerExtenderProfileState | undefined;
            resourceInputs["allowaccess"] = state ? state.allowaccess : undefined;
            resourceInputs["bandwidthLimit"] = state ? state.bandwidthLimit : undefined;
            resourceInputs["cellular"] = state ? state.cellular : undefined;
            resourceInputs["enforceBandwidth"] = state ? state.enforceBandwidth : undefined;
            resourceInputs["extension"] = state ? state.extension : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["lanExtension"] = state ? state.lanExtension : undefined;
            resourceInputs["loginPassword"] = state ? state.loginPassword : undefined;
            resourceInputs["loginPasswordChange"] = state ? state.loginPasswordChange : undefined;
            resourceInputs["model"] = state ? state.model : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ExtenderControllerExtenderProfileArgs | undefined;
            resourceInputs["allowaccess"] = args ? args.allowaccess : undefined;
            resourceInputs["bandwidthLimit"] = args ? args.bandwidthLimit : undefined;
            resourceInputs["cellular"] = args ? args.cellular : undefined;
            resourceInputs["enforceBandwidth"] = args ? args.enforceBandwidth : undefined;
            resourceInputs["extension"] = args ? args.extension : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["lanExtension"] = args ? args.lanExtension : undefined;
            resourceInputs["loginPassword"] = args ? args.loginPassword : undefined;
            resourceInputs["loginPasswordChange"] = args ? args.loginPasswordChange : undefined;
            resourceInputs["model"] = args ? args.model : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExtenderControllerExtenderProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExtenderControllerExtenderProfile resources.
 */
export interface ExtenderControllerExtenderProfileState {
    allowaccess?: pulumi.Input<string>;
    bandwidthLimit?: pulumi.Input<number>;
    cellular?: pulumi.Input<inputs.ExtenderControllerExtenderProfileCellular>;
    enforceBandwidth?: pulumi.Input<string>;
    extension?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    lanExtension?: pulumi.Input<inputs.ExtenderControllerExtenderProfileLanExtension>;
    loginPassword?: pulumi.Input<string>;
    loginPasswordChange?: pulumi.Input<string>;
    model?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExtenderControllerExtenderProfile resource.
 */
export interface ExtenderControllerExtenderProfileArgs {
    allowaccess?: pulumi.Input<string>;
    bandwidthLimit?: pulumi.Input<number>;
    cellular?: pulumi.Input<inputs.ExtenderControllerExtenderProfileCellular>;
    enforceBandwidth?: pulumi.Input<string>;
    extension?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    lanExtension?: pulumi.Input<inputs.ExtenderControllerExtenderProfileLanExtension>;
    loginPassword?: pulumi.Input<string>;
    loginPasswordChange?: pulumi.Input<string>;
    model?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

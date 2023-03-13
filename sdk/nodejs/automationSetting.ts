// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AutomationSetting extends pulumi.CustomResource {
    /**
     * Get an existing AutomationSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutomationSettingState, opts?: pulumi.CustomResourceOptions): AutomationSetting {
        return new AutomationSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/automationSetting:AutomationSetting';

    /**
     * Returns true if the given object is an instance of AutomationSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutomationSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutomationSetting.__pulumiType;
    }

    public readonly maxConcurrentStitches!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a AutomationSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AutomationSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutomationSettingArgs | AutomationSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutomationSettingState | undefined;
            resourceInputs["maxConcurrentStitches"] = state ? state.maxConcurrentStitches : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as AutomationSettingArgs | undefined;
            resourceInputs["maxConcurrentStitches"] = args ? args.maxConcurrentStitches : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AutomationSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutomationSetting resources.
 */
export interface AutomationSettingState {
    maxConcurrentStitches?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AutomationSetting resource.
 */
export interface AutomationSettingArgs {
    maxConcurrentStitches?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
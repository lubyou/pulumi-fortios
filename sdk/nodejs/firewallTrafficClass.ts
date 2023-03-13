// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FirewallTrafficClass extends pulumi.CustomResource {
    /**
     * Get an existing FirewallTrafficClass resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallTrafficClassState, opts?: pulumi.CustomResourceOptions): FirewallTrafficClass {
        return new FirewallTrafficClass(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallTrafficClass:FirewallTrafficClass';

    /**
     * Returns true if the given object is an instance of FirewallTrafficClass.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallTrafficClass {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallTrafficClass.__pulumiType;
    }

    public readonly classId!: pulumi.Output<number>;
    public readonly className!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallTrafficClass resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallTrafficClassArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallTrafficClassArgs | FirewallTrafficClassState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallTrafficClassState | undefined;
            resourceInputs["classId"] = state ? state.classId : undefined;
            resourceInputs["className"] = state ? state.className : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallTrafficClassArgs | undefined;
            resourceInputs["classId"] = args ? args.classId : undefined;
            resourceInputs["className"] = args ? args.className : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallTrafficClass.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallTrafficClass resources.
 */
export interface FirewallTrafficClassState {
    classId?: pulumi.Input<number>;
    className?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallTrafficClass resource.
 */
export interface FirewallTrafficClassArgs {
    classId?: pulumi.Input<number>;
    className?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

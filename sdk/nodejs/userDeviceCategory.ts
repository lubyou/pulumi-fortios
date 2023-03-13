// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class UserDeviceCategory extends pulumi.CustomResource {
    /**
     * Get an existing UserDeviceCategory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserDeviceCategoryState, opts?: pulumi.CustomResourceOptions): UserDeviceCategory {
        return new UserDeviceCategory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/userDeviceCategory:UserDeviceCategory';

    /**
     * Returns true if the given object is an instance of UserDeviceCategory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserDeviceCategory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserDeviceCategory.__pulumiType;
    }

    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly desc!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a UserDeviceCategory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserDeviceCategoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserDeviceCategoryArgs | UserDeviceCategoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserDeviceCategoryState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["desc"] = state ? state.desc : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as UserDeviceCategoryArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["desc"] = args ? args.desc : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserDeviceCategory.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserDeviceCategory resources.
 */
export interface UserDeviceCategoryState {
    comment?: pulumi.Input<string>;
    desc?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserDeviceCategory resource.
 */
export interface UserDeviceCategoryArgs {
    comment?: pulumi.Input<string>;
    desc?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

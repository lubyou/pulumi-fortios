// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ExtensionControllerFortigate extends pulumi.CustomResource {
    /**
     * Get an existing ExtensionControllerFortigate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtensionControllerFortigateState, opts?: pulumi.CustomResourceOptions): ExtensionControllerFortigate {
        return new ExtensionControllerFortigate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/extensionControllerFortigate:ExtensionControllerFortigate';

    /**
     * Returns true if the given object is an instance of ExtensionControllerFortigate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExtensionControllerFortigate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExtensionControllerFortigate.__pulumiType;
    }

    public readonly authorized!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string>;
    public readonly deviceId!: pulumi.Output<number>;
    public readonly fosid!: pulumi.Output<string>;
    public readonly hostname!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly profile!: pulumi.Output<string>;
    public readonly vdom!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a ExtensionControllerFortigate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExtensionControllerFortigateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtensionControllerFortigateArgs | ExtensionControllerFortigateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExtensionControllerFortigateState | undefined;
            resourceInputs["authorized"] = state ? state.authorized : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["profile"] = state ? state.profile : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ExtensionControllerFortigateArgs | undefined;
            resourceInputs["authorized"] = args ? args.authorized : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExtensionControllerFortigate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExtensionControllerFortigate resources.
 */
export interface ExtensionControllerFortigateState {
    authorized?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    deviceId?: pulumi.Input<number>;
    fosid?: pulumi.Input<string>;
    hostname?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    profile?: pulumi.Input<string>;
    vdom?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExtensionControllerFortigate resource.
 */
export interface ExtensionControllerFortigateArgs {
    authorized?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    deviceId?: pulumi.Input<number>;
    fosid?: pulumi.Input<string>;
    hostname?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    profile?: pulumi.Input<string>;
    vdom?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

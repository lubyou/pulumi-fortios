// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemStp extends pulumi.CustomResource {
    /**
     * Get an existing SystemStp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemStpState, opts?: pulumi.CustomResourceOptions): SystemStp {
        return new SystemStp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemStp:SystemStp';

    /**
     * Returns true if the given object is an instance of SystemStp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemStp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemStp.__pulumiType;
    }

    public readonly forwardDelay!: pulumi.Output<number>;
    public readonly helloTime!: pulumi.Output<number>;
    public readonly maxAge!: pulumi.Output<number>;
    public readonly maxHops!: pulumi.Output<number>;
    public readonly switchPriority!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemStp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemStpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemStpArgs | SystemStpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemStpState | undefined;
            resourceInputs["forwardDelay"] = state ? state.forwardDelay : undefined;
            resourceInputs["helloTime"] = state ? state.helloTime : undefined;
            resourceInputs["maxAge"] = state ? state.maxAge : undefined;
            resourceInputs["maxHops"] = state ? state.maxHops : undefined;
            resourceInputs["switchPriority"] = state ? state.switchPriority : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemStpArgs | undefined;
            resourceInputs["forwardDelay"] = args ? args.forwardDelay : undefined;
            resourceInputs["helloTime"] = args ? args.helloTime : undefined;
            resourceInputs["maxAge"] = args ? args.maxAge : undefined;
            resourceInputs["maxHops"] = args ? args.maxHops : undefined;
            resourceInputs["switchPriority"] = args ? args.switchPriority : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemStp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemStp resources.
 */
export interface SystemStpState {
    forwardDelay?: pulumi.Input<number>;
    helloTime?: pulumi.Input<number>;
    maxAge?: pulumi.Input<number>;
    maxHops?: pulumi.Input<number>;
    switchPriority?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemStp resource.
 */
export interface SystemStpArgs {
    forwardDelay?: pulumi.Input<number>;
    helloTime?: pulumi.Input<number>;
    maxAge?: pulumi.Input<number>;
    maxHops?: pulumi.Input<number>;
    switchPriority?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

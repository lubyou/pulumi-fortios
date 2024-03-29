// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemFssoPolling extends pulumi.CustomResource {
    /**
     * Get an existing SystemFssoPolling resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemFssoPollingState, opts?: pulumi.CustomResourceOptions): SystemFssoPolling {
        return new SystemFssoPolling(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemFssoPolling:SystemFssoPolling';

    /**
     * Returns true if the given object is an instance of SystemFssoPolling.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemFssoPolling {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemFssoPolling.__pulumiType;
    }

    public readonly authPassword!: pulumi.Output<string | undefined>;
    public readonly authentication!: pulumi.Output<string>;
    public readonly listeningPort!: pulumi.Output<number>;
    public readonly status!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemFssoPolling resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemFssoPollingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemFssoPollingArgs | SystemFssoPollingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemFssoPollingState | undefined;
            resourceInputs["authPassword"] = state ? state.authPassword : undefined;
            resourceInputs["authentication"] = state ? state.authentication : undefined;
            resourceInputs["listeningPort"] = state ? state.listeningPort : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemFssoPollingArgs | undefined;
            resourceInputs["authPassword"] = args?.authPassword ? pulumi.secret(args.authPassword) : undefined;
            resourceInputs["authentication"] = args ? args.authentication : undefined;
            resourceInputs["listeningPort"] = args ? args.listeningPort : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["authPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SystemFssoPolling.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemFssoPolling resources.
 */
export interface SystemFssoPollingState {
    authPassword?: pulumi.Input<string>;
    authentication?: pulumi.Input<string>;
    listeningPort?: pulumi.Input<number>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemFssoPolling resource.
 */
export interface SystemFssoPollingArgs {
    authPassword?: pulumi.Input<string>;
    authentication?: pulumi.Input<string>;
    listeningPort?: pulumi.Input<number>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemSflow extends pulumi.CustomResource {
    /**
     * Get an existing SystemSflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemSflowState, opts?: pulumi.CustomResourceOptions): SystemSflow {
        return new SystemSflow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemSflow:SystemSflow';

    /**
     * Returns true if the given object is an instance of SystemSflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemSflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemSflow.__pulumiType;
    }

    public readonly collectorIp!: pulumi.Output<string>;
    public readonly collectorPort!: pulumi.Output<number>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly sourceIp!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemSflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemSflowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemSflowArgs | SystemSflowState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemSflowState | undefined;
            resourceInputs["collectorIp"] = state ? state.collectorIp : undefined;
            resourceInputs["collectorPort"] = state ? state.collectorPort : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemSflowArgs | undefined;
            if ((!args || args.collectorIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'collectorIp'");
            }
            resourceInputs["collectorIp"] = args ? args.collectorIp : undefined;
            resourceInputs["collectorPort"] = args ? args.collectorPort : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemSflow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemSflow resources.
 */
export interface SystemSflowState {
    collectorIp?: pulumi.Input<string>;
    collectorPort?: pulumi.Input<number>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemSflow resource.
 */
export interface SystemSflowArgs {
    collectorIp: pulumi.Input<string>;
    collectorPort?: pulumi.Input<number>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemNetflow extends pulumi.CustomResource {
    /**
     * Get an existing SystemNetflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemNetflowState, opts?: pulumi.CustomResourceOptions): SystemNetflow {
        return new SystemNetflow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemNetflow:SystemNetflow';

    /**
     * Returns true if the given object is an instance of SystemNetflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemNetflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemNetflow.__pulumiType;
    }

    public readonly activeFlowTimeout!: pulumi.Output<number>;
    public readonly collectorIp!: pulumi.Output<string>;
    public readonly collectorPort!: pulumi.Output<number>;
    public readonly inactiveFlowTimeout!: pulumi.Output<number>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly sourceIp!: pulumi.Output<string>;
    public readonly templateTxCounter!: pulumi.Output<number>;
    public readonly templateTxTimeout!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemNetflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemNetflowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemNetflowArgs | SystemNetflowState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemNetflowState | undefined;
            resourceInputs["activeFlowTimeout"] = state ? state.activeFlowTimeout : undefined;
            resourceInputs["collectorIp"] = state ? state.collectorIp : undefined;
            resourceInputs["collectorPort"] = state ? state.collectorPort : undefined;
            resourceInputs["inactiveFlowTimeout"] = state ? state.inactiveFlowTimeout : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["templateTxCounter"] = state ? state.templateTxCounter : undefined;
            resourceInputs["templateTxTimeout"] = state ? state.templateTxTimeout : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemNetflowArgs | undefined;
            resourceInputs["activeFlowTimeout"] = args ? args.activeFlowTimeout : undefined;
            resourceInputs["collectorIp"] = args ? args.collectorIp : undefined;
            resourceInputs["collectorPort"] = args ? args.collectorPort : undefined;
            resourceInputs["inactiveFlowTimeout"] = args ? args.inactiveFlowTimeout : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["templateTxCounter"] = args ? args.templateTxCounter : undefined;
            resourceInputs["templateTxTimeout"] = args ? args.templateTxTimeout : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemNetflow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemNetflow resources.
 */
export interface SystemNetflowState {
    activeFlowTimeout?: pulumi.Input<number>;
    collectorIp?: pulumi.Input<string>;
    collectorPort?: pulumi.Input<number>;
    inactiveFlowTimeout?: pulumi.Input<number>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    templateTxCounter?: pulumi.Input<number>;
    templateTxTimeout?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemNetflow resource.
 */
export interface SystemNetflowArgs {
    activeFlowTimeout?: pulumi.Input<number>;
    collectorIp?: pulumi.Input<string>;
    collectorPort?: pulumi.Input<number>;
    inactiveFlowTimeout?: pulumi.Input<number>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    templateTxCounter?: pulumi.Input<number>;
    templateTxTimeout?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

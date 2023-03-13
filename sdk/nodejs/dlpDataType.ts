// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class DlpDataType extends pulumi.CustomResource {
    /**
     * Get an existing DlpDataType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DlpDataTypeState, opts?: pulumi.CustomResourceOptions): DlpDataType {
        return new DlpDataType(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/dlpDataType:DlpDataType';

    /**
     * Returns true if the given object is an instance of DlpDataType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DlpDataType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DlpDataType.__pulumiType;
    }

    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly lookAhead!: pulumi.Output<number>;
    public readonly lookBack!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly pattern!: pulumi.Output<string>;
    public readonly transform!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly verify!: pulumi.Output<string>;
    public readonly verifyTransformedPattern!: pulumi.Output<string>;

    /**
     * Create a DlpDataType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DlpDataTypeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DlpDataTypeArgs | DlpDataTypeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DlpDataTypeState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["lookAhead"] = state ? state.lookAhead : undefined;
            resourceInputs["lookBack"] = state ? state.lookBack : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pattern"] = state ? state.pattern : undefined;
            resourceInputs["transform"] = state ? state.transform : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["verify"] = state ? state.verify : undefined;
            resourceInputs["verifyTransformedPattern"] = state ? state.verifyTransformedPattern : undefined;
        } else {
            const args = argsOrState as DlpDataTypeArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["lookAhead"] = args ? args.lookAhead : undefined;
            resourceInputs["lookBack"] = args ? args.lookBack : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pattern"] = args ? args.pattern : undefined;
            resourceInputs["transform"] = args ? args.transform : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["verify"] = args ? args.verify : undefined;
            resourceInputs["verifyTransformedPattern"] = args ? args.verifyTransformedPattern : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DlpDataType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DlpDataType resources.
 */
export interface DlpDataTypeState {
    comment?: pulumi.Input<string>;
    lookAhead?: pulumi.Input<number>;
    lookBack?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    pattern?: pulumi.Input<string>;
    transform?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    verify?: pulumi.Input<string>;
    verifyTransformedPattern?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DlpDataType resource.
 */
export interface DlpDataTypeArgs {
    comment?: pulumi.Input<string>;
    lookAhead?: pulumi.Input<number>;
    lookBack?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    pattern?: pulumi.Input<string>;
    transform?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    verify?: pulumi.Input<string>;
    verifyTransformedPattern?: pulumi.Input<string>;
}
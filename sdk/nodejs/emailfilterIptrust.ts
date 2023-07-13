// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class EmailfilterIptrust extends pulumi.CustomResource {
    /**
     * Get an existing EmailfilterIptrust resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EmailfilterIptrustState, opts?: pulumi.CustomResourceOptions): EmailfilterIptrust {
        return new EmailfilterIptrust(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/emailfilterIptrust:EmailfilterIptrust';

    /**
     * Returns true if the given object is an instance of EmailfilterIptrust.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmailfilterIptrust {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmailfilterIptrust.__pulumiType;
    }

    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly entries!: pulumi.Output<outputs.EmailfilterIptrustEntry[] | undefined>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a EmailfilterIptrust resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EmailfilterIptrustArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EmailfilterIptrustArgs | EmailfilterIptrustState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EmailfilterIptrustState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as EmailfilterIptrustArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EmailfilterIptrust.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EmailfilterIptrust resources.
 */
export interface EmailfilterIptrustState {
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    entries?: pulumi.Input<pulumi.Input<inputs.EmailfilterIptrustEntry>[]>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EmailfilterIptrust resource.
 */
export interface EmailfilterIptrustArgs {
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    entries?: pulumi.Input<pulumi.Input<inputs.EmailfilterIptrustEntry>[]>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

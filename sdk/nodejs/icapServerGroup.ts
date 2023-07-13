// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class IcapServerGroup extends pulumi.CustomResource {
    /**
     * Get an existing IcapServerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IcapServerGroupState, opts?: pulumi.CustomResourceOptions): IcapServerGroup {
        return new IcapServerGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/icapServerGroup:IcapServerGroup';

    /**
     * Returns true if the given object is an instance of IcapServerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IcapServerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IcapServerGroup.__pulumiType;
    }

    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly ldbMethod!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly serverLists!: pulumi.Output<outputs.IcapServerGroupServerList[] | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a IcapServerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IcapServerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IcapServerGroupArgs | IcapServerGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IcapServerGroupState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["ldbMethod"] = state ? state.ldbMethod : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["serverLists"] = state ? state.serverLists : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as IcapServerGroupArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["ldbMethod"] = args ? args.ldbMethod : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serverLists"] = args ? args.serverLists : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IcapServerGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IcapServerGroup resources.
 */
export interface IcapServerGroupState {
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    ldbMethod?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    serverLists?: pulumi.Input<pulumi.Input<inputs.IcapServerGroupServerList>[]>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IcapServerGroup resource.
 */
export interface IcapServerGroupArgs {
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    ldbMethod?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    serverLists?: pulumi.Input<pulumi.Input<inputs.IcapServerGroupServerList>[]>;
    vdomparam?: pulumi.Input<string>;
}

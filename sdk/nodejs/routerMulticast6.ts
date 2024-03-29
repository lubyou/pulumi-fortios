// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class RouterMulticast6 extends pulumi.CustomResource {
    /**
     * Get an existing RouterMulticast6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterMulticast6State, opts?: pulumi.CustomResourceOptions): RouterMulticast6 {
        return new RouterMulticast6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerMulticast6:RouterMulticast6';

    /**
     * Returns true if the given object is an instance of RouterMulticast6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterMulticast6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterMulticast6.__pulumiType;
    }

    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly interfaces!: pulumi.Output<outputs.RouterMulticast6Interface[] | undefined>;
    public readonly multicastPmtu!: pulumi.Output<string>;
    public readonly multicastRouting!: pulumi.Output<string>;
    public readonly pimSmGlobal!: pulumi.Output<outputs.RouterMulticast6PimSmGlobal>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterMulticast6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouterMulticast6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterMulticast6Args | RouterMulticast6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterMulticast6State | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["multicastPmtu"] = state ? state.multicastPmtu : undefined;
            resourceInputs["multicastRouting"] = state ? state.multicastRouting : undefined;
            resourceInputs["pimSmGlobal"] = state ? state.pimSmGlobal : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RouterMulticast6Args | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["multicastPmtu"] = args ? args.multicastPmtu : undefined;
            resourceInputs["multicastRouting"] = args ? args.multicastRouting : undefined;
            resourceInputs["pimSmGlobal"] = args ? args.pimSmGlobal : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterMulticast6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterMulticast6 resources.
 */
export interface RouterMulticast6State {
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    interfaces?: pulumi.Input<pulumi.Input<inputs.RouterMulticast6Interface>[]>;
    multicastPmtu?: pulumi.Input<string>;
    multicastRouting?: pulumi.Input<string>;
    pimSmGlobal?: pulumi.Input<inputs.RouterMulticast6PimSmGlobal>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterMulticast6 resource.
 */
export interface RouterMulticast6Args {
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    interfaces?: pulumi.Input<pulumi.Input<inputs.RouterMulticast6Interface>[]>;
    multicastPmtu?: pulumi.Input<string>;
    multicastRouting?: pulumi.Input<string>;
    pimSmGlobal?: pulumi.Input<inputs.RouterMulticast6PimSmGlobal>;
    vdomparam?: pulumi.Input<string>;
}

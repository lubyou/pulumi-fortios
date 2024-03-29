// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class RouterRipng extends pulumi.CustomResource {
    /**
     * Get an existing RouterRipng resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterRipngState, opts?: pulumi.CustomResourceOptions): RouterRipng {
        return new RouterRipng(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerRipng:RouterRipng';

    /**
     * Returns true if the given object is an instance of RouterRipng.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterRipng {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterRipng.__pulumiType;
    }

    public readonly aggregateAddresses!: pulumi.Output<outputs.RouterRipngAggregateAddress[] | undefined>;
    public readonly defaultInformationOriginate!: pulumi.Output<string>;
    public readonly defaultMetric!: pulumi.Output<number>;
    public readonly distances!: pulumi.Output<outputs.RouterRipngDistance[] | undefined>;
    public readonly distributeLists!: pulumi.Output<outputs.RouterRipngDistributeList[] | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly garbageTimer!: pulumi.Output<number>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly interfaces!: pulumi.Output<outputs.RouterRipngInterface[] | undefined>;
    public readonly maxOutMetric!: pulumi.Output<number>;
    public readonly neighbors!: pulumi.Output<outputs.RouterRipngNeighbor[] | undefined>;
    public readonly networks!: pulumi.Output<outputs.RouterRipngNetwork[] | undefined>;
    public readonly offsetLists!: pulumi.Output<outputs.RouterRipngOffsetList[] | undefined>;
    public readonly passiveInterfaces!: pulumi.Output<outputs.RouterRipngPassiveInterface[] | undefined>;
    public readonly redistributes!: pulumi.Output<outputs.RouterRipngRedistribute[] | undefined>;
    public readonly timeoutTimer!: pulumi.Output<number>;
    public readonly updateTimer!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterRipng resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouterRipngArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterRipngArgs | RouterRipngState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterRipngState | undefined;
            resourceInputs["aggregateAddresses"] = state ? state.aggregateAddresses : undefined;
            resourceInputs["defaultInformationOriginate"] = state ? state.defaultInformationOriginate : undefined;
            resourceInputs["defaultMetric"] = state ? state.defaultMetric : undefined;
            resourceInputs["distances"] = state ? state.distances : undefined;
            resourceInputs["distributeLists"] = state ? state.distributeLists : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["garbageTimer"] = state ? state.garbageTimer : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["maxOutMetric"] = state ? state.maxOutMetric : undefined;
            resourceInputs["neighbors"] = state ? state.neighbors : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["offsetLists"] = state ? state.offsetLists : undefined;
            resourceInputs["passiveInterfaces"] = state ? state.passiveInterfaces : undefined;
            resourceInputs["redistributes"] = state ? state.redistributes : undefined;
            resourceInputs["timeoutTimer"] = state ? state.timeoutTimer : undefined;
            resourceInputs["updateTimer"] = state ? state.updateTimer : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RouterRipngArgs | undefined;
            resourceInputs["aggregateAddresses"] = args ? args.aggregateAddresses : undefined;
            resourceInputs["defaultInformationOriginate"] = args ? args.defaultInformationOriginate : undefined;
            resourceInputs["defaultMetric"] = args ? args.defaultMetric : undefined;
            resourceInputs["distances"] = args ? args.distances : undefined;
            resourceInputs["distributeLists"] = args ? args.distributeLists : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["garbageTimer"] = args ? args.garbageTimer : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["maxOutMetric"] = args ? args.maxOutMetric : undefined;
            resourceInputs["neighbors"] = args ? args.neighbors : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["offsetLists"] = args ? args.offsetLists : undefined;
            resourceInputs["passiveInterfaces"] = args ? args.passiveInterfaces : undefined;
            resourceInputs["redistributes"] = args ? args.redistributes : undefined;
            resourceInputs["timeoutTimer"] = args ? args.timeoutTimer : undefined;
            resourceInputs["updateTimer"] = args ? args.updateTimer : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterRipng.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterRipng resources.
 */
export interface RouterRipngState {
    aggregateAddresses?: pulumi.Input<pulumi.Input<inputs.RouterRipngAggregateAddress>[]>;
    defaultInformationOriginate?: pulumi.Input<string>;
    defaultMetric?: pulumi.Input<number>;
    distances?: pulumi.Input<pulumi.Input<inputs.RouterRipngDistance>[]>;
    distributeLists?: pulumi.Input<pulumi.Input<inputs.RouterRipngDistributeList>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    garbageTimer?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    interfaces?: pulumi.Input<pulumi.Input<inputs.RouterRipngInterface>[]>;
    maxOutMetric?: pulumi.Input<number>;
    neighbors?: pulumi.Input<pulumi.Input<inputs.RouterRipngNeighbor>[]>;
    networks?: pulumi.Input<pulumi.Input<inputs.RouterRipngNetwork>[]>;
    offsetLists?: pulumi.Input<pulumi.Input<inputs.RouterRipngOffsetList>[]>;
    passiveInterfaces?: pulumi.Input<pulumi.Input<inputs.RouterRipngPassiveInterface>[]>;
    redistributes?: pulumi.Input<pulumi.Input<inputs.RouterRipngRedistribute>[]>;
    timeoutTimer?: pulumi.Input<number>;
    updateTimer?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterRipng resource.
 */
export interface RouterRipngArgs {
    aggregateAddresses?: pulumi.Input<pulumi.Input<inputs.RouterRipngAggregateAddress>[]>;
    defaultInformationOriginate?: pulumi.Input<string>;
    defaultMetric?: pulumi.Input<number>;
    distances?: pulumi.Input<pulumi.Input<inputs.RouterRipngDistance>[]>;
    distributeLists?: pulumi.Input<pulumi.Input<inputs.RouterRipngDistributeList>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    garbageTimer?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    interfaces?: pulumi.Input<pulumi.Input<inputs.RouterRipngInterface>[]>;
    maxOutMetric?: pulumi.Input<number>;
    neighbors?: pulumi.Input<pulumi.Input<inputs.RouterRipngNeighbor>[]>;
    networks?: pulumi.Input<pulumi.Input<inputs.RouterRipngNetwork>[]>;
    offsetLists?: pulumi.Input<pulumi.Input<inputs.RouterRipngOffsetList>[]>;
    passiveInterfaces?: pulumi.Input<pulumi.Input<inputs.RouterRipngPassiveInterface>[]>;
    redistributes?: pulumi.Input<pulumi.Input<inputs.RouterRipngRedistribute>[]>;
    timeoutTimer?: pulumi.Input<number>;
    updateTimer?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

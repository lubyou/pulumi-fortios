// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SystemSdwan extends pulumi.CustomResource {
    /**
     * Get an existing SystemSdwan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemSdwanState, opts?: pulumi.CustomResourceOptions): SystemSdwan {
        return new SystemSdwan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemSdwan:SystemSdwan';

    /**
     * Returns true if the given object is an instance of SystemSdwan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemSdwan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemSdwan.__pulumiType;
    }

    public readonly duplicationMaxNum!: pulumi.Output<number>;
    public readonly duplications!: pulumi.Output<outputs.SystemSdwanDuplication[] | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly failAlertInterfaces!: pulumi.Output<outputs.SystemSdwanFailAlertInterface[] | undefined>;
    public readonly failDetect!: pulumi.Output<string>;
    public readonly healthChecks!: pulumi.Output<outputs.SystemSdwanHealthCheck[] | undefined>;
    public readonly loadBalanceMode!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<outputs.SystemSdwanMember[] | undefined>;
    public readonly neighborHoldBootTime!: pulumi.Output<number>;
    public readonly neighborHoldDown!: pulumi.Output<string>;
    public readonly neighborHoldDownTime!: pulumi.Output<number>;
    public readonly neighbors!: pulumi.Output<outputs.SystemSdwanNeighbor[] | undefined>;
    public readonly services!: pulumi.Output<outputs.SystemSdwanService[] | undefined>;
    public readonly speedtestBypassRouting!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly zones!: pulumi.Output<outputs.SystemSdwanZone[]>;

    /**
     * Create a SystemSdwan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemSdwanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemSdwanArgs | SystemSdwanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemSdwanState | undefined;
            resourceInputs["duplicationMaxNum"] = state ? state.duplicationMaxNum : undefined;
            resourceInputs["duplications"] = state ? state.duplications : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["failAlertInterfaces"] = state ? state.failAlertInterfaces : undefined;
            resourceInputs["failDetect"] = state ? state.failDetect : undefined;
            resourceInputs["healthChecks"] = state ? state.healthChecks : undefined;
            resourceInputs["loadBalanceMode"] = state ? state.loadBalanceMode : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["neighborHoldBootTime"] = state ? state.neighborHoldBootTime : undefined;
            resourceInputs["neighborHoldDown"] = state ? state.neighborHoldDown : undefined;
            resourceInputs["neighborHoldDownTime"] = state ? state.neighborHoldDownTime : undefined;
            resourceInputs["neighbors"] = state ? state.neighbors : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["speedtestBypassRouting"] = state ? state.speedtestBypassRouting : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as SystemSdwanArgs | undefined;
            resourceInputs["duplicationMaxNum"] = args ? args.duplicationMaxNum : undefined;
            resourceInputs["duplications"] = args ? args.duplications : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["failAlertInterfaces"] = args ? args.failAlertInterfaces : undefined;
            resourceInputs["failDetect"] = args ? args.failDetect : undefined;
            resourceInputs["healthChecks"] = args ? args.healthChecks : undefined;
            resourceInputs["loadBalanceMode"] = args ? args.loadBalanceMode : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["neighborHoldBootTime"] = args ? args.neighborHoldBootTime : undefined;
            resourceInputs["neighborHoldDown"] = args ? args.neighborHoldDown : undefined;
            resourceInputs["neighborHoldDownTime"] = args ? args.neighborHoldDownTime : undefined;
            resourceInputs["neighbors"] = args ? args.neighbors : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["speedtestBypassRouting"] = args ? args.speedtestBypassRouting : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemSdwan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemSdwan resources.
 */
export interface SystemSdwanState {
    duplicationMaxNum?: pulumi.Input<number>;
    duplications?: pulumi.Input<pulumi.Input<inputs.SystemSdwanDuplication>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    failAlertInterfaces?: pulumi.Input<pulumi.Input<inputs.SystemSdwanFailAlertInterface>[]>;
    failDetect?: pulumi.Input<string>;
    healthChecks?: pulumi.Input<pulumi.Input<inputs.SystemSdwanHealthCheck>[]>;
    loadBalanceMode?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<inputs.SystemSdwanMember>[]>;
    neighborHoldBootTime?: pulumi.Input<number>;
    neighborHoldDown?: pulumi.Input<string>;
    neighborHoldDownTime?: pulumi.Input<number>;
    neighbors?: pulumi.Input<pulumi.Input<inputs.SystemSdwanNeighbor>[]>;
    services?: pulumi.Input<pulumi.Input<inputs.SystemSdwanService>[]>;
    speedtestBypassRouting?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    zones?: pulumi.Input<pulumi.Input<inputs.SystemSdwanZone>[]>;
}

/**
 * The set of arguments for constructing a SystemSdwan resource.
 */
export interface SystemSdwanArgs {
    duplicationMaxNum?: pulumi.Input<number>;
    duplications?: pulumi.Input<pulumi.Input<inputs.SystemSdwanDuplication>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    failAlertInterfaces?: pulumi.Input<pulumi.Input<inputs.SystemSdwanFailAlertInterface>[]>;
    failDetect?: pulumi.Input<string>;
    healthChecks?: pulumi.Input<pulumi.Input<inputs.SystemSdwanHealthCheck>[]>;
    loadBalanceMode?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<inputs.SystemSdwanMember>[]>;
    neighborHoldBootTime?: pulumi.Input<number>;
    neighborHoldDown?: pulumi.Input<string>;
    neighborHoldDownTime?: pulumi.Input<number>;
    neighbors?: pulumi.Input<pulumi.Input<inputs.SystemSdwanNeighbor>[]>;
    services?: pulumi.Input<pulumi.Input<inputs.SystemSdwanService>[]>;
    speedtestBypassRouting?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    zones?: pulumi.Input<pulumi.Input<inputs.SystemSdwanZone>[]>;
}

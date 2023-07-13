// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class LogSyslogdFilter extends pulumi.CustomResource {
    /**
     * Get an existing LogSyslogdFilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogSyslogdFilterState, opts?: pulumi.CustomResourceOptions): LogSyslogdFilter {
        return new LogSyslogdFilter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logSyslogdFilter:LogSyslogdFilter';

    /**
     * Returns true if the given object is an instance of LogSyslogdFilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogSyslogdFilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogSyslogdFilter.__pulumiType;
    }

    public readonly anomaly!: pulumi.Output<string>;
    public readonly dns!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly filter!: pulumi.Output<string>;
    public readonly filterType!: pulumi.Output<string>;
    public readonly forwardTraffic!: pulumi.Output<string>;
    public readonly freeStyles!: pulumi.Output<outputs.LogSyslogdFilterFreeStyle[] | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly gtp!: pulumi.Output<string>;
    public readonly localTraffic!: pulumi.Output<string>;
    public readonly multicastTraffic!: pulumi.Output<string>;
    public readonly netscanDiscovery!: pulumi.Output<string>;
    public readonly netscanVulnerability!: pulumi.Output<string>;
    public readonly severity!: pulumi.Output<string>;
    public readonly snifferTraffic!: pulumi.Output<string>;
    public readonly ssh!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly voip!: pulumi.Output<string>;
    public readonly ztnaTraffic!: pulumi.Output<string>;

    /**
     * Create a LogSyslogdFilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LogSyslogdFilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogSyslogdFilterArgs | LogSyslogdFilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogSyslogdFilterState | undefined;
            resourceInputs["anomaly"] = state ? state.anomaly : undefined;
            resourceInputs["dns"] = state ? state.dns : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["filterType"] = state ? state.filterType : undefined;
            resourceInputs["forwardTraffic"] = state ? state.forwardTraffic : undefined;
            resourceInputs["freeStyles"] = state ? state.freeStyles : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["gtp"] = state ? state.gtp : undefined;
            resourceInputs["localTraffic"] = state ? state.localTraffic : undefined;
            resourceInputs["multicastTraffic"] = state ? state.multicastTraffic : undefined;
            resourceInputs["netscanDiscovery"] = state ? state.netscanDiscovery : undefined;
            resourceInputs["netscanVulnerability"] = state ? state.netscanVulnerability : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["snifferTraffic"] = state ? state.snifferTraffic : undefined;
            resourceInputs["ssh"] = state ? state.ssh : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["voip"] = state ? state.voip : undefined;
            resourceInputs["ztnaTraffic"] = state ? state.ztnaTraffic : undefined;
        } else {
            const args = argsOrState as LogSyslogdFilterArgs | undefined;
            resourceInputs["anomaly"] = args ? args.anomaly : undefined;
            resourceInputs["dns"] = args ? args.dns : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["filterType"] = args ? args.filterType : undefined;
            resourceInputs["forwardTraffic"] = args ? args.forwardTraffic : undefined;
            resourceInputs["freeStyles"] = args ? args.freeStyles : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["gtp"] = args ? args.gtp : undefined;
            resourceInputs["localTraffic"] = args ? args.localTraffic : undefined;
            resourceInputs["multicastTraffic"] = args ? args.multicastTraffic : undefined;
            resourceInputs["netscanDiscovery"] = args ? args.netscanDiscovery : undefined;
            resourceInputs["netscanVulnerability"] = args ? args.netscanVulnerability : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["snifferTraffic"] = args ? args.snifferTraffic : undefined;
            resourceInputs["ssh"] = args ? args.ssh : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["voip"] = args ? args.voip : undefined;
            resourceInputs["ztnaTraffic"] = args ? args.ztnaTraffic : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogSyslogdFilter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogSyslogdFilter resources.
 */
export interface LogSyslogdFilterState {
    anomaly?: pulumi.Input<string>;
    dns?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    filter?: pulumi.Input<string>;
    filterType?: pulumi.Input<string>;
    forwardTraffic?: pulumi.Input<string>;
    freeStyles?: pulumi.Input<pulumi.Input<inputs.LogSyslogdFilterFreeStyle>[]>;
    getAllTables?: pulumi.Input<string>;
    gtp?: pulumi.Input<string>;
    localTraffic?: pulumi.Input<string>;
    multicastTraffic?: pulumi.Input<string>;
    netscanDiscovery?: pulumi.Input<string>;
    netscanVulnerability?: pulumi.Input<string>;
    severity?: pulumi.Input<string>;
    snifferTraffic?: pulumi.Input<string>;
    ssh?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    voip?: pulumi.Input<string>;
    ztnaTraffic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogSyslogdFilter resource.
 */
export interface LogSyslogdFilterArgs {
    anomaly?: pulumi.Input<string>;
    dns?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    filter?: pulumi.Input<string>;
    filterType?: pulumi.Input<string>;
    forwardTraffic?: pulumi.Input<string>;
    freeStyles?: pulumi.Input<pulumi.Input<inputs.LogSyslogdFilterFreeStyle>[]>;
    getAllTables?: pulumi.Input<string>;
    gtp?: pulumi.Input<string>;
    localTraffic?: pulumi.Input<string>;
    multicastTraffic?: pulumi.Input<string>;
    netscanDiscovery?: pulumi.Input<string>;
    netscanVulnerability?: pulumi.Input<string>;
    severity?: pulumi.Input<string>;
    snifferTraffic?: pulumi.Input<string>;
    ssh?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    voip?: pulumi.Input<string>;
    ztnaTraffic?: pulumi.Input<string>;
}

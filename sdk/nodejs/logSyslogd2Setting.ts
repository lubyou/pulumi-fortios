// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class LogSyslogd2Setting extends pulumi.CustomResource {
    /**
     * Get an existing LogSyslogd2Setting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogSyslogd2SettingState, opts?: pulumi.CustomResourceOptions): LogSyslogd2Setting {
        return new LogSyslogd2Setting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logSyslogd2Setting:LogSyslogd2Setting';

    /**
     * Returns true if the given object is an instance of LogSyslogd2Setting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogSyslogd2Setting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogSyslogd2Setting.__pulumiType;
    }

    public readonly certificate!: pulumi.Output<string>;
    public readonly customFieldNames!: pulumi.Output<outputs.LogSyslogd2SettingCustomFieldName[] | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly encAlgorithm!: pulumi.Output<string>;
    public readonly facility!: pulumi.Output<string>;
    public readonly format!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly maxLogRate!: pulumi.Output<number>;
    public readonly mode!: pulumi.Output<string>;
    public readonly port!: pulumi.Output<number>;
    public readonly priority!: pulumi.Output<string>;
    public readonly server!: pulumi.Output<string>;
    public readonly sourceIp!: pulumi.Output<string>;
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly syslogType!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a LogSyslogd2Setting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LogSyslogd2SettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogSyslogd2SettingArgs | LogSyslogd2SettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogSyslogd2SettingState | undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["customFieldNames"] = state ? state.customFieldNames : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["encAlgorithm"] = state ? state.encAlgorithm : undefined;
            resourceInputs["facility"] = state ? state.facility : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["maxLogRate"] = state ? state.maxLogRate : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["syslogType"] = state ? state.syslogType : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LogSyslogd2SettingArgs | undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["customFieldNames"] = args ? args.customFieldNames : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["encAlgorithm"] = args ? args.encAlgorithm : undefined;
            resourceInputs["facility"] = args ? args.facility : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["maxLogRate"] = args ? args.maxLogRate : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["syslogType"] = args ? args.syslogType : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogSyslogd2Setting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogSyslogd2Setting resources.
 */
export interface LogSyslogd2SettingState {
    certificate?: pulumi.Input<string>;
    customFieldNames?: pulumi.Input<pulumi.Input<inputs.LogSyslogd2SettingCustomFieldName>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    encAlgorithm?: pulumi.Input<string>;
    facility?: pulumi.Input<string>;
    format?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    maxLogRate?: pulumi.Input<number>;
    mode?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    priority?: pulumi.Input<string>;
    server?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sslMinProtoVersion?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    syslogType?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogSyslogd2Setting resource.
 */
export interface LogSyslogd2SettingArgs {
    certificate?: pulumi.Input<string>;
    customFieldNames?: pulumi.Input<pulumi.Input<inputs.LogSyslogd2SettingCustomFieldName>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    encAlgorithm?: pulumi.Input<string>;
    facility?: pulumi.Input<string>;
    format?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    maxLogRate?: pulumi.Input<number>;
    mode?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    priority?: pulumi.Input<string>;
    server?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sslMinProtoVersion?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    syslogType?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

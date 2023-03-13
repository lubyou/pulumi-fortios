// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class LogFortianalyzerOverrideSetting extends pulumi.CustomResource {
    /**
     * Get an existing LogFortianalyzerOverrideSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogFortianalyzerOverrideSettingState, opts?: pulumi.CustomResourceOptions): LogFortianalyzerOverrideSetting {
        return new LogFortianalyzerOverrideSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logFortianalyzerOverrideSetting:LogFortianalyzerOverrideSetting';

    /**
     * Returns true if the given object is an instance of LogFortianalyzerOverrideSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogFortianalyzerOverrideSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogFortianalyzerOverrideSetting.__pulumiType;
    }

    public readonly __changeIp!: pulumi.Output<number>;
    public readonly accessConfig!: pulumi.Output<string>;
    public readonly certificate!: pulumi.Output<string>;
    public readonly certificateVerification!: pulumi.Output<string>;
    public readonly connTimeout!: pulumi.Output<number>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly encAlgorithm!: pulumi.Output<string>;
    public readonly fazType!: pulumi.Output<number>;
    public readonly hmacAlgorithm!: pulumi.Output<string>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly ipsArchive!: pulumi.Output<string>;
    public readonly maxLogRate!: pulumi.Output<number>;
    public readonly mgmtName!: pulumi.Output<string>;
    public readonly monitorFailureRetryPeriod!: pulumi.Output<number>;
    public readonly monitorKeepalivePeriod!: pulumi.Output<number>;
    public readonly override!: pulumi.Output<string>;
    public readonly presharedKey!: pulumi.Output<string>;
    public readonly priority!: pulumi.Output<string>;
    public readonly reliable!: pulumi.Output<string>;
    public readonly serials!: pulumi.Output<outputs.LogFortianalyzerOverrideSettingSerial[] | undefined>;
    public readonly server!: pulumi.Output<string>;
    public readonly sourceIp!: pulumi.Output<string>;
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly uploadDay!: pulumi.Output<string>;
    public readonly uploadInterval!: pulumi.Output<string>;
    public readonly uploadOption!: pulumi.Output<string>;
    public readonly uploadTime!: pulumi.Output<string>;
    public readonly useManagementVdom!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a LogFortianalyzerOverrideSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LogFortianalyzerOverrideSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogFortianalyzerOverrideSettingArgs | LogFortianalyzerOverrideSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogFortianalyzerOverrideSettingState | undefined;
            resourceInputs["__changeIp"] = state ? state.__changeIp : undefined;
            resourceInputs["accessConfig"] = state ? state.accessConfig : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["certificateVerification"] = state ? state.certificateVerification : undefined;
            resourceInputs["connTimeout"] = state ? state.connTimeout : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["encAlgorithm"] = state ? state.encAlgorithm : undefined;
            resourceInputs["fazType"] = state ? state.fazType : undefined;
            resourceInputs["hmacAlgorithm"] = state ? state.hmacAlgorithm : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["ipsArchive"] = state ? state.ipsArchive : undefined;
            resourceInputs["maxLogRate"] = state ? state.maxLogRate : undefined;
            resourceInputs["mgmtName"] = state ? state.mgmtName : undefined;
            resourceInputs["monitorFailureRetryPeriod"] = state ? state.monitorFailureRetryPeriod : undefined;
            resourceInputs["monitorKeepalivePeriod"] = state ? state.monitorKeepalivePeriod : undefined;
            resourceInputs["override"] = state ? state.override : undefined;
            resourceInputs["presharedKey"] = state ? state.presharedKey : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["reliable"] = state ? state.reliable : undefined;
            resourceInputs["serials"] = state ? state.serials : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["uploadDay"] = state ? state.uploadDay : undefined;
            resourceInputs["uploadInterval"] = state ? state.uploadInterval : undefined;
            resourceInputs["uploadOption"] = state ? state.uploadOption : undefined;
            resourceInputs["uploadTime"] = state ? state.uploadTime : undefined;
            resourceInputs["useManagementVdom"] = state ? state.useManagementVdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LogFortianalyzerOverrideSettingArgs | undefined;
            resourceInputs["__changeIp"] = args ? args.__changeIp : undefined;
            resourceInputs["accessConfig"] = args ? args.accessConfig : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["certificateVerification"] = args ? args.certificateVerification : undefined;
            resourceInputs["connTimeout"] = args ? args.connTimeout : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["encAlgorithm"] = args ? args.encAlgorithm : undefined;
            resourceInputs["fazType"] = args ? args.fazType : undefined;
            resourceInputs["hmacAlgorithm"] = args ? args.hmacAlgorithm : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["ipsArchive"] = args ? args.ipsArchive : undefined;
            resourceInputs["maxLogRate"] = args ? args.maxLogRate : undefined;
            resourceInputs["mgmtName"] = args ? args.mgmtName : undefined;
            resourceInputs["monitorFailureRetryPeriod"] = args ? args.monitorFailureRetryPeriod : undefined;
            resourceInputs["monitorKeepalivePeriod"] = args ? args.monitorKeepalivePeriod : undefined;
            resourceInputs["override"] = args ? args.override : undefined;
            resourceInputs["presharedKey"] = args ? args.presharedKey : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["reliable"] = args ? args.reliable : undefined;
            resourceInputs["serials"] = args ? args.serials : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["uploadDay"] = args ? args.uploadDay : undefined;
            resourceInputs["uploadInterval"] = args ? args.uploadInterval : undefined;
            resourceInputs["uploadOption"] = args ? args.uploadOption : undefined;
            resourceInputs["uploadTime"] = args ? args.uploadTime : undefined;
            resourceInputs["useManagementVdom"] = args ? args.useManagementVdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogFortianalyzerOverrideSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogFortianalyzerOverrideSetting resources.
 */
export interface LogFortianalyzerOverrideSettingState {
    __changeIp?: pulumi.Input<number>;
    accessConfig?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    certificateVerification?: pulumi.Input<string>;
    connTimeout?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    encAlgorithm?: pulumi.Input<string>;
    fazType?: pulumi.Input<number>;
    hmacAlgorithm?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    ipsArchive?: pulumi.Input<string>;
    maxLogRate?: pulumi.Input<number>;
    mgmtName?: pulumi.Input<string>;
    monitorFailureRetryPeriod?: pulumi.Input<number>;
    monitorKeepalivePeriod?: pulumi.Input<number>;
    override?: pulumi.Input<string>;
    presharedKey?: pulumi.Input<string>;
    priority?: pulumi.Input<string>;
    reliable?: pulumi.Input<string>;
    serials?: pulumi.Input<pulumi.Input<inputs.LogFortianalyzerOverrideSettingSerial>[]>;
    server?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sslMinProtoVersion?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    uploadDay?: pulumi.Input<string>;
    uploadInterval?: pulumi.Input<string>;
    uploadOption?: pulumi.Input<string>;
    uploadTime?: pulumi.Input<string>;
    useManagementVdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogFortianalyzerOverrideSetting resource.
 */
export interface LogFortianalyzerOverrideSettingArgs {
    __changeIp?: pulumi.Input<number>;
    accessConfig?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    certificateVerification?: pulumi.Input<string>;
    connTimeout?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    encAlgorithm?: pulumi.Input<string>;
    fazType?: pulumi.Input<number>;
    hmacAlgorithm?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    ipsArchive?: pulumi.Input<string>;
    maxLogRate?: pulumi.Input<number>;
    mgmtName?: pulumi.Input<string>;
    monitorFailureRetryPeriod?: pulumi.Input<number>;
    monitorKeepalivePeriod?: pulumi.Input<number>;
    override?: pulumi.Input<string>;
    presharedKey?: pulumi.Input<string>;
    priority?: pulumi.Input<string>;
    reliable?: pulumi.Input<string>;
    serials?: pulumi.Input<pulumi.Input<inputs.LogFortianalyzerOverrideSettingSerial>[]>;
    server?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sslMinProtoVersion?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    uploadDay?: pulumi.Input<string>;
    uploadInterval?: pulumi.Input<string>;
    uploadOption?: pulumi.Input<string>;
    uploadTime?: pulumi.Input<string>;
    useManagementVdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

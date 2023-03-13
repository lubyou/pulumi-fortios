// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LogFortiguardSetting extends pulumi.CustomResource {
    /**
     * Get an existing LogFortiguardSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogFortiguardSettingState, opts?: pulumi.CustomResourceOptions): LogFortiguardSetting {
        return new LogFortiguardSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logFortiguardSetting:LogFortiguardSetting';

    /**
     * Returns true if the given object is an instance of LogFortiguardSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogFortiguardSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogFortiguardSetting.__pulumiType;
    }

    public readonly accessConfig!: pulumi.Output<string>;
    public readonly connTimeout!: pulumi.Output<number>;
    public readonly encAlgorithm!: pulumi.Output<string>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly maxLogRate!: pulumi.Output<number>;
    public readonly priority!: pulumi.Output<string>;
    public readonly sourceIp!: pulumi.Output<string>;
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly uploadDay!: pulumi.Output<string>;
    public readonly uploadInterval!: pulumi.Output<string>;
    public readonly uploadOption!: pulumi.Output<string>;
    public readonly uploadTime!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a LogFortiguardSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LogFortiguardSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogFortiguardSettingArgs | LogFortiguardSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogFortiguardSettingState | undefined;
            resourceInputs["accessConfig"] = state ? state.accessConfig : undefined;
            resourceInputs["connTimeout"] = state ? state.connTimeout : undefined;
            resourceInputs["encAlgorithm"] = state ? state.encAlgorithm : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["maxLogRate"] = state ? state.maxLogRate : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["uploadDay"] = state ? state.uploadDay : undefined;
            resourceInputs["uploadInterval"] = state ? state.uploadInterval : undefined;
            resourceInputs["uploadOption"] = state ? state.uploadOption : undefined;
            resourceInputs["uploadTime"] = state ? state.uploadTime : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LogFortiguardSettingArgs | undefined;
            resourceInputs["accessConfig"] = args ? args.accessConfig : undefined;
            resourceInputs["connTimeout"] = args ? args.connTimeout : undefined;
            resourceInputs["encAlgorithm"] = args ? args.encAlgorithm : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["maxLogRate"] = args ? args.maxLogRate : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["uploadDay"] = args ? args.uploadDay : undefined;
            resourceInputs["uploadInterval"] = args ? args.uploadInterval : undefined;
            resourceInputs["uploadOption"] = args ? args.uploadOption : undefined;
            resourceInputs["uploadTime"] = args ? args.uploadTime : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogFortiguardSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogFortiguardSetting resources.
 */
export interface LogFortiguardSettingState {
    accessConfig?: pulumi.Input<string>;
    connTimeout?: pulumi.Input<number>;
    encAlgorithm?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    maxLogRate?: pulumi.Input<number>;
    priority?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sslMinProtoVersion?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    uploadDay?: pulumi.Input<string>;
    uploadInterval?: pulumi.Input<string>;
    uploadOption?: pulumi.Input<string>;
    uploadTime?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogFortiguardSetting resource.
 */
export interface LogFortiguardSettingArgs {
    accessConfig?: pulumi.Input<string>;
    connTimeout?: pulumi.Input<number>;
    encAlgorithm?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    maxLogRate?: pulumi.Input<number>;
    priority?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sslMinProtoVersion?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    uploadDay?: pulumi.Input<string>;
    uploadInterval?: pulumi.Input<string>;
    uploadOption?: pulumi.Input<string>;
    uploadTime?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

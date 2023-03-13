// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LogTacacsAccounting2Filter extends pulumi.CustomResource {
    /**
     * Get an existing LogTacacsAccounting2Filter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogTacacsAccounting2FilterState, opts?: pulumi.CustomResourceOptions): LogTacacsAccounting2Filter {
        return new LogTacacsAccounting2Filter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logTacacsAccounting2Filter:LogTacacsAccounting2Filter';

    /**
     * Returns true if the given object is an instance of LogTacacsAccounting2Filter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogTacacsAccounting2Filter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogTacacsAccounting2Filter.__pulumiType;
    }

    public readonly cliCmdAudit!: pulumi.Output<string>;
    public readonly configChangeAudit!: pulumi.Output<string>;
    public readonly loginAudit!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a LogTacacsAccounting2Filter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LogTacacsAccounting2FilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogTacacsAccounting2FilterArgs | LogTacacsAccounting2FilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogTacacsAccounting2FilterState | undefined;
            resourceInputs["cliCmdAudit"] = state ? state.cliCmdAudit : undefined;
            resourceInputs["configChangeAudit"] = state ? state.configChangeAudit : undefined;
            resourceInputs["loginAudit"] = state ? state.loginAudit : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LogTacacsAccounting2FilterArgs | undefined;
            resourceInputs["cliCmdAudit"] = args ? args.cliCmdAudit : undefined;
            resourceInputs["configChangeAudit"] = args ? args.configChangeAudit : undefined;
            resourceInputs["loginAudit"] = args ? args.loginAudit : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogTacacsAccounting2Filter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogTacacsAccounting2Filter resources.
 */
export interface LogTacacsAccounting2FilterState {
    cliCmdAudit?: pulumi.Input<string>;
    configChangeAudit?: pulumi.Input<string>;
    loginAudit?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogTacacsAccounting2Filter resource.
 */
export interface LogTacacsAccounting2FilterArgs {
    cliCmdAudit?: pulumi.Input<string>;
    configChangeAudit?: pulumi.Input<string>;
    loginAudit?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

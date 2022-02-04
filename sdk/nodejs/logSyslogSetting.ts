// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to configure logging to remote Syslog logging servers.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.LogSyslogdSetting`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const test2 = new fortios.LogSyslogSetting("test2", {
 *     facility: "local7",
 *     format: "csv",
 *     mode: "udp",
 *     port: "514",
 *     server: "2.2.2.2",
 *     sourceIp: "10.2.2.199",
 *     status: "enable",
 * });
 * ```
 */
export class LogSyslogSetting extends pulumi.CustomResource {
    /**
     * Get an existing LogSyslogSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogSyslogSettingState, opts?: pulumi.CustomResourceOptions): LogSyslogSetting {
        return new LogSyslogSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logSyslogSetting:LogSyslogSetting';

    /**
     * Returns true if the given object is an instance of LogSyslogSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogSyslogSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogSyslogSetting.__pulumiType;
    }

    /**
     * Remote syslog facility.
     */
    public readonly facility!: pulumi.Output<string>;
    /**
     * Log format.
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * Remote syslog logging over UDP/Reliable TCP.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Server listen port.
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * Address of remote syslog server.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Source IP address of syslog.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Enable/disable remote syslog logging.
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a LogSyslogSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogSyslogSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogSyslogSettingArgs | LogSyslogSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogSyslogSettingState | undefined;
            resourceInputs["facility"] = state ? state.facility : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as LogSyslogSettingArgs | undefined;
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["facility"] = args ? args.facility : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogSyslogSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogSyslogSetting resources.
 */
export interface LogSyslogSettingState {
    /**
     * Remote syslog facility.
     */
    facility?: pulumi.Input<string>;
    /**
     * Log format.
     */
    format?: pulumi.Input<string>;
    /**
     * Remote syslog logging over UDP/Reliable TCP.
     */
    mode?: pulumi.Input<string>;
    /**
     * Server listen port.
     */
    port?: pulumi.Input<string>;
    /**
     * Address of remote syslog server.
     */
    server?: pulumi.Input<string>;
    /**
     * Source IP address of syslog.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable remote syslog logging.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogSyslogSetting resource.
 */
export interface LogSyslogSettingArgs {
    /**
     * Remote syslog facility.
     */
    facility?: pulumi.Input<string>;
    /**
     * Log format.
     */
    format?: pulumi.Input<string>;
    /**
     * Remote syslog logging over UDP/Reliable TCP.
     */
    mode?: pulumi.Input<string>;
    /**
     * Server listen port.
     */
    port?: pulumi.Input<string>;
    /**
     * Address of remote syslog server.
     */
    server?: pulumi.Input<string>;
    /**
     * Source IP address of syslog.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable remote syslog logging.
     */
    status: pulumi.Input<string>;
}

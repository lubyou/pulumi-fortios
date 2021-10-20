// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure WCCP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemWccp("trname", {
 *     assignmentBucketFormat: "cisco-implementation",
 *     assignmentDstaddrMask: "0.0.0.0",
 *     assignmentMethod: "HASH",
 *     assignmentSrcaddrMask: "0.0.23.65",
 *     assignmentWeight: 0,
 *     authentication: "disable",
 *     cacheEngineMethod: "GRE",
 *     cacheId: "1.1.1.1",
 *     forwardMethod: "GRE",
 *     groupAddress: "0.0.0.0",
 *     primaryHash: "dst-ip",
 *     priority: 0,
 *     protocol: 0,
 *     returnMethod: "GRE",
 *     routerId: "1.1.1.1",
 *     routerList: "\"1.0.0.0\" ",
 *     serverType: "forward",
 *     serviceId: "1",
 *     serviceType: "auto",
 * });
 * ```
 *
 * ## Import
 *
 * System Wccp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemWccp:SystemWccp labelname {{service_id}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemWccp extends pulumi.CustomResource {
    /**
     * Get an existing SystemWccp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemWccpState, opts?: pulumi.CustomResourceOptions): SystemWccp {
        return new SystemWccp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemWccp:SystemWccp';

    /**
     * Returns true if the given object is an instance of SystemWccp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemWccp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemWccp.__pulumiType;
    }

    /**
     * Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
     */
    public readonly assignmentBucketFormat!: pulumi.Output<string>;
    /**
     * Assignment destination address mask.
     */
    public readonly assignmentDstaddrMask!: pulumi.Output<string>;
    /**
     * Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
     */
    public readonly assignmentMethod!: pulumi.Output<string>;
    /**
     * Assignment source address mask.
     */
    public readonly assignmentSrcaddrMask!: pulumi.Output<string>;
    /**
     * Assignment of hash weight/ratio for the WCCP cache engine.
     */
    public readonly assignmentWeight!: pulumi.Output<number>;
    /**
     * Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
     */
    public readonly authentication!: pulumi.Output<string>;
    /**
     * Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
     */
    public readonly cacheEngineMethod!: pulumi.Output<string>;
    /**
     * IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
     */
    public readonly cacheId!: pulumi.Output<string>;
    /**
     * Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
     */
    public readonly forwardMethod!: pulumi.Output<string>;
    /**
     * IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
     */
    public readonly groupAddress!: pulumi.Output<string>;
    /**
     * Password for MD5 authentication.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Service ports.
     */
    public readonly ports!: pulumi.Output<string>;
    /**
     * Match method. Valid values: `source`, `destination`.
     */
    public readonly portsDefined!: pulumi.Output<string>;
    /**
     * Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
     */
    public readonly primaryHash!: pulumi.Output<string>;
    /**
     * Service priority.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Service protocol.
     */
    public readonly protocol!: pulumi.Output<number>;
    /**
     * Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
     */
    public readonly returnMethod!: pulumi.Output<string>;
    /**
     * IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
     */
    public readonly routerId!: pulumi.Output<string>;
    /**
     * IP addresses of one or more WCCP routers.
     */
    public readonly routerList!: pulumi.Output<string>;
    /**
     * IP addresses and netmasks for up to four cache servers.
     */
    public readonly serverList!: pulumi.Output<string>;
    /**
     * Cache server type. Valid values: `forward`, `proxy`.
     */
    public readonly serverType!: pulumi.Output<string>;
    /**
     * Service ID.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
     */
    public readonly serviceType!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemWccp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemWccpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemWccpArgs | SystemWccpState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemWccpState | undefined;
            inputs["assignmentBucketFormat"] = state ? state.assignmentBucketFormat : undefined;
            inputs["assignmentDstaddrMask"] = state ? state.assignmentDstaddrMask : undefined;
            inputs["assignmentMethod"] = state ? state.assignmentMethod : undefined;
            inputs["assignmentSrcaddrMask"] = state ? state.assignmentSrcaddrMask : undefined;
            inputs["assignmentWeight"] = state ? state.assignmentWeight : undefined;
            inputs["authentication"] = state ? state.authentication : undefined;
            inputs["cacheEngineMethod"] = state ? state.cacheEngineMethod : undefined;
            inputs["cacheId"] = state ? state.cacheId : undefined;
            inputs["forwardMethod"] = state ? state.forwardMethod : undefined;
            inputs["groupAddress"] = state ? state.groupAddress : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["ports"] = state ? state.ports : undefined;
            inputs["portsDefined"] = state ? state.portsDefined : undefined;
            inputs["primaryHash"] = state ? state.primaryHash : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["returnMethod"] = state ? state.returnMethod : undefined;
            inputs["routerId"] = state ? state.routerId : undefined;
            inputs["routerList"] = state ? state.routerList : undefined;
            inputs["serverList"] = state ? state.serverList : undefined;
            inputs["serverType"] = state ? state.serverType : undefined;
            inputs["serviceId"] = state ? state.serviceId : undefined;
            inputs["serviceType"] = state ? state.serviceType : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemWccpArgs | undefined;
            inputs["assignmentBucketFormat"] = args ? args.assignmentBucketFormat : undefined;
            inputs["assignmentDstaddrMask"] = args ? args.assignmentDstaddrMask : undefined;
            inputs["assignmentMethod"] = args ? args.assignmentMethod : undefined;
            inputs["assignmentSrcaddrMask"] = args ? args.assignmentSrcaddrMask : undefined;
            inputs["assignmentWeight"] = args ? args.assignmentWeight : undefined;
            inputs["authentication"] = args ? args.authentication : undefined;
            inputs["cacheEngineMethod"] = args ? args.cacheEngineMethod : undefined;
            inputs["cacheId"] = args ? args.cacheId : undefined;
            inputs["forwardMethod"] = args ? args.forwardMethod : undefined;
            inputs["groupAddress"] = args ? args.groupAddress : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["ports"] = args ? args.ports : undefined;
            inputs["portsDefined"] = args ? args.portsDefined : undefined;
            inputs["primaryHash"] = args ? args.primaryHash : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["returnMethod"] = args ? args.returnMethod : undefined;
            inputs["routerId"] = args ? args.routerId : undefined;
            inputs["routerList"] = args ? args.routerList : undefined;
            inputs["serverList"] = args ? args.serverList : undefined;
            inputs["serverType"] = args ? args.serverType : undefined;
            inputs["serviceId"] = args ? args.serviceId : undefined;
            inputs["serviceType"] = args ? args.serviceType : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemWccp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemWccp resources.
 */
export interface SystemWccpState {
    /**
     * Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
     */
    assignmentBucketFormat?: pulumi.Input<string>;
    /**
     * Assignment destination address mask.
     */
    assignmentDstaddrMask?: pulumi.Input<string>;
    /**
     * Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
     */
    assignmentMethod?: pulumi.Input<string>;
    /**
     * Assignment source address mask.
     */
    assignmentSrcaddrMask?: pulumi.Input<string>;
    /**
     * Assignment of hash weight/ratio for the WCCP cache engine.
     */
    assignmentWeight?: pulumi.Input<number>;
    /**
     * Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
     */
    cacheEngineMethod?: pulumi.Input<string>;
    /**
     * IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
     */
    cacheId?: pulumi.Input<string>;
    /**
     * Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
     */
    forwardMethod?: pulumi.Input<string>;
    /**
     * IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
     */
    groupAddress?: pulumi.Input<string>;
    /**
     * Password for MD5 authentication.
     */
    password?: pulumi.Input<string>;
    /**
     * Service ports.
     */
    ports?: pulumi.Input<string>;
    /**
     * Match method. Valid values: `source`, `destination`.
     */
    portsDefined?: pulumi.Input<string>;
    /**
     * Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
     */
    primaryHash?: pulumi.Input<string>;
    /**
     * Service priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * Service protocol.
     */
    protocol?: pulumi.Input<number>;
    /**
     * Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
     */
    returnMethod?: pulumi.Input<string>;
    /**
     * IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
     */
    routerId?: pulumi.Input<string>;
    /**
     * IP addresses of one or more WCCP routers.
     */
    routerList?: pulumi.Input<string>;
    /**
     * IP addresses and netmasks for up to four cache servers.
     */
    serverList?: pulumi.Input<string>;
    /**
     * Cache server type. Valid values: `forward`, `proxy`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Service ID.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
     */
    serviceType?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemWccp resource.
 */
export interface SystemWccpArgs {
    /**
     * Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
     */
    assignmentBucketFormat?: pulumi.Input<string>;
    /**
     * Assignment destination address mask.
     */
    assignmentDstaddrMask?: pulumi.Input<string>;
    /**
     * Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
     */
    assignmentMethod?: pulumi.Input<string>;
    /**
     * Assignment source address mask.
     */
    assignmentSrcaddrMask?: pulumi.Input<string>;
    /**
     * Assignment of hash weight/ratio for the WCCP cache engine.
     */
    assignmentWeight?: pulumi.Input<number>;
    /**
     * Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
     */
    cacheEngineMethod?: pulumi.Input<string>;
    /**
     * IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
     */
    cacheId?: pulumi.Input<string>;
    /**
     * Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
     */
    forwardMethod?: pulumi.Input<string>;
    /**
     * IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
     */
    groupAddress?: pulumi.Input<string>;
    /**
     * Password for MD5 authentication.
     */
    password?: pulumi.Input<string>;
    /**
     * Service ports.
     */
    ports?: pulumi.Input<string>;
    /**
     * Match method. Valid values: `source`, `destination`.
     */
    portsDefined?: pulumi.Input<string>;
    /**
     * Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
     */
    primaryHash?: pulumi.Input<string>;
    /**
     * Service priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * Service protocol.
     */
    protocol?: pulumi.Input<number>;
    /**
     * Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
     */
    returnMethod?: pulumi.Input<string>;
    /**
     * IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
     */
    routerId?: pulumi.Input<string>;
    /**
     * IP addresses of one or more WCCP routers.
     */
    routerList?: pulumi.Input<string>;
    /**
     * IP addresses and netmasks for up to four cache servers.
     */
    serverList?: pulumi.Input<string>;
    /**
     * Cache server type. Valid values: `forward`, `proxy`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Service ID.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
     */
    serviceType?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
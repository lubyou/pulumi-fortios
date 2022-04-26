// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure logical storage.
 *
 * ## Import
 *
 * System Storage can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemStorage:SystemStorage labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemStorage:SystemStorage labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemStorage extends pulumi.CustomResource {
    /**
     * Get an existing SystemStorage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemStorageState, opts?: pulumi.CustomResourceOptions): SystemStorage {
        return new SystemStorage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemStorage:SystemStorage';

    /**
     * Returns true if the given object is an instance of SystemStorage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemStorage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemStorage.__pulumiType;
    }

    /**
     * Partition device.
     */
    public readonly device!: pulumi.Output<string>;
    /**
     * The physical status of current media. Valid values: `enable`, `disable`, `fail`.
     */
    public readonly mediaStatus!: pulumi.Output<string>;
    /**
     * Storage name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Set storage order.
     */
    public readonly order!: pulumi.Output<number>;
    /**
     * Label of underlying partition.
     */
    public readonly partition!: pulumi.Output<string>;
    /**
     * Partition size.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * Enable/disable storage. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Use hard disk for logging or WAN Optimization (default = log).
     */
    public readonly usage!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
     */
    public readonly wanoptMode!: pulumi.Output<string>;

    /**
     * Create a SystemStorage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemStorageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemStorageArgs | SystemStorageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemStorageState | undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["mediaStatus"] = state ? state.mediaStatus : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["order"] = state ? state.order : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["usage"] = state ? state.usage : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wanoptMode"] = state ? state.wanoptMode : undefined;
        } else {
            const args = argsOrState as SystemStorageArgs | undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["mediaStatus"] = args ? args.mediaStatus : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["order"] = args ? args.order : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["usage"] = args ? args.usage : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wanoptMode"] = args ? args.wanoptMode : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemStorage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemStorage resources.
 */
export interface SystemStorageState {
    /**
     * Partition device.
     */
    device?: pulumi.Input<string>;
    /**
     * The physical status of current media. Valid values: `enable`, `disable`, `fail`.
     */
    mediaStatus?: pulumi.Input<string>;
    /**
     * Storage name.
     */
    name?: pulumi.Input<string>;
    /**
     * Set storage order.
     */
    order?: pulumi.Input<number>;
    /**
     * Label of underlying partition.
     */
    partition?: pulumi.Input<string>;
    /**
     * Partition size.
     */
    size?: pulumi.Input<number>;
    /**
     * Enable/disable storage. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Use hard disk for logging or WAN Optimization (default = log).
     */
    usage?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
     */
    wanoptMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemStorage resource.
 */
export interface SystemStorageArgs {
    /**
     * Partition device.
     */
    device?: pulumi.Input<string>;
    /**
     * The physical status of current media. Valid values: `enable`, `disable`, `fail`.
     */
    mediaStatus?: pulumi.Input<string>;
    /**
     * Storage name.
     */
    name?: pulumi.Input<string>;
    /**
     * Set storage order.
     */
    order?: pulumi.Input<number>;
    /**
     * Label of underlying partition.
     */
    partition?: pulumi.Input<string>;
    /**
     * Partition size.
     */
    size?: pulumi.Input<number>;
    /**
     * Enable/disable storage. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Use hard disk for logging or WAN Optimization (default = log).
     */
    usage?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
     */
    wanoptMode?: pulumi.Input<string>;
}

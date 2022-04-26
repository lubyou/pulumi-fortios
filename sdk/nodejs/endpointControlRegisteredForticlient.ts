// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Registered FortiClient list. Applies to FortiOS Version `<= 6.2.0`.
 *
 * ## Import
 *
 * EndpointControl RegisteredForticlient can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/endpointControlRegisteredForticlient:EndpointControlRegisteredForticlient labelname {{uid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/endpointControlRegisteredForticlient:EndpointControlRegisteredForticlient labelname {{uid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class EndpointControlRegisteredForticlient extends pulumi.CustomResource {
    /**
     * Get an existing EndpointControlRegisteredForticlient resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointControlRegisteredForticlientState, opts?: pulumi.CustomResourceOptions): EndpointControlRegisteredForticlient {
        return new EndpointControlRegisteredForticlient(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/endpointControlRegisteredForticlient:EndpointControlRegisteredForticlient';

    /**
     * Returns true if the given object is an instance of EndpointControlRegisteredForticlient.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointControlRegisteredForticlient {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointControlRegisteredForticlient.__pulumiType;
    }

    /**
     * FortiClient registration flag.
     */
    public readonly flag!: pulumi.Output<number>;
    /**
     * Endpoint IP address.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Endpoint MAC address.
     */
    public readonly mac!: pulumi.Output<string>;
    /**
     * Registering FortiGate SN.
     */
    public readonly regFortigate!: pulumi.Output<string>;
    /**
     * FortiClient registration status.
     */
    public readonly status!: pulumi.Output<number>;
    /**
     * FortiClient UID.
     */
    public readonly uid!: pulumi.Output<string>;
    /**
     * Registering vdom.
     */
    public readonly vdom!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a EndpointControlRegisteredForticlient resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EndpointControlRegisteredForticlientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointControlRegisteredForticlientArgs | EndpointControlRegisteredForticlientState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointControlRegisteredForticlientState | undefined;
            resourceInputs["flag"] = state ? state.flag : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["mac"] = state ? state.mac : undefined;
            resourceInputs["regFortigate"] = state ? state.regFortigate : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as EndpointControlRegisteredForticlientArgs | undefined;
            resourceInputs["flag"] = args ? args.flag : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["mac"] = args ? args.mac : undefined;
            resourceInputs["regFortigate"] = args ? args.regFortigate : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["uid"] = args ? args.uid : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EndpointControlRegisteredForticlient.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointControlRegisteredForticlient resources.
 */
export interface EndpointControlRegisteredForticlientState {
    /**
     * FortiClient registration flag.
     */
    flag?: pulumi.Input<number>;
    /**
     * Endpoint IP address.
     */
    ip?: pulumi.Input<string>;
    /**
     * Endpoint MAC address.
     */
    mac?: pulumi.Input<string>;
    /**
     * Registering FortiGate SN.
     */
    regFortigate?: pulumi.Input<string>;
    /**
     * FortiClient registration status.
     */
    status?: pulumi.Input<number>;
    /**
     * FortiClient UID.
     */
    uid?: pulumi.Input<string>;
    /**
     * Registering vdom.
     */
    vdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndpointControlRegisteredForticlient resource.
 */
export interface EndpointControlRegisteredForticlientArgs {
    /**
     * FortiClient registration flag.
     */
    flag?: pulumi.Input<number>;
    /**
     * Endpoint IP address.
     */
    ip?: pulumi.Input<string>;
    /**
     * Endpoint MAC address.
     */
    mac?: pulumi.Input<string>;
    /**
     * Registering FortiGate SN.
     */
    regFortigate?: pulumi.Input<string>;
    /**
     * FortiClient registration status.
     */
    status?: pulumi.Input<number>;
    /**
     * FortiClient UID.
     */
    uid?: pulumi.Input<string>;
    /**
     * Registering vdom.
     */
    vdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

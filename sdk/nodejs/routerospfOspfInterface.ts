// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class RouterospfOspfInterface extends pulumi.CustomResource {
    /**
     * Get an existing RouterospfOspfInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterospfOspfInterfaceState, opts?: pulumi.CustomResourceOptions): RouterospfOspfInterface {
        return new RouterospfOspfInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerospfOspfInterface:RouterospfOspfInterface';

    /**
     * Returns true if the given object is an instance of RouterospfOspfInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterospfOspfInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterospfOspfInterface.__pulumiType;
    }

    public readonly authentication!: pulumi.Output<string>;
    public readonly authenticationKey!: pulumi.Output<string | undefined>;
    public readonly bfd!: pulumi.Output<string>;
    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly cost!: pulumi.Output<number>;
    public readonly databaseFilterOut!: pulumi.Output<string>;
    public readonly deadInterval!: pulumi.Output<number>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly helloInterval!: pulumi.Output<number>;
    public readonly helloMultiplier!: pulumi.Output<number>;
    public readonly interface!: pulumi.Output<string>;
    public readonly ip!: pulumi.Output<string>;
    public readonly keychain!: pulumi.Output<string>;
    public readonly md5Key!: pulumi.Output<string>;
    public readonly md5Keychain!: pulumi.Output<string>;
    public readonly md5Keys!: pulumi.Output<outputs.RouterospfOspfInterfaceMd5Key[] | undefined>;
    public readonly mtu!: pulumi.Output<number>;
    public readonly mtuIgnore!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly networkType!: pulumi.Output<string>;
    public readonly prefixLength!: pulumi.Output<number>;
    public readonly priority!: pulumi.Output<number>;
    public readonly resyncTimeout!: pulumi.Output<number>;
    public readonly retransmitInterval!: pulumi.Output<number>;
    public readonly status!: pulumi.Output<string>;
    public readonly transmitDelay!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterospfOspfInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouterospfOspfInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterospfOspfInterfaceArgs | RouterospfOspfInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterospfOspfInterfaceState | undefined;
            resourceInputs["authentication"] = state ? state.authentication : undefined;
            resourceInputs["authenticationKey"] = state ? state.authenticationKey : undefined;
            resourceInputs["bfd"] = state ? state.bfd : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["cost"] = state ? state.cost : undefined;
            resourceInputs["databaseFilterOut"] = state ? state.databaseFilterOut : undefined;
            resourceInputs["deadInterval"] = state ? state.deadInterval : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["helloInterval"] = state ? state.helloInterval : undefined;
            resourceInputs["helloMultiplier"] = state ? state.helloMultiplier : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["keychain"] = state ? state.keychain : undefined;
            resourceInputs["md5Key"] = state ? state.md5Key : undefined;
            resourceInputs["md5Keychain"] = state ? state.md5Keychain : undefined;
            resourceInputs["md5Keys"] = state ? state.md5Keys : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["mtuIgnore"] = state ? state.mtuIgnore : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["prefixLength"] = state ? state.prefixLength : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["resyncTimeout"] = state ? state.resyncTimeout : undefined;
            resourceInputs["retransmitInterval"] = state ? state.retransmitInterval : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["transmitDelay"] = state ? state.transmitDelay : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RouterospfOspfInterfaceArgs | undefined;
            resourceInputs["authentication"] = args ? args.authentication : undefined;
            resourceInputs["authenticationKey"] = args ? args.authenticationKey : undefined;
            resourceInputs["bfd"] = args ? args.bfd : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["cost"] = args ? args.cost : undefined;
            resourceInputs["databaseFilterOut"] = args ? args.databaseFilterOut : undefined;
            resourceInputs["deadInterval"] = args ? args.deadInterval : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["helloInterval"] = args ? args.helloInterval : undefined;
            resourceInputs["helloMultiplier"] = args ? args.helloMultiplier : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["keychain"] = args ? args.keychain : undefined;
            resourceInputs["md5Key"] = args ? args.md5Key : undefined;
            resourceInputs["md5Keychain"] = args ? args.md5Keychain : undefined;
            resourceInputs["md5Keys"] = args ? args.md5Keys : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["mtuIgnore"] = args ? args.mtuIgnore : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["prefixLength"] = args ? args.prefixLength : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["resyncTimeout"] = args ? args.resyncTimeout : undefined;
            resourceInputs["retransmitInterval"] = args ? args.retransmitInterval : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["transmitDelay"] = args ? args.transmitDelay : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterospfOspfInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterospfOspfInterface resources.
 */
export interface RouterospfOspfInterfaceState {
    authentication?: pulumi.Input<string>;
    authenticationKey?: pulumi.Input<string>;
    bfd?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    cost?: pulumi.Input<number>;
    databaseFilterOut?: pulumi.Input<string>;
    deadInterval?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    helloInterval?: pulumi.Input<number>;
    helloMultiplier?: pulumi.Input<number>;
    interface?: pulumi.Input<string>;
    ip?: pulumi.Input<string>;
    keychain?: pulumi.Input<string>;
    md5Key?: pulumi.Input<string>;
    md5Keychain?: pulumi.Input<string>;
    md5Keys?: pulumi.Input<pulumi.Input<inputs.RouterospfOspfInterfaceMd5Key>[]>;
    mtu?: pulumi.Input<number>;
    mtuIgnore?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    networkType?: pulumi.Input<string>;
    prefixLength?: pulumi.Input<number>;
    priority?: pulumi.Input<number>;
    resyncTimeout?: pulumi.Input<number>;
    retransmitInterval?: pulumi.Input<number>;
    status?: pulumi.Input<string>;
    transmitDelay?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterospfOspfInterface resource.
 */
export interface RouterospfOspfInterfaceArgs {
    authentication?: pulumi.Input<string>;
    authenticationKey?: pulumi.Input<string>;
    bfd?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    cost?: pulumi.Input<number>;
    databaseFilterOut?: pulumi.Input<string>;
    deadInterval?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    helloInterval?: pulumi.Input<number>;
    helloMultiplier?: pulumi.Input<number>;
    interface?: pulumi.Input<string>;
    ip?: pulumi.Input<string>;
    keychain?: pulumi.Input<string>;
    md5Key?: pulumi.Input<string>;
    md5Keychain?: pulumi.Input<string>;
    md5Keys?: pulumi.Input<pulumi.Input<inputs.RouterospfOspfInterfaceMd5Key>[]>;
    mtu?: pulumi.Input<number>;
    mtuIgnore?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    networkType?: pulumi.Input<string>;
    prefixLength?: pulumi.Input<number>;
    priority?: pulumi.Input<number>;
    resyncTimeout?: pulumi.Input<number>;
    retransmitInterval?: pulumi.Input<number>;
    status?: pulumi.Input<string>;
    transmitDelay?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

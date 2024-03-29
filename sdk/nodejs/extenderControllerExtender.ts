// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ExtenderControllerExtender extends pulumi.CustomResource {
    /**
     * Get an existing ExtenderControllerExtender resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtenderControllerExtenderState, opts?: pulumi.CustomResourceOptions): ExtenderControllerExtender {
        return new ExtenderControllerExtender(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/extenderControllerExtender:ExtenderControllerExtender';

    /**
     * Returns true if the given object is an instance of ExtenderControllerExtender.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExtenderControllerExtender {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExtenderControllerExtender.__pulumiType;
    }

    public readonly aaaSharedSecret!: pulumi.Output<string | undefined>;
    public readonly accessPointName!: pulumi.Output<string>;
    public readonly admin!: pulumi.Output<string>;
    public readonly allowaccess!: pulumi.Output<string>;
    public readonly atDialScript!: pulumi.Output<string>;
    public readonly authorized!: pulumi.Output<string>;
    public readonly bandwidthLimit!: pulumi.Output<number>;
    public readonly billingStartDay!: pulumi.Output<number>;
    public readonly cdmaAaaSpi!: pulumi.Output<string>;
    public readonly cdmaHaSpi!: pulumi.Output<string>;
    public readonly cdmaNai!: pulumi.Output<string>;
    public readonly connStatus!: pulumi.Output<number>;
    public readonly controllerReport!: pulumi.Output<outputs.ExtenderControllerExtenderControllerReport>;
    public readonly description!: pulumi.Output<string>;
    public readonly deviceId!: pulumi.Output<number>;
    public readonly dialMode!: pulumi.Output<string>;
    public readonly dialStatus!: pulumi.Output<number>;
    public readonly enforceBandwidth!: pulumi.Output<string>;
    public readonly extName!: pulumi.Output<string>;
    public readonly extensionType!: pulumi.Output<string>;
    public readonly fosid!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly haSharedSecret!: pulumi.Output<string | undefined>;
    public readonly ifname!: pulumi.Output<string>;
    public readonly initiatedUpdate!: pulumi.Output<string>;
    public readonly loginPassword!: pulumi.Output<string | undefined>;
    public readonly loginPasswordChange!: pulumi.Output<string>;
    public readonly mode!: pulumi.Output<string>;
    public readonly modem1!: pulumi.Output<outputs.ExtenderControllerExtenderModem1>;
    public readonly modem2!: pulumi.Output<outputs.ExtenderControllerExtenderModem2>;
    public readonly modemPasswd!: pulumi.Output<string | undefined>;
    public readonly modemType!: pulumi.Output<string>;
    public readonly multiMode!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly overrideAllowaccess!: pulumi.Output<string>;
    public readonly overrideEnforceBandwidth!: pulumi.Output<string>;
    public readonly overrideLoginPasswordChange!: pulumi.Output<string>;
    public readonly pppAuthProtocol!: pulumi.Output<string>;
    public readonly pppEchoRequest!: pulumi.Output<string>;
    public readonly pppPassword!: pulumi.Output<string | undefined>;
    public readonly pppUsername!: pulumi.Output<string>;
    public readonly primaryHa!: pulumi.Output<string>;
    public readonly profile!: pulumi.Output<string>;
    public readonly quotaLimitMb!: pulumi.Output<number>;
    public readonly redial!: pulumi.Output<string>;
    public readonly redundantIntf!: pulumi.Output<string>;
    public readonly roaming!: pulumi.Output<string>;
    public readonly role!: pulumi.Output<string>;
    public readonly secondaryHa!: pulumi.Output<string>;
    public readonly simPin!: pulumi.Output<string | undefined>;
    public readonly vdom!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly wanExtension!: pulumi.Output<outputs.ExtenderControllerExtenderWanExtension>;
    public readonly wimaxAuthProtocol!: pulumi.Output<string>;
    public readonly wimaxCarrier!: pulumi.Output<string>;
    public readonly wimaxRealm!: pulumi.Output<string>;

    /**
     * Create a ExtenderControllerExtender resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExtenderControllerExtenderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtenderControllerExtenderArgs | ExtenderControllerExtenderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExtenderControllerExtenderState | undefined;
            resourceInputs["aaaSharedSecret"] = state ? state.aaaSharedSecret : undefined;
            resourceInputs["accessPointName"] = state ? state.accessPointName : undefined;
            resourceInputs["admin"] = state ? state.admin : undefined;
            resourceInputs["allowaccess"] = state ? state.allowaccess : undefined;
            resourceInputs["atDialScript"] = state ? state.atDialScript : undefined;
            resourceInputs["authorized"] = state ? state.authorized : undefined;
            resourceInputs["bandwidthLimit"] = state ? state.bandwidthLimit : undefined;
            resourceInputs["billingStartDay"] = state ? state.billingStartDay : undefined;
            resourceInputs["cdmaAaaSpi"] = state ? state.cdmaAaaSpi : undefined;
            resourceInputs["cdmaHaSpi"] = state ? state.cdmaHaSpi : undefined;
            resourceInputs["cdmaNai"] = state ? state.cdmaNai : undefined;
            resourceInputs["connStatus"] = state ? state.connStatus : undefined;
            resourceInputs["controllerReport"] = state ? state.controllerReport : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["dialMode"] = state ? state.dialMode : undefined;
            resourceInputs["dialStatus"] = state ? state.dialStatus : undefined;
            resourceInputs["enforceBandwidth"] = state ? state.enforceBandwidth : undefined;
            resourceInputs["extName"] = state ? state.extName : undefined;
            resourceInputs["extensionType"] = state ? state.extensionType : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["haSharedSecret"] = state ? state.haSharedSecret : undefined;
            resourceInputs["ifname"] = state ? state.ifname : undefined;
            resourceInputs["initiatedUpdate"] = state ? state.initiatedUpdate : undefined;
            resourceInputs["loginPassword"] = state ? state.loginPassword : undefined;
            resourceInputs["loginPasswordChange"] = state ? state.loginPasswordChange : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["modem1"] = state ? state.modem1 : undefined;
            resourceInputs["modem2"] = state ? state.modem2 : undefined;
            resourceInputs["modemPasswd"] = state ? state.modemPasswd : undefined;
            resourceInputs["modemType"] = state ? state.modemType : undefined;
            resourceInputs["multiMode"] = state ? state.multiMode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overrideAllowaccess"] = state ? state.overrideAllowaccess : undefined;
            resourceInputs["overrideEnforceBandwidth"] = state ? state.overrideEnforceBandwidth : undefined;
            resourceInputs["overrideLoginPasswordChange"] = state ? state.overrideLoginPasswordChange : undefined;
            resourceInputs["pppAuthProtocol"] = state ? state.pppAuthProtocol : undefined;
            resourceInputs["pppEchoRequest"] = state ? state.pppEchoRequest : undefined;
            resourceInputs["pppPassword"] = state ? state.pppPassword : undefined;
            resourceInputs["pppUsername"] = state ? state.pppUsername : undefined;
            resourceInputs["primaryHa"] = state ? state.primaryHa : undefined;
            resourceInputs["profile"] = state ? state.profile : undefined;
            resourceInputs["quotaLimitMb"] = state ? state.quotaLimitMb : undefined;
            resourceInputs["redial"] = state ? state.redial : undefined;
            resourceInputs["redundantIntf"] = state ? state.redundantIntf : undefined;
            resourceInputs["roaming"] = state ? state.roaming : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["secondaryHa"] = state ? state.secondaryHa : undefined;
            resourceInputs["simPin"] = state ? state.simPin : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wanExtension"] = state ? state.wanExtension : undefined;
            resourceInputs["wimaxAuthProtocol"] = state ? state.wimaxAuthProtocol : undefined;
            resourceInputs["wimaxCarrier"] = state ? state.wimaxCarrier : undefined;
            resourceInputs["wimaxRealm"] = state ? state.wimaxRealm : undefined;
        } else {
            const args = argsOrState as ExtenderControllerExtenderArgs | undefined;
            if ((!args || args.admin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'admin'");
            }
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["aaaSharedSecret"] = args?.aaaSharedSecret ? pulumi.secret(args.aaaSharedSecret) : undefined;
            resourceInputs["accessPointName"] = args ? args.accessPointName : undefined;
            resourceInputs["admin"] = args ? args.admin : undefined;
            resourceInputs["allowaccess"] = args ? args.allowaccess : undefined;
            resourceInputs["atDialScript"] = args ? args.atDialScript : undefined;
            resourceInputs["authorized"] = args ? args.authorized : undefined;
            resourceInputs["bandwidthLimit"] = args ? args.bandwidthLimit : undefined;
            resourceInputs["billingStartDay"] = args ? args.billingStartDay : undefined;
            resourceInputs["cdmaAaaSpi"] = args ? args.cdmaAaaSpi : undefined;
            resourceInputs["cdmaHaSpi"] = args ? args.cdmaHaSpi : undefined;
            resourceInputs["cdmaNai"] = args ? args.cdmaNai : undefined;
            resourceInputs["connStatus"] = args ? args.connStatus : undefined;
            resourceInputs["controllerReport"] = args ? args.controllerReport : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["dialMode"] = args ? args.dialMode : undefined;
            resourceInputs["dialStatus"] = args ? args.dialStatus : undefined;
            resourceInputs["enforceBandwidth"] = args ? args.enforceBandwidth : undefined;
            resourceInputs["extName"] = args ? args.extName : undefined;
            resourceInputs["extensionType"] = args ? args.extensionType : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["haSharedSecret"] = args?.haSharedSecret ? pulumi.secret(args.haSharedSecret) : undefined;
            resourceInputs["ifname"] = args ? args.ifname : undefined;
            resourceInputs["initiatedUpdate"] = args ? args.initiatedUpdate : undefined;
            resourceInputs["loginPassword"] = args ? args.loginPassword : undefined;
            resourceInputs["loginPasswordChange"] = args ? args.loginPasswordChange : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["modem1"] = args ? args.modem1 : undefined;
            resourceInputs["modem2"] = args ? args.modem2 : undefined;
            resourceInputs["modemPasswd"] = args?.modemPasswd ? pulumi.secret(args.modemPasswd) : undefined;
            resourceInputs["modemType"] = args ? args.modemType : undefined;
            resourceInputs["multiMode"] = args ? args.multiMode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrideAllowaccess"] = args ? args.overrideAllowaccess : undefined;
            resourceInputs["overrideEnforceBandwidth"] = args ? args.overrideEnforceBandwidth : undefined;
            resourceInputs["overrideLoginPasswordChange"] = args ? args.overrideLoginPasswordChange : undefined;
            resourceInputs["pppAuthProtocol"] = args ? args.pppAuthProtocol : undefined;
            resourceInputs["pppEchoRequest"] = args ? args.pppEchoRequest : undefined;
            resourceInputs["pppPassword"] = args?.pppPassword ? pulumi.secret(args.pppPassword) : undefined;
            resourceInputs["pppUsername"] = args ? args.pppUsername : undefined;
            resourceInputs["primaryHa"] = args ? args.primaryHa : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["quotaLimitMb"] = args ? args.quotaLimitMb : undefined;
            resourceInputs["redial"] = args ? args.redial : undefined;
            resourceInputs["redundantIntf"] = args ? args.redundantIntf : undefined;
            resourceInputs["roaming"] = args ? args.roaming : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["secondaryHa"] = args ? args.secondaryHa : undefined;
            resourceInputs["simPin"] = args?.simPin ? pulumi.secret(args.simPin) : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wanExtension"] = args ? args.wanExtension : undefined;
            resourceInputs["wimaxAuthProtocol"] = args ? args.wimaxAuthProtocol : undefined;
            resourceInputs["wimaxCarrier"] = args ? args.wimaxCarrier : undefined;
            resourceInputs["wimaxRealm"] = args ? args.wimaxRealm : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["aaaSharedSecret", "haSharedSecret", "modemPasswd", "pppPassword", "simPin"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ExtenderControllerExtender.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExtenderControllerExtender resources.
 */
export interface ExtenderControllerExtenderState {
    aaaSharedSecret?: pulumi.Input<string>;
    accessPointName?: pulumi.Input<string>;
    admin?: pulumi.Input<string>;
    allowaccess?: pulumi.Input<string>;
    atDialScript?: pulumi.Input<string>;
    authorized?: pulumi.Input<string>;
    bandwidthLimit?: pulumi.Input<number>;
    billingStartDay?: pulumi.Input<number>;
    cdmaAaaSpi?: pulumi.Input<string>;
    cdmaHaSpi?: pulumi.Input<string>;
    cdmaNai?: pulumi.Input<string>;
    connStatus?: pulumi.Input<number>;
    controllerReport?: pulumi.Input<inputs.ExtenderControllerExtenderControllerReport>;
    description?: pulumi.Input<string>;
    deviceId?: pulumi.Input<number>;
    dialMode?: pulumi.Input<string>;
    dialStatus?: pulumi.Input<number>;
    enforceBandwidth?: pulumi.Input<string>;
    extName?: pulumi.Input<string>;
    extensionType?: pulumi.Input<string>;
    fosid?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    haSharedSecret?: pulumi.Input<string>;
    ifname?: pulumi.Input<string>;
    initiatedUpdate?: pulumi.Input<string>;
    loginPassword?: pulumi.Input<string>;
    loginPasswordChange?: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    modem1?: pulumi.Input<inputs.ExtenderControllerExtenderModem1>;
    modem2?: pulumi.Input<inputs.ExtenderControllerExtenderModem2>;
    modemPasswd?: pulumi.Input<string>;
    modemType?: pulumi.Input<string>;
    multiMode?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    overrideAllowaccess?: pulumi.Input<string>;
    overrideEnforceBandwidth?: pulumi.Input<string>;
    overrideLoginPasswordChange?: pulumi.Input<string>;
    pppAuthProtocol?: pulumi.Input<string>;
    pppEchoRequest?: pulumi.Input<string>;
    pppPassword?: pulumi.Input<string>;
    pppUsername?: pulumi.Input<string>;
    primaryHa?: pulumi.Input<string>;
    profile?: pulumi.Input<string>;
    quotaLimitMb?: pulumi.Input<number>;
    redial?: pulumi.Input<string>;
    redundantIntf?: pulumi.Input<string>;
    roaming?: pulumi.Input<string>;
    role?: pulumi.Input<string>;
    secondaryHa?: pulumi.Input<string>;
    simPin?: pulumi.Input<string>;
    vdom?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
    wanExtension?: pulumi.Input<inputs.ExtenderControllerExtenderWanExtension>;
    wimaxAuthProtocol?: pulumi.Input<string>;
    wimaxCarrier?: pulumi.Input<string>;
    wimaxRealm?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExtenderControllerExtender resource.
 */
export interface ExtenderControllerExtenderArgs {
    aaaSharedSecret?: pulumi.Input<string>;
    accessPointName?: pulumi.Input<string>;
    admin: pulumi.Input<string>;
    allowaccess?: pulumi.Input<string>;
    atDialScript?: pulumi.Input<string>;
    authorized?: pulumi.Input<string>;
    bandwidthLimit?: pulumi.Input<number>;
    billingStartDay?: pulumi.Input<number>;
    cdmaAaaSpi?: pulumi.Input<string>;
    cdmaHaSpi?: pulumi.Input<string>;
    cdmaNai?: pulumi.Input<string>;
    connStatus?: pulumi.Input<number>;
    controllerReport?: pulumi.Input<inputs.ExtenderControllerExtenderControllerReport>;
    description?: pulumi.Input<string>;
    deviceId?: pulumi.Input<number>;
    dialMode?: pulumi.Input<string>;
    dialStatus?: pulumi.Input<number>;
    enforceBandwidth?: pulumi.Input<string>;
    extName?: pulumi.Input<string>;
    extensionType?: pulumi.Input<string>;
    fosid: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    haSharedSecret?: pulumi.Input<string>;
    ifname?: pulumi.Input<string>;
    initiatedUpdate?: pulumi.Input<string>;
    loginPassword?: pulumi.Input<string>;
    loginPasswordChange?: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    modem1?: pulumi.Input<inputs.ExtenderControllerExtenderModem1>;
    modem2?: pulumi.Input<inputs.ExtenderControllerExtenderModem2>;
    modemPasswd?: pulumi.Input<string>;
    modemType?: pulumi.Input<string>;
    multiMode?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    overrideAllowaccess?: pulumi.Input<string>;
    overrideEnforceBandwidth?: pulumi.Input<string>;
    overrideLoginPasswordChange?: pulumi.Input<string>;
    pppAuthProtocol?: pulumi.Input<string>;
    pppEchoRequest?: pulumi.Input<string>;
    pppPassword?: pulumi.Input<string>;
    pppUsername?: pulumi.Input<string>;
    primaryHa?: pulumi.Input<string>;
    profile?: pulumi.Input<string>;
    quotaLimitMb?: pulumi.Input<number>;
    redial?: pulumi.Input<string>;
    redundantIntf?: pulumi.Input<string>;
    roaming?: pulumi.Input<string>;
    role: pulumi.Input<string>;
    secondaryHa?: pulumi.Input<string>;
    simPin?: pulumi.Input<string>;
    vdom?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
    wanExtension?: pulumi.Input<inputs.ExtenderControllerExtenderWanExtension>;
    wimaxAuthProtocol?: pulumi.Input<string>;
    wimaxCarrier?: pulumi.Input<string>;
    wimaxRealm?: pulumi.Input<string>;
}

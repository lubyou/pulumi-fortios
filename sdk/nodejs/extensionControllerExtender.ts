// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ExtensionControllerExtender extends pulumi.CustomResource {
    /**
     * Get an existing ExtensionControllerExtender resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtensionControllerExtenderState, opts?: pulumi.CustomResourceOptions): ExtensionControllerExtender {
        return new ExtensionControllerExtender(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/extensionControllerExtender:ExtensionControllerExtender';

    /**
     * Returns true if the given object is an instance of ExtensionControllerExtender.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExtensionControllerExtender {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExtensionControllerExtender.__pulumiType;
    }

    public readonly allowaccess!: pulumi.Output<string>;
    public readonly authorized!: pulumi.Output<string>;
    public readonly bandwidthLimit!: pulumi.Output<number>;
    public readonly description!: pulumi.Output<string>;
    public readonly deviceId!: pulumi.Output<number>;
    public readonly enforceBandwidth!: pulumi.Output<string>;
    public readonly extName!: pulumi.Output<string>;
    public readonly extensionType!: pulumi.Output<string>;
    public readonly firmwareProvisionLatest!: pulumi.Output<string>;
    public readonly fosid!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly loginPassword!: pulumi.Output<string | undefined>;
    public readonly loginPasswordChange!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly overrideAllowaccess!: pulumi.Output<string>;
    public readonly overrideEnforceBandwidth!: pulumi.Output<string>;
    public readonly overrideLoginPasswordChange!: pulumi.Output<string>;
    public readonly profile!: pulumi.Output<string>;
    public readonly vdom!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly wanExtension!: pulumi.Output<outputs.ExtensionControllerExtenderWanExtension>;

    /**
     * Create a ExtensionControllerExtender resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExtensionControllerExtenderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtensionControllerExtenderArgs | ExtensionControllerExtenderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExtensionControllerExtenderState | undefined;
            resourceInputs["allowaccess"] = state ? state.allowaccess : undefined;
            resourceInputs["authorized"] = state ? state.authorized : undefined;
            resourceInputs["bandwidthLimit"] = state ? state.bandwidthLimit : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["enforceBandwidth"] = state ? state.enforceBandwidth : undefined;
            resourceInputs["extName"] = state ? state.extName : undefined;
            resourceInputs["extensionType"] = state ? state.extensionType : undefined;
            resourceInputs["firmwareProvisionLatest"] = state ? state.firmwareProvisionLatest : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["loginPassword"] = state ? state.loginPassword : undefined;
            resourceInputs["loginPasswordChange"] = state ? state.loginPasswordChange : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overrideAllowaccess"] = state ? state.overrideAllowaccess : undefined;
            resourceInputs["overrideEnforceBandwidth"] = state ? state.overrideEnforceBandwidth : undefined;
            resourceInputs["overrideLoginPasswordChange"] = state ? state.overrideLoginPasswordChange : undefined;
            resourceInputs["profile"] = state ? state.profile : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wanExtension"] = state ? state.wanExtension : undefined;
        } else {
            const args = argsOrState as ExtensionControllerExtenderArgs | undefined;
            resourceInputs["allowaccess"] = args ? args.allowaccess : undefined;
            resourceInputs["authorized"] = args ? args.authorized : undefined;
            resourceInputs["bandwidthLimit"] = args ? args.bandwidthLimit : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["enforceBandwidth"] = args ? args.enforceBandwidth : undefined;
            resourceInputs["extName"] = args ? args.extName : undefined;
            resourceInputs["extensionType"] = args ? args.extensionType : undefined;
            resourceInputs["firmwareProvisionLatest"] = args ? args.firmwareProvisionLatest : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["loginPassword"] = args ? args.loginPassword : undefined;
            resourceInputs["loginPasswordChange"] = args ? args.loginPasswordChange : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrideAllowaccess"] = args ? args.overrideAllowaccess : undefined;
            resourceInputs["overrideEnforceBandwidth"] = args ? args.overrideEnforceBandwidth : undefined;
            resourceInputs["overrideLoginPasswordChange"] = args ? args.overrideLoginPasswordChange : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wanExtension"] = args ? args.wanExtension : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExtensionControllerExtender.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExtensionControllerExtender resources.
 */
export interface ExtensionControllerExtenderState {
    allowaccess?: pulumi.Input<string>;
    authorized?: pulumi.Input<string>;
    bandwidthLimit?: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    deviceId?: pulumi.Input<number>;
    enforceBandwidth?: pulumi.Input<string>;
    extName?: pulumi.Input<string>;
    extensionType?: pulumi.Input<string>;
    firmwareProvisionLatest?: pulumi.Input<string>;
    fosid?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    loginPassword?: pulumi.Input<string>;
    loginPasswordChange?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    overrideAllowaccess?: pulumi.Input<string>;
    overrideEnforceBandwidth?: pulumi.Input<string>;
    overrideLoginPasswordChange?: pulumi.Input<string>;
    profile?: pulumi.Input<string>;
    vdom?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
    wanExtension?: pulumi.Input<inputs.ExtensionControllerExtenderWanExtension>;
}

/**
 * The set of arguments for constructing a ExtensionControllerExtender resource.
 */
export interface ExtensionControllerExtenderArgs {
    allowaccess?: pulumi.Input<string>;
    authorized?: pulumi.Input<string>;
    bandwidthLimit?: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    deviceId?: pulumi.Input<number>;
    enforceBandwidth?: pulumi.Input<string>;
    extName?: pulumi.Input<string>;
    extensionType?: pulumi.Input<string>;
    firmwareProvisionLatest?: pulumi.Input<string>;
    fosid?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    loginPassword?: pulumi.Input<string>;
    loginPasswordChange?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    overrideAllowaccess?: pulumi.Input<string>;
    overrideEnforceBandwidth?: pulumi.Input<string>;
    overrideLoginPasswordChange?: pulumi.Input<string>;
    profile?: pulumi.Input<string>;
    vdom?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
    wanExtension?: pulumi.Input<inputs.ExtensionControllerExtenderWanExtension>;
}

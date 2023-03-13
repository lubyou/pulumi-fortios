// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemLicenseVM extends pulumi.CustomResource {
    /**
     * Get an existing SystemLicenseVM resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemLicenseVMState, opts?: pulumi.CustomResourceOptions): SystemLicenseVM {
        return new SystemLicenseVM(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemLicenseVM:SystemLicenseVM';

    /**
     * Returns true if the given object is an instance of SystemLicenseVM.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemLicenseVM {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemLicenseVM.__pulumiType;
    }

    public readonly fileContent!: pulumi.Output<string>;

    /**
     * Create a SystemLicenseVM resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemLicenseVMArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemLicenseVMArgs | SystemLicenseVMState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemLicenseVMState | undefined;
            resourceInputs["fileContent"] = state ? state.fileContent : undefined;
        } else {
            const args = argsOrState as SystemLicenseVMArgs | undefined;
            if ((!args || args.fileContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileContent'");
            }
            resourceInputs["fileContent"] = args ? args.fileContent : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemLicenseVM.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemLicenseVM resources.
 */
export interface SystemLicenseVMState {
    fileContent?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemLicenseVM resource.
 */
export interface SystemLicenseVMArgs {
    fileContent: pulumi.Input<string>;
}

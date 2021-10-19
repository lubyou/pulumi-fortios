# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FortimanagerFirewallSecurityPolicyPackageArgs', 'FortimanagerFirewallSecurityPolicyPackage']

@pulumi.input_type
class FortimanagerFirewallSecurityPolicyPackageArgs:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 inspection_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FortimanagerFirewallSecurityPolicyPackage resource.
        :param pulumi.Input[str] adom: Source ADOM name. default is 'root'
        :param pulumi.Input[str] inspection_mode: Inspection Mode. Enum:[flow, proxy]. default is 'flow'
        :param pulumi.Input[str] name: Security policy package name.
        :param pulumi.Input[str] target: The installation target.
        :param pulumi.Input[str] vdom: Vdom of managed device. default is 'root'
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if inspection_mode is not None:
            pulumi.set(__self__, "inspection_mode", inspection_mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        """
        Source ADOM name. default is 'root'
        """
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="inspectionMode")
    def inspection_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Inspection Mode. Enum:[flow, proxy]. default is 'flow'
        """
        return pulumi.get(self, "inspection_mode")

    @inspection_mode.setter
    def inspection_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inspection_mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Security policy package name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        """
        The installation target.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        """
        Vdom of managed device. default is 'root'
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)


@pulumi.input_type
class _FortimanagerFirewallSecurityPolicyPackageState:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 inspection_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FortimanagerFirewallSecurityPolicyPackage resources.
        :param pulumi.Input[str] adom: Source ADOM name. default is 'root'
        :param pulumi.Input[str] inspection_mode: Inspection Mode. Enum:[flow, proxy]. default is 'flow'
        :param pulumi.Input[str] name: Security policy package name.
        :param pulumi.Input[str] target: The installation target.
        :param pulumi.Input[str] vdom: Vdom of managed device. default is 'root'
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if inspection_mode is not None:
            pulumi.set(__self__, "inspection_mode", inspection_mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        """
        Source ADOM name. default is 'root'
        """
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="inspectionMode")
    def inspection_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Inspection Mode. Enum:[flow, proxy]. default is 'flow'
        """
        return pulumi.get(self, "inspection_mode")

    @inspection_mode.setter
    def inspection_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inspection_mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Security policy package name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        """
        The installation target.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        """
        Vdom of managed device. default is 'root'
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)


class FortimanagerFirewallSecurityPolicyPackage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 inspection_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource supports Create/Read/Update/Delete firewall security policypackage on FortiManager which could be installed to the FortiGate later

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test1 = fortios.FortimanagerFirewallSecurityPolicyPackage("test1", target="FGVM64-test")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adom: Source ADOM name. default is 'root'
        :param pulumi.Input[str] inspection_mode: Inspection Mode. Enum:[flow, proxy]. default is 'flow'
        :param pulumi.Input[str] name: Security policy package name.
        :param pulumi.Input[str] target: The installation target.
        :param pulumi.Input[str] vdom: Vdom of managed device. default is 'root'
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FortimanagerFirewallSecurityPolicyPackageArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource supports Create/Read/Update/Delete firewall security policypackage on FortiManager which could be installed to the FortiGate later

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test1 = fortios.FortimanagerFirewallSecurityPolicyPackage("test1", target="FGVM64-test")
        ```

        :param str resource_name: The name of the resource.
        :param FortimanagerFirewallSecurityPolicyPackageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortimanagerFirewallSecurityPolicyPackageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 inspection_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortimanagerFirewallSecurityPolicyPackageArgs.__new__(FortimanagerFirewallSecurityPolicyPackageArgs)

            __props__.__dict__["adom"] = adom
            __props__.__dict__["inspection_mode"] = inspection_mode
            __props__.__dict__["name"] = name
            __props__.__dict__["target"] = target
            __props__.__dict__["vdom"] = vdom
        super(FortimanagerFirewallSecurityPolicyPackage, __self__).__init__(
            'fortios:index/fortimanagerFirewallSecurityPolicyPackage:FortimanagerFirewallSecurityPolicyPackage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adom: Optional[pulumi.Input[str]] = None,
            inspection_mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[str]] = None,
            vdom: Optional[pulumi.Input[str]] = None) -> 'FortimanagerFirewallSecurityPolicyPackage':
        """
        Get an existing FortimanagerFirewallSecurityPolicyPackage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adom: Source ADOM name. default is 'root'
        :param pulumi.Input[str] inspection_mode: Inspection Mode. Enum:[flow, proxy]. default is 'flow'
        :param pulumi.Input[str] name: Security policy package name.
        :param pulumi.Input[str] target: The installation target.
        :param pulumi.Input[str] vdom: Vdom of managed device. default is 'root'
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortimanagerFirewallSecurityPolicyPackageState.__new__(_FortimanagerFirewallSecurityPolicyPackageState)

        __props__.__dict__["adom"] = adom
        __props__.__dict__["inspection_mode"] = inspection_mode
        __props__.__dict__["name"] = name
        __props__.__dict__["target"] = target
        __props__.__dict__["vdom"] = vdom
        return FortimanagerFirewallSecurityPolicyPackage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def adom(self) -> pulumi.Output[Optional[str]]:
        """
        Source ADOM name. default is 'root'
        """
        return pulumi.get(self, "adom")

    @property
    @pulumi.getter(name="inspectionMode")
    def inspection_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Inspection Mode. Enum:[flow, proxy]. default is 'flow'
        """
        return pulumi.get(self, "inspection_mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Security policy package name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[Optional[str]]:
        """
        The installation target.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def vdom(self) -> pulumi.Output[Optional[str]]:
        """
        Vdom of managed device. default is 'root'
        """
        return pulumi.get(self, "vdom")


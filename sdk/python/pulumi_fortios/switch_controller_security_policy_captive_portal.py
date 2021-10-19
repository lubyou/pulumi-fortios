# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SwitchControllerSecurityPolicyCaptivePortalArgs', 'SwitchControllerSecurityPolicyCaptivePortal']

@pulumi.input_type
class SwitchControllerSecurityPolicyCaptivePortalArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerSecurityPolicyCaptivePortal resource.
        :param pulumi.Input[str] name: Policy name.
        :param pulumi.Input[str] policy_type: Policy type. Valid values: `captive-portal`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan: Names of VLANs that use captive portal authentication.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Policy type. Valid values: `captive-portal`.
        """
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def vlan(self) -> Optional[pulumi.Input[str]]:
        """
        Names of VLANs that use captive portal authentication.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan", value)


@pulumi.input_type
class _SwitchControllerSecurityPolicyCaptivePortalState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerSecurityPolicyCaptivePortal resources.
        :param pulumi.Input[str] name: Policy name.
        :param pulumi.Input[str] policy_type: Policy type. Valid values: `captive-portal`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan: Names of VLANs that use captive portal authentication.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Policy type. Valid values: `captive-portal`.
        """
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def vlan(self) -> Optional[pulumi.Input[str]]:
        """
        Names of VLANs that use captive portal authentication.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan", value)


class SwitchControllerSecurityPolicyCaptivePortal(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Names of VLANs that use captive portal authentication. Applies to FortiOS Version `<= 6.2.0`.

        ## Import

        SwitchControllerSecurityPolicy CaptivePortal can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/switchControllerSecurityPolicyCaptivePortal:SwitchControllerSecurityPolicyCaptivePortal labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Policy name.
        :param pulumi.Input[str] policy_type: Policy type. Valid values: `captive-portal`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan: Names of VLANs that use captive portal authentication.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchControllerSecurityPolicyCaptivePortalArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Names of VLANs that use captive portal authentication. Applies to FortiOS Version `<= 6.2.0`.

        ## Import

        SwitchControllerSecurityPolicy CaptivePortal can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/switchControllerSecurityPolicyCaptivePortal:SwitchControllerSecurityPolicyCaptivePortal labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SwitchControllerSecurityPolicyCaptivePortalArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerSecurityPolicyCaptivePortalArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[str]] = None,
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
            __props__ = SwitchControllerSecurityPolicyCaptivePortalArgs.__new__(SwitchControllerSecurityPolicyCaptivePortalArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["policy_type"] = policy_type
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vlan"] = vlan
        super(SwitchControllerSecurityPolicyCaptivePortal, __self__).__init__(
            'fortios:index/switchControllerSecurityPolicyCaptivePortal:SwitchControllerSecurityPolicyCaptivePortal',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vlan: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerSecurityPolicyCaptivePortal':
        """
        Get an existing SwitchControllerSecurityPolicyCaptivePortal resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Policy name.
        :param pulumi.Input[str] policy_type: Policy type. Valid values: `captive-portal`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan: Names of VLANs that use captive portal authentication.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerSecurityPolicyCaptivePortalState.__new__(_SwitchControllerSecurityPolicyCaptivePortalState)

        __props__.__dict__["name"] = name
        __props__.__dict__["policy_type"] = policy_type
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vlan"] = vlan
        return SwitchControllerSecurityPolicyCaptivePortal(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Output[str]:
        """
        Policy type. Valid values: `captive-portal`.
        """
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vlan(self) -> pulumi.Output[str]:
        """
        Names of VLANs that use captive portal authentication.
        """
        return pulumi.get(self, "vlan")


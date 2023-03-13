# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SwitchControllerLldpProfileArgs', 'SwitchControllerLldpProfile']

@pulumi.input_type
class SwitchControllerLldpProfileArgs:
    def __init__(__self__, *,
                 auto_isl: Optional[pulumi.Input[str]] = None,
                 auto_isl_hello_timer: Optional[pulumi.Input[int]] = None,
                 auto_isl_port_group: Optional[pulumi.Input[int]] = None,
                 auto_isl_receive_timeout: Optional[pulumi.Input[int]] = None,
                 auto_mclag_icl: Optional[pulumi.Input[str]] = None,
                 custom_tlvs: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileCustomTlvArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 med_location_services: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedLocationServiceArgs']]]] = None,
                 med_network_policies: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedNetworkPolicyArgs']]]] = None,
                 med_tlvs: Optional[pulumi.Input[str]] = None,
                 n8021_tlvs: Optional[pulumi.Input[str]] = None,
                 n8023_tlvs: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerLldpProfile resource.
        """
        if auto_isl is not None:
            pulumi.set(__self__, "auto_isl", auto_isl)
        if auto_isl_hello_timer is not None:
            pulumi.set(__self__, "auto_isl_hello_timer", auto_isl_hello_timer)
        if auto_isl_port_group is not None:
            pulumi.set(__self__, "auto_isl_port_group", auto_isl_port_group)
        if auto_isl_receive_timeout is not None:
            pulumi.set(__self__, "auto_isl_receive_timeout", auto_isl_receive_timeout)
        if auto_mclag_icl is not None:
            pulumi.set(__self__, "auto_mclag_icl", auto_mclag_icl)
        if custom_tlvs is not None:
            pulumi.set(__self__, "custom_tlvs", custom_tlvs)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if med_location_services is not None:
            pulumi.set(__self__, "med_location_services", med_location_services)
        if med_network_policies is not None:
            pulumi.set(__self__, "med_network_policies", med_network_policies)
        if med_tlvs is not None:
            pulumi.set(__self__, "med_tlvs", med_tlvs)
        if n8021_tlvs is not None:
            pulumi.set(__self__, "n8021_tlvs", n8021_tlvs)
        if n8023_tlvs is not None:
            pulumi.set(__self__, "n8023_tlvs", n8023_tlvs)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="autoIsl")
    def auto_isl(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_isl")

    @auto_isl.setter
    def auto_isl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_isl", value)

    @property
    @pulumi.getter(name="autoIslHelloTimer")
    def auto_isl_hello_timer(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "auto_isl_hello_timer")

    @auto_isl_hello_timer.setter
    def auto_isl_hello_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_isl_hello_timer", value)

    @property
    @pulumi.getter(name="autoIslPortGroup")
    def auto_isl_port_group(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "auto_isl_port_group")

    @auto_isl_port_group.setter
    def auto_isl_port_group(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_isl_port_group", value)

    @property
    @pulumi.getter(name="autoIslReceiveTimeout")
    def auto_isl_receive_timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "auto_isl_receive_timeout")

    @auto_isl_receive_timeout.setter
    def auto_isl_receive_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_isl_receive_timeout", value)

    @property
    @pulumi.getter(name="autoMclagIcl")
    def auto_mclag_icl(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_mclag_icl")

    @auto_mclag_icl.setter
    def auto_mclag_icl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_mclag_icl", value)

    @property
    @pulumi.getter(name="customTlvs")
    def custom_tlvs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileCustomTlvArgs']]]]:
        return pulumi.get(self, "custom_tlvs")

    @custom_tlvs.setter
    def custom_tlvs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileCustomTlvArgs']]]]):
        pulumi.set(self, "custom_tlvs", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="medLocationServices")
    def med_location_services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedLocationServiceArgs']]]]:
        return pulumi.get(self, "med_location_services")

    @med_location_services.setter
    def med_location_services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedLocationServiceArgs']]]]):
        pulumi.set(self, "med_location_services", value)

    @property
    @pulumi.getter(name="medNetworkPolicies")
    def med_network_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedNetworkPolicyArgs']]]]:
        return pulumi.get(self, "med_network_policies")

    @med_network_policies.setter
    def med_network_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedNetworkPolicyArgs']]]]):
        pulumi.set(self, "med_network_policies", value)

    @property
    @pulumi.getter(name="medTlvs")
    def med_tlvs(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "med_tlvs")

    @med_tlvs.setter
    def med_tlvs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "med_tlvs", value)

    @property
    @pulumi.getter(name="n8021Tlvs")
    def n8021_tlvs(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "n8021_tlvs")

    @n8021_tlvs.setter
    def n8021_tlvs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "n8021_tlvs", value)

    @property
    @pulumi.getter(name="n8023Tlvs")
    def n8023_tlvs(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "n8023_tlvs")

    @n8023_tlvs.setter
    def n8023_tlvs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "n8023_tlvs", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SwitchControllerLldpProfileState:
    def __init__(__self__, *,
                 auto_isl: Optional[pulumi.Input[str]] = None,
                 auto_isl_hello_timer: Optional[pulumi.Input[int]] = None,
                 auto_isl_port_group: Optional[pulumi.Input[int]] = None,
                 auto_isl_receive_timeout: Optional[pulumi.Input[int]] = None,
                 auto_mclag_icl: Optional[pulumi.Input[str]] = None,
                 custom_tlvs: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileCustomTlvArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 med_location_services: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedLocationServiceArgs']]]] = None,
                 med_network_policies: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedNetworkPolicyArgs']]]] = None,
                 med_tlvs: Optional[pulumi.Input[str]] = None,
                 n8021_tlvs: Optional[pulumi.Input[str]] = None,
                 n8023_tlvs: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerLldpProfile resources.
        """
        if auto_isl is not None:
            pulumi.set(__self__, "auto_isl", auto_isl)
        if auto_isl_hello_timer is not None:
            pulumi.set(__self__, "auto_isl_hello_timer", auto_isl_hello_timer)
        if auto_isl_port_group is not None:
            pulumi.set(__self__, "auto_isl_port_group", auto_isl_port_group)
        if auto_isl_receive_timeout is not None:
            pulumi.set(__self__, "auto_isl_receive_timeout", auto_isl_receive_timeout)
        if auto_mclag_icl is not None:
            pulumi.set(__self__, "auto_mclag_icl", auto_mclag_icl)
        if custom_tlvs is not None:
            pulumi.set(__self__, "custom_tlvs", custom_tlvs)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if med_location_services is not None:
            pulumi.set(__self__, "med_location_services", med_location_services)
        if med_network_policies is not None:
            pulumi.set(__self__, "med_network_policies", med_network_policies)
        if med_tlvs is not None:
            pulumi.set(__self__, "med_tlvs", med_tlvs)
        if n8021_tlvs is not None:
            pulumi.set(__self__, "n8021_tlvs", n8021_tlvs)
        if n8023_tlvs is not None:
            pulumi.set(__self__, "n8023_tlvs", n8023_tlvs)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="autoIsl")
    def auto_isl(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_isl")

    @auto_isl.setter
    def auto_isl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_isl", value)

    @property
    @pulumi.getter(name="autoIslHelloTimer")
    def auto_isl_hello_timer(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "auto_isl_hello_timer")

    @auto_isl_hello_timer.setter
    def auto_isl_hello_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_isl_hello_timer", value)

    @property
    @pulumi.getter(name="autoIslPortGroup")
    def auto_isl_port_group(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "auto_isl_port_group")

    @auto_isl_port_group.setter
    def auto_isl_port_group(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_isl_port_group", value)

    @property
    @pulumi.getter(name="autoIslReceiveTimeout")
    def auto_isl_receive_timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "auto_isl_receive_timeout")

    @auto_isl_receive_timeout.setter
    def auto_isl_receive_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_isl_receive_timeout", value)

    @property
    @pulumi.getter(name="autoMclagIcl")
    def auto_mclag_icl(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_mclag_icl")

    @auto_mclag_icl.setter
    def auto_mclag_icl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_mclag_icl", value)

    @property
    @pulumi.getter(name="customTlvs")
    def custom_tlvs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileCustomTlvArgs']]]]:
        return pulumi.get(self, "custom_tlvs")

    @custom_tlvs.setter
    def custom_tlvs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileCustomTlvArgs']]]]):
        pulumi.set(self, "custom_tlvs", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="medLocationServices")
    def med_location_services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedLocationServiceArgs']]]]:
        return pulumi.get(self, "med_location_services")

    @med_location_services.setter
    def med_location_services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedLocationServiceArgs']]]]):
        pulumi.set(self, "med_location_services", value)

    @property
    @pulumi.getter(name="medNetworkPolicies")
    def med_network_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedNetworkPolicyArgs']]]]:
        return pulumi.get(self, "med_network_policies")

    @med_network_policies.setter
    def med_network_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerLldpProfileMedNetworkPolicyArgs']]]]):
        pulumi.set(self, "med_network_policies", value)

    @property
    @pulumi.getter(name="medTlvs")
    def med_tlvs(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "med_tlvs")

    @med_tlvs.setter
    def med_tlvs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "med_tlvs", value)

    @property
    @pulumi.getter(name="n8021Tlvs")
    def n8021_tlvs(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "n8021_tlvs")

    @n8021_tlvs.setter
    def n8021_tlvs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "n8021_tlvs", value)

    @property
    @pulumi.getter(name="n8023Tlvs")
    def n8023_tlvs(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "n8023_tlvs")

    @n8023_tlvs.setter
    def n8023_tlvs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "n8023_tlvs", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SwitchControllerLldpProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_isl: Optional[pulumi.Input[str]] = None,
                 auto_isl_hello_timer: Optional[pulumi.Input[int]] = None,
                 auto_isl_port_group: Optional[pulumi.Input[int]] = None,
                 auto_isl_receive_timeout: Optional[pulumi.Input[int]] = None,
                 auto_mclag_icl: Optional[pulumi.Input[str]] = None,
                 custom_tlvs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerLldpProfileCustomTlvArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 med_location_services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerLldpProfileMedLocationServiceArgs']]]]] = None,
                 med_network_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerLldpProfileMedNetworkPolicyArgs']]]]] = None,
                 med_tlvs: Optional[pulumi.Input[str]] = None,
                 n8021_tlvs: Optional[pulumi.Input[str]] = None,
                 n8023_tlvs: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SwitchControllerLldpProfile resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchControllerLldpProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SwitchControllerLldpProfile resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SwitchControllerLldpProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerLldpProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_isl: Optional[pulumi.Input[str]] = None,
                 auto_isl_hello_timer: Optional[pulumi.Input[int]] = None,
                 auto_isl_port_group: Optional[pulumi.Input[int]] = None,
                 auto_isl_receive_timeout: Optional[pulumi.Input[int]] = None,
                 auto_mclag_icl: Optional[pulumi.Input[str]] = None,
                 custom_tlvs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerLldpProfileCustomTlvArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 med_location_services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerLldpProfileMedLocationServiceArgs']]]]] = None,
                 med_network_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerLldpProfileMedNetworkPolicyArgs']]]]] = None,
                 med_tlvs: Optional[pulumi.Input[str]] = None,
                 n8021_tlvs: Optional[pulumi.Input[str]] = None,
                 n8023_tlvs: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SwitchControllerLldpProfileArgs.__new__(SwitchControllerLldpProfileArgs)

            __props__.__dict__["auto_isl"] = auto_isl
            __props__.__dict__["auto_isl_hello_timer"] = auto_isl_hello_timer
            __props__.__dict__["auto_isl_port_group"] = auto_isl_port_group
            __props__.__dict__["auto_isl_receive_timeout"] = auto_isl_receive_timeout
            __props__.__dict__["auto_mclag_icl"] = auto_mclag_icl
            __props__.__dict__["custom_tlvs"] = custom_tlvs
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["med_location_services"] = med_location_services
            __props__.__dict__["med_network_policies"] = med_network_policies
            __props__.__dict__["med_tlvs"] = med_tlvs
            __props__.__dict__["n8021_tlvs"] = n8021_tlvs
            __props__.__dict__["n8023_tlvs"] = n8023_tlvs
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchControllerLldpProfile, __self__).__init__(
            'fortios:index/switchControllerLldpProfile:SwitchControllerLldpProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_isl: Optional[pulumi.Input[str]] = None,
            auto_isl_hello_timer: Optional[pulumi.Input[int]] = None,
            auto_isl_port_group: Optional[pulumi.Input[int]] = None,
            auto_isl_receive_timeout: Optional[pulumi.Input[int]] = None,
            auto_mclag_icl: Optional[pulumi.Input[str]] = None,
            custom_tlvs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerLldpProfileCustomTlvArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            med_location_services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerLldpProfileMedLocationServiceArgs']]]]] = None,
            med_network_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerLldpProfileMedNetworkPolicyArgs']]]]] = None,
            med_tlvs: Optional[pulumi.Input[str]] = None,
            n8021_tlvs: Optional[pulumi.Input[str]] = None,
            n8023_tlvs: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerLldpProfile':
        """
        Get an existing SwitchControllerLldpProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerLldpProfileState.__new__(_SwitchControllerLldpProfileState)

        __props__.__dict__["auto_isl"] = auto_isl
        __props__.__dict__["auto_isl_hello_timer"] = auto_isl_hello_timer
        __props__.__dict__["auto_isl_port_group"] = auto_isl_port_group
        __props__.__dict__["auto_isl_receive_timeout"] = auto_isl_receive_timeout
        __props__.__dict__["auto_mclag_icl"] = auto_mclag_icl
        __props__.__dict__["custom_tlvs"] = custom_tlvs
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["med_location_services"] = med_location_services
        __props__.__dict__["med_network_policies"] = med_network_policies
        __props__.__dict__["med_tlvs"] = med_tlvs
        __props__.__dict__["n8021_tlvs"] = n8021_tlvs
        __props__.__dict__["n8023_tlvs"] = n8023_tlvs
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchControllerLldpProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoIsl")
    def auto_isl(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auto_isl")

    @property
    @pulumi.getter(name="autoIslHelloTimer")
    def auto_isl_hello_timer(self) -> pulumi.Output[int]:
        return pulumi.get(self, "auto_isl_hello_timer")

    @property
    @pulumi.getter(name="autoIslPortGroup")
    def auto_isl_port_group(self) -> pulumi.Output[int]:
        return pulumi.get(self, "auto_isl_port_group")

    @property
    @pulumi.getter(name="autoIslReceiveTimeout")
    def auto_isl_receive_timeout(self) -> pulumi.Output[int]:
        return pulumi.get(self, "auto_isl_receive_timeout")

    @property
    @pulumi.getter(name="autoMclagIcl")
    def auto_mclag_icl(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auto_mclag_icl")

    @property
    @pulumi.getter(name="customTlvs")
    def custom_tlvs(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerLldpProfileCustomTlv']]]:
        return pulumi.get(self, "custom_tlvs")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="medLocationServices")
    def med_location_services(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerLldpProfileMedLocationService']]]:
        return pulumi.get(self, "med_location_services")

    @property
    @pulumi.getter(name="medNetworkPolicies")
    def med_network_policies(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerLldpProfileMedNetworkPolicy']]]:
        return pulumi.get(self, "med_network_policies")

    @property
    @pulumi.getter(name="medTlvs")
    def med_tlvs(self) -> pulumi.Output[str]:
        return pulumi.get(self, "med_tlvs")

    @property
    @pulumi.getter(name="n8021Tlvs")
    def n8021_tlvs(self) -> pulumi.Output[str]:
        return pulumi.get(self, "n8021_tlvs")

    @property
    @pulumi.getter(name="n8023Tlvs")
    def n8023_tlvs(self) -> pulumi.Output[str]:
        return pulumi.get(self, "n8023_tlvs")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")


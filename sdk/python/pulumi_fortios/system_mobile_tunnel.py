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

__all__ = ['SystemMobileTunnelArgs', 'SystemMobileTunnel']

@pulumi.input_type
class SystemMobileTunnelArgs:
    def __init__(__self__, *,
                 hash_algorithm: pulumi.Input[str],
                 home_agent: pulumi.Input[str],
                 lifetime: pulumi.Input[int],
                 n_mhae_key_type: pulumi.Input[str],
                 n_mhae_spi: pulumi.Input[int],
                 reg_interval: pulumi.Input[int],
                 reg_retry: pulumi.Input[int],
                 renew_interval: pulumi.Input[int],
                 roaming_interface: pulumi.Input[str],
                 tunnel_mode: pulumi.Input[str],
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 home_address: Optional[pulumi.Input[str]] = None,
                 n_mhae_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input['SystemMobileTunnelNetworkArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemMobileTunnel resource.
        """
        pulumi.set(__self__, "hash_algorithm", hash_algorithm)
        pulumi.set(__self__, "home_agent", home_agent)
        pulumi.set(__self__, "lifetime", lifetime)
        pulumi.set(__self__, "n_mhae_key_type", n_mhae_key_type)
        pulumi.set(__self__, "n_mhae_spi", n_mhae_spi)
        pulumi.set(__self__, "reg_interval", reg_interval)
        pulumi.set(__self__, "reg_retry", reg_retry)
        pulumi.set(__self__, "renew_interval", renew_interval)
        pulumi.set(__self__, "roaming_interface", roaming_interface)
        pulumi.set(__self__, "tunnel_mode", tunnel_mode)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if home_address is not None:
            pulumi.set(__self__, "home_address", home_address)
        if n_mhae_key is not None:
            pulumi.set(__self__, "n_mhae_key", n_mhae_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="hashAlgorithm")
    def hash_algorithm(self) -> pulumi.Input[str]:
        return pulumi.get(self, "hash_algorithm")

    @hash_algorithm.setter
    def hash_algorithm(self, value: pulumi.Input[str]):
        pulumi.set(self, "hash_algorithm", value)

    @property
    @pulumi.getter(name="homeAgent")
    def home_agent(self) -> pulumi.Input[str]:
        return pulumi.get(self, "home_agent")

    @home_agent.setter
    def home_agent(self, value: pulumi.Input[str]):
        pulumi.set(self, "home_agent", value)

    @property
    @pulumi.getter
    def lifetime(self) -> pulumi.Input[int]:
        return pulumi.get(self, "lifetime")

    @lifetime.setter
    def lifetime(self, value: pulumi.Input[int]):
        pulumi.set(self, "lifetime", value)

    @property
    @pulumi.getter(name="nMhaeKeyType")
    def n_mhae_key_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "n_mhae_key_type")

    @n_mhae_key_type.setter
    def n_mhae_key_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "n_mhae_key_type", value)

    @property
    @pulumi.getter(name="nMhaeSpi")
    def n_mhae_spi(self) -> pulumi.Input[int]:
        return pulumi.get(self, "n_mhae_spi")

    @n_mhae_spi.setter
    def n_mhae_spi(self, value: pulumi.Input[int]):
        pulumi.set(self, "n_mhae_spi", value)

    @property
    @pulumi.getter(name="regInterval")
    def reg_interval(self) -> pulumi.Input[int]:
        return pulumi.get(self, "reg_interval")

    @reg_interval.setter
    def reg_interval(self, value: pulumi.Input[int]):
        pulumi.set(self, "reg_interval", value)

    @property
    @pulumi.getter(name="regRetry")
    def reg_retry(self) -> pulumi.Input[int]:
        return pulumi.get(self, "reg_retry")

    @reg_retry.setter
    def reg_retry(self, value: pulumi.Input[int]):
        pulumi.set(self, "reg_retry", value)

    @property
    @pulumi.getter(name="renewInterval")
    def renew_interval(self) -> pulumi.Input[int]:
        return pulumi.get(self, "renew_interval")

    @renew_interval.setter
    def renew_interval(self, value: pulumi.Input[int]):
        pulumi.set(self, "renew_interval", value)

    @property
    @pulumi.getter(name="roamingInterface")
    def roaming_interface(self) -> pulumi.Input[str]:
        return pulumi.get(self, "roaming_interface")

    @roaming_interface.setter
    def roaming_interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "roaming_interface", value)

    @property
    @pulumi.getter(name="tunnelMode")
    def tunnel_mode(self) -> pulumi.Input[str]:
        return pulumi.get(self, "tunnel_mode")

    @tunnel_mode.setter
    def tunnel_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "tunnel_mode", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="homeAddress")
    def home_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "home_address")

    @home_address.setter
    def home_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "home_address", value)

    @property
    @pulumi.getter(name="nMhaeKey")
    def n_mhae_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "n_mhae_key")

    @n_mhae_key.setter
    def n_mhae_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "n_mhae_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemMobileTunnelNetworkArgs']]]]:
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemMobileTunnelNetworkArgs']]]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemMobileTunnelState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 hash_algorithm: Optional[pulumi.Input[str]] = None,
                 home_address: Optional[pulumi.Input[str]] = None,
                 home_agent: Optional[pulumi.Input[str]] = None,
                 lifetime: Optional[pulumi.Input[int]] = None,
                 n_mhae_key: Optional[pulumi.Input[str]] = None,
                 n_mhae_key_type: Optional[pulumi.Input[str]] = None,
                 n_mhae_spi: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input['SystemMobileTunnelNetworkArgs']]]] = None,
                 reg_interval: Optional[pulumi.Input[int]] = None,
                 reg_retry: Optional[pulumi.Input[int]] = None,
                 renew_interval: Optional[pulumi.Input[int]] = None,
                 roaming_interface: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tunnel_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemMobileTunnel resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if hash_algorithm is not None:
            pulumi.set(__self__, "hash_algorithm", hash_algorithm)
        if home_address is not None:
            pulumi.set(__self__, "home_address", home_address)
        if home_agent is not None:
            pulumi.set(__self__, "home_agent", home_agent)
        if lifetime is not None:
            pulumi.set(__self__, "lifetime", lifetime)
        if n_mhae_key is not None:
            pulumi.set(__self__, "n_mhae_key", n_mhae_key)
        if n_mhae_key_type is not None:
            pulumi.set(__self__, "n_mhae_key_type", n_mhae_key_type)
        if n_mhae_spi is not None:
            pulumi.set(__self__, "n_mhae_spi", n_mhae_spi)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if reg_interval is not None:
            pulumi.set(__self__, "reg_interval", reg_interval)
        if reg_retry is not None:
            pulumi.set(__self__, "reg_retry", reg_retry)
        if renew_interval is not None:
            pulumi.set(__self__, "renew_interval", renew_interval)
        if roaming_interface is not None:
            pulumi.set(__self__, "roaming_interface", roaming_interface)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tunnel_mode is not None:
            pulumi.set(__self__, "tunnel_mode", tunnel_mode)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="hashAlgorithm")
    def hash_algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hash_algorithm")

    @hash_algorithm.setter
    def hash_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hash_algorithm", value)

    @property
    @pulumi.getter(name="homeAddress")
    def home_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "home_address")

    @home_address.setter
    def home_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "home_address", value)

    @property
    @pulumi.getter(name="homeAgent")
    def home_agent(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "home_agent")

    @home_agent.setter
    def home_agent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "home_agent", value)

    @property
    @pulumi.getter
    def lifetime(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "lifetime")

    @lifetime.setter
    def lifetime(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lifetime", value)

    @property
    @pulumi.getter(name="nMhaeKey")
    def n_mhae_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "n_mhae_key")

    @n_mhae_key.setter
    def n_mhae_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "n_mhae_key", value)

    @property
    @pulumi.getter(name="nMhaeKeyType")
    def n_mhae_key_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "n_mhae_key_type")

    @n_mhae_key_type.setter
    def n_mhae_key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "n_mhae_key_type", value)

    @property
    @pulumi.getter(name="nMhaeSpi")
    def n_mhae_spi(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "n_mhae_spi")

    @n_mhae_spi.setter
    def n_mhae_spi(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "n_mhae_spi", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemMobileTunnelNetworkArgs']]]]:
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemMobileTunnelNetworkArgs']]]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter(name="regInterval")
    def reg_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "reg_interval")

    @reg_interval.setter
    def reg_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reg_interval", value)

    @property
    @pulumi.getter(name="regRetry")
    def reg_retry(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "reg_retry")

    @reg_retry.setter
    def reg_retry(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reg_retry", value)

    @property
    @pulumi.getter(name="renewInterval")
    def renew_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "renew_interval")

    @renew_interval.setter
    def renew_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "renew_interval", value)

    @property
    @pulumi.getter(name="roamingInterface")
    def roaming_interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "roaming_interface")

    @roaming_interface.setter
    def roaming_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "roaming_interface", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="tunnelMode")
    def tunnel_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tunnel_mode")

    @tunnel_mode.setter
    def tunnel_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tunnel_mode", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemMobileTunnel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 hash_algorithm: Optional[pulumi.Input[str]] = None,
                 home_address: Optional[pulumi.Input[str]] = None,
                 home_agent: Optional[pulumi.Input[str]] = None,
                 lifetime: Optional[pulumi.Input[int]] = None,
                 n_mhae_key: Optional[pulumi.Input[str]] = None,
                 n_mhae_key_type: Optional[pulumi.Input[str]] = None,
                 n_mhae_spi: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemMobileTunnelNetworkArgs']]]]] = None,
                 reg_interval: Optional[pulumi.Input[int]] = None,
                 reg_retry: Optional[pulumi.Input[int]] = None,
                 renew_interval: Optional[pulumi.Input[int]] = None,
                 roaming_interface: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tunnel_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SystemMobileTunnel resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemMobileTunnelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SystemMobileTunnel resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SystemMobileTunnelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemMobileTunnelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 hash_algorithm: Optional[pulumi.Input[str]] = None,
                 home_address: Optional[pulumi.Input[str]] = None,
                 home_agent: Optional[pulumi.Input[str]] = None,
                 lifetime: Optional[pulumi.Input[int]] = None,
                 n_mhae_key: Optional[pulumi.Input[str]] = None,
                 n_mhae_key_type: Optional[pulumi.Input[str]] = None,
                 n_mhae_spi: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemMobileTunnelNetworkArgs']]]]] = None,
                 reg_interval: Optional[pulumi.Input[int]] = None,
                 reg_retry: Optional[pulumi.Input[int]] = None,
                 renew_interval: Optional[pulumi.Input[int]] = None,
                 roaming_interface: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tunnel_mode: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemMobileTunnelArgs.__new__(SystemMobileTunnelArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if hash_algorithm is None and not opts.urn:
                raise TypeError("Missing required property 'hash_algorithm'")
            __props__.__dict__["hash_algorithm"] = hash_algorithm
            __props__.__dict__["home_address"] = home_address
            if home_agent is None and not opts.urn:
                raise TypeError("Missing required property 'home_agent'")
            __props__.__dict__["home_agent"] = home_agent
            if lifetime is None and not opts.urn:
                raise TypeError("Missing required property 'lifetime'")
            __props__.__dict__["lifetime"] = lifetime
            __props__.__dict__["n_mhae_key"] = None if n_mhae_key is None else pulumi.Output.secret(n_mhae_key)
            if n_mhae_key_type is None and not opts.urn:
                raise TypeError("Missing required property 'n_mhae_key_type'")
            __props__.__dict__["n_mhae_key_type"] = n_mhae_key_type
            if n_mhae_spi is None and not opts.urn:
                raise TypeError("Missing required property 'n_mhae_spi'")
            __props__.__dict__["n_mhae_spi"] = n_mhae_spi
            __props__.__dict__["name"] = name
            __props__.__dict__["networks"] = networks
            if reg_interval is None and not opts.urn:
                raise TypeError("Missing required property 'reg_interval'")
            __props__.__dict__["reg_interval"] = reg_interval
            if reg_retry is None and not opts.urn:
                raise TypeError("Missing required property 'reg_retry'")
            __props__.__dict__["reg_retry"] = reg_retry
            if renew_interval is None and not opts.urn:
                raise TypeError("Missing required property 'renew_interval'")
            __props__.__dict__["renew_interval"] = renew_interval
            if roaming_interface is None and not opts.urn:
                raise TypeError("Missing required property 'roaming_interface'")
            __props__.__dict__["roaming_interface"] = roaming_interface
            __props__.__dict__["status"] = status
            if tunnel_mode is None and not opts.urn:
                raise TypeError("Missing required property 'tunnel_mode'")
            __props__.__dict__["tunnel_mode"] = tunnel_mode
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["nMhaeKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SystemMobileTunnel, __self__).__init__(
            'fortios:index/systemMobileTunnel:SystemMobileTunnel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            hash_algorithm: Optional[pulumi.Input[str]] = None,
            home_address: Optional[pulumi.Input[str]] = None,
            home_agent: Optional[pulumi.Input[str]] = None,
            lifetime: Optional[pulumi.Input[int]] = None,
            n_mhae_key: Optional[pulumi.Input[str]] = None,
            n_mhae_key_type: Optional[pulumi.Input[str]] = None,
            n_mhae_spi: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemMobileTunnelNetworkArgs']]]]] = None,
            reg_interval: Optional[pulumi.Input[int]] = None,
            reg_retry: Optional[pulumi.Input[int]] = None,
            renew_interval: Optional[pulumi.Input[int]] = None,
            roaming_interface: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tunnel_mode: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemMobileTunnel':
        """
        Get an existing SystemMobileTunnel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemMobileTunnelState.__new__(_SystemMobileTunnelState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["hash_algorithm"] = hash_algorithm
        __props__.__dict__["home_address"] = home_address
        __props__.__dict__["home_agent"] = home_agent
        __props__.__dict__["lifetime"] = lifetime
        __props__.__dict__["n_mhae_key"] = n_mhae_key
        __props__.__dict__["n_mhae_key_type"] = n_mhae_key_type
        __props__.__dict__["n_mhae_spi"] = n_mhae_spi
        __props__.__dict__["name"] = name
        __props__.__dict__["networks"] = networks
        __props__.__dict__["reg_interval"] = reg_interval
        __props__.__dict__["reg_retry"] = reg_retry
        __props__.__dict__["renew_interval"] = renew_interval
        __props__.__dict__["roaming_interface"] = roaming_interface
        __props__.__dict__["status"] = status
        __props__.__dict__["tunnel_mode"] = tunnel_mode
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemMobileTunnel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="hashAlgorithm")
    def hash_algorithm(self) -> pulumi.Output[str]:
        return pulumi.get(self, "hash_algorithm")

    @property
    @pulumi.getter(name="homeAddress")
    def home_address(self) -> pulumi.Output[str]:
        return pulumi.get(self, "home_address")

    @property
    @pulumi.getter(name="homeAgent")
    def home_agent(self) -> pulumi.Output[str]:
        return pulumi.get(self, "home_agent")

    @property
    @pulumi.getter
    def lifetime(self) -> pulumi.Output[int]:
        return pulumi.get(self, "lifetime")

    @property
    @pulumi.getter(name="nMhaeKey")
    def n_mhae_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "n_mhae_key")

    @property
    @pulumi.getter(name="nMhaeKeyType")
    def n_mhae_key_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "n_mhae_key_type")

    @property
    @pulumi.getter(name="nMhaeSpi")
    def n_mhae_spi(self) -> pulumi.Output[int]:
        return pulumi.get(self, "n_mhae_spi")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> pulumi.Output[Optional[Sequence['outputs.SystemMobileTunnelNetwork']]]:
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="regInterval")
    def reg_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "reg_interval")

    @property
    @pulumi.getter(name="regRetry")
    def reg_retry(self) -> pulumi.Output[int]:
        return pulumi.get(self, "reg_retry")

    @property
    @pulumi.getter(name="renewInterval")
    def renew_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "renew_interval")

    @property
    @pulumi.getter(name="roamingInterface")
    def roaming_interface(self) -> pulumi.Output[str]:
        return pulumi.get(self, "roaming_interface")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tunnelMode")
    def tunnel_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tunnel_mode")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")


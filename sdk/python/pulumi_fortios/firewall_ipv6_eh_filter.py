# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallIpv6EhFilterArgs', 'FirewallIpv6EhFilter']

@pulumi.input_type
class FirewallIpv6EhFilterArgs:
    def __init__(__self__, *,
                 auth: Optional[pulumi.Input[str]] = None,
                 dest_opt: Optional[pulumi.Input[str]] = None,
                 fragment: Optional[pulumi.Input[str]] = None,
                 hdopt_type: Optional[pulumi.Input[int]] = None,
                 hop_opt: Optional[pulumi.Input[str]] = None,
                 no_next: Optional[pulumi.Input[str]] = None,
                 routing: Optional[pulumi.Input[str]] = None,
                 routing_type: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallIpv6EhFilter resource.
        :param pulumi.Input[str] auth: Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dest_opt: Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] fragment: Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[int] hdopt_type: Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        :param pulumi.Input[str] hop_opt: Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] no_next: Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
        :param pulumi.Input[str] routing: Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[int] routing_type: Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth is not None:
            pulumi.set(__self__, "auth", auth)
        if dest_opt is not None:
            pulumi.set(__self__, "dest_opt", dest_opt)
        if fragment is not None:
            pulumi.set(__self__, "fragment", fragment)
        if hdopt_type is not None:
            pulumi.set(__self__, "hdopt_type", hdopt_type)
        if hop_opt is not None:
            pulumi.set(__self__, "hop_opt", hop_opt)
        if no_next is not None:
            pulumi.set(__self__, "no_next", no_next)
        if routing is not None:
            pulumi.set(__self__, "routing", routing)
        if routing_type is not None:
            pulumi.set(__self__, "routing_type", routing_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter(name="destOpt")
    def dest_opt(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dest_opt")

    @dest_opt.setter
    def dest_opt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dest_opt", value)

    @property
    @pulumi.getter
    def fragment(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fragment")

    @fragment.setter
    def fragment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fragment", value)

    @property
    @pulumi.getter(name="hdoptType")
    def hdopt_type(self) -> Optional[pulumi.Input[int]]:
        """
        Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        """
        return pulumi.get(self, "hdopt_type")

    @hdopt_type.setter
    def hdopt_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hdopt_type", value)

    @property
    @pulumi.getter(name="hopOpt")
    def hop_opt(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "hop_opt")

    @hop_opt.setter
    def hop_opt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hop_opt", value)

    @property
    @pulumi.getter(name="noNext")
    def no_next(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "no_next")

    @no_next.setter
    def no_next(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "no_next", value)

    @property
    @pulumi.getter
    def routing(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "routing")

    @routing.setter
    def routing(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routing", value)

    @property
    @pulumi.getter(name="routingType")
    def routing_type(self) -> Optional[pulumi.Input[int]]:
        """
        Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        """
        return pulumi.get(self, "routing_type")

    @routing_type.setter
    def routing_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "routing_type", value)

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


@pulumi.input_type
class _FirewallIpv6EhFilterState:
    def __init__(__self__, *,
                 auth: Optional[pulumi.Input[str]] = None,
                 dest_opt: Optional[pulumi.Input[str]] = None,
                 fragment: Optional[pulumi.Input[str]] = None,
                 hdopt_type: Optional[pulumi.Input[int]] = None,
                 hop_opt: Optional[pulumi.Input[str]] = None,
                 no_next: Optional[pulumi.Input[str]] = None,
                 routing: Optional[pulumi.Input[str]] = None,
                 routing_type: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallIpv6EhFilter resources.
        :param pulumi.Input[str] auth: Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dest_opt: Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] fragment: Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[int] hdopt_type: Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        :param pulumi.Input[str] hop_opt: Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] no_next: Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
        :param pulumi.Input[str] routing: Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[int] routing_type: Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auth is not None:
            pulumi.set(__self__, "auth", auth)
        if dest_opt is not None:
            pulumi.set(__self__, "dest_opt", dest_opt)
        if fragment is not None:
            pulumi.set(__self__, "fragment", fragment)
        if hdopt_type is not None:
            pulumi.set(__self__, "hdopt_type", hdopt_type)
        if hop_opt is not None:
            pulumi.set(__self__, "hop_opt", hop_opt)
        if no_next is not None:
            pulumi.set(__self__, "no_next", no_next)
        if routing is not None:
            pulumi.set(__self__, "routing", routing)
        if routing_type is not None:
            pulumi.set(__self__, "routing_type", routing_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter(name="destOpt")
    def dest_opt(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dest_opt")

    @dest_opt.setter
    def dest_opt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dest_opt", value)

    @property
    @pulumi.getter
    def fragment(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fragment")

    @fragment.setter
    def fragment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fragment", value)

    @property
    @pulumi.getter(name="hdoptType")
    def hdopt_type(self) -> Optional[pulumi.Input[int]]:
        """
        Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        """
        return pulumi.get(self, "hdopt_type")

    @hdopt_type.setter
    def hdopt_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hdopt_type", value)

    @property
    @pulumi.getter(name="hopOpt")
    def hop_opt(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "hop_opt")

    @hop_opt.setter
    def hop_opt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hop_opt", value)

    @property
    @pulumi.getter(name="noNext")
    def no_next(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "no_next")

    @no_next.setter
    def no_next(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "no_next", value)

    @property
    @pulumi.getter
    def routing(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "routing")

    @routing.setter
    def routing(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routing", value)

    @property
    @pulumi.getter(name="routingType")
    def routing_type(self) -> Optional[pulumi.Input[int]]:
        """
        Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        """
        return pulumi.get(self, "routing_type")

    @routing_type.setter
    def routing_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "routing_type", value)

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


class FirewallIpv6EhFilter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 dest_opt: Optional[pulumi.Input[str]] = None,
                 fragment: Optional[pulumi.Input[str]] = None,
                 hdopt_type: Optional[pulumi.Input[int]] = None,
                 hop_opt: Optional[pulumi.Input[str]] = None,
                 no_next: Optional[pulumi.Input[str]] = None,
                 routing: Optional[pulumi.Input[str]] = None,
                 routing_type: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPv6 extension header filter.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallIpv6EhFilter("trname",
            auth="disable",
            dest_opt="disable",
            fragment="disable",
            hop_opt="disable",
            no_next="disable",
            routing="enable")
        ```

        ## Import

        Firewall Ipv6EhFilter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallIpv6EhFilter:FirewallIpv6EhFilter labelname FirewallIpv6EhFilter
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth: Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dest_opt: Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] fragment: Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[int] hdopt_type: Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        :param pulumi.Input[str] hop_opt: Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] no_next: Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
        :param pulumi.Input[str] routing: Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[int] routing_type: Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallIpv6EhFilterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPv6 extension header filter.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallIpv6EhFilter("trname",
            auth="disable",
            dest_opt="disable",
            fragment="disable",
            hop_opt="disable",
            no_next="disable",
            routing="enable")
        ```

        ## Import

        Firewall Ipv6EhFilter can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallIpv6EhFilter:FirewallIpv6EhFilter labelname FirewallIpv6EhFilter
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallIpv6EhFilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallIpv6EhFilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 dest_opt: Optional[pulumi.Input[str]] = None,
                 fragment: Optional[pulumi.Input[str]] = None,
                 hdopt_type: Optional[pulumi.Input[int]] = None,
                 hop_opt: Optional[pulumi.Input[str]] = None,
                 no_next: Optional[pulumi.Input[str]] = None,
                 routing: Optional[pulumi.Input[str]] = None,
                 routing_type: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallIpv6EhFilterArgs.__new__(FirewallIpv6EhFilterArgs)

            __props__.__dict__["auth"] = auth
            __props__.__dict__["dest_opt"] = dest_opt
            __props__.__dict__["fragment"] = fragment
            __props__.__dict__["hdopt_type"] = hdopt_type
            __props__.__dict__["hop_opt"] = hop_opt
            __props__.__dict__["no_next"] = no_next
            __props__.__dict__["routing"] = routing
            __props__.__dict__["routing_type"] = routing_type
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallIpv6EhFilter, __self__).__init__(
            'fortios:index/firewallIpv6EhFilter:FirewallIpv6EhFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth: Optional[pulumi.Input[str]] = None,
            dest_opt: Optional[pulumi.Input[str]] = None,
            fragment: Optional[pulumi.Input[str]] = None,
            hdopt_type: Optional[pulumi.Input[int]] = None,
            hop_opt: Optional[pulumi.Input[str]] = None,
            no_next: Optional[pulumi.Input[str]] = None,
            routing: Optional[pulumi.Input[str]] = None,
            routing_type: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallIpv6EhFilter':
        """
        Get an existing FirewallIpv6EhFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth: Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dest_opt: Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] fragment: Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[int] hdopt_type: Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        :param pulumi.Input[str] hop_opt: Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] no_next: Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
        :param pulumi.Input[str] routing: Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
        :param pulumi.Input[int] routing_type: Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallIpv6EhFilterState.__new__(_FirewallIpv6EhFilterState)

        __props__.__dict__["auth"] = auth
        __props__.__dict__["dest_opt"] = dest_opt
        __props__.__dict__["fragment"] = fragment
        __props__.__dict__["hdopt_type"] = hdopt_type
        __props__.__dict__["hop_opt"] = hop_opt
        __props__.__dict__["no_next"] = no_next
        __props__.__dict__["routing"] = routing
        __props__.__dict__["routing_type"] = routing_type
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallIpv6EhFilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def auth(self) -> pulumi.Output[str]:
        """
        Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auth")

    @property
    @pulumi.getter(name="destOpt")
    def dest_opt(self) -> pulumi.Output[str]:
        """
        Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dest_opt")

    @property
    @pulumi.getter
    def fragment(self) -> pulumi.Output[str]:
        """
        Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fragment")

    @property
    @pulumi.getter(name="hdoptType")
    def hdopt_type(self) -> pulumi.Output[int]:
        """
        Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        """
        return pulumi.get(self, "hdopt_type")

    @property
    @pulumi.getter(name="hopOpt")
    def hop_opt(self) -> pulumi.Output[str]:
        """
        Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "hop_opt")

    @property
    @pulumi.getter(name="noNext")
    def no_next(self) -> pulumi.Output[str]:
        """
        Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "no_next")

    @property
    @pulumi.getter
    def routing(self) -> pulumi.Output[str]:
        """
        Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "routing")

    @property
    @pulumi.getter(name="routingType")
    def routing_type(self) -> pulumi.Output[int]:
        """
        Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        """
        return pulumi.get(self, "routing_type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")


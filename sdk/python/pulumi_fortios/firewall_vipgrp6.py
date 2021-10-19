# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FirewallVipgrp6Args', 'FirewallVipgrp6']

@pulumi.input_type
class FirewallVipgrp6Args:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input['FirewallVipgrp6MemberArgs']]],
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallVipgrp6 resource.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallVipgrp6MemberArgs']]] members: Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        :param pulumi.Input[int] color: Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: IPv6 VIP name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "members", members)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallVipgrp6MemberArgs']]]:
        """
        Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallVipgrp6MemberArgs']]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        """
        Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 VIP name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)

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
class _FirewallVipgrp6State:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVipgrp6MemberArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallVipgrp6 resources.
        :param pulumi.Input[int] color: Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallVipgrp6MemberArgs']]] members: Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        :param pulumi.Input[str] name: IPv6 VIP name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        """
        Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVipgrp6MemberArgs']]]]:
        """
        Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVipgrp6MemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 VIP name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)

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


class FirewallVipgrp6(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVipgrp6MemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPv6 virtual IP groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname1 = fortios.FirewallVip6("trname1",
            arp_reply="enable",
            color=0,
            extip="2001:1:1:2::100",
            extport="0-65535",
            fosid=0,
            http_cookie_age=60,
            http_cookie_domain_from_host="disable",
            http_cookie_generation=0,
            http_cookie_share="same-ip",
            http_ip_header="disable",
            http_multiplex="disable",
            https_cookie_secure="disable",
            ldb_method="static",
            mappedip="2001:1:1:2::200",
            mappedport="0-65535",
            max_embryonic_connections=1000,
            outlook_web_access="disable",
            persistence="none",
            portforward="disable",
            protocol="tcp",
            ssl_algorithm="high",
            ssl_client_fallback="enable",
            ssl_client_renegotiation="secure",
            ssl_client_session_state_max=1000,
            ssl_client_session_state_timeout=30,
            ssl_client_session_state_type="both",
            ssl_dh_bits="2048",
            ssl_hpkp="disable",
            ssl_hpkp_age=5184000,
            ssl_hpkp_include_subdomains="disable",
            ssl_hsts="disable",
            ssl_hsts_age=5184000,
            ssl_hsts_include_subdomains="disable",
            ssl_http_location_conversion="disable",
            ssl_http_match_host="enable",
            ssl_max_version="tls-1.2",
            ssl_min_version="tls-1.1",
            ssl_mode="half",
            ssl_pfs="require",
            ssl_send_empty_frags="enable",
            ssl_server_algorithm="client",
            ssl_server_max_version="client",
            ssl_server_min_version="client",
            ssl_server_session_state_max=100,
            ssl_server_session_state_timeout=60,
            ssl_server_session_state_type="both",
            type="static-nat",
            weblogic_server="disable",
            websphere_server="disable")
        trname = fortios.FirewallVipgrp6("trname",
            color=0,
            members=[fortios.FirewallVipgrp6MemberArgs(
                name=trname1.name,
            )])
        ```

        ## Import

        Firewall Vipgrp6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallVipgrp6:FirewallVipgrp6 labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] color: Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVipgrp6MemberArgs']]]] members: Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        :param pulumi.Input[str] name: IPv6 VIP name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallVipgrp6Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPv6 virtual IP groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname1 = fortios.FirewallVip6("trname1",
            arp_reply="enable",
            color=0,
            extip="2001:1:1:2::100",
            extport="0-65535",
            fosid=0,
            http_cookie_age=60,
            http_cookie_domain_from_host="disable",
            http_cookie_generation=0,
            http_cookie_share="same-ip",
            http_ip_header="disable",
            http_multiplex="disable",
            https_cookie_secure="disable",
            ldb_method="static",
            mappedip="2001:1:1:2::200",
            mappedport="0-65535",
            max_embryonic_connections=1000,
            outlook_web_access="disable",
            persistence="none",
            portforward="disable",
            protocol="tcp",
            ssl_algorithm="high",
            ssl_client_fallback="enable",
            ssl_client_renegotiation="secure",
            ssl_client_session_state_max=1000,
            ssl_client_session_state_timeout=30,
            ssl_client_session_state_type="both",
            ssl_dh_bits="2048",
            ssl_hpkp="disable",
            ssl_hpkp_age=5184000,
            ssl_hpkp_include_subdomains="disable",
            ssl_hsts="disable",
            ssl_hsts_age=5184000,
            ssl_hsts_include_subdomains="disable",
            ssl_http_location_conversion="disable",
            ssl_http_match_host="enable",
            ssl_max_version="tls-1.2",
            ssl_min_version="tls-1.1",
            ssl_mode="half",
            ssl_pfs="require",
            ssl_send_empty_frags="enable",
            ssl_server_algorithm="client",
            ssl_server_max_version="client",
            ssl_server_min_version="client",
            ssl_server_session_state_max=100,
            ssl_server_session_state_timeout=60,
            ssl_server_session_state_type="both",
            type="static-nat",
            weblogic_server="disable",
            websphere_server="disable")
        trname = fortios.FirewallVipgrp6("trname",
            color=0,
            members=[fortios.FirewallVipgrp6MemberArgs(
                name=trname1.name,
            )])
        ```

        ## Import

        Firewall Vipgrp6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallVipgrp6:FirewallVipgrp6 labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallVipgrp6Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallVipgrp6Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVipgrp6MemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
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
            __props__ = FirewallVipgrp6Args.__new__(FirewallVipgrp6Args)

            __props__.__dict__["color"] = color
            __props__.__dict__["comments"] = comments
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["uuid"] = uuid
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallVipgrp6, __self__).__init__(
            'fortios:index/firewallVipgrp6:FirewallVipgrp6',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[int]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVipgrp6MemberArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallVipgrp6':
        """
        Get an existing FirewallVipgrp6 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] color: Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVipgrp6MemberArgs']]]] members: Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        :param pulumi.Input[str] name: IPv6 VIP name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallVipgrp6State.__new__(_FirewallVipgrp6State)

        __props__.__dict__["color"] = color
        __props__.__dict__["comments"] = comments
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["uuid"] = uuid
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallVipgrp6(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[int]:
        """
        Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence['outputs.FirewallVipgrp6Member']]:
        """
        Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        IPv6 VIP name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")


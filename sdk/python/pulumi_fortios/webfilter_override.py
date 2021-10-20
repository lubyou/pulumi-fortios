# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WebfilterOverrideArgs', 'WebfilterOverride']

@pulumi.input_type
class WebfilterOverrideArgs:
    def __init__(__self__, *,
                 expires: pulumi.Input[str],
                 new_profile: pulumi.Input[str],
                 old_profile: pulumi.Input[str],
                 user: pulumi.Input[str],
                 fosid: Optional[pulumi.Input[int]] = None,
                 initiator: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user_group: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebfilterOverride resource.
        :param pulumi.Input[str] expires: Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
        :param pulumi.Input[str] new_profile: Name of the new web filter profile used by the override.
        :param pulumi.Input[str] old_profile: Name of the web filter profile which the override applies.
        :param pulumi.Input[str] user: Name of the user which the override applies.
        :param pulumi.Input[int] fosid: Override rule ID.
        :param pulumi.Input[str] initiator: Initiating user of override (read-only setting).
        :param pulumi.Input[str] ip: IPv4 address which the override applies.
        :param pulumi.Input[str] ip6: IPv6 address which the override applies.
        :param pulumi.Input[str] scope: Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
        :param pulumi.Input[str] status: Enable/disable override rule. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] user_group: Specify the user group for which the override applies.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "expires", expires)
        pulumi.set(__self__, "new_profile", new_profile)
        pulumi.set(__self__, "old_profile", old_profile)
        pulumi.set(__self__, "user", user)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if initiator is not None:
            pulumi.set(__self__, "initiator", initiator)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ip6 is not None:
            pulumi.set(__self__, "ip6", ip6)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if user_group is not None:
            pulumi.set(__self__, "user_group", user_group)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def expires(self) -> pulumi.Input[str]:
        """
        Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
        """
        return pulumi.get(self, "expires")

    @expires.setter
    def expires(self, value: pulumi.Input[str]):
        pulumi.set(self, "expires", value)

    @property
    @pulumi.getter(name="newProfile")
    def new_profile(self) -> pulumi.Input[str]:
        """
        Name of the new web filter profile used by the override.
        """
        return pulumi.get(self, "new_profile")

    @new_profile.setter
    def new_profile(self, value: pulumi.Input[str]):
        pulumi.set(self, "new_profile", value)

    @property
    @pulumi.getter(name="oldProfile")
    def old_profile(self) -> pulumi.Input[str]:
        """
        Name of the web filter profile which the override applies.
        """
        return pulumi.get(self, "old_profile")

    @old_profile.setter
    def old_profile(self, value: pulumi.Input[str]):
        pulumi.set(self, "old_profile", value)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        Name of the user which the override applies.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Override rule ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def initiator(self) -> Optional[pulumi.Input[str]]:
        """
        Initiating user of override (read-only setting).
        """
        return pulumi.get(self, "initiator")

    @initiator.setter
    def initiator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initiator", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address which the override applies.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def ip6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 address which the override applies.
        """
        return pulumi.get(self, "ip6")

    @ip6.setter
    def ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip6", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable override rule. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="userGroup")
    def user_group(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the user group for which the override applies.
        """
        return pulumi.get(self, "user_group")

    @user_group.setter
    def user_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_group", value)

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
class _WebfilterOverrideState:
    def __init__(__self__, *,
                 expires: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 initiator: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 new_profile: Optional[pulumi.Input[str]] = None,
                 old_profile: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 user_group: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebfilterOverride resources.
        :param pulumi.Input[str] expires: Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
        :param pulumi.Input[int] fosid: Override rule ID.
        :param pulumi.Input[str] initiator: Initiating user of override (read-only setting).
        :param pulumi.Input[str] ip: IPv4 address which the override applies.
        :param pulumi.Input[str] ip6: IPv6 address which the override applies.
        :param pulumi.Input[str] new_profile: Name of the new web filter profile used by the override.
        :param pulumi.Input[str] old_profile: Name of the web filter profile which the override applies.
        :param pulumi.Input[str] scope: Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
        :param pulumi.Input[str] status: Enable/disable override rule. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] user: Name of the user which the override applies.
        :param pulumi.Input[str] user_group: Specify the user group for which the override applies.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if expires is not None:
            pulumi.set(__self__, "expires", expires)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if initiator is not None:
            pulumi.set(__self__, "initiator", initiator)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ip6 is not None:
            pulumi.set(__self__, "ip6", ip6)
        if new_profile is not None:
            pulumi.set(__self__, "new_profile", new_profile)
        if old_profile is not None:
            pulumi.set(__self__, "old_profile", old_profile)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if user is not None:
            pulumi.set(__self__, "user", user)
        if user_group is not None:
            pulumi.set(__self__, "user_group", user_group)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def expires(self) -> Optional[pulumi.Input[str]]:
        """
        Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
        """
        return pulumi.get(self, "expires")

    @expires.setter
    def expires(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Override rule ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def initiator(self) -> Optional[pulumi.Input[str]]:
        """
        Initiating user of override (read-only setting).
        """
        return pulumi.get(self, "initiator")

    @initiator.setter
    def initiator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initiator", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address which the override applies.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def ip6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 address which the override applies.
        """
        return pulumi.get(self, "ip6")

    @ip6.setter
    def ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip6", value)

    @property
    @pulumi.getter(name="newProfile")
    def new_profile(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the new web filter profile used by the override.
        """
        return pulumi.get(self, "new_profile")

    @new_profile.setter
    def new_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "new_profile", value)

    @property
    @pulumi.getter(name="oldProfile")
    def old_profile(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the web filter profile which the override applies.
        """
        return pulumi.get(self, "old_profile")

    @old_profile.setter
    def old_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "old_profile", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable override rule. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the user which the override applies.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter(name="userGroup")
    def user_group(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the user group for which the override applies.
        """
        return pulumi.get(self, "user_group")

    @user_group.setter
    def user_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_group", value)

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


class WebfilterOverride(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 initiator: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 new_profile: Optional[pulumi.Input[str]] = None,
                 old_profile: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 user_group: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiGuard Web Filter administrative overrides.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WebfilterOverride("trname",
            expires="2021/03/16 19:18:25",
            fosid=1,
            ip="69.101.119.0",
            ip6="4565:7700::",
            new_profile="monitor-all",
            old_profile="default",
            scope="user",
            status="disable",
            user="Eew")
        ```

        ## Import

        Webfilter Override can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/webfilterOverride:WebfilterOverride labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expires: Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
        :param pulumi.Input[int] fosid: Override rule ID.
        :param pulumi.Input[str] initiator: Initiating user of override (read-only setting).
        :param pulumi.Input[str] ip: IPv4 address which the override applies.
        :param pulumi.Input[str] ip6: IPv6 address which the override applies.
        :param pulumi.Input[str] new_profile: Name of the new web filter profile used by the override.
        :param pulumi.Input[str] old_profile: Name of the web filter profile which the override applies.
        :param pulumi.Input[str] scope: Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
        :param pulumi.Input[str] status: Enable/disable override rule. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] user: Name of the user which the override applies.
        :param pulumi.Input[str] user_group: Specify the user group for which the override applies.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebfilterOverrideArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiGuard Web Filter administrative overrides.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WebfilterOverride("trname",
            expires="2021/03/16 19:18:25",
            fosid=1,
            ip="69.101.119.0",
            ip6="4565:7700::",
            new_profile="monitor-all",
            old_profile="default",
            scope="user",
            status="disable",
            user="Eew")
        ```

        ## Import

        Webfilter Override can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/webfilterOverride:WebfilterOverride labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WebfilterOverrideArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebfilterOverrideArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 initiator: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip6: Optional[pulumi.Input[str]] = None,
                 new_profile: Optional[pulumi.Input[str]] = None,
                 old_profile: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 user_group: Optional[pulumi.Input[str]] = None,
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
            __props__ = WebfilterOverrideArgs.__new__(WebfilterOverrideArgs)

            if expires is None and not opts.urn:
                raise TypeError("Missing required property 'expires'")
            __props__.__dict__["expires"] = expires
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["initiator"] = initiator
            __props__.__dict__["ip"] = ip
            __props__.__dict__["ip6"] = ip6
            if new_profile is None and not opts.urn:
                raise TypeError("Missing required property 'new_profile'")
            __props__.__dict__["new_profile"] = new_profile
            if old_profile is None and not opts.urn:
                raise TypeError("Missing required property 'old_profile'")
            __props__.__dict__["old_profile"] = old_profile
            __props__.__dict__["scope"] = scope
            __props__.__dict__["status"] = status
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__.__dict__["user"] = user
            __props__.__dict__["user_group"] = user_group
            __props__.__dict__["vdomparam"] = vdomparam
        super(WebfilterOverride, __self__).__init__(
            'fortios:index/webfilterOverride:WebfilterOverride',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expires: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            initiator: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            ip6: Optional[pulumi.Input[str]] = None,
            new_profile: Optional[pulumi.Input[str]] = None,
            old_profile: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None,
            user_group: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WebfilterOverride':
        """
        Get an existing WebfilterOverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expires: Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
        :param pulumi.Input[int] fosid: Override rule ID.
        :param pulumi.Input[str] initiator: Initiating user of override (read-only setting).
        :param pulumi.Input[str] ip: IPv4 address which the override applies.
        :param pulumi.Input[str] ip6: IPv6 address which the override applies.
        :param pulumi.Input[str] new_profile: Name of the new web filter profile used by the override.
        :param pulumi.Input[str] old_profile: Name of the web filter profile which the override applies.
        :param pulumi.Input[str] scope: Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
        :param pulumi.Input[str] status: Enable/disable override rule. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] user: Name of the user which the override applies.
        :param pulumi.Input[str] user_group: Specify the user group for which the override applies.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebfilterOverrideState.__new__(_WebfilterOverrideState)

        __props__.__dict__["expires"] = expires
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["initiator"] = initiator
        __props__.__dict__["ip"] = ip
        __props__.__dict__["ip6"] = ip6
        __props__.__dict__["new_profile"] = new_profile
        __props__.__dict__["old_profile"] = old_profile
        __props__.__dict__["scope"] = scope
        __props__.__dict__["status"] = status
        __props__.__dict__["user"] = user
        __props__.__dict__["user_group"] = user_group
        __props__.__dict__["vdomparam"] = vdomparam
        return WebfilterOverride(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def expires(self) -> pulumi.Output[str]:
        """
        Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
        """
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        Override rule ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def initiator(self) -> pulumi.Output[str]:
        """
        Initiating user of override (read-only setting).
        """
        return pulumi.get(self, "initiator")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        IPv4 address which the override applies.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def ip6(self) -> pulumi.Output[str]:
        """
        IPv6 address which the override applies.
        """
        return pulumi.get(self, "ip6")

    @property
    @pulumi.getter(name="newProfile")
    def new_profile(self) -> pulumi.Output[str]:
        """
        Name of the new web filter profile used by the override.
        """
        return pulumi.get(self, "new_profile")

    @property
    @pulumi.getter(name="oldProfile")
    def old_profile(self) -> pulumi.Output[str]:
        """
        Name of the web filter profile which the override applies.
        """
        return pulumi.get(self, "old_profile")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable override rule. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        Name of the user which the override applies.
        """
        return pulumi.get(self, "user")

    @property
    @pulumi.getter(name="userGroup")
    def user_group(self) -> pulumi.Output[str]:
        """
        Specify the user group for which the override applies.
        """
        return pulumi.get(self, "user_group")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

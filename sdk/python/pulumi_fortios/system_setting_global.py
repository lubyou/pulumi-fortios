# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemSettingGlobalArgs', 'SystemSettingGlobal']

@pulumi.input_type
class SystemSettingGlobalArgs:
    def __init__(__self__, *,
                 hostname: pulumi.Input[str],
                 admin_scp: Optional[pulumi.Input[str]] = None,
                 admin_sport: Optional[pulumi.Input[str]] = None,
                 admin_ssh_port: Optional[pulumi.Input[str]] = None,
                 admintimeout: Optional[pulumi.Input[str]] = None,
                 timezone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemSettingGlobal resource.
        """
        pulumi.set(__self__, "hostname", hostname)
        if admin_scp is not None:
            pulumi.set(__self__, "admin_scp", admin_scp)
        if admin_sport is not None:
            pulumi.set(__self__, "admin_sport", admin_sport)
        if admin_ssh_port is not None:
            pulumi.set(__self__, "admin_ssh_port", admin_ssh_port)
        if admintimeout is not None:
            pulumi.set(__self__, "admintimeout", admintimeout)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Input[str]:
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="adminScp")
    def admin_scp(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "admin_scp")

    @admin_scp.setter
    def admin_scp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_scp", value)

    @property
    @pulumi.getter(name="adminSport")
    def admin_sport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "admin_sport")

    @admin_sport.setter
    def admin_sport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_sport", value)

    @property
    @pulumi.getter(name="adminSshPort")
    def admin_ssh_port(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "admin_ssh_port")

    @admin_ssh_port.setter
    def admin_ssh_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_ssh_port", value)

    @property
    @pulumi.getter
    def admintimeout(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "admintimeout")

    @admintimeout.setter
    def admintimeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admintimeout", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)


@pulumi.input_type
class _SystemSettingGlobalState:
    def __init__(__self__, *,
                 admin_scp: Optional[pulumi.Input[str]] = None,
                 admin_sport: Optional[pulumi.Input[str]] = None,
                 admin_ssh_port: Optional[pulumi.Input[str]] = None,
                 admintimeout: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 timezone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemSettingGlobal resources.
        """
        if admin_scp is not None:
            pulumi.set(__self__, "admin_scp", admin_scp)
        if admin_sport is not None:
            pulumi.set(__self__, "admin_sport", admin_sport)
        if admin_ssh_port is not None:
            pulumi.set(__self__, "admin_ssh_port", admin_ssh_port)
        if admintimeout is not None:
            pulumi.set(__self__, "admintimeout", admintimeout)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)

    @property
    @pulumi.getter(name="adminScp")
    def admin_scp(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "admin_scp")

    @admin_scp.setter
    def admin_scp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_scp", value)

    @property
    @pulumi.getter(name="adminSport")
    def admin_sport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "admin_sport")

    @admin_sport.setter
    def admin_sport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_sport", value)

    @property
    @pulumi.getter(name="adminSshPort")
    def admin_ssh_port(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "admin_ssh_port")

    @admin_ssh_port.setter
    def admin_ssh_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_ssh_port", value)

    @property
    @pulumi.getter
    def admintimeout(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "admintimeout")

    @admintimeout.setter
    def admintimeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admintimeout", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)


class SystemSettingGlobal(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_scp: Optional[pulumi.Input[str]] = None,
                 admin_sport: Optional[pulumi.Input[str]] = None,
                 admin_ssh_port: Optional[pulumi.Input[str]] = None,
                 admintimeout: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SystemSettingGlobal resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemSettingGlobalArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SystemSettingGlobal resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SystemSettingGlobalArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemSettingGlobalArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_scp: Optional[pulumi.Input[str]] = None,
                 admin_sport: Optional[pulumi.Input[str]] = None,
                 admin_ssh_port: Optional[pulumi.Input[str]] = None,
                 admintimeout: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemSettingGlobalArgs.__new__(SystemSettingGlobalArgs)

            __props__.__dict__["admin_scp"] = admin_scp
            __props__.__dict__["admin_sport"] = admin_sport
            __props__.__dict__["admin_ssh_port"] = admin_ssh_port
            __props__.__dict__["admintimeout"] = admintimeout
            if hostname is None and not opts.urn:
                raise TypeError("Missing required property 'hostname'")
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["timezone"] = timezone
        super(SystemSettingGlobal, __self__).__init__(
            'fortios:index/systemSettingGlobal:SystemSettingGlobal',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_scp: Optional[pulumi.Input[str]] = None,
            admin_sport: Optional[pulumi.Input[str]] = None,
            admin_ssh_port: Optional[pulumi.Input[str]] = None,
            admintimeout: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            timezone: Optional[pulumi.Input[str]] = None) -> 'SystemSettingGlobal':
        """
        Get an existing SystemSettingGlobal resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemSettingGlobalState.__new__(_SystemSettingGlobalState)

        __props__.__dict__["admin_scp"] = admin_scp
        __props__.__dict__["admin_sport"] = admin_sport
        __props__.__dict__["admin_ssh_port"] = admin_ssh_port
        __props__.__dict__["admintimeout"] = admintimeout
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["timezone"] = timezone
        return SystemSettingGlobal(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminScp")
    def admin_scp(self) -> pulumi.Output[str]:
        return pulumi.get(self, "admin_scp")

    @property
    @pulumi.getter(name="adminSport")
    def admin_sport(self) -> pulumi.Output[str]:
        return pulumi.get(self, "admin_sport")

    @property
    @pulumi.getter(name="adminSshPort")
    def admin_ssh_port(self) -> pulumi.Output[str]:
        return pulumi.get(self, "admin_ssh_port")

    @property
    @pulumi.getter
    def admintimeout(self) -> pulumi.Output[str]:
        return pulumi.get(self, "admintimeout")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def timezone(self) -> pulumi.Output[str]:
        return pulumi.get(self, "timezone")


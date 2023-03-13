# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['EndpointControlSettingsArgs', 'EndpointControlSettings']

@pulumi.input_type
class EndpointControlSettingsArgs:
    def __init__(__self__, *,
                 download_custom_link: Optional[pulumi.Input[str]] = None,
                 download_location: Optional[pulumi.Input[str]] = None,
                 forticlient_avdb_update_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_dereg_unsupported_client: Optional[pulumi.Input[str]] = None,
                 forticlient_disconnect_unsupported_client: Optional[pulumi.Input[str]] = None,
                 forticlient_ems_rest_api_call_timeout: Optional[pulumi.Input[int]] = None,
                 forticlient_keepalive_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_offline_grace: Optional[pulumi.Input[str]] = None,
                 forticlient_offline_grace_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_reg_key: Optional[pulumi.Input[str]] = None,
                 forticlient_reg_key_enforce: Optional[pulumi.Input[str]] = None,
                 forticlient_reg_timeout: Optional[pulumi.Input[int]] = None,
                 forticlient_sys_update_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_user_avatar: Optional[pulumi.Input[str]] = None,
                 forticlient_warning_interval: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EndpointControlSettings resource.
        """
        if download_custom_link is not None:
            pulumi.set(__self__, "download_custom_link", download_custom_link)
        if download_location is not None:
            pulumi.set(__self__, "download_location", download_location)
        if forticlient_avdb_update_interval is not None:
            pulumi.set(__self__, "forticlient_avdb_update_interval", forticlient_avdb_update_interval)
        if forticlient_dereg_unsupported_client is not None:
            pulumi.set(__self__, "forticlient_dereg_unsupported_client", forticlient_dereg_unsupported_client)
        if forticlient_disconnect_unsupported_client is not None:
            pulumi.set(__self__, "forticlient_disconnect_unsupported_client", forticlient_disconnect_unsupported_client)
        if forticlient_ems_rest_api_call_timeout is not None:
            pulumi.set(__self__, "forticlient_ems_rest_api_call_timeout", forticlient_ems_rest_api_call_timeout)
        if forticlient_keepalive_interval is not None:
            pulumi.set(__self__, "forticlient_keepalive_interval", forticlient_keepalive_interval)
        if forticlient_offline_grace is not None:
            pulumi.set(__self__, "forticlient_offline_grace", forticlient_offline_grace)
        if forticlient_offline_grace_interval is not None:
            pulumi.set(__self__, "forticlient_offline_grace_interval", forticlient_offline_grace_interval)
        if forticlient_reg_key is not None:
            pulumi.set(__self__, "forticlient_reg_key", forticlient_reg_key)
        if forticlient_reg_key_enforce is not None:
            pulumi.set(__self__, "forticlient_reg_key_enforce", forticlient_reg_key_enforce)
        if forticlient_reg_timeout is not None:
            pulumi.set(__self__, "forticlient_reg_timeout", forticlient_reg_timeout)
        if forticlient_sys_update_interval is not None:
            pulumi.set(__self__, "forticlient_sys_update_interval", forticlient_sys_update_interval)
        if forticlient_user_avatar is not None:
            pulumi.set(__self__, "forticlient_user_avatar", forticlient_user_avatar)
        if forticlient_warning_interval is not None:
            pulumi.set(__self__, "forticlient_warning_interval", forticlient_warning_interval)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="downloadCustomLink")
    def download_custom_link(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "download_custom_link")

    @download_custom_link.setter
    def download_custom_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "download_custom_link", value)

    @property
    @pulumi.getter(name="downloadLocation")
    def download_location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "download_location")

    @download_location.setter
    def download_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "download_location", value)

    @property
    @pulumi.getter(name="forticlientAvdbUpdateInterval")
    def forticlient_avdb_update_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_avdb_update_interval")

    @forticlient_avdb_update_interval.setter
    def forticlient_avdb_update_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_avdb_update_interval", value)

    @property
    @pulumi.getter(name="forticlientDeregUnsupportedClient")
    def forticlient_dereg_unsupported_client(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_dereg_unsupported_client")

    @forticlient_dereg_unsupported_client.setter
    def forticlient_dereg_unsupported_client(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_dereg_unsupported_client", value)

    @property
    @pulumi.getter(name="forticlientDisconnectUnsupportedClient")
    def forticlient_disconnect_unsupported_client(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_disconnect_unsupported_client")

    @forticlient_disconnect_unsupported_client.setter
    def forticlient_disconnect_unsupported_client(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_disconnect_unsupported_client", value)

    @property
    @pulumi.getter(name="forticlientEmsRestApiCallTimeout")
    def forticlient_ems_rest_api_call_timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_ems_rest_api_call_timeout")

    @forticlient_ems_rest_api_call_timeout.setter
    def forticlient_ems_rest_api_call_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_ems_rest_api_call_timeout", value)

    @property
    @pulumi.getter(name="forticlientKeepaliveInterval")
    def forticlient_keepalive_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_keepalive_interval")

    @forticlient_keepalive_interval.setter
    def forticlient_keepalive_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_keepalive_interval", value)

    @property
    @pulumi.getter(name="forticlientOfflineGrace")
    def forticlient_offline_grace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_offline_grace")

    @forticlient_offline_grace.setter
    def forticlient_offline_grace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_offline_grace", value)

    @property
    @pulumi.getter(name="forticlientOfflineGraceInterval")
    def forticlient_offline_grace_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_offline_grace_interval")

    @forticlient_offline_grace_interval.setter
    def forticlient_offline_grace_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_offline_grace_interval", value)

    @property
    @pulumi.getter(name="forticlientRegKey")
    def forticlient_reg_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_reg_key")

    @forticlient_reg_key.setter
    def forticlient_reg_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_reg_key", value)

    @property
    @pulumi.getter(name="forticlientRegKeyEnforce")
    def forticlient_reg_key_enforce(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_reg_key_enforce")

    @forticlient_reg_key_enforce.setter
    def forticlient_reg_key_enforce(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_reg_key_enforce", value)

    @property
    @pulumi.getter(name="forticlientRegTimeout")
    def forticlient_reg_timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_reg_timeout")

    @forticlient_reg_timeout.setter
    def forticlient_reg_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_reg_timeout", value)

    @property
    @pulumi.getter(name="forticlientSysUpdateInterval")
    def forticlient_sys_update_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_sys_update_interval")

    @forticlient_sys_update_interval.setter
    def forticlient_sys_update_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_sys_update_interval", value)

    @property
    @pulumi.getter(name="forticlientUserAvatar")
    def forticlient_user_avatar(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_user_avatar")

    @forticlient_user_avatar.setter
    def forticlient_user_avatar(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_user_avatar", value)

    @property
    @pulumi.getter(name="forticlientWarningInterval")
    def forticlient_warning_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_warning_interval")

    @forticlient_warning_interval.setter
    def forticlient_warning_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_warning_interval", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _EndpointControlSettingsState:
    def __init__(__self__, *,
                 download_custom_link: Optional[pulumi.Input[str]] = None,
                 download_location: Optional[pulumi.Input[str]] = None,
                 forticlient_avdb_update_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_dereg_unsupported_client: Optional[pulumi.Input[str]] = None,
                 forticlient_disconnect_unsupported_client: Optional[pulumi.Input[str]] = None,
                 forticlient_ems_rest_api_call_timeout: Optional[pulumi.Input[int]] = None,
                 forticlient_keepalive_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_offline_grace: Optional[pulumi.Input[str]] = None,
                 forticlient_offline_grace_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_reg_key: Optional[pulumi.Input[str]] = None,
                 forticlient_reg_key_enforce: Optional[pulumi.Input[str]] = None,
                 forticlient_reg_timeout: Optional[pulumi.Input[int]] = None,
                 forticlient_sys_update_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_user_avatar: Optional[pulumi.Input[str]] = None,
                 forticlient_warning_interval: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EndpointControlSettings resources.
        """
        if download_custom_link is not None:
            pulumi.set(__self__, "download_custom_link", download_custom_link)
        if download_location is not None:
            pulumi.set(__self__, "download_location", download_location)
        if forticlient_avdb_update_interval is not None:
            pulumi.set(__self__, "forticlient_avdb_update_interval", forticlient_avdb_update_interval)
        if forticlient_dereg_unsupported_client is not None:
            pulumi.set(__self__, "forticlient_dereg_unsupported_client", forticlient_dereg_unsupported_client)
        if forticlient_disconnect_unsupported_client is not None:
            pulumi.set(__self__, "forticlient_disconnect_unsupported_client", forticlient_disconnect_unsupported_client)
        if forticlient_ems_rest_api_call_timeout is not None:
            pulumi.set(__self__, "forticlient_ems_rest_api_call_timeout", forticlient_ems_rest_api_call_timeout)
        if forticlient_keepalive_interval is not None:
            pulumi.set(__self__, "forticlient_keepalive_interval", forticlient_keepalive_interval)
        if forticlient_offline_grace is not None:
            pulumi.set(__self__, "forticlient_offline_grace", forticlient_offline_grace)
        if forticlient_offline_grace_interval is not None:
            pulumi.set(__self__, "forticlient_offline_grace_interval", forticlient_offline_grace_interval)
        if forticlient_reg_key is not None:
            pulumi.set(__self__, "forticlient_reg_key", forticlient_reg_key)
        if forticlient_reg_key_enforce is not None:
            pulumi.set(__self__, "forticlient_reg_key_enforce", forticlient_reg_key_enforce)
        if forticlient_reg_timeout is not None:
            pulumi.set(__self__, "forticlient_reg_timeout", forticlient_reg_timeout)
        if forticlient_sys_update_interval is not None:
            pulumi.set(__self__, "forticlient_sys_update_interval", forticlient_sys_update_interval)
        if forticlient_user_avatar is not None:
            pulumi.set(__self__, "forticlient_user_avatar", forticlient_user_avatar)
        if forticlient_warning_interval is not None:
            pulumi.set(__self__, "forticlient_warning_interval", forticlient_warning_interval)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="downloadCustomLink")
    def download_custom_link(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "download_custom_link")

    @download_custom_link.setter
    def download_custom_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "download_custom_link", value)

    @property
    @pulumi.getter(name="downloadLocation")
    def download_location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "download_location")

    @download_location.setter
    def download_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "download_location", value)

    @property
    @pulumi.getter(name="forticlientAvdbUpdateInterval")
    def forticlient_avdb_update_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_avdb_update_interval")

    @forticlient_avdb_update_interval.setter
    def forticlient_avdb_update_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_avdb_update_interval", value)

    @property
    @pulumi.getter(name="forticlientDeregUnsupportedClient")
    def forticlient_dereg_unsupported_client(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_dereg_unsupported_client")

    @forticlient_dereg_unsupported_client.setter
    def forticlient_dereg_unsupported_client(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_dereg_unsupported_client", value)

    @property
    @pulumi.getter(name="forticlientDisconnectUnsupportedClient")
    def forticlient_disconnect_unsupported_client(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_disconnect_unsupported_client")

    @forticlient_disconnect_unsupported_client.setter
    def forticlient_disconnect_unsupported_client(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_disconnect_unsupported_client", value)

    @property
    @pulumi.getter(name="forticlientEmsRestApiCallTimeout")
    def forticlient_ems_rest_api_call_timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_ems_rest_api_call_timeout")

    @forticlient_ems_rest_api_call_timeout.setter
    def forticlient_ems_rest_api_call_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_ems_rest_api_call_timeout", value)

    @property
    @pulumi.getter(name="forticlientKeepaliveInterval")
    def forticlient_keepalive_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_keepalive_interval")

    @forticlient_keepalive_interval.setter
    def forticlient_keepalive_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_keepalive_interval", value)

    @property
    @pulumi.getter(name="forticlientOfflineGrace")
    def forticlient_offline_grace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_offline_grace")

    @forticlient_offline_grace.setter
    def forticlient_offline_grace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_offline_grace", value)

    @property
    @pulumi.getter(name="forticlientOfflineGraceInterval")
    def forticlient_offline_grace_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_offline_grace_interval")

    @forticlient_offline_grace_interval.setter
    def forticlient_offline_grace_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_offline_grace_interval", value)

    @property
    @pulumi.getter(name="forticlientRegKey")
    def forticlient_reg_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_reg_key")

    @forticlient_reg_key.setter
    def forticlient_reg_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_reg_key", value)

    @property
    @pulumi.getter(name="forticlientRegKeyEnforce")
    def forticlient_reg_key_enforce(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_reg_key_enforce")

    @forticlient_reg_key_enforce.setter
    def forticlient_reg_key_enforce(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_reg_key_enforce", value)

    @property
    @pulumi.getter(name="forticlientRegTimeout")
    def forticlient_reg_timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_reg_timeout")

    @forticlient_reg_timeout.setter
    def forticlient_reg_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_reg_timeout", value)

    @property
    @pulumi.getter(name="forticlientSysUpdateInterval")
    def forticlient_sys_update_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_sys_update_interval")

    @forticlient_sys_update_interval.setter
    def forticlient_sys_update_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_sys_update_interval", value)

    @property
    @pulumi.getter(name="forticlientUserAvatar")
    def forticlient_user_avatar(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forticlient_user_avatar")

    @forticlient_user_avatar.setter
    def forticlient_user_avatar(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forticlient_user_avatar", value)

    @property
    @pulumi.getter(name="forticlientWarningInterval")
    def forticlient_warning_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forticlient_warning_interval")

    @forticlient_warning_interval.setter
    def forticlient_warning_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forticlient_warning_interval", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class EndpointControlSettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 download_custom_link: Optional[pulumi.Input[str]] = None,
                 download_location: Optional[pulumi.Input[str]] = None,
                 forticlient_avdb_update_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_dereg_unsupported_client: Optional[pulumi.Input[str]] = None,
                 forticlient_disconnect_unsupported_client: Optional[pulumi.Input[str]] = None,
                 forticlient_ems_rest_api_call_timeout: Optional[pulumi.Input[int]] = None,
                 forticlient_keepalive_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_offline_grace: Optional[pulumi.Input[str]] = None,
                 forticlient_offline_grace_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_reg_key: Optional[pulumi.Input[str]] = None,
                 forticlient_reg_key_enforce: Optional[pulumi.Input[str]] = None,
                 forticlient_reg_timeout: Optional[pulumi.Input[int]] = None,
                 forticlient_sys_update_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_user_avatar: Optional[pulumi.Input[str]] = None,
                 forticlient_warning_interval: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a EndpointControlSettings resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EndpointControlSettingsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a EndpointControlSettings resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EndpointControlSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointControlSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 download_custom_link: Optional[pulumi.Input[str]] = None,
                 download_location: Optional[pulumi.Input[str]] = None,
                 forticlient_avdb_update_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_dereg_unsupported_client: Optional[pulumi.Input[str]] = None,
                 forticlient_disconnect_unsupported_client: Optional[pulumi.Input[str]] = None,
                 forticlient_ems_rest_api_call_timeout: Optional[pulumi.Input[int]] = None,
                 forticlient_keepalive_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_offline_grace: Optional[pulumi.Input[str]] = None,
                 forticlient_offline_grace_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_reg_key: Optional[pulumi.Input[str]] = None,
                 forticlient_reg_key_enforce: Optional[pulumi.Input[str]] = None,
                 forticlient_reg_timeout: Optional[pulumi.Input[int]] = None,
                 forticlient_sys_update_interval: Optional[pulumi.Input[int]] = None,
                 forticlient_user_avatar: Optional[pulumi.Input[str]] = None,
                 forticlient_warning_interval: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointControlSettingsArgs.__new__(EndpointControlSettingsArgs)

            __props__.__dict__["download_custom_link"] = download_custom_link
            __props__.__dict__["download_location"] = download_location
            __props__.__dict__["forticlient_avdb_update_interval"] = forticlient_avdb_update_interval
            __props__.__dict__["forticlient_dereg_unsupported_client"] = forticlient_dereg_unsupported_client
            __props__.__dict__["forticlient_disconnect_unsupported_client"] = forticlient_disconnect_unsupported_client
            __props__.__dict__["forticlient_ems_rest_api_call_timeout"] = forticlient_ems_rest_api_call_timeout
            __props__.__dict__["forticlient_keepalive_interval"] = forticlient_keepalive_interval
            __props__.__dict__["forticlient_offline_grace"] = forticlient_offline_grace
            __props__.__dict__["forticlient_offline_grace_interval"] = forticlient_offline_grace_interval
            __props__.__dict__["forticlient_reg_key"] = None if forticlient_reg_key is None else pulumi.Output.secret(forticlient_reg_key)
            __props__.__dict__["forticlient_reg_key_enforce"] = forticlient_reg_key_enforce
            __props__.__dict__["forticlient_reg_timeout"] = forticlient_reg_timeout
            __props__.__dict__["forticlient_sys_update_interval"] = forticlient_sys_update_interval
            __props__.__dict__["forticlient_user_avatar"] = forticlient_user_avatar
            __props__.__dict__["forticlient_warning_interval"] = forticlient_warning_interval
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["forticlientRegKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(EndpointControlSettings, __self__).__init__(
            'fortios:index/endpointControlSettings:EndpointControlSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            download_custom_link: Optional[pulumi.Input[str]] = None,
            download_location: Optional[pulumi.Input[str]] = None,
            forticlient_avdb_update_interval: Optional[pulumi.Input[int]] = None,
            forticlient_dereg_unsupported_client: Optional[pulumi.Input[str]] = None,
            forticlient_disconnect_unsupported_client: Optional[pulumi.Input[str]] = None,
            forticlient_ems_rest_api_call_timeout: Optional[pulumi.Input[int]] = None,
            forticlient_keepalive_interval: Optional[pulumi.Input[int]] = None,
            forticlient_offline_grace: Optional[pulumi.Input[str]] = None,
            forticlient_offline_grace_interval: Optional[pulumi.Input[int]] = None,
            forticlient_reg_key: Optional[pulumi.Input[str]] = None,
            forticlient_reg_key_enforce: Optional[pulumi.Input[str]] = None,
            forticlient_reg_timeout: Optional[pulumi.Input[int]] = None,
            forticlient_sys_update_interval: Optional[pulumi.Input[int]] = None,
            forticlient_user_avatar: Optional[pulumi.Input[str]] = None,
            forticlient_warning_interval: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'EndpointControlSettings':
        """
        Get an existing EndpointControlSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndpointControlSettingsState.__new__(_EndpointControlSettingsState)

        __props__.__dict__["download_custom_link"] = download_custom_link
        __props__.__dict__["download_location"] = download_location
        __props__.__dict__["forticlient_avdb_update_interval"] = forticlient_avdb_update_interval
        __props__.__dict__["forticlient_dereg_unsupported_client"] = forticlient_dereg_unsupported_client
        __props__.__dict__["forticlient_disconnect_unsupported_client"] = forticlient_disconnect_unsupported_client
        __props__.__dict__["forticlient_ems_rest_api_call_timeout"] = forticlient_ems_rest_api_call_timeout
        __props__.__dict__["forticlient_keepalive_interval"] = forticlient_keepalive_interval
        __props__.__dict__["forticlient_offline_grace"] = forticlient_offline_grace
        __props__.__dict__["forticlient_offline_grace_interval"] = forticlient_offline_grace_interval
        __props__.__dict__["forticlient_reg_key"] = forticlient_reg_key
        __props__.__dict__["forticlient_reg_key_enforce"] = forticlient_reg_key_enforce
        __props__.__dict__["forticlient_reg_timeout"] = forticlient_reg_timeout
        __props__.__dict__["forticlient_sys_update_interval"] = forticlient_sys_update_interval
        __props__.__dict__["forticlient_user_avatar"] = forticlient_user_avatar
        __props__.__dict__["forticlient_warning_interval"] = forticlient_warning_interval
        __props__.__dict__["vdomparam"] = vdomparam
        return EndpointControlSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="downloadCustomLink")
    def download_custom_link(self) -> pulumi.Output[str]:
        return pulumi.get(self, "download_custom_link")

    @property
    @pulumi.getter(name="downloadLocation")
    def download_location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "download_location")

    @property
    @pulumi.getter(name="forticlientAvdbUpdateInterval")
    def forticlient_avdb_update_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "forticlient_avdb_update_interval")

    @property
    @pulumi.getter(name="forticlientDeregUnsupportedClient")
    def forticlient_dereg_unsupported_client(self) -> pulumi.Output[str]:
        return pulumi.get(self, "forticlient_dereg_unsupported_client")

    @property
    @pulumi.getter(name="forticlientDisconnectUnsupportedClient")
    def forticlient_disconnect_unsupported_client(self) -> pulumi.Output[str]:
        return pulumi.get(self, "forticlient_disconnect_unsupported_client")

    @property
    @pulumi.getter(name="forticlientEmsRestApiCallTimeout")
    def forticlient_ems_rest_api_call_timeout(self) -> pulumi.Output[int]:
        return pulumi.get(self, "forticlient_ems_rest_api_call_timeout")

    @property
    @pulumi.getter(name="forticlientKeepaliveInterval")
    def forticlient_keepalive_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "forticlient_keepalive_interval")

    @property
    @pulumi.getter(name="forticlientOfflineGrace")
    def forticlient_offline_grace(self) -> pulumi.Output[str]:
        return pulumi.get(self, "forticlient_offline_grace")

    @property
    @pulumi.getter(name="forticlientOfflineGraceInterval")
    def forticlient_offline_grace_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "forticlient_offline_grace_interval")

    @property
    @pulumi.getter(name="forticlientRegKey")
    def forticlient_reg_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "forticlient_reg_key")

    @property
    @pulumi.getter(name="forticlientRegKeyEnforce")
    def forticlient_reg_key_enforce(self) -> pulumi.Output[str]:
        return pulumi.get(self, "forticlient_reg_key_enforce")

    @property
    @pulumi.getter(name="forticlientRegTimeout")
    def forticlient_reg_timeout(self) -> pulumi.Output[int]:
        return pulumi.get(self, "forticlient_reg_timeout")

    @property
    @pulumi.getter(name="forticlientSysUpdateInterval")
    def forticlient_sys_update_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "forticlient_sys_update_interval")

    @property
    @pulumi.getter(name="forticlientUserAvatar")
    def forticlient_user_avatar(self) -> pulumi.Output[str]:
        return pulumi.get(self, "forticlient_user_avatar")

    @property
    @pulumi.getter(name="forticlientWarningInterval")
    def forticlient_warning_interval(self) -> pulumi.Output[int]:
        return pulumi.get(self, "forticlient_warning_interval")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")


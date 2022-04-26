# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SwitchController8021XSettingsArgs', 'SwitchController8021XSettings']

@pulumi.input_type
class SwitchController8021XSettingsArgs:
    def __init__(__self__, *,
                 link_down_auth: Optional[pulumi.Input[str]] = None,
                 max_reauth_attempt: Optional[pulumi.Input[int]] = None,
                 reauth_period: Optional[pulumi.Input[int]] = None,
                 tx_period: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchController8021XSettings resource.
        :param pulumi.Input[str] link_down_auth: Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
        :param pulumi.Input[int] max_reauth_attempt: Maximum number of authentication attempts (0 - 15, default = 3).
        :param pulumi.Input[int] reauth_period: Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
        :param pulumi.Input[int] tx_period: 802.1X Tx period (seconds, default=30).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if link_down_auth is not None:
            pulumi.set(__self__, "link_down_auth", link_down_auth)
        if max_reauth_attempt is not None:
            pulumi.set(__self__, "max_reauth_attempt", max_reauth_attempt)
        if reauth_period is not None:
            pulumi.set(__self__, "reauth_period", reauth_period)
        if tx_period is not None:
            pulumi.set(__self__, "tx_period", tx_period)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="linkDownAuth")
    def link_down_auth(self) -> Optional[pulumi.Input[str]]:
        """
        Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
        """
        return pulumi.get(self, "link_down_auth")

    @link_down_auth.setter
    def link_down_auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_down_auth", value)

    @property
    @pulumi.getter(name="maxReauthAttempt")
    def max_reauth_attempt(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of authentication attempts (0 - 15, default = 3).
        """
        return pulumi.get(self, "max_reauth_attempt")

    @max_reauth_attempt.setter
    def max_reauth_attempt(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_reauth_attempt", value)

    @property
    @pulumi.getter(name="reauthPeriod")
    def reauth_period(self) -> Optional[pulumi.Input[int]]:
        """
        Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
        """
        return pulumi.get(self, "reauth_period")

    @reauth_period.setter
    def reauth_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reauth_period", value)

    @property
    @pulumi.getter(name="txPeriod")
    def tx_period(self) -> Optional[pulumi.Input[int]]:
        """
        802.1X Tx period (seconds, default=30).
        """
        return pulumi.get(self, "tx_period")

    @tx_period.setter
    def tx_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tx_period", value)

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
class _SwitchController8021XSettingsState:
    def __init__(__self__, *,
                 link_down_auth: Optional[pulumi.Input[str]] = None,
                 max_reauth_attempt: Optional[pulumi.Input[int]] = None,
                 reauth_period: Optional[pulumi.Input[int]] = None,
                 tx_period: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchController8021XSettings resources.
        :param pulumi.Input[str] link_down_auth: Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
        :param pulumi.Input[int] max_reauth_attempt: Maximum number of authentication attempts (0 - 15, default = 3).
        :param pulumi.Input[int] reauth_period: Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
        :param pulumi.Input[int] tx_period: 802.1X Tx period (seconds, default=30).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if link_down_auth is not None:
            pulumi.set(__self__, "link_down_auth", link_down_auth)
        if max_reauth_attempt is not None:
            pulumi.set(__self__, "max_reauth_attempt", max_reauth_attempt)
        if reauth_period is not None:
            pulumi.set(__self__, "reauth_period", reauth_period)
        if tx_period is not None:
            pulumi.set(__self__, "tx_period", tx_period)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="linkDownAuth")
    def link_down_auth(self) -> Optional[pulumi.Input[str]]:
        """
        Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
        """
        return pulumi.get(self, "link_down_auth")

    @link_down_auth.setter
    def link_down_auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link_down_auth", value)

    @property
    @pulumi.getter(name="maxReauthAttempt")
    def max_reauth_attempt(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of authentication attempts (0 - 15, default = 3).
        """
        return pulumi.get(self, "max_reauth_attempt")

    @max_reauth_attempt.setter
    def max_reauth_attempt(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_reauth_attempt", value)

    @property
    @pulumi.getter(name="reauthPeriod")
    def reauth_period(self) -> Optional[pulumi.Input[int]]:
        """
        Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
        """
        return pulumi.get(self, "reauth_period")

    @reauth_period.setter
    def reauth_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reauth_period", value)

    @property
    @pulumi.getter(name="txPeriod")
    def tx_period(self) -> Optional[pulumi.Input[int]]:
        """
        802.1X Tx period (seconds, default=30).
        """
        return pulumi.get(self, "tx_period")

    @tx_period.setter
    def tx_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tx_period", value)

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


class SwitchController8021XSettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 link_down_auth: Optional[pulumi.Input[str]] = None,
                 max_reauth_attempt: Optional[pulumi.Input[int]] = None,
                 reauth_period: Optional[pulumi.Input[int]] = None,
                 tx_period: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure global 802.1X settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SwitchController8021XSettings("trname",
            link_down_auth="set-unauth",
            max_reauth_attempt=3,
            reauth_period=12)
        ```

        ## Import

        SwitchController 8021XSettings can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/switchController8021XSettings:SwitchController8021XSettings labelname SwitchController8021XSettings
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/switchController8021XSettings:SwitchController8021XSettings labelname SwitchController8021XSettings
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] link_down_auth: Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
        :param pulumi.Input[int] max_reauth_attempt: Maximum number of authentication attempts (0 - 15, default = 3).
        :param pulumi.Input[int] reauth_period: Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
        :param pulumi.Input[int] tx_period: 802.1X Tx period (seconds, default=30).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchController8021XSettingsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure global 802.1X settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SwitchController8021XSettings("trname",
            link_down_auth="set-unauth",
            max_reauth_attempt=3,
            reauth_period=12)
        ```

        ## Import

        SwitchController 8021XSettings can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/switchController8021XSettings:SwitchController8021XSettings labelname SwitchController8021XSettings
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/switchController8021XSettings:SwitchController8021XSettings labelname SwitchController8021XSettings
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SwitchController8021XSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchController8021XSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 link_down_auth: Optional[pulumi.Input[str]] = None,
                 max_reauth_attempt: Optional[pulumi.Input[int]] = None,
                 reauth_period: Optional[pulumi.Input[int]] = None,
                 tx_period: Optional[pulumi.Input[int]] = None,
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
            __props__ = SwitchController8021XSettingsArgs.__new__(SwitchController8021XSettingsArgs)

            __props__.__dict__["link_down_auth"] = link_down_auth
            __props__.__dict__["max_reauth_attempt"] = max_reauth_attempt
            __props__.__dict__["reauth_period"] = reauth_period
            __props__.__dict__["tx_period"] = tx_period
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchController8021XSettings, __self__).__init__(
            'fortios:index/switchController8021XSettings:SwitchController8021XSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            link_down_auth: Optional[pulumi.Input[str]] = None,
            max_reauth_attempt: Optional[pulumi.Input[int]] = None,
            reauth_period: Optional[pulumi.Input[int]] = None,
            tx_period: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchController8021XSettings':
        """
        Get an existing SwitchController8021XSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] link_down_auth: Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
        :param pulumi.Input[int] max_reauth_attempt: Maximum number of authentication attempts (0 - 15, default = 3).
        :param pulumi.Input[int] reauth_period: Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
        :param pulumi.Input[int] tx_period: 802.1X Tx period (seconds, default=30).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchController8021XSettingsState.__new__(_SwitchController8021XSettingsState)

        __props__.__dict__["link_down_auth"] = link_down_auth
        __props__.__dict__["max_reauth_attempt"] = max_reauth_attempt
        __props__.__dict__["reauth_period"] = reauth_period
        __props__.__dict__["tx_period"] = tx_period
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchController8021XSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="linkDownAuth")
    def link_down_auth(self) -> pulumi.Output[str]:
        """
        Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
        """
        return pulumi.get(self, "link_down_auth")

    @property
    @pulumi.getter(name="maxReauthAttempt")
    def max_reauth_attempt(self) -> pulumi.Output[int]:
        """
        Maximum number of authentication attempts (0 - 15, default = 3).
        """
        return pulumi.get(self, "max_reauth_attempt")

    @property
    @pulumi.getter(name="reauthPeriod")
    def reauth_period(self) -> pulumi.Output[int]:
        """
        Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
        """
        return pulumi.get(self, "reauth_period")

    @property
    @pulumi.getter(name="txPeriod")
    def tx_period(self) -> pulumi.Output[int]:
        """
        802.1X Tx period (seconds, default=30).
        """
        return pulumi.get(self, "tx_period")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")


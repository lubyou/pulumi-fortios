# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LogEventfilterArgs', 'LogEventfilter']

@pulumi.input_type
class LogEventfilterArgs:
    def __init__(__self__, *,
                 cifs: Optional[pulumi.Input[str]] = None,
                 compliance_check: Optional[pulumi.Input[str]] = None,
                 connector: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 event: Optional[pulumi.Input[str]] = None,
                 fortiextender: Optional[pulumi.Input[str]] = None,
                 ha: Optional[pulumi.Input[str]] = None,
                 rest_api: Optional[pulumi.Input[str]] = None,
                 router: Optional[pulumi.Input[str]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 security_rating: Optional[pulumi.Input[str]] = None,
                 switch_controller: Optional[pulumi.Input[str]] = None,
                 system: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vpn: Optional[pulumi.Input[str]] = None,
                 wan_opt: Optional[pulumi.Input[str]] = None,
                 webproxy: Optional[pulumi.Input[str]] = None,
                 wireless_activity: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LogEventfilter resource.
        """
        if cifs is not None:
            pulumi.set(__self__, "cifs", cifs)
        if compliance_check is not None:
            pulumi.set(__self__, "compliance_check", compliance_check)
        if connector is not None:
            pulumi.set(__self__, "connector", connector)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if event is not None:
            pulumi.set(__self__, "event", event)
        if fortiextender is not None:
            pulumi.set(__self__, "fortiextender", fortiextender)
        if ha is not None:
            pulumi.set(__self__, "ha", ha)
        if rest_api is not None:
            pulumi.set(__self__, "rest_api", rest_api)
        if router is not None:
            pulumi.set(__self__, "router", router)
        if sdwan is not None:
            pulumi.set(__self__, "sdwan", sdwan)
        if security_rating is not None:
            pulumi.set(__self__, "security_rating", security_rating)
        if switch_controller is not None:
            pulumi.set(__self__, "switch_controller", switch_controller)
        if system is not None:
            pulumi.set(__self__, "system", system)
        if user is not None:
            pulumi.set(__self__, "user", user)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vpn is not None:
            pulumi.set(__self__, "vpn", vpn)
        if wan_opt is not None:
            pulumi.set(__self__, "wan_opt", wan_opt)
        if webproxy is not None:
            pulumi.set(__self__, "webproxy", webproxy)
        if wireless_activity is not None:
            pulumi.set(__self__, "wireless_activity", wireless_activity)

    @property
    @pulumi.getter
    def cifs(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cifs")

    @cifs.setter
    def cifs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cifs", value)

    @property
    @pulumi.getter(name="complianceCheck")
    def compliance_check(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "compliance_check")

    @compliance_check.setter
    def compliance_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compliance_check", value)

    @property
    @pulumi.getter
    def connector(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "connector")

    @connector.setter
    def connector(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connector", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def event(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "event")

    @event.setter
    def event(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event", value)

    @property
    @pulumi.getter
    def fortiextender(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fortiextender")

    @fortiextender.setter
    def fortiextender(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortiextender", value)

    @property
    @pulumi.getter
    def ha(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ha")

    @ha.setter
    def ha(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ha", value)

    @property
    @pulumi.getter(name="restApi")
    def rest_api(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rest_api")

    @rest_api.setter
    def rest_api(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rest_api", value)

    @property
    @pulumi.getter
    def router(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "router")

    @router.setter
    def router(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "router", value)

    @property
    @pulumi.getter
    def sdwan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sdwan")

    @sdwan.setter
    def sdwan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdwan", value)

    @property
    @pulumi.getter(name="securityRating")
    def security_rating(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "security_rating")

    @security_rating.setter
    def security_rating(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_rating", value)

    @property
    @pulumi.getter(name="switchController")
    def switch_controller(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "switch_controller")

    @switch_controller.setter
    def switch_controller(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch_controller", value)

    @property
    @pulumi.getter
    def system(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "system")

    @system.setter
    def system(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "system", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def vpn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpn")

    @vpn.setter
    def vpn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn", value)

    @property
    @pulumi.getter(name="wanOpt")
    def wan_opt(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "wan_opt")

    @wan_opt.setter
    def wan_opt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wan_opt", value)

    @property
    @pulumi.getter
    def webproxy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "webproxy")

    @webproxy.setter
    def webproxy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "webproxy", value)

    @property
    @pulumi.getter(name="wirelessActivity")
    def wireless_activity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "wireless_activity")

    @wireless_activity.setter
    def wireless_activity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wireless_activity", value)


@pulumi.input_type
class _LogEventfilterState:
    def __init__(__self__, *,
                 cifs: Optional[pulumi.Input[str]] = None,
                 compliance_check: Optional[pulumi.Input[str]] = None,
                 connector: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 event: Optional[pulumi.Input[str]] = None,
                 fortiextender: Optional[pulumi.Input[str]] = None,
                 ha: Optional[pulumi.Input[str]] = None,
                 rest_api: Optional[pulumi.Input[str]] = None,
                 router: Optional[pulumi.Input[str]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 security_rating: Optional[pulumi.Input[str]] = None,
                 switch_controller: Optional[pulumi.Input[str]] = None,
                 system: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vpn: Optional[pulumi.Input[str]] = None,
                 wan_opt: Optional[pulumi.Input[str]] = None,
                 webproxy: Optional[pulumi.Input[str]] = None,
                 wireless_activity: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogEventfilter resources.
        """
        if cifs is not None:
            pulumi.set(__self__, "cifs", cifs)
        if compliance_check is not None:
            pulumi.set(__self__, "compliance_check", compliance_check)
        if connector is not None:
            pulumi.set(__self__, "connector", connector)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if event is not None:
            pulumi.set(__self__, "event", event)
        if fortiextender is not None:
            pulumi.set(__self__, "fortiextender", fortiextender)
        if ha is not None:
            pulumi.set(__self__, "ha", ha)
        if rest_api is not None:
            pulumi.set(__self__, "rest_api", rest_api)
        if router is not None:
            pulumi.set(__self__, "router", router)
        if sdwan is not None:
            pulumi.set(__self__, "sdwan", sdwan)
        if security_rating is not None:
            pulumi.set(__self__, "security_rating", security_rating)
        if switch_controller is not None:
            pulumi.set(__self__, "switch_controller", switch_controller)
        if system is not None:
            pulumi.set(__self__, "system", system)
        if user is not None:
            pulumi.set(__self__, "user", user)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vpn is not None:
            pulumi.set(__self__, "vpn", vpn)
        if wan_opt is not None:
            pulumi.set(__self__, "wan_opt", wan_opt)
        if webproxy is not None:
            pulumi.set(__self__, "webproxy", webproxy)
        if wireless_activity is not None:
            pulumi.set(__self__, "wireless_activity", wireless_activity)

    @property
    @pulumi.getter
    def cifs(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cifs")

    @cifs.setter
    def cifs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cifs", value)

    @property
    @pulumi.getter(name="complianceCheck")
    def compliance_check(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "compliance_check")

    @compliance_check.setter
    def compliance_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compliance_check", value)

    @property
    @pulumi.getter
    def connector(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "connector")

    @connector.setter
    def connector(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connector", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def event(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "event")

    @event.setter
    def event(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event", value)

    @property
    @pulumi.getter
    def fortiextender(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fortiextender")

    @fortiextender.setter
    def fortiextender(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortiextender", value)

    @property
    @pulumi.getter
    def ha(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ha")

    @ha.setter
    def ha(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ha", value)

    @property
    @pulumi.getter(name="restApi")
    def rest_api(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rest_api")

    @rest_api.setter
    def rest_api(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rest_api", value)

    @property
    @pulumi.getter
    def router(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "router")

    @router.setter
    def router(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "router", value)

    @property
    @pulumi.getter
    def sdwan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sdwan")

    @sdwan.setter
    def sdwan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdwan", value)

    @property
    @pulumi.getter(name="securityRating")
    def security_rating(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "security_rating")

    @security_rating.setter
    def security_rating(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_rating", value)

    @property
    @pulumi.getter(name="switchController")
    def switch_controller(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "switch_controller")

    @switch_controller.setter
    def switch_controller(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch_controller", value)

    @property
    @pulumi.getter
    def system(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "system")

    @system.setter
    def system(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "system", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def vpn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpn")

    @vpn.setter
    def vpn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn", value)

    @property
    @pulumi.getter(name="wanOpt")
    def wan_opt(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "wan_opt")

    @wan_opt.setter
    def wan_opt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wan_opt", value)

    @property
    @pulumi.getter
    def webproxy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "webproxy")

    @webproxy.setter
    def webproxy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "webproxy", value)

    @property
    @pulumi.getter(name="wirelessActivity")
    def wireless_activity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "wireless_activity")

    @wireless_activity.setter
    def wireless_activity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wireless_activity", value)


class LogEventfilter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cifs: Optional[pulumi.Input[str]] = None,
                 compliance_check: Optional[pulumi.Input[str]] = None,
                 connector: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 event: Optional[pulumi.Input[str]] = None,
                 fortiextender: Optional[pulumi.Input[str]] = None,
                 ha: Optional[pulumi.Input[str]] = None,
                 rest_api: Optional[pulumi.Input[str]] = None,
                 router: Optional[pulumi.Input[str]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 security_rating: Optional[pulumi.Input[str]] = None,
                 switch_controller: Optional[pulumi.Input[str]] = None,
                 system: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vpn: Optional[pulumi.Input[str]] = None,
                 wan_opt: Optional[pulumi.Input[str]] = None,
                 webproxy: Optional[pulumi.Input[str]] = None,
                 wireless_activity: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LogEventfilter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LogEventfilterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LogEventfilter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LogEventfilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogEventfilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cifs: Optional[pulumi.Input[str]] = None,
                 compliance_check: Optional[pulumi.Input[str]] = None,
                 connector: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 event: Optional[pulumi.Input[str]] = None,
                 fortiextender: Optional[pulumi.Input[str]] = None,
                 ha: Optional[pulumi.Input[str]] = None,
                 rest_api: Optional[pulumi.Input[str]] = None,
                 router: Optional[pulumi.Input[str]] = None,
                 sdwan: Optional[pulumi.Input[str]] = None,
                 security_rating: Optional[pulumi.Input[str]] = None,
                 switch_controller: Optional[pulumi.Input[str]] = None,
                 system: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vpn: Optional[pulumi.Input[str]] = None,
                 wan_opt: Optional[pulumi.Input[str]] = None,
                 webproxy: Optional[pulumi.Input[str]] = None,
                 wireless_activity: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogEventfilterArgs.__new__(LogEventfilterArgs)

            __props__.__dict__["cifs"] = cifs
            __props__.__dict__["compliance_check"] = compliance_check
            __props__.__dict__["connector"] = connector
            __props__.__dict__["endpoint"] = endpoint
            __props__.__dict__["event"] = event
            __props__.__dict__["fortiextender"] = fortiextender
            __props__.__dict__["ha"] = ha
            __props__.__dict__["rest_api"] = rest_api
            __props__.__dict__["router"] = router
            __props__.__dict__["sdwan"] = sdwan
            __props__.__dict__["security_rating"] = security_rating
            __props__.__dict__["switch_controller"] = switch_controller
            __props__.__dict__["system"] = system
            __props__.__dict__["user"] = user
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vpn"] = vpn
            __props__.__dict__["wan_opt"] = wan_opt
            __props__.__dict__["webproxy"] = webproxy
            __props__.__dict__["wireless_activity"] = wireless_activity
        super(LogEventfilter, __self__).__init__(
            'fortios:index/logEventfilter:LogEventfilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cifs: Optional[pulumi.Input[str]] = None,
            compliance_check: Optional[pulumi.Input[str]] = None,
            connector: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            event: Optional[pulumi.Input[str]] = None,
            fortiextender: Optional[pulumi.Input[str]] = None,
            ha: Optional[pulumi.Input[str]] = None,
            rest_api: Optional[pulumi.Input[str]] = None,
            router: Optional[pulumi.Input[str]] = None,
            sdwan: Optional[pulumi.Input[str]] = None,
            security_rating: Optional[pulumi.Input[str]] = None,
            switch_controller: Optional[pulumi.Input[str]] = None,
            system: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vpn: Optional[pulumi.Input[str]] = None,
            wan_opt: Optional[pulumi.Input[str]] = None,
            webproxy: Optional[pulumi.Input[str]] = None,
            wireless_activity: Optional[pulumi.Input[str]] = None) -> 'LogEventfilter':
        """
        Get an existing LogEventfilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogEventfilterState.__new__(_LogEventfilterState)

        __props__.__dict__["cifs"] = cifs
        __props__.__dict__["compliance_check"] = compliance_check
        __props__.__dict__["connector"] = connector
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["event"] = event
        __props__.__dict__["fortiextender"] = fortiextender
        __props__.__dict__["ha"] = ha
        __props__.__dict__["rest_api"] = rest_api
        __props__.__dict__["router"] = router
        __props__.__dict__["sdwan"] = sdwan
        __props__.__dict__["security_rating"] = security_rating
        __props__.__dict__["switch_controller"] = switch_controller
        __props__.__dict__["system"] = system
        __props__.__dict__["user"] = user
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vpn"] = vpn
        __props__.__dict__["wan_opt"] = wan_opt
        __props__.__dict__["webproxy"] = webproxy
        __props__.__dict__["wireless_activity"] = wireless_activity
        return LogEventfilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cifs(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cifs")

    @property
    @pulumi.getter(name="complianceCheck")
    def compliance_check(self) -> pulumi.Output[str]:
        return pulumi.get(self, "compliance_check")

    @property
    @pulumi.getter
    def connector(self) -> pulumi.Output[str]:
        return pulumi.get(self, "connector")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def event(self) -> pulumi.Output[str]:
        return pulumi.get(self, "event")

    @property
    @pulumi.getter
    def fortiextender(self) -> pulumi.Output[str]:
        return pulumi.get(self, "fortiextender")

    @property
    @pulumi.getter
    def ha(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ha")

    @property
    @pulumi.getter(name="restApi")
    def rest_api(self) -> pulumi.Output[str]:
        return pulumi.get(self, "rest_api")

    @property
    @pulumi.getter
    def router(self) -> pulumi.Output[str]:
        return pulumi.get(self, "router")

    @property
    @pulumi.getter
    def sdwan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sdwan")

    @property
    @pulumi.getter(name="securityRating")
    def security_rating(self) -> pulumi.Output[str]:
        return pulumi.get(self, "security_rating")

    @property
    @pulumi.getter(name="switchController")
    def switch_controller(self) -> pulumi.Output[str]:
        return pulumi.get(self, "switch_controller")

    @property
    @pulumi.getter
    def system(self) -> pulumi.Output[str]:
        return pulumi.get(self, "system")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        return pulumi.get(self, "user")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vpn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpn")

    @property
    @pulumi.getter(name="wanOpt")
    def wan_opt(self) -> pulumi.Output[str]:
        return pulumi.get(self, "wan_opt")

    @property
    @pulumi.getter
    def webproxy(self) -> pulumi.Output[str]:
        return pulumi.get(self, "webproxy")

    @property
    @pulumi.getter(name="wirelessActivity")
    def wireless_activity(self) -> pulumi.Output[str]:
        return pulumi.get(self, "wireless_activity")


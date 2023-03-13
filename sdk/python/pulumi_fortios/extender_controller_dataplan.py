# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ExtenderControllerDataplanArgs', 'ExtenderControllerDataplan']

@pulumi.input_type
class ExtenderControllerDataplanArgs:
    def __init__(__self__, *,
                 apn: Optional[pulumi.Input[str]] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 billing_date: Optional[pulumi.Input[int]] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 carrier: Optional[pulumi.Input[str]] = None,
                 iccid: Optional[pulumi.Input[str]] = None,
                 modem_id: Optional[pulumi.Input[str]] = None,
                 monthly_fee: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overage: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 pdn: Optional[pulumi.Input[str]] = None,
                 preferred_subnet: Optional[pulumi.Input[int]] = None,
                 private_network: Optional[pulumi.Input[str]] = None,
                 signal_period: Optional[pulumi.Input[int]] = None,
                 signal_threshold: Optional[pulumi.Input[int]] = None,
                 slot: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ExtenderControllerDataplan resource.
        """
        if apn is not None:
            pulumi.set(__self__, "apn", apn)
        if auth_type is not None:
            pulumi.set(__self__, "auth_type", auth_type)
        if billing_date is not None:
            pulumi.set(__self__, "billing_date", billing_date)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if carrier is not None:
            pulumi.set(__self__, "carrier", carrier)
        if iccid is not None:
            pulumi.set(__self__, "iccid", iccid)
        if modem_id is not None:
            pulumi.set(__self__, "modem_id", modem_id)
        if monthly_fee is not None:
            pulumi.set(__self__, "monthly_fee", monthly_fee)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if overage is not None:
            pulumi.set(__self__, "overage", overage)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if pdn is not None:
            pulumi.set(__self__, "pdn", pdn)
        if preferred_subnet is not None:
            pulumi.set(__self__, "preferred_subnet", preferred_subnet)
        if private_network is not None:
            pulumi.set(__self__, "private_network", private_network)
        if signal_period is not None:
            pulumi.set(__self__, "signal_period", signal_period)
        if signal_threshold is not None:
            pulumi.set(__self__, "signal_threshold", signal_threshold)
        if slot is not None:
            pulumi.set(__self__, "slot", slot)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def apn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "apn")

    @apn.setter
    def apn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "apn", value)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter(name="billingDate")
    def billing_date(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "billing_date")

    @billing_date.setter
    def billing_date(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "billing_date", value)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def carrier(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "carrier")

    @carrier.setter
    def carrier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "carrier", value)

    @property
    @pulumi.getter
    def iccid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "iccid")

    @iccid.setter
    def iccid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iccid", value)

    @property
    @pulumi.getter(name="modemId")
    def modem_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "modem_id")

    @modem_id.setter
    def modem_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modem_id", value)

    @property
    @pulumi.getter(name="monthlyFee")
    def monthly_fee(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "monthly_fee")

    @monthly_fee.setter
    def monthly_fee(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "monthly_fee", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def overage(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "overage")

    @overage.setter
    def overage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "overage", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def pdn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "pdn")

    @pdn.setter
    def pdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pdn", value)

    @property
    @pulumi.getter(name="preferredSubnet")
    def preferred_subnet(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "preferred_subnet")

    @preferred_subnet.setter
    def preferred_subnet(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "preferred_subnet", value)

    @property
    @pulumi.getter(name="privateNetwork")
    def private_network(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "private_network")

    @private_network.setter
    def private_network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_network", value)

    @property
    @pulumi.getter(name="signalPeriod")
    def signal_period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "signal_period")

    @signal_period.setter
    def signal_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "signal_period", value)

    @property
    @pulumi.getter(name="signalThreshold")
    def signal_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "signal_threshold")

    @signal_threshold.setter
    def signal_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "signal_threshold", value)

    @property
    @pulumi.getter
    def slot(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "slot")

    @slot.setter
    def slot(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slot", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _ExtenderControllerDataplanState:
    def __init__(__self__, *,
                 apn: Optional[pulumi.Input[str]] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 billing_date: Optional[pulumi.Input[int]] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 carrier: Optional[pulumi.Input[str]] = None,
                 iccid: Optional[pulumi.Input[str]] = None,
                 modem_id: Optional[pulumi.Input[str]] = None,
                 monthly_fee: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overage: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 pdn: Optional[pulumi.Input[str]] = None,
                 preferred_subnet: Optional[pulumi.Input[int]] = None,
                 private_network: Optional[pulumi.Input[str]] = None,
                 signal_period: Optional[pulumi.Input[int]] = None,
                 signal_threshold: Optional[pulumi.Input[int]] = None,
                 slot: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ExtenderControllerDataplan resources.
        """
        if apn is not None:
            pulumi.set(__self__, "apn", apn)
        if auth_type is not None:
            pulumi.set(__self__, "auth_type", auth_type)
        if billing_date is not None:
            pulumi.set(__self__, "billing_date", billing_date)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if carrier is not None:
            pulumi.set(__self__, "carrier", carrier)
        if iccid is not None:
            pulumi.set(__self__, "iccid", iccid)
        if modem_id is not None:
            pulumi.set(__self__, "modem_id", modem_id)
        if monthly_fee is not None:
            pulumi.set(__self__, "monthly_fee", monthly_fee)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if overage is not None:
            pulumi.set(__self__, "overage", overage)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if pdn is not None:
            pulumi.set(__self__, "pdn", pdn)
        if preferred_subnet is not None:
            pulumi.set(__self__, "preferred_subnet", preferred_subnet)
        if private_network is not None:
            pulumi.set(__self__, "private_network", private_network)
        if signal_period is not None:
            pulumi.set(__self__, "signal_period", signal_period)
        if signal_threshold is not None:
            pulumi.set(__self__, "signal_threshold", signal_threshold)
        if slot is not None:
            pulumi.set(__self__, "slot", slot)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def apn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "apn")

    @apn.setter
    def apn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "apn", value)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter(name="billingDate")
    def billing_date(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "billing_date")

    @billing_date.setter
    def billing_date(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "billing_date", value)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def carrier(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "carrier")

    @carrier.setter
    def carrier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "carrier", value)

    @property
    @pulumi.getter
    def iccid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "iccid")

    @iccid.setter
    def iccid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iccid", value)

    @property
    @pulumi.getter(name="modemId")
    def modem_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "modem_id")

    @modem_id.setter
    def modem_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modem_id", value)

    @property
    @pulumi.getter(name="monthlyFee")
    def monthly_fee(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "monthly_fee")

    @monthly_fee.setter
    def monthly_fee(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "monthly_fee", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def overage(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "overage")

    @overage.setter
    def overage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "overage", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def pdn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "pdn")

    @pdn.setter
    def pdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pdn", value)

    @property
    @pulumi.getter(name="preferredSubnet")
    def preferred_subnet(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "preferred_subnet")

    @preferred_subnet.setter
    def preferred_subnet(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "preferred_subnet", value)

    @property
    @pulumi.getter(name="privateNetwork")
    def private_network(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "private_network")

    @private_network.setter
    def private_network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_network", value)

    @property
    @pulumi.getter(name="signalPeriod")
    def signal_period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "signal_period")

    @signal_period.setter
    def signal_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "signal_period", value)

    @property
    @pulumi.getter(name="signalThreshold")
    def signal_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "signal_threshold")

    @signal_threshold.setter
    def signal_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "signal_threshold", value)

    @property
    @pulumi.getter
    def slot(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "slot")

    @slot.setter
    def slot(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slot", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class ExtenderControllerDataplan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apn: Optional[pulumi.Input[str]] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 billing_date: Optional[pulumi.Input[int]] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 carrier: Optional[pulumi.Input[str]] = None,
                 iccid: Optional[pulumi.Input[str]] = None,
                 modem_id: Optional[pulumi.Input[str]] = None,
                 monthly_fee: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overage: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 pdn: Optional[pulumi.Input[str]] = None,
                 preferred_subnet: Optional[pulumi.Input[int]] = None,
                 private_network: Optional[pulumi.Input[str]] = None,
                 signal_period: Optional[pulumi.Input[int]] = None,
                 signal_threshold: Optional[pulumi.Input[int]] = None,
                 slot: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ExtenderControllerDataplan resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ExtenderControllerDataplanArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ExtenderControllerDataplan resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ExtenderControllerDataplanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExtenderControllerDataplanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apn: Optional[pulumi.Input[str]] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 billing_date: Optional[pulumi.Input[int]] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 carrier: Optional[pulumi.Input[str]] = None,
                 iccid: Optional[pulumi.Input[str]] = None,
                 modem_id: Optional[pulumi.Input[str]] = None,
                 monthly_fee: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overage: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 pdn: Optional[pulumi.Input[str]] = None,
                 preferred_subnet: Optional[pulumi.Input[int]] = None,
                 private_network: Optional[pulumi.Input[str]] = None,
                 signal_period: Optional[pulumi.Input[int]] = None,
                 signal_threshold: Optional[pulumi.Input[int]] = None,
                 slot: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExtenderControllerDataplanArgs.__new__(ExtenderControllerDataplanArgs)

            __props__.__dict__["apn"] = apn
            __props__.__dict__["auth_type"] = auth_type
            __props__.__dict__["billing_date"] = billing_date
            __props__.__dict__["capacity"] = capacity
            __props__.__dict__["carrier"] = carrier
            __props__.__dict__["iccid"] = iccid
            __props__.__dict__["modem_id"] = modem_id
            __props__.__dict__["monthly_fee"] = monthly_fee
            __props__.__dict__["name"] = name
            __props__.__dict__["overage"] = overage
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["pdn"] = pdn
            __props__.__dict__["preferred_subnet"] = preferred_subnet
            __props__.__dict__["private_network"] = private_network
            __props__.__dict__["signal_period"] = signal_period
            __props__.__dict__["signal_threshold"] = signal_threshold
            __props__.__dict__["slot"] = slot
            __props__.__dict__["type"] = type
            __props__.__dict__["username"] = username
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ExtenderControllerDataplan, __self__).__init__(
            'fortios:index/extenderControllerDataplan:ExtenderControllerDataplan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            apn: Optional[pulumi.Input[str]] = None,
            auth_type: Optional[pulumi.Input[str]] = None,
            billing_date: Optional[pulumi.Input[int]] = None,
            capacity: Optional[pulumi.Input[int]] = None,
            carrier: Optional[pulumi.Input[str]] = None,
            iccid: Optional[pulumi.Input[str]] = None,
            modem_id: Optional[pulumi.Input[str]] = None,
            monthly_fee: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            overage: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            pdn: Optional[pulumi.Input[str]] = None,
            preferred_subnet: Optional[pulumi.Input[int]] = None,
            private_network: Optional[pulumi.Input[str]] = None,
            signal_period: Optional[pulumi.Input[int]] = None,
            signal_threshold: Optional[pulumi.Input[int]] = None,
            slot: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'ExtenderControllerDataplan':
        """
        Get an existing ExtenderControllerDataplan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExtenderControllerDataplanState.__new__(_ExtenderControllerDataplanState)

        __props__.__dict__["apn"] = apn
        __props__.__dict__["auth_type"] = auth_type
        __props__.__dict__["billing_date"] = billing_date
        __props__.__dict__["capacity"] = capacity
        __props__.__dict__["carrier"] = carrier
        __props__.__dict__["iccid"] = iccid
        __props__.__dict__["modem_id"] = modem_id
        __props__.__dict__["monthly_fee"] = monthly_fee
        __props__.__dict__["name"] = name
        __props__.__dict__["overage"] = overage
        __props__.__dict__["password"] = password
        __props__.__dict__["pdn"] = pdn
        __props__.__dict__["preferred_subnet"] = preferred_subnet
        __props__.__dict__["private_network"] = private_network
        __props__.__dict__["signal_period"] = signal_period
        __props__.__dict__["signal_threshold"] = signal_threshold
        __props__.__dict__["slot"] = slot
        __props__.__dict__["type"] = type
        __props__.__dict__["username"] = username
        __props__.__dict__["vdomparam"] = vdomparam
        return ExtenderControllerDataplan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def apn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "apn")

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auth_type")

    @property
    @pulumi.getter(name="billingDate")
    def billing_date(self) -> pulumi.Output[int]:
        return pulumi.get(self, "billing_date")

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Output[int]:
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def carrier(self) -> pulumi.Output[str]:
        return pulumi.get(self, "carrier")

    @property
    @pulumi.getter
    def iccid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "iccid")

    @property
    @pulumi.getter(name="modemId")
    def modem_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "modem_id")

    @property
    @pulumi.getter(name="monthlyFee")
    def monthly_fee(self) -> pulumi.Output[int]:
        return pulumi.get(self, "monthly_fee")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def overage(self) -> pulumi.Output[str]:
        return pulumi.get(self, "overage")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def pdn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "pdn")

    @property
    @pulumi.getter(name="preferredSubnet")
    def preferred_subnet(self) -> pulumi.Output[int]:
        return pulumi.get(self, "preferred_subnet")

    @property
    @pulumi.getter(name="privateNetwork")
    def private_network(self) -> pulumi.Output[str]:
        return pulumi.get(self, "private_network")

    @property
    @pulumi.getter(name="signalPeriod")
    def signal_period(self) -> pulumi.Output[int]:
        return pulumi.get(self, "signal_period")

    @property
    @pulumi.getter(name="signalThreshold")
    def signal_threshold(self) -> pulumi.Output[int]:
        return pulumi.get(self, "signal_threshold")

    @property
    @pulumi.getter
    def slot(self) -> pulumi.Output[str]:
        return pulumi.get(self, "slot")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")


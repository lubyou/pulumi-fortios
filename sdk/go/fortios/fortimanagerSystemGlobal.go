// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports modifying system global setting for FortiManager.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFortimanagerSystemGlobal(ctx, "test1", &fortios.FortimanagerSystemGlobalArgs{
// 			AdomMode:            pulumi.String("advanced"),
// 			AdomStatus:          pulumi.String("disable"),
// 			FortianalyzerStatus: pulumi.String("disable"),
// 			Hostname:            pulumi.String("FMG-VM64-test"),
// 			Timezone:            pulumi.String("09"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FortimanagerSystemGlobal struct {
	pulumi.CustomResourceState

	// Adom Mode.
	AdomMode pulumi.StringPtrOutput `pulumi:"adomMode"`
	// Enable/Disable ADOM.
	AdomStatus pulumi.StringPtrOutput `pulumi:"adomStatus"`
	// Enable/Disable FortiAnalyzer feature.
	FortianalyzerStatus pulumi.StringPtrOutput `pulumi:"fortianalyzerStatus"`
	// Hostname.
	Hostname pulumi.StringPtrOutput `pulumi:"hostname"`
	// TimeZone. 00 - (GMT-12:00) Eniwetak, Kwajalein.01 - (GMT-11:00) Midway Island, Samoa.02 - (GMT-10:00) Hawaii.03 - (GMT-9:00) Alaska.04 - (GMT-8:00) Pacific Time (US & Canada).05 - (GMT-7:00) Arizona.06 - (GMT-7:00) Mountain Time (US & Canada).07 - (GMT-6:00) Central America.08 - (GMT-6:00) Central Time (US & Canada).09 - (GMT-6:00) Mexico City.10 - (GMT-6:00) Saskatchewan.11 - (GMT-5:00) Bogota, Lima, Quito.12 - (GMT-5:00) Eastern Time (US & Canada).13 - (GMT-5:00) Indiana (East).14 - (GMT-4:00) Atlantic Time (Canada).15 - (GMT-4:00) La Paz.16 - (GMT-4:00) Santiago.17 - (GMT-3:30) Newfoundland.18 - (GMT-3:00) Brasilia.19 - (GMT-3:00) Buenos Aires, Georgetown.20 - (GMT-3:00) Nuuk (Greenland).21 - (GMT-2:00) Mid-Atlantic.22 - (GMT-1:00) Azores.23 - (GMT-1:00) Cape Verde Is.24 - (GMT) Monrovia.25 - (GMT) Greenwich Mean Time:Dublin, Edinburgh, Lisbon, London.26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna.27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague.28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris.29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb.30 - (GMT+1:00) West Central Africa.31 - (GMT+2:00) Athens, Sofia, Vilnius.32 - (GMT+2:00) Bucharest.33 - (GMT+2:00) Cairo.34 - (GMT+2:00) Harare, Pretoria.35 - (GMT+2:00) Helsinki, Riga,Tallinn.36 - (GMT+2:00) Jerusalem.37 - (GMT+3:00) Baghdad.38 - (GMT+3:00) Kuwait, Riyadh.39 - (GMT+3:00) St.Petersburg, Volgograd.40 - (GMT+3:00) Nairobi.41 - (GMT+3:30) Tehran.42 - (GMT+4:00) Abu Dhabi, Muscat.43 - (GMT+4:00) Baku.44 - (GMT+4:30) Kabul.45 - (GMT+5:00) Ekaterinburg.46 - (GMT+5:00) Islamabad, Karachi,Tashkent.47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi.48 - (GMT+5:45) Kathmandu.49 - (GMT+6:00) Almaty, Novosibirsk.50 - (GMT+6:00) Astana, Dhaka.51 - (GMT+6:00) Sri Jayawardenapura.52 - (GMT+6:30) Rangoon.53 - (GMT+7:00) Bangkok, Hanoi, Jakarta.54 - (GMT+7:00) Krasnoyarsk.55 - (GMT+8:00) Beijing,ChongQing, HongKong,Urumqi.56 - (GMT+8:00) Irkutsk, Ulaanbaatar.57 - (GMT+8:00) Kuala Lumpur, Singapore.58 - (GMT+8:00) Perth.59 - (GMT+8:00) Taipei.60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul.61 - (GMT+9:00) Yakutsk.62 - (GMT+9:30) Adelaide.63 - (GMT+9:30) Darwin.64 - (GMT+10:00) Brisbane.65 - (GMT+10:00) Canberra, Melbourne, Sydney.66 - (GMT+10:00) Guam, Port Moresby.67 - (GMT+10:00) Hobart.68 - (GMT+10:00) Vladivostok.69 - (GMT+11:00) Magadan.70 - (GMT+11:00) Solomon Is., New Caledonia.71 - (GMT+12:00) Auckland, Wellington.72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is.73 - (GMT+13:00) Nuku'alofa.74 - (GMT-4:30) Caracas.75 - (GMT+1:00) Namibia.76 - (GMT-5:00) Brazil-Acre.77 - (GMT-4:00) Brazil-West.78 - (GMT-3:00) Brazil-East.79 - (GMT-2:00) Brazil-DeNoronha.80 - (GMT+14:00) Kiritimati.81 - (GMT-7:00) Baja California Sur, Chihuahua.82 - (GMT+12:45) Chatham Islands.83 - (GMT+3:00) Minsk.84 - (GMT+13:00) Samoa.85 - (GMT+3:00) Istanbul.86 - (GMT-4:00) Paraguay.87 - (GMT) Casablanca.88 - (GMT+3:00) Moscow.89 - (GMT) Greenwich Mean Time.
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
}

// NewFortimanagerSystemGlobal registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemGlobal(ctx *pulumi.Context,
	name string, args *FortimanagerSystemGlobalArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemGlobal, error) {
	if args == nil {
		args = &FortimanagerSystemGlobalArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FortimanagerSystemGlobal
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemGlobal:FortimanagerSystemGlobal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemGlobal gets an existing FortimanagerSystemGlobal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemGlobal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemGlobalState, opts ...pulumi.ResourceOption) (*FortimanagerSystemGlobal, error) {
	var resource FortimanagerSystemGlobal
	err := ctx.ReadResource("fortios:index/fortimanagerSystemGlobal:FortimanagerSystemGlobal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemGlobal resources.
type fortimanagerSystemGlobalState struct {
	// Adom Mode.
	AdomMode *string `pulumi:"adomMode"`
	// Enable/Disable ADOM.
	AdomStatus *string `pulumi:"adomStatus"`
	// Enable/Disable FortiAnalyzer feature.
	FortianalyzerStatus *string `pulumi:"fortianalyzerStatus"`
	// Hostname.
	Hostname *string `pulumi:"hostname"`
	// TimeZone. 00 - (GMT-12:00) Eniwetak, Kwajalein.01 - (GMT-11:00) Midway Island, Samoa.02 - (GMT-10:00) Hawaii.03 - (GMT-9:00) Alaska.04 - (GMT-8:00) Pacific Time (US & Canada).05 - (GMT-7:00) Arizona.06 - (GMT-7:00) Mountain Time (US & Canada).07 - (GMT-6:00) Central America.08 - (GMT-6:00) Central Time (US & Canada).09 - (GMT-6:00) Mexico City.10 - (GMT-6:00) Saskatchewan.11 - (GMT-5:00) Bogota, Lima, Quito.12 - (GMT-5:00) Eastern Time (US & Canada).13 - (GMT-5:00) Indiana (East).14 - (GMT-4:00) Atlantic Time (Canada).15 - (GMT-4:00) La Paz.16 - (GMT-4:00) Santiago.17 - (GMT-3:30) Newfoundland.18 - (GMT-3:00) Brasilia.19 - (GMT-3:00) Buenos Aires, Georgetown.20 - (GMT-3:00) Nuuk (Greenland).21 - (GMT-2:00) Mid-Atlantic.22 - (GMT-1:00) Azores.23 - (GMT-1:00) Cape Verde Is.24 - (GMT) Monrovia.25 - (GMT) Greenwich Mean Time:Dublin, Edinburgh, Lisbon, London.26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna.27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague.28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris.29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb.30 - (GMT+1:00) West Central Africa.31 - (GMT+2:00) Athens, Sofia, Vilnius.32 - (GMT+2:00) Bucharest.33 - (GMT+2:00) Cairo.34 - (GMT+2:00) Harare, Pretoria.35 - (GMT+2:00) Helsinki, Riga,Tallinn.36 - (GMT+2:00) Jerusalem.37 - (GMT+3:00) Baghdad.38 - (GMT+3:00) Kuwait, Riyadh.39 - (GMT+3:00) St.Petersburg, Volgograd.40 - (GMT+3:00) Nairobi.41 - (GMT+3:30) Tehran.42 - (GMT+4:00) Abu Dhabi, Muscat.43 - (GMT+4:00) Baku.44 - (GMT+4:30) Kabul.45 - (GMT+5:00) Ekaterinburg.46 - (GMT+5:00) Islamabad, Karachi,Tashkent.47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi.48 - (GMT+5:45) Kathmandu.49 - (GMT+6:00) Almaty, Novosibirsk.50 - (GMT+6:00) Astana, Dhaka.51 - (GMT+6:00) Sri Jayawardenapura.52 - (GMT+6:30) Rangoon.53 - (GMT+7:00) Bangkok, Hanoi, Jakarta.54 - (GMT+7:00) Krasnoyarsk.55 - (GMT+8:00) Beijing,ChongQing, HongKong,Urumqi.56 - (GMT+8:00) Irkutsk, Ulaanbaatar.57 - (GMT+8:00) Kuala Lumpur, Singapore.58 - (GMT+8:00) Perth.59 - (GMT+8:00) Taipei.60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul.61 - (GMT+9:00) Yakutsk.62 - (GMT+9:30) Adelaide.63 - (GMT+9:30) Darwin.64 - (GMT+10:00) Brisbane.65 - (GMT+10:00) Canberra, Melbourne, Sydney.66 - (GMT+10:00) Guam, Port Moresby.67 - (GMT+10:00) Hobart.68 - (GMT+10:00) Vladivostok.69 - (GMT+11:00) Magadan.70 - (GMT+11:00) Solomon Is., New Caledonia.71 - (GMT+12:00) Auckland, Wellington.72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is.73 - (GMT+13:00) Nuku'alofa.74 - (GMT-4:30) Caracas.75 - (GMT+1:00) Namibia.76 - (GMT-5:00) Brazil-Acre.77 - (GMT-4:00) Brazil-West.78 - (GMT-3:00) Brazil-East.79 - (GMT-2:00) Brazil-DeNoronha.80 - (GMT+14:00) Kiritimati.81 - (GMT-7:00) Baja California Sur, Chihuahua.82 - (GMT+12:45) Chatham Islands.83 - (GMT+3:00) Minsk.84 - (GMT+13:00) Samoa.85 - (GMT+3:00) Istanbul.86 - (GMT-4:00) Paraguay.87 - (GMT) Casablanca.88 - (GMT+3:00) Moscow.89 - (GMT) Greenwich Mean Time.
	Timezone *string `pulumi:"timezone"`
}

type FortimanagerSystemGlobalState struct {
	// Adom Mode.
	AdomMode pulumi.StringPtrInput
	// Enable/Disable ADOM.
	AdomStatus pulumi.StringPtrInput
	// Enable/Disable FortiAnalyzer feature.
	FortianalyzerStatus pulumi.StringPtrInput
	// Hostname.
	Hostname pulumi.StringPtrInput
	// TimeZone. 00 - (GMT-12:00) Eniwetak, Kwajalein.01 - (GMT-11:00) Midway Island, Samoa.02 - (GMT-10:00) Hawaii.03 - (GMT-9:00) Alaska.04 - (GMT-8:00) Pacific Time (US & Canada).05 - (GMT-7:00) Arizona.06 - (GMT-7:00) Mountain Time (US & Canada).07 - (GMT-6:00) Central America.08 - (GMT-6:00) Central Time (US & Canada).09 - (GMT-6:00) Mexico City.10 - (GMT-6:00) Saskatchewan.11 - (GMT-5:00) Bogota, Lima, Quito.12 - (GMT-5:00) Eastern Time (US & Canada).13 - (GMT-5:00) Indiana (East).14 - (GMT-4:00) Atlantic Time (Canada).15 - (GMT-4:00) La Paz.16 - (GMT-4:00) Santiago.17 - (GMT-3:30) Newfoundland.18 - (GMT-3:00) Brasilia.19 - (GMT-3:00) Buenos Aires, Georgetown.20 - (GMT-3:00) Nuuk (Greenland).21 - (GMT-2:00) Mid-Atlantic.22 - (GMT-1:00) Azores.23 - (GMT-1:00) Cape Verde Is.24 - (GMT) Monrovia.25 - (GMT) Greenwich Mean Time:Dublin, Edinburgh, Lisbon, London.26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna.27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague.28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris.29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb.30 - (GMT+1:00) West Central Africa.31 - (GMT+2:00) Athens, Sofia, Vilnius.32 - (GMT+2:00) Bucharest.33 - (GMT+2:00) Cairo.34 - (GMT+2:00) Harare, Pretoria.35 - (GMT+2:00) Helsinki, Riga,Tallinn.36 - (GMT+2:00) Jerusalem.37 - (GMT+3:00) Baghdad.38 - (GMT+3:00) Kuwait, Riyadh.39 - (GMT+3:00) St.Petersburg, Volgograd.40 - (GMT+3:00) Nairobi.41 - (GMT+3:30) Tehran.42 - (GMT+4:00) Abu Dhabi, Muscat.43 - (GMT+4:00) Baku.44 - (GMT+4:30) Kabul.45 - (GMT+5:00) Ekaterinburg.46 - (GMT+5:00) Islamabad, Karachi,Tashkent.47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi.48 - (GMT+5:45) Kathmandu.49 - (GMT+6:00) Almaty, Novosibirsk.50 - (GMT+6:00) Astana, Dhaka.51 - (GMT+6:00) Sri Jayawardenapura.52 - (GMT+6:30) Rangoon.53 - (GMT+7:00) Bangkok, Hanoi, Jakarta.54 - (GMT+7:00) Krasnoyarsk.55 - (GMT+8:00) Beijing,ChongQing, HongKong,Urumqi.56 - (GMT+8:00) Irkutsk, Ulaanbaatar.57 - (GMT+8:00) Kuala Lumpur, Singapore.58 - (GMT+8:00) Perth.59 - (GMT+8:00) Taipei.60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul.61 - (GMT+9:00) Yakutsk.62 - (GMT+9:30) Adelaide.63 - (GMT+9:30) Darwin.64 - (GMT+10:00) Brisbane.65 - (GMT+10:00) Canberra, Melbourne, Sydney.66 - (GMT+10:00) Guam, Port Moresby.67 - (GMT+10:00) Hobart.68 - (GMT+10:00) Vladivostok.69 - (GMT+11:00) Magadan.70 - (GMT+11:00) Solomon Is., New Caledonia.71 - (GMT+12:00) Auckland, Wellington.72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is.73 - (GMT+13:00) Nuku'alofa.74 - (GMT-4:30) Caracas.75 - (GMT+1:00) Namibia.76 - (GMT-5:00) Brazil-Acre.77 - (GMT-4:00) Brazil-West.78 - (GMT-3:00) Brazil-East.79 - (GMT-2:00) Brazil-DeNoronha.80 - (GMT+14:00) Kiritimati.81 - (GMT-7:00) Baja California Sur, Chihuahua.82 - (GMT+12:45) Chatham Islands.83 - (GMT+3:00) Minsk.84 - (GMT+13:00) Samoa.85 - (GMT+3:00) Istanbul.86 - (GMT-4:00) Paraguay.87 - (GMT) Casablanca.88 - (GMT+3:00) Moscow.89 - (GMT) Greenwich Mean Time.
	Timezone pulumi.StringPtrInput
}

func (FortimanagerSystemGlobalState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemGlobalState)(nil)).Elem()
}

type fortimanagerSystemGlobalArgs struct {
	// Adom Mode.
	AdomMode *string `pulumi:"adomMode"`
	// Enable/Disable ADOM.
	AdomStatus *string `pulumi:"adomStatus"`
	// Enable/Disable FortiAnalyzer feature.
	FortianalyzerStatus *string `pulumi:"fortianalyzerStatus"`
	// Hostname.
	Hostname *string `pulumi:"hostname"`
	// TimeZone. 00 - (GMT-12:00) Eniwetak, Kwajalein.01 - (GMT-11:00) Midway Island, Samoa.02 - (GMT-10:00) Hawaii.03 - (GMT-9:00) Alaska.04 - (GMT-8:00) Pacific Time (US & Canada).05 - (GMT-7:00) Arizona.06 - (GMT-7:00) Mountain Time (US & Canada).07 - (GMT-6:00) Central America.08 - (GMT-6:00) Central Time (US & Canada).09 - (GMT-6:00) Mexico City.10 - (GMT-6:00) Saskatchewan.11 - (GMT-5:00) Bogota, Lima, Quito.12 - (GMT-5:00) Eastern Time (US & Canada).13 - (GMT-5:00) Indiana (East).14 - (GMT-4:00) Atlantic Time (Canada).15 - (GMT-4:00) La Paz.16 - (GMT-4:00) Santiago.17 - (GMT-3:30) Newfoundland.18 - (GMT-3:00) Brasilia.19 - (GMT-3:00) Buenos Aires, Georgetown.20 - (GMT-3:00) Nuuk (Greenland).21 - (GMT-2:00) Mid-Atlantic.22 - (GMT-1:00) Azores.23 - (GMT-1:00) Cape Verde Is.24 - (GMT) Monrovia.25 - (GMT) Greenwich Mean Time:Dublin, Edinburgh, Lisbon, London.26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna.27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague.28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris.29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb.30 - (GMT+1:00) West Central Africa.31 - (GMT+2:00) Athens, Sofia, Vilnius.32 - (GMT+2:00) Bucharest.33 - (GMT+2:00) Cairo.34 - (GMT+2:00) Harare, Pretoria.35 - (GMT+2:00) Helsinki, Riga,Tallinn.36 - (GMT+2:00) Jerusalem.37 - (GMT+3:00) Baghdad.38 - (GMT+3:00) Kuwait, Riyadh.39 - (GMT+3:00) St.Petersburg, Volgograd.40 - (GMT+3:00) Nairobi.41 - (GMT+3:30) Tehran.42 - (GMT+4:00) Abu Dhabi, Muscat.43 - (GMT+4:00) Baku.44 - (GMT+4:30) Kabul.45 - (GMT+5:00) Ekaterinburg.46 - (GMT+5:00) Islamabad, Karachi,Tashkent.47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi.48 - (GMT+5:45) Kathmandu.49 - (GMT+6:00) Almaty, Novosibirsk.50 - (GMT+6:00) Astana, Dhaka.51 - (GMT+6:00) Sri Jayawardenapura.52 - (GMT+6:30) Rangoon.53 - (GMT+7:00) Bangkok, Hanoi, Jakarta.54 - (GMT+7:00) Krasnoyarsk.55 - (GMT+8:00) Beijing,ChongQing, HongKong,Urumqi.56 - (GMT+8:00) Irkutsk, Ulaanbaatar.57 - (GMT+8:00) Kuala Lumpur, Singapore.58 - (GMT+8:00) Perth.59 - (GMT+8:00) Taipei.60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul.61 - (GMT+9:00) Yakutsk.62 - (GMT+9:30) Adelaide.63 - (GMT+9:30) Darwin.64 - (GMT+10:00) Brisbane.65 - (GMT+10:00) Canberra, Melbourne, Sydney.66 - (GMT+10:00) Guam, Port Moresby.67 - (GMT+10:00) Hobart.68 - (GMT+10:00) Vladivostok.69 - (GMT+11:00) Magadan.70 - (GMT+11:00) Solomon Is., New Caledonia.71 - (GMT+12:00) Auckland, Wellington.72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is.73 - (GMT+13:00) Nuku'alofa.74 - (GMT-4:30) Caracas.75 - (GMT+1:00) Namibia.76 - (GMT-5:00) Brazil-Acre.77 - (GMT-4:00) Brazil-West.78 - (GMT-3:00) Brazil-East.79 - (GMT-2:00) Brazil-DeNoronha.80 - (GMT+14:00) Kiritimati.81 - (GMT-7:00) Baja California Sur, Chihuahua.82 - (GMT+12:45) Chatham Islands.83 - (GMT+3:00) Minsk.84 - (GMT+13:00) Samoa.85 - (GMT+3:00) Istanbul.86 - (GMT-4:00) Paraguay.87 - (GMT) Casablanca.88 - (GMT+3:00) Moscow.89 - (GMT) Greenwich Mean Time.
	Timezone *string `pulumi:"timezone"`
}

// The set of arguments for constructing a FortimanagerSystemGlobal resource.
type FortimanagerSystemGlobalArgs struct {
	// Adom Mode.
	AdomMode pulumi.StringPtrInput
	// Enable/Disable ADOM.
	AdomStatus pulumi.StringPtrInput
	// Enable/Disable FortiAnalyzer feature.
	FortianalyzerStatus pulumi.StringPtrInput
	// Hostname.
	Hostname pulumi.StringPtrInput
	// TimeZone. 00 - (GMT-12:00) Eniwetak, Kwajalein.01 - (GMT-11:00) Midway Island, Samoa.02 - (GMT-10:00) Hawaii.03 - (GMT-9:00) Alaska.04 - (GMT-8:00) Pacific Time (US & Canada).05 - (GMT-7:00) Arizona.06 - (GMT-7:00) Mountain Time (US & Canada).07 - (GMT-6:00) Central America.08 - (GMT-6:00) Central Time (US & Canada).09 - (GMT-6:00) Mexico City.10 - (GMT-6:00) Saskatchewan.11 - (GMT-5:00) Bogota, Lima, Quito.12 - (GMT-5:00) Eastern Time (US & Canada).13 - (GMT-5:00) Indiana (East).14 - (GMT-4:00) Atlantic Time (Canada).15 - (GMT-4:00) La Paz.16 - (GMT-4:00) Santiago.17 - (GMT-3:30) Newfoundland.18 - (GMT-3:00) Brasilia.19 - (GMT-3:00) Buenos Aires, Georgetown.20 - (GMT-3:00) Nuuk (Greenland).21 - (GMT-2:00) Mid-Atlantic.22 - (GMT-1:00) Azores.23 - (GMT-1:00) Cape Verde Is.24 - (GMT) Monrovia.25 - (GMT) Greenwich Mean Time:Dublin, Edinburgh, Lisbon, London.26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna.27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague.28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris.29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb.30 - (GMT+1:00) West Central Africa.31 - (GMT+2:00) Athens, Sofia, Vilnius.32 - (GMT+2:00) Bucharest.33 - (GMT+2:00) Cairo.34 - (GMT+2:00) Harare, Pretoria.35 - (GMT+2:00) Helsinki, Riga,Tallinn.36 - (GMT+2:00) Jerusalem.37 - (GMT+3:00) Baghdad.38 - (GMT+3:00) Kuwait, Riyadh.39 - (GMT+3:00) St.Petersburg, Volgograd.40 - (GMT+3:00) Nairobi.41 - (GMT+3:30) Tehran.42 - (GMT+4:00) Abu Dhabi, Muscat.43 - (GMT+4:00) Baku.44 - (GMT+4:30) Kabul.45 - (GMT+5:00) Ekaterinburg.46 - (GMT+5:00) Islamabad, Karachi,Tashkent.47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi.48 - (GMT+5:45) Kathmandu.49 - (GMT+6:00) Almaty, Novosibirsk.50 - (GMT+6:00) Astana, Dhaka.51 - (GMT+6:00) Sri Jayawardenapura.52 - (GMT+6:30) Rangoon.53 - (GMT+7:00) Bangkok, Hanoi, Jakarta.54 - (GMT+7:00) Krasnoyarsk.55 - (GMT+8:00) Beijing,ChongQing, HongKong,Urumqi.56 - (GMT+8:00) Irkutsk, Ulaanbaatar.57 - (GMT+8:00) Kuala Lumpur, Singapore.58 - (GMT+8:00) Perth.59 - (GMT+8:00) Taipei.60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul.61 - (GMT+9:00) Yakutsk.62 - (GMT+9:30) Adelaide.63 - (GMT+9:30) Darwin.64 - (GMT+10:00) Brisbane.65 - (GMT+10:00) Canberra, Melbourne, Sydney.66 - (GMT+10:00) Guam, Port Moresby.67 - (GMT+10:00) Hobart.68 - (GMT+10:00) Vladivostok.69 - (GMT+11:00) Magadan.70 - (GMT+11:00) Solomon Is., New Caledonia.71 - (GMT+12:00) Auckland, Wellington.72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is.73 - (GMT+13:00) Nuku'alofa.74 - (GMT-4:30) Caracas.75 - (GMT+1:00) Namibia.76 - (GMT-5:00) Brazil-Acre.77 - (GMT-4:00) Brazil-West.78 - (GMT-3:00) Brazil-East.79 - (GMT-2:00) Brazil-DeNoronha.80 - (GMT+14:00) Kiritimati.81 - (GMT-7:00) Baja California Sur, Chihuahua.82 - (GMT+12:45) Chatham Islands.83 - (GMT+3:00) Minsk.84 - (GMT+13:00) Samoa.85 - (GMT+3:00) Istanbul.86 - (GMT-4:00) Paraguay.87 - (GMT) Casablanca.88 - (GMT+3:00) Moscow.89 - (GMT) Greenwich Mean Time.
	Timezone pulumi.StringPtrInput
}

func (FortimanagerSystemGlobalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemGlobalArgs)(nil)).Elem()
}

type FortimanagerSystemGlobalInput interface {
	pulumi.Input

	ToFortimanagerSystemGlobalOutput() FortimanagerSystemGlobalOutput
	ToFortimanagerSystemGlobalOutputWithContext(ctx context.Context) FortimanagerSystemGlobalOutput
}

func (*FortimanagerSystemGlobal) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemGlobal)(nil)).Elem()
}

func (i *FortimanagerSystemGlobal) ToFortimanagerSystemGlobalOutput() FortimanagerSystemGlobalOutput {
	return i.ToFortimanagerSystemGlobalOutputWithContext(context.Background())
}

func (i *FortimanagerSystemGlobal) ToFortimanagerSystemGlobalOutputWithContext(ctx context.Context) FortimanagerSystemGlobalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemGlobalOutput)
}

// FortimanagerSystemGlobalArrayInput is an input type that accepts FortimanagerSystemGlobalArray and FortimanagerSystemGlobalArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemGlobalArrayInput` via:
//
//          FortimanagerSystemGlobalArray{ FortimanagerSystemGlobalArgs{...} }
type FortimanagerSystemGlobalArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemGlobalArrayOutput() FortimanagerSystemGlobalArrayOutput
	ToFortimanagerSystemGlobalArrayOutputWithContext(context.Context) FortimanagerSystemGlobalArrayOutput
}

type FortimanagerSystemGlobalArray []FortimanagerSystemGlobalInput

func (FortimanagerSystemGlobalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemGlobal)(nil)).Elem()
}

func (i FortimanagerSystemGlobalArray) ToFortimanagerSystemGlobalArrayOutput() FortimanagerSystemGlobalArrayOutput {
	return i.ToFortimanagerSystemGlobalArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemGlobalArray) ToFortimanagerSystemGlobalArrayOutputWithContext(ctx context.Context) FortimanagerSystemGlobalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemGlobalArrayOutput)
}

// FortimanagerSystemGlobalMapInput is an input type that accepts FortimanagerSystemGlobalMap and FortimanagerSystemGlobalMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemGlobalMapInput` via:
//
//          FortimanagerSystemGlobalMap{ "key": FortimanagerSystemGlobalArgs{...} }
type FortimanagerSystemGlobalMapInput interface {
	pulumi.Input

	ToFortimanagerSystemGlobalMapOutput() FortimanagerSystemGlobalMapOutput
	ToFortimanagerSystemGlobalMapOutputWithContext(context.Context) FortimanagerSystemGlobalMapOutput
}

type FortimanagerSystemGlobalMap map[string]FortimanagerSystemGlobalInput

func (FortimanagerSystemGlobalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemGlobal)(nil)).Elem()
}

func (i FortimanagerSystemGlobalMap) ToFortimanagerSystemGlobalMapOutput() FortimanagerSystemGlobalMapOutput {
	return i.ToFortimanagerSystemGlobalMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemGlobalMap) ToFortimanagerSystemGlobalMapOutputWithContext(ctx context.Context) FortimanagerSystemGlobalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemGlobalMapOutput)
}

type FortimanagerSystemGlobalOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemGlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemGlobal)(nil)).Elem()
}

func (o FortimanagerSystemGlobalOutput) ToFortimanagerSystemGlobalOutput() FortimanagerSystemGlobalOutput {
	return o
}

func (o FortimanagerSystemGlobalOutput) ToFortimanagerSystemGlobalOutputWithContext(ctx context.Context) FortimanagerSystemGlobalOutput {
	return o
}

type FortimanagerSystemGlobalArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemGlobalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemGlobal)(nil)).Elem()
}

func (o FortimanagerSystemGlobalArrayOutput) ToFortimanagerSystemGlobalArrayOutput() FortimanagerSystemGlobalArrayOutput {
	return o
}

func (o FortimanagerSystemGlobalArrayOutput) ToFortimanagerSystemGlobalArrayOutputWithContext(ctx context.Context) FortimanagerSystemGlobalArrayOutput {
	return o
}

func (o FortimanagerSystemGlobalArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemGlobalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerSystemGlobal {
		return vs[0].([]*FortimanagerSystemGlobal)[vs[1].(int)]
	}).(FortimanagerSystemGlobalOutput)
}

type FortimanagerSystemGlobalMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemGlobalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemGlobal)(nil)).Elem()
}

func (o FortimanagerSystemGlobalMapOutput) ToFortimanagerSystemGlobalMapOutput() FortimanagerSystemGlobalMapOutput {
	return o
}

func (o FortimanagerSystemGlobalMapOutput) ToFortimanagerSystemGlobalMapOutputWithContext(ctx context.Context) FortimanagerSystemGlobalMapOutput {
	return o
}

func (o FortimanagerSystemGlobalMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemGlobalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerSystemGlobal {
		return vs[0].(map[string]*FortimanagerSystemGlobal)[vs[1].(string)]
	}).(FortimanagerSystemGlobalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemGlobalInput)(nil)).Elem(), &FortimanagerSystemGlobal{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemGlobalArrayInput)(nil)).Elem(), FortimanagerSystemGlobalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemGlobalMapInput)(nil)).Elem(), FortimanagerSystemGlobalMap{})
	pulumi.RegisterOutputType(FortimanagerSystemGlobalOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemGlobalArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemGlobalMapOutput{})
}

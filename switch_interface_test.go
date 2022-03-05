package metalcloud

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"

	. "github.com/onsi/gomega"
)

func TestSwitchInterfaceUnmarshalTestWithEmptyStuff(t *testing.T) {
	RegisterTestingT(t)

	var obj map[string]searchResultResponseWrapperForSwitchInterfaces
	err := json.Unmarshal([]byte(_switchInterfaceFixture), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())

	r := obj["_switch_interfaces"].Rows[0]
	Expect(r.NetworkEquipmentID).To(Equal(7))

}

func TestSwitchInterfaceUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var obj map[string]searchResultResponseWrapperForSwitchInterfaces
	err := json.Unmarshal([]byte(_switchInterfaceFixture2), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())

	r := obj["_switch_interfaces"].Rows[2]
	Expect(r.NetworkEquipmentID).To(Equal(2))
	Expect(r.IP[0][1]).To(Equal("192.168.78.4"))

}

func TestSwitchInterfaceMarshalAsYamlTest(t *testing.T) {
	RegisterTestingT(t)

	obj := SwitchInterfaceSearchResult{
		ServerID:           100,
		NetworkEquipmentID: 100,
		NetworkType:        []string{"WAN"},
	}

	dcBytes, err := yaml.Marshal(obj)

	Expect(dcBytes).NotTo(BeNil())
	Expect(err).To(BeNil())

}

const _switchInterfaceFixture = `{
        "_switch_interfaces": {
            "duration_milliseconds": 0.07028317451477051,
            "rows": [
                {
                    "network_equipment_interface_id": 80,
                    "network_equipment_identifier_string": "dtsy1sppdal001",
                    "network_equipment_id": 7,
                    "network_equipment_interface_identifier_string": "ethernet1\/1\/5",
                    "server_interface_mac_address": "b4:96:91:be:bb:78",
                    "server_interface_index": 0,
                    "server_interface_capacity_mbps": 25000,
                    "ip_human_readable": [
                        [
                            null
                        ]
                    ],
                    "subnet_range_start_human_readable": [
                        [
                            null
                        ]
                    ],
                    "network_type": [
                        null
                    ],
                    "network_id": [
                        null
                    ],
                    "subnet_prefix_size": [
                        [
                            null
                        ]
                    ],
                    "subnet_pool_id": [
                        [
                            null
                        ]
                    ],
                    "instance_interface_id": [
                        null
                    ],
                    "ip_id": [
                        [
                            null
                        ]
                    ]
                },
                {
                    "network_equipment_interface_id": 81,
                    "network_equipment_identifier_string": "dtsy1sppdal002",
                    "network_equipment_id": 8,
                    "network_equipment_interface_identifier_string": "ethernet1\/1\/5",
                    "server_interface_mac_address": "b4:96:91:be:bb:79",
                    "server_interface_index": 1,
                    "server_interface_capacity_mbps": 25000,
                    "ip_human_readable": [
                        [
                            null
                        ]
                    ],
                    "subnet_range_start_human_readable": [
                        [
                            null
                        ]
                    ],
                    "network_type": [
                        null
                    ],
                    "network_id": [
                        null
                    ],
                    "subnet_prefix_size": [
                        [
                            null
                        ]
                    ],
                    "subnet_pool_id": [
                        [
                            null
                        ]
                    ],
                    "instance_interface_id": [
                        null
                    ],
                    "ip_id": [
                        [
                            null
                        ]
                    ]
                },
                {
                    "network_equipment_interface_id": 82,
                    "network_equipment_identifier_string": "dtsy1sppstl005",
                    "network_equipment_id": 9,
                    "network_equipment_interface_identifier_string": "ethernet1\/1\/5",
                    "server_interface_mac_address": "b4:96:91:c3:82:6e",
                    "server_interface_index": 2,
                    "server_interface_capacity_mbps": 25000,
                    "ip_human_readable": [
                        [
                            null
                        ]
                    ],
                    "subnet_range_start_human_readable": [
                        [
                            null
                        ]
                    ],
                    "network_type": [
                        null
                    ],
                    "network_id": [
                        null
                    ],
                    "subnet_prefix_size": [
                        [
                            null
                        ]
                    ],
                    "subnet_pool_id": [
                        [
                            null
                        ]
                    ],
                    "instance_interface_id": [
                        null
                    ],
                    "ip_id": [
                        [
                            null
                        ]
                    ]
                },
                {
                    "network_equipment_interface_id": 83,
                    "network_equipment_identifier_string": "dtsy1sppstl006",
                    "network_equipment_id": 10,
                    "network_equipment_interface_identifier_string": "ethernet1\/1\/5",
                    "server_interface_mac_address": "b4:96:91:c3:82:6f",
                    "server_interface_index": 3,
                    "server_interface_capacity_mbps": 25000,
                    "ip_human_readable": [
                        [
                            null
                        ]
                    ],
                    "subnet_range_start_human_readable": [
                        [
                            null
                        ]
                    ],
                    "network_type": [
                        null
                    ],
                    "network_id": [
                        null
                    ],
                    "subnet_prefix_size": [
                        [
                            null
                        ]
                    ],
                    "subnet_pool_id": [
                        [
                            null
                        ]
                    ],
                    "instance_interface_id": [
                        null
                    ],
                    "ip_id": [
                        [
                            null
                        ]
                    ]
                },
                {
                    "network_equipment_interface_id": 88,
                    "network_equipment_identifier_string": "dtsy1sppdal003",
                    "network_equipment_id": 15,
                    "network_equipment_interface_identifier_string": "ethernet1\/1\/1",
                    "server_interface_mac_address": "b4:96:91:ba:19:d4",
                    "server_interface_index": 0,
                    "server_interface_capacity_mbps": 25000,
                    "ip_human_readable": [
                        [
                            null
                        ]
                    ],
                    "subnet_range_start_human_readable": [
                        [
                            null
                        ]
                    ],
                    "network_type": [
                        null
                    ],
                    "network_id": [
                        null
                    ],
                    "subnet_prefix_size": [
                        [
                            null
                        ]
                    ],
                    "subnet_pool_id": [
                        [
                            null
                        ]
                    ],
                    "instance_interface_id": [
                        null
                    ],
                    "ip_id": [
                        [
                            null
                        ]
                    ]
                },
                {
                    "network_equipment_interface_id": 89,
                    "network_equipment_identifier_string": "dtsy1sppdal004",
                    "network_equipment_id": 16,
                    "network_equipment_interface_identifier_string": "ethernet1\/1\/1",
                    "server_interface_mac_address": "b4:96:91:ba:19:d5",
                    "server_interface_index": 1,
                    "server_interface_capacity_mbps": 25000,
                    "ip_human_readable": [
                        [
                            null
                        ]
                    ],
                    "subnet_range_start_human_readable": [
                        [
                            null
                        ]
                    ],
                    "network_type": [
                        null
                    ],
                    "network_id": [
                        null
                    ],
                    "subnet_prefix_size": [
                        [
                            null
                        ]
                    ],
                    "subnet_pool_id": [
                        [
                            null
                        ]
                    ],
                    "instance_interface_id": [
                        null
                    ],
                    "ip_id": [
                        [
                            null
                        ]
                    ]
                },
                {
                    "network_equipment_interface_id": 90,
                    "network_equipment_identifier_string": "dtsy1sppstl008",
                    "network_equipment_id": 17,
                    "network_equipment_interface_identifier_string": "ethernet1\/1\/1",
                    "server_interface_mac_address": "b4:96:91:b9:2f:7f",
                    "server_interface_index": 2,
                    "server_interface_capacity_mbps": 25000,
                    "ip_human_readable": [
                        [
                            null
                        ]
                    ],
                    "subnet_range_start_human_readable": [
                        [
                            null
                        ]
                    ],
                    "network_type": [
                        null
                    ],
                    "network_id": [
                        null
                    ],
                    "subnet_prefix_size": [
                        [
                            null
                        ]
                    ],
                    "subnet_pool_id": [
                        [
                            null
                        ]
                    ],
                    "instance_interface_id": [
                        null
                    ],
                    "ip_id": [
                        [
                            null
                        ]
                    ]
                },
                {
                    "network_equipment_interface_id": 91,
                    "network_equipment_identifier_string": "dtsy1sppstl007",
                    "network_equipment_id": 18,
                    "network_equipment_interface_identifier_string": "ethernet1\/1\/1",
                    "server_interface_mac_address": "b4:96:91:b9:2f:7e",
                    "server_interface_index": 3,
                    "server_interface_capacity_mbps": 25000,
                    "ip_human_readable": [
                        [
                            null
                        ]
                    ],
                    "subnet_range_start_human_readable": [
                        [
                            null
                        ]
                    ],
                    "network_type": [
                        null
                    ],
                    "network_id": [
                        null
                    ],
                    "subnet_prefix_size": [
                        [
                            null
                        ]
                    ],
                    "subnet_pool_id": [
                        [
                            null
                        ]
                    ],
                    "instance_interface_id": [
                        null
                    ],
                    "ip_id": [
                        [
                            null
                        ]
                    ]
                }
            ],
            "rows_total": 8,
            "rows_order": [
                [
                    "network_equipment_interface_id",
                    "ASC"
                ],
                [
                    "instance_interface_id",
                    "ASC"
                ],
                [
                    "ip_id",
                    "ASC"
                ]
            ],
            "xls_for_pivot_tables_download_url": ""
        }
    }`

const _switchInterfaceFixture2 = `{
	"_switch_interfaces": {
		"duration_milliseconds": 0.07420182228088379,
		"rows": [
			{
				"network_equipment_interface_id": 17,
				"network_equipment_identifier_string": "ddd12",
				"network_equipment_id": 4,
				"network_equipment_interface_identifier_string": "ethernet1\/1\/9:1",
				"server_interface_mac_address": "68:05:ca:e0:5b:b6",
				"server_interface_index": 3,
				"server_interface_capacity_mbps": 10000,
				"ip_human_readable": [
					[
						null
					]
				],
				"subnet_range_start_human_readable": [
					[
						null
					]
				],
				"network_type": [
					null
				],
				"network_id": [
					null
				],
				"subnet_prefix_size": [
					[
						null
					]
				],
				"subnet_pool_id": [
					[
						null
					]
				],
				"instance_interface_id": [
					null
				],
				"ip_id": [
					[
						null
					]
				]
			},
			{
				"network_equipment_interface_id": 33,
				"network_equipment_identifier_string": "dasd",
				"network_equipment_id": 1,
				"network_equipment_interface_identifier_string": "ethernet1\/1\/5",
				"server_interface_mac_address": "b4:96:91:b9:2f:57",
				"server_interface_index": 0,
				"server_interface_capacity_mbps": 25000,
				"ip_human_readable": [
					[
						null
					]
				],
				"subnet_range_start_human_readable": [
					[
						null
					]
				],
				"network_type": [
					"wan"
				],
				"network_id": [
					249
				],
				"subnet_prefix_size": [
					[
						null
					]
				],
				"subnet_pool_id": [
					[
						null
					]
				],
				"instance_interface_id": [
					902
				],
				"ip_id": [
					[
						null
					]
				]
			},
			{
				"network_equipment_interface_id": 34,
				"network_equipment_identifier_string": "asdasd",
				"network_equipment_id": 2,
				"network_equipment_interface_identifier_string": "ethernet1\/1\/5",
				"server_interface_mac_address": "b4:96:91:ba:0b:83",
				"server_interface_index": 1,
				"server_interface_capacity_mbps": 25000,
				"ip_human_readable": [
					[
						"fd50:d68d:add7:0001:0000:0000:0000:0004",
						"192.168.78.4"
					]
				],
				"subnet_range_start_human_readable": [
					[
						"fd50:d68d:add7:0001:0000:0000:0000:0000",
						"192.168.78.0"
					]
				],
				"network_type": [
					"wan"
				],
				"network_id": [
					249
				],
				"subnet_prefix_size": [
					[
						64,
						29
					]
				],
				"subnet_pool_id": [
					[
						2,
						1
					]
				],
				"instance_interface_id": [
					903
				],
				"ip_id": [
					[
						1245,
						1246
					]
				]
			},
			{
				"network_equipment_interface_id": 35,
				"network_equipment_identifier_string": "dtsy1spmstl001",
				"network_equipment_id": 3,
				"network_equipment_interface_identifier_string": "ethernet1\/1\/5",
				"server_interface_mac_address": "b4:96:91:b9:2f:56",
				"server_interface_index": 2,
				"server_interface_capacity_mbps": 25000,
				"ip_human_readable": [
					[
						"fd64:ed65:bef3:0000:0000:0000:0000:0006"
					]
				],
				"subnet_range_start_human_readable": [
					[
						"fd64:ed65:bef3:0000:0000:0000:0000:0000"
					]
				],
				"network_type": [
					"san"
				],
				"network_id": [
					250
				],
				"subnet_prefix_size": [
					[
						64
					]
				],
				"subnet_pool_id": [
					[
						3
					]
				],
				"instance_interface_id": [
					904
				],
				"ip_id": [
					[
						1249
					]
				]
			},
			{
				"network_equipment_interface_id": 36,
				"network_equipment_identifier_string": "asdasd",
				"network_equipment_id": 4,
				"network_equipment_interface_identifier_string": "ethernet1\/1\/5",
				"server_interface_mac_address": "b4:96:91:ba:0b:82",
				"server_interface_index": 3,
				"server_interface_capacity_mbps": 25000,
				"ip_human_readable": [
					[
						"fd64:ed65:bef3:0001:0000:0000:0000:0006"
					]
				],
				"subnet_range_start_human_readable": [
					[
						"fd64:ed65:bef3:0001:0000:0000:0000:0000"
					]
				],
				"network_type": [
					"san"
				],
				"network_id": [
					250
				],
				"subnet_prefix_size": [
					[
						64
					]
				],
				"subnet_pool_id": [
					[
						3
					]
				],
				"instance_interface_id": [
					905
				],
				"ip_id": [
					[
						1250
					]
				]
			},
			{
				"network_equipment_interface_id": 37,
				"network_equipment_identifier_string": "asdasd",
				"network_equipment_id": 1,
				"network_equipment_interface_identifier_string": "ethernet1\/1\/3",
				"server_interface_mac_address": "b4:96:91:b9:2f:af",
				"server_interface_index": 0,
				"server_interface_capacity_mbps": 25000,
				"ip_human_readable": [
					[
						null
					]
				],
				"subnet_range_start_human_readable": [
					[
						null
					]
				],
				"network_type": [
					"wan"
				],
				"network_id": [
					249
				],
				"subnet_prefix_size": [
					[
						null
					]
				],
				"subnet_pool_id": [
					[
						null
					]
				],
				"instance_interface_id": [
					918
				],
				"ip_id": [
					[
						null
					]
				]
			},
			{
				"network_equipment_interface_id": 38,
				"network_equipment_identifier_string": "asdasdasd",
				"network_equipment_id": 2,
				"network_equipment_interface_identifier_string": "ethernet1\/1\/3",
				"server_interface_mac_address": "b4:96:91:ba:12:83",
				"server_interface_index": 1,
				"server_interface_capacity_mbps": 25000,
				"ip_human_readable": [
					[
						"fd50:d68d:add7:0001:0000:0000:0000:0006",
						"192.168.78.6"
					]
				],
				"subnet_range_start_human_readable": [
					[
						"fd50:d68d:add7:0001:0000:0000:0000:0000",
						"192.168.78.0"
					]
				],
				"network_type": [
					"wan"
				],
				"network_id": [
					249
				],
				"subnet_prefix_size": [
					[
						64,
						29
					]
				],
				"subnet_pool_id": [
					[
						2,
						1
					]
				],
				"instance_interface_id": [
					919
				],
				"ip_id": [
					[
						1277,
						1278
					]
				]
			},
			{
				"network_equipment_interface_id": 39,
				"network_equipment_identifier_string": "asdasd",
				"network_equipment_id": 3,
				"network_equipment_interface_identifier_string": "ethernet1\/1\/3",
				"server_interface_mac_address": "b4:96:91:b9:2f:ae",
				"server_interface_index": 2,
				"server_interface_capacity_mbps": 25000,
				"ip_human_readable": [
					[
						"fd64:ed65:bef3:0000:0000:0000:0000:0008"
					]
				],
				"subnet_range_start_human_readable": [
					[
						"fd64:ed65:bef3:0000:0000:0000:0000:0000"
					]
				],
				"network_type": [
					"san"
				],
				"network_id": [
					250
				],
				"subnet_prefix_size": [
					[
						64
					]
				],
				"subnet_pool_id": [
					[
						3
					]
				],
				"instance_interface_id": [
					920
				],
				"ip_id": [
					[
						1279
					]
				]
			},
			{
				"network_equipment_interface_id": 40,
				"network_equipment_identifier_string": "asdasd",
				"network_equipment_id": 4,
				"network_equipment_interface_identifier_string": "ethernet1\/1\/3",
				"server_interface_mac_address": "b4:96:91:ba:12:82",
				"server_interface_index": 3,
				"server_interface_capacity_mbps": 25000,
				"ip_human_readable": [
					[
						"fd64:ed65:bef3:0001:0000:0000:0000:0008"
					]
				],
				"subnet_range_start_human_readable": [
					[
						"fd64:ed65:bef3:0001:0000:0000:0000:0000"
					]
				],
				"network_type": [
					"san"
				],
				"network_id": [
					250
				],
				"subnet_prefix_size": [
					[
						64
					]
				],
				"subnet_pool_id": [
					[
						3
					]
				],
				"instance_interface_id": [
					921
				],
				"ip_id": [
					[
						1280
					]
				]
			}
		],
		"rows_total": 9,
		"rows_order": [
			[
				"network_equipment_interface_id",
				"ASC"
			],
			[
				"instance_interface_id",
				"ASC"
			],
			[
				"ip_id",
				"ASC"
			]
		],
		"xls_for_pivot_tables_download_url": ""
	}
}`

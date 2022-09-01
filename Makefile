.PHONY: build create_testnet run_testnet

build:
	bash ./scripts/go_build_executables.sh

run_testnet:
	bash ./scripts/run_testnet.sh

run_ahltest:
	bash ./scripts/run_ahl.sh

run_elrondtest:
	bash ./scripts/run_elrond.sh

run_haechitest:
	bash ./scripts/run_haechi.sh

#haechi
run_haechi_2shard_4node_test:
	bash ./scripts/run_haechi_2shard_4node.sh

run_haechi_4shard_4node_test:
	bash ./scripts/run_haechi_4shard_4node.sh

run_haechi_6shard_4node_test:
	bash ./scripts/run_haechi_6shard_4node.sh

run_haechi_8shard_4node_test:
	bash ./scripts/run_haechi_8shard_4node.sh

run_haechi_10shard_4node_test:
	bash ./scripts/run_haechi_10shard_4node.sh

run_haechi_12shard_4node_test:
	bash ./scripts/run_haechi_12shard_4node.sh

run_haechi_14shard_4node_test:
	bash ./scripts/run_haechi_14shard_4node.sh

run_haechi_16shard_4node_test:
	bash ./scripts/run_haechi_16shard_4node.sh

run_haechi_16shard_4node_8beacon_test:
	bash ./scripts/run_haechi_16shard_4node_8beacon.sh

run_haechi_16shard_4node_12beacon_test:
	bash ./scripts/run_haechi_16shard_4node_12beacon.sh

run_haechi_16shard_4node_16beacon_test:
	bash ./scripts/run_haechi_16shard_4node_16beacon.sh

run_haechi_16shard_4node_20beacon_test:
	bash ./scripts/run_haechi_16shard_4node_20beacon.sh

run_haechi_16shard_8node_test:
	bash ./scripts/run_haechi_16shard_8node.sh

run_haechi_16shard_12node_test:
	bash ./scripts/run_haechi_16shard_12node.sh

run_haechi_16shard_16node_test:
	bash ./scripts/run_haechi_16shard_16node.sh

run_haechi_16shard_20node_test:
	bash ./scripts/run_haechi_16shard_20node.sh

# ahl
run_ahl_2shard_4node_test:
	bash ./scripts/run_ahl_2shard_4node.sh

run_ahl_4shard_4node_test:
	bash ./scripts/run_ahl_4shard_4node.sh

run_ahl_6shard_4node_test:
	bash ./scripts/run_ahl_6shard_4node.sh

run_ahl_8shard_4node_test:
	bash ./scripts/run_ahl_8shard_4node.sh

run_ahl_10shard_4node_test:
	bash ./scripts/run_ahl_10shard_4node.sh

run_ahl_12shard_4node_test:
	bash ./scripts/run_ahl_12shard_4node.sh

run_ahl_14shard_4node_test:
	bash ./scripts/run_ahl_14shard_4node.sh

run_ahl_16shard_4node_test:
	bash ./scripts/run_ahl_16shard_4node.sh

run_ahl_16shard_8node_test:
	bash ./scripts/run_ahl_16shard_8node.sh

run_ahl_16shard_12node_test:
	bash ./scripts/run_ahl_16shard_12node.sh

run_ahl_16shard_16node_test:
	bash ./scripts/run_ahl_16shard_16node.sh

run_ahl_16shard_20node_test:
	bash ./scripts/run_ahl_16shard_20node.sh

#elrond
run_elrond_2shard_4node_test:
	bash ./scripts/run_elrond_2shard_4node.sh

run_elrond_4shard_4node_test:
	bash ./scripts/run_elrond_4shard_4node.sh

run_elrond_6shard_4node_test:
	bash ./scripts/run_elrond_6shard_4node.sh

run_elrond_8shard_4node_test:
	bash ./scripts/run_elrond_8shard_4node.sh

run_elrond_10shard_4node_test:
	bash ./scripts/run_elrond_10shard_4node.sh

run_elrond_12shard_4node_test:
	bash ./scripts/run_elrond_12shard_4node.sh

run_elrond_14shard_4node_test:
	bash ./scripts/run_elrond_14shard_4node.sh

run_elrond_16shard_4node_test:
	bash ./scripts/run_elrond_16shard_4node.sh

run_elrond_16shard_8node_test:
	bash ./scripts/run_elrond_16shard_8node.sh

run_elrond_16shard_12node_test:
	bash ./scripts/run_elrond_16shard_12node.sh

run_elrond_16shard_16node_test:
	bash ./scripts/run_elrond_16shard_16node.sh

run_elrond_16shard_20node_test:
	bash ./scripts/run_elrond_16shard_20node.sh

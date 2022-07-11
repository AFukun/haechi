.PHONY: build create_testnet run_testnet

build:
	bash ./scripts/go_build_executables.sh

run_testnet:
	bash ./scripts/run_testnet.sh

run_elrondtest:
	bash ./scripts/run_elrond.sh

run_elrondtest:
	bash ./scripts/run_haechi.sh

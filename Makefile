.PHONY: build_testnet build_ahl build_elrond build_haechi run_testnet run_ahl run_elrond run_haechi

build: build_testnet build_ahl build_elrond build_haechi

build_testnet:
	bash ./scripts/go_build_testnet.sh

build_ahl:
	bash ./scripts/go_build_ahl.sh

build_elrond:
	bash ./scripts/go_build_elrond.sh

build_haechi:
	bash ./scripts/go_build_haechi.sh

run_testnet:
	bash ./scripts/run_testnet.sh

run_ahl:
	bash ./scripts/run_ahl.sh

run_elrond:
	bash ./scripts/run_elrond.sh

run_haechi:
	bash ./scripts/run_haechi.sh

clean:
	rm -rf build

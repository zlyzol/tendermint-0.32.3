build:
	@sh -c "'$(CURDIR)/build.sh'"

push:
	@sh -c "'$(CURDIR)/push.sh'"

build_develop:
	docker build -t "zlyzol/tendermint-0.32.3:develop" -f Dockerfile.develop .

build_testing:
	docker build --tag tendermint/testing -f ./Dockerfile.testing .

push_develop:
	docker push "zlyzol/tendermint-0.32.3:develop"

.PHONY: build build_develop push push_develop

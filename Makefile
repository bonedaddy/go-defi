GIT_VERSION=`git describe --tags`

# NOTE: if you dont want to use the sqlite database you can remove the json1 tag
.PHONY: cli
cli:
	go build --tags "json1" -ldflags "-X main.Version=$(GIT_VERSION)" .

.PHONY: docker
docker:
	docker build --build-arg VERSION=$(GIT_VERSION) -t bonedaddy/go-defi:$(GIT_VERSION) .
	docker image tag bonedaddy/go-defi:$(GIT_VERSION) bonedaddy/go-defi:latest

.PHONY: flatten-dependencies
flatten-dependencies:
	@echo "generating abis"
	(cd v3-core ; npm run compile ; cp -r abis/contracts/* ../abi/uniswap/v3/core)
	(cd v3-periphery ; npm run compile ; cp -r abis/contracts/* ../abi/uniswap/v3/periphery)

.PHONY: abigen
abigen:
	abigen \
		--abi abi/uniswap/v3/core/UniswapV3Factory.sol/UniswapV3Factory.json \
		--bin bin/uniswap/v3/core/UniswapV3Factory.bin \
		--pkg v3factory \
		--out bindings/uniswap/v3/factory/bindings.go
	abigen \
		--abi abi/uniswap/v3/core/UniswapV3Pool.sol/UniswapV3Pool.json \
		--bin bin/uniswap/v3/core/UniswapV3Pool.bin \
		--pkg v3pool \
		--out bindings/uniswap/v3/pool/bindings.go
	abigen \
		--abi abi/uniswap/v3/core/UniswapV3PoolDeployer.sol/UniswapV3PoolDeployer.json \
		--bin bin/uniswap/v3/core/UniswapV3PoolDeployer.bin \
		--pkg v3pooldeployer \
		--out bindings/uniswap/v3/pooldeployer/bindings.go
	abigen \
		--abi abi/uniswap/v3/periphery/NonfungiblePositionManager.sol/NonfungiblePositionManager.json \
		--bin bin/uniswap/v3/periphery/NonfungiblePositionManager.bin \
		--pkg v3nftposmanager \
		--out bindings/uniswap/v3/nftposmanager/bindings.go
	abigen \
		--abi abi/uniswap/v3/periphery/SwapRouter.sol/SwapRouter.json \
		--bin bin/uniswap/v3/periphery/SwapRouter.bin \
		--pkg v3swaprouter \
		--out bindings/uniswap/v3/swaprouter/bindings.go

.PHONY: release
release:
	./scripts/release.sh

.PHONY: tests
tests:
	go test --tags "json1" -count 1 -cover ./...
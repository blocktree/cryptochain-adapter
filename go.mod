module github.com/blocktree/cryptochain-adapter

go 1.12

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet/v2 v2.0.10
	github.com/cosmos/cosmos-sdk v0.41.4
	github.com/ethereum/go-ethereum v1.9.9
	github.com/graarh/golang-socketio v0.0.0-20170510162725-2c44953b9b5f
	github.com/imroc/req v0.2.4
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
	github.com/tidwall/gjson v1.3.5
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

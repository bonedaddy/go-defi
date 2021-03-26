// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v3swaprouter

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// V3swaprouterABI is the input ABI used to generate the binding from.
const V3swaprouterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// V3swaprouterBin is the compiled bytecode used for deploying new contracts.
var V3swaprouterBin = "0x6080604052600436106100cb5760003560e01c8063c2e3140a11610074578063f28c04981161004e578063f28c04981461025a578063f3995c671461026d578063fa461e331461028057610176565b8063c2e3140a1461021f578063c45a015514610232578063df2ab5bb1461024757610176565b8063a4a78f0c116100a5578063a4a78f0c146101cc578063ac9650d8146101df578063c04b8d59146101ff57610176565b80634659a4941461017b57806349404b7c1461018e5780634aa4a4fc146101a157610176565b36610176573373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461017457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742057455448390000000000000000000000000000000000000000000000604482015290519081900360640190fd5b005b600080fd5b610174610189366004611efd565b6102a0565b61017461019c36600461223f565b610360565b3480156101ad57600080fd5b506101b661052c565b6040516101c391906122c4565b60405180910390f35b6101746101da366004611efd565b610550565b6101f26101ed366004611f5b565b61062d565b6040516101c39190612337565b61021261020d3660046120d2565b610790565b6040516101c3919061247e565b61017461022d366004611efd565b6108d1565b34801561023e57600080fd5b506101b6610986565b610174610255366004611ec2565b6109aa565b61021261026836600461217d565b610acd565b61017461027b366004611efd565b610c22565b34801561028c57600080fd5b5061017461029b366004611fed565b610cba565b604080517f8fcbaf0c00000000000000000000000000000000000000000000000000000000815233600482015230602482015260448101879052606481018690526001608482015260ff851660a482015260c4810184905260e48101839052905173ffffffffffffffffffffffffffffffffffffffff881691638fcbaf0c9161010480830192600092919082900301818387803b15801561034057600080fd5b505af1158015610354573d6000803e3d6000fd5b50505050505050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156103e957600080fd5b505afa1580156103fd573d6000803e3d6000fd5b505050506040513d602081101561041357600080fd5b50519050821561048c578281101561048c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f496e73756666696369656e742057455448390000000000000000000000000000604482015290519081900360640190fd5b8015610527577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561050557600080fd5b505af1158015610519573d6000803e3d6000fd5b505050506105278282610de1565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b604080517fdd62ed3e00000000000000000000000000000000000000000000000000000000815233600482015230602482015290517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9173ffffffffffffffffffffffffffffffffffffffff89169163dd62ed3e91604480820192602092909190829003018186803b1580156105e557600080fd5b505afa1580156105f9573d6000803e3d6000fd5b505050506040513d602081101561060f57600080fd5b50511015610625576106258686868686866102a0565b505050505050565b60608167ffffffffffffffff8111801561064657600080fd5b5060405190808252806020026020018201604052801561067a57816020015b60608152602001906001900390816106655790505b50905060005b82811015610789576000803086868581811061069857fe5b90506020028101906106aa9190612487565b6040516106b89291906122b4565b600060405180830381855af49150503d80600081146106f3576040519150601f19603f3d011682016040523d82523d6000602084013e6106f8565b606091505b5091509150816107675760448151101561071157600080fd5b6004810190508080602001905181019061072b9190612068565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075e91906123b5565b60405180910390fd5b8084848151811061077457fe5b60209081029190910101525050600101610680565b5092915050565b60008160400151806107a0610f2f565b111561080d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f5472616e73616374696f6e20746f6f206f6c6400000000000000000000000000604482015290519081900360640190fd5b600061081c8460000151610f33565b905061085d846060015182610835578560200151610837565b305b604051806040016040528061084f8960000151610f3e565b815233602090910152610f53565b6060850152801561087a57835161087390611096565b8452610887565b836060015192505061088d565b5061080d565b82608001518210156108cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075e906123ff565b50919050565b604080517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523360048201523060248201529051869173ffffffffffffffffffffffffffffffffffffffff89169163dd62ed3e91604480820192602092909190829003018186803b15801561094657600080fd5b505afa15801561095a573d6000803e3d6000fd5b505050506040513d602081101561097057600080fd5b5051101561062557610625868686868686610c22565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610a1357600080fd5b505afa158015610a27573d6000803e3d6000fd5b505050506040513d6020811015610a3d57600080fd5b505190508215610ab65782811015610ab657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f496e73756666696369656e7420746f6b656e0000000000000000000000000000604482015290519081900360640190fd5b8015610ac757610ac78483836110cb565b50505050565b6000816040013580610add610f2f565b1115610b4a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f5472616e73616374696f6e20746f6f206f6c6400000000000000000000000000604482015290519081900360640190fd5b610bba6060840135610b626040860160208701611ea1565b6040805180820190915280610b778880612487565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250336020909101526112a7565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff909155915060808301358211156108cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075e906123c8565b604080517fd505accf000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018790526064810186905260ff8516608482015260a4810184905260c48101839052905173ffffffffffffffffffffffffffffffffffffffff88169163d505accf9160e480830192600092919082900301818387803b15801561034057600080fd5b6000610cc8828401846121b5565b90506000806000610cdc84600001516113d2565b925092509250610d0e7f0000000000000000000000000000000000000000000000000000000000000000848484611403565b5060008060008a13610d4f578473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161089610d80565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16108a5b915091508115610d9f57610d9a8587602001513384611422565b610354565b8551610daa90610f33565b15610dc7578551610dba90611096565b8652610d9a8133886112a7565b806000819055508394506103548587602001513384611422565b6040805160008082526020820190925273ffffffffffffffffffffffffffffffffffffffff84169083906040518082805190602001908083835b60208310610e5857805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610e1b565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114610eba576040519150601f19603f3d011682016040523d82523d6000602084013e610ebf565b606091505b505090508061052757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600360248201527f5354450000000000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b4290565b8051602b105b919050565b6060610f4d826000602b611605565b92915050565b600080600080610f6685600001516113d2565b9194509250905073ffffffffffffffffffffffffffffffffffffffff80831690841610600080610f978686866117ec565b73ffffffffffffffffffffffffffffffffffffffff1663128acb088a85610fbd8e61182a565b87610fdc5773fffd8963efd1fc6a506488495d951d5263988d25610fe3565b6401000276a45b8d604051602001610ff49190612436565b6040516020818303038152906040526040518663ffffffff1660e01b81526004016110239594939291906122e5565b6040805180830381600087803b15801561103c57600080fd5b505af1158015611050573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110749190611fca565b91509150826110835781611085565b805b6000039a9950505050505050505050565b8051606090610f4d9083906017907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe901611605565b6040805173ffffffffffffffffffffffffffffffffffffffff8481166024830152604480830185905283518084039091018152606490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001781529251825160009485949389169392918291908083835b602083106111a057805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101611163565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611202576040519150601f19603f3d011682016040523d82523d6000602084013e611207565b606091505b5091509150818015611235575080511580611235575080806020019051602081101561123257600080fd5b50515b6112a057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600260248201527f5354000000000000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b5050505050565b60008060006112b984600001516113d2565b9194509250905073ffffffffffffffffffffffffffffffffffffffff808416908316106112e78385846117ec565b73ffffffffffffffffffffffffffffffffffffffff1663128acb08878361130d8b61182a565b6000038561132f5773fffd8963efd1fc6a506488495d951d5263988d25611336565b6401000276a45b8a6040516020016113479190612436565b6040516020818303038152906040526040518663ffffffff1660e01b81526004016113769594939291906122e5565b6040805180830381600087803b15801561138f57600080fd5b505af11580156113a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c79190611fca565b505050505050505050565b600080806113e0848261185c565b92506113ed84601461195c565b90506113fa84601761185c565b91509193909250565b600061141985611414868686611a4c565b611ac9565b95945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161480156114825750814791508110155b156115cb577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b1580156114ef57600080fd5b505af1158015611503573d6000803e3d6000fd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561159957600080fd5b505af11580156115ad573d6000803e3d6000fd5b505050506040513d60208110156115c357600080fd5b506112a09050565b73ffffffffffffffffffffffffffffffffffffffff84163014156115f9576115f48584846110cb565b6112a0565b6112a085858585611af9565b60608182601f01101561167957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f736c6963655f6f766572666c6f77000000000000000000000000000000000000604482015290519081900360640190fd5b8282840110156116ea57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f736c6963655f6f766572666c6f77000000000000000000000000000000000000604482015290519081900360640190fd5b8183018451101561175c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f736c6963655f6f75744f66426f756e6473000000000000000000000000000000604482015290519081900360640190fd5b60608215801561177b57604051915060008252602082016040526117e3565b6040519150601f8416801560200281840101858101878315602002848b0101015b818310156117b457805183526020928301920161179c565b5050858452601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016604052505b50949350505050565b60006118227f000000000000000000000000000000000000000000000000000000000000000061181d868686611a4c565b611cd6565b949350505050565b60007f8000000000000000000000000000000000000000000000000000000000000000821061185857600080fd5b5090565b6000818260140110156118d057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f746f416464726573735f6f766572666c6f770000000000000000000000000000604482015290519081900360640190fd5b816014018351101561194357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f746f416464726573735f6f75744f66426f756e64730000000000000000000000604482015290519081900360640190fd5b5001602001516c01000000000000000000000000900490565b6000818260030110156119d057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f746f55696e7432345f6f766572666c6f77000000000000000000000000000000604482015290519081900360640190fd5b8160030183511015611a4357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f746f55696e7432345f6f75744f66426f756e6473000000000000000000000000604482015290519081900360640190fd5b50016003015190565b611a54611e0c565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161115611a8c579192915b506040805160608101825273ffffffffffffffffffffffffffffffffffffffff948516815292909316602083015262ffffff169181019190915290565b6000611ad58383611cd6565b90503373ffffffffffffffffffffffffffffffffffffffff821614610f4d57600080fd5b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528481166044830152606480830185905283518084039091018152608490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000178152925182516000948594938a169392918291908083835b60208310611bd657805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101611b99565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611c38576040519150601f19603f3d011682016040523d82523d6000602084013e611c3d565b606091505b5091509150818015611c6b575080511580611c6b5750808060200190516020811015611c6857600080fd5b50515b61062557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600360248201527f5354460000000000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6000816020015173ffffffffffffffffffffffffffffffffffffffff16826000015173ffffffffffffffffffffffffffffffffffffffff1610611d1857600080fd5b508051602080830151604093840151845173ffffffffffffffffffffffffffffffffffffffff94851681850152939091168385015262ffffff166060808401919091528351808403820181526080840185528051908301207fff0000000000000000000000000000000000000000000000000000000000000060a085015294901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660a183015260b58201939093527f01d4d358e07707f4db84b6a7527455b06f95ee89b5d059b4a1298ada7b6c7d6760d5808301919091528251808303909101815260f5909101909152805191012090565b604080516060810182526000808252602082018190529181019190915290565b803573ffffffffffffffffffffffffffffffffffffffff81168114610f3957600080fd5b600082601f830112611e60578081fd5b8135611e73611e6e82612515565b6124f1565b818152846020838601011115611e87578283fd5b816020850160208301379081016020019190915292915050565b600060208284031215611eb2578081fd5b611ebb82611e2c565b9392505050565b600080600060608486031215611ed6578182fd5b611edf84611e2c565b925060208401359150611ef460408501611e2c565b90509250925092565b60008060008060008060c08789031215611f15578182fd5b611f1e87611e2c565b95506020870135945060408701359350606087013560ff81168114611f41578283fd5b9598949750929560808101359460a0909101359350915050565b60008060208385031215611f6d578182fd5b823567ffffffffffffffff80821115611f84578384fd5b818501915085601f830112611f97578384fd5b813581811115611fa5578485fd5b8660208083028501011115611fb8578485fd5b60209290920196919550909350505050565b60008060408385031215611fdc578182fd5b505080516020909101519092909150565b60008060008060608587031215612002578384fd5b8435935060208501359250604085013567ffffffffffffffff80821115612027578384fd5b818701915087601f83011261203a578384fd5b813581811115612048578485fd5b886020828501011115612059578485fd5b95989497505060200194505050565b600060208284031215612079578081fd5b815167ffffffffffffffff81111561208f578182fd5b8201601f8101841361209f578182fd5b80516120ad611e6e82612515565b8181528560208385010111156120c1578384fd5b611419826020830160208601612555565b6000602082840312156120e3578081fd5b813567ffffffffffffffff808211156120fa578283fd5b9083019060a0828603121561210d578283fd5b60405160a08101818110838211171561212257fe5b604052823582811115612133578485fd5b61213f87828601611e50565b82525061214e60208401611e2c565b602082015260408301356040820152606083013560608201526080830135608082015280935050505092915050565b60006020828403121561218e578081fd5b813567ffffffffffffffff8111156121a4578182fd5b820160a08185031215611ebb578182fd5b6000602082840312156121c6578081fd5b813567ffffffffffffffff808211156121dd578283fd5b90830190604082860312156121f0578283fd5b60405160408101818110838211171561220557fe5b604052823582811115612216578485fd5b61222287828601611e50565b82525061223160208401611e2c565b602082015295945050505050565b60008060408385031215612251578182fd5b8235915061226160208401611e2c565b90509250929050565b60008151808452612282816020860160208601612555565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6000828483379101908152919050565b73ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b600073ffffffffffffffffffffffffffffffffffffffff8088168352861515602084015285604084015280851660608401525060a0608083015261232c60a083018461226a565b979650505050505050565b6000602080830181845280855180835260408601915060408482028701019250838701855b828110156123a8577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc088860301845261239685835161226a565b9450928501929085019060010161235c565b5092979650505050505050565b600060208252611ebb602083018461226a565b60208082526012908201527f546f6f206d756368207265717565737465640000000000000000000000000000604082015260600190565b60208082526013908201527f546f6f206c6974746c6520726563656976656400000000000000000000000000604082015260600190565b600060208252825160406020840152612452606084018261226a565b905073ffffffffffffffffffffffffffffffffffffffff60208501511660408401528091505092915050565b90815260200190565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126124bb578283fd5b83018035915067ffffffffffffffff8211156124d5578283fd5b6020019150368190038213156124ea57600080fd5b9250929050565b60405181810167ffffffffffffffff8111828210171561250d57fe5b604052919050565b600067ffffffffffffffff82111561252957fe5b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b83811015612570578181015183820152602001612558565b83811115610ac7575050600091015256fea164736f6c6343000706000a"

// DeployV3swaprouter deploys a new Ethereum contract, binding an instance of V3swaprouter to it.
func DeployV3swaprouter(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _WETH9 common.Address) (common.Address, *types.Transaction, *V3swaprouter, error) {
	parsed, err := abi.JSON(strings.NewReader(V3swaprouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(V3swaprouterBin), backend, _factory, _WETH9)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &V3swaprouter{V3swaprouterCaller: V3swaprouterCaller{contract: contract}, V3swaprouterTransactor: V3swaprouterTransactor{contract: contract}, V3swaprouterFilterer: V3swaprouterFilterer{contract: contract}}, nil
}

// V3swaprouter is an auto generated Go binding around an Ethereum contract.
type V3swaprouter struct {
	V3swaprouterCaller     // Read-only binding to the contract
	V3swaprouterTransactor // Write-only binding to the contract
	V3swaprouterFilterer   // Log filterer for contract events
}

// V3swaprouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3swaprouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3swaprouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3swaprouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3swaprouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3swaprouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3swaprouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3swaprouterSession struct {
	Contract     *V3swaprouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V3swaprouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3swaprouterCallerSession struct {
	Contract *V3swaprouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// V3swaprouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3swaprouterTransactorSession struct {
	Contract     *V3swaprouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// V3swaprouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3swaprouterRaw struct {
	Contract *V3swaprouter // Generic contract binding to access the raw methods on
}

// V3swaprouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3swaprouterCallerRaw struct {
	Contract *V3swaprouterCaller // Generic read-only contract binding to access the raw methods on
}

// V3swaprouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3swaprouterTransactorRaw struct {
	Contract *V3swaprouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3swaprouter creates a new instance of V3swaprouter, bound to a specific deployed contract.
func NewV3swaprouter(address common.Address, backend bind.ContractBackend) (*V3swaprouter, error) {
	contract, err := bindV3swaprouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3swaprouter{V3swaprouterCaller: V3swaprouterCaller{contract: contract}, V3swaprouterTransactor: V3swaprouterTransactor{contract: contract}, V3swaprouterFilterer: V3swaprouterFilterer{contract: contract}}, nil
}

// NewV3swaprouterCaller creates a new read-only instance of V3swaprouter, bound to a specific deployed contract.
func NewV3swaprouterCaller(address common.Address, caller bind.ContractCaller) (*V3swaprouterCaller, error) {
	contract, err := bindV3swaprouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3swaprouterCaller{contract: contract}, nil
}

// NewV3swaprouterTransactor creates a new write-only instance of V3swaprouter, bound to a specific deployed contract.
func NewV3swaprouterTransactor(address common.Address, transactor bind.ContractTransactor) (*V3swaprouterTransactor, error) {
	contract, err := bindV3swaprouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3swaprouterTransactor{contract: contract}, nil
}

// NewV3swaprouterFilterer creates a new log filterer instance of V3swaprouter, bound to a specific deployed contract.
func NewV3swaprouterFilterer(address common.Address, filterer bind.ContractFilterer) (*V3swaprouterFilterer, error) {
	contract, err := bindV3swaprouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3swaprouterFilterer{contract: contract}, nil
}

// bindV3swaprouter binds a generic wrapper to an already deployed contract.
func bindV3swaprouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(V3swaprouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3swaprouter *V3swaprouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3swaprouter.Contract.V3swaprouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3swaprouter *V3swaprouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3swaprouter.Contract.V3swaprouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3swaprouter *V3swaprouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3swaprouter.Contract.V3swaprouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3swaprouter *V3swaprouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3swaprouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3swaprouter *V3swaprouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3swaprouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3swaprouter *V3swaprouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3swaprouter.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3swaprouter *V3swaprouterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3swaprouter.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3swaprouter *V3swaprouterSession) WETH9() (common.Address, error) {
	return _V3swaprouter.Contract.WETH9(&_V3swaprouter.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3swaprouter *V3swaprouterCallerSession) WETH9() (common.Address, error) {
	return _V3swaprouter.Contract.WETH9(&_V3swaprouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3swaprouter *V3swaprouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3swaprouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3swaprouter *V3swaprouterSession) Factory() (common.Address, error) {
	return _V3swaprouter.Contract.Factory(&_V3swaprouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3swaprouter *V3swaprouterCallerSession) Factory() (common.Address, error) {
	return _V3swaprouter.Contract.Factory(&_V3swaprouter.CallOpts)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_V3swaprouter *V3swaprouterTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _V3swaprouter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_V3swaprouter *V3swaprouterSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _V3swaprouter.Contract.ExactInput(&_V3swaprouter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_V3swaprouter *V3swaprouterTransactorSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _V3swaprouter.Contract.ExactInput(&_V3swaprouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_V3swaprouter *V3swaprouterTransactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _V3swaprouter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_V3swaprouter *V3swaprouterSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _V3swaprouter.Contract.ExactOutput(&_V3swaprouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_V3swaprouter *V3swaprouterTransactorSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _V3swaprouter.Contract.ExactOutput(&_V3swaprouter.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3swaprouter *V3swaprouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _V3swaprouter.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3swaprouter *V3swaprouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.Multicall(&_V3swaprouter.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3swaprouter *V3swaprouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.Multicall(&_V3swaprouter.TransactOpts, data)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.SelfPermit(&_V3swaprouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.SelfPermit(&_V3swaprouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.SelfPermitAllowed(&_V3swaprouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.SelfPermitAllowed(&_V3swaprouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.SelfPermitAllowedIfNecessary(&_V3swaprouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.SelfPermitAllowedIfNecessary(&_V3swaprouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.SelfPermitIfNecessary(&_V3swaprouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3swaprouter *V3swaprouterTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.SelfPermitIfNecessary(&_V3swaprouter.TransactOpts, token, value, deadline, v, r, s)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3swaprouter *V3swaprouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3swaprouter.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3swaprouter *V3swaprouterSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3swaprouter.Contract.SweepToken(&_V3swaprouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3swaprouter *V3swaprouterTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3swaprouter.Contract.SweepToken(&_V3swaprouter.TransactOpts, token, amountMinimum, recipient)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_V3swaprouter *V3swaprouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3swaprouter.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_V3swaprouter *V3swaprouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.UniswapV3SwapCallback(&_V3swaprouter.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_V3swaprouter *V3swaprouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3swaprouter.Contract.UniswapV3SwapCallback(&_V3swaprouter.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_V3swaprouter *V3swaprouterTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3swaprouter.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_V3swaprouter *V3swaprouterSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3swaprouter.Contract.UnwrapWETH9(&_V3swaprouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_V3swaprouter *V3swaprouterTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3swaprouter.Contract.UnwrapWETH9(&_V3swaprouter.TransactOpts, amountMinimum, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3swaprouter *V3swaprouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3swaprouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3swaprouter *V3swaprouterSession) Receive() (*types.Transaction, error) {
	return _V3swaprouter.Contract.Receive(&_V3swaprouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3swaprouter *V3swaprouterTransactorSession) Receive() (*types.Transaction, error) {
	return _V3swaprouter.Contract.Receive(&_V3swaprouter.TransactOpts)
}

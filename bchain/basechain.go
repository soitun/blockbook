package bchain

import (
	"errors"
	"math/big"
)

// BaseChain is base type for bchain.BlockChain
type BaseChain struct {
	Parser  BlockChainParser
	Testnet bool
	Network string
}

// TODO more bchain.BlockChain methods

// GetChainParser returns BlockChainParser
func (b *BaseChain) GetChainParser() BlockChainParser {
	return b.Parser
}

// IsTestnet returns true if the blockchain is testnet
func (b *BaseChain) IsTestnet() bool {
	return b.Testnet
}

// GetNetworkName returns network name
func (b *BaseChain) GetNetworkName() string {
	return b.Network
}

// GetBlockRaw is not supported by default
func (b *BaseChain) GetBlockRaw(hash string) (string, error) {
	return "", errors.New("GetBlockRaw: not supported")
}

// GetMempoolEntry is not supported by default
func (b *BaseChain) GetMempoolEntry(txid string) (*MempoolEntry, error) {
	return nil, errors.New("GetMempoolEntry: not supported")
}

// LongTermFeeRate returns smallest fee rate from historic blocks.
func (b *BaseChain) LongTermFeeRate() (*LongTermFeeRate, error) {
	return nil, errors.New("not supported")
}

// EthereumTypeGetBalance is not supported
func (b *BaseChain) EthereumTypeGetBalance(addrDesc AddressDescriptor) (*big.Int, error) {
	return nil, errors.New("not supported")
}

// EthereumTypeGetNonce is not supported
func (b *BaseChain) EthereumTypeGetNonce(addrDesc AddressDescriptor) (uint64, error) {
	return 0, errors.New("not supported")
}

// EthereumTypeEstimateGas is not supported
func (b *BaseChain) EthereumTypeEstimateGas(params map[string]interface{}) (uint64, error) {
	return 0, errors.New("not supported")
}

// EthereumTypeGetEip1559Fees is not supported
func (b *BaseChain) EthereumTypeGetEip1559Fees() (*Eip1559Fees, error) {
	return nil, errors.New("not supported")
}

// GetContractInfo is not supported
func (b *BaseChain) GetContractInfo(contractDesc AddressDescriptor) (*ContractInfo, error) {
	return nil, errors.New("not supported")
}

// EthereumTypeGetErc20ContractBalance is not supported
func (b *BaseChain) EthereumTypeGetErc20ContractBalance(addrDesc, contractDesc AddressDescriptor) (*big.Int, error) {
	return nil, errors.New("not supported")
}

// GetTokenURI returns URI of non fungible or multi token defined by token id
func (p *BaseChain) GetTokenURI(contractDesc AddressDescriptor, tokenID *big.Int) (string, error) {
	return "", errors.New("not supported")
}

func (b *BaseChain) EthereumTypeGetSupportedStakingPools() []string {
	return nil
}

func (b *BaseChain) EthereumTypeGetStakingPoolsData(addrDesc AddressDescriptor) ([]StakingPoolData, error) {
	return nil, errors.New("not supported")
}

// EthereumTypeRpcCall calls eth_call with given data and to address
func (b *BaseChain) EthereumTypeRpcCall(data, to, from string) (string, error) {
	return "", errors.New("not supported")
}

func (b *BaseChain) EthereumTypeGetRawTransaction(txid string) (string, error) {
	return "", errors.New("not supported")
}

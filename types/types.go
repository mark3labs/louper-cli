package types

import "encoding/json"

type DiamondResponse struct {
	Chain      string  `json:"chain"`
	Diamond    Diamond `json:"diamond"`
	DiamondAbi []struct {
		Name            string `json:"name,omitempty"`
		Type            string `json:"type"`
		StateMutability string `json:"stateMutability,omitempty"`
		Inputs          []any  `json:"inputs,omitempty"`
		Outputs         []any  `json:"outputs,omitempty"`
		Anonymous       bool   `json:"anonymous,omitempty"`
	} `json:"diamondAbi"`
}

type Diamond struct {
	Name string `json:"name"`
	Abi  []struct {
		StateMutability string `json:"stateMutability,omitempty"`
		Type            string `json:"type"`
		Name            string `json:"name,omitempty"`
		Inputs          []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
		} `json:"inputs,omitempty"`
	} `json:"abi"`
	Address string `json:"address"`
	Facets  []struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Abi     []struct {
			Name            string `json:"name"`
			Type            string `json:"type"`
			StateMutability string `json:"stateMutability,omitempty"`
			Inputs          []any  `json:"inputs"`
			Outputs         []any  `json:"outputs,omitempty"`
			Anonymous       bool   `json:"anonymous,omitempty"`
		} `json:"abi"`
	} `json:"facets"`
}

type Chain struct {
	BlockExplorers struct {
		Default struct {
			Name   string `json:"name"`
			URL    string `json:"url"`
			APIURL string `json:"apiUrl"`
		} `json:"default"`
	} `json:"blockExplorers"`
	Name    string `json:"name"`
	Network string `json:"network"`
	RPCUrls struct {
		Public struct {
			HTTP      []string `json:"http"`
			WebSocket []string `json:"webSocket"`
		} `json:"public"`
		Default struct {
			HTTP      []string `json:"http"`
			WebSocket []string `json:"webSocket"`
		} `json:"default"`
	} `json:"rpcUrls"`
	NativeCurrency struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
	} `json:"nativeCurrency"`
	ID      int  `json:"id"`
	Testnet bool `json:"testnet"`
}

type Chains map[string]Chain

type AbiResponse struct {
	Name string `json:"name"`
	Abi  []struct {
		Name            string        `json:"name"`
		Type            string        `json:"type"`
		StateMutability string        `json:"stateMutability,omitempty"`
		Inputs          []interface{} `json:"inputs"`
		Outputs         []interface{} `json:"outputs,omitempty"`
		Anonymous       bool          `json:"anonymous,omitempty"`
	} `json:"abi"`
}

type DiamondJson struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Facets  []struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"facets"`
}

type FacetJson struct {
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	Functions []FunctionJson `json:"functions"`
}

type FunctionJson struct {
	Name      string `json:"name"`
	Signature string `json:"signature"`
	Selector  string `json:"selector"`
}

func (d DiamondJson) String() string {
	j, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return ""
	}
	return string(j)
}

func (f FacetJson) String() string {
	j, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return ""
	}
	return string(j)
}

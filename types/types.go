package types

type Diamond struct {
	Chain   string `json:"chain"`
	Diamond struct {
		Name string `json:"name"`
		Abi  []struct {
			Inputs []struct {
				InternalType string `json:"internalType"`
				Name         string `json:"name"`
				Type         string `json:"type"`
			} `json:"inputs,omitempty"`
			StateMutability string `json:"stateMutability,omitempty"`
			Type            string `json:"type"`
			Name            string `json:"name,omitempty"`
		} `json:"abi"`
		Address string `json:"address"`
		Facets  []struct {
			Name string `json:"name"`
			Abi  []struct {
				Inputs          []any  `json:"inputs"`
				Name            string `json:"name"`
				Type            string `json:"type"`
				Anonymous       bool   `json:"anonymous,omitempty"`
				Outputs         []any  `json:"outputs,omitempty"`
				StateMutability string `json:"stateMutability,omitempty"`
			} `json:"abi"`
			Address string `json:"address"`
		} `json:"facets"`
	} `json:"diamond"`
	DiamondAbi []struct {
		Inputs          []any  `json:"inputs,omitempty"`
		Name            string `json:"name,omitempty"`
		Type            string `json:"type"`
		Anonymous       bool   `json:"anonymous,omitempty"`
		Outputs         []any  `json:"outputs,omitempty"`
		StateMutability string `json:"stateMutability,omitempty"`
	} `json:"diamondAbi"`
}

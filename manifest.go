package assets

type Manifest map[string]ManifestEntry

type ManifestEntry struct {
	File           string   `json:"file"`
	Src            string   `json:"src"`
	IsEntry        bool     `json:"isEntry"`
	Css            []string `json:"css"`
	DynamicImports []string `json:"dynamicImports"`
}

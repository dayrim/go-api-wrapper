package cafa

type Configuration struct {
	//Application (s) name that uses CAFA to store configs.
	Application string `json:"application" example:"erply-to-woo-sync"`
	// Level is the operating level where the config will be stored. Has to be one of Company, Warehouse, Pos, User
	Level string `json:"level" example:"Pos"`
	//LevelID is an unique value, used to identify the instance of the level.
	//for User level, user id
	//for Pos - Point of Sale ID. Each Pos belongs to Warehouse. When we query CAFA for configuration on level: 'Pos', and add 'Look-Deeper: true' header, if CAFA do not find any configuration with the query parameters, CAFA will use the level_id(PointOfSaleID), to find which Warehouse the Pos belongs to and check for configuration with the same name(and type if present)
	//Warehouse can use any unique string to identify the instance of its level, beacuse the higher level for each Warehouse is Company level, which can only be one insatnce per account
	//Company level does not need level_id because there can only be one company per account in Erply.
	LevelID string `json:"level_id" example:"1"`
	// Type is an optional grouping property, in this case its 'payment',
	//indicating that this configuration belongs to payment integrations group
	//used to query groups of related configurations
	Type string `json:"type" example:"payment"`
	//Name of the configuration
	Name string `json:"name" example:"adyen"`
	//Value stored by the configuration. Supported types: JSON, JSON array, string
	Value interface{} `json:"value" example:"{\"username\": \"tatata\", \"password\": \"password\"}"`
}

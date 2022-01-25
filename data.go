package data

import "time"

type productdata struct {
	ID          int    `json: "id"`
	NAME        string `json: "name"`
	COCOA       int    `json: "cocoa"`
	CALORIE     int    `json: "calorie"`
	INGREDIENTS string `json: "ingredients"`
	MFD         int    `json: "-"`
	EXPDT       int    `json: "-"`
}

//type prod []*productdata

//func (p* prod) tojson(w io.Writer) error {

//	enco := json.NewEncoder(w)
//	return enco.Encode(p)
//}

var Includes = []*productdata{
	&productdata{
		ID:          0001,
		NAME:        "DARKMELOWE",
		COCOA:       70,
		CALORIE:     6,
		INGREDIENTS: "cocoa solids, cocoa butter, sugar, milk solids",
		MFD:         time.Now().UTC().Day(),
		EXPDT:       20 / 01 / 2025,
	},
	&productdata{
		ID:          0002,
		NAME:        "DARKNUTEE",
		COCOA:       70,
		CALORIE:     6,
		INGREDIENTS: "cocoa solids, cocoa butter, sugar, milk solids",
		MFD:         time.Now().UTC().Day(),
		EXPDT:       20 / 01 / 2025,
	},
}

func export_Data() []*productdata {

	return Includes

}

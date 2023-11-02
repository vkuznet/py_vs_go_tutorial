package main

type Dataset struct {
	Dataset string `json:"dataset"`
}

type PrimDS struct {
	PrimaryDSName string `json:"primary_ds_name"`
}

type File struct {
	LFN string `json:"logical_file_name"`
}

type Data struct {
	Dataset Dataset `json:"dataset`
	PrimDS  PrimDS  `json:"primds"`
	Files   []File  `json:"files"`
}

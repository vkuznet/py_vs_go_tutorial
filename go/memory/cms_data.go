package main

type Data struct {
	Dataset        Dataset        `json:"dataset`
	Block          Block          `json:"block"`
	ProcessingEra  ProcessingEra  `json:"processing_era"`
	AcquisitionEra AcquisitionEra `json:"acquisition_era"`
	PrimDS         PrimDS         `json:"primds"`
	Files          []File         `json:"files"`
	DatasetConf    []DatasetConf  `json:"dataset_conf_list"`
	FileConf       []FileConf     `json:"file_conf_list"`
}

type Dataset struct {
	GName      string `json:"physics_group_name"`
	PName      string `json:"processed_ds_name"`
	Tier       string `json:"data_tier_name"`
	AccessType string `json:"dataset_access_type"`
	Dataset    string `json:"dataset"`
	CreateBy   string `json:"create_by"`
	ModifiedBy string `json:"last_modified_by"`
	CDate      int64  `json:"creation_date" validate:"required,number,gt=0"`
	LDate      int64  `json:"last_modification_date" validate:"required,number,gt=0"`
}

type PrimDS struct {
	Name         string `json:"primary_ds_name"`
	Type         string `json:"primary_ds_type"`
	CreateBy     string `json:"create_by"`
	CreationDate int64  `json:"creation_date"`
}

type File struct {
	FileType         string     `json:"file_type"`
	Lfn              string     `json:"logical_file_name"`
	FileSize         int64      `json:"file_size"`
	EventCount       int64      `json:"event_count"`
	LastModifiedBy   string     `json:"last_modified_by"`
	LastModifiedDate int64      `json:"last_modification-date"`
	CrossSection     float64    `json:"auto_cross_section"`
	CHeckSum         string     `json:"check_sum"`
	Adler32          string     `json:"adler32"`
	FileLumis        []FileLumi `json:"file_lumi_list"`
}

type FileLumi struct {
	LumiNumber int `json:"lumi_section_num"`
	RunNumber  int `json:"run_num"`
	EventCount int `json:"event_count"`
}

type DatasetConf struct {
	Release      string `json:"release_version"`
	PSet         string `json:"pset_hash"`
	AppName      string `json:"app_name"`
	OutputModule string `json:"output_module_label"`
	GlobalTag    string `json:"global_tag"`
}
type FileConf struct {
	Release      string `json:"release_version"`
	PSet         string `json:"pset_hash"`
	AppName      string `json:"app_name"`
	OutputModule string `json:"output_module_label"`
	GlobalTag    string `json:"global_tag"`
	Lfn          string `json:"lfn"`
}

type Block struct {
	BlockName         string   `json:"block_name"`
	SiteName          string   `json:"origin_site_name"`
	Open              int      `json:"open_for_writing"`
	CreateBy          string   `json:"create_by"`
	CreationDate      int64    `json:"creation_date"`
	BlockSize         int64    `json:"block_size"`
	FIleCount         int      `json:"file_count"`
	FileParentList    []string `json:"file_parent_list"`
	DatasetParentList []string `json:"dataset_parent_list"`
}

type ProcessingEra struct {
	Version     int    `json:"processing_version"`
	CreateBy    string `json:"create_by"`
	Description string `json:"description"`
}
type AcquisitionEra struct {
	Name string `json:"acquisition_era_name"`
	Date int64  `json:"start_date"`
}

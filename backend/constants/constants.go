package constants

const (
	RecordStatusINAU = "INAU"
	RecordStatusREVE = "REVE"
	RecordStatusRNAU = "RNAU"
	RecordStatusLIVE = "LIVE"
	RecordStatusHIS  = "HIS"

	RecordStatusParameterNAU     = "nau"
	RecordStatusParameterLIVE    = "live"
	RecordStatusParameterHIS     = "his"
	RecordStatusParameterFullNau = "full-nau"
	RecordStatusParameterPending = "pending"

	RecordStatusHold   = "IHLD"
	RecordStatusReject = "RJCT"
	RecordStatusRna    = "RNA"
	RecordStatusIna    = "INA"
)

const (
	PartnerRecidAll    = "ALL"
	PartnerRecidPublic = "PUBLIC"

	BranchCodeAll = "ALL"

	SUCCESS = "SUCCESS"
	ERROR   = "ERROR"
	SYSTEM  = "SYSTEM"
	PENDING = "PENDING"
	SEND    = "SEND"

	POST = "POST"
	PUT  = "PUT"
)

const (
	DefaultProfileImageMinio = "DEFAULT_PROFILE_MINIO"
	MargaChineseDefault      = "Marga-中文"

	ILIKE       = "ILIKE"
	LIKE        = "LIKE"
	BETWEEN     = "BETWEEN"
	Contains    = "@>"
	Placeholder = "?|"
	EXISTS      = "EXISTS"
	JSONB       = "JSONB"

	Age          = "age"
	Job          = "job"
	MargaChinese = "marga-chinese"
	MargaLatin   = "marga-latin"
	Interest     = "interest"
	Location     = "location"
	Jarak        = "jarak"
)

const (
	JobContract = "job_contract"
	Skill       = "skill"
	JobType     = "type"
)

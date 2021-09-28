package config

import "time"

type Config struct {
	AccessKey              string
	APNInfo                *APNInfo
	AssumeRole             *AssumeRole
	CallerDocumentationURL string
	CallerName             string
	DebugLogging           bool
	HTTPProxy              string
	IamEndpoint            string
	Insecure               bool
	MaxRetries             int
	Profile                string
	Region                 string
	SecretKey              string
	SharedCredentialsFiles []string
	SharedConfigFiles      []string
	SkipCredsValidation    bool
	SkipMetadataApiCheck   bool
	StsEndpoint            string
	Token                  string
}

type APNInfo struct {
	PartnerName string
	Products    []APNProduct
}

type APNProduct struct {
	Name    string
	Version string
	Comment string
}

type AssumeRole struct {
	RoleARN           string
	Duration          time.Duration
	ExternalID        string
	Policy            string
	PolicyARNs        []string
	SessionName       string
	Tags              map[string]string
	TransitiveTagKeys []string
}

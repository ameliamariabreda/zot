package errors

import "errors"

var (
	ErrBadConfig               = errors.New("config: invalid config")
	ErrCliBadConfig            = errors.New("cli: bad config")
	ErrRepoNotFound            = errors.New("repository: not found")
	ErrRepoIsNotDir            = errors.New("repository: not a directory")
	ErrRepoBadVersion          = errors.New("repository: unsupported layout version")
	ErrManifestNotFound        = errors.New("manifest: not found")
	ErrBadManifest             = errors.New("manifest: invalid contents")
	ErrBadIndex                = errors.New("index: invalid contents")
	ErrUploadNotFound          = errors.New("uploads: not found")
	ErrBadUploadRange          = errors.New("uploads: bad range")
	ErrBlobNotFound            = errors.New("blob: not found")
	ErrBadBlob                 = errors.New("blob: bad blob")
	ErrBadBlobDigest           = errors.New("blob: bad blob digest")
	ErrUnknownCode             = errors.New("error: unknown error code")
	ErrBadCACert               = errors.New("tls: invalid ca cert")
	ErrBadUser                 = errors.New("auth: non-existent user")
	ErrEntriesExceeded         = errors.New("ldap: too many entries returned")
	ErrLDAPEmptyPassphrase     = errors.New("ldap: empty passphrase")
	ErrLDAPBadConn             = errors.New("ldap: bad connection")
	ErrLDAPConfig              = errors.New("config: invalid LDAP configuration")
	ErrCacheRootBucket         = errors.New("cache: unable to create/update root bucket")
	ErrCacheNoBucket           = errors.New("cache: unable to find bucket")
	ErrCacheMiss               = errors.New("cache: miss")
	ErrRequireCred             = errors.New("ldap: bind credentials required")
	ErrInvalidCred             = errors.New("ldap: invalid credentials")
	ErrEmptyJSON               = errors.New("cli: config json is empty")
	ErrInvalidArgs             = errors.New("cli: Invalid Arguments")
	ErrInvalidFlagsCombination = errors.New("cli: Invalid combination of flags")
	ErrInvalidURL              = errors.New("cli: invalid URL format")
	ErrUnauthorizedAccess      = errors.New("auth: unauthorized access. check credentials")
	ErrCannotResetConfigKey    = errors.New("cli: cannot reset given config key")
	ErrConfigNotFound          = errors.New("cli: config with the given name does not exist")
	ErrNoURLProvided           = errors.New("cli: no URL provided in argument or via config")
	ErrIllegalConfigKey        = errors.New("cli: given config key is not allowed")
	ErrScanNotSupported        = errors.New("search: scanning of image media type not supported")
	ErrCLITimeout              = errors.New("cli: Query timed out while waiting for results")
	ErrDuplicateConfigName     = errors.New("cli: cli config name already added")
	ErrInvalidRoute            = errors.New("routes: invalid route prefix")
	ErrImgStoreNotFound        = errors.New("routes: image store not found corresponding to given route")
	ErrEmptyValue              = errors.New("cache: empty value")
	ErrEmptyRepoList           = errors.New("search: no repository found")
	ErrInvalidRepositoryName   = errors.New("routes: not a repository name")
	ErrSyncMissingCatalog      = errors.New("sync: couldn't fetch upstream registry's catalog")
	ErrMethodNotSupported      = errors.New("storage: method not supported")
	ErrInvalidMetric           = errors.New("metrics: invalid metric func")
	ErrInjected                = errors.New("test: injected failure")
	ErrSyncInvalidUpstreamURL  = errors.New("sync: upstream url not found in sync config")
	ErrRegistryNoContent       = errors.New("sync: could not find a Content that matches localRepo")
	ErrSyncSignatureNotFound   = errors.New("sync: couldn't find any upstream notary/cosign signatures")
	ErrSyncSignature           = errors.New("sync: couldn't get upstream notary/cosign signatures")
	ErrImageLintAnnotations    = errors.New("routes: lint checks failed")
	ErrParsingAuthHeader       = errors.New("auth: failed parsing authorization header")
	ErrBadType                 = errors.New("core: invalid type")
	ErrParsingHTTPHeader       = errors.New("routes: invalid HTTP header")
	ErrBadRange                = errors.New("storage: bad range")
	ErrManifestMetaNotFound    = errors.New("repodb: image metadata not found for given manifest digest")
	ErrRepoMetaNotFound        = errors.New("repodb: repo metadata not found for given repo name")
	ErrTypeAssertionFailed     = errors.New("storage: failed DatabaseDriver type assertion")
)

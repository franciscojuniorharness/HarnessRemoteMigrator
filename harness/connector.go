package harness

// Generated by https://quicktype.io

type Connector struct {
	Status        string        `json:"status"`
	Data          ConnectorData `json:"data"`
	MetaData      interface{}   `json:"metaData"`
	CorrelationID string        `json:"correlationId"`
}

type ConnectorData struct {
	Connector             ConnectorClass                 `json:"connector"`
	CreatedAt             int64                          `json:"createdAt"`
	LastModifiedAt        int64                          `json:"lastModifiedAt"`
	Status                ConnectorStatus                `json:"status"`
	ActivityDetails       ActivityDetails                `json:"activityDetails"`
	HarnessManaged        bool                           `json:"harnessManaged"`
	GitDetails            map[string]interface{}         `json:"gitDetails"`
	EntityValidityDetails ConnectorEntityValidityDetails `json:"entityValidityDetails"`
	GovernanceMetadata    interface{}                    `json:"governanceMetadata"`
	IsFavorite            interface{}                    `json:"isFavorite"`
}

type ActivityDetails struct {
	LastActivityTime int64 `json:"lastActivityTime"`
}

type ConnectorClass struct {
	Name              string        `json:"name"`
	Identifier        string        `json:"identifier"`
	Description       string        `json:"description"`
	AccountIdentifier string        `json:"accountIdentifier"`
	OrgIdentifier     interface{}   `json:"orgIdentifier"`
	ProjectIdentifier interface{}   `json:"projectIdentifier"`
	Tags              ConnectorTags `json:"tags"`
	Type              string        `json:"type"`
	Spec              ConnectorSpec `json:"spec"`
}

type ConnectorSpec struct {
	URL               string         `json:"url"`
	ValidationRepo    interface{}    `json:"validationRepo"`
	Authentication    Authentication `json:"authentication"`
	APIAccess         APIAccess      `json:"apiAccess"`
	DelegateSelectors []interface{}  `json:"delegateSelectors"`
	ExecuteOnDelegate bool           `json:"executeOnDelegate"`
	Type              string         `json:"type"`
}

type APIAccess struct {
	Type string        `json:"type"`
	Spec APIAccessSpec `json:"spec"`
}

type APIAccessSpec struct {
	TokenRef string `json:"tokenRef"`
}

type Authentication struct {
	Type string             `json:"type"`
	Spec AuthenticationSpec `json:"spec"`
}

type AuthenticationSpec struct {
	Type string   `json:"type"`
	Spec SpecSpec `json:"spec"`
}

type SpecSpec struct {
	Username    string      `json:"username"`
	UsernameRef interface{} `json:"usernameRef"`
	TokenRef    string      `json:"tokenRef"`
}

type ConnectorTags struct {
}

type ConnectorEntityValidityDetails struct {
	Valid       bool        `json:"valid"`
	InvalidYAML interface{} `json:"invalidYaml"`
}

type ConnectorStatus struct {
	Status          string      `json:"status"`
	ErrorSummary    interface{} `json:"errorSummary"`
	Errors          interface{} `json:"errors"`
	TestedAt        int64       `json:"testedAt"`
	LastTestedAt    int64       `json:"lastTestedAt"`
	LastConnectedAt int64       `json:"lastConnectedAt"`
}
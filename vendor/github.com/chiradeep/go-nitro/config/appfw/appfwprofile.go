package appfw

type Appfwprofile struct {
	Addcookieflags                             string      `json:"addcookieflags,omitempty"`
	Archivename                                string      `json:"archivename,omitempty"`
	Autodeployrule                             bool        `json:"autodeployrule,omitempty"`
	Bufferoverflowaction                       interface{} `json:"bufferoverflowaction,omitempty"`
	Bufferoverflowmaxcookielength              int         `json:"bufferoverflowmaxcookielength,omitempty"`
	Bufferoverflowmaxheaderlength              int         `json:"bufferoverflowmaxheaderlength,omitempty"`
	Bufferoverflowmaxurllength                 int         `json:"bufferoverflowmaxurllength,omitempty"`
	Builtin                                    bool        `json:"builtin,omitempty"`
	Canonicalizehtmlresponse                   string      `json:"canonicalizehtmlresponse,omitempty"`
	Checkrequestheaders                        string      `json:"checkrequestheaders,omitempty"`
	Comment                                    string      `json:"comment,omitempty"`
	Contenttypeaction                          interface{} `json:"contenttypeaction,omitempty"`
	Cookieconsistencyaction                    interface{} `json:"cookieconsistencyaction,omitempty"`
	Cookieencryption                           string      `json:"cookieencryption,omitempty"`
	Cookieproxying                             string      `json:"cookieproxying,omitempty"`
	Cookietransforms                           string      `json:"cookietransforms,omitempty"`
	Creditcard                                 interface{} `json:"creditcard,omitempty"`
	Creditcardaction                           interface{} `json:"creditcardaction,omitempty"`
	Creditcardmaxallowed                       int         `json:"creditcardmaxallowed,omitempty"`
	Creditcardxout                             string      `json:"creditcardxout,omitempty"`
	Crosssitescriptingaction                   interface{} `json:"crosssitescriptingaction,omitempty"`
	Crosssitescriptingcheckcompleteurls        string      `json:"crosssitescriptingcheckcompleteurls,omitempty"`
	Crosssitescriptingtransformunsafehtml      string      `json:"crosssitescriptingtransformunsafehtml,omitempty"`
	Csrftag                                    string      `json:"csrftag,omitempty"`
	Csrftagaction                              interface{} `json:"csrftagaction,omitempty"`
	Customsettings                             string      `json:"customsettings,omitempty"`
	Defaultcharset                             string      `json:"defaultcharset,omitempty"`
	Defaultfieldformatmaxlength                int         `json:"defaultfieldformatmaxlength,omitempty"`
	Defaultfieldformatminlength                int         `json:"defaultfieldformatminlength,omitempty"`
	Defaultfieldformattype                     string      `json:"defaultfieldformattype,omitempty"`
	Defaults                                   string      `json:"defaults,omitempty"`
	Denyurlaction                              interface{} `json:"denyurlaction,omitempty"`
	Dosecurecreditcardlogging                  string      `json:"dosecurecreditcardlogging,omitempty"`
	Dynamiclearning                            interface{} `json:"dynamiclearning,omitempty"`
	Enableformtagging                          string      `json:"enableformtagging,omitempty"`
	Errorurl                                   string      `json:"errorurl,omitempty"`
	Excludefileuploadfromchecks                string      `json:"excludefileuploadfromchecks,omitempty"`
	Exemptclosureurlsfromsecuritychecks        string      `json:"exemptclosureurlsfromsecuritychecks,omitempty"`
	Fieldconsistencyaction                     interface{} `json:"fieldconsistencyaction,omitempty"`
	Fieldformataction                          interface{} `json:"fieldformataction,omitempty"`
	Fileuploadmaxnum                           int         `json:"fileuploadmaxnum,omitempty"`
	Fileuploadtypesaction                      interface{} `json:"fileuploadtypesaction,omitempty"`
	Htmlerrorobject                            string      `json:"htmlerrorobject,omitempty"`
	Inspectcontenttypes                        interface{} `json:"inspectcontenttypes,omitempty"`
	Invalidpercenthandling                     string      `json:"invalidpercenthandling,omitempty"`
	Jsondosaction                              interface{} `json:"jsondosaction,omitempty"`
	Jsonerrorobject                            string      `json:"jsonerrorobject,omitempty"`
	Jsonsqlinjectionaction                     interface{} `json:"jsonsqlinjectionaction,omitempty"`
	Jsonsqlinjectiontype                       string      `json:"jsonsqlinjectiontype,omitempty"`
	Jsonxssaction                              interface{} `json:"jsonxssaction,omitempty"`
	Logeverypolicyhit                          string      `json:"logeverypolicyhit,omitempty"`
	Multipleheaderaction                       interface{} `json:"multipleheaderaction,omitempty"`
	Name                                       string      `json:"name,omitempty"`
	Optimizepartialreqs                        string      `json:"optimizepartialreqs,omitempty"`
	Percentdecoderecursively                   string      `json:"percentdecoderecursively,omitempty"`
	Postbodylimit                              int         `json:"postbodylimit,omitempty"`
	Postbodylimitsignature                     int         `json:"postbodylimitsignature,omitempty"`
	Refererheadercheck                         string      `json:"refererheadercheck,omitempty"`
	Requestcontenttype                         string      `json:"requestcontenttype,omitempty"`
	Responsecontenttype                        string      `json:"responsecontenttype,omitempty"`
	Rfcprofile                                 string      `json:"rfcprofile,omitempty"`
	Semicolonfieldseparator                    string      `json:"semicolonfieldseparator,omitempty"`
	Sessionlessfieldconsistency                string      `json:"sessionlessfieldconsistency,omitempty"`
	Sessionlessurlclosure                      string      `json:"sessionlessurlclosure,omitempty"`
	Signatures                                 string      `json:"signatures,omitempty"`
	Sqlinjectionaction                         interface{} `json:"sqlinjectionaction,omitempty"`
	Sqlinjectionchecksqlwildchars              string      `json:"sqlinjectionchecksqlwildchars,omitempty"`
	Sqlinjectiononlycheckfieldswithsqlchars    string      `json:"sqlinjectiononlycheckfieldswithsqlchars,omitempty"`
	Sqlinjectionparsecomments                  string      `json:"sqlinjectionparsecomments,omitempty"`
	Sqlinjectiontransformspecialchars          string      `json:"sqlinjectiontransformspecialchars,omitempty"`
	Sqlinjectiontype                           string      `json:"sqlinjectiontype,omitempty"`
	Starturlaction                             interface{} `json:"starturlaction,omitempty"`
	Starturlclosure                            string      `json:"starturlclosure,omitempty"`
	State                                      string      `json:"state,omitempty"`
	Streaming                                  string      `json:"streaming,omitempty"`
	Stripcomments                              string      `json:"stripcomments,omitempty"`
	Striphtmlcomments                          string      `json:"striphtmlcomments,omitempty"`
	Stripxmlcomments                           string      `json:"stripxmlcomments,omitempty"`
	Trace                                      string      `json:"trace,omitempty"`
	Type                                       interface{} `json:"type,omitempty"`
	Urldecoderequestcookies                    string      `json:"urldecoderequestcookies,omitempty"`
	Usehtmlerrorobject                         string      `json:"usehtmlerrorobject,omitempty"`
	Verboseloglevel                            string      `json:"verboseloglevel,omitempty"`
	Xmlattachmentaction                        interface{} `json:"xmlattachmentaction,omitempty"`
	Xmldosaction                               interface{} `json:"xmldosaction,omitempty"`
	Xmlerrorobject                             string      `json:"xmlerrorobject,omitempty"`
	Xmlformataction                            interface{} `json:"xmlformataction,omitempty"`
	Xmlsoapfaultaction                         interface{} `json:"xmlsoapfaultaction,omitempty"`
	Xmlsqlinjectionaction                      interface{} `json:"xmlsqlinjectionaction,omitempty"`
	Xmlsqlinjectionchecksqlwildchars           string      `json:"xmlsqlinjectionchecksqlwildchars,omitempty"`
	Xmlsqlinjectiononlycheckfieldswithsqlchars string      `json:"xmlsqlinjectiononlycheckfieldswithsqlchars,omitempty"`
	Xmlsqlinjectionparsecomments               string      `json:"xmlsqlinjectionparsecomments,omitempty"`
	Xmlsqlinjectiontype                        string      `json:"xmlsqlinjectiontype,omitempty"`
	Xmlvalidationaction                        interface{} `json:"xmlvalidationaction,omitempty"`
	Xmlwsiaction                               interface{} `json:"xmlwsiaction,omitempty"`
	Xmlxssaction                               interface{} `json:"xmlxssaction,omitempty"`
}

package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// WindowsUpdateIdentity represents the WindowsUpdateIdentity schema from the OpenAPI specification
type WindowsUpdateIdentity struct {
	Updateid string `json:"updateId,omitempty"` // The revision independent identifier of the update.
	Revision int `json:"revision,omitempty"` // The revision number of the update.
}

// V1DeploymentNote represents the V1DeploymentNote schema from the OpenAPI specification
type V1DeploymentNote struct {
	Resourceuri []string `json:"resourceUri,omitempty"` // Required. Resource URI for the artifact being deployed.
}

// V1SlsaProvenanceZeroTwoSlsaMetadata represents the V1SlsaProvenanceZeroTwoSlsaMetadata schema from the OpenAPI specification
type V1SlsaProvenanceZeroTwoSlsaMetadata struct {
	Reproducible bool `json:"reproducible,omitempty"`
	Buildfinishedon string `json:"buildFinishedOn,omitempty"`
	Buildinvocationid string `json:"buildInvocationId,omitempty"`
	Buildstartedon string `json:"buildStartedOn,omitempty"`
	Completeness V1SlsaProvenanceZeroTwoSlsaCompleteness `json:"completeness,omitempty"` // Indicates that the builder claims certain fields in this message to be complete.
}

// ProtobufAny represents the ProtobufAny schema from the OpenAPI specification
type ProtobufAny struct {
	TypeField string `json:"@type,omitempty"` // A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one "/" character. The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading "." is not accepted). In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows: * If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield a [google.protobuf.Type][] value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the URL, or have them precompiled into a binary to avoid any lookup. Therefore, binary compatibility needs to be preserved on changes to types. (Use versioned type names to manage breaking changes.) Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com. Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics.
}

// V1Source represents the V1Source schema from the OpenAPI specification
type V1Source struct {
	Artifactstoragesourceuri string `json:"artifactStorageSourceUri,omitempty"` // If provided, the input binary artifacts for the build came from this location.
	Context V1SourceContext `json:"context,omitempty"` // A SourceContext is a reference to a tree of files. A SourceContext together with a path point to a unique revision of a single file or directory.
	Filehashes map[string]interface{} `json:"fileHashes,omitempty"` // Hash(es) of the build source, which can be used to verify that the original source integrity was maintained in the build. The keys to this map are file paths used as build source and the values contain the hash values for those files. If the build source came in a single package such as a gzipped tarfile (.tar.gz), the FileHash will be for the single path to that file.
	Additionalcontexts []V1SourceContext `json:"additionalContexts,omitempty"` // If provided, some of the source code used for the build may be found in these locations, in the case where the source repository had multiple remotes or submodules. This list will not include the context specified in the context field.
}

// V1UpgradeOccurrence represents the V1UpgradeOccurrence schema from the OpenAPI specification
type V1UpgradeOccurrence struct {
	Parsedversion V1Version `json:"parsedVersion,omitempty"` // Version contains structured information about the version of a package.
	Windowsupdate V1WindowsUpdate `json:"windowsUpdate,omitempty"` // Windows Update represents the metadata about the update for the Windows operating system. The fields in this message come from the Windows Update API documented at https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdate.
	Distribution V1UpgradeDistribution `json:"distribution,omitempty"` // The Upgrade Distribution represents metadata about the Upgrade for each operating system (CPE). Some distributions have additional metadata around updates, classifying them into various categories and severities.
	PackageField string `json:"package,omitempty"` // Required for non-Windows OS. The package this Upgrade is for.
}

// V1SbomReferenceIntotoPayload represents the V1SbomReferenceIntotoPayload schema from the OpenAPI specification
type V1SbomReferenceIntotoPayload struct {
	Subject []V1Subject `json:"subject,omitempty"` // Set of software artifacts that the attestation applies to. Each element represents a single software artifact.
	TypeField string `json:"_type,omitempty"` // Identifier for the schema of the Statement.
	Predicate V1SbomReferenceIntotoPredicate `json:"predicate,omitempty"` // A predicate which describes the SBOM being referenced.
	Predicatetype string `json:"predicateType,omitempty"` // URI identifying the type of the Predicate.
}

// V1SecretLocation represents the V1SecretLocation schema from the OpenAPI specification
type V1SecretLocation struct {
	Filelocation V1FileLocation `json:"fileLocation,omitempty"` // Indicates the location at which a package was found.
}

// SlsaProvenanceZeroTwoSlsaConfigSource represents the SlsaProvenanceZeroTwoSlsaConfigSource schema from the OpenAPI specification
type SlsaProvenanceZeroTwoSlsaConfigSource struct {
	Digest map[string]interface{} `json:"digest,omitempty"`
	Entrypoint string `json:"entryPoint,omitempty"`
	Uri string `json:"uri,omitempty"`
}

// WindowsUpdateCategory represents the WindowsUpdateCategory schema from the OpenAPI specification
type WindowsUpdateCategory struct {
	Name string `json:"name,omitempty"` // The localized name of the category.
	Categoryid string `json:"categoryId,omitempty"` // The identifier of the category.
}

// V1BaseImage represents the V1BaseImage schema from the OpenAPI specification
type V1BaseImage struct {
	Repository string `json:"repository,omitempty"` // The repository name in which the base image is from.
	Layercount int `json:"layerCount,omitempty"` // The number of layers that the base image is composed of.
	Name string `json:"name,omitempty"` // The name of the base image.
}

// V1ComplianceOccurrence represents the V1ComplianceOccurrence schema from the OpenAPI specification
type V1ComplianceOccurrence struct {
	Noncompliancereason string `json:"nonComplianceReason,omitempty"`
	Noncompliantfiles []V1NonCompliantFile `json:"nonCompliantFiles,omitempty"`
	Version V1ComplianceVersion `json:"version,omitempty"` // Describes the CIS benchmark version that is applicable to a given OS and os version.
}

// V1SBOMReferenceNote represents the V1SBOMReferenceNote schema from the OpenAPI specification
type V1SBOMReferenceNote struct {
	Format string `json:"format,omitempty"` // The format that SBOM takes. E.g. may be spdx, cyclonedx, etc...
	Version string `json:"version,omitempty"` // The version of the format that the SBOM takes. E.g. if the format is spdx, the version may be 2.3.
}

// ComplianceNoteCisBenchmark represents the ComplianceNoteCisBenchmark schema from the OpenAPI specification
type ComplianceNoteCisBenchmark struct {
	Severity string `json:"severity,omitempty"` // Note provider assigned severity/impact ranking. - SEVERITY_UNSPECIFIED: Unknown. - MINIMAL: Minimal severity. - LOW: Low severity. - MEDIUM: Medium severity. - HIGH: High severity. - CRITICAL: Critical severity.
	Profilelevel int `json:"profileLevel,omitempty"`
}

// V1SecretNote represents the V1SecretNote schema from the OpenAPI specification
type V1SecretNote struct {
}

// V1DSSEAttestationOccurrence represents the V1DSSEAttestationOccurrence schema from the OpenAPI specification
type V1DSSEAttestationOccurrence struct {
	Envelope V1Envelope `json:"envelope,omitempty"` // MUST match https://github.com/secure-systems-lab/dsse/blob/master/envelope.proto. An authenticated message of arbitrary type.
	Statement V1InTotoStatement `json:"statement,omitempty"` // Spec defined at https://github.com/in-toto/attestation/tree/main/spec#statement The serialized InTotoStatement will be stored as Envelope.payload. Envelope.payloadType is always "application/vnd.in-toto+json".
}

// V1Envelope represents the V1Envelope schema from the OpenAPI specification
type V1Envelope struct {
	Signatures []V1EnvelopeSignature `json:"signatures,omitempty"`
	Payload string `json:"payload,omitempty"`
	Payloadtype string `json:"payloadType,omitempty"`
}

// V1InTotoStatement represents the V1InTotoStatement schema from the OpenAPI specification
type V1InTotoStatement struct {
	Provenance V1InTotoProvenance `json:"provenance,omitempty"`
	Slsaprovenance V1SlsaProvenance `json:"slsaProvenance,omitempty"`
	Slsaprovenancezerotwo V1SlsaProvenanceZeroTwo `json:"slsaProvenanceZeroTwo,omitempty"`
	Subject []V1Subject `json:"subject,omitempty"`
	TypeField string `json:"_type,omitempty"` // Always `https://in-toto.io/Statement/v0.1`.
	Predicatetype string `json:"predicateType,omitempty"` // `https://slsa.dev/provenance/v0.1` for SlsaProvenance.
}

// V1Layer represents the V1Layer schema from the OpenAPI specification
type V1Layer struct {
	Arguments string `json:"arguments,omitempty"` // The recovered arguments to the Dockerfile directive.
	Directive string `json:"directive,omitempty"` // Required. The recovered Dockerfile directive used to construct this layer. See https://docs.docker.com/engine/reference/builder/ for more information.
}

// InTotoSlsaProvenanceV1SlsaProvenanceV1 represents the InTotoSlsaProvenanceV1SlsaProvenanceV1 schema from the OpenAPI specification
type InTotoSlsaProvenanceV1SlsaProvenanceV1 struct {
	Builddefinition InTotoSlsaProvenanceV1BuildDefinition `json:"buildDefinition,omitempty"`
	Rundetails InTotoSlsaProvenanceV1RunDetails `json:"runDetails,omitempty"`
}

// V1Metadata represents the V1Metadata schema from the OpenAPI specification
type V1Metadata struct {
	Buildfinishedon string `json:"buildFinishedOn,omitempty"` // The timestamp of when the build completed.
	Buildinvocationid string `json:"buildInvocationId,omitempty"` // Identifies the particular build invocation, which can be useful for finding associated logs or other ad-hoc analysis. The value SHOULD be globally unique, per in-toto Provenance spec.
	Buildstartedon string `json:"buildStartedOn,omitempty"` // The timestamp of when the build started.
	Completeness V1Completeness `json:"completeness,omitempty"` // Indicates that the builder claims certain fields in this message to be complete.
	Reproducible bool `json:"reproducible,omitempty"` // If true, the builder claims that running the recipe on materials will produce bit-for-bit identical output.
}

// V1License represents the V1License schema from the OpenAPI specification
type V1License struct {
	Comments string `json:"comments,omitempty"`
	Expression string `json:"expression,omitempty"` // Often a single license can be used to represent the licensing terms. Sometimes it is necessary to include a choice of one or more licenses or some combination of license identifiers. Examples: "LGPL-2.1-only OR MIT", "LGPL-2.1-only AND MIT", "GPL-2.0-or-later WITH Bison-exception-2.2".
}

// V1InTotoSlsaProvenanceV1ResourceDescriptor represents the V1InTotoSlsaProvenanceV1ResourceDescriptor schema from the OpenAPI specification
type V1InTotoSlsaProvenanceV1ResourceDescriptor struct {
	Downloadlocation string `json:"downloadLocation,omitempty"`
	Mediatype string `json:"mediaType,omitempty"`
	Name string `json:"name,omitempty"`
	Uri string `json:"uri,omitempty"`
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	Content string `json:"content,omitempty"`
	Digest map[string]interface{} `json:"digest,omitempty"`
}

// VulnerabilityAssessmentNoteAssessment represents the VulnerabilityAssessmentNoteAssessment schema from the OpenAPI specification
type VulnerabilityAssessmentNoteAssessment struct {
	Shortdescription string `json:"shortDescription,omitempty"` // A one sentence description of this Vex.
	Cve string `json:"cve,omitempty"` // Holds the MITRE standard Common Vulnerabilities and Exposures (CVE) tracking number for the vulnerability. Deprecated: Use vulnerability_id instead to denote CVEs.
	Relateduris []V1RelatedUrl `json:"relatedUris,omitempty"` // Holds a list of references associated with this vulnerability item and assessment. These uris have additional information about the vulnerability and the assessment itself. E.g. Link to a document which details how this assessment concluded the state of this vulnerability.
	State string `json:"state,omitempty"` // Provides the state of this Vulnerability assessment. - STATE_UNSPECIFIED: No state is specified. - AFFECTED: This product is known to be affected by this vulnerability. - NOT_AFFECTED: This product is known to be not affected by this vulnerability. - FIXED: This product contains a fix for this vulnerability. - UNDER_INVESTIGATION: It is not known yet whether these versions are or are not affected by the vulnerability. However, it is still under investigation.
	Vulnerabilityid string `json:"vulnerabilityId,omitempty"` // The vulnerability identifier for this Assessment. Will hold one of common identifiers e.g. CVE, GHSA etc.
	Justification AssessmentJustification `json:"justification,omitempty"` // Justification provides the justification when the state of the assessment if NOT_AFFECTED.
	Remediations []AssessmentRemediation `json:"remediations,omitempty"` // Specifies details on how to handle (and presumably, fix) a vulnerability.
	Impacts []string `json:"impacts,omitempty"` // Contains information about the impact of this vulnerability, this will change with time.
	Longdescription string `json:"longDescription,omitempty"` // A detailed description of this Vex.
}

// ProjectListProjectsResponse represents the ProjectListProjectsResponse schema from the OpenAPI specification
type ProjectListProjectsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The next pagination token in the list response. It should be used as `page_token` for the following request. An empty value means no more results.
	Projects []ProjectProject `json:"projects,omitempty"` // The projects requested.
}

// V1VulnerabilityNote represents the V1VulnerabilityNote schema from the OpenAPI specification
type V1VulnerabilityNote struct {
	Sourceupdatetime string `json:"sourceUpdateTime,omitempty"` // The time this information was last changed at the source. This is an upstream timestamp from the underlying information source - e.g. Ubuntu security tracker.
	Windowsdetails []VulnerabilityNoteWindowsDetail `json:"windowsDetails,omitempty"` // Windows details get their own format because the information format and model don't match a normal detail. Specifically Windows updates are done as patches, thus Windows vulnerabilities really are a missing package, rather than a package being at an incorrect version.
	Cvssscore float32 `json:"cvssScore,omitempty"` // The CVSS score of this vulnerability. CVSS score is on a scale of 0 - 10 where 0 indicates low severity and 10 indicates high severity.
	Cvssv2 V1CVSS `json:"cvssV2,omitempty"` // Common Vulnerability Scoring System. For details, see https://www.first.org/cvss/specification-document This is a message we will try to use for storing various versions of CVSS rather than making a separate proto for storing a specific version.
	Cvssv3 V1CVSSv3 `json:"cvssV3,omitempty"`
	Cvssversion string `json:"cvssVersion,omitempty"` // CVSS Version.
	Details []VulnerabilityNoteDetail `json:"details,omitempty"` // Details of all known distros and packages affected by this vulnerability.
	Severity string `json:"severity,omitempty"` // Note provider assigned severity/impact ranking. - SEVERITY_UNSPECIFIED: Unknown. - MINIMAL: Minimal severity. - LOW: Low severity. - MEDIUM: Medium severity. - HIGH: High severity. - CRITICAL: Critical severity.
}

// SlsaProvenanceSlsaRecipe represents the SlsaProvenanceSlsaRecipe schema from the OpenAPI specification
type SlsaProvenanceSlsaRecipe struct {
	Arguments ProtobufAny `json:"arguments,omitempty"` // `Any` contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a message in C++. Foo foo = ...; Any any; any.PackFrom(foo); ... if (any.UnpackTo(&foo)) { ... } Example 2: Pack and unpack a message in Java. Foo foo = ...; Any any = Any.pack(foo); ... if (any.is(Foo.class)) { foo = any.unpack(Foo.class); } Example 3: Pack and unpack a message in Python. foo = Foo(...) any = Any() any.Pack(foo) ... if any.Is(Foo.DESCRIPTOR): any.Unpack(foo) ... Example 4: Pack and unpack a message in Go foo := &pb.Foo{...} any, err := ptypes.MarshalAny(foo) ... foo := &pb.Foo{} if err := ptypes.UnmarshalAny(any, foo); err != nil { ... } The pack methods provided by protobuf library will by default use 'type.googleapis.com/full.type.name' as the type URL and the unpack methods only use the fully qualified type name after the last '/' in the type URL, for example "foo.bar.com/x/y.z" will yield type name "y.z". JSON ==== The JSON representation of an `Any` value uses the regular representation of the deserialized, embedded message, with an additional field `@type` which contains the type URL. Example: package google.profile; message Person { string first_name = 1; string last_name = 2; } { "@type": "type.googleapis.com/google.profile.Person", "firstName": <string>, "lastName": <string> } If the embedded message type is well-known and has a custom JSON representation, that representation will be embedded adding a field `value` which holds the custom JSON in addition to the `@type` field. Example (for message [google.protobuf.Duration][]): { "@type": "type.googleapis.com/google.protobuf.Duration", "value": "1.212s" }
	Definedinmaterial string `json:"definedInMaterial,omitempty"` // Index in materials containing the recipe steps that are not implied by recipe.type. For example, if the recipe type were "make", then this would point to the source containing the Makefile, not the make program itself. Set to -1 if the recipe doesn't come from a material, as zero is default unset value for int64.
	Entrypoint string `json:"entryPoint,omitempty"` // String identifying the entry point into the build. This is often a path to a configuration file and/or a target label within that file. The syntax and meaning are defined by recipe.type. For example, if the recipe type were "make", then this would reference the directory in which to run make as well as which target to use.
	Environment ProtobufAny `json:"environment,omitempty"` // `Any` contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a message in C++. Foo foo = ...; Any any; any.PackFrom(foo); ... if (any.UnpackTo(&foo)) { ... } Example 2: Pack and unpack a message in Java. Foo foo = ...; Any any = Any.pack(foo); ... if (any.is(Foo.class)) { foo = any.unpack(Foo.class); } Example 3: Pack and unpack a message in Python. foo = Foo(...) any = Any() any.Pack(foo) ... if any.Is(Foo.DESCRIPTOR): any.Unpack(foo) ... Example 4: Pack and unpack a message in Go foo := &pb.Foo{...} any, err := ptypes.MarshalAny(foo) ... foo := &pb.Foo{} if err := ptypes.UnmarshalAny(any, foo); err != nil { ... } The pack methods provided by protobuf library will by default use 'type.googleapis.com/full.type.name' as the type URL and the unpack methods only use the fully qualified type name after the last '/' in the type URL, for example "foo.bar.com/x/y.z" will yield type name "y.z". JSON ==== The JSON representation of an `Any` value uses the regular representation of the deserialized, embedded message, with an additional field `@type` which contains the type URL. Example: package google.profile; message Person { string first_name = 1; string last_name = 2; } { "@type": "type.googleapis.com/google.profile.Person", "firstName": <string>, "lastName": <string> } If the embedded message type is well-known and has a custom JSON representation, that representation will be embedded adding a field `value` which holds the custom JSON in addition to the `@type` field. Example (for message [google.protobuf.Duration][]): { "@type": "type.googleapis.com/google.protobuf.Duration", "value": "1.212s" }
	TypeField string `json:"type,omitempty"` // URI indicating what type of recipe was performed. It determines the meaning of recipe.entryPoint, recipe.arguments, recipe.environment, and materials.
}

// V1AttestationNote represents the V1AttestationNote schema from the OpenAPI specification
type V1AttestationNote struct {
	Hint AttestationNoteHint `json:"hint,omitempty"` // This submessage provides human-readable hints about the purpose of the authority. Because the name of a note acts as its resource reference, it is important to disambiguate the canonical name of the Note (which might be a UUID for security purposes) from "readable" names more suitable for debug output. Note that these hints should not be used to look up authorities in security sensitive contexts, such as when looking up attestations to verify.
}

// SlsaProvenanceMaterial represents the SlsaProvenanceMaterial schema from the OpenAPI specification
type SlsaProvenanceMaterial struct {
	Digest map[string]interface{} `json:"digest,omitempty"`
	Uri string `json:"uri,omitempty"`
}

// DiscoveryOccurrenceSBOMStatus represents the DiscoveryOccurrenceSBOMStatus schema from the OpenAPI specification
type DiscoveryOccurrenceSBOMStatus struct {
	Sbomstate string `json:"sbomState,omitempty"` // An enum indicating the progress of the SBOM generation. - SBOM_STATE_UNSPECIFIED: Default unknown state. - PENDING: SBOM scanning is pending. - COMPLETE: SBOM scanning has completed.
	ErrorField string `json:"error,omitempty"` // If there was an error generating an SBOM, this will indicate what that error was.
}

// V1SourceContext represents the V1SourceContext schema from the OpenAPI specification
type V1SourceContext struct {
	Cloudrepo V1CloudRepoSourceContext `json:"cloudRepo,omitempty"` // A CloudRepoSourceContext denotes a particular revision in a Google Cloud Source Repo.
	Gerrit V1GerritSourceContext `json:"gerrit,omitempty"` // A SourceContext referring to a Gerrit project.
	Git V1GitSourceContext `json:"git,omitempty"` // A GitSourceContext denotes a particular revision in a third party Git repository (e.g., GitHub).
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels with user defined metadata.
}

// V1BuildOccurrence represents the V1BuildOccurrence schema from the OpenAPI specification
type V1BuildOccurrence struct {
	Intotoprovenance V1InTotoProvenance `json:"intotoProvenance,omitempty"`
	Intotostatement V1InTotoStatement `json:"intotoStatement,omitempty"` // Spec defined at https://github.com/in-toto/attestation/tree/main/spec#statement The serialized InTotoStatement will be stored as Envelope.payload. Envelope.payloadType is always "application/vnd.in-toto+json".
	Provenance V1BuildProvenance `json:"provenance,omitempty"` // Provenance of a build. Contains all information needed to verify the full details about the build from source to completion.
	Provenancebytes string `json:"provenanceBytes,omitempty"` // Serialized JSON representation of the provenance, used in generating the build signature in the corresponding build note. After verifying the signature, `provenance_bytes` can be unmarshalled and compared to the provenance to confirm that it is unchanged. A base64-encoded string representation of the provenance bytes is used for the signature in order to interoperate with openssl which expects this format for signature verification. The serialized form is captured both to avoid ambiguity in how the provenance is marshalled to json as well to prevent incompatibilities with future changes.
	Intotoslsaprovenancev1 V1InTotoSlsaProvenanceV1 `json:"inTotoSlsaProvenanceV1,omitempty"`
}

// V1CVSSv3 represents the V1CVSSv3 schema from the OpenAPI specification
type V1CVSSv3 struct {
	Confidentialityimpact string `json:"confidentialityImpact,omitempty"`
	Privilegesrequired string `json:"privilegesRequired,omitempty"`
	Availabilityimpact string `json:"availabilityImpact,omitempty"`
	Scope string `json:"scope,omitempty"`
	Attackcomplexity string `json:"attackComplexity,omitempty"`
	Basescore float32 `json:"baseScore,omitempty"` // The base score is a function of the base metric scores.
	Exploitabilityscore float32 `json:"exploitabilityScore,omitempty"`
	Impactscore float32 `json:"impactScore,omitempty"`
	Integrityimpact string `json:"integrityImpact,omitempty"`
	Attackvector string `json:"attackVector,omitempty"`
	Userinteraction string `json:"userInteraction,omitempty"`
}

// InTotoSlsaProvenanceV1BuildMetadata represents the InTotoSlsaProvenanceV1BuildMetadata schema from the OpenAPI specification
type InTotoSlsaProvenanceV1BuildMetadata struct {
	Finishedon string `json:"finishedOn,omitempty"`
	Invocationid string `json:"invocationId,omitempty"`
	Startedon string `json:"startedOn,omitempty"`
}

// V1DiscoveryNote represents the V1DiscoveryNote schema from the OpenAPI specification
type V1DiscoveryNote struct {
	Analysiskind string `json:"analysisKind,omitempty"` // Kind represents the kinds of notes supported. - NOTE_KIND_UNSPECIFIED: Default value. This value is unused. - VULNERABILITY: The note and occurrence represent a package vulnerability. - BUILD: The note and occurrence assert build provenance. - IMAGE: This represents an image basis relationship. - PACKAGE: This represents a package installed via a package manager. - DEPLOYMENT: The note and occurrence track deployment events. - DISCOVERY: The note and occurrence track the initial discovery status of a resource. - ATTESTATION: This represents a logical "role" that can attest to artifacts. - UPGRADE: This represents an available package upgrade. - COMPLIANCE: This represents a Compliance Note - DSSE_ATTESTATION: This represents a DSSE attestation Note - VULNERABILITY_ASSESSMENT: This represents a Vulnerability Assessment. - SBOM_REFERENCE: This represents an SBOM Reference. - SECRET: This represents a secret.
}

// V1SecretOccurrence represents the V1SecretOccurrence schema from the OpenAPI specification
type V1SecretOccurrence struct {
	Locations []V1SecretLocation `json:"locations,omitempty"` // Locations where the secret is detected.
	Statuses []V1SecretStatus `json:"statuses,omitempty"` // Status of the secret.
	Kind string `json:"kind,omitempty"` // Kind of secret. - SECRET_KIND_UNSPECIFIED: Unspecified - SECRET_KIND_UNKNOWN: The secret kind is unknown. - SECRET_KIND_GCP_SERVICE_ACCOUNT_KEY: A GCP service account key per: https://cloud.google.com/iam/docs/creating-managing-service-account-keys
}

// V1BuildProvenance represents the V1BuildProvenance schema from the OpenAPI specification
type V1BuildProvenance struct {
	Buildoptions map[string]interface{} `json:"buildOptions,omitempty"` // Special options applied to this build. This is a catch-all field where build providers can enter any desired additional details.
	Createtime string `json:"createTime,omitempty"` // Time at which the build was created.
	Builderversion string `json:"builderVersion,omitempty"` // Version string of the builder at the time this build was executed.
	Creator string `json:"creator,omitempty"` // E-mail address of the user who initiated this build. Note that this was the user's e-mail address at the time the build was initiated; this address may not represent the same end-user for all time.
	Endtime string `json:"endTime,omitempty"` // Time at which execution of the build was finished.
	Id string `json:"id,omitempty"` // Required. Unique identifier of the build.
	Logsuri string `json:"logsUri,omitempty"` // URI where any logs for this provenance were written.
	Sourceprovenance V1Source `json:"sourceProvenance,omitempty"` // Source describes the location of the source used for the build.
	Builtartifacts []V1Artifact `json:"builtArtifacts,omitempty"` // Output of the build.
	Commands []V1Command `json:"commands,omitempty"` // Commands requested by the build.
	Projectid string `json:"projectId,omitempty"` // ID of the project.
	Starttime string `json:"startTime,omitempty"` // Time at which execution of the build was started.
	Triggerid string `json:"triggerId,omitempty"` // Trigger identifier if the build was triggered automatically; empty if not.
}

// V1BuildNote represents the V1BuildNote schema from the OpenAPI specification
type V1BuildNote struct {
	Builderversion string `json:"builderVersion,omitempty"` // Required. Immutable. Version of the builder which produced this build.
}

// V1GitSourceContext represents the V1GitSourceContext schema from the OpenAPI specification
type V1GitSourceContext struct {
	Revisionid string `json:"revisionId,omitempty"` // Git commit hash.
	Url string `json:"url,omitempty"` // Git repository URL.
}

// V1ListOccurrencesResponse represents the V1ListOccurrencesResponse schema from the OpenAPI specification
type V1ListOccurrencesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The next pagination token in the list response. It should be used as `page_token` for the following request. An empty value means no more results.
	Occurrences []V1Occurrence `json:"occurrences,omitempty"` // The occurrences requested.
}

// VulnerabilityNoteDetail represents the VulnerabilityNoteDetail schema from the OpenAPI specification
type VulnerabilityNoteDetail struct {
	Fixedpackage string `json:"fixedPackage,omitempty"` // The distro recommended package to update to that contains a fix for this vulnerability. It is possible for this to be different from the affected_package.
	Source string `json:"source,omitempty"` // The source from which the information in this Detail was obtained.
	Vendor string `json:"vendor,omitempty"` // The name of the vendor of the product.
	Description string `json:"description,omitempty"` // A vendor-specific description of this vulnerability.
	Fixedcpeuri string `json:"fixedCpeUri,omitempty"` // The distro recommended [CPE URI](https://cpe.mitre.org/specification/) to update to that contains a fix for this vulnerability. It is possible for this to be different from the affected_cpe_uri.
	Packagetype string `json:"packageType,omitempty"` // The type of package; whether native or non native (e.g., ruby gems, node.js packages, etc.).
	Severityname string `json:"severityName,omitempty"` // The distro assigned severity of this vulnerability.
	Sourceupdatetime string `json:"sourceUpdateTime,omitempty"` // The time this information was last changed at the source. This is an upstream timestamp from the underlying information source - e.g. Ubuntu security tracker.
	Isobsolete bool `json:"isObsolete,omitempty"` // Whether this detail is obsolete. Occurrences are expected not to point to obsolete details.
	Affectedversionend V1Version `json:"affectedVersionEnd,omitempty"` // Version contains structured information about the version of a package.
	Affectedversionstart V1Version `json:"affectedVersionStart,omitempty"` // Version contains structured information about the version of a package.
	Fixedversion V1Version `json:"fixedVersion,omitempty"` // Version contains structured information about the version of a package.
	Affectedcpeuri string `json:"affectedCpeUri,omitempty"` // Required. The [CPE URI](https://cpe.mitre.org/specification/) this vulnerability affects.
	Affectedpackage string `json:"affectedPackage,omitempty"` // Required. The package this vulnerability affects.
}

// V1WindowsUpdate represents the V1WindowsUpdate schema from the OpenAPI specification
type V1WindowsUpdate struct {
	Description string `json:"description,omitempty"` // The localized description of the update.
	Identity WindowsUpdateIdentity `json:"identity,omitempty"` // The unique identifier of the update.
	Kbarticleids []string `json:"kbArticleIds,omitempty"` // The Microsoft Knowledge Base article IDs that are associated with the update.
	Lastpublishedtimestamp string `json:"lastPublishedTimestamp,omitempty"` // The last published timestamp of the update.
	Supporturl string `json:"supportUrl,omitempty"` // The hyperlink to the support information for the update.
	Title string `json:"title,omitempty"` // The localized title of the update.
	Categories []WindowsUpdateCategory `json:"categories,omitempty"` // The list of categories to which the update belongs.
}

// V1AttestationOccurrence represents the V1AttestationOccurrence schema from the OpenAPI specification
type V1AttestationOccurrence struct {
	Signatures []V1Signature `json:"signatures,omitempty"` // One or more signatures over `serialized_payload`. Verifier implementations should consider this attestation message verified if at least one `signature` verifies `serialized_payload`. See `Signature` in common.proto for more details on signature structure and verification.
	Jwts []V1Jwt `json:"jwts,omitempty"` // One or more JWTs encoding a self-contained attestation. Each JWT encodes the payload that it verifies within the JWT itself. Verifier implementation SHOULD ignore the `serialized_payload` field when verifying these JWTs. If only JWTs are present on this AttestationOccurrence, then the `serialized_payload` SHOULD be left empty. Each JWT SHOULD encode a claim specific to the `resource_uri` of this Occurrence, but this is not validated by Grafeas metadata API implementations. The JWT itself is opaque to Grafeas.
	Serializedpayload string `json:"serializedPayload,omitempty"` // Required. The serialized payload that is verified by one or more `signatures`.
}

// V1ImageNote represents the V1ImageNote schema from the OpenAPI specification
type V1ImageNote struct {
	Resourceurl string `json:"resourceUrl,omitempty"` // Required. Immutable. The resource_url for the resource representing the basis of associated occurrence images.
	Fingerprint V1Fingerprint `json:"fingerprint,omitempty"` // A set of properties that uniquely identify a given Docker image.
}

// InTotoSlsaProvenanceV1RunDetails represents the InTotoSlsaProvenanceV1RunDetails schema from the OpenAPI specification
type InTotoSlsaProvenanceV1RunDetails struct {
	Byproducts []V1InTotoSlsaProvenanceV1ResourceDescriptor `json:"byproducts,omitempty"`
	Metadata InTotoSlsaProvenanceV1BuildMetadata `json:"metadata,omitempty"`
	Builder InTotoSlsaProvenanceV1ProvenanceBuilder `json:"builder,omitempty"`
}

// Grafeasv1Location represents the Grafeasv1Location schema from the OpenAPI specification
type Grafeasv1Location struct {
	Cpeuri string `json:"cpeUri,omitempty"`
	Path string `json:"path,omitempty"` // The path from which we gathered that this package/version is installed.
	Version V1Version `json:"version,omitempty"` // Version contains structured information about the version of a package.
}

// ProjectProject represents the ProjectProject schema from the OpenAPI specification
type ProjectProject struct {
	Name string `json:"name,omitempty"` // The name of the project in the form of `projects/{PROJECT_ID}`.
}

// V1RepoId represents the V1RepoId schema from the OpenAPI specification
type V1RepoId struct {
	Projectrepoid V1ProjectRepoId `json:"projectRepoId,omitempty"` // Selects a repo using a Google Cloud Platform project ID (e.g., winged-cargo-31) and a repo name within that project.
	Uid string `json:"uid,omitempty"` // A server-assigned, globally unique identifier.
}

// V1ComplianceVersion represents the V1ComplianceVersion schema from the OpenAPI specification
type V1ComplianceVersion struct {
	Version string `json:"version,omitempty"` // The version of the benchmark. This is set to the version of the OS-specific CIS document the benchmark is defined in.
	Benchmarkdocument string `json:"benchmarkDocument,omitempty"` // The name of the document that defines this benchmark, e.g. "CIS Container-Optimized OS".
	Cpeuri string `json:"cpeUri,omitempty"` // The CPE URI (https://cpe.mitre.org/specification/) this benchmark is applicable to.
}

// V1UpgradeDistribution represents the V1UpgradeDistribution schema from the OpenAPI specification
type V1UpgradeDistribution struct {
	Cpeuri string `json:"cpeUri,omitempty"` // Required - The specific operating system this metadata applies to. See https://cpe.mitre.org/specification/.
	Cve []string `json:"cve,omitempty"` // The cve tied to this Upgrade.
	Severity string `json:"severity,omitempty"` // The severity as specified by the upstream operating system.
	Classification string `json:"classification,omitempty"`
}

// V1Command represents the V1Command schema from the OpenAPI specification
type V1Command struct {
	Id string `json:"id,omitempty"` // Optional unique identifier for this command, used in wait_for to reference this command as a dependency.
	Name string `json:"name,omitempty"` // Required. Name of the command, as presented on the command line, or if the command is packaged as a Docker container, as presented to `docker pull`.
	Waitfor []string `json:"waitFor,omitempty"` // The ID(s) of the command(s) that this command depends on.
	Args []string `json:"args,omitempty"` // Command-line arguments used when executing this command.
	Dir string `json:"dir,omitempty"` // Working directory (relative to project source root) used when running this command.
	Env []string `json:"env,omitempty"` // Environment variables set before running this command.
}

// V1ImageOccurrence represents the V1ImageOccurrence schema from the OpenAPI specification
type V1ImageOccurrence struct {
	Layerinfo []V1Layer `json:"layerInfo,omitempty"` // This contains layer-specific metadata, if populated it has length "distance" and is ordered with [distance] being the layer immediately following the base image and [1] being the final layer.
	Baseresourceurl string `json:"baseResourceUrl,omitempty"` // Output only. This contains the base image URL for the derived image occurrence.
	Distance int `json:"distance,omitempty"` // Output only. The number of layers by which this image differs from the associated image basis.
	Fingerprint V1Fingerprint `json:"fingerprint,omitempty"` // A set of properties that uniquely identify a given Docker image.
}

// AssessmentRemediation represents the AssessmentRemediation schema from the OpenAPI specification
type AssessmentRemediation struct {
	Details string `json:"details,omitempty"` // Contains a comprehensive human-readable discussion of the remediation.
	Remediationtype string `json:"remediationType,omitempty"` // The type of remediation that can be applied. - REMEDIATION_TYPE_UNSPECIFIED: No remediation type specified. - MITIGATION: A MITIGATION is available. - NO_FIX_PLANNED: No fix is planned. - NONE_AVAILABLE: Not available. - VENDOR_FIX: A vendor fix is available. - WORKAROUND: A workaround is available.
	Remediationuri V1RelatedUrl `json:"remediationUri,omitempty"` // Metadata for any related URL information.
}

// V1Hash represents the V1Hash schema from the OpenAPI specification
type V1Hash struct {
	Value string `json:"value,omitempty"` // Required. The hash value.
	TypeField string `json:"type,omitempty"` // Required. The type of hash that was performed, e.g. "SHA-256".
}

// V1CloudRepoSourceContext represents the V1CloudRepoSourceContext schema from the OpenAPI specification
type V1CloudRepoSourceContext struct {
	Aliascontext V1AliasContext `json:"aliasContext,omitempty"` // An alias to a repo revision.
	Repoid V1RepoId `json:"repoId,omitempty"` // A unique identifier for a Cloud Repo.
	Revisionid string `json:"revisionId,omitempty"` // A revision ID.
}

// AssessmentJustification represents the AssessmentJustification schema from the OpenAPI specification
type AssessmentJustification struct {
	Justificationtype string `json:"justificationType,omitempty"` // Provides the type of justification. - JUSTIFICATION_TYPE_UNSPECIFIED: JUSTIFICATION_TYPE_UNSPECIFIED. - COMPONENT_NOT_PRESENT: The vulnerable component is not present in the product. - VULNERABLE_CODE_NOT_PRESENT: The vulnerable code is not present. Typically this case occurs when source code is configured or built in a way that excludes the vulnerable code. - VULNERABLE_CODE_NOT_IN_EXECUTE_PATH: The vulnerable code can not be executed. Typically this case occurs when the product includes the vulnerable code but does not call or use the vulnerable code. - VULNERABLE_CODE_CANNOT_BE_CONTROLLED_BY_ADVERSARY: The vulnerable code cannot be controlled by an attacker to exploit the vulnerability. - INLINE_MITIGATIONS_ALREADY_EXIST: The product includes built-in protections or features that prevent exploitation of the vulnerability. These built-in protections cannot be subverted by the attacker and cannot be configured or disabled by the user. These mitigations completely prevent exploitation based on known attack vectors.
	Details string `json:"details,omitempty"` // Additional details on why this justification was chosen.
}

// V1SlsaProvenanceZeroTwo represents the V1SlsaProvenanceZeroTwo schema from the OpenAPI specification
type V1SlsaProvenanceZeroTwo struct {
	Buildtype string `json:"buildType,omitempty"`
	Builder V1SlsaProvenanceZeroTwoSlsaBuilder `json:"builder,omitempty"` // Identifies the entity that executed the recipe, which is trusted to have correctly performed the operation and populated this provenance.
	Invocation SlsaProvenanceZeroTwoSlsaInvocation `json:"invocation,omitempty"` // Identifies the event that kicked off the build.
	Materials []SlsaProvenanceZeroTwoSlsaMaterial `json:"materials,omitempty"`
	Metadata V1SlsaProvenanceZeroTwoSlsaMetadata `json:"metadata,omitempty"` // Other properties of the build.
	Buildconfig map[string]interface{} `json:"buildConfig,omitempty"`
}

// V1SbomReferenceIntotoPredicate represents the V1SbomReferenceIntotoPredicate schema from the OpenAPI specification
type V1SbomReferenceIntotoPredicate struct {
	Digest map[string]interface{} `json:"digest,omitempty"` // A map of algorithm to digest of the contents of the SBOM.
	Location string `json:"location,omitempty"` // The location of the SBOM.
	Mimetype string `json:"mimeType,omitempty"` // The mime type of the SBOM.
	Referrerid string `json:"referrerId,omitempty"` // The person or system referring this predicate to the consumer.
}

// V1NonCompliantFile represents the V1NonCompliantFile schema from the OpenAPI specification
type V1NonCompliantFile struct {
	Displaycommand string `json:"displayCommand,omitempty"` // Command to display the non-compliant files.
	Path string `json:"path,omitempty"` // Empty if `display_command` is set.
	Reason string `json:"reason,omitempty"` // Explains why a file is non compliant for a CIS check.
}

// V1SlsaProvenanceSlsaBuilder represents the V1SlsaProvenanceSlsaBuilder schema from the OpenAPI specification
type V1SlsaProvenanceSlsaBuilder struct {
	Id string `json:"id,omitempty"`
}

// V1BatchCreateOccurrencesResponse represents the V1BatchCreateOccurrencesResponse schema from the OpenAPI specification
type V1BatchCreateOccurrencesResponse struct {
	Occurrences []V1Occurrence `json:"occurrences,omitempty"` // The occurrences that were created.
}

// V1Jwt represents the V1Jwt schema from the OpenAPI specification
type V1Jwt struct {
	Compactjwt string `json:"compactJwt,omitempty"`
}

// V1GerritSourceContext represents the V1GerritSourceContext schema from the OpenAPI specification
type V1GerritSourceContext struct {
	Gerritproject string `json:"gerritProject,omitempty"` // The full project name within the host. Projects may be nested, so "project/subproject" is a valid project name. The "repo name" is the hostURI/project.
	Hosturi string `json:"hostUri,omitempty"` // The URI of a running Gerrit instance.
	Revisionid string `json:"revisionId,omitempty"` // A revision (commit) ID.
	Aliascontext V1AliasContext `json:"aliasContext,omitempty"` // An alias to a repo revision.
}

// V1Recipe represents the V1Recipe schema from the OpenAPI specification
type V1Recipe struct {
	Arguments []ProtobufAny `json:"arguments,omitempty"` // Collection of all external inputs that influenced the build on top of recipe.definedInMaterial and recipe.entryPoint. For example, if the recipe type were "make", then this might be the flags passed to make aside from the target, which is captured in recipe.entryPoint. Since the arguments field can greatly vary in structure, depending on the builder and recipe type, this is of form "Any".
	Definedinmaterial string `json:"definedInMaterial,omitempty"` // Index in materials containing the recipe steps that are not implied by recipe.type. For example, if the recipe type were "make", then this would point to the source containing the Makefile, not the make program itself. Set to -1 if the recipe doesn't come from a material, as zero is default unset value for int64.
	Entrypoint string `json:"entryPoint,omitempty"` // String identifying the entry point into the build. This is often a path to a configuration file and/or a target label within that file. The syntax and meaning are defined by recipe.type. For example, if the recipe type were "make", then this would reference the directory in which to run make as well as which target to use.
	Environment []ProtobufAny `json:"environment,omitempty"` // Any other builder-controlled inputs necessary for correctly evaluating the recipe. Usually only needed for reproducing the build but not evaluated as part of policy. Since the environment field can greatly vary in structure, depending on the builder and recipe type, this is of form "Any".
	TypeField string `json:"type,omitempty"` // URI indicating what type of recipe was performed. It determines the meaning of recipe.entryPoint, recipe.arguments, recipe.environment, and materials.
}

// V1ListNotesResponse represents the V1ListNotesResponse schema from the OpenAPI specification
type V1ListNotesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The next pagination token in the list response. It should be used as `page_token` for the following request. An empty value means no more results.
	Notes []V1Note `json:"notes,omitempty"` // The notes requested.
}

// V1SBOMReferenceOccurrence represents the V1SBOMReferenceOccurrence schema from the OpenAPI specification
type V1SBOMReferenceOccurrence struct {
	Payload V1SbomReferenceIntotoPayload `json:"payload,omitempty"` // The actual payload that contains the SBOM Reference data. The payload follows the intoto statement specification. See https://github.com/in-toto/attestation/blob/main/spec/v1.0/statement.md for more details.
	Payloadtype string `json:"payloadType,omitempty"` // The kind of payload that SbomReferenceIntotoPayload takes. Since it's in the intoto format, this value is expected to be 'application/vnd.in-toto+json'.
	Signatures []V1EnvelopeSignature `json:"signatures,omitempty"` // The signatures over the payload.
}

// V1Digest represents the V1Digest schema from the OpenAPI specification
type V1Digest struct {
	Digestbytes string `json:"digestBytes,omitempty"` // Value of the digest.
	Algo string `json:"algo,omitempty"` // `SHA1`, `SHA512` etc.
}

// V1ProjectRepoId represents the V1ProjectRepoId schema from the OpenAPI specification
type V1ProjectRepoId struct {
	Reponame string `json:"repoName,omitempty"` // The name of the repo. Leave empty for the default repo.
	Projectid string `json:"projectId,omitempty"` // The ID of the project.
}

// WindowsDetailKnowledgeBase represents the WindowsDetailKnowledgeBase schema from the OpenAPI specification
type WindowsDetailKnowledgeBase struct {
	Url string `json:"url,omitempty"` // A link to the KB in the [Windows update catalog] (https://www.catalog.update.microsoft.com/).
	Name string `json:"name,omitempty"` // The KB name (generally of the form KB[0-9]+ (e.g., KB123456)).
}

// V1ListNoteOccurrencesResponse represents the V1ListNoteOccurrencesResponse schema from the OpenAPI specification
type V1ListNoteOccurrencesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to provide to skip to a particular spot in the list.
	Occurrences []V1Occurrence `json:"occurrences,omitempty"` // The occurrences attached to the specified note.
}

// V1BuilderConfig represents the V1BuilderConfig schema from the OpenAPI specification
type V1BuilderConfig struct {
	Id string `json:"id,omitempty"`
}

// V1Occurrence represents the V1Occurrence schema from the OpenAPI specification
type V1Occurrence struct {
	PackageField V1PackageOccurrence `json:"package,omitempty"` // Details on how a particular software package was installed on a system.
	Sbomreference V1SBOMReferenceOccurrence `json:"sbomReference,omitempty"` // The occurrence representing an SBOM reference as applied to a specific resource. The occurrence follows the DSSE specification. See https://github.com/secure-systems-lab/dsse/blob/master/envelope.md for more details.
	Createtime string `json:"createTime,omitempty"` // Output only. The time this occurrence was created.
	Name string `json:"name,omitempty"` // Output only. The name of the occurrence in the form of `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.
	Deployment V1DeploymentOccurrence `json:"deployment,omitempty"` // The period during which some deployable was active in a runtime.
	Notename string `json:"noteName,omitempty"` // Required. Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
	Remediation string `json:"remediation,omitempty"` // A description of actions that can be taken to remedy the note.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time this occurrence was last updated.
	Vulnerability V1VulnerabilityOccurrence `json:"vulnerability,omitempty"` // An occurrence of a severity vulnerability on a resource.
	Compliance V1ComplianceOccurrence `json:"compliance,omitempty"` // An indication that the compliance checks in the associated ComplianceNote were not satisfied for particular resources or a specified reason.
	Resourceuri string `json:"resourceUri,omitempty"` // Required. Immutable. A URI that represents the resource for which the occurrence applies. For example, `https://gcr.io/project/image@sha256:123abc` for a Docker image.
	Attestation V1AttestationOccurrence `json:"attestation,omitempty"` // Occurrence that represents a single "attestation". The authenticity of an attestation can be verified using the attached signature. If the verifier trusts the public key of the signer, then verifying the signature is sufficient to establish trust. In this circumstance, the authority to which this attestation is attached is primarily useful for lookup (how to find this attestation if you already know the authority and artifact to be verified) and intent (for which authority this attestation was intended to sign.
	Build V1BuildOccurrence `json:"build,omitempty"` // Details of a build occurrence.
	Discovery V1DiscoveryOccurrence `json:"discovery,omitempty"` // Provides information about the analysis status of a discovered resource.
	Envelope V1Envelope `json:"envelope,omitempty"` // MUST match https://github.com/secure-systems-lab/dsse/blob/master/envelope.proto. An authenticated message of arbitrary type.
	Secret V1SecretOccurrence `json:"secret,omitempty"` // The occurrence provides details of a secret.
	Dsseattestation V1DSSEAttestationOccurrence `json:"dsseAttestation,omitempty"` // Deprecated. Prefer to use a regular Occurrence, and populate the Envelope at the top level of the Occurrence.
	Image V1ImageOccurrence `json:"image,omitempty"` // Details of the derived image portion of the DockerImage relationship. This image would be produced from a Dockerfile with FROM <DockerImage.Basis in attached Note>.
	Upgrade V1UpgradeOccurrence `json:"upgrade,omitempty"` // An Upgrade Occurrence represents that a specific resource_url could install a specific upgrade. This presence is supplied via local sources (i.e. it is present in the mirror and the running system has noticed its availability). For Windows, both distribution and windows_update contain information for the Windows update.
	Kind string `json:"kind,omitempty"` // Kind represents the kinds of notes supported. - NOTE_KIND_UNSPECIFIED: Default value. This value is unused. - VULNERABILITY: The note and occurrence represent a package vulnerability. - BUILD: The note and occurrence assert build provenance. - IMAGE: This represents an image basis relationship. - PACKAGE: This represents a package installed via a package manager. - DEPLOYMENT: The note and occurrence track deployment events. - DISCOVERY: The note and occurrence track the initial discovery status of a resource. - ATTESTATION: This represents a logical "role" that can attest to artifacts. - UPGRADE: This represents an available package upgrade. - COMPLIANCE: This represents a Compliance Note - DSSE_ATTESTATION: This represents a DSSE attestation Note - VULNERABILITY_ASSESSMENT: This represents a Vulnerability Assessment. - SBOM_REFERENCE: This represents an SBOM Reference. - SECRET: This represents a secret.
}

// SlsaProvenanceZeroTwoSlsaInvocation represents the SlsaProvenanceZeroTwoSlsaInvocation schema from the OpenAPI specification
type SlsaProvenanceZeroTwoSlsaInvocation struct {
	Configsource SlsaProvenanceZeroTwoSlsaConfigSource `json:"configSource,omitempty"` // Describes where the config file that kicked off the build came from. This is effectively a pointer to the source where buildConfig came from.
	Environment map[string]interface{} `json:"environment,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// VulnerabilityAssessmentNotePublisher represents the VulnerabilityAssessmentNotePublisher schema from the OpenAPI specification
type VulnerabilityAssessmentNotePublisher struct {
	Issuingauthority string `json:"issuingAuthority,omitempty"` // Provides information about the authority of the issuing party to release the document, in particular, the party's constituency and responsibilities or other obligations.
	Name string `json:"name,omitempty"` // Name of the publisher. Examples: 'Google', 'Google Cloud Platform'.
	Publishernamespace string `json:"publisherNamespace,omitempty"`
}

// V1PackageNote represents the V1PackageNote schema from the OpenAPI specification
type V1PackageNote struct {
	Digest []V1Digest `json:"digest,omitempty"` // Hash value, typically a file digest, that allows unique identification a specific package.
	Distribution []V1Distribution `json:"distribution,omitempty"` // Deprecated. The various channels by which a package is distributed.
	Cpeuri string `json:"cpeUri,omitempty"` // The cpe_uri in [CPE format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package. The cpe_uri will be blank for language packages.
	Architecture string `json:"architecture,omitempty"` // Instruction set architectures supported by various package managers. - ARCHITECTURE_UNSPECIFIED: Unknown architecture. - X86: X86 architecture. - X64: X64 architecture.
	License V1License `json:"license,omitempty"` // License information.
	Packagetype string `json:"packageType,omitempty"` // The type of package; whether native or non native (e.g., ruby gems, node.js packages, etc.).
	Url string `json:"url,omitempty"` // The homepage for this package.
	Description string `json:"description,omitempty"` // The description of this package.
	Maintainer string `json:"maintainer,omitempty"` // A freeform text denoting the maintainer of this package.
	Name string `json:"name"` // The name of the package.
	Version V1Version `json:"version,omitempty"` // Version contains structured information about the version of a package.
}

// V1Signature represents the V1Signature schema from the OpenAPI specification
type V1Signature struct {
	Signature string `json:"signature,omitempty"` // The content of the signature, an opaque bytestring. The payload that this signature verifies MUST be unambiguously provided with the Signature during verification. A wrapper message might provide the payload explicitly. Alternatively, a message might have a canonical serialization that can always be unambiguously computed to derive the payload.
	Publickeyid string `json:"publicKeyId,omitempty"` // The identifier for the public key that verifies this signature. * The `public_key_id` is required. * The `public_key_id` SHOULD be an RFC3986 conformant URI. * When possible, the `public_key_id` SHOULD be an immutable reference, such as a cryptographic digest. Examples of valid `public_key_id`s: OpenPGP V4 public key fingerprint: * "openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA" See https://www.iana.org/assignments/uri-schemes/prov/openpgp4fpr for more details on this scheme. RFC6920 digest-named SubjectPublicKeyInfo (digest of the DER serialization): * "ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU" * "nih:///sha-256;703f68f42aba2c6de30f488a5ea122fef76324679c9bf89791ba95a1271589a5"
}

// V1InTotoProvenance represents the V1InTotoProvenance schema from the OpenAPI specification
type V1InTotoProvenance struct {
	Recipe V1Recipe `json:"recipe,omitempty"` // Steps taken to build the artifact. For a TaskRun, typically each container corresponds to one step in the recipe.
	Builderconfig V1BuilderConfig `json:"builderConfig,omitempty"`
	Materials []string `json:"materials,omitempty"` // The collection of artifacts that influenced the build including sources, dependencies, build tools, base images, and so on. This is considered to be incomplete unless metadata.completeness.materials is true. Unset or null is equivalent to empty.
	Metadata V1Metadata `json:"metadata,omitempty"` // Other properties of the build.
}

// VulnerabilityNoteWindowsDetail represents the VulnerabilityNoteWindowsDetail schema from the OpenAPI specification
type VulnerabilityNoteWindowsDetail struct {
	Cpeuri string `json:"cpeUri,omitempty"` // Required. The [CPE URI](https://cpe.mitre.org/specification/) this vulnerability affects.
	Description string `json:"description,omitempty"` // The description of this vulnerability.
	Fixingkbs []WindowsDetailKnowledgeBase `json:"fixingKbs,omitempty"` // Required. The names of the KBs which have hotfixes to mitigate this vulnerability. Note that there may be multiple hotfixes (and thus multiple KBs) that mitigate a given vulnerability. Currently any listed KBs presence is considered a fix.
	Name string `json:"name,omitempty"` // Required. The name of this vulnerability.
}

// V1CVSS represents the V1CVSS schema from the OpenAPI specification
type V1CVSS struct {
	Exploitabilityscore float32 `json:"exploitabilityScore,omitempty"`
	Impactscore float32 `json:"impactScore,omitempty"`
	Integrityimpact string `json:"integrityImpact,omitempty"`
	Scope string `json:"scope,omitempty"`
	Authentication string `json:"authentication,omitempty"`
	Confidentialityimpact string `json:"confidentialityImpact,omitempty"`
	Attackvector string `json:"attackVector,omitempty"`
	Userinteraction string `json:"userInteraction,omitempty"`
	Basescore float32 `json:"baseScore,omitempty"` // The base score is a function of the base metric scores.
	Privilegesrequired string `json:"privilegesRequired,omitempty"`
	Attackcomplexity string `json:"attackComplexity,omitempty"`
	Availabilityimpact string `json:"availabilityImpact,omitempty"`
}

// V1SlsaProvenanceSlsaMetadata represents the V1SlsaProvenanceSlsaMetadata schema from the OpenAPI specification
type V1SlsaProvenanceSlsaMetadata struct {
	Buildstartedon string `json:"buildStartedOn,omitempty"` // The timestamp of when the build started.
	Completeness V1SlsaProvenanceSlsaCompleteness `json:"completeness,omitempty"` // Indicates that the builder claims certain fields in this message to be complete.
	Reproducible bool `json:"reproducible,omitempty"` // If true, the builder claims that running the recipe on materials will produce bit-for-bit identical output.
	Buildfinishedon string `json:"buildFinishedOn,omitempty"` // The timestamp of when the build completed.
	Buildinvocationid string `json:"buildInvocationId,omitempty"` // Identifies the particular build invocation, which can be useful for finding associated logs or other ad-hoc analysis. The value SHOULD be globally unique, per in-toto Provenance spec.
}

// V1Artifact represents the V1Artifact schema from the OpenAPI specification
type V1Artifact struct {
	Id string `json:"id,omitempty"` // Artifact ID, if any; for container images, this will be a URL by digest like `gcr.io/projectID/imagename@sha256:123456`.
	Names []string `json:"names,omitempty"` // Related artifact names. This may be the path to a binary or jar file, or in the case of a container build, the name used to push the container image to Google Container Registry, as presented to `docker push`. Note that a single Artifact ID can have multiple names, for example if two tags are applied to one image.
	Checksum string `json:"checksum,omitempty"` // Hash or checksum value of a binary, or Docker Registry 2.0 digest of a container.
}

// V1DeploymentOccurrence represents the V1DeploymentOccurrence schema from the OpenAPI specification
type V1DeploymentOccurrence struct {
	Config string `json:"config,omitempty"` // Configuration used to create this deployment.
	Deploytime string `json:"deployTime,omitempty"` // Required. Beginning of the lifetime of this deployment.
	Platform string `json:"platform,omitempty"` // Types of platforms. - PLATFORM_UNSPECIFIED: Unknown. - GKE: Google Container Engine. - FLEX: Google App Engine: Flexible Environment. - CUSTOM: Custom user-defined platform.
	Resourceuri []string `json:"resourceUri,omitempty"` // Output only. Resource URI for the artifact being deployed taken from the deployable field with the same name.
	Undeploytime string `json:"undeployTime,omitempty"` // End of the lifetime of this deployment.
	Useremail string `json:"userEmail,omitempty"` // Identity of the user that triggered this deployment.
	Address string `json:"address,omitempty"` // Address of the runtime element hosting this deployment.
}

// V1SlsaProvenance represents the V1SlsaProvenance schema from the OpenAPI specification
type V1SlsaProvenance struct {
	Materials []SlsaProvenanceMaterial `json:"materials,omitempty"` // The collection of artifacts that influenced the build including sources, dependencies, build tools, base images, and so on. This is considered to be incomplete unless metadata.completeness.materials is true. Unset or null is equivalent to empty.
	Metadata V1SlsaProvenanceSlsaMetadata `json:"metadata,omitempty"` // Other properties of the build.
	Recipe SlsaProvenanceSlsaRecipe `json:"recipe,omitempty"` // Steps taken to build the artifact. For a TaskRun, typically each container corresponds to one step in the recipe.
	Builder V1SlsaProvenanceSlsaBuilder `json:"builder,omitempty"`
}

// V1SlsaProvenanceZeroTwoSlsaCompleteness represents the V1SlsaProvenanceZeroTwoSlsaCompleteness schema from the OpenAPI specification
type V1SlsaProvenanceZeroTwoSlsaCompleteness struct {
	Parameters bool `json:"parameters,omitempty"`
	Environment bool `json:"environment,omitempty"`
	Materials bool `json:"materials,omitempty"`
}

// V1SlsaProvenanceZeroTwoSlsaBuilder represents the V1SlsaProvenanceZeroTwoSlsaBuilder schema from the OpenAPI specification
type V1SlsaProvenanceZeroTwoSlsaBuilder struct {
	Id string `json:"id,omitempty"`
}

// V1ComplianceNote represents the V1ComplianceNote schema from the OpenAPI specification
type V1ComplianceNote struct {
	Remediation string `json:"remediation,omitempty"` // A description of remediation steps if the compliance check fails.
	Scaninstructions string `json:"scanInstructions,omitempty"` // Serialized scan instructions with a predefined format.
	Title string `json:"title,omitempty"` // The title that identifies this compliance check.
	Version []V1ComplianceVersion `json:"version,omitempty"` // The OS and config versions the benchmark applies to.
	Cisbenchmark ComplianceNoteCisBenchmark `json:"cisBenchmark,omitempty"` // A compliance check that is a CIS benchmark.
	Description string `json:"description,omitempty"` // A description about this compliance check.
	Impact string `json:"impact,omitempty"`
	Rationale string `json:"rationale,omitempty"` // A rationale for the existence of this compliance check.
}

// GooglerpcStatus represents the GooglerpcStatus schema from the OpenAPI specification
type GooglerpcStatus struct {
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Details []ProtobufAny `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
}

// V1UpgradeNote represents the V1UpgradeNote schema from the OpenAPI specification
type V1UpgradeNote struct {
	Distributions []V1UpgradeDistribution `json:"distributions,omitempty"` // Metadata about the upgrade for each specific operating system.
	PackageField string `json:"package,omitempty"` // Required for non-Windows OS. The package this Upgrade is for.
	Version V1Version `json:"version,omitempty"` // Version contains structured information about the version of a package.
	Windowsupdate V1WindowsUpdate `json:"windowsUpdate,omitempty"` // Windows Update represents the metadata about the update for the Windows operating system. The fields in this message come from the Windows Update API documented at https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdate.
}

// V1InTotoSlsaProvenanceV1 represents the V1InTotoSlsaProvenanceV1 schema from the OpenAPI specification
type V1InTotoSlsaProvenanceV1 struct {
	Predicate InTotoSlsaProvenanceV1SlsaProvenanceV1 `json:"predicate,omitempty"` // Keep in sync with schema at https://github.com/slsa-framework/slsa/blob/main/docs/provenance/schema/v1/provenance.proto Builder renamed to ProvenanceBuilder because of Java conflicts.
	Predicatetype string `json:"predicateType,omitempty"`
	Subject []V1Subject `json:"subject,omitempty"`
	TypeField string `json:"_type,omitempty"`
}

// SlsaProvenanceZeroTwoSlsaMaterial represents the SlsaProvenanceZeroTwoSlsaMaterial schema from the OpenAPI specification
type SlsaProvenanceZeroTwoSlsaMaterial struct {
	Digest map[string]interface{} `json:"digest,omitempty"`
	Uri string `json:"uri,omitempty"`
}

// V1SlsaProvenanceSlsaCompleteness represents the V1SlsaProvenanceSlsaCompleteness schema from the OpenAPI specification
type V1SlsaProvenanceSlsaCompleteness struct {
	Arguments bool `json:"arguments,omitempty"` // If true, the builder claims that recipe.arguments is complete, meaning that all external inputs are properly captured in the recipe.
	Environment bool `json:"environment,omitempty"` // If true, the builder claims that recipe.environment is claimed to be complete.
	Materials bool `json:"materials,omitempty"` // If true, the builder claims that materials are complete, usually through some controls to prevent network access. Sometimes called "hermetic".
}

// V1EnvelopeSignature represents the V1EnvelopeSignature schema from the OpenAPI specification
type V1EnvelopeSignature struct {
	Keyid string `json:"keyid,omitempty"`
	Sig string `json:"sig,omitempty"`
}

// DSSEAttestationNoteDSSEHint represents the DSSEAttestationNoteDSSEHint schema from the OpenAPI specification
type DSSEAttestationNoteDSSEHint struct {
	Humanreadablename string `json:"humanReadableName,omitempty"` // Required. The human readable name of this attestation authority, for example "cloudbuild-prod".
}

// V1FileLocation represents the V1FileLocation schema from the OpenAPI specification
type V1FileLocation struct {
	Filepath string `json:"filePath,omitempty"` // For jars that are contained inside .war files, this filepath can indicate the path to war file combined with the path to jar file.
	Layerdetails V1LayerDetails `json:"layerDetails,omitempty"` // Details about the layer a package was found in.
}

// VulnerabilityAssessmentNoteProduct represents the VulnerabilityAssessmentNoteProduct schema from the OpenAPI specification
type VulnerabilityAssessmentNoteProduct struct {
	Genericuri string `json:"genericUri,omitempty"` // Contains a URI which is vendor-specific. Example: The artifact repository URL of an image.
	Id string `json:"id,omitempty"` // Token that identifies a product so that it can be referred to from other parts in the document. There is no predefined format as long as it uniquely identifies a group in the context of the current document.
	Name string `json:"name,omitempty"` // Name of the product.
}

// V1Note represents the V1Note schema from the OpenAPI specification
type V1Note struct {
	Vulnerability V1VulnerabilityNote `json:"vulnerability,omitempty"` // A security vulnerability that can be found in resources.
	Kind string `json:"kind,omitempty"` // Kind represents the kinds of notes supported. - NOTE_KIND_UNSPECIFIED: Default value. This value is unused. - VULNERABILITY: The note and occurrence represent a package vulnerability. - BUILD: The note and occurrence assert build provenance. - IMAGE: This represents an image basis relationship. - PACKAGE: This represents a package installed via a package manager. - DEPLOYMENT: The note and occurrence track deployment events. - DISCOVERY: The note and occurrence track the initial discovery status of a resource. - ATTESTATION: This represents a logical "role" that can attest to artifacts. - UPGRADE: This represents an available package upgrade. - COMPLIANCE: This represents a Compliance Note - DSSE_ATTESTATION: This represents a DSSE attestation Note - VULNERABILITY_ASSESSMENT: This represents a Vulnerability Assessment. - SBOM_REFERENCE: This represents an SBOM Reference. - SECRET: This represents a secret.
	Createtime string `json:"createTime,omitempty"` // Output only. The time this note was created. This field can be used as a filter in list requests.
	Sbomreference V1SBOMReferenceNote `json:"sbomReference,omitempty"` // The note representing an SBOM reference.
	Vulnerabilityassessment V1VulnerabilityAssessmentNote `json:"vulnerabilityAssessment,omitempty"` // A single VulnerabilityAssessmentNote represents one particular product's vulnerability assessment for one CVE.
	Build V1BuildNote `json:"build,omitempty"` // Note holding the version of the provider's builder and the signature of the provenance message in the build details occurrence.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time this note was last updated. This field can be used as a filter in list requests.
	Relatednotenames []string `json:"relatedNoteNames,omitempty"` // Other notes related to this note.
	Longdescription string `json:"longDescription,omitempty"` // A detailed description of this note.
	Shortdescription string `json:"shortDescription,omitempty"` // A one sentence description of this note.
	Upgrade V1UpgradeNote `json:"upgrade,omitempty"` // An Upgrade Note represents a potential upgrade of a package to a given version. For each package version combination (i.e. bash 4.0, bash 4.1, bash 4.1.2), there will be an Upgrade Note. For Windows, windows_update field represents the information related to the update.
	Image V1ImageNote `json:"image,omitempty"` // Basis describes the base image portion (Note) of the DockerImage relationship. Linked occurrences are derived from this or an equivalent image via: FROM <Basis.resource_url> Or an equivalent reference, e.g., a tag of the resource_url.
	Name string `json:"name,omitempty"` // Output only. The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
	Deployment V1DeploymentNote `json:"deployment,omitempty"` // An artifact that can be deployed in some runtime.
	Attestation V1AttestationNote `json:"attestation,omitempty"` // Note kind that represents a logical attestation "role" or "authority". For example, an organization might have one `Authority` for "QA" and one for "build". This note is intended to act strictly as a grouping mechanism for the attached occurrences (Attestations). This grouping mechanism also provides a security boundary, since IAM ACLs gate the ability for a principle to attach an occurrence to a given note. It also provides a single point of lookup to find all attached attestation occurrences, even if they don't all live in the same project.
	Discovery V1DiscoveryNote `json:"discovery,omitempty"` // A note that indicates a type of analysis a provider would perform. This note exists in a provider's project. A `Discovery` occurrence is created in a consumer's project at the start of analysis.
	Expirationtime string `json:"expirationTime,omitempty"` // Time of expiration for this note. Empty if note does not expire.
	Relatedurl []V1RelatedUrl `json:"relatedUrl,omitempty"` // URLs associated with this note.
	Compliance V1ComplianceNote `json:"compliance,omitempty"`
	PackageField V1PackageNote `json:"package,omitempty"` // PackageNote represents a particular package version.
	Secret V1SecretNote `json:"secret,omitempty"` // The note representing a secret.
	Dsseattestation V1DSSEAttestationNote `json:"dsseAttestation,omitempty"`
}

// VulnerabilityOccurrencePackageIssue represents the VulnerabilityOccurrencePackageIssue schema from the OpenAPI specification
type VulnerabilityOccurrencePackageIssue struct {
	Fixedpackage string `json:"fixedPackage,omitempty"` // The package this vulnerability was fixed in. It is possible for this to be different from the affected_package.
	Fixedversion V1Version `json:"fixedVersion,omitempty"` // Version contains structured information about the version of a package.
	Affectedpackage string `json:"affectedPackage,omitempty"` // Required. The package this vulnerability was found in.
	Affectedversion V1Version `json:"affectedVersion,omitempty"` // Version contains structured information about the version of a package.
	Filelocation []V1FileLocation `json:"fileLocation,omitempty"` // The location at which this package was found.
	Packagetype string `json:"packageType,omitempty"` // The type of package (e.g. OS, MAVEN, GO).
	Affectedcpeuri string `json:"affectedCpeUri,omitempty"` // Required. The [CPE URI](https://cpe.mitre.org/specification/) this vulnerability was found in.
	Effectiveseverity string `json:"effectiveSeverity,omitempty"` // Note provider assigned severity/impact ranking. - SEVERITY_UNSPECIFIED: Unknown. - MINIMAL: Minimal severity. - LOW: Low severity. - MEDIUM: Medium severity. - HIGH: High severity. - CRITICAL: Critical severity.
	Fixavailable bool `json:"fixAvailable,omitempty"` // Output only. Whether a fix is available for this package.
	Fixedcpeuri string `json:"fixedCpeUri,omitempty"` // The [CPE URI](https://cpe.mitre.org/specification/) this vulnerability was fixed in. It is possible for this to be different from the affected_cpe_uri.
}

// V1AliasContext represents the V1AliasContext schema from the OpenAPI specification
type V1AliasContext struct {
	Name string `json:"name,omitempty"` // The alias name.
	Kind string `json:"kind,omitempty"` // The type of an alias. - KIND_UNSPECIFIED: Unknown. - FIXED: Git tag. - MOVABLE: Git branch. - OTHER: Used to specify non-standard aliases. For example, if a Git repo has a ref named "refs/foo/bar".
}

// V1LayerDetails represents the V1LayerDetails schema from the OpenAPI specification
type V1LayerDetails struct {
	Index int `json:"index,omitempty"` // The index of the layer in the container image.
	Baseimages []V1BaseImage `json:"baseImages,omitempty"` // The base images the layer is found within.
	Chainid string `json:"chainId,omitempty"`
	Command string `json:"command,omitempty"` // The layer build command that was used to build the layer. This may not be found in all layers depending on how the container image is built.
	Diffid string `json:"diffId,omitempty"` // The diff ID (typically a sha256 hash) of the layer in the container image.
}

// V1Completeness represents the V1Completeness schema from the OpenAPI specification
type V1Completeness struct {
	Arguments bool `json:"arguments,omitempty"` // If true, the builder claims that recipe.arguments is complete, meaning that all external inputs are properly captured in the recipe.
	Environment bool `json:"environment,omitempty"` // If true, the builder claims that recipe.environment is claimed to be complete.
	Materials bool `json:"materials,omitempty"` // If true, the builder claims that materials are complete, usually through some controls to prevent network access. Sometimes called "hermetic".
}

// V1FileHashes represents the V1FileHashes schema from the OpenAPI specification
type V1FileHashes struct {
	Filehash []V1Hash `json:"fileHash,omitempty"` // Required. Collection of file hashes.
}

// AttestationNoteHint represents the AttestationNoteHint schema from the OpenAPI specification
type AttestationNoteHint struct {
	Humanreadablename string `json:"humanReadableName,omitempty"` // Required. The human readable name of this attestation authority, for example "qa".
}

// DiscoveryOccurrenceAnalysisCompleted represents the DiscoveryOccurrenceAnalysisCompleted schema from the OpenAPI specification
type DiscoveryOccurrenceAnalysisCompleted struct {
	Analysistype []string `json:"analysisType,omitempty"`
}

// V1BatchCreateNotesResponse represents the V1BatchCreateNotesResponse schema from the OpenAPI specification
type V1BatchCreateNotesResponse struct {
	Notes []V1Note `json:"notes,omitempty"` // The notes that were created.
}

// V1Subject represents the V1Subject schema from the OpenAPI specification
type V1Subject struct {
	Digest map[string]interface{} `json:"digest,omitempty"`
	Name string `json:"name,omitempty"`
}

// DiscoveryOccurrenceVulnerabilityAttestation represents the DiscoveryOccurrenceVulnerabilityAttestation schema from the OpenAPI specification
type DiscoveryOccurrenceVulnerabilityAttestation struct {
	ErrorField string `json:"error,omitempty"` // If failure, the error reason for why the attestation generation failed.
	Lastattempttime string `json:"lastAttemptTime,omitempty"` // The last time we attempted to generate an attestation.
	State string `json:"state,omitempty"` // An enum indicating the state of the attestation generation. - VULNERABILITY_ATTESTATION_STATE_UNSPECIFIED: Default unknown state. - SUCCESS: Attestation was successfully generated and stored. - FAILURE: Attestation was unsuccessfully generated and stored.
}

// VulnerabilityOccurrenceVexAssessment represents the VulnerabilityOccurrenceVexAssessment schema from the OpenAPI specification
type VulnerabilityOccurrenceVexAssessment struct {
	Vulnerabilityid string `json:"vulnerabilityId,omitempty"` // The vulnerability identifier for this Assessment. Will hold one of common identifiers e.g. CVE, GHSA etc.
	Cve string `json:"cve,omitempty"` // Holds the MITRE standard Common Vulnerabilities and Exposures (CVE) tracking number for the vulnerability. Deprecated: Use vulnerability_id instead to denote CVEs.
	Impacts []string `json:"impacts,omitempty"` // Contains information about the impact of this vulnerability, this will change with time.
	Justification AssessmentJustification `json:"justification,omitempty"` // Justification provides the justification when the state of the assessment if NOT_AFFECTED.
	Notename string `json:"noteName,omitempty"`
	Relateduris []V1RelatedUrl `json:"relatedUris,omitempty"` // Holds a list of references associated with this vulnerability item and assessment.
	Remediations []AssessmentRemediation `json:"remediations,omitempty"` // Specifies details on how to handle (and presumably, fix) a vulnerability.
	State string `json:"state,omitempty"` // Provides the state of this Vulnerability assessment. - STATE_UNSPECIFIED: No state is specified. - AFFECTED: This product is known to be affected by this vulnerability. - NOT_AFFECTED: This product is known to be not affected by this vulnerability. - FIXED: This product contains a fix for this vulnerability. - UNDER_INVESTIGATION: It is not known yet whether these versions are or are not affected by the vulnerability. However, it is still under investigation.
}

// V1VulnerabilityAssessmentNote represents the V1VulnerabilityAssessmentNote schema from the OpenAPI specification
type V1VulnerabilityAssessmentNote struct {
	Longdescription string `json:"longDescription,omitempty"` // A detailed description of this Vex.
	Product VulnerabilityAssessmentNoteProduct `json:"product,omitempty"`
	Publisher VulnerabilityAssessmentNotePublisher `json:"publisher,omitempty"`
	Shortdescription string `json:"shortDescription,omitempty"` // A one sentence description of this Vex.
	Title string `json:"title,omitempty"`
	Assessment VulnerabilityAssessmentNoteAssessment `json:"assessment,omitempty"` // Assessment provides all information that is related to a single vulnerability for this product.
	Languagecode string `json:"languageCode,omitempty"` // Identifies the language used by this document, corresponding to IETF BCP 47 / RFC 5646.
}

// V1Version represents the V1Version schema from the OpenAPI specification
type V1Version struct {
	Revision string `json:"revision,omitempty"` // The iteration of the package build from the above version.
	Epoch int `json:"epoch,omitempty"` // Used to correct mistakes in the version numbering scheme.
	Fullname string `json:"fullName,omitempty"` // Human readable version string. This string is of the form <epoch>:<name>-<revision> and is only set when kind is NORMAL.
	Inclusive bool `json:"inclusive,omitempty"` // Whether this version is specifying part of an inclusive range. Grafeas does not have the capability to specify version ranges; instead we have fields that specify start version and end versions. At times this is insufficient - we also need to specify whether the version is included in the range or is excluded from the range. This boolean is expected to be set to true when the version is included in a range.
	Kind string `json:"kind,omitempty"` // Whether this is an ordinary package version or a sentinel MIN/MAX version. - VERSION_KIND_UNSPECIFIED: Unknown. - NORMAL: A standard package version. - MINIMUM: A special version representing negative infinity. - MAXIMUM: A special version representing positive infinity.
	Name string `json:"name,omitempty"` // Required only when version kind is NORMAL. The main part of the version name.
}

// V1SecretStatus represents the V1SecretStatus schema from the OpenAPI specification
type V1SecretStatus struct {
	Message string `json:"message,omitempty"` // Optional message about the status code.
	Status string `json:"status,omitempty"` // The status of the secret. - STATUS_UNSPECIFIED: Unspecified - UNKNOWN: The status of the secret is unknown. - VALID: The secret is valid. - INVALID: The secret is invalid.
	Updatetime string `json:"updateTime,omitempty"` // The time the secret status was last updated.
}

// V1DSSEAttestationNote represents the V1DSSEAttestationNote schema from the OpenAPI specification
type V1DSSEAttestationNote struct {
	Hint DSSEAttestationNoteDSSEHint `json:"hint,omitempty"` // This submessage provides human-readable hints about the purpose of the authority. Because the name of a note acts as its resource reference, it is important to disambiguate the canonical name of the Note (which might be a UUID for security purposes) from "readable" names more suitable for debug output. Note that these hints should not be used to look up authorities in security sensitive contexts, such as when looking up attestations to verify.
}

// V1PackageOccurrence represents the V1PackageOccurrence schema from the OpenAPI specification
type V1PackageOccurrence struct {
	Name string `json:"name"` // The name of the installed package.
	Packagetype string `json:"packageType,omitempty"` // The type of package; whether native or non native (e.g., ruby gems, node.js packages, etc.).
	Version V1Version `json:"version,omitempty"` // Version contains structured information about the version of a package.
	Architecture string `json:"architecture,omitempty"` // Instruction set architectures supported by various package managers. - ARCHITECTURE_UNSPECIFIED: Unknown architecture. - X86: X86 architecture. - X64: X64 architecture.
	Cpeuri string `json:"cpeUri,omitempty"` // The cpe_uri in [CPE format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package. The cpe_uri will be blank for language packages.
	License V1License `json:"license,omitempty"` // License information.
	Location []Grafeasv1Location `json:"location,omitempty"` // All of the places within the filesystem versions of this package have been found.
}

// V1RelatedUrl represents the V1RelatedUrl schema from the OpenAPI specification
type V1RelatedUrl struct {
	Label string `json:"label,omitempty"` // Label to describe usage of the URL.
	Url string `json:"url,omitempty"` // Specific URL associated with the resource.
}

// V1Distribution represents the V1Distribution schema from the OpenAPI specification
type V1Distribution struct {
	Description string `json:"description,omitempty"` // The distribution channel-specific description of this package.
	Latestversion V1Version `json:"latestVersion,omitempty"` // Version contains structured information about the version of a package.
	Maintainer string `json:"maintainer,omitempty"` // A freeform string denoting the maintainer of this package.
	Url string `json:"url,omitempty"` // The distribution channel-specific homepage for this package.
	Architecture string `json:"architecture,omitempty"` // Instruction set architectures supported by various package managers. - ARCHITECTURE_UNSPECIFIED: Unknown architecture. - X86: X86 architecture. - X64: X64 architecture.
	Cpeuri string `json:"cpeUri"` // The cpe_uri in [CPE format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package.
}

// InTotoSlsaProvenanceV1ProvenanceBuilder represents the InTotoSlsaProvenanceV1ProvenanceBuilder schema from the OpenAPI specification
type InTotoSlsaProvenanceV1ProvenanceBuilder struct {
	Version map[string]interface{} `json:"version,omitempty"`
	Builderdependencies []V1InTotoSlsaProvenanceV1ResourceDescriptor `json:"builderDependencies,omitempty"`
	Id string `json:"id,omitempty"`
}

// V1DiscoveryOccurrence represents the V1DiscoveryOccurrence schema from the OpenAPI specification
type V1DiscoveryOccurrence struct {
	Archivetime string `json:"archiveTime,omitempty"` // The time occurrences related to this discovery occurrence were archived.
	Analysiscompleted DiscoveryOccurrenceAnalysisCompleted `json:"analysisCompleted,omitempty"` // Indicates which analysis completed successfully. Multiple types of analysis can be performed on a single resource.
	Continuousanalysis string `json:"continuousAnalysis,omitempty"` // Whether the resource is continuously analyzed. - CONTINUOUS_ANALYSIS_UNSPECIFIED: Unknown. - ACTIVE: The resource is continuously analyzed. - INACTIVE: The resource is ignored for continuous analysis.
	Vulnerabilityattestation DiscoveryOccurrenceVulnerabilityAttestation `json:"vulnerabilityAttestation,omitempty"` // The status of an vulnerability attestation generation.
	Analysiserror []GooglerpcStatus `json:"analysisError,omitempty"` // Indicates any errors encountered during analysis of a resource. There could be 0 or more of these errors.
	Analysisstatuserror GooglerpcStatus `json:"analysisStatusError,omitempty"` // - Simple to use and understand for most users - Flexible enough to meet unexpected needs # Overview The `Status` message contains three pieces of data: error code, error message, and error details. The error code should be an enum value of [google.rpc.Code][google.rpc.Code], but it may accept additional error codes if needed. The error message should be a developer-facing English message that helps developers *understand* and *resolve* the error. If a localized user-facing error message is needed, put the localized message in the error details or localize it in the client. The optional error details may contain arbitrary information about the error. There is a predefined set of error detail types in the package `google.rpc` that can be used for common error conditions. # Language mapping The `Status` message is the logical representation of the error model, but it is not necessarily the actual wire format. When the `Status` message is exposed in different client libraries and different wire protocols, it can be mapped differently. For example, it will likely be mapped to some exceptions in Java, but more likely mapped to some error codes in C. # Other uses The error model and the `Status` message can be used in a variety of environments, either with or without APIs, to provide a consistent developer experience across different environments. Example uses of this error model include: - Partial errors. If a service needs to return partial errors to the client, it may embed the `Status` in the normal response to indicate the partial errors. - Workflow errors. A typical workflow has multiple steps. Each step may have a `Status` message for error reporting. - Batch operations. If a client uses batch request and batch response, the `Status` message should be used directly inside batch response, one for each error sub-response. - Asynchronous operations. If an API call embeds asynchronous operation results in its response, the status of those operations should be represented directly using the `Status` message. - Logging. If some API errors are stored in logs, the message `Status` could be used directly after any stripping needed for security/privacy reasons.
	Analysisstatus string `json:"analysisStatus,omitempty"` // Analysis status for a resource. Currently for initial analysis only (not updated in continuous analysis). - ANALYSIS_STATUS_UNSPECIFIED: Unknown. - PENDING: Resource is known but no action has been taken yet. - SCANNING: Resource is being analyzed. - FINISHED_SUCCESS: Analysis has finished successfully. - COMPLETE: Analysis has completed. - FINISHED_FAILED: Analysis has finished unsuccessfully, the analysis itself is in a bad state. - FINISHED_UNSUPPORTED: The resource is known not to be supported.
	Cpe string `json:"cpe,omitempty"` // The CPE of the resource being scanned.
	Lastscantime string `json:"lastScanTime,omitempty"` // The last time this resource was scanned.
	Sbomstatus DiscoveryOccurrenceSBOMStatus `json:"sbomStatus,omitempty"` // The status of an SBOM generation.
}

// V1VulnerabilityOccurrence represents the V1VulnerabilityOccurrence schema from the OpenAPI specification
type V1VulnerabilityOccurrence struct {
	Cvssv3 V1CVSS `json:"cvssv3,omitempty"` // Common Vulnerability Scoring System. For details, see https://www.first.org/cvss/specification-document This is a message we will try to use for storing various versions of CVSS rather than making a separate proto for storing a specific version.
	Effectiveseverity string `json:"effectiveSeverity,omitempty"` // Note provider assigned severity/impact ranking. - SEVERITY_UNSPECIFIED: Unknown. - MINIMAL: Minimal severity. - LOW: Low severity. - MEDIUM: Medium severity. - HIGH: High severity. - CRITICAL: Critical severity.
	Packageissue []VulnerabilityOccurrencePackageIssue `json:"packageIssue,omitempty"` // Required. The set of affected locations and their fixes (if available) within the associated resource.
	Severity string `json:"severity,omitempty"` // Note provider assigned severity/impact ranking. - SEVERITY_UNSPECIFIED: Unknown. - MINIMAL: Minimal severity. - LOW: Low severity. - MEDIUM: Medium severity. - HIGH: High severity. - CRITICAL: Critical severity.
	Cvssv2 V1CVSS `json:"cvssV2,omitempty"` // Common Vulnerability Scoring System. For details, see https://www.first.org/cvss/specification-document This is a message we will try to use for storing various versions of CVSS rather than making a separate proto for storing a specific version.
	Fixavailable bool `json:"fixAvailable,omitempty"` // Output only. Whether at least one of the affected packages has a fix available.
	Longdescription string `json:"longDescription,omitempty"` // Output only. A detailed description of this vulnerability.
	Extradetails string `json:"extraDetails,omitempty"` // Occurrence-specific extra details about the vulnerability.
	Relatedurls []V1RelatedUrl `json:"relatedUrls,omitempty"` // Output only. URLs related to this vulnerability.
	Shortdescription string `json:"shortDescription,omitempty"` // Output only. A one sentence description of this vulnerability.
	TypeField string `json:"type,omitempty"` // The type of package; whether native or non native (e.g., ruby gems, node.js packages, etc.).
	Vexassessment VulnerabilityOccurrenceVexAssessment `json:"vexAssessment,omitempty"` // VexAssessment provides all publisher provided Vex information that is related to this vulnerability.
	Cvssscore float32 `json:"cvssScore,omitempty"` // Output only. The CVSS score of this vulnerability. CVSS score is on a scale of 0 - 10 where 0 indicates low severity and 10 indicates high severity.
	Cvssversion string `json:"cvssVersion,omitempty"` // CVSS Version.
}

// InTotoSlsaProvenanceV1BuildDefinition represents the InTotoSlsaProvenanceV1BuildDefinition schema from the OpenAPI specification
type InTotoSlsaProvenanceV1BuildDefinition struct {
	Buildtype string `json:"buildType,omitempty"`
	Externalparameters map[string]interface{} `json:"externalParameters,omitempty"`
	Internalparameters map[string]interface{} `json:"internalParameters,omitempty"`
	Resolveddependencies []V1InTotoSlsaProvenanceV1ResourceDescriptor `json:"resolvedDependencies,omitempty"`
}

// V1Fingerprint represents the V1Fingerprint schema from the OpenAPI specification
type V1Fingerprint struct {
	V2name string `json:"v2Name,omitempty"` // Output only. The name of the image's v2 blobs computed via: [bottom] := v2_blob[bottom] [N] := sha256(v2_blob[N] + " " + v2_name[N+1]) Only the name of the final blob is kept.
	V1name string `json:"v1Name,omitempty"` // Required. The layer ID of the final layer in the Docker image's v1 representation.
	V2blob []string `json:"v2Blob,omitempty"` // Required. The ordered list of v2 blobs that represent a given image.
}

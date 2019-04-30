package idnah

import (
	"internal/x/net/idna"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BidiRuleKey = "BidiRule"

	MapForLookupKey = "MapForLookup"

	NewKey = "New"

	RemoveLeadingDotsKey = "RemoveLeadingDots"

	StrictDomainNameKey = "StrictDomainName"

	ToASCIIKey = "ToASCII"

	ToUnicodeKey = "ToUnicode"

	TransitionalKey = "Transitional"

	ValidateForRegistrationKey = "ValidateForRegistration"

	ValidateLabelsKey = "ValidateLabels"

	VerifyDNSLengthKey = "VerifyDNSLength"
)

func New() hctx.Map {
	return hctx.Map{

		BidiRuleKey: BidiRule,

		MapForLookupKey: MapForLookup,

		NewKey: New,

		RemoveLeadingDotsKey: RemoveLeadingDots,

		StrictDomainNameKey: StrictDomainName,

		ToASCIIKey: ToASCII,

		ToUnicodeKey: ToUnicode,

		TransitionalKey: Transitional,

		ValidateForRegistrationKey: ValidateForRegistration,

		ValidateLabelsKey: ValidateLabels,

		VerifyDNSLengthKey: VerifyDNSLength,
	}
}

// BidiRule enables the Bidi rule as defined in RFC 5893. Any application
// that relies on proper validation of labels should include this rule.
var BidiRule = idna.BidiRule

// MapForLookup sets validation and mapping options such that a given IDN is
// transformed for domain name lookup according to the requirements set out in
// Section 5 of RFC 5891. The mappings follow the recommendations of RFC 5894,
// RFC 5895 and UTS 46. It does not add the Bidi Rule. Use the BidiRule option
// to add this check.
//
// The mappings include normalization and mapping case, width and other
// compatibility mappings.
var MapForLookup = idna.MapForLookup

// New creates a new Profile.
//
// With no options, the returned Profile is the most permissive and equals the
// Punycode Profile. Options can be passed to further restrict the Profile. The
// MapForLookup and ValidateForRegistration options set a collection of options,
// for lookup and registration purposes respectively, which can be tailored by
// adding more fine-grained options, where later options override earlier
// options.
var New = idna.New

// RemoveLeadingDots removes leading label separators. Leading runes that map to
// dots, such as U+3002 IDEOGRAPHIC FULL STOP, are removed as well.
//
// This is the behavior suggested by the UTS #46 and is adopted by some
// browsers.
var RemoveLeadingDots = idna.RemoveLeadingDots

// StrictDomainName limits the set of permissible ASCII characters to those
// allowed in domain names as defined in RFC 1034 (A-Z, a-z, 0-9 and the
// hyphen). This is set by default for MapForLookup and ValidateForRegistration.
//
// This option is useful, for instance, for browsers that allow characters
// outside this range, for example a &#39;_&#39; (U+005F LOW LINE). See
// http://www.rfc-editor.org/std/std3.txt for more details This option
// corresponds to the UseSTD3ASCIIRules option in UTS #46.
var StrictDomainName = idna.StrictDomainName

// ToASCII is a wrapper for Punycode.ToASCII.
var ToASCII = idna.ToASCII

// ToUnicode is a wrapper for Punycode.ToUnicode.
var ToUnicode = idna.ToUnicode

// Transitional sets a Profile to use the Transitional mapping as defined in UTS
// #46. This will cause, for example, &#34;ß&#34; to be mapped to &#34;ss&#34;. Using the
// transitional mapping provides a compromise between IDNA2003 and IDNA2008
// compatibility. It is used by most browsers when resolving domain names. This
// option is only meaningful if combined with MapForLookup.
var Transitional = idna.Transitional

// ValidateForRegistration sets validation options to verify that a given IDN is
// properly formatted for registration as defined by Section 4 of RFC 5891.
var ValidateForRegistration = idna.ValidateForRegistration

// ValidateLabels sets whether to check the mandatory label validation criteria
// as defined in Section 5.4 of RFC 5891. This includes testing for correct use
// of hyphens (&#39;-&#39;), normalization, validity of runes, and the context rules.
var ValidateLabels = idna.ValidateLabels

// VerifyDNSLength sets whether a Profile should fail if any of the IDN parts
// are longer than allowed by the RFC.
var VerifyDNSLength = idna.VerifyDNSLength

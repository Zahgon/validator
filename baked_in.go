package validator

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

type Func func(fl FieldLevel) bool

type FuncCtx func(ctx context.Context, fl FieldLevel) bool

func wrapFunc(fn Func) FuncCtx { _ = "STUB: not implemented"; return *new(FuncCtx) }

var (
	restrictedTags = map[string]struct{}{
		diveTag:           {},
		keysTag:           {},
		endKeysTag:        {},
		structOnlyTag:     {},
		omitzero:          {},
		omitempty:         {},
		omitnil:           {},
		skipValidationTag: {},
		utf8HexComma:      {},
		utf8Pipe:          {},
		noStructLevelTag:  {},
		requiredTag:       {},
		isdefault:         {},
	}

	bakedInAliases = map[string]string{
		"iscolor":         "hexcolor|rgb|rgba|hsl|hsla|cmyk",
		"country_code":    "iso3166_1_alpha2|iso3166_1_alpha3|iso3166_1_alpha_numeric",
		"eu_country_code": "iso3166_1_alpha2_eu|iso3166_1_alpha3_eu|iso3166_1_alpha_numeric_eu",
	}

	bakedInValidators = map[string]Func{
		"required":                      hasValue,
		"required_if":                   requiredIf,
		"required_unless":               requiredUnless,
		"skip_unless":                   skipUnless,
		"required_with":                 requiredWith,
		"required_with_all":             requiredWithAll,
		"required_without":              requiredWithout,
		"required_without_all":          requiredWithoutAll,
		"excluded_if":                   excludedIf,
		"excluded_unless":               excludedUnless,
		"excluded_with":                 excludedWith,
		"excluded_with_all":             excludedWithAll,
		"excluded_without":              excludedWithout,
		"excluded_without_all":          excludedWithoutAll,
		"isdefault":                     isDefault,
		"len":                           hasLengthOf,
		"min":                           hasMinOf,
		"max":                           hasMaxOf,
		"eq":                            isEq,
		"eq_ignore_case":                isEqIgnoreCase,
		"ne":                            isNe,
		"ne_ignore_case":                isNeIgnoreCase,
		"lt":                            isLt,
		"lte":                           isLte,
		"gt":                            isGt,
		"gte":                           isGte,
		"eqfield":                       isEqField,
		"eqcsfield":                     isEqCrossStructField,
		"necsfield":                     isNeCrossStructField,
		"gtcsfield":                     isGtCrossStructField,
		"gtecsfield":                    isGteCrossStructField,
		"ltcsfield":                     isLtCrossStructField,
		"ltecsfield":                    isLteCrossStructField,
		"nefield":                       isNeField,
		"gtefield":                      isGteField,
		"gtfield":                       isGtField,
		"ltefield":                      isLteField,
		"ltfield":                       isLtField,
		"fieldcontains":                 fieldContains,
		"fieldexcludes":                 fieldExcludes,
		"alpha":                         isAlpha,
		"alphaspace":                    isAlphaSpace,
		"alphanum":                      isAlphanum,
		"alphanumspace":                 isAlphaNumericSpace,
		"alphaunicode":                  isAlphaUnicode,
		"alphanumunicode":               isAlphanumUnicode,
		"boolean":                       isBoolean,
		"numeric":                       isNumeric,
		"number":                        isNumber,
		"hexadecimal":                   isHexadecimal,
		"hexcolor":                      isHEXColor,
		"rgb":                           isRGB,
		"rgba":                          isRGBA,
		"hsl":                           isHSL,
		"hsla":                          isHSLA,
		"cmyk":                          isCMYK,
		"e164":                          isE164,
		"email":                         isEmail,
		"url":                           isURL,
		"http_url":                      isHttpURL,
		"https_url":                     isHttpsURL,
		"uri":                           isURI,
		"origin":                        isOrigin,
		"urn_rfc8141":                   isUrnRFC8141,
		"urn_rfc2141":                   isUrnRFC2141,
		"file":                          isFile,
		"filepath":                      isFilePath,
		"base32":                        isBase32,
		"base64":                        isBase64,
		"base64url":                     isBase64URL,
		"base64rawurl":                  isBase64RawURL,
		"contains":                      contains,
		"containsany":                   containsAny,
		"containsrune":                  containsRune,
		"excludes":                      excludes,
		"excludesall":                   excludesAll,
		"excludesrune":                  excludesRune,
		"startswith":                    startsWith,
		"endswith":                      endsWith,
		"startsnotwith":                 startsNotWith,
		"endsnotwith":                   endsNotWith,
		"image":                         isImage,
		"mimetype":                      isMIMEType,
		"isbn":                          isISBN,
		"isbn10":                        isISBN10,
		"isbn13":                        isISBN13,
		"issn":                          isISSN,
		"eth_addr":                      isEthereumAddress,
		"eth_addr_checksum":             isEthereumAddressChecksum,
		"btc_addr":                      isBitcoinAddress,
		"btc_addr_bech32":               isBitcoinBech32Address,
		"uuid":                          isUUID,
		"uuid3":                         isUUID3,
		"uuid4":                         isUUID4,
		"uuid5":                         isUUID5,
		"uuid_rfc4122":                  isUUIDRFC4122,
		"uuid3_rfc4122":                 isUUID3RFC4122,
		"uuid4_rfc4122":                 isUUID4RFC4122,
		"uuid5_rfc4122":                 isUUID5RFC4122,
		"ulid":                          isULID,
		"md4":                           isMD4,
		"md5":                           isMD5,
		"sha256":                        isSHA256,
		"sha384":                        isSHA384,
		"sha512":                        isSHA512,
		"ripemd128":                     isRIPEMD128,
		"ripemd160":                     isRIPEMD160,
		"tiger128":                      isTIGER128,
		"tiger160":                      isTIGER160,
		"tiger192":                      isTIGER192,
		"ascii":                         isASCII,
		"printascii":                    isPrintableASCII,
		"multibyte":                     hasMultiByteCharacter,
		"datauri":                       isDataURI,
		"latitude":                      isLatitude,
		"longitude":                     isLongitude,
		"ssn":                           isSSN,
		"ipv4":                          isIPv4,
		"ipv6":                          isIPv6,
		"ip":                            isIP,
		"cidrv4":                        isCIDRv4,
		"cidrv6":                        isCIDRv6,
		"cidr":                          isCIDR,
		"tcp4_addr":                     isTCP4AddrResolvable,
		"tcp6_addr":                     isTCP6AddrResolvable,
		"tcp_addr":                      isTCPAddrResolvable,
		"udp4_addr":                     isUDP4AddrResolvable,
		"udp6_addr":                     isUDP6AddrResolvable,
		"udp_addr":                      isUDPAddrResolvable,
		"ip4_addr":                      isIP4AddrResolvable,
		"ip6_addr":                      isIP6AddrResolvable,
		"ip_addr":                       isIPAddrResolvable,
		"unix_addr":                     isUnixAddrResolvable,
		"uds_exists":                    isUnixDomainSocketExists,
		"mac":                           isMAC,
		"hostname":                      isHostnameRFC952,
		"hostname_rfc1123":              isHostnameRFC1123,
		"fqdn":                          isFQDN,
		"unique":                        isUnique,
		"oneof":                         isOneOf,
		"oneofci":                       isOneOfCI,
		"noneof":                        isNoneOf,
		"noneofci":                      isNoneOfCI,
		"html":                          isHTML,
		"html_encoded":                  isHTMLEncoded,
		"url_encoded":                   isURLEncoded,
		"dir":                           isDir,
		"dirpath":                       isDirPath,
		"json":                          isJSON,
		"jwt":                           isJWT,
		"hostname_port":                 isHostnamePort,
		"port":                          isPort,
		"lowercase":                     isLowercase,
		"uppercase":                     isUppercase,
		"datetime":                      isDatetime,
		"timezone":                      isTimeZone,
		"iso3166_1_alpha2":              isIso3166Alpha2,
		"iso3166_1_alpha2_eu":           isIso3166Alpha2EU,
		"iso3166_1_alpha3":              isIso3166Alpha3,
		"iso3166_1_alpha3_eu":           isIso3166Alpha3EU,
		"iso3166_1_alpha_numeric":       isIso3166AlphaNumeric,
		"iso3166_1_alpha_numeric_eu":    isIso3166AlphaNumericEU,
		"iso3166_2":                     isIso31662,
		"iso4217":                       isIso4217,
		"iso4217_numeric":               isIso4217Numeric,
		"bcp47_language_tag":            isBCP47LanguageTag,
		"bcp47_strict_language_tag":     isBCP47StrictLanguageTag,
		"postcode_iso3166_alpha2":       isPostcodeByIso3166Alpha2,
		"postcode_iso3166_alpha2_field": isPostcodeByIso3166Alpha2Field,
		"bic_iso_9362_2014":             isIsoBic2014Format,
		"bic":                           isIsoBic2022Format,
		"semver":                        isSemverFormat,
		"dns_rfc1035_label":             isDnsRFC1035LabelFormat,
		"credit_card":                   isCreditCard,
		"cve":                           isCveFormat,
		"luhn_checksum":                 hasLuhnChecksum,
		"mongodb":                       isMongoDBObjectId,
		"mongodb_connection_string":     isMongoDBConnectionString,
		"cron":                          isCron,
		"spicedb":                       isSpiceDB,
		"ein":                           isEIN,
		"validateFn":                    isValidateFn,
	}
)

var (
	oneofValsCache       = map[string][]string{}
	oneofValsCacheRWLock = sync.RWMutex{}

	bcp47LanguageTagRe = regexp.MustCompile(strings.Join([]string{

		`^(`,

		`en-gb-oed|i-ami|i-bnn|i-default|i-enochian|i-hak|i-klingon|i-lux|i-mingo|i-navajo|i-pwn|i-tao|i-tay|i-tsu|`,
		`sgn-be-fr|sgn-be-nl|sgn-ch-de|`,

		`art-lojban|cel-gaulish|no-bok|no-nyn|zh-guoyu|zh-hakka|zh-min|zh-min-nan|zh-xiang|`,

		`x-[a-z0-9]{1,8}`,
		`)$`,

		`|`,

		`^`,
		`((?:[a-z]{2,3}(?:-[a-z]{3}){0,3})|[a-z]{4}|[a-z]{5,8})`,
		`(?:-([a-z]{4}))?`,
		`(?:-([a-z]{2}|[0-9]{3}))?`,
		`(?:-((?:[a-z0-9]{5,8}|[0-9][a-z0-9]{3})(?:-(?:[a-z0-9]{5,8}|[0-9][a-z0-9]{3}))*))?`,
		`(?:-((?:[a-wyz0-9](?:-[a-z0-9]{2,8})+)(?:-(?:[a-wyz0-9](?:-[a-z0-9]{2,8})+))*))?`,
		`(?:-x(?:-[a-z0-9]{1,8})+)?`,
		`$`,
	}, ""))
)

func parseOneOfParam2(s string) []string { _ = "STUB: not implemented"; return nil }

func isURLEncoded(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHTMLEncoded(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHTML(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isOneOf(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isOneOfCI(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isNoneOf(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isNoneOfCI(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUnique(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isMAC(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isCIDRv4(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isCIDRv6(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isCIDR(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIPv4(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIPv6(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIP(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isSSN(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isLongitude(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isLatitude(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isDataURI(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func hasMultiByteCharacter(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isPrintableASCII(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isASCII(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUUID5(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUUID4(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUUID3(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUUID(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUUID5RFC4122(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUUID4RFC4122(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUUID3RFC4122(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUUIDRFC4122(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isULID(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isMD4(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isMD5(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isSHA256(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isSHA384(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isSHA512(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isRIPEMD128(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isRIPEMD160(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isTIGER128(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isTIGER160(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isTIGER192(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isISBN(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isISBN13(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isISBN10(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isISSN(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isEthereumAddress(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isEthereumAddressChecksum(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isBitcoinAddress(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isBitcoinBech32Address(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func excludesRune(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func excludesAll(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func excludes(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func containsRune(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func containsAny(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func contains(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func startsWith(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func endsWith(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func startsNotWith(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func endsNotWith(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func fieldContains(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func fieldExcludes(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isNeField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isNe(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isNeIgnoreCase(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isLteCrossStructField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isLtCrossStructField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isGteCrossStructField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isGtCrossStructField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isNeCrossStructField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isEqCrossStructField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isEqField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isEq(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isEqIgnoreCase(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isPostcodeByIso3166Alpha2(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isPostcodeByIso3166Alpha2Field(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isBase32(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isBase64(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isBase64URL(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isBase64RawURL(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isURI(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isOrigin(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isURL(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHttpURL(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHttpsURL(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUrnRFC8141(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUrnRFC2141(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isFile(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func detectFileMIMEType(field reflect.Value) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func matchesMIMEType(mime, expected string) bool { _ = "STUB: not implemented"; return false }

func isMIMEType(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isImage(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isFilePath(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isE164(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isEmail(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHSLA(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isCMYK(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHSL(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isRGBA(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isRGB(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHEXColor(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHexadecimal(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isNumber(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isNumeric(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isAlphanum(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isAlpha(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isAlphanumUnicode(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isAlphaSpace(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isAlphaNumericSpace(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isAlphaUnicode(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isBoolean(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isDefault(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func hasValue(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func hasNotZeroValue(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func requireCheckFieldKind(fl FieldLevel, param string, defaultNotFoundValue bool) bool {
	_ = "STUB: not implemented"
	return false
}

func requireCheckFieldValue(
	fl FieldLevel, param string, value string, defaultNotFoundValue bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

func requiredIf(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func excludedIf(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func requiredUnless(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func skipUnless(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func excludedUnless(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func excludedWith(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func requiredWith(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func excludedWithAll(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func requiredWithAll(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func excludedWithout(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func requiredWithout(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func excludedWithoutAll(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func requiredWithoutAll(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isGteField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isGtField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isGte(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isGt(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func hasLengthOf(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func hasMinOf(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isLteField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isLtField(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isLte(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isLt(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func hasMaxOf(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isTCP4AddrResolvable(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isTCP6AddrResolvable(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isTCPAddrResolvable(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUDP4AddrResolvable(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUDP6AddrResolvable(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUDPAddrResolvable(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIP4AddrResolvable(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIP6AddrResolvable(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIPAddrResolvable(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUnixAddrResolvable(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUnixDomainSocketExists(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isAbstractSocketExists(sockpath string) bool { _ = "STUB: not implemented"; return false }

func isIP4Addr(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIP6Addr(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHostnameRFC952(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHostnameRFC1123(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isFQDN(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isDir(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isDirPath(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isJSON(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isJWT(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isHostnamePort(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isPort(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isLowercase(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isUppercase(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isDatetime(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isTimeZone(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIso3166Alpha2(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIso3166Alpha2EU(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIso3166Alpha3(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIso3166Alpha3EU(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIso3166AlphaNumeric(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIso3166AlphaNumericEU(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIso31662(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIso4217(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIso4217Numeric(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isBCP47LanguageTag(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isBCP47StrictLanguageTag(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIsoBic2014Format(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isIsoBic2022Format(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isSemverFormat(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isCveFormat(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isDnsRFC1035LabelFormat(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func digitsHaveLuhnChecksum(digits []string) bool { _ = "STUB: not implemented"; return false }

func isMongoDBObjectId(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isMongoDBConnectionString(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isSpiceDB(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isCreditCard(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func hasLuhnChecksum(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isCron(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func isEIN(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

var (
	errMethodNotFound          = errors.New(`method not found`)
	errMethodReturnNoValues    = errors.New(`method return o values (void)`)
	errMethodReturnInvalidType = errors.New(`method should return invalid type`)
)

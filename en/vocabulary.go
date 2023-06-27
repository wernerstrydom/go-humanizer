package en

import (
    "fmt"
    "log"
    "regexp"
    "sort"
    "strings"
    "unicode"
)

type rule struct {
    regex *regexp.Regexp
    repl  string
}

func newRule(regex string, repl string) *rule {
    return &rule{
        regex: regexp.MustCompile(fmt.Sprintf("(?i)%s", regex)),
        repl:  repl,
    }
}

func (r *rule) apply(s string) (string, bool) {
    if r.regex.MatchString(s) {

        result := r.regex.ReplaceAllString(s, r.repl)
        return result, true
    }
    return "", false
}

type vocabulary struct {
    plurals            []*rule
    singulars          []*rule
    irregularPlurals   []*rule
    irregularSingulars []*rule
    uncountables       []string
    letterSRe          *regexp.Regexp
}

func newVocabulary() vocabulary {
    v := vocabulary{
        letterSRe: regexp.MustCompile(`^([sS])[sS]*$`),
    }
    v.addPlural("(ax|test)is$", "${1}es")
    v.addPlural("(octop|alumn|fung|cact|foc|hippopotam|radi|stimul|syllab|nucle)us$", "${1}i")
    v.addPlural("(alias|bias|iris|status|campus|apparatus|virus|walrus|trellis)$", "${1}es")
    v.addPlural("(buffal|tomat|volcan|ech|embarg|her|mosquit|potat|torped|vet)o$", "${1}oes")
    v.addPlural("([dti])um$", "${1}a")
    v.addPlural("sis$", "ses")
    v.addPlural("(?:([^f])fe|([lr])f)$", "${1}${2}ves")
    v.addPlural("(hive)$", "${1}s")
    v.addPlural("([^aeiouy]|qu)y$", "${1}ies")
    v.addPlural("(matr|vert|ind|d)(ix|ex)$", "${1}ices")
    v.addPlural("(x)$", "${1}es")
    v.addPlural("(ch)$", "${1}es")
    v.addPlural("(ss)$", "${1}es")
    v.addPlural("(sh)$", "${1}es")
    v.addPlural("(^[m|l])ouse$", "${1}ice")
    v.addPlural("^(ox)$", "${1}en")
    v.addPlural("(quiz)$", "${1}zes")
    v.addPlural("(buz|blit|walt)z$", "${1}zes")
    v.addPlural("(hoo|lea|loa|thie)f$", "${1}ves")
    v.addPlural("(alumn|alg|larv|vertebr)a$", "${1}ae")
    v.addPlural("(criteri|phenomen)on$", "${1}a")
    v.addPlural("s$", "s")
    v.addPlural("$", "s")

    v.addSingular("s$", "")
    v.addSingular("(n)ews$", "${1}ews")
    v.addSingular("([dti])a$", "${1}um")
    v.addSingular(
        "(analy|ba|diagno|parenthe|progno|synop|the|ellip|empha|neuro|oa|paraly)ses$",
        "${1}sis",
    )
    // v.addSingular("([^f])ves$", "${1}fe")
    v.addSingular("(wi)ves$", "${1}fe")
    v.addSingular("(hive)s$", "${1}")
    v.addSingular("(tive)s$", "${1}")
    v.addSingular("([lr]|hoo|lea|loa|thie)ves$", "${1}f")
    v.addSingular("(^zomb)?([^aeiouy]|qu)ies$", "${2}y")
    v.addSingular("(s)eries$", "${1}eries")
    v.addSingular("(m)ovies$", "${1}ovie")
    v.addSingular("(x|ch|ss|sh)es$", "${1}")
    v.addSingular("(^[m|l])ice$", "${1}ouse")
    // v.addSingular("(?<!^[a-z])(o)es$", "${1}")
    v.addSingular("(shoe)s$", "${1}")
    v.addSingular("(cris|ax|test)es$", "${1}is")
    v.addSingular("(octop|vir|alumn|fung|cact|foc|hippopotam|radi|stimul|syllab|nucle)i$", "${1}us")
    v.addSingular("(alias|bias|iris|status|campus|apparatus|virus|walrus|trellis)es$", "${1}")
    v.addSingular("^(ox)en", "${1}")
    v.addSingular("(matr|d)ices$", "${1}ix")
    v.addSingular("(vert|ind)ices$", "${1}ex")
    v.addSingular("(quiz)zes$", "${1}")
    v.addSingular("(buz|blit|walt)zes$", "${1}z")
    v.addSingular("(alumn|alg|larv|vertebr)ae$", "${1}a")
    v.addSingular("(criteri|phenomen)a$", "${1}on")
    v.addSingular("([t|b|r|c]ook|room|smooth)ies$", "${1}ie")
    v.addSingular("(buffal|tomat|volcan|ech|embarg|her|mosquit|potat|torped|vet)oes$", "${1}o")

    v.addIrregular("person", "people", true)
    v.addIrregular("man", "men", true)
    v.addIrregular("human", "humans", true)
    v.addIrregular("child", "children", true)
    v.addIrregular("sex", "sexes", true)
    v.addIrregular("glove", "gloves", true)
    v.addIrregular("move", "moves", true)
    v.addIrregular("goose", "geese", true)
    v.addIrregular("wave", "waves", true)
    v.addIrregular("foot", "feet", true)
    v.addIrregular("tooth", "teeth", true)
    v.addIrregular("curriculum", "curricula", true)
    v.addIrregular("database", "databases", true)
    v.addIrregular("zombie", "zombies", true)
    v.addIrregular("personnel", "personnel", true)
    v.addIrregular("cache", "caches", true)

    v.addIrregular("ex", "exes", false)
    v.addIrregular("is", "are", false)
    v.addIrregular("that", "those", false)
    v.addIrregular("this", "these", false)
    v.addIrregular("bus", "buses", false)
    v.addIrregular("die", "dice", false)
    v.addIrregular("tie", "ties", false)
    v.addIrregular("movie", "movies", false)
    v.addIrregular("safe", "saves", false)

    v.addUncountable("staff")
    v.addUncountable("training")
    v.addUncountable("equipment")
    v.addUncountable("information")
    v.addUncountable("corn")
    v.addUncountable("milk")
    v.addUncountable("rice")
    v.addUncountable("money")
    v.addUncountable("species")
    v.addUncountable("series")
    v.addUncountable("fish")
    v.addUncountable("sheep")
    v.addUncountable("deer")
    v.addUncountable("aircraft")
    v.addUncountable("oz")
    v.addUncountable("tsp")
    v.addUncountable("tbsp")
    v.addUncountable("ml")
    v.addUncountable("l")
    v.addUncountable("water")
    v.addUncountable("waters")
    v.addUncountable("semen")
    v.addUncountable("sperm")
    v.addUncountable("bison")
    v.addUncountable("grass")
    v.addUncountable("hair")
    v.addUncountable("mud")
    v.addUncountable("elk")
    v.addUncountable("luggage")
    v.addUncountable("moose")
    v.addUncountable("offspring")
    v.addUncountable("salmon")
    v.addUncountable("shrimp")
    v.addUncountable("someone")
    v.addUncountable("swine")
    v.addUncountable("trout")
    v.addUncountable("tuna")
    v.addUncountable("corps")
    v.addUncountable("scissors")
    v.addUncountable("means")
    v.addUncountable("mail")

    // Fix 1132
    v.addUncountable("metadata")

    sort.Slice(
        v.plurals, func(i, j int) bool {
            return compare(v.plurals[i].regex, v.plurals[j].regex) > 0
        },
    )
    sort.Slice(
        v.singulars, func(i, j int) bool {
            return compare(v.singulars[i].regex, v.singulars[j].regex) > 0
        },
    )
    sort.Slice(
        v.irregularPlurals, func(i, j int) bool {
            return compare(v.irregularPlurals[i].regex, v.irregularPlurals[j].regex) > 0
        },
    )
    sort.Slice(
        v.irregularSingulars, func(i, j int) bool {
            return compare(v.irregularSingulars[i].regex, v.irregularSingulars[j].regex) > 0
        },
    )
    return v
}

func compare(re1, re2 *regexp.Regexp) int {
    s1 := re1.String()
    s2 := re2.String()

    // check if the regular express represents an irregular word, if so, it should be sorted first
    if strings.HasPrefix(s1, "(?i)^") && !strings.HasSuffix(s2, "(?i)^") {
        return 1
    }

    if strings.HasPrefix(s2, "(?i)^") && !strings.HasSuffix(s1, "(?i)^") {
        return -1
    }

    if len(s1) < len(s2) {
        return -1
    }
    if len(s1) > len(s2) {
        return 1
    }
    return strings.Compare(s1, s2)
}

func (v *vocabulary) applyRules(skipFirstRule bool, word string, rules []*rule) (string, bool) {
    if len(word) < 1 {
        return word, true
    }

    if v.isUncountable(word) {
        return word, true
    }

    var ok bool
    var result string
    if skipFirstRule {
        rules = rules[1:]
    }

    for _, rule := range rules {
        if result, ok = rule.apply(word); ok && result != "" {
            log.Println("applied rule", rule.regex, "to", word, "=>", result)
            break
        } else {
            log.Println("rule failed", rule.regex, "to", word)
        }
    }

    if result == "" {
        return word, false
    }

    return v.matchUpperCase(word, result), true
}

func (v *vocabulary) matchUpperCase(word, replacement string) string {
    if unicode.IsUpper(rune(word[0])) && unicode.IsLower(rune(replacement[0])) {
        // replace first letter with upper case
        return string(unicode.ToUpper(rune(replacement[0]))) + replacement[1:]
    }
    return replacement
}

func (v *vocabulary) letterS(value string) (string, bool) {
    groups := v.letterSRe.FindStringSubmatch(value)
    if len(groups) > 0 {
        return groups[1], true
    }
    return "", false
}

func (v *vocabulary) addPlural(regex string, repl string) {
    v.plurals = append(v.plurals, newRule(regex, repl))
}

func (v *vocabulary) addSingular(regex string, repl string) {
    v.singulars = append(v.singulars, newRule(regex, repl))
}

func (v *vocabulary) addIrregularPlural(regex string, repl string) {
    v.irregularPlurals = append(v.irregularPlurals, newRule(regex, repl))
}

func (v *vocabulary) addIrregularSingular(regex string, repl string) {
    v.irregularSingulars = append(v.irregularSingulars, newRule(regex, repl))
}

func (v *vocabulary) addUncountable(word string) {
    v.uncountables = append(v.uncountables, strings.ToLower(word))
}

func (v *vocabulary) addIrregular(singular, plural string, matchEnding bool) string {
    if matchEnding {
        v.addIrregularPlural(fmt.Sprintf(`(%c)%s$`, singular[0], singular[1:]), fmt.Sprintf(`${1}%s`, plural[1:]))
        v.addIrregularSingular(fmt.Sprintf(`(%c)%s$`, plural[0], plural[1:]), fmt.Sprintf(`${1}%s`, singular[1:]))
    } else {
        v.addIrregularPlural(`^`+singular+`$`, plural)
        v.addIrregularSingular(`^`+plural+`$`, singular)
    }
    return plural
}

func (v *vocabulary) pluralize(value string) string {
    var result string
    var ok bool

    // is it the letter s? return it
    s, ok := v.letterS(value)
    if ok {
        return s + "s"
    }

    result, ok = v.applyRules(false, value, v.irregularPlurals)
    if ok {
        return result
    }

    result, ok = v.applyRules(false, value, v.plurals)
    if ok {
        return result
    }

    var asSingular, ok1 = v.applyRules(false, value, v.singulars)
    var asSingularAsPlural, _ = v.applyRules(false, asSingular, v.plurals)

    if ok1 && asSingular != value && asSingular+"s" != value && asSingularAsPlural == value && result != value {
        return value
    }

    return result

}

func (v *vocabulary) singularize(value string) string {
    var result string
    var ok bool

    // is it the letter s? return it
    s, ok := v.letterS(value)
    if ok {
        return s
    }

    result, ok = v.applyRules(false, value, v.irregularSingulars)
    if ok {
        return result
    }

    result, ok = v.applyRules(false, value, v.singulars)
    if ok {
        return result
    }

    var asPlural, _ = v.applyRules(false, value, v.plurals)
    var asPluralAsSingular, _ = v.applyRules(false, asPlural, v.singulars)

    if (asPlural != value) && (asPlural != value+"s") && (asPluralAsSingular == value) && (result != value) {
        return value
    }

    if result == "" {
        return value
    }

    return result
}

func (v *vocabulary) isUncountable(s string) bool {
    for _, word := range v.uncountables {
        if word == s {
            return true
        }
    }
    return false
}

var defaultVocabulary vocabulary

func init() {
    defaultVocabulary = newVocabulary()
}

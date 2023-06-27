package en

import (
    "testing"

    "github.com/stretchr/testify/suite"
)

var testCases = map[string]string{
    "search":      "searches",
    "switch":      "switches",
    "fix":         "fixes",
    "box":         "boxes",
    "process":     "processes",
    "address":     "addresses",
    "case":        "cases",
    "stack":       "stacks",
    "wish":        "wishes",
    "fish":        "fish",
    "category":    "categories",
    "query":       "queries",
    "ability":     "abilities",
    "agency":      "agencies",
    "movie":       "movies",
    "archive":     "archives",
    "index":       "indices",
    "wife":        "wives",
    "safe":        "saves",
    "half":        "halves",
    "glove":       "gloves",
    "move":        "moves",
    "salesperson": "salespeople",
    "person":      "people",

    "spokesman":    "spokesmen",
    "man":          "men",
    "woman":        "women",
    "freshman":     "freshmen",
    "chairman":     "chairmen",
    "human":        "humans",
    "personnel":    "personnel",
    "staff":        "staff",
    "training":     "training",
    "basis":        "bases",
    "diagnosis":    "diagnoses",
    "datum":        "data",
    "medium":       "media",
    "analysis":     "analyses",
    "node_child":   "node_children",
    "child":        "children",
    "experience":   "experiences",
    "day":          "days",
    "comment":      "comments",
    "foobar":       "foobars",
    "newsletter":   "newsletters",
    "old_news":     "old_news",
    "news":         "news",
    "series":       "series",
    "species":      "species",
    "quiz":         "quizzes",
    "perspective":  "perspectives",
    "ox":           "oxen",
    "photo":        "photos",
    "buffalo":      "buffaloes",
    "tomato":       "tomatoes",
    "dwarf":        "dwarves",
    "elf":          "elves",
    "information":  "information",
    "equipment":    "equipment",
    "bus":          "buses",
    "status":       "statuses",
    "status_code":  "status_codes",
    "mouse":        "mice",
    "louse":        "lice",
    "house":        "houses",
    "octopus":      "octopi",
    "alias":        "aliases",
    "portfolio":    "portfolios",
    "criterion":    "criteria",
    "vertex":       "vertices",
    "matrix":       "matrices",
    "axis":         "axes",
    "testis":       "testes",
    "crisis":       "crises",
    "corn":         "corn",
    "milk":         "milk",
    "rice":         "rice",
    "shoe":         "shoes",
    "horse":        "horses",
    "prize":        "prizes",
    "edge":         "edges",
    "goose":        "geese",
    "deer":         "deer",
    "sheep":        "sheep",
    "wolf":         "wolves",
    "volcano":      "volcanoes",
    "aircraft":     "aircraft",
    "alumna":       "alumnae",
    "alumnus":      "alumni",
    "fungus":       "fungi",
    "water":        "water",
    "waters":       "waters",
    "semen":        "semen",
    "sperm":        "sperm",
    "wave":         "waves",
    "campus":       "campuses",
    "is":           "are",
    "addendum":     "addenda",
    "alga":         "algae",
    "apparatus":    "apparatuses",
    "appendix":     "appendices",
    "bias":         "biases",
    "bison":        "bison",
    "blitz":        "blitzes",
    "buzz":         "buzzes",
    "cactus":       "cacti",
    "corps":        "corps",
    "curriculum":   "curricula",
    "database":     "databases",
    "die":          "dice",
    "echo":         "echoes",
    "ellipsis":     "ellipses",
    "elk":          "elk",
    "emphasis":     "emphases",
    "embargo":      "embargoes",
    "focus":        "foci",
    "foot":         "feet",
    "fuse":         "fuses",
    "grass":        "grass",
    "hair":         "hair",
    "hero":         "heroes",
    "hippopotamus": "hippopotami",
    "hoof":         "hooves",
    "iris":         "irises",
    "larva":        "larvae",
    "leaf":         "leaves",
    "loaf":         "loaves",
    "luggage":      "luggage",
    "means":        "means",
    "mail":         "mail",
    "millennium":   "millennia",
    "moose":        "moose",
    "mosquito":     "mosquitoes",
    "mud":          "mud",
    "nucleus":      "nuclei",
    "neurosis":     "neuroses",
    "oasis":        "oases",
    "offspring":    "offspring",
    "paralysis":    "paralyses",
    "phenomenon":   "phenomena",
    "potato":       "potatoes",
    "radius":       "radii",
    "salmon":       "salmon",
    "scissors":     "scissors",
    "shrimp":       "shrimp",
    "someone":      "someone",
    "stimulus":     "stimuli",
    "swine":        "swine",
    "syllabus":     "syllabi",
    "that":         "those",
    "thief":        "thieves",
    "this":         "these",
    "tie":          "ties",
    "tooth":        "teeth",
    "torpedo":      "torpedoes",
    "trellis":      "trellises",
    "trout":        "trout",
    "tuna":         "tuna",
    "vertebra":     "vertebrae",
    "veto":         "vetoes",
    "virus":        "viruses",
    "walrus":       "walruses",
    "waltz":        "waltzes",
    "zombie":       "zombies",
    "cookie":       "cookies",
    "bookie":       "bookies",
    "rookie":       "rookies",
    "roomie":       "roomies",
    "smoothie":     "smoothies",
    "cache":        "caches",
    "ex":           "exes",
    "":             "",
    "doe":          "does",
    "hoe":          "hoes",
    "toe":          "toes",
    "woe":          "woes",
    "metadata":     "metadata",
    "a":            "as",
    "A":            "As",
    "s":            "ss",
    "S":            "Ss",
    "z":            "zs",
    "Z":            "Zs",
    "1":            "1s",
}

type InflectionTestSuite struct {
    suite.Suite
}

func (suite *InflectionTestSuite) SetupTest() {
}

func TestInflectionTestSuite(t *testing.T) {
    suite.Run(t, new(InflectionTestSuite))
}

func (suite *InflectionTestSuite) TestPluralize() {

    type args struct {
        value string
    }

    var tests []struct {
        name string
        args args
        want string
    }

    for target, expected := range testCases {
        tests = append(
            tests, struct {
                name string
                args args
                want string
            }{
                name: target,
                args: args{value: target},
                want: expected,
            },
        )

    }

    for _, tt := range tests {
        suite.Run(
            tt.name, func() {
                if got := Plural(tt.args.value); got != tt.want {
                    suite.T().Errorf("Plural(`%v`) = %v, want %v", tt.args.value, got, tt.want)
                } else {
                    suite.T().Logf("Plural(`%v`) = %v, want %v", tt.args.value, got, tt.want)
                }
            },
        )
    }
}

func (suite *InflectionTestSuite) TestSingularize() {

    type args struct {
        value string
    }

    var tests []struct {
        name string
        args args
        want string
    }

    for expected, target := range testCases {
        tests = append(
            tests, struct {
                name string
                args args
                want string
            }{
                name: target,
                args: args{value: target},
                want: expected,
            },
        )
    }

    for _, tt := range tests {
        suite.Run(
            tt.name, func() {
                if got := Singular(tt.args.value); got != tt.want {
                    suite.T().Errorf("Singular(`%v`) = %v, want %v", tt.args.value, got, tt.want)
                } else {
                    suite.T().Logf("Singular(`%v`) = %v, want %v", tt.args.value, got, tt.want)
                }
            },
        )
    }
}

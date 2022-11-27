package sometests

import "time"

// Typescript: TSDeclaration= MySecialArray<T> = T[];

type UserRegisterResponse struct {
	Token string      `json:"token" `
	User  string      `json:"user,omitempty"`
	c     chan string // not exported
}

type HttpSessions struct {
	ID         int64     `json:"id"`
	Key        string    `json:"key"`
	Data       string    `json:"data"`
	CreatedOn  time.Time `json:"created" ts:"type=Date"`
	ModifiedOn time.Time `json:"modified" ts:"type=Date"`
	ExpiresOn  time.Time `json:"expire" ts:"type=Date"`
}

// Typescript: interface
type TestComposition struct {
	UserRegisterResponse `ts:"expand"`
	Info                 string    `json:"info"`
	Count                string    // not exported
	TypeName             string    `json:"typename" ts:"type=MyType"`
	CreatedOn            time.Time `json:"created" ts:"type=Date"`
}

// Typescript: TStype=  MyType = number

// Typescript: type
type TestIntArray []int

// Typescript: type
type TestTypeMap map[string]map[int]string

// Typescript: type=Date
type TestTypeTime time.Time // force Date

// Typescript: interface
type TestStruct struct {
	CreatedOn  time.Time    `json:"created" ts:"type=Date"` // force Date
	TestT      TestIntArray `json:"testIntArray"`
	Session    HttpSessions `json:"Session"`
	ID         int64        `json:"id"`
	Key        []string     `json:"key"`
	Data       *string      `json:"data"`
	DataPTR    *[]string    `json:"dataPointer"`
	ModifiedOn time.Time    `json:"modified" ts:"type=Date"`
	ExpiresOn  time.Time    `json:"-"` // not exported

	MapsArray     []map[string]TestTypeTime    `json:"mapsarray" `
	Maps          map[string]time.Time         `json:"maps" ts:"type=Record<string, Date>"`
	MapsNested    map[string]map[int]string    `json:"MapsNested"`
	MapsNestedPTR map[string]map[int]*[]string `json:"MapsNestedPtr"`
	TestTypeMap   TestTypeMap                  `json:"TestTypeMap"`
	EnumTest      Direction                    `json:"direction"`
	EnumSeason    Season                       `json:"season"`
}

// Typescript: type
type TestTypeStruct TestStruct

// Typescript: const
const Timeout = 1000

// Typescript: const
const (
	Uno   string = "uno"
	Cento int    = 100
)

type Season string

// Typescript: enum=Season
const (
	Summer Season = "summer"
	Autumn        = "autumn"
	Winter        = "winter"
	Spring        = "spring"
)

// Typescript: enum=Test
const (
	A int = iota
	B
	C
	D
)

type Direction int

// Typescript: enum=Direction
const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

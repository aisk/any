package any

// The Type you want to hold any thing.
type Type = interface{}

// The Map which can have any type key as it's key or value.
type Map = map[Type]Type

// The Slice can have any type as it's element.
type Slice = []Type

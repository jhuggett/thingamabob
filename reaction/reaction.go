package reaction

type Resource interface {
	DebugName() string
}

type Reaction interface {
	ReactionType() ReactionType

	SetUnregister(func())
	Unregister()

	SetEnabled(enabled bool)
	IsEnabled() bool

	TryPerform(event *Event, data any) error

	SetDepth(depth []int)
	Depth() []int

	// Essentially the related doodad
	Resource() Resource
	SetResource(Resource)
}

func NewReaction[T any](
	reactionType ReactionType,
	condition func(T) bool,
	callback func(T),
) Reaction {
	return &basicReaction[T]{
		Type:      reactionType,
		Condition: condition,
		Callback:  callback,
		Enabled:   true,
	}
}

package reaction

type Reactions struct {
	Reactions []Reaction

	resource Resource
}

func (r *Reactions) Enable() {
	for _, reaction := range r.Reactions {
		reaction.SetEnabled(true)
	}
}

func (r *Reactions) Disable() {
	for _, reaction := range r.Reactions {
		reaction.SetEnabled(false)
	}
}

func (r *Reactions) Register(gesturer Gesturer, atDepth []int) {
	for _, r := range r.Reactions {
		r.SetUnregister(gesturer.Register(r, atDepth))
	}
}

func (r *Reactions) Add(reactions ...Reaction) {
	for _, reaction := range reactions {
		reaction.SetResource(r.resource)
	}
	r.Reactions = append(r.Reactions, reactions...)
}

func (r *Reactions) Unregister() {
	for _, reaction := range r.Reactions {
		if reaction != nil {
			reaction.Unregister()
		}
	}

	r.Reactions = nil

}

func (r *Reactions) SetResource(res Resource) {
	r.resource = res
	for _, reaction := range r.Reactions {
		reaction.SetResource(res)
	}
}

func (r *Reactions) Resource() Resource {
	return r.resource
}

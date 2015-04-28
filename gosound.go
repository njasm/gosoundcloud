package gosound

type Saver interface {
	Save(s *SoundcloudApi)
}

type Updater interface {
	Update(s *SoundcloudApi)
}

type Deleter interface {
	Delete(s *SoundcloudApi)
}

type Resourcer interface {
	Id()    uint64
	Kind()  string
	IsNew() bool
}

type resource struct {
	id   uint64
	kind string
}

func (r *resource) Id() uint64 {
	return r.id
}

func (r *resource) Kind() string {
	return r.kind
}

func (r *resource) IsNew() bool {
	if r.id > 0 {
		return false
	}
	return true
}

func newResource() *resource {
	return &resource{}
}




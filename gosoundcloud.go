package gosoundcloud

type Saver interface {
	Save(s *SoundcloudApi) error
}

type Updater interface {
	Update(s *SoundcloudApi) error
}

type Deleter interface {
	Delete(s *SoundcloudApi) error
}

type Resourcer interface {
	GetId()     uint64
	GetKind()   string
	IsNew()     bool
}

package art

type Service interface {
	FindAllArt() ([]Art, error)
	FindArtByID(ID int) (Art, error)
	CreateArt(creator Arts) (Art, error)
	UpdateArt(ID int, creator Arts) (Art, error)
	DeleteArt(ID int) (Art, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllArt() ([]Art, error) {
	service, err := s.repository.FindAll()
	return service, err
	// return s.repository.FindAll()
}

func (s *service) FindArtByID(ID int) (Art, error) {
	service, err := s.repository.FindByID(ID)
	return service, err
}

func (s *service) CreateArt(artist Arts) (Art, error) {
	service := Art{
		Name_Art:      artist.Name_Art,
		Name_Artist:   artist.Name_Artist,
		Type:          artist.Type,
		Location:      artist.Location,
		Creation_Date: artist.Creation_Date,
	}
	service, err := s.repository.Create(service)
	return service, err
}

func (s *service) UpdateArt(ID int, artist Arts) (Art, error) {
	services, err := s.repository.FindByID(ID)

	services.Name_Art = artist.Name_Art
	services.Name_Artist = artist.Name_Artist
	services.Type = artist.Type
	services.Location = artist.Location
	services.Creation_Date = artist.Creation_Date

	newservice, err := s.repository.Update(services)
	return newservice, err
}

func (s *service) DeleteArt(ID int) (Art, error) {
	services, err := s.repository.FindByID(ID)
	newservice, err := s.repository.Delete(services)
	return newservice, err
}

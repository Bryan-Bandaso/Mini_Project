package artist

type Service interface {
	FindAll() ([]Artist, error)
	FindByID(ID int) (Artist, error)
	Create(creator Creator) (Artist, error)
	Update(ID int, creator Creator) (Artist, error)
	Delete(ID int) (Artist, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Artist, error) {
	service, err := s.repository.FindAll()
	return service, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Artist, error) {
	service, err := s.repository.FindByID(ID)
	return service, err
}

func (s *service) Create(artist Creator) (Artist, error) {
	service := Artist{
		Name:        artist.Name,
		Nationality: artist.Nationality,
		Description: artist.Description,
		Birth_year:  artist.Birth_year,
		Death_year:  artist.Death_year,
	}
	service, err := s.repository.Create(service)
	return service, err
}

func (s *service) Update(ID int, artist Creator) (Artist, error) {
	services, err := s.repository.FindByID(ID)

	services.Name = artist.Name
	services.Nationality = artist.Nationality
	services.Description = artist.Description
	services.Birth_year = artist.Birth_year
	services.Death_year = artist.Death_year

	newservice, err := s.repository.Update(services)
	return newservice, err
}

func (s *service) Delete(ID int) (Artist, error) {
	services, err := s.repository.FindByID(ID)
	newservice, err := s.repository.Delete(services)
	return newservice, err
}

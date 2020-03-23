package main

// serviceは外部から呼ばれる
// serviceはrepositoryを呼ぶ

type UserService struct {
	repo UserRepository
}

// serviceとrepositoryが依存する
// initRepositoryをテストするとき、All()関数の挙動を変えれない（モックできない）
//func (s *UserService) initRepository() {
//	s.repo = NewUserRepository()
//}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) All() ([]User, error){
	return s.repo.All()
}



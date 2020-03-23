package main

// repositoryはDBとのやり取りを担当

// 抽象化することで内部実装に気を使うことがなくなり、結合度が下がる
// repository以外ではこいつが呼ばれる
// repository以外の関数をテストするとき、このインターフェースをモックできるのでテスタビリティが向上
type UserRepository interface {
	All() ([]User, error)
}

// 抽象化したものの実体
type userRepository struct {}

// 実体をnewして渡す
func NewUserRepository() *userRepository {
	return &userRepository{}
}

// 実体の振る舞い
func (u *userRepository) All() ([]User, error) {
	// DBへ読み書きなど
	return []User{}, nil
}

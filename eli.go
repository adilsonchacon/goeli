package eli

import "github.com/adilsonchacon/eli/entities"

type Letmein interface {
	SignIn(email, password string) (string, int, error)
	SignedIn(sessionToken string) (bool, error)
	CurrentUser(sessionToken string) (*entities.User, int, error)
	SignOut(sessionToken string) (int, error)
	Refresh(sessionToken string) (string, int, error)
	Unlock(unlockToken string) (int, error)
	Confirm(confirmationToken string) (int, error)
	RequestPasswordRecovery(appToken, email string) (int, error)
	RecoverPassword(token, password, passwordConfirmation string) (int, error)
}

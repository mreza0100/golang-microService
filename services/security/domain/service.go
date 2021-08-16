package domain

import (
	"errors"
	"freeSociety/configs"
	"freeSociety/services/security/instances"
	"freeSociety/services/security/models"
	"freeSociety/services/security/repository/postgres"
	"freeSociety/services/security/repository/redis"
	"freeSociety/services/security/utils"
	"freeSociety/utils/security"
	"time"

	generalUtils "freeSociety/utils"

	"github.com/mreza0100/golog"
)

type NewOpts struct {
	Lgr *golog.Core
}

func New(opts *NewOpts) instances.Sevice {
	return &service{
		lgr:          opts.Lgr.With("In domain->"),
		redisRepo:    redis.New(opts.Lgr),
		postgresRepo: postgres.New(opts.Lgr),
	}
}

type service struct {
	redisRepo    *instances.Repo_Redis
	postgresRepo *instances.Repo_Postgres
	lgr          *golog.Core
}

func (s *service) NewUser(userId uint64, device, password string) (token string, err error) {
	debug, success := s.lgr.DebugPKG("NewUser", false)

	{
		hashPass := security.HashSha1(password)
		err = s.postgresRepo.Write.NewUser(userId, hashPass)
		if debug("after s.postgresRepo.NewUser")(err) != nil {
			return "", err
		}
	}
	{
		token = utils.CreateToken()
		expire := generalUtils.ParseDateForDb(time.Now().Add(configs.Token_expire))

		_, err = s.postgresRepo.Write.NewSession(userId, device, token, expire)
		if debug("after s.postgresRepo..NewSession")(err) != nil {
			return "", err
		}
	}
	{
		err = s.redisRepo.Write.NewSession(token, userId)
		if debug("after s.redisDAOS.NewSession")(err) != nil {
			return "", err
		}
	}

	success(token)
	return token, nil
}

func (s *service) Login(userId uint64, device, password string) (string, error) {
	var (
		token string
		err   error
	)

	{
		hashPass, err := s.postgresRepo.Read.GetHashPass(userId)
		if err != nil {
			return "", errors.New("email or password is wrong")
		}
		if !security.HashSha1Compare(hashPass, password) {
			return "", errors.New("email or password is wrong")
		}
	}
	{
		token = utils.CreateToken()
	}
	{
		expire := generalUtils.ParseDateForDb(time.Now().Add(configs.Token_expire))

		_, err = s.postgresRepo.Write.NewSession(userId, device, token, expire)
		if err != nil {
			return "", err
		}
	}
	{
		err = s.redisRepo.Write.NewSession(token, userId)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}

func (s *service) Logout(token string) (err error) {
	{
		err = s.redisRepo.Write.DeleteSession(token)
	}
	{
		err = s.postgresRepo.Write.DeleteSessionByToken(token)
	}
	return err
}

func (s *service) GetUserId(token string) (uint64, error) {
	return s.redisRepo.Read.GetSession(token)
}

func (s *service) PurgeUser(userId uint64) error {
	var (
		tokens []string
		chErr  = make(chan error)
	)

	{
		sessions, err := s.postgresRepo.Write.DeleteUserSessions(userId)
		if err != nil {
			return err
		}
		tokens = make([]string, len(sessions))
		for idx, i := range sessions {
			tokens[idx] = i.Token
		}
	}
	{
		go func(ch chan error) {
			ch <- s.redisRepo.Write.DeleteSession(tokens...)
		}(chErr)
	}
	{
		go func(ch chan error) {
			ch <- s.postgresRepo.Write.DeletePassword(userId)
		}(chErr)
	}

	for i := 0; i < 2; i++ {
		if err := <-chErr; err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetSessions(userId uint64) ([]*models.Session, error) {
	return s.postgresRepo.Read.GetSessions(userId)
}

func (s *service) DeleteSession(sessionId uint64) (err error) {
	var (
		token string
	)

	{
		session, err := s.postgresRepo.Write.DeleteSessionById(sessionId)
		if err != nil {
			return err
		}
		token = session.Token
	}
	{
		err = s.redisRepo.Write.DeleteSession(token)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) ChangePassword(userId uint64, prevPassword, newPassword string) error {
	var (
		tokens []string
	)

	{
		hashPass, err := s.postgresRepo.Read.GetHashPass(userId)
		if err != nil {
			return err
		}

		if !security.HashSha1Compare(hashPass, prevPassword) {
			return errors.New("password is wrong")
		}
	}

	{
		err := s.postgresRepo.Write.ChangeHashPass(userId, security.HashSha1(newPassword))
		if err != nil {
			return err
		}
	}

	{
		sessions, err := s.postgresRepo.Write.DeleteUserSessions(userId)
		if err != nil {
			return err
		}

		tokens = make([]string, len(sessions))
		for idx, i := range sessions {
			tokens[idx] = i.Token
		}
	}

	return s.redisRepo.Write.DeleteSession(tokens...)
}

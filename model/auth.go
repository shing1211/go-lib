package model

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/alexedwards/argon2id"
	"github.com/shing1211/go-lib/config"
)

// Auth model - `auths` table
type Auth struct {
	AuthID    uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string         `json:"Email"`
	Password  string         `json:"Password"`
	Users     User           `gorm:"foreignkey:IDAuth;references:AuthID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// UnmarshalJSON ...
func (v *Auth) UnmarshalJSON(b []byte) error {
	aux := struct {
		AuthID   uint64 `json:"AuthID"`
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}{}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	v.AuthID = aux.AuthID
	v.Email = aux.Email
	if v.Password = HashPass(aux.Password); v.Password == "error" {
		return errors.New("HashPass failed")
	}

	return nil
}

// HashPass ...
func HashPass(pass string) string {
	configure := config.Config()
	params := &argon2id.Params{
		Memory:      configure.Server.ServerHashPass.Memory * 1024, // the amount of memory used by the Argon2 algorithm (in kibibytes)
		Iterations:  configure.Server.ServerHashPass.Iterations,    // the number of iterations (or passes) over the memory
		Parallelism: configure.Server.ServerHashPass.Parallelism,   // the number of threads (or lanes) used by the algorithm
		SaltLength:  configure.Server.ServerHashPass.SaltLength,    // length of the random salt. 16 bytes is recommended for password hashing
		KeyLength:   configure.Server.ServerHashPass.KeyLength,     // length of the generated key (or password hash). 16 bytes or more is recommended
	}
	h, err := argon2id.CreateHash(pass, params)
	if err != nil {
		return "error"
	}
	return h
}

// MarshalJSON ...
func (v Auth) MarshalJSON() ([]byte, error) {
	aux := struct {
		AuthID uint64 `json:"AuthId"`
		Email  string `json:"Email"`
	}{
		AuthID: v.AuthID,
		Email:  v.Email,
	}

	return json.Marshal(aux)
}

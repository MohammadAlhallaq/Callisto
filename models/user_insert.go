package models

type UserInsert struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

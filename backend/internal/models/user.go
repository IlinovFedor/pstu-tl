package models

import (
	"strings"
	"uuid"
)

type UserRole int

const (
	UserRoleTeacher UserRole = iota
	UserRoleEditor
	UserRoleAdmin
)

type User struct {
	name                string
	passwordHash        string
	role                UserRole
	associatedTeacherId *uuid.UUID
}

func NewUser(name string, passwordHash string, role UserRole, associatedTeacherId *uuid.UUID) *User {
	name = strings.ToLower(name)
	return &User{name: strings.TrimSpace(name), passwordHash: passwordHash, role: role, associatedTeacherId: associatedTeacherId}
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) SetPasswordHash(passwordHash string) {
	u.passwordHash = passwordHash
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Role() UserRole {
	return u.role
}

func (u *User) SetRole(role UserRole) {
	u.role = role
}

func (u *User) AssociatedTeacherId() *uuid.UUID {
	return u.associatedTeacherId
}

func (u *User) SetAssociatedTeacherId(associatedTeacherId *uuid.UUID) {
	u.associatedTeacherId = associatedTeacherId
}

func (u UserRole) IsAdminOrEditor() bool {
	if u == UserRoleAdmin || u == UserRoleEditor {
		return true
	}
	return false
}

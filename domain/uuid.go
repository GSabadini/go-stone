package domain

import (
	gouuid "github.com/google/uuid"
)

//NewUUID cria um UUID v4
func NewUUID() string {
	return gouuid.New().String()
}

//IsValidUUID retorna um UUID válido
func IsValidUUID(uuid string) bool {
	_, err := gouuid.Parse(uuid)
	return err == nil
}

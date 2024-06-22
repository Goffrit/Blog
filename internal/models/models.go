// internal/models/models.go

package models

import (
    "database/sql"
)

type User struct {
    UserID      interface{}    `json:"user_id"`
    Username    string         `json:"username"`
    Password    string         `json:"password"`
    Email       string         `json:"email"`
    FullName    sql.NullString `json:"full_name"`
    DateOfBirth sql.NullTime   `json:"date_of_birth"`
    CreatedAt   sql.NullTime   `json:"created_at"`
    UpdatedAt   interface{}    `json:"updated_at"`
}

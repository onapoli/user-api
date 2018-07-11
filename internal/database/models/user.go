package models

import (
  "time"

  "github.com/satori/go.uuid"
)

type User struct {
  Id uuid.UUID `sql:"type:uuid,default:uuid_generate_v4()"`
  CreatedAt time.Time `sql:"default:now()"`
  UpdatedAt time.Time
  FullName string `sql:",notnull"`
  DisplayName string `sql:",unique"` // TODO (notnull) remove as user is private by default
  FirstName string
  LastName string
  Email string `sql:",unique,notnull"`
  Username string `sql:",unique,notnull"` // TODO (notnull?)
  Member bool `sql:",notnull"`
  Avatar []byte
  NewsletterNotification bool

  ResidenceAddressId uuid.UUID  `sql:"type:uuid,notnull"`
  // ResidenceAddress *StreetAddress

  FavoriteTracks []uuid.UUID `sql:",type:uuid[]" pg:",array"`
  FollowedGroups []uuid.UUID `sql:",type:uuid[]" pg:",array"`
  // Playlists

  OwnerOfGroups []UserGroup `pg:"fk:owner_id"`
  // OwnerOfGroups []uuid.UUID `sql:",type:uuid[]" pg:",array"`
  // MemberOfGroups

  // Tags?

  // Shares => Membership API?
}

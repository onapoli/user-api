package model

import "github.com/satori/go.uuid"

type UserGroupPrivacy struct {
  Id uuid.UUID `sql:"type:uuid,default:uuid_generate_v4()"`
  Private bool `sql:",notnull"`
  OwnedTracks bool `sql:",notnull"`
  SupportedArtists bool `sql:",notnull"`
}

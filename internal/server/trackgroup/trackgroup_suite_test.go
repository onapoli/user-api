package trackgroupserver_test

import (
	"testing"
	"time"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/go-pg/pg"
	"github.com/satori/go.uuid"

	"user-api/pkg/config"
	"user-api/pkg/postgres"
	"user-api/internal/model"
	trackgroupserver "user-api/internal/server/trackgroup"
)

var (
	db *pg.DB
	service *trackgroupserver.Server
	newUser *model.User
	newTrack *model.Track
	playlistTrack *model.Track
	newAlbum *model.TrackGroup
	newPlaylist *model.TrackGroup
	newArtistGroupTaxonomy *model.GroupTaxonomy
	newLabelGroupTaxonomy *model.GroupTaxonomy
	newArtistUserGroup *model.UserGroup
	newLabelUserGroup *model.UserGroup
	label *model.UserGroup
	artist *model.UserGroup
	newGenreTag *model.Tag
)

func TestTrack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trackgroup server Suite")
}

var _ = BeforeSuite(func() {
	var err error

	cfgPath, err := filepath.Abs("./../../../conf.local.yaml")
	Expect(err).NotTo(HaveOccurred())

	cfg, err := config.Load(cfgPath)
	Expect(err).NotTo(HaveOccurred())

	db, err = pgsql.New(cfg.DB.Test.PSN, cfg.DB.Test.LogQueries, cfg.DB.Test.TimeoutSeconds)
	Expect(err).NotTo(HaveOccurred())

	service = trackgroupserver.NewServer(db)

	newGenreTag = &model.Tag{Type: "genre", Name: "pop"}
	err = db.Insert(newGenreTag)
	Expect(err).NotTo(HaveOccurred())

	newAddress := &model.StreetAddress{Data: map[string]string{"some": "data"}}
	err = db.Insert(newAddress)
	Expect(err).NotTo(HaveOccurred())

	newArtistGroupTaxonomy = &model.GroupTaxonomy{Type: "artist", Name: "Artist"}
	err = db.Insert(newArtistGroupTaxonomy)
	Expect(err).NotTo(HaveOccurred())

	newLabelGroupTaxonomy = &model.GroupTaxonomy{Type: "label", Name: "Label"}
	err = db.Insert(newLabelGroupTaxonomy)
	Expect(err).NotTo(HaveOccurred())

	newUser = &model.User{Username: "username", FullName: "full name", Email: "email@fake.com"}
	err = db.Insert(newUser)
	Expect(err).NotTo(HaveOccurred())

	// Create new user_groups
	avatar := make([]byte, 5)
	newArtistUserGroup = &model.UserGroup{
		DisplayName: "artist",
		Avatar: avatar,
		OwnerId: newUser.Id,
		TypeId: newArtistGroupTaxonomy.Id,
		AddressId: newAddress.Id,
	}
	err = db.Insert(newArtistUserGroup)
	Expect(err).NotTo(HaveOccurred())

	newLabelUserGroup = &model.UserGroup{
		DisplayName: "label",
		Avatar: avatar,
		OwnerId: newUser.Id,
		TypeId: newLabelGroupTaxonomy.Id,
		AddressId: newAddress.Id,
	}
	err = db.Insert(newLabelUserGroup)
	Expect(err).NotTo(HaveOccurred())

	artist = &model.UserGroup{
		DisplayName: "new artist",
		Avatar: avatar,
		OwnerId: newUser.Id,
		TypeId: newArtistGroupTaxonomy.Id,
		AddressId: newAddress.Id,
	}
	err = db.Insert(artist)
	Expect(err).NotTo(HaveOccurred())

	label = &model.UserGroup{
		DisplayName: "new label",
		Avatar: avatar,
		OwnerId: newUser.Id,
		TypeId: newLabelGroupTaxonomy.Id,
		AddressId: newAddress.Id,
	}
	err = db.Insert(label)
	Expect(err).NotTo(HaveOccurred())

	// Create a new track
	tagIds := []uuid.UUID{newGenreTag.Id}
	newTrack = &model.Track{
		CreatorId: newUser.Id,
		UserGroupId: newArtistUserGroup.Id,
		Artists: []uuid.UUID{newArtistUserGroup.Id},
		Title: "track title",
		Status: "free",
		Tags: tagIds,
	}
	err = db.Insert(newTrack)
	Expect(err).NotTo(HaveOccurred())

	playlistTrack = &model.Track{
		CreatorId: newUser.Id,
		UserGroupId: newArtistUserGroup.Id,
		Artists: []uuid.UUID{newArtistUserGroup.Id},
		Title: "playlist track",
		Status: "free",
	}
	err = db.Insert(playlistTrack)
	Expect(err).NotTo(HaveOccurred())

	favoritingUser := &model.User{Username: "fav", FullName: "fav name", Email: "fav@fake.com", FavoriteTracks: []uuid.UUID{newTrack.Id}}
	err = db.Insert(favoritingUser)
	Expect(err).NotTo(HaveOccurred())

	// Create track groups
	// tracks := map[string]string{
	// 	"1": newTrack.Id.String(),
	// }
	tracks := []uuid.UUID{newTrack.Id}
	newAlbum = &model.TrackGroup{
		CreatorId: newUser.Id,
		UserGroupId: newArtistUserGroup.Id,
		LabelId: newLabelUserGroup.Id,
		Title: "album title",
		ReleaseDate: time.Now(),
		Type: "lp",
		Cover: avatar,
		Tracks: tracks,
		Tags: tagIds,
	}
	err = db.Insert(newAlbum)
	Expect(err).NotTo(HaveOccurred())

	newPlaylist = &model.TrackGroup{
		CreatorId: newUser.Id,
		// UserGroupId: uuid.UUID{},
		// LabelId: uuid.UUID{},
		Title: "playlist title",
		ReleaseDate: time.Now(),
		Type: "playlist",
		Cover: avatar,
		Tracks: tracks,
		Private: true,
	}
	err = db.Insert(newPlaylist)
	Expect(err).NotTo(HaveOccurred())

	newUser.Playlists = []uuid.UUID{newPlaylist.Id}
	_, err = db.Model(newUser).Column("playlists").WherePK().Update()
	Expect(err).NotTo(HaveOccurred())

	newTrack.TrackGroups = []uuid.UUID{newAlbum.Id, newPlaylist.Id}
	newTrack.FavoriteOfUsers = []uuid.UUID{favoritingUser.Id}
	_, err = db.Model(newTrack).Column("track_groups", "favorite_of_users").WherePK().Update()
	Expect(err).NotTo(HaveOccurred())

	newArtistUserGroup.ArtistOfTracks = []uuid.UUID{newTrack.Id, playlistTrack.Id}
	_, err = db.Model(newArtistUserGroup).Column("artist_of_tracks").WherePK().Update()
	Expect(err).NotTo(HaveOccurred())

	// newLabelUserGroup.TrackGroups = []uuid.UUID{newAlbum.Id}
	// _, err = db.Model(newLabelUserGroup).Column("tracks", "track_groups").WherePK().Update()
	// Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	// Delete all tracks
	var tracks []model.Track
	err := db.Model(&tracks).Select()
	Expect(err).NotTo(HaveOccurred())
	if len(tracks) > 0 {
		_, err = db.Model(&tracks).Delete()
		Expect(err).NotTo(HaveOccurred())
	}

	// Delete all track groups
	var trackGroups []model.TrackGroup
	err = db.Model(&trackGroups).Select()
	Expect(err).NotTo(HaveOccurred())
	_, err = db.Model(&trackGroups).Delete()
	Expect(err).NotTo(HaveOccurred())

	// Delete all userGroups
	var userGroups []model.UserGroup
	err = db.Model(&userGroups).Select()
	Expect(err).NotTo(HaveOccurred())
	_, err = db.Model(&userGroups).Delete()
	Expect(err).NotTo(HaveOccurred())

	// Delete all streetAddresses
	var streetAddresses []model.StreetAddress
	err = db.Model(&streetAddresses).Select()
	Expect(err).NotTo(HaveOccurred())
	_, err = db.Model(&streetAddresses).Delete()
	Expect(err).NotTo(HaveOccurred())

	// Delete all users
	var users []model.User
	err = db.Model(&users).Select()
	Expect(err).NotTo(HaveOccurred())
	_, err = db.Model(&users).Delete()
	Expect(err).NotTo(HaveOccurred())

	// Delete all groupTaxonomies
	var groupTaxonomies []model.GroupTaxonomy
	err = db.Model(&groupTaxonomies).Select()
	Expect(err).NotTo(HaveOccurred())
	_, err = db.Model(&groupTaxonomies).Delete()
	Expect(err).NotTo(HaveOccurred())

	// Delete all tags
	var tags []model.Tag
	err = db.Model(&tags).Select()
	Expect(err).NotTo(HaveOccurred())
	if len(tags) > 0 {
		_, err = db.Model(&tags).Delete()
		Expect(err).NotTo(HaveOccurred())
	}
})

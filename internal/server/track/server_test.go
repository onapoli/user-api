package trackserver_test

import (
	// "fmt"
	// "reflect"
	"context"

	"github.com/go-pg/pg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/twitchtv/twirp"
	"github.com/satori/go.uuid"

	pb "user-api/rpc/track"
	tagpb "user-api/rpc/tag"
	"user-api/internal/model"
)

var _ = Describe("Track server", func() {
	const already_exists_code twirp.ErrorCode = "already_exists"
	const invalid_argument_code twirp.ErrorCode = "invalid_argument"
	const not_found_code twirp.ErrorCode = "not_found"

	Describe("GetTracks", func() {
		Context("with valid uuids", func() {
			It("should respond with tracks if exist", func() {
				track := &pb.Track{Id: newTrack.Id.String()}
				tracks := &pb.TracksList{Tracks: []*pb.Track{track}}

				res, err := service.GetTracks(context.Background(), tracks)

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(len(res.Tracks)).To(Equal(1))
				Expect(res.Tracks[0].Title).To(Equal(newTrack.Title))
				Expect(res.Tracks[0].Status).To(Equal(newTrack.Status))
				Expect(res.Tracks[0].Duration).To(Equal(newTrack.Duration))
				Expect(res.Tracks[0].TrackNumber).To(Equal(newTrack.TrackNumber))
				Expect(res.Tracks[0].TrackServerId).To(Equal(newTrack.TrackServerId.String()))

				// Expect(len(res.Tags)).To(Equal(1))
				// Expect(res.Tags[0].Id).To(Equal(newGenreTag.Id.String()))
				// Expect(res.Tags[0].Type).To(Equal(newGenreTag.Type))
				// Expect(res.Tags[0].Name).To(Equal(newGenreTag.Name))

				Expect(len(res.Tracks[0].Artists)).To(Equal(1))
				Expect(res.Tracks[0].Artists[0].Id).To(Equal(newArtistUserGroup.Id.String()))
				Expect(res.Tracks[0].Artists[0].Avatar).To(Equal(newArtistUserGroup.Avatar))
				Expect(res.Tracks[0].Artists[0].DisplayName).To(Equal(newArtistUserGroup.DisplayName))

				Expect(len(res.Tracks[0].TrackGroups)).To(Equal(1))
				Expect(res.Tracks[0].TrackGroups[0].Id).To(Equal(newAlbum.Id.String()))
				Expect(res.Tracks[0].TrackGroups[0].Cover).To(Equal(newAlbum.Cover))
				Expect(res.Tracks[0].TrackGroups[0].Title).To(Equal(newAlbum.Title))
			})
			It("should not respond with tracks that don't exist", func() {
				id := uuid.NewV4()
				for id == newTrack.Id {
					id = uuid.NewV4()
				}
				track := &pb.Track{Id: id.String()}
				tracks := &pb.TracksList{Tracks: []*pb.Track{track}}

				resp, err := service.GetTracks(context.Background(), tracks)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp).NotTo(BeNil())
				Expect(len(resp.Tracks)).To(Equal(0))

			})
		})
		Context("with invalid uuid", func() {
			It("should respond with invalid_argument error", func() {
				id := "45"
				track := &pb.Track{Id: id}
				tracks := &pb.TracksList{Tracks: []*pb.Track{track}}

				resp, err := service.GetTracks(context.Background(), tracks)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())

				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
				Expect(twerr.Meta("argument")).To(Equal("id"))
			})
		})
	})

	Describe("SearchTracks", func() {
    Context("with valid query", func() {
      It("should respond with tracks", func() {
        q := &tagpb.Query{Query: "title"}
        res, err := service.SearchTracks(context.Background(), q)

        Expect(err).NotTo(HaveOccurred())
        Expect(res).NotTo(BeNil())
        Expect(len(res.Tracks)).To(Equal(1))
        Expect(res.Tracks[0].Id).To(Equal(newTrack.Id.String()))
        Expect(res.Tracks[0].Title).To(Equal(newTrack.Title))

				Expect(len(res.Tracks[0].TrackGroups)).To(Equal(1))
				Expect(res.Tracks[0].TrackGroups[0].Id).To(Equal(newAlbum.Id.String()))
				Expect(res.Tracks[0].TrackGroups[0].Title).To(Equal(newAlbum.Title))
				Expect(res.Tracks[0].TrackGroups[0].Cover).To(Equal(newAlbum.Cover))
				Expect(res.Tracks[0].TrackGroups[0].DisplayArtist).To(Equal(newAlbum.DisplayArtist))
				Expect(res.Tracks[0].TrackGroups[0].Type).To(Equal(newAlbum.Type))
				Expect(res.Tracks[0].TrackGroups[0].About).To(Equal(newAlbum.About))
        Expect(res.Tracks[0].TrackGroups[0].TotalTracks).To(Equal(int32(len(newAlbum.Tracks))))

				Expect(len(res.Tracks[0].Artists)).To(Equal(1))
        Expect(res.Tracks[0].Artists[0]).NotTo(BeNil())
        Expect(res.Tracks[0].Artists[0].Id).To(Equal(newArtistUserGroup.Id.String()))
        Expect(res.Tracks[0].Artists[0].DisplayName).To(Equal(newArtistUserGroup.DisplayName))
        Expect(res.Tracks[0].Artists[0].Avatar).To(Equal(newArtistUserGroup.Avatar))
      })
    })
    Context("with invalid query", func() {
      It("should respond with invalid error", func() {
        q := &tagpb.Query{Query: "ti"}
        resp, err := service.SearchTracks(context.Background(), q)

        Expect(resp).To(BeNil())
        Expect(err).To(HaveOccurred())

        twerr := err.(twirp.Error)
        Expect(twerr.Code()).To(Equal(invalid_argument_code))
        Expect(twerr.Meta("argument")).To(Equal("query"))
      })
    })
  })

	Describe("UpdateTrack", func() {
		Context("with valid uuid", func() {
			It("should update track if it exists", func() {
				trackServerId := uuid.NewV4()
				track := &pb.Track{
					Id: newTrack.Id.String(),
					Title: "best track ever",
					TrackNumber: 1,
					Status: "paid",
					Duration: 250.12,
					CreatorId: newUser.Id.String(),
					UserGroupId: artist.Id.String(),
					TrackServerId: trackServerId.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: artist.Id.String(),
						},
						&tagpb.RelatedUserGroup{
							Id: featuredArtist.Id.String(),
						},
					},
					Tags: []*tagpb.Tag{
						&tagpb.Tag{
							Type: "genre",
							Name: "pop",
						},
					},
				}
				_, err := service.UpdateTrack(context.Background(), track)

				Expect(err).NotTo(HaveOccurred())

				t := new(model.Track)
				err = db.Model(t).Where("id = ?", newTrack.Id).Select()
				Expect(err).NotTo(HaveOccurred())
				Expect(t.Title).To(Equal(track.Title))
				Expect(t.TrackNumber).To(Equal(track.TrackNumber))
				Expect(t.Status).To(Equal(track.Status))
				Expect(t.Duration).To(Equal(track.Duration))
				Expect(t.TrackServerId.String()).To(Equal(track.TrackServerId))

				Expect(t.UserGroupId).To(Equal(artist.Id))
        // oldUserGroup := new(model.UserGroup)
        // err = db.Model(oldUserGroup).Where("id = ?", newArtistUserGroup.Id).Select()
        // Expect(err).NotTo(HaveOccurred())
        // Expect(oldUserGroup.Tracks).NotTo(ContainElement(newTrack.Id))
        // newUserGroup := new(model.UserGroup)
        // err = db.Model(newUserGroup).Where("id = ?", artist.Id).Select()
        // Expect(err).NotTo(HaveOccurred())
        // Expect(newUserGroup.Tracks).To(ContainElement(newTrack.Id))

				Expect(len(t.Artists)).To(Equal(2))
				Expect(t.Artists).To(ContainElement(artist.Id))
				Expect(t.Artists).To(ContainElement(featuredArtist.Id))
				newArtist := new(model.UserGroup)
				err = db.Model(newArtist).Where("id = ?", artist.Id).Select()
				Expect(err).NotTo(HaveOccurred())
				Expect(newArtist.ArtistOfTracks).To(ContainElement(newTrack.Id))
				newFeaturedArtist := new(model.UserGroup)
				err = db.Model(newFeaturedArtist).Where("id = ?", featuredArtist.Id).Select()
				Expect(err).NotTo(HaveOccurred())
				Expect(newFeaturedArtist.ArtistOfTracks).To(ContainElement(newTrack.Id))

				Expect(len(t.Tags)).To(Equal(1))
				addedTag := model.Tag{Id: t.Tags[0]}
				err = db.Model(&addedTag).WherePK().Returning("*").Select()
				Expect(addedTag.Type).To(Equal("genre"))
				Expect(addedTag.Name).To(Equal("pop"))

				// unchanged
				Expect(t.CreatorId.String()).To(Equal(track.CreatorId))
				Expect(len(t.TrackGroups)).To(Equal(2))
			})
			It("should respond with not_found error if one of the artists does not exist", func() {
				userGroupId := uuid.NewV4()
				for userGroupId == newLabelUserGroup.Id || userGroupId == newArtistUserGroup.Id || userGroupId == artist.Id || userGroupId == featuredArtist.Id {
					userGroupId = uuid.NewV4()
				}
				track := &pb.Track{
					Id: newTrack.Id.String(),
					Title: "best track ever",
					TrackNumber: 1,
					Status: "paid",
					Duration: 250.12,
					CreatorId:  newTrack.Id.String(),
					UserGroupId: newArtistUserGroup.Id.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: userGroupId.String(),
						},
						&tagpb.RelatedUserGroup{
							Id: featuredArtist.Id.String(),
						},
					},
				}
				resp, err := service.UpdateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())

				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(not_found_code))
			})
			It("should respond with not_found error if user group does not exist", func() {
				userGroupId := uuid.NewV4()
				for userGroupId == newLabelUserGroup.Id || userGroupId == newArtistUserGroup.Id || userGroupId == artist.Id || userGroupId == featuredArtist.Id {
					userGroupId = uuid.NewV4()
				}
				track := &pb.Track{
					Id: newTrack.Id.String(),
					Title: "best track ever",
					TrackNumber: 1,
					Status: "paid",
					Duration: 250.12,
					CreatorId: newUser.Id.String(),
					UserGroupId: userGroupId.String(),
				}
				resp, err := service.UpdateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(not_found_code))
			})
			It("should respond with not_found error if track does not exist", func() {
				id := uuid.NewV4()
				for id == newTrack.Id {
					id = uuid.NewV4()
				}
				track := &pb.Track{
					Id: id.String(),
					Title: "best track ever",
					TrackNumber: 1,
					Status: "paid",
					Duration: 250.12,
					CreatorId: newUser.Id.String(),
					UserGroupId: newArtistUserGroup.Id.String(),
				}
				resp, err := service.UpdateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())

				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(not_found_code))
			})

		})
		Context("with invalid uuid", func() {
			It("should respond with invalid_argument error", func() {
				id := "45"
				track := &pb.Track{
					Id: id,
					Title: "best track ever",
					TrackNumber: 1,
					Status: "paid",
					Duration: 250.12,
					CreatorId: newUser.Id.String(),
					UserGroupId: newArtistUserGroup.Id.String(),
				}
				resp, err := service.UpdateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())

				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
				Expect(twerr.Meta("argument")).To(Equal("id"))
			})
			It("should not update track if user_group_id is invalid", func() {
				track := &pb.Track{
					Id: newTrack.Id.String(),
					Title: "best track ever",
					TrackNumber: 1,
					Status: "paid",
					Duration: 250.12,
					CreatorId: newUser.Id.String(),
					UserGroupId: "",
				}
				resp, err := service.UpdateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())

				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
				Expect(twerr.Meta("argument")).To(Equal("user_group_id"))
			})
			It("should not update track if one of the artists' ids is invalid", func() {
				track := &pb.Track{
					Id: newTrack.Id.String(),
					Title: "best track ever",
					TrackNumber: 1,
					Status: "paid",
					Duration: 250.12,
					CreatorId: newUser.Id.String(),
					UserGroupId: newArtistUserGroup.Id.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: "2a5",
						},
						&tagpb.RelatedUserGroup{
							Id: featuredArtist.Id.String(),
						},
					},
				}
				resp, err := service.UpdateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())

				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
				Expect(twerr.Meta("argument")).To(Equal("user_group id"))
			})
		})
	})

	Describe("CreateTrack", func() {
		Context("with all required attributes", func() {
			It("should create a new track", func() {
				track := &pb.Track{
					Title: "best track ever",
					TrackNumber: 1,
					Status: "paid",
					CreatorId: newUser.Id.String(),
					UserGroupId: newArtistUserGroup.Id.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: newArtistUserGroup.Id.String(),
						},
					},
					Tags: []*tagpb.Tag{
						&tagpb.Tag{
							Type: "genre",
							Name: "rock",
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp).NotTo(BeNil())

				Expect(resp.Id).NotTo(BeNil())
				Expect(resp.Title).To(Equal(track.Title))
				Expect(resp.TrackNumber).To(Equal(track.TrackNumber))
				Expect(resp.Status).To(Equal(track.Status))
				Expect(resp.CreatorId).To(Equal(track.CreatorId))
				Expect(resp.UserGroupId).To(Equal(track.UserGroupId))
				Expect(len(resp.Tags)).To(Equal(1))
				Expect(resp.Tags[0].Id).NotTo(Equal(""))
				Expect(resp.Tags[0].Type).To(Equal("genre"))
				Expect(resp.Tags[0].Name).To(Equal("rock"))
				Expect(len(resp.Artists)).To(Equal(1))
				Expect(resp.Artists[0].Id).To(Equal(newArtistUserGroup.Id.String()))
				Expect(resp.Artists[0].DisplayName).To(Equal(newArtistUserGroup.DisplayName))
				Expect(resp.Artists[0].Avatar).To(Equal(newArtistUserGroup.Avatar))

				artist := new(model.UserGroup)
				err = db.Model(artist).Where("id = ?", newArtistUserGroup.Id).Select()
				Expect(err).NotTo(HaveOccurred())
				trackId, err := uuid.FromString(resp.Id)
				Expect(err).NotTo(HaveOccurred())
				Expect(artist.ArtistOfTracks).To(ContainElement(trackId))
			})
		})

		Context("with missing required attributes", func() {
			It("should not create a track without title", func() {
				track := &pb.Track{
					Title: "",
					TrackNumber: 1,
					Status: "paid",
					CreatorId: newUser.Id.String(),
					UserGroupId: newArtistUserGroup.Id.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: newArtistUserGroup.Id.String(),
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())

				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
				Expect(twerr.Meta("argument")).To(Equal("title"))
			})
			It("should not create a track without status", func() {
				track := &pb.Track{
					Title: "title",
					TrackNumber: 1,
					Status: "",
					CreatorId: newUser.Id.String(),
					UserGroupId: newArtistUserGroup.Id.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: newArtistUserGroup.Id.String(),
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
				Expect(twerr.Meta("argument")).To(Equal("status"))
			})
			It("should not create a track without creator_id", func() {
				track := &pb.Track{
					Title: "title",
					TrackNumber: 1,
					Status: "paid",
					CreatorId: "",
					UserGroupId: newArtistUserGroup.Id.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: newArtistUserGroup.Id.String(),
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())

				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
				Expect(twerr.Meta("argument")).To(Equal("creator_id"))
			})
			It("should not create a track without user_group_id", func() {
				track := &pb.Track{
					Title: "title",
					TrackNumber: 1,
					Status: "paid",
					CreatorId: newUser.Id.String(),
					UserGroupId: "",
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: newArtistUserGroup.Id.String(),
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
				Expect(twerr.Meta("argument")).To(Equal("user_group_id"))
			})
			It("should not create a track without track_number", func() {
				track := &pb.Track{
					Title: "title",
					TrackNumber: 0,
					Status: "free",
					CreatorId: newUser.Id.String(),
					UserGroupId: newArtistUserGroup.Id.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: newArtistUserGroup.Id.String(),
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
				Expect(twerr.Meta("argument")).To(Equal("track_number"))
			})
		})

		Context("with invalid attributes", func() {
			It("should not create a track if creator_id is invalid", func() {
				track := &pb.Track{
					Title: "title",
					TrackNumber: 1,
					Status: "free",
					CreatorId: "12a",
					UserGroupId: newArtistUserGroup.Id.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: newArtistUserGroup.Id.String(),
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
			})
			It("should not create a track if user_group_id is invalid", func() {
				track := &pb.Track{
					Title: "title",
					TrackNumber: 1,
					Status: "free",
					CreatorId: newUser.Id.String(),
					UserGroupId: "abc1",
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: newArtistUserGroup.Id.String(),
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
			})
			It("should not create a track if creator does not exist", func() {
				userId := uuid.NewV4()
				for userId == newUser.Id {
					userId = uuid.NewV4()
				}
				track := &pb.Track{
					Title: "title",
					TrackNumber: 1,
					Status: "free",
					CreatorId: userId.String(),
					UserGroupId: newArtistUserGroup.Id.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: newArtistUserGroup.Id.String(),
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(not_found_code))
			})
			It("should not create a track if user_group does not exist", func() {
				userGroupId := uuid.NewV4()
				for userGroupId == newLabelUserGroup.Id || userGroupId == newArtistUserGroup.Id || userGroupId == artist.Id || userGroupId == featuredArtist.Id {
					userGroupId = uuid.NewV4()
				}
				track := &pb.Track{
					Title: "title",
					TrackNumber: 1,
					Status: "free",
					CreatorId: newUser.Id.String(),
					UserGroupId: userGroupId.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: newArtistUserGroup.Id.String(),
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(not_found_code))
			})
			It("should not create a track if one of the artists does not exist", func() {
				userGroupId := uuid.NewV4()
				for userGroupId == newLabelUserGroup.Id || userGroupId == newArtistUserGroup.Id || userGroupId == artist.Id || userGroupId == featuredArtist.Id {
					userGroupId = uuid.NewV4()
				}
				track := &pb.Track{
					Title: "title",
					TrackNumber: 1,
					Status: "free",
					CreatorId: newUser.Id.String(),
					UserGroupId: userGroupId.String(),
					Artists: []*tagpb.RelatedUserGroup{
						&tagpb.RelatedUserGroup{
							Id: userGroupId.String(),
						},
					},
				}
				resp, err := service.CreateTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())
				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(not_found_code))
			})
		})
	})

	Describe("DeleteTrack", func() {
		Context("with valid uuid", func() {
			It("should delete track if it exists", func() {
				track := &pb.Track{Id: newTrack.Id.String()}

				trackToDelete := new(model.Track)
				err := db.Model(trackToDelete).Where("id = ?", newTrack.Id).Select()
				Expect(err).NotTo(HaveOccurred())

				_, err = service.DeleteTrack(context.Background(), track)

				Expect(err).NotTo(HaveOccurred())

				// owner := new(model.UserGroup)
				// err = db.Model(owner).Where("id = ?", trackToDelete.UserGroupId).Select()
				// Expect(err).NotTo(HaveOccurred())
				// Expect(owner.Tracks).NotTo(ContainElement(trackToDelete.Id))

				var users []*model.User
				err = db.Model(&users).
					Where("id in (?)", pg.In(trackToDelete.FavoriteOfUsers)).
					Select()
				for _, user := range users {
					Expect(user.FavoriteTracks).NotTo(ContainElement(trackToDelete.Id))
				}

				var artists []*model.UserGroup
				err = db.Model(&artists).
					Where("id in (?)", pg.In(trackToDelete.Artists)).
					Select()
				for _, artist := range artists {
					Expect(artist.ArtistOfTracks).NotTo(ContainElement(trackToDelete.Id))
				}

				var trackGroups []*model.TrackGroup
				err = db.Model(&trackGroups).
					Where("id in (?)", pg.In(trackToDelete.TrackGroups)).
					Select()
				for _, trackGroup := range trackGroups {
					Expect(trackGroup.Tracks).NotTo(ContainElement(trackToDelete.Id))
				}

				var tracks []model.Track
				err = db.Model(&tracks).
					Where("id in (?)", pg.In([]uuid.UUID{trackToDelete.Id})).
					Select()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(tracks)).To(Equal(0))
			})
			It("should respond with not_found error if track does not exist", func() {
				id := uuid.NewV4()
				for id == newTrack.Id {
					id = uuid.NewV4()
				}
				track := &pb.Track{Id: id.String()}
				resp, err := service.DeleteTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())

				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(not_found_code))
			})
		})
		Context("with invalid uuid", func() {
			It("should respond with invalid_argument error", func() {
				id := "45"
				track := &pb.Track{Id: id}
				resp, err := service.DeleteTrack(context.Background(), track)

				Expect(resp).To(BeNil())
				Expect(err).To(HaveOccurred())

				twerr := err.(twirp.Error)
				Expect(twerr.Code()).To(Equal(invalid_argument_code))
				Expect(twerr.Meta("argument")).To(Equal("id"))
			})
		})
	})
})

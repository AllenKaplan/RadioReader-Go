package RadioReader

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAllSongs(context.Context, *GetRequest) (*Response, error) {
	var songs []*Song
	if err := repo.db.Find(&songs).Error; err != nil {
		return nil, err
	}
	return &Response{
		Created: true,
		Songs: songs,
	}, nil
}

func (repo *UserRepository) GetSongPlays(ctx context.Context, song *Song) (*SongCountResponse, error) {
	var plays int32
	if err := repo.db.Where(fmt.Sprintf("artist = %s AND title = %s", song.Artist, song.Title)).Count(&plays).Error; err != nil {
		return nil, err
	}
	return &SongCountResponse{
		Song: song,
		Plays: plays,
	}, nil
}

func (repo *UserRepository) GetSongsByArtist(ctx context.Context, req *GetSongsByArtistRequest) (*Response, error) {
	var songs []*Song
	if err := repo.db.Where(fmt.Sprintf("artist = %s", req.Artist)).Error; err != nil {
		return nil, err
	}
	return &Response{
		Created: true,
		Songs: songs,
	}, nil
}

func (repo *UserRepository) GetSongsByTitle(ctx context.Context, req *GetSongsByTitleRequest) (*Response, error) {
	var songs []*Song
	if err := repo.db.Where(fmt.Sprintf("title = %s", req.Title)).Error; err != nil {
		return nil, err
	}
	return &Response{
		Created: true,
		Songs: songs,
	}, nil
}

func (repo *UserRepository) AddSong(ctx context.Context, in *Song) (*Response, error) {
	if err := repo.db.Create(in).Error; err != nil {
		return &Response{
			Created: false,
			Song: in,
		},err
	}
	return &Response{
		Created: true,
		Song: in,
	},nil
}
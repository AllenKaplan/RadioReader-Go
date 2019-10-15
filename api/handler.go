package RadioReader

import (
	"context"
	"log"

)

type Server struct {
	Songs []*Song
}

func main(){
	db, err := CreateConnection()
	defer db.Close()
}

func (s *Server) GetAllSongs(context.Context, *GetRequest) (*Response, error) {
	return &Response{Created: true, Songs: s.Songs}, nil
}

func (s *Server) GetSongPlays(context.Context, *Song) (*SongCountResponse, error) {
	panic("implement me")
}

func (s *Server) GetSongsByArtist(context.Context, *GetSongsByArtistRequest) (*Response, error) {
	panic("implement me")
}

func (s *Server) GetSongsByTitle(context.Context, *GetSongsByTitleRequest) (*Response, error) {
	panic("implement me")
}

func (s *Server) AddSong(ctx context.Context, in *Song) (*Response, error) {
	log.Printf("Recieved song %s by %s", in.Title, in.Artist)
	song := &Song{
		Title: in.Title,
		Artist: in.Artist,
	}
	s.Songs = append(s.Songs, song)

	response := &Response{Created: true, Song: song}
	return response, nil
}

func (s *Server) CountSongs(ctx context.Context, in *Song) (*SongCountResponse, error){
	var count int32 = 0
	for _, song := range s.Songs {
		if in.Title == song.Title {
			count++
		}
	}

	return &SongCountResponse{Song: in, Plays: count}, nil
}
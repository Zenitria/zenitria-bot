package api

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func StartAPI(s *discordgo.Session) {
	r := chi.NewRouter()
	r.Route("/send-dm", func(r chi.Router) {
		r.Post("/{id}", func(w http.ResponseWriter, r *http.Request) {
			var data struct {
				Title       string `json:"title"`
				Description string `json:"description"`
				Thumbnail   string `json:"thumbnail"`
			}

			id := chi.URLParam(r, "id")

			err := json.NewDecoder(r.Body).Decode(&data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("BAD REQUEST"))
				return
			}

			embed := &discordgo.MessageEmbed{
				Title:       data.Title,
				Description: data.Description,
				Color:       0xBE4DFF,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: data.Thumbnail,
				},
			}

			channel, _ := s.UserChannelCreate(id)
			s.ChannelMessageSendEmbed(channel.ID, embed)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("SUCCESS"))
		})
	})
	r.Route("/claim-reminder", func(r chi.Router) {
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")

			embed := &discordgo.MessageEmbed{
				Title:       "⏰・Next Claim",
				Description: fmt.Sprintf("You can claim your 15 minute reward now! [Click here](https://zenitria.dev/bot/claim/%s) to claim it.", id),
				Color:       0xBE4DFF,
			}

			channel, _ := s.UserChannelCreate(id)

			go func() {
				time.AfterFunc(15*time.Minute, func() {
					s.ChannelMessageSendEmbed(channel.ID, embed)
				})
			}()

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("SUCCESS"))
		})
	})

	fmt.Println("API is running on :9999")
	http.ListenAndServe(":9999", r)
}

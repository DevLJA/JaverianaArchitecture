package controllers

import (
	"CircuitBreaker-LauraJoya/dtos"
	"encoding/json"
	"github.com/cjsaylor/chessimage"
	"image/png"
	"net/http"
)

type ChessController interface {
	GetImageFEN(w http.ResponseWriter, r *http.Request)
}

func NewChessController() ChessController {
	return &chessControllerImp{}
}

type chessControllerImp struct {
}

func positionToTile(pos string) chessimage.Tile {
	if len(pos) != 2 {
		return chessimage.NoTile // invalid input
	}

	col := pos[0] - 'a' // Set 'a' a 0, 'b' a 1, etc.
	row := pos[1] - '1' // Set '1' a 0, '2' a 1, etc.

	// Get index position
	return chessimage.Tile(8*(7-row) + col)
}

func (u chessControllerImp) GetImageFEN(w http.ResponseWriter, r *http.Request) {
	var fenRequest dtos.InputFen
	err := json.NewDecoder(r.Body).Decode(&fenRequest)
	if err != nil {
		http.Error(w, "error trying decode input", http.StatusBadRequest)
		return
	}

	board, _ := chessimage.NewRendererFromFEN(fenRequest.Fen)
	if len(fenRequest.LastMove) > 0 {
		from := positionToTile(fenRequest.LastMove[:2])
		to := positionToTile(fenRequest.LastMove[2:])
		board.SetLastMove(chessimage.LastMove{
			From: from,
			To:   to,
		})
	}

	image, _ := board.Render(chessimage.Options{AssetPath: "./assets/"})
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, image)
}

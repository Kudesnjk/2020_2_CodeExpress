package consts

import (
	"errors"
	"net/http"
)

const (
	ErrInternal = iota
	ErrBadRequest
	ErrEmailAlreadyExist
	ErrNameAlreadyExist
	ErrIncorrectLoginOrPassword
	ErrNotAuthorized
	ErrNoEmail
	ErrNoUsername
	ErrNoAvatar
	ErrWrongOldPassword
	ErrNewPasswordIsOld
	ErrArtistNotExist
	ErrTrackNotExist
	ErrAlbumNotExist
	ErrTitleAlreadyExist
	ErrNoFavoritesTracks
<<<<<<< HEAD
	ErrPlaylistNotExist
=======
	ErrEmptySearchQuery
>>>>>>> ff8a9b4... CP-92: Реализован поиск по названию альбома, имени артиста, и названию трека.
)

var Errors = map[int]error{
	ErrInternal:                 errors.New("Internal server error"),
	ErrBadRequest:               errors.New("Bad request received"),
	ErrEmailAlreadyExist:        errors.New("Email already exists"),
	ErrNameAlreadyExist:         errors.New("Name already exists"),
	ErrIncorrectLoginOrPassword: errors.New("Incorrect login or password"),
	ErrNotAuthorized:            errors.New("Not authorized"),
	ErrNoEmail:                  errors.New("No email field"),
	ErrNoUsername:               errors.New("No username field"),
	ErrNoAvatar:                 errors.New("No avatar field"),
	ErrWrongOldPassword:         errors.New("Wrong old password"),
	ErrNewPasswordIsOld:         errors.New("New password matches old"),
	ErrArtistNotExist:           errors.New("Artist not found"),
	ErrTrackNotExist:            errors.New("Track not found"),
	ErrAlbumNotExist:            errors.New("Album not found"),
	ErrTitleAlreadyExist:        errors.New("Title already exists"),
	ErrNoFavoritesTracks:        errors.New("User has no favorite tracks"),
<<<<<<< HEAD
	ErrPlaylistNotExist:         errors.New("Playlist not found"),
=======
	ErrEmptySearchQuery:         errors.New("Empty search query"),
>>>>>>> ff8a9b4... CP-92: Реализован поиск по названию альбома, имени артиста, и названию трека.
}

var StatusCodes = map[int]int{
	ErrInternal:                 http.StatusInternalServerError,
	ErrBadRequest:               http.StatusBadRequest,
	ErrEmailAlreadyExist:        http.StatusForbidden,
	ErrNameAlreadyExist:         http.StatusForbidden,
	ErrIncorrectLoginOrPassword: http.StatusNotFound,
	ErrNotAuthorized:            http.StatusNotFound,
	ErrNoEmail:                  http.StatusBadRequest,
	ErrNoUsername:               http.StatusBadRequest,
	ErrNoAvatar:                 http.StatusBadRequest,
	ErrWrongOldPassword:         http.StatusBadRequest,
	ErrNewPasswordIsOld:         http.StatusBadRequest,
	ErrArtistNotExist:           http.StatusNotFound,
	ErrTrackNotExist:            http.StatusNotFound,
	ErrAlbumNotExist:            http.StatusNotFound,
	ErrTitleAlreadyExist:        http.StatusForbidden,
	ErrNoFavoritesTracks:        http.StatusNotFound,
<<<<<<< HEAD
	ErrPlaylistNotExist:         http.StatusNotFound,
=======
	ErrEmptySearchQuery:         http.StatusBadRequest,
>>>>>>> ff8a9b4... CP-92: Реализован поиск по названию альбома, имени артиста, и названию трека.
}

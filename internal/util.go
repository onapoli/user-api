package internal

import (
  "fmt"
  "strings"

	"github.com/satori/go.uuid"
  "github.com/go-pg/pg"
  "github.com/twitchtv/twirp"
)

func ConvertUuidToStrArray(uuids []uuid.UUID) ([]string) {
  strArray := make([]string, len(uuids))
  for i := range uuids {
    strArray[i] = uuids[i].String()
  }
  return strArray
}

func CheckError(err error, table string) (twirp.Error) {
	if err != nil {
		if err == pg.ErrNoRows {
			return twirp.NotFoundError(fmt.Sprintf("%s does not exist", table))
		}
    twerr, ok := err.(twirp.Error)
    if ok && twerr.Meta("argument") == "id" {
      argument := "id"
      if table != "" {
        argument = table + " id"
      }
      return twirp.InvalidArgumentError(argument, "must be a valid uuid")
    }
		pgerr, ok := err.(pg.Error)
		if ok {
			code := pgerr.Field('C')
			name := pgerr.Field('n')
			var message string
			if code == "23505" { // unique_violation
				message = strings.TrimPrefix(strings.TrimSuffix(name, "_key"), fmt.Sprintf("%ss_", table))
				return twirp.NewError("already_exists", message)
      } else if code == "23503" { // foreign_key_violation
        message = pgerr.Field('M')
        return twirp.NotFoundError(message)
      } else {
				message = pgerr.Field('M')
        fmt.Println(twirp.NewError("unknown", message))
				return twirp.NewError("unknown", message)
			}
		}
		return twirp.NewError("unknown", err.Error())
	}
	return nil
}

func GetUuidFromString(id string) (uuid.UUID, twirp.Error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return uuid.UUID{}, twirp.InvalidArgumentError("id", "must be a valid uuid")
	}
	return uid, nil
}

func Difference(a, b []uuid.UUID) []uuid.UUID {
    mb := map[uuid.UUID]bool{}
    for _, x := range b {
        mb[x] = true
    }
    ab := []uuid.UUID{}
    for _, x := range a {
        if _, ok := mb[x]; !ok {
            ab = append(ab, x)
        }
    }
    return ab
}

func RemoveDuplicates(elements []uuid.UUID) []uuid.UUID {
  // Use map to record duplicates as we find them
  encountered := map[uuid.UUID]bool{}
  result := []uuid.UUID{}

  for v := range elements {
    if encountered[elements[v]] == true {
      // Do not add duplicate
    } else {
      // Record this element as an encountered element
      encountered[elements[v]] = true
      // Append to result slice
      result = append(result, elements[v])
    }
  }
  // Return the new slice.
  return result
}

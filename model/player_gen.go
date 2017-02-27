// Auto generated by genmodel. DO NOT EDIT.
package model

import (
	"encoding/json"

	"github.com/noctf/scoreserver/db"
)

func (player Player) MarshalJSON() ([]byte, error) {
	buf, err := json.Marshal(player)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (player *Player) Load(tx *db.Tx, id string) error {
	pdb := db.Player{}
	if err := pdb.LoadByID(tx, id); err != nil {
		return err
	}
	if err := player.FromRow(&pdb); err != nil {
		return err
	}
	return nil
}

func (player *Player) FromRow(pdb *db.Player) error {
	player.ID = pdb.ID
	player.Name = pdb.Name
	if pdb.Email.Valid {
		player.Email = pdb.Email.string
	}
	player.TeamID = pdb.TeamID
	player.Password = pdb.Password
	return nil
}

func (player *Player) ToRow(pdb *db.Player) error {
	pdb.ID = player.ID
	pdb.Name = player.Name
	pdb.Email.Valid = true
	pdb.Email = player.Email
	pdb.TeamID = player.TeamID
	pdb.Password = player.Password
	return nil
}

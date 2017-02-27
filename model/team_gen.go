// Auto generated by genmodel. DO NOT EDIT.
package model

import (
	"encoding/json"

	"github.com/noctf/scoreserver/db"
)

func (team Team) MarshalJSON() ([]byte, error) {
	buf, err := json.Marshal(team)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (team *Team) Load(tx *db.Tx, id string) error {
	pdb := db.Team{}
	if err := pdb.LoadByID(tx, id); err != nil {
		return err
	}
	if err := team.FromRow(&pdb); err != nil {
		return err
	}
	return nil
}

func (team *Team) FromRow(pdb *db.Team) error {
	team.ID = pdb.ID
	team.Name = pdb.Name
	return nil
}

func (team *Team) ToRow(pdb *db.Player) error {
	pdb.ID = team.ID
	pdb.Name = team.Name
	return nil
}
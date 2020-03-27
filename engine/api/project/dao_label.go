package project

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	"github.com/ovh/cds/sdk"
)

type dbLabel sdk.Label

// Labels return list of labels given a project ID
func Labels(db gorp.SqlExecutor, projectID int64) ([]sdk.Label, error) {
	var labels []sdk.Label
	query := `
	SELECT project_label.*
		FROM project_label
		WHERE project_label.project_id = $1
		ORDER BY project_label.name
	`
	if _, err := db.Select(&labels, query, projectID); err != nil {
		if err == sql.ErrNoRows {
			return labels, nil
		}
		return labels, sdk.WrapError(err, "Cannot load labels")
	}

	return labels, nil
}

// LabelByName return a label given his name and project id
func LabelByName(db gorp.SqlExecutor, projectID int64, labelName string) (sdk.Label, error) {
	var label sdk.Label
	err := db.SelectOne(&label, "SELECT project_label.* FROM project_label WHERE project_id = $1 AND name = $2", projectID, labelName)

	return label, err
}

// DeleteLabel delete a label given a label ID
func DeleteLabel(db gorp.SqlExecutor, labelID int64) error {
	query := "DELETE FROM project_label WHERE id = $1"
	if _, err := db.Exec(query, labelID); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return sdk.WrapError(err, "Cannot delete labels")
	}

	return nil
}

// InsertLabel insert a label
func InsertLabel(db gorp.SqlExecutor, label *sdk.Label) error {
	if err := label.IsValid(); err != nil {
		return err
	}

	lbl := dbLabel(*label)
	if err := db.Insert(&lbl); err != nil {
		return sdk.WrapError(err, "Cannot insert labels")
	}
	*label = sdk.Label(lbl)

	return nil
}

// UpdateLabel update a label
func UpdateLabel(db gorp.SqlExecutor, label *sdk.Label) error {
	if err := label.IsValid(); err != nil {
		return err
	}

	lbl := dbLabel(*label)
	if _, err := db.Update(&lbl); err != nil {
		return sdk.WrapError(err, "Cannot update labels")
	}
	*label = sdk.Label(lbl)

	return nil
}
package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
)

func UpdateConfigurationItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		dec      *json.Decoder
		item     *ConfigurationItem
		lookupID string
		details  *somaproto.DeploymentDetails
		err      error
	)

	if _, err = uuid.FromString(params.ByName("item")); err != nil {
		dispatchBadRequest(&w, err.Error())
		return
	}

	dec = json.NewDecoder(r.Body)
	if err = dec.Decode(details); err != nil {
		dispatchUnprocessable(&w, err.Error())
		return
	}

	if lookupID, item, err = Itemize(details); err != nil {
		dispatchUnprocessable(&w, err.Error())
		return
	}

	if item.ConfigurationItemId.String() != params.ByName("item") {
		dispatchBadRequest(&w, "Mismatching ConfigurationItemID")
		return
	}

	if err = updateItem(item, lookupID); err != nil {
		dispatchInternalServerError(&w, err.Error())
		return
	}
	dispatchNoContent(&w)
}

func updateItem(item *ConfigurationItem, lookupID string) error {
	var (
		itemID string
		err    error
		jsonb  []byte
	)

	if jsonb, err = json.Marshal(item); err != nil {
		return err
	}

	// since this was an explicit update request, non-existance is a
	// hard error
	if err = Eye.run.check_item.QueryRow(item.ConfigurationItemId.String()).Scan(&itemID); err != nil {
		return err
	}
	if itemID != item.ConfigurationItemId.String() {
		panic(`Database corrupted`)
	}

	_, err = Eye.run.update_item.Exec(
		item.ConfigurationItemId.String(),
		lookupID,
		jsonb,
	)
	return err
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix
package	routes

import (
	"time"
	"net/http"
	"strings"
	"appengine"
	"appengine/datastore"

	view "viewengine"

	"github.com/Shaked/gomobiledetect"
)
type RenderData struct{
	Mode string
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	detect := gomobiledetect.NewMobileDetect(r, nil)
	if detect.IsMobile() {
		// do some mobile stuff 
		view.Render(w, "main", RenderData{Mode: "mobile"});
		return;
	}
	/*
	if detect.IsTablet() {
		// do some tablet stuff 
		view.Render(w, "main", RenderData{Mode:"pc"});
		return;
	}
	deviceProperty := "iPhone"
	if detect.VersionFloat(deviceProperty) > 6 {
		// do something with iPhone v6 
		view.Render(w, "mobile", RenderData{Mode: "mobile"});
		return;
	}*/
	view.Render(w, "main", RenderData{Mode:"pc"});
	return;
}
type Team struct {
	Date time.Time

	Name string
	Email string
	Phonenumber string

	PartnerName string
	PartnerEmail string
	PartnerPhonenumber string

	Tech string

	TeamDescrition string
	TeamWish string
}
func register(w http.ResponseWriter, req *http.Request){
	req.ParseForm();
	context := appengine.NewContext(req);

	team := Team {
		Date : time.Now(),
		Name : req.FormValue("name"),
		Email : req.FormValue("email"),
		Phonenumber : req.FormValue("phone"),

		PartnerName : req.FormValue("partner-name"),
		PartnerEmail : req.FormValue("partner-email"),
		PartnerPhonenumber : req.FormValue("partner-phone"),
		Tech : strings.Join(req.Form["tech"],","),
		TeamDescrition : req.FormValue("descript"),
		TeamWish : req.FormValue("wish"),
	}
	// TODO make vaildater
	key := datastore.NewIncompleteKey(context, "3rd", getDataStoreKey(context));

	_, err := datastore.Put(context, key, &team);

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		view.Render(w, "main", RenderData{Mode:"pc"});
		return
	}
}
func getDataStoreKey(context appengine.Context) *datastore.Key {
	return datastore.NewKey(context, "opencamp", "third", 0, nil);
}

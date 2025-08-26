package main

import (
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	"github.com/pistolricks/riman-api/internal/profile_v1"
)

func (app *application) getLocalizationGetTimezones(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.LocalizationApiLocalizationGetTimezonesOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetTimezones(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"timezones": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationInvalidateLocalization(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.LocalizationApiLocalizationInvalidateLocalizationOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationInvalidateLocalization(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetAllCountries(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.LocalizationApiLocalizationGetAllCountriesOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetAllCountries(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"countries": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetCountryGroups(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.LocalizationApiLocalizationGetCountryGroupsOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetCountryGroups(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"groups": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetGatewayAddress(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.LocalizationApiLocalizationGetGatewayAddressOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetGatewayAddress(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"gatewayAddress": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetStreetsForArea(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	areaIdStr := q.Get("areaId")
	areaId64, _ := strconv.ParseInt(areaIdStr, 10, 32)
	opts := &profile_v1.LocalizationApiLocalizationGetStreetsForAreaOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetStreetsForArea(r.Context(), int32(areaId64), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"streets": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetBillingCountries(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.LocalizationApiLocalizationGetBillingCountriesOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetBillingCountries(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"billingCountries": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetEuGatewayCountries(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.LocalizationApiLocalizationGetEuGatewayCountriesOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetEuGatewayCountries(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"euGatewayCountries": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetRegions(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	opts := &profile_v1.LocalizationApiLocalizationGetRegionsOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetRegions(r.Context(), countryCode, opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"regions": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationComunaCheck(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("stateOrProvince")
	opts := &profile_v1.LocalizationApiLocalizationComunaCheckOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationComunaCheck(r.Context(), state, opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetMinEnrollmentAge(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	opts := &profile_v1.LocalizationApiLocalizationGetMinEnrollmentAgeOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetMinEnrollmentAge(r.Context(), country, opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"minEnrollmentAge": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetAreasForCity(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	cityId := q.Get("cityId")
	countryCode := q.Get("countryCode")
	opts := &profile_v1.LocalizationApiLocalizationGetAreasForCityOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetAreasForCity(r.Context(), cityId, countryCode, opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"areas": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetCitiesForCountry(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	countryCode := q.Get("countryCode")
	isoCode := q.Get("isoCode")
	opts := &profile_v1.LocalizationApiLocalizationGetCitiesForCountryOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetCitiesForCountry(r.Context(), countryCode, isoCode, opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"cities": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetShippingCountries(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	opts := &profile_v1.LocalizationApiLocalizationGetShippingCountriesOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetShippingCountries(r.Context(), countryCode, opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"shippingCountries": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationLookupJPAddressByPostalCode(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &profile_v1.LocalizationApiLocalizationLookupJPAddressByPostalCodeOpts{}
	if v := q.Get("postalCode"); v != "" { opts.PostalCode = optional.NewString(v) }
	res, h, err := app.profileV1.LocalizationApi.LocalizationLookupJPAddressByPostalCode(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"jpAddress": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetCountryZipFormats(w http.ResponseWriter, r *http.Request) {
	cultureCode := r.URL.Query().Get("cultureCode")
	opts := &profile_v1.LocalizationApiLocalizationGetCountryZipFormatsOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetCountryZipFormats(r.Context(), cultureCode, opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"zipFormats": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetIdTypes(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	opts := &profile_v1.LocalizationApiLocalizationGetIdTypesOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetIdTypes(r.Context(), country, opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"idTypes": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getLocalizationGetCountries(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.LocalizationApiLocalizationGetCountriesOpts{}
	res, h, err := app.profileV1.LocalizationApi.LocalizationGetCountries(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"countries": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}

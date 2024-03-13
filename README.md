# trackerGPS

## Summary
Create an application consisting of a Go server and a Vue user interface which includes data retrieval, storage, and presentation.

# Design
## Frontend
- Made dashboard showing the device map publicly available for the purposes of making this as accessible as possible for the interview process, however, in a production environment, I would protect this route behind a login.
## Backend
### One Step GPS API
- Didn't use [One Step GPS Webhooks](https://track.onestepgps.com/v3/apidoc-webhooks/) because there didn't seem to be enough information to implement it.

## Preliminary Stuff
- [ ] Research the OneStepGPS API and determine what data is available and how to access it.
- [ ] Implement authentication (basic)
- [ ] Make two Google Maps API keys (one for development and one for production)
- [ ] 404 Missing pages should have links to all of the actual pages (home, dashboard, about, etc)

## The Go backend should do the following:
- [ ] Request information from https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=[APIKEY]
- [ ] Provide some storage mechanism for user preferences or any other relevant data (exact storage mechanism is up to you)
- [ ] Provide endpoints for anything else you need in the UI

## The Vue UI should do the following:
- [ ] List information regarding the devices returned from the API. Useful information would be the name, the current position, and its active state or drive status.
- [ ] Allow the user to edit preferences and save those preferences to the server. These preferences can be anything you want, some recommendations are: sort order, hiding specific devices from view, user-uploaded icons for each device.
- [ ] Show the current location of the devices on a map (Google Maps is a good choice) and update the map in real time as the API returns new data.

## Useful Links
- [Creating a Vue.js 3 App with Vite.js and Vue Router](https://ochner.com.br/posts/creating-a-vuejs3-app-with-vitejs-and-vue-router)
- [Flowbite - Tailwind CSS component library](https://flowbite.com/docs/getting-started/introduction/)
- [Google Maps Javascript API Docs](https://developers.google.com/maps/documentation/javascript)

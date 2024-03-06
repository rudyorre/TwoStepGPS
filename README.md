# trackerGPS

## Summary
Create an application consisting of a Go server and a Vue user interface which includes data retrieval, storage, and presentation.

## Preliminary Stuff
- [ ] Research the OneStepGPS API and determine what data is available and how to access it.
- [ ] Implement authentication (basic)

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

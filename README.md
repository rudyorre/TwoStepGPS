# TwoStepGPS

## Summary
A comprehensive GPS tracking dashboard that provides real-time visualiztion of GPS devices on a map.

## Features
- Dashboard for viewing GPS devices on a map
- Short polling to get the latest device locations
- Hide/show devices on the map
- Nickname devices
- Change the device color on the map
- Authentication (to save the above preferences)

## Architecture

### Frontend

The frontend is a Vue3 + Vite application. It's a Single Page Application (SPA) that utilizes the Google Maps API for the dashboard view. It allows signed in users to customize several properties about the dashboard.

### Backend

The backend is built with Go and uses PostgresDB for data storage. It handles JWT authentication, integrates with the OneStepGPS API, and interfaces with the Postgres database.

### Infrastructure

The frontend is hosted as a Vue/Vite app on Vercel, while the backend Go server is on a serverless fly.io VM which mounts the Postgres volume on startup. This is analogous to AWS Lambda paired with RDS.

## Design Decisions
- Made dashboard showing the device map publicly available for the purposes of making this as accessible as possible for the interview process, however, in a real production environment, I would protect this route behind a login.
- Didn't use [One Step GPS Webhooks](https://track.onestepgps.com/v3/apidoc-webhooks/) because there didn't seem to be enough information to implement it, however I would use this in a production environment to avoid short polling.
- The responsiveness is not perfect, as the mobile view is not optimized.

## Helpful Resources
- [Creating a Vue.js 3 App with Vite.js and Vue Router](https://ochner.com.br/posts/creating-a-vuejs3-app-with-vitejs-and-vue-router)
- [Flowbite - Tailwind CSS component library](https://flowbite.com/docs/getting-started/introduction/)
- [Google Maps Javascript API Docs](https://developers.google.com/maps/documentation/javascript)
- [Postman](https://www.postman.com/) for testing the backend APIs

# htmx-note-app
Simple note app using HTMX, Go &amp; TEMPL.

## What's What
### App
The code for the app is located in the `app/` directory. The README in that directory has more information on how to run the app. 

This app is fairly simple, and just has basic CRUD operations, so at this time there are no unit tests. If any logic is added that has more complicated functionality than simple creation / deletion of records then it should have unit tests! 

However there are still e2e tests created for this project.

**Quickstart**:  
Run `cp .env_example .env && make run` in the `app/` directory to run the app.

### E2E Tests
There are a couple e2e tests written using playwright, in the `e2e/` directory. They can be ran with `npm run test`. However, this assumes you have the app running (on port 8080).

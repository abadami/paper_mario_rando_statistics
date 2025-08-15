## Paper Mario 64 Randomizer Race Statistics

A web server with a simple web page attached that collects data from racetime.gg (specifically paper mario 64 randomizer races) and gives specific race statistics and raw data for specific racers when queried.

### Quick Start

1. Clone this repository
2. Copy `~/.env.example` to `~/.env`.
3. Run `pnpm install` in the `~/client` directory for (you'll need [NodeJS](https://nodejs.org/en/download) installed and [pnpm](https://pnpm.io/installation) installed).
4. To run the server, simply run `go run ./app` in the root folder or `go run .` in the app folder (You will need [Golang](https://go.dev/doc/install) installed).
5. To run the webpage, simply run `pnpm dev` in the `~/client` directory.
6. (Optional) If you want to just run the server and have the server run locally, you can run `pnpm build` in the `~/client` folder. You can access the site via `http://localhost:3000`.

If you run both the webpage and the server separately, the server is accessible via `http://localhost:3000` and the webpage is accessible via `http://localhost:5173` when running.

### TODO List

[] - Proper handling of error states
[] - Better loading state
[] - Chart showing player stats over time
[] - Bar Chart for racer placement in community races
[] - W/L/T for league races
[] - See what we can do with data from the Paper Mario 64 Randomizer Spoiler Log

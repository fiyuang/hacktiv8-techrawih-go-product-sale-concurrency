<div align="center">
      <h1>Go Techrawih <i>(Hacktiv8 Workshop)</i><br>Optimize Version</h1>
     </div>


# Description
A simple import data demo to process and analyze >10,000 data in a CSV file with <b>goroutine, channel, sync.Waitgroup</b> implementation. There is a feature to process and analyze product sales data from an online store over a period of 1 year, and the feature is Import Sales Data. The goal of the analysis is to determine the total net sales and total gross sales for each product available.
 
# Tech Used
 ![GO](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
 ![GIN](https://img.shields.io/badge/Gin-3390d1?style=for-the-badge&logo=GIN&logoColor=white)
 ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
      
# Getting Start:
Before you running the program, make sure you've run this command:
- `cp .env.example .env`
- `go mod tidy`
- `make db-migrate-seed`

### Run the program
- `make dev`
- `make dev-race-condition` for build-in race detector

The program will run on http://localhost:8000

### API Route List
| Method | URL                                | Description |
| ----------- |------------------------------------| ----------- | 
| POST | localhost:8000/api/v1/sales/import | Import File CSV of Sales |
      
<!-- </> with 💛 by readMD (https://readmd.itsvg.in) -->
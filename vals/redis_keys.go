package vals

const ( // naming convention roughly: <service><type><name>
	ScraperRedisKeyLastRun    = "SCRAPER:LAST_RUN"
	ScraperRedisKeyInProgress = "SCRAPER:RUNNING"

	JanitorRedisKeyLastRun    = "JANITOR:LAST_RUN"
	JanitorRedisKeyInProgress = "JANITOR:RUNNING"
)

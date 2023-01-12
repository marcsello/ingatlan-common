package vals

const ( // naming convention roughly: <service><type><name>
	// -- Scraper --
	ScraperSourceIngatlanCom = "ING"
	ScraperSourceJofogasHu   = "JOF"

	ScraperRedisKeyLastRun = "SCRAPER:LAST_RUN"

	// -- Janitor --
	JanitorRedisKeyLastRun = "JANITOR:LAST_RUN"
)

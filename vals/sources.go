package vals

const ( // naming convention roughly: <service><type><name>
	ScraperSourceIngatlanCom = "ING"
	ScraperSourceJofogasHu   = "JOF"
)

func isScraperSourceValid(src string) bool {
	switch src {
	case ScraperSourceIngatlanCom:
	case ScraperSourceJofogasHu:
		return true
	}
	return false
}

func getValidSources() []string {
	return []string{ScraperSourceIngatlanCom, ScraperSourceJofogasHu}
}

func getValidSourcesWithUrlBase() map[string]string {
	return map[string]string{
		ScraperSourceIngatlanCom: "https://ingatlan.com/",
		ScraperSourceJofogasHu:   "https://ingatlan.jofogas.hu/",
	}
}

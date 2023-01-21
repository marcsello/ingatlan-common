package vals

const ( // naming convention roughly: <service><type><name>
	ScraperSourceIngatlanCom = "ING"
	ScraperSourceJofogasHu   = "JOF"
)

func IsScraperSourceValid(src string) bool {
	switch src {
	case ScraperSourceIngatlanCom:
	case ScraperSourceJofogasHu:
		return true
	}
	return false
}

func GetValidSources() []string {
	return []string{ScraperSourceIngatlanCom, ScraperSourceJofogasHu}
}

func GetValidSourcesWithUrlBase() map[string]string {
	return map[string]string{
		ScraperSourceIngatlanCom: "https://ingatlan.com/",
		ScraperSourceJofogasHu:   "https://ingatlan.jofogas.hu/",
	}
}

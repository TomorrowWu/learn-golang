package engine

import "log"

type SimpleEngine struct {
}

// Run starts crawling
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, e := Worker(r)
		if e != nil {
			continue
		}

		//The next level url to request
		for _, r := range parseResult.Requests {
			if isDuplicate(r.Url) {
				continue
			}
			requests = append(requests, r)
		}

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

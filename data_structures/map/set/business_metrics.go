package data_structures

// https://carloarg02.medium.com/my-favorite-coding-question-to-give-candidates-17ea4758880c

// Let’s say we have a website and we keep track of what pages customers are viewing, for things like business metrics.

// Every time somebody comes to the website, we write a record to a log file consisting of Timestamp, PageId, CustomerId.
// At the end of the day we have a big log file with many entries in that format. And for every day we have a new file.

// Now, given two log files (log file from day 1 and log file from day 2) we want to generate a list of ‘loyal customers’
// that meet the criteria of: (a) they came on both days, and (b) they visited at least two unique pages.

// Clarifying questions
// * unique pages per day or overall?
// * scale?

// Approaches
// 1. came on both days -> use a map on first day and then loop over second - costs O(2n) time and O(n) space
// 2. visited at least two unique pages -> use a map of set of pages O(n) space for map but then O(n) space for set O(2n) total
// we have a set of pages for the second day so need to consider that.

func business_metrics() {
	for file2.hasNext() {
		pageView := file2.next()
		customerId := pageView.customerId

		if pagesDay1, ok := day1Customers[customerId]; ok {
			// customer came on both days but did they visit at least two unique pages?
			if pagesDay1.size >= 2 || !pagesDay1[pageView.pageId] {
				loyalCustomers.add(customerId)
				// dont need to do additional work if see this customers page again
				remove(day1Customers, customerId)
			}
		}
	}
}

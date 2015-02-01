package testdiffman


func RunDiff() {

	/*
	Is blueprint flag set
		YES -> Read in API blueprint and compare
		NO -> Stop, what's the point in diffing against nothing?

	Execute request with json.

	Build expected JSON
	Iterate through fields
		Is it even in the API blueprint?
			No -> New info, add to report

		Does it fit the API blueprint response? (NOTE we can have regexes in the body that we need to check)
			NO -> Is the field REQUIRED?
				Yes -> Failed the diff, add to report
				No -> Missing optional info, add to report
			YES -> continue

	Construct simple CLI report
	If html flag is set -> Construct full beautiful html report and output
	 */
}

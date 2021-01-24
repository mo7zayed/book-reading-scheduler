# Schedule generator for reading a book (Challenge)
Senario:  A student required to finish a book of thirty chapters, he is allowed to choose when he starts, days he will be attending every week, and a starting date.

Inputs:
- Starting date: yyy-mm-dd
- Days: int array with a number of days per week assuming the start of the week is Saturday. Example: {2,4,6}
- Sessions Number: How many sessions required to finish one chapter. Example: {6}

The response will be a JSON representing the sessions scheduled for this student until he finishes the 30 chapters.

Example `[2019-08-01, 2019-09-1, ….]`

# run
- Import the postman collection
- `make serve` or `make dev` for development NOTE: if you are gonna run `make dev` see https://github.com/gravityblast/fresh first.

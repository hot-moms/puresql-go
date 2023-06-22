# puresql-go  #
***Tiny (pure bytes.Buffer) implementation of SQL Query Builder***

Download:
```go get github,com/hot-moms/puresql-go@latest```

Package *puresql-go* is a lightweight implementation of an SQL query builder written purely in Golang's bytes.Buffer.

This package provides an easy-to-use interface for building SQL queries programmatically, allowing developers to quickly and efficiently generate complex queries with minimal coding.

With puresql-go, building SQL queries is as simple as writing out the desired query in Golang code, making it an ideal solution for teams working with databases powered by SQL.

By utilizing Golang's powerful byte-based buffer system, puresql-go provides superior performance and precision in generating valid SQL scripts. Get started quickly and with ease by leveraging this versatile tool for all your SQL query building needs.



### Usage ###

#### Query:
```
const _sql = `SELECT hub FROM hubs`

	var (
		name            = "Github"
		companies       = []string{"Apple", "Microsoft", "Github"}
		types           = []string{"Commercial", "Non-commercial"}
		year_from int32 = 1990
		year_to   int32 = 2023
	)

	sB := sql.Init(_sql)
	sB.ContainedByStrings("company", companies)
	sB.BetweenInts("year", &year_from, &year_to)
	sB.EqualToString("name", &name)
	sB.AnyOfStrings("type", types)
	sB.OrderBy(nil, nil, "last_update DESC")
	sB.Limit(12)

	fmt.Println(sB.ToSQL())
```

### Result ###

__Query__:  `SELECT hub FROM hubs WHERE (company @> ($1)) AND (year BETWEEN $2 AND $3) AND (name = $4) AND (type = ANY($5)) ORDER BY last_update DESC LIMIT 12 `

__Arguments__: `[Apple Microsoft Github], 1990, 2023, Github, [Commercial Non-commercial]`

---

_2023, Archie Iwakura (hot-moms)_

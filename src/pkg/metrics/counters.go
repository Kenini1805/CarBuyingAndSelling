package metrics

var DbCall = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "db_calls_total",
		Help: "Number of database calls",
	}, []string{"type_name", "operation_name", "status"},
)

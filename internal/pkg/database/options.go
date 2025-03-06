package database

type Options func(ctx *pgContext)

func SetConfig(config *PostgresConfig) Options {
	return func(ctx *pgContext) {
		ctx.config = config
	}
}

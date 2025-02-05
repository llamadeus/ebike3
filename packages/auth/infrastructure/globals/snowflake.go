package globals

import "github.com/llamadeus/ebike3/packages/auth/infrastructure/utils"

var (
	_snowflake *utils.SnowflakeGenerator
)

func SetSnowflake(snowflake *utils.SnowflakeGenerator) {
	_snowflake = snowflake
}

func GetSnowflake() *utils.SnowflakeGenerator {
	return _snowflake
}

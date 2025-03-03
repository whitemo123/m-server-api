package snowflake

import "github.com/bwmarrin/snowflake"

func GenSnowflakeId() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0
	}
	// Generate a snowflake ID.
	id := node.Generate()
	return id.Int64()
}

package utils

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func PGTypeUUIDToString(myUUID pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", myUUID.Bytes[0:4], myUUID.Bytes[4:6], myUUID.Bytes[6:8], myUUID.Bytes[8:10], myUUID.Bytes[10:16])
}

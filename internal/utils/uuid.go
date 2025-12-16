package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func UUID() {
	id := uuid.New()
	fmt.Println(id.String())
}

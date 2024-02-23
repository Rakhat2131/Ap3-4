module architecture_go

go 1.18

require (
    github.com/lib/pq v1.10.2
)

replace (
    architecture_go/services/contact/internal/adapters => ./services/contact/internal/adapters
    architecture_go/services/contact/internal/delivery => ./services/contact/internal/delivery
    architecture_go/services/contact/internal/repository => ./services/contact/internal/repository
    architecture_go/services/contact/internal/usecase => ./services/contact/internal/usecase
    architecture_go/pkg/store/postgres => ./pkg/store/postgres
    architecture_go/services/contact/internal => ./services/contact/internal
)

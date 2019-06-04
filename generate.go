package shortener

//go:generate mockgen -source=internal/service/shortener/service.go -package=shortener -destination=internal/service/shortener/service_mock_test.go
//go:generate mockgen -source=internal/service/statistic/service.go -package=statistic -destination=internal/service/statistic/service_mock_test.go

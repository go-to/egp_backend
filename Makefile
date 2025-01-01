run:
	go run main.go
grpcurl-test:
	grpcurl -plaintext localhost:8080 egp.EgpService.GetShops
clean-branch:
	git switch main && git branch | xargs git branch -d

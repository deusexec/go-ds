test:
	@go test -count=1 \
	./queue \
	./stack \
	./graph

testv:
	@go test -v -count=1 \
	./queue \
	./stack \
	./graph

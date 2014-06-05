# Public Domain (-) 2014 The Espians Website Authors.
# See the Espians Website UNLICENSE file for details.

site: sync
	@assetgen
	@go run gensite.go

sync:
	@mkdir -p www
	@echo ">> Syncing with the espian-website-assets bucket on S3"
	@cd www && s3cmd sync -r --exclude '*_$$folder$$' s3://espians-website-assets .

clean:
	@assetgen --clean
	@rm -rf www

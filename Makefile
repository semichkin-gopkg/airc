image_version := v0.0.11

build:
	docker buildx build \
		--platform linux/amd64,linux/arm64,linux/arm/v7 \
		-t semichkin/airc:${image_version} \
		-t semichkin/airc:latest ./docker

push:
	docker buildx build \
		--platform linux/amd64,linux/arm64,linux/arm/v7 \
		-t semichkin/airc:${image_version} \
		-t semichkin/airc:latest \
		--push ./docker

init-builder:
	docker run --privileged --rm tonistiigi/binfmt --install all && \
	docker buildx create --name airc && \
	docker buildx use airc && \
	docker buildx inspect --bootstrap

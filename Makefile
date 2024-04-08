font:
	/home/quasilyte/CODE/bitfontier/bitfontier \
		--on-missing=stub \
		--data-dir=_data \
		--out-dir=_build \
		--pkgname=bitsweetfont \
		--generate-info \
		-v
	@ cp _build/* .
	@ rm -rf _build
	echo "Build completed!"

.PHONY: font

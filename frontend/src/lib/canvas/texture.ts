const imageCache: Record<number, HTMLImageElement> = {};

export function getTexture(index: number) {
	if (imageCache[index]) {
		return imageCache[index];
	}

	const img = new Image();
	img.src = "/assets/blocks/" + index + ".png";

	imageCache[index] = img;
	return img;
}

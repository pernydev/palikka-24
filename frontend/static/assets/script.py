import os
import operator
import struct
import array
import colorsys

def get_average_hue_saturation(image_path):
    with open(image_path, 'rb') as f:
        f.read(8)  # Skip header
        width, height = struct.unpack('>II', f.read(8))
        f.read(4)  # Skip png info

        rgb_values = array.array('B')
        rgb_values.fromfile(f, width * height * 3)

        hue_saturation_sum = [0, 0]
        for x in range(width):
            for y in range(height):
                r = rgb_values[3 * (x + y * width)]
                g = rgb_values[3 * (x + y * width) + 1]
                b = rgb_values[3 * (x + y * width) + 2]
                hue, saturation, value = colorsys.rgb_to_hsv(r / 255.0, g / 255.0, b / 255.0)
                hue_saturation_sum[0] += hue
                hue_saturation_sum[1] += saturation

        average_hue = hue_saturation_sum[0] / (width * height)
        average_saturation = hue_saturation_sum[1] / (width * height)

        return average_hue, average_saturation

def sort_images_by_hue_saturation(folder_path):
    image_paths = [os.path.join(folder_path, filename) for filename in os.listdir(folder_path) if filename.endswith('.png')]
    image_hue_saturation = {}

    for image_path in image_paths:
        average_hue, average_saturation = get_average_hue_saturation(image_path)
        image_hue_saturation[image_path] = (average_hue, average_saturation)

    sorted_images = sorted(image_hue_saturation.items(), key=operator.itemgetter(1))

    return sorted_images

folder_path = 'blocks'
sorted_images = sort_images_by_hue_saturation(folder_path)

for i, (image_path, hue_saturation) in enumerate(sorted_images, start=1):
    filename = os.path.basename(image_path)
    new_name = f'{i}.png'
    new_path = os.path.join(folder_path, new_name)
    os.rename(image_path, new_path)
    print(f'Renamed {filename} to {new_name}')

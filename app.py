from flask import Flask, render_template
import matplotlib.pyplot as plt
import numpy as np

from skimage import data, io, img_as_ubyte
from skimage.filters import threshold_multiotsu
from skimage.color import rgb2gray
import cv2

def detect_tumor(image):
    image = rgb2gray(image)
    thresholds = threshold_multiotsu(image, classes=5)
    regions = np.digitize(image, bins=thresholds)
    cv2.imwrite('res.jpg',regions)
    return
    output = img_as_ubyte(regions)
    fig, ax = plt.subplots(nrows=1, ncols=3, figsize=(10, 3.5))

    # Plotting the original image.
    ax[0].imshow(image, cmap='gray')
    ax[0].set_title('Original')
    ax[0].axis('off')

    # Plotting the histogram and the two thresholds obtained from
    # multi-Otsu.
    ax[1].hist(image.ravel(), bins=255)
    ax[1].set_title('Histogram')
    for thresh in thresholds:
        ax[1].axvline(thresh, color='r')

    # Plotting the Multi Otsu result.
    ax[2].imshow(regions, cmap='Accent')
    ax[2].set_title('Multi-Otsu result')
    ax[2].axis('off')

    plt.subplots_adjust()

    plt.show()
    
im = cv2.imread('test.jpg');
detect_tumor(im)



# app = Flask(__name__)

# @app.route('/')
# def index():
#     return render_template('index.html')


# if __name__== "__main__":
#     app.run(debug=True)


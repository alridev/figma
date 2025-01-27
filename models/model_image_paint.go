package models

// ImagePaint struct for ImagePaint
type ImagePaint struct {
	// Is the paint enabled?
	Visible bool `json:"visible,omitempty"`
	// Overall opacity of paint (colors within the paint can also have opacity values which would blend with this)
	Opacity float64 `json:"opacity,omitempty"`
	// How this node blends with nodes behind it in the scene
	BlendMode BlendMode `json:"blendMode"`
	// The string literal \"IMAGE\" representing the paint's type. Always check the `type` before reading other properties.
	Type string `json:"type"`
	// Image scaling mode.
	ScaleMode string `json:"scaleMode"`
	// A reference to an image embedded in this node. To download the image using this reference, use the `GET file images` endpoint to retrieve the mapping from image references to image URLs.
	ImageRef string `json:"imageRef"`
	// A transformation matrix is standard way in computer graphics to represent translation and rotation. These are the top two rows of a 3x3 matrix. The bottom row of the matrix is assumed to be [0, 0, 1]. This is known as an affine transform and is enough to represent translation, rotation, and skew.  The identity transform is [[1, 0, 0], [0, 1, 0]].  A translation matrix will typically look like:  ``` [[1, 0, tx],   [0, 1, ty]] ```  and a rotation matrix will typically look like:  ``` [[cos(angle), sin(angle), 0],   [-sin(angle), cos(angle), 0]] ```  Another way to think about this transform is as three vectors:  - The x axis (t[0][0], t[1][0]) - The y axis (t[0][1], t[1][1]) - The translation offset (t[0][2], t[1][2])  The most common usage of the Transform matrix is the `relativeTransform property`. This particular usage of the matrix has a few additional restrictions. The translation offset can take on any value but we do enforce that the axis vectors are unit vectors (i.e. have length 1). The axes are not required to be at 90° angles to each other.
	ImageTransform [][]float64 `json:"imageTransform,omitempty"`
	// Amount image is scaled by in tiling, only present if scaleMode is `TILE`.
	ScalingFactor float64 `json:"scalingFactor,omitempty"`
	// Defines what image filters have been applied to this paint, if any. If this property is not defined, no filters have been applied.
	Filters *ImageFilters `json:"filters,omitempty"`
	// Image rotation, in degrees.
	Rotation float64 `json:"rotation,omitempty"`
	// A reference to an animated GIF embedded in this node. To download the image using this reference, use the `GET file images` endpoint to retrieve the mapping from image references to image URLs.
	GifRef string `json:"gifRef,omitempty"`
}

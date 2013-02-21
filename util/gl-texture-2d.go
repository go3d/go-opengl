package ugl

import (
	"image"
	"math"

	gl "github.com/go3d/go-opengl/core"
)

//	Represents a 2-dimensional texture image.
type Texture2D struct {
	//	Common texture params
	TextureBase

	//	Width and height of mip-level 0 for this 2-dimensional texture object.
	Width, Height gl.Sizei
}

//	Initializes --but does not Recreate()-- a new Texture2D, calls its Init() method and returns it.
func NewTexture2D() (me *Texture2D) {
	me = &Texture2D{}
	me.Init()
	return
}

//	Sets GlTarget to gl.TEXTURE_2D and initializes TextureBase with defaults.
func (me *Texture2D) Init() {
	me.GlTarget = gl.TEXTURE_2D
	me.TextureBase.init()
}

//	Returns the maximum number of possible MIP map levels for this 2-dimensional texture image according to its current Width and Height.
func (me *Texture2D) MaxNumMipLevels() gl.Sizei {
	return gl.Sizei(texture2DMaxNumMipLevels(float64(me.Width), float64(me.Height)))
}

//	Prepares this Texture2D for uploading the specified Image via Recreate().
//	This sets all of the following fields to applicable values:
//	me.PixelData.Type, me.PixelData.Format, me.PixelData.Ptrs[0], me.Width, me.Height, me.MipMap.NumLevels, me.SizedInternalFormat
func (me *Texture2D) PrepFromImage(bgra, uintRev bool, img image.Image) (err error) {
	me.Width, me.Height = gl.Sizei(img.Bounds().Dx()), gl.Sizei(img.Bounds().Dy())
	me.MipMap.NumLevels = me.MaxNumMipLevels()
	err = me.prepFromImages(bgra, uintRev, img)
	return
}

//	Deletes and (re)creates the texture object based on its current params.
//	Uploads the image data at me.PixelData.Ptrs[0], if any.
//	Generates the MIP map if me.MipMap.AutoGen is true and me.MipMap.NumLevels isn't 1.
//	If me.MipMap.NumLevels is 0 (or just smaller than 1), then me.MaxNumMipLevels() is used.
func (me *Texture2D) Recreate() (err error) {
	err = me.onBeforeRecreate()
	defer me.onAfterRecreate()
	if err == nil {
		hasPixData, numLevels := me.PixelData.Ptrs[0] != gl.Ptr(nil), me.MipMap.NumLevels
		if numLevels < 1 {
			numLevels = me.MaxNumMipLevels()
		}
		if me.immutable() {
			if err = Try.TexStorage2D(me.GlTarget, numLevels, me.SizedInternalFormat, me.Width, me.Height); hasPixData && (err == nil) {
				err = Try.TexSubImage2D(me.GlTarget, 0, 0, 0, me.Width, me.Height, me.PixelData.Format, me.PixelData.Type, me.PixelData.Ptrs[0])
			}
		} else {
			err = Try.TexImage2D(me.GlTarget, 0, gl.Int(me.SizedInternalFormat), me.Width, me.Height, 0, me.PixelData.Format, me.PixelData.Type, me.PixelData.Ptrs[0])
		}
		if (err == nil) && hasPixData && me.MipMap.AutoGen {
			err = Try.GenerateMipmap(me.GlTarget)
		}
	}
	return
}

//	Returns the maximum number of possible MIP map levels for a 2-dimensional texture image with the specified width and height.
func texture2DMaxNumMipLevels(width, height float64) float64 {
	return math.Log2(math.Max(width, height)) + 1
}

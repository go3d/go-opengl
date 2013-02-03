go-opengl
=========

Go packages for working with OpenGL:


core:
=====


An OpenGL binding inspired by and 90% similar to [GoGL](https://github.com/chsc/gogl) but generated by [my own hacky binding-generator](https://github.com/go3d/go-opengl/tree/master/cmd/gen-opengl-bindings) and fine-tuned to the needs of [go:ngine](http://github.com/go3d/go-ngine).

(Better to call it a blatant rip-off as a little hacking exercise to grok what's going on behind the curtains in a CGO binding and to be able to customize it slightly.)

If you need a general OpenGL binding, you may first want to try [the pre-generated GoGL bindings](https://github.com/chsc/gogl) or GoGL's binding-generator -- because this go3d/go-opengl/core binding will always be fine-tuned to the needs of go:ngine and its users first and foremost, rather than being a general-use OpenGL binding.

Specifically, this binding:

- includes strictly only GL core profile functionality from version 3.3 onwards (up until 4.3) -- note, this does not mean that a core profile context is *required*, although it is probably advisable on any current-gen GPU with recent drivers
- no compatibility-profile functionality or any that was deprecated or removed in 3.3 or before
- Init() and stolen-from-GoGL utility functions moved to a *Util* struct: so every exported (non-method) function is a direct CGO binding function. (Just minor cosmetics really.)
- added to *Util*: EnumName(), ErrorFlags(), Error() methods
- runtime checking for individual function availability via a *Supports* struct
- function wrappers that check for GL errors and return Go errors via a *Try* struct
- the above items add massively to the generated binding pkg (core.a is 7.8MB whereas GoGL's is only 2MB), but hey, the Go linker strips unused code anyway.
- embeds the following GL extensions right inside the same binding package:
	- EXT_texture_filter_anisotropic
	- EXT_texture_compression_s3tc
	- EXT_texture_sRGB


util:
=====


Based on the above **core** GL binding, provides higher-level utility constructs and a few sane slim wrappers (without going full-on "OO"-overboard) for Go.



cmd:
====


- **gen-opengl-bindings**: can be used to generate  OpenGL bindings like the above **core** package
- **opengl-minimal-app**: a "minimal" program using OpenGL core profile to draw a triangle and a quad, to test the above binding with GLFW
- **gogl-minimal-app**: same but with the GoGL (gl43) binding

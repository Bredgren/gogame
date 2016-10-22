package gogame

// import (
// 	"image/color"

// 	"github.com/Bredgren/gogame/composite"
// 	"github.com/Bredgren/gogame/geo"
// )

// // Surface represents an image or drawable surface.
// type Surface interface {
// 	// Blit draws the given surface to this one at the given position. Source's top-left corner
// 	// (according to Surface.Rect() fill be drawn at (x, y).
// 	Blit(source Surface, x, y float64)
// 	// BlitArea draws the given portion of the source surface defined by the Rect to this
// 	// one with its top-left corner (according to Surface.Rect()) at the given position.
// 	BlitArea(source Surface, area geo.Rect, x, y float64)
// 	// BlitComp is like Blit but takes a composite operation to use.
// 	BlitComp(source Surface, x, y float64, c composite.Op)
// 	// BlitAreaComp is like BlitArea but takes a composite operation to use.
// 	BlitAreaComp(source Surface, area geo.Rect, x, y float64, c composite.Op)
// 	// Fill fills the whole surface with the given style.
// 	Fill(*FillStyle)
// 	// Width returns the unrotated, unscaled width of the surface in pixels. To get the width
// 	// after scaling and rotating use Surface.Rect.
// 	Width() int
// 	// Height returns the unrotated, unscaled height of the surface in pixels. To get the height
// 	// after scaling and rotating use Surface.Rect.
// 	Height() int
// 	// Copy returns a new Surface that is identical to this one.
// 	Copy() Surface
// 	// SubSurface creates a surface that is a reference to a sub area, defined by the given rect,
// 	// of this surface. Draw operations to either surface affect the other. Note that calling
// 	// Canvas() on a sub-surface returns the parent's canvas.
// 	SubSurface(geo.Rect) Surface
// 	// Parent returns the parent surface if this one is a sub-surface, otherwise nil.
// 	Parent() Surface
// 	// GetAt returns the color of the pixel at (x, y).
// 	GetAt(x, y int) color.Color
// 	// SetAt sets the color of the pixel at (x, y).
// 	SetAt(x, y int, c color.Color)
// 	// SetClip defines a rectangular region of the surface where only the enclosed pixels can
// 	// be affected by draw operations. Note that a parent surface and a subsurface cannot be
// 	// clipped at the same time.
// 	SetClip(geo.Rect)
// 	// SetClipPath defines an arbitrary polygon where only the enclosed pixels can be affected
// 	// by draw operations. The pointList is a list of xy-coordinates.
// 	SetClipPath(pointList [][2]float64)
// 	// ClearClip resets the clipping region to the whole canvas.
// 	ClearClip()
// 	// Scaled returns a new Surface that is equivalent to this one scaled by the given amount.
// 	Scaled(x, y float64) Surface
// 	// Rotated returns a new Surface that is equivalent to this one but rotated counter-clockwise
// 	// by the given amount.
// 	Rotated(radians float64) Surface
// 	// Rect returns the bouding rectangle for the surface, taking into acount rotation and scale.
// 	Rect() geo.Rect
// 	// DrawRect draws a rectangle on the surface.
// 	DrawRect(geo.Rect, Styler)
// 	// DrawCircle draws a circle on the surface.
// 	DrawCircle(posX, posY, radius float64, s Styler)
// 	// DrawEllipse draws an ellipse on the canvas within the given Rect.
// 	DrawEllipse(geo.Rect, Styler)
// 	// DrawArc draws an arc on the canvas within the given Rect between the given angles.
// 	// Angles are counter-clockwise.
// 	DrawArc(r geo.Rect, startRadians, stopRadians float64, s Styler)
// 	// DrawLine draws a line on the surface.
// 	DrawLine(startX, startY, endX, endY float64, s Styler)
// 	// DrawLines draws multiple connected lines to the surface.
// 	DrawLines(points [][2]float64, s Styler)
// 	// DrawText draws the given text to the surface.
// 	DrawText(text string, x, y float64, font *Font, style *TextStyle)
// 	// DrawQuadraticCurve draws a quadratic curve to the surface from (startX, startY) to
// 	// (endX, endY) with the control point (cpX, cpY).
// 	DrawQuadraticCurve(startX, startY, endX, endY, cpX, cpY float64, s Styler)
// 	// DrawQuadraticCurves draws multiple connected quadratic curves to the surface. The list
// 	// of points alternaes between end points and control points, i.e. [startpoint, control1,
// 	// point2, control2, point3, etc.]. If the list has an even number of points then the last
// 	// point is ignored. If there are less than 3 points then nothing will drawn.
// 	DrawQuadraticCurves(points [][2]float64, s Styler)
// 	// DrawBezierCurve draws a cubic bezier curve to the surface from (startX, startY) to
// 	// (endX, endY) with the given respective control points..
// 	DrawBezierCurve(startX, startY, endX, endY, cpStartX, cpStartY, cpEndX, cpEndY float64, s Styler)
// 	// DrawBezierCurves draws multiple connected cubic quadratic curves to the surface. The list
// 	// of points alternaes between end points and both control points, i.e. [startpoint, control1,
// 	// control2, point2, control3, control4, point3, etc.]. If there are not enough points at
// 	// the end of the list then it will stop at the last possible point. If there are less than
// 	// 4 points then nothing will drawn.
// 	DrawBezierCurves(points [][2]float64, s Styler)
// }

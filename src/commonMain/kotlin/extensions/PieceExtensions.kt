package extensions

import com.soywiz.kds.getExtra
import com.soywiz.kds.setExtra
import com.soywiz.korma.geom.Point
import containers.Piece

var Piece.lastBoardPosition
	get() = this.getExtra("lastBoardPosition") as? Point?
	set(value)
	{
		this.setExtra("lastBoardPosition", value)
	}
package extensions

import com.soywiz.klogger.Logger

val Any.logger get() = Logger(this::class.simpleName!!)
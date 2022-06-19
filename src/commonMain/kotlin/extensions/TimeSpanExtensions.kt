package extensions

import com.soywiz.klock.PerformanceCounter
import com.soywiz.klock.TimeSpan

fun TimeSpan.Companion.now() = TimeSpan(PerformanceCounter.milliseconds)
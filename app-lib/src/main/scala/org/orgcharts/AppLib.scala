package org.orgcharts

import cats.effect.IO
import scala.scalajs.js
import scala.scalajs.js.annotation._
import cats.effect.unsafe.implicits.global

object AppLib:
  @JSExportTopLevel("sampleTask")
  def sampleTask(): js.Promise[Int] =
    val task =
      IO.println("Hello!") *>
        IO(100)

    task.unsafeRunSyncToPromise()

package org.orgcharts

import cats.effect.IO
import cats.effect.unsafe.implicits.global
import io.circe.Codec
import org.http4s.ember.client.EmberClientBuilder
import org.http4s.{Uri, Method, Request}
import org.http4s.circe.jsonOf
import org.http4s.circe.CirceEntityCodec._

import scala.scalajs.js
import scala.scalajs.js.annotation._

object AppLib:
  val clientResource = EmberClientBuilder
    .default[IO]
    .build

  @JSExportTopLevel("sampleTask")
  def sampleTask(): js.Promise[Int] =
    val task = for
      _ <- IO.println("Hello!")
      num <- IO(100)
    yield num

    task.unsafeToPromise()

  final case class InitializeUploadRequest(
    id: String,
    name: String,
    parts: Int
  ) derives Codec

  final case class InitializeUploadResponse() derives Codec

  @JSExportTopLevel("initializeUpload")
  def initializeUpload(
    urlBase: String,
    id: String,
    name: String,
    parts: Int
  ): js.Promise[Unit] =
    val task = for
      _ <- IO.println(
        ">>> PROGRAM >>> INITIALIZING UPLOAD >>> " +
        s"ID ... $id >>> NAME ... $name >>> PARTS ... $parts"
      )
      _ <- clientResource.use(client =>
        client.expect[InitializeUploadResponse](
          Request[IO](
            Method.POST,
            Uri.unsafeFromString(urlBase) / "initialize-upload"
          ).withEntity(
            InitializeUploadRequest(
              id,
              name,
              parts
            )
          )
        )
      )
    yield ()

    task.unsafeToPromise()
